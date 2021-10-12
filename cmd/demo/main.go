package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := registerCtrlC(context.Background())

	err := runThreads(ctx)

	log.Printf("Threads exited.  Error: %v", err)
}

// registerCtrlC returns a context that is canceled when the user CTRL-C's the program
func registerCtrlC(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}

// runThreads runs all tasks in separate goroutines.  If an error is encountered, all routines are stopped, and the
// first error encountered by any of them is returned.
func runThreads(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(repeater(ctx, 1*time.Second, func() {
		log.Printf("Doing task a")
	}))

	return g.Wait()
}

// repeater takes in a function and runs it continually with the given delay, until the given context is closed.
func repeater(ctx context.Context, delay time.Duration, f func()) func() error {
	return func() error {
		t := time.NewTicker(delay)
		for {
			select {
			case <-t.C:
				f() // Do the thing
			case <-ctx.Done():
				t.Stop()
				return nil
			}
		}
	}
}

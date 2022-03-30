package main

import (
	"context"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := registerCtrlC(context.Background())
	registerMetrics()

	log.Printf("Starting threads...\n")

	err := runThreads(ctx)

	log.Printf("Threads exited.  Error: %v\n", err)
}

// registerCtrlC returns a context that is canceled when the user CTRL-C's the program
func registerCtrlC(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Printf("Gracefully shutting down the gobin service\n")
		cancel()
	}()

	return ctx
}

// runThreads runs all tasks in separate goroutines.  If an error is encountered, all routines are stopped, and the
// first error encountered by any of them is returned.
func runThreads(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// Start the metric server
	g.Go(func() error { return runMetricsServer(ctx) })

	constDelay := func(d time.Duration) func() time.Duration {
		return func() time.Duration {
			return d
		}
	}
	randDelay := func(mean, stddev float64) func() time.Duration {
		return func() time.Duration {
			return time.Duration(rand.NormFloat64()*stddev+mean) * time.Second
		}
	}

	// Counter (goes up by 1 ever second)
	g.Go(repeatFunc(ctx, randDelay(1, 1), func() {
		log.Printf("counter: increment\n")

		metricCounter.Inc()
		// functions:
		// .Inc()  // +1
		// .Add(n) // +n (n _must_ be positive)
	}))

	// Gauge (new value every second, somewhere between 0-10)
	g.Go(repeatFunc(ctx, constDelay(1*time.Second), func() {
		val := rand.Float64() * 10
		log.Printf("gauger: %v\n", val)

		metricGauge.Set(val)
		// Functions:
		// .Set(n) // Set's current value to n
		// .Inc()  // +1
		// .Dec()  // -1
		// .Add(n) // +n (n can be negative)
		// .Sub(n) // -n
	}))

	// Histogram (new value roughly every 1s)
	g.Go(repeatFunc(ctx, randDelay(1, 1), func() {
		// *1000 to convert to ms
		// /3 to give us 300ms as 1stddev
		val := math.Abs(rand.NormFloat64() * 1000 / 3)
		log.Printf("histogram: %v", val)

		metricHistogram.Observe(val)
		// Functions:
		// .Observe(n) // n will get put into the correct bucket for you.
	}))

	return g.Wait()
}

// repeatFunc takes in a function and runs it continually with the given delay, until the given context is closed.
func repeatFunc(ctx context.Context, delayCalc func() time.Duration, f func()) func() error {
	return func() error {
		for {
			select {
			case <-time.After(delayCalc()):
				f() // Do the thing
			case <-ctx.Done():
				return nil
			}
		}
	}
}

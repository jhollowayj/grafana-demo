bin/demo:
	GOOS=linux go build -o bin/demo cmd/demo/main.go

build: bin/demo
	docker build --rm -t jhollowayj/grafana-demo .

run:
	docker run --rm jhollowayj/grafana-demo

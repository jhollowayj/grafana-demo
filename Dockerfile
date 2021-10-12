# Build image
FROM golang:alpine AS builder

WORKDIR /app
ADD . /app
RUN go build -o /bin/demo /app/cmd/demo

# Deploy image

FROM alpine

COPY --from=builder /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]

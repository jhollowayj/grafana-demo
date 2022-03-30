# Build image
FROM golang:alpine AS builder

WORKDIR /app
ADD . /app

# Note: CGO_ENABLED=0 needed for deploying in scratch images
RUN CGO_ENABLED=0 go build -o /bin/demo /app/cmd/demo
RUN CGO_ENABLED=0 go build -o /bin/promdemo /app/cmd/prometheusexample


# Deploy image

FROM scratch

COPY --from=builder /bin/demo /bin/demo
COPY --from=builder /bin/promdemo /bin/promdemo
ENTRYPOINT ["/bin/demo"]

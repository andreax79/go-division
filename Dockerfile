FROM golang:1.19.5-buster AS builder

# Build the application
WORKDIR /build
COPY . .
RUN go build -v -o /build/division -ldflags "-X main.version=$(cat VERSION)"
RUN go test
WORKDIR /dist
RUN cp /build/division ./division

# Create the runtime image
FROM scratch
COPY --chown=0:0 --from=builder /dist/division /division
COPY --chown=0:0 LICENSE /LICENSE
USER 65534
WORKDIR /

LABEL org.opencontainers.image.source=https://github.com/andreax79/go-division/
LABEL org.opencontainers.image.description="Column Division"
LABEL org.opencontainers.image.licenses="Apache-2.0"

ENTRYPOINT ["/division"]

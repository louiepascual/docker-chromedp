FROM golang:1.21.6-bookworm as build

WORKDIR /app

COPY go.* ./
RUN go mod download -x all

COPY main.go .

# CGO_ENABLED=0 disables calling C code (import "c")
#   + since the final image is based on bullseye
#   there are missing glibc packages used in bookworm
#RUN go env -w CGO_ENABLED=0

RUN go build -o docker-chromedp -v main.go

# arm64 fork of chromedp/headless-shell
FROM louiepascual/headless-shell-arm64:121.0.6167.160 AS final

WORKDIR /app
COPY --from=build /app/docker-chromedp ./
ENTRYPOINT ["/app/docker-chromedp"]

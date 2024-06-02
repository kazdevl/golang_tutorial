FROM golang:1.22.2 as build

WORKDIR /workdir
COPY go.mod go.sum /workdir/
RUN go mod download

COPY . /workdir
RUN go build -ldflags "-s -w" -o app ./context/value/.

FROM debian:bookworm-slim
COPY --from=build /workdir/app /app
ENTRYPOINT ["/app"]
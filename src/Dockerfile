FROM golang:1.22.2 as build

WORKDIR /workdir
RUN go env -w GOMODCACHE=/root/.cache/go-build
COPY go.mod go.sum /workdir/
RUN --mount=type=cache,target=/root/.cache/go-build go mod download

COPY . /workdir
RUN --mount=type=cache,target=/root/.cache/go-build go build -ldflags "-s -w" -o app ./cicd/cache/heavy/.

FROM debian:bookworm-slim
COPY --from=build /workdir/app /app
ENTRYPOINT ["/app"]
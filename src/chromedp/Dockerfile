FROM golang:1.22 as builder
WORKDIR /src

COPY go.mod go.sum chromedp/test_server/main.go ./
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o app main.go

FROM scratch
WORKDIR /src
COPY --from=builder /src/app .
CMD ["./app"]
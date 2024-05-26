FROM golang:1.22.3-alpine3.19 as build

WORKDIR /app
COPY . /app

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app -ldflags="-extldflags=-static" ./cmd/cli/main.go

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=build /app/app /app/app

ENTRYPOINT [ "/app/app" ]

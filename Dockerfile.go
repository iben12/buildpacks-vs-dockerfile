FROM golang:1.17 as build

WORKDIR /app

COPY go .

RUN CGO_ENABLED=0 go build -o ./bin/main main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/bin/main .

CMD ["/app/main"]
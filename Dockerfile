FROM golang:1.20.4-alpine as build-stage

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.18

COPY --from=build-stage /app/main .
COPY --from=build-stage /app/.env .

EXPOSE 8080

CMD ["./main"]

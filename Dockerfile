FROM golang:1.23-alpine AS builder

WORKDIR /addons-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /app main.go

FROM gcr.io/distroless/static-debian12

COPY --chown=nonroot:nonroot --from=builder /app /app

USER nonroot

EXPOSE 3000 3001

ENTRYPOINT ["/app"]
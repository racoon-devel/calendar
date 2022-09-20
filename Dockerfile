FROM golang as builder
WORKDIR /go/src/github.com/racoon-devel/calendar
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o calendar cmd/calendar/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/racoon-devel/calendar/calendar .
CMD ["./calendar"]
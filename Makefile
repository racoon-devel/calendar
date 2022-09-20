all: calendar test

calendar:
	go build -o calendar.out cmd/calendar/main.go

clean:
	go clean
	rm -f calendar.out

test:
	go test
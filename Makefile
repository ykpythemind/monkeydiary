test:
		go test -v

run:
		go run .

deps:
		go mod verify

cron: deps run

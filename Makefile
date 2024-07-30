.PHONY: clean build run cov

build: 
	go test ./...
	go build -o shimmy main.go

run: 
	go run main.go

clean:
	rm -f shimmy nohup.out ./mover/mover.o ./mover/mover_program 
	go clean

cov:
	go test -coverpkg=./... ./...
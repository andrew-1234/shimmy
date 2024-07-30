.PHONY: clean build run

all: build

mover.o: mover.c mover.h
	gcc -c -o mover.o mover.c

build: mover.o
	go build -o mover_program mover.go

run: mover.o
	go run mover.go

clean:
	rm -f mover.o mover_program
	go clean
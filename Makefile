build:
	mkdir -p bin
	go build -o bin/dazy

run: build
	./bin/dazy
	

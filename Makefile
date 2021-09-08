


clean: 
	rm -rf build

test: 
	go test ./...

build:
	go build -o build/uniq main.go 





clean: 
	rm -rf build


test-coverage:
	go test -coverprofile=coverage.out.tmp ./...
	go tool cover -html=coverage.out.tmp

test: 
	go test ./...

build:
	go build -o build/uniq main.go 


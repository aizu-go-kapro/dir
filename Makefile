clean: 
	rm -rf ./cronl
	rm -rf ./cronr

test: 
	go test -v 

build:
	go build -o cronl ./cmd/local/main.go   #学内用 
	go build -o cronr ./cmd/remote/main.go  #リモート用

run: 
	go run ./cmd/local/main.go
	go run ./cmd/remote/main.go	

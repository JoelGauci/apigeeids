#Makefile for apigeeids module
build:
	@echo "build apigeeids"
	@go build -o bin/apigeeids main.go

run:
	@echo "run apigeeids\n-----\n"
	@go run main.go

clean:
	@echo "clean all apigeeids binary files"
	@rm bin/*

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=386 go build -o bin/apigeeids-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/apigeeids-linux-386 main.go
	GOOS=linux GOARCH=arm go build -o bin/apigeeids-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/apigeeids-linux-arm64 main.go
	GOOS=linux GOARCH=ppc64 go build -o bin/apigeeids-linux-ppc64 main.go
	GOOS=linux GOARCH=ppc64le go build -o bin/apigeeids-linux-ppc64le main.go
	GOOS=linux GOARCH=mips go build -o bin/apigeeids-linux-mips main.go
	GOOS=linux GOARCH=mipsle go build -o bin/apigeeids-linux-mipsle main.go
	GOOS=linux GOARCH=mips64le go build -o bin/apigeeids-linux-mips64le main.go

	GOOS=freebsd GOARCH=386 go build -o bin/apigeeids-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/apigeeids-freebsd-amd64 main.go
	GOOS=freebsd GOARCH=arm go build -o bin/apigeeids-freebsd-arm main.go

	GOOS=darwin GOARCH=amd64 go build -o bin/apigeeids-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/apigeeids-darwin-arm64 main.go

	GOOS=windows GOARCH=386 go build -o bin/apigeeids-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/apigeeids-windows-amd64 main.go

.PHONY: build clean deploy

test: build
	go test ./... -cover -count 1 -coverprofile=coverage.out

build:
	GOARCH=amd64 GOOS=linux go build -gcflags="-N -l" -o bin/handler handler/handler.go
	GOARCH=amd64 GOOS=windows go build -gcflags="-N -l" -o devbin/server.exe server/server.go
	GOARCH=amd64 GOOS=linux go build -gcflags="-N -l" -o devbin/server server/server.go

	if [ -a .serverless/jiraTagger.zip ]; then rm -rf .serverless/jiraTagger.zip; fi;
	mkdir -p .serverless
	zip .serverless/jiraTagger.zip bin/*
	
clean:
	rm -rf ./bin ./vendor Gopkg.lock ./devbin

deploy: clean build
	sls deploy --verbosed 

slsapi: clean build
	sls offline --useDocker --host 192.168.0.15

api: clean build
	./devbin/server

undeploy:
	sls undeploy --verbosed

.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	GOARCH=amd64 GOOS=linux go build -gcflags="-N -l" -o bin/receive receive/main.go
	if [ -a .serverless/jiraTagger.zip ]; then rm -rf .serverless/jiraTagger.zip; fi;
	mkdir -p .serverless
	zip .serverless/jiraTagger.zip bin/*
	
buildServer:
	export GO111MODULE=on
	GOARCH=amd64 GOOS=windows go build -gcflags="-N -l" -o bin/server server/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbosed 

slsapi: clean build
	sls offline --useDocker --host 192.168.0.15

api: clean buildServer
	./bin/server

undeploy:
	sls undeploy --verbosed

test:
	go test ./... -cover -count 1
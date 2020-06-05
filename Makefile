.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	GOARCH=amd64 GOOS=linux go build -gcflags="-N -l" -o bin/receive receive/main.go
	if [ -a .serverless/jiraTagger.zip ]; then rm -rf .serverless/jiraTagger.zip; fi;
	mkdir -p .serverless
	zip .serverless/jiraTagger.zip bin/*

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbosed 

api: clean build
	sls offline --useDocker --host 192.168.0.15

undeploy:
	sls undeploy --verbosed

test:
	go test ./... -cover -count 1
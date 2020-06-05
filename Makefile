.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	GOARCH=amd64 GOOS=linux go build -gcflags="-N -l" -o bin/receive receive/main.go
### yeoman hook ###
## Don't touch this comment, the subgenerator needs it'
	if [ -a .serverless/jiraTagger.zip ]; then rm -rf .serverless/jiraTagger.zip; fi;
	mkdir -p .serverless
	zip .serverless/jiraTagger.zip bin/*

access:
	chmod -R u+x ./scripts/

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbosed

sam: build
	sls sam export --output template.yaml

local-api: sam
	./scripts/samapi.sh '' '$(network)'

api: local-api

debug-api: sam
	./scripts/samapi.sh 'debug' '$(network)'

local-invoke: clean build access
	./scripts/localinvoke.sh '$(func)' '$(event)' '$(network)'

sam-invoke: sam
	./scripts/samlocal.sh '$(func)' '$(event)' '$(network)' 'export'

debug: sam access
	./scripts/samdebug.sh '$(func)' '$(event)' '$(network)' 'export'

sam-debug: sam access
	./scripts/samdebug.sh '$(func)' '$(event)' '$(network)' 'export'

undeploy:
	sls undeploy --verbosed

test: gomodgen
	go test ./... -cover -count 1
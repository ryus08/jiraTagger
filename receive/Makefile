.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o ../bin/receive main.go

clean:
	rm -rf ../bin/receive ../vendor ../Gopkg.lock

deploy: clean build
	cd ..
	sls deploy --function receive --verbose


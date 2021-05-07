.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/services services/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/canoe canoe/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/jetski jetski/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/kayak kayak/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/scuba scuba/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/waterski waterski/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/windsurf windsurf/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

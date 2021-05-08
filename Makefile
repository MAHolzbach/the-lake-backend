.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/boats boats/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/breeze breeze/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/lancer lancer/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/nina nina/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/pinta pinta/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/santaMaria santaMaria/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/waverunner waverunner/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

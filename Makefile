.PHONY: all build clean upload

all:
	@echo "What do you want to do?"

build:
	GOOS=linux GOARCH=mipsle GOMIPS=hardfloat CGO_ENABLED=0 go build -o linkit-lab -v

clean:
	go clean

upload:
	scp ./linkit-lab root@mylinkit.local:./
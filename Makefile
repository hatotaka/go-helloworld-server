DOCKER_USER?=$(shell whoami)

ifdef DOCKER_REGISTRY
IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_USER)/helloworld-server
else
IMAGE_NAME=$(DOCKER_USER)/helloworld-server
endif

all: build-linux build-docker

build:
	go build -o go-helloworld-server

build-linux:
	GOOS=linux GOARCH=amd64 go build -o go-helloworld-server

build-docker:
	sudo docker build -t $(IMAGE_NAME) .

run:
	docker run -d -p 8080:8080 $(IMAGE_NAME)

clean:
	rm -f go-helloworld-server

push:
	sudo docker push $(IMAGE_NAME)


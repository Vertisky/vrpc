SHELL := /bin/zsh

.PHONY: setup

# setup will create the rabbitmq docker container
setup:
	docker-compose up -d

tidy:
	go mod tidy && cd ./service && go mod tidy 

.PHONY: version ps desc build up down build_up in deploy start test

APP_VERSION=${shell git rev-parse --short HEAD}

version:
	echo 'Version: ${APP_VERSION}'

#dockerのコマンドの設定
ps:
	docker-compose ps

desc:
	@ echo ""; echo "====コンテナ====" && docker ps -a && echo ""; echo "====イメージ====" && docker images &&  echo ""; echo "====ボリューム====" && docker volume ls

build:
	docker-compose build

up:
	docker-compose up

up_db:
	docker-compose up db

down:
	docker-compose down

build_up:
	docker-compose up --build

in:
	#docker container run -it bc_app bash
	docker-compose exec app bash

in_db:
	docker-compose exec db bash

reset:
	docker-compose down && docker system prune -a --volumes

deploy:
	docker-compose push app


# Goのコマンドの設定
start:
	go run ./src/cmd/main.go

go_build:
	go build ./src/cmd/main.go

test:
	go test -v -cover  ./src/usecase/... | grep -v -e "input" -e "output"

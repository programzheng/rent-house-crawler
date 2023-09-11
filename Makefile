#讀取.env
include ./.env
export $(shell sed 's/=.*//' ./.env)

#當前年-月-日
DATE=$(shell date +"%F")
COMPOSE=docker compose
BASH?=bash
SERVICES=rent-house-crawler

.PHONY: dev, up, init, down
bash:
	$(COMPOSE) exec $(SERVICES) $(BASH)

#重新編譯
dev:
	$(COMPOSE) build $(SERVICES)
	$(COMPOSE) up $(SERVICES)

#啟動服務
up:
	$(COMPOSE) up -d $(SERVICES)

#重啟服務
restart:
	$(COMPOSE) restart

#初始化
init:
	$(COMPOSE) build --force-rm --no-cache
	$(MAKE) up
#列出容器列表
ps:
	$(COMPOSE) ps

#關閉所有服務
down:
	$(COMPOSE) down

build-image:
	docker build -t programzheng/rent-house-crawler -f Dockerfile.linux --platform linux/amd64 .

push-image:
	docker push programzheng/rent-house-crawler
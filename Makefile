.PHONY: up down build
up:
	docker-compose up -d
down:
	docker-compose down
build:
	docker-compose build

.PHONY: go go-logs
go:
	docker-compose exec go bash
go-logs:
	docker-compose logs go -f

.PHONY: node node-logs npm storybook
node:
	docker-compose exec node bash
npm:
	docker-compose exec node npm i
node-logs:
	docker-compose logs node -f.PHONY: mysql
storybook:
	docker-compose exec node npm run storybook

.PHONY: mysql mysql-logs
mysql:
	docker-compose exec mysql bash
mysql-logs:
	docker-compose logs mysql -f

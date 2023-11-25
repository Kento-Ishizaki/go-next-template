.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: build
build:
	docker-compose build

.PHONY: go
go:
	docker-compose exec go bash
.PHONY: go-logs
go-logs:
	docker-compose logs go -f

.PHONY: node
node:
	docker-compose exec node bash
.PHONY: node-logs
node-logs:
	docker-compose logs node -f

.PHONY: mysql
mysql:
	docker-compose exec mysql bash
.PHONY: mysql-logs
mysql-logs:
	docker-compose logs mysql -f

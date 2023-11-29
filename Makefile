step := 0
name :=

.PHONY: up down build logs
up:
	docker compose up -d
down:
	docker compose down
build:
	docker compose build
logs:
	docker compose logs -f

.PHONY: go go-logs migrate
go:
	docker compose exec go bash
go-logs:
	docker compose logs go -f
gen-migration:
	docker compose exec go migrate create -ext sql -dir db/migrations -seq ${name}
migrate-up:
	docker compose exec go migrate -path db/migrations -database "mysql://db_user:password@tcp(mysql:3306)/go_next?multiStatements=true" up ${step}
migrate-down:
	docker compose exec go migrate -path db/migrations -database "mysql://db_user:password@tcp(mysql:3306)/go_next?multiStatements=true" down ${step}
migrate-force:
	docker compose exec go migrate -path db/migrations -database "mysql://db_user:password@tcp(mysql:3306)/go_next" force ${step}
gen-init:
	docker compose exec go gqlgen init

.PHONY: node node-logs npm storybook
node:
	docker compose exec node bash
npm:
	docker compose exec node npm i
node-logs:
	docker compose logs node -f.PHONY: mysql
storybook:
	docker compose exec node npm run storybook

.PHONY: mysql mysql-logs
mysql:
	docker compose exec mysql bash
mysql-logs:
	docker compose logs mysql -f

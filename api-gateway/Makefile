dev:
	go mod tidy && wgo run .

debug:
	go mod tidy && dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --continue --accept-multiclient .

watch:
	go mod tidy && find . -name "*.go" | entr -r dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --continue .

generate:
	sqlc generate

# Migrations and seed
migrate-up:
	sql-migrate up -env="database"

migrate-down:
	sql-migrate down -env="database"

migrate-status:
	sql-migrate status -env="database"

migrate-new:
	sql-migrate new -env="database" $(name)

seed-up:
	sql-migrate up -env="seed"

seed-down:
	sql-migrate down -env="seed"

seed-status:
	sql-migrate status -env="seed"

seed-new:
	sql-migrate new -env="seed" $(name) 
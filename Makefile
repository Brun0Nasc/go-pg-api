include app.env
export

dockerUp:
	sudo docker compose up -d

dockerDown:
	sudo docker compose down

createMigrations:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrateUp:
	migrate -path db/migrations -database "${POSTGRES_SOURCE}" -verbose up

migrateDown:
	migrate -path db/migrations -database "${POSTGRES_SOURCE}" -verbose down

.PHONY: 
.SILENT: 

create-table :
	migrate create -ext sql -dir db/migrations -seq ${name}

migrate:
	migrate -database postgres://postgres:postgres@localhost:5432/sports_day?sslmode=disable -path db/migrations up
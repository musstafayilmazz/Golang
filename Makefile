createdb:
	docker exec -it simple_bank_postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simple_bank_postgres dropdb simple_bank

migrate:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

rollback:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down


.PHONY: createdb dropdb migrate rollback

	
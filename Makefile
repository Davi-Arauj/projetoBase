postgres:
	sudo docker run --name academia_back -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
createdb:
	sudo docker exec -it academia_back createdb --username=root --owner=root academia
dropdb:
	sudo docker exec -it academia_back dropdb academia
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/academia?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/academia?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination /home/davi/go/src/github.com/academia_back/db/mock/store.go pdv/db/sqlc Store
	
.PHONY:postgres createdb dropdb migrateup migratedown test server mock
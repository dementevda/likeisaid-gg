build:
	go build -v ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v -race -timeout 60s -cover ./...

# docker run --name likeisaid -p 27017:27017 -d -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root mongo mongod --auth
# db.createUsers({user: "likeisaid", pwd: "saidilike", roles: [{role: "readWrite", db: "likeisaid"}]})
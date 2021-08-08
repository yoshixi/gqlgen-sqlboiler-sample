# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install


BINARY_NAME=gqlgen-sqlboiler-sample
BINARY_UNIX=$(BINARY_NAME)_unix
BUILD_PATH=./server.go

RUN_SEVER=$(GORUN) $(BUILD_PATH)

all: deps db run_server
run: refresh run_server
deps: install_gqlgen install_dbmate
db:
	dbcreate
	dbmigrate
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(BUILD_PATH)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
start:
	$(GOBUILD) -o $(BINARY_NAME) -v $(BUILD_PATH)
	./$(BINARY_NAME)
run_server:
	$(RUN_SEVER)
refresh:
	$(GORUN) github.com/99designs/gqlgen generate --verbose
dbmate:
	bin/dbmate
dbcreate:
	dbmate create && dbmate -e TEST_DATABASE_URL --no-dump-schema create
dbmigrate:
	dbmate migrate && dbmate -e TEST_DATABASE_URL --no-dump-schema migrate
dbrollback:
	dbmate rollback && dbmate -e TEST_DATABASE_URL --no-dump-schema rollback
dbdrop:
	dbmate drop && dbmate -e TEST_DATABASE_URL --no-dump-schema drop
install_dbmate:
	mkdir -p bin && curl -fsSL -o bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64 && chmod +x bin/dbmate
install_gqlgen:
	$(GOGET) github.com/99designs/gqlgen

# https://github.com/volatiletech/sqlboiler/issues/743
sqlboiler:
	go run github.com/volatiletech/sqlboiler/v4 psql --struct-tag-casing camel

# クロスコンパイル
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

TAG=`git describe --always --tags | cut -c 2-`
PROJECT=stori_challenge

run:
ifdef PORT
	@ go run cmd/api/*.go --port $(PORT)
else
	@ go run cmd/api/*.go
endif

build:
	@ go build -o ./.bin/api/server cmd/api/*.go

mocking:
	@ rm -r mocks
	@ mockery

test:
ifdef FILE
	@ dotenv -e .env.test -- go run github.com/onsi/ginkgo/v2/ginkgo $(FILE)
else
	@ dotenv -e .env.test -- go run github.com/onsi/ginkgo/v2/ginkgo -v --label-filter="!integration" ./...
endif

test_integration:
	@ dotenv -e .env.test -- go run github.com/onsi/ginkgo/v2/ginkgo -v --label-filter="integration" ./...

lint:
	@ golangci-lint run

build_image:
	@ docker build . -t "$(PROJECT):latest"
	@ docker build . -t "$(PROJECT):$(TAG)"

gen_swag:
	@ swag init --dir cmd/api --output cmd/api/docs --parseInternal --parseDependency --parseDepth 1

fmt:
	@ go fmt ./...
	@ swag fmt
	@ npx prisma format

version:
	@ cz version -p

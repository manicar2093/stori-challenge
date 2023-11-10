TAG=`git describe --always --tags | cut -c 2-`
PROJECT=stori_challenge

run:
ifdef PORT
	@ dotenv -e .env.dev -- go run cmd/api/*.go --port $(PORT)
else
	@ dotenv -e .env.dev -- go run cmd/api/*.go
endif

build:
	@ go build -o ./.bin/api/server cmd/api/*.go

mocking:
	@ rm -r mocks
	@ mockery

test:
	@ dotenv -e .env.test -- npx prisma db push
	@ dotenv -e .env.test -- npx prisma db seed
ifdef FILE
	@ dotenv -e .env.test -- go run github.com/onsi/ginkgo/v2/ginkgo $(FILE)
else
	@ dotenv -e .env.test -- go run github.com/onsi/ginkgo/v2/ginkgo -v --label-filter="!integration" ./...
endif
	@ dotenv -e .env.test -- npx prisma migrate reset --force

test_integration:
	@ dotenv -e .env.test -- go run github.com/onsi/ginkgo/v2/ginkgo -v --label-filter="integration" ./...

lint:
	@ golangci-lint run

build_image:
	@ docker build . -t "$(PROJECT):latest"
	@ docker build . -t "$(PROJECT):$(TAG)"

migrate_postgres:
ifdef ENV
	@ dotenv -e .env.$(ENV) -- npx prisma migrate dev --skip-generate --skip-seed
else
	@ dotenv -e .env.dev -- npx prisma migrate dev --skip-generate --skip-seed
endif

migrate_create:
ifdef ENV
	@ dotenv -e .env.$(ENV) -- npx prisma migrate dev --skip-generate --skip-seed --create-only
else
	@ dotenv -e .env.dev -- npx prisma migrate dev --skip-generate --skip-seed --create-only
endif

validate_postgres_migration:
	@ dotenv -e .env.test -- npx prisma migrate dev --skip-generate --skip-seed --create-only

reset_dev_db:
	@ dotenv -e .env.dev -- npx prisma migrate reset

gen_swag:
	@ swag init --dir cmd/api --output cmd/api/docs --parseInternal --parseDependency --parseDepth 1

fmt:
	@ go fmt ./...
	@ swag fmt
	@ npx prisma format

version:
	@ cz version -p

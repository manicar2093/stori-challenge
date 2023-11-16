TAG=`git describe --always --tags | cut -c 2-`
PROJECT=stori_challenge

run:
ifdef PORT
	@ go run cmd/api/*.go --port $(PORT)
else
	@ go run cmd/api/*.go
endif

run_demo: build_image
	@ docker run --rm --env-file ./env.demo -p 8000:5001 -v demo_vol:/data --name stori_challenge-demo stori_challenge:latest

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

version:
	@ cz version -p

deploy_sls:
	@ ./scripts/aws/01-aws-lambda-compiler.sh
	@ dotenv -e .env.dev -- serverless deploy --aws-profile $(AWS_PROFILE_DEPLOY)

remove_sls:
	@ dotenv -e .env.dev -- serverless remove --aws-profile $(AWS_PROFILE_DEPLOY)

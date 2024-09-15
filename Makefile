test:
	@go test -v ./...

build:
	@go build -o bin/ecommerce cmd/main.go

run: build
	@./bin/ecommerce

dcbuild:
	@docker-compose build

dcup:
	@docker-compose up -d

dcrun:
	@docker-compose build
	@docker-compose up -d

dcdown:
	@docker-compose down

template:
	@npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css
	@templ generate

tailwind:
	@npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

templ:
	@templ generate -watch
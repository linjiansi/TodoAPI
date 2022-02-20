.PHONY: build
build:
	go build -o bin/app cmd/main.go

.PHONY: clean
clean:
	rm bin/app

.PHONY: compose/up
compose/up: compose/up/db
	docker-compose up -d app

.PHONY: compose/up/db
compose/up/db: compose/down
	docker-compose up -d db

.PHONY: compose/down
compose/down:
	docker-compose down

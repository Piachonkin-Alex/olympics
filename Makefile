runserver:
	./cmd/runserver/runserver --config "./configs/config.yaml"

start-db:
	docker-compose up -d db-migrate

stop-db:
	docker-compose down

build:
	bash  `cd cmd/runserver && go build -o runserver.exe`

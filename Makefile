start:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
stop:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml down

start-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build

stop-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml down

	
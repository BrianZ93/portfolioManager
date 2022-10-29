default: build

run-local:
	docker login
	docker build --target="frontend" --tag frontend .
	docker build --target="backend" --tag backend .
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up


stop-local:
	docker login
	docker-compose down --remove-orphans
	docker image prune -a
	docker network prune




default: build

run-local:
	docker login
	docker build --file=blackswan/Dockerfile  -t web-frontend .
	docker build --file=backend/Dockerfile  -t web-backend .
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up

stop-local:
	docker-compose -f docker-compose.yml down




default: build

run-local:
	docker login
	docker build --target="frontend" --tag frontend .
	docker build --target="backend" --tag backend .
	# docker build --target="webscraper" --tag webscraper .
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up


clean:
	docker login
	docker-compose down --remove-orphans 
	docker image prune -a
	docker network prune

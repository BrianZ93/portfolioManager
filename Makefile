default: build

run-local:
	docker login
	docker build --target="frontend" --tag frontend .
	docker build --target="backend" --tag backend .
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up


clean:
	docker login
	docker-compose down --remove-orphans 
	docker image prune -a
	docker network prune

# air:
# 	alias air='$(go env GOPATH)/bin/air'
# 	air

stop-appArmor:
	sudo aa-status
	sudo systemctl disable apparmor.service --now 
	sudo service apparmor teardown
	sudo aa-status

restart-appArmor:
	sudo systemctl start apparmor
	sudo apparmor_parser -r /var/lib/snapd/apparmor/profiles/*
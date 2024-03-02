build-dev:
	docker-compose --verbose -f docker-compose-dev.yml -p paybill build

run-dev:
	docker-compose --verbose  -f docker-compose-dev.yml -p paybill up --build --remove-orphans

prune-dev:
	docker-compose -f docker-compose-dev.yml -p paybill down --remove-orphans --volumes --rmi all
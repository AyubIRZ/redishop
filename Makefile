build:
	docker-compose -f docker-compose.yml up -d --build --force-recreate

build_app:
	docker-compose -f docker-compose.yml up -d --build --no-deps app

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down
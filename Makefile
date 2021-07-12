dbup:
	docker-compose -f docker-compose.yml --env-file ./.env.dev up -d
dbdown:
	docker-compose -f docker-compose.yml down

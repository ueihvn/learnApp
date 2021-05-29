dbup:
	docker-compose -f docker-compose.yml up -d
dbdown:
	docker stop learnapp_db_1

init:
	docker-compose up -d postgres_db
	docker-compose up -d

app_restart:
	docker-compose up -d --build app
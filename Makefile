include .env

dockerRun :
	docker run --name mysqlIntern -p ${DATABASE_PORT}:${DATABASE_PORT} -e MYSQL_ROOT_PASSWORD=${DATABASE_PASSWORD} -d mysql:8.2

createdb:
	docker exec -it mysqlIntern mysql -e "CREATE DATABASE ${DATABASE_NAME};" -u ${DATABASE_USERNAME} -p

dropdb:
	docker exec -it mysqlIntern mysql -e "DROP DATABASE ${DATABASE_NAME};" -u ${DATABASE_USERNAME} -p

phony : dockerRun createdb dropdb
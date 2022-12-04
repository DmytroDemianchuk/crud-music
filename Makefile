build:
	docker-compose build crud-musics

run:
	docker-compose up crud-musics

# migrate:
# 	migrate \
# 		-source file://docker/migrations \
# 		-database postgres://postgres:qwerty@localhost:5432/musics?sslmode=disable \
# 		up

migrate:
	migrate -path ./docker/migrations -database postgres://postgres:goLANGninja@localhost:5432/postgres?sslmode=disable up
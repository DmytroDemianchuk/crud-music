run:
	docker-compose up -d --build

stop:
	docker-compose down

db-logs:
	docker-compose logs -f db

movies-app-logs:
	docker-compose logs -f musics-app

migrates-up:
	migrate \
		-source file://docker/migrations \
		-database postgres://postgres:goLANGn1nja@localhost:5432/musics?sslmode=disable \
		up

migrates-down:
	migrate \
		-source file://docker/migrations \
		-database postgres://postgres:goLANGn1nja@localhost:5432/musics?sslmode=disable \
		down
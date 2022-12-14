run:
	docker-compose up -d --build

stop:
	docker-compose down

db-logs:
	docker-compose logs -f db

musics-app-logs:
	docker-compose logs -f musics-app

migrates-up:
	migrate \
		-source file://docker/migratios \
		-database postgres://postgres:goLANGninja@localhost:5432/musics?sslmode=disable \
		up

migrates-down:
	migrate \
		-source file://docker/migratios \
		-database postgres://postgres:goLANGninja@localhost:5432/musics?sslmode=disable \
		down
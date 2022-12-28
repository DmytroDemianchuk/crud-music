# REST API Crud-music

Use command: `git clone https://github.com/dmytrodemianchuk/crud-music`

### Build and Start


```
source .env && go build -o app cmd/main.go && ./app
```

For postgres use Docker

```
docker run -d --name ninja-db -e POSTGRES_PASSWORD=qwerty123 -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres
```


## CRUD operations:
POST - "/music" - create music

GET - "/music/:id" - get music by id

GET "/musics" - get all musics

PUT "/music/:id" - update music by id

DELETE "/music/:id" - delete music by id

## Example of creating an music:
In Postman you choose "Body" menu, POST `localhost:8080/music` and type for example:

```
{
"name": "Spiral",  
"performer": "21 Savage",  
"realise_year": 2021,  
"genre": "Hip-Hop"
}
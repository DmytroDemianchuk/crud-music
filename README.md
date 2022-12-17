# REST API FOR crud-music

Use command: `git clone https://github.com/dmytrodemianchuk/crud-music`

## Commands of application
Use this command in the directory
- `make run` - to run an application
- `make stop` - to stop an application

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
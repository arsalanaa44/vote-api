## API for poll-CRUD golang project

### <a id="getting-started">Getting Started</a>

Clone project, install Docker and Docker Compose. In order to get the project up and running follow steps below :

set your configs in `app.env` :

```bash
cp app.env.sample app.env
```
docker containers :
```bash
docker compose up -d 
```
database migration:
```bash
go run migrate/migrate.go 

```
go run 
get golang packages :
```bash
go mod tidy
```
install [air](https://github.com/cosmtrek/air) to autorun the server by this command :
```bash
air
```
Now the API should be up and running :)  
you can use my simple [postman collection](https://github.com/arsalanaa44/vote-api/blob/master/voting-system.postman_collection.json)   
good luck !

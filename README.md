## Restfull API written in GO

Me trying to learn golang by coding a rest api with clean archetecture
following the SOLID principles. So far I've refactored and extracted all
dependencies and implemented interfaces so that its easy to swap out db
(postgres) and routing (gorilla/mux). Then added SQLite as a DB which I (will)
use for testing.

### Status:
* [x] "/" => "Hello world!"
* [x] Serve static page
* [ ] Serve some static pages (home, errors)
* [x] Database (postgres in docker)
* [x] GET request (get all)
* [x] GET request (get by id)
* [x] POST request
* [x] Cleaner archetecture
* [x] Tests (service layer)
* [x] Dockerfile
* [x] docker-compose
* [x] Improove README.md
* [x] Add comments to code
* [ ] Tests (controller) - in progress
* [ ] Alternative DB (SQLITE) - in progress
* [ ] PUT request
* [ ] PATCH request
* [x] DELETE request
* [ ] MOAR Tests!
* [ ] Implement swagger
* [ ] Add Authorization


### How to use:
Run db and server in docker with docker-compose: 
```shell
docker-compose up
```

Testing with curl:

Get all books: 
```shell
curl localhost:8080/api/books
```
Get book with id = 1: 
```shell
curl localhost:8080/api/books/1
```
Create a book:
```shell
curl localhost:8080/api/books -H "application/json" -d '{"title":"<title>","author":"<author>", "year": 2022}' -X POST
```
Delete a book with id = 1:
```shell
curl localhost:8080/api/books/delete/1 -X DELETE
```

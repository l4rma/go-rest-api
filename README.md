## Restfull API written in GO

Me trying to learn golang by coding a rest api using gorilla/mux and postgres.

### Functionality:
Create, Read, ~~Update~~ and Delete books.

### How to use:
Run db and server in docker with docker-compose: ``$ docker-compose up``

#### Testing with curl:
* GET all books : ``curl localhost:8080/api/books``
* GET book with id = 1: ``curl localhost:8080/api/books/1``
* Create a book: ``curl localhost:8080/api/books -H "application/json" -d '{"title":"<title>","author":"<author>", "year": 2022}' -X POST``
* Delete a book with id = 1: ``curl localhost:8080/api/books/delete/1 -X DELETE``

### Todo:

* [x] "/" => "Hello world!"
* [ ] Serve some static pages (home, error)
* [x] Database
* [x] GET request (get all)
* [x] GET request (get by id)
* [x] POST request
* [x] Cleaner archetecture
* [x] Tests (service layer)
* [x] Dockerfile
* [x] docker-compose
* [ ] PUT request
* [ ] PATCH request
* [x] DELETE request
* [ ] MOAR Tests!
* [ ] Implement swagger
* [ ] Add Authorization


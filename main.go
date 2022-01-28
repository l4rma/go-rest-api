package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/l4rma/go-rest-api/server/controller"
	"github.com/l4rma/go-rest-api/server/db/repository"
	router "github.com/l4rma/go-rest-api/server/http"
	"github.com/l4rma/go-rest-api/server/service"
	_ "github.com/lib/pq"
)

var (
	bookRepository repository.BookRepository = repository.NewPostgresRepository()
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	r := bookRepository
	err := r.Open()
	if err != nil {
		panic(err)
	}
	defer r.Close()

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	httpRouter.GET("/books", bookController.GetBooks)
	httpRouter.POST("/books", bookController.AddBook)
	httpRouter.DELETE("/api/books/delete/{id}", bookController.DeleteBookById)

	books, err := r.FindAll()
	if len(books) < 1 {
		r.InsertDummyData(r)
	}

	httpRouter.SERVE(os.Getenv("PORT"))
}

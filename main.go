package main

import (
	"os"

	"github.com/l4rma/go-rest-api/server/controller"
	"github.com/l4rma/go-rest-api/server/db/repository"
	router "github.com/l4rma/go-rest-api/server/http"
	"github.com/l4rma/go-rest-api/server/service"
)

var (
	bookRepository repository.BookRepository = repository.NewPostgresRepository()
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {

	//httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Hello World!")
	//})

	httpRouter.GET("/api/books", bookController.GetBooks)
	httpRouter.GET("/api/books/{id}", bookController.GetBookById)
	httpRouter.POST("/api/books", bookController.AddBook)
	httpRouter.DELETE("/api/books/delete/{id}", bookController.DeleteBookById)

	httpRouter.SERVE(os.Getenv("PORT"))
}

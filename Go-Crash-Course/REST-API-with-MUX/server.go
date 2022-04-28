package main

import (
	"fmt"
	"net/http"

	"github.com/Param-Scalent/Go_Practice/Go-Crash-Course/REST-API-with-MUX/controller"
	router "github.com/Param-Scalent/Go_Practice/Go-Crash-Course/REST-API-with-MUX/http"
	"github.com/Param-Scalent/Go_Practice/Go-Crash-Course/REST-API-with-MUX/repository"
	"github.com/Param-Scalent/Go_Practice/Go-Crash-Course/REST-API-with-MUX/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and Running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}

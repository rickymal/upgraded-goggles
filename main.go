package main
import handlers "athena/handlers"
import database "athena/models"
import gin "github.com/gin-gonic/gin"

func main() {
	database.InitializeDatabase()

	router := gin.Default()
	router.POST("/movies", handlers.CreateMovie)
	router.GET("/movies", handlers.GetAllMovies)
	router.GET("/movies/:id", handlers.GetMovieByID)
	router.DELETE("/movies/:id", handlers.DeleteMovieByID)
	router.PUT("/movies/:id", handlers.UpdateMovie)
	router.Run(":8080")
}
package handlers
import database "athena/models"
import "net/http"
import "github.com/gin-gonic/gin"
import "strconv"


func GetAllMovies(context *gin.Context) {
	context.JSON(http.StatusOK, database.DataMovies) 
}

func CreateMovie(context *gin.Context) {
    var movie database.Movie
	err := context.ShouldBindJSON(&movie)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H {"error": err.Error()} )
        return
    }

	idx := database.AddMovie(movie)
    context.JSON(http.StatusCreated, gin.H {"id" : idx, "message": "Movie created successfully"} )
}



func GetMovieByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error" : "Invalid 'id'. This movie realy exists?",
		})
	}

	for _, movie := range database.DataMovies {
		if movie.Id == id {
			context.JSON(http.StatusOK, movie)
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H {
		"error" : "Movie not founded",
	})
}


func UpdateMovie(context *gin.Context) {
	var movieUpdate database.Movie

	errJsonParsing := context.ShouldBindJSON(&movieUpdate)
	if errJsonParsing != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error" :  errJsonParsing.Error,
		})
		return
	}

	idStr := context.Param("id")
	id, errConversion := strconv.Atoi(idStr)
	if errConversion != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error" : "Invalid 'id'. This movie realy exists?",
		})
		return
	}

	for idx, actualMovie := range database.DataMovies {
		if actualMovie.Id == id {
			database.UpdateMovie(idx, movieUpdate)
			
			context.JSON(http.StatusOK, gin.H {
				"message" : "Movie updated successfully",
			})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H {
		"error" : "Movie not founded",
	})
}

func DeleteMovieByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error" : "Invalid 'id'. This movie realy exists?",
		})
		return
	}

	for idx, movie := range database.DataMovies {
		if movie.Id == id {
			database.DeleteMovie(idx)
			context.JSON(http.StatusOK, gin.H {
				"message" : "Movie deleted successfully",
			})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H {
		"error" : "Movie not founded",
	})
}
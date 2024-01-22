package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


type Todo struct {
	Name string
	Description string
	ID string
	IsDone bool
}
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album
	if err :=c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums,newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}



var todos = []Todo {
	{ID: "1", Name: "In-house", Description: "Bring trash Outside"},
	{ID: "2", Name: "Outside", Description: "Clean Patio"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var addTodo Todo
	
	todos = append(todos, addTodo)
	c.IndentedJSON(http.StatusOK, addTodo)
}

func main() {
	router := gin.Default()
	router.GET("/", getTodos)
	router.POST("/", addTodo)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums",postAlbums)

	router.Run("localhost:8080")
}
package main
import (
  "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/ping", pingHandler)
	router.GET("/query", queryHandler)
	
	router.GET("/books/:id", booksHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	
	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})
}

func postBooksHandler(c *gin.Context) {
	var bookRequest struct {
		Title string
		Price int
		SubTitle string `json:"sub_title"`
	}

	err := c.ShouldBindJSON(&bookRequest)
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"title": bookRequest.Title,
		"price": bookRequest.Price,
		"sub_title": bookRequest.SubTitle,
	})
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Beer struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

var albums = []Beer{
	{Name: "Augustiner Hell", Price: "56.99€"},
	{Name: "Giesinger", Price: "17.99€"},
}

// handler, normally database request
func getBeers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	apiGroup := router.Group("/api")
	apiGroup.GET("/beers", getBeers)
	router.Run("0.0.0.0:8080")
}

// package main

// import (
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     // Create a new Gin router
//     r := gin.Default()

//     // Create a route group with the /api prefix
//     apiGroup := r.Group("/api")
//     {
//         apiGroup.GET("/users", func(c *gin.Context) {
//             c.JSON(200, gin.H{
//                 "message": "List of users",
//             })
//         })

//         apiGroup.POST("/users", func(c *gin.Context) {
//             c.JSON(201, gin.H{
//                 "message": "User created",
//             })
//         })

//         apiGroup.GET("/posts", func(c *gin.Context) {
//             c.JSON(200, gin.H{
//                 "message": "List of posts",
//             })
//         })
//     }

//     // Start the server
//     r.Run() // By default, it will listen and serve on 0.0.0.0:8080
// }

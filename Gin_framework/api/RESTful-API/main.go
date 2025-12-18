package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-go-learning/docs" // Swagger docs (อยู่ที่ api/docs)
)

// openBrowser opens the specified URL in the default browser
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}
	if err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}
}

// @title Album API
// @version 1.0
// @description RESTful API สำหรับจัดการข้อมูล Albums
// @host localhost:8082
// @BasePath /

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id" example:"1"`
	Title  string  `json:"title" example:"Blue Train"`
	Artist string  `json:"artist" example:"John Coltrane"`
	Price  float64 `json:"price" example:"56.99"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
// @Summary Get all albums
// @Description ดึงรายการ albums ทั้งหมด
// @Tags albums
// @Produce json
// @Success 200 {array} Album
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album from JSON received in the request body.
// @Summary Create a new album
// @Description เพิ่ม album ใหม่
// @Tags albums
// @Accept json
// @Produce json
// @Param album body Album true "Album data"
// @Success 201 {object} Album
// @Failure 400 {object} map[string]string
// @Router /albums [post]
func postAlbum(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// @Summary Get album by ID
// @Description ดึง album ตาม ID
// @Tags albums
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} Album
// @Failure 404 {object} map[string]string
// @Router /albums/{id} [get]
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

func main() {
	fmt.Println("Starting server on port 8082...")
	router := gin.Default()
	
	// Root handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin Server"})
	})

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handle 404 (Not Found)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
			"path":    c.Request.URL.Path,
		})
	})

	// Auto open Swagger UI in browser
	go func() {
		time.Sleep(500 * time.Millisecond) // รอให้ server พร้อม
		fmt.Println("\n🚀 Opening Swagger UI at http://localhost:8082/swagger/index.html")
		openBrowser("http://localhost:8082/swagger/index.html")
	}()

	router.Run(":8082")
}

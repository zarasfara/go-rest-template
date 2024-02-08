package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/go-rest-template/internal/delivery/http/v1/book"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	
	// Here we can use global middleware such cors...
	// router.Use()

	// init auth handlers
	_ = router.Group("/auth")

	api := router.Group("/api")
	api.GET("/health", HealthHandler)

	v1 := api.Group("/v1")

	book.RegisterHandlers(v1.Group("/books"), book.NewService(book.NewPostgresRepository(db)))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
	})

	return router
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
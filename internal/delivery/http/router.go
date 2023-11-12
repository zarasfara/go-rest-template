package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/go-rest-template/internal/delivery/http/v1/book"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	// init auth handlers
	_ = router.Group("/auth")

	api := router.Group("/api")
	v1 := api.Group("/v1")

	book.RegisterHandlers(v1.Group("/books"), book.NewService(book.NewRepository(db)))

	return router
}

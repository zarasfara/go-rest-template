package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarasfara/go-rest-template/internal/entity"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.GET("/", res.get)

}

type resource struct {
	service Service
}

func (r resource) get(c *gin.Context) {
	var books []entity.Book
	books = append(books, entity.Book{Id: 1})

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
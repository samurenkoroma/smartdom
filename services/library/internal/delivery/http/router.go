package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	jsonBook "smartdom/services/library/internal/delivery/http/book"
)

func (d *Delivery) routerBooks(router *gin.RouterGroup) {
	router.GET("/", d.ListContact)
	router.GET("/:id", d.GetBook)
}

func (d *Delivery) ListContact(c *gin.Context) {
	books, err := d.ucBook.List()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, books)
}

func (d *Delivery) GetBook(c *gin.Context) {
	var id jsonBook.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	book, err := d.ucBook.GetBookByID(id.Value)
	if err != nil {
		return
	}
	c.FileAttachment(book.Path.String(), book.Title.String())
}

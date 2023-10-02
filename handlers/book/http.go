package book

import (
	"Book_Management_System/handlers"
	"Book_Management_System/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	bookService handlers.BookSvc
}

func New(bookService handlers.BookSvc) handler {
	return handler{bookService: bookService}
}

func (h handler) Create(c *gin.Context) {
	var book models.Book

	ctx := c.Request.Context()
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	createBook, _, err := h.bookService.Create(ctx, &book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, createBook)
}

func (h handler) Update(c *gin.Context) {
	var book models.Book

	ctx := c.Request.Context()
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id := c.Param("id")

	book.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updateBook, err := h.bookService.Update(ctx, &book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, updateBook)
}

func (h handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	books, err := h.bookService.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, books)
}

func (h handler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book, err := h.bookService.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, book)
}

func (h handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.bookService.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusNoContent)
}

package bookIssue

import (
	"Book_Management_System/handlers"
	"Book_Management_System/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	issueSvc handlers.BookIssue
}

func New(issueSvc handlers.BookIssue) handler {
	return handler{
		issueSvc: issueSvc,
	}
}

func (h handler) Create(c *gin.Context) {
	var book models.Issue

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	createBook, err := h.issueSvc.Create(ctx, &book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, createBook)
}

func (h handler) Update(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()

	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.issueSvc.Update(ctx, bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func (h handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	books, err := h.issueSvc.GetAll(ctx)
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

	book, err := h.issueSvc.GetByID(ctx, id)
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

	err = h.issueSvc.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusNoContent)
}

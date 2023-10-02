package users

import (
	"Book_Management_System/handlers"
	"Book_Management_System/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	userService handlers.UserSvc
}

func New(userService handlers.UserSvc) handler {
	return handler{
		userService: userService,
	}
}

func (h handler) Create(c *gin.Context) {
	var user models.User

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	createUser, err := h.userService.Create(ctx, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, createUser)
}

func (h handler) Update(c *gin.Context) {
	var user models.User

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id := c.Param("id")

	user.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updateUser, err := h.userService.Update(ctx, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, updateUser)
}

func (h handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.userService.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}

func (h handler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.userService.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func (h handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.userService.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusNoContent)
}

package admin

import (
	"Book_Management_System/handlers"
	"Book_Management_System/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	adminService handlers.AdminSvc
}

func New(adminService handlers.AdminSvc) handler {
	return handler{adminService: adminService}
}

func (h handler) Login(c *gin.Context) {
	var (
		admin models.UserInput
	)

	ctx := c.Request.Context()
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	adminDetails, err := h.adminService.Login(ctx, &admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, adminDetails)
}

func (h handler) Logout(c *gin.Context) {
	email := fmt.Sprintf("%v", c.Value("abc"))

	err := h.adminService.Logout(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusNoContent)
}

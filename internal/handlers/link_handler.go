// Link handler
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vishwanththalla/linkme/internal/services"
)

type CreateLinkInput struct {
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required,url"`
}

func CreateLink(c *gin.Context) {
	var input CreateLinkInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	err := services.CreateLink(input.Title, input.URL, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link created successfully"})
}

func GetLinks(c *gin.Context) {
	userID := c.GetUint("user_id")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	links, err := services.GetUserLinks(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch links"})
		return
	}

	c.JSON(http.StatusOK, links)
}

type UpdateLinkInput struct {
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required,url"`
}

func UpdateLink(c *gin.Context) {
	linkIDParam := c.Param("id")
	userID := c.GetUint("user_id")

	var input UpdateLinkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	linkID, err := strconv.Atoi(linkIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid link id"})
		return
	}

	err = services.UpdateLink(uint(linkID), input.Title, input.URL, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link updated"})
}

func DeleteLink(c *gin.Context) {
	linkIDParam := c.Param("id")
	userID := c.GetUint("user_id")

	linkID, err := strconv.Atoi(linkIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid link id"})
		return
	}

	err = services.DeleteLink(uint(linkID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted"})
}

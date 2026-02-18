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
        respondError(c, http.StatusBadRequest, err.Error())
        return
    }

    userID := getUserID(c)

    if err := services.CreateLink(input.Title, input.URL, userID); err != nil {
        respondError(c, http.StatusInternalServerError, err.Error())
}

func GetLinks(c *gin.Context) {
    userID := getUserID(c)

    page, limit := parsePagination(c)

    links, err := services.GetUserLinks(userID, page, limit)
    if err != nil {
        respondError(c, http.StatusInternalServerError, "failed to fetch links")
	}

	c.JSON(http.StatusOK, links)
}

type UpdateLinkInput struct {
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required,url"`
}

func UpdateLink(c *gin.Context) {
	linkIDParam := c.Param("id")
    userID := getUserID(c)

    var input UpdateLinkInput
    if err := c.ShouldBindJSON(&input); err != nil {
        respondError(c, http.StatusBadRequest, err.Error())
        return
    }

    linkID, err := strconv.Atoi(linkIDParam)
    if err != nil {
        respondError(c, http.StatusBadRequest, "invalid link id")
        return
    }

    if err := services.UpdateLink(uint(linkID), input.Title, input.URL, userID); err != nil {
        respondError(c, http.StatusNotFound, "link not found")

	c.JSON(http.StatusOK, gin.H{"message": "Link updated"})
}

func DeleteLink(c *gin.Context) {
	linkIDParam := c.Param("id")
    userID := getUserID(c)

    linkID, err := strconv.Atoi(linkIDParam)
    if err != nil {
        respondError(c, http.StatusBadRequest, "invalid link id")
        return
    }

    if err := services.DeleteLink(uint(linkID), userID); err != nil {
        respondError(c, http.StatusNotFound, "link not found")

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted"})
}

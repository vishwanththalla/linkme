// Handler utility helpers
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// respondError sends a JSON error response with a consistent structure.
func respondError(c *gin.Context, code int, message interface{}) {
    c.JSON(code, gin.H{"error": message})
}

// getUserID extracts the authenticated user's ID from the context.
func getUserID(c *gin.Context) uint {
    return c.GetUint("user_id")
}

// parsePagination reads page/limit query parameters with sane defaults and bounds.
func parsePagination(c *gin.Context) (page, limit int) {
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")

    page, _ = strconv.Atoi(pageStr)
    limit, _ = strconv.Atoi(limitStr)

    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 100 {
        limit = 10
    }

    return
}

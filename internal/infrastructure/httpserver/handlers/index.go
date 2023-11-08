package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/wojciechpawlinow/find-indexes/internal/application/service"
	"github.com/wojciechpawlinow/find-indexes/pkg/logger"
)

type IndexHTTPHandler struct {
	Srv service.IndexPort
}

const (
	exactMatch   = "exact"
	closestMatch = "nearest"
	errNotFound  = "value not found"
)

// FindIndex is an HTTP controller method for finding an index by value
// value - url param, integer 0 - 1000000
func (h *IndexHTTPHandler) FindIndex(c *gin.Context) {
	vParam := c.Param("value")

	// basic validation
	v, err := strconv.Atoi(vParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value must be an integer"})
		return
	}

	if v < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value must be greater than zero"})
		return
	}

	logger.Debug(fmt.Sprintf("searching for value: %d", v))

	// pass the request further to the application layer
	idx, value, directMatch, err := h.Srv.FindByValue(c.Request.Context(), v)
	if err != nil {
		// most likely empty slice and no data to search through
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	if idx == -1 {
		logger.Info(fmt.Sprintf("value %d not found", v))
		c.JSON(http.StatusNotFound, gin.H{"error": errNotFound})
		return
	}

	resp := gin.H{
		"index": idx,
		"value": value,
		"match": exactMatch,
	}

	if !directMatch {
		resp["match"] = closestMatch
	}

	logger.Info(fmt.Sprintf("value %d found at index %d with exact match(%t)", v, idx, directMatch))

	c.JSON(http.StatusOK, resp)
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/wojciechpawlinow/find-indexes/internal/application/service"
)

type IndexHTTPHandler struct {
	Srv service.IndexPort
}

const (
	directMatch   = "exact"
	indirectMatch = "near"
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

	if v <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value must be greater than zero"})
		return
	}

	// pass the request further to the application layer
	result, nonDirect, err := h.Srv.FindByValue(c.Request.Context(), v)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	resp := gin.H{
		"index": result,
		"value": v,
		"match": directMatch,
	}

	if nonDirect {
		resp["match"] = indirectMatch
	}

	c.JSON(http.StatusOK, resp)
}

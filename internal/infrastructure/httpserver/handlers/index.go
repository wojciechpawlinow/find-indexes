package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"

	"github.com/wojciechpawlinow/find-indexes/internal/application/service"
)

type IndexFinderHTTPHandler struct {
	ctn di.Container
}

const (
	directMatch   = "exact"
	indirectMatch = "proximity"
)

func (h *IndexFinderHTTPHandler) FindIndex(c *gin.Context) {
	vParam := c.Param("value")

	v, err := strconv.Atoi(vParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value must be an integer"})
		return
	}

	if v <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value must be greater than zero"})
		return
	}

	s := h.ctn.Get("service-index").(service.IndexFinderPort)

	result, nonDirect, err := s.Find(c.Request.Context(), v)
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

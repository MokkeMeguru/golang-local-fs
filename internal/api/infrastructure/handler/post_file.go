package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostFileHandler struct {
	HostDir   string
	OverWrite bool
}

func (h *PostFileHandler) PostFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileName := c.Param("file-name")

	if strings.Index(fileName, "..") != -1 ||
		strings.Index(fileName, "\\") != -1 ||
		strings.Index(fileName, "/") != -1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1300})
		return
	}

	imageFilePath := filepath.Join(h.HostDir, fileName)

	_, err := os.Stat(imageFilePath)
	if err == nil && h.OverWrite {
		fmt.Printf("Error: Duplicate File is Uploaded")
		c.JSON(http.StatusBadRequest, gin.H{"code": 1500})
		return
	}

	if err := c.SaveUploadedFile(file, imageFilePath); err != nil {
		fmt.Printf("Error: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1501})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

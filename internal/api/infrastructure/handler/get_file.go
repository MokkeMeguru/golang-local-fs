package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetFileHandler struct {
	HostDir string
}

func (h *GetFileHandler) GetFile(c *gin.Context) {
	fileName := c.Param("file-name")

	if strings.Index(fileName, "..") != -1 ||
		strings.Index(fileName, "\\") != -1 ||
		strings.Index(fileName, "/") != -1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1300})
		return
	}

	imageFilePath := filepath.Join(h.HostDir, fileName)

	file, err := os.Stat(imageFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1400})
		return
	}

	if file.IsDir() {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1401})
		return
	}

	fileBytes, err := ioutil.ReadFile(imageFilePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1402,
			"message": err,
		})
		return
	}

	contentType := http.DetectContentType(fileBytes)

	c.Writer.Header().Add("Content-Disposition", "inline")
	c.Writer.Header().Add("Content-Type", contentType)
	c.File(imageFilePath)
}

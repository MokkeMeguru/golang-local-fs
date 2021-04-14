package router

import (
	"github.com/MokkeMeguru/golang-local-fs/internal/api/infrastructure/env"
	"github.com/MokkeMeguru/golang-local-fs/internal/api/infrastructure/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	getFileHandler  handler.GetFileHandler
	postFileHandler handler.PostFileHandler
}

func NewRouter(env *env.Env) (*Router, error) {

	getFileHandler := handler.GetFileHandler{
		HostDir: env.LocalFileRoot,
	}

	postFileHandler := handler.PostFileHandler{
		HostDir:   env.LocalFileRoot,
		OverWrite: env.OverWriteFile,
	}
	return &Router{
		getFileHandler:  getFileHandler,
		postFileHandler: postFileHandler,
	}, nil
}

func (r *Router) Execute() {
	router := gin.Default()
	router.GET("/files/:file-name", r.getFileHandler.GetFile)
	router.POST("/files/:file-name", r.postFileHandler.PostFile)
	router.Run(":5577")
}

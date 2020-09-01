package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faagerholm/fileserver/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	conf := config.GetConfig()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/files", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		_ = c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", conf.MEDIA_ROOT, file.Filename))

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! to %s/%s", file.Filename, conf.MEDIA_ROOT, file.Filename))
	})
	r.GET("/files", func(c *gin.Context) {
		filepath := c.Query("filename")
		c.File(fmt.Sprintf("%s/%s", conf.MEDIA_ROOT, filepath))
	})

	return r
}

package datafile

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/setting"
)

// DownloadGET impliments `/api/datafile/download/:filename`
func DownloadGET(c *gin.Context) {
	filename := filepath.Base(c.Param("filename"))
	if filename == "" {
		c.JSON(http.StatusBadRequest, "url is not in the format like `/api/datafile/donwload/:filename`")
		return
	}
	path, _ := setting.ResolveTextdatasPath(filename)

	s, err := os.Stat(path)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	f, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	defer f.Close()

	reader := f
	contentLength := s.Size()
	contentType := "text/csv"
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="` + filename + `"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

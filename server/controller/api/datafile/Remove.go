package datafile

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/setting"
)

// RemovePOST impliments `/api/datafile/remove/:filename`
func RemovePOST(c *gin.Context) {
	filename := filepath.Base(c.Param("filename"))
	if filename == "" {
		c.JSON(http.StatusBadRequest, "url is not in the format like `/api/datafile/remove/:filename`")
		return
	}
	path, _ := setting.ResolveTextdatasPath(filename)

	_, err := os.Stat(path)
	if err != nil {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	os.Remove(path)
	c.AbortWithStatus(http.StatusAccepted)
}

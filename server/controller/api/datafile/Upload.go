package datafile

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/domain/textdata"
	"github.com/pakuyuya/dbsql-meta-viewer/server/setting"
)

type UploadPOSTForm struct {
	Merge string `form:"merge"`
}

func UploadPOST(c *gin.Context) {
	p := UploadPOSTForm{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate file
	_, err = textdata.LoadAllCsv(f.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(f.Filename)
	path, _ := setting.ResolveTextdatasPath(filename)

	err = c.SaveUploadedFile(f, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go func() {
		path, _ := setting.ResolveTextdatasPath("*")
		datas, err := textdata.LoadAllCsv(path)
		if err != nil {
			return
		}
		textdata.SwitchRepository(&datas)
	}()

	c.JSON(http.StatusOK, gin.H{})
}

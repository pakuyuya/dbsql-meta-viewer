package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/domain/textdata"
	"github.com/pakuyuya/dbsql-meta-viewer/server/setting"
)

type ListGETResponse struct {
	List []datafileInfo `json:"list"`
}

type datafileInfo struct {
	Filename   string `json:"filename"`
	LastUpdate string `json:"lastupdate"`
}

func ListGET(c *gin.Context) {
	p, _ := setting.ResolveTextdatasPath()
	lfp := textdata.ListupFilePathes(p)

	list := make([]datafileInfo, 0)
	for _, fp := range lfp {
		info := datafileInfo{}
		info.Filename = filepath.Base(fp)

		fi, err := os.Stat(fp)
		if err == nil {
			info.LastUpdate = fi.ModTime().Format("2006-01-02 15:04:05")
		}

		list = append(list, info)
	}

	response := ListGETResponse{List: list}
	c.JSON(http.StatusOK, response)
}

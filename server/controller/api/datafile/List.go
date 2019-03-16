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
	Datas []datafileInfo `json:"datas"`
}

type datafileInfo struct {
	Filename   string `json:"filename"`
	LastUpdate string `json:"lastupdate"`
}

func ListGET(c *gin.Context) {
	p, _ := setting.ResolveTextdatasPath()
	lfp := textdata.ListupFilePathes(p)

	datas := make([]datafileInfo, 0)
	for _, fp := range lfp {
		info := datafileInfo{}
		info.Filename = filepath.Base(fp)

		fi, err := os.Stat(fp)
		if err == nil {
			info.LastUpdate = fi.ModTime().Format("2006-01-02 15:04:05")
		}

		datas = append(datas, info)
	}

	response := ListGETResponse{Datas: datas}
	c.JSON(http.StatusOK, response)
}

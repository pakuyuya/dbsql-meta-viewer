package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/domain/textdata"
)

type SearchtextGETParam struct {
	Query        string `form:"query" binding:"required"`
	WordOnly     string `form:"wordOnly"`
	IgnoreCase   string `form:"ignoreCase"`
	Regex        string `form:"regex"`
	BodyOnly     string `form:"bodyOnly"`
	ResponseBody string `form:"responseBody"`
}

type SearchtextGETResponse struct {
	Count int                   `json:"count"`
	Datas []TextdataForResponse `json:"datas"`
}
type TextdataForResponse struct {
	Idx     int    `json:"idx"`
	Caption string `json:"caption"`
	Body    string `json:"body"`
	Since   string `json:"since"`
}

func SearchtextGET(c *gin.Context) {
	var param SearchtextGETParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	datas, err := textdata.Search(param.Query,
		textdata.WithWordOnly(param.WordOnly == "true"),
		textdata.WithIgnoreCase(param.IgnoreCase != "false"),
		textdata.WithRegex(param.Regex == "true"),
		textdata.WithBodyOnly(param.BodyOnly == "true"),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	resdatas := make([]TextdataForResponse, 0)
	for i, data := range datas {
		resdata := TextdataForResponse{}
		resdata.Idx = i
		resdata.Caption = data.Caption
		if param.ResponseBody == "true" {
			resdata.Body = data.Body
		} else {
			resdata.Body = ""
		}
		resdata.Since = data.Since

		resdatas = append(resdatas, resdata)
	}

	response := SearchtextGETResponse{Count: len(resdatas), Datas: resdatas}
	c.JSON(http.StatusOK, response)
}

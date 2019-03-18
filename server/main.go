package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/controller/api"
	api_datafile "github.com/pakuyuya/dbsql-meta-viewer/server/controller/api/datafile"
	"github.com/pakuyuya/dbsql-meta-viewer/server/controller/page"
	"github.com/pakuyuya/dbsql-meta-viewer/server/domain/textdata"
	"github.com/pakuyuya/dbsql-meta-viewer/server/setting"
)

func main() {
	var err error
	err = setting.LoadSetting()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if setting.ServerSetting.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := loadDefaultTextdata(); err != nil {
		fmt.Println(err.Error())
		return
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	router.LoadHTMLGlob("templates/*")

	// static
	gpage := router.Group("/")
	{
		gpage.GET("/", page.IndexGET)
		gpage.Static("/static", "./static")
	}
	// api
	gapi := router.Group("/api")
	{
		gapi.GET("/server/status", api.ServerStatusGET)
		gapi.GET("/searchtext", api.SearchtextGET)
		gapi.GET("/datafile/list", api_datafile.ListGET)
		gapi.POST("/datafile/upload", api_datafile.UploadPOST)
		gapi.GET("/datafile/download/:filename", api_datafile.DownloadGET)
	}

	server := &http.Server{
		Addr:           ":" + setting.ServerSetting.Port,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 18,
	}
	if setting.ServerSetting.Debug {
		fmt.Println("run http server at " + server.Addr)
	}
	server.ListenAndServe()
}

func loadDefaultTextdata() error {
	// load from `/path/to/textdata/dir/*`
	path, _ := setting.ResolveTextdatasPath("*")
	datas, err := textdata.LoadAllCsv(path)
	if err != nil {
		return err
	}

	// switch repository on memory
	textdata.SwitchRepository(&datas)

	return nil
}

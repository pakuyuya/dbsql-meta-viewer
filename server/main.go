package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/dbsql-meta-viewer/server/controller/api"
	"github.com/pakuyuya/dbsql-meta-viewer/server/controller/page"
	"github.com/pakuyuya/dbsql-meta-viewer/server/domain/textdata"
	"github.com/pakuyuya/gopathmatch"
	yaml "gopkg.in/yaml.v2"
)

type ServerSettingType struct {
	Port            string `yaml:"Port"`
	Debug           bool   `yaml:"Debug"`
	DefaultTextdata string `yaml:"DefaultTextdata"`
}

var ServerSetting *ServerSettingType

func main() {
	var err error
	ServerSetting, err = loadSetting()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if ServerSetting.Debug {
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
	}
	gapi := router.Group("/api")
	{
		gapi.GET("/server/status", api.ServerStatusGET)
		gapi.GET("/searchtext", api.SearchtextGET)
	}

	server := &http.Server{
		Addr:           ":" + ServerSetting.Port,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 18,
	}
	if ServerSetting.Debug {
		fmt.Println("run http server at " + server.Addr)
	}
	server.ListenAndServe()
}

func loadSetting() (*ServerSettingType, error) {
	bytes, err := ioutil.ReadFile(`./setting.yml`)
	if err != nil {
		return nil, err
	}

	s := ServerSettingType{}
	err = yaml.Unmarshal(bytes, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func loadDefaultTextdata() error {
	files := gopathmatch.Listup(ServerSetting.DefaultTextdata, gopathmatch.FlgFileOnly)
	datas := make([]textdata.Textdata, 0)

	for _, file := range files {
		if ServerSetting.Debug {
			fmt.Printf("load default textdata from `%s`\r\n", file)
		}
		file, err := os.Open(file)
		if err != nil {
			return err
		}
		defer file.Close()

		dats, err := textdata.DecodeCSV(file)
		if err != nil {
			return err
		}
		datas = append(datas, dats...)
	}

	textdata.SwitchRepository(&datas)

	return nil
}

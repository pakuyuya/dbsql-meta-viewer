package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/simple-greptool/server/controller/api"
	"github.com/pakuyuya/simple-greptool/server/controller/page"
	yaml "gopkg.in/yaml.v2"
)

type ServerSettingType struct {
	Port  string `yaml:"Port"`
	Debug bool   `yaml:Debug`
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

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")

	// static
	gpage := router.Group("/")
	{
		gpage.GET("/", page.IndexGET)
	}
	gapi := router.Group("/api")
	{
		gapi.GET("/server/status", api.ServerStatusGET)
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

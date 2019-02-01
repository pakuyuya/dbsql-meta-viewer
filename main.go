package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pakuyuya/go-text-grep-app/controller/page"
	yaml "gopkg.in/yaml.v2"
)

type Setting struct {
	Port string
}

func main() {
	s, err := loadSetting()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// static
	gp := router.Group("/")
	{
		gp.GET("/", page.IndexGET)
	}

	server := &http.Server{
		Addr:           ":" + s.Port,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 18,
	}
	server.ListenAndServe()
}

func loadSetting() (*Setting, error) {
	bytes, err := ioutil.ReadFile(`./setting.yml`)
	if err != nil {
		return nil, err
	}

	s := Setting{}
	err = yaml.Unmarshal(bytes, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

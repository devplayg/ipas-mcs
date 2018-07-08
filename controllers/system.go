package controllers

import (
	"time"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
)

type SystemController struct {
	baseController
}

func (c *SystemController) GetNews() {

	news := map[string]interface{}{
		"time":     time.Now().Format(time.RFC3339),
		"message":  nil,
		"resource": nil,
	}

	filter := objs.Server{
		Category1: 1,
		Category2: 1,
	}
	server, err := models.GetServer(filter)
	if err != nil {
		log.Error(err)
	}
	news["resource"] = server

	// 리소스 조회

	// 메시지 조회

	c.Data["json"] = news
	c.ServeJSON()
}

package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"time"
)

type SystemController struct {
	baseController
}

func (c *SystemController) GetNews() {

	t := time.Now()

	// 뉴스
	news := map[string]interface{}{
		"time":     t.Format(time.RFC3339),
		"message":  nil,
		"resource": nil,
	}

	// 리소스 조회
	if server, err := models.GetServer(objs.Server{
		Category1: 1,
		Category2: 1,
	}); err != nil {
		log.Error(err)
	} else {
		news["resource"] = server
	}

	// 메시지 조회
	msgFilter := objs.MessageFilter{
		//PagingFilter: {
		//	StartDate: time.Now().Format(""),
		//},
		Status: objs.UnreadMessage,
	}
	msgFilter.EndDate = t.Format(objs.SearchTimeFormat)
	msgFilter.StartDate = t.AddDate(0, 0, -1).Format(objs.SearchTimeFormat)
	msgFilter.FastPaging = "on"
	msgFilter.Sort = "date"
	msgFilter.Order = "desc"
	msgFilter.Offset = 0
	msgFilter.Limit = 5
	if messages, _, err := models.GetMessage(msgFilter, c.member); err != nil {
		log.Error(err)
	} else {
		news["message"] = messages
	}

	c.Data["json"] = news
	c.ServeJSON()
}

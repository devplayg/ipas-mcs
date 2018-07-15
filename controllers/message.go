package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"time"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type MessageController struct {
	baseController
}

func (c *MessageController) CtrlPrepare() {
	// 추가 언어 키워드
	c.addToFrontLang("ipas.start,shock,speeding,proximity")

	// 권한 부여
	c.grant(objs.User)
}

func (c *MessageController) GetMessage() {
	filter := c.getFilter()
	messages, total, err := c.getMessage(filter, c.member)
	if err != nil {
		log.Error(err)
	}
	c.serveResultJson(messages, total, err, filter.FastPaging)
}

func (c *MessageController) GetUnreadMessage() {
	filter := c.getFilter()
	filter.Status = 1
	messages, total, err := c.getMessage(filter, c.member)
	if err != nil {
		log.Error(err)
	}
	c.serveResultJson(messages, total, err, filter.FastPaging)
}

func (c *MessageController) getMessage(filter objs.MessageFilter, member *objs.Member) ([]objs.Message, int64, error) {
	return models.GetMessage(filter, c.member)
}

func (c *MessageController) getFilter() objs.MessageFilter {

	// 요청값 분류
	filter := objs.MessageFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정
	if filter.StartDate == "" || filter.EndDate == "" {
		t := time.Now()
		filter.StartDate = t.AddDate(0, 0, -7).Format(objs.SearchTimeFormat)
		filter.EndDate = t.Format(objs.SearchTimeFormat)
	}

	// 페이징 처리
	if filter.Sort == "" {
		filter.Sort = "date"
	}
	if filter.Order == "" {
		filter.Order = "desc"
	}
	if filter.Limit < 1 {
		filter.Limit = 20
	}
	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}
	return filter
}

func (c *MessageController) GotIt() {
	messageId, _ := strconv.Atoi(c.Ctx.Input.Param(":messageId"))
	err := models.MarkMessageAsRead(messageId, c.member)
	if err != nil {
		log.Error(err)
	}
	c.ServeJSON()
}

func (c *MessageController) MarkAllAsRead() {
	err := models.MarkAllMessageAsRead(c.member)
	if err != nil {
		log.Error(err)
	}
	c.ServeJSON()
}

//func SendMessageToAll() error {
//	return nil
//}
//
//func SendMessage() error {
//	return nil
//}
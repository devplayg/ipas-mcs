package controllers

import (
	log "github.com/sirupsen/logrus"
)

type LoginController struct {
	baseController
}

func (c *LoginController) Get() {
	log.Info("test logging")
	c.setTpl("login.tpl")
}

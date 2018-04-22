package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
)

type ConfigController struct {
	baseController
}

func (c *ConfigController) Get() {
	var config = make(map[string]map[string]objs.MultiValue)
	rows, err := models.GetSystemConfig()
	if err == nil {
		for _, r := range rows {
			if m, ok := config[r.Section]; ok {
			} else {
				m = make(map[string]objs.MultiValue)
				config[r.Section] = m
			}

			config[r.Section][r.Keyword] = objs.MultiValue{r.ValueS, r.ValueN}
		}
	}
	c.Data["config"] = config
	c.setTpl("config.tpl")

}

func (c *ConfigController) Patch() {
	// 입력 파싱 및 수치 제한
	filter := c.getFilter()

	// 시스템
	config := make([]objs.SysConfig, 0)
	config = append(config,
		objs.SysConfig{"system",
			"data_retention_days",
			objs.MultiValue{"", filter.DataRetentionDays},
		})
	config = append(config, objs.SysConfig{
		"system",
		"use_namecard",
		objs.MultiValue{filter.UseNameCard, 0},
	})

	// 로그인
	config = append(config, objs.SysConfig{
		"login",
		"max_failed_login_attempts",
		objs.MultiValue{"", filter.MaxFailedLoginAttempts},
	})
	config = append(config, objs.SysConfig{
		"login",
		"allow_multiple_login",
		objs.MultiValue{filter.AllowMultipleLogin, 0},
	})
	err := models.UpdateSystemConfig(config)

	dbResult := objs.NewDbResult()
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
		dbResult.Message = c.Tr("msg.success_updated")
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}

func (c *ConfigController) getFilter() *objs.Configuration {
	filter := objs.Configuration{}
	if err := c.ParseForm(&filter); err != nil {
	}

	if filter.DataRetentionDays > 999 {
		filter.DataRetentionDays = 999
	} else if filter.DataRetentionDays < 60 {
		filter.DataRetentionDays = 60
	}

	return &filter
}

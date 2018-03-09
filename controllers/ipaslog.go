package controllers


type IpaslogController struct {
	baseController
}

func (c *IpaslogController) Get() {
	c.setTpl("ipaslog.tpl")
}


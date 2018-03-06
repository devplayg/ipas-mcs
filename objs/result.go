package objs

import "github.com/astaxie/beego/orm"

type Result struct {
	State   bool        `json:"state"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{false, "", ""}
}

type DbResult struct {
	State        bool         `json:"state"`
	Message      string       `json:"message"`
	Rows         []orm.Params `json:"rows"`
	AffectedRows int64        `json:"affected_rows"`
	Total        int64        `json:"total"`
	LastInsertId int64        `json:"last_insert_id"`
}

func newDbResult() *DbResult {
	return &DbResult{}
}

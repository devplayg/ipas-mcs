package objs

type Result struct {
	State   bool        `json:"state"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{false, "", ""}
}

type DbResult struct {
	State        bool        `json:"state"`
	Message      string      `json:"message"`
	Rows         interface{} `json:"rows"`
	Data         interface{} `json:"data"`
	AffectedRows int64       `json:"affected_rows"`
	Total        int64       `json:"total"`
	LastInsertId int64       `json:"last_insert_id"`
}

func NewDbResult() *DbResult {
	return &DbResult{false, "", nil, nil, 0, 0, 0}
}

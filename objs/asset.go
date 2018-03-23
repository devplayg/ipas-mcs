package objs

import (
	"time"
)

type Asset struct {
	AssetId     int
	Class       int `json:"-"`
	ParentId    int `json:"parent_id"`
	Name        string
	Type        string `json:"type"`
	Type1       int
	Type2       int
	Hostname    string
	IP          string    `json:"-"`
	Cidr        int       `json:"-"`
	Guid        string    `json:"-"`
	Mac         string    `json:"-"`
	Port        uint16    `json:"-"`
	PortSub1    uint16    `json:"-"`
	PortSub2    uint16    `json:"-"`
	Version     string    `json:"-"`
	Username    string    `json:"-"`
	Password    string    `json:"-"`
	Usage_cpu   float32   `json:"-"`
	Usage_mem   float32   `json:"-"`
	Usage_disk1 float32   `json:"-"`
	Usage_disk2 float32   `json:"-"`
	State       uint8     `json:"-"`
	N1          int       `json:"-"`
	N2          int       `json:"-"`
	S1          string    `json:"-"`
	S2          string    `json:"-"`
	F1          float32   `json:"-"`
	F2          float64   `json:"-"`
	Created     time.Time `json:"-"`
	Updated     time.Time `json:"-"`

	Text     string   `json:"text"`
	Id       string   `json:"id"` // for tree
	Icon     string   `json:"-"`
	Children []*Asset `json:"children"`
}

package objs

import (
	"time"
)

type treeState struct {
	Opened   bool `json:"opened"`
	Disabled bool `json:"disabled"`
	Selected bool `json:"selected"`
}

type Asset struct {
	AssetId  int    `json:"asset_id" form:"asset_id"`
	Class    int    `json:"-" form:"class"`
	ParentId int    `json:"parent_id" form:"parent_id"`
	Name     string `json:"name" form:"name"`
	Type     int    `json:"type"`
	Type1    int    `json:"-" form:"type1"`
	Type2    int    `json:"-" formn:"type2"`
	//Hostname    string
	//IP          string    `json:"-"`
	//Cidr        int       `json:"-"`
	//Guid        string    `json:"-"`
	//Mac         string    `json:"-"`
	//Port        uint16    `json:"-"`
	//PortSub1    uint16    `json:"-"`
	//PortSub2    uint16    `json:"-"`
	//Version     string    `json:"-"`
	//Username    string    `json:"-"`
	//Password    string    `json:"-"`
	//Usage_cpu   float32   `json:"-"`
	//Usage_mem   float32   `json:"-"`
	//Usage_disk1 float32   `json:"-"`
	//Usage_disk2 float32   `json:"-"`
	//State       uint8     `json:"-"`
	//N1          int       `json:"-"`
	//N2          int       `json:"-"`
	//S1          string    `json:"-"`
	//S2          string    `json:"-"`
	//F1          float32   `json:"-"`
	//F2          float64   `json:"-"`
	Created time.Time `json:"-"`
	Updated time.Time `json:"-"`

	Text     string    `json:"text"`
	Id       int       `json:"id"` // for tree
	Icon     string    `json:"-"`
	Children []*Asset  `json:"children"`
	State    treeState `json:"state"`
}

var AssetClass = map[int]string{
	1: "IPAS",
}

type AssetMap map[int]*Asset

func NewRootAsset(class int) *Asset {
	str := "Unknown"
	if _, ok := AssetClass[class]; ok {
		str = AssetClass[class]
	}

	root := Asset{
		AssetId:  0,
		Id:       0,
		Class:    1,
		ParentId: -1,
		Type1:    0,
		//Type2:    0,
		Text:     str,
		Children: nil,
	}

	return &root
}

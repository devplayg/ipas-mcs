package objs

import (
	"time"
)

type HighchartData struct {
	PointStart    int64         `json:"pointStart"`
	PointInterval time.Duration `json:"pointInterval"`
	Data          []interface{} `json:"data"`
	TimeMap       map[int64]int `json:"-"`
	//YAxis         int           `json:"yAxis"`
	//Type          string        `json:"type"`
	//Name          string        `json:"name"`
	//Index         int           `json:"index"`
	//DashStyle     string        `json:"dashStyle"`
	//LineWidth     int           `json:"lineWidth"`
}
//
//func NewHighchartsData(from, to time.Time, pointInterval int, loc *time.Location) *HighchartData {
//	spew.Dump(from)
//	spew.Dump(to)
//	spew.Dump(pointInterval)
//	spew.Dump(loc)
//	c := HighchartData{
//		PointStart:    from.In(loc).Unix() * 1000,
//		PointInterval: time.Duration(pointInterval) * 1000,
//		TimeMap:       make(map[int64]int),
//	}
//	for i := from; i.Before(to); i = i.Add(time.Duration(pointInterval) * time.Second) {
//		// 시간맵 초기화
//		c.TimeMap[i.Unix()] = 0
//	}
//
//	return &c
//}
//
//func (c *HighchartData) Sort() {
//	int64arr := make(Int64Slice, 0, len(c.TimeMap))
//	for k := range c.TimeMap {
//		int64arr = append(int64arr, k)
//	}
//	sort.Sort(int64arr)
//	c.Data = make([]interface{}, 0)
//	for _, k := range int64arr {
//		c.Data = append(c.Data, c.TimeMap[k])
//	}
//}

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// E-Chart
type EChartData struct {
	PointStart    int64         `json:"pointStart"`
	PointInterval time.Duration `json:"pointInterval"`
	Data          []interface{} `json:"data"`
	TimeMap       map[int64]int `json:"-"`
	//YAxis         int           `json:"yAxis"`
	//Type          string        `json:"type"`
	//Name          string        `json:"name"`
	//Index         int           `json:"index"`
	//DashStyle     string        `json:"dashStyle"`
	//LineWidth     int           `json:"lineWidth"`
}

func NewTimeLineData(from, to time.Time, pointInterval int) map[int64]int {
	m := make(map[int64]int)
	for i := from; i.Before(to); i = i.Add(time.Duration(pointInterval) * time.Second) {
		m[i.Unix()] = 0
	}
	return m
}

//
//func NewEChartData(from, to time.Time, pointInterval int, loc *time.Location) *EChartData {
//	spew.Dump(from)
//	spew.Dump(to)
//	spew.Dump(pointInterval)
//	spew.Dump(loc)
//	c := EChartData{
//		//PointStart:    from.In(loc).Unix() * 1000,
//		//PointInterval: time.Duration(pointInterval) * 1000,
//		TimeMap:       make(map[int64]int),
//	}
//	for i := from; i.Before(to); i = i.Add(time.Duration(pointInterval) * time.Second) {
//		// 시간맵 초기화
//		c.TimeMap[i.Unix()] = 0
//	}
//
//	return &c
//}
//
//func (c *EChartData) Sort() {
//	int64arr := make(int64arr, 0, len(c.TimeMap))
//	for k := range c.TimeMap {
//		int64arr = append(int64arr, k)
//	}
//	sort.Sort(int64arr)
//	c.Data = make([]interface{}, 0)
//	for _, k := range int64arr {
//		c.Data = append(c.Data, c.TimeMap[k])
//	}
//}

//
//type TimeLine struct {
//	Timestamp int64
//	Count     int64
//}
//type TimelineList []TimeLine
//
//func (p TimelineList) Len() int { return len(p) }
//func (p TimelineList) Less(i, j int) bool {
//	return p[i].Timestamp < p[j].Timestamp
//}
//func (p TimelineList) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}

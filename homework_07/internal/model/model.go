package model

type Record struct {
	Id     string `json:"id" redis:"id"`
	Data   string `json:"data" redis:"data"`
	Binary string
}

type Result struct {
	Name          string  `json:"name"`
	Size          int     `json:"size"`
	WriteDuration float64 `json:"write_duration"`
	ReadDuration  float64 `json:"read_duration"`
}

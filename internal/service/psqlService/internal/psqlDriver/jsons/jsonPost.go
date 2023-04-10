package jsons

type PostReq struct {
	Table string `json:"table"`
	Body  string `json:"body"`
}

type PostResp struct {
	Key int `json:"key"`
}

type GetReq struct {
	Table string `json:"table"`
	Key   int    `json:"key"`
}

type GetResp struct {
	Key  int    `json:"key"`
	Body string `json:"body"`
}

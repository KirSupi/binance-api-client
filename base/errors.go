package base

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

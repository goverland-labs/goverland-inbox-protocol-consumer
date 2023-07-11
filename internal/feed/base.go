package feed

type Treasury struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
}

type Strategy struct {
	Name    string `json:"name"`
	Network string `json:"network"`
}

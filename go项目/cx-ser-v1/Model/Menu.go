package Model

type Menu struct {
	Id           int    `json:"id"`
	Url          string `json:"url"`
	Path         string `json:"path"`
	Component    string `json:"component"`
	Name         string `json:"name"`
	Icon_cls     string `json:"icon_cls"`
	Keep_alive   int    `json:"keep_alive"`
	Require_auth int    `json:"require_auth"`
	Parent_id    int    `json:"parent_id"`
	Enabled      int    `json:"enabled"`
}

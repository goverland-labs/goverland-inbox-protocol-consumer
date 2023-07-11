package feed

type Voting struct {
	Delay       int     `json:"delay"`
	Period      int     `json:"period"`
	Type        string  `json:"type"`
	Quorum      float32 `json:"quorum"`
	Blind       bool    `json:"blind"`
	HideAbstain bool    `json:"hide_abstain"`
	Privacy     string  `json:"privacy"`
	Aliased     bool    `json:"aliased"`
}

type Dao struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Private        bool       `json:"private"`
	About          string     `json:"about"`
	Avatar         string     `json:"avatar"`
	Terms          string     `json:"terms"`
	Location       string     `json:"location"`
	Website        string     `json:"website"`
	Twitter        string     `json:"twitter"`
	Github         string     `json:"github"`
	Coingecko      string     `json:"coingecko"`
	Email          string     `json:"email"`
	Network        string     `json:"network"`
	Symbol         string     `json:"symbol"`
	Skin           string     `json:"skin"`
	Domain         string     `json:"domain"`
	Strategies     []Strategy `json:"strategies"`
	Admins         []string   `json:"admins"`
	Members        []string   `json:"members"`
	Moderators     []string   `json:"moderators"`
	Voting         Voting     `json:"voting"`
	Categories     []string   `json:"categories"`
	Treasures      []Treasury `json:"treasures"`
	FollowersCount int        `json:"followers_count"`
	ProposalsCount int        `json:"proposals_count"`
	Guidelines     string     `json:"guidelines"`
	Template       string     `json:"template"`
	ParentID       string     `json:"parent_id"`
}

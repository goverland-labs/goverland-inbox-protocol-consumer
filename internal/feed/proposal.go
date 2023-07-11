package feed

type Proposal struct {
	ID            string     `json:"id"`
	Ipfs          string     `json:"ipfs"`
	Author        string     `json:"author"`
	Created       int        `json:"created"`
	DaoID         string     `json:"dao_id"`
	Network       string     `json:"network"`
	Symbol        string     `json:"symbol"`
	Type          string     `json:"type"`
	Strategies    []Strategy `json:"strategies"`
	Title         string     `json:"title"`
	Body          string     `json:"body"`
	Discussion    string     `json:"discussion"`
	Choices       []string   `json:"choices"`
	Start         int        `json:"start"`
	End           int        `json:"end"`
	Quorum        float64    `json:"quorum"`
	Privacy       string     `json:"privacy"`
	Snapshot      string     `json:"snapshot"`
	State         string     `json:"state"`
	Link          string     `json:"link"`
	App           string     `json:"app"`
	Scores        []float32  `json:"scores"`
	ScoresState   string     `json:"scores_state"`
	ScoresTotal   float32    `json:"scores_total"`
	ScoresUpdated int        `json:"scores_updated"`
	Votes         int        `json:"votes"`
}
package contract

type NewCampaing struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Emails  []string `json:"emails"`
}

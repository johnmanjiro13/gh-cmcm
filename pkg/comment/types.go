package comment

type Comment struct {
	Body    string `json:"body"`
	Author  string `json:"author"`
	HTMLURL string `json:"html_url"`
}

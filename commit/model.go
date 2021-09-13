package commit

type CommitModel struct {
	Id             string   `json:"id"`              //sample : "ed899a2f4b50b4370feeea94676502b42383c746",
	ShortId        string   `json:"short_id"`        //sample : "ed899a2f4b5",
	Title          string   `json:"title"`           //sample : "Replace sanitize with escape once",
	AuthorName     string   `json:"author_name"`     //sample : "Example User",
	AuthorEmail    string   `json:"author_email"`    //sample : "user@example.com",
	AuthoredDate   string   `json:"authored_date"`   //sample : "2012-09-20T11:50:22+03:00",
	CommitterName  string   `json:"committer_name"`  //sample : "Administrator",
	CommitterEmail string   `json:"committer_email"` //sample : "admin@example.com",
	CommittedDate  string   `json:"committed_date"`  //sample : "2012-09-20T11:50:22+03:00",
	CreatedAt      string   `json:"created_at"`      //sample : "2012-09-20T11:50:22+03:00",
	Message        string   `json:"message"`         //sample : "Replace sanitize with escape once",
	ParentIds      []string `json:"parent_ids"`      //sample : ["6104942438c14ec7bd21c6cd5bd995272b3faff6"]
	WebUrl         string   `json:"web_url"`         //sample : "https://gitlab.example.com/thedude/gitlab-foss/-/commit/ed899a2f4b50b4370feeea94676502b42383c746"
}

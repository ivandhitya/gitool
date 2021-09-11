package tag

type TagModel struct {
	Commit struct {
		ID             string   `json:"id"`
		ShortID        string   `json:"short_id"`
		Title          string   `json:"title"`
		CreatedAt      string   `json:"created_at"`
		ParentIDs      []string `json:"parent_ids"`
		Message        string   `json:"message"`
		AuthorName     string   `json:"author_name"`
		AuthorEmail    string   `json:"author_email"`
		AuthoredDate   string   `json:"authored_date"`
		CommitterName  string   `json:"committer_name"`
		CommitterEmail string   `json:"committer_email"`
		CommittedDate  string   `json:"committed_date"`
	} `json:"commit"`
	Release struct {
		TagName     string `json:"tag_name"`
		Description string `json:"description"`
	} `json:"release"`
	Name      string  `json:"name"`
	Target    string  `json:"target"`
	Message   *string `json:"message"`
	Protected bool    `json:"protected"`
}

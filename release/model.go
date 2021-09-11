package release

type ReleaseModel struct {
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	ReleasedAt  string `json:"released_at"`
	Author      struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"author"`
	Commit struct {
		Id             string   `json:"id"`
		ShortId        string   `json:"short_id"`
		Title          string   `json:"title"`
		CreatedAt      string   `json:"created_at"`
		ParentIds      []string `json:"parent_ids"`
		Message        string   `json:"message"`
		AuthorName     string   `json:"author_name"`
		AuthorEmail    string   `json:"author_email"`
		AuthoredDate   string   `json:"authored_date"`
		CommitterName  string   `json:"committer_name"`
		CommitterEmail string   `json:"committer_email"`
		CommittedDate  string   `json:"committed_date"`
	} `json:"commit"`
	Milestones []struct {
		Id          int    `json:"id"`
		Iid         int    `json:"iid"`
		ProjectId   int    `json:"project_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		State       string `json:"state"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
		DueDate     string `json:"due_date"`
		StartDate   string `json:"start_date"`
		WebUrl      string `json:"web_url"`
		IssueStats  struct {
			Opened int `json:"opened"`
			Closed int `json:"closed"`
		} `json:"issue_stats"`
	} `json:"milestones"`

	CommitPath  string `json:"commit_path"`
	TagPath     string `json:"tag_path"`
	EvidenceSha string `json:"evidence_sha"`
	Assets      struct {
		Count   int `json:"count"`
		Sources []struct {
			Format string `json:"format"`
			Url    string `json:"url"`
		} `json:"sources"`
		Links            []string `json:"links"`
		EvidenceFilePath string   `json:"evidence_file_path"`
	} `json:"assets"`
}

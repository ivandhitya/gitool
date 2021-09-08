package mr

// RespMR is a type for merge request response
type RespMR struct {
	ID          int      `json:"id"`
	IID         int      `json:"iid"`
	ProjectID   int      `json:"project_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Message     []string `json:"message"`
	User        struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
}

type RespAcceptMR struct {
	Title        string   `json:"title"`
	TargetBranch string   `json:"target_branch"`
	SourceBranch string   `json:"source_branch"`
	Message      []string `json:"message"`
	MergedAt     string   `json:"merged_at"`
	MergeError   string   `json:"merge_error"`
	State        string   `json:"state"`
}

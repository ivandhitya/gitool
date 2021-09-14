package commit

type ApprovalRuleModel struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	RuleType          string `json:"rule_type"`
	EligibleApprovers []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"eligible_approvers"`

	ApprovalsRequired int `json:"approvals_required"`
	Users             []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"users"`

	Groups               []interface{} `json:"groups"`
	ContainsHiddenGroups bool          `json:"contains_hidden_groups"`
	Section              interface{}   `json:"section"`
	SourceRule           struct {
		ApprovalsRequired int `json:"approvals_required"`
	} `json:"source_rule"`
	Overridden bool `json:"overridden"`
}

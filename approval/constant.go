package commit

const (
	KeyDataID                = "id"                 //integer or string	yes	The ID or URL-encoded path of a project
	KeyDataMergeRequestIID   = "merge_request_iid"  //integer	yes	The IID of MR
	KeyDataName              = "name"               //string	yes	The name of the approval rule
	KeyDataApprovalsRequired = "approvals_required" //integer	yes	The number of required approvals for this rule
	KeyDataUserIDs           = "user_ids"           //Array	no	The ids of users as approvers
	KeyDataGroupIDs          = "group_ids"          //Array	no	The ids of groups as approvers
)

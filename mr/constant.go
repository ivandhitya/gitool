package mr

const (
	KeyDataID                        = "id"                           //yes The project ID
	KeyDataSourceBranch              = "source_branch"                //yes	The source branch.
	KeyDataTargetBranch              = "target_branch"                //yes	The target branch.
	KeyDataTitle                     = "title"                        //yes	Title of MR.
	KeyDataAssigneeID                = "assignee_id"                  //no	Assignee user ID.
	KeyDataAssigneeIDs               = "assignee_ids"                 //array	no	The ID of the user(s) to assign the MR to. Set to 0 or provide an empty value to unassign all assignees.
	KeyDataReviewerIDs               = "reviewer_ids"                 //array	no	The ID of the user(s) added as a reviewer to the MR. If set to 0 or left empty, no reviewers are added.
	KeyDataDescription               = "description"                  //no	Description of MR. Limited to 1,048,576 characters.
	KeyDataTargetProjectID           = "target_project_id"            //no	The target project (numeric ID).
	KeyDataLabels                    = "labels"                       //no	Labels for MR as a comma-separated list.
	KeyDataMilestoneID               = "milestone_id"                 //no	The global ID of a milestone.
	KeyDataRemoveSourceBranch        = "remove_source_branch"         //no	Flag indicating if a merge request should remove the source branch when merging.
	KeyDataAllowCollaboration        = "allow_collaboration"          //no	Allow commits from members who can merge to the target branch.
	KeyDataAllowMaintainerToPush     = "allow_maintainer_to_push"     //no	Deprecated, see allow_collaboration.
	KeyDataSquash                    = "squash"                       //no	Squash commits into a single commit when merging.
	KeyDataApprovalsRequired         = "approvals_required"           //Approvals required before MR can be merged. Deprecated in 12.0 in favor of Approval Rules API
	KeyDataMergeRequestIID           = "merge_request_iid"            //The IID of MR
	KeyDataMergeCommitMessage        = "merge_commit_message"         //Custom merge commit message
	KeyDataSquashCommitMessage       = "squash_commit_message"        //Custom squash commit message
	KeyDataShouldRemoveSourceBranch  = "should_remove_source_branch"  //If true removes the source branch.
	KeyDataMergeWhenPipelineSucceeds = "merge_when_pipeline_succeeds" //If true the MR is merged when the pipeline succeeds.
	KeyDataSHA                       = "sha"                          //If present, then this SHA must match the HEAD of the source branch, otherwise the merge fails.
)

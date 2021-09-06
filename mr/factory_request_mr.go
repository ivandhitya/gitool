package mr

import (
	"strconv"
)

// ReqMR is a type for merge request data form
type ReqMR map[string]string

// AddProjectID is a function for Add ProjectID value in ReqMR type
func (r *ReqMR) AddProjectID(projectID int) *ReqMR {
	(*r)[KeyDataID] = strconv.Itoa(projectID)
	return r
}

// AddSourceBranch is a function for Add SourceBranch value in ReqMR type
func (r *ReqMR) AddSourceBranch(sourceBranch string) *ReqMR {
	(*r)[KeyDataSourceBranch] = sourceBranch
	return r
}

// AddTargetBranch is a function for Add TargetBranch value in ReqMR type
func (r *ReqMR) AddTargetBranch(targetBranch string) *ReqMR {
	(*r)[KeyDataTargetBranch] = targetBranch
	return r
}

// AddTitle is a function for Add Title value in ReqMR type
func (r *ReqMR) AddTitle(title string) *ReqMR {
	(*r)[KeyDataTitle] = title
	return r
}

// AddAssigneeID is a function for Add AssigneeID value in ReqMR type
func (r *ReqMR) AddAssigneeID(assigneeID int) *ReqMR {
	(*r)[KeyDataAssigneeID] = strconv.Itoa(assigneeID)
	return r
}

// AddAssigneeIDs is a function for Add AssigneeIDs value in ReqMR type
func (r *ReqMR) AddAssigneeIDs(assigneeIDs int) *ReqMR {
	(*r)[KeyDataAssigneeIDs] = strconv.Itoa(assigneeIDs)
	return r
}

// AddReviewerIDs is a function for Add ReviewerIDs value in ReqMR type
func (r *ReqMR) AddReviewerIDs(reviewerIDs int) *ReqMR {
	(*r)[KeyDataReviewerIDs] = strconv.Itoa(reviewerIDs)
	return r
}

// AddDescription is a function for Add Description value in ReqMR type
func (r *ReqMR) AddDescription(description string) *ReqMR {
	(*r)[KeyDataDescription] = description
	return r
}

// AddTargetProjectID is a function for Add TargetProjectID value in ReqMR type
func (r *ReqMR) AddTargetProjectID(targetProjectID int) *ReqMR {
	(*r)[KeyDataTargetProjectID] = strconv.Itoa(targetProjectID)
	return r
}

// AddLabels is a function for Add Labels value in ReqMR type
func (r *ReqMR) AddLabels(labels string) *ReqMR {
	(*r)[KeyDataLabels] = labels
	return r
}

// AddMilestoneID is a function for Add MilestoneID value in ReqMR type
func (r *ReqMR) AddMilestoneID(milestoneID int) *ReqMR {
	(*r)[KeyDataMilestoneID] = strconv.Itoa(milestoneID)
	return r
}

// AddRemoveSourceBranch is a function for Add RemoveSourceBranch value in ReqMR type
func (r *ReqMR) AddRemoveSourceBranch(removeSourceBranch bool) *ReqMR {
	(*r)[KeyDataRemoveSourceBranch] = strconv.FormatBool(removeSourceBranch)
	return r
}

// AddAllowCollaboration is a function for Add AllowCollaboration value in ReqMR type
func (r *ReqMR) AddAllowCollaboration(allowCollaboration bool) *ReqMR {
	(*r)[KeyDataAllowCollaboration] = strconv.FormatBool(allowCollaboration)
	return r
}

// AddAllowMaintainerToPush is a function for Add AllowMaintainerToPush value in ReqMR type
func (r *ReqMR) AddAllowMaintainerToPush(allowMaintainerToPush bool) *ReqMR {
	(*r)[KeyDataAllowMaintainerToPush] = strconv.FormatBool(allowMaintainerToPush)
	return r
}

// AddSquash is a function for Add Squash value in ReqMR type
func (r *ReqMR) AddSquash(squash bool) *ReqMR {
	(*r)[KeyDataSquash] = strconv.FormatBool(squash)
	return r
}

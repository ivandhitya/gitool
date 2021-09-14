package mr

import (
	"fmt"
	"strconv"
)

// ReqMR is a type for merge request data form
type ReqMR map[string]string

// AddProjectID is a function for Add ProjectID value in ReqMR type (*required)
func (r *ReqMR) AddProjectID(val interface{}) *ReqMR {
	(*r)[KeyDataID] = fmt.Sprintf("%v", val)
	return r
}

// AddSourceBranch is a function for Add SourceBranch value in ReqMR type (*required)
func (r *ReqMR) AddSourceBranch(val string) *ReqMR {
	(*r)[KeyDataSourceBranch] = val
	return r
}

// AddTargetBranch is a function for Add TargetBranch value in ReqMR type (*required)
func (r *ReqMR) AddTargetBranch(val string) *ReqMR {
	(*r)[KeyDataTargetBranch] = val
	return r
}

// AddTitle is a function for Add Title value in ReqMR type (*required)
func (r *ReqMR) AddTitle(val string) *ReqMR {
	(*r)[KeyDataTitle] = val
	return r
}

// AddAssigneeID is a function for Add AssigneeID value in ReqMR type
func (r *ReqMR) AddAssigneeID(val int) *ReqMR {
	(*r)[KeyDataAssigneeID] = strconv.Itoa(val)
	return r
}

// AddAssigneeIDs is a function for Add AssigneeIDs value in ReqMR type
func (r *ReqMR) AddAssigneeIDs(val int) *ReqMR {
	(*r)[KeyDataAssigneeIDs] = strconv.Itoa(val)
	return r
}

// AddReviewerIDs is a function for Add ReviewerIDs value in ReqMR type
func (r *ReqMR) AddReviewerIDs(val int) *ReqMR {
	(*r)[KeyDataReviewerIDs] = strconv.Itoa(val)
	return r
}

// AddDescription is a function for Add Description value in ReqMR type
func (r *ReqMR) AddDescription(val string) *ReqMR {
	(*r)[KeyDataDescription] = val
	return r
}

// AddTargetProjectID is a function for Add TargetProjectID value in ReqMR type
func (r *ReqMR) AddTargetProjectID(val int) *ReqMR {
	(*r)[KeyDataTargetProjectID] = strconv.Itoa(val)
	return r
}

// AddLabels is a function for Add Labels value in ReqMR type
func (r *ReqMR) AddLabels(val string) *ReqMR {
	(*r)[KeyDataLabels] = val
	return r
}

// AddMilestoneID is a function for Add MilestoneID value in ReqMR type
func (r *ReqMR) AddMilestoneID(val int) *ReqMR {
	(*r)[KeyDataMilestoneID] = strconv.Itoa(val)
	return r
}

// AddRemoveSourceBranch is a function for Add RemoveSourceBranch value in ReqMR type
func (r *ReqMR) AddRemoveSourceBranch(val bool) *ReqMR {
	(*r)[KeyDataRemoveSourceBranch] = strconv.FormatBool(val)
	return r
}

// AddAllowCollaboration is a function for Add AllowCollaboration value in ReqMR type
func (r *ReqMR) AddAllowCollaboration(val bool) *ReqMR {
	(*r)[KeyDataAllowCollaboration] = strconv.FormatBool(val)
	return r
}

// AddAllowMaintainerToPush is a function for Add AllowMaintainerToPush value in ReqMR type
func (r *ReqMR) AddAllowMaintainerToPush(val bool) *ReqMR {
	(*r)[KeyDataAllowMaintainerToPush] = strconv.FormatBool(val)
	return r
}

// AddSquash is a function for Add Squash value in ReqMR type
func (r *ReqMR) AddSquash(val bool) *ReqMR {
	(*r)[KeyDataSquash] = strconv.FormatBool(val)
	return r
}

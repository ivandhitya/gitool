package mr

import "strconv"

// ReqAcceptMR is a type for accept merge request data form
type ReqAcceptMR map[string]string

// AddProjectID is a function for Add ProjectID value in ReqAcceptMR type (*required)
func (r *ReqAcceptMR) AddProjectID(projectID int) *ReqAcceptMR {
	(*r)[KeyDataID] = strconv.Itoa(projectID)
	return r
}

//AddMergeRequestIID is a fuction for Add MergeRequestIID value in ReqAcceptMR type (*required)
func (r *ReqAcceptMR) AddMergeRequestIID(val int) *ReqAcceptMR {
	(*r)[KeyDataMergeRequestIID] = strconv.Itoa(val)
	return r
}

//AddMergeCommitMessage is a fuction for Add MergeCommitMessage value in ReqAcceptMR type
func (r *ReqAcceptMR) AddMergeCommitMessage(val string) *ReqAcceptMR {
	(*r)[KeyDataMergeCommitMessage] = val
	return r
}

//AddSquashCommitMessage is a fuction for Add SquashCommitMessage value in ReqAcceptMR type
func (r *ReqAcceptMR) AddSquashCommitMessage(val string) *ReqAcceptMR {
	(*r)[KeyDataSquashCommitMessage] = val
	return r
}

//AddShouldRemoveSourceBranch is a fuction for Add ShouldRemoveSourceBranch value in ReqAcceptMR type
func (r *ReqAcceptMR) AddShouldRemoveSourceBranch(val string) *ReqAcceptMR {
	(*r)[KeyDataShouldRemoveSourceBranch] = val
	return r
}

//AddMergeWhenPipelineSucceeds is a fuction for Add MergeWhenPipelineSucceeds value in ReqAcceptMR type
func (r *ReqAcceptMR) AddMergeWhenPipelineSucceeds(val string) *ReqAcceptMR {
	(*r)[KeyDataMergeWhenPipelineSucceeds] = val
	return r
}

//AddSHA is a fuction for Add SHA value in ReqAcceptMR type
func (r *ReqAcceptMR) AddSHA(val string) *ReqAcceptMR {
	(*r)[KeyDataSHA] = val
	return r
}

package tag

import (
	"strconv"
)

type ReqGetTagList map[string]string

// AddProjectID is a function for Add ProjectID value in ReqGetTagList type (*required)
func (r *ReqGetTagList) AddProjectID(projectID int) *ReqGetTagList {
	(*r)[KeyDataID] = strconv.Itoa(projectID)
	return r
}

// AddOrderBy is a function for Add OrderBy value in ReqGetTagList type
func (r *ReqGetTagList) AddOrderBy(val OrderBy) *ReqGetTagList {
	(*r)[KeyDataOrderBy] = string(val)
	return r
}

// AddSort is a function for Add Sort value in ReqGetTagList type
func (r *ReqGetTagList) AddSort(val Sort) *ReqGetTagList {
	(*r)[KeyDataSort] = string(val)
	return r
}

// AddSearch is a function for Add Search value in ReqGetTagList type
func (r *ReqGetTagList) AddSearch(val string) *ReqGetTagList {
	(*r)[KeyDataSearch] = val
	return r
}

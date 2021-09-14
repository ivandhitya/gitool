package commit

import (
	"fmt"
	"strconv"
)

type ReqGetCommitList map[string]string

// AddProjectID is a function for Add ProjectID value in ReqGetCommitList type (*required)
func (r *ReqGetCommitList) AddProjectID(val interface{}) *ReqGetCommitList {
	(*r)[KeyDataID] = fmt.Sprintf("%v", val)
	return r
}

func (r *ReqGetCommitList) AddRefName(val string) *ReqGetCommitList {
	(*r)[KeyDataRefName] = val
	return r
}

func (r *ReqGetCommitList) AddSince(val string) *ReqGetCommitList {
	(*r)[KeyDataSince] = val
	return r
}

func (r *ReqGetCommitList) AddUntil(val string) *ReqGetCommitList {
	(*r)[KeyDataUntil] = val
	return r
}

func (r *ReqGetCommitList) AddPath(val string) *ReqGetCommitList {
	(*r)[KeyDataPath] = val
	return r
}

func (r *ReqGetCommitList) AddAll(val bool) *ReqGetCommitList {
	(*r)[KeyDataAll] = strconv.FormatBool(val)
	return r
}

func (r *ReqGetCommitList) AddWithStats(val bool) *ReqGetCommitList {
	(*r)[KeyDataWithStats] = strconv.FormatBool(val)
	return r
}

func (r *ReqGetCommitList) AddFirstParent(val bool) *ReqGetCommitList {
	(*r)[KeyDataFirstParent] = strconv.FormatBool(val)
	return r
}

func (r *ReqGetCommitList) AddOrder(val string) *ReqGetCommitList {
	(*r)[KeyDataOrder] = val
	return r
}

func (r *ReqGetCommitList) AddTrailers(val bool) *ReqGetCommitList {
	(*r)[KeyDataTrailers] = strconv.FormatBool(val)
	return r
}

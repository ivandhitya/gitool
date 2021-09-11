package release

import "strconv"

type ReqCreateRelease map[string]string

// AddProjectID is a function for Add ProjectID value in ReqCreateRelease type (*required)
func (r *ReqCreateRelease) AddProjectID(projectID int) *ReqCreateRelease {
	(*r)[KeyDataID] = strconv.Itoa(projectID)
	return r
}

func (r *ReqCreateRelease) AddTagName(val string) *ReqCreateRelease {
	(*r)[KeyDataTagName] = val
	return r
}

func (r *ReqCreateRelease) AddName(val string) *ReqCreateRelease {
	(*r)[KeyDataName] = val
	return r
}

func (r *ReqCreateRelease) AddDescription(val string) *ReqCreateRelease {
	(*r)[KeyDataDescription] = val
	return r
}

func (r *ReqCreateRelease) AddReleasedAt(val string) *ReqCreateRelease {
	(*r)[KeyDataReleasedAt] = val
	return r
}

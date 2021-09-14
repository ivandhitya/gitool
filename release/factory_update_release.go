package release

import "fmt"

type ReqCreateRelease map[string]string

// AddProjectID is a function for Add ProjectID value in ReqCreateRelease type (*required)
func (r *ReqCreateRelease) AddProjectID(val interface{}) *ReqCreateRelease {
	(*r)[KeyDataID] = fmt.Sprintf("%v", val)
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

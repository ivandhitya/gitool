package release

import "fmt"

type ReqUpdateRelease map[string]string

// AddProjectID is a function for Add ProjectID value in ReqUpdateRelease type (*required)
func (r *ReqUpdateRelease) AddProjectID(val interface{}) *ReqUpdateRelease {
	(*r)[KeyDataID] = fmt.Sprintf("%v", val)
	return r
}

func (r *ReqUpdateRelease) AddTagName(val string) *ReqUpdateRelease {
	(*r)[KeyDataTagName] = val
	return r
}

func (r *ReqUpdateRelease) AddName(val string) *ReqUpdateRelease {
	(*r)[KeyDataName] = val
	return r
}

func (r *ReqUpdateRelease) AddRef(val string) *ReqUpdateRelease {
	(*r)[KeyDataRef] = val
	return r
}

func (r *ReqUpdateRelease) AddDescription(val string) *ReqUpdateRelease {
	(*r)[KeyDataDescription] = val
	return r
}

func (r *ReqUpdateRelease) AddReleasedAt(val string) *ReqUpdateRelease {
	(*r)[KeyDataReleasedAt] = val
	return r
}

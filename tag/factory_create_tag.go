package tag

import "fmt"

type ReqCreateTag map[string]string

// AddProjectID is a function for Add ProjectID value in ReqCreateTag type (*required)
func (r *ReqCreateTag) AddProjectID(val interface{}) *ReqCreateTag {
	(*r)[KeyDataID] = fmt.Sprintf("%v", val)
	return r
}

func (r *ReqCreateTag) AddTagName(val string) *ReqCreateTag {
	(*r)[KeyDataTagName] = val
	return r
}
func (r *ReqCreateTag) AddRef(val string) *ReqCreateTag {
	(*r)[KeyDataRef] = val
	return r
}
func (r *ReqCreateTag) AddMessage(val string) *ReqCreateTag {
	(*r)[KeyDataMessage] = val
	return r
}

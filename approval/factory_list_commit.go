package commit

import (
	"fmt"
)

type ReqUpdateMRApprovalRules map[string]string

// AddApprovalsRequired is a function for Add ApprovalsRequired value in ReqUpdateMRApprovalRules type (*required)
func (r *ReqUpdateMRApprovalRules) AddApprovalsRequired(val int) *ReqUpdateMRApprovalRules {
	(*r)[KeyDataApprovalsRequired] = fmt.Sprintf("%v", val)
	return r
}

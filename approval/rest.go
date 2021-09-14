package commit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ivandhitya/gitool/model"
	"github.com/sirupsen/logrus"
)

const (
	PATH_GET_MR_APPROVAL_RULE    = "%s/api/v4/projects/%d/merge_requests/%d/approval_rules"
	PATH_UPDATE_MR_APPROVAL_RULE = "%s/api/v4/projects/%d/merge_requests/%d/approval_rules/%d"
)

type RestMRApprovalRule struct {
	gitConfig *model.GitConfig
}

func NewRestMRApprovalRule(gitConfig *model.GitConfig) RestMRApprovalRule {
	return RestMRApprovalRule{gitConfig}
}

func (r *RestMRApprovalRule) GetMRApprovalRules(projectID interface{}, mergeRequestIID int) ([]ApprovalRuleModel, error) {
	path := fmt.Sprintf(PATH_GET_MR_APPROVAL_RULE, r.gitConfig.URL, projectID, mergeRequestIID)
	resp := []ApprovalRuleModel{}

	respOrigin, err := r.gitConfig.Client.R().SetAuthToken(r.gitConfig.Token).Get(path)
	if err != nil {
		logrus.Trace(err)
		return resp, err
	}
	respBody := respOrigin.Body()
	if err := json.Unmarshal(respBody, &resp); err != nil {
		errr := errors.New(fmt.Sprintf("%v %s", err, string(respBody)))
		logrus.Trace(errr)
		return resp, errr
	}
	status := respOrigin.StatusCode()
	if status != http.StatusAccepted && status != http.StatusCreated && status != http.StatusOK {
		errr := errors.New(fmt.Sprintf("http status: %s", respOrigin.Status()))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

func (r *RestMRApprovalRule) UpdateMRApprovalRules(projectID interface{}, mergeRequestIID int, approvalRuleID int, formData ReqUpdateMRApprovalRules) (ApprovalRuleModel, error) {
	path := fmt.Sprintf(PATH_UPDATE_MR_APPROVAL_RULE, r.gitConfig.URL, projectID, mergeRequestIID, approvalRuleID)
	resp := ApprovalRuleModel{}

	respOrigin, err := r.gitConfig.Client.R().SetAuthToken(r.gitConfig.Token).SetFormData(formData).Put(path)
	if err != nil {
		logrus.Trace(err)
		return resp, err
	}
	respBody := respOrigin.Body()
	if err := json.Unmarshal(respBody, &resp); err != nil {
		errr := errors.New(fmt.Sprintf("%v %s", err, string(respBody)))
		logrus.Trace(errr)
		return resp, errr
	}
	status := respOrigin.StatusCode()
	if status != http.StatusAccepted && status != http.StatusCreated && status != http.StatusOK {
		errr := errors.New(fmt.Sprintf("http status: %s", respOrigin.Status()))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

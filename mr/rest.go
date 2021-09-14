package mr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ivandhitya/gitool/model"
	"github.com/sirupsen/logrus"
)

const (
	PATH_MR        = "%s/api/v4/projects/%d/merge_requests"
	PATH_ACCEPT_MR = "%s/api/v4/projects/%d/merge_requests/%d/merge"
	PATH_DELETE_MR = "%s/api/v4/projects/%d/merge_requests/%d"
)

type RestMergeRequest struct {
	gitConfig *model.GitConfig
}

func NewRestMergeRequest(gitConfig *model.GitConfig) RestMergeRequest {
	return RestMergeRequest{gitConfig}
}

func (mr *RestMergeRequest) CreateMR(projectID interface{}, formData ReqMR) (RespMR, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_MR, mr.gitConfig.URL, projectID)
	resp := RespMR{}

	respOrigin, err := mr.gitConfig.Client.R().SetAuthToken(mr.gitConfig.Token).SetFormData(formData).Post(path)
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
		errr := errors.New(fmt.Sprintf("http status: %s, message: %v", respOrigin.Status(), resp.Message))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

func (mr *RestMergeRequest) AcceptMR(projectID interface{}, mergeIID int, formData ReqAcceptMR) (RespAcceptMR, error) {
	path := fmt.Sprintf(PATH_ACCEPT_MR, mr.gitConfig.URL, projectID, mergeIID)
	resp := RespAcceptMR{}

	respOrigin, err := mr.gitConfig.Client.R().SetAuthToken(mr.gitConfig.Token).SetFormData(formData).Put(path)
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
		errr := errors.New(fmt.Sprintf("http status: %s, message: %v", respOrigin.Status(), resp.Message))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

func (mr *RestMergeRequest) DeleteMR(projectID interface{}, mergeIID int) (RespAcceptMR, error) {
	path := fmt.Sprintf(PATH_DELETE_MR, mr.gitConfig.URL, projectID, mergeIID)
	resp := RespAcceptMR{}

	respOrigin, err := mr.gitConfig.Client.R().SetAuthToken(mr.gitConfig.Token).Delete(path)
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
		errr := errors.New(fmt.Sprintf("http status: %s, message: %v", respOrigin.Status(), resp.Message))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

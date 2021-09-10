package mr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

const (
	PATH_MR        = "%s/api/v4/projects/%d/merge_requests"
	PATH_ACCEPT_MR = "%s/api/v4/projects/%d/merge_requests/%d/merge"
	PATH_DELETE_MR = "%s/api/v4/projects/%d/merge_requests/%d"
)

type RestMergeRequest struct {
	url    string
	token  string
	client *resty.Client
}

func NewRestMergeRequest(url, token string, client *resty.Client) RestMergeRequest {
	return RestMergeRequest{url, token, client}
}

func (mr *RestMergeRequest) CreateMR(projectID int, formData ReqMR) (RespMR, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_MR, mr.url, projectID)
	resp := RespMR{}

	respOrigin, err := mr.client.R().SetAuthToken(mr.token).SetFormData(formData).Post(path)
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

func (mr *RestMergeRequest) AcceptMR(projectID, mergeIID int, formData ReqAcceptMR) (RespAcceptMR, error) {
	path := fmt.Sprintf(PATH_ACCEPT_MR, mr.url, projectID, mergeIID)
	resp := RespAcceptMR{}

	respOrigin, err := mr.client.R().SetAuthToken(mr.token).SetFormData(formData).Put(path)
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

func (mr *RestMergeRequest) DeleteMR(projectID, mergeIID int) (RespAcceptMR, error) {
	path := fmt.Sprintf(PATH_DELETE_MR, mr.url, projectID, mergeIID)
	resp := RespAcceptMR{}

	respOrigin, err := mr.client.R().SetAuthToken(mr.token).Delete(path)
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

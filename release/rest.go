package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

const (
	PATH_CREATE_RELEASE = "%s/api/v4/projects/%d/releases"
	PATH_DEF_RELEASE    = "%s/api/v4/projects/%d/releases/%s"
)

type RestRelease struct {
	url    string
	token  string
	client *resty.Client
}

func NewRestRelease(url, token string, client *resty.Client) RestRelease {
	return RestRelease{url, token, client}
}

func (r *RestRelease) UpdateRelease(projectID int, tagName string, formData ReqUpdateRelease) (ReleaseModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_DEF_RELEASE, r.url, projectID, tagName)
	resp := ReleaseModel{}

	respOrigin, err := r.client.R().SetAuthToken(r.token).SetFormData(formData).Put(path)
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

func (r *RestRelease) CreateRelease(projectID int, formData ReqCreateRelease) (ReleaseModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_CREATE_RELEASE, r.url, projectID)
	resp := ReleaseModel{}

	respOrigin, err := r.client.R().SetAuthToken(r.token).SetFormData(formData).Post(path)
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

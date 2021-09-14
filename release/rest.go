package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ivandhitya/gitool/model"
	"github.com/sirupsen/logrus"
)

const (
	PATH_CREATE_RELEASE = "%s/api/v4/projects/%d/releases"
	PATH_DEF_RELEASE    = "%s/api/v4/projects/%d/releases/%s"
)

type RestRelease struct {
	gitConfig *model.GitConfig
}

func NewRestRelease(gitConfig *model.GitConfig) RestRelease {
	return RestRelease{gitConfig}
}

func (r *RestRelease) UpdateRelease(projectID interface{}, tagName string, formData ReqUpdateRelease) (ReleaseModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_DEF_RELEASE, r.gitConfig.URL, projectID, tagName)
	resp := ReleaseModel{}

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

func (r *RestRelease) CreateRelease(projectID interface{}, formData ReqCreateRelease) (ReleaseModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_CREATE_RELEASE, r.gitConfig.URL, projectID)
	resp := ReleaseModel{}

	respOrigin, err := r.gitConfig.Client.R().SetAuthToken(r.gitConfig.Token).SetFormData(formData).Post(path)
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

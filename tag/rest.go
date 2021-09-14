package tag

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ivandhitya/gitool/model"
	"github.com/sirupsen/logrus"
)

const (
	PATH_DEF_TAG = "%s/api/v4/projects/%d/repository/tags"
)

type RestTag struct {
	gitConfig *model.GitConfig
}

func NewRestTag(gitConfig *model.GitConfig) RestTag {
	return RestTag{gitConfig}
}

func (r *RestTag) GetAllTag(projectID interface{}, formData ReqGetTagList) ([]TagModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_DEF_TAG, r.gitConfig.URL, projectID)
	resp := []TagModel{}

	respOrigin, err := r.gitConfig.Client.R().SetAuthToken(r.gitConfig.Token).SetFormData(formData).Get(path)
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

func (r *RestTag) CreateTag(projectID interface{}, formData ReqCreateTag) (TagModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_DEF_TAG, r.gitConfig.URL, projectID)
	resp := TagModel{}

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
		errr := errors.New(fmt.Sprintf("http status: %s, message: %s", respOrigin.Status(), *resp.Message))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

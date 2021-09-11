package tag

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

const (
	PATH_DEF_TAG = "%s/api/v4/projects/%d/repository/tags"
)

type RestTag struct {
	url    string
	token  string
	client *resty.Client
}

func NewRestTag(url, token string, client *resty.Client) RestTag {
	return RestTag{url, token, client}
}

func (r *RestTag) GetAllTag(projectID int, formData ReqGetTagList) ([]Tag, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_DEF_TAG, r.url, projectID)
	resp := []Tag{}

	respOrigin, err := r.client.R().SetAuthToken(r.token).SetFormData(formData).Get(path)
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

func (r *RestTag) CreateTag(projectID int, formData ReqCreateTag) (Tag, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_DEF_TAG, r.url, projectID)
	resp := Tag{}

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
		errr := errors.New(fmt.Sprintf("http status: %s, message: %s", respOrigin.Status(), *resp.Message))
		logrus.Trace(errr)
		return resp, errr
	}

	return resp, nil
}

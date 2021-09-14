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
	PATH_LIST_COMMIT = "%s/api/v4/projects/%d/repository/commits"
)

type RestCommit struct {
	gitConfig *model.GitConfig
}

func NewRestCommit(gitConfig *model.GitConfig) RestCommit {
	return RestCommit{gitConfig}
}

func (r *RestCommit) GetCommit(projectID interface{}, formData ReqGetCommitList) ([]CommitModel, error) {
	formData.AddProjectID(projectID)
	path := fmt.Sprintf(PATH_LIST_COMMIT, r.gitConfig.URL, projectID)
	resp := []CommitModel{}

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

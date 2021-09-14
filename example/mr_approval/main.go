package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-resty/resty/v2"
	approval "github.com/ivandhitya/gitool/approval"
	"github.com/ivandhitya/gitool/example"
	"github.com/ivandhitya/gitool/model"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := example.Conf{}
	conf.GetConf("../../.gitool.yaml")

	// logrus formater example
	logrus.SetReportCaller(true)
	formatter := &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)

	client := resty.New().SetAllowGetMethodPayload(true)
	gitConfig := &model.GitConfig{
		Client: client,
		URL:    conf.Gitlab.Address,
		Token:  conf.Gitlab.Token,
	}

	approvalClient := approval.NewRestMRApprovalRule(gitConfig)
	projectID := 683
	mrID := 25
	// Get merge approval rules
	resp, err := approvalClient.GetMRApprovalRules(projectID, mrID)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(resp)

	req := make(approval.ReqUpdateMRApprovalRules)
	req.AddApprovalsRequired(0)
	for _, val := range resp {
		resp, err := approvalClient.UpdateMRApprovalRules(projectID, mrID, val.Id, req)
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Debug(resp)
		}

	}
}

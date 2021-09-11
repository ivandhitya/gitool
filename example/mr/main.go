package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/example"
	"github.com/ivandhitya/gitool/mr"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := example.Conf{}
	conf.GetConf()

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

	client := resty.New()
	mrClient := mr.NewRestMergeRequest(conf.Gitlab.Address, conf.Gitlab.Token, client)
	projectID := 17619669

	req := make(mr.ReqMR)

	// Merge Request
	req.AddSourceBranch("DC-223_test3").
		AddTargetBranch("development").
		AddTitle("testing merge by api test3").
		AddDescription("just testing 3")

	resp, err := mrClient.CreateMR(projectID, req)
	if err != nil {
		logrus.Error(err)
		resp, err := mrClient.DeleteMR(projectID, resp.IID)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Info(resp)
		return
	}
	logrus.Info(resp)

	// Accept MR
	reqAccept := make(mr.ReqAcceptMR)
	reqAccept.AddShouldRemoveSourceBranch(true)

	respAccept, err := mrClient.AcceptMR(projectID, resp.IID, reqAccept)
	if err != nil {
		logrus.Error(err)
		resp, err := mrClient.DeleteMR(projectID, resp.IID)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Info(resp)
		return
	}
	logrus.Info(respAccept)

}

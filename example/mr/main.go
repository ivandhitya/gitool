package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/mr"
	"github.com/sirupsen/logrus"
)

const (
	TOKEN   = ""                   // add your token here
	ADDRESS = "https://gitlab.com" // add your gitlab repository here
)

func main() {
	// logrus formater example
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		DisableColors: false,
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)

	client := resty.New()
	mrClient := mr.NewRestMergeRequest(ADDRESS, TOKEN, client)
	projectID := 17619669

	req := make(mr.ReqMR)

	// Merge Request
	req.AddSourceBranch("DC-222_test1").
		AddTargetBranch("development").
		AddTitle("testing merge by api test2").
		AddDescription("just testing 2")

	resp, err := mrClient.CreateMR(projectID, req)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(resp)

	// Accept MR
	reqAccept := make(mr.ReqAcceptMR)
	reqAccept.AddMergeRequestIID(resp.IID).
		AddShouldRemoveSourceBranch(true)

	respAccept, err := mrClient.AcceptMR(projectID, resp.IID, reqAccept)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(respAccept)

}

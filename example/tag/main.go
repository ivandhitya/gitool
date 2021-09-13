package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/example"
	"github.com/ivandhitya/gitool/model"
	"github.com/ivandhitya/gitool/tag"
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
	client := resty.New()
	gitConfig := &model.GitConfig{
		Client: client,
		URL:    conf.Gitlab.Address,
		Token:  conf.Gitlab.Token,
	}
	tagClient := tag.NewRestTag(gitConfig)
	projectID := 17619669

	// Get All Tag
	req := make(tag.ReqGetTagList)
	resp, err := tagClient.GetAllTag(projectID, req)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(resp)

	// Create new tag
	reqCreateTag := make(tag.ReqCreateTag)
	reqCreateTag.AddRef("master")
	reqCreateTag.AddMessage("test tag from api")
	reqCreateTag.AddTagName("v1.1.1")
	respCreate, err := tagClient.CreateTag(projectID, reqCreateTag)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(respCreate)

}

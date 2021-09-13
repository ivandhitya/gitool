package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/example"
	"github.com/ivandhitya/gitool/model"
	"github.com/ivandhitya/gitool/release"
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
	gitConfig := &model.GitConfig{
		Client: client,
		URL:    conf.Gitlab.Address,
		Token:  conf.Gitlab.Token,
	}

	releaseClient := release.NewRestRelease(gitConfig)
	projectID := 17619669

	// Create release
	req := make(release.ReqCreateRelease)
	tagName := "v1.1.1"
	releaseNote := `- [DC-518: [Tech Initiate] Drop Groot database connection in bernoulli](https://tcashsquad.atlassian.net/browse/DC-518)  
test new line with API`
	req.AddDescription(releaseNote).
		AddTagName(tagName).
		AddName(tagName)
	resp, err := releaseClient.CreateRelease(projectID, req)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(resp)
}

package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/tag"
	"github.com/sirupsen/logrus"
)

const (
	TOKEN   = "vYRx6yhxtdtTTseF_KfW" // add your token here
	ADDRESS = "https://gitlab.com"   // add your gitlab repository here
)

func main() {
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
	tagClient := tag.NewRestTag(ADDRESS, TOKEN, client)
	projectID := 17619669

	req := make(tag.ReqGetTagList)

	resp, err := tagClient.GetAllTag(projectID, req)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(resp[0])
}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/mr"
)

const (
	TOKEN   = ""                           // add your token here
	ADDRESS = "https://gitlab.linkaja.com" // add your gitlab repository here
)

func main() {
	client := resty.New()
	mrClient := mr.NewRestMergeRequest(ADDRESS, TOKEN, client)
	req := make(mr.ReqMR)
	projectID := 683
	req.AddSourceBranch("test-api-to-merge").
		AddTargetBranch("test-api").
		AddTitle("testing merge by api").
		AddDescription("just testing")

	resp, err := mrClient.CreateMR(projectID, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

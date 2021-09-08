package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/mr"
)

const (
	TOKEN   = ""                   // add your token here
	ADDRESS = "https://gitlab.com" // add your gitlab repository here
)

func main() {
	client := resty.New()
	mrClient := mr.NewRestMergeRequest(ADDRESS, TOKEN, client)
	projectID := 17619669

	req := make(mr.ReqMR)

	// Merge Request
	req.AddSourceBranch("DC-222_test1").
		AddTargetBranch("development").
		AddTitle("testing merge by api test1").
		AddDescription("just testing 1")

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

	// Accept MR
	reqAcceept := make(mr.ReqAcceptMR)
	reqAcceept.AddMergeRequestIID(resp.IID)

	respAccept, err := mrClient.AcceptMR(projectID, resp.IID, reqAcceept)
	if err != nil {
		fmt.Println(err)
		return
	}

	b1, err := json.Marshal(respAccept)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b1))
}

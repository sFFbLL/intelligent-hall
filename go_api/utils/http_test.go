package utils

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	var LoginData = map[string]string{"xxx1": "GetToken",
		"xxx2": "default",
		"xxx3": "root",
		"xxx4": "e10adc3949ba59abbe56e057f20f883e"}
	//fmt.Println(HttpPost(config.ApplicationConfig.TerminalHttpLogin, settings.LoginData, "application/javascript; charset=UTF-8"))
	fmt.Println(HttpPost("http://49.234.29.171:8060/mock/11/api/backend_mgt/v2/logon", LoginData, "application/json"))
}

func TestList(t *testing.T) {
	type pager struct {
		Total   int    `json:"total"`
		PerPage int    `json:"per_page"`
		Page    int    `json:"page"`
		OrderBy string `json:"orderby"`
		SortBy  string `json:"sortby"`
		KeyWord string `json:"keyword"`
		Status  string `json:"status"`
	}

	type ListData struct {
		ProjectName string `json:"project_name"`
		Action      string `json:"action"`
		CategoryID  string `json:"categoryID"`
		Pager       pager  `json:"Pager"`
		User        string `json:"user"`
		Token       string `json:"token"`
	}

	var TerminalList = ListData{
		ProjectName: "default",
		Action:      "getTermList",
		Pager: pager{Total: -1, PerPage: 10, Page: 1, OrderBy: "",
			SortBy: "", KeyWord: "", Status: "running"},
		User:  "root",
		Token: "1111"}
	//fmt.Println(HttpPost(config.ApplicationConfig.TerminalHttpList, settings.TerminalList, "application/javascript; charset=UTF-8"))
	fmt.Println(HttpPost("http://49.234.29.171:8060/mock/11/api/backend_mgt/v2/termcategory", TerminalList, "application/json"))
}

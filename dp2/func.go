package dp2

import (
	"encoding/json"
	"log"

	"github.com/fatih/structs"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
)

type SystemConfig struct {
	ServerUrl   string
	Dp2Username string
	Dp2Password string
	Session     *string
}

type Method interface {
	Login() LoginResult
	Logout() LogoutResult
	GetReaderInfo(barcode string, retType string) GetReaderInfoResult
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (s SystemConfig) Login() LoginResult {
	// Data
	raw := LoginRequest{StrUserName: s.Dp2Username, StrPassword: s.Dp2Password, StrParameters: "type=worker,client=REST|0.01"}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	// Post
	ret, err := requests.Post(s.ServerUrl+"/login", req)
	checkErr(err)
	// Save this session
	*s.Session = ret.Cookies[0].String()
	// Return
	var res LoginResult
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) Logout() LogoutResult {
	// Data
	req := url.NewRequest()
	req.Cookies = url.ParseCookies(s.ServerUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServerUrl+"/logout", req)
	checkErr(err)
	// Return
	var res LogoutResult
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetReaderInfo(barcode string, retType string) GetReaderInfoResult {
	// Data
	raw := GetReaderInfoRequest{StrBarcode: barcode, StrResultTypeList: retType}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServerUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServerUrl+"/getReaderInfo", req)
	checkErr(err)
	// Return
	var res GetReaderInfoResult
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

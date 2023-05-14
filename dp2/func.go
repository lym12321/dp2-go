package dp2

import (
	"encoding/json"
	"log"

	"github.com/fatih/structs"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
)

type SystemConfig struct {
	ServiceUrl  string
	Dp2Username string
	Dp2Password string
	Session     *string
}

type Method interface {
	Login() LoginReturn
	Logout() LogoutReturn
	GetReaderInfo(barcode string, retType string) GetReaderInfoReturn
	GetBiblioSummary(barcode string) GetBiblioSummaryReturn
	GetBiblioInfo(recPath string, retType string) GetBiblioInfoReturn
	Borrow(reader string, item string, renew bool) BorrowReturn
	Return(item string) ReturnReturn
	SearchReader(resultSetName string, keyword string) SearchReaderReturn
	SearchBiblio(resultSetName string, keyword string) SearchBiblioReturn
	GetSearchResult(resultSetName string, lStart int, lCount int) GetSearchResultReturn
	GetEntities(recPath string) GetEntitiesReturn
	GetItemInfo(barcode string, retType string) GetItemInfoReturn
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (s SystemConfig) Login() LoginReturn {
	// Data
	raw := LoginRequest{StrUserName: s.Dp2Username, StrPassword: s.Dp2Password, StrParameters: "type=worker,client=REST|0.01"}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/login", req)
	checkErr(err)
	// Save this session
	*s.Session = ret.Cookies[0].String()
	// Return
	var res LoginReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) Logout() LogoutReturn {
	// Data
	req := url.NewRequest()
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/logout", req)
	checkErr(err)
	// Return
	var res LogoutReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetReaderInfo(barcode string, retType string) GetReaderInfoReturn {
	// Data
	raw := GetReaderInfoRequest{StrBarcode: barcode, StrResultTypeList: retType}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/getReaderInfo", req)
	checkErr(err)
	// Return
	var res GetReaderInfoReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetBiblioSummary(barcode string) GetBiblioSummaryReturn {
	// Data
	raw := GetBiblioSummaryRequest{StrItemBarcode: barcode}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/getBiblioSummary", req)
	checkErr(err)
	// Return
	var res GetBiblioSummaryReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetBiblioInfo(recPath string, retType string) GetBiblioInfoReturn {
	// Data
	raw := GetBiblioInfoRequest{StrBiblioRecPath: recPath, StrBiblioType: retType}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/getBiblioInfo", req)
	checkErr(err)
	// Return
	var res GetBiblioInfoReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) Borrow(reader string, item string, renew bool) BorrowReturn {
	// renew 为 true 时，执行续借操作
	// Data
	raw := BorrowRequest{StrReaderBarcode: reader, StrItemBarcode: item, BRenew: renew, BForce: false}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/Borrow", req)
	checkErr(err)
	// Return
	var res BorrowReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) Return(item string) ReturnReturn {
	// Data
	raw := ReturnRequest{StrItemBarcode: item, StrAction: "return"}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/Return", req)
	checkErr(err)
	// Return
	var res ReturnReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) SearchReader(resultSetName string, keyword string) SearchReaderReturn {
	// 这里写死了查找方式 middle，以简化操作方式
	// Data
	raw := SearchReaderRequest{
		StrQueryWord:     keyword,
		StrReaderDbNames: "<all>",
		NPerMax:          -1,
		StrFrom:          "<all>",
		StrMatchStyle:    "middle",
		StrResultSetName: resultSetName,
	}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/searchReader", req)
	checkErr(err)
	// Return
	var res SearchReaderReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) SearchBiblio(resultSetName string, keyword string) SearchBiblioReturn {
	// 这里写死了查找方式 middle，以简化操作方式
	// Data
	raw := SearchBiblioRequest{
		StrQueryWord:     keyword,
		StrReaderDbNames: "<all>",
		NPerMax:          -1,
		StrFrom:          "<all>",
		StrMatchStyle:    "middle",
		StrResultSetName: resultSetName,
	}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/searchBiblio", req)
	checkErr(err)
	// Return
	var res SearchBiblioReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetSearchResult(resultSetName string, lStart int, lCount int) GetSearchResultReturn {
	raw := GetSearchResultRequest{StrResultSetName: resultSetName, LStart: lStart, LCount: lCount, StrBrowseInfoStyle: "id,cols"}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/getSearchResult", req)
	checkErr(err)
	// Return
	var res GetSearchResultReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetEntities(recPath string) GetEntitiesReturn {
	// 考虑到同一路径下的书不会太多，这里默认返回全部信息
	raw := GetEntitiesRequest{StrBiblioRecPath: recPath, LStart: 0, LCount: -1, StrStyle: "onlygetpath"}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/getEntities", req)
	checkErr(err)
	// Return
	var res GetEntitiesReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

func (s SystemConfig) GetItemInfo(barcode string, retType string) GetItemInfoReturn {
	raw := GetItemInfoRequest{StrBarcode: barcode, StrResultType: retType}
	req := url.NewRequest()
	req.Json = structs.Map(&raw)
	req.Cookies = url.ParseCookies(s.ServiceUrl, *s.Session)
	// Post
	ret, err := requests.Post(s.ServiceUrl+"/getItemInfo", req)
	checkErr(err)
	// Return
	var res GetItemInfoReturn
	err = json.Unmarshal(ret.Content, &res)
	checkErr(err)
	return res
}

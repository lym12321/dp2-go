package dp2

// Data Structure

type LoginRequest struct {
	StrUserName   string `structs:"strUserName"`
	StrPassword   string `structs:"strPassword"`
	StrParameters string `structs:"strParameters"` //'type=worker,client=REST|0.01'
}

type LoginResult struct {
	LoginResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"LoginResult"`
	StrOutputUserName string `json:"strOutputUserName"`
	StrRights         string `json:"strRights"`
	StrLibraryCode    string `json:"strLibraryCode"`
}

type LogoutResult struct {
	LogoutResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"LogoutResult"`
}

type GetReaderInfoRequest struct {
	StrBarcode        string `structs:"strBarcode"`
	StrResultTypeList string `structs:"strResultTypeList"`
}

type GetReaderInfoResult struct {
	GetReaderInfoResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"GetReaderInfoResult"`
	Results     []string `json:"results"`
	StrRecPath  string   `json:"strRecPath"`
	BaTimestamp []int    `json:"baTimestamp"`
}

type GetBiblioSummaryRequest struct {
	StrItemBarcode string `structs:"strItemBarcode"`
}

type GetBiblioInfoRequest struct {
	StrBiblioRecPath string `structs:"strBiblioRecPath"`
	StrBiblioType    string `structs:"strBiblioType"`
}

type BorrowRequest struct {
	BRenew           bool   `structs:"bRenew"`
	StrReaderBarcode string `structs:"strReaderBarcode"`
	StrItemBarcode   string `structs:"strItemBarcode"`
	BForce           bool   `structs:"bForce"` // false
}

type ReturnRequest struct {
	StrAction      string `structs:"strAction"` // "return"
	StrItemBarcode string `structs:"strItemBarcode"`
}

type SearchReaderRequest struct {
	StrReaderDbNames string `structs:"strReaderDbNames"` // "<all>"
	StrQueryWord     string `structs:"strQueryWord"`
	NPerMax          int    `structs:"nPerMax"` // -1
	StrFrom          string `structs:"strFrom"`
	StrMatchStyle    string `structs:"strMatchStyle"`    // middle
	StrResultSetName string `structs:"strResultSetName"` // default
}

type SearchBiblioRequest struct {
	StrReaderDbNames string `structs:"strReaderDbNames"` // "<all>"
	StrQueryWord     string `structs:"strQueryWord"`
	NPerMax          int    `structs:"nPerMax"` // -1
	StrFrom          string `structs:"strFrom"`
	StrMatchStyle    string `structs:"strMatchStyle"`    // middle
	StrResultSetName string `structs:"strResultSetName"` // default
}

type GetSearchResultRequest struct {
	StrResultSetName   string `structs:"strResultSetName"`
	LStart             int    `structs:"lStart"`
	LCount             int    `structs:"lCount"`
	StrBorrowInfoStyle string `structs:"strBorrowInfoStyle"` // "id,cols"
}

type GetEntitiesRequest struct {
	StrBiblioRecPath string `structs:"strBiblioRecPath"`
	LStart           int    `structs:"lStart"`
	LCount           int    `structs:"lCount"`
	StrStyle         string `structs:"strStyle"` // "onlygetpath"
}

type GetItemInfoRequest struct {
	StrBarcode    string `structs:"strBarcode"`
	StrResultType string `structs:"strResultType"`
}

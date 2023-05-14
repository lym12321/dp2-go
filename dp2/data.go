package dp2

// Data Structure

type LoginRequest struct {
	StrUserName   string `structs:"strUserName"`
	StrPassword   string `structs:"strPassword"`
	StrParameters string `structs:"strParameters"` //'type=worker,client=REST|0.01'
}

type LoginReturn struct {
	LoginResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"LoginResult"`
	StrOutputUserName string `json:"strOutputUserName"`
	StrRights         string `json:"strRights"`
	StrLibraryCode    string `json:"strLibraryCode"`
}

type LogoutReturn struct {
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

type GetReaderInfoReturn struct {
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

type GetBiblioSummaryReturn struct {
	GetBiblioSummaryResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"GetBiblioSummaryResult"`
	StrBiblioRecPath string `json:"strBiblioRecPath"`
	StrSummary       string `json:"strSummary"`
}

type GetBiblioInfoRequest struct {
	StrBiblioRecPath string `structs:"strBiblioRecPath"`
	StrBiblioType    string `structs:"strBiblioType"`
}

type GetBiblioInfoReturn struct {
	GetBiblioInfoResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"GetBiblioInfoResult"`
	StrBiblio string `json:"strBiblio"`
}

type BorrowRequest struct {
	BRenew           bool   `structs:"bRenew"`
	StrReaderBarcode string `structs:"strReaderBarcode"`
	StrItemBarcode   string `structs:"strItemBarcode"`
	BForce           bool   `structs:"bForce"` // false
}

type BorrowReturn struct {
	BorrowResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"BorrowResult"`
	ItemRecords   interface{} `json:"item_records"`
	ReaderRecords interface{} `json:"reader_records"`
	BiblioRecords interface{} `json:"biblio_records"`
	BorrowInfo    struct {
		BorrowCount      int         `json:"BorrowCount"`
		BorrowID         string      `json:"BorrowID"`
		BorrowOperator   string      `json:"BorrowOperator"`
		DenyPeriod       string      `json:"DenyPeriod"`
		ItemBarcode      string      `json:"ItemBarcode"`
		LatestReturnTime string      `json:"LatestReturnTime"`
		Overflows        interface{} `json:"Overflows"`
		Period           string      `json:"Period"`
	} `json:"borrow_info"`
	ADupPath               interface{} `json:"aDupPath"`
	StrOutputReaderBarcode string      `json:"strOutputReaderBarcode"`
}

type ReturnRequest struct {
	StrAction      string `structs:"strAction"` // "return"
	StrItemBarcode string `structs:"strItemBarcode"`
}

type ReturnReturn struct {
	ReturnResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"ReturnResult"`
	ItemRecords            interface{} `json:"item_records"`
	ReaderRecords          interface{} `json:"reader_records"`
	BiblioRecords          interface{} `json:"biblio_records"`
	ADupPath               interface{} `json:"aDupPath"`
	StrOutputReaderBarcode string      `json:"strOutputReaderBarcode"`
	ReturnInfo             struct {
		BookType         string      `json:"BookType"`
		BorrowCount      int         `json:"BorrowCount"`
		BorrowID         string      `json:"BorrowID"`
		BorrowOperator   string      `json:"BorrowOperator"`
		BorrowTime       string      `json:"BorrowTime"`
		Borrower         string      `json:"Borrower"`
		ItemBarcode      string      `json:"ItemBarcode"`
		LatestReturnTime string      `json:"LatestReturnTime"`
		Location         string      `json:"Location"`
		OverdueString    string      `json:"OverdueString"`
		Period           string      `json:"Period"`
		ReturnOperator   string      `json:"ReturnOperator"`
		Volume           interface{} `json:"Volume"`
	} `json:"return_info"`
}

type SearchReaderRequest struct {
	StrReaderDbNames string `structs:"strReaderDbNames"` // "<all>"
	StrQueryWord     string `structs:"strQueryWord"`
	NPerMax          int    `structs:"nPerMax"` // -1
	StrFrom          string `structs:"strFrom"`
	StrMatchStyle    string `structs:"strMatchStyle"`    // middle
	StrResultSetName string `structs:"strResultSetName"` // default
}

type SearchReaderReturn struct {
	SearchReaderResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"SearchReaderResult"`
}

type SearchBiblioRequest struct {
	StrReaderDbNames string `structs:"strReaderDbNames"` // "<all>"
	StrQueryWord     string `structs:"strQueryWord"`
	NPerMax          int    `structs:"nPerMax"` // -1
	StrFrom          string `structs:"strFrom"`
	StrMatchStyle    string `structs:"strMatchStyle"`    // middle
	StrResultSetName string `structs:"strResultSetName"` // default
}

type SearchBiblioReturn struct {
	SearchBiblioResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"SearchBiblioResult"`
	StrQueryXML string `json:"strQueryXml"`
}

type GetSearchResultRequest struct {
	StrResultSetName   string `structs:"strResultSetName"`
	LStart             int    `structs:"lStart"`
	LCount             int    `structs:"lCount"`
	StrBrowseInfoStyle string `structs:"strBrowseInfoStyle"` // "id,cols"
}

type GetSearchResultReturn struct {
	GetSearchResultResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"GetSearchResultResult"`
	Searchresults []struct {
		Cols       []string    `json:"Cols"`
		Keys       interface{} `json:"Keys"`
		Path       string      `json:"Path"`
		RecordBody struct {
			Metadata  string      `json:"Metadata"`
			Path      string      `json:"Path"`
			Result    interface{} `json:"Result"`
			Timestamp interface{} `json:"Timestamp"`
			XML       interface{} `json:"Xml"`
		} `json:"RecordBody"`
	} `json:"searchresults"`
}

type GetEntitiesRequest struct {
	StrBiblioRecPath string `structs:"strBiblioRecPath"`
	LStart           int    `structs:"lStart"`
	LCount           int    `structs:"lCount"`
	StrStyle         string `structs:"strStyle"` // "onlygetpath"
}

type GetEntitiesReturn struct {
	GetEntitiesResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"GetEntitiesResult"`
	Entityinfos []struct {
		Action       string      `json:"Action"`
		ErrorCode    int         `json:"ErrorCode"`
		ErrorInfo    string      `json:"ErrorInfo"`
		NewRecPath   string      `json:"NewRecPath"`
		NewRecord    string      `json:"NewRecord"`
		NewTimestamp interface{} `json:"NewTimestamp"`
		OldRecPath   string      `json:"OldRecPath"`
		OldRecord    string      `json:"OldRecord"`
		OldTimestamp interface{} `json:"OldTimestamp"`
		RefID        string      `json:"RefID"`
		Style        string      `json:"Style"`
	} `json:"entityinfos"`
}

type GetItemInfoRequest struct {
	StrBarcode    string `structs:"strBarcode"`
	StrResultType string `structs:"strResultType"`
}

type GetItemInfoReturn struct {
	GetItemInfoResult struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
		Value     int    `json:"Value"`
	} `json:"GetItemInfoResult"`
	StrResult        string `json:"strResult"`
	StrItemRecPath   string `json:"strItemRecPath"`
	ItemTimestamp    []int  `json:"item_timestamp"`
	StrBiblio        string `json:"strBiblio"`
	StrBiblioRecPath string `json:"strBiblioRecPath"`
}

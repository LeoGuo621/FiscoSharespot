package utils

type Resp struct {
	Recode string `json:"recode"`
	Msg string `json:"msg"`
	// valid data
	Data interface{} `json:"data"`
}

// scheduled some predictable err recodes
const (
	RecodeOk        = "0"
	RecodeDBErr     = "4001"
	RecodeLoginErr  = "4002"
	RecodeParamErr  = "4003"
	RecodeSysErr     = "4004"
	RecodeFiscoErr   = "4105"
	RecodeContractErr   = "4106"
	RecodeUnknownErr = "4107"
)

var recodeText = map[string]string{
	RecodeOk:         "Success",
	RecodeDBErr:      "Database operation error",
	RecodeLoginErr:   "User login error",
	RecodeParamErr:   "Parameter error",
	RecodeSysErr:     "System error",
	RecodeFiscoErr:   "Error interacting with Fisco",
	RecodeContractErr:   "Error interacting with contract",
	RecodeUnknownErr: "Unknown error",
}

func RecodeText(code string) string {
	if str, ok := recodeText[code]; ok {
		return str
	}
	return recodeText[RecodeUnknownErr]
}

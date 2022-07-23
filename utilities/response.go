package utilities

type ReturnJSON struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func ResponseWithModel(code string, model interface{}, message, errType string) ReturnJSON {
	return ReturnJSON{
		Code:  code,
		Msg:   message,
		Model: model,
	}
}

func ResponseWithError(code string, message string) ReturnJSON {

	return ReturnJSON{
		Code:  code,
		Msg:   message,
		Model: make(map[string]string),
	}
}

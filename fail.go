package jpush

import "fmt"

type Fail struct {
	Error struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (f *Fail) String() string {
	return fmt.Sprintf("error_code=%d,error_msg=%s", f.Error.Code, f.Error.Message)
}

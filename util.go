package jpush

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Auth(appKey, masterSecret string) string {
	str := appKey + ":" + masterSecret
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(str)))
}

func Request(method, url string, body io.Reader, auth string) (result []byte, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = rsp.Body.Close()
	}()

	result, err = ioutil.ReadAll(rsp.Body)
	return
}

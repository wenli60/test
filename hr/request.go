package hr

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ToJson(v interface{}) string {
	str, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return str
}

// RequestDw 请求hrdw服务接口
func RequestDw(reqUrl string) ([]byte, error) {
	addArgs := AddData{
		QueryCondition: QueryCondition{ArgMap: ArgMap{}},
	}

	argStr := ToJson(addArgs)
	request, _ := http.NewRequest("POST", reqUrl, strings.NewReader(argStr))
	headers := hrDwHeaders()
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	var requestRsp *http.Response
	requestRsp, err := http.DefaultClient.Do(request)
	defer requestRsp.Body.Close()
	var rspBytes []byte
	rspBytes, err = ioutil.ReadAll(requestRsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBytes, nil
}

func hrDwHeaders() map[string]string {
	token := "9a9d5136-bf06-4662-8a91-2a311967d2a7"
	appName := "IEG-CROS"
	timestamp := fmt.Sprintf("%v", time.Now().Unix())
	signature := fmt.Sprintf("%x", sha256.Sum256([]byte(appName+token+timestamp)))
	return map[string]string{
		"Content-Type":   "application/json",
		"hrgw-appname":   appName,
		"hrgw-timestamp": timestamp,
		"hrgw-signature": signature,
		"requestId":      uuid.New().String(),
	}
}

type AddData struct {
	QueryCondition QueryCondition `json:"queryCondition"`
}

type QueryCondition struct {
	ArgMap ArgMap `json:"argMap"`
}

type ArgMap struct {
}

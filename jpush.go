package jpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ApiPush      = "https://api.jpush.cn/v3/push"
	ApiGroupPush = "https://api.jpush.cn/v3/grouppush"
	ApiCid       = "https://api.jpush.cn/v3/push/cid"
	ApiValidate  = "https://api.jpush.cn/v3/push/validate"
	ApiFile      = "https://api.jpush.cn/v3/push/file"
	ApiCancel    = "https://api.jpush.cn/v3/push/"

	CidTypePush     = "push"
	CidTypeSchedule = "schedule"
)

type Jpush struct {
	AppKey       string
	MasterSecret string
	GroupKey     string
	Auth         string
	GroupAuth    string
}

func NewJpush(appKey, masterSecret string, groupKey ...string) *Jpush {
	jp := &Jpush{
		AppKey:       appKey,
		MasterSecret: masterSecret,
	}
	jp.Auth = Auth(appKey, masterSecret)
	if len(groupKey) > 0 {
		jp.GroupKey = GroupAuth(groupKey[0], masterSecret)
	}
	return jp
}

// 普通推送
func (j *Jpush) Push(request *PushRequest) (result PushResult, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	data, err := Request(http.MethodPost, ApiPush, bytes.NewBuffer(body), j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 应用分组推送
func (j *Jpush) GroupPush(request *PushRequest) (result PushResult, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	data, err := Request(http.MethodPost, ApiGroupPush, bytes.NewBuffer(body), j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 获取推送唯一标识符
func (j *Jpush) Cid(count int64, types string) (result CidResult, err error) {
	url := fmt.Sprintf("%s?count=%d&type=%s", ApiCid, count, types)
	data, err := Request(http.MethodGet, url, nil, j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 推送校验
func (j *Jpush) Validate(request *PushRequest) (result PushResult, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	data, err := Request(http.MethodPost, ApiValidate, bytes.NewBuffer(body), j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 文件推送
func (j *Jpush) File(request *PushRequest) (result PushResult, err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	data, err := Request(http.MethodPost, ApiFile, bytes.NewBuffer(body), j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 推送撤销
func (j *Jpush) Cancel(msgId string) (result PushResult, err error) {
	url := ApiCancel + msgId
	data, err := Request(http.MethodDelete, url, nil, j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

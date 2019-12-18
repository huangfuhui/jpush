package jpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	ApiPush      = "https://api.jpush.cn/v3/push"
	ApiGroupPush = "https://api.jpush.cn/v3/grouppush"
	ApiCid       = "https://api.jpush.cn/v3/push/cid"
	ApiValidate  = "https://api.jpush.cn/v3/push/validate"
	ApiFile      = "https://api.jpush.cn/v3/push/file"
	ApiCancel    = "https://api.jpush.cn/v3/push/"
	ApiDevice    = "https://device.jpush.cn/v3/devices/"
	ApiAlias     = "https://device.jpush.cn/v3/aliases/"
	ApiTags      = "https://device.jpush.cn/v3/tags/"

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

// 查询设备的别名与标签
func (j *Jpush) GetTagAndAlias(registrationId string) (result DeviceResult, err error) {
	url := ApiDevice + registrationId
	data, err := Request(http.MethodGet, url, nil, j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 设置设备的别名与标签
func (j *Jpush) AddTagAndAlias(registrationId string, request *TagAndAliasRequest) (err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	url := ApiDevice + registrationId
	_, err = Request(http.MethodPost, url, bytes.NewBuffer(body), j.Auth)

	return
}

// 查询别名
func (j *Jpush) GetAlias(alias string, platform ...string) (result AliasResult, err error) {
	url := ApiAlias + alias
	if len(platform) > 0 {
		url += "?platform=" + strings.Join(platform, ",")
	}
	data, err := Request(http.MethodGet, url, nil, j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 删除别名
func (j *Jpush) DeleteAlias(alias string, platform ...string) (err error) {
	url := ApiAlias + alias
	if len(platform) > 0 {
		url += "?platform=" + strings.Join(platform, ",")
	}
	_, err = Request(http.MethodDelete, url, nil, j.Auth)

	return
}

// 查询标签列表
func (j *Jpush) GetTags() (result TagsResult, err error) {
	data, err := Request(http.MethodGet, ApiTags, nil, j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 判断设备与标签绑定关系
func (j *Jpush) DeviceAndTag(tag, registrationId string) (result DeviceAndTagResult, err error) {
	url := fmt.Sprintf("%s%s/registration_ids/%s", ApiTags, tag, registrationId)
	data, err := Request(http.MethodGet, url, nil, j.Auth)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	return
}

// 更新标签
func (j *Jpush) UpdateTag(tag string, request *UpdateTagRequest) (err error) {
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	url := ApiTags + tag
	_, err = Request(http.MethodPost, url, bytes.NewBuffer(body), j.Auth)

	return
}

// 删除标签
func (j *Jpush) DeleteTag(tag string, platform ...string) (err error) {
	url := ApiTags + tag
	if len(platform) > 0 {
		url += "?platform=" + strings.Join(platform, ",")
	}
	_, err = Request(http.MethodDelete, url, nil, j.Auth)

	return
}

package jpush

const (
	PlatformAll      = "all"
	PlatformAndroid  = "android"
	PlatformIos      = "ios"
	PlatFormWinphone = "winphone"
)

type PushRequest struct {
	Platform        interface{}      `json:"platform"`
	Audience        interface{}      `json:"audience"`
	Notification    *Notification    `json:"notification,omitempty"`
	Message         *Message         `json:"message,omitempty"`
	Notification3rd *Notification3rd `json:"notification_3rd,omitempty"`
	SmsMessage      *SmsMessage      `json:"sms_message,omitempty"`
	Options         *Options         `json:"options,omitempty"`
	Callback        *Callback        `json:"callback,omitempty"`
	Cid             string           `json:"cid,omitempty"`
}

type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	ABTest         []string `json:"abtest,omitempty"`
}

type Notification struct {
	Alert    string                `json:"alert,omitempty"`
	Android  *NotificationAndroid  `json:"android,omitempty"`
	IOS      *NotificationIos      `json:"ios,omitempty"`
	WinPhone *NotificationWinphone `json:"winphone,omitempty"`
}

type NotificationAndroid struct {
	Alert         string      `json:"alert"`
	Title         string      `json:"title,omitempty"`
	BuilderId     int64       `json:"builder_id,omitempty"`
	ChannelId     string      `json:"channel_id,omitempty"`
	Priority      int64       `json:"priority,omitempty"`
	Category      string      `json:"category,omitempty"`
	Style         int64       `json:"style,omitempty"`
	AlertType     int64       `json:"alert_type,omitempty"`
	BigText       string      `json:"big_text,omitempty"`
	Inbox         interface{} `json:"inbox,omitempty"`
	BigPicPath    string      `json:"big_pic_path,omitempty"`
	Extras        interface{} `json:"extras,omitempty"`
	LargeIcon     string      `json:"large_icon,omitempty"`
	Intent        interface{} `json:"intent,omitempty"`
	UriActivity   string      `json:"uri_activity,omitempty"`
	UriAction     string      `json:"uri_action,omitempty"`
	BadgeAddNum   int64       `json:"badge_add_num,omitempty"`
	BadgeClass    string      `json:"badge_class,omitempty"`
	Sound         string      `json:"sound,omitempty"`
	ShowBeginTime string      `json:"show_begin_time,omitempty"`
	ShowEndTime   string      `json:"show_end_time,omitempty"`
}

type NotificationIos struct {
	Alert            interface{} `json:"alert"`
	Sound            string      `json:"sound,omitempty"`
	Badge            interface{} `json:"badge,omitempty"`
	ContentAvailable bool        `json:"content-available,omitempty"`
	MutableContent   bool        `json:"mutable-content,omitempty"`
	Category         string      `json:"category,omitempty"`
	Extras           interface{} `json:"extras,omitempty"`
	ThreadId         string      `json:"thread_id,omitempty"`
}

type NotificationWinphone struct {
	Alert    string      `json:"alert"`
	Title    string      `json:"title,omitempty"`
	OpenPage string      `json:"_open_page,omitempty"`
	Extras   interface{} `json:"extras,omitempty"`
}

type Message struct {
	MsgContent  string      `json:"msg_content"`
	Title       string      `json:"title,omitempty"`
	ContentType string      `json:"content_type,omitempty"`
	Extras      interface{} `json:"extras,omitempty"`
}

type Notification3rd struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content"`
}

type SmsMessage struct {
	DelayTime    int64       `json:"delay_time"`
	SignId       int64       `json:"signid,omitempty"`
	TempId       int64       `json:"temp_id"`
	TempPara     interface{} `json:"temp_para,omitempty"`
	ActiveFilter bool        `json:"active_filter,omitempty"`
}

type Options struct {
	SendNo            int64       `json:"sendno,omitempty"`
	TimeToLive        int64       `json:"time_to_live,omitempty"`
	OverrideMsgId     int64       `json:"override_msg_id,int64,omitempty"`
	ApnsProduction    bool        `json:"apns_production"`
	ApnsCollapseId    string      `json:"apns_collapse_id,omitempty"`
	BigPushDuration   int64       `json:"big_push_duration,omitempty"`
	ThirdPartyChannel interface{} `json:"third_party_channel,omitempty"`
}

type Callback struct {
	Url    string      `json:"url,omitempty"`
	Params interface{} `json:"params,omitempty"`
	Type   string      `json:"type,omitempty"`
}

type PushResult struct {
	SendNo string `json:"sendno"`
	MsgId  string `json:"msg_id"`
}

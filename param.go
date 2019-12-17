package jpush

type PushParam struct {
	Platform     string        `json:"platform"`
	Audience     interface{}   `json:"audience,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	SmsMessage   *SmsMessage   `json:"sms_message,omitempty"`
	Options      *PushOptions  `json:"options,omitempty"`
}

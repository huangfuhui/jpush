package jpush

const (
	ApiPush = "https://api.jpush.cn/v3/push"
)

type Jpush struct {
	AppKey       string
	MasterSecret string
}

func NewJpush(appKey, masterSecret string) *Jpush {
	return &Jpush{
		AppKey:       appKey,
		MasterSecret: masterSecret,
	}
}

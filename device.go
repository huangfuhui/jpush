package jpush

type DeviceResult struct {
	Tags   []string `json:"tags"`
	Alias  string   `json:"alias"`
	Mobile string   `json:"mobile"`
}

type TagAndAliasRequest struct {
	Tags struct {
		Add    []string `json:"add"`
		Remove []string `json:"remove"`
	} `json:"tags"`
	Alias  string `json:"alias"`
	Mobile string `json:"mobile"`
}

type AliasResult struct {
	RegistrationIds []string `json:"registration_ids"`
}

type TagsResult struct {
	Tags []string `json:"tags"`
}

type DeviceAndTagResult struct {
	Result bool `json:"result"`
}

type UpdateTagRequest struct {
	RegistrationIds struct {
		Add    []string `json:"add"`
		Remove []string `json:"remove"`
	} `json:"registration_ids"`
}

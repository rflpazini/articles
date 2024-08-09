package shortener

import (
	jsoniter "github.com/json-iterator/go"
)

type URLInfo struct {
	Url       string `json:"url"`
	Short     string `json:"short"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func (m URLInfo) UnmarshalBinary(data []byte) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, m)
}

func (m URLInfo) MarshalBinary() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(m)
}

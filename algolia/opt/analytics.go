package opt

import "encoding/json"

type AnalyticsOption struct {
	value bool
}

func Analytics(v bool) AnalyticsOption {
	return AnalyticsOption{v}
}

func (o AnalyticsOption) Get() bool {
	return o.value
}

func (o AnalyticsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AnalyticsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

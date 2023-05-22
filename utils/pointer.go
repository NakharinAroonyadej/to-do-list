package utils

func GetStringValue(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func NewString(s string) *string {
	return &s
}
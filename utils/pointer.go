package utils

func GetStringValue(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
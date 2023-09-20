package utils

func NewString(s string) *string {
	res := new(string)
	*res = s
	return res
}

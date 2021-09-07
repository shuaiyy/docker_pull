package code

import "encoding/base64"

func UrlEncode(in string) string {
	toString := base64.RawURLEncoding.EncodeToString([]byte(in))
	return toString
}

func UrlDecode(in string) string {
	res, err := base64.RawURLEncoding.DecodeString(in)
	if err != nil{
		return ""
	}
	return string(res)

}

package code

import "testing"

func TestUrlEncode(t *testing.T) {
	got := UrlEncode("k8s.gcr.io/kublete/adm:v1.18.6");
	t.Log(got)
	t.Log(UrlDecode(got))
}
package textdata

import (
	"strings"
	"testing"
)

// Search is a function search textdata in shared repository
func TestDecodeCSV(t *testing.T) {
	gotest := func(csv string, expect []Textdata) {
		r := strings.NewReader(csv)
		datas, err := DecodeCSV(r)
		if err != nil {
			msg := "error occured"
			t.Error(msg)
		}

		if len(expect) != len(datas) {
			t.Errorf("expect %d records, but returns %d.\r\n[expect]:%s\r\n[result]:%s",
				len(expect), len(datas), expect, datas)
			return
		}
	}

	genExpect := func(namespace string, caption string, body string) Textdata {
		return Textdata{Namespace: namespace, Caption: caption, Body: body, CreateAt: ""}
	}

	gotest(`f,a,b`, []Textdata{genExpect("f", "a", "b")})
	gotest(`f,"a",b`, []Textdata{genExpect("f", "a", "b")})
	gotest(`f,a,"b"`, []Textdata{genExpect("f", "a", "b")})
	gotest(`f,"""a""","b"`, []Textdata{genExpect("f", "\"a\"", "b")})
	gotest(`f,"""a\r\nb\rc\nd""","b"`, []Textdata{genExpect("f", "\"a\r\nb\rc\nd\"", "b")})
}

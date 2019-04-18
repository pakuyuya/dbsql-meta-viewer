package textdata

import (
	"bytes"
	"fmt"
	"testing"
)

// Search is a function search textdata in shared repository
func TestEncodeCSV(t *testing.T) {
	linebreak := fmt.Sprintln("")
	gotest := func(datas []Textdata, expect string) {
		w := bytes.Buffer{}

		err := EncodeCSV(&w, datas)
		if err != nil {
			msg := "error occured"
			t.Error(msg)
		}

		s := w.String()
		if s != expect {
			t.Errorf("[expect]:%s\r\n[result]:%s", expect, s)
			return
		}
	}

	genExpect := func(namespace string, caption string, body string) Textdata {
		return Textdata{Namespace: namespace, Caption: caption, Body: body, CreateAt: ""}
	}

	gotest([]Textdata{genExpect("f", "a", "b")}, `f,a,b`+linebreak)
	gotest([]Textdata{genExpect("f", "a", "b"), genExpect("f", "b", "c")}, `f,a,b`+linebreak+`f,b,c`+linebreak)
	gotest([]Textdata{genExpect("f", "\"a\"", "b")}, `f,"""a""",b`+linebreak)
	gotest([]Textdata{genExpect("f", "\"a\r\nb\rc\nd\"", "b")}, "f,\"\"\"a\r\nb\rc\nd\"\"\",b"+linebreak)
}

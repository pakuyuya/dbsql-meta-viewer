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

	genExpect := func(caption string, body string) Textdata {
		return Textdata{Caption: caption, Body: body, CreateAt: ""}
	}

	gotest([]Textdata{genExpect("a", "b")}, `a,b`+linebreak)
	gotest([]Textdata{genExpect("a", "b"), genExpect("b", "c")}, `a,b`+linebreak+`b,c`+linebreak)
	gotest([]Textdata{genExpect("\"a\"", "b")}, `"""a""",b`+linebreak)
	gotest([]Textdata{genExpect("\"a\r\nb\rc\nd\"", "b")}, "\"\"\"a\r\nb\rc\nd\"\"\",b"+linebreak)
}

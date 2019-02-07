package textdata

import (
	"testing"
)

// Search is a function search textdata in shared repository
func TestSearch(t *testing.T) {
	repos := []Textdata{
		Textdata{Caption: "caption", Body: "this is content", Since: "2019-01-01T11:22:33"},
		Textdata{Caption: "caption2", Body: "this is content2", Since: "2019-01-01T11:22:33"},
		Textdata{Caption: "caption3", Body: "this is content3", Since: "2019-01-01T11:22:33"},
	}
	SwitchRepository(&repos)

	res1, err := Search("caption")
	if err != nil {
		t.Fatalf("failed case1 err:%s", err.Error())
	}
	if len(res1) != 3 {
		t.Fatal("failed case1")
	}

	res2, err := Search("^caption$", WithRegex(true))
	if err != nil {
		t.Fatalf("failed case2 err:%s", err.Error())
	}
	if len(res2) != 1 || res2[0].Caption != "caption" {
		t.Fatalf("failed case2 %s", res2)
	}

	res3, err := Search("")
	if err != nil {
		t.Fatalf("failed case3 err:%s", err.Error())
	}
	if len(res3) != 3 {
		t.Fatalf("failed case3 returns: %s", res3)
	}

	res4, err := Search("", WithWordOnly(true))
	if err != nil {
		t.Fatalf("failed case4 err:%s", err.Error())
	}
	if len(res4) != 3 {
		t.Fatalf("failed case4 returns: %s", res3)
	}

	res5, err := Search("content", WithWordOnly(true))
	if err != nil {
		t.Fatalf("failed case4 err:%s", err.Error())
	}
	if len(res5) != 1 || res5[0].Caption != "caption" {
		t.Fatalf("failed case4 returns: %s", res3)
	}
}

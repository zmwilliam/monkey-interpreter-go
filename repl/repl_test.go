package repl_test

import (
	"bytes"
	"strings"
	"testing"

	"monkey/repl"
)

func TestStartEmptyInput(t *testing.T) {
	in := strings.NewReader("")
	var out bytes.Buffer

	repl.Start(in, &out)

	want := ">>"
	got := out.String()
	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

func TestStart(t *testing.T) {
	in := strings.NewReader("let ten = 10;")
	var out bytes.Buffer

	repl.Start(in, &out)

	want := `>>{Type:LET Literal:let}
{Type:IDENT Literal:ten}
{Type:= Literal:=}
{Type:INT Literal:10}
{Type:; Literal:;}
>>`
	got := out.String()
	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

package src

import (
	"testing"
)

func  TestNoArgs(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks",
			"I love music of Kartik.",
			"I love music of Kartik."}, 
		"I love music.\n \nI love music of Kartik.\nThanks\nI love music of Kartik.\n"},
		
	}
	for _, c := range cases {
		got:= Uniq(c.in,false,false,false,false,0,0);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want)
		}
	}
}

func  TestCount(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks",
			"I love music of Kartik.",
			"I love music of Kartik."}, 
			"3 I love music.\n1  \n2 I love music of Kartik.\n1 Thanks\n2 I love music of Kartik.\n"},
		
	}
	for _, c := range cases {
		got:= Uniq(c.in,true,false,false,false,0,0);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want)
		}
	}
}

func  TestDelete(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
					 "I love music.",
			         "I love music.",
					 " ",
					 "I love music of Kartik.",
					 "I love music of Kartik.",
					 "Thanks",
					 "I love music of Kartik.",
					 "I love music of Kartik."},

			"I love music.\nI love music of Kartik.\nI love music of Kartik.\n"},
	}
	for _, c := range cases {
		got:= Uniq(c.in,false,true,false,false,0,0);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want);
		}
	}
}

func  TestUnique(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
					 "I love music.",
			         "I love music.",
					 " ",
					 "I love music of Kartik.",
					 "I love music of Kartik.",
					 "Thanks",
					 "I love music of Kartik.",
					 "I love music of Kartik."},

			" \nThanks\n"},
	}
	for _, c := range cases {
		got:= Uniq(c.in,false,false,true,false,0,0);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want);
		}
	}
}

func  TestCaseInsesitive(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
			"I love music.",
			"I love MUSIC.",
			" ",
			"I love muSIc of Kartik.",
			"I love MusiC of KarTiK.",
			"Thanks",
			"I love musiC of Kartik.",
			"I love mUsiC of Kartik."}, 
		"I love MUSIC.\n \nI love MusiC of KarTiK.\nThanks\nI love mUsiC of Kartik.\n"},
	}
	for _, c := range cases {
		got:= Uniq(c.in,false,false,false,true,0,0);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want);
		}
	}
}


func  TestIgnoreWords(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"We love music.",
			"I love music.",
			"They love music.",
			" ",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks",
			}, 
		"They love music.\n \nWe love music of Kartik.\nThanks\n"},
	}
	for _, c := range cases {
		got:= Uniq(c.in,false,false,false,false,1,0);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want);
		}
	}
}

func  TestIgnoreSymbols(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
			"A love music.",
			"C love music.",
			" ",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks",
			}, 
		"C love music.\n \nI love music of Kartik.\nWe love music of Kartik.\nThanks\n"},
	}
	for _, c := range cases {
		got:= Uniq(c.in,false,false,false,false,0,1);
		if got != c.want {
			t.Errorf("uniq(%q) == %q, want %q", c.in,got, c.want);
		}
	}
}

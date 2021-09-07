package main

import "testing"

func  TestNoArgs(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{
			[]string{"I love music.",
			"I love music.\n",
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
		got:= Uniq(c.in,false,false,false,0,0);
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in,got, c.want)
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
			"I love music.\n",
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
		got:= Uniq(c.in,true,false,false,0,0);
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in,got, c.want)
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
					 "I love music.\n",
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
		got:= Uniq(c.in,false,true,false,0,0);
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in,got, c.want)
		}
	}
}



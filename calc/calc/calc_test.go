package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func  TestComplete(t *testing.T) {
	cases := []struct {
		in string
		want float64
	}{
		{"(12+15)/2*4", 54.0},
		{"2*2", 4.0},
		{"350*15*((150-10+5)/1000)*20*(3-6)", -45675.0},
	}
	for _, c := range cases {
		got, err:= Calc(c.in)
		if err!=nil {
			panic(err)
		}
		if !assert.Equal(t,got, c.want, "Should be equal") {
			t.Errorf("%v == %f, want %f", c.in,got, c.want)
		}
	}
}

func TestWrongInput(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"(12+15)///2*4"},
		{"2A*2"},
		{"350*15*((150-10+5)/1000)*20*(3-6))))))"},
	}
	for _, c := range cases {
		got, err:= Calc(c.in)
		if err==nil {
			t.Errorf("%v == %f", c.in,got)
		}
	}
}
package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func  TestPushPop(t *testing.T) {
	var stack Stack;
	tokens := []Token{
			{isOperation: false,Value: "1.2332"},
			{isOperation: true,Value: "+"}, 
	}
	stack.Push(tokens[0]);
	stack.Push(tokens[1]);
	stack.Pop()
	tokens0Value, err := stack.Pop();
	if err!=nil {
		panic(err)
	}
	assert.Equal(t, tokens[0], tokens0Value, "they should be equal")
}
package calc

import (
	"errors"
)

// Структура, чтобы хранить токены математического выражения
type Token struct {
	isOperation bool
	Value string
}

type Stack []Token

// Проверяет пустой ли стэк
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Пушит в стэк
func (s *Stack) Push(t Token) {
	*s = append(*s, t)
}

func (s *Stack) Peek() (Token, error) {
	if s.IsEmpty() {
		return Token{}, errors.New("Stack is empty. Can`t peek")
 	} 
	return (*s)[len(*s)-1], nil
}

// Возвращает и удаляет последний элемент из стэка
func (s *Stack) Pop() (Token, error) {
	if s.IsEmpty() {
		return Token{}, errors.New("Stack is empty. Can`t pop")
	} else {
		index := len(*s) - 1 
		element := (*s)[index]
		*s = (*s)[:index] 
		return element, nil
	}
}

// Возвращает и удаляет первый элемент из стэка
func (s *Stack) Shift() (Token, error) {
	var t Token;
	if !s.IsEmpty() {
		t = (*s)[0]
		*s = (*s)[1:]
	} else {
		return Token{}, errors.New("Stack is empty. Can`t shift");
	}
	return t, nil
}

type CalcStack []float64

// Проверяет пустой ли стэк
func (s *CalcStack) IsEmpty() bool {
	return len(*s) == 0
}

// Пушит в стэк
func (s *CalcStack) Push(t float64) {
	*s = append(*s, t) 
}

// Возвращает и удаляет первый элемент из стэка
func (s *CalcStack) Pop() (float64, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty. Can`t pop")
	} else {
		index := len(*s) - 1 
		element := (*s)[index]
		*s = (*s)[:index] 
		return element, nil
	}
}
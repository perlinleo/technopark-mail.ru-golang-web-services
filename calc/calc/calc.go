package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var priorityMap = map[string]int {
	"(": 1,
	")": 1,
	"+": 2,
	"-": 2,
	"*": 3,
	"/": 3,
}


// Разбивает строку с математическим выражением
// на составные части
func GetTokenStack(expression string) (Stack, error) { 
	var result Stack;
	var numberAccum string;

	expression = strings.ReplaceAll(expression," ", "");
	if strings.Count(expression,"(") != strings.Count(expression,")") {
		return nil, errors.New("Amount of ')' doesn`t match with amount of '('")
	}
	for index ,char := range expression {
		if priorityMap[string(char)]>0 {
			if numberAccum!="" {
				result.Push(Token{isOperation: false, Value: numberAccum});
				
				numberAccum = "";
			}
			if !result.IsEmpty() {
				topElem, err := result.Peek()
				if err!=nil{
					return nil, err
				}
				if !topElem.isOperation || topElem.Value==")" || 
				topElem.Value=="(" || char=='(' || char==')' {  
					result.Push(Token{isOperation: true, Value: string(char)});
				} else {
					return nil, errors.New("Can`t have two operators in a row")
				}
			} else {
				result.Push(Token{isOperation: true, Value: string(char)})
			}
		} else if unicode.IsDigit(char) || char == '.' || char == ',' {
			numberAccum+=string(char);
			if index+1==len(expression) {
				result.Push(Token{isOperation: false, Value: numberAccum});
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Unrecognised symbol %c", char));
		}
	}
	
	return result, nil;
}

// Составляет постфиксную форму для токенов математического выражения
func PostFixStack(tokensStack Stack) (Stack,error) {
	var opeStack Stack;
	var tokens Stack;

	for !tokensStack.IsEmpty() {
		t , err:= tokensStack.Shift()
		if err!=nil {
			return nil, err
		}

		tokenPriority := priorityMap[t.Value]

		var stackPriority int;
		if !opeStack.IsEmpty() {
			stackPriority = priorityMap[opeStack[len(opeStack)-1].Value]
		}
		if t.isOperation == false {
			tokens.Push(t)
		} else {
			if opeStack.IsEmpty() || tokenPriority > stackPriority || t.Value == "(" {
				opeStack.Push(t)
			}  else if t.Value == ")" {
				var ope Token

				ope, err = opeStack.Pop();
				if err != nil {
					return nil, err
				}

				for ope.Value != "(" && len(opeStack) != 0 {
					tokens.Push(ope)
					ope, err = opeStack.Pop();
					if err!=nil {
						return nil, err
					}
				}
			} else {
				for tokenPriority <= stackPriority {
					
					token,err := opeStack.Pop();
					if err != nil {
						return nil, err
					}
					
					tokens.Push(token)

					stackPriority = 0
					if !opeStack.IsEmpty() {
						stackPriority = priorityMap[opeStack[len(opeStack)-1].Value]
					}
				}

			opeStack.Push(t)
			}
		}
	}

	for !opeStack.IsEmpty() {
		token,err := opeStack.Pop()
		if err!=nil {
			return nil, err
		}
		tokens.Push(token);
	}

	return tokens, nil
}

// Вычисляет выражение в постфиксной форме
func SolvePostfix(tokens Stack) (float64,error) {
	var calcStack CalcStack

	for !tokens.IsEmpty() {
		t , err:= tokens.Shift();
		if err!=nil {
			return 0, err
		}
		if !t.isOperation {
			val, err := strconv.ParseFloat(t.Value, 64)
			if err != nil {
				return 0, err
			}
			calcStack.Push(val);
		} else {
			var calc float64;
			l,err := calcStack.Pop()
			if err!=nil {
				return 0, err
			}

			r,err := calcStack.Pop()
			if err!=nil{
				return 0, err
			}
			
			switch t.Value {
			case "+":
				calc = r + l
			case "-":
				calc = r - l
			case "*":
				calc = r * l
			case "/":
				calc = r / l
			}
			calcStack.Push(calc)
		}
	}

	result, err:= calcStack.Pop()
	if err!=nil {
		return 0, err
	}

	return result, nil
}

func Calc(expression string) (float64, error) {
	tokenStack, err:=GetTokenStack(expression)
	if err!=nil {
		return 0, err
	}
	postfixStack, err:=PostFixStack(tokenStack)
	if err!=nil {
		return 0, err
	}
	result, err:=SolvePostfix(postfixStack)
	if err!=nil {
		return 0, err
	}
	return result, nil
}



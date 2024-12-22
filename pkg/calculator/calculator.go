package calculator

import (
	"errors"
	"strconv"
)

type StackByte []byte
type StackFloat64 []float64

var (
	ErrInvalidParentheses       = errors.New("invalid parentheses")
	ErrInvalidSymbol            = errors.New("invalid symbol")
	ErrInvalidSymbolCombination = errors.New("invalid symbol combination")
	ErrInvalidNumber            = errors.New("invalid number")
	ErrDivisionByZero           = errors.New("division by zero")
	ErrEmptyExpression          = errors.New("empty expression")
)

func (s *StackByte) Pop() byte {
	res := byte(0)
	if len(*s) > 0 {
		res = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	}
	return byte(res)
}

func (s *StackByte) Push(b byte) {
	*s = append(*s, b)
}

func (s StackByte) Top() byte {
	res := byte(0)
	if len(s) > 0 {
		res = s[len(s)-1]
	}
	return byte(res)
}

func (s StackByte) IsEmpty() bool {
	return len(s) == 0
}

func (s *StackFloat64) Push(f float64) {
	*s = append(*s, f)
}

func (s *StackFloat64) Pop() float64 {
	res := 0.0
	if len(*s) > 0 {
		res = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	}
	return res
}

func (s StackFloat64) Top() float64 {
	res := 0.0
	if len(s) > 0 {
		res = s[len(s)-1]
	}
	return res
}

func (s StackFloat64) IsEmpty() bool {
	return len(s) == 0
}

func ValidParentheses(expression string) error {
	st := StackByte{}
	for i := range expression {
		if expression[i] == '(' {
			st.Push(expression[i])
		} else if expression[i] == ')' {
			if st.IsEmpty() {
				return ErrInvalidParentheses
			} else {
				_ = st.Pop()
			}
		}
	}
	if st.IsEmpty() {
		return nil
	} else {
		return ErrInvalidParentheses
	}
}

func Find(a []byte, c byte) int {
	for i := range a {
		if a[i] == c {
			return i
		}
	}
	return -1
}

func ValidExpression(expression string) error {
	can := []byte{'-', '+', '*', '.', '/', ')', '('}
	size := len(can)
	for i := range expression {
		if Find(can, expression[i]) == -1 && expression[i]-'0' > 9 {
			return ErrInvalidSymbol
		}
	}
	if Find(can, expression[len(expression)-1]) != -1 && expression[len(expression)-1] != ')' {
		return ErrInvalidSymbolCombination
	}
	if Find(can, expression[0]) != -1 && expression[0] != '(' && expression[0] != '-' {
		return ErrInvalidSymbolCombination
	}
	for i := 1; i < len(expression); i++ {
		currPos := Find(can, expression[i])
		beforePos := Find(can, expression[i-1])
		if currPos != -1 && beforePos != -1 {
			if currPos == size-1 && beforePos == size-2 || currPos == size && beforePos == size-4 {
				return ErrInvalidSymbolCombination
			} else if beforePos == size-2 && currPos == size-1 || beforePos == size-2 && currPos == size-4 {
				return ErrInvalidSymbolCombination
			} else if currPos != size-1 && beforePos != size-2 {
				return ErrInvalidSymbolCombination
			}
		} else if Find(can, expression[i-1]) != -1 && Find(can, expression[i-1]) != len(can)-2 && Find(can, expression[i]) == len(can)-2 {
			return ErrInvalidSymbolCombination
		} else if Find(can, expression[i]) == len(can)-1 && expression[i-1]-'0' <= 9 {
			return ErrInvalidSymbolCombination
		}
	}
	return nil
}

func Valid(expression string) error {
	err := ValidParentheses(expression)
	if err != nil {
		return err
	}
	err = ValidExpression(expression)
	return err
}

func GetNumber(expression string, pos *int) (string, error) {
	number := ""
	for ; *pos < len(expression); *pos = *pos + 1 {
		num := expression[*pos]
		t := num - '0'
		if t <= 9 || num == '.' {
			number = number + string(num)
		} else {
			*pos = *pos - 1
			break
		}
	}
	pointCounter := 0
	for i := range number {
		if number[i] == '.' {
			pointCounter++
		}
	}
	if pointCounter >= 2 {
		return "", ErrInvalidNumber
	} else {
		return number, nil
	}
}

func ToPostfix(expression string) (string, error) {
	st := StackByte{}
	postfix := ""
	for i := 0; i < len(expression); i++ {
		c := expression[i]
		t := c - '0'
		if t <= 9 {
			number, err := GetNumber(expression, &i)
			if err != nil {
				return "", err
			}
			postfix += number + " "
		} else if c == '(' {
			st.Push(c)
		} else if c == ')' {
			for !st.IsEmpty() && st.Top() != '(' {
				postfix += string(st.Pop())
			}
			if !st.IsEmpty() {
				_ = st.Pop()
			}
		} else if c != ' ' {
			if c == '-' && (i == 0 || expression[i-1] == '(') {
				c = '~'
			}
			currentPriority := 0
			if c == '-' || c == '+' {
				currentPriority = 1
			} else if c == '*' || c == '/' {
				currentPriority = 2
			} else {
				currentPriority = 3
			}
			for !st.IsEmpty() {
				stackPriority := 0
				stackTop := st.Top()
				if stackTop == '(' {
					stackPriority = 0
				} else if stackTop == '-' || stackTop == '+' {
					stackPriority = 1
				} else if stackTop == '*' || stackTop == '/' {
					stackPriority = 2
				} else {
					stackPriority = 3
				}
				if stackPriority >= currentPriority {
					postfix += string(st.Pop())
				} else {
					break
				}
			}
			st.Push(c)
		}
	}
	for !st.IsEmpty() {
		postfix += string(st.Pop())
	}
	return postfix, nil
}

func Execute(op byte, first, second float64) (float64, error) {
	res := 0.0
	switch op {
	case '-':
		res = first - second
	case '+':
		res = first + second
	case '*':
		res = first * second
	case '/':
		if second == 0.0 {
			return 0.0, ErrDivisionByZero
		}
		res = first / second
	}
	return res, nil
}

func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0.0, ErrEmptyExpression
	}
	err := Valid(expression)
	if err != nil {
		return 0.0, err
	}
	var postfix string
	postfix, err = ToPostfix(expression)
	if err != nil {
		return 0.0, err
	}
	st := StackFloat64{}
	for i := 0; i < len(postfix); i++ {
		c := postfix[i]
		t := c - '0'
		if t <= 9 {
			number, _ := GetNumber(postfix, &i)
			num, _ := strconv.ParseFloat(number, 64)
			st.Push(num)
		} else if c != ' ' {
			if c == '~' {
				second := st.Pop()
				res, _ := Execute('-', 0.0, second)
				st.Push(res)
			} else {
				second := st.Pop()
				first := st.Pop()
				var res float64
				res, err = Execute(c, first, second)
				if err != nil {
					return 0.0, err
				}
				st.Push(res)
			}
		}
	}
	return st.Pop(), nil
}

// about 40 minutes
// I was told that I should use stack
// Google for example of stack realisation in Go
package main

import "fmt"

func main() {
	fmt.Println(isValidString("([{}))))"))
	fmt.Println(isValidString("([{}[]()])))"))
	fmt.Println(isValidString("([{}["))
}

func isValidString(str string) bool {
	var stringInvalid bool
	stk := new(stack)

	for i := 0; i < len(str); i++ {
		if stringInvalid {
			break
		}

		currentChar := bracket(str[i])
		prevChar := stk.Pop()

		if prevChar == "" {
			stk.Push(currentChar)
			continue
		}

		if currentChar.GetType() == prevChar.GetType() && prevChar.IsClosing() != currentChar.IsClosing() {
			continue
		}

		if !prevChar.IsClosing() && currentChar.IsClosing() && currentChar.GetType() != prevChar.GetType() {
			stringInvalid = true
			break
		}

		stk.Push(prevChar)
		stk.Push(currentChar)
	}

	if !stringInvalid && !stk.IsEmpty() {
		stringInvalid = true
	}

	return !stringInvalid
}

type bracket string
func (b *bracket) GetType() string {
	switch *b {
	case "[", "]":
		return "square"
	case "{", "}":
		return "figure"
	default:
		return "round"
	}
}
func (b *bracket) IsClosing() bool {
	switch *b {
	case "]", "}", ")":
		return true
	default:
		return false
	}
}

type stack [] bracket
func (s *stack) Push(value bracket) bool {
	*s = append(*s, value)
	return true
}
func (s *stack) Pop() bracket {
	if len(*s) == 0 {
		return ""
	}

	idx := len(*s) - 1
	element := (*s)[idx]
	*s = (*s)[:idx]

	return element
}
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *stack) GetLength() int {
	return len(*s)
}
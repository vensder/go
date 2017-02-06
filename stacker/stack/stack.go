package stack

import "fmt"
import "errors"

type Stack []interface{}

func Hello() {
	fmt.Println("Hello from stack package from Hello func")
}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return (len(stack) == 0)
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	//if stack.IsEmpty() {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	//return stack[stack.Len()-1], nil
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	//if stack.IsEmpty() {
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	//x := theStack[stack.Len()-1]
	//*stack = theStack[:stack.Len()-1]
	return x, nil
}

package Utils

import "fmt"

type StringStack struct {
	strings []string
}

//Pushes the given string onto the stack
//@param s -- string to add to stack
func (ss *StringStack)Push(s string) {
	ss.strings = append(ss.strings, s)
}

//Removes the last added item on the stack
//@return string last added to stack, error if stack is empty
func (ss *StringStack)Pop() (string, error) {
	length := len(ss.strings)
	if length < 1 { return "", fmt.Errorf("stringstack is empty") }
	toReturn := ss.strings[length - 1]
	ss.strings = ss.strings[:length - 1]
	return toReturn, nil
}

//Checks if the stack is empty
//@return true if empty, otherwise false
func (ss *StringStack)IsEmpty() bool {
	return len(ss.strings) == 0
}
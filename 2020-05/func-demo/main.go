package main

var fu *funcTest

type funcTest struct {
	s map[string]Inter
}

// Inter asdsa
type Inter interface {
	Run()
}

func (f *funcTest) InitFunc() {
	fu = &funcTest{
		s: make(map[string]Inter),
	}
}

func main() {

}

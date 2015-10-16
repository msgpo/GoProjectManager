/*
Test program
*/

package main

import (
	"C"
	"os"
)

var (
	Var1 int
	Var2 bool
)

const (
	Const1 = "123"
	Const2 = 222
)

type Test struct {
	test1 string
	test2 string
}

type ITest interface {
}

type TTest *os.File
type TTest1 *Test
type TTest2 bool
type TTest3 string
type TTest4 *string
type TTest5 *[]string
type TTest6 *struct {
	A int
	b int
}

func NewTest(fileName string) *Test {
	p := Test{}
	//	Tes
	p.test1 = "test11"
	p.test2 = "test22"
	return &p
}

func (t *Test) Test1(t1 Test) (x int, y int) {
	return 12, 32
}

func (t Test) Test2(t1 *Test) bool {
	return false
}

func (t *Test) Test3(int) int {
	return 0
}

//export test
func test() {

}

func main() {

}

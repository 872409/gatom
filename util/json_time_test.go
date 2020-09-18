package util

import (
	"fmt"
	"testing"
)

type people interface {
	speak1()
	speak2()
}

type student struct {
	name string
	age  int
}

func (stu student) speak1() {
	fmt.Println("I am a student1, I am ", stu.age)
}
func (stu *student) speak2() {
	fmt.Println("I am a student2, I am ", stu.age)
}
func f(out people) {

	if out != nil { // out!=nil
		fmt.Println("surprise!")
	}

}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
func TestJSONTime_FormatDate4(t1 *testing.T) {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
func TestJSONTime_FormatDate3(t1 *testing.T) {
	var s1 student
	var s2 = &s1
	s1.speak1()
	s1.speak2()
	s2.speak1()
	s2.speak2()
	var p people
	// p = &student{name: "RyuGou", age: 12} // 如果不是用指针会报错
	// p.speak1()
	// p.speak2()
	if p != nil { // out!=nil
		fmt.Println("surprise!")
	}
	f(p)
}

func TestJSONTime_FormatDate(t1 *testing.T) {
	var date JSONTime = JSONTime{}
	fmt.Println(date.FormatDate())
}

func TestJSONTime_FormatDate2(t1 *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	var stus []student

	stus = []student{
		{Name: "one", Age: 18},
		{Name: "two", Age: 19},
	}

	data1 := make(map[int]student)
	data2 := make(map[int]*student)

	for i, _ := range stus {
		vv := stus[i]
		fmt.Printf("%d, %v, %X\n", i, vv, &vv)
		data1[i] = vv
		data2[i] = &vv // 应该改为：data[i] = &stus[i]
	}

	for i := range data1 {
		fmt.Printf("key=%d, value=%v \n", i, data1[i])
		fmt.Printf("key=%d, value=%v \n", i, data2[i])
	}
}

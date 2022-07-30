package main

import (
	_ "awesomeProject/lib1"
	. "awesomeProject/lib2"
	//	mylib2 "awesomeProject/lib2"
	"fmt"
	"time"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
}

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

type Dog struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

type Human struct {
	name string
	sex  string
}

type SuperMan struct {
	Human
	level int
}

type Book struct {
	title string
	auth  string
}

// Hero 类名首字母大写，外部包能访问，属性也是如此, 否则只能内部包访问
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat().......")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk().......")
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat().......")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly().......")
}

// GetName 方法名大写，外部包可访问，否则，只能本包内访问
func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(name string) {
	this.Name = name
}

func changeBook(book *Book) {
	book.auth = "bc"
}

func changeValue(a *int) {
	*a = 10
}

func swap(m, n *int) {
	var tmp int
	tmp = *m
	*m = *n
	*n = tmp
}

//多个返回值, 匿名
func foo(a int, b int) (int, int) {
	a = 100
	b = 200
	return 100, 200
}

//多个返回值，有名称
func foo1(a int, b int) (r1, r2 int) {
	a = 100
	b = 200

	r1 = 300
	r2 = 400
	return
}

func printArray(myArray [5]int) {
	myArray[4] = 100
	for i := 0; i < len(myArray); i++ {
		print(myArray[i], " ")
	}
	println("")
}

func printDyArray(dyArray []int) {
	for _, value := range dyArray {
		fmt.Println("value is", value)
	}
	dyArray[0] = 500
}

func updateArray(array []int) {
	array[0] = 1000
}

func deferfunc() {
	fmt.Println("this is deferfunc")
}

func returnfunc() int {
	fmt.Println("this is return func")
	return 0
}

func deferandreturn() int {
	defer deferfunc()
	return returnfunc()
}

func changeMap(citymap map[string]string) {
	citymap["China"] = "HongKong"
	citymap["Aus"] = "NewYork"
}

func main() {
	fmt.Println("hello")
	println("hello,world")
	time.Sleep(1 * time.Second)

	var a int
	println("a =", a)

	var b int = 100
	println("b =", b)

	var c = 100
	fmt.Printf("c is %d, type of c is %T\n", c, c)

	d := 200
	fmt.Println("d is", d)

	var xx, yy = 100, 200
	fmt.Println("xx, yy is", xx, yy)

	var (
		mm = "abc"
		vv = true
	)
	fmt.Println("mm, vv is", mm, vv)

	const name = "china"
	println(name)

	const (
		BEIGIN = 10 * iota //iota = 0
		SHANG
		SHENG = 30 * iota //iota = 2
	)
	fmt.Println("BEIGIN =", BEIGIN)
	fmt.Println("SHANG =", SHANG)
	fmt.Println("SHENG =", SHENG)

	ret1, ret2 := foo(2, 3)
	println("foo------")
	println("ret1 is", ret1)
	println("ret2 is", ret2)

	ret1, ret2 = foo1(12, 23)
	println("foo1------")
	println("ret1 is", ret1)
	println("ret2 is", ret2)

	// _ package name ->call lib1 init method and other method is not called the package
	//	lib1.LibTest1()

	// lib2.LibTest2()

	// mylib2.LibTest2() // mylib2 is the alias name of lib2
	fmt.Println("welcome china")

	// . package is imported to current directory, lib2 is not needed
	LibTest2()

	h := 1
	changeValue(&h)

	//h is changed, go-pointer
	println("h is", h)

	m := 2
	n := 3
	swap(&m, &n)
	println("m is", m)
	println("n is", n)

	var p *int = &m
	println(p)
	println(&m)

	//second level pointer **pp
	var pp = &p
	println(pp)
	println(&p)
	println(**pp)

	// defer stack, FIFO, 先进后出
	defer println("main end1")
	defer println("main end2")

	//defer 和 return， return 先执行， defer后执行
	var myArray [3]int

	for i := 0; i < len(myArray); i++ {
		println(myArray[i])
	}

	//static array has length, and rather than dynamic array
	myArray2 := [5]int{1, 2, 3, 4}
	for index, value := range myArray2 {
		println("index is", index, "value is", value)
	}

	fmt.Printf("myArray type is %T\n", myArray)
	fmt.Printf("myArray2 type is %T\n", myArray2)

	// array is value pass, 传递的是副本，不能在函数中更改原数组值
	printArray(myArray2)
	println(myArray2[4])

	//dy array, slice
	dyArray := []int{1, 2, 3, 4}
	fmt.Printf("type of dyArray is %T\n", dyArray)

	//dyArray 内部含有指针，golang都是值传递
	printDyArray(dyArray)
	println("now dyArray[0] is", dyArray[0], "and dyArray length is", len(dyArray))

	//%v, 打印详细信息
	fmt.Printf("dyArray is %v\n", dyArray)

	//slice 声明无容量
	//	var dyslice []int
	//make give the dyslice length, 类型int，个数3
	//	dyslice = make([]int, 3)
	// var dyslice = make([]int, 3), 合二为一
	dyslice := make([]int, 3)
	dyslice[0] = 100
	fmt.Printf("dyslice is %v\n", dyslice)

	//判断切片是否为空
	if dyslice == nil {
		println("slice has no length")
	} else {
		println("slice has length")
	}

	var number = make([]int, 3, 5)
	updateArray(number)
	fmt.Printf("number length is %d, cap is %d, slice = %v\n", len(number), cap(number), number)

	number = append(number, 4)
	number = append(number, 5)
	fmt.Printf("number length is %d, cap is %d, slice = %v\n", len(number), cap(number), number)

	//append超过cap, cap变2倍
	num := append(number, 6)
	fmt.Printf("number length is %d, cap is %d, slice = %v\n", len(num), cap(num), num)

	slice1 := make([]int, 3)
	fmt.Printf("slice1 length is %d, cap is %d, slice1 = %v\n", len(slice1), cap(slice1), slice1)

	//不写容量，容量与length相同，append后，容量增加一倍
	slice1 = append(slice1, 43)
	fmt.Printf("slice1 length is %d, cap is %d, slice1 = %v\n", len(slice1), cap(slice1), slice1)

	s1 := slice1[0:]
	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(slice1, s1)

	//左闭右开，index=2取不到,that is index = (0, 1)
	s2 := s1[0:2]
	fmt.Println(s2)

	//取全部
	s3 := slice1[:]
	fmt.Println(s3)

	//index = (0, 1, 2), index=3 not got
	s4 := slice1[:3]
	fmt.Println(s4)

	aa := make(
		[]int,
		5,
	)
	fmt.Println(aa)

	//return first, defer second
	deferandreturn()

	//first declare
	var mymap1 map[string]string
	if mymap1 == nil {
		fmt.Println("this is empty map")
	}
	mymap1 = make(map[string]string, 3)
	mymap1["China"] = "Taiwan"
	mymap1["Japan"] = "Tokyo"
	mymap1["USA"] = "DC"
	fmt.Println(mymap1)

	//second declare
	mymap2 := make(map[int]string)
	mymap2[1] = "a"
	mymap2[2] = "b"
	mymap2[3] = "c"
	fmt.Println(mymap2)

	//third declare
	mymap3 := map[string]string{
		"one":   "Java",
		"two":   "Go",
		"three": "C++",
	}
	fmt.Println(mymap3)

	fmt.Println("-------------------")

	//遍历
	for key, value := range mymap1 {
		fmt.Println("key is", key)
		fmt.Println("value is", value)
	}

	//delete
	delete(mymap1, "USA")

	//modify
	mymap1["China"] = "Shanghai"

	fmt.Println("-----------")
	fmt.Println("after modify,", mymap1)

	changeMap(mymap1)
	fmt.Println("-----------")
	fmt.Println("after change,", mymap1)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "abc"
	fmt.Printf("%v\n", book1)

	fmt.Println("----------------")
	changeBook(&book1) //传地址 &book1
	fmt.Printf("after change book,%v\n", book1)

	fmt.Println("--------------------------")
	hero := Hero{Name: "LiuBang", Ad: 100, Level: 1}
	fmt.Println(hero)

	fmt.Println("get hero name is", hero.GetName())
	hero.SetName("QIN")
	fmt.Println("after set,name is", hero.GetName())

	fmt.Println("--------------------------")

	hh := Human{"zhang3", "female"}
	hh.Eat()
	hh.Walk()

	fmt.Println("--------------------------")
	//	h2 := SuperMan{Human{"li4", "female"}, 88}
	var h2 SuperMan
	h2.sex = "male"
	h2.name = "li4"
	h2.level = 88
	h2.Fly()
	h2.Eat()

	fmt.Println("--------------------------")
	// 多态，指针
	var animal AnimalIF
	animal = &Cat{color: "Green"}
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())

	fmt.Println("--------------------------")
	animal = &Dog{"Yellow"}
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())

	fmt.Println("--------------------------")
	cat := &Cat{color: "green"}
	dog := &Dog{
		"black",
	}
	showAnimal(cat)
	showAnimal(dog)

	fmt.Println("--------------------------")
	myFunc("abc")
	myFunc(1)
	myFunc(3.14)

	fmt.Println("--------------------------")
	bb := &Book{"ab", "cc"}
	var r Reader
	r = bb
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()

	fmt.Println("--------------------------")

	//移动副号：13817124914
	//移动主号：15901819077
	//0571 87013615 江苏银行
	//0523-84993855 徐法官

	//622848 3429 3285 57770
	//8359.65
	//11030
	//623066 7031 008390177

	//徐:2000, 李：2000
	//广发：570  招商：1204
	//华夏：1042.33
	//房租：2000
	//桔多多：

}

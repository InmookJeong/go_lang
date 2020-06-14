package main

import (
	"fmt"
	"strings"

	"github.com/mook/learngo/accounts"
)

// Create struct
type person struct {
	name    string
	age     int
	favFood []string
}

// java : constructor()
// python : __init__
// Go에는 해당 기능 없음 -> constructor나 Method나 동등하게 인식
// channels and go routines는 미니프로젝트 하면서 학습

// func func_name, func_parameter, return type
func main() {
	fmt.Println("Hello World...!")

	// something.SayHello()
	// something.SayBye()

	const constName string = "mook(const)"
	fmt.Println(constName)

	// var name string = "mook"
	name := "mook" // 축약형은 func 안에서만 사용 가능
	fmt.Println(name)

	fmt.Println("-------------------")
	fmt.Println(multiply(2, 2))

	// 변수 하나만 만들어주면 error 발생
	// totalLength, _로 변수들 선언해주면 하나만 리턴 받는거 가능
	// _가 변수를 무시해주는 역할
	totalLength, upperName := lenAndUpper("mook")
	fmt.Println(totalLength, upperName)

	fmt.Println("mook", "jim", "james", "inmook")

	fmt.Println(lenAndUpper2("mooka"))

	total := loopFunc(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	fmt.Println(canIDrink(16))

	fmt.Println(canIDrink2(18))

	fmt.Println(canIDrink3(18))

	// Pointer
	a := 2
	b := a
	a = 10
	// &을 붙이면 메모리 주소 확인 가능
	fmt.Println(a, b)
	fmt.Println(&a, &b)

	aa := 2
	bb := &aa // aa의 주소 참조
	aa = 7
	fmt.Println(aa, bb, *bb) // *은 참조된 메모리의 값을 열람
	*bb = 10                 // 참조된 주소의 값을 수정할 수 있음
	fmt.Println(aa, *bb)

	// Array
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alala3"
	names[4] = "alala4"
	fmt.Println("names -> ", names)

	// non length Array
	names2 := []string{"nico", "lynn", "dal"}
	// names2[3] = "alala3"		--> intext out of range Error
	// slice(=첫번째 파라미터)에 value(두번째 파라미터)를 추가해서 새로운 slice를 Return
	names2 = append(names2, "alala-3")
	names2 = append(names2, "alala-4")
	names2 = append(names2, "alala-5")
	names2 = append(names2, "alala-6")
	fmt.Println("names2 --> ", names2)

	// map[key]value{key:value, ---}
	sampleMap := map[string]string{"name": "nico", "age": "20"}
	fmt.Println(sampleMap)
	for key, value := range sampleMap {
		fmt.Println(key, " : ", value)
	}
	fmt.Println(sampleMap["name"])
	sampleMap["nation"] = "Korea"
	fmt.Println("map --> ", sampleMap)
	// delete map
	delete(sampleMap, "nation")
	fmt.Println("map --> ", sampleMap)
	// key check
	val, exists := sampleMap["age"]
	fmt.Println(val, " exists : ", exists)
	val2, exists2 := sampleMap["nation"]
	fmt.Println(val2, " exists : ", exists2)
	if !exists2 {
		println("no....")
	}

	// struct --> structure(구조체), 일종의 Object Type Data or Class
	favFood := []string{"test", "sample"}
	sampleStruct := person{name: "nico", age: 29, favFood: favFood}
	fmt.Println(sampleStruct)
	fmt.Println(sampleStruct.age)

	// create account
	fmt.Println("---Account Test---")
	account := accounts.NewAccount("mook")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
}

// return integer
func multiply(a, b int) int {
	return a * b
}

// multi return = integer, string
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Print("repeatMe...")
	fmt.Println(words)
}

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	// defer -> 현재 func이 값을 return한 다음에 실행
	defer fmt.Println("I'm done")
	fmt.Println("lenAndUpper2...")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// loop는 for만 존재
// range -> array에 loop 적용, index와 value를 같이 줌
func loopFunc(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func canIDrink(age int) bool {
	fmt.Println("canIDring...")
	// if age < 20 {
	// variable expression 생성
	// koranAge를 block 안에서만 사용
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	fmt.Println("canIDring2...")
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIDrink3(age int) bool {
	fmt.Println("canIDring3...")
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	}
	return false
}

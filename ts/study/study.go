package study

// import "fmt"

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"
}

// GetMusic : Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
	return pop[singer]
}

// func getKeys() { // 내부에서만 호출 가능
//
//		for _, kv := range pop {
//			fmt.Println(kv)
//		}
//	}
func PassByRef(a *int) {
	println(a)
}
func PassByValue(a int) {
	println(&a)
}

func SayString(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func SayInt(msg ...int) {
	println(msg)
	// 가변길이 인자로 다시 호출하기
	ret, mx := SumAndMax(msg...)
	// ret := sum(msg...)

	println(ret, mx)

}

func Sum(nums ...int) int {
	var ret int
	for _, v := range nums {
		ret += v
	}
	return ret
}

func SumAndMax(nums ...int) (sumOfElems int, MaxElem int) {

	for _, v := range nums {
		sumOfElems += v
		if MaxElem < v {
			MaxElem = v
		}
	}
	return sumOfElems, MaxElem
}

func IfStatement(k int, i int) {
	//if statement
	if k == 1 {
		println("One")
	} else if k == 2 { //같은 라인
		println("Two")
	} else { //같은 라인
		println("Other")
	}

	if val := i * 2; val < k {
		println(val)
	}
	// 아래 처럼 사용하면 Scope 벗어나 에러
	//val++
}

func SwitchStatment(score int) {
	//특징 -> break 가 필요없음
	//fallthrough -> 스위치 걸리는것부터 아래까지 실행
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

func ForStatement(num int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	println(sum)
}

func ForButWhileStatement(to int) {
	n := 1
	for n < to {
		n *= 2
		//if n > 90 {
		//   break
		//}
	}
	println(n)
}

// func main() {

// 	IfStatement(2, 1)
// 	SwitchStatment(64)
// 	ForStatement(10)
// 	ForButWhileStatement(100)

// 	var a int = 100
// 	println(&a)
// 	//go 는 포인터만 존재 , C++ 처럼 레퍼런스를 사용할수 없다
// 	PassByValue(a)
// 	PassByRef(&a)

// 	var ap *int = &a
// 	println(ap)
// 	println(*ap)

// 	//가변인자 함수
// 	SayString("This", "is", "a", "book")
// 	SayString("Hi")
// 	SayInt(0, 1, 2, 3, 4)

// }

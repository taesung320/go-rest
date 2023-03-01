package source

import "go-ts/ts/study"

func main() {
	var a int = 10
	study.PassByRef(&a)
}

package dep2

import dep1 "github.com/faja/gomodules-dep1"

func Func1() string {
	return "[dep2][v1.0.0] Func1"
}

func Func2_from_dep1() string {
	return dep1.Func1()
}

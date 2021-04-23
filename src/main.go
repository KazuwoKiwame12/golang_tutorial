package main

import (
	"app/basics"
	"app/concurrency"
	"app/methods"
	"fmt"
)

func main() {
	// 参考url: https://go-tour-jp.appspot.com/basics/11
	// checkBasic()
	// checkFlowControl()
	// checkTypes()
	// checkMethods()
	checkConcurrency()
}

func checkBasic() {
	// 1
	fmt.Println(basics.Add(10, 20))
	// 2
	a := "Sing"
	b := "Song"
	a, b = basics.Swap(a, b)
	fmt.Println(a, b)
	// 3
	basics.Variables()
	// 4
	basics.VariablesWithInitializers()
	// 5
	basics.ShortVariableDeclarations()
	// 6
	basics.TypeConversions()
	// 7
	basics.Constants()
}

func checkFlowControl() {
	// 1
	basics.For()
	// 2
	fmt.Println(basics.IfWithAShortStatement(3, 2, 10))
	// 3
	basics.Switch()
	// 4
	basics.Defer()
}

func checkTypes() {
	// 1
	basics.Pointers()
	// 2
	basics.FunctionValue()
	// 3
	basics.FunctionClosures()
}

func checkMethods() {
	//1
	methods.Methods()
	// 2
	methods.InterfaceImplementImplicity()
	// 3
	methods.Interface()
	// 4
	methods.InterfaceValuesWithNil()
	// 5...error起こしているので、コメントアウト
	// methods.NilInterfaceValues()
	// 6
	methods.TypeAssertions()
	// 7
	methods.TypeSwitches()

}

func checkConcurrency() {
	// 1
	concurrency.GoRoutines()
	// 2
	concurrency.Channels()
	// 3
	concurrency.ChannelsForDecentralize()
	// 4
	concurrency.RangeAndClose()
	// 5
	concurrency.Select()
}

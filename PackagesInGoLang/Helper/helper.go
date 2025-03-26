package helper

import "fmt"

// if you want to exprt this function such that it can be used in other packages, start the function name with a capital letter.
// CommanFunc is an exported function
func CommanFunc() {
	fmt.Println("This is a common function")
}

// Same can be done with variables as well
// Thus three types of variables can be defined in GoLang:
// 1. Local Variables: Defined inside a function and can be accessed only inside that function.
// 2. Global Variables: Defined outside any function and can be accessed from any function in the package.
// 3. Exported Variables(in another Package): Defined outside any function and can be accessed from any package. Start the variable name with a capital letter to export it.
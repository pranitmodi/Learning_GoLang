// A package is a collection of Go files that are in the same directory.
// Under a directory, you can have multiple go files, just the main() function should be in the main.go file.

// You can multiple packages in a single directory, but only one package can be declared as main.

package main

import "PackagesInGoLang/helper" // importing helper package from the common PackagesInGoLang package defined in go.mod file

func main() {
	helper.CommanFunc()
}

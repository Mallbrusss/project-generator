package template

var defaultTemplate = `package {{ .PackageName }}
import "fmt"

// {{ .Description }}

func {{ .FuncName }}(){
	fmt.Println("This is the {{ .PackageName }} package")
}`

var templates = map[string]string{
	"main.go": `package main
import "fmt"

func main(){
	fmt.Println("hello, world")
}
	`,
}

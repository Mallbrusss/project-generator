package template

var defaultTemplate = `package {{ .PackageName }}
import "fmt"

// {{ .Description }}

func {{ .PackageName }}{{ .FuncName }}(){
	fmt.Println("This is the {{ .PackageName }} package")
}`
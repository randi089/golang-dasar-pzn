package helper

import "fmt"

var (
	version     = "1.0.0"
	Application = "golang"
)

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() string {
	fmt.Println(version)
	return sayGoodBye("Randi")
}

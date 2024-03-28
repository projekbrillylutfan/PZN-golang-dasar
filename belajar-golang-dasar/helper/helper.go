package helper
// agar syntax bisa di akses di luar package, nama harus di awali dengan huruf besar

var version = "1.0.0" // tidak bisa di akses diluar package
var Application = "golang" // tidak bisa di akses diluar package

func sayGoodBye(name string) string {
	return "good bye " + name
}

func SayHello(name string) string {
	return "hello " + name
}
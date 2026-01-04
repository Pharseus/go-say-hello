package gosayhello

// func SayHello() string {
// 	return "Hello World"
// }

// major change
// karena merusak kode sebelunya yang awlanya tidak ada parameter
func SayHello(name string) string {
	return "Hello World" + name
}

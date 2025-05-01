package structures

//go:generate easyjson -all
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

package main

type User struct {
	ID int64
	Name string
	Avatar string
}
func GetUserInfo() string {
	var id int64=1000
	println(&id)
	var user= &User{ID: id, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
	println(user)
	return "aaa"
}
func main() {
	_ = GetUserInfo()

}
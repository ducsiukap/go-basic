package exported

import "fmt"

// identifier viết hoa chữ cái đầu
// -> các package khác có thể gọi
// ~ public

// ngược lại, chữ cái đầu của Identifier là chữ thương
// -> chỉ dùng trong package nơi nó được khai báo (ex: exported)
// ~ private
type user struct { // user is private
	id_card string // id_card is private
	Name    string // public
	Age     int    // public
}

// vì user là private -> không thể tạo trực tiếp từ bên ngoài package
// -> cung cấp hàm tạo 1 instance
func CreateUser(id_card, name string, age int) user {
	u := user{id_card: id_card, Name: name, Age: age}
	return u
}

func (u user) Introduction() { // public
	fmt.Printf("Hi, im %s. %dyo!\n", u.Name, u.Age)
}

// cung cấp setter/getter cho các private field của struct
// tùy trường hợp có thể cung cấp cả 2, 1 trong 2 hoặc không cung cấp
func (u user) setIDCard(id_card string) { // private
	u.id_card = id_card
}
func (u user) GetIDCard() string { // private
	return u.id_card
}

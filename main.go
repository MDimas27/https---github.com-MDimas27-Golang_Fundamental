package main

import "fmt"

type User struct {
	ID				int
	FirstName		string
	LastName		string
	Email			string
	isActive		bool

} 

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

type Group struct {
	Name		string
	Admin		User
	Users		[]User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}

}


func main() {

	user := User{1, "Dimas", "Surya","dimas@email.com", true }
	result := user.display()
	fmt.Println(result)

	user2 := User{2, "Dede", "Nugraha","surya@email.com", true }
	fmt.Println(user2.display())

	users := []User{user, user2}
	group := Group{"Gamer", user, users, true}

	// Quiz/5 Membuat function menjadi Method
	// displayGroup(group)
	group.DisplayGroup()

}




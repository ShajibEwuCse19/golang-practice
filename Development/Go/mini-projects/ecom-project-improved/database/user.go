package database

type User struct {
	ID          uint8  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}
	u.ID = uint8(len(users) + 1)
	users = append(users, u) // add new user
	return u                 // return new user
}

func GetAllUsers(u User) []User {
	return users
}

func FindUserByEmailAndPassword(email string, password string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}

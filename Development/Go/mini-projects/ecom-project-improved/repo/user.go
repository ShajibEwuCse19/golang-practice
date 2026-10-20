package repo

type User struct {
	ID          uint8  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	List() ([]*User, error)
	FindUserByEmailAndPassword(email string, password string) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return nil, nil
	}
	user.ID = uint8(len(u.userList) + 1)
	u.userList = append(u.userList, &user) 
	return &user, nil
}
func (u *userRepo) List() ([]*User, error) {
	return u.userList, nil
}
func (u *userRepo) FindUserByEmailAndPassword(email string, password string) (*User, error) {
	for _, user := range u.userList {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, nil
}

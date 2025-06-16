package types

type User struct {
	Id        int    `json:"id,string"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SmallUser struct {
	Id        int    `json:"id,string"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

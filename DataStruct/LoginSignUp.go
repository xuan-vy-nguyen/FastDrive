package DataStruct

// LoginAccount type of login data
type LoginAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

// use for SignUp and DB
type SignUpAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
	UserName string `json:"username"`
	BirthDay string `json:"birthday"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	IsActive bool `json:"isactive"`
	CreateAt string `json:"createat"`
	IsDeleted bool `json:"isdeleted"`
}

// use for DB
type LoginDB struct {
	Mail string `json:"mail"`
	Token string `json:"token"`
}
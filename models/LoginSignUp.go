package datastruct

// LoginAccount type of login data
type LoginAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

// SignUpAccount is used for SignUp and DB
type SignUpAccount struct {
	Mail        string `json:"mail"`
	Pass        string `json:"pass"`
	UserName    string `json:"username"`
	BirthDay    string `json:"birthday"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	IsActive    bool   `json:"isactive"`
	CreateAt    string `json:"createat"`
	IsDeleted   bool   `json:"isdeleted"`
}

// LoginDB is used for DB
type LoginDB struct {
	Mail  string `json:"mail"`
	Token string `json:"token"`
}

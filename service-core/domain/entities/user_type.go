package entities

type UserType struct {
	Id   string
	Name string
}

func NewUserType(Id string, Name string) *UserType {
	return &UserType{
		Id:   Id,
		Name: Name,
	}
}

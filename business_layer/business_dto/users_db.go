package business_dto

type UserDb struct {
	Id string
	FirstName string
	LastName string
	Email string
	CreationDate string
	LastUpdateDate string
	Password string
	Version int
}

func (u UserDb) IsValid() bool{
	return len(u.Id) > 0;
}
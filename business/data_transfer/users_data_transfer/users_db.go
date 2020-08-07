package users_data_transfer

/*
Comentar que con DD y repository pattern esto no es requerido
pero en apis como users donde a veces el modelo difiere de la persistencia
es mas conveniente este enfoque
 */

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
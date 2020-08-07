package users_data_transfer

type PrivateUserDto struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	CreationDate string `json:"creation_date"`
}
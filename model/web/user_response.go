package web

type UserResponse struct {
	Id   	int 	`json:"id"`
	IdRole 	int 	`json:"id_role"`
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
}

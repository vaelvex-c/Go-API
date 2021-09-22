package web

type UserCreateRequest struct {
	IdRole 	int 	`validate:"required" json:"id_role"`
	Name 	string 	`validate:"required,min=1,max=200" json:"name"`
	Email 	string 	`validate:"required,email" json:"email"`
	Password string `validate:"required,password" json:"password"`
}

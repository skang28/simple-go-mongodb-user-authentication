package form

//NewUser defines the structure of the user form, use binding to make these fields required
type NewUser struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
}

//LoginUser defines structure of user login
type LoginUser struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
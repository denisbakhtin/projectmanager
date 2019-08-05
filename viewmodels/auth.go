package viewmodels

//LoginVM is a login view model
type LoginVM struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ActivateVM is an activation view model
type ActivateVM struct {
	Token string `json:"token" binding:"required"`
}

//RegisterVM is a registration view model
type RegisterVM struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ForgotVM is a view model for forgotten password request
type ForgotVM struct {
	Email string `json:"email" binding:"required"`
}

//ResetVM is a view model for password reset requests
type ResetVM struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

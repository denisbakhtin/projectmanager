package viewmodels

//AccountVM is an account view model
type AccountVM struct {
	Name            string `json:"name" binding:"required"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

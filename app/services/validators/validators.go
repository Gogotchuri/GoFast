package validators

/*SignInRequestT defines sign in form request*/
type SignInRequestT struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*Validate validates sign in request*/
func (sir *SignInRequestT) Validate() *[]string {
	var errs []string

	if sir.Email == "" {
		errs = append(errs, "Email can't be empty!")
	}

	if sir.Password == "" {
		errs = append(errs, "Password can't be empty!")
	}

	if len(errs) == 0 {
		return nil
	}
	return &errs
}

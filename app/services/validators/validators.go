package validators

/*VerificationRequestT defines verification code request*/
type VerificationRequestT struct {
	Email string `json:"email"`
}

/*SignInRequestT defines sign in form request*/
type SignInRequestT struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*SignInRequestT defines sign in form request*/
type SignUpRequestT struct {
	SignInRequestT
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

/*Validate validates verification request*/
func (vr *VerificationRequestT) Validate() *[]string {
	var errs []string

	if vr.Email == "" {
		errs = append(errs, "Email can't be empty!")
	}

	if len(errs) == 0 {
		return nil
	}
	return &errs
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

/*Validate validates sign in request*/
func (sur *SignUpRequestT) Validate() *[]string {
	var errs []string

	// Validate email and password
	signInErrs := sur.SignInRequestT.Validate()
	if signInErrs != nil {
		errs = *signInErrs
	}

	if sur.FirstName == "" {
		errs = append(errs, "First name can't be empty!")
	}

	if sur.LastName == "" {
		errs = append(errs, "Last name can't be empty!")
	}

	if len(errs) == 0 {
		return nil
	}
	return &errs
}

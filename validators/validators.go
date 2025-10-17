package validators

type AuthFormValidator struct {
	Emailempty    bool
	Passwordempty bool
	Confirmbad    bool
	Email         string
	Exists        bool
}

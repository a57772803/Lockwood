package models

type PasswordEntry struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Note     string `json:"note,omitempty"`
}

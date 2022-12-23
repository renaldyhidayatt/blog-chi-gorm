package request

import "errors"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (e *LoginRequest) Validate() error {
	if e.Username == "" {
		return errors.New("username tidak boleh kosong")
	} else if e.Password == "" {
		return errors.New("password tidak boleh kosong")
	} else if len(e.Password) < 6 || len(e.Password) > 18 {
		return errors.New("password minimal 6 karakter dan maksimal 18 karakter")
	}

	return nil

}

type ForgetPasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (e *ForgetPasswordRequest) Validate() error {
	if e.OldPassword == "" {
		return errors.New("password lama tidak boleh kosong")
	} else if len(e.OldPassword) < 6 || len(e.OldPassword) > 18 {
		return errors.New("password lama minimal 6 karakter dan maksimal 18 karakter")
	} else if e.NewPassword == "" {
		return errors.New("password baru tidak boleh kosong")
	} else if len(e.NewPassword) < 6 || len(e.NewPassword) > 18 {
		return errors.New("password baru minimal 6 karakter dan maksimal 18 karakter")
	}

	return nil
}

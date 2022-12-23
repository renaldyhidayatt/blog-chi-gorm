package request

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type UserRequest struct {
	ID        int64  `json:"id_user" gorm:"column:id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	NoTelp    string `json:"no_telp"`
	Photo     string `json:"photo"`
	Status    bool   `json:"status"`
	IdRole    int64  `json:"id_role" gorm:"column:id_role"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
}

func (e *UserRequest) TableName() string {
	return "tb_users"
}

func (e *UserRequest) Prepare() {
	e.Username = strings.ToLower(strings.ReplaceAll(e.Username, " ", ""))
	e.Password = strings.ReplaceAll(e.Password, " ", "")
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *UserRequest) Validate() error {
	matched, _ := regexp.MatchString(`[!@#~$%^&*()+|_-]{1}`, e.Username)

	if e.FirstName == "" {
		return errors.New("firstname tidak boleh kosong")
	} else if e.Username == "" {
		return errors.New("username tidak boleh kosong")
	} else if matched {
		return errors.New("username tidak boleh memiliki simbol")
	} else if len(e.Username) < 4 || len(e.Username) > 16 {
		return errors.New("username minimal 4 karakter dan maksimal 16 karakter")
	} else if e.Password == "" {
		return errors.New("password tidak boleh kosong")
	} else if len(e.Password) < 6 || len(e.Password) > 18 {
		return errors.New("password minimal 6 karakter dan maksimal 18 karakter")
	} else if e.Email == "" {
		return errors.New("email tidak boleh kosong")
	} else if err := checkmail.ValidateFormat(e.Email); err != nil {
		return errors.New("format email tidak sesuai")
	}

	return nil
}

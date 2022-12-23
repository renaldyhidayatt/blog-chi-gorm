package request

import (
	"errors"
	"time"

	"github.com/gosimple/slug"
)

type MenuRequest struct {
	ID       int64  `json:"id_menu" gorm:"column:id_menu"`
	NamaMenu string `json:"nama_menu"`
	Slug     string `json:"slug"`
	Icon     string `json:"icon"`
	Path     string `json:"path"`
	Status   bool   `json:"status"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *MenuRequest) TableName() string {
	return "tb_menus"
}

func (e *MenuRequest) Prepare() {
	e.Slug = slug.Make(e.NamaMenu)
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *MenuRequest) Validate() error {
	if e.NamaMenu == "" {
		return errors.New("nama Menu tidak boleh kosong")
	}

	if e.Icon == "" {
		return errors.New("icon tidak boleh kosong")
	}

	if e.Path == "" {
		return errors.New("path tidak boleh kosong")
	}

	return nil
}

package models

import (
	"time"
)

type (
	User struct {
		BaseModel
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password" valid:"password"`
	}

	UserPaginationResponse struct {
		Meta PaginationResponse `json:"meta"`
		Data []User             `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	UserFilterable struct {
		Name     string `json:"name"`
		Username string `json:"username"`
	}
)

var (
	_page = 1
	_rp   = 25
)

// Callback before update user
func (m *User) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create user
func (m *User) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func (m *User) Create() (*User, error) {
	var err error
	err = Create(&m)
	return m, err
}

// Update
func (m *User) Update() error {
	var err error
	err = Save(&m)
	return err
}

// Delete
func (m *User) Delete() error {
	var err error
	err = Delete(&m)
	return err
}

// FindByID
func (m *User) FindByID(id int) (*User, error) {
	var (
		err error
	)
	err = FindOneByID(&m, id)
	return m, err
}

// GetList
func (m *User) GetList(page int, rp int, orderby string, sort string, filters interface{}) (interface{}, error) {
	var (
		users []User
		err   error
	)

	resp, err := FindAllWithPage(&users, page, rp, orderby, sort, filters)
	return resp, err
}

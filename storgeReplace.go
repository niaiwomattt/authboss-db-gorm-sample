package main

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/volatiletech/authboss"
)


// DBStorer stores users in memory
type DBStorer struct {
	DB     *gorm.DB
	Users  map[string]User
	Tokens map[string][]string
}

// NewDBStorer constructor
func NewDBStorer() *DBStorer {
	db := openDb()
	initCreate()

	return &DBStorer{
		DB:     db,
		Users:  make(map[string]User),
		Tokens: make(map[string][]string),
	}
}

// Save the user
func (m DBStorer) Save(ctx context.Context, user authboss.User) error {
	u := user.(*User)
	m.Users[u.Email] = *u

	debugln("Saved user:", u.Name)
	// 保存到数据库
	if err := m.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

// Load the user
func (m DBStorer) Load(ctx context.Context, key string) (user authboss.User, err error) {

	// 从数据库加载
	var u User
	if err := m.DB.Where(&User{Email: key}).First(&u).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &u, nil
}

// New user creation
func (m DBStorer) New(ctx context.Context) authboss.User {
	return &User{}
}

// Create the user
func (m DBStorer) Create(ctx context.Context, user authboss.User) error {
	u := user.(*User)
	db.Create(u)
	debugln("Created new user:", u.Name)
	return nil
}

// LoadByConfirmSelector looks a user up by confirmation token
func (m DBStorer) LoadByConfirmSelector(ctx context.Context, selector string) (user authboss.ConfirmableUser, err error) {

	var u  User
	db.Where(&User{ConfirmSelector: selector}).First(&u)
	if u.ConfirmSelector != "" {
		return &u, nil
	}
	return nil, authboss.ErrUserNotFound
}

// LoadByRecoverSelector looks a user up by confirmation selector
func (m DBStorer) LoadByRecoverSelector(ctx context.Context, selector string) (user authboss.RecoverableUser, err error) {

	var u  User
	db.Where(&User{RecoverSelector: selector}).First(&u)
	if u.ConfirmSelector != "" {
		return &u, nil
	}
	return nil, authboss.ErrUserNotFound
}
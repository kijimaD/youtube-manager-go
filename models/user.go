package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key"`
	UID       string     `json:"-"` // 内部でのみ使用する情報であり、レスポンスには含めない。`そのためjson:"-"`を宣言して、レスポンスには含めない
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`

	Favorites []Favorite // リレーション定義。usersとfavoritesは1対Nのリレーション
}

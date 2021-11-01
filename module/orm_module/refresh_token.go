package orm_module

import "time"

type RefreshTokens struct {
	Id        string          `xorm:"pk notnull char(64)"` // OAuth Access Token
	CreatedAt time.Time       `xorm:"notnull"`             // OAuth Token Created Time
	UpdatedAt time.Time       `xorm:"notnull"`             // OAuth Token Updated Time
	ExpiredAt time.Time       `xorm:"notnull"`             // OAuth Token Expired Time
	User      string          `xorm:"notnull char(32)"`    // OAuth Token UserID
	Client    string          `xorm:"notnull char(32)"`    // OAuth Token ClientID
}

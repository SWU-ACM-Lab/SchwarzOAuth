package orm_module

import "time"

type OAuthUsers struct {
	Id        string    `xorm:"pk notnull unique char(32)"` // OAuth User ID
	CreatedAt time.Time `xorm:"notnull"`                    // OAuth User Created Time
	UpdatedAt time.Time `xorm:"notnull"`                    // OAuth User Updated Time
	RoleId    string    `xorm:"notnull char(32)"`           // OAuth User Role ID -> OAuthRoles: Id
	Name      string    `xorm:"notnull varchar(15)"`        // OAuth User Name
	Password  string    `xorm:"notnull char(64)"`           // OAuth User Password
	Email     string    `xorm:"notnull unique varchar(255)"`        // OAuth User Email
	Avatar    string    `xorm:"varchar(1023)"`              // OAuth User Avatar URI
	Status    int       `xorm:"notnull char(32)"`           // OAuth User Status
	Verified  bool      `xorm:"notnull"`                    // OAuth Verify Status
	SubId     string    `xorm:"varchar(255)"`               // OAuth User SubId
	RealName  string    `xorm:"varchar(15)"`                // OAuth User Real Name
}

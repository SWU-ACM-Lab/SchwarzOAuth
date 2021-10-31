package orm_module

import "time"

type OAuthClients struct {
	Id          string    `xorm:"pk notnull unique char(32)"`  // OAuth Client ID
	CreatedAt   time.Time `xorm:"notnull"`                     // OAuth Client Created Time
	UpdatedAt   time.Time `xorm:"notnull"`                     // OAuth Client Updated Time
	Name        string    `xorm:"notnull varchar(15)"`         // OAuth Client Name
	Description string    `xorm:"notnull varchar(255)"`        // OAuth Client Description
	AdminEmail  string    `xorm:"notnull unique varchar(255)"` // OAuth Client Administrator Email
	Secret      string    `xorm:"notnull char(32)"`            // OAuth Client Secret
	RedirectUri string    `xorm:"varchar(255)"`                // OAuth Client Redirect URI
	Avatar      string    `xorm:"varchar(1023)"`               // OAuth Client Avatar URI
	Scope       string    `xorm:"notnull char(32)"`            // OAuth Client Scope
	Status      int       `xorm:"notnull char(32)"`            // OAuth Client Status
	Verified    bool      `xorm:"notnull"`                     // OAuth Client Verify Status
}

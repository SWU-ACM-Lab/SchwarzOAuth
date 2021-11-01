package orm_module

type OAuthRoles struct {
	Id string `xorm:"pk notnull char(32)"` // OAuth Role ID
	Tag string `xorm:"notnull varchar(15)"` // OAuth Role Tag
	Rights map[string]bool `xorm:"notnull json"` // OAuth Role Mapping Rights
}
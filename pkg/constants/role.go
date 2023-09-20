package constants

type Role uint

const (
	SuperAdmin Role = iota
	Admin
	User
	Guest
)

func (r Role) String() string {
	switch r {
	case SuperAdmin:
		return "超级管理员"
	case Admin:
		return "管理员"
	case User:
		return "用户"
	case Guest:
		return "游客"
	}
	return ""
}

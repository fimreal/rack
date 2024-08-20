package models

type DBLayer interface {
	// 用户信息查询（精确查询）
	CheckUserExistence(username, email, phoneNumber string) (string, error)
	FindUser(query string, args ...interface{}) (*User, error)    // 查询返回第一匹配用户
	FindUsers(query string, args ...interface{}) ([]*User, error) // 查询返回所有匹配用户
	FindUsersByNickname(nickname string) ([]*User, error)         // 按照昵称查询用户
	FindUsersByRole(role int) ([]*User, error)                    // 按照角色查询用户
	FindUsersByRoleAndStatus(role int, status int) ([]*User, error)
	FindUsersWithPagination(limit, offset int) ([]*User, error)
	// 获取用户信息
	GetUser(user interface{}) (*User, error)
	GetUserByID(id uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByGithubID(githubID string) (*User, error)
	GetUserByGoogleID(googleID string) (*User, error)
	GetUserByWechatID(wechatID string) (*User, error)
	GetUserByQQID(qqID string) (*User, error)
	// 用户管理
	CreateUser(user *User) error
	UpdateUser(user *User) error
	UpdateUserStatus(id uint, status int) error
	UpdateUserPassword(id uint, password string) error
	UpdateUserNickname(id uint, nickname string) error
	UpdateUserEmail(id uint, email string) error
	UpdateUserRole(id uint, role int) error
	DeleteUser(id uint) error
}

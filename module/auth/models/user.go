package models

import (
	"database/sql"
	"fmt"

	"github.com/fimreal/rack/pkg/utils"
	"gorm.io/gorm"
)

// 定义角色常量
const (
	RoleSuperAdmin  = iota // 超级管理员
	RoleAdmin              // 管理员
	RoleContributor        // 贡献者
	RoleGuest              // 游客
	RoleUser               // 普通用户
	RoleSubscriber         // 订阅者
	RoleBanned             // 封禁用户
)

// User 表结构
type User struct {
	gorm.Model
	Username    string         `gorm:"column:username;uniqueIndex" validate:"required,min=1,max=50" json:"username"`       // 登录的用户名
	Password    string         `gorm:"column:password;" validate:"min=6,max=32" json:"-"`                                  // 不返回给客户端
	Email       sql.NullString `gorm:"column:email;uniqueIndex" validate:"email" json:"email,omitempty"`                   // 邮箱
	PhoneNumber sql.NullString `gorm:"column:phone_number;unique" validate:"numeric,len=11" json:"phone_number,omitempty"` // 11 位手机号码
	Nickname    string         `gorm:"column:nickname" json:"nickname,omitempty"`                                          // 用户昵称
	Avatar      string         `gorm:"column:avatar" json:"avatar,omitempty"`                                              // 用户头像 url
	Role        int            `gorm:"column:role;default:3" json:"role"`                                                  // 默认角色为游客
	Status      int            `gorm:"column:status;default:1" json:"status"`                                              // 默认状态为启用
}

// IsUsernameExists 检查用户名是否已存在
func (o *ORM) IsUsernameExists(username string) (bool, error) {
	var count int64
	if err := o.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err // 返回查询错误
	}
	return count > 0, nil // 返回是否存在
}

// IsEmailExists 检查邮箱是否已存在
func (o *ORM) IsEmailExists(email string) (bool, error) {
	if email == "" {
		return false, nil
	}
	var count int64
	if err := o.Model(&User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// IsPhoneNumberExists 检查手机号是否已存在
func (o *ORM) IsPhoneNumberExists(phone string) (bool, error) {
	if phone == "" {
		return false, nil
	}
	var count int64
	if err := o.Model(&User{}).Where("phone_number = ?", phone).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckUserExistence 检查用户是否存在（包括用户名、邮箱和手机号）
// 返回已占用的字段，如果不存在则返回空字符串
func (o *ORM) CheckUserExistence(username, email, phoneNumber string) (string, error) {
	// 检查用户名是否已存在
	exists, err := o.IsUsernameExists(username)
	if err != nil {
		return "", err
	}
	if exists {
		return "username", nil // 用户名已被占用
	}

	// 检查邮箱是否已存在
	exists, err = o.IsEmailExists(email)
	if err != nil {
		return "", err
	}
	if exists {
		return "email", nil // 邮箱已被占用
	}

	// 检查手机号是否已存在
	exists, err = o.IsPhoneNumberExists(phoneNumber)
	if err != nil {
		return "", err
	}
	if exists {
		return "phone_number", nil // 手机号已被占用
	}

	return "", nil // 所有字段均未被占用
}

// 获取当前所有用户数量
func (o *ORM) GetUserCount() (int64, error) {
	var count int64
	if err := o.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// FindUser 查询用户，可以传入条件和参数
func (o *ORM) FindUser(query string, args ...interface{}) (*User, error) {
	var u User
	if err := o.Where(query, args...).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByID 根据 ID 查询用户
func (o *ORM) GetUserByID(id uint) (*User, error) {
	return o.FindUser("id = ?", id)
}

// GetUserByUsername 根据用户名查询用户
func (o *ORM) GetUserByUsername(username string) (*User, error) {
	return o.FindUser("username = ?", username)
}

// GetUserByEmail 根据邮箱查询用户
func (o *ORM) GetUserByEmail(email string) (*User, error) {
	return o.FindUser("email = ?", email)
}

// GetUserByGithubID 根据 Github ID 查询用户
func (o *ORM) GetUserByGithubID(githubID string) (*User, error) {
	return o.FindUser("github_id = ?", githubID)
}

// GetUserByGoogleID 根据 Google ID 查询用户
func (o *ORM) GetUserByGoogleID(googleID string) (*User, error) {
	return o.FindUser("google_id = ?", googleID)
}

// GetUserByWechatID 根据微信 ID 查询用户
func (o *ORM) GetUserByWechatID(wechatID string) (*User, error) {
	return o.FindUser("wechat_id = ?", wechatID)
}

// GetUserByQQID 根据 QQ ID 查询用户
func (o *ORM) GetUserByQQID(qqID string) (*User, error) {
	return o.FindUser("qq_id = ?", qqID)
}

// GetUser 根据 ID、用户名或邮箱查询用户信息
func (o *ORM) GetUser(account interface{}) (*User, error) {
	switch t := account.(type) {
	case uint: // 如果传入的是 ID
		return o.GetUserByID(t)
	case string: // 如果传入的是用户名或邮箱
		if utils.IsEmail(t) {
			return o.GetUserByEmail(t)
		}
		return o.GetUserByUsername(t)
	default:
		return nil, fmt.Errorf("unsupported account type: %T", t)
	}
}

// FindUsers 查询符合条件的多个用户
func (o *ORM) FindUsers(query string, args ...interface{}) ([]*User, error) {
	var users []*User
	if err := o.Where(query, args...).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindUsersWithPagination 分页查询用户
func (o *ORM) FindUsersWithPagination(limit, offset int) ([]*User, error) {
	var users []*User
	if err := o.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID 根据 Nickname 查询用户
func (o *ORM) FindUsersByNickname(nickname string) ([]*User, error) {
	return o.FindUsers("nickname = ?", nickname)
}

// GetUserByRole 根据角色查询用户
func (o *ORM) FindUsersByRole(role int) ([]*User, error) {
	return o.FindUsers("role = ?", role)
}

// FindUsersByRoleAndStatus 根据角色和状态查询用户
func (o *ORM) FindUsersByRoleAndStatus(role int, status int) ([]*User, error) {
	return o.FindUsers("role = ? AND status = ?", role, status)
}

// CreateUser 创建新的用户
func (o *ORM) CreateUser(user *User) error {
	return o.Create(user).Error
}

// UpdateUser 更新用户信息
func (o *ORM) UpdateUser(user *User) error {
	return o.Save(user).Error
}

// DeleteUser 根据 ID 删除用户
func (o *ORM) DeleteUser(id uint) error {
	return o.Delete(&User{}, id).Error
}

// UpdateUserStatus 根据 ID 更新用户状态
func (o *ORM) UpdateUserStatus(id uint, status int) error {
	return o.Model(&User{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateUserPassword 根据 ID 更新用户密码
func (o *ORM) UpdateUserPassword(id uint, password string) error {
	return o.Model(&User{}).Where("id = ?", id).Update("password", password).Error
}

// UpdateUserNickname 根据 ID 更新用户昵称
func (o *ORM) UpdateUserNickname(id uint, nickname string) error {
	return o.Model(&User{}).Where("id = ?", id).Update("nickname", nickname).Error
}

// UpdateUserEmail 根据 ID 更新用户邮箱
func (o *ORM) UpdateUserEmail(id uint, email string) error {
	return o.Model(&User{}).Where("id = ?", id).Update("email", email).Error
}

// UpdateUserRole 根据 ID 更新用户角色
// id 对应角色
// 1:超级管理员
// 2: 管理员
// 3：贡献者
// 4: 游客（默认角色）
// 5: 普通用户
// 6: 订阅者
// 7: 封禁用户
func (o *ORM) UpdateUserRole(id uint, role int) error {
	return o.Model(&User{}).Where("id = ?", id).Update("role", role).Error
}

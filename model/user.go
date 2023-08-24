package model

import (
	"time"

	"go-web-example/conf"
	"go-web-example/errorcode"
)

// User return user info
type User struct {
	ID         int       `json:"id" gorm:"column:id; PRIMARY_KEY; AUTO_INCREMENT"`
	UserPhone  string    `json:"user_phone" gorm:"column:user_phone; type:char(11); default:''; unique"`
	UserPwd    string    `json:"user_pwd" gorm:"column:user_pwd; type:varchar(255); default:''"`
	UserName   string    `json:"user_name" gorm:"column:user_name;type:varchar(20)"`
	NickName   string    `json:"nick_name" gorm:"column:nick_name;type:varchar(20)"`
	ISUse      int       `json:"is_use" gorm:"column:is_use; type:tinyint(2); default:1"`
	Status     int       `json:"status" gorm:"column:status; type:tinyint(2); default:1"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time; type:int(11); not null"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time; type:int(11); not null"`
}

// TableName return table name
func (User) TableName() string {
	return "user"
}

// Create add a new row
func (u *User) Create() error {
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()

	db := conf.MysqlClient.Create(u)

	return db.Error
}

// Update update a new row
func (u *User) Update(params map[string]interface{}) error {
	params["update_time"] = time.Now()

	db := conf.MysqlClient.Model(&User{}).Where("id = ?", u.ID).Updates(params)

	return db.Error
}

// UserSwitch 开启或停止用户使用
func UserSwitch(ID, isUse int) error {
	db := conf.MysqlClient.Model(&User{}).Where("id = ?", ID).Updates(
		map[string]interface{}{
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
			"is_use":      isUse,
		})

	return db.Error
}

// UserDeleteByID 删除用户
func UserDeleteByID(ID int) error {
	db := conf.MysqlClient.Model(&User{}).Where("id = ?", ID).Updates(
		map[string]interface{}{
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
			"status":      ModelsDelete,
		})

	return db.Error
}

// UpdatePassByID 修改用户密码
func UpdatePassByID(ID int, pass string) error {
	db := conf.MysqlClient.Model(&User{}).Where("id = ?", ID).Updates(
		map[string]interface{}{
			"user_pwd":    pass,
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
		})

	return db.Error
}

// VerifyUserID 校验用户是否存在
func VerifyUserID(userID int) (bool, error) {
	var res int64
	db := conf.MysqlClient.Model(&User{}).Where("id = ? and status = ?", userID, ModelsActive).Count(&res)

	if db.Error != nil {
		return false, db.Error
	}

	if res == 0 {
		return false, errorcode.New(errorcode.ErrBizRecordNotFound, errorcode.RecordNotFoundMsg, nil)
	}

	return true, nil
}

// VerifyUserPhone 检测用户电话号是否存在
func VerifyUserPhone(phone string) (bool, error) {
	var res User
	db := conf.MysqlClient.Model(&User{}).Select("id").Where("user_phone = ? and status = ?", phone, ModelsActive).Find(&res)

	if db.Error != nil {
		return false, db.Error
	}

	if db.RowsAffected == 0 {
		return true, nil
	}

	return false, errorcode.New(errorcode.ErrBizRecordDuplicate, errorcode.RecordDuplicateMsg, nil)

}

// GetInfoByPhone return user info by phone
func GetInfoByPhone(phone string) (User, error) {
	var res User
	db := conf.MysqlClient.Where("user_phone = ? and status = ?", phone, ModelsActive).Find(&res)

	//判断是否存在
	if db.RowsAffected == 0 {
		return res, errorcode.New(errorcode.ErrBizRecordNotFound, "用户不存在", nil)
	}
	return res, db.Error
}

// GetUserInfoByParams return user info
func GetUserInfoByParams(params interface{}, limit, offset int) ([]User, error) {
	var res []User

	db := conf.MysqlClient.Model(&User{}).Where(params).Limit(limit).Offset(offset).Find(&res)

	return res, db.Error
}

// GetUserCountByParams return user count
func GetUserCountByParams(params interface{}) (int64, error) {
	var res int64

	db := conf.MysqlClient.Model(&User{}).Where(params).Count(&res)

	return res, db.Error
}

// GetUserInfoByID return user info
func GetUserInfoByID(ID int) (User, error) {
	var res User
	db := conf.MysqlClient.Model(&User{}).Where("id = ? and status = ?", ID, ModelsActive).Find(&res)

	return res, db.Error
}

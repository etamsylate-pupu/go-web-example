package systembiz

import (
	"strconv"

	"go-web-example/model"
	"go-web-example/request/authrequest"
	"go-web-example/utils"
)

// UserAdd add new  user
func UserAdd(params authrequest.RegisterInputParams) error {
	userModel := &model.User{}

	userModel.UserName = params.UserName
	userModel.NickName = params.NickName
	userModel.UserPhone = params.Phone
	userModel.UserPwd = utils.Sha1([]byte(params.Password))

	if err := userModel.Create(); err != nil {
		return err
	}

	return nil
}

// UserEdit edit a row user
func UserEdit(params map[string]string) error {
	userModel := &model.User{}
	userModel.ID, _ = strconv.Atoi(params["user_id"])

	updateParam := map[string]interface{}{}
	updateParam["nick_name"] = params["nick_name"]

	if err := userModel.Update(updateParam); err != nil {
		return err
	}

	return nil
}

// UserDelete delete a user
func UserDelete(userID int) error {

	if err := model.UserDeleteByID(userID); err != nil {
		return err
	}

	return nil
}

// UserInfo return user info
func UserInfo(userID int) (interface{}, error) {
	res, err := model.GetUserInfoByID(userID)
	if err != nil {
		return nil, err
	}

	info := map[string]interface{}{
		"user_name":  res.UserName,
		"user_phone": res.UserPhone,
	}

	return info, nil
}

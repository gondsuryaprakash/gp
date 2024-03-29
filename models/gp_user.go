package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/gondsuryaprakash/gondpariwar/logger"
)

type GpUser struct {
	Id         int    `orm:"column(id);auto" json:"id"`
	Name       string `orm:"column(name);null" json:"name"`
	Password   string `orm:"column(password);null" json:"-"`
	Email      string `orm:"column(email);null" json:"email"`
	Mobile     string `orm:"column(mobile);size(10);null;" json:"mobile"`
	Gender     string `orm:"column(gender);null;" json:"gender"`
	Age        string `orm:"column(age); null;" json:"age"`
	Religion   string `orm:"column(religion); null" json:"religion"`
	Dob        string `orm:"column(dob); null" json:"dob"`
	FatherName string `orm:"column(fathername); null" json:"fathername"`
	MotherName string `orm:"column(mothername); null" json:"mothername"`
}

func (u *GpUser) TableName() string {
	return "gp_user"
}

func init() {
	funcName := "models.init()"
	orm.RegisterModel(new(GpUser))
	logger.I(funcName)
}

// Adduser function insert user in database.
func AddUser(user *GpUser) (err error) {
	funcName := "models.AddUser"
	logger.I(funcName)
	o := orm.NewOrm()
	if _, err := o.Insert(user); err != nil {
		logger.E("Err", err)
		return err
	}
	return nil
}

// Get User ById and return User
func GetUserById(userId int) (v *GpUser, err error) {
	funcName := "models.GetUserById"
	logger.I(funcName)
	o := orm.NewOrm()
	v = &GpUser{Id: userId}
	if err := o.Read(v); err != nil {
		if err == orm.ErrMissPK {
			logger.D("No primary key found.")
		} else if err == orm.ErrNoRows {
			logger.D("No Row Found.")
		}
		logger.D("GetUserById", err)
		return nil, err
	}
	return v, nil
}

// Get User ById and return User
func GetUserByEmailId(email string) (v *GpUser, err error) {
	funcName := "models.GetUserByEmailId"
	logger.I(funcName)
	o := orm.NewOrm()
	v = &GpUser{Email: email}
	if err := o.Read(v, "Email"); err != nil {
		if err == orm.ErrMissPK {
			logger.D("No primary key found.")
		} else if err == orm.ErrNoRows {
			logger.D("No Row Found.")
		}
		logger.D("GetUserById", err)
		return nil, err
	}
	logger.I(funcName, v)
	return v, nil
}

func IsUserExistByEmail(email string) (bool, error) {
	funcName := "models.isUserExistByEmail"
	logger.I(funcName)
	o := orm.NewOrm()
	v := &GpUser{}
	err := o.QueryTable(new(GpUser)).Filter("email", email).One(v)
	if err != nil && err != orm.ErrNoRows {
		return true, nil
	}
	if err == orm.ErrNoRows {
		return false, err
	}
	return true, nil
}

func UpdateUserById(m *GpUser) (err error) {

	funcName := "models.UpdateUserById"
	logger.I(funcName)
	o := orm.NewOrm()
	v := &GpUser{Id: m.Id}
	if err = o.Read(v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			logger.D("User updated with number of columns", num)
		}
	}
	return
}

func UpdateUserByIdWithColumns(m *GpUser, colums ...string) (err error) {

	funcName := "models.UpdateUserByIdWithColumns"
	logger.I(funcName)
	o := orm.NewOrm()
	v := &GpUser{Id: m.Id}
	if err = o.Read(v); err == nil {
		// var num int64
		// if num, err = o.Update(m); err == nil {
		// 	logger.D("User updated with number of columns", num)
		// }
		if err = UpdateRowByColumns(m, colums...); err == nil {
			logger.E(err)
			return
		}
	}
	return
}
func DeleteUserById(userId int) (err error) {
	funcName := "models.DeleteUserById"
	logger.I(funcName)

	o := orm.NewOrm()
	v := &GpUser{Id: userId}
	if err = o.Read(v); err == nil {
		var num int64
		if num, err = o.Delete(&GpUser{Id: v.Id}); err == nil {
			logger.D("User Deleted with records", num)
		}
	}

	return
}

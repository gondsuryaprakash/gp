package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/gondsuryaprakash/gondpariwar/logger"
)

func UpdateRowByColumns(m interface{}, columns ...string) error {
	funcName := "models.UpdateRowByColumns"
	logger.I(funcName)
	o := orm.NewOrm()
	return UpdateRowByColumnsByOrm(m, o, columns...)
}

func UpdateRowByColumnsByOrm(m interface{}, o orm.Ormer, columns ...string) error {
	funcName := "models.UpdateRowByColumnsByOrm"
	logger.I(funcName)
	if _, err := o.Update(m, columns...); err != nil {
		logger.E(err)
		return err
	}
	return nil
}

package gazo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"kanri/orm"
	"kanri/orm/model"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
)

func jOutput(dst io.Writer, src interface{}) error {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return errs.Wrap(err)
	}
	if _, err := io.Copy(os.Stdout, buf); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func CreateGazo(prod_cd string, customer_cd string, customer_name string, comment string, img_name string) exitcode.ExitCode {
	// create gorm.DB instance for PostgreSQL service
	gormCtx, err := orm.NewGORM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitcode.Abnormal
	}
	defer gormCtx.Close()

	if img_name == "undefined" {
		img_name = ""
	}

	data := &model.Gazo{
		Prod_cd:       prod_cd,
		Customer_cd:   customer_cd,
		Customer_name: customer_name,
		Comment:       comment,
		Img_name:      img_name,
	}

	tx := gormCtx.GetDb().WithContext(context.TODO()).Create(data)
	if tx.Error != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

func FindGazo(id string, prod_cd string, customer_cd string, customer_name string) ([]map[string]interface{}, exitcode.ExitCode) {
	// create gorm.DB instance for PostgreSQL service
	gormCtx, err := orm.NewGORM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, exitcode.Abnormal
	}

	//data := []model.Gazo{}
	var data []map[string]interface{}
	gazo := model.Gazo{}

	chain := gormCtx.GetDb().Model(&gazo).WithContext(context.TODO()).Where("")

	if id != "" {
		chain.Where("id = ?", id)
	}

	if prod_cd != "" {
		chain.Where("prod_cd = ?", prod_cd)
	}
	if customer_cd != "" {
		chain.Where("customer_cd = ?", customer_cd)
	}
	if customer_name != "" {
		chain.Where("customer_name LIKE ?", "%"+customer_name+"%")
	}

	tx := chain.Find(&data)

	if tx.Error != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return nil, exitcode.Abnormal
	}

	if err := jOutput(os.Stdout, data); err != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return nil, exitcode.Abnormal
	}

	return data, exitcode.Normal
}

func UpdateGazo(id string, prod_cd string, customer_cd string, customer_name string, comment string, img_name string) exitcode.ExitCode {
	// create gorm.DB instance for PostgreSQL service
	gormCtx, err := orm.NewGORM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitcode.Abnormal
	}
	defer gormCtx.Close()

	data := &model.Gazo{
		Prod_cd:       prod_cd,
		Customer_cd:   customer_cd,
		Customer_name: customer_name,
		Comment:       comment,
		Img_name:      img_name,
	}

	tx := gormCtx.GetDb().WithContext(context.TODO()).Where(id).Updates(&data)
	if tx.Error != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

func DeleteGazo(id string) exitcode.ExitCode {
	gormCtx, err := orm.NewGORM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitcode.Abnormal
	}
	defer gormCtx.Close()

	gazo := model.Gazo{}

	tx := gormCtx.GetDb().WithContext(context.TODO()).Where(id).Delete(&gazo)
	if tx.Error != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

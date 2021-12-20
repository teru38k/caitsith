package main

import (
	"context"
	"fmt"
	"kanri/orm"
	"kanri/orm/model"
	"os"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
)

func Run() exitcode.ExitCode {
	//create gorm.DB instace for PostgreSQL service
	gormCtx, err := orm.NewGORM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitcode.Abnormal
	}
	defer gormCtx.Close()

	//migration
	if err := gormCtx.GetDb().WithContext(context.TODO()).AutoMigrate(&model.Gazo{}); err != nil {
		gormCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}
	return exitcode.Normal
}

func main() {
	Run().Exit()
}

package repository

import (
	"encoding/json"
	"fmt"
	"investment/domain/po/custom"
	"investment/server/conf"
	"time"

	"github.com/8treenet/freedom"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getSatUnitTest() freedom.UnitTest {
	//创建单元测试工具
	unitTest := freedom.NewUnitTest()
	unitTest.InstallDB(func() interface{} {
		conf := conf.Get().DB
		db, err := gorm.Open(sqlite.Open("/Users/ys/work/my/investment/investment.db"), &gorm.Config{})
		if err != nil {
			freedom.Logger().Fatal(err.Error())
		}
		db.AutoMigrate(&custom.ClsDepthArticle{})

		sqlDB, err := db.DB()
		if err != nil {
			freedom.Logger().Fatal(err)
		}
		if err = sqlDB.Ping(); err != nil {
			freedom.Logger().Fatal(err)
		}

		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)
		return db
	})
	return unitTest
}
func jsonLog(data interface{}) {
	jdata, _ := json.MarshalIndent(data, "  ", "  ")
	fmt.Println(string(jdata))
}

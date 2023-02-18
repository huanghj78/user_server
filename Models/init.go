package Models

import (
	"database/sql"
	"fmt"
	"user_server/Config"
	"user_server/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	DBHelper *gorm.DB
	DBSetting *sql.DB
	err			error
)

type BaseModels struct {
	ID 		  uint 		`gorm:"primary_key" json:"id"`
	CreatedAt time.Time	`json:"create_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.Address,
		Config.Database.Port,
		Config.Database.Database)
	DBHelper, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		utils.Logger.Fatal("MySQL Connect Failed ", err)
	}

	if DBHelper.Error != nil {
		utils.Logger.Fatal("DB Error ", err)
	}

	DBSetting, err = DBHelper.DB()
	DBSetting.SetMaxIdleConns(10)
	DBSetting.SetMaxOpenConns(100)
	DBSetting.SetConnMaxLifetime(time.Hour)

	migrate()
}

func migrate() {
	DBHelper.AutoMigrate(&UserInfo{})
}

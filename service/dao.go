package service

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zzsds/micro-store/user-service/conf"
	"github.com/zzsds/micro-store/user-service/models"
)

type Dao struct {
	db   *gorm.DB
	user *models.User
}

var (
	instance *Dao
	Once     sync.Once
)

func NewDao() *Dao {
	Once.Do(func() {
		instance = &Dao{}
	})
	instance.connection()
	instance.AutoMigrate()
	return instance
}

func (dao *Dao) Db() *gorm.DB {
	return dao.db
}

func (dao *Dao) Close() {
	dao.Db().Close()
}

func (dao *Dao) connection() (err error) {
	dao.db, err = gorm.Open("mysql", dao.GetDbConnMsg())
	if err != nil {
		log.Fatalf("models.Connection err: %v", err)
		return err
	}
	dao.db.DB().SetMaxIdleConns(10)
	dao.db.DB().SetMaxOpenConns(100)
	dao.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8")
	dao.db.LogMode(conf.Conf.Db.Debug)
	return
}

func (dao *Dao) GetDbConnMsg() string {
	cfg := conf.Conf
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Name, cfg.Db.Charset)
}

func (dao *Dao) AutoMigrate() {
	// 数据迁移
	if err := dao.Db().AutoMigrate(dao.user).Error; err != nil {
		panic(fmt.Sprintf("sync create mysql table failed: %v", err))
	}
}

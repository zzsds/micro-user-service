package service

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/store"
	Sync "github.com/micro/go-micro/v2/sync"
	"github.com/micro/go-micro/v2/sync/lock"
	memoryLock "github.com/micro/go-micro/v2/sync/lock/memory"
	"github.com/zzsds/micro-user-service/conf"
	"github.com/zzsds/micro-user-service/models"
)

const (
	ID     = "id"
	CODE   = "code"
	MOBILE = "mobile"
)

var (
	instance *Dao
	Once     sync.Once
	MapData  = []string{ID, CODE, MOBILE}
)

type Options struct {
	Store store.Store
	Lock  lock.Lock
	Code  Sync.Map
}

type Option func(o *Options)

func NewOption(opts ...Option) *Options {
	var options = &Options{
		Store: store.DefaultStore,
		Lock:  memoryLock.NewLock(),
	}
	for _, o := range opts {
		o(options)
	}
	mapping := Sync.NewMap(Sync.WithStore(options.Store), Sync.WithLock(options.Lock))
	options.Code = mapping
	return options
}

// WithStore sets the store implementation option
func WithStore(s store.Store) Option {
	return func(o *Options) {
		o.Store = s
	}
}

// WithLock sets the locking implementation option
func WithLock(l lock.Lock) Option {
	return func(o *Options) {
		o.Lock = l
	}
}

type Dao struct {
	options *Options
	mu      sync.Mutex
	db      *gorm.DB
	user    *models.User
}

// NewDao ...
func NewDao(opts ...Option) *Dao {
	Once.Do(func() {
		instance = &Dao{
			options: NewOption(opts...),
		}
	})
	instance.connection()
	instance.AutoMigrate()
	// 初始化数据到map列表
	instance.initCodeList()
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

func (dao *Dao) initCodeList(users ...*models.User) error {
	if len(users) <= 0 {
		if dao.Db().Model(dao.user).Select("id, code").Find(&users).RecordNotFound() {
			return nil
		}
	}

	opt := dao.options
	for _, v := range users {
		opt.Code.Write(v.Code, v.ID)
	}
	return nil
}

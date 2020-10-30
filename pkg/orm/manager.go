/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/24 21:38
 */
package orm

import (
	"context"
	"errors"
	"fmt"

	"github.com/nopsky/project-demo/pkg/config"
	"github.com/nopsky/project-demo/pkg/logger"
	perrors "github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Manager struct {
	opts *Options
	dbs  map[string]*DB
}

func New() (*Manager, error) {

	dm := &Manager{
		opts: newOptions(),
		dbs:  make(map[string]*DB),
	}

	if err := config.UnmarshalKey("mysql", &dm.opts.Databases); err != nil {
		return nil, perrors.Wrap(err, "读取数据库配置")
	}

	logger.Debug(dm.opts)

	if _, ok := dm.opts.Databases["default"]; !ok {
		return nil, errors.New("db:default not found")
	}

	for name, o := range dm.opts.Databases {
		engine, err := gorm.Open(mysql.Open(o.Dns), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: NewLogger(),
		})
		if err != nil {
			return nil, err
		}

		logger.Infof("db:%s mode:%s", name, o.Mode)

		if o.Mode == "debug" {
			engine = engine.Debug()
		}

		db, err := engine.DB()
		db.SetMaxIdleConns(o.MaxIdleConns)
		db.SetMaxOpenConns(o.MaxOpenConns)

		dm.dbs[name] = &DB{
			name: name,
			orm:  engine,
		}
	}

	return dm, nil
}

func NewWithOptions(optFn ...OptionFn) (*Manager, error) {
	dm := &Manager{
		opts: newOptions(),
		dbs:  make(map[string]*DB),
	}

	for _, fn := range optFn {
		fn(dm.opts)
	}

	logger.Debug(dm.opts)

	for name, o := range dm.opts.Databases {
		engine, err := gorm.Open(mysql.Open(o.Dns), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: NewLogger(),
		})
		if err != nil {
			return nil, err
		}

		logger.Infof("db:%s mode:%s", name, o.Mode)

		if o.Mode == "debug" {
			engine = engine.Debug()
		}

		db, err := engine.DB()
		db.SetMaxIdleConns(o.MaxIdleConns)
		db.SetMaxOpenConns(o.MaxOpenConns)

		dm.dbs[name] = &DB{
			name: name,
			orm:  engine,
		}
	}

	return dm, nil
}

func (d Manager) GetDB(name string) (*DB, error) {
	if db, ok := d.dbs[name]; ok {
		return db, nil
	}
	return nil, fmt.Errorf("db:%s not found", name)
}

func (d Manager) NewOrm(ctx context.Context) (*Session, error) {
	db, ok := d.dbs["default"]
	if !ok {
		return nil, errors.New("db:`default` not found")
	}
	return db.NewOrm(ctx)
}

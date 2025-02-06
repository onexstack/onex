package registry

import (
	"sync"

	"gorm.io/gorm"
)

// 定义一个Registry类型用于管理模型注册  
type Registry struct {
	models []interface{}
}

var (
	globalRegistry *Registry
	once           sync.Once
)

// NewRegistry 创建并返回一个新的Registry实例  
func NewRegistry() *Registry {
	return &Registry{
		models: make([]interface{}, 0),
	}
}

func Register(model interface{}) {
	once.Do(func() {
		globalRegistry = NewRegistry()
	})
	globalRegistry.Register(model)
}

// Register 添加新的模型到Registry  
func (r *Registry) Register(model interface{}) {
	r.models = append(r.models, model)
}

func Migrate(db *gorm.DB) error {
	if globalRegistry == nil {
		return nil
	}

	return globalRegistry.Migrate(db)
}

// Migrate 执行所有注册模型的迁移
func (r *Registry) Migrate(db *gorm.DB) error {
	for _, model := range r.models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

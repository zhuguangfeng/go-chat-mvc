package mysqlx

import (
	"gorm.io/gorm"
	"reflect"
)

type Builder struct {
	DB *gorm.DB
}

func NewDaoBuilder(db *gorm.DB) *Builder {
	return &Builder{
		DB: db,
	}
}

func (b *Builder) WithPagination(pageNum, pageSize int) *Builder {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset, limit := (pageNum-1)*pageSize, pageSize
	b.DB = b.DB.Offset(offset).Limit(limit)
	return b
}

func (b *Builder) WithLike(key, val string) *Builder {
	if val != "" {
		b.DB = b.DB.Where("? like ?", key, "%"+key+"%")
	}
	return b
}

func (b *Builder) WithEqual(key string, val any) *Builder {
	if b.IsNull(val) {
		return b
	}

	b.DB = b.DB.Where(key+" = ?", val)
	return b
}

func (b *Builder) WithLte(key string, val any) *Builder {
	if b.IsNull(val) {
		return b
	}
	b.DB = b.DB.Where(key+" <= ?", val)
	return b
}

func (b *Builder) WithGte(key string, val any) *Builder {
	if b.IsNull(val) {
		return b
	}
	b.DB = b.DB.Where(key+" >= ?", val)
	return b
}

func (b *Builder) WithFind(res any) *Builder {
	b.DB = b.DB.Find(&res)
	return b
}

func (b *Builder) IsNull(val any) bool {
	// 获取 val 的反射值
	valValue := reflect.ValueOf(val)
	// 如果 val 是 nil 或者是类型的零值，则不添加 Where 条件
	return val == nil || valValue.Kind() == reflect.Ptr && valValue.IsNil() || valValue.IsZero()
}

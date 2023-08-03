// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gin-mall/internal/gen/model"
)

func newSpec(db *gorm.DB, opts ...gen.DOOption) spec {
	_spec := spec{}

	_spec.specDo.UseDB(db, opts...)
	_spec.specDo.UseModel(&model.Spec{})

	tableName := _spec.specDo.TableName()
	_spec.ALL = field.NewAsterisk(tableName)
	_spec.ID = field.NewInt(tableName, "id")
	_spec.CategoryID = field.NewInt(tableName, "category_id")
	_spec.GroupID = field.NewInt(tableName, "group_id")
	_spec.Name = field.NewString(tableName, "name")
	_spec.IsGenerate = field.NewBool(tableName, "is_generate")
	_spec.Searching = field.NewBool(tableName, "searching")
	_spec.CreatedAt = field.NewTime(tableName, "created_at")
	_spec.UpdatedAt = field.NewTime(tableName, "updated_at")

	_spec.fillFieldMap()

	return _spec
}

type spec struct {
	specDo specDo

	ALL        field.Asterisk
	ID         field.Int
	CategoryID field.Int    // 商品分类id
	GroupID    field.Int    // 分组id
	Name       field.String // 参数名称
	IsGenerate field.Bool   // 是否通用
	Searching  field.Bool   // 是否用于搜索
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (s spec) Table(newTableName string) *spec {
	s.specDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s spec) As(alias string) *spec {
	s.specDo.DO = *(s.specDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *spec) updateTableName(table string) *spec {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.CategoryID = field.NewInt(table, "category_id")
	s.GroupID = field.NewInt(table, "group_id")
	s.Name = field.NewString(table, "name")
	s.IsGenerate = field.NewBool(table, "is_generate")
	s.Searching = field.NewBool(table, "searching")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *spec) WithContext(ctx context.Context) ISpecDo { return s.specDo.WithContext(ctx) }

func (s spec) TableName() string { return s.specDo.TableName() }

func (s spec) Alias() string { return s.specDo.Alias() }

func (s spec) Columns(cols ...field.Expr) gen.Columns { return s.specDo.Columns(cols...) }

func (s *spec) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *spec) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["category_id"] = s.CategoryID
	s.fieldMap["group_id"] = s.GroupID
	s.fieldMap["name"] = s.Name
	s.fieldMap["is_generate"] = s.IsGenerate
	s.fieldMap["searching"] = s.Searching
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s spec) clone(db *gorm.DB) spec {
	s.specDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s spec) replaceDB(db *gorm.DB) spec {
	s.specDo.ReplaceDB(db)
	return s
}

type specDo struct{ gen.DO }

type ISpecDo interface {
	gen.SubQuery
	Debug() ISpecDo
	WithContext(ctx context.Context) ISpecDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISpecDo
	WriteDB() ISpecDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISpecDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISpecDo
	Not(conds ...gen.Condition) ISpecDo
	Or(conds ...gen.Condition) ISpecDo
	Select(conds ...field.Expr) ISpecDo
	Where(conds ...gen.Condition) ISpecDo
	Order(conds ...field.Expr) ISpecDo
	Distinct(cols ...field.Expr) ISpecDo
	Omit(cols ...field.Expr) ISpecDo
	Join(table schema.Tabler, on ...field.Expr) ISpecDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISpecDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISpecDo
	Group(cols ...field.Expr) ISpecDo
	Having(conds ...gen.Condition) ISpecDo
	Limit(limit int) ISpecDo
	Offset(offset int) ISpecDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISpecDo
	Unscoped() ISpecDo
	Create(values ...*model.Spec) error
	CreateInBatches(values []*model.Spec, batchSize int) error
	Save(values ...*model.Spec) error
	First() (*model.Spec, error)
	Take() (*model.Spec, error)
	Last() (*model.Spec, error)
	Find() ([]*model.Spec, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Spec, err error)
	FindInBatches(result *[]*model.Spec, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Spec) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISpecDo
	Assign(attrs ...field.AssignExpr) ISpecDo
	Joins(fields ...field.RelationField) ISpecDo
	Preload(fields ...field.RelationField) ISpecDo
	FirstOrInit() (*model.Spec, error)
	FirstOrCreate() (*model.Spec, error)
	FindByPage(offset int, limit int) (result []*model.Spec, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISpecDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s specDo) Debug() ISpecDo {
	return s.withDO(s.DO.Debug())
}

func (s specDo) WithContext(ctx context.Context) ISpecDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s specDo) ReadDB() ISpecDo {
	return s.Clauses(dbresolver.Read)
}

func (s specDo) WriteDB() ISpecDo {
	return s.Clauses(dbresolver.Write)
}

func (s specDo) Session(config *gorm.Session) ISpecDo {
	return s.withDO(s.DO.Session(config))
}

func (s specDo) Clauses(conds ...clause.Expression) ISpecDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s specDo) Returning(value interface{}, columns ...string) ISpecDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s specDo) Not(conds ...gen.Condition) ISpecDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s specDo) Or(conds ...gen.Condition) ISpecDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s specDo) Select(conds ...field.Expr) ISpecDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s specDo) Where(conds ...gen.Condition) ISpecDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s specDo) Order(conds ...field.Expr) ISpecDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s specDo) Distinct(cols ...field.Expr) ISpecDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s specDo) Omit(cols ...field.Expr) ISpecDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s specDo) Join(table schema.Tabler, on ...field.Expr) ISpecDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s specDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISpecDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s specDo) RightJoin(table schema.Tabler, on ...field.Expr) ISpecDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s specDo) Group(cols ...field.Expr) ISpecDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s specDo) Having(conds ...gen.Condition) ISpecDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s specDo) Limit(limit int) ISpecDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s specDo) Offset(offset int) ISpecDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s specDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISpecDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s specDo) Unscoped() ISpecDo {
	return s.withDO(s.DO.Unscoped())
}

func (s specDo) Create(values ...*model.Spec) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s specDo) CreateInBatches(values []*model.Spec, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s specDo) Save(values ...*model.Spec) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s specDo) First() (*model.Spec, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spec), nil
	}
}

func (s specDo) Take() (*model.Spec, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spec), nil
	}
}

func (s specDo) Last() (*model.Spec, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spec), nil
	}
}

func (s specDo) Find() ([]*model.Spec, error) {
	result, err := s.DO.Find()
	return result.([]*model.Spec), err
}

func (s specDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Spec, err error) {
	buf := make([]*model.Spec, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s specDo) FindInBatches(result *[]*model.Spec, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s specDo) Attrs(attrs ...field.AssignExpr) ISpecDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s specDo) Assign(attrs ...field.AssignExpr) ISpecDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s specDo) Joins(fields ...field.RelationField) ISpecDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s specDo) Preload(fields ...field.RelationField) ISpecDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s specDo) FirstOrInit() (*model.Spec, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spec), nil
	}
}

func (s specDo) FirstOrCreate() (*model.Spec, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spec), nil
	}
}

func (s specDo) FindByPage(offset int, limit int) (result []*model.Spec, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s specDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s specDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s specDo) Delete(models ...*model.Spec) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *specDo) withDO(do gen.Dao) *specDo {
	s.DO = *do.(*gen.DO)
	return s
}

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

	"apu/pkg/store/mysql/model"
)

func newWechatHeader(db *gorm.DB, opts ...gen.DOOption) wechatHeader {
	_wechatHeader := wechatHeader{}

	_wechatHeader.wechatHeaderDo.UseDB(db, opts...)
	_wechatHeader.wechatHeaderDo.UseModel(&model.WechatHeader{})

	tableName := _wechatHeader.wechatHeaderDo.TableName()
	_wechatHeader.ALL = field.NewAsterisk(tableName)
	_wechatHeader.ID = field.NewInt32(tableName, "id")
	_wechatHeader.Wxuin = field.NewString(tableName, "wxuin")
	_wechatHeader.Type = field.NewString(tableName, "type")
	_wechatHeader.Headers = field.NewString(tableName, "headers")
	_wechatHeader.Cookie = field.NewString(tableName, "cookie")
	_wechatHeader.Status = field.NewString(tableName, "status")
	_wechatHeader.CreateTime = field.NewTime(tableName, "create_time")

	_wechatHeader.fillFieldMap()

	return _wechatHeader
}

type wechatHeader struct {
	wechatHeaderDo

	ALL        field.Asterisk
	ID         field.Int32
	Wxuin      field.String
	Type       field.String
	Headers    field.String
	Cookie     field.String
	Status     field.String
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (w wechatHeader) Table(newTableName string) *wechatHeader {
	w.wechatHeaderDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wechatHeader) As(alias string) *wechatHeader {
	w.wechatHeaderDo.DO = *(w.wechatHeaderDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wechatHeader) updateTableName(table string) *wechatHeader {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.Wxuin = field.NewString(table, "wxuin")
	w.Type = field.NewString(table, "type")
	w.Headers = field.NewString(table, "headers")
	w.Cookie = field.NewString(table, "cookie")
	w.Status = field.NewString(table, "status")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *wechatHeader) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wechatHeader) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["wxuin"] = w.Wxuin
	w.fieldMap["type"] = w.Type
	w.fieldMap["headers"] = w.Headers
	w.fieldMap["cookie"] = w.Cookie
	w.fieldMap["status"] = w.Status
	w.fieldMap["create_time"] = w.CreateTime
}

func (w wechatHeader) clone(db *gorm.DB) wechatHeader {
	w.wechatHeaderDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wechatHeader) replaceDB(db *gorm.DB) wechatHeader {
	w.wechatHeaderDo.ReplaceDB(db)
	return w
}

type wechatHeaderDo struct{ gen.DO }

type IWechatHeaderDo interface {
	gen.SubQuery
	Debug() IWechatHeaderDo
	WithContext(ctx context.Context) IWechatHeaderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWechatHeaderDo
	WriteDB() IWechatHeaderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWechatHeaderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWechatHeaderDo
	Not(conds ...gen.Condition) IWechatHeaderDo
	Or(conds ...gen.Condition) IWechatHeaderDo
	Select(conds ...field.Expr) IWechatHeaderDo
	Where(conds ...gen.Condition) IWechatHeaderDo
	Order(conds ...field.Expr) IWechatHeaderDo
	Distinct(cols ...field.Expr) IWechatHeaderDo
	Omit(cols ...field.Expr) IWechatHeaderDo
	Join(table schema.Tabler, on ...field.Expr) IWechatHeaderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWechatHeaderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWechatHeaderDo
	Group(cols ...field.Expr) IWechatHeaderDo
	Having(conds ...gen.Condition) IWechatHeaderDo
	Limit(limit int) IWechatHeaderDo
	Offset(offset int) IWechatHeaderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWechatHeaderDo
	Unscoped() IWechatHeaderDo
	Create(values ...*model.WechatHeader) error
	CreateInBatches(values []*model.WechatHeader, batchSize int) error
	Save(values ...*model.WechatHeader) error
	First() (*model.WechatHeader, error)
	Take() (*model.WechatHeader, error)
	Last() (*model.WechatHeader, error)
	Find() ([]*model.WechatHeader, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WechatHeader, err error)
	FindInBatches(result *[]*model.WechatHeader, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WechatHeader) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWechatHeaderDo
	Assign(attrs ...field.AssignExpr) IWechatHeaderDo
	Joins(fields ...field.RelationField) IWechatHeaderDo
	Preload(fields ...field.RelationField) IWechatHeaderDo
	FirstOrInit() (*model.WechatHeader, error)
	FirstOrCreate() (*model.WechatHeader, error)
	FindByPage(offset int, limit int) (result []*model.WechatHeader, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWechatHeaderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wechatHeaderDo) Debug() IWechatHeaderDo {
	return w.withDO(w.DO.Debug())
}

func (w wechatHeaderDo) WithContext(ctx context.Context) IWechatHeaderDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wechatHeaderDo) ReadDB() IWechatHeaderDo {
	return w.Clauses(dbresolver.Read)
}

func (w wechatHeaderDo) WriteDB() IWechatHeaderDo {
	return w.Clauses(dbresolver.Write)
}

func (w wechatHeaderDo) Session(config *gorm.Session) IWechatHeaderDo {
	return w.withDO(w.DO.Session(config))
}

func (w wechatHeaderDo) Clauses(conds ...clause.Expression) IWechatHeaderDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wechatHeaderDo) Returning(value interface{}, columns ...string) IWechatHeaderDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wechatHeaderDo) Not(conds ...gen.Condition) IWechatHeaderDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wechatHeaderDo) Or(conds ...gen.Condition) IWechatHeaderDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wechatHeaderDo) Select(conds ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wechatHeaderDo) Where(conds ...gen.Condition) IWechatHeaderDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wechatHeaderDo) Order(conds ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wechatHeaderDo) Distinct(cols ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wechatHeaderDo) Omit(cols ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wechatHeaderDo) Join(table schema.Tabler, on ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wechatHeaderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wechatHeaderDo) RightJoin(table schema.Tabler, on ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wechatHeaderDo) Group(cols ...field.Expr) IWechatHeaderDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wechatHeaderDo) Having(conds ...gen.Condition) IWechatHeaderDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wechatHeaderDo) Limit(limit int) IWechatHeaderDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wechatHeaderDo) Offset(offset int) IWechatHeaderDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wechatHeaderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWechatHeaderDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wechatHeaderDo) Unscoped() IWechatHeaderDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wechatHeaderDo) Create(values ...*model.WechatHeader) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wechatHeaderDo) CreateInBatches(values []*model.WechatHeader, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wechatHeaderDo) Save(values ...*model.WechatHeader) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wechatHeaderDo) First() (*model.WechatHeader, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatHeader), nil
	}
}

func (w wechatHeaderDo) Take() (*model.WechatHeader, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatHeader), nil
	}
}

func (w wechatHeaderDo) Last() (*model.WechatHeader, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatHeader), nil
	}
}

func (w wechatHeaderDo) Find() ([]*model.WechatHeader, error) {
	result, err := w.DO.Find()
	return result.([]*model.WechatHeader), err
}

func (w wechatHeaderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WechatHeader, err error) {
	buf := make([]*model.WechatHeader, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wechatHeaderDo) FindInBatches(result *[]*model.WechatHeader, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wechatHeaderDo) Attrs(attrs ...field.AssignExpr) IWechatHeaderDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wechatHeaderDo) Assign(attrs ...field.AssignExpr) IWechatHeaderDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wechatHeaderDo) Joins(fields ...field.RelationField) IWechatHeaderDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wechatHeaderDo) Preload(fields ...field.RelationField) IWechatHeaderDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wechatHeaderDo) FirstOrInit() (*model.WechatHeader, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatHeader), nil
	}
}

func (w wechatHeaderDo) FirstOrCreate() (*model.WechatHeader, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatHeader), nil
	}
}

func (w wechatHeaderDo) FindByPage(offset int, limit int) (result []*model.WechatHeader, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wechatHeaderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wechatHeaderDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wechatHeaderDo) Delete(models ...*model.WechatHeader) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wechatHeaderDo) withDO(do gen.Dao) *wechatHeaderDo {
	w.DO = *do.(*gen.DO)
	return w
}
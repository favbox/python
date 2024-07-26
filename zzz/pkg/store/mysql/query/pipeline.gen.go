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

func newPipeline(db *gorm.DB, opts ...gen.DOOption) pipeline {
	_pipeline := pipeline{}

	_pipeline.pipelineDo.UseDB(db, opts...)
	_pipeline.pipelineDo.UseModel(&model.Pipeline{})

	tableName := _pipeline.pipelineDo.TableName()
	_pipeline.ALL = field.NewAsterisk(tableName)
	_pipeline.ID = field.NewInt64(tableName, "id")
	_pipeline.NoteID = field.NewUint64(tableName, "note_id")
	_pipeline.IsCounted = field.NewBool(tableName, "is_counted")
	_pipeline.IsDetailed = field.NewBool(tableName, "is_detailed")
	_pipeline.IsStructured = field.NewBool(tableName, "is_structured")
	_pipeline.IsEmbedded = field.NewBool(tableName, "is_embedded")
	_pipeline.IsIndexed = field.NewBool(tableName, "is_indexed")
	_pipeline.CreateTime = field.NewTime(tableName, "create_time")
	_pipeline.UpdateTime = field.NewTime(tableName, "update_time")

	_pipeline.fillFieldMap()

	return _pipeline
}

type pipeline struct {
	pipelineDo

	ALL          field.Asterisk
	ID           field.Int64
	NoteID       field.Uint64
	IsCounted    field.Bool // 交互量已采集
	IsDetailed   field.Bool // 详情页已采集
	IsStructured field.Bool // 内容已结构化
	IsEmbedded   field.Bool // 内容已向量化
	IsIndexed    field.Bool // 内容已入索引
	CreateTime   field.Time
	UpdateTime   field.Time

	fieldMap map[string]field.Expr
}

func (p pipeline) Table(newTableName string) *pipeline {
	p.pipelineDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pipeline) As(alias string) *pipeline {
	p.pipelineDo.DO = *(p.pipelineDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pipeline) updateTableName(table string) *pipeline {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.NoteID = field.NewUint64(table, "note_id")
	p.IsCounted = field.NewBool(table, "is_counted")
	p.IsDetailed = field.NewBool(table, "is_detailed")
	p.IsStructured = field.NewBool(table, "is_structured")
	p.IsEmbedded = field.NewBool(table, "is_embedded")
	p.IsIndexed = field.NewBool(table, "is_indexed")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateTime = field.NewTime(table, "update_time")

	p.fillFieldMap()

	return p
}

func (p *pipeline) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pipeline) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["note_id"] = p.NoteID
	p.fieldMap["is_counted"] = p.IsCounted
	p.fieldMap["is_detailed"] = p.IsDetailed
	p.fieldMap["is_structured"] = p.IsStructured
	p.fieldMap["is_embedded"] = p.IsEmbedded
	p.fieldMap["is_indexed"] = p.IsIndexed
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_time"] = p.UpdateTime
}

func (p pipeline) clone(db *gorm.DB) pipeline {
	p.pipelineDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pipeline) replaceDB(db *gorm.DB) pipeline {
	p.pipelineDo.ReplaceDB(db)
	return p
}

type pipelineDo struct{ gen.DO }

type IPipelineDo interface {
	gen.SubQuery
	Debug() IPipelineDo
	WithContext(ctx context.Context) IPipelineDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPipelineDo
	WriteDB() IPipelineDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPipelineDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPipelineDo
	Not(conds ...gen.Condition) IPipelineDo
	Or(conds ...gen.Condition) IPipelineDo
	Select(conds ...field.Expr) IPipelineDo
	Where(conds ...gen.Condition) IPipelineDo
	Order(conds ...field.Expr) IPipelineDo
	Distinct(cols ...field.Expr) IPipelineDo
	Omit(cols ...field.Expr) IPipelineDo
	Join(table schema.Tabler, on ...field.Expr) IPipelineDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPipelineDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPipelineDo
	Group(cols ...field.Expr) IPipelineDo
	Having(conds ...gen.Condition) IPipelineDo
	Limit(limit int) IPipelineDo
	Offset(offset int) IPipelineDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPipelineDo
	Unscoped() IPipelineDo
	Create(values ...*model.Pipeline) error
	CreateInBatches(values []*model.Pipeline, batchSize int) error
	Save(values ...*model.Pipeline) error
	First() (*model.Pipeline, error)
	Take() (*model.Pipeline, error)
	Last() (*model.Pipeline, error)
	Find() ([]*model.Pipeline, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Pipeline, err error)
	FindInBatches(result *[]*model.Pipeline, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Pipeline) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPipelineDo
	Assign(attrs ...field.AssignExpr) IPipelineDo
	Joins(fields ...field.RelationField) IPipelineDo
	Preload(fields ...field.RelationField) IPipelineDo
	FirstOrInit() (*model.Pipeline, error)
	FirstOrCreate() (*model.Pipeline, error)
	FindByPage(offset int, limit int) (result []*model.Pipeline, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPipelineDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pipelineDo) Debug() IPipelineDo {
	return p.withDO(p.DO.Debug())
}

func (p pipelineDo) WithContext(ctx context.Context) IPipelineDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pipelineDo) ReadDB() IPipelineDo {
	return p.Clauses(dbresolver.Read)
}

func (p pipelineDo) WriteDB() IPipelineDo {
	return p.Clauses(dbresolver.Write)
}

func (p pipelineDo) Session(config *gorm.Session) IPipelineDo {
	return p.withDO(p.DO.Session(config))
}

func (p pipelineDo) Clauses(conds ...clause.Expression) IPipelineDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pipelineDo) Returning(value interface{}, columns ...string) IPipelineDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pipelineDo) Not(conds ...gen.Condition) IPipelineDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pipelineDo) Or(conds ...gen.Condition) IPipelineDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pipelineDo) Select(conds ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pipelineDo) Where(conds ...gen.Condition) IPipelineDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pipelineDo) Order(conds ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pipelineDo) Distinct(cols ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pipelineDo) Omit(cols ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pipelineDo) Join(table schema.Tabler, on ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pipelineDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pipelineDo) RightJoin(table schema.Tabler, on ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pipelineDo) Group(cols ...field.Expr) IPipelineDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pipelineDo) Having(conds ...gen.Condition) IPipelineDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pipelineDo) Limit(limit int) IPipelineDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pipelineDo) Offset(offset int) IPipelineDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pipelineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPipelineDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pipelineDo) Unscoped() IPipelineDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pipelineDo) Create(values ...*model.Pipeline) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pipelineDo) CreateInBatches(values []*model.Pipeline, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pipelineDo) Save(values ...*model.Pipeline) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pipelineDo) First() (*model.Pipeline, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pipeline), nil
	}
}

func (p pipelineDo) Take() (*model.Pipeline, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pipeline), nil
	}
}

func (p pipelineDo) Last() (*model.Pipeline, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pipeline), nil
	}
}

func (p pipelineDo) Find() ([]*model.Pipeline, error) {
	result, err := p.DO.Find()
	return result.([]*model.Pipeline), err
}

func (p pipelineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Pipeline, err error) {
	buf := make([]*model.Pipeline, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pipelineDo) FindInBatches(result *[]*model.Pipeline, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pipelineDo) Attrs(attrs ...field.AssignExpr) IPipelineDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pipelineDo) Assign(attrs ...field.AssignExpr) IPipelineDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pipelineDo) Joins(fields ...field.RelationField) IPipelineDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pipelineDo) Preload(fields ...field.RelationField) IPipelineDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pipelineDo) FirstOrInit() (*model.Pipeline, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pipeline), nil
	}
}

func (p pipelineDo) FirstOrCreate() (*model.Pipeline, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Pipeline), nil
	}
}

func (p pipelineDo) FindByPage(offset int, limit int) (result []*model.Pipeline, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pipelineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pipelineDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pipelineDo) Delete(models ...*model.Pipeline) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pipelineDo) withDO(do gen.Dao) *pipelineDo {
	p.DO = *do.(*gen.DO)
	return p
}

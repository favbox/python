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

func newInteraction(db *gorm.DB, opts ...gen.DOOption) interaction {
	_interaction := interaction{}

	_interaction.interactionDo.UseDB(db, opts...)
	_interaction.interactionDo.UseModel(&model.Interaction{})

	tableName := _interaction.interactionDo.TableName()
	_interaction.ALL = field.NewAsterisk(tableName)
	_interaction.ID = field.NewInt64(tableName, "id")
	_interaction.DocID = field.NewInt64(tableName, "doc_id")
	_interaction.Day = field.NewInt32(tableName, "day")
	_interaction.ReadNum = field.NewInt32(tableName, "read_num")
	_interaction.LikeNum = field.NewInt32(tableName, "like_num")
	_interaction.CollectNum = field.NewInt32(tableName, "collect_num")
	_interaction.DownloadNum = field.NewInt32(tableName, "download_num")
	_interaction.CreateTime = field.NewTime(tableName, "create_time")

	_interaction.fillFieldMap()

	return _interaction
}

type interaction struct {
	interactionDo

	ALL         field.Asterisk
	ID          field.Int64
	DocID       field.Int64
	Day         field.Int32
	ReadNum     field.Int32
	LikeNum     field.Int32
	CollectNum  field.Int32
	DownloadNum field.Int32
	CreateTime  field.Time

	fieldMap map[string]field.Expr
}

func (i interaction) Table(newTableName string) *interaction {
	i.interactionDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i interaction) As(alias string) *interaction {
	i.interactionDo.DO = *(i.interactionDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *interaction) updateTableName(table string) *interaction {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.DocID = field.NewInt64(table, "doc_id")
	i.Day = field.NewInt32(table, "day")
	i.ReadNum = field.NewInt32(table, "read_num")
	i.LikeNum = field.NewInt32(table, "like_num")
	i.CollectNum = field.NewInt32(table, "collect_num")
	i.DownloadNum = field.NewInt32(table, "download_num")
	i.CreateTime = field.NewTime(table, "create_time")

	i.fillFieldMap()

	return i
}

func (i *interaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *interaction) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 8)
	i.fieldMap["id"] = i.ID
	i.fieldMap["doc_id"] = i.DocID
	i.fieldMap["day"] = i.Day
	i.fieldMap["read_num"] = i.ReadNum
	i.fieldMap["like_num"] = i.LikeNum
	i.fieldMap["collect_num"] = i.CollectNum
	i.fieldMap["download_num"] = i.DownloadNum
	i.fieldMap["create_time"] = i.CreateTime
}

func (i interaction) clone(db *gorm.DB) interaction {
	i.interactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i interaction) replaceDB(db *gorm.DB) interaction {
	i.interactionDo.ReplaceDB(db)
	return i
}

type interactionDo struct{ gen.DO }

type IInteractionDo interface {
	gen.SubQuery
	Debug() IInteractionDo
	WithContext(ctx context.Context) IInteractionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInteractionDo
	WriteDB() IInteractionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInteractionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInteractionDo
	Not(conds ...gen.Condition) IInteractionDo
	Or(conds ...gen.Condition) IInteractionDo
	Select(conds ...field.Expr) IInteractionDo
	Where(conds ...gen.Condition) IInteractionDo
	Order(conds ...field.Expr) IInteractionDo
	Distinct(cols ...field.Expr) IInteractionDo
	Omit(cols ...field.Expr) IInteractionDo
	Join(table schema.Tabler, on ...field.Expr) IInteractionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInteractionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInteractionDo
	Group(cols ...field.Expr) IInteractionDo
	Having(conds ...gen.Condition) IInteractionDo
	Limit(limit int) IInteractionDo
	Offset(offset int) IInteractionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInteractionDo
	Unscoped() IInteractionDo
	Create(values ...*model.Interaction) error
	CreateInBatches(values []*model.Interaction, batchSize int) error
	Save(values ...*model.Interaction) error
	First() (*model.Interaction, error)
	Take() (*model.Interaction, error)
	Last() (*model.Interaction, error)
	Find() ([]*model.Interaction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Interaction, err error)
	FindInBatches(result *[]*model.Interaction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Interaction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInteractionDo
	Assign(attrs ...field.AssignExpr) IInteractionDo
	Joins(fields ...field.RelationField) IInteractionDo
	Preload(fields ...field.RelationField) IInteractionDo
	FirstOrInit() (*model.Interaction, error)
	FirstOrCreate() (*model.Interaction, error)
	FindByPage(offset int, limit int) (result []*model.Interaction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInteractionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i interactionDo) Debug() IInteractionDo {
	return i.withDO(i.DO.Debug())
}

func (i interactionDo) WithContext(ctx context.Context) IInteractionDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i interactionDo) ReadDB() IInteractionDo {
	return i.Clauses(dbresolver.Read)
}

func (i interactionDo) WriteDB() IInteractionDo {
	return i.Clauses(dbresolver.Write)
}

func (i interactionDo) Session(config *gorm.Session) IInteractionDo {
	return i.withDO(i.DO.Session(config))
}

func (i interactionDo) Clauses(conds ...clause.Expression) IInteractionDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i interactionDo) Returning(value interface{}, columns ...string) IInteractionDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i interactionDo) Not(conds ...gen.Condition) IInteractionDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i interactionDo) Or(conds ...gen.Condition) IInteractionDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i interactionDo) Select(conds ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i interactionDo) Where(conds ...gen.Condition) IInteractionDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i interactionDo) Order(conds ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i interactionDo) Distinct(cols ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i interactionDo) Omit(cols ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i interactionDo) Join(table schema.Tabler, on ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i interactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i interactionDo) RightJoin(table schema.Tabler, on ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i interactionDo) Group(cols ...field.Expr) IInteractionDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i interactionDo) Having(conds ...gen.Condition) IInteractionDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i interactionDo) Limit(limit int) IInteractionDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i interactionDo) Offset(offset int) IInteractionDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i interactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInteractionDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i interactionDo) Unscoped() IInteractionDo {
	return i.withDO(i.DO.Unscoped())
}

func (i interactionDo) Create(values ...*model.Interaction) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i interactionDo) CreateInBatches(values []*model.Interaction, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i interactionDo) Save(values ...*model.Interaction) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i interactionDo) First() (*model.Interaction, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interaction), nil
	}
}

func (i interactionDo) Take() (*model.Interaction, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interaction), nil
	}
}

func (i interactionDo) Last() (*model.Interaction, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interaction), nil
	}
}

func (i interactionDo) Find() ([]*model.Interaction, error) {
	result, err := i.DO.Find()
	return result.([]*model.Interaction), err
}

func (i interactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Interaction, err error) {
	buf := make([]*model.Interaction, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i interactionDo) FindInBatches(result *[]*model.Interaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i interactionDo) Attrs(attrs ...field.AssignExpr) IInteractionDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i interactionDo) Assign(attrs ...field.AssignExpr) IInteractionDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i interactionDo) Joins(fields ...field.RelationField) IInteractionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i interactionDo) Preload(fields ...field.RelationField) IInteractionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i interactionDo) FirstOrInit() (*model.Interaction, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interaction), nil
	}
}

func (i interactionDo) FirstOrCreate() (*model.Interaction, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interaction), nil
	}
}

func (i interactionDo) FindByPage(offset int, limit int) (result []*model.Interaction, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i interactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i interactionDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i interactionDo) Delete(models ...*model.Interaction) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *interactionDo) withDO(do gen.Dao) *interactionDo {
	i.DO = *do.(*gen.DO)
	return i
}

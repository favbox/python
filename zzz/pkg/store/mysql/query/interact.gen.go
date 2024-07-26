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

func newInteract(db *gorm.DB, opts ...gen.DOOption) interact {
	_interact := interact{}

	_interact.interactDo.UseDB(db, opts...)
	_interact.interactDo.UseModel(&model.Interact{})

	tableName := _interact.interactDo.TableName()
	_interact.ALL = field.NewAsterisk(tableName)
	_interact.ID = field.NewInt64(tableName, "id")
	_interact.NoteID = field.NewUint64(tableName, "note_id")
	_interact.Day = field.NewInt(tableName, "day")
	_interact.ReadCount = field.NewInt(tableName, "read_count")
	_interact.LikedCount = field.NewInt(tableName, "liked_count")
	_interact.CollectedCount = field.NewInt(tableName, "collected_count")
	_interact.CommentCount = field.NewInt(tableName, "comment_count")
	_interact.ShareCount = field.NewInt(tableName, "share_count")
	_interact.CreateTime = field.NewTime(tableName, "create_time")

	_interact.fillFieldMap()

	return _interact
}

type interact struct {
	interactDo

	ALL            field.Asterisk
	ID             field.Int64
	NoteID         field.Uint64
	Day            field.Int
	ReadCount      field.Int // 阅读量
	LikedCount     field.Int // 点赞量
	CollectedCount field.Int // 收藏量
	CommentCount   field.Int // 评论量
	ShareCount     field.Int // 分享量
	CreateTime     field.Time

	fieldMap map[string]field.Expr
}

func (i interact) Table(newTableName string) *interact {
	i.interactDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i interact) As(alias string) *interact {
	i.interactDo.DO = *(i.interactDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *interact) updateTableName(table string) *interact {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.NoteID = field.NewUint64(table, "note_id")
	i.Day = field.NewInt(table, "day")
	i.ReadCount = field.NewInt(table, "read_count")
	i.LikedCount = field.NewInt(table, "liked_count")
	i.CollectedCount = field.NewInt(table, "collected_count")
	i.CommentCount = field.NewInt(table, "comment_count")
	i.ShareCount = field.NewInt(table, "share_count")
	i.CreateTime = field.NewTime(table, "create_time")

	i.fillFieldMap()

	return i
}

func (i *interact) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *interact) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["id"] = i.ID
	i.fieldMap["note_id"] = i.NoteID
	i.fieldMap["day"] = i.Day
	i.fieldMap["read_count"] = i.ReadCount
	i.fieldMap["liked_count"] = i.LikedCount
	i.fieldMap["collected_count"] = i.CollectedCount
	i.fieldMap["comment_count"] = i.CommentCount
	i.fieldMap["share_count"] = i.ShareCount
	i.fieldMap["create_time"] = i.CreateTime
}

func (i interact) clone(db *gorm.DB) interact {
	i.interactDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i interact) replaceDB(db *gorm.DB) interact {
	i.interactDo.ReplaceDB(db)
	return i
}

type interactDo struct{ gen.DO }

type IInteractDo interface {
	gen.SubQuery
	Debug() IInteractDo
	WithContext(ctx context.Context) IInteractDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInteractDo
	WriteDB() IInteractDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInteractDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInteractDo
	Not(conds ...gen.Condition) IInteractDo
	Or(conds ...gen.Condition) IInteractDo
	Select(conds ...field.Expr) IInteractDo
	Where(conds ...gen.Condition) IInteractDo
	Order(conds ...field.Expr) IInteractDo
	Distinct(cols ...field.Expr) IInteractDo
	Omit(cols ...field.Expr) IInteractDo
	Join(table schema.Tabler, on ...field.Expr) IInteractDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInteractDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInteractDo
	Group(cols ...field.Expr) IInteractDo
	Having(conds ...gen.Condition) IInteractDo
	Limit(limit int) IInteractDo
	Offset(offset int) IInteractDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInteractDo
	Unscoped() IInteractDo
	Create(values ...*model.Interact) error
	CreateInBatches(values []*model.Interact, batchSize int) error
	Save(values ...*model.Interact) error
	First() (*model.Interact, error)
	Take() (*model.Interact, error)
	Last() (*model.Interact, error)
	Find() ([]*model.Interact, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Interact, err error)
	FindInBatches(result *[]*model.Interact, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Interact) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInteractDo
	Assign(attrs ...field.AssignExpr) IInteractDo
	Joins(fields ...field.RelationField) IInteractDo
	Preload(fields ...field.RelationField) IInteractDo
	FirstOrInit() (*model.Interact, error)
	FirstOrCreate() (*model.Interact, error)
	FindByPage(offset int, limit int) (result []*model.Interact, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInteractDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i interactDo) Debug() IInteractDo {
	return i.withDO(i.DO.Debug())
}

func (i interactDo) WithContext(ctx context.Context) IInteractDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i interactDo) ReadDB() IInteractDo {
	return i.Clauses(dbresolver.Read)
}

func (i interactDo) WriteDB() IInteractDo {
	return i.Clauses(dbresolver.Write)
}

func (i interactDo) Session(config *gorm.Session) IInteractDo {
	return i.withDO(i.DO.Session(config))
}

func (i interactDo) Clauses(conds ...clause.Expression) IInteractDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i interactDo) Returning(value interface{}, columns ...string) IInteractDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i interactDo) Not(conds ...gen.Condition) IInteractDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i interactDo) Or(conds ...gen.Condition) IInteractDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i interactDo) Select(conds ...field.Expr) IInteractDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i interactDo) Where(conds ...gen.Condition) IInteractDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i interactDo) Order(conds ...field.Expr) IInteractDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i interactDo) Distinct(cols ...field.Expr) IInteractDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i interactDo) Omit(cols ...field.Expr) IInteractDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i interactDo) Join(table schema.Tabler, on ...field.Expr) IInteractDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i interactDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInteractDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i interactDo) RightJoin(table schema.Tabler, on ...field.Expr) IInteractDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i interactDo) Group(cols ...field.Expr) IInteractDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i interactDo) Having(conds ...gen.Condition) IInteractDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i interactDo) Limit(limit int) IInteractDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i interactDo) Offset(offset int) IInteractDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i interactDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInteractDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i interactDo) Unscoped() IInteractDo {
	return i.withDO(i.DO.Unscoped())
}

func (i interactDo) Create(values ...*model.Interact) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i interactDo) CreateInBatches(values []*model.Interact, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i interactDo) Save(values ...*model.Interact) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i interactDo) First() (*model.Interact, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interact), nil
	}
}

func (i interactDo) Take() (*model.Interact, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interact), nil
	}
}

func (i interactDo) Last() (*model.Interact, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interact), nil
	}
}

func (i interactDo) Find() ([]*model.Interact, error) {
	result, err := i.DO.Find()
	return result.([]*model.Interact), err
}

func (i interactDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Interact, err error) {
	buf := make([]*model.Interact, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i interactDo) FindInBatches(result *[]*model.Interact, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i interactDo) Attrs(attrs ...field.AssignExpr) IInteractDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i interactDo) Assign(attrs ...field.AssignExpr) IInteractDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i interactDo) Joins(fields ...field.RelationField) IInteractDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i interactDo) Preload(fields ...field.RelationField) IInteractDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i interactDo) FirstOrInit() (*model.Interact, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interact), nil
	}
}

func (i interactDo) FirstOrCreate() (*model.Interact, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interact), nil
	}
}

func (i interactDo) FindByPage(offset int, limit int) (result []*model.Interact, count int64, err error) {
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

func (i interactDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i interactDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i interactDo) Delete(models ...*model.Interact) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *interactDo) withDO(do gen.Dao) *interactDo {
	i.DO = *do.(*gen.DO)
	return i
}
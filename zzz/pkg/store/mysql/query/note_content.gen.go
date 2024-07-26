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

func newNoteContent(db *gorm.DB, opts ...gen.DOOption) noteContent {
	_noteContent := noteContent{}

	_noteContent.noteContentDo.UseDB(db, opts...)
	_noteContent.noteContentDo.UseModel(&model.NoteContent{})

	tableName := _noteContent.noteContentDo.TableName()
	_noteContent.ALL = field.NewAsterisk(tableName)
	_noteContent.ID = field.NewUint64(tableName, "id")
	_noteContent.NoteID = field.NewInt64(tableName, "note_id")
	_noteContent.Text = field.NewString(tableName, "text")
	_noteContent.UpdateTime = field.NewTime(tableName, "update_time")

	_noteContent.fillFieldMap()

	return _noteContent
}

type noteContent struct {
	noteContentDo

	ALL        field.Asterisk
	ID         field.Uint64
	NoteID     field.Int64
	Text       field.String
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (n noteContent) Table(newTableName string) *noteContent {
	n.noteContentDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n noteContent) As(alias string) *noteContent {
	n.noteContentDo.DO = *(n.noteContentDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *noteContent) updateTableName(table string) *noteContent {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewUint64(table, "id")
	n.NoteID = field.NewInt64(table, "note_id")
	n.Text = field.NewString(table, "text")
	n.UpdateTime = field.NewTime(table, "update_time")

	n.fillFieldMap()

	return n
}

func (n *noteContent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *noteContent) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 4)
	n.fieldMap["id"] = n.ID
	n.fieldMap["note_id"] = n.NoteID
	n.fieldMap["text"] = n.Text
	n.fieldMap["update_time"] = n.UpdateTime
}

func (n noteContent) clone(db *gorm.DB) noteContent {
	n.noteContentDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n noteContent) replaceDB(db *gorm.DB) noteContent {
	n.noteContentDo.ReplaceDB(db)
	return n
}

type noteContentDo struct{ gen.DO }

type INoteContentDo interface {
	gen.SubQuery
	Debug() INoteContentDo
	WithContext(ctx context.Context) INoteContentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INoteContentDo
	WriteDB() INoteContentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INoteContentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INoteContentDo
	Not(conds ...gen.Condition) INoteContentDo
	Or(conds ...gen.Condition) INoteContentDo
	Select(conds ...field.Expr) INoteContentDo
	Where(conds ...gen.Condition) INoteContentDo
	Order(conds ...field.Expr) INoteContentDo
	Distinct(cols ...field.Expr) INoteContentDo
	Omit(cols ...field.Expr) INoteContentDo
	Join(table schema.Tabler, on ...field.Expr) INoteContentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INoteContentDo
	RightJoin(table schema.Tabler, on ...field.Expr) INoteContentDo
	Group(cols ...field.Expr) INoteContentDo
	Having(conds ...gen.Condition) INoteContentDo
	Limit(limit int) INoteContentDo
	Offset(offset int) INoteContentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INoteContentDo
	Unscoped() INoteContentDo
	Create(values ...*model.NoteContent) error
	CreateInBatches(values []*model.NoteContent, batchSize int) error
	Save(values ...*model.NoteContent) error
	First() (*model.NoteContent, error)
	Take() (*model.NoteContent, error)
	Last() (*model.NoteContent, error)
	Find() ([]*model.NoteContent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NoteContent, err error)
	FindInBatches(result *[]*model.NoteContent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NoteContent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INoteContentDo
	Assign(attrs ...field.AssignExpr) INoteContentDo
	Joins(fields ...field.RelationField) INoteContentDo
	Preload(fields ...field.RelationField) INoteContentDo
	FirstOrInit() (*model.NoteContent, error)
	FirstOrCreate() (*model.NoteContent, error)
	FindByPage(offset int, limit int) (result []*model.NoteContent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INoteContentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n noteContentDo) Debug() INoteContentDo {
	return n.withDO(n.DO.Debug())
}

func (n noteContentDo) WithContext(ctx context.Context) INoteContentDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n noteContentDo) ReadDB() INoteContentDo {
	return n.Clauses(dbresolver.Read)
}

func (n noteContentDo) WriteDB() INoteContentDo {
	return n.Clauses(dbresolver.Write)
}

func (n noteContentDo) Session(config *gorm.Session) INoteContentDo {
	return n.withDO(n.DO.Session(config))
}

func (n noteContentDo) Clauses(conds ...clause.Expression) INoteContentDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n noteContentDo) Returning(value interface{}, columns ...string) INoteContentDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n noteContentDo) Not(conds ...gen.Condition) INoteContentDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n noteContentDo) Or(conds ...gen.Condition) INoteContentDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n noteContentDo) Select(conds ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n noteContentDo) Where(conds ...gen.Condition) INoteContentDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n noteContentDo) Order(conds ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n noteContentDo) Distinct(cols ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n noteContentDo) Omit(cols ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n noteContentDo) Join(table schema.Tabler, on ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n noteContentDo) LeftJoin(table schema.Tabler, on ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n noteContentDo) RightJoin(table schema.Tabler, on ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n noteContentDo) Group(cols ...field.Expr) INoteContentDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n noteContentDo) Having(conds ...gen.Condition) INoteContentDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n noteContentDo) Limit(limit int) INoteContentDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n noteContentDo) Offset(offset int) INoteContentDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n noteContentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INoteContentDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n noteContentDo) Unscoped() INoteContentDo {
	return n.withDO(n.DO.Unscoped())
}

func (n noteContentDo) Create(values ...*model.NoteContent) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n noteContentDo) CreateInBatches(values []*model.NoteContent, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n noteContentDo) Save(values ...*model.NoteContent) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n noteContentDo) First() (*model.NoteContent, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoteContent), nil
	}
}

func (n noteContentDo) Take() (*model.NoteContent, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoteContent), nil
	}
}

func (n noteContentDo) Last() (*model.NoteContent, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoteContent), nil
	}
}

func (n noteContentDo) Find() ([]*model.NoteContent, error) {
	result, err := n.DO.Find()
	return result.([]*model.NoteContent), err
}

func (n noteContentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NoteContent, err error) {
	buf := make([]*model.NoteContent, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n noteContentDo) FindInBatches(result *[]*model.NoteContent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n noteContentDo) Attrs(attrs ...field.AssignExpr) INoteContentDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n noteContentDo) Assign(attrs ...field.AssignExpr) INoteContentDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n noteContentDo) Joins(fields ...field.RelationField) INoteContentDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n noteContentDo) Preload(fields ...field.RelationField) INoteContentDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n noteContentDo) FirstOrInit() (*model.NoteContent, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoteContent), nil
	}
}

func (n noteContentDo) FirstOrCreate() (*model.NoteContent, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoteContent), nil
	}
}

func (n noteContentDo) FindByPage(offset int, limit int) (result []*model.NoteContent, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n noteContentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n noteContentDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n noteContentDo) Delete(models ...*model.NoteContent) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *noteContentDo) withDO(do gen.Dao) *noteContentDo {
	n.DO = *do.(*gen.DO)
	return n
}
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package persistence

import (
	"context"
	"pigeon2/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newWorkNode(db *gorm.DB, opts ...gen.DOOption) workNode {
	_workNode := workNode{}

	_workNode.workNodeDo.UseDB(db, opts...)
	_workNode.workNodeDo.UseModel(&model.WorkNode{})

	tableName := _workNode.workNodeDo.TableName()
	_workNode.ALL = field.NewAsterisk(tableName)
	_workNode.ID = field.NewInt64(tableName, "id")
	_workNode.ParentID = field.NewInt(tableName, "parent_id")
	_workNode.WorkFlowID = field.NewInt64(tableName, "work_flow_id")
	_workNode.Type = field.NewInt(tableName, "type")
	_workNode.Name = field.NewString(tableName, "name")

	_workNode.fillFieldMap()

	return _workNode
}

type workNode struct {
	workNodeDo

	ALL        field.Asterisk
	ID         field.Int64
	ParentID   field.Int
	WorkFlowID field.Int64
	Type       field.Int
	Name       field.String

	fieldMap map[string]field.Expr
}

func (w workNode) Table(newTableName string) *workNode {
	w.workNodeDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w workNode) As(alias string) *workNode {
	w.workNodeDo.DO = *(w.workNodeDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *workNode) updateTableName(table string) *workNode {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.ParentID = field.NewInt(table, "parent_id")
	w.WorkFlowID = field.NewInt64(table, "work_flow_id")
	w.Type = field.NewInt(table, "type")
	w.Name = field.NewString(table, "name")

	w.fillFieldMap()

	return w
}

func (w *workNode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *workNode) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["id"] = w.ID
	w.fieldMap["parent_id"] = w.ParentID
	w.fieldMap["work_flow_id"] = w.WorkFlowID
	w.fieldMap["type"] = w.Type
	w.fieldMap["name"] = w.Name
}

func (w workNode) clone(db *gorm.DB) workNode {
	w.workNodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w workNode) replaceDB(db *gorm.DB) workNode {
	w.workNodeDo.ReplaceDB(db)
	return w
}

type workNodeDo struct{ gen.DO }

type IWorkNodeDo interface {
	gen.SubQuery
	Debug() IWorkNodeDo
	WithContext(ctx context.Context) IWorkNodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWorkNodeDo
	WriteDB() IWorkNodeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWorkNodeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWorkNodeDo
	Not(conds ...gen.Condition) IWorkNodeDo
	Or(conds ...gen.Condition) IWorkNodeDo
	Select(conds ...field.Expr) IWorkNodeDo
	Where(conds ...gen.Condition) IWorkNodeDo
	Order(conds ...field.Expr) IWorkNodeDo
	Distinct(cols ...field.Expr) IWorkNodeDo
	Omit(cols ...field.Expr) IWorkNodeDo
	Join(table schema.Tabler, on ...field.Expr) IWorkNodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWorkNodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWorkNodeDo
	Group(cols ...field.Expr) IWorkNodeDo
	Having(conds ...gen.Condition) IWorkNodeDo
	Limit(limit int) IWorkNodeDo
	Offset(offset int) IWorkNodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWorkNodeDo
	Unscoped() IWorkNodeDo
	Create(values ...*model.WorkNode) error
	CreateInBatches(values []*model.WorkNode, batchSize int) error
	Save(values ...*model.WorkNode) error
	First() (*model.WorkNode, error)
	Take() (*model.WorkNode, error)
	Last() (*model.WorkNode, error)
	Find() ([]*model.WorkNode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WorkNode, err error)
	FindInBatches(result *[]*model.WorkNode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WorkNode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWorkNodeDo
	Assign(attrs ...field.AssignExpr) IWorkNodeDo
	Joins(fields ...field.RelationField) IWorkNodeDo
	Preload(fields ...field.RelationField) IWorkNodeDo
	FirstOrInit() (*model.WorkNode, error)
	FirstOrCreate() (*model.WorkNode, error)
	FindByPage(offset int, limit int) (result []*model.WorkNode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWorkNodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w workNodeDo) Debug() IWorkNodeDo {
	return w.withDO(w.DO.Debug())
}

func (w workNodeDo) WithContext(ctx context.Context) IWorkNodeDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w workNodeDo) ReadDB() IWorkNodeDo {
	return w.Clauses(dbresolver.Read)
}

func (w workNodeDo) WriteDB() IWorkNodeDo {
	return w.Clauses(dbresolver.Write)
}

func (w workNodeDo) Session(config *gorm.Session) IWorkNodeDo {
	return w.withDO(w.DO.Session(config))
}

func (w workNodeDo) Clauses(conds ...clause.Expression) IWorkNodeDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w workNodeDo) Returning(value interface{}, columns ...string) IWorkNodeDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w workNodeDo) Not(conds ...gen.Condition) IWorkNodeDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w workNodeDo) Or(conds ...gen.Condition) IWorkNodeDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w workNodeDo) Select(conds ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w workNodeDo) Where(conds ...gen.Condition) IWorkNodeDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w workNodeDo) Order(conds ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w workNodeDo) Distinct(cols ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w workNodeDo) Omit(cols ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w workNodeDo) Join(table schema.Tabler, on ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w workNodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w workNodeDo) RightJoin(table schema.Tabler, on ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w workNodeDo) Group(cols ...field.Expr) IWorkNodeDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w workNodeDo) Having(conds ...gen.Condition) IWorkNodeDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w workNodeDo) Limit(limit int) IWorkNodeDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w workNodeDo) Offset(offset int) IWorkNodeDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w workNodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWorkNodeDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w workNodeDo) Unscoped() IWorkNodeDo {
	return w.withDO(w.DO.Unscoped())
}

func (w workNodeDo) Create(values ...*model.WorkNode) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w workNodeDo) CreateInBatches(values []*model.WorkNode, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w workNodeDo) Save(values ...*model.WorkNode) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w workNodeDo) First() (*model.WorkNode, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkNode), nil
	}
}

func (w workNodeDo) Take() (*model.WorkNode, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkNode), nil
	}
}

func (w workNodeDo) Last() (*model.WorkNode, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkNode), nil
	}
}

func (w workNodeDo) Find() ([]*model.WorkNode, error) {
	result, err := w.DO.Find()
	return result.([]*model.WorkNode), err
}

func (w workNodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WorkNode, err error) {
	buf := make([]*model.WorkNode, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w workNodeDo) FindInBatches(result *[]*model.WorkNode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w workNodeDo) Attrs(attrs ...field.AssignExpr) IWorkNodeDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w workNodeDo) Assign(attrs ...field.AssignExpr) IWorkNodeDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w workNodeDo) Joins(fields ...field.RelationField) IWorkNodeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w workNodeDo) Preload(fields ...field.RelationField) IWorkNodeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w workNodeDo) FirstOrInit() (*model.WorkNode, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkNode), nil
	}
}

func (w workNodeDo) FirstOrCreate() (*model.WorkNode, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkNode), nil
	}
}

func (w workNodeDo) FindByPage(offset int, limit int) (result []*model.WorkNode, count int64, err error) {
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

func (w workNodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w workNodeDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w workNodeDo) Delete(models ...*model.WorkNode) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *workNodeDo) withDO(do gen.Dao) *workNodeDo {
	w.DO = *do.(*gen.DO)
	return w
}

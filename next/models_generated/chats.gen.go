// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newChat(db *gorm.DB, opts ...gen.DOOption) chat {
	_chat := chat{}

	_chat.chatDo.UseDB(db, opts...)
	_chat.chatDo.UseModel(&model.Chat{})

	tableName := _chat.chatDo.TableName()
	_chat.ALL = field.NewAsterisk(tableName)
	_chat.ID = field.NewString(tableName, "id")
	_chat.UserID = field.NewString(tableName, "user_id")
	_chat.Payload = field.NewString(tableName, "payload")
	_chat.CreatedAt = field.NewTime(tableName, "created_at")
	_chat.ProjectID = field.NewString(tableName, "project_id")

	_chat.fillFieldMap()

	return _chat
}

type chat struct {
	chatDo

	ALL       field.Asterisk
	ID        field.String
	UserID    field.String
	Payload   field.String
	CreatedAt field.Time
	ProjectID field.String

	fieldMap map[string]field.Expr
}

func (c chat) Table(newTableName string) *chat {
	c.chatDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chat) As(alias string) *chat {
	c.chatDo.DO = *(c.chatDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chat) updateTableName(table string) *chat {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewString(table, "id")
	c.UserID = field.NewString(table, "user_id")
	c.Payload = field.NewString(table, "payload")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.ProjectID = field.NewString(table, "project_id")

	c.fillFieldMap()

	return c
}

func (c *chat) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chat) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["payload"] = c.Payload
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["project_id"] = c.ProjectID
}

func (c chat) clone(db *gorm.DB) chat {
	c.chatDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chat) replaceDB(db *gorm.DB) chat {
	c.chatDo.ReplaceDB(db)
	return c
}

type chatDo struct{ gen.DO }

type IChatDo interface {
	gen.SubQuery
	Debug() IChatDo
	WithContext(ctx context.Context) IChatDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChatDo
	WriteDB() IChatDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChatDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChatDo
	Not(conds ...gen.Condition) IChatDo
	Or(conds ...gen.Condition) IChatDo
	Select(conds ...field.Expr) IChatDo
	Where(conds ...gen.Condition) IChatDo
	Order(conds ...field.Expr) IChatDo
	Distinct(cols ...field.Expr) IChatDo
	Omit(cols ...field.Expr) IChatDo
	Join(table schema.Tabler, on ...field.Expr) IChatDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChatDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChatDo
	Group(cols ...field.Expr) IChatDo
	Having(conds ...gen.Condition) IChatDo
	Limit(limit int) IChatDo
	Offset(offset int) IChatDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChatDo
	Unscoped() IChatDo
	Create(values ...*model.Chat) error
	CreateInBatches(values []*model.Chat, batchSize int) error
	Save(values ...*model.Chat) error
	First() (*model.Chat, error)
	Take() (*model.Chat, error)
	Last() (*model.Chat, error)
	Find() ([]*model.Chat, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Chat, err error)
	FindInBatches(result *[]*model.Chat, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Chat) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChatDo
	Assign(attrs ...field.AssignExpr) IChatDo
	Joins(fields ...field.RelationField) IChatDo
	Preload(fields ...field.RelationField) IChatDo
	FirstOrInit() (*model.Chat, error)
	FirstOrCreate() (*model.Chat, error)
	FindByPage(offset int, limit int) (result []*model.Chat, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChatDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c chatDo) Debug() IChatDo {
	return c.withDO(c.DO.Debug())
}

func (c chatDo) WithContext(ctx context.Context) IChatDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chatDo) ReadDB() IChatDo {
	return c.Clauses(dbresolver.Read)
}

func (c chatDo) WriteDB() IChatDo {
	return c.Clauses(dbresolver.Write)
}

func (c chatDo) Session(config *gorm.Session) IChatDo {
	return c.withDO(c.DO.Session(config))
}

func (c chatDo) Clauses(conds ...clause.Expression) IChatDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chatDo) Returning(value interface{}, columns ...string) IChatDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chatDo) Not(conds ...gen.Condition) IChatDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chatDo) Or(conds ...gen.Condition) IChatDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chatDo) Select(conds ...field.Expr) IChatDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chatDo) Where(conds ...gen.Condition) IChatDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chatDo) Order(conds ...field.Expr) IChatDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chatDo) Distinct(cols ...field.Expr) IChatDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chatDo) Omit(cols ...field.Expr) IChatDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chatDo) Join(table schema.Tabler, on ...field.Expr) IChatDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chatDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChatDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chatDo) RightJoin(table schema.Tabler, on ...field.Expr) IChatDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chatDo) Group(cols ...field.Expr) IChatDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chatDo) Having(conds ...gen.Condition) IChatDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chatDo) Limit(limit int) IChatDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chatDo) Offset(offset int) IChatDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chatDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChatDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chatDo) Unscoped() IChatDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chatDo) Create(values ...*model.Chat) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chatDo) CreateInBatches(values []*model.Chat, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chatDo) Save(values ...*model.Chat) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chatDo) First() (*model.Chat, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chat), nil
	}
}

func (c chatDo) Take() (*model.Chat, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chat), nil
	}
}

func (c chatDo) Last() (*model.Chat, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chat), nil
	}
}

func (c chatDo) Find() ([]*model.Chat, error) {
	result, err := c.DO.Find()
	return result.([]*model.Chat), err
}

func (c chatDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Chat, err error) {
	buf := make([]*model.Chat, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chatDo) FindInBatches(result *[]*model.Chat, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chatDo) Attrs(attrs ...field.AssignExpr) IChatDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chatDo) Assign(attrs ...field.AssignExpr) IChatDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chatDo) Joins(fields ...field.RelationField) IChatDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chatDo) Preload(fields ...field.RelationField) IChatDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chatDo) FirstOrInit() (*model.Chat, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chat), nil
	}
}

func (c chatDo) FirstOrCreate() (*model.Chat, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chat), nil
	}
}

func (c chatDo) FindByPage(offset int, limit int) (result []*model.Chat, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chatDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chatDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chatDo) Delete(models ...*model.Chat) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chatDo) withDO(do gen.Dao) *chatDo {
	c.DO = *do.(*gen.DO)
	return c
}

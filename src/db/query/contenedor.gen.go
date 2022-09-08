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

	"Mindia/Stock1/Stock/src/db/model"
)

func newContenedor(db *gorm.DB) contenedor {
	_contenedor := contenedor{}

	_contenedor.contenedorDo.UseDB(db)
	_contenedor.contenedorDo.UseModel(&model.Contenedor{})

	tableName := _contenedor.contenedorDo.TableName()
	_contenedor.ALL = field.NewAsterisk(tableName)
	_contenedor.ID = field.NewInt32(tableName, "id")
	_contenedor.Nombre = field.NewString(tableName, "nombre")

	_contenedor.fillFieldMap()

	return _contenedor
}

type contenedor struct {
	contenedorDo

	ALL    field.Asterisk
	ID     field.Int32
	Nombre field.String

	fieldMap map[string]field.Expr
}

func (c contenedor) Table(newTableName string) *contenedor {
	c.contenedorDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c contenedor) As(alias string) *contenedor {
	c.contenedorDo.DO = *(c.contenedorDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *contenedor) updateTableName(table string) *contenedor {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.Nombre = field.NewString(table, "nombre")

	c.fillFieldMap()

	return c
}

func (c *contenedor) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *contenedor) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["id"] = c.ID
	c.fieldMap["nombre"] = c.Nombre
}

func (c contenedor) clone(db *gorm.DB) contenedor {
	c.contenedorDo.ReplaceDB(db)
	return c
}

type contenedorDo struct{ gen.DO }

func (c contenedorDo) Debug() *contenedorDo {
	return c.withDO(c.DO.Debug())
}

func (c contenedorDo) WithContext(ctx context.Context) *contenedorDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c contenedorDo) ReadDB() *contenedorDo {
	return c.Clauses(dbresolver.Read)
}

func (c contenedorDo) WriteDB() *contenedorDo {
	return c.Clauses(dbresolver.Write)
}

func (c contenedorDo) Clauses(conds ...clause.Expression) *contenedorDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c contenedorDo) Returning(value interface{}, columns ...string) *contenedorDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c contenedorDo) Not(conds ...gen.Condition) *contenedorDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c contenedorDo) Or(conds ...gen.Condition) *contenedorDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c contenedorDo) Select(conds ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c contenedorDo) Where(conds ...gen.Condition) *contenedorDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c contenedorDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *contenedorDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c contenedorDo) Order(conds ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c contenedorDo) Distinct(cols ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c contenedorDo) Omit(cols ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c contenedorDo) Join(table schema.Tabler, on ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c contenedorDo) LeftJoin(table schema.Tabler, on ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c contenedorDo) RightJoin(table schema.Tabler, on ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c contenedorDo) Group(cols ...field.Expr) *contenedorDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c contenedorDo) Having(conds ...gen.Condition) *contenedorDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c contenedorDo) Limit(limit int) *contenedorDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c contenedorDo) Offset(offset int) *contenedorDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c contenedorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *contenedorDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c contenedorDo) Unscoped() *contenedorDo {
	return c.withDO(c.DO.Unscoped())
}

func (c contenedorDo) Create(values ...*model.Contenedor) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c contenedorDo) CreateInBatches(values []*model.Contenedor, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c contenedorDo) Save(values ...*model.Contenedor) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c contenedorDo) First() (*model.Contenedor, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Contenedor), nil
	}
}

func (c contenedorDo) Take() (*model.Contenedor, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Contenedor), nil
	}
}

func (c contenedorDo) Last() (*model.Contenedor, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Contenedor), nil
	}
}

func (c contenedorDo) Find() ([]*model.Contenedor, error) {
	result, err := c.DO.Find()
	return result.([]*model.Contenedor), err
}

func (c contenedorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Contenedor, err error) {
	buf := make([]*model.Contenedor, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c contenedorDo) FindInBatches(result *[]*model.Contenedor, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c contenedorDo) Attrs(attrs ...field.AssignExpr) *contenedorDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c contenedorDo) Assign(attrs ...field.AssignExpr) *contenedorDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c contenedorDo) Joins(fields ...field.RelationField) *contenedorDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c contenedorDo) Preload(fields ...field.RelationField) *contenedorDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c contenedorDo) FirstOrInit() (*model.Contenedor, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Contenedor), nil
	}
}

func (c contenedorDo) FirstOrCreate() (*model.Contenedor, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Contenedor), nil
	}
}

func (c contenedorDo) FindByPage(offset int, limit int) (result []*model.Contenedor, count int64, err error) {
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

func (c contenedorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c contenedorDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c contenedorDo) Delete(models ...*model.Contenedor) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *contenedorDo) withDO(do gen.Dao) *contenedorDo {
	c.DO = *do.(*gen.DO)
	return c
}

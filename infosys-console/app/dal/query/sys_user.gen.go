// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"com.alex.infosys/app/dal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newSysUser(db *gorm.DB) sysUser {
	_sysUser := sysUser{}

	_sysUser.sysUserDo.UseDB(db)
	_sysUser.sysUserDo.UseModel(&model.SysUser{})

	tableName := _sysUser.sysUserDo.TableName()
	_sysUser.ALL = field.NewField(tableName, "*")
	_sysUser.ID = field.NewInt64(tableName, "id")
	_sysUser.UserName = field.NewString(tableName, "user_name")
	_sysUser.Password = field.NewString(tableName, "password")
	_sysUser.Phone = field.NewString(tableName, "phone")
	_sysUser.RealName = field.NewString(tableName, "real_name")
	_sysUser.Pid = field.NewInt64(tableName, "pid")
	_sysUser.RoleID = field.NewString(tableName, "role_id")
	_sysUser.Idcard = field.NewString(tableName, "idcard")
	_sysUser.OrgnicationID = field.NewInt32(tableName, "orgnication_id")
	_sysUser.CreateAt = field.NewTime(tableName, "create_at")
	_sysUser.UpdateAt = field.NewTime(tableName, "update_at")
	_sysUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_sysUser.CreateBy = field.NewInt32(tableName, "create_by")

	_sysUser.fillFieldMap()

	return _sysUser
}

type sysUser struct {
	sysUserDo

	ALL           field.Field
	ID            field.Int64
	UserName      field.String
	Password      field.String
	Phone         field.String
	RealName      field.String
	Pid           field.Int64
	RoleID        field.String
	Idcard        field.String
	OrgnicationID field.Int32
	CreateAt      field.Time
	UpdateAt      field.Time
	DeletedAt     field.Field
	CreateBy      field.Int32

	fieldMap map[string]field.Expr
}

func (s sysUser) Table(newTableName string) *sysUser {
	s.sysUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUser) As(alias string) *sysUser {
	s.sysUserDo.DO = *(s.sysUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUser) updateTableName(table string) *sysUser {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.UserName = field.NewString(table, "user_name")
	s.Password = field.NewString(table, "password")
	s.Phone = field.NewString(table, "phone")
	s.RealName = field.NewString(table, "real_name")
	s.Pid = field.NewInt64(table, "pid")
	s.RoleID = field.NewString(table, "role_id")
	s.Idcard = field.NewString(table, "idcard")
	s.OrgnicationID = field.NewInt32(table, "orgnication_id")
	s.CreateAt = field.NewTime(table, "create_at")
	s.UpdateAt = field.NewTime(table, "update_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.CreateBy = field.NewInt32(table, "create_by")

	s.fillFieldMap()

	return s
}

func (s *sysUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["id"] = s.ID
	s.fieldMap["user_name"] = s.UserName
	s.fieldMap["password"] = s.Password
	s.fieldMap["phone"] = s.Phone
	s.fieldMap["real_name"] = s.RealName
	s.fieldMap["pid"] = s.Pid
	s.fieldMap["role_id"] = s.RoleID
	s.fieldMap["idcard"] = s.Idcard
	s.fieldMap["orgnication_id"] = s.OrgnicationID
	s.fieldMap["create_at"] = s.CreateAt
	s.fieldMap["update_at"] = s.UpdateAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["create_by"] = s.CreateBy
}

func (s sysUser) clone(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceDB(db)
	return s
}

type sysUserDo struct{ gen.DO }

func (s sysUserDo) Debug() *sysUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserDo) WithContext(ctx context.Context) *sysUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserDo) Clauses(conds ...clause.Expression) *sysUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserDo) Returning(value interface{}, columns ...string) *sysUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserDo) Not(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserDo) Or(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserDo) Select(conds ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserDo) Where(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysUserDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysUserDo) Order(conds ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserDo) Distinct(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserDo) Omit(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserDo) Join(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserDo) Group(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserDo) Having(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserDo) Limit(limit int) *sysUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserDo) Offset(offset int) *sysUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserDo) Unscoped() *sysUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserDo) Create(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserDo) CreateInBatches(values []*model.SysUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserDo) Save(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserDo) First() (*model.SysUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Take() (*model.SysUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Last() (*model.SysUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Find() ([]*model.SysUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUser), err
}

func (s sysUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUser, err error) {
	buf := make([]*model.SysUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserDo) FindInBatches(result *[]*model.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserDo) Attrs(attrs ...field.AssignExpr) *sysUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserDo) Assign(attrs ...field.AssignExpr) *sysUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserDo) Joins(field field.RelationField) *sysUserDo {
	return s.withDO(s.DO.Joins(field))
}

func (s sysUserDo) Preload(field field.RelationField) *sysUserDo {
	return s.withDO(s.DO.Preload(field))
}

func (s sysUserDo) FirstOrInit() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FirstOrCreate() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FindByPage(offset int, limit int) (result []*model.SysUser, count int64, err error) {
	if limit <= 0 {
		count, err = s.Count()
		return
	}

	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s *sysUserDo) withDO(do gen.Dao) *sysUserDo {
	s.DO = *do.(*gen.DO)
	return s
}

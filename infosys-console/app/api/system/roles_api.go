package system

import (
	"com.alex.infosys/app/dal/model"
	"com.alex.infosys/app/dal/query"
	"com.alex.infosys/pkg/g"
	"com.alex.infosys/pkg/utils/upage"
	"gorm.io/gorm"
)

type roles struct {
	q  *query.Query
	db *gorm.DB
}

//@Summary 角色
//@accept json
//@Produce json
//@Tags SysRole
//@Param post body model.SysRole true "json"
//@Success 200 {object} g.Response{} "result"
//@Router /api/system/roles/updates [post]
func (o *roles) createOrupdate(ctx *g.Context) (interface{}, error) {
	parmsModel := &model.SysRole{}
	t := o.q.SysRole
	if err := ctx.BindParams(parmsModel); err != nil {
		return nil, err
	}
	if parmsModel.ID == 0 {
		return nil, t.Create(parmsModel)
	} else {
		_, err := t.Where(t.ID.Eq(parmsModel.ID)).Updates(&parmsModel)
		return nil, err
	}
}

//@Summary 角色
//@accept json
//@Produce json
//@Tags SysRole
//@Param post body upage.PageQuery true "json"
//@Success 200 {object} g.Response{} "result"
//@Router/api/system/roles/query [post]
func (o *roles) conditionQuery(ctx *g.Context) (interface{}, error) {
	pageQuery := &upage.PageQuery{}
	var (
		datas []model.SysRole
	)

	t := o.q.SysRole
	if err := ctx.BindParams(pageQuery); err != nil {
		return nil, err
	}
	total, err := upage.GenSql(t.TableName(), &datas, pageQuery)

	if err != nil {
		return nil, err
	}

	return &g.Response{
		Total: total,
		Data:  datas,
	}, nil
}

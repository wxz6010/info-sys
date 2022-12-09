package system

import (
	"com.alex.infosys/app/dal/model"
	"com.alex.infosys/app/dal/query"
	"com.alex.infosys/pkg/g"
	"com.alex.infosys/pkg/utils/upage"
	"gorm.io/gorm"
)

type orgs struct {
	q  *query.Query
	db *gorm.DB
}

//@Summary 组织机构
//@accept json
//@Produce json
//@Tags organization
//@Param post body model.SysOrgnization true "json"
//@Success 200 {object} g.Response{} "result"
//@Router /api/system/orgs/updates [post]
func (o orgs) createOrupdate(ctx *g.Context) (interface{}, error) {
	parmsModel := &model.SysOrgnization{}
	t := o.q.SysOrgnization
	if err := ctx.BindParams(parmsModel); err != nil {
		return nil, err
	}
	if parmsModel.ID == 0 {
		return nil, t.Create(parmsModel)
	} else {
		_, err := t.Where(t.ID.Eq(parmsModel.ID)).Updates(parmsModel)
		return nil, err
	}
}

//@Summary 组织机构
//@accept json
//@Produce json
//@Tags organization
//@Param post body upage.PageQuery true "json"
//@Success 200 {object} g.Response{data=[]model.SysOrgnization} "result"
//@Router /api/system/orgs/query [post]
func (o *orgs) conditionQuery(ctx *g.Context) (interface{}, error) {
	pageQuery := &upage.PageQuery{}
	var (
		datas []model.SysOrgnization
	)

	t := o.q.SysOrgnization
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

func New() {

}

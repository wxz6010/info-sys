package system

import (
	"com.alex.infosys/app/dal/model"
	"com.alex.infosys/app/dal/query"
	"com.alex.infosys/pkg/g"
	"com.alex.infosys/pkg/utils/upage"
	"gorm.io/gorm"
)

type orgType struct {
	q  *query.Query
	db *gorm.DB
}

//@Summary 机构类型
//@accept json
//@Produce json
//@Tags orgnizationType
//@Param post body model.OrgnizationType true "json"
//@Success 200 {object} g.Response{} "result"
//@Router /api/system/orgs/type/updates [post]
func (o orgType) createOrupdate(ctx *g.Context) (interface{}, error) {
	parmsModel := &model.OrgnizationType{}
	t := o.q.OrgnizationType
	if err := ctx.BindParams(parmsModel); err != nil {
		return nil, err
	}
	if parmsModel.ID == 0 {
		return nil, t.Create(parmsModel)
	} else {
		_, err := t.Updates(&parmsModel)
		return nil, err
	}
}

//@Summary 机构类型
//@accept json
//@Produce json
//@Tags orgnizationType
//@Param post body upage.PageQuery true "json"
//@Success 200 {object} g.Response{} "result"
//@Router /api/system/orgs/type/query [post]
func (o *orgType) conditionQuery(ctx *g.Context) (interface{}, error) {
	pageQuery := &upage.PageQuery{}
	var (
		datas model.OrgnizationType
	)

	t := o.q.OrgnizationType
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

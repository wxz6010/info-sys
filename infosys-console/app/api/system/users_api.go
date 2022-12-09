package system

import (
	"com.alex.infosys/app/dal/model"
	"com.alex.infosys/app/dal/query"
	"com.alex.infosys/pkg/g"
	"com.alex.infosys/pkg/utils/upage"
	"gorm.io/gorm"
)

type users struct {
	q  *query.Query
	db *gorm.DB
}

func (o *users) createOrupdate(ctx *g.Context) (interface{}, error) {
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

func (o *users) conditionQuery(ctx *g.Context) (interface{}, error) {
	pageQuery := &upage.PageQuery{}
	var (
		datas model.OrgnizationType
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

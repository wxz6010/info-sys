package system

import (
	"com.alex.infosys/app/dal/model"
	"com.alex.infosys/app/dal/query"
	"com.alex.infosys/pkg/g"
	"com.alex.infosys/pkg/utils/upage"
	"gorm.io/gorm"
)

type dictionary struct {
	q  *query.Query
	db *gorm.DB
}

//@Summary 系统数据字典
//@accept json
//@Produce json
//@Tags dictionary
//@Param post body model.SysDictionary true "json"
//@Success 200 {object} g.Response{} "result"
//@Router /api/system/dic/updates [post]
func (o *dictionary) createOrupdate(ctx *g.Context) (interface{}, error) {
	parmsModel := &model.SysDictionary{}
	t := o.q.SysDictionary
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

//@Summary 系统数据字典
//@accept json
//@Produce json
//@Tags dictionary
//@Param post body  upage.PageQuery true "json"
//@Success 200 {object} g.Response{} "result"
//@Router /api/system/dic/query [post]
func (o *dictionary) conditionQuery(ctx *g.Context) (interface{}, error) {
	pageQuery := &upage.PageQuery{}
	var (
		datas []model.SysDictionary
	)

	t := o.q.SysDictionary
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

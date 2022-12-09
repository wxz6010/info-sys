package upage

import (
	"com.alex.infosys/conf"
	"com.alex.infosys/pkg/g"
	"fmt"
	"strings"
)

type PageQuery struct {
	PageSize int         `json:"pageSize"`
	Current  int         `json:"current"`
	Sorts    g.MapStrStr `json:"sorts"`
	Query    g.MapStrAny `json:"query"`
} //@name PageQuery

func GenSql(tableName string, datas interface{}, query *PageQuery) (int, error) {
	return GenSelectsSql(tableName, "*", datas, query)
}

//构建分页查询Sql 支持 mysql mssql
func GenSelectsSql(tableName string, selects string, datas interface{}, page *PageQuery) (int, error) {
	var (
		wheres   []string
		sorts    []string
		sql      string
		totalSql string
		total    int
	)

	delete(page.Query, "current")
	delete(page.Query, "pageSize")

	for s, i := range page.Query {
		switch i.(type) {
		case string:
			wheres = append(wheres, s+" like @"+s+" ")
			page.Query[s] = fmt.Sprintf("%%%s%%", i)
		case int, float32, float64:
			wheres = append(wheres, " "+s+" = @"+s+" ")
		}
	}

	page.Query["limit"] = page.Current * page.PageSize
	page.Query["offset"] = (page.Current - 1) * page.PageSize

	for s, i := range page.Sorts {
		switch i {
		case "descend":
			sorts = append(sorts, s+" desc")
		case "ascend":
			sorts = append(sorts, s+"asc")
		}
	}

	if conf.DB_TYPE == "MYSQLNDS" {
		sql = "select " + selects + " from " + tableName
		totalSql = "select count(*) as total from " + tableName
		if len(wheres) > 0 {
			where := " where " + strings.Join(wheres, " and ")
			sql += where
			totalSql += where
		}
		if len(sorts) > 0 {
			sql += " order by " + strings.Join(sorts, " , ")
		} else {
			sql += " order by id desc"
		}
		sql += " LIMIT @limit offset @offset "
	}

	if conf.DB_TYPE == "MSSQLDNS" {
		var orderBy string
		where := ""
		if len(sorts) > 0 {
			orderBy = "order by" + strings.Join(sorts, " , ")
		} else {
			orderBy = "order by id"
		}
		if len(wheres) > 0 {
			where = " where " + strings.Join(wheres, " and ")
		}
		totalSql = "select ROW_NUMBER() OVER(Order by Id desc) AS RowNumber from " + tableName + "" + where
		sql += "SELECT  * FROM (SELECT " + selects + ", ROW_NUMBER() OVER(" + orderBy + " ) AS RowNumber from  " + tableName + "" + where +
			" ) as b where RowNumber between @offset and @limit"
	}

	if err := g.DB().Raw(sql, page.Query).Scan(datas).Error; err != nil {
		return 0, err
	}

	if err := g.DB().Raw(totalSql).Scan(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

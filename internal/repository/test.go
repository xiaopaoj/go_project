package repository

import (
	"go_project/internal/model"
)

func (d *repository) Find() (res model.TTable, err error) {
	//rows, err := d.db.Query("select id, name, create_time, utime from t_table limit 1")
	//if err != nil {
	//	panic(err)
	//	return res, err
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err = rows.Scan(&res.ID, &res.Name, &res.CreateTime, &res.Utime)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(res)
	//}
	if err = d.orm.Select("id, name").Where("id > 0").Limit(1).Find(&res).Error;
	err != nil {
		return
	}

	//d.orm.Raw("select id, name, create_time, utime from t_table limit 1").Scan(&res)
	return res, nil
}

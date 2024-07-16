package dao

import (
	"log"
	"wy-go-blog/models"
)

func GetAllCategory() ([]models.Category, error) {
	query, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
		return nil, err
	}
	var categorys []models.Category
	for query.Next() {
		var category models.Category
		row := query.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if row != nil {
			log.Println("GetAllCategory 取值出错：", row)
			return nil, row
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

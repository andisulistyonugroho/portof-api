package controller

import (
	"portof-api/model"

	"github.com/gofiber/fiber/v3"
)

func (cdh *DuckHandler) ListPage(c fiber.Ctx) error {

	rows, err := cdh.DB.Query(`SELECT id,title,parent_id FROM pages WHERE is_active = true`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var p []model.Page

	for rows.Next() {
		each := model.Page{}
		err := rows.Scan(&each.ID, &each.Title, &each.ParentId)
		if err != nil {
			return err
		}

		p = append(p, each)
	}

	return c.JSON(p)
}

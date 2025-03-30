package controller

import (
	"fmt"
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

	return c.JSON(fiber.Map{
		"data": p,
	})
}

func (cdh *DuckHandler) CreatePage(c fiber.Ctx) error {
	type PageRequest struct {
		Title    string `json:"title"`
		ParentId int    `json:"parent_id"`
	}

	req := new(PageRequest)

	if err := c.Bind().JSON(req); err != nil {
		return fiber.ErrBadRequest
	}

	_, err := cdh.DB.Exec(`INSERT INTO pages (title, parent_id) VALUES (?, ?)`, req.Title, req.ParentId)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}

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

	return c.JSON(fiber.Map{
		"data": p,
	})
}

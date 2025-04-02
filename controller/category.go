package controller

import (
	"fmt"
	"portof-api/model"

	"github.com/gofiber/fiber/v2"
)

func (cdh *DuckHandler) ListCategory(c *fiber.Ctx) error {
	rows, err := cdh.DB.Query(`SELECT id,page_id,title,parent_id FROM categories WHERE is_active = true`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var cat []model.Category

	for rows.Next() {
		each := model.Category{}
		err := rows.Scan(&each.ID, &each.PageId, &each.Title, &each.ParentId)
		if err != nil {
			return err
		}

		cat = append(cat, each)
	}

	return c.JSON(fiber.Map{
		"data": cat,
	})
}

func (cdh *DuckHandler) CreateCategory(c *fiber.Ctx) error {
	type TheRequest struct {
		Title    string `json:"title"`
		PageId   int    `json:"page_id"`
		ParentId int    `json:"parent_id"`
	}

	req := new(TheRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	_, err := cdh.DB.Exec(`INSERT INTO categories (title, page_id, parent_id) VALUES (?, ?, ?)`, req.Title, req.PageId, req.ParentId)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}

	rows, err := cdh.DB.Query(`SELECT id,page_id,title,parent_id FROM categories WHERE is_active = true`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var p []model.Category

	for rows.Next() {
		each := model.Category{}
		err := rows.Scan(&each.ID, &each.PageId, &each.Title, &each.ParentId)
		if err != nil {
			return err
		}

		p = append(p, each)
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}

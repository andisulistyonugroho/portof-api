package controller

import (
	"fmt"
	"portof-api/model"

	"github.com/gofiber/fiber/v2"
)

func (cdh *DuckHandler) ListPortfolio(c *fiber.Ctx) error {
	// page_id 1 = portfolio
	rows, err := cdh.DB.Query(`SELECT id,title,short_desc,the_body,display_picture_url,is_active,
	created_at FROM post WHERE is_active = true AND page_id = 1`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var p []model.Post

	for rows.Next() {
		each := model.Post{}
		err := rows.Scan(&each.ID, &each.PageId, &each.Title, &each.ShortDesc, &each.TheBody,
			&each.DisplayPictureUrl, &each.IsActive, &each.CreatedAt)
		if err != nil {
			return err
		}

		p = append(p, each)
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}

func (cdh *DuckHandler) ListBlog(c *fiber.Ctx) error {
	// page_id 2 = portfolio
	rows, err := cdh.DB.Query(`SELECT id,title,short_desc,the_body,display_picture_url,is_active,
	created_at FROM post WHERE is_active = true AND page_id = 2`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var p []model.Post

	for rows.Next() {
		each := model.Post{}
		err := rows.Scan(&each.ID, &each.PageId, &each.Title, &each.ShortDesc, &each.TheBody,
			&each.DisplayPictureUrl, &each.IsActive, &each.CreatedAt)
		if err != nil {
			return err
		}

		p = append(p, each)
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}

func (cdh *DuckHandler) CreatePost(c *fiber.Ctx) error {
	type Request struct {
		Title             string `json:"title"`
		CategoryId        int    `json:"category_id"`
		PageId            int    `json:"page_id"`
		ShortDesc         string `json:"short_desc"`
		TheBody           string `json:"the_body"`
		DisplayPictureUrl string `json:"display_picture_url"`
	}

	req := new(Request)

	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	_, err := cdh.DB.Exec(`INSERT INTO post (page_id,category_id,title,short_desc,the_body,display_picture_url)
	VALUES (?,?,?,?,?,?)`, req.PageId, req.CategoryId, req.Title, req.ShortDesc, req.TheBody, req.DisplayPictureUrl)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}

	rows, err := cdh.DB.Query(`SELECT id,title,created_at FROM pages WHERE is_active = true AND page_id = ?`, req.PageId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var p []model.Post

	for rows.Next() {
		each := model.Post{}
		err := rows.Scan(&each.ID, &each.Title, &each.CreatedAt)
		if err != nil {
			return err
		}

		p = append(p, each)
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}

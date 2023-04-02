package handler

// delivery, controller

import (
	"challange-8/helper"
	"challange-8/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) Addbook(c *gin.Context) {
	in := model.Books{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.Addbook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.GetBookByID(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h HttpServer) GetBooks(c *gin.Context) {
	var data []model.Books
	res, err := h.app.GetBooks(data)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Books{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in.ID = idInt
	// call service
	res, err := h.app.UpdateBook(in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.OkWithMessage(c, "Successfully deleted Book")
}

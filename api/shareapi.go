package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

const (
	image = "image"
	video = "video"
	other = "other"
)

func (obj *API) getGroupMessages(c *fiber.Ctx) error {
	id := obj.getGroupIDFromToken(c)
	messagesDTO := obj.service.GetGroupMessages(id)
	return c.JSON(messagesDTO)
}

func (obj *API) postMessage(c *fiber.Ctx) error {
	var messageInDTO model.MessageInDTO
	if err := c.BodyParser(&messageInDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	gid := obj.getGroupIDFromToken(c)
	uid := obj.getUserIDFromToken(c)
	obj.service.CreateMessage(gid, uid, messageInDTO)
	return c.SendStatus(fiber.StatusCreated)
}

func (obj *API) postImage(c *fiber.Ctx) error {
	data, err := c.FormFile(image)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	file, err := data.Open()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	gid := obj.getGroupIDFromToken(c)
	uid := obj.getUserIDFromToken(c)
	obj.service.CreateImage(gid, uid, data.Filename, file)
	return c.SendStatus(fiber.StatusCreated)
}

func (obj *API) postVideo(c *fiber.Ctx) error {
	data, err := c.FormFile(video)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	file, err := data.Open()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	gid := obj.getGroupIDFromToken(c)
	uid := obj.getUserIDFromToken(c)
	obj.service.CreateVideo(gid, uid, data.Filename, file)
	return c.SendStatus(fiber.StatusCreated)
}

func (obj *API) postOther(c *fiber.Ctx) error {
	data, err := c.FormFile(other)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	file, err := data.Open()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	gid := obj.getGroupIDFromToken(c)
	uid := obj.getUserIDFromToken(c)
	obj.service.CreateOther(gid, uid, data.Filename, file)
	return c.SendStatus(fiber.StatusCreated)
}

func (obj *API) deletemessage(c *fiber.Ctx) error {
	id := obj.getIDFromParam(c)
	uid := obj.getUserIDFromToken(c)
	obj.service.DeleteMessage(id, uid)
	return c.SendStatus(fiber.StatusNoContent)
}

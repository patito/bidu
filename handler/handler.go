package handler

import "github.com/patito/bidu/model"

// Handler ..
type Handler struct {
	model *model.Model
}

// New creates a new instance of Hander struct
func New(m *model.Model) *Handler {
	return &Handler{model: m}
}

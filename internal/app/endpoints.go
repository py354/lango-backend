package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebHandler struct {
	repo Repository
}

func New(repo Repository) *WebHandler {
	return &WebHandler{repo: repo}
}

type settingsRequest struct {
	Token               string `json:"token"`
	EducationMode       int    `json:"education_mode"`
	EducationComplexity int    `json:"education_complexity"`
}

func (wh *WebHandler) SetSettings(c *gin.Context) {
	var s settingsRequest
	if err := c.BindJSON(&s); err != nil {
		return
	}

	err := wh.repo.SetUserSettings(context.Background(), s.Token, s.EducationMode, s.EducationComplexity)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, s)
}

type getLessonsRequest struct {
	Token string `json:"token"`
}

func (wh *WebHandler) GetUserLessons(c *gin.Context) {
	var s getLessonsRequest
	if err := c.BindJSON(&s); err != nil {
		return
	}

	lessons, err := wh.repo.GetUserLessons(context.Background(), s.Token)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, lessons)
}

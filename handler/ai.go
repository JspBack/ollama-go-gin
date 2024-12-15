package handler

import (
	"ai-thing/schemas"
	"ai-thing/utils"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	HTTPClient *http.Client
	Logger     *slog.Logger
}

func NewHandler() *Handler {
	return &Handler{
		HTTPClient: &http.Client{},
		Logger: slog.New(
			slog.Default().Handler(),
		),
	}
}

// chat func
//
//	@Summary		Chat with AI
//	@Description	Chat with AI
//	@Tags			AI
//	@Accept			json
//	@Produce		json
//	@Param			message	formData	string	true	"Message"
//	@Param			model	formData	string	true	"Model"
//	@Param			stream	formData	bool	true	"Stream"
//	@Success		200		{string}	string				"ok"
//	@Router			/api/chat	[post]
func (h *Handler) Chat(c *gin.Context) {
	if err := utils.ValidateForms(c, "message", "model", "stream"); err != nil {
		h.Logger.Error("Failed to validate forms", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ur := schemas.AiRequest{
		Model:  c.PostForm("model"),
		Stream: c.PostForm("stream") == "true",
		Messages: []schemas.AiMessage{
			{
				Role:    "user",
				Content: c.PostForm("message"),
			},
		},
	}

	resp, err := utils.DoRequest(h.HTTPClient, ur)
	if err != nil {
		h.Logger.Error("Failed to do chat request", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})
}

// image func
//
//	@Summary		Image with AI
//	@Description	You have to send as multipart/form-data
//	@Tags			AI
//	@Accept			json
//	@Produce		json
//	@Param			message	formData	string	true	"Message"
//	@Param			image	formData	file	true	"Image file"
//	@Param			model	formData	string	true	"Model"
//	@Param			stream	formData	bool	true	"Stream"
//	@Success		200		{string}	string	"ok"
//	@Router			/api/image	[post]
func (h *Handler) Image(c *gin.Context) {
	if err := utils.ValidateForms(c, "message", "model", "stream"); err != nil {
		h.Logger.Error("Failed to validate forms", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		h.Logger.Error("Failed to get image", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := image.Open()
	if err != nil {
		h.Logger.Error("Failed to open image", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer r.Close()

	imageBytes, err := io.ReadAll(r)
	if err != nil {
		h.Logger.Error("Failed to read image", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ur := schemas.AiRequest{
		Model:  c.PostForm("model"),
		Stream: c.PostForm("stream") == "true",
		Messages: []schemas.AiMessage{
			{
				Role:    "user",
				Content: c.PostForm("message"),
				Image:   imageBytes,
			},
		},
	}

	resp, err := utils.DoRequest(h.HTTPClient, ur)
	if err != nil {
		h.Logger.Error("Failed to do image request", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})
}

// health func
//
//	@Summary		Health check
//	@Description	Health check
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	string	"ok"
//	@Router			/api/health	[get]
func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

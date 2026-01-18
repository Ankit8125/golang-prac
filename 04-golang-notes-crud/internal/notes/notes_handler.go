package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler is going to hold all the dependencies that are needed by my HTTP endpoints.
// Repo talkes to DB. Handler talks to DB via Repo.
type Handler struct {
	repo *Repo
}

func NewHandler (repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote (c *gin.Context) {
	// per req. object from gin. It will contain request info, params, body helpers, response helpers, etc
	// c.req, c.param (this "c" is going to hold every information)
	var req CreateNoteRequest

	// ShouldBindJSON: It is going to parse our JSON into request struck && It is going to validate our struct based on "binding:required" that we have passed; in 'CreateNoteRequest'
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}
  
	now := time.Now().UTC()

	note := Note {
		ID: primitive.NewObjectID(),
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := h.repo.Create(c.Request.Context(), note) // c.Request.Context() = parent context
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create Note here!",
		})
		return 
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) ListNotes (c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch all notes.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})

	
}
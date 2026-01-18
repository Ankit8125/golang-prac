package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (h* Handler) GetNoteByID (c *gin.Context) {

	idStr := c.Param("id")

	// Convert 24-char hex string into a mongo ObjectID Type
	objId, err := primitive.ObjectIDFromHex(idStr)
	// Why ?
	// String vs. BSON: In your URL, the ID is just a string ("notes/65a..."). But inside MongoDB, the _id field is stored as a special binary type called ObjectID.
	// Type Safety: MongoDB drivers won't find the document if you search for a String "65a..." against an ObjectID field. You must convert the string to an ObjectID first.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID.",
		})
		return
	}

	note, err := h.repo.GetByID(c.Request.Context(), objId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Note not found for that given ID.",
			})
			return 
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch the note.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note": note,
	})
}
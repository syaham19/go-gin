package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

// respondError mengirim pesan error dengan format seragam {"error": "..."}
func respondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// validationMessage mengubah error binding menjadi pesan yang ramah
func validationMessage(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) && len(ve) > 0 {
		fe := ve[0]
		switch fe.Tag() {
		case "required":
			return fe.Field() + " wajib diisi"
		case "min":
			return fe.Field() + " minimal " + fe.Param() + " karakter"
		case "email":
			return fe.Field() + " harus berupa email yang valid"
		}
		return fe.Field() + " tidak valid"
	}
	return "Format JSON tidak valid"
}

// parseID mengambil path param :id dan mengubahnya menjadi angka
func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, http.StatusBadRequest, "ID harus berupa angka")
		return 0, false
	}
	return id, true
}

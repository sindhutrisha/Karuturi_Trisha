package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sindhutrisha/Karuturi_Trisha/karuturi_trisha/pkg/rest/server/models"
	"github.com/sindhutrisha/Karuturi_Trisha/karuturi_trisha/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type KaruturiTrishaController struct {
	karuturiTrishaService *services.KaruturiTrishaService
}

func NewKaruturiTrishaController() (*KaruturiTrishaController, error) {
	karuturiTrishaService, err := services.NewKaruturiTrishaService()
	if err != nil {
		return nil, err
	}
	return &KaruturiTrishaController{
		karuturiTrishaService: karuturiTrishaService,
	}, nil
}

func (karuturiTrishaController *KaruturiTrishaController) CreateKaruturiTrisha(context *gin.Context) {
	// validate input
	var input models.KaruturiTrisha
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiTrisha creation
	if _, err := karuturiTrishaController.karuturiTrishaService.CreateKaruturiTrisha(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "KaruturiTrisha created successfully"})
}

func (karuturiTrishaController *KaruturiTrishaController) UpdateKaruturiTrisha(context *gin.Context) {
	// validate input
	var input models.KaruturiTrisha
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiTrisha update
	if _, err := karuturiTrishaController.karuturiTrishaService.UpdateKaruturiTrisha(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "KaruturiTrisha updated successfully"})
}

func (karuturiTrishaController *KaruturiTrishaController) FetchKaruturiTrisha(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiTrisha fetching
	karuturiTrisha, err := karuturiTrishaController.karuturiTrishaService.GetKaruturiTrisha(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, karuturiTrisha)
}

func (karuturiTrishaController *KaruturiTrishaController) DeleteKaruturiTrisha(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiTrisha deletion
	if err := karuturiTrishaController.karuturiTrishaService.DeleteKaruturiTrisha(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "KaruturiTrisha deleted successfully",
	})
}

func (karuturiTrishaController *KaruturiTrishaController) ListKaruturiTrishas(context *gin.Context) {
	// trigger all karuturiTrishas fetching
	karuturiTrishas, err := karuturiTrishaController.karuturiTrishaService.ListKaruturiTrishas()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, karuturiTrishas)
}

func (*KaruturiTrishaController) PatchKaruturiTrisha(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*KaruturiTrishaController) OptionsKaruturiTrisha(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*KaruturiTrishaController) HeadKaruturiTrisha(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}

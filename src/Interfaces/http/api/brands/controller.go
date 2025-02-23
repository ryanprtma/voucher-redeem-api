package brands

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"voucher-redeem-api/src/Applications/services"
	commons "voucher-redeem-api/src/Commons"
	"voucher-redeem-api/src/Commons/exceptions"
)

type handler struct {
	svs        services.Service
	translator *exceptions.DomainErrorTranslator
}

func newHandler(svc services.Service, translator *exceptions.DomainErrorTranslator) handler {
	return handler{
		svs:        svc,
		translator: translator,
	}
}

func (h handler) createBrand(c *gin.Context) {
	var request createBrandRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand, err := h.svs.CreateNewBrand(request.Name)

	if err != nil {
		err = h.translator.Translate(err)
		var clientErr *exceptions.ClientError
		var invariantErr *exceptions.InvariantError
		if errors.As(err, &clientErr) {
			commons.WriteError(c, clientErr.Message, clientErr.StatusCode)
		} else if errors.As(err, &invariantErr) {
			commons.WriteError(c, invariantErr.Message, invariantErr.StatusCode)
		} else {
			c.JSON(http.StatusInternalServerError, " Internal Server Error")
		}
		return
	}

	data := brandResponses{
		ID:        brand.ID,
		Name:      brand.Name,
		CreatedAt: brand.CreatedAt,
		UpdatedAt: brand.UpdatedAt,
	}

	commons.WriteSuccess(c, data, http.StatusCreated)
}

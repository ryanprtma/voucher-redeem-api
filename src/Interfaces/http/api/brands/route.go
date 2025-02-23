package brands

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"voucher-redeem-api/src/Applications/services"
	"voucher-redeem-api/src/Commons/exceptions"
	"voucher-redeem-api/src/Infrastructures/repository"
)

func SetupRoutesWith(route *gin.RouterGroup, db *sqlx.DB) {
	repo := repository.NewPostgresDB(db)
	service := services.NewBrandService(repo)
	domainErrorTranslator := exceptions.NewDomainErrorTranslator()
	controller := newHandler(service, domainErrorTranslator)

	v1 := route.Group("/v1/brand")
	{
		v1.POST("", controller.createBrand)
	}
}

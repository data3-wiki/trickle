package service

import (
	"net/http"
	"strings"

	"github.com/dereference-xyz/trickle/model"
	"github.com/dereference-xyz/trickle/store"
	"github.com/dereference-xyz/trickle/swagger"
	"github.com/gin-gonic/gin"
)

type Service struct {
	accountStore *store.AccountStore
	programType  *model.ProgramType
}

func NewService(accountStore *store.AccountStore, programType *model.ProgramType) *Service {
	return &Service{
		accountStore: accountStore,
		programType:  programType,
	}
}

func (srv *Service) Router() *gin.Engine {
	router := gin.Default()
	api := router.Group("api")
	v1 := api.Group("v1")
	{
		v1.GET("swagger/spec.json", srv.v1SwaggerSpec)
		solana := v1.Group("solana")
		{
			account := solana.Group("account")
			account.GET("read/:accountType", srv.v1SolanaAccountRead)
		}
	}
	router.Static("/swagger", "./js/swagger-ui/dist")
	return router
}

func errorToJSON(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func sendErrorResponse(c *gin.Context, err error) {
	switch err.(type) {
	case *model.InputValidationError:
		c.JSON(http.StatusBadRequest, errorToJSON(err))
	default:
		// TODO: Log error.
		c.Status(http.StatusInternalServerError)
	}
}

func (srv *Service) v1SolanaAccountRead(c *gin.Context) {
	typeName := c.Param("accountType")

	accountType, exists := srv.programType.AccountType(typeName)
	if !exists {
		candidates := []string{}
		for _, acc := range srv.programType.AccountTypes {
			candidates = append(candidates, acc.Name)
		}
		sendErrorResponse(c, model.NewInputValidationError(
			"Account type '%s' does not exist. Please choose one of the following: %s",
			typeName,
			strings.Join(candidates, ",")))
		return
	}

	predicates := make(map[string]interface{})
	for _, propertyType := range accountType.PropertyTypes {
		value, exists := c.GetQuery(propertyType.Name)
		if exists {
			// TODO: Convert based on DataType.
			predicates[propertyType.Name] = value
		}
	}

	accounts, err := srv.accountStore.Read(accountType, predicates)
	if err != nil {
		sendErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

func (srv *Service) v1SwaggerSpec(c *gin.Context) {
	spec, err := swagger.Generate(srv.programType)
	if err != nil {
		sendErrorResponse(c, err)
		return
	}

	serialized, err := spec.MarshalJSON()
	if err != nil {
		sendErrorResponse(c, err)
		return
	}

	c.Data(http.StatusOK, "application/json", serialized)
}

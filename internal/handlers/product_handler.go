package handlers

import (
	"net/http"
	"pharmacy-store/internal/domain/entities"
	"pharmacy-store/internal/domain/services"
	"pharmacy-store/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := h.Service.CreateProduct(c.Request.Context(), product); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusCreated, product)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	product, err := h.Service.GetProductByID(c.Request.Context(), id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID = id

	if err := h.Service.UpdateProduct(c.Request.Context(), product); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	if err := h.Service.DeleteProduct(c.Request.Context(), id); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	filter := c.Query("filter")
	sort := c.Query("sort")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	products, err := h.Service.ListProducts(c.Request.Context(), filter, sort, limit, offset)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, products)
}

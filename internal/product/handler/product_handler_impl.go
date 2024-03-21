package handler

import (
	"math"

	"github.com/dafailyasa/learn-golang-template/internal/product/model"
	"github.com/dafailyasa/learn-golang-template/internal/product/service"
	"github.com/dafailyasa/learn-golang-template/pkg/token"
	"github.com/dafailyasa/learn-golang-template/pkg/validator"
	util "github.com/dafailyasa/learn-golang-template/utils"
	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	ProductService service.ProductService
	Validator      validator.ValidatorApplication
}

func NewProductHandler(productService service.ProductService, validate validator.ValidatorApplication) productHandler {
	return productHandler{
		ProductService: productService,
		Validator:      validate,
	}
}

func (h *productHandler) CreateProduct(ctx *fiber.Ctx) error {
	body := new(model.ProductCreateRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	if err := h.Validator.ValidateStruct(body); len(err) > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(util.ApiResponse{Errors: err})
	}

	authLocals := ctx.Locals("auth")
	auth := authLocals.(*token.CustomClaim)

	if err := h.ProductService.Create(body, auth.Email); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON("OK")
}

func (h *productHandler) Search(ctx *fiber.Ctx) error {
	params := &model.ProductSearchParams{
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
		Search: ctx.Query("search", ""),
	}
	authLocals := ctx.Locals("auth")
	auth := authLocals.(*token.CustomClaim)

	products, total, err := h.ProductService.Search(auth.Email, params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	meta := util.PageMetadata{
		Page:      params.Page,
		Size:      params.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(params.Size))),
	}

	response := util.PaginatioResponse{
		Data: products,
		Meta: meta,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

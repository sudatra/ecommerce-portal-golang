package handlers

import (
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/dto"
	"go-ecommerce/internal/repository"
	"go-ecommerce/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoute(rh *rest.RestHandler) {
	app := rh.App;
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}
	handler := UserHandler{
		svc: svc,
	}

	publicRoutes := app.Group("/users");
	privateRoutes := app.Group("/users/private", rh.Auth.Authorize);

	// public endpoints
	publicRoutes.Post("/register", handler.Register);
	publicRoutes.Post("/login", handler.Login);

	// private endpoints
	privateRoutes.Get("/verify", handler.GetVerificationCode);
	privateRoutes.Post("/verify", handler.Verify);
	privateRoutes.Post("/profile", handler.CreateProfile);
	privateRoutes.Get("/profile", handler.GetProfile);

	privateRoutes.Post("/cart", handler.AddToCart);
	privateRoutes.Get("/cart", handler.GetCart);
	privateRoutes.Get("/order", handler.GetOrders);
	privateRoutes.Get("/order/:id", handler.GetOrder);

	privateRoutes.Post("/become-seller", handler.BecomeSeller);
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Provide Valid Input",
		})
	}

	token, err := h.svc.Signup(user);

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Unable to Register",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token": token,
	})
}



func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Provide Valid Input",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password);

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unable to login",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token": token,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify code get",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx);
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"user": user,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create order",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
	})
}
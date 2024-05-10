package http

import (
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	v  *validator.Validate
	uc UseCase
}

func NewHandler(uc UseCase) user.Handler {
	return &handler{
		uc: uc,
		v:  validator.New(),
	}
}

func (h *handler) UserSignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req SignUpRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		err := h.uc.CreateUser(ctx.Context(), req.Login, req.Password)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) UserSignIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Cookies("access-token")
		if token != "" {
			err := h.uc.DeleteSession(ctx.Context(), token)
			if err != nil {
				return err.ToHTTPError(ctx)
			}
		}

		var req SignInRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		err := h.uc.CheckPassword(ctx.Context(), req.Login, req.Password)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		t, expireAt, err := h.uc.CreateSession(ctx.Context(), req.Login)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		ctx.Cookie(&fiber.Cookie{
			Name:    "access-token",
			Value:   t,
			Expires: expireAt,
		})

		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) UserSignOut() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Cookies("access-token")
		if token == "" {
			return de.ErrUnauthorized.ToHTTPError(ctx)
		}

		err := h.uc.DeleteSession(ctx.Context(), token)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		ctx.ClearCookie("access-token")
		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) GetUserInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID, ok := (ctx.Locals("userID")).(int)
		if !ok {
			return de.ErrInvalidUserID.ToHTTPError(ctx)
		}

		login, roleID, err := h.uc.GetUserLoginAndRoleID(ctx.Context(), userID)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		return ctx.Status(fiber.StatusOK).JSON(GetUserInfoResponse{
			Name:   login,
			TypeID: roleID,
		})
	}
}

func (h *handler) AuthMW() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Cookies("access-token")
		if token == "" {
			return de.ErrUnauthorized.ToHTTPError(ctx)
		}

		userID, err := h.uc.GetUserIDFromSession(ctx.Context(), token)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		ctx.Locals("userID", userID)
		return ctx.Next()
	}
}

func (h *handler) RoleMW(role user.Role) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID, ok := (ctx.Locals("userID")).(int)
		if !ok {
			return de.ErrInvalidUserID.ToHTTPError(ctx)
		}

		roleMatch, err := h.uc.CheckUserRole(ctx.Context(), userID, int(role))
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		if !roleMatch {
			return de.ErrForbidden.ToHTTPError(ctx)
		}

		return ctx.Next()
	}
}

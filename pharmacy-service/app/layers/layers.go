package layers

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/delivery/http"
	useruc "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"
)

type Adapters struct {
	UserPGAdp useruc.PGAdapter
}

type UseCases struct {
	UserUC http.UseCase
}

type Handlers struct {
	UserH user.Handler
}

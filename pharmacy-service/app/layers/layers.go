package layers

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine"
	http2 "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/delivery/http"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/delivery/http"
	useruc "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"
)

type Adapters struct {
	UserPGAdp         useruc.PGAdapter
	UserRedisAdp      useruc.RedisAdapter
	MedicinePGAdapter usecase.PGAdapter
}

type UseCases struct {
	UserUC     http.UseCase
	MedicineUC http2.UseCase
}

type Handlers struct {
	UserH     user.Handler
	MedicineH medicine.Handler
}

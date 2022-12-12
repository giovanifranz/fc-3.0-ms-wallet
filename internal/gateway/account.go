package gateway

import "github.com.br/giovanifranz/fc-3.0-ms-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}

package gateway

import "github.com.br/giovanifranz/fc-3.0-ms-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

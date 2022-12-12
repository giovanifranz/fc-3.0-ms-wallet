package gateway

import "github.com.br/giovanifranz/fc-3.0-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

package proxy

import "GoProxyService/internal/entity"

type (
	// Для service
	Service interface {
		SaveData(body []byte) error
	}

	// Для storage
	Storage interface {
		SaveData(data []entity.ProxyData) error
	}
)

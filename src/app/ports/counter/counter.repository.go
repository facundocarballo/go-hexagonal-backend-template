package ports

import "github.com/facundocarballo/go-hexagonal-backend-template/src/app/entities"

type CounterRepository interface {
	GetActualCounterValue() entities.Counter
}

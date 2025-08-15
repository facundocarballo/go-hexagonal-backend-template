package usecases

import (
	"github.com/facundocarballo/go-hexagonal-backend-template/src/app/entities"
	ports "github.com/facundocarballo/go-hexagonal-backend-template/src/app/ports/counter"
)

func HealthUseCase(
	counterRepository ports.CounterRepository,
) entities.Counter {
	counter := counterRepository.GetActualCounterValue()

	return counter
}

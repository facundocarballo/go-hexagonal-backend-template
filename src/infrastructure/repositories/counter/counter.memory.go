package repositories

import (
	"os"
	"strconv"
	"strings"

	"github.com/facundocarballo/go-hexagonal-backend-template/src/app/entities"
)

type CounterMemory struct{}

func (counterMemory CounterMemory) GetActualCounterValue() entities.Counter {
	filename := "counter.txt"

	data, _ := os.ReadFile(filename)
	n, _ := strconv.Atoi(strings.TrimSpace(string(data)))

	n++

	os.WriteFile(filename, []byte(strconv.Itoa(n)), 0644)

	return entities.Counter{
		Value: n,
	}
}

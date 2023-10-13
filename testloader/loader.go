package testloader

import (
	"github.com/lechefer/gonkey/models"
)

type LoaderInterface interface {
	Load() ([]models.TestInterface, error)
}

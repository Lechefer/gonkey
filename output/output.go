package output

import (
	"github.com/lechefer/gonkey/models"
)

type OutputInterface interface {
	Process(models.TestInterface, *models.Result) error
}

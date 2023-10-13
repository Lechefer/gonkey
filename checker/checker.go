package checker

import "github.com/lechefer/gonkey/models"

type CheckerInterface interface {
	Check(models.TestInterface, *models.Result) ([]error, error)
}

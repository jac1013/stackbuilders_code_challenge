package useCases

import . "stackbuilders_pizza/business/models"

type StatisticOperations interface {
	Compute() Statistic
}

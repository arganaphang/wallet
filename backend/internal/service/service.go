package service

type Service struct {
	Transaction ITransactionService
	Category    ICategoryService
}

package repository

type Repository struct {
	Transaction ITransactionRepository
	Category    ICategoryRepository
}

package handler

type Handler struct {
	Transaction ITransactionHandler
	Category    ICategoryHandler
}

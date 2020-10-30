package handler

type (
	// Request data
	CreateTransferRequest struct {
		ID        string
		PayerID   string
		PayeeID   string
		Value     int64
		CreatedAt string
	}
)

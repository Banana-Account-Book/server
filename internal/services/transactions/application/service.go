package application

import (
	appError "banana-account-book.com/internal/libs/app-error"
	transaction "banana-account-book.com/internal/services/transactions/domain"
	"banana-account-book.com/internal/services/transactions/dto"
	"banana-account-book.com/internal/services/transactions/infrastructure"
	"github.com/google/uuid"
)

type TransactionService struct {
	transactionRepository infrastructure.TransactionRepository
}

func NewTransactionService(transactionRepository infrastructure.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
	}
}

func (s *TransactionService) Add(userId, accountBookId uuid.UUID, args dto.CreateTransactionRequest) error {
	transaction, err := transaction.New(transaction.TransactionDetails{
		UserId:        userId,
		AccountBookId: accountBookId,
		Title:         args.Title,
		Description:   args.Description,
		PeriodStartOn: args.PeriodStartOn,
		PeriodEndOn:   args.PeriodEndOn,
		Type:          args.Type,
		RepeatType:    args.RepeatType,
		Amount:        args.Amount,
	})
	if err != nil {
		return appError.Wrap(err)
	}

	return s.transactionRepository.Save(nil, transaction)
}

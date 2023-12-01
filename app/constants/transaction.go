package constants

const (
	// Transaction Success messages
	TransactionCreatedSuccessfully = "Transaction created successfully"
	TransactionUpdatedSuccessfully = "Transaction updated successfully"
	TransactionDeletedSuccessfully = "Transaction deleted successfully"
	// TransactionRetrievedSuccessfully = "Transaction retrieved successfully"
	// TransactionsListedSuccessfully   = "Transactions listed successfully"
	TransactionCreditedSuccessfully = "Transaction credited successfully"
	TransactionDebitedSuccessfully  = "Transaction debited successfully"

	// Transaction Error messages\
	FailedToCreditAmount         = "Failed to credit amount"
	FailedToDebitAmount          = "Failed to debit amount"
	FailedToGetTransaction       = "Failed to get transaction details"
	FailedToGetBalance           = "Failed to get user balance"
	FailedToListTransactions     = "Failed to list transactions"
	InvalidTransactionRequest    = "Invalid transaction request"
	InvalidCreditRequest         = "Invalid credit request"
	InvalidDebitRequest          = "Invalid dedit request"
	FailedToCreateTransaction    = "Failed to create transaction"
	FailedToUpdateTransaction    = "Failed to update transaction"
	FailedToDeleteTransaction    = "Failed to delete transaction"
	FailedToRetrieveBalance      = "Failed to retrieve balance"
	FailedToRetrieveTransaction  = "Failed to retrieve transaction"
	FailedToRetrieveTransactions = "Failed to retrieve transactions"
	InsufficientFunds            = "Insufficient funds for the debit"
	TransactionNotFound          = "Transaction not found"
)

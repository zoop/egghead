package services

import (
	"egghead/app/common"
	"egghead/app/models"
)

type ProductPaginateResult struct {
	common.PaginatedResult
	Products []models.Products `json:"products"`
}

type UserPaginateResult struct {
	common.PaginatedResult
	Users []models.Users `json:"users"`
}

type TransactionPaginateResult struct {
	common.PaginatedResult
	Transactions []models.TransactionHistory `json:"transactions"`
}

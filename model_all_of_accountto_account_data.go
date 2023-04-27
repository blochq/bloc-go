/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type AllOfAccounttoAccountData struct {
	Reference string `json:"reference"`
	AccountId string `json:"account_id"`
	RecipientAccountNumber string `json:"recipient_account_number"`
	RecipientAccountName string `json:"recipient_account_name"`
	Amount int32 `json:"amount"`
	Narration string `json:"narration"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	CreateAt string `json:"create_at"`
}
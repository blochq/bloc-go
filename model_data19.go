/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type Data19 struct {
	Id string `json:"id"`
	CustomerId string `json:"customer_id"`
	OrganisationId string `json:"organisation_id"`
	Environment string `json:"environment"`
	CardId string `json:"card_id"`
	TransactionId string `json:"transaction_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	TransactionAmount int32 `json:"transaction_amount"`
	Reason string `json:"reason"`
	Explanation string `json:"explanation"`
	Status string `json:"status"`
	Metadata *AllOfData19Metadata `json:"metadata"`
}

/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type AllTransactions struct {
	Success bool `json:"success"`
	Data []Data50 `json:"data"`
	Message string `json:"message"`
	Metadata *AllOfAllTransactionsMetadata `json:"metadata"`
}

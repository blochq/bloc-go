/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type TransfersBulkBody struct {
	AccountId string `json:"account_id"`
	BulkList []BulkList `json:"bulk_list"`
}

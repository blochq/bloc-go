/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type Getaccounts struct {
	Success bool `json:"success"`
	Data []Data `json:"data"`
	Message string `json:"message"`
	Metadata *AllOfGetaccountsMetadata `json:"metadata"`
}

/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type UpgradedCustomerT3 struct {
	Success bool `json:"success"`
	Data *AllOfUpgradedCustomerT3Data `json:"data"`
	Message string `json:"message"`
}
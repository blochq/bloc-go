/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type AllOfGetCardSecureDataData struct {
	Type_ string `json:"type"`
	Brand string `json:"brand"`
	Number string `json:"number"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear string `json:"expiryYear"`
	Cvv string `json:"cvv"`
	Cvv2 string `json:"cvv2"`
	DefaultPin string `json:"defaultPin"`
}

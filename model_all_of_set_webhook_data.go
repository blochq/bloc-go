/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type AllOfSetWebhookData struct {
	OrganizationId string `json:"organization_id"`
	Environment string `json:"environment"`
	Url string `json:"url"`
	PortalUrl string `json:"portal_url"`
	SecretKey string `json:"secret_key"`
}

/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type UpdateCardDisputeRequest struct {
	Reason string `json:"reason"`
	Explanation string `json:"explanation"`
	MetaData *interface{} `json:"meta_data"`
}

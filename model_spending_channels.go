/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type SpendingChannels struct {
	Mobile bool `json:"mobile"`
	Web bool `json:"web"`
	Pos bool `json:"pos"`
	Atm bool `json:"atm"`
}
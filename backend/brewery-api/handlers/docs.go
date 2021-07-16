// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package handlers

import "brewery/api/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// a product returned in the response
// swagger:response productResponse
type productResponseWrapper struct {
	// a product in the system
	// in: body
	Body data.Product
}

// swagger:parameters deleteProduct updateProduct listProduct
type productIDParametersWrapper struct {
	// The id of the product
	// in: path
	// required: true
	ID int `json:"id"`
}

/*
 * API for interacting with GameStonks
 *
 * GameStonks API 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// StockApiService is a service that implents the logic for the StockApiServicer
// This service should implement the business logic for every endpoint for the StockApi API. 
// Include any external packages or services that will be required by this service.
type StockApiService struct {
}

// NewStockApiService creates a default api service
func NewStockApiService() StockApiServicer {
	return &StockApiService{}
}

// AddNewComment - Add a comment to a stock
func (s *StockApiService) AddNewComment(ctx context.Context, stockTicker string) (ImplResponse, error) {
	// TODO - update AddNewComment with the required logic for this service method.
	// Add api_stock_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddNewComment method not implemented")
}

// AddNewVote - Add a vote to a stock
func (s *StockApiService) AddNewVote(ctx context.Context, stockTicker string) (ImplResponse, error) {
	// TODO - update AddNewVote with the required logic for this service method.
	// Add api_stock_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddNewVote method not implemented")
}

// GetAllStocks - Gets all stocks on the platform within a date range
func (s *StockApiService) GetAllStocks(ctx context.Context, since string, until string) (ImplResponse, error) {
	// TODO - update GetAllStocks with the required logic for this service method.
	// Add api_stock_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAllStocks method not implemented")
}

// GetIndividualStock - Returns a stock&#39;s vote count and comments
func (s *StockApiService) GetIndividualStock(ctx context.Context, stockTicker string) (ImplResponse, error) {
	// TODO - update GetIndividualStock with the required logic for this service method.
	// Add api_stock_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]int32{}) or use other options such as http.Ok ...
	//return Response(200, map[string]int32{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetIndividualStock method not implemented")
}

// InsertIndividualStock - Create a new stock ticker for tracking
func (s *StockApiService) InsertIndividualStock(ctx context.Context, stockTicker string) (ImplResponse, error) {
	// TODO - update InsertIndividualStock with the required logic for this service method.
	// Add api_stock_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("InsertIndividualStock method not implemented")
}


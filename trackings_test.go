package trackingmore

import (
	"context"
	"errors"
	"testing"
)

func TestCreateTracking(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validParams := CreateTrackingParams{
		TrackingNumber: "9400111899562537683155",
		CourierCode:    "usps",
	}

	response, err := client.CreateTracking(context.Background(), validParams)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Error("Expected a response, got nil")
	}

	var _, ok = response.Data.(*Tracking)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	trackingNumberEmptyParams := CreateTrackingParams{
		TrackingNumber: "",
		CourierCode:    "usps",
	}

	response, err = client.CreateTracking(context.Background(), trackingNumberEmptyParams)
	if err == nil || !errors.Is(err, ErrMissingTrackingNumber) {
		t.Errorf("Expected ErrMissingTrackingNumber error, got %v", err)
	}
	if response != nil {
		t.Error("Expected nil response, got a response")
	}

	courierCodeEmptyParams := CreateTrackingParams{
		TrackingNumber: "1234227890",
		CourierCode:    "",
	}

	response, err = client.CreateTracking(context.Background(), courierCodeEmptyParams)
	if err == nil || !errors.Is(err, ErrMissingCourierCode) {
		t.Errorf("Expected ErrMissingTrackingNumber error, got %v", err)
	}
	if response != nil {
		t.Error("Expected nil response, got a response")
	}
}

func TestGetTrackingResults(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validParams := GetTrackingResultsParams{
		TrackingNumbers: "9400111899562537683155",
		CourierCode:     "usps",
	}

	response, err := client.GetTrackingResults(context.Background(), validParams)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Error("Expected a response, got nil")
	}

	var _, ok = response.Data.(*[]Tracking)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

}

func TestBatchCreateTrackings(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validParams := []CreateTrackingParams{
		{
			TrackingNumber: "123456789",
			CourierCode:    "usps",
		},
		{
			TrackingNumber: "987654321",
			CourierCode:    "fedex",
		},
	}

	response, err := client.BatchCreateTrackings(context.Background(), validParams)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if response == nil {
		t.Error("Expected a response, but got nil")
	}

	var _, ok = response.Data.(*BatchResults)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	invalidParams := make([]CreateTrackingParams, 41)
	_, err = client.BatchCreateTrackings(context.Background(), invalidParams)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if !errors.Is(err, ErrMaxTrackingNumbersExceeded) {
		t.Errorf("Expected error: %v, but got: %v", ErrMaxTrackingNumbersExceeded, err)
	}

	missingTrackingNumber := []CreateTrackingParams{
		{
			TrackingNumber: "",
			CourierCode:    "usps",
		},
	}
	_, err = client.BatchCreateTrackings(context.Background(), missingTrackingNumber)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if !errors.Is(err, ErrMissingTrackingNumber) {
		t.Errorf("Expected error: %v, but got: %v", ErrMissingTrackingNumber, err)
	}

	missingCourierCode := []CreateTrackingParams{
		{
			TrackingNumber: "123456789",
			CourierCode:    "",
		},
	}
	_, err = client.BatchCreateTrackings(context.Background(), missingCourierCode)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if !errors.Is(err, ErrMissingCourierCode) {
		t.Errorf("Expected error: %v, but got: %v", ErrMissingCourierCode, err)
	}
}

func TestUpdateTrackingByID(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validID := "9a1d6e2a33bb9c867bf9614b575ef3fe"
	validParams := UpdateTrackingParams{
		CustomerName: "New name",
		Note:         "New tests order note",
	}

	response, err := client.UpdateTrackingByID(context.Background(), validID, validParams)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if response == nil {
		t.Error("Expected a response, but got nil")
	}

	var _, ok = response.Data.(*UpdateAfterResult)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	emptyID := ""
	_, err = client.UpdateTrackingByID(context.Background(), emptyID, validParams)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if !errors.Is(err, ErrEmptyId) {
		t.Errorf("Expected error: %v, but got: %v", ErrEmptyId, err)
	}
}

func TestDeleteTrackingByID(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validID := "9a1d6e2a33bb9c867bf9614b575ef3fe"

	response, err := client.DeleteTrackingByID(context.Background(), validID)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if response == nil {
		t.Error("Expected a response, but got nil")
	}

	var _, ok = response.Data.(*Tracking)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	emptyID := ""
	_, err = client.DeleteTrackingByID(context.Background(), emptyID)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if !errors.Is(err, ErrEmptyId) {
		t.Errorf("Expected error: %v, but got: %v", ErrEmptyId, err)
	}
}

func TestRetrackTrackingByID(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validID := "9a1d6e2a33bb9c867bf9614b575ef3fe"

	response, err := client.RetrackTrackingByID(context.Background(), validID)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if response == nil {
		t.Error("Expected a response, but got nil")
	}

	var _, ok = response.Data.(*Tracking)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	emptyID := ""
	_, err = client.RetrackTrackingByID(context.Background(), emptyID)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if !errors.Is(err, ErrEmptyId) {
		t.Errorf("Expected error: %v, but got: %v", ErrEmptyId, err)
	}
}

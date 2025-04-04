package trackingmore

import "errors"

var (
	// ErrEmptyAPIKey api key is empty
	ErrEmptyAPIKey = errors.New("API Key is missing")

	// ErrMissingTrackingNumber miss tracking number
	ErrMissingTrackingNumber = errors.New("Tracking number cannot be empty")

	// ErrMissingCourierCode miss courier code
	ErrMissingCourierCode = errors.New("Courier Code cannot be empty")

	// ErrMissingAwbNumber awb number empty
	ErrMissingAwbNumber = errors.New("Awb number cannot be empty")

	// ErrMaxTrackingNumbersExceeded more than 40 tracking numbers
	ErrMaxTrackingNumbersExceeded = errors.New("Max. 40 tracking numbers create in one call")

	// ErrEmptyId id empty
	ErrEmptyId = errors.New("Id cannot be empty")

	// ErrInvalidAirWaybillFormat air waybill number invalid
	ErrInvalidAirWaybillFormat = errors.New("The air waybill number format is invalid")
)

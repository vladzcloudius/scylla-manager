// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigFdInitialValueMsReader is a Reader for the FindConfigFdInitialValueMs structure.
type FindConfigFdInitialValueMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigFdInitialValueMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigFdInitialValueMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigFdInitialValueMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigFdInitialValueMsOK creates a FindConfigFdInitialValueMsOK with default headers values
func NewFindConfigFdInitialValueMsOK() *FindConfigFdInitialValueMsOK {
	return &FindConfigFdInitialValueMsOK{}
}

/*FindConfigFdInitialValueMsOK handles this case with default header values.

Config value
*/
type FindConfigFdInitialValueMsOK struct {
	Payload int64
}

func (o *FindConfigFdInitialValueMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigFdInitialValueMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigFdInitialValueMsDefault creates a FindConfigFdInitialValueMsDefault with default headers values
func NewFindConfigFdInitialValueMsDefault(code int) *FindConfigFdInitialValueMsDefault {
	return &FindConfigFdInitialValueMsDefault{
		_statusCode: code,
	}
}

/*FindConfigFdInitialValueMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigFdInitialValueMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config fd initial value ms default response
func (o *FindConfigFdInitialValueMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigFdInitialValueMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigFdInitialValueMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigFdInitialValueMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

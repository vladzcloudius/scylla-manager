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

// FindConfigAPIUIDirReader is a Reader for the FindConfigAPIUIDir structure.
type FindConfigAPIUIDirReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAPIUIDirReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAPIUIDirOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAPIUIDirDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAPIUIDirOK creates a FindConfigAPIUIDirOK with default headers values
func NewFindConfigAPIUIDirOK() *FindConfigAPIUIDirOK {
	return &FindConfigAPIUIDirOK{}
}

/*FindConfigAPIUIDirOK handles this case with default header values.

Config value
*/
type FindConfigAPIUIDirOK struct {
	Payload string
}

func (o *FindConfigAPIUIDirOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigAPIUIDirOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAPIUIDirDefault creates a FindConfigAPIUIDirDefault with default headers values
func NewFindConfigAPIUIDirDefault(code int) *FindConfigAPIUIDirDefault {
	return &FindConfigAPIUIDirDefault{
		_statusCode: code,
	}
}

/*FindConfigAPIUIDirDefault handles this case with default header values.

unexpected error
*/
type FindConfigAPIUIDirDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config api ui dir default response
func (o *FindConfigAPIUIDirDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAPIUIDirDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAPIUIDirDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAPIUIDirDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

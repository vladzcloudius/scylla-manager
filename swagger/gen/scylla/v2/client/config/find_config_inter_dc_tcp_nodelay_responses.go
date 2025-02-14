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

// FindConfigInterDcTCPNodelayReader is a Reader for the FindConfigInterDcTCPNodelay structure.
type FindConfigInterDcTCPNodelayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigInterDcTCPNodelayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigInterDcTCPNodelayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigInterDcTCPNodelayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigInterDcTCPNodelayOK creates a FindConfigInterDcTCPNodelayOK with default headers values
func NewFindConfigInterDcTCPNodelayOK() *FindConfigInterDcTCPNodelayOK {
	return &FindConfigInterDcTCPNodelayOK{}
}

/*FindConfigInterDcTCPNodelayOK handles this case with default header values.

Config value
*/
type FindConfigInterDcTCPNodelayOK struct {
	Payload bool
}

func (o *FindConfigInterDcTCPNodelayOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigInterDcTCPNodelayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigInterDcTCPNodelayDefault creates a FindConfigInterDcTCPNodelayDefault with default headers values
func NewFindConfigInterDcTCPNodelayDefault(code int) *FindConfigInterDcTCPNodelayDefault {
	return &FindConfigInterDcTCPNodelayDefault{
		_statusCode: code,
	}
}

/*FindConfigInterDcTCPNodelayDefault handles this case with default header values.

unexpected error
*/
type FindConfigInterDcTCPNodelayDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config inter dc tcp nodelay default response
func (o *FindConfigInterDcTCPNodelayDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigInterDcTCPNodelayDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigInterDcTCPNodelayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigInterDcTCPNodelayDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

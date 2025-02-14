// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// HintedHandoffMetricsCreateHintByAddrGetReader is a Reader for the HintedHandoffMetricsCreateHintByAddrGet structure.
type HintedHandoffMetricsCreateHintByAddrGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HintedHandoffMetricsCreateHintByAddrGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHintedHandoffMetricsCreateHintByAddrGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHintedHandoffMetricsCreateHintByAddrGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetOK creates a HintedHandoffMetricsCreateHintByAddrGetOK with default headers values
func NewHintedHandoffMetricsCreateHintByAddrGetOK() *HintedHandoffMetricsCreateHintByAddrGetOK {
	return &HintedHandoffMetricsCreateHintByAddrGetOK{}
}

/*HintedHandoffMetricsCreateHintByAddrGetOK handles this case with default header values.

HintedHandoffMetricsCreateHintByAddrGetOK hinted handoff metrics create hint by addr get o k
*/
type HintedHandoffMetricsCreateHintByAddrGetOK struct {
	Payload int32
}

func (o *HintedHandoffMetricsCreateHintByAddrGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *HintedHandoffMetricsCreateHintByAddrGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHintedHandoffMetricsCreateHintByAddrGetDefault creates a HintedHandoffMetricsCreateHintByAddrGetDefault with default headers values
func NewHintedHandoffMetricsCreateHintByAddrGetDefault(code int) *HintedHandoffMetricsCreateHintByAddrGetDefault {
	return &HintedHandoffMetricsCreateHintByAddrGetDefault{
		_statusCode: code,
	}
}

/*HintedHandoffMetricsCreateHintByAddrGetDefault handles this case with default header values.

internal server error
*/
type HintedHandoffMetricsCreateHintByAddrGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the hinted handoff metrics create hint by addr get default response
func (o *HintedHandoffMetricsCreateHintByAddrGetDefault) Code() int {
	return o._statusCode
}

func (o *HintedHandoffMetricsCreateHintByAddrGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *HintedHandoffMetricsCreateHintByAddrGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *HintedHandoffMetricsCreateHintByAddrGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

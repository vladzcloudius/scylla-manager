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

// StorageServiceIncrementalBackupsPostReader is a Reader for the StorageServiceIncrementalBackupsPost structure.
type StorageServiceIncrementalBackupsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceIncrementalBackupsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceIncrementalBackupsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceIncrementalBackupsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceIncrementalBackupsPostOK creates a StorageServiceIncrementalBackupsPostOK with default headers values
func NewStorageServiceIncrementalBackupsPostOK() *StorageServiceIncrementalBackupsPostOK {
	return &StorageServiceIncrementalBackupsPostOK{}
}

/*StorageServiceIncrementalBackupsPostOK handles this case with default header values.

StorageServiceIncrementalBackupsPostOK storage service incremental backups post o k
*/
type StorageServiceIncrementalBackupsPostOK struct {
}

func (o *StorageServiceIncrementalBackupsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceIncrementalBackupsPostDefault creates a StorageServiceIncrementalBackupsPostDefault with default headers values
func NewStorageServiceIncrementalBackupsPostDefault(code int) *StorageServiceIncrementalBackupsPostDefault {
	return &StorageServiceIncrementalBackupsPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceIncrementalBackupsPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceIncrementalBackupsPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service incremental backups post default response
func (o *StorageServiceIncrementalBackupsPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceIncrementalBackupsPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceIncrementalBackupsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceIncrementalBackupsPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

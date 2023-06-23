// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSilenceReader is a Reader for the DeleteSilence structure.
type DeleteSilenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSilenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSilenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteSilenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSilenceOK creates a DeleteSilenceOK with default headers values
func NewDeleteSilenceOK() *DeleteSilenceOK {
	return &DeleteSilenceOK{}
}

/*
DeleteSilenceOK describes a response with status code 200, with default header values.

Delete silence response
*/
type DeleteSilenceOK struct {
}

// IsSuccess returns true when this delete silence o k response has a 2xx status code
func (o *DeleteSilenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete silence o k response has a 3xx status code
func (o *DeleteSilenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete silence o k response has a 4xx status code
func (o *DeleteSilenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete silence o k response has a 5xx status code
func (o *DeleteSilenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete silence o k response a status code equal to that given
func (o *DeleteSilenceOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteSilenceOK) Error() string {
	return fmt.Sprintf("[DELETE /silence/{silenceID}][%d] deleteSilenceOK ", 200)
}

func (o *DeleteSilenceOK) String() string {
	return fmt.Sprintf("[DELETE /silence/{silenceID}][%d] deleteSilenceOK ", 200)
}

func (o *DeleteSilenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSilenceInternalServerError creates a DeleteSilenceInternalServerError with default headers values
func NewDeleteSilenceInternalServerError() *DeleteSilenceInternalServerError {
	return &DeleteSilenceInternalServerError{}
}

/*
DeleteSilenceInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteSilenceInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this delete silence internal server error response has a 2xx status code
func (o *DeleteSilenceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete silence internal server error response has a 3xx status code
func (o *DeleteSilenceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete silence internal server error response has a 4xx status code
func (o *DeleteSilenceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete silence internal server error response has a 5xx status code
func (o *DeleteSilenceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete silence internal server error response a status code equal to that given
func (o *DeleteSilenceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteSilenceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /silence/{silenceID}][%d] deleteSilenceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSilenceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /silence/{silenceID}][%d] deleteSilenceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSilenceInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *DeleteSilenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
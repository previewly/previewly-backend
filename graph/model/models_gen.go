// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ImageProcess struct {
	ImageID   int                `json:"imageId"`
	Processes []*OneImageProcess `json:"processes"`
	Status    Status             `json:"status"`
	Error     *string            `json:"error,omitempty"`
}

type ImageProcessOption struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

type ImageProcessOptionInput struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

type ImageProcessesInput struct {
	Type    ImageProcessType           `json:"type"`
	Options []*ImageProcessOptionInput `json:"options"`
}

type Mutation struct {
}

type OneImageProcess struct {
	Type    ImageProcessType      `json:"type"`
	Options []*ImageProcessOption `json:"options"`
}

type PreviewData struct {
	ID     int     `json:"id"`
	URL    string  `json:"url"`
	Status Status  `json:"status"`
	Image  string  `json:"image"`
	Error  *string `json:"error,omitempty"`
	Title  *string `json:"title,omitempty"`
}

type Query struct {
}

type UploadImageStatus struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Status Status  `json:"status"`
	Error  *string `json:"error,omitempty"`
}

type ImageProcessType string

const (
	ImageProcessTypeResize ImageProcessType = "resize"
)

var AllImageProcessType = []ImageProcessType{
	ImageProcessTypeResize,
}

func (e ImageProcessType) IsValid() bool {
	switch e {
	case ImageProcessTypeResize:
		return true
	}
	return false
}

func (e ImageProcessType) String() string {
	return string(e)
}

func (e *ImageProcessType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ImageProcessType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ImageProcessType", str)
	}
	return nil
}

func (e ImageProcessType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
	StatusPending Status = "pending"
)

var AllStatus = []Status{
	StatusSuccess,
	StatusError,
	StatusPending,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusSuccess, StatusError, StatusPending:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

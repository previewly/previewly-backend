// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"wsw/backend/ent/errorresult"
	"wsw/backend/ent/imageprocess"
	"wsw/backend/ent/schema"
	"wsw/backend/ent/stat"
	"wsw/backend/ent/uploadimage"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	errorresultFields := schema.ErrorResult{}.Fields()
	_ = errorresultFields
	// errorresultDescCreatedAt is the schema descriptor for created_at field.
	errorresultDescCreatedAt := errorresultFields[0].Descriptor()
	// errorresult.DefaultCreatedAt holds the default value on creation for the created_at field.
	errorresult.DefaultCreatedAt = errorresultDescCreatedAt.Default.(func() time.Time)
	imageprocessFields := schema.ImageProcess{}.Fields()
	_ = imageprocessFields
	// imageprocessDescCreatedAt is the schema descriptor for created_at field.
	imageprocessDescCreatedAt := imageprocessFields[3].Descriptor()
	// imageprocess.DefaultCreatedAt holds the default value on creation for the created_at field.
	imageprocess.DefaultCreatedAt = imageprocessDescCreatedAt.Default.(func() time.Time)
	// imageprocessDescUpdatedAt is the schema descriptor for updated_at field.
	imageprocessDescUpdatedAt := imageprocessFields[4].Descriptor()
	// imageprocess.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	imageprocess.DefaultUpdatedAt = imageprocessDescUpdatedAt.Default.(func() time.Time)
	// imageprocess.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	imageprocess.UpdateDefaultUpdatedAt = imageprocessDescUpdatedAt.UpdateDefault.(func() time.Time)
	statFields := schema.Stat{}.Fields()
	_ = statFields
	// statDescCreatedAt is the schema descriptor for created_at field.
	statDescCreatedAt := statFields[0].Descriptor()
	// stat.DefaultCreatedAt holds the default value on creation for the created_at field.
	stat.DefaultCreatedAt = statDescCreatedAt.Default.(func() time.Time)
	uploadimageFields := schema.UploadImage{}.Fields()
	_ = uploadimageFields
	// uploadimageDescFilename is the schema descriptor for filename field.
	uploadimageDescFilename := uploadimageFields[0].Descriptor()
	// uploadimage.FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	uploadimage.FilenameValidator = uploadimageDescFilename.Validators[0].(func(string) error)
	// uploadimageDescDestinationPath is the schema descriptor for destination_path field.
	uploadimageDescDestinationPath := uploadimageFields[1].Descriptor()
	// uploadimage.DestinationPathValidator is a validator for the "destination_path" field. It is called by the builders before save.
	uploadimage.DestinationPathValidator = uploadimageDescDestinationPath.Validators[0].(func(string) error)
	// uploadimageDescOriginalFilename is the schema descriptor for original_filename field.
	uploadimageDescOriginalFilename := uploadimageFields[2].Descriptor()
	// uploadimage.OriginalFilenameValidator is a validator for the "original_filename" field. It is called by the builders before save.
	uploadimage.OriginalFilenameValidator = uploadimageDescOriginalFilename.Validators[0].(func(string) error)
	// uploadimageDescType is the schema descriptor for type field.
	uploadimageDescType := uploadimageFields[3].Descriptor()
	// uploadimage.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	uploadimage.TypeValidator = uploadimageDescType.Validators[0].(func(string) error)
}

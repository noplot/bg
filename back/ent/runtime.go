// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/biouser"
	"back/ent/enttiktokuser"
	"back/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	biouserFields := schema.BioUser{}.Fields()
	_ = biouserFields
	// biouserDescName is the schema descriptor for name field.
	biouserDescName := biouserFields[0].Descriptor()
	// biouser.DefaultName holds the default value on creation for the name field.
	biouser.DefaultName = biouserDescName.Default.(string)
	enttiktokuserFields := schema.EntTikTokUser{}.Fields()
	_ = enttiktokuserFields
	// enttiktokuserDescName is the schema descriptor for name field.
	enttiktokuserDescName := enttiktokuserFields[0].Descriptor()
	// enttiktokuser.DefaultName holds the default value on creation for the name field.
	enttiktokuser.DefaultName = enttiktokuserDescName.Default.(string)
}

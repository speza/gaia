package resourcehelper

import (
	"errors"
	"fmt"

	"github.com/gaia-pipeline/gaia"
	"gopkg.in/yaml.v2"
)

type Marshaller interface {
	Marshal(in interface{}) ([]byte, error)
	Unmarshal(bytes []byte, out interface{}) error
}

type marshaller struct{}

func NewMarshaller() *marshaller {
	return &marshaller{}
}

// Marshal is currently just a wrapper around the yaml.Marshal func. Since we have the concept of a versioned
// policy specifications its likely we will need to handle these here in the future.
func (f marshaller) Marshal(in interface{}) ([]byte, error) {
	bts, err := yaml.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal policy: %w", err)
	}

	return bts, nil
}

// Unmarshal is a wrapper around the yaml.Unmarshal func that allows us to unmarshal and validate the specification.
func (f marshaller) Unmarshal(bytes []byte, out interface{}) error {
	if err := yaml.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("failed to unmarhsal policy: %w", err)
	}

	// Check for version mismatches.
	if v1, ok := out.(*gaia.AuthPolicyResourceV1); ok {
		if v1.Type != "authorization.policy" {
			return errors.New("failed to unmarshal policy type: not an authorization.policy type")
		}
		if v1.Version != "v1" {
			return errors.New("version does not match struct AuthPolicyResourceV1")
		}
		// TODO: add more validation?
		return nil
	}

	return errors.New("policy specification struct not found")
}
package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEndPointFoi(t *testing.T) {
	// arrange
	ep := CreateFeaturesOfInterestEndpoint("http://www.nu.nl")
	ep.Name = "yo"

	// assert
	assert.True(t, ep != nil)
	assert.True(t, ep.GetName() == "yo")
}

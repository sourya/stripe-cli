package resource

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewResourceCmd(t *testing.T) {
	parentCmd := &cobra.Command{Annotations: make(map[string]string)}

	rc := NewResourceCmd(parentCmd, "foo")

	assert.Equal(t, "foo", rc.Name)
	assert.True(t, parentCmd.HasSubCommands())
	assert.Equal(t, "resource", parentCmd.Annotations["foo"])
}

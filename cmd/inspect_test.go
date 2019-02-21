package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestInspect(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	err := inspectFunc(&emptyCmd, []string{"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"})
	if err != nil {
		t.Error(err)
	}
}

package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestInspect(t *testing.T) {
	setCLIConfig()
	emptyCmd := cobra.Command{}
	inspectFunc(&emptyCmd, []string{"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"})
}

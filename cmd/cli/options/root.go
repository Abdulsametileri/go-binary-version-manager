package options

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const verbosityUsage = `PanicLevel (0),
Fatal (1) .
Error (2) .
Warn (3) 
Info (4) 
Debug (5)
Trace (6)`

type RootOptions struct {
	Verbosity uint32 // --verbosity=3 | -v=3
}

func (ro *RootOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Uint32VarP(&ro.Verbosity, "verbosity", "v", uint32(log.WarnLevel), verbosityUsage)
}

package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/thecodeteam/gocsi/csi"
)

var supportedVersCmd = &cobra.Command{
	Use:     "supportedversions",
	Aliases: []string{"ver", "version", "gets"},
	Short:   `invokes the rpc "GetSupportedVersions"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx, cancel := context.WithTimeout(root.ctx, root.timeout)
		defer cancel()

		rep, err := identity.client.GetSupportedVersions(
			ctx,
			&csi.GetSupportedVersionsRequest{})
		if err != nil {
			return err
		}

		return root.tpl.Execute(os.Stdout, rep)
	},
}

func init() {
	identityCmd.AddCommand(supportedVersCmd)
}

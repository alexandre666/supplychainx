package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/ltacker/supplychainx/x/scx/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	scxTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	scxTxCmd.AddCommand(flags.PostCommands(
		GetCmdAppendOrganization(cdc),
	)...)

	return scxTxCmd
}

// Transactions sent by authority to append new organization
func GetCmdAppendOrganization(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "append-organization [organization-addr] [organization-name] [flags]",
		Short: "Append a new organization to interact with the ledger",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// The authority is the sender
			accAddress := cliCtx.GetFromAddress()
			if accAddress.Empty() {
				return fmt.Errorf("Account address empty")
			}
			authorityAddress := sdk.ValAddress(accAddress)

			// Get organization address
			organizationAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			name := args[1]
			description, _ := cmd.Flags().GetString(FlagOrganizationDescription)

			organization := types.NewOrganization(organizationAddr, name, description)

			msg := types.NewMsgAppendOrganization(organization, authorityAddress)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(FlagSetOrganizationDescriptionCreate())

	return cmd
}

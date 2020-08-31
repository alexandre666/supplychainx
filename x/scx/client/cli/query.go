package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ltacker/supplychainx/x/scx/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group scx queries under a subcommand
	scxQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	scxQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdQueryOrganization(queryRoute, cdc),
			GetCmdQueryOrganizations(queryRoute, cdc),
			GetCmdQueryProduct(queryRoute, cdc),
			GetCmdQueryProductUnits(queryRoute, cdc),
			GetCmdQueryUnit(queryRoute, cdc),
			GetCmdQueryUnitTrace(queryRoute, cdc),
			GetCmdQueryUnitComponents(queryRoute, cdc),
		)...,
	)

	return scxQueryCmd
}

func GetCmdQueryOrganization(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "organization [organization-addr]",
		Short: "Query an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get address
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := types.NewQueryOrganizationParams(addr)

			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryOrganization), bz)
			if err != nil {
				fmt.Printf("could not resolve %s %s \n", types.QueryOrganization, addr)
				return nil
			}

			var out types.Organization
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryOrganizations(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "organizations",
		Short: "Query all organizations",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryOrganizations), nil)
			if err != nil {
				fmt.Printf("could not resolve %s \n", types.QueryOrganizations)
				return nil
			}

			var out []types.Organization
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryProduct(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "product [product-name]",
		Short: "Query product description",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get product name
			params := types.NewQueryProductParams(args[0])
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryProduct), bz)
			if err != nil {
				fmt.Printf("could not resolve %s \n", types.QueryProduct)
				return nil
			}

			var out types.Product
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryProductUnits(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "product-units [product-name]",
		Short: "Query all the units of a product",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get product name
			params := types.NewQueryProductParams(args[0])
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryProductUnits), bz)
			if err != nil {
				fmt.Printf("could not resolve %s \n", types.QueryProductUnits)
				return nil
			}

			var out []types.Unit
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryUnit(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unit [unit-reference]",
		Short: "Query unit description",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get unit reference
			params := types.NewQueryUnitParams(args[0])
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryUnit), bz)
			if err != nil {
				fmt.Printf("could not resolve %s \n", types.QueryUnit)
				return nil
			}

			var out types.Unit
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryUnitTrace(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unit-trace [unit-reference]",
		Short: "Query the history of all the holding organizations of the unit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get unit reference
			params := types.NewQueryUnitParams(args[0])
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryUnitTrace), bz)
			if err != nil {
				fmt.Printf("could not resolve %s \n", types.QueryUnitTrace)
				return nil
			}

			var out []types.Organization
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryUnitComponents(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unit-components [unit-reference]",
		Short: "Query the descriptions of all the components composing the unit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get unit reference
			params := types.NewQueryUnitParams(args[0])
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryUnitComponents), bz)
			if err != nil {
				fmt.Printf("could not resolve %s \n", types.QueryUnitComponents)
				return nil
			}

			var out []types.Unit
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

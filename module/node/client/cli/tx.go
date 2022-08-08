package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/bianjieai/spartan-cosmos/module/node/client/utils"
)

// NewCreateValidatorProposalTxCmd returns a CLI command handler for creating
// a validator create proposal governance transaction.
func NewCreateValidatorProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-validator [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a validator create proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Example:
$ %s tx gov submit-proposal create-validator proposal.json --from mykey

Where proposal.json contains:

{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"name":"node1",
		"certificate":"-----BEGIN CERTIFICATE-----\nMIIBbDCCAR4CFGX5t8OFPrR8Ul/1JoXj3CFBBrBvMAUGAytlcDBYMQswCQYDVQQG\nEwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v\ndDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMjA4MDUxMTA0Mzda\nFw0yMzA4MDUxMTA0MzdaMFkxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w\nCwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ4wDAYD\nVQQDDAVub2RlMDAqMAUGAytlcAMhACJAXK+iRqiKcapdgM3aKrP39Mwp/gGjUDPS\nWo3+qM8vMAUGAytlcANBAP0JsV3l8PW65sxgxx3dBAQbagCDimvcxVtwIIF/rWPo\ndLTSmKoX6XNtyq5O5bU8Xj8mypeWqQd0SQe65MQqvQc=\n-----END CERTIFICATE-----\n",
	    "power":"100",
	    "description":"my node1",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := utils.ParseValidatorCreateJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&proposal.Content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// NewUpdateValidatorProposalTxCmd returns a CLI command handler for creating
// a validator create proposal governance transaction.
func NewUpdateValidatorProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update-validator [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a validator update proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Example:
$ %s tx gov submit-proposal update-validator proposal.json --from mykey

Where proposal.json contains:

{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"id":"xxxxxx",
		"name":"node1",
		"certificate":"-----BEGIN CERTIFICATE-----\nMIIBbDCCAR4CFGX5t8OFPrR8Ul/1JoXj3CFBBrBvMAUGAytlcDBYMQswCQYDVQQG\nEwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v\ndDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMjA4MDUxMTA0Mzda\nFw0yMzA4MDUxMTA0MzdaMFkxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w\nCwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ4wDAYD\nVQQDDAVub2RlMDAqMAUGAytlcAMhACJAXK+iRqiKcapdgM3aKrP39Mwp/gGjUDPS\nWo3+qM8vMAUGAytlcANBAP0JsV3l8PW65sxgxx3dBAQbagCDimvcxVtwIIF/rWPo\ndLTSmKoX6XNtyq5O5bU8Xj8mypeWqQd0SQe65MQqvQc=\n-----END CERTIFICATE-----\n",
	    "power":"100",
	    "description":"my node1",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := utils.ParseValidatorUpdateJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&proposal.Content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// NewRemoveValidatorProposalTxCmd returns a CLI command handler for creating
// a validator create proposal governance transaction.
func NewRemoveValidatorProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove-validator [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a validator remove proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Example:
$ %s tx gov submit-proposal remove-validator proposal.json --from mykey

Where proposal.json contains:

{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"id":"xxxxxx",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := utils.ParseValidatorRemoveJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&proposal.Content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

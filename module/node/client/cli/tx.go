package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

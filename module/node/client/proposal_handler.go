package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/bianjieai/spartan-cosmos/module/node/client/cli"
)

var (
	CreateValidatorProposalHandler = govclient.NewProposalHandler(cli.NewCreateValidatorProposalTxCmd, nil)
	UpdateValidatorProposalHandler = govclient.NewProposalHandler(cli.NewUpdateValidatorProposalTxCmd, nil)
	RemoveValidatorProposalHandler = govclient.NewProposalHandler(cli.NewRemoveValidatorProposalTxCmd, nil)
)

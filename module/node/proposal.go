package node

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

const (
	ProposalTypeCreateValidator string = "CreateValidator"
	ProposalTypeUpdateValidator string = "UpdateValidator"
	ProposalTypeRemoveValidator string = "RemoveValidator"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateValidator)
	govtypes.RegisterProposalType(ProposalTypeUpdateValidator)
	govtypes.RegisterProposalType(ProposalTypeRemoveValidator)
}

// Implements Proposal Interface
var (
	_ govtypes.Content = &CreateValidatorProposal{}
	_ govtypes.Content = &UpdateValidatorProposal{}
	_ govtypes.Content = &RemoveValidatorProposal{}
)

func (sup *CreateValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *CreateValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *CreateValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *CreateValidatorProposal) ProposalType() string   { return ProposalTypeCreateValidator }
func (sup *CreateValidatorProposal) ValidateBasic() error {
	//TODO
	return nil
}

func (sup *UpdateValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *UpdateValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *UpdateValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *UpdateValidatorProposal) ProposalType() string   { return ProposalTypeUpdateValidator }
func (sup *UpdateValidatorProposal) ValidateBasic() error {
	//TODO
	return nil
}

func (sup *RemoveValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *RemoveValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *RemoveValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *RemoveValidatorProposal) ProposalType() string   { return ProposalTypeRemoveValidator }
func (sup *RemoveValidatorProposal) ValidateBasic() error {
	//TODO
	return nil
}

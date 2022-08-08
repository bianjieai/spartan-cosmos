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

func (sup *CreateValidatorProposal) ToMsgCreateValidator() types.MsgCreateValidator {
	return types.MsgCreateValidator{
		Name:        sup.Name,
		Certificate: sup.Certificate,
		Power:       sup.Power,
		Description: sup.Description,
		Operator:    sup.Operator,
	}
}

func (sup *CreateValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *CreateValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *CreateValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *CreateValidatorProposal) ProposalType() string   { return ProposalTypeCreateValidator }
func (sup *CreateValidatorProposal) ValidateBasic() error {
	return sup.ToMsgCreateValidator().ValidateBasic()
}

func (sup *UpdateValidatorProposal) ToMsgUpdateValidator() types.MsgUpdateValidator {
	return types.MsgUpdateValidator{
		Id:          sup.Id,
		Name:        sup.Name,
		Certificate: sup.Certificate,
		Power:       sup.Power,
		Description: sup.Description,
		Operator:    sup.Operator,
	}
}
func (sup *UpdateValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *UpdateValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *UpdateValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *UpdateValidatorProposal) ProposalType() string   { return ProposalTypeUpdateValidator }
func (sup *UpdateValidatorProposal) ValidateBasic() error {
	return sup.ToMsgUpdateValidator().ValidateBasic()
}

func (sup *RemoveValidatorProposal) ToMsgRemoveValidator() types.MsgRemoveValidator {
	return types.MsgRemoveValidator{
		Id:       sup.Id,
		Operator: sup.Operator,
	}
}
func (sup *RemoveValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *RemoveValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *RemoveValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *RemoveValidatorProposal) ProposalType() string   { return ProposalTypeRemoveValidator }
func (sup *RemoveValidatorProposal) ValidateBasic() error {
	return sup.ToMsgRemoveValidator().ValidateBasic()
}

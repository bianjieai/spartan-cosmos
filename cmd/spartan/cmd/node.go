package cmd

import (
	"fmt"
	"github.com/gogo/protobuf/proto"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/p2p"

	"github.com/cosmos/cosmos-sdk/client"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/version"
)

type NodeInfo struct {
	NodeID    string             `json:"node_id"`
	PubKey    cryptotypes.PubKey `json:"pub_key"`
	Signature []byte             `json:"signature"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}

// NodeCmd creates a main CLI command
func NodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Tool for helping with debugging your application",
		RunE:  client.ValidateCmd,
	}

	cmd.AddCommand(SignCmd())
	cmd.AddCommand(VerifyCmd())
	return cmd
}

func SignCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sign",
		Short: "sign with nodeID and export signature",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			cfg := serverCtx.Config

			nodeKey, err := p2p.LoadNodeKey(cfg.NodeKeyFile())
			if err != nil {
				return err
			}

			pubKey := nodeKey.PubKey()
			nodeID := string(p2p.PubKeyToID(pubKey))

			sdkPK, err := cryptocodec.FromTmPubKeyInterface(pubKey)
			if err != nil {
				return err
			}

			sigByte, err := nodeKey.PrivKey.Sign([]byte(nodeID))
			if err != nil {
				return err
			}

			nodeInfo := NodeInfo{
				NodeID:    nodeID,
				PubKey:    sdkPK,
				Signature: sigByte,
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			bz, err := clientCtx.Codec.MarshalInterfaceJSON(&nodeInfo)
			if err != nil {
				return err
			}
			fmt.Println(string(bz))
			return nil
		},
	}
}

func VerifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pubkey [pubkey]",
		Short: "Decode a pubkey from proto JSON",
		Long: fmt.Sprintf(`Decode a pubkey from proto JSON and display it's address.

Example:
$ %s debug pubkey '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AurroA7jvfPd1AadmmOvWM2rJSwipXfRf8yD6pLbA2DJ"}'
			`, version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

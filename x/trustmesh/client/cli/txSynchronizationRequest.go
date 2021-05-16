package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	//"strconv"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	baseledgerTypes "github.com/example/baseledger/x/baseledger/types"
	"github.com/example/baseledger/x/trustmesh/proxy"
	"github.com/example/baseledger/x/trustmesh/types"
	uuid "github.com/kthomas/go.uuid"
)

func CmdCreateSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-SynchronizationRequest [WorkgroupID] [Recipient] [WorkstepType] [BusinessObjectType] [BaseledgerBusinessObjectID] [BusinessObject] [ReferencedBaseledgerBusinessObjectID]",
		Short: "Creates a new SynchronizationRequest",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {

			var createSyncReq types.SynchronizationRequest
			createSyncReq.WorkgroupID = string(args[0])
			createSyncReq.Recipient = string(args[1])
			createSyncReq.WorkstepType = string(args[2])
			createSyncReq.BusinessObjectType = string(args[3])
			createSyncReq.BaseledgerBusinessObjectID = string(args[4])
			createSyncReq.BusinessObject = string(args[5])
			createSyncReq.ReferencedBaseledgerBusinessObjectID = string(args[6])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			payload := proxy.SynchronizeBusinessObjectCLI(createSyncReq)
			baseId, _ := uuid.NewV4()

			fmt.Printf("creator address2 %s\n", clientCtx.GetFromAddress().String())

			msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), baseId.String(), string(payload))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			fmt.Printf("msg with encrypted payload to be broadcasted %s\n", msg)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

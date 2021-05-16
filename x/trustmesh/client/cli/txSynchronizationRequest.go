package cli

import (
	"github.com/spf13/cobra"

	//"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
	// "github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/example/baseledger/x/trustmesh/proxy"
	"github.com/example/baseledger/x/trustmesh/types"
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

			// clientCtx, err := client.GetClientTxContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// msg := types.NewMsgCreateSynchronizationRequest(clientCtx.GetFromAddress().String(), string(argsWorkgroupID), string(argsRecipient), string(argsWorkstepType), string(argsBusinessObjectType), string(argsBaseledgerBusinessObjectID), string(argsBusinessObject), string(argsReferencedBaseledgerBusinessObjectID))
			// if err := msg.ValidateBasic(); err != nil {
			// 	return err
			// }

			proxy.SynchronizeBusinessObjectCLI(createSyncReq)
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

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
)

func CmdCreateSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-SynchronizationRequest [WorkgroupId] [Recipient] [WorkstepType] [BusinessObjectType] [BaseledgerBusinessObjectId] [BusinessObject] [ReferencedBaseledgerBusinessObjectId]",
		Short: "Creates a new SynchronizationRequest",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {

			createSyncReq := &types.SynchronizationRequest{
				WorkgroupId:                          string(args[0]),
				Recipient:                            string(args[1]),
				WorkstepType:                         string(args[2]),
				BusinessObjectType:                   string(args[3]),
				BaseledgerBusinessObjectId:           string(args[4]),
				BusinessObject:                       string(args[5]),
				ReferencedBaseledgerBusinessObjectId: string(args[6]),
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			payload, transactionId := proxy.SynchronizeBusinessObject(createSyncReq)

			msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), transactionId, string(payload))
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

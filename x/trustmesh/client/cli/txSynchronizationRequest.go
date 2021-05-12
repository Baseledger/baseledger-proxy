package cli

import (
	"github.com/spf13/cobra"
	//"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/example/baseledger/x/trustmesh/types"
)

func CmdCreateSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-SynchronizationRequest [WorkgroupID] [Recipient] [WorkstepType] [BusinessObjectType] [BaseledgerBusinessObjectID] [BusinessObject] [ReferencedBaseledgerBusinessObjectID]",
		Short: "Creates a new SynchronizationRequest",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsWorkgroupID := string(args[0])
			argsRecipient := string(args[1])
			argsWorkstepType := string(args[2])
			argsBusinessObjectType := string(args[3])
			argsBaseledgerBusinessObjectID := string(args[4])
			argsBusinessObject := string(args[5])
			argsReferencedBaseledgerBusinessObjectID := string(args[6])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSynchronizationRequest(clientCtx.GetFromAddress().String(), string(argsWorkgroupID), string(argsRecipient), string(argsWorkstepType), string(argsBusinessObjectType), string(argsBaseledgerBusinessObjectID), string(argsBusinessObject), string(argsReferencedBaseledgerBusinessObjectID))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

/* func CmdUpdateSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-SynchronizationRequest [id] [WorkgroupID] [Recipient] [WorkstepType] [BusinessObjectType] [BaseledgerBusinessObjectID] [BusinessObject] [ReferencedBaseledgerBusinessObjectID]",
		Short: "Update a SynchronizationRequest",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsWorkgroupID := string(args[1])
			argsRecipient := string(args[2])
			argsWorkstepType := string(args[3])
			argsBusinessObjectType := string(args[4])
			argsBaseledgerBusinessObjectID := string(args[5])
			argsBusinessObject := string(args[6])
			argsReferencedBaseledgerBusinessObjectID := string(args[7])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSynchronizationRequest(clientCtx.GetFromAddress().String(), id, string(argsWorkgroupID), string(argsRecipient), string(argsWorkstepType), string(argsBusinessObjectType), string(argsBaseledgerBusinessObjectID), string(argsBusinessObject), string(argsReferencedBaseledgerBusinessObjectID))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-SynchronizationRequest [id] [WorkgroupID] [Recipient] [WorkstepType] [BusinessObjectType] [BaseledgerBusinessObjectID] [BusinessObject] [ReferencedBaseledgerBusinessObjectID]",
		Short: "Delete a SynchronizationRequest by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSynchronizationRequest(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
} */

package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/example/baseledger/x/trustmesh/types"
)

func CmdCreateSynchronizationFeedback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-SynchronizationFeedback [Approved] [BusinessObject] [BaseledgerBusinessObjectIDOfApprovedObject] [Workgroup] [Recipient] [HashOfObjectToApprove] [OriginalBaseledgerTransactionID] [OriginalOffchainProcessMessageID] [FeedbackMessage]",
		Short: "Creates a new SynchronizationFeedback",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsApproved := string(args[0])
			argsBusinessObject := string(args[1])
			argsBaseledgerBusinessObjectIDOfApprovedObject := string(args[2])
			argsWorkgroup := string(args[3])
			argsRecipient := string(args[4])
			argsHashOfObjectToApprove := string(args[5])
			argsOriginalBaseledgerTransactionID := string(args[6])
			argsOriginalOffchainProcessMessageID := string(args[7])
			argsFeedbackMessage := string(args[8])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSynchronizationFeedback(clientCtx.GetFromAddress().String(), string(argsApproved), string(argsBusinessObject), string(argsBaseledgerBusinessObjectIDOfApprovedObject), string(argsWorkgroup), string(argsRecipient), string(argsHashOfObjectToApprove), string(argsOriginalBaseledgerTransactionID), string(argsOriginalOffchainProcessMessageID), string(argsFeedbackMessage))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateSynchronizationFeedback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-SynchronizationFeedback [id] [Approved] [BusinessObject] [BaseledgerBusinessObjectIDOfApprovedObject] [Workgroup] [Recipient] [HashOfObjectToApprove] [OriginalBaseledgerTransactionID] [OriginalOffchainProcessMessageID] [FeedbackMessage]",
		Short: "Update a SynchronizationFeedback",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsApproved := string(args[1])
			argsBusinessObject := string(args[2])
			argsBaseledgerBusinessObjectIDOfApprovedObject := string(args[3])
			argsWorkgroup := string(args[4])
			argsRecipient := string(args[5])
			argsHashOfObjectToApprove := string(args[6])
			argsOriginalBaseledgerTransactionID := string(args[7])
			argsOriginalOffchainProcessMessageID := string(args[8])
			argsFeedbackMessage := string(args[9])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSynchronizationFeedback(clientCtx.GetFromAddress().String(), id, string(argsApproved), string(argsBusinessObject), string(argsBaseledgerBusinessObjectIDOfApprovedObject), string(argsWorkgroup), string(argsRecipient), string(argsHashOfObjectToApprove), string(argsOriginalBaseledgerTransactionID), string(argsOriginalOffchainProcessMessageID), string(argsFeedbackMessage))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteSynchronizationFeedback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-SynchronizationFeedback [id] [Approved] [BusinessObject] [BaseledgerBusinessObjectIDOfApprovedObject] [Workgroup] [Recipient] [HashOfObjectToApprove] [OriginalBaseledgerTransactionID] [OriginalOffchainProcessMessageID] [FeedbackMessage]",
		Short: "Delete a SynchronizationFeedback by id",
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

			msg := types.NewMsgDeleteSynchronizationFeedback(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

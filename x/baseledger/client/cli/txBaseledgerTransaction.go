package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/unibrightio/baseledger/x/baseledger/types"
)

func CmdCreateBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-BaseledgerTransaction [baseledgerTransactionId] [payload]",
		Short: "Create a new BaseledgerTransaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBaseledgerTransactionId, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsPayload, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), argsBaseledgerTransactionId, argsPayload)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-BaseledgerTransaction [id] [baseledgerTransactionId] [payload]",
		Short: "Update a BaseledgerTransaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBaseledgerTransactionId, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsPayload, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBaseledgerTransaction(clientCtx.GetFromAddress().String(), id, argsBaseledgerTransactionId, argsPayload)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-BaseledgerTransaction [id]",
		Short: "Delete a BaseledgerTransaction by id",
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

			msg := types.NewMsgDeleteBaseledgerTransaction(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

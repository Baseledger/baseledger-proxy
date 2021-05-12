package cli

import (
	"github.com/spf13/cobra"
	//"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/example/baseledger/x/baseledger/types"
)

func CmdCreateBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-BaseledgerTransaction [baseid] [payload]",
		Short: "Creates a new BaseledgerTransaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBaseid := string(args[0])
			argsPayload := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), string(argsBaseid), string(argsPayload))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

/* func CmdUpdateBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-BaseledgerTransaction [id] [baseid] [payload]",
		Short: "Update a BaseledgerTransaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBaseid := string(args[1])
			argsPayload := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBaseledgerTransaction(clientCtx.GetFromAddress().String(), id, string(argsBaseid), string(argsPayload))
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
		Use:   "delete-BaseledgerTransaction [id] [baseid] [payload]",
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
} */

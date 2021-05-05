package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/example/baseledger/x/baseledger/types"
	"github.com/spf13/cobra"
)

func CmdListBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-BaseledgerTransaction",
		Short: "list all BaseledgerTransaction",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBaseledgerTransactionRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BaseledgerTransactionAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowBaseledgerTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-BaseledgerTransaction [id]",
		Short: "shows a BaseledgerTransaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetBaseledgerTransactionRequest{
				Id: id,
			}

			res, err := queryClient.BaseledgerTransaction(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

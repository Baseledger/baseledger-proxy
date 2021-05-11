package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/example/baseledger/x/trustmesh/types"
	"github.com/spf13/cobra"
)

func CmdListSynchronizationFeedback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-SynchronizationFeedback",
		Short: "list all SynchronizationFeedback",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSynchronizationFeedbackRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SynchronizationFeedbackAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSynchronizationFeedback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-SynchronizationFeedback [id]",
		Short: "shows a SynchronizationFeedback",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetSynchronizationFeedbackRequest{
				Id: id,
			}

			res, err := queryClient.SynchronizationFeedback(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

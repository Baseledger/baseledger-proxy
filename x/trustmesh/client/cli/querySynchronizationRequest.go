package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/example/baseledger/x/trustmesh/types"
	"github.com/spf13/cobra"
)

func CmdListSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-SynchronizationRequest",
		Short: "list all SynchronizationRequest",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSynchronizationRequestRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SynchronizationRequestAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-SynchronizationRequest [id]",
		Short: "shows a SynchronizationRequest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetSynchronizationRequestRequest{
				Id: id,
			}

			res, err := queryClient.SynchronizationRequest(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

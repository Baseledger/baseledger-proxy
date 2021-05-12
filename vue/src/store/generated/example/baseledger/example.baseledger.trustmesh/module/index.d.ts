import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgUpdateSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgCreateSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgDeleteSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgDeleteSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgUpdateSynchronizationFeedback } from "./types/trustmesh/tx";
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgCreateSynchronizationRequest: (data: MsgCreateSynchronizationRequest) => EncodeObject;
    msgUpdateSynchronizationRequest: (data: MsgUpdateSynchronizationRequest) => EncodeObject;
    msgCreateSynchronizationFeedback: (data: MsgCreateSynchronizationFeedback) => EncodeObject;
    msgDeleteSynchronizationFeedback: (data: MsgDeleteSynchronizationFeedback) => EncodeObject;
    msgDeleteSynchronizationRequest: (data: MsgDeleteSynchronizationRequest) => EncodeObject;
    msgUpdateSynchronizationFeedback: (data: MsgUpdateSynchronizationFeedback) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };

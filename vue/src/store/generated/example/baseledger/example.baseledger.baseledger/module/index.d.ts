import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateBaseledgerTransaction } from "./types/baseledger/tx";
import { MsgCreateBaseledgerTransaction } from "./types/baseledger/tx";
import { MsgDeleteBaseledgerTransaction } from "./types/baseledger/tx";
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgUpdateBaseledgerTransaction: (data: MsgUpdateBaseledgerTransaction) => EncodeObject;
    msgCreateBaseledgerTransaction: (data: MsgCreateBaseledgerTransaction) => EncodeObject;
    msgDeleteBaseledgerTransaction: (data: MsgDeleteBaseledgerTransaction) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };

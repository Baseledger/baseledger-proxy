// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgUpdateSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgDeleteSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgCreateSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgDeleteSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgUpdateSynchronizationRequest } from "./types/trustmesh/tx";
const types = [
    ["/example.baseledger.trustmesh.MsgCreateSynchronizationFeedback", MsgCreateSynchronizationFeedback],
    ["/example.baseledger.trustmesh.MsgUpdateSynchronizationFeedback", MsgUpdateSynchronizationFeedback],
    ["/example.baseledger.trustmesh.MsgDeleteSynchronizationFeedback", MsgDeleteSynchronizationFeedback],
    ["/example.baseledger.trustmesh.MsgCreateSynchronizationRequest", MsgCreateSynchronizationRequest],
    ["/example.baseledger.trustmesh.MsgDeleteSynchronizationRequest", MsgDeleteSynchronizationRequest],
    ["/example.baseledger.trustmesh.MsgUpdateSynchronizationRequest", MsgUpdateSynchronizationRequest],
];
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw new Error("wallet is required");
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee = defaultFee, memo = null }) => memo ? client.signAndBroadcast(address, msgs, fee, memo) : client.signAndBroadcast(address, msgs, fee),
        msgCreateSynchronizationFeedback: (data) => ({ typeUrl: "/example.baseledger.trustmesh.MsgCreateSynchronizationFeedback", value: data }),
        msgUpdateSynchronizationFeedback: (data) => ({ typeUrl: "/example.baseledger.trustmesh.MsgUpdateSynchronizationFeedback", value: data }),
        msgDeleteSynchronizationFeedback: (data) => ({ typeUrl: "/example.baseledger.trustmesh.MsgDeleteSynchronizationFeedback", value: data }),
        msgCreateSynchronizationRequest: (data) => ({ typeUrl: "/example.baseledger.trustmesh.MsgCreateSynchronizationRequest", value: data }),
        msgDeleteSynchronizationRequest: (data) => ({ typeUrl: "/example.baseledger.trustmesh.MsgDeleteSynchronizationRequest", value: data }),
        msgUpdateSynchronizationRequest: (data) => ({ typeUrl: "/example.baseledger.trustmesh.MsgUpdateSynchronizationRequest", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };

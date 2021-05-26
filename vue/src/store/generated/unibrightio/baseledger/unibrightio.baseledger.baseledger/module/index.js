// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteBaseledgerTransaction } from "./types/baseledger/tx";
import { MsgCreateBaseledgerTransaction } from "./types/baseledger/tx";
import { MsgUpdateBaseledgerTransaction } from "./types/baseledger/tx";
const types = [
    ["/unibrightio.baseledger.baseledger.MsgDeleteBaseledgerTransaction", MsgDeleteBaseledgerTransaction],
    ["/unibrightio.baseledger.baseledger.MsgCreateBaseledgerTransaction", MsgCreateBaseledgerTransaction],
    ["/unibrightio.baseledger.baseledger.MsgUpdateBaseledgerTransaction", MsgUpdateBaseledgerTransaction],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgDeleteBaseledgerTransaction: (data) => ({ typeUrl: "/unibrightio.baseledger.baseledger.MsgDeleteBaseledgerTransaction", value: data }),
        msgCreateBaseledgerTransaction: (data) => ({ typeUrl: "/unibrightio.baseledger.baseledger.MsgCreateBaseledgerTransaction", value: data }),
        msgUpdateBaseledgerTransaction: (data) => ({ typeUrl: "/unibrightio.baseledger.baseledger.MsgUpdateBaseledgerTransaction", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };

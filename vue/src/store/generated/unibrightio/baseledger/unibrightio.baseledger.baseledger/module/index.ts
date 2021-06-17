// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateBaseledgerTransaction } from "./types/baseledger/tx";
import { MsgCreateBaseledgerTransaction } from "./types/baseledger/tx";
import { MsgDeleteBaseledgerTransaction } from "./types/baseledger/tx";


const types = [
  ["/unibrightio.baseledger.baseledger.MsgUpdateBaseledgerTransaction", MsgUpdateBaseledgerTransaction],
  ["/unibrightio.baseledger.baseledger.MsgCreateBaseledgerTransaction", MsgCreateBaseledgerTransaction],
  ["/unibrightio.baseledger.baseledger.MsgDeleteBaseledgerTransaction", MsgDeleteBaseledgerTransaction],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateBaseledgerTransaction: (data: MsgUpdateBaseledgerTransaction): EncodeObject => ({ typeUrl: "/unibrightio.baseledger.baseledger.MsgUpdateBaseledgerTransaction", value: data }),
    msgCreateBaseledgerTransaction: (data: MsgCreateBaseledgerTransaction): EncodeObject => ({ typeUrl: "/unibrightio.baseledger.baseledger.MsgCreateBaseledgerTransaction", value: data }),
    msgDeleteBaseledgerTransaction: (data: MsgDeleteBaseledgerTransaction): EncodeObject => ({ typeUrl: "/unibrightio.baseledger.baseledger.MsgDeleteBaseledgerTransaction", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};

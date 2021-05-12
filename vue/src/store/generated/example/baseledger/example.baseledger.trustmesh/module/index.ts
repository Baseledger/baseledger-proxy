// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgUpdateSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgCreateSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgDeleteSynchronizationFeedback } from "./types/trustmesh/tx";
import { MsgDeleteSynchronizationRequest } from "./types/trustmesh/tx";
import { MsgUpdateSynchronizationFeedback } from "./types/trustmesh/tx";


const types = [
  ["/example.baseledger.trustmesh.MsgCreateSynchronizationRequest", MsgCreateSynchronizationRequest],
  ["/example.baseledger.trustmesh.MsgUpdateSynchronizationRequest", MsgUpdateSynchronizationRequest],
  ["/example.baseledger.trustmesh.MsgCreateSynchronizationFeedback", MsgCreateSynchronizationFeedback],
  ["/example.baseledger.trustmesh.MsgDeleteSynchronizationFeedback", MsgDeleteSynchronizationFeedback],
  ["/example.baseledger.trustmesh.MsgDeleteSynchronizationRequest", MsgDeleteSynchronizationRequest],
  ["/example.baseledger.trustmesh.MsgUpdateSynchronizationFeedback", MsgUpdateSynchronizationFeedback],
  
];

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
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgCreateSynchronizationRequest: (data: MsgCreateSynchronizationRequest): EncodeObject => ({ typeUrl: "/example.baseledger.trustmesh.MsgCreateSynchronizationRequest", value: data }),
    msgUpdateSynchronizationRequest: (data: MsgUpdateSynchronizationRequest): EncodeObject => ({ typeUrl: "/example.baseledger.trustmesh.MsgUpdateSynchronizationRequest", value: data }),
    msgCreateSynchronizationFeedback: (data: MsgCreateSynchronizationFeedback): EncodeObject => ({ typeUrl: "/example.baseledger.trustmesh.MsgCreateSynchronizationFeedback", value: data }),
    msgDeleteSynchronizationFeedback: (data: MsgDeleteSynchronizationFeedback): EncodeObject => ({ typeUrl: "/example.baseledger.trustmesh.MsgDeleteSynchronizationFeedback", value: data }),
    msgDeleteSynchronizationRequest: (data: MsgDeleteSynchronizationRequest): EncodeObject => ({ typeUrl: "/example.baseledger.trustmesh.MsgDeleteSynchronizationRequest", value: data }),
    msgUpdateSynchronizationFeedback: (data: MsgUpdateSynchronizationFeedback): EncodeObject => ({ typeUrl: "/example.baseledger.trustmesh.MsgUpdateSynchronizationFeedback", value: data }),
    
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

import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "unibrightio.baseledger.baseledger";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateBaseledgerTransaction {
    creator: string;
    id: string;
    baseledgerTransactionId: string;
    payload: string;
}
export interface MsgCreateBaseledgerTransactionResponse {
    id: string;
}
export interface MsgUpdateBaseledgerTransaction {
    creator: string;
    id: string;
    baseledgerTransactionId: string;
    payload: string;
}
export interface MsgUpdateBaseledgerTransactionResponse {
}
export interface MsgDeleteBaseledgerTransaction {
    creator: string;
    id: string;
}
export interface MsgDeleteBaseledgerTransactionResponse {
}
export declare const MsgCreateBaseledgerTransaction: {
    encode(message: MsgCreateBaseledgerTransaction, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateBaseledgerTransaction;
    fromJSON(object: any): MsgCreateBaseledgerTransaction;
    toJSON(message: MsgCreateBaseledgerTransaction): unknown;
    fromPartial(object: DeepPartial<MsgCreateBaseledgerTransaction>): MsgCreateBaseledgerTransaction;
};
export declare const MsgCreateBaseledgerTransactionResponse: {
    encode(message: MsgCreateBaseledgerTransactionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateBaseledgerTransactionResponse;
    fromJSON(object: any): MsgCreateBaseledgerTransactionResponse;
    toJSON(message: MsgCreateBaseledgerTransactionResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateBaseledgerTransactionResponse>): MsgCreateBaseledgerTransactionResponse;
};
export declare const MsgUpdateBaseledgerTransaction: {
    encode(message: MsgUpdateBaseledgerTransaction, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateBaseledgerTransaction;
    fromJSON(object: any): MsgUpdateBaseledgerTransaction;
    toJSON(message: MsgUpdateBaseledgerTransaction): unknown;
    fromPartial(object: DeepPartial<MsgUpdateBaseledgerTransaction>): MsgUpdateBaseledgerTransaction;
};
export declare const MsgUpdateBaseledgerTransactionResponse: {
    encode(_: MsgUpdateBaseledgerTransactionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateBaseledgerTransactionResponse;
    fromJSON(_: any): MsgUpdateBaseledgerTransactionResponse;
    toJSON(_: MsgUpdateBaseledgerTransactionResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateBaseledgerTransactionResponse>): MsgUpdateBaseledgerTransactionResponse;
};
export declare const MsgDeleteBaseledgerTransaction: {
    encode(message: MsgDeleteBaseledgerTransaction, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteBaseledgerTransaction;
    fromJSON(object: any): MsgDeleteBaseledgerTransaction;
    toJSON(message: MsgDeleteBaseledgerTransaction): unknown;
    fromPartial(object: DeepPartial<MsgDeleteBaseledgerTransaction>): MsgDeleteBaseledgerTransaction;
};
export declare const MsgDeleteBaseledgerTransactionResponse: {
    encode(_: MsgDeleteBaseledgerTransactionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteBaseledgerTransactionResponse;
    fromJSON(_: any): MsgDeleteBaseledgerTransactionResponse;
    toJSON(_: MsgDeleteBaseledgerTransactionResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteBaseledgerTransactionResponse>): MsgDeleteBaseledgerTransactionResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateBaseledgerTransaction(request: MsgCreateBaseledgerTransaction): Promise<MsgCreateBaseledgerTransactionResponse>;
    UpdateBaseledgerTransaction(request: MsgUpdateBaseledgerTransaction): Promise<MsgUpdateBaseledgerTransactionResponse>;
    DeleteBaseledgerTransaction(request: MsgDeleteBaseledgerTransaction): Promise<MsgDeleteBaseledgerTransactionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateBaseledgerTransaction(request: MsgCreateBaseledgerTransaction): Promise<MsgCreateBaseledgerTransactionResponse>;
    UpdateBaseledgerTransaction(request: MsgUpdateBaseledgerTransaction): Promise<MsgUpdateBaseledgerTransactionResponse>;
    DeleteBaseledgerTransaction(request: MsgDeleteBaseledgerTransaction): Promise<MsgDeleteBaseledgerTransactionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

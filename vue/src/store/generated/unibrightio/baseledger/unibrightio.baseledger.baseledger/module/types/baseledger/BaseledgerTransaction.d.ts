import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "unibrightio.baseledger.baseledger";
export interface BaseledgerTransaction {
    creator: string;
    id: number;
    baseledgerTransactionId: string;
    payload: string;
}
export declare const BaseledgerTransaction: {
    encode(message: BaseledgerTransaction, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BaseledgerTransaction;
    fromJSON(object: any): BaseledgerTransaction;
    toJSON(message: BaseledgerTransaction): unknown;
    fromPartial(object: DeepPartial<BaseledgerTransaction>): BaseledgerTransaction;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

import { Reader, Writer } from "protobufjs/minimal";
import { BaseledgerTransaction } from "../baseledger/BaseledgerTransaction";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "example.baseledger.baseledger";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetBaseledgerTransactionRequest {
    id: number;
}
export interface QueryGetBaseledgerTransactionResponse {
    BaseledgerTransaction: BaseledgerTransaction | undefined;
}
export interface QueryAllBaseledgerTransactionRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllBaseledgerTransactionResponse {
    BaseledgerTransaction: BaseledgerTransaction[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetBaseledgerTransactionRequest: {
    encode(message: QueryGetBaseledgerTransactionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBaseledgerTransactionRequest;
    fromJSON(object: any): QueryGetBaseledgerTransactionRequest;
    toJSON(message: QueryGetBaseledgerTransactionRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBaseledgerTransactionRequest>): QueryGetBaseledgerTransactionRequest;
};
export declare const QueryGetBaseledgerTransactionResponse: {
    encode(message: QueryGetBaseledgerTransactionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBaseledgerTransactionResponse;
    fromJSON(object: any): QueryGetBaseledgerTransactionResponse;
    toJSON(message: QueryGetBaseledgerTransactionResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBaseledgerTransactionResponse>): QueryGetBaseledgerTransactionResponse;
};
export declare const QueryAllBaseledgerTransactionRequest: {
    encode(message: QueryAllBaseledgerTransactionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBaseledgerTransactionRequest;
    fromJSON(object: any): QueryAllBaseledgerTransactionRequest;
    toJSON(message: QueryAllBaseledgerTransactionRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllBaseledgerTransactionRequest>): QueryAllBaseledgerTransactionRequest;
};
export declare const QueryAllBaseledgerTransactionResponse: {
    encode(message: QueryAllBaseledgerTransactionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBaseledgerTransactionResponse;
    fromJSON(object: any): QueryAllBaseledgerTransactionResponse;
    toJSON(message: QueryAllBaseledgerTransactionResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllBaseledgerTransactionResponse>): QueryAllBaseledgerTransactionResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** this line is used by starport scaffolding # 2 */
    BaseledgerTransaction(request: QueryGetBaseledgerTransactionRequest): Promise<QueryGetBaseledgerTransactionResponse>;
    BaseledgerTransactionAll(request: QueryAllBaseledgerTransactionRequest): Promise<QueryAllBaseledgerTransactionResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    BaseledgerTransaction(request: QueryGetBaseledgerTransactionRequest): Promise<QueryGetBaseledgerTransactionResponse>;
    BaseledgerTransactionAll(request: QueryAllBaseledgerTransactionRequest): Promise<QueryAllBaseledgerTransactionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

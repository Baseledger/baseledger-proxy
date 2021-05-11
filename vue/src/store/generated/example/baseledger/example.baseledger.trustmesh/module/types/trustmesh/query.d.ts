import { Reader, Writer } from "protobufjs/minimal";
import { SynchronizationFeedback } from "../trustmesh/SynchronizationFeedback";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { SynchronizationRequest } from "../trustmesh/SynchronizationRequest";
export declare const protobufPackage = "example.baseledger.trustmesh";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetSynchronizationFeedbackRequest {
    id: number;
}
export interface QueryGetSynchronizationFeedbackResponse {
    SynchronizationFeedback: SynchronizationFeedback | undefined;
}
export interface QueryAllSynchronizationFeedbackRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSynchronizationFeedbackResponse {
    SynchronizationFeedback: SynchronizationFeedback[];
    pagination: PageResponse | undefined;
}
export interface QueryGetSynchronizationRequestRequest {
    id: number;
}
export interface QueryGetSynchronizationRequestResponse {
    SynchronizationRequest: SynchronizationRequest | undefined;
}
export interface QueryAllSynchronizationRequestRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSynchronizationRequestResponse {
    SynchronizationRequest: SynchronizationRequest[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetSynchronizationFeedbackRequest: {
    encode(message: QueryGetSynchronizationFeedbackRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSynchronizationFeedbackRequest;
    fromJSON(object: any): QueryGetSynchronizationFeedbackRequest;
    toJSON(message: QueryGetSynchronizationFeedbackRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSynchronizationFeedbackRequest>): QueryGetSynchronizationFeedbackRequest;
};
export declare const QueryGetSynchronizationFeedbackResponse: {
    encode(message: QueryGetSynchronizationFeedbackResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSynchronizationFeedbackResponse;
    fromJSON(object: any): QueryGetSynchronizationFeedbackResponse;
    toJSON(message: QueryGetSynchronizationFeedbackResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSynchronizationFeedbackResponse>): QueryGetSynchronizationFeedbackResponse;
};
export declare const QueryAllSynchronizationFeedbackRequest: {
    encode(message: QueryAllSynchronizationFeedbackRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSynchronizationFeedbackRequest;
    fromJSON(object: any): QueryAllSynchronizationFeedbackRequest;
    toJSON(message: QueryAllSynchronizationFeedbackRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSynchronizationFeedbackRequest>): QueryAllSynchronizationFeedbackRequest;
};
export declare const QueryAllSynchronizationFeedbackResponse: {
    encode(message: QueryAllSynchronizationFeedbackResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSynchronizationFeedbackResponse;
    fromJSON(object: any): QueryAllSynchronizationFeedbackResponse;
    toJSON(message: QueryAllSynchronizationFeedbackResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSynchronizationFeedbackResponse>): QueryAllSynchronizationFeedbackResponse;
};
export declare const QueryGetSynchronizationRequestRequest: {
    encode(message: QueryGetSynchronizationRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSynchronizationRequestRequest;
    fromJSON(object: any): QueryGetSynchronizationRequestRequest;
    toJSON(message: QueryGetSynchronizationRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSynchronizationRequestRequest>): QueryGetSynchronizationRequestRequest;
};
export declare const QueryGetSynchronizationRequestResponse: {
    encode(message: QueryGetSynchronizationRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSynchronizationRequestResponse;
    fromJSON(object: any): QueryGetSynchronizationRequestResponse;
    toJSON(message: QueryGetSynchronizationRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSynchronizationRequestResponse>): QueryGetSynchronizationRequestResponse;
};
export declare const QueryAllSynchronizationRequestRequest: {
    encode(message: QueryAllSynchronizationRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSynchronizationRequestRequest;
    fromJSON(object: any): QueryAllSynchronizationRequestRequest;
    toJSON(message: QueryAllSynchronizationRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSynchronizationRequestRequest>): QueryAllSynchronizationRequestRequest;
};
export declare const QueryAllSynchronizationRequestResponse: {
    encode(message: QueryAllSynchronizationRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSynchronizationRequestResponse;
    fromJSON(object: any): QueryAllSynchronizationRequestResponse;
    toJSON(message: QueryAllSynchronizationRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSynchronizationRequestResponse>): QueryAllSynchronizationRequestResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** this line is used by starport scaffolding # 2 */
    SynchronizationFeedback(request: QueryGetSynchronizationFeedbackRequest): Promise<QueryGetSynchronizationFeedbackResponse>;
    SynchronizationFeedbackAll(request: QueryAllSynchronizationFeedbackRequest): Promise<QueryAllSynchronizationFeedbackResponse>;
    SynchronizationRequest(request: QueryGetSynchronizationRequestRequest): Promise<QueryGetSynchronizationRequestResponse>;
    SynchronizationRequestAll(request: QueryAllSynchronizationRequestRequest): Promise<QueryAllSynchronizationRequestResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    SynchronizationFeedback(request: QueryGetSynchronizationFeedbackRequest): Promise<QueryGetSynchronizationFeedbackResponse>;
    SynchronizationFeedbackAll(request: QueryAllSynchronizationFeedbackRequest): Promise<QueryAllSynchronizationFeedbackResponse>;
    SynchronizationRequest(request: QueryGetSynchronizationRequestRequest): Promise<QueryGetSynchronizationRequestResponse>;
    SynchronizationRequestAll(request: QueryAllSynchronizationRequestRequest): Promise<QueryAllSynchronizationRequestResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

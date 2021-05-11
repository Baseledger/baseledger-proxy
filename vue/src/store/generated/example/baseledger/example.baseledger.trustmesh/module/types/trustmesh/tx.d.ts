import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "example.baseledger.trustmesh";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateSynchronizationFeedback {
    creator: string;
    Approved: string;
    BusinessObject: string;
    BaseledgerBusinessObjectIDOfApprovedObject: string;
    Workgroup: string;
    Recipient: string;
    HashOfObjectToApprove: string;
    OriginalBaseledgerTransactionID: string;
    OriginalOffchainProcessMessageID: string;
    FeedbackMessage: string;
}
export interface MsgCreateSynchronizationFeedbackResponse {
    id: number;
}
export interface MsgUpdateSynchronizationFeedback {
    creator: string;
    id: number;
    Approved: string;
    BusinessObject: string;
    BaseledgerBusinessObjectIDOfApprovedObject: string;
    Workgroup: string;
    Recipient: string;
    HashOfObjectToApprove: string;
    OriginalBaseledgerTransactionID: string;
    OriginalOffchainProcessMessageID: string;
    FeedbackMessage: string;
}
export interface MsgUpdateSynchronizationFeedbackResponse {
}
export interface MsgDeleteSynchronizationFeedback {
    creator: string;
    id: number;
}
export interface MsgDeleteSynchronizationFeedbackResponse {
}
export interface MsgCreateSynchronizationRequest {
    creator: string;
    WorkgroupID: string;
    Recipient: string;
    WorkstepType: string;
    BusinessObjectType: string;
    BaseledgerBusinessObjectID: string;
    BusinessObject: string;
    ReferencedBaseledgerBusinessObjectID: string;
}
export interface MsgCreateSynchronizationRequestResponse {
    id: number;
}
export interface MsgUpdateSynchronizationRequest {
    creator: string;
    id: number;
    WorkgroupID: string;
    Recipient: string;
    WorkstepType: string;
    BusinessObjectType: string;
    BaseledgerBusinessObjectID: string;
    BusinessObject: string;
    ReferencedBaseledgerBusinessObjectID: string;
}
export interface MsgUpdateSynchronizationRequestResponse {
}
export interface MsgDeleteSynchronizationRequest {
    creator: string;
    id: number;
}
export interface MsgDeleteSynchronizationRequestResponse {
}
export declare const MsgCreateSynchronizationFeedback: {
    encode(message: MsgCreateSynchronizationFeedback, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSynchronizationFeedback;
    fromJSON(object: any): MsgCreateSynchronizationFeedback;
    toJSON(message: MsgCreateSynchronizationFeedback): unknown;
    fromPartial(object: DeepPartial<MsgCreateSynchronizationFeedback>): MsgCreateSynchronizationFeedback;
};
export declare const MsgCreateSynchronizationFeedbackResponse: {
    encode(message: MsgCreateSynchronizationFeedbackResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSynchronizationFeedbackResponse;
    fromJSON(object: any): MsgCreateSynchronizationFeedbackResponse;
    toJSON(message: MsgCreateSynchronizationFeedbackResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateSynchronizationFeedbackResponse>): MsgCreateSynchronizationFeedbackResponse;
};
export declare const MsgUpdateSynchronizationFeedback: {
    encode(message: MsgUpdateSynchronizationFeedback, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSynchronizationFeedback;
    fromJSON(object: any): MsgUpdateSynchronizationFeedback;
    toJSON(message: MsgUpdateSynchronizationFeedback): unknown;
    fromPartial(object: DeepPartial<MsgUpdateSynchronizationFeedback>): MsgUpdateSynchronizationFeedback;
};
export declare const MsgUpdateSynchronizationFeedbackResponse: {
    encode(_: MsgUpdateSynchronizationFeedbackResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSynchronizationFeedbackResponse;
    fromJSON(_: any): MsgUpdateSynchronizationFeedbackResponse;
    toJSON(_: MsgUpdateSynchronizationFeedbackResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateSynchronizationFeedbackResponse>): MsgUpdateSynchronizationFeedbackResponse;
};
export declare const MsgDeleteSynchronizationFeedback: {
    encode(message: MsgDeleteSynchronizationFeedback, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSynchronizationFeedback;
    fromJSON(object: any): MsgDeleteSynchronizationFeedback;
    toJSON(message: MsgDeleteSynchronizationFeedback): unknown;
    fromPartial(object: DeepPartial<MsgDeleteSynchronizationFeedback>): MsgDeleteSynchronizationFeedback;
};
export declare const MsgDeleteSynchronizationFeedbackResponse: {
    encode(_: MsgDeleteSynchronizationFeedbackResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSynchronizationFeedbackResponse;
    fromJSON(_: any): MsgDeleteSynchronizationFeedbackResponse;
    toJSON(_: MsgDeleteSynchronizationFeedbackResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteSynchronizationFeedbackResponse>): MsgDeleteSynchronizationFeedbackResponse;
};
export declare const MsgCreateSynchronizationRequest: {
    encode(message: MsgCreateSynchronizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSynchronizationRequest;
    fromJSON(object: any): MsgCreateSynchronizationRequest;
    toJSON(message: MsgCreateSynchronizationRequest): unknown;
    fromPartial(object: DeepPartial<MsgCreateSynchronizationRequest>): MsgCreateSynchronizationRequest;
};
export declare const MsgCreateSynchronizationRequestResponse: {
    encode(message: MsgCreateSynchronizationRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSynchronizationRequestResponse;
    fromJSON(object: any): MsgCreateSynchronizationRequestResponse;
    toJSON(message: MsgCreateSynchronizationRequestResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateSynchronizationRequestResponse>): MsgCreateSynchronizationRequestResponse;
};
export declare const MsgUpdateSynchronizationRequest: {
    encode(message: MsgUpdateSynchronizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSynchronizationRequest;
    fromJSON(object: any): MsgUpdateSynchronizationRequest;
    toJSON(message: MsgUpdateSynchronizationRequest): unknown;
    fromPartial(object: DeepPartial<MsgUpdateSynchronizationRequest>): MsgUpdateSynchronizationRequest;
};
export declare const MsgUpdateSynchronizationRequestResponse: {
    encode(_: MsgUpdateSynchronizationRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSynchronizationRequestResponse;
    fromJSON(_: any): MsgUpdateSynchronizationRequestResponse;
    toJSON(_: MsgUpdateSynchronizationRequestResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateSynchronizationRequestResponse>): MsgUpdateSynchronizationRequestResponse;
};
export declare const MsgDeleteSynchronizationRequest: {
    encode(message: MsgDeleteSynchronizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSynchronizationRequest;
    fromJSON(object: any): MsgDeleteSynchronizationRequest;
    toJSON(message: MsgDeleteSynchronizationRequest): unknown;
    fromPartial(object: DeepPartial<MsgDeleteSynchronizationRequest>): MsgDeleteSynchronizationRequest;
};
export declare const MsgDeleteSynchronizationRequestResponse: {
    encode(_: MsgDeleteSynchronizationRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSynchronizationRequestResponse;
    fromJSON(_: any): MsgDeleteSynchronizationRequestResponse;
    toJSON(_: MsgDeleteSynchronizationRequestResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteSynchronizationRequestResponse>): MsgDeleteSynchronizationRequestResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateSynchronizationFeedback(request: MsgCreateSynchronizationFeedback): Promise<MsgCreateSynchronizationFeedbackResponse>;
    UpdateSynchronizationFeedback(request: MsgUpdateSynchronizationFeedback): Promise<MsgUpdateSynchronizationFeedbackResponse>;
    DeleteSynchronizationFeedback(request: MsgDeleteSynchronizationFeedback): Promise<MsgDeleteSynchronizationFeedbackResponse>;
    CreateSynchronizationRequest(request: MsgCreateSynchronizationRequest): Promise<MsgCreateSynchronizationRequestResponse>;
    UpdateSynchronizationRequest(request: MsgUpdateSynchronizationRequest): Promise<MsgUpdateSynchronizationRequestResponse>;
    DeleteSynchronizationRequest(request: MsgDeleteSynchronizationRequest): Promise<MsgDeleteSynchronizationRequestResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateSynchronizationFeedback(request: MsgCreateSynchronizationFeedback): Promise<MsgCreateSynchronizationFeedbackResponse>;
    UpdateSynchronizationFeedback(request: MsgUpdateSynchronizationFeedback): Promise<MsgUpdateSynchronizationFeedbackResponse>;
    DeleteSynchronizationFeedback(request: MsgDeleteSynchronizationFeedback): Promise<MsgDeleteSynchronizationFeedbackResponse>;
    CreateSynchronizationRequest(request: MsgCreateSynchronizationRequest): Promise<MsgCreateSynchronizationRequestResponse>;
    UpdateSynchronizationRequest(request: MsgUpdateSynchronizationRequest): Promise<MsgUpdateSynchronizationRequestResponse>;
    DeleteSynchronizationRequest(request: MsgDeleteSynchronizationRequest): Promise<MsgDeleteSynchronizationRequestResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

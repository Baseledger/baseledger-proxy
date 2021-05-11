import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "example.baseledger.trustmesh";
export interface SynchronizationRequest {
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
export declare const SynchronizationRequest: {
    encode(message: SynchronizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SynchronizationRequest;
    fromJSON(object: any): SynchronizationRequest;
    toJSON(message: SynchronizationRequest): unknown;
    fromPartial(object: DeepPartial<SynchronizationRequest>): SynchronizationRequest;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

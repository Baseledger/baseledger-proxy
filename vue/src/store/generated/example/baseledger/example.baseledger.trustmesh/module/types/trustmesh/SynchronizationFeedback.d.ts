import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "example.baseledger.trustmesh";
export interface SynchronizationFeedback {
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
export declare const SynchronizationFeedback: {
    encode(message: SynchronizationFeedback, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SynchronizationFeedback;
    fromJSON(object: any): SynchronizationFeedback;
    toJSON(message: SynchronizationFeedback): unknown;
    fromPartial(object: DeepPartial<SynchronizationFeedback>): SynchronizationFeedback;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

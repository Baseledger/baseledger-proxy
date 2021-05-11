/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "example.baseledger.trustmesh";

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

const baseSynchronizationFeedback: object = {
  creator: "",
  id: 0,
  Approved: "",
  BusinessObject: "",
  BaseledgerBusinessObjectIDOfApprovedObject: "",
  Workgroup: "",
  Recipient: "",
  HashOfObjectToApprove: "",
  OriginalBaseledgerTransactionID: "",
  OriginalOffchainProcessMessageID: "",
  FeedbackMessage: "",
};

export const SynchronizationFeedback = {
  encode(
    message: SynchronizationFeedback,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.Approved !== "") {
      writer.uint32(26).string(message.Approved);
    }
    if (message.BusinessObject !== "") {
      writer.uint32(34).string(message.BusinessObject);
    }
    if (message.BaseledgerBusinessObjectIDOfApprovedObject !== "") {
      writer
        .uint32(42)
        .string(message.BaseledgerBusinessObjectIDOfApprovedObject);
    }
    if (message.Workgroup !== "") {
      writer.uint32(50).string(message.Workgroup);
    }
    if (message.Recipient !== "") {
      writer.uint32(58).string(message.Recipient);
    }
    if (message.HashOfObjectToApprove !== "") {
      writer.uint32(66).string(message.HashOfObjectToApprove);
    }
    if (message.OriginalBaseledgerTransactionID !== "") {
      writer.uint32(74).string(message.OriginalBaseledgerTransactionID);
    }
    if (message.OriginalOffchainProcessMessageID !== "") {
      writer.uint32(82).string(message.OriginalOffchainProcessMessageID);
    }
    if (message.FeedbackMessage !== "") {
      writer.uint32(90).string(message.FeedbackMessage);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SynchronizationFeedback {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSynchronizationFeedback,
    } as SynchronizationFeedback;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.Approved = reader.string();
          break;
        case 4:
          message.BusinessObject = reader.string();
          break;
        case 5:
          message.BaseledgerBusinessObjectIDOfApprovedObject = reader.string();
          break;
        case 6:
          message.Workgroup = reader.string();
          break;
        case 7:
          message.Recipient = reader.string();
          break;
        case 8:
          message.HashOfObjectToApprove = reader.string();
          break;
        case 9:
          message.OriginalBaseledgerTransactionID = reader.string();
          break;
        case 10:
          message.OriginalOffchainProcessMessageID = reader.string();
          break;
        case 11:
          message.FeedbackMessage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SynchronizationFeedback {
    const message = {
      ...baseSynchronizationFeedback,
    } as SynchronizationFeedback;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.Approved !== undefined && object.Approved !== null) {
      message.Approved = String(object.Approved);
    } else {
      message.Approved = "";
    }
    if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
      message.BusinessObject = String(object.BusinessObject);
    } else {
      message.BusinessObject = "";
    }
    if (
      object.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
      object.BaseledgerBusinessObjectIDOfApprovedObject !== null
    ) {
      message.BaseledgerBusinessObjectIDOfApprovedObject = String(
        object.BaseledgerBusinessObjectIDOfApprovedObject
      );
    } else {
      message.BaseledgerBusinessObjectIDOfApprovedObject = "";
    }
    if (object.Workgroup !== undefined && object.Workgroup !== null) {
      message.Workgroup = String(object.Workgroup);
    } else {
      message.Workgroup = "";
    }
    if (object.Recipient !== undefined && object.Recipient !== null) {
      message.Recipient = String(object.Recipient);
    } else {
      message.Recipient = "";
    }
    if (
      object.HashOfObjectToApprove !== undefined &&
      object.HashOfObjectToApprove !== null
    ) {
      message.HashOfObjectToApprove = String(object.HashOfObjectToApprove);
    } else {
      message.HashOfObjectToApprove = "";
    }
    if (
      object.OriginalBaseledgerTransactionID !== undefined &&
      object.OriginalBaseledgerTransactionID !== null
    ) {
      message.OriginalBaseledgerTransactionID = String(
        object.OriginalBaseledgerTransactionID
      );
    } else {
      message.OriginalBaseledgerTransactionID = "";
    }
    if (
      object.OriginalOffchainProcessMessageID !== undefined &&
      object.OriginalOffchainProcessMessageID !== null
    ) {
      message.OriginalOffchainProcessMessageID = String(
        object.OriginalOffchainProcessMessageID
      );
    } else {
      message.OriginalOffchainProcessMessageID = "";
    }
    if (
      object.FeedbackMessage !== undefined &&
      object.FeedbackMessage !== null
    ) {
      message.FeedbackMessage = String(object.FeedbackMessage);
    } else {
      message.FeedbackMessage = "";
    }
    return message;
  },

  toJSON(message: SynchronizationFeedback): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.Approved !== undefined && (obj.Approved = message.Approved);
    message.BusinessObject !== undefined &&
      (obj.BusinessObject = message.BusinessObject);
    message.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
      (obj.BaseledgerBusinessObjectIDOfApprovedObject =
        message.BaseledgerBusinessObjectIDOfApprovedObject);
    message.Workgroup !== undefined && (obj.Workgroup = message.Workgroup);
    message.Recipient !== undefined && (obj.Recipient = message.Recipient);
    message.HashOfObjectToApprove !== undefined &&
      (obj.HashOfObjectToApprove = message.HashOfObjectToApprove);
    message.OriginalBaseledgerTransactionID !== undefined &&
      (obj.OriginalBaseledgerTransactionID =
        message.OriginalBaseledgerTransactionID);
    message.OriginalOffchainProcessMessageID !== undefined &&
      (obj.OriginalOffchainProcessMessageID =
        message.OriginalOffchainProcessMessageID);
    message.FeedbackMessage !== undefined &&
      (obj.FeedbackMessage = message.FeedbackMessage);
    return obj;
  },

  fromPartial(
    object: DeepPartial<SynchronizationFeedback>
  ): SynchronizationFeedback {
    const message = {
      ...baseSynchronizationFeedback,
    } as SynchronizationFeedback;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.Approved !== undefined && object.Approved !== null) {
      message.Approved = object.Approved;
    } else {
      message.Approved = "";
    }
    if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
      message.BusinessObject = object.BusinessObject;
    } else {
      message.BusinessObject = "";
    }
    if (
      object.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
      object.BaseledgerBusinessObjectIDOfApprovedObject !== null
    ) {
      message.BaseledgerBusinessObjectIDOfApprovedObject =
        object.BaseledgerBusinessObjectIDOfApprovedObject;
    } else {
      message.BaseledgerBusinessObjectIDOfApprovedObject = "";
    }
    if (object.Workgroup !== undefined && object.Workgroup !== null) {
      message.Workgroup = object.Workgroup;
    } else {
      message.Workgroup = "";
    }
    if (object.Recipient !== undefined && object.Recipient !== null) {
      message.Recipient = object.Recipient;
    } else {
      message.Recipient = "";
    }
    if (
      object.HashOfObjectToApprove !== undefined &&
      object.HashOfObjectToApprove !== null
    ) {
      message.HashOfObjectToApprove = object.HashOfObjectToApprove;
    } else {
      message.HashOfObjectToApprove = "";
    }
    if (
      object.OriginalBaseledgerTransactionID !== undefined &&
      object.OriginalBaseledgerTransactionID !== null
    ) {
      message.OriginalBaseledgerTransactionID =
        object.OriginalBaseledgerTransactionID;
    } else {
      message.OriginalBaseledgerTransactionID = "";
    }
    if (
      object.OriginalOffchainProcessMessageID !== undefined &&
      object.OriginalOffchainProcessMessageID !== null
    ) {
      message.OriginalOffchainProcessMessageID =
        object.OriginalOffchainProcessMessageID;
    } else {
      message.OriginalOffchainProcessMessageID = "";
    }
    if (
      object.FeedbackMessage !== undefined &&
      object.FeedbackMessage !== null
    ) {
      message.FeedbackMessage = object.FeedbackMessage;
    } else {
      message.FeedbackMessage = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

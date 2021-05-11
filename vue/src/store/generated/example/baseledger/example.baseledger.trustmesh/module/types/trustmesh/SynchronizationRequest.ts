/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "example.baseledger.trustmesh";

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

const baseSynchronizationRequest: object = {
  creator: "",
  id: 0,
  WorkgroupID: "",
  Recipient: "",
  WorkstepType: "",
  BusinessObjectType: "",
  BaseledgerBusinessObjectID: "",
  BusinessObject: "",
  ReferencedBaseledgerBusinessObjectID: "",
};

export const SynchronizationRequest = {
  encode(
    message: SynchronizationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.WorkgroupID !== "") {
      writer.uint32(26).string(message.WorkgroupID);
    }
    if (message.Recipient !== "") {
      writer.uint32(34).string(message.Recipient);
    }
    if (message.WorkstepType !== "") {
      writer.uint32(42).string(message.WorkstepType);
    }
    if (message.BusinessObjectType !== "") {
      writer.uint32(50).string(message.BusinessObjectType);
    }
    if (message.BaseledgerBusinessObjectID !== "") {
      writer.uint32(58).string(message.BaseledgerBusinessObjectID);
    }
    if (message.BusinessObject !== "") {
      writer.uint32(66).string(message.BusinessObject);
    }
    if (message.ReferencedBaseledgerBusinessObjectID !== "") {
      writer.uint32(74).string(message.ReferencedBaseledgerBusinessObjectID);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SynchronizationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSynchronizationRequest } as SynchronizationRequest;
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
          message.WorkgroupID = reader.string();
          break;
        case 4:
          message.Recipient = reader.string();
          break;
        case 5:
          message.WorkstepType = reader.string();
          break;
        case 6:
          message.BusinessObjectType = reader.string();
          break;
        case 7:
          message.BaseledgerBusinessObjectID = reader.string();
          break;
        case 8:
          message.BusinessObject = reader.string();
          break;
        case 9:
          message.ReferencedBaseledgerBusinessObjectID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SynchronizationRequest {
    const message = { ...baseSynchronizationRequest } as SynchronizationRequest;
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
    if (object.WorkgroupID !== undefined && object.WorkgroupID !== null) {
      message.WorkgroupID = String(object.WorkgroupID);
    } else {
      message.WorkgroupID = "";
    }
    if (object.Recipient !== undefined && object.Recipient !== null) {
      message.Recipient = String(object.Recipient);
    } else {
      message.Recipient = "";
    }
    if (object.WorkstepType !== undefined && object.WorkstepType !== null) {
      message.WorkstepType = String(object.WorkstepType);
    } else {
      message.WorkstepType = "";
    }
    if (
      object.BusinessObjectType !== undefined &&
      object.BusinessObjectType !== null
    ) {
      message.BusinessObjectType = String(object.BusinessObjectType);
    } else {
      message.BusinessObjectType = "";
    }
    if (
      object.BaseledgerBusinessObjectID !== undefined &&
      object.BaseledgerBusinessObjectID !== null
    ) {
      message.BaseledgerBusinessObjectID = String(
        object.BaseledgerBusinessObjectID
      );
    } else {
      message.BaseledgerBusinessObjectID = "";
    }
    if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
      message.BusinessObject = String(object.BusinessObject);
    } else {
      message.BusinessObject = "";
    }
    if (
      object.ReferencedBaseledgerBusinessObjectID !== undefined &&
      object.ReferencedBaseledgerBusinessObjectID !== null
    ) {
      message.ReferencedBaseledgerBusinessObjectID = String(
        object.ReferencedBaseledgerBusinessObjectID
      );
    } else {
      message.ReferencedBaseledgerBusinessObjectID = "";
    }
    return message;
  },

  toJSON(message: SynchronizationRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.WorkgroupID !== undefined &&
      (obj.WorkgroupID = message.WorkgroupID);
    message.Recipient !== undefined && (obj.Recipient = message.Recipient);
    message.WorkstepType !== undefined &&
      (obj.WorkstepType = message.WorkstepType);
    message.BusinessObjectType !== undefined &&
      (obj.BusinessObjectType = message.BusinessObjectType);
    message.BaseledgerBusinessObjectID !== undefined &&
      (obj.BaseledgerBusinessObjectID = message.BaseledgerBusinessObjectID);
    message.BusinessObject !== undefined &&
      (obj.BusinessObject = message.BusinessObject);
    message.ReferencedBaseledgerBusinessObjectID !== undefined &&
      (obj.ReferencedBaseledgerBusinessObjectID =
        message.ReferencedBaseledgerBusinessObjectID);
    return obj;
  },

  fromPartial(
    object: DeepPartial<SynchronizationRequest>
  ): SynchronizationRequest {
    const message = { ...baseSynchronizationRequest } as SynchronizationRequest;
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
    if (object.WorkgroupID !== undefined && object.WorkgroupID !== null) {
      message.WorkgroupID = object.WorkgroupID;
    } else {
      message.WorkgroupID = "";
    }
    if (object.Recipient !== undefined && object.Recipient !== null) {
      message.Recipient = object.Recipient;
    } else {
      message.Recipient = "";
    }
    if (object.WorkstepType !== undefined && object.WorkstepType !== null) {
      message.WorkstepType = object.WorkstepType;
    } else {
      message.WorkstepType = "";
    }
    if (
      object.BusinessObjectType !== undefined &&
      object.BusinessObjectType !== null
    ) {
      message.BusinessObjectType = object.BusinessObjectType;
    } else {
      message.BusinessObjectType = "";
    }
    if (
      object.BaseledgerBusinessObjectID !== undefined &&
      object.BaseledgerBusinessObjectID !== null
    ) {
      message.BaseledgerBusinessObjectID = object.BaseledgerBusinessObjectID;
    } else {
      message.BaseledgerBusinessObjectID = "";
    }
    if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
      message.BusinessObject = object.BusinessObject;
    } else {
      message.BusinessObject = "";
    }
    if (
      object.ReferencedBaseledgerBusinessObjectID !== undefined &&
      object.ReferencedBaseledgerBusinessObjectID !== null
    ) {
      message.ReferencedBaseledgerBusinessObjectID =
        object.ReferencedBaseledgerBusinessObjectID;
    } else {
      message.ReferencedBaseledgerBusinessObjectID = "";
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

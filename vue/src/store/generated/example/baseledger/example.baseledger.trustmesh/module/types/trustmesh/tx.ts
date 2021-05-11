/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "example.baseledger.trustmesh";

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

export interface MsgUpdateSynchronizationFeedbackResponse {}

export interface MsgDeleteSynchronizationFeedback {
  creator: string;
  id: number;
}

export interface MsgDeleteSynchronizationFeedbackResponse {}

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

export interface MsgUpdateSynchronizationRequestResponse {}

export interface MsgDeleteSynchronizationRequest {
  creator: string;
  id: number;
}

export interface MsgDeleteSynchronizationRequestResponse {}

const baseMsgCreateSynchronizationFeedback: object = {
  creator: "",
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

export const MsgCreateSynchronizationFeedback = {
  encode(
    message: MsgCreateSynchronizationFeedback,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.Approved !== "") {
      writer.uint32(18).string(message.Approved);
    }
    if (message.BusinessObject !== "") {
      writer.uint32(26).string(message.BusinessObject);
    }
    if (message.BaseledgerBusinessObjectIDOfApprovedObject !== "") {
      writer
        .uint32(34)
        .string(message.BaseledgerBusinessObjectIDOfApprovedObject);
    }
    if (message.Workgroup !== "") {
      writer.uint32(42).string(message.Workgroup);
    }
    if (message.Recipient !== "") {
      writer.uint32(50).string(message.Recipient);
    }
    if (message.HashOfObjectToApprove !== "") {
      writer.uint32(58).string(message.HashOfObjectToApprove);
    }
    if (message.OriginalBaseledgerTransactionID !== "") {
      writer.uint32(66).string(message.OriginalBaseledgerTransactionID);
    }
    if (message.OriginalOffchainProcessMessageID !== "") {
      writer.uint32(74).string(message.OriginalOffchainProcessMessageID);
    }
    if (message.FeedbackMessage !== "") {
      writer.uint32(82).string(message.FeedbackMessage);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSynchronizationFeedback {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSynchronizationFeedback,
    } as MsgCreateSynchronizationFeedback;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.Approved = reader.string();
          break;
        case 3:
          message.BusinessObject = reader.string();
          break;
        case 4:
          message.BaseledgerBusinessObjectIDOfApprovedObject = reader.string();
          break;
        case 5:
          message.Workgroup = reader.string();
          break;
        case 6:
          message.Recipient = reader.string();
          break;
        case 7:
          message.HashOfObjectToApprove = reader.string();
          break;
        case 8:
          message.OriginalBaseledgerTransactionID = reader.string();
          break;
        case 9:
          message.OriginalOffchainProcessMessageID = reader.string();
          break;
        case 10:
          message.FeedbackMessage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSynchronizationFeedback {
    const message = {
      ...baseMsgCreateSynchronizationFeedback,
    } as MsgCreateSynchronizationFeedback;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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

  toJSON(message: MsgCreateSynchronizationFeedback): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
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
    object: DeepPartial<MsgCreateSynchronizationFeedback>
  ): MsgCreateSynchronizationFeedback {
    const message = {
      ...baseMsgCreateSynchronizationFeedback,
    } as MsgCreateSynchronizationFeedback;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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

const baseMsgCreateSynchronizationFeedbackResponse: object = { id: 0 };

export const MsgCreateSynchronizationFeedbackResponse = {
  encode(
    message: MsgCreateSynchronizationFeedbackResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSynchronizationFeedbackResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSynchronizationFeedbackResponse,
    } as MsgCreateSynchronizationFeedbackResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSynchronizationFeedbackResponse {
    const message = {
      ...baseMsgCreateSynchronizationFeedbackResponse,
    } as MsgCreateSynchronizationFeedbackResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSynchronizationFeedbackResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSynchronizationFeedbackResponse>
  ): MsgCreateSynchronizationFeedbackResponse {
    const message = {
      ...baseMsgCreateSynchronizationFeedbackResponse,
    } as MsgCreateSynchronizationFeedbackResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateSynchronizationFeedback: object = {
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

export const MsgUpdateSynchronizationFeedback = {
  encode(
    message: MsgUpdateSynchronizationFeedback,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSynchronizationFeedback {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSynchronizationFeedback,
    } as MsgUpdateSynchronizationFeedback;
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

  fromJSON(object: any): MsgUpdateSynchronizationFeedback {
    const message = {
      ...baseMsgUpdateSynchronizationFeedback,
    } as MsgUpdateSynchronizationFeedback;
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

  toJSON(message: MsgUpdateSynchronizationFeedback): unknown {
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
    object: DeepPartial<MsgUpdateSynchronizationFeedback>
  ): MsgUpdateSynchronizationFeedback {
    const message = {
      ...baseMsgUpdateSynchronizationFeedback,
    } as MsgUpdateSynchronizationFeedback;
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

const baseMsgUpdateSynchronizationFeedbackResponse: object = {};

export const MsgUpdateSynchronizationFeedbackResponse = {
  encode(
    _: MsgUpdateSynchronizationFeedbackResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSynchronizationFeedbackResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSynchronizationFeedbackResponse,
    } as MsgUpdateSynchronizationFeedbackResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateSynchronizationFeedbackResponse {
    const message = {
      ...baseMsgUpdateSynchronizationFeedbackResponse,
    } as MsgUpdateSynchronizationFeedbackResponse;
    return message;
  },

  toJSON(_: MsgUpdateSynchronizationFeedbackResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSynchronizationFeedbackResponse>
  ): MsgUpdateSynchronizationFeedbackResponse {
    const message = {
      ...baseMsgUpdateSynchronizationFeedbackResponse,
    } as MsgUpdateSynchronizationFeedbackResponse;
    return message;
  },
};

const baseMsgDeleteSynchronizationFeedback: object = { creator: "", id: 0 };

export const MsgDeleteSynchronizationFeedback = {
  encode(
    message: MsgDeleteSynchronizationFeedback,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSynchronizationFeedback {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSynchronizationFeedback,
    } as MsgDeleteSynchronizationFeedback;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteSynchronizationFeedback {
    const message = {
      ...baseMsgDeleteSynchronizationFeedback,
    } as MsgDeleteSynchronizationFeedback;
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
    return message;
  },

  toJSON(message: MsgDeleteSynchronizationFeedback): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteSynchronizationFeedback>
  ): MsgDeleteSynchronizationFeedback {
    const message = {
      ...baseMsgDeleteSynchronizationFeedback,
    } as MsgDeleteSynchronizationFeedback;
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
    return message;
  },
};

const baseMsgDeleteSynchronizationFeedbackResponse: object = {};

export const MsgDeleteSynchronizationFeedbackResponse = {
  encode(
    _: MsgDeleteSynchronizationFeedbackResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSynchronizationFeedbackResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSynchronizationFeedbackResponse,
    } as MsgDeleteSynchronizationFeedbackResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteSynchronizationFeedbackResponse {
    const message = {
      ...baseMsgDeleteSynchronizationFeedbackResponse,
    } as MsgDeleteSynchronizationFeedbackResponse;
    return message;
  },

  toJSON(_: MsgDeleteSynchronizationFeedbackResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSynchronizationFeedbackResponse>
  ): MsgDeleteSynchronizationFeedbackResponse {
    const message = {
      ...baseMsgDeleteSynchronizationFeedbackResponse,
    } as MsgDeleteSynchronizationFeedbackResponse;
    return message;
  },
};

const baseMsgCreateSynchronizationRequest: object = {
  creator: "",
  WorkgroupID: "",
  Recipient: "",
  WorkstepType: "",
  BusinessObjectType: "",
  BaseledgerBusinessObjectID: "",
  BusinessObject: "",
  ReferencedBaseledgerBusinessObjectID: "",
};

export const MsgCreateSynchronizationRequest = {
  encode(
    message: MsgCreateSynchronizationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.WorkgroupID !== "") {
      writer.uint32(18).string(message.WorkgroupID);
    }
    if (message.Recipient !== "") {
      writer.uint32(26).string(message.Recipient);
    }
    if (message.WorkstepType !== "") {
      writer.uint32(34).string(message.WorkstepType);
    }
    if (message.BusinessObjectType !== "") {
      writer.uint32(42).string(message.BusinessObjectType);
    }
    if (message.BaseledgerBusinessObjectID !== "") {
      writer.uint32(50).string(message.BaseledgerBusinessObjectID);
    }
    if (message.BusinessObject !== "") {
      writer.uint32(58).string(message.BusinessObject);
    }
    if (message.ReferencedBaseledgerBusinessObjectID !== "") {
      writer.uint32(66).string(message.ReferencedBaseledgerBusinessObjectID);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSynchronizationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSynchronizationRequest,
    } as MsgCreateSynchronizationRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.WorkgroupID = reader.string();
          break;
        case 3:
          message.Recipient = reader.string();
          break;
        case 4:
          message.WorkstepType = reader.string();
          break;
        case 5:
          message.BusinessObjectType = reader.string();
          break;
        case 6:
          message.BaseledgerBusinessObjectID = reader.string();
          break;
        case 7:
          message.BusinessObject = reader.string();
          break;
        case 8:
          message.ReferencedBaseledgerBusinessObjectID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSynchronizationRequest {
    const message = {
      ...baseMsgCreateSynchronizationRequest,
    } as MsgCreateSynchronizationRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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

  toJSON(message: MsgCreateSynchronizationRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
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
    object: DeepPartial<MsgCreateSynchronizationRequest>
  ): MsgCreateSynchronizationRequest {
    const message = {
      ...baseMsgCreateSynchronizationRequest,
    } as MsgCreateSynchronizationRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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

const baseMsgCreateSynchronizationRequestResponse: object = { id: 0 };

export const MsgCreateSynchronizationRequestResponse = {
  encode(
    message: MsgCreateSynchronizationRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSynchronizationRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSynchronizationRequestResponse,
    } as MsgCreateSynchronizationRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSynchronizationRequestResponse {
    const message = {
      ...baseMsgCreateSynchronizationRequestResponse,
    } as MsgCreateSynchronizationRequestResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSynchronizationRequestResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSynchronizationRequestResponse>
  ): MsgCreateSynchronizationRequestResponse {
    const message = {
      ...baseMsgCreateSynchronizationRequestResponse,
    } as MsgCreateSynchronizationRequestResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateSynchronizationRequest: object = {
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

export const MsgUpdateSynchronizationRequest = {
  encode(
    message: MsgUpdateSynchronizationRequest,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSynchronizationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSynchronizationRequest,
    } as MsgUpdateSynchronizationRequest;
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

  fromJSON(object: any): MsgUpdateSynchronizationRequest {
    const message = {
      ...baseMsgUpdateSynchronizationRequest,
    } as MsgUpdateSynchronizationRequest;
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

  toJSON(message: MsgUpdateSynchronizationRequest): unknown {
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
    object: DeepPartial<MsgUpdateSynchronizationRequest>
  ): MsgUpdateSynchronizationRequest {
    const message = {
      ...baseMsgUpdateSynchronizationRequest,
    } as MsgUpdateSynchronizationRequest;
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

const baseMsgUpdateSynchronizationRequestResponse: object = {};

export const MsgUpdateSynchronizationRequestResponse = {
  encode(
    _: MsgUpdateSynchronizationRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSynchronizationRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSynchronizationRequestResponse,
    } as MsgUpdateSynchronizationRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateSynchronizationRequestResponse {
    const message = {
      ...baseMsgUpdateSynchronizationRequestResponse,
    } as MsgUpdateSynchronizationRequestResponse;
    return message;
  },

  toJSON(_: MsgUpdateSynchronizationRequestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSynchronizationRequestResponse>
  ): MsgUpdateSynchronizationRequestResponse {
    const message = {
      ...baseMsgUpdateSynchronizationRequestResponse,
    } as MsgUpdateSynchronizationRequestResponse;
    return message;
  },
};

const baseMsgDeleteSynchronizationRequest: object = { creator: "", id: 0 };

export const MsgDeleteSynchronizationRequest = {
  encode(
    message: MsgDeleteSynchronizationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSynchronizationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSynchronizationRequest,
    } as MsgDeleteSynchronizationRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteSynchronizationRequest {
    const message = {
      ...baseMsgDeleteSynchronizationRequest,
    } as MsgDeleteSynchronizationRequest;
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
    return message;
  },

  toJSON(message: MsgDeleteSynchronizationRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteSynchronizationRequest>
  ): MsgDeleteSynchronizationRequest {
    const message = {
      ...baseMsgDeleteSynchronizationRequest,
    } as MsgDeleteSynchronizationRequest;
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
    return message;
  },
};

const baseMsgDeleteSynchronizationRequestResponse: object = {};

export const MsgDeleteSynchronizationRequestResponse = {
  encode(
    _: MsgDeleteSynchronizationRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSynchronizationRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSynchronizationRequestResponse,
    } as MsgDeleteSynchronizationRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteSynchronizationRequestResponse {
    const message = {
      ...baseMsgDeleteSynchronizationRequestResponse,
    } as MsgDeleteSynchronizationRequestResponse;
    return message;
  },

  toJSON(_: MsgDeleteSynchronizationRequestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSynchronizationRequestResponse>
  ): MsgDeleteSynchronizationRequestResponse {
    const message = {
      ...baseMsgDeleteSynchronizationRequestResponse,
    } as MsgDeleteSynchronizationRequestResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateSynchronizationFeedback(
    request: MsgCreateSynchronizationFeedback
  ): Promise<MsgCreateSynchronizationFeedbackResponse>;
  UpdateSynchronizationFeedback(
    request: MsgUpdateSynchronizationFeedback
  ): Promise<MsgUpdateSynchronizationFeedbackResponse>;
  DeleteSynchronizationFeedback(
    request: MsgDeleteSynchronizationFeedback
  ): Promise<MsgDeleteSynchronizationFeedbackResponse>;
  CreateSynchronizationRequest(
    request: MsgCreateSynchronizationRequest
  ): Promise<MsgCreateSynchronizationRequestResponse>;
  UpdateSynchronizationRequest(
    request: MsgUpdateSynchronizationRequest
  ): Promise<MsgUpdateSynchronizationRequestResponse>;
  DeleteSynchronizationRequest(
    request: MsgDeleteSynchronizationRequest
  ): Promise<MsgDeleteSynchronizationRequestResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateSynchronizationFeedback(
    request: MsgCreateSynchronizationFeedback
  ): Promise<MsgCreateSynchronizationFeedbackResponse> {
    const data = MsgCreateSynchronizationFeedback.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Msg",
      "CreateSynchronizationFeedback",
      data
    );
    return promise.then((data) =>
      MsgCreateSynchronizationFeedbackResponse.decode(new Reader(data))
    );
  }

  UpdateSynchronizationFeedback(
    request: MsgUpdateSynchronizationFeedback
  ): Promise<MsgUpdateSynchronizationFeedbackResponse> {
    const data = MsgUpdateSynchronizationFeedback.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Msg",
      "UpdateSynchronizationFeedback",
      data
    );
    return promise.then((data) =>
      MsgUpdateSynchronizationFeedbackResponse.decode(new Reader(data))
    );
  }

  DeleteSynchronizationFeedback(
    request: MsgDeleteSynchronizationFeedback
  ): Promise<MsgDeleteSynchronizationFeedbackResponse> {
    const data = MsgDeleteSynchronizationFeedback.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Msg",
      "DeleteSynchronizationFeedback",
      data
    );
    return promise.then((data) =>
      MsgDeleteSynchronizationFeedbackResponse.decode(new Reader(data))
    );
  }

  CreateSynchronizationRequest(
    request: MsgCreateSynchronizationRequest
  ): Promise<MsgCreateSynchronizationRequestResponse> {
    const data = MsgCreateSynchronizationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Msg",
      "CreateSynchronizationRequest",
      data
    );
    return promise.then((data) =>
      MsgCreateSynchronizationRequestResponse.decode(new Reader(data))
    );
  }

  UpdateSynchronizationRequest(
    request: MsgUpdateSynchronizationRequest
  ): Promise<MsgUpdateSynchronizationRequestResponse> {
    const data = MsgUpdateSynchronizationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Msg",
      "UpdateSynchronizationRequest",
      data
    );
    return promise.then((data) =>
      MsgUpdateSynchronizationRequestResponse.decode(new Reader(data))
    );
  }

  DeleteSynchronizationRequest(
    request: MsgDeleteSynchronizationRequest
  ): Promise<MsgDeleteSynchronizationRequestResponse> {
    const data = MsgDeleteSynchronizationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Msg",
      "DeleteSynchronizationRequest",
      data
    );
    return promise.then((data) =>
      MsgDeleteSynchronizationRequestResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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

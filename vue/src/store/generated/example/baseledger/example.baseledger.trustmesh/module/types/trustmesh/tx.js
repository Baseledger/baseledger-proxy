/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
export const protobufPackage = "example.baseledger.trustmesh";
const baseMsgCreateSynchronizationFeedback = {
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
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateSynchronizationFeedback,
        };
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
    fromJSON(object) {
        const message = {
            ...baseMsgCreateSynchronizationFeedback,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.Approved !== undefined && object.Approved !== null) {
            message.Approved = String(object.Approved);
        }
        else {
            message.Approved = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = String(object.BusinessObject);
        }
        else {
            message.BusinessObject = "";
        }
        if (object.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
            object.BaseledgerBusinessObjectIDOfApprovedObject !== null) {
            message.BaseledgerBusinessObjectIDOfApprovedObject = String(object.BaseledgerBusinessObjectIDOfApprovedObject);
        }
        else {
            message.BaseledgerBusinessObjectIDOfApprovedObject = "";
        }
        if (object.Workgroup !== undefined && object.Workgroup !== null) {
            message.Workgroup = String(object.Workgroup);
        }
        else {
            message.Workgroup = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = String(object.Recipient);
        }
        else {
            message.Recipient = "";
        }
        if (object.HashOfObjectToApprove !== undefined &&
            object.HashOfObjectToApprove !== null) {
            message.HashOfObjectToApprove = String(object.HashOfObjectToApprove);
        }
        else {
            message.HashOfObjectToApprove = "";
        }
        if (object.OriginalBaseledgerTransactionID !== undefined &&
            object.OriginalBaseledgerTransactionID !== null) {
            message.OriginalBaseledgerTransactionID = String(object.OriginalBaseledgerTransactionID);
        }
        else {
            message.OriginalBaseledgerTransactionID = "";
        }
        if (object.OriginalOffchainProcessMessageID !== undefined &&
            object.OriginalOffchainProcessMessageID !== null) {
            message.OriginalOffchainProcessMessageID = String(object.OriginalOffchainProcessMessageID);
        }
        else {
            message.OriginalOffchainProcessMessageID = "";
        }
        if (object.FeedbackMessage !== undefined &&
            object.FeedbackMessage !== null) {
            message.FeedbackMessage = String(object.FeedbackMessage);
        }
        else {
            message.FeedbackMessage = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = {
            ...baseMsgCreateSynchronizationFeedback,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.Approved !== undefined && object.Approved !== null) {
            message.Approved = object.Approved;
        }
        else {
            message.Approved = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = object.BusinessObject;
        }
        else {
            message.BusinessObject = "";
        }
        if (object.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
            object.BaseledgerBusinessObjectIDOfApprovedObject !== null) {
            message.BaseledgerBusinessObjectIDOfApprovedObject =
                object.BaseledgerBusinessObjectIDOfApprovedObject;
        }
        else {
            message.BaseledgerBusinessObjectIDOfApprovedObject = "";
        }
        if (object.Workgroup !== undefined && object.Workgroup !== null) {
            message.Workgroup = object.Workgroup;
        }
        else {
            message.Workgroup = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = object.Recipient;
        }
        else {
            message.Recipient = "";
        }
        if (object.HashOfObjectToApprove !== undefined &&
            object.HashOfObjectToApprove !== null) {
            message.HashOfObjectToApprove = object.HashOfObjectToApprove;
        }
        else {
            message.HashOfObjectToApprove = "";
        }
        if (object.OriginalBaseledgerTransactionID !== undefined &&
            object.OriginalBaseledgerTransactionID !== null) {
            message.OriginalBaseledgerTransactionID =
                object.OriginalBaseledgerTransactionID;
        }
        else {
            message.OriginalBaseledgerTransactionID = "";
        }
        if (object.OriginalOffchainProcessMessageID !== undefined &&
            object.OriginalOffchainProcessMessageID !== null) {
            message.OriginalOffchainProcessMessageID =
                object.OriginalOffchainProcessMessageID;
        }
        else {
            message.OriginalOffchainProcessMessageID = "";
        }
        if (object.FeedbackMessage !== undefined &&
            object.FeedbackMessage !== null) {
            message.FeedbackMessage = object.FeedbackMessage;
        }
        else {
            message.FeedbackMessage = "";
        }
        return message;
    },
};
const baseMsgCreateSynchronizationFeedbackResponse = { id: 0 };
export const MsgCreateSynchronizationFeedbackResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateSynchronizationFeedbackResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgCreateSynchronizationFeedbackResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgCreateSynchronizationFeedbackResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgUpdateSynchronizationFeedback = {
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
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateSynchronizationFeedback,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
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
    fromJSON(object) {
        const message = {
            ...baseMsgUpdateSynchronizationFeedback,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.Approved !== undefined && object.Approved !== null) {
            message.Approved = String(object.Approved);
        }
        else {
            message.Approved = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = String(object.BusinessObject);
        }
        else {
            message.BusinessObject = "";
        }
        if (object.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
            object.BaseledgerBusinessObjectIDOfApprovedObject !== null) {
            message.BaseledgerBusinessObjectIDOfApprovedObject = String(object.BaseledgerBusinessObjectIDOfApprovedObject);
        }
        else {
            message.BaseledgerBusinessObjectIDOfApprovedObject = "";
        }
        if (object.Workgroup !== undefined && object.Workgroup !== null) {
            message.Workgroup = String(object.Workgroup);
        }
        else {
            message.Workgroup = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = String(object.Recipient);
        }
        else {
            message.Recipient = "";
        }
        if (object.HashOfObjectToApprove !== undefined &&
            object.HashOfObjectToApprove !== null) {
            message.HashOfObjectToApprove = String(object.HashOfObjectToApprove);
        }
        else {
            message.HashOfObjectToApprove = "";
        }
        if (object.OriginalBaseledgerTransactionID !== undefined &&
            object.OriginalBaseledgerTransactionID !== null) {
            message.OriginalBaseledgerTransactionID = String(object.OriginalBaseledgerTransactionID);
        }
        else {
            message.OriginalBaseledgerTransactionID = "";
        }
        if (object.OriginalOffchainProcessMessageID !== undefined &&
            object.OriginalOffchainProcessMessageID !== null) {
            message.OriginalOffchainProcessMessageID = String(object.OriginalOffchainProcessMessageID);
        }
        else {
            message.OriginalOffchainProcessMessageID = "";
        }
        if (object.FeedbackMessage !== undefined &&
            object.FeedbackMessage !== null) {
            message.FeedbackMessage = String(object.FeedbackMessage);
        }
        else {
            message.FeedbackMessage = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = {
            ...baseMsgUpdateSynchronizationFeedback,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.Approved !== undefined && object.Approved !== null) {
            message.Approved = object.Approved;
        }
        else {
            message.Approved = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = object.BusinessObject;
        }
        else {
            message.BusinessObject = "";
        }
        if (object.BaseledgerBusinessObjectIDOfApprovedObject !== undefined &&
            object.BaseledgerBusinessObjectIDOfApprovedObject !== null) {
            message.BaseledgerBusinessObjectIDOfApprovedObject =
                object.BaseledgerBusinessObjectIDOfApprovedObject;
        }
        else {
            message.BaseledgerBusinessObjectIDOfApprovedObject = "";
        }
        if (object.Workgroup !== undefined && object.Workgroup !== null) {
            message.Workgroup = object.Workgroup;
        }
        else {
            message.Workgroup = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = object.Recipient;
        }
        else {
            message.Recipient = "";
        }
        if (object.HashOfObjectToApprove !== undefined &&
            object.HashOfObjectToApprove !== null) {
            message.HashOfObjectToApprove = object.HashOfObjectToApprove;
        }
        else {
            message.HashOfObjectToApprove = "";
        }
        if (object.OriginalBaseledgerTransactionID !== undefined &&
            object.OriginalBaseledgerTransactionID !== null) {
            message.OriginalBaseledgerTransactionID =
                object.OriginalBaseledgerTransactionID;
        }
        else {
            message.OriginalBaseledgerTransactionID = "";
        }
        if (object.OriginalOffchainProcessMessageID !== undefined &&
            object.OriginalOffchainProcessMessageID !== null) {
            message.OriginalOffchainProcessMessageID =
                object.OriginalOffchainProcessMessageID;
        }
        else {
            message.OriginalOffchainProcessMessageID = "";
        }
        if (object.FeedbackMessage !== undefined &&
            object.FeedbackMessage !== null) {
            message.FeedbackMessage = object.FeedbackMessage;
        }
        else {
            message.FeedbackMessage = "";
        }
        return message;
    },
};
const baseMsgUpdateSynchronizationFeedbackResponse = {};
export const MsgUpdateSynchronizationFeedbackResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateSynchronizationFeedbackResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgUpdateSynchronizationFeedbackResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdateSynchronizationFeedbackResponse,
        };
        return message;
    },
};
const baseMsgDeleteSynchronizationFeedback = { creator: "", id: 0 };
export const MsgDeleteSynchronizationFeedback = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== 0) {
            writer.uint32(16).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteSynchronizationFeedback,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgDeleteSynchronizationFeedback,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgDeleteSynchronizationFeedback,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgDeleteSynchronizationFeedbackResponse = {};
export const MsgDeleteSynchronizationFeedbackResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteSynchronizationFeedbackResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgDeleteSynchronizationFeedbackResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeleteSynchronizationFeedbackResponse,
        };
        return message;
    },
};
const baseMsgCreateSynchronizationRequest = {
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
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateSynchronizationRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseMsgCreateSynchronizationRequest,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.WorkgroupID !== undefined && object.WorkgroupID !== null) {
            message.WorkgroupID = String(object.WorkgroupID);
        }
        else {
            message.WorkgroupID = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = String(object.Recipient);
        }
        else {
            message.Recipient = "";
        }
        if (object.WorkstepType !== undefined && object.WorkstepType !== null) {
            message.WorkstepType = String(object.WorkstepType);
        }
        else {
            message.WorkstepType = "";
        }
        if (object.BusinessObjectType !== undefined &&
            object.BusinessObjectType !== null) {
            message.BusinessObjectType = String(object.BusinessObjectType);
        }
        else {
            message.BusinessObjectType = "";
        }
        if (object.BaseledgerBusinessObjectID !== undefined &&
            object.BaseledgerBusinessObjectID !== null) {
            message.BaseledgerBusinessObjectID = String(object.BaseledgerBusinessObjectID);
        }
        else {
            message.BaseledgerBusinessObjectID = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = String(object.BusinessObject);
        }
        else {
            message.BusinessObject = "";
        }
        if (object.ReferencedBaseledgerBusinessObjectID !== undefined &&
            object.ReferencedBaseledgerBusinessObjectID !== null) {
            message.ReferencedBaseledgerBusinessObjectID = String(object.ReferencedBaseledgerBusinessObjectID);
        }
        else {
            message.ReferencedBaseledgerBusinessObjectID = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = {
            ...baseMsgCreateSynchronizationRequest,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.WorkgroupID !== undefined && object.WorkgroupID !== null) {
            message.WorkgroupID = object.WorkgroupID;
        }
        else {
            message.WorkgroupID = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = object.Recipient;
        }
        else {
            message.Recipient = "";
        }
        if (object.WorkstepType !== undefined && object.WorkstepType !== null) {
            message.WorkstepType = object.WorkstepType;
        }
        else {
            message.WorkstepType = "";
        }
        if (object.BusinessObjectType !== undefined &&
            object.BusinessObjectType !== null) {
            message.BusinessObjectType = object.BusinessObjectType;
        }
        else {
            message.BusinessObjectType = "";
        }
        if (object.BaseledgerBusinessObjectID !== undefined &&
            object.BaseledgerBusinessObjectID !== null) {
            message.BaseledgerBusinessObjectID = object.BaseledgerBusinessObjectID;
        }
        else {
            message.BaseledgerBusinessObjectID = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = object.BusinessObject;
        }
        else {
            message.BusinessObject = "";
        }
        if (object.ReferencedBaseledgerBusinessObjectID !== undefined &&
            object.ReferencedBaseledgerBusinessObjectID !== null) {
            message.ReferencedBaseledgerBusinessObjectID =
                object.ReferencedBaseledgerBusinessObjectID;
        }
        else {
            message.ReferencedBaseledgerBusinessObjectID = "";
        }
        return message;
    },
};
const baseMsgCreateSynchronizationRequestResponse = { id: 0 };
export const MsgCreateSynchronizationRequestResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateSynchronizationRequestResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgCreateSynchronizationRequestResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgCreateSynchronizationRequestResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgUpdateSynchronizationRequest = {
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
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateSynchronizationRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
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
    fromJSON(object) {
        const message = {
            ...baseMsgUpdateSynchronizationRequest,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.WorkgroupID !== undefined && object.WorkgroupID !== null) {
            message.WorkgroupID = String(object.WorkgroupID);
        }
        else {
            message.WorkgroupID = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = String(object.Recipient);
        }
        else {
            message.Recipient = "";
        }
        if (object.WorkstepType !== undefined && object.WorkstepType !== null) {
            message.WorkstepType = String(object.WorkstepType);
        }
        else {
            message.WorkstepType = "";
        }
        if (object.BusinessObjectType !== undefined &&
            object.BusinessObjectType !== null) {
            message.BusinessObjectType = String(object.BusinessObjectType);
        }
        else {
            message.BusinessObjectType = "";
        }
        if (object.BaseledgerBusinessObjectID !== undefined &&
            object.BaseledgerBusinessObjectID !== null) {
            message.BaseledgerBusinessObjectID = String(object.BaseledgerBusinessObjectID);
        }
        else {
            message.BaseledgerBusinessObjectID = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = String(object.BusinessObject);
        }
        else {
            message.BusinessObject = "";
        }
        if (object.ReferencedBaseledgerBusinessObjectID !== undefined &&
            object.ReferencedBaseledgerBusinessObjectID !== null) {
            message.ReferencedBaseledgerBusinessObjectID = String(object.ReferencedBaseledgerBusinessObjectID);
        }
        else {
            message.ReferencedBaseledgerBusinessObjectID = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = {
            ...baseMsgUpdateSynchronizationRequest,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.WorkgroupID !== undefined && object.WorkgroupID !== null) {
            message.WorkgroupID = object.WorkgroupID;
        }
        else {
            message.WorkgroupID = "";
        }
        if (object.Recipient !== undefined && object.Recipient !== null) {
            message.Recipient = object.Recipient;
        }
        else {
            message.Recipient = "";
        }
        if (object.WorkstepType !== undefined && object.WorkstepType !== null) {
            message.WorkstepType = object.WorkstepType;
        }
        else {
            message.WorkstepType = "";
        }
        if (object.BusinessObjectType !== undefined &&
            object.BusinessObjectType !== null) {
            message.BusinessObjectType = object.BusinessObjectType;
        }
        else {
            message.BusinessObjectType = "";
        }
        if (object.BaseledgerBusinessObjectID !== undefined &&
            object.BaseledgerBusinessObjectID !== null) {
            message.BaseledgerBusinessObjectID = object.BaseledgerBusinessObjectID;
        }
        else {
            message.BaseledgerBusinessObjectID = "";
        }
        if (object.BusinessObject !== undefined && object.BusinessObject !== null) {
            message.BusinessObject = object.BusinessObject;
        }
        else {
            message.BusinessObject = "";
        }
        if (object.ReferencedBaseledgerBusinessObjectID !== undefined &&
            object.ReferencedBaseledgerBusinessObjectID !== null) {
            message.ReferencedBaseledgerBusinessObjectID =
                object.ReferencedBaseledgerBusinessObjectID;
        }
        else {
            message.ReferencedBaseledgerBusinessObjectID = "";
        }
        return message;
    },
};
const baseMsgUpdateSynchronizationRequestResponse = {};
export const MsgUpdateSynchronizationRequestResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateSynchronizationRequestResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgUpdateSynchronizationRequestResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdateSynchronizationRequestResponse,
        };
        return message;
    },
};
const baseMsgDeleteSynchronizationRequest = { creator: "", id: 0 };
export const MsgDeleteSynchronizationRequest = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== 0) {
            writer.uint32(16).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteSynchronizationRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgDeleteSynchronizationRequest,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgDeleteSynchronizationRequest,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgDeleteSynchronizationRequestResponse = {};
export const MsgDeleteSynchronizationRequestResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteSynchronizationRequestResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgDeleteSynchronizationRequestResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeleteSynchronizationRequestResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateSynchronizationFeedback(request) {
        const data = MsgCreateSynchronizationFeedback.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Msg", "CreateSynchronizationFeedback", data);
        return promise.then((data) => MsgCreateSynchronizationFeedbackResponse.decode(new Reader(data)));
    }
    UpdateSynchronizationFeedback(request) {
        const data = MsgUpdateSynchronizationFeedback.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Msg", "UpdateSynchronizationFeedback", data);
        return promise.then((data) => MsgUpdateSynchronizationFeedbackResponse.decode(new Reader(data)));
    }
    DeleteSynchronizationFeedback(request) {
        const data = MsgDeleteSynchronizationFeedback.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Msg", "DeleteSynchronizationFeedback", data);
        return promise.then((data) => MsgDeleteSynchronizationFeedbackResponse.decode(new Reader(data)));
    }
    CreateSynchronizationRequest(request) {
        const data = MsgCreateSynchronizationRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Msg", "CreateSynchronizationRequest", data);
        return promise.then((data) => MsgCreateSynchronizationRequestResponse.decode(new Reader(data)));
    }
    UpdateSynchronizationRequest(request) {
        const data = MsgUpdateSynchronizationRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Msg", "UpdateSynchronizationRequest", data);
        return promise.then((data) => MsgUpdateSynchronizationRequestResponse.decode(new Reader(data)));
    }
    DeleteSynchronizationRequest(request) {
        const data = MsgDeleteSynchronizationRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Msg", "DeleteSynchronizationRequest", data);
        return promise.then((data) => MsgDeleteSynchronizationRequestResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}

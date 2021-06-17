/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
export const protobufPackage = 'unibrightio.baseledger.baseledger';
const baseMsgCreateBaseledgerTransaction = { creator: '', baseledgerTransactionId: '', payload: '' };
export const MsgCreateBaseledgerTransaction = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.baseledgerTransactionId !== '') {
            writer.uint32(18).string(message.baseledgerTransactionId);
        }
        if (message.payload !== '') {
            writer.uint32(26).string(message.payload);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateBaseledgerTransaction };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.baseledgerTransactionId = reader.string();
                    break;
                case 3:
                    message.payload = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
            message.baseledgerTransactionId = String(object.baseledgerTransactionId);
        }
        else {
            message.baseledgerTransactionId = '';
        }
        if (object.payload !== undefined && object.payload !== null) {
            message.payload = String(object.payload);
        }
        else {
            message.payload = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.baseledgerTransactionId !== undefined && (obj.baseledgerTransactionId = message.baseledgerTransactionId);
        message.payload !== undefined && (obj.payload = message.payload);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
            message.baseledgerTransactionId = object.baseledgerTransactionId;
        }
        else {
            message.baseledgerTransactionId = '';
        }
        if (object.payload !== undefined && object.payload !== null) {
            message.payload = object.payload;
        }
        else {
            message.payload = '';
        }
        return message;
    }
};
const baseMsgCreateBaseledgerTransactionResponse = { id: 0 };
export const MsgCreateBaseledgerTransactionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateBaseledgerTransactionResponse };
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
        const message = { ...baseMsgCreateBaseledgerTransactionResponse };
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
        const message = { ...baseMsgCreateBaseledgerTransactionResponse };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    }
};
const baseMsgUpdateBaseledgerTransaction = { creator: '', id: 0, baseledgerTransactionId: '', payload: '' };
export const MsgUpdateBaseledgerTransaction = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== 0) {
            writer.uint32(16).uint64(message.id);
        }
        if (message.baseledgerTransactionId !== '') {
            writer.uint32(26).string(message.baseledgerTransactionId);
        }
        if (message.payload !== '') {
            writer.uint32(34).string(message.payload);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateBaseledgerTransaction };
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
                    message.baseledgerTransactionId = reader.string();
                    break;
                case 4:
                    message.payload = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
            message.baseledgerTransactionId = String(object.baseledgerTransactionId);
        }
        else {
            message.baseledgerTransactionId = '';
        }
        if (object.payload !== undefined && object.payload !== null) {
            message.payload = String(object.payload);
        }
        else {
            message.payload = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        message.baseledgerTransactionId !== undefined && (obj.baseledgerTransactionId = message.baseledgerTransactionId);
        message.payload !== undefined && (obj.payload = message.payload);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
            message.baseledgerTransactionId = object.baseledgerTransactionId;
        }
        else {
            message.baseledgerTransactionId = '';
        }
        if (object.payload !== undefined && object.payload !== null) {
            message.payload = object.payload;
        }
        else {
            message.payload = '';
        }
        return message;
    }
};
const baseMsgUpdateBaseledgerTransactionResponse = {};
export const MsgUpdateBaseledgerTransactionResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateBaseledgerTransactionResponse };
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
        const message = { ...baseMsgUpdateBaseledgerTransactionResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdateBaseledgerTransactionResponse };
        return message;
    }
};
const baseMsgDeleteBaseledgerTransaction = { creator: '', id: 0 };
export const MsgDeleteBaseledgerTransaction = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
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
        const message = { ...baseMsgDeleteBaseledgerTransaction };
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
        const message = { ...baseMsgDeleteBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
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
        const message = { ...baseMsgDeleteBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    }
};
const baseMsgDeleteBaseledgerTransactionResponse = {};
export const MsgDeleteBaseledgerTransactionResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteBaseledgerTransactionResponse };
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
        const message = { ...baseMsgDeleteBaseledgerTransactionResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDeleteBaseledgerTransactionResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateBaseledgerTransaction(request) {
        const data = MsgCreateBaseledgerTransaction.encode(request).finish();
        const promise = this.rpc.request('unibrightio.baseledger.baseledger.Msg', 'CreateBaseledgerTransaction', data);
        return promise.then((data) => MsgCreateBaseledgerTransactionResponse.decode(new Reader(data)));
    }
    UpdateBaseledgerTransaction(request) {
        const data = MsgUpdateBaseledgerTransaction.encode(request).finish();
        const promise = this.rpc.request('unibrightio.baseledger.baseledger.Msg', 'UpdateBaseledgerTransaction', data);
        return promise.then((data) => MsgUpdateBaseledgerTransactionResponse.decode(new Reader(data)));
    }
    DeleteBaseledgerTransaction(request) {
        const data = MsgDeleteBaseledgerTransaction.encode(request).finish();
        const promise = this.rpc.request('unibrightio.baseledger.baseledger.Msg', 'DeleteBaseledgerTransaction', data);
        return promise.then((data) => MsgDeleteBaseledgerTransactionResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}

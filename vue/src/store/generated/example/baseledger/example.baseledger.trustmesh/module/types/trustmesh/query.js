/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { SynchronizationFeedback } from "../trustmesh/SynchronizationFeedback";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { SynchronizationRequest } from "../trustmesh/SynchronizationRequest";
export const protobufPackage = "example.baseledger.trustmesh";
const baseQueryGetSynchronizationFeedbackRequest = { id: 0 };
export const QueryGetSynchronizationFeedbackRequest = {
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
            ...baseQueryGetSynchronizationFeedbackRequest,
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
            ...baseQueryGetSynchronizationFeedbackRequest,
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
            ...baseQueryGetSynchronizationFeedbackRequest,
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
const baseQueryGetSynchronizationFeedbackResponse = {};
export const QueryGetSynchronizationFeedbackResponse = {
    encode(message, writer = Writer.create()) {
        if (message.SynchronizationFeedback !== undefined) {
            SynchronizationFeedback.encode(message.SynchronizationFeedback, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetSynchronizationFeedbackResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SynchronizationFeedback = SynchronizationFeedback.decode(reader, reader.uint32());
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
            ...baseQueryGetSynchronizationFeedbackResponse,
        };
        if (object.SynchronizationFeedback !== undefined &&
            object.SynchronizationFeedback !== null) {
            message.SynchronizationFeedback = SynchronizationFeedback.fromJSON(object.SynchronizationFeedback);
        }
        else {
            message.SynchronizationFeedback = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.SynchronizationFeedback !== undefined &&
            (obj.SynchronizationFeedback = message.SynchronizationFeedback
                ? SynchronizationFeedback.toJSON(message.SynchronizationFeedback)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetSynchronizationFeedbackResponse,
        };
        if (object.SynchronizationFeedback !== undefined &&
            object.SynchronizationFeedback !== null) {
            message.SynchronizationFeedback = SynchronizationFeedback.fromPartial(object.SynchronizationFeedback);
        }
        else {
            message.SynchronizationFeedback = undefined;
        }
        return message;
    },
};
const baseQueryAllSynchronizationFeedbackRequest = {};
export const QueryAllSynchronizationFeedbackRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllSynchronizationFeedbackRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
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
            ...baseQueryAllSynchronizationFeedbackRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSynchronizationFeedbackRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllSynchronizationFeedbackResponse = {};
export const QueryAllSynchronizationFeedbackResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.SynchronizationFeedback) {
            SynchronizationFeedback.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllSynchronizationFeedbackResponse,
        };
        message.SynchronizationFeedback = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SynchronizationFeedback.push(SynchronizationFeedback.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
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
            ...baseQueryAllSynchronizationFeedbackResponse,
        };
        message.SynchronizationFeedback = [];
        if (object.SynchronizationFeedback !== undefined &&
            object.SynchronizationFeedback !== null) {
            for (const e of object.SynchronizationFeedback) {
                message.SynchronizationFeedback.push(SynchronizationFeedback.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.SynchronizationFeedback) {
            obj.SynchronizationFeedback = message.SynchronizationFeedback.map((e) => e ? SynchronizationFeedback.toJSON(e) : undefined);
        }
        else {
            obj.SynchronizationFeedback = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSynchronizationFeedbackResponse,
        };
        message.SynchronizationFeedback = [];
        if (object.SynchronizationFeedback !== undefined &&
            object.SynchronizationFeedback !== null) {
            for (const e of object.SynchronizationFeedback) {
                message.SynchronizationFeedback.push(SynchronizationFeedback.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryGetSynchronizationRequestRequest = { id: 0 };
export const QueryGetSynchronizationRequestRequest = {
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
            ...baseQueryGetSynchronizationRequestRequest,
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
            ...baseQueryGetSynchronizationRequestRequest,
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
            ...baseQueryGetSynchronizationRequestRequest,
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
const baseQueryGetSynchronizationRequestResponse = {};
export const QueryGetSynchronizationRequestResponse = {
    encode(message, writer = Writer.create()) {
        if (message.SynchronizationRequest !== undefined) {
            SynchronizationRequest.encode(message.SynchronizationRequest, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetSynchronizationRequestResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SynchronizationRequest = SynchronizationRequest.decode(reader, reader.uint32());
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
            ...baseQueryGetSynchronizationRequestResponse,
        };
        if (object.SynchronizationRequest !== undefined &&
            object.SynchronizationRequest !== null) {
            message.SynchronizationRequest = SynchronizationRequest.fromJSON(object.SynchronizationRequest);
        }
        else {
            message.SynchronizationRequest = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.SynchronizationRequest !== undefined &&
            (obj.SynchronizationRequest = message.SynchronizationRequest
                ? SynchronizationRequest.toJSON(message.SynchronizationRequest)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetSynchronizationRequestResponse,
        };
        if (object.SynchronizationRequest !== undefined &&
            object.SynchronizationRequest !== null) {
            message.SynchronizationRequest = SynchronizationRequest.fromPartial(object.SynchronizationRequest);
        }
        else {
            message.SynchronizationRequest = undefined;
        }
        return message;
    },
};
const baseQueryAllSynchronizationRequestRequest = {};
export const QueryAllSynchronizationRequestRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllSynchronizationRequestRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
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
            ...baseQueryAllSynchronizationRequestRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSynchronizationRequestRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllSynchronizationRequestResponse = {};
export const QueryAllSynchronizationRequestResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.SynchronizationRequest) {
            SynchronizationRequest.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllSynchronizationRequestResponse,
        };
        message.SynchronizationRequest = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SynchronizationRequest.push(SynchronizationRequest.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
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
            ...baseQueryAllSynchronizationRequestResponse,
        };
        message.SynchronizationRequest = [];
        if (object.SynchronizationRequest !== undefined &&
            object.SynchronizationRequest !== null) {
            for (const e of object.SynchronizationRequest) {
                message.SynchronizationRequest.push(SynchronizationRequest.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.SynchronizationRequest) {
            obj.SynchronizationRequest = message.SynchronizationRequest.map((e) => e ? SynchronizationRequest.toJSON(e) : undefined);
        }
        else {
            obj.SynchronizationRequest = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSynchronizationRequestResponse,
        };
        message.SynchronizationRequest = [];
        if (object.SynchronizationRequest !== undefined &&
            object.SynchronizationRequest !== null) {
            for (const e of object.SynchronizationRequest) {
                message.SynchronizationRequest.push(SynchronizationRequest.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    SynchronizationFeedback(request) {
        const data = QueryGetSynchronizationFeedbackRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Query", "SynchronizationFeedback", data);
        return promise.then((data) => QueryGetSynchronizationFeedbackResponse.decode(new Reader(data)));
    }
    SynchronizationFeedbackAll(request) {
        const data = QueryAllSynchronizationFeedbackRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Query", "SynchronizationFeedbackAll", data);
        return promise.then((data) => QueryAllSynchronizationFeedbackResponse.decode(new Reader(data)));
    }
    SynchronizationRequest(request) {
        const data = QueryGetSynchronizationRequestRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Query", "SynchronizationRequest", data);
        return promise.then((data) => QueryGetSynchronizationRequestResponse.decode(new Reader(data)));
    }
    SynchronizationRequestAll(request) {
        const data = QueryAllSynchronizationRequestRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.trustmesh.Query", "SynchronizationRequestAll", data);
        return promise.then((data) => QueryAllSynchronizationRequestResponse.decode(new Reader(data)));
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

/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { BaseledgerTransaction } from "../baseledger/BaseledgerTransaction";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
export const protobufPackage = "example.baseledger.baseledger";
const baseQueryGetBaseledgerTransactionRequest = { id: 0 };
export const QueryGetBaseledgerTransactionRequest = {
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
            ...baseQueryGetBaseledgerTransactionRequest,
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
            ...baseQueryGetBaseledgerTransactionRequest,
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
            ...baseQueryGetBaseledgerTransactionRequest,
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
const baseQueryGetBaseledgerTransactionResponse = {};
export const QueryGetBaseledgerTransactionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.BaseledgerTransaction !== undefined) {
            BaseledgerTransaction.encode(message.BaseledgerTransaction, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetBaseledgerTransactionResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.BaseledgerTransaction = BaseledgerTransaction.decode(reader, reader.uint32());
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
            ...baseQueryGetBaseledgerTransactionResponse,
        };
        if (object.BaseledgerTransaction !== undefined &&
            object.BaseledgerTransaction !== null) {
            message.BaseledgerTransaction = BaseledgerTransaction.fromJSON(object.BaseledgerTransaction);
        }
        else {
            message.BaseledgerTransaction = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.BaseledgerTransaction !== undefined &&
            (obj.BaseledgerTransaction = message.BaseledgerTransaction
                ? BaseledgerTransaction.toJSON(message.BaseledgerTransaction)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetBaseledgerTransactionResponse,
        };
        if (object.BaseledgerTransaction !== undefined &&
            object.BaseledgerTransaction !== null) {
            message.BaseledgerTransaction = BaseledgerTransaction.fromPartial(object.BaseledgerTransaction);
        }
        else {
            message.BaseledgerTransaction = undefined;
        }
        return message;
    },
};
const baseQueryAllBaseledgerTransactionRequest = {};
export const QueryAllBaseledgerTransactionRequest = {
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
            ...baseQueryAllBaseledgerTransactionRequest,
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
            ...baseQueryAllBaseledgerTransactionRequest,
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
            ...baseQueryAllBaseledgerTransactionRequest,
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
const baseQueryAllBaseledgerTransactionResponse = {};
export const QueryAllBaseledgerTransactionResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.BaseledgerTransaction) {
            BaseledgerTransaction.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllBaseledgerTransactionResponse,
        };
        message.BaseledgerTransaction = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.BaseledgerTransaction.push(BaseledgerTransaction.decode(reader, reader.uint32()));
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
            ...baseQueryAllBaseledgerTransactionResponse,
        };
        message.BaseledgerTransaction = [];
        if (object.BaseledgerTransaction !== undefined &&
            object.BaseledgerTransaction !== null) {
            for (const e of object.BaseledgerTransaction) {
                message.BaseledgerTransaction.push(BaseledgerTransaction.fromJSON(e));
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
        if (message.BaseledgerTransaction) {
            obj.BaseledgerTransaction = message.BaseledgerTransaction.map((e) => e ? BaseledgerTransaction.toJSON(e) : undefined);
        }
        else {
            obj.BaseledgerTransaction = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllBaseledgerTransactionResponse,
        };
        message.BaseledgerTransaction = [];
        if (object.BaseledgerTransaction !== undefined &&
            object.BaseledgerTransaction !== null) {
            for (const e of object.BaseledgerTransaction) {
                message.BaseledgerTransaction.push(BaseledgerTransaction.fromPartial(e));
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
    BaseledgerTransaction(request) {
        const data = QueryGetBaseledgerTransactionRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.baseledger.Query", "BaseledgerTransaction", data);
        return promise.then((data) => QueryGetBaseledgerTransactionResponse.decode(new Reader(data)));
    }
    BaseledgerTransactionAll(request) {
        const data = QueryAllBaseledgerTransactionRequest.encode(request).finish();
        const promise = this.rpc.request("example.baseledger.baseledger.Query", "BaseledgerTransactionAll", data);
        return promise.then((data) => QueryAllBaseledgerTransactionResponse.decode(new Reader(data)));
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

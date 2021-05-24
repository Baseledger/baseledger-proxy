/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
import { BaseledgerTransaction } from '../baseledger/BaseledgerTransaction';
export const protobufPackage = 'unibrightio.baseledger.baseledger';
const baseGenesisState = { BaseledgerTransactionCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.BaseledgerTransactionList) {
            BaseledgerTransaction.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.BaseledgerTransactionCount !== 0) {
            writer.uint32(16).uint64(message.BaseledgerTransactionCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.BaseledgerTransactionList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.BaseledgerTransactionList.push(BaseledgerTransaction.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.BaseledgerTransactionCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.BaseledgerTransactionList = [];
        if (object.BaseledgerTransactionList !== undefined && object.BaseledgerTransactionList !== null) {
            for (const e of object.BaseledgerTransactionList) {
                message.BaseledgerTransactionList.push(BaseledgerTransaction.fromJSON(e));
            }
        }
        if (object.BaseledgerTransactionCount !== undefined && object.BaseledgerTransactionCount !== null) {
            message.BaseledgerTransactionCount = Number(object.BaseledgerTransactionCount);
        }
        else {
            message.BaseledgerTransactionCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.BaseledgerTransactionList) {
            obj.BaseledgerTransactionList = message.BaseledgerTransactionList.map((e) => (e ? BaseledgerTransaction.toJSON(e) : undefined));
        }
        else {
            obj.BaseledgerTransactionList = [];
        }
        message.BaseledgerTransactionCount !== undefined && (obj.BaseledgerTransactionCount = message.BaseledgerTransactionCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.BaseledgerTransactionList = [];
        if (object.BaseledgerTransactionList !== undefined && object.BaseledgerTransactionList !== null) {
            for (const e of object.BaseledgerTransactionList) {
                message.BaseledgerTransactionList.push(BaseledgerTransaction.fromPartial(e));
            }
        }
        if (object.BaseledgerTransactionCount !== undefined && object.BaseledgerTransactionCount !== null) {
            message.BaseledgerTransactionCount = object.BaseledgerTransactionCount;
        }
        else {
            message.BaseledgerTransactionCount = 0;
        }
        return message;
    }
};
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

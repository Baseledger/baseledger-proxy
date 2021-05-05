/* eslint-disable */
import { BaseledgerTransaction } from "../baseledger/BaseledgerTransaction";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "example.baseledger.baseledger";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.BaseledgerTransactionList) {
            BaseledgerTransaction.encode(v, writer.uint32(10).fork()).ldelim();
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
        if (object.BaseledgerTransactionList !== undefined &&
            object.BaseledgerTransactionList !== null) {
            for (const e of object.BaseledgerTransactionList) {
                message.BaseledgerTransactionList.push(BaseledgerTransaction.fromJSON(e));
            }
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
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.BaseledgerTransactionList = [];
        if (object.BaseledgerTransactionList !== undefined &&
            object.BaseledgerTransactionList !== null) {
            for (const e of object.BaseledgerTransactionList) {
                message.BaseledgerTransactionList.push(BaseledgerTransaction.fromPartial(e));
            }
        }
        return message;
    },
};

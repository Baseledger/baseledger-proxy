/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'unibrightio.baseledger.baseledger';
const baseBaseledgerTransaction = { creator: '', id: '', baseledgerTransactionId: '', payload: '' };
export const BaseledgerTransaction = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== '') {
            writer.uint32(18).string(message.id);
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
        const message = { ...baseBaseledgerTransaction };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = reader.string();
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
        const message = { ...baseBaseledgerTransaction };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = '';
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
        const message = { ...baseBaseledgerTransaction };
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
            message.id = '';
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

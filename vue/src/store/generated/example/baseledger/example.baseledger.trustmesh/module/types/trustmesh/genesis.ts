/* eslint-disable */
import { SynchronizationFeedback } from "../trustmesh/SynchronizationFeedback";
import { SynchronizationRequest } from "../trustmesh/SynchronizationRequest";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "example.baseledger.trustmesh";

/** GenesisState defines the trustmesh module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  SynchronizationFeedbackList: SynchronizationFeedback[];
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  SynchronizationRequestList: SynchronizationRequest[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.SynchronizationFeedbackList) {
      SynchronizationFeedback.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.SynchronizationRequestList) {
      SynchronizationRequest.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.SynchronizationFeedbackList = [];
    message.SynchronizationRequestList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.SynchronizationFeedbackList.push(
            SynchronizationFeedback.decode(reader, reader.uint32())
          );
          break;
        case 1:
          message.SynchronizationRequestList.push(
            SynchronizationRequest.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.SynchronizationFeedbackList = [];
    message.SynchronizationRequestList = [];
    if (
      object.SynchronizationFeedbackList !== undefined &&
      object.SynchronizationFeedbackList !== null
    ) {
      for (const e of object.SynchronizationFeedbackList) {
        message.SynchronizationFeedbackList.push(
          SynchronizationFeedback.fromJSON(e)
        );
      }
    }
    if (
      object.SynchronizationRequestList !== undefined &&
      object.SynchronizationRequestList !== null
    ) {
      for (const e of object.SynchronizationRequestList) {
        message.SynchronizationRequestList.push(
          SynchronizationRequest.fromJSON(e)
        );
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.SynchronizationFeedbackList) {
      obj.SynchronizationFeedbackList = message.SynchronizationFeedbackList.map(
        (e) => (e ? SynchronizationFeedback.toJSON(e) : undefined)
      );
    } else {
      obj.SynchronizationFeedbackList = [];
    }
    if (message.SynchronizationRequestList) {
      obj.SynchronizationRequestList = message.SynchronizationRequestList.map(
        (e) => (e ? SynchronizationRequest.toJSON(e) : undefined)
      );
    } else {
      obj.SynchronizationRequestList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.SynchronizationFeedbackList = [];
    message.SynchronizationRequestList = [];
    if (
      object.SynchronizationFeedbackList !== undefined &&
      object.SynchronizationFeedbackList !== null
    ) {
      for (const e of object.SynchronizationFeedbackList) {
        message.SynchronizationFeedbackList.push(
          SynchronizationFeedback.fromPartial(e)
        );
      }
    }
    if (
      object.SynchronizationRequestList !== undefined &&
      object.SynchronizationRequestList !== null
    ) {
      for (const e of object.SynchronizationRequestList) {
        message.SynchronizationRequestList.push(
          SynchronizationRequest.fromPartial(e)
        );
      }
    }
    return message;
  },
};

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

/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { BaseledgerTransaction } from '../baseledger/BaseledgerTransaction'

export const protobufPackage = 'unibrightio.baseledger.baseledger'

/** GenesisState defines the baseledger module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  BaseledgerTransactionList: BaseledgerTransaction[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  BaseledgerTransactionCount: number
}

const baseGenesisState: object = { BaseledgerTransactionCount: 0 }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.BaseledgerTransactionList) {
      BaseledgerTransaction.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.BaseledgerTransactionCount !== 0) {
      writer.uint32(16).uint64(message.BaseledgerTransactionCount)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.BaseledgerTransactionList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.BaseledgerTransactionList.push(BaseledgerTransaction.decode(reader, reader.uint32()))
          break
        case 2:
          message.BaseledgerTransactionCount = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.BaseledgerTransactionList = []
    if (object.BaseledgerTransactionList !== undefined && object.BaseledgerTransactionList !== null) {
      for (const e of object.BaseledgerTransactionList) {
        message.BaseledgerTransactionList.push(BaseledgerTransaction.fromJSON(e))
      }
    }
    if (object.BaseledgerTransactionCount !== undefined && object.BaseledgerTransactionCount !== null) {
      message.BaseledgerTransactionCount = Number(object.BaseledgerTransactionCount)
    } else {
      message.BaseledgerTransactionCount = 0
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.BaseledgerTransactionList) {
      obj.BaseledgerTransactionList = message.BaseledgerTransactionList.map((e) => (e ? BaseledgerTransaction.toJSON(e) : undefined))
    } else {
      obj.BaseledgerTransactionList = []
    }
    message.BaseledgerTransactionCount !== undefined && (obj.BaseledgerTransactionCount = message.BaseledgerTransactionCount)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.BaseledgerTransactionList = []
    if (object.BaseledgerTransactionList !== undefined && object.BaseledgerTransactionList !== null) {
      for (const e of object.BaseledgerTransactionList) {
        message.BaseledgerTransactionList.push(BaseledgerTransaction.fromPartial(e))
      }
    }
    if (object.BaseledgerTransactionCount !== undefined && object.BaseledgerTransactionCount !== null) {
      message.BaseledgerTransactionCount = object.BaseledgerTransactionCount
    } else {
      message.BaseledgerTransactionCount = 0
    }
    return message
  }
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}

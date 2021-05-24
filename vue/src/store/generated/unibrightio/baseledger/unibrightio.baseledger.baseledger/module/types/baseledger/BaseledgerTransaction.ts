/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'unibrightio.baseledger.baseledger'

export interface BaseledgerTransaction {
  creator: string
  id: number
  baseId: string
  payload: string
}

const baseBaseledgerTransaction: object = { creator: '', id: 0, baseId: '', payload: '' }

export const BaseledgerTransaction = {
  encode(message: BaseledgerTransaction, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.baseId !== '') {
      writer.uint32(26).string(message.baseId)
    }
    if (message.payload !== '') {
      writer.uint32(34).string(message.payload)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): BaseledgerTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseBaseledgerTransaction } as BaseledgerTransaction
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.baseId = reader.string()
          break
        case 4:
          message.payload = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): BaseledgerTransaction {
    const message = { ...baseBaseledgerTransaction } as BaseledgerTransaction
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.baseId !== undefined && object.baseId !== null) {
      message.baseId = String(object.baseId)
    } else {
      message.baseId = ''
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = String(object.payload)
    } else {
      message.payload = ''
    }
    return message
  },

  toJSON(message: BaseledgerTransaction): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.baseId !== undefined && (obj.baseId = message.baseId)
    message.payload !== undefined && (obj.payload = message.payload)
    return obj
  },

  fromPartial(object: DeepPartial<BaseledgerTransaction>): BaseledgerTransaction {
    const message = { ...baseBaseledgerTransaction } as BaseledgerTransaction
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.baseId !== undefined && object.baseId !== null) {
      message.baseId = object.baseId
    } else {
      message.baseId = ''
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = object.payload
    } else {
      message.payload = ''
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

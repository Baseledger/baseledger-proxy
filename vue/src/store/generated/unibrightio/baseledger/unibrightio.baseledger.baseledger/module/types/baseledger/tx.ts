/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'unibrightio.baseledger.baseledger'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateBaseledgerTransaction {
  creator: string
  baseledgerTransactionId: string
  payload: string
}

export interface MsgCreateBaseledgerTransactionResponse {
  id: number
}

export interface MsgUpdateBaseledgerTransaction {
  creator: string
  id: number
  baseledgerTransactionId: string
  payload: string
}

export interface MsgUpdateBaseledgerTransactionResponse {}

export interface MsgDeleteBaseledgerTransaction {
  creator: string
  id: number
}

export interface MsgDeleteBaseledgerTransactionResponse {}

const baseMsgCreateBaseledgerTransaction: object = { creator: '', baseledgerTransactionId: '', payload: '' }

export const MsgCreateBaseledgerTransaction = {
  encode(message: MsgCreateBaseledgerTransaction, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.baseledgerTransactionId !== '') {
      writer.uint32(18).string(message.baseledgerTransactionId)
    }
    if (message.payload !== '') {
      writer.uint32(26).string(message.payload)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBaseledgerTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateBaseledgerTransaction } as MsgCreateBaseledgerTransaction
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.baseledgerTransactionId = reader.string()
          break
        case 3:
          message.payload = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateBaseledgerTransaction {
    const message = { ...baseMsgCreateBaseledgerTransaction } as MsgCreateBaseledgerTransaction
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
      message.baseledgerTransactionId = String(object.baseledgerTransactionId)
    } else {
      message.baseledgerTransactionId = ''
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = String(object.payload)
    } else {
      message.payload = ''
    }
    return message
  },

  toJSON(message: MsgCreateBaseledgerTransaction): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.baseledgerTransactionId !== undefined && (obj.baseledgerTransactionId = message.baseledgerTransactionId)
    message.payload !== undefined && (obj.payload = message.payload)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateBaseledgerTransaction>): MsgCreateBaseledgerTransaction {
    const message = { ...baseMsgCreateBaseledgerTransaction } as MsgCreateBaseledgerTransaction
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
      message.baseledgerTransactionId = object.baseledgerTransactionId
    } else {
      message.baseledgerTransactionId = ''
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = object.payload
    } else {
      message.payload = ''
    }
    return message
  }
}

const baseMsgCreateBaseledgerTransactionResponse: object = { id: 0 }

export const MsgCreateBaseledgerTransactionResponse = {
  encode(message: MsgCreateBaseledgerTransactionResponse, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBaseledgerTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateBaseledgerTransactionResponse } as MsgCreateBaseledgerTransactionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateBaseledgerTransactionResponse {
    const message = { ...baseMsgCreateBaseledgerTransactionResponse } as MsgCreateBaseledgerTransactionResponse
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: MsgCreateBaseledgerTransactionResponse): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateBaseledgerTransactionResponse>): MsgCreateBaseledgerTransactionResponse {
    const message = { ...baseMsgCreateBaseledgerTransactionResponse } as MsgCreateBaseledgerTransactionResponse
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseMsgUpdateBaseledgerTransaction: object = { creator: '', id: 0, baseledgerTransactionId: '', payload: '' }

export const MsgUpdateBaseledgerTransaction = {
  encode(message: MsgUpdateBaseledgerTransaction, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.baseledgerTransactionId !== '') {
      writer.uint32(26).string(message.baseledgerTransactionId)
    }
    if (message.payload !== '') {
      writer.uint32(34).string(message.payload)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBaseledgerTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateBaseledgerTransaction } as MsgUpdateBaseledgerTransaction
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
          message.baseledgerTransactionId = reader.string()
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

  fromJSON(object: any): MsgUpdateBaseledgerTransaction {
    const message = { ...baseMsgUpdateBaseledgerTransaction } as MsgUpdateBaseledgerTransaction
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
    if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
      message.baseledgerTransactionId = String(object.baseledgerTransactionId)
    } else {
      message.baseledgerTransactionId = ''
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = String(object.payload)
    } else {
      message.payload = ''
    }
    return message
  },

  toJSON(message: MsgUpdateBaseledgerTransaction): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.baseledgerTransactionId !== undefined && (obj.baseledgerTransactionId = message.baseledgerTransactionId)
    message.payload !== undefined && (obj.payload = message.payload)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateBaseledgerTransaction>): MsgUpdateBaseledgerTransaction {
    const message = { ...baseMsgUpdateBaseledgerTransaction } as MsgUpdateBaseledgerTransaction
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
    if (object.baseledgerTransactionId !== undefined && object.baseledgerTransactionId !== null) {
      message.baseledgerTransactionId = object.baseledgerTransactionId
    } else {
      message.baseledgerTransactionId = ''
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = object.payload
    } else {
      message.payload = ''
    }
    return message
  }
}

const baseMsgUpdateBaseledgerTransactionResponse: object = {}

export const MsgUpdateBaseledgerTransactionResponse = {
  encode(_: MsgUpdateBaseledgerTransactionResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBaseledgerTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateBaseledgerTransactionResponse } as MsgUpdateBaseledgerTransactionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateBaseledgerTransactionResponse {
    const message = { ...baseMsgUpdateBaseledgerTransactionResponse } as MsgUpdateBaseledgerTransactionResponse
    return message
  },

  toJSON(_: MsgUpdateBaseledgerTransactionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateBaseledgerTransactionResponse>): MsgUpdateBaseledgerTransactionResponse {
    const message = { ...baseMsgUpdateBaseledgerTransactionResponse } as MsgUpdateBaseledgerTransactionResponse
    return message
  }
}

const baseMsgDeleteBaseledgerTransaction: object = { creator: '', id: 0 }

export const MsgDeleteBaseledgerTransaction = {
  encode(message: MsgDeleteBaseledgerTransaction, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBaseledgerTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteBaseledgerTransaction } as MsgDeleteBaseledgerTransaction
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeleteBaseledgerTransaction {
    const message = { ...baseMsgDeleteBaseledgerTransaction } as MsgDeleteBaseledgerTransaction
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
    return message
  },

  toJSON(message: MsgDeleteBaseledgerTransaction): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteBaseledgerTransaction>): MsgDeleteBaseledgerTransaction {
    const message = { ...baseMsgDeleteBaseledgerTransaction } as MsgDeleteBaseledgerTransaction
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
    return message
  }
}

const baseMsgDeleteBaseledgerTransactionResponse: object = {}

export const MsgDeleteBaseledgerTransactionResponse = {
  encode(_: MsgDeleteBaseledgerTransactionResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBaseledgerTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteBaseledgerTransactionResponse } as MsgDeleteBaseledgerTransactionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgDeleteBaseledgerTransactionResponse {
    const message = { ...baseMsgDeleteBaseledgerTransactionResponse } as MsgDeleteBaseledgerTransactionResponse
    return message
  },

  toJSON(_: MsgDeleteBaseledgerTransactionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteBaseledgerTransactionResponse>): MsgDeleteBaseledgerTransactionResponse {
    const message = { ...baseMsgDeleteBaseledgerTransactionResponse } as MsgDeleteBaseledgerTransactionResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateBaseledgerTransaction(request: MsgCreateBaseledgerTransaction): Promise<MsgCreateBaseledgerTransactionResponse>
  UpdateBaseledgerTransaction(request: MsgUpdateBaseledgerTransaction): Promise<MsgUpdateBaseledgerTransactionResponse>
  DeleteBaseledgerTransaction(request: MsgDeleteBaseledgerTransaction): Promise<MsgDeleteBaseledgerTransactionResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateBaseledgerTransaction(request: MsgCreateBaseledgerTransaction): Promise<MsgCreateBaseledgerTransactionResponse> {
    const data = MsgCreateBaseledgerTransaction.encode(request).finish()
    const promise = this.rpc.request('unibrightio.baseledger.baseledger.Msg', 'CreateBaseledgerTransaction', data)
    return promise.then((data) => MsgCreateBaseledgerTransactionResponse.decode(new Reader(data)))
  }

  UpdateBaseledgerTransaction(request: MsgUpdateBaseledgerTransaction): Promise<MsgUpdateBaseledgerTransactionResponse> {
    const data = MsgUpdateBaseledgerTransaction.encode(request).finish()
    const promise = this.rpc.request('unibrightio.baseledger.baseledger.Msg', 'UpdateBaseledgerTransaction', data)
    return promise.then((data) => MsgUpdateBaseledgerTransactionResponse.decode(new Reader(data)))
  }

  DeleteBaseledgerTransaction(request: MsgDeleteBaseledgerTransaction): Promise<MsgDeleteBaseledgerTransactionResponse> {
    const data = MsgDeleteBaseledgerTransaction.encode(request).finish()
    const promise = this.rpc.request('unibrightio.baseledger.baseledger.Msg', 'DeleteBaseledgerTransaction', data)
    return promise.then((data) => MsgDeleteBaseledgerTransactionResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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

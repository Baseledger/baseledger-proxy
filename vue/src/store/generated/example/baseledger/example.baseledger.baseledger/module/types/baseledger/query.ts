/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { BaseledgerTransaction } from "../baseledger/BaseledgerTransaction";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "example.baseledger.baseledger";

/** this line is used by starport scaffolding # 3 */
export interface QueryGetBaseledgerTransactionRequest {
  id: number;
}

export interface QueryGetBaseledgerTransactionResponse {
  BaseledgerTransaction: BaseledgerTransaction | undefined;
}

export interface QueryAllBaseledgerTransactionRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBaseledgerTransactionResponse {
  BaseledgerTransaction: BaseledgerTransaction[];
  pagination: PageResponse | undefined;
}

const baseQueryGetBaseledgerTransactionRequest: object = { id: 0 };

export const QueryGetBaseledgerTransactionRequest = {
  encode(
    message: QueryGetBaseledgerTransactionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetBaseledgerTransactionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBaseledgerTransactionRequest,
    } as QueryGetBaseledgerTransactionRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBaseledgerTransactionRequest {
    const message = {
      ...baseQueryGetBaseledgerTransactionRequest,
    } as QueryGetBaseledgerTransactionRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetBaseledgerTransactionRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBaseledgerTransactionRequest>
  ): QueryGetBaseledgerTransactionRequest {
    const message = {
      ...baseQueryGetBaseledgerTransactionRequest,
    } as QueryGetBaseledgerTransactionRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetBaseledgerTransactionResponse: object = {};

export const QueryGetBaseledgerTransactionResponse = {
  encode(
    message: QueryGetBaseledgerTransactionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.BaseledgerTransaction !== undefined) {
      BaseledgerTransaction.encode(
        message.BaseledgerTransaction,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetBaseledgerTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBaseledgerTransactionResponse,
    } as QueryGetBaseledgerTransactionResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.BaseledgerTransaction = BaseledgerTransaction.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBaseledgerTransactionResponse {
    const message = {
      ...baseQueryGetBaseledgerTransactionResponse,
    } as QueryGetBaseledgerTransactionResponse;
    if (
      object.BaseledgerTransaction !== undefined &&
      object.BaseledgerTransaction !== null
    ) {
      message.BaseledgerTransaction = BaseledgerTransaction.fromJSON(
        object.BaseledgerTransaction
      );
    } else {
      message.BaseledgerTransaction = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBaseledgerTransactionResponse): unknown {
    const obj: any = {};
    message.BaseledgerTransaction !== undefined &&
      (obj.BaseledgerTransaction = message.BaseledgerTransaction
        ? BaseledgerTransaction.toJSON(message.BaseledgerTransaction)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBaseledgerTransactionResponse>
  ): QueryGetBaseledgerTransactionResponse {
    const message = {
      ...baseQueryGetBaseledgerTransactionResponse,
    } as QueryGetBaseledgerTransactionResponse;
    if (
      object.BaseledgerTransaction !== undefined &&
      object.BaseledgerTransaction !== null
    ) {
      message.BaseledgerTransaction = BaseledgerTransaction.fromPartial(
        object.BaseledgerTransaction
      );
    } else {
      message.BaseledgerTransaction = undefined;
    }
    return message;
  },
};

const baseQueryAllBaseledgerTransactionRequest: object = {};

export const QueryAllBaseledgerTransactionRequest = {
  encode(
    message: QueryAllBaseledgerTransactionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllBaseledgerTransactionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBaseledgerTransactionRequest,
    } as QueryAllBaseledgerTransactionRequest;
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

  fromJSON(object: any): QueryAllBaseledgerTransactionRequest {
    const message = {
      ...baseQueryAllBaseledgerTransactionRequest,
    } as QueryAllBaseledgerTransactionRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBaseledgerTransactionRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBaseledgerTransactionRequest>
  ): QueryAllBaseledgerTransactionRequest {
    const message = {
      ...baseQueryAllBaseledgerTransactionRequest,
    } as QueryAllBaseledgerTransactionRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBaseledgerTransactionResponse: object = {};

export const QueryAllBaseledgerTransactionResponse = {
  encode(
    message: QueryAllBaseledgerTransactionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.BaseledgerTransaction) {
      BaseledgerTransaction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllBaseledgerTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBaseledgerTransactionResponse,
    } as QueryAllBaseledgerTransactionResponse;
    message.BaseledgerTransaction = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.BaseledgerTransaction.push(
            BaseledgerTransaction.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllBaseledgerTransactionResponse {
    const message = {
      ...baseQueryAllBaseledgerTransactionResponse,
    } as QueryAllBaseledgerTransactionResponse;
    message.BaseledgerTransaction = [];
    if (
      object.BaseledgerTransaction !== undefined &&
      object.BaseledgerTransaction !== null
    ) {
      for (const e of object.BaseledgerTransaction) {
        message.BaseledgerTransaction.push(BaseledgerTransaction.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBaseledgerTransactionResponse): unknown {
    const obj: any = {};
    if (message.BaseledgerTransaction) {
      obj.BaseledgerTransaction = message.BaseledgerTransaction.map((e) =>
        e ? BaseledgerTransaction.toJSON(e) : undefined
      );
    } else {
      obj.BaseledgerTransaction = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBaseledgerTransactionResponse>
  ): QueryAllBaseledgerTransactionResponse {
    const message = {
      ...baseQueryAllBaseledgerTransactionResponse,
    } as QueryAllBaseledgerTransactionResponse;
    message.BaseledgerTransaction = [];
    if (
      object.BaseledgerTransaction !== undefined &&
      object.BaseledgerTransaction !== null
    ) {
      for (const e of object.BaseledgerTransaction) {
        message.BaseledgerTransaction.push(
          BaseledgerTransaction.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** this line is used by starport scaffolding # 2 */
  BaseledgerTransaction(
    request: QueryGetBaseledgerTransactionRequest
  ): Promise<QueryGetBaseledgerTransactionResponse>;
  BaseledgerTransactionAll(
    request: QueryAllBaseledgerTransactionRequest
  ): Promise<QueryAllBaseledgerTransactionResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  BaseledgerTransaction(
    request: QueryGetBaseledgerTransactionRequest
  ): Promise<QueryGetBaseledgerTransactionResponse> {
    const data = QueryGetBaseledgerTransactionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.baseledger.Query",
      "BaseledgerTransaction",
      data
    );
    return promise.then((data) =>
      QueryGetBaseledgerTransactionResponse.decode(new Reader(data))
    );
  }

  BaseledgerTransactionAll(
    request: QueryAllBaseledgerTransactionRequest
  ): Promise<QueryAllBaseledgerTransactionResponse> {
    const data = QueryAllBaseledgerTransactionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.baseledger.Query",
      "BaseledgerTransactionAll",
      data
    );
    return promise.then((data) =>
      QueryAllBaseledgerTransactionResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

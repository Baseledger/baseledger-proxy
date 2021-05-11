/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { SynchronizationFeedback } from "../trustmesh/SynchronizationFeedback";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { SynchronizationRequest } from "../trustmesh/SynchronizationRequest";

export const protobufPackage = "example.baseledger.trustmesh";

/** this line is used by starport scaffolding # 3 */
export interface QueryGetSynchronizationFeedbackRequest {
  id: number;
}

export interface QueryGetSynchronizationFeedbackResponse {
  SynchronizationFeedback: SynchronizationFeedback | undefined;
}

export interface QueryAllSynchronizationFeedbackRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSynchronizationFeedbackResponse {
  SynchronizationFeedback: SynchronizationFeedback[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSynchronizationRequestRequest {
  id: number;
}

export interface QueryGetSynchronizationRequestResponse {
  SynchronizationRequest: SynchronizationRequest | undefined;
}

export interface QueryAllSynchronizationRequestRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSynchronizationRequestResponse {
  SynchronizationRequest: SynchronizationRequest[];
  pagination: PageResponse | undefined;
}

const baseQueryGetSynchronizationFeedbackRequest: object = { id: 0 };

export const QueryGetSynchronizationFeedbackRequest = {
  encode(
    message: QueryGetSynchronizationFeedbackRequest,
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
  ): QueryGetSynchronizationFeedbackRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSynchronizationFeedbackRequest,
    } as QueryGetSynchronizationFeedbackRequest;
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

  fromJSON(object: any): QueryGetSynchronizationFeedbackRequest {
    const message = {
      ...baseQueryGetSynchronizationFeedbackRequest,
    } as QueryGetSynchronizationFeedbackRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetSynchronizationFeedbackRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSynchronizationFeedbackRequest>
  ): QueryGetSynchronizationFeedbackRequest {
    const message = {
      ...baseQueryGetSynchronizationFeedbackRequest,
    } as QueryGetSynchronizationFeedbackRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetSynchronizationFeedbackResponse: object = {};

export const QueryGetSynchronizationFeedbackResponse = {
  encode(
    message: QueryGetSynchronizationFeedbackResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SynchronizationFeedback !== undefined) {
      SynchronizationFeedback.encode(
        message.SynchronizationFeedback,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSynchronizationFeedbackResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSynchronizationFeedbackResponse,
    } as QueryGetSynchronizationFeedbackResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SynchronizationFeedback = SynchronizationFeedback.decode(
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

  fromJSON(object: any): QueryGetSynchronizationFeedbackResponse {
    const message = {
      ...baseQueryGetSynchronizationFeedbackResponse,
    } as QueryGetSynchronizationFeedbackResponse;
    if (
      object.SynchronizationFeedback !== undefined &&
      object.SynchronizationFeedback !== null
    ) {
      message.SynchronizationFeedback = SynchronizationFeedback.fromJSON(
        object.SynchronizationFeedback
      );
    } else {
      message.SynchronizationFeedback = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSynchronizationFeedbackResponse): unknown {
    const obj: any = {};
    message.SynchronizationFeedback !== undefined &&
      (obj.SynchronizationFeedback = message.SynchronizationFeedback
        ? SynchronizationFeedback.toJSON(message.SynchronizationFeedback)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSynchronizationFeedbackResponse>
  ): QueryGetSynchronizationFeedbackResponse {
    const message = {
      ...baseQueryGetSynchronizationFeedbackResponse,
    } as QueryGetSynchronizationFeedbackResponse;
    if (
      object.SynchronizationFeedback !== undefined &&
      object.SynchronizationFeedback !== null
    ) {
      message.SynchronizationFeedback = SynchronizationFeedback.fromPartial(
        object.SynchronizationFeedback
      );
    } else {
      message.SynchronizationFeedback = undefined;
    }
    return message;
  },
};

const baseQueryAllSynchronizationFeedbackRequest: object = {};

export const QueryAllSynchronizationFeedbackRequest = {
  encode(
    message: QueryAllSynchronizationFeedbackRequest,
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
  ): QueryAllSynchronizationFeedbackRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSynchronizationFeedbackRequest,
    } as QueryAllSynchronizationFeedbackRequest;
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

  fromJSON(object: any): QueryAllSynchronizationFeedbackRequest {
    const message = {
      ...baseQueryAllSynchronizationFeedbackRequest,
    } as QueryAllSynchronizationFeedbackRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSynchronizationFeedbackRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSynchronizationFeedbackRequest>
  ): QueryAllSynchronizationFeedbackRequest {
    const message = {
      ...baseQueryAllSynchronizationFeedbackRequest,
    } as QueryAllSynchronizationFeedbackRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSynchronizationFeedbackResponse: object = {};

export const QueryAllSynchronizationFeedbackResponse = {
  encode(
    message: QueryAllSynchronizationFeedbackResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.SynchronizationFeedback) {
      SynchronizationFeedback.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllSynchronizationFeedbackResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSynchronizationFeedbackResponse,
    } as QueryAllSynchronizationFeedbackResponse;
    message.SynchronizationFeedback = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SynchronizationFeedback.push(
            SynchronizationFeedback.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllSynchronizationFeedbackResponse {
    const message = {
      ...baseQueryAllSynchronizationFeedbackResponse,
    } as QueryAllSynchronizationFeedbackResponse;
    message.SynchronizationFeedback = [];
    if (
      object.SynchronizationFeedback !== undefined &&
      object.SynchronizationFeedback !== null
    ) {
      for (const e of object.SynchronizationFeedback) {
        message.SynchronizationFeedback.push(
          SynchronizationFeedback.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSynchronizationFeedbackResponse): unknown {
    const obj: any = {};
    if (message.SynchronizationFeedback) {
      obj.SynchronizationFeedback = message.SynchronizationFeedback.map((e) =>
        e ? SynchronizationFeedback.toJSON(e) : undefined
      );
    } else {
      obj.SynchronizationFeedback = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSynchronizationFeedbackResponse>
  ): QueryAllSynchronizationFeedbackResponse {
    const message = {
      ...baseQueryAllSynchronizationFeedbackResponse,
    } as QueryAllSynchronizationFeedbackResponse;
    message.SynchronizationFeedback = [];
    if (
      object.SynchronizationFeedback !== undefined &&
      object.SynchronizationFeedback !== null
    ) {
      for (const e of object.SynchronizationFeedback) {
        message.SynchronizationFeedback.push(
          SynchronizationFeedback.fromPartial(e)
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

const baseQueryGetSynchronizationRequestRequest: object = { id: 0 };

export const QueryGetSynchronizationRequestRequest = {
  encode(
    message: QueryGetSynchronizationRequestRequest,
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
  ): QueryGetSynchronizationRequestRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSynchronizationRequestRequest,
    } as QueryGetSynchronizationRequestRequest;
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

  fromJSON(object: any): QueryGetSynchronizationRequestRequest {
    const message = {
      ...baseQueryGetSynchronizationRequestRequest,
    } as QueryGetSynchronizationRequestRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetSynchronizationRequestRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSynchronizationRequestRequest>
  ): QueryGetSynchronizationRequestRequest {
    const message = {
      ...baseQueryGetSynchronizationRequestRequest,
    } as QueryGetSynchronizationRequestRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetSynchronizationRequestResponse: object = {};

export const QueryGetSynchronizationRequestResponse = {
  encode(
    message: QueryGetSynchronizationRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SynchronizationRequest !== undefined) {
      SynchronizationRequest.encode(
        message.SynchronizationRequest,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSynchronizationRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSynchronizationRequestResponse,
    } as QueryGetSynchronizationRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SynchronizationRequest = SynchronizationRequest.decode(
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

  fromJSON(object: any): QueryGetSynchronizationRequestResponse {
    const message = {
      ...baseQueryGetSynchronizationRequestResponse,
    } as QueryGetSynchronizationRequestResponse;
    if (
      object.SynchronizationRequest !== undefined &&
      object.SynchronizationRequest !== null
    ) {
      message.SynchronizationRequest = SynchronizationRequest.fromJSON(
        object.SynchronizationRequest
      );
    } else {
      message.SynchronizationRequest = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSynchronizationRequestResponse): unknown {
    const obj: any = {};
    message.SynchronizationRequest !== undefined &&
      (obj.SynchronizationRequest = message.SynchronizationRequest
        ? SynchronizationRequest.toJSON(message.SynchronizationRequest)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSynchronizationRequestResponse>
  ): QueryGetSynchronizationRequestResponse {
    const message = {
      ...baseQueryGetSynchronizationRequestResponse,
    } as QueryGetSynchronizationRequestResponse;
    if (
      object.SynchronizationRequest !== undefined &&
      object.SynchronizationRequest !== null
    ) {
      message.SynchronizationRequest = SynchronizationRequest.fromPartial(
        object.SynchronizationRequest
      );
    } else {
      message.SynchronizationRequest = undefined;
    }
    return message;
  },
};

const baseQueryAllSynchronizationRequestRequest: object = {};

export const QueryAllSynchronizationRequestRequest = {
  encode(
    message: QueryAllSynchronizationRequestRequest,
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
  ): QueryAllSynchronizationRequestRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSynchronizationRequestRequest,
    } as QueryAllSynchronizationRequestRequest;
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

  fromJSON(object: any): QueryAllSynchronizationRequestRequest {
    const message = {
      ...baseQueryAllSynchronizationRequestRequest,
    } as QueryAllSynchronizationRequestRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSynchronizationRequestRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSynchronizationRequestRequest>
  ): QueryAllSynchronizationRequestRequest {
    const message = {
      ...baseQueryAllSynchronizationRequestRequest,
    } as QueryAllSynchronizationRequestRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSynchronizationRequestResponse: object = {};

export const QueryAllSynchronizationRequestResponse = {
  encode(
    message: QueryAllSynchronizationRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.SynchronizationRequest) {
      SynchronizationRequest.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllSynchronizationRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSynchronizationRequestResponse,
    } as QueryAllSynchronizationRequestResponse;
    message.SynchronizationRequest = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SynchronizationRequest.push(
            SynchronizationRequest.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllSynchronizationRequestResponse {
    const message = {
      ...baseQueryAllSynchronizationRequestResponse,
    } as QueryAllSynchronizationRequestResponse;
    message.SynchronizationRequest = [];
    if (
      object.SynchronizationRequest !== undefined &&
      object.SynchronizationRequest !== null
    ) {
      for (const e of object.SynchronizationRequest) {
        message.SynchronizationRequest.push(SynchronizationRequest.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSynchronizationRequestResponse): unknown {
    const obj: any = {};
    if (message.SynchronizationRequest) {
      obj.SynchronizationRequest = message.SynchronizationRequest.map((e) =>
        e ? SynchronizationRequest.toJSON(e) : undefined
      );
    } else {
      obj.SynchronizationRequest = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSynchronizationRequestResponse>
  ): QueryAllSynchronizationRequestResponse {
    const message = {
      ...baseQueryAllSynchronizationRequestResponse,
    } as QueryAllSynchronizationRequestResponse;
    message.SynchronizationRequest = [];
    if (
      object.SynchronizationRequest !== undefined &&
      object.SynchronizationRequest !== null
    ) {
      for (const e of object.SynchronizationRequest) {
        message.SynchronizationRequest.push(
          SynchronizationRequest.fromPartial(e)
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
  SynchronizationFeedback(
    request: QueryGetSynchronizationFeedbackRequest
  ): Promise<QueryGetSynchronizationFeedbackResponse>;
  SynchronizationFeedbackAll(
    request: QueryAllSynchronizationFeedbackRequest
  ): Promise<QueryAllSynchronizationFeedbackResponse>;
  SynchronizationRequest(
    request: QueryGetSynchronizationRequestRequest
  ): Promise<QueryGetSynchronizationRequestResponse>;
  SynchronizationRequestAll(
    request: QueryAllSynchronizationRequestRequest
  ): Promise<QueryAllSynchronizationRequestResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SynchronizationFeedback(
    request: QueryGetSynchronizationFeedbackRequest
  ): Promise<QueryGetSynchronizationFeedbackResponse> {
    const data = QueryGetSynchronizationFeedbackRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Query",
      "SynchronizationFeedback",
      data
    );
    return promise.then((data) =>
      QueryGetSynchronizationFeedbackResponse.decode(new Reader(data))
    );
  }

  SynchronizationFeedbackAll(
    request: QueryAllSynchronizationFeedbackRequest
  ): Promise<QueryAllSynchronizationFeedbackResponse> {
    const data = QueryAllSynchronizationFeedbackRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Query",
      "SynchronizationFeedbackAll",
      data
    );
    return promise.then((data) =>
      QueryAllSynchronizationFeedbackResponse.decode(new Reader(data))
    );
  }

  SynchronizationRequest(
    request: QueryGetSynchronizationRequestRequest
  ): Promise<QueryGetSynchronizationRequestResponse> {
    const data = QueryGetSynchronizationRequestRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Query",
      "SynchronizationRequest",
      data
    );
    return promise.then((data) =>
      QueryGetSynchronizationRequestResponse.decode(new Reader(data))
    );
  }

  SynchronizationRequestAll(
    request: QueryAllSynchronizationRequestRequest
  ): Promise<QueryAllSynchronizationRequestResponse> {
    const data = QueryAllSynchronizationRequestRequest.encode(request).finish();
    const promise = this.rpc.request(
      "example.baseledger.trustmesh.Query",
      "SynchronizationRequestAll",
      data
    );
    return promise.then((data) =>
      QueryAllSynchronizationRequestResponse.decode(new Reader(data))
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

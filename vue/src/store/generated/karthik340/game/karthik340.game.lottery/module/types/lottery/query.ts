/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../lottery/params";
import { Round } from "../lottery/round";
import { TxnCounter } from "../lottery/txn_counter";
import { Bet } from "../lottery/bet";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { ValidatorsWinner } from "../lottery/validators_winner";

export const protobufPackage = "karthik340.game.lottery";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetRoundRequest {}

export interface QueryGetRoundResponse {
  Round: Round | undefined;
}

export interface QueryGetTxnCounterRequest {}

export interface QueryGetTxnCounterResponse {
  TxnCounter: TxnCounter | undefined;
}

export interface QueryGetBetRequest {
  sender: string;
}

export interface QueryGetBetResponse {
  bet: Bet | undefined;
}

export interface QueryAllBetRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBetResponse {
  bet: Bet[];
  pagination: PageResponse | undefined;
}

export interface QueryGetWinnerByRoundRequest {
  round: number;
}

export interface QueryGetWinnerByRoundResponse {
  winner: string;
}

export interface QueryGetValidatorsWinnerRequest {}

export interface QueryGetValidatorsWinnerResponse {
  ValidatorsWinner: ValidatorsWinner | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetRoundRequest: object = {};

export const QueryGetRoundRequest = {
  encode(_: QueryGetRoundRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRoundRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetRoundRequest } as QueryGetRoundRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetRoundRequest {
    const message = { ...baseQueryGetRoundRequest } as QueryGetRoundRequest;
    return message;
  },

  toJSON(_: QueryGetRoundRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetRoundRequest>): QueryGetRoundRequest {
    const message = { ...baseQueryGetRoundRequest } as QueryGetRoundRequest;
    return message;
  },
};

const baseQueryGetRoundResponse: object = {};

export const QueryGetRoundResponse = {
  encode(
    message: QueryGetRoundResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Round !== undefined) {
      Round.encode(message.Round, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRoundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetRoundResponse } as QueryGetRoundResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Round = Round.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRoundResponse {
    const message = { ...baseQueryGetRoundResponse } as QueryGetRoundResponse;
    if (object.Round !== undefined && object.Round !== null) {
      message.Round = Round.fromJSON(object.Round);
    } else {
      message.Round = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRoundResponse): unknown {
    const obj: any = {};
    message.Round !== undefined &&
      (obj.Round = message.Round ? Round.toJSON(message.Round) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRoundResponse>
  ): QueryGetRoundResponse {
    const message = { ...baseQueryGetRoundResponse } as QueryGetRoundResponse;
    if (object.Round !== undefined && object.Round !== null) {
      message.Round = Round.fromPartial(object.Round);
    } else {
      message.Round = undefined;
    }
    return message;
  },
};

const baseQueryGetTxnCounterRequest: object = {};

export const QueryGetTxnCounterRequest = {
  encode(
    _: QueryGetTxnCounterRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTxnCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTxnCounterRequest,
    } as QueryGetTxnCounterRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetTxnCounterRequest {
    const message = {
      ...baseQueryGetTxnCounterRequest,
    } as QueryGetTxnCounterRequest;
    return message;
  },

  toJSON(_: QueryGetTxnCounterRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTxnCounterRequest>
  ): QueryGetTxnCounterRequest {
    const message = {
      ...baseQueryGetTxnCounterRequest,
    } as QueryGetTxnCounterRequest;
    return message;
  },
};

const baseQueryGetTxnCounterResponse: object = {};

export const QueryGetTxnCounterResponse = {
  encode(
    message: QueryGetTxnCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TxnCounter !== undefined) {
      TxnCounter.encode(message.TxnCounter, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTxnCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTxnCounterResponse,
    } as QueryGetTxnCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TxnCounter = TxnCounter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTxnCounterResponse {
    const message = {
      ...baseQueryGetTxnCounterResponse,
    } as QueryGetTxnCounterResponse;
    if (object.TxnCounter !== undefined && object.TxnCounter !== null) {
      message.TxnCounter = TxnCounter.fromJSON(object.TxnCounter);
    } else {
      message.TxnCounter = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTxnCounterResponse): unknown {
    const obj: any = {};
    message.TxnCounter !== undefined &&
      (obj.TxnCounter = message.TxnCounter
        ? TxnCounter.toJSON(message.TxnCounter)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTxnCounterResponse>
  ): QueryGetTxnCounterResponse {
    const message = {
      ...baseQueryGetTxnCounterResponse,
    } as QueryGetTxnCounterResponse;
    if (object.TxnCounter !== undefined && object.TxnCounter !== null) {
      message.TxnCounter = TxnCounter.fromPartial(object.TxnCounter);
    } else {
      message.TxnCounter = undefined;
    }
    return message;
  },
};

const baseQueryGetBetRequest: object = { sender: "" };

export const QueryGetBetRequest = {
  encode(
    message: QueryGetBetRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBetRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBetRequest } as QueryGetBetRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBetRequest {
    const message = { ...baseQueryGetBetRequest } as QueryGetBetRequest;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: QueryGetBetRequest): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetBetRequest>): QueryGetBetRequest {
    const message = { ...baseQueryGetBetRequest } as QueryGetBetRequest;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseQueryGetBetResponse: object = {};

export const QueryGetBetResponse = {
  encode(
    message: QueryGetBetResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bet !== undefined) {
      Bet.encode(message.bet, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBetResponse } as QueryGetBetResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bet = Bet.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBetResponse {
    const message = { ...baseQueryGetBetResponse } as QueryGetBetResponse;
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = Bet.fromJSON(object.bet);
    } else {
      message.bet = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBetResponse): unknown {
    const obj: any = {};
    message.bet !== undefined &&
      (obj.bet = message.bet ? Bet.toJSON(message.bet) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetBetResponse>): QueryGetBetResponse {
    const message = { ...baseQueryGetBetResponse } as QueryGetBetResponse;
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = Bet.fromPartial(object.bet);
    } else {
      message.bet = undefined;
    }
    return message;
  },
};

const baseQueryAllBetRequest: object = {};

export const QueryAllBetRequest = {
  encode(
    message: QueryAllBetRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBetRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBetRequest } as QueryAllBetRequest;
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

  fromJSON(object: any): QueryAllBetRequest {
    const message = { ...baseQueryAllBetRequest } as QueryAllBetRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBetRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllBetRequest>): QueryAllBetRequest {
    const message = { ...baseQueryAllBetRequest } as QueryAllBetRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBetResponse: object = {};

export const QueryAllBetResponse = {
  encode(
    message: QueryAllBetResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.bet) {
      Bet.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBetResponse } as QueryAllBetResponse;
    message.bet = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bet.push(Bet.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllBetResponse {
    const message = { ...baseQueryAllBetResponse } as QueryAllBetResponse;
    message.bet = [];
    if (object.bet !== undefined && object.bet !== null) {
      for (const e of object.bet) {
        message.bet.push(Bet.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBetResponse): unknown {
    const obj: any = {};
    if (message.bet) {
      obj.bet = message.bet.map((e) => (e ? Bet.toJSON(e) : undefined));
    } else {
      obj.bet = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllBetResponse>): QueryAllBetResponse {
    const message = { ...baseQueryAllBetResponse } as QueryAllBetResponse;
    message.bet = [];
    if (object.bet !== undefined && object.bet !== null) {
      for (const e of object.bet) {
        message.bet.push(Bet.fromPartial(e));
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

const baseQueryGetWinnerByRoundRequest: object = { round: 0 };

export const QueryGetWinnerByRoundRequest = {
  encode(
    message: QueryGetWinnerByRoundRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.round !== 0) {
      writer.uint32(8).uint64(message.round);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetWinnerByRoundRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetWinnerByRoundRequest,
    } as QueryGetWinnerByRoundRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.round = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWinnerByRoundRequest {
    const message = {
      ...baseQueryGetWinnerByRoundRequest,
    } as QueryGetWinnerByRoundRequest;
    if (object.round !== undefined && object.round !== null) {
      message.round = Number(object.round);
    } else {
      message.round = 0;
    }
    return message;
  },

  toJSON(message: QueryGetWinnerByRoundRequest): unknown {
    const obj: any = {};
    message.round !== undefined && (obj.round = message.round);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetWinnerByRoundRequest>
  ): QueryGetWinnerByRoundRequest {
    const message = {
      ...baseQueryGetWinnerByRoundRequest,
    } as QueryGetWinnerByRoundRequest;
    if (object.round !== undefined && object.round !== null) {
      message.round = object.round;
    } else {
      message.round = 0;
    }
    return message;
  },
};

const baseQueryGetWinnerByRoundResponse: object = { winner: "" };

export const QueryGetWinnerByRoundResponse = {
  encode(
    message: QueryGetWinnerByRoundResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.winner !== "") {
      writer.uint32(10).string(message.winner);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetWinnerByRoundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetWinnerByRoundResponse,
    } as QueryGetWinnerByRoundResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWinnerByRoundResponse {
    const message = {
      ...baseQueryGetWinnerByRoundResponse,
    } as QueryGetWinnerByRoundResponse;
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    return message;
  },

  toJSON(message: QueryGetWinnerByRoundResponse): unknown {
    const obj: any = {};
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetWinnerByRoundResponse>
  ): QueryGetWinnerByRoundResponse {
    const message = {
      ...baseQueryGetWinnerByRoundResponse,
    } as QueryGetWinnerByRoundResponse;
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    return message;
  },
};

const baseQueryGetValidatorsWinnerRequest: object = {};

export const QueryGetValidatorsWinnerRequest = {
  encode(
    _: QueryGetValidatorsWinnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetValidatorsWinnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetValidatorsWinnerRequest,
    } as QueryGetValidatorsWinnerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetValidatorsWinnerRequest {
    const message = {
      ...baseQueryGetValidatorsWinnerRequest,
    } as QueryGetValidatorsWinnerRequest;
    return message;
  },

  toJSON(_: QueryGetValidatorsWinnerRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetValidatorsWinnerRequest>
  ): QueryGetValidatorsWinnerRequest {
    const message = {
      ...baseQueryGetValidatorsWinnerRequest,
    } as QueryGetValidatorsWinnerRequest;
    return message;
  },
};

const baseQueryGetValidatorsWinnerResponse: object = {};

export const QueryGetValidatorsWinnerResponse = {
  encode(
    message: QueryGetValidatorsWinnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ValidatorsWinner !== undefined) {
      ValidatorsWinner.encode(
        message.ValidatorsWinner,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetValidatorsWinnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetValidatorsWinnerResponse,
    } as QueryGetValidatorsWinnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ValidatorsWinner = ValidatorsWinner.decode(
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

  fromJSON(object: any): QueryGetValidatorsWinnerResponse {
    const message = {
      ...baseQueryGetValidatorsWinnerResponse,
    } as QueryGetValidatorsWinnerResponse;
    if (
      object.ValidatorsWinner !== undefined &&
      object.ValidatorsWinner !== null
    ) {
      message.ValidatorsWinner = ValidatorsWinner.fromJSON(
        object.ValidatorsWinner
      );
    } else {
      message.ValidatorsWinner = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetValidatorsWinnerResponse): unknown {
    const obj: any = {};
    message.ValidatorsWinner !== undefined &&
      (obj.ValidatorsWinner = message.ValidatorsWinner
        ? ValidatorsWinner.toJSON(message.ValidatorsWinner)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetValidatorsWinnerResponse>
  ): QueryGetValidatorsWinnerResponse {
    const message = {
      ...baseQueryGetValidatorsWinnerResponse,
    } as QueryGetValidatorsWinnerResponse;
    if (
      object.ValidatorsWinner !== undefined &&
      object.ValidatorsWinner !== null
    ) {
      message.ValidatorsWinner = ValidatorsWinner.fromPartial(
        object.ValidatorsWinner
      );
    } else {
      message.ValidatorsWinner = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Round by index. */
  Round(request: QueryGetRoundRequest): Promise<QueryGetRoundResponse>;
  /** Queries a TxnCounter by index. */
  TxnCounter(
    request: QueryGetTxnCounterRequest
  ): Promise<QueryGetTxnCounterResponse>;
  /** Queries a Bet by index. */
  Bet(request: QueryGetBetRequest): Promise<QueryGetBetResponse>;
  /** Queries a list of Bet items. */
  BetAll(request: QueryAllBetRequest): Promise<QueryAllBetResponse>;
  /** Queries a list of GetWinnerByRound items. */
  GetWinnerByRound(
    request: QueryGetWinnerByRoundRequest
  ): Promise<QueryGetWinnerByRoundResponse>;
  /** Queries a ValidatorsWinner by index. */
  ValidatorsWinner(
    request: QueryGetValidatorsWinnerRequest
  ): Promise<QueryGetValidatorsWinnerResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Round(request: QueryGetRoundRequest): Promise<QueryGetRoundResponse> {
    const data = QueryGetRoundRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "Round",
      data
    );
    return promise.then((data) =>
      QueryGetRoundResponse.decode(new Reader(data))
    );
  }

  TxnCounter(
    request: QueryGetTxnCounterRequest
  ): Promise<QueryGetTxnCounterResponse> {
    const data = QueryGetTxnCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "TxnCounter",
      data
    );
    return promise.then((data) =>
      QueryGetTxnCounterResponse.decode(new Reader(data))
    );
  }

  Bet(request: QueryGetBetRequest): Promise<QueryGetBetResponse> {
    const data = QueryGetBetRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "Bet",
      data
    );
    return promise.then((data) => QueryGetBetResponse.decode(new Reader(data)));
  }

  BetAll(request: QueryAllBetRequest): Promise<QueryAllBetResponse> {
    const data = QueryAllBetRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "BetAll",
      data
    );
    return promise.then((data) => QueryAllBetResponse.decode(new Reader(data)));
  }

  GetWinnerByRound(
    request: QueryGetWinnerByRoundRequest
  ): Promise<QueryGetWinnerByRoundResponse> {
    const data = QueryGetWinnerByRoundRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "GetWinnerByRound",
      data
    );
    return promise.then((data) =>
      QueryGetWinnerByRoundResponse.decode(new Reader(data))
    );
  }

  ValidatorsWinner(
    request: QueryGetValidatorsWinnerRequest
  ): Promise<QueryGetValidatorsWinnerResponse> {
    const data = QueryGetValidatorsWinnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Query",
      "ValidatorsWinner",
      data
    );
    return promise.then((data) =>
      QueryGetValidatorsWinnerResponse.decode(new Reader(data))
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

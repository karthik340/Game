/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "karthik340.game.lottery";

export interface MsgPlaceBet {
  sender: string;
  fee: Coin | undefined;
  bet: Coin | undefined;
}

export interface MsgPlaceBetResponse {}

export interface MsgSetWinner {
  validator: string;
  winner: string;
}

export interface MsgSetWinnerResponse {}

const baseMsgPlaceBet: object = { sender: "" };

export const MsgPlaceBet = {
  encode(message: MsgPlaceBet, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.fee !== undefined) {
      Coin.encode(message.fee, writer.uint32(18).fork()).ldelim();
    }
    if (message.bet !== undefined) {
      Coin.encode(message.bet, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlaceBet {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlaceBet } as MsgPlaceBet;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.fee = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.bet = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlaceBet {
    const message = { ...baseMsgPlaceBet } as MsgPlaceBet;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Coin.fromJSON(object.fee);
    } else {
      message.fee = undefined;
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = Coin.fromJSON(object.bet);
    } else {
      message.bet = undefined;
    }
    return message;
  },

  toJSON(message: MsgPlaceBet): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.fee !== undefined &&
      (obj.fee = message.fee ? Coin.toJSON(message.fee) : undefined);
    message.bet !== undefined &&
      (obj.bet = message.bet ? Coin.toJSON(message.bet) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlaceBet>): MsgPlaceBet {
    const message = { ...baseMsgPlaceBet } as MsgPlaceBet;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Coin.fromPartial(object.fee);
    } else {
      message.fee = undefined;
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = Coin.fromPartial(object.bet);
    } else {
      message.bet = undefined;
    }
    return message;
  },
};

const baseMsgPlaceBetResponse: object = {};

export const MsgPlaceBetResponse = {
  encode(_: MsgPlaceBetResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlaceBetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlaceBetResponse } as MsgPlaceBetResponse;
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

  fromJSON(_: any): MsgPlaceBetResponse {
    const message = { ...baseMsgPlaceBetResponse } as MsgPlaceBetResponse;
    return message;
  },

  toJSON(_: MsgPlaceBetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgPlaceBetResponse>): MsgPlaceBetResponse {
    const message = { ...baseMsgPlaceBetResponse } as MsgPlaceBetResponse;
    return message;
  },
};

const baseMsgSetWinner: object = { validator: "", winner: "" };

export const MsgSetWinner = {
  encode(message: MsgSetWinner, writer: Writer = Writer.create()): Writer {
    if (message.validator !== "") {
      writer.uint32(10).string(message.validator);
    }
    if (message.winner !== "") {
      writer.uint32(18).string(message.winner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetWinner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetWinner } as MsgSetWinner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator = reader.string();
          break;
        case 2:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetWinner {
    const message = { ...baseMsgSetWinner } as MsgSetWinner;
    if (object.validator !== undefined && object.validator !== null) {
      message.validator = String(object.validator);
    } else {
      message.validator = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    return message;
  },

  toJSON(message: MsgSetWinner): unknown {
    const obj: any = {};
    message.validator !== undefined && (obj.validator = message.validator);
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetWinner>): MsgSetWinner {
    const message = { ...baseMsgSetWinner } as MsgSetWinner;
    if (object.validator !== undefined && object.validator !== null) {
      message.validator = object.validator;
    } else {
      message.validator = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    return message;
  },
};

const baseMsgSetWinnerResponse: object = {};

export const MsgSetWinnerResponse = {
  encode(_: MsgSetWinnerResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetWinnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetWinnerResponse } as MsgSetWinnerResponse;
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

  fromJSON(_: any): MsgSetWinnerResponse {
    const message = { ...baseMsgSetWinnerResponse } as MsgSetWinnerResponse;
    return message;
  },

  toJSON(_: MsgSetWinnerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSetWinnerResponse>): MsgSetWinnerResponse {
    const message = { ...baseMsgSetWinnerResponse } as MsgSetWinnerResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  PlaceBet(request: MsgPlaceBet): Promise<MsgPlaceBetResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetWinner(request: MsgSetWinner): Promise<MsgSetWinnerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  PlaceBet(request: MsgPlaceBet): Promise<MsgPlaceBetResponse> {
    const data = MsgPlaceBet.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Msg",
      "PlaceBet",
      data
    );
    return promise.then((data) => MsgPlaceBetResponse.decode(new Reader(data)));
  }

  SetWinner(request: MsgSetWinner): Promise<MsgSetWinnerResponse> {
    const data = MsgSetWinner.encode(request).finish();
    const promise = this.rpc.request(
      "karthik340.game.lottery.Msg",
      "SetWinner",
      data
    );
    return promise.then((data) =>
      MsgSetWinnerResponse.decode(new Reader(data))
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

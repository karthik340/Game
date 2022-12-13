/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "game.lottery";

export interface MsgPlaceBet {
  sender: string;
  fee: Coin | undefined;
  bet: Coin | undefined;
}

export interface MsgPlaceBetResponse {}

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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PlaceBet(request: MsgPlaceBet): Promise<MsgPlaceBetResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  PlaceBet(request: MsgPlaceBet): Promise<MsgPlaceBetResponse> {
    const data = MsgPlaceBet.encode(request).finish();
    const promise = this.rpc.request("game.lottery.Msg", "PlaceBet", data);
    return promise.then((data) => MsgPlaceBetResponse.decode(new Reader(data)));
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

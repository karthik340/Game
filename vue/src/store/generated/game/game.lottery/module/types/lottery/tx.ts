/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "game.lottery";

export interface MsgPlaceBet {
  creator: string;
  fee: string;
  bet: string;
}

export interface MsgPlaceBetResponse {}

const baseMsgPlaceBet: object = { creator: "", fee: "", bet: "" };

export const MsgPlaceBet = {
  encode(message: MsgPlaceBet, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    if (message.bet !== "") {
      writer.uint32(26).string(message.bet);
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
          message.creator = reader.string();
          break;
        case 2:
          message.fee = reader.string();
          break;
        case 3:
          message.bet = reader.string();
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = String(object.bet);
    } else {
      message.bet = "";
    }
    return message;
  },

  toJSON(message: MsgPlaceBet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fee !== undefined && (obj.fee = message.fee);
    message.bet !== undefined && (obj.bet = message.bet);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlaceBet>): MsgPlaceBet {
    const message = { ...baseMsgPlaceBet } as MsgPlaceBet;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = object.bet;
    } else {
      message.bet = "";
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

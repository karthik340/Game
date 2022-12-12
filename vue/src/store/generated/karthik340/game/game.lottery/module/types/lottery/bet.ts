/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "game.lottery";

export interface Bet {
  sender: string;
  fee: Coin | undefined;
  bet: Coin | undefined;
  status: boolean;
  txNum: number;
}

const baseBet: object = { sender: "", status: false, txNum: 0 };

export const Bet = {
  encode(message: Bet, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.fee !== undefined) {
      Coin.encode(message.fee, writer.uint32(18).fork()).ldelim();
    }
    if (message.bet !== undefined) {
      Coin.encode(message.bet, writer.uint32(26).fork()).ldelim();
    }
    if (message.status === true) {
      writer.uint32(32).bool(message.status);
    }
    if (message.txNum !== 0) {
      writer.uint32(40).uint64(message.txNum);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Bet {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBet } as Bet;
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
        case 4:
          message.status = reader.bool();
          break;
        case 5:
          message.txNum = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Bet {
    const message = { ...baseBet } as Bet;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = Boolean(object.status);
    } else {
      message.status = false;
    }
    if (object.txNum !== undefined && object.txNum !== null) {
      message.txNum = Number(object.txNum);
    } else {
      message.txNum = 0;
    }
    return message;
  },

  toJSON(message: Bet): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.fee !== undefined &&
      (obj.fee = message.fee ? Coin.toJSON(message.fee) : undefined);
    message.bet !== undefined &&
      (obj.bet = message.bet ? Coin.toJSON(message.bet) : undefined);
    message.status !== undefined && (obj.status = message.status);
    message.txNum !== undefined && (obj.txNum = message.txNum);
    return obj;
  },

  fromPartial(object: DeepPartial<Bet>): Bet {
    const message = { ...baseBet } as Bet;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = false;
    }
    if (object.txNum !== undefined && object.txNum !== null) {
      message.txNum = object.txNum;
    } else {
      message.txNum = 0;
    }
    return message;
  },
};

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

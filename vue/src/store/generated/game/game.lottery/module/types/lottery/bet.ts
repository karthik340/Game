/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "game.lottery";

export interface Bet {
  sender: string;
  betSize: number;
  fee: number;
  status: boolean;
}

const baseBet: object = { sender: "", betSize: 0, fee: 0, status: false };

export const Bet = {
  encode(message: Bet, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.betSize !== 0) {
      writer.uint32(16).uint64(message.betSize);
    }
    if (message.fee !== 0) {
      writer.uint32(24).uint64(message.fee);
    }
    if (message.status === true) {
      writer.uint32(32).bool(message.status);
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
          message.betSize = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.fee = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.status = reader.bool();
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
    if (object.betSize !== undefined && object.betSize !== null) {
      message.betSize = Number(object.betSize);
    } else {
      message.betSize = 0;
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Number(object.fee);
    } else {
      message.fee = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = Boolean(object.status);
    } else {
      message.status = false;
    }
    return message;
  },

  toJSON(message: Bet): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.betSize !== undefined && (obj.betSize = message.betSize);
    message.fee !== undefined && (obj.fee = message.fee);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Bet>): Bet {
    const message = { ...baseBet } as Bet;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.betSize !== undefined && object.betSize !== null) {
      message.betSize = object.betSize;
    } else {
      message.betSize = 0;
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = false;
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

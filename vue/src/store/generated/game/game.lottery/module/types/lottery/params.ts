/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "game.lottery";

/** Params defines the parameters for the module. */
export interface Params {
  minFee: number;
  minBet: number;
  maxBet: number;
  minTxn: number;
}

const baseParams: object = { minFee: 0, minBet: 0, maxBet: 0, minTxn: 0 };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.minFee !== 0) {
      writer.uint32(8).uint64(message.minFee);
    }
    if (message.minBet !== 0) {
      writer.uint32(16).uint64(message.minBet);
    }
    if (message.maxBet !== 0) {
      writer.uint32(24).uint64(message.maxBet);
    }
    if (message.minTxn !== 0) {
      writer.uint32(32).uint64(message.minTxn);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minFee = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.minBet = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.maxBet = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.minTxn = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.minFee !== undefined && object.minFee !== null) {
      message.minFee = Number(object.minFee);
    } else {
      message.minFee = 0;
    }
    if (object.minBet !== undefined && object.minBet !== null) {
      message.minBet = Number(object.minBet);
    } else {
      message.minBet = 0;
    }
    if (object.maxBet !== undefined && object.maxBet !== null) {
      message.maxBet = Number(object.maxBet);
    } else {
      message.maxBet = 0;
    }
    if (object.minTxn !== undefined && object.minTxn !== null) {
      message.minTxn = Number(object.minTxn);
    } else {
      message.minTxn = 0;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.minFee !== undefined && (obj.minFee = message.minFee);
    message.minBet !== undefined && (obj.minBet = message.minBet);
    message.maxBet !== undefined && (obj.maxBet = message.maxBet);
    message.minTxn !== undefined && (obj.minTxn = message.minTxn);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.minFee !== undefined && object.minFee !== null) {
      message.minFee = object.minFee;
    } else {
      message.minFee = 0;
    }
    if (object.minBet !== undefined && object.minBet !== null) {
      message.minBet = object.minBet;
    } else {
      message.minBet = 0;
    }
    if (object.maxBet !== undefined && object.maxBet !== null) {
      message.maxBet = object.maxBet;
    } else {
      message.maxBet = 0;
    }
    if (object.minTxn !== undefined && object.minTxn !== null) {
      message.minTxn = object.minTxn;
    } else {
      message.minTxn = 0;
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

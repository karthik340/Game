/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "karthik340.game.lottery";

export interface TxnCounter {
  val: number;
}

const baseTxnCounter: object = { val: 0 };

export const TxnCounter = {
  encode(message: TxnCounter, writer: Writer = Writer.create()): Writer {
    if (message.val !== 0) {
      writer.uint32(8).uint64(message.val);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TxnCounter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTxnCounter } as TxnCounter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.val = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TxnCounter {
    const message = { ...baseTxnCounter } as TxnCounter;
    if (object.val !== undefined && object.val !== null) {
      message.val = Number(object.val);
    } else {
      message.val = 0;
    }
    return message;
  },

  toJSON(message: TxnCounter): unknown {
    const obj: any = {};
    message.val !== undefined && (obj.val = message.val);
    return obj;
  },

  fromPartial(object: DeepPartial<TxnCounter>): TxnCounter {
    const message = { ...baseTxnCounter } as TxnCounter;
    if (object.val !== undefined && object.val !== null) {
      message.val = object.val;
    } else {
      message.val = 0;
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

/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "karthik340.game.lottery";

export interface ValidatorsWinner {
  winner: string;
}

const baseValidatorsWinner: object = { winner: "" };

export const ValidatorsWinner = {
  encode(message: ValidatorsWinner, writer: Writer = Writer.create()): Writer {
    if (message.winner !== "") {
      writer.uint32(10).string(message.winner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ValidatorsWinner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseValidatorsWinner } as ValidatorsWinner;
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

  fromJSON(object: any): ValidatorsWinner {
    const message = { ...baseValidatorsWinner } as ValidatorsWinner;
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    return message;
  },

  toJSON(message: ValidatorsWinner): unknown {
    const obj: any = {};
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial(object: DeepPartial<ValidatorsWinner>): ValidatorsWinner {
    const message = { ...baseValidatorsWinner } as ValidatorsWinner;
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    return message;
  },
};

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

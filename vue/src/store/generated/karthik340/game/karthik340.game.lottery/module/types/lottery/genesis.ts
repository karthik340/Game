/* eslint-disable */
import { Params } from "../lottery/params";
import { Round } from "../lottery/round";
import { TxnCounter } from "../lottery/txn_counter";
import { Bet } from "../lottery/bet";
import { ValidatorsWinner } from "../lottery/validators_winner";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "karthik340.game.lottery";

/** GenesisState defines the lottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  round: Round | undefined;
  txnCounter: TxnCounter | undefined;
  betList: Bet[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  validatorsWinner: ValidatorsWinner[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.round !== undefined) {
      Round.encode(message.round, writer.uint32(18).fork()).ldelim();
    }
    if (message.txnCounter !== undefined) {
      TxnCounter.encode(message.txnCounter, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.betList) {
      Bet.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.validatorsWinner) {
      ValidatorsWinner.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.betList = [];
    message.validatorsWinner = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.round = Round.decode(reader, reader.uint32());
          break;
        case 3:
          message.txnCounter = TxnCounter.decode(reader, reader.uint32());
          break;
        case 4:
          message.betList.push(Bet.decode(reader, reader.uint32()));
          break;
        case 5:
          message.validatorsWinner.push(
            ValidatorsWinner.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.betList = [];
    message.validatorsWinner = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.round !== undefined && object.round !== null) {
      message.round = Round.fromJSON(object.round);
    } else {
      message.round = undefined;
    }
    if (object.txnCounter !== undefined && object.txnCounter !== null) {
      message.txnCounter = TxnCounter.fromJSON(object.txnCounter);
    } else {
      message.txnCounter = undefined;
    }
    if (object.betList !== undefined && object.betList !== null) {
      for (const e of object.betList) {
        message.betList.push(Bet.fromJSON(e));
      }
    }
    if (
      object.validatorsWinner !== undefined &&
      object.validatorsWinner !== null
    ) {
      for (const e of object.validatorsWinner) {
        message.validatorsWinner.push(ValidatorsWinner.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.round !== undefined &&
      (obj.round = message.round ? Round.toJSON(message.round) : undefined);
    message.txnCounter !== undefined &&
      (obj.txnCounter = message.txnCounter
        ? TxnCounter.toJSON(message.txnCounter)
        : undefined);
    if (message.betList) {
      obj.betList = message.betList.map((e) => (e ? Bet.toJSON(e) : undefined));
    } else {
      obj.betList = [];
    }
    if (message.validatorsWinner) {
      obj.validatorsWinner = message.validatorsWinner.map((e) =>
        e ? ValidatorsWinner.toJSON(e) : undefined
      );
    } else {
      obj.validatorsWinner = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.betList = [];
    message.validatorsWinner = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.round !== undefined && object.round !== null) {
      message.round = Round.fromPartial(object.round);
    } else {
      message.round = undefined;
    }
    if (object.txnCounter !== undefined && object.txnCounter !== null) {
      message.txnCounter = TxnCounter.fromPartial(object.txnCounter);
    } else {
      message.txnCounter = undefined;
    }
    if (object.betList !== undefined && object.betList !== null) {
      for (const e of object.betList) {
        message.betList.push(Bet.fromPartial(e));
      }
    }
    if (
      object.validatorsWinner !== undefined &&
      object.validatorsWinner !== null
    ) {
      for (const e of object.validatorsWinner) {
        message.validatorsWinner.push(ValidatorsWinner.fromPartial(e));
      }
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

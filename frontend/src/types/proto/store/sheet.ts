/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Engine, engineFromJSON, engineToJSON, engineToNumber } from "./common";
import { DatabaseConfig } from "./database";

export const protobufPackage = "bytebase.store";

export interface SheetPayload {
  /** The snapshot of the database config when creating the sheet, be used to compare with the baseline_database_config and apply the diff to the database. */
  databaseConfig:
    | DatabaseConfig
    | undefined;
  /** The snapshot of the baseline database config when creating the sheet. */
  baselineDatabaseConfig:
    | DatabaseConfig
    | undefined;
  /** The SQL dialect. */
  engine: Engine;
  /** The start and end position of each command in the sheet statement. */
  commands: SheetCommand[];
}

export interface SheetCommand {
  start: number;
  end: number;
}

function createBaseSheetPayload(): SheetPayload {
  return {
    databaseConfig: undefined,
    baselineDatabaseConfig: undefined,
    engine: Engine.ENGINE_UNSPECIFIED,
    commands: [],
  };
}

export const SheetPayload = {
  encode(message: SheetPayload, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.databaseConfig !== undefined) {
      DatabaseConfig.encode(message.databaseConfig, writer.uint32(10).fork()).ldelim();
    }
    if (message.baselineDatabaseConfig !== undefined) {
      DatabaseConfig.encode(message.baselineDatabaseConfig, writer.uint32(18).fork()).ldelim();
    }
    if (message.engine !== Engine.ENGINE_UNSPECIFIED) {
      writer.uint32(24).int32(engineToNumber(message.engine));
    }
    for (const v of message.commands) {
      SheetCommand.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SheetPayload {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSheetPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.databaseConfig = DatabaseConfig.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.baselineDatabaseConfig = DatabaseConfig.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.engine = engineFromJSON(reader.int32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.commands.push(SheetCommand.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SheetPayload {
    return {
      databaseConfig: isSet(object.databaseConfig) ? DatabaseConfig.fromJSON(object.databaseConfig) : undefined,
      baselineDatabaseConfig: isSet(object.baselineDatabaseConfig)
        ? DatabaseConfig.fromJSON(object.baselineDatabaseConfig)
        : undefined,
      engine: isSet(object.engine) ? engineFromJSON(object.engine) : Engine.ENGINE_UNSPECIFIED,
      commands: globalThis.Array.isArray(object?.commands)
        ? object.commands.map((e: any) => SheetCommand.fromJSON(e))
        : [],
    };
  },

  toJSON(message: SheetPayload): unknown {
    const obj: any = {};
    if (message.databaseConfig !== undefined) {
      obj.databaseConfig = DatabaseConfig.toJSON(message.databaseConfig);
    }
    if (message.baselineDatabaseConfig !== undefined) {
      obj.baselineDatabaseConfig = DatabaseConfig.toJSON(message.baselineDatabaseConfig);
    }
    if (message.engine !== Engine.ENGINE_UNSPECIFIED) {
      obj.engine = engineToJSON(message.engine);
    }
    if (message.commands?.length) {
      obj.commands = message.commands.map((e) => SheetCommand.toJSON(e));
    }
    return obj;
  },

  create(base?: DeepPartial<SheetPayload>): SheetPayload {
    return SheetPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SheetPayload>): SheetPayload {
    const message = createBaseSheetPayload();
    message.databaseConfig = (object.databaseConfig !== undefined && object.databaseConfig !== null)
      ? DatabaseConfig.fromPartial(object.databaseConfig)
      : undefined;
    message.baselineDatabaseConfig =
      (object.baselineDatabaseConfig !== undefined && object.baselineDatabaseConfig !== null)
        ? DatabaseConfig.fromPartial(object.baselineDatabaseConfig)
        : undefined;
    message.engine = object.engine ?? Engine.ENGINE_UNSPECIFIED;
    message.commands = object.commands?.map((e) => SheetCommand.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSheetCommand(): SheetCommand {
  return { start: 0, end: 0 };
}

export const SheetCommand = {
  encode(message: SheetCommand, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.start !== 0) {
      writer.uint32(8).int32(message.start);
    }
    if (message.end !== 0) {
      writer.uint32(16).int32(message.end);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SheetCommand {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSheetCommand();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.start = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.end = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SheetCommand {
    return {
      start: isSet(object.start) ? globalThis.Number(object.start) : 0,
      end: isSet(object.end) ? globalThis.Number(object.end) : 0,
    };
  },

  toJSON(message: SheetCommand): unknown {
    const obj: any = {};
    if (message.start !== 0) {
      obj.start = Math.round(message.start);
    }
    if (message.end !== 0) {
      obj.end = Math.round(message.end);
    }
    return obj;
  },

  create(base?: DeepPartial<SheetCommand>): SheetCommand {
    return SheetCommand.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SheetCommand>): SheetCommand {
    const message = createBaseSheetCommand();
    message.start = object.start ?? 0;
    message.end = object.end ?? 0;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

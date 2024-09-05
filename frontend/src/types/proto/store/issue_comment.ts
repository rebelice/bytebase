/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "bytebase.store";

export interface IssueCommentPayload {
  comment: string;
  approval?: IssueCommentPayload_Approval | undefined;
  issueUpdate?: IssueCommentPayload_IssueUpdate | undefined;
  stageEnd?: IssueCommentPayload_StageEnd | undefined;
  taskUpdate?: IssueCommentPayload_TaskUpdate | undefined;
  taskPriorBackup?: IssueCommentPayload_TaskPriorBackup | undefined;
}

export interface IssueCommentPayload_Approval {
  status: IssueCommentPayload_Approval_Status;
}

export enum IssueCommentPayload_Approval_Status {
  STATUS_UNSPECIFIED = "STATUS_UNSPECIFIED",
  PENDING = "PENDING",
  APPROVED = "APPROVED",
  REJECTED = "REJECTED",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function issueCommentPayload_Approval_StatusFromJSON(object: any): IssueCommentPayload_Approval_Status {
  switch (object) {
    case 0:
    case "STATUS_UNSPECIFIED":
      return IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED;
    case 1:
    case "PENDING":
      return IssueCommentPayload_Approval_Status.PENDING;
    case 2:
    case "APPROVED":
      return IssueCommentPayload_Approval_Status.APPROVED;
    case 3:
    case "REJECTED":
      return IssueCommentPayload_Approval_Status.REJECTED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IssueCommentPayload_Approval_Status.UNRECOGNIZED;
  }
}

export function issueCommentPayload_Approval_StatusToJSON(object: IssueCommentPayload_Approval_Status): string {
  switch (object) {
    case IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED:
      return "STATUS_UNSPECIFIED";
    case IssueCommentPayload_Approval_Status.PENDING:
      return "PENDING";
    case IssueCommentPayload_Approval_Status.APPROVED:
      return "APPROVED";
    case IssueCommentPayload_Approval_Status.REJECTED:
      return "REJECTED";
    case IssueCommentPayload_Approval_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function issueCommentPayload_Approval_StatusToNumber(object: IssueCommentPayload_Approval_Status): number {
  switch (object) {
    case IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED:
      return 0;
    case IssueCommentPayload_Approval_Status.PENDING:
      return 1;
    case IssueCommentPayload_Approval_Status.APPROVED:
      return 2;
    case IssueCommentPayload_Approval_Status.REJECTED:
      return 3;
    case IssueCommentPayload_Approval_Status.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface IssueCommentPayload_IssueUpdate {
  fromTitle?: string | undefined;
  toTitle?: string | undefined;
  fromDescription?: string | undefined;
  toDescription?: string | undefined;
  fromStatus?: IssueCommentPayload_IssueUpdate_IssueStatus | undefined;
  toStatus?: IssueCommentPayload_IssueUpdate_IssueStatus | undefined;
  fromLabels: string[];
  toLabels: string[];
}

export enum IssueCommentPayload_IssueUpdate_IssueStatus {
  ISSUE_STATUS_UNSPECIFIED = "ISSUE_STATUS_UNSPECIFIED",
  OPEN = "OPEN",
  DONE = "DONE",
  CANCELED = "CANCELED",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function issueCommentPayload_IssueUpdate_IssueStatusFromJSON(
  object: any,
): IssueCommentPayload_IssueUpdate_IssueStatus {
  switch (object) {
    case 0:
    case "ISSUE_STATUS_UNSPECIFIED":
      return IssueCommentPayload_IssueUpdate_IssueStatus.ISSUE_STATUS_UNSPECIFIED;
    case 1:
    case "OPEN":
      return IssueCommentPayload_IssueUpdate_IssueStatus.OPEN;
    case 2:
    case "DONE":
      return IssueCommentPayload_IssueUpdate_IssueStatus.DONE;
    case 3:
    case "CANCELED":
      return IssueCommentPayload_IssueUpdate_IssueStatus.CANCELED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IssueCommentPayload_IssueUpdate_IssueStatus.UNRECOGNIZED;
  }
}

export function issueCommentPayload_IssueUpdate_IssueStatusToJSON(
  object: IssueCommentPayload_IssueUpdate_IssueStatus,
): string {
  switch (object) {
    case IssueCommentPayload_IssueUpdate_IssueStatus.ISSUE_STATUS_UNSPECIFIED:
      return "ISSUE_STATUS_UNSPECIFIED";
    case IssueCommentPayload_IssueUpdate_IssueStatus.OPEN:
      return "OPEN";
    case IssueCommentPayload_IssueUpdate_IssueStatus.DONE:
      return "DONE";
    case IssueCommentPayload_IssueUpdate_IssueStatus.CANCELED:
      return "CANCELED";
    case IssueCommentPayload_IssueUpdate_IssueStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function issueCommentPayload_IssueUpdate_IssueStatusToNumber(
  object: IssueCommentPayload_IssueUpdate_IssueStatus,
): number {
  switch (object) {
    case IssueCommentPayload_IssueUpdate_IssueStatus.ISSUE_STATUS_UNSPECIFIED:
      return 0;
    case IssueCommentPayload_IssueUpdate_IssueStatus.OPEN:
      return 1;
    case IssueCommentPayload_IssueUpdate_IssueStatus.DONE:
      return 2;
    case IssueCommentPayload_IssueUpdate_IssueStatus.CANCELED:
      return 3;
    case IssueCommentPayload_IssueUpdate_IssueStatus.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface IssueCommentPayload_StageEnd {
  stage: string;
}

export interface IssueCommentPayload_TaskUpdate {
  tasks: string[];
  /** Format: projects/{project}/sheets/{sheet} */
  fromSheet?:
    | string
    | undefined;
  /** Format: projects/{project}/sheets/{sheet} */
  toSheet?: string | undefined;
  fromEarliestAllowedTime?: Date | undefined;
  toEarliestAllowedTime?: Date | undefined;
  toStatus?: IssueCommentPayload_TaskUpdate_Status | undefined;
}

export enum IssueCommentPayload_TaskUpdate_Status {
  STATUS_UNSPECIFIED = "STATUS_UNSPECIFIED",
  PENDING = "PENDING",
  RUNNING = "RUNNING",
  DONE = "DONE",
  FAILED = "FAILED",
  SKIPPED = "SKIPPED",
  CANCELED = "CANCELED",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function issueCommentPayload_TaskUpdate_StatusFromJSON(object: any): IssueCommentPayload_TaskUpdate_Status {
  switch (object) {
    case 0:
    case "STATUS_UNSPECIFIED":
      return IssueCommentPayload_TaskUpdate_Status.STATUS_UNSPECIFIED;
    case 1:
    case "PENDING":
      return IssueCommentPayload_TaskUpdate_Status.PENDING;
    case 2:
    case "RUNNING":
      return IssueCommentPayload_TaskUpdate_Status.RUNNING;
    case 3:
    case "DONE":
      return IssueCommentPayload_TaskUpdate_Status.DONE;
    case 4:
    case "FAILED":
      return IssueCommentPayload_TaskUpdate_Status.FAILED;
    case 5:
    case "SKIPPED":
      return IssueCommentPayload_TaskUpdate_Status.SKIPPED;
    case 6:
    case "CANCELED":
      return IssueCommentPayload_TaskUpdate_Status.CANCELED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IssueCommentPayload_TaskUpdate_Status.UNRECOGNIZED;
  }
}

export function issueCommentPayload_TaskUpdate_StatusToJSON(object: IssueCommentPayload_TaskUpdate_Status): string {
  switch (object) {
    case IssueCommentPayload_TaskUpdate_Status.STATUS_UNSPECIFIED:
      return "STATUS_UNSPECIFIED";
    case IssueCommentPayload_TaskUpdate_Status.PENDING:
      return "PENDING";
    case IssueCommentPayload_TaskUpdate_Status.RUNNING:
      return "RUNNING";
    case IssueCommentPayload_TaskUpdate_Status.DONE:
      return "DONE";
    case IssueCommentPayload_TaskUpdate_Status.FAILED:
      return "FAILED";
    case IssueCommentPayload_TaskUpdate_Status.SKIPPED:
      return "SKIPPED";
    case IssueCommentPayload_TaskUpdate_Status.CANCELED:
      return "CANCELED";
    case IssueCommentPayload_TaskUpdate_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function issueCommentPayload_TaskUpdate_StatusToNumber(object: IssueCommentPayload_TaskUpdate_Status): number {
  switch (object) {
    case IssueCommentPayload_TaskUpdate_Status.STATUS_UNSPECIFIED:
      return 0;
    case IssueCommentPayload_TaskUpdate_Status.PENDING:
      return 1;
    case IssueCommentPayload_TaskUpdate_Status.RUNNING:
      return 2;
    case IssueCommentPayload_TaskUpdate_Status.DONE:
      return 3;
    case IssueCommentPayload_TaskUpdate_Status.FAILED:
      return 4;
    case IssueCommentPayload_TaskUpdate_Status.SKIPPED:
      return 5;
    case IssueCommentPayload_TaskUpdate_Status.CANCELED:
      return 6;
    case IssueCommentPayload_TaskUpdate_Status.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface IssueCommentPayload_TaskPriorBackup {
  task: string;
  tables: IssueCommentPayload_TaskPriorBackup_Table[];
  originalLine?: number | undefined;
  database: string;
  error: string;
}

export interface IssueCommentPayload_TaskPriorBackup_Table {
  schema: string;
  table: string;
}

function createBaseIssueCommentPayload(): IssueCommentPayload {
  return {
    comment: "",
    approval: undefined,
    issueUpdate: undefined,
    stageEnd: undefined,
    taskUpdate: undefined,
    taskPriorBackup: undefined,
  };
}

export const IssueCommentPayload = {
  encode(message: IssueCommentPayload, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.comment !== "") {
      writer.uint32(10).string(message.comment);
    }
    if (message.approval !== undefined) {
      IssueCommentPayload_Approval.encode(message.approval, writer.uint32(18).fork()).ldelim();
    }
    if (message.issueUpdate !== undefined) {
      IssueCommentPayload_IssueUpdate.encode(message.issueUpdate, writer.uint32(26).fork()).ldelim();
    }
    if (message.stageEnd !== undefined) {
      IssueCommentPayload_StageEnd.encode(message.stageEnd, writer.uint32(34).fork()).ldelim();
    }
    if (message.taskUpdate !== undefined) {
      IssueCommentPayload_TaskUpdate.encode(message.taskUpdate, writer.uint32(42).fork()).ldelim();
    }
    if (message.taskPriorBackup !== undefined) {
      IssueCommentPayload_TaskPriorBackup.encode(message.taskPriorBackup, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.comment = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.approval = IssueCommentPayload_Approval.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.issueUpdate = IssueCommentPayload_IssueUpdate.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.stageEnd = IssueCommentPayload_StageEnd.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.taskUpdate = IssueCommentPayload_TaskUpdate.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.taskPriorBackup = IssueCommentPayload_TaskPriorBackup.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload {
    return {
      comment: isSet(object.comment) ? globalThis.String(object.comment) : "",
      approval: isSet(object.approval) ? IssueCommentPayload_Approval.fromJSON(object.approval) : undefined,
      issueUpdate: isSet(object.issueUpdate) ? IssueCommentPayload_IssueUpdate.fromJSON(object.issueUpdate) : undefined,
      stageEnd: isSet(object.stageEnd) ? IssueCommentPayload_StageEnd.fromJSON(object.stageEnd) : undefined,
      taskUpdate: isSet(object.taskUpdate) ? IssueCommentPayload_TaskUpdate.fromJSON(object.taskUpdate) : undefined,
      taskPriorBackup: isSet(object.taskPriorBackup)
        ? IssueCommentPayload_TaskPriorBackup.fromJSON(object.taskPriorBackup)
        : undefined,
    };
  },

  toJSON(message: IssueCommentPayload): unknown {
    const obj: any = {};
    if (message.comment !== "") {
      obj.comment = message.comment;
    }
    if (message.approval !== undefined) {
      obj.approval = IssueCommentPayload_Approval.toJSON(message.approval);
    }
    if (message.issueUpdate !== undefined) {
      obj.issueUpdate = IssueCommentPayload_IssueUpdate.toJSON(message.issueUpdate);
    }
    if (message.stageEnd !== undefined) {
      obj.stageEnd = IssueCommentPayload_StageEnd.toJSON(message.stageEnd);
    }
    if (message.taskUpdate !== undefined) {
      obj.taskUpdate = IssueCommentPayload_TaskUpdate.toJSON(message.taskUpdate);
    }
    if (message.taskPriorBackup !== undefined) {
      obj.taskPriorBackup = IssueCommentPayload_TaskPriorBackup.toJSON(message.taskPriorBackup);
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload>): IssueCommentPayload {
    return IssueCommentPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IssueCommentPayload>): IssueCommentPayload {
    const message = createBaseIssueCommentPayload();
    message.comment = object.comment ?? "";
    message.approval = (object.approval !== undefined && object.approval !== null)
      ? IssueCommentPayload_Approval.fromPartial(object.approval)
      : undefined;
    message.issueUpdate = (object.issueUpdate !== undefined && object.issueUpdate !== null)
      ? IssueCommentPayload_IssueUpdate.fromPartial(object.issueUpdate)
      : undefined;
    message.stageEnd = (object.stageEnd !== undefined && object.stageEnd !== null)
      ? IssueCommentPayload_StageEnd.fromPartial(object.stageEnd)
      : undefined;
    message.taskUpdate = (object.taskUpdate !== undefined && object.taskUpdate !== null)
      ? IssueCommentPayload_TaskUpdate.fromPartial(object.taskUpdate)
      : undefined;
    message.taskPriorBackup = (object.taskPriorBackup !== undefined && object.taskPriorBackup !== null)
      ? IssueCommentPayload_TaskPriorBackup.fromPartial(object.taskPriorBackup)
      : undefined;
    return message;
  },
};

function createBaseIssueCommentPayload_Approval(): IssueCommentPayload_Approval {
  return { status: IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED };
}

export const IssueCommentPayload_Approval = {
  encode(message: IssueCommentPayload_Approval, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED) {
      writer.uint32(8).int32(issueCommentPayload_Approval_StatusToNumber(message.status));
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload_Approval {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload_Approval();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.status = issueCommentPayload_Approval_StatusFromJSON(reader.int32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload_Approval {
    return {
      status: isSet(object.status)
        ? issueCommentPayload_Approval_StatusFromJSON(object.status)
        : IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED,
    };
  },

  toJSON(message: IssueCommentPayload_Approval): unknown {
    const obj: any = {};
    if (message.status !== IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED) {
      obj.status = issueCommentPayload_Approval_StatusToJSON(message.status);
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload_Approval>): IssueCommentPayload_Approval {
    return IssueCommentPayload_Approval.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IssueCommentPayload_Approval>): IssueCommentPayload_Approval {
    const message = createBaseIssueCommentPayload_Approval();
    message.status = object.status ?? IssueCommentPayload_Approval_Status.STATUS_UNSPECIFIED;
    return message;
  },
};

function createBaseIssueCommentPayload_IssueUpdate(): IssueCommentPayload_IssueUpdate {
  return {
    fromTitle: undefined,
    toTitle: undefined,
    fromDescription: undefined,
    toDescription: undefined,
    fromStatus: undefined,
    toStatus: undefined,
    fromLabels: [],
    toLabels: [],
  };
}

export const IssueCommentPayload_IssueUpdate = {
  encode(message: IssueCommentPayload_IssueUpdate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fromTitle !== undefined) {
      writer.uint32(10).string(message.fromTitle);
    }
    if (message.toTitle !== undefined) {
      writer.uint32(18).string(message.toTitle);
    }
    if (message.fromDescription !== undefined) {
      writer.uint32(26).string(message.fromDescription);
    }
    if (message.toDescription !== undefined) {
      writer.uint32(34).string(message.toDescription);
    }
    if (message.fromStatus !== undefined) {
      writer.uint32(40).int32(issueCommentPayload_IssueUpdate_IssueStatusToNumber(message.fromStatus));
    }
    if (message.toStatus !== undefined) {
      writer.uint32(48).int32(issueCommentPayload_IssueUpdate_IssueStatusToNumber(message.toStatus));
    }
    for (const v of message.fromLabels) {
      writer.uint32(58).string(v!);
    }
    for (const v of message.toLabels) {
      writer.uint32(66).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload_IssueUpdate {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload_IssueUpdate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.fromTitle = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.toTitle = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.fromDescription = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.toDescription = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.fromStatus = issueCommentPayload_IssueUpdate_IssueStatusFromJSON(reader.int32());
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.toStatus = issueCommentPayload_IssueUpdate_IssueStatusFromJSON(reader.int32());
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.fromLabels.push(reader.string());
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.toLabels.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload_IssueUpdate {
    return {
      fromTitle: isSet(object.fromTitle) ? globalThis.String(object.fromTitle) : undefined,
      toTitle: isSet(object.toTitle) ? globalThis.String(object.toTitle) : undefined,
      fromDescription: isSet(object.fromDescription) ? globalThis.String(object.fromDescription) : undefined,
      toDescription: isSet(object.toDescription) ? globalThis.String(object.toDescription) : undefined,
      fromStatus: isSet(object.fromStatus)
        ? issueCommentPayload_IssueUpdate_IssueStatusFromJSON(object.fromStatus)
        : undefined,
      toStatus: isSet(object.toStatus)
        ? issueCommentPayload_IssueUpdate_IssueStatusFromJSON(object.toStatus)
        : undefined,
      fromLabels: globalThis.Array.isArray(object?.fromLabels)
        ? object.fromLabels.map((e: any) => globalThis.String(e))
        : [],
      toLabels: globalThis.Array.isArray(object?.toLabels) ? object.toLabels.map((e: any) => globalThis.String(e)) : [],
    };
  },

  toJSON(message: IssueCommentPayload_IssueUpdate): unknown {
    const obj: any = {};
    if (message.fromTitle !== undefined) {
      obj.fromTitle = message.fromTitle;
    }
    if (message.toTitle !== undefined) {
      obj.toTitle = message.toTitle;
    }
    if (message.fromDescription !== undefined) {
      obj.fromDescription = message.fromDescription;
    }
    if (message.toDescription !== undefined) {
      obj.toDescription = message.toDescription;
    }
    if (message.fromStatus !== undefined) {
      obj.fromStatus = issueCommentPayload_IssueUpdate_IssueStatusToJSON(message.fromStatus);
    }
    if (message.toStatus !== undefined) {
      obj.toStatus = issueCommentPayload_IssueUpdate_IssueStatusToJSON(message.toStatus);
    }
    if (message.fromLabels?.length) {
      obj.fromLabels = message.fromLabels;
    }
    if (message.toLabels?.length) {
      obj.toLabels = message.toLabels;
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload_IssueUpdate>): IssueCommentPayload_IssueUpdate {
    return IssueCommentPayload_IssueUpdate.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IssueCommentPayload_IssueUpdate>): IssueCommentPayload_IssueUpdate {
    const message = createBaseIssueCommentPayload_IssueUpdate();
    message.fromTitle = object.fromTitle ?? undefined;
    message.toTitle = object.toTitle ?? undefined;
    message.fromDescription = object.fromDescription ?? undefined;
    message.toDescription = object.toDescription ?? undefined;
    message.fromStatus = object.fromStatus ?? undefined;
    message.toStatus = object.toStatus ?? undefined;
    message.fromLabels = object.fromLabels?.map((e) => e) || [];
    message.toLabels = object.toLabels?.map((e) => e) || [];
    return message;
  },
};

function createBaseIssueCommentPayload_StageEnd(): IssueCommentPayload_StageEnd {
  return { stage: "" };
}

export const IssueCommentPayload_StageEnd = {
  encode(message: IssueCommentPayload_StageEnd, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stage !== "") {
      writer.uint32(10).string(message.stage);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload_StageEnd {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload_StageEnd();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.stage = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload_StageEnd {
    return { stage: isSet(object.stage) ? globalThis.String(object.stage) : "" };
  },

  toJSON(message: IssueCommentPayload_StageEnd): unknown {
    const obj: any = {};
    if (message.stage !== "") {
      obj.stage = message.stage;
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload_StageEnd>): IssueCommentPayload_StageEnd {
    return IssueCommentPayload_StageEnd.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IssueCommentPayload_StageEnd>): IssueCommentPayload_StageEnd {
    const message = createBaseIssueCommentPayload_StageEnd();
    message.stage = object.stage ?? "";
    return message;
  },
};

function createBaseIssueCommentPayload_TaskUpdate(): IssueCommentPayload_TaskUpdate {
  return {
    tasks: [],
    fromSheet: undefined,
    toSheet: undefined,
    fromEarliestAllowedTime: undefined,
    toEarliestAllowedTime: undefined,
    toStatus: undefined,
  };
}

export const IssueCommentPayload_TaskUpdate = {
  encode(message: IssueCommentPayload_TaskUpdate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tasks) {
      writer.uint32(10).string(v!);
    }
    if (message.fromSheet !== undefined) {
      writer.uint32(18).string(message.fromSheet);
    }
    if (message.toSheet !== undefined) {
      writer.uint32(26).string(message.toSheet);
    }
    if (message.fromEarliestAllowedTime !== undefined) {
      Timestamp.encode(toTimestamp(message.fromEarliestAllowedTime), writer.uint32(34).fork()).ldelim();
    }
    if (message.toEarliestAllowedTime !== undefined) {
      Timestamp.encode(toTimestamp(message.toEarliestAllowedTime), writer.uint32(42).fork()).ldelim();
    }
    if (message.toStatus !== undefined) {
      writer.uint32(48).int32(issueCommentPayload_TaskUpdate_StatusToNumber(message.toStatus));
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload_TaskUpdate {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload_TaskUpdate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tasks.push(reader.string());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.fromSheet = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.toSheet = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.fromEarliestAllowedTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.toEarliestAllowedTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.toStatus = issueCommentPayload_TaskUpdate_StatusFromJSON(reader.int32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload_TaskUpdate {
    return {
      tasks: globalThis.Array.isArray(object?.tasks) ? object.tasks.map((e: any) => globalThis.String(e)) : [],
      fromSheet: isSet(object.fromSheet) ? globalThis.String(object.fromSheet) : undefined,
      toSheet: isSet(object.toSheet) ? globalThis.String(object.toSheet) : undefined,
      fromEarliestAllowedTime: isSet(object.fromEarliestAllowedTime)
        ? fromJsonTimestamp(object.fromEarliestAllowedTime)
        : undefined,
      toEarliestAllowedTime: isSet(object.toEarliestAllowedTime)
        ? fromJsonTimestamp(object.toEarliestAllowedTime)
        : undefined,
      toStatus: isSet(object.toStatus) ? issueCommentPayload_TaskUpdate_StatusFromJSON(object.toStatus) : undefined,
    };
  },

  toJSON(message: IssueCommentPayload_TaskUpdate): unknown {
    const obj: any = {};
    if (message.tasks?.length) {
      obj.tasks = message.tasks;
    }
    if (message.fromSheet !== undefined) {
      obj.fromSheet = message.fromSheet;
    }
    if (message.toSheet !== undefined) {
      obj.toSheet = message.toSheet;
    }
    if (message.fromEarliestAllowedTime !== undefined) {
      obj.fromEarliestAllowedTime = message.fromEarliestAllowedTime.toISOString();
    }
    if (message.toEarliestAllowedTime !== undefined) {
      obj.toEarliestAllowedTime = message.toEarliestAllowedTime.toISOString();
    }
    if (message.toStatus !== undefined) {
      obj.toStatus = issueCommentPayload_TaskUpdate_StatusToJSON(message.toStatus);
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload_TaskUpdate>): IssueCommentPayload_TaskUpdate {
    return IssueCommentPayload_TaskUpdate.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IssueCommentPayload_TaskUpdate>): IssueCommentPayload_TaskUpdate {
    const message = createBaseIssueCommentPayload_TaskUpdate();
    message.tasks = object.tasks?.map((e) => e) || [];
    message.fromSheet = object.fromSheet ?? undefined;
    message.toSheet = object.toSheet ?? undefined;
    message.fromEarliestAllowedTime = object.fromEarliestAllowedTime ?? undefined;
    message.toEarliestAllowedTime = object.toEarliestAllowedTime ?? undefined;
    message.toStatus = object.toStatus ?? undefined;
    return message;
  },
};

function createBaseIssueCommentPayload_TaskPriorBackup(): IssueCommentPayload_TaskPriorBackup {
  return { task: "", tables: [], originalLine: undefined, database: "", error: "" };
}

export const IssueCommentPayload_TaskPriorBackup = {
  encode(message: IssueCommentPayload_TaskPriorBackup, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.task !== "") {
      writer.uint32(10).string(message.task);
    }
    for (const v of message.tables) {
      IssueCommentPayload_TaskPriorBackup_Table.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.originalLine !== undefined) {
      writer.uint32(24).int32(message.originalLine);
    }
    if (message.database !== "") {
      writer.uint32(34).string(message.database);
    }
    if (message.error !== "") {
      writer.uint32(42).string(message.error);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload_TaskPriorBackup {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload_TaskPriorBackup();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.task = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.tables.push(IssueCommentPayload_TaskPriorBackup_Table.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.originalLine = reader.int32();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.database = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.error = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload_TaskPriorBackup {
    return {
      task: isSet(object.task) ? globalThis.String(object.task) : "",
      tables: globalThis.Array.isArray(object?.tables)
        ? object.tables.map((e: any) => IssueCommentPayload_TaskPriorBackup_Table.fromJSON(e))
        : [],
      originalLine: isSet(object.originalLine) ? globalThis.Number(object.originalLine) : undefined,
      database: isSet(object.database) ? globalThis.String(object.database) : "",
      error: isSet(object.error) ? globalThis.String(object.error) : "",
    };
  },

  toJSON(message: IssueCommentPayload_TaskPriorBackup): unknown {
    const obj: any = {};
    if (message.task !== "") {
      obj.task = message.task;
    }
    if (message.tables?.length) {
      obj.tables = message.tables.map((e) => IssueCommentPayload_TaskPriorBackup_Table.toJSON(e));
    }
    if (message.originalLine !== undefined) {
      obj.originalLine = Math.round(message.originalLine);
    }
    if (message.database !== "") {
      obj.database = message.database;
    }
    if (message.error !== "") {
      obj.error = message.error;
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload_TaskPriorBackup>): IssueCommentPayload_TaskPriorBackup {
    return IssueCommentPayload_TaskPriorBackup.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IssueCommentPayload_TaskPriorBackup>): IssueCommentPayload_TaskPriorBackup {
    const message = createBaseIssueCommentPayload_TaskPriorBackup();
    message.task = object.task ?? "";
    message.tables = object.tables?.map((e) => IssueCommentPayload_TaskPriorBackup_Table.fromPartial(e)) || [];
    message.originalLine = object.originalLine ?? undefined;
    message.database = object.database ?? "";
    message.error = object.error ?? "";
    return message;
  },
};

function createBaseIssueCommentPayload_TaskPriorBackup_Table(): IssueCommentPayload_TaskPriorBackup_Table {
  return { schema: "", table: "" };
}

export const IssueCommentPayload_TaskPriorBackup_Table = {
  encode(message: IssueCommentPayload_TaskPriorBackup_Table, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.schema !== "") {
      writer.uint32(10).string(message.schema);
    }
    if (message.table !== "") {
      writer.uint32(18).string(message.table);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueCommentPayload_TaskPriorBackup_Table {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueCommentPayload_TaskPriorBackup_Table();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.schema = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.table = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IssueCommentPayload_TaskPriorBackup_Table {
    return {
      schema: isSet(object.schema) ? globalThis.String(object.schema) : "",
      table: isSet(object.table) ? globalThis.String(object.table) : "",
    };
  },

  toJSON(message: IssueCommentPayload_TaskPriorBackup_Table): unknown {
    const obj: any = {};
    if (message.schema !== "") {
      obj.schema = message.schema;
    }
    if (message.table !== "") {
      obj.table = message.table;
    }
    return obj;
  },

  create(base?: DeepPartial<IssueCommentPayload_TaskPriorBackup_Table>): IssueCommentPayload_TaskPriorBackup_Table {
    return IssueCommentPayload_TaskPriorBackup_Table.fromPartial(base ?? {});
  },
  fromPartial(
    object: DeepPartial<IssueCommentPayload_TaskPriorBackup_Table>,
  ): IssueCommentPayload_TaskPriorBackup_Table {
    const message = createBaseIssueCommentPayload_TaskPriorBackup_Table();
    message.schema = object.schema ?? "";
    message.table = object.table ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof globalThis.Date) {
    return o;
  } else if (typeof o === "string") {
    return new globalThis.Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

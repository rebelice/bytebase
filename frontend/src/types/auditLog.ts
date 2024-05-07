import { LogEntity_Action } from "@/types/proto/v1/logging_service";

export const AuditActivityTypeList = [
  // Workspace level.
  //
  // Member related
  LogEntity_Action.ACTION_MEMBER_CREATE,
  LogEntity_Action.ACTION_MEMBER_ROLE_UPDATE,
  LogEntity_Action.ACTION_MEMBER_ACTIVATE,
  LogEntity_Action.ACTION_MEMBER_DEACTIVE,
  // Project level.
  //
  // Project related
  LogEntity_Action.ACTION_PROJECT_DATABASE_TRANSFER,
  // SQL Editor related.
  LogEntity_Action.ACTION_DATABASE_SQL_EDITOR_QUERY,
  LogEntity_Action.ACTION_DATABASE_SQL_EXPORT,
];

export const AuditActivityTypeI18nNameMap: Map<LogEntity_Action, string> =
  new Map([
    [
      LogEntity_Action.ACTION_MEMBER_CREATE,
      "audit-log.type.workspace.member-create",
    ],
    [
      LogEntity_Action.ACTION_MEMBER_ROLE_UPDATE,
      "audit-log.type.workspace.member-role-update",
    ],
    [
      LogEntity_Action.ACTION_MEMBER_ACTIVATE,
      "audit-log.type.workspace.member-activate",
    ],
    [
      LogEntity_Action.ACTION_MEMBER_DEACTIVE,
      "audit-log.type.workspace.member-deactivate",
    ],
    [
      LogEntity_Action.ACTION_PROJECT_DATABASE_TRANSFER,
      "audit-log.type.project.database-transfer",
    ],
    [
      LogEntity_Action.ACTION_DATABASE_SQL_EDITOR_QUERY,
      "audit-log.type.sql-editor-query",
    ],
    [LogEntity_Action.ACTION_DATABASE_SQL_EXPORT, "audit-log.type.sql-export"],
  ]);

export interface SearchAuditLogsParams {
  parent?: string;
  creatorEmail?: string;
  order?: "asc" | "desc";
  pageSize?: number;
  pageToken?: string;
}

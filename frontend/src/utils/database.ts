import { keyBy } from "lodash-es";
import { Engine } from "@/types/proto/v1/common";
import type { Environment as EnvironmentV1 } from "@/types/proto/v1/environment_service";
import type { Database, Environment } from "../types";
import { semverCompare } from "./util";

// Sort the list to put prod items first.
export function sortDatabaseList(
  list: Database[],
  environmentList: Environment[]
): Database[] {
  return list.sort((a: Database, b: Database) => {
    let aEnvIndex = -1;
    let bEnvIndex = -1;

    for (let i = 0; i < environmentList.length; i++) {
      if (environmentList[i].id == a.instance.environment.id) {
        aEnvIndex = i;
      }
      if (environmentList[i].id == b.instance.environment.id) {
        bEnvIndex = i;
      }

      if (aEnvIndex != -1 && bEnvIndex != -1) {
        break;
      }
    }
    return bEnvIndex - aEnvIndex;
  });
}

// Sort the list to put prod items first.
export function sortDatabaseListByEnvironmentV1(
  list: Database[],
  environmentList: EnvironmentV1[]
): Database[] {
  const environmentMap = keyBy(environmentList, (env) => env.uid);
  return list.sort((a: Database, b: Database) => {
    const aEnvOrder =
      environmentMap[String(a.instance.environment.id)]?.order ?? -1;
    const bEnvOrder =
      environmentMap[String(b.instance.environment.id)]?.order ?? -1;

    return bEnvOrder - aEnvOrder;
  });
}

const MIN_GHOST_SUPPORT_MYSQL_VERSION = "5.7.0";

export function allowGhostMigration(databaseList: Database[]): boolean {
  return databaseList.every((db) => {
    return (
      db.instance.engine === "MYSQL" &&
      semverCompare(db.instance.engineVersion, MIN_GHOST_SUPPORT_MYSQL_VERSION)
    );
  });
}

type DatabaseFilterFields = "name" | "project" | "instance" | "environment";
export function filterDatabaseByKeyword(
  db: Database,
  keyword: string,
  columns: DatabaseFilterFields[] = ["name"]
): boolean {
  keyword = keyword.trim().toLowerCase();
  if (!keyword) {
    // Skip the filter
    return true;
  }

  if (columns.includes("name") && db.name.toLowerCase().includes(keyword)) {
    return true;
  }

  if (
    columns.includes("project") &&
    db.project.name.toLowerCase().includes(keyword)
  ) {
    return true;
  }

  if (
    columns.includes("instance") &&
    db.instance.name.toLowerCase().includes(keyword)
  ) {
    return true;
  }

  if (
    columns.includes("environment") &&
    db.instance.environment.name.toLowerCase().includes(keyword)
  ) {
    return true;
  }

  return false;
}

export function isArchivedDatabase(db: Database): boolean {
  if (db.instance.rowStatus === "ARCHIVED") {
    return true;
  }
  if (db.instance.environment.rowStatus === "ARCHIVED") {
    return true;
  }
  return false;
}

export const hasSchemaProperty = (databaseEngine: Engine) => {
  return (
    databaseEngine === Engine.POSTGRES ||
    databaseEngine === Engine.SNOWFLAKE ||
    databaseEngine === Engine.ORACLE ||
    databaseEngine === Engine.MSSQL ||
    databaseEngine === Engine.REDSHIFT ||
    databaseEngine === Engine.RISINGWAVE
  );
};

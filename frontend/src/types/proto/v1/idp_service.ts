/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";
import { FieldMask } from "../google/protobuf/field_mask";
import { State, stateFromJSON, stateToJSON } from "./common";

export const protobufPackage = "bytebase.v1";

export enum IdentityProviderType {
  IDENTITY_PROVIDER_TYPE_UNSPECIFIED = 0,
  OAUTH2 = 1,
  OIDC = 2,
  LDAP = 3,
  UNRECOGNIZED = -1,
}

export function identityProviderTypeFromJSON(object: any): IdentityProviderType {
  switch (object) {
    case 0:
    case "IDENTITY_PROVIDER_TYPE_UNSPECIFIED":
      return IdentityProviderType.IDENTITY_PROVIDER_TYPE_UNSPECIFIED;
    case 1:
    case "OAUTH2":
      return IdentityProviderType.OAUTH2;
    case 2:
    case "OIDC":
      return IdentityProviderType.OIDC;
    case 3:
    case "LDAP":
      return IdentityProviderType.LDAP;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IdentityProviderType.UNRECOGNIZED;
  }
}

export function identityProviderTypeToJSON(object: IdentityProviderType): string {
  switch (object) {
    case IdentityProviderType.IDENTITY_PROVIDER_TYPE_UNSPECIFIED:
      return "IDENTITY_PROVIDER_TYPE_UNSPECIFIED";
    case IdentityProviderType.OAUTH2:
      return "OAUTH2";
    case IdentityProviderType.OIDC:
      return "OIDC";
    case IdentityProviderType.LDAP:
      return "LDAP";
    case IdentityProviderType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum OAuth2AuthStyle {
  OAUTH2_AUTH_STYLE_UNSPECIFIED = 0,
  /**
   * IN_PARAMS - IN_PARAMS sends the "client_id" and "client_secret" in the POST body
   * as application/x-www-form-urlencoded parameters.
   */
  IN_PARAMS = 1,
  /**
   * IN_HEADER - IN_HEADER sends the client_id and client_password using HTTP Basic Authorization.
   * This is an optional style described in the OAuth2 RFC 6749 section 2.3.1.
   */
  IN_HEADER = 2,
  UNRECOGNIZED = -1,
}

export function oAuth2AuthStyleFromJSON(object: any): OAuth2AuthStyle {
  switch (object) {
    case 0:
    case "OAUTH2_AUTH_STYLE_UNSPECIFIED":
      return OAuth2AuthStyle.OAUTH2_AUTH_STYLE_UNSPECIFIED;
    case 1:
    case "IN_PARAMS":
      return OAuth2AuthStyle.IN_PARAMS;
    case 2:
    case "IN_HEADER":
      return OAuth2AuthStyle.IN_HEADER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OAuth2AuthStyle.UNRECOGNIZED;
  }
}

export function oAuth2AuthStyleToJSON(object: OAuth2AuthStyle): string {
  switch (object) {
    case OAuth2AuthStyle.OAUTH2_AUTH_STYLE_UNSPECIFIED:
      return "OAUTH2_AUTH_STYLE_UNSPECIFIED";
    case OAuth2AuthStyle.IN_PARAMS:
      return "IN_PARAMS";
    case OAuth2AuthStyle.IN_HEADER:
      return "IN_HEADER";
    case OAuth2AuthStyle.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface GetIdentityProviderRequest {
  name: string;
}

export interface ListIdentityProvidersRequest {
  /**
   * The maximum number of identity providers to return. The service may return fewer than
   * this value.
   * If unspecified, at most 50 will be returned.
   * The maximum value is 1000; values above 1000 will be coerced to 1000.
   */
  pageSize: number;
  /**
   * A page token, received from a previous `ListIdentityProviders` call.
   * Provide this to retrieve the subsequent page.
   *
   * When paginating, all other parameters provided to `ListIdentityProviders` must match
   * the call that provided the page token.
   */
  pageToken: string;
  /** Show deleted identity providers if specified. */
  showDeleted: boolean;
}

export interface ListIdentityProvidersResponse {
  /** The identity providers from the specified request. */
  identityProviders: IdentityProvider[];
  /**
   * A token, which can be sent as `page_token` to retrieve the next page.
   * If this field is omitted, there are no subsequent pages.
   */
  nextPageToken: string;
}

export interface CreateIdentityProviderRequest {
  /** The identity provider to create. */
  identityProvider?:
    | IdentityProvider
    | undefined;
  /**
   * The ID to use for the identity provider, which will become the final component of
   * the identity provider's resource name.
   *
   * This value should be 4-63 characters, and valid characters
   * are /[a-z][0-9]-/.
   */
  identityProviderId: string;
}

export interface UpdateIdentityProviderRequest {
  /**
   * The identity provider to update.
   *
   * The identity provider's `name` field is used to identify the identity provider to update.
   * Format: idps/{identity_provider}
   */
  identityProvider?:
    | IdentityProvider
    | undefined;
  /** The list of fields to update. */
  updateMask?: string[] | undefined;
}

export interface DeleteIdentityProviderRequest {
  /**
   * The name of the identity provider to delete.
   * Format: idps/{identity_provider}
   */
  name: string;
}

export interface UndeleteIdentityProviderRequest {
  /**
   * The name of the deleted identity provider.
   * Format: idps/{identity_provider}
   */
  name: string;
}

export interface TestIdentityProviderRequest {
  /** The identity provider to test connection including uncreated. */
  identityProvider?: IdentityProvider | undefined;
  oauth2Context?: OAuth2IdentityProviderTestRequestContext | undefined;
}

export interface OAuth2IdentityProviderTestRequestContext {
  /** Authorize code from website. */
  code: string;
}

export interface TestIdentityProviderResponse {
}

export interface IdentityProvider {
  /**
   * The name of the identity provider.
   * Format: idps/{identity_provider}
   */
  name: string;
  /** The system-assigned, unique identifier for a resource. */
  uid: string;
  state: State;
  title: string;
  domain: string;
  type: IdentityProviderType;
  config?: IdentityProviderConfig | undefined;
}

export interface IdentityProviderConfig {
  oauth2Config?: OAuth2IdentityProviderConfig | undefined;
  oidcConfig?: OIDCIdentityProviderConfig | undefined;
  ldapConfig?: LDAPIdentityProviderConfig | undefined;
}

/** OAuth2IdentityProviderConfig is the structure for OAuth2 identity provider config. */
export interface OAuth2IdentityProviderConfig {
  authUrl: string;
  tokenUrl: string;
  userInfoUrl: string;
  clientId: string;
  clientSecret: string;
  scopes: string[];
  fieldMapping?: FieldMapping | undefined;
  skipTlsVerify: boolean;
  authStyle: OAuth2AuthStyle;
}

/** OIDCIdentityProviderConfig is the structure for OIDC identity provider config. */
export interface OIDCIdentityProviderConfig {
  issuer: string;
  clientId: string;
  clientSecret: string;
  scopes: string[];
  fieldMapping?: FieldMapping | undefined;
  skipTlsVerify: boolean;
  authStyle: OAuth2AuthStyle;
}

/** LDAPIdentityProviderConfig is the structure for LDAP identity provider config. */
export interface LDAPIdentityProviderConfig {
  /**
   * Host is the hostname or IP address of the LDAP server, e.g.
   * "ldap.example.com".
   */
  host: string;
  /**
   * Port is the port number of the LDAP server, e.g. 389. When not set, the
   * default port of the corresponding security protocol will be used, i.e. 389
   * for StartTLS and 636 for LDAPS.
   */
  port: number;
  /** SkipTLSVerify controls whether to skip TLS certificate verification. */
  skipTlsVerify: boolean;
  /**
   * BindDN is the DN of the user to bind as a service account to perform
   * search requests.
   */
  bindDn: string;
  /** BindPassword is the password of the user to bind as a service account. */
  bindPassword: string;
  /** BaseDN is the base DN to search for users, e.g. "ou=users,dc=example,dc=com". */
  baseDn: string;
  /** UserFilter is the filter to search for users, e.g. "(uid=%s)". */
  userFilter: string;
  /**
   * SecurityProtocol is the security protocol to be used for establishing
   * connections with the LDAP server. It should be either StartTLS or LDAPS, and
   * cannot be empty.
   */
  securityProtocol: string;
  /**
   * FieldMapping is the mapping of the user attributes returned by the LDAP
   * server.
   */
  fieldMapping?: FieldMapping | undefined;
}

/**
 * FieldMapping saves the field names from user info API of identity provider.
 * As we save all raw json string of user info response data into `principal.idp_user_info`,
 * we can extract the relevant data based with `FieldMapping`.
 *
 * e.g. For GitHub authenticated user API, it will return `login`, `name` and `email` in response.
 * Then the identifier of FieldMapping will be `login`, display_name will be `name`,
 * and email will be `email`.
 * reference: https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
 */
export interface FieldMapping {
  /** Identifier is the field name of the unique identifier in 3rd-party idp user info. Required. */
  identifier: string;
  /** DisplayName is the field name of display name in 3rd-party idp user info. */
  displayName: string;
  /** Email is the field name of primary email in 3rd-party idp user info. */
  email: string;
  /** Phone is the field name of primary phone in 3rd-party idp user info. */
  phone: string;
}

function createBaseGetIdentityProviderRequest(): GetIdentityProviderRequest {
  return { name: "" };
}

export const GetIdentityProviderRequest = {
  encode(message: GetIdentityProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetIdentityProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetIdentityProviderRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: GetIdentityProviderRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create(base?: DeepPartial<GetIdentityProviderRequest>): GetIdentityProviderRequest {
    return GetIdentityProviderRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<GetIdentityProviderRequest>): GetIdentityProviderRequest {
    const message = createBaseGetIdentityProviderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseListIdentityProvidersRequest(): ListIdentityProvidersRequest {
  return { pageSize: 0, pageToken: "", showDeleted: false };
}

export const ListIdentityProvidersRequest = {
  encode(message: ListIdentityProvidersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pageSize !== 0) {
      writer.uint32(8).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(18).string(message.pageToken);
    }
    if (message.showDeleted === true) {
      writer.uint32(24).bool(message.showDeleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListIdentityProvidersRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListIdentityProvidersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pageToken = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.showDeleted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListIdentityProvidersRequest {
    return {
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? String(object.pageToken) : "",
      showDeleted: isSet(object.showDeleted) ? Boolean(object.showDeleted) : false,
    };
  },

  toJSON(message: ListIdentityProvidersRequest): unknown {
    const obj: any = {};
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.pageToken !== undefined && (obj.pageToken = message.pageToken);
    message.showDeleted !== undefined && (obj.showDeleted = message.showDeleted);
    return obj;
  },

  create(base?: DeepPartial<ListIdentityProvidersRequest>): ListIdentityProvidersRequest {
    return ListIdentityProvidersRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListIdentityProvidersRequest>): ListIdentityProvidersRequest {
    const message = createBaseListIdentityProvidersRequest();
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    message.showDeleted = object.showDeleted ?? false;
    return message;
  },
};

function createBaseListIdentityProvidersResponse(): ListIdentityProvidersResponse {
  return { identityProviders: [], nextPageToken: "" };
}

export const ListIdentityProvidersResponse = {
  encode(message: ListIdentityProvidersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.identityProviders) {
      IdentityProvider.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListIdentityProvidersResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListIdentityProvidersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProviders.push(IdentityProvider.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListIdentityProvidersResponse {
    return {
      identityProviders: Array.isArray(object?.identityProviders)
        ? object.identityProviders.map((e: any) => IdentityProvider.fromJSON(e))
        : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: ListIdentityProvidersResponse): unknown {
    const obj: any = {};
    if (message.identityProviders) {
      obj.identityProviders = message.identityProviders.map((e) => e ? IdentityProvider.toJSON(e) : undefined);
    } else {
      obj.identityProviders = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  create(base?: DeepPartial<ListIdentityProvidersResponse>): ListIdentityProvidersResponse {
    return ListIdentityProvidersResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<ListIdentityProvidersResponse>): ListIdentityProvidersResponse {
    const message = createBaseListIdentityProvidersResponse();
    message.identityProviders = object.identityProviders?.map((e) => IdentityProvider.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseCreateIdentityProviderRequest(): CreateIdentityProviderRequest {
  return { identityProvider: undefined, identityProviderId: "" };
}

export const CreateIdentityProviderRequest = {
  encode(message: CreateIdentityProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.identityProvider !== undefined) {
      IdentityProvider.encode(message.identityProvider, writer.uint32(10).fork()).ldelim();
    }
    if (message.identityProviderId !== "") {
      writer.uint32(18).string(message.identityProviderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateIdentityProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProvider = IdentityProvider.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.identityProviderId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateIdentityProviderRequest {
    return {
      identityProvider: isSet(object.identityProvider) ? IdentityProvider.fromJSON(object.identityProvider) : undefined,
      identityProviderId: isSet(object.identityProviderId) ? String(object.identityProviderId) : "",
    };
  },

  toJSON(message: CreateIdentityProviderRequest): unknown {
    const obj: any = {};
    message.identityProvider !== undefined &&
      (obj.identityProvider = message.identityProvider ? IdentityProvider.toJSON(message.identityProvider) : undefined);
    message.identityProviderId !== undefined && (obj.identityProviderId = message.identityProviderId);
    return obj;
  },

  create(base?: DeepPartial<CreateIdentityProviderRequest>): CreateIdentityProviderRequest {
    return CreateIdentityProviderRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<CreateIdentityProviderRequest>): CreateIdentityProviderRequest {
    const message = createBaseCreateIdentityProviderRequest();
    message.identityProvider = (object.identityProvider !== undefined && object.identityProvider !== null)
      ? IdentityProvider.fromPartial(object.identityProvider)
      : undefined;
    message.identityProviderId = object.identityProviderId ?? "";
    return message;
  },
};

function createBaseUpdateIdentityProviderRequest(): UpdateIdentityProviderRequest {
  return { identityProvider: undefined, updateMask: undefined };
}

export const UpdateIdentityProviderRequest = {
  encode(message: UpdateIdentityProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.identityProvider !== undefined) {
      IdentityProvider.encode(message.identityProvider, writer.uint32(10).fork()).ldelim();
    }
    if (message.updateMask !== undefined) {
      FieldMask.encode(FieldMask.wrap(message.updateMask), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateIdentityProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProvider = IdentityProvider.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.updateMask = FieldMask.unwrap(FieldMask.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateIdentityProviderRequest {
    return {
      identityProvider: isSet(object.identityProvider) ? IdentityProvider.fromJSON(object.identityProvider) : undefined,
      updateMask: isSet(object.updateMask) ? FieldMask.unwrap(FieldMask.fromJSON(object.updateMask)) : undefined,
    };
  },

  toJSON(message: UpdateIdentityProviderRequest): unknown {
    const obj: any = {};
    message.identityProvider !== undefined &&
      (obj.identityProvider = message.identityProvider ? IdentityProvider.toJSON(message.identityProvider) : undefined);
    message.updateMask !== undefined && (obj.updateMask = FieldMask.toJSON(FieldMask.wrap(message.updateMask)));
    return obj;
  },

  create(base?: DeepPartial<UpdateIdentityProviderRequest>): UpdateIdentityProviderRequest {
    return UpdateIdentityProviderRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateIdentityProviderRequest>): UpdateIdentityProviderRequest {
    const message = createBaseUpdateIdentityProviderRequest();
    message.identityProvider = (object.identityProvider !== undefined && object.identityProvider !== null)
      ? IdentityProvider.fromPartial(object.identityProvider)
      : undefined;
    message.updateMask = object.updateMask ?? undefined;
    return message;
  },
};

function createBaseDeleteIdentityProviderRequest(): DeleteIdentityProviderRequest {
  return { name: "" };
}

export const DeleteIdentityProviderRequest = {
  encode(message: DeleteIdentityProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteIdentityProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteIdentityProviderRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: DeleteIdentityProviderRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create(base?: DeepPartial<DeleteIdentityProviderRequest>): DeleteIdentityProviderRequest {
    return DeleteIdentityProviderRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<DeleteIdentityProviderRequest>): DeleteIdentityProviderRequest {
    const message = createBaseDeleteIdentityProviderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseUndeleteIdentityProviderRequest(): UndeleteIdentityProviderRequest {
  return { name: "" };
}

export const UndeleteIdentityProviderRequest = {
  encode(message: UndeleteIdentityProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UndeleteIdentityProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUndeleteIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UndeleteIdentityProviderRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: UndeleteIdentityProviderRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create(base?: DeepPartial<UndeleteIdentityProviderRequest>): UndeleteIdentityProviderRequest {
    return UndeleteIdentityProviderRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UndeleteIdentityProviderRequest>): UndeleteIdentityProviderRequest {
    const message = createBaseUndeleteIdentityProviderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseTestIdentityProviderRequest(): TestIdentityProviderRequest {
  return { identityProvider: undefined, oauth2Context: undefined };
}

export const TestIdentityProviderRequest = {
  encode(message: TestIdentityProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.identityProvider !== undefined) {
      IdentityProvider.encode(message.identityProvider, writer.uint32(10).fork()).ldelim();
    }
    if (message.oauth2Context !== undefined) {
      OAuth2IdentityProviderTestRequestContext.encode(message.oauth2Context, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TestIdentityProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTestIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProvider = IdentityProvider.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.oauth2Context = OAuth2IdentityProviderTestRequestContext.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TestIdentityProviderRequest {
    return {
      identityProvider: isSet(object.identityProvider) ? IdentityProvider.fromJSON(object.identityProvider) : undefined,
      oauth2Context: isSet(object.oauth2Context)
        ? OAuth2IdentityProviderTestRequestContext.fromJSON(object.oauth2Context)
        : undefined,
    };
  },

  toJSON(message: TestIdentityProviderRequest): unknown {
    const obj: any = {};
    message.identityProvider !== undefined &&
      (obj.identityProvider = message.identityProvider ? IdentityProvider.toJSON(message.identityProvider) : undefined);
    message.oauth2Context !== undefined && (obj.oauth2Context = message.oauth2Context
      ? OAuth2IdentityProviderTestRequestContext.toJSON(message.oauth2Context)
      : undefined);
    return obj;
  },

  create(base?: DeepPartial<TestIdentityProviderRequest>): TestIdentityProviderRequest {
    return TestIdentityProviderRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<TestIdentityProviderRequest>): TestIdentityProviderRequest {
    const message = createBaseTestIdentityProviderRequest();
    message.identityProvider = (object.identityProvider !== undefined && object.identityProvider !== null)
      ? IdentityProvider.fromPartial(object.identityProvider)
      : undefined;
    message.oauth2Context = (object.oauth2Context !== undefined && object.oauth2Context !== null)
      ? OAuth2IdentityProviderTestRequestContext.fromPartial(object.oauth2Context)
      : undefined;
    return message;
  },
};

function createBaseOAuth2IdentityProviderTestRequestContext(): OAuth2IdentityProviderTestRequestContext {
  return { code: "" };
}

export const OAuth2IdentityProviderTestRequestContext = {
  encode(message: OAuth2IdentityProviderTestRequestContext, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OAuth2IdentityProviderTestRequestContext {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuth2IdentityProviderTestRequestContext();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.code = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OAuth2IdentityProviderTestRequestContext {
    return { code: isSet(object.code) ? String(object.code) : "" };
  },

  toJSON(message: OAuth2IdentityProviderTestRequestContext): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  create(base?: DeepPartial<OAuth2IdentityProviderTestRequestContext>): OAuth2IdentityProviderTestRequestContext {
    return OAuth2IdentityProviderTestRequestContext.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<OAuth2IdentityProviderTestRequestContext>): OAuth2IdentityProviderTestRequestContext {
    const message = createBaseOAuth2IdentityProviderTestRequestContext();
    message.code = object.code ?? "";
    return message;
  },
};

function createBaseTestIdentityProviderResponse(): TestIdentityProviderResponse {
  return {};
}

export const TestIdentityProviderResponse = {
  encode(_: TestIdentityProviderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TestIdentityProviderResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTestIdentityProviderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): TestIdentityProviderResponse {
    return {};
  },

  toJSON(_: TestIdentityProviderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<TestIdentityProviderResponse>): TestIdentityProviderResponse {
    return TestIdentityProviderResponse.fromPartial(base ?? {});
  },

  fromPartial(_: DeepPartial<TestIdentityProviderResponse>): TestIdentityProviderResponse {
    const message = createBaseTestIdentityProviderResponse();
    return message;
  },
};

function createBaseIdentityProvider(): IdentityProvider {
  return { name: "", uid: "", state: 0, title: "", domain: "", type: 0, config: undefined };
}

export const IdentityProvider = {
  encode(message: IdentityProvider, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    if (message.state !== 0) {
      writer.uint32(24).int32(message.state);
    }
    if (message.title !== "") {
      writer.uint32(34).string(message.title);
    }
    if (message.domain !== "") {
      writer.uint32(42).string(message.domain);
    }
    if (message.type !== 0) {
      writer.uint32(48).int32(message.type);
    }
    if (message.config !== undefined) {
      IdentityProviderConfig.encode(message.config, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IdentityProvider {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProvider();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.uid = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.state = reader.int32() as any;
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.title = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.domain = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.config = IdentityProviderConfig.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IdentityProvider {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      uid: isSet(object.uid) ? String(object.uid) : "",
      state: isSet(object.state) ? stateFromJSON(object.state) : 0,
      title: isSet(object.title) ? String(object.title) : "",
      domain: isSet(object.domain) ? String(object.domain) : "",
      type: isSet(object.type) ? identityProviderTypeFromJSON(object.type) : 0,
      config: isSet(object.config) ? IdentityProviderConfig.fromJSON(object.config) : undefined,
    };
  },

  toJSON(message: IdentityProvider): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.uid !== undefined && (obj.uid = message.uid);
    message.state !== undefined && (obj.state = stateToJSON(message.state));
    message.title !== undefined && (obj.title = message.title);
    message.domain !== undefined && (obj.domain = message.domain);
    message.type !== undefined && (obj.type = identityProviderTypeToJSON(message.type));
    message.config !== undefined &&
      (obj.config = message.config ? IdentityProviderConfig.toJSON(message.config) : undefined);
    return obj;
  },

  create(base?: DeepPartial<IdentityProvider>): IdentityProvider {
    return IdentityProvider.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<IdentityProvider>): IdentityProvider {
    const message = createBaseIdentityProvider();
    message.name = object.name ?? "";
    message.uid = object.uid ?? "";
    message.state = object.state ?? 0;
    message.title = object.title ?? "";
    message.domain = object.domain ?? "";
    message.type = object.type ?? 0;
    message.config = (object.config !== undefined && object.config !== null)
      ? IdentityProviderConfig.fromPartial(object.config)
      : undefined;
    return message;
  },
};

function createBaseIdentityProviderConfig(): IdentityProviderConfig {
  return { oauth2Config: undefined, oidcConfig: undefined, ldapConfig: undefined };
}

export const IdentityProviderConfig = {
  encode(message: IdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.oauth2Config !== undefined) {
      OAuth2IdentityProviderConfig.encode(message.oauth2Config, writer.uint32(10).fork()).ldelim();
    }
    if (message.oidcConfig !== undefined) {
      OIDCIdentityProviderConfig.encode(message.oidcConfig, writer.uint32(18).fork()).ldelim();
    }
    if (message.ldapConfig !== undefined) {
      LDAPIdentityProviderConfig.encode(message.ldapConfig, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.oauth2Config = OAuth2IdentityProviderConfig.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.oidcConfig = OIDCIdentityProviderConfig.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.ldapConfig = LDAPIdentityProviderConfig.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IdentityProviderConfig {
    return {
      oauth2Config: isSet(object.oauth2Config) ? OAuth2IdentityProviderConfig.fromJSON(object.oauth2Config) : undefined,
      oidcConfig: isSet(object.oidcConfig) ? OIDCIdentityProviderConfig.fromJSON(object.oidcConfig) : undefined,
      ldapConfig: isSet(object.ldapConfig) ? LDAPIdentityProviderConfig.fromJSON(object.ldapConfig) : undefined,
    };
  },

  toJSON(message: IdentityProviderConfig): unknown {
    const obj: any = {};
    message.oauth2Config !== undefined &&
      (obj.oauth2Config = message.oauth2Config ? OAuth2IdentityProviderConfig.toJSON(message.oauth2Config) : undefined);
    message.oidcConfig !== undefined &&
      (obj.oidcConfig = message.oidcConfig ? OIDCIdentityProviderConfig.toJSON(message.oidcConfig) : undefined);
    message.ldapConfig !== undefined &&
      (obj.ldapConfig = message.ldapConfig ? LDAPIdentityProviderConfig.toJSON(message.ldapConfig) : undefined);
    return obj;
  },

  create(base?: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    return IdentityProviderConfig.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    const message = createBaseIdentityProviderConfig();
    message.oauth2Config = (object.oauth2Config !== undefined && object.oauth2Config !== null)
      ? OAuth2IdentityProviderConfig.fromPartial(object.oauth2Config)
      : undefined;
    message.oidcConfig = (object.oidcConfig !== undefined && object.oidcConfig !== null)
      ? OIDCIdentityProviderConfig.fromPartial(object.oidcConfig)
      : undefined;
    message.ldapConfig = (object.ldapConfig !== undefined && object.ldapConfig !== null)
      ? LDAPIdentityProviderConfig.fromPartial(object.ldapConfig)
      : undefined;
    return message;
  },
};

function createBaseOAuth2IdentityProviderConfig(): OAuth2IdentityProviderConfig {
  return {
    authUrl: "",
    tokenUrl: "",
    userInfoUrl: "",
    clientId: "",
    clientSecret: "",
    scopes: [],
    fieldMapping: undefined,
    skipTlsVerify: false,
    authStyle: 0,
  };
}

export const OAuth2IdentityProviderConfig = {
  encode(message: OAuth2IdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authUrl !== "") {
      writer.uint32(10).string(message.authUrl);
    }
    if (message.tokenUrl !== "") {
      writer.uint32(18).string(message.tokenUrl);
    }
    if (message.userInfoUrl !== "") {
      writer.uint32(26).string(message.userInfoUrl);
    }
    if (message.clientId !== "") {
      writer.uint32(34).string(message.clientId);
    }
    if (message.clientSecret !== "") {
      writer.uint32(42).string(message.clientSecret);
    }
    for (const v of message.scopes) {
      writer.uint32(50).string(v!);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(58).fork()).ldelim();
    }
    if (message.skipTlsVerify === true) {
      writer.uint32(64).bool(message.skipTlsVerify);
    }
    if (message.authStyle !== 0) {
      writer.uint32(72).int32(message.authStyle);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OAuth2IdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuth2IdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.authUrl = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.tokenUrl = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.userInfoUrl = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.clientId = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.clientSecret = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.scopes.push(reader.string());
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.skipTlsVerify = reader.bool();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.authStyle = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OAuth2IdentityProviderConfig {
    return {
      authUrl: isSet(object.authUrl) ? String(object.authUrl) : "",
      tokenUrl: isSet(object.tokenUrl) ? String(object.tokenUrl) : "",
      userInfoUrl: isSet(object.userInfoUrl) ? String(object.userInfoUrl) : "",
      clientId: isSet(object.clientId) ? String(object.clientId) : "",
      clientSecret: isSet(object.clientSecret) ? String(object.clientSecret) : "",
      scopes: Array.isArray(object?.scopes) ? object.scopes.map((e: any) => String(e)) : [],
      fieldMapping: isSet(object.fieldMapping) ? FieldMapping.fromJSON(object.fieldMapping) : undefined,
      skipTlsVerify: isSet(object.skipTlsVerify) ? Boolean(object.skipTlsVerify) : false,
      authStyle: isSet(object.authStyle) ? oAuth2AuthStyleFromJSON(object.authStyle) : 0,
    };
  },

  toJSON(message: OAuth2IdentityProviderConfig): unknown {
    const obj: any = {};
    message.authUrl !== undefined && (obj.authUrl = message.authUrl);
    message.tokenUrl !== undefined && (obj.tokenUrl = message.tokenUrl);
    message.userInfoUrl !== undefined && (obj.userInfoUrl = message.userInfoUrl);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.clientSecret !== undefined && (obj.clientSecret = message.clientSecret);
    if (message.scopes) {
      obj.scopes = message.scopes.map((e) => e);
    } else {
      obj.scopes = [];
    }
    message.fieldMapping !== undefined &&
      (obj.fieldMapping = message.fieldMapping ? FieldMapping.toJSON(message.fieldMapping) : undefined);
    message.skipTlsVerify !== undefined && (obj.skipTlsVerify = message.skipTlsVerify);
    message.authStyle !== undefined && (obj.authStyle = oAuth2AuthStyleToJSON(message.authStyle));
    return obj;
  },

  create(base?: DeepPartial<OAuth2IdentityProviderConfig>): OAuth2IdentityProviderConfig {
    return OAuth2IdentityProviderConfig.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<OAuth2IdentityProviderConfig>): OAuth2IdentityProviderConfig {
    const message = createBaseOAuth2IdentityProviderConfig();
    message.authUrl = object.authUrl ?? "";
    message.tokenUrl = object.tokenUrl ?? "";
    message.userInfoUrl = object.userInfoUrl ?? "";
    message.clientId = object.clientId ?? "";
    message.clientSecret = object.clientSecret ?? "";
    message.scopes = object.scopes?.map((e) => e) || [];
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
    message.skipTlsVerify = object.skipTlsVerify ?? false;
    message.authStyle = object.authStyle ?? 0;
    return message;
  },
};

function createBaseOIDCIdentityProviderConfig(): OIDCIdentityProviderConfig {
  return {
    issuer: "",
    clientId: "",
    clientSecret: "",
    scopes: [],
    fieldMapping: undefined,
    skipTlsVerify: false,
    authStyle: 0,
  };
}

export const OIDCIdentityProviderConfig = {
  encode(message: OIDCIdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.clientId !== "") {
      writer.uint32(18).string(message.clientId);
    }
    if (message.clientSecret !== "") {
      writer.uint32(26).string(message.clientSecret);
    }
    for (const v of message.scopes) {
      writer.uint32(34).string(v!);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(42).fork()).ldelim();
    }
    if (message.skipTlsVerify === true) {
      writer.uint32(48).bool(message.skipTlsVerify);
    }
    if (message.authStyle !== 0) {
      writer.uint32(56).int32(message.authStyle);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OIDCIdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOIDCIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.issuer = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.clientId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.clientSecret = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.scopes.push(reader.string());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.skipTlsVerify = reader.bool();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.authStyle = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OIDCIdentityProviderConfig {
    return {
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      clientId: isSet(object.clientId) ? String(object.clientId) : "",
      clientSecret: isSet(object.clientSecret) ? String(object.clientSecret) : "",
      scopes: Array.isArray(object?.scopes) ? object.scopes.map((e: any) => String(e)) : [],
      fieldMapping: isSet(object.fieldMapping) ? FieldMapping.fromJSON(object.fieldMapping) : undefined,
      skipTlsVerify: isSet(object.skipTlsVerify) ? Boolean(object.skipTlsVerify) : false,
      authStyle: isSet(object.authStyle) ? oAuth2AuthStyleFromJSON(object.authStyle) : 0,
    };
  },

  toJSON(message: OIDCIdentityProviderConfig): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.clientSecret !== undefined && (obj.clientSecret = message.clientSecret);
    if (message.scopes) {
      obj.scopes = message.scopes.map((e) => e);
    } else {
      obj.scopes = [];
    }
    message.fieldMapping !== undefined &&
      (obj.fieldMapping = message.fieldMapping ? FieldMapping.toJSON(message.fieldMapping) : undefined);
    message.skipTlsVerify !== undefined && (obj.skipTlsVerify = message.skipTlsVerify);
    message.authStyle !== undefined && (obj.authStyle = oAuth2AuthStyleToJSON(message.authStyle));
    return obj;
  },

  create(base?: DeepPartial<OIDCIdentityProviderConfig>): OIDCIdentityProviderConfig {
    return OIDCIdentityProviderConfig.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<OIDCIdentityProviderConfig>): OIDCIdentityProviderConfig {
    const message = createBaseOIDCIdentityProviderConfig();
    message.issuer = object.issuer ?? "";
    message.clientId = object.clientId ?? "";
    message.clientSecret = object.clientSecret ?? "";
    message.scopes = object.scopes?.map((e) => e) || [];
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
    message.skipTlsVerify = object.skipTlsVerify ?? false;
    message.authStyle = object.authStyle ?? 0;
    return message;
  },
};

function createBaseLDAPIdentityProviderConfig(): LDAPIdentityProviderConfig {
  return {
    host: "",
    port: 0,
    skipTlsVerify: false,
    bindDn: "",
    bindPassword: "",
    baseDn: "",
    userFilter: "",
    securityProtocol: "",
    fieldMapping: undefined,
  };
}

export const LDAPIdentityProviderConfig = {
  encode(message: LDAPIdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.host !== "") {
      writer.uint32(10).string(message.host);
    }
    if (message.port !== 0) {
      writer.uint32(16).int32(message.port);
    }
    if (message.skipTlsVerify === true) {
      writer.uint32(24).bool(message.skipTlsVerify);
    }
    if (message.bindDn !== "") {
      writer.uint32(34).string(message.bindDn);
    }
    if (message.bindPassword !== "") {
      writer.uint32(42).string(message.bindPassword);
    }
    if (message.baseDn !== "") {
      writer.uint32(50).string(message.baseDn);
    }
    if (message.userFilter !== "") {
      writer.uint32(58).string(message.userFilter);
    }
    if (message.securityProtocol !== "") {
      writer.uint32(66).string(message.securityProtocol);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LDAPIdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLDAPIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.host = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.port = reader.int32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.skipTlsVerify = reader.bool();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.bindDn = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.bindPassword = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.baseDn = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.userFilter = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.securityProtocol = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LDAPIdentityProviderConfig {
    return {
      host: isSet(object.host) ? String(object.host) : "",
      port: isSet(object.port) ? Number(object.port) : 0,
      skipTlsVerify: isSet(object.skipTlsVerify) ? Boolean(object.skipTlsVerify) : false,
      bindDn: isSet(object.bindDn) ? String(object.bindDn) : "",
      bindPassword: isSet(object.bindPassword) ? String(object.bindPassword) : "",
      baseDn: isSet(object.baseDn) ? String(object.baseDn) : "",
      userFilter: isSet(object.userFilter) ? String(object.userFilter) : "",
      securityProtocol: isSet(object.securityProtocol) ? String(object.securityProtocol) : "",
      fieldMapping: isSet(object.fieldMapping) ? FieldMapping.fromJSON(object.fieldMapping) : undefined,
    };
  },

  toJSON(message: LDAPIdentityProviderConfig): unknown {
    const obj: any = {};
    message.host !== undefined && (obj.host = message.host);
    message.port !== undefined && (obj.port = Math.round(message.port));
    message.skipTlsVerify !== undefined && (obj.skipTlsVerify = message.skipTlsVerify);
    message.bindDn !== undefined && (obj.bindDn = message.bindDn);
    message.bindPassword !== undefined && (obj.bindPassword = message.bindPassword);
    message.baseDn !== undefined && (obj.baseDn = message.baseDn);
    message.userFilter !== undefined && (obj.userFilter = message.userFilter);
    message.securityProtocol !== undefined && (obj.securityProtocol = message.securityProtocol);
    message.fieldMapping !== undefined &&
      (obj.fieldMapping = message.fieldMapping ? FieldMapping.toJSON(message.fieldMapping) : undefined);
    return obj;
  },

  create(base?: DeepPartial<LDAPIdentityProviderConfig>): LDAPIdentityProviderConfig {
    return LDAPIdentityProviderConfig.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<LDAPIdentityProviderConfig>): LDAPIdentityProviderConfig {
    const message = createBaseLDAPIdentityProviderConfig();
    message.host = object.host ?? "";
    message.port = object.port ?? 0;
    message.skipTlsVerify = object.skipTlsVerify ?? false;
    message.bindDn = object.bindDn ?? "";
    message.bindPassword = object.bindPassword ?? "";
    message.baseDn = object.baseDn ?? "";
    message.userFilter = object.userFilter ?? "";
    message.securityProtocol = object.securityProtocol ?? "";
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
    return message;
  },
};

function createBaseFieldMapping(): FieldMapping {
  return { identifier: "", displayName: "", email: "", phone: "" };
}

export const FieldMapping = {
  encode(message: FieldMapping, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.identifier !== "") {
      writer.uint32(10).string(message.identifier);
    }
    if (message.displayName !== "") {
      writer.uint32(18).string(message.displayName);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    if (message.phone !== "") {
      writer.uint32(34).string(message.phone);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FieldMapping {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFieldMapping();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identifier = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.displayName = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.email = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.phone = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FieldMapping {
    return {
      identifier: isSet(object.identifier) ? String(object.identifier) : "",
      displayName: isSet(object.displayName) ? String(object.displayName) : "",
      email: isSet(object.email) ? String(object.email) : "",
      phone: isSet(object.phone) ? String(object.phone) : "",
    };
  },

  toJSON(message: FieldMapping): unknown {
    const obj: any = {};
    message.identifier !== undefined && (obj.identifier = message.identifier);
    message.displayName !== undefined && (obj.displayName = message.displayName);
    message.email !== undefined && (obj.email = message.email);
    message.phone !== undefined && (obj.phone = message.phone);
    return obj;
  },

  create(base?: DeepPartial<FieldMapping>): FieldMapping {
    return FieldMapping.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<FieldMapping>): FieldMapping {
    const message = createBaseFieldMapping();
    message.identifier = object.identifier ?? "";
    message.displayName = object.displayName ?? "";
    message.email = object.email ?? "";
    message.phone = object.phone ?? "";
    return message;
  },
};

export type IdentityProviderServiceDefinition = typeof IdentityProviderServiceDefinition;
export const IdentityProviderServiceDefinition = {
  name: "IdentityProviderService",
  fullName: "bytebase.v1.IdentityProviderService",
  methods: {
    getIdentityProvider: {
      name: "GetIdentityProvider",
      requestType: GetIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([19, 18, 17, 47, 118, 49, 47, 123, 110, 97, 109, 101, 61, 105, 100, 112, 115, 47, 42, 125]),
          ],
        },
      },
    },
    listIdentityProviders: {
      name: "ListIdentityProviders",
      requestType: ListIdentityProvidersRequest,
      requestStream: false,
      responseType: ListIdentityProvidersResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([0])],
          578365826: [new Uint8Array([10, 18, 8, 47, 118, 49, 47, 105, 100, 112, 115])],
        },
      },
    },
    createIdentityProvider: {
      name: "CreateIdentityProvider",
      requestType: CreateIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([0])],
          578365826: [
            new Uint8Array([
              29,
              58,
              17,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              34,
              8,
              47,
              118,
              49,
              47,
              105,
              100,
              112,
              115,
            ]),
          ],
        },
      },
    },
    updateIdentityProvider: {
      name: "UpdateIdentityProvider",
      requestType: UpdateIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [
            new Uint8Array([
              29,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              44,
              117,
              112,
              100,
              97,
              116,
              101,
              95,
              109,
              97,
              115,
              107,
            ]),
          ],
          578365826: [
            new Uint8Array([
              56,
              58,
              17,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              50,
              35,
              47,
              118,
              49,
              47,
              123,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              46,
              110,
              97,
              109,
              101,
              61,
              105,
              100,
              112,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
    deleteIdentityProvider: {
      name: "DeleteIdentityProvider",
      requestType: DeleteIdentityProviderRequest,
      requestStream: false,
      responseType: Empty,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([19, 42, 17, 47, 118, 49, 47, 123, 110, 97, 109, 101, 61, 105, 100, 112, 115, 47, 42, 125]),
          ],
        },
      },
    },
    undeleteIdentityProvider: {
      name: "UndeleteIdentityProvider",
      requestType: UndeleteIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              31,
              58,
              1,
              42,
              34,
              26,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              105,
              100,
              112,
              115,
              47,
              42,
              125,
              58,
              117,
              110,
              100,
              101,
              108,
              101,
              116,
              101,
            ]),
          ],
        },
      },
    },
    testIdentityProvider: {
      name: "TestIdentityProvider",
      requestType: TestIdentityProviderRequest,
      requestStream: false,
      responseType: TestIdentityProviderResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              20,
              58,
              1,
              42,
              34,
              15,
              47,
              118,
              49,
              47,
              105,
              100,
              112,
              115,
              47,
              42,
              58,
              116,
              101,
              115,
              116,
            ]),
          ],
        },
      },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

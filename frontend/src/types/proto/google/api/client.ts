/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Duration } from "../protobuf/duration";
import { LaunchStage, launchStageFromJSON, launchStageToJSON, launchStageToNumber } from "./launch_stage";

export const protobufPackage = "google.api";

/**
 * The organization for which the client libraries are being published.
 * Affects the url where generated docs are published, etc.
 */
export enum ClientLibraryOrganization {
  /** CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED - Not useful. */
  CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED = "CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED",
  /** CLOUD - Google Cloud Platform Org. */
  CLOUD = "CLOUD",
  /** ADS - Ads (Advertising) Org. */
  ADS = "ADS",
  /** PHOTOS - Photos Org. */
  PHOTOS = "PHOTOS",
  /** STREET_VIEW - Street View Org. */
  STREET_VIEW = "STREET_VIEW",
  /** SHOPPING - Shopping Org. */
  SHOPPING = "SHOPPING",
  /** GEO - Geo Org. */
  GEO = "GEO",
  /** GENERATIVE_AI - Generative AI - https://developers.generativeai.google */
  GENERATIVE_AI = "GENERATIVE_AI",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function clientLibraryOrganizationFromJSON(object: any): ClientLibraryOrganization {
  switch (object) {
    case 0:
    case "CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED":
      return ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED;
    case 1:
    case "CLOUD":
      return ClientLibraryOrganization.CLOUD;
    case 2:
    case "ADS":
      return ClientLibraryOrganization.ADS;
    case 3:
    case "PHOTOS":
      return ClientLibraryOrganization.PHOTOS;
    case 4:
    case "STREET_VIEW":
      return ClientLibraryOrganization.STREET_VIEW;
    case 5:
    case "SHOPPING":
      return ClientLibraryOrganization.SHOPPING;
    case 6:
    case "GEO":
      return ClientLibraryOrganization.GEO;
    case 7:
    case "GENERATIVE_AI":
      return ClientLibraryOrganization.GENERATIVE_AI;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ClientLibraryOrganization.UNRECOGNIZED;
  }
}

export function clientLibraryOrganizationToJSON(object: ClientLibraryOrganization): string {
  switch (object) {
    case ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED:
      return "CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED";
    case ClientLibraryOrganization.CLOUD:
      return "CLOUD";
    case ClientLibraryOrganization.ADS:
      return "ADS";
    case ClientLibraryOrganization.PHOTOS:
      return "PHOTOS";
    case ClientLibraryOrganization.STREET_VIEW:
      return "STREET_VIEW";
    case ClientLibraryOrganization.SHOPPING:
      return "SHOPPING";
    case ClientLibraryOrganization.GEO:
      return "GEO";
    case ClientLibraryOrganization.GENERATIVE_AI:
      return "GENERATIVE_AI";
    case ClientLibraryOrganization.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function clientLibraryOrganizationToNumber(object: ClientLibraryOrganization): number {
  switch (object) {
    case ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED:
      return 0;
    case ClientLibraryOrganization.CLOUD:
      return 1;
    case ClientLibraryOrganization.ADS:
      return 2;
    case ClientLibraryOrganization.PHOTOS:
      return 3;
    case ClientLibraryOrganization.STREET_VIEW:
      return 4;
    case ClientLibraryOrganization.SHOPPING:
      return 5;
    case ClientLibraryOrganization.GEO:
      return 6;
    case ClientLibraryOrganization.GENERATIVE_AI:
      return 7;
    case ClientLibraryOrganization.UNRECOGNIZED:
    default:
      return -1;
  }
}

/** To where should client libraries be published? */
export enum ClientLibraryDestination {
  /**
   * CLIENT_LIBRARY_DESTINATION_UNSPECIFIED - Client libraries will neither be generated nor published to package
   * managers.
   */
  CLIENT_LIBRARY_DESTINATION_UNSPECIFIED = "CLIENT_LIBRARY_DESTINATION_UNSPECIFIED",
  /**
   * GITHUB - Generate the client library in a repo under github.com/googleapis,
   * but don't publish it to package managers.
   */
  GITHUB = "GITHUB",
  /** PACKAGE_MANAGER - Publish the library to package managers like nuget.org and npmjs.com. */
  PACKAGE_MANAGER = "PACKAGE_MANAGER",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function clientLibraryDestinationFromJSON(object: any): ClientLibraryDestination {
  switch (object) {
    case 0:
    case "CLIENT_LIBRARY_DESTINATION_UNSPECIFIED":
      return ClientLibraryDestination.CLIENT_LIBRARY_DESTINATION_UNSPECIFIED;
    case 10:
    case "GITHUB":
      return ClientLibraryDestination.GITHUB;
    case 20:
    case "PACKAGE_MANAGER":
      return ClientLibraryDestination.PACKAGE_MANAGER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ClientLibraryDestination.UNRECOGNIZED;
  }
}

export function clientLibraryDestinationToJSON(object: ClientLibraryDestination): string {
  switch (object) {
    case ClientLibraryDestination.CLIENT_LIBRARY_DESTINATION_UNSPECIFIED:
      return "CLIENT_LIBRARY_DESTINATION_UNSPECIFIED";
    case ClientLibraryDestination.GITHUB:
      return "GITHUB";
    case ClientLibraryDestination.PACKAGE_MANAGER:
      return "PACKAGE_MANAGER";
    case ClientLibraryDestination.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function clientLibraryDestinationToNumber(object: ClientLibraryDestination): number {
  switch (object) {
    case ClientLibraryDestination.CLIENT_LIBRARY_DESTINATION_UNSPECIFIED:
      return 0;
    case ClientLibraryDestination.GITHUB:
      return 10;
    case ClientLibraryDestination.PACKAGE_MANAGER:
      return 20;
    case ClientLibraryDestination.UNRECOGNIZED:
    default:
      return -1;
  }
}

/** Required information for every language. */
export interface CommonLanguageSettings {
  /**
   * Link to automatically generated reference documentation.  Example:
   * https://cloud.google.com/nodejs/docs/reference/asset/latest
   *
   * @deprecated
   */
  referenceDocsUri: string;
  /** The destination where API teams want this client library to be published. */
  destinations: ClientLibraryDestination[];
}

/** Details about how and where to publish client libraries. */
export interface ClientLibrarySettings {
  /**
   * Version of the API to apply these settings to. This is the full protobuf
   * package for the API, ending in the version element.
   * Examples: "google.cloud.speech.v1" and "google.spanner.admin.database.v1".
   */
  version: string;
  /** Launch stage of this version of the API. */
  launchStage: LaunchStage;
  /**
   * When using transport=rest, the client request will encode enums as
   * numbers rather than strings.
   */
  restNumericEnums: boolean;
  /** Settings for legacy Java features, supported in the Service YAML. */
  javaSettings:
    | JavaSettings
    | undefined;
  /** Settings for C++ client libraries. */
  cppSettings:
    | CppSettings
    | undefined;
  /** Settings for PHP client libraries. */
  phpSettings:
    | PhpSettings
    | undefined;
  /** Settings for Python client libraries. */
  pythonSettings:
    | PythonSettings
    | undefined;
  /** Settings for Node client libraries. */
  nodeSettings:
    | NodeSettings
    | undefined;
  /** Settings for .NET client libraries. */
  dotnetSettings:
    | DotnetSettings
    | undefined;
  /** Settings for Ruby client libraries. */
  rubySettings:
    | RubySettings
    | undefined;
  /** Settings for Go client libraries. */
  goSettings: GoSettings | undefined;
}

/**
 * This message configures the settings for publishing [Google Cloud Client
 * libraries](https://cloud.google.com/apis/docs/cloud-client-libraries)
 * generated from the service config.
 */
export interface Publishing {
  /**
   * A list of API method settings, e.g. the behavior for methods that use the
   * long-running operation pattern.
   */
  methodSettings: MethodSettings[];
  /**
   * Link to a *public* URI where users can report issues.  Example:
   * https://issuetracker.google.com/issues/new?component=190865&template=1161103
   */
  newIssueUri: string;
  /**
   * Link to product home page.  Example:
   * https://cloud.google.com/asset-inventory/docs/overview
   */
  documentationUri: string;
  /**
   * Used as a tracking tag when collecting data about the APIs developer
   * relations artifacts like docs, packages delivered to package managers,
   * etc.  Example: "speech".
   */
  apiShortName: string;
  /** GitHub label to apply to issues and pull requests opened for this API. */
  githubLabel: string;
  /**
   * GitHub teams to be added to CODEOWNERS in the directory in GitHub
   * containing source code for the client libraries for this API.
   */
  codeownerGithubTeams: string[];
  /**
   * A prefix used in sample code when demarking regions to be included in
   * documentation.
   */
  docTagPrefix: string;
  /** For whom the client library is being published. */
  organization: ClientLibraryOrganization;
  /**
   * Client library settings.  If the same version string appears multiple
   * times in this list, then the last one wins.  Settings from earlier
   * settings with the same version string are discarded.
   */
  librarySettings: ClientLibrarySettings[];
  /**
   * Optional link to proto reference documentation.  Example:
   * https://cloud.google.com/pubsub/lite/docs/reference/rpc
   */
  protoReferenceDocumentationUri: string;
}

/** Settings for Java client libraries. */
export interface JavaSettings {
  /**
   * The package name to use in Java. Clobbers the java_package option
   * set in the protobuf. This should be used **only** by APIs
   * who have already set the language_settings.java.package_name" field
   * in gapic.yaml. API teams should use the protobuf java_package option
   * where possible.
   *
   * Example of a YAML configuration::
   *
   *  publishing:
   *    java_settings:
   *      library_package: com.google.cloud.pubsub.v1
   */
  libraryPackage: string;
  /**
   * Configure the Java class name to use instead of the service's for its
   * corresponding generated GAPIC client. Keys are fully-qualified
   * service names as they appear in the protobuf (including the full
   * the language_settings.java.interface_names" field in gapic.yaml. API
   * teams should otherwise use the service name as it appears in the
   * protobuf.
   *
   * Example of a YAML configuration::
   *
   *  publishing:
   *    java_settings:
   *      service_class_names:
   *        - google.pubsub.v1.Publisher: TopicAdmin
   *        - google.pubsub.v1.Subscriber: SubscriptionAdmin
   */
  serviceClassNames: { [key: string]: string };
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

export interface JavaSettings_ServiceClassNamesEntry {
  key: string;
  value: string;
}

/** Settings for C++ client libraries. */
export interface CppSettings {
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

/** Settings for Php client libraries. */
export interface PhpSettings {
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

/** Settings for Python client libraries. */
export interface PythonSettings {
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

/** Settings for Node client libraries. */
export interface NodeSettings {
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

/** Settings for Dotnet client libraries. */
export interface DotnetSettings {
  /** Some settings. */
  common:
    | CommonLanguageSettings
    | undefined;
  /**
   * Map from original service names to renamed versions.
   * This is used when the default generated types
   * would cause a naming conflict. (Neither name is
   * fully-qualified.)
   * Example: Subscriber to SubscriberServiceApi.
   */
  renamedServices: { [key: string]: string };
  /**
   * Map from full resource types to the effective short name
   * for the resource. This is used when otherwise resource
   * named from different services would cause naming collisions.
   * Example entry:
   * "datalabeling.googleapis.com/Dataset": "DataLabelingDataset"
   */
  renamedResources: { [key: string]: string };
  /**
   * List of full resource types to ignore during generation.
   * This is typically used for API-specific Location resources,
   * which should be handled by the generator as if they were actually
   * the common Location resources.
   * Example entry: "documentai.googleapis.com/Location"
   */
  ignoredResources: string[];
  /**
   * Namespaces which must be aliased in snippets due to
   * a known (but non-generator-predictable) naming collision
   */
  forcedNamespaceAliases: string[];
  /**
   * Method signatures (in the form "service.method(signature)")
   * which are provided separately, so shouldn't be generated.
   * Snippets *calling* these methods are still generated, however.
   */
  handwrittenSignatures: string[];
}

export interface DotnetSettings_RenamedServicesEntry {
  key: string;
  value: string;
}

export interface DotnetSettings_RenamedResourcesEntry {
  key: string;
  value: string;
}

/** Settings for Ruby client libraries. */
export interface RubySettings {
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

/** Settings for Go client libraries. */
export interface GoSettings {
  /** Some settings. */
  common: CommonLanguageSettings | undefined;
}

/** Describes the generator configuration for a method. */
export interface MethodSettings {
  /**
   * The fully qualified name of the method, for which the options below apply.
   * This is used to find the method to apply the options.
   */
  selector: string;
  /**
   * Describes settings to use for long-running operations when generating
   * API methods for RPCs. Complements RPCs that use the annotations in
   * google/longrunning/operations.proto.
   *
   * Example of a YAML configuration::
   *
   *  publishing:
   *    method_settings:
   *      - selector: google.cloud.speech.v2.Speech.BatchRecognize
   *        long_running:
   *          initial_poll_delay:
   *            seconds: 60 # 1 minute
   *          poll_delay_multiplier: 1.5
   *          max_poll_delay:
   *            seconds: 360 # 6 minutes
   *          total_poll_timeout:
   *             seconds: 54000 # 90 minutes
   */
  longRunning: MethodSettings_LongRunning | undefined;
}

/**
 * Describes settings to use when generating API methods that use the
 * long-running operation pattern.
 * All default values below are from those used in the client library
 * generators (e.g.
 * [Java](https://github.com/googleapis/gapic-generator-java/blob/04c2faa191a9b5a10b92392fe8482279c4404803/src/main/java/com/google/api/generator/gapic/composer/common/RetrySettingsComposer.java)).
 */
export interface MethodSettings_LongRunning {
  /**
   * Initial delay after which the first poll request will be made.
   * Default value: 5 seconds.
   */
  initialPollDelay:
    | Duration
    | undefined;
  /**
   * Multiplier to gradually increase delay between subsequent polls until it
   * reaches max_poll_delay.
   * Default value: 1.5.
   */
  pollDelayMultiplier: number;
  /**
   * Maximum time between two subsequent poll requests.
   * Default value: 45 seconds.
   */
  maxPollDelay:
    | Duration
    | undefined;
  /**
   * Total polling timeout.
   * Default value: 5 minutes.
   */
  totalPollTimeout: Duration | undefined;
}

function createBaseCommonLanguageSettings(): CommonLanguageSettings {
  return { referenceDocsUri: "", destinations: [] };
}

export const CommonLanguageSettings = {
  encode(message: CommonLanguageSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceDocsUri !== "") {
      writer.uint32(10).string(message.referenceDocsUri);
    }
    writer.uint32(18).fork();
    for (const v of message.destinations) {
      writer.int32(clientLibraryDestinationToNumber(v));
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CommonLanguageSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommonLanguageSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.referenceDocsUri = reader.string();
          continue;
        case 2:
          if (tag === 16) {
            message.destinations.push(clientLibraryDestinationFromJSON(reader.int32()));

            continue;
          }

          if (tag === 18) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.destinations.push(clientLibraryDestinationFromJSON(reader.int32()));
            }

            continue;
          }

          break;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CommonLanguageSettings {
    return {
      referenceDocsUri: isSet(object.referenceDocsUri) ? globalThis.String(object.referenceDocsUri) : "",
      destinations: globalThis.Array.isArray(object?.destinations)
        ? object.destinations.map((e: any) => clientLibraryDestinationFromJSON(e))
        : [],
    };
  },

  toJSON(message: CommonLanguageSettings): unknown {
    const obj: any = {};
    if (message.referenceDocsUri !== "") {
      obj.referenceDocsUri = message.referenceDocsUri;
    }
    if (message.destinations?.length) {
      obj.destinations = message.destinations.map((e) => clientLibraryDestinationToJSON(e));
    }
    return obj;
  },

  create(base?: DeepPartial<CommonLanguageSettings>): CommonLanguageSettings {
    return CommonLanguageSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CommonLanguageSettings>): CommonLanguageSettings {
    const message = createBaseCommonLanguageSettings();
    message.referenceDocsUri = object.referenceDocsUri ?? "";
    message.destinations = object.destinations?.map((e) => e) || [];
    return message;
  },
};

function createBaseClientLibrarySettings(): ClientLibrarySettings {
  return {
    version: "",
    launchStage: LaunchStage.LAUNCH_STAGE_UNSPECIFIED,
    restNumericEnums: false,
    javaSettings: undefined,
    cppSettings: undefined,
    phpSettings: undefined,
    pythonSettings: undefined,
    nodeSettings: undefined,
    dotnetSettings: undefined,
    rubySettings: undefined,
    goSettings: undefined,
  };
}

export const ClientLibrarySettings = {
  encode(message: ClientLibrarySettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.version !== "") {
      writer.uint32(10).string(message.version);
    }
    if (message.launchStage !== LaunchStage.LAUNCH_STAGE_UNSPECIFIED) {
      writer.uint32(16).int32(launchStageToNumber(message.launchStage));
    }
    if (message.restNumericEnums === true) {
      writer.uint32(24).bool(message.restNumericEnums);
    }
    if (message.javaSettings !== undefined) {
      JavaSettings.encode(message.javaSettings, writer.uint32(170).fork()).ldelim();
    }
    if (message.cppSettings !== undefined) {
      CppSettings.encode(message.cppSettings, writer.uint32(178).fork()).ldelim();
    }
    if (message.phpSettings !== undefined) {
      PhpSettings.encode(message.phpSettings, writer.uint32(186).fork()).ldelim();
    }
    if (message.pythonSettings !== undefined) {
      PythonSettings.encode(message.pythonSettings, writer.uint32(194).fork()).ldelim();
    }
    if (message.nodeSettings !== undefined) {
      NodeSettings.encode(message.nodeSettings, writer.uint32(202).fork()).ldelim();
    }
    if (message.dotnetSettings !== undefined) {
      DotnetSettings.encode(message.dotnetSettings, writer.uint32(210).fork()).ldelim();
    }
    if (message.rubySettings !== undefined) {
      RubySettings.encode(message.rubySettings, writer.uint32(218).fork()).ldelim();
    }
    if (message.goSettings !== undefined) {
      GoSettings.encode(message.goSettings, writer.uint32(226).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClientLibrarySettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClientLibrarySettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.version = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.launchStage = launchStageFromJSON(reader.int32());
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.restNumericEnums = reader.bool();
          continue;
        case 21:
          if (tag !== 170) {
            break;
          }

          message.javaSettings = JavaSettings.decode(reader, reader.uint32());
          continue;
        case 22:
          if (tag !== 178) {
            break;
          }

          message.cppSettings = CppSettings.decode(reader, reader.uint32());
          continue;
        case 23:
          if (tag !== 186) {
            break;
          }

          message.phpSettings = PhpSettings.decode(reader, reader.uint32());
          continue;
        case 24:
          if (tag !== 194) {
            break;
          }

          message.pythonSettings = PythonSettings.decode(reader, reader.uint32());
          continue;
        case 25:
          if (tag !== 202) {
            break;
          }

          message.nodeSettings = NodeSettings.decode(reader, reader.uint32());
          continue;
        case 26:
          if (tag !== 210) {
            break;
          }

          message.dotnetSettings = DotnetSettings.decode(reader, reader.uint32());
          continue;
        case 27:
          if (tag !== 218) {
            break;
          }

          message.rubySettings = RubySettings.decode(reader, reader.uint32());
          continue;
        case 28:
          if (tag !== 226) {
            break;
          }

          message.goSettings = GoSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ClientLibrarySettings {
    return {
      version: isSet(object.version) ? globalThis.String(object.version) : "",
      launchStage: isSet(object.launchStage)
        ? launchStageFromJSON(object.launchStage)
        : LaunchStage.LAUNCH_STAGE_UNSPECIFIED,
      restNumericEnums: isSet(object.restNumericEnums) ? globalThis.Boolean(object.restNumericEnums) : false,
      javaSettings: isSet(object.javaSettings) ? JavaSettings.fromJSON(object.javaSettings) : undefined,
      cppSettings: isSet(object.cppSettings) ? CppSettings.fromJSON(object.cppSettings) : undefined,
      phpSettings: isSet(object.phpSettings) ? PhpSettings.fromJSON(object.phpSettings) : undefined,
      pythonSettings: isSet(object.pythonSettings) ? PythonSettings.fromJSON(object.pythonSettings) : undefined,
      nodeSettings: isSet(object.nodeSettings) ? NodeSettings.fromJSON(object.nodeSettings) : undefined,
      dotnetSettings: isSet(object.dotnetSettings) ? DotnetSettings.fromJSON(object.dotnetSettings) : undefined,
      rubySettings: isSet(object.rubySettings) ? RubySettings.fromJSON(object.rubySettings) : undefined,
      goSettings: isSet(object.goSettings) ? GoSettings.fromJSON(object.goSettings) : undefined,
    };
  },

  toJSON(message: ClientLibrarySettings): unknown {
    const obj: any = {};
    if (message.version !== "") {
      obj.version = message.version;
    }
    if (message.launchStage !== LaunchStage.LAUNCH_STAGE_UNSPECIFIED) {
      obj.launchStage = launchStageToJSON(message.launchStage);
    }
    if (message.restNumericEnums === true) {
      obj.restNumericEnums = message.restNumericEnums;
    }
    if (message.javaSettings !== undefined) {
      obj.javaSettings = JavaSettings.toJSON(message.javaSettings);
    }
    if (message.cppSettings !== undefined) {
      obj.cppSettings = CppSettings.toJSON(message.cppSettings);
    }
    if (message.phpSettings !== undefined) {
      obj.phpSettings = PhpSettings.toJSON(message.phpSettings);
    }
    if (message.pythonSettings !== undefined) {
      obj.pythonSettings = PythonSettings.toJSON(message.pythonSettings);
    }
    if (message.nodeSettings !== undefined) {
      obj.nodeSettings = NodeSettings.toJSON(message.nodeSettings);
    }
    if (message.dotnetSettings !== undefined) {
      obj.dotnetSettings = DotnetSettings.toJSON(message.dotnetSettings);
    }
    if (message.rubySettings !== undefined) {
      obj.rubySettings = RubySettings.toJSON(message.rubySettings);
    }
    if (message.goSettings !== undefined) {
      obj.goSettings = GoSettings.toJSON(message.goSettings);
    }
    return obj;
  },

  create(base?: DeepPartial<ClientLibrarySettings>): ClientLibrarySettings {
    return ClientLibrarySettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ClientLibrarySettings>): ClientLibrarySettings {
    const message = createBaseClientLibrarySettings();
    message.version = object.version ?? "";
    message.launchStage = object.launchStage ?? LaunchStage.LAUNCH_STAGE_UNSPECIFIED;
    message.restNumericEnums = object.restNumericEnums ?? false;
    message.javaSettings = (object.javaSettings !== undefined && object.javaSettings !== null)
      ? JavaSettings.fromPartial(object.javaSettings)
      : undefined;
    message.cppSettings = (object.cppSettings !== undefined && object.cppSettings !== null)
      ? CppSettings.fromPartial(object.cppSettings)
      : undefined;
    message.phpSettings = (object.phpSettings !== undefined && object.phpSettings !== null)
      ? PhpSettings.fromPartial(object.phpSettings)
      : undefined;
    message.pythonSettings = (object.pythonSettings !== undefined && object.pythonSettings !== null)
      ? PythonSettings.fromPartial(object.pythonSettings)
      : undefined;
    message.nodeSettings = (object.nodeSettings !== undefined && object.nodeSettings !== null)
      ? NodeSettings.fromPartial(object.nodeSettings)
      : undefined;
    message.dotnetSettings = (object.dotnetSettings !== undefined && object.dotnetSettings !== null)
      ? DotnetSettings.fromPartial(object.dotnetSettings)
      : undefined;
    message.rubySettings = (object.rubySettings !== undefined && object.rubySettings !== null)
      ? RubySettings.fromPartial(object.rubySettings)
      : undefined;
    message.goSettings = (object.goSettings !== undefined && object.goSettings !== null)
      ? GoSettings.fromPartial(object.goSettings)
      : undefined;
    return message;
  },
};

function createBasePublishing(): Publishing {
  return {
    methodSettings: [],
    newIssueUri: "",
    documentationUri: "",
    apiShortName: "",
    githubLabel: "",
    codeownerGithubTeams: [],
    docTagPrefix: "",
    organization: ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED,
    librarySettings: [],
    protoReferenceDocumentationUri: "",
  };
}

export const Publishing = {
  encode(message: Publishing, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.methodSettings) {
      MethodSettings.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.newIssueUri !== "") {
      writer.uint32(810).string(message.newIssueUri);
    }
    if (message.documentationUri !== "") {
      writer.uint32(818).string(message.documentationUri);
    }
    if (message.apiShortName !== "") {
      writer.uint32(826).string(message.apiShortName);
    }
    if (message.githubLabel !== "") {
      writer.uint32(834).string(message.githubLabel);
    }
    for (const v of message.codeownerGithubTeams) {
      writer.uint32(842).string(v!);
    }
    if (message.docTagPrefix !== "") {
      writer.uint32(850).string(message.docTagPrefix);
    }
    if (message.organization !== ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED) {
      writer.uint32(856).int32(clientLibraryOrganizationToNumber(message.organization));
    }
    for (const v of message.librarySettings) {
      ClientLibrarySettings.encode(v!, writer.uint32(874).fork()).ldelim();
    }
    if (message.protoReferenceDocumentationUri !== "") {
      writer.uint32(882).string(message.protoReferenceDocumentationUri);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Publishing {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePublishing();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.methodSettings.push(MethodSettings.decode(reader, reader.uint32()));
          continue;
        case 101:
          if (tag !== 810) {
            break;
          }

          message.newIssueUri = reader.string();
          continue;
        case 102:
          if (tag !== 818) {
            break;
          }

          message.documentationUri = reader.string();
          continue;
        case 103:
          if (tag !== 826) {
            break;
          }

          message.apiShortName = reader.string();
          continue;
        case 104:
          if (tag !== 834) {
            break;
          }

          message.githubLabel = reader.string();
          continue;
        case 105:
          if (tag !== 842) {
            break;
          }

          message.codeownerGithubTeams.push(reader.string());
          continue;
        case 106:
          if (tag !== 850) {
            break;
          }

          message.docTagPrefix = reader.string();
          continue;
        case 107:
          if (tag !== 856) {
            break;
          }

          message.organization = clientLibraryOrganizationFromJSON(reader.int32());
          continue;
        case 109:
          if (tag !== 874) {
            break;
          }

          message.librarySettings.push(ClientLibrarySettings.decode(reader, reader.uint32()));
          continue;
        case 110:
          if (tag !== 882) {
            break;
          }

          message.protoReferenceDocumentationUri = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Publishing {
    return {
      methodSettings: globalThis.Array.isArray(object?.methodSettings)
        ? object.methodSettings.map((e: any) => MethodSettings.fromJSON(e))
        : [],
      newIssueUri: isSet(object.newIssueUri) ? globalThis.String(object.newIssueUri) : "",
      documentationUri: isSet(object.documentationUri) ? globalThis.String(object.documentationUri) : "",
      apiShortName: isSet(object.apiShortName) ? globalThis.String(object.apiShortName) : "",
      githubLabel: isSet(object.githubLabel) ? globalThis.String(object.githubLabel) : "",
      codeownerGithubTeams: globalThis.Array.isArray(object?.codeownerGithubTeams)
        ? object.codeownerGithubTeams.map((e: any) => globalThis.String(e))
        : [],
      docTagPrefix: isSet(object.docTagPrefix) ? globalThis.String(object.docTagPrefix) : "",
      organization: isSet(object.organization)
        ? clientLibraryOrganizationFromJSON(object.organization)
        : ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED,
      librarySettings: globalThis.Array.isArray(object?.librarySettings)
        ? object.librarySettings.map((e: any) => ClientLibrarySettings.fromJSON(e))
        : [],
      protoReferenceDocumentationUri: isSet(object.protoReferenceDocumentationUri)
        ? globalThis.String(object.protoReferenceDocumentationUri)
        : "",
    };
  },

  toJSON(message: Publishing): unknown {
    const obj: any = {};
    if (message.methodSettings?.length) {
      obj.methodSettings = message.methodSettings.map((e) => MethodSettings.toJSON(e));
    }
    if (message.newIssueUri !== "") {
      obj.newIssueUri = message.newIssueUri;
    }
    if (message.documentationUri !== "") {
      obj.documentationUri = message.documentationUri;
    }
    if (message.apiShortName !== "") {
      obj.apiShortName = message.apiShortName;
    }
    if (message.githubLabel !== "") {
      obj.githubLabel = message.githubLabel;
    }
    if (message.codeownerGithubTeams?.length) {
      obj.codeownerGithubTeams = message.codeownerGithubTeams;
    }
    if (message.docTagPrefix !== "") {
      obj.docTagPrefix = message.docTagPrefix;
    }
    if (message.organization !== ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED) {
      obj.organization = clientLibraryOrganizationToJSON(message.organization);
    }
    if (message.librarySettings?.length) {
      obj.librarySettings = message.librarySettings.map((e) => ClientLibrarySettings.toJSON(e));
    }
    if (message.protoReferenceDocumentationUri !== "") {
      obj.protoReferenceDocumentationUri = message.protoReferenceDocumentationUri;
    }
    return obj;
  },

  create(base?: DeepPartial<Publishing>): Publishing {
    return Publishing.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Publishing>): Publishing {
    const message = createBasePublishing();
    message.methodSettings = object.methodSettings?.map((e) => MethodSettings.fromPartial(e)) || [];
    message.newIssueUri = object.newIssueUri ?? "";
    message.documentationUri = object.documentationUri ?? "";
    message.apiShortName = object.apiShortName ?? "";
    message.githubLabel = object.githubLabel ?? "";
    message.codeownerGithubTeams = object.codeownerGithubTeams?.map((e) => e) || [];
    message.docTagPrefix = object.docTagPrefix ?? "";
    message.organization = object.organization ?? ClientLibraryOrganization.CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED;
    message.librarySettings = object.librarySettings?.map((e) => ClientLibrarySettings.fromPartial(e)) || [];
    message.protoReferenceDocumentationUri = object.protoReferenceDocumentationUri ?? "";
    return message;
  },
};

function createBaseJavaSettings(): JavaSettings {
  return { libraryPackage: "", serviceClassNames: {}, common: undefined };
}

export const JavaSettings = {
  encode(message: JavaSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.libraryPackage !== "") {
      writer.uint32(10).string(message.libraryPackage);
    }
    Object.entries(message.serviceClassNames).forEach(([key, value]) => {
      JavaSettings_ServiceClassNamesEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JavaSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJavaSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.libraryPackage = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          const entry2 = JavaSettings_ServiceClassNamesEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.serviceClassNames[entry2.key] = entry2.value;
          }
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): JavaSettings {
    return {
      libraryPackage: isSet(object.libraryPackage) ? globalThis.String(object.libraryPackage) : "",
      serviceClassNames: isObject(object.serviceClassNames)
        ? Object.entries(object.serviceClassNames).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined,
    };
  },

  toJSON(message: JavaSettings): unknown {
    const obj: any = {};
    if (message.libraryPackage !== "") {
      obj.libraryPackage = message.libraryPackage;
    }
    if (message.serviceClassNames) {
      const entries = Object.entries(message.serviceClassNames);
      if (entries.length > 0) {
        obj.serviceClassNames = {};
        entries.forEach(([k, v]) => {
          obj.serviceClassNames[k] = v;
        });
      }
    }
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<JavaSettings>): JavaSettings {
    return JavaSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<JavaSettings>): JavaSettings {
    const message = createBaseJavaSettings();
    message.libraryPackage = object.libraryPackage ?? "";
    message.serviceClassNames = Object.entries(object.serviceClassNames ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = globalThis.String(value);
        }
        return acc;
      },
      {},
    );
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBaseJavaSettings_ServiceClassNamesEntry(): JavaSettings_ServiceClassNamesEntry {
  return { key: "", value: "" };
}

export const JavaSettings_ServiceClassNamesEntry = {
  encode(message: JavaSettings_ServiceClassNamesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JavaSettings_ServiceClassNamesEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJavaSettings_ServiceClassNamesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): JavaSettings_ServiceClassNamesEntry {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: JavaSettings_ServiceClassNamesEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create(base?: DeepPartial<JavaSettings_ServiceClassNamesEntry>): JavaSettings_ServiceClassNamesEntry {
    return JavaSettings_ServiceClassNamesEntry.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<JavaSettings_ServiceClassNamesEntry>): JavaSettings_ServiceClassNamesEntry {
    const message = createBaseJavaSettings_ServiceClassNamesEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCppSettings(): CppSettings {
  return { common: undefined };
}

export const CppSettings = {
  encode(message: CppSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CppSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCppSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CppSettings {
    return { common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined };
  },

  toJSON(message: CppSettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<CppSettings>): CppSettings {
    return CppSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CppSettings>): CppSettings {
    const message = createBaseCppSettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBasePhpSettings(): PhpSettings {
  return { common: undefined };
}

export const PhpSettings = {
  encode(message: PhpSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PhpSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhpSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PhpSettings {
    return { common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined };
  },

  toJSON(message: PhpSettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<PhpSettings>): PhpSettings {
    return PhpSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PhpSettings>): PhpSettings {
    const message = createBasePhpSettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBasePythonSettings(): PythonSettings {
  return { common: undefined };
}

export const PythonSettings = {
  encode(message: PythonSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PythonSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePythonSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PythonSettings {
    return { common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined };
  },

  toJSON(message: PythonSettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<PythonSettings>): PythonSettings {
    return PythonSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PythonSettings>): PythonSettings {
    const message = createBasePythonSettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBaseNodeSettings(): NodeSettings {
  return { common: undefined };
}

export const NodeSettings = {
  encode(message: NodeSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NodeSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNodeSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): NodeSettings {
    return { common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined };
  },

  toJSON(message: NodeSettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<NodeSettings>): NodeSettings {
    return NodeSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<NodeSettings>): NodeSettings {
    const message = createBaseNodeSettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBaseDotnetSettings(): DotnetSettings {
  return {
    common: undefined,
    renamedServices: {},
    renamedResources: {},
    ignoredResources: [],
    forcedNamespaceAliases: [],
    handwrittenSignatures: [],
  };
}

export const DotnetSettings = {
  encode(message: DotnetSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    Object.entries(message.renamedServices).forEach(([key, value]) => {
      DotnetSettings_RenamedServicesEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    Object.entries(message.renamedResources).forEach(([key, value]) => {
      DotnetSettings_RenamedResourcesEntry.encode({ key: key as any, value }, writer.uint32(26).fork()).ldelim();
    });
    for (const v of message.ignoredResources) {
      writer.uint32(34).string(v!);
    }
    for (const v of message.forcedNamespaceAliases) {
      writer.uint32(42).string(v!);
    }
    for (const v of message.handwrittenSignatures) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DotnetSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDotnetSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          const entry2 = DotnetSettings_RenamedServicesEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.renamedServices[entry2.key] = entry2.value;
          }
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          const entry3 = DotnetSettings_RenamedResourcesEntry.decode(reader, reader.uint32());
          if (entry3.value !== undefined) {
            message.renamedResources[entry3.key] = entry3.value;
          }
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.ignoredResources.push(reader.string());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.forcedNamespaceAliases.push(reader.string());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.handwrittenSignatures.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DotnetSettings {
    return {
      common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined,
      renamedServices: isObject(object.renamedServices)
        ? Object.entries(object.renamedServices).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      renamedResources: isObject(object.renamedResources)
        ? Object.entries(object.renamedResources).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      ignoredResources: globalThis.Array.isArray(object?.ignoredResources)
        ? object.ignoredResources.map((e: any) => globalThis.String(e))
        : [],
      forcedNamespaceAliases: globalThis.Array.isArray(object?.forcedNamespaceAliases)
        ? object.forcedNamespaceAliases.map((e: any) => globalThis.String(e))
        : [],
      handwrittenSignatures: globalThis.Array.isArray(object?.handwrittenSignatures)
        ? object.handwrittenSignatures.map((e: any) => globalThis.String(e))
        : [],
    };
  },

  toJSON(message: DotnetSettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    if (message.renamedServices) {
      const entries = Object.entries(message.renamedServices);
      if (entries.length > 0) {
        obj.renamedServices = {};
        entries.forEach(([k, v]) => {
          obj.renamedServices[k] = v;
        });
      }
    }
    if (message.renamedResources) {
      const entries = Object.entries(message.renamedResources);
      if (entries.length > 0) {
        obj.renamedResources = {};
        entries.forEach(([k, v]) => {
          obj.renamedResources[k] = v;
        });
      }
    }
    if (message.ignoredResources?.length) {
      obj.ignoredResources = message.ignoredResources;
    }
    if (message.forcedNamespaceAliases?.length) {
      obj.forcedNamespaceAliases = message.forcedNamespaceAliases;
    }
    if (message.handwrittenSignatures?.length) {
      obj.handwrittenSignatures = message.handwrittenSignatures;
    }
    return obj;
  },

  create(base?: DeepPartial<DotnetSettings>): DotnetSettings {
    return DotnetSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<DotnetSettings>): DotnetSettings {
    const message = createBaseDotnetSettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    message.renamedServices = Object.entries(object.renamedServices ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = globalThis.String(value);
        }
        return acc;
      },
      {},
    );
    message.renamedResources = Object.entries(object.renamedResources ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = globalThis.String(value);
        }
        return acc;
      },
      {},
    );
    message.ignoredResources = object.ignoredResources?.map((e) => e) || [];
    message.forcedNamespaceAliases = object.forcedNamespaceAliases?.map((e) => e) || [];
    message.handwrittenSignatures = object.handwrittenSignatures?.map((e) => e) || [];
    return message;
  },
};

function createBaseDotnetSettings_RenamedServicesEntry(): DotnetSettings_RenamedServicesEntry {
  return { key: "", value: "" };
}

export const DotnetSettings_RenamedServicesEntry = {
  encode(message: DotnetSettings_RenamedServicesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DotnetSettings_RenamedServicesEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDotnetSettings_RenamedServicesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DotnetSettings_RenamedServicesEntry {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: DotnetSettings_RenamedServicesEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create(base?: DeepPartial<DotnetSettings_RenamedServicesEntry>): DotnetSettings_RenamedServicesEntry {
    return DotnetSettings_RenamedServicesEntry.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<DotnetSettings_RenamedServicesEntry>): DotnetSettings_RenamedServicesEntry {
    const message = createBaseDotnetSettings_RenamedServicesEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseDotnetSettings_RenamedResourcesEntry(): DotnetSettings_RenamedResourcesEntry {
  return { key: "", value: "" };
}

export const DotnetSettings_RenamedResourcesEntry = {
  encode(message: DotnetSettings_RenamedResourcesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DotnetSettings_RenamedResourcesEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDotnetSettings_RenamedResourcesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DotnetSettings_RenamedResourcesEntry {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: DotnetSettings_RenamedResourcesEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create(base?: DeepPartial<DotnetSettings_RenamedResourcesEntry>): DotnetSettings_RenamedResourcesEntry {
    return DotnetSettings_RenamedResourcesEntry.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<DotnetSettings_RenamedResourcesEntry>): DotnetSettings_RenamedResourcesEntry {
    const message = createBaseDotnetSettings_RenamedResourcesEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseRubySettings(): RubySettings {
  return { common: undefined };
}

export const RubySettings = {
  encode(message: RubySettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RubySettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRubySettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RubySettings {
    return { common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined };
  },

  toJSON(message: RubySettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<RubySettings>): RubySettings {
    return RubySettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<RubySettings>): RubySettings {
    const message = createBaseRubySettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBaseGoSettings(): GoSettings {
  return { common: undefined };
}

export const GoSettings = {
  encode(message: GoSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.common !== undefined) {
      CommonLanguageSettings.encode(message.common, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GoSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGoSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.common = CommonLanguageSettings.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GoSettings {
    return { common: isSet(object.common) ? CommonLanguageSettings.fromJSON(object.common) : undefined };
  },

  toJSON(message: GoSettings): unknown {
    const obj: any = {};
    if (message.common !== undefined) {
      obj.common = CommonLanguageSettings.toJSON(message.common);
    }
    return obj;
  },

  create(base?: DeepPartial<GoSettings>): GoSettings {
    return GoSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GoSettings>): GoSettings {
    const message = createBaseGoSettings();
    message.common = (object.common !== undefined && object.common !== null)
      ? CommonLanguageSettings.fromPartial(object.common)
      : undefined;
    return message;
  },
};

function createBaseMethodSettings(): MethodSettings {
  return { selector: "", longRunning: undefined };
}

export const MethodSettings = {
  encode(message: MethodSettings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.selector !== "") {
      writer.uint32(10).string(message.selector);
    }
    if (message.longRunning !== undefined) {
      MethodSettings_LongRunning.encode(message.longRunning, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MethodSettings {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMethodSettings();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.selector = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.longRunning = MethodSettings_LongRunning.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MethodSettings {
    return {
      selector: isSet(object.selector) ? globalThis.String(object.selector) : "",
      longRunning: isSet(object.longRunning) ? MethodSettings_LongRunning.fromJSON(object.longRunning) : undefined,
    };
  },

  toJSON(message: MethodSettings): unknown {
    const obj: any = {};
    if (message.selector !== "") {
      obj.selector = message.selector;
    }
    if (message.longRunning !== undefined) {
      obj.longRunning = MethodSettings_LongRunning.toJSON(message.longRunning);
    }
    return obj;
  },

  create(base?: DeepPartial<MethodSettings>): MethodSettings {
    return MethodSettings.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MethodSettings>): MethodSettings {
    const message = createBaseMethodSettings();
    message.selector = object.selector ?? "";
    message.longRunning = (object.longRunning !== undefined && object.longRunning !== null)
      ? MethodSettings_LongRunning.fromPartial(object.longRunning)
      : undefined;
    return message;
  },
};

function createBaseMethodSettings_LongRunning(): MethodSettings_LongRunning {
  return { initialPollDelay: undefined, pollDelayMultiplier: 0, maxPollDelay: undefined, totalPollTimeout: undefined };
}

export const MethodSettings_LongRunning = {
  encode(message: MethodSettings_LongRunning, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.initialPollDelay !== undefined) {
      Duration.encode(message.initialPollDelay, writer.uint32(10).fork()).ldelim();
    }
    if (message.pollDelayMultiplier !== 0) {
      writer.uint32(21).float(message.pollDelayMultiplier);
    }
    if (message.maxPollDelay !== undefined) {
      Duration.encode(message.maxPollDelay, writer.uint32(26).fork()).ldelim();
    }
    if (message.totalPollTimeout !== undefined) {
      Duration.encode(message.totalPollTimeout, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MethodSettings_LongRunning {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMethodSettings_LongRunning();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.initialPollDelay = Duration.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 21) {
            break;
          }

          message.pollDelayMultiplier = reader.float();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.maxPollDelay = Duration.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.totalPollTimeout = Duration.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MethodSettings_LongRunning {
    return {
      initialPollDelay: isSet(object.initialPollDelay) ? Duration.fromJSON(object.initialPollDelay) : undefined,
      pollDelayMultiplier: isSet(object.pollDelayMultiplier) ? globalThis.Number(object.pollDelayMultiplier) : 0,
      maxPollDelay: isSet(object.maxPollDelay) ? Duration.fromJSON(object.maxPollDelay) : undefined,
      totalPollTimeout: isSet(object.totalPollTimeout) ? Duration.fromJSON(object.totalPollTimeout) : undefined,
    };
  },

  toJSON(message: MethodSettings_LongRunning): unknown {
    const obj: any = {};
    if (message.initialPollDelay !== undefined) {
      obj.initialPollDelay = Duration.toJSON(message.initialPollDelay);
    }
    if (message.pollDelayMultiplier !== 0) {
      obj.pollDelayMultiplier = message.pollDelayMultiplier;
    }
    if (message.maxPollDelay !== undefined) {
      obj.maxPollDelay = Duration.toJSON(message.maxPollDelay);
    }
    if (message.totalPollTimeout !== undefined) {
      obj.totalPollTimeout = Duration.toJSON(message.totalPollTimeout);
    }
    return obj;
  },

  create(base?: DeepPartial<MethodSettings_LongRunning>): MethodSettings_LongRunning {
    return MethodSettings_LongRunning.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MethodSettings_LongRunning>): MethodSettings_LongRunning {
    const message = createBaseMethodSettings_LongRunning();
    message.initialPollDelay = (object.initialPollDelay !== undefined && object.initialPollDelay !== null)
      ? Duration.fromPartial(object.initialPollDelay)
      : undefined;
    message.pollDelayMultiplier = object.pollDelayMultiplier ?? 0;
    message.maxPollDelay = (object.maxPollDelay !== undefined && object.maxPollDelay !== null)
      ? Duration.fromPartial(object.maxPollDelay)
      : undefined;
    message.totalPollTimeout = (object.totalPollTimeout !== undefined && object.totalPollTimeout !== null)
      ? Duration.fromPartial(object.totalPollTimeout)
      : undefined;
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

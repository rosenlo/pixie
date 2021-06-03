/* tslint:disable */
/* eslint-disable */
import { GraphQLResolveInfo } from 'graphql';
/**
 * This file is auto-generated by graphql-schema-typescript
 * Please note that any changes in this file may be overwritten
 */
 

/*******************************
 *                             *
 *          TYPE DEFS          *
 *                             *
 *******************************/
export interface GQLUserInfo {
  id: string;
  name: string;
  email: string;
  picture: string;
  orgName: string;
  orgID: string;
  isApproved: boolean;
}

export interface GQLOrgInfo {
  id: string;
  name: string;
  enableApprovals: boolean;
}

export interface GQLUserSetting {
  key: string;
  value: string;
}

export const enum GQLArtifactType {
  AT_UNKNOWN = 'AT_UNKNOWN',
  AT_LINUX_AMD64 = 'AT_LINUX_AMD64',
  AT_DARWIN_AMD64 = 'AT_DARWIN_AMD64',
  AT_CONTAINER_SET_YAMLS = 'AT_CONTAINER_SET_YAMLS',
  AT_CONTAINER_SET_LINUX_AMD64 = 'AT_CONTAINER_SET_LINUX_AMD64',
  AT_CONTAINER_SET_TEMPLATE_YAMLS = 'AT_CONTAINER_SET_TEMPLATE_YAMLS'
}

export const enum GQLAutocompleteEntityState {
  AES_UNKNOWN = 'AES_UNKNOWN',
  AES_PENDING = 'AES_PENDING',
  AES_RUNNING = 'AES_RUNNING',
  AES_FAILED = 'AES_FAILED',
  AES_TERMINATED = 'AES_TERMINATED'
}

export const enum GQLAutocompleteActionType {
  AAT_UNKNOWN = 'AAT_UNKNOWN',
  AAT_EDIT = 'AAT_EDIT',
  AAT_SELECT = 'AAT_SELECT'
}

export const enum GQLAutocompleteEntityKind {
  AEK_UNKNOWN = 'AEK_UNKNOWN',
  AEK_POD = 'AEK_POD',
  AEK_SVC = 'AEK_SVC',
  AEK_SCRIPT = 'AEK_SCRIPT',
  AEK_NAMESPACE = 'AEK_NAMESPACE'
}

export interface GQLDeploymentKey {
  id: string;
  key: string;
  createdAtMs: number;
  desc: string;
}

export interface GQLAPIKey {
  id: string;
  key: string;
  createdAtMs: number;
  desc: string;
}

export interface GQLQuery {
  user: GQLUserInfo;
  org: GQLOrgInfo;
  userSettings: Array<GQLUserSetting>;
  orgUsers: Array<GQLUserInfo | null>;
  cluster: GQLClusterInfo;
  clusterByName: GQLClusterInfo;
  clusters: Array<GQLClusterInfo>;
  clusterConnection: GQLClusterConnectionInfo;
  cliArtifact: GQLCLIArtifact;
  artifacts: GQLArtifactsInfo;
  autocomplete: GQLAutocompleteResult;
  autocompleteField?: Array<GQLAutocompleteSuggestion | null>;
  liveViews: Array<GQLLiveViewMetadata>;
  liveViewContents: GQLLiveViewContents;
  scripts: Array<GQLScriptMetadata>;
  scriptContents: GQLScriptContents;
  deploymentKeys: Array<GQLDeploymentKey>;
  deploymentKey: GQLDeploymentKey;
  apiKeys: Array<GQLAPIKey>;
  apiKey: GQLAPIKey;
}

export interface GQLAutocompleteResult {
  formattedInput?: string;
  isExecutable?: boolean;
  tabSuggestions?: Array<GQLTabSuggestion | null>;
}

export interface GQLTabSuggestion {
  tabIndex?: number;
  executableAfterSelect?: boolean;
  suggestions?: Array<GQLAutocompleteSuggestion | null>;
}

export interface GQLAutocompleteSuggestion {
  kind?: GQLAutocompleteEntityKind;
  name?: string;
  description?: string;
  matchedIndexes?: Array<number | null>;
  state?: GQLAutocompleteEntityState;
}

export const enum GQLClusterStatus {
  CS_UNKNOWN = 'CS_UNKNOWN',
  CS_HEALTHY = 'CS_HEALTHY',
  CS_UNHEALTHY = 'CS_UNHEALTHY',
  CS_DISCONNECTED = 'CS_DISCONNECTED',
  CS_UPDATING = 'CS_UPDATING',
  CS_CONNECTED = 'CS_CONNECTED',
  CS_UPDATE_FAILED = 'CS_UPDATE_FAILED'
}

export interface GQLVizierConfig {
  passthroughEnabled?: boolean;
}

export interface GQLContainerStatus {
  name: string;
  createdAtMs: number;
  state: string;
  message?: string;
  reason?: string;
}

export interface GQLK8sEvent {
  message?: string;
  firstTimeMs: number;
  lastTimeMs: number;
}

export interface GQLPodStatus {
  name: string;
  createdAtMs: number;
  status: string;
  message?: string;
  reason?: string;
  containers: Array<GQLContainerStatus>;
  events: Array<GQLK8sEvent>;
}

export interface GQLClusterInfo {
  id: string;
  status: GQLClusterStatus;
  lastHeartbeatMs: number;
  vizierConfig: GQLVizierConfig;
  vizierVersion?: string;
  clusterVersion?: string;
  clusterName?: string;
  prettyClusterName?: string;
  clusterUID?: string;
  controlPlanePodStatuses: Array<GQLPodStatus>;
  numNodes: number;
  numInstrumentedNodes: number;
}

export interface GQLClusterConnectionInfo {
  ipAddress: string;
  token: string;
}

export interface GQLUserInvite {
  email: string;
  inviteLink: string;
}

export interface GQLMutation {
  
  /**
   * 
   * @deprecated Clusters are now created via px deploy
   */
  CreateCluster?: GQLClusterInfo;
  UpdateVizierConfig: boolean;
  CreateDeploymentKey: GQLDeploymentKey;
  DeleteDeploymentKey: boolean;
  CreateAPIKey: GQLAPIKey;
  DeleteAPIKey: boolean;
  UpdateUserSettings: Array<GQLUserSetting>;
  InviteUser: GQLUserInvite;
  UpdateUserPermissions: GQLUserInfo;
  UpdateOrgSettings: GQLOrgInfo;
}

export interface GQLEditableUserPermissions {
  isApproved?: boolean;
}

export interface GQLEditableOrgSettings {
  enableApprovals?: boolean;
}

export interface GQLCLIArtifact {
  url: string;
  sha256: string;
}

export interface GQLArtifactsInfo {
  items?: Array<GQLArtifact | null>;
}

export interface GQLArtifact {
  version: string;
  changelog: string;
  timestampMs: number;
}

export interface GQLLiveViewMetadata {
  id: string;
  name: string;
  desc: string;
}

export interface GQLLiveViewContents {
  metadata: GQLLiveViewMetadata;
  pxlContents: string;
  visJSON: string;
}

export interface GQLScriptMetadata {
  id: string;
  name: string;
  desc: string;
  hasLiveView: boolean;
}

export interface GQLScriptContents {
  metadata: GQLScriptMetadata;
  contents: string;
}

/*********************************
 *                               *
 *         TYPE RESOLVERS        *
 *                               *
 *********************************/
/**
 * This interface define the shape of your resolver
 * Note that this type is designed to be compatible with graphql-tools resolvers
 * However, you can still use other generated interfaces to make your resolver type-safed
 */
export interface GQLResolver {
  UserInfo?: GQLUserInfoTypeResolver;
  OrgInfo?: GQLOrgInfoTypeResolver;
  UserSetting?: GQLUserSettingTypeResolver;
  DeploymentKey?: GQLDeploymentKeyTypeResolver;
  APIKey?: GQLAPIKeyTypeResolver;
  Query?: GQLQueryTypeResolver;
  AutocompleteResult?: GQLAutocompleteResultTypeResolver;
  TabSuggestion?: GQLTabSuggestionTypeResolver;
  AutocompleteSuggestion?: GQLAutocompleteSuggestionTypeResolver;
  VizierConfig?: GQLVizierConfigTypeResolver;
  ContainerStatus?: GQLContainerStatusTypeResolver;
  K8sEvent?: GQLK8sEventTypeResolver;
  PodStatus?: GQLPodStatusTypeResolver;
  ClusterInfo?: GQLClusterInfoTypeResolver;
  ClusterConnectionInfo?: GQLClusterConnectionInfoTypeResolver;
  UserInvite?: GQLUserInviteTypeResolver;
  Mutation?: GQLMutationTypeResolver;
  CLIArtifact?: GQLCLIArtifactTypeResolver;
  ArtifactsInfo?: GQLArtifactsInfoTypeResolver;
  Artifact?: GQLArtifactTypeResolver;
  LiveViewMetadata?: GQLLiveViewMetadataTypeResolver;
  LiveViewContents?: GQLLiveViewContentsTypeResolver;
  ScriptMetadata?: GQLScriptMetadataTypeResolver;
  ScriptContents?: GQLScriptContentsTypeResolver;
}
export interface GQLUserInfoTypeResolver<TParent = any> {
  id?: UserInfoToIdResolver<TParent>;
  name?: UserInfoToNameResolver<TParent>;
  email?: UserInfoToEmailResolver<TParent>;
  picture?: UserInfoToPictureResolver<TParent>;
  orgName?: UserInfoToOrgNameResolver<TParent>;
  orgID?: UserInfoToOrgIDResolver<TParent>;
  isApproved?: UserInfoToIsApprovedResolver<TParent>;
}

export interface UserInfoToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInfoToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInfoToEmailResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInfoToPictureResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInfoToOrgNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInfoToOrgIDResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInfoToIsApprovedResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLOrgInfoTypeResolver<TParent = any> {
  id?: OrgInfoToIdResolver<TParent>;
  name?: OrgInfoToNameResolver<TParent>;
  enableApprovals?: OrgInfoToEnableApprovalsResolver<TParent>;
}

export interface OrgInfoToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface OrgInfoToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface OrgInfoToEnableApprovalsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLUserSettingTypeResolver<TParent = any> {
  key?: UserSettingToKeyResolver<TParent>;
  value?: UserSettingToValueResolver<TParent>;
}

export interface UserSettingToKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserSettingToValueResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLDeploymentKeyTypeResolver<TParent = any> {
  id?: DeploymentKeyToIdResolver<TParent>;
  key?: DeploymentKeyToKeyResolver<TParent>;
  createdAtMs?: DeploymentKeyToCreatedAtMsResolver<TParent>;
  desc?: DeploymentKeyToDescResolver<TParent>;
}

export interface DeploymentKeyToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface DeploymentKeyToKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface DeploymentKeyToCreatedAtMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface DeploymentKeyToDescResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLAPIKeyTypeResolver<TParent = any> {
  id?: APIKeyToIdResolver<TParent>;
  key?: APIKeyToKeyResolver<TParent>;
  createdAtMs?: APIKeyToCreatedAtMsResolver<TParent>;
  desc?: APIKeyToDescResolver<TParent>;
}

export interface APIKeyToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface APIKeyToKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface APIKeyToCreatedAtMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface APIKeyToDescResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLQueryTypeResolver<TParent = any> {
  user?: QueryToUserResolver<TParent>;
  org?: QueryToOrgResolver<TParent>;
  userSettings?: QueryToUserSettingsResolver<TParent>;
  orgUsers?: QueryToOrgUsersResolver<TParent>;
  cluster?: QueryToClusterResolver<TParent>;
  clusterByName?: QueryToClusterByNameResolver<TParent>;
  clusters?: QueryToClustersResolver<TParent>;
  clusterConnection?: QueryToClusterConnectionResolver<TParent>;
  cliArtifact?: QueryToCliArtifactResolver<TParent>;
  artifacts?: QueryToArtifactsResolver<TParent>;
  autocomplete?: QueryToAutocompleteResolver<TParent>;
  autocompleteField?: QueryToAutocompleteFieldResolver<TParent>;
  liveViews?: QueryToLiveViewsResolver<TParent>;
  liveViewContents?: QueryToLiveViewContentsResolver<TParent>;
  scripts?: QueryToScriptsResolver<TParent>;
  scriptContents?: QueryToScriptContentsResolver<TParent>;
  deploymentKeys?: QueryToDeploymentKeysResolver<TParent>;
  deploymentKey?: QueryToDeploymentKeyResolver<TParent>;
  apiKeys?: QueryToApiKeysResolver<TParent>;
  apiKey?: QueryToApiKeyResolver<TParent>;
}

export interface QueryToUserResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToOrgResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToUserSettingsArgs {
  keys: Array<string>;
}
export interface QueryToUserSettingsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToUserSettingsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToOrgUsersResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToClusterArgs {
  id: string;
}
export interface QueryToClusterResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToClusterArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToClusterByNameArgs {
  name: string;
}
export interface QueryToClusterByNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToClusterByNameArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToClustersResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToClusterConnectionArgs {
  id: string;
}
export interface QueryToClusterConnectionResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToClusterConnectionArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToCliArtifactArgs {
  artifactType?: GQLArtifactType;
}
export interface QueryToCliArtifactResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToCliArtifactArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToArtifactsArgs {
  artifactName?: string;
}
export interface QueryToArtifactsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToArtifactsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToAutocompleteArgs {
  input?: string;
  cursorPos?: number;
  action?: GQLAutocompleteActionType;
  clusterUID?: string;
}
export interface QueryToAutocompleteResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToAutocompleteArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToAutocompleteFieldArgs {
  input?: string;
  fieldType?: GQLAutocompleteEntityKind;
  requiredArgTypes?: Array<GQLAutocompleteEntityKind | null>;
  clusterUID?: string;
}
export interface QueryToAutocompleteFieldResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToAutocompleteFieldArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToLiveViewsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToLiveViewContentsArgs {
  id: string;
}
export interface QueryToLiveViewContentsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToLiveViewContentsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToScriptsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToScriptContentsArgs {
  id: string;
}
export interface QueryToScriptContentsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToScriptContentsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToDeploymentKeysResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToDeploymentKeyArgs {
  id: string;
}
export interface QueryToDeploymentKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToDeploymentKeyArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToApiKeysResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface QueryToApiKeyArgs {
  id: string;
}
export interface QueryToApiKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: QueryToApiKeyArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLAutocompleteResultTypeResolver<TParent = any> {
  formattedInput?: AutocompleteResultToFormattedInputResolver<TParent>;
  isExecutable?: AutocompleteResultToIsExecutableResolver<TParent>;
  tabSuggestions?: AutocompleteResultToTabSuggestionsResolver<TParent>;
}

export interface AutocompleteResultToFormattedInputResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface AutocompleteResultToIsExecutableResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface AutocompleteResultToTabSuggestionsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLTabSuggestionTypeResolver<TParent = any> {
  tabIndex?: TabSuggestionToTabIndexResolver<TParent>;
  executableAfterSelect?: TabSuggestionToExecutableAfterSelectResolver<TParent>;
  suggestions?: TabSuggestionToSuggestionsResolver<TParent>;
}

export interface TabSuggestionToTabIndexResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface TabSuggestionToExecutableAfterSelectResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface TabSuggestionToSuggestionsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLAutocompleteSuggestionTypeResolver<TParent = any> {
  kind?: AutocompleteSuggestionToKindResolver<TParent>;
  name?: AutocompleteSuggestionToNameResolver<TParent>;
  description?: AutocompleteSuggestionToDescriptionResolver<TParent>;
  matchedIndexes?: AutocompleteSuggestionToMatchedIndexesResolver<TParent>;
  state?: AutocompleteSuggestionToStateResolver<TParent>;
}

export interface AutocompleteSuggestionToKindResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface AutocompleteSuggestionToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface AutocompleteSuggestionToDescriptionResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface AutocompleteSuggestionToMatchedIndexesResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface AutocompleteSuggestionToStateResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLVizierConfigTypeResolver<TParent = any> {
  passthroughEnabled?: VizierConfigToPassthroughEnabledResolver<TParent>;
}

export interface VizierConfigToPassthroughEnabledResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLContainerStatusTypeResolver<TParent = any> {
  name?: ContainerStatusToNameResolver<TParent>;
  createdAtMs?: ContainerStatusToCreatedAtMsResolver<TParent>;
  state?: ContainerStatusToStateResolver<TParent>;
  message?: ContainerStatusToMessageResolver<TParent>;
  reason?: ContainerStatusToReasonResolver<TParent>;
}

export interface ContainerStatusToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ContainerStatusToCreatedAtMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ContainerStatusToStateResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ContainerStatusToMessageResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ContainerStatusToReasonResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLK8sEventTypeResolver<TParent = any> {
  message?: K8sEventToMessageResolver<TParent>;
  firstTimeMs?: K8sEventToFirstTimeMsResolver<TParent>;
  lastTimeMs?: K8sEventToLastTimeMsResolver<TParent>;
}

export interface K8sEventToMessageResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface K8sEventToFirstTimeMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface K8sEventToLastTimeMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLPodStatusTypeResolver<TParent = any> {
  name?: PodStatusToNameResolver<TParent>;
  createdAtMs?: PodStatusToCreatedAtMsResolver<TParent>;
  status?: PodStatusToStatusResolver<TParent>;
  message?: PodStatusToMessageResolver<TParent>;
  reason?: PodStatusToReasonResolver<TParent>;
  containers?: PodStatusToContainersResolver<TParent>;
  events?: PodStatusToEventsResolver<TParent>;
}

export interface PodStatusToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface PodStatusToCreatedAtMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface PodStatusToStatusResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface PodStatusToMessageResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface PodStatusToReasonResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface PodStatusToContainersResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface PodStatusToEventsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLClusterInfoTypeResolver<TParent = any> {
  id?: ClusterInfoToIdResolver<TParent>;
  status?: ClusterInfoToStatusResolver<TParent>;
  lastHeartbeatMs?: ClusterInfoToLastHeartbeatMsResolver<TParent>;
  vizierConfig?: ClusterInfoToVizierConfigResolver<TParent>;
  vizierVersion?: ClusterInfoToVizierVersionResolver<TParent>;
  clusterVersion?: ClusterInfoToClusterVersionResolver<TParent>;
  clusterName?: ClusterInfoToClusterNameResolver<TParent>;
  prettyClusterName?: ClusterInfoToPrettyClusterNameResolver<TParent>;
  clusterUID?: ClusterInfoToClusterUIDResolver<TParent>;
  controlPlanePodStatuses?: ClusterInfoToControlPlanePodStatusesResolver<TParent>;
  numNodes?: ClusterInfoToNumNodesResolver<TParent>;
  numInstrumentedNodes?: ClusterInfoToNumInstrumentedNodesResolver<TParent>;
}

export interface ClusterInfoToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToStatusResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToLastHeartbeatMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToVizierConfigResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToVizierVersionResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToClusterVersionResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToClusterNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToPrettyClusterNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToClusterUIDResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToControlPlanePodStatusesResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToNumNodesResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterInfoToNumInstrumentedNodesResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLClusterConnectionInfoTypeResolver<TParent = any> {
  ipAddress?: ClusterConnectionInfoToIpAddressResolver<TParent>;
  token?: ClusterConnectionInfoToTokenResolver<TParent>;
}

export interface ClusterConnectionInfoToIpAddressResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ClusterConnectionInfoToTokenResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLUserInviteTypeResolver<TParent = any> {
  email?: UserInviteToEmailResolver<TParent>;
  inviteLink?: UserInviteToInviteLinkResolver<TParent>;
}

export interface UserInviteToEmailResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface UserInviteToInviteLinkResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLMutationTypeResolver<TParent = any> {
  CreateCluster?: MutationToCreateClusterResolver<TParent>;
  UpdateVizierConfig?: MutationToUpdateVizierConfigResolver<TParent>;
  CreateDeploymentKey?: MutationToCreateDeploymentKeyResolver<TParent>;
  DeleteDeploymentKey?: MutationToDeleteDeploymentKeyResolver<TParent>;
  CreateAPIKey?: MutationToCreateAPIKeyResolver<TParent>;
  DeleteAPIKey?: MutationToDeleteAPIKeyResolver<TParent>;
  UpdateUserSettings?: MutationToUpdateUserSettingsResolver<TParent>;
  InviteUser?: MutationToInviteUserResolver<TParent>;
  UpdateUserPermissions?: MutationToUpdateUserPermissionsResolver<TParent>;
  UpdateOrgSettings?: MutationToUpdateOrgSettingsResolver<TParent>;
}

export interface MutationToCreateClusterResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToUpdateVizierConfigArgs {
  clusterID: string;
  passthroughEnabled?: boolean;
}
export interface MutationToUpdateVizierConfigResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToUpdateVizierConfigArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToCreateDeploymentKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToDeleteDeploymentKeyArgs {
  id: string;
}
export interface MutationToDeleteDeploymentKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToDeleteDeploymentKeyArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToCreateAPIKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToDeleteAPIKeyArgs {
  id: string;
}
export interface MutationToDeleteAPIKeyResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToDeleteAPIKeyArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToUpdateUserSettingsArgs {
  keys: Array<string>;
  values: Array<string>;
}
export interface MutationToUpdateUserSettingsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToUpdateUserSettingsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToInviteUserArgs {
  email: string;
  firstName: string;
  lastName: string;
}
export interface MutationToInviteUserResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToInviteUserArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToUpdateUserPermissionsArgs {
  userID: string;
  userPermissions: GQLEditableUserPermissions;
}
export interface MutationToUpdateUserPermissionsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToUpdateUserPermissionsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface MutationToUpdateOrgSettingsArgs {
  orgID: string;
  orgSettings: GQLEditableOrgSettings;
}
export interface MutationToUpdateOrgSettingsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: MutationToUpdateOrgSettingsArgs, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLCLIArtifactTypeResolver<TParent = any> {
  url?: CLIArtifactToUrlResolver<TParent>;
  sha256?: CLIArtifactToSha256Resolver<TParent>;
}

export interface CLIArtifactToUrlResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface CLIArtifactToSha256Resolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLArtifactsInfoTypeResolver<TParent = any> {
  items?: ArtifactsInfoToItemsResolver<TParent>;
}

export interface ArtifactsInfoToItemsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLArtifactTypeResolver<TParent = any> {
  version?: ArtifactToVersionResolver<TParent>;
  changelog?: ArtifactToChangelogResolver<TParent>;
  timestampMs?: ArtifactToTimestampMsResolver<TParent>;
}

export interface ArtifactToVersionResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ArtifactToChangelogResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ArtifactToTimestampMsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLLiveViewMetadataTypeResolver<TParent = any> {
  id?: LiveViewMetadataToIdResolver<TParent>;
  name?: LiveViewMetadataToNameResolver<TParent>;
  desc?: LiveViewMetadataToDescResolver<TParent>;
}

export interface LiveViewMetadataToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface LiveViewMetadataToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface LiveViewMetadataToDescResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLLiveViewContentsTypeResolver<TParent = any> {
  metadata?: LiveViewContentsToMetadataResolver<TParent>;
  pxlContents?: LiveViewContentsToPxlContentsResolver<TParent>;
  visJSON?: LiveViewContentsToVisJSONResolver<TParent>;
}

export interface LiveViewContentsToMetadataResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface LiveViewContentsToPxlContentsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface LiveViewContentsToVisJSONResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLScriptMetadataTypeResolver<TParent = any> {
  id?: ScriptMetadataToIdResolver<TParent>;
  name?: ScriptMetadataToNameResolver<TParent>;
  desc?: ScriptMetadataToDescResolver<TParent>;
  hasLiveView?: ScriptMetadataToHasLiveViewResolver<TParent>;
}

export interface ScriptMetadataToIdResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ScriptMetadataToNameResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ScriptMetadataToDescResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ScriptMetadataToHasLiveViewResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface GQLScriptContentsTypeResolver<TParent = any> {
  metadata?: ScriptContentsToMetadataResolver<TParent>;
  contents?: ScriptContentsToContentsResolver<TParent>;
}

export interface ScriptContentsToMetadataResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

export interface ScriptContentsToContentsResolver<TParent = any, TResult = any> {
  (parent: TParent, args: {}, context: any, info: GraphQLResolveInfo): TResult;
}

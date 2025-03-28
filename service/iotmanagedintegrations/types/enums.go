// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AbortCriteriaAction string

// Enum values for AbortCriteriaAction
const (
	AbortCriteriaActionCancel AbortCriteriaAction = "CANCEL"
)

// Values returns all known values for AbortCriteriaAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AbortCriteriaAction) Values() []AbortCriteriaAction {
	return []AbortCriteriaAction{
		"CANCEL",
	}
}

type AbortCriteriaFailureType string

// Enum values for AbortCriteriaFailureType
const (
	AbortCriteriaFailureTypeFailed   AbortCriteriaFailureType = "FAILED"
	AbortCriteriaFailureTypeRejected AbortCriteriaFailureType = "REJECTED"
	AbortCriteriaFailureTypeTimedOut AbortCriteriaFailureType = "TIMED_OUT"
	AbortCriteriaFailureTypeAll      AbortCriteriaFailureType = "ALL"
)

// Values returns all known values for AbortCriteriaFailureType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AbortCriteriaFailureType) Values() []AbortCriteriaFailureType {
	return []AbortCriteriaFailureType{
		"FAILED",
		"REJECTED",
		"TIMED_OUT",
		"ALL",
	}
}

type AuthMaterialType string

// Enum values for AuthMaterialType
const (
	AuthMaterialTypeWifiSetupQrBarCode AuthMaterialType = "WIFI_SETUP_QR_BAR_CODE"
	AuthMaterialTypeZwaveQrBarCode     AuthMaterialType = "ZWAVE_QR_BAR_CODE"
	AuthMaterialTypeZigbeeQrBarCode    AuthMaterialType = "ZIGBEE_QR_BAR_CODE"
)

// Values returns all known values for AuthMaterialType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthMaterialType) Values() []AuthMaterialType {
	return []AuthMaterialType{
		"WIFI_SETUP_QR_BAR_CODE",
		"ZWAVE_QR_BAR_CODE",
		"ZIGBEE_QR_BAR_CODE",
	}
}

type ConfigurationState string

// Enum values for ConfigurationState
const (
	ConfigurationStateEnabled          ConfigurationState = "ENABLED"
	ConfigurationStateUpdateInProgress ConfigurationState = "UPDATE_IN_PROGRESS"
	ConfigurationStateUpdateFailed     ConfigurationState = "UPDATE_FAILED"
)

// Values returns all known values for ConfigurationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationState) Values() []ConfigurationState {
	return []ConfigurationState{
		"ENABLED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type DeliveryDestinationType string

// Enum values for DeliveryDestinationType
const (
	DeliveryDestinationTypeKinesis DeliveryDestinationType = "KINESIS"
)

// Values returns all known values for DeliveryDestinationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryDestinationType) Values() []DeliveryDestinationType {
	return []DeliveryDestinationType{
		"KINESIS",
	}
}

type DeviceDiscoveryStatus string

// Enum values for DeviceDiscoveryStatus
const (
	DeviceDiscoveryStatusRunning   DeviceDiscoveryStatus = "RUNNING"
	DeviceDiscoveryStatusSucceeded DeviceDiscoveryStatus = "SUCCEEDED"
	DeviceDiscoveryStatusFailed    DeviceDiscoveryStatus = "FAILED"
	DeviceDiscoveryStatusTimedOut  DeviceDiscoveryStatus = "TIMED_OUT"
)

// Values returns all known values for DeviceDiscoveryStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceDiscoveryStatus) Values() []DeviceDiscoveryStatus {
	return []DeviceDiscoveryStatus{
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
	}
}

type DisconnectReasonValue string

// Enum values for DisconnectReasonValue
const (
	DisconnectReasonValueAuthError                 DisconnectReasonValue = "AUTH_ERROR"
	DisconnectReasonValueClientInitiatedDisconnect DisconnectReasonValue = "CLIENT_INITIATED_DISCONNECT"
	DisconnectReasonValueClientError               DisconnectReasonValue = "CLIENT_ERROR"
	DisconnectReasonValueConnectionLost            DisconnectReasonValue = "CONNECTION_LOST"
	DisconnectReasonValueDuplicateClientid         DisconnectReasonValue = "DUPLICATE_CLIENTID"
	DisconnectReasonValueForbiddenAccess           DisconnectReasonValue = "FORBIDDEN_ACCESS"
	DisconnectReasonValueMqttKeepAliveTimeout      DisconnectReasonValue = "MQTT_KEEP_ALIVE_TIMEOUT"
	DisconnectReasonValueServerError               DisconnectReasonValue = "SERVER_ERROR"
	DisconnectReasonValueServerInitiatedDisconnect DisconnectReasonValue = "SERVER_INITIATED_DISCONNECT"
	DisconnectReasonValueThrottled                 DisconnectReasonValue = "THROTTLED"
	DisconnectReasonValueWebsocketTtlExpiration    DisconnectReasonValue = "WEBSOCKET_TTL_EXPIRATION"
	DisconnectReasonValueCustomauthTtlExpiration   DisconnectReasonValue = "CUSTOMAUTH_TTL_EXPIRATION"
	DisconnectReasonValueUnknown                   DisconnectReasonValue = "UNKNOWN"
	DisconnectReasonValueNone                      DisconnectReasonValue = "NONE"
)

// Values returns all known values for DisconnectReasonValue. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DisconnectReasonValue) Values() []DisconnectReasonValue {
	return []DisconnectReasonValue{
		"AUTH_ERROR",
		"CLIENT_INITIATED_DISCONNECT",
		"CLIENT_ERROR",
		"CONNECTION_LOST",
		"DUPLICATE_CLIENTID",
		"FORBIDDEN_ACCESS",
		"MQTT_KEEP_ALIVE_TIMEOUT",
		"SERVER_ERROR",
		"SERVER_INITIATED_DISCONNECT",
		"THROTTLED",
		"WEBSOCKET_TTL_EXPIRATION",
		"CUSTOMAUTH_TTL_EXPIRATION",
		"UNKNOWN",
		"NONE",
	}
}

type DiscoveryAuthMaterialType string

// Enum values for DiscoveryAuthMaterialType
const (
	DiscoveryAuthMaterialTypeZwaveInstallCode DiscoveryAuthMaterialType = "ZWAVE_INSTALL_CODE"
)

// Values returns all known values for DiscoveryAuthMaterialType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DiscoveryAuthMaterialType) Values() []DiscoveryAuthMaterialType {
	return []DiscoveryAuthMaterialType{
		"ZWAVE_INSTALL_CODE",
	}
}

type DiscoveryType string

// Enum values for DiscoveryType
const (
	DiscoveryTypeZwave  DiscoveryType = "ZWAVE"
	DiscoveryTypeZigbee DiscoveryType = "ZIGBEE"
	DiscoveryTypeCloud  DiscoveryType = "CLOUD"
)

// Values returns all known values for DiscoveryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DiscoveryType) Values() []DiscoveryType {
	return []DiscoveryType{
		"ZWAVE",
		"ZIGBEE",
		"CLOUD",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeManagedIntegrationsDefaultEncryption EncryptionType = "MANAGED_INTEGRATIONS_DEFAULT_ENCRYPTION"
	EncryptionTypeCustomerKeyEncryption                EncryptionType = "CUSTOMER_KEY_ENCRYPTION"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"MANAGED_INTEGRATIONS_DEFAULT_ENCRYPTION",
		"CUSTOMER_KEY_ENCRYPTION",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeDeviceCommand        EventType = "DEVICE_COMMAND"
	EventTypeDeviceCommandRequest EventType = "DEVICE_COMMAND_REQUEST"
	EventTypeDeviceEvent          EventType = "DEVICE_EVENT"
	EventTypeDeviceLifeCycle      EventType = "DEVICE_LIFE_CYCLE"
	EventTypeDeviceState          EventType = "DEVICE_STATE"
	EventTypeDeviceOta            EventType = "DEVICE_OTA"
	EventTypeConnectorAssociation EventType = "CONNECTOR_ASSOCIATION"
	EventTypeConnectorErrorReport EventType = "CONNECTOR_ERROR_REPORT"
)

// Values returns all known values for EventType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"DEVICE_COMMAND",
		"DEVICE_COMMAND_REQUEST",
		"DEVICE_EVENT",
		"DEVICE_LIFE_CYCLE",
		"DEVICE_STATE",
		"DEVICE_OTA",
		"CONNECTOR_ASSOCIATION",
		"CONNECTOR_ERROR_REPORT",
	}
}

type HubNetworkMode string

// Enum values for HubNetworkMode
const (
	HubNetworkModeStandard             HubNetworkMode = "STANDARD"
	HubNetworkModeNetworkWideExclusion HubNetworkMode = "NETWORK_WIDE_EXCLUSION"
)

// Values returns all known values for HubNetworkMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HubNetworkMode) Values() []HubNetworkMode {
	return []HubNetworkMode{
		"STANDARD",
		"NETWORK_WIDE_EXCLUSION",
	}
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelError LogLevel = "ERROR"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
)

// Values returns all known values for LogLevel. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogLevel) Values() []LogLevel {
	return []LogLevel{
		"DEBUG",
		"ERROR",
		"INFO",
		"WARN",
	}
}

type OtaMechanism string

// Enum values for OtaMechanism
const (
	OtaMechanismPush OtaMechanism = "PUSH"
)

// Values returns all known values for OtaMechanism. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OtaMechanism) Values() []OtaMechanism {
	return []OtaMechanism{
		"PUSH",
	}
}

type OtaProtocol string

// Enum values for OtaProtocol
const (
	OtaProtocolHttp OtaProtocol = "HTTP"
)

// Values returns all known values for OtaProtocol. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OtaProtocol) Values() []OtaProtocol {
	return []OtaProtocol{
		"HTTP",
	}
}

type OtaStatus string

// Enum values for OtaStatus
const (
	OtaStatusInProgress         OtaStatus = "IN_PROGRESS"
	OtaStatusCanceled           OtaStatus = "CANCELED"
	OtaStatusCompleted          OtaStatus = "COMPLETED"
	OtaStatusDeletionInProgress OtaStatus = "DELETION_IN_PROGRESS"
	OtaStatusScheduled          OtaStatus = "SCHEDULED"
)

// Values returns all known values for OtaStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OtaStatus) Values() []OtaStatus {
	return []OtaStatus{
		"IN_PROGRESS",
		"CANCELED",
		"COMPLETED",
		"DELETION_IN_PROGRESS",
		"SCHEDULED",
	}
}

type OtaTaskExecutionStatus string

// Enum values for OtaTaskExecutionStatus
const (
	OtaTaskExecutionStatusQueued     OtaTaskExecutionStatus = "QUEUED"
	OtaTaskExecutionStatusInProgress OtaTaskExecutionStatus = "IN_PROGRESS"
	OtaTaskExecutionStatusSucceeded  OtaTaskExecutionStatus = "SUCCEEDED"
	OtaTaskExecutionStatusFailed     OtaTaskExecutionStatus = "FAILED"
	OtaTaskExecutionStatusTimedOut   OtaTaskExecutionStatus = "TIMED_OUT"
	OtaTaskExecutionStatusRejected   OtaTaskExecutionStatus = "REJECTED"
	OtaTaskExecutionStatusRemoved    OtaTaskExecutionStatus = "REMOVED"
	OtaTaskExecutionStatusCanceled   OtaTaskExecutionStatus = "CANCELED"
)

// Values returns all known values for OtaTaskExecutionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OtaTaskExecutionStatus) Values() []OtaTaskExecutionStatus {
	return []OtaTaskExecutionStatus{
		"QUEUED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
		"REJECTED",
		"REMOVED",
		"CANCELED",
	}
}

type OtaType string

// Enum values for OtaType
const (
	OtaTypeOneTime    OtaType = "ONE_TIME"
	OtaTypeContinuous OtaType = "CONTINUOUS"
)

// Values returns all known values for OtaType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OtaType) Values() []OtaType {
	return []OtaType{
		"ONE_TIME",
		"CONTINUOUS",
	}
}

type ProvisioningStatus string

// Enum values for ProvisioningStatus
const (
	ProvisioningStatusUnassociated     ProvisioningStatus = "UNASSOCIATED"
	ProvisioningStatusPreAssociated    ProvisioningStatus = "PRE_ASSOCIATED"
	ProvisioningStatusDiscovered       ProvisioningStatus = "DISCOVERED"
	ProvisioningStatusActivated        ProvisioningStatus = "ACTIVATED"
	ProvisioningStatusDeletionFailed   ProvisioningStatus = "DELETION_FAILED"
	ProvisioningStatusDeleteInProgress ProvisioningStatus = "DELETE_IN_PROGRESS"
	ProvisioningStatusIsolated         ProvisioningStatus = "ISOLATED"
	ProvisioningStatusDeleted          ProvisioningStatus = "DELETED"
)

// Values returns all known values for ProvisioningStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProvisioningStatus) Values() []ProvisioningStatus {
	return []ProvisioningStatus{
		"UNASSOCIATED",
		"PRE_ASSOCIATED",
		"DISCOVERED",
		"ACTIVATED",
		"DELETION_FAILED",
		"DELETE_IN_PROGRESS",
		"ISOLATED",
		"DELETED",
	}
}

type ProvisioningType string

// Enum values for ProvisioningType
const (
	ProvisioningTypeFleetProvisioning ProvisioningType = "FLEET_PROVISIONING"
	ProvisioningTypeJitr              ProvisioningType = "JITR"
)

// Values returns all known values for ProvisioningType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProvisioningType) Values() []ProvisioningType {
	return []ProvisioningType{
		"FLEET_PROVISIONING",
		"JITR",
	}
}

type RetryCriteriaFailureType string

// Enum values for RetryCriteriaFailureType
const (
	RetryCriteriaFailureTypeFailed   RetryCriteriaFailureType = "FAILED"
	RetryCriteriaFailureTypeTimedOut RetryCriteriaFailureType = "TIMED_OUT"
	RetryCriteriaFailureTypeAll      RetryCriteriaFailureType = "ALL"
)

// Values returns all known values for RetryCriteriaFailureType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RetryCriteriaFailureType) Values() []RetryCriteriaFailureType {
	return []RetryCriteriaFailureType{
		"FAILED",
		"TIMED_OUT",
		"ALL",
	}
}

type Role string

// Enum values for Role
const (
	RoleController Role = "CONTROLLER"
	RoleDevice     Role = "DEVICE"
)

// Values returns all known values for Role. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Role) Values() []Role {
	return []Role{
		"CONTROLLER",
		"DEVICE",
	}
}

type SchedulingConfigEndBehavior string

// Enum values for SchedulingConfigEndBehavior
const (
	SchedulingConfigEndBehaviorStopRollout SchedulingConfigEndBehavior = "STOP_ROLLOUT"
	SchedulingConfigEndBehaviorCancel      SchedulingConfigEndBehavior = "CANCEL"
	SchedulingConfigEndBehaviorForceCancel SchedulingConfigEndBehavior = "FORCE_CANCEL"
)

// Values returns all known values for SchedulingConfigEndBehavior. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchedulingConfigEndBehavior) Values() []SchedulingConfigEndBehavior {
	return []SchedulingConfigEndBehavior{
		"STOP_ROLLOUT",
		"CANCEL",
		"FORCE_CANCEL",
	}
}

type SchemaVersionFormat string

// Enum values for SchemaVersionFormat
const (
	SchemaVersionFormatAws       SchemaVersionFormat = "AWS"
	SchemaVersionFormatZcl       SchemaVersionFormat = "ZCL"
	SchemaVersionFormatConnector SchemaVersionFormat = "CONNECTOR"
)

// Values returns all known values for SchemaVersionFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaVersionFormat) Values() []SchemaVersionFormat {
	return []SchemaVersionFormat{
		"AWS",
		"ZCL",
		"CONNECTOR",
	}
}

type SchemaVersionType string

// Enum values for SchemaVersionType
const (
	SchemaVersionTypeCapability SchemaVersionType = "capability"
	SchemaVersionTypeDefinition SchemaVersionType = "definition"
)

// Values returns all known values for SchemaVersionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaVersionType) Values() []SchemaVersionType {
	return []SchemaVersionType{
		"capability",
		"definition",
	}
}

type SchemaVersionVisibility string

// Enum values for SchemaVersionVisibility
const (
	SchemaVersionVisibilityPublic  SchemaVersionVisibility = "PUBLIC"
	SchemaVersionVisibilityPrivate SchemaVersionVisibility = "PRIVATE"
)

// Values returns all known values for SchemaVersionVisibility. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaVersionVisibility) Values() []SchemaVersionVisibility {
	return []SchemaVersionVisibility{
		"PUBLIC",
		"PRIVATE",
	}
}

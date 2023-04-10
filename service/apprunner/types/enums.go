// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AutoScalingConfigurationStatus string

// Enum values for AutoScalingConfigurationStatus
const (
	AutoScalingConfigurationStatusActive   AutoScalingConfigurationStatus = "ACTIVE"
	AutoScalingConfigurationStatusInactive AutoScalingConfigurationStatus = "INACTIVE"
)

// Values returns all known values for AutoScalingConfigurationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AutoScalingConfigurationStatus) Values() []AutoScalingConfigurationStatus {
	return []AutoScalingConfigurationStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type CertificateValidationRecordStatus string

// Enum values for CertificateValidationRecordStatus
const (
	CertificateValidationRecordStatusPendingValidation CertificateValidationRecordStatus = "PENDING_VALIDATION"
	CertificateValidationRecordStatusSuccess           CertificateValidationRecordStatus = "SUCCESS"
	CertificateValidationRecordStatusFailed            CertificateValidationRecordStatus = "FAILED"
)

// Values returns all known values for CertificateValidationRecordStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CertificateValidationRecordStatus) Values() []CertificateValidationRecordStatus {
	return []CertificateValidationRecordStatus{
		"PENDING_VALIDATION",
		"SUCCESS",
		"FAILED",
	}
}

type ConfigurationSource string

// Enum values for ConfigurationSource
const (
	ConfigurationSourceRepository ConfigurationSource = "REPOSITORY"
	ConfigurationSourceApi        ConfigurationSource = "API"
)

// Values returns all known values for ConfigurationSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationSource) Values() []ConfigurationSource {
	return []ConfigurationSource{
		"REPOSITORY",
		"API",
	}
}

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusPendingHandshake ConnectionStatus = "PENDING_HANDSHAKE"
	ConnectionStatusAvailable        ConnectionStatus = "AVAILABLE"
	ConnectionStatusError            ConnectionStatus = "ERROR"
	ConnectionStatusDeleted          ConnectionStatus = "DELETED"
)

// Values returns all known values for ConnectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionStatus) Values() []ConnectionStatus {
	return []ConnectionStatus{
		"PENDING_HANDSHAKE",
		"AVAILABLE",
		"ERROR",
		"DELETED",
	}
}

type CustomDomainAssociationStatus string

// Enum values for CustomDomainAssociationStatus
const (
	CustomDomainAssociationStatusCreating                        CustomDomainAssociationStatus = "CREATING"
	CustomDomainAssociationStatusCreateFailed                    CustomDomainAssociationStatus = "CREATE_FAILED"
	CustomDomainAssociationStatusActive                          CustomDomainAssociationStatus = "ACTIVE"
	CustomDomainAssociationStatusDeleting                        CustomDomainAssociationStatus = "DELETING"
	CustomDomainAssociationStatusDeleteFailed                    CustomDomainAssociationStatus = "DELETE_FAILED"
	CustomDomainAssociationStatusPendingCertificateDnsValidation CustomDomainAssociationStatus = "PENDING_CERTIFICATE_DNS_VALIDATION"
	CustomDomainAssociationStatusBindingCertificate              CustomDomainAssociationStatus = "BINDING_CERTIFICATE"
)

// Values returns all known values for CustomDomainAssociationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CustomDomainAssociationStatus) Values() []CustomDomainAssociationStatus {
	return []CustomDomainAssociationStatus{
		"CREATING",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETING",
		"DELETE_FAILED",
		"PENDING_CERTIFICATE_DNS_VALIDATION",
		"BINDING_CERTIFICATE",
	}
}

type EgressType string

// Enum values for EgressType
const (
	EgressTypeDefault EgressType = "DEFAULT"
	EgressTypeVpc     EgressType = "VPC"
)

// Values returns all known values for EgressType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EgressType) Values() []EgressType {
	return []EgressType{
		"DEFAULT",
		"VPC",
	}
}

type HealthCheckProtocol string

// Enum values for HealthCheckProtocol
const (
	HealthCheckProtocolTcp  HealthCheckProtocol = "TCP"
	HealthCheckProtocolHttp HealthCheckProtocol = "HTTP"
)

// Values returns all known values for HealthCheckProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HealthCheckProtocol) Values() []HealthCheckProtocol {
	return []HealthCheckProtocol{
		"TCP",
		"HTTP",
	}
}

type ImageRepositoryType string

// Enum values for ImageRepositoryType
const (
	ImageRepositoryTypeEcr       ImageRepositoryType = "ECR"
	ImageRepositoryTypeEcrPublic ImageRepositoryType = "ECR_PUBLIC"
)

// Values returns all known values for ImageRepositoryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageRepositoryType) Values() []ImageRepositoryType {
	return []ImageRepositoryType{
		"ECR",
		"ECR_PUBLIC",
	}
}

type ObservabilityConfigurationStatus string

// Enum values for ObservabilityConfigurationStatus
const (
	ObservabilityConfigurationStatusActive   ObservabilityConfigurationStatus = "ACTIVE"
	ObservabilityConfigurationStatusInactive ObservabilityConfigurationStatus = "INACTIVE"
)

// Values returns all known values for ObservabilityConfigurationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ObservabilityConfigurationStatus) Values() []ObservabilityConfigurationStatus {
	return []ObservabilityConfigurationStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusPending            OperationStatus = "PENDING"
	OperationStatusInProgress         OperationStatus = "IN_PROGRESS"
	OperationStatusFailed             OperationStatus = "FAILED"
	OperationStatusSucceeded          OperationStatus = "SUCCEEDED"
	OperationStatusRollbackInProgress OperationStatus = "ROLLBACK_IN_PROGRESS"
	OperationStatusRollbackFailed     OperationStatus = "ROLLBACK_FAILED"
	OperationStatusRollbackSucceeded  OperationStatus = "ROLLBACK_SUCCEEDED"
)

// Values returns all known values for OperationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		"PENDING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"ROLLBACK_IN_PROGRESS",
		"ROLLBACK_FAILED",
		"ROLLBACK_SUCCEEDED",
	}
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeStartDeployment OperationType = "START_DEPLOYMENT"
	OperationTypeCreateService   OperationType = "CREATE_SERVICE"
	OperationTypePauseService    OperationType = "PAUSE_SERVICE"
	OperationTypeResumeService   OperationType = "RESUME_SERVICE"
	OperationTypeDeleteService   OperationType = "DELETE_SERVICE"
	OperationTypeUpdateService   OperationType = "UPDATE_SERVICE"
)

// Values returns all known values for OperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationType) Values() []OperationType {
	return []OperationType{
		"START_DEPLOYMENT",
		"CREATE_SERVICE",
		"PAUSE_SERVICE",
		"RESUME_SERVICE",
		"DELETE_SERVICE",
		"UPDATE_SERVICE",
	}
}

type ProviderType string

// Enum values for ProviderType
const (
	ProviderTypeGithub ProviderType = "GITHUB"
)

// Values returns all known values for ProviderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProviderType) Values() []ProviderType {
	return []ProviderType{
		"GITHUB",
	}
}

type Runtime string

// Enum values for Runtime
const (
	RuntimePython3    Runtime = "PYTHON_3"
	RuntimeNodejs12   Runtime = "NODEJS_12"
	RuntimeNodejs14   Runtime = "NODEJS_14"
	RuntimeCorretto8  Runtime = "CORRETTO_8"
	RuntimeCorretto11 Runtime = "CORRETTO_11"
	RuntimeNodejs16   Runtime = "NODEJS_16"
	RuntimeGo1        Runtime = "GO_1"
	RuntimeDotnet6    Runtime = "DOTNET_6"
	RuntimePhp81      Runtime = "PHP_81"
	RuntimeRuby31     Runtime = "RUBY_31"
)

// Values returns all known values for Runtime. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Runtime) Values() []Runtime {
	return []Runtime{
		"PYTHON_3",
		"NODEJS_12",
		"NODEJS_14",
		"CORRETTO_8",
		"CORRETTO_11",
		"NODEJS_16",
		"GO_1",
		"DOTNET_6",
		"PHP_81",
		"RUBY_31",
	}
}

type ServiceStatus string

// Enum values for ServiceStatus
const (
	ServiceStatusCreateFailed        ServiceStatus = "CREATE_FAILED"
	ServiceStatusRunning             ServiceStatus = "RUNNING"
	ServiceStatusDeleted             ServiceStatus = "DELETED"
	ServiceStatusDeleteFailed        ServiceStatus = "DELETE_FAILED"
	ServiceStatusPaused              ServiceStatus = "PAUSED"
	ServiceStatusOperationInProgress ServiceStatus = "OPERATION_IN_PROGRESS"
)

// Values returns all known values for ServiceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceStatus) Values() []ServiceStatus {
	return []ServiceStatus{
		"CREATE_FAILED",
		"RUNNING",
		"DELETED",
		"DELETE_FAILED",
		"PAUSED",
		"OPERATION_IN_PROGRESS",
	}
}

type SourceCodeVersionType string

// Enum values for SourceCodeVersionType
const (
	SourceCodeVersionTypeBranch SourceCodeVersionType = "BRANCH"
)

// Values returns all known values for SourceCodeVersionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceCodeVersionType) Values() []SourceCodeVersionType {
	return []SourceCodeVersionType{
		"BRANCH",
	}
}

type TracingVendor string

// Enum values for TracingVendor
const (
	TracingVendorAwsxray TracingVendor = "AWSXRAY"
)

// Values returns all known values for TracingVendor. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TracingVendor) Values() []TracingVendor {
	return []TracingVendor{
		"AWSXRAY",
	}
}

type VpcConnectorStatus string

// Enum values for VpcConnectorStatus
const (
	VpcConnectorStatusActive   VpcConnectorStatus = "ACTIVE"
	VpcConnectorStatusInactive VpcConnectorStatus = "INACTIVE"
)

// Values returns all known values for VpcConnectorStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VpcConnectorStatus) Values() []VpcConnectorStatus {
	return []VpcConnectorStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type VpcIngressConnectionStatus string

// Enum values for VpcIngressConnectionStatus
const (
	VpcIngressConnectionStatusAvailable       VpcIngressConnectionStatus = "AVAILABLE"
	VpcIngressConnectionStatusPendingCreation VpcIngressConnectionStatus = "PENDING_CREATION"
	VpcIngressConnectionStatusPendingUpdate   VpcIngressConnectionStatus = "PENDING_UPDATE"
	VpcIngressConnectionStatusPendingDeletion VpcIngressConnectionStatus = "PENDING_DELETION"
	VpcIngressConnectionStatusFailedCreation  VpcIngressConnectionStatus = "FAILED_CREATION"
	VpcIngressConnectionStatusFailedUpdate    VpcIngressConnectionStatus = "FAILED_UPDATE"
	VpcIngressConnectionStatusFailedDeletion  VpcIngressConnectionStatus = "FAILED_DELETION"
	VpcIngressConnectionStatusDeleted         VpcIngressConnectionStatus = "DELETED"
)

// Values returns all known values for VpcIngressConnectionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (VpcIngressConnectionStatus) Values() []VpcIngressConnectionStatus {
	return []VpcIngressConnectionStatus{
		"AVAILABLE",
		"PENDING_CREATION",
		"PENDING_UPDATE",
		"PENDING_DELETION",
		"FAILED_CREATION",
		"FAILED_UPDATE",
		"FAILED_DELETION",
		"DELETED",
	}
}

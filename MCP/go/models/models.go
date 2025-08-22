package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListSecurityPoliciesResponse represents the ListSecurityPoliciesResponse schema from the OpenAPI specification
type ListSecurityPoliciesResponse struct {
	Securitypolicynames interface{} `json:"SecurityPolicyNames"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ServiceMetadata represents the ServiceMetadata schema from the OpenAPI specification
type ServiceMetadata struct {
	Userdetails interface{} `json:"UserDetails"`
}

// IdentityProviderDetails represents the IdentityProviderDetails schema from the OpenAPI specification
type IdentityProviderDetails struct {
	Invocationrole interface{} `json:"InvocationRole,omitempty"`
	Sftpauthenticationmethods interface{} `json:"SftpAuthenticationMethods,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Directoryid interface{} `json:"DirectoryId,omitempty"`
	Function interface{} `json:"Function,omitempty"`
}

// CustomStepDetails represents the CustomStepDetails schema from the OpenAPI specification
type CustomStepDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Sourcefilelocation interface{} `json:"SourceFileLocation,omitempty"`
	Target interface{} `json:"Target,omitempty"`
	Timeoutseconds interface{} `json:"TimeoutSeconds,omitempty"`
}

// DescribedSecurityPolicy represents the DescribedSecurityPolicy schema from the OpenAPI specification
type DescribedSecurityPolicy struct {
	Sshciphers interface{} `json:"SshCiphers,omitempty"`
	Sshkexs interface{} `json:"SshKexs,omitempty"`
	Sshmacs interface{} `json:"SshMacs,omitempty"`
	Tlsciphers interface{} `json:"TlsCiphers,omitempty"`
	Fips interface{} `json:"Fips,omitempty"`
	Securitypolicyname interface{} `json:"SecurityPolicyName"`
}

// DescribeWorkflowRequest represents the DescribeWorkflowRequest schema from the OpenAPI specification
type DescribeWorkflowRequest struct {
	Workflowid interface{} `json:"WorkflowId"`
}

// StartServerRequest represents the StartServerRequest schema from the OpenAPI specification
type StartServerRequest struct {
	Serverid interface{} `json:"ServerId"`
}

// ListedWorkflow represents the ListedWorkflow schema from the OpenAPI specification
type ListedWorkflow struct {
	Workflowid interface{} `json:"WorkflowId,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// ListWorkflowsRequest represents the ListWorkflowsRequest schema from the OpenAPI specification
type ListWorkflowsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// As2ConnectorConfig represents the As2ConnectorConfig schema from the OpenAPI specification
type As2ConnectorConfig struct {
	Basicauthsecretid interface{} `json:"BasicAuthSecretId,omitempty"`
	Encryptionalgorithm interface{} `json:"EncryptionAlgorithm,omitempty"`
	Mdnresponse interface{} `json:"MdnResponse,omitempty"`
	Partnerprofileid interface{} `json:"PartnerProfileId,omitempty"`
	Signingalgorithm interface{} `json:"SigningAlgorithm,omitempty"`
	Mdnsigningalgorithm interface{} `json:"MdnSigningAlgorithm,omitempty"`
	Compression interface{} `json:"Compression,omitempty"`
	Localprofileid interface{} `json:"LocalProfileId,omitempty"`
	Messagesubject interface{} `json:"MessageSubject,omitempty"`
}

// ProtocolDetails represents the ProtocolDetails schema from the OpenAPI specification
type ProtocolDetails struct {
	Tlssessionresumptionmode interface{} `json:"TlsSessionResumptionMode,omitempty"`
	As2transports interface{} `json:"As2Transports,omitempty"`
	Passiveip interface{} `json:"PassiveIp,omitempty"`
	Setstatoption interface{} `json:"SetStatOption,omitempty"`
}

// S3InputFileLocation represents the S3InputFileLocation schema from the OpenAPI specification
type S3InputFileLocation struct {
	Key interface{} `json:"Key,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
}

// DeleteStepDetails represents the DeleteStepDetails schema from the OpenAPI specification
type DeleteStepDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Sourcefilelocation interface{} `json:"SourceFileLocation,omitempty"`
}

// SendWorkflowStepStateResponse represents the SendWorkflowStepStateResponse schema from the OpenAPI specification
type SendWorkflowStepStateResponse struct {
}

// DescribeProfileResponse represents the DescribeProfileResponse schema from the OpenAPI specification
type DescribeProfileResponse struct {
	Profile interface{} `json:"Profile"`
}

// UpdateServerRequest represents the UpdateServerRequest schema from the OpenAPI specification
type UpdateServerRequest struct {
	Endpointdetails interface{} `json:"EndpointDetails,omitempty"`
	Endpointtype interface{} `json:"EndpointType,omitempty"`
	Identityproviderdetails interface{} `json:"IdentityProviderDetails,omitempty"`
	Serverid interface{} `json:"ServerId"`
	Workflowdetails interface{} `json:"WorkflowDetails,omitempty"`
	Postauthenticationloginbanner interface{} `json:"PostAuthenticationLoginBanner,omitempty"`
	Preauthenticationloginbanner interface{} `json:"PreAuthenticationLoginBanner,omitempty"`
	Protocols interface{} `json:"Protocols,omitempty"`
	Structuredlogdestinations interface{} `json:"StructuredLogDestinations,omitempty"`
	Protocoldetails interface{} `json:"ProtocolDetails,omitempty"`
	Securitypolicyname interface{} `json:"SecurityPolicyName,omitempty"`
	Hostkey interface{} `json:"HostKey,omitempty"`
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
}

// ListSecurityPoliciesRequest represents the ListSecurityPoliciesRequest schema from the OpenAPI specification
type ListSecurityPoliciesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// HomeDirectoryMapEntry represents the HomeDirectoryMapEntry schema from the OpenAPI specification
type HomeDirectoryMapEntry struct {
	Target interface{} `json:"Target"`
	Entry interface{} `json:"Entry"`
}

// ListProfilesResponse represents the ListProfilesResponse schema from the OpenAPI specification
type ListProfilesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Profiles interface{} `json:"Profiles"`
}

// ListAgreementsResponse represents the ListAgreementsResponse schema from the OpenAPI specification
type ListAgreementsResponse struct {
	Agreements interface{} `json:"Agreements"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribedServer represents the DescribedServer schema from the OpenAPI specification
type DescribedServer struct {
	Protocoldetails interface{} `json:"ProtocolDetails,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
	Identityprovidertype interface{} `json:"IdentityProviderType,omitempty"`
	Preauthenticationloginbanner interface{} `json:"PreAuthenticationLoginBanner,omitempty"`
	Usercount interface{} `json:"UserCount,omitempty"`
	Hostkeyfingerprint interface{} `json:"HostKeyFingerprint,omitempty"`
	Endpointtype interface{} `json:"EndpointType,omitempty"`
	State interface{} `json:"State,omitempty"`
	Arn interface{} `json:"Arn"`
	Domain interface{} `json:"Domain,omitempty"`
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Structuredlogdestinations interface{} `json:"StructuredLogDestinations,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Identityproviderdetails interface{} `json:"IdentityProviderDetails,omitempty"`
	Postauthenticationloginbanner interface{} `json:"PostAuthenticationLoginBanner,omitempty"`
	Endpointdetails interface{} `json:"EndpointDetails,omitempty"`
	Serverid interface{} `json:"ServerId,omitempty"`
	Workflowdetails interface{} `json:"WorkflowDetails,omitempty"`
	Protocols interface{} `json:"Protocols,omitempty"`
	Securitypolicyname interface{} `json:"SecurityPolicyName,omitempty"`
}

// ListServersResponse represents the ListServersResponse schema from the OpenAPI specification
type ListServersResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Servers interface{} `json:"Servers"`
}

// DeleteCertificateRequest represents the DeleteCertificateRequest schema from the OpenAPI specification
type DeleteCertificateRequest struct {
	Certificateid interface{} `json:"CertificateId"`
}

// DeleteServerRequest represents the DeleteServerRequest schema from the OpenAPI specification
type DeleteServerRequest struct {
	Serverid interface{} `json:"ServerId"`
}

// ListHostKeysRequest represents the ListHostKeysRequest schema from the OpenAPI specification
type ListHostKeysRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverid interface{} `json:"ServerId"`
}

// DeleteSshPublicKeyRequest represents the DeleteSshPublicKeyRequest schema from the OpenAPI specification
type DeleteSshPublicKeyRequest struct {
	Username interface{} `json:"UserName"`
	Serverid interface{} `json:"ServerId"`
	Sshpublickeyid interface{} `json:"SshPublicKeyId"`
}

// ListedConnector represents the ListedConnector schema from the OpenAPI specification
type ListedConnector struct {
	Arn interface{} `json:"Arn,omitempty"`
	Connectorid interface{} `json:"ConnectorId,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// ListUsersResponse represents the ListUsersResponse schema from the OpenAPI specification
type ListUsersResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverid interface{} `json:"ServerId"`
	Users interface{} `json:"Users"`
}

// DescribeSecurityPolicyRequest represents the DescribeSecurityPolicyRequest schema from the OpenAPI specification
type DescribeSecurityPolicyRequest struct {
	Securitypolicyname interface{} `json:"SecurityPolicyName"`
}

// StartFileTransferRequest represents the StartFileTransferRequest schema from the OpenAPI specification
type StartFileTransferRequest struct {
	Remotedirectorypath interface{} `json:"RemoteDirectoryPath,omitempty"`
	Retrievefilepaths interface{} `json:"RetrieveFilePaths,omitempty"`
	Sendfilepaths interface{} `json:"SendFilePaths,omitempty"`
	Connectorid interface{} `json:"ConnectorId"`
	Localdirectorypath interface{} `json:"LocalDirectoryPath,omitempty"`
}

// ImportHostKeyRequest represents the ImportHostKeyRequest schema from the OpenAPI specification
type ImportHostKeyRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Hostkeybody interface{} `json:"HostKeyBody"`
	Serverid interface{} `json:"ServerId"`
}

// CreateServerRequest represents the CreateServerRequest schema from the OpenAPI specification
type CreateServerRequest struct {
	Structuredlogdestinations interface{} `json:"StructuredLogDestinations,omitempty"`
	Endpointtype interface{} `json:"EndpointType,omitempty"`
	Identityproviderdetails interface{} `json:"IdentityProviderDetails,omitempty"`
	Workflowdetails interface{} `json:"WorkflowDetails,omitempty"`
	Preauthenticationloginbanner interface{} `json:"PreAuthenticationLoginBanner,omitempty"`
	Protocoldetails interface{} `json:"ProtocolDetails,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
	Endpointdetails interface{} `json:"EndpointDetails,omitempty"`
	Hostkey interface{} `json:"HostKey,omitempty"`
	Domain interface{} `json:"Domain,omitempty"`
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Identityprovidertype interface{} `json:"IdentityProviderType,omitempty"`
	Postauthenticationloginbanner interface{} `json:"PostAuthenticationLoginBanner,omitempty"`
	Protocols interface{} `json:"Protocols,omitempty"`
	Securitypolicyname interface{} `json:"SecurityPolicyName,omitempty"`
}

// DeleteAgreementRequest represents the DeleteAgreementRequest schema from the OpenAPI specification
type DeleteAgreementRequest struct {
	Agreementid interface{} `json:"AgreementId"`
	Serverid interface{} `json:"ServerId"`
}

// LoggingConfiguration represents the LoggingConfiguration schema from the OpenAPI specification
type LoggingConfiguration struct {
	Loggroupname interface{} `json:"LogGroupName,omitempty"`
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
}

// UpdateConnectorResponse represents the UpdateConnectorResponse schema from the OpenAPI specification
type UpdateConnectorResponse struct {
	Connectorid interface{} `json:"ConnectorId"`
}

// DescribeHostKeyResponse represents the DescribeHostKeyResponse schema from the OpenAPI specification
type DescribeHostKeyResponse struct {
	Hostkey interface{} `json:"HostKey"`
}

// UpdateCertificateResponse represents the UpdateCertificateResponse schema from the OpenAPI specification
type UpdateCertificateResponse struct {
	Certificateid interface{} `json:"CertificateId"`
}

// PosixProfile represents the PosixProfile schema from the OpenAPI specification
type PosixProfile struct {
	Gid interface{} `json:"Gid"`
	Secondarygids interface{} `json:"SecondaryGids,omitempty"`
	Uid interface{} `json:"Uid"`
}

// DescribeAccessRequest represents the DescribeAccessRequest schema from the OpenAPI specification
type DescribeAccessRequest struct {
	Externalid interface{} `json:"ExternalId"`
	Serverid interface{} `json:"ServerId"`
}

// CreateWorkflowResponse represents the CreateWorkflowResponse schema from the OpenAPI specification
type CreateWorkflowResponse struct {
	Workflowid interface{} `json:"WorkflowId"`
}

// ListCertificatesRequest represents the ListCertificatesRequest schema from the OpenAPI specification
type ListCertificatesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// S3Tag represents the S3Tag schema from the OpenAPI specification
type S3Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ListUsersRequest represents the ListUsersRequest schema from the OpenAPI specification
type ListUsersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverid interface{} `json:"ServerId"`
}

// DescribedUser represents the DescribedUser schema from the OpenAPI specification
type DescribedUser struct {
	Policy interface{} `json:"Policy,omitempty"`
	Arn interface{} `json:"Arn"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
	Homedirectorymappings interface{} `json:"HomeDirectoryMappings,omitempty"`
	Posixprofile interface{} `json:"PosixProfile,omitempty"`
	Role interface{} `json:"Role,omitempty"`
	Sshpublickeys interface{} `json:"SshPublicKeys,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DecryptStepDetails represents the DecryptStepDetails schema from the OpenAPI specification
type DecryptStepDetails struct {
	TypeField interface{} `json:"Type"`
	Destinationfilelocation interface{} `json:"DestinationFileLocation"`
	Name interface{} `json:"Name,omitempty"`
	Overwriteexisting interface{} `json:"OverwriteExisting,omitempty"`
	Sourcefilelocation interface{} `json:"SourceFileLocation,omitempty"`
}

// ImportCertificateRequest represents the ImportCertificateRequest schema from the OpenAPI specification
type ImportCertificateRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Usage interface{} `json:"Usage"`
	Activedate interface{} `json:"ActiveDate,omitempty"`
	Certificate interface{} `json:"Certificate"`
	Certificatechain interface{} `json:"CertificateChain,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Inactivedate interface{} `json:"InactiveDate,omitempty"`
	Privatekey interface{} `json:"PrivateKey,omitempty"`
}

// WorkflowDetails represents the WorkflowDetails schema from the OpenAPI specification
type WorkflowDetails struct {
	Onpartialupload interface{} `json:"OnPartialUpload,omitempty"`
	Onupload interface{} `json:"OnUpload,omitempty"`
}

// UpdateConnectorRequest represents the UpdateConnectorRequest schema from the OpenAPI specification
type UpdateConnectorRequest struct {
	As2config interface{} `json:"As2Config,omitempty"`
	Connectorid interface{} `json:"ConnectorId"`
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Sftpconfig interface{} `json:"SftpConfig,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Accessrole interface{} `json:"AccessRole,omitempty"`
}

// DeleteProfileRequest represents the DeleteProfileRequest schema from the OpenAPI specification
type DeleteProfileRequest struct {
	Profileid interface{} `json:"ProfileId"`
}

// CreateAccessResponse represents the CreateAccessResponse schema from the OpenAPI specification
type CreateAccessResponse struct {
	Externalid interface{} `json:"ExternalId"`
	Serverid interface{} `json:"ServerId"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Arn interface{} `json:"Arn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EfsFileLocation represents the EfsFileLocation schema from the OpenAPI specification
type EfsFileLocation struct {
	Path interface{} `json:"Path,omitempty"`
	Filesystemid interface{} `json:"FileSystemId,omitempty"`
}

// TagStepDetails represents the TagStepDetails schema from the OpenAPI specification
type TagStepDetails struct {
	Sourcefilelocation interface{} `json:"SourceFileLocation,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// CreateConnectorResponse represents the CreateConnectorResponse schema from the OpenAPI specification
type CreateConnectorResponse struct {
	Connectorid interface{} `json:"ConnectorId"`
}

// ListAccessesResponse represents the ListAccessesResponse schema from the OpenAPI specification
type ListAccessesResponse struct {
	Serverid interface{} `json:"ServerId"`
	Accesses interface{} `json:"Accesses"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListedProfile represents the ListedProfile schema from the OpenAPI specification
type ListedProfile struct {
	Profileid interface{} `json:"ProfileId,omitempty"`
	Profiletype interface{} `json:"ProfileType,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	As2id interface{} `json:"As2Id,omitempty"`
}

// DescribeServerResponse represents the DescribeServerResponse schema from the OpenAPI specification
type DescribeServerResponse struct {
	Server interface{} `json:"Server"`
}

// ListedAccess represents the ListedAccess schema from the OpenAPI specification
type ListedAccess struct {
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
	Role interface{} `json:"Role,omitempty"`
	Externalid interface{} `json:"ExternalId,omitempty"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
}

// ListProfilesRequest represents the ListProfilesRequest schema from the OpenAPI specification
type ListProfilesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Profiletype interface{} `json:"ProfileType,omitempty"`
}

// DeleteWorkflowRequest represents the DeleteWorkflowRequest schema from the OpenAPI specification
type DeleteWorkflowRequest struct {
	Workflowid interface{} `json:"WorkflowId"`
}

// ImportCertificateResponse represents the ImportCertificateResponse schema from the OpenAPI specification
type ImportCertificateResponse struct {
	Certificateid interface{} `json:"CertificateId"`
}

// ListHostKeysResponse represents the ListHostKeysResponse schema from the OpenAPI specification
type ListHostKeysResponse struct {
	Hostkeys interface{} `json:"HostKeys"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverid interface{} `json:"ServerId"`
}

// CreateAccessRequest represents the CreateAccessRequest schema from the OpenAPI specification
type CreateAccessRequest struct {
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
	Posixprofile PosixProfile `json:"PosixProfile,omitempty"` // The full POSIX identity, including user ID (<code>Uid</code>), group ID (<code>Gid</code>), and any secondary groups IDs (<code>SecondaryGids</code>), that controls your users' access to your Amazon EFS file systems. The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	Role interface{} `json:"Role"`
	Serverid interface{} `json:"ServerId"`
	Externalid interface{} `json:"ExternalId"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Homedirectorymappings interface{} `json:"HomeDirectoryMappings,omitempty"`
}

// S3FileLocation represents the S3FileLocation schema from the OpenAPI specification
type S3FileLocation struct {
	Bucket interface{} `json:"Bucket,omitempty"`
	Etag interface{} `json:"Etag,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
}

// UpdateAccessRequest represents the UpdateAccessRequest schema from the OpenAPI specification
type UpdateAccessRequest struct {
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
	Posixprofile PosixProfile `json:"PosixProfile,omitempty"` // The full POSIX identity, including user ID (<code>Uid</code>), group ID (<code>Gid</code>), and any secondary groups IDs (<code>SecondaryGids</code>), that controls your users' access to your Amazon EFS file systems. The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	Role interface{} `json:"Role,omitempty"`
	Serverid interface{} `json:"ServerId"`
	Externalid interface{} `json:"ExternalId"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Homedirectorymappings interface{} `json:"HomeDirectoryMappings,omitempty"`
}

// CreateWorkflowRequest represents the CreateWorkflowRequest schema from the OpenAPI specification
type CreateWorkflowRequest struct {
	Steps interface{} `json:"Steps"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Onexceptionsteps interface{} `json:"OnExceptionSteps,omitempty"`
}

// DescribeConnectorRequest represents the DescribeConnectorRequest schema from the OpenAPI specification
type DescribeConnectorRequest struct {
	Connectorid interface{} `json:"ConnectorId"`
}

// UpdateAgreementResponse represents the UpdateAgreementResponse schema from the OpenAPI specification
type UpdateAgreementResponse struct {
	Agreementid interface{} `json:"AgreementId"`
}

// DescribeAgreementResponse represents the DescribeAgreementResponse schema from the OpenAPI specification
type DescribeAgreementResponse struct {
	Agreement interface{} `json:"Agreement"`
}

// UpdateProfileResponse represents the UpdateProfileResponse schema from the OpenAPI specification
type UpdateProfileResponse struct {
	Profileid interface{} `json:"ProfileId"`
}

// CopyStepDetails represents the CopyStepDetails schema from the OpenAPI specification
type CopyStepDetails struct {
	Sourcefilelocation interface{} `json:"SourceFileLocation,omitempty"`
	Destinationfilelocation interface{} `json:"DestinationFileLocation,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Overwriteexisting interface{} `json:"OverwriteExisting,omitempty"`
}

// DescribeExecutionRequest represents the DescribeExecutionRequest schema from the OpenAPI specification
type DescribeExecutionRequest struct {
	Workflowid interface{} `json:"WorkflowId"`
	Executionid interface{} `json:"ExecutionId"`
}

// UpdateUserRequest represents the UpdateUserRequest schema from the OpenAPI specification
type UpdateUserRequest struct {
	Username interface{} `json:"UserName"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Homedirectorymappings interface{} `json:"HomeDirectoryMappings,omitempty"`
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
	Posixprofile interface{} `json:"PosixProfile,omitempty"`
	Role interface{} `json:"Role,omitempty"`
	Serverid interface{} `json:"ServerId"`
}

// CreateProfileResponse represents the CreateProfileResponse schema from the OpenAPI specification
type CreateProfileResponse struct {
	Profileid interface{} `json:"ProfileId"`
}

// ExecutionStepResult represents the ExecutionStepResult schema from the OpenAPI specification
type ExecutionStepResult struct {
	ErrorField interface{} `json:"Error,omitempty"`
	Outputs interface{} `json:"Outputs,omitempty"`
	Steptype interface{} `json:"StepType,omitempty"`
}

// DescribedWorkflow represents the DescribedWorkflow schema from the OpenAPI specification
type DescribedWorkflow struct {
	Onexceptionsteps interface{} `json:"OnExceptionSteps,omitempty"`
	Steps interface{} `json:"Steps,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Workflowid interface{} `json:"WorkflowId,omitempty"`
	Arn interface{} `json:"Arn"`
	Description interface{} `json:"Description,omitempty"`
}

// UpdateAgreementRequest represents the UpdateAgreementRequest schema from the OpenAPI specification
type UpdateAgreementRequest struct {
	Partnerprofileid interface{} `json:"PartnerProfileId,omitempty"`
	Serverid interface{} `json:"ServerId"`
	Status interface{} `json:"Status,omitempty"`
	Accessrole interface{} `json:"AccessRole,omitempty"`
	Agreementid interface{} `json:"AgreementId"`
	Basedirectory interface{} `json:"BaseDirectory,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Localprofileid interface{} `json:"LocalProfileId,omitempty"`
}

// DescribeAccessResponse represents the DescribeAccessResponse schema from the OpenAPI specification
type DescribeAccessResponse struct {
	Access interface{} `json:"Access"`
	Serverid interface{} `json:"ServerId"`
}

// CreateUserRequest represents the CreateUserRequest schema from the OpenAPI specification
type CreateUserRequest struct {
	Serverid interface{} `json:"ServerId"`
	Policy interface{} `json:"Policy,omitempty"`
	Posixprofile interface{} `json:"PosixProfile,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Role interface{} `json:"Role"`
	Sshpublickeybody interface{} `json:"SshPublicKeyBody,omitempty"`
	Homedirectorymappings interface{} `json:"HomeDirectoryMappings,omitempty"`
	Username interface{} `json:"UserName"`
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
}

// SshPublicKey represents the SshPublicKey schema from the OpenAPI specification
type SshPublicKey struct {
	Dateimported interface{} `json:"DateImported"`
	Sshpublickeybody interface{} `json:"SshPublicKeyBody"`
	Sshpublickeyid interface{} `json:"SshPublicKeyId"`
}

// DescribedHostKey represents the DescribedHostKey schema from the OpenAPI specification
type DescribedHostKey struct {
	Arn interface{} `json:"Arn"`
	Dateimported interface{} `json:"DateImported,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Hostkeyfingerprint interface{} `json:"HostKeyFingerprint,omitempty"`
	Hostkeyid interface{} `json:"HostKeyId,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ExecutionResults represents the ExecutionResults schema from the OpenAPI specification
type ExecutionResults struct {
	Onexceptionsteps interface{} `json:"OnExceptionSteps,omitempty"`
	Steps interface{} `json:"Steps,omitempty"`
}

// DescribedAgreement represents the DescribedAgreement schema from the OpenAPI specification
type DescribedAgreement struct {
	Status interface{} `json:"Status,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Agreementid interface{} `json:"AgreementId,omitempty"`
	Basedirectory interface{} `json:"BaseDirectory,omitempty"`
	Accessrole interface{} `json:"AccessRole,omitempty"`
	Arn interface{} `json:"Arn"`
	Localprofileid interface{} `json:"LocalProfileId,omitempty"`
	Partnerprofileid interface{} `json:"PartnerProfileId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Serverid interface{} `json:"ServerId,omitempty"`
}

// DescribeAgreementRequest represents the DescribeAgreementRequest schema from the OpenAPI specification
type DescribeAgreementRequest struct {
	Agreementid interface{} `json:"AgreementId"`
	Serverid interface{} `json:"ServerId"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Arn interface{} `json:"Arn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// WorkflowDetail represents the WorkflowDetail schema from the OpenAPI specification
type WorkflowDetail struct {
	Executionrole interface{} `json:"ExecutionRole"`
	Workflowid interface{} `json:"WorkflowId"`
}

// ListedExecution represents the ListedExecution schema from the OpenAPI specification
type ListedExecution struct {
	Servicemetadata interface{} `json:"ServiceMetadata,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
	Initialfilelocation interface{} `json:"InitialFileLocation,omitempty"`
}

// DescribeUserRequest represents the DescribeUserRequest schema from the OpenAPI specification
type DescribeUserRequest struct {
	Username interface{} `json:"UserName"`
	Serverid interface{} `json:"ServerId"`
}

// ImportSshPublicKeyRequest represents the ImportSshPublicKeyRequest schema from the OpenAPI specification
type ImportSshPublicKeyRequest struct {
	Username interface{} `json:"UserName"`
	Serverid interface{} `json:"ServerId"`
	Sshpublickeybody interface{} `json:"SshPublicKeyBody"`
}

// CreateAgreementResponse represents the CreateAgreementResponse schema from the OpenAPI specification
type CreateAgreementResponse struct {
	Agreementid interface{} `json:"AgreementId"`
}

// SendWorkflowStepStateRequest represents the SendWorkflowStepStateRequest schema from the OpenAPI specification
type SendWorkflowStepStateRequest struct {
	Status interface{} `json:"Status"`
	Token interface{} `json:"Token"`
	Workflowid interface{} `json:"WorkflowId"`
	Executionid interface{} `json:"ExecutionId"`
}

// UpdateAccessResponse represents the UpdateAccessResponse schema from the OpenAPI specification
type UpdateAccessResponse struct {
	Externalid interface{} `json:"ExternalId"`
	Serverid interface{} `json:"ServerId"`
}

// DescribeServerRequest represents the DescribeServerRequest schema from the OpenAPI specification
type DescribeServerRequest struct {
	Serverid interface{} `json:"ServerId"`
}

// ListAgreementsRequest represents the ListAgreementsRequest schema from the OpenAPI specification
type ListAgreementsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverid interface{} `json:"ServerId"`
}

// DescribeCertificateRequest represents the DescribeCertificateRequest schema from the OpenAPI specification
type DescribeCertificateRequest struct {
	Certificateid interface{} `json:"CertificateId"`
}

// ListCertificatesResponse represents the ListCertificatesResponse schema from the OpenAPI specification
type ListCertificatesResponse struct {
	Certificates interface{} `json:"Certificates"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteConnectorRequest represents the DeleteConnectorRequest schema from the OpenAPI specification
type DeleteConnectorRequest struct {
	Connectorid interface{} `json:"ConnectorId"`
}

// ImportSshPublicKeyResponse represents the ImportSshPublicKeyResponse schema from the OpenAPI specification
type ImportSshPublicKeyResponse struct {
	Serverid interface{} `json:"ServerId"`
	Sshpublickeyid interface{} `json:"SshPublicKeyId"`
	Username interface{} `json:"UserName"`
}

// ListedHostKey represents the ListedHostKey schema from the OpenAPI specification
type ListedHostKey struct {
	Description interface{} `json:"Description,omitempty"`
	Fingerprint interface{} `json:"Fingerprint,omitempty"`
	Hostkeyid interface{} `json:"HostKeyId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn"`
	Dateimported interface{} `json:"DateImported,omitempty"`
}

// StopServerRequest represents the StopServerRequest schema from the OpenAPI specification
type StopServerRequest struct {
	Serverid interface{} `json:"ServerId"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Arn interface{} `json:"Arn"`
	Tags interface{} `json:"Tags"`
}

// ListAccessesRequest represents the ListAccessesRequest schema from the OpenAPI specification
type ListAccessesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverid interface{} `json:"ServerId"`
}

// ListExecutionsResponse represents the ListExecutionsResponse schema from the OpenAPI specification
type ListExecutionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Workflowid interface{} `json:"WorkflowId"`
	Executions interface{} `json:"Executions"`
}

// ListServersRequest represents the ListServersRequest schema from the OpenAPI specification
type ListServersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateUserResponse represents the CreateUserResponse schema from the OpenAPI specification
type CreateUserResponse struct {
	Serverid interface{} `json:"ServerId"`
	Username interface{} `json:"UserName"`
}

// ListedUser represents the ListedUser schema from the OpenAPI specification
type ListedUser struct {
	Role interface{} `json:"Role,omitempty"`
	Sshpublickeycount interface{} `json:"SshPublicKeyCount,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
	Arn interface{} `json:"Arn"`
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
}

// ExecutionError represents the ExecutionError schema from the OpenAPI specification
type ExecutionError struct {
	TypeField interface{} `json:"Type"`
	Message interface{} `json:"Message"`
}

// CreateAgreementRequest represents the CreateAgreementRequest schema from the OpenAPI specification
type CreateAgreementRequest struct {
	Accessrole interface{} `json:"AccessRole"`
	Basedirectory interface{} `json:"BaseDirectory"`
	Description interface{} `json:"Description,omitempty"`
	Localprofileid interface{} `json:"LocalProfileId"`
	Partnerprofileid interface{} `json:"PartnerProfileId"`
	Serverid interface{} `json:"ServerId"`
	Status interface{} `json:"Status,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// TestConnectionResponse represents the TestConnectionResponse schema from the OpenAPI specification
type TestConnectionResponse struct {
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Connectorid interface{} `json:"ConnectorId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// UpdateCertificateRequest represents the UpdateCertificateRequest schema from the OpenAPI specification
type UpdateCertificateRequest struct {
	Activedate interface{} `json:"ActiveDate,omitempty"`
	Certificateid interface{} `json:"CertificateId"`
	Description interface{} `json:"Description,omitempty"`
	Inactivedate interface{} `json:"InactiveDate,omitempty"`
}

// DescribedExecution represents the DescribedExecution schema from the OpenAPI specification
type DescribedExecution struct {
	Executionrole interface{} `json:"ExecutionRole,omitempty"`
	Initialfilelocation interface{} `json:"InitialFileLocation,omitempty"`
	Loggingconfiguration interface{} `json:"LoggingConfiguration,omitempty"`
	Posixprofile PosixProfile `json:"PosixProfile,omitempty"` // The full POSIX identity, including user ID (<code>Uid</code>), group ID (<code>Gid</code>), and any secondary groups IDs (<code>SecondaryGids</code>), that controls your users' access to your Amazon EFS file systems. The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	Results interface{} `json:"Results,omitempty"`
	Servicemetadata interface{} `json:"ServiceMetadata,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
}

// UpdateHostKeyRequest represents the UpdateHostKeyRequest schema from the OpenAPI specification
type UpdateHostKeyRequest struct {
	Description interface{} `json:"Description"`
	Hostkeyid interface{} `json:"HostKeyId"`
	Serverid interface{} `json:"ServerId"`
}

// SftpConnectorConfig represents the SftpConnectorConfig schema from the OpenAPI specification
type SftpConnectorConfig struct {
	Trustedhostkeys interface{} `json:"TrustedHostKeys,omitempty"`
	Usersecretid interface{} `json:"UserSecretId,omitempty"`
}

// DescribeWorkflowResponse represents the DescribeWorkflowResponse schema from the OpenAPI specification
type DescribeWorkflowResponse struct {
	Workflow interface{} `json:"Workflow"`
}

// UpdateHostKeyResponse represents the UpdateHostKeyResponse schema from the OpenAPI specification
type UpdateHostKeyResponse struct {
	Hostkeyid interface{} `json:"HostKeyId"`
	Serverid interface{} `json:"ServerId"`
}

// DeleteHostKeyRequest represents the DeleteHostKeyRequest schema from the OpenAPI specification
type DeleteHostKeyRequest struct {
	Hostkeyid interface{} `json:"HostKeyId"`
	Serverid interface{} `json:"ServerId"`
}

// UpdateProfileRequest represents the UpdateProfileRequest schema from the OpenAPI specification
type UpdateProfileRequest struct {
	Certificateids interface{} `json:"CertificateIds,omitempty"`
	Profileid interface{} `json:"ProfileId"`
}

// CreateConnectorRequest represents the CreateConnectorRequest schema from the OpenAPI specification
type CreateConnectorRequest struct {
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Sftpconfig interface{} `json:"SftpConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Url interface{} `json:"Url"`
	Accessrole interface{} `json:"AccessRole"`
	As2config interface{} `json:"As2Config,omitempty"`
}

// DescribedCertificate represents the DescribedCertificate schema from the OpenAPI specification
type DescribedCertificate struct {
	Certificateid interface{} `json:"CertificateId,omitempty"`
	Inactivedate interface{} `json:"InactiveDate,omitempty"`
	Notbeforedate interface{} `json:"NotBeforeDate,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Certificatechain interface{} `json:"CertificateChain,omitempty"`
	Notafterdate interface{} `json:"NotAfterDate,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Usage interface{} `json:"Usage,omitempty"`
	Arn interface{} `json:"Arn"`
	Description interface{} `json:"Description,omitempty"`
	Serial interface{} `json:"Serial,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Activedate interface{} `json:"ActiveDate,omitempty"`
}

// DescribeSecurityPolicyResponse represents the DescribeSecurityPolicyResponse schema from the OpenAPI specification
type DescribeSecurityPolicyResponse struct {
	Securitypolicy interface{} `json:"SecurityPolicy"`
}

// EndpointDetails represents the EndpointDetails schema from the OpenAPI specification
type EndpointDetails struct {
	Vpcid interface{} `json:"VpcId,omitempty"`
	Addressallocationids interface{} `json:"AddressAllocationIds,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Vpcendpointid interface{} `json:"VpcEndpointId,omitempty"`
}

// UserDetails represents the UserDetails schema from the OpenAPI specification
type UserDetails struct {
	Serverid interface{} `json:"ServerId"`
	Sessionid interface{} `json:"SessionId,omitempty"`
	Username interface{} `json:"UserName"`
}

// ListExecutionsRequest represents the ListExecutionsRequest schema from the OpenAPI specification
type ListExecutionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Workflowid interface{} `json:"WorkflowId"`
}

// DescribeCertificateResponse represents the DescribeCertificateResponse schema from the OpenAPI specification
type DescribeCertificateResponse struct {
	Certificate interface{} `json:"Certificate"`
}

// ListedServer represents the ListedServer schema from the OpenAPI specification
type ListedServer struct {
	Arn interface{} `json:"Arn"`
	Domain interface{} `json:"Domain,omitempty"`
	Endpointtype interface{} `json:"EndpointType,omitempty"`
	Identityprovidertype interface{} `json:"IdentityProviderType,omitempty"`
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Serverid interface{} `json:"ServerId,omitempty"`
	State interface{} `json:"State,omitempty"`
	Usercount interface{} `json:"UserCount,omitempty"`
}

// TestIdentityProviderRequest represents the TestIdentityProviderRequest schema from the OpenAPI specification
type TestIdentityProviderRequest struct {
	Userpassword interface{} `json:"UserPassword,omitempty"`
	Serverid interface{} `json:"ServerId"`
	Serverprotocol interface{} `json:"ServerProtocol,omitempty"`
	Sourceip interface{} `json:"SourceIp,omitempty"`
	Username interface{} `json:"UserName"`
}

// DeleteUserRequest represents the DeleteUserRequest schema from the OpenAPI specification
type DeleteUserRequest struct {
	Serverid interface{} `json:"ServerId"`
	Username interface{} `json:"UserName"`
}

// WorkflowStep represents the WorkflowStep schema from the OpenAPI specification
type WorkflowStep struct {
	Tagstepdetails interface{} `json:"TagStepDetails,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Copystepdetails interface{} `json:"CopyStepDetails,omitempty"`
	Customstepdetails interface{} `json:"CustomStepDetails,omitempty"`
	Decryptstepdetails interface{} `json:"DecryptStepDetails,omitempty"`
	Deletestepdetails interface{} `json:"DeleteStepDetails,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateProfileRequest represents the CreateProfileRequest schema from the OpenAPI specification
type CreateProfileRequest struct {
	As2id interface{} `json:"As2Id"`
	Certificateids interface{} `json:"CertificateIds,omitempty"`
	Profiletype interface{} `json:"ProfileType"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListedAgreement represents the ListedAgreement schema from the OpenAPI specification
type ListedAgreement struct {
	Localprofileid interface{} `json:"LocalProfileId,omitempty"`
	Partnerprofileid interface{} `json:"PartnerProfileId,omitempty"`
	Serverid interface{} `json:"ServerId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Agreementid interface{} `json:"AgreementId,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// CreateServerResponse represents the CreateServerResponse schema from the OpenAPI specification
type CreateServerResponse struct {
	Serverid interface{} `json:"ServerId"`
}

// DescribedAccess represents the DescribedAccess schema from the OpenAPI specification
type DescribedAccess struct {
	Homedirectory interface{} `json:"HomeDirectory,omitempty"`
	Homedirectorymappings interface{} `json:"HomeDirectoryMappings,omitempty"`
	Homedirectorytype interface{} `json:"HomeDirectoryType,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
	Posixprofile PosixProfile `json:"PosixProfile,omitempty"` // The full POSIX identity, including user ID (<code>Uid</code>), group ID (<code>Gid</code>), and any secondary groups IDs (<code>SecondaryGids</code>), that controls your users' access to your Amazon EFS file systems. The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	Role interface{} `json:"Role,omitempty"`
	Externalid interface{} `json:"ExternalId,omitempty"`
}

// TestIdentityProviderResponse represents the TestIdentityProviderResponse schema from the OpenAPI specification
type TestIdentityProviderResponse struct {
	Message interface{} `json:"Message,omitempty"`
	Response interface{} `json:"Response,omitempty"`
	Statuscode interface{} `json:"StatusCode"`
	Url interface{} `json:"Url"`
}

// DescribedProfile represents the DescribedProfile schema from the OpenAPI specification
type DescribedProfile struct {
	Arn interface{} `json:"Arn"`
	As2id interface{} `json:"As2Id,omitempty"`
	Certificateids interface{} `json:"CertificateIds,omitempty"`
	Profileid interface{} `json:"ProfileId,omitempty"`
	Profiletype interface{} `json:"ProfileType,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListedCertificate represents the ListedCertificate schema from the OpenAPI specification
type ListedCertificate struct {
	TypeField interface{} `json:"Type,omitempty"`
	Usage interface{} `json:"Usage,omitempty"`
	Activedate interface{} `json:"ActiveDate,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Certificateid interface{} `json:"CertificateId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Inactivedate interface{} `json:"InactiveDate,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeConnectorResponse represents the DescribeConnectorResponse schema from the OpenAPI specification
type DescribeConnectorResponse struct {
	Connector interface{} `json:"Connector"`
}

// DescribeHostKeyRequest represents the DescribeHostKeyRequest schema from the OpenAPI specification
type DescribeHostKeyRequest struct {
	Hostkeyid interface{} `json:"HostKeyId"`
	Serverid interface{} `json:"ServerId"`
}

// ListConnectorsRequest represents the ListConnectorsRequest schema from the OpenAPI specification
type ListConnectorsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeUserResponse represents the DescribeUserResponse schema from the OpenAPI specification
type DescribeUserResponse struct {
	User interface{} `json:"User"`
	Serverid interface{} `json:"ServerId"`
}

// DescribeExecutionResponse represents the DescribeExecutionResponse schema from the OpenAPI specification
type DescribeExecutionResponse struct {
	Execution interface{} `json:"Execution"`
	Workflowid interface{} `json:"WorkflowId"`
}

// FileLocation represents the FileLocation schema from the OpenAPI specification
type FileLocation struct {
	Efsfilelocation interface{} `json:"EfsFileLocation,omitempty"`
	S3filelocation interface{} `json:"S3FileLocation,omitempty"`
}

// DescribeProfileRequest represents the DescribeProfileRequest schema from the OpenAPI specification
type DescribeProfileRequest struct {
	Profileid interface{} `json:"ProfileId"`
}

// DescribedConnector represents the DescribedConnector schema from the OpenAPI specification
type DescribedConnector struct {
	Loggingrole interface{} `json:"LoggingRole,omitempty"`
	Sftpconfig interface{} `json:"SftpConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Accessrole interface{} `json:"AccessRole,omitempty"`
	Arn interface{} `json:"Arn"`
	As2config interface{} `json:"As2Config,omitempty"`
	Connectorid interface{} `json:"ConnectorId,omitempty"`
}

// DeleteAccessRequest represents the DeleteAccessRequest schema from the OpenAPI specification
type DeleteAccessRequest struct {
	Serverid interface{} `json:"ServerId"`
	Externalid interface{} `json:"ExternalId"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ListConnectorsResponse represents the ListConnectorsResponse schema from the OpenAPI specification
type ListConnectorsResponse struct {
	Connectors interface{} `json:"Connectors"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ImportHostKeyResponse represents the ImportHostKeyResponse schema from the OpenAPI specification
type ImportHostKeyResponse struct {
	Hostkeyid interface{} `json:"HostKeyId"`
	Serverid interface{} `json:"ServerId"`
}

// TestConnectionRequest represents the TestConnectionRequest schema from the OpenAPI specification
type TestConnectionRequest struct {
	Connectorid interface{} `json:"ConnectorId"`
}

// UpdateServerResponse represents the UpdateServerResponse schema from the OpenAPI specification
type UpdateServerResponse struct {
	Serverid interface{} `json:"ServerId"`
}

// UpdateUserResponse represents the UpdateUserResponse schema from the OpenAPI specification
type UpdateUserResponse struct {
	Serverid interface{} `json:"ServerId"`
	Username interface{} `json:"UserName"`
}

// ListWorkflowsResponse represents the ListWorkflowsResponse schema from the OpenAPI specification
type ListWorkflowsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Workflows interface{} `json:"Workflows"`
}

// InputFileLocation represents the InputFileLocation schema from the OpenAPI specification
type InputFileLocation struct {
	Efsfilelocation interface{} `json:"EfsFileLocation,omitempty"`
	S3filelocation interface{} `json:"S3FileLocation,omitempty"`
}

// StartFileTransferResponse represents the StartFileTransferResponse schema from the OpenAPI specification
type StartFileTransferResponse struct {
	Transferid interface{} `json:"TransferId"`
}

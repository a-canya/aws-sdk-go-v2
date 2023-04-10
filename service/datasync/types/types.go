// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Represents a single entry in a list (or array) of DataSync agents when you call
// the ListAgents (https://docs.aws.amazon.com/datasync/latest/userguide/API_ListAgents.html)
// operation.
type AgentListEntry struct {

	// The Amazon Resource Name (ARN) of a DataSync agent.
	AgentArn *string

	// The name of an agent.
	Name *string

	// The status of an agent. For more information, see DataSync agent statuses (https://docs.aws.amazon.com/datasync/latest/userguide/understand-agent-statuses.html)
	// .
	Status AgentStatus

	noSmithyDocumentSerde
}

// The subnet and security groups that DataSync uses to access your Amazon EFS
// file system.
type Ec2Config struct {

	// Specifies the Amazon Resource Names (ARNs) of the security groups associated
	// with an Amazon EFS file system's mount target.
	//
	// This member is required.
	SecurityGroupArns []string

	// Specifies the ARN of a subnet where DataSync creates the network interfaces (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// for managing traffic during your transfer. The subnet must be located:
	//   - In the same virtual private cloud (VPC) as the Amazon EFS file system.
	//   - In the same Availability Zone as at least one mount target for the Amazon
	//   EFS file system.
	// You don't need to specify a subnet that includes a file system mount target.
	//
	// This member is required.
	SubnetArn *string

	noSmithyDocumentSerde
}

// Specifies which files, folders, and objects to include or exclude when
// transferring files from source to destination.
type FilterRule struct {

	// The type of filter rule to apply. DataSync only supports the SIMPLE_PATTERN
	// rule type.
	FilterType FilterType

	// A single filter string that consists of the patterns to include or exclude. The
	// patterns are delimited by "|" (that is, a pipe), for example: /folder1|/folder2
	Value *string

	noSmithyDocumentSerde
}

// Specifies the data transfer protocol that DataSync uses to access your Amazon
// FSx file system.
type FsxProtocol struct {

	// Specifies the Network File System (NFS) protocol configuration that DataSync
	// uses to access your FSx for OpenZFS file system or FSx for ONTAP file system's
	// storage virtual machine (SVM).
	NFS *FsxProtocolNfs

	// Specifies the Server Message Block (SMB) protocol configuration that DataSync
	// uses to access your FSx for ONTAP file system's SVM.
	SMB *FsxProtocolSmb

	noSmithyDocumentSerde
}

// Specifies the Network File System (NFS) protocol configuration that DataSync
// uses to access your Amazon FSx for OpenZFS or Amazon FSx for NetApp ONTAP file
// system.
type FsxProtocolNfs struct {

	// Specifies how DataSync can access a location using the NFS protocol.
	MountOptions *NfsMountOptions

	noSmithyDocumentSerde
}

// Specifies the Server Message Block (SMB) protocol configuration that DataSync
// uses to access your Amazon FSx for NetApp ONTAP file system. For more
// information, see Accessing FSx for ONTAP file systems (https://docs.aws.amazon.com/datasync/latest/userguide/create-ontap-location.html#create-ontap-location-access)
// .
type FsxProtocolSmb struct {

	// Specifies the password of a user who has permission to access your SVM.
	//
	// This member is required.
	Password *string

	// Specifies a user name that can mount the location and access the files,
	// folders, and metadata that you need in the SVM. If you provide a user in your
	// Active Directory, note the following:
	//   - If you're using Directory Service for Microsoft Active Directory, the user
	//   must be a member of the Amazon Web Services Delegated FSx Administrators group.
	//   - If you're using a self-managed Active Directory, the user must be a member
	//   of either the Domain Admins group or a custom group that you specified for file
	//   system administration when you created your file system.
	// Make sure that the user has the permissions it needs to copy the data you want:
	//   - SE_TCB_NAME : Required to set object ownership and file metadata. With this
	//   privilege, you also can copy NTFS discretionary access lists (DACLs).
	//   - SE_SECURITY_NAME : May be needed to copy NTFS system access control lists
	//   (SACLs). This operation specifically requires the Windows privilege, which is
	//   granted to members of the Domain Admins group. If you configure your task to
	//   copy SACLs, make sure that the user has the required privileges. For information
	//   about copying SACLs, see Ownership and permissions-related options (https://docs.aws.amazon.com/datasync/latest/userguide/create-task.html#configure-ownership-and-permissions)
	//   .
	//
	// This member is required.
	User *string

	// Specifies the fully qualified domain name (FQDN) of the Microsoft Active
	// Directory that your storage virtual machine (SVM) belongs to.
	Domain *string

	// Specifies the version of the Server Message Block (SMB) protocol that DataSync
	// uses to access an SMB file server.
	MountOptions *SmbMountOptions

	noSmithyDocumentSerde
}

// The NameNode of the Hadoop Distributed File System (HDFS). The NameNode manages
// the file system's namespace. The NameNode performs operations such as opening,
// closing, and renaming files and directories. The NameNode contains the
// information to map blocks of data to the DataNodes.
type HdfsNameNode struct {

	// The hostname of the NameNode in the HDFS cluster. This value is the IP address
	// or Domain Name Service (DNS) name of the NameNode. An agent that's installed
	// on-premises uses this hostname to communicate with the NameNode in the network.
	//
	// This member is required.
	Hostname *string

	// The port that the NameNode uses to listen to client requests.
	//
	// This member is required.
	Port *int32

	noSmithyDocumentSerde
}

// Narrow down the list of resources returned by ListLocations . For example, to
// see all your Amazon S3 locations, create a filter using "Name": "LocationType" ,
// "Operator": "Equals" , and "Values": "S3" . For more information, see filtering
// resources (https://docs.aws.amazon.com/datasync/latest/userguide/query-resources.html)
// .
type LocationFilter struct {

	// The name of the filter being used. Each API call supports a list of filters
	// that are available for it (for example, LocationType for ListLocations ).
	//
	// This member is required.
	Name LocationFilterName

	// The operator that is used to compare filter values (for example, Equals or
	// Contains ).
	//
	// This member is required.
	Operator Operator

	// The values that you want to filter for. For example, you might want to display
	// only Amazon S3 locations.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Represents a single entry in a list of locations. LocationListEntry returns an
// array that contains a list of locations when the ListLocations (https://docs.aws.amazon.com/datasync/latest/userguide/API_ListLocations.html)
// operation is called.
type LocationListEntry struct {

	// The Amazon Resource Name (ARN) of the location. For Network File System (NFS)
	// or Amazon EFS, the location is the export path. For Amazon S3, the location is
	// the prefix path that you want to mount and use as the root of the location.
	LocationArn *string

	// Represents a list of URIs of a location. LocationUri returns an array that
	// contains a list of locations when the ListLocations (https://docs.aws.amazon.com/datasync/latest/userguide/API_ListLocations.html)
	// operation is called. Format: TYPE://GLOBAL_ID/SUBDIR . TYPE designates the type
	// of location (for example, nfs or s3 ). GLOBAL_ID is the globally unique
	// identifier of the resource that backs the location. An example for EFS is
	// us-east-2.fs-abcd1234 . An example for Amazon S3 is the bucket name, such as
	// myBucket . An example for NFS is a valid IPv4 address or a hostname that is
	// compliant with Domain Name Service (DNS). SUBDIR is a valid file system path,
	// delimited by forward slashes as is the *nix convention. For NFS and Amazon EFS,
	// it's the export path to mount the location. For Amazon S3, it's the prefix path
	// that you mount to and treat as the root of the location.
	LocationUri *string

	noSmithyDocumentSerde
}

// Specifies how DataSync can access a location using the NFS protocol.
type NfsMountOptions struct {

	// Specifies the NFS version that you want DataSync to use when mounting your NFS
	// share. If the server refuses to use the version specified, the task fails. You
	// can specify the following options:
	//   - AUTOMATIC (default): DataSync chooses NFS version 4.1.
	//   - NFS3 : Stateless protocol version that allows for asynchronous writes on the
	//   server.
	//   - NFSv4_0 : Stateful, firewall-friendly protocol version that supports
	//   delegations and pseudo file systems.
	//   - NFSv4_1 : Stateful protocol version that supports sessions, directory
	//   delegations, and parallel data processing. NFS version 4.1 also includes all
	//   features available in version 4.0.
	// DataSync currently only supports NFS version 3 with Amazon FSx for NetApp ONTAP
	// locations.
	Version NfsVersion

	noSmithyDocumentSerde
}

// A list of Amazon Resource Names (ARNs) of agents to use for a Network File
// System (NFS) location.
type OnPremConfig struct {

	// ARNs of the agents to use for an NFS location.
	//
	// This member is required.
	AgentArns []string

	noSmithyDocumentSerde
}

// Configures your DataSync task settings. These options include how DataSync
// handles files, objects, and their associated metadata. You also can specify how
// DataSync verifies data integrity, set bandwidth limits for your task, among
// other options. Each task setting has a default value. Unless you need to, you
// don't have to configure any of these Options before starting your task.
type Options struct {

	// Specifies whether to preserve metadata indicating the last time a file was read
	// or written to. If you set Atime to BEST_EFFORT , DataSync attempts to preserve
	// the original Atime attribute on all source files (that is, the version before
	// the PREPARING phase of the task execution). The behavior of Atime isn't fully
	// standard across platforms, so DataSync can only do this on a best-effort basis.
	// Default value: BEST_EFFORT BEST_EFFORT : Attempt to preserve the per-file Atime
	// value (recommended). NONE : Ignore Atime . If Atime is set to BEST_EFFORT ,
	// Mtime must be set to PRESERVE . If Atime is set to NONE , Mtime must also be
	// NONE .
	Atime Atime

	// Limits the bandwidth used by a DataSync task. For example, if you want DataSync
	// to use a maximum of 1 MB, set this value to 1048576 ( =1024*1024 ).
	BytesPerSecond *int64

	// Specifies the POSIX group ID (GID) of the file's owners. For more information,
	// see Metadata copied by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html#metadata-copied)
	// . Default value: INT_VALUE . This preserves the integer value of the ID.
	// INT_VALUE : Preserve the integer value of user ID (UID) and GID (recommended).
	// NONE : Ignore UID and GID.
	Gid Gid

	// Specifies the type of logs that DataSync publishes to a Amazon CloudWatch Logs
	// log group. To specify the log group, see CloudWatchLogGroupArn (https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateTask.html#DataSync-CreateTask-request-CloudWatchLogGroupArn)
	// . If you set LogLevel to OFF , no logs are published. BASIC publishes logs on
	// errors for individual files transferred. TRANSFER publishes logs for every file
	// or object that is transferred and integrity checked.
	LogLevel LogLevel

	// Specifies whether to preserve metadata indicating the last time that a file was
	// written to before the PREPARING phase of your task execution. This option is
	// required when you need to run the a task more than once. Default Value: PRESERVE
	// PRESERVE : Preserve original Mtime (recommended) NONE : Ignore Mtime . If Mtime
	// is set to PRESERVE , Atime must be set to BEST_EFFORT . If Mtime is set to NONE
	// , Atime must also be set to NONE .
	Mtime Mtime

	// Specifies whether object tags are preserved when transferring between object
	// storage systems. If you want your DataSync task to ignore object tags, specify
	// the NONE value. Default Value: PRESERVE
	ObjectTags ObjectTags

	// Specifies whether data at the destination location should be overwritten or
	// preserved. If set to NEVER , a destination file for example will not be replaced
	// by a source file (even if the destination file differs from the source file). If
	// you modify files in the destination and you sync the files, you can use this
	// value to protect against overwriting those changes. Some storage classes have
	// specific behaviors that can affect your Amazon S3 storage cost. For detailed
	// information, see Considerations when working with Amazon S3 storage classes in
	// DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes)
	// .
	OverwriteMode OverwriteMode

	// Specifies which users or groups can access a file for a specific purpose such
	// as reading, writing, or execution of the file. For more information, see
	// Metadata copied by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html#metadata-copied)
	// . Default value: PRESERVE PRESERVE : Preserve POSIX-style permissions
	// (recommended). NONE : Ignore permissions. DataSync can preserve extant
	// permissions of a source location.
	PosixPermissions PosixPermissions

	// Specifies whether files in the destination location that don't exist in the
	// source should be preserved. This option can affect your Amazon S3 storage cost.
	// If your task deletes objects, you might incur minimum storage duration charges
	// for certain storage classes. For detailed information, see Considerations when
	// working with Amazon S3 storage classes in DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes)
	// . Default value: PRESERVE PRESERVE : Ignore such destination files
	// (recommended). REMOVE : Delete destination files that aren’t present in the
	// source. If you set this parameter to REMOVE , you can't set TransferMode to ALL
	// . When you transfer all data, DataSync doesn't scan your destination location
	// and doesn't know what to delete.
	PreserveDeletedFiles PreserveDeletedFiles

	// Specifies whether DataSync should preserve the metadata of block and character
	// devices in the source location and recreate the files with that device name and
	// metadata on the destination. DataSync copies only the name and metadata of such
	// devices. DataSync can't copy the actual contents of these devices because
	// they're nonterminal and don't return an end-of-file (EOF) marker. Default value:
	// NONE NONE : Ignore special devices (recommended). PRESERVE : Preserve character
	// and block device metadata. This option currently isn't supported for Amazon EFS.
	PreserveDevices PreserveDevices

	// Specifies which components of the SMB security descriptor are copied from
	// source to destination objects. This value is only used for transfers between SMB
	// and Amazon FSx for Windows File Server locations or between two FSx for Windows
	// File Server locations. For more information, see how DataSync handles metadata (https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html)
	// . Default value: OWNER_DACL OWNER_DACL : For each copied object, DataSync copies
	// the following metadata:
	//   - The object owner.
	//   - NTFS discretionary access control lists (DACLs), which determine whether to
	//   grant access to an object. DataSync won't copy NTFS system access control lists
	//   (SACLs) with this option.
	// OWNER_DACL_SACL : For each copied object, DataSync copies the following
	// metadata:
	//   - The object owner.
	//   - NTFS discretionary access control lists (DACLs), which determine whether to
	//   grant access to an object.
	//   - SACLs, which are used by administrators to log attempts to access a secured
	//   object. Copying SACLs requires granting additional permissions to the Windows
	//   user that DataSync uses to access your SMB location. For information about
	//   choosing a user that ensures sufficient permissions to files, folders, and
	//   metadata, see user .
	// NONE : None of the SMB security descriptor components are copied. Destination
	// objects are owned by the user that was provided for accessing the destination
	// location. DACLs and SACLs are set based on the destination server’s
	// configuration.
	SecurityDescriptorCopyFlags SmbSecurityDescriptorCopyFlags

	// Specifies whether tasks should be queued before executing the tasks. The
	// default is ENABLED , which means the tasks will be queued. If you use the same
	// agent to run multiple tasks, you can enable the tasks to run in series. For more
	// information, see Queueing task executions (https://docs.aws.amazon.com/datasync/latest/userguide/run-task.html#queue-task-execution)
	// .
	TaskQueueing TaskQueueing

	// Determines whether DataSync transfers only the data and metadata that differ
	// between the source and the destination location or transfers all the content
	// from the source (without comparing what's in the destination). CHANGED :
	// DataSync copies only data or metadata that is new or different content from the
	// source location to the destination location. ALL : DataSync copies all source
	// location content to the destination (without comparing what's in the
	// destination).
	TransferMode TransferMode

	// Specifies the POSIX user ID (UID) of the file's owner. For more information,
	// see Metadata copied by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html#metadata-copied)
	// . Default value: INT_VALUE . This preserves the integer value of the ID.
	// INT_VALUE : Preserve the integer value of UID and group ID (GID) (recommended).
	// NONE : Ignore UID and GID.
	Uid Uid

	// Specifies how and when DataSync checks the integrity of your data during a
	// transfer. Default value: POINT_IN_TIME_CONSISTENT ONLY_FILES_TRANSFERRED
	// (recommended): DataSync calculates the checksum of transferred files and
	// metadata at the source location. At the end of the transfer, DataSync then
	// compares this checksum to the checksum calculated on those files at the
	// destination. We recommend this option when transferring to S3 Glacier Flexible
	// Retrieval or S3 Glacier Deep Archive storage classes. For more information, see
	// Storage class considerations with Amazon S3 locations (https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes)
	// . POINT_IN_TIME_CONSISTENT : At the end of the transfer, DataSync scans the
	// entire source and destination to verify that both locations are fully
	// synchronized. You can't use this option when transferring to S3 Glacier Flexible
	// Retrieval or S3 Glacier Deep Archive storage classes. For more information, see
	// Storage class considerations with Amazon S3 locations (https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes)
	// . NONE : DataSync doesn't run additional verification at the end of the
	// transfer. All data transmissions are still integrity-checked with checksum
	// verification during the transfer.
	VerifyMode VerifyMode

	noSmithyDocumentSerde
}

// The VPC endpoint, subnet, and security group that an agent uses to access IP
// addresses in a VPC (Virtual Private Cloud).
type PrivateLinkConfig struct {

	// The private endpoint that is configured for an agent that has access to IP
	// addresses in a PrivateLink (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html)
	// . An agent that is configured with this endpoint will not be accessible over the
	// public internet.
	PrivateLinkEndpoint *string

	// The Amazon Resource Names (ARNs) of the security groups that are configured for
	// the EC2 resource that hosts an agent activated in a VPC or an agent that has
	// access to a VPC endpoint.
	SecurityGroupArns []string

	// The Amazon Resource Names (ARNs) of the subnets that are configured for an
	// agent activated in a VPC or an agent that has access to a VPC endpoint.
	SubnetArns []string

	// The ID of the VPC endpoint that is configured for an agent. An agent that is
	// configured with a VPC endpoint will not be accessible over the public internet.
	VpcEndpointId *string

	noSmithyDocumentSerde
}

// The Quality of Protection (QOP) configuration specifies the Remote Procedure
// Call (RPC) and data transfer privacy settings configured on the Hadoop
// Distributed File System (HDFS) cluster.
type QopConfiguration struct {

	// The data transfer protection setting configured on the HDFS cluster. This
	// setting corresponds to your dfs.data.transfer.protection setting in the
	// hdfs-site.xml file on your Hadoop cluster.
	DataTransferProtection HdfsDataTransferProtection

	// The RPC protection setting configured on the HDFS cluster. This setting
	// corresponds to your hadoop.rpc.protection setting in your core-site.xml file on
	// your Hadoop cluster.
	RpcProtection HdfsRpcProtection

	noSmithyDocumentSerde
}

// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
// used to access an Amazon S3 bucket. For detailed information about using such a
// role, see Creating a Location for Amazon S3 in the DataSync User Guide.
type S3Config struct {

	// The ARN of the IAM role for accessing the S3 bucket.
	//
	// This member is required.
	BucketAccessRoleArn *string

	noSmithyDocumentSerde
}

// Specifies the version of the Server Message Block (SMB) protocol that DataSync
// uses to access an SMB file server.
type SmbMountOptions struct {

	// By default, DataSync automatically chooses an SMB protocol version based on
	// negotiation with your SMB file server. You also can configure DataSync to use a
	// specific SMB version, but we recommend doing this only if DataSync has trouble
	// negotiating with the SMB file server automatically. These are the following
	// options for configuring the SMB version:
	//   - AUTOMATIC (default): DataSync and the SMB file server negotiate a protocol
	//   version that they mutually support. (DataSync supports SMB versions 1.0 and
	//   later.) This is the recommended option. If you instead choose a specific version
	//   that your file server doesn't support, you may get an Operation Not Supported
	//   error.
	//   - SMB3 : Restricts the protocol negotiation to only SMB version 3.0.2.
	//   - SMB2 : Restricts the protocol negotiation to only SMB version 2.1.
	//   - SMB2_0 : Restricts the protocol negotiation to only SMB version 2.0.
	//   - SMB1 : Restricts the protocol negotiation to only SMB version 1.0. The SMB1
	//   option isn't available when creating an Amazon FSx for NetApp ONTAP location (https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateLocationFsxOntap.html)
	//   .
	Version SmbVersion

	noSmithyDocumentSerde
}

// A key-value pair representing a single tag that's been applied to an Amazon Web
// Services resource.
type TagListEntry struct {

	// The key for an Amazon Web Services resource tag.
	//
	// This member is required.
	Key *string

	// The value for an Amazon Web Services resource tag.
	Value *string

	noSmithyDocumentSerde
}

// Represents a single entry in a list of task executions. TaskExecutionListEntry
// returns an array that contains a list of specific invocations of a task when the
// ListTaskExecutions (https://docs.aws.amazon.com/datasync/latest/userguide/API_ListTaskExecutions.html)
// operation is called.
type TaskExecutionListEntry struct {

	// The status of a task execution.
	Status TaskExecutionStatus

	// The Amazon Resource Name (ARN) of the task that was executed.
	TaskExecutionArn *string

	noSmithyDocumentSerde
}

// Describes the detailed result of a TaskExecution operation. This result
// includes the time in milliseconds spent in each phase, the status of the task
// execution, and the errors encountered.
type TaskExecutionResultDetail struct {

	// Errors that DataSync encountered during execution of the task. You can use this
	// error code to help troubleshoot issues.
	ErrorCode *string

	// Detailed description of an error that was encountered during the task
	// execution. You can use this information to help troubleshoot issues.
	ErrorDetail *string

	// The total time in milliseconds that DataSync spent in the PREPARING phase.
	PrepareDuration *int64

	// The status of the PREPARING phase.
	PrepareStatus PhaseStatus

	// The total time in milliseconds that DataSync took to transfer the file from the
	// source to the destination location.
	TotalDuration *int64

	// The total time in milliseconds that DataSync spent in the TRANSFERRING phase.
	TransferDuration *int64

	// The status of the TRANSFERRING phase.
	TransferStatus PhaseStatus

	// The total time in milliseconds that DataSync spent in the VERIFYING phase.
	VerifyDuration *int64

	// The status of the VERIFYING phase.
	VerifyStatus PhaseStatus

	noSmithyDocumentSerde
}

// You can use API filters to narrow down the list of resources returned by
// ListTasks . For example, to retrieve all tasks on a source location, you can use
// ListTasks with filter name LocationId and Operator Equals with the ARN for the
// location. For more information, see filtering DataSync resources (https://docs.aws.amazon.com/datasync/latest/userguide/query-resources.html)
// .
type TaskFilter struct {

	// The name of the filter being used. Each API call supports a list of filters
	// that are available for it. For example, LocationId for ListTasks .
	//
	// This member is required.
	Name TaskFilterName

	// The operator that is used to compare filter values (for example, Equals or
	// Contains ).
	//
	// This member is required.
	Operator Operator

	// The values that you want to filter for. For example, you might want to display
	// only tasks for a specific destination location.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Represents a single entry in a list of tasks. TaskListEntry returns an array
// that contains a list of tasks when the ListTasks (https://docs.aws.amazon.com/datasync/latest/userguide/API_ListTasks.html)
// operation is called. A task includes the source and destination file systems to
// sync and the options to use for the tasks.
type TaskListEntry struct {

	// The name of the task.
	Name *string

	// The status of the task.
	Status TaskStatus

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	noSmithyDocumentSerde
}

// Specifies the schedule you want your task to use for repeated executions. For
// more information, see Schedule Expressions for Rules (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
// .
type TaskSchedule struct {

	// A cron expression that specifies when DataSync initiates a scheduled transfer
	// from a source to a destination location.
	//
	// This member is required.
	ScheduleExpression *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

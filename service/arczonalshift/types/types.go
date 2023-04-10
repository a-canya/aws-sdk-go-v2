// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A complex structure for a managed resource in an account. A managed resource is
// a Network Load Balancer or Application Load Balancer that has been registered
// with Route 53 ARC by Elastic Load Balancing. You can start a zonal shift in
// Route 53 ARC for a managed resource to temporarily move traffic for the resource
// away from an Availability Zone in an AWS Region. At this time, you can only
// start a zonal shift for Network Load Balancers and Application Load Balancers
// with cross-zone load balancing turned off.
type ManagedResourceSummary struct {

	// The Availability Zones that a resource is deployed in.
	//
	// This member is required.
	AvailabilityZones []string

	// The Amazon Resource Name (ARN) for the managed resource.
	Arn *string

	// The name of the managed resource.
	Name *string

	noSmithyDocumentSerde
}

// A complex structure that lists the zonal shifts for a managed resource and
// their statuses for the resource.
type ZonalShiftInResource struct {

	// An appliedStatus for a zonal shift for a resource can have one of two values:
	// APPLIED or NOT_APPLIED .
	//
	// This member is required.
	AppliedStatus AppliedStatus

	// The Availability Zone that traffic is moved away from for a resource when you
	// start a zonal shift. Until the zonal shift expires or you cancel it, traffic for
	// the resource is instead moved to other Availability Zones in the AWS Region.
	//
	// This member is required.
	AwayFrom *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. That is, a new comment overwrites
	// any existing comment string.
	//
	// This member is required.
	Comment *string

	// The expiry time (expiration time) for the zonal shift. A zonal shift is
	// temporary and must be set to expire when you start the zonal shift. You can
	// initially set a zonal shift to expire in a maximum of three days (72 hours).
	// However, you can update a zonal shift to set a new expiration at any time. When
	// you start a zonal shift, you specify how long you want it to be active, which
	// Route 53 ARC converts to an expiry time (expiration time). You can cancel a
	// zonal shift, for example, if you're ready to restore traffic to the Availability
	// Zone. Or you can update the zonal shift to specify another length of time to
	// expire in.
	//
	// This member is required.
	ExpiryTime *time.Time

	// The identifier for the resource to include in a zonal shift. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, you can only
	// start a zonal shift for Network Load Balancers and Application Load Balancers
	// with cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	// The time (UTC) when the zonal shift is started.
	//
	// This member is required.
	StartTime *time.Time

	// The identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	noSmithyDocumentSerde
}

// You start a zonal shift to temporarily move load balancer traffic away from an
// Availability Zone in a AWS Region. A zonal shift helps your application recover
// immediately, for example, from a developer's bad code deployment or from an AWS
// infrastructure failure in a single Availability Zone. You can start a zonal
// shift in Route 53 ARC only for managed resources in your account in an AWS
// Region. Supported AWS resources are automatically registered with Route 53 ARC.
// Zonal shifts are temporary. A zonal shift can be active for up to three days (72
// hours). When you start a zonal shift, you specify how long you want it to be
// active, which Amazon Route 53 Application Recovery Controller converts to an
// expiry time (expiration time). You can cancel a zonal shift, for example, if
// you're ready to restore traffic to the Availability Zone. Or you can extend the
// zonal shift by updating the expiration so the zonal shift is active longer.
type ZonalShiftSummary struct {

	// The Availability Zone that traffic is moved away from for a resource when you
	// start a zonal shift. Until the zonal shift expires or you cancel it, traffic for
	// the resource is instead moved to other Availability Zones in the AWS Region.
	//
	// This member is required.
	AwayFrom *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. That is, a new comment overwrites
	// any existing comment string.
	//
	// This member is required.
	Comment *string

	// The expiry time (expiration time) for the zonal shift. A zonal shift is
	// temporary and must be set to expire when you start the zonal shift. You can
	// initially set a zonal shift to expire in a maximum of three days (72 hours).
	// However, you can update a zonal shift to set a new expiration at any time. When
	// you start a zonal shift, you specify how long you want it to be active, which
	// Route 53 ARC converts to an expiry time (expiration time). You can cancel a
	// zonal shift, for example, if you're ready to restore traffic to the Availability
	// Zone. Or you can update the zonal shift to specify another length of time to
	// expire in.
	//
	// This member is required.
	ExpiryTime *time.Time

	// The identifier for the resource to include in a zonal shift. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, you can only
	// start a zonal shift for Network Load Balancers and Application Load Balancers
	// with cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	// The time (UTC) when the zonal shift is started.
	//
	// This member is required.
	StartTime *time.Time

	// A status for a zonal shift. The Status for a zonal shift can have one of the
	// following values:
	//   - ACTIVE: The zonal shift is started and active.
	//   - EXPIRED: The zonal shift has expired (the expiry time was exceeded).
	//   - CANCELED: The zonal shift was canceled.
	//
	// This member is required.
	Status ZonalShiftStatus

	// The identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

type CaseAttachmentAttributes struct {

	//
	//
	// This member is required.
	AttachmentId *string

	//
	//
	// This member is required.
	AttachmentStatus CaseAttachmentStatus

	//
	//
	// This member is required.
	CreatedDate *time.Time

	//
	//
	// This member is required.
	Creator *string

	//
	//
	// This member is required.
	FileName *string

	noSmithyDocumentSerde
}

type CaseEditItem struct {

	//
	Action *string

	//
	EventTimestamp *time.Time

	//
	Message *string

	//
	Principal *string

	noSmithyDocumentSerde
}

type GetMembershipAccountDetailError struct {

	//
	//
	// This member is required.
	AccountId *string

	//
	//
	// This member is required.
	Error *string

	//
	//
	// This member is required.
	Message *string

	noSmithyDocumentSerde
}

type GetMembershipAccountDetailItem struct {

	//
	AccountId *string

	//
	RelationshipStatus MembershipAccountRelationshipStatus

	//
	RelationshipType MembershipAccountRelationshipType

	noSmithyDocumentSerde
}

type ImpactedAwsRegion struct {

	//
	//
	// This member is required.
	Region AwsRegion

	noSmithyDocumentSerde
}

type IncidentResponder struct {

	//
	//
	// This member is required.
	Email *string

	//
	//
	// This member is required.
	JobTitle *string

	//
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type ListCasesItem struct {

	//
	//
	// This member is required.
	CaseId *string

	//
	CaseArn *string

	//
	CaseStatus CaseStatus

	//
	ClosedDate *time.Time

	//
	CreatedDate *time.Time

	//
	EngagementType EngagementType

	//
	LastUpdatedDate *time.Time

	//
	PendingAction PendingAction

	//
	ResolverType ResolverType

	//
	Title *string

	noSmithyDocumentSerde
}

type ListCommentsItem struct {

	//
	//
	// This member is required.
	CommentId *string

	//
	Body *string

	//
	CreatedDate *time.Time

	//
	Creator *string

	//
	LastUpdatedBy *string

	//
	LastUpdatedDate *time.Time

	noSmithyDocumentSerde
}

type ListMembershipItem struct {

	//
	//
	// This member is required.
	MembershipId *string

	//
	AccountId *string

	//
	MembershipArn *string

	//
	MembershipStatus MembershipStatus

	//
	Region AwsRegion

	noSmithyDocumentSerde
}

type OptInFeature struct {

	//
	//
	// This member is required.
	FeatureName OptInFeatureName

	//
	//
	// This member is required.
	IsEnabled *bool

	noSmithyDocumentSerde
}

type ThreatActorIp struct {

	//
	//
	// This member is required.
	IpAddress *string

	//
	UserAgent *string

	noSmithyDocumentSerde
}

type ValidationExceptionField struct {

	//
	//
	// This member is required.
	Message *string

	//
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type Watcher struct {

	//
	//
	// This member is required.
	Email *string

	//
	JobTitle *string

	//
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

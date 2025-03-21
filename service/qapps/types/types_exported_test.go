// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/qapps/types"
	"time"
)

func ExampleCard_outputUsage() {
	var union types.Card
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CardMemberFileUpload:
		_ = v.Value // Value is types.FileUploadCard

	case *types.CardMemberFormInput:
		_ = v.Value // Value is types.FormInputCard

	case *types.CardMemberQPlugin:
		_ = v.Value // Value is types.QPluginCard

	case *types.CardMemberQQuery:
		_ = v.Value // Value is types.QQueryCard

	case *types.CardMemberTextInput:
		_ = v.Value // Value is types.TextInputCard

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.QQueryCard
var _ *types.FileUploadCard
var _ *types.FormInputCard
var _ *types.TextInputCard
var _ *types.QPluginCard

func ExampleCardInput_outputUsage() {
	var union types.CardInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CardInputMemberFileUpload:
		_ = v.Value // Value is types.FileUploadCardInput

	case *types.CardInputMemberFormInput:
		_ = v.Value // Value is types.FormInputCardInput

	case *types.CardInputMemberQPlugin:
		_ = v.Value // Value is types.QPluginCardInput

	case *types.CardInputMemberQQuery:
		_ = v.Value // Value is types.QQueryCardInput

	case *types.CardInputMemberTextInput:
		_ = v.Value // Value is types.TextInputCardInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FileUploadCardInput
var _ *types.QQueryCardInput
var _ *types.QPluginCardInput
var _ *types.FormInputCardInput
var _ *types.TextInputCardInput

func ExampleDocumentAttributeValue_outputUsage() {
	var union types.DocumentAttributeValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DocumentAttributeValueMemberDateValue:
		_ = v.Value // Value is time.Time

	case *types.DocumentAttributeValueMemberLongValue:
		_ = v.Value // Value is int64

	case *types.DocumentAttributeValueMemberStringListValue:
		_ = v.Value // Value is []string

	case *types.DocumentAttributeValueMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *int64
var _ []string
var _ *time.Time

func ExamplePredictQAppInputOptions_outputUsage() {
	var union types.PredictQAppInputOptions
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PredictQAppInputOptionsMemberConversation:
		_ = v.Value // Value is []types.ConversationMessage

	case *types.PredictQAppInputOptionsMemberProblemStatement:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ []types.ConversationMessage

// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelExportTask struct {
}

func (*validateOpCancelExportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelExportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelExportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelExportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelImportTask struct {
}

func (*validateOpCancelImportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelImportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelQuery struct {
}

func (*validateOpCancelQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGraph struct {
}

func (*validateOpCreateGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGraphSnapshot struct {
}

func (*validateOpCreateGraphSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGraphSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGraphSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGraphSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGraphUsingImportTask struct {
}

func (*validateOpCreateGraphUsingImportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGraphUsingImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGraphUsingImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGraphUsingImportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePrivateGraphEndpoint struct {
}

func (*validateOpCreatePrivateGraphEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePrivateGraphEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePrivateGraphEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePrivateGraphEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteGraph struct {
}

func (*validateOpDeleteGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteGraphSnapshot struct {
}

func (*validateOpDeleteGraphSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteGraphSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteGraphSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteGraphSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePrivateGraphEndpoint struct {
}

func (*validateOpDeletePrivateGraphEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePrivateGraphEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePrivateGraphEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePrivateGraphEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExecuteQuery struct {
}

func (*validateOpExecuteQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExecuteQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExecuteQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExecuteQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExportTask struct {
}

func (*validateOpGetExportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetGraph struct {
}

func (*validateOpGetGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetGraphSnapshot struct {
}

func (*validateOpGetGraphSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetGraphSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetGraphSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetGraphSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetGraphSummary struct {
}

func (*validateOpGetGraphSummary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetGraphSummary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetGraphSummaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetGraphSummaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetImportTask struct {
}

func (*validateOpGetImportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetImportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPrivateGraphEndpoint struct {
}

func (*validateOpGetPrivateGraphEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPrivateGraphEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPrivateGraphEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPrivateGraphEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQuery struct {
}

func (*validateOpGetQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPrivateGraphEndpoints struct {
}

func (*validateOpListPrivateGraphEndpoints) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPrivateGraphEndpoints) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPrivateGraphEndpointsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPrivateGraphEndpointsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListQueries struct {
}

func (*validateOpListQueries) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListQueries) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListQueriesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListQueriesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResetGraph struct {
}

func (*validateOpResetGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResetGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResetGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResetGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRestoreGraphFromSnapshot struct {
}

func (*validateOpRestoreGraphFromSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRestoreGraphFromSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RestoreGraphFromSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRestoreGraphFromSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExportTask struct {
}

func (*validateOpStartExportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartImportTask struct {
}

func (*validateOpStartImportTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartImportTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateGraph struct {
}

func (*validateOpUpdateGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelExportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelExportTask{}, middleware.After)
}

func addOpCancelImportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelImportTask{}, middleware.After)
}

func addOpCancelQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelQuery{}, middleware.After)
}

func addOpCreateGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGraph{}, middleware.After)
}

func addOpCreateGraphSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGraphSnapshot{}, middleware.After)
}

func addOpCreateGraphUsingImportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGraphUsingImportTask{}, middleware.After)
}

func addOpCreatePrivateGraphEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePrivateGraphEndpoint{}, middleware.After)
}

func addOpDeleteGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteGraph{}, middleware.After)
}

func addOpDeleteGraphSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteGraphSnapshot{}, middleware.After)
}

func addOpDeletePrivateGraphEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePrivateGraphEndpoint{}, middleware.After)
}

func addOpExecuteQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExecuteQuery{}, middleware.After)
}

func addOpGetExportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExportTask{}, middleware.After)
}

func addOpGetGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetGraph{}, middleware.After)
}

func addOpGetGraphSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetGraphSnapshot{}, middleware.After)
}

func addOpGetGraphSummaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetGraphSummary{}, middleware.After)
}

func addOpGetImportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetImportTask{}, middleware.After)
}

func addOpGetPrivateGraphEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPrivateGraphEndpoint{}, middleware.After)
}

func addOpGetQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQuery{}, middleware.After)
}

func addOpListPrivateGraphEndpointsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPrivateGraphEndpoints{}, middleware.After)
}

func addOpListQueriesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListQueries{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpResetGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResetGraph{}, middleware.After)
}

func addOpRestoreGraphFromSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRestoreGraphFromSnapshot{}, middleware.After)
}

func addOpStartExportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartExportTask{}, middleware.After)
}

func addOpStartImportTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartImportTask{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateGraph{}, middleware.After)
}

func validateImportOptions(v types.ImportOptions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportOptions"}
	switch uv := v.(type) {
	case *types.ImportOptionsMemberNeptune:
		if err := validateNeptuneImportOptions(&uv.Value); err != nil {
			invalidParams.AddNested("[neptune]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNeptuneImportOptions(v *types.NeptuneImportOptions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NeptuneImportOptions"}
	if v.S3ExportPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3ExportPath"))
	}
	if v.S3ExportKmsKeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3ExportKmsKeyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVectorSearchConfiguration(v *types.VectorSearchConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VectorSearchConfiguration"}
	if v.Dimension == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Dimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelExportTaskInput(v *CancelExportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelExportTaskInput"}
	if v.TaskIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelImportTaskInput(v *CancelImportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelImportTaskInput"}
	if v.TaskIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelQueryInput(v *CancelQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelQueryInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGraphInput(v *CreateGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGraphInput"}
	if v.GraphName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphName"))
	}
	if v.VectorSearchConfiguration != nil {
		if err := validateVectorSearchConfiguration(v.VectorSearchConfiguration); err != nil {
			invalidParams.AddNested("VectorSearchConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ProvisionedMemory == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProvisionedMemory"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGraphSnapshotInput(v *CreateGraphSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGraphSnapshotInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.SnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGraphUsingImportTaskInput(v *CreateGraphUsingImportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGraphUsingImportTaskInput"}
	if v.GraphName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphName"))
	}
	if v.VectorSearchConfiguration != nil {
		if err := validateVectorSearchConfiguration(v.VectorSearchConfiguration); err != nil {
			invalidParams.AddNested("VectorSearchConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ImportOptions != nil {
		if err := validateImportOptions(v.ImportOptions); err != nil {
			invalidParams.AddNested("ImportOptions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePrivateGraphEndpointInput(v *CreatePrivateGraphEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePrivateGraphEndpointInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteGraphInput(v *DeleteGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteGraphInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.SkipSnapshot == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SkipSnapshot"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteGraphSnapshotInput(v *DeleteGraphSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteGraphSnapshotInput"}
	if v.SnapshotIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePrivateGraphEndpointInput(v *DeletePrivateGraphEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePrivateGraphEndpointInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.VpcId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExecuteQueryInput(v *ExecuteQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExecuteQueryInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if len(v.Language) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Language"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExportTaskInput(v *GetExportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExportTaskInput"}
	if v.TaskIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetGraphInput(v *GetGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetGraphInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetGraphSnapshotInput(v *GetGraphSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetGraphSnapshotInput"}
	if v.SnapshotIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetGraphSummaryInput(v *GetGraphSummaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetGraphSummaryInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetImportTaskInput(v *GetImportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetImportTaskInput"}
	if v.TaskIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPrivateGraphEndpointInput(v *GetPrivateGraphEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPrivateGraphEndpointInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.VpcId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueryInput(v *GetQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueryInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPrivateGraphEndpointsInput(v *ListPrivateGraphEndpointsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPrivateGraphEndpointsInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListQueriesInput(v *ListQueriesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListQueriesInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.MaxResults == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MaxResults"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResetGraphInput(v *ResetGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResetGraphInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.SkipSnapshot == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SkipSnapshot"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRestoreGraphFromSnapshotInput(v *RestoreGraphFromSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RestoreGraphFromSnapshotInput"}
	if v.SnapshotIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotIdentifier"))
	}
	if v.GraphName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExportTaskInput(v *StartExportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExportTaskInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if v.Destination == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Destination"))
	}
	if v.KmsKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KmsKeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartImportTaskInput(v *StartImportTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartImportTaskInput"}
	if v.ImportOptions != nil {
		if err := validateImportOptions(v.ImportOptions); err != nil {
			invalidParams.AddNested("ImportOptions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateGraphInput(v *UpdateGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateGraphInput"}
	if v.GraphIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

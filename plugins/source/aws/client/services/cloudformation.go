// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

//go:generate mockgen -package=mocks -destination=../mocks/cloudformation.go -source=cloudformation.go CloudformationClient
type CloudformationClient interface {
	DescribeAccountLimits(context.Context, *cloudformation.DescribeAccountLimitsInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeChangeSet(context.Context, *cloudformation.DescribeChangeSetInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetHooks(context.Context, *cloudformation.DescribeChangeSetHooksInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetHooksOutput, error)
	DescribeOrganizationsAccess(context.Context, *cloudformation.DescribeOrganizationsAccessInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeOrganizationsAccessOutput, error)
	DescribePublisher(context.Context, *cloudformation.DescribePublisherInput, ...func(*cloudformation.Options)) (*cloudformation.DescribePublisherOutput, error)
	DescribeStackDriftDetectionStatus(context.Context, *cloudformation.DescribeStackDriftDetectionStatusInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
	DescribeStackEvents(context.Context, *cloudformation.DescribeStackEventsInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackInstance(context.Context, *cloudformation.DescribeStackInstanceInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackResource(context.Context, *cloudformation.DescribeStackResourceInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceDrifts(context.Context, *cloudformation.DescribeStackResourceDriftsInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceDriftsOutput, error)
	DescribeStackResources(context.Context, *cloudformation.DescribeStackResourcesInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackSet(context.Context, *cloudformation.DescribeStackSetInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetOperation(context.Context, *cloudformation.DescribeStackSetOperationInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStacks(context.Context, *cloudformation.DescribeStacksInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error)
	DescribeType(context.Context, *cloudformation.DescribeTypeInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeOutput, error)
	DescribeTypeRegistration(context.Context, *cloudformation.DescribeTypeRegistrationInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeRegistrationOutput, error)
	GetStackPolicy(context.Context, *cloudformation.GetStackPolicyInput, ...func(*cloudformation.Options)) (*cloudformation.GetStackPolicyOutput, error)
	GetTemplate(context.Context, *cloudformation.GetTemplateInput, ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error)
	GetTemplateSummary(context.Context, *cloudformation.GetTemplateSummaryInput, ...func(*cloudformation.Options)) (*cloudformation.GetTemplateSummaryOutput, error)
	ListChangeSets(context.Context, *cloudformation.ListChangeSetsInput, ...func(*cloudformation.Options)) (*cloudformation.ListChangeSetsOutput, error)
	ListExports(context.Context, *cloudformation.ListExportsInput, ...func(*cloudformation.Options)) (*cloudformation.ListExportsOutput, error)
	ListImports(context.Context, *cloudformation.ListImportsInput, ...func(*cloudformation.Options)) (*cloudformation.ListImportsOutput, error)
	ListStackInstances(context.Context, *cloudformation.ListStackInstancesInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackInstancesOutput, error)
	ListStackResources(context.Context, *cloudformation.ListStackResourcesInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackResourcesOutput, error)
	ListStackSetOperationResults(context.Context, *cloudformation.ListStackSetOperationResultsInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperations(context.Context, *cloudformation.ListStackSetOperationsInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSets(context.Context, *cloudformation.ListStackSetsInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackSetsOutput, error)
	ListStacks(context.Context, *cloudformation.ListStacksInput, ...func(*cloudformation.Options)) (*cloudformation.ListStacksOutput, error)
	ListTypeRegistrations(context.Context, *cloudformation.ListTypeRegistrationsInput, ...func(*cloudformation.Options)) (*cloudformation.ListTypeRegistrationsOutput, error)
	ListTypeVersions(context.Context, *cloudformation.ListTypeVersionsInput, ...func(*cloudformation.Options)) (*cloudformation.ListTypeVersionsOutput, error)
	ListTypes(context.Context, *cloudformation.ListTypesInput, ...func(*cloudformation.Options)) (*cloudformation.ListTypesOutput, error)
}

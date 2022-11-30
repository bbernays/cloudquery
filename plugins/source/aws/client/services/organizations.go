// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

//go:generate mockgen -package=mocks -destination=../mocks/organizations.go -source=organizations.go OrganizationsClient
type OrganizationsClient interface {
	DescribeAccount(context.Context, *organizations.DescribeAccountInput, ...func(*organizations.Options)) (*organizations.DescribeAccountOutput, error)
	DescribeCreateAccountStatus(context.Context, *organizations.DescribeCreateAccountStatusInput, ...func(*organizations.Options)) (*organizations.DescribeCreateAccountStatusOutput, error)
	DescribeEffectivePolicy(context.Context, *organizations.DescribeEffectivePolicyInput, ...func(*organizations.Options)) (*organizations.DescribeEffectivePolicyOutput, error)
	DescribeHandshake(context.Context, *organizations.DescribeHandshakeInput, ...func(*organizations.Options)) (*organizations.DescribeHandshakeOutput, error)
	DescribeOrganization(context.Context, *organizations.DescribeOrganizationInput, ...func(*organizations.Options)) (*organizations.DescribeOrganizationOutput, error)
	DescribeOrganizationalUnit(context.Context, *organizations.DescribeOrganizationalUnitInput, ...func(*organizations.Options)) (*organizations.DescribeOrganizationalUnitOutput, error)
	DescribePolicy(context.Context, *organizations.DescribePolicyInput, ...func(*organizations.Options)) (*organizations.DescribePolicyOutput, error)
	DescribeResourcePolicy(context.Context, *organizations.DescribeResourcePolicyInput, ...func(*organizations.Options)) (*organizations.DescribeResourcePolicyOutput, error)
	ListAWSServiceAccessForOrganization(context.Context, *organizations.ListAWSServiceAccessForOrganizationInput, ...func(*organizations.Options)) (*organizations.ListAWSServiceAccessForOrganizationOutput, error)
	ListAccounts(context.Context, *organizations.ListAccountsInput, ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
	ListAccountsForParent(context.Context, *organizations.ListAccountsForParentInput, ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error)
	ListChildren(context.Context, *organizations.ListChildrenInput, ...func(*organizations.Options)) (*organizations.ListChildrenOutput, error)
	ListCreateAccountStatus(context.Context, *organizations.ListCreateAccountStatusInput, ...func(*organizations.Options)) (*organizations.ListCreateAccountStatusOutput, error)
	ListDelegatedAdministrators(context.Context, *organizations.ListDelegatedAdministratorsInput, ...func(*organizations.Options)) (*organizations.ListDelegatedAdministratorsOutput, error)
	ListDelegatedServicesForAccount(context.Context, *organizations.ListDelegatedServicesForAccountInput, ...func(*organizations.Options)) (*organizations.ListDelegatedServicesForAccountOutput, error)
	ListHandshakesForAccount(context.Context, *organizations.ListHandshakesForAccountInput, ...func(*organizations.Options)) (*organizations.ListHandshakesForAccountOutput, error)
	ListHandshakesForOrganization(context.Context, *organizations.ListHandshakesForOrganizationInput, ...func(*organizations.Options)) (*organizations.ListHandshakesForOrganizationOutput, error)
	ListOrganizationalUnitsForParent(context.Context, *organizations.ListOrganizationalUnitsForParentInput, ...func(*organizations.Options)) (*organizations.ListOrganizationalUnitsForParentOutput, error)
	ListParents(context.Context, *organizations.ListParentsInput, ...func(*organizations.Options)) (*organizations.ListParentsOutput, error)
	ListPolicies(context.Context, *organizations.ListPoliciesInput, ...func(*organizations.Options)) (*organizations.ListPoliciesOutput, error)
	ListPoliciesForTarget(context.Context, *organizations.ListPoliciesForTargetInput, ...func(*organizations.Options)) (*organizations.ListPoliciesForTargetOutput, error)
	ListRoots(context.Context, *organizations.ListRootsInput, ...func(*organizations.Options)) (*organizations.ListRootsOutput, error)
	ListTagsForResource(context.Context, *organizations.ListTagsForResourceInput, ...func(*organizations.Options)) (*organizations.ListTagsForResourceOutput, error)
	ListTargetsForPolicy(context.Context, *organizations.ListTargetsForPolicyInput, ...func(*organizations.Options)) (*organizations.ListTargetsForPolicyOutput, error)
}

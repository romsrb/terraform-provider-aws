// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package inspector

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	inspector_sdkv1 "github.com/aws/aws-sdk-go/service/inspector"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceRulesPackages,
			TypeName: "aws_inspector_rules_packages",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAssessmentTarget,
			TypeName: "aws_inspector_assessment_target",
		},
		{
			Factory:  ResourceAssessmentTemplate,
			TypeName: "aws_inspector_assessment_template",
			Name:     "Assessment Template",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceResourceGroup,
			TypeName: "aws_inspector_resource_group",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Inspector
}

func (p *servicePackage) Configure(ctx context.Context, config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*inspector_sdkv1.Inspector, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return inspector_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}

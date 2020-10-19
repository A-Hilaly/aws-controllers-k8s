// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package function

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/lambda/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.Lambda{}
	_ = &svcapitypes.Function{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetFunctionWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Code != nil {
		f0 := &svcapitypes.FunctionCode{}
		ko.Spec.Code = f0
	}
	if resp.Tags != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.Tags {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		ko.Spec.Tags = f3
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.FunctionName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetFunctionInput, error) {
	res := &svcsdk.GetFunctionInput{}

	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.ListFunctionsInput, error) {
	res := &svcsdk.ListFunctionsInput{}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateFunctionWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.CodeSha256 != nil {
		ko.Status.CodeSha256 = resp.CodeSha256
	}
	if resp.CodeSize != nil {
		ko.Status.CodeSize = resp.CodeSize
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.FunctionArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.FunctionArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.LastModified != nil {
		ko.Status.LastModified = resp.LastModified
	}
	if resp.LastUpdateStatus != nil {
		ko.Status.LastUpdateStatus = resp.LastUpdateStatus
	}
	if resp.LastUpdateStatusReason != nil {
		ko.Status.LastUpdateStatusReason = resp.LastUpdateStatusReason
	}
	if resp.LastUpdateStatusReasonCode != nil {
		ko.Status.LastUpdateStatusReasonCode = resp.LastUpdateStatusReasonCode
	}
	if resp.MasterArn != nil {
		ko.Status.MasterARN = resp.MasterArn
	}
	if resp.RevisionId != nil {
		ko.Status.RevisionID = resp.RevisionId
	}
	if resp.State != nil {
		ko.Status.State = resp.State
	}
	if resp.StateReason != nil {
		ko.Status.StateReason = resp.StateReason
	}
	if resp.StateReasonCode != nil {
		ko.Status.StateReasonCode = resp.StateReasonCode
	}
	if resp.Version != nil {
		ko.Status.Version = resp.Version
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateFunctionInput, error) {
	res := &svcsdk.CreateFunctionInput{}

	if r.ko.Spec.Code != nil {
		f0 := &svcsdk.FunctionCode{}
		if r.ko.Spec.Code.S3Bucket != nil {
			f0.SetS3Bucket(*r.ko.Spec.Code.S3Bucket)
		}
		if r.ko.Spec.Code.S3Key != nil {
			f0.SetS3Key(*r.ko.Spec.Code.S3Key)
		}
		if r.ko.Spec.Code.S3ObjectVersion != nil {
			f0.SetS3ObjectVersion(*r.ko.Spec.Code.S3ObjectVersion)
		}
		if r.ko.Spec.Code.ZipFile != nil {
			f0.SetZipFile(r.ko.Spec.Code.ZipFile)
		}
		res.SetCode(f0)
	}
	if r.ko.Spec.DeadLetterConfig != nil {
		f1 := &svcsdk.DeadLetterConfig{}
		if r.ko.Spec.DeadLetterConfig.TargetARN != nil {
			f1.SetTargetArn(*r.ko.Spec.DeadLetterConfig.TargetARN)
		}
		res.SetDeadLetterConfig(f1)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.Environment != nil {
		f3 := &svcsdk.Environment{}
		if r.ko.Spec.Environment.Variables != nil {
			f3f0 := map[string]*string{}
			for f3f0key, f3f0valiter := range r.ko.Spec.Environment.Variables {
				var f3f0val string
				f3f0val = *f3f0valiter
				f3f0[f3f0key] = &f3f0val
			}
			f3.SetVariables(f3f0)
		}
		res.SetEnvironment(f3)
	}
	if r.ko.Spec.FileSystemConfigs != nil {
		f4 := []*svcsdk.FileSystemConfig{}
		for _, f4iter := range r.ko.Spec.FileSystemConfigs {
			f4elem := &svcsdk.FileSystemConfig{}
			if f4iter.ARN != nil {
				f4elem.SetArn(*f4iter.ARN)
			}
			if f4iter.LocalMountPath != nil {
				f4elem.SetLocalMountPath(*f4iter.LocalMountPath)
			}
			f4 = append(f4, f4elem)
		}
		res.SetFileSystemConfigs(f4)
	}
	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Spec.Handler != nil {
		res.SetHandler(*r.ko.Spec.Handler)
	}
	if r.ko.Spec.KMSKeyARN != nil {
		res.SetKMSKeyArn(*r.ko.Spec.KMSKeyARN)
	}
	if r.ko.Spec.Layers != nil {
		f8 := []*string{}
		for _, f8iter := range r.ko.Spec.Layers {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		res.SetLayers(f8)
	}
	if r.ko.Spec.MemorySize != nil {
		res.SetMemorySize(*r.ko.Spec.MemorySize)
	}
	if r.ko.Spec.Publish != nil {
		res.SetPublish(*r.ko.Spec.Publish)
	}
	if r.ko.Spec.Role != nil {
		res.SetRole(*r.ko.Spec.Role)
	}
	if r.ko.Spec.Runtime != nil {
		res.SetRuntime(*r.ko.Spec.Runtime)
	}
	if r.ko.Spec.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range r.ko.Spec.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		res.SetTags(f13)
	}
	if r.ko.Spec.Timeout != nil {
		res.SetTimeout(*r.ko.Spec.Timeout)
	}
	if r.ko.Spec.TracingConfig != nil {
		f15 := &svcsdk.TracingConfig{}
		if r.ko.Spec.TracingConfig.Mode != nil {
			f15.SetMode(*r.ko.Spec.TracingConfig.Mode)
		}
		res.SetTracingConfig(f15)
	}
	if r.ko.Spec.VPCConfig != nil {
		f16 := &svcsdk.VpcConfig{}
		if r.ko.Spec.VPCConfig.SecurityGroupIDs != nil {
			f16f0 := []*string{}
			for _, f16f0iter := range r.ko.Spec.VPCConfig.SecurityGroupIDs {
				var f16f0elem string
				f16f0elem = *f16f0iter
				f16f0 = append(f16f0, &f16f0elem)
			}
			f16.SetSecurityGroupIds(f16f0)
		}
		if r.ko.Spec.VPCConfig.SubnetIDs != nil {
			f16f1 := []*string{}
			for _, f16f1iter := range r.ko.Spec.VPCConfig.SubnetIDs {
				var f16f1elem string
				f16f1elem = *f16f1iter
				f16f1 = append(f16f1, &f16f1elem)
			}
			f16.SetSubnetIds(f16f1)
		}
		res.SetVpcConfig(f16)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	return rm.customUpdateFunction(ctx, desired, latest, diffReporter)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteFunctionWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFunctionInput, error) {
	res := &svcsdk.DeleteFunctionInput{}

	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Function,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}
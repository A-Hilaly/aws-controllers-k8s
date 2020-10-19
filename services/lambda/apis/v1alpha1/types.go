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

package v1alpha1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

type AccountLimit struct {
	CodeSizeUnzipped *int64 `json:"codeSizeUnzipped,omitempty"`
	CodeSizeZipped *int64 `json:"codeSizeZipped,omitempty"`
	TotalCodeSize *int64 `json:"totalCodeSize,omitempty"`
}

type AccountUsage struct {
	FunctionCount *int64 `json:"functionCount,omitempty"`
	TotalCodeSize *int64 `json:"totalCodeSize,omitempty"`
}

type AliasRoutingConfiguration struct {
	AdditionalVersionWeights map[string]*float64 `json:"additionalVersionWeights,omitempty"`
}

type DeadLetterConfig struct {
	TargetARN *string `json:"targetARN,omitempty"`
}

type DestinationConfig struct {
	OnFailure *OnFailure `json:"onFailure,omitempty"`
	OnSuccess *OnSuccess `json:"onSuccess,omitempty"`
}

type Environment struct {
	Variables map[string]*string `json:"variables,omitempty"`
}

type EnvironmentError struct {
	ErrorCode *string `json:"errorCode,omitempty"`
	Message *string `json:"message,omitempty"`
}

type EnvironmentResponse struct {
	Error *EnvironmentError `json:"error,omitempty"`
	Variables map[string]*string `json:"variables,omitempty"`
}

type FileSystemConfig struct {
	ARN *string `json:"arn,omitempty"`
	LocalMountPath *string `json:"localMountPath,omitempty"`
}

type FunctionCode struct {
	S3Bucket *string `json:"s3Bucket,omitempty"`
	S3Key *string `json:"s3Key,omitempty"`
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty"`
	ZipFile []byte `json:"zipFile,omitempty"`
}

type FunctionCodeLocation struct {
	Location *string `json:"location,omitempty"`
	RepositoryType *string `json:"repositoryType,omitempty"`
}

type FunctionEventInvokeConfig struct {
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty"`
	FunctionARN *string `json:"functionARN,omitempty"`
	LastModified *metav1.Time `json:"lastModified,omitempty"`
}

type Layer struct {
	ARN *string `json:"arn,omitempty"`
	CodeSize *int64 `json:"codeSize,omitempty"`
}

type LayerVersionContentInput struct {
	S3Bucket *string `json:"s3Bucket,omitempty"`
	S3Key *string `json:"s3Key,omitempty"`
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty"`
	ZipFile []byte `json:"zipFile,omitempty"`
}

type LayerVersionContentOutput struct {
	CodeSha256 *string `json:"codeSha256,omitempty"`
	CodeSize *int64 `json:"codeSize,omitempty"`
	Location *string `json:"location,omitempty"`
}

type LayerVersionsListItem struct {
	CreatedDate *string `json:"createdDate,omitempty"`
	Description *string `json:"description,omitempty"`
	LayerVersionARN *string `json:"layerVersionARN,omitempty"`
}

type OnFailure struct {
	Destination *string `json:"destination,omitempty"`
}

type OnSuccess struct {
	Destination *string `json:"destination,omitempty"`
}

type ProvisionedConcurrencyConfigListItem struct {
	FunctionARN *string `json:"functionARN,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	StatusReason *string `json:"statusReason,omitempty"`
}

type TracingConfig struct {
	Mode *string `json:"mode,omitempty"`
}

type TracingConfigResponse struct {
	Mode *string `json:"mode,omitempty"`
}

type VPCConfig struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	SubnetIDs []*string `json:"subnetIDs,omitempty"`
}

type VPCConfigResponse struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	SubnetIDs []*string `json:"subnetIDs,omitempty"`
	VPCID *string `json:"vpcID,omitempty"`
}
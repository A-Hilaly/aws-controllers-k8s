{{- template "boilerplate" }}

// Code generated by ack-generate. DO NOT EDIT.

package {{ .APIVersion }}

import (
    "k8s.io/apimachinery/pkg/runtime/schema"
    "sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "{{ .APIGroup }}"
	Version = "{{ .APIVersion }}"
)

var (
    // GroupVersion is the API Group Version used to register the objects
    GroupVersion = schema.GroupVersion{Group: Group, Version: Version}

    // SchemeBuilder is used to add go types to the GroupVersionKind scheme
    SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

    // AddToScheme adds the types in this group-version to the given scheme.
    AddToScheme = SchemeBuilder.AddToScheme
)
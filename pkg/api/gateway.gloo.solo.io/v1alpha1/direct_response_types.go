package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DirectResponse contains configuration for defining direct response routes.
//
// +kubebuilder:object:root=true
// +kubebuilder:metadata:labels={app=gloo-gateway,app.kubernetes.io/name=gloo-gateway}
// +kubebuilder:resource:categories=gloo-gateway,shortName=dr
// +kubebuilder:subresource:status
type DirectResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectResponseSpec   `json:"spec,omitempty"`
	Status DirectResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type DirectResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectResponse `json:"items"`
}

// DirectResponseSpec describes the desired state of a DirectResponse.
type DirectResponseSpec struct {
	// StatusCode defines the HTTP status code to return for this route.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=200
	// +kubebuilder:validation:Maximum=599
	StatusCode uint32 `json:"status"`
	// Body defines the content to be returned in the HTTP response body.
	// The maximum length of the body is restricted to prevent excessively large responses.
	//
	// +kubebuilder:validation:MaxLength=4096
	// +kubebuilder:validation:Optional
	Body string `json:"body,omitempty"`
}

// DirectResponseStatus defines the observed state of a DirectResponse.
type DirectResponseStatus struct{}

// GetStatus returns the HTTP status code to return for this route.
func (in *DirectResponse) GetStatusCode() uint32 {
	if in == nil {
		return 0
	}
	return in.Spec.StatusCode
}

// GetBody returns the content to be returned in the HTTP response body.
func (in *DirectResponse) GetBody() string {
	if in == nil {
		return ""
	}
	return in.Spec.Body
}

func init() {
	SchemeBuilder.Register(&DirectResponse{}, &DirectResponseList{})
}

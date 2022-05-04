package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TFTemplateSpec defines the desired state of TFTemplate
// TODO: TemplateSpec -> Extract into a struct
type TFTemplateSpec struct {
	Description       string             `json:"description,omitempty"`
	Params            []TemplateParam    `json:"params,omitempty"` // Described above
	ResourceTemplates []ResourceTemplate `json:"resourcetemplates,omitempty"`
}

// Param is a parameter that can be templated into a struct.
type TemplateParam struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Required    bool     `json:"required,omitempty"`
	Options     []string `json:"options,omitempty"`
}

//+kubebuilder:pruning:PreserveUnknownFields

// ResourceTemplate describes a resource to create.
type ResourceTemplate struct {
	runtime.RawExtension `json:",inline"`
}

// TFTemplateStatus defines the observed state of TFTemplate
type TFTemplateStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TFTemplate is the Schema for the TFTemplates API
type TFTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TFTemplateSpec   `json:"spec,omitempty"`
	Status TFTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TFTemplateList contains a list of TFTemplate
type TFTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TFTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TFTemplate{}, &TFTemplateList{})
}

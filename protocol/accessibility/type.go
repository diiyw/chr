package accessibility

import (
	"github.com/diiyw/cuto/protocol/dom"
)

// Unique accessibility node identifier.
type AXNodeId string

// Enum of possible property types.
type AXValueType string

// Enum of possible property sources.
type AXValueSourceType string

// Enum of possible native property sources (as a subtype of a particular AXValueSourceType).
type AXValueNativeSourceType string

// A single source for a computed AX property.
type AXValueSource  struct {

	// What type of source this is.
	Type	AXValueSourceType	`json:"type"`

	// The value of this property source.
	Value	AXValue	`json:"value,omitempty"`

	// The name of the relevant attribute, if any.
	Attribute	string	`json:"attribute,omitempty"`

	// The value of the relevant attribute, if any.
	AttributeValue	AXValue	`json:"attributeValue,omitempty"`

	// Whether this source is superseded by a higher priority source.
	Superseded	bool	`json:"superseded,omitempty"`

	// The native markup source for this value, e.g. a <label> element.
	NativeSource	AXValueNativeSourceType	`json:"nativeSource,omitempty"`

	// The value, such as a node or node list, of the native source.
	NativeSourceValue	AXValue	`json:"nativeSourceValue,omitempty"`

	// Whether the value for this property is invalid.
	Invalid	bool	`json:"invalid,omitempty"`

	// Reason for the value being invalid, if it is.
	InvalidReason	string	`json:"invalidReason,omitempty"`
}

// 
type AXRelatedNode  struct {

	// The BackendNodeId of the related DOM node.
	BackendDOMNodeId	dom.BackendNodeId	`json:"backendDOMNodeId"`

	// The IDRef value provided, if any.
	Idref	string	`json:"idref,omitempty"`

	// The text alternative of this node in the current context.
	Text	string	`json:"text,omitempty"`
}

// 
type AXProperty  struct {

	// The name of this property.
	Name	AXPropertyName	`json:"name"`

	// The value of this property.
	Value	AXValue	`json:"value"`
}

// A single computed AX property.
type AXValue  struct {

	// The type of this value.
	Type	AXValueType	`json:"type"`

	// The computed value of this property.
	Value	interface{}	`json:"value,omitempty"`

	// One or more related nodes, if applicable.
	RelatedNodes	[]*AXRelatedNode	`json:"relatedNodes,omitempty"`

	// The sources which contributed to the computation of this property.
	Sources	[]*AXValueSource	`json:"sources,omitempty"`
}

// Values of AXProperty name:
	// - from 'busy' to 'roledescription': states which apply to every AX node
	// - from 'live' to 'root': attributes which apply to nodes in live regions
	// - from 'autocomplete' to 'valuetext': attributes which apply to widgets
	// - from 'checked' to 'selected': states which apply to widgets
	// - from 'activedescendant' to 'owns' - relationships between elements other than parent/child/sibling.
type AXPropertyName string

// A node in the accessibility tree.
type AXNode  struct {

	// Unique identifier for this node.
	NodeId	AXNodeId	`json:"nodeId"`

	// Whether this node is ignored for accessibility
	Ignored	bool	`json:"ignored"`

	// Collection of reasons why this node is hidden.
	IgnoredReasons	[]*AXProperty	`json:"ignoredReasons,omitempty"`

	// This `Node`'s role, whether explicit or implicit.
	Role	AXValue	`json:"role,omitempty"`

	// The accessible name for this `Node`.
	Name	AXValue	`json:"name,omitempty"`

	// The accessible description for this `Node`.
	Description	AXValue	`json:"description,omitempty"`

	// The value for this `Node`.
	Value	AXValue	`json:"value,omitempty"`

	// All other properties
	Properties	[]*AXProperty	`json:"properties,omitempty"`

	// IDs for each of this node's child nodes.
	ChildIds	[]*AXNodeId	`json:"childIds,omitempty"`

	// The backend ID for the associated DOM node, if any.
	BackendDOMNodeId	dom.BackendNodeId	`json:"backendDOMNodeId,omitempty"`
}

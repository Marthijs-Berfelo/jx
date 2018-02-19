package domdebugger

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// DOMBreakpointType DOM breakpoint type.
type DOMBreakpointType string

// String returns the DOMBreakpointType as string value.
func (t DOMBreakpointType) String() string {
	return string(t)
}

// DOMBreakpointType values.
const (
	DOMBreakpointTypeSubtreeModified   DOMBreakpointType = "subtree-modified"
	DOMBreakpointTypeAttributeModified DOMBreakpointType = "attribute-modified"
	DOMBreakpointTypeNodeRemoved       DOMBreakpointType = "node-removed"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DOMBreakpointType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DOMBreakpointType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DOMBreakpointType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch DOMBreakpointType(in.String()) {
	case DOMBreakpointTypeSubtreeModified:
		*t = DOMBreakpointTypeSubtreeModified
	case DOMBreakpointTypeAttributeModified:
		*t = DOMBreakpointTypeAttributeModified
	case DOMBreakpointTypeNodeRemoved:
		*t = DOMBreakpointTypeNodeRemoved

	default:
		in.AddError(errors.New("unknown DOMBreakpointType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DOMBreakpointType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// EventListener object event listener.
type EventListener struct {
	Type            string                `json:"type"`                      // EventListener's type.
	UseCapture      bool                  `json:"useCapture"`                // EventListener's useCapture.
	Passive         bool                  `json:"passive"`                   // EventListener's passive flag.
	Once            bool                  `json:"once"`                      // EventListener's once flag.
	ScriptID        runtime.ScriptID      `json:"scriptId"`                  // Script id of the handler code.
	LineNumber      int64                 `json:"lineNumber"`                // Line number in the script (0-based).
	ColumnNumber    int64                 `json:"columnNumber"`              // Column number in the script (0-based).
	Handler         *runtime.RemoteObject `json:"handler,omitempty"`         // Event handler function value.
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"` // Event original handler function value.
	BackendNodeID   cdp.BackendNodeID     `json:"backendNodeId,omitempty"`   // Node the listener is added to (if any).
}
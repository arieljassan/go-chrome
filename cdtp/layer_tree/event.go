package layerTree

import (
	"github.com/mkenney/go-chrome/cdtp/dom"
)

/*
LayerPaintedEvent represents LayerTree.layerPainted event data.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
*/
type LayerPaintedEvent struct {
	// The ID of the painted layer.
	LayerID LayerID `json:"layerId"`

	// Clip rectangle.
	Clip *dom.Rect `json:"clip"`
}

/*
DidChangeEvent represents LayerTree.layerTreeDidChange event data.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
*/
type DidChangeEvent struct {
	// Optional. Layer tree, absent if not in the comspositing mode.
	Layers []*Layer `json:"layers,omitempty"`
}
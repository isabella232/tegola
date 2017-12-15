package atlas

import (
	"github.com/airmap/tegola"
	"github.com/airmap/tegola/mvt"
)

type Layer struct {
	//	optional. if not set, the ProviderLayerName will be used
	Name              string
	ProviderLayerName string
	MinZoom           int
	MaxZoom           int
	//	instantiated provider
	Provider mvt.Provider
	//	default tags to include when encoding the layer. provider tags take precedence
	DefaultTags map[string]interface{}
	GeomType    tegola.Geometry
	//	if true, ignore the layer when encoding
	Disabled bool
}

//	MVTName will return the value that will be encoded in the Name field when the layer is encoded as MVT
func (l *Layer) MVTName() string {
	if l.Name != "" {
		return l.Name
	}

	return l.ProviderLayerName
}

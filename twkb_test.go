package twkb

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodePoint(t *testing.T) {
	data, _ := hex.DecodeString("01000204")
	geom, err := Decode(bytes.NewReader(data))

	if err != nil {
		t.Fatalf("Failed to decode point geometry: err = %s", err)
	}

	t.Logf("Geom: %+v\n", geom)

	p, ok := geom.(*Point)

	if !ok {
		t.Fatalf("Expected Point geometry")
	}

	assert.Equal(t, POINT, p.Type())
	assert.Equal(t, XY, p.Dim())
	assert.InDelta(t, 1, p.Coord[0], 1e-6)
	assert.InDelta(t, 2, p.Coord[1], 1e-6)
}

func TestDecodeLineString(t *testing.T) {
	data, _ := hex.DecodeString("02000202020808")
	geom, err := Decode(bytes.NewReader(data))

	if err != nil {
		t.Fatalf("Failed to decode point geometry: err = %s", err)
	}

	t.Logf("Geom: %+v\n", geom)

	p, ok := geom.(*LineString)

	if !ok {
		t.Fatalf("Expected LineString geometry")
	}

	assert.Equal(t, LINESTRING, p.Type())
	assert.Equal(t, XY, p.Dim())
	assert.Equal(t, 2, len(p.Coords))

	assert.InDelta(t, 1, p.Coords[0][0], 1e-6)
	assert.InDelta(t, 1, p.Coords[0][1], 1e-6)

	assert.InDelta(t, 5, p.Coords[1][0], 1e-6)
	assert.InDelta(t, 5, p.Coords[1][1], 1e-6)
}

func TestDecodePolygon(t *testing.T) {
	data, _ := hex.DecodeString("03031b000400040205000004000004030000030500000002020000010100")
	geom, err := Decode(bytes.NewReader(data))

	if err != nil {
		t.Fatalf("Failed to decode point geometry: err = %s", err)
	}

	t.Logf("Geom: %+v\n", geom)

	p, ok := geom.(*Polygon)

	if !ok {
		t.Fatalf("Expected Polygon geometry")
	}

	assert.Equal(t, POLYGON, p.Type())
	assert.Equal(t, XY, p.Dim())
	assert.Equal(t, 2, len(p.Rings))

	// assert.InDelta(t, 1, p.Coords[0][0], 1e-6)
	// assert.InDelta(t, 1, p.Coords[0][1], 1e-6)
	//
	// assert.InDelta(t, 5, p.Coords[1][0], 1e-6)
	// assert.InDelta(t, 5, p.Coords[1][1], 1e-6)
}

func TestDecodeMultiPoint(t *testing.T) {

	data, _ := hex.DecodeString("04070b0004020402000200020404")
	geom, err := Decode(bytes.NewReader(data))

	if err != nil {
		t.Fatalf("Failed to decode multipoint geometry: err = %s", err)
	}

	t.Logf("Geom: %+v\n", geom)

	p, ok := geom.(*MultiPoint)

	if !ok {
		t.Fatalf("Expected Polygon geometry")
	}

	assert.Equal(t, MULTIPOINT, p.Type())
	assert.Equal(t, XY, p.Dim())
}

func TestDecodeGeometryCollection(t *testing.T) {

	data, _ := hex.DecodeString("070402000201000002020002080a0404")
	geom, err := Decode(bytes.NewReader(data))

	if err != nil {
		t.Fatalf("Failed to decode geom collection: err = %s", err)
	}

	t.Logf("Geom: %+v\n", geom)

	p, ok := geom.(*GeometryCollection)

	if !ok {
		t.Fatalf("Expected Polygon geometry")
	}

	assert.Equal(t, GEOMETRYCOLLECTION, p.Type())
	assert.Equal(t, XY, p.Dim())
	assert.Equal(t, 2, len(p.Geometries))

	t.Logf("Geom 0: %+v: POINT ? = %+v\n", p.Geometries[0], p.Geometries[0].Type() == POINT)
	t.Logf("Geom 0: %+v: LINESTRING ? = %+v\n", p.Geometries[1], p.Geometries[1].Type() == LINESTRING)
}

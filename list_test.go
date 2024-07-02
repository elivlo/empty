package empty

import (
	"bytes"
	"encoding/json"
	"testing"
)

type testStruct struct {
	Name string `json:"naMe,omitempty"`
}

func TestEmpty_MarshalJSON(t *testing.T) {
	emptyPrimitive := List[int]{}
	m, err := json.Marshal(emptyPrimitive)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte("[]")) {
		t.Error("Empty array should marshal with brackets")
	}

	var emptyPrimitivePointer List[int]
	m, err = json.Marshal(emptyPrimitivePointer)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte("[]")) {
		t.Error("Nil array should marshal with brackets")
	}

	emptyStruct := List[testStruct]{}
	m, err = json.Marshal(emptyStruct)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte("[]")) {
		t.Error("Empty array should marshal with brackets")
	}

	var emptyStructPointer List[int]
	m, err = json.Marshal(emptyStructPointer)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte("[]")) {
		t.Error("Nil array should marshal with brackets")
	}

	notEmptyPrimitive := List[int]{10, 42}
	m, err = json.Marshal(notEmptyPrimitive)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte("[10,42]")) {
		t.Error("Filled array should marshal normally")
	}

	notEmptyPrimitivePointer := &List[int]{10, 42}
	m, err = json.Marshal(notEmptyPrimitivePointer)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte("[10,42]")) {
		t.Error("Filled array pointer should marshal normally")
	}

	notEmptyStruct := List[testStruct]{{Name: "test1"}, {Name: ""}}
	m, err = json.Marshal(notEmptyStruct)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte(`[{"naMe":"test1"},{}]`)) {
		t.Error("Filled array should marshal normally")
	}

	notEmptyStructPointer := &List[testStruct]{{Name: "test1"}, {Name: ""}}
	m, err = json.Marshal(notEmptyStructPointer)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(m, []byte(`[{"naMe":"test1"},{}]`)) {
		t.Error("Filled array pointer should marshal normally")
	}
}

func TestEmpty_UnmarshalJSON(t *testing.T) {
	emptyPrimitiveJson := []byte(`null`)
	emptyPrimitive := &List[int]{}
	err := json.Unmarshal(emptyPrimitiveJson, emptyPrimitive)
	if err != nil {
		t.Error(err)
	}
	if emptyPrimitive == nil {
		t.Error("Null json array should return empty go slice")
	}

	bracketPrimitiveJson := []byte(`[]`)
	bracketPrimitive := &List[int]{}
	err = json.Unmarshal(bracketPrimitiveJson, bracketPrimitive)
	if err != nil {
		t.Error(err)
	}
	if emptyPrimitive == nil {
		t.Error("Empty json array should return empty go slice")
	}

	primitiveJson := []byte(`[10,42]`)
	primitive := &List[int]{}
	err = json.Unmarshal(primitiveJson, primitive)
	if err != nil {
		t.Error(err)
	}
	if emptyPrimitive == nil {
		t.Error("Empty json array should return empty go slice")
	}

	emptyStructJson := []byte(`null`)
	emptyStruct := &List[testStruct]{}
	err = json.Unmarshal(emptyStructJson, emptyStruct)
	if err != nil {
		t.Error(err)
	}
	if emptyPrimitive == nil {
		t.Error("Null json array should return empty go slice")
	}

	bracketStructJson := []byte(`[]`)
	bracketStruct := &List[testStruct]{}
	err = json.Unmarshal(bracketStructJson, bracketStruct)
	if err != nil {
		t.Error(err)
	}
	if emptyPrimitive == nil {
		t.Error("Empty json array should return empty go slice")
	}

	structJson := []byte(`[{"naMe": "test42"},{}]`)
	parsedStruct := List[testStruct]{}
	err = json.Unmarshal(structJson, &parsedStruct)
	if err != nil {
		t.Error(err)
	}
	if len(parsedStruct) != 2 {
		t.Error("Json array must have len of 2")
	}
	if parsedStruct[0].Name != "test42" {
		t.Error("Did not parse naMe field with value test42 correctly")
	}
	if parsedStruct[1].Name != "" {
		t.Error("Did not parse naMe field with zero value correctly")
	}
}

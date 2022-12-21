package runtime

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	Nil = NewModelTanna("Nil")
	All = NewModelTanna("All")

	Bit = NewModelPlatformWithAlias("Bit", reflect.TypeOf(false))

	Int = NewModelPlatformWithAlias("Int", reflect.TypeOf(int64(0)))
	Dec = NewModelPlatformWithAlias("Dec", reflect.TypeOf(float64(0)))
	Let = NewModelPlatformWithAlias("Let", reflect.TypeOf(' '))
	Txt = NewModelPlatformWithAlias("Txt", reflect.TypeOf(" "))

	RInt  = NewModelPlatformWithAlias("RInt", reflect.TypeOf(IntRange{}))
	RDec  = NewModelPlatformWithAlias("RDec", reflect.TypeOf(DecRange{}))
	RIInt = NewModelPlatformWithAlias("RIInt", reflect.TypeOf(IntRangeIter{}))
	RIDec = NewModelPlatformWithAlias("RIDec", reflect.TypeOf(DecRangeIter{}))

	BuiltInModels = []Model{Nil, All, Bit, Int, Dec, Let, Txt, RInt, RDec, RIInt, RIDec}
)

func FindBuiltInModel(name string) Model {
	for _, model := range BuiltInModels {
		if model.Name() == name {
			return model
		}
	}

	return nil
}

type Model interface {
	Name() string

	Same(other Model) bool
}

type ModelTanna struct {
	name string
}

func (m *ModelTanna) Name() string {
	return m.name
}

func (m *ModelTanna) Same(other Model) bool {
	return m.Name() == "All" || m.Name() == other.Name()
}

type ModelPlatform struct {
	tannaAliased *string
	platformType reflect.Type
}

func (m *ModelPlatform) Name() string {
	if m.tannaAliased != nil {
		return *m.tannaAliased
	} else {
		return m.platformType.Name()
	}
}

func (m *ModelPlatform) Same(other Model) bool {
	if t, ok := other.(*ModelPlatform); !ok {
		return false
	} else {
		return m.platformType.AssignableTo(t.platformType)
	}
}

type ModelTuple struct {
	models []Model
}

func (m *ModelTuple) Name() string {
	names := make([]string, len(m.models))

	for index, model := range m.models {
		names[index] = model.Name()
	}

	return fmt.Sprintf("(%s)", strings.Join(names, ", "))
}

func (m *ModelTuple) Same(other Model) bool {
	if t, ok := other.(*ModelTuple); !ok {
		return false
	} else {
		if len(m.models) != len(t.models) {
			return false
		}

		for i := 0; i < len(m.models); i++ {
			thisTupleType := m.models[i]
			thatTupleType := t.models[i]

			if !thisTupleType.Same(thatTupleType) {
				return false
			}
		}

		return true
	}
}

func (m *ModelTuple) Parts() []Model {
	return m.models
}

func NewModelTanna(name string) Model {
	return &ModelTanna{name: name}
}

func NewModelPlatform(platformType reflect.Type) Model {
	return &ModelPlatform{platformType: platformType}
}

func NewModelPlatformWithAlias(alias string, platformType reflect.Type) Model {
	return &ModelPlatform{platformType: platformType, tannaAliased: &alias}
}

func NewModelTuple(models []Model) Model {
	return &ModelTuple{models: models}
}

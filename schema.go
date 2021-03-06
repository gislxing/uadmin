package uadmin

import (
	"strings"
)

// ModelSchema for a form
type ModelSchema struct {
	Name          string // Name of the Model
	DisplayName   string // Display Name of the model
	ModelName     string // URL
	ModelID       uint
	Inlines       []*ModelSchema
	InlinesData   []listData
	Fields        []F
	IncludeFormJS []string
	IncludeListJS []string
}

// FieldByName returns a field from a ModelSchema by name or nil if
// it doen't exist
func (s ModelSchema) FieldByName(a string) *F {
	for i := range s.Fields {
		if strings.ToLower(s.Fields[i].Name) == strings.ToLower(a) {
			return &s.Fields[i]
		}
	}
	return &F{}
}

// F is a field
type F struct {
	Name              string
	DisplayName       string
	Type              string
	Value             interface{}
	Help              string
	Max               interface{}
	Min               interface{}
	Format            string
	DefaultValue      string
	Required          bool
	Pattern           string
	PatternMsg        string
	Hidden            bool
	ReadOnly          string
	Searchable        bool
	Filter            bool
	ListDisplay       bool
	FormDisplay       bool
	CategoricalFilter bool
	Translations      []translation
	Choices           []Choice
	IsMethod          bool
	ErrMsg            string
	ProgressBar       map[float64]string
	LimitChoicesTo    func(interface{}, *User) []Choice
	UploadTo          string
	Encrypt           bool
}

// Choice is a struct for list choices
type Choice struct {
	V        string
	K        uint
	Selected bool
}

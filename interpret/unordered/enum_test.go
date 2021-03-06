package unordered_test

import (
	"reflect"
	"testing"

	"github.com/yoheimuta/go-protoparser/interpret/unordered"
	"github.com/yoheimuta/go-protoparser/parser"
)

func TestInterpretEnum(t *testing.T) {
	tests := []struct {
		name      string
		inputEnum *parser.Enum
		wantEnum  *unordered.Enum
		wantErr   bool
	}{
		{
			name: "interpreting a nil",
		},
		{
			name: "interpreting an excerpt from the official reference with comments",
			inputEnum: &parser.Enum{
				EnumName: "EnumAllowingAlias",
				EnumBody: []interface{}{
					&parser.Option{
						OptionName: "allow_alias",
						Constant:   "true",
					},
					&parser.EnumField{
						Ident:  "UNKNOWN",
						Number: "0",
					},
					&parser.EnumField{
						Ident:  "STARTED",
						Number: "1",
					},
					&parser.EnumField{
						Ident:  "RUNNING",
						Number: "2",
						EnumValueOptions: []*parser.EnumValueOption{
							{
								OptionName: "(custom_option)",
								Constant:   `"hello world"`,
							},
						},
					},
				},
				Comments: []*parser.Comment{
					{
						Raw: "// enum",
					},
				},
			},
			wantEnum: &unordered.Enum{
				EnumName: "EnumAllowingAlias",
				EnumBody: &unordered.EnumBody{
					Options: []*parser.Option{
						{
							OptionName: "allow_alias",
							Constant:   "true",
						},
					},
					EnumFields: []*parser.EnumField{
						{
							Ident:  "UNKNOWN",
							Number: "0",
						},
						{
							Ident:  "STARTED",
							Number: "1",
						},
						{
							Ident:  "RUNNING",
							Number: "2",
							EnumValueOptions: []*parser.EnumValueOption{
								{
									OptionName: "(custom_option)",
									Constant:   `"hello world"`,
								},
							},
						},
					},
				},
				Comments: []*parser.Comment{
					{
						Raw: "// enum",
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := unordered.InterpretEnum(test.inputEnum)
			switch {
			case test.wantErr:
				if err == nil {
					t.Errorf("got err nil, but want err, parsed=%v", got)
				}
				return
			case !test.wantErr && err != nil:
				t.Errorf("got err %v, but want nil", err)
				return
			}

			if !reflect.DeepEqual(got, test.wantEnum) {
				t.Errorf("got %v, but want %v", got, test.wantEnum)
			}
		})
	}

}

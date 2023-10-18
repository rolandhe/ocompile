// Code generated from Service.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Service

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ServiceParser struct {
	*antlr.BaseParser
}

var ServiceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func serviceParserInit() {
	staticData := &ServiceParserStaticData
	staticData.LiteralNames = []string{
		"", "'with_client'", "'='", "'namespace'", "'go_import'", "'true'",
		"'false'", "'struct'", "'{'", "'}'", "'service'", "'required'", "'optional'",
		"'POST'", "'('", "')'", "'GET'", "'url'", "'list'", "'<'", "'>'", "'void'",
		"'not_login'", "'description'", "'map'", "'['", "']'", "':'", "';'",
		"", "", "", "'bool'", "'byte'", "'i16'", "'i32'", "'i64'", "'double'",
		"'string'", "", "", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "INTEGER", "HEX_INTEGER",
		"DOUBLE", "TYPE_BOOL", "TYPE_BYTE", "TYPE_I16", "TYPE_I32", "TYPE_I64",
		"TYPE_DOUBLE", "TYPE_STRING", "LITERAL", "IDENTIFIER", "COMMA", "WS",
		"SL_COMMENT", "ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"document", "header", "namespace_", "golang_import_", "with_client_optional",
		"definition", "struct_", "service", "field", "field_req", "method_",
		"post_", "get_", "url_", "method_type", "struct_type_list", "get_param_",
		"next_simple_param_", "real_base_type_list_", "void_", "method_param_",
		"single_struct_param", "simple_param_", "real_base_type_parm", "real_base_type_list_parm",
		"not_login", "method_description", "method_description_content", "type_annotations",
		"type_annotation", "field_annotations", "field_annotation", "annotation_value",
		"field_type", "base_type", "container_type", "struct_type", "map_type",
		"list_type", "const_value", "integer", "const_list", "const_map_entry",
		"const_map", "list_separator", "real_base_type", "map_key_type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 399, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		5, 0, 96, 8, 0, 10, 0, 12, 0, 99, 9, 0, 1, 0, 5, 0, 102, 8, 0, 10, 0, 12,
		0, 105, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 114, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 3, 5, 130, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 136, 8,
		6, 10, 6, 12, 6, 139, 9, 6, 1, 6, 1, 6, 3, 6, 143, 8, 6, 1, 7, 1, 7, 1,
		7, 3, 7, 148, 8, 7, 1, 7, 1, 7, 5, 7, 152, 8, 7, 10, 7, 12, 7, 155, 9,
		7, 1, 7, 1, 7, 3, 7, 159, 8, 7, 1, 8, 3, 8, 162, 8, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 167, 8, 8, 1, 8, 3, 8, 170, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10,
		176, 8, 10, 1, 11, 3, 11, 179, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 187, 8, 11, 1, 11, 1, 11, 3, 11, 191, 8, 11, 1, 11, 3, 11,
		194, 8, 11, 1, 12, 3, 12, 197, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 205, 8, 12, 1, 12, 1, 12, 3, 12, 209, 8, 12, 1, 12, 3, 12,
		212, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 223, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 5, 16, 233, 8, 16, 10, 16, 12, 16, 236, 9, 16, 3, 16, 238, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 254, 8, 20, 1, 21, 1, 21, 1, 21, 1, 22,
		3, 22, 260, 8, 22, 1, 22, 1, 22, 3, 22, 264, 8, 22, 1, 22, 3, 22, 267,
		8, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 5, 27, 284, 8, 27, 10, 27, 12, 27,
		287, 9, 27, 1, 28, 1, 28, 5, 28, 291, 8, 28, 10, 28, 12, 28, 294, 9, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29, 301, 8, 29, 1, 29, 3, 29, 304,
		8, 29, 1, 30, 1, 30, 5, 30, 308, 8, 30, 10, 30, 12, 30, 311, 9, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 319, 8, 31, 1, 32, 1, 32,
		3, 32, 323, 8, 32, 1, 33, 1, 33, 1, 33, 3, 33, 328, 8, 33, 1, 34, 1, 34,
		3, 34, 332, 8, 34, 1, 35, 1, 35, 3, 35, 336, 8, 35, 1, 35, 3, 35, 339,
		8, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 3, 39, 362, 8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 3, 41, 369,
		8, 41, 5, 41, 371, 8, 41, 10, 41, 12, 41, 374, 9, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 42, 1, 42, 3, 42, 382, 8, 42, 1, 43, 1, 43, 5, 43, 386, 8,
		43, 10, 43, 12, 43, 389, 9, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 0, 0, 47, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0,
		7, 1, 0, 39, 40, 1, 0, 5, 6, 1, 0, 11, 12, 1, 0, 29, 30, 2, 0, 28, 28,
		41, 41, 1, 0, 32, 38, 1, 0, 33, 38, 405, 0, 97, 1, 0, 0, 0, 2, 113, 1,
		0, 0, 0, 4, 115, 1, 0, 0, 0, 6, 119, 1, 0, 0, 0, 8, 125, 1, 0, 0, 0, 10,
		129, 1, 0, 0, 0, 12, 131, 1, 0, 0, 0, 14, 144, 1, 0, 0, 0, 16, 161, 1,
		0, 0, 0, 18, 171, 1, 0, 0, 0, 20, 175, 1, 0, 0, 0, 22, 178, 1, 0, 0, 0,
		24, 196, 1, 0, 0, 0, 26, 213, 1, 0, 0, 0, 28, 222, 1, 0, 0, 0, 30, 224,
		1, 0, 0, 0, 32, 237, 1, 0, 0, 0, 34, 239, 1, 0, 0, 0, 36, 242, 1, 0, 0,
		0, 38, 247, 1, 0, 0, 0, 40, 253, 1, 0, 0, 0, 42, 255, 1, 0, 0, 0, 44, 266,
		1, 0, 0, 0, 46, 268, 1, 0, 0, 0, 48, 271, 1, 0, 0, 0, 50, 274, 1, 0, 0,
		0, 52, 276, 1, 0, 0, 0, 54, 285, 1, 0, 0, 0, 56, 288, 1, 0, 0, 0, 58, 297,
		1, 0, 0, 0, 60, 305, 1, 0, 0, 0, 62, 314, 1, 0, 0, 0, 64, 322, 1, 0, 0,
		0, 66, 327, 1, 0, 0, 0, 68, 329, 1, 0, 0, 0, 70, 335, 1, 0, 0, 0, 72, 340,
		1, 0, 0, 0, 74, 343, 1, 0, 0, 0, 76, 350, 1, 0, 0, 0, 78, 361, 1, 0, 0,
		0, 80, 363, 1, 0, 0, 0, 82, 365, 1, 0, 0, 0, 84, 377, 1, 0, 0, 0, 86, 383,
		1, 0, 0, 0, 88, 392, 1, 0, 0, 0, 90, 394, 1, 0, 0, 0, 92, 396, 1, 0, 0,
		0, 94, 96, 3, 2, 1, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 103, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0,
		100, 102, 3, 10, 5, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 106, 107, 5, 0, 0, 1, 107, 1, 1, 0, 0, 0, 108, 114, 3, 4, 2,
		0, 109, 114, 3, 6, 3, 0, 110, 111, 5, 1, 0, 0, 111, 112, 5, 2, 0, 0, 112,
		114, 3, 8, 4, 0, 113, 108, 1, 0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 110,
		1, 0, 0, 0, 114, 3, 1, 0, 0, 0, 115, 116, 5, 3, 0, 0, 116, 117, 5, 40,
		0, 0, 117, 118, 7, 0, 0, 0, 118, 5, 1, 0, 0, 0, 119, 121, 5, 4, 0, 0, 120,
		122, 5, 40, 0, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 124, 5, 39, 0, 0, 124, 7, 1, 0, 0, 0, 125, 126, 7, 1,
		0, 0, 126, 9, 1, 0, 0, 0, 127, 130, 3, 12, 6, 0, 128, 130, 3, 14, 7, 0,
		129, 127, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 11, 1, 0, 0, 0, 131, 132,
		5, 7, 0, 0, 132, 133, 5, 40, 0, 0, 133, 137, 5, 8, 0, 0, 134, 136, 3, 16,
		8, 0, 135, 134, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0,
		137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140,
		142, 5, 9, 0, 0, 141, 143, 3, 56, 28, 0, 142, 141, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 13, 1, 0, 0, 0, 144, 145, 5, 10, 0, 0, 145, 147, 5, 40,
		0, 0, 146, 148, 3, 26, 13, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0,
		0, 148, 149, 1, 0, 0, 0, 149, 153, 5, 8, 0, 0, 150, 152, 3, 20, 10, 0,
		151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 158,
		5, 9, 0, 0, 157, 159, 3, 56, 28, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1,
		0, 0, 0, 159, 15, 1, 0, 0, 0, 160, 162, 3, 18, 9, 0, 161, 160, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 3, 66, 33, 0,
		164, 166, 5, 40, 0, 0, 165, 167, 3, 60, 30, 0, 166, 165, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 170, 3, 88, 44, 0, 169, 168,
		1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 17, 1, 0, 0, 0, 171, 172, 7, 2,
		0, 0, 172, 19, 1, 0, 0, 0, 173, 176, 3, 22, 11, 0, 174, 176, 3, 24, 12,
		0, 175, 173, 1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 21, 1, 0, 0, 0, 177,
		179, 3, 52, 26, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180,
		1, 0, 0, 0, 180, 181, 5, 13, 0, 0, 181, 182, 3, 26, 13, 0, 182, 183, 3,
		28, 14, 0, 183, 184, 5, 40, 0, 0, 184, 186, 5, 14, 0, 0, 185, 187, 3, 40,
		20, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 190, 5, 15, 0, 0, 189, 191, 3, 50, 25, 0, 190, 189, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 194, 3, 88, 44, 0, 193, 192,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 23, 1, 0, 0, 0, 195, 197, 3, 52,
		26, 0, 196, 195, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0,
		198, 199, 5, 16, 0, 0, 199, 200, 3, 26, 13, 0, 200, 201, 3, 28, 14, 0,
		201, 202, 5, 40, 0, 0, 202, 204, 5, 14, 0, 0, 203, 205, 3, 32, 16, 0, 204,
		203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208,
		5, 15, 0, 0, 207, 209, 3, 50, 25, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1,
		0, 0, 0, 209, 211, 1, 0, 0, 0, 210, 212, 3, 88, 44, 0, 211, 210, 1, 0,
		0, 0, 211, 212, 1, 0, 0, 0, 212, 25, 1, 0, 0, 0, 213, 214, 5, 17, 0, 0,
		214, 215, 5, 2, 0, 0, 215, 216, 5, 39, 0, 0, 216, 27, 1, 0, 0, 0, 217,
		223, 3, 90, 45, 0, 218, 223, 3, 36, 18, 0, 219, 223, 3, 38, 19, 0, 220,
		223, 3, 72, 36, 0, 221, 223, 3, 30, 15, 0, 222, 217, 1, 0, 0, 0, 222, 218,
		1, 0, 0, 0, 222, 219, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0,
		0, 0, 223, 29, 1, 0, 0, 0, 224, 225, 5, 18, 0, 0, 225, 226, 5, 19, 0, 0,
		226, 227, 3, 72, 36, 0, 227, 228, 5, 20, 0, 0, 228, 31, 1, 0, 0, 0, 229,
		238, 3, 42, 21, 0, 230, 234, 3, 44, 22, 0, 231, 233, 3, 34, 17, 0, 232,
		231, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235,
		1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 229, 1, 0,
		0, 0, 237, 230, 1, 0, 0, 0, 238, 33, 1, 0, 0, 0, 239, 240, 5, 41, 0, 0,
		240, 241, 3, 44, 22, 0, 241, 35, 1, 0, 0, 0, 242, 243, 5, 18, 0, 0, 243,
		244, 5, 19, 0, 0, 244, 245, 3, 90, 45, 0, 245, 246, 5, 20, 0, 0, 246, 37,
		1, 0, 0, 0, 247, 248, 5, 21, 0, 0, 248, 39, 1, 0, 0, 0, 249, 254, 3, 42,
		21, 0, 250, 251, 3, 30, 15, 0, 251, 252, 5, 40, 0, 0, 252, 254, 1, 0, 0,
		0, 253, 249, 1, 0, 0, 0, 253, 250, 1, 0, 0, 0, 254, 41, 1, 0, 0, 0, 255,
		256, 3, 72, 36, 0, 256, 257, 5, 40, 0, 0, 257, 43, 1, 0, 0, 0, 258, 260,
		3, 18, 9, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 1, 0,
		0, 0, 261, 267, 3, 46, 23, 0, 262, 264, 3, 18, 9, 0, 263, 262, 1, 0, 0,
		0, 263, 264, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 3, 48, 24, 0,
		266, 259, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0, 267, 45, 1, 0, 0, 0, 268, 269,
		3, 90, 45, 0, 269, 270, 5, 40, 0, 0, 270, 47, 1, 0, 0, 0, 271, 272, 3,
		36, 18, 0, 272, 273, 5, 40, 0, 0, 273, 49, 1, 0, 0, 0, 274, 275, 5, 22,
		0, 0, 275, 51, 1, 0, 0, 0, 276, 277, 5, 14, 0, 0, 277, 278, 5, 23, 0, 0,
		278, 279, 5, 2, 0, 0, 279, 280, 3, 54, 27, 0, 280, 281, 5, 15, 0, 0, 281,
		53, 1, 0, 0, 0, 282, 284, 3, 64, 32, 0, 283, 282, 1, 0, 0, 0, 284, 287,
		1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 55, 1, 0,
		0, 0, 287, 285, 1, 0, 0, 0, 288, 292, 5, 14, 0, 0, 289, 291, 3, 58, 29,
		0, 290, 289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292,
		293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296,
		5, 15, 0, 0, 296, 57, 1, 0, 0, 0, 297, 300, 5, 40, 0, 0, 298, 299, 5, 2,
		0, 0, 299, 301, 3, 64, 32, 0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0,
		0, 301, 303, 1, 0, 0, 0, 302, 304, 3, 88, 44, 0, 303, 302, 1, 0, 0, 0,
		303, 304, 1, 0, 0, 0, 304, 59, 1, 0, 0, 0, 305, 309, 5, 14, 0, 0, 306,
		308, 3, 62, 31, 0, 307, 306, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307,
		1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 312, 1, 0, 0, 0, 311, 309, 1, 0,
		0, 0, 312, 313, 5, 15, 0, 0, 313, 61, 1, 0, 0, 0, 314, 315, 5, 40, 0, 0,
		315, 316, 5, 2, 0, 0, 316, 318, 5, 39, 0, 0, 317, 319, 3, 88, 44, 0, 318,
		317, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 63, 1, 0, 0, 0, 320, 323, 3,
		80, 40, 0, 321, 323, 5, 39, 0, 0, 322, 320, 1, 0, 0, 0, 322, 321, 1, 0,
		0, 0, 323, 65, 1, 0, 0, 0, 324, 328, 3, 68, 34, 0, 325, 328, 3, 72, 36,
		0, 326, 328, 3, 70, 35, 0, 327, 324, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0,
		327, 326, 1, 0, 0, 0, 328, 67, 1, 0, 0, 0, 329, 331, 3, 90, 45, 0, 330,
		332, 3, 56, 28, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 69,
		1, 0, 0, 0, 333, 336, 3, 74, 37, 0, 334, 336, 3, 76, 38, 0, 335, 333, 1,
		0, 0, 0, 335, 334, 1, 0, 0, 0, 336, 338, 1, 0, 0, 0, 337, 339, 3, 56, 28,
		0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 71, 1, 0, 0, 0, 340,
		341, 5, 7, 0, 0, 341, 342, 5, 40, 0, 0, 342, 73, 1, 0, 0, 0, 343, 344,
		5, 24, 0, 0, 344, 345, 5, 19, 0, 0, 345, 346, 3, 92, 46, 0, 346, 347, 5,
		41, 0, 0, 347, 348, 3, 66, 33, 0, 348, 349, 5, 20, 0, 0, 349, 75, 1, 0,
		0, 0, 350, 351, 5, 18, 0, 0, 351, 352, 5, 19, 0, 0, 352, 353, 3, 66, 33,
		0, 353, 354, 5, 20, 0, 0, 354, 77, 1, 0, 0, 0, 355, 362, 3, 80, 40, 0,
		356, 362, 5, 31, 0, 0, 357, 362, 5, 39, 0, 0, 358, 362, 5, 40, 0, 0, 359,
		362, 3, 82, 41, 0, 360, 362, 3, 86, 43, 0, 361, 355, 1, 0, 0, 0, 361, 356,
		1, 0, 0, 0, 361, 357, 1, 0, 0, 0, 361, 358, 1, 0, 0, 0, 361, 359, 1, 0,
		0, 0, 361, 360, 1, 0, 0, 0, 362, 79, 1, 0, 0, 0, 363, 364, 7, 3, 0, 0,
		364, 81, 1, 0, 0, 0, 365, 372, 5, 25, 0, 0, 366, 368, 3, 78, 39, 0, 367,
		369, 3, 88, 44, 0, 368, 367, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 371,
		1, 0, 0, 0, 370, 366, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0,
		0, 0, 372, 373, 1, 0, 0, 0, 373, 375, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0,
		375, 376, 5, 26, 0, 0, 376, 83, 1, 0, 0, 0, 377, 378, 3, 78, 39, 0, 378,
		379, 5, 27, 0, 0, 379, 381, 3, 78, 39, 0, 380, 382, 3, 88, 44, 0, 381,
		380, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 85, 1, 0, 0, 0, 383, 387, 5,
		8, 0, 0, 384, 386, 3, 84, 42, 0, 385, 384, 1, 0, 0, 0, 386, 389, 1, 0,
		0, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 390, 1, 0, 0, 0,
		389, 387, 1, 0, 0, 0, 390, 391, 5, 9, 0, 0, 391, 87, 1, 0, 0, 0, 392, 393,
		7, 4, 0, 0, 393, 89, 1, 0, 0, 0, 394, 395, 7, 5, 0, 0, 395, 91, 1, 0, 0,
		0, 396, 397, 7, 6, 0, 0, 397, 93, 1, 0, 0, 0, 45, 97, 103, 113, 121, 129,
		137, 142, 147, 153, 158, 161, 166, 169, 175, 178, 186, 190, 193, 196, 204,
		208, 211, 222, 234, 237, 253, 259, 263, 266, 285, 292, 300, 303, 309, 318,
		322, 327, 331, 335, 338, 361, 368, 372, 381, 387,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ServiceParserInit initializes any static state used to implement ServiceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewServiceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ServiceParserInit() {
	staticData := &ServiceParserStaticData
	staticData.once.Do(serviceParserInit)
}

// NewServiceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewServiceParser(input antlr.TokenStream) *ServiceParser {
	ServiceParserInit()
	this := new(ServiceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ServiceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Service.g4"

	return this
}

// ServiceParser tokens.
const (
	ServiceParserEOF         = antlr.TokenEOF
	ServiceParserT__0        = 1
	ServiceParserT__1        = 2
	ServiceParserT__2        = 3
	ServiceParserT__3        = 4
	ServiceParserT__4        = 5
	ServiceParserT__5        = 6
	ServiceParserT__6        = 7
	ServiceParserT__7        = 8
	ServiceParserT__8        = 9
	ServiceParserT__9        = 10
	ServiceParserT__10       = 11
	ServiceParserT__11       = 12
	ServiceParserT__12       = 13
	ServiceParserT__13       = 14
	ServiceParserT__14       = 15
	ServiceParserT__15       = 16
	ServiceParserT__16       = 17
	ServiceParserT__17       = 18
	ServiceParserT__18       = 19
	ServiceParserT__19       = 20
	ServiceParserT__20       = 21
	ServiceParserT__21       = 22
	ServiceParserT__22       = 23
	ServiceParserT__23       = 24
	ServiceParserT__24       = 25
	ServiceParserT__25       = 26
	ServiceParserT__26       = 27
	ServiceParserT__27       = 28
	ServiceParserINTEGER     = 29
	ServiceParserHEX_INTEGER = 30
	ServiceParserDOUBLE      = 31
	ServiceParserTYPE_BOOL   = 32
	ServiceParserTYPE_BYTE   = 33
	ServiceParserTYPE_I16    = 34
	ServiceParserTYPE_I32    = 35
	ServiceParserTYPE_I64    = 36
	ServiceParserTYPE_DOUBLE = 37
	ServiceParserTYPE_STRING = 38
	ServiceParserLITERAL     = 39
	ServiceParserIDENTIFIER  = 40
	ServiceParserCOMMA       = 41
	ServiceParserWS          = 42
	ServiceParserSL_COMMENT  = 43
	ServiceParserML_COMMENT  = 44
)

// ServiceParser rules.
const (
	ServiceParserRULE_document                   = 0
	ServiceParserRULE_header                     = 1
	ServiceParserRULE_namespace_                 = 2
	ServiceParserRULE_golang_import_             = 3
	ServiceParserRULE_with_client_optional       = 4
	ServiceParserRULE_definition                 = 5
	ServiceParserRULE_struct_                    = 6
	ServiceParserRULE_service                    = 7
	ServiceParserRULE_field                      = 8
	ServiceParserRULE_field_req                  = 9
	ServiceParserRULE_method_                    = 10
	ServiceParserRULE_post_                      = 11
	ServiceParserRULE_get_                       = 12
	ServiceParserRULE_url_                       = 13
	ServiceParserRULE_method_type                = 14
	ServiceParserRULE_struct_type_list           = 15
	ServiceParserRULE_get_param_                 = 16
	ServiceParserRULE_next_simple_param_         = 17
	ServiceParserRULE_real_base_type_list_       = 18
	ServiceParserRULE_void_                      = 19
	ServiceParserRULE_method_param_              = 20
	ServiceParserRULE_single_struct_param        = 21
	ServiceParserRULE_simple_param_              = 22
	ServiceParserRULE_real_base_type_parm        = 23
	ServiceParserRULE_real_base_type_list_parm   = 24
	ServiceParserRULE_not_login                  = 25
	ServiceParserRULE_method_description         = 26
	ServiceParserRULE_method_description_content = 27
	ServiceParserRULE_type_annotations           = 28
	ServiceParserRULE_type_annotation            = 29
	ServiceParserRULE_field_annotations          = 30
	ServiceParserRULE_field_annotation           = 31
	ServiceParserRULE_annotation_value           = 32
	ServiceParserRULE_field_type                 = 33
	ServiceParserRULE_base_type                  = 34
	ServiceParserRULE_container_type             = 35
	ServiceParserRULE_struct_type                = 36
	ServiceParserRULE_map_type                   = 37
	ServiceParserRULE_list_type                  = 38
	ServiceParserRULE_const_value                = 39
	ServiceParserRULE_integer                    = 40
	ServiceParserRULE_const_list                 = 41
	ServiceParserRULE_const_map_entry            = 42
	ServiceParserRULE_const_map                  = 43
	ServiceParserRULE_list_separator             = 44
	ServiceParserRULE_real_base_type             = 45
	ServiceParserRULE_map_key_type               = 46
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllHeader() []IHeaderContext
	Header(i int) IHeaderContext
	AllDefinition() []IDefinitionContext
	Definition(i int) IDefinitionContext

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(ServiceParserEOF, 0)
}

func (s *DocumentContext) AllHeader() []IHeaderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeaderContext); ok {
			len++
		}
	}

	tst := make([]IHeaderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeaderContext); ok {
			tst[i] = t.(IHeaderContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Header(i int) IHeaderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *ServiceParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ServiceParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26) != 0 {
		{
			p.SetState(94)
			p.Header()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ServiceParserT__6 || _la == ServiceParserT__9 {
		{
			p.SetState(100)
			p.Definition()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(ServiceParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Namespace_() INamespace_Context
	Golang_import_() IGolang_import_Context
	With_client_optional() IWith_client_optionalContext

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_header
	return p
}

func InitEmptyHeaderContext(p *HeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_header
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Namespace_() INamespace_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_Context)
}

func (s *HeaderContext) Golang_import_() IGolang_import_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGolang_import_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGolang_import_Context)
}

func (s *HeaderContext) With_client_optional() IWith_client_optionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWith_client_optionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWith_client_optionalContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *ServiceParser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ServiceParserRULE_header)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Namespace_()
		}

	case ServiceParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Golang_import_()
		}

	case ServiceParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(ServiceParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(ServiceParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.With_client_optional()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamespace_Context is an interface to support dynamic dispatch.
type INamespace_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LITERAL() antlr.TerminalNode

	// IsNamespace_Context differentiates from other interfaces.
	IsNamespace_Context()
}

type Namespace_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespace_Context() *Namespace_Context {
	var p = new(Namespace_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_namespace_
	return p
}

func InitEmptyNamespace_Context(p *Namespace_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_namespace_
}

func (*Namespace_Context) IsNamespace_Context() {}

func NewNamespace_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_Context {
	var p = new(Namespace_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_namespace_

	return p
}

func (s *Namespace_Context) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ServiceParserIDENTIFIER)
}

func (s *Namespace_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, i)
}

func (s *Namespace_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ServiceParserLITERAL, 0)
}

func (s *Namespace_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterNamespace_(s)
	}
}

func (s *Namespace_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitNamespace_(s)
	}
}

func (p *ServiceParser) Namespace_() (localctx INamespace_Context) {
	localctx = NewNamespace_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ServiceParserRULE_namespace_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(ServiceParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ServiceParserLITERAL || _la == ServiceParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGolang_import_Context is an interface to support dynamic dispatch.
type IGolang_import_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsGolang_import_Context differentiates from other interfaces.
	IsGolang_import_Context()
}

type Golang_import_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGolang_import_Context() *Golang_import_Context {
	var p = new(Golang_import_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_golang_import_
	return p
}

func InitEmptyGolang_import_Context(p *Golang_import_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_golang_import_
}

func (*Golang_import_Context) IsGolang_import_Context() {}

func NewGolang_import_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Golang_import_Context {
	var p = new(Golang_import_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_golang_import_

	return p
}

func (s *Golang_import_Context) GetParser() antlr.Parser { return s.parser }

func (s *Golang_import_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ServiceParserLITERAL, 0)
}

func (s *Golang_import_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Golang_import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Golang_import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Golang_import_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterGolang_import_(s)
	}
}

func (s *Golang_import_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitGolang_import_(s)
	}
}

func (p *ServiceParser) Golang_import_() (localctx IGolang_import_Context) {
	localctx = NewGolang_import_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ServiceParserRULE_golang_import_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(ServiceParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserIDENTIFIER {
		{
			p.SetState(120)
			p.Match(ServiceParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(123)
		p.Match(ServiceParserLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWith_client_optionalContext is an interface to support dynamic dispatch.
type IWith_client_optionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWith_client_optionalContext differentiates from other interfaces.
	IsWith_client_optionalContext()
}

type With_client_optionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWith_client_optionalContext() *With_client_optionalContext {
	var p = new(With_client_optionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_with_client_optional
	return p
}

func InitEmptyWith_client_optionalContext(p *With_client_optionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_with_client_optional
}

func (*With_client_optionalContext) IsWith_client_optionalContext() {}

func NewWith_client_optionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *With_client_optionalContext {
	var p = new(With_client_optionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_with_client_optional

	return p
}

func (s *With_client_optionalContext) GetParser() antlr.Parser { return s.parser }
func (s *With_client_optionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *With_client_optionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *With_client_optionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterWith_client_optional(s)
	}
}

func (s *With_client_optionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitWith_client_optional(s)
	}
}

func (p *ServiceParser) With_client_optional() (localctx IWith_client_optionalContext) {
	localctx = NewWith_client_optionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ServiceParserRULE_with_client_optional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ServiceParserT__4 || _la == ServiceParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct_() IStruct_Context
	Service() IServiceContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_definition
	return p
}

func InitEmptyDefinitionContext(p *DefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_definition
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Struct_() IStruct_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_Context)
}

func (s *DefinitionContext) Service() IServiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *ServiceParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ServiceParserRULE_definition)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Struct_()
		}

	case ServiceParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Service()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_Context is an interface to support dynamic dispatch.
type IStruct_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext
	Type_annotations() IType_annotationsContext

	// IsStruct_Context differentiates from other interfaces.
	IsStruct_Context()
}

type Struct_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_Context() *Struct_Context {
	var p = new(Struct_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_struct_
	return p
}

func InitEmptyStruct_Context(p *Struct_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_struct_
}

func (*Struct_Context) IsStruct_Context() {}

func NewStruct_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_Context {
	var p = new(Struct_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_struct_

	return p
}

func (s *Struct_Context) GetParser() antlr.Parser { return s.parser }

func (s *Struct_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Struct_Context) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *Struct_Context) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Struct_Context) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Struct_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterStruct_(s)
	}
}

func (s *Struct_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitStruct_(s)
	}
}

func (p *ServiceParser) Struct_() (localctx IStruct_Context) {
	localctx = NewStruct_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ServiceParserRULE_struct_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(ServiceParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(ServiceParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&545477892224) != 0 {
		{
			p.SetState(134)
			p.Field()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(ServiceParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(141)
			p.Type_annotations()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Url_() IUrl_Context
	AllMethod_() []IMethod_Context
	Method_(i int) IMethod_Context
	Type_annotations() IType_annotationsContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_service
	return p
}

func InitEmptyServiceContext(p *ServiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_service
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *ServiceContext) Url_() IUrl_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrl_Context)
}

func (s *ServiceContext) AllMethod_() []IMethod_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethod_Context); ok {
			len++
		}
	}

	tst := make([]IMethod_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethod_Context); ok {
			tst[i] = t.(IMethod_Context)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) Method_(i int) IMethod_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_Context)
}

func (s *ServiceContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *ServiceParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ServiceParserRULE_service)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(ServiceParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__16 {
		{
			p.SetState(146)
			p.Url_()
		}

	}
	{
		p.SetState(149)
		p.Match(ServiceParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&90112) != 0 {
		{
			p.SetState(150)
			p.Method_()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(ServiceParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(157)
			p.Type_annotations()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	IDENTIFIER() antlr.TerminalNode
	Field_req() IField_reqContext
	Field_annotations() IField_annotationsContext
	List_separator() IList_separatorContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *FieldContext) Field_req() IField_reqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_reqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_reqContext)
}

func (s *FieldContext) Field_annotations() IField_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_annotationsContext)
}

func (s *FieldContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ServiceParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ServiceParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__10 || _la == ServiceParserT__11 {
		{
			p.SetState(160)
			p.Field_req()
		}

	}
	{
		p.SetState(163)
		p.Field_type()
	}
	{
		p.SetState(164)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(165)
			p.Field_annotations()
		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
		{
			p.SetState(168)
			p.List_separator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_reqContext is an interface to support dynamic dispatch.
type IField_reqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsField_reqContext differentiates from other interfaces.
	IsField_reqContext()
}

type Field_reqContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_reqContext() *Field_reqContext {
	var p = new(Field_reqContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_req
	return p
}

func InitEmptyField_reqContext(p *Field_reqContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_req
}

func (*Field_reqContext) IsField_reqContext() {}

func NewField_reqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_reqContext {
	var p = new(Field_reqContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_field_req

	return p
}

func (s *Field_reqContext) GetParser() antlr.Parser { return s.parser }
func (s *Field_reqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_reqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_reqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterField_req(s)
	}
}

func (s *Field_reqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitField_req(s)
	}
}

func (p *ServiceParser) Field_req() (localctx IField_reqContext) {
	localctx = NewField_reqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ServiceParserRULE_field_req)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ServiceParserT__10 || _la == ServiceParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_Context is an interface to support dynamic dispatch.
type IMethod_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Post_() IPost_Context
	Get_() IGet_Context

	// IsMethod_Context differentiates from other interfaces.
	IsMethod_Context()
}

type Method_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_Context() *Method_Context {
	var p = new(Method_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_
	return p
}

func InitEmptyMethod_Context(p *Method_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_
}

func (*Method_Context) IsMethod_Context() {}

func NewMethod_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_Context {
	var p = new(Method_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_method_

	return p
}

func (s *Method_Context) GetParser() antlr.Parser { return s.parser }

func (s *Method_Context) Post_() IPost_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPost_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPost_Context)
}

func (s *Method_Context) Get_() IGet_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_Context)
}

func (s *Method_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMethod_(s)
	}
}

func (s *Method_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMethod_(s)
	}
}

func (p *ServiceParser) Method_() (localctx IMethod_Context) {
	localctx = NewMethod_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ServiceParserRULE_method_)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Post_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Get_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPost_Context is an interface to support dynamic dispatch.
type IPost_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Url_() IUrl_Context
	Method_type() IMethod_typeContext
	IDENTIFIER() antlr.TerminalNode
	Method_description() IMethod_descriptionContext
	Method_param_() IMethod_param_Context
	Not_login() INot_loginContext
	List_separator() IList_separatorContext

	// IsPost_Context differentiates from other interfaces.
	IsPost_Context()
}

type Post_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPost_Context() *Post_Context {
	var p = new(Post_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_post_
	return p
}

func InitEmptyPost_Context(p *Post_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_post_
}

func (*Post_Context) IsPost_Context() {}

func NewPost_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Post_Context {
	var p = new(Post_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_post_

	return p
}

func (s *Post_Context) GetParser() antlr.Parser { return s.parser }

func (s *Post_Context) Url_() IUrl_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrl_Context)
}

func (s *Post_Context) Method_type() IMethod_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_typeContext)
}

func (s *Post_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Post_Context) Method_description() IMethod_descriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_descriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_descriptionContext)
}

func (s *Post_Context) Method_param_() IMethod_param_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_param_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_param_Context)
}

func (s *Post_Context) Not_login() INot_loginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INot_loginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INot_loginContext)
}

func (s *Post_Context) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Post_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Post_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Post_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterPost_(s)
	}
}

func (s *Post_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitPost_(s)
	}
}

func (p *ServiceParser) Post_() (localctx IPost_Context) {
	localctx = NewPost_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ServiceParserRULE_post_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(177)
			p.Method_description()
		}

	}
	{
		p.SetState(180)
		p.Match(ServiceParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Url_()
	}
	{
		p.SetState(182)
		p.Method_type()
	}
	{
		p.SetState(183)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(ServiceParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__6 || _la == ServiceParserT__17 {
		{
			p.SetState(185)
			p.Method_param_()
		}

	}
	{
		p.SetState(188)
		p.Match(ServiceParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__21 {
		{
			p.SetState(189)
			p.Not_login()
		}

	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
		{
			p.SetState(192)
			p.List_separator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_Context is an interface to support dynamic dispatch.
type IGet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Url_() IUrl_Context
	Method_type() IMethod_typeContext
	IDENTIFIER() antlr.TerminalNode
	Method_description() IMethod_descriptionContext
	Get_param_() IGet_param_Context
	Not_login() INot_loginContext
	List_separator() IList_separatorContext

	// IsGet_Context differentiates from other interfaces.
	IsGet_Context()
}

type Get_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_Context() *Get_Context {
	var p = new(Get_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_get_
	return p
}

func InitEmptyGet_Context(p *Get_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_get_
}

func (*Get_Context) IsGet_Context() {}

func NewGet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_Context {
	var p = new(Get_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_get_

	return p
}

func (s *Get_Context) GetParser() antlr.Parser { return s.parser }

func (s *Get_Context) Url_() IUrl_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrl_Context)
}

func (s *Get_Context) Method_type() IMethod_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_typeContext)
}

func (s *Get_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Get_Context) Method_description() IMethod_descriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_descriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_descriptionContext)
}

func (s *Get_Context) Get_param_() IGet_param_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_param_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_param_Context)
}

func (s *Get_Context) Not_login() INot_loginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INot_loginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INot_loginContext)
}

func (s *Get_Context) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Get_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterGet_(s)
	}
}

func (s *Get_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitGet_(s)
	}
}

func (p *ServiceParser) Get_() (localctx IGet_Context) {
	localctx = NewGet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ServiceParserRULE_get_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(195)
			p.Method_description()
		}

	}
	{
		p.SetState(198)
		p.Match(ServiceParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Url_()
	}
	{
		p.SetState(200)
		p.Method_type()
	}
	{
		p.SetState(201)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(ServiceParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&545461115008) != 0 {
		{
			p.SetState(203)
			p.Get_param_()
		}

	}
	{
		p.SetState(206)
		p.Match(ServiceParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__21 {
		{
			p.SetState(207)
			p.Not_login()
		}

	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
		{
			p.SetState(210)
			p.List_separator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUrl_Context is an interface to support dynamic dispatch.
type IUrl_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode

	// IsUrl_Context differentiates from other interfaces.
	IsUrl_Context()
}

type Url_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrl_Context() *Url_Context {
	var p = new(Url_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_url_
	return p
}

func InitEmptyUrl_Context(p *Url_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_url_
}

func (*Url_Context) IsUrl_Context() {}

func NewUrl_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Url_Context {
	var p = new(Url_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_url_

	return p
}

func (s *Url_Context) GetParser() antlr.Parser { return s.parser }

func (s *Url_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ServiceParserLITERAL, 0)
}

func (s *Url_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Url_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Url_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterUrl_(s)
	}
}

func (s *Url_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitUrl_(s)
	}
}

func (p *ServiceParser) Url_() (localctx IUrl_Context) {
	localctx = NewUrl_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ServiceParserRULE_url_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(ServiceParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(ServiceParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(ServiceParserLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_typeContext is an interface to support dynamic dispatch.
type IMethod_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type() IReal_base_typeContext
	Real_base_type_list_() IReal_base_type_list_Context
	Void_() IVoid_Context
	Struct_type() IStruct_typeContext
	Struct_type_list() IStruct_type_listContext

	// IsMethod_typeContext differentiates from other interfaces.
	IsMethod_typeContext()
}

type Method_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_typeContext() *Method_typeContext {
	var p = new(Method_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_type
	return p
}

func InitEmptyMethod_typeContext(p *Method_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_type
}

func (*Method_typeContext) IsMethod_typeContext() {}

func NewMethod_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_typeContext {
	var p = new(Method_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_method_type

	return p
}

func (s *Method_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_typeContext) Real_base_type() IReal_base_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_typeContext)
}

func (s *Method_typeContext) Real_base_type_list_() IReal_base_type_list_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_type_list_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_type_list_Context)
}

func (s *Method_typeContext) Void_() IVoid_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVoid_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVoid_Context)
}

func (s *Method_typeContext) Struct_type() IStruct_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_typeContext)
}

func (s *Method_typeContext) Struct_type_list() IStruct_type_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_type_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_type_listContext)
}

func (s *Method_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMethod_type(s)
	}
}

func (s *Method_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMethod_type(s)
	}
}

func (p *ServiceParser) Method_type() (localctx IMethod_typeContext) {
	localctx = NewMethod_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ServiceParserRULE_method_type)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Real_base_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Real_base_type_list_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.Void_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Struct_type()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(221)
			p.Struct_type_list()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_type_listContext is an interface to support dynamic dispatch.
type IStruct_type_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct_type() IStruct_typeContext

	// IsStruct_type_listContext differentiates from other interfaces.
	IsStruct_type_listContext()
}

type Struct_type_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_type_listContext() *Struct_type_listContext {
	var p = new(Struct_type_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_struct_type_list
	return p
}

func InitEmptyStruct_type_listContext(p *Struct_type_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_struct_type_list
}

func (*Struct_type_listContext) IsStruct_type_listContext() {}

func NewStruct_type_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_type_listContext {
	var p = new(Struct_type_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_struct_type_list

	return p
}

func (s *Struct_type_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_type_listContext) Struct_type() IStruct_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_typeContext)
}

func (s *Struct_type_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_type_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_type_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterStruct_type_list(s)
	}
}

func (s *Struct_type_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitStruct_type_list(s)
	}
}

func (p *ServiceParser) Struct_type_list() (localctx IStruct_type_listContext) {
	localctx = NewStruct_type_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ServiceParserRULE_struct_type_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(ServiceParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(ServiceParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Struct_type()
	}
	{
		p.SetState(227)
		p.Match(ServiceParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_param_Context is an interface to support dynamic dispatch.
type IGet_param_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Single_struct_param() ISingle_struct_paramContext
	Simple_param_() ISimple_param_Context
	AllNext_simple_param_() []INext_simple_param_Context
	Next_simple_param_(i int) INext_simple_param_Context

	// IsGet_param_Context differentiates from other interfaces.
	IsGet_param_Context()
}

type Get_param_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_param_Context() *Get_param_Context {
	var p = new(Get_param_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_get_param_
	return p
}

func InitEmptyGet_param_Context(p *Get_param_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_get_param_
}

func (*Get_param_Context) IsGet_param_Context() {}

func NewGet_param_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_param_Context {
	var p = new(Get_param_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_get_param_

	return p
}

func (s *Get_param_Context) GetParser() antlr.Parser { return s.parser }

func (s *Get_param_Context) Single_struct_param() ISingle_struct_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_struct_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_struct_paramContext)
}

func (s *Get_param_Context) Simple_param_() ISimple_param_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_param_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_param_Context)
}

func (s *Get_param_Context) AllNext_simple_param_() []INext_simple_param_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INext_simple_param_Context); ok {
			len++
		}
	}

	tst := make([]INext_simple_param_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INext_simple_param_Context); ok {
			tst[i] = t.(INext_simple_param_Context)
			i++
		}
	}

	return tst
}

func (s *Get_param_Context) Next_simple_param_(i int) INext_simple_param_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INext_simple_param_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INext_simple_param_Context)
}

func (s *Get_param_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_param_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_param_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterGet_param_(s)
	}
}

func (s *Get_param_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitGet_param_(s)
	}
}

func (p *ServiceParser) Get_param_() (localctx IGet_param_Context) {
	localctx = NewGet_param_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ServiceParserRULE_get_param_)
	var _la int

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Single_struct_param()
		}

	case ServiceParserT__10, ServiceParserT__11, ServiceParserT__17, ServiceParserTYPE_BOOL, ServiceParserTYPE_BYTE, ServiceParserTYPE_I16, ServiceParserTYPE_I32, ServiceParserTYPE_I64, ServiceParserTYPE_DOUBLE, ServiceParserTYPE_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Simple_param_()
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ServiceParserCOMMA {
			{
				p.SetState(231)
				p.Next_simple_param_()
			}

			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INext_simple_param_Context is an interface to support dynamic dispatch.
type INext_simple_param_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Simple_param_() ISimple_param_Context

	// IsNext_simple_param_Context differentiates from other interfaces.
	IsNext_simple_param_Context()
}

type Next_simple_param_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNext_simple_param_Context() *Next_simple_param_Context {
	var p = new(Next_simple_param_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_next_simple_param_
	return p
}

func InitEmptyNext_simple_param_Context(p *Next_simple_param_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_next_simple_param_
}

func (*Next_simple_param_Context) IsNext_simple_param_Context() {}

func NewNext_simple_param_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Next_simple_param_Context {
	var p = new(Next_simple_param_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_next_simple_param_

	return p
}

func (s *Next_simple_param_Context) GetParser() antlr.Parser { return s.parser }

func (s *Next_simple_param_Context) COMMA() antlr.TerminalNode {
	return s.GetToken(ServiceParserCOMMA, 0)
}

func (s *Next_simple_param_Context) Simple_param_() ISimple_param_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_param_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_param_Context)
}

func (s *Next_simple_param_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Next_simple_param_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Next_simple_param_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterNext_simple_param_(s)
	}
}

func (s *Next_simple_param_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitNext_simple_param_(s)
	}
}

func (p *ServiceParser) Next_simple_param_() (localctx INext_simple_param_Context) {
	localctx = NewNext_simple_param_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ServiceParserRULE_next_simple_param_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(ServiceParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Simple_param_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReal_base_type_list_Context is an interface to support dynamic dispatch.
type IReal_base_type_list_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type() IReal_base_typeContext

	// IsReal_base_type_list_Context differentiates from other interfaces.
	IsReal_base_type_list_Context()
}

type Real_base_type_list_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_base_type_list_Context() *Real_base_type_list_Context {
	var p = new(Real_base_type_list_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type_list_
	return p
}

func InitEmptyReal_base_type_list_Context(p *Real_base_type_list_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type_list_
}

func (*Real_base_type_list_Context) IsReal_base_type_list_Context() {}

func NewReal_base_type_list_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_base_type_list_Context {
	var p = new(Real_base_type_list_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_real_base_type_list_

	return p
}

func (s *Real_base_type_list_Context) GetParser() antlr.Parser { return s.parser }

func (s *Real_base_type_list_Context) Real_base_type() IReal_base_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_typeContext)
}

func (s *Real_base_type_list_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_base_type_list_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_base_type_list_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterReal_base_type_list_(s)
	}
}

func (s *Real_base_type_list_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitReal_base_type_list_(s)
	}
}

func (p *ServiceParser) Real_base_type_list_() (localctx IReal_base_type_list_Context) {
	localctx = NewReal_base_type_list_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ServiceParserRULE_real_base_type_list_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(ServiceParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(ServiceParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Real_base_type()
	}
	{
		p.SetState(245)
		p.Match(ServiceParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVoid_Context is an interface to support dynamic dispatch.
type IVoid_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVoid_Context differentiates from other interfaces.
	IsVoid_Context()
}

type Void_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoid_Context() *Void_Context {
	var p = new(Void_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_void_
	return p
}

func InitEmptyVoid_Context(p *Void_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_void_
}

func (*Void_Context) IsVoid_Context() {}

func NewVoid_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Void_Context {
	var p = new(Void_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_void_

	return p
}

func (s *Void_Context) GetParser() antlr.Parser { return s.parser }
func (s *Void_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Void_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Void_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterVoid_(s)
	}
}

func (s *Void_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitVoid_(s)
	}
}

func (p *ServiceParser) Void_() (localctx IVoid_Context) {
	localctx = NewVoid_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ServiceParserRULE_void_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(ServiceParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_param_Context is an interface to support dynamic dispatch.
type IMethod_param_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Single_struct_param() ISingle_struct_paramContext
	Struct_type_list() IStruct_type_listContext
	IDENTIFIER() antlr.TerminalNode

	// IsMethod_param_Context differentiates from other interfaces.
	IsMethod_param_Context()
}

type Method_param_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_param_Context() *Method_param_Context {
	var p = new(Method_param_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_param_
	return p
}

func InitEmptyMethod_param_Context(p *Method_param_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_param_
}

func (*Method_param_Context) IsMethod_param_Context() {}

func NewMethod_param_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_param_Context {
	var p = new(Method_param_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_method_param_

	return p
}

func (s *Method_param_Context) GetParser() antlr.Parser { return s.parser }

func (s *Method_param_Context) Single_struct_param() ISingle_struct_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_struct_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_struct_paramContext)
}

func (s *Method_param_Context) Struct_type_list() IStruct_type_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_type_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_type_listContext)
}

func (s *Method_param_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Method_param_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_param_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_param_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMethod_param_(s)
	}
}

func (s *Method_param_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMethod_param_(s)
	}
}

func (p *ServiceParser) Method_param_() (localctx IMethod_param_Context) {
	localctx = NewMethod_param_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ServiceParserRULE_method_param_)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Single_struct_param()
		}

	case ServiceParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Struct_type_list()
		}
		{
			p.SetState(251)
			p.Match(ServiceParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingle_struct_paramContext is an interface to support dynamic dispatch.
type ISingle_struct_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct_type() IStruct_typeContext
	IDENTIFIER() antlr.TerminalNode

	// IsSingle_struct_paramContext differentiates from other interfaces.
	IsSingle_struct_paramContext()
}

type Single_struct_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_struct_paramContext() *Single_struct_paramContext {
	var p = new(Single_struct_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_single_struct_param
	return p
}

func InitEmptySingle_struct_paramContext(p *Single_struct_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_single_struct_param
}

func (*Single_struct_paramContext) IsSingle_struct_paramContext() {}

func NewSingle_struct_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_struct_paramContext {
	var p = new(Single_struct_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_single_struct_param

	return p
}

func (s *Single_struct_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_struct_paramContext) Struct_type() IStruct_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_typeContext)
}

func (s *Single_struct_paramContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Single_struct_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_struct_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_struct_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterSingle_struct_param(s)
	}
}

func (s *Single_struct_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitSingle_struct_param(s)
	}
}

func (p *ServiceParser) Single_struct_param() (localctx ISingle_struct_paramContext) {
	localctx = NewSingle_struct_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ServiceParserRULE_single_struct_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Struct_type()
	}
	{
		p.SetState(256)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_param_Context is an interface to support dynamic dispatch.
type ISimple_param_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type_parm() IReal_base_type_parmContext
	Field_req() IField_reqContext
	Real_base_type_list_parm() IReal_base_type_list_parmContext

	// IsSimple_param_Context differentiates from other interfaces.
	IsSimple_param_Context()
}

type Simple_param_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_param_Context() *Simple_param_Context {
	var p = new(Simple_param_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_simple_param_
	return p
}

func InitEmptySimple_param_Context(p *Simple_param_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_simple_param_
}

func (*Simple_param_Context) IsSimple_param_Context() {}

func NewSimple_param_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_param_Context {
	var p = new(Simple_param_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_simple_param_

	return p
}

func (s *Simple_param_Context) GetParser() antlr.Parser { return s.parser }

func (s *Simple_param_Context) Real_base_type_parm() IReal_base_type_parmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_type_parmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_type_parmContext)
}

func (s *Simple_param_Context) Field_req() IField_reqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_reqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_reqContext)
}

func (s *Simple_param_Context) Real_base_type_list_parm() IReal_base_type_list_parmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_type_list_parmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_type_list_parmContext)
}

func (s *Simple_param_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_param_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_param_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterSimple_param_(s)
	}
}

func (s *Simple_param_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitSimple_param_(s)
	}
}

func (p *ServiceParser) Simple_param_() (localctx ISimple_param_Context) {
	localctx = NewSimple_param_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ServiceParserRULE_simple_param_)
	var _la int

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ServiceParserT__10 || _la == ServiceParserT__11 {
			{
				p.SetState(258)
				p.Field_req()
			}

		}
		{
			p.SetState(261)
			p.Real_base_type_parm()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ServiceParserT__10 || _la == ServiceParserT__11 {
			{
				p.SetState(262)
				p.Field_req()
			}

		}
		{
			p.SetState(265)
			p.Real_base_type_list_parm()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReal_base_type_parmContext is an interface to support dynamic dispatch.
type IReal_base_type_parmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type() IReal_base_typeContext
	IDENTIFIER() antlr.TerminalNode

	// IsReal_base_type_parmContext differentiates from other interfaces.
	IsReal_base_type_parmContext()
}

type Real_base_type_parmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_base_type_parmContext() *Real_base_type_parmContext {
	var p = new(Real_base_type_parmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type_parm
	return p
}

func InitEmptyReal_base_type_parmContext(p *Real_base_type_parmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type_parm
}

func (*Real_base_type_parmContext) IsReal_base_type_parmContext() {}

func NewReal_base_type_parmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_base_type_parmContext {
	var p = new(Real_base_type_parmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_real_base_type_parm

	return p
}

func (s *Real_base_type_parmContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_base_type_parmContext) Real_base_type() IReal_base_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_typeContext)
}

func (s *Real_base_type_parmContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Real_base_type_parmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_base_type_parmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_base_type_parmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterReal_base_type_parm(s)
	}
}

func (s *Real_base_type_parmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitReal_base_type_parm(s)
	}
}

func (p *ServiceParser) Real_base_type_parm() (localctx IReal_base_type_parmContext) {
	localctx = NewReal_base_type_parmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ServiceParserRULE_real_base_type_parm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Real_base_type()
	}
	{
		p.SetState(269)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReal_base_type_list_parmContext is an interface to support dynamic dispatch.
type IReal_base_type_list_parmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type_list_() IReal_base_type_list_Context
	IDENTIFIER() antlr.TerminalNode

	// IsReal_base_type_list_parmContext differentiates from other interfaces.
	IsReal_base_type_list_parmContext()
}

type Real_base_type_list_parmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_base_type_list_parmContext() *Real_base_type_list_parmContext {
	var p = new(Real_base_type_list_parmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type_list_parm
	return p
}

func InitEmptyReal_base_type_list_parmContext(p *Real_base_type_list_parmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type_list_parm
}

func (*Real_base_type_list_parmContext) IsReal_base_type_list_parmContext() {}

func NewReal_base_type_list_parmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_base_type_list_parmContext {
	var p = new(Real_base_type_list_parmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_real_base_type_list_parm

	return p
}

func (s *Real_base_type_list_parmContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_base_type_list_parmContext) Real_base_type_list_() IReal_base_type_list_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_type_list_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_type_list_Context)
}

func (s *Real_base_type_list_parmContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Real_base_type_list_parmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_base_type_list_parmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_base_type_list_parmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterReal_base_type_list_parm(s)
	}
}

func (s *Real_base_type_list_parmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitReal_base_type_list_parm(s)
	}
}

func (p *ServiceParser) Real_base_type_list_parm() (localctx IReal_base_type_list_parmContext) {
	localctx = NewReal_base_type_list_parmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ServiceParserRULE_real_base_type_list_parm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Real_base_type_list_()
	}
	{
		p.SetState(272)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INot_loginContext is an interface to support dynamic dispatch.
type INot_loginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNot_loginContext differentiates from other interfaces.
	IsNot_loginContext()
}

type Not_loginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_loginContext() *Not_loginContext {
	var p = new(Not_loginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_not_login
	return p
}

func InitEmptyNot_loginContext(p *Not_loginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_not_login
}

func (*Not_loginContext) IsNot_loginContext() {}

func NewNot_loginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_loginContext {
	var p = new(Not_loginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_not_login

	return p
}

func (s *Not_loginContext) GetParser() antlr.Parser { return s.parser }
func (s *Not_loginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_loginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_loginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterNot_login(s)
	}
}

func (s *Not_loginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitNot_login(s)
	}
}

func (p *ServiceParser) Not_login() (localctx INot_loginContext) {
	localctx = NewNot_loginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ServiceParserRULE_not_login)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(ServiceParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_descriptionContext is an interface to support dynamic dispatch.
type IMethod_descriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Method_description_content() IMethod_description_contentContext

	// IsMethod_descriptionContext differentiates from other interfaces.
	IsMethod_descriptionContext()
}

type Method_descriptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_descriptionContext() *Method_descriptionContext {
	var p = new(Method_descriptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_description
	return p
}

func InitEmptyMethod_descriptionContext(p *Method_descriptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_description
}

func (*Method_descriptionContext) IsMethod_descriptionContext() {}

func NewMethod_descriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_descriptionContext {
	var p = new(Method_descriptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_method_description

	return p
}

func (s *Method_descriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_descriptionContext) Method_description_content() IMethod_description_contentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_description_contentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_description_contentContext)
}

func (s *Method_descriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_descriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_descriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMethod_description(s)
	}
}

func (s *Method_descriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMethod_description(s)
	}
}

func (p *ServiceParser) Method_description() (localctx IMethod_descriptionContext) {
	localctx = NewMethod_descriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ServiceParserRULE_method_description)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(ServiceParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(ServiceParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(ServiceParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Method_description_content()
	}
	{
		p.SetState(280)
		p.Match(ServiceParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_description_contentContext is an interface to support dynamic dispatch.
type IMethod_description_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnnotation_value() []IAnnotation_valueContext
	Annotation_value(i int) IAnnotation_valueContext

	// IsMethod_description_contentContext differentiates from other interfaces.
	IsMethod_description_contentContext()
}

type Method_description_contentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_description_contentContext() *Method_description_contentContext {
	var p = new(Method_description_contentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_description_content
	return p
}

func InitEmptyMethod_description_contentContext(p *Method_description_contentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_method_description_content
}

func (*Method_description_contentContext) IsMethod_description_contentContext() {}

func NewMethod_description_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_description_contentContext {
	var p = new(Method_description_contentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_method_description_content

	return p
}

func (s *Method_description_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_description_contentContext) AllAnnotation_value() []IAnnotation_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotation_valueContext); ok {
			len++
		}
	}

	tst := make([]IAnnotation_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotation_valueContext); ok {
			tst[i] = t.(IAnnotation_valueContext)
			i++
		}
	}

	return tst
}

func (s *Method_description_contentContext) Annotation_value(i int) IAnnotation_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_valueContext)
}

func (s *Method_description_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_description_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_description_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMethod_description_content(s)
	}
}

func (s *Method_description_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMethod_description_content(s)
	}
}

func (p *ServiceParser) Method_description_content() (localctx IMethod_description_contentContext) {
	localctx = NewMethod_description_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ServiceParserRULE_method_description_content)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&551366426624) != 0 {
		{
			p.SetState(282)
			p.Annotation_value()
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_annotationsContext is an interface to support dynamic dispatch.
type IType_annotationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_annotation() []IType_annotationContext
	Type_annotation(i int) IType_annotationContext

	// IsType_annotationsContext differentiates from other interfaces.
	IsType_annotationsContext()
}

type Type_annotationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationsContext() *Type_annotationsContext {
	var p = new(Type_annotationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_type_annotations
	return p
}

func InitEmptyType_annotationsContext(p *Type_annotationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_type_annotations
}

func (*Type_annotationsContext) IsType_annotationsContext() {}

func NewType_annotationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationsContext {
	var p = new(Type_annotationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_type_annotations

	return p
}

func (s *Type_annotationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationsContext) AllType_annotation() []IType_annotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_annotationContext); ok {
			len++
		}
	}

	tst := make([]IType_annotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_annotationContext); ok {
			tst[i] = t.(IType_annotationContext)
			i++
		}
	}

	return tst
}

func (s *Type_annotationsContext) Type_annotation(i int) IType_annotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationContext)
}

func (s *Type_annotationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterType_annotations(s)
	}
}

func (s *Type_annotationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitType_annotations(s)
	}
}

func (p *ServiceParser) Type_annotations() (localctx IType_annotationsContext) {
	localctx = NewType_annotationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ServiceParserRULE_type_annotations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(ServiceParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ServiceParserIDENTIFIER {
		{
			p.SetState(289)
			p.Type_annotation()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(295)
		p.Match(ServiceParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_annotationContext is an interface to support dynamic dispatch.
type IType_annotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Annotation_value() IAnnotation_valueContext
	List_separator() IList_separatorContext

	// IsType_annotationContext differentiates from other interfaces.
	IsType_annotationContext()
}

type Type_annotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationContext() *Type_annotationContext {
	var p = new(Type_annotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_type_annotation
	return p
}

func InitEmptyType_annotationContext(p *Type_annotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_type_annotation
}

func (*Type_annotationContext) IsType_annotationContext() {}

func NewType_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationContext {
	var p = new(Type_annotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_type_annotation

	return p
}

func (s *Type_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Type_annotationContext) Annotation_value() IAnnotation_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_valueContext)
}

func (s *Type_annotationContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Type_annotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterType_annotation(s)
	}
}

func (s *Type_annotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitType_annotation(s)
	}
}

func (p *ServiceParser) Type_annotation() (localctx IType_annotationContext) {
	localctx = NewType_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ServiceParserRULE_type_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__1 {
		{
			p.SetState(298)
			p.Match(ServiceParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Annotation_value()
		}

	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
		{
			p.SetState(302)
			p.List_separator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_annotationsContext is an interface to support dynamic dispatch.
type IField_annotationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField_annotation() []IField_annotationContext
	Field_annotation(i int) IField_annotationContext

	// IsField_annotationsContext differentiates from other interfaces.
	IsField_annotationsContext()
}

type Field_annotationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_annotationsContext() *Field_annotationsContext {
	var p = new(Field_annotationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_annotations
	return p
}

func InitEmptyField_annotationsContext(p *Field_annotationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_annotations
}

func (*Field_annotationsContext) IsField_annotationsContext() {}

func NewField_annotationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_annotationsContext {
	var p = new(Field_annotationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_field_annotations

	return p
}

func (s *Field_annotationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_annotationsContext) AllField_annotation() []IField_annotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IField_annotationContext); ok {
			len++
		}
	}

	tst := make([]IField_annotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IField_annotationContext); ok {
			tst[i] = t.(IField_annotationContext)
			i++
		}
	}

	return tst
}

func (s *Field_annotationsContext) Field_annotation(i int) IField_annotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_annotationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_annotationContext)
}

func (s *Field_annotationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_annotationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_annotationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterField_annotations(s)
	}
}

func (s *Field_annotationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitField_annotations(s)
	}
}

func (p *ServiceParser) Field_annotations() (localctx IField_annotationsContext) {
	localctx = NewField_annotationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ServiceParserRULE_field_annotations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(ServiceParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ServiceParserIDENTIFIER {
		{
			p.SetState(306)
			p.Field_annotation()
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(312)
		p.Match(ServiceParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_annotationContext is an interface to support dynamic dispatch.
type IField_annotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	List_separator() IList_separatorContext

	// IsField_annotationContext differentiates from other interfaces.
	IsField_annotationContext()
}

type Field_annotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_annotationContext() *Field_annotationContext {
	var p = new(Field_annotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_annotation
	return p
}

func InitEmptyField_annotationContext(p *Field_annotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_annotation
}

func (*Field_annotationContext) IsField_annotationContext() {}

func NewField_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_annotationContext {
	var p = new(Field_annotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_field_annotation

	return p
}

func (s *Field_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_annotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Field_annotationContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ServiceParserLITERAL, 0)
}

func (s *Field_annotationContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Field_annotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_annotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_annotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterField_annotation(s)
	}
}

func (s *Field_annotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitField_annotation(s)
	}
}

func (p *ServiceParser) Field_annotation() (localctx IField_annotationContext) {
	localctx = NewField_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ServiceParserRULE_field_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Match(ServiceParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Match(ServiceParserLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
		{
			p.SetState(317)
			p.List_separator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotation_valueContext is an interface to support dynamic dispatch.
type IAnnotation_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	LITERAL() antlr.TerminalNode

	// IsAnnotation_valueContext differentiates from other interfaces.
	IsAnnotation_valueContext()
}

type Annotation_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_valueContext() *Annotation_valueContext {
	var p = new(Annotation_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_annotation_value
	return p
}

func InitEmptyAnnotation_valueContext(p *Annotation_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_annotation_value
}

func (*Annotation_valueContext) IsAnnotation_valueContext() {}

func NewAnnotation_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_valueContext {
	var p = new(Annotation_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_annotation_value

	return p
}

func (s *Annotation_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_valueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Annotation_valueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ServiceParserLITERAL, 0)
}

func (s *Annotation_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterAnnotation_value(s)
	}
}

func (s *Annotation_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitAnnotation_value(s)
	}
}

func (p *ServiceParser) Annotation_value() (localctx IAnnotation_valueContext) {
	localctx = NewAnnotation_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ServiceParserRULE_annotation_value)
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserINTEGER, ServiceParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Integer()
		}

	case ServiceParserLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.Match(ServiceParserLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Base_type() IBase_typeContext
	Struct_type() IStruct_typeContext
	Container_type() IContainer_typeContext

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_type
	return p
}

func InitEmptyField_typeContext(p *Field_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_field_type
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) Base_type() IBase_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBase_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBase_typeContext)
}

func (s *Field_typeContext) Struct_type() IStruct_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_typeContext)
}

func (s *Field_typeContext) Container_type() IContainer_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_typeContext)
}

func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (p *ServiceParser) Field_type() (localctx IField_typeContext) {
	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ServiceParserRULE_field_type)
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserTYPE_BOOL, ServiceParserTYPE_BYTE, ServiceParserTYPE_I16, ServiceParserTYPE_I32, ServiceParserTYPE_I64, ServiceParserTYPE_DOUBLE, ServiceParserTYPE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.Base_type()
		}

	case ServiceParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.Struct_type()
		}

	case ServiceParserT__17, ServiceParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)
			p.Container_type()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBase_typeContext is an interface to support dynamic dispatch.
type IBase_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type() IReal_base_typeContext
	Type_annotations() IType_annotationsContext

	// IsBase_typeContext differentiates from other interfaces.
	IsBase_typeContext()
}

type Base_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBase_typeContext() *Base_typeContext {
	var p = new(Base_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_base_type
	return p
}

func InitEmptyBase_typeContext(p *Base_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_base_type
}

func (*Base_typeContext) IsBase_typeContext() {}

func NewBase_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Base_typeContext {
	var p = new(Base_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_base_type

	return p
}

func (s *Base_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Base_typeContext) Real_base_type() IReal_base_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_typeContext)
}

func (s *Base_typeContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Base_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Base_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Base_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterBase_type(s)
	}
}

func (s *Base_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitBase_type(s)
	}
}

func (p *ServiceParser) Base_type() (localctx IBase_typeContext) {
	localctx = NewBase_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ServiceParserRULE_base_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Real_base_type()
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(330)
			p.Type_annotations()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainer_typeContext is an interface to support dynamic dispatch.
type IContainer_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Map_type() IMap_typeContext
	List_type() IList_typeContext
	Type_annotations() IType_annotationsContext

	// IsContainer_typeContext differentiates from other interfaces.
	IsContainer_typeContext()
}

type Container_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_typeContext() *Container_typeContext {
	var p = new(Container_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_container_type
	return p
}

func InitEmptyContainer_typeContext(p *Container_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_container_type
}

func (*Container_typeContext) IsContainer_typeContext() {}

func NewContainer_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_typeContext {
	var p = new(Container_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_container_type

	return p
}

func (s *Container_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_typeContext) Map_type() IMap_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMap_typeContext)
}

func (s *Container_typeContext) List_type() IList_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_typeContext)
}

func (s *Container_typeContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Container_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterContainer_type(s)
	}
}

func (s *Container_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitContainer_type(s)
	}
}

func (p *ServiceParser) Container_type() (localctx IContainer_typeContext) {
	localctx = NewContainer_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ServiceParserRULE_container_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserT__23:
		{
			p.SetState(333)
			p.Map_type()
		}

	case ServiceParserT__17:
		{
			p.SetState(334)
			p.List_type()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__13 {
		{
			p.SetState(337)
			p.Type_annotations()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_typeContext is an interface to support dynamic dispatch.
type IStruct_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsStruct_typeContext differentiates from other interfaces.
	IsStruct_typeContext()
}

type Struct_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_typeContext() *Struct_typeContext {
	var p = new(Struct_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_struct_type
	return p
}

func InitEmptyStruct_typeContext(p *Struct_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_struct_type
}

func (*Struct_typeContext) IsStruct_typeContext() {}

func NewStruct_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_typeContext {
	var p = new(Struct_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_struct_type

	return p
}

func (s *Struct_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Struct_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterStruct_type(s)
	}
}

func (s *Struct_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitStruct_type(s)
	}
}

func (p *ServiceParser) Struct_type() (localctx IStruct_typeContext) {
	localctx = NewStruct_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ServiceParserRULE_struct_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(ServiceParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Match(ServiceParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMap_typeContext is an interface to support dynamic dispatch.
type IMap_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Map_key_type() IMap_key_typeContext
	COMMA() antlr.TerminalNode
	Field_type() IField_typeContext

	// IsMap_typeContext differentiates from other interfaces.
	IsMap_typeContext()
}

type Map_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_typeContext() *Map_typeContext {
	var p = new(Map_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_map_type
	return p
}

func InitEmptyMap_typeContext(p *Map_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_map_type
}

func (*Map_typeContext) IsMap_typeContext() {}

func NewMap_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_typeContext {
	var p = new(Map_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_map_type

	return p
}

func (s *Map_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_typeContext) Map_key_type() IMap_key_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_key_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMap_key_typeContext)
}

func (s *Map_typeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ServiceParserCOMMA, 0)
}

func (s *Map_typeContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Map_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMap_type(s)
	}
}

func (s *Map_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMap_type(s)
	}
}

func (p *ServiceParser) Map_type() (localctx IMap_typeContext) {
	localctx = NewMap_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ServiceParserRULE_map_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(ServiceParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(ServiceParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Map_key_type()
	}
	{
		p.SetState(346)
		p.Match(ServiceParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Field_type()
	}
	{
		p.SetState(348)
		p.Match(ServiceParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_typeContext is an interface to support dynamic dispatch.
type IList_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext

	// IsList_typeContext differentiates from other interfaces.
	IsList_typeContext()
}

type List_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_typeContext() *List_typeContext {
	var p = new(List_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_list_type
	return p
}

func InitEmptyList_typeContext(p *List_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_list_type
}

func (*List_typeContext) IsList_typeContext() {}

func NewList_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_typeContext {
	var p = new(List_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_list_type

	return p
}

func (s *List_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *List_typeContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *List_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterList_type(s)
	}
}

func (s *List_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitList_type(s)
	}
}

func (p *ServiceParser) List_type() (localctx IList_typeContext) {
	localctx = NewList_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ServiceParserRULE_list_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(ServiceParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(ServiceParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.Field_type()
	}
	{
		p.SetState(353)
		p.Match(ServiceParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_valueContext is an interface to support dynamic dispatch.
type IConst_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	DOUBLE() antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Const_list() IConst_listContext
	Const_map() IConst_mapContext

	// IsConst_valueContext differentiates from other interfaces.
	IsConst_valueContext()
}

type Const_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_valueContext() *Const_valueContext {
	var p = new(Const_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_value
	return p
}

func InitEmptyConst_valueContext(p *Const_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_value
}

func (*Const_valueContext) IsConst_valueContext() {}

func NewConst_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_valueContext {
	var p = new(Const_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_const_value

	return p
}

func (s *Const_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_valueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Const_valueContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ServiceParserDOUBLE, 0)
}

func (s *Const_valueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ServiceParserLITERAL, 0)
}

func (s *Const_valueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENTIFIER, 0)
}

func (s *Const_valueContext) Const_list() IConst_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_listContext)
}

func (s *Const_valueContext) Const_map() IConst_mapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_mapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_mapContext)
}

func (s *Const_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterConst_value(s)
	}
}

func (s *Const_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitConst_value(s)
	}
}

func (p *ServiceParser) Const_value() (localctx IConst_valueContext) {
	localctx = NewConst_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ServiceParserRULE_const_value)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserINTEGER, ServiceParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Integer()
		}

	case ServiceParserDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Match(ServiceParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ServiceParserLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(357)
			p.Match(ServiceParserLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ServiceParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(358)
			p.Match(ServiceParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ServiceParserT__24:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(359)
			p.Const_list()
		}

	case ServiceParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(360)
			p.Const_map()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	HEX_INTEGER() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ServiceParserINTEGER, 0)
}

func (s *IntegerContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(ServiceParserHEX_INTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *ServiceParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ServiceParserRULE_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ServiceParserINTEGER || _la == ServiceParserHEX_INTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_listContext is an interface to support dynamic dispatch.
type IConst_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_value() []IConst_valueContext
	Const_value(i int) IConst_valueContext
	AllList_separator() []IList_separatorContext
	List_separator(i int) IList_separatorContext

	// IsConst_listContext differentiates from other interfaces.
	IsConst_listContext()
}

type Const_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_listContext() *Const_listContext {
	var p = new(Const_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_list
	return p
}

func InitEmptyConst_listContext(p *Const_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_list
}

func (*Const_listContext) IsConst_listContext() {}

func NewConst_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_listContext {
	var p = new(Const_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_const_list

	return p
}

func (s *Const_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_listContext) AllConst_value() []IConst_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_valueContext); ok {
			len++
		}
	}

	tst := make([]IConst_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_valueContext); ok {
			tst[i] = t.(IConst_valueContext)
			i++
		}
	}

	return tst
}

func (s *Const_listContext) Const_value(i int) IConst_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_listContext) AllList_separator() []IList_separatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_separatorContext); ok {
			len++
		}
	}

	tst := make([]IList_separatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_separatorContext); ok {
			tst[i] = t.(IList_separatorContext)
			i++
		}
	}

	return tst
}

func (s *Const_listContext) List_separator(i int) IList_separatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterConst_list(s)
	}
}

func (s *Const_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitConst_list(s)
	}
}

func (p *ServiceParser) Const_list() (localctx IConst_listContext) {
	localctx = NewConst_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ServiceParserRULE_const_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(ServiceParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1653059092736) != 0 {
		{
			p.SetState(366)
			p.Const_value()
		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
			{
				p.SetState(367)
				p.List_separator()
			}

		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(375)
		p.Match(ServiceParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_map_entryContext is an interface to support dynamic dispatch.
type IConst_map_entryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_value() []IConst_valueContext
	Const_value(i int) IConst_valueContext
	List_separator() IList_separatorContext

	// IsConst_map_entryContext differentiates from other interfaces.
	IsConst_map_entryContext()
}

type Const_map_entryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_map_entryContext() *Const_map_entryContext {
	var p = new(Const_map_entryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_map_entry
	return p
}

func InitEmptyConst_map_entryContext(p *Const_map_entryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_map_entry
}

func (*Const_map_entryContext) IsConst_map_entryContext() {}

func NewConst_map_entryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_map_entryContext {
	var p = new(Const_map_entryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_const_map_entry

	return p
}

func (s *Const_map_entryContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_map_entryContext) AllConst_value() []IConst_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_valueContext); ok {
			len++
		}
	}

	tst := make([]IConst_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_valueContext); ok {
			tst[i] = t.(IConst_valueContext)
			i++
		}
	}

	return tst
}

func (s *Const_map_entryContext) Const_value(i int) IConst_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_map_entryContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_map_entryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_map_entryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_map_entryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterConst_map_entry(s)
	}
}

func (s *Const_map_entryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitConst_map_entry(s)
	}
}

func (p *ServiceParser) Const_map_entry() (localctx IConst_map_entryContext) {
	localctx = NewConst_map_entryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ServiceParserRULE_const_map_entry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Const_value()
	}
	{
		p.SetState(378)
		p.Match(ServiceParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Const_value()
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ServiceParserT__27 || _la == ServiceParserCOMMA {
		{
			p.SetState(380)
			p.List_separator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_mapContext is an interface to support dynamic dispatch.
type IConst_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_map_entry() []IConst_map_entryContext
	Const_map_entry(i int) IConst_map_entryContext

	// IsConst_mapContext differentiates from other interfaces.
	IsConst_mapContext()
}

type Const_mapContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_mapContext() *Const_mapContext {
	var p = new(Const_mapContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_map
	return p
}

func InitEmptyConst_mapContext(p *Const_mapContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_const_map
}

func (*Const_mapContext) IsConst_mapContext() {}

func NewConst_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_mapContext {
	var p = new(Const_mapContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_const_map

	return p
}

func (s *Const_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_mapContext) AllConst_map_entry() []IConst_map_entryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_map_entryContext); ok {
			len++
		}
	}

	tst := make([]IConst_map_entryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_map_entryContext); ok {
			tst[i] = t.(IConst_map_entryContext)
			i++
		}
	}

	return tst
}

func (s *Const_mapContext) Const_map_entry(i int) IConst_map_entryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_map_entryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_map_entryContext)
}

func (s *Const_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterConst_map(s)
	}
}

func (s *Const_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitConst_map(s)
	}
}

func (p *ServiceParser) Const_map() (localctx IConst_mapContext) {
	localctx = NewConst_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ServiceParserRULE_const_map)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(ServiceParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1653059092736) != 0 {
		{
			p.SetState(384)
			p.Const_map_entry()
		}

		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(390)
		p.Match(ServiceParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_separatorContext is an interface to support dynamic dispatch.
type IList_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode

	// IsList_separatorContext differentiates from other interfaces.
	IsList_separatorContext()
}

type List_separatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_separatorContext() *List_separatorContext {
	var p = new(List_separatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_list_separator
	return p
}

func InitEmptyList_separatorContext(p *List_separatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_list_separator
}

func (*List_separatorContext) IsList_separatorContext() {}

func NewList_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_separatorContext {
	var p = new(List_separatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_list_separator

	return p
}

func (s *List_separatorContext) GetParser() antlr.Parser { return s.parser }

func (s *List_separatorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ServiceParserCOMMA, 0)
}

func (s *List_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_separatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterList_separator(s)
	}
}

func (s *List_separatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitList_separator(s)
	}
}

func (p *ServiceParser) List_separator() (localctx IList_separatorContext) {
	localctx = NewList_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ServiceParserRULE_list_separator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ServiceParserT__27 || _la == ServiceParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReal_base_typeContext is an interface to support dynamic dispatch.
type IReal_base_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_BOOL() antlr.TerminalNode
	TYPE_BYTE() antlr.TerminalNode
	TYPE_I16() antlr.TerminalNode
	TYPE_I32() antlr.TerminalNode
	TYPE_I64() antlr.TerminalNode
	TYPE_DOUBLE() antlr.TerminalNode
	TYPE_STRING() antlr.TerminalNode

	// IsReal_base_typeContext differentiates from other interfaces.
	IsReal_base_typeContext()
}

type Real_base_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_base_typeContext() *Real_base_typeContext {
	var p = new(Real_base_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type
	return p
}

func InitEmptyReal_base_typeContext(p *Real_base_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_real_base_type
}

func (*Real_base_typeContext) IsReal_base_typeContext() {}

func NewReal_base_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_base_typeContext {
	var p = new(Real_base_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_real_base_type

	return p
}

func (s *Real_base_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_base_typeContext) TYPE_BOOL() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_BOOL, 0)
}

func (s *Real_base_typeContext) TYPE_BYTE() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_BYTE, 0)
}

func (s *Real_base_typeContext) TYPE_I16() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_I16, 0)
}

func (s *Real_base_typeContext) TYPE_I32() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_I32, 0)
}

func (s *Real_base_typeContext) TYPE_I64() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_I64, 0)
}

func (s *Real_base_typeContext) TYPE_DOUBLE() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_DOUBLE, 0)
}

func (s *Real_base_typeContext) TYPE_STRING() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_STRING, 0)
}

func (s *Real_base_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_base_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_base_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterReal_base_type(s)
	}
}

func (s *Real_base_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitReal_base_type(s)
	}
}

func (p *ServiceParser) Real_base_type() (localctx IReal_base_typeContext) {
	localctx = NewReal_base_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ServiceParserRULE_real_base_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&545460846592) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMap_key_typeContext is an interface to support dynamic dispatch.
type IMap_key_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_BYTE() antlr.TerminalNode
	TYPE_I16() antlr.TerminalNode
	TYPE_I32() antlr.TerminalNode
	TYPE_I64() antlr.TerminalNode
	TYPE_DOUBLE() antlr.TerminalNode
	TYPE_STRING() antlr.TerminalNode

	// IsMap_key_typeContext differentiates from other interfaces.
	IsMap_key_typeContext()
}

type Map_key_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_key_typeContext() *Map_key_typeContext {
	var p = new(Map_key_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_map_key_type
	return p
}

func InitEmptyMap_key_typeContext(p *Map_key_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_map_key_type
}

func (*Map_key_typeContext) IsMap_key_typeContext() {}

func NewMap_key_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_key_typeContext {
	var p = new(Map_key_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_map_key_type

	return p
}

func (s *Map_key_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_key_typeContext) TYPE_BYTE() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_BYTE, 0)
}

func (s *Map_key_typeContext) TYPE_I16() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_I16, 0)
}

func (s *Map_key_typeContext) TYPE_I32() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_I32, 0)
}

func (s *Map_key_typeContext) TYPE_I64() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_I64, 0)
}

func (s *Map_key_typeContext) TYPE_DOUBLE() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_DOUBLE, 0)
}

func (s *Map_key_typeContext) TYPE_STRING() antlr.TerminalNode {
	return s.GetToken(ServiceParserTYPE_STRING, 0)
}

func (s *Map_key_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_key_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_key_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.EnterMap_key_type(s)
	}
}

func (s *Map_key_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ServiceListener); ok {
		listenerT.ExitMap_key_type(s)
	}
}

func (p *ServiceParser) Map_key_type() (localctx IMap_key_typeContext) {
	localctx = NewMap_key_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ServiceParserRULE_map_key_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541165879296) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

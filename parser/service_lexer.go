// Code generated from Service.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ServiceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ServiceLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func servicelexerLexerInit() {
	staticData := &ServiceLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "INTEGER", "HEX_INTEGER", "DOUBLE", "TYPE_BOOL",
		"TYPE_BYTE", "TYPE_I16", "TYPE_I32", "TYPE_I64", "TYPE_DOUBLE", "TYPE_STRING",
		"LITERAL", "ESC_SEQ", "IDENTIFIER", "COMMA", "LETTER", "DIGIT", "HEX_DIGIT",
		"WS", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 414, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 3, 28, 246, 8, 28, 1, 28, 4, 28,
		249, 8, 28, 11, 28, 12, 28, 250, 1, 29, 3, 29, 254, 8, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 4, 29, 260, 8, 29, 11, 29, 12, 29, 261, 1, 30, 3, 30, 265,
		8, 30, 1, 30, 4, 30, 268, 8, 30, 11, 30, 12, 30, 269, 1, 30, 1, 30, 4,
		30, 274, 8, 30, 11, 30, 12, 30, 275, 3, 30, 278, 8, 30, 1, 30, 1, 30, 4,
		30, 282, 8, 30, 11, 30, 12, 30, 283, 3, 30, 286, 8, 30, 1, 30, 1, 30, 3,
		30, 290, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 5,
		38, 331, 8, 38, 10, 38, 12, 38, 334, 9, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		5, 38, 340, 8, 38, 10, 38, 12, 38, 343, 9, 38, 1, 38, 3, 38, 346, 8, 38,
		1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 3, 40, 353, 8, 40, 1, 40, 1, 40, 1,
		40, 5, 40, 358, 8, 40, 10, 40, 12, 40, 361, 9, 40, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 3, 44, 371, 8, 44, 1, 45, 1, 45, 1,
		45, 1, 45, 4, 45, 377, 8, 45, 11, 45, 12, 45, 378, 1, 45, 1, 45, 1, 46,
		1, 46, 1, 46, 3, 46, 386, 8, 46, 1, 46, 5, 46, 389, 8, 46, 10, 46, 12,
		46, 392, 9, 46, 1, 46, 3, 46, 395, 8, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		47, 1, 47, 1, 47, 1, 47, 5, 47, 405, 8, 47, 10, 47, 12, 47, 408, 9, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 406, 0, 48, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 0, 81,
		40, 83, 41, 85, 0, 87, 0, 89, 0, 91, 42, 93, 43, 95, 44, 1, 0, 10, 2, 0,
		43, 43, 45, 45, 2, 0, 69, 69, 101, 101, 2, 0, 34, 34, 92, 92, 2, 0, 39,
		39, 92, 92, 6, 0, 34, 34, 39, 39, 92, 92, 110, 110, 114, 114, 116, 116,
		2, 0, 46, 46, 95, 95, 2, 0, 65, 90, 97, 122, 2, 0, 65, 70, 97, 102, 2,
		0, 9, 9, 32, 32, 1, 0, 10, 10, 437, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0,
		0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0,
		0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0,
		0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1,
		0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 1, 97,
		1, 0, 0, 0, 3, 109, 1, 0, 0, 0, 5, 111, 1, 0, 0, 0, 7, 121, 1, 0, 0, 0,
		9, 131, 1, 0, 0, 0, 11, 136, 1, 0, 0, 0, 13, 142, 1, 0, 0, 0, 15, 149,
		1, 0, 0, 0, 17, 151, 1, 0, 0, 0, 19, 153, 1, 0, 0, 0, 21, 161, 1, 0, 0,
		0, 23, 170, 1, 0, 0, 0, 25, 179, 1, 0, 0, 0, 27, 184, 1, 0, 0, 0, 29, 186,
		1, 0, 0, 0, 31, 188, 1, 0, 0, 0, 33, 192, 1, 0, 0, 0, 35, 196, 1, 0, 0,
		0, 37, 201, 1, 0, 0, 0, 39, 203, 1, 0, 0, 0, 41, 205, 1, 0, 0, 0, 43, 210,
		1, 0, 0, 0, 45, 220, 1, 0, 0, 0, 47, 232, 1, 0, 0, 0, 49, 236, 1, 0, 0,
		0, 51, 238, 1, 0, 0, 0, 53, 240, 1, 0, 0, 0, 55, 242, 1, 0, 0, 0, 57, 245,
		1, 0, 0, 0, 59, 253, 1, 0, 0, 0, 61, 264, 1, 0, 0, 0, 63, 291, 1, 0, 0,
		0, 65, 296, 1, 0, 0, 0, 67, 301, 1, 0, 0, 0, 69, 305, 1, 0, 0, 0, 71, 309,
		1, 0, 0, 0, 73, 313, 1, 0, 0, 0, 75, 320, 1, 0, 0, 0, 77, 345, 1, 0, 0,
		0, 79, 347, 1, 0, 0, 0, 81, 352, 1, 0, 0, 0, 83, 362, 1, 0, 0, 0, 85, 364,
		1, 0, 0, 0, 87, 366, 1, 0, 0, 0, 89, 370, 1, 0, 0, 0, 91, 376, 1, 0, 0,
		0, 93, 385, 1, 0, 0, 0, 95, 400, 1, 0, 0, 0, 97, 98, 5, 119, 0, 0, 98,
		99, 5, 105, 0, 0, 99, 100, 5, 116, 0, 0, 100, 101, 5, 104, 0, 0, 101, 102,
		5, 95, 0, 0, 102, 103, 5, 99, 0, 0, 103, 104, 5, 108, 0, 0, 104, 105, 5,
		105, 0, 0, 105, 106, 5, 101, 0, 0, 106, 107, 5, 110, 0, 0, 107, 108, 5,
		116, 0, 0, 108, 2, 1, 0, 0, 0, 109, 110, 5, 61, 0, 0, 110, 4, 1, 0, 0,
		0, 111, 112, 5, 110, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5, 109, 0,
		0, 114, 115, 5, 101, 0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5, 112, 0,
		0, 117, 118, 5, 97, 0, 0, 118, 119, 5, 99, 0, 0, 119, 120, 5, 101, 0, 0,
		120, 6, 1, 0, 0, 0, 121, 122, 5, 103, 0, 0, 122, 123, 5, 111, 0, 0, 123,
		124, 5, 95, 0, 0, 124, 125, 5, 105, 0, 0, 125, 126, 5, 109, 0, 0, 126,
		127, 5, 112, 0, 0, 127, 128, 5, 111, 0, 0, 128, 129, 5, 114, 0, 0, 129,
		130, 5, 116, 0, 0, 130, 8, 1, 0, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133,
		5, 114, 0, 0, 133, 134, 5, 117, 0, 0, 134, 135, 5, 101, 0, 0, 135, 10,
		1, 0, 0, 0, 136, 137, 5, 102, 0, 0, 137, 138, 5, 97, 0, 0, 138, 139, 5,
		108, 0, 0, 139, 140, 5, 115, 0, 0, 140, 141, 5, 101, 0, 0, 141, 12, 1,
		0, 0, 0, 142, 143, 5, 115, 0, 0, 143, 144, 5, 116, 0, 0, 144, 145, 5, 114,
		0, 0, 145, 146, 5, 117, 0, 0, 146, 147, 5, 99, 0, 0, 147, 148, 5, 116,
		0, 0, 148, 14, 1, 0, 0, 0, 149, 150, 5, 123, 0, 0, 150, 16, 1, 0, 0, 0,
		151, 152, 5, 125, 0, 0, 152, 18, 1, 0, 0, 0, 153, 154, 5, 115, 0, 0, 154,
		155, 5, 101, 0, 0, 155, 156, 5, 114, 0, 0, 156, 157, 5, 118, 0, 0, 157,
		158, 5, 105, 0, 0, 158, 159, 5, 99, 0, 0, 159, 160, 5, 101, 0, 0, 160,
		20, 1, 0, 0, 0, 161, 162, 5, 114, 0, 0, 162, 163, 5, 101, 0, 0, 163, 164,
		5, 113, 0, 0, 164, 165, 5, 117, 0, 0, 165, 166, 5, 105, 0, 0, 166, 167,
		5, 114, 0, 0, 167, 168, 5, 101, 0, 0, 168, 169, 5, 100, 0, 0, 169, 22,
		1, 0, 0, 0, 170, 171, 5, 111, 0, 0, 171, 172, 5, 112, 0, 0, 172, 173, 5,
		116, 0, 0, 173, 174, 5, 105, 0, 0, 174, 175, 5, 111, 0, 0, 175, 176, 5,
		110, 0, 0, 176, 177, 5, 97, 0, 0, 177, 178, 5, 108, 0, 0, 178, 24, 1, 0,
		0, 0, 179, 180, 5, 80, 0, 0, 180, 181, 5, 79, 0, 0, 181, 182, 5, 83, 0,
		0, 182, 183, 5, 84, 0, 0, 183, 26, 1, 0, 0, 0, 184, 185, 5, 40, 0, 0, 185,
		28, 1, 0, 0, 0, 186, 187, 5, 41, 0, 0, 187, 30, 1, 0, 0, 0, 188, 189, 5,
		71, 0, 0, 189, 190, 5, 69, 0, 0, 190, 191, 5, 84, 0, 0, 191, 32, 1, 0,
		0, 0, 192, 193, 5, 117, 0, 0, 193, 194, 5, 114, 0, 0, 194, 195, 5, 108,
		0, 0, 195, 34, 1, 0, 0, 0, 196, 197, 5, 108, 0, 0, 197, 198, 5, 105, 0,
		0, 198, 199, 5, 115, 0, 0, 199, 200, 5, 116, 0, 0, 200, 36, 1, 0, 0, 0,
		201, 202, 5, 60, 0, 0, 202, 38, 1, 0, 0, 0, 203, 204, 5, 62, 0, 0, 204,
		40, 1, 0, 0, 0, 205, 206, 5, 118, 0, 0, 206, 207, 5, 111, 0, 0, 207, 208,
		5, 105, 0, 0, 208, 209, 5, 100, 0, 0, 209, 42, 1, 0, 0, 0, 210, 211, 5,
		110, 0, 0, 211, 212, 5, 111, 0, 0, 212, 213, 5, 116, 0, 0, 213, 214, 5,
		95, 0, 0, 214, 215, 5, 108, 0, 0, 215, 216, 5, 111, 0, 0, 216, 217, 5,
		103, 0, 0, 217, 218, 5, 105, 0, 0, 218, 219, 5, 110, 0, 0, 219, 44, 1,
		0, 0, 0, 220, 221, 5, 100, 0, 0, 221, 222, 5, 101, 0, 0, 222, 223, 5, 115,
		0, 0, 223, 224, 5, 99, 0, 0, 224, 225, 5, 114, 0, 0, 225, 226, 5, 105,
		0, 0, 226, 227, 5, 112, 0, 0, 227, 228, 5, 116, 0, 0, 228, 229, 5, 105,
		0, 0, 229, 230, 5, 111, 0, 0, 230, 231, 5, 110, 0, 0, 231, 46, 1, 0, 0,
		0, 232, 233, 5, 109, 0, 0, 233, 234, 5, 97, 0, 0, 234, 235, 5, 112, 0,
		0, 235, 48, 1, 0, 0, 0, 236, 237, 5, 91, 0, 0, 237, 50, 1, 0, 0, 0, 238,
		239, 5, 93, 0, 0, 239, 52, 1, 0, 0, 0, 240, 241, 5, 58, 0, 0, 241, 54,
		1, 0, 0, 0, 242, 243, 5, 59, 0, 0, 243, 56, 1, 0, 0, 0, 244, 246, 7, 0,
		0, 0, 245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248, 1, 0, 0, 0,
		247, 249, 3, 87, 43, 0, 248, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 58, 1, 0, 0, 0, 252, 254, 5,
		45, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0,
		0, 255, 256, 5, 48, 0, 0, 256, 257, 5, 120, 0, 0, 257, 259, 1, 0, 0, 0,
		258, 260, 3, 89, 44, 0, 259, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 60, 1, 0, 0, 0, 263, 265, 7,
		0, 0, 0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 285, 1, 0, 0,
		0, 266, 268, 3, 87, 43, 0, 267, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0,
		269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 277, 1, 0, 0, 0, 271,
		273, 5, 46, 0, 0, 272, 274, 3, 87, 43, 0, 273, 272, 1, 0, 0, 0, 274, 275,
		1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 278, 1, 0,
		0, 0, 277, 271, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 286, 1, 0, 0, 0,
		279, 281, 5, 46, 0, 0, 280, 282, 3, 87, 43, 0, 281, 280, 1, 0, 0, 0, 282,
		283, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 286,
		1, 0, 0, 0, 285, 267, 1, 0, 0, 0, 285, 279, 1, 0, 0, 0, 286, 289, 1, 0,
		0, 0, 287, 288, 7, 1, 0, 0, 288, 290, 3, 57, 28, 0, 289, 287, 1, 0, 0,
		0, 289, 290, 1, 0, 0, 0, 290, 62, 1, 0, 0, 0, 291, 292, 5, 98, 0, 0, 292,
		293, 5, 111, 0, 0, 293, 294, 5, 111, 0, 0, 294, 295, 5, 108, 0, 0, 295,
		64, 1, 0, 0, 0, 296, 297, 5, 98, 0, 0, 297, 298, 5, 121, 0, 0, 298, 299,
		5, 116, 0, 0, 299, 300, 5, 101, 0, 0, 300, 66, 1, 0, 0, 0, 301, 302, 5,
		105, 0, 0, 302, 303, 5, 49, 0, 0, 303, 304, 5, 54, 0, 0, 304, 68, 1, 0,
		0, 0, 305, 306, 5, 105, 0, 0, 306, 307, 5, 51, 0, 0, 307, 308, 5, 50, 0,
		0, 308, 70, 1, 0, 0, 0, 309, 310, 5, 105, 0, 0, 310, 311, 5, 54, 0, 0,
		311, 312, 5, 52, 0, 0, 312, 72, 1, 0, 0, 0, 313, 314, 5, 100, 0, 0, 314,
		315, 5, 111, 0, 0, 315, 316, 5, 117, 0, 0, 316, 317, 5, 98, 0, 0, 317,
		318, 5, 108, 0, 0, 318, 319, 5, 101, 0, 0, 319, 74, 1, 0, 0, 0, 320, 321,
		5, 115, 0, 0, 321, 322, 5, 116, 0, 0, 322, 323, 5, 114, 0, 0, 323, 324,
		5, 105, 0, 0, 324, 325, 5, 110, 0, 0, 325, 326, 5, 103, 0, 0, 326, 76,
		1, 0, 0, 0, 327, 332, 5, 34, 0, 0, 328, 331, 3, 79, 39, 0, 329, 331, 8,
		2, 0, 0, 330, 328, 1, 0, 0, 0, 330, 329, 1, 0, 0, 0, 331, 334, 1, 0, 0,
		0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 335, 1, 0, 0, 0, 334,
		332, 1, 0, 0, 0, 335, 346, 5, 34, 0, 0, 336, 341, 5, 39, 0, 0, 337, 340,
		3, 79, 39, 0, 338, 340, 8, 3, 0, 0, 339, 337, 1, 0, 0, 0, 339, 338, 1,
		0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0,
		0, 342, 344, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 346, 5, 39, 0, 0, 345,
		327, 1, 0, 0, 0, 345, 336, 1, 0, 0, 0, 346, 78, 1, 0, 0, 0, 347, 348, 5,
		92, 0, 0, 348, 349, 7, 4, 0, 0, 349, 80, 1, 0, 0, 0, 350, 353, 3, 85, 42,
		0, 351, 353, 5, 95, 0, 0, 352, 350, 1, 0, 0, 0, 352, 351, 1, 0, 0, 0, 353,
		359, 1, 0, 0, 0, 354, 358, 3, 85, 42, 0, 355, 358, 3, 87, 43, 0, 356, 358,
		7, 5, 0, 0, 357, 354, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 356, 1, 0,
		0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0,
		360, 82, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 363, 5, 44, 0, 0, 363,
		84, 1, 0, 0, 0, 364, 365, 7, 6, 0, 0, 365, 86, 1, 0, 0, 0, 366, 367, 2,
		48, 57, 0, 367, 88, 1, 0, 0, 0, 368, 371, 3, 87, 43, 0, 369, 371, 7, 7,
		0, 0, 370, 368, 1, 0, 0, 0, 370, 369, 1, 0, 0, 0, 371, 90, 1, 0, 0, 0,
		372, 377, 7, 8, 0, 0, 373, 374, 5, 13, 0, 0, 374, 377, 5, 10, 0, 0, 375,
		377, 5, 10, 0, 0, 376, 372, 1, 0, 0, 0, 376, 373, 1, 0, 0, 0, 376, 375,
		1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0,
		0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 6, 45, 0, 0, 381, 92, 1, 0, 0, 0,
		382, 383, 5, 47, 0, 0, 383, 386, 5, 47, 0, 0, 384, 386, 5, 35, 0, 0, 385,
		382, 1, 0, 0, 0, 385, 384, 1, 0, 0, 0, 386, 390, 1, 0, 0, 0, 387, 389,
		8, 9, 0, 0, 388, 387, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 390, 388, 1, 0,
		0, 0, 390, 391, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0,
		393, 395, 5, 13, 0, 0, 394, 393, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395,
		396, 1, 0, 0, 0, 396, 397, 5, 10, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399,
		6, 46, 0, 0, 399, 94, 1, 0, 0, 0, 400, 401, 5, 47, 0, 0, 401, 402, 5, 42,
		0, 0, 402, 406, 1, 0, 0, 0, 403, 405, 9, 0, 0, 0, 404, 403, 1, 0, 0, 0,
		405, 408, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 407,
		409, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 410, 5, 42, 0, 0, 410, 411,
		5, 47, 0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 6, 47, 0, 0, 413, 96, 1, 0,
		0, 0, 27, 0, 245, 250, 253, 261, 264, 269, 275, 277, 283, 285, 289, 330,
		332, 339, 341, 345, 352, 357, 359, 370, 376, 378, 385, 390, 394, 406, 1,
		0, 1, 0,
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

// ServiceLexerInit initializes any static state used to implement ServiceLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewServiceLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ServiceLexerInit() {
	staticData := &ServiceLexerLexerStaticData
	staticData.once.Do(servicelexerLexerInit)
}

// NewServiceLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewServiceLexer(input antlr.CharStream) *ServiceLexer {
	ServiceLexerInit()
	l := new(ServiceLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ServiceLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Service.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ServiceLexer tokens.
const (
	ServiceLexerT__0        = 1
	ServiceLexerT__1        = 2
	ServiceLexerT__2        = 3
	ServiceLexerT__3        = 4
	ServiceLexerT__4        = 5
	ServiceLexerT__5        = 6
	ServiceLexerT__6        = 7
	ServiceLexerT__7        = 8
	ServiceLexerT__8        = 9
	ServiceLexerT__9        = 10
	ServiceLexerT__10       = 11
	ServiceLexerT__11       = 12
	ServiceLexerT__12       = 13
	ServiceLexerT__13       = 14
	ServiceLexerT__14       = 15
	ServiceLexerT__15       = 16
	ServiceLexerT__16       = 17
	ServiceLexerT__17       = 18
	ServiceLexerT__18       = 19
	ServiceLexerT__19       = 20
	ServiceLexerT__20       = 21
	ServiceLexerT__21       = 22
	ServiceLexerT__22       = 23
	ServiceLexerT__23       = 24
	ServiceLexerT__24       = 25
	ServiceLexerT__25       = 26
	ServiceLexerT__26       = 27
	ServiceLexerT__27       = 28
	ServiceLexerINTEGER     = 29
	ServiceLexerHEX_INTEGER = 30
	ServiceLexerDOUBLE      = 31
	ServiceLexerTYPE_BOOL   = 32
	ServiceLexerTYPE_BYTE   = 33
	ServiceLexerTYPE_I16    = 34
	ServiceLexerTYPE_I32    = 35
	ServiceLexerTYPE_I64    = 36
	ServiceLexerTYPE_DOUBLE = 37
	ServiceLexerTYPE_STRING = 38
	ServiceLexerLITERAL     = 39
	ServiceLexerIDENTIFIER  = 40
	ServiceLexerCOMMA       = 41
	ServiceLexerWS          = 42
	ServiceLexerSL_COMMENT  = 43
	ServiceLexerML_COMMENT  = 44
)

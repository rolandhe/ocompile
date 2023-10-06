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
		"", "'with_client'", "'='", "'namespace'", "'true'", "'false'", "'struct'",
		"'{'", "'}'", "'service'", "'required'", "'optional'", "'POST'", "'('",
		"')'", "'GET'", "'url'", "'list'", "'<'", "'>'", "'void'", "'not_login'",
		"'map'", "'['", "']'", "':'", "';'", "", "", "", "'bool'", "'byte'",
		"'i16'", "'i32'", "'i64'", "'double'", "'string'", "'binary'", "", "",
		"','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "INTEGER", "HEX_INTEGER", "DOUBLE",
		"TYPE_BOOL", "TYPE_BYTE", "TYPE_I16", "TYPE_I32", "TYPE_I64", "TYPE_DOUBLE",
		"TYPE_STRING", "TYPE_BINARY", "LITERAL", "IDENTIFIER", "COMMA", "WS",
		"SL_COMMENT", "ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "INTEGER", "HEX_INTEGER", "DOUBLE", "TYPE_BOOL", "TYPE_BYTE",
		"TYPE_I16", "TYPE_I32", "TYPE_I64", "TYPE_DOUBLE", "TYPE_STRING", "TYPE_BINARY",
		"LITERAL", "ESC_SEQ", "IDENTIFIER", "COMMA", "LETTER", "DIGIT", "HEX_DIGIT",
		"WS", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 397, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 3, 26, 222, 8, 26, 1, 26, 4, 26, 225, 8, 26, 11, 26,
		12, 26, 226, 1, 27, 3, 27, 230, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 4, 27,
		236, 8, 27, 11, 27, 12, 27, 237, 1, 28, 3, 28, 241, 8, 28, 1, 28, 4, 28,
		244, 8, 28, 11, 28, 12, 28, 245, 1, 28, 1, 28, 4, 28, 250, 8, 28, 11, 28,
		12, 28, 251, 3, 28, 254, 8, 28, 1, 28, 1, 28, 4, 28, 258, 8, 28, 11, 28,
		12, 28, 259, 3, 28, 262, 8, 28, 1, 28, 1, 28, 3, 28, 266, 8, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 5, 37, 314, 8, 37, 10, 37, 12, 37, 317, 9, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 5, 37, 323, 8, 37, 10, 37, 12, 37, 326, 9, 37,
		1, 37, 3, 37, 329, 8, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 3, 39, 336,
		8, 39, 1, 39, 1, 39, 1, 39, 5, 39, 341, 8, 39, 10, 39, 12, 39, 344, 9,
		39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 3, 43, 354,
		8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 4, 44, 360, 8, 44, 11, 44, 12, 44, 361,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 3, 45, 369, 8, 45, 1, 45, 5, 45, 372,
		8, 45, 10, 45, 12, 45, 375, 9, 45, 1, 45, 3, 45, 378, 8, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 388, 8, 46, 10, 46,
		12, 46, 391, 9, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 389, 0, 47, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 0, 79, 39, 81, 40, 83, 0, 85, 0, 87, 0, 89, 41, 91, 42, 93, 43, 1,
		0, 10, 2, 0, 43, 43, 45, 45, 2, 0, 69, 69, 101, 101, 2, 0, 34, 34, 92,
		92, 2, 0, 39, 39, 92, 92, 6, 0, 34, 34, 39, 39, 92, 92, 110, 110, 114,
		114, 116, 116, 2, 0, 46, 46, 95, 95, 2, 0, 65, 90, 97, 122, 2, 0, 65, 70,
		97, 102, 2, 0, 9, 9, 32, 32, 1, 0, 10, 10, 420, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 1, 95, 1, 0,
		0, 0, 3, 107, 1, 0, 0, 0, 5, 109, 1, 0, 0, 0, 7, 119, 1, 0, 0, 0, 9, 124,
		1, 0, 0, 0, 11, 130, 1, 0, 0, 0, 13, 137, 1, 0, 0, 0, 15, 139, 1, 0, 0,
		0, 17, 141, 1, 0, 0, 0, 19, 149, 1, 0, 0, 0, 21, 158, 1, 0, 0, 0, 23, 167,
		1, 0, 0, 0, 25, 172, 1, 0, 0, 0, 27, 174, 1, 0, 0, 0, 29, 176, 1, 0, 0,
		0, 31, 180, 1, 0, 0, 0, 33, 184, 1, 0, 0, 0, 35, 189, 1, 0, 0, 0, 37, 191,
		1, 0, 0, 0, 39, 193, 1, 0, 0, 0, 41, 198, 1, 0, 0, 0, 43, 208, 1, 0, 0,
		0, 45, 212, 1, 0, 0, 0, 47, 214, 1, 0, 0, 0, 49, 216, 1, 0, 0, 0, 51, 218,
		1, 0, 0, 0, 53, 221, 1, 0, 0, 0, 55, 229, 1, 0, 0, 0, 57, 240, 1, 0, 0,
		0, 59, 267, 1, 0, 0, 0, 61, 272, 1, 0, 0, 0, 63, 277, 1, 0, 0, 0, 65, 281,
		1, 0, 0, 0, 67, 285, 1, 0, 0, 0, 69, 289, 1, 0, 0, 0, 71, 296, 1, 0, 0,
		0, 73, 303, 1, 0, 0, 0, 75, 328, 1, 0, 0, 0, 77, 330, 1, 0, 0, 0, 79, 335,
		1, 0, 0, 0, 81, 345, 1, 0, 0, 0, 83, 347, 1, 0, 0, 0, 85, 349, 1, 0, 0,
		0, 87, 353, 1, 0, 0, 0, 89, 359, 1, 0, 0, 0, 91, 368, 1, 0, 0, 0, 93, 383,
		1, 0, 0, 0, 95, 96, 5, 119, 0, 0, 96, 97, 5, 105, 0, 0, 97, 98, 5, 116,
		0, 0, 98, 99, 5, 104, 0, 0, 99, 100, 5, 95, 0, 0, 100, 101, 5, 99, 0, 0,
		101, 102, 5, 108, 0, 0, 102, 103, 5, 105, 0, 0, 103, 104, 5, 101, 0, 0,
		104, 105, 5, 110, 0, 0, 105, 106, 5, 116, 0, 0, 106, 2, 1, 0, 0, 0, 107,
		108, 5, 61, 0, 0, 108, 4, 1, 0, 0, 0, 109, 110, 5, 110, 0, 0, 110, 111,
		5, 97, 0, 0, 111, 112, 5, 109, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114,
		5, 115, 0, 0, 114, 115, 5, 112, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117,
		5, 99, 0, 0, 117, 118, 5, 101, 0, 0, 118, 6, 1, 0, 0, 0, 119, 120, 5, 116,
		0, 0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 117, 0, 0, 122, 123, 5, 101,
		0, 0, 123, 8, 1, 0, 0, 0, 124, 125, 5, 102, 0, 0, 125, 126, 5, 97, 0, 0,
		126, 127, 5, 108, 0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5, 101, 0, 0,
		129, 10, 1, 0, 0, 0, 130, 131, 5, 115, 0, 0, 131, 132, 5, 116, 0, 0, 132,
		133, 5, 114, 0, 0, 133, 134, 5, 117, 0, 0, 134, 135, 5, 99, 0, 0, 135,
		136, 5, 116, 0, 0, 136, 12, 1, 0, 0, 0, 137, 138, 5, 123, 0, 0, 138, 14,
		1, 0, 0, 0, 139, 140, 5, 125, 0, 0, 140, 16, 1, 0, 0, 0, 141, 142, 5, 115,
		0, 0, 142, 143, 5, 101, 0, 0, 143, 144, 5, 114, 0, 0, 144, 145, 5, 118,
		0, 0, 145, 146, 5, 105, 0, 0, 146, 147, 5, 99, 0, 0, 147, 148, 5, 101,
		0, 0, 148, 18, 1, 0, 0, 0, 149, 150, 5, 114, 0, 0, 150, 151, 5, 101, 0,
		0, 151, 152, 5, 113, 0, 0, 152, 153, 5, 117, 0, 0, 153, 154, 5, 105, 0,
		0, 154, 155, 5, 114, 0, 0, 155, 156, 5, 101, 0, 0, 156, 157, 5, 100, 0,
		0, 157, 20, 1, 0, 0, 0, 158, 159, 5, 111, 0, 0, 159, 160, 5, 112, 0, 0,
		160, 161, 5, 116, 0, 0, 161, 162, 5, 105, 0, 0, 162, 163, 5, 111, 0, 0,
		163, 164, 5, 110, 0, 0, 164, 165, 5, 97, 0, 0, 165, 166, 5, 108, 0, 0,
		166, 22, 1, 0, 0, 0, 167, 168, 5, 80, 0, 0, 168, 169, 5, 79, 0, 0, 169,
		170, 5, 83, 0, 0, 170, 171, 5, 84, 0, 0, 171, 24, 1, 0, 0, 0, 172, 173,
		5, 40, 0, 0, 173, 26, 1, 0, 0, 0, 174, 175, 5, 41, 0, 0, 175, 28, 1, 0,
		0, 0, 176, 177, 5, 71, 0, 0, 177, 178, 5, 69, 0, 0, 178, 179, 5, 84, 0,
		0, 179, 30, 1, 0, 0, 0, 180, 181, 5, 117, 0, 0, 181, 182, 5, 114, 0, 0,
		182, 183, 5, 108, 0, 0, 183, 32, 1, 0, 0, 0, 184, 185, 5, 108, 0, 0, 185,
		186, 5, 105, 0, 0, 186, 187, 5, 115, 0, 0, 187, 188, 5, 116, 0, 0, 188,
		34, 1, 0, 0, 0, 189, 190, 5, 60, 0, 0, 190, 36, 1, 0, 0, 0, 191, 192, 5,
		62, 0, 0, 192, 38, 1, 0, 0, 0, 193, 194, 5, 118, 0, 0, 194, 195, 5, 111,
		0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5, 100, 0, 0, 197, 40, 1, 0, 0,
		0, 198, 199, 5, 110, 0, 0, 199, 200, 5, 111, 0, 0, 200, 201, 5, 116, 0,
		0, 201, 202, 5, 95, 0, 0, 202, 203, 5, 108, 0, 0, 203, 204, 5, 111, 0,
		0, 204, 205, 5, 103, 0, 0, 205, 206, 5, 105, 0, 0, 206, 207, 5, 110, 0,
		0, 207, 42, 1, 0, 0, 0, 208, 209, 5, 109, 0, 0, 209, 210, 5, 97, 0, 0,
		210, 211, 5, 112, 0, 0, 211, 44, 1, 0, 0, 0, 212, 213, 5, 91, 0, 0, 213,
		46, 1, 0, 0, 0, 214, 215, 5, 93, 0, 0, 215, 48, 1, 0, 0, 0, 216, 217, 5,
		58, 0, 0, 217, 50, 1, 0, 0, 0, 218, 219, 5, 59, 0, 0, 219, 52, 1, 0, 0,
		0, 220, 222, 7, 0, 0, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		224, 1, 0, 0, 0, 223, 225, 3, 85, 42, 0, 224, 223, 1, 0, 0, 0, 225, 226,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 54, 1, 0,
		0, 0, 228, 230, 5, 45, 0, 0, 229, 228, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 5, 48, 0, 0, 232, 233, 5, 120, 0, 0, 233,
		235, 1, 0, 0, 0, 234, 236, 3, 87, 43, 0, 235, 234, 1, 0, 0, 0, 236, 237,
		1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 56, 1, 0,
		0, 0, 239, 241, 7, 0, 0, 0, 240, 239, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0,
		241, 261, 1, 0, 0, 0, 242, 244, 3, 85, 42, 0, 243, 242, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 253,
		1, 0, 0, 0, 247, 249, 5, 46, 0, 0, 248, 250, 3, 85, 42, 0, 249, 248, 1,
		0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0,
		0, 252, 254, 1, 0, 0, 0, 253, 247, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		262, 1, 0, 0, 0, 255, 257, 5, 46, 0, 0, 256, 258, 3, 85, 42, 0, 257, 256,
		1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0,
		0, 0, 260, 262, 1, 0, 0, 0, 261, 243, 1, 0, 0, 0, 261, 255, 1, 0, 0, 0,
		262, 265, 1, 0, 0, 0, 263, 264, 7, 1, 0, 0, 264, 266, 3, 53, 26, 0, 265,
		263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 58, 1, 0, 0, 0, 267, 268, 5,
		98, 0, 0, 268, 269, 5, 111, 0, 0, 269, 270, 5, 111, 0, 0, 270, 271, 5,
		108, 0, 0, 271, 60, 1, 0, 0, 0, 272, 273, 5, 98, 0, 0, 273, 274, 5, 121,
		0, 0, 274, 275, 5, 116, 0, 0, 275, 276, 5, 101, 0, 0, 276, 62, 1, 0, 0,
		0, 277, 278, 5, 105, 0, 0, 278, 279, 5, 49, 0, 0, 279, 280, 5, 54, 0, 0,
		280, 64, 1, 0, 0, 0, 281, 282, 5, 105, 0, 0, 282, 283, 5, 51, 0, 0, 283,
		284, 5, 50, 0, 0, 284, 66, 1, 0, 0, 0, 285, 286, 5, 105, 0, 0, 286, 287,
		5, 54, 0, 0, 287, 288, 5, 52, 0, 0, 288, 68, 1, 0, 0, 0, 289, 290, 5, 100,
		0, 0, 290, 291, 5, 111, 0, 0, 291, 292, 5, 117, 0, 0, 292, 293, 5, 98,
		0, 0, 293, 294, 5, 108, 0, 0, 294, 295, 5, 101, 0, 0, 295, 70, 1, 0, 0,
		0, 296, 297, 5, 115, 0, 0, 297, 298, 5, 116, 0, 0, 298, 299, 5, 114, 0,
		0, 299, 300, 5, 105, 0, 0, 300, 301, 5, 110, 0, 0, 301, 302, 5, 103, 0,
		0, 302, 72, 1, 0, 0, 0, 303, 304, 5, 98, 0, 0, 304, 305, 5, 105, 0, 0,
		305, 306, 5, 110, 0, 0, 306, 307, 5, 97, 0, 0, 307, 308, 5, 114, 0, 0,
		308, 309, 5, 121, 0, 0, 309, 74, 1, 0, 0, 0, 310, 315, 5, 34, 0, 0, 311,
		314, 3, 77, 38, 0, 312, 314, 8, 2, 0, 0, 313, 311, 1, 0, 0, 0, 313, 312,
		1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0,
		0, 0, 316, 318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 329, 5, 34, 0, 0,
		319, 324, 5, 39, 0, 0, 320, 323, 3, 77, 38, 0, 321, 323, 8, 3, 0, 0, 322,
		320, 1, 0, 0, 0, 322, 321, 1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322,
		1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 327, 1, 0, 0, 0, 326, 324, 1, 0,
		0, 0, 327, 329, 5, 39, 0, 0, 328, 310, 1, 0, 0, 0, 328, 319, 1, 0, 0, 0,
		329, 76, 1, 0, 0, 0, 330, 331, 5, 92, 0, 0, 331, 332, 7, 4, 0, 0, 332,
		78, 1, 0, 0, 0, 333, 336, 3, 83, 41, 0, 334, 336, 5, 95, 0, 0, 335, 333,
		1, 0, 0, 0, 335, 334, 1, 0, 0, 0, 336, 342, 1, 0, 0, 0, 337, 341, 3, 83,
		41, 0, 338, 341, 3, 85, 42, 0, 339, 341, 7, 5, 0, 0, 340, 337, 1, 0, 0,
		0, 340, 338, 1, 0, 0, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342,
		340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 80, 1, 0, 0, 0, 344, 342, 1,
		0, 0, 0, 345, 346, 5, 44, 0, 0, 346, 82, 1, 0, 0, 0, 347, 348, 7, 6, 0,
		0, 348, 84, 1, 0, 0, 0, 349, 350, 2, 48, 57, 0, 350, 86, 1, 0, 0, 0, 351,
		354, 3, 85, 42, 0, 352, 354, 7, 7, 0, 0, 353, 351, 1, 0, 0, 0, 353, 352,
		1, 0, 0, 0, 354, 88, 1, 0, 0, 0, 355, 360, 7, 8, 0, 0, 356, 357, 5, 13,
		0, 0, 357, 360, 5, 10, 0, 0, 358, 360, 5, 10, 0, 0, 359, 355, 1, 0, 0,
		0, 359, 356, 1, 0, 0, 0, 359, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361,
		359, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 364,
		6, 44, 0, 0, 364, 90, 1, 0, 0, 0, 365, 366, 5, 47, 0, 0, 366, 369, 5, 47,
		0, 0, 367, 369, 5, 35, 0, 0, 368, 365, 1, 0, 0, 0, 368, 367, 1, 0, 0, 0,
		369, 373, 1, 0, 0, 0, 370, 372, 8, 9, 0, 0, 371, 370, 1, 0, 0, 0, 372,
		375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 377,
		1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 378, 5, 13, 0, 0, 377, 376, 1, 0,
		0, 0, 377, 378, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 380, 5, 10, 0, 0,
		380, 381, 1, 0, 0, 0, 381, 382, 6, 45, 0, 0, 382, 92, 1, 0, 0, 0, 383,
		384, 5, 47, 0, 0, 384, 385, 5, 42, 0, 0, 385, 389, 1, 0, 0, 0, 386, 388,
		9, 0, 0, 0, 387, 386, 1, 0, 0, 0, 388, 391, 1, 0, 0, 0, 389, 390, 1, 0,
		0, 0, 389, 387, 1, 0, 0, 0, 390, 392, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0,
		392, 393, 5, 42, 0, 0, 393, 394, 5, 47, 0, 0, 394, 395, 1, 0, 0, 0, 395,
		396, 6, 46, 0, 0, 396, 94, 1, 0, 0, 0, 27, 0, 221, 226, 229, 237, 240,
		245, 251, 253, 259, 261, 265, 313, 315, 322, 324, 328, 335, 340, 342, 353,
		359, 361, 368, 373, 377, 389, 1, 0, 1, 0,
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
	ServiceLexerINTEGER     = 27
	ServiceLexerHEX_INTEGER = 28
	ServiceLexerDOUBLE      = 29
	ServiceLexerTYPE_BOOL   = 30
	ServiceLexerTYPE_BYTE   = 31
	ServiceLexerTYPE_I16    = 32
	ServiceLexerTYPE_I32    = 33
	ServiceLexerTYPE_I64    = 34
	ServiceLexerTYPE_DOUBLE = 35
	ServiceLexerTYPE_STRING = 36
	ServiceLexerTYPE_BINARY = 37
	ServiceLexerLITERAL     = 38
	ServiceLexerIDENTIFIER  = 39
	ServiceLexerCOMMA       = 40
	ServiceLexerWS          = 41
	ServiceLexerSL_COMMENT  = 42
	ServiceLexerML_COMMENT  = 43
)
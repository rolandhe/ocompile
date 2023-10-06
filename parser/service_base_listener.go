// Code generated from Service.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Service

import "github.com/antlr4-go/antlr/v4"

// BaseServiceListener is a complete listener for a parse tree produced by ServiceParser.
type BaseServiceListener struct{}

var _ ServiceListener = &BaseServiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseServiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseServiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseServiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseServiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseServiceListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseServiceListener) ExitDocument(ctx *DocumentContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseServiceListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseServiceListener) ExitHeader(ctx *HeaderContext) {}

// EnterNamespace_ is called when production namespace_ is entered.
func (s *BaseServiceListener) EnterNamespace_(ctx *Namespace_Context) {}

// ExitNamespace_ is called when production namespace_ is exited.
func (s *BaseServiceListener) ExitNamespace_(ctx *Namespace_Context) {}

// EnterWith_client_optional is called when production with_client_optional is entered.
func (s *BaseServiceListener) EnterWith_client_optional(ctx *With_client_optionalContext) {}

// ExitWith_client_optional is called when production with_client_optional is exited.
func (s *BaseServiceListener) ExitWith_client_optional(ctx *With_client_optionalContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseServiceListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseServiceListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterStruct_ is called when production struct_ is entered.
func (s *BaseServiceListener) EnterStruct_(ctx *Struct_Context) {}

// ExitStruct_ is called when production struct_ is exited.
func (s *BaseServiceListener) ExitStruct_(ctx *Struct_Context) {}

// EnterService is called when production service is entered.
func (s *BaseServiceListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseServiceListener) ExitService(ctx *ServiceContext) {}

// EnterField is called when production field is entered.
func (s *BaseServiceListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseServiceListener) ExitField(ctx *FieldContext) {}

// EnterField_req is called when production field_req is entered.
func (s *BaseServiceListener) EnterField_req(ctx *Field_reqContext) {}

// ExitField_req is called when production field_req is exited.
func (s *BaseServiceListener) ExitField_req(ctx *Field_reqContext) {}

// EnterMethod_ is called when production method_ is entered.
func (s *BaseServiceListener) EnterMethod_(ctx *Method_Context) {}

// ExitMethod_ is called when production method_ is exited.
func (s *BaseServiceListener) ExitMethod_(ctx *Method_Context) {}

// EnterPost_ is called when production post_ is entered.
func (s *BaseServiceListener) EnterPost_(ctx *Post_Context) {}

// ExitPost_ is called when production post_ is exited.
func (s *BaseServiceListener) ExitPost_(ctx *Post_Context) {}

// EnterGet_ is called when production get_ is entered.
func (s *BaseServiceListener) EnterGet_(ctx *Get_Context) {}

// ExitGet_ is called when production get_ is exited.
func (s *BaseServiceListener) ExitGet_(ctx *Get_Context) {}

// EnterUrl_ is called when production url_ is entered.
func (s *BaseServiceListener) EnterUrl_(ctx *Url_Context) {}

// ExitUrl_ is called when production url_ is exited.
func (s *BaseServiceListener) ExitUrl_(ctx *Url_Context) {}

// EnterMethod_type is called when production method_type is entered.
func (s *BaseServiceListener) EnterMethod_type(ctx *Method_typeContext) {}

// ExitMethod_type is called when production method_type is exited.
func (s *BaseServiceListener) ExitMethod_type(ctx *Method_typeContext) {}

// EnterStruct_type_list is called when production struct_type_list is entered.
func (s *BaseServiceListener) EnterStruct_type_list(ctx *Struct_type_listContext) {}

// ExitStruct_type_list is called when production struct_type_list is exited.
func (s *BaseServiceListener) ExitStruct_type_list(ctx *Struct_type_listContext) {}

// EnterGet_param_ is called when production get_param_ is entered.
func (s *BaseServiceListener) EnterGet_param_(ctx *Get_param_Context) {}

// ExitGet_param_ is called when production get_param_ is exited.
func (s *BaseServiceListener) ExitGet_param_(ctx *Get_param_Context) {}

// EnterReal_base_type_list_ is called when production real_base_type_list_ is entered.
func (s *BaseServiceListener) EnterReal_base_type_list_(ctx *Real_base_type_list_Context) {}

// ExitReal_base_type_list_ is called when production real_base_type_list_ is exited.
func (s *BaseServiceListener) ExitReal_base_type_list_(ctx *Real_base_type_list_Context) {}

// EnterVoid_ is called when production void_ is entered.
func (s *BaseServiceListener) EnterVoid_(ctx *Void_Context) {}

// ExitVoid_ is called when production void_ is exited.
func (s *BaseServiceListener) ExitVoid_(ctx *Void_Context) {}

// EnterMethod_param_ is called when production method_param_ is entered.
func (s *BaseServiceListener) EnterMethod_param_(ctx *Method_param_Context) {}

// ExitMethod_param_ is called when production method_param_ is exited.
func (s *BaseServiceListener) ExitMethod_param_(ctx *Method_param_Context) {}

// EnterSingle_struct_param is called when production single_struct_param is entered.
func (s *BaseServiceListener) EnterSingle_struct_param(ctx *Single_struct_paramContext) {}

// ExitSingle_struct_param is called when production single_struct_param is exited.
func (s *BaseServiceListener) ExitSingle_struct_param(ctx *Single_struct_paramContext) {}

// EnterSimple_param_ is called when production simple_param_ is entered.
func (s *BaseServiceListener) EnterSimple_param_(ctx *Simple_param_Context) {}

// ExitSimple_param_ is called when production simple_param_ is exited.
func (s *BaseServiceListener) ExitSimple_param_(ctx *Simple_param_Context) {}

// EnterReal_base_type_parm is called when production real_base_type_parm is entered.
func (s *BaseServiceListener) EnterReal_base_type_parm(ctx *Real_base_type_parmContext) {}

// ExitReal_base_type_parm is called when production real_base_type_parm is exited.
func (s *BaseServiceListener) ExitReal_base_type_parm(ctx *Real_base_type_parmContext) {}

// EnterReal_base_type_list_parm is called when production real_base_type_list_parm is entered.
func (s *BaseServiceListener) EnterReal_base_type_list_parm(ctx *Real_base_type_list_parmContext) {}

// ExitReal_base_type_list_parm is called when production real_base_type_list_parm is exited.
func (s *BaseServiceListener) ExitReal_base_type_list_parm(ctx *Real_base_type_list_parmContext) {}

// EnterNot_login is called when production not_login is entered.
func (s *BaseServiceListener) EnterNot_login(ctx *Not_loginContext) {}

// ExitNot_login is called when production not_login is exited.
func (s *BaseServiceListener) ExitNot_login(ctx *Not_loginContext) {}

// EnterType_annotations is called when production type_annotations is entered.
func (s *BaseServiceListener) EnterType_annotations(ctx *Type_annotationsContext) {}

// ExitType_annotations is called when production type_annotations is exited.
func (s *BaseServiceListener) ExitType_annotations(ctx *Type_annotationsContext) {}

// EnterType_annotation is called when production type_annotation is entered.
func (s *BaseServiceListener) EnterType_annotation(ctx *Type_annotationContext) {}

// ExitType_annotation is called when production type_annotation is exited.
func (s *BaseServiceListener) ExitType_annotation(ctx *Type_annotationContext) {}

// EnterField_annotations is called when production field_annotations is entered.
func (s *BaseServiceListener) EnterField_annotations(ctx *Field_annotationsContext) {}

// ExitField_annotations is called when production field_annotations is exited.
func (s *BaseServiceListener) ExitField_annotations(ctx *Field_annotationsContext) {}

// EnterField_annotation is called when production field_annotation is entered.
func (s *BaseServiceListener) EnterField_annotation(ctx *Field_annotationContext) {}

// ExitField_annotation is called when production field_annotation is exited.
func (s *BaseServiceListener) ExitField_annotation(ctx *Field_annotationContext) {}

// EnterAnnotation_value is called when production annotation_value is entered.
func (s *BaseServiceListener) EnterAnnotation_value(ctx *Annotation_valueContext) {}

// ExitAnnotation_value is called when production annotation_value is exited.
func (s *BaseServiceListener) ExitAnnotation_value(ctx *Annotation_valueContext) {}

// EnterField_type is called when production field_type is entered.
func (s *BaseServiceListener) EnterField_type(ctx *Field_typeContext) {}

// ExitField_type is called when production field_type is exited.
func (s *BaseServiceListener) ExitField_type(ctx *Field_typeContext) {}

// EnterBase_type is called when production base_type is entered.
func (s *BaseServiceListener) EnterBase_type(ctx *Base_typeContext) {}

// ExitBase_type is called when production base_type is exited.
func (s *BaseServiceListener) ExitBase_type(ctx *Base_typeContext) {}

// EnterContainer_type is called when production container_type is entered.
func (s *BaseServiceListener) EnterContainer_type(ctx *Container_typeContext) {}

// ExitContainer_type is called when production container_type is exited.
func (s *BaseServiceListener) ExitContainer_type(ctx *Container_typeContext) {}

// EnterStruct_type is called when production struct_type is entered.
func (s *BaseServiceListener) EnterStruct_type(ctx *Struct_typeContext) {}

// ExitStruct_type is called when production struct_type is exited.
func (s *BaseServiceListener) ExitStruct_type(ctx *Struct_typeContext) {}

// EnterMap_type is called when production map_type is entered.
func (s *BaseServiceListener) EnterMap_type(ctx *Map_typeContext) {}

// ExitMap_type is called when production map_type is exited.
func (s *BaseServiceListener) ExitMap_type(ctx *Map_typeContext) {}

// EnterList_type is called when production list_type is entered.
func (s *BaseServiceListener) EnterList_type(ctx *List_typeContext) {}

// ExitList_type is called when production list_type is exited.
func (s *BaseServiceListener) ExitList_type(ctx *List_typeContext) {}

// EnterConst_value is called when production const_value is entered.
func (s *BaseServiceListener) EnterConst_value(ctx *Const_valueContext) {}

// ExitConst_value is called when production const_value is exited.
func (s *BaseServiceListener) ExitConst_value(ctx *Const_valueContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseServiceListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseServiceListener) ExitInteger(ctx *IntegerContext) {}

// EnterConst_list is called when production const_list is entered.
func (s *BaseServiceListener) EnterConst_list(ctx *Const_listContext) {}

// ExitConst_list is called when production const_list is exited.
func (s *BaseServiceListener) ExitConst_list(ctx *Const_listContext) {}

// EnterConst_map_entry is called when production const_map_entry is entered.
func (s *BaseServiceListener) EnterConst_map_entry(ctx *Const_map_entryContext) {}

// ExitConst_map_entry is called when production const_map_entry is exited.
func (s *BaseServiceListener) ExitConst_map_entry(ctx *Const_map_entryContext) {}

// EnterConst_map is called when production const_map is entered.
func (s *BaseServiceListener) EnterConst_map(ctx *Const_mapContext) {}

// ExitConst_map is called when production const_map is exited.
func (s *BaseServiceListener) ExitConst_map(ctx *Const_mapContext) {}

// EnterList_separator is called when production list_separator is entered.
func (s *BaseServiceListener) EnterList_separator(ctx *List_separatorContext) {}

// ExitList_separator is called when production list_separator is exited.
func (s *BaseServiceListener) ExitList_separator(ctx *List_separatorContext) {}

// EnterReal_base_type is called when production real_base_type is entered.
func (s *BaseServiceListener) EnterReal_base_type(ctx *Real_base_typeContext) {}

// ExitReal_base_type is called when production real_base_type is exited.
func (s *BaseServiceListener) ExitReal_base_type(ctx *Real_base_typeContext) {}

// EnterMap_key_type is called when production map_key_type is entered.
func (s *BaseServiceListener) EnterMap_key_type(ctx *Map_key_typeContext) {}

// ExitMap_key_type is called when production map_key_type is exited.
func (s *BaseServiceListener) ExitMap_key_type(ctx *Map_key_typeContext) {}

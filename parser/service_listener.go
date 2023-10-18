// Code generated from Service.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Service

import "github.com/antlr4-go/antlr/v4"

// ServiceListener is a complete listener for a parse tree produced by ServiceParser.
type ServiceListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterNamespace_ is called when entering the namespace_ production.
	EnterNamespace_(c *Namespace_Context)

	// EnterGolang_import_ is called when entering the golang_import_ production.
	EnterGolang_import_(c *Golang_import_Context)

	// EnterWith_client_optional is called when entering the with_client_optional production.
	EnterWith_client_optional(c *With_client_optionalContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterStruct_ is called when entering the struct_ production.
	EnterStruct_(c *Struct_Context)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterField_req is called when entering the field_req production.
	EnterField_req(c *Field_reqContext)

	// EnterMethod_ is called when entering the method_ production.
	EnterMethod_(c *Method_Context)

	// EnterPost_ is called when entering the post_ production.
	EnterPost_(c *Post_Context)

	// EnterGet_ is called when entering the get_ production.
	EnterGet_(c *Get_Context)

	// EnterUrl_ is called when entering the url_ production.
	EnterUrl_(c *Url_Context)

	// EnterMethod_type is called when entering the method_type production.
	EnterMethod_type(c *Method_typeContext)

	// EnterStruct_type_list is called when entering the struct_type_list production.
	EnterStruct_type_list(c *Struct_type_listContext)

	// EnterGet_param_ is called when entering the get_param_ production.
	EnterGet_param_(c *Get_param_Context)

	// EnterNext_simple_param_ is called when entering the next_simple_param_ production.
	EnterNext_simple_param_(c *Next_simple_param_Context)

	// EnterReal_base_type_list_ is called when entering the real_base_type_list_ production.
	EnterReal_base_type_list_(c *Real_base_type_list_Context)

	// EnterVoid_ is called when entering the void_ production.
	EnterVoid_(c *Void_Context)

	// EnterMethod_param_ is called when entering the method_param_ production.
	EnterMethod_param_(c *Method_param_Context)

	// EnterSingle_struct_param is called when entering the single_struct_param production.
	EnterSingle_struct_param(c *Single_struct_paramContext)

	// EnterSimple_param_ is called when entering the simple_param_ production.
	EnterSimple_param_(c *Simple_param_Context)

	// EnterReal_base_type_parm is called when entering the real_base_type_parm production.
	EnterReal_base_type_parm(c *Real_base_type_parmContext)

	// EnterReal_base_type_list_parm is called when entering the real_base_type_list_parm production.
	EnterReal_base_type_list_parm(c *Real_base_type_list_parmContext)

	// EnterNot_login is called when entering the not_login production.
	EnterNot_login(c *Not_loginContext)

	// EnterMethod_description is called when entering the method_description production.
	EnterMethod_description(c *Method_descriptionContext)

	// EnterMethod_description_content is called when entering the method_description_content production.
	EnterMethod_description_content(c *Method_description_contentContext)

	// EnterType_annotations is called when entering the type_annotations production.
	EnterType_annotations(c *Type_annotationsContext)

	// EnterType_annotation is called when entering the type_annotation production.
	EnterType_annotation(c *Type_annotationContext)

	// EnterField_annotations is called when entering the field_annotations production.
	EnterField_annotations(c *Field_annotationsContext)

	// EnterField_annotation is called when entering the field_annotation production.
	EnterField_annotation(c *Field_annotationContext)

	// EnterAnnotation_value is called when entering the annotation_value production.
	EnterAnnotation_value(c *Annotation_valueContext)

	// EnterField_type is called when entering the field_type production.
	EnterField_type(c *Field_typeContext)

	// EnterBase_type is called when entering the base_type production.
	EnterBase_type(c *Base_typeContext)

	// EnterContainer_type is called when entering the container_type production.
	EnterContainer_type(c *Container_typeContext)

	// EnterStruct_type is called when entering the struct_type production.
	EnterStruct_type(c *Struct_typeContext)

	// EnterMap_type is called when entering the map_type production.
	EnterMap_type(c *Map_typeContext)

	// EnterList_type is called when entering the list_type production.
	EnterList_type(c *List_typeContext)

	// EnterConst_value is called when entering the const_value production.
	EnterConst_value(c *Const_valueContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterConst_list is called when entering the const_list production.
	EnterConst_list(c *Const_listContext)

	// EnterConst_map_entry is called when entering the const_map_entry production.
	EnterConst_map_entry(c *Const_map_entryContext)

	// EnterConst_map is called when entering the const_map production.
	EnterConst_map(c *Const_mapContext)

	// EnterList_separator is called when entering the list_separator production.
	EnterList_separator(c *List_separatorContext)

	// EnterReal_base_type is called when entering the real_base_type production.
	EnterReal_base_type(c *Real_base_typeContext)

	// EnterMap_key_type is called when entering the map_key_type production.
	EnterMap_key_type(c *Map_key_typeContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitNamespace_ is called when exiting the namespace_ production.
	ExitNamespace_(c *Namespace_Context)

	// ExitGolang_import_ is called when exiting the golang_import_ production.
	ExitGolang_import_(c *Golang_import_Context)

	// ExitWith_client_optional is called when exiting the with_client_optional production.
	ExitWith_client_optional(c *With_client_optionalContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitStruct_ is called when exiting the struct_ production.
	ExitStruct_(c *Struct_Context)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitField_req is called when exiting the field_req production.
	ExitField_req(c *Field_reqContext)

	// ExitMethod_ is called when exiting the method_ production.
	ExitMethod_(c *Method_Context)

	// ExitPost_ is called when exiting the post_ production.
	ExitPost_(c *Post_Context)

	// ExitGet_ is called when exiting the get_ production.
	ExitGet_(c *Get_Context)

	// ExitUrl_ is called when exiting the url_ production.
	ExitUrl_(c *Url_Context)

	// ExitMethod_type is called when exiting the method_type production.
	ExitMethod_type(c *Method_typeContext)

	// ExitStruct_type_list is called when exiting the struct_type_list production.
	ExitStruct_type_list(c *Struct_type_listContext)

	// ExitGet_param_ is called when exiting the get_param_ production.
	ExitGet_param_(c *Get_param_Context)

	// ExitNext_simple_param_ is called when exiting the next_simple_param_ production.
	ExitNext_simple_param_(c *Next_simple_param_Context)

	// ExitReal_base_type_list_ is called when exiting the real_base_type_list_ production.
	ExitReal_base_type_list_(c *Real_base_type_list_Context)

	// ExitVoid_ is called when exiting the void_ production.
	ExitVoid_(c *Void_Context)

	// ExitMethod_param_ is called when exiting the method_param_ production.
	ExitMethod_param_(c *Method_param_Context)

	// ExitSingle_struct_param is called when exiting the single_struct_param production.
	ExitSingle_struct_param(c *Single_struct_paramContext)

	// ExitSimple_param_ is called when exiting the simple_param_ production.
	ExitSimple_param_(c *Simple_param_Context)

	// ExitReal_base_type_parm is called when exiting the real_base_type_parm production.
	ExitReal_base_type_parm(c *Real_base_type_parmContext)

	// ExitReal_base_type_list_parm is called when exiting the real_base_type_list_parm production.
	ExitReal_base_type_list_parm(c *Real_base_type_list_parmContext)

	// ExitNot_login is called when exiting the not_login production.
	ExitNot_login(c *Not_loginContext)

	// ExitMethod_description is called when exiting the method_description production.
	ExitMethod_description(c *Method_descriptionContext)

	// ExitMethod_description_content is called when exiting the method_description_content production.
	ExitMethod_description_content(c *Method_description_contentContext)

	// ExitType_annotations is called when exiting the type_annotations production.
	ExitType_annotations(c *Type_annotationsContext)

	// ExitType_annotation is called when exiting the type_annotation production.
	ExitType_annotation(c *Type_annotationContext)

	// ExitField_annotations is called when exiting the field_annotations production.
	ExitField_annotations(c *Field_annotationsContext)

	// ExitField_annotation is called when exiting the field_annotation production.
	ExitField_annotation(c *Field_annotationContext)

	// ExitAnnotation_value is called when exiting the annotation_value production.
	ExitAnnotation_value(c *Annotation_valueContext)

	// ExitField_type is called when exiting the field_type production.
	ExitField_type(c *Field_typeContext)

	// ExitBase_type is called when exiting the base_type production.
	ExitBase_type(c *Base_typeContext)

	// ExitContainer_type is called when exiting the container_type production.
	ExitContainer_type(c *Container_typeContext)

	// ExitStruct_type is called when exiting the struct_type production.
	ExitStruct_type(c *Struct_typeContext)

	// ExitMap_type is called when exiting the map_type production.
	ExitMap_type(c *Map_typeContext)

	// ExitList_type is called when exiting the list_type production.
	ExitList_type(c *List_typeContext)

	// ExitConst_value is called when exiting the const_value production.
	ExitConst_value(c *Const_valueContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitConst_list is called when exiting the const_list production.
	ExitConst_list(c *Const_listContext)

	// ExitConst_map_entry is called when exiting the const_map_entry production.
	ExitConst_map_entry(c *Const_map_entryContext)

	// ExitConst_map is called when exiting the const_map production.
	ExitConst_map(c *Const_mapContext)

	// ExitList_separator is called when exiting the list_separator production.
	ExitList_separator(c *List_separatorContext)

	// ExitReal_base_type is called when exiting the real_base_type production.
	ExitReal_base_type(c *Real_base_typeContext)

	// ExitMap_key_type is called when exiting the map_key_type production.
	ExitMap_key_type(c *Map_key_typeContext)
}

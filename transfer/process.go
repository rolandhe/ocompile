package transfer

import (
	"errors"
	"github.com/rolandhe/ocompile/parser"
	"strings"
)

type WalkListener struct {
	*parser.BaseServiceListener
	def *Definition
}

func NewWalkListener(def *Definition) *WalkListener {
	return &WalkListener{
		def: def,
	}
}

func (this *WalkListener) EnterNamespace_(ctx *parser.Namespace_Context) {
	if ctx.LITERAL() != nil {
		this.def.Namespace = trimDoubleQuote(ctx.LITERAL().GetText())
		return
	}
	this.def.Namespace = ctx.IDENTIFIER(1).GetText()
}

func (this *WalkListener) EnterWith_client_optional(ctx *parser.With_client_optionalContext) {
	if ctx.GetText() == "true" {
		this.def.WithClient = true
	}
}

func (this *WalkListener) EnterStruct_(ctx *parser.Struct_Context) {
	if this.def.Err != nil {
		return
	}
	st := &StructDefine{
		Name: ctx.IDENTIFIER().GetText(),
	}
	fields := ctx.AllField()
	for _, f := range fields {
		st.addField(createField(f))
	}
	this.def.addStruct(st)
}

func (this *WalkListener) EnterService(ctx *parser.ServiceContext) {
	if this.def.Err != nil {
		return
	}
	svc := &ServiceDefine{
		Name: ctx.IDENTIFIER().GetText(),
	}
	if ctx.Url_() != nil {
		svc.RootUrl = trimDoubleQuote(ctx.Url_().LITERAL().GetText())
	}
	for _, m := range ctx.AllMethod_() {
		if m.Get_() != nil {
			g, err := createSvcGet(m.Get_())
			if err != nil {
				this.def.Err = err
				return
			}
			svc.addGet(g)
			continue
		}
		if m.Post_() != nil {
			p, err := createSvcPost(m.Post_())
			if err != nil {
				this.def.Err = err
				return
			}
			svc.addPost(p)
			continue
		}
		this.def.Err = errors.New("invalid method in service:" + svc.Name + "," + m.GetText())
		return
	}

	this.def.addService(svc)
}

func createSvcGet(ctx parser.IGet_Context) (*GetMethod, error) {
	g := &GetMethod{}
	g.Name = ctx.IDENTIFIER().GetText()
	g.Url = trimDoubleQuote(ctx.Url_().LITERAL().GetText())
	if ctx.Not_login() != nil {
		g.NotLogin = true
	}
	procMethodType(&g.BaseMethod, ctx.Method_type())

	if ctx.Get_param_() == nil {
		g.Params.IsEmpty = true
		return g, nil
	}
	if ctx.Get_param_().Single_struct_param() != nil {
		g.Params.IsSingleStruct = true
		g.Params.StructName = ctx.Get_param_().Single_struct_param().Struct_type().IDENTIFIER().GetText()
		g.Params.StructParamName = ctx.Get_param_().Single_struct_param().IDENTIFIER().GetText()
		return g, nil
	}
	if ctx.Get_param_().AllSimple_param_() == nil {
		return nil, errors.New("invalid parse in " + g.Name + ":" + ctx.Get_param_().GetText())
	}
	for _, simpleCtx := range ctx.Get_param_().AllSimple_param_() {
		bp := &BasicGetParam{}
		g.Params.addBasicParams(bp)
		if simpleCtx.Real_base_type_parm() != nil {
			bp.IsList = false
			bp.TypeName = simpleCtx.Real_base_type_parm().Real_base_type().GetText()
			bp.ParamName = simpleCtx.Real_base_type_parm().IDENTIFIER().GetText()
			continue
		}
		if simpleCtx.Real_base_type_list_parm() != nil {
			baseType := simpleCtx.Real_base_type_list_parm().Real_base_type_list_().Real_base_type()
			if baseType.TYPE_I16() == nil && baseType.TYPE_BYTE() == nil && baseType.TYPE_I32() == nil && baseType.TYPE_I64() == nil && baseType.TYPE_STRING() == nil {
				return nil, errors.New("get list param must be int or string in " + g.Name + ":" + simpleCtx.GetText())
			}
			bp.IsList = true
			bp.ParamName = simpleCtx.Real_base_type_list_parm().IDENTIFIER().GetText()
			bp.TypeName = simpleCtx.Real_base_type_list_parm().Real_base_type_list_().Real_base_type().GetText()
			continue
		}
		return nil, errors.New("invalid parse in " + g.Name + ":" + simpleCtx.GetText())
	}

	return g, nil
}

func createSvcPost(ctx parser.IPost_Context) (*PostMethod, error) {
	p := &PostMethod{}
	p.Name = ctx.IDENTIFIER().GetText()
	p.Url = trimDoubleQuote(ctx.Url_().LITERAL().GetText())
	if ctx.Not_login() != nil {
		p.NotLogin = true
	}
	procMethodType(&p.BaseMethod, ctx.Method_type())
	if ctx.Method_param_() == nil {
		p.Params.IsEmpty = true
		return p, nil
	}

	if ctx.Method_param_().Single_struct_param() != nil {
		p.Params.IsList = false
		p.Params.StructName = ctx.Method_param_().Single_struct_param().Struct_type().IDENTIFIER().GetText()
		p.Params.ParamName = ctx.Method_param_().Single_struct_param().IDENTIFIER().GetText()
		return p, nil
	}
	if ctx.Method_param_().Struct_type_list() != nil {
		p.Params.IsList = true
		p.Params.ParamName = ctx.Method_param_().IDENTIFIER().GetText()
		p.Params.StructName = ctx.Method_param_().Struct_type_list().Struct_type().IDENTIFIER().GetText()
		return p, nil
	}

	return nil, errors.New("invalid parse:" + ctx.GetText())
}

func procMethodType(bm *BaseMethod, mt parser.IMethod_typeContext) {
	if mt.Real_base_type() != nil {
		bm.TypeName = mt.Real_base_type().GetText()
		return
	}
	if mt.Real_base_type_list_() != nil {
		bm.IsList = true
		bm.TypeName = mt.Real_base_type_list_().Real_base_type().GetText()
		return
	}
	if mt.Void_() != nil {
		bm.IsVoid = true
		return
	}
	if mt.Struct_type() != nil {
		bm.IsStruct = true
		bm.TypeName = mt.Struct_type().IDENTIFIER().GetText()
		return
	}
	bm.IsList = true
	bm.IsStruct = true
	bm.TypeName = mt.Struct_type_list().Struct_type().IDENTIFIER().GetText()
}

func createField(f parser.IFieldContext) *Field {
	field := &Field{
		Name: f.IDENTIFIER().GetText(),
	}
	if f.Field_req() != nil {
		field.ReqDefine = f.Field_req().GetText()
	}

	if f.Field_annotations() != nil {
		for _, anno := range f.Field_annotations().AllField_annotation() {
			fa := &Annotation{
				Key:   anno.IDENTIFIER().GetText(),
				Value: trimDoubleQuote(anno.LITERAL().GetText()),
			}
			field.addAnnotation(fa)
		}
	}

	field.Tp = createFieldType(f.Field_type())

	return field
}

func createFieldType(ftCtx parser.IField_typeContext) *FieldType {
	ft := &FieldType{}
	if ftCtx.Base_type() != nil {
		ft.IsBasic = true
		ft.TypeName = ftCtx.Base_type().GetText()
		return ft
	}

	if ftCtx.Struct_type() != nil {
		ft.IsStruct = true
		ft.TypeName = ftCtx.Struct_type().IDENTIFIER().GetText()
		return ft
	}

	if ftCtx.Container_type().Map_type() != nil {
		ft.IsMap = true
		ft.TypeName = ftCtx.Container_type().Map_type().Map_key_type().GetText()
		ft.ValueType = createFieldType(ftCtx.Container_type().Map_type().Field_type())
		return ft
	}

	if ftCtx.Container_type().List_type() != nil {
		ft.IsList = true
		ft.ValueType = createFieldType(ftCtx.Container_type().List_type().Field_type())
		return ft
	}

	return ft
}

func trimDoubleQuote(v string) string {
	return strings.TrimRight(strings.TrimLeft(v, "\""), "\"")
}
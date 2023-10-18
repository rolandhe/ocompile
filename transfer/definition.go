package transfer

type Definition struct {
	Namespace  string
	GoImports  []string
	WithClient bool
	Structs    []*StructDefine
	Services   []*ServiceDefine
	err        error
}

func (def *Definition) addStruct(st *StructDefine) {
	def.Structs = append(def.Structs, st)
}

func (def *Definition) addService(svc *ServiceDefine) {
	def.Services = append(def.Services, svc)
}

type StructDefine struct {
	Name   string
	Fields []*Field
}

func (st *StructDefine) addField(field *Field) {
	st.Fields = append(st.Fields, field)
}

type Field struct {
	ReqDefine   string
	Tp          *FieldType
	Name        string
	Annotations []*Annotation
}

func (f *Field) addAnnotation(anno *Annotation) {
	f.Annotations = append(f.Annotations, anno)
}

type FieldType struct {
	TypeName string
	IsStruct bool
	IsBasic  bool
	IsList   bool
	IsMap    bool

	ValueType *FieldType
}

type Annotation struct {
	Key   string
	Value string
}

type ServiceDefine struct {
	Name    string
	RootUrl string
	Posts   []*PostMethod
	Gets    []*GetMethod
}

func (svc *ServiceDefine) addGet(g *GetMethod) {
	svc.Gets = append(svc.Gets, g)
}

func (svc *ServiceDefine) addPost(p *PostMethod) {
	svc.Posts = append(svc.Posts, p)
}

type BaseMethod struct {
	Name        string
	Url         string
	NotLogin    bool
	Description string
	MethodReturnType
}

type PostMethod struct {
	BaseMethod
	Params PostParam
}

type PostParam struct {
	IsEmpty    bool
	IsList     bool
	StructName string
	ParamName  string
}

type MethodReturnType struct {
	IsList   bool
	IsStruct bool
	IsVoid   bool
	TypeName string
}

type GetMethod struct {
	BaseMethod
	Params GetParam
}
type GetParam struct {
	IsEmpty         bool
	IsSingleStruct  bool
	StructName      string
	StructParamName string

	BasicParams []*BasicGetParam
}

func (gp *GetParam) addBasicParams(bp *BasicGetParam) {
	gp.BasicParams = append(gp.BasicParams, bp)
}

type BasicGetParam struct {
	ReqDefine string
	IsList    bool
	TypeName  string
	ParamName string
}

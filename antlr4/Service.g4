grammar Service;

document
    : header* definition* EOF
    ;

header
    : namespace_
    | golang_import_
    | 'with_client' '=' with_client_optional
    ;


namespace_
    : 'namespace' IDENTIFIER (IDENTIFIER | LITERAL)
    ;

golang_import_
    : 'go_import' IDENTIFIER? LITERAL
    ;



with_client_optional
    : 'true'
    | 'false'
    ;

definition
    : struct_ | service
    ;


struct_
    : 'struct' IDENTIFIER '{' field* '}' type_annotations?
    ;


service
    : 'service' IDENTIFIER url_? '{' method_* '}' type_annotations?
    ;


field
    : field_req? field_type IDENTIFIER field_annotations? list_separator?
    ;


field_req
    : 'required'
    | 'optional'
    ;

method_
    : post_
    | get_
    ;

post_
    : method_description? 'POST' url_ method_type IDENTIFIER '(' method_param_? ')' not_login?  list_separator?
    ;

get_
    : method_description? 'GET' url_ method_type IDENTIFIER '(' get_param_? ')' not_login?  list_separator?
    ;

url_
    : 'url' '=' LITERAL
    ;

method_type
    : real_base_type
    | real_base_type_list_
    | void_
    | struct_type
    | struct_type_list
    ;

struct_type_list
    : 'list' '<' struct_type '>'
    ;

get_param_
    : single_struct_param
    | simple_param_ next_simple_param_*
    ;

next_simple_param_
    : ',' simple_param_
    ;

real_base_type_list_
    : 'list' '<' real_base_type '>'
    ;
void_
    : 'void'
    ;
method_param_
    : single_struct_param
    | struct_type_list IDENTIFIER
    ;

single_struct_param
   : struct_type IDENTIFIER
   ;

simple_param_
    : field_req? real_base_type_parm
    | field_req? real_base_type_list_parm
    ;
real_base_type_parm
    : real_base_type IDENTIFIER
    ;

real_base_type_list_parm
    : real_base_type_list_ IDENTIFIER
    ;

not_login
    : 'not_login'
    ;

method_description
   : '(' 'description' '=' method_description_content  ')'
   ;

method_description_content
  : annotation_value*
  ;

type_annotations
    : '(' type_annotation* ')'
    ;

type_annotation
    : IDENTIFIER ('=' annotation_value)? list_separator?
    ;

field_annotations
    : '(' field_annotation* ')'
    ;

field_annotation
    : IDENTIFIER '=' LITERAL list_separator?
    ;

annotation_value
    : integer | LITERAL
    ;


field_type
    : base_type | struct_type | container_type
    ;

base_type
    : real_base_type type_annotations?
    ;

container_type
    : (map_type | list_type) type_annotations?
    ;

struct_type
   : 'struct' IDENTIFIER
   ;

map_type
    : 'map' '<' map_key_type COMMA field_type '>'
    ;

list_type
    : 'list' '<' field_type '>'
    ;


const_value
    : integer | DOUBLE | LITERAL | IDENTIFIER | const_list | const_map
    ;

integer
    : INTEGER | HEX_INTEGER
    ;

INTEGER
    : ('+' | '-')? DIGIT+
    ;

HEX_INTEGER
    : '-'? '0x' HEX_DIGIT+
    ;

DOUBLE
    : ('+' | '-')? ( DIGIT+ ('.' DIGIT+)? | '.' DIGIT+ ) (('E' | 'e') INTEGER)?
    ;

const_list
    : '[' (const_value list_separator?)* ']'
    ;

const_map_entry
    : const_value ':' const_value list_separator?
    ;

const_map
    : '{' const_map_entry* '}'
    ;

list_separator
    : COMMA | ';'
    ;

real_base_type
    :  TYPE_BOOL | TYPE_BYTE | TYPE_I16 | TYPE_I32 | TYPE_I64 | TYPE_DOUBLE | TYPE_STRING
    ;

map_key_type
    : TYPE_BYTE | TYPE_I16 | TYPE_I32 | TYPE_I64 | TYPE_DOUBLE | TYPE_STRING
    ;

TYPE_BOOL: 'bool';
TYPE_BYTE: 'byte';
TYPE_I16: 'i16';
TYPE_I32: 'i32';
TYPE_I64: 'i64';
TYPE_DOUBLE: 'double';
TYPE_STRING: 'string';
//TYPE_BINARY: 'binary';

LITERAL
    : '"' ( ESC_SEQ | ~[\\"] )* '"'
    | '\'' ( ESC_SEQ | ~[\\'] )* '\''
    ;

fragment ESC_SEQ : '\\' [rnt"'\\] ;


IDENTIFIER
    : (LETTER | '_') (LETTER | DIGIT | '.' | '_')*
    ;


COMMA
    : ','
    ;

fragment LETTER
    : 'A'..'Z' | 'a'..'z'
    ;

fragment DIGIT
    : '0'..'9'
    ;

fragment HEX_DIGIT
    : DIGIT | 'A'..'F' | 'a'..'f'
    ;

WS
    : (' ' | '\t' | '\r' '\n' | '\n')+ -> channel(HIDDEN)
    ;

SL_COMMENT
    : ('//' | '#') (~'\n')* ('\r')? '\n' -> channel(HIDDEN)
    ;

ML_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;
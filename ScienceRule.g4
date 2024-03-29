grammar ScienceRule;

sciencerule
    : '(' sciencerule ')'                                                      #parenExp
    | sciencerule SP LOGICAL_OPERATOR SP sciencerule                           #logicalExp
    | attrPath SP 'pr'                                                        #presentExp
    | lhs=sciencerule SP op=( ADD | SUB | MUL | DIV | MOD ) SP rhs=sciencerule       #arithmeticExp
    | lhs=sciencerule SP op=( EQ | NE | GT | LT | GE | LE | CO ) SP rhs=sciencerule       #compareExp
    | atom=value                                                            #valueExp
    ;

value
   : BOOLEAN           #boolean
   | NULL              #null
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
   | attrPath          #topic
   ;

ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';

EQ : 'eq' | 'EQ' | '==';
NE : 'ne' | 'NE' | '!=';
GT : 'gt' | 'GT' | '>';
LT : 'lt' | 'LT' | '<';
GE : 'ge' | 'GE' | '>=';
LE : 'le' | 'LE' | '<=';
CO : 'co' | 'CO';

LOGICAL_OPERATOR
    : 'and' | 'or'
    ;

BOOLEAN
    : 'true' | 'false'
    ;

NULL
    : 'null'
    ;

attrPath
    : ATTRNAME subAttr?
    ;

subAttr
    : '.' attrPath
    ;

ATTRNAME
    : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

COMMA
   : ',' ' '*;

SP
   : ' ' NEWLINE*
   ;
grammar TinyLanguage; // rename to distinguish from Expr.g4

prog: block;

block: ( statement | functionDecl)* ( Return expression ';')?;

statement:
	assignment ';'
	| functionCall ';'
	| ifStatement
	| forStatement
	| whileStatement;

assignment: Identifier indexes? '=' expression;

functionDecl: Def Identifier '(' idList? ')' block End;

forStatement:
	For Identifier '=' expression To expression Do block End;

whileStatement: While expression Do block End;

functionCall:
	Identifier '(' exprList? ')' # identifierFunctionCall
	// built-in
	| Println '(' expression? ')' # printlnFunctionCall;

ifStatement: ifStat elseIfStat* elseStat? End;

ifStat: If expression Do block;

elseIfStat: Else If expression Do block;

elseStat: Else Do block;

idList: Identifier ( ',' Identifier)*;

exprList: expression ( ',' expression)*;

expression:
	'-' expression											# unaryMinusExpression
	| <assoc = right> expression '^' expression				# powerExpression
	| expression op = ('*' | '/' | '%') expression			# multExpression
	| expression op = ('+' | '-') expression				# addExpression
	| expression op = ('>=' | '<=' | '>' | '<') expression	# compExpression
	| expression op = ('==' | '!=') expression				# eqExpression
	| expression '&&' expression							# andExpression
	| expression '||' expression							# orExpression
	| Number												# numberExpression
	| Bool													# boolExpression
	| functionCall											# functionCallExpression
	| list indexes?											# listExpression
	| Identifier indexes?									# identifierExpression
	| StringLiteral											# stringExpression
	| '(' expression ')'									# expressionExpression;

list: '[' exprList? ']';

indexes: ( '[' expression ']')+;
Println: 'println';
Def: 'def';
If: 'if';
Else: 'else';
Do: 'do';
Return: 'return';
For: 'for';
While: 'while';
To: 'to';
End: 'end';

Bool: 'true' | 'false';

Number: Int ( '.' Digit*)?;

Identifier: [a-zA-Z] [a-zA-Z_0-9]*;

StringLiteral:
	["] (~["\r\n\\] | '\\' ~[\r\n])* ["]
	| ['] ( ~['\r\n\\] | '\\' ~[\r\n])* ['];

// Comment: ( '//' ~[\r\n]* | '/*' .*? '*/') -> skip ;

Space: [ \t\r\n\u000C] -> skip;

fragment Int: [1-9] Digit* | '0';

fragment Digit: [0-9];

Or: '||';
And: '&&';
Equals: '==';
NEquals: '!=';
GTEquals: '>=';
LTEquals: '<=';
Power: '^';
Add: '+';
Subtract: '-';
Mult: '*';
Divide: '/';
Mod: '%';
GT: '>';
LT: '<';
grammar TinyLanguage; // rename to distinguish from Expr.g4

prog: block;

block: ( statement )*;
// block: ( statement | functionDecl)* ( Return expression ';')?;

statement:
	assignment ';'
	| functionCall ';'
	;

assignment: Identifier '=' expression;

functionCall:
	// Identifier '(' exprList? ')'	# identifierFunctionCall
    // built-in
	Println '(' expression? ')'	# printlnFunctionCall
    ;

// functionDecl: Def Identifier '(' idList? ')' block End;

// idList: Identifier ( ',' Identifier)*;

// exprList: expression ( ',' expression)*;

expression:
	// '-' expression											# unaryMinusExpression
	// | <assoc = right> expression '^' expression				# powerExpression
	expression op = ('*' | '/' | '%') expression			# multExpression
	| expression op = ('+' | '-') expression				# addExpression
	| Number												# numberExpression
	| functionCall								# functionCallExpression
	| Identifier									# identifierExpression
	| '(' expression ')'							# expressionExpression
    ;

Println: 'println';
// Def: 'def';
// Return: 'return';
// End: 'end';

Number: Int ( '.' Digit*)?
    ;

Identifier: [a-zA-Z] [a-zA-Z_0-9]*
    ;

StringLiteral: ["] (~["\r\n\\] | '\\' ~[\r\n])* ["]
	| ['] ( ~['\r\n\\] | '\\' ~[\r\n])* [']
    ;

// Comment: ( '//' ~[\r\n]* | '/*' .*? '*/') -> skip
//     ;

Space: [ \t\r\n\u000C] -> skip
    ;

fragment Int: [1-9] Digit* 
    | '0'
    ;

fragment Digit: [0-9]
    ;

Add      : '+';
Subtract : '-';
Mult    : '*';
Divide 	: '/';
Mod 	: '%';
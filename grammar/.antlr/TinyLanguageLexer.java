// Generated from e:\Projects\tiny-language-antlr4-llvm-ir\grammar\TinyLanguage.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TinyLanguageLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, Println=5, Number=6, Identifier=7, StringLiteral=8, 
		Space=9, Add=10, Subtract=11;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "Println", "Number", "Identifier", "StringLiteral", 
			"Space", "Int", "Digit", "Add", "Subtract"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'='", "'('", "')'", "'println'", null, null, null, null, 
			"'+'", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "Println", "Number", "Identifier", "StringLiteral", 
			"Space", "Add", "Subtract"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public TinyLanguageLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "TinyLanguage.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\rh\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\7\7\61\n\7\f\7\16\7\64\13\7\5\7\66"+
		"\n\7\3\b\3\b\7\b:\n\b\f\b\16\b=\13\b\3\t\3\t\3\t\3\t\7\tC\n\t\f\t\16\t"+
		"F\13\t\3\t\3\t\3\t\3\t\3\t\7\tM\n\t\f\t\16\tP\13\t\3\t\5\tS\n\t\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\7\13[\n\13\f\13\16\13^\13\13\3\13\5\13a\n\13\3\f"+
		"\3\f\3\r\3\r\3\16\3\16\2\2\17\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\2\27\2\31\f\33\r\3\2\f\4\2C\\c|\6\2\62;C\\aac|\3\2$$\6\2\f\f\17\17"+
		"$$^^\4\2\f\f\17\17\3\2))\6\2\f\f\17\17))^^\5\2\13\f\16\17\"\"\3\2\63;"+
		"\3\2\62;\2o\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2"+
		"\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\31\3\2\2\2\2\33"+
		"\3\2\2\2\3\35\3\2\2\2\5\37\3\2\2\2\7!\3\2\2\2\t#\3\2\2\2\13%\3\2\2\2\r"+
		"-\3\2\2\2\17\67\3\2\2\2\21R\3\2\2\2\23T\3\2\2\2\25`\3\2\2\2\27b\3\2\2"+
		"\2\31d\3\2\2\2\33f\3\2\2\2\35\36\7=\2\2\36\4\3\2\2\2\37 \7?\2\2 \6\3\2"+
		"\2\2!\"\7*\2\2\"\b\3\2\2\2#$\7+\2\2$\n\3\2\2\2%&\7r\2\2&\'\7t\2\2\'(\7"+
		"k\2\2()\7p\2\2)*\7v\2\2*+\7n\2\2+,\7p\2\2,\f\3\2\2\2-\65\5\25\13\2.\62"+
		"\7\60\2\2/\61\5\27\f\2\60/\3\2\2\2\61\64\3\2\2\2\62\60\3\2\2\2\62\63\3"+
		"\2\2\2\63\66\3\2\2\2\64\62\3\2\2\2\65.\3\2\2\2\65\66\3\2\2\2\66\16\3\2"+
		"\2\2\67;\t\2\2\28:\t\3\2\298\3\2\2\2:=\3\2\2\2;9\3\2\2\2;<\3\2\2\2<\20"+
		"\3\2\2\2=;\3\2\2\2>D\t\4\2\2?C\n\5\2\2@A\7^\2\2AC\n\6\2\2B?\3\2\2\2B@"+
		"\3\2\2\2CF\3\2\2\2DB\3\2\2\2DE\3\2\2\2EG\3\2\2\2FD\3\2\2\2GS\t\4\2\2H"+
		"N\t\7\2\2IM\n\b\2\2JK\7^\2\2KM\n\6\2\2LI\3\2\2\2LJ\3\2\2\2MP\3\2\2\2N"+
		"L\3\2\2\2NO\3\2\2\2OQ\3\2\2\2PN\3\2\2\2QS\t\7\2\2R>\3\2\2\2RH\3\2\2\2"+
		"S\22\3\2\2\2TU\t\t\2\2UV\3\2\2\2VW\b\n\2\2W\24\3\2\2\2X\\\t\n\2\2Y[\5"+
		"\27\f\2ZY\3\2\2\2[^\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]a\3\2\2\2^\\\3\2\2\2"+
		"_a\7\62\2\2`X\3\2\2\2`_\3\2\2\2a\26\3\2\2\2bc\t\13\2\2c\30\3\2\2\2de\7"+
		"-\2\2e\32\3\2\2\2fg\7/\2\2g\34\3\2\2\2\r\2\62\65;BDLNR\\`\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
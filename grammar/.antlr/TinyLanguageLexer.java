// Generated from /Users/yidongl/Desktop/Study/compiler/tiny-language-antlr4/grammar/TinyLanguage.g4 by ANTLR 4.9.2
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
		Space=9, Add=10, Subtract=11, Mult=12, Divide=13, Mod=14;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "Println", "Number", "Identifier", "StringLiteral", 
			"Space", "Int", "Digit", "Add", "Subtract", "Mult", "Divide", "Mod"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'='", "'('", "')'", "'println'", null, null, null, null, 
			"'+'", "'-'", "'*'", "'/'", "'%'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "Println", "Number", "Identifier", "StringLiteral", 
			"Space", "Add", "Subtract", "Mult", "Divide", "Mod"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\20t\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2\3"+
		"\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\7\7"+
		"\67\n\7\f\7\16\7:\13\7\5\7<\n\7\3\b\3\b\7\b@\n\b\f\b\16\bC\13\b\3\t\3"+
		"\t\3\t\3\t\7\tI\n\t\f\t\16\tL\13\t\3\t\3\t\3\t\3\t\3\t\7\tS\n\t\f\t\16"+
		"\tV\13\t\3\t\5\tY\n\t\3\n\3\n\3\n\3\n\3\13\3\13\7\13a\n\13\f\13\16\13"+
		"d\13\13\3\13\5\13g\n\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20"+
		"\3\21\3\21\2\2\22\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\2\27\2\31"+
		"\f\33\r\35\16\37\17!\20\3\2\f\4\2C\\c|\6\2\62;C\\aac|\3\2$$\6\2\f\f\17"+
		"\17$$^^\4\2\f\f\17\17\3\2))\6\2\f\f\17\17))^^\5\2\13\f\16\17\"\"\3\2\63"+
		";\3\2\62;\2{\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2"+
		"\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\31\3\2\2\2\2"+
		"\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\3#\3\2\2\2\5%\3\2\2\2"+
		"\7\'\3\2\2\2\t)\3\2\2\2\13+\3\2\2\2\r\63\3\2\2\2\17=\3\2\2\2\21X\3\2\2"+
		"\2\23Z\3\2\2\2\25f\3\2\2\2\27h\3\2\2\2\31j\3\2\2\2\33l\3\2\2\2\35n\3\2"+
		"\2\2\37p\3\2\2\2!r\3\2\2\2#$\7=\2\2$\4\3\2\2\2%&\7?\2\2&\6\3\2\2\2\'("+
		"\7*\2\2(\b\3\2\2\2)*\7+\2\2*\n\3\2\2\2+,\7r\2\2,-\7t\2\2-.\7k\2\2./\7"+
		"p\2\2/\60\7v\2\2\60\61\7n\2\2\61\62\7p\2\2\62\f\3\2\2\2\63;\5\25\13\2"+
		"\648\7\60\2\2\65\67\5\27\f\2\66\65\3\2\2\2\67:\3\2\2\28\66\3\2\2\289\3"+
		"\2\2\29<\3\2\2\2:8\3\2\2\2;\64\3\2\2\2;<\3\2\2\2<\16\3\2\2\2=A\t\2\2\2"+
		">@\t\3\2\2?>\3\2\2\2@C\3\2\2\2A?\3\2\2\2AB\3\2\2\2B\20\3\2\2\2CA\3\2\2"+
		"\2DJ\t\4\2\2EI\n\5\2\2FG\7^\2\2GI\n\6\2\2HE\3\2\2\2HF\3\2\2\2IL\3\2\2"+
		"\2JH\3\2\2\2JK\3\2\2\2KM\3\2\2\2LJ\3\2\2\2MY\t\4\2\2NT\t\7\2\2OS\n\b\2"+
		"\2PQ\7^\2\2QS\n\6\2\2RO\3\2\2\2RP\3\2\2\2SV\3\2\2\2TR\3\2\2\2TU\3\2\2"+
		"\2UW\3\2\2\2VT\3\2\2\2WY\t\7\2\2XD\3\2\2\2XN\3\2\2\2Y\22\3\2\2\2Z[\t\t"+
		"\2\2[\\\3\2\2\2\\]\b\n\2\2]\24\3\2\2\2^b\t\n\2\2_a\5\27\f\2`_\3\2\2\2"+
		"ad\3\2\2\2b`\3\2\2\2bc\3\2\2\2cg\3\2\2\2db\3\2\2\2eg\7\62\2\2f^\3\2\2"+
		"\2fe\3\2\2\2g\26\3\2\2\2hi\t\13\2\2i\30\3\2\2\2jk\7-\2\2k\32\3\2\2\2l"+
		"m\7/\2\2m\34\3\2\2\2no\7,\2\2o\36\3\2\2\2pq\7\61\2\2q \3\2\2\2rs\7\'\2"+
		"\2s\"\3\2\2\2\r\28;AHJRTXbf\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
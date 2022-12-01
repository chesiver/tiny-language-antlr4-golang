// Generated from e:\Projects\tiny-language-antlr4\grammar\TinyLanguage.g4 by ANTLR 4.9.2
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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, Println=6, Def=7, Return=8, End=9, 
		Number=10, Identifier=11, StringLiteral=12, Space=13, Power=14, Add=15, 
		Subtract=16, Mult=17, Divide=18, Mod=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "Println", "Def", "Return", "End", 
			"Number", "Identifier", "StringLiteral", "Space", "Int", "Digit", "Power", 
			"Add", "Subtract", "Mult", "Divide", "Mod"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'='", "'('", "')'", "','", "'println'", "'def'", "'return'", 
			"'end'", null, null, null, null, "'^'", "'+'", "'-'", "'*'", "'/'", "'%'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, "Println", "Def", "Return", "End", 
			"Number", "Identifier", "StringLiteral", "Space", "Power", "Add", "Subtract", 
			"Mult", "Divide", "Mod"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\25\u0091\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\3\3\3\3\3\4\3\4"+
		"\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13\7\13R\n\13\f\13"+
		"\16\13U\13\13\5\13W\n\13\3\f\3\f\7\f[\n\f\f\f\16\f^\13\f\3\r\3\r\3\r\3"+
		"\r\7\rd\n\r\f\r\16\rg\13\r\3\r\3\r\3\r\3\r\3\r\7\rn\n\r\f\r\16\rq\13\r"+
		"\3\r\5\rt\n\r\3\16\3\16\3\16\3\16\3\17\3\17\7\17|\n\17\f\17\16\17\177"+
		"\13\17\3\17\5\17\u0082\n\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\25\3\25\3\26\3\26\2\2\27\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n"+
		"\23\13\25\f\27\r\31\16\33\17\35\2\37\2!\20#\21%\22\'\23)\24+\25\3\2\f"+
		"\4\2C\\c|\6\2\62;C\\aac|\3\2$$\6\2\f\f\17\17$$^^\4\2\f\f\17\17\3\2))\6"+
		"\2\f\f\17\17))^^\5\2\13\f\16\17\"\"\3\2\63;\3\2\62;\2\u0098\2\3\3\2\2"+
		"\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3"+
		"\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2"+
		"\2\2\33\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2"+
		"\2\2+\3\2\2\2\3-\3\2\2\2\5/\3\2\2\2\7\61\3\2\2\2\t\63\3\2\2\2\13\65\3"+
		"\2\2\2\r\67\3\2\2\2\17?\3\2\2\2\21C\3\2\2\2\23J\3\2\2\2\25N\3\2\2\2\27"+
		"X\3\2\2\2\31s\3\2\2\2\33u\3\2\2\2\35\u0081\3\2\2\2\37\u0083\3\2\2\2!\u0085"+
		"\3\2\2\2#\u0087\3\2\2\2%\u0089\3\2\2\2\'\u008b\3\2\2\2)\u008d\3\2\2\2"+
		"+\u008f\3\2\2\2-.\7=\2\2.\4\3\2\2\2/\60\7?\2\2\60\6\3\2\2\2\61\62\7*\2"+
		"\2\62\b\3\2\2\2\63\64\7+\2\2\64\n\3\2\2\2\65\66\7.\2\2\66\f\3\2\2\2\67"+
		"8\7r\2\289\7t\2\29:\7k\2\2:;\7p\2\2;<\7v\2\2<=\7n\2\2=>\7p\2\2>\16\3\2"+
		"\2\2?@\7f\2\2@A\7g\2\2AB\7h\2\2B\20\3\2\2\2CD\7t\2\2DE\7g\2\2EF\7v\2\2"+
		"FG\7w\2\2GH\7t\2\2HI\7p\2\2I\22\3\2\2\2JK\7g\2\2KL\7p\2\2LM\7f\2\2M\24"+
		"\3\2\2\2NV\5\35\17\2OS\7\60\2\2PR\5\37\20\2QP\3\2\2\2RU\3\2\2\2SQ\3\2"+
		"\2\2ST\3\2\2\2TW\3\2\2\2US\3\2\2\2VO\3\2\2\2VW\3\2\2\2W\26\3\2\2\2X\\"+
		"\t\2\2\2Y[\t\3\2\2ZY\3\2\2\2[^\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]\30\3\2\2"+
		"\2^\\\3\2\2\2_e\t\4\2\2`d\n\5\2\2ab\7^\2\2bd\n\6\2\2c`\3\2\2\2ca\3\2\2"+
		"\2dg\3\2\2\2ec\3\2\2\2ef\3\2\2\2fh\3\2\2\2ge\3\2\2\2ht\t\4\2\2io\t\7\2"+
		"\2jn\n\b\2\2kl\7^\2\2ln\n\6\2\2mj\3\2\2\2mk\3\2\2\2nq\3\2\2\2om\3\2\2"+
		"\2op\3\2\2\2pr\3\2\2\2qo\3\2\2\2rt\t\7\2\2s_\3\2\2\2si\3\2\2\2t\32\3\2"+
		"\2\2uv\t\t\2\2vw\3\2\2\2wx\b\16\2\2x\34\3\2\2\2y}\t\n\2\2z|\5\37\20\2"+
		"{z\3\2\2\2|\177\3\2\2\2}{\3\2\2\2}~\3\2\2\2~\u0082\3\2\2\2\177}\3\2\2"+
		"\2\u0080\u0082\7\62\2\2\u0081y\3\2\2\2\u0081\u0080\3\2\2\2\u0082\36\3"+
		"\2\2\2\u0083\u0084\t\13\2\2\u0084 \3\2\2\2\u0085\u0086\7`\2\2\u0086\""+
		"\3\2\2\2\u0087\u0088\7-\2\2\u0088$\3\2\2\2\u0089\u008a\7/\2\2\u008a&\3"+
		"\2\2\2\u008b\u008c\7,\2\2\u008c(\3\2\2\2\u008d\u008e\7\61\2\2\u008e*\3"+
		"\2\2\2\u008f\u0090\7\'\2\2\u0090,\3\2\2\2\r\2SV\\cemos}\u0081\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
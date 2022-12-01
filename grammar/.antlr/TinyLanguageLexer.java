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
		Space=9, Power=10, Add=11, Subtract=12, Mult=13, Divide=14, Mod=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "Println", "Number", "Identifier", "StringLiteral", 
			"Space", "Int", "Digit", "Power", "Add", "Subtract", "Mult", "Divide", 
			"Mod"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'='", "'('", "')'", "'println'", null, null, null, null, 
			"'^'", "'+'", "'-'", "'*'", "'/'", "'%'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "Println", "Number", "Identifier", "StringLiteral", 
			"Space", "Power", "Add", "Subtract", "Mult", "Divide", "Mod"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\21x\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3"+
		"\7\3\7\7\79\n\7\f\7\16\7<\13\7\5\7>\n\7\3\b\3\b\7\bB\n\b\f\b\16\bE\13"+
		"\b\3\t\3\t\3\t\3\t\7\tK\n\t\f\t\16\tN\13\t\3\t\3\t\3\t\3\t\3\t\7\tU\n"+
		"\t\f\t\16\tX\13\t\3\t\5\t[\n\t\3\n\3\n\3\n\3\n\3\13\3\13\7\13c\n\13\f"+
		"\13\16\13f\13\13\3\13\5\13i\n\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3"+
		"\20\3\20\3\21\3\21\3\22\3\22\2\2\23\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n"+
		"\23\13\25\2\27\2\31\f\33\r\35\16\37\17!\20#\21\3\2\f\4\2C\\c|\6\2\62;"+
		"C\\aac|\3\2$$\6\2\f\f\17\17$$^^\4\2\f\f\17\17\3\2))\6\2\f\f\17\17))^^"+
		"\5\2\13\f\16\17\"\"\3\2\63;\3\2\62;\2\177\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2"+
		"!\3\2\2\2\2#\3\2\2\2\3%\3\2\2\2\5\'\3\2\2\2\7)\3\2\2\2\t+\3\2\2\2\13-"+
		"\3\2\2\2\r\65\3\2\2\2\17?\3\2\2\2\21Z\3\2\2\2\23\\\3\2\2\2\25h\3\2\2\2"+
		"\27j\3\2\2\2\31l\3\2\2\2\33n\3\2\2\2\35p\3\2\2\2\37r\3\2\2\2!t\3\2\2\2"+
		"#v\3\2\2\2%&\7=\2\2&\4\3\2\2\2\'(\7?\2\2(\6\3\2\2\2)*\7*\2\2*\b\3\2\2"+
		"\2+,\7+\2\2,\n\3\2\2\2-.\7r\2\2./\7t\2\2/\60\7k\2\2\60\61\7p\2\2\61\62"+
		"\7v\2\2\62\63\7n\2\2\63\64\7p\2\2\64\f\3\2\2\2\65=\5\25\13\2\66:\7\60"+
		"\2\2\679\5\27\f\28\67\3\2\2\29<\3\2\2\2:8\3\2\2\2:;\3\2\2\2;>\3\2\2\2"+
		"<:\3\2\2\2=\66\3\2\2\2=>\3\2\2\2>\16\3\2\2\2?C\t\2\2\2@B\t\3\2\2A@\3\2"+
		"\2\2BE\3\2\2\2CA\3\2\2\2CD\3\2\2\2D\20\3\2\2\2EC\3\2\2\2FL\t\4\2\2GK\n"+
		"\5\2\2HI\7^\2\2IK\n\6\2\2JG\3\2\2\2JH\3\2\2\2KN\3\2\2\2LJ\3\2\2\2LM\3"+
		"\2\2\2MO\3\2\2\2NL\3\2\2\2O[\t\4\2\2PV\t\7\2\2QU\n\b\2\2RS\7^\2\2SU\n"+
		"\6\2\2TQ\3\2\2\2TR\3\2\2\2UX\3\2\2\2VT\3\2\2\2VW\3\2\2\2WY\3\2\2\2XV\3"+
		"\2\2\2Y[\t\7\2\2ZF\3\2\2\2ZP\3\2\2\2[\22\3\2\2\2\\]\t\t\2\2]^\3\2\2\2"+
		"^_\b\n\2\2_\24\3\2\2\2`d\t\n\2\2ac\5\27\f\2ba\3\2\2\2cf\3\2\2\2db\3\2"+
		"\2\2de\3\2\2\2ei\3\2\2\2fd\3\2\2\2gi\7\62\2\2h`\3\2\2\2hg\3\2\2\2i\26"+
		"\3\2\2\2jk\t\13\2\2k\30\3\2\2\2lm\7`\2\2m\32\3\2\2\2no\7-\2\2o\34\3\2"+
		"\2\2pq\7/\2\2q\36\3\2\2\2rs\7,\2\2s \3\2\2\2tu\7\61\2\2u\"\3\2\2\2vw\7"+
		"\'\2\2w$\3\2\2\2\r\2:=CJLTVZdh\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
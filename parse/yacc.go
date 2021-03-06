//line yacc.y:2
package parse

import __yyfmt__ "fmt"

//line yacc.y:3
import (
	"github.com/nelhage/gollum"
)

var keywords = map[string]token{
	"if":   tokIf,
	"else": tokElse,
	"let":  tokLet,
	"rec":  tokRec,
	"in":   tokIn,

	"fn":    tokFunc,
	"true":  tokBoolean,
	"false": tokBoolean,
}

//line yacc.y:23
type yySymType struct {
	yys  int
	ast  gollum.AST
	asts []gollum.AST
	tok  *tokenStruct
}

const tokFunc = 57346
const tokIf = 57347
const tokElse = 57348
const tokLet = 57349
const tokRec = 57350
const tokIn = 57351
const tokBoolean = 57352
const tokNumber = 57353
const tokStr = 57354
const tokIdent = 57355
const tokArrow = 57356
const tokIntType = 57357

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokFunc",
	"tokIf",
	"tokElse",
	"tokLet",
	"tokRec",
	"tokIn",
	"tokBoolean",
	"tokNumber",
	"tokStr",
	"tokIdent",
	"tokArrow",
	"tokIntType",
	"','",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"':'",
	"'+'",
	"'-'",
	"'/'",
	"'*'",
	"'!'",
	"'='",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc.y:321

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 121

var yyAct = [...]int{

	4, 10, 39, 38, 35, 61, 62, 24, 26, 20,
	70, 25, 77, 23, 31, 29, 30, 66, 32, 65,
	42, 34, 24, 41, 64, 67, 25, 47, 48, 49,
	50, 51, 26, 33, 53, 71, 56, 19, 17, 43,
	21, 59, 68, 14, 15, 16, 18, 60, 3, 42,
	13, 40, 20, 36, 1, 73, 63, 58, 40, 22,
	72, 69, 75, 37, 74, 44, 55, 54, 46, 79,
	45, 80, 81, 12, 82, 83, 76, 11, 19, 17,
	78, 21, 9, 8, 14, 15, 16, 18, 7, 26,
	6, 13, 57, 20, 28, 27, 29, 30, 5, 26,
	22, 20, 2, 0, 28, 27, 29, 30, 26, 52,
	0, 0, 26, 28, 27, 29, 30, 28, 27, 29,
	30,
}
var yyPact = [...]int{

	33, -1000, -1000, 9, 95, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 74, -1000, -1000, -1000, 74, -1000, 16,
	74, 45, 74, 35, -1000, 9, 74, 74, 74, 74,
	74, 91, 82, 38, 72, 48, 38, 31, -1000, -22,
	-15, 15, 9, 6, 1, 7, 26, 95, -9, -9,
	15, 15, -1000, 55, -8, 19, -1000, -1000, -10, 46,
	38, 74, 9, 35, -1000, -1000, -6, -1000, 74, -10,
	-10, 38, -1000, -10, -1000, 95, 35, -1000, 35, 95,
	-1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 102, 0, 1, 98, 90, 88, 83, 82, 77,
	73, 70, 68, 67, 66, 2, 13, 65, 4, 63,
	3, 54,
}
var yyR1 = [...]int{

	0, 21, 21, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 4, 4, 4, 5, 6, 7,
	13, 13, 13, 14, 14, 15, 15, 16, 16, 16,
	16, 16, 17, 17, 8, 11, 11, 11, 12, 12,
	9, 9, 18, 18, 18, 19, 19, 20, 10, 10,
	10, 10, 10,
}
var yyR2 = [...]int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 1, 1, 1, 5, 1, 5,
	0, 1, 2, 1, 3, 1, 3, 1, 3, 3,
	3, 4, 1, 3, 4, 0, 1, 2, 1, 3,
	4, 5, 0, 1, 2, 1, 3, 3, 3, 3,
	3, 3, 2,
}
var yyChk = [...]int{

	-1000, -21, -1, 15, -2, -4, -5, -6, -7, -8,
	-3, -9, -10, 17, 10, 11, 12, 5, 13, 4,
	19, 7, 26, -16, 13, 17, 17, 23, 22, 24,
	25, -2, -2, 17, -2, -18, 8, -19, -20, -15,
	13, -2, 14, -16, -17, -11, -12, -2, -2, -2,
	-2, -2, 18, -3, -13, -14, -15, 20, 9, -18,
	16, 27, 21, -16, 18, 18, 16, 18, 16, 6,
	18, 16, -3, 9, -20, -2, -16, 18, -16, -2,
	-3, -3, -15, -3,
}
var yyDef = [...]int{

	0, -2, 1, 0, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 0, 14, 15, 16, 0, 18, 0,
	0, 42, 0, 2, 27, 0, 35, 0, 0, 0,
	0, 0, 0, 20, 0, 0, 42, 43, 45, 0,
	25, 52, 0, 32, 0, 0, 36, 38, 48, 49,
	50, 51, 12, 0, 0, 21, 23, 13, 0, 0,
	44, 0, 0, 28, 29, 30, 0, 34, 37, 0,
	0, 22, 40, 0, 46, 47, 26, 31, 33, 39,
	17, 19, 24, 41,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 26, 3, 3, 3, 3, 3, 3,
	17, 18, 25, 22, 16, 23, 3, 24, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 21, 3,
	3, 27, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 3, 20,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:65
		{
			yylex.(*lexer).expression = yyDollar[1].ast
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc.y:69
		{
			yylex.(*lexer).ty = yyDollar[2].ast
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:85
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:91
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:98
		{
			yyVAL.ast = &gollum.Boolean{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string) == "true",
			}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:105
		{
			yyVAL.ast = &gollum.Integer{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(int64),
			}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:112
		{
			yyVAL.ast = &gollum.String{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string),
			}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:121
		{
			yyVAL.ast = &gollum.If{
				Loc:        extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Condition:  yyDollar[2].ast,
				Consequent: yyDollar[3].ast,
				Alternate:  yyDollar[5].ast,
			}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:131
		{
			yyVAL.ast = &gollum.Variable{
				Loc: yyDollar[1].tok.loc,
				Var: yyDollar[1].tok.val.(string),
			}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:140
		{
			yyVAL.ast = &gollum.Abstraction{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Vars: yyDollar[3].asts,
				Body: yyDollar[5].ast,
			}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:149
		{
			yyVAL.asts = []gollum.AST{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:156
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:160
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:166
		{
			yyVAL.ast = &gollum.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: nil,
			}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:174
		{
			yyVAL.ast = &gollum.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: yyDollar[3].ast,
			}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:184
		{
			yyVAL.ast = &gollum.TyName{
				Loc:  yyDollar[1].tok.loc,
				Type: yyDollar[1].tok.val.(string),
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:191
		{
			yyVAL.ast = &gollum.TyArrow{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Dom:   yyDollar[1].ast,
				Range: yyDollar[3].ast,
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:199
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:203
		{
			yyVAL.ast = &gollum.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:210
		{
			yyVAL.ast = &gollum.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:219
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:223
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:229
		{
			yyVAL.ast = &gollum.Application{
				Loc:  extend(yyDollar[1].ast.Location(), yyDollar[4].tok.loc),
				Func: yyDollar[1].ast,
				Args: yyDollar[3].asts,
			}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:238
		{
			yyVAL.asts = []gollum.AST{}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:246
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:250
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:255
		{
			yyVAL.ast = &gollum.Let{
				Loc:      extend(yyDollar[1].tok.loc, yyDollar[4].ast.Location()),
				Bindings: yyDollar[2].asts,
				Body:     yyDollar[4].ast,
			}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:263
		{
			yyVAL.ast = &gollum.Let{
				Loc:       extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Bindings:  yyDollar[3].asts,
				Body:      yyDollar[5].ast,
				Recursive: true,
			}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:273
		{
			yyVAL.asts = []gollum.AST{}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:281
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:285
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:291
		{
			yyVAL.ast = &gollum.NameBinding{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Var:   yyDollar[1].ast,
				Value: yyDollar[3].ast,
			}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:301
		{
			yyVAL.ast = arithmetic(yyDollar[2].tok, yyDollar[1].ast, yyDollar[3].ast)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:305
		{
			yyVAL.ast = arithmetic(yyDollar[2].tok, yyDollar[1].ast, yyDollar[3].ast)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:309
		{
			yyVAL.ast = arithmetic(yyDollar[2].tok, yyDollar[1].ast, yyDollar[3].ast)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:313
		{
			yyVAL.ast = arithmetic(yyDollar[2].tok, yyDollar[1].ast, yyDollar[3].ast)
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc.y:317
		{
			yyVAL.ast = arithmetic(yyDollar[1].tok, yyDollar[2].ast)
		}
	}
	goto yystack /* stack new state and value */
}

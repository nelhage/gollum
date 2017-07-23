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
	"'='",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc.y:294

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 90

var yyAct = [...]int{

	10, 33, 32, 4, 29, 50, 51, 24, 19, 24,
	46, 19, 59, 21, 22, 55, 25, 54, 23, 66,
	26, 35, 22, 28, 56, 53, 23, 42, 40, 45,
	24, 41, 24, 27, 60, 48, 57, 36, 18, 16,
	49, 20, 35, 34, 13, 14, 15, 17, 61, 52,
	30, 12, 63, 19, 64, 34, 62, 47, 58, 69,
	70, 68, 71, 72, 1, 65, 31, 18, 16, 67,
	20, 37, 44, 13, 14, 15, 17, 43, 3, 39,
	12, 38, 19, 11, 9, 8, 7, 6, 5, 2,
}
var yyPact = [...]int{

	63, -1000, -1000, 9, 15, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 34, -1000, -1000, -1000, 34, -1000, 16, 34,
	42, 28, -1000, 9, 34, 13, -8, 30, -10, 48,
	30, 24, -1000, -17, -15, 9, 7, -1, 6, 20,
	15, -1000, 52, -6, 18, -1000, -1000, -11, 47, 30,
	34, 9, 28, -1000, -1000, 1, -1000, 34, -11, -11,
	30, -1000, -11, -1000, 15, 28, -1000, 28, 15, -1000,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 89, 3, 0, 88, 87, 86, 85, 84, 83,
	81, 79, 77, 72, 1, 13, 71, 4, 66, 2,
	64,
}
var yyR1 = [...]int{

	0, 20, 20, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 4, 4, 4, 5, 6, 7, 12,
	12, 12, 13, 13, 14, 14, 15, 15, 15, 15,
	15, 16, 16, 8, 10, 10, 10, 11, 11, 9,
	9, 17, 17, 17, 18, 18, 19,
}
var yyR2 = [...]int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 1, 1, 1, 5, 1, 5, 0,
	1, 2, 1, 3, 1, 3, 1, 3, 3, 3,
	4, 1, 3, 4, 0, 1, 2, 1, 3, 4,
	5, 0, 1, 2, 1, 3, 3,
}
var yyChk = [...]int{

	-1000, -20, -1, 15, -2, -4, -5, -6, -7, -8,
	-3, -9, 17, 10, 11, 12, 5, 13, 4, 19,
	7, -15, 13, 17, 17, -2, -2, 17, -2, -17,
	8, -18, -19, -14, 13, 14, -15, -16, -10, -11,
	-2, 18, -3, -12, -13, -14, 20, 9, -17, 16,
	22, 21, -15, 18, 18, 16, 18, 16, 6, 18,
	16, -3, 9, -19, -2, -15, 18, -15, -2, -3,
	-3, -14, -3,
}
var yyDef = [...]int{

	0, -2, 1, 0, 3, 4, 5, 6, 7, 8,
	9, 10, 0, 13, 14, 15, 0, 17, 0, 0,
	41, 2, 26, 0, 34, 0, 0, 19, 0, 0,
	41, 42, 44, 0, 24, 0, 31, 0, 0, 35,
	37, 11, 0, 0, 20, 22, 12, 0, 0, 43,
	0, 0, 27, 28, 29, 0, 33, 36, 0, 0,
	21, 39, 0, 45, 46, 25, 30, 32, 38, 16,
	18, 23, 40,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	17, 18, 3, 3, 16, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 21, 3,
	3, 22, 3, 3, 3, 3, 3, 3, 3, 3,
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
		//line yacc.y:61
		{
			yylex.(*lexer).expression = yyDollar[1].ast
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc.y:65
		{
			yylex.(*lexer).ty = yyDollar[2].ast
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:80
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:86
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:93
		{
			yyVAL.ast = &gollum.Boolean{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string) == "true",
			}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:100
		{
			yyVAL.ast = &gollum.Integer{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(int64),
			}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:107
		{
			yyVAL.ast = &gollum.String{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string),
			}
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:116
		{
			yyVAL.ast = &gollum.If{
				Loc:        extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Condition:  yyDollar[2].ast,
				Consequent: yyDollar[3].ast,
				Alternate:  yyDollar[5].ast,
			}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:126
		{
			yyVAL.ast = &gollum.Variable{
				Loc: yyDollar[1].tok.loc,
				Var: yyDollar[1].tok.val.(string),
			}
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:135
		{
			yyVAL.ast = &gollum.Abstraction{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Vars: yyDollar[3].asts,
				Body: yyDollar[5].ast,
			}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:144
		{
			yyVAL.asts = []gollum.AST{}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:151
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:155
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:161
		{
			yyVAL.ast = &gollum.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: nil,
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:169
		{
			yyVAL.ast = &gollum.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: yyDollar[3].ast,
			}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:179
		{
			yyVAL.ast = &gollum.TyName{
				Loc:  yyDollar[1].tok.loc,
				Type: yyDollar[1].tok.val.(string),
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:186
		{
			yyVAL.ast = &gollum.TyArrow{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Dom:   yyDollar[1].ast,
				Range: yyDollar[3].ast,
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:194
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:198
		{
			yyVAL.ast = &gollum.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:205
		{
			yyVAL.ast = &gollum.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:214
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:218
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:224
		{
			yyVAL.ast = &gollum.Application{
				Loc:  extend(yyDollar[1].ast.Location(), yyDollar[4].tok.loc),
				Func: yyDollar[1].ast,
				Args: yyDollar[3].asts,
			}
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:233
		{
			yyVAL.asts = []gollum.AST{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:241
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:245
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:250
		{
			yyVAL.ast = &gollum.Let{
				Loc:      extend(yyDollar[1].tok.loc, yyDollar[4].ast.Location()),
				Bindings: yyDollar[2].asts,
				Body:     yyDollar[4].ast,
			}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:258
		{
			yyVAL.ast = &gollum.Let{
				Loc:       extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Bindings:  yyDollar[3].asts,
				Body:      yyDollar[5].ast,
				Recursive: true,
			}
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:268
		{
			yyVAL.asts = []gollum.AST{}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:276
		{
			yyVAL.asts = []gollum.AST{yyDollar[1].ast}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:280
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:286
		{
			yyVAL.ast = &gollum.NameBinding{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Var:   yyDollar[1].ast,
				Value: yyDollar[3].ast,
			}
		}
	}
	goto yystack /* stack new state and value */
}

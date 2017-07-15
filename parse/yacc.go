//line yacc.y:2
package parse

import __yyfmt__ "fmt"

//line yacc.y:3
import (
	"nelhage.com/lambda"
)

//line yacc.y:11
type yySymType struct {
	yys  int
	ast  lambda.AST
	asts []lambda.AST
	tok  *tokenStruct
}

const tokFunc = 57346
const tokIf = 57347
const tokThen = 57348
const tokElse = 57349
const tokEnd = 57350
const tokBoolean = 57351
const tokNumber = 57352
const tokStr = 57353
const tokIdent = 57354
const tokArrow = 57355

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokFunc",
	"tokIf",
	"tokThen",
	"tokElse",
	"tokEnd",
	"tokBoolean",
	"tokNumber",
	"tokStr",
	"tokIdent",
	"tokArrow",
	"','",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc.y:212

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 34
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 59

var yyAct = [...]int{

	41, 2, 29, 36, 16, 16, 52, 25, 51, 45,
	17, 18, 50, 16, 24, 19, 15, 13, 23, 48,
	34, 10, 11, 12, 14, 31, 16, 8, 33, 9,
	42, 38, 26, 43, 37, 20, 39, 35, 40, 16,
	44, 16, 32, 45, 46, 30, 49, 1, 47, 28,
	27, 22, 21, 53, 7, 6, 5, 4, 3,
}
var yyPact = [...]int{

	12, -1000, -10, -1000, -1000, -1000, -1000, -1000, 12, 12,
	-1000, -1000, -1000, 12, -1000, 20, 12, -2, -11, 26,
	33, 9, 28, -10, -1000, -1000, 12, 4, 23, -1000,
	-16, -1000, 12, 24, 12, 33, 18, -10, 12, -10,
	-1000, 30, -1000, 18, 11, 18, -4, -8, -1000, 30,
	-1000, -1000, 18, 30,
}
var yyPgo = [...]int{

	0, 1, 58, 57, 56, 55, 54, 52, 51, 50,
	49, 2, 0, 48, 47,
}
var yyR1 = [...]int{

	0, 14, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 3, 4, 5, 9, 9, 9, 10, 10,
	11, 11, 12, 12, 12, 12, 13, 13, 6, 7,
	7, 7, 8, 8,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 3, 3, 1,
	1, 1, 7, 1, 5, 0, 1, 2, 1, 3,
	1, 3, 1, 3, 3, 3, 1, 3, 4, 0,
	1, 2, 1, 3,
}
var yyChk = [...]int{

	-1000, -14, -1, -2, -3, -4, -5, -6, 15, 17,
	9, 10, 11, 5, 12, 4, 15, -1, -1, -1,
	15, -7, -8, -1, 16, 18, 6, -9, -10, -11,
	12, 16, 14, -1, 16, 14, 19, -1, 7, -1,
	-11, -12, 12, 15, -1, 13, -12, -13, 8, -12,
	16, 16, 14, -12,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 0, 0,
	9, 10, 11, 0, 13, 0, 29, 0, 0, 0,
	15, 0, 30, 32, 7, 8, 0, 0, 16, 18,
	20, 28, 31, 0, 0, 17, 0, 33, 0, 14,
	19, 21, 22, 0, 0, 0, 26, 0, 12, 23,
	24, 25, 0, 27,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 3, 14, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 19, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 17, 3, 18,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
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
		//line yacc.y:41
		{
			yylex.(*lexer).result = yyDollar[1].ast
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:52
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:56
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:63
		{
			yyVAL.ast = &lambda.Boolean{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string) == "true",
			}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:70
		{
			yyVAL.ast = &lambda.Integer{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(int64),
			}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:77
		{
			yyVAL.ast = &lambda.String{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string),
			}
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line yacc.y:86
		{
			yyVAL.ast = &lambda.If{
				Loc:        extend(yyDollar[1].tok.loc, yyDollar[6].ast.Location()),
				Condition:  yyDollar[2].ast,
				Consequent: yyDollar[4].ast,
				Alternate:  yyDollar[6].ast,
			}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:96
		{
			yyVAL.ast = &lambda.Variable{
				Loc: yyDollar[1].tok.loc,
				Var: yyDollar[1].tok.val.(string),
			}
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:105
		{
			yyVAL.ast = &lambda.Abstraction{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Vars: yyDollar[3].asts,
				Body: yyDollar[5].ast,
			}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:114
		{
			yyVAL.asts = []lambda.AST{}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:121
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:125
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:131
		{
			yyVAL.ast = &lambda.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: nil,
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:139
		{
			yyVAL.ast = &lambda.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: yyDollar[3].ast,
			}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:149
		{
			yyVAL.ast = &lambda.TyName{
				Loc:  yyDollar[1].tok.loc,
				Type: yyDollar[1].tok.val.(string),
			}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:156
		{
			yyVAL.ast = &lambda.TyArrow{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Dom:   yyDollar[1].ast,
				Range: yyDollar[3].ast,
			}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:164
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:168
		{
			yyVAL.ast = &lambda.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:177
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:181
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:187
		{
			yyVAL.ast = &lambda.Application{
				Loc:  extend(yyDollar[1].ast.Location(), yyDollar[4].tok.loc),
				Func: yyDollar[1].ast,
				Args: yyDollar[3].asts,
			}
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:196
		{
			yyVAL.asts = []lambda.AST{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:204
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:208
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	}
	goto yystack /* stack new state and value */
}

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
const tokElse = 57348
const tokLet = 57349
const tokIn = 57350
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
	"tokElse",
	"tokLet",
	"tokIn",
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
	"'='",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc.y:266

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 73

var yyAct = [...]int{

	50, 8, 27, 26, 2, 16, 14, 40, 18, 41,
	11, 12, 13, 15, 19, 20, 10, 37, 17, 21,
	17, 45, 23, 33, 31, 36, 51, 19, 32, 52,
	64, 19, 63, 17, 62, 57, 51, 42, 61, 52,
	47, 19, 22, 48, 46, 49, 54, 55, 53, 56,
	43, 39, 57, 58, 28, 38, 44, 1, 60, 25,
	24, 59, 35, 34, 65, 30, 29, 9, 7, 6,
	5, 4, 3,
}
var yyPact = [...]int{

	1, -1000, 26, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	1, -1000, -1000, -1000, 1, -1000, 27, 1, 42, 1,
	12, 16, 42, -1, 47, 37, -1000, -13, -10, 21,
	36, 26, -1000, 50, 5, 30, -1000, -1000, 3, 42,
	1, 24, -1000, 1, 3, 3, 42, -1000, -1000, 26,
	39, -1000, 24, 26, -1000, -1000, -1000, 24, 22, 18,
	39, -1000, -1000, 14, -1000, 39,
}
var yyPgo = [...]int{

	0, 4, 1, 72, 71, 70, 69, 68, 67, 66,
	65, 63, 62, 2, 0, 61, 60, 59, 3, 57,
}
var yyR1 = [...]int{

	0, 19, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 3, 3, 3, 4, 5, 6, 11, 11, 11,
	12, 12, 13, 13, 14, 14, 14, 14, 14, 15,
	15, 7, 9, 9, 9, 10, 10, 8, 16, 16,
	16, 17, 17, 18,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	3, 1, 1, 1, 5, 1, 5, 0, 1, 2,
	1, 3, 1, 3, 1, 3, 3, 3, 4, 1,
	3, 4, 0, 1, 2, 1, 3, 4, 0, 1,
	2, 1, 3, 3,
}
var yyChk = [...]int{

	-1000, -19, -1, -3, -4, -5, -6, -7, -2, -8,
	15, 9, 10, 11, 5, 12, 4, 17, 7, 15,
	-1, -1, 15, -1, -16, -17, -18, -13, 12, -9,
	-10, -1, 16, -2, -11, -12, -13, 18, 8, 14,
	20, 19, 16, 14, 6, 16, 14, -2, -18, -1,
	-14, 12, 15, -1, -2, -2, -13, 13, -14, -15,
	-14, 16, 16, 14, 16, -14,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 11, 12, 13, 0, 15, 0, 0, 38, 32,
	0, 0, 17, 0, 0, 39, 41, 0, 22, 0,
	33, 35, 9, 0, 0, 18, 20, 10, 0, 40,
	0, 0, 31, 34, 0, 0, 19, 37, 42, 43,
	23, 24, 0, 36, 14, 16, 21, 0, 29, 0,
	25, 26, 27, 0, 28, 30,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 3, 14, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 19, 3,
	3, 20, 3, 3, 3, 3, 3, 3, 3, 3,
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
		//line yacc.y:48
		{
			yylex.(*lexer).result = yyDollar[1].ast
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:61
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:67
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:74
		{
			yyVAL.ast = &lambda.Boolean{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string) == "true",
			}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:81
		{
			yyVAL.ast = &lambda.Integer{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(int64),
			}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:88
		{
			yyVAL.ast = &lambda.String{
				Loc:   yyDollar[1].tok.loc,
				Value: yyDollar[1].tok.val.(string),
			}
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:97
		{
			yyVAL.ast = &lambda.If{
				Loc:        extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Condition:  yyDollar[2].ast,
				Consequent: yyDollar[3].ast,
				Alternate:  yyDollar[5].ast,
			}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:107
		{
			yyVAL.ast = &lambda.Variable{
				Loc: yyDollar[1].tok.loc,
				Var: yyDollar[1].tok.val.(string),
			}
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yacc.y:116
		{
			yyVAL.ast = &lambda.Abstraction{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[5].ast.Location()),
				Vars: yyDollar[3].asts,
				Body: yyDollar[5].ast,
			}
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:125
		{
			yyVAL.asts = []lambda.AST{}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:132
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:136
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:142
		{
			yyVAL.ast = &lambda.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: nil,
			}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:150
		{
			yyVAL.ast = &lambda.TypedName{
				Loc:  yyDollar[1].tok.loc,
				Name: yyDollar[1].tok.val.(string),
				Type: yyDollar[3].ast,
			}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:160
		{
			yyVAL.ast = &lambda.TyName{
				Loc:  yyDollar[1].tok.loc,
				Type: yyDollar[1].tok.val.(string),
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:167
		{
			yyVAL.ast = &lambda.TyArrow{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Dom:   yyDollar[1].ast,
				Range: yyDollar[3].ast,
			}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:175
		{
			yyVAL.ast = yyDollar[2].ast
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:179
		{
			yyVAL.ast = &lambda.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:186
		{
			yyVAL.ast = &lambda.TyTuple{
				Loc:  extend(yyDollar[1].tok.loc, yyDollar[3].tok.loc),
				Elts: yyDollar[2].asts,
			}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:195
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:199
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:205
		{
			yyVAL.ast = &lambda.Application{
				Loc:  extend(yyDollar[1].ast.Location(), yyDollar[4].tok.loc),
				Func: yyDollar[1].ast,
				Args: yyDollar[3].asts,
			}
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:214
		{
			yyVAL.asts = []lambda.AST{}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:222
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:226
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yacc.y:231
		{
			yyVAL.ast = &lambda.Let{
				Loc:      extend(yyDollar[1].tok.loc, yyDollar[4].ast.Location()),
				Bindings: yyDollar[2].asts,
				Body:     yyDollar[4].ast,
			}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yacc.y:240
		{
			yyVAL.asts = []lambda.AST{}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yacc.y:248
		{
			yyVAL.asts = []lambda.AST{yyDollar[1].ast}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:252
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:258
		{
			yyVAL.ast = &lambda.NameBinding{
				Loc:   extend(yyDollar[1].ast.Location(), yyDollar[3].ast.Location()),
				Var:   yyDollar[1].ast,
				Value: yyDollar[3].ast,
			}
		}
	}
	goto yystack /* stack new state and value */
}

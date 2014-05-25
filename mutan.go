
//line mutan.y:2
package mutan
import __yyfmt__ "fmt"
//line mutan.y:2
		
import (
	_"fmt"
)

var Tree *SyntaxTree


//line mutan.y:12
type yySymType struct {
	yys int
	num int
	str string
	tnode *SyntaxTree
}

const ASSIGN = 57346
const EQUAL = 57347
const IF = 57348
const ELSE = 57349
const FOR = 57350
const LEFT_BRACES = 57351
const RIGHT_BRACES = 57352
const STORE = 57353
const LEFT_BRACKET = 57354
const RIGHT_BRACKET = 57355
const ASM = 57356
const LEFT_PAR = 57357
const RIGHT_PAR = 57358
const STOP = 57359
const ADDR = 57360
const ORIGIN = 57361
const CALLER = 57362
const CALLVAL = 57363
const CALLDATALOAD = 57364
const CALLDATASIZE = 57365
const GASPRICE = 57366
const DOT = 57367
const THIS = 57368
const ARRAY = 57369
const CALL = 57370
const COMMA = 57371
const SIZEOF = 57372
const QUOTE = 57373
const END_STMT = 57374
const RETURN = 57375
const CREATE = 57376
const TRANSACT = 57377
const NIL = 57378
const BALANCE = 57379
const VAR_ASSIGN = 57380
const LAMBDA = 57381
const COLON = 57382
const DIFFICULTY = 57383
const PREVHASH = 57384
const TIMESTAMP = 57385
const BLOCKNUM = 57386
const COINBASE = 57387
const GAS = 57388
const VAR = 57389
const ID = 57390
const NUMBER = 57391
const INLINE_ASM = 57392
const OP = 57393
const DOP = 57394
const TYPE = 57395
const STR = 57396
const BOOLEAN = 57397
const CODE = 57398

var yyToknames = []string{
	"ASSIGN",
	"EQUAL",
	"IF",
	"ELSE",
	"FOR",
	"LEFT_BRACES",
	"RIGHT_BRACES",
	"STORE",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"ASM",
	"LEFT_PAR",
	"RIGHT_PAR",
	"STOP",
	"ADDR",
	"ORIGIN",
	"CALLER",
	"CALLVAL",
	"CALLDATALOAD",
	"CALLDATASIZE",
	"GASPRICE",
	"DOT",
	"THIS",
	"ARRAY",
	"CALL",
	"COMMA",
	"SIZEOF",
	"QUOTE",
	"END_STMT",
	"RETURN",
	"CREATE",
	"TRANSACT",
	"NIL",
	"BALANCE",
	"VAR_ASSIGN",
	"LAMBDA",
	"COLON",
	"DIFFICULTY",
	"PREVHASH",
	"TIMESTAMP",
	"BLOCKNUM",
	"COINBASE",
	"GAS",
	"VAR",
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"DOP",
	"TYPE",
	"STR",
	"BOOLEAN",
	"CODE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:225



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 44,
	-1, 57,
	51, 41,
	52, 41,
	-2, 48,
}

const yyNprod = 69
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 392

var yyAct = []int{

	21, 4, 11, 2, 20, 59, 48, 36, 35, 60,
	29, 68, 123, 27, 75, 40, 41, 155, 69, 34,
	120, 30, 54, 33, 28, 128, 12, 32, 31, 27,
	47, 168, 162, 150, 10, 148, 102, 170, 58, 61,
	19, 15, 22, 119, 64, 65, 63, 67, 23, 18,
	13, 101, 14, 99, 166, 71, 73, 74, 8, 36,
	35, 29, 36, 35, 94, 93, 46, 95, 97, 62,
	34, 57, 30, 153, 33, 28, 9, 12, 32, 31,
	27, 36, 35, 5, 152, 163, 56, 116, 36, 35,
	149, 19, 15, 22, 96, 151, 141, 154, 140, 23,
	139, 138, 125, 127, 124, 90, 126, 137, 136, 132,
	115, 36, 35, 77, 78, 79, 80, 81, 85, 142,
	135, 145, 36, 35, 146, 147, 42, 134, 36, 35,
	133, 88, 36, 35, 43, 82, 83, 84, 86, 87,
	89, 131, 130, 129, 103, 92, 13, 70, 14, 114,
	165, 158, 113, 157, 8, 98, 112, 29, 160, 161,
	111, 164, 44, 167, 110, 109, 34, 108, 30, 169,
	33, 28, 9, 12, 32, 31, 27, 106, 105, 5,
	104, 53, 13, 52, 14, 51, 156, 19, 15, 22,
	8, 91, 50, 29, 49, 23, 38, 117, 107, 100,
	3, 37, 34, 144, 30, 159, 33, 28, 9, 12,
	32, 31, 27, 39, 122, 5, 55, 66, 13, 45,
	14, 143, 121, 19, 15, 22, 8, 7, 24, 29,
	26, 23, 17, 16, 76, 25, 6, 1, 34, 0,
	30, 0, 33, 28, 9, 12, 32, 31, 27, 0,
	0, 5, 0, 0, 13, 0, 14, 0, 118, 19,
	15, 22, 8, 0, 0, 29, 0, 23, 0, 0,
	0, 0, 0, 0, 34, 0, 30, 0, 33, 28,
	9, 12, 32, 31, 27, 0, 0, 5, 0, 0,
	13, 0, 14, 0, 0, 19, 15, 22, 8, 0,
	0, 29, 0, 23, 0, 0, 0, 0, 0, 0,
	34, 0, 30, 0, 33, 28, 9, 12, 32, 31,
	27, 0, 0, 5, 0, 0, 0, 29, 0, 0,
	0, 19, 15, 22, 0, 0, 34, 0, 30, 23,
	33, 28, 0, 0, 32, 31, 27, 0, 0, 0,
	0, 0, 0, 29, 0, 0, 0, 19, 15, 22,
	0, 0, 34, 0, 30, 23, 33, 28, 0, 0,
	32, 31, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 22, 0, 0, 0, 0,
	0, 23,
}
var yyPact = []int{

	-1000, -1000, 284, -1000, -44, 189, -1000, -1000, 181, -1000,
	-1000, -1000, 284, -7, -7, 122, 215, -1000, -1000, 18,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -48, 179,
	177, 170, 168, 166, -3, -1000, -7, -51, -41, -1000,
	30, 37, -7, -7, 213, -7, -1000, -38, -13, 131,
	336, 336, 336, -34, 94, -1000, -1000, -1000, -44, 178,
	129, -1000, -7, -1000, -44, 81, -7, -44, 142, -1000,
	-1000, 24, 187, 22, 7, 128, -1000, 165, 163, 162,
	186, 152, 150, 149, 145, 141, 137, 134, 95, 72,
	185, -1000, -1000, 248, 11, 212, 210, -44, -36, 336,
	-7, 336, -23, -1000, 127, 126, 125, -7, 114, 111,
	104, 92, 91, 85, 84, 82, 80, -7, 196, -7,
	-1000, -1000, 310, -1000, 6, 77, 4, 79, -1000, -1000,
	-1000, -1000, 71, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 60, -1000, 88, 8, 176, -1000, 336, -1000,
	-23, -1000, -1000, 201, -1000, -1000, -1000, 3, 69, -7,
	140, 44, -23, -1000, -44, -1000, -1000, 2, -23, 21,
	-1000,
}
var yyPgo = []int{

	0, 237, 3, 200, 1, 2, 49, 4, 236, 34,
	235, 234, 233, 232, 232, 232, 230, 228, 227, 221,
	0, 216,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	10, 10, 10, 10, 10, 10, 14, 14, 15, 15,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 8, 19, 19, 18, 18,
	18, 4, 4, 4, 4, 9, 9, 21, 21, 5,
	5, 5, 5, 5, 5, 5, 12, 13, 6, 7,
	7, 7, 7, 7, 7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 4, 1,
	3, 12, 8, 6, 4, 3, 3, 0, 1, 0,
	3, 3, 3, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 6, 6, 4, 0, 9, 7,
	5, 1, 1, 2, 0, 2, 3, 1, 1, 3,
	6, 3, 4, 1, 1, 1, 2, 5, 1, 1,
	1, 4, 1, 1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 14, 32,
	-9, -5, 33, 6, 8, 48, -12, -13, -6, 47,
	-7, -20, 49, 55, -17, -10, -16, 36, 31, 17,
	28, 35, 34, 30, 26, 52, 51, 12, 15, -3,
	-4, -4, 4, 12, 40, 4, 48, 12, 54, 15,
	15, 15, 15, 15, 25, -21, -6, -9, -4, 56,
	50, 9, 32, 9, -4, -4, 4, -4, 49, 31,
	16, -7, 48, -7, -7, 48, -11, 19, 20, 21,
	22, 23, 41, 42, 43, 24, 44, 45, 37, 46,
	11, 13, 16, -2, -4, -2, 13, -4, 13, 29,
	12, 29, 29, 16, 15, 15, 15, 12, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 12, 10, 32,
	9, 10, 4, 48, -7, -4, -7, -20, 48, 16,
	16, 16, -4, 16, 16, 16, 16, 16, 16, 16,
	16, 16, -4, -19, 7, -4, -2, -5, 29, 13,
	29, 16, 13, 13, 9, 9, 10, -7, -20, 4,
	-2, -2, 29, 16, -4, 10, 10, -20, 29, -20,
	16,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 0, 6, 7, 0, 9,
	41, 42, 44, 44, 44, 67, 53, 54, 55, 0,
	58, 59, 60, 62, 63, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 0, 45, 44, 0, 0, 43,
	0, 0, 44, 44, 0, 44, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 46, 47, -2, 0, 0,
	0, 3, 44, 3, 49, 0, 44, 51, 0, 68,
	10, 0, 67, 0, 0, 0, 15, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 5, 8, 44, 0, 44, 61, 52, 0, 0,
	44, 0, 0, 14, 0, 0, 0, 44, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 37, 44,
	3, 40, 0, 57, 0, 0, 0, 0, 67, 20,
	21, 22, 0, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 0, 35, 0, 0, 44, 50, 0, 61,
	0, 13, 23, 33, 3, 3, 39, 0, 0, 44,
	44, 44, 0, 12, 34, 36, 38, 0, 0, 0,
	11,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line mutan.y:30
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:34
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:35
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:39
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:40
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line mutan.y:41
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:42
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:43
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 9:
		//line mutan.y:44
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 10:
		//line mutan.y:49
		{ yyVAL.tnode = NewNode(StopTy) }
	case 11:
		//line mutan.y:52
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 12:
		//line mutan.y:56
		{
		  	  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
	          }
	case 13:
		//line mutan.y:59
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 14:
		//line mutan.y:60
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 15:
		//line mutan.y:61
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 16:
		//line mutan.y:65
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 17:
		//line mutan.y:66
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 18:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 19:
		//line mutan.y:71
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 20:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 21:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 22:
		//line mutan.y:77
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 23:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 24:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 25:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 26:
		//line mutan.y:81
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 27:
		//line mutan.y:82
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 28:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 29:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 30:
		//line mutan.y:85
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 31:
		//line mutan.y:86
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 32:
		//line mutan.y:87
		{ yyVAL.tnode = NewNode(GasTy) }
	case 33:
		//line mutan.y:88
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 34:
		//line mutan.y:90
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 35:
		//line mutan.y:98
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 36:
		//line mutan.y:108
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 37:
		//line mutan.y:111
		{ yyVAL.tnode = nil }
	case 38:
		//line mutan.y:116
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 39:
		//line mutan.y:121
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 40:
		//line mutan.y:126
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 41:
		//line mutan.y:132
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 42:
		//line mutan.y:133
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 43:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 44:
		//line mutan.y:135
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 45:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 46:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 47:
		//line mutan.y:146
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 48:
		//line mutan.y:147
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 49:
		//line mutan.y:152
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 50:
		//line mutan.y:158
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 51:
		//line mutan.y:162
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 52:
		//line mutan.y:168
		{
		  	node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
		  	varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 53:
		//line mutan.y:174
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line mutan.y:175
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 55:
		//line mutan.y:176
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 56:
		//line mutan.y:181
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      //$$.VarType = $1
	  }
	case 57:
		//line mutan.y:190
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 58:
		//line mutan.y:200
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 59:
		//line mutan.y:204
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 60:
		//line mutan.y:205
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 61:
		//line mutan.y:206
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 62:
		//line mutan.y:207
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 63:
		//line mutan.y:208
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 64:
		//line mutan.y:209
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 65:
		//line mutan.y:213
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 66:
		//line mutan.y:214
		{ yyVAL.tnode = NewNode(NilTy) }
	case 67:
		//line mutan.y:218
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 68:
		//line mutan.y:222
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

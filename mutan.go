
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
const DIFFICULTY = 57375
const PREVHASH = 57376
const TIMESTAMP = 57377
const BLOCKNUM = 57378
const COINBASE = 57379
const GAS = 57380
const ID = 57381
const NUMBER = 57382
const INLINE_ASM = 57383
const OP = 57384
const DOP = 57385
const TYPE = 57386
const STR = 57387

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
	"DIFFICULTY",
	"PREVHASH",
	"TIMESTAMP",
	"BLOCKNUM",
	"COINBASE",
	"GAS",
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"DOP",
	"TYPE",
	"STR",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:198



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 37,
}

const yyNprod = 58
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 255

var yyAct = []int{

	20, 2, 19, 5, 40, 10, 30, 29, 46, 54,
	39, 138, 103, 60, 3, 32, 33, 55, 139, 134,
	99, 125, 11, 81, 12, 44, 136, 16, 141, 118,
	8, 117, 96, 25, 45, 116, 129, 38, 50, 51,
	52, 53, 28, 98, 26, 57, 27, 24, 115, 75,
	114, 77, 76, 30, 29, 13, 21, 11, 113, 12,
	18, 135, 16, 127, 112, 8, 49, 126, 25, 30,
	29, 111, 110, 95, 79, 109, 108, 28, 107, 26,
	78, 27, 24, 94, 104, 47, 105, 106, 84, 48,
	13, 21, 30, 29, 74, 18, 30, 29, 56, 30,
	29, 122, 121, 30, 29, 93, 124, 123, 92, 30,
	29, 91, 11, 90, 12, 89, 130, 16, 30, 29,
	8, 80, 88, 25, 87, 86, 85, 43, 131, 42,
	132, 133, 28, 41, 26, 137, 27, 24, 31, 83,
	140, 82, 37, 128, 120, 13, 21, 11, 34, 12,
	18, 100, 16, 102, 101, 8, 35, 36, 25, 62,
	63, 64, 65, 66, 70, 119, 7, 28, 22, 26,
	15, 27, 24, 67, 68, 69, 71, 72, 73, 14,
	13, 21, 11, 61, 12, 18, 97, 16, 23, 11,
	8, 12, 9, 25, 16, 6, 17, 8, 4, 1,
	25, 0, 28, 0, 26, 0, 27, 24, 0, 28,
	0, 26, 0, 27, 24, 13, 21, 0, 16, 0,
	18, 0, 13, 21, 25, 59, 0, 18, 0, 0,
	0, 25, 0, 28, 0, 26, 0, 27, 24, 0,
	28, 0, 26, 0, 27, 24, 13, 21, 0, 0,
	0, 18, 0, 58, 21,
}
var yyPact = []int{

	-18, -1000, 183, -1000, -1000, -36, -1000, -1000, 123, -1000,
	-1000, 207, 207, 144, 153, -1000, 130, -1000, -2, -1000,
	-1000, -1000, -1000, -1000, -41, 118, 114, 112, 0, -1000,
	207, -33, 76, 57, 207, 207, 207, 207, -1000, -31,
	-14, 82, 214, -26, 140, -36, 78, -18, 207, -18,
	-36, 67, -36, 61, 108, -1000, -1000, -6, 129, 127,
	72, -1000, 111, 110, 109, 107, 100, 98, 96, 93,
	90, 68, 58, 17, -1000, 176, 11, 141, 150, 149,
	-27, 214, 207, 207, -1000, 62, 60, 59, 56, 55,
	48, 42, 34, 32, 19, 15, 13, 137, 207, -18,
	-1000, 207, 207, -1000, -8, 54, 50, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	134, 27, 106, -1000, -36, 214, -1000, -1000, -18, -18,
	-1000, -10, 51, 16, -28, -1000, -1000, -11, -1000, -28,
	12, -1000,
}
var yyPgo = []int{

	0, 199, 1, 198, 3, 5, 196, 2, 195, 192,
	188, 183, 179, 170, 170, 170, 0, 168, 166, 165,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 10,
	10, 10, 10, 14, 14, 15, 15, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 8,
	19, 19, 18, 18, 18, 4, 4, 4, 9, 9,
	5, 5, 5, 5, 5, 5, 5, 12, 13, 6,
	7, 7, 7, 7, 7, 7, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 1, 0, 1, 1, 1, 4, 3,
	12, 4, 3, 3, 0, 1, 0, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 6,
	4, 0, 9, 7, 5, 1, 1, 0, 2, 3,
	3, 6, 3, 1, 1, 6, 1, 2, 5, 1,
	1, 1, 4, 4, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 32, -3, -4, -8, -18, 14, -9,
	-5, 6, 8, 39, -12, -13, 11, -6, 44, -7,
	-16, 40, -17, -10, 31, 17, 28, 30, 26, 43,
	42, 15, -4, -4, 4, 12, 4, 12, 39, 12,
	45, 15, 15, 15, 25, -4, 41, 9, 32, 9,
	-4, -4, -4, -4, 40, 31, 16, -7, 39, 11,
	39, -11, 19, 20, 21, 22, 23, 33, 34, 35,
	24, 36, 37, 38, 16, -2, -4, -2, 13, 13,
	13, 29, 12, 12, 16, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 10, 32, 9,
	10, 4, 4, 39, -7, -4, -4, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, -19,
	7, -4, -2, -5, -4, 29, 13, 13, 9, 9,
	10, -7, -2, -2, 29, 10, 10, -16, 39, 29,
	-16, 16,
}
var yyDef = []int{

	4, -2, -2, 3, 2, 5, 6, 7, 0, 35,
	36, 37, 37, 56, 43, 44, 0, 46, 0, 49,
	50, 51, 54, 55, 0, 0, 0, 0, 0, 38,
	37, 0, 0, 0, 37, 37, 37, 37, 47, 0,
	0, 0, 0, 0, 0, 39, 0, 4, 37, 4,
	40, 0, 42, 0, 0, 57, 9, 0, 56, 0,
	0, 12, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 8, 37, 0, 37, 52, 53,
	0, 0, 37, 37, 11, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 31, 37, 4,
	34, 0, 37, 48, 0, 0, 0, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	0, 0, 37, 41, 45, 0, 52, 53, 4, 4,
	33, 0, 37, 37, 0, 30, 32, 0, 56, 0,
	0, 10,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
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
		//line mutan.y:36
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 5:
		//line mutan.y:40
		{ yyVAL.tnode = yyS[yypt-0].tnode }
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
		//line mutan.y:48
		{ yyVAL.tnode = NewNode(StopTy) }
	case 10:
		//line mutan.y:51
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 11:
		//line mutan.y:54
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 12:
		//line mutan.y:55
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 13:
		//line mutan.y:59
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 14:
		//line mutan.y:60
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 15:
		//line mutan.y:64
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 16:
		//line mutan.y:65
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 17:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 18:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 19:
		//line mutan.y:71
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 20:
		//line mutan.y:72
		{ yyVAL.tnode = NewNode(CallDataLoadTy) }
	case 21:
		//line mutan.y:73
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 22:
		//line mutan.y:74
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 23:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 24:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 25:
		//line mutan.y:77
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 26:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 27:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 28:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(GasTy) }
	case 29:
		//line mutan.y:85
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 30:
		//line mutan.y:95
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 31:
		//line mutan.y:98
		{ yyVAL.tnode = nil }
	case 32:
		//line mutan.y:103
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 33:
		//line mutan.y:108
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 34:
		//line mutan.y:113
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 35:
		//line mutan.y:119
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 36:
		//line mutan.y:120
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 37:
		//line mutan.y:121
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 38:
		//line mutan.y:125
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 39:
		//line mutan.y:126
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 40:
		//line mutan.y:132
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 41:
		//line mutan.y:138
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 42:
		//line mutan.y:142
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 43:
		//line mutan.y:147
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 44:
		//line mutan.y:148
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 45:
		//line mutan.y:150
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 46:
		//line mutan.y:154
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 47:
		//line mutan.y:159
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 48:
		//line mutan.y:168
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 49:
		//line mutan.y:178
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 50:
		//line mutan.y:182
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 51:
		//line mutan.y:183
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 52:
		//line mutan.y:184
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 53:
		//line mutan.y:185
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 54:
		//line mutan.y:186
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 55:
		//line mutan.y:187
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 56:
		//line mutan.y:191
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 57:
		//line mutan.y:195
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

package settings

type LimitMinMax struct {
	Min int
	Max int
}

type LtLRule struct {
	Radius int
	S      LimitMinMax // S: Survive
	B      LimitMinMax // B: Born
}

var LtLRuleMap = map[string]LtLRule{
	"Life": {
		Radius: 1,
		S: LimitMinMax{
			Min: 2,
			Max: 3,
		},
		B: LimitMinMax{
			Min: 3,
			Max: 3,
		},
	},
	"Gnarl": {
		Radius: 1,
		S: LimitMinMax{
			Min: 0,
			Max: 0,
		},
		B: LimitMinMax{
			Min: 1,
			Max: 1,
		},
	},
	"Majority": {
		Radius: 4,
		S: LimitMinMax{
			Min: 40,
			Max: 80,
		},
		B: LimitMinMax{
			Min: 41,
			Max: 81,
		},
	},
	"Bosco's Rule": {
		Radius: 5,
		S: LimitMinMax{
			Min: 33,
			Max: 57,
		},
		B: LimitMinMax{
			Min: 34,
			Max: 45,
		},
	},
	"Waffle": {
		Radius: 7,
		S: LimitMinMax{
			Min: 99,
			Max: 199,
		},
		B: LimitMinMax{
			Min: 75,
			Max: 170,
		},
	},
	"Globe": {
		Radius: 8,
		S: LimitMinMax{
			Min: 163,
			Max: 223,
		},
		B: LimitMinMax{
			Min: 74,
			Max: 252,
		},
	},
	"BugsMovie": {
		Radius: 10,
		S: LimitMinMax{
			Min: 122,
			Max: 211,
		},
		B: LimitMinMax{
			Min: 123,
			Max: 170,
		},
	},
	// todo: fix this (color: 255)
	"ModernArt": {
		Radius: 10,
		S: LimitMinMax{
			Min: 1,
			Max: 2,
		},
		B: LimitMinMax{
			Min: 3,
			Max: 3,
		},
	},
	//
	"Hash": {
		Radius: 2,
		S: LimitMinMax{
			Min: 4,
			Max: 6,
		},
		B: LimitMinMax{
			Min: 5,
			Max: 6,
		},
	},
}

func GetLtLRule(ruleName string) LtLRule {
	return LtLRuleMap[ruleName]
}

func GetLtLRuleNames() []string {
	ruleNames := make([]string, 0, len(LtLRuleMap))

	for k := range LtLRuleMap {
		ruleNames = append(ruleNames, k)
	}

	// reverse the slice
	//for i := len(ruleNames)/2 - 1; i >= 0; i-- {
	//	opp := len(ruleNames) - 1 - i
	//	ruleNames[i], ruleNames[opp] = ruleNames[opp], ruleNames[i]
	//}

	return ruleNames
}

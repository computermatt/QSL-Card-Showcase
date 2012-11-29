package main

type bounds struct {
	upper float64
	lower float64
}

var (
	bandPlan = map[string]bounds{
		"160":  bounds{2, 1.8},
		"80":   bounds{4, 3.5},
		"60":   bounds{5.4, 5.3},
		"40":   bounds{7.3, 7},
		"30":   bounds{10.15, 10.1},
		"20":   bounds{14.35, 14},
		"17":   bounds{18.168, 18.068},
		"15":   bounds{21.450, 21},
		"12":   bounds{24.990, 24.890},
		"10":   bounds{29.7, 28},
		"6":    bounds{54, 50},
		"2":    bounds{148, 144},
		"1.25": bounds{225, 222},
		"70":   bounds{450, 420},
		"33":   bounds{928, 902},
	}
)

package rivergo

type Trend int

const (
	TrendUnknown Trend = iota
	TrendUp
	TrendDown
	TrendUnchange
	TrendMissing
)

func (t Trend) String() string {
	switch t {
	case TrendUp:
		return "waterLevelUp"
	case TrendDown:
		return "waterLevelDown"
	case TrendUnchange:
		return "waterLevelUnchange"
	case TrendMissing:
		return "waterLevelMissing"
	default:
		return "unknown"
	}
}

func (t Trend) Description() string {
	switch t {
	case TrendUp:
		return "上昇"
	case TrendDown:
		return "下降"
	case TrendUnchange:
		return "変化なし"
	case TrendMissing:
		return "欠測"
	default:
		return "不明"
	}
}

type Condition int

const (
	ConditionUnknown Condition = iota
	ConditionMissing
	ConditionNone
	ConditionNormal
	ConditionFloodStandBy
	ConditionFloodAdvisory
	ConditionFloodEvacuation
	ConditionFloodHazard
)

func (c Condition) String() string {
	switch c {
	case ConditionMissing:
		return "levelMissing"
	case ConditionNone:
		return "levelNone"
	case ConditionNormal:
		return "levelNormal"
	case ConditionFloodStandBy:
		return "floodStandBy"
	case ConditionFloodAdvisory:
		return "floodAdvisory"
	case ConditionFloodEvacuation:
		return "floodEvacuation"
	case ConditionFloodHazard:
		return "floodHazard"
	default:
		return "unknown"
	}
}

func (c Condition) Description() string {
	switch c {
	case ConditionMissing:
		return "欠測"
	case ConditionNone:
		return "基準未設定"
	case ConditionNormal:
		return "平常"
	case ConditionFloodStandBy:
		return "水防団待機水位"
	case ConditionFloodAdvisory:
		return "氾濫注意水位"
	case ConditionFloodEvacuation:
		return "避難判断水位"
	case ConditionFloodHazard:
		return "氾濫危険水位"
	default:
		return "不明"
	}
}

type Station struct {
	ID        string
	Name      string
	Trend     Trend
	Condition Condition
}

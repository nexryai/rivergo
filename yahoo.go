package rivergo

import (
	"strings"

	"github.com/antchfx/htmlquery"
)

func parseWaterLevelTable(htmlStr string) ([]Station, error) {
	doc, err := htmlquery.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	conditionMap := map[string]Condition{
		"levelMissing":    ConditionMissing,
		"levelNone":       ConditionNone,
		"levelNormal":     ConditionNormal,
		"floodStandBy":    ConditionFloodStandBy,
		"floodAdvisory":   ConditionFloodAdvisory,
		"floodEvacuation": ConditionFloodEvacuation,
		"floodHazard":     ConditionFloodHazard,
	}

	nodes, err := htmlquery.QueryAll(doc, "//table[@id='waterLevelStationList']//tr[contains(@class, 'largeLine')]")
	if err != nil {
		return nil, err
	}

	var stations []Station
	for _, node := range nodes {
		var s Station

		if nameNode := htmlquery.FindOne(node, ".//span[@class='name']"); nameNode != nil {
			s.Name = htmlquery.InnerText(nameNode)
		}

		if linkNode := htmlquery.FindOne(node, ".//a[@class='stationLink']"); linkNode != nil {
			s.ID = htmlquery.SelectAttr(linkNode, "data-obsrvtn")
		}

		if imgNode := htmlquery.FindOne(node, ".//span[@class='trendIcon']/img"); imgNode != nil {
			src := htmlquery.SelectAttr(imgNode, "src")
			switch {
			case strings.Contains(src, "waterLevelUp"):
				s.Trend = TrendUp
			case strings.Contains(src, "waterLevelDown"):
				s.Trend = TrendDown
			case strings.Contains(src, "waterLevelUnchange"):
				s.Trend = TrendUnchange
			case strings.Contains(src, "waterLevelMissing"):
				s.Trend = TrendMissing
			default:
				s.Trend = TrendUnknown
			}
		}

		classAttr := htmlquery.SelectAttr(node, "class")
		for _, class := range strings.Fields(classAttr) {
			if c, exists := conditionMap[class]; exists {
				s.Condition = c
				break
			}
		}

		stations = append(stations, s)
	}

	return stations, nil
}

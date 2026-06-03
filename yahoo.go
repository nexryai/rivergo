package rivergo

import (
	"strings"

	"github.com/antchfx/htmlquery"
)

type Station struct {
	ID        string
	Name      string
	Trend     string
	Condition string
}

func parseWaterLevelTable(htmlStr string) ([]Station, error) {
	doc, err := htmlquery.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	conditionMap := map[string]string{
		"levelMissing":    "欠測",
		"levelNone":       "基準未設定",
		"levelNormal":     "平常",
		"floodStandBy":    "水防団待機水位",
		"floodAdvisory":   "氾濫注意水位",
		"floodEvacuation": "避難判断水位",
		"floodHazard":     "氾濫危険水位",
	}

	nodes, err := htmlquery.QueryAll(doc, "//table[@id='waterLevelStationList']//tr[contains(@class, 'largeLine')]")
	if err != nil {
		return nil, err
	}

	var stations []Station
	for _, node := range nodes {
		var s Station

		// 観測所名の取得
		if nameNode := htmlquery.FindOne(node, ".//span[@class='name']"); nameNode != nil {
			s.Name = htmlquery.InnerText(nameNode)
		}

		// 観測所IDの取得
		if linkNode := htmlquery.FindOne(node, ".//a[@class='stationLink']"); linkNode != nil {
			s.ID = htmlquery.SelectAttr(linkNode, "data-obsrvtn")
		}

		// trendIconのsrcから水位の傾向を判断
		if imgNode := htmlquery.FindOne(node, ".//span[@class='trendIcon']/img"); imgNode != nil {
			src := htmlquery.SelectAttr(imgNode, "src")
			switch {
			case strings.Contains(src, "waterLevelUp"):
				s.Trend = "上昇"
			case strings.Contains(src, "waterLevelDown"):
				s.Trend = "下降"
			case strings.Contains(src, "waterLevelMissing"):
				s.Trend = "欠測"
			default:
				s.Trend = "不明"
			}
		}

		// クラス名から水位状況を判断
		classAttr := htmlquery.SelectAttr(node, "class")
		for _, class := range strings.Fields(classAttr) {
			if cond, exists := conditionMap[class]; exists {
				s.Condition = cond
				break
			}
		}

		stations = append(stations, s)
	}

	return stations, nil
}

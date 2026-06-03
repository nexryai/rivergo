package rivergo

import (
	"testing"
)

func TestParseWaterLevelTable(t *testing.T) {
	tests := []struct {
		name     string
		htmlData string
		expected []Station
	}{
		{
			name: "パターン1：氾濫注意",
			htmlData: `<table class="waterLevelStationList target_modules" id="waterLevelStationList">
      <tbody>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">間黒</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0204900400078">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">加賀田</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0204900400079">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelMissing">
        <td class="stationName"><span class="name">川根橋</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0204900400106">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelMissing_l.png"></span>
              <span class="waterLevelLabel">欠測</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine floodAdvisory">
        <td class="stationName"><span class="name">高橋</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0204900400001">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_black_l_02.png"></span>
              <span class="waterLevelLabel">氾濫注意水位</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNone">
        <td class="stationName"><span class="name">下石崎</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="2127100400023">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUp_black_l_02.png"></span>
              <span class="waterLevelLabel">上昇中</span>
            </div>
          </a>
        </td>
      </tr>
      </tbody>
    </table>`,
			expected: []Station{
				{ID: "0204900400078", Name: "間黒", Trend: TrendDown, Condition: ConditionNormal},
				{ID: "0204900400079", Name: "加賀田", Trend: TrendDown, Condition: ConditionNormal},
				{ID: "0204900400106", Name: "川根橋", Trend: TrendMissing, Condition: ConditionMissing},
				{ID: "0204900400001", Name: "高橋", Trend: TrendDown, Condition: ConditionFloodAdvisory},
				{ID: "2127100400023", Name: "下石崎", Trend: TrendUp, Condition: ConditionNone},
			},
		},
		{
			name: "パターン2：変化なし（waterLevelUnchange）を含む",
			htmlData: `<table class="waterLevelStationList target_modules" id="waterLevelStationList">
      <tbody>
      <tr class="largeLine levelMissing">
        <td class="stationName"><span class="name">桂川城南橋</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0486500400042">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelMissing_l.png"></span>
              <span class="waterLevelLabel">欠測</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">桂川強瀬</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0486500400049">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNone">
        <td class="stationName"><span class="name">大月</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0358500400001">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUnchange_black_l.png"></span>
              <span class="waterLevelLabel">変化なし</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">上依知</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0358500400002">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUnchange_white_l.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">磯部</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0358500400003">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUnchange_white_l.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">相模大橋</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="2132000400033">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">厚木</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0358500400004">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUnchange_white_l.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">神川橋</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="2132000400032">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png"></span>
              <span class="waterLevelLabel">平常</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNone">
        <td class="stationName"><span class="name">神川橋（下）</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="2132000400212">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_black_l_02.png"></span>
              <span class="waterLevelLabel">下降中</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="largeLine levelNone">
        <td class="stationName"><span class="name">馬入橋</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="2132000400034">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUp_black_l_02.png"></span>
              <span class="waterLevelLabel">上昇中</span>
            </div>
          </a>
        </td>
      </tr>
      </tbody>
    </table>`,
			expected: []Station{
				{ID: "0486500400042", Name: "桂川城南橋", Trend: TrendMissing, Condition: ConditionMissing},
				{ID: "0486500400049", Name: "桂川強瀬", Trend: TrendDown, Condition: ConditionNormal},
				{ID: "0358500400001", Name: "大月", Trend: TrendUnchange, Condition: ConditionNone},
				{ID: "0358500400002", Name: "上依知", Trend: TrendUnchange, Condition: ConditionNormal},
				{ID: "0358500400003", Name: "磯部", Trend: TrendUnchange, Condition: ConditionNormal},
				{ID: "2132000400033", Name: "相模大橋", Trend: TrendDown, Condition: ConditionNormal},
				{ID: "0358500400004", Name: "厚木", Trend: TrendUnchange, Condition: ConditionNormal},
				{ID: "2132000400032", Name: "神川橋", Trend: TrendDown, Condition: ConditionNormal},
				{ID: "2132000400212", Name: "神川橋（下）", Trend: TrendDown, Condition: ConditionNone},
				{ID: "2132000400034", Name: "馬入橋", Trend: TrendUp, Condition: ConditionNone},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stations, err := parseWaterLevelTable(tt.htmlData)
			if err != nil {
				t.Fatalf("ParseWaterLevelTable() failed: %v", err)
			}

			if len(stations) != len(tt.expected) {
				t.Fatalf("got %d stations, want %d", len(stations), len(tt.expected))
			}

			for i, want := range tt.expected {
				got := stations[i]
				if got != want {
					t.Errorf("stations[%d] = %+v; want %+v", i, got, want)
				} else {
					t.Logf("stations[%d] passed: ID=%s, Name=%s, Trend=%v, Condition=%v", i, got.ID, got.Name, got.Trend, got.Condition)
				}
			}
		})
	}
}

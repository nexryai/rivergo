package rivergo

import (
	"testing"
)

func TestParseWaterLevelTable(t *testing.T) {
	htmlData := `<table class="waterLevelStationList target_modules" id="waterLevelStationList">
      <tbody>
      <tr class="waterLevelStationListHeader">
        <th class="stationNameHeader"><span>観測所名</span></th>
        <th class="waterLevelHeader"><span>水位状況</span></th>
      </tr>
      <tr class="riverUpper">
        <td colspan="2">
          <span class="publishedDate">2026/6/3 20:30更新</span>
          <span class="riverUpperLabel">上流</span>
        </td>
      </tr>
      <tr class="largeLine levelNormal">
        <td class="stationName"><span class="name">間黒</span></td>
        <td>
          <a class="stationLink" href="javascript:void(0);" data-obsrvtn="0204900400078">
            <div class="waterLevel">
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png" width="20" height="20"></span>
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
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_white_l_02.png" width="20" height="20"></span>
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
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelMissing_l.png" width="20" height="20"></span>
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
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelDown_black_l_02.png" width="20" height="20"></span>
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
              <span class="trendIcon"><img src="//s.yimg.jp/images/weather/smp/v2/img/river/detail_stationList_ico_waterLevelUp_black_l_02.png" width="20" height="20"></span>
              <span class="waterLevelLabel">上昇中</span>
            </div>
          </a>
        </td>
      </tr>
      <tr class="riverDown">
        <td colspan="2">下流</td>
      </tr>
      </tbody>
    </table>`

	expected := []Station{
		{ID: "0204900400078", Name: "間黒", Trend: "下降", Condition: "平常"},
		{ID: "0204900400079", Name: "加賀田", Trend: "下降", Condition: "平常"},
		{ID: "0204900400106", Name: "川根橋", Trend: "欠測", Condition: "欠測"},
		{ID: "0204900400001", Name: "高橋", Trend: "下降", Condition: "氾濫注意水位"},
		{ID: "2127100400023", Name: "下石崎", Trend: "上昇", Condition: "基準未設定"},
	}

	stations, err := parseWaterLevelTable(htmlData)
	if err != nil {
		t.Fatalf("ParseWaterLevelTable() failed: %v", err)
	}

	if len(stations) != len(expected) {
		t.Fatalf("got %d stations, want %d", len(stations), len(expected))
	}

	for i, want := range expected {
		got := stations[i]
		if got != want {
			t.Errorf("stations[%d] = %+v; want %+v", i, got, want)
		} else {
			t.Logf("stations[%d] passed: ID=%s, Name=%s, Trend=%s, Condition=%s", i, got.ID, got.Name, got.Trend, got.Condition)
		}
	}
}

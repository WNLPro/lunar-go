// @Title HolidayUtil
// @Description 法定节假日工具（自2001年12月29日起）
// @Author 6tail
package HolidayUtil

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/calendar"
	"strings"
)

const size = 18
const zero rune = '0'
const data = "200112290020020101200112300020020101200201010120020101200201020120020101200201030120020101200202091020020212200202101020020212200202121120020212200202131120020212200202141120020212200202151120020212200202161120020212200202171120020212200202181120020212200204273020020501200204283020020501200205013120020501200205023120020501200205033120020501200205043120020501200205053120020501200205063120020501200205073120021001200209286020021001200209296020021001200210016120021001200210026120021001200210036120021001200210046120021001200210056120021001200210066120021001200210076120021001200301010120030101200302011120030201200302021120030201200302031120030201200302041120030201200302051120030201200302061120030201200302071120030201200302081020030201200302091020030201200304263020030501200304273020030501200305013120030501200305023120030501200305033120030501200305043120030501200305053120030501200305063120030501200305073120031001200309276020031001200309286020031001200310016120031001200310026120031001200310036120031001200310046120031001200310056120031001200310066120031001200310076120031001200401010120040101200401171020040122200401181020040122200401221120040122200401231120040122200401241120040122200401251120040122200401261120040122200401271120040122200401281120040122200405013120040501200405023120040501200405033120040501200405043120040501200405053120040501200405063120040501200405073120041001200405083020040501200405093020040501200410016120041001200410026120041001200410036120041001200410046120041001200410056120041001200410066120041001200410076120041001200410096020041001200410106020041001200501010120050101200501020120050101200501030120050101200502051020050209200502061020050209200502091120050209200502101120050209200502111120050209200502121120050209200502131120050209200502141120050209200502151120050209200504303020050501200505013120050501200505023120050501200505033120050501200505043120050501200505053120050501200505063120050501200505073120051001200505083020050501200510016120051001200510026120051001200510036120051001200510046120051001200510056120051001200510066120051001200510076120051001200510086020051001200510096020051001200512310020060101200601010120060101200601020120060101200601030120060101200601281020060129200601291120060129200601301120060129200601311120060129200602011120060129200602021120060129200602031120060129200602041120060129200602051020060129200604293020060501200604303020060501200605013120060501200605023120060501200605033120060501200605043120060501200605053120060501200605063120060501200605073120061001200609306020061001200610016120061001200610026120061001200610036120061001200610046120061001200610056120061001200610066120061001200610076120061001200610086020061001200612300020070101200612310020070101200701010120070101200701020120070101200701030120070101200702171020070218200702181120070218200702191120070218200702201120070218200702211120070218200702221120070218200702231120070218200702241120070218200702251020070218200704283020070501200704293020070501200705013120070501200705023120070501200705033120070501200705043120070501200705053120070501200705063120070501200705073120070501200709296020071001200709306020071001200710016120071001200710026120071001200710036120071001200710046120071001200710056120071001200710066120071001200710076120071001200712290020080101200712300120080101200712310120080101200801010120080101200802021020080206200802031020080206200802061120080206200802071120080206200802081120080206200802091120080206200802101120080206200802111120080206200802121120080206200804042120080404200804052120080404200804062120080404200805013120080501200805023120080501200805033120080501200805043020080501200806074120080608200806084120080608200806094120080608200809135120080914200809145120080914200809155120080914200809276020081001200809286020081001200809296120081001200809306120081001200810016120081001200810026120081001200810036120081001200810046120081001200810056120081001200901010120090101200901020120090101200901030120090101200901040020090101200901241020090125200901251120090125200901261120090125200901271120090125200901281120090125200901291120090125200901301120090125200901311120090125200902011020090125200904042120090404200904052120090404200904062120090404200905013120090501200905023120090501200905033120090501200905284120090528200905294120090528200905304120090528200905314020090528200909276020091001200910016120091001200910026120091001200910036120091001200910046120091001200910055120091003200910065120091003200910075120091003200910085120091003200910105020091003201001010120100101201001020120100101201001030120100101201002131120100213201002141120100213201002151120100213201002161120100213201002171120100213201002181120100213201002191120100213201002201020100213201002211020100213201004032120100405201004042120100405201004052120100405201005013120100501201005023120100501201005033120100501201006124020100616201006134020100616201006144120100616201006154120100616201006164120100616201009195020100922201009225120100922201009235120100922201009245120100922201009255020100922201009266020101001201010016120101001201010026120101001201010036120101001201010046120101001201010056120101001201010066120101001201010076120101001201010096020101001201101010120110101201101020120110101201101030120110101201101301020110203201102021120110203201102031120110203201102041120110203201102051120110203201102061120110203201102071120110203201102081120110203201102121020110203201104022020110405201104032120110405201104042120110405201104052120110405201104303120110501201105013120110501201105023120110501201106044120110606201106054120110606201106064120110606201109105120110912201109115120110912201109125120110912201110016120111001201110026120111001201110036120111001201110046120111001201110056120111001201110066120111001201110076120111001201110086020111001201110096020111001201112310020120101201201010120120101201201020120120101201201030120120101201201211020120123201201221120120123201201231120120123201201241120120123201201251120120123201201261120120123201201271120120123201201281120120123201201291020120123201203312020120404201204012020120404201204022120120404201204032120120404201204042120120404201204283020120501201204293120120501201204303120120501201205013120120501201205023020120501201206224120120623201206234120120623201206244120120623201209295020120930201209305120120930201210016120121001201210026120121001201210036120121001201210046120121001201210056120121001201210066120121001201210076120121001201210086020121001201301010120130101201301020120130101201301030120130101201301050020130101201301060020130101201302091120130210201302101120130210201302111120130210201302121120130210201302131120130210201302141120130210201302151120130210201302161020130210201302171020130210201304042120130404201304052120130404201304062120130404201304273020130501201304283020130501201304293120130501201304303120130501201305013120130501201306084020130612201306094020130612201306104120130612201306114120130612201306124120130612201309195120130919201309205120130919201309215120130919201309225020130919201309296020131001201310016120131001201310026120131001201310036120131001201310046120131001201310056120131001201310066120131001201310076120131001201401010120140101201401261020140131201401311120140131201402011120140131201402021120140131201402031120140131201402041120140131201402051120140131201402061120140131201402081020140131201404052120140405201404062120140405201404072120140405201405013120140501201405023120140501201405033120140501201405043020140501201405314120140602201406014120140602201406024120140602201409065120140908201409075120140908201409085120140908201409286020141001201410016120141001201410026120141001201410036120141001201410046120141004201410056120141001201410066120141001201410076120141001201410116020141001201501010120150101201501020120150101201501030120150101201501040020150101201502151020150219201502181120150219201502191120150219201502201120150219201502211120150219201502221120150219201502231120150219201502241120150219201502281020150219201504042120150405201504052120150405201504062120150405201505013120150501201505023120150501201505033120150501201506204120150620201506214120150620201506224120150620201509038120150903201509048120150903201509058120150903201509068020150903201509265120150927201509275120150927201510016120151001201510026120151001201510036120151001201510046120151004201510056120151001201510066120151001201510076120151001201510106020151001201601010120160101201601020120160101201601030120160101201602061020160208201602071120160208201602081120160208201602091120160208201602101120160208201602111120160208201602121120160208201602131120160208201602141020160208201604022120160404201604032120160404201604042120160404201604303120160501201605013120160501201605023120160501201606094120160609201606104120160609201606114120160609201606124020160609201609155120160915201609165120160915201609175120160915201609185020160915201610016120161001201610026120161001201610036120161001201610046120161001201610056120161001201610066120161001201610076120161001201610086020161001201610096020161001201612310120170101201701010120170101201701020120170101201701221020170128201701271120170128201701281120170128201701291120170128201701301120170128201701311120170128201702011120170128201702021120170128201702041020170128201704012020170404201704022120170404201704032120170404201704042120170404201704293120170501201704303120170501201705013120170501201705274020170530201705284120170530201705294120170530201705304120170530201709306020171001201710016120171001201710026120171001201710036120171001201710045120171004201710056120171001201710066120171001201710076120171001201710086120171001201712300120180101201712310120180101201801010120180101201802111020180216201802151120180216201802161120180216201802171120180216201802181120180216201802191120180216201802201120180216201802211120180216201802241020180216201804052120180405201804062120180405201804072120180405201804082020180405201804283020180501201804293120180501201804303120180501201805013120180501201806164120180618201806174120180618201806184120180618201809225120180924201809235120180924201809245120180924201809296020181001201809306020181001201810016120181001201810026120181001201810036120181001201810046120181001201810056120181001201810066120181001201810076120181001201812290020190101201812300120190101201812310120190101201901010120190101201902021020190205201902031020190205201902041120190205201902051120190205201902061120190205201902071120190205201902081120190205201902091120190205201902101120190205201904052120190405201904062120190405201904072120190405201904283020190501201905013120190501201905023120190501201905033120190501201905043120190501201905053020190501201906074120190607201906084120190607201906094120190607201909135120190913201909145120190913201909155120190913201909296020191001201910016120191001201910026120191001201910036120191001201910046120191001201910056120191001201910066120191001201910076120191001201910126020191001202001010120200101202001191020200125202001241120200125202001251120200125202001261120200125202001271120200125202001281120200125202001291120200125202001301120200125202001311120200125202002011120200125202002021120200125202004042120200404202004052120200404202004062120200404202004263020200501202005013120200501202005023120200501202005033120200501202005043120200501202005053120200501202005093020200501202006254120200625202006264120200625202006274120200625202006284020200625202009277020201001202010017120201001202010026120201001202010036120201001202010046120201001202010056120201001202010066120201001202010076120201001202010086120201001202010106020201001202101010120210101202101020120210101202101030120210101202102071020210212202102111120210212202102121120210212202102131120210212202102141120210212202102151120210212202102161120210212202102171120210212202102201020210212202104032120210404202104042120210404202104052120210404202104253020210501202105013120210501202105023120210501202105033120210501202105043120210501202105053120210501202105083020210501202106124120210614202106134120210614202106144120210614202109185020210921202109195120210921202109205120210921202109215120210921202109266020211001202110016120211001202110026120211001202110036120211001202110046120211001202110056120211001202110066120211001202110076120211001202110096020211001"

var NAMES = []string{"元旦节", "春节", "清明节", "劳动节", "端午节", "中秋节", "国庆节", "国庆中秋", "抗战胜利日"}

var namesInUse = NAMES
var dataInUse = data

func buildHolidayForward(s string) *calendar.Holiday {
	day := s[0:8]
	name := namesInUse[[]rune(s[8:9])[0]-zero]
	work := []rune(s[9:10])[0] == zero
	target := s[10:size]
	return calendar.NewHoliday(day, name, work, target)
}

func buildHolidayBackward(s string) *calendar.Holiday {
	length := len(s)
	day := s[length-18 : length-10]
	name := namesInUse[[]rune(s[length-10 : length-9])[0]-zero]
	work := []rune(s[length-9 : length-8])[0] == zero
	target := s[length-8:]
	return calendar.NewHoliday(day, name, work, target)
}

func findForward(key string) string {
	start := strings.Index(dataInUse, key)
	if start < 0 {
		return ""
	}
	right := dataInUse[start:]
	n := len(right) % size
	if n > 0 {
		right = right[n:]
	}
	for {
		if len(right) < size {
			break
		}
		if strings.HasPrefix(right, key) {
			break
		}
		right = right[size:]
	}
	return right
}

func findBackward(key string) string {
	start := strings.LastIndex(dataInUse, key)
	if start < 0 {
		return ""
	}
	left := dataInUse[:start+len(key)]
	length := len(left)
	n := length % size
	if n > 0 {
		left = left[:length-n]
	}
	length = len(left)
	for {
		if length < size {
			break
		}
		if strings.HasSuffix(left, key) {
			break
		}
		left = left[:length-size]
		length = len(left)
	}
	return left
}

func findHolidaysForward(key string) *list.List {
	l := list.New()
	s := findForward(key)
	if "" == s {
		return l
	}
	for {
		if !strings.HasPrefix(s, key) {
			break
		}
		l.PushBack(buildHolidayForward(s))
		s = s[size:]
	}
	return l
}

func findHolidaysBackward(key string) *list.List {
	l := list.New()
	s := findBackward(key)
	if "" == s {
		return l
	}
	for {
		if !strings.HasSuffix(s, key) {
			break
		}
		l.PushFront(buildHolidayBackward(s))
		s = s[0 : len(s)-size]
	}
	return l
}

func GetHoliday(ymd string) *calendar.Holiday {
	l := findHolidaysForward(strings.Replace(ymd, "-", "", -1))
	if l.Len() < 1 {
		return nil
	}
	return l.Front().Value.(*calendar.Holiday)
}

func GetHolidayByYmd(year int, month int, day int) *calendar.Holiday {
	return GetHoliday(fmt.Sprintf("%d%02d%02d", year, month, day))
}

func GetHolidaysByYm(year int, month int) *list.List {
	return findHolidaysForward(fmt.Sprintf("%d%02d", year, month))
}

func GetHolidaysByYear(year int) *list.List {
	return findHolidaysForward(fmt.Sprintf("%d", year))
}

func GetHolidays(ymd string) *list.List {
	return findHolidaysForward(strings.Replace(ymd, "-", "", -1))
}

func GetHolidaysByTargetYmd(year int, month int, day int) *list.List {
	return findHolidaysBackward(fmt.Sprintf("%d%02d%02d", year, month, day))
}

func GetHolidaysByTarget(ymd string) *list.List {
	return findHolidaysBackward(strings.Replace(ymd, "-", "", -1))
}

func Fix(nms []string, dt string) {
	if nil != nms {
		namesInUse = nms
	}
	if "" == dt {
		return
	}
	appends := ""
	for {
		if len(dt) < size {
			break
		}
		segment := dt[:size]
		day := segment[:8]
		holiday := GetHoliday(day)
		if nil == holiday {
			appends += segment
		} else {
			nameIndex := -1
			for i := 0; i < len(namesInUse); i++ {
				if strings.Compare(namesInUse[i], holiday.GetName()) == 0 {
					nameIndex = i
					break
				}
			}
			if nameIndex > -1 {
				old := day + string(rune(nameIndex+int(zero)))
				if holiday.IsWork() {
					old += "0"
				} else {
					old += "1"
				}
				old += strings.Replace(holiday.GetTarget(), "-", "", -1)
				dataInUse = strings.Replace(dataInUse, old, segment, -1)
			}
		}
		dt = dt[size:]
	}
	if len(appends) > 0 {
		dataInUse += appends
	}
}

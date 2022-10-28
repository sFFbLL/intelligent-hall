package settings

// ---------------------------HTTP数据---------------------------
var Token string
var OperateStatus = map[string]string{"pre": "pre", "next": "next",
	"start": "start", "resume": "resume", "stop": "stop",
	"back": "back", "forward": "forward", "volumn": "volumn"}

// 登录
var LoginData = map[string]string{"xxx1": "GetToken",
	"xxx2": "default",
	"xxx3": "root",
	"xxx4": "e10adc3949ba59abbe56e057f20f883e"}

// 终端列表
type pager struct {
	Total   int    `json:"total"`
	PerPage int    `json:"per_page"`
	Page    int    `json:"page"`
	OrderBy string `json:"orderby"`
	SortBy  string `json:"sortby"`
	KeyWord string `json:"keyword"`
	Status  string `json:"status"`
}

type ListData struct {
	ProjectName string `json:"project_name"`
	Action      string `json:"action"`
	CategoryID  string `json:"categoryID"`
	Pager       pager  `json:"Pager"`
	User        string `json:"user"`
	Token       string `json:"token"`
}

var TerminalList = ListData{
	ProjectName: "default",
	Action:      "getTermList",
	Pager: pager{Total: -1, PerPage: 10, Page: 1, OrderBy: "",
		SortBy: "", KeyWord: "", Status: "running"},
	User: "root"}

// ---------------------------TCP数据---------------------------
var CommandData Command

type Command struct {
	// 机柜
	PowerOpen  []byte
	PowerClose []byte
	// 电视
	TvOpen  []byte
	TvClose []byte
	// 灯光
	LightAllOpen  []byte
	LightAllClose []byte
	// DVD
	DvdOpen  []byte
	DvdClose []byte
	// 灯光区域
	Light1Open       []byte
	Light1Close      []byte
	Spotlight1Open   []byte
	Spotlight1Close  []byte
	LightStrip1Open  []byte
	LightStrip1Close []byte
	Light2Open       []byte
	Light2Close      []byte
	Spotlight2Open   []byte
	Spotlight2Close  []byte
	LightStrip2Open  []byte
	LightStrip2Close []byte
	// DVD控制
	DvdPower       []byte
	DvdSetting     []byte
	DvdChannelAdd  []byte
	DvdChannelSub  []byte
	DvdUp          []byte
	DvdDown        []byte
	DvdLeft        []byte
	DvdRight       []byte
	DvdOk          []byte
	DvdVolumeAdd   []byte
	DvdVolumeSub   []byte
	DvdReturn      []byte
	DvdRetreat     []byte
	DvdHomeMenu    []byte
	DvdAdvance     []byte
	DvdMediaCenter []byte
}

func DataInitialize() {
	// ------------------------------------机柜------------------------------------
	//98 110 153 236 200 200 200 200 200 200 200 200 149 244 171 141
	CommandData.PowerOpen = []byte{'b', 'n', 0x99, 0xec,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x95, 0xf4, 0xab, 0x8d}
	//98 110 27 64 200 200 200 200 200 200 200 200 107 112 171 141
	//98 110 11 92 200 200 200 200 200 200 200 200 119 100 171 141
	//98 110 252 1 200 200 200 200 200 200 200 200 13 208 171 141
	CommandData.PowerClose = []byte{'b', 'n', 0x1b, 0x40,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x6b, 0x70, 0xab, 0x8d}
	// ------------------------------------电视------------------------------------
	//98 110 57 136 200 200 200 200 200 200 200 200 209 76 171 141
	//98 110 162 108 200 200 200 200 200 200 200 200 30 244 171 141
	//98 110 114 176 200 200 200 200 200 200 200 200 50 184 171 141
	CommandData.TvOpen = []byte{'b', 'n', 0x39, 0x88,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xd1, 0x4c, 0xab, 0x8d}
	//98 110 88 193 200 200 200 200 200 200 200 200 41 148 171 141
	//98 110 215 66 200 200 200 200 200 200 200 200 41 13 171 141
	//98 110 47 216 200 200 200 200 200 200 200 200 23 25 171 141
	CommandData.TvClose = []byte{'b', 'n', 0x58, 0xc1,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x29, 0x94, 0xab, 0x8d}
	// ------------------------------------灯光------------------------------------
	//98 110 28 204 200 200 200 200 200 200 200 200 248 112 171 141
	//98 110 19 72 200 200 200 200 200 200 200 200 107 64 171 141
	//98 110 57 184 200 200 200 200 200 200 200 200 1 76 171 141
	CommandData.LightAllOpen = []byte{'b', 'n', 0x1c, 0xcc,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xf8, 0x70, 0xab, 0x8d}
	// 98 110 36 2 200 200 200 200 200 200 200 200 54 220 171 141
	// 98 110 163 56 200 200 200 200 200 200 200 200 235 244 171 141
	// 98 110 76 208 200 200 200 200 200 200 200 200 44 232 171 141
	CommandData.LightAllClose = []byte{'b', 'n', 0x24, 0x2,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x36, 0xdc, 0xab, 0x8d}
	// ------------------------------------DVD------------------------------------
	//98 110 71 81 200 200 200 200 200 200 200 200 168 148 171 141
	//98 110 49 254 200 200 200 200 200 200 200 200 63 172 171 141
	//98 110 145 192 200 200 200 200 200 200 200 200 97 88 171 141
	CommandData.DvdOpen = []byte{'b', 'n', 0x47, 0x51,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xa8, 0x94, 0xab, 0x8d}
	//98 110 71 225 200 200 200 200 200 200 200 200 56 148 171 141
	//98 110 50 20 200 200 200 200 200 200 200 200 86 172 171 141
	//98 110 63 200 200 200 200 200 200 200 200 200 23 124 171 141
	CommandData.DvdClose = []byte{'b', 'n', 0x47, 0xe1,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x38, 0x94, 0xab, 0x8d}
	// 灯光区域1
	//98 110 65 176 200 200 200 200 200 200 200 200 1 232 171 141
	//98 110 139 148 200 200 200 200 200 200 200 200 47 244 171 141
	CommandData.Light1Open = []byte{'b', 'n', 0x41, 0xb0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x1, 0xe8, 0xab, 0x8d}
	//98 110 213 240 200 200 200 200 200 200 200 200 213 196 171 141
	//98 110 98 192 200 200 200 200 200 200 200 200 50 184 171 141
	//98 110 40 97 200 200 200 200 200 200 200 200 153 148 171 141
	CommandData.Light1Close = []byte{'b', 'n', 0xd5, 0xf0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xd5, 0xc4, 0xab, 0x8d}
	//98 110 41 172 200 200 200 200 200 200 200 200 229 25 171 141
	//98 110 4 58 200 200 200 200 200 200 200 200 78 52 171 141
	CommandData.Spotlight1Open = []byte{'b', 'n', 0x29, 0xac,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xe5, 0x19, 0xab, 0x8d}
	//98 110 134 0 200 200 200 200 200 200 200 200 150 88 171 141
	//98 110 45 244 200 200 200 200 200 200 200 200 49 172 171 141
	//98 110 211 193 200 200 200 200 200 200 200 200 164 208 171 141
	CommandData.Spotlight1Close = []byte{'b', 'n', 0x86, 0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x96, 0x58, 0xab, 0x8d}
	//98 110 166 192 200 200 200 200 200 200 200 200 118 88 171 141
	//98 110 51 252 200 200 200 200 200 200 200 200 63 25 171 141
	//98 110 62 152 200 200 200 200 200 200 200 200 230 76 171 141
	CommandData.LightStrip1Open = []byte{'b', 'n', 0xa6, 0xc0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x76, 0x58, 0xab, 0x8d}
	//98 110 83 80 200 200 200 200 200 200 200 200 179 232 171 141
	//98 110 20 240 200 200 200 200 200 200 200 200 20 64 171 141
	//98 110 23 99 200 200 200 200 200 200 200 200 138 13 171 141
	CommandData.LightStrip1Close = []byte{'b', 'n', 0x53, 0x50,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xb3, 0xe8, 0xab, 0x8d}
	// 灯光区域2
	//98 110 62 56 200 200 200 200 200 200 200 200 134 76 171 141
	//98 110 116 17 200 200 200 200 200 200 200 200 149 1 171 141
	//98 110 116 81 200 200 200 200 200 200 200 200 213 148 171 141
	CommandData.Light2Open = []byte{'b', 'n', 0x3e, 0x38,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x86, 0x4c, 0xab, 0x8d}
	//98 110 31 24 200 200 200 200 200 200 200 200 71 112 171 141
	//98 110 166 0 200 200 200 200 200 200 200 200 182 88 171 141
	CommandData.Light2Close = []byte{'b', 'n', 0x1f, 0x18,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x47, 0x70, 0xab, 0x8d}
	//98 110 32 81 200 200 200 200 200 200 200 200 129 220 171 141
	//98 110 146 124 200 200 200 200 200 200 200 200 30 244 171 141
	//98 110 25 228 200 200 200 200 200 200 200 200 13 112 171 141
	CommandData.Spotlight2Open = []byte{'b', 'n', 0x20, 0x51,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x81, 0xdc, 0xab, 0x8d}
	//98 110 60 128 200 200 200 200 200 200 200 200 204 124 171 141
	//98 110 227 129 200 200 200 200 200 200 200 200 116 208 171 141
	//98 110 146 192 200 200 200 200 200 200 200 200 98 244 171 141
	CommandData.Spotlight2Close = []byte{'b', 'n', 0x3c, 0x80,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xcc, 0x7c, 0xab, 0x8d}
	//98 110 20 248 200 200 200 200 200 200 200 200 28 64 171 141
	//98 110 125 48 200 200 200 200 200 200 200 200 189 184 171 141
	//98 110 24 147 200 200 200 200 200 200 200 200 187 13 171 141
	CommandData.LightStrip2Open = []byte{'b', 'n', 0x14, 0xf8,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x1c, 0x40, 0xab, 0x8d}
	//98 110 57 120 200 200 200 200 200 200 200 200 193 172 171 141
	//98 110 177 144 200 200 200 200 200 200 200 200 81 244 171 141
	//98 110 73 48 200 200 200 200 200 200 200 200 137 124 171 141
	CommandData.LightStrip2Close = []byte{'b', 'n', 0x39, 0x78,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xc1, 0xac, 0xab, 0x8d}
	// ------------------------------------DVD控制1------------------------------------
	//98 110 0 80 200 200 200 200 200 200 200 200 96 232 171 141
	//98 110 0 128 200 200 200 200 200 200 200 200 144 184 171 141
	CommandData.DvdPower = []byte{'b', 'n', 0x0, 0x80,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x60, 0xe8, 0xab, 0x8d}
	//98 110 5 192 200 200 200 200 200 200 200 200 213 136 171 141
	//98 110 0 208 200 200 200 200 200 200 200 200 224 64 171 141
	//98 110 26 160 200 200 200 200 200 200 200 200 202 13 171 141
	CommandData.DvdSetting = []byte{'b', 'n', 0x5, 0xc0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xd5, 0x88, 0xab, 0x8d}
	// ------------------------------------DVD控制2------------------------------------
	//98 110 0 70 200 200 200 200 200 200 200 200 86 52 171 141
	//98 110 2 200 200 200 200 200 200 200 200 200 218 76 171 141
	//98 110 24 32 200 200 200 200 200 200 200 200 72 208 171 141
	CommandData.DvdChannelAdd = []byte{'b', 'n', 0x0, 0x46,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x56, 0x34, 0xab, 0x8d}
	//98 110 2 16 200 200 200 200 200 200 200 200 34 25 171 141
	//98 110 7 8 200 200 200 200 200 200 200 200 31 244 171 141
	CommandData.DvdChannelSub = []byte{'b', 'n', 0x2, 0x10,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x22, 0x19, 0xab, 0x8d}
	//98 110 15 240 200 200 200 200 200 200 200 200 15 13 171 141
	//98 110 2 144 200 200 200 200 200 200 200 200 162 184 171 141
	CommandData.DvdUp = []byte{'b', 'n', 0xf, 0xf0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xf, 0xd, 0xab, 0x8d}
	//98 110 1 68 200 200 200 200 200 200 200 200 85 172 171 141
	//98 110 4 0 200 200 200 200 200 200 200 200 20 88 171 141
	CommandData.DvdDown = []byte{'b', 'n', 0x1, 0x44,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x55, 0xac, 0xab, 0x8d}
	//98 110 1 168 200 200 200 200 200 200 200 200 185 76 171 141
	//98 110 8 240 200 200 200 200 200 200 200 200 8 148 171 141
	//98 110 0 83 200 200 200 200 200 200 200 200 99 100 171 141
	CommandData.DvdLeft = []byte{'b', 'n', 0x1, 0xa8,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xb9, 0x4c, 0xab, 0x8d}
	//98 110 0 160 200 200 200 200 200 200 200 200 176 64 171 141
	//98 110 1 72 200 200 200 200 200 200 200 200 89 25 171 141
	//98 110 4 128 200 200 200 200 200 200 200 200 148 88 171 141
	CommandData.DvdRight = []byte{'b', 'n', 0x0, 0xa0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xb0, 0x40, 0xab, 0x8d}
	//98 110 0 50 200 200 200 200 200 200 200 200 66 52 171 141
	//98 110 4 128 200 200 200 200 200 200 200 200 148 136 171 141
	//98 110 0 220 200 200 200 200 200 200 200 200 236 112 171 141
	CommandData.DvdOk = []byte{'b', 'n', 0x0, 0x32,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x42, 0x34, 0xab, 0x8d}
	//98 110 4 16 200 200 200 200 200 200 200 200 36 13 171 141
	//98 110 0 23 200 200 200 200 200 200 200 200 39 100 171 141
	//98 110 3 32 200 200 200 200 200 200 200 200 51 208 171 141
	CommandData.DvdVolumeAdd = []byte{'b', 'n', 0x4, 0x10,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x24, 0xd, 0xab, 0x8d}
	//98 110 0 24 200 200 200 200 200 200 200 200 40 52 171 141
	//98 110 0 64 200 200 200 200 200 200 200 200 80 112 171 141
	//98 110 0 88 200 200 200 200 200 200 200 200 104 25 171 141
	CommandData.DvdVolumeSub = []byte{'b', 'n', 0x0, 0x24,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x28, 0x34, 0xab, 0x8d}
	// ------------------------------------DVD控制3------------------------------------
	//98 110 1 112 200 200 200 200 200 200 200 200 129 124 171 141
	//98 110 1 24 200 200 200 200 200 200 200 200 41 172 171 141
	//98 110 0 248 200 200 200 200 200 200 200 200 8 25 171 141
	CommandData.DvdReturn = []byte{'b', 'n', 0x1, 0x70,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x81, 0x7c, 0xab, 0x8d}
	//98 110 33 192 200 200 200 200 200 200 200 200 241 13 171 141
	//98 110 5 96 200 200 200 200 200 200 200 200 117 184 171 141
	//98 110 16 64 200 200 200 200 200 200 200 200 96 148 171 141
	CommandData.DvdRetreat = []byte{'b', 'n', 0x21, 0xc0,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0xf1, 0xd, 0xab, 0x8d}
	//98 110 0 28 200 200 200 200 200 200 200 200 44 52 171 141
	//98 110 0 148 200 200 200 200 200 200 200 200 164 172 171 141
	//98 110 2 0 200 200 200 200 200 200 200 200 18 88 171 141
	CommandData.DvdHomeMenu = []byte{'b', 'n', 0x0, 0x1c,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x2c, 0x34, 0xab, 0x8d}
	//98 110 1 8 200 200 200 200 200 200 200 200 25 64 171 141
	//98 110 34 240 200 200 200 200 200 200 200 200 34 13 171 141
	CommandData.DvdAdvance = []byte{'b', 'n', 0x1, 0x8,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x19, 0x40, 0xab, 0x8d}
	//98 110 1 16 200 200 200 200 200 200 200 200 33 184 171 141
	//98 110 4 224 200 200 200 200 200 200 200 200 244 208 171 141
	//98 110 0 33 200 200 200 200 200 200 200 200 49 100 171 141
	CommandData.DvdMediaCenter = []byte{'b', 'n', 0x1, 0x10,
		0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8,
		0x21, 0xb8, 0xab, 0x8d}
}

package data

type JewelType uint

const (
	GloriousVanity = JewelType(iota + 1)
	LethalPride
	BrutalRestraint
	MilitantFaith
	ElegantHubris
)

func (t JewelType) String() string {
	switch t {
	case GloriousVanity:
		return "輝煌的虛榮"
	case LethalPride:
		return "致命的驕傲"
	case BrutalRestraint:
		return "殘酷的紀律"
	case MilitantFaith:
		return "激進的信仰"
	case ElegantHubris:
		return "優雅的高傲"
	default:
		return "N/A"
	}
}

type Conqueror string

const (
	Xibaqua = Conqueror("賽巴昆")
	Zerphi  = Conqueror("澤佛伊")
	Ahuana  = Conqueror("阿呼阿娜")
	Doryani = Conqueror("多里亞尼")

	Kaom    = Conqueror("岡姆")
	Rakiata = Conqueror("拉其塔")
	Kiloava = Conqueror("基洛瓦")
	Akoya   = Conqueror("阿冦亞")

	Deshret = Conqueror("迪虛瑞特")
	Balbala = Conqueror("貝爾巴拉")
	Asenath = Conqueror("安賽娜絲")
	Nasima  = Conqueror("納西瑪")

	Venarius = Conqueror("維那利斯")
	Maxarius = Conqueror("瑪薩里歐斯")
	Dominus  = Conqueror("神主")
	Avarius  = Conqueror("伊爾莉斯")

	Cadiro   = Conqueror("卡迪羅")
	Victario = Conqueror("維多里奧")
	Chitus   = Conqueror("切特斯")
	Caspiro  = Conqueror("卡斯皮羅")
)

var TimelessJewelConquerors = map[JewelType]map[Conqueror]*TimelessJewelConqueror{
	GloriousVanity: {
		Xibaqua: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Zerphi: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Ahuana: &TimelessJewelConqueror{
			Index:   2,
			Version: 1,
		},
		Doryani: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
	},
	LethalPride: {
		Kaom: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Rakiata: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Kiloava: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
		Akoya: &TimelessJewelConqueror{
			Index:   3,
			Version: 1,
		},
	},
	BrutalRestraint: {
		Deshret: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Balbala: &TimelessJewelConqueror{
			Index:   1,
			Version: 1,
		},
		Asenath: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Nasima: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
	},
	MilitantFaith: {
		Venarius: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Maxarius: &TimelessJewelConqueror{
			Index:   1,
			Version: 1,
		},
		Dominus: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Avarius: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
	},
	ElegantHubris: {
		Cadiro: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Victario: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Chitus: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
		Caspiro: &TimelessJewelConqueror{
			Index:   3,
			Version: 1,
		},
	},
}

type Range struct {
	Min     uint32
	Max     uint32
	Special bool
}

var TimelessJewelSeedRanges = map[JewelType]Range{
	GloriousVanity: {
		Min: 100,
		Max: 8000,
	},
	LethalPride: {
		Min: 10000,
		Max: 18000,
	},
	BrutalRestraint: {
		Min: 500,
		Max: 8000,
	},
	MilitantFaith: {
		Min: 2000,
		Max: 10000,
	},
	ElegantHubris: {
		Min:     2000,
		Max:     160000,
		Special: true,
	},
}

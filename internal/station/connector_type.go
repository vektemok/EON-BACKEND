package station

type Connector string

const (
	Type1    Connector = "Type1"    // Yazaki — AC, распространён в Азии и США
	Type2    Connector = "Type2"    // Mennekes — AC, стандарт в Европе
	CHAdeMO  Connector = "CHAdeMO"  // DC, японский стандарт (Nissan, Mitsubishi)
	CCS1     Connector = "CCS1"     // Combined Charging System, США
	CCS2     Connector = "CCS2"     // Combined Charging System, Европа
	GB_T_AC  Connector = "GB/T-AC"  // Китайский стандарт переменного тока
	GB_T_DC  Connector = "GB/T-DC"  // Китайский стандарт постоянного тока
	Tesla    Connector = "Tesla"    // Проприетарный разъём (в США)
	Schuko   Connector = "Schuko"   // Обычная бытовая розетка (EU)
	CEE_Red  Connector = "CEE-Red"  // Промышленная розетка 3-фазная
	Wireless Connector = "Wireless" // Беспроводная зарядка (редко, прототипы)
)

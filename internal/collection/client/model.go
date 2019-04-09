package client

type Boxver struct {
	Boxid string `json:"boxid"`
	Tick  string `json:"tick"`
	Mcuid string `json:"mcuid"`
	Imei  string `json:"imei"`
	Iccid string `json:"iccid"`
	Imsi  string `json:"imsi"`
	Boxhw string `json:"boxhw"`
	Boxsw string `json:"boxsw"`
	Obdhw string `json:"obdhw"`
	Obdsw string `json:"obdsw"`
}

type Protocol struct {
	Boxid   string      `json:"boxid"`
	Tick    string      `json:"tick"`
	Vin     string      `json:"vin"`
	Obdtype string      `json:"obdtype"`
	Page    int         `json:"page"`
	Info    interface{} `json:"info"`
}

type Dbinfo struct {
	Boxid   string `json:"boxid"`
	Tick    string `json:"tick"`
	Dbname  string `json:"dbname"`
	Dbver   string `json:"dbver"`
	Support string `json:"support"`
}

type Boxstate struct {
	Boxid    string      `json:"boxid"`
	Tick     string      `json:"tick"`
	Hpflag   int         `json:"hpflag"`
	Hpangle  float64     `json:"hpangle"`
	Dpflag   int         `json:"dpflag"`
	Dpangle  float64     `json:"dpangle"`
	Download int         `json:"download"`
	Boxdtc   interface{} `json:"boxdtc"`
}

type GPS struct {
	Boxid string `json:"boxid"`
	Data  []GPSData
}

type GPSData struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Flag     int     `json:"flag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Height   int     `json:"height"`
	Speed    float64 `json:"speed"`
	Direct   int     `json:"direct"`
	Rpm      int     `json:"rpm"`
	Vs       int     `json:"vs"`
	Dpkm     int     `json:"dpkm"`
	Ttkm     int     `json:"ttkm"`
	Tttime   int     `json:"tttime"`
	Sgkm     int     `json:"sgkm"`
	Sgtime   int     `json:"sgtime"`
}

type Engstate struct {
	Boxid     string  `json:"boxid"`
	Tick      string  `json:"tick"`
	Travelid  string  `json:"travelid"`
	Eventid   int     `json:"eventid"`
	Gpsflag   int     `json:"gpsflag"`
	Lon       float64 `json:"lon"`
	Lat       float64 `json:"lat"`
	Fuel      float64 `json:"fuel"`
	Idlefuel  float64 `json:"idlefuel"`
	Mileage   float64 `json:"mileage"`
	Maxspeed  int     `json:"maxspeed"`
	Runtime   int     `json:"runtime"`
	Hottime   int     `json:"hottime"`
	Idletime  int     `json:"idletime"`
	Speed1    float64 `json:"speed1"`
	Speed2    float64 `json:"speed2"`
	Speed3    float64 `json:"speed3"`
	Speed4    float64 `json:"speed4"`
	Speed5    float64 `json:"speed5"`
	Speed6    float64 `json:"speed6"`
	Comfort   int     `json:"comfort"`
	Dashboard int     `json:"dashboard"`
	Fuelflag  int     `json:"fuelflag"`
	Fuellevel float64 `json:"fuellevel"`
}

type Drive struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Svs      int     `json:"svs"`
	Evs      int     `json:"evs"`
	Time     int     `json:"time"`
	Acce     int     `json:"acce"`
}

type Run struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Flag     int     `json:"flag"`
	Value    int     `json:"value"`
	Maxv     int     `json:"maxv"`
	Time     int     `json:"time"`
}

type Traffic struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Part     int     `json:"part"`
	Level    byte    `json:"level"`
	Vss      int     `json:"vss"`
	Vse      int     `json:"vse"`
	Distance int     `json:"distance"`
	Time     int     `json:"time"`
	Direct   int     `json:"direct"`
}

type Box struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Flag     int     `json:"flag"`
	Time     int     `json:"time"`
}

type Tyre struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Flag     int     `json:"flag"`
	Time     int     `json:"time"`
}

type Fuel struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Value1   int     `json:"value1"`
	Value2   int     `json:"value2"`
	Fuelflag int     `json:"fuelflag"`
}

type Dtc struct {
	Boxid    string      `json:"boxid"`
	Tick     string      `json:"tick"`
	Travelid string      `json:"travelid"`
	Eventid  int         `json:"eventid"`
	Gpsflag  int         `json:"gpsflag"`
	Lon      float64     `json:"lon"`
	Lat      float64     `json:"lat"`
	Dtc      interface{} `json:"dtc"`
}

type Area struct {
	Boxid    string  `json:"boxid"`
	Tick     string  `json:"tick"`
	Travelid string  `json:"travelid"`
	Eventid  int     `json:"eventid"`
	Gpsflag  int     `json:"gpsflag"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Areaid   int     `json:"areaid"`
	Flag     int     `json:"flag"`
	Time     int     `json:"time"`
}

type Statistic struct {
	Boxid     string  `json:"boxid"`
	Tick      string  `json:"tick"`
	Travelid  string  `json:"travelid"`
	Fuel      float64 `json:"fuel"`
	Idlefuel  float64 `json:"idlefuel"`
	Mileage   float64 `json:"mileage"`
	Maxspeed  int     `json:"maxspeed"`
	Runtime   int     `json:"runtime"`
	Hottime   int     `json:"hottime"`
	Idletime  int     `json:"idletime"`
	Speed1    float64 `json:"speed1"`
	Speed2    float64 `json:"speed2"`
	Speed3    float64 `json:"speed3"`
	Speed4    float64 `json:"speed4"`
	Speed5    float64 `json:"speed5"`
	Speed6    float64 `json:"speed6"`
	Comfort   int     `json:"comfort"`
	Dashboard int     `json:"dashboard"`
	Fuelflag  int     `json:"fuelflag"`
	Fuellevel float64 `json:"fuellevel"`
}

type Boxbody struct {
	Boxid    string `json:"boxid"`
	Tick     string `json:"tick"`
	Travelid string `json:"travelid"`
	Value    int    `json:"value"`
}

type Cardata struct {
	Boxid      string  `json:"boxid"`
	Tick       string  `json:"tick"`
	Travelid   string  `json:"travelid"`
	Rpm        int     `json:"rpm"`
	Speed      int     `json:"speed"`
	Batv       float64 `json:"batv"`
	Ect        float64 `json:"ect"`
	Mileage    int     `json:"mileage"`
	Unit       int     `json:"unit"`
	Fuel       float64 `json:"fuel"`
	Periodkm   int     `json:"periodkm"`
	Periodday  int     `json:"periodday"`
	Accumkm    int     `json:"accumkm"`
	Accumfuel  float64 `json:"accumfuel"`
	Centerlock int     `json:"centerlock"`
	Tyre       int     `json:"tyre"`
	Oiltemp    float64 `json:"oiltemp"`
	Oilqua     int     `json:"oilqua"`
	Oillevel   int     `json:"oillevel"`
	Fuelqua    int     `json:"fuelqua"`
	Fuelpre    float64 `json:"fuelpre"`
	Fuelcon    int     `json:"fuelcon"`
	Batcamax   int     `json:"batcamax"`
	Batcanow   int     `json:"batcanow"`
	Bathealth  int     `json:"bathealth"`
	Batimp     float64 `json:"batimp"`
	Drvf       int     `json:"drvf"`
	Drvr       int     `json:"drvr"`
	Passf      int     `json:"passf"`
	Passr      int     `json:"passr"`
	Airflow    float64 `json:"airflow"`
	Tps        float64 `json:"tps"`
	Ignd       float64 `json:"ignd"`
	Oilpress   float64 `json:"oilpress"`
	Fuelpress  float64 `json:"fuelpress"`
	Batcurrent float64 `json:"batcurrent"`
	Battemp    float64 `json:"battemp"`
	Chargev    int     `json:"chargev"`
	Chargep    float64 `json:"chargep"`
	Charget    int     `json:"cherget"`
	Chargeh    float64 `json:"chergeh"`
	Mileagebat int     `json:"mileagebat"`
	Motorrpm   float64 `json:"motorrpm"`
	Clutchp    int     `json:"clutchp"`
	Clutchg    int     `json:"clutchg"`
	Drvsf      int     `json:"drv_sf"`
	Drvsr      int     `json:"drv_sr"`
	Passsf     int     `json:"pass_sf"`
	Passsr     int     `json:"pass_sr"`
	Fzd        int     `json:"fzd"`
}

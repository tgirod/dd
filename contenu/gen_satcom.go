package main

import (
	"fmt"
	"math/rand"

	"gopkg.in/yaml.v3"
)

// Register représente registre mémoire qui peut être modifié pour contrôler quelque chose
type Register struct {
	Server      string
	Group       string
	Description string
	State       RegisterState   // état actuel
	Options     []RegisterState // valeurs possible
}
type RegisterState string

type EntrySat struct {
	Code      string
	splitKeys []string
	Owner     string
	title     string
	inside    string
}

var satFred = []RegisterState{"613MHz", "827MHz", "1010MHz", "1198MHz", "1399MHz", "1632MHz"}

var satEntry = []EntrySat{
	{"GEO-EU-D01", []string{"GEO", "EU", "D01"}, "", "Europole D01", `azimut:222.473862
altitude:57.545902
Europole District D01`},
	{"GEO-EU-D02", []string{"GEO", "EU", "D02"}, "", "Europole D02", `azimut:239.324897
altitude:40.278407
Europole District D02`},
	{"GEO-EU-D03", []string{"GEO", "EU", "D03"}, "", "Europole D03", `azimut:109.366561
altitude:76.071807
Europole District D03`},
	{"GEO-EU-D04", []string{"GEO", "EU", "D04"}, "", "Europole D04", `azimut:329.297929
altitude:27.991250
Europole District D04`},
	{"GEO-EU-D05", []string{"GEO", "EU", "D05"}, "", "Europole D05", `azimut:196.971308
altitude:89.923900
Europole District D05`},
	{"GEO-EU-D06", []string{"GEO", "EU", "D06"}, "", "Europole D06", `azimut:128.213744
altitude:72.986507
Europole District D06`},
	{"GEO-EU-D07", []string{"GEO", "EU", "D07"}, "", "Europole D07", `azimut:188.000197
altitude:25.475341
Europole District D07`},
	{"GEO-EU-D08", []string{"GEO", "EU", "D08"}, "", "Europole D08", `azimut:18.387379
altitude:39.189764
Europole District D08`},
	{"GEO-EU-D09", []string{"GEO", "EU", "D09"}, "", "Europole D09", `azimut:25.997548
altitude:2.442340
Europole District D09`},
	{"GEO-EU-D10", []string{"GEO", "EU", "D10"}, "", "Europole D10", `azimut:63.563148
altitude:15.524259
Europole District D10`},
	{"GEO-EU-D11", []string{"GEO", "EU", "D11"}, "", "Europole D11", `azimut:45.742992
altitude:66.014371
Europole District D11`},
	{"GEO-EU-D12", []string{"GEO", "EU", "D12"}, "", "Europole D12", `azimut:347.864768
altitude:52.184109
Europole District D12`},
	{"GEO-EU-D13", []string{"GEO", "EU", "D13"}, "", "Europole D13", `azimut:243.115459
altitude:82.627421
Europole District D13`},
	{"GEO-EU-D14", []string{"GEO", "EU", "D14"}, "", "Europole D14", `azimut:116.483503
altitude:18.400539
Europole District D14`},
	{"GEO-EU-D15", []string{"GEO", "EU", "D15"}, "", "Europole D15", `azimut:340.027907
altitude:44.906145
Europole District D15`},
	{"GEO-EU-D16", []string{"GEO", "EU", "D16"}, "", "Europole D16", `azimut:156.679333
altitude:59.225136
Europole District D16`},
	{"GEO-EU-D17", []string{"GEO", "EU", "D17"}, "", "Europole D17", `azimut:98.859250
altitude:70.193535
Europole District D17`},
	{"GEO-EU-D18", []string{"GEO", "EU", "D18"}, "", "Europole D18", `azimut:249.598879
altitude:35.274047
Europole District D18`},
	{"GEO-EU-D19", []string{"GEO", "EU", "D19"}, "", "Europole D19", `azimut:39.792230
altitude:84.093000
Europole District D19`},
	{"GEO-EU-D20", []string{"GEO", "EU", "D20"}, "", "Europole D20", `azimut:181.817280
altitude:44.512595
Europole District D20`},
	{"GEO-EU-D21", []string{"GEO", "EU", "D21"}, "", "Europole D21", `azimut:150.167960
altitude:85.991215
Europole District D21`},
	{"GEO-EU-D22", []string{"GEO", "EU", "D22"}, "", "Europole D22", `azimut:239.977281
altitude:75.689278
Europole District D22`},
	{"GEO-EU-D23", []string{"GEO", "EU", "D23"}, "", "Europole D23", `azimut:351.246429
altitude:34.655470
Europole District D23`},
	{"GEO-EU-D24", []string{"GEO", "EU", "D24"}, "", "Europole D24", `azimut:160.687062
altitude:65.748652
Europole District D24`},
	{"GEO-EU-D25", []string{"GEO", "EU", "D25"}, "", "Europole D25", `azimut:199.657318
altitude:16.504889
Europole District D25`},
	{"GEO-EU-D26", []string{"GEO", "EU", "D26"}, "", "Europole D26", `azimut:113.082227
altitude:70.553254
Europole District D26`},
	{"GEO-EU-D27", []string{"GEO", "EU", "D27"}, "", "Europole D27", `azimut:149.928442
altitude:38.723569
Europole District D27`},
	{"GEO-EU-D28", []string{"GEO", "EU", "D28"}, "", "Europole D28", `azimut:195.343852
altitude:66.718099
Europole District D28`},
	{"GEO-EU-D29", []string{"GEO", "EU", "D29"}, "", "Europole D29", `azimut:68.053002
altitude:59.244626
Europole District D29`},
	{"GEO-EU-D30", []string{"GEO", "EU", "D30"}, "", "Europole D30", `azimut:109.773570
altitude:75.527002
Europole District D30`},
	{"GEO-AM-D01", []string{"GEO", "AM", "D01"}, "", "Amerique D01", `azimut:73.597028
altitude:49.878709
Amerique District D01`},
	{"GEO-AM-D02", []string{"GEO", "AM", "D02"}, "", "Amerique D02", `azimut:123.021633
altitude:17.279446
Amerique District D02`},
	{"GEO-AM-D03", []string{"GEO", "AM", "D03"}, "", "Amerique D03", `azimut:343.543004
altitude:40.210107
Amerique District D03`},
	{"GEO-AM-D04", []string{"GEO", "AM", "D04"}, "", "Amerique D04", `azimut:118.704682
altitude:30.886185
Amerique District D04`},
	{"GEO-AM-D05", []string{"GEO", "AM", "D05"}, "", "Amerique D05", `azimut:332.719647
altitude:66.663091
Amerique District D05`},
	{"GEO-AM-D06", []string{"GEO", "AM", "D06"}, "", "Amerique D06", `azimut:170.711919
altitude:38.124117
Amerique District D06`},
	{"GEO-AM-D07", []string{"GEO", "AM", "D07"}, "", "Amerique D07", `azimut:95.659856
altitude:19.927787
Amerique District D07`},
	{"GEO-AM-D08", []string{"GEO", "AM", "D08"}, "", "Amerique D08", `azimut:112.172492
altitude:19.548745
Amerique District D08`},
	{"GEO-AM-D09", []string{"GEO", "AM", "D09"}, "", "Amerique D09", `azimut:45.613917
altitude:51.208722
Amerique District D09`},
	{"GEO-AM-D10", []string{"GEO", "AM", "D10"}, "", "Amerique D10", `azimut:330.125659
altitude:73.166312
Amerique District D10`},
	{"GEO-AM-D11", []string{"GEO", "AM", "D11"}, "", "Amerique D11", `azimut:87.738024
altitude:46.632757
Amerique District D11`},
	{"GEO-AM-D12", []string{"GEO", "AM", "D12"}, "", "Amerique D12", `azimut:155.679631
altitude:89.617381
Amerique District D12`},
	{"GEO-AM-D13", []string{"GEO", "AM", "D13"}, "", "Amerique D13", `azimut:175.463825
altitude:13.228532
Amerique District D13`},
	{"GEO-AM-D14", []string{"GEO", "AM", "D14"}, "", "Amerique D14", `azimut:182.310405
altitude:12.549442
Amerique District D14`},
	{"GEO-AM-D15", []string{"GEO", "AM", "D15"}, "", "Amerique D15", `azimut:79.390452
altitude:5.071440
Amerique District D15`},
	{"GEO-AM-D16", []string{"GEO", "AM", "D16"}, "", "Amerique D16", `azimut:110.318744
altitude:33.526340
Amerique District D16`},
	{"GEO-AM-D17", []string{"GEO", "AM", "D17"}, "", "Amerique D17", `azimut:298.195798
altitude:84.808764
Amerique District D17`},
	{"GEO-AM-D18", []string{"GEO", "AM", "D18"}, "", "Amerique D18", `azimut:2.873332
altitude:51.596388
Amerique District D18`},
	{"GEO-AM-D19", []string{"GEO", "AM", "D19"}, "", "Amerique D19", `azimut:296.794890
altitude:71.677123
Amerique District D19`},
	{"GEO-AM-D20", []string{"GEO", "AM", "D20"}, "", "Amerique D20", `azimut:231.170081
altitude:30.647222
Amerique District D20`},
	{"GEO-AM-D21", []string{"GEO", "AM", "D21"}, "", "Amerique D21", `azimut:89.500920
altitude:30.522516
Amerique District D21`},
	{"GEO-AM-D22", []string{"GEO", "AM", "D22"}, "", "Amerique D22", `azimut:323.958919
altitude:30.437744
Amerique District D22`},
	{"GEO-AM-D23", []string{"GEO", "AM", "D23"}, "", "Amerique D23", `azimut:253.869255
altitude:30.920316
Amerique District D23`},
	{"GEO-AM-D24", []string{"GEO", "AM", "D24"}, "", "Amerique D24", `azimut:186.124318
altitude:62.824878
Amerique District D24`},
	{"GEO-AM-D25", []string{"GEO", "AM", "D25"}, "", "Amerique D25", `azimut:341.876998
altitude:86.569408
Amerique District D25`},
	{"GEO-AM-D26", []string{"GEO", "AM", "D26"}, "", "Amerique D26", `azimut:345.379005
altitude:56.438396
Amerique District D26`},
	{"GEO-AM-D27", []string{"GEO", "AM", "D27"}, "", "Amerique D27", `azimut:286.611533
altitude:62.474308
Amerique District D27`},
	{"GEO-AM-D28", []string{"GEO", "AM", "D28"}, "", "Amerique D28", `azimut:153.424171
altitude:2.944232
Amerique District D28`},
	{"GEO-AM-D29", []string{"GEO", "AM", "D29"}, "", "Amerique D29", `azimut:300.060128
altitude:70.495054
Amerique District D29`},
	{"GEO-AM-D30", []string{"GEO", "AM", "D30"}, "", "Amerique D30", `azimut:213.636218
altitude:8.239724
Amerique District D30`},
	{"GEO-AS-D01", []string{"GEO", "AS", "D01"}, "", "Asie D01", `azimut:342.051745
altitude:42.218323
Asie District D01`},
	{"GEO-AS-D02", []string{"GEO", "AS", "D02"}, "", "Asie D02", `azimut:250.259304
altitude:71.493491
Asie District D02`},
	{"GEO-AS-D03", []string{"GEO", "AS", "D03"}, "", "Asie D03", `azimut:95.251740
altitude:48.210691
Asie District D03`},
	{"GEO-AS-D04", []string{"GEO", "AS", "D04"}, "", "Asie D04", `azimut:183.859707
altitude:76.592457
Asie District D04`},
	{"GEO-AS-D05", []string{"GEO", "AS", "D05"}, "", "Asie D05", `azimut:333.600815
altitude:61.490801
Asie District D05`},
	{"GEO-AS-D06", []string{"GEO", "AS", "D06"}, "", "Asie D06", `azimut:33.457971
altitude:64.460645
Asie District D06`},
	{"GEO-AS-D07", []string{"GEO", "AS", "D07"}, "", "Asie D07", `azimut:205.287459
altitude:21.342705
Asie District D07`},
	{"GEO-AS-D08", []string{"GEO", "AS", "D08"}, "", "Asie D08", `azimut:290.712095
altitude:83.393417
Asie District D08`},
	{"GEO-AS-D09", []string{"GEO", "AS", "D09"}, "", "Asie D09", `azimut:129.804248
altitude:56.810036
Asie District D09`},
	{"GEO-AS-D10", []string{"GEO", "AS", "D10"}, "", "Asie D10", `azimut:181.933236
altitude:82.832447
Asie District D10`},
	{"GEO-AS-D11", []string{"GEO", "AS", "D11"}, "", "Asie D11", `azimut:1.056958
altitude:46.733304
Asie District D11`},
	{"GEO-AS-D12", []string{"GEO", "AS", "D12"}, "", "Asie D12", `azimut:130.245444
altitude:38.701721
Asie District D12`},
	{"GEO-AS-D13", []string{"GEO", "AS", "D13"}, "", "Asie D13", `azimut:11.754200
altitude:29.779863
Asie District D13`},
	{"GEO-AS-D14", []string{"GEO", "AS", "D14"}, "", "Asie D14", `azimut:63.656253
altitude:50.797300
Asie District D14`},
	{"GEO-AS-D15", []string{"GEO", "AS", "D15"}, "", "Asie D15", `azimut:240.014087
altitude:64.780596
Asie District D15`},
	{"GEO-AS-D16", []string{"GEO", "AS", "D16"}, "", "Asie D16", `azimut:185.696595
altitude:44.346859
Asie District D16`},
	{"GEO-AS-D17", []string{"GEO", "AS", "D17"}, "", "Asie D17", `azimut:204.779937
altitude:13.827560
Asie District D17`},
	{"GEO-AS-D18", []string{"GEO", "AS", "D18"}, "", "Asie D18", `azimut:92.243999
altitude:9.715407
Asie District D18`},
	{"GEO-AS-D19", []string{"GEO", "AS", "D19"}, "", "Asie D19", `azimut:199.837823
altitude:58.591530
Asie District D19`},
	{"GEO-AS-D20", []string{"GEO", "AS", "D20"}, "", "Asie D20", `azimut:53.849202
altitude:36.692392
Asie District D20`},
	{"GEO-AS-D21", []string{"GEO", "AS", "D21"}, "", "Asie D21", `azimut:142.217905
altitude:76.522163
Asie District D21`},
	{"GEO-AS-D22", []string{"GEO", "AS", "D22"}, "", "Asie D22", `azimut:75.695987
altitude:62.240956
Asie District D22`},
	{"GEO-AS-D23", []string{"GEO", "AS", "D23"}, "", "Asie D23", `azimut:285.624536
altitude:57.488524
Asie District D23`},
	{"GEO-AS-D24", []string{"GEO", "AS", "D24"}, "", "Asie D24", `azimut:274.254944
altitude:45.720608
Asie District D24`},
	{"GEO-AS-D25", []string{"GEO", "AS", "D25"}, "", "Asie D25", `azimut:357.594847
altitude:36.241697
Asie District D25`},
	{"GEO-AS-D26", []string{"GEO", "AS", "D26"}, "", "Asie D26", `azimut:87.273324
altitude:84.761897
Asie District D26`},
	{"GEO-AS-D27", []string{"GEO", "AS", "D27"}, "", "Asie D27", `azimut:192.987743
altitude:39.505095
Asie District D27`},
	{"GEO-AS-D28", []string{"GEO", "AS", "D28"}, "", "Asie D28", `azimut:330.694254
altitude:67.430962
Asie District D28`},
	{"GEO-AS-D29", []string{"GEO", "AS", "D29"}, "", "Asie D29", `azimut:127.146828
altitude:19.723466
Asie District D29`},
	{"GEO-AS-D30", []string{"GEO", "AS", "D30"}, "", "Asie D30", `azimut:74.864500
altitude:75.241127
Asie District D30`},
	{"GEO-AU-D01", []string{"GEO", "AU", "D01"}, "", "Australie D01", `azimut:292.083689
altitude:88.194980
Australie District D01`},
	{"GEO-AU-D02", []string{"GEO", "AU", "D02"}, "", "Australie D02", `azimut:82.831966
altitude:21.731413
Australie District D02`},
	{"GEO-AU-D03", []string{"GEO", "AU", "D03"}, "", "Australie D03", `azimut:154.231428
altitude:1.419308
Australie District D03`},
	{"GEO-AU-D04", []string{"GEO", "AU", "D04"}, "", "Australie D04", `azimut:24.530170
altitude:89.837985
Australie District D04`},
	{"GEO-AU-D05", []string{"GEO", "AU", "D05"}, "", "Australie D05", `azimut:57.894963
altitude:45.076034
Australie District D05`},
	{"GEO-AU-D06", []string{"GEO", "AU", "D06"}, "", "Australie D06", `azimut:342.363100
altitude:21.464879
Australie District D06`},
	{"GEO-AU-D07", []string{"GEO", "AU", "D07"}, "", "Australie D07", `azimut:78.005193
altitude:5.087446
Australie District D07`},
	{"GEO-AU-D08", []string{"GEO", "AU", "D08"}, "", "Australie D08", `azimut:225.866362
altitude:67.567458
Australie District D08`},
	{"GEO-AU-D09", []string{"GEO", "AU", "D09"}, "", "Australie D09", `azimut:167.444574
altitude:61.857512
Australie District D09`},
	{"GEO-AU-D10", []string{"GEO", "AU", "D10"}, "", "Australie D10", `azimut:255.402402
altitude:36.556867
Australie District D10`},
	{"GEO-AU-D11", []string{"GEO", "AU", "D11"}, "", "Australie D11", `azimut:140.726251
altitude:11.197472
Australie District D11`},
	{"GEO-AU-D12", []string{"GEO", "AU", "D12"}, "", "Australie D12", `azimut:161.236059
altitude:36.436500
Australie District D12`},
	{"GEO-AU-D13", []string{"GEO", "AU", "D13"}, "", "Australie D13", `azimut:17.379168
altitude:33.483234
Australie District D13`},
	{"GEO-AU-D14", []string{"GEO", "AU", "D14"}, "", "Australie D14", `azimut:51.559107
altitude:18.235380
Australie District D14`},
	{"GEO-AU-D15", []string{"GEO", "AU", "D15"}, "", "Australie D15", `azimut:108.356927
altitude:73.259122
Australie District D15`},
	{"GEO-AU-D16", []string{"GEO", "AU", "D16"}, "", "Australie D16", `azimut:165.263461
altitude:1.147638
Australie District D16`},
	{"GEO-AU-D17", []string{"GEO", "AU", "D17"}, "", "Australie D17", `azimut:316.051053
altitude:70.737982
Australie District D17`},
	{"GEO-AU-D18", []string{"GEO", "AU", "D18"}, "", "Australie D18", `azimut:76.158201
altitude:64.184684
Australie District D18`},
	{"GEO-AU-D19", []string{"GEO", "AU", "D19"}, "", "Australie D19", `azimut:110.508745
altitude:49.884585
Australie District D19`},
	{"GEO-AU-D20", []string{"GEO", "AU", "D20"}, "", "Australie D20", `azimut:255.635603
altitude:77.332306
Australie District D20`},
	{"GEO-AU-D21", []string{"GEO", "AU", "D21"}, "", "Australie D21", `azimut:277.905640
altitude:72.924568
Australie District D21`},
	{"GEO-AU-D22", []string{"GEO", "AU", "D22"}, "", "Australie D22", `azimut:239.945763
altitude:45.988657
Australie District D22`},
	{"GEO-AU-D23", []string{"GEO", "AU", "D23"}, "", "Australie D23", `azimut:44.763104
altitude:87.511096
Australie District D23`},
	{"GEO-AU-D24", []string{"GEO", "AU", "D24"}, "", "Australie D24", `azimut:288.379136
altitude:78.406113
Australie District D24`},
	{"GEO-AU-D25", []string{"GEO", "AU", "D25"}, "", "Australie D25", `azimut:130.566795
altitude:26.739230
Australie District D25`},
	{"GEO-AU-D26", []string{"GEO", "AU", "D26"}, "", "Australie D26", `azimut:122.319196
altitude:86.855428
Australie District D26`},
	{"GEO-AU-D27", []string{"GEO", "AU", "D27"}, "", "Australie D27", `azimut:259.975290
altitude:21.280427
Australie District D27`},
	{"GEO-AU-D28", []string{"GEO", "AU", "D28"}, "", "Australie D28", `azimut:96.620961
altitude:41.302436
Australie District D28`},
	{"GEO-AU-D29", []string{"GEO", "AU", "D29"}, "", "Australie D29", `azimut:146.349130
altitude:81.275968
Australie District D29`},
	{"GEO-AU-D30", []string{"GEO", "AU", "D30"}, "", "Australie D30", `azimut:75.093528
altitude:80.022612
Australie District D30`},
	{"LEO-SATCOM", []string{"LEO", "SATCOM"}, "", "Constellation SATCOM orbite LEO", "6a8e2a76-b0b7-42e4-aec4-9af7d0b1339e"},
	{"LEO-STARLINK", []string{"LEO", "STARLINK"}, "", "Constellation STARLINK orbite LEO", "39a4c6ac-3710-4aca-b802-8ab7bce8b6fa"},
	{"LEO-VIASAT", []string{"LEO", "VIASAT"}, "", "Constellation VIASAT orbite LEO", "843521bc-30f7-4d73-b68c-a1ea707da880"},
	{"LEO-IRIDIUM", []string{"LEO", "IRIDIUM"}, "", "Constellation IRIDIUM orbite LEO", "10e77c5c-ee5e-4324-9547-f2856ea3e3ac"},
	{"MEO-SATCOM", []string{"MEO", "SATCOM"}, "", "Constellation SATCOM orbite MEO", "fbc01335-8224-43e0-a175-298e43832f96"},
	{"MEO-STARLINK", []string{"MEO", "STARLINK"}, "", "Constellation STARLINK orbite MEO", "704cf5b2-1ccd-42ae-a93a-054fa65f7950"},
	{"MEO-VIASAT", []string{"MEO", "VIASAT"}, "", "Constellation VIASAT orbite MEO", "7ba9b084-442f-445e-97ce-b3936c368079"},
	{"MEO-IRIDIUM", []string{"MEO", "IRIDIUM"}, "", "Constellation IRIDIUM orbite MEO", "c65c0fe2-0ada-4555-b641-439239426488"},
}

func rndState(states []RegisterState) RegisterState {
	return states[rand.Intn(len(states))]
}

func genSatRegistry(sat []EntrySat, states []RegisterState) []Register {
	var res []Register

	for _, satEntry := range sat {
		key := satEntry.Code
		state := rndState(states)

		reg := Register{
			Server:      "satcom.legba.d22.eu",
			Group:       "sat",
			Description: key,
			State:       state,
			Options:     states,
		}
		//fmt.Printf("REG %v\n", reg)
		res = append(res, reg)
	}
	return res
}

var regDevice []Register

// génère du YAML pour les registres de Satcom.Legba
// go run contenu/gen_satcom.go > reg_satcom.yaml
func main() {
	rand.Seed(0)

	regDevice = append(regDevice, genSatRegistry(satEntry, satFred)...)

	yamlDevice, err := yaml.Marshal(regDevice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("## Devices *****\n%s\n", yamlDevice)
}

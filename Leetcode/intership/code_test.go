package intership

import (
	"fmt"
	"testing"
)

func TestIntegerReplacement(t *testing.T) {
	for i := 2; i <= 100; i++ {
		fmt.Printf("%d %d\n", 65535, IntegerReplacement(65535))
	}
}

func TestCountGoodSubstrings(t *testing.T) {
	str := "xyzzaz"
	println(CountGoodSubstrings(str))
}

func TestOriginalDigits(t *testing.T) {
	str := "ertfsxvxttiorseenivsoiwosefeinoinwoioousiieihtfirrnioeenwiortwsefwnnniseoisontieiitnvsovthrenwfitfenoiwouwtrtnxneisieinovhuieenrveenoiuoireooieiuittefotnoeeftfsetooeniuvrsnvetwieetvoneeoiieeiiirhftnrntihenseveeorioriononnuhweixefoiotowxreossentetresvzoerfeonvhfoestveooxesintstvefewrifsneeexveefoxonvueituwoieefvofwrfistixeesooxoeweteortfevtsuffohnfruiwitnuxexrevrhfenwenofnzfvsiofeneesoshoefxoefieiieeoueofoivshvtrsoneenenwiexnifttfeoooetesnouewtneisfeenwiiwnvuntnwuesniffweeoretviftseeexeoetzntefnfoxesfforvrwrrhrewtegwsnvinotonitvxeerwooeefswxxtfixenwfxioevierefevinviwvxonfitouevwxnshiewuntwunniennvunvfnnxinovfsvohrresenffintfissstnrfwinsegrrnesentoitoieneiiinuenrtssgosetoufnniuesneseonesetnstiofieeosvoowooeueietrheiienstifweeueitioretiwefwvonsernsxrtteeiisofttwnorovenoisnrsoxseoeeoeeovnoweeowhfirovesxeutisnnviieneetvuxnreoenxinfnrvxntooeonroewxeoirxenfeeehietriuoofxiexwntotointhwvtnfwtnrieeofifinevsfxioninsvntenrhwiwovfenwfrffiofeiehuriieeeswooeofienuoeifoffrfetewoofeusoeiireootseisouvoofntexsnorovrersoovnoovtvvreineerrieeooxtoeseeowixtfeoettiornofunintonsrenoooosxifsrviuitonnetosfwnewnetoznvfhueovwrnnooexnnfeieninetihvwnrexiinvnieevexnwnvtfewexftixiernoonfeinffisofoefevviownoeionforfevufoefreisetvovixxoehuiniivvirewsrsoxoiohtoovvveioertfhveoitoxoiexniieneteiwerngfiesiiifwehsifrteieiwoueenwssnrtxnoevefrtxtoesnferxveeoeifoxuivononeeentsohofwexinrweuovoutoenhnnxenveoihfstturoeeeitrwiirisiinuhrosooiinffnoriweixeiuuenxtoouwosnnshiixnnefgereothioiinrfnneioeenitofieerwtreniortiweoinniewinxeerxtneeneoofutiiirrnsvoefoetssxonfxenowirutohnxerisweewuntnwfwounnentnohevnooxwwiifsoooevtviwsfeounfeitfxostgfgtnnfsounireiiingofxwesfroeofsxefirstrfnerfonitorneneiuifvhnninfvwsiffrffowxetvnfsrwweruiirfoooeeneetuioveoiioowxwsihvowfsrefifvetnwnzfoesnnevfoitrioniehtovwhreonensnhixsifnioeuushswtnnnfnofvneeetooeenheinurowzvftsoofwwgewsunnioffwersrtttohsrfxsoooitvnnsgruieuvteznwrueiisrineoewnxooiwnrreeofnxtuieesrnvtwotfiiewtoteriifwoofieriesoteofterverrroievoitnnnoweineehiowwvsweruooioffovoonrxvsresitrsnfinnostheeffosieeewntewotwfehonvfonrwrennwxeouontwxroitefnveneessrneeiiioeufeevofnxiifxwhefevxxtiosieonoexfrnvtfntnonereioiiferoutwtweouxoeneiirwotenrexvoouiooztnnvreotweeetveefwzeiueseiittitivsenrfseeisieosinnsreifeewnffiiffisfoiwfirtevewwuforiroteiefetiexwioenshneiiooesirhnonnnoriieoioruevroesntoffnuvievtinierfroxveriutofiovweivnifuwrwniteneefoeirxtnfiovogvfoinnrnnitviwiointtwneitnhetihuetnogwvfvevtxwefooiihevxnvxefftnnihfrrieoeortsseeoetnoioeiisowixnfowioseffexvoiwtteouieoexrfeeefeioxotniintoxwtfnofvxoenneoxowwoviionefnnssoifeoxrssovteeietoiirfeeretnerhoegoiesfvenwnnovzonifetorititohrtoixsinvieivfeivueufnisfoeinetioeoeonwiustxnstenwisnownornsievevtuwnontsntuonseeeneeeeheorivoeigoxtnwessuvfviniuuivoeeoetwvnxixxrnuvfweosofseohivuvvofivowiveinwvouueoeettwoxswnervtnovtfxionosweuisooiwnewwxnruowootisinrveowieoonewroefnsietfeentxoonoxvezroeutvefnueioihhxeoeeuxxifotwneeihewuefwexnftsnonofvfveoufniegionswreotrwfeowfeexxeofxtnuevieffoeeeeivteeeseofeniorevrftenrtnrnotfworsvifuesfvofxweoffvienitnonfeiotwnofnruswfonwottoenxorntfnxeoeeveotnouwetnfteuifnoxnvnnffssxfnuwiefveowexwiioffeewuwnewroiegnwzsniiinoixxeiitstsiztoeoefeesnvfniehotxouwtsrieernouwueuovfwreernovniwofrsvenioevnwoeotnuenfxientsooneexwioueioeeneouoeouteeweiitrreuioheonteuweroewiuiooftooutfeueesfenuwvniueniniforotviewinwxnhnteeftotfnoxiruonutofvrevvsutvrtnexeovoxffoxnfeiiirevesrooueeonnwnxetfofseoioenwnoxeeiwiofxnfteottufnfowseteeeeofwnxreevnvovuurroeefefenuwtoriofxwuiestonhreeuieeoonirefieeostieeeeossnveueerosinweefrsihrnveuiwnesoeetwvietoretzxonnuseixhsrtgnofnietseoenwvonnoieftvoefotwrooetxetueinruizotwuooiooveonoiftuevvenhroerenrfhfsfifwreoewwrtsesoewtvtfueiszviewieourioeinurnsrethertivnsenswenffhnrsvosuftffetwfoenneseffrenfnewsneneeisfwfoettfsstnnweteffotnfitsriotoswsuiioeenternsvfxivtiiwtnfosrwwretsfitsoioohfveosfovtifneeieuivtnexfuoeeiewirioruzonrenvitrsovoovvterswfweiiioorewhsefrieinerenexsirivifrxrsnsexoosveexvotoiofitiexfoioiveteitroizeisiehoeeoniuiewooeisuifxtotitehoswvtfntiiwseouxnhvfnteiwvrooioisnhesowreosiswenrnseeiweisieiewfsxfoxtreeinxxixueuunhwoetosxnvrfeeewhtxeuufiowifefnnneesonfxitveeeooiivetnovtiwtneoortoioeuroosfxfewnxeeinnetnnrfwrfhvnieniwtivxteussoieeotetuooxnuseioiwxsorroeuenrentnsnvosrufioeonteofiniiexfuisseooeiownsnvutwsrfeeieeseiiwvfxrivnfhhteftuvfuexifweentfoovsooveruutoinsofeswuxofiowififnsennsrixveovowfeeueevtinvifrooftirofwvuoutnuxwifweteoiosieeusesoiirneuiusnnxiiwxontverxferuefxsnnorrenrzeruxxnsvooetfxiiseirrneetnnniioovnxieesweewfieeuftzoovnoifofwheexofxsurweesinweooxnnoeeownieiisoowteiofonewxwreoeevrswretotvnewetvxtutufeoeoieeuhnoiieuiheweiteeniftxerntsenoxrittievosineuuxesrwfeneeineruxiosiveeeenftenvinivveroxhtetfwnttgftvninhnonesteoexnoooriowoisenoofowwrtwowffttfrrtervfvufeotffhnxfeueftwwfotrifwenfrfetennisftefeiewvnniiowveooiirorwivvxftieoweeufotorsnrorveiftsirvofoontixfnooteoneeiorneosurnhnxfioiiteoieineeerwftxviiwnteotefesutreeenenvutxieiohvhzrironixovonrsiiteefeeeiiotwfrnhutouitfreiresefeofofiniifnennowoifuevteoioiixtnweeveroeeunetoeeeteesewewfxunioeufoefroreofwnrznvifewoeeouxxetoirnoinnsttwiiigrnoufrerhovfiveefreteoswsfvsrseietneewtwsionisxotortnsvfsvfriioisnfuuvifiwssfofnxienxftisxeheoffesveioiitnseowfoeniofeetfesninfvusiootetixfwewwvvxniurefwioheriiooiwfvsoveifenetroeosonhfvfwooeononfghuonesefoofnfstxeevoenieeuofuvonniueoooinizxfieexutwxvnosroorfovtoruoetenrsfexnweiinxotfnvfieiiinixitvwwweoevnhveftrrorseeeetouefiuxoewtxnuotvoieesivefrvivfweeihwwxiexnwotffnxiwerxtwnwiixufeoeenrieeoeouseftifxensioiiegornezonxenuxwotvrsiivsowffnwoinhfooxriuninioownftoewvsoeteuefxorexfeevnfinnineitwtiietiufitntoowiutthrtoiefsewsuonvveetefoeixiexovriiitnfrfeuntwwnsgiuworuosvenoniswitrernnerfvoftihonhoxwfosevirfnwftseghxrsitiirvfvvefhvvufiooerowtunewthfvfeoeoiwouvefffiftietvnsoinforoegsrsexiowefoeoufeofiinoiioifgvoiteieorenevieteeitwuuienoioxvriowoewioenrrxooseexiuzeewoeirwotozetvoreortxefuvewofeftnueisoneeivroniewuwniveheoohsinonnfnvofovxetveiifniuiwvwfieoseseisveitsghtteeeiottsivovonveinvwtteofifevwnexieevfvvefhniwiueurieefeifoevnnexwivieewnoeerinewoeefofeoohnteneeenofoewvifvixnoetvtitoiwovsifxeeueoexeeiwioeeeonxnotriveionseioxvewnfvxfsrnstwwvivvehvnesevereivfftonrvoviosethueevhientinxvefiintvisoiefeiteexfsuonotieniievereeinexrvstnoosnrwhexoffrweuehonuetoutrsxisthrrfwixueeuurvnseiefeifrtenivtonovrontesnrouenweiisreetunwonvethifewesegiifwnevfotihwinenffnxtnirgownxhifiiwfeewfienrviixnitoufeoiiwioueovueungvoxeienuffifroxfifninoxetteetioxwoxexwfevswtvonrtrsentierusuuitiotevnonootriioivoonwvxiewivwotetuevnefentxfivwxxrwivrwevuheueeuohsfesxnivonwnsonriteirwehfehtrhunioofverwssuvstssenxifneriienneretevowererfnifnnriewfwsiriffseuvvvforisxvosntetrixofffewiiuurnerninurfhsfvxonforeexoeirxtneeifvootxnioewfettfeiseninwuxtutsnrfeonewioioiufutvviweuoennowfvtieiuuunnitttxoinitnwtoiiiifxnftenoioterfofofnhofxitvvisvuinnuveeefrowhioouioisxoenevthuinfeifixesrherienftfsiuvereuooetofsiofeuhxvnhsrirtvuwxtrnisstxtnoioewixfieootovefintegtieioeftifrnnextgeeirnfhoiznesxefrvvinnosetoiftennoonnxenwwhtiuwshoxirfefeeeeneoivoifnfvwsneeotnueuusohxwvnnweosinefefeoeoeeoovnwifeunfwfooosoostiroertosfvtetxixffertxvnrnxetfnunehwsxtirnueviowuffrureufuieiiefweneritrufxxnronresussnueiwetxxiofxntvrtxinfuuesroivnntiieosoffiroohtetoroeotfirfvnhoffieuonexewsxnssrvottexiifweohfeeevtfewonotnhixeiintetzrtweeeevenwunoonfefnontnsvviunxvritfohtenuuvoioouroeieovnesoeotnsevffvueuonwneeeoutioeieeoeeitwieuiteoesesheinvouffrerinnnoonesiwreoietfxeirforotninfonnvttfnvooenessnsuhnonnoxnterifwxefrfiruwvseshrwressufifoeixtentnsneuninivnioeonwxsoofnvrenwsoifrivvewnosreweewreveetfiinhtifeoxuoienosovoffexoinwnvieeoteneosuoeeisfnoeotooeveeiitfeonifesnineeeriiiwffeovuoofrwwevioisftstwveowetwsrusefhriexnfnwewsevenneinvxoiswefesetffonioxsowwveefsneeeooerxnefxiovsvoesewiifvxeeoxxxuteofoeosinifiestoroffosswwuinroixvoeiirzeesfsfvfreffifwveiroeiisieeeifwexfewhewxuwnofftihwovevstvisnnzorefnoteneiirirfffiuonuuuveoexiiotefoneevsfonsrifueoerfonneteoostieiiitfuuoeuufveuerorxtvnnfentreoofunfieoessxosfwotonuxsfwniounwuwiwioxoierfewioirvtseutooostnifreesofinefefnnoeneesioinfhoeuieoeweueoonstutentnffoesvfoneoeixrffuivrnooeiewieenfvieeiowrroiirrseoeieiiovneoiroetetooeexovositwneuvetiotxvtentursoethwiievreeixoiieieeerwieviuefefowvewoerxeivffntriouonveoinsvotfienneietfinutoovexennvinieofxoveixvoettvxeoewsfsiisxefuegsooiniieuffoveoewfnoxefoewuisgtnessnetftifvwfeftxfwtnenzvtnrinofeteinvxuftoueisofitixrweixvviounenvsgtfnnietnufohvrfxrfesrfoixeievxtseioeieeoesseetxeioiswnuseoofsefenfeitewottoiutofssreniextevsreiesfrvwnifhfueeuitfxoneivnrvhfiwxoiiinxoneiosvwuevverehexinsnoonsreuietnoeouiiwitwfvutovitsuienriifeotirxvnieenfoftooxovorvstetfoeontexfnnnhtrvniooohtewxooruiuhsfnerexvosnnotwvvoeeowvtiiiuiviisonroxffwooetevneeeuotntofwvietotsseeusuevnosutswoeirofeietvitvfonneuouuiseriixurooforvnneeefgvurivfgxfierfwwoveoootxnefefwssvvvisoohevfnowntuniseieewunofentxieoiiseonexnxofuefeniiixttwtesioifsenevtoontreevwwrveefeeenooerensooeetovivvvineiuisfofnfrwnvfvonsefvexwvtntifueveswwfxiennrnwfegiishniinesenuweieeofonsvnesvfrwnexweetooeteuxovnrowwsenoxunfwntonithwnenfifnsrnffisnwneeirhnuvfurroewoeeeoroteoxoxifwxerrfvrnefoixeeweooiohitnnnesrnisositsesfuuisvovfretgxivvofeuefvniisveenuoreenwswtwnvewuevwurfevfennoveoixeuuuifnfronverieivsntthrorhoroxootzoonoriofisiueenoievseveffoffunvfftiwofnnessfevuerevuinenffntsooftnoieirnoioshtnnoionoewritoinfseheoweoitiesneiusrinvifnieouohexnrtsintxooninrseoexrfsniefvexoeirxorwtefoefexovewwoftetrtixfxeuwtfefsrnntoiteifiewxoioveeviftefoiirinivwnitswxifvuieonvworstofoeovoienxnwfwuesoonviuivuniwnwrrwrfiexunnerofnnxoovezvosneeeftoeowewooegoenwoevowfttniifoosneweixsoenttrtutftsreenioerroovewenfetovrrfeefwitseoeevhfifsiioeerrftxowexfrrseevtheisefonofofneifutfeseenwrsswoovetfxofvroevxnnrtviowxvtssriwinfnhoonsfsifwewornuehfnrnvsnffiohoneoufnreiertfueenifofeoxtowiiivnioreouexfhnvsiosnnsifxfoeeooofeefnneotionhieinvhhniefxtsinxeneerinterrnfnvniionsviniroheositifuisiienvisntenininxeoiifnrruineowsiveewtieieurvhuteioeinnoovreeeiexiiffieufetiiossvtnosotesxifofitevusrforunrffniintfefnwiefxwiieoufooneiiusnoiwvohoivuiufuosiennnffnztiirnxetiiiieifotviiowfiveexnoivenfnitfoinrivtvtuiffstwxvfofwfiefxfnowroofveeeftoegrwsetfixtnefoouurionitonnrvsvwouirittoeefiooietifurrnfinwfvivewxuoxxftionrnrowinuuietnenriunwfiefoutnroxtewhotinvoritnvnoeerwnffveuiorixeoefefeiiniueiowxsosexortfreevtoxnfiwvnorzsiuoiiioftevhieerxeoisneneeutfgoitwftoeiervneessoreooviuusfioonsxtroviroxoenovrnowntienooiiuitroxortsfueinuwiwfofnoevveorvweixvefnfsnronhirrfrevonevsvviioieiiviireovffeeetiutevsnowoetsfevfnuxetxtnfssswxetftueirfefnrwixixxvtoieeninxofvooiveeooxteseeveeneenexeeoxtoieinoowuofoexveeiotosexosffrnexoernuvoiniohiivinoirwntstniriieiooeeeioietrxxeshooetienxeefneotifnnevriefnnsfouvvenfeoeotisienweneexufffffwntffutnnsuennxnotinrvooisfesweeoiownvxvwroeehrirotehoveorviwenxeuexhfrxwovsivonuvifiinfweeefoonoexhnefoxserxvoxovoeeettfoinevvfesoeonffiifieneuneheeiirueowsuvtefiovinoneowxeiveieoifieoieirnrenefnetevtoteowusrweiiioisrvrewtssgoeisxwfueewneeefiseexfntnfxtiirtiweservxtsxwnixvwniruotentgtxozouooruivfniihweuireizoenssevurfiniorueneuifoeonenoooiwxxewtfeewoxifiiosoifeeunuinfxvinfuovfwiuvnxnoxsonxsrteiitnrvneewrnefsufeovorowenoouretrsgotinvveeesotoeieissnuisenffvnouovssoooitntifeiioseosrfnotnsunrxeveitviiihoivfirxroreisiwovwfsoowosuwrsssnvuoninssnxreoenetfuifnoeoozoiroxeowinieeuhvsoixostzftefwxxeuxiusowsuseoeesevefseieexnorevsniiitwoiitttixvirirorfwveovwnnttuivrivtvuovuvitnnrsrseefeewnnvoifieniusonvfosnxeivixrifttreeotefsvitoneenivisovsineototvfvfnrnivovieiixnuhteirootixousetvfvenenveoevhtoviensretiioinionifoswffnriifhixevifuennivsxeeisfnotvrofrtooerxwvheitiuoitrrnueeoereenfitsnneiugtffoeenoeswxiowfnvivfnefxwusuretvesngiuuoensruxhnvnfwietwitfixernvwfntuioiiwiftrsfoifntinotfufinwuhtefoooiitvuoivoneieowntesrfioitnwuwtoinitioeonefxufsreersonnixxenftvueoeesfiefuifxnnfverixriiorxstwstofttffonwnvnsnehfeefxtnsvootvftwixneieoiiueennofeontoefrwioefeosrwnifrtixotwerffniunxoeositeneffeevtewisisosvsenzeinovtesoxoerefssoennwiestoxfexinfftnheouxtihirexeuewoouixihwfeviffiiftietsseneweooviofonnifiohswrfoxoveesiovxxeviveoneeeoifinennvefvrfvwoxnufrfuewfruitnhsirinvffohweoeenxifxeirnoientfiuuwnitfnrfnuivvsevwneerievrtontunoneienexuwtnuuioiofttiunvxevieevierxoixosnwiofrffnsevnoiexoffsntfftnoewnsefnoieeeonxxxviuefshterenotioonfzuuristrioeniwhreevusfnnroooverneieifrvtitnnntiotuurofsfxfuerxesefrehexefniffofotnxiinexortizffwiosfusuoteusesesxvinewoeofnvseiseinnxeeofneeoeuirioiievirffiisefrftioewrutnivsoeefwexfxorhvfeensnnnxvgvsrvvxvxeeeewufuivfvsisfsveiirftxooefrfesfiifrxooerevooentrotisshnixeitftovxfveusnreietfeueieinwienuivirfwsfvneerfeioesissvvvsrtfonvnsfffiveeieonviihtgsoroewnwfohxexhnirsoowinexiiiiooieisohiunwnofennfvsswneoeeioionieoxwowitfofofriefxevnsrowieuistfnixtuxniwseiuwixotstiientfvoeorosntextonsexnsiitiffrovonionnoftevehewuirwtsfnuifiiiveoifrfoeexfsneeionnersrnioiueuoteonortiftthnoueivnrwisxuonieofervieohrfxhvnetsnnfxirernxesnotoenshofvfvoxevvouononhxroresiivreriesoeiueuffifiivtieihihurtweesfnsieffntnvvieineufnxnnoinneivseoisneifnononfteoofifsnuntwfooorrsifunoixueuenivifeofroswinhwrsrfiwswuoronnevexxvxonsvieeeteinsovtwoxieievoxfsteeeheweeensnshenvvftnoivwihnnxveeooxofierisnhnnuxeninnvoerxisersethhnztxouexutefxuxztrexrorintnivufxenwfsrxruxuxexxooneotvrsifwiuvwwnuxufefvtifsfeiitsiuxnwfietoeffernwiniteeeserswxooueonzvinefuowvneiexoestneeoigsoofwiovveoiviirwnnntoffovonieeeotwfnwvtevfevfinwworeehetfxiixtxnievoveifofnsvwooxvxerewnfeuuifiwinoouovifwvwtrfeoieiisnfvnnfsvoenenxvfftswivxowfiuherisnssrnseionsosxsfofevouneenifeefieuexfefnorrhoeetuoweveunwnevseoouevwrxnrwesnensxefeofeitfiewxveoniesiiosnwifnenesnsrnwnvfvowrxsxixfenriffoxroxrinoiweviotofiiufnwwfeotsnihexvvorinnewvoseixovnetisefrfohhveeeteoofosetteereotfwnneieeiivrfrtnihxfeeirrifuerfrusofsefesrooovritrrxxfwsevtviinfoseouofitnnewfoffowiseeitiottvegvnvvieseixoirofnerwfvoneeiiofroeronehovfrnweofworonxoexroenoeseoweoinfveeenegrsxenettxfoofieoinesftfoxnirszeoifeesoixtonntevixesevoitenfntnurvtnftiooeniueirnifftiweeienfoiixviinsuefzfwovnvwtvwfefuiefnoineeveeeiwvtoooevnisfienninoutvesfvneeisxeoeovoesfeheuriieuotugesvvninnvrusrteonifhszoofnioniefwxtioisneuzuxnfseveovvenewoevitfvnstniewtxwoefeeuoevvoiresewieivsitvixrfienooeuurwvitonretfenvowetohontnixteeufesniinseeufixxfowxxxvvvfvoeoeoertieieeeoeeftfvsoefioieirsfrseotiuexiiisefrotonoenownwufeerwheeonxwsisiienwiootiefnxtiefizuvivovereefonsenovfnofiweewoornevofieoxinnixietfuxreiinfxexevfoenvteosfrstrufvtvrvvievxiiffftoiefvhisofeonihwwftenvoieveeinftoweouefteinfeetnoffrwtiseinrnivexosunfeohotoevnoiooisegvseniivsniozieoueufroeiveoenesfsonnvnfteshnfxnionwuixooeioxiroivheiihoeeeouooixieetvvfsieeuhnegfeteeeofrifofevteeeosvoveuwietusvxtonrteiteehsnofsefetsionnfiioiiiexefiveuixiofnnsisviuxhsfxwteoftwureuninrniwexunsvivvivofeunfesxofonsxwefisxvnfnefsivoxiotsfeefiitnsefviftrtovoenifirvifoshfuxeetgreenoneoeefoxonfhofixfeiietnsstvsseenuoiweouisfftfwiiwoosreresnoneuriiunoffoftivvetiiineftznszsfisveoovinenssfeeesoseehfnfvnvnsoetvxsewerireenonwfnfftoveiesseeniztownietwtoniroosefextreifieososexxefnevivfninexeienhvehrsvnetewefowetoiinfiiuxoetfnnxirxnioesfnvieuifouwnvufiinniviwnfesnonnwwefieoxoewiuoinsreoiffsennxowovxieixiewnovfgisunfvfntrieoossonsietoeitifteovinfnveeeevoieeofeesserefoxrisfeoienitvievtifweiexitunirefurnvwnvotnfoononhxeeisehoovnieieesiwwuetufnnvefeeeoeinevefoioeitivovniriiinxosivfeftoiieofriofouoweuxiofifveeeenizeoxxveeneovwtonifeooiiirinixnifosfvioisfvthevxnnvessuzffntsiwiinftiiwoeexsrrneoxiveonxeoeeonifseesfxtsisoiiooiefoieotntenoswhfsxnssewseuftfeeifevsftvrxnosfweionxeonfoevofixesvsefoovoveenfveiwonofsetnefeenzofossronwftnefeetoetvnitrivsoinrnuwgiinssveiteseoosfnveteinieefnwttefveiewiiuenfusutfiniseiowfiiginietreusvenoewwofvwitfeefrxrtfoivioeninoooivsfetoinxfswtinoviroiixtefoextooriewoiesofesvvvesitixernofixnxteeeiehohtiivteexoosnooeninieuxeefvowenortofwefooieefexfrenisezenswivotsusfffivoneofvwninnirfiesouwrowiiftsvewwuottoenveeiriivevienfvnfnuwxenwtffstevhovitnioiirnonierivorhnfniewnnviiffixtefteenoiiosfeotfnxwrivenoinfnxnnieeffheeeweetvexnrofivxortuftfsnetxwwtfioisirinnniseeiehvofueishtwnoooftooewisesoirnusoixeinioxreseuenrsstrevvxsiwfvsvoenwveisuveeizusxftwvioiiveoerweusfehefouwvunwixieeixfwuovinvintvontvenvrueieoooxnufsxevevxrevnivnievrirnrevxewxforviitieoiiosnwfioisixuixefeersxhfvrrrioeeeoeooveuienoofisoxefrxenviorfneoivifeeostnfeeorswnxouvnevieorxwioehoifiohieftufttrxnexfitruoneheexerwrtnseuenvtesrioeveuivfseoouvevnrttetrofswxuiexnetoefionoinegeooionevrevteivevoewefxreierseevnffrfeworhooiiesxeiguwesnsoosoeeotvevreoznfhoiwxouefrthiiixiiiifiosiunrxfniexstsoeeeiofuforveitsnnxinhsoxevnvxvnxwngxneofvefrroxfvsvfnisniooxefuunisvtnuvexxssieseiiexefxuuiostfoewnnoofonovoufuwovsxovnsihiixvnieosieewistsiovwuooweesefounefruwtoifhvsestennrosruetreeftfionwtenvtfeiesfeovfeeiiivtifsufirfnwevgtreeuiixievxnonfninvoiiutteioesrtvuovvrvfsowrineoteiiisuizvtisitoivosenieeiiieoittreoowfsoeffnenxosowvrnvswnnensfexeuwwwutwuefhffxienviontiwfffneioeonsrenxixeihouisiifsowoneexrewvsfgfswfssfotinsxneienfeeinerweoetesiosofeeonoswuitfvneeisetitofiohgninvwinuhrieownveiufetxvensifvftenrieoxorfnsfoexnoesenxnotfeovxorioeouvviexnsvvvsnfiooeuoeiineosuiohxiixgteivifwihofriooeonneoeuioxftonvsrwnixftsvviinfroeirfiexeiowfineseneiowwtexoeoxsnieufneuiteifxifxieorwesznnwnfeoueiiuvexotexrnosvxinrruveeivfnieniexsnnfiexeeeresorvveoonwwissiuruxeishxveiwtonxfeivnnviffeixifowtositwwxrfieiisitewsoesoeinioihxewfiovffxsevvowifftoeieeneinttneixiioteiixeeiofiottoofeoxferneettnftesrnuievfouoruntiiernoheosroxtonfwhrhitinoowonfiresounoirwiwixttfssxwieoerertfiswsfnenvnewofeeotnuuirntvutesxtixfuonexivfrxfrutfthoiiiovunfeonsiinxieiooioeuwuoeznixvnffgeresueiereiweofusttttnxnrvnenfxzieosftoiuwesoxonnnoixsixivxvixferststiefiztitoiefsxnnetuxnnunfoeooioiooiewrtwtionntisoretrovniseonvivrofnutioewriieiefnwrsretuevogfeuiffworiioointtnsinewehfsuoesneonituinvetofineifienteivtiwereosffontstxesfitxwieirotsxrriihuiexuvuffnrovhnthsfotwxhiniiornueoritntfvexsonfusfetnuofeirfewnovwfifwsorioossezftwvfuteuieivxnenffeitreeetnvfetfrrnvoeervnvsefeofroooenovorvfoiinuinrxoeeeoisixevwxtzisewiesoeiiiuisoenexitsriosizwsiwseroiiixunfreontxeixineeffnrtiiioeonsoswuiioerruwxinevewxiiienfnifnvrxnvnnnsnvexeuoneoniunvtenxtorfeewsvnevxvoesiosuvtonvisweuotioeooeonnfvoforontnooxnxrfiuihisufnitsnffxteeeoenwwinwxettiioirevuveitfvovovnueeifxffwfieoeiisovenoieioonfinesoooeiiournfvooveethnnenvsienovvunseeitiffxsirfeonttwofeeinfnevriuvtneenneeriirhxwfnoeewvosrvssnvntiossxtvevwniwfoxsoiuefiuwieeiiirioteewegovineorvtnoinnuffrvtnntensxofvvruivrifftfieintruooioevsivrwiesoxisnrsrgosfxoenosxirnoierowiivfexxeeiseiierfiefnoefiineiveitfxwxiewvtwwueneinssvinstnvfniwxnivooeeeinvvniegvsvsioxineeeofeeneennvhitxeetioseezgnsuutoiveeitenstsnirreseroevenwissfinnnzirierrevuwerfefiiovfnniiioirirninfrnfnefnnvftvuerveooftoxniiifoieiviufrfnoxwxnornfhxeeiexenineifnnvtftsweeiuioosoioofeveexixonuifuinrowioeixeitrxxnswioeonissefnxeueinexnvesioitonevieirzvorexoeooxeisiioiieofvnxwiinexfieiifsneoifituefrixwnxoefvuiifeeffsiirfuntiiegieoisiufiwiexnnoifxfroeffoforvooivvtoeieesxttnuiwffeeexueiesotifiefvofniiueoixrnitsnefosvroxieenerowsevnwvswiexnisxnseouxtrsntrrooeeofriofofwivnseivrhvteiiooexetusrvnvtffentieenfnvxtnfofifwueffnfteeheoeiweiefuniehnneituevieiiiereowifvonuohvvrtsofeeovifoiixsefuwsesierfiensohewsfefsvinvifxfeietiwofinnoxssfofeihnifivtfniewiensenrvoreefovxeeioofewifniieovwrernontrrfrtsiotoixnvifivnrousouwsiiueitsfeihnfwrxrntsttviwxunwoexuufuoerofeerfnheenvneiitfverefiirseeostehteswufesfuetexivfeevfofieesfeineevfvnosehfoiovwnsirnerevieouxfuuniefnfhuowontninnowoernffxistrisrioeueiofevursxeewotosxfertxseoinsnnefithnenxrotevtieonnxeoissitieuoeniteefnnxenxnesnsfevefffxniteertitnoreenfzexfsxxfsivenxoeevswntxeiwvfoiheetneuohnsnoerefwexewoieneeooviiinhrfwxvoeixiinnitfowifrieetvivxtxsvfftovneooexoheeosnwroorowiisuuensrsiesnrviwiirevifeffeirnrfnniwtxioonoetfsvwuswewseiushneseionnisrvxrixtnrfuettfeitutheetfexnereiforestournefweseiuwtinfwsrotiofiennunseuuiteuuhfxosnwrwvewwserfnriifieoswihstsefvvnexnwieueesewsuooniixuootirsvusxwfeunoinxexesnnnfvvizrrwiffteoenfevetxnsegffienexsotenextiefuxienxtfsoioeuoxeehifixteftnioeintnrieeortivwfitxtieenefetxntioixegsinrosxxrnnoofeovooihterfonvusiissnuseentoeozosffueixwnhrfewiioinnorovennvhxvxsrvswntosinrfiierieneriitonxivtsevesnoroviwxitnfeniivefesietovsxuiwnnuwsfosiiufuifiiefoeetveofxnvioefrsvnnfeeeososnvxnentxrvxoithunrnrsfwtoeveeeoeexoiueiftiveifofsrriesneovntsviueztettxinoiohiinfuiewefeeeufvvweiinfeseveeiifuhhoeivisevuueoeifinhrfxerxitxfnnssuvrvfentesteuswesnsffeoinevniirtseonisvivowtfivxifiotifiiexuwirsfnvtnrfeeiiwotnuixosovoniovnvzvttseuesvvnseirsewffrofxfeetsotvwoefneinsosoifiixeniowtrnnxeewfefvefoohxewfeuxxnsveewievozseoofhtonfnniiutvzxsxnroeofexovewxxfsetenxresoivsfiexfigfevoexifeneotsrfsefovsneusiewronwivuwveeionsrroffroerfzweetthniefuoetivhinnexieinewnfoovnirhweoetntitfxooneouvrifneteeniestfntfooxveotinresusoffiiivsisffvefnefnnffwiiriiixrensifxvuiiievewinwxvexosrrnoeoitterniseeuteftnniezeneewviosefnevionsnnfosirenxeneeniverioefxoerwosvowvonseteoenxinftsensetvisnongoennfutixeorrtevnxitxnowutnswiftviinwoewtzwnoeetfniiwvioonveevwrewxsenoiiennrfeiteisfezefoivififitxeorunrttiinxsniotinotretsegeoniiieeotouroiotovwfiirxiiwieoexrxnivetfinienieotirnwfxeevftnoohstsveoiniionsosireiethrexoeuhhvofeeexeoizvtvsfvvvnsvtovevooiierhoniwoiiougexwrneowshriixinoofvnunientoiixsfiioestoiintoixioruesrwuonenineieefnnuiveoixixesetofexhoeiuenteezoeerinsiiniosusfvftvxoevewitfinfvwtnreixoftnnnososiovionfhveoxeueoonxoxiioitnfoorntnennnefffsesnnovuffrxvunreexewretuenxwonisofnfxxwvtrnerwriitvnefexfteennixvientxunfonoottoetsenwofefvitwnefiweotvsoohxwseiiftwotftinxsnniefenfixeieoveeooonoofeioenevesoifiieoirrorxoxteneionvtntiersoixftnxvsihoxuosowrevoteeniigwsinrisvostniniiureoniuxoezfexvnosniuxttwitroeeuseexosoiofvunottxfforsnotreesewoersnireoivsrusvsixnrvoferrxiuretevenevixooetheiotwiustxwneeitxoivereeeoofevhnoinetneeenxfnwfiofxniwnieoeseixotrveiexorfsxwwivxnooovfrstoiehsonfvsrioififttfeehxuoeveieogntoieerseeffieiixxxoieotreioirxfseisxxvregneossostiooxuousteetoxfwoxiofngtniivnrnentnxfsonnseuivisevsevooiieerxivovxeixoeotuonifwtisewtsooeetuxoennsewxsifxttviietxovxoxxnwsoeteiieefevtnoiteihswsihoxoiofwveooiwoosowrninfntenrnisvieiiseniwxvuxixewroeeoniufeuvnuetwosfeseeeristovssiroeftrxirwwrfwneoffeieeeoeneornseniveoineernfswfttntiinntowtnowxtttetenonoeexsneineeneivnvtxisiwefefnehonoosrunsxssvsonnxtneteeseuxeoixrnsrwfnowsfioiiissioweesfonsroinxoseriuoninfnwuwuxtoixsfnnfevrreoitfoefnnrswiixfeofswwivzrniixfoutveoesennewvenefnvfexitfosvtefsvwvnstvverifnorioehvtevooroiiexeruwfsrxewunxiwxwnithfnvtotefoieiewiwisvfwtoxoiniivsenxnssoteeenfnxrteievseevoeueonnuteesfwvrnfeiuuwoirrvnnvovoenownfotoeooeewrwueseeroosveoeefsfnnixvieivrrnnfnnfeutenewxofnensoonvofxxfeinireinverooevsoeiirrosoevuenotrvitoserstivsfvoixvvfitenieftfteinoienesuvhnxvtotwoeorffeneoinoxtiwioitoxneneeontwxvsvnetrveoesieethixfeiossohiixoonsoirxtiesrwzwnesneotritwonsuneofiiofxiiueeoiinirvfuutrhtreiieeeevfeeswrvovovuiwioeefnwffooefoeiesiiiiinneveoxxoxxonttfxfeionnesgiinnwfniiefwntswteonwiiwterioosnenrinernvfheveirfsvnerwfesnwfvseifxevnrieveesetvhsisrnxfeistewxstfooniwruxsoifnhoirinsrneersirgfneoirithvieeftuunevvroerueustnfuewethonxtivuvitvfwneisttioxioinetivstfforwfinerxioieexrixtrneneeeeieviiineisevexewrhwiittixrnwvoiesseoesweortfvfxitofuwvivninwsnweeevfronotneeozwrwrxoneereneveiiofieffvfriivvweifnuiwefoxgefrfneiiiownsnnerisoxssvreoeotfnfxxirfwreeivioswovnnxxvneeoieoostievixnseooioessvooiuruxfoueroooessioeiiffwrwnfenwhfeihitfennoiiwoigsorrrwwotftxevxwvenxfvroovotssixtwoieuotioisnonrxnsnnvesntofvioesioovretonwitvsvxthovfwevfuwofrxnfexxnitronxwnsnueooxeeoxoeetveotwiooinioxviixefiovreinioivxxsnoeihuefotsrfrtieeetesxisoeneuwennfneieisnweneonhiefniihorfswefrnoseowrveeoieftowsoffvxnxtveervewoesvessesiswiwrntnwifwuuierfoiivetxeetetretnfnnnwrofensxwnwxseseesixeifereeihonffwsowntvihntioeorvovinoieeetsisfftgoofexfnwsivxievofeisiesretvezuiwifxihevtievtonniweixniriseessoooonxvirfitxuoeoeowvoontiuowterxeruofxetisuffroitfieeefxefoovtfxwxioxniifvxounvxeioiifoengixutsstnoixeifeivifxnrvsnftsfioritefosiozowenetexsnowwieonrfvnoxfueownrstnifsievvortoofinereeonneieseiinreetesrfrxxniexvvnnsfxxtwoiefwxniowiursrrttreunsforisernsvffinvexrrftseionixiesnixfoersentiteioiiierneienfssrinirxvwtovtifurneseneonnrsonsvrwrnoievieiowwvootnnfootiouotvtnnfeeefexrnhrninsoriounneoveiteessnutnxhtsovtfusxonovnrsuxrsxviinioonvoinniuussfrtennnonvevtooffrsonxvvrtrwovetesoteonougnuefnnioosfoexrxtorfnxeewteeiiiofvotenxseeezwrwitxhhxefrtfiinofvhiioeftsittwifhxvoeoitntfswwnuoititreeoeveseeriowviifnzexeefwvefssnsxesinonixxxsieerefxenfeeuiuefifwvxoniowenetorsnoxuosoinwiieweineeifesxewieixrhfteoiiszooiexfevioefwxvxetehnuuifextvfoiitntioviorutsowenoitrxnntxwsrvnouseiwifeiteeeennneeoffeoevvtuihtiwnuifniefioorxrfnfexsevntxnnsewosetutnxseseextonvvsiuxieivnoniiihoxxtextssirwowvoxfsvsirhoitrwvxifeoenoesoixenfruivtniweoxivsnfieniorwnesfviiservoeooeeonenehsnirieixofsieeiirfnoeioxfssftfwtiousiofwooeinsfnrvfvteeseievfffvowsesveviixxiretiieoewhvrovowienenvifreowsfvouteeiohsnxnirffnfsxeoeeiestfviiosoiohntwneifhfrfeniefnieeeixioiifrsoswxesofnwsietentoioifneniseuioeswroeiowwexvtwtnstrfgoeigiiwwiierwiexneefiswsxeewiuhvsefvsveoneietsxeiirxsiieuiefeissrieuoonreewtiooiwinrfuvuvsoersssovnvxonwiifvrroeoxuseoovfietxoteenvivtosofiorxxsernsexwnfeefihnseseoifrewfofossirieeivostwiugnoeeoifresntvixooieunfwuxrnxwivntsnfsnefeeorxnusrwexiefovvweneietfxvfeneesoxxeiioeiuutxeetexiuooiifeuieswnfiensewftfotsnuenrfoerszxrufvexxnoexniownotrnoftifnvrivinnnsxiioxosorwihsfninieetrfgefntfeeswfnrxssevseveennfnifgfetefhuitsenfnxsiefsooneivseoeoohesixesnotwwwivixxintfnixonftswwrerweetoxxnixosowuronenivhrsfutenfnuweeenioehneifntseeexreifeeftetwneexixeiesfennonnneoetoionfvnrenxeutiexunvoxwvriifvsoxsuxxnfvuxoiniroinesuioirusfurvstesrioewseixneowoineeveinevivtoeinsevtiwwswioriosfviveifrxoriisvioofosnornivfvonftnixfefvfwofnofevnfevvhnsieowiseonxrrsnxrifesioixxsuveotieixenffreoivxgofofnooroiisoteuuevefvisisfonoovtsefxneezsivvfrheeiteisuieivfnixfneounvfixfnervenfntesnsteototieeoiotvnvervxerrwewowsisxisrsuieuexxoeexoistiwirvwtfefnresfiifteoefviseussivexixteoewivouefoeuxfoexeotfiiiuevvewziexneefnonifveoihworexeunietnvvxnosoeoonrieovevoinoisvwionfrnvtfifeeooeiietintsosswenoeeetnvxuxfzrufeeinwwevwoiuusuueiviswfeznrsereuowsuwsirnznsvxfoeueruioeofeeeowsvetounixsvvvewiouxueeewevennxftwvfivniwufsxuoenewifeoeeewxtfxeoswvtisveovsivurwexfxeeotfxnwxtsvvsftesroeeiowfitoinunfnvwnwnswsiefivxttrveseivoxiinuntefenuewitfftxnrwsixsoheinxeenvetoownxeefneenwsoeoxoofoweohfrtexnsnviigiienefvxnveioiitttooviiooohsiiiwtsieinowretixniieorfofoeenrufhzxisisvonsnxeexxrwnxreresensitfeiohxihwesnuouwevtntioteofexxixorviioooxxtxoiwosuinvesisnethfneniwtxfiiuxihroiwvteentuoxoeseeesfweevixivvoeivetnswxishooioinexevxrreirerfntvfioresowursznoeneoeesihsvensenonswuvsfnsionovvrwfsffneeeixnofiuxtrxvnninhnstiownnefnnheiweieioixiiseeefsvvxeoonuofrniswxzuneeffteewhwofuesrrttseeoeonnoeesxirnfefigeveefiififvnuoinssfervxrrenesewniienoexfsiiiivvnineisinuwwsinsiveeewhisuxnineuufxrorntxtnxwewneexeewioeetoosieirxoinexieiefevofnstoonfstftfiseinexsneiesinetveftefeefneoeoxsfvxuironuvtevnesfvtwvrfvnfsooeeosfeeiotteoserfvwfivseinuoietofvhnxhsifonioowoowvosiffoseooefeeotewuortioionrrrussvvosxieiteniifroenuvotuonnseertnsfoxnweffsongswoxreievtxfrnhrtsssnoievfzwnnwiixrerfixrtioesneweeioooeevvenveuooigtwroxioontriwsewuewroeniorstneftrvoeeeseiirnxuvivxnfewxnnnsfnfuneeoevveiffnswviiwwieinnteoxvierefeioxnwonnxotrxenetxertsnfonntwtfwxxessreiesxoeixihxihnsnnxieuioifvnvirnsrusnnnivennihixtiinfeiiinonvgsiseiuieoesvreffwunutionsernstrfrvweivovtzritieoserevoonowfiieieviriivviinneouxfvnuvvrzoiieovnfwizeefwxnfwoewsnoweesftefhtheioiinnreeesxosxreexeevexieefeseifoifgooeensnfiifwwvteeeoonnexwrvioxossfwhzinixsoieevnnuvertrriiwvriwnuntftvnssesuntneifeeefewseexoierwsesnixxofitrhewxiiieizeiieiiissosnnxoeoewneviffwnoeoesontvnfgifiiiiwewiieiifevrvtrzifnovfennifiiinwfsefnssonvweooiweeoufwxeifisiieewevifovtwofvseirnonovrfninennvuuiofifsusnnnoivonrwesueeoirueiofuwrvvnovtiuxtnevtnoivfwivnsistsouffiouvtvnenuowvvsthrfniiwisxriinnefuresosisenfnsoreefuinixotioienoiwntovrewfiisonenetosovrfvotwenvsefowrvounvxihwnefunienrowenueeesieoeffsntvexnsnssoeniieevennxtunovevnevexenieifiinotsifofvissufewxotreeisnivxnfsveveefnosnhwefxxwniieiuevxutneiwtowniieoeieifwevtetssenovfnvtoiriseonooxtininwffewoifehwvitnwefwieonvefetveeetftesszefxioerftusrwxeefsverooisiwswsrinnsrunrrsetsvewrevoiesgoxreioieowveovoteiriefwnesoiiuinoneooeswsosiieuwtooeisueofefeeneuxovntxrnrxwworweoesffoeneoiossroefinoreritseeeixwisfisetrnwtfvrifesieereoinwxrifsostfveerxivrosofeiftfieisonvnnvwuieeiieixxsoieentsfnootuwvronsuooenrornxenunfntetnssftieeniineuvixifeeooihfvsfeeoinwereetinsoufsenefiisfrsfeneewtxeotnntooofexhixswenfevftvonixiifufintniviievfeoowveotvtexronrtesegtwuvsneennifxineohxwrionvtionoeinewnxrviseetegsoetexnunweienevveiexnoofietoofxvoirirothrfiennorienoneffoefsfveewwsfiweoetfeorooriieviwnotfutnnrsertrviwooefeoworefeureoenvwxrnneuntvevinoeereeoenrvnheexsnttrsnniveerenvueuovosxxuooefnesfrextenieonrsxeetowehnonsiieeietrenhfxeeresveiringefttowuxifovseooewiiuostixvsewrnevvogoiufirssfwnfisonroeewsvvoonexwvfowintetvnsrietxrexrvfwnxwfnwneieennrreffenureefitoeffrotxoxoenoteeenviootffoxrevifeoeiivtrrtvwiofsovssnonuosiiirstxtewfieotovsioetntinoovnitvfsoeeswvotnisinitoeietefvtxfofixsgefevwotixvoeswfvnferevfooioffseeoiifnefvwguioevofxxuovfvsuenotviuvxseeshfoxroixifvoweoesoxesiwtefvneueoeeiievwnvoitnrvtifeeifoeninvertxteriieeoxinuouhuiesvinihisfrnfoxxfwintrevnenfeeeosonffnxxentioitrretihiewszowoffeensrooeooisrexnsxeesfxufereirrntiiunfhzrhvxszeenoeniferfssnisnonxreoifsooxeinizfsefeefeioxvnoitiiefsninfifoeoenwueoeuingeeeusvnwruooxwsfoenwitiriitivsefheifoewxetrofswnoexntvvfhsiwtsonoooniwexeenfiiuitveeexwifntnniwvrufeofoioneevgsifoowoifevihoteuuongtfiuewesonvxiweusvuoooeiineehsfxssneisnoootstshfuirtsfierveiivewfniooioeenornufewevofinonifxvfwuveuutwnwtgnoetssoitiwninnwfnseonxehrxefesoiefveswnixtexefxexxxtenieefonhrrsisitfewnfwvwreswetexsusxhswvtnotwwiwvfeoiutiinrnoxxneoiueenwsiireehtvxwsreeeefvteoeievnnxfevforofooiwefvovwofvosxrntevnoeeitieoetniiittoiheieethfffxvrfeonnsrfeossetoxnvfehvnteiiioufixhftssniohnowfnfovnhtftiitureieonfeiesnwvnotuehutefsenntvwvuoffoeeewuitnorienrereeeefsvsextsonifgxszefiiertefhnroteiisxtoweifeexeoerevesriuxexosowiefffenoiienenoosieeotugunevuenfffsiivftnennsiivxnrxwfsoiestoeuvoieonnxivoeeneeivrveoefxtnoefxtsnrxsnhoiennoveitnoefxffieievieuseeorenhixoievinenertetfrnifrserosueeenwtwhiunhowtfinenssisnwrtefrnnxgxnetieiovwosoiiiirinevwovffwnrniiifivwtvtnesievnortennnffnvwsosorxentiofteehirirsotreoosevoxsvtewinovnenoinnvxwfxfxnoeeiinvefvnoxnznxfonniwoeeuetnerernfootrfeooiievireoseixotriehefonfvrhootioforvfweetfexfrninftntorotfeitixfeeneheiwinnrinretoetiofwinevxehrrxoswvtewroniiuxienevoifunouxefonfotsfioiteuietesfeinfvrsiivetrnxevoveirtfoevunseuufirvsnurteuoweuigreetovtvieonnenifsiooonnefntnnoovssnveennssvvfuntiiefixiixnveiseiseewinnowffneiwvviesittontesevxvtnvoieoentienenserwnniefxitnwevneiwonxeieeroeftoeiovtteriniennsiesnfrnnnxniofnheeeevnsnonnfwxetvnntnehiestfrioiwxnuwexseinfvoeftnvhiersexfisneieziexnexufwxtehenwieiuenowneinfirsxuovoieifrusieswieefofiwfoousiuevtsivvieeseoxtvtstuorfrfexeneneeieooneivieeewefntiviuhtvfeftooxewuwintfxeorosiiwonrefsxnonoiwenrstrnernfenienowteverfseeofwiwrteetweofoxseervwteethsivvnoviweowoieerhionvfoosoeowfiwhrsufooirweeewniuunrwsenrxoeoxfeoevnitiztreufxtuoettvoiefioeheineoeetitnnefntxwvixfiiiunexvniewxoniviftfifxivtenhvsftufninnosuoeeoeoionvennneinerieresostostwoovzwohoftreioeoixoxeininessneeetnntetieoieueiwrruieooowfsixeienefexehfuxxnhisosoeevevioinfoexniooefeinivfnviioneinofoeinsffifvesoooeonovtrterhewiisfvnoteutesetuooueiiufhnoitvotisovsieieovtoiwsieueieeosefouitevoertnotfoxonoetrnxoivhixisweoteonreiooffhivfesfwshueriwonowenseveoxftsivefihoeoofnoewtiorofuveofrnsinewvwtesvnrtsitiuouoeexweoenveiwothoowoeowexiioovvenierfxineiwxwxerenoetenieenenninsfineetnfnnsffvetsiitoffsssneseeeevftxoiiorwforznfestunurinvofevoeefenertovefiiofftreiooevxisuseosroeexsefisrsovftihfnsxttuforenoueosnonterounonfssntouieenesurewfieiovtsnoxteweenienuirfvitnonenwoeortnwiseutxestwiwuseetuefeuetwoevorinefnixsiseenstttovewireeooererotfiveoieinfnvvnevwsnsioxtvsnsterxereftnteeoeesfrxwneerfwtenrxifroeenexfonnsxovouiuroefvtofnneioeeessetntewvtxeifhvonewvnrinheneeeoifxwiffheszsnxonwifsfnfeoxxfeoiiefsfoxsxneeuieteveeeseieswoefzeeeofxownfeneeerxnufxextxwwierfiiifwineterwffifriurvoeofvwtrnronfrsrfieotfsseovsusfioexnsrenfnfsiiiwoeeiefifoxnenixnniefiesuxuvifeuwxuosfieosxfhnwfnfvvfexeinnntixwfohsoeenrheueieeeveotnvzoiwuoroxnoefsfovesiifewxxnsztevienvorovfseineofoofsvfooeriniooivessrxwfonwnveeuvsssiewxorsioveintnetetoiifsteeienooonisvnfrnttwieheinfifssorxnovoetonrroeswiutiosiuerwteitfweeroovfserneenofiionfiioteoievzeviieffoesowiwnoxoetssotteeeuvfneixnuzvetntivfrgeuerfxnxnifeftrntirvenefvsuivivnwiwfrvnoixxtrveefisizorrioeeitohesvuefnsotvtfrooowxfvonrovetosxiroeevoifoxfseevnfnoienetoenrsiineoessireoeroiifousfxerftswiievuiwrneisfinwonigfofwoxoiexionieroooetoiveefrwtovsfisexeonrnneiofuvewnusounnoffofvrieizoosnteswnfnofesifteuivnueisivrfxoofrutnitxwwnftxiuwxsouwfooxowiefruosooxiixvtrvotegntiosrveffhroieitxnetnfsiuifesfionevfsfoirefiviffittrwwevturooixfvinevhrtswneetritvteosttesineueoovoiexotufivutuoooofiuvwhfsviuteehoffofewoviwefooeivexuofvhiieienwfuwsehsriuntienrweoivoxvwouevsoifeeoixsnioeoenunofiexrixuexteereoovvurvvoreorftfneoftisesefowswrtfoorrxhifuivoveetfirseixfvituriwvwrreiioifevtixeionwffeeouisotfvfniwnfowrifftvneitooeevsvwwsoifueioxosutronenesnssiritoifffexfrnitetvrxveeufvtserzviseoneroossxrnuitteonnouinxrtiseettefrxvfxfinfoiieeivwrweenwerswxioefeetioivfwewtswvewvinfteeueofioohueseoootoorhufxownexirnoiiifisffiteofhxsreffeotiftiwieoxxevihusffsosrniviuwuxoexerosoievisiioeoetowtxostfrsfeefussoffwithihxtwrfifioououwensrisewesofroefuseresooxsooxvxtsixtosxgxeofieuoovfiwieoxuoenissieoouuifhuvuvwierrrsfosioieiseoioivioeefefonsfvxiifvweoorotuevofvotroresewffuieunseftvoexereofetneuifsxoviettxoxesxrvriesvvvfnwroioifvrhreffxueunentvueeeoivfoweeeeoosufrefrxextifexxfhsutefifwsievihirvsioiiiotffersetieweeioreiirfviefiriseeifvwottfvsweuneffhvfoewtioteoiooerxswioesivfetxeixvsoetexwxsuoorenrvoiortrefifixsfvoervvtrevnonsoovntoiievfrxsoivesvivisrewshuoetennsxexeoiorvgwoouuefvifstfriirveseriooeevvioxrzvtttieifhrohvtirrsiwvrfrosfwxxxfwsextesiriveiseiirsunwoinefwiwsexirusoowwufoesruxiesuvtssvwoufeneifenfwegoiovtvtooixifreeoonshnxwrhorotofewrnfwvenoxwweheshehetxvvneoittisvsewvsiwnwwuofovvwexiofonnfetrivwrftsowrhotroiuerxeiizsoeswxofehuviotenevthoessuofntoterneiowrtooreeivwezooivtvfuwifinisfttrvxvexunerewioretfrtissoevienuievofssueowniewwvovrfeeoxetvissxouxunwueivhewoorrtixseinofiexneentieeooeifetiiexreftfifeieoooeonfoetiihreninfifhtfesfoosisoirtruwfneesuwixxerxiriofiizgxiteixwrofneunvfvffxtrsreexfxirnhrwhsofnwveiiffrxfferifwiwnnxnsutvxffuevoeeerivtfthiiouieweiheeiioneusfenittvofxezrnehxotenehoewwiifiitenwisenvifxveeinitervifrseooiiveesutifrvewxntvneiifnfeoiootesnwutexihefgwniteohhoorfefixsefoexoioewvnveisviihsweofstrsioennfiwisiffixoeeuoisnfoeenusorfofenxowrrwiniifvuefortnrenfwwereetfrfterueeoetxnesieiwioxefteventrorrvieuovoxtffiretioxursnzovovoxtoefenevrrusesehtifnntvovsftuoviuoihvniiivxtoooneoeuufeixfeefxosowfnuetsinitfvsoonxsfxnxetheiuvifeoxosvniiesvifexsxvotouienvhneseofeoitovnuvfsisforwuifteserifrtxohnnoourhffexfoeotvtfsohofooeoxwrwstteinxursftsioesinnrsfrnitfiwieiosefsiooorxnevioieiisoeeuoeoinfuvtfefxwonttnvixorxtoveifenwsnnooovsiefwfieoeeeeneweexohiuefsvtxenrereestiofsofsofnettorzroxsiffewofseoeoxoiitostonvfvsxtverisovtvewwrsetehetwfeiuffniirirteeifrxerrvsgnteosettieiuoeveuwosvftsnwririueeoeeznnoovinwitetnvtvfuwsifieveionfoixrnhefrfefvevotohveievntxxeeuohfsvstterefuefxtoeoruoxxeeioosoevfifnfinosvivoteriivfosuehxtrnoteeiifosovfoentnnstrtsiievoiixooiinfosuvsfoonvxtvuoozouefeieetxhrnxixoxiereoenitefiehetenowviooihfxeeewivioofinftxtowoorneffeeonofoxeeihvniewxuiivrsinvrihftiiiihvotnwfeoonfweeinoriotofiwuxsuiinsieewhwnefeehwisrieenvieeriteiiiroesfeferofvwovfixowfwefevovrfvxeeeootnfrioreititrfvohoitoreiovotvfnnvnseusfxeoeinrotueooeuvxfoeofosxstetxsftonrvtesftfsovexotxronvieifeeueewiervuoeonuesxefienowoonfvntoufvewvietiiioorsnrhnfevefoowoixeefesexsttenwieuttnxoouootfiweofeoooxneutfiveonshxovwioetoisnxftohrusoeouhtfnevhxsfifuesfnihoofwexsovnerweoiezwxeueixiwvfwoeorffeefxrvexenrirviuueiruxovoeneioxuuiiuotttfvtxvsnevsisxftwestffnnsxonofnoifitoeiwintieisevsesgieorxonooxvhofusfuoteofonhvevovsenienoeerfoisffiiuinoeehnenoxtwnxwevwtufteusftvnsntxooesiweostvoerxitoxsnesefunioivnsesvzfofrgsirerefowfnsvvsferiveiofrofwssvneowoehoofiriouefovesvfeuonrwnineseuhfirvtetnivnefxeevwiofrhtenzeivoneinntwxevituostfeiowsweniivitnntooenerssoeriuoeitftfwevxruusuvowurexosvutowrffxoeneswoeifeoiioevehotuwuewnnwwintiuntfontovifffwfofivvenevhrrenfixeieweusevrfzohuovrtisooeexioonorioorvtfnnfteseirtuefestsweouxtxvvnrefwoitiefooveihsnweswnxniewffeexnrnvhsxeenhtizeiitzvivifeefoxwosifwfxeiixwosunttesvvrxrwesxioeeeexferohetoixtvtosereitfnerrwshuonhwoieeiofeeiioieefrfeixhexevnfosestfnnxvnsresowvwnunennietiiwototeunrissouoveitotnewroofeotfoonewfewxifvuefnrnnvseosffexfisetttuvrweovuzittiivtowvoiiofvoeevefxoisnwxneoooiiweeetwfvituofxvttfteweeifnreiisooreitotniwenonetffenvhzteviioinnoioixewnieefuwtfvnireefrfxennooesvuuwnesisvxfxfeoneonfoewoosneoooiiefuxnezxfvvffsteierovoneiunewrfeeiouuwtwooifiernetsooefxifoxeinzrieuieeunrienvnnroeioxvxweonxsennsxeuhzerwseewnoefsinnuxvovswevieuffsexieeewoifoennssioooufniixesensxnftegetnuiieeivfteoervxiinxeinfetsnnetixnewifnvieinnotroerwfessuooiewfrtooxofetivrnsohosoeseweseffnrxsneeintrfneoieioeixwneervvxxitewufiieeniieftsffisefooitxeteeenurvieefhieorneivvheunewevfxuoonfseffofrfnnfxwxsueffnirsiftirnfiwonoonoieonzeivheivfsxuoeerrtoiooeeszvweitnevstrxfeurtsrooevfiseonhnetnnxoinfeeixtvtrfrziieeenineevenftwirnixwuenwenteiiieufoffounfesfehuxiveinoifiofwfeiiexnfixivnetoixrieiueeeofoowteenrtoxunoenffisnefivvnosfvwewwexvoveovitrvnxoofnireovvfeeihierexntihinresexwnontteexrveeefeoneissswwfnsooisxevwfiivosonuusuowenusoennowiersorirrefufixueewofoeerioevnoinvietfnifxvsssnntoxeeeteniiiiehoefixnereiioefeoufifieesioseveviinwriiosoiuivesrnrwwsxoeiiiotnxnevsxrusoeeefiorihsoewxoifeuietexsweixfifowsefeifofowtfnevnfhvfotwsnevwifon"
	println(OriginalDigits(str))
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	ConstructMaximumBinaryTree(nums)
}

func TestGetAverages(t *testing.T) {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	GetAverages(nums, 3)
}

func TestFindNthDigit(t *testing.T) {
	println(FindNthDigit(2147483647))
}

func TestLargestSumAfterKNegations(t *testing.T) {
	println(LargestSumAfterKNegations([]int{-4, -2, -3}, 4))
}

func TestSuperPow(t *testing.T) {

}

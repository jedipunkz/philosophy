---
source: "https://arxiv.org/abs/0704.1020v1"
title: "The on-line shortest path problem under partial monitoring"
author: "Andras Gyorgy, Tamas Linder, Gabor Lugosi, Gyorgy Ottucsak"
year: "2007"
publication: "arXiv preprint / cs.LG"
download: "https://arxiv.org/pdf/0704.1020v1"
pdf: "https://arxiv.org/pdf/0704.1020v1"
captured_at: "2026-06-18T11:15:23Z"
updated_at: "2026-06-18T11:15:23Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "György Lukács"
tags:
  - "現代思想"
  - "マルクス主義"
  - "物象化論"
  - "西洋マルクス主義"
status: raw
---

# The on-line shortest path problem under partial monitoring

- 著者: Andras Gyorgy, Tamas Linder, Gabor Lugosi, Gyorgy Ottucsak
- 年: 2007
- 掲載情報: arXiv preprint / cs.LG
- 情報源: [arxiv](https://arxiv.org/abs/0704.1020v1)
- ダウンロード: https://arxiv.org/pdf/0704.1020v1
- PDF: https://arxiv.org/pdf/0704.1020v1

## Obsidian Links

- 研究動向: [[研究動向/ルカーチ・ジェルジュ-現代研究動向|ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代思想]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[物象化論]]
- 関連分野: [[西洋マルクス主義]]
- 関連タグ: #現代思想 #マルクス主義 #物象化論 #西洋マルクス主義

## Abstract

The on-line shortest path problem is considered under various models of partial monitoring. Given a weighted directed acyclic graph whose edge weights can change in an arbitrary (adversarial) way, a decision maker has to choose in each round of a game a path between two distinguished vertices such that the loss of the chosen path (defined as the sum of the weights of its composing edges) be as small as possible. In a setting generalizing the multi-armed bandit problem, after choosing a path, the decision maker learns only the weights of those edges that belong to the chosen path. For this problem, an algorithm is given whose average cumulative loss in n rounds exceeds that of the best path, matched off-line to the entire sequence of the edge weights, by a quantity that is proportional to 1/\sqrt{n} and depends only polynomially on the number of edges of the graph. The algorithm can be implemented with linear complexity in the number of rounds n and in the number of edges. An extension to the so-called label efficient setting is also given, in which the decision maker is informed about the weights of the edges corresponding to the chosen path at a total of m << n time instances. Another extension is shown where the decision maker competes against a time-varying path, a generalization of the problem of tracking the best expert. A version of the multi-armed bandit setting for shortest path is also discussed where the decision maker learns only the total weight of the chosen path but not the weights of the individual edges on the path. Applications to routing in packet switched networks along with simulation results are also presented.

## PDF Text

arXiv:0704.1020v1 [cs.LG] 8 Apr 2007
Theon-lineshortestpathproblemunderpartialmonitoringAndr´asGy¨orgyTam´asLinderG´aborLugosiGy¨orgyOttucs´ak∗October27,2018AbstractTheon-lineshortestpathproblemisconsideredundervariousmodelsofpartialmonitor-ing.Givenaweighteddirectedacyclicgraphwhoseedgeweightscanchangeinanarbitrary(adversarial)way,adecisionmakerhastochooseineachroundofagameapathbetweentwodistinguishedverticessuchthatthelossofthechosenpath(deﬁnedasthesumoftheweightsofitscomposingedges)beassmallaspossible.Inasettinggeneralizingthemulti-armedbanditproblem,afterchoosingapath,thedecisionmakerlearnsonlytheweightsofthoseedgesthatbelongtothechosenpath.Forthisproblem,analgorithmisgivenwhoseaveragecumulativelossinnroundsexceedsthatofthebestpath,matchedoﬀ-linetotheentiresequenceoftheedgeweights,byaquantitythatisproportionalto1/√
nanddependsonlypolynomiallyonthenumberofedgesofthegraph.Thealgorithmcanbeimplementedwithlinearcomplexityinthenumberofroundsnandinthenumberofedges.Anextensiontotheso-calledlabeleﬃcientsettingisalsogiven,inwhichthedecisionmakerisinformedabouttheweightsoftheedgescorrespondingtothechosenpathatatotalofmntimeinstances.Anotherextensionisshownwherethedecisionmakercompetesagainstatime-varyingpath,ageneralizationoftheproblemoftrackingthebestexpert.Aversionofthemulti-armedbanditsettingforshortestpathisalsodiscussedwherethedecisionmakerlearnsonlythetotalweightofthechosenpathbutnottheweightsoftheindividualedgesonthepath.Applicationstoroutinginpacketswitchednetworksalongwithsimulationresultsarealsopresented.IndexTerms:On-linelearning,shortestpathproblem,multi-armedbanditproblem.
∗A.Gy¨orgyiswiththeMachineLearningResearchGroup,ComputerandAutomationResearchInstituteoftheHungarianAcademyofSciences,Kendeu.11-13,Budapest,Hungary,H-1111(email:gya@szit.bme.hu).T.LinderiswiththeDepartmentofMathematicsandStatistics,Queen’sUniversity,Kingston,Ontario,CanadaK7L3N6(email:linder@mast.queensu.ca).G.LugosiiswithICREAandDepartmentofEconomics,UniversitatPompeuFabra,RamonTriasFargas25-27,08005Barcelona,Spain(email:gabor.lugosi@gmail.com).Gy.Ottucs´akiswithDepartmentofComputerScienceandInformationTheory,BudapestUniversityofTechnologyandEconomics,MagyarTud´osokK¨or´utja2.,Budapest,Hungary,H-1117.HeisalsowiththeMachineLearningResearchGroup,ComputerandAutomationResearchInstituteoftheHungarianAcademyofSciences(email:oti@szit.bme.hu).ThisresearchwassupportedinpartbytheJ´anosBolyaiResearchScholarshipoftheHungarianAcademyofSciences,theMobileInnovationCenterofHungary,bytheNaturalSciencesandEngineeringResearchCouncil(NSERC)ofCanada,bytheSpanishMinistryofScienceandTechnologygrantMTM2006-05650,bythePASCALNetworkofExcellenceunderECgrantno.506778andbytheHighSpeedNetworksLaboratory,DepartmentofTelecommuni-cationsandMediaInformatics,BudapestUniversityofTechnologyandEconomics.PartsofthispaperhavebeenpresentedatCOLT’06.
1Introduction
Inasequentialdecisionproblem,adecisionmaker(orforecaster)performsasequenceofactions.Aftereachactionthedecisionmakersuﬀerssomeloss,dependingontheresponse(orstate)oftheenvironment,anditsgoalistominimizeitscumulativelossoveracertainperiodoftime.Inthesettingconsideredhere,noprobabilisticassumptionismadeonhowthelossescorrespondingtodiﬀerentactionsaregenerated.Inparticular,thelossesmaydependonthepreviousactionsofthedecisionmaker,whosegoalistoperformwellrelativetoasetofreferenceforecasters(theso-called“experts”)foranypossiblebehavioroftheenvironment.Moreprecisely,theaimofthedecisionmakeristoachieveasymptoticallythesameaverage(perround)lossasthebestexpert.Researchintothisproblemstartedinthe1950s(see,forexample,Blackwell[5]andHannan[18]forsomeofthebasicresults)andgainednewlifeinthe1990sfollowingtheworkofVovk[29],LittlestoneandWarmuth[24],andCesa-Bianchietal.[7].Theseresultsshowthatforanyboundedlossfunction,ifthedecisionmakerhasaccesstothepastlossesofallexperts,thenitispossibletoconstructon-linealgorithmsthatperform,foranypossiblebehavioroftheenvironment,almostaswellasthebestofNexperts.Moreprecisely,theperroundcumulativelossofthesealgorithmsisatmostaslargeasthatofthebestexpertplusaquantityproportionaltop lnN/nforanyboundedlossfunction,wherenisthenumberofroundsinthedecisiongame.Thelogarithmicdependenceonthenumberofexpertsmakesitpossibletoobtainmeaningfulboundsevenifthepoolofexpertsisverylarge.Incertainsituationsthedecisionmakerhasonlylimitedknowledgeaboutthelossesofallpossibleactions.Forexample,itisoftennaturaltoassumethatthedecisionmakergetstoknowonlythelosscorrespondingtotheactionithasmade,andhasnoinformationaboutthelossitwouldhavesuﬀeredhaditmadeadiﬀerentdecision.Thissetupisreferredtoasthemulti-armedbanditproblem,andwasconsidered,intheadversarialsetting,byAueretal.[1]whogaveanalgorithmwhosenormalizedregret(thediﬀerenceofthealgorithm’saveragelossandthatofthebestexpert)isupperboundedbyaquantitywhichisproportionaltop
NlnN/n.Notethat,comparedtothefullinformationcasedescribedabovewherethelossesofallpossibleactionsarerevealedtothedecisionmaker,thereisanextra√
Nfactorintheperformancebound,whichseriouslylimitstheusefulnessoftheboundifthenumberofexpertsislarge.Anotherinterestingexampleforthelimitedinformationcaseistheso-calledlabeleﬃcientdecisionproblem(seeHelmboldandPanizza[22])inwhichitistoocostlytoobservethestateoftheenvironment,andsothedecisionmakercanquerythelossesofallpossibleactionsforonlyalimitednumberoftimes.ArecentresultofCesa-Bianchi,Lugosi,andStoltz[9]showsthatinthiscase,ifthedecisionmakercanquerythelossesmtimesduringaperiodoflengthn,thenitcanachieveO(p lnN/m)averageexcesslossrelativetothebestexpert.Inmanyapplicationsthesetofexpertshasacertainstructurethatmaybeexploitedtoconstructeﬃcienton-linedecisionalgorithms.Theconstructionofsuchalgorithmshasbeenofgreatinterestincomputationallearningtheory.ApartiallistofworksdealingwiththisproblemincludesHerbsterandWarmuth[19],Vovk[30],BousquetandWarmuth[6],HelmboldandSchapire[27],TakimotoandWarmuth[28],KalaiandVempala[23],Gy¨orgyatal.[13,14,15].Foramorecompletesurvey,werefertoCesa-BianchiandLugosi[8,Chapter5].Inthispaperwestudytheon-lineshortestpathproblem,arepresentativeexampleofstructuredexpertclassesthathasreceivedattentionintheliteratureforitsmanyapplications,including,amongothers,routingincommunicationnetworks;see,e.g.,TakimotoandWarmuth[28],Awerbuchetal.[3],orGy¨orgyandOttucs´ak[17],andadaptivequantizerdesigninzero-delaylossysource2
coding;see,Gy¨orgyetal.[13,14,16].Inthisproblem,aweighteddirected(acyclic)graphisgivenwhoseedgeweightscanchangeinanarbitrarymanner,andthedecisionmakerhastopickineachroundapathbetweentwogivenvertices,suchthattheweightofthispath(thesumoftheweightsofitscomposingedges)beassmallaspossible.Eﬃcientsolutions,withtimeandspacecomplexityproportionaltothenumberofedgesratherthantothenumberofpaths(thelattertypicallybeingexponentialinthenumberofedges),havebeengiveninthefullinformationcase,whereineachroundtheweightsofalltheedgesarerevealedafterapathhasbeenchosen;see,forexample,Mohri[26],TakimotoandWarmuth[28],KalaiandVempala[23],andGy¨orgyetal.[15].Inthebanditsettingonlytheweightsoftheedgesorjustthesumoftheweightsoftheedgescomposingthechosenpatharerevealedtothedecisionmaker.IfoneappliesthegeneralbanditalgorithmofAueretal.[1],theresultingboundwillbetoolargetobeofpracticalusebecauseofitssquare-root-typedependenceonthenumberofpathsN.Ontheotherhand,usingthespecialgraphstructureintheproblem,AwerbuchandKleinberg[4]andMcMahanandBlum[25]managedtogetridoftheexponentialdependenceonthenumberofedgesintheperformancebound.Theyachievedthisbyextendingtheexponentiallyweightedaveragepredictorandthefollow-the-perturbed-leaderalgorithmofHannan[18]tothegeneralizationofthemulti-armedbanditsettingforshortestpaths,whenonlythesumoftheweightsoftheedgesisavailableforthealgorithm.However,thedependenceoftheboundsobtainedin[4]and[25]onthenumberofroundsnissigniﬁcantlyworsethantheO(1/√
n)boundofAueretal.[1].AwerbuchandKleinberg[4]considerthemodelof“non-oblivious”adversariesforshortestpath(i.e.,thelossesassignedtotheedgescandependonthepreviousactionsoftheforecaster)andproveanO(n−1/3)boundfortheexpectedper-roundregret.McMahanandBlum[25]giveasimpleralgorithmthanin[4]howeverobtainaboundoftheorderofO(n−1/4)fortheexpectedregret.InthispaperweprovideanextensionofthebanditalgorithmofAueretal.[1]unifyingtheadvantagesoftheaboveapproaches,withaperformanceboundthatispolynomialinthenumberofedges,andconvergestozeroattherightO(1/√
n)rateasthenumberofroundsincreases.Weachievethisboundinamodelwhichassumesthatthelossesofalledgesonthepathchosenbytheforecasterareavailableseparatelyaftermakingthedecision.Wealsodiscussthecase(consideredby[4]and[25])inwhichonlythetotalloss(i.e.,thesumofthelossesonthechosenpath)isknowntothedecisionmaker.WeexhibitasimplealgorithmwhichachievesanO(n−1/3)per-roundregretwithhighprobabilityagainst“non-oblivious”adversary.InthiscaseitremainsanopenproblemtoﬁndanalgorithmwhosecumulativelossispolynomialinthenumberofedgesofthegraphanddecreasesasO(n−1/2)withthenumberofrounds.InSection2weformallydeﬁnetheon-lineshortestpathproblem,whichisextendedtothemulti-armedbanditsettinginSection3.OurnewalgorithmfortheshortestpathprobleminthebanditsettingisgiveninSection4togetherwithitsperformanceanalysis.Thealgorithmisextendedtosolvetheshortestpathprobleminacombinedlabeleﬃcientmulti-armedbanditsettinginSection5.Anotherextension,whenthealgorithmcompetesagainstatime-varyingpathisstudiedinSection6.Analgorithmforthe“restricted”multi-armedbanditsetting(whenonlythesumsofthelossesoftheedgesareavailable)isgiveninSection7.SimulationresultsarepresentedinSection8.3
2Theshortestpathproblem
Consideranetworkrepresentedbyasetofverticesconnectedbyedges,andassumethatwehavetosendastreamofpacketsfromadistinguishedvertex,calledsource,toanotherdistinguishedvertex,calleddestination.Ateachtimeslotapacketissentalongachosenrouteconnectingsourceanddestination.Dependingonthetraﬃc,eachedgeinthenetworkmayhaveadiﬀerentdelay,andthetotaldelaythepacketsuﬀersonthechosenrouteisthesumofdelaysoftheedgescomposingtheroute.Thedelaysmaychangefromonetimeslottothenextoneinanarbitraryway,andourgoalistoﬁndawayofchoosingtherouteineachtimeslotsuchthatthesumofthetotaldelaysovertimeisnotsigniﬁcantlymorethanthatofthebestﬁxedrouteinthenetwork.Thisadversarialversionoftheroutingproblemismostusefulwhenthedelaysontheedgescanchangedynamically,evendependingonourpreviousroutingdecisions.Thisisthesituationinthecaseofad-hocnetworks,wherethenetworktopologycanchangerapidly,orincertainsecurenetworks,wherethealgorithmhastobepreparedtohandledenialofserviceattacks,thatis,situationswherewillinglymalfunctioningverticesandlinksincreasethedelay;see,e.g.,Awerbuchetal.[3].Thisproblemcanbecastnaturallyasasequentialdecisionprobleminwhicheachpossiblerouteisrepresentedbyanaction.However,thenumberofroutesistypicallyexponentiallylargeinthenumberofedges,andthereforecomputationallyeﬃcientalgorithmsarecalledfor.Twosolutionsofdiﬀerentﬂavorhavebeenproposed.Oneofthemisbasedonafollow-the-perturbed-leaderforecaster,seeKalaiandVempala[23],whiletheotherisbasedonaneﬃcientcomputationoftheexponentiallyweightedaverageforecaster,see,forexample,TakimotoandWarmuth[28].Bothsolutionshavediﬀerentadvantagesandmaybegeneralizedindiﬀerentdirections.Toformalizetheproblem,considera(ﬁnite)directedacyclicgraphwithasetofedgesE={e1,...,e|E|}andasetofverticesV.Thus,eachedgee∈Eisanorderedpairofvertices(v1,v2).LetuandvbetwodistinguishedverticesinV.Apathfromutovisasequenceofedgese(1),...,e(k)suchthate(1)=(u,v1),e(j)=(vj−1,vj)forallj=2,...,k−1,ande(k)=(vk−1,v).LetP={i1,...,iN}denotethesetofallsuchpaths.Forsimplicity,weassumethateveryedgeinEisonsomepathfromutovandeveryvertexinVisanendpointofanedge(seeFigure1forexamples).
PSfragreplacements uuvvFigure1:Twoexamplesofdirectedacyclicgraphsfortheshortestpathproblem.(a)(b)Ineachroundt=1,...,nofthedecisiongame,thedecisionmakerchoosesapathItamongallpathsfromutov.Thenaloss`e,t∈[0,1]isassignedtoeachedgee∈E.Wewritee∈iiftheedgee∈Ebelongstothepathi∈P,andwithaslightabuseofnotationthelossofapathiat4
timeslottisalsorepresentedby`i,t.Then`i,tisgivenas`i,t=Xe∈i`e,tandthereforethecumulativelossuptotimetofeachpathitakestheadditiveformLi,t=tXs=1`i,s=Xe∈itXs=1`e,swheretheinnersumontheright-handsideisthelossaccumulatedbyedgeeduringtheﬁrsttroundsofthegame.ThecumulativelossofthealgorithmisbLt=tXs=1`Is,s=tXs=1Xe∈Is`e,s.Itiswellknownthatforagenerallosssequence,thedecisionmakermustbeallowedtouserandomizationtobeabletoapproximatetheperformanceofthebestexpert(see,e.g.,Cesa-BianchiandLugosi[8]).Therefore,thepathItischosenrandomlyaccordingtosomedistributionptoverallpathsfromutov.Westudythenormalizedregretovernroundsofthegame1
nbLn−mini∈PLi,nwheretheminimumistakenoverallpathsifromutov.Forexample,theexponentiallyweightedaverageforecaster([29],[24],[7]),calculatedoverallpossiblepaths,hasregret1
nbLn−mini∈PLi,n≤K r lnN
2n+r ln(1/δ)
2n!withprobabilityatleast1−δ,whereNisthetotalnumberofpathsfromutovinthegraphandKisthelengthofthelongestpath.3Themulti-armedbanditsetting
Inthissectionwediscussthe“bandit”versionoftheshortestpathproblem.Inthissetup,whichismorerealisticinmanyapplications,thedecisionmakerhasonlyaccesstothelossescorrespondingtothepathsithaschosen.Forexample,intheroutingproblemthismeansthatinformationisavailableonthedelayoftheroutethepacketissenton,andnotonotherroutesinthenetwork.Wedistinguishbetweentwotypesofbanditproblems,bothofwhicharenaturalgeneralizationsofthesimplebanditproblemtotheshortestpathproblem.Intheﬁrstvariant,thedecisionmakerhasaccesstothelossesofthoseedgesthatareonthepathithaschosen.Thatis,afterchoosingapathItattimet,thevalueoftheloss`e,tisrevealedtothedecisionmakerifandonlyife∈It.WestudythiscaseanditsextensionsinSections4,5,and6.Thesecondvariantisamorerestrictedversioninwhichthelossofthechosenpathisobserved,butnoinformationisavailableontheindividuallossesoftheedgesbelongingtothepath.That5
is,afterchoosingapathItattimet,onlythevalueofthelossofthepath`It,tisrevealedtothedecisionmaker.Furtheronwecallthissettingastherestrictedbanditproblemforshortestpath.WeconsiderthisrestrictedprobleminSection7.Formally,theon-lineshortestpathprobleminthemulti-armedbanditsettingisdescribedasfollows:ateachtimeinstancet=1,...,n,thedecisionmakerpicksapathIt∈Pfromutov.Thentheenvironmentassignsloss`e,t∈[0,1]toeachedgee∈E,andthedecisionmakersuﬀersloss`It,t=Pe∈It`e,t.Intheunrestrictedcasethelosses`e,tarerevealedforalle∈It,whileintherestrictedcaseonly`It,tisrevealed.Notethatinbothcases`e,tmaydependonI1,...,It−1,theearlierchoicesofthedecisionmaker.Forthebasicmulti-armedbanditproblem,Aueretal.[1]gaveanalgorithm,basedonexponen-tialweightingwithabiasedestimateofthegains(deﬁned,inourcase,asgi,t=K−`i,t),combinedwithuniformexploration.Applyingtheiralgorithmtotheon-lineshortestpathprobleminthebanditsettingresultsinaperformancethatcanbebounded,forany0<δ<1andﬁxedtimehorizonn,withprobabilityatleast1−δ,by1
nbLn−mini∈PLi,n≤11K
2r
Nln(N/δ)
n+KlnN
2n.(Theconstantsfollowfromaslightlyimprovedversion;seeCesa-BianchiandLugosi[8].)However,fortheshortestpathproblemthisboundisunacceptablylargebecause,unlikeinthefullinformationcase,herethedependenceonthenumberofallpathsNisnotmerelylogarithmic,whileNistypicallyexponentiallylargeinthesizeofthegraph(asinthetwosimpleexamplesofFigure1).Inordertoachieveaboundthatdoesnotgrowexponentiallywiththenumberofedgesofthegraph,itisimperativetomakeuseofthedependencestructureofthelossesofthediﬀerentactions(i.e.,paths).AwerbuchandKleinberg[4]andMcMahanandBlum[25]dothisbyextendinglowcomplexitypredictors,suchasthefollow-the-perturbed-leaderforecaster[18],[23]totherestrictedbanditsetting.However,inbothcasesthepricetopayforthepolynomialdependenceonthenumberofedgesisaworsedependenceonthelengthnofthegame.4Abanditalgorithmforshortestpaths
Inthissectionwedescribeavariantofthebanditalgorithmof[1]whichachievesthedesiredperformancefortheshortestpathproblem.Thenewalgorithmusesthefactthatwhenthelossesoftheedgesofthechosenpatharerevealed,thenthisalsoprovidessomeinformationaboutthelossesofeachpathsharingcommonedgeswiththechosenpath.Foreachedgee∈E,andt=1,2,...,introducethegainge,t=1−`e,t,andforeachpathi∈P,letthegainbethesumofthegainsoftheedgesonthepath,thatis,gi,t=Xe∈ige,t.Theconversionfromlossestogainsisdoneinordertofacilitatethesubsequentperformanceanalysis.Tosimplifytheconversion,weassumethateachpathi∈PisofthesamelengthKforsomeK>0.Notethatalthoughthisassumptionmayseemtoberestrictiveattheﬁrstglance,fromeachacyclicdirectedgraph(V,E)onecanconstructanewgraphbyaddingatmost(K−2)(|V|−2)+1verticesandedges(withconstantweightzero)tothegraphwithoutmodifyingtheweightsofthepathssuchthateachpathfromutovwillbeoflengthK,whereKdenotesthe6
lengthofthelongestpathoftheoriginalgraph.Ifthenumberofedgesisquadraticinthenumberofvertices,thesizeofthegraphisnotincreasedsubstantially.Amainfeatureofthealgorithmbelowisthatthegainsareestimatedforeachedgeandnotforeachpath.Thismodiﬁcationresultsinanimprovedupperboundontheperformancewiththenumberofedgesinplaceofthenumberofpaths.Moreover,usingdynamicprogrammingasinTakimotoandWarmuth[28],thealgorithmcanbecomputedeﬃciently.Anotherimportantingredientofthealgorithmisthatoneneedstomakesurethateveryedgeissampledsuﬃcientlyoften.Tothisend,weintroduceasetCofcoveringpathswiththepropertythatforeachedgee∈Ethereisapathi∈Csuchthate∈i.Observethatonecanalwaysﬁndsuchacoveringsetofcardinality|C|≤|E|.Wenotethatthealgorithmof[1]isaspecialcaseofthealgorithmbelow:Foranymulti-armedbanditproblemwithNexperts,onecandeﬁneagraphwithtwoverticesuandv,andNdirectededgesfromutovwithweightscorrespondingtothelossesoftheexperts.Thesolutionoftheshortestpathprobleminthiscaseisequivalenttothatoftheoriginalbanditproblemwithchoosingexpertiifthecorrespondingedgeischosen.Forthisgraph,ouralgorithmreducestotheoriginalalgorithmof[1].ABANDITALGORITHMFORSHORTESTPATHSParameters:realnumbersβ>0,0<η,γ<1.Initialization:Setwe,0=1foreache∈E, wi,0=1foreachi∈P,and
W0=N.Foreachroundt=1,2,...(a)ChooseapathItatrandomaccordingtothedistributionptonP,deﬁnedbypi,t=

(1−γ)
wi,t−1
Wt−1+γ
|C|ifi∈C(1−γ)
wi,t−1
Wt−1ifi6∈C.(b)Computetheprobabilityofchoosingeachedgeeasqe,t=Xi:e∈ipi,t=(1−γ)Pi:e∈i wi,t−1
Wt−1+γ|{i∈C:e∈i}|
|C|.(c)Calculatetheestimatedgainsg0e,t=(ge,t+β
qe,tife∈Itβ
qe,totherwise.(d)Computetheupdatedweightswe,t=we,t−1eηg0e,t wi,t=Ye∈iwe,t=
wi,t−1eηg0i,twhereg0i,t=Pe∈ig0e,t,andthesumofthetotalweightsofthepaths
Wt=X
i∈P
wi,t.
7
Themainresultofthepaperisthefollowingperformanceboundfortheshortest-pathbanditalgorithm.Itstatesthattheperroundregretofthealgorithm,afternroundsofplay,is,roughly,oftheorderofKp
|E|lnN/nwhere|E|isthenumberofedgesofthegraph,Kisthelengthofthepaths,andNisthetotalnumberofpaths.Theorem1Foranyδ∈(0,1)andparameters0≤γ<1/2,0<β≤1,andη>0satisfying2ηK|C|≤γ,theperformanceofthealgorithmdeﬁnedabovecanbebounded,withprobabilityatleast1−δ,as1
nbLn−mini∈PLi,n≤Kγ+2ηK2|C|+K
nβln|E|
δ+lnN
nη+|E|β.Inparticular,choosingβ=q
K
n|E|ln|E|
δ,γ=2ηK|C|,andη=q lnN
4nK2|C|yieldsforalln≥maxnK
|E|ln|E|
δ,4|C|lnNo,1
nbLn−mini∈PLi,n≤2r
K
n p
4K|C|lnN+r
|E|ln|E|
δ!.Theproofofthetheoremisbasedontheanalysisoftheoriginalalgorithmof[1]withnecessarymodiﬁcationsrequiredtotransformpartsoftheargumentfrompathstoedges,andtousetheconnectionbetweenthegainsofpathssharingcommonedges.Fortheanalysisweintroducesomenotation:Gi,n=nXt=1gi,tandG0
i,n=nXt=1g0i,tforeachi∈PandGe,n=nXt=1ge,tandG0
e,n=nXt=1g0e,tforeache∈E,andbGn=nXt=1gIt,t.Thefollowinglemma,showsthatthedeviationofthetruecumulativegainfromtheestimatedcumulativegainisoftheorderof√
n.Theproofisamodiﬁcationof[8,Lemma6.7].Lemma1Foranyδ∈(0,1),0≤β<1ande∈EwehavePGe,n>G0
e,n+1
βln|E|
δ≤δ
|E|.Proof.Fixe∈E.Foranyu>0andc>0,bytheChernoﬀboundwehaveP[Ge,n>G0
e,n+u]≤e−cuEec(Ge,n−G0
e,n).(1)8
Lettingu=ln(|E|/δ)/βandc=β,wegete−cuEec(Ge,n−G0
e,n)=e−ln(|E|/δ)Eeβ(Ge,n−G0
e,n)=δ
|E|Eeβ(Ge,n−G0
e,n),soitsuﬃcestoprovethatEeβ(Ge,n−G0
e,n)≤1foralln.Tothisend,introduceZt=eβ(Ge,t−G0
e,t).BelowweshowthatEt[Zt]≤Zt−1fort≥2whereEtdenotestheconditionalexpectationE[·|I1,...,It−1].Clearly,Zt=Zt−1expβge,t−1{e∈It}ge,t+β
qe,t.Takingconditionalexpectations,weobtainEt[Zt]=Zt−1Etexpβge,t−1{e∈It}ge,t+β
qe,t=Zt−1e−β2
qe,tEtexpβge,t−1{e∈It}ge,t qe,t≤Zt−1e−β2
qe,tEt"1+βge,t−1{e∈It}ge,t qe,t+β2ge,t−1{e∈It}ge,t qe,t2#(2)=Zt−1e−β2
qe,tEt"1+β2ge,t−1{e∈It}ge,t qe,t2#(3)≤Zt−1e−β2
qe,tEt"1+β21{e∈It}ge,t qe,t2#≤Zt−1e−β2
qe,t1+β2
qe,t≤Zt−1.(4)Here(2)holdssinceβ≤1,ge,t−1{e∈It}ge,t qe,t≤1andex≤1+x+x2forx≤1.(3)followsfromEth1{e∈It}ge,t qe,ti=ge,t.Finally,(4)holdsbytheinequality1+x≤ex.TakingexpectationsonbothsidesprovesE[Zt]≤E[Zt−1].AsimilarargumentshowsthatE[Z1]≤1,implyingE[Zn]≤1asdesired.2ProofofTheorem1.Asusualintheanalysisofexponentiallyweightedaverageforecasters,westartwithboundingthequantityln
Wn
W0.Ontheonehand,wehavethelowerboundln
Wn
W0=lnX
i∈PeηG0
i,n−lnN≥ηmaxi∈PG0
i,n−lnN.(5)9
Toderiveasuitableupperbound,ﬁrstnoticethattheconditionη≤γ
2K|C|impliesηg0i,t≤1foralliandt,sinceηg0i,t=ηXe∈ig0e,t≤ηXe∈i1+β
qe,t≤ηK(1+β)|C|
γ≤1wherethesecondinequalityfollowsbecauseqe,t≥γ/|C|foreache∈E.Therefore,usingthefactthatex≤1+x+x2forallx≤1,forallt=1,2,...wehaveln
Wt
Wt−1=lnX
i∈P
wi,t−1
Wt−1eηg0i,t=ln X
i∈Ppi,t−γ
|C|1{i∈C}
1−γeηg0i,t!(6)≤ln X
i∈Ppi,t−γ
|C|1{i∈C}
1−γ1+ηg0i,t+η2g02
i,t!≤ln 1+Xi∈Ppi,t
1−γηg0i,t+η2g02
i,t!≤η
1−γXi∈Ppi,tg0i,t+η2
1−γXi∈Ppi,tg02
i,t(7)where(6)followsformthedeﬁnitionofpi,t,and(7)holdsbytheinequalityln(1+x)≤xforallx>−1.Nextweboundthesumsin(7).Ontheonehand,X
i∈Ppi,tg0i,t=X
i∈Ppi,tXe∈ig0e,t=Xe∈Eg0e,tXi∈P:e∈ipi,t=Xe∈Eg0e,tqe,t=gIt,t+|E|β.Ontheotherhand,X
i∈Ppi,tg02
i,t=X
i∈Ppi,t Xe∈ig0e,t!2≤X
i∈Ppi,tKXe∈ig02
e,t=KXe∈Eg02
e,tXi∈P:e∈ipi,t=KXe∈Eg02
e,tqe,t=KXe∈Eqe,tg0e,tβ+1{e∈It}ge,t qe,t≤K(1+β)Xe∈Eg0e,t10
wheretheﬁrstinequalityisduetotheinequalitybetweenthearithmeticandquadraticmean,andthesecondoneholdsbecausege,t≤1.Therefore,ln
Wt
Wt−1≤η
1−γ(gIt,t+|E|β)+η2K(1+β)
1−γXe∈Eg0e,t.Summingfort=1,...,n,weobtainln
Wn
W0≤η
1−γbGn+n|E|β+η2K(1+β)
1−γXe∈EG0
e,n≤η
1−γbGn+n|E|β+η2K(1+β)
1−γ|C|maxi∈PG0
i,n(8)wherethesecondinequalityfollowssincePe∈EG0
e,n≤Pi∈CG0
i,n.Combiningtheupperboundwiththelowerbound(5),weobtainbGn≥(1−γ−ηK(1+β)|C|)maxi∈PG0
i,n−1−γ
ηlnN−n|E|β.(9)NowusingLemma1andapplyingtheunionbound,foranyδ∈(0,1)wehavethat,withprobabilityatleast1−δ,bGn≥(1−γ−ηK(1+β)|C|)maxi∈PGi,n−K
βln|E|
δ−1−γ
ηlnN−n|E|β,whereweused1−γ−ηK(1+β)|C|≥0whichfollowsfromtheassumptionsofthetheorem.SincebGn=Kn−bLnandGi,n=Kn−Li,nforalli∈P,wehavebLn≤Kn(γ+η(1+β)K|C|)+(1−γ−η(1+β)K|C|)mini∈PLi,n+(1−γ−η(1+β)K|C|)K
βln|E|
δ+1−γ
ηlnN+n|E|βwithprobabilityatleast1−δ.ThisimpliesbLn−mini∈PLi,n≤Knγ+η(1+β)nK2|C|+K
βln|E|
δ+1−γ
ηlnN+n|E|β≤Knγ+2ηnK2|C|+K
βln|E|
δ+lnN
η+n|E|βwithprobabilityatleast1−δ,whichistheﬁrststatementofthetheorem.Settingβ=s
K
n|E|ln|E|
δandγ=2ηK|C|resultsintheinequalitybLn−mini∈PLi,n≤4ηnK2|C|+lnN
η+2r nK|E|ln|E|
δ11
whichholdswithprobabilityatleast1−δifn≥(K/|E|)ln(|E|/δ)(toensureβ≤1).Finally,settingη=s lnN
4nK2|C|yieldsthelaststatementofthetheorem(n≥4lnN|C|isrequiredtoensureγ≤1/2).2Nextweanalyzethecomputationalcomplexityofthealgorithm.Thenextresultshowsthatthealgorithmisfeasibleasitscomplexityislinearinthesize(numberofedges)ofthegraph.Theorem2TheproposedalgorithmcanbeimplementedeﬃcientlywithtimecomplexityO(n|E|)andspacecomplexityO(|E|).Proof.Thetwocomplexstepsofthealgorithmaresteps(a)and(b),bothofwhichcanbecomputed,similarlytoTakimotoandWarmuth[28],usingdynamicprogramming.Toperformthesestepseﬃciently,ﬁrstweordertheverticesofthegraph.Sincewehaveanacyclicdirectedgraph,itsverticescanbelabeled(inO(|E|)time)from1to|V|suchthatu=1,v=|V|,andif(v1,v2)∈E,thenv1<v2.Foranypairofverticesu1<v1letPu1,v1denotethesetofpathsfromu1tov1,andforanyvertexs∈V,letHt(s)=Xi∈Ps,vYe∈iwe,tandbHt(s)=Xi∈Pu,sYe∈iwe,t.Giventheedgeweights{we,t},Ht(s)canbecomputedrecursivelyfors=|V|−1,...,1,andbHt(s)canbecomputedrecursivelyfors=2,...,|V|inO(|E|)time(lettingHt(v)=bHt(u)=1bydeﬁnition).Instep(a),ﬁrstonehastodecidewithprobabilityγwhetherItisgeneratedaccordingtothegraphweights,oritischosenuniformlyfromC.IfItistobedrawnaccordingtothegraphweights,itcanbeshownthatitsverticescanbechosenonebyonesuchthatiftheﬁrstkverticesofItarev0=u,v1,...,vk−1,thenthenextvertexofItcanbechosentobeanyvk>vk−1,satisfying(vk−1,vk)∈E,withprobabilityw(vk−1,vk),t−1Ht−1(vk)/Ht−1(vk−1).Theothercomputationallydemandingstep,namelystep(b),canbeperformedeasilybynotingthatforanyedge(v1,v2),q(v1,v2),t=(1−γ)bHt−1(v1)w(v1,v2),t−1Ht−1(v2)
Ht−1(u)+γ|{i∈C:(v1,v2)∈i}|
|C|asdesired.212
5Acombinationofthelabeleﬃcientandbanditsettings
Inthissectionweinvestigateacombinationofthemulti-armedbanditandthelabeleﬃcientproblems.Thismeansthatthedecisionmakeronlyhasaccesstothelossofthechosenpathuponrequestandthetotalnumberofrequestsmustbeboundedbyaconstantm.Thiscombinationismotivatedbysomeapplications,inwhichfeedbackinformationiscostlytoobtain.Inthegenerallabeleﬃcientdecisionproblem,aftertakinganaction,thedecisionmakerhastheoptiontoquerythelossesofallpossibleactions.Forthisproblem,Cesa-Bianchietal.[9]provedanupperboundonthenormalizedregretoforderO(Kp ln(4N/δ)/(m))whichholdswithprobabilityatleast1−δ.Ourmodelofthelabel-eﬃcientbanditproblemforshortestpathsismotivatedbyanapplicationtoaparticularpacketswitchednetworkmodel.Thismodel,calledthecognitivepacketnetwork,wasintroducedbyGelenbeetal.[11,12].Inthesenetworksaparticulartypeofpackets,calledsmartpackets,areusedtoexplorethenetwork(e.g.,thedelayofthechosenpath).Thesepacketsdonotcarryanyusefuldata;theyaremerelyusedforexploringthenetwork.Theothertypeofpacketsarethedatapackets,whichdonotcollectanyinformationabouttheirpaths.Thetaskofthedecisionmakeristosendpacketsfromthesourcetothedestinationoverrouteswithminimumaveragetransmissiondelay(orpacketloss).Inthisscenario,smartpacketsareusedtoquerythedelay(orloss)ofthechosenpath.However,asthesepacketsdonottransportinformation,thereisatradeoﬀbetweenthenumberofqueriesandtheusageofthenetwork.Ifdatapacketsareontheaverageαtimeslargerthansmartpackets(notethattypicallyα1)andistheproportionoftimeinstanceswhensmartpacketsareusedtoexplorethenetwork,then/(+α(1−))istheproportionofthebandwidthsacriﬁcedforwellinformedroutingdecisions.Westudyacombinedalgorithmwhich,ateachtimeslott,queriesthelossofthechosenpathwithprobability(asinthesolutionofthelabeleﬃcientproblemproposedin[9]),and,similarlytothemulti-armedbanditcase,computesbiasedestimatesg0i,tofthetruegainsgi,t.Justasintheprevioussection,itisassumedthateachpathofthegraphisofthesamelengthK.Thealgorithmdiﬀersfromourbanditalgorithmoftheprevioussectiononlyinstep(c),whichismodiﬁedinthespiritof[9].Themodiﬁedstepisgivenbelow:MODIFIEDSTEPFORTHELABELEFFICIENTBANDITALGORITHMFORSHORTESTPATHS(c’)DrawaBernoullirandomvariableStwithP(St=1)=,andcomputetheestimatedgainsg0e,t=(ge,t+β
qe,tStife∈Itβ
qe,tStife/∈It.
Theperformanceofthealgorithmisanalyzedinthenexttheorem,whichcanbeviewedasacombinationofTheorem1intheprecedingsectionandTheorem2of[9].13
Theorem3Foranyδ∈(0,1),∈(0,1],parametersη=q
lnN
4nK2|C|,γ=2ηK|C|
≤1/2,andβ=q
K
n|E|ln2|E|
δ≤1,andforalln≥1
maxK2ln2(2|E|/δ)
|E|lnN,|E|ln(2|E|/δ)
K,4|C|lnNtheperformanceofthealgorithmdeﬁnedabovecanbebounded,withprobabilityatleast1−δ,as1
n bLn−mini∈PnXt=1`i,t!≤r
K
n 4p
K|C|lnN+5r
|E|ln2|E|
δ+r
8Kln2
δ!+4K
3nln2N
δ≤27K
2s
|E|ln2N
δ
n.Ifischosenas(m−p
2mln(1/δ))/nthen,withprobabilityatleast1−δ,thetotalnumberofqueriesisboundedbym(see[8,Lemma6.1])andtheperformanceboundaboveisoftheorderofKp
|E|ln(N/δ)/m.SimilarlytoTheorem1,weneedalemmawhichrevealstheconnectionbetweenthetrueandtheestimatedcumulativelosses.However,hereweneedamorecarefulanalysisbecausethe“shiftingterm”β
qe,tSt,isarandomvariable.Lemma2Forany0<δ<1,0<≤1,foranyn≥1
maxK2ln2(2|E|/δ)
|E|lnN,Kln(2|E|/δ)
|E|,parameters2ηK|C|
≤γ,η=s
lnN
4nK2|C|andβ=s
K
n|E|ln2|E|
δ≤1,ande∈E,wehavePGe,n>G0
e,n+4
βln2|E|
δ≤δ
2|E|.Proof.Fixe∈E.Using(1)withu=4
βln2|E|
δandc=β
4,itsuﬃcestoproveforallnthatEhec(Ge,n−G0
e,n)i≤1.SimilarlytoLemma1weintroduceZt=ec(Ge,t−G0
e,t)andweshowthatZ1,...,Znisasuper-martingale,thatisEt[Zt]≤Zt−1fort≥2whereEtdenotesE[·|(I1,S1),...,(It−1,St−1)].Takingconditionalexpectations,weobtainEt[Zt]=Zt−1Et"ec„ge,t−1{e∈It}Stge,t+Stβ
qe,t«#14
≤Zt−1Et1+cge,t−1{e∈It}Stge,t+Stβ
qe,t+c2ge,t−1{e∈It}Stge,t+Stβ
qe,t2#.(10)SinceEtge,t−1{e∈It}Stge,t+Stβ
qe,t=−β
qe,tandEt"ge,t−1{e∈It}Stge,t qe,t2#≤Et"1{e∈It}Stge,t qe,t2#≤1
qe,twegetfrom(10)thatEt[Zt]≤Zt−1Et"1−cβ
qe,t+c2
qe,t+c2 21{e∈It}Stge,tβ
q2e,t2−2ge,tStβ
qe,t+Stβ2
q2e,t2!#≤Zt−11+c qe,t−β+c
+cβ2
+β
qe,t.(11)Sincec=β/4wehave−β+c
+cβ2
+β
qe,t=−3β
4+β2
42
+β
qe,t=−3β
4+β2
2+β3
4qe,t≤−β
4+β3
4qe,t≤−β
4+β3|C|
4γ(12)≤0,(13)where(12)followsfromqe,t≥γ
|C|and(13)holdsbyβ2|C|
γ≤β2
2ηK≤1,andthelastinequalityisensuredbyn≥K2ln2(2|E|/δ)
|E|lnN,theassumptionofthelemma.Combining(11)and(13)wegetthatEt[Zt]≤Zt−1.Takingexpectationsonbothsidesoftheinequality,wegetE[Zt]≤E[Zt−1]andsinceE[Z1]≤1,weobtainE[Zn]≤1asdesired.2ProofofTheorem3.TheproofofthetheoremisageneralizationofthatofTheorem1,andfollowsthesamelineswithsomeextratechnicalitiestohandletheeﬀectsofthemodiﬁedstep(c’).15
Therefore,inthefollowingweemphasizeonlythediﬀerences.Firstnotethat(5)and(7)alsoholdinthiscase.Boundingthesumsin(7),oneobtainsX
i∈Ppi,tg0i,t=St
(gIt,t+|E|β)andX
i∈Ppi,tg02
i,t≤1
K(1+β)Xe∈Eg0e,t.Pluggingtheseboundsinto(7)andsummingfort=1,...,n,weobtainln
Wn
W0≤η
1−γnXt=1St
(gIt,t+|E|β)+η2K(1+β)
(1−γ)|C|maxi∈PG0
i,n.Combiningtheupperboundwiththelowerbound(5),weobtainnXt=1St
(gIt,t+|E|β)≥1−γ−ηK(1+β)|C|
maxi∈PG0
i,n−lnN
η.(14)Torelatetheleft-handsideoftheaboveinequalitytotherealgainPn t=1gIt,t,noticethatXt=St
(gIt,t+|E|β)−(gIt,t+|E|β)isamartingalediﬀerencesequencewithrespectto(I1,S1),(I2,S2),....Nowforallt=1,...,n,wehavetheboundEX2t|(I1,S1),...,(It−1,St−1)≤ESt
2(gIt,t+|E|β)2

(I1,S1),...,(It−1,St−1)≤(K+|E|β)2
≤4K2
def=σ2,(15)where(15)holdsbyn≥|E|ln(2|E|/δ)
K(toensureβ|E|≤K).WeknowthatXt∈−2K,1
−12Kforallt.NowapplyBernstein’sinequalityformartingalediﬀerences(seeLemma9intheAppendix)toobtainP"nXt=1Xt>u#≤δ
2,(16)whereu=s
2n4K2
ln2
δ+4K
3ln2
δ.16
From(16)wegetP"nXt=1St
(gIt,t+|E|β)≥bGn+βn|E|+u#≤δ
2.(17)NowLemma2,theunionbound,and(17)combinedwith(14)yield,withprobabilityatleast1−δ,bGn≥1−γ−ηK(1+β)|C|
maxi∈PGi,n−4K
βln2|E|
δ−lnN
η−βn|E|−usincethecoeﬃcientofG0
i,nisgreaterthanzerobytheassumptionsofthetheorem.SincebGn=Kn−bLnandGi,n=Kn−Li,n,wehavebLn≤1−γ−K(1+β)η|C|
mini∈PLi,n+Knγ+K(1+β)η|C|
+1−γ−K(1+β)η|C|
4K
βln2|E|
δ+βn|E|+lnN
η+u≤mini∈PLi,n+Knγ+K(1+β)η|C|
+5βn|E|+lnN
η+u,whereweusedthefactthatK
βln2|E|
δ=βn|E|.Substitutingthevalueofβ,ηandγ,wehavebLn−mini∈PLi,n≤Kn2Kη|C|
+Kn2Kη|C|
+lnN
η+5βn|E|+u≤4Kr n|C|lnN
+5r n|E|Kln(2|E|/δ)
+u≤r nK
4p
K|C|lnN+5p
|E|ln(2|E|/δ)+p
8Kln(2/δ)+4K
3ln(2/δ)asdesired.26Abanditalgorithmfortrackingtheshortestpath
Ourgoalinthissectionistoextendthebanditalgorithmsothatitisabletocompetewithtime-varyingpathsundersmallcomputationalcomplexity.Thisisavariantoftheproblemknownastrackingthebestexpert;see,forexample,HerbsterandWarmuth[19],Vovk[30],AuerandWarmuth[2],BousquetandWarmuth[6],HerbsterandWarmuth[20].Todescribethelossthedecisionmakeriscomparedto,considerthefollowing“m-partition”predictionscheme:thesequenceofpathsispartitionedintom+1contiguoussegments,andoneachsegmenttheschemeassignsexactlyoneoftheNpaths.Formally,anm-partitionPart(n,m,t,i
)ofthenpathsisgivenbyanm-tuplet=(t1,...,tm)suchthatt0=1<t1<···<tm<n+1=tm+1,17
andan(m+1)-vectori
=(i0,...,im)whereij∈P.Ateachtimeinstantt,tj≤t<tj+1,pathijisusedtopredictthebestpath.ThecumulativelossofapartitionPart(n,m,t,i
)isL(Part(n,m,t,i
))=mXj=0tj+1−1Xt=tj`ij,t=mXj=0tj+1−1Xt=tjXe∈ij`e,t.Thegoalofthedecisionmakeristoperformaswellasthebesttime-varyingpath(partition),thatis,tokeepthenormalizedregret1
nbLn−mint,i
L(Part(n,m,t,i
))assmallaspossible(withhighprobability)forallpossibleoutcomesequences.Inthe“classical”trackingproblemthereisarelativelysmallnumberof“base”expertsandthegoalofthedecisionmakeristopredictaswellasthebest“compound”expert(i.e.,time-varyingexpert).Howeverinourcase,baseexpertscorrespondtoallpathsofthegraphbetweensourceanddestinationwhosenumberistypicallyexponentiallylargeinthenumberofedges,andthereforewecannotdirectlyapplythecomputationallyeﬃcientmethodsfortrackingthebestexpert.Gy¨orgy,Linder,andLugosi[15]developeﬃcientalgorithmsfortrackingthebestexpertforcertainlargeandstructuredclassesofbaseexperts,includingtheshortestpathproblem.Thepurposeofthefollowingalgorithmistoextendthemethodsof[15]tothebanditsettingwhentheforecasteronlyobservesthelossesoftheedgesonthechosenpath.18
ABANDITALGORITHMFORTRACKINGSHORTESTPATHSParameters:realnumbersβ>0,0<η,γ<1,0≤α≤1.Initialization:Setwe,0=1foreache∈E, wi,0=1foreachi∈P,and
W0=N.Foreachroundt=1,2,...(a)ChooseapathItaccordingtothedistributionptdeﬁnedbypi,t=

(1−γ)
wi,t−1
Wt−1+γ
|C|ifi∈C;(1−γ)
wi,t−1
Wt−1ifi6∈C.(b)Computetheprobabilityofchoosingeachedgeeasqe,t=Xi:e∈ipi,t=(1−γ)Pi:e∈i wi,t−1
Wt−1+γ|{i∈C:e∈i}|
|C|.(c)Calculatetheestimatedgainsg0e,t=(ge,t+β
qe,tife∈It;β
qe,totherwise.(d)Computetheupdatedweights vi,t=
wi,t−1eηg0i,t wi,t=(1−α)
vi,t+α
N
Wtwhereg0i,t=Pe∈ig0e,tand
Wtisthesumofthetotalweightsofthepaths,thatis,
Wt=Xi∈P
vi,t=Xi∈P
wi,t.
Thefollowingperformanceboundsshowsthatthenormalizedregretwithrespecttothebesttime-varyingpathwhichisallowedtoswitchpathsmtimesisroughlyoftheorderofKp
(m/n)|C|lnN.Theorem4Foranyδ∈(0,1)andparameters0≤γ<1/2,α,β∈[0,1],andη>0satisfying2ηK|C|≤γ,theperformanceofthealgorithmdeﬁnedabovecanbebounded,withprobabilityatleast1−δ,asbLn−mint,i
L(Part(n,m,t,i
))≤Kn(γ+η(1+β)K|C|)+K(m+1)
βln|E|(m+1)
δ+βn|E|+1
ηlnNm+1
αm(1−α)n−m−1.19
Inparticular,choosingβ=s
K(m+1)
n|E|ln|E|(m+1)
δ,γ=2ηK|C|,α=m n−1,andη=s
(m+1)lnN+mlne(n−1)
m
4nK2|C|wehave,foralln≥maxnK(m+1)
|E|ln|E|(m+1)
δ,4|C|Do,bLn−mint,i
L(Part(n,m,t,i
))≤2√
nK p
4K|C|D+r
|E|(m+1)ln|E|(m+1)
δ!,whereD=(m+1)lnN+m1+lnn−1
m.TheproofofthetheoremisacombinationofthatofourTheorem1andTheorem1of[15].Wewillneedthefollowingthreelemmas.
Lemma3Forany1≤t≤t0≤nandanyi∈P, vi,t0
wi,t−1≥eηG0
i,[t,t0](1−α)t0−twhereG0
i,[t,t0]=Pt0τ=tg0i,τ.Proof.TheproofisastraightforwardmodiﬁcationoftheoneinHerbsterandWarmuth[19].Fromthedeﬁnitionsofvi,tandwi,t(seestep(d)ofthealgorithm)itisclearthatforanyτ≥1, wi,τ=(1−α)
vi,τ+α
N
Wτ≥(1−α)eηg0i,τ
wi,τ−1.Applyingthisequationiterativelyforτ=t,t+1,...,t0−1,andthedeﬁnitionof vi,t(step(d))forτ=t0,weobtain vi,t0=
wi,t0−1eηg0i,t0≥eηg0i,t0t0−1Yτ=t(1−α)eηg0i,τ
wi,t−1=eηG0
i,[t,t0](1−α)t0−t wi,t−1whichimpliesthestatementofthelemma.2Lemma4Foranyt≥1andi,j∈P,wehave wi,t vj,t≥α
N20
Proof.Bythedeﬁnitionof wi,twehave wi,t=(1−α)
vi,t+α
N
Wt≥α
N
Wt≥α
N
vj,t.Thiscompletestheproofofthelemma.2ThenextlemmaisasimplecorollaryofLemma1.Lemma5Foranyδ∈(0,1),0≤β≤1,t≥1ande∈EwehavePGe,t>G0
e,t+1
βln|E|(m+1)
δ≤δ
|E|(m+1).ProofofTheorem4.ThetheoremisprovedthesamewayasTheorem1until(8),thatis,ln
Wn
W0≤η
1−γbGn+n|E|β+η2K(1+β)
1−γ|C|maxi∈PG0
i,n.(18)LetPart(n,m,t,i
)beanarbitrarypartition.Thenthelowerboundisobtainedasln
Wn
W0=lnXj∈P
wj,n
N=lnXj∈P
vj,n
N≥ln vim,n
N(19)(recallthatimdenotesthepathusedinthelastsegmentofthepartition).Nowvim,ncanberewrittenintheformofthefollowingtelescopingproduct vim,n=
wi0,t0−1
vi0,t1−1
wi0,t0−1mYj=1
wij,tj−1
vij−1,tj−1
vij,tj+1−1
wij,tj−1.Therefore,applyingLemmas3and4,wehave vim,n≥
wi0,t0−1α
NmmYj=0(1−α)tj+1−1−tjeηG0
ij,[tj,tj+1−1]=α
NmeηG0(Part(n,m,t,i
))(1−α)n−m−1.Combiningthelowerboundwiththeupperbound(18),wehavelnαm(1−α)n−m−1
Nm+1+maxt,i
ηG0(Part(n,m,t,i
))≤η
1−γbGn+n|E|β+η2K(1+β)
1−γ|C|maxi∈PG0
i,n,whereweusedthefactthatPart(n,m,t,i
)isanarbitrarypartition.Afterrearrangingandusingmaxi∈PG0
i,n≤maxt,i
G0(Part(n,m,t,i
))wegetbGn≥(1−γ−ηK(1+β)|C|)maxt,i
G0(Part(n,m,t,i
))−n|E|β−1−γ
ηlnNm+1
αm(1−α)n−m−1.21
Nowsince1−γ−ηK(1+β)|C|≥0,bytheassumptionsofthetheoremandfromLemma5withanapplicationoftheunionboundweobtainthat,withprobabilityatleast1−δ,bGn≥(1−γ−ηK(1+β)|C|)maxt,i
G(Part(n,m,t,i
))−K(m+1)
βln|E|(m+1)
δ−n|E|β−1−γ
ηlnNm+1
αm(1−α)n−m−1.SincebGn=Kn−bLnandG(Part(n,m,t,i
))=Kn−L(Part(n,m,t,i
)),wehavebLn≤(1−γ−ηK(1+β)|C|)mint,iL(Part(n,m,t,i
))+Kn(γ+η(1+β)K|C|)+(1−γ−η(1+β)K|C|)K(m+1)
βln|E|(m+1)
δ+n|E|β+1
ηlnNm+1
αm(1−α)n−m−1.Thisimpliesthat,withprobabilityatleast1−δ,bLn−mint,i
L(Part(n,m,t,i
))≤Kn(γ+η(1+β)K|C|)+K(m+1)
βln|E|(m+1)
δ+n|E|β+1
ηlnNm+1
αm(1−α)n−m−1.(20)Toprovethesecondstatement,letH(p)=−plnp−(1−p)ln(1−p)andD(pkq)=plnp q+(1−p)ln1−p
1−q.Optimizingthevalueofαinthelasttermof(20)gives1
ηlnNm+1
αm(1−α)n−m−1=1
η(m+1)ln(N)+mln1
α+(n−m−1)ln1
1−α=1
η (m+1)ln(N)+(n−1)(Db(α∗kα)+Hb(α∗))whereα∗=m n−1.Forα=α∗weobtain1
ηlnNm+1
αm(1−α)n−m−1=1
η((m+1)ln(N)+(n−1)(Hb(α∗)))=1
η((m+1)ln(N)+mln((n−1)/m)+(n−m−1)ln(1+m/(n−m−1)))≤1
η((m+1)ln(N)+mln((n−1)/m)+m)22
=1
η(m+1)ln(N)+mlne(n−1)
mdef=1
ηDwheretheinequalityfollowssinceln(1+x)≤xforx>0.ThereforebLn−mint,i
L(Part(n,m,t,i
))≤Kn(γ+η(1+β)K|C|)+K(m+1)
βln|E|(m+1)
δ+n|E|β+1
ηD.whichistheﬁrststatementofthetheorem.Settingβ=s
K(m+1)
n|E|ln|E|(m+1)
δ,γ=2ηK|C|,andη=s
D
4nK2|C|resultsinthesecondstatementofthetheorem,thatis,bLn−mint,i
L(Part(n,m,t,i
))≤2√
nK p
4K|C|D+r
|E|(m+1)ln|E|(m+1)
δ!.2Similarlyto[15],theproposedalgorithmhasanalternativeversion,whichiseﬃcientlycom-putable:ANALTERNATIVEBANDITALGORITHMFORTRACKINGSHORTESTPATHSFort=1,chooseI1uniformlyfromthesetP.Fort≥2,(a)DrawaBernoullirandomvariableΓtwithP(Γt=1)=γ.(b)IfΓt=1,thenchooseItuniformlyfromC.(c)IfΓt=0,(c1)chooseτtrandomlyaccordingtothedistributionP{τt=t0}=

(1−α)t−1Z1,t−1
Wt−1fort0=1α(1−α)t−t0Wt0Zt0,t−1
NWtfort0=2,...,twhereZt0,t−1=Pi∈PeηG0
i,[t0,t−1]fort0=1,...,t−1,andZt,t−1=N;(c2)givenτt=t0,chooseItrandomlyaccordingtotheprobabilitiesP{It=i|τt=t0}=

eηG0
i,[t0,t−1]
Zt0,t−1fort0=1,...,t−11
Nfort0=t.
23
Inawaycompletelyanalogousto[15],inthisalternativeformulationofthealgorithmonecancomputetheprobabilitiesP{It=i|τt=t0}andthenormalizationfactorsZt0,t−1eﬃciently.UsingthefactthatthebaselinebanditalgorithmforshortestpathshasanO(n|E|)timecomplexitybyTheorem2,itfollowsfromTheorem3of[15]thatthetimecomplexityofthealternativebanditalgorithmfortrackingtheshortestpathisO(n2|E|).7Analgorithmfortherestrictedmulti-armedbanditproblemInthissectionweconsiderthesituationwherethedecisionmakerreceivesinformationonlyabouttheperformanceofthewholechosenpath,buttheindividualedgelossesarenotavailable.Thatis,theforecasterhasaccesstothesum`It,toflossesoverthechosenpathItbutnottothelosses{`e,t}e∈ItoftheedgesbelongingtoIt.ThisistheproblemformulationconsideredbyMcMahanandBlum[25]andAwerbuchandKleinberg[4].McMahanandBlumprovidedarelativelysimplealgorithmwhoseregretisatmostoftheorderofn−1/4,whileAwerbuchandKleinberggaveamorecomplexalgorithmtoachieveO(n−1/3)regret.Inthissectionwecombinethestrengthsofthesepapers,andproposeasimplealgorithmwithregretatmostoftheorderofn−1/3.Moreover,ourboundholdswithhighprobability,whiletheabove-mentionedpapersproveboundsfortheexpectedregretonly.TheproposedalgorithmusesideasverysimilartothoseofMcMahanandBlum[25].Thealgorithmalternatesbetweenchoosingapathfroma“basis”Btoobtainunbiasedestimatesoftheloss(explorationstep),andchoosingapathaccordingtoexponentialweightingbasedontheseestimates.Asimplewaytodescribeapathi∈Pisabinaryrowvectorwith|E|componentswhichareindexedbytheedgesofthegraphsuchthat,foreache∈E,theethentryofthevectoris1ife∈iand0otherwise.Withaslightabuseofnotationwewillalsodenotebyithebinaryrowvectorrepresentingpathi.Intheprevioussections,wherethelossofeachedgealongthechosenpathisavailabletothedecisionmaker,thecomplexitystemmingfromthelargenumberofpathswasreducedbyrepresentingallinformationintermsoftheedges,asthesetofedgesspansthesetofpaths.Thatis,thevectorcorrespondingtoagivenpathcanbeexpressedasthelinearcombinationoftheunitvectorsassociatedwiththeedges(theethcomponentoftheunitvectorrepresentingedgeeis1,whiletheothercomponentsare0).Whilethelossescorrespondingtosuchaspanningsetarenotobservableintherestrictedsettingofthissection,onecanchooseasubsetofPthatformsabasis,thatis,acollectionofbpathswhicharelinearlyindependentandeachpathinPcanbeexpressedasalinearcombinationofthepathsinthebasis.WedenotebyBtheb×|E|matrixwhoserowsb1,...,bbrepresentthepathsinthebasis.Notethatbisequaltothemaximumnumberoflinearlyindependentvectorsin{i:i∈P},sob≤|E|.Let`(E)tdenotethe(column)vectoroftheedgelosses{`e,t}e∈Eattimet,andlet`(B)t=(`b1,t,...,`bb,t)Tbeab-dimensionalcolumnvectorwhosecomponentsarethelossesofthepathsinthebasisBattimet.Ifα(i,B)b1,...,α(i,B)bbarethecoeﬃcientsinthelinearcombinationofthebasispathsexpressingpathi∈P,thatis,i=Pb j=1α(i,B)bjbj,thenthelossofpathi∈Pattimetisgivenby`i,t=hi,`(E)ti=bXj=1α(i,B)bjhbj,`(E)ti=bXj=1α(i,B)bj`bj,t(21)whereh·,·idenotesthestandardinnerproductinR|E|.Inthealgorithmweobtainestimates˜`bj,t24
ofthelossesofthebasispathsanduse(21)toestimatethelossofanyi∈Pas˜`i,t=bXj=1α(i,B)bj˜`bj,t.(22)Itisalgorithmicallyadvantageoustocalculatetheestimatedpathlosses˜`i,tfromanintermediateestimateoftheindividualedgelosses.LetB+denotethetheMoore-PenroseinverseofBdeﬁnedbyB+=BT(BBT)−1,whereBTdenotesthetransposeofBandBBTisinvertiblesincetherowsofBarelinearlyindependent.(NotethatB+=B−1ifb=|E|).Thenletting˜`(B)t=(˜`b1,t,...,˜`bb,t)Tand˜`(E)t=B+˜`(B)titiseasytoseethat˜`i,tin(22)canbeobtainedas˜`i,t=hi,˜`(E)ti,orequivalently˜`i,t=Xe∈i˜`e,t.Thisformofthepathlossesallowsforaneﬃcientimplementationofexponentialweightingviadynamicprogramming[28].Toanalyzethealgorithmweneedanupperboundonthemagnitudeofthecoeﬃcientsα(i,B)bj.Forthis,weinvokethedeﬁnitionofabarycentricspannerfrom[4]:thebasisBiscalledaC-barycentricspannerif|α(i,B)bj|≤Cforalli∈Pandj=1,...,b.AwerbuchandKleinberg[4]showthata1-barycentricspannerexistsifBisasquarematrix(i.e.,b=|E|)andgivealow-complexityalgorithmwhichﬁndsaC-barycentricspannerforC>1.Weusetheirtechniquetoshowthata1-barycentricspanneralsoexistsincaseofanon-squareB,whenthebasisischosentomaximizetheabsolutevalueofthedeterminantofBBT.Asbefore,bdenotesthemaximumnumberoflinearlyindependentvectors(paths)inP.Lemma6Foradirectedacyclicgraph,thesetofpathsPbetweentwodedicatednodeshasa1-barycentricspanner.Moreover,letBbeab×|E|matrixwithrowsfromPsuchthatdet[BBT]6=0.IfB−j,iisthematrixobtainedfromBbyreplacingitsjthrowbyi∈Pand
detB−j,iBT
−j,i
≤C2
detBBT
(23)forallj=1,...,bandi∈P,thenBisaC-barycentricspanner.Proof.LetBbeabasisofPwithrowsb1,...,bb∈Pthatmaximizes|det[BBT]|.Then,foranypathi∈P,wehavei=Pb j=1α(i,B)bjbjforsomecoeﬃcients{α(i,B)bj}.NowforthematrixB−1,i=[iT,(b2)T,...,(bb)T]Twehave
detB−1,iBT
−1,i
=

dethB−1,iiT,B−1,i(b2)T,B−1,i(b3)T,...,B−1,i(bb)Ti

=

det


bXj=1α(i,B)bjB−1,ibj
T,B−1,i(b2)T,B−1,i(b3)T,...,B−1,i(bb)T



25
=

bXj=1α(i,B)bjdethB−1,i(bj)T,B−1,i(b2)T,B−1,i(b3)T,...,B−1,i(bb)Ti

=|α(i,B)b1|
detB−1,iBT
=α(i,B)b12
detBBT
wherelastequalityfollowsbythesameargumentthepenultimateequalitywasobtained.RepeatingthesameargumentforB−j,i,j=2,...,bweobtain
detB−j,iBT
−j,i
=α(i,B)bj2
detBBT
.(24)Thusthemaximalpropertyof|det[BBT]|implies|α(i,B)bj|≤1forallj=1,...,b.Thesecondstatementfollowstriviallyfrom(23)and(24).2AwerbuchandKleinberg[4]alsopresentaniterativealgorithmtoﬁndaC-barycentricspannerifBisasquarematrix.Startingfromtheidentitymatrix,theiralgorithmreplacesarowofthematrixineachstepbymaximizingthedeterminantwithrespecttothegivenrow.Thisisdonebycallinganoraclefunction,anditisshownthattheoracleiscalledO(blogCb)times.IncaseBisnotasquarematrix,thealgorithmcarriesoverifwehaveaccesstoanalternativeoraclethatcanmaximize|det[BBT]|:StartingfromanarbitrarybasisBwecaniterativelyreplaceonerowineachstep,usingtheoracle,tomaximizethedeterminant|det[BBT]|until(23)issatisﬁedforalljandi.ByLemma6,thisresultsinaC-barycentricspanner.Similarlyto[4],itcanbeshownthattheoracleiscalledO(blogCb)timesforC>1.Forsimplicity(toavoidcarryingtheconstantC),assumethatwehavea2-barycentricspannerB.Basedontheideasoflabeleﬃcientprediction,thenextalgorithmgivesasimplesolutiontotherestrictedshortestpathproblem.Thealgorithmisverysimilartothatofthealgorithminthelabeleﬃcientcase,butherewecannotestimatetheedgelossesdirectly.Therefore,wequerythelossofa(random)basisvectorfromtimetotime,andcreateunbiasedestimates˜`bj,tofthelossesofbasispaths`bj,t,whicharethentransformedintoedge-lossestimates.26
ABANDITALGORITHMFORTHERESTRICTEDSHORTESTPATHPROBLEMParameters:0<,η≤1.Initialization:Setwe,0=1foreache∈E, wi,0=1foreachi∈P,
W0=N.FixabasisB,whichisa2-barycentricspanner.Foreachroundt=1,2,...(a)DrawaBernoullirandomvariableStsuchthatP(St=1)=;(b)IfSt=1,thenchoosethepathItuniformlyfromthebasisB.IfSt=0,thenchooseItaccordingtothedistribution{pi,t},deﬁnedbypi,t=
wi,t−1
Wt−1.(c)Calculatetheestimatedlossofalledgesaccordingto˜`(E)t=B+˜`(B)t,where˜`(E)t={˜`(E)e,t}e∈E,and˜`(B)t=(˜`(B)b1,t,...,˜`(B)bb,t)isthevectoroftheestimatedlosses˜`bj,t=St
`bj,t1{It=bj}bforj=1,...,b.(d)Computetheupdatedweightswe,t=we,t−1e−η˜`e,t, wi,t=Ye∈iwe,t=
wi,t−1e−ηPe∈i˜`e,t,andthesumofthetotalweightsofthepaths
Wt=X
i∈P
wi,t.
Theperformanceofthealgorithmisanalyzedinthenexttheorem.Theprooffollowstheargu-mentofCesa-Bianchietal.[9],butwealsohavetodealwithsomeadditionaltechnicaldiﬃculties.Notethatinthetheoremwedonotassumethatallpathsbetweenuandvhaveequallength.Theorem5LetKdenotethelengthofthelongestpathinthegraph.Foranyδ∈(0,1),parameters0<≤1
Kandη>0satisfyingη≤2,andn≥8b
2ln4bN
δ,theperformanceofthealgorithmdeﬁnedabovecanbebounded,withprobabilityatleast1−δ,asbLn−mini∈PLi,n≤K
ηb
Kn+r n
2ln4
δ+n+q
2nln4
δ
K+16
3br
2nb
ln4bN
δ
+lnN
η27
Inparticular,choosing=Kb nln4bN
δ1/3andη=2weobtainbLn−mini∈PLi,n≤9.1K2b(Kbln(4bN/δ))1/3n2/3.Thetheoremisprovedusingthefollowingtwolemmas.TheﬁrstoneisaneasyconsequenceofBernstein’sinequality:
Lemma7UndertheassumptionsofTheorem5,theprobabilitythatthealgorithmqueriesthebasismorethann+q
2nln4
δtimesisatmostδ/4.Usingtheestimatedlossofapathi∈Pgivenin(22),wecanestimatethecumulativelossofiuptotimenas˜Li,n=nXt=1˜`i,t.Thenextlemmademonstratesthequalityoftheseestimates.
Lemma8Let0<δ<1andassumen≥8b
ln4bN
δ.Foranyi∈P,withprobabilityatleast1−δ/4,nXt=1Xi∈Ppi,t`i,t−nXt=1X
i∈Ppi,t˜`i,t≤8
3br
2nbK2
ln4b
δ.Furthermore,withprobabilityatleast1−δ/(4N),˜Li,n−Li,n≤8
3br
2nbK2
ln4bN
δ.Proof.WemaywritenXt=1X
i∈Ppi,t`i,t−nXt=1Xi∈Ppi,t˜`i,t=nXt=1X
i∈Ppi,tbXj=1α(i,B)bj`bj,t−˜`bj,t=bXj=1nXt=1"Xi∈Ppi,tα(i,B)bj`bj,t−˜`bj,t#def=bXj=1nXt=1Xbj,t.(25)Notethatforanybj,Xbj,t,t=1,2,...isamartingalediﬀerencesequencewithrespectto(It,St),t=1,2,...asEt˜`b,t=`b,t.Also,Et[X2bj,t]≤ X
i∈Ppi,tα(i,B)bj!2Et˜`bj,t2≤X
i∈Ppi,tα(i,B)bj2K2b
≤4K2b
(26)28
and|Xbj,t|≤

Xi∈Ppi,tα(i,B)bj

`bj,t−˜`bj,t

≤X
i∈Ppi,t

α(i,B)bj

Kb
≤2Kb
(27)wherethelastinequalitiesinbothcasesfollowfromthefactthatBisa2-barycentricspanner.Then,usingBernstein’sinequalityformartingalediﬀerences(Lemma9),wehave,foranyﬁxedbj,P"nXt=1Xbj,t≥8
3r
2nbK2
ln4b
δ#≤δ
4b(28)whereweused(26),(27)andtheassumptionofthelemmaonn.Theproofoftheﬁrststatementisﬁnishedwithanapplicationoftheunionboundanditscombinationwith(25).Forthesecondstatementweuseasimilarargument,thatis,nXt=1(˜`i,t−`i,t)=bXj=1α(i,B)bjnXt=1(˜`bj,t−`bj,t)≤Xj=1

α(i,B)bj

nXt=1(˜`bj,t−`bj,t)

≤2bXj=1

nXt=1(˜`bj,t−`bj,t)

.(29)NowapplyingLemma9foraﬁxedbjwegetP"nXt=1(˜`bj,t−`bj,t)≥4
3r
2nK2b
ln4bN
δ#≤δ
4bN(30)becauseofEt[(˜`bj,t−`bj,t)2]≤K2b
and−K≤˜`bj,t−`bj,t≤K b
−1.Theproofiscompletedbyapplyingtheunionboundto(30)andcombiningtheresultwith(29).2ProofofTheorem5.Similarlytoearlierproofs,wefollowtheevolutionofthetermln
Wn
W0.Inthesamewayasweobtained(5)and(7),wehaveln
Wn
W0≥−ηmini∈PeLi,n−lnNandln
Wn
W0≤nXt=1 −ηX
i∈Ppi,t˜`i,t+η2
2X
i∈Ppi,t˜`2
i,t!.Combiningthesebounds,weobtain−mini∈PeLi,n−lnN
η≤nXt=1 −Xi∈Ppi,t˜`i,t+η
2Xi∈Ppi,t˜`2
i,t!≤−1+ηKb
nXt=1X
i∈Ppi,t˜`i,t,29
because0≤˜`i,t≤2Kb
.ApplyingtheresultsofLemma8andtheunionbound,wehave,withprobability1−δ/2,−mini∈PLi,n−8
3br
2nK2b
ln4bN
δ≤−1+ηKb
 nXt=1Xi∈Ppi,t`i,t−8
3br
2nK2b
ln4b
δ!+lnN
η≤−1+ηKb
nXt=1X
i∈Ppi,t`i,t+8
3br
2nK2b
ln4b
δ+lnN
η.(31)IntroducethesetsTndef={t:1≤t≤nandSt=0}and
Tndef={t:1≤t≤nandSt=1}of“exploitation”and“exploration”steps,respectively.Then,bytheHoeﬀding-Azumainequality[21]weobtainthat,withprobabilityatleast1−δ/4,Xt∈TnXi∈Ppi,t`i,t≥Xt∈Tn`It,t−r
|Tn|K2
2ln4
δ.Notethatfortheexplorationstepst∈
Tn,asthealgorithmplaysaccordingtoauniformdistribu-tioninsteadofpi,t,wecanonlyusethetriviallowerboundzeroonthelosses,thatis,Xt∈
TnX
i∈Ppi,t`i,t≥Xt∈
Tn`It,t−K|
Tn|.ThelasttwoinequalitiesimplynXt=1X
i∈Ppi,t`i,t≥bLn−r
|Tn|K2
2ln4
δ−K|
Tn|.(32)Then,by(31),(32)andLemma7weobtain,withprobabilityatleast1−δ,bLn−mini∈PLi,n≤K
ηb
Kn+r n
2ln4
δ+n+q
2nln4
δ
K+16
3br
2nb
ln4bN
δ
+lnN
ηwhereweusedbLn≤Knand|Tn|≤n.SubstitutingthevaluesofandηgivesbLn−mini∈PLi,n≤K2bn+1
4Kn+Kn+1
2n+16
3b√
Kn+n≤9.1K2bnwhereweusedq n
2ln4
δ≤1
4n,q
2nln4
δ≤1
2n,q nbK
ln4N
δ=n,andlnN
η≤n(fromtheassumptionsofthetheorem).230
8Simulationresults
Tofurtherinvestigateournewalgorithms,wehaveconductedsomesimplesimulations.Asthemainmotivationofthisworkistoimproveearlieralgorithmsincasethenumberofpathsisexponentiallylargeinthenumberofedges,wetestedthealgorithmsonthesmallgraphshowninFigure1(b),whichhasoneofthesimpleststructureswithexponentiallymany(namely2|E|/2)paths.Thelossesontheedgesweregeneratedbyasequenceofindependentanduniformrandomvariables,withvaluesfrom[0,1]ontheupperedges,andfrom[0.32,1]ontheloweredges,resultingina(long-term)optimalpathconsistingoftheupperedges.Weranthetestsforn=10000steps,withconﬁdencevalueδ=0.001.Toestablishbaselineperformance,wealsotestedtheEXP3algorithmofAueretal.[1](notethatthisalgorithmdoesnotneededgelosses,onlythelossofthechosenpath).Fortheversionofourbanditalgorithmthatisinformedoftheindividualedgelosses(edge-bandit),weusedthesimple2-elementcoversetofthepathsconsistingoftheupperandloweredges,respectively(other2-elementcoversetsgivesimilarperformance).Forourrestrictedshortestpathalgorithm(path-bandit)thebasis{uuuuu,uuuul,uuull,uulll,ullll}wasused,whereu(resp.l)inthekthpositiondenotestheupper(resp.lower)edgeconnectingvk−1andvk.Inthisexampletheperformanceofthealgorithmappearedtobeindependentoftheactualchoiceofthebasis;however,ingeneralwedonotexpectthisbehavior.TwoversionsofthealgorithmofAwerbuchandKleinberg[4]werealsosimulated.Withitsoriginalparametersetting(AwKl),thealgorithmdidnotperformwell.However,afteroptimizingitsparametersoﬀ-line(AwKltuned),substantiallybetterperformancewasachieved.Thenormalizedregretoftheabovealgorithms,averagedover30runs,aswellastheregretoftheﬁxedpathsinthegraphareshowninFigure2.Althoughallalgorithmsshowedbetterperformancethantheboundfortheedge-banditalgo-rithm,thelattershowedtheexpectedsuperiorperformanceinthesimulations.Furthermore,ouralgorithmfortherestrictedshortestpathproblemoutperformedAwerbuchandKleinberg’s(AwKl)algorithm,whi

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

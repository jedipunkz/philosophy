---
source: "https://arxiv.org/abs/1006.3302v2"
title: "Quasirandom Load Balancing"
author: "Tobias Friedrich, Martin Gairing, Thomas Sauerwald"
year: "2010"
publication: "arXiv preprint / cs.DS"
download: "https://arxiv.org/pdf/1006.3302v2"
pdf: "https://arxiv.org/pdf/1006.3302v2"
captured_at: "2026-05-09T12:06:49Z"
updated_at: "2026-05-09T12:06:49Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Friedrich Nietzsche"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Quasirandom Load Balancing

- 著者: Tobias Friedrich, Martin Gairing, Thomas Sauerwald
- 年: 2010
- 掲載情報: arXiv preprint / cs.DS
- 情報源: [arxiv](https://arxiv.org/abs/1006.3302v2)
- ダウンロード: https://arxiv.org/pdf/1006.3302v2
- PDF: https://arxiv.org/pdf/1006.3302v2

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

We propose a simple distributed algorithm for balancing indivisible tokens on graphs. The algorithm is completely deterministic, though it tries to imitate (and enhance) a random algorithm by keeping the accumulated rounding errors as small as possible. Our new algorithm surprisingly closely approximates the idealized process (where the tokens are divisible) on important network topologies. On d-dimensional torus graphs with n nodes it deviates from the idealized process only by an additive constant. In contrast to that, the randomized rounding approach of Friedrich and Sauerwald (2009) can deviate up to Omega(polylog(n)) and the deterministic algorithm of Rabani, Sinclair and Wanka (1998) has a deviation of Omega(n^{1/d}). This makes our quasirandom algorithm the first known algorithm for this setting which is optimal both in time and achieved smoothness. We further show that also on the hypercube our algorithm has a smaller deviation from the idealized process than the previous algorithms.

## PDF Text

arXiv:1006.3302v2 [cs.DS] 19 Apr 2013
QUASIRANDOMLOADBALANCING∗TOBIASFRIEDRICH†,MARTINGAIRING‡,ANDTHOMASSAUERWALD§Abstract.Weproposeasimpledistributedalgorithmforbalancingindivisibletokensongraphs.Thealgorithmiscompletelydeterministic,thoughittriestoimitate(andenhance)arandomizedalgorithmbykeepingtheaccumulatedroundingerrorsassmallaspossible.Ournewalgorithmsurprisinglycloselyapproximatestheidealizedprocess(wherethetokensaredivisible)onimportantnetworktopologies.Ond-dimensionaltorusgraphswithnnodesitdeviatesfromtheidealizedprocessonlybyanadditiveconstant.Incontrast,therandomizedroundingapproachofFriedrichandSauerwald[11]candeviateuptoΩ(polylog(n))andthedeterministicalgorithmofRabani,SinclairandWanka[33]hasadeviationofΩ(n1/d).Thismakesourquasirandomalgorithmtheﬁrstknownalgorithmforthissettingwhichisoptimalbothintimeandachievedsmoothness.Wefurthershowthatonthehypercubeaswell,ouralgorithmhasasmallerdeviationfromtheidealizedprocessthanthepreviousalgorithms.Toprovetheseresults,wederiveseveralcombinatorialandprobabilisticresultsthatwebelievetobeofindependentinterest.Inparticular,weshowthatﬁrst-passageprobabilitiesofarandomwalkonapathwitharbitraryweightscanbeexpressedasaconvolutionofindependentgeometricprobabilitydistributions.1.Introduction.Loadbalancingisarequisitefortheeﬃcientutilizationofcom-putationalresourcesinparallelanddistributedsystems.Theaimistoreallocatetheloadsuchthatafterward,eachnodehasapproximatelythesameload.Loadbalancingproblemshavevariousapplications,e.g.,forscheduling[37],routing[5],andnumericalcomputation[38,39].Typically,loadbalancingalgorithmsiterativelyexchangeloadalongedgesofanundirectedconnectedgraph.Inthenaturaldiﬀusionparadigm,anarbitraryamountofloadcanbesentalongeachedgeateachstep[31,33].Fortheidealizedcaseofdivisibleload,apopulardiﬀusionalgorithmistheﬁrst-order-schemebySubramanianandScherson[36]whoseconvergencerateisfairlywellcapturedintermsofthespectralgap[27].However,formanyapplicationstheassumptionofdivisibleloadmaybeinvalid.Therefore,weconsiderthediscretecasewheretheloadcanonlybedecomposedintoin-divisibleunit-sizetokens.Averynaturalquestionishowmuchthisintegralityassump-tiondecreasestheeﬃciencyofloadbalancing.Infact,ﬁndingaprecisequantitativerelationshipbetweenthediscreteandtheidealizedcaseisanopenproblemposedbymanyauthors,e.g.,[9,11,15,16,28,31,33,36].AsimplemethodforapproximatingtheidealizedprocesswasanalyzedbyRabani,Sinclair,andWanka[33].Theirapproach(whichwewillcall“RSWalgorithm”)istorounddownthefractionalﬂowoftheidealizedprocess.Theyintroduceaveryusefulpa-rameterofthegraphcalledlocaldivergenceandprovethatitgivestightupperboundsonthedeviationbetweentheidealizedprocessandtheirdiscreteprocess.However,onedrawbackoftheRSWalgorithmisthatitcanendupinratherunbalancedstates(cf.Proposition6.1).Toovercomethisproblem,FriedrichandSauerwaldanalyzedanewalgorithmbasedonrandomizedrounding[11].Onmanygraphs,thisalgorithmapproxi-matestheidealizedcasemuchbetterthanRSWalgorithm’sapproachofalwaysroundingdown.Anaturalquestioniswhetherthisrandomizedalgorithmcanbederandomized
∗Apreliminaryconferenceversion[12]appearedinthe21stACM-SIAMSymposiumonDiscreteAlgorithms(SODA2010).Afullversion[13]appearedinSIAMJournalonComputing.ThisworkwasdonewhileallauthorswerepostdoctoralfellowsattheInternationalComputerScienceInstitute(ICSI),BerkeleysupportedbytheGermanAcademicExchangeService(DAAD).†Friedrich-Schiller-Universit?tJena,Germany.‡UniversityofLiverpool,UnitedKingdom.§Max-Planck-Institutf¨urInformatik,Saarbr¨ucken,Germany.1
withoutsacriﬁcingonitsperformance.Forthegraphsconsideredinthiswork,weanswerthisquestionintheaﬃrmative.Weintroduceaquasirandomloadbalancingalgorithmwhichroundsupordowndeterministicallysuchthattheaccumulatedroundingerrorsoneachedgeareminimized.Ourapproachfollowstheconceptofquasirandomnessasitdeterministicallyimitatestheexpectedbehaviorofitsrandomcounterpart.OurResults.Wefocusontwonetworktopologies:hypercubesandtorusgraphs.Bothhavebeenintensivelystudiedinthecontextofloadbalancing(seee.g.,[11,14,20,32,33]).Wemeasurethesmoothnessoftheloadbytheso-calleddiscrepancy(seee.g.[9,11,16,33])whichisthediﬀerencebetweenthemaximumandminimumloadamongallnodes.Ford-dimensionaltorusgraphsweprovethatourquasirandomalgorithmapproxi-matestheidealizedprocessuptoanadditiveconstant(Theorem5.4).Moreprecisely,forallinitialloaddistributionsandtimesteps,theloadofanyvertexinthediscreteprocessdiﬀersfromtherespectiveloadintheidealizedprocessonlybyaconstant.Thisholdsevenfornon-uniformtorusgraphswithdiﬀerentsidelengths(cf.Deﬁnition5.1).FortheuniformtorusgraphourresultsaretobecomparedwithadeviationofΩ(polylog(n))fortherandomizedroundingapproach(Theorem6.3)andΩ(n1/d)fortheRSWalgorithm(Proposition6.1).Hence,despitethefactthatourapproachisdeterministic,itstillimprovesoveritsrandomcounterpart.StartingwithaninitialdiscrepancyofK,theidealizedprocessreachesaconstantdiscrepancyafterO(n2/dlog(Kn))steps(cf.Corol-lary3.2).Hencethesameholdsforourquasirandomalgorithm,whichmakesittheﬁrstalgorithmforthediscretecasewhichisoptimalbothintimeanddiscrepancyuptoconstantfactors.Forthehypercube,weprovethatthedeviationofourquasirandomalgorithmfromtheidealizedprocessisbetweenΩ(logn)andO(log3/2n)(Theorem4.2).Notethattheanalysisfortheupperboundinthispaperﬁxesabuginthecorrespondingproofoftheconferenceversion[12],whereweclaimedanupperboundofO(logn).ForthehypercubewealsoshowthatthedeviationoftherandomapproachisΩ(logn)(Theorem6.2)whilethedeviationoftheRSWalgorithmisΩ(log2n)(Proposition6.1).Again,ourquasirandomalgorithmissubstantiallybetterthantheRSWalgorithm[33].OurTechniques.Insteadofanalyzingourquasirandomalgorithmdirectly,weexam-ineanewgenericclassofloadbalancingalgorithmsthatwecallboundederrordiﬀusion(BED).Roughlyspeaking,inaBEDalgorithmtheaccumulatedroundingerroroneachedgeisboundedbysomeconstantatalltimes.Thisclassincludesourquasirandomalgorithm.Thestartingpointof[33]and[11]aswellasthatofourpaperistoexpressthedeviationfromtheidealizedcasebyacertainsumofweightedroundingerrors(equa-tion(3.1)).Inthissum,theroundingerrorsareweightedbytransitionprobabilitiesofacertainrandomwalk.Roughlyspeaking,Rabanietal.[33]estimatethissumdirectlybyaddingupalltransitionprobabilities.Intherandomizedapproachof[11],thesumisboundedbyChernoﬀ-typeinequalitiesrelyingonindependentroundingdecisions.Wetakeacompletelydiﬀerentapproachandprovethatthetransitionprobabilitiesbetweentwoﬁxedverticesareunimodalintime(cf.Theorem4.9forthehypercube).Thisallowsustoupperboundthecompletesumbyitsmaximalsummand(Lemma3.6)forBEDal-gorithms.Theintriguingcombinatorialpropertyofunimodalityistheheartofourproofandseemstobethemainreasonwhywecanoutperformthepreviousapproaches.Eventhoughunimodalityhasaone-linedeﬁnition,ithasbecomeapparentthatprovingitcanbeaverychallengingtaskrequiringintricatecombinatorialconstructionsorreﬁnedmathematicaltools(seee.g.Stanley’ssurvey[35]).Itturnsoutthatthisisalsothecaseforthetransitionprobabilitiesoftorusgraphsandhypercubesconsideredhere.Thereasonisthatexplicitformulasseemtobein-2
tractable,andtypicalapproximations(e.g.Poissonization[6])arefartooloosetocom-pareconsecutivetransitionprobabilities.Forthed-dimensionaltorus,weusealocalcentrallimittheoremtoapproximatethetransitionprobabilitiesbyamultivariatenor-maldistributionwhichisknowntobeunimodal.Onhypercubes,theabovemethodfails,asseveralinequalitiesforthetorusgraphareonlytrueforconstantd.However,wecanemploytheadditionalsymmetriesofthehypercubetoproveunimodalityofthetransitionprobabilitiesbyrelatingittoarandomwalkonaweightedpath.Somewhatsurprisingly,thisintriguingpropertywaspreviouslyunknown,althoughrandomwalksonhypercubeshavebeenintensivelystud-ied(seee.g.[6,22,29]).Weprovethisunimodalityresultbyestablishinganinterestingresultconcerningﬁrst-passageprobabilitiesofarandomwalkonpathswitharbitrarytransitionprob-abilities:Iftheloopprobabilitiesareatleast1/2,thentheﬁrst-passageprobabilitydistributioncanbeexpressedasaconvolutionofindependentgeometricdistributions.Inparticular,thisimpliesthattheseprobabilitiesarelog-concave.Reducingtherandomwalkonahypercubetoarandomwalkonaweightedpath,weobtaintheresultthatthetransitionprobabilitiesonthehypercubeareunimodal.Estimatingthemaximumprob-abilitiesviaaballs-and-bins-process,weﬁnallyobtainourupperboundonthedeviationforthehypercube.Webelievethatourprobabilisticresultforpathsisofindependentinterest,asran-domwalksonthepathsareamongthemostextensivelystudiedstochasticprocesses.Moreover,manyanalysesofrandomizedalgorithmscanbereducedtosuchrandomwalks(seee.g.[30,Thm.6.1]).RelatedWork.IntheapproachofEls¨asserandSauerwald[8]certaininteractingrandomwalksareusedtoreducetheloaddeviation.ThisrandomizedalgorithmachievesaconstantadditiveerrorbetweenthemaximumandaverageloadonhypercubesandtorusgraphsintimeO(log(Kn)/(1−λ2)),whereλ2isthesecondlargesteigenvalueofthediﬀusionmatrix.However,incontrasttoourdeterministicalgorithm,thisalgorithmislessnaturalandmorecomplicated(e.g.,thenodesmusthaveanaccurateestimateoftheaverageload).Aielloetal.[1]andGhoshetal.[16]studiedbalancingalgorithmswhere,ineachtimestep,atmostonetokenistransmittedovereachedge.Duetothisrestriction,thesealgorithmstakesubstantiallymoretime,i.e.,theyrunintimeatleastlinearintheinitialdiscrepancyK.Nonetheless,thebestknownboundsonthediscrepancyareonlypolynomialinnforthetorusandΩ(log5n)forthehypercube[16].Inanothercommonmodel,nodesareonlyallowedtoexchangeloadwithatmostoneneighborineachtimestep,seee.g.,[11,15,33].Infact,theaforementionedrandomizedroundingapproach[11]wasanalyzedinthismodel.However,theideaofrandomlyroundingthefractionalﬂowsuchthattheexpectederroriszeronaturallyextendstoourdiﬀusivesettingwhereanodemayexchangeloadwithallneighborssimultaneously.Quasirandomnessdescribesadeterministicprocesswhichimitatescertainpropertiesofarandomprocess.Ourquasirandomloadbalancingalgorithmimitatesthepropertythatroundingupanddowntheﬂowbetweentwoverticesoccursroughlyequallyoften,usingadeterministicprocesswhichminimizestheseroundingerrorsdirectly.Inthisway,wekeepthedesiredpropertythatthe“expected”accumulatedroundingerroriszero,butremovealmostallofits(undesired)variance.Similarconceptshavebeenusedfordeterministicrandomwalks[4],externalmergesort[2],andquasirandomrumorspread-ing[7].Thelatterworkpresentsaquasirandomalgorithmwhichisabletobroadcastapieceofinformationatleastasfastasitsrandomcounterpartonthehypercubeandmostrandomgraphs.However,inthecaseofrumorspreading,thequasirandomprotocolisjustslightlyfasterthantherandomprotocol,whilethequasirandomload-balancing3
algorithmpresentedheresubstantiallyoutperformsitsrandomcounterpart.Organizationofthepaper.InSection2,wegiveadescriptionofourboundederrordiﬀusion(BED)model.Forabettercomparison,wepresentsomeresultsforthepreviousalgorithmsof[11]and[33]inSection6.InSection3,weintroduceourbasicmethodwhichisusedinSections4and5toanalyzeBEDalgorithmsonhypercubesandtorusgraphs,respectively.2.Modelandalgorithms.Weaimforbalancingloadonaconnected,undirectedgraphG=(V,E).Denotebydeg(i)thedegreeofnodei∈Vandlet∆=∆(G)=maxi∈Vdeg(i)bethemaximumdegreeofG.Thebalancingprocessisgovernedbyanergodic,doubly-stochasticdiﬀusionmatrixPwithPi,j=



1
2∆if{i,j}∈E,1−deg(i)
2∆ifi=j,0otherwise.Letx(t)betheload-vectoroftheverticesatstept(ormoreprecisely,afterthecompletionofthebalancingprocedureatstept).Thediscrepancyofsucha(row)vectorxismaxi,j(xi−xj),andthediscrepancyatstep0iscalledinitialdiscrepancyK.Theidealizedprocess.Inonetimestepeachpair(i,j)ofadjacentverticesshiftsdivisibletokensbetweeniandj.Wehavethefollowingiteration,x(t)=x(t−1)Pandinductively,x(t)=x(0)Pt.Equivalently,foranyedge{i,j}∈Eandstept,theﬂowfromitojatsteptisPi,jx(t−1)i−Pj,ix(t−1)j.NotethatthesymmetryofPimpliesthatfort→∞,x(t)convergestowardstheuniformvector(1/n,1/n,...,1/n).Thediscreteprocess.Therearediﬀerentwaystohandlenon-divisibletokens.Wedeﬁnethefollowingboundederrordiﬀusion(BED)model.LetΦ(t)i,jdenotetheintegralﬂowfromitojattimet.AsΦ(t)i,j=−Φ(t)j,i,wehavex(t)i=x(t−1)i−Pj:{i,j}∈EΦ(t)i,j.Lete(t)i,j:= Pi,jx(t−1)i−Pj,ix(t−1)j−Φ(t)i,jbetheexcessloadallocatedtoiasaresultofroundingonedge{i,j}intimestept.Anegativevalueofe(t)i,jsigniﬁesadeﬁcitofload.Notethatforallverticesi,x(t)i=(x(t−1)P)i+Pj:{i,j}∈Ee(t)i,j.Now,letΛbeanupperboundfortheaccumulatedroundingerrors(deviationfromtheidealizedprocess),thatis,
Pt s=1e(s)i,j
6Λforallt∈Nand{i,j}∈E.AllourboundsstillholdifΛisafunctionofnand/ort,butweonlysaythatanalgorithmisaBEDalgorithmifΛisaconstant.ForPi,jx(t)i>Pj,ix(t)j,ournewquasirandomdiﬀusionalgorithmchoosestheﬂowΦ(t)i,jfromitojtobeeitherΦ(t)i,j=Pi,jx(t)i−Pj,ix(t)jorΦ(t)i,j=Pi,jx(t)i−Pj,ix(t)jsuchthat
Pt s=1e(s)i,j
isminimized.ThisyieldsaBEDalgorithmwithΛ61/2,whichcanbeimplementedwithO(log∆)storageperedge.Notethatonecanimaginevariousothernatural(deterministicorrandomized)BEDalgorithms.Toachievethis,thealgorithmonlyhastoensurethattheerrorsdonotadduptomorethanaconstant.Withtheabovenotation,theRSWalgorithmusesΦ(t)i,j=Pi,jx(t)i−Pj,ix(t)j,providedthatPi,jx(t)i>Pj,ix(t)j.Inotherwords,theﬂowoneachedgeisalwaysroundeddown.InourBEDframeworkthiswouldimplyaΛoforderTafterTtimesteps.AsimplerandomizedroundingdiﬀusionalgorithmchoosesforPi,jx(t)i>Pj,ix(t)jtheﬂowΦ(t)i,jastherandomizedroundingofPi,jx(t)i−Pj,ix(t)j,thatis,itroundsupwithprobability(Pi,jx(t)i−Pj,ix(t)j)−Pi,jx(t)i−Pj,ix(t)jandroundsdownotherwise.ThistypicallyachievesanerrorΛoforder√
T
afterTtimesteps.4
HandlingNegativeLoads.Unlessthereisalowerboundontheminimumloadofavertex,negativeloadsmayoccurduringthebalancingprocedure.Inwhatfollows,wedescribeasimpleapproachforcopingwiththisproblem.ConsideragraphGforwhichwecanproveadeviationofatmostγfromtheidealizedprocess.Letx(0)betheinitialloadvectorwithanaverageloadof¯x.Thenatthebeginningofthebalancingprocedure,eachnodegeneratesγadditional(virtual)tokens.Duringthebalancingprocedure,thesetokensareregardedascommontokens,butattheendtheyareignored.Firstobservethatsincetheminimumloadateachnodeintheidealizedprocessisatleastγ,itfollowsthatateachstep,everynodehasatleastaloadofzerointhediscreteprocess.Sinceeachnodehasaloadof¯x+O(γ)attheend,weendupwithaloaddistributionwherethemaximumloadisstill¯x+O(γ)(ignoringthevirtualtokens).3.Basicmethodtoanalyzeourquasirandomalgorithm.ToboundruntimeanddiscrepancyofaBEDalgorithm,wealwaysboundthedeviationbetweentheide-alizedmodelandthediscretemodel,whichisanimportantmeasureinitsownright.Forthisdiscussion,letx(t)`denotetheloadonvertex`insteptinthediscretemodelandξ(t)`denotetheloadonvertex`insteptintheidealizedmodel.Weassumethatthediscreteandidealizedmodelstartwiththesameinitialload,thatis,x(0)=ξ(0).AsderivedinRabanietal.[33],theirdiﬀerencecanbewrittenasx(t)`−ξ(t)`=t−1Xs=0X[i:j]∈Ee(t−s)i,j(Ps
`,i−Ps
`,j).(3.1)where[i:j]referstoanedge{i,j}∈Ewithi<j,and“<”issomearbitrarybutﬁxedorderingontheverticesV.Itwillbesuﬃcienttoboundequation(3.1),astheconvergencespeedoftheidealizedprocesscanbeboundedintermsofthesecondlargesteigenvalue.Theorem3.1(e.g.,[33,Thm.1]).Onallgraphswithsecondlargesteigenvalueinabsolutevalueλ2=λ2(P),theidealizedprocesswithdivisibletokensreducesaninitialdiscrepancyKto`within2
1−λ2ln Kn2
`timesteps.Asλ2=1−Θ(log−1n)forthehypercubeandλ2=1−Θ(n−2/d)forthed-dimensionaltorus[15],oneimmediatelyobtainsthefollowingcorollary.Corollary3.2.TheidealizedprocessreducesaninitialdiscrepancyofKto1withinO(n2/dlog(Kn))timestepsonthed-dimensionaltorusandwithinO(lognlog(Kn))timestepsonthehypercube.Animportantpropertyofallexaminedgraphclasseswillbeunimodalityorlog-concavityofcertaintransitionprobabilities.Definition3.3.Afunctionf:N→R≥0islog-concaveiff(i)2≥f(i−1)·f(i+1)foralli∈N>0.Definition3.4.Afunctionf:N→Risunimodalifthereisas∈Nsuchthatf|x≤saswellasf|x≥saremonotone.Log-concavefunctionsaresometimesalsocalledstronglyunimodal[24].Wesumma-rizesomeclassicalresultsregardinglog-concavityandunimodality.Fact3.5.1.Letfbealog-concavefunction.Then,fisalsoaunimodalfunction(e.g.[23,24]).5
2.Hoggar’stheorem[19]:Letfandgbelog-concavefunctions.Thentheirconvo-lution(f∗g)(k)=Pk i=0f(i)g(k−i)isalsolog-concave.3.Letfbealog-concavefunctionandgbeaunimodalfunction.Thentheircon-volutionf∗gisaunimodalfunction[24].Ourinterestinunimodalityisbasedonthefactthatanalternatingsumoveraunimodalfunctioncanbeboundedbyitsmaximum.Moreprecisely,foranon-negativeandunimodalfunctionf:X→Randt0,...,tk∈Xwitht06···6tk,thefollowingholds:

kXi=0(−1)if(ti)

≤maxx∈Xf(x).ThisisaspecialcaseofAbel’sinequality.Wegeneralizebothinthefollowinglemma.Lemma3.6.Letf:X→Rbenon-negativewithX⊆R.LetA0,...,Ak∈Randt0,...,tk∈Xsuchthatt06···6tkand|Pk i=aAi|6kforall06a6k.Iffhas`localextrema,then

kXi=0Aif(ti)

≤(`+1)kkmaxj=0f(tj).Proof.Letusstartwiththeassumptionthatf(ti),06i6k,ismonotoneincreasing.Withf(t−1):=0,ittheniseasytoseethat
Pk i=0Aif(ti)
=
Pk i=0Pi j=0Ai(f(tj)−f(tj−1))
=
Pk j=0Pk i=jAi(f(tj)−f(tj−1))
6Pk j=0
f(tj)−f(tj−1)

Pk i=jAi
6Pk j=0
f(tj)−f(tj−1)
k=kkmaxj=0f(tj).Thesameholdsiff(ti),06i6k,ismonotonedecreasing.Iff(x)has`localextrema,wesplitthesumin(`+1)partssuchthatf(x)ismonotoneoneachpartandapplytheabovearguments.
RandomWalks.Toexaminethediﬀusionprocess,itwillbeusefultodeﬁnearandomwalkbasedonP.Foranypairofverticesi,j,Pt i,jistheprobabilitythatarandomwalkguidedbyPstartingfromiislocatedatjatstept.InSection4itwillbeusefultosetPi,j(t):=Pt i,jandtodenotewithfi,j(t)fori6=jtheﬁrst-passageprobabilities,thatis,theprobabilitythatarandomwalkstartingfromivisitsthevertexjatsteptfortheﬁrsttime.4.Analysisonthehypercube.Weﬁrstgivethedeﬁnitionofthehypercube.Definition4.1.Ad-dimensionalhypercubewithn=2dverticeshasvertexsetV={0,1}dandedgesetE={{i,j}|iandjdiﬀerinonebit}.Inthissectionweprovethefollowingresult.
Theorem4.2.Forallinitialloadvectorsonthed-dimensionalhypercubewithnvertices,thedeviationbetweentheidealizedprocessandadiscreteprocesswithaccu-mulatedroundingerrorsatmostΛisupperboundedbyO(Λ·(logn)3/2)atalltimes,andthereareloadvectorsforwhichthisdeviationcanbe(logn)/2.6
RecallthatforBEDalgorithmsΛ=O(1).WithTheorem3.1itfollowsthatanyBEDalgorithm(andinparticularourquasirandomalgorithm)reducesthediscrepancyofanyinitialloadvectorwithdiscrepancyKtoO(logn)withinO(lognlog(Kn))timesteps.4.1.Log-concavepassagetimeonpaths.ToproveTheorem4.2,weﬁrstcon-sideradiscrete-timerandomwalkonapathP=(0,1,...,d)startingatnode0.Wemakeuseofaspecialgeneratingfunction,calledz-transform.Thez-transformofafunctiong:N7→R≥0isdeﬁnedbyG(z)=P∞
i=0g(i)·z−i.Wewillusethefactthataconvolutionreducestomultiplicationinthez-plane.Moreformally,ifG1andG2arethez-transformsofg1andg2,respectively,thentheproductG1·G2isthez-transformoftheirconvolutiong1∗g2.Insteadofthez-transformonecouldcarryoutasimilaranalysisusingtheprobabilitygeneratingfunction.Wechoosetousethez-transformheresinceitleadstoslightlysimplerarithmeticexpressions.Ouranalysisalsousesthegeometricdistributionwithparameterp,whichisdeﬁnedbyGeo(p)(t)=(1−p)t−1pfort∈N\{0}andGeo(p)(0)=0.ItiseasytocheckthatGeo(p)islog-concave.Moreover,thez-transformofGeo(p)is∞Xi=1Geo(p)(i)·z−i=p z−(1−p).Foreachnodei∈P,letαibetheloopprobabilityatnodeiandβibetheupwardprobability,i.e.,theprobabilityofmovingtonodei+1.Then,thedownwardprobabilityatnodeiis1−αi−βi.Wecanassumethatβi>0foralli∈P\{d}.Weareinterestedintheﬁrst-passageprobabilitiesf0,d(t).Observethatf0,d(t)=(f0,1∗f1,2∗···∗fd−1,d)(t).(4.1)Inthefollowingargument,wewillshowthatf0,d(t)islog-concave.Indeed,weshowamuchstrongerresult:Theorem4.3.ConsiderarandomwalkonapathP=(0,1,...,d)startingatnode0.Ifαi≥1
2forallnodesi∈P,thenf0,dcanbeexpressedasconvolutionofdindependentgeometricdistributions.Asthegeometricdistributionislog-concaveandtheconvolutionoflog-concavefunc-tionsisagainlog-concave(cf.Fact3.5),weimmediatelygetthefollowingcorollary.Corollary4.4.ConsiderarandomwalkonapathP=(0,1,...,d)startingatnode0.Ifαi≥1
2forallnodesi∈P,thenf0,d(t)islog-concaveint.NotethatTheorem4.3followsdirectlyfromTheorem1.2ofFill[10].AsTheorem4.3isacrucialingredientforprovingourmainresult(Theorem4.2)forthehypercube,wegiveadiﬀerentalternativeproofofthestatement.WhileFill’sproofispurelystochastic,ourproofisbasedonfunctionalanalysisofthez-transform.Ouranalysisforthediscrete-timerandomwalkshouldalsobecomparedwithKeilson’sanalysisofthecontinuous-timeprocess[23].Thecontinuous-timeprocesswasindependentlyconsideredevenearlierbyKarlinandMcGregor[21].Beforeprovingthetheorem,wewillshowhowtoobtainf0,d(t)byarecursivear-gument.Todothis,supposeweareatnodei∈P\{d}.Thenextstepisaloopwithprobabilityαi.Moreover,thenextsubsequentnon-loopmoveendsati+1withprobabilityβi
1−αiandati−1withprobability1−βi−αi
1−αi.Thus,foralli∈P\{d},fi,i+1(t)=βi
1−αi·Geo(1−αi)(t)+1−βi−αi
1−αi·(Geo(1−αi)∗fi−1,i∗fi,i+1)(t),7
withcorrespondingz-transformFi,i+1(z)=βi
1−αi·1−αi z−αi+1−βi−αi
1−αi·1−αi z−αi·Fi−1,i(z)·Fi,i+1(z).RearrangingtermsyieldsFi,i+1(z)=βi z−αi−(1−βi−αi)·Fi−1,i(z),(4.2)foralli∈P\{d}.SoFi,i+1(z)isobtainedrecursivelywithF0,1(z)=β0
z−(1−β0).Finally,thez-transformofequation(4.1)isF0,d(z)=F0,1(z)·F1,2(z)·...·Fd−1,d(z).(4.3)WewillnowprovesomepropertiesofFi,i+1(z)fori∈P\{d}.Lemma4.5.Exceptforsingularities,Fi,i+1(z)ismonotonedecreasinginz,foralli∈P\{d}.Proof.Wewillshowtheclaimbyinductiononi.Itiseasytoseethattheclaimholdsforthebasecase(i=0)sinceF0,1(z)=β0
z−(1−β0).AssumeinductivelythattheclaimholdsforFi−1,i(z).With1−βi−αi≥0thisdirectlyimpliesthatthedenominatorofequation(4.2)isincreasinginz.TheclaimforFi,i+1(z)follows.
Lemma4.6.Foralli∈P\{d},Fi,i+1(z)hasexactlyi+1poleswhichareallintheinterval(0,1).ThepolesofFi,i+1(z)aredistinctfromthepolesofFi−1,i(z).Proof.Beforeprovingtheclaimsofthelemma,wewillshowthatFi,i+1(0)≥−1andFi,i+1(1)=1foralli∈P\{d}.Observe,thatF0,1(0)=β0
−(1−β0)=1−α0
−α0≥−1,sinceα0≥1
2.AlsoobservethatF0,1(1)=1.Assume,inductivelythatFi−1,i(0)≥−1andFi−1,i(1)=1.Thenwithequation(4.2),Fi,i+1(0)≥βi
−αi−(1−βi−αi)·(−1)=βi
1−2αi−βi≥−1,since1−2αi≤0.Moreover,Fi,i+1(1)=βi
1−αi−(1−αi−βi)=1.Thus,Fi,i+1(0)≥−1andFi,i+1(1)=1foralli∈P\{d}.Wewillnowshowtheclaimsofthelemmabyinduction.Forthebasecase,observethatF0,1(z)=β0
z−(1−β0)hasonepoleatz=1−β0>0andF−1,0isnotdeﬁned.Thisimpliestheclaimfori=0.SupposetheclaimholdsforFi−1,i(z),andletz1,z2,...zibethepolesofFi−1,i(z).Withoutlossofgenerality,weassume0<z1<z2<···<zi<1.Letgi(z)bethedenominatorofequation(4.2),thatis,gi(z):=z−αi−(1−βi−αi)·Fi−1,i(z).Observethat(i)gi(z)hasthesamesetofpolesasFi−1,i(z),(ii)limz→−∞gi(z)=−∞,and(iii)limz→∞gi(z)=∞.Byequation(4.2),Fi,i+1(z)hasitspolesatthezerosofgi(z).Lemma4.5showsthatineachinterval(zj,zj+1)with1≤j≤i−1,gi(z)isincreasinginz.Usingfact(i)thisimpliesthatgi(z)hasexactlyonezeroineachinterval(zj,zj+1).ThusFi,i+1(z)hasexactlyonepoleineachinterval(zj,zj+1).Similarly,Lemma4.5,togetherwithfacts(i),(ii)and(iii),impliesthatFi,i+1(z)hasexactlyonepole,sayz0,intheinterval[−∞,z1)andonepole,sayz00,intheinterval(zi,∞].ThisimpliesthatFi,i+1(z)hasexactlyi+1poleswhicharealldistinctfromthepolesofFi−1,i(z).Itremainstoshowthatz0>0andz00<1.8
SinceFi−1,i(0)≥−1andlimz→−∞Fi−1,i(z)=−0,itfollowsfromLemma4.5that−1≤Fi−1,i(z)≤0forallrealz<0.Sogi(z)<0forallrealz<0.Itfollowsthatz0>0.Similarly,sinceFi−1,i(1)=1andlimz→∞Fi−1,i(z)=+0,itfollowsbyLemma4.5that0≤Fi−1,i(z)≤1forallrealz>1.Sogi(z)>0forallrealz>1.Itfollowsthatz00<1.Thisﬁnishestheproofofourinductivestep.Theclaimfollows.
Thedistinctnessofthei+1polesofFi−1,i(z),establishedinLemma4.6,iscrucialfortheproofofthefollowinglemma.Lemma4.7.Let(bj,i)i j=0bethepolesofFi,i+1(z),i∈P\{d},anddeﬁnePi(z)=Qi j=0(z−bj,i).ThenFi,i+1(z)=βi·Pi−1(z)
Pi(z)foralli∈P\{d}.Proof.Ourproofisbyinductiononi.Forthebasecase(i=0),observethatP−1(z)=1andthusF0,1(z)hasthedesiredform.SupposetheclaimholdsforFi−1,i(z).Thenequation(4.2)impliesFi,i+1(z)=βi z−αi−(1−βi−αi)·βi−1·Pi−2(z)
Pi−1(z)=βi·Pi−1(z)
(z−αi)·Pi−1(z)−(1−βi−αi)·βi−1·Pi−2(z).(4.4)Observethat(z−αi)·Pi−1(z)isapolynomialofdegreei+1wheretheleadingtermhasacoeﬃcientof1.Thisalsoholdsforthedenominatorofequation(4.4),sincethere,weonlysubtractapolynomialoforderi−1.ByLemma4.6weknowthatFi,i+1(z)hasexactlyi+1realpositivepoleswhicharealldistinct.Sothedenominatorofequation(4.4)hasexactlyd+1zeros–thebj,i’s.Theonlypolynomialoforderi+1thathasexactlythosezerosandleadingcoeﬃcient1isPi(z).Itfollowsthatthedenominatorofequation(4.4)isequaltoPi(z).Theclaimfollows.
WearenowreadytoproveTheorem4.3.ProofofTheorem4.3.Byequation(4.3)andLemma4.7,wegetF0,d(z)=d−1Yi=0Fi,i+1(z)=d−1Yi=0βi·Pi−1(z)
Pi(z)=Qd−1i=0βi
Pd−1(z)=Kd·d−1Y
i=01−bi,d−1
z−bi,d−1,where(bi,d−1)d−1i=0arethepolesofFd−1,d(z)asdeﬁnedinLemma4.7andKd=Qd−1i=0βi
1−bi,d−1.ByLemma4.6,bi,d−1∈(0,1)foralli.Nowforeachitheterm1−bi,d−1
z−bi,d−1isthez-transformofthegeometricdistributionwithparameter1−bi,d−1,i.e.,Geo(1−bi,d−1)(t).Thus,f0,d(t)canbeexpressedastheconvolutionofdindependentgeometricdistri-butionsf0,d(t)=Kd·[Geo(1−b0,d−1)∗Geo(1−b1,d−1)∗...∗Geo(1−bd−1,d−1)](t).Moreover,sincef0,disaprobabilitydistributionovertandtheconvolutionofprobabilitydistributionsisagainaprobabilitydistribution,wehaveKd=1.Thetheoremfollows.
Itshouldbenotedthatitfollowsfrom[10,Theorem1.2]thattheparameters(bi,d−1)d−1i=0inthegeometricdistributionsaretheeigenvaluesoftheunderlyingtran-sitionmatrix.RecallthatouraimistoproveunimodalityforthefunctionPtΩ0,j(int).UsingthesimpleconvolutionformulaP0,j=f0,j∗Pj,jandthelog-concavityoff0,j,itsuﬃces9
toprovethatPj,jisunimodal(cf.Fact3.5).Next,wewillprovethatPj,jisalsonon-increasingint.Lemma4.8.LetPbethe(d+1)×(d+1)-transitionmatrixdeﬁninganergodicMarkovchainonapathP=(0,...,d).IfPii≥1
2forall06i6dthenforall06i6d,Pt i,iisnon-increasingint.Proof.ItiswellknownthatergodicMarkovchainsonpathsaretimereversible(seee.g.Section4.8ofRoss[34]).Toseethis,letπ=(π0,...,πd)bethestationarydistribution.Thenforall0≤i≤d−1therateatwhichtheprocessgoesfromitoi+1(namely,πiPi,i+1)isequaltotherateatwhichtheprocessgoesfromi+1toi(namely,πi+1Pi+1,i).Thus,Pistime-reversible.Oneusefulpropertyofatime-reversiblematrixisthatitseigenvaluesareallreal.TheGer˘sgorindisctheoremstatesthateveryeigenvalueλj,0≤j≤d,satisﬁesthecondition|λj−Pii|≤1−Pii,forsome0≤i≤d.SincePii≥1
2,thisdirectlyimpliesthatalleigenvaluesareintheinterval[0,1].Itiswell-knownthatthereisanorthonormalbaseofRd+1whichisformedbytheeigenvectorsv0,v1,...,vd(seee.g.[18]).Thenforanyn-dimensionalvectorw∈Rd+1,w=Pd j=0hw,vjivj,whereh·,·idenotestheinnerproduct.Applyingthistothei-thunitvectoreiandusing[·]itodenotethei-thentryofavectorinRd+1,weobtainei=dXj=0hei,vjivj=dXj=0[vj]ivj.Thus,Ptei=PtdXj=0[vj]ivj=dXj=0[vj]iPtvj=dXj=0[vj]iλt jvjandﬁnallyPt i,i=Pteii=dXj=0[vj]iλt j[vj]i=dXj=0λt j[vj]2
i,whichisnon-increasingintas[vj]i∈Rand0≤λj≤1forall0≤j≤d.
4.2.Unimodaltransitionprobabilitiesonthehypercube.CombiningLemma4.8andTheorem4.3andthenprojectingtherandomwalkonthehypercubeonarandomwalkonapath,weobtainthefollowingresult.Theorem4.9.Leti,j∈Vbetwoverticesofad-dimensionalhypercube.Then,Pi,j(t)isunimodal.Proof.Weusethefollowingprojectionofarandomwalkonad-dimensionalhypercubewithloopprobability1/2toarandomwalkonapathwithdvertices,againwithloopprobability1/2.Theinducedrandomwalkisobtainedfromthemappingx7→|x|1,thatis,verticesin{0,1}dwiththesamenumberofonesareequivalent.Itiseasytocheckthatthisnewrandomwalkisarandomwalkonapathwithvertices0,1,...,dthatmovesrightwithprobabilityλk=d−k
2k,leftwithprobabilityµk=d
2k,andloopswithprobability1
2.(ThisprocessisalsoknownastheEhrenfestchain[17]).10
Considernowtherandomwalkonthepathwithvertexset{0,1,...,d}andletjbeanarbitrarynumberwith06j6d.RecallthatP0,jcanbeexpressedasaconvolution(cf.[17])ofPandfasfollows,P0,j=f0,j∗Pj,j.ByCorollary4.4,f0,j(t)islog-concave.Moreover,Lemma4.8impliesthatPj,j(t)isnon-increasingintandhenceunimodal.Astheconvolutionofanylog-concavefunctionwithanyunimodalfunctionisagainunimodal(cf.Fact3.5),itfollowsthatP0,j(t)isunimodalint.Nowﬁxtwoverticesi,jofthed-dimensionalhypercube.Bysymmetry,wemayassumethati=0d≡0.Conditionedontheeventthattheprojectedrandomwalkislocatedatavertexwith|j|1onesatstept,everyvertexwith|j|1onesisequallylikely.ThisgivesP0,j(t)=P0,|j|1(t)/ d|j|1,andthereforetheunimodalityofP0,|j|1(t)impliesdirectlytheunimodalityofP0,j(t),asrequired.
Withmoredirectmethods,onecanprovethefollowingsupplementaryresultgivingfurtherinsightsintothedistributionofPi,j(t).Astheresultisnotrequiredforouranalysis,theproofisgivenintheappendix.Proposition4.10.Leti,j∈Vbetwoverticesofthed-dimensionalhypercubewithdist(i,j)>d/2.ThenPi,j(t)ismonotoneincreasing.4.3.Analysisofthediscretealgorithm.Wearenowreadytoproveourmainresultforhypercubes.
ProofofTheorem4.2.Bysymmetry,itsuﬃcestoboundthedeviationatthevertex0≡0d.Hencebyequation(3.1)wehavetobound
x(t)0−ξ(t)0
6
Pt−1s=0P[i:j]∈Ee(t−s)i,j(P0,i(s)−P0,j(s))
6
Pt−1s=0P[i:j]∈Ee(t−s)i,jP0,i(s)
+
Pt−1s=0P[i:j]∈Ee(t−s)i,jP0,j(s)
6P[i:j]∈E
Pt−1s=0e(t−s)i,jP0,i(s)
+P[i:j]∈E
Pt−1s=0e(t−s)i,jP0,j(s)
.UsingTheorem4.9,weknowthatthesequencesP0,i(s)andP0,j(s)areunimodalinsandhencewecanboundbothsummandsbyLemma3.6(where`=1)toobtainthat
x(t)0−ξ(t)0
62ΛP[i:j]∈Emaxt−1s=0P0,i(s)+2ΛP[i:j]∈Emaxt−1s=0P0,j(s)=2ΛdPi∈Vmaxt−1s=0P0,i(s).(4.5)Toboundthelastterm,weconceptualizetherandomwalkasthefollowingprocess,sim-ilartoaballs-and-binsprocess.Ineachstept∈Nwechooseacoordinatei∈{1,...,d}uniformlyatrandom.Thenwithprobability1/2weﬂipthebitofthiscoordinate;oth-erwiseweleaveitunchanged(equivalently,wesetthebitto1withprobability1/2andtozerootherwise).Nowwepartitiontherandomwalk’sdistributionatsteptaccordingtothenumberofdiﬀerentcoordinateschosen(notnecessarilyﬂipped)uptostept.ConsiderP0,x(t)foravertexx∈{0,1}d.Notethatbythesymmetryofthehypercube,P0,x(t)isthesameforallx∈{0,1}dwiththesame|x|1.Hence,letusﬁxavalue`with06`6dandletusconsiderP0,`(t),whichistheprobabilityforvisitingthevertex,say,1`0d−`from0≡0dinroundt.Since(i)thekchosencoordinatesmustcontainthe`bitswhichshouldbeoneand(ii)allkchosencoordinatesmustbesettothecorrectvalue,wehaveP0,`(t)=Pd k=`Pr[exactlykcoordinateschosenintsteps]·2−k d−`k−` dk.(4.6)11
UsingthistoestimateP0,i(s),wecanboundequation(4.5)by
x(t)0−ξ(t)0
62ΛdPd
`=0 d`max∞
s=0P0,`(s)=2Λd· 1+Pd
`=1 d`max∞
s=0Pd k=`Pr[exactlykcoordinateschoseninssteps]·(d−`k−`)
(d k)·2−k!62Λd· 1+Pd
`=1maxd k=`(d−`k−`)(d`)
(dk)·2−k!=2Λd· 1+Pd
`=1maxd k=`n k`·2−ko!62Λd· 1+Pd
`=1maxd k=`n kdk/2e·2−ko!62Λd· 1+Pd
`=1maxd k=`n(1+o(1))·2k
√
πk
·2−ko!=OΛd1+Pd
`=11
√
`
=O Λd3/2,wherewehaveusedthesimpleconsequenceofStirling’sformulathat kdk
2e6(1+o(1))·2k/√
πk
.Thisprovestheﬁrstclaimofthetheorem.Thesecondclaimfollowsbythefollowingsimpleconstruction.Deﬁnealoadvectorx(0)suchthatx(0)
v:=dforallverticesv=(v1,v2,...,vd)∈{0,1}dwithv1=0,andx(0)
v:=0otherwise.Thenforeachedge{i,j}∈Ewith0=i16=j1thefractionalﬂowatstep1is Pi,jx(0)
i−Pi,jx(0)
j=+1
2.Sinceintheﬁrsttimestepnoroundingerrorshaveyetbeenincurred,eachedgeisallowedtoroundupanddownarbitrarily.Hencewecanletalltheseedgesroundtowardsj,i.e.,Φ(1)
i,j:=1foreachsuchedge{i,j}∈E.Bydeﬁnition,thisimpliesforthecorrespondingroundingerrorthate(1)
i,j=−1
2.Moreover,wehavethefollowingloaddistributionafterstep1.Wehavex(1)
v=0forallverticesvifv1=0,andx(1)
v=dotherwise.Similarly,thefractionalﬂowforeachedge{i,j}∈Ewith0=i16=j1is Pi,jx(0)
i−Pi,jx(0)
j=−1
2.Sincee(1)
i,j=−1
2,
P2
s=1e(s)i,j
willbeminimizedife(2)
i,j=1
2.HencewecansetΦ(2)
i,j:=−1.Thisimpliesthatweendupinexactlythesamesituationasatthebeginning:theloadvectoristhesameandalsothesumoverthepreviousroundingerrorsalongeachedgeiszero.WeconcludethatthereisaninstanceoftheBEDalgorithmforwhichx(t)=x(tmod2),whichprovesthesecondclaimoftheTheorem4.2.
5.Analysisond-dimensionaltorusgraphs.Westartthissectionwiththeformaldeﬁnitionofad-dimensionaltorus.Definition5.1.Ad-dimensionaltorusT(n1,n2,...,nd)withn=n1·n2·...·ndverticeshasvertexsetV={0,1,...,n1−1}×{0,1,...,n2−1}×...×{0,1,...,nd−1}andeveryvertex(i1,i2,...,id)∈Vhas2dneighbors((i1±1)modn1,i2,...,id),(i1,(i2±1)modn2,i3,...,id),...,(i1,i2,...,id−1,(id±1)modnd).Henceforth,wewillalwaysassumethatd=O(1).Wecallatorusuniformifn1=n2=...=nd=d√
n.Withoutlossofgeneralitywewillassumeintheremainderthatn16n26···6nd.Bythesymmetryofthetorus,thisdoesnotrestrictourresults.12
Recallthatλ2denotesthesecondlargesteigenvalueinabsolutevalue.Beforeweanalyzethedeviationbetweentheidealizedanddiscreteprocess,weestimate(1−λ2)−1forgeneraltorusgraphs.Lemma5.2.Forad-dimensionaltorusT=T(n1,n2,...,nd),(1−λ2)−1=Θ n2
d.Proof.Followingthenotationof[3],forak-regulargraphG,letL(G)bethematrixgivenbyLu,u(G)=1,Lu,v(G)=−1
kif{u,v}∈E(G)andLu,v(G)=0otherwise.LetCqbeacyclewithqvertices.Asshownin[3,Example1.4],theeigenvaluesofL(Cq)are1−cos 2πr qwhere06r6q−1.Inparticular,thesecondsmallesteigenvalueofL(Cq)denotedbyτisgivenby1−cos 2π
q.Let×denotetheCartesianproductofgraphs,thatis,foranytwographsG1=(V1,E1),G2=(V2,E2)thegraphG:=G1×G2withG=(V,E)isdeﬁnedbyV=V1×V2andE:= (u1,u2),(v1,u2):u2∈V2∧{u1,v1}∈E1 ∪ (u1,u2),(u1,v2):u1∈V1∧{u2,v2}∈E2 .ItisstraightforwardtogeneralizethisdeﬁnitiontotheCartesianproductofmorethantwographsanditistheneasytocheckthatT(n1,n2,...,nd)=Cn1×Cn2×...×Cnd.ThefollowingtheoremexpressesthesecondsmallesteigenvalueoftheCartesianproductofgraphsintermsofthesecondsmallesteigenvalueoftherespectivegraphs.Theorem5.3([3,Theorem2.12]).LetG1,G2,...,Gdbedgraphsandletτ1,τ2,...,τdbetherespectivesecondsmallesteigenvalueofL(G1),L(G2),...,L(Gd).ThenthesecondsmallesteigenvalueτofL(G1×G2×...×Gd)satisﬁesτ=1
dmind k=1τk.Applyingthistheoremtooursetting,itfollowsthatthesecondsmallesteigenvalueτofL(T)isτ=1
d 1−cos 2π
nd.Asnd>d√
n,wehavecos 2π
nd=1−Θ 1
n2
d.Usingthisandthefactthatdisaconstant,weobtainτ=Θ 1
n2
d.AsTisak-regulargraph,thetransitionmatrixP(T)canbeexpressedasP(T)=I−1
2L(T).ThisimpliesforthesecondsmallesteigenvalueofL(T),τ,andthesecondlargesteigenvalueofthetransitionmatrixP(T),λ2,thatλ2=1−1
2τ.Henceλ2=1−Θ 1
n2
d,whichcompletestheproof.
Notethatthecorrespondingresultsof[11,33]onlyholdforuniformtorusgraphswhilethefollowingresultforouralgorithmholdsforgeneraltorusgraphs.Theorem5.4.Forallinitialloadvectorsonthe(notnecessarilyuniform)d-dimensionaltorusgraphwithnvertices,thedeviationbetweentheidealizedprocessandadiscreteprocesswithaccumulatedroundingerroratmostΛisO(Λ)atalltimes.Foranytorusgraph,weknowthat(1−λ2)−1=Θ(n2
d)byLemma5.2.WithTheorem3.1itfollowsthatanyBEDalgorithm(andinparticularourquasirandomalgorithm)reducesthediscrepancyofanyinitialloadvectorwithdiscrepancyKtoO(1)withinO(n2
dlog(Kn))timesteps(foruniformtorusgraphs,thisnumberoftimestepsisO(n2/dlog(Kn))).ProofofTheorem5.4.Bysymmetryofthetorusgraph,wehavePi,j=P0,i−j.HencewesetPi=P0,i.WewillﬁrstreducetherandomwalkPi,jontheﬁnited-dimensionaltorustoarandomwalkontheinﬁnitegridZd,bothwithloopprobability1/2.Let
Pi,jbethetransitionprobabilityfromitojonZddeﬁnedby
Pi,j=1/(4d)if|i−j|1=1,
Pi,i=1/2,and0otherwise.Fori=(i1,...,id)∈VwesetH(i):=(i1+n1Z,i2+n2Z,...,id+ndZ)⊂Zd.13
With
Pi:=
P0,i,weobservethatPs i=Xk∈H(i)
Ps kforalls>0andi∈V.Weextendthedeﬁnitionofei,jinthenaturalwaybysettingek,`:=ei,jforalli,j∈Vandk∈H(i),`∈H(j).LetARR={±u`|`∈{1,...,d}}∈Zdwithu`beingthe`-thunitvector.Followingequation(3.1)andusingthefactthatbysymmetryitsuﬃcestoboundthedeviationatthevertex0:=0d,wegetx(t)0−ξ(t)0=1
2t−1Xs=0Xi∈VXz∈ARRe(t−s)i,i+z(Ps i−Ps i+z)=1
2t−1Xs=0Xi∈VXz∈ARRe(t−s)i,i+z Xk∈H(i)
Ps k−X`∈H(i+z)
Ps
`!=1
2t−1Xs=0Xz∈ARRXi∈Ve(t−s)i,i+z Xk∈H(i)
Ps k−
Ps k+z!=1
2Xi∈VXz∈ARRXk∈H(i)t−1Xs=0e(t−s)k,k+z 
Ps k−
Ps k+zAsZd=Si∈VH(i)isadisjointunion,wecanalsowritex(t)0−ξ(t)0=1
2Xk∈ZdXz∈ARRt−1Xs=0e(t−s)k,k+z 
Ps k−
Ps k+z.(5.1)Wenowcarefullybreakdownthesumsofequation(5.1)andshowthateachpartcanbeboundedbyO(Λ).Forthisargument,ourmaintoolwillbeLemma3.6.Aswecannotproveunimodalityof 
Ps k−
Ps k+zdirectly,wewilluseappropriatelocalcentrallimittheoremstoapproximatethetransitionprobabilities
Ps kofZdwithamultivariatenormaldistribution.ToderivethelimitingdistributionePs kofourrandomwalk
Pi,j,wefollowLawlerandLimic[26]andletX=(X1,...,Xd)beaZd-valuedrandomvariablewithPr[X=z]=1/(4d)forallz∈ARRandPrX=0d=1/2.ObservethatE[XjXk]=0forj6=ksincebothofthemcannotbenon-zerosimultaneously.Moreover,E[XjXj]=1
4d(−1)2+1
4d(+1)2=1
2dforall16j6d.HencethecovariancematrixisΓ:=hE[XjXk]i16j,k6d=(2d)−1I.FromEq.(2.2)ofLawlerandLimic[26]wegetePs k=1
(2π)dsd/2ZRdexpix·k
√
s
exp−x·Γx
2ddx14
wherei=√
−1
denotestheimaginaryunit.Fromhere,wecanfurtherconcludethatePs k=1
(2π)dsd/2ZRdexpix·k
√
s
−x·Γx
2ddx=1
(2π)dsd/2ZRdexpix·k
√
s
−kxk2Ω2
4dddx=1
(2π)dsd/2ZRdexp−1
4dkxk2Ω2−2i2d
√
s x·kddx.(5.2)Toevaluatetheintegralwecompletethesquare,whichyieldsZRdexp−1
4dkxk2Ω2−2i2d
√
s x·kddx=ZRdexp−1
4dkxk2Ω2−2i2d
√
s x·k−4d2
skkk2Ω2+4d2
skkk2Ω2ddx=exp−d skkk2Ω2ZRdexp −1
4d

x−i2d
√
s k

2Ω2!ddx.(5.3)Bysubstitutingz=x−i2d
√
s kwegetZRdexp −1
4d

x−i2d
√
s k

2Ω2!ddx=ZRdexp−1
4d kzk2Ω2ddz=Z···ZRdexp −1
4d dXi=1z2i!!dzd...dz1=Z···ZRd−1exp −1
4d d−1Xi=1z2i!!ZRexp−1
4dz2ddzddzd−1...dz1=2√
πd
·Z···ZRd−1exp −1
4d d−1Xi=1z2i!!dzd−1...dz1=2√
πd
d.(5.4)Combiningequations(5.2),(5.3),and(5.4),wegetePs k=1
(2π)dsd/2exp−d skkk2Ω22√
πd
d=d
πsd/2exp−dkkk2Ω2
s.(5.5)ItfollowsdirectlyfromClaims4and5ofCooperandSpencer[4]thatforallk∈Zd,z∈ARR,ePs k−ePs k+z=O(kkk−(d+1)2)foralls,(5.6)(s7→ePs k−ePs k+z)hasonlyaconstantnumberoflocalextrema.(5.7)15
Thisgivestheintuitionthatbyapproximating 
Ps k−
Ps k+zwith ePs k−ePs k+z,wecanboundequation(5.1)forsuﬃcientlylargekandsbyLemma3.6.Thisapproximationismadeprecisebythefollowinglocalcentrallimittheorems.Theorem2.3.6ofLawlerandLimic[26]givesforallk∈Zd,z∈ARR,s>0,
 
Ps k−
Ps k+z− ePs k−ePs k+z
=O(s−(d+3)/2).(5.8)Wenowstarttobreakdownequation(5.1).Astheﬁrststep,weconsiderthespecialcasekwithkkk2<3,whichisrelativelystraightforward,sinceitonlyinvolvesaconstantnumberofvertices.WithZd
>3:={k∈Zd:kkk2>3}andZd
<3:={k∈Zd:kkk2<3},
x(t)0−ξ(t)0
61
2Xk∈Zd
<3(5.9a)z
}|
{

Xz∈ARRt−1Xs=0e(t−s)0,0+z
Ps k−
Ps k+z

+1
2

Xk∈Zd
>3Xz∈ARRt−1Xs=0e(t−s)k,k+z
Ps k−
Ps k+z

|
{z
}(5.9b)(5.9)Nowwecanapplythelocalcentrallimittheoremgiveninequation(5.8)to(5.9a)andget(5.9a)=

Xz∈ARRt−1Xs=0e(t−s)k,k+z ePsΩ0−ePsΩ0+z

+

Xz∈ARRt−1Xs=0O(s−(d+3)/2)

=O(Λ),wherethelastequalityfollowsbyLemma3.6combinedwithequation(5.7)andtheproperty
Pt s=1e(s)i,j
6Λ.Inordertoanalyzetheverticesk∈Zd
>3inequation(5.1),weproceedbyﬁxingacutoﬀpointT(k):=Ckkk2
2
ln2(kkk2),k∈Zd
>3,oftheinnermostsumof(5.9b)forsomesuﬃcientlysmallconstantC>0.NotethatthecutoﬀpointischoseninsuchawaythatitisveryunlikelyforarandomwalktoreachavertexwithdistancekwithinlessthanT(k)rounds.Hencethecontributionfromallthosesummands(5.10a)isonlyaconstant.Forboundingtheremainingsum,(5.10b),weusethelocalcentrallimittheorem(5.8)whoseapproximationerrorissmall,sinceweonlyneedtoconsiderroundslargerthanthecutoﬀT(k).Proceedingwiththeformalproof,wehave(5.9b)6(5.10a)z
}|
{

Xk∈Zd
>3Xz∈ARRT(k)Xs=0e(t−s)k,k+z
Ps k−
Ps k+z

+

Xk∈Zd
>3Xz∈ARRt−1Xs=T(k)e(t−s)k,k+z
Ps k−
Ps k+z

|
{z
}(5.10b).(5.10)Notethatthesummandwiths=0iszeroandcanbeignored(sincekkk2>3).Hence16
theﬁrstsummand(5.10a)canbeboundedby(5.10a)=O Xk∈Zd
>3Xz∈ARRT(k)Xs=1
Ps k+
Ps k+z!.(5.11)ItisknownfromLawler[25,Lem.1.5.1(a)]thatforrandomwalksoninﬁnitegrids,Pkkk2>λ√
s
Ps k=O(e−λ)foralls>0andλ>0.Hence,itisalsotruethat
Ps k=O exp −kkk2/√
s
foralls>0,k∈Zd.Withthat,wecannowboundtheterm 
Ps k+
Ps k+zfromequation(5.11).For0<s6T(k),k∈Zd
>3,z∈ARR,andsuﬃcientlysmallC>0,
Ps k+
Ps k+z=Oexp−kkk2
√
s
+exp−kk+zk2
√
s
=Oexp−kkk2
√
3
·√
s
=Oexp−2ln(kkk2)kkk2
√
3·C
kkk2=O kkk−(d+4)2,whereinthesecondlinewehaveusedthatkkk2Ω2−kk+zk2Ω26kkk2Ω2− kkk2Ω2−2kkk2·kzk2+kzk2Ω262kkk2,implyingdirectlythatkk+zk2>q kkk2Ω2−2kkk2
=p kkk2·(kkk2−2)
>kkk2
√
3
,sincekkk2>3.Pluggingthisintoequation(5.11),weobtain(5.10a)=O Xk∈Zd
>3T(k)kkk−(d+4)2!=O Xk∈Zd
>3kkk−(d+2)2ln−2(kkk2)!=O(1).Tobound(5.10b),weapproximatethetransitionprobabilitiesofZdwiththemul-tivariatenormaldistributionofequation(5.5)bythelocalcentrallimittheoremstatedinequation(5.8):(5.10b)=

Xk∈Zd
>3Xz∈ARRt−1Xs=T(k)e(t−s)k,k+zePs k−ePs k+z+Xk∈Zd
>3Xz∈ARRt−1Xs=T(k)e(t−s)k,k+z
Ps k−
Ps k+z−ePs k−ePs k+z

6(5.12a)z
}|
{

Xk∈Zd
>3Xz∈ARRt−1Xs=T(k)e(t−s)k,k+zePs k−ePs k+z

17
+

Xk∈Zd
>3Xz∈ARRt−1Xs=T(k)e(t−s)k,k+zO(s−(d+3)/2)

|
{z
}(5.12b).(5.12)Wecanboundthesecondterm(5.12b),signifyingtheapproximationerrorfromthelocalcentrallimittheorem,by(5.12b)=O dXk∈Zd
>3∞Xs=T(k)s−(d+3)/2!=O Xk∈Zd
>3T(k)−(d+1)/2!=O Xk∈Zd
>3lnd+1(kkk2)
kkkd+12!.AsthereareconstantsC0>0andε>0suchthatlnd+1(kkk2)6C0kkk1−ε2forallk∈Zd
>3,weobtain(5.12b)=O Xk∈Zd
>3kkk−(d+ε)2!.ToseethatthiscanbeboundedbyO(1),observethatwithNd
>3:={k∈Nd:kkk2>3},Xk∈Zd
>3kkk−(d+ε)262dXk∈Nd
>3(k21+···+k2d)−(d+ε)/2.Byconvexityofx7→x2,k21+···+k2d>1
d(k1+···+kd)2,wethenget(5.12b)=O Xk∈Nd
>3(k1+···+kd)−(d+ε)!=O ∞Xx=1Xk∈Ndkkk1=xx−(d+ε)!=O ∞Xx=1xd−1·x−(d+ε)!=O ∞X
x=1x−(1+ε)!=O(1).Finally,tobound(5.12a),weapplyequation(5.7).WealsousethefactthatePs k−ePs k+zcanbeboundedbyO(kkk−(d+1)2)accordingtoequation(5.6).As|Pt s=1e(s)i,j|6Λ,applyingLemma3.6yields(5.12a)=O Xk∈Zd
>3Xz∈ARRΛt−1maxs=T(k)ePs k−ePs k+z!=O Xk∈Zd
>3Xz∈ARRΛkkk−(d+1)2!=O ΛdXk∈Zd
>3kkk−(d+1)2!=O(Λ).18
Combiningalltheabovebounds,wecanconcludethat
x(t)0−ξ(t)0
=O(Λ),meaningthatthedeviationbetweentheidealizedprocessandthediscreteprocessatanytimeandvertexisatmostO(Λ).
6.Lowerboundsforpreviousalgorithms.Forabettercomparisonwithpre-viousalgorithms,thissectiongiveslowerboundsforotherdiscretediﬀusionprocesses.First,weobservethefollowinggenerallowerboundonthediscrepancyfortheRSWalgorithm.Proposition6.1.OnallgraphsGwithmaximumdegree∆,thereisaninitialload-vectorx(0)withdiscrepancy∆diam(G)suchthatfortheRSWalgorithm,x(t)=x(t−1)forallt∈N.Proof.Fixapairofverticesiandjwithdist(i,j)=diam(G).Deﬁneaninitialload-vectorx(0)byx(0)
k:=dist(k,i)·∆.Clearly,theinitialdiscrepancyisx(0)
j−x(0)
i=∆diam(G).Weclaimthatx(1)=x(0).Consideranarbitraryedge{r,s}∈E(G).Then,
Pr,sx(1)
r−Ps,rx(1)
s
=1
2∆
x(0)
r−x(0)
s
61
2∆∆=1
2.Hencetheintegralﬂowonanyedge{r,s}∈E(G)isb1
2c=0andtheload-vectorremainsunchanged.Theclaimfollows.
Intheremainderofthissectionwepresenttwolowerboundsforthedeviationbetweentherandomizedroundingdiﬀusionalgorithmandtheidealizedprocess.First,weproveaboundofΩ(logn)forthehypercube.TogetherwithTheorem4.2thisimpliesthatonhypercubesthequasirandomapproachisasgoodastherandomizedroundingdiﬀusionalgorithm.Theorem6.2.Thereisaninitialloadvectorofthed-dimensionalhypercubewithn=2dverticessuchthatthedeviationoftherandomizedroundingdiﬀusionalgorithmandtheidealizedprocessisatleastlogn/4withprobability1−nΩ(1).Proof.Wedeﬁneaninitialloadvectorx(0)asfollows.Foreveryvertexv=(v1,v2,...,vd)∈{0,1}dwithv1=0wesetx(0)
v=ξ(0)v=0andifv1=1wesetx(0)
v=ξ(0)v=d.Hence,theidealizedprocesswillsendaﬂowofd/(2d)=1/2fromeveryvertexv=(1,v2,v3,...,vd)∈{0,1}dto(0,v2,v3,...,vd).Hencefortheidealizedprocess,ξ(1)v=(1/2)d,thatis,allverticeshavealoadof(1/2)dafteronestepandtheloadisperfectlybalanced.Letusnowconsiderthediscreteprocess.LetV0bethesetofverticeswhosebitstringbeginswith0.Consideranynodev∈V0.Notethatallneighborsofvhavealoadof1andtheintegralﬂowfromanyofthoseneighborsequals1withindependentprobability1/2.Hencetheloadofvinthenextstepisjustabinomialrandomvariable,andusingthefactthat r s>(r/s)s,weobtainPrx(1)
v=3
4d>Prx(1)
v>3
4d>d(3/4)d2−d>4
3(3/4)d2−d>n−1+C,forsomeconstantC>0sinced=log2n.AsthemaximumdegreeofthegraphislognandthesizeofV0isn/2,itfollowsthatthereisasubsetS⊆V0ofsizeΩ n log4nin19
thehypercubesuchthateverypairinShasdistanceatleast4.Byconstruction,therespectiveeventsx(1)
v>(3/4)dareindependentforallverticesv∈S.HencePr∃v∈S:x(1)
v>3
4d>1− 1−n−1+CΩn log4n>1−n−C0,where0<C0<Cisanotherconstant.Thismeansthatwithprobabilityatleast1−n−C0theloadatvertexuatstep1willbeatleast(3/4)dinthediscreteprocess,butequals(1/2)dintheidealizedprocess.Thiscompletestheproof.
Itremainstogivealowerboundforthedeviationbetweentherandomizedroundingandtheidealizedprocessfortorusgraphs.Thefollowingtheoremprovesapolylogarith-miclowerboundfortherandomizedroundingalgorithm,whichshouldbecomparedtotheconstantupperboundforthequasirandomapproachofTheorem5.4.Similarresultscanalsobederivedfornon-uniformtorusgraphs.Theorem6.3.Thereisaninitialloadvectorofthed-dimensionaluniformtorusgraphwithnverticessuchthatthedeviationbetweentherandomizedroundingdiﬀusionalgorithmandtheidealizedprocessisΩ(polylog(n))withprobability1−o(1).Proof.LetnbeasuﬃcientlylargeintegerandTbead-dimensionaltorusgraphwithnverticesandsidelengthd√
n∈N.LetBk(u):={v∈V:kv−uk∞6k}and∂Bk(u):={v∈V:kv−uk∞=k}.Foreveryvertexv∈V(T),wedeﬁne|B`/2(v)|=`d=(logn)1/4with`:=(logn)1/(4d),whereweassumew.l.o.g.that`isanoddinteger.For`0:=(logn)2/(3d),deﬁneasetS⊆VbyS:=(x1`0,x2`0,...,xd`0)
16x1,x2,...,xd<d√
n/`0−1 ,thatis,everypairofdistinctverticesinShasacoordinate-wisedistancewhichisamultipleof`0.Notethat|S|=Ω(n/`0d).Deﬁnetheinitialloadvectorasx(0)
i=ξ(0)i:=2d·max{0,`/2−dist(i,S)},i∈V.Clearly,theinitialdiscrepancyisK=2d·`/2.TheideaisnowtodecomposeTintosmallersubgraphscenteredarounds∈S,sincetheupperboundontheconvergencerategivenbyTheorem3.1hasastrongdependenceonthesizeofthegraph.Thenwerelatethesimultaneousconvergenceoneachofthesmallergraphstotheconvergenceontheoriginalgraph.AnillustrationofourdecompositionofTcanbefoundinFigure6.1.Fixsomes∈S.ThenthesubgraphinducedbytheverticesB`/2(s)isad-dimensionalgridwithexactlyn0:=(logn)1/4vertices.LetT0=T0(s)denotethecorrespondingd-dimensionaltorusgraphwiththesamevertices,butadditionalwrap-aroundedgesbetweenverticesof∂B`/2(s).W.l.o.g.weassumethatthesidelengthd√
nofTisamultipleofthesidelength`ofT0(s).LetP0bethediﬀusionmatrixofT0(s).Letusdenotebyξ0(0)(x0(0))theprojectionoftheloadvectorξ(0)(x(0))fromTontoT0(s).ByCorollary3.2,theidealizedprocessreducesthediscrepancyonT0(s)fromK=(logn)1/(4d)/2to1withint0:=O((n0)2/dlog(Kn0))=O(loglog(n)(logn)1/(2d))timesteps.WenowwanttoarguethatthisalsohappensontheoriginalgraphTwithnvertices.NotethattheconvergenceoftheidealizedprocessonT0(s)implieskξ0(t0)−
ξ0k∞=kP0t0ξ0(0)−
ξ0k∞61.(6.1)Furthermore,notethattheaverageload
ξ0ineachT0(s)satisﬁes
ξ062d·`/4.Ournextobservationisthatforanytwoverticesu,v∈T0(s),Pt0u,v6P0t0u,v(6.2)20
``0−`
Fig.6.1.OverviewofthedecompositionofTintovariousT0(s)forthetwo-dimensionalcased=2.TheinnerrectanglesrepresentthevarioussmallergridsT0(s)withs∈S.Thedarknessindicatestheamountoftheinitialload.NotethattheinitialloadofverticesoutsidetheT0(s)’sis0.asarandomwalkonT0(s)canbeexpressedasaprojectionofarandomwalkonT(byassigningeachvertexinT0(s)toasetofverticesinT).Withtheobservations•forv∈T0(s):ξ(0)v=ξ0v(0),•forv∈Tand`/26dist(v,s)6t0:ξ(0)v=0(ast0=o(`0−`/2)),•forv∈Tanddist(v,s)>t0:Pt0v,s=0,weobtainforanyvertexv∈B`/2(s),ξ(t0)s= P(t0)·ξ(0)s=Xv∈Tξ(0)vP(t0)v,s=Xv∈T0(s)ξ0v(0)P(t0)v,s.Byﬁrstapplyingequation(6.2)andthenequation(6.1),wegetξ(t0)s6Xv∈T0(s)ξ0s(0)P0(t0)v,s=ξ0s(t0)6
ξ0+1.Thismeansthataftert0timesteps,theidealizedprocessachievesagoodbalancingats.Ontheotherhand,thediscreteprocessmayfailwithint0timestepsifthereisanssuchthatalledgesinT0(s)roundtowardssatalltimestepst6t0.(Notethatbyconstruction,noloadfromanotherT0(s0),s0∈S\{s},canreachT0(s)withint0steps,sincethedistancebetweenanyvertexinT0(s)andT0(s0)is`0−2`>t0.)Moreover,bydeﬁnitionofx(0),|x(0)
u−x(0)
v|∈{0,2d}if{u,v}∈E(T).Hencethefractionalﬂowintheﬁrststepis∈{0,1
2}andforﬁxedstheprobabilitythatx(0)
u=x(1)
uforallu∈T0(s)isatleast2−|B`/2(s)|.Byinduction,forﬁxedstheprobabilitythatx(0)
u=x(t0)uholdsforallu∈T0(s)isatleast2−|B`/2(s)|t0=2−(logn)1/4·O(loglog(n)(logn)1/(2d))>2−(logn)4/5.Aswehave|S|=Ω(n/`0d)=Ω(poly(n))independentevents,itfollowsthatthereisatleastones∈Swithx(t0)s=x(0)
s=`/2·2dwithprobability1−1−2−(logn)4/5Ω(poly(n))>1−n−C,21
whereC>0issomeconstant.Ifthishappens,thenthedeviationbetweenthediscreteandidealizedprocessatvertexs∈Satstept0isatleast
x(t0)s−ξ(t0)s
>
2d·`/2−(2d·`/4+1)
=Ω((logn)1/(4d)),andtheclaimfollows.
7.Conclusions.Weproposeandanalyzeanewdeterministicalgorithmforbal-ancingindivisibletokens.Byachievingaconstantdiscrepancyinoptimaltimeonalltorusgraphs,ouralgorithmimprovesuponallpreviousdeterministicandrandomap-proacheswithrespecttobothrunningtimeanddiscrepancy.ForhypercubesweproveadiscrepancyofΘ(logn),whichisalsosigniﬁcantlybetterthanthe(deterministic)RSWalgorithm,whichachievesadiscrepancyofΩ(log2n).Onaconcretelevel,itwouldbeinterestingtoextendtheseresultstoothernetworktopologies.Fromahigher-levelperspective,ournewalgorithmprovidesastrikingexam-pleofquasirandomnessinalgorithmics.Devisingandanalyzingsimilaralgorithmsforothertaskssuchasrouting,schedulingorsynchronizationremainsaninterestingopenproblem.References.[1]W.Aiello,B.Awerbuch,B.M.Maggs,andS.Rao.Approximateloadbalancingondynamicandasynchronousnetworks.In25thAnnualACMSymposiumonTheoryofComputing(STOC’93),pages632–641,1993.[2]R.D.Barve,E.F.Grove,andJ.S.Vitter.Simplerandomizedmergesortonparalleldisks.ParallelComputing,23(4-5):601–631,1997.[3]F.Chung.SpectralGraphTheory.AMS,revisededition,2006.URLhttp://www.math.ucsd.edu/∼fan/research/revised.html.[4]J.CooperandJ.Spencer.Simulatingarandomwalkwithconstanterror.Combi-natorics,Probability&Computing,15:815–822,2006.[5]G.Cybenko.Loadbalancingfordistributedmemorymultiprocessors.JournalofParallelandDistributedComputing,7:279–301,1989.[6]P.Diaconis,R.Graham,andJ.Morrison.Asymptoticanalysisofarandomwalkonahypercubewithmanydimensions.RandomStructuresandAlgorithms,1(1):51–72,1990.[7]B.Doerr,T.Friedrich,andT.Sauerwald.Quasirandomrumorspreading.In19thAnnualACM-SIAMSymposiumonDiscreteAlgorithms(SODA’08),pages773–781,2008.[8]R.Els¨asserandT.Sauerwald.Discreteloadbalancingis(almost)aseasyascon-tinuousloadbalancing.In29thAnnualACMPrinciplesofDistributedComputing(PODC’10),pages346–354,2010.[9]R.Els¨asser,B.Monien,andS.Schamberger.Distributingunitsizeworkloadpack-agesinheterogenousnetworks.JournalofGraphAlgorithms&Applications,10(1):51–68,2006.[10]J.A.Fill.Thepassagetimedistributionforabirth-and-deathchain:Strongsta-tionarydualitygivesaﬁrststochasticproof.JournalofTheoreticalProbability,22(3):543–557,2009.[11]T.FriedrichandT.Sauerwald.Near-perfectloadbalancingbyrandomizedrounding.In41stAnnualACMSymposiumonTheoryofComputing(STOC’09),pages121–130,2009.[12]T.Friedrich,M.Gairing,andT.Sauerwald.Quasirandomloadbalancing.In21stAnnualACM-SIAMSymposiumonDiscreteAlgorithms(SODA’10),pages1620–1629.SIAM,2010.22
[13]T.Friedrich,M.Gairing,andT.Sauerwald.Quasirandomloadbalancing.SIAMJournalonComputing,41(4):747–771,2012.[14]J.Gehrke,C.Plaxton,andR.Rajaraman.Rapidconvergenceofalocalloadbal-ancingalgorithmforasynchronousrings.TheoreticalComputerScience,220(1):247–265,1999.[15]B.GhoshandS.Muthukrishnan.Dynamicloadbalancingbyrandommatchings.JournalofComputerandSystemSciences,53(3):357–370,1996.[16]B.Ghosh,F.T.Leighton,B.M.Maggs,S.Muthukrishnan,C.G.Plaxton,R.Ra-jaraman,A.W.Richa,R.E.Tarjan,andD.Zuckerman.Tightanalysesoftwolocalloadbalancingalgorithms.SIAMJournalonComputing,29(1):29–64,1999.[17]G.GrimmetandD.Stirzaker.ProbabilityandRandomProcesses.OxfordUniversityPress

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

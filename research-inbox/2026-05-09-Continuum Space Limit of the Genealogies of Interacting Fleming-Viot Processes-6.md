---
source: "https://arxiv.org/abs/1508.07169v2"
title: "Continuum Space Limit of the Genealogies of Interacting Fleming-Viot Processes on $\\Z$"
author: "Andreas Greven, Rongfeng Sun, Anita Winter"
year: "2015"
publication: "arXiv preprint / math.PR"
download: "https://arxiv.org/pdf/1508.07169v2"
pdf: "https://arxiv.org/pdf/1508.07169v2"
captured_at: "2026-05-09T13:10:22Z"
updated_at: "2026-05-09T13:10:22Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche genealogy of morals"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Continuum Space Limit of the Genealogies of Interacting Fleming-Viot Processes on $\Z$

- 著者: Andreas Greven, Rongfeng Sun, Anita Winter
- 年: 2015
- 掲載情報: arXiv preprint / math.PR
- 情報源: [arxiv](https://arxiv.org/abs/1508.07169v2)
- ダウンロード: https://arxiv.org/pdf/1508.07169v2
- PDF: https://arxiv.org/pdf/1508.07169v2

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

We study the evolution of genealogies of a population of individuals, whose type frequencies result in an interacting Fleming-Viot process on $\Z$. We construct and analyze the genealogical structure of the population in this genealogy-valued Fleming-Viot process as a marked metric measure space, with each individual carrying its spatial location as a mark. We then show that its time evolution converges to that of the genealogy of a continuum-sites stepping stone model on $\R$, if space and time are scaled diffusively. We construct the genealogies of the continuum-sites stepping stone model as functionals of the Brownian web, and furthermore, we show that its evolution solves a martingale problem. The generator for the continuum-sites stepping stone model has a singular feature: at each time, the resampling of genealogies only affects a set of individuals of measure $0$. Along the way, we prove some negative correlation inequalities for coalescing Brownian motions, as well as extend the theory of marked metric measure spaces (developed recently by Depperschmidt, Greven and Pfaffelhuber [DGP12]) from the case of probability measures to measures that are finite on bounded sets.

## PDF Text

arXiv:1508.07169v2 [math.PR] 6 Jan 2017
CONTINUUMSPACELIMITOFTHEGENEALOGIESOFINTERACTINGFLEMING-VIOTPROCESSESONZANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3Abstract.Westudytheevolutionofgenealogiesofapopulationofindividuals,whosetypefrequenciesresultinaninteractingFleming-ViotprocessonZ.Weconstructandanalyzethegenealogicalstructureofthepopulationinthisgenealogy-valuedFleming-Viotprocessasamarkedmetricmeasurespace,witheachindividualcarryingitsspatiallocationasamark.Wethenshowthatitstimeevolutionconvergestothatofthegenealogyofacontinuum-sitessteppingstonemodelonR,ifspaceandtimearescaleddiﬀusively.Weconstructthegenealogiesofthecontinuum-sitessteppingstonemodelasfunctionalsoftheBrownianweb,andfurthermore,weshowthatitsevolutionsolvesamartingaleproblem.Thegeneratorforthecontinuum-sitessteppingstonemodelhasasingularfeature:ateachtime,theresamplingofgenealogiesonlyaﬀectsasetofindividualsofmeasure0.Alongtheway,weprovesomenegativecorrelationinequalitiesforcoalescingBrownianmotions,aswellasextendthetheoryofmarkedmetricmeasurespaces(developedrecentlybyDepperschmidt,GrevenandPfaﬀelhuber[DGP12])fromthecaseofprobabilitymeasurestomeasuresthatareﬁniteonboundedsets.AMS2010subjectclassiﬁcation:60K35,60J65,60J70,92D25.Keywords:Brownianweb,continuum-sitessteppingstonemodel,evolvinggenealogies,interactingFleming-Viotprocess,markedmetricmeasurespace,martingaleproblems,negativecorrelationinequalities,spatialcontinuumlimit.Contents1.Introductionandmainresults21.1.BackgroundandOverview21.2.MarkedMetricMeasureSpaces51.3.InteractingFleming-Viot(IFV)GenealogyProcesses81.4.GenealogiesofContinuum-sitesSteppingStoneModel(CSSM)onR131.5.ConvergenceofRescaledIFVGenealogies171.6.MartingaleProblemforCSSMGenealogyProcesses181.7.Outline202.ProofofTheorems1.12and1.17213.ProofsofthepropertiesofCSSMGenealogyProcesses264.ProofofconvergenceofRescaledIFVGenealogies294.1.ConvergenceofFinite-dimensionalDistributions304.2.Tightness324.3.ProofofconvergenceofrescaledIFVprocesses(Theorem1.32)405.MartingaleProblemforCSSMGenealogyProcesses405.1.GeneratorActiononRegularTestFunctions415.2.GeneratorActiononGeneralTestFunctions505.3.ProofofTheorem1.4054AppendixA.Proofsonmarkedmetricmeasurespaces55AppendixB.TheBrownianweb56AppendixC.CorrelationInequalitiesforCoalescingBrownianmotions57References62
Date:October13,2018.1Universit¨atErlangen-N¨urnberg,Germany.Email:greven@mi.uni-erlangen.de2NationalUniversityofSingapore,Singapore.Email:matsr@nus.edu.sg3Universit¨atDuisburg-Essen,Germany.Email:anita.winter@uni-due.de1
2ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER31.IntroductionandmainresultsInthestudyofspatialpopulationmodelsondiscretegeographicspaces(forexampleZd),suchasbranchingprocesses,votermodels,orinteractingFisher-Wrightdiﬀusions(Fleming-Viotmodels),thetechniqueofpassingtothespatialcontinuumlimithasproventobeusefulingaininginsightintothequalitativebehaviouroftheseprocesses.AkeyexampleisbranchingrandomwalksonZd,leadingtotheDawson-Watanabeprocess[Daw77]onRdandFisher-Wrightdiﬀusions;catalyticbranchingandmutuallycatalyticbranchingonZ,leadingtoSPDEonR[KS88,EF96,DP98,DEF+02b,DEF+02a].Thegoalofthispaperistocarryoutthisprogramatthelevelofgenealogies,ratherthanjusttypeormassconﬁgurations.WefocushereoninteractingFleming-ViotmodelsonZ.1.1.BackgroundandOverview.Wesummarizebelowthemainresultsofthispaper,recallsomehistoricalbackground,aswellasstatesomeopenproblems.Summaryofresults.Thepurposeofthispaperistwofold.Ontheonehand,wewanttounderstandtheformationoflargelocalone-familyclustersinFleming-ViotpopulationsonthegeographicspaceZ1,bytakingaspace-timecontinuumlimitofthegenealogicalconﬁgurationsequippedwithtypes.Ontheotherhandweusethisexampletodevelopthetheoryoftree-valueddynamicsviamartingaleproblemsinsomenewdirections.Inparticular,thisistheﬁrststudyofatree-valueddynamicsonanunboundedgeograph-icalspacewithinﬁnitesamplingmeasure,whichrequiresustoextendboththenotionofmarkedmetricmeasurespacesin[GPW09,DGP11]andthemartingaleproblemfor-mulationsin[GPW13,DGP12]tomarkedmetricmeasurespaceswithinﬁnitesamplingmeasuresthatareboundedlyﬁnite(i.e.,ﬁniteonboundedsets).Hereisasummaryofourmainresults:(1)Weextendthetheoryofmarkedmetricmeasurespaces[GPW09,DGP12]fromprobabilitysamplingmeasurestoinﬁnitesamplingmeasuresthatareboundedlyﬁnite,whichserveasthestatespaceofmarkedgenealogiesofspatialpopulationmodels.SeeSection1.2.(2)WecharacterizetheevolutionofthegenealogiesofinteractingFleming-Viot(IFV)modelsbywell-posedmartingaleproblemsonspacesofmarkedultrametricmeasurespaces.SeeSection1.3.(3)WegiveagraphicalconstructionofthespatialcontinuumlimitoftheIFVgenealogyprocess,whichisthegenealogyprocessoftheso-calledContinuumSitesStepping-stoneModel(CSSM),takingvaluesinthespaceofultrametricmeasurespaceswithspatialmarksandaninﬁnitetotalpopulation.Thegraphicalconstructionisbasedonthe(dual)Brownianweb[FINR04].TheCSSMgenealogyprocesshasthefeaturethat,assoonast>0,theprocessentersaregularsubsetofthestatespacethatisnotclosedunderthetopology.Onlyonthissubsetwecanevaluatetheactionoftheoperatorofthemartingaleprobleminitsactionontestfunctions.Theniceaspectisthatthesetofthesestatesarepreservedunderthedynamic.
Thisleadstoasingularstructurewithcomplicationsfortheassociatedmartingaleproblemandforthestudyofcontinuityoftheprocessattime0.SeeSection1.4.(4)Weprovethatundersuitablescaling,theIFVgenealogyprocessesconvergetotheCSSMgenealogyprocess.Theproofisbasedondualitywithspatialcoales-cents,togetherwithanovelapproachofcontrollingthegenealogystructureusingaweakerconvergenceresultonthecorrespondingmeasure-valuedprocesses,withmeasuresonthegeographicandtypespace(withnogenealogies).SeeSection1.5.(5)WeshowthattheCSSMgenealogyprocesssolvesamartingaleproblemwithasingulargenerator.Moreprecisely,thegeneratoractioninvolvesindividuals,which
CONTINUUMSPACELIMITOFIFVGENEALOGIES3
arenottypicalunderthesamplingmeasure,sothatthedynamicisdrivenbyatypicalindividualsatatypicallocations.Inparticular,thegeneratorisonlydeﬁnedonaregularsubsetofthestatespace.SeeSection1.6.(6)WeprovesomenegativecorrelationinequalitiesforcoalescingBrownianmotions,whichareofindependentinterest.SeeAppendixC.
Besidesthedescriptionofthegenealogiesofthecurrentpopulation,wealsopreparethegroundforthetreatmentofallindividualseveralive,i.e.fossils,movingfromthestatespaceofmarkedultrametricmeasurespacestothestatespaceofmarkedmeasureR-trees,whichwillbecarriedoutelsewhere.
Historyoftheproblem:Whyareweparticularlyinterestedinone-dimensionalge-ographicspacesforourscalingresults?Manyinteractingspatialsystemsthatmodelevolvingpopulations,i.e.,MarkovprocesseswithstatespacesIG(I=R,N,[0,1],etc.,andG=ZdorthehierarchicalgroupΩN)thatevolvebyamigrationmechanismbe-tweensitesandastochasticmechanismactinglocallyateachsite,exhibitadichotomyintheirlongtimebehavior.Forexample,whenG=Zdandthemigrationisinducedbythesimplesymmetricrandomwalk:indimensiond≤2,oneobservesconvergencetolawsconcentratedonthetrapsofthedynamic;whileind≥3,nontrivialequilibriumstatesareapproachedandtheextremalinvariantmeasuresarespatiallyhomogeneousergodicmeasurescharacterizedbytheintensityoftheconﬁguration.Typicalexamplesincludethevotermodel,branchingrandomwalks,spatialMoranmodels,orsystemsofinteract-ingdiﬀusions(e.g.,Feller,Fisher-WrightorAndersondiﬀusions).Oneobtainsuniversaldimension-dependentscalinglimitsforthesemodelsifanadditionalcontinuumspatiallimitistaken,resultingin,forexample,superBrownianmotion(seeLiggett[Lig85]orDawson[Daw93]).Inthelow-dimensionregime,thecasesd=1andd=2areverydiﬀerent.Ind=2,oneobservesforexampleinthevotermodeltheformationofmono-typeclustersonspatialscalestα/2witharandomα∈[0,1],aphenomenoncalleddiﬀusiveclustering(seeCoxandGriﬀeath[CG86]).Intheone-dimensionalvotermodel,theclustershaveanextensionofaﬁxedorderofmagnitudebutexhibitrandomfactorsinthatscale.Moreprecisely,inspace-timescales(√
t,t)fort→∞,wegetannihilatingBrownianmotions.Similarresultshavebeenobtainedforlow-dimensionalbranchingsystems(Klenke[Kle00],Winter[Win02]),systemsofinteractingFisher-Wrightdiﬀusions(FleischmannandGreven[FG94],[FG96]andsubsequently[Zho03],[DEF+00])andfortheMoranmodelind=2(Greven,Limic,Winter[GLW05]).Inallthesemodels,onecangofurtherandstudythecompletespace-timegenealogystructureoftheclusterformationanddescribethisphenomenonasymptoticallybythespatialcontinuumlimit.Inparticular,thedescriptionfortheone-dimensionalvotermodelcanbeextendedtothecompletespace-timegenealogystructure,obtainingasscalinglimittheBrownianweb[Arr79,Arr81,TW98,FINR04](seeAppendixB,andtherecentsurvey[SSS15]).Moreprecisely,theBrownianwebisdeﬁnedbyconsideringinstantaneouslycoalescingone-dimensionalBrownianmotionsstartingfromeveryspace-timepointinR×R.Itarisesasthediﬀusivescalinglimitofcontinuoustimecoalescingsimplesymmetricrandomwalksstartingfromeveryspace-timepointinZ×R,whichrepresentthespace-timegenealogiesofthevotermodel.Thisisanalogoustothestudyofhistoricalprocessforbranchingprocesses,whichapproximatestheancestralpathsofbranchingrandomwalksbythatofsuperBrownianmotion(seee.g.[DP91,FG96,GLW05]).Thebasicideabehindallthisisthat,wecanidentifythegenealogicalrelationshipbetweentheindividualsofthepopulationlivingatdiﬀerenttimesanddiﬀerentlocations.Thisraisesthequestionofwhetheronecanobtainadescriptionoftheasymptoticbehavior
4ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3ofthecompletegenealogicalstructureoftheprocessonlargespace-timescales,whichwillinturnallowforasymptoticdescriptionsofinterestinggenealogicalstatisticsthatarenotexpressibleinanaturalwayintermsoftheconﬁgurationprocess.TheseobservationsonthegenealogicalstructuregoesbacktothegraphicalconstructionofthevotermodelduetoHarris,andcontinuesuptothehistoricalprocessofDawsonandPerkinsforbranchingmodels[DP91],orrepresentationbycontourprocesses[GJ98,Ald93].Tobetterdescribegenealogies,thenotionofR-trees,markedR-treesormarkedmeasureR-treesweredevelopedasaframework[Eva97,EPW06].Theseobjectscontaintherelevantinformationabstractedfromthelabeledgenealogytree,whereeveryindividualiscodedwithitslifespananditslocationsateachtime.Suchacodingmeansinparticularthatallmembersofthepopulationaredistinguished,whichinformationismostlynotneeded.Inthelargepopulationlimit,itsuﬃcestoconsiderthestatisticsofthepopulationviasampling.Forthispurpose,oneequipsthepopulationwithametric(genealogicaldistance),aprobabilitymeasure(theso-calledsamplingmeasure,whichallowstodrawtypicalﬁnitesamplesfromthepopulation)togetherwithamark(specifyingtypesandlocations).Thisdescriptionintermsofrandommarkedmetricmeasurespaces(infact,themetricdeﬁnesatree)isthemostnaturalframeworktodiscusstheasymptoticanalysisofpopulationmodels,sinceitcomprisesexactlytheinformationonewantstokeepfortheanalysisinthelimitsofpopulationswithevenlocallyinﬁnitelymanyindividuals.Theevolutionofaprocesswithsuchastatespaceisdescribedbymartingaleproblems[GPW13,DGP12].Openproblemsandconjectures.Weshowinthispaperthatthespatialcontinuumlimitoftheone-dimensionalinteractingFleming-Viotgenealogyprocesssolvesamartin-galeproblem.However,duetothesingularnatureofthegenerator,theuniquenessofthismartingaleproblemisnon-standard,andweleaveitforafuturepaper.Inparticular,itisdiﬃculttoestablishthedualityrelationatthelevelofgenerators,becausethegeneratorsfortheprocessanditsdualareonlydeﬁnedforspecialtestfunctionsatspecialpoints.InsteadofZ,resp.R,asgeographicspaces,onecouldconsiderthehierarchicalgroupΩN=⊕NZN,withZNbeingthecyclicalgroupoforderN,respectivelythecontinuumhierarchicalgroupΩ∞
N=LZZN.BrownianmotiononRcanbereplacedbysuitableL´evyprocessesonΩ∞
Nandtheprogramofthispapercanthenbecarriedout.TheBrownianwebwouldhavetobereplacedbyanobjectbasedonL´evyprocessesasstudiedin[EF96],[DEF+00].Weconjecturethattheanaloguesofourtheoremswouldholdinthiscontext.AfurtherchallengewouldbetogiveauniﬁedtreatmentoftheseresultsonR,Ω∞
N.AnotherdirectionistoconsiderthegenealogyprocessesofinteractingFellerdiﬀusions,catalyticormutuallycatalyticdiﬀusions,interactinglogisticFellerdiﬀusions,andderivetheirgenealogicalcontinuumlimits.Amorediﬃcultextensionwouldbetoincludean-cestralpathsasmarks,whichraisesnewchallengesrelatedtotopologicalpropertiesofthestatespace.Namely,asthespaceofpathsisaPolishspace,itisnotaHeine-Borelspaceasclosedballsaroundapatharenotcompact.
OutlineofSection1.Theremainderoftheintroductionisorganizedasfollows.InSubsection1.2,werecallandextendthenotionofmarkedmetricmeasurespacesneededtodescribethegenealogies.InSubsection1.3,wedeﬁnetheinteractingFleming-Viot(IFV)genealogyprocessviaamartingaleproblemandgiveadualrepresentationintermsofaspatialcoalescent.InSubsection1.4,wegive,basedontheBrownianweb,agraphicalconstructionforthecontinuum-sitessteppingstonemodel(CSSM)onRanditsmarkedgenealogyprocess,whichisthecontinuumanalogueandscalinglimitoftheinteract-ingFleming-ViotgenealogyprocessonZ,underdiﬀusivescalingofspaceandtimeand
CONTINUUMSPACELIMITOFIFVGENEALOGIES5normalizingofmeasure,afactwhichwestateinSubsection1.5.InSubsection1.6,weformulateamartingaleproblemfortheCSSMgenealogyprocess.InSubsection1.7weoutlinetherestofthepaper.Ω1.2.MarkedMetricMeasureSpaces.InthissubsectionweintroducethestatespaceofthegenealogiesofinteractingFleming-Viotprocesses.Wewanttodescribeevolvinggenealogiesofthewholepopulationofallindividualscurrentlyaliveallowingtosamplefromthispopulation.Wealsowanttoincludetheindividuals’positionsingeographicspaceandpossiblegenetictypes.Wethereforeregardgenealogiesas(equivalenceclasses)ofmarkedmetricmeasurespaces.Asourgeographicspaceisinﬁnite(andnotcompact),itwon’tbepossibletosampleindividualsbymeansofaﬁnite(orafterrenormalizingofaprobability)measure.Weratherrequirethesamplingmeasuretobeﬁniteonallsubpopulationswhichcanbeobtainedbyrestrictingtoﬁnitegeographicsubspaces.Thefollowingdeﬁnitionofmarkedmetricmeasurespacesextendstheoneintroducedin[DGP11]whichconsideredprobabilitymeasuresonly.
Deﬁnition1.1(V-mmm-spaces).Let(V,rV)beacompleteseparablemetricspacewithmetricrV,andletobeadistinguishedpointinV.1.Wecall(X,r,µ)aV-markedmetricmeasurespace(V-mmmspaceforshort)if:(i)(X,r)isacompleteseparablemetricspace,(ii)µisameasureontheBorelσ-algebraofX×V,withµ(X×Bo(R))<∞foreachballBo(R)⊂VofﬁniteradiusRcenteredato.2.WesaytwoV-mmmspaces(X,rX,µX)and(Y,rY,µY)areequivalentifthereexistsameasureablemapϕ:X→Y,suchthat(1.1)rX(x1,x2)=rY(ϕ(x1),ϕ(x2))forallx1,x2∈supp(µX(·×V)),andifeϕ:X×V→Y×Visdeﬁnedbyeϕ(x,v):=(ϕ(x),v),then(1.2)µY=µX◦eϕ−1.Inotherwords,ϕisanisometrybetweensupp(µX(·×V))andsupp(µY(·×V)),andtheinducedmapeϕismarkandmeasurepreserving.Wedenotetheequivalenceclassof(X,r,µ)by(1.3)
(X,r,µ).3.Thespaceof(equivalenceclassesof)V-mmmspacesisdenotedby(1.4)MV:=
(X,r,µ):(X,r,µ)isaV-mmmspace .4.Thesubspaceof(equivalenceclassesof)V-mmmspaceswhichadmitamarkfunc-tionisdenotedby(1.5)MV
fct:={
(X,r,µ)∈MV:∃measurableκ:X→Vs.t.µ(d(x,·))=µ(dx×V)⊗δκ(x)}.NotethatMVdependsbothonthesetVandthemetricrVsincethelatterdeﬁnesthesetsonwhichthemeasuremustbeﬁnite.Markedmetricmeasurespaceswereintroducedin[DGP11],whichextendsthenotionofmetricmeasurespacesstudiedearlierin[GPW09].Deﬁnition1.1isexactlytheanalogueof[DGP11,Def.2.1],whereµisaprobabilitymeasure.Thebasicinterpretationinourcontextisthat:Xisthespaceofindividuals;r(x1,x2)measuresthegenealogicaldistancebetweentwoindividualsx1andx2inX;themeasureµisameasureontheindividualsandthemarkstheycarry(whichcanbespatiallocationaswellastype,orevenancestral
6ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3pathsuptonow),allowingustodrawsamplesfromindividualswithmarksinaboundedset.TodeﬁneatopologyonMVthatmakesitaPolishspace,wewillmakeuseofthemarkedGromov-weaktopologyintroducedin[DGP11,Def.2.4]forV-mmmspacesequippedwithprobabilitymeasures.Inthistopologyasequenceconvergesiﬀwecanembeddtheinvolvedmetricspacesisometricallyintoonemetricspacesuchthattheimagesofthe“sampling”measuresconvergesweakly.
ThebasicideatoextendthistoMVisthat,givenouras-sumptiononµinDeﬁnition1.1.1.(ii),wecanlocalizeµtoﬁniteballsinVtoreduceµtoaﬁnitemeasure.WecanthenapplythemarkedGromov-weaktopology(whichalsoappliestoﬁnitemeasuresinsteadofprobabilitymeasures)torequireconvergenceforeachsuchlocalizedversionoftheV-mmmspaces.WewillcallsuchatopologyV-markedGromov-weak#topology,replacingweakbyweak#,followingtheterminologyin[DVJ03,SectionA2.6]fortheconvergenceofmeasuresthatareboundedonboundedsubsetsofacompleteseparablemetricspace.Notethatvagueconvergenceisformeasuresthatareﬁniteoncompactratherthanboundedsubsets.BothnotionsagreeonHeine-Borelspaces(compare,[ALW14]).
Deﬁnition1.2(V-markedGromov-weak#Topology).Fixasequenceofcontinuousfunc-tionsψk:V→[0,1],k∈N,suchthatψk=1onBo(k),theballofradiuskcenteredato∈V,andψk=0onBco(k+1).Letχ:=
(X,r,µ)andχn:=
(Xn,rn,µn),n∈N,beelementsofMV.Letψk·µbethemeasureonX×Vdeﬁnedby(ψk·µ)(d(x,v)):=ψk(v)µ(d(x,v)),andletψk·µnbedeﬁnedsimilarly.Wesaythatχn−→n→∞χintheV-markedGromov-weak#topologyifandonlyif:Ω(1.6)
(Xn,rn,ψk·µn)=⇒n→∞
(X,r,ψk·µ)intheGromov-weaktopologyforeachk∈N.WhenV=Rd,wemaychooseψktobeinﬁnitelydiﬀerentiable.Remark1.3(Dependenceonoand(ψk)k∈N).NotethattheV-markedGromov-weak#topologydoesnotdependonthechoiceo∈Vandthesequence(ψk)k∈N,aslongasψkhasboundedsupportandAk:={v:ψk(v)=1}increasestoVask→∞.Remark1.4(MVasasubspaceof(MV
f)N).LetMV
fdenotethespaceof(equivalentclassesof)V-mmmspaceswithﬁnitemeasures,equippedwiththeV-markedGromov-weaktopologyasintroducedin[DGP11,Def.2.4].Thenitisawell-knownfactthateachelement
(X,r,µ)∈MVcanbeidentiﬁedwithasequence(
(X,r,ψ1·µ),
(X,r,ψ2·µ),...)intheproductspace(MV
f)N,equippedwiththeproducttopology.ThisidentiﬁcationallowsustoeasilydeducemanypropertiesofMVfrompropertiesofMV
fthatwereestablishedin[DGP11].Inparticular,wecanmetrizetheV-markedGromov-weak#topologyonMVbyintroducingametric(whichcanbecalledV-markedGromov-Prohorov#metric)(1.7)dMGP#(
(X1,r1,µ1),
(X2,r2,µ2)):=∞Xk=1dMGP(
(X1,r1,ψk·µ1),
(X2,r2,ψk·µ2))∧1
2k,wheredMGPisthemarkedGromov-ProhorovmetriconMV
f,whichwasintroducedin[DGP11,Def.3.1]andmetrizesthemarkedGromov-weaktopology.TheproofofthefollowingresultisinAppendixA.Theorem1.5(Polishspace).ThespaceMV,equippedwiththeV-markedGromov-weak#topology,isaPolishspace.
CONTINUUMSPACELIMITOFIFVGENEALOGIES7
PointsinMV,aswellasweakconvergenceofMV-valuedrandomvariables,canbedeterminedbytheso-calledpolynomialsonMV,whicharedeﬁnedviasamplingaﬁnitesubsetontheV-mmmspace.Deﬁnition1.6(Polynomials).Letn∈N.Letg∈Cbb(Vn,R),thespaceofreal-valuedboundedcontinuousfunctiononVnwithboundedsupport.Fork∈N∪{0,∞},letφ∈Ckb(R(n2)+,R),thespaceofk-timescontinuouslydiﬀerentiablereal-valuedfunctionsonR(n2)+withboundedderivativesuptoorderk.WecallthefunctionΦn,φ,g:MV→Rdeﬁnedby(1.8)Φn,φ,g(
(X,r,µ)):=Z···Zφ(r
)g(v
)µ⊗n(d(x
,v
)),amonomialofordern,wherev
:=(v1,...,vn),x
:=(x1,...,xn),r
=r
(x
):=(r(xi,xj))1≤i<j≤n,andµ⊗n(d(x
,v
))denotesthen-foldproductmeasureofµon(X×V)n.(a)LetΠk n:={Φn,φ,g:φ∈Ckb(R(n2)+,R),g∈Cbb(Vn,R)},whichwecallthespaceofmonomialsofordern(withdiﬀerentiabilityoforderk).LetΠk
0bethesetofconstantfunctions.WethendenotebyΠk:=∪n∈N0Πk nthesetofallmonomials(withdiﬀerentiabilityoforderk).(b)ForV=Rdandk,l∈N∪{0,∞},wedeﬁneΠk,l n:={Φn,φ,g:φ∈Ckb(R(n2)+,R),g∈Clbb(Vn,R)},andΠk,l:=∪n∈N0Πk,l n.(c)Wecallthelinearspaces(1.9)eΠk,`generatedbyΠk,`,thepolynomials(withdiﬀerentiabilityofφ,resp.g,oforderk,resp.`).Remark1.7.Notethatthepolynomialsformanalgebraofboundedcontinuousfunc-tionssincetheproductoftwomonomialscanbeseenasanewmonomialasdeﬁnedin(1.8).However,thesumoftwomonomialsingeneralisnotamonomial.Throughoutthepaperweareofteninterestedinthefollowingsub-spaceofMV(com-pare,e.g.,Deﬁnition1.34ofso-calledregularspaceofstates).LetbbeameasurablefunctiononR+,andwrite(1.10)M(V,≤b):=
(X,r,µ)∈MV:µ(X×Br(o))≤b(r) .IfβisameasureonB(V)whichisﬁniteonboundedsubsets,and(1.11)M(V,β):=
(X,r,µ)∈MV:µ(X×·)=β ,thenobviouslyM(V,β)isaclosedsubspaceofM(V,≤b)withb(r):=β(Br(o)).Theorem1.8(Convergencedeterminingclass).Fixameasurablefunctionβ:R+→R+.WehavethefollowingpropertiesforΠk,foreachk∈N∪{0,∞}:(i)ΠkisconvergencedetermininginM(V,≤b).Namely,
(Xn,rn,µn)→
(X,r,µ)inM(V,≤b)ifandonlyifΦ(
(Xn,rn,µn))→Φ(
(X,r,µ))asn→∞forallΦ∈Πk.(ii)ΠkisalsoconvergencedetermininginM1(M(V,≤b)),thespaceofprobabilitymea-suresonM(V,≤b).Namely,asequenceofM(V,≤b)-valuedrandomvariables(Xn)n∈NconvergesweaklytoaM(V,≤b)-valuedrandomvariableXifandonlyifE[Φ(Xn)]→E[Φ(X)]asn→∞forallΦ∈Πk.
8ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
(iii)ForV=Rdandeachk,l∈N∪{0,∞},Πk,lisalsoconvergencedetermininginM(V,≤b)andM1(M(V,≤b)).WedefertheproofofTheorems1.5and1.8,aswellassomeadditionalpropertiesofV-mmmspaces,toAppendixA.Remark1.9.Forthemodelsweconsider,thegenealogieslieincertainparticularPolishspaceswhichariseasclosedsubspacesofMV.Notethatthecurrentpopulationalivecorrespondstotheleavesofagenealogicaltree,andtheassociatedV-mmmspaceisultrametric.WewilldenotethespaceofV-markedultrametricmeasurespacesbyUV.TheyformaclosedsubspaceofMVandhenceUVisPolish.ThesameholdsforM(V,β),forsomeBorelmeasureβonVwhichisﬁniteonboundedsubsets.Remark1.10.RecallMV
fctfromDeﬁnition1.1,andnoticethatMV
fctisnotclosed,andthatwethereforechoosethebiggerspaceMVasthestatespace.ThespaceMVallowsanindividualx∈Xtocarryasetofmarks,equippedwiththeconditionalmeasureofµonVgivenx∈X.Ifeachindividualcarriesonlyasinglemarkwhichwecanidentifyviaamarkfunctionκ:X→V,thecorrespondingmarkedmetricmeasurespaceisanelementofMV
fct.ThiswillbethecasefortheinteractingFleming-Viotprocessthatwewillstudy.ItcanbeshownthateveryelementinMV
fctisanelementoftheclosedspace(1.12)MV
h:=[Abdd\δ>0(X,r,µ)∈MV:∃µ0
A∈Mf(X×A):µ0≤µ(·×A),kµ0−µ(·×A)kTV≤hA(δ),(µ0
A)⊗2{r(x1,x2)<δ,d(v1,v2)>hA(δ)}=0 ,forsomehA∈H,where(1.13)H:={h:R+→R+∪{∞}:hiscontinuousandincreasing,andh(0)=0}([KL15,Lemma2.8]).Ω1.3.InteractingFleming-Viot(IFV)GenealogyProcesses.Wenowstudythege-nealogiesofthemeasure-valuedinteractingFleming-Viot(IFV)processesonacountablegeographicspaceandwithallelictypes,typicallytakenfromthetypespace[0,1](see[DGV95]fordetailsonIFV),whichismotivatedbythefollowingindividual-basedmodel,theso-calledMoranmodel.Considerapopulationofindividuals,X,withlocationsindexedbyacountableadditivegroupV(forusthislaterwillbeZ).Theindividualsmigrateindependentlyaccordingtorateonecontinuoustimerandomwalkswithtransitionprobabilitykernela(·,·),(1.14)a(v1,v2)=a(0,v2−v1)forallv1,v2∈V.Wedenotethetransitionkernelofthetimereversedrandomwalksby(1.15)¯a(v1,v2):=a(v2,v1)=a(0,v1−v2).Individualsfurthermorereproducebyresampling,whereeverypairofindividualsatthesamesitediestogetheratexponentialrateγ>0,andwithequalprobability,oneofthetwoindividualsischosentogivebirthtotwonewindividualsatthesamesitewiththesametypeastheparent.Thisnaturallyinducesagenealogicalstructure.Thegenealogicaldistance,r,oftwoindividualsattimetis2min{t,T}plusthedistanceoftheancestorsattime0,whereTisthetimeittakestogobacktothemostrecentcommonancestor.ImposingtheHaarmeasure(herethecountingmeasure),µ,onthepopulationwitheachindividualcarryingitslocationasamark,weobtainaV-mmmspaceanditsequivalenceclassisthestateofthegenealogyprocess.
CONTINUUMSPACELIMITOFIFVGENEALOGIES9
Lettingnowthenumberofindividualspersitetendtoinﬁnityandnormalizingthemeasuresuchthateachsitecarriespopulationmassoforderone,weobtainadiﬀusionmodel,theinteractingFleming-Viot(IFV)genealogyprocesswithstatespace(1.16)UV
1:=U(V,n),wherendenotestheHaarmeasureonthecountablegeographicspaceV,the1indicatingthatthemeasurerestrictedtoeachcolonyisaprobabilitymeasure.This(seeRemark1.9)isagainaPolishspace.(ForthediﬀusionlimitinthecaseofaﬁnitegeographicspaceV,see[GPW13],[DGP12]).Remark1.11.Ifweintroduceasmarks(besideslocationsfromacountablegeographicspaceG)alsoallelictypesfromsomesetK(typicallytakenasaclosedsubsetof[0,1]),thenthetypeisinheritedatreproductionandV=K×Gistheproductoftypespaceandgeographicspace.InthiscasethelocalizationprocedureinDeﬁnition1.2appliestothegeographicspace,sinceKiscompact.1.3.1.ThegenealogicalIFVmartingaleproblem.WenowdeﬁnetheinteractingFleming-Viot(IFV)genealogyprocessviaamartingaleproblemforalinearoperatorLFVonCb(UV
1,R),actingonpolynomials.Forsimplicityweﬁrstleaveoutallelictypes,whichweintroducelaterinRemark1.13.ThelinearoperatorLFVonCb(UV
1,R),withdomainΠ1,0asintroducedinDeﬁni-tion1.6(b)withd=1,consistsofthreeterms,correspondingrespectivelytoaging,migration,andreproductionbyresampling.WithX=
(X,r,µ)andΦ=Φn,φ,g∈Π1,0,LFVΦ(X)=2Z(X×V)nµ⊗n(d(x
,v
))g(v
)X1≤k<`≤n∂
∂rk,`φ(r
)+Z(X×V)nµ⊗n(d(x
,v
))φ(r
)nXj=1Xv0∈V¯a(vj,v0)(Mvj,v0g−g)(v
)
(1.17)
+2γZ(X×V)nµ⊗n(d(x
,v
))g(v
)X1≤k<`≤n1{vk=v`}(θk,`φ−φ)(r
),wherex
=(x1,···,xn),v
=(v1,···,vn),r
:=(rk,`)1≤k<`≤n:=(r(xk,x`)1≤k<`≤n),and(1.18)(Mvj,v0g)(v1,···,vn):=g(v1,···,vj−1,v0,vj+1,···,vn)encodesthereplacementbymigrationofthej-thsampledindividualcorrespondingtoajumpfromlocationvjtov0,whileθk,`encodesthereplacementofthe`-thindividualbythek-thindividual(bothatthesamesite).Moreprecisely,(1.19)(θk,`φ)(r
):=φ(θk,`r
),with(θk,`r
)i,j:=



r(xi,xj)ifi,j6=`,r(xi,xk)ifj=`,r(xk,xj)ifi=`.TheﬁrstresultstatesthatthereisauniqueUV
1-valueddiﬀusionprocessassociatedwiththisoperator.
Theorem1.12(MartingaleproblemcharacterizationofIVFGenealogyprocesses).ForanyX0=
(X0,r0,µ0)∈UV
1,wehave:(i)The(LFV,Π1,0,δX0)-martingaleproblemiswell-posed,i.e.thereexistsaUV
1-valuedprocessXFV:=(XFVt)t≥0,uniqueinitsdistribution,whichhasinitialcondition
10ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
X0andc`adl`agpath,suchthatforallΦ∈Π1,0andw.r.t.thenaturalﬁltrationgeneratedby(XFVt)t≥0,(1.20)Φ(XFVt)−Φ(XFV0)−tZ0(LFVΦ)(XFVs)dst≥0isamartingale.(ii)Thesolutions(forvaryinginitialconditions)deﬁneastrongMarkovwithcontin-uouspath.ThisMarkovprocesshastheFellerproperty,i.e.,theone-dimensionaldistributionsdependcontinuouslyontheinitialdistribution.(iii)Iftheinitialstateadmitsamarkfunction,thensodoesthepathalmostsurely.Remark1.13.Ifweaddthetypeofanindividualasanadditionalmark,i.e.,V:=G×Kwithgeographicandtypespacerespectively,thenthesameresultholdsifwemodifyasfollows.WerequirethestatestosatisfytheconstraintthattheprojectionofµonthegeographicspaceVisstillthecountingmeasure.ThetestfunctionsΦshouldbemodiﬁedsothatwemultiplyg:Gn→R,actingonthelocationsofthensampledindividuals,byanotherbounded,continuousfactorgtyp:Kn→R,actingonthetypesoftheindividuals.ThegeneratorLFVshouldbemodiﬁedaccordingly,sothatgtypchangesatresamplingfromgtyptogtyp◦eθk,`,witheθk,`replacingthe`-thsampledindividualbythek-thone,see[DGP12]).
Remark1.14.TheprocessXFVt=
(Xt,rt,µt)hasthepropertythatthemeasure-valuedprocessbXtgivenbythecollection{µt(Xt×{i}×·),i∈G},istheIFVprocesson(M1(K))G.Remark1.15.FromSection1.4onward,wewillchooseV=Z.Howeverinthesubsequentanalysis,itisimportanttoobservethatwecanembedZintoRandviewUZasaclosedsubspaceofUR,andviewtheIFVgenealogyprocessasUR-valuedprocess.1.3.2.Duality.TheIFVgenealogyprocessXFVcanbecharacterizedbyadualprocess,thespatialcoalescent.Theformulationgivenherecanalsoincorporatemutationandselection(see[DGP12]).Foreachn∈N,let˜Sndenotethespaceofpartitionsoftheset{1,...,n},i.e.,π={p;p∈π}∈˜Snisacollectionofdisjointsubsetsp⊆{1,...,n},referredtoaspartitionelements(orblocks),suchthat]p∈πp={1,...,n}.Moreover,let˜SV
ndenotethespaceofmarkedpartitionsoftheset{1,...,n}withmarkspaceV,i.e.,weregard{(p,vp);π={p;p∈π}∈˜Sn,vp∈V}∈˜SV
nasapartitionof{1,...,n}whereeachpartitionelementisassignedamarkinV.Finally,letSn:=˜SV
n×R(n2)+denotethespaceofhistoricalmarkedpartitionsoftheset{1,...,n}withmarkspaceV.Thatis,weregard({(p,vp);p∈π},(rij)1≤i<j≤n)∈Snasconsistingofamarkedpartitionandamatrixofmutualdistances.Thedualprocess(Kt)t≥0willtakevaluesinS:=Sn∈NSnwiththefollowingdynamics:givenaﬁnitehistoricalmarkedpartition,independentlyeverypairofpartitionelementswiththesamemarkinVmergesatrateγ.Untilapairofpartitionelementsmerges,themarksmigrateindependentlyofeachotheronVaccordingtoacontinuoustimerandomwalkwithtransitionkernel¯a.Aftermergingthemarksofthetwoinvolvedpartitionelementswillmovetogetherforever.Attimet,wedeﬁnethegenealogicaldistancert(i,j)ofiandjin{1,2,...,n}as2min{t,Ti,j},whereTi,jistheﬁrsttimethatiandjbelongtothesamepartitionelement.
CONTINUUMSPACELIMITOFIFVGENEALOGIES11
Foreachn,φ∈C∞b(R(n2)+,R),andg∈Cbb(Vn,R),wedeﬁneadualityfunctionH:UV
1×Sn→Rwith(1.21)H X,K:=Z(X×V)|π|µ⊗|π|(d(x
,v
))1{vp=ξ0p;∀p∈π}·g(v
π)φ r
π(x
)+r
0,whereX=
(X,r,µ)∈UV
1,andK=({(p,ξ0p);p∈π},r
0)∈Sn,andv
π=(vp(i))i=1,...,n,andr
π(x
):=(r(xp(i),xp(j)))1≤i<j≤n,withp(i)beingthepartitionelementofπcontainingi∈{1,...,n}.Remark1.16.Notethat{H(·,K):K∈Sn,φ∈C∞b(R(n2)+,R),g∈Cbb(Vn,R),n∈N}islaw-determiningandconvergence-determiningonUV
1.TheIFVgenealogyprocess(XFVt)t≥0isdualtothecoalescent(Kt)t≥0,anditslawandbehaviorast→∞canbedeterminedasfollows.Theorem1.17(Dualityandlongtimebehaviour).ThefollowingpropertiesholdfortheIFVgenealogyprocess(XFVt)t≥0:(a)ForeveryXFV0∈UV
1andK0∈S,wehave(1.22)E[H(XFVt,K0)]=E[H(XFV0,Kt)],t≥0.(b)Ifba(·,·)=1
2(a(·,·)+¯a(·,·))isrecurrent,then(1.23)L[XFVt]=⇒t→∞Γ∈M1(UV
1),whereΓistheuniqueinvariantmeasureoftheprocessXFVonUV
1.Remark1.18.Ifbaistransient,thenwecandecomposeXFVt=
(Xt,rt,µt)insuchawaythatXtisthecountableunionofdisjointsetsXit,µtisthesumofmeasuresµi t,and{(Xit,rt|Xit×Xit,µi t)}i∈NareV-mmmspacessuchthatforx∈Xit,x0∈Xjtwithi6=j,rt(x,x0)divergesinprobabilityast→∞,andeach(Xit,rt|Xit×Xit,µi t)convergesinlawtoalimitingV-mmmspace.Alternativelywecantransformdistancesr:r→1−e−randobtainauniqueequilibriuminthatcase.Seealso[GM].
Remark1.19(Strongduality).AstheinteractingFleming-Viotprocessisapopulationmodel,itsevolvinggenealogy(Ut,rt,µt)t≥0canberepresentedbyaV-markedR-tree(X,r).Thatis,wecanﬁnda0-hyperbolicmetricspace(withdistances∞possibleequalto)(X,r)suchthatanytwox,y∈Xofﬁnitedistanceareconnectedbyapath,andisometries(ϕt)t≥0withϕt:Ut→Xsuchthatforallx∈Xthereisat≥0withϕ−1t(x)∈supp(µt(·×V)).Wewon’twritedownthisR-treerepresentationexplicitlyherebutitcanbederivedfromthelook-downconstructionpresentedin[GLW05]inastraightforwardway.WewouldliketopointoutthatthisrepresentationbyanR-treeallowstodeﬁnetheancestorAT
sofxattimeTbackattimeT−s,andtheabovedualityrelationcanactuallybestatedinastrongsense.Thatis,wecanconstructforeachT>0ourmodeltogetherwithadualprocess˜KTonthesameprobabilityspacesuchthatforallX0⊆UTwith#X0<∞,φ∈C∞b(R(#X02)+,R),andg∈Cbb(V#X0,R),andH=H#X0,ϕ,gfrom(1.21),foralls∈[0,T],(1.24)H UT−s,˜KTs≡H UT,˜KT0,almostsurely.Noticethatwheneveradualityrelationholdsalmostsurelyratherthaninexpectation,onereferstoitasastrongformofduality.
12ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
Indeed,deﬁneforﬁxedT≥0andanyﬁnitesubsetX0:={x1,...,x#X0}⊂UTthemap˜KT:=(˜KTs:=(({(p,ξ0p);p∈πTs},r0,Ts))s∈[0,T]whichsendsU:=(Ut:=(Ut,rt,µt))t≥0toapathwithvaluesinthespaceofhistoricalmarkedpartitionsS#X0deﬁnedby(1.25)r0,Ts(i,j)=r ϕT(xi),ϕT(xj)∧(T−s),forall1≤i<j≤#X0,amarkedpartitionof{1,2,...,n}deﬁnedthroughtheequivalencerelationΩ(1.26)i≡T
sjiﬀr0,Ts(i,j)<(T−s),andafamilyofpositionfunctionsonϕT(supp(µT(·×V)))suchthat(1.27)ξ0,Ts(p):=κT−s ϕ−1T−s(p).Byconstruction,˜KTisthedualspatialcoalescent.Moreover,foralls∈[0,T],and(x
,v
)∈(XT×V)|π0|,(1.28)r
πsT−s(AT
s(x
))+r
0T
s=r
π0T(x
)+r
0T
0,andsκT−s(AT
s(x))hasunderµTthesamedistributionastypesunderµT−s.Thistogetherimpliesthestrongformofdualityasstatedin(1.24).
Asafurtherconsequenceofrelation(1.22)wecanspecifytheﬁnitedimensionaldistri-butionsof(XFVt)t≥0completelyintermsofthespatialcoalescentasfollows.FixatimehorizonT>0.TheﬁnitedimensionaldistributionsaredeterminedbytheexpectationofalltestfunctionseΦ:C([0,∞),MV)ofthefollowingform:(1.29)eΦ (XFVt)t≥0:=Π`
k=1Φk XFVtk,forsome`∈N,0≤t1<t2<...<t`=tandΦk=Φnk,gk,φk∈Π1,0ofordernk∈N,φk∈C∞b(R(nk2)+,R),andgk∈Cbb(Vnk,R),foreachk=1,...,`.Thedualisthespatialcoalescentwithfrozenparticles(eKs)s∈[0,T],forwhichthetimeindexs∈[0,T]runsintheoppositedirectionfromthetimeindexofXFV.Namely,lookingbackwardfromtimeT,foreach1≤k≤`,westartnkparticlesattimeT−tkatlocationsξ1tk:=(ξ1tk,1,...,ξ1tk,nk)∈Vnk,eachformingitsownpartitionelementinthepartitionπ∈˜Sn,n:=n1+...+n`.TheparticlesperformtheusualdynamicsofthespatialcoalescentwiththerestrictionthatallparticlesstartingattimeT−tkwerekeptfrozenbeforetimeT−tk.Attimes,thegenealogicaldistancers(i,j)betweentwoindividualsiandj,startedrespectivelyattimesT−tiandT−tj,isdeﬁnedtobe(1.30)2min{s,Ti,j}−min{s,(T−ti)}−min{s,(T−tj)},whereTi,jistheﬁrsttimethetwoindividualscoalesce.Denotethisnewspatialcoalescentprocesswithfrozenparticlesby(eKt)t≥0.ThestatespaceSisoncemorethespaceofhistoricalmarkedpartitions.WethendeﬁnethedualityfunctionH:UV
1×S→RwhichdeterminestheﬁnitedimensionaldistributionsofXforvaryingeKasin(1.21)butnowwith(1.31)φ(r
):=`Yk=1φk (ri,j)n1+···+nk−1<i<j≤n1+···+nkandg v
π:=`Yk=1gk tk,v
πk,whereπkarethepartitionelementsstartedattimetk.Thefollowingspace-timedualityisanimmediateconsequenceoftheMarkovpropertyandthedualityappliedsuccessivelytothetimeintervals[t`−1,t`],[t`−2,t`−1],...,[0,t1].
CONTINUUMSPACELIMITOFIFVGENEALOGIES13
Corollary1.20(Space-timeduality).LetXFV0∈UV
1,andeK0∈Sbeasin(??)withtk≤Tforall1≤k≤`.LeteΦbedeﬁnedasabove(1.30).Then(1.32)EeΦ (XFVs)s∈[0,T]=EH XFV0,eKT.Inwords,thegenealogicaldistancesamongthen=n1+···+n`individualssampledfrom(XFVs)s∈[0,T],withnkindividualssampledattimetkatspeciﬁedlocations,canberecoveredbylettingtheseindividualsevolvebackwardintimeasaspatialcoalescentuntiltime0,atwhichpointwesamplefromXFV0accordingtothelocationofeachpartitionelementinthespatialcoalescent.Ω1.4.GenealogiesofContinuum-sitesSteppingStoneModel(CSSM)onR.Ifwerescalespaceandtimediﬀusively,themeasure-valuedinteractingFleming-ViotprocessonZconvergestoacontinuumspacelimit,theso-calledcontinuum-sitessteppingstonemodel(CSSM).Formally,CSSMisameasure-valuedprocessν:=(νt)t≥0onR×[0,1],whereRisthegeographicalspaceand[0,1]isthetypespace.WemightthinkofindividualsinthepopulationwhichundergoindependentBrownianmotions,andwhenevertwoindividualsmeet,oneofthetwoindividualsischosenwithequalprobabilityandswitchesitstypetothatofthesecondindividual.Providedthatν0(·×[0,1])istheLebesguemeasureonR,CSSMwasrigorouslyconstructedin[EF96,Eva97,DEF+00,Zho03]viaamomentdualitywithcoalescingBrownianmotions.Inparticular,νt(·×[0,1])istheLebesguemeasureonR,forallt>0.InthissubsectionwewillconstructtheevolvinggenealogiesoftheCSSMbasedondual-itytothe(dual)Brownianweb,andestablishproperties(Proposition1.25,Theorem1.27).ForthevotermodelonZ,thejointgenealogylinesofindividualsatallspace-timepointsinZ×[0,∞)aredistributedasacollectionofcoalescingrandomwalksevolvingbackwardintime(see[Lig85]).Analogously,fortheCSSMonR,whennindividualsaresampledfromthepopulationatpossiblydiﬀerenttimes,theirjointgenealogylinesevolvebackwardintimeascoalescingBrownianmotions.Uponreachingtime0,eachsurvivinggenealogylinethenindependentlyselectsanancestraltypebysamplingaccordingtotheconditionaldistributionofν0onthetypespace[0,1],conditionedonthespatiallocationofthegenealogyline.Furthermorethejointgenealogylinesofindividualsatallspace-timepointsinR×[0,∞)aredistributedasacollectionofcoalescingBrownianmotionsevolvingbackwardintime.ThereforetheCSSMonRisexactlythecontinuumanalogueoftheinteractingFleming-Viotprocess(aswellasthevotermodel)onZ.Althoughhavinganuncountablenumber(startingfromeveryspace-timepointinR×[0,∞))ofcoalescingBrownianmotionsseemsproblematic,thisobjecthasbeencon-structedrigorouslyandisnowknownasthe(dual)BrownianwebcW[FINR04,FINR06],dualtoaforwardBrownianwebWconstructedonthesameprobabilityspace.The(dual)BrownianwebisessentiallyacollectionofcoalescingBrownianmotionsonR,startingfromeverypointinthespace-timeplaneR×R.In[FINR04],theBrownianwebWisconstructedasarandomvariablewhereeachrealizationofΩ(1.33)WisaclosedsubsetofΠ:=∪s∈RC([s,∞),R),thespaceofcontinuouspathsinRwithanystartingtimes∈R.In[FINR04],thetopologyisdeﬁnedbyﬁrstcompactifyingthespaceR2suitably,andthenbychoosingforΠthetopologyoflocaluniformconvergenceandrequiringtheinitialtimestobecloseifpathsareclose.Foreachz:=(x,t)∈R,wewillletW(z):=W(x,t)denotethesubsetofpathsinWwithstartingpositionxandstartingtimet.WecanconstructWbyﬁrstﬁxingacountabledensesubsetD⊂R2,andthenconstructacollectionofcoalescingBrownian
14ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
motions{W(z):z∈D},withoneBrownianmotionstartingfromeachz∈D.TheBrownianwebWisthenobtainedbytakingasuitableclosureof{W(z):z∈D}inΠ,whichgivesrisetoasetofpathsstartingfromeverypointinthespace-timeplaneR2.ItcanbeshownthatthelawofWdoesnotdependonthechoiceofD(seeTheoremB.1inAppendixB).TheBrownianwebWhasagraphicaldualcalledthedualBrownianweb,whichwedenotebycW.Formally,(1.34)cWisarandomclosedsubsetofbΠ:=∪s∈RC((−∞,s],R),thespaceofcontinuouspathsinRstartingatanytimes∈RandrunningbackwardintimeascoalescingBrownianmotions.ThejointdistributionofcWandWisuniquelydeterminedbytherequirementthatthepathsofcWnevercrosspathsofW(see,TheoremB.3inAppendixB).Thus,jointly,theBrownianwebanditsdualisarandomvariabletakingvaluesinaPolishspace,with(1.35)(W,cW)∈Π×bΠ.InterpretingcoalescingBrownianmotionsinthe(dual)Brownianwebasancestrallinesspecifyingthegenealogies,wecanthengiveanalmostsuregraphicalconstructionoftheCSSM,insteadofrelyingonmomentdualityrelationsasin[EF96,Eva97,DEF+00,Zho03],whichneverthelesswegetascorollaryofthegraphicalconstruction.Theclassicalmeasure-valuedCSSMprocesscanberecoveredfrom(XCSt)t≥0byprojectingthesamplingmeasure(µCS
t)t≥0tothemarkspaceV,ifVischosentobetheproductofgeographicalspaceRandtypespace[0,1].Inwhatfollows,wewilltakeVtobeonlythegeographicalspaceR,sincetypeshavenoinﬂuenceontheevolutionofgenealogies.WenextexplicitlyconstructaversionoftheCSSMgenealogyprocess(1.36)XCS:=(XCSt)t≥0,XCSt=
(XCSt,rCSt,µCS
t),t≥0,asafunctionalofthe(dual)Brownianweb(W,cW).Toavoidadisruptionoftheﬂowofpresentation,backgrounddetailsonthe(dual)BrownianwebwewillneedarecollectedinAppendixB.Weproceedinthreesteps:Step1(Initialstates).AssumethatXCS0belongstothefollowingclosedsubspaceofUR:(1.37)UR
1:={
(X,r,µ)∈UR:µ(X×·)istheLebesguemeasureonR}.Inotherwords,UR
1isthesetofR-markedultrametricmeasurespaceswheretheprojectionofthemeasureonthemarkspace(geographicspace)RistheLebesguemeasure.ThisisnecessaryforthedualitybetweenCSSMandcoalescingBrownianmotions.WewillseethatalmostsurelyXCSt∈UR
1forallt≥0.Step2(Thetime-tgenealogyasametricspace).Todeﬁne(XCSt,rCSt)foreveryt>0,letusﬁxarealizationof(W,cW),(seetheAppendixBformoredetails).Foreacht>0,let(1.38)At:=v∈R:cW(v,t)containsasinglepathˆf(v,t) ,Et:=R\At.ByLemmaB.4ontheclassiﬁcationofpointsinR2w.r.t.WandcW,almostsurely,Etisacountablesetforeacht>0.Foreachv∈At,weinterpretˆf(v,t)asthegenealogylineoftheindividualatthespace-timecoordinate(v,t).Genealogylinesofindividuals
CONTINUUMSPACELIMITOFIFVGENEALOGIES15
atdiﬀerentspace-timecoordinatesevolvebackwardintimeandcoalescewitheachother.Attime0,eachgenealogylinetracesbacktoexactlyonespatiallocationintheset(1.39)bE:={ˆf(v,s)(0):v∈R,s>0,ˆf∈cW},wherewenotethatbEisalmostsurelyacountableset,becausebyTheoremB.1andLemmaB.2,pathsincWcanbeapproximatedinastrongsensebyacountablesubsetofpathsincW.Foreachv∈bE,wethenidentifyacommonancestorξ(v)foralltheindividualswhosegenealogylinestracebacktospatiallocationvattime0,bysamplinganindividualξ(v)∈XCS0accordingtotheconditionaldistributionofµCS
0onXCS0,conditionedonthespatialmarkintheproductspaceXCS0×Rbeingequaltov.Wenextcharacterizeindividualsbypointsinspace.Notethatthereisanaturalge-nealogicaldistancebetweenpointsinAt.Forindividualsx,y∈At,ifˆf(x,t)andˆf(y,t)coalesceattimeˆτ<t,thendenotingu:=ˆf(x,t)(0)andv:=ˆf(y,t)(0),wedeﬁnethedistancebetweenxandyby(1.40)rt(x,y):=(2(t−ˆτ)ifˆτ≥0,2t+rCS0(ξ(u),ξ(v))ifˆτ<0.Firstdeﬁne(XCSt,rCSt)astheclosureofAtw.r.t.themetricrtdeﬁnedin(1.40).Notethat(XCSt,rCSt)isultrametric,andbyconstructionPolish.Remark1.21.Wemayevenextendthedistancerttoadistancerbetween(x,t)withx∈Atandt>0,and(y,s)withy∈Asands>0.Moreprecisely,let(1.41)r((x,t),(y,s)):=(s+t−2ˆτifˆτ≥0,s+t+rCS0(ξ(u),ξ(v))ifˆτ<0.Step3(Addingthesamplingmeasure).WenowdeﬁneXCSt:=
(XCSt,rCSt,µCS
t).Forthatpurpose,wewillrepresentnextXCStasanenrichedcopyofR(see(1.42)below).Byidentifyingeachx∈Atwiththepathˆf(x,t)∈cW,wecanalsoidentifyXCStwiththeclosureof{ˆf(x,t)∈cW:x∈At}inbΠ,becauseasequencexn∈AtisaCauchysequencew.r.t.themetricrtifandonlyifthesequenceofpathsˆf(xn,t)isaCauchysequencewhenthedistancebetweentwopathsismeasuredbythetimetocoalescence,whichbyLemmaB.2,isalsoequivalentto(ˆf(xn,t))n∈NbeingaCauchysequenceinbΠ.Whenwetaketheclosureof{ˆf(x,t)∈cW:x∈At}inbΠ,onlyacountablenumberofpathsincWareadded,whicharepreciselytheleftmostandrightmostpathsincW(x,t),whencW(x,t)containsmorethanonepath.Namelyforeachx∈Et(recallfrom(1.38)),letx+,x−denotethetwocopiesofxobtainedbytakinglimitsofxn∈Atwitheitherxn↓xorxn↑xinR,andletE±t:={x±:x∈Et}.Wecanthentake(1.42)XCSt:=At∪E+t∪E−t,equippedwithametricrCSt,whichistheextensionofrtfromAttoitsclosureXCSt,givingaPolishspace(XCSt,rCSt).Nexttogetasamplingmeasure,notethateachﬁniteballin(XCSt,rCSt)withradiuslessthantcanbeidentiﬁedwithanintervalinR(moduloasubsetofEt∪E+t∪E−t),andhencecanbeassignedtheLebesguemeasureofthisinterval.Thatis,wedeﬁnetheBorelmeasureeµCS
ton(XCSt,rCSt)by(1.43)eµCS
t {x0∈XCSt:rCSt(x0,x)<δ}:=R(x,δ)−L(x,δ),δ<t,x∈XCSt
16ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
if{x0∈XCSt:rCSt(x0,x)<δ}=(L(x,δ),R(x,δ))⊆R.WethendeﬁnethesamplingmeasureµCS
tonB((XCSt,rCSt)×(R,deucl))by(1.44)µCS
t(dxdv):=eµCS
t(dx)δx(dv).ThiscompletesourconstructionoftheCSSMgenealogyprocess(XCSt)t≥0.Remark1.22(Notationalsimpliﬁcation).Inthesequelwewillapplytheexistenceofamarkfunctionandtheembeddingofthebasicsetintheenrichedrealstosimplifynotation.Namely,wewillwriteintegralsw.r.t.thesamplingmeasureoverX×Vasanintegralw.r.t.theLebesguemeasureoverR.Akeyfeatureisagainduality.WecanreplaceintheDeﬁnitionoftheprocess(Kt)t≥0fromSubsubsection1.3.2therandomwalksbyBrownianmotionstoobtain(KBrt)t≥0,whichisdualtotheprocess(XCSt)t≥0byconstruction:Corollary1.23.(H-duality)Thetree-valuedCS-processandmarkedtree-valuedcoalesc-ingBrownianmotionsareinH-duality,i.e.,foreachHoftheform(1.21),(1.45)E[H(XCSt,KBr0)]=E[H(XCS0,KBrt)].Furthermorebytheaboveconstructionstrongdualityholds.Remark1.24.NoteinthecontinuumcasethefunctionΦn,g,ϕ(·,κBr)isnotapoly-nomialsinceweﬁxthelocationsweconsider.Inordertogetapolynomialweneedtoconsiderafunctiongonmarkspacewithg∈C2bb(R,R)overwhichweintegratew.r.t.thesamplingmeasure.WecollectbelowsomebasicpropertiesfortheCSSMgenealogyprocessesthatwejustconstructed.
Proposition1.25(Regularityofstates).LetXCS:=(XCSt)t≥0betheCSSMgenealogyprocessconstructedfromthedualBrownianwebcW,withXCS0∈UR
1.Thenalmostsurely,foreveryt>0:(a)Thereexistsacontinuous(mark)functionκt:XCSt→R,i.e.,µCS
t(dxdv)=µCS
t(dx×R)δκt(x)(dv);(b)Foreach`∈(0,t),XCStisthedisjointunionofballs(B`i)i∈Zofradius`.Fur-thermore,thereexistsalocallyﬁnitesetE`t:={vi}i∈Z⊂Rwithvi−1<viforalli∈Z,suchthat{κ(x):x∈B`i}=[vi−1,vi]andµCS
t(B`i)=vi−vi−1.WecanidentifyE`tfromcWby(1.46)E`t:={x∈Et:ˆf+(x,t)andˆf−(x,t)coalesceatsometimes≤t−`},wheref+(x,t)andf−(x,t)arerespectivelytherightmostandleftmostpathincW(x,t);(c)XCSt=
(XCSt,rCSt,µCS
t)∈UR
1,i.e.,µCS
t(XCSt×·)istheLebesguemeasureonR.Remark1.26.UsingthedualitybetweentheBrownianwebWandcW,ascharacterizedinAppendixB,itiseasilyseenthatwecanalsowriteΩ(1.47)Elt={f(t):f∈W(x,t−s)forsomex∈R,s≥l}.Theorem1.27(Markovproperty,pathcontinuity,Fellerproperty).LetXCSbeasinProposition1.25.Then(a)(XCSt)t≥0isaUR
1-valuedMarkovprocess;(b)Almostsurely,XCStiscontinuousint≥0;
CONTINUUMSPACELIMITOFIFVGENEALOGIES17
(c)Foreachm∈N,letXCS,(m)beaCSSMgenealogyprocesswithXCS,(m)0∈UR
1.IfXCS,(m)0→XCS0inUR
1,thenforanysequencetm→t≥0,wehaveweakconvergenceinlaw,i.e.XCS,(m)tm⇒XCSt,asm→∞.(d)ForeachinitialstateinUR
1,L[XCSt]convergesweaklytotheuniqueequilibriumdistributiononUR
1ast→∞.Remark1.28.Noticethatthisequilibriumstatecanberepresentedintermsofafunc-tionaloftheBrownianweb.
Remark1.29.Wenotethatifweallowtypesaswell,weenlargethemarkspacefromRtoR×[0,1],whereeachindividualcarriesatypein[0,1]thatisinheriteduponresampling.Theorem1.27stillholdsinthiscase.WewillconsidersuchanextendedmarkspaceinTheorem1.32below.
Remark1.30.Proposition1.25showsthateventhoughXCS0canbeanystateinUR
1,fort>0,XCStcanonlytakeonasmallsubsetofthestatespaceUR
1.Thisintroducescomplicationsinestablishingthecontinuityoftheprocessatt=0,anditwillalsobeanimportantpointwhenwediscussthegeneratoroftheassociatedmartingaleproblem.1.5.ConvergenceofRescaledIFVGenealogies.Inthissubsection,weestablishtheconvergenceoftheinteractingFleming-ViotgenealogyprocessesonZtothatoftheCSSM,whereweviewthestatesaselementsofUR(seeRemark1.15).Weassumethatthetransitionprobabilitykernela(·,·)inthedeﬁnitionoftheIFVprocesssatisﬁes(1.48)Xx∈Za(0,x)x=0andσ2:=X
x∈Za(0,x)x2=X
x∈Z¯a(0,x)x2∈(0,∞).Foreach>0,wethendeﬁneascalingmapS=Sσ:UR→UR(dependingonσ)asfollows.LetX=
(X,r,µ)∈UR.Then(1.49)SX:=
(X,Sr,Sµ),where(Sr)(x,y):=2r(x,y)forallx,y∈X,andSµisthemeasureonX×Rinducedbyµandthemap(x,v)∈X×R→(x,σ−1v),andthenthemassrescaledbyafactorofσ−1.Moreprecisely,(1.50)(Sµ)(F)=σ−1µ{(x,−1σv):(x,v)∈F}forallmeasurableF⊂X×R.WehavethefollowingconvergenceresultforrescaledIFVgenealogyprocesses.Theorem1.31(ConvergenceofRescaledIFVGenealogies).LetXFV,:=(XFV,t)t≥0beanIFVgenealogyprocessonZwithinitialconditionXFV,0∈UZ
1,indexedby>0.Assumethata(·,·)satisﬁes(1.48),andSXFV,0→XCS0forsomeXCS0∈UR
1as→0.Then(SXFV,−2t)t≥0convergesasC([0,∞),UR)-valuedrandomvariableweaklytotheCSSMgenealogyprocessXCS:=(XCSt)t≥0as→0.ToproveTheorem1.31,wewillneedanauxiliaryresultofinterestinitsownontheconvergenceofrescaledmeasure-valuedIFVprocessestothemeasure-valuedCSSM.TheIFVprocesswithmarkspaceR×[0,1]isameasure-valuedprocess(bXFVt)t≥0,wherebXFVtisameasureonR×[0,1],anditsprojectiononRisthecountingmeasureonZ.Similarly,theCSSMwithmarkspaceR×[0,1]isameasure-valuedprocess(bXCSt)t≥0,wherebXCStisameasureonR×[0,1],anditsprojectiononRistheLebesguemeasureonR.DeﬁnebXFVandbXCSrespectivelyastheprojectionofthemeasurecomponentofXFVandXCS,projectedontothemarkspaceR×[0,1].Wethenhave
18ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
Theorem1.32(ConvergenceofRescaledIFVProcesses).LetbXFV,:=(bXFV,t)t≥0beameasure-valuedIFVprocessonZ,indexedby>0.Assumethata(·,·)satisﬁes(1.48),andSbXFV,0convergestobXCS0w.r.t.thevaguetopologyforsomebXCS0as→0.Then(SbXFV,−2t)t≥0convergesasC([0,∞),M(R×[0,1]))-valuedrandomvariableweaklytotheCSSMprocessbXCS:=(bXCSt)t≥0as→0.Asimilarconvergenceresulthaspreviouslybeenestablishedforthevotermodelin[AS11].Remark1.33.AscanbeseenfromtheaboveconvergenceresultsandtheregularitypropertiesofthelimitprocessinProposition1.25,onamacroscopicscale,thereareonlylocallyﬁnitelymanyindividualswithdescendantssurvivingforamacroscopictimeofδormore.Thisphenomenonleadsinthecontinuumlimittoadynamicdrivenbyathinsubsetofhotspotsonly.Forsimilareﬀectsinotherpopulationmodels,seeforexample[DEF+02b,DEF+02a].1.6.MartingaleProblemforCSSMGenealogyProcesses.Inthissection,weshowthattheCSSMgenealogyprocessessolvesamartingaleproblemwithasingulargenerator.ToidentifythegeneratorLCSfortheCSSMgenealogyprocess(XCSt)t≥0,wenotethatforallt>0,XCStsatisﬁestheregularitypropertiesestablishedinProposition1.25.WewillseethatLCSisonlywell-deﬁnedonΦ∈Π1,2evaluatedatpointsX∈URwithsuitableregularitypropertiesforΦandX.WenowformalizethesubsetofregularpointsinUR
1asfollows,whichsatisﬁesexactlythepropertiesinProposition1.25.
Deﬁnition1.34(RegularclassofstatesUR
r).LetUR
rdenotethesetofX=
(X,r,µ)∈URwhichsatisﬁesthefollowingregularityproperties:(a)X∈UR
1,i.e.,µ(X×·)istheLebesguemeasureonR;(b)thereexistsamarkfunctionκ:X→R,withµ(dxdv)=µ(dx×R)δκ(x)(dv);(c)thereexistsδ>0suchthatforeachl∈(0,δ),Xisthedisjointunionofballs(Bli)i∈Zofradiusl.Furthermore,thereexistsalocallyﬁnitesetEl:={vi}i∈Z⊂Rwithvi−1<viforalli∈Z,suchthat{κ(x):x∈Bli}=[vi−1,vi]andµ(Bli)=vi−vi−1.ByProposition1.25,UR
risclosedunderthedynamicsofXCS(i.e.,XCS0∈UR
rimpliesthatXCSt∈UR
rforallt>0),separable,metricmeasurablesubsetofthePolishspaceUR
1.However,itisnotcomplete.Notethattheﬁrstrequirementgivesrisetoaclosedset,thesecondrequirementisknowntogenerateameasurableset[KL15],anditisnothardtoseethatthethirdrequirementalsogeneratesameasurableset.Remark1.35.Similartothediscussionleadingto(1.42),forX∈UR
r,wecangivearepresentationonanenrichedcopyofRasfollows.Property(c)inDeﬁnition1.34impliesthatanytwodisjointballsinXaremappedbyκtotwointervals,whichoverlapatatmostasinglepointinElforsomel>0.Thereforeκ−1(x)mustcontainasinglepointforallx∈Rwithxnotin(1.51)E:=∪l>0El,andκ−1(x)containingtwoormorepointsimpliesthatxisinκ(B1)∩κ(B2)fortwodisjointballsinX.Bythesamereasoning,foreachx∈E,κ−1(x)mustcontainexactlytwopoints,whichwedenotebyx±,wherex+isalimitpointof{κ−1(w):w>x}andx−isalimitpointof{κ−1(w):w<x}.Similarto(1.42),wecanthenidentifyXwith(R\E)∪E+∪E−,whereE±:={x±:x∈E}.Withthisidentiﬁcationandwith(1.44),wecansimplifyournotation(withaslightabuse)andletµbethemeasureon
CONTINUUMSPACELIMITOFIFVGENEALOGIES19
(R\E)∪E+∪E−,whichassignsnomasstoE±andisequaltotheLebesguemeasureonR\E.WenowintroducearegularsubsetofΠ1,2neededtodeﬁnethegeneratorLCS.Deﬁnition1.36(RegularclassoffunctionsΠ1,2r).LetΠ1,2rdenotethesetofregulartestfunctionsΦn,φ,g∈Π1,2,deﬁnedasin(1.8),withthepropertythat:(1.52)∃δ>0s.t.∀i6=j∈{1,2,···,n},∂φ
∂ri,j((rk,l)1≤k<l≤n)=0∀ri,j∈[0,δ].WecannowspecifytheactionofthegeneratorLCSonregularΦevaluatedatregularpoints,namelyLΦ(X)existsforΦ=Φn,φ,g∈Π1,2randX=
(X,r,µ)∈UR
r.ByRemark1.35,wecanidentifyXwith(R\E)∪E+∪E−,whileµisidentiﬁedwiththeLebesguemeasureonR\E.ThegeneratorLCSisgivenby(1.53)LCSΦ(X):=LCS
dΦ(X)+LCS
aΦ(X)+LCS
rΦ(X),withthecomponentforthemassﬂow(migration)ofthepopulationonRgivenby(1.54)LCS
dΦ(X):=1
2ZXnφ(r
)∆g(x
)dx
,andthecomponentforagingofindividualsgivenbyΩ(1.55)LCS
aΦ(X):=2ZXng(x
)X1≤i<j≤n∂φ
∂ri,j(r
)dx
.Theseoperatorsarelinearoperatorsonthespaceofboundedcontinuousfunctions,Cb(UR
1,R),withdomainΠ1,2,andmapspolynomialstopolynomialsofthesameorder.Thecomponentforresamplingis,withθk,lφdeﬁnedasin(1.19),givenby(1.56)LCS
rΦ(X):=X1≤k6=l≤nZXn1{κ(xk)=κ(xl)}g(x
)(θk,lφ−φ)(r
)µ∗(dxk)µ∗(dxl)Y1≤i≤ni6=k,ldxi,witheﬀectiveresamplingmeasureandmarkfunctionsΩ(1.57)µ∗:=Xx∈Eδx++Xx∈Eδx−,(1.58)κ(x±)=xforx∈E,whereE,andx±forx∈E,aredeﬁnedasinRemark1.35.Remark1.37.NotethatLCS
rissingular.Firstbecausetheeﬀectiveresamplingmeasureµ∗issupportedonacountablesubsetofXandissingularw.r.t.thesamplingmeasureµonX.Secondly,µ∗islocallyinﬁnitebecauseE∩(a,b)containsinﬁnitelymanypointsforanya<b.Thereforether.h.s.of(1.56)isnowinprincipleasumofcountablymanymonomialsofordern−2.Indeed,aswepartitionX=(R\E)∪E+∪E−intoballsofradiusl,withl↓0,theballsmustcorrespondtosmallerandsmallerintervalsonRsoasnottocontradictthefactthateachpointinXisassignedonevalueinR.Nevertheless,LCS
rΦ(X)in(1.56)iswell-deﬁnedatleastonUR
rbecausebyourassumptionthatΦ∈Π1,2randcondition(1.52)onφ,wehaveθk,lφ=φiftheresamplingiscarriedoutbetweentwoindividualsx+andx−forsomex∈E,withr(x+,x−)≤δ.Thusonlyresamplinginvolvingx∈Ewithr(x+,x−)>δremainsintheintegrationw.r.t.µ∗,andsuchxarecontainedinthelocally
20ANDREASGREVEN1,RONGFENGSUN2,ANITAWINTER3
ﬁnitesetEδintroducedinDeﬁnition1.34(c

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

---
source: "https://arxiv.org/abs/1407.7720v1"
title: "Sample genealogy and mutational patterns for critical branching populations"
author: "G. Achaz, C. Delaporte, A. Lambert"
year: "2014"
publication: "arXiv preprint / math.PR"
download: "https://arxiv.org/pdf/1407.7720v1"
pdf: "https://arxiv.org/pdf/1407.7720v1"
captured_at: "2026-05-09T12:57:37Z"
updated_at: "2026-05-09T12:57:37Z"
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

# Sample genealogy and mutational patterns for critical branching populations

- 著者: G. Achaz, C. Delaporte, A. Lambert
- 年: 2014
- 掲載情報: arXiv preprint / math.PR
- 情報源: [arxiv](https://arxiv.org/abs/1407.7720v1)
- ダウンロード: https://arxiv.org/pdf/1407.7720v1
- PDF: https://arxiv.org/pdf/1407.7720v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

We study a universal object for the genealogy of a sample in populations with mutations: the critical birth-death process with Poissonian mutations, conditioned on its population size at a fixed time horizon. We show how this process arises as the law of the genealogy of a sample in a large class of critical branching populations with mutations at birth, namely populations converging, in a large population asymptotic, towards the continuum random tree. We extend this model to populations with random foundation times, with (potentially improper) prior distributions g_i: x\mapsto x^{-i}, i\in\Z_+, including the so-called uniform (i=0) and log-uniform (i=1) priors. We first investigate the mutational patterns arising from these models, by studying the site frequency spectrum of a sample with fixed size, i.e. the number of mutations carried by k individuals in the sample. Explicit formulae for the expected frequency spectrum of a sample are provided, in the cases of a fixed foundation time, and of a uniform and log-uniform prior on the foundation time. Second, we establish the convergence in distribution, for large sample sizes, of the (suitably renormalized) tree spanned by the sample genealogy with prior g_i on the time of origin. We finally prove that the limiting genealogies with different priors can all be embedded in the same realization of a given Poisson point measure.

## PDF Text

SamplegenealogyandmutationalpatternsforcriticalbranchingpopulationsGuillaumeAchaz1;2;3CécileDelaporte3;4AmauryLambert3;4AbstractWestudyauniversalobjectforthegenealogyofasampleinpopulationswithmutations:thecriticalbirth-deathprocesswithPoissonianmutations,conditionedonitspopulationsizeataﬁxedtimehorizon.Weshowhowthisprocessarisesasthelawofthegenealogyofasampleinalargeclassofcriticalbranchingpopulationswithmutationsatbirth,namelypopulationsconverging,inalargepopulationasymptotic,towardsthecontinuumrandomtree.Weextendthismodeltopopulationswithrandomfoundationtimes,with(potentiallyimproper)priordistributionsgi:x7!x�i,i2Z+,includingtheso-calleduniform(i=0)andlog-uniform(i=1)priors.Weﬁrstinvestigatethemutationalpatternsarisingfromthesemodels,bystudyingthesitefrequencyspectrumofasamplewithﬁxedsize,i.e.thenumberofmutationscarriedbykindividualsinthesample.Explicitformulaefortheexpectedfrequencyspectrumofasampleareprovided,inthecasesofaﬁxedfoundationtime,andofauniformandlog-uniformprioronthefoundationtime.Second,weestablishtheconvergenceindistribution,forlargesamplesizes,ofthe(suitablyrenormalized)treespannedbythesamplegenealogywithpriorgionthetimeoforigin.WeﬁnallyprovethatthelimitinggenealogieswithdiﬀerentpriorscanallbeembeddedinthesamerealizationofagivenPoissonpointmeasure.Keywordsandphrases:criticalbirth-deathprocess;sampling;coalescentpointprocess;sitefrequencyspectrum;inﬁnite-sitemodel;Poissonpointmeasure;invarianceprinciple2010AMSClassiﬁcation:92D10,60J80(Primary),92D25,60F17,60G55,60G57,60J85(Sec-ondary)1IntroductionAmajorconcerninpopulationgeneticsisthepredictionofpatternsofgeneticvariationwithhelpofstochasticmodels.ThereferencemodelcurrentlyusedbybiologiststoanswerthisquestionistheKingmancoalescentmodel[16,15]coupledwithPoissonianmutationsonthelin-eages.Asthescalinglimitofnumerousconstantpopulationsizemodels,suchasWright-FisherandMoranmodels,itencompassesthetwopopulationmodelsthataremostcommonlyusedbybiologists.Thegenealogicalstructureofasample(ratherthanofthetotalpopulation)is
1UMR7138,EvolutionParis-Seine,UPMC&CNRS,Paris.2AtelierdeBioinformatique,UPMC,Paris.3UMR7241,CentreInterdisciplinairedeRechercheenBiologie,CollègedeFrance,Paris.4UMR7599,LaboratoiredeProbabilitésetModèlesAléatoires,UPMC&CNRS,Paris.Correspondingauthor:amaury.lambert@upmc.fr1
arXiv:1407.7720v1 [math.PR] 29 Jul 2014
well-known(equivalentlygivenbytheKingmancoalescent),andexplicitresultsontheallelicpartitiongeneratedbyrare,neutralmutations(equivalenttoaKingmancoalescentwithPoisso-nianmutations)areprovidedbyEwens’samplingformula[9,8].Inthiswork,weintendtostudythegenealogicalandmutationalpatternsofasamplefromabranchingpopulation,inordertooﬀeranalternativemodelwheretheconstantpopulationsizeassumptionisreleased,withnoaprioriassumptiononthevariationofthepopulationsizeovertime.Thesamplingishereessen-tialtomakethemodelapplicabletorealdataandcomparabletotheKingmancoalescentmodel.ThegenealogyofbranchingpopulationswasinparticularstudiedbyL.Popovicin[19],inthesettingofthecriticalbirth-deathprocessconditionedonitspopulationsizeataﬁxedtimehorizon,andlaterbyA.Lambertin[18]inthemoregeneralframeworkofsplittingtrees.Thegenealogyoftheextantindividualsisdescribedasarandompointprocess,calledcoalescentpointprocess,whichdistributionischaracterizedbyasequenceofi.i.d.randomvariables.Herewewanttofocusonthegenealogyofasampleratherthanofthetotalextantpopu-lation.Thequestionofsamplinginbirth-deathmodelshasalreadybeenapproachedwithtwodiﬀerentpointsofview.Ontheonehand,[19]and[22]dealwithBernoullisamplingofthetotalpopulation.Thisapproachratherappliestothespeciesscale,forexampleinthecaseofincompletephylogenies.Ontheotherhand,in[21]and[22],T.Stadlerconsidersthecaseofauniformsampleofmindividualsamongtheextantones,inthebirth-deathprocesscon-ditionedonitspopulationsizeatpresenttime,withuniformprioronitstimeoforigin.OurapproachisbasedonBernoullisamplingwithconditioningonthesamplesize,inordertogetauniformsamplewithﬁxedsizewithouthavingtoconditiononthetotalextantpopulationsize.WeﬁrstconsiderinSection1.1samplegenealogiesinageneralframeworkofbranchingpopu-lationswithneutralmutationsatbirth.Wemakeuseofconvergenceresultsobtainedbyoneoftheauthors[7]toshowhowabroadclassofsuchpopulationsallresultinthesamedistributionforthegenealogyofasample,namelythelawofacriticalbirth-deathmodelwithPoissonianmutationsonthelineages.WethenspecifyinSection1.2themodelthatweadoptfortherestofthepaper.WeﬁnallypresentinSection1.3theoutlineandthemainresultsofthiswork:inSection1.3.1,weinvestigatethelawofthegenealogyofasampleinthecriticalbirth-deathmodelconditionedonitssamplesize,withvariouspriordistributionsonthefoundationtimeofthepopulation.WeprovideinSection1.3.2explicitformulaefortheexpectedsitefrequencyspectrumofthesample.Section1.3.3isthendevotedtotheconvergenceindistributionofthesamplegenealogy,asthesamplesizegetslarge.Furthermore,westatethatthelimitinggenealogieswithdiﬀerentpriorscanallbeembeddedinthesamerealizationofagivenPoissonpointmeasure.1.1Genealogiesandsamplinginbranchingpopulationsconditionedonsur-vivalLetusﬁrstconsideraverygeneralmodelofbranchingpopulationswithmutations:let(TN)N2Nbeasequenceofsplittingtrees,i.e.randomtreeswhereindividualshavelifetimesthatdonotnecessarilyfollowanexponentialdistribution,duringwhichtheygivebirthatconstantratetoi.i.dcopiesofthemselves[10,11,18].ForanyN,TNischaracterizedbyitsso-calledlifespanmeasure�N,whichisa˙-ﬁnitemeasureon(0;1)suchthatR(1^r)�N(dr)<1.WefurtherassumethatanyindividualinTNexperiences,conditionalonherlifetimer,amuta-tionatbirthwithprobabilityfN(r),wherefNisacontinuousfunctionfromR�+to[0;1]calledmutationfunction.Weadopttheclassicalassumptionsofneutralmutations(i.e.mutations2
donotaﬀectthepopulationdynamics)andoftheinﬁnite-sitemodel[14]:eachindividualisassociatedtoaDNAsequence,andeachmutationoccursatasitethathasnevermutatedbefore.Finally,weﬁxt>0,andweconditionTNonsurvivalattimeNt.WeworklaterinatimescalewhereaunitoftimeisproportionaltoN:thefactorNcanthusbeseenasacounterpartoftheconstantpopulationsizeoftheWright-Fishermodel.WeassumethateachindividualaliveatNtisindependentlysampledwithprobabilitypN2(0;1).Individualsarelabeledaccordingtotheorderdeﬁnedin[17,Sec.1.1](“lefttoright”orderassociatedtotheplanarrepresentationofthetreewhendaughtersallsprouttotherightoftheirmother),andwedenotebyIN=(INj)jthesequenceofindicesofthesampledindividuals.SeeFigure1foragraphicalrepresentationofTN,andofsomeobjectshereafterdeﬁned.
Figure1:Inthethreepanels(a),(b),(c),theverticalaxisindicatestime.Thehorizontal(dotted)linesshowﬁliation.Mutationsaresymbolizedby?andsampledindividualsby�.(a)AnexampleoftherescaledtreeTNwith7extantindividualsatt,where4individualsaresampled.(b)its(marked)coalescentpointprocess(laterreferredtoas�N),(c)andthe(marked)coalescentpointprocessofthesampledindividuals.WearehereinterestedinthedistributionofthegenealogyofthesampledindividualsinTN,andweconsiderthemodelundertwoslightlydiﬀerentpointsofview:incase(I),relyingonresultsof[7],weconsiderascalinglimitinalargepopulationasymptotic,whileincase(II),weconsidertheexampleofthecriticalbirth-deathprocess,forwhichresultscanbeobtainedwithoutnecessarilyhavingtoconsiderN!1.Weshowherehowthesetwosettingsleadtothesamedistributionforthegenealogyofasample,justifyinghencethemodelwelaterconsiderfortherestofthepaper.TothisaimwerescaletimeinTNbymultipyingalltheedgelengthsofTNbyafactor1=N.ThisrescaledtreeisstilldenotedbyTN,andisnoworiginatingattimet.Thenweintroduce,foranyN2N,thesocalledmarkedcoalescentpointprocess�N[7],i.e.thetreespannedbythege-nealogyoftheextantpopulationofTNattimet,enrichedwiththemutationalhistoryofextantindividuals.Moreprecisely,�Nisapointmeasurethatcanbeexpressedas�N=PN�1i=1�(i;˙Ni)whereNisthenumberofextantindividualsattimet,andforany1�i�N�1,˙Niisitselfapointmeasure,whosesetofatomscontains,inadditiontothecoalescencetimebetweenin-dividualsiandi+1,allthetimesatwhichamutationoccurredonthei-thlineage(seeFigure1).3
(I)Scalinglimit.First,weassumethat(TN)converges,asN!1,towardsaBrowniantree(seee.g.[1]):foranyN2N,forany��0,deﬁne N(�):=���R(0;1)(1�e��r)�N(dr).Weassumethatthesequence(TN)follows(aparticularcaseof)AssumptionAin[7]:AssumptionA:Thereexistsasequenceofpositiverealnumbers(dN)N�1suchthatasN!1,thesequence(dN N(�=N))convergestowards�7!�2,i.e.theLaplaceexponentofaBrownianmotion.Thisassumptionhastobeinterpretedastheconvergenceinlawoftheso-calledjumpingchrono-logicalcontourprocessoftherescaledtreeTN,whichdistributionischaracterizedbyaLévyprocesswithﬁnitevariation,drift�1andLévymeasure�N[18].Second,weﬁx�2R+andwesupposethatthesequenceofmutationfunctions(fN)satisﬁesoneofthefollowingconvergenceassumptions[7]:AssumptionB.1:ForallN�1,forallu2R+,fN(u)=�N,where�N2[0;1]issuchthatdN
N�N�!N!1�.AssumptionB.2:Thesequence�u7!fN(Nu)
1^u�convergesuniformlytou7!f(u)
1^uonR�+,wherefisacontinuousfunctionfromR+toR+satisfyingf(u)=u!�asu!0+.Thenwehavethefollowingconvergence.Theorem.[7,Th.3.2]The(spacerescaled)pointmeasure�N=PN�1i=1�(idN
N;˙Ni)convergesindistribution,asN!1,towardsaPoissonpointprocesson[0;e]�(0;t)withintensitydlx�2dx,whereeisanindependentexponentialvariablewithparameter1=t,withindependentPoissonianmutationsatrate�onthelineages.Besides,weassumethatthesamplingprobabilityisgivenbypN=pN=dN,wherepisaﬁxedpositiverealnumbersuchthatpN2(0;1)forNlargeenough.Thentherescaledsequence(dN
NIN)ofindicesofthesampledindividuals(independentof(TN)),convergestowardsthesequenceofjumptimesofanindependentPoissonprocesswithratep.Thejointconvergenceof�NwithdN
NINisofcourseprovidedbytheirindependence.Asaconsequence,from[17]wededucethatthecoalescentpointprocessofthesampledindi-vidualsisthendistributedasthecoalescentpointprocessofacriticalbirth-deathmodelwithratepconditionedonsurvivalattimet,withindependentPoissonianmutationsatrate�onthelineages.(II)Criticalbirth-deathtree.Second,ﬁxN2N,p2(0;N),andconsidertheexamplewhereTNisacriticalbirth-deathtreewithrateNconditionedonsurvivalattimet.Then,setpN=p=NandassumethatthemutationfunctionfNisconstant,equalto�=N.Thisisinfactaparticularcaseof(I)(Assump-tionsAandB.1aresatisﬁedwithdN=N2),butherewedonotneedtoletN!1.ForanyN2N,themarkedcoalescentpointprocess�Nisdistributedasthecoalescentpointprocessofacriticalbirth-deathmodelwithrate1conditionedonsurvivalattimet,withPoissonianmutationsatrate�onthelineages(see[19,Sec.3]and[7,Ex.1]).Finally,from[17],wegetthatthecoalescentpointprocessofthesampleisthendistributed,exactlyasabove,asthecoalescentpointprocessofacriticalbirth-deathmodelwithratepconditionedonsurvivalattimet,withindependentPoissonianmutationsatratethetaonthelineages.4
Sincethetwocases(I)and(II)resultinthesamedistributionforthegenealogyofasam-ple,welimitourstudytocase(II).Besides,sincethemutationschemesariseasindependentofgenealogies,theresultsconcerningdistributionsofgenealogiesarestatedwithoutreferencetomutations.1.2ModelwithconditioningonthesamplesizeFromnowon,considerTacriticalbirth-deathtreewithrate1.Timeisnowcountedback-wardsintothepast,i.e.“presenttime”isnowtime0,and“uunitsoftimebeforepresent”isnowtimeu.Webeginwiththecaseofaﬁxedfoundationtimeofthepopulation.Themodelhasfourparameters:atimet2R�+,ascalingfactorN2R�+,apositiveintegern(thesamplesize),andasamplingparameterp2(0;N).AssumeﬁrstthatThasbeenfoundedNtunitsoftimeago.Aspreviously,individualsareindependentlysampledatpresenttime,withprobabilityp=N.Besides,werescaletimebyafactor1=N(alltheedgelengthsarethenmultipliedbyafactor1=N).WekeepthenotationTfortherescaledtree,sothatTisnowacriticalbirth-deathtreewithrateN,originatingattimet.
Figure2:Inbothﬁgures(a)and(b),theverticalaxisindicatestime(runningbackwards).(a)Agraphicalrepresentationofthecoalescentpointprocessatpresenttimeofa(rescaled)treeToriginatingattimetwith15extantindividualsandn=4sampledindividuals(symbolizedby�).Thehorizontallinesshowﬁliation.(b)Agraphicalrepresentationofthecoalescentpointprocessˇn=Pn�1k=1�(k n;H�k)ofthesamplerepresentedin(a).Wenowintroducetheconditioningonthesamplesize:weconditionTonhavingnsampledindividualsatpresenttime.Notethatafterconditioning,thedistributionofthensampledin-dividualswithinthetotalextantpopulationdoesnotdependonp,andisaposterioriequivalenttouniform,sequentialsampling.Thegenealogyofthensampledindividualsischaracterizedbyitscoalescentpointprocessˇn=n�1Xk=1�(k n;H�k);5
wherefor1�k�n�1,H�kisthedivergencetimebetweenthek-thandthe(k+1)-thsampledindividualintherescaledtreeT(seeFigure2).Thespacerescalingbyafactor1=nensuresinparticularthatthesupportsofthemeasuresˇnconvergeasn!1,whichisrequiredbytheresultslaterestablishedinthelargesamplesizeasymptotic.Besides,recallthatthankstotheirindependencewiththegenealogy,mutationsarefornowdeliberatelyomitted.Finally,wedeﬁne(Tn;k)1�k�n�1thedecreasingreorderingofthedivergencetimes(H�k)1�k�n�1.1.3OutlineandstatementofresultsTheﬁrstpurposeofthispaperistostudythedistributionofˇn,undervarioushypothesesontheoriginoftheprocess:wedenoteby-PtnthelawoftherescaledtreeTwithﬁxedtimeoforigintandsamplesizen,-P(1)nthelawofTwithinﬁnitetimeoforiginandsamplesizen,-P(i)nthelawofTwithrandomtimeoforigin,with(potentiallyimproper)priordistributiongi:x7!x�i,i2Z+,andsamplesizen.Notethatthecasei=0correspondstothecaseofauniformpriorinvestigatedin[2]and[12].Thisstudy,presentedinSection2,willthenenableustoderiveresultsconcerningmutationalpatternsofthesample(Section3),andthenconcerningthebehaviourofthegenealogyasthesamplesizegetslarge(Section4).1.3.1AuniversallawforthegenealogyofasampleFirst,inthecaseofaﬁxedtimeoforigin,thelawofˇnunderPtnisindependentofNandisspeciﬁedbythefollowingresult(Theorem2.1):Theorem.UnderPtn,(H�i)1�i�n�1isasequenceofi.i.d.randomvariableswithprobabilitydensityfunctionx7!p
(1+px)21+pt pt1(0;t)(x).Inotherwords,thecoalescentpointprocessˇnhasthelawofthegenealogyofacriticalbirth-deathtreewithratepconditionedonhavingnextantindividualsattimet.Wethenprovethatthisequalityinlawstillholdswhenlettingthetimegotoinﬁnityorwhenrandomizingthetimeoforigin(withpriordistributiongi,i2Z+)inbothprocesses:forexample,underP(i)nthecoalescentpointprocessˇnhasthelawofthegenealogyofacriticalbirth-deathtreewithratep,withpriorgionitstimeoforigin,andconditionedonhavingnextantindividualsatpresenttime.Hencewhatevertheassumptiononthefoundationtimeofthepopulation,thestudyofthegenealogyofthesampleboilsdowntothesameobject:thegenealogyofacriticalbirth-deathprocesswithratep,withextantpopulationsizen.Followingonfromresultsprovidedby[12]inthecaseofauniformprior,wethenobtainthefollowingpropertyforthesuccessivedivergencetimes(Tn;k)1�k�n�1(Proposition2.9):Proposition.UnderP(i)n,thetimeTn;ktothek-thmostrecentcommonancestorhasﬁnitemomentofordermiﬀm�k+i.Althoughwelimitedhereourstudytotheframework(II)introducedearlier,onecouldcertainlygeneralizetheseresults(andtheupcomingones)tothescalinglimitofcase(I).Toprovethis,onewouldhavetoconsiderasequenceoftreesconditionedontheirsamplesize,andthentoestablishtheconvergence,inthelargepopulationasymptotic,ofthemarkedcoalescentpointprocessofthesample.Thisishoweverbeyondthescopeofthepresentpaper.6
1.3.2MutationalpatternsInSection3,westudytheso-calledsitefrequencyspectrumofthesample,i.e.the(n�1)-tuple(˘1;:::;˘n�1),where˘kisthenumberofmutationscarriedbykindividualsinthesample.Variousresultsforthefrequencyspectrumintheframeworkofgeneralbranchingprocessesareestablishedin[17,3,4,20].Oneoftheauthorsinvestigatesin[17]thecaseofcoalescentpointprocesseswithPoissonianmutationsongermlines,andobtainsasymptoticresultsforthesiteandallelefrequencyspectrumoflargesamples.Explicitformulaefortheexpectedallelefre-quencyspectrumofasplittingtreewithnindividualsatﬁxedtimehorizontareprovidedbyN.ChampagnatandthisauthorinthecaseofPoissonianmutationsonthelineages[3],andbyM.Richardinthecaseofmutationsatbirth[20].Theirresultsarecomparedin[5]intheparticularcaseofbirth-deathprocesses.Furtherresultsabouttheasymptoticbehaviour,ast!1,oflarge(resp.old)families,i.e.familieswithmostfrequent(resp.oldest)types,aredevelopedin[4].Inthisarticle,wegetexplicitformulaefortheexpectedsitefrequencyspectrum(˘k)1�k�n�1ofthesampleunderPtn,P(1)n,P(0)nandP(1)n.AccordingtoSection1.1,mutationsareassumedtooccuratconstantrate�onthelineages.Twodiﬀerentmethodsareusedtoobtaintheexpectationofthe˘k.Ontheonehand,thesimilarityofthemodelwith[17]allowsustomakeuseofaproofmethoddevelopedinthisarticle.Indeed,accordingtotheresultsofSection2,theframeworkusedin[17]coversoursettinginthecaseofaninﬁnitetimeoforigin.Ontheotherhand,foreachk,E(˘k)canbeexpressedasalinearcombinationoftheexpectationsofbranchingtimes[23].Althoughtheﬁrstmethodcouldbeusedtoprovealltheresultsofthissection,thesecondoneprovidesveryshortproofsinthecasesofaninﬁnitetimeoforiginandofauniformprior.FirstunderP(1)n,theabsenceofaﬁrstmomentforthetimetothemostrecentcommonancestoryieldsimmediatelythefollowingresult(Proposition3.2).Proposition.Foranyk2f1;:::;n�1g,E(1)n(˘k)isinﬁnite.Second,usingthefactthattheexpecteddivergencetimes,undertheKingmancoalescentmodel,andunderthe(suitablyrescaled)criticalbirth-deathprocesswithuniformprioronitstimeoforigin,areequal[12],wededucethattheexpectedsitefrequencyspectrumunderP(0)nisthatofasampleoftheKingmancoalescent[23,(4.20)](Proposition3.4).Proposition.Foranyk2f1;:::;n�1g,E(0)n(˘k)=n�=kp.Finally,theformulasobtainedintheremainingtwocasesarethefollowing(Propositions3.1and3.5).Proposition.Foranyk2f1;:::;n�1g,t2R�+,deﬁning˝:=pt,wehaveEtn(˘k)=�
p(n�3k�1
k+(n�k�1)(k+1)
k˝+(1+˝)k�1
˝k+1h2˝2�(n�2k�1)2˝�(n�k�1)(k+1)i�ln(1+˝)�k�1Xi=11
i�˝
1+˝�i�):Proposition.Foranyk2f1;:::;n�3g,E(1)n(˘k)=�
pn(n�1)
(n�k)(n�k�2)"n+k�2
k�2(n�1)
n�k�1(Hn�1�Hk)#;whereforanyk2N,Hk=Pkj=1j�1.7
1.3.3ConvergenceofgenealogiesforlargesamplesizesWeinvestigateinSection4theasymptoticbehaviourofthecoalescentpointprocessˇn,asn!1.Wetakeinspirationfromasymptoticresultspresentedin[19]and[2].First,L.Popovicobtainsin[19]theconvergenceofthe(suitablyrescaled)coalescentpointprocessofacriticalbirth-deathprocessconditionedonitspopulationsizeattimettowardsacertainPoissonpointmeasureon(0;1)�(0;t).Usingthisresult,shethenobtainswithD.Aldousin[2]asimilarconvergenceforthemodelwithuniformprioronthetimeoforigin.Hereweextendthistothecasesofaninﬁnitetimeoforigin,andofarandomtimeoforiginwithpriorgi,i2N.Obtainingsuchasymptoticresultsrequirestoletthesamplingparameterpdependonninsuchawaythatp=n=�,with�>0.Itensuresindeedthattheexpectednumberofsampledindividualsisoftheorderofthesamplesizen.Wethenobtainthefollowingconvergences(Theorem4.1).Theorem.DenotebyˇtthePoissonpointmeasurewithintensity�dlx�2dx1(l;x)2(0;1)�(0;�t).a)UnderP(1)n,thecoalescentpointprocessˇnconvergesinlaw,asn!1,towardsthePoissonpointmeasureˇwithintensitymeasure�dlx�2dxon(0;1)�R�+.b)Foranyi2Z+,underP(i)n,thejointlawofthetimeoforigin,alongwithˇn,convergesasn!1towardsapair(T(i)or;ˇ(i)),suchthatT(i)orfollowsaninverse-gammadistributionwithparameters(i+1;�),andconditionalonT(i)or=t,ˇ(i)isdistributedasˇt.Thelastresultweobtaindescribesthelinksbetweenthediﬀerentrandommeasuresobtainedinthelimit.Letusordertheatomsofourpointprocessesw.r.t.theirsecondcoordinate.WeprovethattherandomvariableT(i)orisdistributedasthe(i+1)-thlargestatomofthePoissonpointprocessˇ,andwethendeducethefollowingtheorem(Theorem4.4).Theorem.Thepointmeasureˇ(i)isdistributedasthepointprocessobtainedfromˇbyremovingitsi+1largestatoms.Inotherwords,genealogieswithdiﬀerentpriorscanallbeembeddedinthesamerealizationofthepointmeasureˇ.2AuniversaldistributionforthegenealogyofasampleLetusconsiderthemodeldeﬁnedinSection1.2andspecifysomenotation.RecallthattherescaledtreeTisacriticalbirth-deathtreewithparameterNoriginatingattimet,andthateachextantindividualinTisindependentlysampledwithprobabilityp=N.WedenotebyNthenumberofextantindividualsatpresenttimeinT,andwelabeltheseindi-vidualsfrom1toN,usingtheorderdeﬁnedin[17,Sec.1.1].Inordertoformalizethesamplingprocess,weintroduceasequence(Ij)j�1ofrandomvariables,suchthat(I1;I2�I1;I3�I2;:::)formsasequenceofi.i.d.geometricrandomvariableswithsuccessprobabilityp=N.ThenforanyjsuchthatIj�N,Ijisthelabelofthej-thsampledindividualintheextantpopulationatpresenttime(inthepreviouslydeﬁnedorder).TheconditioningonthesamplesizetobeequaltonmeansthusconditioningonfIn�N<In+1g.8
Figure3:(a)Thecoalescentpointprocessatpresenttimeofa(rescaled)populationoriginatingattimetwithnsampledindividuals(symbolizedby�).TheNverticalbranchesrepresentthesequence(Hi)1�i�N.(b)Thecoalescentpointprocessˇnofthesamplerepresentedinﬁgure(a).TheequalityH�2=maxfHI2+1;:::;HI3gisillustratedbyboldlines.Letusnowexplainthelinkbetweenthegenealogyofthetotalextantpopulationandthegenealogyofthesample.Denoteby(Hi)1�i�N�1thesequenceofnodedepthsofthecoalescentpointprocessofthetotalextantpopulation,i.e.forany1�i�N�1,Hiisthedivergencetimebetweenindividualiandindividuali+1intherescaledtreeT.Weknowfrom[18,Th.5.4]thatforany1�i<j�N,thedivergencetimebetweenindividualiandjisgivenbythemaximumofthenodedepthsfHi+1;:::;Hjg.Asaconsequence,thedivergencetimeH�ibetweenindividualIiandindividualIi+1inT,1�i�n�1,isgivenbyH�i=maxfHIi+1;:::;HIi+1g:Finallywerecallthedeﬁnitionofthepointmeasureˇn:ˇn=n�1Xk=1�(k n;H�k):Inthesequelweequallycall“coalescentpointprocess”ofthesample,themeasureˇnandthesequence(H�i)1�i�n�1.SeeFigure3foragraphicalrepresentationoftheobjectsdeﬁnedabove.Theaimofthissectionistocharacterizethelawofthegenealogyofthesample,underdiﬀerentassumptionsonthetimeoforigin.Section2.1establishesthedistributionofˇninthecaseofaﬁxed(possiblyinﬁnite)timeoforigin.InSection2.2,werandomizethetimeoforiginbygivingitapriordistributionoftheformx7!x�i,i2Z+.2.1FixedtimeoforiginWedenotebyPtthelawoftherescaledtreeToriginatingattimet,andwerecallthatPtndenotesthelawofToriginatingattimetandconditionedonhavingnsampledindividualsatpresenttime,i.eonfIn�N<In+1g.ThefollowingtheoremspeciﬁesthelawofthesamplegenealogyunderPtn.9
Theorem2.1.UnderPtn,thecoalescentpointprocess(H�i)1�i�n�1isasequenceofi.i.d.ran-domvariableswithprobabilitydensityfunctionx7!p
(1+px)21+pt pt1(0;t)(x):Remark2.2.Accordingto[19,Lem.3],therescaledcoalescentpointprocessofthensampledindividualsisthusdistributedasthecoalescentpointprocessofthepopulationattimetofacriticalbranchingprocesswithratep,conditionedonhavingnextantindividualsattimet–orequivalently,asthecoalescentpointprocessofthepopulationattimeptofacriticalbranchingprocesswithrate1,conditionedonhavingnextantindividualsattimept,andthenrescaledbyafactor1=p.Remark2.3.Itisinterestingtonotethattheindependencew.r.t.NofthelawofˇnunderPtnimpliesthattheparameterNhasonlyascalingeﬀectonthelawofthegenealogy.Onthecontrary,theparameterspandtbothaﬀectthebranchlengthsratios,throughtheconditioningonthepopulationsizeataﬁxedtime.Weextendthetheoremtothelimitingcaset!1:recallthatP(1)n(T2�)=limt!1Ptn(T2�).Wehavethefollowingstatement.Proposition2.4.UnderP(1)n,(H�i)1�i�n�1isasequenceofi.i.d.randomvariableswithdensityfunctionx7!p
(1+px)21R+(x).Recallthatforany1�k�n�1,Tn;kisdeﬁnedasthek-thorderstatisticofthesequence(H�i)1�i�n�1.Inparticular,Tn;1isthetimetothemostrecentcommonancestorofthesample.Thefollowingpropositionprovidesthem-thmomentofTn;kunderP(1)n.Proposition2.5.Forany1�k�n�1andm�1,them-thmomentofTn;kunderP(1)nisﬁniteiﬀm�k�1.Speciﬁcally,form�k�1,E(1)n((Tn;k)m)=�n�k+m�1m�
pm�k�1m�:Inparticular,thetimetothemostrecentcommonancestorhasinﬁniteexpectationunderP(1)n.ProofofProposition2.5:UsingthedeﬁnitionofTn;kasthek-thorderstatisticofthei.i.d.randomvariables(H�i)1�i�n�1withdensityfunctionx7!p
(1+px)21R+(x),alongwith[6,2.1.6],wegetthatthedensityfunctionofTn;kunderP(1)niss7!p(n�k)�n�1n�k�(ps)n�k�1
(1+ps)n1s�0.ThenE(1)n((Tn;k)m)=p�m(n�k)�n�1n�k�Z10sn+m�k�1
(1+s)nds:WeconcludeusingPropositionA.2intheAppendix.�ProofofTheorem2.1:Forany(t1;:::;tn�1)2(R+)n�1,writePt(H�1�t1;:::;H�n�1�tn�1;In�N�In+1jN�1)=Xk0;:::;kn�1Pt(H�1�t1;:::;H�n�1�tn�1;In�N�In+1;I1=k0;:::;In=k0+:::+knjN�1):10
Nowrecallfrom[18,Th.5.4]thatconditionalonN�1,thesequence(Hi)1�i�N�1isdistributedasasequenceofi.i.d.randomvariablessatisfyingPt(Hi�u)=Nu
1+Nu,stoppedattheﬁrstoneexceedingt.RememberingthatH�i=maxfHIi+1;:::;HIi+1g,andfromthedeﬁnitionofthesequence(Ii)i�1,Pt(H�1�t1;:::;H�n�1�tn�1;In�N�In+1jN�1)=Xk0;:::;kn�1 nYi=0p
N�1�p
N�ki�1�Pt(max1�i<l0Hi�t;maxl0�i<l1Hi�t1;:::;maxln�2�i<ln�1Hi�tn�1;maxln�1�i<lnHi>t)!=Xk0;:::;kn�1"nYi=0p
N�1�p
N�ki�1#�Nt
1+Nt�k0�1"1��Nt
1+Nt�kn#n�1Yi=1�N(ti^t)
1+N(ti^t)�ki;whereforany0�i�n,li:=k0+:::+ki.Now8u2R+,Xk�1p
N�1�p
N�k�1�Nu
1+Nu�k=pu
1+pu;andXkn�1p
N�1�p
N�kn�1 1��Nu
1+Nu�kn!=1
1+pu:ThuswehavePt(H�1�t1;:::;H�n�1�tn�1;In�N�In+1jN�1)=1+Nt
Nt1
1+ptpt
1+ptn�1Yi=1p(ti^t)
1+p(ti^t):(1)Finally,bytakingti=tforall1�i�n�1in(1),wegetPt(In�N�In+1jN�1)=1+Nt
Nt1
1+pt�pt
1+pt�n:(2)Asaconsequence,wehaveforany(t1;:::;tn�1)2(R+)n�1,Ptn(H�1�t1;:::;H�n�1�tn�1)=�1+pt pt�n�1n�1Yi=1p(ti^t)
1+p(ti^t);whichleadstotheannouncedresult.�2.2RandomtimeoforiginWenowwanttorandomizethetimeoforigin.Tothisaim,wegivea(potentiallyimproper)priordistributiontothetimeoforigininthemodeldeﬁnedabove.Weinvestigateherepriorswithdensityfunctiongi:u7!u�i1R�+(u),i2Z+.Thecasei=0(resp.i=1)isusually11
referredtoasuniform(resp.log-uniform)prioron(0;1).Forany0�i<n,recallthatP(i)ndenotesthelawoftherescaledtreeT,withpriorgionitstimeoforigin,andconditionedonhavingnsampledindividualsatpresenttime:P(i)n(T2�)=R+10Ptn(T2�)Pt(In�N�In+1)gi(t)dt
R+10Pt(In�N�In+1)gi(t)dt:NotethatwewouldhaveobtainedthesamedistributionP(i)nifwehadrandomizedthetimeoforiginbeforehavingrescaledtimeintheprocess.Proposition2.6.Forany0�i<n,thelawofTunderP(i)nisgivenbyP(i)n(T2�)=Z+10Ptn(T2�)h(i)n(t)dt;whereh(i)n:t7!pn�n�1i�(pt)n�i�1
(1+pt)n+11R+(t);i.e.,thetimeoforiginofTunderP(i)nisarandomvariableTorwithposteriordistributioncharacterizedbyitsprobabilitydensityfunctionh(i)n.ProofofProposition2.6:From(2)andfromPt(N�1)=(1+Nt)�1,weknowthatforallt>0,Pt(In�N�In+1)=1
Nt(pt)n
(1+pt)n+1.Thus,Z+10Pt(In�N�In+1)gi(t)dt=pi
NZ+10(pt)n�i�1
(1+pt)n+1pdt=pi
N1
(i+1)�ni+1�=pi nN�n�1i��1;usingPropositionA.2intheAppendix.FinallybydeﬁnitionofP(i)n,P(i)n(T2�)=Nn pi�n�1i�Z+10Ptn(T2�)p
N(pt)n�1
(1+pt)n+1dt ti=Z+10Ptn(T2�)pn�n�1i�(pt)n�i�1
(1+pt)n+1dt;whichgivestheexpectedresult.�Asacorollary,wehavethatthegenealogyofthesamplehasthelawofthegenealogyofabirth-deathprocesswithﬁxedsize:Corollary2.7.Foranyi2Z+,therescaledcoalescentpointprocessˇnisdistributedunderP(i)nasthecoalescentpointprocessofacriticalbirth-deathprocesswithparameterp,withpriorgionitstimeoforigin,andconditionedonhavingnextantindividualsatpresenttime.Remark2.8.FromthecorollaryitiseasytoseethatthesamplingparameterponlyhasascalingeﬀectontimeregardingthedistributionofˇnunderP(i)n.ThisremainstrueunderP(1)n,butnotunderPtnbecauseoftheconditioningonthepopulationsizeattimet(seeRemark2.3).12
ProofofCorollary2.7:Theprobabilityforacriticalbirth-deathprocesswithparameterpofhavingnextantindividualsattimetis(pt)n�1
(1+pt)n+1(see[2,(1)]),henceitdiﬀersfromPt(In�N�In+1)onlybyafactorp=N,andaneasyadaptationofthecalculationsintheproofofProposition2.6givestheexpectedresult.�Finallywestudythemomentsofthedivergencetimes(Tn;k)1�k�n�1.Thefollowingpropositionstatesanecessaryandsuﬃcientconditionfortheexistenceofthem-thmomentofTn;kunderP(i)n.Inthecaseofauniformprior(i=0),wealsorecalltheexplicitformulaestablishedin[12,Cor.2.2].Proposition2.9.Forany0�i<n,1�k�n�1andm�1,them-thmomentofTn;kunderP(i)nisﬁniteiﬀm�k+i.Besides,forany1�k�n�1andm�k,E(0)n((Tn;k)m)=�n�k+m�1m�
pm�km�:Proof:FromTheorem2.1,weknowthatunderPtn,therandomvariables(H�i)1�i�n�1arei.i.d.Henceweobtainfrom[6,2.1.6]thattherandomvariableTn;k,deﬁnedasthek-thorderstatisticofthesequence(H�i)1�i�n�1,hasdensityfunctionftn;k:s7!p(n�k)�n�1n�k�(ps)n�k�1
(1+ps)n(1+pt)n�k
(pt)n�1(pt�ps)k�11s�tunderPtn.Asaconsequence,wehaveE(i)n((Tn;k)m)=Z10sm�Z10ftn;k(s)h(i)n(t)dt�ds;andthenE(i)n((Tn;k)m)<1,Z10(ps)n�k�1+m
(1+ps)n�Z1ps(pt�ps)k�1
(pt)i(1+pt)k+1dt�ds<1,Z10sn�k�1+m
(1+s)n�Z1s(t�s)k�1
ti(1+t)k+1dt�ds<1:LetusﬁrstcharacterizetheintegrabilityofthefunctionF:s7!sn�k�1+m
(1+s)n�R1s(t�s)k�1
ti(1+t)k+1dt�intheneighbourhoodof+1.WeproveherethatR1s(t�s)k�1
ti(1+t)k+1dt˘s!+1cs�i�1,wherecisa(positive)constantw.r.t.s.Expanding(t�s)k�1,wehaveZ1s(t�s)k�1
ti(1+t)k+1dt=k�1Xj=0(�s)k�1�jZ1sdt ti�j(1+t)k+1:Notingthatforany0�j�k�1,1
(k+i�j)(1+s)k+i�j�Z1sdt ti�j(1+t)k+1�1
(k+i�j)sk+i�j;13
weobtaink�1Xj=0(�1)k�1�j k+i�j�k�1j�sk+i�j
(1+s)k+i�j�si+1Z1s(t�s)k�1
ti(1+t)k+1dt�k�1Xj=0(�1)k�1�j k+i�j�k�1j�;Lettings!1leadstotheannouncedequivalent.Asaconsequence,F(s)˘+1csm�k�i�2,andFisintegrableintheneighbourhoodof+1iﬀm�k�i�0.Ontheotherhand,inthecasem�k�i�0,theintegrabilityofFonanycompactsetofR+isclear.ThusE(i)n((Tn;k)m)isﬁniteiﬀm�k�i�0.�3Expectedfrequencyspectrum3.1MutationsettingRecallfromSection1thatweassumePoissonianmutationsatrate�2R+onthelineages.Weadoptthenotationintroducedin[17],whoseframeworkisveryclosetoours.Let(Pj)j2f0;:::;n�1gbeindependentPoissonmeasuresonR�+withparameter�.ForeachjwedenotetheatomlocationsofPjby`j1<`j2<:::.Thebranchlengths(H�0;H�1;:::;H�n�1),wherewesetH�0:=Tor,characterizethegenealogyofthenindividuals(labeledaccordinglyfrom0ton�1)jointlywiththefoundationtimeofthepopulation.Thenthetimes`jlsatisfying`jl<H�jareinterpretedasmutationevents,andforallk2f0;:::;n�j�1g,individualj+kbearsmutation`jlifmaxfH�j+1;:::;H�j+kg<`jl<H�j;wheremax?=0(seeFigure4).Theﬁrstinequalityexpressesthefactthatamutationonbranchjinthecoalescentpointprocessiscarriedbyindividualj+kifthetimeatwhichitappearsisgreaterthanthedivergencetimeofindividualsjandj+k(recallthattimeisrunningbackwards).Thesecondinequalitymeansthatallthevalues`jlthataregreaterthanthej-thnodedepthH�jarenottakenintoaccount.Foranyk2f1;:::;n�1g,recallthatwedenoteby(˘k)1�k�n�1thesitefrequencyspectrumofthesample,i.e.˘kisthenumberofmutationscarriedbykindividualsamongthensampledindividuals.ThesumS=Pn�1k=1˘kistheso-callednumberofpolymorphicsites,alsoknownassinglenucleotidepolymorphismsinpopulationgenomics.3.2ResultsInthissectionwegiveexplicitformulaefortheexpectedsitefrequencyspectruminthecaseofaﬁxedtimeoforiginandinthecaseofauniformorlog-uniformprioronthetimeoforigin.Theproofsarebasedontwodiﬀerentmethods,dependingontheassumptiononTor,andareexpandedinthenextsection.3.2.1Fixed(ﬁnite)timeoforiginTheexpectedsitefrequencyspectrumofthensampledindividualsunderPtnisgivenby14
Figure4:Thecoalescentpointprocessofasampleofsizen,withmutationssymbolizedbystars.Mutations`j2and`j3arecarriedbyindividualj+kwhilemutation`j1isnot.Since`j4>H�j,itisnotconsideredasamutationevent.Onlymutations`01,`02and`j1arecarriedbytwoindividuals,sothathere˘2=3.Proposition3.1.Foranyk2f1;:::;n�1g,t2R�+,deﬁning˝:=pt,wehaveEtn(˘k)=�
p(n�3k�1
k+(n�k�1)(k+1)
k˝+(1+˝)k�1
˝k+1h2˝2�(n�2k�1)2˝�(n�k�1)(k+1)i�ln(1+˝)�k�1Xi=11
i�˝
1+˝�i�):3.2.2InﬁnitetimeoforiginThefollowingtwopropositionsaredirectconsequencesofProposition3.1.HowevernotethatProposition3.2canbeprovedindependentlyfromtheformulaprovidedbyProposition3.1,aswillbeexplainedinSection3.3.Proposition3.2.Foranyk2f1;:::;n�1g,˘khasinﬁniteexpectationunderP(1)n.Theinﬁniteexpectationof˘kunderP(1)nleadstoconsideritsrenormalizationbytheexpectednumberofpolymorphicsites.Thepropositionbelowshowsthatlettingthetimeoforigingoto+1ﬂattenstherenormalizedexpectedfrequencyspectrum.AhintforthisresultisgiveninSection3.3,whileweproveitherebylettingt!1inProposition3.1.Proposition3.3.Foranyk2f1;:::;n�1g,limt!1Etn(˘k)
Etn(S)=1
n�1:Proof:Fixk2f1;:::;n�1g.OnecaneasilyseefromProposition3.1thatast!+1,Etn(˘k)˘2�ln(t)andEtn(S)˘2�(n�1)ln(t);15
Figure5:Thenormalizedexpectedsitefrequencyspectrumofasampleofn=10individuals,underPtn,for˝=pt2f1;10;100;1000g.Thehorizontaldottedlinehasequationy=1=(n�1)whichleadstotheresult.�3.2.3RandomtimeoforiginWeprovideexplicitformulaefortheexpectedfrequencyspectrumfortwoparticularcasesofpriors:theuniformprior(casei=0)andthelog-uniformprior(casei=1).Proposition3.4.Foranyk2f1;:::;n�1g,E(0)n(˘k)=n�=pk.Proposition3.5.Foranyk2f1;:::;n�3g,E(1)n(˘k)=�
pn(n�1)
(n�k)(n�k�2)"n+k�2
k�2(n�1)
n�k�1(Hn�1�Hk)#;whereforanyk2N,Hk=Pkj=1j�1.Remark3.6.TheformulaeobtainedforE(1)n(˘n�2)andE(1)n(˘n�3),whichwechosenottodisplayhere,involvenonexplicitintegrals.GraphicalrepresentationsoftheexpectedfrequencyspectrumunderPtn,P(0)n,P(1)nareprovidedinFigures5and6.16
Figure6:Thenormalizedexpectedsitefrequencyspectrumofasampleofn=10individuals,underP(0)nandP(1)n.3.3ProofsDependingontheassumptiononTor,twodiﬀerentmethodscanbeused.Theﬁrstonereliesonanexpressionoftheexpectednumberofmutationscarriedbykindividualsasafunctionoftheexpectedcoalescencetimesofthetree[23,pp.103-105].Thesecondonedecomposesitscomputationintothesumofthemutationspresentonlineagej,1�j�n�k,carriedbykindividuals[17].AlthoughthesecondonecouldbeusedtoprovealltheresultsofSection3.2,theﬁrstoneprovidesaveryshortproofinthecasesofaninﬁnitetimeoforiginandofauniformprior.3.3.1InﬁnitetimeoforiginanduniformpriorWebaseourproofofPropositions3.2and3.4onFormula[23,(4.22)],whichgivesforany1�k�n�1andanyi2Z+[f1gE(i)n(˘k)=�2
k�n�1k��1n�k+1Xj=2�j2��n�jk�1�E(i)n(^Tn;j);(3)where^Tn;j:=Tn;j�Tn;j�1denotesthetimeelapsedbetweenthe(j�1)-thandthej-thcoales-cence.Whenthetimeoforiginissettobeinﬁnitea.s.,fromProposition2.5theexpectedtimetothemostrecentcommonancestorisinﬁnite,whichentailsdirectly,alongwith(3),thatE(1)n(˘k)isinﬁniteforanyk2f1;:::;n�1g.FromEquation(3)wecanalsogiveanintuitiveexplanationoftheresultofProposition3.3,whichestablishesthatlimt!1Etn(˘k)=Etn(S)=1=(n�1)forany1�k�n�1.Indeed,using(3)tocomputeE(1)n(˘k),fromProposition2.5weknowthat17
E(1)n(^Tn;2)istheonlyinﬁnitecontributiontoE(1)n(˘k).ThiscontributionisthussupportedbytheﬁrstorderstatisticTn;1of(H�i)1�i�n�1(i.e.thelargestdivergencetimeinthecoalescentpointprocess).ConditionalonTn;1=H�i0,˘iisﬁnitea.s.foranyi6=n�i0.NowunderP(1)n,(H�i)1�i�n�1isasequenceofi.i.d.randomvariables,sothattheindexi0isuniformlydistributedinf1;:::;n�1g.Thisexplainstheindependenceoflimt!1Etn(˘k)=Etn(S)w.r.t.k.Inthecaseofauniformprioronthetimeoforigin,weuseacomparisonwiththeverydocumentedKingmancoalescentmodel.DenotebyPKthelawofthegenealogyofasampleofsizenundertheKingmancoalescentmodelwithmutationsatrate�.Firstfrom[12]weknowthatforanyj2f2;:::;ng,theinter-coalescencetime^Tn;jhaveproportionalexpectationunderP(0)nandundertheKingmancoalescentmodel:E(0)n(^Tn;j)=n
2pEK(^Tn;j).Second,from[23,(4.20)],foranyk2f1;:::;n�1g,EK(˘k)=2�
k.Asaconsequence,using(3)(whichisalsovalidunderPK)weobtainforany1�k�n�1,E(0)n(˘k)=n p�
k.ThisendstheproofofProposition3.4.3.3.2Fixed(ﬁnite)timeoforiginandlog-uniformpriorWhenTorisﬁxed(andﬁnite),orinthecaseofanonuniformprioronTor(i2N),theequality(3)doesnotleadtoanexplicitexpressionoftheexpectedfrequencyspectrum.TheformulaestatedinProposition3.1(caseTor=t2R�+)andProposition3.5(caseofalog-uniformprioronTor)areobtainedusingamethoddevelopedin[17](seeproofofTheorem2.3formoredetails).ProofofProposition3.1:Fixt>0.Decomposing˘kintothesumofthenumberofmutationsonthej-thbranchcarriedbyexactlykindividuals,from[17](seeproofofTheorem2.3),weknowthatEtn(˘k)=�n�kXj=0Etn��minfH�j;H�j+kg�maxfH�j+1;:::;H�j+k�1g�+�;(4)wherewehavesetH�n:=+1.Twoparticularcasesappear,namelyj=0,whereminfH�j;H�j+kg=H�ka.s.,andj+k=n,whereminfH�j;H�j+kg=H�ja.s.Henceusingthei.i.d.propertyof(H�j)1�j�n�1,itfollowsforany1�j�n�k�1,Q:=Etn��minfH�j;H�j+kg�maxfH�j+1;:::;H�j+k�1g�+�=EtnZ101maxfH�j+1;:::;H�j+k�1g<x<minfH�j;H�j+kgdx=Z10Ptn(H�1>x)2Ptn(H�1<x)k�1dx;andsimilarlyforj2f0;n�kg,R:=Etn��minfH�j;H�j+kg�maxfH�j+1;:::;H�j+k�1g�+�=Z10Ptn(H�1>x)Ptn(H�1<x)k�1dx:FromTheorem2.1,weknowthatPtn(H�1<x)=p(x^t)
1+p(x^t)1+pt pt.Thisentails,afterachangeofvariables,andrecallingthatwedeﬁned˝=pt,18
Q=1
p�1+˝
˝�k�1Z˝0�x
1+x�k�1�1�x(1+˝)
˝(1+x)�2dx=1
p(1+˝)k�1
˝k+1�˝2Ik+1;2(˝)�2˝Ik+1;1(˝)+Ik+1;0(˝)�;andR=1
p�1+˝
˝�k�1Z˝0�x
1+x�k�1�1�x(1+˝)
˝(1+x)�dx=1
p(1+˝)k�1
˝k+1�˝2Ik;1(˝)�˝Ik;0(˝)�;whereforanyu2R�+,k2Z+,l2Z,Ik;l(u):=Ru0xk�l
(1+x)kdx.UsingEquation(4),thisleadstoEtn(˘k)=�
p(1+˝)k�1
˝k+1h(n�k�1)�˝2Ik+1;2(˝)�2˝Ik+1;1(˝)+Ik+1;0(˝)�+2�˝2Ik;1(˝)�˝Ik;0(˝)�iFinally,usingtheformulaeprovidedbyPropositionA.1forIk;l,l2f0;1;2g,weﬁnallygetaftersomerearrangementsEtn(˘k)=�
p(1+˝)k�1
˝k+1�ln(1+˝)�2˝2�2(n�2k�1)˝�(k+1)(n�k�1))��2˝2+˝(n�k�1)+n�k�1
k˝k+2
(1+˝)k+(�1)k�1n�k�1
k�1�1
(1+˝)k�(2˝+1)+k�1Xj=1�k�1j�(�1)j j�1�1
(1+˝)j��2˝2+2˝k j+1�(n�k�1)�2˝+k+1
j+1�k k�j��:(5)Toobtaintheﬁnalform,wedecomposethesuminther.h.s.asfollows:Firstdeﬁne,forx2Randk2N˚1;k(x):=Pkj=1�kj�xj j;˚2;k(x):=Pkj=1�kj�xj+1
j(j+1) 1;k(x):=Pkj=1�kj�xj j(k�j); 2;k(x):=Pkj=1�kj�xj+1
j(j+1)(k�j):ThenwehaveS:=k�1Xj=1�k�1j�(�1)j j�1�1
(1+˝)j��2˝2+2˝k j+1�(n�k�1)�2˝+k+1
j+1�k k�j�=2˝2�˚1;k�1(�1)�˚1;k�1(�(1+˝)�1)��2˝k�˚2;k�1(�1)�(1+˝)˚2;k�1(�(1+˝)�1)��(n�k�1)2˝k� 1;k�1(�1)� 1;k�1(�(1+˝)�1)�+(n�k�1)(k+1)k� 2;k�1(�1)�(1+˝) 2;k�1(�(1+˝)�1)�:19
Letusnowreexpressthefunctions˚1;k,˚2;k, 1;kand 2;k.Fixx2Randk2N.Thefunction˚1;kisdiﬀerentiableatxandwehave˚01;k(x)=Pkj=1�kj�xj�1=x�1�(1+x)k�1�:Thisleadsbysimpleintegrationcalculusto˚1;k(x)=Pkj=1(1+x)j j�Hk,whereHk=Pkj=1j�1.Thennotingthat˚02;k(x)=˚1;k(x),weobtain˚2;k(x)=Pkj=1(1+x)j+1
j(j+1)�xHk+1
k+1�1.Finally,itiseasytoshowthat 1;k(x)=1
k+1�˚1;k+1(x)�xk+1
k+1�and 2;k(x)=1
k+1�˚2;k+1(x)�xk+2
(k+1)(k+2)�.ThisyieldsS=�2˝2k�1Xi=11
i�˝
1+˝�i+2˝k"k�1
k˝+˝k�1Xi=11
i(i+1)�˝
1+˝�i#+2(n�k�1)˝"kXi=11
i�˝
1+˝�i+(�1)k k�1�1
(1+˝)k�#+(n�k�1)(k+1)"k k+1˝�˝kXi=11
i(i+1)�˝
1+˝�i+(�1)k k(k+1)�1�1
(1+˝)k�#=(n�k�1)k˝�2(k�1)˝2+n�k�1
k˝k+1
(1+˝)k+(�1)kn�k�1
k�1�1
(1+˝)k�(1+2˝)+(2˝(n�k�1)�2˝2)k�1Xi=11
i�˝
1+˝�i+(2˝2k�(n�k�1)(k+1)˝)k�1Xi=11
i(i+1)�˝
1+˝�i=2˝2�˝(n�k�1)+n�k�1
k˝k+1
(1+˝)k+(�1)kn�k�1
k�1�1
(1+˝)k�(2˝+1)�1
k�˝
1+˝�k�1�2k˝2�(n�k�1)(k+1)˝�+�2(n�2k�1)˝�2˝2+(k+1)(n�k�1)�k�1Xi=11
i�˝
1+˝�i;wherethelastequalitywasobtainedbywriting1
i(i+1)=1
i�1
i+1.Itsuﬃcesnowtoreinjectthisformulaintoequation(5)toobtaintheannouncedresult.�ProofofProposition3.5:ReasoningasintheproofofProposition3.1,weexpressE(1)n(˘k)asE(1)n(˘k)=�n�kXj=0E(1)n��minfH�j;H�j+kg�maxfH�j+1;:::;H�j+k�1g�+�;(6)withforany1�j�n�k�1,Q:=E(1)n��minfH�j;H�j+kg�maxfH�j+1;:::;H�j+k�1g�+�=Z10�Z10h(1)n(˝)Ptn(H�1>x)2Ptn(H�1<x)k�1d˝�dx;20
andforj2f0;n�kg,R:=E(1)n��minfH�j;H�j+kg�maxfH�j+1;:::;H�j+k�1g�+�=Z10�Z10h(1)n(˝)Ptn(H�1>x)Ptn(H�1<x)k�1d˝�dx:FromTheorem2.1,weknowthatPtn(H�1<x)=p(x^t)
1+p(x^t)1+pt pt,andfromProposition2.6,forallt�0,h(1)n(t)=pn(n�1)(pt)n�2
(1+pt)n+1.Afterachangeofvariables,thisleadstoQ=1
pn(n�1)Z10�x
1+x�k�1 Z1xtn�k�1
(1+t)n�k+2�1�x(1+t)
t(1+x)�2dt!dx=1
pn(n�1)Z10xk�1
(1+x)k+1�Jn�k+2;3(x)�2xJn�k+2;4(x)+x2Jn�k+2;5(x)�dx;R=1
pn(n�1)Z10�x
1+x�k�1�Z1xtn�k�1
(1+t)n�k+2�1�x(1+t)
t(1+x)�dt�dx=1
pn(n�1)Z10xk�1
(1+x)k�Jn�k+2;3(x)�xJn�k+2;4(x)�dx;whereforanyintegersk�l�2andforanypositiverealnumberx,Jk;l(x):=R1xuk�l
(1+u)kdu.Nowusing(9)inPropositionA.2toexpresstheintegralsJk;linRandQ,andusingagainPropositionA.2tocalculatetheremainingintegrals,weobtainforanyk�n�3,Q=1
p2n(n�1)
(n�k)(n�k+1)�n�k�1Xj=0j+1
(j+k)(j+k+1)(j+k+2)�2
(n�k�1)n�k�2Xj=0(j+1)(j+2)
(j+k+1)(j+k+2)(j+k+3)+1
(n�k�1)(n�k�2)n�k�3Xj=0(j+1)(j+2)(j+3)
(j+k+2)(j+k+3)(j+k+4)#;R=1
pn(n�1)
(n�k)(n�k+1)�n�k�1Xj=0j+1
(j+k)(j+k+1)+1
n�k�1n�k�2Xj=0(j+1)(j+2)
(j+k+1)(j+k+2)#:Finally,usingpartialfractiondecompositionstocalculatethesumsintheexpressionsofQandR,Q=1
pn(n�1)
(n�k)(n�k+1)�1
k+6
n�k�2�2(2n+k�1)
(n�k�1)(n�k�2)(Hn�1�Hk)�;R=1
pn(n�1)
(n�k�1)(n�k+1)�n+k�1
n�k(Hn�1�Hk�1)�2�:Reinjectingtheseexpressionsintoequation(6)leadstoE(1)n(˘k)=�
p[(n�k�1)Q+2R]=�
pn(n�1)
(n�k)(n�k�2)�n+k�2
k�2(n�1)
n�k�1(Hn�1�Hk)�;forany1�k�n�3,whichendstheproof.�21
4ConvergenceofgenealogiesinthelargesampleasymptoticInthissectionweprovideconvergenceresultsforthedistributionofthesuitablyrescaledgeneal-ogyofasampleofsizen,asn!1.Obtainingsuchasymptoticresultsrequiresanadditionalassumptiononthesamplingprobability:weassumethatthesamplingparameterpdependsonninsuchawaythatp=n=�,where�2R�+.Thisassumptionarisesnaturally,sinceitensuresthattheexpectednumberofsampledindividualsisofordern.Besides,accordingtoRemark2.8,notethattheparameter�willonlyhaveascalingeﬀectontime.Inthesequel,thesymbolL=meansanequalityinlaw,andforanyn>i�0,L(�;P(i)n)referstothedistributionofarandomvariableoraprocessunderP(i)n.Finally,)denotestheconvergenceindistribution.Recallfrom[13,Th.16.16]that,if(n)isasequenceofrandommeasuresonRdandasimplepointprocessonRd,n)iﬀn(B))(B)foranycompactsetBsuchthat(@B)=0,where@BdenotestheboundaryofB.4.1ResultsConvergenceofgenealogiesFirstdeﬁne,foranyt>0,ˇt(respˇ)asthePoissonpointmeasureon(0;1)�(0;�t)(resp.(0;1)�R�+)withintensity�dlx�2dx1(l;x)2(0;1)�(0;�t)(resp.�dlx�2dx1(l;x)2(0;1)�R�+).Let(ˆi)i�0beasequenceofi.i.d.exponentialrandomvariableswithparameter1=�,anddeﬁneforalli�0theinverse-gammarandomvariableei:=(ˆ0+:::+ˆi)�1.Thenfori2Z+,deﬁnethepair(ˇ(i);T(i)or),whereT(i)orisapositiverandomvariable,andˇ(i)isaCoxprocessˇ(i),asP(T(i)or2dt;ˇ(i)2�)=P(ei2dt)P(ˇt2�):Inparticular,conditionalonT(i)or=t,ˇ(i)hasthelawofthePoissonpointmeasureˇt.TheﬁrsttheoremstatestheconvergenceindistributionoftherandommeasureˇnunderP(i)n,i2Z+[f1g.ThisresultisageneralizationofCorollary2in[2],whichprovidesconvergenceindistributionofˇnunderP(0)ntowardsˇ(0).Theproofofthisconvergence,aswellastheproofofthegeneralizationwepropose,mainlyrelyontheconvergenceofˇnunderPtntowardsˇt,whichisestablishedinTheorem5in[19].Theorem4.1.Wehavethefollowingconvergencesindistributionasn!1:a)L(ˇn;P(1)n))ˇ;b)andforanyi�0,L((ˇn;Tor);P(i)n))(ˇ(i);T(i)or):Asacorollaryofthistheorem,westatetheﬁnitedimensionalconvergenceofthedivergencetimesofˇnunderP(i)n,i2Z+[f1g.Wedenoteby(Tk)k�1(resp.(T(i)k)k�1)thedecreasingreorderingofthesecondcoordinatesoftheatomsofˇ(resp.ˇ(i)).Corollary4.2.Fixk2N.Wehavethefollowingconvergencesindistributionasn!1:a)L((Tn;1;:::;Tn;k);P(1)n))(T1;:::;Tk);b)andforanyi�0,L((Tor;Tn;1;:::;Tn;k);P(i)n))(T(i)or;T(i)1;:::;T(i)k):22
Besides,thelimitingdistributionsappearinginCorollary4.2arespeciﬁedinthefollowingpropo-sition.Proposition4.3.a)Foranyk2N,thek-tuple(T1;:::;Tk)isdistributedas(e0;:::;ek�1).b)Foranyi2Z+,k2N,thek+1-tuple(T(i)or;T(i)1;:::;T(i)k)isdistributedas(ei;:::;ei+k).Thelasttheoremdescribesthelinksbetweenthediﬀerentrandommeasuresobtainedinthelimit,inTheorem4.1.Beforestatingthisresult,letusclarifysomedeﬁnition.Consider�anyrandommeasureamongˇ,ˇ(i)(i2Z+).Conditionalon�=Pt2A�(t;yt),whereAˆ[0;1]isacountableset,denotingby(u;yu)itslargestatom,wherewerefertotheorderw.r.t.thesecondcoordinate,wedeﬁnetherandommeasurePt2Anfug�(t;yt)astherandommeasureobtainedfrom�byremovingitslargestatom.Proposition4.3establishesinparticularthatforanyi2Z+,thetimeoforiginT(i)orisdistributedasthe(i+1)-thlargestatomoftherandommeasureˇ.Thefollowingstatementisadirectconsequenceofthisresult.Theorem4.4.Foranyi2Z+,themeasureˇ(i)hasthedistributionoftherandommeasureobtainedfromˇbyremovingitsi+1largestatoms.Inparticular,foranyi2N,themeasureˇ(i)hasthedistributionoftherandommeasureobtainedfromˇ(i�1)byremovingitslargestatom.Asaconclusion,inthelimitn!1,genealogieswithdiﬀerentpriorsonthetimeoforigincanallbeembeddedinthesamerealizationofthemeasureˇ:arealizationofthelimitingcoalescentpointprocesswithgivenpriorcanbeobtainedbyremovingfromarealizationofˇagivennumberofitslargestatoms.ConvergenceoftheexpectedsitefrequencyspectrumRecallthatmutationsareassumedtooccuratrate�onthelineages.WededucethefollowingpropositionfromtheresultsofSection3.2.Proposition4.5.Foranyt2R�+andanyi2f0;1g,foranyk2Nwehavelimn!1Etn(˘k)=��=kandlimn!1E(i)n(˘k)=��=k:Inotherwords,underPtn,P(0)nandP(1)n,theexpectedsitefrequencyspectrumofthesampleconverges,asthesizeofthesamplegetslarge,towardstheexpectedfrequencyspectrumoftheKingmancoalescent[23,(4.20)].4.2ProofsTobeginwith,westatetheconvergence,asn!1,oftheposteriordistributionofthetimeoforiginTorunderP(i)n.ThisresultisessentialtoobtainotherconvergenceresultsunderP(i)n,sincetheposteriordensityfunctionh(i)nofTorisdirectlyinvolvedinthedeﬁnitionofthelawP(i)n.Proposition4.6.Foranyi2Z+,wehavethefollowingconvergenceinlawL(Tor;P(i)n))T(i)or:23
Lemma4.7.Foranyi2Z+,therandomvariableeihasdensityfunctionh(i):t7!�i+1e��=t i!ti+21t>0(i.e.eifollowsaninverse-gammadistributionwithparameters(i+1;�)).Proof:Fixi2Z+.Therandomvariableeiistheinverseofthesumofi+1independentexponentialvariableswithparameter��1,i.e.theinverseofaGammavariablewithparameters(i+1;��1).Fromtheknowndensityfunctiont7!�i+1tie��t i!1t>0ofa�(i+1;��1)-variable,wededucethateihasdensityfunctiont7!�i+1e��=t i!ti+21t>0.�ProofofProposition4.6:Recallﬁrstthatbydeﬁnition,T(i)orisdistributedasei,andasaconsequence,hasdensityfunctionh(i).FromProposition2.6,recallingthatp=n=�,thedensityfunctionofTorunderP(i)nisgivenby:forallt>0,h(i)n(t)=n2
�i!(n�1):::(n�i)�nt=�
1+nt=��n+11
(nt=�)i+2=(n�1):::(n�i)
nie�(n+1)ln(1+�
nt)
i!�(t=�)i+2;(7)andhenceforallt>0h(i)n(t)!n!1�i+1e��=t i!ti+2=h(i)(t);andtheconvergenceofthedensityfunctions(h(i)n)towardsh(i)ensurestheconvergenceinlawunderP(i)nofTortowardsT(i)or.�ToproveTheorem4.1,weﬁrstrecallTheorem5of[19],whichcanbestatedasfollows.Lemma4.8.Foranyt>0,asn!1,L(ˇn;Ptn))ˇt:ProofofTheorem4.1:a)FromProposition2.4,underP(1)ntherandommeasureˇnisasimplepointprocesswithin-tensityPn�1i=1�fi=ng(dl)ndx
�(1+nx=�)2.Asn!1,thisintensitymeasureconvergesweaklytowards�dlx�2dx1(l;x)2(0;1)�(0;+1),whichistheintensitymeasureofthePoissonprocessˇ.From[13,Th.16.18],thisissuﬃcienttoprovetheconvergenceindistribution,underP(1)n,ofˇntowardsˇ.b)Fixi2Z+.ForanycompactsetAof[0;1]�R+,BBorelsetofR+ofzeroLebesguemeasureboundary,andk2Z+,wehaveP(i)n(ˇn(A)=k;Tor2B)=ZBPtn(ˇn(A)=k)h(i)n(t)dt:Inordertoapplythedominatedconvergencetheorem,weﬁrstremarkthatforallt>0,foranyn>i,h(i)n(t)�f(i)(t):=�i+1
i!ti+2;(8)ascaneasilybeseenfrom(7).Besides,studyingthevariationsofh(i)nyieldsinparticularthath(i)nisanonnegativefunctionthatincreaseson�0;�n�i�1
n(i+2)�.Nowthereexists�>0suchthatfornlargeenough,�n�i�1
n(i+2)��.Finallywehavefrom(8)thatforanynlargeenoughandforallt>0,jh(i)n(t)j�f(i)(�)1t��+f(i)(t)1t>�;whichisintegrableonR+.24
ItsuﬃcesnowtoinvokeLemma4.8andProposition4.6todeducebydominatedconvergencethatP(i)n(ˇn(A)=k;Tor2B)�!n!1ZBP(ˇt(A)=k)h(i)(t)dt=P(ˇ(i)(A)=k;ei2B);andthisendstheproof.�ProofofCorollary4.2:Hereweonlyprovea)sincetheproofofb)isidentical.Fixk2NandA1;:::Ak,BorelsetsofR�+ofzeroLebesguemeasureboundary,satisfyingsupAi=infAi�1foranyi2f2;:::;kg.WesetBi=(0;1)�Aiforany1�i�k.ThenP(1)n(Tn;12A1;:::;Tn;k2Ak)=P(1)n(ˇn(B1)=1;:::;ˇn(Bk)�1)�!n!1P(ˇ(B1)=1;:::;ˇ(Bk)�1)=P(T12A1;:::;Tk2Ak);wheretheconvergencefollowsfromTheorem4.1.Furthermore,thisresultclearlystillholdsifthesets(Ai)satisfysupAi�infAi�1insteadofsupAi=infAi�1.ToobtaintheresultinthecasewhereA1;:::Akarenonnecessarilypairwisedisjointsets,itsuﬃcestotodecompose[ki=1AiintoapartitionofdisjointBorelsetsandtoapplythesamereasoningasabove.Letusprovethisinthesimplecasek=2:P(Tn;12A1;Tn;22A2)=P(Tn;12A1\A2;Tn;22A1\A2)+P(Tn;12A1\A2;Tn;22A2nA1)+P(Tn;12A1nA2;Tn;22A1\A2)+P(Tn;12A1nA2;Tn;22A2nA1)=P(ˇn(A1nA2)=0;ˇn(A1\A2)�2)+P(ˇn(A1nA2)=0;ˇn(A1\A2)=1;ˇn(A2nA1)�1)+P(ˇn(A1nA2)=1;ˇn(A1\A2)�1)+P(ˇn(A1nA2)=1;ˇn(A1\A2)=0;ˇn(A2nA1)�1);andweconcludeasabove,usingTheorem4.1.�ProofofProposition4.3:a)WebaseourreasoningonthefactthataPoissonpointmeasureonR�+withintensitymeasure�x�2dxisthepushforwardmeasurebythecontinuousmappingx7!x�1ofaPoissonprocesswithparameter��1.Let�besuchaPoissonprocess.Thenforanya2R�+,recallingthatˆ0isanexponentialvariablewithparameter��1ande0=ˆ�10a.s.,P(T1�a)=P(ˇ((0;1)�(a;+1))�1)=P(�(0;a�1)�1)=P(ˆ0�a�1)=P(e0�a);andhenceT1L=e0.Asimilarreasoningshowsthatforanyk�1,A1;:::Ak,BorelsetsofR�+ofzeroLebesguemeasureboundary,satisfyingsupAi<infAi�1foranyi2f2;:::;kg,P(T12A1;:::;Tk2Ak)=P(e02A1;:::;ek�12Ak):Asinthepreviousproof,thecasewherethesets(Ai)arenonpairwisedisjointcanbeprovedwiththesamereasoning,decomposing[ki=1Aiintoapartitionofdisjointsets.Wecanthenconcludethat(T1;:::;Tk)isdistributedas(e0;:::;ek�1).25
b)Inthesameway,foranyt2R�+,aPoissonpointmeasureon(0;t)withintensitymeasure�x�2dx1(0;t)(x)isthepushforwardmeasurebythemappingx7!x�1oftherestrictionto(t�1;+1)ofaPoissonprocesswithparameter��1.Thenbydeﬁnitionofˇ(i),foranya;b2R�+,P(T(i)or�b;T(i)1�a)=Z+1a_bP(Tt1�a)P(ei2dt)=Z+1a_bP(ˇ((0;1)�(a;t))�1)P(ei2dt)=Za�1^b�10P(�((u;a�1))�1)P(e�1i2du)=P(�((e�1i;a�1))�1;e�1i�b�1):Nowforanyi�0,e�1iisdistributedasthe(i+1)-thatomof�,henceP(�((e�1i;a�1))�1;e�1i�b�1)=P(e�1i�b�1;e�1i+1�a�1):Asaconclusion,wehave(T(i)or;T(i)1)L=(ei;ei+1).Withasimilarreasoningweobtaintheequalityinlaw,foranyk�1,between(T(i)or;T(i)1;:::;T(i)k)and(ei;:::;ei+k).�ProofofTheorem4.4:Wedenoteby�ˇ(i)therandompointmeasureobtainedfromˇbyremovingitsilargestatoms.BytherestrictionpropertyofthePoissonpointmeasures,conditionalonTi=t,wehave�ˇ(i)L=ˇt.RecallingfromProposition4.3.(i)thatTiL=ei�1,foranya2R�+andk�0wehaveP(�ˇ(i)((a;+1))=k)=Z+10P(�ˇ(i)((a;+1))=kjTi=t)h(i�1)(t)dt=Z+1aP(ˇt((a;t))=k)h(i�1)(t)dt+1k=0Za0h(i�1)(t)dt=P(ˇ(i�1)((a;+1))=k);wherethelastequalityfollowsfromthedeﬁnitionoftheCoxprocessˇ(i�1).�AAppendixPropositionA.1.Foranyk2N,l2Zsatisfyingk�l,andx2R+,wedeﬁneIk;l(x):=Zt0tk�l
(1+t)kdtThenwehave(a)fork�0,Ik;0(x)=Rx0tk
(1+t)kdt=x�kln(1+x)+Pk�1j=1(�1)j�1
j�kj+1�(1�(1+x)�j),(b)fork�1,Ik;1(x)=Rx0tk�1
(1+t)kdt=ln(1+x)+Pk�1j=1(�1)j j�k�1j�(1�(1+x)�j),(c)fork�2,Ik;2(x)=Rx0tk�2
(1+t)kdt=1
k�1�x
1+x�k�1.26
Proof:Usingthebinomialtheoremtoexpandtk
(1+t)k=�1�1
1+t�k,wegetIk;0(x)=kXj=0�kj�(�1)jZx0(1+t)�jdt;andIk;1(x)=k�1Xj=0�k�1j�(�1)jZx0(1+t)�j�1dt;whicheasilyleadsto(a)and(b).�PropositionA.2.Foranyk2N,l2Zsatisfyingk�l,andanyx2R+,deﬁneJk;l(x):=Z1xtk�l
(1+t)kdt:Thenforanyt2R+,Jk;l(t)<1ifandonlyifl�2.InthiscasewehaveJk;l(x)=k�lXj=0xj
(1+x)j+l�1(j+1):::(j+l�2)
(k�1):::(k�l+1);(9)andinparticularJk;l(0)=(l�2)!
(k�1):::(k�l+1)=h(l�1)�k�1l�1�i�1:Proof:Firstforanyl�2andx>0,Jl;l(x)=R1xdt
(1+t)l=1
l�1(1+x)1�l.Foranyk�l�2andx�0,anintegrationbypartsgivesJk;l(x)=k k�l+1Jk+1;l�xk�l+1
(k�l+1)(1+x)k:Then,assumingthatJk;l(x)=Pk�lj=0xj
(1+x)j+l�1(j+1):::(j+l�2)
(k�1):::(k�l+1),weobtainJk+1;l(x)=k�lXj=0xj
(1+x)j+l�1(j+1):::(j+l�2)
k:::(k�l+2)+xk�l+1
k(1+x)k=k�l+1Xj=0xj
(1+x)j+l�1(j+1):::(j+l�2)
k:::(k�l+2);and(9)isthenprovedbyinductiononk.�References[1]D.Aldous.ThecontinuumrandomtreeIII.TheAnnalsofProbability,pages248–289,1993.[2]D.AldousandL.Popovic.Acriticalbranchingprocessmodelforbiodiversity.Advancesinappliedprobability,37(4):1094–1115,2005.[3]N.ChampagnatandA.Lambert.SplittingtreeswithneutralPoissonianmutationsI:Smallfamilies.StochasticProcessesandtheirApplications,122(3):1003–1033,2012.27
[4]N.ChampagnatandA.Lambert.SplittingtreeswithneutralPoissonianmutationsII:LargestandOldestfamilies.StochasticProcessesandtheirApplications,2012.[5]N.Champagnat,A.Lambert,andM.Richard.Birthanddeathprocesseswithneutralmutations.InternationalJournalofStochasticAnalysis,2012,2012.[6]H.A.DavidandH.N.Nagaraja.Orderstatistics.WileyOnlineLibrary,1970.[7]C.Delaporte.LévyprocesseswithmarksII:Invarianceprincipleforbranchingprocesseswithmutations.EprintarXiv:1305.6491,2013.[8]R.Durrett.Probabilit

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

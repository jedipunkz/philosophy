---
source: "https://arxiv.org/abs/2307.08988v1"
title: "EVIL: Evidential Inference Learning for Trustworthy Semi-supervised Medical Image Segmentation"
author: "Yingyu Chen, Ziyuan Yang, Chenyu Shen, Zhiwen Wang, Yang Qin, Yi Zhang"
year: "2023"
publication: "arXiv preprint / cs.CV"
download: "https://arxiv.org/pdf/2307.08988v1"
pdf: "https://arxiv.org/pdf/2307.08988v1"
captured_at: "2026-05-09T12:12:50Z"
updated_at: "2026-05-09T12:12:50Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche Beyond Good and Evil"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# EVIL: Evidential Inference Learning for Trustworthy Semi-supervised Medical Image Segmentation

- 著者: Yingyu Chen, Ziyuan Yang, Chenyu Shen, Zhiwen Wang, Yang Qin, Yi Zhang
- 年: 2023
- 掲載情報: arXiv preprint / cs.CV
- 情報源: [arxiv](https://arxiv.org/abs/2307.08988v1)
- ダウンロード: https://arxiv.org/pdf/2307.08988v1
- PDF: https://arxiv.org/pdf/2307.08988v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Recently, uncertainty-aware methods have attracted increasing attention in semi-supervised medical image segmentation. However, current methods usually suffer from the drawback that it is difficult to balance the computational cost, estimation accuracy, and theoretical support in a unified framework. To alleviate this problem, we introduce the Dempster-Shafer Theory of Evidence (DST) into semi-supervised medical image segmentation, dubbed Evidential Inference Learning (EVIL). EVIL provides a theoretically guaranteed solution to infer accurate uncertainty quantification in a single forward pass. Trustworthy pseudo labels on unlabeled data are generated after uncertainty estimation. The recently proposed consistency regularization-based training paradigm is adopted in our framework, which enforces the consistency on the perturbed predictions to enhance the generalization with few labeled data. Experimental results show that EVIL achieves competitive performance in comparison with several state-of-the-art methods on the public dataset.

## PDF Text

EVIL:EVIDENTIALINFERENCELEARNINGFORTRUSTWORTHYSEMI-SUPERVISEDMEDICALIMAGESEGMENTATIONYingyuChen1ZiyuanYang1ChenyuShen1ZhiwenWang1YangQin1YiZhang2,1,SeniorMember,IEEE1CollegeofComputerScience,SichuanUniversity,Chengdu,China2SchoolofCyberScienceandEngineering,SichuanUniversity,Chengdu,ChinaABSTRACTRecently,uncertainty-awaremethodshaveattractedin-creasingattentioninsemi-supervisedmedicalimageseg-mentation.However,currentmethodsusuallysufferfromthedrawbackthatitisdifﬁculttobalancethecomputationalcost,estimationaccuracy,andtheoreticalsupportinauni-ﬁedframework.Toalleviatethisproblem,weintroducetheDempster–ShaferTheoryofEvidence(DST)intosemi-supervisedmedicalimagesegmentation,dubbedEVidentialInferenceLearning(EVIL).EVILprovidesatheoreticallyguaranteedsolutiontoinferaccurateuncertaintyquantiﬁca-tioninasingleforwardpass.Trustworthypseudolabelsonunlabeleddataaregeneratedafteruncertaintyestimation.Therecentlyproposedconsistencyregularization-basedtrainingparadigmisadoptedinourframework,whichenforcestheconsistencyontheperturbedpredictionstoenhancethegen-eralizationwithfewlabeleddata.ExperimentalresultsshowthatEVILachievescompetitiveperformanceincomparisonwithseveralstate-of-the-artmethodsonthepublicdataset.IndexTerms—MedicalImageSegmentation,Semi-SupervisedLearning,EvidentialLearning1.INTRODUCTIONMedicalimagesegmentationplaysanessentialroleinsub-sequentclinicalorcomputer-aideddiagnosisandfully-supervisedlearninghasachievedgreatsuccessintheﬁeldofautomaticimagesegmentation[1].However,annotatingmedicalimagesislaboriousandrequiresrichprofessionalknowledge[2].Semi-supervisedlearning(SSL)hasshowngreatpoten-tialtoalleviatethisproblembyleveragingalargesetofun-labeleddataaccompaniedwithalimitednumberoflabeleddata.Thesemethodscanberoughlycategorizedintotwotypes:(1)pseudo-labelretraining,whichincorporatespseudolabelsonunlabeleddataforretraining[3,4,5];and(2)con-sistencyregularization,whichenforcesthepredictionconsis-tencytoenhancegeneralizationwithvariousperturbations,
YiZhangisthecorrespondingauthorsuchasinputperturbation,featureperturbation,andnetworkperturbation[6,7,8].However,sincethesemethodsrelyheavilyonthepredic-tionofpseudolabel,falsepredictionswillseverelydegradethesegmentationperformance.Toimprovethequalityofpseudolabels,someuncertainty-awaremethodshavebeenproposed,includingMonteCarlodropout(MC-dropout)-based[9],Information-Entropy-based[10],andPredictionVariance-based[11]methods.However,thesemethodssufferfromsomeproblems:(1)AlthoughMC-dropoutismathe-maticallyguaranteedbyBayesiantheory,itstrainingprocessiscostlyduetothemultiplesamplingoperations;(2)Duetothelimitedsamplingtimes,MC-dropoutcan’tobtainac-curateuncertaintyquantiﬁcation;(3)Othertwouncertaintyestimationmethodshaveadvantagesincomputationalcost,buttheylacktheoreticalsupport,leadingtounstablepseudolabelgeneration.Tohandletheaboveissues,weintroducetheDemp-ster–ShaferTheoryofEvidence(DST)intosemi-supervisedmedicalimagesegmentation,providingatheoreticallyguar-anteedsingle-passsolutionforuncertaintyquantiﬁcationinference,dubbedEVidentialInferenceLearning(EVIL).Followingthetrainingparadigmproposedin[7],EVILbe-longstotheconsistencyregularizationmethodwithnetworkperturbation,whichimposesthepredictionconsistencyontwonetworksperturbedwithdifferentinitialization.Inpar-ticular,thetwonetworksplaydifferentroles.Oneisavanillasegmentationnetwork(S-Net)whichdirectlygeneratesthesegmentationresult.Theothernetworkcalledevidential-network(E-Net)isbuiltfromtheperspectiveofDST,whichistheoreticallyguaranteedforreliablepredictions.DifferentfromS-Net,theoutputofE-NetisregardedastheevidenceandparameterizedintoaDirichletdistributiononsegmen-tationprobabilities.SubjectiveLogictheory(SL)[12]isemployedtoquantifythepredictionsanduncertaintiesofdifferentcategorieswiththeDirichletdistributioninasin-gleinference,whichsigniﬁcantlyreducesthetrainingtime.Then,thetrustworthypseudolabelsonunlabeleddataaregenerated.Insummary,therearethreemeritsforourpro-posedEVIL:lowercomputationcostduetothesingle-pass arXiv:2307.08988v1 [cs.CV] 18 Jul 2023
Fig.1.TheoverviewofourEVidentialInferenceLearningframework(EVIL),whereMdenotesuncertaintymapesti-matedbyE-Netand⊙denoteselement-wiseproduct.‘→’presentsforwardoperation,‘���’presentssupervisionlossoperationand‘//’on‘→’presentsstop-gradient.operation,accurateuncertaintyestimationbasedonSLandtheoreticalguaranteebasedonDST.Themaincontributionsofthisworkaresummarizedas:1)weintroduceDSTintoSSLandprovideafastaccurateuncertaintyestimationwiththeoreticalguaranteeinauniﬁedframework;2)anovelnetworkperturbationstrategyispro-posed,whichallowsdifferentinitializednetworkoptimizedwithdifferentobjectives;and3)extensiveexperimentsareconductedtovalidatetheeffectivenessofourproposedEVIL.2.METHODGivenalabeledsetDl={(xi,yi)}Nli=1withNlsamplesandanunlabeledsetDu={xi}Nui=1withNusamples,whereNu≫Nlinsemi-supervisedtask.AsillustratedinFig.1,EVILhastwodifferentlyinitial-izednetworks,E-NetF1withparametersetθ1andS-NetF2withparametersetθ2.Forlabeleddata,S-Netisoptimizedwithtraditionaljointcross-entropylossanddiceloss,whileE-NetmodelsaDirichletdistributionandisoptimizedwithevidentialsegmentationloss.P1,P2arethesegmentationpredictionsandY1,Y2arethecorrespondingpseudolabelsgeneratedbyargmaxfunction.Forunlabeleddata,E-NetgeneratespseudolabelsY1andaccurateuncertaintyestima-tionsMsimultaneously.Then,thetrustworthypseudolabelsarecalculatedbyM⊙Y1andusedtoguidethetrainingofS-Net.Reversely,thepseudolabelsY2generatedbyS-NetisleveragedforE-Nettoexploremorepotentialevidencetoimprovethegenerationofpseudolabelsfromunlabeleddata.2.1.UncertaintyModelingInthissection,weutilizeDSTtomodelthesegmentationun-certaintyandgeneratetrustworthyprediction.ForaK-classsegmentationtask,givenaninputxi,theevidencevectoreiisobtainedwithatransformfunctiong,whichisdeﬁnedin[13]:ei=g(F1(xi))=exp(tanhF1(xi)/τ),(1)
Fig.2.SubjectiveLogicmodel,whereui+�Kk=1bki=1.where0<τ<1isascalingparametersetto1/K.F1(xi)istheoutputofE-Netwithinputxi.SubjectiveLogic[12]computesthebeliefmassforcategorykanduncertaintyas:bki=eki
Si=αki−1
Siandui=K
Si,(2)whereSi=�Kk=1(eki+1),ui+�Kk=1bki=1andαki=eki+1.αi=�α1i,...,αKi�canberegardedastheparametersofDirichletdistribution,whichmodelsthedensityofsegmen-tationprobabilityanduncertainty[14].Thedensityfunctionisdeﬁnedas:D(pi|αi)={1
B(αi)�Kk=1pαki−1iforpi∈SK,0otherwise,(3)wherepiisthesegmentationprobability,B(αi)istheK-dimensionalmultinomialbetafunctionforparameterαi,andSKistheK-dimensionalsimplex.2.2.EvidentialNet(E-Net)Wefollow[14]andusecross-entropylosstomaketheseg-mentationprobabilitiespiapproachtheground-truthyi.No-tably,thedensityofpifollowstheDirichletdistributionpa-rameterizedwithαi.Thelosscanbeformulatedas:Ldig=��K�k=1−ykilog�pki��D(pi|αi)dpi=K�k=1yki�ψ(Si)−ψ�αki��,(4)whereψ(·)isthedigammafunction.ByoptimizingLdig,theevidenceofdifferentclassesforpositivesamplesisgenerated.However,Ldigcannotguaranteethatnegativesamplesgener-ateevidenceascloseaszero.Therefore,Kullback-Leibler(KL)divergenceisincorporatedintoourlossfunctiontope-nalizethedivergencefromnegativesamples,whichisdeﬁnedas:LKL=KL[D(pi|˜αi)∥D(pi|1)]=logΓ��Kk=1�αki�
Γ(K)�Kk=1Γ��αki�+K�k=1��αki−1��ψ��αki�−ψ�K�k=1�αki��,(5)whereΓ(·)isthegammafunction,D(pi|1)istheuniformDirichletdistribution,and˜αi=yi+(1−yi)⊙αi.
(a)image&GT
(b)MT
(c)UAMT
(d)ICT
(e)CPS
(f)URPC
(g)EVIL(ours)Fig.3.Visualcomparisonofsegmentationresultswithdifferentmethodswith10%labeledimages.Forsegmentationtask,theevidenceeiisobtainedwithxi.Then,αi=ei+1isparameterizedintothecorrespondingDirichletdistributionandtheevidentiallossis:Levi=Ldig+βLKL,(6)whereβisaannealingcoefﬁcientandissettoβ(t)=min(1.0,t
0.5tmax).tisthecurrentepochandtmaxisthetotalnumberoftrainingepochs.AsshowninFig.2,theSubjectiveLogicmodelhastwoparts,thecertainpartcalledbeliefmassbiandtheuncertainpartui.Theevidentiallossgeneratesevidencetoreducetheuncertainty.However,sincethecross-entropybasedeviden-tiallossisbasedonpixellevel,whichignorestherelation-shipsbetweenpixelsinsegmentationtask,weusetheDicelossonthecertainpartandthecertainlossisdeﬁnedas:Lcertain=1−2�Kk=1yki�Kk=1ˆpki
�Kk=1yki+�Kk=1ˆpki,(7)whereˆpi=softmax(bi)presentsasimplextransformedfromthebeliefmassbiwithasoftmaxfunction.Then,ouroverallevidentialsegmentationlossisdeﬁned:LEseg=Levi+γLcertain,(8)whereLeviandLcertaindenotetheevidentiallossandthecertainloss,respectively.γdenotestheweightingparame-ter,whichissetto1.ByoptimizingLevi,E-Netgeneratestheevidenceforpositivesamples,whilereducestheevidencefornegativesamples.Lcertainisleveragedtoconstraintherelationshipbetweendifferentpredictedpixels.2.3.EVILFrameworkThetotallossLforourwholeframeworkcontainstwocom-ponents:supervisedlossLsuponlabeleddataandconsis-tencylossLcononunlabeleddata:L=Lsup+λLcon,(9)whereλisthebalancingparameter.WeuseGaussianramp-upfunctionλ(t)=λmax∗e−5(1.0−t tmax)2andλmax=0.1.Thesupervisionlossisformulatedas:Lsup=LEseg(F1(x),y)+LSseg(F2(x),y),(10)whereLSseg=1
2(Lce+Ldice)denotesthelosscomponentforS-Net.LceandLdicearethecross-entropylossanddiceloss,respectively.ThepseudolabelcanbecalculatedasY1=argmax(bi)forE-NetandY2=argmax(F2(x))forS-Net.Theconsis-tencylossontheunlabeleddataiswrittenas:Lcon=Levi(F1(x),Y2)+Lce(F2(x),M⊙Y1).(11)whereM=u<TisthemasktoﬁlterouthighuncertainresultswiththresholdT=0.2.Weonlyusetheevidentialandcross-entropylossesinconsistencylosstermduetothemaskoperationwhichpreservesonlythereliablepseudopixellabels.TheconsistencylossencouragesE-NettogeneratepotentialevidencefromS-NetusingLeviandS-NettolearnthereliablepseudolabelsusingLcefromE-Net.3.EXPERIMENT3.1.ExperimentSetupWeevaluateourmethodontheAutomatedCardiacDiagno-sisChallenge(ACDC)[15]datasetwhichcontains200anno-tatedshort-axiscardiacMR-cineimagesfrom100patients.Weleverage70patients(140scans)fortraining,10patients(20scans)forvalidationand20patients(40scans)fortest-ing.Allshort-axissliceswithin3Dscansareresizedto256×256as2Dimages.SeeSSL4MIS1formoredetails.Forsemi-supervisedexperiments,imagesfrom7patients,14pa-tientsand21patientsaresetaslabeledratio10%,20%and30%inthetrainingset,respectively.Standarddataaugmen-tation,includingrandomcropping,randomrotating,andran-domﬂipping,isusedtoenlargethetrainingset.Threewidely
1https://github.com/HiLab-git/SSL4MIS
Table1.ThecomparisonofdifferentmethodsonACDCdatasetondifferentsemi-supervisedlabeleddataratiosettings.
Method
10%
20%
30%
DSC↑HD95↓ASD↓
DSC↑HD95↓ASD↓
DSC↑HD95↓ASD↓
Unet
80.057.412.38
84.908.942.52
87.076.611.95E-Net(ours)
81.0511.173.26
85.687.392.12
87.458.122.23
MT
81.0610.172.64
86.018.132.40
87.374.811.49UA-MT
80.8111.733.52
85.387.772.70
87.536.322.05ICT
83.548.422.46
85.285.651.64
87.498.252.23CPS
84.708.252.35
87.475.981.74
88.216.491.90URPC
82.075.621.88
85.135.711.75
86.994.431.31EVIL(ours)
85.913.911.36
88.224.011.21
89.433.841.07
Table2.Thecomparisonoftrainingtime.
Method
Num
Uncertainty
Time
Cost
Unet
1
×
0.076s
-ICT
1
×
0.090s
+18.42%URPC
1
√
0.089s
+17.11%E-Net(ours)
1
√
0.085s
+11.84%
MT
2
×
0.101s
-CPS
2
×
0.137s
+35.64%UA-MT
2
√
0.337s
+233.66%EVIL(ours)
2
√
0.148s
+46.53%
usedmetrics,DiceCoefﬁcient(DSC),HausdorffDistance95(HD95)andAverageSurfaceDistance(ASD)areemployedtoevaluatetheperformanceofourmethod.Forthesakeoffairness,Unet[1]ischosenastheback-boneinallmethods,andSGDisadoptedastheoptimizer.Theinitiallearningrateissetto0.01,andpolynomialsched-ulerstrategyisemployedtoupdatethelearningrate.Weim-plementtheproposedframeworkwithPyTorch,usingasingleNVIDIAGTX1080TiGPU.Thebatchsizeissetto24,where12imagesarelabeled.Allmethodsperform30000iterationsduringtraining.3.2.ExperimentalResultsSeveralrecentlyproposedsemi-supervisedsegmentationmethodsarecompared,including:Mean-Teacher(MT)[6],Uncertainty-AwareMeanTeacher(UA-MT)[9],Interpola-tionConsistencyTraining(ICT)[16],CrossPseudoSuper-vision(CPS)[7],andUncertaintyRectiﬁedPyramidCon-sistency(URPC)[17].Forallcompetingmethods,ofﬁcialparametersettingsareadopted.Tab.1illustratesthequantitativeresultsonACDC.Theﬁrstandsecondrowslistthequantitativeresultsofsuper-visedUnetandE-Net.Indifferentlabeleddataratiosettings,EVILoutperformsalltheothermethods.Whenonly10%ofdataarelabeled,ourmethodimprovesDSCbymorethan3%comparedwithotherSOTAuncertainty-awaremethods(UAMTandURPC).Moreover,weachieve4pointsimprove-mentinHD95and1pointinASDcomparedwithCPS.Es-pecially,wecanseethattheperformanceofEVILusing20%labeleddatahassurpassedallcomparedmethodsusing30%labeleddata.
(a)Input
(b)Label
(c)Ours
(d)EVIL
(e)S=2
(f)S=8
(g)S=64
(h)UA-MTFig.4.Visualizationofuncertaintyestimation.‘S’denotestheMC-dropoutsamplingtimes.Fig.3visualizesthesegmentationresultsoftwocasesus-ingdifferentmethodswith10%labeleddata.Itiseasytoseethatthecomparedmethodsmis-classifymanypixelswhileEVILobtainsmoreaccurateprediction.AsshowninFig.4,samplingtimesaffecttheuncertaintyestimationqualityofMC-dropoutandourE-Nethasbestaccurateestimation.Tab.2showsthetrainingtimewithﬁxedbatchsize=24,where‘Num’,‘Uncertainty’,‘Time’,‘Cost’denotesthenetworknumber,uncertainty-basedornot,timeconsuming,andtheadditionaltimeconsumingcostrespectively.WetreatUnetastheupperboundofsinglenetworkmethodandMTasthebaselineofthemulti-networkframeworksinceitisthefastestmethodcomparedtoothers.Specially,wecanseethattheproposedmethodimprovesigniﬁcantlywithoutintroduc-ingtoomuchcomputationoverhead.4.CONCLUSIONInthispaper,weproposeanoveluncertainty-awaresemi-supervisedmedicalimagesegmentationframework.Thepro-posedEVILintroducesDSTintotheconsistencyregulariza-tiontrainingparadigmandachievesfastaccurateuncertaintyestimationwithsolidtheoreticalguarantee.Extensiveexperi-mentsdemonstratethatEVILachievesstate-of-the-artperfor-manceonthewidelyusedACDCdataset.5.CONFLICTSOFINTERESTTheauthorsdeclarethattheyhavenoconﬂictsofinterest.
6.COMPLIANCEWITHETHICALSTANDARDSThisresearchstudywasconductedretrospectivelyusingrealclinicalexamsacquiredattheUniversityHospitalofDijon.Ethicalapprovalwasnotrequiredasconﬁrmedbythelicenseattachedwiththeopenaccessdata.7.ACKNOWLEDGEMENTThisworkwassupportedinpartbytheNationalNaturalSci-enceFoundationofChinaunderGrant62271335;inpartbytheSichuanScienceandTechnologyProgramunderGrant2021JDJQ0024;andinpartbytheSichuanUniver-sity“From0to1”InnovativeResearchProgramunderGrant2022SCUH0016.8.REFERENCES[1]OlafRonneberger,PhilippFischer,andThomasBrox,“U-net:Convolutionalnetworksforbiomedicalimagesegmentation,”inInternationalConferenceonMedicalImageComputingandComputer-AssistedIntervention(MICCAI).Springer,2015,pp.234–241.[2]KeZou,XuedongYuan,XiaojingShen,MengWang,andHuazhuFu,“Tbrats:Trustedbraintumorsegmen-tation,”inInternationalConferenceonMedicalImageComputingandComputer-AssistedIntervention(MIC-CAI).Springer,2022,pp.503–513.[3]BarretZoph,GolnazGhiasi,Tsung-YiLin,YinCui,HanxiaoLiu,EkinDogusCubuk,andQuocLe,“Re-thinkingpre-trainingandself-training,”inAdvancesinNeuralInformationProcessingSystems,2020,vol.33,pp.3833–3845.[4]ZhengyangFeng,QianyuZhou,GuangliangCheng,XinTan,JianpingShi,andLizhuangMa,“Semi-supervisedsemanticsegmentationviadynamicself-trainingandclassbalancedcurriculum,”arXivpreprintarXiv:2004.08514,vol.1,no.2,pp.5,2020.[5]MostafaSIbrahim,ArashVahdat,ManiRanjbar,andWilliamGMacready,“Semi-supervisedsemanticim-agesegmentationwithself-correctingnetworks,”inIEEE/CVFConferenceonComputerVisionandPatternRecognition(CVPR),2020,pp.12715–12725.[6]AnttiTarvainenandHarriValpola,“Meanteachersarebetterrolemodels:Weight-averagedconsistencytargetsimprovesemi-superviseddeeplearningresults,”inAd-vancesinNeuralInformationProcessingSystems,2017,vol.30.[7]XiaokangChen,YuhuiYuan,GangZeng,andJingdongWang,“Semi-supervisedsemanticsegmentationwithcrosspseudosupervision,”inIEEE/CVFConferenceonComputerVisionandPatternRecognition(CVPR).IEEE,2021,pp.2613–2622.[8]YassineOuali,C´elineHudelot,andMyriamTami,“Semi-supervisedsemanticsegmentationwithcross-consistencytraining,”inIEEE/CVFConferenceonComputerVisionandPatternRecognition(CVPR),2020,pp.12674–12684.[9]LequanYu,ShujunWang,XiaomengLi,Chi-WingFu,andPheng-AnnHeng,“Uncertainty-awareself-ensemblingmodelforsemi-supervised3dleftatriumsegmentation,”inInternationalConferenceonMedicalImageComputingandComputer-AssistedIntervention(MICCAI).Springer,2019,pp.605–613.[10]TaoWang,JianglinLu,ZhihuiLai,JiajunWen,andHengKong,“Uncertainty-guidedpixelcontrastivelearningforsemi-supervisedmedicalimagesegmenta-tion,”inInternationalJointConferencesonArtiﬁcialIntelligence,2022.[11]ZhedongZhengandYiYang,“Rectifyingpseudolabellearningviauncertaintyestimationfordomainadaptivesemanticsegmentation,”InternationalJournalofCom-puterVision,vol.129,pp.1106–1120,2021.[12]AudunJsang,SubjectiveLogic:Aformalismforrea-soningunderuncertainty,Springer,2018.[13]YangQin,DezhongPeng,XiPeng,XuWang,andPengHu,“Deepevidentiallearningwithnoisycorrespon-denceforcross-modalretrieval,”inACMInternationalConferenceonMultimedia,2022,p.4948–4956.[14]MuratSensoy,LanceKaplan,andMelihKandemir,“Evidentialdeeplearningtoquantifyclassiﬁcationun-certainty,”inAdvancesinNeuralInformationProcess-ingSystems,2018,vol.31.[15]OlivierBernard,AlainLalande,ClementZotti,Fred-erickCervenansky,XinYang,Pheng-AnnHeng,IremCetin,KarimLekadir,OscarCamara,MiguelAn-gelGonzalezBallester,etal.,“Deeplearningtechniquesforautomaticmricardiacmulti-structuressegmentationanddiagnosis:istheproblemsolved?,”IEEETrans-actionsonMedicalImaging,vol.37,pp.2514–2525,2018.[16]VikasVerma,KenjiKawaguchi,AlexLamb,JuhoKan-nala,ArnoSolin,YoshuaBengio,andDavidLopez-Paz,“Interpolationconsistencytrainingforsemi-supervisedlearning,”NeuralNetworks,vol.145,pp.90–106,2022.[17]XiangdeLuo,GuotaiWang,WenjunLiao,JienengChen,TaoSong,YinanChen,ShichuanZhang,Dim-itrisNMetaxas,andShaotingZhang,“Semi-supervisedmedicalimagesegmentationviauncertaintyrectiﬁedpyramidconsistency,”MedicalImageAnalysis,vol.80,pp.102517,2022.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

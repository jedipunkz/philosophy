---
source: "https://arxiv.org/abs/0705.1480v1"
title: "Anomalous Diffusion of particles with inertia in external potentials"
author: "S. Eule, R. Friedrich, F. Jenko"
year: "2007"
publication: "arXiv preprint / cond-mat.stat-mech"
download: "https://arxiv.org/pdf/0705.1480v1"
pdf: "https://arxiv.org/pdf/0705.1480v1"
captured_at: "2026-05-09T12:30:17Z"
updated_at: "2026-05-09T12:30:17Z"
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

# Anomalous Diffusion of particles with inertia in external potentials

- 著者: S. Eule, R. Friedrich, F. Jenko
- 年: 2007
- 掲載情報: arXiv preprint / cond-mat.stat-mech
- 情報源: [arxiv](https://arxiv.org/abs/0705.1480v1)
- ダウンロード: https://arxiv.org/pdf/0705.1480v1
- PDF: https://arxiv.org/pdf/0705.1480v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Recently a new type of Kramers-Fokker-Planck Equation has been proposed [R. Friedrich et al. Phys. Rev. Lett. {\bf 96}, 230601 (2006)] describing anomalous diffusion in external potentials. In the present paper the explicit cases of a harmonic potential and a velocity-dependend damping are incorporated. Exact relations for moments for these cases are presented and the asymptotic behaviour for long times is discussed. Interestingly the bounding potential and the additional damping by itself lead to a subdiffussive behaviour, while acting together the particle becomes localized for long times.

## PDF Text

arXiv:0705.1480v1 [cond-mat.stat-mech] 10 May 2007
AnomalousDiﬀusionofparticleswithinertiainexternalpotentialsS.Eule,R.FriedrichandF.JenkoNovember29,2018AbstractRecentlyanewtypeofKramers-Fokker-PlanckEquationhasbeenpro-posed[R.Friedrichetal.Phys.Rev.Lett.96,230601(2006)]describinganomalousdiﬀusioninexternalpotentials.Inthepresentpapertheex-plicitcasesofaharmonicpotentialandavelocity-dependenddampingareincorporated.Exactrelationsformomentsforthesecasesarepresentedandtheasymptoticbehaviourforlongtimesisdiscussed.Interestinglytheboundingpotentialandtheadditionaldampingbyitselfleadtoasub-diﬀussivebehaviour,whileactingtogethertheparticlebecomeslocalizedforlongtimes.1Introduction
InrecentyearstheconceptofContinousTimeRandomWalks(CTRWs)[1]hasbeenextensivelyusedtomodelsystemsthatexhibitthephenomenonwhichhasbeentermed”AnomalousDiﬀusion”.Variousmechanismsinphysicalsys-temsareknownleadingtosub-andsuperdiﬀusionrespectively.CTRWsarethussuitabletomodelavarietyofphysical,chemical,biologicalandevenmed-icalsystems,likechargetransportindisorderedsystems[2],proteindynamics[3],transportinlow-dimensionalchaoticsystems[4],anomaloustransportinplasmas[5,6,7]andthespreadofapandemic[8].Forareviewseee.g.[9].CTRWsinanaturalwayleadtoadescriptionwithintheframeworkoffrac-tionalkineticequationsandinparticularFractional-Fokker-PlanckEquations(FFPEs).
Inarecentpublication[10],[11]theconceptofCTRWshasbeenextendedtoparticleswithﬁniteinertiaundertheinﬂuenceofastochasticforce.Inthecontextofthismodeltheparticleissubjectedtoaseriesofrandomkicks,suchthattheparticle’smotionisdeterministicmostofthetime.Thisleadstoagen-eralizedequationofKramers-Fokker-Planck(KFP)typewhichischaracterizedbyacollisionoperatorthatisnonlocalinspaceaswellasintime.Foraspecialchoiceofthetime-evolutionkernelaFractionalKramers-Fokker-PlanckEqua-tionhasbeenobtained.Fractionalkineticequationsingeneralandfractionaldiﬀusionequations,inparticular,haveawiderangeofapplications.WeshouldmentionthatoverthepastyearsseveralgeneralizedKFP-equations[12],[13],[14]havebeenproposed(forareviewincludingdiscussionsofvariousapplica-tionsseee.g.,Ref.[15]).Inarecentpaper[16]theconnectionoftheseequationsisestablishedunsinganapproachbasedontheunderlyingLangevinequations.1
Inthispaperwewanttoextendthediscussionsoftheforce-freecasestartedin[11]anddiscusstheeﬀectsofaninclusionofalinearforceandanadditionaldampingbetweenthekicks.Thiscaseshouldbeofimportanceforvariouspracticalapplicationsinthecontextofanomalousdiﬀusioninrandomandac-cordinglycomplexenvironments.
Forresultsconcerninganomaloustransportofparticleswithoutinertiainexter-nalﬁeldsthereaderisreferedto[17].RecentresultsonCTRWsinanoscillatingexternalﬁeldaregivenin[18].
Thispaperisorganizedasfollows.Firstwepresentthemodelunderconsidera-tionandincorporatetheeﬀectsoftheharmonicpotentialandthedamping.Aswewillsee,theresultingequationwillberathercomplicated.Wethusconcen-trateontheequationsforthemomentsanddiscussinthiscontextsomelimitingcases.Wewillpresentanalyticalandnumericalresultsforthetime-evolutionofthesecond-ordermoments.
2Themodel
Some70yearsagoKramersinvestigatedthemotionofBrownianparticlesinaﬂuid.Inaseminalpaper[19]heconsideredthejoint-probability-distributionofpositionandvelocityf(x,u,t)forwhichhecouldderivethewell-knownequation:∂
∂t+∇x·u+∇u·A(x,u)f(x,u,t)=LFPf(x,u,t)(1)whereA(x,u)istheaccelerationduetoanexternalpotentialandLFPistheFokker-Planckcollisionoperator[20],[21]LFPf=γ∇u·(uf)+K∆uf.(2)Notethatthisnotationindicatesthatthedivergencesactontheproductofthefollowingfunctionandtheprobabilitydistribution.
In[10]anewtypeoftheKramers-Fokker-PlanckEquationhasbeenproposed.Inthecontextofthismodelthemotionoftheparticleisdeterministicwhilefromtimetotimeitissubjectedtoakickinarandomdirectionwithrandomvelocity.Asthetimeτbetweenthekicksaswellasthenewvelocityuisgivenbysomeprobabilitydistribution,aCTRW-modelinthevelocity-spacehasbeenobtained.Asamainresultfromthefollowingmasterequation∂
∂t+∇x·u+∇u·A(x,u)f(x,u,t)=Zt0dt0Φ(t−t0)Zdu0F(u;u0)Pt,t0f(x,u0,t0)−Zt0dt0Φ(t−t0)Pt,t0f(x,u,t0)(3)thegeneralizedKramersequation∂
∂t+∇x·u+∇u·A(x,u)f(x,u,t)==LFPZt0dt0Φ(t−t0)Pt,t0f(x,u,t0),(4)2
hasbeenderived,whereF(u;u0)du0denotestheprobabilitythattheparticle’svelocitywillendupinthevelocityspaceelementduaboutuandΦ(t−t0)issomememorykernel.ThedeterministicpropagationbetweentherandomkicksisgovernedbytheFrobenius-PerronoperatorPt,t0f(x,u,t0)=e−(t−t0)[∇x·u+∇u·A(x,u)]f(x,u,t0).(5)Eq.(4)generalizesEq.(1)intwoaspects.FirstitincorporateseﬀectswhicharenonlocalintimeandseconditcontainsretardationeﬀectsduetoPt,t0whichrenderstheequationnonlocalinspaceaswell.Forthespecialcaseofanasymptoticpower-lawwaiting-timedistributionΨ(τ)∼τ−(1+δ)oneobtainsRt0Φ(t−t0)∼0D1−δtwhere0D1−δtistheRiemann-Liouvillefractionalderiva-tive.ForanexcellentaccountonCTRWsthereaderisreferredto[7].AdetailedderivationofEq.(4)withanexplicitdeﬁnitionofF(u;u0)isfoundin[10].3Dampedmotioninaharmonicpotential
Inthissectionweshallconsiderthespecialcase,wheretheexternalpotentialisgivenbyA(x,u)=−2ηu−ω20x.Thecorrespondingequationsofmotionforanindividualparticlewithouttheimpactofﬂuctuationsread:˙x(t)=u(t),˙u(t)=A(x,u)=−2ηu−ω20x.(6)Thissetofequationsdeterminesthebehaviorofanindividualparticlebetweentwosuccesivekicks.AswehavealreadymentionedtheeﬀectofthePerron-Frobenius-Operatoristoprojectthepositionandvelocityoftheparticlebacktothelastkick.ForthisreasonwehavetosolveEq.(6)andinvertthissolutiontoretaintheretardationoftheprobabilitydistributionfunction(pdf)ontherighthandsideofEq.(4).Thesolutionreadsx(t)=e−ηtx0cos(ω0(t−t0))+u0+ηx0
ω0sin(ω0(t−t0))u(t)=e−ηt −ηu0−(η2+ω02)xo
ω0sin(ω0(t−t0)+u0cos(ω0(t−t0)!.(7)Herebyweimposedtheinitialconditionsx(t0)=x0,u(t0)=u0andusedtheabbrevationω0=p
ω20−η2.Notethatforη>ω0,ω0becomespurelyimaginaryandwehaveanexponentialrelaxationofx(t)andu(t).InvertingthisequationsgivesarathercomplicatedexpressionandsolvingEq.(4)withthisretardationwouldbeahopelesstask.Intheremainderofthispaperwethusﬁrstrestrictourselvestothespecialcaseswherewehaveeithersetω20=0oroneofthedampingconstantsη=0,γ=0respectively.Forthesespecialcasesrigorousresultsarepresentedinthenextsections.
4Dampedmotionbetweenthekicks
Inthissectionthelimitingcaseω0ηisstudiedi.e.wewillanalyzepurelydampedmotionbetweenthekicks.Themotionofsuchaparticleinphase-space3
isvisualizedinﬁgure1.Themotionofanindividualparticlebetweenthekicksisgivenby˙x(t)=u(t),˙u(t)=A(u)=−2ηu,(8)withthesolutionx(t)=x0+u01
2η1−e−2η(t−t0)u(t)=u0e−2η(t−t0).(9)AsalreadymentionedwehavetoinvertEq.(9)togettheretardationofthepdfinEq.(4),i.e.wehavetoregardtheeﬀectofthePerron-Frobenius-operatorx0=x(t)−u(t)1
2η1−e2η(t−t0)u0=u(t)e2η(t−t0).(10)TocalculatethetemporalbehaviorofthemomentsweuseEq.(4).Herewehavetobecareful.Onthelefthandsideof(4)wehavetoaverageoverthevaluesattwhileontherighthandsideweaverageoverretardedvaluesatt0.Consequentlywegetretardationeﬀectsinthemomentequations.Assumingthesystemtobeinﬁniteandspatiallyhomogeneousweobtainforthemomentsoflowestorder∂
∂tq(t)=Aq(t)+Zt0Φ(t−t0)(Bq(t0)+I)dt0.(11)Herewehaveintroducedq=
hx2ihuxihu2i
,I=
0
02K
,A=
020
0−2η100−4η
and(12)B=−γ

000
0e−2η(t−t0)e−2η(t−t0)1−e−2η(t−t0)00−2e−4η(t−t0)

.(13)Itisstraightforwardtogeneralizetheseequationsformomentsofanyorder.DuetotheoccurenceoftheconvolutionintegralsitisappropriatetoswitchtotheLaplacedomain,wherethissetofequationscanbesolved.hu2i(λ)=2KΦ(λ)
λ+4η+2γΦ(2η+λ),(14)huxi(λ)=1−γ
2ηΦ(2η+λ)+γ
2ηΦ(4η+λ)
λ+2η+γΦ(2η+λ)hu2i(λ)(15)hx2i(λ)=2
λhuxi(λ).(16)HerebywehaveusedtheLaplaceshiftingtheorem.Itisapparenthowthead-ditionaldampingaﬀectsthetimeevolutionkernel.Obviouslytheasymptoticbehaviourforlargetimes,i.e.thesmallλ-behaviour,isgivenbytheLaplacein-versionofΦ(λ)forhu2i(λ)andhuxi(λ)andbyΦ(λ)
λforhx2i(λ)respectively.4
Toobtainconcreteresultsweconsidertheimportantcaseofafractionalevolu-tionkernel,i.e.Φ(λ)=λ1−δ.Wegethu2i∼K
2η+γ(2η)1−δλ−δ(17)whichcorrespondstotheasymptoticbehaviourhu2i(t)∼K
Γ(δ)(2η+γ(2η)1−δ)t1−δ,(18)huxi(t)∼γ((2η)1−δ+(4η)1−δ)
4η2+γ(2η)2−δK
Γ(δ)(2η+γ(2η)1−δ)t1−δ.(19)Forthemean-square-displacementweobtainhx2i(t)∼2γ((2η)1−δ+(4η)1−δ)
4η2+γ(2η)2−δK
Γ(1+δ)(2η+γ(2η)1−δ)tδ.(20)Theadditionaldampingthusshiftsthesubballistic-superdiﬀusivebehaviourtothesubdiﬀusiveregime.FromEqs.(14)-(16)itisclearthatthemagnitudeofthedampingconstantiscrucialfortherateofconvergencetotheasymptotics.Intheclassicallimit,i.e.δ=1standarddiﬀusivebehaviourhx2i(t)∼tisobtained.
Tocompleteourdiscussionswementionthatforη→0wegettheresultsdescribedin[11].Forγ→0weobtainhu2i∼K
Γ(δ)2ηt1−δ,(21)huxi∼K
Γ(δ)4η2t1−δ,(22)andhx2i∼K
Γ(1+δ)4η2tδ.(23)Weseethateventheadditionaldampingaloneleadstosubdiﬀusivebehaviour.Thisisincontrasttothefreeﬂightcase,whereevenifγ6=0superdiﬀusive-suballisticbehaviourisobtained.
5Responsetoanexternalharmonicpotential
Letusnowexaminethemotionoftheparticlesexposedtoanexternallinearrestoringforce.Inotherwordswewillstudytheeﬀectsofanexternalharmonicpotentialactingontheparticlesbetweenthekicks.Asmentionedbeforedis-cussingthisproblemwithbothdampingsactingontheparticleleadstobulkyequations.Forthisreasonweconsiderﬁrstthesituationwhenηω20.Eq.(7)thenreadsx(t)=x0cos(ω(t−t0))+u0
ωsin(ω(t−t0))u(t)=−ωx0sin(ω(t−t0))+u0cos(ω(t−t0)),(24)5
ux
Figure1:Visualizationofdampedmotionbetweenthekicksinphase-spacewherewehaveimposedtheinitialconditionsx(t0)=x0andu(t0)=u0.Inthiscontextawordofcautionshallbeappropriate.Onehastobecarefulinsettingtheinitialconditions.
Wecannowrepeattheprocedureofthepreviouschapterforthecalculationofthemoments.Thenon-localityinspaceprovidesforamixingofthespaceandvelocitycoordinates.
Weobtain∂
∂tq(t)=Aq(t)+Zt0Φ(t−t0)(Bq(t0)+I)dt0.(25)Herewehaveintroducedq=
hx2ihuxihu2i
,I=
0
02K
,A=
020−ω0010−2ω00
and(26)B=−γ

000−ω20
2sin(2ω0(t−t0))cos(2ω0(t−t0))1
2ω0sin(2ω0(t−t0))2ω20sin2(ω0(t−t0))−2ω0sin(2ω0(t−t0))cos2(ω0(t−t0))

.(27)SwitchingtoLaplacespaceweobtainalinearalgebraicsystemλq(λ)=C(λ)q(λ)+I(28)whereC=
020γω0
4i(A−B)−ω0−γ
2(A+B)−γ
4iω0(A−B)+1γω20
2(A+B)−γω20Φ(λ)γ
i(A−B)−2ω20−γ
2(A+B)+γΦ(λ)
(29)andI=0,0,2KΦ(λ)
λ.Herebywehaveusedtheshort-handnotationA=Φ(λ−2iω0)andB=Φ(λ+2iω0).NotethattherelationshipΦ(z∗)=Φ∗(z)6
ux
Figure2:VisualizationofthemotioninaharmonicpotentialholdsandthusA∗=B,wherethe(·)∗meanscomplexconjugation.ItfollowsinthiscasethatChasonlyreal-valuedentries.IntheLaplace-domainweobtainalinearsystemofequationswhichissolvablebutwehavenotyetbeenabletotransformtheseexpressionsbackortoextractusefullimitingcases.Oneshouldnoticethatforthissystemtheknowledgeofjustthelong-timebehaviourisnolongersuﬃcient.Inparticulartheoscillationsofthemomentsforsmalltareignored.Employingthefractionaltimekernelwegetforδ=1theordinaryKramersequation.Itisinterestingtonotehowinthiscasetheadditionalcouplingsinthedynamicalsystem(25)disappear.Letusnowconcentrateonthegeneralcaseforγ=0.TheequationsforthemomentsuptosecondorderarethennolongerdiﬀerentfromaretardationfreeversionofEq.(4)(i.e.thefollowingresultsforthemomentsstatedherealsoholdfortheKramersequationproposedbyBarkaiandSilbey[12]).Thisallowsustoobtainconcreteresultsonthebehaviourofthemoments.
ItisimportanttoﬁrstconsiderthephysicalmeaningofγinEq.(2).Intheclas-sicalKramersequationthedampingoperatesbetweentwosuccessiverandomevents.InthegeneralizedKramersequationEq.(4)thedampingactsduringthekicks.ApositiveγaﬀectstheprobabilityF(u,u0)inawaythatitismoreprobabletohavesmallerabsolutevalueofu0afterakick.Thiseﬀectisvisual-izedinﬁgure3.Withoutthisdampingtheenergyoftheparticlewouldtendtoinﬁnityfortheforce-freecase.IfwehaveanadditionaldampingbetweenthekicksinthegeneralizedKramerseq.thoughwecansetγ=0withoutlossofphysicalsigniﬁcance.Thusthelimitingcaseγ→0isnotsolelyofacademicinterest.Thephase-portraitofsuchamotionisshowninﬁgure4.ThesetofequationsforthesecondordermomentsreadsinLaplacespacehu2i(λ)=2KΦ(λ)(λ2+2ηλ+2ω20)
λ(2η+λ)(λ2+4ηλ+4ω20),(30)huxi(λ)=2KΦ(λ)
(2η+λ)(λ2+4ηλ+4ω20),(31)7
ux
Figure3:TheeﬀectofγinEq.(4)visualizedforfreemotionbetweenthekicks ux
Figure4:Thecombinedeﬀectoftheharmonicpotentialanddampinginphase-space8
0
2
4
6
8
10
t
0.00
0.05
0.10
0.15
0.20
0.25
<x2>
Figure5:η=0,ω0=4hx2i(λ)=4KΦ(λ)
λ(2η+λ)(λ2+4ηλ+4ω20).(32)ToobtainconcreteresultsweconsideronceagainthecaseΦ(λ)=λ1−δhu2i(λ)=2Kλ−δ(λ2+2ηλ+2ω20)
(2η+λ)(λ2+4ηλ+4ω20),(33)huxi(λ)=2Kλ1−δ
(2η+λ)(λ2+4ηλ+4ω20),(34)hx2i(λ)=4Kλ−δ
(2η+λ)(λ2+4ηλ+4ω20).(35)TheinverseLaplace-Transformsoftheseexpressionscannowbeexplicitlycal-culated.
Forthemean-square-displacement(MSD)hx2i(t)weobtainhx2i(t)=1
ˆωΓ(δ)2−δ−1e−2t(η+√
ˆω)Ktδ2e2t√
ˆω(Γ(δ,−2ηt)−Γ(δ))(ηt)−δ+e4t√
ˆω(36)(t(√
ˆω−η))−δ(Γ(δ)−Γ(δ,(2t√
ˆω−η)))+(−t(√
ˆω−η))−δ(Γ(δ)−Γ(δ,(−2t√
ˆω−η)))Herebyˆωisgivenbyˆω=η2−ω20andΓ(α,β)istheincompleteGamma-function.(TheoccurenceoftheseincompleteGamma-functions,whichareintegralexpres-sions,isduetothemultipleoccurenceofconvolution-integralswhenbacktrans-formingthemoment-equations)
Wedonotwanttostatetheothercompleteexpressionsherebutratherdiscusssomequalitativefeaturesoftheresults.Letusinthefollowingfocusouratten-tiononthemean-square-displacement.Wepresentnumericalresultsforsomeinterestinglimitingcasesanddiscusstheseresults.Thefollowingresultswereobtainedforδ=0.5andK=1.Inﬁgure(5)weseetheMSDoscillatingaroundtheasymptotics∼t1
2.Theboundingpotentialthusleadstosubdiﬀusivebehaviour.
Thenextspecialcaseisω0=0,i.e.noexternalpotentialispresent.Inﬁgure6weseethefrom(20)expectedsubdiﬀusivebehaviour.Letusnextconsiderthecasewhenthemotionbetweenthekicksisover-damped,i.e.ω02<0.Weseefromﬁgure(7)anexponentialdecayoftheMSD.Theparticlelocalizesattheattractingcenter.9
0
2
4
6
8
10
t
0.0
0.1
0.2
0.3
0.4
<x2>
Figure6:ω0=0,η=2
0
2
4
6
8
10
t
0
0.005
0.01
0.015
0.02
<x2>
Figure7:Overdampedmotionbetweenthekicks,η=4,ω0=210
0
1
2
3
4
t
0
0.001
0.002
0.003
0.004
0.005
<x2>
Figure8:Dampedoscillationsbetweenthekicks,η=1,ω0=10Finallyweconsiderthecaseω02>0.Thiscorrespondsintheclassicalpicturetodampedoscillatingmotionbetweenthekicks.ThebehaviouroftheMSDisapparentlysimilartothebehaviouroftheclas-sicalpositionx(t)oftheHarmonicOscillator.Itisquiteinterestinghowevertoseethatforη=0aswellasforω0=0theMSDshowssubdiﬀusivebe-haviour,whilewhenbotheﬀectsareactivetheMSDexponentiallydecaysandsolocalizes.Inthiscontextwewanttomentionthatthelong-timebehaviourofthesecond-ordermomentsfortheKramersequationproposedbyBarkaietal.iscalcuableevenforγ6=0andleadstoloacalizationevenifη=0.AstheKramersequationofMetzleretal.describessubdiﬀusion,localizationishereobtainedtoo.
Summarizingtheboundingforceaswellastheadditionaldampingturnsthesuperdiﬀusive-ballisticbehaviouroftheforce-freecaseintothesubdiﬀusiveregime.Ifbotheﬀectsarepresenttheparticlelocalizeseventuallyatthecenterofattraction.
6Conclusions
WehaveinvestigatedthegeneralizedKramersequationproposedin[10]foraparticlemovinginaaharmonicpotentialundertheinﬂuenceofdamping.WepresentedtheappropriateKramersequation.Forafractionaltimeevolu-tionkernelwediscussedlimitingcasesforthebehaviourofthemoments,i.e.ω0=0,t∼∞andthegeneralcaseforγ=0.Forω0=0weobtainedatransitionfromthesuperdiﬀusive-ballistictothesubdiﬀusiveregime.Therateofconvergencedependsonthevalueoftheadditionaldamping.
Forγ=0wecouldobtainaclosedexpressionforthemeansquaredisplace-ment.Wepresentednumericalresultsforthemeansquaredisplacementanddiscussedqualitativefeatures.Inthiscasetheharmonicpotentialaswellastheadditionaldampingshifttheballistic-superdiﬀusivebehaviour,presentintheforce-freecase,tothesubdiﬀusiveregime.Forω0=0thisalsoholdstrueforthegeneralcase.Interestinglyonlybotheﬀectstogetherleadtoalocalisationoftheparticle.11
References
[1]E.W.MontrollandG.Weiss,J.Math.Phys.6,167(1965)[2]H.ScherandE.Montroll,Phys.Rev.B12,2455(1975)[3]W.G.Gl¨ockle,andT.F.Nonnenmacher,Macromolecules68,46(1995)[4]G.M.Zaslavsky,M.EdelmanandB.A.Niyasov,Chaos7,159(1997)[5]D.del-Castillo-Negrete,B.A.CarrerasandV.E.Lynch,Phys.Rev.Lett.94,065003(2005)[6]R.Sanchez,B.Ph.vanMilligenB.A.Carreras,PhysicsofPlasmas12,056105(2005)[7]R.Balescu,AspectsofAnomalousTransportinPlasmas(IoPPublishing,Bristol,2005)[8]D.Brockmann,L.Hufnagel,andT.Geisel,Nature439,462-465(2006)[9]R.MetzlerandJ.Klafter,Phys.Rep.339,1(2000),R.MetzlerandJ.Klafter,J.Phys.A:Math.Gen.37,R161(2004)[10]R.Friedrich,F.Jenko,A.BauleandS.EulePhys.Rev.Lett.96,230601(2006)[11]R.Friedrich,F.Jenko,A.Baule,andS.Eule,Phys.Rev.E74,041103(2006)[12]E.BarkaiandR.Silbey,J.Chem.Phys.B104,3866(2000)[13]R.MetzlerandJ.Klafter,J.Chem.Phys.B104,3851(2000)[14]R.MetzlerandI.M.Sokolov,Europhys.Lett.,58,(2002)[15]W.T.Coﬀey,Y.P.KalmykovandJ.T.Waldron,TheLangevinEquation(WorldScientiﬁc,Singapore,2004)[16]R.Friedrich,S.EuleandF.Jenko,LangevinApproachtoFractionalDif-fusionEquationsincludingInertialEﬀects,tobepublished[17]R.Metzler,J.KlafterandI.M.Sokolov,Phys.Rev.E58,1621(1998)[18]I.M.SokolovandJ.Klafter,Chaos,SolitonsandFractals34(2007)[19]H.A.Kramers,Physica7,284(1940)[20]H.Risken,TheFokker-PlanckEquation(Springer,Berlin,1989)[21]N.G.vanKampen,StochasticProcessesinPhysicsandChemistry(North-Holland,Amsterdam,1981)12

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

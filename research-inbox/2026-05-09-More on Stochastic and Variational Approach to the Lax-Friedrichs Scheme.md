---
source: "https://arxiv.org/abs/1210.2178v2"
title: "More on Stochastic and Variational Approach to the Lax-Friedrichs Scheme"
author: "Kohei Soga"
year: "2012"
publication: "arXiv preprint / math.NA"
download: "https://arxiv.org/pdf/1210.2178v2"
pdf: "https://arxiv.org/pdf/1210.2178v2"
captured_at: "2026-05-09T12:02:12Z"
updated_at: "2026-05-09T12:02:12Z"
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

# More on Stochastic and Variational Approach to the Lax-Friedrichs Scheme

- 著者: Kohei Soga
- 年: 2012
- 掲載情報: arXiv preprint / math.NA
- 情報源: [arxiv](https://arxiv.org/abs/1210.2178v2)
- ダウンロード: https://arxiv.org/pdf/1210.2178v2
- PDF: https://arxiv.org/pdf/1210.2178v2

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

A stochastic and variational aspect of the Lax-Friedrichs scheme was applied to hyperbolic scalar conservation laws by Soga [arXiv: 1205.2167v1]. The results for the Lax-Friedrichs scheme are extended here to show its time-global stability, the large-time behavior, and error estimates. The proofs essentially rely on the calculus of variations in the Lax-Friedrichs scheme and on the theory of viscosity solutions of Hamilton-Jacobi equations corresponding to the hyperbolic scalar conservation laws. Also provided are basic facts that are useful in the numerical analysis and simulation of the weak Kolmogorov-Arnold-Moser (KAM) theory. As one application, a finite difference approximation to KAM tori is rigorously treated.

## PDF Text

arXiv:1210.2178v2 [math.NA] 16 Apr 2013
MoreonStochasticandVariationalApproachtotheLax-FriedrichsSchemeKoheiSoga∗AbstractAstochasticandvariationalaspectoftheLax-FriedrichsschemewasappliedtohyperbolicscalarconservationlawsbySoga[arXiv:1205.2167v1].TheresultsfortheLax-Friedrichsschemeareextendedheretoshowitstime-globalstability,thelarge-timebehavior,anderrorestimates.TheproofsessentiallyrelyonthecalculusofvariationsintheLax-Friedrichsschemeandonthetheoryofviscos-itysolutionsofHamilton-Jacobiequationscorrespondingtothehyperbolicscalarconservationlaws.Alsoprovidedarebasicfactsthatareusefulinthenumeri-calanalysisandsimulationoftheweakKolmogorov-Arnold-Moser(KAM)theory.Asoneapplication,aﬁnitediﬀerenceapproximationtoKAMtoriisrigorouslytreated.
Keywords:Lax-Friedrichsscheme;scalarconservationlaw;Hamilton-Jacobiequa-tion;calculusofvariations;randomwalk;weakKAMtheory
AMSsubjectclassiﬁcations:65M06;35L65;49L25;60G50;37J501IntroductionΩWeinvestigatetheLax-Friedrichsschemeappliedtoinitialvalueproblemsofhyperbolicscalarconservationlawswithaconstantc,ut+H(x,t,c+u)x=0.(1.1)ΩThereisavastliteratureonthestabilityandconvergenceofthescheme.ThestandardtechniqueisbasedontheL1-frameworkwithaprioriestimatesandthecompactnessoffunctionsofboundedvariation,wheremesh-sizeindependentboundednessofboththediﬀerencesolutionsandtheirtotalvariationmustbeveriﬁed.SincetheLax-Friedrichsschemeisverysimple,detailsofapproximationcanbesuccessfullyanalyzed,particularlyinthecaseofaﬂuxfunctionwiththesimpleformH(x,t,p)=H(p).Wereferto[6],[19],[23],andthestudiescitedtherein.However,inthecaseofageneralﬂuxfunctiondependingonbothxandt,theproblembecomesfarmorediﬃcultandoftenrequiresundesirableassumptions.Theresultsofthegeneralcaseﬁrstappearedin[18],wherestabilityandL1-convergenceareprovedwitharestrictedtimeintervalthatisdetermined
∗DepartmentofPureandAppliedMathematics,WasedaUniversity,Tokyo169-8555,Japan(kohei-math@toki.waseda.jp).1
bythegrowthofH(x,t,p)withrespecttop.In[17],time-globalstabilityandL1-convergencewithinarbitrarytimeintervalsareprovedforaﬂuxfunctionoftheformH(x,t,p)=f(p)+F(x,t)intheperiodicsetting,withmanydetailsofthelarge-timebehavioroftheLax-Friedrichsscheme.Still,itseemsverydiﬃculttoobtainresultssimilartothosein[17]formoregeneralﬂuxfunctionsbythestandardapproachbasedontheL1-framework.Recently,astochasticandvariationalapproachtotheLax-Friedrichsschemewasannounced[22].Stabilityandconvergencewereprovedonthebasisof1)thelawoflargenumbersinthehyperbolicscalinglimitofrandomwalks,and2)thecalculusofvariationsinthetheoryofviscositysolutionsoftheHamilton-Jacobiequationswithconstantscandh(c),vt+H(x,t,c+vx)=h(c).(1.2)ΩThisisaﬁnitediﬀerenceversionofthestochasticandvariationalapproachtothevan-ishingviscositymethodin[11].Nowwebrieﬂyreviewthestochasticandvariationalapproachin[22].Considerinitialvalueproblemsoftheinviscidhyperbolicscalarcon-servationlawandthecorrespondingHamilton-Jacobiequation

ut+H(x,t,c+u)x=0inT×(0,T],u(x,0)=u0(x)∈L∞(T)onT,ZTu0(x)dx=0,ku0kL∞≤r,(1.3)(vt+H(x,t,c+vx)=h(c)inT×(0,T],v(x,0)=v0(x)∈Lip(T)onT,kv0xkL∞≤r,(1.4)Ωwherec∈[c0,c1]isavaryingparameter,T:=R/Zisthestandardtorus,h(c)isacontinuousfunction,andr>0isaconstant.WearbitrarilyﬁxT,r,and[c0,c1].Notethat(1.3)and(1.4)areequivalentinthesensethattheentropysolutionuorviscositysolutionvisderivedfromtheotherifu0=v0x.Inparticular,wehaveu=vx(e.g.,see[1]).Hereinafterweassumethatu0=v0x.TheﬂuxfunctionHisassumedtosatisfythefollowing:(A1)H(x,t,p):T2×R→R,C2(A2)Hpp>0(A3)lim|p|→+∞H(x,t,p)
|p|=+∞.From(A1)–(A3)weobtaintheLegendretransformL(x,t,ξ)ofH(x,t,·),whichisgivenbyL(x,t,ξ)=supp∈R{ξp−H(x,t,p)}andsatisﬁes(A1)0L(x,t,ξ):T2×R→R,C2(A2)0Lξξ>0(A3)0lim|ξ|→+∞L(x,t,ξ)
|ξ|=+∞.Theﬁnalassumptionisthefollowing:(A4)Thereexistsα>0suchthat|Lx|≤α(|L|+1).Notethat(A4)impliescompletenessoftheEuler-LagrangeﬂowgeneratedbyLandHamiltonianﬂowgeneratedbyH.2
Wediscretizetheequationin(1.3)bytheLax-Friedrichsschemeasfollows:uk+1m+1−(uk m+uk m+2)
2
∆t+H(xm+2,tk,c+uk m+2)−H(xm,tk,c+uk m)
2∆x=0.(1.5)ΩWealsodiscretizetheequationin(1.4)bythefollowingscheme:vk+1m−(vkm−1+vkm+1)
2
∆t+H(xm,tk,c+vkm+1−vkm−1
2∆x)=h(c).(1.6)ΩNotethat(1.5)and(1.6)arealsoequivalentinthesensethatuk morvkm+1isderivedfromtheother.Inparticular,wehaveuk m=vkm+1−vkm−1
2∆x,whichisanimportantrelationinthispaper.Inthestochasticandvariationalapproach,thestochasticcomesfromthenumericalviscosityintrinsicto(1.5)and(1.6),whilethevariationalcomesfromthevariationalstructuresofHamilton-Jacobiequations.Thestochasticandvariationalapproachin[22]ledtoseveralresults:(1)Stochasticandvariationalrepresentationformulas(valuefunctions)forvkm+1anduk mwereobtained.(2)StabilityoftheLax-FriedrichsschemeuptoarbitraryT>0wasderivedbyvari-ationaltechniques.(3)Pointwiseconvergenceofuk mtou=vxwasprovedalmosteverywhere.Inpartic-ular,thisyieldeduniformconvergence,exceptforneighborhoodsofshockswitharbitrarilysmallmeasure.(4)Uniformconvergenceofvkm+1tovwithanerrorO(√
∆x)wasprovedfromastochas-ticviewpoint.(5)Randomwalksplayedaroleascharacteristiccurvesofthediﬀerenceequations,whichconvergedtothegenuinecharacteristiccurvesof(1.1)and(1.2).ThepurposeofthispaperistoshowfurtherresultsfortheLax-Friedrichsschemeonthebasisof(1)–(5)undertheassumptions(A1)–(A4)withtechniquesfromthetheoryofviscositysolutionsofHamilton-Jacobiequations.WereferonlytotheresultsfortheLax-Friedrichsscheme,butsimilarresultsforotherﬁnitediﬀerenceschemeswithnumericalviscosity(e.g.,theupwind/downwindscheme)areavailableaswell.Themainresultsareonthefollowing:1.Time-globalstabilityoftheLax-Friedrichsschemewithaﬁxedmeshsize.2.Errorestimatesforentropysolutions.ΩItisprovedthatgenuineentropysolutionsatt=1areuniformlybounded,regardlessofthemagnitudesoftheinitialdata.Sincethegenuinesolutionsarewellapproximatedbythediﬀerenceentropysolutionsforsmallmeshsizes,thediﬀerenceentropysolutionsatt=1arealsouniformlybounded.Duetotheperiodicsetting,iterationofthetime-1analysisyieldstime-globalproperties.Combiningthesefacts,weobtaintime-global3
stabilityoftheLax-Friedrichsscheme.Asaresult,wecanshowthatthelarge-timebehavioroftheLax-Friedrichsschemeissuchthatanysolutionsassociatedwitheachcfallintothetimeperiodicstateuniquelydeterminedbyeachc.Thismeansthatforeachcweobtaintheuniquespace-timeperiodicdiﬀerenceentropysolutionandtheunique(uptoaconstant)space-timeperiodicdiﬀerenceviscositysolution.TheseapproximatethegenuineZ2-periodicentropy(resp.viscosity)solutionof(1.1)(resp.(1.2)).Fortheperiodicstates,wenaturallyhavethenotionoftheeﬀectiveHamiltonianforthediﬀerenceHamilton-Jacobiequation(1.6).WerevealitspropertiesandprovethatitconvergestotheeﬀectiveHamiltonianfortheexactequation(1.2)withanerrorestimateofO(√
∆x).ItisknownthattheoptimalestimateoftheL1-errorbetweenuk manduisO(√
∆x)inthecaseofH(x,t,p)=H(p)[19].TheupperboundO(√
∆x)isduetopropertiesoffunctionsofboundedvariation[14].Itisnotclearwhethertheresultin[14]isapplicabletothecaseofourgeneralﬂuxfunctions.Throughadiﬀerentapproach,weobtainanL1-errorestimateofO(∆x1
4).ThiserrorestimateisbasedonO(√
∆x),whicharisesastheerrorbetweenrandomwalksandtheirspace-timecontinuouslimitunderhyperbolicscaling(i.e.,abackwardcharacteristiccurve).Foratechnicalreason,welosetheexponent1/4inthecaseofthegeneralﬂuxfunctionH(x,t,p).Inaddition,weshowthatifthegenuineentropysolutionisLipschitz,thenaC0-errorestimateofO(∆x1
4)isavailable.Unlikethecaseforinitialvalueproblems,itischallengingtoshowconvergenceoffullsequencesandestimatetheerrorforZ2-periodicentropy(resp.viscosity)solutionsof(1.1)(resp.(1.2)),becausetheuniquenessofsuchgenuineZ2-periodicsolutionswithrespecttocisnotvalidingeneral.However,wecanmanagethespecialcaseinwhichagenuineZ2-periodicentropysolution¯uwithsomecisC1andthedynamicsofitscharacteristiccurvesC∗(s):=(q(s)mod1,smod1)areC1-conjugatetothedynamicsofthelinearﬂowonT2withaDiophantinerotationvector.Suchasolution¯uisknownasaKAMtorusinHamiltoniandynamics(e.g.,see[13],[16],[20]).WeshowaC0-errorestimatedependingontheDiophantinenatureoftherotationvector,whichisarigorousresultonﬁnitediﬀerenceapproximationofKAMtori.OurproofisbasedonthefactthatoneorbitofthelinearﬂowonT2withaDiophantinerotationvectorisergodiconT2andhencesoiseachC∗(s).Finally,wenotethatourmotivationcomesnotonlyfromtheviewpointofPDEsincontinuummechanicsbutalsofromtherecenttheoryofLagrangianandHamiltoniandynamicsthatiscalledtheAubry-MathertheoryortheweakKAMtheory[9],[8],[13].Ourperiodicsettingisstandard,andZ2-periodicentropy(resp.viscosity)solutionsof(1.1)(resp.(1.2))andtheeﬀectiveHamiltonianplaycentralrolesintheweakKAMtheory.TheresultsofthispaperprovidebasictoolsfornumericalanalysisoftheweakKAMtheorythroughﬁnitediﬀerenceapproximation.Weremarkthatfromthestand-pointofaccuracyitisbettertoapproximateentropysolutionsandcharacteristiccurvesaswellasviscositysolutions,becausethecentralobjectsintheweakKAMtheory,suchasKAMtori,Aubry-Mathersets,eﬀectiveHamiltonians,andcalibratedcurves,areobtainedfromthederivativesofviscositysolutionsorentropysolutions.The“deriva-tives”ofnumericalviscositysolutionsobtainedthroughaschemethathasnorelationtoentropysolutionsarenotaccurateingeneral.Somedevelopmentsinﬁnitediﬀer-enceapproximationmethodsandnumericalsimulationsfortheweakKAMtheoryare4
foundin[17].However,theresultsaremathematicallyrestrictedbytheabsenceofthestochasticandvariationalapproachtotheLax-Friedrichsscheme.Wealsopointto[2]and[13]forresultsonsmoothapproximationmethodsfortheweakKAMtheorybasedonthevanishingviscositymethod.Inparticular,[2]successfullyappliesthestochasticandvariationalapproachtothevanishingviscositymethodgivenin[11],wherethegen-uinecharacteristiccurvesareapproximatedbysolutionsofstochasticODEswiththestandardBrownianmotion.Theadvantageofourstochasticandvariationalapproachisthatstructuresandprop-ertiessimilartothoseoftheexactequations(1.1)and(1.2)areavailableinthemostcommonﬁnitediﬀerenceschemes,whichprovidesmuchmoreinformationontheschemes.Inparticular,wecantracegenuinecharacteristiccurvesbymeansofrandomwalks.ThisenablesfurtherdevelopmentofﬁnitediﬀerenceapproximationmethodsfortheclassicalandweakKAMtheories.
2PreliminaryResultsΩInthissection,westateseveralimportantpreliminaryresults.
2.1EntropySolutionandViscositySolutionΩItiswellknownthattheviscositysolutionvof(1.4)isLipschitzandischaracterizedbythecalculusofvariations.Thevalueofvateachpoint(x,t)∈T×(0,T],T∈(0,∞),isgivenbyv(x,t)=infγ∈AC,γ(t)=xZt0L(c)(γ(s),s,γ0(s))ds+v0(γ(0))+h(c)t,(2.1)ΩwhereACisthefamilyofabsolutelycontinuouscurvesγ:[0,t]→TandL(c)(x,t,ξ):=L(x,t,ξ)−cξistheLegendretransformofH(x,t,c+·).Wecanﬁndaminimizingcurveγ∗of(2.1)thatisabackwardcharacteristiccurveof(1.1)and(1.2)aswellasaC2-solutionoftheEuler-LagrangeequationgeneratedbytheLagrangianL(c).Oneachminimizingcurve,visdiﬀerentiablewithrespecttox:vx(γ∗(s),s)=L(c)ξ(γ∗(s),s,γ∗0(s))for0<s<t.(2.2)ΩWesaythatapoint(x,t)isaregularpointofv,orregular,ifvx(x,t)exists.SincevisLipschitz,almosteverypointisregular.Inparticular,if(x,t)isregular,theminimizingcurveγ∗for(2.1)isuniqueand(2.2)holdsfors=t.Usually,theentropysolutionuof(1.3)isdeﬁnedasanelementofC0((0,T];L1(T)).Herewealwaystaketherepresentativeelementgivenbyvx,whichisstilldenotedbyu.If(x,t)isregularandγ∗istheuniqueminimizingcurveforv(x,t),thevalueoftheentropysolutionu=vxatthepoint(x,t)isgivenbyu(x,t)=Zt0L(c)x(γ∗(s),s,γ∗0(s))ds+u0(γ∗(0)),5
whereu0isassumedtoberarefaction-free,esssupx6=yu0(x)−u0(y)
x−y≤MforsomeM>0(i.e.,aone-sidedLipschitzcondition),orequivalentlyv0issemiconcave,v0(x+h)+v0(x−h)−2v0(x)≤Mh2forallx,h.Otherwise,u0(γ∗(0))mustbereplacedwithL(c)ξ(γ∗(0),0,γ∗0(0)).Inparticular,foranyτ∈[0,t)wehaveu(x,t)=ZtτL(c)x(γ∗(s),s,γ∗0(s))ds+L(c)ξ(γ∗(τ),τ,γ∗0(τ)).Formoredetailssee,e.g.,[1]or[5].Weintroducethesolutionoperatorsof(1.3)and(1.4)asfollows:φt:L∞
r,0(T)3u07→u(·,t)∈L∞(T),ψt:Lipr(T)3v07→v(·,t)∈Lip(T),whereL∞
r,0(T)isthesetofallfunctionsu0∈L∞(T)withku0kL∞≤randRTu0dx=0,whileLipr(T)isthesetofallLipschitzfunctionsonTwithaLipschitzconstantboundedbyr.Whenwespecifythevalueofc,wewriteφt(·;c),ψt∆(·;c),u(c),v(c).Wewouldliketoproveaprioriboundednessofu(x,t)=vx(x,t).Thisiscloselyrelatedtoaprioricompactnessofminimizersfor(2.1).WeremarkthataprioricompactnessofminimizersplaysanimportantroleintheAubry-MathertheoryandtheweakKAMtheory,anddetailsareknownformoregeneralsettings(e.g.,[15],[12]).Thebasicassumptionsforthisare(A1)0–(A3)0andcompletenessoftheEuler-Lagrangeﬂow.Hereweadopt(A4),whichisstrongerthanthecompletenessassumption.Weneedthistoobtaincompactnessofminimizersforourstochasticandvariationalproblems,whichdonotsatisfytheEuler-LagrangeequationgeneratedbyL(c).Inordertoprovideaself-containedtreatment,wegivebriefproofsbymodifyingSection4.1of[10].Proposition2.1.Foreacht∈(0,T],thereexistsaconstantβ1(t)>0(independentofr,c∈[c0,c1],andtheinitialdatav0,u0)forwhichkφt(u0;c)kL∞≤β1(t),kψt(v0;c)xkL∞≤β1(t).Proof.Fixt∈(0,T].If(x,t)isregular,then(2.2)holdsfors=t.Thus,itissuﬃcienttoestimateL(c)ξ(γ∗(t),t,γ∗0(t))foreachminimizingcurveγ∗of(2.1).Wenowpreparetwolemmas.
Lemma2.2.Letγ∗beaminimizingcurveforv(x,t).Sety:=γ∗(0).Then,γ∗attainsinfγ∈AC,γ(t)=x,γ(0)=yZt0L(c)(γ(s),s,γ0(s))ds.Proof.Ifnot,thereexistsγ]suchthatZt0L(c)(γ](s),s,γ]0(s))ds<Zt0L(c)(γ∗(s),s,γ∗0(s))ds.6
Sincev0(γ](0))=y=v0(γ∗(0)),wehaveZt0L(c)(γ](s),s,γ]0(s))ds+v0(γ](0))<Zt0L(c)(γ∗(s),s,γ∗0(s))ds+v0(γ∗(0)).Therefore,γ∗isnotaminimizingcurveforv(x,t),whichisacontradiction.
Wedeﬁnethefollowingset:Γ(t):=γ(c)|γ(c)attainsinfγ(t)=x,γ(0)=yZt0L(c)(γ(s),s,γ0(s))ds,x,y∈T,c∈[c0,c1].ByLemma2.2,anyminimizingcurveγ∗forv(x,t),x∈T,belongstoΓ(t).(Actually,weshouldtakeγ∗mod1,butthisisnotimportantduetotheperiodicsetting.)Lemma2.3.1.ThereexistsaconstantC1(t)>0suchthatforanyx,y∈TwehaveaC1-curveγthatsatisﬁesγ(t)=x,γ(0)=y,Zt0L(c)(γ(s),s,γ0(s))ds≤C1(t).Inparticular,anyγ(c)∈Γ(t)satisﬁesZt0L(c)(γ(c)(s),s,γ(c)0(s))ds≤C1(t).2.ThereexistsaconstantC2(t)>0suchthatforanyγ(c)∈Γ(t)wehaveτ∈(0,t)thatsatisﬁes|γ(c)0(τ)|≤C2(t).3.ThereexistsaconstantC3(t)>0suchthatforanyγ(c)∈Γ(t)wehave|L(c)ξ(γ(c)(s),s,γ(c)0(s))|≤C3(t),s∈[0,t].Proof.1.Considerγ(s):=x+x−y t(s−t).Since|x−y|≤1,wehave|γ0(s)|≤t−1.Therefore,weobtainZt0L(c)(γ(s),s,γ0(s))ds≤supx,s∈T,|ξ|≤t−1,c∈[c0,c1]|L(c)(x,s,ξ)|t.SetC1(t):=supx,s∈T,|ξ|≤t−1,c∈[c0,c1]|L(c)(x,s,ξ)|tandClaim1isproved.2.DuetoClaim1andtheminimizingpropertyofγ(c),wehaveτ∈(0,t)thatsatisﬁesC1(t)≥Zt0L(c)(γ(c)(s),s,γ(c)0(s))ds=L(c)(γ(c)(τ),τ,γ(c)0(τ))t.By(A3),|γ(c)0(τ)|mustbeboundedbyaconstantC2(t)independentofγ(c)∈Γ(t).3.Notethatγ(c)isaC2-solutionofthefollowingEuler-LagrangeequationgeneratedbyL(c):d dtL(c)ξ(γ(c)(s),s,γ(c)0(s))=Lx(γ(c)(s),s,γ(c)0(s)).7
Itfollowsfrom(A1)–(A4)thatthereexistsα1forwhich|L(c)x|≤α1(|L(c)|+1)foranyc∈[c0,c1]andthatL∗:=|min{0,infx,s,ξ,cL(c)}|isbounded.Wehaveτ∗∈[0,t],whichattainsthemaximumof|L(c)ξ(γ(c)(s),s,γ(c)0(s))|,0≤s≤t.Supposethatτ∗6=τ,whereτisthevalueinClaim2.Then,|Zτ∗τd dtL(c)ξ(γ(c)(s),s,γ(c)0(s))ds|=|L(c)ξ(γ(c)(τ∗),τ∗,γ(c)0(τ∗))−L(c)ξ(γ(c)(τ),τ,γ(c)0(τ))|≤Zt0|L(c)x(γ(c)(s),s,γ(c)0(s))|ds≤Zt0α1(1+|L(c)(γ(c)(s),s,γ(c)0(s))|)ds≤α1Zt01+(L(c)(γ(c)(s),s,γ(c)0(s))+L∗)+L∗ds=α1(2L∗+1)t+α1Zt0L(c)(γ(c)(s),s,γ(c)0(s))ds≤α1(2L∗+1)t+α1C1(t).Therefore,settingC3(t):=α1(2L∗+1)t+α1C1(t)+supx,s∈T,|ξ|≤C2(t),c∈[c0,c1]|L(c)ξ(x,s,ξ)|,for0≤s≤tweobtain|L(c)ξ(γ(c)(s),s,γ(c)0(s))|≤|L(c)ξ(γ(c)(τ∗),τ∗,γ(c)0(τ∗))|≤C3(t).Thecaseτ∗=τisincludedbytheaboveinequality.
Since(x,t)isregularforalmosteveryx∈Twitheachﬁxedtandvx(x,t)=L(c)ξ(γ∗(t),t,γ∗0(t))holdsforalmosteveryx∈T,weobtainProposition2.1bysettingβ1(t):=C3(t).
Weshowcontinuityofφt(v0x;c)andψt(v0;c)withrespecttov0andc.Proposition2.4.Fixt∈(0,T].Foreachsequencev0j→v0uniformlyandcj→casj→∞(v0jxisnotnecessarilyconvergent),wehaveψt(v0j;cj)→ψt(v0;c)uniformly,φt(v0jx;cj)→φt(v0x;c)inL1(T)asj→∞.Proof.Bythevariationalrepresentation,wehaveψt(v0;c)(x)=Zt0L(γ∗(s),s,γ∗0(s))−cγ∗0(s)ds+v0(γ∗(0))+h(c)t,ψt(v0j;cj)(x)=Zt0L(γ∗j(s),s,γ∗j0(s))−cjγ∗j0(s)ds+v0j(γ∗j(0))+h(cj)t8
andhenceψt(v0j;cj)(x)−ψt(v0;c)(x)≤Zt0−(cj−c)γ∗0(s)ds+v0j(γ∗(0))−v0(γ∗(0))+(h(cj)−h(c))t,ψt(v0j;cj)(x)−ψt(v0;c)(x)≥Zt0−(cj−c)γ∗j0(s)ds+v0j(γ∗j(0))−v0(γ∗j(0))+(h(cj)−h(c))t.ItfollowsfromClaim3ofLemma2.3thatanyminimizingcurvesforv(x,t)areLipschitzwithacommonLipschitzconstantforallx∈Tandv0∈Lipr(T).Sincehiscontinuous,weconcludethatψt(v0j;cj)→ψt(v0;c)uniformlyasj→∞.Letx∈Tbeacommonregularpointofallψt(v0j;cj),j=1,2,3,....Almosteverypointissuchapoint.Throughavariationaltechnique,weﬁndthatγ∗j→γ∗uniformlyandγ∗j0→γ∗0inL2asj→∞(e.g.,seeLemma3.4in[22]).Notethatforeach0≤τ<twehaveφt(v0x;c)(x)=ψt(v0;c)x(x)=ZtτLx(γ∗(s),s,γ∗0(s))ds+Lξ(γ∗(τ),τ,γ∗0(τ))−c,φt(v0jx;cj)(x)=ψt(v0j;cj)x(x)=ZtτLx(γ∗j(s),s,γ∗j0(s))ds+Lξ(γ∗j(τ),τ,γ∗j0(τ))−cj.Foranyε>0,thereexistsaJsuchthat,ifj≥J,wehavekγ∗j−γ∗kC0≤εandkγ∗j0−γ∗0kL2≤ε√
t.Notethatwehaveτ(dependingonj≥J)suchthat|γ∗j0(τ)−γ∗0(τ)|≤ε.Therefore,weconcludethatφt(v0jx;cj)→φt(v0x;c)pointwisealmosteverywhere.ThisimmediatelyleadstoL1(T)-convergence.
2.2StochasticandVariationalApproachtotheLax-FriedrichsSchemeInthissubsection,westateseveralresultsofthestochasticandvariationalapproachtotheLax-Friedrichsschemethatareshownin[22].LetN,KbenaturalnumberswithN≤K.Themeshsize∆=(∆x,∆t)isdeﬁnedby∆x:=(2N)−1and∆t:=(2K)−1.Wesetλ:=∆t/∆x.Wealsosetxm:=m∆xform∈Zandtk:=k∆tfork=0,1,2,....Forx∈Randt>0,thenotationm(x),k(t)denotestheintegersm,kforwhichx∈[xm,xm+2∆x)andt∈[tk,tk+∆t),respectively.Let(∆xZ)×(∆tZ≥0)bethesetofall(xm,tk),andletGeven⊂(∆xZ)×(∆tZ≥0),Godd⊂(∆xZ)×(∆tZ≥0)bethesetofall(xm,tk)withk=0,1,2,...andm∈Zsuchthatm+kiseven(odd),whichiscalledtheevengrid(oddgrid).Weconsiderthediscretizationof(1.3)bytheLax-FriedrichsschemeinGeven:









uk+1m+1−(uk m+uk m+2)
2
∆t+H(xm+2,tk,c+uk m+2)−H(xm,tk,c+uk m)
2∆x=0,u0
m=u0Ω∆(xm),uk m±2N=uk m,(2.3)9
whereformevenu0Ω∆(x):=1
2∆xZxm+∆xxm−∆xu0(y)dyforx∈[xm−∆x,xm+∆x).(2.4)ΩNotethatX{m|0≤m<2N,m+keven}uk m·2∆xisconservativewithrespecttokandiszeroforu0thathaszeromean.Wealsodiscretize(1.4)inGodd:









vk+1m−(vkm−1+vkm+1)
2
∆t+H(xm,tk,c+vkm+1−vkm−1
2∆x)=h(c),v0m+1=v0∆(xm+1),vkm+1±2N=vkm+1,(2.5)Ωwhere,inadditiontou0=v0x,weassumethatv0∆(x):=v0(−∆x)+Zx−∆xu0Ω∆(y)dy(i.e.,v0∆(xm+1)=v0(xm+1)formeven).(2.6)ΩNotethatu0Ω∆→u0inL1(T)andv0∆→v0uniformlywithkv0∆−v0kC0≤ku0kL∞·2∆x,as∆→0.Weintroducethefollowingdiﬀerenceoperators:Dtwk+1m:=wk+1m−wkm−1+wkm+1
2
∆t,Dxwkm+1:=wkm+1−wkm−1
2∆x.Thetwoproblems(2.3)and(2.5)areequivalentunder(2.4)and(2.6).Inparticular,wehaveDxvkm+1=uk m[22].Letu∆bethestepfunctionderivedfromthesolutionuk mof(2.3);namely,u∆(x,t):=uk mfor(x,t)∈[xm−1,xm+1)×[tk,tk+1).Letv∆bethelinearinterpolationwithrespecttothespacevariablederivedfromthesolutionvkm+1of(2.5);namely,v∆(x,t):=vkm−1+Dxvkm+1·(x−xm−1)for(x,t)∈[xm−1,xm+1)×[tk,tk+1).Weremarkthatv∆(x,·)isastepfunctionforeachﬁxedxandthat(v∆)x=u∆.Weintroducespace-timeinhomogeneousrandomwalksinGodd,whichcorrespondtocharacteristiccurvesof(1.3)and(1.4).Foreachpoint(xn,tl+1)∈Godd,weintroducebackwardrandomwalksγthatstartfromxnattl+1andmoveby±∆xineachbackwardtimestep:γ={γk}k=0,1,...,l+1,γl+1=xn,γk+1−γk=±∆x.Moreprecisely,foreach(xn,tl+1)∈Goddweintroducethefollowing:Xk:={xm|(xm,tk)∈Godd,|xm−xn|≤(l+1−k)∆x}fork≤l+1,G:=[1≤k≤l+1 Xk×{tk}⊂Godd,ξ:G3(xm,tk)7→ξkm∈[−λ−1,λ−1],λ=∆t/∆x,¯Ω¯ρ:G3(xm,tk)7→¯Ω¯ρk m:=1
2−1
2λξkm∈[0,1],¯ρ:G3(xm,tk)7→¯ρk m:=1
2+1
2λξkm∈[0,1],γ:{0,1,2,...,l+1}3k7→γk∈Xk,γl+1=xn,γk+1−γk=±∆x,Ω:thefamilyoftheaboveγ.10
Weregard¯Ω¯ρk m(resp.¯ρk m)astheprobabilityoftransitionfrom(xm,tk)to(xm+∆x,tk−∆t)(resp.from(xm,tk)to(xm−∆x,tk−∆t)).Notethatξisacontrolforrandomwalks,whichplaystheroleofavelocityﬁeldonthegrid.Wedeﬁnethedensityofeachpathγ∈Ωasµ(γ):=Y1≤k≤l+1ρ(γk,γk−1),whereρ(γk,γk−1)=¯Ω¯ρk m(γk)(resp.¯ρk m(γk))ifγk−γk−1=−∆x(resp.∆x).Thedensityµ(·)=µ(·;ξ)yieldsaprobabilitymeasureforΩ;namely,prob(A)=Xγ∈Aµ(γ;ξ)forA⊂Ω.TheexpectationwithrespecttothisprobabilitymeasureisdenotedbyEµ(·;ξ);namely,forarandomvariablef:Ω→RwehaveEµ(·;ξ)[f(γ)]:=Xγ∈Ωµ(γ;ξ)f(γ).Weuseγasthesymbolforrandomwalksorasamplepath.Ifnecessary,wewriteγ=γ(xn,tl+1;ξ)inordertospecifyitsinitialpointandcontrol.Wenowstateanimportantresultonthescalinglimitofinhomogeneousrandomwalks.Letη(γ)={ηk(γ)}k=0,1,2,...,l+1,γ∈Ωbearandomvariablethatisinducedbyarandomwalkγ=γ(xn,tl+1;ξ)andisdeﬁnedbyηl+1:=γl+1,ηk(γ):=γl+1−Xk<k0≤l+1ξ(γk0,tk0)∆tfor0≤k≤l.Proposition2.5.([21])Set˜σk:=Eµ(·;ξ)[|γk−ηk(γ)|2]and˜dk:=Eµ(·;ξ)[|γk−ηk(γ)|]for0≤k≤l+1.Then,wehave(˜dk)2≤˜σk≤tl+1−tk
λ∆x.Ifwetakethehyperbolicscalinglimit,inwhich∆=(∆x,∆t)→0under0<λ0≤λ=∆t/∆x<λ1,then˜dkand√
˜σkalwaystendtozerowithO(√
∆x).Notethatthevariancedoesnotnecessarilydosoforinhomogeneousrandomwalks.Wereferto[21]formoredetailsofthehyperbolicscalinglimitofinhomogeneousrandomwalks.Notethatwealwaystakethelimit∆→0underhyperbolicscaling.NowwestateresultsforthestochasticandvariationalapproachtotheLax-Friedrichsscheme.
Theorem2.6([22]).Thereexistsλ1>0(dependingonT,[c0,c1],andr)suchthatforanysmall∆=(∆x,∆t)withλ=∆t/∆x<λ1wehavethefollowing:1.TheexpectationEµ(·;ξ)hX0<k≤l+1L(c)(γk,tk−1,ξkm(γk))∆t+v0∆(γ0)i+h(c)tl+111
whichisgivenbyγ=γ(xn,tl+1;ξ),hasaninﬁmumwithrespecttoξ:G→[−λ−1,λ−1]foreachn∈Zand0<l+1≤k(T).Theinﬁmumisattainedbytheξ∗thatsatisﬁes|ξ∗|≤λ−11<λ−1.2.Foreachn∈Zand0<l+1≤k(T)thesolutionof(2.5)satisﬁesvl+1n=infξEµ(·;ξ)hX0<k≤l+1L(c)(γk,tk−1,ξkm(γk))∆t+v0∆(γ0)i+h(c)tl+1.3.Foreachvl+1ntheminimizingvelocityﬁeldξ∗isuniqueandinGsatisﬁesL(c)ξ(xm,tk,ξ∗k+1m)=Dxvkm+1(⇔ξ∗k+1m=Hp(xm,tk,c+Dxvkm+1)).4.Letξ∗(resp.˜ξ∗)betheminimizingvelocityﬁeldforvl+1n(resp.vl+1n+2).Letγ=γ(xn,tl+1;ξ∗)andµ(·;ξ∗)(resp.˜γ=γ(xn+2,tl+1;˜ξ∗)and˜µ(·;˜ξ∗))betheminimizingrandomwalkanditsprobabilitymeasuregeneratedbyξ∗(resp.˜ξ∗).Then,ul+1n+1=Dxvl+1n+2satisﬁesul+1n+1≤Eµ(·;ξ∗)hX0<k≤l+1L(c)x(γk,tk−1,ξ∗k m(γk))∆t+u0Ω∆(γ0+∆x)i+O(∆x),ul+1n+1≥E˜µ(·;˜ξ∗)hX0<k≤l+1L(c)x(˜γk,tk−1,˜ξ∗k m(˜γk))∆t+u0Ω∆(˜γ0−∆x)i+O(∆x),whereO(∆x)standsforanumberof(−θ∆x,θ∆x)withθ>0independentof∆x.Nowwetakethehyperbolicscalinglimit.5.Letvbetheviscositysolutionof(1.4).Then,foreacht∈[0,T]wehavev∆(·,t)→v(·,t)uniformlyonTas∆→0.Inparticular,wehaveanerrorestimate.Thatis,thereexistsβ2>0(independentof∆,c∈[c0,c1],andtheinitialdatav0∈Lipr(T))suchthatsupt∈[0,T]kv∆(·,t)−v(·,t)kC0(T)≤β2√
∆x.6.Let(x,t)∈T×(0,T]bearegularpointandletγ∗:[0,t]→Rbetheminimizingcurveforv(x,t).Let(xn,tl+1)beapointof[x−2∆x,x+2∆x)×[t−∆t,t+∆t)andletγ∆:[0,t]→Rbethelinearinterpolationoftherandomwalkγ=γ(xn,tl+1;ξ∗)givenbytheminimizingvelocityﬁeldξ∗forvl+1n.Then,γ∆→γ∗uniformlyon[0,t]inprobabilityas∆→0.Inparticular,theaverageofγ∆convergesuniformlytoγ∗as∆→0.7.Letu=vxbetheentropysolutionof(1.3).Then,foreachregularpoint(x,t)∈T×[0,T]wehaveu∆(x,t)→u(x,t)as∆→0.Inparticular,u∆convergesuniformlytouon(T×[0,T])\Θ,whereΘisaneigh-borhoodofthesetofpointsofdiscontinuityofuwithanarbitrarilysmallmeasure.12
NotethatClaims1and3givethestabilityconditionoftheLax-Friedrichsscheme,|λHp(xm,tk,c+uk m))|<1,whichiscalledtheCFLcondition.WenextstatefurtherpreliminaryresultsfortheLax-Friedrichsscheme.Thesolutionoperatorsof(2.3)and(2.5)areintroducedasφtΩ∆:L∞
r,0(T)3u07→u∆(·,t)∈L∞(T),ψt∆:Lipr(T)3v07→v∆(·,t)∈Lip(T).Whenwespecifythevalueofc,wewriteφtΩ∆(·;c),ψt∆(·;c),u(c)∆,uk m(c),v(c)∆,vkm+1(c).Notethatweﬁrstobtainthestepfunctionu0Ω∆fromu0with(2.4)andthenwemapu0Ω∆tou∆(·,t)withφtΩ∆.Similarly,weﬁrstobtainthepiecewiselinearfunctionv0∆fromv0with(2.6),inwhichu0=v0x,andthenwemapv0∆tov∆(·,t)withψt∆.Proposition2.7.Fixt∈[0,T].Foreachsequencev0j→v0uniformlyandcj→casj→∞(v0jxisnotnecessarilyconvergent),wehaveψt∆(v0j;cj)→ψt∆(v0;c)uniformly,φtΩ∆(v0jx;cj)→φtΩ∆(v0x;c)inL1(T)asj→∞.Proof.Itissuﬃcienttoshowthatψtl+1∆(v0j;cj)(xn)→ψtl+1∆(v0;c)(xn)uniformlywithrespecttoxnasj→∞.Usingthestochasticandvariationalrepresentation,wehaveψtl+1∆(v0;c)(xn)=Eµ(·;ξ∗)hX0<k≤l+1L(γk,tk−1,ξ∗k m(γk))−cξ∗k m(γk)∆t+v0∆(γ0)i+h(c)tl+1,ψtl+1∆(v0j;cj)(xn)=Eµ(·;ξ∗j)hX0<k≤l+1L(γk,tk−1,ξ∗jk m(γk))−cjξ∗jk m(γk)∆t+v0j∆(γ0)i+h(cj)tl+1,whereξ∗,ξ∗jareminimizingvelocityﬁelds.Hence,bythestochasticandvariationalrepresentationagain,wehaveψtl+1∆(v0j;cj)(xn)−ψtl+1∆(v0;c)(xn)≤Eµ(·;ξ∗)hX0<k≤l+1−(cj−c)ξ∗k m(γk)∆t+v0j∆(γ0)−v0∆(γ0)i+(h(cj)−h(c))tl+1,ψtl+1∆(v0j;cj)(xn)−ψtl+1∆(v0;c)(xn)≥Eµ(·;ξ∗j)hX0<k≤l+1−(cj−c)ξ∗jk m(γk)∆t+v0j∆(γ0)−v0∆(γ0)i+(h(cj)−h(c))tl+1.Sinceξ∗,ξ∗jareuniformlybounded,wehavedemonstratedtheassertion.Thesecondconvergencefollowsfromtheﬁrstoneandthefollowingrelation:φtΩ∆(v0jx;cj)(xm)=ψt∆(v0j;cj)(xm+1)−ψt∆(v0j;cj)(xm−1)
2∆x.
13
Wenowshowdetailsoftheone-sidedLipschitzconditiononuk m,orequivalentlythesemiconcavepropertyofvkm+1;namely,weobtainthe∆-independentupperboundednessofEk∆:=maxmuk m+2−uk m
2∆x=maxmvkm+3+vkm−1−2vkm+1
(2∆x)2.Thisleadstotheentropyconditiononu(·,t)andsemiconcavityofv(·,t).Theone-sidedLipschitzconditiononuk misessentialinthestandardL1-frameworkofdiﬀerenceapproximation,becausethisconditionyields∆-independentboundednessofthetotalvariationofu∆(·,t)andthenL1-convergenceoftheapproximationfollowswiththeaidofthecompactnessoffunctionsofboundedvariation.WeremarkthatTheorem2.6wasprovedindependentlyoftheconditionandwithoutsuchcompactness.Inthesectionsbelow,weusetheone-sidedLipschitzconditiononuk mfordiﬀerentpurposes.Ifweassumethatv0issemiconcave,itiseasytoﬁndanupperboundforEk∆throughthesemiconcavityofvkm+1duetoitsvariationalstructure.However,wewouldliketoavoidthatassumptionandknowaboutthek-dependenceoftheupperbound.Therefore,weuseadirectmethodsimilartothatofLemma2in[18].ThedirectmethodisavailableforarbitraryT>0,becausewealreadyknowbyTheorem2.6thatthediﬀerencesolutionsareboundeduptoT.Weintroducethefollowingnotationwiththeλ1inTheorem2.6:u∗:=supx,t∈T,c∈[c0,c1],|ξ|≤λ−11|L(c)ξ(x,t,ξ)|(notethat|uk m|≤u∗),H∗xx:=supx,t∈T,c∈[c0,c1],|u|≤u∗|Hxx(x,t,c+u)|,H∗xp:=supx,t∈T,c∈[c0,c1],|u|≤u∗|Hxp(x,t,c+u)|,H∗pp:=infx,t∈T,c∈[c0,c1],|u|≤u∗|Hpp(x,t,c+u)|(H∗pp>0dueto(A2)),η:=max{2H∗xp+H∗pp,1
2H∗pp+H∗xx},E∗:=2H∗xp
H∗pp+s
4H∗xp
H∗pp2+2H∗xx
H∗pp.Beforegivingdetails,wesummarizeourstrategyasfollows:WeestimateEk+1∆fromEk∆byusingthediﬀerenceequation.WeﬁndthateachEk+1∆−Ek∆isboundedfromabovebyP(Ek∆),whereP(y)isaconcaveparabolawhosezeropointontheright-handsideisE∗;i.e.,P(y)>0for0≤y<E∗,P(E∗)=0,andP(y)<0fory>E∗(see(2.8)below).Hence,ifEk∆>E∗(resp.Ek∆<E∗),thenEk+1∆decreasesbyatleastP(Ek∆)<0(resp.increasesbyatmostP(Ek∆)>0)andcanremainnearE∗forlargek≤k(T).IfE0∆isverylarge,Ek∆decaysrapidlyatﬁrstinawaysimilartothatofsolutionstow0(s)=−(w(s))2,wherew(s)∼1/s.Proposition2.8.Letλ1>0bethatofTheorem2.6.Supposethat∆=(∆x,∆t)satisﬁesλ=∆t/∆x<λ1,∆t<min{(2η)−1,(E∗H∗pp+2H∗xp)−1},supx,t∈T,c∈[c0,c1],|u|≤u∗λ(|Hp(x,t,c+u)|+H∗xp·2∆x)<1,λ≤1−2H∗xp∆t rH∗pp+(1+H∗pp)∆x.(2.7)
Then,thefollowinghold:1.For1≤k≤k(T)wehaveEk∆=maxmuk m+2−uk m
2∆x≤2eηtk
H∗pp1
tk(tk=k∆t).14
2.IfE0∆≤E∗,wehaveEk∆≤E∗for1≤k≤k(T).3.Ifk>k(η−1),wehaveEk∆≤4eη
H∗pp.4.Ifuk misextendedtok→∞with|uk m|≤u∗,wehavelimsupk→∞Ek∆≤E∗.Proof.UsingthediﬀerenceequationandTaylor’sformula,weobtainanestimateofEk+1∆fromEk∆.Forbrevity,theremaindersinTaylor’sformulaaredenotedbyHpp,Hxx,andHxp,whichsatisfyHpp≥H∗pp,|Hxx|≤H∗xx,|Hxp|≤H∗xp.Setzkm:=uk m+2−uk m.Then,wehavezk+1m+1=zkm+zkm+2
2−∆t
2∆x{H(xm+4,tk,c+uk m+4)−H(xm+2,tk,c+uk m+4)+H(xm+2,tk,c+uk m+4)−H(xm+2,tk,c+uk m+2)+H(xm,tk,c+uk m)−H(xm+2,tk,c+uk m)+H(xm+2,tk,c+uk m)−H(xm+2,tk,c+uk m+2)}=(1
2+λ
2Hp(xm+2,tk,c+uk m+2))zkm+(1
2−λ
2Hp(xm+2,tk,c+uk m+2))zkm+2−∆t
2∆x{(Hx(xm+2,tk,c+uk m+4)−Hx(xm+2,tk,c+uk m))(2∆x)+1
2Hpp·(zkm+2)2+1
2Hpp·(zkm)2+1
2Hxx·(2∆x)2+1
2Hxx·(2∆x)2}={1
2+λ
2Hp(xm+2,tk,c+uk m+2)−λ
2Hxp·2∆x}zkm+{1
2−λ
2Hp(xm+2,tk,c+uk m+2)−λ
2Hxp·2∆x}zkm+2−∆t
2∆x{1
2Hpp·(zkm+2)2+1
2Hpp·(zkm)2+1
2Hxx·(2∆x)2+1
2Hxx·(2∆x)2}.Bytheﬁrstinequalityin(2.7),itfollowsthat{1
2±λ
2Hp(xm+2,tk,c+uk m+2)−λ
2Hxp·2∆x}>0.Hence,setting˜zkm:=max{zkm,zkm+2},weobtainzk+1m+1≤(1−2Hxp∆t)˜zkm+H∗xx·2∆x∆t−H∗pp
2∆t
2∆x(˜zkm)2,zk+1m+1
2∆x≤(1−2Hxp∆t)˜zkm
2∆x+H∗xx∆t−H∗pp
2∆t(˜zkm
2∆x)2.Notethatg(y):=(1−2Hxp∆t)y+H∗xx∆t−(H∗pp
2∆t)y2ismonotonicallyincreasingify≤1−2Hxp∆t
H∗pp∆t.15
Fromthesecondinequalityin(2.7)itfollowsthatλ≤(1−2H∗xp∆t)/(rH∗pp+∆x)andhenceE0∆≤2r
2∆x≤1−2H∗xp∆t
H∗pp∆t≤1−2Hxp∆t
H∗pp∆tforallinitialdatainL∞
r,0(T).SupposethatEk∆≤(1−2H∗xp∆t)/(H∗pp∆t).Then,weobtainEk+1∆≤Ek∆+P(Ek∆),P(y):=−∆t(H∗pp
2y2−2H∗xpy−H∗xx).(2.8)ΩFrom∆t<(E∗H∗pp+2H∗xp)−1itfollowsthatE∗<(1−2H∗xp∆t)/(H∗pp∆t),andfrom∆t<(2η)−1itfollowsthat|y−E∗|≥|P(y)|forall0≤y≤(1−2H∗xp∆t)/(H∗pp∆t).Hence,wehavetwocases:(1)IfEk∆≤E∗,wemayhaveEk+1∆≥Ek∆,butwecertainlyhaveEk+1∆≤Ek∆+P(Ek∆)≤E∗.(2)IfE∗<Ek∆,wehaveEk+1∆<Ek∆.Therefore,wehaveEk+1∆≤(1−2H∗xp∆t)/(H∗pp∆t)and,byinduction,itfollowsthatEk∆≤(1−2H∗xp∆t)/(H∗pp∆t)forall0≤k≤k(T).Thus,(2.8)holdsforall0≤k<k(T).Itisnoweasytoverifythat1)ifE0∆≤E∗,thenEk∆mayincreasebutneverexceedE∗,and2)ifE0∆>E∗,thentheEk∆areboundedfromabovebyamonotonicallydecreasingsequence.Claim4isalsoclear.NowwefollowLemma2in[18].SetVk:=Ek∆+1≥1.Then,by(2.8)wehaveVk+1≤(1+η∆t)Vk−H∗pp
2∆t(Vk)2WesetWk:=(1−η∆t)kVkfork≥0(1−η∆t>0holdssince∆t<(2η)−1).Then,fork≥1wehaveWk+1≤(1−η∆t)(1+η∆t)Wk−H∗pp
2∆t(Wk)2(1−η∆t)−k+1≤Wk−H∗pp
2∆t(Wk)2.Considerw0(t)=−H∗pp
2(w(t))2,withw(0)=w0:=2/(H∗pp∆t).Thesolutionsatisﬁesw(t)=1
H∗pp
2t+1
w0≤2
H∗ppt.WecanshowthatWk≤w(k∆t)fork≥1bynotingthatw(∆t)=1/(H∗pp∆t)andW1=(1−η∆t)(E1∆+1)≤r/∆x+1.Fromthesecondinequalityin(2.7)itfollowsthatλ≤1/{(r+∆x)H∗pp}andhencethatW1≤w(∆t).SupposethatWk≤w(k∆t)forsomek≥1.Then,sinceg(y):=y−H∗pp∆t
2y2ismonotonicallyincreasingfory≤1/(H∗pp∆t),w(k∆t)≤1/(H∗pp∆t),andw00>0,wehaveWk+1≤Wk−H∗pp∆t
2(Wk)2≤w(k∆t)−H∗pp∆t
2(w(k∆t))2=w(k∆t)+∆tw0(k∆t)=w(k∆t+∆t)−1
2w00(k∆t+θ∆t)·(∆t)2≤w((k+1)∆t)(θ∈(0,1)).16
Thus,weobtainEk∆≤(1−η∆t)−k2
H∗ppk∆t≤(1−η∆t)−ηk∆t
η∆t2
H∗ppk∆t≤2eηtk
H∗pp1
tk.Settingf(t):=2eηt
H∗pp1
t,theminimumoffbecomesf(η−1)=2eη
H∗pp,whichisgreaterthanE∗.Therefore,duetoCases(1)and(2),theEk∆areboundedfromabovebyf(tk)(>E∗)fork≤k(η−1)andneverexceedf(η−1−∆t)≤4eη
H∗ppfork>k(η−1).Thisdemonstratestheproposition.
3Time-GlobalStabilityandLarge-TimeBehaviorΩWeprovetime-globalstabilityoftheLax-Friedrichsschemewithaﬁxedmeshsize.Then,weshowthelarge-timebehavioroftheschemeinwhicheachdiﬀerencesolutionfallsintoatimeperiodicstatewithunitperiod.Eachtimeperiodicstatecorrespondstoaspace-timeperiodicdiﬀerencesolution.TherearisesthenotionoftheeﬀectiveHamiltonianof(1.6).
3.1Time-GlobalStabilityΩThemainresultofthissectionisthefollowingtheorem.
Theorem3.1.Thereexistλ1>0andδ>0suchthat,if∆=(∆x,∆t)satisﬁes0<λ0≤λ=∆t/∆x<λ1and∆x≤δ,theLax-Friedrichsschemestartingfromanyu0∈L∞
r,0(T)succeedsuptoanarbitrarytimeindexandsatisﬁestheCFLcondition|Hp(xm,tk,c+uk m)|≤λ−11<λ−1forallm∈Zandk∈Z+.Inordertoprovethistheorem,weneeduniformboundednessofkφ1Ω∆(u0;c)kL∞withrespectto(u0;c)similartothatinProposition2.1.First,weobservethefollowinglemma.
Lemma3.2.Letλ1>0bethatofTheorem2.6andlet∆=(∆x,∆t)besuchthat0<λ0≤λ=∆t/∆x<λ1.Fixt∈(0,T]arbitrarily.Then,foranyε>0thereexistsδ=δ(ε;t)>0suchthat,if∆x≤δ,wehavesupu0∈L∞
r,0(T),c∈[c0,c1]kφtΩ∆(u0;c)−φt(u0;c)kL1(T)≤ε.Proof.Ifnot,thenforsomeε0>0andδj→0asj→∞,wehave∆xj≤δjsuchthatsupu0∈L∞
r,0(T),c∈[c0,c1]kφtΩ∆j(u0;c)−φt(u0;c)kL1(T)>ε0,(3.1)Ωwhere∆j=(∆xj,λ∆xj).Weshowthatforeachjthereexists(u0,c)∈L∞
r,0(T)×[c0,c1]thatattainsthesupremum(3.1)denotedbyb.Let(u0
i,ci)bethesequenceforwhichkφtΩ∆j(u0
i;ci)−φt(u0
i;ci)kL1(T)→basi→∞.Letv0ibeaprimitiveofu0
ithatbelongstoLipr(T)andisboundedbyr.BytheArzela-Ascolitheoremwehaveasubsequenceof17
(v0i,ci),stilldenotedby(v0i,ci),whichconvergesto(v0,c).ByPropositions2.4and2.7,wehaveφtΩ∆j(v0ix;ci)→φtΩ∆j(v0x;c)inL1(T)andφt(v0ix;ci)→φt(v0x;c)inL1(T)asi→∞.Therefore,weobtainkφtΩ∆j(v0x;c)−φt(v0x;c)kL1(T)=b.Let(u0
j,cj)attainthesupremum(3.1).Letv0jbeaprimitiveofu0
jthatbelongstoLipr(T)andisboundedbyr.Wehaveasubsequenceof(v0j,cj),stilldenotedby(v0j,cj),whichconvergesto(v0,c).ItfollowsfromClaim7ofTheorem2.6thatthereexistsδ0>0suchthat,if∆x≤δ0,wehavekφtΩ∆(v0x;c)−φt(v0x;c)kL1(T)<ε0
2.Hence,ε0
2>kφtΩ∆(v0x;c)−φt(v0x;c)kL1(T)≥kφtΩ∆(v0jx;cj)−φt(v0jx;cj)kL1(T)−kφtΩ∆(v0x;c)−φtΩ∆(v0jx;cj)kL1(T)−kφt(v0jx;cj)−φt(v0x;c)kL1(T).ByPropositions2.4and2.7,wehavekφtΩ∆(v0x;c)−φtΩ∆(v0jx;cj)kL1(T)+kφt(v0jx;cj)−φt(v0x;c)kL1(T)≤ε0
2forlargej.Therefore,wehavekφtΩ∆(v0jx;cj)−φt(v0jx;cj)kL1(T)<ε0forany∆x≤δ0,whichisacontradiction.
Next,weseethattheconvergencekφ1Ω∆(u0;c)−φ1(u0;c)kL1(T)→0as∆→0,whichisuniformwithrespectto(u0,c),yieldsuniformboundednessofkφ1Ω∆(u0;c)kL∞withtheaidoftheone-sidedLipschitzcondition.
Proposition3.3.Letλ1>0bethatofTheorem2.6withT=1.Let∆=(∆x,∆t)besuchthat0<λ0≤λ=∆t/∆x<λ1,satisfyingtheconditionsinProposition2.8.Then,thereexistsδ>0suchthat,if∆x≤δ,withβ1ofProposition2.1wehavesupu0∈L∞
r,0(T),c∈[c0,c1]kφ1Ω∆(u0;c)kL∞≤β1(1)+1.Proof.Let0<ε<1besuchthat1−√
ε
3√
ε>2eη/H∗pp≥E2K∆,where2K∆t=1.Withthisεandt=1,wehaveδ>0inLemma3.2.Wetake∆x≤min{δ,√
ε}.ConsiderA:={y∈T||φ1Ω∆(u0;c)(y)−φ1(u0;c)(y)|>√
ε}.Sincekφ1Ω∆(u0;c)−φ1(u0;c)kL1(T)≤ε,wehavemeas[A]≤√
ε.Hence,fory∈Athereexistsx,˜x∈R\Asuchthat0<y−x≤√
εand0<˜x−y≤√
ε.Forx∈T\A,wehave|φ1Ω∆(u0;c)(x)−φ1(u0;c)(x)|≤√
εand|φ1Ω∆(u0;c)(x)|≤|φ1(u0;c)(x)|+√
ε≤β1(t)+1.Consider˜A:={y∈A||φ1Ω∆(u0;c)(y)|>β1(t)+1}.Supposethat˜Aisnotempty.Then,thereexistsxn∈˜A∩(∆xZ)suchthatu2Kn>β1(1)+1(resp.u2Kn<−β1(1)−1).Sincethereexistxm,xm0∈(R\A)∩(∆xZ)suchthat0<xn−xm≤√
ε+2∆x≤3√
εand0<xm0−xn≤√
ε+2∆x≤3√
ε,wehaveu2Kn−u2Km xn−xm>β1(1)+1−(β1(1)+√
ε)
3√
ε=1−√
ε
3√
ε>E2K∆,resp.u2Km0−u2Kn xm0−xn>−(β1(1)+√
ε)−(−β1(1)−1)
3√
ε=1−√
ε
3√
ε>E2K∆.Thesetwoinequalitiescontradicttheone-sidedLipschitzcondition.
18
Remark.Claim7ofTheorem2.6statesthatφ1Ω∆(u0;c)convergestoφ1(u0;c)uniformlyonT\Θas∆→0,whereΘisanarbitrarysmallneighborhoodofshocks.However,wecannotusethisfactforProposition3.3,becauseuniformityoftheconvergencewithrespectto(u0;c)isunveriﬁed.ProofofTheorem3.1.Letλ1>0bethatofTheorem2.6withT=1andr≥β1(1)+1.Letδ>0bethatofProposition3.3.Then,˜u0:=φ1Ω∆(u0;c)belongstoL∞
β1(1)+1,0(T)foranyu0∈L∞
r,0(T).Hence,bythechoiceofλ1,weareguaranteedthatφ1Ω∆(˜u0;c)=φ2Ω∆(u0;c)iswelldeﬁnedandboundedbyβ1(1)+1again.Inthisway,φtΩ∆(u0;c)canbedeﬁnedfort→∞withtheCFLcondition.
3.2Large-TimeBehaviorΩIfwetaker≥β1(1)+1,thenφ1Ω∆(u0;c)belongstoL∞
r,0(T).Therefore,φ1Ω∆(·;c)mapsL∞
r,0(T)intoitself.Wecanﬁndtheﬁxedpointsofthemapforeachc.Inthissubsection,weconsidertheﬁxedpointsandtheirstability,whichmakesclearthelarge-timebehavioroftheLax-Friedrichsscheme.NotethattheLax-FriedrichsschemehasacontractionpropertyundertheCFLcondition.Thatis,for0≤t≤t0wehavekφt0∆(u0;c)−φt0∆(˜u0;c)kL1(T)≤kφtΩ∆(u0;c)−φtΩ∆(˜u0;c)kL1(T).Thiscanbereﬁnedtobecomeastrictcontractionproperty.LetPm;kdenotesummationwithrespectto{m|0≤m<2N,m+keven}foreachﬁxedk,andletkxk1:=P1≤j≤n|xj|forx∈Rn.Proposition3.4.Thefamilyofmaps{φtΩ∆(·;c)}t≥0hasastrictcontractionpropertywithintheunittimeperiod.Thatis,foranytwodistinctinitialdatau0and˜u0,wehavekφt+1∆(u0;c)−φt+1∆(˜u0;c)kL1(T)<kφtΩ∆(u0;c)−φtΩ∆(˜u0;c)kL1(T).Proof.Itissuﬃcienttoshowthatforallk≥0anytwodiﬀerencesolutionsukand˜ukof(2.3)satisfykuk+2K−˜uk+2Kk1<kuk−˜ukk1.Setzkm:=uk m−˜uk mandσkm:=signzkm=1or−1(sign0:=1).Then,kuk−˜ukk1=Pm;k|zkm|=Pm;kσkmzkm.Bythediﬀerenceequationof(2.3),wehaveXm;k|zk+1m+1|=Xm;kσk+1m+1zk+1m+1=Xm;kσk+1m+1n1
2zkm+2(1−λδkm+2)+1
2zkm(1+λδkm)o,whereδkm:=Hp(xm,tk,c+uk m+θkm)withaconstantθkmderivedfromTaylor’sformula.Switchingtheorderofthesummationsabove,weobtainXm;k|zk+1m+1|=Xm;kzkm1
2σk+1m−1(1−λδkm)+1
2σk+1m+1(1+λδkm)=Xm;k|zkm|+Xm;k|zkm|−1+σkm1
2σk+1m−1(1−λδkm)+1
2σk+1m+1(1+λδkm).19
LetRkdenotethesecondsuminthesecondlineoftheaboveequality.WeﬁndthatRk≤0,sinceforeachtermofRkthefactor[]belongstooneoftwocases:(1)Ifσk+1m−1+σk+1m+1=±2,then[]=−1±σkm=0or−2.(2)Ifσk+1m−1+σk+1m+1=0,then[]=−1±λδkm<0duetotheCFLcondition.Sinceukand˜ukeachhavezeromeanandu06=˜u0,thesignofzkmnecessarilychangesandCase(2)occurs.Itseemspossiblethateventhoughukand˜ukaresuch,wemayhaveRk=0;namely,zkm=0foralltheintegersmforwhichCase(2)occurs.However,afterfurtherk∗-timeevolution(k∗<N<2K),Case(2)certainlyoccursandRk+k∗<0,becausesuchzero-pointsdisappearaskincreasesduetothemonotonicityoftheLax-FriedrichsschemeundertheCFLcondition(seealsoRemark2.5in[17]).
Wenowshowthattimeperiodicdiﬀerencesolutionsnotonlyexistbutarestable,whichprovidesthelarge-timebehavioroftheLax-Friedrichsscheme.Theorem3.5.Taker≥β1(1)+1andﬁx∆=(∆x,∆t)sothatTheorems2.6and3.1hold.Then,foreachcthereexistsaﬁxedpoint¯u0Ω∆∈L∞
r,0(T)ofφ1Ω∆(·;c),whichyieldsatimeperiodicdiﬀerencesolutionφtΩ∆(¯u0Ω∆;c).Suchaperiodicsolutionisuniquewithrespecttoc.AnyothersolutionφtΩ∆(u0;c)exponentiallyfallsintotheperiodicstate;namely,thereexistρ∈(0,1)anda>0dependingon∆,butindependentofu0,suchthatkφtΩ∆(u0;c)−φtΩ∆(¯u0Ω∆;c)kL∞≤aρtfort∈N.Proof.Themapφ1Ω∆(·;c)isactuallyamapfromRNtoRN,sincethestepfunctionshaveonlyNdiﬀerentvaluesatmost.LetBrbethesetofallx∈RNwithkxk∞≤r.Ifr≥β1(1)+1,thenthemapφ1Ω∆(·;c)isactuallyamapfromBrtoBr.Therefore,weobtainaﬁxedpointthroughBrouwer’sﬁxedpointtheorem.ByProposition3.4,periodicsolutionsmustbeunique.Exponentialdecaycanbeprovedinthesamewayas(5)ofTheorem2.1in[17].
Remark.Itislikelythatingeneralρbecomesarbitrarilyclosetounityas∆tendstozero.Numericalexperimentsimplysuchapropertyofρ[17].Theuniquenessdoesnotholdfortheexactequation(1.1)ingeneral.Theremayexisttimeperiodicentropysolutionsof(1.1)withtheminimumperiodgreaterthanunity[1].ΩThefollowingtheoremforthediscreteHamilton-JacobiequationisliketheweakKAMtheorem.
Theorem3.6.Taker≥β1(1)+1andﬁx∆=(∆x,∆t)sothatTheorems2.6and3.1hold.Then,foreachcthereexistsaconstant¯h∆(c)∈Rsuchthatifh(c)=¯h∆(c),wehaveaﬁxedpoint¯v0∆∈Lipr(T)ofψ1∆(·;c),whichyieldsatimeperiodicdiﬀerencesolutionψt∆(¯v0∆;c).Suchaperiodicsolutionisuniquewithrespecttocuptoconstants.Anyothersolutionψt∆(v0;c)exponentiallyfallsintoaperiodicstate;namely,fortheρ∈(0,1)anda>0inTheorem3.5andford∈Rdependingon(v0;c)wehavekψt∆(v0;c)−ψt∆(¯v0∆+dI1;c)kC0≤aρtfort∈N,whereI1(x):=1andψt∆(v0+dI1;c)=ψt∆(v0;c)+dI1.Proof.WeimitatetheproofoftheweakKAMtheorem[9].Letuswritev∼wforv,w∈C0(T)ifthereexistsb∈Rsuchthatw=v+bI1.Weintroduceˆv:={w∈C0(T)|w∼v},20
kˆvk:=infw∈ˆvkwkC0(T),ˆC0(T):=C0(T)/∼,andˆLipr(T):=Lipr(T)/∼.FromtheArzela-AscolitheoremitfollowsthatˆLipr(T)isacompactconvexsubsetoftheBanachspaceˆC0(T).Duetothepropertyψt∆(v0+dI1;c)=ψt∆(v0;c)+dI1,wehaveψ1∆(v0;c)∼ψ1∆(w0;c)forallv0,w0∈ˆv0.Hence,themapˆψ1∆(·;c):ˆLipr(T)→ˆLipr(T),ˆψ1∆(ˆv0;c):={w∈Lipr(T)|w∼ψ1∆(v0;c)}(v0∈ˆv0)iswelldeﬁnedandcontinuous.BySchauder’sﬁxedpointtheorem,wehaveaﬁxedpointˆΩ¯v0∆satisfyingˆψ1∆(ˆΩ¯v0∆;c)=ˆΩ¯v0∆.Therefore,wehaveanelement¯v0∆∈ˆΩ¯v0∆andconstantb(c)∈Rsuchthat¯v0∆=ψ1∆(¯v0∆;c)+b(c)I1.Thisrelationmeansthat¯v0∆yieldsatimeperiodicsolutionof(2.5)withh(c)+b(c)insteadofh(c).Notethatψt∆(v0;c)≤ψt∆(˜v0;c)ifv0≤˜v0.Leta0,b0beconstantssuchthatforallx∈Twehave¯v0∆(x)+b0≤v0(x)≤¯v0∆(x)+a0,withatleastonepointattainingtheequalityineachinequality.Then,wehave¯v0∆(x)+b0≤ψ1∆(v0;c)(x)≤¯v0∆(x)+a0forallx∈T.Leta1,b1beconstantssuchthatforallx∈Twehave¯v0∆(x)+b1≤ψ1∆(v0;c)(x)≤¯v0∆(x)+a1,withatleastonepointattainingtheequalityineachinequality.Notethata1≤a0andb1≥b0.Then,wehave¯v0∆(x)+b1≤ψ2∆(v0;c)(x)≤¯v0∆(x)+a1forallx∈T.Inthisway,weobtaintheboundedmonotonesequencesajandbj.Takedsuchthatlimj→∞bj≤d≤limj→∞aj.Then,¯v0∆+dI1andψt∆(v0;c)coincideforatleastonepointandforanyt∈N.Letx0∈Tbesuchthat¯v0∆(x0)+d=ψt∆(v0;c)(x0).Then,forallx∈Tandt∈Nweobtain|ψt∆(v0;c)(x)−ψt∆(¯v0∆+dI1;c)(x)|≤

Zxx0|φtΩ∆(v0x;c)−φtΩ∆(¯u0Ω∆;c)|dy

≤aρt.
Weintroducethemap¯h∆(c):c7→h(c)+b(c),whichistheeﬀectiveHamiltonianofthediﬀerenceHamilton-Jacobiequation(1.6).Weremarkthat¯h∆(c)playsanimportantroleinthenumericalanalysisoftheweakKAMtheory.Hence,itspropertiesaremeaningfultoinvestigate.
3.3EﬀectiveHamiltonianΩBelowisthecharacterizationof¯h∆(c),whichisverysimilartothatoftheeﬀectiveHamiltonian¯h(c)oftheexactHamilton-Jacobiequations(1.2).Wereferto[1]forthecharacterizationof¯h(c).Theorem3.7.1.h(c)=¯h∆(c)istheuniquevalueforwhich(1.6)admitsaspace-timeperiodicdiﬀerencesolution.21
2.¯h∆(c)istheaveragedHamiltonian.Thatis,forthespace-timeperiodicdiﬀerencesolution¯uk mof(1.5)wehave¯h∆(c)=X0≤k<2KXm;kH(xm,tk,c+¯uk m(c))·2∆x∆t.3.Letvl+1n(c)beatime-globalsolutionofthediﬀerenceequationDtvk+1m+H(xm,tk,c+Dxvkm+1)=0.(3.2)
Then,forallnwehaveliml→∞vl+1n(c)
tl+1=−¯h∆(c).4.¯h∆(c)isaconvexC1-function.5.¯h∆(c)uniformlyconvergestotheexacteﬀectiveHamiltonian¯h(c)of(1.2)as∆→0:supc∈[c0,c1]|¯h∆(c)−¯h(c)|≤β3√
∆x.Proof.1.Let˜Ω¯vkm+1beanotherspace-timeperiodicsolutionof(1.6)withh(c)=˜Ω¯h∆(c).Extendingtheperiodicsolutionstotheentireoddgrid,wehavethefollowingstochasticandvariationalrepresentationformulasuptoanynegativetimeindexl0:¯vl+1n=Eµ(·;ξ∗)hXl0<k≤l+1L(c)(γk,tk−1,ξ∗k m(γk))∆t+¯vl0m(γl0)i+¯h∆(c)(tl+1−tl0),˜Ω¯vl+1n=Eµ(·;˜ξ∗)hXl0<k≤l+1L(c)(γk,tk−1,˜ξ∗k m(γk))∆t+˜Ω¯vl0m(γl0)i+˜Ω¯h∆(c)(tl+1−tl0).Bythevariationalproperty,wehave˜Ω¯vl+1n−¯vl+1n≤Eµ(·;ξ∗)h˜Ω¯vl0m(γl0)−¯vl0m(γl0)i+(˜Ω¯h∆(c)−¯
h∆(c))(tl+1−tl0).(3.3)ΩNotethat¯vkm+1,˜Ω¯vkm+1areperiodicandhencebounded.Dividing(3.3)bytl+1−tl0andlettingl0→−∞,weobtain0≤˜Ω¯h∆(c)−¯h∆(c).Similarreasoningyieldstheconverseinequality.2.Since¯vkm+1satisﬁesDt¯vk+1m+H(xm,tk,c+Dx¯vkm+1)=¯h∆(c),wehave¯h∆(c)=X0≤k<2KXm;kDt¯vk+1m·2∆x∆t+X0≤k<2KXm;kH(xm,tk,c+Dx¯vkm+1)·2∆x∆t.Theﬁrsttermontheright-handsideisequaltozeroduetotheperiodicityof¯vkm+1.3.Let˜vl+1n(c)bethesolutionofDt˜vk+1m+H(xm,tk,c+Dx˜vkm+1)=¯h∆(c)withthesamemeshsizeas(3.2)andwith˜v0m+1=v0m+1.FromTheorem3.6itfollowsthatwe22
have|˜vl+1n(c)−¯vl+1n(c)|→0asl→∞,addingaconstantifnecessary.Therefore,˜vl+1nisboundedforl→∞.Sincevl+1n(c)=infξEµ(·;ξ)hX0<k≤l+1L(c)(γk,tk−1,ξkm(γk))∆t+v0m(γ0)i,˜vl+1n(c)=infξEµ(·;ξ)hX0<k≤l+1L(c)(γk,tk−1,ξkm(γk))∆t+v0m(γ0)i+¯h∆(c)tl+1,andtheminimizingvelocityﬁeldsofthesearethesame,weobtainvl+1n(c)−˜vl+1n(c)=−¯h∆(c)tl+1.4.Followingtheproofof(6)ofTheorem2.1in[17],wecanprovethatc+¯uk m(c)isaC1-functionofcforeachm,k.Therefore,Claim2yieldsC1-regularityof¯h∆.Letvl+1n(c)beasolutionof(3.2)andﬁxn.Weshowthatthemapc7→vl+1n(c)isaconcavefunctionforeachl+1≥1.Letξ∗betheminimizingvelocityﬁeldforvl+1n(c∗)withc∗:=θc+(1−θ)˜c,θ∈[0,1].Then,wehavevl+1n(c∗)−{θvl+1n(c)+(1−θ)vl+1n(˜c)}≥θEµ(·;ξ∗)hX0<k≤l+1−(c∗−c)ξ∗k m(γk)∆ti+(1−θ)Eµ(·;ξ∗)hX0<k≤l+1−(c∗−˜c)ξ∗k m(γk)∆ti=0.Therefore,themapc7→vl+1n(c)/tl+1isalsoaconcavefunctionand¯h∆(c)=−liml→∞vl+1n(c)
tl+1isaconvexfunction.5.Hereinafterb1,b2,...arepositiveconstantsindependentof∆andc.Foreachx∈T,wehavemsuchthatx∈[xm+1,xm+3).Notethat¯v2Kn∗+1≤¯v∆(x,1)≤¯v2Kn∗+1with(n∗,n∗)=(m,m+2)or(n∗,n∗)=(m+2,m)and¯v2Kn∗+1−¯v(xn∗+1,1)−2r∆x≤¯v∆(x,1)−¯v(x,1)≤¯v2Kn∗+1−¯v(xn∗+1,1)+2r∆x.(3.4)Letx∈Tattainmaxy∈T(¯v∆(y,1)−¯v(y,1))andletn∗bedeﬁnedintheabovemannerwiththisx.Letγ∗beaminimizingcurvefor¯v(xn∗+1,t).Deﬁneξasξkm:=γ∗0(tk).Notethattheηk(γ)deﬁnedbythisξsatisﬁes|ηk(γ)−γ∗(tk)|≤b1∆xforany0≤k≤2K.BytherepresentationformulasandProposition2.5,wehave¯v(xn∗+1,1)=Z10L(c)(γ∗(s),s,γ∗0(s))ds+¯v(γ∗(0),0)+¯h(c),¯v2Kn∗+1≤Eµ(·;ξ)hX0<k≤2KL(c)(γk,tk−1,ξkm(γk))∆t+¯v∆(γ0,0)i+¯h∆(c)(3.5)≤Eµ(·;ξ)hX0<k≤2KL(c)(ηk(γ),tk−1,ξkm(γk))∆t+¯v∆(η0(γ),0)i+¯h∆(c)+b2√
∆x≤Eµ(·;ξ)hX0<k≤2KL(c)(γ∗(tk),tk−1,γ∗0(tk))∆t+¯v∆(γ∗(0),0)i+¯
h∆(c)+b3√
∆x≤Z10L(c)(γ∗(s),s,γ∗0(s))ds+¯v∆(γ∗(0),0)+¯h∆(c)+b4√
∆x.23
Therefore,noting(3.4),wehave¯v∆(x,1)−¯v(x,1)≤¯v∆(γ∗(0),0)−¯v(γ∗(0),0)+¯h∆(c)−¯h(c)+b5√
∆x.Fromtheperiodicityof¯v∆,¯vandtheabovechoiceofxitfollowsthat(¯v∆(x,1)−¯v(x,1))−(¯v∆(γ∗(0),0)−¯v(γ∗(0),0))≥0.Therefore,weobtain−b5√
∆x≤¯h∆(c)−¯h(c).Letx∈Tattainminy∈T(¯v∆(y,1)−¯v(y,1))andletn∗bedeﬁnedintheabovemannerwiththisx.Letξ∗betheminimizingvelocityﬁeldfor¯v2Kn∗+1.Then,wehave¯v2Kn∗+1=Eµ(·;ξ∗)hX0<k≤2KL(c)(γk,tk−1,ξ∗k m(γk))∆t+¯v∆(γ0,0)i+¯h∆(c),≥Eµ(·;ξ∗)hX0<k≤2KL(c)(ηk(γ),tk−1,ξ∗k m(γk))∆t+¯v∆(η0(γ),0)i+¯h∆(c)−b6√
∆x.Letη∆(γ)bethelinearinterpolationofηk(γ).Notethatη∆(γ)0(t)=ξ∗k m(γk)fort∈(tk−1,tk).Foreachγwehave¯v(xn∗+1,1)≤Z10L(c)(η∆(γ)(s),s,η∆(γ)0(s))ds+¯v(η∆(γ)(0),0)+¯h(c)(3.6)≤X0<k≤2KL(c)(ηk(γ),tk−1,ξ∗k m(γk))∆t+¯v(η0(γ),0)+¯h(c)+b7∆x.Therefore,noting(3.4),wehave¯v∆(x,1)−¯v(x,1)≥Eµ(·;ξ∗)h¯v∆(η0(γ),0)−¯v(η0(γ),0)i+¯h∆(c)−¯h(c)−b8√
∆x.Fromtheperiodicityof¯v∆,¯vandtheabovechoiceofxitfollowsthat(¯v∆(x,1)−¯v(x,1))−(¯v∆(η0(γ),0)−¯v(η0(γ),0))≤0forallγ.Thus,weobtain¯h∆(c)−¯h(c)≤b8√
∆x.
3.4ConvergenceofPeriodicSolutionsΩWeprovethatforspace-timeperiodicsolutionsthediﬀerencesolutionsconvergetotheexactonesuptoasubsequence.Notethatviscositysolutionsandentropysolutionswithspace-timeperiodicityarenotuniquewithrespecttocingeneral.Theselectionprobleminﬁnitediﬀerenceapproximationremainsopen.Itisalsochallengingtoinvestigatedetailsoftheconvergenceeveninthecasewheretheuniquenessholds.Wewillmakesomeprogresswiththisissueinthenextsection.
Theorem3.8.Thereexistsasequence∆=(∆x,∆t)→0forwhich{¯v(c)∆}and{¯u(c)∆}convergetoaZ2-periodicviscositysolution¯vof(1.2)withh(c)=¯h(c)andtoaZ2-periodicentropysolution¯u=¯vxof(1.1),respectively:supt∈Tk¯v(c)∆(·,t)−¯v(·,t)kC0→0,supt∈Tk¯u(c)∆(·,t)−¯u(·,t)kL1(T)→0.24
Proof.Ifnecessary,weaddaconstantsothat¯v(c)∆(·,0)isboundedbyr.Then,{¯v(c)∆(·,0)}isafamilyoffunctionsthatareuniformlyboundedandequicontinuous.Wehaveacon-vergentsubsequence,stilldenotedby¯v(c)∆(·,0):¯v(c)∆(·,0)→¯v0.Let¯vbetheviscositysolutionof(1.4)withv0=¯v0andh(c)=¯h(c).Then,wehaveaminimizingcurvesuchthat¯v(xn,tl+1)=Ztl+10L(c)(γ∗(s),s,γ∗0(s))ds+¯v0(γ∗(0))+¯h(c)tl+1.Byanestimatesimilarto(3.5),wehave¯v∆(xn,tl+1)≤Ztl+10L(c)(γ∗(s),s,γ∗0(s))ds+¯v∆(γ∗(0),0)+¯h∆(c)tl+1+b1√
∆x.Since¯h∆(c)→¯h(c),weobtainlimsup∆→0{¯v(c)∆(xn,tl+1)−¯v(xn,tl+1)}≤0,whichisuniformwithrespectto(xn,tl+1)∈T2.Byanestimatesimilarto(3.6),weobtainliminf∆→0{¯v(c)∆(xn,tl+1)−¯v(xn,tl+1)}≥0,whichisuniformwithrespectto(xn,tl+1)∈T2.Therefore,weconcludethat¯v(c)∆→¯vuniformlyonT2and¯visZ2-periodicduetotheperiodicityof¯v∆.ThroughreasoningsimilartotheproofofTheorem2.8in[22],itfollowsthat¯u(c)∆:=(¯v(c)∆)xconvergesto¯u=¯vxpointwisealmosteverywhereinT2,where{¯v(c)∆}istheconvergentsubsequenceabove.Hence,wehavek¯u(c)∆(·,t)−¯u(·,t)kL1(T)→0foreacht.ThroughreasoningsimilartotheproofofProposition2.14in[17],itfollowsthat¯u(c)∆satisﬁesk¯u(c)∆(·,tk)−¯u(c)∆(·,tk0)kL1(T)≤b2|tk−tk0|withaconstantb2independentofk,k0,c,and∆.Therefore,¯u∈Lip(T;L1(T))withtheLipschitzconstantb2.Thus,wehavedemonstratedthetheorem.
4ErrorEstimatesΩWeshowerrorestimatesforentropysolutionsofinitialvalueproblemsandforZ2-periodicentropysolutionsinthespecialcasewheretheyareassociatedwithKAMtori.ThelatterisarigorousresultonﬁnitediﬀerenceapproximationofKAMtori.Wereferto[2]foranerrorestimateforZ2-periodicentropysolutionsassociatedwithKAMtoriinthevanishingviscositymethod.
4.1ErrorEstimatesforInitialValueProblemΩThefollowingtheoremprovideserrorestimatesfortheinitialvalueproblem.Theorem4.1.Let∆=(∆x,∆t)satisfytheconditionsinTheorem2.6andProposi-tion2.8.Letubetheentropysolutionof(1.3)andu∆begivenbythediﬀerencesolutionof(2.3).Then,thefollowinghold:25
1.ForanyT∈(0,∞),foreacht∈(0,T],andindependentoftheinitialdata,thereexistsaconstantβ4(t)>0forwhichku∆(·,t)−u(·,t)kL1(T)≤β4(t)∆x1
4.Inparticular,ifu0israrefaction-free,thenthereexistsaconstantβ5>0forwhichsup0≤t≤Tku∆(·,t)−u(·,t)kL1(T)≤β5∆x1
4.2.IfuisLipschitzinT×[0,T],thenthereexistsaconstantβ6>0forwhichsup(x,t)∈T×[0,T]|u∆(x,t)−u(x,t)|≤β6∆x1
4.Proof.1.Letv∆andvcorrespondtou∆andu,respectively.ByTheorem2.6,forallt∈[0,T]andallinitialdata,wehavekv∆(·,t)−v(·,t)kC0≤β2√
∆x.(4.1)ΩByProposition2.8,foreacht∈[∆t,T]andallinitialdata,wehaveu∆(xm+2,t)−u∆(xm,t)
2∆x≤Ek(t)∆.(4.2)ΩSinceu∆(·,t)haszeromean,wehaveXm;k(t){u∆(xm+2,t)−u∆(xm,t)}=Xm:+{u∆(xm+2,t)−u∆(xm,t)}+Xm:−{u∆(xm+2,t)−u∆(xm,t)}=0,wherePm:+(resp.Pm:−)standsforthesummationwithrespecttomforwhichu∆(xm+2,t)−u∆(xm,t)≥0(resp.u∆(xm+2,t)−u∆(xm,t)<0).Hence,itfollowsfrom(4.2)thatthetotalvariationofu∆(·,t)onTisbounded:Xm;k(t)|u∆(xm+2,t)−u∆(xm,t)|=2Xm:+{u∆(xm+2,t)−u∆(xm,t)}≤2Ek(t)∆.(4.3)Foranyε>0,thereexists˜∆=(˜∆x,˜∆t)suchthatku˜∆(·,t)−u(·,t)kL1(T)≤ε.Inparticular,wetakesucha˜∆=(˜∆x,˜∆t)thatsatisﬁes˜∆t/˜∆x=∆t/∆x,˜∆x≤(β−12ε)4,and∆x/˜∆x=3pforsomep∈N.Thelastrelationguaranteesthatthepointsofdiscontinuityofu∆arealsothoseofu˜∆.Then,wehaveku∆(·,t)−u(·,t)kL1(T)≤ku˜∆(·,t)−u∆(·,t)kL1(T)+ε,kv˜∆(·,t)−v∆(·,t)kC0≤β2√
∆x+β2p
˜∆x≤2β2√
∆x.(4.4)26
Nowweestimateku˜∆(·,t)−u∆(·,t)kL1(T).Weintroducew∆:=u˜∆(·,t)−u∆(·,t),˜w∆:=v˜∆(·,t)−v∆(·,t)and˜k(t):=3pk(t).Letxm∈˜∆xZandsetxm0:=0for˜k(t)evenorxm0:=˜∆xfor˜k(t)odd.Wedivide˜∆xZaccordingtothesignofw∆.Thatis,I1,I2,...,In+1aredeﬁnedasI1:={xm0,xm0+2,···,xm1}onwhichw∆(x)≥0(or<0),I2:={xm1+2,xm1+4,···,xm2}onwhichw∆(x)<0(or≥0),I3:={xm2+2,xm2+4,···,xm3}onwhichw∆(x)≥0(or<0),···,In:={xmn−1+2,xmn−1+4,···,xmn}onwhichw∆(x)<0(or≥0),In+1:={xmn+2,xmn+4,···,xm0+1}onwhichw∆(x)≥0(or<0),wherenisevenandxmn≤xm0+1−2˜∆x.WethenredeﬁneI1asI1:={xm0,xm0+2,···,xm1}withxm0:=xmn+2−1.Notethatw∆(x)≥0(or<0)onI1.Setting|I1|:=xm1−xm0+2˜∆xand|Ij|:=xmj−xmj−1+2+2∆xforj>1,wehavePn j=1|Ij|=1.ForeachIjonwhichw∆(x)≥0(resp.<0),wehaveayj∈Ijforwhichw∆(x)takesthemaximum(resp.minimum)withinIj.Supposethatw∆(x)≥0onI1.Intheothercase,theargumentisparallel.Notethatkw∆(x)kL1(T)=n/2Xj=1

Xx∈I2j−1w∆(x)·2˜∆x−Xx∈I2jw∆(x)·2˜∆x

.IntroducingJ:={j|0≤j≤n/2,max{|I2j−1|,|I2j|}<∆x1/4}and˜J:={j|0≤j≤n/2,max{|I2j−1|,|I2j|}≥∆x1/4},wehave]˜J·∆x1/4≤1and]˜J≤∆x−1/4.Therefore,noting(4.3)and(4.4)aswellasw∆=(˜w∆)x,weobtainkw∆(x)kL1(T)=Xj∈J

Xx∈I2j−1w(x)·2˜∆x−Xx∈I2jw∆(x)·2˜∆x

+Xj∈˜J

Xx∈I2j−1w∆(x)·2˜∆x−Xx∈I2jw∆(x)·2˜∆x

≤Xj∈J|w∆(y2j−1)−w∆(y2j)|∆x1
4+Xj∈˜Jh{˜w∆(xm2j−1+˜∆x)−˜w∆(xm2j−2+2−˜∆x)}−{˜w∆(xm2j+˜∆x)−˜w∆(xm2j−1+2−˜∆x)}i≤(2E˜k(t)˜∆+2Ek(t)∆)∆x1
4+]˜J·4·2β2√
∆x≤4Ek(t)∆∆x1
4+8β2∆x1
4,Sinceεisarbitrary,weconcludethatku∆(·,t)−u(·,t)kL1(T)≤(4Ek(t)∆+8β2)∆x1
4.27
Ifu0israrefaction-free,wehaveM>0suchthatEk(t)∆≤max{M,E∗}forall0≤t≤T.2.Fixt∈[0,T]arbitrarily.By(4.1),foranyx,x0∈Twehave|Zxx0u∆(y,t)−u(y,t)dy|≤2β2√
∆x.From(4.2)itfollowsthatk˜u∆−u∆(·,t)kL1(T)≤b1∆x,where˜u∆(x)denotesthelinearinterpolationofuk(t)mwithrespecttothespacevariable.Hence,settingw∆:=˜u∆−u(·,t),forallx,x0∈Twehave|Zxx0w∆(y)dy|≤b2√
∆x.(4.5)ΩSinceuisLipschitz,w∆stillsatisﬁestheone-sidedLipschitzconditionw∆(x1)−w∆(x2)
x1−x2≤b3.Notethatw∆doesnotnecessarilysatisfyanyLipschitzcondition,because˜u∆doesnotnecessarilysatisfyanyLipschitzcondition.Supposethat|w∆(¯x)|>b4∆x1
4with(b4)2/(4b2)>b3forsome¯x.LetI3¯xbeaconnectedintervalonwhoseboundarywehave|w∆(x)|=b4
2∆x1
4.By(4.5),weﬁndthat|I|≤2b2
b4∆x1
4.Ifw∆(¯x)>0(resp.<0),andwiththeleft(resp.right)boundaryofIdenotedbyx,wehavew∆(¯x)−w∆(x)
¯x−x≥(b4)2
4b2>b3resp.w∆(x)−w∆(¯x)
x−¯x≥(b4)2
4b2>b3,whichisacontradiction.Therefore,weobtainkw∆kC0≤b4∆x1
4.Since|u∆(x,t)−u(x,t)|=|uk(t)m−u(x,t)|≤|uk(t)m−u(xm,t)|+b5∆x=|w∆(xm)|+b5∆x,wehavedemonstratedthetheorem.
4.2ErrorEstimateforKAMToriΩLet¯u(c)=¯v(c)xbeaZ2-periodicentropysolutionoftheC1-class.Weremarkthere-lationshipbetweensucha¯u(c)andHamiltoniandynamics.Considerthetime-1mapf:T×R→T×RoftheHamiltonianﬂowgeneratedbytheﬂuxfunctionH(x,t,p)withtheinitialtimeequaltozero.Then,{(x,c+¯u(c)(x,0))|x∈T}∼=Tisasmoothinvarianttorusoff.AccordingtotheclassicalresultofPoincar´e,thereexistsarotationnumberω1.LetusregardthenonautonomousHamiltoniandynamicsgeneratedbyH(x,t,p)astheautonomousdynamicsgeneratedbyH(q1,q2,p1,p2):=p2+H(q1,q2,p1)intheextendedphasespaceT2×R2.WedeﬁneI(¯u(c)):={(q,g(q))|q=(q1,q2)∈T2}∼Ω=T2,28
whereg(q):=(c+¯u(c)(q1,q2),¯
h(c)−H(q1,q2,c+¯u(c)(q1,q2)).Then,I(¯u(c))isasmoothinvarianttorusoftheHamiltonianﬂowϕs
HgeneratedbyH.LetC(s):=(γ∗(s),s)bethecharacteristiccurvesof¯u(c),whichsatisfyγ∗0(s)=Hp(γ∗(s),s,c+¯u(c)(γ∗(s),s)

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

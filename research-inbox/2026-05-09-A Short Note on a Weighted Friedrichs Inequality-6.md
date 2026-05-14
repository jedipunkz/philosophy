---
source: "https://arxiv.org/abs/1809.04577v3"
title: "A Short Note on a Weighted Friedrichs Inequality"
author: "Immanuel Anjam, Dirk Pauly"
year: "2018"
publication: "arXiv preprint / math.AP"
download: "https://arxiv.org/pdf/1809.04577v3"
pdf: "https://arxiv.org/pdf/1809.04577v3"
captured_at: "2026-05-09T12:41:08Z"
updated_at: "2026-05-09T12:41:08Z"
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

# A Short Note on a Weighted Friedrichs Inequality

- 著者: Immanuel Anjam, Dirk Pauly
- 年: 2018
- 掲載情報: arXiv preprint / math.AP
- 情報源: [arxiv](https://arxiv.org/abs/1809.04577v3)
- ダウンロード: https://arxiv.org/pdf/1809.04577v3
- PDF: https://arxiv.org/pdf/1809.04577v3

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

In this short note we derive, for bounded domains, an upper bound for a Friedrichs type constant in a weighted Friedrichs type inequality. This upper bound generalizes a well known upper bound of the Friedrichs constant. This upper bound is also used to improve an upper bound of a Maxwell type constant for convex domains in $\mathbb{R}^3$. A simple numerical application is also given: we apply the main result in a posteriori error estimation for an elliptic problem.

## PDF Text

arXiv:1809.04577v3 [math.AP] 4 Mar 2019
ASHORTNOTEONAWEIGHTEDFRIEDRICHSINEQUALITYIMMANUELANJAMANDDIRKPAULYAbstract.Inthisnotewederiveanupperboundfortheconstantcf,α>0intheweightedFriedrichstypeinequality∀ϕ∈˚H1(Ω)|ϕ|L2≤cf,αp hα∇ϕ,∇ϕiL2,whereΩ⊂Rd,d≥1isaboundeddomain,andαabounded,self-adjoint,anduniformlypositivedeﬁnitematrixvaluedfunction.Thecontentsofthisnotefollowinastraightforwardmannerfromwellknownresults.Inparticular,foraconstantdiagonalmatrixαweobtaintheboundcf,α≤πr
α1
l21+···+αd l2d−1,whereliarethesidelengthsofad-intervalencompassingΩ,andαiarethediagonalentriesofα.Extensionstocasesofunboundeddomainsandpartialhomogeneousboundaryconditionsareremarkedupon.Wealsoapplythemainresultinaposteriorierrorestimationforanellipticproblemandpresentsomenumericalresults.Lastly,weusethemainresulttoderiveanimprovedupperboundofthetangentialMaxwellconstantforconvexdomains.Contents1.Introduction22.AWeightedFriedrichsInequality33.TheTangentialMaxwellInequalityforConvexDomainsinR38References11AppendixA.PositiveSemi-DeﬁniteMatricesα11
Date:4March2019.2010MathematicsSubjectClassiﬁcation.35A23,26D10,35Q61.Keywordsandphrases.Friedrichsconstant,Friedrichsinequality,Maxwellconstant,Maxwellinequality,boundeddomain.1
2IMMANUELANJAMANDDIRKPAULY1.IntroductionWedenotebyx:=(x1,...,xd)theEuclideancoordinatesinRd,d≥1,andbyΩ⊂Rdaboundeddomain.Thecalculationsperformedinthisnoteareinvariantwithrespecttotranslationsofthedomain,sowithoutlossofgeneralityweassumeΩtobecontainedintheopend-intervalI:=dYi=1(0,li),0<li<∞.Thespaceofsmoothscalar-orvector-valuedfunctionsvanishingontheboundaryofthedomainisdenotedby˚C∞(Ω).Wedenotebyh·,·iL2and|·|L2theinnerproductandnormforscalar-orvector-valuedfunctionsinL2(Ω).Weintroducethenotationh·,·iL2,ρ:=hρ·,·iL2(Ω),whichinduces|·|L2,ρ,whereρbelongstothespaceofessentiallyboundedfunctionsL∞(Ω).Ifρisself-adjointanduniformlypositivedeﬁnite,theybecomeaninnerproductandanorminL2(Ω),respectively.Thespaceofscalar-valuedfunctionsinL2(Ω)withzeromeanisdeﬁnedasL2Ω0(Ω):=ϕ∈L2(Ω)

ZΩϕdx=0,andasusual,foravector-valuedfunctionφwewriteφ∈L2Ω0(Ω)ifallitscomponentsbelongtoL2Ω0(Ω).IntherestofthispaperwemaydropΩinournotationsforbrevity,i.e.,L2:=L2(Ω).WedeﬁnetheusualSobolevspacesH1:={ϕ∈L2|∇ϕ∈L2},˚H1:=
˚C∞H1,D:={φ∈L2|divφ∈L2},˚D:=
˚C∞D,whichareHilbertspaces.Notethatontheformerspacesthediﬀerentialoperatorsarenowdeﬁnedintheusualweaksense.Thelatterspaces,wheretheclosuresaretakenwithrespecttographnorms,generalizetheclassicalhomogeneousscalarandnormalboundaryconditions,respectively.TheFriedrichsinequalityreadsas∀ϕ∈˚H1|ϕ|L2≤cf|∇ϕ|L2,wherecf=cf(Ω)>0iscalledtheFriedrichsconstant.Notethatcfisassumedtobethebestpossible,i.e.,smallestpossibleconstantforwhichtheFriedrichsinequalityholds.Acommonlyutilizedupperboundforcfis[10](1.1)cf≤ πs
1
l21+···+1
l2d!−1.Thisnoteisdedicatedtoﬁndingupperboundsfortheconstantcf,α=cf,α(Ω,α)>0intheweightedFriedrichstypeinequalityΩ(1.2)∀ϕ∈˚H1|ϕ|L2≤cf,α|∇ϕ|L2,αforboundedΩ.Hereα∈L∞isaself-adjoint(i.e.,equaltoitsconjugatetranspose),uniformlypositivedeﬁnitematrixvaluedfunctionα:Ω→Rd×d,i.e.,itsatisﬁes(1.3)∃α
>0∀φ∈L2α
|φ|2
L2≤hαφ,φiL2.Estimatesforcf,αcanbecalculatedbyusingestimatesforcf,sinceobviously|ϕ|L2≤cf|∇ϕ|L2≤cf
√
α
|∇ϕ|L2,αholds.NotethatsincetheﬁrstestimationstepisdoneusingtheFriedrichsinequality,andcontainsthefullgradientontherighthandside,itisinevitablethattheﬁnalestimationstep
ASHORTNOTEONAWEIGHTEDFRIEDRICHSINEQUALITY3involvesadivisionwiththesmallesteigenvalueofα.Estimatingfurtherbyusing(1.1)weobtaintheestimateΩ(1.4)cf,α≤ πs
α
1
l21+···+1
l2d!−1,whichblowsupasα
approacheszero.HavingcomputableupperboundsofFriedrichs,Poincar´e,andMaxwelltypeconstantsrelatedtobothweightedandnon-weightedvariantsofcorrespondinginequalitiesisimportantinaposteriorierrorestimation.Errorupperboundstypicallycontaintheseconstants,andareespeciallyimportantforfunctionaltypeerrorestimates,whereguaranteedupperboundsoftheexacterroraredesired.Inthisnoteweomitaliteratureoverviewofaposteriorierrorestimation,andinsteadreferthereadertothebooks[2,9,11,16,19].SomereferenceswithupperboundsofFriedrichsandPoincar´etypeconstantsarethebook[10]andthepaper[15](seealso[5]).Wealsocitetheinterestingsurveyarticle[7].Somemorerecentworkonthesubjectinclude[18],whereaweightedFriedrichsinequalitysimilarto(1.2)isconsidered.Theauthorcalculatesnumericallytwo-sidedboundsofaFriedrichstypeconstantinweightednorms.Thisapproachallowsformixedboundaryconditions.In[17]FriedrichsandPoincar´einequalitiesinnon-weightednormswithmixedboundaryconditionsareconsidered.ThisapproachinvolvesdecomposingthedomainintosmallersubdomainsforwhichFriedrichsandPoincar´econstantsareknown.Theresultingupperboundsdependonthedecomposition.ComputableupperboundsofMaxwellconstantsforconvexdomainshavebeenstudiedbytheauthorsofthepresentnote.Inthesecondauthor’spapers[12–14]itisshownthattheMaxwellconstantsareboundedbyabovebythePoincar´econstant,andasmallimprovementtotheseresultscanbefoundintheﬁrstauthor’spaper[4].InthisnoteweshowthatinthecaseoffullhomogeneousDirichletboundaryconditions,thereisasimplewaytoobtainanupperboundofcf,αwithbetterpropertiesthan(1.4).Theupperbound,derivedinSection2,followsfromwellknownresults.Inthissectionwealsodemonstratethebeneﬁtofusingtheimprovedupperboundofcf,αbyanumericalexamplewhereweperformaposteriorierrorestimationofanellipticproblem.InSection3weusethisupperboundtoimproveanupperboundofthetangentialMaxwellconstantforconvexdomainsinR3.2.AWeightedFriedrichsInequalityThecalculationsofthissectionarebasedonthewellknownone-dimensionalinequality(2.1)∀ϕ∈˚H1((0,l))Zl0|ϕ(y)|2dy≤l2
π2Zl0|ϕ0(y)|2dy,where0<l<∞.UsingthisinequalityonecanproofaFriedrichstypeinequalityinvolvingonlyonepartialderivative,andbyanadditionalestimationstepobtainaninequalityinvolvingthefullgradient.Inthecaseofboundeddomains,thiswouldresultintheestimate(1.1).However,wewillneedtheintermediateresultinvolvingonlyonepartialderivative.Note,thatsincewewanttocontrolallpartialderivativesseparately(withrespecttothealreadychosencoordinatesystem),wecannotrotatethedomain.
Lemma2.1.LetΩbebounded,andi∈{1,...,d}.Thenwehavetheestimate∀ϕ∈˚H1|ϕ|L2≤li
π|∂iϕ|L2.Proof.Considerﬁrsttherealvaluedcaseandi=1.Foranyϕ∈˚C∞(Ω)itszero-extensionˆϕ:I→Rbelongsto˚C∞(I).Forany˜x:=(x2,...,xd)belongingto˜I:=(0,l2)×···×(0,ld),
4IMMANUELANJAMANDDIRKPAULYthefunctionˆϕ(x1,˜x)isarealvaluedfunctionofonevariablevanishingattheendpointsoftheinterval(0,l1),soby(2.1)wehaveZl10|ˆϕ(x1,˜x)|2dx1≤l21
π2Zl10|∂1ˆϕ(x1,˜x)|2dx1.Byintegratingtheabovewithrespectto˜xin˜I,weobtain|ˆϕ|2
L2(I)≤l21
π2|∂1ˆϕ|2
L2(I)⇒|ϕ|2
L2(Ω)≤l21
π2|∂1ϕ|2
L2(Ω),sincethenormsarenonzeroonlyinΩ.Bydensitytheaboveholdsforanyϕ∈˚H1(Ω).Byanidenticalproceduretheassertionfollowsfori∈{2,...,d}forrealvaluedfunctions.Havingestablishedtheassertionforrealvaluedfunctions,itisclearthatitholdsforcomplexvaluedfunctionsaswell.Wenowconsidertheconstantcf,αintheinequality(1.2).Weassumethatα∈L∞isaself-adjointdiagonalmatrixΩ(2.2)α:=
α10...0αd
satisfyinguniformpositivedeﬁniteness(1.3),whichinthiscaseisequivalentto(2.3)∃α
i>0∀ϕ∈L2α
i|ϕ|2
L2≤hαiϕ,ϕiL2,i=1,...,d.Notethatsuchanαhasnoimaginarypart,andthatα
=min{α
1,...,α
d}.Theorem2.2.LetΩbeboundedandα∈L∞beaself-adjointdiagonalmatrixsatisfying(2.2)–(2.3).Thenwehavetheestimatecf,α≤πr
α
1
l21+···+α
d l2d−1.Proof.Letϕ∈˚H1.Sinceαisdiagonal,theweightednormcanbewrittenas|∇ϕ|2
L2,α=|∂1ϕ|2
L2,α1+···+|∂dϕ|2
L2,αd.Lemma2.1gives|ϕ|2
L2≤l2i
π2|∂iϕ|2
L2≤l2i
π2α
i|∂iϕ|2
L2,αiforanyi∈{1,...,d}.Bymultiplyingtheabovebyα
i/l2iandsummingupthedinequalities,weobtainα
1
l21+···+α
d l2d|ϕ|2
L2≤1
π2|∇ϕ|2
L2,α,whichimpliestheassertion.Remark2.3.(i)Theorem2.2withα=idresultsintheestimate(1.1).(ii)ItiseasytoseethattheupperboundofTheorem2.2isalwayssmallerorequaltotheupperbound(1.4).(iii)Theaboveprocedurefurnishesupperboundsofcf,αevenwhenthediagonalmatrixαisnotuniformlypositivedeﬁnite(seeAppendixA).
ASHORTNOTEONAWEIGHTEDFRIEDRICHSINEQUALITY5(iv)TheresultofLemma2.1,andthusalsoTheorem2.2,holdsalsoforanunboundeddomainlyingbetweentwoparallelhyperplanes,providedthatthehyperplanesarenotparalleltoanycoordinateaxes.ThereasonforthislimitationisbecauseintheproofofLemma2.1weneed(2.1)inthedirectionofallthecoordinateaxes.Notethatinthiscasetheconstantslidenotethedistanceofthesetwohyperplanesmeasuredbyalineparalleltothexiaxis.Theproofofthisresultisonlyslightlymoreinvolved.(v)AnupperboundsimilartoTheorem2.2forboundeddomainscanbeobtainedprovidedthatthehomogeneousDirichletboundaryconditionsholdinatleastonedirection.I.e.,fortheunitsquareinR2,itisenoughthatonepairofopposingboundarieshavetheboundarycondition;ifthispairistheoneparalleltothex2-axis,thentheboundaryconditionisinthedirectionofthex1-axis,andwehavecf,α≤(πp
α
1/l21)−1.However,sincetheopposingboundarypartsmusthavea”straightlineofsight”toeachother,thepossibledomainsarequitelimited.Morevarietyindomainsisachieved,ifoneuses(insteadof(2.1))Zl0|ϕ(y)|2dy≤l2
2Zl0|ϕ0(y)|2dy,whichholdsforallfunctionsϕ∈H1((0,l))vanishingeitheronthebeginningortheendoftheinterval.Inthiswayonlyoneoftheopposingboundariesneedtohavetheboundarycondition.Forthisinequalitysee,e.g.,[1,p.158].Undercertainconditionsnon-diagonalαcanbehandledaswell.Forreadabilityweconsideronlythethreedimensionalcase.Foranyself-adjointα(x)={αij(x)}3
i,j=1fromL∞wedeﬁne(2.4)˜α:=
˜α1000˜α2000˜α3
,˜α1:=α11−(|<α12|+|=α12|+|<α13|+|=α13|),˜α2:=α22−(|<α12|+|=α12|+|<α23|+|=α23|),˜α3:=α33−(|<α13|+|=α13|+|<α23|+|=α23|).Itiseasytoverifythat˜αisself-adjoint,andthat(2.5)∀φ∈L2|φ|L2,˜α≤|φ|L2,αholds.If˜αisalsouniformlypositivedeﬁnite,i.e.,itsatisﬁes∃˜α
i>0∀ϕ∈L2˜α
i|ϕ|2
L2≤h˜αiϕ,ϕiL2,i=1,2,3,wecandirectlyuseTheorem2.2toobtainanestimateofcf,α.Theorem2.4.LetΩ⊂R3beboundedandα∈L∞beaself-adjointmatrixvaluedfunctionforwhich˜α,deﬁnedby(2.4),isuniformlypositivedeﬁnite.Thenwehavetheestimatecf,α≤ πs
˜α
1
l21+˜α
2
l22+˜α
3
l23!−1.Proof.Letϕ∈˚H1.Theorem2.2gives|ϕ|L2≤ πs
˜α
1
l21+˜α
2
l22+˜α
3
l23!−1|∇ϕ|L2,˜α,andwith(2.5)wehavetheassertion.Remark2.5.For˜αtobeuniformlypositivedeﬁnitewouldrequirethattheoﬀ-diagonalentriesofαbecomparativelysmallcomparedtoitsdiagonalentries.However,nowRemark2.3(iii)holdswithrespectto˜α.Inparticular,foranupperboundofcf,αitisenoughthat˜αispositivesemi-deﬁnitesuchthatoneofthediagonalentriesof˜αisuniformlypositivedeﬁnite.Wedemonstratethederivedresultsintherealvaluedsettingthroughsomeexamples.
6IMMANUELANJAMANDDIRKPAULYExample1(Diagonalmatrixα).LetΩ⊂(0,1)2andαbetheuniformlypositivedeﬁniteconstantmatrixα=10Ω0δ,δ>0.Theestimate(1.4)givestheupperboundΩ(2.6)cf,α≤ πp
2min{1,δ}−1,andTheorem2.2givesΩ(2.7)cf,α≤ π√
1+δ−1.Itiseasytoseethatthelatterdoesnotblowupasδbecomessmaller.Table1showsthevaluesoftheboundswithdiﬀerentδ.Table1.Example1:Valuesoftheupperbounds(2.6)and(2.7)withdiﬀerentδ.δ
10−6
10−4
10−2
1
102
104
106
(2.6)
225.07908
22.50791
2.25079
0.22508
0.22508
0.22508
0.22508(2.7)
0.31831
0.31829
0.31673
0.22508
0.03167
0.00318
0.00032Example2(Solutiontheoryforareaction-diﬀusionproblem).Considerthefollowingreaction-diﬀusionproblem:ﬁndu∈˚H1satisfying−divα∇u+ρu=f,wheref∈L2,ρ∈R,andα∈L∞isasymmetricuniformlypositivedeﬁnitematrixvaluedfunction.Thevariationalformulationofthisproblemreadsas∀ϕ∈˚H1h∇u,∇ϕiL2,α+hu,ϕiL2,ρ=hf,ϕiL2.Bysettingϕ=uinthebilinearformonthelefthandside,weobtainh∇u,∇uiL2,α+hu,uiL2,ρ=(1−)|∇u|2
L2,α+|∇u|2
L2,α+|u|2
L2,ρ≥(1−)α
|∇u|2
L2+
c2Ωf,α+ρ|u|2
L2,where0<<1.Weobservethatthisformiscoerciveprovidedthat
c2Ωf,α+ρ>0holds,andunderthisconditionauniquesolutionexistsin˚H1bytheRieszrepresentationtheorem.LetΩ⊂(0,1)2andα=10Ω0100.Using(1.4)toestimatecf,α(seeExample1),weseethatforexistenceanduniquenessofasolution,thenecessaryconditionisρ>−2π2,butusingTheorem2.2thenecessaryconditionbecomesρ>−101π2,allowingforalargerrangeofadmissibleρ.Example3(Non-diagonalmatrixα).LetΩ⊂(0,1)3,andα=
311Ω13001Ω113
,˜α=
100Ω02980Ω001
,where˜αiscalculatedaccordingto(2.4).Nowα
=2,and(1.4)givestheboundcf,α≤ π√
6−1.Theorem2.4givestheupperboundcf,α≤ π√
300−1,whichissharper.
ASHORTNOTEONAWEIGHTEDFRIEDRICHSINEQUALITY7Asstatedintheintroduction,themotivationforderivingcomputableupperboundsfortheconstantcf,αisthatitisessentialinaposteriorierrorestimationfornumericalapproximationsofellipticpartialdiﬀerentialequations.Asanexampleweconsiderthediﬀusionproblemintherealvaluedsetting,inaboundeddomainΩ,withhomogeneousDirichletboundaryconditionsonthewholeboundary:ﬁndu∈˚H1satisfying−divα∇u=f,wheref∈L2andα∈L∞isasymmetricuniformlypositivedeﬁnitematrixvaluedfunction.ThevariationalformulationforthisproblemreadsasΩ(2.8)∀ϕ∈˚H1h∇u,∇ϕiL2,α=hf,ϕiL2.Since(1.2)issatisﬁed,auniquesolutionu∈˚H1existsbytheRieszrepresentationtheorem.Bysettingϕ=uin(2.8)weseethatthesolutiondependscontinuouslyontherighthandside:|∇u|2
L2,α=hf,uiL2≤|f|L2|u|L2≤cf,α|f|L2|∇u|L2,α⇒|∇u|L2,α≤cf,α|f|L2.Wenowpresentthefunctionaltypeaposteriorierrorupperbound,whichcanbefoundin,e.g.,thebooks[9,11,16].
Theorem2.6.Let˜u∈˚H1beanarbitraryapproximationofu,and˜cf,αbeanyapproximationofcf,αfromabove.Thenwehavetheestimate∀y∈D|∇(u−˜u)|L2,α≤˜cf,α|f+divy|L2+|y−α∇˜u|L2,α−1:=M(˜cf,α,˜u,y).Proof.Webeginbysubtractingthetermh∇˜u,∇ϕiL2,αfrombothsidesof(2.8)andobtainh∇(u−˜u),∇ϕiL2,α=hf,ϕiL2−h∇˜u,∇ϕiL2,α=hf+divy,ϕiL2+hy−α∇˜u,∇ϕiL2≤|f+divy|L2|ϕ|L2+|y−α∇˜u|L2,α−1|∇ϕ|L2,α≤ cf,α|f+divy|L2+|y−α∇˜u|L2,α−1|∇ϕ|L2,α,whereweusedhdivy,ϕiL2+hy,∇ϕiL2=0and(1.2).Settingϕ=u−˜uﬁnishestheproof.Remark2.7.Byusingtheupperboundcf,α≤cf/√
α
forthevalueof˜cf,αweobtainthemostcommonlyusedformofthisfunctionaltypeaposteriorierrorupperboundforthediﬀusionproblem.Notethattheaboveestimateissharp,i.e.,theoreticallythereisnogapbetweentheexacterrorandtheestimate.Thisisseenbysettingy=α∇u∈D.Theﬁrsttermoftheerrorfunc-tionalMvanishes,anditbecomesapparentthatsharpnessdoesnotdependoncf,α.However,obtaininggooderrorboundsrequiresnotonlychoosingyclosetotheexactﬂuxα∇u,butalsohavinggoodupperboundsfortheunknownconstantcf,α.Especiallyinthecasewhen−divyisnotclosetof,alargeover-estimationoftheconstantcf,αwillleadtoalargeover-estimationoftheerror,aswewillnowdemonstrate.
Example4(ErrorestimationwithRaviart-Thomasaveraging).Wesolvethediﬀusionproblem(2.8)inintheL-shapeddomainΩ=(0,1)2\[(1/2,1)×(0,1/2)]withα=10Ω010−4,f=1.Weuselinearnodalﬁniteelementsintrianglestosolve(2.8),anddenotetheapproximationby˜u.ThefunctionyinthefunctionalMisobtainedbyaveragingα∇˜utotheedgesofthemeshresultinginafunctionfromthelinearRaviart-Thomasﬁniteelementspace,whichisasubspaceofD.WedenotethisaveragingoperatorbyGRT.Using(1.4)toestimatethevalueofcf,α(seeExample1),wehavetheestimate(2.9)|∇(u−˜u)|L2,α≤M(22.50791,˜u,GRT(α∇˜u)),
8IMMANUELANJAMANDDIRKPAULYandbyusingTheorem2.2weobtaintheestimateΩ(2.10)|∇(u−˜u)|L2,α≤M(0.31829,˜u,GRT(α∇˜u)).Since−divGRT(α∇˜u)isonlyaroughapproximationoff,thequalityofthelatterestimateisbetter,asisseenfromTable2.Table2.Example4:Valuesoftheupperbounds(2.9)and(2.10)withdiﬀerentmeshes.#elements
384
1536
6144
24576
98304
(2.9)
18.4444
17.1419
16.1891
14.9832
13.2664(2.10)
1.5563
0.9166
0.5705
0.3809
0.26953.TheTangentialMaxwellInequalityforConvexDomainsinR3Inthissection,afterintroducingsomeadditionalnotation,weimproveanupperboundofthetangentialMaxwellconstantusingTheorem2.2.Throughout,weworkinthreedimensions,i.e.,thedomainΩbelongstoR3.AsidefromtheSobolevspacesalreadydeﬁnedintheintroduction,wealsodeﬁneR:={φ∈L2|rotφ∈L2},˚R:=
˚C∞R,whichareHilbertspaces.Asbefore,ontheformerspacethediﬀerentialoperatorrotisdeﬁnedintheusualweaksense.Thelatterspace,wheretheclosureistakenwithrespecttothegraphnorm,generalizestheclassicaltangentialboundarycondition.Notethattherotationrotisoftenwrittenascurlor∇×intheliterature.Letε∈L∞beaself-adjointuniformlypositivedeﬁnitefunctionε:Ω→R3×3,i.e.,itsatisﬁes∃ε
,
ε>0∀φ∈L2ε
|φ|2
L2≤hεφ,φiL2≤
ε|φ|2
L2.Notethatheretheoverlinein
εdoesnotdenotecomplexconjugation.Inthissectionthepropertiesassumedfromεaresimilartowhatwasassumedforαintheprevioussection.Weuseεinsteadofαtoconformtotheusualnotationusedinelectromagnetictheorywhereεdenotestheelectricpermittivityofthemedia.SinceΩisconvex,itisalsoLipschitz[6],andRellich’sselectiontheoremandWeck’sselectiontheorem[20]hold.ThusthespacesinthewellknownHelmholtzdecompositions(see,e.g.,[8])L2=∇˚H1⊕εε−1D0=˚R0⊕εε−1rotR,∇˚H1=˚R0,D0=rotR,(3.1)L2=∇H1⊕εε−1˚D0=R0⊕εε−1rot˚R,∇H1=R0,˚D0=rot˚R(3.2)Ωareclosed.Here⊕εdenotesorthogonalsumwithrespecttotheweightedscalarproducth·,·iL2,ε.Ifε=idinthesedecompositions,weomitit,i.e.,wewrite⊕insteadof⊕id.ThePoincar´eandtangentialMaxwellestimatesreadas∀ϕ∈H1∩L2Ω0|ϕ|L2≤cp|∇ϕ|L2,∀φ∈˚R∩ε−1D|φ|L2,ε≤cm,εq
|divεφ|2
L2+|rotφ|2
L2,wherecp=cp(Ω)>0isthePoincar´econstantandcm,ε=cm,ε(Ω,ε)>0thetangentialMaxwellconstant.Forconvexdomainswehavethecomputableupperbound(3.3)cp≤diam(Ω)
πbyPayneandWeinberger[15](seealso[5]).In[13](seealso[4,12,14])itwasshownthat(3.4)cm,ε≤maxcf
√
ε
,√
εcp,
ASHORTNOTEONAWEIGHTEDFRIEDRICHSINEQUALITY9which,togetherwith(1.1)and(3.3),givesthecomputableupperbound(3.5)cm,ε≤1
πmax(s
ε
1
l21+1
l22+1
l23−1,√
εdiam(Ω)).Asε
approacheszero,thisboundblowsup,asdoes(1.4).Inthefollowingweimprovethisboundsothatthisdoesnothappen.Therestofthissectionessentiallyfollowsthesequencefoundin[13],withafewsmalldiﬀerencestoallowfortheimprovedbound.Lemma3.1.LetΩ⊂R3beconvexandφ∈˚R∩Dorφ∈R∩˚D.Thenφ∈H1and|∇φ|2
L2≤|divφ|2
L2+|rotφ|2
L2.Proof.See[3,Thm.2.17].Lemma3.2.LetΩ⊂Rd.Theinclusion˚D0⊂L2Ω0holds.Proof.Letφ∈˚D0.Thenhφi,1iL2=hφ,∇xiiL2=−hdivφ,xiiL2=0foranyi∈{1,...,d}.Lemma3.3.LetΩ⊂R3beconvex,φ∈˚R0∩ε−1Dandψ∈˚R∩ε−1D0.Thentheestimates|φ|L2,ε≤cf,ε|divεφ|L2and|ψ|L2,ε≤√
εcp|rotψ|L2hold.Proof.Weprovetheﬁrstestimatefollowing[13].Letφbelongto˚R0∩ε−1D.Since∇˚H1=˚R0,thereexistsascalarpotentialϕ∈˚H1suchthatφ=∇ϕ,andwehave|φ|2
L2,ε=hεφ,φiL2=hεφ,∇ϕiL2=−hdivεφ,ϕiL2≤|divεφ|L2|ϕ|L2≤cf,ε|divεφ|L2|∇ϕ|L2,ε,wherewehaveapplied(1.2)withα=ε.Thisprovestheﬁrstestimate.Thesecondestimateweproveinthewaypresentedin[4].Letψbelongto˚R∩ε−1D0.From(3.1)–(3.2)wededuceD0=rotR=rot(R∩˚D0).ThelatteridentityisseenbydecomposingRusingL2=∇H1⊕˚D0.Thus,sinceεψ∈D0,thereexistsavectorpotentialξ∈R∩˚D0suchthatεψ=rotξ.WithLemmas3.1and3.2weseethatξalsobelongstoH1∩L2Ω0,andwecanwrite|ψ|2
L2,ε=hεψ,ψiL2=hrotξ,ψiL2=hξ,rotψiL2≤|ξ|L2|rotψ|L2≤cp|∇ξ|L2|rotψ|L2=cp|rotξ|L2|rotψ|L2=cp|εψ|L2|rotψ|L2≤√
εcp|ψ|L2,ε|rotψ|L2,whereweusedthePoincar´eestimateandLemma3.1.Thisprovesthesecondestimate.Remark3.4.ThesecondestimateinLemma3.3wasprovenwithaslightlybetterconstantin[4].ThiswasachievedbyreﬁningtheglobalzeromeanpropertyofLemma3.2intoalocalequivalent.Weskipthisimprovementinthisnote,sinceitdoesnotchangetheoverallbehaviouroftheforthcomingcomputableestimatewithrespecttosmallε
.Wenowsatethereﬁnementof(3.4).Theorem3.5.LetΩ⊂R3beconvexandε∈L∞beaself-adjointuniformlypositivedeﬁnitematrix.Thencm,ε≤max{cf,ε,√
εcp}holds.Proof.Weﬁrstuse(3.1)–(3.2)todecomposethespace˚R∩ε−1D.Bydecomposing˚RusingL2=˚R0⊕εε−1D0weobtain˚R=˚R0⊕ε(˚R∩ε−1D0).Thus,wededuce˚R∩ε−1D=(˚R0∩ε−1D)⊕ε(˚R∩ε−1D0).Usingtheabovedecompositionwecanwriteφ∈˚R∩ε−1Dasthesumφ=φ1+φ2,whereφ1∈˚R0∩ε−1D,φ2∈˚R∩ε−1D0,andfurthermore,divεφ=divεφ1androtφ=rotφ2hold.ByusingLemma3.3wethenhave|φ|2
L2,ε=|φ1|2
L2,ε+|φ2|2
L2,ε≤c2Ωf,ε|divεφ1|2
L2+
εc2Ωp|rotφ2|2
L2=c2Ωf,ε|divεφ|2
L2+
εc2Ωp|rotφ|2
L2,fromwhichtheassertionfollows.
10IMMANUELANJAMANDDIRKPAULYTheimprovementof(3.5),inthecaseofdiagonalε,thenlookslikefollows.Theorem3.6.LetΩ⊂R3beconvexandε∈L∞beaself-adjointdiagonalmatrixsuchthatproperties(2.2)–(2.3)withα=εhold.Thenwehavetheestimatecm,ε≤1
πmax(s
ε
1
l21+ε
2
l22+ε
3
l23−1,√
εdiam(Ω)).Proof.Theorem3.5togetherwithTheorem2.2and(3.3)resultintheassertion.Remark3.7.(i)ItiseasytoseethattheupperboundofTheorem3.6isalwayssmallerorequaltotheupperbound(3.5).Moreover,withεaconstantrealnumber,theseestimatescoincide.(ii)Theaboveprocedurefurnishesupperboundsofcm,εevenwhenthediagonalmatrixεisnotuniformlypositivedeﬁnite(seeAppendixA).Undercertainconditionsnon-diagonalεcanbehandledusingTheorem2.4.Below˜εisthediagonalmatrixrelatedtoεsatisfyingtheproperties(2.4)–(2.5).Theorem3.8.LetΩ⊂R3beconvexandε∈L∞beaself-adjointmatrixvaluedfunctionforwhich˜εisuniformlypositivedeﬁnite.Thenwehavetheestimatecm,ε≤1
πmax(s
˜ε
1
l21+˜ε
2
l22+˜ε
3
l23−1,√
εdiam(Ω)).Proof.Theorem3.5togetherwithTheorem2.4and(3.3)resultintheassertion.Remark3.9.For˜εtobeuniformlypositivedeﬁnitewouldrequirethattheoﬀ-diagonalentriesofεbecomparativelysmallcomparedtoitsdiagonalentries.However,nowRemark3.7(ii)holdswithrespectto˜ε.Inparticular,foranupperboundofcm,εitisenoughthat˜εispositivesemi-deﬁnitesuchthatoneofthediagonalentriesof˜εisuniformlypositivedeﬁnite.Weconcludewithanexample.Example5(Diagonalmatrixε).LetΩ⊂(0,1)3andεbetheuniformlypositivedeﬁniteconstantmatrixε=
100Ω010Ω00δ
,δ>0.Theestimate(3.5)givestheupperboundΩ(3.6)cm,ε≤√
3
πmax1
3√
δ,1,andTheorem3.6givesΩ(3.7)cm,ε≤√
3
π.Obviouslythelatterdoesnotblowupasδbecomessmaller.Table3showsthevaluesoftheboundswithdiﬀerentδ.Table3.Example5:Valuesoftheupperbounds(3.6)and(3.7)withdiﬀerentδ.δ
10−6
10−4
10−2
1
102
104
106
(3.6)
183.77630
18.37763
1.83776
0.55133
0.55133
0.55133
0.55133
(3.7)
√
3/π≈0.55133
ASHORTNOTEONAWEIGHTEDFRIEDRICHSINEQUALITY11References[1]R.A.Adams.Sobolevspaces.AcademicPress,1975.[2]M.AinsworthandJ.Oden.Aposteriorierrorestimationinﬁniteelementanalysis.JohnWiley&Sons,2000.[3]C.Amrouche,C.Bernardi,M.Dauge,andV.Girault.Vectorpotentialsinthree-dimensionalnon-smoothdomains.Math.MethodsAppl.Sci.,21(9):823–864,1998.[4]I.Anjam.AshortnoteonHelmholtzdecompositionsforboundeddomainsinR3.Math.Scand.,2018.Accepted.[5]M.Bebendorf.AnoteonthePoincar´einequalityforconvexdomains.Z.Anal.Anwendungen,22(4):751–756,2003.[6]P.Grisvard.Ellipticproblemsinnonsmoothdomains.Pitman,Boston,1985.[7]N.KuznetsovandA.Nazarov.SharpconstantsinthePoincar´e,Steklovandrelatedinequalities(asurvey).Mathematika,pages1–17,May2015.[8]R.Leis.Initialboundaryvalueproblemsinmathematicalphysics.Teubner,Stuttgart,1986.[9]O.Mali,P.Neittaanm¨aki,andS.Repin.Accuracyveriﬁcationmethods,theoryandalgorithms.Springer,2014.[10]S.G.Mikhlin.Constantsinsomeinequalitiesofanalysis.Wiley,1986.[11]P.Neittaanm¨akiandS.Repin.Reliablemethodsforcomputersimulation,errorcontrolandaposterioriestimates.Elsevier,2004.[12]D.Pauly.OntheMaxwellinequalitiesforboundedandconvexdomains.ZapiskiPOMI,435:46–54,2014.[13]D.Pauly.OnMaxwell’sandPoincar´e’sconstants.DiscreteContin.Dyn.Syst.Ser.S,8(3):607–618,2015.[14]D.Pauly.OntheMaxwell’sconstantsin3D.MathematicalMethodsintheAppliedSciences,40(2):435–447,2017.[15]L.E.PayneandH.F.Weinberger.AnoptimalPoincar´einequalityforconvexdomains.Arch.RationalMech.Anal.,5:286–292,1960.[16]S.Repin.Aposterioriestimatesforpartialdiﬀerentialequations.DeGruyter,2008.[17]S.Repin.ComputablemajorantsofconstantsinthePoincar´eandFriedrichsinequalities.J.Math.Sci.(N.Y.),186(2):307–321,2012.[18]T.Vejchodsk´y.ComputingupperboundsonFriedrichs’constant.InJ.Brandts,J.Chleboun,S.Korotov,K.Segeth,J.ˇS´ıstek,andT.Vejchodsk´y,editors,ProceedingsoftheInternationalConferenceApplicationsofMathematics2012,pages278–289,2012.[19]R.Verf¨urth.Aposteriorierrorestimationtechniquesforﬁniteelementmethods.OxfordUniversityPress,2013.[20]N.Weck.Maxwell’sboundaryvalueproblemsonRiemannianmanifoldswithnonsmoothboundaries.J.Math.Anal.Appl.,46(2):410–437,1974.AppendixA.PositiveSemi-DefiniteMatricesαWeshortlydemonstratethatupperboundsforcf,αcanbeobtainedeveniftheself-adjointα∈L∞isonlypositivesemi-deﬁnite:LetΩ⊂(−1,1)3andαbetheconstantmatrixα=
000Ω000Ω001
.Thesmallesteigenvalueisα
=0,sotheestimate(1.4)cannotbeused.However,asstatedinRemark2.3(iii),wecanstillobtainupperboundsforcf,α;wecanapplyLemma2.1withi=3toobtain|ϕ|L2≤2
π|∂3ϕ|L2=2
π|∇ϕ|L2,α⇒cf,α≤2
π.Thequantity|∇·|L2,αisnownotanorm,though.E-mailaddress,ImmanuelAnjam:immanuel.anjam@gmail.comFacultyofMathematics,UniversityofDuisburg-Essen,CampusEssen,GermanyE-mailaddress,DirkPauly:dirk.pauly@uni-due.de

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

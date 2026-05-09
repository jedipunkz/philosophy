---
source: "https://arxiv.org/abs/math/0408128v1"
title: "Dual random fragmentation and coagulation and an application to the genealogy of Yule processes"
author: "Jean Bertoin, Christina Goldschmidt"
year: "2004"
publication: "arXiv preprint / math.PR"
download: "https://arxiv.org/pdf/math/0408128v1"
pdf: "https://arxiv.org/pdf/math/0408128v1"
captured_at: "2026-05-09T12:58:02Z"
updated_at: "2026-05-09T12:58:02Z"
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

# Dual random fragmentation and coagulation and an application to the genealogy of Yule processes

- 著者: Jean Bertoin, Christina Goldschmidt
- 年: 2004
- 掲載情報: arXiv preprint / math.PR
- 情報源: [arxiv](https://arxiv.org/abs/math/0408128v1)
- ダウンロード: https://arxiv.org/pdf/math/0408128v1
- PDF: https://arxiv.org/pdf/math/0408128v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

The purpose of this work is to describe a duality between a fragmentation associated to certain Dirichlet distributions and a natural random coagulation. The dual fragmentation and coalescent chains arising in this setting appear in the description of the genealogy of Yule processes.

## PDF Text

arXiv:math/0408128v1 [math.PR] 10 Aug 2004
DualrandomfragmentationandcoagulationandanapplicationtothegenealogyofYuleprocessesJeanBertoin∗andChristinaGoldschmidt†August22,2018AbstractThepurposeofthisworkistodescribeadualitybetweenafragmentationassociatedtocertainDirichletdistributionsandanaturalrandomcoagulation.ThedualfragmentationandcoalescentchainsarisinginthissettingappearinthedescriptionofthegenealogyofYuleprocesses.
1Introduction
Atanaivelevel,fragmentationandcoagulationareinversephenomena,inthatasimpletime-reversalchangesoneintotheother.However,stochasticmodelsforfragmentationandcoalescenceusuallyimposestronghypothesesonthedynamicsoftheprocesses,suchasthebranchingpropertyforfragmentation(distinctfragmentsevolveindependentlyastimepasses),andtheserequirementsdonottendtobecompatiblewithtime-reversal.Thus,ingeneral,thetime-reversalofacoalescentprocessisnotafragmentationprocess.Nonetheless,thereareafewspecialcasesinwhichtime-reversaldoestransformacoalescentprocessintoafragmentationprocess.ProbablythemostimportantexamplewasdiscoveredbyPitman[17];itisrelatedtotheso-calledcascadesofRuelleandtheBolthausen-Sznitmancoalescent[7],andalsohasanaturalinterpretationintermsofthegenealogyofaremarkablebranchingprocessconsideredbyNeveu,see[4]and[6].Theﬁrstpurposeofthisnoteistopointoutothersimpleinstancesofsuchduality,whichrelyoncertainDirichletandPoisson-Dirichletdistributions.Then,inthesecondpart,weshallshowthattheseexamplesarerelatedtothegenealogyofYuleprocesses.
∗LaboratoiredeProbabilit´esetMod`elesAl´eatoiresandInstitutuniversitairedeFrance,Universit´ePierreetMarieCurie,175,rueduChevaleret,F-75013Paris,France.†LaboratoiredeProbabilit´esetMod`elesAl´eatoires,Universit´ePierreetMarieCurie,175,rueduChevaleret,F-75013Paris,France.1
2Dualfragmentationandcoagulation
2.1Somenotation
Foreveryintegern≥1,weconsiderthesimplex∆n:=(x=(x1,...,xn+1):xi≥0foreveryi=1,...,n+1andn+1Xi=1xi=1).Itwillalsobeconvenienttoagreethat∆0:={1}.Weshalloftenrefertothecoordinatesx1,...,xn+1ofpointsxin∆nasmasses.Werecallthatthen-dimensionalDirichletdistributionwithparameter(α1,...,αn+1)istheprobabilitymeasureonthesimplex∆nwithdensityΓ(α1+···+αn+1)
Γ(α1)···Γ(αn+1)xα1−11···xαn+1−1n+1.Thespecialcasewhenα1=...=αn+1:=α∈]0,∞[willhaveanimportantroleinthiswork;itwillbeconvenienttowriteDirn(α)forthisdistribution.Werecallthefollowingwell-knownconstruction:letγ1,...,γn+1bei.i.d.gammavariableswithparameters(α,c).Set¯γ=γ1+···+γn+1,sothat¯γhasagammadistributionwithparameters(α(n+1),c).Thenthe(n+1)-tuple(γ1/¯γ,...,γn+1/¯γ)hasthedistributionDirn(α)andisindependentof¯γ.Wealsodeﬁnethe(ranked)inﬁnitesimplex∆∞:=(x=(x1,...):x1≥x2≥...≥0and∞Xi=1xi=1)andrecallthatthePoisson-Dirichletdistributionwithparameterθ>0,whichwillbedenotedbyPD(θ)inthesequel,isthelawoftherandomsequenceξ:=a1
P∞
i=1ai,a2
P∞
i=1ai,...,wherea1≥a2≥...>0aretheatomsofaPoissonrandommeasureon]0,∞[withinten-sityθy−1e−ydy.WealsorecallthatξisindependentofP∞
i=1ai,andthatthelatterhasthegammadistributionwithparameters(θ,1).BythecelebratedL´evy-Itˆodecompositionofsubordinators,wemayalsorephrasethisconstructionasfollows:ifγ=(γ(t),t≥0)isastandardgammaprocessand,foreachﬁxedθ>0,δ1≥δ2≥...denotesthesequenceofsizesofthejumpsofγonthetimeinterval[0,θ],thenδ1
γ(θ),δ2
γ(θ),...hasthePD(θ)distributionandisindependentofγ(θ).2
2.2Twodualrandomtransformations
Wenowdeﬁnetworandomtransformations:Fragk:∆n→∆n+kandCoagk:∆n+k→∆n,wherek,nareintegers.First,weﬁxx=(x1,...,xn+1)∈∆nandpickanindexI∈{1,...,n+1}atrandomaccordingtothedistributionP(I=i)=xi,i=1,...,n+1,sothatxIisasize-biasedpickfromthesequencex.Letη=(η1,...,ηk+1)bearandomvariablewithvaluesin∆kwhichisdistributedaccordingtoDirk(1/k)andindependentofI.ThenwesplittheIthmassofxaccordingtoηandweobtainarandomvariablein∆n+k:Fragk(x):=(x1,...,xI−1,xIη1,...,xIηk+1,xI+1,...,xn+1).Second,weﬁxx=(x1,...,xn+k+1)∈∆n+kandpickanindexJ∈{1,...,n+1}uniformlyatrandom.Wemergethek+1massesxJ,xJ+1...,xJ+ktoformasinglemassPJ+ki=Jxiandleavetheothermassesunchanged.Weobtainarandomvariablein∆n:Coagk(x)= x1,...,xJ−1,J+kXi=Jxi,xJ+k+1,...,xn+k+1!.Remark.Considerthefollowingalternativerandomcoagulationofx=(x1,...,xn+k+1)∈∆n+k.Pickk+1indicesi1,...,ik+1from{1,...,n+k+1}uniformlyatrandomwithoutreplacement,mergethemassesxi1,...,xik+1,leavetheothermassesunchangedandlet]Coagk(x)bethesequenceobtainedbyrankingtheresultingmassesindecreasingorder.WritealsoCoag↓
k(x)forthesequenceCoagk(x)re-arrangedindecreasingorder.Thenifξisexchangeablethepairs(ξ,Coag↓
k(ξ))and(ξ,]Coagk(ξ))havethesamedistribution.ThisremarkappliesinparticulartothecasewhenξhasthelawDirn+k(1/k),andcanthusbecombinedwithforthcomingProposition1.ThestartingpointofthisworkliesintheobservationofasimplerelationofdualitywhichlinksthesetworandomtransformationsviaDirichletlaws.Proposition1.Letk,n≥1betwointegers,andξ,ξ0tworandomvariableswithvaluesin∆nand∆n+k,respectively.Thefollowingassertionsarethenequivalent:(i)ξhasthelawDirn(1/k)and,conditionallyonξ,ξ0isdistributedasFragk(ξ).(ii)ξ0hasthelawDirn+k(1/k)and,conditionallyonξ0,ξisdistributedasCoagk(ξ0).IthasbeenobservedbyKingman[13]thatfork=1,ifξ0isuniformlydistributedonthesimplex∆n+1(i.e.hasthelawDirn+1(1)),thenCoag1(ξ0)isuniformlydistributedon∆n.Clearly,thisagreeswithourstatement.Proof:Letγ1,γ2,...,γn+1beindependentGamma(1/k,1)randomvariablesandset¯γ=n+1Xi=1γiandξ=γ1
¯γ,...,γn+1
¯γ,sothatξhaslawDirn(1/k)andisindependentof¯γ.SupposethatηisaDirk(1/k)randomvariablewhichisindependentoftheγi’s,andletΦ:Rn+k+1→Rbeabounded3
measurablefunction.LetIbeanindexpickedatrandomfrom{1,...,n+1}accordingtotheconditionaldistributionP(I=i|γ1,...,γn+1)=γi/¯γ,i=1,...,n+1,anddenotebyFragk(ξ)therandomsequenceobtainedfromξafterthefragmentationofitsIthmassaccordingtoη.WehaveE(Φ(Fragk(ξ)),I=i)=Eγi
¯γΦ((γl/¯γ)l<i,γiη/¯γ,(γl/¯γ)l>i).Now,usingtheindependenceof¯γandξandthefactthat¯γhasthelawGamma((n+1)/k,1),weseethatthelastexpressionisequaltok n+1E[γiΦ((γl/¯γ)l<i,γiη/¯γ,(γl/¯γ)l>i)]=k n+1EZ∞0xΦ (γl)l<i x+Pj6=iγj,xη
x+Pj6=iγj,(γl)l>i x+Pj6=iγj!1
Γ(1/k)x1/k−1e−xdx=1
n+1E"Φ (γl)l<i
γ0+Pj6=iγj,γ0η
γ0+Pj6=iγj,(γl)l>i
γ0+Pj6=iγj!#whereγ0∼Gamma((k+1)/k,1),independentlyofηand(γj)j6=i.Butthenγ0ηisacollectionofk+1independentGamma(1/k,1)randomvariables,soFragk(ξ)hasthelawDirn+k(1/k)andisindependentoftherandomindexIwhichisuniformlydistributedon{1,...,n+1}.SincewecanrecoverξfromFragk(ξ)andIbyanobviouscoagulation,thiscompletestheproof.Nextweturnourattentiontotheinﬁniterankedsimplexanddeﬁnetworandomtransformations,Frag∞:∆∞→∆∞andCoaga:∆∞→∆∞,wherea∈[0,1]issomeparameter.Thefragmentationtransformationontheinﬁnitesimplexsimplymim-icsthatontheﬁnitesimplex;inthisdirection,recallthatthePoisson-DirichletPD(1)arisesastheweaklimitask→∞ofsequenceofDirk(1/k)variablesafterobviousre-ordering.Moreprecisely,givenx=(x1,...)∈∆∞,wepickamassxIatrandombysize-biasedsamplingandsplitxIusinganindependentvariableη=(η1,...)withlawPD(1).Inotherwords,Frag∞(x)istherankedsequencewithunorderedtermsx1,...,xI−1,xIη1,xIη2,...,xI+1,....Next,considerasequenceU1,U2,...ofi.i.d.uniformrandomvariablesanda∈[0,1].Startingagainfromsomeﬁxedx∈∆∞,wemergethemassesxiforwhichUi≤aintoasinglemassandleavetheothersunchanged.WedenotebyCoaga(x)therandomsequenceobtainedbyputtingtheresultingmassesindecreasingorder.WethenhavethefollowinganalogueofProposition1,whichisreminiscentofCorollary13ofPitman[17].Proposition2.Letξ,ξ0betworandomvariableswithvaluesin∆∞.Foreveryθ>0,thefollowingassertionsareequivalent:(i)ξhasthelawPD(θ)and,conditionallyonξ,ξ0isdistributedasFrag∞(ξ).(ii)ξ0hasthelawPD(θ+1)and,conditionallyonξ0,ξisdistributedasCoag1/(θ+1)(ξ0).Proof:Letγ=(γ(t),t≥0)beastandardgammaprocessandsetDt=γ((θ+1)t)/γ(θ+1),4
for0≤t≤1,sothat(Dt,0≤t≤1)isaDirichletprocessofparameterθ+1.(ThevectoroforderedjumpsofthisDirichletprocesshasthePD(θ+1)distribution.)ConsiderthefollowingalternativewayofthinkingoftherandomcoagulationoperatorCoag1/(θ+1):pickapointVuniformlyin[0,1]anddeﬁneanewprocess(D0t,0≤t≤1)byD0t=(Dθt/(θ+1)ift<VD(1+θt)/(θ+1)ift≥V.AsthetimesofthejumpsofDareuniformlydistributedon[0,1],thispicksaproportion1/(θ+1)ofthemandcoalescesthemintoasinglejump(sayβ∗=D(1+θV)/(θ+1)−DθV/(θ+1))atV.Letβ1≥β2≥...>0bethesequenceofotherjumpsofD0andU1,U2,...thecorrespondingjumptimes.Letβ01≥β02≥...>0bethesequenceofjumpsofDintheinterval[θV/(θ+1),(1+θV)/(θ+1)],sothatβ∗=P∞
i=1β0i.WewishtoshowthatD0isaDirichletprocesswithparameterθ,sothatthevector(β∗,β1,β2,...)ofitsjumps(re-arrangedinthedecreasingorder)hasthePD(θ)distribution.Wewillalsoshowthatthemassβ∗resultingfromthecoalescenceconstitutesasize-biasedpickfromthisvector.Letγ1(t)=(γ(t)ift<Vθγ(t+1)−(γ(Vθ+1)−γ(Vθ))ifVθ≤t≤θγ2(t)=γ(Vθ+t)−γ(Vθ)for0≤t≤1.Thenγ1andγ2areindependentprocesseswithγ1d=(γ(t),0≤t≤θ)andγ2d=(γ(t),0≤t≤1),independentlyofV.Writeδ1≥δ2≥...fortheorderedsequenceofjumpsofγ1andT1,T2,...forthecorrespondingtimesofthesejumps.Writeδ01≥δ02≥...fortheorderedsequenceofjumpsofγ2.Then(i)U1=T1/θ,U2=T2/θ,...arei.i.d.U[0,1],(ii)β∗=γ2(1)/γ(1+θ)andsohasaBeta(θ,1)distribution,(iii)1
β∗(β01,β02,...)=1
γ2(1)(δ01,δ02,...)andsohasthePD(1)distribution,(iv)1
1−β∗(β1,β2,...)=1
γ1(θ)(δ1,δ2,...)andsohasthePD(θ)distribution.Furthermore,therandomvariablesin(i)to(iv)aboveareindependent.Thefactthatβ∗isasize-biasedpickfrom(β∗,β1,β2,...)andthePD(θ)distributionofthelatterfollowfrom(i)and(iii)andthestick-breakingscheme(see,forinstance,Deﬁnition1inPitmanandYor[19]).ThatD0isaDirichletprocessofparameterθthenfollowsfrom(iv)andtheindependence.Thecoagulationoperatorusedherecanbere-phrasedasfollows:startingwithx∈∆∞,takeasequenceV,V1,V2,...ofi.i.d.U[0,1]randomvariables,mergethemassesxiforwhichVi∈[θV/(θ+1),(1+θV)/(θ+1)],leavetheothermassesunchangedand,ﬁnally,ranktheresultingsequenceindecreasingorder.Callthisoperator]Coag1/(θ+1).Thenitisclearthatwheneverξ0isarandomexchangeablesequencein∆∞,(ξ0,Coag1/(θ+1)(ξ0))and(ξ0,]Coag1/(θ+1)(ξ0))havethesamedistribution.Ourclaimfollowsnowreadilyfromtheseresults.Remark.ItmaybeinterestingtocheckProposition2asfollows.ConsiderPoissonrandommeasureMon(0,∞)withintensityθx−1e−xdx.Leta1,a2,...betheatomsof5
Mindecreasingorder,sothat a1
P∞
j=1aj,a2
P∞
j=1aj,...!hasdistributionPD(θ),independentlyofP∞
j=1aj.Letη∼PD(1),independentlyofMandsupposethatΦ:∆∞→Risanysymmetricboundedmeasurablefunction.Thenifξ∼PD(θ),usingindependencewehavethatE[Φ(Frag∞(ξ))]=1
EhP∞
j=1ajiE"∞Xi=1aiΦ aiη
P∞
j=1aj,(al)l6=i
P∞
j=1aj!#.BythePalmformula,thisisequalto1
θEZ∞0xΦ xη
x+P∞
j=1aj,(al)∞
l=1
x+P∞
j=1aj!θx−1e−xdx=E"Φ a0η
a0+P∞
j=1aj,(al)∞
l=1
a0+P∞
j=1aj!#wherea0∼Exp(1),independentlyofMandη.Butthena0ηhasthedistributionoftheatomsofaPoissonrandommeasurewithintensityx−1e−xdxarrangedindecreasingorderandsoweseethattakingtheseatomstogetherwiththoseofM,weobtainaPoissonrandommeasureofintensity(θ+1)x−1e−1dx.Hence,Frag∞(ξ)hasthelawPD(θ+1).2.3Dualfragmentationandcoagulationchains
ThedualfragmentationandcoagulationoperatorsthatweredeﬁnedintheprecedingsectioninciteustointroduceMarkovfragmentationandcoagulationchainsindualitybytime-reversal.Speciﬁcally,weconsiderforeachintegerk≥1achainX(k)(0),X(k)(1),X(k)(2),...,whereX(k)(n)isarandomvariablewithvaluesin∆nk(inparticularX(k)(0)=1),andtheconditionaldistributionofX(k)(n+1)givenX(k)(n)=xisthelawofFragk(x).WededucefromProposition1byinductionthatforeachn,X(k)(n)hasthedistributionDirnk(1/k).Thetime-reversedcoagulationchain...,X(k)(n+1),X(k)(n),...,X(k)(1),X(k)(0)isalsoMarkov;moreprecisely,theconditionaldistributionofX(k)(n)givenX(k)(n+1)=xisthelawofCoagk(x).Notethatfork=1,thishasthedistributionofthejumpchainofKingman’scoalescent[13].Analogously,fork=∞,wecandeﬁneaMarkovfragmentationchainon∆∞,X(∞)(0),X(∞)(1),X(∞)(2),...,suchthattheconditionaldistributionofX(∞)(n+1)givenX(∞)(n)=xisthelawofFrag∞(x).WededucebyinductionfromProposition2thatforeveryθ>0,ifthe6
distributionoftheinitialstateX(∞)(0)isPD(θ)then,foreveryintegern,X(∞)(n)hasthedistributionPD(θ+n).Moreover,inthiscase,thetime-reversedcoagulationchain...,X(∞)(n+1),X(∞)(n),...,X(∞)(1),X(∞)(0)isalsoMarkov;moreprecisely,theconditionaldistributionofX(∞)(n)givenX(∞)(n+1)=xisthelawofCoag1/(n+1+θ)(x).Remarks.(a)RecallthattheparameterθcanberecoveredfromasampleξofaPD(θ)randomvariableasfollows:θ=limε→0+1
log1/εmax{n:ξn>ε}.ThisshowsthattheabovedescriptionforthereversedcoagulationchainisindeedMarko-vian.(b)Thereissimplerepresentationforthek=∞fragmentationchainintermsofcompoundbridgeswithexchangeableincrementswhichisinspiredby[5].LetU0,U1,...beasequenceofindependentuniformvariableson[0,1].Foreachn,weconsidertheelementarybridgebn:[0,1]→[0,1]deﬁnedbybn(t)=n n+1t+1
n+11{t>Un},t∈[0,1].Thenisiseasytocheckthatforeveryn∈N,thesequencebn◦bn+1◦···◦bn+iconvergespointwisealmostsurelyasi→∞toabridgewithexchangeableincrementsBnwhichhasnodriftandinﬁnitelymanyjumpsa.s.Ifwewriteβn∈∆nforthesequenceofthesizesofthejumpsofBnrankedthedecreasingorder,thenthechain(βn,n∈N)hasthesamelawasX(∞).Wereferto[5]forthenecessarytechnicalbackground.3ThegenealogyofYuleprocesses
Weshallnowshowthatthedualfragmentationandcoagulationchainswhichweintro-ducedintheprecedingsectionarenaturallyconnectedtothegenealogyofYuleprocesses.3.1Discretesetting
Foreveryintegerk≥1,wewriteY(k)=Y(k)t,t≥0fortheYuleprocessstartedfromY(k)0=1:Y(k)tgivesthenumberofindividualsaliveatimetinabranchingprocessinwhicheachindividuallivesforanexponentialtimeofparameter1andgivesbirthatitsdeathtok+1children,whichthenevolveindependentlyofoneanotheraccordingtothesamerulesastheirparent.Weagreetolabeleachchildofanindividualbyanintegerin{1,...,k+1},whichallowsustoorderindividualsatanygenerationinaconsistentway:giventwodistinctindividuals,wemayconsidertheirmostrecentcommonancestor.Plainly,twodiﬀerentchildrenofthisancestorareancestorsofexactlyoneofthesetwoindividuals,andthelabellingofthechildrenofthemostrecentcommonancestorinducestheorderoftheindividuals.
Lemma3.Theprocessexp(−kt)Y(k)t,t≥0isauniformlyintegrablemartingaleanditslimitW(k)hastheGamma(1/k,1/k)distribution.7
Proof:AsimilarlimitresultisstatedinAthreya&Ney[1]onpage130;however,thelimitingdistributiongiventhereisincorrectandsoweshallprovidehereadetailedproof.Themartingalepropertyisclassical,sowefocusonthedistributionofthelimitW(k).DeﬁneΦt(s):=EsY(k)t.Thebackwardequationimpliesthat∂
∂tΦt(s)=Φk+1t(s)−Φt(s),Φ0(s)=s.ThisequationhassolutionΦt(s)=se−t1−1−e−ktsk−1/k.Hence,forθ<0,Eexpθe−ktY(k)t=expθe−kte−t1−1−e−ktexpθke−kt−1/k=hektexp−θke−kt−ekt+1i−1/k,andwhent→∞,thisquantityconvergesto(1−kθ)−1/k=1/k
1/k−θ1/k,whichisthemomentgeneratingfunctionofagammarandomvariablewithparameters(1/k,1/k).WethinkofW(k)asthesizeoftheterminalpopulation.Foreveryt≥0,byapplicationofthebranchingpropertyattimet,wemaydecomposetheterminalpopulationintosub-populationshavingthesameancestorattimet.Speciﬁcally,W(k)=Y(k)tXi=1W(k)i(t),whereW(k)i(t)isthesizeoftheterminalsub-populationdescendingfromtheithindividualinthepopulationattimet.ObservethatconditionallyonY(k)t,thevariablesW(k)i(t)areindependentandallhavethesamelawase−ktW(k).Finally,wedeﬁnethegenealogicalprocessG(k)= G(k)(t),t≥0associatedtoY(k)byG(k)(t)=W(k)1(t),...,W(k)Y(k)t(t).ThegenealogicalstructureoftheYuleprocesscanbedescribedintermsofthefrag-mentationchainX(k)ofSection2.3asfollows.Theorem4.LetN=(Nt,t≥0)beastandardPoissonprocesswhichisindependentofthechainX(k).Thenforeachw>0,thecompoundchainwX(k)(Nwt),t≥08
hasthesamelawasthetime-changedprocessG(k)1
klog(1+kt),t≥0conditionedonW(k)=w.Remark.TheoremIofKendall[12]statesthatgivenW(1),Y(1)log(1+t/W(1)),t≥0isaPoissonprocesswithunitparameter.ThisisclearlyanaspectofTheorem4.Moreover,onpage130ofAthreya&Ney[1],itissuggestedthatnogeneralizationofKendall’sresulttoamoregeneralcontinuous-timeMarkovbranchingprocessisknown;Theorem4constitutesasmallsuchgeneralization.
Proof:Setτ(t):=1
klog(1+kt)andletTbethetimeoftheﬁrstbirthintheYuleprocess,whichisalsothetimeoftheﬁrstdislocationofG(k).Thek+1fragmentsofG(k)(T)canbewrittenase−kTZ1,...,e−kTZk+1where,bythebranchingproperty,Z1,...,Zk+1arei.i.d.Gamma(1/k,1/k)randomvariables,independentofTwhichisExp(1).DeﬁneachangeofvariablesbyS=τ−1(T)=(ekT−1)/kU1=e−kTZ1,...,Uk=e−kTZk,W=e−kT(Z1+···+Zk+1).Itisstraightforwardtoseethatthejointdensityof(T,Z1,...,Zk+1)isf(t,z1,...,zk+1)=e−tΓ(1/k)−(k+1)(1/k)(k+1)/k(z1z2...zk+1)−(k−1)/kexp(−(z1+···+zk+1)/k)andsothejointdensityof(S,U1,...,Uk,W)isg(s,u1,...,uk,w)=we−ws·(1/k)Γ(1/k)−kw−1/k(u1u2...uk(w−u1−···−uk))1/k−1·(1/k)1/kΓ(1/k)−1w−(k−1)/kexp(−w/k).Hence,W∼Gamma(1/k,1/k)(aswealreadyknew)and,conditionalonW=w,wehaveS∼Exp(w)and(U1,U2,...,Uk,W−U1−···−Uk)∼wDirk(1/k)independentlyofS.Thus,theﬁrstdislocationhasthecorrectdynamics.Butbythebranchingproperty,subsequentdislocationsareindependentfordiﬀerentsub-populationsandthetotalrateoffragmentationisalwaysw.Henceresult.Intheterminologyof[2],Theorem4statesthatthetime-changedgenealogicalprocessG(k)◦τisaself-similarfragmentationwithindex1,dislocationlawDirk(1/k)anderosioncoeﬃcient0.Itmaybeinterestingtoobservethatinthespecialcasek=1,thisresultcanalsobederivedasfollows.ConsiderarealBrownianmotionBstartedfrom1andkilledwhenitreaches0(attimeT0=inf{t≥0:Bt=0}).Foreveryu∈[0,1[,leteYudenotethenumberofexcursionsofBawayfrom1whichgobelowlevelu.Then(Y(1)−log(1−u))0≤u<1isaversionof(eYu)0≤u<1.Toseethis,letusconsidertheevolutionofeY.Firstly,eY0=1,correspondingtothesingleexcursionbelow1whichreaches0.LetD=sup{t<T0:Bt=1},thestartingtimeoftheﬁnalexcursionwhichhits0,letU=inf0≤t≤DBtbethelevelreachedbythedeepestexcursionbelow1beforeDandletTUbethetimeatwhichitisreached.Then,9
byWilliams’pathdecompositiontheorem(TheoremVII.4.9ofRevuzandYor[20]),Uisdistributeduniformlyon[0,1[and,conditionalonU,(Bt)0≤t<TUisaBrownianmotionstartedat1andstoppedwhenitﬁrsthitslevelU.Bysymmetry,(BD−t)0≤t<D−TUisanotherindependentBrownianmotionstartedat1andstoppedwhenitﬁrsthitslevelU.Thus,eYuisequalto1on[0,U[,eYU=2and(eYU+v)0≤v<1−UevolvesasthesumoftwoindependentprocesseswhicharethesameaseYexceptthatthetimesuntiltheﬁrstjumpsarenowuniformon[0,U[ratherthanon[0,1[.(ThisisTheorem8ofLeGall[16],repeatedhereforcompleteness.)Time-changingY(1)withu→−log(1−u)meansthatitsexponentialinter-jumptimesbecomeuniformandsowedo,indeed,have(Y(1)−log(1−u))0≤u<1d=(eYu)0≤u<1.AmoreelegantwayofexpressingtheprecedingargumentistosaythattheBrownianpathencodesacontinuous-statebranchingprocesswithquadraticbranchingmechanism.Thelocaltimeatlevel1,L1
T0,satisﬁesL1
T0=limu→1−2(1−u)eYu.Inthiscontext,1
2L1
T0correspondstothesizeofthepopulationattime1inthecontinuous-statebranchingprocessgeneratedbyasingleancestorconditionedtohavedescendentsuptotime1.Theso-calledreducedtreeassociatedwiththepopulationattime1isdescribeduptothedeterministictime-changeu→−log(1−u)bytheYuleprocessY(1).See,forinstance,Section2.7inDuquesneandLeGall[8],andFleischmannandSiegmund-Schultze[9].Notethatthewell-knownfactthat1
2L1
T0hasanexponentialdistributionwithmean1(PropositionVI.4.6ofRevuzandYor[20])givesanotherderivationofthelimitingdistributioninLemma3,sinceW(1)=limt→∞e−tY(1)t=limu→1−(1−u)Y(1)−log(1−u)d=1
2L1
T0.Itisknownfromexcursiontheorythatinthescaleofthelocaltimeatlevel1,therateofexcursionsofBawayfrom1whichreachlevelu∈]0,1[butdonotexceedu−duis(1−u)−2du.Notethatthemaps→1−1
1+sfromR+to[0,1[hasinverseu→1
1−u−1and,thus,transformsLebesguemeasureonR+intothemeasure(1−u)−2duon[0,1[.Supposethatwesplitthelocaltimeatlevel1accordingtotheoccurrenceofexcursionsexceedinglevelu,sothatweobtainthesequencefW(u)=fW1(u),...,fWeYu(u),wherefW(u)isthesequenceoftheincrementsofthelocaltimeatlevel1onthemaximaltimeintervalssuchthatatthebeginningandendofeachintervalBisat1andduringtheintervalremainsabovelevelu.Thenitfollowseasilythatthetime-changedprocessfW1−1
1+s,s≥0isafragmentationinwhicheachmass,sayx,splitsatratexintoxUandx(1−U)whereUisuniform.Inotherwords,conditionallyon1
2L1
T0=w,theprocessfW1−1
1+s,s≥0isdistributedasthecompoundfragmentationchain wX(1)(Nws),s≥0,whereNisanindependentstandardPoissonprocess.Finally,thecompositionofthetwotime-changeswhichappearinthisanalysisyieldss→−log1−1−1
1+s=log(1+s),s∈R+,10
andsowerecoverTheorem4inthespecialcasek=1.Unfortunately,itdoesnotseemthattherearesimilarinterpretationsfork≥2.Corollary5.Wehavethat1
W(k)G(k)1
klog1+ke−t/W(k),t∈Risatime-homogeneousMarkovcoalescentprocesswhichisindependentofW(k).Foranyn≥1,giventhatitisinstatex∈∆nk,itwaitsanexponentialtimeofparameternandthenjumpstoavariabledistributedasCoagk(x),independentlyoftheexponentialtime.Notethatthecasek=1ofthisresultgivesavariationofKingman’scoalescent.Thejump-chainsareidentical,aswehavealreadynoted,butheretherateofcoalescenceoftwoblocksdependsonthetotalnumberofblockspresent,whereasinKingman’scoalescentitdoesnot.
Proof:Firstly,wenotethatbyTheorem4,1
W(k)G(k)1
klog(1+ke−t/W(k)),t∈RhasthesamelawasX(k)(Ne−t),t∈Randsowewillworkwiththelatterprocessinstead.Thek=1caseisessentiallytreatedin[3]andtheproofproceedsinthesamewayhere.Thejumpchainclearlybehavesinthecorrectmannerandsoitremainstocheckthattheinter-jumptimesareasclaimed.Let0≤T1≤T2≤...bethejumptimesof(Nt)t≥0.ThentheﬁrstinstantthatX(k)(Ne−t)hasexactlynk+1termsisinf{t∈R:Ne−t=n}=−logTn+1.Thesequenceofinter-jumptimesis...,logTn+1−logTn,logTn−logTn−1,...,logT2−logT1anditiseasilyshownthatthisisasequenceofindependentexponentialrandomvariableswithparameters...,n,n−1,...,1respectively.3.2Continuoussetting
Continuous-statebranchingprocesses(orCSBP’s)wereintroducedbyLamperti[14,15]aslimitsofrescaledbranchingprocesses.Typically,aCSBPisatime-homogeneousMarkovprocesswithvaluesinR+,Z=(Z(t,a),t≥0anda≥0),(wheretheparametertreferstotimeandtheparameteratothestartingpointi.e.Z(0,a)=aa.s.)whichfulﬁlsthebranchingproperty:thepath-valuedprocess(Z(·,x),x≥11
0)hasindependentandstationaryincrements.Inparticular,ifeZ(·,y)isanindependentcopyofZ(·,y),thenZ(·,x)+eZ(·,y)hasthelawofZ(·,x+y).Thereisasimplere-lationconnectingCSBP’sandBochner’ssubordinationforsubordinatorswhichenablesustodeﬁnetheirgenealogy;werefertheinterestedreaderto[4]forheuristics,detailedargumentsetc.WecallacontinuousstateYuleprocessaCSBPY=(Y(t,a),t≥0anda≥0),whichevolvesasfollows:foreacha>0,theprocessY(·,a)waitsanexponentialtimewithparameteraandthenjumpstoa+1.Itthenevolvesindependentlyasifithadbeenstartedinstatea+1.Intermsofthegenealogy,thesub-populationofsize1whichisbornatajumptimehasaparentwhichischosenuniformlyatrandomfromthepopulationpresentbeforethejump.Notethatthisgenealogyiseasytodescribeinaconsistentmannerfordiﬀerentvaluesaofthestartingpopulation.Itisimmediatethatforanintegerstartingpointa∈N,theprocess(Y(t,a),t≥0)isaYuleprocessY(1)with2oﬀspring,asconsideredintheprecedingsection.However,westressthatitsgenealogyisnotthesameasthatofY(1),aswearedealingwithacontinuouspopulationintheﬁrstcaseandadiscretepopulationinthesecond.WehavethefollowinganalogueofLemma3:Lemma6.Foreverya≥0,theprocess e−tY(t,a),t≥0isauniformlyintegrablemartingale.Itslimit,sayγ(a),viewedasaprocessinthevariablea,hasthesameﬁnitedimensionallawsasastandardgammaprocess.
Proof:Fora=1,weseefromLemma3andtheidentityindistributionY(·,1)L=Y(1)(·)that e−tY(t,1),t≥0isauniformlyintegrablemartingaleandthatitslimithasthestandardexponentialdistribution.Theproofiseasilycompletedbyanappealtothebranchingproperty.Remark.ThelimitingdistributioninLemma6isessentiallyacorollaryofTheorem3ofGrey[10].Justasintheprecedingsection,wethinkofγ(a)asthesizeoftheterminalpopulationwhentheinitialpopulationhassizea.Wecanexpressγ(a)asγ(a)=Xb≤aδb,whereδ:=(δb,b≥0)isthejumpprocessofγ,whichcorrespondstodecomposingtheterminalpopulationintosub-populationshavingthesameancestorattheinitialtime.WewriteG(0,a)forthesequenceofthejumpsofγon[0,a],rankedindecreasingorder,andwededucefromLemma6thatconditionallyonγ(a)=g,G(0,a)/ghasdistributionPD(a).Moregenerally,bythebranchingproperty,wecandecomposetheterminalpopulationintosub-populationshavingthesameancestoratanygiventimet.Thisgivesγ(a)=Xb≤Y(t,a)e−tδ(t)b,whereδ(t):=(δ(t)b,b≥0)isthejumpprocessofastandardgammaprocessγ(t)whichisindependentoftheYuleprocessuptotimet,(Y(s,c),s∈[0,t]andc≥0).Thisenables12
ustodeﬁneforeacha>0thegenealogicalprocessassociatedtoaYuleprocessY(·,a),G(·,a)=(G(t,a),t≥0),whereetG(t,a)istherankedsequenceofthesizesofthejumpsofthesubordinatorγ(t)ontheinterval[0,Y(t,a)].AneasyvariationoftheargumentsfortheproofofTheorem4showsthatthege-nealogicalstructureoftheYuleprocesscanbedescribedintermsofthefragmentationchainX(∞)ofSection2.3asfollows.Theorem7.Fixa,g>0andletthechainX(∞)haveinitialdistributionPD(a).In-troduceastandardPoissonprocess,N=(Nt,t≥0),whichisindependentofthechainX(∞).ThenthecompoundchaingX(∞)(Ngt),t≥0hasthesamelawasthetime-changedprocess(G(log(1+t),a),t≥0)conditionedonγ(a)=g.Likewise,theanalogueofCorollary5isasfollows.Corollary8.Fixa>0.Then1
γ(a)G log(1+e−t/γ(a)),a,t∈Risatime-homogeneousMarkovcoalescentprocesswhichisindependentofγ(a).Supposethatitisinstatex∈∆∞andrecallRemark(a)ofSection2.3.Theniflim→0+1
log1/max{i:xi>}=n+a,theprocesswaitsanexponentialtimeofparameternandthenjumpstoavariabledis-tributedasCoag1/(n+a)(x),independentlyoftheexponentialtime.References[1]Athreya,K.B.andNey,P.E.(1972).Branchingprocesses.Springer-Verlag,Berlin-Heidelberg-NewYork[2]Bertoin,J.(2002).Self-similarfragmentations.Ann.Inst.HenriPoincar´e38,319-340[3]Bertoin,J.(2003).RandomcoveringofanintervalandavariationofKingman’scoalescent.ToappearinRandomStructuresAlgorithms.AlsoavailableasPreprintPMA-794athttp://www.proba.jussieu.fr/mathdoc/preprints/index.html[4]Bertoin,J.andLeGall,J.-F.(2000).TheBolthausen-Sznitmancoalescentandthegenealogyofcontinuous-statebranchingprocesses.Probab.TheoryRelat.Fields117,249-26613
[5]Bertoin,J.andLeGall,J.-F.(2003).Stochasticﬂowsassociatedtocoalescentpro-cesses.Probab.TheoryRelat.Fields126,261-288[6]Bertoin,J.andPitman,J.(2000).Twocoalescentsderivedfromtherangesofstablesubordinators.Elect.J.Probab.5,1-17.Availableviahttp://www.math.u-psud.fr/~ejpecp/ejp5contents.html[7]Bolthausen,E.andSznitman,A.S.(1998).OnRuelle’sprobabilitycascadesandanabstractcavitymethod.Comm.Math.Physics197,247-276[8]Duquesne,T.andLeGall,J.-F.(2002).Randomtrees,L´evyprocessesandspatialbranchingprocesses.Ast´erisque281[9]Fleischmann,K.andSiegmund-Schultze,R.(1977)ThestructureofreducedcriticalGalton-Watsonprocesses.Math.Nachr.79,233-241[10]Grey,D.R.(1974).Asymptoticbehaviourofcontinuoustime,continuousstate-spacebranchingprocesses.J.Appl.Probab.11,669-677[11]Kallenberg,O.(1973).Canonicalrepresentationsandconvergencecriteriaforpro-cesseswithinterchangeableincrements.Z.Wahrsch.verw.Gebiete27,23-36[12]Kendall,D.G.(1966).Branchingprocessessince1873.J.LondonMath.Soc.41,385-406[13]Kingman,J.F.C.(1982).Thecoalescent.StochasticProcess.Appl.13,235-248[14]Lamperti,J.(1967).Thelimitofasequenceofbranchingprocesses.Z.Wahrsch.verw.Gebiete7,271-288[15]Lamperti,J.(1967).Continuous-statebranchingprocesses.Bull.Amer.Math.Soc.73,382-386[16]LeGall,J.-F.(1989).Marchesal´eatoires,mouvementbrownienetprocessusdebranchement.S´eminairedeProbabilit´esXXIII,LectureNotesinMath.,1372,Springer,Berlin,258-274[17]Pitman,J.(1999).Coalescentswithmultiplecollisions.Ann.Probab.27,1870-1902[18]Pitman,J.(2002).CombinatorialStochasticProcesses.LecturenotesfortheStFloursummerschool.Toappear.Availableat http://stat-www.berkeley.edu/users/pitman/621.ps.Z[19]Pitman,J.andYor,M.(1997).Thetwo-parameterPoisson-Dirichletdistributionderivedfromastablesubordinator.Ann.Probab.25,855-900[20]Revuz,D.andYor,M.(1999).ContinuousmartingalesandBrownianmotion,Thirdedition.Springer-Verlag,Berlin14

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

---
source: "https://arxiv.org/abs/1912.11949v1"
title: "Emergence of stochastic flocking for the discrete Cucker-Smale model with randomly switching topologies"
author: "Jiu-Gang Dong, Seung-Yeal Ha, Jinwook Jung, Doheon Kim"
year: "2019"
publication: "arXiv preprint / math.DS"
download: "https://arxiv.org/pdf/1912.11949v1"
pdf: "https://arxiv.org/pdf/1912.11949v1"
captured_at: "2026-05-09T13:02:31Z"
updated_at: "2026-05-09T13:02:31Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代思想"
status: raw
---

# Emergence of stochastic flocking for the discrete Cucker-Smale model with randomly switching topologies

- 著者: Jiu-Gang Dong, Seung-Yeal Ha, Jinwook Jung, Doheon Kim
- 年: 2019
- 掲載情報: arXiv preprint / math.DS
- 情報源: [arxiv](https://arxiv.org/abs/1912.11949v1)
- ダウンロード: https://arxiv.org/pdf/1912.11949v1
- PDF: https://arxiv.org/pdf/1912.11949v1

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

We study emergent dynamics of the discrete Cucker-Smale (in short, DCS) model with randomly switching network topologies. For this, we provide a sufficient framework leading to the stochastic flocking with probability one. Our sufficient framework is formulated in terms of an admissible set of network topologies realized by digraphs and probability density function for random switching times. As examples for the law of switching times, we use the Poisson process and the geometric process and show that these two processes satisfy the required conditions in a given framework so that we have a stochastic flocking with probability one. As a corollary of our flocking analysis, we improve the earlier result [J.-G. Dong, S.-Y. Ha, J. Jung and D. Kim: On the stochastic flocking of the Cucker-Smale flock with randomly switching topologies. arXiv:1911.07390.] on the continuous C-S model.

## PDF Text

arXiv:1912.11949v1 [math.DS] 27 Dec 2019
EMERGENCEOFSTOCHASTICFLOCKINGFORTHEDISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIESJIU-GANGDONG,SEUNG-YEALHA,JINWOOKJUNG,ANDDOHEONKIMAbstract.WestudyemergentdynamicsofthediscreteCucker-Smale(inshort,DCS)modelwithrandomlyswitchingnetworktopologies.Forthis,weprovideasuﬃcientframe-workleadingtothestochasticﬂockingwithprobabilityone.Oursuﬃcientframeworkisformulatedintermsofanadmissiblesetofnetworktopologiesrealizedbydigraphsandprobabilitydensityfunctionforrandomswitchingtimes.Asexamplesforthelawofswitch-ingtimes,weusethePoissonprocessandthegeometricprocessandshowthatthesetwoprocessessatisfytherequiredconditionsinagivenframeworksothatwehaveastochasticﬂockingwithprobabilityone.Asacorollaryofourﬂockinganalysis,weimprovetheearlierresult[15]onthecontinuousC-Smodel.1.IntroductionThejargon“ﬂocking”denotesacollectivephenomenoninwhichself-propelledparticlesareorganizedinanorderedmotionwiththesamevelocityusingonlyenvironmentalinfor-mation,anditissometimesusedinterchangeablyinliteraturewithotherjargonssuchasswarmingandherding.Flockingphenomenaareoftenobservedinmanybiologicalcomplexsystems[3,7,37,43,44,45]andfoundinapplicationsintheengineeringdomainsuchassensornetworks[31].In2007,CuckerandSmale[14]introducedasimplesecond-orderNetwon-likemodelinthespiritofVicsek’smodel[44]asananalyticallymanageablemodel,andtheyderivedsuﬃcientframeworksleadingtotheglobalandlocalasymptoticﬂockingonthecompletegraph.Aftertheirwork,theCucker-Smale(CS)modelhasbeenextensivelystudiedinappliedmathcommunityfromseveralaspects,e.g.,stochasticnoises[1,4,13,27],discreteapproximation[17,23],randomfailures[19,20,30,41],time-delayedcommunica-tions[8,9,21,34],collisionavoidance[10],networktopologies[11,12,18,32,33,39,42],mean-ﬁeldlimit[26,28],kineticandhydrodynamicdescriptions[5,6,22,29,40],nonsym-metriccommunicationnetworks[36],bondingforce[38],extrainternalvariables[24,25],etc(seerecentsurveyarticles[2,7]).Inthispaper,amongothers,wefocusontheeﬀectoftherandomlyswitchingnetworktopologiesinasymptoticﬂocking.Toﬁxtheidea,consideragroupofmigratingbirds.Then,oneeasilyobservesthatleaderscanchangeovertime,birds’ﬂockcansplitintoseverallocalgroups,andthensomeofthemcanmergetoformabiggergroupfromtimetotime.Thus,
Date:December30,2019.2010MathematicsSubjectClassiﬁcation.39A12,34D05,68M10.Keywordsandphrases.Cucker-Smalemodel,randomlyswitchingtopology,directedgraphs,ﬂocking.Acknowledgment.TheworkofS.-Y.HaissupportedbytheNationalResearchFoundationofKorea(NRF-2017R1A2B2001864).TheworkofJ.JungissupportedbytheNationalResearchFounda-tionofKorea(NRF)grantfundedbytheKoreagovernment(MSIP):NRF-2016K2A9A2A13003815.1
2DONG,HA,JUNG,ANDKIMtheunderlyingnetworktopologieskeepchangingovertime.Hence,itisnaturaltoconsiderﬂockingmodelswithrandomlyswitchingnetworktopologiesforrealisticﬂockingmodeling.Inthisdirection,authorsrecentlystudiedtheemergentdynamicsofthecontinuousCSmodelwithrandomlyswitchingnetworktopologiesandprovidedasuﬃcientframeworkforasymptoticglobalﬂockingin[15](seerelatedworks[35,46]forlinearconsensusmodelswithrandomlyswitchingtopologies).ThegoalofthisworkistorevisitthediscreteCucker-Smalemodelwithrandomlyswitch-ingnetworktopologies,andprovideasuﬃcientframeworkforglobalﬂockingandimprovetheearlierresult[15]onthecontinuousmodel.Toaddressourproblem,weﬁrstdiscussourproblemsetting.LetSG:={G1,···,GNG}betheadmissiblesetofnetworktopologies(digraphs)satisfyingacertainconnectivityconditiontobespeciﬁedlater,andletxiandvibethepositionandvelocityofthei-thC-Sparticleundertheinﬂuenceofrandomlyswitchingnetworktopologies.Then,theirdynamicsisgovernedbytheCSmodel:(1.1)





˙xi=vi,t>0,1≤i≤N,˙vi=1
NNXj=1χσ
ijφ(kxj−xik)(vj−vi),wherek·kdenotesthestandard`2-norminRdandthecommunicationweightφ:[0,∞)→R+isbounded,Lipschitzcontinuousandmonotonicallydecreasing:0<φ(r)≤φ(0)=:κ,[φ]Lip:=supx6=y|φ(x)−φ(y)|
|x−y|<∞,(φ(r)−φ(s))(r−s)≤0,r,s≥0.(1.2)ΩTheadjacentmatrix(χσ
ij)denotesthetime-dependentnetworktopologycorrespondingtotherandomlyswitchinglawσ:[0,∞)×Ω→{1,···,NG}.Thelawσisassociatedwithasequence{t`}`≥0ofrandomvariables,calledrandomswitchinginstants,suchthatforeachﬁxedω∈Ω,σ(ω):[0,∞)→{1,···,NG}ispiecewiseconstantandright-continuouswhosediscontinuitiesare{t`(ω)}`≥0.Inaddition,therandomlawσassigns(1.1)withanetworktopologyateachinstant.Speciﬁcally,onceweﬁxω∈Ωandaninstantt>0,thenσ(t,ω)=σ(t`(ω))=kforsome1≤k≤NGand`≥0,andweequipsystem(1.1)withthenetworktopology(χσ(t,ω)ij)givenbythe0-1adjacencymatrixofthek-thdigraphGk∈SG.Next,weconsiderthediscreteanalogofCSwhoseemergentdynamicswewillstudy.ThediscretecounterpartofCSwillbegivenbytheforwardﬁrst-orderEulermethod.Leth:=∆tbethetime-step,andfornotationalsimplicity,weusethefollowinghandynotation:xi[t]:=xi(th),vi[t]:=vi(th)andσ[t]:=σ(th),t∈{0}∪Z+.Then,thediscreteC-Smodelreadsasfollows:Ω(1.3)





xi[t+1]=xi[t]+hvi[t],t=0,1,···,1≤i≤N,vi[t+1]=vi[t]+h
NNXj=1χσ[t]ijφ(kxj[t]−xi[t]k)(vj[t]−vi[t]).Inthispaper,welookforasuﬃcientframeworkleadingtotheasymptoticglobalﬂocking:sup0≤t<∞max1≤i,j≤Nkxi[t]−xj[t]k<∞,limt→∞max1≤i,j≤Nkvi[t]−vj[t]k=0.
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES3Beforewepresentourmainresult,weintroduceconﬁgurationdiameters:X:=(x1,···,xN),V:=(v1,···,vN),D(X):=max1≤i,j≤Nkxi−xjk,D(V):=max1≤i,j≤Nkvi−vjk.Below,webrieﬂydiscussourmainresultfor(1.3).Asaforementioned,weneedtoimposecertainconditionsontheadmissiblesetSGfornetworktopologiesandtheprobabilitydis-tributionoftherandomswitchingtimes{t`}.Asdiscussedin[15],therandomdwellingtimesbetweenimmediateswitchingsneedtobemadenottooshortandnottoolong.IfthishappensandchoiceofnetworktopologiesintheadmissiblesetSGismaderandomly,thentheunionofnetworktopologiesinSGwillgoverntheasymptoticdynamicsinasuf-ﬁcientlylongﬁnitetimeinterval,duetothelawoflargenumbers.Hence,weneedtoassumethattheuniongraphcontainsaspanningtreesothatanypairoftwoparticleswillcommunicateeventuallyinaﬁnite-timeinterval.TheseheuristicargumentscanbemaderigorouslyinSection3.1.Thenundertheseassumptionsonthenetworktopologiesandrandomswitchingtimes,ifthetime-steph,choiceprobabilitypkforthenetworktopologyGkandcommunicationweightφsatisfyasetofassumptions:0<hκ<1,log1
1−hκ
min1≤k≤NGlog1
1−pk1,1
φ(r)=O(rε)asr→∞,whereεisasmallpositiveconstanttobedeterminedlater,thenwecanderiveasymptoticﬂockingwithprobabilityone(seeTheorem3.1):Pω:∃x∞>0s.tsup0≤t<∞D(X[t,ω])≤x∞,andlimt→∞D(V[t,ω])=0=1.Therestofthispaperisorganizedasfollows.InSection2,webrieﬂyreviewbasictermi-nologiesandlemmastobeusedthroughoutthispaper,suchasdirectedgraphs,scramblingmatrices,andtheirbasicproperties.Wealsoprovideamonotonicitylemmaforvelocitydi-ameter.InSection3,weprovideaframeworkandourmainresultsofthispaper.InSection4,weprovidetechnicalproofforourmainresult.InSection5,weconsidertwoconcretestochasticprocessesfordwellingtimes(Poissonprocessandgeometricprocess)andthenshowthattheysatisfytheproposedframeworkwhichresultsintheasymptoticﬂocking.Wealsodiscusssomeimprovementsforthecontinuousmodelbyremovingthecompactsupportassumptionontheprobabilitydensityfunctionofswitchingtimes.Finally,Section6isdevotedtoasummaryofourmainresultsandsomefutureissuestobediscussedlater.NotationThroughoutthepaper,(Ω,F,P)denotesagenericprobabilityspace.Forma-tricesA=(aij)N×NandB=(bij)N×N,wedenoteA≥Bifaij≥bijforalliandj.Moreover,wedeﬁnetheFrobeniusnormofAasfollows.kAkF:=p tr(A∗A),andId=diag(1,···,1)isthed×didentitymatrix.2.PreliminariesInthissection,webrieﬂyreviewbasicmaterialsonthedirectedgraphsandmatrixalgebra,andthenstudythedissipativestructureofsystem(1.3).
4DONG,HA,JUNG,ANDKIM2.1.Directedgraphsandmatrixalgebra.Forthemodelingofnetworktopology,weusedirectedgraphs(digraphs).AdigraphG=(V,E)consistsofavertexsetV={1,···,N}andanedgesetE⊂V×Vconsistingoforderedconnectedpairsofvertices.Wesayjisaneighborofiif(j,i)∈EandNi:={j:(j,i)∈E}denotestheneighborsetofi.If(i,i)∈E,thenGissaidtohaveaself-loopati.ForagivendigraphG=(V,E),weconsiderits0-1adjacencymatrixχ=(χij),whereχij:=1if(j,i)∈E,0if(j,i)/∈E.ApathinGfromitojisasequenceofdistinctverticesi0=i,i1,···,in=jsuchthat(im−1,im)∈Eforevery1≤m≤n.Ifthereisapathfromitoj,thenwesayjisreachablefromi.Moreover,adigraphGissaidtohaveaspanningtree(orrooted)ifGhasavertexifromwhichanyotherverticesarereachable.Next,wereviewsomeconceptsandresultsinmatrixalgebrawhichwillbecruciallyusedinlatersections.First,weintroduceseveralconceptsofnonnegativematricesinthefollowingdeﬁnition.
Deﬁnition2.1.LetA=(aij)beanonnegativeN×Nmatrix.(1)Aisastochasticmatrix,ifitsrow-sumisequaltounity:NXj=1aij=1,1≤i≤N.(2)Aisascramblingmatrix,ifforeachpairofindicesiandj,thereexistsanindexksuchthataik>0andajk>0.(3)AisaweightedadjacencymatrixofadigraphGifthefollowingholds:aij>0⇐⇒(j,i)∈E.Inthiscase,wewriteG=G(A).(4)TheergodicitycoeﬃcientofAisdeﬁnedasfollows.µ(A):=mini,jNXk=1min{aik,ajk}.Remark2.1.Thefollowingassertionseasilyfollowfromtheabovedeﬁnition:(1)Aisscramblingifandonlyifµ(A)>0.(2)FornonnegativematricesAandB,A≥B=⇒µ(A)≥µ(B).Inthesequel,wepresenttwolemmaswithoutproofs.Theselemmaswillbecruciallyusedinnextsections.
Lemma2.1.(Lemma2.2,[16])SupposethatanonnegativeN×NmatrixA=(aij)isstochastic,andletB=(bj i),Z=(zji)andW=(wji)beN×dmatricessuchthatW=AZ+B.Then,wehavemaxi,kkwi−wkk≤(1−µ(A))maxl,mkzl−zmk+√
2kBkF,
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES5wherezi:=(z1i,···,zdi),bi:=(b1
i,···,bd i),wi:=(w1i,···,wdi),i=1,···,N.Lemma2.2.(Theorem5.1,[46])IfAiarenonnegativeN×Nmatriceswithpositivediago-nalelementssuchthatG(Ai)hasaspanningtreeforall1≤i≤N−1,thenA1A2...AN−1isscrambling.Ω2.2.MatrixformulationoftheDCSmodel.Inthissubsection,weprovideaprioriestimateforthevelocitydiameteralongthesamplepath.Forconvenience,wesuppressω-dependenceinthesequel.First,weconsidertheR.H.S.of(1.3)2:RHS2:=vi[t]+h
NNXj=1χσ[t]ijφ(kxj[t]−xi[t]k)(vj[t]−vi[t])=1−h
NNXj=1χσ[t]ijφ(kxj[t]−xi[t]k)+h
Nχσ[t]iiφ(kxj[t]−xi[t]k)vi[t]+h
NXj6=iχσ[t]ijφ(kxj[t]−xi[t]k)vj[t].(2.1)ΩTorewrite(2.1)asamatrixform,weintroduceN×NmatrixLk[t](k=1,···,NG):Lk[t]:=Dk[t]−Ak[t],wherethematricesAk[t]andDk(t)aredeﬁnedasfollows.Ak[t]=ak ij[t],Dk(t)=diagdk
1[t],···,dk
N[t],ak ij[t]:=χk ijφ(kxi[t]−xj[t]k)anddk i[t]=NXj=1χk ijφ(kxi[t]−xj[t]k).(2.2)ΩFinally,wecombine(2.1)and(2.2)torewritesystem(1.3)asamatrixform:

X[t+1]=X[t]+hV[t],t=0,1,···,V[t+1]=Id−h
NLσ[t][t]V[t],(2.3)ΩanddeﬁneΩ(2.4)Φ[s1,s0,ω]:=s1−1Yt=s0Id−h
NLσ[t,ω].Next,westudythemonotonicityofD(V[t]).Lemma2.3.Supposethatthetime-stepandcouplingstrengthsatisfy0<hκ<1,andlet(X[t],V[t])beasolutionto(2.3).Thenforeachﬁxedω∈Ω,wehaveD(V[t+1])≤D(V[t]),D(X[t+1])≤D(X[t])+hD(V[t]),foranyt∈N∪{0}.
6DONG,HA,JUNG,ANDKIMProof.First,weuseφ(r)≤φ(0)=κin(1.2)andhκ<1toseethatforeachi=1,···,Nandk=1,···,NG,h
Ndk i[t]=h
NNXj=1χk ijφ(kxi[t]−xj[t]k)≤hκ<1.ThisimpliesthatId−h
NLσ[t][t]isalwaysnonnegativeanditslowsumisone,i.e.,stochastic.Sincetheergodicitycoeﬃcientofanynonnegativematricesisnonnegative,weapplyLemma2.1togetD(V[t+1])≤1−µId−h
NLσ[t][t]D(V[t])≤D(V[t]).Theotherestimateforthespatialdiameterreadilyfollowsfrom(2.3)1.3.AframeworkandmainresultInthissection,webrieﬂypresentaframeworkforstochasticﬂockingintermsofadmis-siblenetworktopologiesandswitchingtimes,andstateourmainresult.3.1.Problemsetup.Let(Ω,F,P)beaprobabilityspace.Then,wesetthesequenceof(possibly)switchingtimes{t`}`≥0tobeasequenceofN-valuedindependentrandomvariableson(Ω,F,P).Then,weassumethattheswitchinglawσ:N×Ω→{1,···,NG}andthesequence{t`+1−t`}`≥0ofdwellingtimessatisfythefollowingconditions:•Foreach`≥0,t`+1−t`isgivenby(3.1)t`+1−t`=1+T`,where{T`}`≥0isagivensequenceofindependentinteger-valuednonnegativerandomvariables.•Foreach`≥0andω∈Ω,σt(ω):=σ(t,ω)isconstantontheintervalt∈[t`(ω),t`+1(ω)).•{σt`}`≥0isasequenceofi.i.d.randomvariablessuchthatforany`≥0,P(σt`=k)=pk,foreachk=1,···,NG,wherep1,···,pNG>0aregivenpositiveconstantssatisfyingp1+···+pNG=1.Webrieﬂycommentontherandomswitchingtimeprocess{t`}asfollows.Supposethatwehaveprocessatt=t`.Then,thenextswitchingtimet`+1isdeterminedbytherelation(3.1)viatheinteger-valuedandnonnegativerandomprocessT`:t`+1=t`+1+T`,InSection5,wewillconsidertwoexplicitprocesses(Poissonandgeometricprocesses)andshowthattheysatisfyourframeworkthatwillbepresentedsoon.Foreachk=1,···,NG,letGk=(V,Ek)bethek-thadmissibledigraph.Then,foreacht≥0andω∈Ω,thetime-dependentnetworktopology(χσ
ij)=(χσ[t,ω]ij)isdeterminedby
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES7χσ[t,ω]ij:=1if(j,i)∈Eσ[t,ω],0if(j,i)/∈Eσ[t,ω].Fortechnicalreasonsandwithoutlossofgenerality,weassumethateachadmissibledigraphGkhasaself-loopateachvertex,andwedeﬁnetheuniongraphofGσ[t,ω]fors0≤t<s1andω∈Ω:G([s0,s1))(ω):=s1−1[t=s0Gσ[t,ω]= V,s1−1[t=s0Eσ[t,ω]!.Notethatthenetworktopologymightremainunchangedatswitchinginstants.Inotherwords,itwouldbepossiblethatσt`+1(ω)=σt`(ω)forsome`≥0andω∈Ω.3.2.Aframeworkandmainresult.Inthissubsection,wepresentasuﬃcientframe-workintermsofthestructureofadmissiblenetworktopologiesandlawsofswitchingtimesandourmainresult.Foreachpositiveintegernandpositiverealnumberc>0,wedeﬁneanincreasingsequence{a`(n,c)}`∈Nofintegersbythefollowingrecursiverelation:(3.2)a0(n,c)=0,a`+1(n,c)=a`(n,c)+n+bclog(`+1)c,(`∈N).Inwhatfollows,wesometimessuppressnandcdependencefora`(n,c):a`:=a`(n,c).Weimposethefollowingassumptions(A)onthesetSGofadmissibledigraphsandtheprobabilitydensityfunctionoftheswitchingtimes:•(A1):TheuniondigraphfromSGhasaspanningtree:[1≤k≤NGGk=
V,[1≤k≤NGEk
hasaspanningtree.•(A2):Thedwellingtimesareboundedinprobability:forsomeM∈N,P
ω:ai(N−1)−1X`=a(i−1)(N−1)T`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
≤˜p(n),where˜p(n)→0asn→∞.Now,wearereadytostateourmaintheoremontheemergenceofﬂocking.Theorem3.1.SupposethatparametersN,h,κ,pk’sandcommunicationweightφsatisfythefollowingconditions:0<hκ<1,(M+N−1)log1
1−hκ
min1≤k≤NGlog1
1−pk<1,1
φ(r)=O(rε)asr→∞,whereMisapositiveconstantin(A2)andεisapositiveconstantsatisfying0≤ε<1
N−1−(M+N−1)log1
1−hκ
(N−1)min1≤k≤NGlog1
1−pk,
8DONG,HA,JUNG,ANDKIMandlet(X,V)beasolutionprocessto(1.3).Then,asymptoticglobalﬂockingisexhibitedwithprobabilityone:Pω:∃x∞>0s.tsup0≤t<∞D(X[t,ω])≤x∞,andlimt→∞D(V[t,ω])=0=1.Proof.Theproofwillbegiveninnextsection.4.EmergenceofstochasticflockingInthissection,weprovideasuﬃcientframeworkfortheemergenceofﬂockinginsystem(1.3),andprovideaproofforTheorem3.1.First,wepresentaprioriassumptions(˜A1)-(˜A3):foraﬁxedω∈Ω,•(˜A1):thereexistn∈N,n>0andc>0suchthatcNlog1
1−hκ<1,andthesubsequence{t∗
`}`∈N⊂{t`}`∈Ndeﬁnedbyt∗
`:=ta`(n,c)in(3.2)satisﬁesG([t∗
`,t∗
`+1))(ω)hasaspanningtreeforall`≥0.•(˜A2):thereexistn∈N,M∈N,andc>0suchthattherandomsequence{T`(ω)}`≥0satisfythefollowingboundednesscondition:(4.1)ai(N−1)−1X`=a(i−1)(N−1)T`(ω)<M(n+bclogi(N−1)c),foreachi∈N.•(˜A3):thepositiondiameterisuniformlyboundedintime:sup0≤t<∞D(X[t,ω])≤x∞<∞.FortheproofofTheorem3.1,wesplititsproofintothreestepsasfollows:•StepA:Aprioriassumptions(˜A1),(˜A2)and(˜A3)implystochasticﬂocking.•StepB:Conditionsoninitialconﬁgurationimpliestheaprioriassumption(˜A3).•StepC:Framework(A1)-(A2)implyconditiononinitialconﬁgurationandaprioriassumptions(˜A1)-(˜A2).Inthefollowingthreesubsections,wepresenteachsteponebyone.4.1.StepA(Aprioriassumptionsimplystochasticﬂocking).Inthissubsection,weshowthattheaprioriassumptions(˜A1),(˜A2)and(˜A3)implystochasticﬂocking.Lemma4.1.Supposethataprioriassumptions(˜A1)and(˜A3)hold,andsystemparameterssatisfy0<hκ<1,andlet(X,V)beasolutionto(1.3).Then,thematrixΦ[t∗
r(N−1),t∗
(r−1)(N−1)]deﬁnedin(2.4)satisﬁes
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES9(1)Φ[t∗
r(N−1),t∗
(r−1)(N−1)]isstochasticforeachr∈N.(2)TheergodicitycoeﬃcientofΦ[t∗
r(N−1),t∗
(r−1)(N−1)]satisﬁesµ(Φ[t∗
r(N−1),t∗
(r−1)(N−1)])≥(1−hκ)t∗
r(N−1)−t∗
(r−1)(N−1)hφ(x∞)
N(1−hκ)N−1,foreachr∈N.Proof.Foreachr∈Nand1≤i≤N−1,weconsiderΦ[t∗
(r−1)(N−1)+i,t∗
(r−1)(N−1)+i−1].Here,weassumethat{t`}`≥0∩[t∗
(r−1)(N−1)+i−1,t∗
(r−1)(N−1)+i]={t`1,···,t`q+1}.Then,wegetΦ[t∗
(r−1)(N−1)+i,t∗
(r−1)(N−1)+i−1]=qYj=1Φ[t`j+1,t`j],andforeachj,wecanchoosekj∈{1,···,NG}suchthatσ[t,ω]=kjforallt∈N∩[t`j,t`j+1).Thus,wehaveΦ[t`j+1,t`j]=t`j+1−1Yt=t`jId−h
NLkj[t]=t`j+1−1Yt=t`jId−h
NDkj[t]+h
NAkj[t]≥t`j+1−1Yt=t`j(1−hκ)Id+h
NAkj[t],andconsequently,Φ[t∗
(r−1)(N−1)+i,t∗
(r−1)(N−1)+i−1]=qYj=1Φ[t`j+1,t`j]≥qYj=1t`j+1−1Yt=t`j(1−hκ)Id+h
NAkj[t]≥(1−hκ)t∗
(r−1)(N−1)+i−t∗
(r−1)(N−1)+i−1−1h
NqXj=1t`j+1−1Xt=t`jAkj[t]≥(1−hκ)t∗
(r−1)(N−1)+i−t∗
(r−1)(N−1)+i−1hφ(x∞)
N(1−hκ)Fi,whereFidenotesthe0-1adjacencymatrixofthegraphG[t∗
(r−1)(N−1)+i−1,t∗
(r−1)(N−1)+i)andthesecondinequalityfollowsfromthefact:mYi=1(aId+Bi)=amId+am−1mXi=1Bi+···≥am−1mXi=1Bi,foralla>0,Bi≥0.Hence,wehaveΦ[t∗
r(N−1),t∗
(r−1)(N−1)]=N−1Yi=1Φ[t∗
(r−1)(N−1)+i,t∗
(r−1)(N−1)+i−1]
10DONG,HA,JUNG,ANDKIM≥(1−hκ)t∗
r(N−1)−t∗
(r−1)(N−1)hφ(x∞)
N(1−hκ)N−1N−1Yi=1Fi.SinceeachFiisthe0-1adjacencymatrixofagraphwithaspanningtree,QN−1i=1FiisscramblingbyLemma2.2.Moreover,wehaveµ N−1Yi=1Fi!≥1.Thus,wecandeducetheresult(2).Forthestochasticity,notethatId−h
NLk[t]isstochasticforanyk=1,···,NG.ThisgivesthestochasticityofΦ[t`j+1,t`j]andconsequently,weobtainthedesiredresult.Next,weprovidethedecayestimatesforthevelocitydiameterbasedonaprioriassump-tions.Inthesequel,wesetX0:=X[0],V0:=V[0].Proposition4.1.(Velocityalignment)Supposethat(˜A1)-(˜A3)hold,andsystemparame-terssatisfy0<hκ<1,andlet(X,V)beasolutionto(1.3).Thenforeacht∈N∩[t∗
r(N−1),t∗
(r+1)(N−1)),D(V[t])≤D(V0)×exp"−(1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)#.Proof.Foreachr≥0,weuseLemma4.1togetD(V[t])≤D(V[t∗
r(N−1)])≤1−µΦ[t∗
r(N−1),t∗
(r−1)(N−1)]D(V[t∗
(r−1)(N−1)])≤rYi=11−µΦ[t∗
i(N−1),t∗
(i−1)(N−1)]D(V0)≤rYi=1 1−(1−hκ)t∗
r(N−1)−t∗
(r−1)(N−1)hφ(x∞)
N(1−hκ)N−1!D(V0).Notethatt∗
i(N−1)−t∗
(i−1)(N−1)=ai(N−1)−1X`=a(i−1)(N−1)(t`+1−t`)=ai(N−1)−a(i−1)(N−1)+ai(N−1)−1X`=a(i−1)(N−1)T`≤(M+N−1)(n+bclogi(N−1)c).Thisfollowsfromthedeﬁnitionofthesequencea`andtheaprioriassumption.Thus,wehave
D(V[t])≤rYi=1"1−(1−hκ)(M+N−1)(n+bclogi(N−1)c)hφ(x∞)
N(1−hκ)N−1#D(V0)
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES11≤rYi=1exp"−(1−hκ)(M+N−1)(n+clogi(N−1))hφ(x∞)
N(1−hκ)N−1#D(V0)=exp"−(1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1rXi=1ic(M+N−1)log(1−hκ)#D(V0)≤exp"−(1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1Zr+11xc(M+N−1)log(1−hκ)dx#D(V0)=exp"−(1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)#D(V0).Thisleadstothedesiredestimate.4.2.StepB(Pathwisely,(˜A1)and(˜A2)implyﬂocking).Inthissubsection,weshowthatthespatialdiameterisuniformlybounded.Forthis,wepresentasimplelemmawithoutaproof.
Lemma4.2.[15]Foranyx>0andδ>0,wehavethefollowinginequality:e−x≤δ
eδx−δ.Now,weshowthattheaprioriassumption(˜A3)onthepositiondiametercanbereplacedbytheassumptiononinitialdata.
Lemma4.3.(Pathwiseuniformboundforspatialdiameter)Supposethatω∈Ωsatisﬁesaprioriassumptions(˜A1)-(˜A2),andsystemparametersandinitialdatasatisfy(i)0<hκ<1.(ii)∃δ>0andx∞>0independentofasamplepointsuchthatD(X0)+D(V0)h(M+N−1)(n+clog(N−1))+hD(V0)∞Xr=1"(M+N−1)(n+clog((r+1)(N−1)))δ
eδ× (1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)!−δ#<x∞,andlet(X,V)beasolutionprocessto(1.3).Then,aprioricondition(˜A3)isfulﬁlled:sup0≤t<∞D(X[t,ω])<x∞.Proof.WeuseLemma4.1,Lemma4.2andthecondition(ii)toseethatforeacht>0,D(X[t])≤D(X0)+ht−1Xs=0D(V[s])≤D(X0)+hD(V0)∞Xr=0"(t∗
(r+1)(N−1)−t∗
r(N−1))
12DONG,HA,JUNG,ANDKIM×exp −(1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)!#≤D(X0)+hD(V0)∞Xr=0"(M+N−1)(n+clog((r+1)(N−1)))×exp −(1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)!#≤D(X0)+hD(V0)(M+N−1)(n+clog(N−1))+hD(V0)∞Xr=1"(M+N−1)(n+clog((r+1)(N−1)))δ
eδ× (1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)!−δ#<x∞.Thisimpliesthedesiredresult.Finally,wecombineProposition4.1with(˜A1)-(˜A2)andaconditiononcommunicationweighttoshowtheemergenceofﬂockingforanyinitialconﬁguration.Proposition4.2.Supposethatω∈Ωsatisﬁesaprioriassumptions(˜A1)-(˜A2),andsystemparametersandcommunicationweightsatisfy0<hκ<1and1
φ(r)=O(rε)asr→∞,whereεisapositiveconstantsatisfyingtherelation:0≤ε<1+c(M+N−1)log(1−hκ)
N−1,andlet(X,V)beasolutionprocessto(1.3).Then,theglobalﬂockingemergesforanyinitialconﬁgurationpathwise:thereexistsx∞>0suchthatsup0≤t<∞D(X(t,ω))≤x∞andlimt→∞D(V(t,ω))=0.Proof.First,wechooseδ>0suchthat(4.2)1
1+c(M+N−1)log(1−hκ)<δ<1
(N−1)ε.Then,theL.H.S.in(4.2)yieldstheconvergenceoftheseriesappearingin(ii)ofLemma4.3:∞Xr=1"(n+clog((r+1)(N−1))) (r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)!−δ#≤C∞Xr=11+log(r+1)
 (r+1)1+c(M+N−1)log(1−hκ)−1δ
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES13=C r∗−1Xr=1+∞Xr=r∗!1+log(r+1)
 (r+1)1+c(M+N−1)log(1−hκ)−1δ≤C"r∗−1Xr=11+log(r+1)
 (r+1)1+c(M+N−1)log(1−hκ)−1δ+∞Xr=r∗1+log(r+1)
r(1+c(M+N−1)log(1−hκ))δ#<∞,whereC>0isapositiveconstantandr∗>0isanaturalnumbersatisfying(r+1)1+c(M+N−1)log(1−hκ)−1≥1
2r1+c(M+N−1)log(1−hκ),a∈(0,1),∀r≥r∗.Notethattheexistenceofr∗isguaranteedsincelimr→∞(r+1)1+c(M+N−1)log(1−hκ)−1
r1+c(M+N−1)log(1−hκ)=1.Furthermore,sincetheseriesin(ii)ofLemma4.3converges,wemaydeducethatD(X0)+hD(V0)(M+N−1)(n+clog(N−1))+hD(V0)∞Xr=1"(M+N−1)(n+clog((r+1)(N−1)))δ
eδ× (1−hκ)(M+N−1)(n+clog(N−1))hφ(x∞)
N(1−hκ)N−1(r+1)1+c(M+N−1)log(1−hκ)−1
1+c(M+N−1)log(1−hκ)!−δ#≤C1+(φ(x∞))−(N−1)δ,whereCisaconstantindependentofx∞.NotethattheR.H.S.in(4.2)givesφ(r)−(N−1)δ=O(r(N−1)δε)asr→∞.ThisimpliesΩ(4.3)limr→∞φ(r)−(N−1)δ
r=0.Thus,itfollowsfrom(4.3)thatwecandeducetheexistenceofx∞satisfyingtherelation(ii)Lemma4.3forδgivenin(4.2).Sincethecondition(ii)isnowsatisﬁed,weobtainourdesiredresultsfromLemma4.2andProposition4.1.4.3.StepC(Framework(A1)−(A2)impliesstochasticﬂocking).Inthissubsection,weshowthataprioriassumptions(˜A1)−(˜A2)canbefulﬁlledunderourframeworkandprovethatﬂockingemergeswithprobabilityone.First,weshowthattheﬁrstaprioriassumption(˜A1)canbesatisﬁedalmostsurelywiththeappropriatechoicesofparameters.Below,wedeﬁneasubsequence{t∗
`}`≥0⊂{t`}`≥0byt∗
`:=ta`(n,c).
14DONG,HA,JUNG,ANDKIMLemma4.4.Supposethatapositiveintegernandarealnumberc>0aresuﬃcientlylargeenoughsuchthatNGXk=1(1−pk)n≤1
2andc>−1
log(1−pk),k=1,···,NG,andlet(X,V)beasolutionprocessto(1.1).Thenwehavethefollowingassertions.(1)Foreachk=1,···,NG,∞X`=0(1−pk)bclog(`+1)c<∞.(2)P(ω:G([t∗
`,t∗
`+1))(ω)hasaspanningtreeforany`≥0)≥exp −(2log2)NGXk=1(1−pk)n∞X`=0(1−pk)bclog(`+1)c!=:p1(n).Proof.TheproofisalmostthesameasthatofProposition4.3in[15].Thus,weomittheproof.ProofofTheorem3.1:Wechooseεsothat0≤ε<1
N−1−(M+N−1)log1
1−hκ
(N−1)min1≤k≤NGlog1
1−pk,orequivalently1
min1≤k≤NGlog1
1−pk<1−ε(N−1)
(M+N−1)log1
1−hκ,andwesetc:=1
2
1
min1≤k≤NGlog1
1−pk+1−ε(N−1)
(M+N−1)log1
1−hκ
.First,notethatc>1
min1≤k≤NGlog1
1−pk,c(M+N−1)log1
1−hκ<1,and0≤ε<1−c(M+N−1)log1
1−hκ
N−1.Now,wechooseanyn∈NsuchthatNGXk=1(1−pk)n≤1
2.DeﬁnetheeventsA,BandFasA:G([t∗
`,t∗
`+1))(ω)hasaspanningtreeforany`≥0,
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES15B:ai(N−1)−1X`=a(i−1)(N−1)T`(ω)<M(n+bclogi(N−1)c),foreachi∈N,F:∃x∞>0s.tsup0≤t<∞D(X[t,ω])≤x∞andlimt→∞D(V[t,ω])=0.Then,itfollowsfromProposition4.1thatP(F|A∩B)=1.Hence,weuseProposition4.2,Lemma4.4toseeP(F)≥P(F∩(A∩B))=P(A∩B)P(F|A∩B)=1−P(Ac∪Bc)≥1−(P(Ac)+P(Bc))=P(A)−P(Bc)≥p1(n)−p2(n)→1asn→∞.Thisimpliesourdesiredresult.5.ThedwellingtimeprocessandthecontinuousCSmodelInthissection,weconsidertwoexplicitprocessesforthedwellingtimeprocessT`whichisequivalenttospecifyingtheswitchingtimeprocessforthechangeofnetworktopologies,andshowthattheysatisfyoneofthekeyconditionsinourframework:limn→∞P
ω:ai(N−1)−1X`=a(i−1)(N−1)T`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
=0.Moreover,wediscussanimprovedstochasticﬂockingestimatecomparedtoauthors’recentwork[15].Ω5.1.Dwellingtimeprocess.Inthissubsection,wepresenttwoexplicitrandomsequence(T`)`≥0satisfying(A2).•(Poisson’sprocess):SupposethatthedwellingtimeprocessT`isgivenbyasequenceofPoisson’srandomvariablewitharateλ`:P(Nλ`=k)=(λ`)k k!e−λ`,k=0,1,2,···.Proposition5.1.SupposethattheprocessT`=Nλ`isasequenceofindependentPoissonrandomvariableswiththeparametersλ`satisfyingλmax:=sup`≥0λ`<∞,andlet(X,V)beasolutionprocessto(1.1).Then,foranyc>0,wecanchooseM∈Nsothatwehavethefollowingestimate:P
ω:ai(N−1)−1X`=a(i−1)(N−1)Nλ`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
≤p2(n),where{p2(n)}isasequencesatisfyinglimn→∞p2(n)=0.
16DONG,HA,JUNG,ANDKIMProof.First,weuseP(∪∞
i=1Bi)≤P∞
i=1P(Bi)togetP
ω:ai(N−1)−1X`=a(i−1)(N−1)Nλ`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
≤∞Xi=1P
ω:ai(N−1)−1X`=a(i−1)(N−1)Nλ`(ω)≥M(n+bclogi(N−1)c)
=∞Xi=1Pω:N˜λi(ω)≥M(n+bclogi(N−1)c),where˜λi=ai(N−1)−1P`=a(i−1)(N−1)λ`andweusedtheindependencetogetai(N−1)−1X`=a(i−1)(N−1)Nλ`∼N˜λi.Forsimplicity,wesetH(i,n):=n+bclogi(N−1)c,andwehave,fori≥1,(5.1)˜λi=ai(N−1)−1X`=a(i−1)(N−1)λ`≤ ai(N−1)−a(i−1)(N−1)λmax≤(N−1)λmaxH(i,n).Then,itfollowsfromthedistributionofPoissonrandomvariablesand(5.1)thatPω:N˜λi(ω)≥MH(i,n)=∞Xr=MH(i,n)(˜λi)re−˜λi r!≤(˜λi)MH(i,n)
(MH(i,n))!≤((N−1)λmaxH(i,n))MH(i,n)
(MH(i,n))!≤((N−1)eλmax/M)MH(i,n)
p
2πMH(i,n),wheretheﬁrstinequalityfollowsfromtheTaylorapproximationfore˜λi:e˜λi−MH(i,n)−1Xr=0(˜λi)r r!=(˜λi)MH(i,n)
(MH(i,n))!dMH(i,n)
dxMH(i,n)ex

x=λ∈(0,˜λi)≤(˜λi)MH(i,n)e˜λi
(MH(i,n))!,andthelastoneisfromtheStirling’sapproximation:n!≥√
2πnn+1
2e−nforalln∈N.ForM>(N−1)eλmax,oneobtains∞Xi=1Pω:N˜λi(ω)≥M(n+bclogi(N−1)c)≤∞Xi=1((N−1)eλmax/M)MH(i,n)
p
2πMH(i,n)
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES17≤1
√
2πMn∞Xi=1((N−1)eλmax/M)M(n+clogi(N−1)−1)=((N−1)eλmax/M)M(n+clog(N−1)−1)
√
2πMn∞Xi=1iMclog((N−1)eλmax/M)=:p2(n)→0,asn→∞.ThelastseriesconvergesifandonlyifMclog((N−1)eλmax/M)<−1.ThisholdsforsuﬃcientlylargeM>0,sincelimM→+∞Mclog((N−1)eλmax/M)=−∞.•(Geometricprocess):Next,weconsiderageometricprocessforT`.SupposethatT`=Yp`isasequenceofindependentrandomvariablesfollowinggeometricdistributions:(5.2)P(Yp`=k)=(1−p`)kp`,k=0,1,2,···,withtheparameterssatisfyingpmin:=inf`≥0p`>1
2.Weﬁrstbeginwithatechnicallemma.
Lemma5.1.ForpositiveintegersAandB,onehas∂
∂xk
Xj`≥0,PB
`=1j`≤A BY`=1(1−x`)j`x`!
>0,for0<x1,···,xB<1,k=1,···,B.Proof.Withoutlossofgenerality,itsuﬃcestoconsiderk=1case.Then,onehasXj`≥0,PB
`=1j`≤A BY`=1(1−x`)j`x`!=Xj`≥0,PB
`=2j`≤AA−PB
`=2j`Xj1=0 BY`=1(1−x`)j`x`!=Xj`≥0,PB
`=2j`≤A BY`=2(1−x`)j`x`!1−(1−x1)A−PB
`=2j`+1.Hence,wehave∂
∂xk
Xj`≥0,P`j`<A BY`=1(1−x`)j`x`!
=Xj`≥0,PB
`=2j`≤A BY`=2(1−x`)j`x`!A−BX`=2j`+1(1−x1)A−PB
`=2j`>0.
18DONG,HA,JUNG,ANDKIMProposition5.2.Supposethat(Yp`)isageometricprocesswhosedistributionsaregivenby(5.2),andlet(X,V)beasolutionprocessto(1.1).Thenforanyc>0,wecanchooseM∈NsuchthatP
ω:ai(N−1)−1X`=a(i−1)(N−1)Yp`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
≤p2(n),wherep2(n)→0asn→∞.Proof.NotethatP
ω:ai(N−1)−1X`=a(i−1)(N−1)Yp`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
≤∞Xi=1P
ω:ai(N−1)−1X`=a(i−1)(N−1)Yp`(ω)≥M(n+bclogi(N−1)c)
.Forsimplicity,wesetJ(i,n):=ai(N−1)−a(i−1)(N−1),H(i,n):=n+bclogi(N−1)c.Then,fori≥1,P
ω:ai(N−1)−1X`=a(i−1)(N−1)Yp`(ω)<MH(i,n)
=Xj`≥0,P`j`<MH(i,n)
ai(N−1)−1Y`=a(i−1)(N−1)(1−p`)j`p`
≥Xj`≥0,P`j`<MH(i,n)
ai(N−1)−1Y`=a(i−1)(N−1)(1−pmin)j`pmin
=MH(i,n)−1Xr=0Xj`≥0,P`j`=r(1−pmin)rpJ(i,n)min=MH(i,n)−1Xr=0r+J(i,n)−1J(i,n)−1(1−pmin)rpJ(i,n)min,(5.3)ΩwhereweusedLemma5.1intheinequality.Then,itfollowsfrom(5.3)thatP
ω:ai(N−1)−1X`=a(i−1)(N−1)Yp`(ω)≥MH(i,n)
=pJ(i,n)min
1
pJ(i,n)min−MH(i,n)−1Xr=0r+J(i,n)−1J(i,n)−1(1−pmin)r

DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES19≤J(i,n)+MH(i,n)−1J(i,n)−1(1−pmin)MH(i,n)
pMH(i,n)min≤J(i,n)+MH(i,n)J(i,n)(1−pmin)MH(i,n)
pMH(i,n)min≤M(N−1)H(i,n)(N−1)(i,n)(1−pmin)MH(i,n)
pMH(i,n)min≤e
√
2π 1+N−1
M1+N−1
MM
N−1N−1
M!MH(i,n)(1−pmin)MH(i,n)
pMH(i,n)min,whereweusedJ(i,n)≤(N−1)H(i,n)andtheﬁrstinequalityfollowsfromtheTaylorapproximationfor(1−x)−J(i,n):1
pJ(i,n)min−MH(i,n)−1Xr=0r+J(i,n)−1r(1−pmin)r=(1−pmin)MH(i,n)
(MH(i,n))!dMH(i,n)
dxMH(i,n)1
(1−x)J(i,n)

x=p∈(0,1−pmin)≤J(i,n)+MH(i,n)−1MH(i,n)(1−pmin)MH(i,n)
pJ(i,n)+MH(i,n)min,andthelastoneisfromthefollowinginequalityobtainedbyStirling’sapproximation:n+mm=(n+m)!
n!m!≤e(n+m)n+m+1
2e−n−m
√
2πnn+1
2e−n·√
2πmm+1
2e−m=e(n+m)n+m+1
2
2πnn+1
2mm+1
2=e(n+m)n+m
2πnnmmr
1
n+1
m≤e(n+m)n+m
√
2πnnmm=e
√
2π1+m n1+m nn mm nn,n,m≥1.Now,wechooseM∈NlargeenoughsothatC(M):=1+N−1
M1+N−1
MM
N−1N−1
M1−pmin pmin<1andMclogC(M)<−1,whichispossiblesincelimM→∞C(M)=1−pmin pmin<1.Then,oneobtains∞Xi=1P
ω:ai(N−1)−1X`=a(i−1)(N−1)Yp`(ω)≥MH(i,n)
≤∞Xi=1e
√
2π(C(M))MH(i,n)≤e
√
2π∞Xi=1(C(M))M(n+clogi(N−1)−1)
20DONG,HA,JUNG,ANDKIM≤e
√
2π(C(M))M(n+clog(N−1)−1)∞Xi=1iMclogC(M):=p2(n)→0,asn→∞.5.2.ThecontinuousCSmodel.Inthissubsection,weuseourresultsforthediscreteCucker-Smalemodeltoextendthepreviousresult[15]forthecontinuousCucker-Smalemodel.Inthepreviouswork,therandomvariablest`+1−t`areassumedtobei.i.dwithaprobabilitydensityfunctionfwithacompactsupport.Moreprecisely,theprobabilitydensityfunctionfisassumedtobesupportedonaﬁniteinterval[a,b]with0<a<b<∞.Basedonpreviousestimatesin[15]andthisworktogetherwithsuitableassumptions,wecanconsiderthecasewhenfissupportedon[a,∞).Insystem(1.1),weadopttheswitchinglawσ:[0,∞)×Ω→{1,···,NG}andthesequence{t`+1−t`}l≥0satisfythefollowingconditions:•Foreach`≥0,t`+1−t`isgivenbyt`+1−t`=a+T`,wherea>0isaconstantandT`’sarenonnegative,independentrandomvariables.•Foreach`≥0andω∈Ω,σt(ω)isconstantontheintervalt∈[t`(ω),t`+1(ω)).•{σt`}`≥0isasequenceofi.i.d.randomvariablessuchthatforany`≥0,P(σt`=k)=pk,foreachk=1,···,NG,wherep1,···,pNG>0aregivenpositiveconstantssatisfyingp1+···+pNG=1.Foreachk=1,···,NG,letGk=(V,Ek)bethek-thadmissibledigraph.Then,foreacht≥0andω∈Ω,thetime-dependentnetworktopology(χσ
ij)=(χσt(ω)ij)isdeterminedbyχσt(ω)ij:=1if(j,i)∈Eσt(ω),0if(j,i)/∈Eσt(ω).Fortechnicalreasonsandwithoutlossofgenerality,weassumethateachadmissibledigraphGkhasaself-loopateachvertex,andwedeﬁnetheuniongraphofGσ[t,ω]fors0≤t<s1andω∈Ω:G([s0,s1))(ω):=s1[t=s0Gσt(ω)= V,s1[t=s0Eσt(ω)!.Now,weimposethefollowingassumptions(B)onthesetofadmissibledigraphsandthedistributionsoftheswitchingtimes:•(B1):TheuniondigraphoftheelementsofShasaspanningtree.Inotherwords,[1≤k≤NGGk=
V,[1≤k≤NGEk
hasaspanningtree.•(B2):ForsomeM>0,therandomvariablessatisfyP
ω:ai(N−1)−1X`=a(i−1)(N−1)T`(ω)≥M(n+bclogi(N−1)c),forsomei∈N
≤˜p(n),
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES21where˜p(n)→0asn→∞.Then,thesimilaranalysistogetherwiththeproofin[15]usedinprevioussectionsyieldsthefollowingimprovedresult.
Theorem5.1.SupposethatparametersN,κ,choiceprobabilitypk’sandcommunicationweightφsatisfy(a(N−1)+M)κ
min1≤k≤NGlog1
1−pk<1,1
φ(r)=O(rε)asr→∞,whereεisapositiveconstantsatisfying0≤ε<1
N−1−(a(N−1)+M)κ
(N−1)min1≤k≤NGlog1
1−pk,andlet(X,V)beasolutionprocessto(1.1).Then,onehasthestochasticﬂocking:Pω:∃x∞>0s.tsup0≤t<∞D(X(t,ω)≤x∞,andlimt→∞D(V(t,ω))=0=1.Remark5.1.Inthepreviouswork[15],weassumedthatthesequenceofdwellingtimes{t`+1−t`}arei.i.dwithaprobabilitydensityfunctionfsupportedonaﬁniteinterval[a,b]with0<a<b<∞.Thiscompactsupportconditionwaspreviouslyassumedtogettheupperboundestimateforthesequence{t∗
r(N−1)−t∗
(r−1)(N−1}r∈NwhichalsoappearsinProposition4.1.However,thankstoourcurrentanalysis,thecompactsupportconditionoffcanbeweakenedandwemayregardfasbeingsupportedon[a,∞),whichisanimprovementfromthepreviouspaper.6.ConclusionInthiswork,weprovidedasuﬃcientframeworkforthestochasticﬂockingforthediscreteCucker-Smalemodelwithrandomlyswitchingtopologies.Oursuﬃcientframeworkinvolvesconditionsontheadmissiblenetworktopologiesandswitchinginstants.First,weassumethattheunionofnetworktopologiesinanadmissiblesetcontainsatleastonespanningtreesothatalltheparticleswillbeeventuallyconnectedinthespace-timedomain.Second,weassumethatthedwellingtimesareshortenoughinaprobabilisticsensesothatalltheparticlescommunicatesuﬃcientlyofteninawholetopological-temporaldomain.Underthesetwoassumptions,weshowedthatthediscreteCucker-Smalemodelexhibitsasto-chasticﬂockingforgenericinitialdataasymptotically.Fordeﬁniteness,wealsoshowthatthePoissonprocessandthegeometricprocessforthedwellingtimeprocessbothsatisfythesecondassumptioninourframework.Althoughourcurrentworkcorrespondstothedis-creteanalogofthecontinuousone,wecanimprovetheearlierresult[15]forthecontinuousC-Smodelbyremovingthecompactsupportassumptionontheprobabilitydistributionfunctionforswitchingtimeprocess.Therearemanyotherissuesthatwedidnotinvestigateinthiswork.Forexample,ouranalysisfocusesonlyonglobalﬂocking.However,asnoticedinliterature,evenforall-to-allcouplings,theCucker-Smaleensemblemaynotexhibitaglobalstochasticﬂockingdependingonthedecayingnatureofthecommunicationweightfunction,namely,formationoflocalﬂocking.Moreover,inthiswork,wedidnotconsiderotherphysicaleﬀectssuchastime-delay,randommedia,interactionwithotherinternalvariables,etc.Theseissueswillbetreatedinourfuturework.
22DONG,HA,JUNG,ANDKIMReferences[1]S.AhnandS.-Y.Ha:StochasticﬂockingdynamicsoftheCucker-Smalemodelwithmultiplicativewhitenoises.J.Math.Phys.51(2010)103301.[2]G.Albi,N.Bellomo,L.Fermo,S.-Y.Ha,J.Kim,L.Pareschi,D.PoyatoandJ.Soler:Vehiculartraﬃc,crowds,andswarms.Onthekinetictheoryapproachtowardsresearchperspectives.Math.ModelsMethodsAppl.Sci.29(2019)1901-2005.[3]M.Ballerini,N.Cabibbo,R.Candelier,A.Cavagna,E.Cisbani,I.Giardina,V.Lecomte,A.Orlandi,G.Parisi,A.Procaccini,M.VialeandV.Zdravkovic:Interactionrulinganimalcollectivebehaviordependsontopologicalratherthanmetricdistance:Evidencefromaﬁeldstudy.Proc.Natl.Acad.Sci.USA105(2008)1232-1237.[4]F.Bolley,J.A.CanizoandJ.A.Carrillo:Stochasticmean-ﬁeldlimit:Non-Lipschitzforcesandswarm-ing.Math.ModelsMethodsAppl.Sci.21(2011)2179-2210.[5]J.A.Canizo,J.A.CarrilloandJ.Rosado:Awell-posednesstheoryinmeasuresforsomekineticmodelsofcollectivemotion.Math.ModelsMethodsAppl.Sci.21(2011)515-539.[6]J.A.Carrillo,M.Fornasier,J.RosadoandG.Toscani:AsymptoticﬂockingdynamicsforthekineticCucker-Smalemodel.SIAMJ.Math.Anal.42(2010)218236.[7]Y.-P.Choi,S.-Y.HaandZ.Li:EmergentdynamicsoftheCucker-Smaleﬂockingmodelanditsvariants,inActiveParticles,Vol.1:Theory,Models,Applications.eds.N.Bellomo,P.DegondandE.Tadmor,ModelingandSimulationinScienceandTechnology(Birkh¨auser,2017),pp.299-331.[8]Y.-P.ChoiandJ.Haskovec:Cucker-Smalemodelwithnormalizedcommunicationweightsandtimedelay.Kinet.Relat.Models10(2017)1011-1033.[9]Y.-P.ChoiandZ.Li:EmergentbehaviorofCucker-Smaleﬂockingparticleswithheterogeneoustimedelays.Appl.Math.Lett.86(2018)49-56.[10]F.CuckerandJ.-G.Dong:Avoidingcollisionsinﬂocks.IEEETrans.Automat.Control55(2010)1238-1243.[11]F.CuckerandJ.-G.Dong:Onﬂocksinﬂuencedbyclosestneighbors.Math.ModelsMethodsAppl.Sci.26(2016)2685-2708.[12]F.CuckerandJ.-G.Dong:Onﬂocksunderswitchingdirectedinteractiontopologies.SIAMJ.Appl.Math79(2019)95-110.[13]F.CuckerandE.Mordecki:Flockinginnoisyenvironments.J.Math.PuresAppl.89(2008)278-296.[14]F.CuckerandS.Smale:Emergentbehaviorinﬂocks.IEEETrans.Automat.Control52(2007)852-862.[15]J.-G.Dong,S.-Y.Ha,J.JungandD.Kim:OnthestochasticﬂockingoftheCucker-Smaleﬂockwithrandomlyswitchingtopologies.availableatarXiv:1911.07390.[16]J.-G.Dong,S.-Y.HaandD.Kim:EmergentbehaviorsofcontinuousanddiscretethermomechanicalCucker-Smalemodelsongeneraldigraphs.Math.ModelsMeth.Appl.Sci.29(2019)589-632.[17]J.-G.Dong,S.-Y.HaandD.Kim:Interplayoftime-delayandvelocityalignmentintheCucker-Smalemodelonageneraldigraph.DiscreteContin.Dyn.Syst.-Ser.B.22(2019)1-28.[18]J.-G.DongandL.Qiu:FlockingoftheCucker-Smalemodelongeneraldigraphs.IEEETrans.Automat.Control62(2017)5234-5239.[19]F.DalmaoandE.Mordecki:Cucker-Smaleﬂockingunderhierarchicalleadershipandrandominterac-tions.SIAMJ.Appl.Math.71(2011),1307-1316.[20]F.DalmaoandE.Mordecki:HierarchicalCucker-Smalemodelsubjecttorandomfailure.IEEETrans.Automat.Control57(2012)1789-1793.[21]R.Erban,J.HaskovecandY.Sun:OnCucker-Smalemodelwithnoiseanddelay.SIAMJ.Appl.Math.76(2016)1535-1557.[22]M.Fornasier,J.HaskovecandG.Toscani:FluiddynamicdescriptionofﬂockingviaPovzner-Boltzmannequation.PhysicaD240(2011)21-31.[23]S.-Y.Ha,D.KimandZ.Li:EmergentﬂockingdynamicsofthediscretethermodynamicCucker-Smalemodel.ToappearinQuart.Appl.Math.[24]S.-Y.Ha,D.Kim,D.KimandW.Shim:Flockingdynamicsoftheinertialspinmodelwithamulti-plicativecommunicationweight.J.NonlinearSci.29(2019)1301-1342.[25]S.-Y.Ha,J.KimandT.Ruggeri:EmergentbehaviorsofthermodynamicCucker-Smaleparticles.SIAMJ.Math.Anal.50(2018)3092-3121.[26]S.-Y.Ha,J.KimandX.Zhang:UniformstabilityoftheCucker-Smalemodelanditsapplicationtothemean-ﬁeldlimit.Kinet.Relat.Mod.11(2018)1157-1181.
DISCRETECUCKER-SMALEMODELWITHRANDOMLYSWITCHINGTOPOLOGIES23[27]S.-Y.Ha,K.LeeandD.Levy:Emergenceoftime-asymptoticﬂockinginastochasticCucker-Smalesystem.Commun.Math.Sci.7(2009)453-469.[28]S.-Y.HaandJ.-G.Liu:AsimpleproofofCucker-Smaleﬂockingdynamicsandmeanﬁeldlimit.Com-mun.Math.Sci.7(2009)297-325.[29]S.-Y.HaandE.Tadmor:Fromparticletokineticandhydrodynamicdescriptionofﬂocking.Kinet.Relat.Models1(2008)415-435.[30]Y.HeandX.Mu:Cucker-Smaleﬂockingsubjecttorandomfailureongeneraldigraphs.Automatica106(2019)54-60.[31]N.E.Leonard,D.A.Paley,F.Lekien,R.Sepulchre,D.M.FratantoniandR.E.Davis:Collectivemotion,sensornetworksandoceansampling.Proc.IEEE95(2007)48-74.[32]Z.LiandS.-Y.Ha:OntheCucker-Smaleﬂockingwithalternatingleaders.Quart.Appl.Math.73(2015)693-709.[33]Z.LiandX.Xue:Cucker-Smaleﬂockingunderrootedleadershipwithﬁxedandswitchingtopologies.SIAMJ.Appl.Math.70(2010)3156-3174.[34]Y.LiuandJ.Wu:FlockingandasymptoticvelocityoftheCucker-Smalemodelwithprocessingdelay.J.Math.Anal.Appl.415(2014)53-61[35]Lu,W.,Atay,F.M.andJost,J.:Consensusandsynchronizationindiscrete-timenetworksonmulti-agentswithstochasticallyswitchingtopologiesandtime-delays..Netw.Heterog.Media62011,329-349.[36]S.MotschandE.Tadmor:Anewmodelforself-organizeddynamicsanditsﬂockingbehavior,J.Stat.Phys.144(2011)923-947.[37]S.MotschandE.Tadmor:Heterophiliousdynamics:Enhancedconsensus.SIAMRev.56(2014)577-621[38]J.Park,H.KimandS.-Y.Ha:Cucker-Smaleﬂockingwithinter-particlebondingforces.IEEETrans.Automat.Control55(2010)2617-2623.[39]C.PignottiandI.R.Vallejo:FlockingestimatesfortheCucker-Smalemodelwithtimelagandhierar-chicalleadership.J.Math.Anal.Appl.464(2017)1313-1332.[40]D.PoyatoandJ.Soler:Euler-typeequationsandcommutatorsinsingularandhyperboliclimitsofkineticCucker-Smalemodels.Math.ModelsMethodsAppl.Sci.27(2017)1089-1152.[41]L.Ru,Z.LiandX.Xue:Cucker-Smaleﬂockingwithrandomlyfailedinteractions.J.FranklinInst.352(2015)1099-1118.[42]J.Shen:Cucker-Smaleﬂockingunderhierarchicalleadership.SIAMJ.Appl.Math.68(2007)694-719.[43]J.TonerandY.Tu:Flocks,herds,andschools:AquantitativetheoryofﬂockingPhys.Rev.E58(1998)4828-4858.[44]T.Vicsek,A.Czir´ok,E.Ben-Jacob,I.CohenandO.Schochet:Noveltypeofphasetransitioninasystemofself-drivenparticles.Phys.Rev.Lett.75(1995)1226-1229.[45]T.VicsekandA.Zefeiris:Collectivemotion.Phys.Rep.517(2012)71-140.[46]C.W.Wu:Synchronizationandconvergenceoflineardynamicsinrandomdirectednetworks.IEEETrans.Automat.Control51(2006),1207–1210.
24DONG,HA,JUNG,ANDKIM(Jiu-GangDong)DepartmentofMathematics,HarbinInstituteofTechnology,Harbin150001,P.R.ChinaE-mailaddress:jgdong@hit.edu.cn(Seung-YealHa)DepartmentofMathematicalSciencesandResearchInstituteofMathematics,SeoulNationalUniversity,Seoul08826
andSchoolofMathematics,KoreaInstituteforAdvancedStudy,Hoegiro85,Seoul02455,Korea(Republicof)E-mailaddress:syha@snu.ac.kr(JinwookJung)DepartmentofMathematicalSciences,SeoulNationalUniversity,Seoul08826,Korea(Republicof)E-mailaddress:warp100@snu.ac.kr(DoheonKim)SchoolofMathematics,KoreaInstituteforAdvancedStudy,
Hoegiro85,Seoul02455,Korea(Republicof)E-mailaddress:doheonkim@kias.re.kr

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

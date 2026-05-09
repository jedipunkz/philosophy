---
source: "https://arxiv.org/abs/2102.07329v1"
title: "Symmetric Operations on Domains of Size at Most 4"
author: "Zarathustra Brady, Holden Mui"
year: "2021"
publication: "arXiv preprint / math.RA"
download: "https://arxiv.org/pdf/2102.07329v1"
pdf: "https://arxiv.org/pdf/2102.07329v1"
captured_at: "2026-05-09T13:05:26Z"
updated_at: "2026-05-09T13:05:26Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche Thus Spoke Zarathustra"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Symmetric Operations on Domains of Size at Most 4

- 著者: Zarathustra Brady, Holden Mui
- 年: 2021
- 掲載情報: arXiv preprint / math.RA
- 情報源: [arxiv](https://arxiv.org/abs/2102.07329v1)
- ダウンロード: https://arxiv.org/pdf/2102.07329v1
- PDF: https://arxiv.org/pdf/2102.07329v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

To convert a fractional solution to an instance of a constraint satisfaction problem into a solution, a rounding scheme is needed, which can be described by a collection of symmetric operations with one of each arity. An intriguing possibility, raised in a recent paper by Carvalho and Krokhin, would imply that any clone of operations on a set $D$ which contains symmetric operations of arities $1, 2, \ldots, \lvert D \rvert$ contains symmetric operations of all arities in the clone. If true, then it is possible to check whether any given family of constraint satisfaction problems is solved by its linear programming relaxation. We characterize all idempotent clones containing symmetric operations of arities $1, 2, \ldots, \lvert D \rvert$ for all sets $D$ with size at most four and prove that each one contains symmetric operations of every arity, proving the conjecture above for $\lvert D \rvert \leq 4$.

## PDF Text

arXiv:2102.07329v1 [math.RA] 15 Feb 2021
SymmetricOperationsonDomainsofSizeatMost4ZarathustraBradyandHoldenMuiFebruary16,2021AbstractToconvertafractionalsolutiontoaninstanceofaconstraintsatisfactionproblemintoasolution,aroundingschemeisneeded,whichcanbedescribedbyacollectionofsymmetricoperationswithoneofeacharity.Anintriguingpossibility,raisedinarecentpaperbyCarvalhoandKrokhin,wouldimplythatanycloneofoperationsonasetDwhichcontainssymmetricoperationsofarities1,2,...,|D|containssymmetricoperationsofallaritiesintheclone.Iftrue,thenitispossibletocheckwhetheranygivenfamilyofconstraintsatisfactionproblemsissolvedbyitslinearprogrammingrelaxation.Wecharacterizeallidempotentclonescontainingsymmetricoperationsofarities1,2,...,|D|forallsetsDwithsizeatmostfourandprovethateachonecontainssymmetricoperationsofeveryarity,provingtheconjectureabovefor|D|≤4.1Introduction
TheConstraintSatisfactionProblem,commonlyabbreviatedastheCSP,isthedecisionproblemwherewearegivenalistofvariablesandalistofconstraintsonthevariables,andwemustdeterminewhetherornotthereexistsanassignmentofthevariableswhichsatisﬁeseveryconstraint.WhilethisdecisionproblemisNP-hardingeneral,certainclassesofCSPscanbesolvedinpolynomialtime.AmongthosearetheCSPsthataresolvedbytheirlinearprogrammingrelaxation;thatis,a“fractionalsolution”toaninstanceofsuchaCSPcanberoundedtoasolution.SuchCSPshavebeencharacterizedastheCSPsforwhichthecloneofoperationspreservingitsrelationscontainssymmetricoperationsofeveryarity.Inthispaper,weinvestigateaconjecturethat,iftrue,givesasuﬃcientconditionforaclonetocontainsymmetricoperationsofeveryarity.
Conjecture1.SupposeacloneoveradomainDcontainssymmetricoperationsofarities1,2,...,|D|.Thenitcontainssymmetricoperationsofeveryarity.Thisconjectureisaweakformofanopenproblemmentionedinsection6of[7]:itsauthorsspeculatethataﬁnitealgebraicstructureAhassymmetrictermsofeveryarityifandonlyifAhasnosubquotientBsuchthattheautomorphismgroupofBcontainsapairofautomorphismswithnocommonﬁxedpoint.Inthispaper,weproveConjecture1for|D|≤4.Theorem2.SupposeacloneoveradomainDwith|D|≤4containssymmetricoperationsofarities1,2,...,|D|.Thenitcontainssymmetricoperationsofeveryarity.Intheappendix,wesketchamoreambitiousconjectureaboutthesolvabilityofcertaintypesofweaklyconsistentconstraintsatisfactionproblemsattachedtoanalgebraicstructuresatisfyingtheassumptionofConjecture1.
1.1Motivation
IfConjecture1istrue,thenitgivesaneﬃcientwayofdeterminingwhetherornotcombinatorialpuzzlescanbesolvedusingsystemsoflinearinequalities.1
TounderstandConjecture1’simplications,itishelpfultocharacterizethetypesofcombinatorialpuzzleswewanttolookat;agoodexampleofsuchapuzzleisSudoku.SudokuisaconstraintsatisfactionproblemoverthedomainD={1,2,3,4,5,6,7,8,9}withtentypesofconstraints.Thereisthe9-aryconstraintwhichassertsthatitsinputsarepairwisedistinct,andtherearenineunaryconstraints,eachoftheform“thisvariableisequaltod”forsomed∈D.Instancesofthisconstraintsatisfactionproblemhave81variables,andthevariablesthattheconstraintsapplytodependonthespeciﬁcinstance.OtherexamplesofCSPsincludeHornSATand3Coloring:HornSATistheconstraintsatisfactionproblemwherethetaskistodeterminewhetherornotasetofHornclauses(implications)admitsanassignmentofthevariablessatisfyingeachHornclause,and3Coloringistheconstraintsatisfactionproblemwhosetaskistodeterminewhetherornotagivengraphadmitsa3-coloring.ThelinearprogrammingrelaxationofaninstanceofaCSPis,informally,thesetofalllocallyconsistentprobabilitydistributionsoveritsvariablesandconstraintrelations,andisdeﬁnedbyacollectionoflinearinequalities(arigorousbuttersedeﬁnitionisgiveninsection2).ACSPissolvedbyitslinearprogrammingrelaxationiﬀthereisawaytoturnpointsinthelinearprogrammingrelaxationofeveryinstanceoftheCSPtoasolutionofthatCSP,knownasaroundingscheme.WhileﬁndingsolutionstoinstancesofageneralCSPisNP-hard,solvingthelinearprogrammingrelax-ationofanyCSPonlytakespolynomialtime,sothereisaneﬃcientwaytoﬁndasolutiontoaninstanceofaCSPifitissolvedbyitslinearprogrammingrelaxation.Usuallythisisnotthecase,butforthesespecialCSPs,localprobabilitydistributionsofsolutions,calledfractionalsolutions,canbeconvertedtotruesolutionsusingtheroundingscheme.TheCSPssolvedbytheirlinearprogrammingrelaxationhavebeencharacterizedin[9].TheyarepreciselytheonesforwhichthecloneofoperationspreservingeachrelationdeﬁningtheCSPcontainssymmetricoperationsofeveryarity.
Theorem3(Theorem2of[9]).TheCSPdeﬁnedbyacollectionofrelationsΓissolvedbythelinearprogrammingrelaxationifandonlyifthecloneofoperationsthatpreserveseachrelationinΓcontainssymmetricoperationsofeveryarity.AnequivalentcharacterizationofcloneswithsymmetricoperationsofeveryarityappearsinarecentarticlebyButtiandDalmauaboutsolvingCSPswithdistributedalgorithms[6].Intheirsetup,eachagenthasaccesstoasinglevariableortoasingleconstraint,agentscancommunicateonlywhenoneownsaconstraintinvolvingthevariableownedbytheother,andtheagentsareanonymous,sothereisnoobviouswaytoelectaleader.
Theorem4(Theorem6of[6]).TheCSPdeﬁnedbyacollectionofrelationsΓcanbesolvedinthedistributedsettingdescribedaboveifandonlyifthecloneofoperationsthatpreserveseachrelationinΓcontainssymmetricoperationsofeveryarity.Ifitcanbesolvedinthedistributedsettingatall,thenitcanbesolvedinthissettinginpolynomialtime.Forbrevity,wecallsuchclonesround,astheycanbeusedtoconstructaroundingschemethatturnsfractionalsolutionsintosolutionsofinstancesofCSPs.However,determiningwhetherornotagivencloneisroundisdiﬃcult(possiblyundecidable);thatis,unlessConjecture1istrue.RecallthatConjecture1statesthattheexistenceofsymmetricoperationsofarities1,2,...,|D|isasuﬃcient(andnecessary,butthisdirectionisobvious)conditionforaclonetocontainsymmetricoperationsofeveryarity.Equivalently,Conjecture1assertsthatgiven|D|symmetricoperationsofarities1,2,...,|D|,onecancreateasymmetricoperationofanydesiredaritybycomposingthe|D|“base”operationsinsomeway.IfConjecture1istrue,thendeterminingwhetherornotacloneisroundbecomesaﬁnitecasecheck!AlthoughConjecture1remainsanopenprobleminfullgenerality,weproveConjecture1forallclonesoveradomainofsizeatmost4.Weachievethisbyclassifyingallminimalidempotentclonesoveradomainofsizeatmost4satisfyingConjecture1’shypothesisandprovingthateachoneisround.Itisourhopethatthisclassiﬁcationwillhelpfutureresearchersformstrongerhypotheses,verifytheirtruthforalargenumberofexamples,andultimatelytakestepsclosertowardsaproof(oradisproof)ofConjecture1.2
1.2RoadMap
Theremainderofourpaperisorganizedasfollows.Insection2,wegooverdeﬁnitions.Insection3,wesummarizerelatedresults.Insection4,weproveresultsthatsimplifytheenumerationofallminimalidempotentroundclonesoveradomainofsizeatmostfour.Insection5,weenumerateallsuchclones.Insection6,wesketchaplausiblelineofattackonthegeneralcaseofConjecture1.Intheappendix,wedescribeconnectionsbetweenthelinearprogrammingrelaxationofaCSPandcertainweakconsistencyconditions,andweconjectureapreciseconnectionbetweenthem.2Deﬁnitions
2.1ConstraintSatisfactionProblems
Adomain,denotedbythecapitalletterD,isasetofvaluesavariablecanbeassignedto.Ak-aryrelationRoveradomainDisasubsetofthek-foldCartesianproductDk:=D×...×D;kisknownastheconstraint’sarity,denotedar(R).Atuple(a1,...,ak)satisﬁesarelationRif(a1,...,ak)∈R.AconstraintisapairconsistingofarelationRandanar(R)-tupleofvariables.AconstraintsatisfactionproblemisapairP=(D,Γ)whereDisitsdomainandΓisasetofrelationsoverD.Aninstanceofaconstraintsatisfactionproblem(D,Γ)isapairI=(X,T)where•X=x1,...,x|X| isaﬁnitesetofvariables,and•TisasetofconstraintsinvolvingthevariablesinX,suchthateachconstraintrelationisanelementofΓ.Formally,Tisasetofpairs(x,R),whereR∈Γandx∈Xar(R)isthetupleofvariablestherelationRisappliedto.Anassignmentx1=a1,...,xn=anofthevariablestoelementsofDisasolutiontothatinstanceif,foreachpair(x,R)∈T,Rissatisﬁedbythetuplexafterreplacingeachxiinxwithai.Thelinearprogrammingrelaxationofak-aryrelationR=r1,...,r|R| overaﬁnitedomainD=d1,...,d|D| isthepolyhedronin R|D|kdeﬁnedbythesetofallpoints  (v1)d1,(v1)d2,...,(v1)d|D|, (v2)d1,(v2)d2,...,(v2)d|D|,... (vk)d1,(vk)d2,...,(vk)d|D|forwhichthereexistrealspr1,...,pr|R|suchthat0≤pr1,...,pr|R|≤1,Xr∈Rpr=1,and(vi)dj=Xr∈R|ri=djpr3
forall1≤i≤kand1≤d≤|D|.ThelinearprogrammingrelaxationofaninstanceI=(X,T)ofaCSPoveradomainDisthepolyhedronin R|D||X|deﬁnedbythesetofallpoints (x1)d1,(x1)d2,...,(x1)d|D|, (x2)d1,(x2)d2,...,(x2)d|D|,... x|X|d1, x|X|d2,..., x|X|d|D|suchthatforeachpair(x,R)∈T,xliesinR’slinearprogrammingrelaxationwheneachvariableinxisreplacedbyitscorresponding|D|-tuple.AfractionalsolutionofaninstanceofaCSPisapointinsideitslinearprogrammingrelaxation.WesaythataCSPissolvedbyitslinearprogrammingrelaxationif,foreveryinstanceIoftheCSP,theexistenceofafractionalsolutionimpliestheexistenceofasolution.LetP=(D,Γ)beaconstraintsatisfactionproblem,letX=x1,...,x|X| beasetofvariables,andletI= X,{(x1,R1),...,(x|R|,R|R|)}beaninstanceofP.Asteps=(k,(i,j))inIisdeﬁnedtobeaconstraintrelation(xk,Rk)andapairofintegers1≤i,j≤ar(Rk);wethinkofthestepsasconnectingthevariable(xk)itothevariable(xk)j.AcyclepinIisaﬁnitesequenceofstepss1,...,s|p|=(k1,(i1,j1)),..., k|p|,(i|p|,j|p|)forwhich(xk1)j1=(xk2)i2,(xk2)j2=(xk3)i3,(xk|p|−1)j|p|−1=(xk|p|)i|p|,and(xk|p|)j|p|=(xk1)i1,wherethesubscriptnoneach(xi)nrepresentsthevariableinXcorrespondingtothenthcoordinateofxi.IfBisasubsetofDands=(k,(i,j))isastepinI,thenwedeﬁnethesumB+sasB+s:={d∈D|∃a∈Rks.t.ai∈B∧aj=d}andthesumB−sasB+s:={d∈D|∃a∈Rks.t.aj∈B∧ai=d},wherethesubscriptaidenotestheithcoordinateofthetuplea.Ifp=s1,...,s|p|isacycleinI,thenwedeﬁnethesumB+pasB+p:=B+s1+...+spandthesumB−pasB−p:=B−sp−...−s1.2.2Clones
LetDbesomedomain.Anoperationfisafunctionf:Dk→Dforsomepositiveintegerk,knownasitsarity.Anoperationwitharity1isunary,anoperationwitharity2isbinary,andanoperationwitharity3isternary.Ingeneral,anoperationfwitharitykisk-ary,anditsarityisdenotedar(f).Theoutputofanoperationfwithinputsx1,x2,...,xkisdenotedf(x1,x2,...,xk).Letf:Dk→Dbeanoperation.WeextendftoanoperationonvectorsinDnbyapplyingitcoordinatewise,i.e.,f







(a1)1(a1)2...(a1)n



,



(a2)1(a2)2...(a2)n



,...,



(ak)1(ak)2...(ak)n







:=



f((a1)1,(a2)1,...,(ak)1)f((a1)2,(a2)2,...,(ak)2)...f((a1)n,(a2)n,...,(ak)n)



.4
Ak-aryoperationf:Dk→DpreservesarelationRiff(r1,...,rk)∈Rforallr1,...,rk∈R(notethatthearityoffhasnothingtodowiththearityofR).Therelationgeneratedbyfwithgeneratorsx1,x2,...isthesmallestrelationcontainingx1,x2,...thatfpreserves,andisdenotedSgf(x1,x2,...).Foreveryk∈Z+andinteger1≤i≤k,theprojectionoperationπki:Dk→DoveradomainDisdeﬁnedasπki(d1,...,dk):=di.Notethatπ11istheidentityoperationoverD.AcloneoveradomainDisasetOofﬁnite-arityoperationsthatcontainseveryprojectionoperationoverDandisclosedundermultiplecomposition;thatis,iff∈Oisanm-aryoperationandg1,...,gm∈Oaren-aryoperations,thentheoperationh(x1,...,xn):=f(g1(x1,...,xn),...,gm(x1,...,xn))isalsoinO.NotethatOisclosedunderany“natural”wayofcomposingoperationsbecauseeveryprojectionoperationisinO.Acloneiscompatiblewithacyclicautomorphismifthereisarenamingdeﬁnedbyacyclicpermutationofthedomainelementssuchthattherenamedclonecontainsthesameoperationsastheoriginalclone.AsubcloneO0ofacloneOisasubsetofOthatisaclone.ThesubcloneisproperifO06=O.Theclonegeneratedbyasetofoperations{f1,f2,...}isthesmallestclonecontaining{f1,f2,...}andisdenotedhf1,f2,...i.AnoperationfoveradomainDisidempotentiff(x,...,x
|
{z
}ar(f)x’s)=xforallx∈D.AcloneOisidempotentifeveryoperationinOisidempotent.Ak-aryoperationfissymmetriciff(x1,...,xk)=f xσ(1),...,xσ(k)forallx1,...,xk∈Dandpermutationsσ:{1,...,k}→{1,...,k}.Wecallacloneroundifitcontainssymmetricoperationsofeveryarity.WecallacloneoveradomainDsemi-roundifitcontainssymmetricoperationsofarities1,2,...,|D|.TherelationgeneratedbyacloneOwithgeneratorsx1,x2,...isthesmallestrelationthateveryoperationinOpreserves,anditisdenotedSgO(x1,x2,...).LetPbeapropertyofaclone.WesayacloneisminimalwithrespecttopropertyPifitdoesnotcontainapropersubclonewithpropertyP.Forexample,acloneisminimallyroundifitdoesnotcontainaproperroundsubclone,andacloneisminimallysemi-roundifitdoesnotcontainapropersemi-roundsubclone.Lastly,aminimalidempotentroundcloneisanidempotentclonethatisminimallyround,andaminimalidempotentsemi-roundcloneisanidempotentclonethatisminimallysemi-round.2.3Algebraicconcepts
AnalgebraicstructureA=(A,f1,f2,...),alsoknownasanalgebra,isadomainA,whichwecalltheunderlyingsetofA,withsomeoperationsfi:Aki→A,whichwecallthebasicoperationsofA.Algebraicstructureswillalwaysbewritteninblackboardbold.Thesequenceofaritiesk1,k2,...iscalledthesignatureofthealgebraA.5
GivenanalgebraicstructureA=(A,f1,f2,...),wedeﬁnethepowerAm=(Am,f1,f2,...)tobeanalgebraicstructureofthesamesignature,whereeachfiactscoordinatewiseonAm.AsubalgebraBofanalgebraA,denotedasB≤A,consistsofasubsetB⊆AclosedunderthebasicoperationsofA,andbasicoperationswhicharerestrictionsofthebasicoperationsofAtoB.IfBisanysubsetofAm,thenwedeﬁnethesubalgebrageneratedbyB,denotedSgAm(B),tobethesmallestsubalgebraofAmwhichcontainsB.ThecloneofanalgebraicstructureA,writtenClo(A),istheclonegeneratedbythebasicoperationsofA.Whenwestudyalgebraicstructures,wewillmainlybeinterestedinpropertieswhichonlydependontheirclonesinsteadoftheirbasicoperations.ArelationR≤AmonanalgebraAwillalwaysrefertoasubalgebraofAmforsomeintegerm,knownasthearityofR.Alternatively,arelationR≤AmisasubsetofAmwhichiscompatiblewiththebasicoperationsofA;thatis,foreachbasicoperationfiofaritykiandforeverychoiceofkituplesa1,...,aki∈R,wehavef(a1,...,aki)∈R,wherefactscoordinatewise.WhetherornotagivensetR⊆AmdeﬁnesarelationcompatiblewiththebasicoperationsofAonlydependsonthecloneofA.AcongruenceθonanalgebraicstructureAisanequivalencerelationθ≤A2compatiblewiththebasicoperationsofA,andthequotientA/θisanalgebraicstructurewithdomainA/θwiththebasicoperationsdeﬁnedinthenaturalway.IfR≤AmisarelationandIisasubsetofthecoordinates{1,...,m},thenwedeﬁnetheexistentialprojectionπI(R)asπI(R):=x∈AI|∃y∈Rs.t.yi=xi∀i∈I .Forbrevity,deﬁneπi,j,...(R)tobeπ{i,j,...}(R).ArelationR≤Amissubdirect,denotedR≤sdAm,iftheithprojectionπi(R)isequaltoAforeveryinteger1≤i≤m.IfR,S≤A2arebinaryrelations,thenwedeﬁnetheircompositionasR◦S:=(x,z)∈A2|∃y∈As.t.(x,y)∈R∧(y,z)∈S .WedeﬁnethereverseofthebinaryrelationR,denotedR−,asR−:={(y,x)∈A2|(x,y)∈R}.IfR≤sdA2,thenwedeﬁnethelinkingcongruenceofRontheﬁrstcoordinatetobethecongruence[n≥1(R◦R−)◦n.Thelinkingcongruenceonthesecondcoordinateisdeﬁnedsimilarly,withRandR−swapped.IfB⊆AisasetandR≤A2isabinaryrelation,thenwedeﬁnethesumB+RasB+R:={y∈A|∃x∈Bs.t.(x,y)∈R}.andwedeﬁnethediﬀerenceB−RasB−R:=B+R−={x∈A|∃y∈Bs.t.(x,y)∈R}.IfAisacollectionofalgebraicstructureswhichallhavethesamesignature,thenwedeﬁneP(A)tobethecollectionofallproductsofalgebrasinA,S(A)tobethecollectionofallsubalgebrasofalgebrasinA,andH(A)tobethecollectionofallhomomorphicimagesofalgebrasinA(whichisthecollectionofallalgebraswhichareisomorphictothequotientofsomealgebraAinAbysomecongruenceonA).IfB∈HS(A),thenwecallBasubquotientofA.AcloneisTaylorifitcontainsidempotentoperationsthatsatisfysomefunctionalequationthatcannotbesatisﬁedbyprojectionoperations.AnalgebraiscalledTaylorifitscloneisTaylor.ByBirkhoﬀ’sHSPtheorem,anidempotentalgebraisTaylorifandonlyifHSP(A)doesnotcontainatwo-elementalgebrawitheachofitsbasicoperationsequaltoaprojection.6
2.4Miscellaneous
ForsetsAandB,thesetA+BisdeﬁnedasA+B:={a+b|a∈A,b∈B}.Thefunctionsgn:R→Risdeﬁnedassgn(x):=



−1ifx<00ifx=01ifx>0.Fora(k−1)-aryoperationf,wedeﬁnecf:Dk→Dkascf((x1,...,xk)):=(f(x2,x3,...,xk−1,xk),f(x1,x3,...,xk−1,xk),f(x1,x2,...,xk−1,xk),...,f(x1,x2,x3,...,xk−1)).Wealsodeﬁnecf(x1,...,xk):=cf((x1,...,xk))forconvenience.Additionally,givenatuplex=(x1,...,xk)andak-aryoperationf,deﬁnef(x):=f(x1,...,xk).WesayanoperationfoveradomainDactslikeaheight-1semilatticeoverasubsetD0⊆Dofitsdomainifitisidempotentandf x1,...,xar(f)isthesameﬁxedvaluec∈D0overallnon-constanttuples x1,...,xar(f)∈(D0)ar(f).Wesayacloneactslikeaheight-1semilatticeoverasubsetD0⊆Dofitsdomainifallitsoperationsactlikeaheight-1semilatticeoverD0,andtheconstantcisthesameacrossalloperations.WesayabinaryoperationfoveradomainDactslinearlyoverasubsetD0⊆Dif|D0|isoddandthedomainelementsoff,whenrestrictedtothedomainD0,canberenamedsuchthatthenewoperationf0:(Z/|D0|Z)2→Z/|D0|Zsatisﬁesf0(x,y)=x+y
2(mod|D0|)forallx,y∈Z/|D0|Z.AbinaryoperationfoveradomainDisasemilatticeoperationifthereexistsaposetonDsuchthatfrepresents“join”;thatis,f(x,y)=x∧yforallx,y∈D.Lastly,foralla,b,c∈{−,0,+}wedeﬁnefabctobethesymmetricbinaryoperationfabc
−0+
−
−cb0
c0a+
ba+,andfora,b,c,d∈{0,1,2,3}wedeﬁnefabcdeftobethesymmetricbinaryoperation7
fabcdef
0123
0
0abc1
a1de2
bd2f3
cef3.3RelatedResults
RecallthatourgoalistoverifyConjecture1forallclonesoveradomainofsize4orless;thatis,ourgoalistoprovethatallsemi-roundclonesoveradomainofsizeatmost4areround.TherelevanceofthisproblemtoCSPshasbeendemonstratedin[9].TheauthorsprovethataCSPissolvedbyitslinearprogrammingrelaxationifandonlyifthecloneofalloperationspreservingtherelationsdeﬁningtheCSPisround.Arelatedresultaboutcyclicoperationsisprovedin[1].TheyprovethateveryTayloralgebracontainsacyclicoperationofeveryprimearitygreaterthanthesizeofitsdomain.Itturnsoutthatthe|D|boundistightwhen|D|isprime.Proposition5.LetDbeadomainwithprimecardinality.ThenthereexistsacloneOoverDcontainingsymmetricoperationsofarities1,2,...,|D|−1thatisnotround.Proof.Byconstruction,wecanforceOtobecompatibleacyclicautomorphism.Letpbethecardinalityofthedomain,letD=Z/pZ,anddeﬁnethek-aryoperationfkfork∈{1,2,...,|D|−1}tobeanysymmetricoperationsatisfyingfk(x1+c,...,xk+c)=fk(x1,...,xk)+cforall(x1,...,xk)∈Dkandc∈D.Forinstance,wecantakefk(x1,...,xk):=x1+···+xk k(modp).Thenhf1,f2,...,fp−1iiscompatiblewithacyclicautomorphism,sonop-arysymmetricoperationfpcanexist,asfp(x1,...,xp)cannotbepreservedbytheautomorphism.
Remark6(Lemma4of[7]).Thetheoremstatementisfalseif|D|isnotrequiredtobeprime;theconstructionin[7]provesthata4-arysymmetricoperationf4(w,x,y,z)canbeconstructedfromabinarysymmetricoperationf2(x,y)andaternarysymmetricoperationf3(x,y,z):f4(w,x,y,z)=f3(f2(f2(w,x),f2(y,z)),f2(f2(w,y),f2(x,z)),f2(f2(w,z),f2(x,y)))issymmetric.8
4PreliminaryTheorems
Theseresultswillhelpusclassifythesemi-roundclonesoveradomainofsizeatmostfour,uptoarenamingofthedomainelements.
Theorem7.LetΣbeasetoffunctionalequations.TheneverycloneOcontainingoperationsthatsi-multaneouslysatisfyeachfunctionalequationinΣcontainsaminimalsubclonecontainingoperationsthatsimultaneouslysatisfyeachfunctionalequationinΣ.Proof.ForasetΣ1offunctionalequations,letpropertyPΣ1ofacloneOdenotetheassertionthatOcontainsasetofoperationssimultaneouslysatisfyingeachequationinΣ1.ForanyﬁnitesubsetΣ0⊆Σ,therecanonlybeﬁnitelymanycombinationsofoperationsinOthatsatisfyΣ,sincethedomainisﬁnite.Therefore,theintersectionofeverychainO1)O2)...ofcloneswithpropertyPΣ0alsohaspropertyPΣ0foreveryﬁnitesubsetΣ0⊆Σ.Bythelogicalcompactnesstheorem,OhaspropertyPΣifOhaspropertyPΣ0foreveryﬁnitesubsetΣ0⊆Σ,whichallowsustoconcludethattheintersectionofeverychainO1)O2)...ofcloneswithpropertyPΣalsohaspropertyPΣ.Therefore,everychainintheposetofsubclonesofOwithpropertyPΣ,orderedbyreverseinclusion,hasanupperbound,soZorn’slemmaimpliestheexistenceofaminimalclonewithpropertyPΣ.
Corollary8.Everyroundclonecontainsaminimalroundsubclone.Proof.LetΣbethesetoffunctionalequationsf2(x,y)=f2(y,x)f3(x,y,z)=f3(x,z,y)=f3(y,x,z)=f3(y,z,x)=f3(z,x,y)=f3(z,y,x)...,whereeachlineassertstheexistenceofasymmetricoperationofsomearity.ThenPΣisequivalenttoroundness,sothereexistsaminimalroundsubclonebyTheorem7.
Corollary9.Everysemi-roundclonecontainsaminimalsemi-roundsubclone.Proof.ThisisthesameastheproofofCorollary8,exceptthesetofequationsΣassertstheexistenceofsymmetricoperationsofarities1,2,...,|D|,whereDisthedomain,insteadofsymmetricoperationsofeveryarity.
Theorem10.SupposeOisaminimalcounterexampletoConjecture1overthesmallestpossibledomain.ThenOisidempotent.Proof.ThisisacorollaryofafolkloreresultthatseemstohaveﬁrstshownupinthecontextofCSPsin[5],butwereproduceithere.ItsuﬃcestoshowthattheonlyunaryoperationinaminimalcounterexampleOisπ11.Todothis,letthedomainbeDandletfkbeak-arysymmetricoperationinOforeachk∈{1,2,...,|D|}.Ifthereisanon-identityunaryoperationu(x)∈O,theneithersomeunaryoperationisnotinjectiveorallunaryoperationsarepermutations.Ifallunaryoperationsarepermutations,thenonecanﬁndidempotentsymmetricoperationsofarities1,2,...,|D|.Todothis,deﬁneuk(x)=fk(x,...,x
|
{z
}ktimes),9
whichmustbeapermutation.ThenuNkk(x)=xforsomepositiveintegerNk.Deﬁnegk(x1,...,xk)=fkuNk−1k(x1),...,uNk−1k(xk)andnotethatgkisidempotentforeachk.Finally,hg1,...,gki(Oisanidempotentsemi-roundpropersubcloneofOwhichisalsoacounterexampletotheconjecture,contra-dictingtheminimalityofO.Ifthereisaunaryoperationu(x)thatisnotinjective,thenu2N(x)=uN(x)forsomepositiveintegerN,anduN(a)=uN(b)=aforsomedistincta,b∈D.Deﬁnegk(x1,...,xk)=uN fk uN(x1),...,uN(xk)andnotethateachgkisasymmetricoperationthatactsonaandbthesameway;thatis,replacingawithborbwithaingk’sinputwillnotchangeitsoutput.Sincehg1,...,gki⊆OiseﬀectivelyacloneonthedomainD\{b},Oisnotminimalwithrespecttodomainsize,contradiction.
Proposition11.Letfbeasemilatticeoperation.Thenhfiisround.Proof.Considertheposetdeterminedbyf.Byinduction,thesymmetrick-aryoperationfkthatreturnsthegreatestupperboundofitskinputsisintheclone.Indeed,f1=π11andf2=f.Nowfk+1(x1,...,xk+1)=f2(fk(x1,...,xk),xk+1)byinduction.
Theorem12.LetObeanidempotentroundcloneoveradomainD,andsupposesomebinaryoperationinOactslikeaheight-1semilatticeoversomesubsetD0⊆D.ThenOcontainsaroundsubclonethatactslikeaheight-1semilatticeoverD0.Proof.Lett2bethegivenbinaryoperation,andforeachk∈Z+,letfkdenoteak-arysymmetricoperationinO.Thenthebinaryoperationg2(x,y)=f2(t2(x,y),t2(y,x))issymmetricandactslikeaheight-1semilatticeoverD0.Byinduction,thek-aryoperationgk(x1,...,xk)=fkc2
gk−1(x1,...,xk)issymmetricandactslikeaheight-1semilatticeoverD0.(Recallthatcgk−1isthefunctionthattakesak-tupleandreturnsthek-tuplewhoseithentryisgk−1appliedonthek−1variablesnotintheithcoordinate,forall1≤i≤k.)Thenhg1,g2,...iisaroundsubcloneofOthatactslikeaheight-1semilatticeoverD0,asdesired.
Corollary13.LetObeanidempotentsemi-roundcloneoveradomainD,andsupposesomebinaryoperationinOactslikeaheight-1semilatticeoversomesubsetD0⊆D.ThenOcontainsasemi-roundsubclonethatactslikeaheight-1semilatticeoverD0.Proof.ThisistheessentiallythesameastheproofofTheorem12,excepttheinductionstopsaftercon-structingthe|D|-arysymmetricoperation.
10
Theorem14.LetObeanidempotentcloneoveradomainDwithabinarysymmetricoperationf2thatactslinearlyonasubsetD0⊆Dwithoddprimecardinality,andsupposeOcontainsa|D0|-arysymmetricoperationf|D0|.Thenonecanﬁndsymmetricoperationsg1,g2,...∈Oandaconstantcsuchthatgk(x1,...,xk)=(xif{x1,x2,...,xk}={x}cotherwiseforallk∈Z+andtuples(x1,x2,...,xk)∈(D0)k.Proof.Weprovethisbyinductiononk.Deﬁnec:=f|D0|(d1,...,d|D0|)whereD0={d1,...,d|D0|},andfork=1noteg1=π11.Forthek=2construction,assumeD0={0,1,...,p−1}forsomeoddprimep.Sincef2(x,y)=x+y
2(modp)overD0,onecanconstructanyoperationthatactslikea
2bx+2b−a
2by(modp)overD0bycomposingf2withitselfforanya,b∈Z+with0≤a≤2b;inparticular,bychoosingb=p−1,theoperationax−(a−1)y(modp)canbeconstructedusingonlyf2foranya∈Z,byFermat’sLittleTheorem.Finally,deﬁneg2(x,y):=fk(x,2x−y,3x−2y,...,px−(p−1)y)(modp);thisworksbecause{x,2x−y,3x−2y,...,px−(p−1)y}={0,1,...,p−1}(modp)wheneverx−y6=0,bytheprimalityofp.Toconstructgk+1givengk,deﬁnegk+1(x1,x2,...,xk+1):=g2(gk(x1,x2,...,xk),gk(f2(x1,xk+1),f2(x2,xk+1),...,f2(xk,xk+1))).Thisworksbycaseworkonwhetherornotx1=x2=...=xk;iftheyarenotallequal,thenbothargumentsofg2arec,andiftheyareallequal,thenbothargumentsofg2aredistinctelementsofD0.
Lemma15.Supposeasemi-roundcloneOoveradomainDwith|D|≤4containsabinarysymmetricoperationthatactslinearlyonathree-elementsetD0⊆D.ThenOisnotminimallysemi-round.Proof.LetcbetheconstantguaranteedbyTheorem14.Ifc∈D0thenwearedonebyCorollary13.ThuswecanassumeD={0,1,2,3},D0={0,1,2},andc=3.ThebinarysymmetricoperationthatactslinearlyoverD0mustbeoftheformf21a0bc
0123
0
021a1
210b2
102c3
abc3andg2mustbeoftheformf33d3ef
0123
0
033d1
313e2
332f3
def3.11
Ifatleasttwoofd,e,fareequalto3,thenf33d3efactslikeaheight-1semilatticeoverasubsetD0⊂Dwithsize3,sowecanapplyCorollary13toshowthatOisnotminimallysemi-round.ThesetD0contains3andthecorrespondingdomainelementsforthetwoofd,e,fthatareequalto3.Ifthereisatmostone3amongd,e,andf,wecanassumethereisatmostone3amonga,b,andcbyCorollary13onatwo-elementsubsetofD;otherwiseOisnotminimallysemi-round.Ifa∈{1,2}thenonecancheckthata a∈SgO1
2,2
1.Thus,somebinaryoperationactslikeaheight-1semilatticeover{1,2},soOisnotminimallysemi-round.Similarly,ifb∈{0,2}orc∈{0,1}thenOisnotminimallysemi-round.Ifexactlyoneofa,bandcisequalto3,thenwithoutlossofgeneralityassumec=3,sothebinaryoperationthatactslinearlyoverD0isforcedtobef210013:f210013
0123
0
02101
21012
10233
0133.Sincewecanassumea/∈{1,2}andb/∈{0,2},byCorollary13eitherOisnotminimallysemi-round,orthebinarysymmetricoperationf330313guaranteedbyTheorem14isinO:f330313
0123
0
03301
31312
33233
0133.Thenonecancheckthat1
1∈SgO1
2,2
1soOisnotminimallysemi-roundbyCorollary13.If3/∈{a,b,c},thenthebinaryoperationthatactslinearlyoverD0isforcedtobef210012:f210012
0123
0
02101
21012
10223
0123.ByCorollary13eitherOisnotminimallysemi-round,orthebinarysymmetricoperationf330312guaranteedbyTheorem14isinO:f330312
0123
0
03301
31312
33223
0123.Thenf330312(f330312(x,y),f330312(f330312(x,z),f330312(y,z)))isaternarysymmetricoperation,sohf330312i(Oisasemi-roundclonebyRemark6,andonecancheckthatitdoesnotcontainasymmetricbinaryoperationotherthanf330312.
12
5Classiﬁcation
Inthissection,wepresentacompletecatalogueofeveryminimalidempotentsemi-roundcloneoveradomainofsizeatmost4,uptoarenamingofthedomainelements.Additionally,weprovethateverysemi-roundcloneoveradomainofsize4isalsoround.Non-idempotentclonesarenotconsidered,asTheorem10statesthataminimalcounterexampletotheconjectureisidempotent,andnon-minimalroundclonesarenotconsidered,asCorollary9guaranteestheexistenceofaminimallysemi-roundsubcloneofeverysemi-roundclone.
5.1DomainofSize1
Thereisonlyonecloneoveradomainofsize1,whichisbothsemi-roundandround.5.2DomainofSize2
Thereisonlyoneminimalsemi-roundcloneoveradomainofsize2,uptoarenamingofthedomainelements.Toprovethis,letD={0,1}andletf2(x,y)betheclone’sbinarysymmetricoperation.Wecan,withoutlossofgenerality,assumethatf2(x,y)=xy:f2
01
0
001
01.Thenhf2iisroundbyProposition11becausef2isasemilatticeoperation,soitistheuniqueminimalsemi-roundcloneover{0,1},uptoarenamingofthedomainelements.5.3DomainofSize3
AssumeD={−1,0,1},whichweabbreviateas{−,0,+}.Byaﬁnitecasecheck,anyminimalclonemustcontain,uptoarenamingofthedomainelements,oneoffollowingsymmetricbinaryoperations:f000
−0+
−
−000
000+
00+f++0
−0+
−
−0+0
00++
+++f+00
−0+
−
−000
00++
0++f00+
−0+
−
−+00
+00+
00+f+0−
−0+
−
−−00
−0++
0++f+−0
−0+
−
−0−0
00++
−++f−0+
−0+
−
−+00
+0−+
0−+.Todeterminetheminimalroundclones,wecaseworkonthesymmetricbinaryoperation.•hf000iisminimallyroundbyProposition11,sincef000isasemilatticeoperation.•hf++0iisminimallyroundbyProposition11,sincef++0isasemilatticeoperation.13
•hf++0i⊆hf+00i,sof+00doesnotneedtobeconsidered,sincef+00(f+00(x,f+00(x,y)),f+00(y,f+00(x,y)))=f++0(x,y).•hf++0i⊆hf00+i,sof00+doesnotneedtobeconsidered,sincef00+(f00+(x,f00+(x,y)),f00+(y,f00+(x,y)))=f++0(x,y).•hf+0−iisminimallysemi-roundandminimallyround,sincealloperationsoftheformfk(x1,...,xk)=sgn(x1+...+xk)areinhf−+0i.Fork=1takef1=π11,andfork=2takef2=f−+0.Toconstructfk+1inductively,theidentityfk+1(x1,...,xk+1)=f2(fk(fk(x2,x3,...,xk,xk+1),fk(x1,x3,...,xk,xk+1),fk(x1,x2,...,xk,xk+1),...,fk(x1,x2,x3,...,xk+1)),fk(x1,...,xk))willsuﬃce.•hf+−0iisnotsemi-roundbecausetheexistenceofasymmetricternaryoperationisforbiddenbytheautomorphismsending0to+,+to−,and−to0.Hence,wewilldetermineallminimalsemi-roundclonesoftheformhf+−0,gi,wheregissomesymmetricidempotentternaryoperation.ByCorollary13theremustbeasymmetricternaryoperationinOoftheformf3(x,y,z)=(f+−0(d1,d2)if∃(d1,d2∈D){x,y,z}={d1,d2}cif{x,y,z}=Dforsomec∈D;withoutlossofgeneralityassumec=0.Now,hf+−0,f3iisroundbecauseeachoperationoftheformfk(x1,...,xk):=(f+−0(d1,d2)if∃(d1,d2∈D){x1,...,xk}={d1,d2}0otherwiseisinhf+−0,f3i.Notethatthebasecasesk∈{1,2,3}aretrue,soitsuﬃcestoconstructfkinductively.Deﬁnex=(x1,...,xk+1),anddeﬁney=cN
fk(x),y0=cN
fk(f2(x,y)),andy00=cN
fk(f2(x,y0))forasuﬃcientlylargeintegerN.Thenfk+1(x)=f3(y,y0,y00)1,wherethesubscriptdenotestheﬁrstelementofthetuple;thisworksbecausey=cN
fk(x)isaconstanttupleforsuﬃcientlylargeN,so{y,y0,y00}={(−,...,−),(0,...,0),(+,...,+)}ory=y0=y00.•Noclonecontainingf−0+isminimallysemi-roundbyLemma15,sothiscasedoesnotneedtobeconsidered.14
5.4DomainofSize4
Thecharacterizationofallminimalidempotentsemi-roundclonesoveradomainofsize4willbedonethroughcaseworkonthesymmetricbinaryoperationitcontains.Thereare4(4
2)=4096suchoperationsbuttheyfallinto192distinctequivalenceclassesunderisomorphism;representativeelementsarelistedbelow.f000000,f000001,f000002,f000011,f000012,f000013,f000021,f000023,f000033,f000111,f000112,f000121,f000122,f000123,f000132,f000321,f001000,f001001,f001002,f001003,f001010,f001011,f001012,f001013,f001020,f001021,f001022,f001023,f001030,f001031,f001032,f001033,f001100,f001101,f001102,f001103,f001110,f001111,f001112,f001113,f001120,f001121,f001122,f001123,f001130,f001131,f001132,f001133,f001200,f001201,f001202,f001203,f001210,f001211,f001212,f001213,f001220,f001221,f001222,f001223,f001230,f001231,f001232,f001233,f001300,f001301,f001302,f001303,f001310,f001311,f001312,f001313,f001320,f001321,f001322,f001323,f001330,f001331,f001332,f003000,f003001,f003002,f003011,f003012,f003013,f003021,f003100,f003101,f003102,f003110,f003112,f003113,f003120,f003121,f003122,f003300,f003301,f003302,f003311,f003312,f003321,f011000,f011001,f011002,f011010,f011011,f011012,f011013,f011020,f011021,f011022,f011023,f011030,f011031,f011032,f011120,f011121,f011122,f011123,f011130,f011131,f011132,f011220,f011223,f011230,f011231,f011320,f011321,f011322,f012000,f012001,f012002,f012003,f012010,f012011,f012012,f012013,f012020,f012021,f012022,f012030,f012031,f012032,f012100,f012101,f012102,f012103,f012120,f012121,f012122,f012130,f012131,f012132,f012200,f012203,f012210,f012213,f012230,f012300,f012301,f012302,f012310,f012311,f012313,f012320,f012321,f012330,f013002,f013010,f013011,f013012,f013021,f013022,f013102,f013310,f013321,f022101,f022301,f022321,f023321,f032000,f032001,f032020,f032021,f032030,f032230,f032320,f032321,f211000,f211020,f211300,f211301.Manyoftheseoperationsdonotneedtobeconsideredbecausetheclonestheygeneratecontainotherbinaryoperations.Forexample,onlyoperationsintheimageofrepeatedcompositionofthemapf(a,b)7→f(f(a,f(a,b)),f(b,f(a,b)))needtobeconsidered,whichresultsinthefollowing37operations:f000000,f000002,f000012,f000013,f000033,f000111,f000112,f000123,f000132,f000321,f001030,f001031,f001032,f001033,f001130,f001132,f001133,f001230,f001231,f001232,f001233,f003012,f003013,f003112,f003113,f003312,f003321,f011231,f011321,f013310,f022101,f023321,f032230,f032320,f032321,f211000,f211020.Underthemapf(a,b)7→f(f(a,f(a,f(a,b))),f(b,f(b,f(a,b))))theoperationsf001030,f001130,f001230,f011321,f013310canberemovedfromthelist.Similarly,underthemapf(a,b)7→f(f(a,f(b,f(a,b))),f(b,f(a,f(a,b))))theoperationsf211000andf022101canberemoved,andf211020canberemovedbyconsideringthemapf(a,b)7→f(a,f(a,f(b,f(b,f(a,b))))).15
Lastly,theoperationsf000321,f003321,f023321,f032230,f032320,andf032321allactlinearlyoverathree-elementsubsetoftheirdomain,sobyLemma15theydon’tneedtobeconsidered.Therefore,onlythefollowing23operationsneedtobeconsideredforanalysis:f000000,f000002,f000012,f000013,f000033,f000111,f000112,f000123,f000132,f001031,f001032,f001033,f001132,f001133,f001231,f001232,f001233,f003012,f003013,f003112,f003113,f003312,f011231.Sixteenoftheseoperationsalreadygenerateminimalroundclones.ThefollowingﬁveoperationsgenerateroundclonesbyProposition11becausetheyaresemilattices:f000000
0123
0
00001
01002
00203
0003f000002
0123
0
00001
01002
00223
0023f000012
0123
0
00001
01012
00223
0123f000111
0123
0
00001
01112
01213
0113f000112
0123
0
00001
01112
01223
0123.Thefollowingeightoperationsalsogenerateroundclones:f000013
0123
0
00001
01012
00233
0133f000033
0123
0
00001
01032
00233
0333f000123
0123
0
00001
01122
01233
0233f001031
0123
0
00011
01032
00213
1313f001133
0123
0
00011
01132
01233
1333f001231
0123
0
00011
01232
02213
1313f001233
0123
0
00011
01232
02233
1333f011231
0123
0
00111
01232
12213
1313.16
Foreachoftheseoperationsf,the(k+1)-aryoperationfk+1(x1,...,xk+1):=f2(fk(fk(x2,x3,...,xk,xk+1),fk(x1,x3,...,xk,xk+1),fk(x1,x2,...,xk,xk+1),...,fk(x1,x2,x3,...,xk+1)),fk(x1,...,xk)),wheref2=f,issymmetric;thisissimilartothef+0−casefromthedomainofsize3enumeration.Lastly,thefollowingthreeoperationsgenerateroundclones:f001032
0123
0
00011
01032
00223
1323f001033
0123
0
00011
01032
00233
1333f001232
0123
0
00011
01232
02223
1323.Foreachoftheseoperations,a(k+1)-arysymmetricoperationfk+1canbeconstructedthroughthefollowinginduction,where(y1,...,yk+1)=cN
fk(x1,...,xk+1)forasuﬃcientlylargeintegerN.fk+1(x1,...,xk+1):=f2(fk(fk(y2,y3,...,yk,yk+1),fk(y1,y3,...,yk,yk+1),fk(y1,y2,...,yk,yk+1),...,fk(y1,y2,y3,...,yk+1)),fk(y1,...,yk)).Thisworksbecause(y1,...,yk+1)alwaysliesin(D0)k+1forsomethree-elementsubsetD0(DforsuﬃcientlylargeN:•Forhf001032i,repeatedlyapplyingcfktothetuple(x1,...,xk+1)willalwaysresultinanelementof{0,1,3}k+1,unless(x1,...,xk+1)∈{2,3}k+1.•Forhf001033i,repeatedlyapplyingcfktothetuple(x1,...,xk+1)willalwaysresultinanelementof{0,1,3}k+1,unless(x1,...,xk+1)=(2,...,2).•Forhf001232i,repeatedlyapplyingcfktothetuple(x1,...,xk+1)willalwaysresultinanelementof{0,1,2}k+1,unless(x1,...,xk+1)∈{1,3}k+1.Thefollowingﬁveoperationsdonotgeneratecloneswithasymmetricternaryoperation,buttheygenerateroundcloneswhensymmetricternaryoperationsareadded:17
f000132
0123
0
00001
01132
01223
0323f003012
0123
0
00031
01012
00223
3123f003013
0123
0
00031
01012
00233
3133f003112
0123
0
00031
01112
01223
3123f003113
0123
0
00031
01112
01233
3133.Wecaseworkoneachone.Fortheremainderofthisparagraph,letf2bethebinarysymmetricoperationandletf3betheternarysymmetricoperation.Theinductionusedtoprovethateachcaseyieldsaroundcloneissimilartothef+−0casefromthedomainofsize3enumeration.Toconstructthek-aryoperationfk+1,deﬁnex:=(x1,...,xk+1),anddeﬁney:=cN
fk(x),y0:=cN
fk(f2(x,y)),andy00:=cN
fk(f2(x,y0))forasuﬃcientlylargeintegerN.Thenfk+1(x):=f3(y,y0,y00)1,wherethesubscriptdenotestheﬁrstelementofthetuple,issymmetric;thisworksforeachcasebe-causey=cN
k+1(x)isaconstanttupleforsuﬃcientlylargeN,soeithery=y0=y00or{y,y0,y00}={(d1,...,d1),(d2,...,d2),(d3,...,d3)},where{d1,d2,d3}ischosensuchthatf2,whenrestrictedtothedo-main{d1,d2,d3}⊂D,canberenamedtof+−0.Thefunctionθthattakesabinaryoperationf2andaternaryoperationf3asinputandoutputsaternaryoperationisdeﬁnedasθ(f2,f3)(x,y,z):=cN
f2 f2 cN
f2(x,y,z),f3 cN
f2(x,y,z)1forasuﬃcientlylargepositiveintegerN.Ineachofthefollowingcases,θreturnsasymmetricternaryoperationthatmodiﬁesonlyoneortwooutputsoff3.Thecaseworkonthebinarysymmetricoperationisbelow.•Supposeacloneisgeneratedbyf000132andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc(x,y,z)=













00∈{x,y,z}1{x,y,z}∈{{1},{1,2}}2{x,y,z}∈{{2},{2,3}}3{x,y,z}∈{{3},{1,3}}c{x,y,z}={1,2,3}forsomec∈D.Sinceθ(f000132,gc)mapsg17→g37→g27→g1,thiscasegivestwodistinctminimalroundclones.18
•Supposeacloneisgeneratedbyf003012andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc,d(x,y,z)=





















0{x,y,z}∈{{0},{0,1},{0,2},{1,2},{0,1,2},{1,2,3}}1{x,y,z}∈{{1},{1,3}}2{x,y,z}∈{{2},{2,3}}3{x,y,z}∈{{3},{0,3}}c{x,y,z}={0,1,3}d{x,y,z}={0,2,3}forsomec,d∈D.Sinceθ(f003012,gc)mapsg0,07→g3,37→g1,27→g0,0g0,27→g3,07→g1,37→g0,2g0,37→g3,27→g1,07→g0,3andmapseachofg0,1,g1,1,g2,0,g2,1,g2,2,g2,3,g3,1tooneoftheabovethreecycles,thiscasegivesthreedistinctminimalroundclones.•Supposeacloneisgeneratedbyf003013andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc(x,y,z)=













0{x,y,z}∈{{0},{0,1},{0,2},{1,2},{0,1,2}}1{x,y,z}∈{{1},{1,3}}2{x,y,z}={2}3{x,y,z}∈{{3},{0,3},{2,3},{0,2,3}}c{x,y,z}∈{{0,1,3},{1,2,3}}forsomec∈D.Sinceθ(f003013,gc)mapsg07→g37→g17→g0andmapsg2totheabovecycle,thiscaseonlygivesoneminimalroundclone.•Supposeacloneisgeneratedbyf003112andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc,d(x,y,z)=

















0{x,y,z}∈{{0},{0,1},{0,2},{0,1,2}}1{x,y,z}∈{{1},{1,2},{1,3},{1,2,3}}2{x,y,z}∈{{2},{2,3}}3{x,y,z}∈{{3},{0,3}}c{x,y,z}={0,1,3}d{x,y,z}={0,2,3}forsomec,d∈D.Sinceθ(f003012,gc)mapsg0,07→g3,37→g1,27→g0,0g0,27→g3,07→g1,37→g0,2g0,37→g3,27→g1,07→g0,3andmapseachofg0,1,g1,1,g2,0,g2,1,g2,2,g2,3,g3,1tooneoftheabovethreecycles,thiscasegivesthreedistinctminimalroundclones.19
•Supposeacloneisgeneratedbyf003113andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc(x,y,z)=













0{x,y,z}∈{{0},{0,1},{0,2},{0,1,2}}1{x,y,z}∈{{1},{1,2},{1,3},{1,2,3}}2{x,y,z}={2}3{x,y,z}∈{{3},{0,3},{2,3},{0,2,3}}c{x,y,z}={0,1,3}forsomec∈D.Sinceθ(f003013,gc)mapsg07→g37→g17→g0andg2totheabovecycle,thiscasegivestwodistinctminimalroundclones.Onecanprovethateachoftheclonesenumeratedabovearedistinctbycomputingrelations;foreachpairofdistinctclonesO1andO2intheabovelist,onecanﬁndarelationthatispreservedbyO1butnotbyO2.Thefollowingtwooperationsalsodon’tgeneratecloneswithasymmetricternaryoperation,buttheygenerateroundcloneswhensymmetricternaryoperationsareadded:f001132
0123
0
00011
01132
01223
1323f003312
0123
0
00031
01312
03223
3123.Fortheremainderofthisparagraph,letf2bethebinarysymmetricoperationandletf3betheternarysymmetricoperation.Toconstructthek-aryoperationfk+1foreachcase,deﬁne:x:=(x1,...,xk+1)y:=cN
fkf2cN
fk(x),cN+1fk(x)y0:=cN
fkf2cN
fk(f2(x,y)),cN+1fk(f2(x,y))y00:=cN
fkf2cN
fk(f2(x,y0)),cN+1fk(f2(x,y0))forasuﬃcientlylargeintegerN.Thenfk+1(x):=f3(y,y0,y00)1,wherethesubscriptdenotestheﬁrstelementofthetuple,issymmetric;thisworksforeachcasebecausey=cN
fkf2cN
fk(x),cN+1fk(x)isaconstanttupleforsuﬃcientlylargeN,soeithery=y0=y00or{y,y0,y00}={(d1,...,d1),(d2,...,d2),(d3,...,d3)},wheref2restrictedtothedomain{d1,d2,d3}⊂Dcanberenamedtof+−0.ThefunctionΘthattakesabinaryoperationf2andaternaryoperationf3asinputandoutputsaternaryoperationisdeﬁnedasΘ(f2,f3)(x,y,z):=cN
f2f2f2(cN
f2(x,y,z),cN+1f2(x,y,z)),f3f2(cN
f2(x,y,z),cN+1f2(x,y,z))1forasuﬃcientlylargepositiveintegerN.Ineachofthefollowingcases,Θreturnsasymmetricternaryoperationthatmodiﬁesonlyoneortwooutputsoff3.20
•Supposeacloneisgeneratedbyf001132andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc(x,y,z)=













0{x,y,z}∈{{0,2},{0,1,2},{0,2,3}}1{x,y,z}={1,2}2{x,y,z}∈{{2},{2,3}}σ−1(sgn(σ(x)+σ(y)+σ(z))){x,y,z}⊆{0,1,3}c{x,y,z}={1,2,3}forsomec∈D,whereσ(0)=−1,σ(1)=0,andσ(3)=1.Thiscaseonlygivesoneminimalroundclone,sinceΘ(f001132,gc)mapsg17→g37→g27→g1.Toprovethathf001132,g0i⊆hf001132,g1i,letx:=(x,y,z)andy:=f2cN
f001132(x),cN+1f001132(x),whichwillbeconstantunless{x,y,z}={1,2,3}.Additionally,lety0:=cf001132(y)andy00:=cf001132(y0)bethecyclicrotationsofy.ThencN
f003312(f2(f2(f2((g0(y),g0(y),g0(y)),y),y0),y00))1=g1(x)wherethesubscriptdenotestheﬁrstelementofthetuple,asdesired.•Supposeacloneisgeneratedbyf003312andasymmetricternaryoperation.UsingCorollary13,onecanforcetheexistenceofasymmetricternaryoperationoftheformgc,d(x,y,z)=













0{x,y,z}∈{{0},{0,1},{0,2}}3{x,y,z}∈{{0,3},{0,1,2}}σ−1(sgn(σ(x)+σ(y)+σ(z))){x,y,z}⊆{1,2,3}c{x,y,z}={0,1,3}d{x,y,z}={0,2,3}forsomec,d∈D,whereσ(1)=−1,σ(2)=1,andσ(3)=0.Thiscasegivesthreeminimalroundclones,sinceΘ(f003312,gc,d)mapsg0,07→g3,37→g1,27→g0,0g0,27→g3,07→g1,37→g0,2g0,37→g3,27→g1,07→g0,3andalsoeventuallymapsallotherstotheabovecycles,sinceg1,17→g0,17→g3,0,g2,27→g2,07→g0,3,g2,17→g1,2,g2,37→g3,2,andg3,17→g1,3.Sincewehaveexhaustedallcases,wehaveestablishedtheproofofTheorem2.5.5DomainofSize5
Withcomputerassistance,ithasbeenshownthateveryidempotentcloneoveradomainofsize5thatcontainssymmetricoperationsofarities1,2,3,4,and5containssymmetricoperationsofaritiesupto20.ThecodeusedtoverifythisisavailableonGithubathttps://github.com/The-Turtle/PRIMES.6FutureWork
Ifwewanttomakeprogressonlargerdomains,weneedawaytodeterminewhetherornotaclonehasasymmetricoperationofagivenaritywithoutexplicitlygeneratingone.21
Deﬁnition16.LetAbeanalgebrawithunderlyingsetA.Foranytuplea=(a1,...,ak)∈Ak,deﬁnethesymmetricrelationonatobethesetSym(a):=SgAk!



aσ1(1)...aσk!(1)

,...,

aσ1(k)...aσk!(k)



,whereσ1,σ2,...,σk!arethek!permutationsofthetuple(1,...,k).Proposition17.LetAbeanalgebrawithunderlyingsetA,andsupposethatforeveryj≤kandeverytuplea∈Aj,thesymmetricrelationSym(a)containsaconstanttuple.ThenAhasasymmetricoperationofeveryaritylessthanorequaltok.Proof.Weprovethisbyinductiononk;forthebasecasek=1,takef1=π11.Bytheinductivehypothesis,therearesymmetricoperationsf1,f2,...,fk−1ofeveryaritystrictlylessthank.Nowsupposethatfisak-aryoperationsuchthatthesetT⊆Akoftuplesforwhichfactssymmetricallyonismaximal.WeclaimthatTmustequalAk;toprovethis,itsuﬃcestoshowthatift∈Ak\Tisatuplewhichfdoesnotactsymmetricallyon,thenthereisak-aryoperationgwhichactssymmetricallyonT∪{t}.Wewillﬁrstconstruct,foreachj<k,anoperationgjwhichactssymmetricallyonTandwhichisunchangedbyeverypermutationofitsﬁrstjvariables.Westartbytakingg1=f,andthenweinductivelydeﬁnegjasgj(x1,...,xk):=fj(gj−1(x1,x2,...,xj−1,xj,xj+1,...,xk),gj−1(x2,x3,...,xj,x1,xj+1,...,xk),...,gj−1(xj,x1,...,xj−2,xj−1,xj+1,...,xk)).Finally,letabethetuplea:=(gk−1(t1,...,tk),gk−1(t2,...,tk,t1),...,gk−1(tk,t1,...,tk−1)).Byassumption,Sym(a)containsaconstanttuple,sotheremustbesomek-aryoperationh∈Clo(A)whichactssymmetricallyona.Thenwedeﬁnegbyg(x1,...,xk):=h(gk−1(x1,...,xk),gk−1(x2,...,xk,x1),...,gk−1(xk,x1,...,xk−1)).
TherelationSym(a)hasausefulspecialproperty.Proposition18.LetAbeanalgebrawithunderlyingsetA.Foranytuplea∈Akandanypairofpermutations(σ,τ)on{1,...,k},letP≤A2bethebinaryrelationπiσ,iτ(Sym(a)),whereiσandiτaretheindicesofσandτasdeﬁnedinDeﬁnition16.ThenforanysubsetB⊆A,wehaveB+P=B=⇒B−P=B.Proof.DeﬁneP◦n:=P◦···◦P|
{z
}nP’s.ThenwehaveP−⊆P◦(k!−1),sinceP◦(k!−1)containsthegeneratorsofP−,soB−P⊆B+P◦(k!−1)=B.Similarly,B=B+P⊆B−P◦k!−1⊆B.HenceB−PmustinfactequalB.
Deﬁnition19.LetAbeanalgebrawithunderlyingsetA.SaythatarelationR≤Amisreversibleifitsatisﬁesthefollowingtwoproperties:22
•foralli,jwehaveπi(R)=πj(R),and•foreverysequencep=((i1,j1),...,(ik,jk))ofpairsofcoordinatesofR,ifwedeﬁnethebinaryrelationPp≤A2byPp:=πi1,j1(R)◦···◦πik,jk(R),thenforeveryB⊆A,wehaveB+Pp=B=⇒B−Pp=B.Proposition20.ForeveryalgebraAwithunderlyingsetAandeverytuplea∈An,therelationSym(a)isreversible.
Proof.SincethemarginaldistributionsofeachcoordinateoftheuniformdistributiononthesetoftuplesinSym(a)areequal,thisfollowsfromtheimplication(e)=⇒(a)ofProposition23below.
Wehavethefollowingstrongreﬁnementofourmainconjecture.Conjecture21.SupposethatAisaﬁniteidempotentalgebra,suchthatforeverysubquotientB∈HS(A)thereissomeelementb∈BwhichisﬁxedbyeveryautomorphismofB.TheneveryreversiblerelationR≤Ancontainsaconstanttuple.Theconditioninvolvingarbitrarycompositionsoftwo-variableprojectionsoftherelationRinthedeﬁ-nitionofreversibilityisnecessary,asdemonstratedbythefollowingexample.Example22.LetA=({−,0,+},sgn(x+y))andletR≤A5betherelationR:=(x1,x2,x3,x4,x5)∈A5|x1+x2+x3≥1∧x4=−x5 .TheneverybinaryprojectionofRisreversible,buttherelationRisnotreversible:wehave{−}+π1,2(R)+π4,5(R)={−}but{−}−π4,5(R)−π1,2(R)={−,0,+}.SinceRdoesnotcontainanyoftheconstanttuples(−,...,−),(0,...,0),or(+,...,+),weneedthestrongerconditionaboutarbitrarycompositionsoftwo-variableprojections.Forbinaryrelations,theconceptofreversibilitysimpliﬁes.Proposition23.IfR≤sdA2isabinarysubdirectrelationonaﬁnitealgebraAwithunderlyingsetA,thenthefollowingareequivalent.(a)ForeveryB⊆A,wehaveB+R=B=⇒B−R=B.(b)IfweconsidertheorderedpairsofRastheedgesofadirectedgraphGwithvertexsetA,theneveryweaklyconnectedcomponentofGisalsostronglyconnected.(c)IfweconsidertheorderedpairsofRastheedgesofadirectedgraphGwithvertexsetA,theneverydirectededgeofGiscontainedinadirectedcycleofG.(d)Thereissomen≥1suchthatR−⊆R◦n.(e)ThereisapositiveprobabilitydistributionwithsupportRsuchthatthemarginaldistributionsontheﬁrstandsecondcoordinatesagree.(f)ThebinaryrelationRisreversible;thatis,everybinaryrelationwhichcanbewrittenasacompositionofcopiesofRandR−satisﬁes(a).23
Proof.(a)=⇒(b):deﬁneaquasiorderonAbyabifthereisanyk≥0suchthat(a,b)∈R◦k.Foranya∈A,thereisa-maximalelementb∈Asuchthatab,bytheﬁnitenessofA.LetBbethesetofallb0suchthatbb0,thenthe-maximalityofbimpliesthatBisastronglyconnectedcomponentofRandthatB+R=B.Then(a)impliesthatwehaveB−R=B,sowehavea∈{b}−R◦k⊆B−R◦k=B,andsimilarlyanyelementintheweaklyconnectedcomponentcontainingaisalsocontainedinB.(b)=⇒(c)isobvious.For(c)=⇒(d),pickforeachdirectededgeofRadirectedcyclecontainingit,andchoosensuchthatn+1isacommonmultipleofthelengthsofallofthesedirectedcycles.(d)=⇒(a)and(f)=⇒(a)arealsoobvious.Toprovethat(c)=⇒(e),ﬁndacollectionCofdirectedcyclesofRthatcontainseveryedgeofRatleastonce.DeﬁneaprobabilitydistributiononRbythefollowingtwostepprocess:ﬁrstpickauniformlyrandomcycleC∈C,thenpickauniformlyrandomedge(x,y)∈C.For(e)=⇒(a),letp(a,b)>0betheprobabilityassignedtoagivenelement(a,b)∈R(andsetp(a,b)=0for(a,b)6∈R),andletpb:=Xa∈Ap(a,b)=Xc∈Ap(b,c)bethemarginalprobabilityofseeingboneithertheﬁrstorsecondcoordinate.ForanysubsetB⊆A,deﬁnep(B)byp(B):=Xb∈Bpb.Thenwehavep(B+R)=Xb∈B+Rpb=Xb∈B+RXa∈{b}−Rp(a,b)≥Xb∈B+RXa∈Bp(a,b)=Xa∈Bpa=p(B),withequalityonlywheneveryelementb∈B+Rhas{b}−R⊆B.IfB+R=B,thenwemusthaveequalityabove,soB−R=B+R−R=B.Giventheequivalencebetween(a)and(e),(e)=⇒(f)followsfromthefactthatifRandSareanypairofbinaryrelationssuchthattherearepositiveprobabilitydistributionspandqsupportedonRandS,respectively,suchthatthemarginalofponthesecondcoordinateequalsthemarginalofqontheﬁrstcoordinate,thenthereisapositiveprobabilitydistribution“p◦q”supportedonR◦Ssuchthatthemarginalsofpandp◦qontheﬁrstcoordinateareequal,andthemarginalsofp◦qandqonthesecondcoordinateareequal.Theequivalenceof(f)canalsobeshownbyproving(d)=⇒(f).LetR0beacompositionoficopiesofRandjcopiesofR−,insomeorder;itsuﬃcestoshowthatR0satisﬁes(d).Ifi>j,thenR0⊇R◦(i−j),soifR−⊆R◦n,then(R−)◦(i−j)⊆R◦n(i−j)⊆R0◦n,andwecanﬁnishsinceR0andR0−areeachcontainedinsomecompositionofR◦(i−j)and(R−)◦(i−j).Thecasei<jissimilar,soweareleftwiththecasei=j.Todealwiththecasei=j,thecasewhereR0isacompositionofanequalnumberofcopiesofRandR−insomeorder,codethesequenceofcopiesofRandR−asasequenceoficopiesof+andicopiesof−.Letaand−bbethelargestvalueandsmallestvalue,respectively,ofthepartialsumsofthesequenceof+’sand−’s.Thenit’seasytoseethatR0containstherelationsR±a:=R◦a◦(R−)◦aandR∓b:=(R−)◦b◦R◦b.Thus,bothR0and(R0)−arecontainedinsomecompositionofcopiesofR±aandR∓b,asdesired.
24
Theorem24.Conjecture21istrueforthealgebraA=({−,0,+},sgn(x+y)).Proof.LetAhaveunderlyingsetA,andletR≤Anbeareversiblerelation.Ifπi(R)6=A,thenπi(R)isasemilattice-weleavethiscasetothereader.Weareleftwiththecaseπi(R)=Aforalli;thatis,thecasewhereRissubdirect.AbruteforceenumerationshowsthateverybinarysubdirectrelationonAisoneofthesevenrelations{(x,y)∈A2|x=y},{(x,y)∈A2|x=−y},{(x,y)∈A2|x≤y},{(x,y)∈A2|x≥y},{(x,y)∈A2|x+y≥0},{(x,y)∈A2|x+y≤0},A2.Inparticular,eachbinarysubdirectrelationS≤sdA2iscompletelydeterminedbytheintersectionS∩{−,+}2;infact,thecompositionofanypairofbinarysubdirectrelationsonAisalsodeterminedbythecompositionoftheirrestrictionsto{−,+}.Amongthesesevenrelations,thetwobinaryrelations{(x,y)∈A2|x≤y}and{(x,y)∈A2|x≥y}arenotreversible.Since{(x,y)∈A2|x=−y}◦{(y,z)∈A2|y+z≥0}={(x,z)∈A2|x≤z}and{(x,y)∈A2|x+y≤0}◦{(y,z)∈A2|y+z≥0}={(x,z)|x≤z},weseethateveryreversiblesubdirectarity-krelationReither(a)hasπi,j(R)∈{(x,y)∈A2|x=y},{(x,y)∈A2|x=−y},A2 forallintegers1≤i,j≤k,(b)hasπi,j(R)∈{(x,y)∈A2|x=y},{(x,y)∈A2|x+y≥0},A2 forallintegers1≤i,j≤k,or(c)hasπi,j(R)∈{(x,y)∈A2|x=y},{(x,y)∈A2|x+y≤0},A2 forallintegers1≤i,j≤k.Wewillshowthatincase(a),wehave(0,...,0)∈R,incase(b)wehave(+,...,+)∈R,andincase(c)wehave(−,...,−)∈R.Bysymmetry,weonlyhavetoconsidercases(a)and(b).Case(b)followsfromthefollowingclaim.Claim:ForanyrelationS≤Ansuchthat(+,+)∈πi,j(S)foralli,j,wehave(+,...,+)∈S.Proof:Wewillprove,byinductionon|I|thatforeverysubsetI⊆{1,2,...,n}thereisatuplesI∈Ssuchthatitsithcoordinateis+foralli∈I.Thebasecase|I|≤2isourassumptiononS.Fortheinductivestep|I|≥3,leti,j,andkbeanythreedistinctelementsofI.ThenwedeﬁnesIinductivelybysI:=sgn sI\{i},sI\{j},sI\{k},usingthefactthatthethree-variableoperationsgn(x+y+z)isintheclonegeneratedbythetwo-variableoperationsgn(x+y),asproveninthehf+0−icaseofsection5.3.Case(a)alsofollowsfromtheclaim.Toseethis,ﬁndamaximalsubsetI⊆{1,2,...,n}suchthatnopairofindicesi,j∈Ihasπi,j(R)={(x,y)∈A2|x=−y}.Thenwecanusetheclaimtoshowthatthetuplesgivenbyπi(s)=(+i∈I−i6∈I,isinR.Bysymmetry,−s∈Raswell.Thereforesgn((s)+(−s))=(0,...,0)∈R,sowearedone.
Usingsomestrongerbackgroundtheory,wecanconﬁrmthatConjecture21istrueforbinaryrelations.Theorem25.Conjecture21holdsforbinaryrelations:ifeverysubquotientofaﬁniteidempotentalgebraAhasanelementﬁxedbyitsautomorphismgroup,theneverybinaryreversiblerelationR≤A2containsaconstanttuple.25
Proof.Foridempotentalgebras,theassumptionimpliesthatAisTaylorbyProposition4.14of[4];infact,amoregeneralformofthisresultisprovedinProposition2.1of[8].LetAandBbetheunderlyingsetsofAandB,respectively.AssumewithoutlossofgeneralitythatRissubdirect;thatis,π1(R)=π2(R)=A.LetθbethelimitofthelinkingcongruenceofthebinaryrelationR◦mwhenmgetslarge.Analternativewaytodescribeθisasfollows:considerRtobetheedgesofadirectedgraphonA,andconsidertwoverticestobeequivalentifthereisanundirectedpathconnectingthemsuchthatthetotalnumberofforwardedgesalongthepathequalsthetotalnumberofbackwardedgesalongthepath.ThenR/θisthegraphofanautomorphismofA/θ,sobyassumptionthereisacongruenceclassBofθwhichisﬁxedbythisautomorphism.SinceAisidempotent,BisasubalgebraofA,andsinceBisﬁxedbythisautomorphismofA/θ,B+R=B.ThefactthatBisacongruenceclassofθisequivalenttotherestrictionofRtoBdeﬁningadirectedgraphof“algebraiclength1,”sowecanapplytheLoopLemmaof[1]toconcludethatRcontainsaconstanttuple(b,b)withb∈B.
7Acknowledgements
WewouldliketothanktheMIT-PRIMESprogram—includingDr.TanyaK

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

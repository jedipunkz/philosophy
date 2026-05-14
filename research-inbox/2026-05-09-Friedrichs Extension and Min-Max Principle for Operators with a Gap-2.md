---
source: "https://arxiv.org/abs/1806.05206v2"
title: "Friedrichs Extension and Min-Max Principle for Operators with a Gap"
author: "Lukas Schimmer, Jan Philip Solovej, Sabiha Tokus"
year: "2018"
publication: "arXiv preprint / math-ph"
download: "https://arxiv.org/pdf/1806.05206v2"
pdf: "https://arxiv.org/pdf/1806.05206v2"
captured_at: "2026-05-09T12:02:43Z"
updated_at: "2026-05-09T12:02:43Z"
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

# Friedrichs Extension and Min-Max Principle for Operators with a Gap

- 著者: Lukas Schimmer, Jan Philip Solovej, Sabiha Tokus
- 年: 2018
- 掲載情報: arXiv preprint / math-ph
- 情報源: [arxiv](https://arxiv.org/abs/1806.05206v2)
- ダウンロード: https://arxiv.org/pdf/1806.05206v2
- PDF: https://arxiv.org/pdf/1806.05206v2

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Semibounded symmetric operators have a distinguished self-adjoint extension, the Friedrichs extension. The eigenvalues of the Friedrichs extension are given by a variational principle that involves only the domain of the symmetric operator. Although Dirac operators describing relativistic particles are not semibounded, the Dirac operator with Coulomb potential is known to have a distinguished extension. Similarly, for Dirac-type operators on manifolds with a boundary a distinguished self-adjoint extension is characterised by the Atiyah--Patodi--Singer boundary condition. In this paper we relate these extensions to a generalisation of the Friedrichs extension to the setting of operators satisfying a gap condition. In addition we prove, in the general setting, that the eigenvalues of this extension are also given by a variational principle that involves only the domain of the symmetric operator.

## PDF Text

arXiv:1806.05206v2 [math-ph] 11 Jan 2019
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEFOROPERATORSWITHAGAPLUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUSAbstract.Semiboundedsymmetricoperatorshaveadistinguishedself-adjointex-tension,theFriedrichsextension.TheeigenvaluesoftheFriedrichsextensionaregivenbyavariationalprinciplethatinvolvesonlythedomainofthesymmetricoperator.AlthoughDiracoperatorsdescribingrelativisticparticlesarenotsemi-bounded,theDiracoperatorwithCoulombpotentialisknowntohaveadistin-guishedextension.Similarly,forDirac-typeoperatorsonmanifoldswithaboundaryadistinguishedself-adjointextensionischaracterisedbytheAtiyah–Patodi–Singerboundarycondition.InthispaperwerelatetheseextensionstoageneralisationoftheFriedrichsextensiontothesettingofoperatorssatisfyingagapcondition.Inadditionweprove,inthegeneralsetting,thattheeigenvaluesofthisextensionarealsogivenbyavariationalprinciplethatinvolvesonlythedomainofthesymmetricoperator.1.IntroductionandMainResultForasymmetric,semiboundedoperatorAwithdensedomainD(A)onaHilbertspaceHthereexistsadistinguishedself-adjointextension,theFriedrichsextensionAF.ThisextensionwasintroducedbyFriedrichs[16]in1934.Itseigenvaluescanbecomputedbyavariationalprinciple.Moreprecisely,ifAisboundedfrombelowbyλ1,whereλ1=infz∈D(A)hz,AziH
kzk2
H>−∞,(1)avariationalprinciple(seee.g.[7,Theorem4.5.2])statesthatthevaluesλk=infV⊂D(A)dimV=ksupz∈Vhz,AziH
kzk2
H(2)fork≥1arethediscretespectrumofAFintheinterval(−∞,supk≥1λk),countedwithmultiplicitiesdk:=#{j≥1:λj=λk}
2010MathematicsSubjectClassiﬁcation.49R05,49S05,47B25,81Q10.Keywordsandphrases.Self-adjointextension;Spectralgap;Variationalprinciple;Schurcomple-ment;Diracoperator.TheauthorsweresupportedbyERCAdvancedGrantno.321029andVILLUMFONDENthroughtheQMATHCentreofExcellence(grantno.10059).1
2LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
aslongasdk<∞.Ifdk=∞thenλkisintheessentialspectrumofAF.Whilesimilarvariationalprinciplesholdforallsemiboundedself-adjointextensionsofA,westressthatin(2)onlythedomainD(A)isneeded,makingthespectrumoftheFriedrichsextensionespeciallyaccessibletonumericalmethods.ThisisaconsequenceofD(A)beingaformcoreforAF.ForsymmetricoperatorsAthatarenotsemibounded,Friedrichs’constructionisnotapplicable.Ofparticularinterestisthecasewheretheself-adjointextensionofAisexpectedtohaveagapinitsspectrum.Inasimilarwaytothesemiboundedcase,onewouldliketosolvethefollowingproblems.
(P1)Deﬁneadistinguishedself-adjointextensionAFofA.(P2)ProvideasimplevariationalprinciplethatallowstocomputetheeigenvaluesofAF,ideallyonlyfromthesymmetricoperatorA.Inthispaper,wewillgeneralisetheconstructionoftheFriedrichsextensionAFtosymmetricoperatorsAwherethelowersemiboundedness(1)isreplacedbyagapcondition.WewillfurthermorerelatetheextensiontoavariationalprinciplethatonlyinvolvesthedomainofthesymmetricoperatorAhenceprovidingsolutionstobothproblems,(P1)and(P2).AnimportantexampleofanoperatorthatourresultsapplytoistheDiracoperatorHνonL2(R3;C4)withCoulombpotential−ν/|x|.TheoperatorHνisnotsemiboundedandforν≥√
3/2itisnotessentiallyself-adjointonthespaceofsmooth,compactlysupportedfunctionsC∞0(R3;C4).OurresultsalsoapplytoDirac-typeoperatorsonmanifoldswithaboundary.Fortheseoperators,thereexistsadistinguishedself-adjointextensionwhichcanbechar-acterisedbyanon-localboundarycondition,asﬁrstintroducedbyAtiyah,PatodiandSingerintheproofoftheirindextheorem[1].Wewillshowthatthisboundaryconditionnaturallyarisesfromtheconstructiongiveninthispaper.Theproblem(P1)hasbeenstudiedalreadybyKrein[20]forasymmetricoperatorAthatisnotsemiboundedbutsatisﬁesagapconditionoftheform(λ0<λ1)

A−λ0+λ1
2!z

H≥λ1−λ0
2kzkH.(3)InKrein’sworkitisprovedthatsuchanoperatorhasaself-adjointextensionthatpre-servesthegap,i.e.theinterval(λ0,λ1)belongstotheresolventsetoftheextension.SubsequentlyBrascheandNeidhart[4]parametrisedallgap-preservingself-adjointextensionsofAbyusingasuitablerepresentationfortheirinverses.Theauthors’parametrisationallowedthemtoidentifyoneoftheextensionsastheFriedrichsex-tensioninthelimitλ0→−∞.Thetypeofoperatorswewishtoconsiderheresatisfyagapconditionwhichisseentoimply(3),aswillbeprovedinRemark2.InanalogytotheFriedrichsextensionpreservingthelower-semiboundedness,ourextensionAFpreserves(3).
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP3Morerecently,diﬀerentformsofgapconditionshavebeenconsidered.EstebanandLoss[13]consideredablock-matrixoperator PQT−S!(4)denselydeﬁnedonadomainD0×D0⊂H0×H0whereP=P∗,S=S∗,Q=T∗andS≥−λ0>0.FurthermoretheyassumedthatP,Q,S,T,S−1TandQS−1TmapD0intoH0.Theirgapconditionwasphrasedintermsoftheassumptionthatforsomeλ1>0andallz∈D0qλ1(z,z):=h(S+λ1)−1Tz,TziH+h(P−λ1)z,ziH≥0.InthecaseofDiracoperatorsHνwithCoulombpotentialsthisassumptionconstitutesaHardyinequalitythatwaspreviouslyprovedanalyticallybyDolbeault,Esteban,LossandVega[8].InthiswayLossandEsteban[12]wereabletodeﬁneadistinguishedself-adjointextensionforHνuptoandincludingthecriticalvalueν=1.Forν<1theirextensioncoincideswiththepreviouslyknowndistinguishedextensionestablishedseparatelybySchmincke[27],Nenciu[24]andWüst[31](whichwereallprovedtobeequalbyKlausandWüst[18]).Regardingthesecondproblem(P2),variationalprincipleshavebeenstudiedbyseveralauthorsforself-adjointoperatorswithgaps.ForDiracoperatorswithnegativepotentialsTalman[28]aswellasDattaandDeviah[6]suggestedawaytocomputetheﬁrsteigenvalue.Theideawastosplittheoptimisationinthevariationalprinciple.DecomposingtheHilbertspaceintoadirectsumL2(R3;C4)=(L2(R3;C2)×{0})⊕({0}×L2(R3;C2))correspondingtotheupperandlowerspinors,theﬁrsteigenvaluewouldbegivenbyﬁrstmaximisingthequadraticformoveronecomponentandthenminimisingovertheother.Moreprecisely,forsuitablychosenspacesF+⊂L2(R3;C2)×{0},F−⊂{0}×L2(R3;C2)theauthorssuggestedthatλ1=infx+∈F+\{0}supy−∈F−hx++y−,A(x++y−)iH
kx++y−k2
H.ForDiracoperatorsHνwithCoulombpotentialssuchavariationalprincipledescribingthediscretespectrumwasprovedbyDolbeault,EstebanandSéré[9]inthecaseofessentiallyself-adjointnessν∈[0,√
3/2)wheretheycouldchooseF+=C∞0(R3;C2)×{0}andF−={0}×C∞0(R3;C2).Theirargumentforν∈(√
3/2,1)wasnotcomplete.Forν<1MorozovandMüller[22,23]showedthatF+=H1/2(R3;C2)×{0}andF−={0}×H1/2(R3;C2)arevalidchoicestoobtainavariationalprincipleforthedistinguishedextension.Inthegeneralsettingofaself-adjointoperatorwithspectralgap,variationalprin-ciplesthatuseanorthogonaldecompositionoftheHilbertspacewereinvestigated
4LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
byGriesemerandSiedentop[17].Abstractvariationalprincipleswerealsoprovedin[9,22]andwithdiﬀerentassumptionsbyKraus,LangerandTretter[19](seealso[30]).Inalltheseresultshowever,theoperatorisa-prioriassumedtobeself-adjointoressentiallyself-adjoint.OnlyrecentlyEsteban,LewinandSéré[11]extendedthevariationalprincipleforDiracoperatorswithCoulombpotentialstoallν∈[0,1]anddiscusseditsconnectionstothedistinguishedself-adjointextension.Buildingupontheresultsof[9]theyshowedthatforanyν∈[0,1]itissuﬃcienttochooseF+=C∞0(R3;C2)×{0}andF−={0}×C∞0(R3;C2)toobtaintheeigenvaluesofthedistinguishedextension,evokingsimilaritiestotheFriedrichsextension.Ourmainresult,Theorem1clariﬁestheconnectionbetweenadistinguishedself-adjointextensionandavariationalprincipleinthecaseofoperatorssatisfyingageneralgapcondition.Itapplies,inparticular,totheDirac–Coulomboperatorgeneralisingtheresultof[11].
Theorem1.LetAbeadenselydeﬁnedsymmetricoperatoronaHilbertspaceHandlethx,AyiHbethecorrespondingrealquadraticformwithformdomainequaltotheoperatordomainD(A).Furthermorethefollowingassumptionsaremade.(i)Orthogonaldecomposition:ThereareorthogonalprojectionsΛ±onHsuchthatH=Λ+H⊕Λ−H=H+⊕H−andF±:=Λ±D(A)⊂D(A).(ii)Gapcondition:supy−∈F−\{0}hy−,Ay−iH
ky−k2
H=:λ0<λ1:=infx+∈F+\{0}supy−∈F−hx++y−,A(x++y−)iH
kx++y−k2
H.(iii)TheoperatorΛ−A|F−:F−→H−isessentiallyself-adjoint.Thenthereexistsaself-adjointextensionAFofAsuchthatfork≥1thenumbersλk:=infV⊂F+dimV=ksupz∈(V⊕F−)\{0}hz,AziH
kzk2
H(5)aretheeigenvaluesofAFintheset(λ0,sup`≥1λ`)countedwithmultiplicitiesdk:=#{j≥1:λj=λk}aslongasdk<∞.Ifdk=∞thenλkisintheessentialspectrumofAF.TheoperatorAFistheuniqueself-adjointextensionwiththepropertythatD(AF)⊂F+⊕H−,forasubspaceF+⊂H+deﬁnedintheproof.
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP5Remark2.Assumptions(i)and(ii)ofTheorem1implythatAsatisﬁesthegapcondition(3).Toseethis,weletx=x++x−∈D(A)andforgivenε>0chooseyε−∈D(A)suchthathx++yε−,A(x++yε−)iH≥(λ1−ε)

x++yε−

2
H.(6)Thenwithλ:=(λ0+λ1)/2k(A−λ)xkH≥supz∈D(A)|<h(A−λ)x,ziH|
kzkH=supz∈D(A)|hx+z,(A−λ)(x+z)iH−hx−z,(A−λ)(x−z)iH|
4kzkH.Choosingz:=x+−x−+2yε−∈D(A)andusing(6)togetherwiththedeﬁnitionofλ0weobtainthelowerboundk(A−λ)xkH≥

hx++yε−,(A−λ)(x++yε−)iH−hx−−yε−,(A−λ)(x−−yε−)iH

kzkH≥ λ1−λ0
2−ε!kx+zk2
H+kx−zk2
H
4kzkH.Usingtheparallelogramlawandthefactthata+1/a≥2foranya>0weobtaink(A−λ)xkH≥ λ1−λ0
2−ε!kxkH
2 kxkH
kzkH+kzkH
kxkH!≥ λ1−λ0
2−ε!kxkH.Sinceε>0wasarbitrary,thegapcondition(3)holds.ByanapplicationofthespectraltheoremthesameholdstruefortheextensionAF.WeconstructAFasananaloguetotheFriedrichsextensionofasemiboundedoperator(seee.g.[25,TheoremVIII.15]and[26,TheoremX.23]aswellas[3,pp.224]).Wecloselyfollow[9],themainideabeingthefollowing.IfAisaboundedself-adjointoperatorsuchthatF±=H±,thenforanyE/∈σ(Λ−A|H−)thedecomposition Λ+A|H+Λ+A|H−Λ−A|H+Λ−A|H−!−EI= I−L∗
E0I! QE00−(B+E)! I0−LEI!(7)holds(seee.g.[30,Proposition1.6.2]),whereB=−Λ−A|H−,LE=(B+E)−1Λ−A|H+,QE=(Λ+A−E)|H++Λ+A|H−(B+E)−1Λ−A|H+.TheoperatorQEisoneoftwoSchurcomplementsofA.InSection2,wewillconstructAFbydeﬁningthesethreeoperators.ThedeﬁnitionofBinSubsection2.1isstraightforwardandyieldsanoperatorwithformdomaindenotedbyF−⊂H−.ComplicationsarisefromthefactthattheSchurcomplementisonlydeﬁnedintermsofaquadraticformqEwhichisnotnecessarilyclosableonH+.
6LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
ThusanewHilbertspaceG+,whichisobtainedwhenconsideringtheclosureLEoftheoperator(B+E)−1Λ−A|F+,hastobeintroducedinSubsection2.2.That(B+E)−1Λ−A|F+isclosableisnon-trivialanddoesnotseemtoholdtruewithoutassumption(iii).Forthisreasonwebelieve(iii)isnecessarytoguaranteethatG+canbeidentiﬁedwithasubspaceofH+.OnG+,wecancloseqEanddeﬁnethecorrespondingoperatorQE,asdoneinSubsection2.3.ParticularconsiderationhastobegiventothefactthattheconstructiondoesnotdependontheexplicitchoiceofE>λ0.InSubsection2.4thedeﬁnitionoftheself-adjointextensionAFisgiveninaformthatresemblestheabovedecomposition(7).InSubsection2.6thevariationalprinciplestatedinTheorem1willbeproved.Table1summarisestheHilbertspacesthatneedtobedeﬁnedwhileTable2listsalltheadditionalspaces.InSection3wewillapplyTheorem1totheDirac–Coulomboperator.InSection4wewillintroducetheAPS-boundaryconditionforgeneralisedDirac-operatorsandprovethattheself-adjointextensionconstructedaccordingtoTheorem1isexactlycharacterisedbytheseboundaryconditions.
Remark3.Ourconstructionofthedistinguishedself-adjointextensiondiﬀersfrom[13].Phrasingourassumptionsintermsoftheblock-matrixnotation(4),wedonotrequirePandStobeself-adjointnorthatQ=T∗.Inaddition,wedonotmakeanyassumptionaboutthedomainofQS−1T.Withthesetupasin[13]thequadraticformqEisclosableonH+anditisclaimed(butnotproved)in[13]thatthedomainoftheclosureisindependentofE.InourconstructiontheintroductionofG+isnecessarytoguaranteeboththatqEisclosableonG+andthatthedomainoftheclosureandhencetheself-adjointextensiondonotdependonthechoiceofE.Neverthelessourconstructionisinspiredbytheapproachin[12]and[13].
Hilbertspace
(Equivalent)norms
Containedin
Description
Page
H
k·kH
p.4
H+
k·kH
H+⊂H
Λ+H
p.4
H−
k·kH
H−⊂H
Λ−H
p.4
F−
k·kF−
F−⊂H−
FormdomainofB
p.7
G+
k·kE,E>λ0
G+⊂H+
DomainofLE
p.8
F+
k·kF+,E,E>λ0
F+⊂G+
FormdomainofQE
p.11
Table1.TherequiredHilbertspaces
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP7
Space
Containedin
Description
Page
D(AF)
D(AF)⊂F+⊕H−
DomainofAF
p.15
D(A)
D(A)⊂D(AF)
DomainofA
p.4
F+
F+⊂F+
Λ+D(A)
p.4
F−
F−⊂F−
Λ−D(A)
p.4
Table2.Theadditionallyrequiredvectorspaces2.TheProofofTheorem12.1.TheDeﬁnitionofB.Westartbysettinghy−,z−iF−:=(λ0+1)hy−,z−iH−hy−,Az−iHwhichbydeﬁnitionofλ0isaninnerproductonF−withcorrespondingnormky−k2
F−=(λ0+1)ky−k2
H−hy−,Ay−iH.Sincethequadraticformh·,·iF−comesfromasymmetricoperatoritisclosable,i.e.itextendstoaclosedquadraticformontheformdomainF−⊂H−,whichistheclosureofF−withrespecttothenormk·kF−.Ifwedenotethecontinuousextensionofthequadraticformh·,·iF−toF−byh·,·iF−,then(F−,h·,·iF−)formsaHilbertspace.ByassumptionΛ−A|F−isessentiallyself-adjoint,hencethereexistsauniqueself-adjointextensiongivenbyitsclosure,whichwewilldenoteby−B.ItisthenclearthatB+λ0+1coincideswiththeself-adjointoperatorassociatedwiththeclosedquadraticformh·,·iF−suchthathy−,z−iF−=hy−,Bz−iH+(λ0+1)hy−,z−iH.TheformdomainofBisF−anditsoperatordomainD(B)isasubsetofF−.ForE>λ0theself-adjointoperatorB+Eisstrictlypositiveanditsinverse(B+E)−1iswell-deﬁnedandboundedonallofH−.Remark4.SinceΛ−A|F−isessentiallyself-adjoint,theoperatorBcoincideswiththeFriedrichsextensionofthesemi-boundedoperator−Λ−A|F−.Fortheconvenienceofthereaderandtoevokeconnectionstoourconstruction,werecallthedeﬁnitionofthisextension.UsingRiesz’theoremweﬁrstdeﬁnetheoperatorbPastheisometricisomorphismbetweentheHilbertspaceF−anditsdualF0−,i.e.foranyz−∈F−wedeﬁnebPz−∈F0−tobetheuniquecontinuousfunctionalsuchthat[bPz−](y−)=hz−,y−iF−.Withtheembeddingj−:H−→F0−givenby[j−(y−)](z−)=hy−,z−iF−(identifyingH−withitsdualspaceH0−)wecanshowthatonthedomainD(P)=nz−∈F−⊂H−:bPz−∈j−(H−)o⊂H−theoperatorP=j−1−◦bPisaself-adjointextensionof−Λ−A|F−+λ0+1.TheFriedrichsextensionBisthendeﬁnedasB=P−λ0−1withdomainD(B)=
8LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUSD(P).Inparticular,thequadraticformhy−,Bz−iHhasacontinuousextensiontoallx−,y−∈F−givenby[bB(y−)](z−)wherebB=bP−(λ0+1)j−.TheformdomainofBisconsequentlyF−.2.2.TheDeﬁnitionofLE.LetE>λ0.Thenforx+∈F+themappingx+7→(B+E)−1Λ−Ax+deﬁnesalinearoperatorfromF+intoH−.Theproofofthesecondpartofthefollowinglemmaisadaptedfrom[9,Lemma2.1].
Lemma5.Theoperator(B+E)−1Λ−AdeﬁnedonF+isclosable.WedenoteitsclosurebyLEwithgraphnormkx+k2
E=kx++LEx+k2
H=kx+k2
H+kLEx+k2
H.Forλ0<E≤E0thenormsk·kEandk·kE0areequivalentonF+withkx+kH≤kx+kE0≤kx+kE≤CE,E0kx+kE0,(8)whereCE,E0=(E0−λ0)/(E−λ0)≥1.Proof.Weﬁrstshowclosability.Considerasequenceofxn∈F+withkxnkH→0andy∈H−withk(B+E)−1Λ−Axn−ykH→0.Wehavetoshowthaty=0.Letz∈(B+E)F−⊂H−.Then|hz,(B+E)−1Λ−AxniH|=|h(B+E)−1z,Λ−AxniH|=|hA(B+E)−1z,xniH|≤

A(B+E)−1z

HkxnkH→0.SinceΛ−A|F−isessentiallyself-adjoint,wecanconcludethat(B+E)F−isdenseinH−andthusy=0.Next,assumeλ0<E≤E0.Thentheﬁrsttwoinequalitiesin(8)followdirectlyfromthedeﬁnitionofthenorms.Foraboundonk·kEintermsofk·kE0wenotethatbythespectraltheoremforx∈F−

(B+E)−1(B+E0)x

2
H≤supλ≥−λ0|λ+E0|2
|λ+E|2kxk2
H≤(E0−λ0)2
(E−λ0)2kxk2
H.AsaconsequenceweobtainwithCE,E0:=(E0−λ0)/(E−λ0)foranyx+∈F+kLEx+kH=

(B+E)−1Λ−Ax+

H≤CE,E0

(B+E0)−1Λ−Ax+

H=CE,E0kLE0x+kH,whichproves(8).WeconcludethatthedomainofLE,meaningtheclosureofF+withrespecttothenormk·kE,canbeidentiﬁedforallvaluesofE>λ0andwewilldenotethisvectorspacebyG+.Togetherwiththeinnerproducthx+,z+iE:=hx+,z+iH+hLEx+,LEz+iH−itformsaHilbertspace(G+,h·,·iE)andwehavethevectorspaceinclusionsF+⊂G+⊂H+,wherethelastequationalsoholdsinthesenseofHilbertspaces.
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP9Viewedasanoperatorfrom(G+,k·kE)to(H−,k·kH),LEisthenbounded.WewilllaterconsidertheLEasanoperatoronanevensmallerHilbertspace,whereitisconsequentlyalsobounded.
2.3.TheDeﬁnitionofQE.ForE>λ0wenowdeﬁnethequadraticformqEonF+×F+qE(x+,z+):=hx+,(A−E)z+iH+hΛ−Ax+,(B+E)−1Λ−Az+iH.ItisthequadraticformrelatedtooneoftheSchurcomplementsofthematrixrep-resentationofA.WewillseethatqEcanbeclosedasalower-semiboundedformontheHilbertspaceG+suchthattheclosureisindependentofE.Tothisendweﬁrstderivethefollowingresultwhichcanalsobefoundin[9,pp.210].Lemma6.ForE>λ0andx+∈F+letϕE,x+:F−→RbethefunctiondeﬁnedasϕE,x+(y−):=hx++y−,A(x++y−)iH−Ekx++y−k2
H.ThequadraticformqEisthenrelatedtoϕE,x+byqE(x+,x+)=supy−∈F−ϕE,x+(y−).Inparticular,ϕE,x+(·)canbeextendedtoF−andtheextensionattainsitsmaximumattheuniquepointymax=LEx+=(B+E)−1Λ−Ax+.Proof.Fory−∈F−wewriteϕE,x+(y−)=hx+,(A−E)x+iH+2<hy−,Ax+iH−hy−,(B+E)y−iH.(9)ItisthenclearthatthefunctionalϕE,x+(·)naturallyextendstoF−,seealsoRemark4.Wedenotethecontinuousextensionby
ϕE,x+.Thequadraticpolynomialf:R→Rwhichwecandeﬁneforanyy−,z−∈F−asf(h)=
ϕE,x+(y−+h(z−−y−)),h∈Risstrictlyconcaveandthuswehavef(1)<f(0)+f0(0).(10)Nowassumethaty−∈F−satisﬁestheEulerequation,thatis
ϕ0
E,x+(y−;(z−−y−))=0forallz−∈F−.Thenwemusthavehw−,Λ−Ax+−(B+E)y−iH=0forallw−∈F−orequivalentlyy−=(B+E)−1Λ−Ax+(11)andby(10)forallz−∈F−,z−6=y−
ϕE,x+(z−)<
ϕE,x+(y−),
10LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
i.e.
ϕE,x+(·)hasauniqueglobalmaximumatthepointymax=(B+E)−1Λ−Ax+∈D(B).Inserting(11)into(9)weobtain
ϕE,x+(ymax)=hx+,(A−E)x+iH+hΛ−Ax+,LEx+iH.ThefollowinglemmaestablishesimportantpropertiesofthequadraticformqE.Theproofcanalsobefoundin[9,Lemma2.1].
Lemma7.Letλ0<E≤E0.(i)OnF+thequadraticformsqEandqE0satisfyqE0(x+,x+)+(E0−E)kx+k2
E0≤qE(x+,x+)≤qE0(x+,x+)+(E0−E)kx+k2
E.(ii)ThequadraticformqEisboundedfrombelowonF+⊂G+byaconstantκE≥1suchthatqE(x+,x+)+κEhx+,x+iE≥hx+,x+iE.(iii)ItholdsthatE<λ1ifandonlyifqE(x+,x+)>0forallx+∈F+\{0},E≤λ1ifandonlyifqE(x+,x+)≥0forallx+∈F+\{0}.Proof.For(i)weusetheresolventidentitytocomputeforanyλ,λ0>λ0,qλ0(x+,x+)=qλ(x+,x+)+(λ−λ0)hkx+k2
H+hΛ−Ax+,(B+λ)−1(B+λ0)−1Λ−Ax+iHi.Theresultthenfollowsbysettingλ=E,λ0=E0andλ=E0,λ0=E,respectively,andusing(B+E0)−1≤(B+E)−1toboundthelastterm.Wecontinuetoshow(iii)andsubsequently(ii).First,letλ0<E<λ1andx+∈F+\{0}arbitrary.Bythedeﬁnitionofλ1,foranyε>0thereexistsayε−∈F−suchthathx++yε−,A(x++yε−)iH
kx++yε−k2
H≥λ1−εandconsequentlyqE(x+,x+)≥ϕE,x+(yε−)=hx++yε−,A(x++yε−)iH−E

x++yε−

2
H≥(λ1−ε−E)

x++yε−

2
H.Wecanconcludethatforallx+∈F+\{0}andallλ0<E<λ1qE(x+,x+)≥(λ1−E)kx+k2
H>0.SettingE0:=λ1andusing(i)wehavethatqλ1(x+,x+)≥qE(x+,x+)−(λ1−E)kx+k2
E,whichinthecaseE→λ1showsqλ1≥0sincekx+kE→kx+kλ1by(8).
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP11IfE>λ1thenagainbydeﬁnitionofλ1,foranyε>0withλ1<E−εthereexistsxε
+∈F+\{0}withh(xε
++y−),A(xε
++y−)iH
kxε
++y−k2
H≤(E−ε)forally−∈F−andconsequentlyqE(xε
+,xε
+)=supy−∈F−ϕE,xε
+(y−)≤−infy−∈F−ε

xε
++y−

2
H≤−ε

xε
+

2
Hwhichﬁnishestheproofof(iii).Thestatementin(ii)isclearforE≤λ1.IfE>λ1weuse(8)tocomputethatqE(x+,x+)≥qλ1(x+,x+)−(E−λ1)kx+k2
λ1≥qλ1(x+,x+)−(E−λ1)Cλ1,Ekx+k2
E.ChoosingκE=1+max(0,(E−λ1)Cλ1,E)thengivestheresult.Remark8.IntheﬁrstpartoftheproofweneededboundsontermsofthetypekLEx+kH.Thiswasdonebyusingthenewnormk·kE.WithoutspecifyingfurtherassumptionsontheoperatorAitisnotpossibletoestimatethediﬀerencebetweenquadraticformsqEfordiﬀerentvaluesofEbyak·kH-normonly.IntroducingtheHilbertspaceG+thusturnsouttobeessential.AnimmediateconsequenceofLemma7togetherwithLemma5isthatthecomple-tionF+ofF+withrespecttothenormk·kF+,Einducedbytheinnerproducthx+,z+iF+,E=qE(x+,z+)+κEhx+,z+iEisindependentofE.IntheremainderweﬁxE>λ0anddenotetheextensionoftheinnerproducth·,·iF+,EtoF+byh·,·iF+,E.Apriori,itisnotclearthatF+isasubspaceofH+.However,thefollowingholds.Lemma9.ThesemiboundedquadraticformqEisclosableontheHilbertspaceG+forE>λ0.Theclosure qEhastheformdomainF+,independentofE,andcanbeidentiﬁedwithasubspaceofG+andsubsequentlyalsoofH+.Proof.Weshowthatthepositiveformh·,·iE=qE(·,·)+κEh·,·iEisclosable.Considerasequencexn∈F+whichisaCauchysequencewithrespecttok·kF+,EandwhichsatisﬁeskxnkE→0.Thenforanyz∈F+|hz,xniF+,E|≤κE|hz,xniE|+|hz,(A−E)xniH|+|hΛ−Az,LExniH|=κE|hz,xniE|+|h(A−E)z,xniH|+|hΛ−Az,LExniH|≤κE(kzkE+k(A−E)zkH+kΛ−AzkH)kxnkEandthushz,xniF+,E→0,whereweagaincruciallyneedtheassumptionthatxn→0inthek·kE-norm,whichisstrongerthanthek·kH-norm.SinceF+isdenseinF+withrespecttok·kF+,E,wecanconcludethatkxnkF+,E→0.
12LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUSWethushavetheHilbertspaceinclusionsF+⊂G+⊂H+,withtheirrespectiveinnerproductsimplicit,andthecorrespondinginclusionsoftheassociateddualspacesH0+⊂G0+⊂F0+.ByRiesz’theoremthereexistsanisometricisomorphismiX→X0(x)=hx,·iXbetweeneachHilbertspaceXanditsdualspaceX0.IngeneralwewillnotexplicitlywritetheisomorphismsiH±→H0
±thusidentifyingHanditsdualspaceH0.FurthermoreforeachoftheHilbertspaceinclusionsX⊂YthereisacorrespondingembeddingofdualspacesjY0→X0:Y0→X0,inthesense[jY0→X0`](x)=`(x)=hi−1Y→Y0`,xiY,`∈Y0,x∈X.Alltheseembeddingsareboundedinnormbyone.Associatedtotheclosedquadraticform qEthereisanoperatordeﬁnedonalloftheformdomainF+,aswellasaself-adjointoperatorwithdomainasubsetofF+.Lemma10.LetE>λ0.ThereexistsanoperatordQE:F+⊂G+→F0+withthefollowingproperties.(i)Forallx+,z+∈F+theclosure qEofqEonF+isgivenby qE(x+,z+)=[dQEx+](z+).(ii)TheoperatordQEisboundedandifadditionallyE<λ1thenitsinversedQE−1isalsobounded.(iii)OnthedensedomainD(QE)=nz+∈F+:dQEz+∈jG0+→F0+(G0+)o⊂G+theoperatorQE:=i−1G+→G0+◦j−1G0+→F0+◦dQE:D(QE)→G+isself-adjointandQE+κE≥1.IfadditionallyE<λ1thenQEisalsopositive.Proof.WedeﬁnebS:F+→F0+usingRiesz’theoremastheuniqueoperatorsuchthat[bSz+](y+)=hz+,y+iF+,E.TheoperatordQE=bS−κEjG0+→F0+iG+→G0+thenhastheclaimedproperties.
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP132.4.TheDeﬁnitionofAF.WeconsideroncemoretheoperatorLE,viewednowasamappingfrom(F+,k·kF+,E)into(H−,k·kH).ThisoperatorisboundedandwedenoteitsadjointbyL0
E:(H−,k·kH)→(F0+,k·kF0+,E),whichisrelatedtotheHilbertadjointL∗
EbyL∗
E=i−1F+→F0+L0
EiH−→H0
−.ThisallowsustodeﬁnetheoperatordRE:D(dRE)⊂F+×H−→F0+×H−forE>λ0asdRE x+y−!= I−L0
E0I! dQE00−(B+E)! I0−LEI! x+y−!= dQEx++L0
E(B+E)(y−−LEx+)−(B+E)(y−−LEx+)!onthedomainD(dRE)=( x+y−!∈F+×H−:y−−LEx+∈D(B))⊂H+×H−.TheconstructionofdREshouldbecomparedtothedecomposition(7).Bytheresolventidentityforanyx+∈F+LEx+−LE0x+=(E0−E)(B+E)−1LE0x+andbyLemma5thisidentityextendstoF+.ThusD(dRE)isindependentofE>λ0andthesameholdsforthecorrespondingsubsetFofHF:={x++y−∈F+⊕H−:y−−LEx+∈D(B)}⊂H+⊕H−.IfE<λ1theoperatordQEisinvertibleandinthiscasedREhasaninversedeﬁnedbydRE−1 x+y−!= I0LEI!
dQE−100−(B+E)−1
 IL0
E0I! `+k−!= dQE−1(`++L0
Ek−)LEdQE−1(`++L0
Ek−)−(B+E)−1k−!forall(`+,k−)∈F0+×H−.ItisstraightforwardtoseethatthisoperatormapsintothedomainofdREandviceversa.WenowdeﬁneD(R)⊂F⊂F+⊕H−tobethesetD(R)=nx++y−∈F:dQEx++L0
E(B+E)(y−−LEx+)∈j+(H+)owhichwillbeprovedtobeindependentofE.OnthisdomainwedeﬁneforE>λ0thefamilyofoperatorsRE:D(R)→HactingasRE(x++y−)=j−1+(dQEx++L0
E(B+E)(y−−LEx+))−(B+E)(y−−LEx+).Hereweusethenotationj+fortheembeddingjH+→F0+.InthefollowingweprovethatREisanextensionofA−E,thatitsdomainisindeedindependentofEandthatitisself-adjoint.
14LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUSToseethatREisanextensionofA−E,wenotethatforx+∈F+,y−∈F−wehavey−−LEx+=y−−(B+E)−1Λ−Ax+∈D(B).Furthermore,foranyu+∈F+wecomputethat[dQEx++L0
E(B+E)(y−−LEx+)](u+)=hx+,(A−E)u+iH+hΛ−Ax+,LEu+iH+h(B+E)(y−−(B+E)−1Λ−Ax+),LEu+iH=hx+,(A−E)u+iH+h(B+E)y−,LEu+iH=h(A−E)x+,u+iH+hAy−,u+iH.Thelinearfunctionalh(A−E)x+,·iH+hAy−,·iHisboundedonF+andextendscontin-uouslytoF+.HencedQEx++L0
E(B+E)(y−−LEx+)∈j+(H+)andF+⊕F−⊂D(AF).Sinceinadditionforanyv−∈F−−h(B+E)(y−−LEx+),v−iH=−h(B+E)y−,v−iH+hΛ−Ax+,v−iH=h(A−E)y−,v−iH+hAx+,v−iHweobtainthatforallx+,u+∈F+,y−,v−∈F−hRE(x++y−),u++v−iH=h(A−E)(x++y−),u++v−iHwhichallowsustoconcludethatREisanextensionofA−E.ToshowthatD(R)isindependentofE,weﬁrstnotethatbytheaboveforanyx+,u+∈F+andy−∈F−
qE(x+,u+)+hy−−LEx+,(B+E)LEu+iH−
qE0(x+,u+)−hy−−LE0x+,(B+E0)LE0u+iH=(E0−E)hx+,u+iH.(12)Letnowx++y−∈D(R)⊂F+⊕H−,suchthatforsomeE>λ0y−−LEx+∈D(B),[dQEx++L0
E(B+E)(y−−LEx+)]∈j+(H+).Wehavealreadyseenthatthenalsoy−−LE0x+∈D(B)foranyE0>λ0.Wecanapproximatex++y−byelementsofx(n)++y(n)−∈F+⊕F−suchthat

x+−x(n)+

F+,E→0,

y−−y(n)−

H→0.Bycontinuity(12)extendstox++y−andweobtainthatforallu+∈F+[dQEx++L0
E(B+E)(y−−LEx+)](u+)−[dQE0x++L0
E0(B+E0)(y−−LE0x+)](u+)=(E0−E)hx+,u+iHandthusalsoy−−LE0x+∈D(B),[dQE0x++L0
E0(B+E0)(y−−LE0x+)]∈j+(H+).
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP15ToprovethatREissymmetric,wecomputethatforgivenu++v−,x++y−∈D(R)hRE(x++y−),u++v−iH=hj−1+(dQEx++L0
E(B+E)(y−−LEx+))−(B+E)(y−−LEx+),u++v−iH=hj−1+(dQEx++L0
E(B+E)(y−−LEx+),u+iH−h(B+E)(y−−LEx+),v−iH=[dQEx+](u+)+h(B+E)(y−−LEx+),LEu+iH−h(B+E)(y−−LEx+),v−iH=
qE(x+,u+)−h(B+E)(y−−LEx+),(v−−LEu+)iH.Thislastexpressionissymmetricininterchangingu++v−andx++y−andhenceREisasymmetricoperator.ForE<λ1,theoperatorR−1E:H=H+⊕H−→HdeﬁnedasR−1E(x++y−)=dQE−1(j+(x+)+L0
Ey−)+LEdQE−1(j+(x+)+L0
Ey−)−(B+E)−1y−istheinverseofRE.ItisitselfsymmetricandsincedeﬁnedonallofH,self-adjoint.BytheHellinger–Toeplitztheoremitisalsoaboundedoperator,henceclosed.ButthenREitselfasabijective,symmetricandclosedoperatorisalsoself-adjoint.Theself-adjointnessthenextendstoREforanyE>λ0.Lastly,wedeﬁnetheextensionAFofAasAF:=RE+EonD(AF):=D(R).2.5.TheUniquenessofAF.LeteAbeanotherself-adjointextensionofAwithD(eA)⊂F+⊕H−.WeﬁrstshowthatthennecessarilyD(eA)⊂F.Forx++y−∈D(eA)andv−∈F−wecomputeh(eA−E)(x++y−),v−iH=hx++y−,(A−E)v−iH=hx++y−,(AF−E)v−iH=−hy−−LEx+,(B+E)v−iHandwecanconcludethatv−7→hy−−LEx+,(B+E)v−iHisacontinuousfunctionalforallv−∈F−.Thisimpliesy−−LEx+∈D(Λ−A|∗
F−)=D(B)andthusx++y−∈F.Takingu+∈F+wefurthercomputeh(eA−E)(x++y−),u+iH=hx++y−,(A−E)u+iH=hx++y−,(AF−E)u+iH=hx+,j−1+(dQEu+−L0
E(B+E)LEu+)iH+hy−,(B+E)LEu+iH=
[dQEu+](x+)−h(B+E)LEu+,LEx+iH+h(B+E)LEu+,y−iH=[dQEx+](u+)+h(B+E)(y−−LEx+),LEu+iH=[dQEx++L0
E(B+E)(y−−LEx+)](u+).Fromthiswecanconcludethatx++y−∈D(AF)=D(R)andthusD(eA)⊂D(AF).Conversely,byself-adjointness,D(AF)=D(A∗
F)⊂D(eA∗)=D(eA),whichprovesthedesiredeA=AF.
16LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
2.6.TheProofoftheVariationalPrinciple.Itremainstoprovethatthevaria-tionalprincipleholds.Themin-maxlevelsofQEon(G+,h·,·iE)aregivenbyµk(QE)=infV⊂F+dimV=ksupx+∈V\{0}
qE(x+,x+)
kx+k2
E=infV⊂F+dimV=ksupx+∈V\{0}qE(x+,x+)
kx+k2
EwhereweusedthatF+isaformcoreof qE.Thenumbersµk(QE)satisfyµk(QE)≤infσess(QE)andifµk(QE)<infσess(QE)thenµkisaneigenvalueofQEwithmulti-plicitymk(QE)=#{j≥1:µj(QE)=µk(QE)}.Weneedthefollowingresult,whichcanbefoundin[9,Lemma2.2].Lemma11.UndertheassumptionsofTheorem1,itholdsthat:(i)Foranyx+∈F+\{0}therealnumberE(x+):=supz∈(span(x+)⊕F−)\{0}hz,AziH
kzk2
Histheuniquesolutionin(λ0,+∞)ofqE(x+,x+)=0,whichmayalsobewrittenasEkx+k2
H=hx+,Ax+iH+hΛ−Ax+,LEx+iH.(ii)Thevariationalprinciple(5)isequivalenttoλk=infV⊂F+dimV=ksupx+∈V\{0}E(x+).(iii)Foranyk≥1therealnumberλkgivenby(5)istheuniquesolutionofµk(Qλ)=0.Proof.FirstnotethatqE(x+,x+)forﬁxedx+∈F+\{0}isastrictlydecreasing,contin-uousfunctionofEwithqλ1(x+,x+)≥0byLemma7andwithlimE→∞qE(x+,x+)=−∞.WecanconcludethatqE(x+,x+)=0haspreciselyonesolutionin[λ1,+∞).DenotethissolutionbyeE(x+).IfE<eE(x+)thennecessarilyqE(x+,x+)>0andthusthereexistsay−∈F−suchthathx++y−,A(x++y−)iH−Ekx++y−k2
H>0.WeobtainthatE(x+)=supz∈(span(x+)⊕F−)\{0}hz,AziH
kzk2
H≥hx++y−,A(x++y−)iH
kx++y−k2
H>E.
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP17IfE>eE(x+)thennecessarilyqE(x+,x+)≤−ε<0forsomeεandthusforally−∈F−hx++y−,A(x++y−)iH−Ekx++y−k2
H≤−ε.Consequently,E(x+)=supz∈(span(x+)⊕F−)\{0}hz,AziH
kzk2
H≤−ε+E<E.ThisprovesthateE(x+)=E(x+).Thestatement(ii)isanimmediateconsequenceofthedeﬁnitionsofE(x+)andλkaswellastheobservationthat,sinceλ0<λ1,foranyk-dimensionalsubspaceV⊂F+supz∈(V⊕F−)\{0}hz,AziH
kzk2
H=supz∈V⊕F−Λ+z6=0hz,AziH
kzk2
H.Notethatµk(Qλ)isacontinuousfunctionofλwithµk(Qλ1)≥0byLemma7.Furthermorelimλ→+∞µk(Qλ)=−∞andwecanconcludethatµk(Qλ)=0hasatleastonesolutionin[λ1,+∞).Denotethissolutionbyeλk.Assumeλ<eλk.ForallV⊂F+withdimV=kthereexistsanxV
+∈V\{0}suchthatqeλk(xV
+,xV
+)≥−eλk−λ
2

xV
+

2eλkandthusbyLemma7qλ(xV
+,xV
+)≥qeλk(xV
+,xV
+)+(eλk−λ)

xV
+

2eλk≥eλk−λ
2

xV
+

2eλk>0.ThisimpliestheexistenceofyV−∈F−suchthatϕλ,xV
+(yV−)=hxV
++yV−,A(xV
++yV−)iH−λ

xV
++yV−

2
H≥0.Weobtainthatsupz∈(V⊕F−)\{0}hz,AziH
kzk2
H≥hxV
++yV−,A(xV
++yV−)iH
kxV
++yV−k2
H≥λandthusλ≤λk.Assumeλ>eλk.ThereexistsavectorspaceV0suchthatforallx+∈V0qeλk(x+,x+)≤λ−eλk
2Ceλk,λkx+k2eλk≤λ−eλk
2kx+k2
λandthusbyLemma7qλ(x+,x+)≤qeλk(x+,x+)−(λ−eλk)kx+k2
λ≤−λ−eλk
2kx+k2
λ<0forallx+∈V0\{0}.Thisimpliesthathx++y−,A(x++y−)iH−λkx++y−k2
H≤0
18LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
forallx+∈V0\{0}andally−∈F−.Wecanconcludethatλk≤supz∈(V⊕F−)\{0}hz,AziH
kzk2
H≤maxsupx+∈V0\{0}supy−∈F−hx++y−,A(x++y−)iH
kx++y−k2
H,supy−∈F−\{0}hy−,Ay−iH
ky−k2
H≤max(λ,λ0)=λandtogetherwiththeaboveλk=eλk.ToprovethattherealnumbersλkareinthespectrumofAF,weuseanargumentpresentedin[9,Section2]andconstructasequenceofsubspacesXnofdimensiondksuchthatlimn→∞supx∈XnkxkH=1supy∈D(AF)\{0}|hx,(AF−λk)yiH|
q kyk2
H+kAFyk2
H=0.Firstnotethatdk:=#{j≥1:λj=λk}=#{j≥1:µj(Qλk)=µk(Qλk)=0}=mk(Qλk)andbythemin-maxprincipleforQλk,thereexistsasequenceofspacesX+n⊂D(Qλk)ofdimensiondksuchthatlimn→∞supx+∈X+nkx+kλk=1kQλkx+kλk=0,whichalsoimpliesthatlimn→∞supx+∈X+nkx+kλk=1

dQλkx

F0+,λk=limn→∞supx+∈X+nkx+kλk=1supy+∈F+|[dQλkx+](y+)|
ky+kF+,λk=0.(13)LetXn:=(1+Lλk)X+n⊂F.Weobservethatforallx+∈F+andy∈D(AF)[dQλk(x+)](Λ+y)=hx++Lλkx+,(AF−λk)yiH.(14)Furthermoreforally∈D(AF)[dQλk(Λ+y)](Λ+y)=hΛ+y+LλkΛ+y,(AF−λk)yiH≤(kykH+kΛ−y−LλkΛ+ykH)(1+|λk|)(kykH+kAFykH)andusingkΛ−y−LλkΛ+ykH=

(B+λk)−1Λ−(AF−λk)y

H≤1+|λk|
λk−λ0(kykH+kAFykH)wecanseethatforally∈D(AF)thereexistsaconstantCλk>0suchthat[dQλk(Λ+y)](Λ+y)≤Cλk(kyk2
H+kAFyk2
H).
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP19ThisallowsustoboundkΛ+yk2
F+,λk=κλkkΛ+y+k2
H+κλkkLλkΛ+yk2
H+[dQλk(Λ+y)](Λ+y)≤C0λk(kyk2
H+kAFyk2
H)withsomeconstantC0λk>0forally∈D(AF).Togetherwith(13),(14)andthefactthatkx++Lλkx+kH=kx+kλkwecanconcludethatlimn→∞supx∈XnkxkH=1supy∈D(AF)\{0}|hx,(AF−λk)yiH|
q kyk2
H+kAFyk2
H=0.(15)As1+Lλk:X+n→XnisasurjectiveisometryweobtainthatdimXn=dk.Considerthecasedk<∞.LetPbethespectralmeasureofAF.Supposenowforsomeε>0wehadthatdimranP((λk−ε,λk+ε))≤dk−1.Thenthereexistsasequenceofxn∈XnwithkxnkH=1andP((λk−ε,λk+ε))xn=0.Wewritexn=wn+znwithwn∈ranP((−∞,λk−ε])andzn∈ranP([λk+ε,∞)).UnlessλkisintheessentialspectrumofAF,wealsoobservethatforsomeν∈(λk−ε,λk+ε)necessarilyν∈ρ(AF).Choosingyn=(AF−ν)−1xn∈D(AF)wecomputethat,ifλk−ε<ν≤λk,hxn,(A−λk)yniH=Zλk−ε−∞λ−λk
λ−νdPwn,wn(λ)+Z∞λk+ελ−λk
λ−νdPzn,zn(λ)≥kwnk2
H+ε
λk+ε−νkznk2
H≥min1,ε
λk+ε−νkxnk2
Handsimilarly,ifλk≤ν<λk+ε,hxn,(A−λk)yniH≥minε
ν−λk+ε,1kxnk2
H.SincekxnkH=1andkynk2
H+kAFynk2
H=

(AF−ν)−1xn

2
H+

xn+ν(AF−ν)−1xn

2
H≤Ckxnk2
HweobtainthatforsomeconstantC0>0|hxn,(AF−λk)yniH|≥C0q kynk2
H+kAFynk2
H,whichcontradicts(15).ThusnecessarilydimranP((λk−ε,λk+ε))≥dkforanyε>0.Inthecasedk=∞,wecanusetheaboveargumenttoconcludethatdimranP(λk−ε,λk+ε)=∞forallε>0.Asaconsequenceλk∈σ(AF)andλkislargerorequaltothek-theigenvalueµk(AF)ofAFin(λ0,sup`≥1λ`).
20LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUSBeforeweprovethattheλkareallthepointsinσ(AF)∩(λ0,sup`≥1λ`),weﬁrstnotethat(seeRemark4)λ0=supy−∈F−\{0}−hy−,By−iH
ky−k2
H=supy−∈F−\{0}−[bBy−](y−)
ky−k2
H=supy−∈F−∩D(AF)\{0}hy−,AFy−iH
ky−k2
H(16)whichisanimmediateconsequenceofthecontinuityofbBwithrespecttok·kF−.Nowassumethatλ∈σ(AF)∩(λ0,sup`≥1λ`)withspectralmultiplicityd.Wehavetoshowthatλ=λkforsomek∈N,orequivalentlythatµk(Qλ)=0forsomek∈N.ByassumptionthereexistspacesXn⊂D(AF)withdimXn=dsuchthatlimn→∞supx∈XnkxkH=1k(AF−λ)xkH=0.Inparticularweobtainthatlimx→∞supx∈Xnk(B+λ)(Λ−x−LλΛ+x)kH
kxkH=0andsince(B+λ)−1isaboundedoperatorlimx→∞supx∈XnkΛ−x−LλΛ+xkH
kxkH=0.WecanconcludethatthereexistsanN∈NsuchthatkΛ−x−LλΛ+xkH≤kxkH
2forallx∈Xnwithn≥N.IntheremainderweassumewithoutlossofgeneralitythatN=1.Notethat0=limn→∞supx∈XnkxkH=1k(AF−λ)xkH=limn→∞supx∈XnkxkH=1supy∈HkykH=1|h(AF−λ)x,yiH|=limn→∞supx∈XnkxkH=1supy∈FkykH=1|[dQλΛ+x](Λ+y)−h(B+λ)(Λ−x−LλΛ+x),(Λ−y−LλΛ+y)iH|andinparticular0=limn→∞supx∈XnkxkH=1supy∈(1+Lλ)F+kykH=1|h(AF−λ)x,yiH|=limn→∞supx∈XnkxkH=1supy+∈F+ky+kλ=1|[dQλ(Λ+x)](y+)|.NowletX+n:=Λ+Xn⊂F+.Ifx∈XnisanelementofF−,thenby(16)h(AF−λ)x,xiH≤(λ0−λ)kxk2
Handthus|h(AF−λ)x,xiH|≥(λ−λ0)kxk2
H
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP21whichisacontradictiontothedeﬁnitionofXn.ThusdimX+n=d.Furthermoreforx∈XnbyanapplicationofthelowertriangleinequalitykΛ+xkλ=kΛ+x++LλΛ+xkH=kx−(Λ−x−LλΛ+x)kH≥kxkH
2.Asaconsequencelimn→∞supx+∈X+nkx+kλ=1supy+∈F+ky+kλ=1|[dQλ(x+)](y+)|=0andthusalsolimn→∞supx+∈X+nkx+kλ=1

dQλx+

F0+,λ=0.(17)ThisimpliesthatzeroisinthespectrumofQλbyageneralisedversionofWeyl’scriterion,whichcanforexamplebefoundin[21].Toprovethistakingintoac-countthemultiplicityd,weletε>0andletPbethespectralmeasureofQλ.IfdimranP((−ε,ε))≤d−1thenwecanﬁndasequenceofxn∈X+nwithkxnkλ=1andP((−ε,ε))xn=0.Usingtheembeddingj=jG0+→F0+◦iG+→G0+wecancomputethatforanyx∈F+

dQλx

F0+,λ=

(dQλ+κλj)−1dQλx

F+,λ=

(Qλ+κλ)1/2(x−κλ(dQλ+κλj)−1j(x))

λ=

(Qλ+κλ)1/2(x−κλ(Qλ+κλ)−1x)

λ=

Qλ(Qλ+κλ)−1/2x

λ.Togetherwiththespectraltheoremweobtainthat

dQλxn

2
F0+,λ=Z−ε−κλ+1t2
t+κλdPxn,xn(t)+Z∞εt2
t+κλdPxn,xn(t)≥ε2
ε+κλkxk2
λwhichisacontradictionto(17).Itremainstoprovethat0=µk(Qλ),forsomek∈N.Sinceλ<sup`≥1λ`thereexistsan`∈Nsuchthatλ<λ`.Bydeﬁnitionµ`(Qλ`)=0andthusforanysubspaceV⊂F+ofdimension`thereexistsanxV
+∈Vsuchthatqλ`(x+,x+)≥−εkx+k2
λ.ByLemma7weobtainthatqλ(xV
+,xV
+)≥qλ`(xV
+,xV
+)≥−ε

xV
+

2
λ`≥−ε

xV
+

2
λ.Thisimpliesthatµ`(Qλ)≥0,andconsequently0=µk(Qλ)forsomek<`.Wecanconcludethatλ=λk,whichcompletestheproofofTheorem1.
22LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS3.ApplicationtotheDirac–CoulombOperatorLetH0=−iα·∇+βbethefreeDiracoperatorwhereα1,α2,α3,β∈C4×4withαiαj+αjαi=2δijIC4,αiβ+βαi=0,β2=IC4.Wechoosetherepresentationαi= 0σiσi0!,β= IC200−IC2!.ThefreeDiracoperatorisessentiallyself-adjointonC∞0(R3;C4).TheDirac–CoulomboperatorHν=H0−ν/|x|issymmetriconD(Hν)=C∞0(R3;C4)⊂L2(R3;C4)=:H.LetΛ±betheTalmanprojections,Λ+ ϕψ!= ϕ0!,Λ− ϕψ!= 0ψ!.ThenclearlyΛ±D(Hν)⊂D(Hν)andthustheﬁrstassumptionofTheorem1issatis-ﬁed.Wefurthercomputethatλ0=supψ∈C∞0(R3;C2)\{0}RR3(−1−ν
|x|)|ψ(x)|2dx kψk2
H=supx∈R3(−1−ν/|x|)=−1.Dolbeault,Esteban,LossandVega[8]provedtheHardyinequalityZR3|σ·∇ψ(x)|2
1+1
|x|dx+ZR3 1−1
|x|!|ψ(x)|2dx≥0(18)forallψ∈H1(R3;C2)byanalyticmethods.Followingsimilarcomputationsin[12,11]wecanuse(18)toprovethatq0(ψ,ψ)≥0forallψ∈C∞0(R3;C2)andallν∈[0,1].HereqEistheSchurcomplementqE(ψ,ψ)=ZR3 1−ν
|x|−E!|ψ(x)|2dx+ZR3|σ·∇ψ(x)|2
1+ν
|x|+Edx.AsaconsequenceofLemma7(iii)weobtainλ1≥0>λ0.Notethatthisstatementcanalsobeprovedbymeansofanabstractcontinuationprinciple[9,Section3]andcantheninturnbeusedtoestablishtheHardyinequality(18)[9,Section4].Sinceinaddition−1−ν/|x|isessentiallyself-adjointonC∞0(R3;C2),alltheconditionsofTheorem1aresatisﬁedandthusforanyν∈[0,1]thereexistsaself-adjointextensionofHνwitheigenvaluesgivenbyλk=infV⊂C∞0(R3;C2)dimV=ksupψ∈(V×C∞0(R3;C2))\{0}hψ,HνψiH
kψk2
H.Theself-adjointextensioncoincideswithextensionconstructedin[12]andthusforν<1alsowiththeextensionsofSchmincke[27],Wüst[31,32]andNenciu[24](whichwereallprovedtobeequalbyKlausandWüst[18]).Thevariationalprincipleforthisdistinguishedextensionisthesameastheoneobtainedin[11].
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP23Toestablishasecondvariationalprinciple,wecanchooseΛ±tobespectralprojec-tionsofthefreeDiracoperator,Λ+=PH0[0,∞),Λ−=PH0(−∞,0).LetHνagaindenotetheDiracoperatorwithCoulombpotentialactingonthedomainD(Hν)=F+⊕F−⊂L2(R3;C4)=:HwithF±=Λ±C∞0(R3;C4).TheoperatorHνissymmetricandtheﬁrstassumptionofTheorem1issatisﬁed.Againwecancomputeλ0tobeλ0=supψ∈F−\{0}hψ,(−√
1−∆−ν/|x|)ψiH
kψk2
H≤supx∈R3(−1−ν/|x|)=−1.Usinganabstractcontinuationprinciple,itwasprovedin[9]thatλ1≥0>λ0forν∈[0,1).Toextendthisresulttotheendpointν=1wenotethatbytheaboveforanyν∈[0,1)theSchurcomplementqEqE(ψ,ψ)=*ψ,√
1−∆−ν
|x|−Eψ+H+*Λ−ψ
|x|,Λ−√
1−∆+ν
|x|+EΛ−−1Λ−ψ
|x|+Hsatisﬁesq0(ψ,ψ)≥0forallψ∈Λ+H1/2(R3).Takingthelimitν→1oneobtains(see[11,Lemma15])theanalogueof(18)*ψ,√
1−∆−1
|x|ψ+H+*Λ−ψ
|x|,Λ−√
1−∆+1
|x|Λ−−1Λ−ψ
|x|+H≥0.IncontrasttothecaseoftheTalmanprojections,wearenotawareofananalyticproofofthisinequality.ByLemma7(iii)thisinequalityprovesthatλ1≥0>λ0stillholdsintheendpointcaseν=1.Asdiscussedintheappendix,theoperatorΛ−(√
1−∆+ν/|x|)|F−isessentiallyself-adjointonΛ−C∞0(R3;C4)⊂Λ−L2(R3;C4).ThusalltheconditionsofTheorem1aresatisﬁedandforanyν∈[0,1]weobtainaaself-adjointextensionofHνwitheigenvaluesgivenbyλk=infV⊂Λ+C∞0(R3;C4)dimV=ksupψ∈(V⊕Λ−C∞0(R3;C4))\{0}hψ,HνψiH
kψkH.Thisisthesameresultasin[11].4.Self-adjointextensionsandtheAPSboundaryconditionWeconsideranoperatoroftheformA=σ(∂x+B)(19)actingonfunctionsinH=L2(R−;K)withKbeingacomplexHilbertspace.TheoperatorBisdenselydeﬁnedonadomainD(B)⊂Kanddoesnotdependonx.Itisself-adjointandhasdiscretespectrum.ThemapσisanautomorphismonK,equallyindependentofx.Furthermore,wemakethefollowingassumptions:
24LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS(i)σ2=−I,σ∗=−σ,(ii){B,σ}=Bσ+σB=0,(iii)dimkerB<∞,kerB=N+⊕N−forsomesubspacesN±⊂D(B)and(iv)σ(N−)=N+.OnthedomainD(A)=C10(R−;D(B)),Aisawell-deﬁnedsymmetricoperator.Conti-nuityanddiﬀerentiabilityonthesetD(A)aredeﬁnedwithrespecttothegraphnormofBonD(B).Remark12.•AsanexamplewecouldtakeKtobetheHilbertspaceofsquareintegrablefunctionsonS1,parametrisedbyavariabley.IfthenBisdeﬁnedonthecon-tinuouslydiﬀerentiableperiodicfunctionsasB=iσ3∂yandσ=iσ2,withthePaulimatricesσi,wemayidentifyAwiththeDiracoperatorontheinﬁnitelylongcylinderboundedfromoneside.•Moregenerally,wemaythinkofBbeinganyﬁrst-orderdiﬀerentialoperatoronsomeclosedmanifoldΣsuchthatArepresentsadiﬀerentialoperator(ofﬁrstorder)onageneralisedcylinder.Infact,anyﬁrst-orderellipticoperatoronacompactmanifoldMwithboundaryΣtakesaformasgivenin(19)onacollarneighbourhoodoftheboundary[1],butBandσarenotnecessarilyindependentofx.Thesespecialcasesareincludedbutwedonotrestrictourselvestothem.Fromassumptions(i),(iii)and(iv)weconcludethatthekernelofBisofevendimension,hencedimkerB=2N0forsomeN0∈N0.Also,thevanishinganticom-mutator{B,σ}impliesthatσmapselementsfromthepositivespectralsubspaceofBtothenegativespectralsubspaceandviceversa.Itiswell-known(cf.e.g.[10],[15])thatA|C10(R−;D(B))hasaself-adjointextensioncharacterisedbyanon-localboundaryconditionknownasthe‘Atiyah–Patodi–Singerboundarycondition’.LetusdenotebyP+B>0theprojectionontothesumofN+andthepositivespectralsubspaceofB.Thenitholds:Proposition13(APS).TheoperatorAfrom(19)isself-adjointonthedomainD(AAPS)={f∈L2(R−;D(B))∩H1(R−;K)|P+B>0f|x=0=0}.(20)Alsohere,D(B)andKaretobeunderstoodasHilbertspaceswiththeirrespectivenorms,inthesensethatf∈D(AAPS)isafunctionsuchthatkfk,kBfkandk∂xfkareallsquare-integrable.Aproofofthispropositionfollowsbyexplicitcalculation.Startingfromthesymmetricoperatorin(19)wecanshowthatitfallsintotheclassofgappedoperatorsforwhichourconstructionofaself-adjointextensionapplies.Indeedweﬁnd:
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP25Theorem14.Theorem1appliestotheoperatorA=σ(∂x+B)deﬁnedonC10(R−;D(B)).Theself-adjointextensionconstructedinthiswaycoincideswiththeAtiyah–Patodi–SingerextensionAAPSfromProposition13.InthissensetheextensionfromTheorem1ischaracterisedbytheglobalboundaryconditionsfrom(20).WeproveTheorem14intheremainderofthissection.First,ifwewriteK±=PB≷0K⊕N±andH±=L2(R−;K±)thenH=H+⊕H−isanorthogonaldecompositionofHwithcorrespondingorthogonalprojectionsΛ±suchthatF±=Λ±D(A)=C10(R−;K±∩D(B))⊂D(A).Itiseasytoseethatσ(K+)=K−andσ(K−)=K+.WecanconcludethatΛ−AΛ−=0andhenceΛ−AΛ−isessentiallyself-adjointonC10(R−;D(B))andλ0=0.Let`k,k∈Z\{0}betheeigenvaluesofBsuchthat`k≥`k0ifk>k0and`k=0for−N0≤k≤N0.WedenotethecorrespondingeigenvectorsofBbyϕkandassumetheyarechosensuchthatσϕk=ϕ−kandσϕ−k=−ϕk.Anyfunctionu∈F+hasthenanexpansionu(x)=Xk>0uk(x)ϕkwithfunctionsuk∈C10(R−;C)andsimilarlyforv∈F−.Wecanthenwriteforanyu∈F+andanyv∈F−hu+v,σ(∂x+B)(u+v)i=Xk>0Xl<0hukϕk,σ(∂x+B)vlϕli+complexconjugate=Xk>0Xl<0hukϕk,(∂x−B)vlϕ−li+c.c.=Xk,l>0hukϕk,(∂x−`l)v−lϕli+c.c.=Xk>0huk,(∂x−`k)v−ki+c.c..Hence,wecanrewritetheexpectationvalueofAasPk>0huk,(∂x−`k)v−ki+c.c.
kuk2+kvk2=Xk>0:uk6=0orv−k6=0kukk2+kv−kk2
kuk2+kvk2huk,(∂x−`k)v−ki+c.c.
kukk2+kv−kk2whichwecanidentifywithanarithmeticmeanofexpectationvaluesforsinglek’sweightedaccordingtotheirnorm.Clearly,thisexpressionisboundedbysupk>0:uk6=0orv−k6=0huk,(∂x−`k)v−ki+c.c.
kukk2+kv−kk2.Takingthesupremumoverv∈F−ﬁnallygivesanupperboundsupv∈F−hu+v,A(u+v)i kuk2+kvk2≤supk>0supv−k6=0huk,(∂x−`k)v−ki+c.c.
kukk2+kv−kk2
26LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
andindeed,equalityholdssincewecanalwaysﬁndasequenceofvapproximatingtheright-hand-side.Clearly,weﬁndthesupremumoverv−kbychoosingv−k=λ(−∂x−`k)ukforsomerealnumberλ.Maximisingoverallvaluesofλweﬁndλ=kukk k(−∂x−`k)ukkandhencesupv∈F−hu+v,A(u+v)i kuk2+kvk2=supk>0:uk6=0k(−∂x−`k)ukk kukk.Notethattheleft-handsideispreciselyE(u)asdeﬁnedinLemma11.Byconstructionthesupremumisachievedatv=LE(u)uwhichcoincideswiththerelationv−k=λ(−∂x−`k)ukaboveforλ=E(u)−1.Takingtheinﬁmumoverallu6=0wethenobtainλ1=infu∈C10(R−;K+∩D(B))u6=0supk>0:uk6=0"`2
k+k−∂xukk2
kukk2#1
2=π>0(21)ifthereisan`k=0.WehaveusedthatbythevariationalprinciplefortheFriedrichsextensionoftheLaplaceoperatorthelasttermgivesthelowesteigenvalueoftheDirichletLaplacian.IfkerB={0},thenλ1>π.Inbothcases,Aisagappedoperatorandallassumptionsforconstructingaself-adjointextensionasinTheorem1aresatisﬁed.ForcomparisonwiththeAPS-extensionweareinterestedinthedomainofAF,thatisinparticularinhowtheHilbertspaceF+appearsinthissetting.RecallthatF+istheclosureofF+inthenormk·kF+,EconstructedfromthequadraticformqEandthegraphnormoftheoperatorLEkuk2
F+,E=qE(u,u)+κEkLEuk2
H,whereinoursettingthequadraticformqE(u,u)=−Ekuk2+1
Ek∂xuk2+kBuk2andkLEuk2=1
E2k∂xuk2+kBuk2.Usingthatλ2
1kuk2≤k∂xuk2+kBuk2,whichfollowsfromcomparisonwith(21),itisdirectlyseenthatqE(u,u)≤(λ1−E)kuk2+(1
E−1
λ1)(k∂xuk2+kBuk2)andhencek·kF+,Eisequivalenttothesumofnormsk·k+k∂x·k+kB·kaslongasE<λ1.ClosingF+inthisnormgivestheHilbertspaceL2(R−;D(B)∩K+)∩H10(R−;K+),ascanbeseenfromthefollowingargument.Givenanf∈L2(R−;D(B)∩K+)∩H10(R−;K+),weusethestandardapproximationbysmooth,compactlysupportedfunctionsbymolliﬁcationofafunctioninH10(R−;K+),whichallowstoconstructasequenceoffn∈C10(R−;D(B)∩K+)thatconvergestofinthenormk·k+k∂x·k,seee.g.[14,Chapter5.5].SinceBf∈L2(R−;K+),itmaybeapproximatedinthesamefashionsuchthatthesequence{fn}willalsoconvergeink·k+kB·k.TheseconsiderationsshowthatindeedD(AAPS)⊂F+⊕H−andthusAAPScoincideswithAFbytheuniquenesspropertyprovedinTheorem1.
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP27AppendixA.EssentialSelf-AdjointnessoftheBrown–RavenhallOperatorLetH0betheself-adjointfreeDiracoperatorwithdomainH1(R3;C4)⊂L2(R3;C4)anddenotebyΛ±theprojectionsontothepositive/negativespectralsubspaceofH0.Forγ∈RtheBrown–Ravenhalloperator[5]isdeﬁnedasBγ=Λ+(H0−γ/|x|)Λ+=Λ+(√
1−∆−γ/|x|)Λ+.ontheHilbertspaceΛ+L2(R3;C4).Foracomprehensivereviewwerefertothetext-bookofBalinskyandEvans[2].Whilethephysicallyrelevantcaseisγ>0,weareinterestedinthecasewhereγ=−ν∈[−1,0].Forγ<3/4theoperatorBγwasprovedtobeself-adjointonΛ+H1(R3;C4)byTix[29].SinceΛ+H1(R3;C4)⊂H1(R3;C4)weobtainfromHardy’sinequality

Λ+γ
|x|Λ+ψ

L2(R3;C4)≤2γ

1
2|x|Λ+ψ

L2(R3;C4)≤2γk∇Λ+ψkL2(R3;C4)foranyψ∈H1(R3;C4).ToprovethatB−νisessentiallyself-adjointonΛ+C∞0(R3;C4),itthussuﬃcestoprovethestatementforB0.ThisisanimmediateconsequenceofthefactthatthefreeDiracoperatorH0isessentiallyself-adjointonC∞0(R3;C4).WecanalsoconcludethattheoperatorΛ−(√
1−∆+ν/|x|)Λ−isessentiallyself-adjointonΛ−C∞0(R3;C4)sinceitisunitarilyequivalenttotheBrown–Ravenhallop-eratorB−νviathetransformU:L2(R3;C4)→L2(R3;C4)"U ψ1ψ2!#(x)= ψ2(−x)ψ1(−x)!.References[1]M.F.Atiyah,V.K.PatodiandI.M.Singer,SpectralasymmetryandRiemannianGeometry.I,Math.Proc.Camb.Phil.Soc.(1975),77,43[2]A.A.BalinskyandW.D.Evans,Spectralanalysisofrelativisticoperators,ImperialCollegePress,London,2011[3]M.Sh.BirmanandM.Z.Solomjak,SpectraltheoryofselfadjointoperatorsinHilbertspace.Translatedfromthe1980RussianoriginalbyS.KhrushandV.Peller.MathematicsanditsApplications(SovietSeries).D.ReidelPublishingCo.,Dordrecht,1987.[4]J.G.BrascheandH.Neidhardt,SomeremarksonKre˘ın’sextensiontheory,Math.Nachr.165(1994),159–181.[5]G.E.BrownandD.G.Ravenhall.OOntheinteractionoftwoelectrons,Proc.Roy.Soc.LondonSer.A.,208(1951),552–559.[6]S.N.DattaandG.Deviah,TheminimaxtechniqueinrelativisticHartree-Fockcalculations,Pramana30(5)(1988),387–405[7]E.B.DaviesSpectraltheoryanddiﬀerentialoperators,CambridgeStudiesinAdvancedMathe-matics,42.CambridgeUniversityPress,Cambridge,1995.[8]J.Dolbeault,M.J.Esteban,M.Loss,andL.Vega,AnanalyticalproofofHardy-likeinequalitiesrelatedtotheDiracoperator,J.Funct.Anal.216(2004),no.1,1–21.[9]J.Dolbeault,M.J.EstebanandE.Séré,OntheEigenvaluesofOperatorswithGaps.ApplicationtoDiracOperators,J.Funct.Anal.174(2000)208–226.
28LUKASSCHIMMER,JANPHILIPSOLOVEJ,ANDSABIHATOKUS
[10]R.G.DouglasandK.P.Wojciechowski,AdiabaticLimitoftheη-Invariants.TheOdd-DimensionalAtiyah-Patodi-SingerProblem,Commun.Math.Phys.142(1991),139–168.[11]M.J.Esteban,M.LewinandE.Séré,DomainsforDirac-Coulombmin-maxlevels,toappearinRev.Mat.Iberoam,arXiv:1702.04976(2017).[12]M.J.EstebanandM.Loss,Self-adjointnessforDiracoperatorsviaHardy-Diracinequalities,J.Math.Phys.48(2007),no.11,112107.[13]M.J.EstebanandM.Loss,Self-adjointnessviapartialHardy-likeinequalities,Mathematicalresultsinquantummechanics,WorldSci.Publ.,Hackensack,NJ,2008,41–47.[14]L.C.EvansPartialDiﬀerentialEquation,SecondEdition,GraduateStudiesinMathematics,Vol.19,AmericanMathematicalSociety,Providence,RhodeIsland.[15]K.FurutaniAtiyah–Patodi–SingerboundaryconditionandasplittingformulaofaspectralﬂowJ.Geom.Phys.56(2006),310–321.[16]K.Friedrichs,SpektraltheoriehalbbeschränkterOperatorenundAnwendungaufdieSpektralzer-legungvonDiﬀerentialoperatoren,Math.Ann.109(1934),no.1,465–487.[17]M.GriesemerandH.Siedentop,Aminimaxprinciplefortheeigenvaluesinspectralgaps,J.LondonMath.Soc.(2)60(1999),no.2,490-500.[18]M.KlausandR.Wüst,CharacterizationanduniquenessofdistinguishedselfadjointextensionsofDiracoperators,Comm.Math.Phys.64(1978/79),no.2,171–176.[19]M.Kraus,M.LangerandC.Tretter,Variationalprinciplesandeigenvalueestimatesforun-boundedblockoperatormatricesandapplications,J.Comput.Appl.Math.171(2004),no.1-2,311–334.[20]M.G.Krein,TheTheoryofSelf-adjointExtensionsofSemiboundedHermitianOperatorsanditsApplications,I.Mat.Sbornik20(1947),no.3,431–495(inRussian).[21]D.KrejčířikandZ.Lu,Locationoftheessentialspectrumincurvedquantumlayers,J.Math.Phys.55(2014),no.8,083520.[22]S.MorozovandD.Müller,OntheminimaxprincipleforCoulomb-Diracoperators,Math.Z.,280(2015),733–747.[23]D.Müller,Minimaxprinciples,Hardy-DiracinequalitiesandoperatorcoresfortwoandthreedimensionalCoulomb-Diracoperators,Doc.Math.21(2016),1151–1169.[24]G.Nenciu,Self-adjointnessandinvarianceoftheessentialspectrumforDiracoperatorsdeﬁnedasquadraticforms,Comm.Math.Phys.48(1976),no.3,235–247.[25]M.ReedandB.Simon,Methodsofmodernmathematicalphysics.I.Functionalanalysis.Aca-demicPress,NewYork-London,1972.[26]M.ReedandB.Simon,Methodsofmodernmathematicalphysics.II.Fourieranalysis,self-adjointness.AcademicPress,NewYork-London,1975.[27]U.-W.Schmincke,EssentialselfadjointnessofDiracoperatorswithastronglysingularpotential,Math.Z.126(1972),71–81.[28]J.D.Talman,MinimaxprinciplefortheDiracequation,Phys.Rev.Lett.57(1986),no.9,1091–1094.[29]C.Tix,Self-adjointnessandspectralpropertiesofapseudo-relativisticHamiltonianduetoBrownandRavenhall,Preprint,mp-arc:97-441,1997.[30]C.Tretter,Spectraltheoryofblockoperatormatricesandapplications,ImperialCollegePress,London,2008.[31]R.Wüst,Distinguishedself-adjointextensionsofDiracoperatorsconstructedbymeansofcut-oﬀpotentials,Math.Z.141(1975),93–98.[32]R.Wüst,Diracoperationswithstronglysingularpotentials.Distinguishedself-adjointextensionsconstructedwithaspectralgaptheoremandcut-oﬀpotentials,Math.Z.152(1977),no.3,259–271.
FRIEDRICHSEXTENSIONANDMIN-MAXPRINCIPLEINAGAP29LukasSchimmer,QMATH,DepartmentofMathematicalSciences,UniversityofCopenhagen,Universitetsparken5,DK-2100CopenhagenØ,DenmarkE-mailaddress:schimmer@math.ku.dkJanPhilipSolovej,QMATH,DepartmentofMathematicalSciences,UniversityofCopenhagen,Universitetsparken5,DK-2100CopenhagenØ,DenmarkE-mailaddress:solovej@math.ku.dkSabihaTokus,QMATH,DepartmentofMathematicalSciences,UniversityofCopen-hagen,Universitetsparken5,DK-2100CopenhagenØ,DenmarkE-mailaddress:sabiha.tokus@math.ku.dk

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

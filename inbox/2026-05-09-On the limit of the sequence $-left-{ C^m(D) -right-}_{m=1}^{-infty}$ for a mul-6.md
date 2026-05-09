---
source: "https://arxiv.org/abs/2211.10978v1"
title: "On the limit of the sequence $\\left\\{ C^m(D) \\right\\}_{m=1}^{\\infty}$ for a multipartite tournament $D$"
author: "Ji-Hwan Jung, Suh-Ryung Kim, Hyesun Yoon"
year: "2022"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/2211.10978v1"
pdf: "https://arxiv.org/pdf/2211.10978v1"
captured_at: "2026-05-09T12:39:57Z"
updated_at: "2026-05-09T12:39:57Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代思想"
status: raw
---

# On the limit of the sequence $\left\{ C^m(D) \right\}_{m=1}^{\infty}$ for a multipartite tournament $D$

- 著者: Ji-Hwan Jung, Suh-Ryung Kim, Hyesun Yoon
- 年: 2022
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/2211.10978v1)
- ダウンロード: https://arxiv.org/pdf/2211.10978v1
- PDF: https://arxiv.org/pdf/2211.10978v1

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

For an integer $k \ge 2$, let $A$ be a Boolean block matrix with blocks $A_{ij}$ for $1 \le i,j \le k$ such that $A_{ii}$ is a zero matrix and $A_{ij}+A_{ji}^T$ is a matrix with all elements $1$ but not both corresponding elements of $A_{ij}$ and $A_{ji}^T$ equal to $1$ for $i \neq j$. Jung~{\em et al.} [Competition periods of multipartite tournaments. {\it Linear and Multilinear Algebra}, https://doi.org/10.1080/03081087.2022.2038057] studied the matrix sequence $\{A^m(A^T)^m\}_{m=1}^{\infty}$. This paper, which is a natural extension of the above paper and was initiated by the observation that $\{A^m(A^T)^m\}_{m=1}^{\infty}$ converges if $A$ has no zero rows, computes the limit of the matrix sequence $\{A^m(A^T)^m\}_{m=1}^{\infty}$ if $A$ has no zero rows. To this end, we take a graph theoretical approach: noting that $A$ is the adjacency matrix of a multipartite tournament $D$, we compute the limit of the graph sequence $\left\{ C^m(D) \right\}_{m=1}^{\infty}$ when $D$ has no sinks.

## PDF Text

arXiv:2211.10978v1 [math.CO] 20 Nov 2022
Onthelimitofthesequence{Cm(D)}∞
m=1foramultipartitetournamentDJi-HwanJung1,Suh-RyungKim2,andHyesunYoon21DepartmentofMathematicsEducation,ChinjuNationalUniversityofEducation,Jinju526732DepartmentofMathematicsEducation,SeoulNationalUniversity,Seoul08826jhjung@cue.ac.kr,srkim@snu.ac.kr,magisakura@snu.ac.krAbstractForanintegerk≥2,letAbeaBooleanblockmatrixwithblocksAijfor1≤i,j≤ksuchthatAiiisazeromatrixandAij+AT
jiisamatrixwithallelements1butnotbothcorrespondingelementsofAijandAT
jiequalto1fori6=j.Jungetal.[Competitionperiodsofmultipartitetournaments.LinearandMulti-linearAlgebra,https://doi.org/10.1080/03081087.2022.2038057]studiedthematrixsequence{Am(AT)m}∞
m=1.Thispaper,whichisanaturalextensionoftheabovepaperandwasinitiatedbytheobservationthat{Am(AT)m}∞
m=1convergesifAhasnozerorows,computesthelimitofthematrixsequence{Am(AT)m}∞
m=1ifAhasnozerorows.Tothisend,wetakeagraphtheoreticalapproach:notingthatAistheadjacencymatrixofamultipartitetournamentD,wecomputethelimitofthegraphsequence{Cm(D)}∞
m=1whenDhasnosinks.Keywords.m-stepcompetitiongraph,multipartitetournament,limitofaBooleanmatrixsequence,limitofagraphsequence,indexofimprimitivity,setsofimprimitivity2010MathematicsSubjectClassiﬁcation.05C20,05C751Introduction
Theunderlyinggraphofeachdigraphinthispaperisassumedtobesimpleunlessother-wisementioned.Weusethenotationu→vfor“thereisanarc(u,v)”andthenotationX→Yfor“thereisanarc(x,y)foreachx∈Xandforeachy∈Y”inthegivendigraph.Forsimplicity,weomitthebracesifXorYisasingletonset.GivenadigraphDandapositiveintegerm,avertexyisanm-steppreyofavertexxifandonlyifthereexistsadirectedwalkfromxtoyoflengthm.Givenapositiveintegerm,them-stepcompetitiongraphofadigraphD,denotedbyCm(D),hasthesamevertex1
setasDandhasanedgebetweenverticesuandvifandonlyifthereexistsanm-stepcommonpreyofuandvinD.Thenotionofanm-stepcompetitiongraphintroducedbyChoetal.[6]isageneralizationofthecompetitiongraphintroducedbyCohen[7](thecompetitiongraphofadigraphDisC1(D)).Sinceitsintroduction,anm-stepcompetitiongraphhasbeenextensivelystudied(see,forexample,[1,9,10,14,16]).Forthetwo-elementBooleanalgebraB={0,1},Bndenotesthesetofalln×nmatricesoverB.UndertheBooleanoperations(1+1=1,0+0=0,1+0=1,1×1=1,0×0=0,1×0=0),matrixadditionandmultiplicationarestillwell-deﬁnedinBn.Throughoutthispaper,amatrixisBooleanunlessotherwisementioned.WenotethattheadjacencymatrixofCm(D)foradigraphDofordernisthematrixA∗
mobtainedfromAm(AT)mbyreplacingeachofdiagonalelementswith0whereAistheadjacencymatrixofD.Toseewhy,wetaketwodistinctverticesuandvofDandsupposethattheithrowandthejthrowofAaretherowscorrespondingtouandv,respectively.ThenuandvareadjacentinCm(D)⇔uandvhaveanm-stepcommonpreyinD⇔theinnerproductoftheithrowandthejthrowofAmis1⇔the(i,j)-entryofA∗
mis1.ItiseasytocheckthatA∗
i=A∗
jifandonlyifAi(AT)i=Aj(AT)jfora(0,1)BooleanmatrixAofordernandanypositiveintegersiandj.Therefore,foradigraphDanditsadjacencymatrixA,Ci(D)=Cj(D)⇔Ai(AT)i=Aj(AT)j(1)foranypositiveintegersiandj.Ak-partitetournamentisanorientationofacompletek-partitegraphforapositiveintegerk.Theadjacencymatrixofak-partitetournamentfork≥2isrepresentedasablockmatrixAwithblocksAijfor1≤i,j≤ksuchthatAiiisazeromatrixandAij+AT
jiisamatrixwithallelements1butnotbothcorrespondingelementsofAijandAT
jiequalto1fori6=j(seeDandAinFigure1foranexample).Wecallak-partitetournamentamultipartitetournamentifk≥2.GivenaBooleanmatrixAofordern,considerthematrixsequence{Am(AT)m}∞
m=1.SincethecardinalityoftheBooleanmatrixsetBnisequalto2n2whichisﬁnite,thereisthesmallestpositiveintegerqsuchthat,forsomepositiveintegerr,Aq+i(AT)q+i=Aq+r+i(AT)q+r+iforeverynonnegativeintegeri.ThenthereisalsothesmallestpositiveintegerpsuchthatAq(AT)q=Aq+p(AT)q+p.ThoseintegersqandparecalledthecompetitionindexandthecompetitionperiodofA,respectively,whichChoandKim[5]introduced.IfDisthedigraphwhoseadjacencymatrixisA,then,by(1),qandparethesmallestintegerssuchthat,forsomepositiveintegerr,Cq+i(D)=Cq+r+i(D)forallnonnegativeintegersiandCq(D)=Cq+p(D),respectively.Inthisrespect,qandparecalledthecompetitionindexandthecompetitionperiodofD,respectively,anddenoted2
v2
v3
v4
v5
v6
v1
D
011111
000001
000001
000010
011000
000100



















v1v2v3v4v5v6v1v2v3v4v5v6A
Figure1:AtripartitetournamentDanditsadjacencymatrixAwhereV1={v1},V2={v2,v3,v4},andV3={v5,v6}arethepartitesetsofD.bycindex(D)andcperiod(D),respectively,whichChoandKim[4]introduced.Referto[12,13]forsomeresultsofcompetitionindicesandcompetitionperiodsofdigraphs.Jungetal.[11]computedthecompetitionperiodofamultipartitetournament.Theorem1.1([11]).Thecompetitionperiodofak-partitetournamentDisatmostthreeforanyintegerk≥2.Further,ifk=2,thenthecompetitionperiodofDisatmosttwo.Asweobservedabove,theyactuallycomputedthecompetitionperiodofaBooleanblockmatrixAwithblocksAijfor1≤i,j≤ksuchthatAiiisazeromatrixandAij+AT
jiisamatrixwithallelements1butnotbothcorrespondingelementsofAijandAT
jiequalto1fori6=j.Inthispaper,wecomputethelimitofthematrixsequence{Am(AT)m}∞
m=1,ifitexists,asanaturalextensionoftheresultobtainedbyJungetal.[11].Tothisend,weinvestigatethelimitof{Cm(D)}∞
m=1forthedigraphDofA,whichisamultipartitetournament.ThegreatestcommondivisorofthelengthsofthecloseddirectedwalksofastronglyconnectednontrivialdigraphDiscalledtheindexofimprimitivityofDanddenotedbyκ(D).Ifκ(D)=1,thenDissaidtobeprimitive.IfAistheadjacencymatrixofaprimitivedigraphD,thenthereexistsapositiveintegerMsuchthatAmbecomesamatrixofall1sforanyintegerm≥M.WecallthesmallestMtheexponentofDanddenoteitbyexp(D).ThevertexsetofDcanbepartitionedintoκ(D)nonemptysubsetsU1,U2,...,Uκ(D)whereeacharcofDgoesoutfromUiandentersUi+1(identifyUκ(D)+1withU1)forsomei∈{1,...,κ(D)}(see[3]).WecallthesetsU1,U2,...,Uκ(D)thesetsofimprimitivityofD.SupposethatDisaweaklyconnecteddigraph.ThenthestrongcomponentsofDcanbearrangedasQ1,...,QssothatthereisnoarcgoingfromQitoQjwhenever3
i>j.WecallthestrongcomponentsarrangedinsuchawayorderedstrongcomponentsofD.Forconvenience,wecallQsthelaststrongcomponentofD.WedenotethesetsofimprimitivityofQibyU(i)1,U(i)2,...,U(i)κ(Qi)foreachintegeri=1,...,s.ForthedigraphDgiveninFigure1,thesubdigraphsQ1andQ2inducedbyV1andV2∪V3,respectively,formtheorderedstrongcomponentsofD.Itiseasytocheckthatκ(Q2)=4,U(2)1={v2,v3},U(2)2={v6},U(2)3={v4},andU(2)4={v5}.Further,111111
111000
111000
100100
100010
100001






















Am(AT)m=foranypositiveintegermwhereAistheadjacencymatrixofD.Thereforetheabovematrixisthelimitofthematrixsequence{Am(AT)m}∞
m=1.WenotethatitcorrespondstoM2inFigure2whereJ(2)andJ(3)donotexist.ThisisnotacoincidenceasshownbyCorollary1.3.Throughoutthispaper,wetakeagraphtheoreticalapproachtoderivethefollowingresultwhosematrixversionisCorollary1.3.
Theorem1.2.Supposethatak-partitetournamentDwithk-partition(V1,V2,...,Vk)hasorderedstrongcomponentsQ1,Q2,...,Qsforsomeintegersk≥2ands≥1.Then,ifQsisnontrivial,thegraphsequence{Cm(D)}∞
m=1convergestoagraphG∼=







K|V(D)|,ifκ(Qs)=1;G1,ifκ(Qs)=2;G3,ifκ(Qs)=3;G2,ifκ(Qs)=4,whereG1,G2,andG3arethegraphsgiveninFigure2.ByProposition2.1,QsbeingnontrivialisequivalenttoDhavingnosinks.ThereforeQsbeingnontrivialisequivalenttotheadjacencymatrixofDhavingeachrowsumnonzero.By(1),Theorem1.2mayberestatedasfollows.
Corollary1.3(Matrixversion).Foranintegerk≥2,letAbeablockmatrixwithblocksAijfor1≤i,j≤ksuchthatAiiisazeromatrixandAij+AT
jiisamatrixofall1sbut4
K(1)
K(3)
K(2)
J(1)JJJJ(2)OJOJ(3)











G1M1
K(1)
K(3)
K(2)
K(6)
K(7)
K(5)
K(4)
J(1)JJJJJJJJ(2)OJJOOJOJ(3)OOJJJJOJ(4)OOOJJOOJ(5)OOJOJOOJ(6)OJOJOOOJ(7)































G2M2
K(4)
K(3)
K(2)
K(7)
K(6)
K(5)
K(1)
J(1)JJJJJJJJ(2)JJJJOJJJ(3)JJOJJJJJ(4)OJJJJJOJ(5)OOJJOJOJ(6)OJOJJOOJ(7)































G3M3Figure2:ThegraphsG1,G2,andG3inTheorem1.2.Ineachgraph,K(i)eitherdoesnotexistorstandsforaclique.ThelinebetweentwocliquesK(i)andK(j)indicatesthatK(i)∪K(j)formsacliquewhiletheabsenceofalinebetweenK(i)andK(j)meansthatthereisnoedgejoiningavertexinK(i)andavertexinK(j).ThematrixMi=A(Gi)+IwhereA(Gi)istheadjacencymatrixofGiforeachi=1,2,3;I,J,andOareanidentitymatrix,amatrixofall1s,andazeromatrix,respectively,ofanappropriateorder;J(s)eitherdoesnotexistorisasquarematrixofall1s.5
notbothcorrespondingelementsofAijandAT
jiequalto1fori6=j.IfeachrowsumofAisnonzero,thenthelimitofthematrixsequence{Am(AT)m}∞
m=1existsandequalsoneofthematricesgiveninFigure2afterasimultaneouspermutationofrowsandcolumns.TherestofthispaperisdevotedtoprovingTheorem1.2.InSection2,wejustifythatitissuﬃcienttoconsiderthecasesκ(Qs)=1,2,3,4whereQsisthelaststrongcomponentofthegivenmultipartitetournament.Section3takescareofthecasesκ(Qs)=1,2,4whileSection4takescareofthecaseκ(Qs)=3.Finally,Section5windsuptheproofofTheorem1.2.
2Preliminaries
Proposition2.1.LetDbeamultipartitetournamentwiththelaststrongcomponentQsforsomepositiveintegers.ThenDhasnosinksifandonlyifQsisnontrivial.Proof.IfQsistrivial,thenthevertexinQsisasinkinD.Toshowtheconverse,supposethatQsisanontrivialstrongcomponentofDandtakeavertexv∈V(D).ThenthereareatleasttwopartitesetsofDthatintersectwithV(Qs).ThereforethereexistsapartitesetXsuchthatX∩V(Qs)6=∅andvdoesnotbelongtoX.Thenvhasanout-neighborinX∩V(Qs)sinceDisamultipartitetournament.HenceDhasnosinks.
ThefollowingresultgivenbyEohetal.[8]wasoriginallyforabipartitetournament.Yet,theirproofisstillvalidforageneraldigraphandsothestatementmayberestatedasfollows.
Proposition2.2([8]).LetDbeadigraphwithoutsinks.IftwoverticesareadjacentinCM(D)forapositiveintegerM,thentheyarealsoadjacentinCm(D)foranypositiveintegerm≥M.Thefollowingisanimmediateconsequenceoftheaboveproposition.Corollary2.3.LetDbeadigraphwithoutsinks.Thenthegraphsequence{Cm(D)}∞
m=1converges.Proposition2.1andCorollary2.3tellusthat,foramultipartitetournamentDwiththenontriviallaststrongcomponent,thegraphsequence{Cm(D)}∞
m=1converges.Inthisvein,throughoutthispaper,weseekforthelimitof{Cm(D)}∞
m=1foramultipartitetournamentDwiththenontriviallaststrongcomponent.WenotethatDhasnosinksbyProposition2.1.Thus,byProposition2.2,(∗)inordertoprovethattwoverticesuandvareadjacentinthelimitof{Cm(D)}∞
m=1,itissuﬃcienttoshowthatuandvhaveanm-stepcommonpreyforsomepositiveintegerm.6
Next,weobservethatthelimitof{Cm(D)}∞
m=1dependsontheindexofimprimitivityofthelaststrongcomponentofamultipartitetournamentD.Proposition2.4.LetDbeastronglyconnectedk-partitetournamentforanintegerk≥2.Thentheindexκ(D)ofimprimitivityofDisasfollows:κ(D)=



2or4,ifk=2;1or3,ifk=3;1,ifk≥4.Proof.Itisshownin[2]that,foranintegerk≥3,eachstronglyconnectedk-partitetournamentcontainsadirectedcycleoflength`foreach`∈{3,4,...,k}.Therefore,ifk≥4,thenDisprimitive.Further,ifk=3,thenDcontainsadirectedcycleoflength3andsoκ(D)is1or3.Supposek=2andlet(V1,V2)beabipartitionofD.SinceDisstronglyconnected,thereexistsacloseddirectedwalkW:=u0→u1→···→ul−1→u0containingalltheverticesofDforsomepositiveintegerl.SinceDisabipartitetournament,thelengthlofWisevenandsol≥4.Withoutlossofgenerality,wemayassumethatavertexonWwithanevenindexbelongstoV1andavertexonWwithanoddindexbelongstoV2.SinceDisabipartitetournament,thereisanarcbetweenuiandul−i−1foreachinteger0≤i≤l/2−1.Thenul/2−1→ul/2.Sinceul−1→u0,thereisthesmallestindexiamong1,...,l/2−1suchthatui→ul−i−1.Thenul−i→ui−1andC:=ui−1→ui→ul−i−1→ul−i→ui−1isacloseddirectedwalkoflength4inD.SinceDhasneitherloopsnormultiplearcs,alltheverticesonCaredistinctandsoCisadirectedcycleinD.SinceDisabipartitetournament,thereisnodirectedcycleofoddlengthinDandsoκ(D)is2or4.
Thefollowingisanimmediateconsequenceoftheaboveproposition.Corollary2.5.LetDbeastronglyconnectedk-partitetournamentwithk-partition(V1,V2,...,Vk)foranintegerk≥2.Thenthefollowingaretrue:(i)ifκ(D)=1,thenk≥3;(ii)ifκ(D)=2,thenk=2andthesetsofimprimitivityofDareV1andV2;(iii)ifκ(D)=3,thenk=3andthesetsofimprimitivityofDareV1,V2,andV3;(iv)ifκ(D)=4,thenk=2andthereexistsapartition{Xi,Yi}ofViforeachi=1,2suchthatX1,Y1,X2,andY2arethesetsofimprimitivityofD.Thefollowingtheoremdescribesthelimitof{Cm(D)}∞
m=1forastronglyconnecteddigraphD.GivenagraphG,wedenotethecliqueofGwiththevertexsetZbyK[Z].7
Theorem2.6([15]).IfDisnontrivialandstronglyconnected,thenthelimitof{Cm(D)}∞
m=1isκ(D)[i=1K[Ui]whereU1,U2,...,Uκ(D)arethesetsofimprimitivityofD.ByProposition2.4,κ(D)≤4forastronglyconnectedmultipartitetournamentDandsowehavethefollowingtheorembyTheorem2.6.
Theorem2.7.LetDbeastronglyconnectedmultipartitetournament.Thenthelimitof{Cm(D)}∞
m=1isthedisjointunionofatmostfourcompletegraphswhosevertexsetsarethesetsofimprimitivityofD.Asthelimitofastronglyconnectedmultipartitetournamentistakencareof,weconsidertheonesthatarenotstronglyconnectedfromnowon.LetDbeamultipartitetournamentwithorderedstrongcomponentsQ1,Q2,...,QsandXandYbepartitesetsofQiandQj,respectively,forsome1≤i,j≤s.WesaythatXandYarepartite-relatedifthepartitesetofDincludingXandthepartitesetofDincludingYarethesame.Bydeﬁnition,ifi<jandXandYarenotpartite-related,thenX→Y.Inaddition,whenwedenotethepartitesetsofDby(V1,V2,...,Vk)forsomepositiveintegerk,unlessotherwisementioned,weassumethatVi∩V(Qs)=



U(s)i,ifκ(Qs)=2;U(s)i,ifκ(Qs)=3;U(s)i∪U(s)i+2,ifκ(Qs)=4,(2)foreachi=1,2,3(i=1,2ifκ(Qs)=2or4).TheaboveobservationtogetherwithTheorem2.7giverisetothefollowingcorollary.Corollary2.8.LetDbeamultipartitetournamentwiththenontriviallaststrongcom-ponentQsforsomeintegers≥2.Thenthesubgraphofthelimitof{Cm(D)}∞
m=1inducedbyV(Qs)isκ(Qs)[i=1KhU(s)ii.3Thecaseκ(Qs)∈{1,2,4}Proposition3.1.LetDbeamultipartitetournamentwiththeprimitivelaststrongcom-ponent.Thenthelimitof{Cm(D)}∞
m=1isacompletegraph.8
Proof.LetQsbethelaststrongcomponentofD.Taketwoverticesu∈V(D)andv∈V(Qs).WeletN=exp(Qs)+1andtakeanintegerm≥N.SinceQsisprimitive,Qsisnontrivialandsohasatleasttwopartitesets.Thereforethereisanarc(u,w)forsomevertexwinQs.Moreover,Qsbeingprimitiveimpliesthatthereare(w,v)-directedwalkoflengthm−1inQs.Since(u,w)isanarcinD,thereisa(u,v)-directedwalkoflengthmandsovisanm-steppreyofu.SinceuwasarbitrarilychosenfromV(D),visanm-steppreyofanyvertexinV(D).Thusthelimitof{Cm(D)}∞
m=1isacompletegraphby(∗).
Bondy[2]showedthat,foranintegerk≥3,astronglyconnectedk-partitetourna-mentcontainsadirectedcycleoflength`foreachinteger3≤`≤k.IfthelaststrongcomponentofamultipartitetournamentDhasatleastfourpartitesets,thenitisprimi-tiveandsothelimitof{Cm(D)}∞
m=1isacompletegraphbyProposition3.1.Thus,fromnowon,weonlyconsideramultipartitetournamentwhoselaststrongcomponenthasatmostthreepartitesets.LetDbeabipartitetournamentwiththenontriviallaststrongcomponentQsforsomeintegers≥2.Thenκ(Qs)∈{2,4}byProposition2.4andthisjustiﬁeswhyTheorem3.2onlyconsidersthecaseκ(Qs)=2andthecaseκ(Qs)=4.Theorem3.2.LetDbeabipartitetournamentwiththenontriviallaststrongcomponentQsforsomeintegers≥2.ThenthelimitGofthegraphsequence{Cm(D)}∞
m=1isasfollows:G∼=(G1,ifκ(Qs)=2;G2,ifκ(Qs)=4,whereG1andG2arethegraphsgiveninFigure2.Proof.Let(V1,V2)beabipartitionofD.SinceQsisnontrivial,κ(Qs)∈{2,4}asweobservedpriortothetheoremstatement.Moreover,V1∩V(Qs)6=∅andV2∩V(Qs)6=∅.Then,byourassumption(2),Vi∩V(Qs)=(U(s)i,ifκ(Qs)=2;U(s)i∪U(s)i+2,ifκ(Qs)=4,foreachi=1,2.ThenthereisanarcfromeachvertexinV1∩V(Qs)(resp.V2∩V(Qs))tosomevertexinV2∩V(Qs)(resp.V1∩V(Qs)).Bytheway,V1∩{V(D)−V(Qs)}→V2∩V(Qs)andV2∩{V(D)−V(Qs)}→V1∩V(Qs).Thereforethefollowingaretrue:(i)anypairofverticesinV1∩{V(D)−V(Qs)}hasacommonpreyinD;(ii)anypairofverticesinV2∩{V(D)−V(Qs)}hasacommonpreyinD;(iii)eachvertexinV1∩V(Qs)andeachvertexinV1∩{V(D)−V(Qs)}haveacommonpreyinD;9
(iv)eachvertexinV2∩V(Qs)andeachvertexinV2∩{V(D)−V(Qs)}haveacommonpreyinD.Therefore,by(∗),wehaveshownthat,foreachi=1,2,•Vi∩{V(D)−V(Qs)}formsacliqueinG;•eachvertexinVi∩V(Qs)andeachvertexinVi∩{V(D)−V(Qs)}areadjacentinG.SinceDisabipartitetournament,thereisnoedgebetweenavertexinV1andavertexinV2inG.Thus,byCorollary2.8,•ifκ(Qs)=2,thenG∼=G1whereK(1)doesnotexist,K(2)=K[V1],andK(3)=K[V2];•ifκ(Qs)=4,thenG∼=G2whereK(1)doesnotexist,K(2)=K[V1−V(Qs)],K(3)=K[V2−V(Qs)],K(4)=KhU(s)1i,K(5)=KhU(s)3i,K(6)=KhU(s)2i,andK(7)=KhU(s)4i.
Thefollowingpropositiongivesadetaileddescriptionofthelimitof{Cm(D)}∞
m=1incasewhereDcontainsastrongcomponentoneofwhosepartitesetsisnotpartite-relatedtoanypartitesetofthelaststrongcomponent.
Proposition3.3.Forsomeintegersk≥3ands≥2,letDbeak-partitetournamentwithorderedstrongcomponentsQ1,Q2,...,Qs.SupposethatQsisnontrivialandthereisanintegert∈{1,...,s−1}suchthatapartitesetofQtisnotpartite-relatedtoanypartitesetofQs.ThenanyvertexinSt i=1V(Qi)andanyvertexinDareadjacentinthelimitofthegraphsequence{Cm(D)}∞
m=1.Proof.Bythehypothesis,thereisapartitesetofQtthatisnotpartite-relatedtoanypartitesetofQs.LetZbethepartitesetofDthatcontainssuchapartitesetofQt.Taketwoverticesu∈St i=1V(Qi)andv∈V(D).SinceQsisnontrivial,Qshasatleasttwopartitesetsandsov→wforsomevertexwinQs.Ifu∈Z,thenu→V(Qs)andsowisacommonpreyofuandv.Supposeu6∈Z.Takeavertexz∈Z∩V(Qt).Thenthereisa(u,z)-directedpathPinD,whichistruebythestrongconnectednessofQtifu∈V(Qt)andbythefactthatubelongstoapartitesetdiﬀerentfromZ(whichimpliesu→Z)ifu∈St−1i=1V(Qi).Let`bethelengthofP.SinceQsisstronglyconnected,thereisa(w,x)-directedwalkP0oflength`forsomevertexx∈V(Qs).BythechoiceofZ,z→x.ThusP→xandv→P0area(u,x)-directedwalkanda(v,x)-directedwalk,respectively,ofthesamelength`+1inD.Thereforexisan(`+1)-stepcommonpreyofuandvinD.Thereforeuandvhaveanm-stepcommonpreyinDforsomepositiveintegermandsothestatementistrueby(∗).
10
Intherestofthispaper,wewillfrequentlyusethenotationSq i=pV(Qi)forpositiveintegersp,q.WewillregardSq i=pV(Qi)asanemptysetifp>q.Theorem3.4.LetDbeak-partitetournamentwithk-partition(V1,V2,...,Vk)andorderedstrongcomponentsQ1,Q2,...,Qsforsomeintegersk≥3ands≥2.SupposethatQsisnontrivialandκ(Qs)∈{2,4}.ThenthelimitGofthegraphsequence{Cm(D)}∞
m=1isasfollows:G∼=(G1,ifκ(Qs)=2;G2,ifκ(Qs)=4,whereG1andG2arethegraphsgiveninFigure2.Proof.Sinceκ(Qs)∈{2,4},QsisabipartitetournamentbyCorollary2.5andsoV(Qs)⊆V1∪V2by(2).Sincek≥3,thereexistsanintegert∈{1,...,s−1}suchthatthereisapartitesetofQtwhichisnotpartite-relatedtoanypartitesetofQs.Wemayassumethattisthelargestamongsuchintegers.ThenV(D)−t[i=1V(Qi)⊆V1∪V2.ThereforethesubdigraphD0inducedbySs i=t+1V(Qi)isabipartitetournamentwithpartitesetsA1:=V1∩Ss i=t+1V(Qi)andA2:=V2∩Ss i=t+1V(Qi).IntheproofofTheorem3.2,itisshownthatthesubgraphG1ofGinducedbyV(D0)containsH1:=K[A1]∪K[A2]ifκ(Qs)=2andH2:=S2
i=1K[Ai−V(Qs)]∨KhU(s)ii∪KhU(s)i+2iifκ(Qs)=4.SinceQt+1,...,Qsarethelastcomponentsamongtheorderedstrongcomponents,G1cannothaveedgesotherthanonesinH1orH2andsoG1=H1orG1=H2.ByProposition3.3,St i=1V(Qi)formsacliqueinGandanyvertexinSt i=1V(Qi)andanyvertexinSs i=t+1V(Qi)areadjacentinG.Thereforewemayconcludethat•ifκ(Qs)=2,thenG∼=G1whereK(1)=KSt i=1V(Qi),K(2)=K[A1],andK(3)=K[A2];•ifκ(Qs)=4,thenG∼=G2whereK(1)=KSt i=1V(Qi),K(2)=K[A1−V(Qs)],K(3)=K[A2−V(Qs)],K(4)=KhU(s)1i,K(5)=KhU(s)3i,K(6)=KhU(s)2i,andK(7)=KhU(s)4i.
4Thecaseκ(Qs)=3ByTheorems2.7,3.2,and3.4,itremainstocharacterizethelimitofthegraphsequence{Cm(D)}∞
m=1foramultipartitetournamentDwithκ(Qs)=3whereQ1,Q2,...,QsareorderedstrongcomponentsofDforsomeintegers≥2.First,weneedthefollowinglemmas.11
Lemma4.1.LetDbeamultipartitetournamentwithorderedstrongcomponentsQ1andQ2.Supposethat,forapositiveintegerLandsomeverticesx∈V(Q1)andy∈V(Q2),thereexistsan(x,y)-directedwalkoflength`foreach`≥L.ThenthereexistsapositiveintegerL0suchthat,foranyvertexuinQ1andforanyvertexvinQ2,thereexistsa(u,v)-directedwalkoflength`foreach`≥L0.Proof.SinceQ1andQ2arestronglyconnected,thereexista(u,x)-directedpathinQ1anda(y,v)-directedpathinQ2foranyu∈V(Q1)andv∈V(Q2).WeletL0=maxu∈V(Q1)d(u,x)+L+maxv∈V(Q2)d(y,v)(d(w,z)standsforthelengthofashortest(w,z)-directedpath),whichisdesired.
LetDbeamultipartitetournamentwithorderedstrongcomponentsQ1andQ2.WesaythatDisunusualifDsatisﬁesthefollowingconditions:•κ(Q1)=κ(Q2)=3;•thereexistsanintegerj∈{0,1,2}suchthatU(1)iandU(2)i+jarepartiterelatedforeachi∈{1,2,3}(identifyU(2)4andU(2)5withU(2)1andU(2)2,respectively).Lemma4.2.SupposethatamultipartitetournamentDhasorderedstrongcomponentsQ1,Q2,...,Qsandκ(Qs)=3forsomeintegers≥2.Then,forthelimitGofthegraphsequence{Cm(D)}∞
m=1,thefollowingaretrue:(i)V(D)−V(Qs)formsacliqueinG;(ii)ifκ(Qs−1)6=3andQs−1isnontrivial,theneachvertexinV(D)−V(Qs)andeachvertexinV(Qs)areadjacentinG;(iii)ifκ(Qs−1)=3,theneachvertexinV(D)−(V(Qs−1)∪V(Qs))andeachvertexinV(Qs)areadjacentinGand,further,incasewhereV(Qs−1)∪V(Qs)doesnotinduceanunusualdigraph,eachvertexinV(Qs−1)andeachvertexinV(Qs)areadjacentinG.Proof.LetDbeak-partitetournamentwithk-partition(V1,V2,...,Vk)foranintegerk≥2.Toshowpart(i),taketwoverticesu1andu2inV(D)−V(Qs).Then,sinceκ(Qs)=3,Vjdoesnotcontainanyofu1andu2forsomej∈{1,2,3}.By(2),U(s)j⊆Vj.ThereforeeachvertexinU(s)jisacommonpreyofu1andu2.Henceu1andu2areadjacentinGby(∗).Sinceu1andu2werearbitrarilychosen,(i)istrue.Toshowparts(ii)and(iii),takeu1∈V(Qs)andassumeu1∈U(s)i12
forsomei∈{1,2,3}.Toshow(ii),supposethatκ(Qs−1)6=3andQs−1isnontrivial,andtakeu2∈V(D)−V(Qs).Then,byCorollary2.5,Qs−1iseitherprimitiveorbipartite,soQs−1hasatleasttwopartitesets.Thereforethereexistsanarcfromu2toavertexinQs−1.Wetakeanout-neighborofu2thatbelongstoQs−1anddenoteitbyϕu2.WeﬁrstconsiderthecaseinwhichQs−1isprimitive.WeletN=exp(Qs−1)+2andtakeanintegerm≥N.SinceQsisstronglyconnected,wemaytakeavertexvinQssuchthatthereisa(u1,v)-directedwalkoflengthm.SinceQs−1hasatleasttwopartitesets,thereisanin-neighborx∈V(Qs−1)ofv.BythechoiceofNandthecaseassumptionthatQs−1isprimitive,thereisa(ϕu2,x)-directedwalkWoflengthm−2.Thenu2→W→visa(u2,v)-directedwalkoflengthm.Thereforevisanm-stepcommonpreyofu1andu2.Henceu1andu2areadjacentinGby(∗).NowweconsiderthecaseinwhichQs−1isbipartite.Then,foreachnonnegativeintegera,thereexistsavertexya∈V(Qs−1)suchthatthereisa(ϕu2,ya)-directedwalkoflength2a.SinceQs−1isbipartite,ϕu2andyabelongtothesamepartiteset,sayY,inQs−1.Sinceκ(Qs)=3,QshasthreepartitesetsbyCorollary2.5.Then,sinceyabelongstoYforanynonnegativeintegera,thereexistsavertexz∈V(Qs)suchthat(ya,z)isanarcforanynonnegativeintegera.Takeavertexw∈U(s)i.Forsomeα∈{1,2,3}andeachnonnegativeintegerb,thereexistsa(z,w)-directedwalkoflength3b+α.Hencethereexistsa(u2,w)-directedwalkoflength2a+3b+2+αforanynonnegativeintegersaandb.SincetheFrobeniusnumberoftheset{2,3}is1,itisguaranteedthatthereisa(u2,w)-directedwalkoflength`forany`≥4+α(recallthattheFrobeniusnumberoftheset{p,q}forrelativelyprimepandqisthelargestintegerbsuchthatbcannotbewrittenasps+qtforanynonnegativeintegerssandtandisknowntobepq−p−q).Sinceu1,w∈U(s)iandκ(Qs)=3,thereexistsa(u1,w)-directedwalkoflength3cforeachpositiveintegerc.Thereforewehaveshownthatthereexista(u1,w)-directedwalkanda(u2,w)-directedwalkofthesamelength.Henceu1andu2areadjacentinGby(∗).Sinceu1andu2werearbitrarilychosenandareadjacentinGineachcase,(ii)istrue.Toshow(iii),supposeκ(Qs−1)=3.ThenQs−1hasthreepartitesetsbyCorollary2.5,sothereexistsapartitesetXofQs−1suchthatX6⊆Xi+1∪Xi+2whereXjisthepartitesetofDincludingU(s)jforeachj∈{1,2,3}(recallu1∈U(s)iandnotethatXj=Vjforj=1,2,3).Foranonnegativeintegeraandforeachj=1,2,3,weidentifyU(s)3a+jwithU(s)jandX3a+jwithXj.Takeu2∈V(D)−(V(Qs−1)∪V(Qs)).Ifu26∈Xi+1,theneachvertexinU(s)i+1isacommonpreyofu1andu2andso,by(∗),u1andu2areadjacentinG.Supposeu2∈Xi+1.Thenthereisanarcfromu2toeachvertexinX.SinceX6⊆Xi+1∪Xi+2,thereisanarcfromanyvertexinXtoanyvertexinU(s)i+2,andsoeachvertexinU(s)i+2isa2-stepcommonpreyofu1andu2.Thusu1andu2areadjacentinGby(∗).Sinceu1andu2werearbitrarilychosen,eachvertexinV(D)−(V(Qs−1)∪V(Qs))andeachvertexinV(Qs)areadjacentinG.13
Now,toshowthe“further”part,supposethatthedigraphinducedbyV(Qs−1)∪V(Qs)isnotunusual.WewillclaimthatthedigraphinducedbyV(Qs−1)∪V(Qs)satisﬁesthehypothesisofLemma4.1ineachofthefollowingcases.Case1.SupposethatthereexistsapartitesetYofQswhichisnotpartite-relatedtoanypartitesetofQs−1.Takeapositiveinteger`,avertexx∈V(Qs−1),andavertexy∈Y.SinceQs−1isanontrivialstrongcomponent,thereexistsan(x,z)-directedwalkoflength`−1inQs−1forsomevertexz∈V(Qs−1).SinceanypartitesetofQs−1andYarenotpartite-related,thereexistsanarcfromztoy,whichresultsinan(x,y)-directedwalkoflength`.Case2.SupposethatanypartitesetofQsisincludedinthepartitesetofDincludingsomepartitesetofQs−1.ThenV(Qs−1)⊆V1∪V2∪V3.SincethedigraphinducedbyV(Qs−1)∪V(Qs)isnotunusual,withoutlossofgenerality,wemayassumethatU(s−1)1⊆V1,U(s−1)2⊆V3,andU(s−1)3⊆V2.Takex∈U(s−1)1andy∈U(s)2.Then(x,y)∈A(D).Sinceκ(Qs−1)=3,thereexistsan(x,x)-directedwalkPaoflength3ainQs−1foreachnonnegativeintegera.Takeverticesz1∈U(s−1)2andz2∈U(s)1.Then(x,z1),(z1,y),(z1,z2),(z2,y)∈A(D).ThereforePa→y,Pa→z1→y,andPa→z1→z2→yare(x,y)-directedwalksoflength3a+1,3a+2,and3a+3,respectively,foreachnonnegativeintegera.Hencethereisadirectedwalkoflength`fromxtoyforeachpositiveinteger`.Consequently,wehaveshownthatthedigraphinducedbyV(Qs−1)∪V(Qs)satisﬁesthehypothesisofLemma4.1ineachcase.Thus,thereexistsapositiveintegerLsuchthatforanyvertexxinQs−1andforanyvertexyinQs,thereexistsan(x,y)-directedwalkoflength`foreach`≥L.Takeu∈V(Qs−1)andv∈V(Qs).SinceQsisstrong,thereexistsa(v,w)-directedwalkoflengthLforsomevertexw∈V(Qs).Bythepreviousclaim,thereexistsa(u,w)-directedwalkoflengthL.HenceuandvareadjacentinGby(∗).Sinceuandvwerearbitrarilychosen,wehaveshownthe“further”partof(iii).
Lemma4.3.LetDbeak-partitetournamenthavingk-partition(V1,V2,...,Vk)andorderedstrongcomponentsQ1,Q2,...,Qswithκ(Qs)=3forsomeintegersk,s≥2andletGbethelimitofthegraphsequence{Cm(D)}∞
m=1.Further,letD1bethesubdigraphinducedbySq i=pV(Qi)∩Vjforsomep,qwith1≤p≤q<sandsomej∈{1,2,3}.(a)everyvertexinD1andeveryvertexinU(s)j∪U(s)j+1areadjacentinG(identifyU(s)iwithU(s)rwherer∈{1,2,3}andr≡i(mod3));(b)SupposethatSq i=pV(Qi)⊆VjandSs−1i=q+1V(Qi)⊆Vj+1(identifyV4withV1).Then,inG,thereisnoedgebetweenavertexinSq i=pV(Qi)andavertexinU(s)j+2.Proof.Toshow(a),takeavertexuinD1andavertexvinU(s)j∪U(s)j+1.Ifv∈U(s)j(resp.U(s)j+1),theneachvertexinU(s)j+1⊆Vj+1(resp.U(s)j+2⊆Vj+2)isacommonpreyofuandv14
(notethat,by(2),U(s)i⊆Viforeachi=1,2,3).ThereforeuandvareadjacentinGby(∗).Toshow(b),supposethatSq i=pV(Qi)⊆VjandSs−1i=q+1V(Qi)⊆Vj+1.Tothecontrary,assumethatavertexu∈Sq i=pV(Qi)⊆Vjandavertexv∈U(s)j+2haveanm-stepcommonpreyzforsomepositiveintegerm.WenotethatvbelongstothelaststrongcomponentQsandU(s)1,U(s)2,andU(s)3arethesetsofimprimitivityofQs.Thereforez∈U(s)j+2+m(3)andthereisan(u,z)-directedwalkWoflengthminD.Letw1andw2betheverticeslocatedrightnexttouandw1,respectively,onW.SinceSq i=pV(Qi)⊆Vj,thereisnoarcbetweenanytwoverticesinSq i=pV(Qi)andsow1∈Ss i=q+1V(Qi).Ifw1∈V(Qs),thenw1∈U(s)j+1∪U(s)j+2sinceu∈Vj(recallκ(Qs)=3).Nowconsiderthecasew1∈Ss−1i=q+1V(Qi)⊆Vj+1.Thenw2∈V(Qs)sincethereisnoarcbetweenanytwoverticesinSs−1i=q+1V(Qi).Sincew1∈Vj+1,w2∈U(s)j∪U(s)j+2.ThenzbelongstoU(s)j+morU(s)j+1+mineachcase,whichcontradicts(3).ThusthereisnoedgebetweenavertexinSq i=pV(Qi)andavertexinU(s)j+2inG.
Theorem4.4.SupposethatamultipartitetournamentDhasorderedstrongcomponentsQ1,Q2,...,QswithtrivialQs−1andκ(Qs)=3forsomeintegers≥2.Thenthelimitofthegraphsequence{Cm(D)}∞
m=1isisomorphictoG3giveninFigure2.Proof.LetGbethelimitofthegraphsequence{Cm(D)}∞
m=1.ByLemma4.2(i)andCorollary2.8,(a)V(D)−V(Qs)formsacliqueinGandV(Qs)inducesS3
i=1KhU(s)iiinG.SinceQs−1istrivial,thedigraphinducedbyV(Qs−1)∪V(Qs)isnotunusual.LetV(Qs−1)={vs−1}.Ifthereisanintegert∈{1,...,s−1}suchthatapartitesetofQtisnotpartite-relatedtoanypartitesetofQs,thenwetaketasthelargestamongsuchintegers.Ifsuchatdoesnotexist,thatis,Disatripartitetournament,thenwelett=0.Ift≥1,then,byProposition3.3,(b)St i=1V(Qi)formsacliqueinGandanyvertexinSt i=1V(Qi)andanyvertexinSs i=t+1V(Qi)areadjacentinG.Especially,ift=s−1,thenG∼
=G3whereK(2),K(3),andK(4)donotexistandK(1)=K"s−1[i=1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3i.Nowassumet<s−1.TheneverypartitesetofQiispartite-relatedtosomepartitesetofQsforeachi=t+1,...,s−1.Recallthat,by(2),Vi∩V(Qs)=U(s)iforeach15
i=1,2,3where(V1,V2,...,Vk)isak-partitionofD.Withoutlossofgenerality,wemayassumethatvs−1∈V2.(4)SupposethatSs−1i=t+1V(Qi)⊆Vjforsomej∈{1,2,3}.Thenj=2and,byLemma4.3,•everyvertexinSs−1i=t+1V(Qi)andeveryvertexinU(s)2∪U(s)3areadjacentinG;•thereisnoedgebetweenavertexinSs−1i=t+1V(Qi)andavertexinU(s)1inG.Therefore,by(a)and(b),G∼
=G3whereK(2)andK(3)donotexistandK(1)=K"t[i=1V(Qi)#,K(4)=K"s−1[i=t+1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3i.NowconsiderthecaseinwhichSs−1i=t+1V(Qi)6⊆Viforanyi∈{1,2,3}.Thenthereisanintegert1suchthatt<t1≤s−1andV(Qt1)6⊆V2.Wemayassumethatt1isthelargestintegeramongsuchintegers.By(4),t16=s−1andSs−1i=t1+1V(Qi)⊆V2.Then,byLemma4.3,(c1)everyvertexinSs−1i=t1+1V(Qi)andeveryvertexinU(s)2∪U(s)3areadjacentinG;(d1)thereisnoedgebetweenavertexinSs−1i=t1+1V(Qi)andavertexinU(s)1inG.ItremainstochecktheadjacencyofavertexbelongingtoSt1i=t+1V(Qi)andavertexbelongingtoV(Qs)inG.Case1.Qt1isnontrivial.SupposethatthedigraphinducedbyV(Qt1)∪V(Qs)isnotunusual.ThenanyvertexinSt1i=t+1V(Qi)andanyvertexinV(Qs)areadjacentinGbyLemma4.2(ii)and(iii)appliedtothesubdigraphinducedbySt1i=t+1V(Qi)∪V(Qs),whichisamultipartitetournamentwithorderedstrongcomponentsQt+1,...,Qt1,Qs.NowsupposethatthedigraphinducedbyV(Qt1)∪V(Qs)isunusual.Thenκ(Qt1)=3andso,ift+1≤t1−1,eachvertexinSt1−1i=t+1V(Qi)andeachvertexinV(Qs)areadjacentinGbyLemma4.2(iii)appliedtothesubdigraphinducedbySt1i=t+1V(Qi)∪V(Qs),whichisamultipartitetournamentwithorderedstrongcomponentsQt+1,...,Qt1,Qs.Ift+1=t1,thenSt1−1i=t+1V(Qi)=∅.Therefore,ineithercase,anyvertexinSt1i=t+1V(Qi)andanyvertexinV(Qs)areadjacentinGaslongasanyvertexinV(Qt1)andanyvertexinV(Qs)areadjacentinG,whichistrueasshownbelow.16
Foreachi=1,2,3,letXibethepartitesetofDincludingU(s)i(notethatXi=Vi)andtakeavertexxiinU(t1)iandavertexyiinU(s)i.Nowﬁxi∈{1,2,3}.Withoutlossofgenerality,wemayassumeU(t1)i⊆Xi.Foranonnegativeintegera,weidentifyU(t1)3a+iwithU(t1)i,U(s)3a+iwithU(s)i,andX3a+iwithXi.Then,byLemma4.3(a)appliedtothedigraphinducedbyV(Qt1)∪V(Qs)andthesubdigraphinducedbyV(Qt1)∩Vi,eachvertexinU(t1)iandeachvertexinU(s)jareadjacentinGifj=iori+1.Accordingly,itremainstoconsiderthecasej=i+2.Thefactκ(Qs)=κ(Qt1)=3guaranteesthatthereexista(yi+2,y1)-directedwalkoflength5−iinQs(identifyy4andy5withy1andy2,respectively)andan(xi,x3)-directedwalkoflength3−iinQt1.Sincevs−1∈V2,therearearcs(x3,vs−1)and(vs−1,y1)inD.Hencethereexistsan(xi,y1)-directedwalkoflength5−iinDandsoy1isa(5−i)-stepcommonpreyofxiandyi+2.Sincei,xi,andyi+2werechosenarbitrarily,anyvertexinU(t1)iandanyvertexinU(s)i+2areadjacentinGby(∗).Therefore,by(a),(b),(c1),and(d1),G∼=G3whereK(2)andK(3)donotexistandK(1)=K"t1[i=1V(Qi)#,K(4)=K"s−1[i=t1+1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3i.Case2.Qt1istrivial.LetV(Qt1)={vt1}.Takeu∈St1i=t+1V(Qi)andv∈V(Qs).Bythechoiceoft,u∈Xjforsomej∈{1,2,3}.Thenthesubdigraphinducedby{u}∪V(Qs)isamultipartitetournamentwithorderedstrongcomponent{u}andQs.Since{u}⊆Xj,itfollowsfromLemma4.3(a)thatuandvareadjacentifv∈Xj∪Xj+1.Nowweassumethatv∈Xj+2.By(4),vs−1→U(s)1∪U(s)3.(5)Subcase2-1.vt1∈X3.Thenvt1→vs−1andvt1→U(s)2.(6)Ifj6=3,thenu∈St1−1i=t+1V(Qi)sincevt1∈V3andsou→vt1.(7)Supposej=1.Thenu∈X1andv∈X3.By(6)and(7),u→vt1→U(s)2.Sinceκ(Qs)=3,wehavev→U(s)1→U(s)2andsoeachvertexinU(s)2isa2-stepcommonpreyofuandv.Supposej=2.Thenu∈X2andv∈X1.By(5),(6),and(7),u→vt1→vs−1→U(s)1.Sinceκ(Qs)=3,wehavev→U(s)2→U(s)3→U(s)1andsoeachvertexinU(s)1isa3-stepcommonpreyofuandv.Supposej=3.Thenu∈X3and17
v∈X2.By(4)and(5),u→vs−1→U(s)1.Sinceκ(Qs)=3,wehavev→U(s)3→U(s)1andsoeachvertexinU(s)1isa2-stepcommonpreyofuandv.Consequently,uandvareadjacentinGby(∗)ifv∈Xj+2.Sinceuandvwerearbitrarilychosen,wemayconcludethatanyvertexinSt1i=t+1V(Qi)andanyvertexinV(Qs)areadjacentinG.Therefore,by(a),(b),(c1),and(d1),G∼=G3whereK(2)andK(3)donotexistandK(1)=K"t1[i=1V(Qi)#,K(4)=K"s−1[i=t1+1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3i.Subcase2-2.vt1∈X1.Lett2bethesmallestnonnegativeintegerlessthant1suchthatSt1i=t2+1V(Qi)⊆X1.Bythechoiceoft,t≤t2.Then,byLemma4.3,(c2)everyvertexinSt1i=t2+1V(Qi)andeveryvertexinU(s)1∪U(s)2areadjacentinG.(d2)thereisnoedgebetweenavertexinSt1i=t2+1V(Qi)andavertexinU(s)3inG.Ift2=t,by(a),(b),(c1),(d1),(c2),and(d2),G∼
=G3whereK(1)=K"t[i=1V(Qi)#,K(3)=K"t1[i=t+1V(Qi)#,K(4)=K"s−1[i=t1+1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3i.Supposet2>t.Takeavertexu∈St2i=t+1V(Qi)andavertexv∈V(Qs).Ifu6∈X1∪X2∪X3,thenthereisanarcfromutoeveryvertexinV(Qs),andsouandvareadjacentinGby(∗).Supposeu∈Xjforsomej∈{1,2,3}.Ifv∈U(s)j(resp.U(s)j+1),theneachvertexinU(s)j+1⊆Xj+1(resp.U(s)j+2⊆Xj+2)isacommonpreyofuandv.ThereforeuandvareadjacentinGby(∗)ifu∈Xjandv∈U(s)j∪U(s)j+1forsomej∈{1,2,3}.Nowconsiderthecaseinwhichu∈Xjandv∈U(s)j+2forsomej∈{1,2,3}.Supposej=1.Thenv→U(s)1→U(s)2→U(s)3.Bythechoiceoft2,V(Qt2)−X16=∅.Ifu∈St2−1i=1V(Qi),thenuhasanout-neighborinV(Qt2)−X1.Supposeu∈V(Qt2).Sinceu∈X1andV(Qt2)−X16=∅,Qt2isanontrivialstrongcomponentandsouhasanout-neighborinQt2−X1.Thereforeu→vt2forsomevertexvt2∈V(Qt2)−X1inbothcases.Sincevt1∈X1,u→vt2→vt1→U(s)3.Ifj=2,thenu→vt1→U(s)3andv→U(s)2→U(s)3;ifj=3,thenu→vs−1→U(s)1andv→U(s)3→U(s)1.Therefore,ineachcase,uandvhaveanm-stepcommonpreyforsomepositiveintegerm.Thus,ifu∈Xjandv∈U(s)j+2forsomej∈{1,2,3},thenuandvareadjacentinGby(∗).Therefore18
everyvertexinSt2i=1V(Qi)andeveryvertexinV(Qs)areadjacentinG.Hence,by(a),(b),(c1),(d1),(c2),and(d2),G∼=G3whereK(1)=K"t2[i=1V(Qi)#,K(3)=K"t1[i=t2+1V(Qi)#,K(4)=K"s−1[i=t1+1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3i.
Theorem4.5.SupposethatamultipartitetournamentDhasorderedstrongcomponentsQ1,Q2,...,Qsandκ(Qs)=3forsomeintegers≥2.Thenthelimitofthegraphsequence{Cm(D)}∞
m=1isisomorphictoG3giveninFigure2.Further,ifV(Qs−1)∪V(Qs)doesnotinduceanunusualdigraph,thenK(2)doesnotexist.Proof.LetGbethelimitofthegraphsequence{Cm(D)}∞
m=1.ByLemma4.2(i),V(D)−V(Qs)formsacliqueinG.Inaddition,byCorollary2.8,V(Qs)inducesS3
i=1KhU(s)iiinG.Case1.ThedigraphinducedbyV(Qs−1)∪V(Qs)isunusual.Thenκ(Qs−1)=3.ByLemma4.2(iii),eachvertexinV(D)−(V(Qs−1)∪V(Qs))andeachvertexinV(Qs)areadjacentinG.Accordingly,weonlyneedtoﬁgureouttheadjacencyofeachvertexinV(Qs−1)andeachvertexinV(Qs).Tothisend,takeavertexxiinU(s−1)iandavertexyjinU(s)jforeachi,j∈{1,2,3}.LetXibethepartitesetofDincludingU(s)iforeachi∈{1,2,3}.Withoutlossofgenerality,wemayassumeU(s−1)i⊆Xiforeachi∈{1,2,3}.Foranonnegativeintegeraandforeachi=1,2,3,weidentifythesubscripts3a+iwithiforU(s),U(s−1),andX.Ifj=i(resp.j=i+1),theneachvertexinU(s)i+1(resp.U(s)i+2)isacommonpreyofxiandyj.Hence,by(∗),xiandyjareadjacentinGifj=iori+1.Supposetothecontrarythatxiandyjhaveanm-stepcommonpreyzforsomepositiveintegermwherej=i+2.WenotethatyjbelongstothelaststrongcomponentQsandU(s)1,U(s)2,andU(s)3arethesetsofimprimitivityofQs.Thereforez∈U(s)i+2+m(8)andthereisan(xi,z)-directedwalkWoflengthminD.LetW=w0→w1→···→wmwherew0=xiandwm=z.Sincexi∈V(Qs−1)andz∈V(Qs),thereisan(bytheway,theonly)arc(wt,wt+1)onWsuchthatwt∈V(Qs−1)andwt+1∈V(Qs)forsomet∈{0,1,...,m−1}.Thenwtandwt+1belongtodistinctpartitesets.Sincew0∈U(s−1)i,wt∈U(s−1)i+t.Thereforewt+1belongstoU(s)i+t+1orU(s)i+t+2.ThenwmbelongstoU(s)i+morU(s)i+m+1whichcontradicts(8).Sincexiandyjwerearbitrarilychosen,wemayconcludethatj≡i+2(mod3)⇔foranypositiveintegerm,novertexinU(s−1)iandnovertexinU(s)jhaveanm-stepcommonprey⇔thereisnoedgebetweenU(s−1)iandU(s)jinG.19
HenceG∼=G3whereK(1)=K"s−2[i=1V(Qi)#,K(2)=KhU(s−1)1i,K(3)=KhU(s−1)3i,K(4)=KhU(s−1)2i,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3iifV(Qs−1)∪V(Qs)inducesanunusualdigraph.Case2.ThedigraphinducedbyV(Qs−1)∪V(Qs)isnotunusual.IfQs−1istrivial,thenG∼=G3byTheorem4.4.IfQs−1isnontrivial,theneachvertexinV(D)−V(Qs)andeachvertexinV(Qs)areadjacentinGby(ii)and(iii)ofLemma4.2.HenceG=G3whereK(2),K(3)andK(4)donotexistandK(1)=K"s−1[i=1V(Qi)#,K(5)=KhU(s)1i,K(6)=KhU(s)2i,andK(7)=KhU(s)3iifV(Qs−1)∪V(Qs)inducesadigraphthatisnotanunusualdigraphandQs−1isnontrivial.
5AproofofTheorem1.2
ProofofTheorem1.2.SupposethatQsisnontrivial.SinceQsisastronglyconnectedmultipartitetournament,κ(Qs)≤4byProposition2.4.Ifκ(Qs)=1,thenG=K[V(D)]byProposition3.1.Thusitremainstoconsiderthecasesκ(Qs)=2,3,4.Supposes=1.Then,byTheorem2.7,thefollowingaretrue:ifκ(Qs)=2,thenG∼=G1whereK(1)doesnotexist;ifκ(Qs)=3,thenG∼=G3whereK(i)doesnotexistforeachpositiveintegeri≤4;ifκ(Qs)=4,thenG∼=G2whereK(i)doesnotexistforeachpositiveintegeri≤3.Nowsupposes≥2.Ifκ(Qs)=2(resp.κ(Qs)=4),thenG∼=G1(resp.G∼=G2)byTheorems3.2and3.4.Ifκ(Qs)=3,thenG∼
=G3byTheorem4.5.
6Disclosurestatement
Nopotentialconﬂictofinterestwasreportedbytheauthors.
7Funding
ThisresearchwassupportedbytheNationalResearchFoundationofKorea(NRF)(NRF-2022R1A2C1009648andNRF-2017R1E1A1A03070489)fundedbytheKoreagovernment(MSIP).20
References[1]E.Belmont.Acompletecharacterizationofpathsthatarem-stepcompetitiongraphs.DiscreteApplMath,159(14):1381–1390,2011.[2]J.A.Bondy.Diconnectedorientationsandaconjectureoflasvergnas.JournaloftheLondonMathematicalSociety,2(2):277–282,1976.[3]R.A.Brualdi,H.J.Ryser,etal.Combinatorialmatrixtheory,volume39.Springer,1991.[4]H.H.ChoandH.K.Kim.Competitionindicesofdigraphs.InProceedingsofworkshopincombinatorics,volume99,pages96–107,2004.[5]H.H.ChoandH.K.Kim.Thecompetitionindexofanearlyreduciblebooleanmatrix.BulletinoftheKoreanMathematicalSociety,50(6):2001–2011,2013.[6]H.H.Cho,S.-R.Kim,andY.Nam.Them-stepcompetitiongraphofadigraph.DiscreteApplMath,105(1):115–127,2000.[7]J.E.Cohen.Intervalgraphsandfoodwebs:aﬁndingandaproblem.RANDCorporationDocument,17696,1968.[8]S.Eoh,S.-R.Kim,andH.Yoon.Onm-stepcompetitiongraphsofbipartitetourna-ments.DiscreteApplMath,283:199–206,2020.[9]G.T.Helleloid.Connectedtriangle-freem-stepcompetitiongraphs.DiscreteApplMath,145(3):376–383,2005.[10]W.Ho.Them-step,same-step,andany-stepcompetitiongraphs.DiscreteApplMath,152(1):159–175,2005.[11]J.-H.Jung,S.-R.Kim,andH.Yoon.Competitionperiodsofmultipartitetournaments.LinearandMultilinearAlgebra,https://doi.org/10.1080/03081087.2022.2038057.[12]H.K.Kim.Generalizedcompetitionindexofaprimitivedigraph.Linearalgebraanditsapplications,433(1):72–79,2010.[13]H.K.Kim.Characterizationofirreduciblebooleanmatriceswiththelargestgener-alizedcompetitionindex.LinearAlgebraAppl,466:218–232,2015.[14]B.Park,J.Y.Lee,andS.-R.Kim.Them-stepcompetitiongraphsofdoublypartialorders.ApplMathLett,24(6):811–816,2011.21
[15]W.Park,B.Park,andS.-R.Kim.Amatrixsequence{Γ(Am)}∞
m=1mightconvergeevenifthematrixAisnotprimitive.LinearAlgebraAppl,438(5):2306–2319,2013.[16]Y.ZhaoandG.J.Chang.Noteonthem-stepcompetitionnumbersofpathsandcycles.DiscreteApplMath,157(8):1953–1958,2009.22

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

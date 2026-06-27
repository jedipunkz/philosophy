---
source: "https://arxiv.org/abs/math/0605149v1"
title: "Notes on duality theories of abelian groups"
author: "Gábor Lukács"
year: "2006"
publication: "arXiv preprint / math.GN"
download: "https://arxiv.org/pdf/math/0605149v1"
pdf: "https://arxiv.org/pdf/math/0605149v1"
captured_at: "2026-06-18T11:15:35Z"
updated_at: "2026-06-18T11:15:35Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "György Lukács"
tags:
  - "現代思想"
  - "マルクス主義"
  - "物象化論"
  - "西洋マルクス主義"
status: raw
---

# Notes on duality theories of abelian groups

- 著者: Gábor Lukács
- 年: 2006
- 掲載情報: arXiv preprint / math.GN
- 情報源: [arxiv](https://arxiv.org/abs/math/0605149v1)
- ダウンロード: https://arxiv.org/pdf/math/0605149v1
- PDF: https://arxiv.org/pdf/math/0605149v1

## Obsidian Links

- 研究動向: [[研究動向/ルカーチ・ジェルジュ-現代研究動向|ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代思想]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[物象化論]]
- 関連分野: [[西洋マルクス主義]]
- 関連タグ: #現代思想 #マルクス主義 #物象化論 #西洋マルクス主義

## Abstract

In this notebook, I present duality theory (or theories) of abelian groups with some categorical and categorical topological flavour. I consider writing this notebook as a longer-term project, and its current content and presentation is "under development." In other words, all questions, comments, suggestions, and criticism is more than usually welcome. In Chapter 1, which is the only more-or-less ready part, the evaluation homomorphism A --> A^^ (of a group A into its bidual A^^) is studied, with particular attention to quasi-convexity and special groups (compact, discrete, LCA) and subgroups (open, compact). The treatment of LCA groups is incomplete, but the ultimate goal is to make it as self-contained as possible.

## PDF Text

arXiv:math/0605149v1 [math.GN] 5 May 2006
Notesondualitytheoriesofabeliangroups1G´aborLuk´acs2July25,201712000MathematicsSubjectClassiﬁcation:22-02(22A0543A40)2IgratefullyacknowledgethegenerousﬁnancialsupportreceivedfromtheKillamTrustsandDalhousieUniversitythatenabledmetodothisresearch.
Contents
Introduction1IThePontryagindual3I.ATheevaluationhomomorphism............................3I.1Pontryagindual................................3
I.2Polar.....................................3
I.6Evaluationhomomorphism..........................5I.10Localquasi-convexity............................7I.21Precompactness................................10I.24Quasi-convexcompactness..........................10I.BSpecialgroupsandsubgroups............................11I.28Metrizablegroups..............................11I.35Compactanddiscretegroups.........................14I.38Bohr-compactiﬁcation............................15I.41Openandcompactsubgroups........................16I.47Locallycompactgroups...........................19IICategoricalandtopologicalgeneralizations22II.AGroupobjectsinacategory..............................22II.1Deﬁninggroupsusingdiagrams.......................22II.2Grp(X)andAb(X)..............................22II.3Cartesianclosedness.............................22II.4Internalhom.................................22II.Bk-groups........................................22II.5Grp(kHaus)..................................22II.6k-groupsofNoble..............................22II.CConvergencegroups..................................23Appendix24ASeparationpropertiesoftopologicalgroups.....................24BExponentiabilityandthecompact-opentopology..................25CHausdorffk-spaces..................................27(iii)
Introduction
InDecember2005,IhadthegoodfortuneofspendingaweekastheguestofReinhardWinkleratTechnischeUnivert¨atWien,andhadthehonourofbeingaskedtodeliveramini-courseonthetopicofdualitytheoryofabeliangroups.Duringthisvisit,IcametorealizethattheonlyelementarytextonthistopicthatIcouldreferagraduatestudentto,namelytheonebyPontryagin[30],waspublished20yearsago,anditswell-organizedcontentishalfacenturyold.Thisisnottosaythattherearenoexcellentbooksthatcontainalotofresultsonabeliantopologicalgroups(suchas[12],[33],[26],or[2]),butmyfeelingisthattheydepartfromtheaimofpresentingbasicnotionsofdualitytheoryinaself-containedandelementarymanner.Therefore,encouragedbymyhostandhisresearchgroupinVienna,IdecidedtowriteupthenotesthatIpreparedfortheminicourse(Chapters1and2),andtokeepdevelopingitasalongertermproject.Inordertoprovideasmootherpresentation,someelementaryresultsofwerecollectedintheAppendix.Iuseitalicsfontforresultsthatappeartobenewandnotpartofthe”commonknowledge,”whileslantedfontisusedforallotherstatements.LetmeknowifIamwrongaboutanyofthem.Pastandfuturework
Thisnotebookisunderdevelopment,soallsuggestions,comments,andquestionsarewarmlywelcome.
Chapter1.Thisisdeﬁnitelytheonlymore-or-lessreadypiece,butIamnothappywithmytreat-mentoflocallycompactabeliangroups.(SimilarlytoRoeder'sapproach(cf.[32]),theonlyresultborrowedfromfunctionalanalysisisthePeter-Weyltheorem.IamstillmissinganiceproofofthestructuretheoremofcompactlygeneratedLCAgroups,andacompleteproofofthePontryaginduality.)
Chapter2.Ihaveonlyasketch.Afteracategoricalintroduction,whichexplainswhycartesianclosedcategoriesaresointeresting,Ihopetocoverk-groups(ofNoble)andconvergencegroupsinthischapter.
Appendix.Itevolvesasthechaptersdevelop.Itisacollectionofresultsmanyreadersmightbefamiliarwith,soIsawnopointinincludingthemintothechapters.Futuretopics.Precompact(abelian)groups[theproblemisthatitrequiresthereadertobefamiliarwithnotionssuchasC-embeddedsubsets,etc];examplesofpathologicalorotherwiseinterestinggroups[although[12]isprobablythebestsourceforthis];localicgroups[itwouldbeaniceexampleofgroupobjectsinacategory].
2G.Luk´acs/NotesondualitytheoriesofabeliangroupsTopicsleftout.Nucleargroups,topologicalvectorspaces[thesearebothimportant,butIamnotsureifIcanpresenttheminanelementaryway].
Acknowledgements
IamdeeplyindebtedtoReinhardtWinklerandhisresearchgroupfortheirhospitalityduringmyvisitatTechnischeUnivert¨atWien,whoinspiredandencouragedmetoorganizemythoughtsandnotesonthetopic.IwishtothankSalvadorHern´andez,KeithJohnson,andJeffEggerforthevaluablediscussionsandhelpfulsuggestionsthatwereofgreatassistanceinwritingthispaper.
ChapterI
ThePontryagindual
I.A.Theevaluationhomomorphism
I.1.Pontryagindual.ForHausdorfftopologicalgroupsGandH,wedenotebyH(G,H)thespaceofcontinuoushomomorphismsϕ:G→Hequippedwiththecompact-opentopology.(Forabriefreviewofthecompact-opentopology,seeAppendixB.)Sincethepropertyofbeingahomomorphismisequationallydeﬁned,H(G,H)isaclosedsubspaceofC(G,H).Indeed,H(G,H)=\g1,g2∈G{f∈C(G,H)|f(g1g2)=f(g1)f(g2)},(I.1)andeachofthesetsintheintersectionisclosed.Topologicalgroupshavenaturaluniformstruc-tures,andthecompact-opentopologycanalsoberealizedasthetopologyofuniformconvergenceoncompacta:ϕα−→ϕinH(G,H)ifϕα|
|Kconvergesuniformlytoϕ||KforeverycompactsubsetK⊆G.ForA∈Ab(Top),acharacterofAisahomomorphismχ:A→T,whereT:=R/Z.Ifχiscontinuous,thenitmustfactorthroughthemaximalHausdorffquotientA/NAofA,becauseTisHausdorff(cf.PropositionA.2).ThePontryagindualˆAofAisthegroupH(A/NA,T)ofallcontinuouscharactersofAequippedwiththecompact-opentopology.SincecompactsubsetsofA/NAarepreciselytheimagesofcompactsubsetsinA,onemayviewˆAasthegroupH(A,T)bothalgebraicallyandtopologically(cf.CorollaryA.4).WeputAdforthegroupAequippedwiththediscretetopology.ThespaceC(Ad,T)coincideswithTAd,andthuscompact.Therefore,itsclosedsubspacecAdisalsocompact;itcarriesthetopologyofpointwiseconvergenceonA,inotherwords,χα→χifandonlyifχα(a)→χ(a)foreverya∈A.Sincethecompact-opentopologyisﬁnerthanthepointwiseconvergenceone,ˆA→cAdiscontinuous.I.2.Polar.ThesetsΛn=[−1
4n,1
4n]formabaseat0forT,andhavethepropertythatifkx∈Λ1fork=1,...,n,thenx∈Λn.(Here,[−1
4n,1
4n]areidentiﬁedwiththeirimagesinT.)ForS⊆A
4G.Luk´acs/NotesondualitytheoriesofabeliangroupsandΦ⊆ˆA,theirpolarsetsaredeﬁnedasSB={χ∈ˆA|χ(S)⊆Λ1},(I.2)ΦC={a∈A|∀χ∈Φ,χ(a)∈Λ1}=\χ∈Φχ−1(Λ1).(I.3)SinceSBisclosedinthetopologyofpointwiseconvergence,inparticular,itisclosedinthecompact-openone.Ontheotherhand,ΦCisclosedinAbecauseeachχ∈Φiscontinuous.Itisimmediatelyseenthat(C,B)isaGalois-connectionbetweensubsetsofAandˆA.Thus,S⊆SBCandΦ⊆ΦCB.I.3.Lemma.Letϕ:A→BbeamorphisminAb(Top).(a)ForS⊆A,ˆϕ−1(SB)=ϕ(S)B.(b)ForΣ⊆ˆB,ϕ−1(ΣC)=ˆϕ(Σ)C.(c)ForS0⊆B,ˆϕ(S0B)⊆ϕ−1(S0)B.(d)ForΣ0⊆ˆA,ϕ(Σ0C)⊆ˆϕ−1(Σ0)C.(e)ForS⊆A,ϕ(SBC)⊆ϕ(S)BC.(f)ForS0⊆B,ϕ−1(S0)BC⊆ϕ−1(S0BC).Proof.Sinceϕiscontinuous,itinducesacontinuoushomomorphismˆϕ:ˆB→ˆAdeﬁnedbyˆϕ(χ)=χ◦ϕ.χ∈ϕ(S)B⇐⇒χ(ϕ(S))⊆Λ1⇐⇒χ◦ϕ∈SB⇐⇒χ∈ˆϕ−1(SB),(I.4)andso(a)follows.g∈ϕ−1(ΣC)⇐⇒ϕ(g)∈ΣC⇐⇒∀χ∈Σ,χ(ϕ(g))∈Λ1(I.5)⇐⇒∀ψ∈ˆϕ(Σ),ψ(g)∈Λ1⇐⇒g∈ˆϕ(Σ)C,(I.6)whichshows(b).(c)Sinceϕ(ϕ−1(S0))⊆S0,onehasS0B⊆ϕ(ϕ−1(S0))B.By(a)appliedtoS=ϕ−1(S0),ϕ(ϕ−1(S0))B=ˆϕ−1(ϕ−1(S0)B).Thus,S0B⊆ˆϕ−1(ϕ−1(S0)B).(d)Sinceˆϕ(ˆϕ−1(Σ0))⊆Σ0,onehasΣ0C⊆ˆϕ(ˆϕ−1(Σ0))C.By(b)appliedtoΣ=ˆϕ−1(Σ0),ˆϕ(ˆϕ−1(Σ0))C=ϕ−1(ˆϕ−1(Σ0)C).Thus,Σ0C⊆ϕ−1(ˆϕ−1(Σ0)C).(e)followsfrom(a)and(d),and(f)followsfrom(b)and(c).
I.4.Lemma.LetA∈Ab(Top),H≤Aitssubgroup,andletπH:A→A/Hbethecanonicalprojection.Then:(a)HBisasubgroup;(b)ˆπH:[A/H→ˆAisinjective,anditsimageisHB;(c)([6,5])ifS0⊆A/Hand0∈S0,thenπ−1H(S0)B=ˆπH(S0B).
G.Luk´acs/Notesondualitytheoriesofabeliangroups5Proof.(a)SinceΛ1containsonlythezerosubgroup,χ(H)⊆Λ1holdsifandonlyifH⊆kerχ.Thus,HBistheannihilatorsubgroupofHinˆA.(b)SinceπHisonto,ˆπHisinjective.Clearly,ImˆπH⊆HB,soinordertoshowtheconverse,letχ∈HB.ThenH⊆kerχ,andthusχinducesacontinuouscharacter¯χ:A/H→TsuchthatˆπH(¯χ)=¯χ◦πH=χ.(c)OnehasH⊆π−1H(S0)because0∈S0,andthusπ−1H(S0)B⊆HB=ImˆπH,by(b).LemmaI.3(c)appliedtoS=π−1H(S0)yieldsˆπ−1H(π−1H(S0)B)=πH(π−1H(S0))B=S0B.Thiscompletestheproof,becauseˆπHisinjective.
ForX∈Top,wedenotebyK(X)thecollectionofcompactsubsetsofX.ForG∈Grp(Top),oneputsN(G)forthecollectionofneighborhoodsofeinG.I.5.Proposition.LetA∈Ab(Top).(a)χ∈hom(A,T)iscontinuousifandonlyifthereisU∈N(A)suchthatχ(U)⊆Λ1.(b)Thecollection{KB|K⊆A,K∈K(A)}isabaseat0toˆA.(c)Σ⊆ˆAisequicontinuousifandonlyifthereisU∈N(A)suchthatΣ⊆UB.Proof.Observethatinordertoestablishcontinuity-likepropertiesofhomomorphismsoftopo-logicalgroups,itsufﬁcestocheckthematasinglepoint,and0isaconvenientchoiceforthatpurpose.(a)Necessityisobvious.So,inordertoshowsufﬁciency,letn∈NandletU∈N(A)besuchthatχ(U)⊆Λ1.BecauseofthecontinuityofadditioninA,thereisV∈N(A)suchthatV+···+V|
{z
}ntimes⊆U,andinparticular,kχ(V)⊆Λ1foreveryk=1,...,n.Therefore,χ(V)⊆Λn,asdesired.(b)LetC∈K(A),andsetK=C∪2C∪···∪nC.ThenK∈K(A),andχ(C)⊆Λnforeveryχ∈KB,becausekχ(a)∈Λ1foreverya∈Candk=1,...,n.Thus,weobtainedKB⊆{χ∈ˆA|χ(C)⊆Λn},asdesired.(c)Becauseofthehomogeneousstructureoftopologicalgroups,equicontinuityofafamilyΣofcharacterscanbecheckedatasinglepoint.Equicontinuityat0meansthatforeveryneigh-borhoodΛnof0thereisV∈N(A)suchthatχ(V)⊆Λnforeveryχ∈Σ.Bytheargumentpresentedin(a),thisconditionissatisﬁedforallnifandonlyifitissatisﬁedforn=1.Sincethisconditionforn=1ispreciselyΣ⊆UB,theproofiscomplete.
I.6.Evaluationhomomorphism.Eacha∈Agivesrisetoacontinuouscharacterˆaofhom(A,T)givenbyevaluation:ˆa(ψ)=ψ(a).Inparticular,ˆaisacontinuouscharacterofˆA(whosetopologyisﬁnerthanthatofhom(A,T)).Thus,ˆa∈ˆ
ˆA,andtheevaluationmapαA:A−→ˆ
ˆA(I.7)a7−→ˆa(I.8)isahomomorphismofgroups.Inthesequel,wepresentnecessaryandsufﬁcientconditionsforαAtobecontinuousandanembedding(cf.PropositionsI.7andI.12),anditwillbecomeclearthatingeneral,αAneednotbecontinuous.
6G.Luk´acs/NotesondualitytheoriesofabeliangroupsAcollectionCofcompactsetsofatopologicalspaceXisacobaseifforeverycompactsubsetK⊆XthereisC∈CsuchthatK⊆C.I.7.Proposition.([28,2.3])LetA∈Ab(Top).(a)ForeveryU∈N(A),UBiscompactinˆA.(b)Thefollowingstatementsareequivalent:(i)αAiscontinuous;(ii)everycompactsubsetofˆAisequicontinuous;(iii){UB|U∈N(A)}isacobaseforˆA;(iv){UBB|U∈N(A)}isabaseforˆ
ˆAat0.(c)kerαA=Tχ∈ˆAkerχ.Proof.(a)Asnotedearlier,UBisclosedinˆA(itisclosedeveninhom(A,T)).ByProposi-tionI.5(c),itisalsoequicontinuous,andthereforeitiscompactbyastandardArzela-Ascolitypeargument.(b)Theequivalenceof(ii)and(iii)isanimmediateconsequenceofPropositionI.5(c).(i)⇒(ii):SupposethatαAiscontinuous,andletΣ⊆ˆAbecompact.ThenΣBisaneighbor-hoodof0inˆ
ˆA.Thus,bycontinuityofαA,thereisU∈N(A)suchthatαA(U)⊆ΣB,whichisequivalenttoΣ⊆UB.Therefore,byPropositionI.5(c),Σisequicontinuous.(iii)⇒(iv):IfV∈N(ˆ
ˆA),thenbyPropositionI.5,thereisacompactsubsetΦ⊆ˆAsuchthatΦB⊆V.Since{UB|U∈N(A)}isacobaseforˆA,thereisU∈N(A)suchthatΦ⊆UB.Therefore,UBB⊆ΦB⊆V,asdesired.(iv)⇒(i):IfW∈N(ˆ
ˆA),thenthereisU∈N(A)suchthatUBB⊆W,andthereforeαA(U)⊆UBB⊆W.Hence,αAiscontinuous.(c)Fora∈A,onehasa∈kerαA⇐⇒αA(a)=0⇐⇒∀χ∈ˆA,(αA(a))(χ)=0⇐⇒∀χ∈ˆA,χ(a)=0,(I.9)whichcompletestheproof.
AcombinationofPropositionsI.5(b)andI.7(b)yields:I.8.Corollary.LetA∈Ab(Top).Thefollowingstatementsareequivalent:(i)αˆAiscontinuous;(ii){KBB|K∈K(A)}isacobaseforˆ
ˆA.
Amapf:X→YbetweenHausdorffspacesissaidtobek-continuousiftherestrictionf|
|KiscontinuousforeverycompactsubsetKofX.AlthoughαAneednotbecontinuous(cf.Propo-sitionI.7),itsrestrictiontoanycompactsubsetofAiscontinuous.I.9.Theorem.LetA∈Ab(Haus).TheevaluationhomomorphismαAisk-continuous.Proof.ByCorollaryB.3,e:A→C(C(A,T),T)isk-continuous.Sincetheimagee(A)iscontainedinH(C(A,T),T),andthemapH(C(A,T),T)→H(H(A,T),T)=ˆ
ˆA(givenbyrestrictiontoH(A,T))iscontinuous,αAisthecompositionofak-continuousmapwithacontinuousone.
G.Luk´acs/Notesondualitytheoriesofabeliangroups7I.10.Localquasi-convexity.LetA∈Ab(Top).FollowingBanaszczyk,S⊆Aisquasi-convexifS=SBC.ThegroupAislocallyquasi-convex(orbrieﬂy,LQC)ifitadmitsabaseofquasi-convexneighborhoodsat0,thatis,if{UBC|U∈N(A)}isabaseat0(cf.[3]).I.11.Proposition.([6,1])ThegroupˆAislocallyquasi-convexforeveryA∈Ab(Top).Proof.ByPropositionI.5,thecollection{KB|K∈K(A)}isabaseat0atˆA.Eachmemberofthecollectionisquasi-convex,becauseKBBC=KB.
Avariantofthenextpropositionappearedin[2,6.10]and[9,4.3],andseemstobeawell-knownresult.
I.12.Proposition.LetA∈Ab(Top).(a)IfαAisanembedding,thenAislocallyquasi-convex.(b)IfAislocallyquasi-convex,thenαAisopenontoitsimage.(c)IfAislocallyquasi-convexandHausdorff,thenαAisinjective.Proof.First,notethatforS⊆A,onehasαA(SBC)=SBB∩αA(A).IfU∈N(A),thenUBiscompactinˆA(cf.PropositionI.7(a)),andsoUBBisopeninˆ
ˆAbyPropositionI.5(b).Therefore,αA(UBC)isopeninαA(A).(a)SinceαAisanembedding,αA(V)isopeninαA(A)foreveryV∈N(A).Thus,byPropositionI.7(b)(iv),thereisU∈N(A)suchthatαA(UBC)=UBB∩αA(A)⊆αA(V),whichimpliesUBC⊆V,becauseαAisinjective.(b)LetV∈N(A),andusingLQCpickU∈N(A)suchthatUBC⊆V.ThenonehasαA(UBC)⊆αA(V),andαA(UBC)isopeninαA(A).(c)Leta∈Abeanon-zeroelement.SinceAisHausdorffandLQC,thereexistsU∈N(A)suchthata6∈UBC.Thus,thereisχ∈UBsuchthat(αA(a))(χ)=χ(a)6∈Λ1.Therefore,αA(a)6=0.Hence,αAisinjective.
I.13.Remarks.(1)InPropositionI.12(b),themapαAneednotbeanembedding,becauseitisnotnecessarilycontinuousorinjective.(2)FollowingNeumannandWigner,wheneverαAisinjective,Aissaidtobemaximallyalmost-periodic(orbrieﬂy,MAP;cf.[27]).(3)ForeveryA∈Ab(Top),αˆAisinjective.Indeed,ifαˆA(χ)=0,thenforeveryξ∈ˆ
ˆA,ξ(χ)=(αˆA(χ))(ξ)=0.Thus,fora∈Aandξ=αA(a),oneobtainsχ(a)=(αA(a))(χ)=0.Therefore,χ=0,asdesired.I.14.Corollary.LetA∈Ab(Haus)besuchthatαAiscontinuous.ThenαAisanembeddingifandonlyifAislocallyquasi-convex.
WedenotebyLQCthefullsubcategoryofAb(Top)formedbythelocallyquasi-convexgroups,andpresentafunctorialmethodof“turning”everygroupintoanLQCgroup.I.15.Theorem.LQCisanepireﬂectivesubcategoryofAb(Top),andthereﬂectionAMisgivenbyequippingthegroupofAwithanewgrouptopologywhosebaseat0is{UBC|U∈N(A)}.
8G.Luk´acs/NotesondualitytheoriesofabeliangroupsInlesscategoricallanguage,TheoremI.15statesthatforeveryA∈Ab(Top):(1)AMisatopologicalgroup;(2)thereisacontinuoussurjectivehomomorphismνA:A→AMthatisnaturalinA;(3)ϕM:AM→BMiscontinuouswheneverϕ:A→Bisacontinuoushomomorphism;(4)Everycontinuoushomomorphismγ:A→CintoanLQCgroupCfactorsuniquelythroughνA,thatis,thereisauniqueγ0:AM→Csuchthatγ=γ0◦νA.Proof.(1)EachsetoftheformUBCissymmetric(i.e.,−UBC=UBC),anditiseasilyseenthat(U∩V)BC⊆UBC∩VBC.LetU∈N(A),andletV∈N(A)besuchthatV+V⊆U.ThenVBC+VBC⊆(V+V)BC⊆UBC.Therefore,theproposedcandidateforN(AM)deﬁnesagrouptopologyontheunderlyinggroupofA,asdesired.(2)ThetopologyofAMiscoarserthanthatofA,sotheidentityhomomorphismA→AMiscontinuous(anditisobviouslynaturalinA).(3)LetUBC∈N(BM).ByLemmaI.3(f),ϕ−1(UBC)⊇ϕ−1(U)BC.Sinceϕiscontinuous,ϕ−1(U)∈N(A),andthereforeϕ−1(U)BC∈N(AM).Hence,ϕMiscontinuous.(4)SinceAandAMhavethesameunderlyingset,uniquenessofγ0isclear.IfC∈LQC,thenC=CM,andthereforeγ0=γM,by(3).
I.16.Corollary.Thelimitofafamilyoflocallyquasi-convexgroupsformedinAb(Top)islocallyquasi-convex,andcoincideswiththelimitformedinLQC.Inparticular,LQCisclosedundertheformationofarbitraryproductsandsubgroups.
Proof.Theﬁrststatementisawell-knowncategorytheoreticalpropertyofreﬂectivesubcate-gories(cf.[24,IV.3,V.5]),anditimpliesthesecondone,asproductsarelimits.Forthethirdstatement,letAbeanLQCgroupandH≤Abeitssubgroup.Thereﬂection(A/H)MhasthesameunderlyinggroupasA/H,andthereforeH=Eq(AπMH//
0//
(A/H)M),(I.10)whereπH:A→A/Histhecanonicalprojection.Sinceequalizersarelimits,thiscompletestheproof.
I.17.Corollary.ForeveryA∈Ab(Top),theunderlyinggroupsofˆAandcAMcoincide.Proof.ByPropositionI.11thegroupˆZ=TisLQC.Thus,byTheoremI.15,everycontinuouscharacterχ:A→TgivesrisetoacontinuousχM:AM→T.Therefore,ˆA⊆cAMassets.Thereverseinclusionisobvious,becauseAMcarriesacoarsertopologythanA.
I.18.Proposition.LetA∈Ab(Top)suchthatαAiscontinuous.Then:(a)ˆAandcAMhavethesamecompactsubsets.(b)AMcoincideswiththegroupAequippedwiththeinitialtopologyinducedbyαA;(c)NAM=kerαAM=kerαA;(d)αAMiscontinuousandopenontoitsimage;(e)ˆ
ˆνA:ˆ
ˆA→c cAMisanembedding;
G.Luk´acs/Notesondualitytheoriesofabeliangroups9Proof.SinceˆνA:cAM→ˆAiscontinuous,everycompactsubsetΦofcAMisalsocompactinˆA,andthus,byPropositionI.7(b),thereexistsU∈N(A)suchthatΦ⊆UB.(a)SupposethatΣ⊆ˆAiscompact.Then,byPropositionI.7(b),Σ⊆UBforsomeU∈N(A).ByPropositionI.7(a),UB=(UBC)BiscompactincAM.SinceΣ=ˆν−1A(Σ)isaclosedsubsetofUB,ΣiscompactincAM.(b)ByPropositionI.7(b),{UBB|U∈N(A)}isabaseforˆ
ˆAat0,becauseαAiscontinuous.Thestatementsfollowsfromα−1A(UBB)=UBC.(c)Theﬁrstequalityfollowsfrom(b)andPropositionA.2,becauseˆ
ˆAisHausdorff.Forthesecondequality,observethatbyPropositionI.7(c),kerαA=Tχ∈ˆAkerχ,andby(a),AandAMhavethesamecontinuouscharacters.(d)By(a),polarsofquasi-convexneighborhoodsformacobasetocAM,andthusαAMiscontin-uous,byPropositionI.7(b).Therefore,thestatementfollowsbyPropositionI.12(b).(e)ByPropositionI.7(b),thecollections{UBB|U∈N(A)}and{VBB|V∈N(AM)}arebasesat0toˆ
ˆAandc cAM,respectively.SinceeveryV∈N(AM)hastheformofUBC,thiscompletestheproof.
I.19.Theorem.LetA∈Ab(Top),andD≤Abeadensesubgroup.Then:(a)theunderlyinggroupsofˆDandˆAcoincide;(b)ifαDiscontinuous,ˆDandˆAhavethesamecompactsubsets,andαAiscontinuous;(c)ifDislocallyquasi-convex,thensoisA.Proof.(a)TheinclusionιD:D→AinducesacontinuoushomomorphismˆιD:ˆA→ˆD,whichisinjective,becauseDisdenseinA.InordertoshowthatˆιDissurjective,letχ∈ˆD.Foreacha∈A,thereisanet{xα}⊆Dsuchthatxα−→a.Thus,xα−xβ−→0,andsoχ(xα)−χ(xβ)−→0.Inotherwords,{χ(xα)}isaCauchynetinthecompletegroupT,andthereforelimχ(xα)exists.Set¯χ(a)=limχ(xα);itiseasilyseenthat¯χiswell-deﬁnedandcontinuous.Hence,ˆιD(¯χ)=χ,asdesired.(b)Itsufﬁcestoshowthatˆι−1D(Φ)iscompactinˆAforeverycompactsubsetΦ⊆ˆD,becauseˆιDiscontinuous.ByPropositionI.7(b),sinceαDiscontinuous,ΦisequicontinuousonthedensesubgroupD,whichmeansthatthereisU∈N(A)suchΦ⊆(U∩D)B(cf.PropositionI.5(b)).ItfollowsfromthedensityofDthatIntU⊆
D∩U,andthereforeΦ⊆(U∩D)B=(
U∩D)B⊆(IntU)B.(I.11)Hence,ˆι−1D(Φ)isaclosedsubsetofthecompactsubset(IntU)BofˆA(cf.PropositionI.5(a)).SinceIntU∈N(A),thisalsoshowsthateverycompactsubsetofˆAiscontainedinUBforsomeU∈N(A),whichmeans,byPropositionI.7(b),thatαAiscontinuous.(c)LetV∈N(A),andpickV1∈N(A)suchthatV1−V1⊆V,so¯V1⊆IntV(cf.Proposi-tionA.1(a)).ThereexistsW1∈N(D)suchthatWBC1∩D⊆V1∩D,becauseV1∩D∈N(D)andDisLQC.(ForΣ⊆ˆD,ΣCinDisˆι−1D(Σ)C∩D,becauseˆD=ˆAassets.)ThereisW2∈N(A)suchthatW1=W2∩D,andsinceDisdense,onehasWB1=WB2.Thus,WBC1=WBC2,andthereforeInt(WBC2)⊆
WBC2∩D=
WBC1∩D⊆
V1∩D=¯V1⊆V.(I.12)
10G.Luk´acs/NotesondualitytheoriesofabeliangroupsForU∈N(A)suchthatU+U⊆W2,onehasUBC+UBC⊆WBC2,andsoUBC⊆Int(WBC2)(cf.PropositionA.1(a)).Hence,wefoundU∈N(A)suchthatUBC⊆V,asdesired.
I.20.Corollary.Thecompletionofeverylocallyquasi-convexgroupislocallyquasi-convex.
I.21.Precompactness.AsubsetS⊆GofG∈Grp(Top)issaidtobeprecompactifforeveryU∈N(A)thereisaﬁnitesubsetF⊆GsuchthatS⊆FU.Similarlytocompactness,ifS,S0⊆GareprecompactandS1⊆S,thensoareS1,¯S,S+S0,andϕ(S)foreverycontinuoushomomorphismϕ:G→H.Ourinterestinthispropertyarisesfromthefollowingtheoremonuniformspaces(cf.[17]and[13,8.3.16]):
I.22.Theorem.Auniformspace(X,U)iscompactifandonlyifitiscompleteandprecompact.I.23.Proposition.LetA∈Ab(Haus)bealocallyquasi-convexgroup.Then:(a)α−1A(E)isprecompactinAforeveryequicontinuoussubsetE⊆ˆ
ˆA;(b)KBCisprecompactinAforeverycompactsubsetK⊆A;(c)ifαˆAiscontinuous,thenα−1A(F)isprecompactforeverycompactF⊆ˆ
ˆA.Proof.First,wenotethatbyPropositionI.12,α−1A:αA(A)→Aiscontinuous.(a)ByPropositionI.5(c),thereisV∈N(ˆA)suchthatE⊆VB.ByPropositionI.7(a),VBiscompactinˆ
ˆA,andsoVB∩αA(A)isprecompact.Thus,thecontinuoushomomorphicimageα−1A(VB∩αA(A))isalsoprecompact,andcontainsα−1A(E).Therefore,α−1A(E)isprecompact.(b)ThesetKBisopeninˆA,andthusKBBiscompactandequicontinuousinˆ
ˆA(cf.Proposi-tionsI.7(a)andI.5(c)).Hence,by(a),KBC=α−1A(KBB)isprecompactinA.(c)ByPropositionsI.8,inthiscase,{KBB|K∈K(A)}isacobaseforˆ
ˆA,andsothestatementfollowsfrom(b).
I.24.Quasi-convexcompactness.Itisawell-knownthattheclosedconvexhullofaweaklycom-pactsubsetofaBanachspaceisweaklycompact.Infact,thepropertyofpreservationofcompact-nessunderformationofclosedconvexhullcharacterizescompletenessinthecategoryofmetriza-blelocallyconvexspaces(cf.[29,2.4]).Motivatedbythis,onesaysthatA∈Ab(Top)hasthequasi-convexcompactnessproperty(orbrieﬂy,AisQCP)ifforeverycompactsubsetK⊆A,thequasi-convexhullKBC(ofKinA)iscompact.I.25.Proposition.LetA∈Ab(Top)besuchthatαAiscontinuous.ThenˆAhasthequasi-convexcompactnessproperty.
Proof.LetΦ⊆ˆAbecompact.ByPropositionI.7(b),thereisU∈N(A)suchthatΦ⊆UB,andsoΦBC⊆UBBC=UB.Therefore,theclosedsetΦBCisasubsetofthecompactsetUB(cf.PropositionI.7(a)).Hence,theresultfollows.
I.26.Proposition.LetA∈Ab(Haus)bealocallyquasi-convexgroup.If(a)Aiscomplete,or(b)αAissurjective,thenAhasthequasi-convexcompactnessproperty.
G.Luk´acs/Notesondualitytheoriesofabeliangroups11Proof.ByPropositionI.12,α−1A:αA(A)→Aiscontinuous.(a)ByPropositionI.23,KBCisprecompactinA,andit,beingaclosedsubspace,iscomplete.Therefore,byTheoremI.22,itiscompact.(b)SinceKBisopeninˆA,KBBiscompactinˆ
ˆA,andthereforeKBC=α−1A(KBB)iscompact,becauseα−1Aiscontinuous.
I.27.Proposition.LetA∈Ab(Haus)besuchthatαAandαˆAarecontinuous.Thefollowingstatementsareequivalent:(i)α−1A(F)iscompactforeverycompactsubsetF⊆ˆ
ˆA;(ii)Ahasthequasi-convexcompactnessproperty.Proof.(i)⇒(ii)isimmediate,becauseKBBiscompactinˆ
ˆAandKBC=α−1A(KBB).(ii)⇒(i):LetF⊆ˆ
ˆAbecompact.SinceFisclosedandαAiscontinuous,α−1A(F)isalsoclosedinA.ByCorollaryI.8,thereisacompactsubsetK⊆AsuchthatF⊆KBB(becauseαˆAiscontinuous).Thus,α−1(F)⊆α−1A(KBB)=KBC,andKBCiscompactbecauseAisQCP.Therefore,α−1(F),beingtheclosedsubsetofacompactset,iscompact.
I.B.Specialgroupsandsubgroups
I.28.Metrizablegroups.EveryHausdorfftopologicalgroupiscompletelyregular(nicerefer-ence???),sobyTychonoff'smetrizationtheorem,ifaHausdorffgroupissecondcountable,thenitismetrizable.Butfortopologicalgroups,secondcountabilityismorethannecessaryinordertowarrantmetrizability.TheoremI.29below(originallyprovedbyKakutaniandBirkhoffin1936)canalsobeobtainedfromwell-knownresultsonuniformspaces(cf.[13,8.1.10,8.1.21]).I.29.Theorem.EveryﬁrstcountableHausdorfftopologicalgroupismetrizable.
Weturntoabelianmetrizablegroups.RecallthataHausdorffspaceXisak-spaceifF⊆XisclosedinXwheneverF∩KisclosedforeveryK∈K(X).(Fordetails,seeAppendixC.)I.30.Theorem.([10,Theorem1])LetA∈Ab(Met).ThenˆAisak-space.AlthoughTheoremI.30hasanon-commutativegeneralization(cf.[23],[22,3.4]),weprovideheretheoriginalproofbyChasco,enrichedwithafewwordsofexplanation.Proof.InordertoshowthatˆAisak-space,letΦ⊆ˆAbesuchthatΦ∩ΞisclosedinΞforeverycompactsubsetΞ⊆ˆA.Weshowthatforeveryζ6∈Φ,thereexistsacompactsubsetK⊆Asuchthat(KB+ζ)∩Φ=∅,andsoΦisclosedinˆA.Withoutlossofgenerality,wemayassumethatζ=0.SinceAismetrizable,0hasadecreasingcountablebase{Un},andwesetU0=A.Ouraimistoconstructinductivelyafamily{Fn}∞
n=0ofﬁnitesubsetsofAsuchthatforeveryn∈N,Fn⊆Un,(I.13)n\k=0FBk∩UBn+1∩Φ=∅.(I.14)
12G.Luk´acs/NotesondualitytheoriesofabeliangroupsByPropositionI.7(a),UB1iscompactinˆA,andthusbyourassumption,UB1∩ΦisclosedinUB1.Thus,UB1∩ΦiscompactinˆA,andsoitiscompactinthepointwisetopologycarriedbyhom(A,T)(whichiscoarserthanthecompact-openone).Inparticular,UB1∩Φisclosedinthepointwisetopology.Abasicneighborhoodof0inthepointwisetopologyhastheformFB,whereF⊆Aisﬁnite,so06∈UB1∩ΦimpliesthatthereexistsaﬁnitesubsetF0⊆A=U0suchthatFB0∩UB1∩Φ=∅.Thiscompletestheproofforn=0.SupposethatwehavealreadyconstructedF0,...,Fn−1suchthat(I.13)and(I.14)hold.Foreachx∈Un,put∆x=n−1\k=0FBk∩{x}B∩UBn+1∩Φ.(I.15)SinceUn⊇Un+1,onehasUBn⊆UBn+1,andthus,bytheinductivehypothesis,\x∈Un∆x=n−1\k=0FBk∩UBn∩UBn+1∩Φ=n−1\k=0FBk∩UBn∩Φ=∅.(I.16)The∆xareclosedsubsetsofthecompactspaceUBn+1,andtheirintersectionisempty.Therefore,theremustbeaﬁnitesubsetFn⊆UnsuchthatTx∈Fn∆x=∅.Hence,\x∈Fn∆x=n−1\k=0FBk∩FBn∩UBn+1∩Φ=∅,(I.17)asdesired.SetK=∞Sn=0Fn∪{0}.Itfollowsfrom(I.13)thatK\Unisﬁniteforeveryn∈N,andthusKissequentiallycompact(becauseeverysequencewithoutaconstantsubsequencemustconvergeto0∈K).Therefore,Kiscompact,becauseAismetrizable.BytheconstructionofK,KB∩UBn∩Φ=∅foreveryn∈N.ByPropositionI.5(a),ˆA=∞Sn=1UBn,andthereforeKB∩Φ=∅,asdesired.
I.31.Corollary.LetA∈Ab(Met).Then:(a)αAiscontinuous;(b)αˆAiscontinuous;(c)([10,Corollary1])ˆ
ˆAiscompleteandmetrizable.Proof.(a)SinceAismetrizable,itstopologyisdeterminedbyconvergentsequencesxn−→x0.Forsuchasequence,{xn|n∈N}∪{x0}iscompact,andthusAisak-space.Therefore,byTheoremI.9(b),αAiscontinuous.(b)ByTheoremI.30,ˆAisak-space,andso,byTheoremI.9(b),αˆAiscontinuous.(c)ByPropositionI.7(b),{UBB|U∈N(A)}isabaseforˆ
ˆAat0,becauseαAiscontinu-ous.Thus,ˆ
ˆAisﬁrst-countable,becauseAisso.(Recallthatfortopologicalgroups,beingﬁrst-countableisequivalenttobeingmetrizable.)ByTheoremI.30,ˆAisak-space,andthusC(ˆA,T)is
G.Luk´acs/Notesondualitytheoriesofabeliangroups13complete,becauseTisso(cf.[20,7.12]).Sinceˆ
ˆAisaclosedsubspaceofC(ˆA,T),thiscompletestheproof.
AcombinationofPropositionI.27andCorollaryI.31yields:I.32.Corollary.LetA∈Ab(Met).ThenAhasthequasi-convexcompactnesspropertyifandonlyifα−1A(F)iscompactinAforeverycompactF⊆ˆ
ˆA.
I.33.Corollary.([10,Theorem2],[22,3.7])LetA∈Ab(Met),andletD≤Abeitsdensesubgroup.ThenˆD=ˆAastopologicalgroups.Proof.ByTheoremI.19,ˆDandˆAhavethesameunderlyinggroupsandcompactsubsets.Ontheotherhand,byTheoremI.30,bothˆDandˆAarek-spaces,becauseDandAaremetrizable.Therefore,ˆAandˆDhavethesametopology,becausethetopologyofak-spaceisdeterminedbyitscompactsubsets.
I.34.Theorem.LetA∈Ab(Met).Thefollowingstatementsareequivalent:(i)Aislocallyquasi-completeandcomplete;(ii)αAisaclosedembedding;(iii)Ahasthequasi-convexcompactnesspropertyandαAisinjective.OstlingandWilanskyshowedthatalocallyconvexmetrizablevectorspaceiscompleteifandonlyiftheabsolutelyconvexclosureofcompactsubsetiscompact(cf.[29,2.4]).Hern´andezprovedafarreachinggeneralizationofthisresultformetrizablegroups,namely,thatametrizableLQCgroupiscompleteifandonlyifithasQCP(cf.[15,Theorem2]).ThesameresultalsoappearsinapaperbyBrugueraandMart´ın-Peinador,whousedCorollaryI.33inordertosimplifytheproof(cf.[8,9]).ItescapedtheattentionoftheseauthorsthatifαAisinjective,thenQCPissufﬁcienttowarrantcompleteness,sotheﬁrsthalfoftheproofof(iii)⇒(i)appearstobenew,whileitssecondhalfusestheideasofBrugueraandMart´ın-Peinador.Proof.(i)⇒(ii):SinceAisLQC(andmetrizable,soHausdorff),αAisanembeddingbyPropo-sitionI.12.Furthermore,ifS⊆AisclosedinA,thenSiscomplete(becauseAiscomplete),andthusαA(S)iscomplete.Therefore,αA(S)isclosedinˆ
ˆA,asdesired.(ii)⇒(iii):Clearly,αAisinjective.LetK⊆Abecompact.ThenKBBiscompactinˆ
ˆA,andsinceαA(A)isclosed,KBB∩αA(A)isalsocompact.Therefore,KBC=α−1A(KBB∩αA(A))iscompacttoo.(iii)⇒(i):SinceαAisinjective,α−1A:αA(A)→Aiswell-deﬁned.ByCorollaryI.32,QCPimpliesthatα−1A(F)iscompactforeverycompactsubsetF⊆αA(A).So,ifK⊆α−1A(F)isclosed,thenKisalsocompact,andthus(α−1A)−1(K)=αA(K)isacompactsubsetofF,becauseαAiscontinuous.Therefore,(α−1A)−1(K)isclosed,andhence,α−1Aisk-continuous.ByCorollaryI.31(c),ˆ
ˆAismetrizable,andsoisitssubspaceαA(A).Inparticular,αA(A)isak-space,andthereforeα−1Aiscontinuous.ThisshowsthatαAisanembedding,andhence,byPropositionI.12(a),AisLQC.InordertoshowthatAiscomplete,letBbethecompletionofA.ByCorollaryI.33,ˆA=ˆB(astopologicalgroups),becauseAisadensesubgroupofthemetrizablegroupB.Let{an}bea
14G.Luk´acs/NotesondualitytheoriesofabeliangroupsCauchy-sequenceinA.Thenan−→bforsomeb∈B,andK1={an|n∈N}∪{b}iscompactinB.So,KB1∈N(ˆB)=N(ˆA),andbyPropositionI.5(b),thereisacompactsubsetK⊆AsuchthatKB⊆KB1.Thus,{an|n∈N}⊆KBC1⊆KBC(where(−)CistakenwithrespecttoA).SinceAhasQCP,KBCiscompact,andsosequentiallycompact(becauseAismetrizable).Therefore,{an}hasaconvergentsubsequenceinKBC,whichshowsthatb∈KBC⊆A.Hence,Aiscomplete,asdesired.
I.35.Compactanddiscretegroups.LetA∈Ab(Top).IfAiscompact,thenbyProposi-tionI.5(b),{0}=AB∈N(ˆA).Ontheotherhand,ifAisdiscrete,then{0}∈N(A),andsobyPropositionI.7(a),ˆA={0}Biscompact.RepeatingtheseargumentsforˆAyields:I.36.Lemma.LetA∈Ab(Top).(a)IfAiscompact,thenˆAisdiscrete.(b)IfAisdiscrete,thenˆAiscompact.(c)IfAiscompact,thensoisˆ
ˆA.(d)IfAisdiscrete,thensoisˆ
ˆA.
I.37.Theorem.LetA∈Ab(Top).IfAis(a)discrete,or(b)compactHausdorff,thenαAisanisomorphismoftopologicalgroups.Proof.(a)Leta∈Abeanon-zeroelement,andletr∈Tbesuchthato(a)=o(r).Thisdeﬁnesagrouphomomorphismχa:hai→Tsuchthatχa(a)6=0.SinceTisaninjectiveabeliangroup,χaextendsto¯χa:A→T,and¯χa∈ˆA,becauseAisdiscrete.Thus,αa(¯χa)=¯χa(a)6=0,whichshowsthatαAisinjective.ByLemmaI.36(d),ˆ
ˆAisdiscrete,andthusαAisanembedding.SupposethatAisﬁnitelygenerated.ThenitdecomposesintothedirectsumofcyclicgroupsA=ha1i⊕···⊕hani.Finitecyclicgroupsareself-dual,andˆZ=TandˆT=Z.Thus,surjectivityofαAfollows,becausethePontryagindualisanadditivefunctor.(Moredetails???)Inthegeneralcase,letξ∈ˆ
ˆA.ByPropositionI.5(a),thereisW∈N(ˆA)suchthatξ(W)⊆Λ1,andWcanbechosentohavetheformFBwhereF⊆Aisﬁnite.PutB=hFi.ByLemmaI.4,BBisasubgroupofˆA,andBBB∼=\ˆA/BB.SinceTisaninjectiveabeliangroup,everyhomomor-phismχ:B→Textendstoahomomorphism¯χ:A→T,andthusthedualˆιB:ˆA→ˆBoftheinclusionιB:B→Aissurjective.ThemapˆιBisalsoclosed,becauseˆAiscompact.Therefore,ˆB∼=ˆA/BB,andhenceBBB∼=\ˆA/BB∼=ˆ
ˆB.SinceF⊆B,onehasξ(BB)⊆ξ(FB)⊆Λ1,andsoξ∈BBB∼=ˆ
ˆB.ThegroupBisﬁnitelygenerated,andthusbywhatwehaveshownsofar,thereisb∈Bsuchthatξ=αB(b).Therefore,ξ=αA(b).(b)SinceAiscompactHausdorff,itisak-space,andsoαAiscontinuous(cf.TheoremI.9).BythePeter-Weyltheorem,foreverynon-zeroa∈A,thereisχ∈ˆAsuchthatχ(a)6=0(cf.[30,Thm.32]).Thus,αAisinjective,andthereforeitisaclosedembedding(becauseAiscompact).InordertoshowthatαAissurjective,assumethecontrary,thatis,assumethatαA(A)isaproperclosedsubgroupofˆ
ˆA.ByLemmaI.36(c),ˆ
ˆAiscompact,andthussoisthequotient
G.Luk´acs/Notesondualitytheoriesofabeliangroups15ˆ
ˆA/αA(A).BythePeter-Weyltheoremappliedtoˆ
ˆA/αA(A),thereexistsanon-zerocontinuouscharacterψ:ˆ
ˆA/αA(A)→T,whichinducesζ∈ˆ
ˆ
ˆAsuchthatζ(αA(A))={0}andζ6=0.SinceˆAisdiscrete(cf.LemmaI.36(a)),by(a),αˆAisanisomorphism.Inparticular,thereisanon-zeroχ∈ˆAsuchthatαˆA(χ)=ζ.Therefore,fora∈A,χ(a)=(αA(a))(χ)=(αˆA(χ))(αA(a))=ζ(αA(a))∈ζ(αA(A))={0}.(I.18)Hence,χ=0(andsoζ=0)contrarytoourassumption,whichcompletestheproof.
I.38.Bohr-compactiﬁcation.TheˇCech-compactiﬁcationβXofaTychonoffspaceXhasthepropertythateverycontinuousmapf:X→KintoacompactHausdorffspaceXextendstoβX,andthusfactorsuniquelythroughthedenseembeddingX→βX.Itappearstobelessknown,however,thatβXexistsevenwhenXisnotTychonoff,inwhichcasethemapX→βXisonlycontinuousbutneednotbeinjectiveoropenontoitsimage.(EverycontinuousmapofXintoacompactHausdorffspacestillfactorsuniquelythroughβX.)Inacategoricallanguage,onesaysthatthecategoryHCompofcompactHausdorffspacesandtheircontinuousmapsisareﬂectivesubcategoryofTop.SimilarlytotherelationshipbetweenTopandHComp,thefullsubcategoryGrp(HComp)ofcompactHausdorffgroupsisreﬂectiveinGrp(Top)(topologicalgroupsandtheircontinuousho-momorphisms).ThereﬂectioniscalledtheBohr-compactiﬁcation,andisdenotedbyρG:G→bG.Weshowitsexistenceforabeliangroups,wheretheconstructionisrathersimpleI.39.Theorem.ForeveryA∈Ab(Top)thereisbA∈Ab(HComp)andacontinuoushomomor-phismρA:A→bAsuchthateverycontinuoushomomorphismϕ:A→KintoK∈Ab(HComp)factorsuniquelythroughρA:AbAρA::::::::AKϕ//
KbAAA∃!˜ϕ(I.19)Moreover,cbA=(ˆA)d,andkerρA=kerαA.Proof.SincebAissupposedtobecompact,byLemmaI.36(a),cbAmustbediscrete.ApplyingtheuniversalpropertyofbAtoK=Tandϕ=χ∈ˆAyieldsthatAandbAhavethesamecontinuouscharacters,thatis,cbA=(ˆA)d(thesubscriptdstandsforthediscretetopology).Usingthisnecessaryconditionasadeﬁnition,onesetbA=d(ˆA)dand(ρA(a))(χ)=χ(a)foreverya∈Aandχ∈(ˆA)d.ByPropositionI.5(b),abasicopensetinbAhastheformΦB,whereΦ⊆ˆAisﬁnite.Thus,ρ−1A(ΦB)=Tχ∈Φχ−1(Λ1)isaneighborhoodof0inA,andthereforeρAiscontinuous.Ifϕ:A→Kisacontinuoushomomorphism,thenitinducesˆϕ:ˆK→ˆA.SinceˆKisdiscrete(cf.LemmaI.36(a)),(ˆϕ)d:ˆK→(ˆA)disalsocontinuous,andsod(ˆϕ)d:d(ˆA)d→ˆ
ˆKiscontinuousaswell.ByTheoremI.37,ˆ
ˆK∼=K,andtherefore˜ϕ=d(ˆϕ)d.Thisconsiderationshowstheuniquenessof˜ϕtoo,because˜ϕisuniquelydeterminedbyitsdual.SinceαA(a)andρA(a)arebothevaluationsata,thelaststatementfollows.
16G.Luk´acs/NotesondualitytheoriesofabeliangroupsAsubgroupH≤AsuchthatHBC=HissaidtobeduallyclosedinA;ifeverycontinuouscharacterofHextendscontinuouslytoA,wesaythatHisduallyembeddedinA.WesawinTheoremI.19(a)thateverydensesubgroupofatopologicalgroupisduallyembedded.I.40.Corollary.LetA∈Ab(Top)besuchthatαAisinjective,andletKbeacompactsubgroupofA.ThenKisduallyclosedandduallyembeddedinA.Proof.SinceαAisinjective,AmustbeHausdorff.First,supposethatAiscompact.InordertoshowthatKisduallyclosed,letx∈A\K.ThequotientA/Kisalsocompact,andtheimage¯xofxinA/Kisnon-zero.Thus,byTheoremI.37,thereis¯χ∈[A/Ksuchthat¯χ(¯x)6=0.Thecharacter¯χinducesχ∈ˆAbysettingχ(a)=¯χ(¯a)(where¯astandsfortheimageofainA/K),andK⊆kerχ,whileχ(x)6=0.Therefore,χ∈KB,andhencex6∈KBC.(NotethatbyLemmaI.4,KBisasubgroupofˆA,andsoKBCcoincideswithitsannihilatorinA.)ThisshowsthatK=KBC.InordertoshowthatKisduallyembedded,considerthediscretegroupˆA.ByLemmaI.4(b),ˆπK:\ˆA/KB→ˆ
ˆAisinjectiveanditsimageisKBB.Sinceˆ
ˆA∼
=A(cf.TheoremI.37),KBB∼=KBC=K.ThequotientˆA/KBisdiscrete,andso\ˆA/KBiscompact.Therefore,ˆπK:\ˆA/KB→AisanembeddingontoK,thatis,K∼=\ˆA/KB.Hence,applyingTheoremI.37tothecompactgroupˆA/KByieldsˆK∼=\
\ˆA/KB∼=ˆA/KB.Inparticular,ˆιA:ˆA→ˆKgivenbyrestrictiontoKissurjective,asdesired.Inthegeneralcase,sinceρAcontinuous,ρA(K)isacompactsubgroupofbA.Therefore,bywhatwehaveshownsofar,itisduallyclosedandduallyembeddedinbA.Thus,foreveryx∈bA\ρA(K),thereisχ∈cbA=(ˆA)dsuchthatχ(ρA(K))=0andχ(x)6=0.Inparticular,thisholdsforeveryx=ρA(a),wherea∈A\K.(SinceρAisinjective,a∈A\Kimpliesx∈bA\ρA(K).)Therefore,χ∈KB,andsoa6∈KBC.Hence,K=KBC,asdesired.Forthesecondstatement,observethatK∼
=ρA(K)(becauseρAisinjective,Kiscompact,andbAisHausdorff).Thus,ifχ∈ˆK=\ρA(K),thenitadmitsacontinuousextensionψtobA,becauseρA(K)isduallyembeddedinbA.Therefore,ψ◦ρA∈ˆAisacontinuousextensionofχ.Hence,KisduallyembeddedinA,asdesired.
I.41.Openandcompactsubgroups.OpensubgroupsofabeliangroupsarecloselyrelatedtocompactsubgroupsofthePontryagindual,andviceversa.Weturntoestablishingthepropertiesofopenandcompactsubgroupofabeliangroups.OurpresentationoftheseresultswasstronglyinﬂuencedbytheworkofBanaszczyk,Chasco,Mart´ın-Peinador,andBruguera(cf.[4]and[7]).I.42.Proposition.([4,2.2])LetA∈Ab(Top),U≤Abeanopensubgroup,andconsidertheexactsequence0//
UιU//
AπU//
A/U//
0.(I.20)(a)ThemapˆπU:dA/U→ˆAisanembeddinganddA/U∼=UB.(b)ThesubgroupUisduallyembeddedinA.Inotherwords,ˆιU:ˆA→ˆUissurjective.(c)ˆιUisopen,andthusaquotient.
G.Luk´acs/Notesondualitytheoriesofabeliangroups17(d)Theinducedsequence0//
dA/UˆπU//
ˆAˆιU//
ˆU//
0(I.21)isexact.Proof.NotethatA/Uisdiscrete,becauseUisopen,andsodA/UiscompactbyLemmaI.36(b).(a)ByLemmaI.4(b),ˆπUisinjective,anditsimageisUB.Thus,ˆπUisanembedding,becauseitiscontinuous,itsdomainiscompactanditscodomainisHausdorff.(b)Letχ∈ˆU,andletV∈N(U)besuchthatχ(V)⊆Λ1(cf.PropositionI.5(a)).SinceTisaninjectiveabeliangroup,χadmitsanextensionψ:A→T.Thus,ψ(V)=χ(V)⊆Λ1,andV∈N(A),becauseUisopeninA.Therefore,byPropositionI.5(a),ψiscontinuousonA.(c)Inordertomakeheproofmoretransparent,wedecomposeditintothreesimplesteps.Step1:A=hU,biforsomeb∈A.LetlbetheorderofbinA/U(possiblyinﬁnite).GivenacompactsubsetK⊆A,itcanbecoveredbyﬁnitelymanytranslatesk1b+U,...,kmb+UofU,andwithoutlossofgeneralitywemayassumethat0≤ki<l.SinceUisanopensubgroup,itisalsoclosed,andso(K−kib)∩Uiscompact.Thus,C=((K−k1b)∪...∪(K−kmb))∩UisacompactsubsetofU,andK⊆(k1b+C)∪...∪(kmb+C).Withoutlossofgenerality,wemayassumethatlb∈Cifl6=∞.SetW={χ∈ˆU|χ(C)⊆Λ2},andweshowthatW⊆ˆιU(KB).Tothatend,letχ∈W.If¯χ∈ˆAisanextensionofχ,then¯χ(K)⊆¯χ(k1b+C)∪...∪¯χ(kmb+C)⊆(k1¯χ(b)+Λ2)∪...(km¯χ(b)+Λ2).(I.22)Iflisﬁnite,thenlb∈C,andsoχ(lb)∈Λ2.Letr∈Λ2l⊆Tbetheclosestpointto0suchthatlr=χ(lb)(ifl=∞,thenr=0).Set¯χ(u+nb)=χ(u)+nr.ThecharacterχiscontinuousonAbecause¯χ|
|U=χiscontinuous,andUisanopensubgroup.Furthermore,onehaski¯χ(b)∈Λ2foreachi,because0≤ki<landr∈Λ2l.Therefore,by(I.22),¯χ(K)⊆Λ2+Λ2⊆Λ1,whichmeansthat¯χ∈KB.Hence,W⊆ˆιU(KB),asdesired.Step2:A/Uisﬁnitelygenerated,thatis,A=hU,b1,...,bniforb1,...bn∈A.SetU0=UandUk=hUk−1,bki.EachUkisanopensubgroupofUk+1(andofA),andthusbystep1,eachˆιUk:ˆUk+1→ˆUkisopen.Therefore,ˆιU=ˆιUk−1◦...◦ˆιU0isopen.Step3:Inthegeneralcase,letK⊆Abecompact.ThenKcanbecoveredbyﬁnitelymanytranslatesofU,K⊆(b1+U)∪···∪(bn+U).PutU0=hU,b1,...,bni.SinceK⊆U0,theimageˆιU0(KB)coincideswiththepolarofKinˆU0withrespecttoU0,andthusopen.Bystep2,ˆιU:ˆU0→ˆUisopen.Therefore,ˆιU(KB)isopeninˆU,asdesired.(d)ExactnessatdA/UandˆUwereshownin(b)and(c),respectively.ExactnessatˆAalsofollowsfrom(b),becauseImˆπU=UB=kerˆιU.
I.43.Proposition.LetA∈Ab(Top),K≤Abeacompactsubgroup,andconsidertheexactsequence0//
KιK//
AπK//
A/K//
0.(I.23)(a)ThemapˆπK:[A/K→ˆAisanembeddingand[A/K∼=KB.
18G.Luk´acs/Notesondualitytheoriesofabeliangroups(b)ThemapˆιK:ˆA→ˆKisopenontoitsimage.(c)Theinducedsequence0//
[A/KˆπK//
ˆAˆιK//
ˆK(I.24)isexact.(d)IfαAisinjective,thenˆιK:ˆA→ˆKissurjective,andsoˆιKisaquotientmap.Proof.NotethatbyLemmaI.36(a),ˆKiscompact.(a)Abasicneighborhoodof0in[A/KisoftheformLB,whereL⊆A/Kiscompact(cf.PropositionI.5(b)).ByLemmaA.3(b),π−1K(L)isacompactsubsetofA,andthusπ−1K(L)Bisabasicneighborhoodof0inˆA.Withoutlossofgenerality,wemayassumethat0∈L,andthenbyLemmaI.4(c),π−1K(L)B=ˆπK(LB),asdesired.(b)SincekerˆιK=KBisopeninˆAandˆKisdiscrete,thestatementisobvious.(c)Exactnessat[A/KandatˆAfollowsfrom(b),becauseImˆπK=KB=kerˆιK.(d)ByCorollaryI.40,KisduallyembeddedinA,whichcompletestheproof.
I.44.Theorem.([4,2.3])LetA∈Ab(Top),andletU≤Abeanopensubgroup.Then:(a)αUisinjective(resp.,surjective)ifandonlyifαAisinjective(resp.,surjective);(b)αUisanisomorphismoftopologicalgroupsifandonlyifαAisso.Proof.ByPropositionI.42,theexactsequence0//
UιU//
AπU//
A/U//
0(I.25)givesrisetoanexactsequence0//
dA/UˆπU//
ˆAˆιU//
ˆU//
0(I.26)withdA/Ucompact,ˆπUanembedding,andˆιUaquotientmap.Thus,theconditionsofProposi-tionI.43arefulﬁlled(byRemarkI.13(3),αˆAisinjective),andso0//
ˆ
ˆUˆ
ˆιU//
ˆ
ˆAˆ
ˆπU//
d dA/U//
0(I.27)isalsoexact,ˆ
ˆιUisanembedding,andˆ
ˆπUisaquotientmap.Sincetheevaluationhomomorphismαisanaturaltransformation,weobtainacommutativediagramwithexactrows:0U//
UAιU//
AA/UπU//
A/U0//
0ˆ
ˆU//
ˆ
ˆUˆ
ˆAˆ
ˆιU//
ˆ
ˆAd dA/Uˆ
ˆπU//
d dA/U0//
Uˆ
ˆUαU
Aˆ
ˆAαA
A/Ud dA/UαA/U
(I.28)ThegroupA/Uisdiscrete,andthusαA/UisatopologicalisomorphismbyTheoremI.37.There-fore,(a)followsfromthewell-knownFiveLemmaforabeliangroups.Inordertoshow(b),observethatαU=αA|
|Uandα−1U=α−1A|
|U,andthatitsufﬁcestocheckthecontinuityofahomo-morphismonaneighborhoodof0.SinceUisopen,thestatementfollows.
G.Luk´acs/Notesondualitytheoriesofabeliangroups19I.45.Theorem.([4,2.6])LetA∈Ab(Haus)besuchthatαAisinjective,andletK≤Abeacompactsubgroup.Then:(a)αA/Kisinjective;(b)αA/KissurjectiveifandonlyifαAisso;(c)αA/KisanisomorphismoftopologicalgroupsifandonlyifαAisso.I.46.Remark.AlthoughinPropositionI.43weassumednoseparationaxioms,herewedo,be-causeαAbeinginjectiveimpliesthatAisHausdorff.Proof.SinceαKisinjective,byPropositionI.43,theexactsequencein(I.23)givesrisetoanexactsequence0//
[A/KˆπK//
ˆAˆιK//
ˆK//
0(I.29)where[A/KisanopensubgroupofˆAandˆιKisaquotient.Therefore,byPropositionI.42,thelowerrowofthecommutativediagrambelowisexact:0K//
KAιK//
AA/KπK//
A/K0//
0ˆ
ˆK//
ˆ
ˆKˆ
ˆAˆ
ˆιK//
ˆ
ˆA[
[A/Kˆ
ˆπK//
[
[A/K0//
Kˆ
ˆKαK
Aˆ
ˆAαA
A/K[
[A/KαA/K
(I.30)SinceKiscompactHausdorff,αKisanisomorphismoftopologicalgroups(cf.PropositionI.37),andthus(a)and(b)followfromtheFiveLemmaforabeliangroups.(c)IfαAisanisomorphismoftopologicalgroups,thenαA/Kisabijectionby(a)and(b),andthetopologiesofA/Kand[
[A/Kcoincide,because[
[A/K∼
=ˆ
ˆA/ˆ
ˆK∼=A/K.Conversely,supposethatαA/Kisatopologicalisomorphism.ThenαAisbijectiveby(a)and(b).WeshowthatαAiscontinuous.Tothatend,letFbeaﬁlterconvergingto0inA.ThenπK(F)−→K,andsoˆ
ˆπK(αA(F))=αA/K(πK(F))−→ˆ
ˆK.ByLemmaA.3(a),theﬁlterαA(F)hasaclusterpointx∈ˆ
ˆK⊆ˆ
ˆA,becauseˆ
ˆKiscompact.So,letG⊇FbeaﬁltersuchthatαA(G)−→x.Forχ∈ˆA,onehasx(χ)=limαA(G)(χ)=limχ(G)=χ(limG)=0(I.31)becauseχiscontinuousandG−→0,andthusx=0.Therefore,0istheuniqueclusterpointofαA(F)foreveryﬁlterF−→0.Hence,αA(F)−→0foreveryF−→0,andαAiscontinuous.Sincethesituationiscompletelysymmetric(onecouldinvertαK,αA,andαA/K),theproofofthecontinuityofα−1Aissimilar.
I.47.Locallycompactgroups.WeconcludethischapterwithmakingasteptowardtheclassicalPontryagindualityforlocallycompactHausdorffabelian(LCA)groups.WedenotebyLCthecategoryoflocallycompactHausdorffspaces,andthusAb(LC)isthecategoryofLCAgroups.I.48.Proposition.LetA∈Ab(LC).Then:(a)ˆA∈Ab(LC);(b)αAiscontinuous;(c)αAisinjective.
20G.Luk´acs/NotesondualitytheoriesofabeliangroupsProof.(a)Clearly,ˆAisHausdorff.LetU∈N(A)suchthat¯Uiscompact.ByPropositionI.7(a),UBiscompactinˆA,andbyPropositionI.5(a),¯UBisabasicneighborhoodof0inˆA.Since¯UB⊆UB,thiscompletestheproof.(b)LetV∈N(A)besuchthat¯Viscompact.ByTheoremI.9,αAisk-continuous,andsoαA|
|¯Viscontinuous.Inparticular,αAiscontinuousat0,andthereforeitiscontinuous.(c)Leta∈Abeanon-zeroelement.First,supposethatAisgeneratedbyV∈N(A)suchthat¯Viscompact.ByreplacingVwithV∪(−V)∪{a,−a}ifnecessary,wemayassumethata∈Vand−V=Vfromtheoutset.By[33,Lemma2.42],thereisasubgroupH≤AsuchthatH∩V={0},H∼=Znforsomen∈N,andA/Hiscompact.Theimagea+HofainA/Hisnon-zero,becausea∈VandV∩H={0}.Thus,byTheoremI.37(b),αA/H(a+H)6=0,andsothereis¯χ∈[A/Hsuchthat¯χ(a+H)6=0.Therefore,onehasαA(a)(χ)=χ(a)=¯χ(a+H)6=0,whereχ=ˆπH(¯χ)∈ˆA.Hence,αA(a)6=0.Inthegeneralcase,pickV∈N(A)suchthat−V=V,a∈V,and¯Viscompact,andputU=hVi,thesubgroupgeneratedbyV.ThenU+V⊆U,andthus¯U⊆IntU(byProposi-tionA.1(a)).Therefore,UisanopensubgroupofA.Bywhatweprovedsofar,αUisinjective,Hence,byTheoremI.44(a),αAisinjective.
I.49.Remark.Theproofof(c)fallsshortofbeingself-containedbecauseitisbasedonaLemmafrom[33].IhopetoﬁndasimpleproofoftherelevantpartsoftheLemma.Suggestions???I.50.Lemma.LetA∈Ab(LC),andsupposethatthereisacountablefamily{Un}⊆N(A)ofneighborhoodsof0inAsuchthatTUn={0}.ThenAismetrizable.Proof.ByreplacingeachUnwithU0n∈N(A)suchthat¯U0n⊆Unand¯U0niscompactifnecessary,wemayassumethatT¯Un={0},and¯U1iscompact(suchU0nexists,becauseAisregular).Furthermore,withoutlossofgenerality,wemayassumethatUn+1⊆Un.ByTheoremI.29,itsufﬁcestoshowthat{Un}isabaseat0forA.Tothatend,letV∈N(A).ThenT(¯Un\IntV)=∅,and{¯Un\IntV}isadecreasingfamilyofclosedsubsetsofthecompactspace¯U1.Therefore,thereisn0suchthat¯Un0\IntV=∅.Hence,¯Un0⊆IntV,asdesired.
I.51.Lemma.LetA∈Ab(LC),andletV∈N(A)besuchthat¯Viscompact.ThenVcontainsacompactsubgroupKofAsuchthatA/Kismetrizable.Proof.Weconstructafamily{Vn}ofneighborhoodsof0inA.PutV1=(−V)∩V,andforn∈N,usingcontinuityofadditioninA,pickVn+1∈N(A)suchthatVn+1+Vn+1⊆Vnand−Vn+1=Vn+1.PutK=∞Tn=1Vn;KisasubgroupofA,becauseK−K⊆Vn+1−Vn+1⊆VnimpliesK−K⊆K.OnehasK+Vn+1⊆Vn+1+Vn+1⊆Vn,andthus¯K⊆IntVn(cf.PropositionA.1(a)).Therefore,¯K⊆K⊆¯V,andKiscompact.TheinclusionK+Vn+1⊆VnalsoyieldsK=∞Tn=1(K+Vn),whichmeansthat∞Tn=1πK(Vn)=KinA/K.Hence,{πK(Vn)}satisﬁestheconditionsofLemmaI.50,andA/Kismetrizable,asdesired.
I.52.Theorem.Thefollowingstatementsareequivalent:
G.Luk´acs/Notesondualitytheoriesofabeliangroups21(i)ForeverysecondcountableA∈Ab(LC),αAisanisomorphismoftopologicalgroups.(ii)ForeveryA∈Ab(LC),αAisanisomorphismoftopologicalgroups.Proof.(i)⇒(ii):LetA∈Ab(LC).First,supposethatAisgeneratedbyV∈N(A)suchthat¯Viscompact.ByreplacingVwith−V∪Vifnecessary,wemayassumethat−V=Vfromoutset.ThenA=∞Sn=1¯V+···+¯V|
{z
}ntimes,soAisσ-compact,andinparticular,AisLindel¨of.LetK⊆VbeacompactsubgroupsuchthatA/Kismetrizable(cf.LemmaI.51).ThequotientA/Kissecondcountable,becauseitisacontinuousimageoftheLindel¨ofspaceA.(ThepropertiesLindel¨of,separable,andsecondcountableareequivalentformetrizablespaces.)Thus,αA/Kisanisomorphismoftopologicalgroupsbyourassumption.ByPropositionI.48(c),αAisinjective,andtherefore,byTheoremI.45(c),αAisanisomorphismoftopologicalgroups.Inthegeneralcase,pickV∈N(A)suchthat¯Viscompact,andputU=hVi,thesubgroupgeneratedbyV.ThenU+V⊆U,andthus¯U⊆IntU(byPropositionA.1(a)).Therefore,UisanopensubgroupofA.Bywhatweprovedsofar,αUisanisomorphismoftopologicalgroups,becauseUisgeneratedbyacompactneighborhoodof0.Hence,byTheoremI.44(b),αAisanisomorphismoftopologicalgroups.
Appendix
A.Separationpropertiesoftopologicalgroups
ForG∈Grp(Top),oneputsN(G)forthecollectionofneighborhoodsofeinG.A.1.Proposition.LetG∈Grp(Top).(a)IfVW⊆UforU,V,W∈N(G),then¯V,¯W⊆IntU.(b)Gisregular.(c)Thefollowingstatementsareequivalent:(i)GisT3(i.e.,T1andregular);(ii)GisHausdorff;(iii)GisT1;(iv)GisT0.Proof.(a)SetW1=(IntW)−1,andletx∈¯V.ThenW1∈N(G)(becausethegroupinversioniscontinuous),andV∩xW16=∅.Thus,thereisv∈Wandw1∈W1suchthatxw1=v,orx=vw−11.Therefore,x∈VW−11=V(IntW).SinceV(IntW)isopen,itiscontainedinInt(VW)⊆IntU.Hence,x∈IntU,asdesired.(b)Bycontinuityofthemultiplicationgroupm:G×G→G,foreveryU∈N(G)thereisV∈N(G)suchthatVV⊆U,andthus¯V⊆IntUby(a).(c)Implications(i)⇒(ii)⇒(iii)⇒(iv)areobvious.Inordertocompletetheproof,observethatby(b),Gisregular,andeveryT0regulartopologicalspaceXisT1.(Indeed,ifx,y∈Xaredistinctpoints,thenoneofthem,sayx,hasaneighborhoodU0suchthaty6∈U0.ThenF=X\IntU0isaclosedsubsetcontainingythatdoesnotcontainx.SinceXisregular,therearedisjointopensubsetsUandVofXsuchthatF⊆Uandx∈V.Thiscompletestheproof.)
A.2.Proposition.LetG∈Grp(Top),andsetNG=TN(G).Then:(a)NGisacompactnormalsubgroupofG;(b)G/NGisHausdorff;(c)foreverycontinuoushomomorphismϕ:G→HintoaHausdorffgroupH,NG⊆kerϕ.Inparticular,ϕfactorsuniquelythroughG→G/NG.Proof.(a)LetU∈N(G)andg∈G.Bycontinuityofthegroupoperations,thereisV∈N(G)suchthatVV−1⊆Uandg−1Vg⊆U.Thus,NGN−1G⊆VV−1⊆Uandg−1NGg⊆g−1Vg⊆U.ThisistrueforeveryU∈N(G),andthereforeNGN−1G⊆NGandg−1NGg⊆NG.Hence,NGisanormalsubgroup.SinceNGiscontainedineveryneighborhoodofe,itscompactnessisclear.
G.Luk´acs/Notesondualitytheoriesofabeliangroups25(b)ForeveryU∈N(G),thereisV∈N(G)suchthatVV⊆U,andthusVNG⊆VV⊆U.Thus,NG⊆T{VNG|V∈N(G)}⊆TN(G)=NG.Therefore,G/NGisT0,andhenceitisHausdorff,byPropositionA.1.(c)IfHisHausdorff,then{eH}=TN(H),andsokerϕ=\{ϕ−1(W)|W∈N(H)}⊇\N(G)=NG,(1)asdesired.
ThegroupG/NGisthemaximalHausdorffquotientofG.Incategoricallanguage,thismeansthatG/NGisthereﬂectionofG∈Grp(Top)inGrp(Haus).A.3.Lemma.LetG∈Grp(Top),K≤Gbeacompactsubgroup,andputπK:G→G/Kforthecanonicalprojection.(a)IfFisaﬁlterinGsuchthatKisaclusterpointofπK(F)inG/K,thenFhasaclusterpointinK.(b)IfL⊆G/Kiscompact,thensoisπ−1K(K).Proof.(a)AssumethatFhasnoclusterpointinK.Theneachpointx∈KhasaneighborhoodUxsuchthatUxdoesnotmeshF,thatis,thereisFx∈FsuchthatUx∩Fx=∅.Thecollection{Ux}x∈KisanopencoverofK,andthusithasaﬁnitesubcoverUx1,...,Uxl,becauseKiscompact.SetU=Ux1∪...∪UxlandF=Fx1∩...∩Fxl.SinceFisaﬁlter,F∈F,andonehasU∩F=∅andK⊆U.Usingthegroupmultiplicationm:G×G→G,thelastinclusioncanbeexpressedas{e}×K⊆m−1(U).Sincemiscontinuous,m−1(U)isopen,andthus,byatube-lemmatypeargument,thereisV∈N(G)suchthatV×K⊆m−1(U).Inotherwords,VK⊆U.Therefore,VK∩F=∅,andsoVK∩FK=∅.Hence,πK(V)∩πK(F)=∅.SinceπK(V)isaneighborhoodofKinG/K,thiscontradictstheassumptionthatKisaclusterpointofπK(F).(b)LetFbeaﬁltermeshingπ−1K(L)(i.e.,F∩π−1K(L)6=∅foreveryF∈F).ThenπK(F)isaﬁlterthatmeshesL,andthushasaclusterpointxK,becauseLiscompact.Equivalently,KisaclusterpointofπK(x−1F),andthereforex−1FhasaclusterpointinK(by(a)).Hence,FhasaclusterpointinxK,whichshowsthatπ−1K(L)iscompact.
AcombinationofPropositionA.2(a)andLemmaA.3yields:A.4.Corollary.CompactsubsetsofG/NGarepreciselytheimagesofcompactsubsetsofG.
B.Exponentiabilityandthecompact-opentopology
LetTbeafullsubcategoryofTop,thecategoryoftopologicalspacesandtheircontinuousmaps.SupposethatThasﬁniteproducts;theymaydifferfromtheproductsinTop,buteachproductinTmusthavethesameunderlyingsetasinTop.Eachfunctionf:X×Y→Zgivesrisetoamap¯f:Y→ZXdeﬁnedby¯f(y)=fy,wherefy∈ZXisgivenbyfy(x)=f(x,y).OnewouldliketoequipthefunctionsetT(X,Z)⊆ZXwithasuitabletopologysuchthat¯fiscontinuouswhenever
26G.Luk´acs/Notesondualitytheoriesofabeliangroupsfisso.Ifweaddthenaturalrequirementoffunctorialityofthetopology,wearriveattheconceptofexponentiability.AspaceX∈Tisexponentiable(inT)ifthefunctorX×−:T−→Thasarightadjoint.SinceweassumeTtohaveﬁniteproducts,ithasaterminalobjectwhichmustthenbetheone-pointspace{∗},becauseTisafullsubcategory.Thus,oneexpectsT(X,Z)tobetheunderlyingsetofthehom-object.Inotherwords,XisexponentiableifthereisawayoftopologizingthesetT(X,Z)foreveryZ∈TsuchthattheresultingspaceisinT,andthemapT(X×Y,Z)−→T(Y,T(X,Z))(2)f7−→¯fisabijectionthatisnaturalinYandZ.ThequestionofdescribingexponentiableHausdorffspaceswasraisedbyHurewiczinaper-sonalcommunicationwithFox,whowastheﬁrsttogiveapartialanswertothequestion(cf.[14]).HeprovedthatforeveryregularlocallycompactspaceXandeveryY∈HausthesetHaus(X,Y)canbetopologizedinawaythatfiscontinuousifandonlyif¯fisso.FoxalsoshowedthatforeveryseparablemetricspaceX,itispossibletotopologizeHaus(X,R)inawaythat(2)holdsforeveryYifandonlyifXislocallycompact.AlthoughFoxdidnotprovethattheconditionoflocalcompactnessissufﬁcient,andheshowednecessityonlyforseparablemetricspacesX,hewasactuallyveryclosetoacompletesolution.ForHausdorffspacesXandY,wedenotebyC(X,Y)thespaceHaus(X,Y)equippedwiththecompact-opentopology:Itssubbaseisthefamily{[K,U]|K⊆Xcompact,U⊆Yopen},where[K,U]={f∈Haus(X,Y)|f(K)⊆U}.Thecompact-opentopologywasalsoinventedbyFox,andfollowinghiswork,ArensstudiedtheseparationpropertiesofC(X,Y)(cf.[1]).ArenswasnotfarfromprovingthatlocalcompactnessisnecessaryforexponentiabilityinHaus(cf.[1,Theorem3]).JacksonprovedthatifXislocallycompact,thenC(X,−)istherightadjointofX×−,moreover,thebijectionC(X×Y,Z)←→C(Y,C(X,Z))(3)isactuallyahomeomorphismforeveryHausdorffspaceYandZ(cf.[19]).ItappearsthatFoxalreadyrecognizedthat(3)isabijection,andJackson'smainachievementisprovingthatitisahomeomorphism.
B.1.Remark.Someauthors,includingIsbell,erroneouslycreditBrownforproving(3)(cf.[18]).ThefactisthatBrownhimselfrefersbothtoFoxandJacksoninhispaperinquestion,andhelaysnoclaimtothe“classical”resultsonspecialcasesoftheexponentiallaw(asapersonalcommunicationwithhimreveals).Atthesametime,BrownwastheﬁrsttoshowthatthecategoryofHausdorffk-spacesiscartesianclosed,butunfortunatelythisseemstobesomewhatforgotten(cf.[5,3.3]).WhiteheadprovedthatifXisalocallycompactHausdorffspace,then1X×gisaquotientmapinHausforeveryquotientmapginHaus(cf.[35]).Michaelprovedthattheconverseisalsotrue:ForaHausdorffspaceX,themap1X×gisaquotientmapinHausforeveryquotientmapginHausifandonlyifXislocallycompact(cf.[25,2.1]).Inotherwords,thefunctorX×−preservesquotientsifandonlyifXislocallycompact.IfXisexponentiable,thenX×−mustpreserve
G.Luk´acs/Notesondualitytheoriesofabeliangroups27everycolimitinHaus([24,V.5.1]),andinparticularithastopreservecoequalizersinHaus(whichisequivalenttopreservationofquotients;fordetailssee[22,1.1]).Thus,byMichael'sresultthelocalcompactnessofXfollows.Therefore,exponentiablespacesinHauscanbecharacterizedasfollows.
B.2.Theorem.AspaceX∈HausisexponentiableinHausifandonlyifXislocallycompact.AproofofsufﬁciencyoflocalcompactnessinTheoremB.2isavailableinstandardtopologytextbooks(cf.[13,3.2]).
B.3.Corollary.LetX,Y∈Haus.Theevaluationmape:X−→C(C(X,Y),Y)(4)x7−→ˆx[ˆx(f)=f(x)](5)isk-continuous,thatis

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

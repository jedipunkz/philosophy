---
source: "https://arxiv.org/abs/2405.05546v1"
title: "Data reification in a concurrent rely-guarantee algebra"
author: "Larissa A. Meinicke, Ian J. Hayes, Cliff B. Jones"
year: "2024"
publication: "arXiv preprint / cs.LO"
download: "https://arxiv.org/pdf/2405.05546v1"
pdf: "https://arxiv.org/pdf/2405.05546v1"
captured_at: "2026-06-18T11:16:10Z"
updated_at: "2026-06-18T11:16:10Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács reification"
tags:
  - "現代思想"
  - "マルクス主義"
  - "物象化論"
  - "西洋マルクス主義"
status: raw
---

# Data reification in a concurrent rely-guarantee algebra

- 著者: Larissa A. Meinicke, Ian J. Hayes, Cliff B. Jones
- 年: 2024
- 掲載情報: arXiv preprint / cs.LO
- 情報源: [arxiv](https://arxiv.org/abs/2405.05546v1)
- ダウンロード: https://arxiv.org/pdf/2405.05546v1
- PDF: https://arxiv.org/pdf/2405.05546v1

## Obsidian Links

- 研究動向: [[研究動向/ルカーチ・ジェルジュ-現代研究動向|ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代思想]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[物象化論]]
- 関連分野: [[西洋マルクス主義]]
- 関連タグ: #現代思想 #マルクス主義 #物象化論 #西洋マルクス主義

## Abstract

Specifications of significant systems can be made short and perspicuous by using abstract data types; data reification can provide a clear, stepwise, development history of programs that use more efficient concrete representations. Data reification (or "refinement") techniques for sequential programs are well established. This paper applies these ideas to concurrency, in particular, an algebraic theory supporting rely-guarantee reasoning about concurrency. A concurrent version of the Galler-Fischer equivalence relation data structure is used as an example.

## PDF Text

Datareiﬁcationinaconcurrentrely-guaranteealgebraLarissaA.Meinicke1[0000−0002−5272−820X],IanJ.Hayes1[0000−0003−3649−392X],andCliffB.Jones2[0000−0002−0038−6623]1SchoolofElectricalEngineeringandComputerScience,TheUniversityofQueensland,Brisbane,Queensland4072,Australia2SchoolofComputingScience,NewcastleUniversity,NewcastleuponTyne,UKAbstract.Speciﬁcationsofsigniﬁcantsystemscanbemadeshortandperspicu-ousbyusingabstractdatatypes;datareiﬁcationcanprovideaclear,stepwise,developmenthistoryofprogramsthatusemoreefﬁcientconcreterepresenta-tions.Datareiﬁcation(or“reﬁnement”)techniquesforsequentialprogramsarewellestablished.Thispaperappliestheseideastoconcurrency,inparticular,analgebraictheorysupportingrely-guaranteereasoningaboutconcurrency.Acon-currentversionoftheGaller-Fischerequivalencerelationdatastructureisusedasanexample.1IntroductionDatareﬁnementtechniquesforsequentialprogramsarewellestablished[21,11,12,16].(Thispapermostlyusestheverb“reify”andnoun“reiﬁcation”inthesenseofmakingconcrete.)Thespeciﬁcationofanabstractdatastructure(ortype)includesbothinitialisationofthedatastructureandasetofoperationsonthedatastructure.Therepresentationoftheabstractdatastructuremayusespeciﬁcationdatatypes(e.g.maps)thatarenotdirectlyavailableinthetargetprogramminglanguageandthespeciﬁcationoftheoperationsmaybeintermsofpreconditionsandpostconditions,ratherthancode.Thedatastructureisencapsulatedinthesensethatitcanonlybeaccessedormodiﬁedusingtheoperationsdeﬁnedonthedatastructure.Becausethedatastructurecannotbedirectlyaccessedbytheprogramusingit,therepresentationusedinanimplementationofthedatastructuremaydifferfromthatinthespeciﬁcationinordertoprovideaconcretedatastructurethatisdeﬁnedintermsofprogramminglanguagedatatypes(e.g.ahashtable)thatprovideefﬁcientimplementationsoftheabstractoperations.Letvrepresentanabstractdatastructurewithinitialisationsatisfyingivandletcvbeaprogramusingtheabstractoperationsonv,andletwrepresentthecorrespondingimplementationdatastructurewithinitialisationsatisfyingiwandletcwbethesameascvbutwithanyoperationonvwithincvreplacedbythecorrespondingoperationonw.Bothvandwcanbetypesthatarerestrictedbydatatypeinvariants.Bothcvandcwmayuseothervariablesincommonthatarenotsubjecttoachangeofrepresentation.Theoverallcorrectnesscriteriafordatareiﬁcationisthatanyprogramrunningusingtheimplementationofthedatatypeisareﬁnementofthecorrespondingprogramusing arXiv:2405.05546v1 [cs.LO] 9 May 2024
2LarissaA.Meinicke,IanJ.Hayes,andCliffB.Jones
✓ϵrϵπϵππϵrrrpqgggσ0σ1σ2σ3σ4σ5σ6σ7Fig.1.AnAczeltracesatisfyingaspeciﬁcationwithpreconditionp,postconditionq,relycondi-tionrandguaranteeconditiong.Iftheinitialstate,σ0,isinpandallenvironmenttransitions(ϵ)satisfyr,thenallprogramtransitions(π)mustsatisfygandtheinitialstateσ0mustberelatedtotheﬁnalstateσ7bythepostconditionq,alsoarelationbetweenstates.theabstractspeciﬁcationofthedatatype,wherev:�iv′�isaspeciﬁcationcommandthatestablishesthepostconditioniv,changingonlyvariablesinv.(varv.v:�iv′�;cv)⪰(varw.w:�iw′�;cw).(1)Datareiﬁcationisatechniquethatallowsthisreﬁnementtobejustiﬁedbyshowingeachabstractoperationisreﬁnedbythecorrespondingconcreteoperation.Todothisacouplinginvariant,I,isintroducedrelatingtheconcreteandabstractrepresentations.Commonly,thecouplinginvariantisintheformofadatatypeinvariantonthecon-cretestateplusa“retrieve”functionthatextractstheabstractionrepresentationfromaconcreterepresentationthatsatisﬁesthedatatypeinvariant.Forshared-memoryconcurrentprograms,themainchangeisthatcvandhencecwcanincludeconcurrency.Theapproachwetaketohandlingreasoningaboutconcurrentprogramsisbasedontheuseofrelyandguaranteeconditions[13,14,15,8].ArelyconditionofathreadTisabinaryrelationrbetweenprogramstates;itrepresentsanassumptionthateverytransitionfromastateσitoastateσi+1madebytheenvironmentofT(i.e.allthethreadsrunninginparallelwithT)satisﬁesr,(i.e.(σi,σi+1)∈r).Com-plementingthisThasaguaranteeconditiong,alsoabinaryrelationbetweenstates,andTensuresallprogramtransitionsitmakessatisfyg,(i.e.(σi,σi+1)∈g).TheguaranteeofTmustimplytherelyconditionsofallthreadsrunninginparallelwithT.Ourse-manticsusesAczeltraces[1,3]thatexplicitlyrepresentenvironment(ϵ)transitionsaswellasprogram(π)transitions(seeFigure1).Itwasobservedin[17]thatdataabstractionandreiﬁcationcanplayanimportantroleinformallydevelopingconcurrentprograms.Thecontributionofthispaperistoextenddatareiﬁcationtoconcurrentprogramsbasedontherely-guaranteestyle.Ourapproachmakesuseof:–programinvariantsthatholdineverystateandcanbeusedtoimposeacouplinginvariantbetweentheabstractandconcretestates[20]–seeSect.3,and–adatareiﬁcationrelationbetweencommandsdeﬁnedintermsofacouplinginvari-ant–seeSect.4.Sect.5usesthesetheoriestoshowthatthecorrectnesscondition(1)isvalidandSect.6appliesourtheorytoaninterestingexample:thedevelopmentofthenon-trivialrepre-sentationofequivalencerelationsduetoGallerandFischer.OurtheoryissupportedbyanextensivelibraryofIsabelle/HOLtheories[22].Thenextsectionintroducesourwide-spectrumlanguage.
Concurrentdatareiﬁcation32Wide-spectrumlanguage2.1LanguageprimitivesOurlanguageconsistsofacompletelatticeofcommandswithpartialorder,c⪰d,meaningcommandcisreﬁned(orimplemented)bycommandd,sothatnon-deterministicchoice(c∨d)isthelatticejoinandstrongconjunction(c∧d)isthelatticemeet.ThelatticeiscompletesothatforasetofcommandsC,non-deterministicchoice�Candstrongconjunction�Caredeﬁnedastheleastupperboundandgreatestlowerbound,respectively.Thelanguageincludesbinaryoperatorsforsequentialcomposition(c;d),parallelcomposition(c∥d),andweakconjunction(c⋒d),whichbehaveslikestrongconjunction(∧)unlesseitheroperandaborts,inwhichcasetheweakconjunctionaborts.Namingandsyntacticprecedenceconventions.Weuseσforprogramstates(i.e.map-pingsfromvariablenamestovalues),canddforcommands;pandIforsetsofprogramstates(withIreservedforinvariants);g,qandrforbinaryrelationsbetweenprogramstates;eforexpressions,andkforvalues.Subscriptedversionsoftheabovenamesfol-lowthesameconvention.Unaryoperationsandfunctionapplicationhavehigherprece-dencethanbinaryoperations.Non-deterministicchoice(∨)hasthelowestprecedence,andsequentialcomposition(;)hasthehighestprecedence,exceptframing(:)ishigher.Weuseparenthesestoresolveallothersyntacticambiguities.Aprogramstateσ∈Σgivesthevaluesoftheprogramvariables.Thelanguageincludesfourprimitivecommands:τpisaninstantaneoustestthatthecurrentstateσsatisﬁesp:ifσ∈pitisano-op,otherwiseitisinfeasible;πrisanatomicprogramcommandthatmayperformaprogramtransitionfromσ0toσif(σ0,σ)∈r,otherwiseitisinfeasibleifσ̸∈domr;ϵrisanatomicenvironmentcommandthatmayperformanenvironmenttransitionfromσ0toσif(σ0,σ)∈r,otherwiseitisinfeasibleifσ̸∈domr;�isDijkstra’sabortcommand[4,5]thatirrecoverablyfailsimmediately.Itisthegreat-estcommand.Thecommandτisthetestthatalwayssucceeds(2).Theassertioncommand,{p},abortsifthecurrentstateisnotinp,where p=Σ−p,otherwiseitisano-op(3).Thecommandαrallowseitheraprogramoranenvironmenttransition,provideditsatisﬁesr(4).Theabbreviationsπandϵallowanyprogram(5)orenvironment(6)transition,respectively,andαallowsanytransition,programorenvironment(7),whereunivistheuniversalrelationbetweenprogramstates.Notetheboldfontsforτ,π,ϵandα.τ�=τΣ(2){p}�=τ∨τ
p;�(3)αr�=πr∨ϵr(4)π�=πuniv(5)ϵ�=ϵuniv(6)α�=αuniv(7)Thebasiccommandssatisfythefollowingreﬁnementproperties.τp1⪰τp2ifp1⊇p2(8){p1}⪰{p2}ifp1⊆p2(9)πr1⪰πr2ifr1⊇r2(10)ϵr1⪰ϵr2ifr1⊇r2(11)
4LarissaA.Meinicke,IanJ.Hayes,andCliffB.JonesAcommand,a,isconsideredatomicifitisoftheforma=πg∨ϵrforsomerelationsgandr.Thatis,acanonlymakeasingletransition,whichmaybeaprogram(π)transitioningoranenvironment(ϵ)inr.Atestcanbepulledoutofabranchofaweakconjunctionprovidedtheotherbranchdoesnotimmediatelyabort,(i.e.dstartswithanatomiccommandordisatest).τp;c⋒d=τp;(c⋒d)ifdisnotimmediatelyaborting(12)2.2DerivedcommandsFiniteiteration,c⋆,andpossiblyinﬁniteiteration,cω,ofacommandcaredeﬁnedastheleastandgreatestﬁxedpoints,respectively,ofthefunction(λy.τ∨c;y).Aprogramguaranteecommand,guarπrforrelationr,requireseveryprogramtransitionsatisﬁesrbutplacesnoconstraintsonenvironmenttransitions(15).Anenvironmentguaranteecommand,guarϵr,requireseveryenvironmenttransitionsatisﬁesr(16).Arelycom-mand,relyr,assumesenvironmenttransitionssatisfyr;ifonedoesnotitaborts(17),inthesamewaythatanassertion{p}abortsiftheinitialstateisnotinp.Thecommandtermonlyperformsaﬁnitenumberofprogramtransitionsbutdoesnotconstrainitsenvironment(18).Therelationidstandsfortheidentityrelationandidv,istheidentityfunctiononasetofvariablesv,and visthesetcomplementofv.Thecommandv:c,restrictscsothatitsprogramtransitionsmayonlymodifyvariablesinthesetv(19).Withintheidentityfunctionandframesavariableustandsforthesingletonset{u}andacommaisusedforunion,sothatforvariablesuandw,idu,w=id{u}∪{w}=id{u,w}.Similarly,id u=id
{u}.Thecommandidlecanperformonlyaﬁnitenumberofidling(nochange)programtransitionsbutdoesnotconstrainitsenvironment(20).Thecommandoptreitherperformsasingleprogramtransitionsatisfyingroritcanterminateim-mediatelyfromstatesσsuchthat(σ,σ)∈r(21).Theatomicspeciﬁcationcommand,⟨r⟩,achievesrviaan(atomic)programtransitionandallowsﬁnitestutteringprogramtransitions(thatdonotchangethestate)beforeandaftertheatomictransition(22);likeoptr,theatomictransitionsatisfyingrmaybeelidedforstatesσinwhich(σ,σ)∈r.c⋆�=µy.τ∨c;y(13)cω�=νy.τ∨c;y(14)guarπr�=(πr∨ϵ)ω(15)guarϵr�=(π∨ϵr)ω(16)relyr�=(α∨ϵ
r;�)ω(17)term�=α⋆;ϵω(18)v:c�=guarπid v⋒c(19)idle�=∅:term(20)optr�=πr∨τ{σ.(σ,σ)∈r}(21)⟨r⟩�=idle;optr;idle(22)Thepostconditionspeciﬁcationcommand,�q�,forarelationbetweenstatesq,guaran-teestoterminateinaﬁnalstateσthatisrelatedtotheinitialstateσ0byq(23).Thecommand[[e]]ksucceedsiftheexpressioneevaluatestothevaluekbutbecomesinfea-sibleotherwise.3Anassignmentcommand,v:=e,isnon-deterministicduetopossibleinterferencefromconcurrentthreadsmodifyingthevaluesofthevariablesusedwithine.Thenotationσ[v�→k]standsforσupdatedsothatvmapstok.Anassignmentevaluatesetosomevaluek,whichisthenatomicallyassignedtov(24);theidleatthe
3See[9]forthecompletedetailsofexpressions.
Concurrentdatareiﬁcation5endallowsforhiddenstutteringstepsintheimplementation.Aconditionaleithereval-uatesbtotrueandexecutesc,ortofalseandexecutesd(25);theidlecommandattheendallowsfor(hidden)programbranchingtransitions.Awhileloopexecutescwhilebevaluatestotrueandterminateswhenbevaluatestofalse(26).�q��=�σ0τ{σ0};term;τ{σ.(σ0,σ)∈q}(23)v:=e�=�k[[e]]k;v:opt{(σ0,σ).σ=σ0[v�→k]};idle(24)ifbthencelsed�=([[b]]true;c∨[[b]]false;d);idle(25)whilebdoc�=([[b]]true;c)ω;[[b]]false(26)2.3LocalvariableblocksLocalvariableblocksaredeﬁnedintermsofamoreprimitiveoperatorsakintoexis-tentialquantiﬁcationofavariableoveracommand.ThealgebraoftheseoperatorsisbasedonTarski’scylindricalgebra4[10],whichintroducesanexistentialoperator∃vxforeachvariablevandvariableandelementx.Anelementxisindependentofvifx=∃vx.Astandardmodeloftheiralgebraissetsofstates,andforasetofstatesp,
E
Svpisthesetofstatesσsuchthatthereexistsastateσp∈p,suchthatσandσpareequalexceptpossiblyforthevalueofv.Ourearlierworkextendedcylindricalgebratoapplytocommandsforsequentialprograms[6]andconcurrentprograms[19],fromwhichthepropertiesusedherearetaken.Wemakeuseoftwoexistentialoperatorsoncommands:5
E
Cvcthathasatracetifthereexistsatracetcofcsuchthatthetwotracesarethesameexceptforthevalueofvineachstate,
L
Cvcthathasatracetifthereexistsatracetcofcsuchthatthetwotracesarethesameexceptforthevalueofvintheﬁnalstate;iftisinﬁnite(i.e.thereisnoﬁnalstate)thentmustbeatraceofc.Thefollowingproperty,inwhichaisanatomiccommandandcisacommand,allowsanexistentialtobedistributedoverasequentialcomposition.
E
Cv(c;aω)=
E
Cvc;(
E
Cva)ωif
L
Cva=
E
Cva(27)Alocalblockintroducesalocalvariablev(28).Ifvexistsasavariableintheoutercontextofthelocalblock,itsvalueisunmodiﬁedbyprogramtransitionsofthelocalblock(i.e.guarπidv).Thevariablevisexistentiallyquantiﬁedoverthescopeofthelocalblock,andwithinthatblock,vcannotbemodiﬁedbyenvironmenttransitions(i.e.guarϵidv).Avariableblock(29)introducesalocalscopeforvbutincludesidlecommandsbeforeandaftertorepresenthiddenactionstoallocateanddeallocatev.(localv.c)�=guarπidv⋒
E
Cv(guarϵidv⋒c)(28)(varv.c)�=idle;(localv.c);idle(29)
4Thename“cylindric”isappropriatefortheirmodelbutnotreallyappropriateforourmodel.5Thereisalsoﬁrststateexistentialoperator,
F
Cvc,butwedonotneedthathere.
6LarissaA.Meinicke,IanJ.Hayes,andCliffB.JonesAcommandcisindependentofavariablevifc=
E
Cvc.Thefollowingpropertiesfollowfromthetheoryin[19].u:c=(localv.v,u:c)ifu̸=vandc=
E
Cvc(30)(varv.(localw.c))=(varw.(localv.c))(31)(localv.(v,u:c);d)=u:c;(localv.d)ifu̸=vandc=
E
Cvc(32)(localv.c⋒d)=(localv.c)⋒difd=
E
Cvd(33)(localv.c⋒(v,u:d))=(localv.c)⋒u:difu̸=vandd=
E
Cvd(34)3InvariantsGivenasetofstatesp,wedeﬁne‵pandp′tobetherelationsbetweeninitialandﬁnalprogramstatesthatsatisfypintheirinitialandﬁnalprogramstates,respectively.Apair(σ0,σ)isinrelationr1⇒r2,ifitisnotinr1oritisinr2.Itisexclusivelyusedhereintheform‵p⇒p′,thatis,thesetofpairs(σ0,σ)suchthateitherσ0/∈porσ∈p.‵p�={(σ0,σ)|σ0∈p}(35)p′�={(σ0,σ)|σ∈p}(36)r1⇒r2�=
r1∪r2(37)ThecommandinvIcantakeanarbitrarynumberofatomictransitions,eachofwhichestablishesIinitsﬁnalstate,assumingtheinitialstatesatisﬁesI.Deﬁnition1(inv).ForasetofstatesI,invI�={I};(αI′)ωAninvariantcanbeintroducedifitholdsinitially(38).AninvariantmaintainsIoneverytransition(39).{I};c⪰invI⋒c(38)invI={I};(α(‵I⇒I′))ω(39)Lemma1(inv-distribute).invI⋒�C=�c∈C(invI⋒c)ifC̸=∅(40)invI⋒(c∨d)=(invI⋒c)∨(invI⋒d)(41)invI⋒(c;d)=(invI⋒c);(invI⋒d)(42)invI⋒(c∥d)=(invI⋒c)∥(invI⋒d)(43)invI⋒(c⋒d)=(invI⋒c)⋒(invI⋒d)(44)invI⋒c⋆=(invI⋒c)⋆(45)invI⋒cω=(invI⋒c)ω(46)invI⋒τp={I};τp(47)invI⋒{p}={I∩p}(48)invI⋒πr={I};π(r∩I′)(49)invI⋒ϵr={I};ϵ(r∩I′)(50)
Concurrentdatareiﬁcation74DatareiﬁcationundercouplingacouplinginvariantWesaythatacommanddisadatareiﬁcationofacommandcundercouplinginvariantI,writtenc⪰Id,whencisreiﬁedbydundertheassumptionthatthecouplinginvariantIisinitiallytrueandIismaintainedbyeveryatomicstep.Deﬁnition2(data-reﬁnes).ForcouplinginvariantIandcommandscandd,c⪰Id�=(invI⋒c⪰invI⋒d)Fromthedeﬁnition,datareiﬁcationisapreorder,thatis,itisreﬂexiveandtransitive,andc⪰Idifc⪰d.Reifyingacompositecommandcanbereducedtoreifyingitscomponents.Lemma2(reify-operators).�C⪰I�Dif∀d∈D.∃c∈C.c⪰Id(51)c1∨c2⪰Id1∨d2ifc1⪰Id1andc2⪰Id2(52)c1;c2⪰Id1;d2ifc1⪰Id1andc2⪰Id2(53)c1∥c2⪰Id1∥d2ifc1⪰Id1andc2⪰Id2(54)c1⋒c2⪰Id1⋒d2ifc1⪰Id1andc2⪰Id2(55)c⋆⪰Id⋆ifc⪰Id(56)cω⪰Idωifc⪰Id(57)Proof.Ineachcasethedeﬁnition⪰Iisunfoldedtomakeuseoftheinvariantcommand,theinvariantisdistributedovertheoperatorusing(40–46),datareiﬁcationofthecom-ponentcommandsfollowsfromthecorrespondingassumption,thedistributionoftheinvariantisreversed,andtheresultisfoldedusingthedeﬁnitionof⪰I.
Lemma3(reify-commands).τp1⪰Iτp2ifp1⊇I∩p2(58)πr1⪰Iπr2ifr1⊇‵I∩r2∩I′(59)ϵr1⪰Iϵr2ifr1⊇‵I∩r2∩I′(60){p1}⪰I{p2}ifI∩p1⊆p2(61)guarπg1⪰Iguarπg2ifg1⊇‵I∩g2∩I′(62)relyr1⪰Irelyr2if‵I∩r1∩I′⊆r2(63)w:c⪰Iguarπg⋒v,w:cifidv⊇‵I∩g∩I′andv̸=w(64)optq1⪰Ioptq2ifq1⊇‵I∩q2∩I′(65){p};�q1�⪰I{I∩p};�q2�ifq1⊇‵I∩‵p∩q2∩I′(66)Proof.For(58),expandingDeﬁnition2(data-reﬁnes),wemustshow,invI⋒τp1⪰invI⋒τp2,orcombininginvariantandtestby(47),{I};τp1⪰{I};τp2,whichfollowsbystrengtheningthetestundertheprecondition,I,asp1⊇I∩p2.
8LarissaA.Meinicke,IanJ.Hayes,andCliffB.JonesWegivetheprooffor(59),theprooffor(60)issimilar.ExpandingDeﬁnition2(data-reﬁnes),wemustshow,invI⋒πr1⪰invI⋒πr2,orcombininginvariantandprogramtransitionby(49),{I};π(r1∩I′)⪰{I};π(r2∩I′),whichfollowsbystrength-eningtherelationunderthecontextofthepreconditionbecauser1⊇‵I∩r2∩I′.For(61),by(48),invI⋒{p1}={I∩p1}⪰{I∩p2}=invI⋒{p2},astheassumptionimpliesI∩p1⊆I∩p2.For(62),by(52)and(59)usingassumptiong1⊇‵I∩g2∩I′fortheprogramtransition,(πg1∨ϵ)⪰I(πg2∨ϵ),andhenceby(57),(πg1∨ϵ)ω⪰I(πg2∨ϵ)ω,andtheresultfollowsbythedeﬁnitionofaguarantee(15).For(63), r1⊇‵I∩
r2∩I′isequivalenttotheassumption‵I∩r1∩I′⊆r2: r1⊇‵I∩
r2∩I′⇔r1⊆
‵I∩
r2∩I′⇔r1⊆‵
I∪r2∪
I′⇔‵I∩r1∩I′⊆r2.By(52),(53)and(60)using r1⊇‵I∩
r2∩I′fortheenvironmenttransition,(π∨ϵ∨ϵ
r1;�)⪰I(π∨ϵ∨ϵ
r2;�)andhenceby(57),(π∨ϵ∨ϵ
r1;�)ω⪰I(π∨ϵ∨ϵ
r2;�)ω,andtheresultfollowsbythedeﬁnitionofarely(17).For(64),thedeﬁnitionofframing(19)isexpandedtoaguarantee,whichisthensplit,w:c=guarπid w⋒c=guarπidv⋒guarπid v,w⋒c=guarπidv⋒v,w:c⪰Iguarπg⋒v,w:cby(62)usingtheassumptionidv⊇‵I∩g∩I′.For(65),theprooffollowsfromthedeﬁnitionofoptq1(21),(52),(59)and(58).optq1=πq1∨τ{σ.(σ,σ)∈q1}⪰Iπq2∨τ{σ.(σ,σ)∈q2}=optq2Fortheapplicationof(58)oneneedstoshow{σ.(σ,σ)∈q1}⊇I∩{σ.(σ,σ)∈q2},giventheassumption,q1⊇‵I∩q2∩I′.{σ.(σ,σ)∈q1}⊇{σ.(σ,σ)∈‵I∩q2∩I′}={σ.(σ,σ)∈q2∧σ∈I}=I∩{σ.(σ,σ)∈q2}For(66),theresultfollowsfromthedeﬁnitionofaspeciﬁcationcommand(23),reﬁningusingtheassumptionq1⊇‵I∩‵p∩q2∩I′,reifyingboththeassertion(61)andthetest(58),notingtheinitialstateisassumedtobeinI∩p,andfoldingusingthedeﬁnitionofaspeciﬁcationcommand(23).{p};�q1�={p};�σ0τ{σ0};term;τ{σ.(σ0,σ)∈q1}⪰{p};�σ0τ{σ0};term;τ{σ.(σ0,σ)∈‵I∩‵p∩q2∩I′}⪰I{I∩p};�σ0τ{σ0};term;τ{σ.(σ0,σ)∈‵I∩‵p∩q2}={I∩p};�σ0τ{σ0};term;τ{σ.(σ0,σ)∈q2}={I∩p};�q2�
Lemma4(reify-rgspec).IfI∩p1⊆p2and‵I∩r1∩I′⊆r2andg1⊇‵I∩g2∩I′andq1⊇‵I∩‵p1∩q2∩I′thenrelyr1⋒guarπg1⋒{p1};�q1�⪰Irelyr2⋒guarπg2⋒{p2};�q2�.
Concurrentdatareiﬁcation9Proof.By(55),(63),(62),(66).
Thenotationeσstandsforthevalueofeinstateσ.Ifanexpressioneissingle-referenceunderarelyconditionr,evaluatingeunderinterferencesatisfyingrcorre-spondstoevaluatingeinoneofthestatesduringitsexecution,andhencethefollowingpropertyholds.6relyr⋒[[e]]k=relyr⋒idle;τ({σ.eσ=k});idle.(67)Lemma5(reify-expr).Ife1issinglereferenceunderr1,e2issinglereferenceunderr2,and‵I∩r1∩I′⊆r2and{σ.e1σ=k1}⊇I∩{σ.e2σ=k2}then,relyr1⋒[[e1]]k1⪰Irelyr2⋒[[e2]]k2.Proof.Theproofuses(67),thenreiﬁestherely(63)andthetest(58).relyr1⋒[[e1]]k1=relyr1⋒idle;τ({σ.e1σ=k1});idle⪰Irelyr2⋒idle;τ({σ.e2σ=k2});idle=relyr2⋒[[e2]]k2
Lemma6(reify-conditional).Ifb1andb2arebooleanexpressionsthataresingle-referenceunderr1andr2,respectively,and‵I∩r1∩I′⊆r2andB1={σ.(b1)σ}andB2={σ.(b2)σ},andI∩B1=I∩B2,andbothc1⪰Id1andc2⪰Id2,relyr1⋒ifb1thenc1elsec2⪰Irelyr2⋒ifb2thend1elsed2Proof.Theproofusesthedeﬁnitionofaconditional(25),thenLemma5asI∩B1=I∩B2,and(52)and(53).relyr1⋒ifb1thenc1elsec2=relyr1⋒([[b1]]true;c1∨[[b1]]false;c2);idle⪰Irelyr2⋒([[b2]]true;d1∨[[b2]]false;d2);idle=relyr2⋒ifb2thend1elsed2
Lemma7(reify-while).Ifb1andb2arebooleanexpressionsthataresingle-referenceunderr1andr2,respectively,and‵I∩r1∩I′⊆r2andB1={σ.(b1)σ}andB2={σ.(b2)σ},andI∩B1=I∩B2,andc1⪰Ic2,relyr1⋒whileb1doc1⪰Irelyr2⋒whileb2doc2Proof.TheproofissimilartothatforLemma6(reify-conditional).
5DatareﬁnementofalocalvariableblockThefollowinglemmaeliminatestheabstractvariable,v,exceptforensuringvisnotchangedbyprogramtransitions.ItisrequiredtoproveTheorem1(data-reiﬁcation),ourmaintheorem.
6See[9]forthecompletedetailsofsingle-referenceexpressionsandthedeﬁnitionofeσ.
10LarissaA.Meinicke,IanJ.Hayes,andCliffB.JonesLemma8(diminish-inv).Ifv̸=wandidw,v⊆(‵I⇒I′)then(localv.τI;invI)⋒guarϵidw⪰τ(
E
SvI);guarπ(‵(
E
SvI)⇒(
E
SvI)′)⋒guarϵidw⋒guarπidv.Proof.(localv.τI;invI)⋒guarϵidw=usinginvariantproperty(39),distributingguarϵidwby(33)asindependentofv(localv.τI;α(‵I⇒I′)ω⋒guarϵidw)=unfoldingthedeﬁnitionofalocalblock(28),andmergingenvironmentguaranteesguarπidv⋒
E
Cv(τI;α(‵I⇒I′)ω⋒guarϵidw,v)=rewritingtheatomicstepiterationguarπidv⋒
E
Cv((τI;(guarπ(‵I⇒I′)⋒guarϵ(‵I⇒I′))⋒guarϵidw,v)=distributingtest,andmergingenvironmentguaranteesguarπidv⋒
E
Cv(τI;(guarπ(‵I⇒I′)⋒guarϵ((‵I⇒I′)∩idw,v)))⪰reﬁningguaranteeandapplyingassumptionidv,w⊆(‵I⇒I′)guarπidv⋒
E
Cv(τI;(guarπ(‵(
E
SvI)⇒I′)⋒guarϵidw,v))=by(27)as
L
Cv(π(‵(
E
SvI)⇒I′)∨ϵidw,v)=π(‵(
E
SvI)⇒(
E
SvI)′)∨ϵidwguarπidv⋒
E
Cv(τI);(guarπ(‵(
E
SvI)⇒(
E
SvI)′)⋒guarϵidw)=distributingtestovernon-abortingcommands,andcommuting⋒
E
Cv(τI);guarπ(‵(
E
SvI)⇒(
E
SvI)′)⋒guarϵidw⋒guarπidvWhere,fortheapplicationofproperty(27),weusethefactthataweakconjunctionofprogramandenvironmentguaranteescanberewrittenastheiterationofanatomicstep:guarπg⋒guarϵr=(πg∨ϵr)ωforanygandr.
OurmaintheoremshowsthataprogramusingtheabstractstateisreﬁnedbythesameprogramusingtheimplementationstateviathecouplinginvariantI.Theorem1(data-reiﬁcation).Givendisjointsetsofvariablesu,vandw,andcou-plinginvariantIrelatingvandw,setsof(initial)statesivandiw,commandscv,cwanddw,ifivandcvareindependentofw,andiw,cwanddwareindependentofv,andifiw∩I⊆iv(68){iw};v,w,u:cv⋒guarϵidv,w⪰Iv,w,u:cw⋒guarϵidv,w(69)idw,v⊆(‵I⇒I′)(70)iw⊆
E
SvI(71)guarπ(‵(
E
SvI)⇒(
E
SvI)′)⋒{iw};w,u:cw⪰w,u:dw(72)holdthen(varv.v:�iv′�;u,v:cv)⪰(varw.w:�iw′�;u,w:dw).
Concurrentdatareiﬁcation11GivenacouplinginvariantI,
E
SvIisthedata-typeinvariantontheimplementationvari-ablesw.Assumption(68)ensurestheconcreteinitialisationcorrespondtoanabstractinitialisation;(69)correspondstoreifyingthebodyoftheblock;(70)ensuresthecou-plinginvariantIismaintainedifneithervnorwismodiﬁed;(71)ensurestheinitialstatesatisﬁesthedatatypeinvariant;and(72)eliminatesthedatatypeinvariant.Proof.(varv.v:�iv′�;v,u:cv)=introducelocalvariablewby(30)w̸=v;ivandcvareindependentofw(varv.(localw.v,w:�iv′�;v,w,u:cv))=interchangevariablesvandwby(31)(varw.(localv.v,w:�iv′�;v,w,u:cv))⪰reﬁneinitializationusing(68)andv,w:�iw′�;τI=v,w:�iw′�;τI;{iw}(varw.(localv.v,w:�iw′�;τI;τiw;v,w,u:cv))=restrictthescopeofvby(32)asv̸=wandiwisindependentofv(varw.w:�iw′�;(localv.τI;τiw;v,w,u:cv))⪰introducecouplinginvariantIby(38)asτI=τI;{I}(varw.w:�iw′�;(localv.τI;(invI⋒τiw;v,w,u:cv)))⪰datareﬁneusingassumption(69)(varw.w:�iw′�;(localv.τI;(invI⋒τiw;v,w,u:cw)))=by(12)assumingτiw;v,w,u:cwisnotimmediatelyaborting(varw.w:�iw′�;(localv.τI;invI⋒τiw;v,w,u:cw))=by(34)asu,v,w:cwisindependentofvandw:�iw′�establishesτiw(varw.w:�iw′�;((localv.τI;invI)⋒w,u:cw))=implicitguarϵidwfromlocalblockforwintolocalblockforv(varw.w:�iw′�;((localv.τI;invI)⋒guarϵidw⋒w,u:cw))=byLemma8(diminish-inv)using(70)andv̸=w(varw.w:�iw′�;(τ(
E
SvI);guarπ(‵(
E
SvI)⇒(
E
SvI)′)⋒guarϵidw⋒guarπidv⋒w,u:cw))=asguarπidv⋒w,u:cw=w,u:cwand(71)(varw.w:�iw′�;(guarπ(‵(
E
SvI)⇒(
E
SvI)′)⋒guarϵidw⋒w,u:cw))=asguarϵidwisimplicitforlocalvariablewand(72)(varw.w:�iw′�;w,u:dw)
6ExampleWeconsideraconcurrentimplementationoftheGaller-Fischerdatastructureforimple-mentinganequivalencerelation[7].Ourtheoryisdevelopedintermsofsetsofstates
12LarissaA.Meinicke,IanJ.Hayes,andCliffB.Jonesandrelationsbetweenstatesbutwhenpresentinganexampleitispreferabletousepred-icatescharacterisingboththesetsofstatesandrelations.Weusethenotation⌜pred⌝todenotetherelationcharacterisedbythepredicatepred.Speciﬁcation.Theabstractstateforthespeciﬁcationisanequivalencerelationeq⊆X×X,forsomeﬁnitesetX,witheqinitiallytheidentityrelation.Alloperationsrelythattheirenvironmentonlygrowstheequivalencerelation(viaequates),i.e.eq0⊆eq,wherea0-subscriptedvariablereferstoitsinitialvalueandandunsubscriptedvariabletoitsﬁnalvalue.Thespeciﬁcationoftheoperationtestisnon-deterministictoallowforthefactthatconcurrentequatesmaytakeplacewhilethetestisexecuting:test(x,y)must returntrueifxandyareequivalentintheinitialrelationeq0,andmay returntrueifxandyarerelationintheﬁnalequivalencerelationeq.Theoperationequate(x,y)ensuresthatxandyareintheequivalencerelationaftertheoperation.Inordertoallowforconcurrenttestandequates,theequatingofxandymustoccuratomically,andhenceanatomicspeciﬁcationcommandisused.Thereﬂexive,transitiveclosureofarelationriswrittenr⋆.Transitiveclosureisusedintheatomicspeciﬁcationtoensurethatforanyuandz,if(u,x)and(y,z)areineq0then(u,z)isineq.t←test(x:X,y:X)�=rely⌜eq0⊆eq⌝⋒t:�⌜((x,y)∈eq0⇒t)∧(t⇒(x,y)∈eq)⌝�equate(x:X,y:X)�=rely⌜eq0⊆eq⌝⋒eq:⟨⌜eq=(eq0∪{(x,y),(y,x)})⋆⌝⟩clean up(x:X)�=rely⌜eq0⊆eq⌝⋒∅:termColletteandJones[2]introducedtheconceptofanevolutioninvariantr,whichcor-respondstorelyr⋒guarπr.Forourexample,eq0⊇eqisanevolutioninvariant.Ourconcurrentrely-guaranteealgebrasupportsevolutioninvariants[20].Implementation.TheGaller-Fischerimplementationdatastructureisaforestoftreesrepresentedasaparentarrayf∈X→X,withdatatypeinvariantthattheonlycyclesareoftheformf[n]=n.Initiallyfistheidentityfunction.Wedeﬁnetherelationthatnisadescendantofmby,n≤fm�=(n,m)∈f⋆.(73)Thedatatypeinvariantonfcorrespondstorequiringtherelation(≤f)tobeapartialorder,thatis,itis(a)reﬂexive,(b)transitive,and(c)anti-symmetric.Both(a)and(b)followdirectlyfromthedeﬁnition(becauseitusesreﬂexive,transitiveclosure).Property(c)isanadditionalconstraint,n≤fm∧m≤fn⇒n=mwhichavoidsanynon-trivialcyclesintheforest,i.e.theonlycyclesareoftheformfn=nfortherootelements.Wemakeuseofthefollowingauxiliarydeﬁnitionsontheforestrepresentation:n≡fmstatesthatnandmareequivalentwithintheforestf,thatis,theyhaveacommonancestorandhencearebothwithinthesametree;rootsfgives
Concurrentdatareiﬁcation13thesetofallelementsinXthataretherootofatree;androotfngivestheuniquerootofthetreewithinfthatnisanelementof.n≡fm�=∃k.n≤fk∧m≤fk(74)rootsf�={n∈X.n=fn}(75)rootfn�=ifn=fnthennelserootf(fn)(76)Notethatn≤fmisapartialorder,son≡fmisnotthesameasn≤fm∧m≤fn,asthelaterimpliesn=mwhiletheformeronlysaysnandmarebothinthesametree.Fromthesedeﬁnitionswecandeducethefollowingproperties.(rootfn=rootfm)⇔(n≡fm)(77)(n≤frootfm)⇒(n≡fm)(78)Thefunction,retr,retrievestheequivalencerelationfromtheforest.retrf�={(n,m)∈X×X.n≡fm}(79)Ourcouplinginvariantcombinestheretrievefunctionwiththedatatypeinvariant.I�=⌜eq=retrf∧partial order(≤f)⌝(80)Forthevariantfunctions,distancefxyisthenumberofarcsonthepathfromxtoywithinf,assumingx≤fy.Thevariantfunctionsareintheformofalexicographicallyorderedpair,withtheﬁrstcomponent#(rootsf)handlingthecasewhenthereisin-terferencethatdecreasesthenumberofroots(viaequates),andthesecondcomponentthenhandlesthecasewherethenumberofrootsdoesnotchange.rx:X←root of(x:X)rx:=x;whilerx̸=f[rx]dovariant(#(rootsf),distancefrx(rootfrx))rx:=f[rx]t:B←test(x:X,y:X)varrx,ry:X;(rx:=x∥ry:=y);whilerx̸=f[rx]∨ry̸=f[ry]dovariant(#(rootsf),distancefrx(rootfrx)+distancefry(rootfry))(rx:=root of(rx)∥ry:=root of(ry));t:=rx=ryequate(x:X,y:X)varrx,ry:X;done:B;(rx:=x∥ry:=y∥done:=false);while¬donedovariant(#(rootsf),distancefrx(rootfrx))(rx:=root of(rx)∥ry:=root of(ry));done:=CAS(f[rx],rx,ry)—orthesymmetricversionTheimplementationalsointroducesanoperationclean up(x)thatreducesthelengthofthepathfromxtoitscorrespondingroot,withoutchangingtheequivalencerelation.clean up(x:X)�=rely⌜eq0⊆eq⌝⋒∅:term(81)Whilemultipletestandequateoperationsmayrunconcurrently,weassumethatthereisonlyoneclean uprunningatanyonetime.
14LarissaA.Meinicke,IanJ.Hayes,andCliffB.JonesReifyingclean up.Theﬁrststepinreifyingtheclean upoperation(81)istoaddtheimplementationstate,f,toitsframe,soitbecomes,rely⌜eq0⊆eq⌝⋒f:term.(82)Wethenreifythiswithcouplinginvariant(80)aboveusing(63)and(64)togive,rely⌜retrf0⊆retrf⌝⋒guarπ⌜retrf0=retrf⌝⋒eq,f:term.(83)Theproofobligationsarestraightforward.Therelyconditionallows,forexample,mod-ifyingf0forwhichf0n=m∧f0m=mtofwithfn=n∧fm=nbutsuchamodi-ﬁcationinvalidatestheassertionn̸∈rootsf(aswellasothers).Fortheimplementationofclean uponeneedsastrongerrelycondition(whichimpliesretrf0⊆retrf),rely⌜rootsf0⊇rootsf∧(∀nm.n≤f0m⇒n≤fm)⌝.(84)Astandardruleinrely-guaranteetheoryisthatarelyconditioncanbeweakenedbuthereweneedtostrengthentherelycondition.Onecanonlystrengthenthisrelycondi-tionifallconcurrentoperationsonfguaranteethiscondition.Here,thisstrongerrelycondition(84)isguaranteedbothbyourimplementationoftest(becauseitdoesnotmodifyf)andequate(becauseitonlyeveraddstof).Note,however,ourimplemen-tationofclean updoesnotguarantee(84),whichiswhyonlyasingleclean upatanyonetimeispermitted.Thefollowinggivesanannotatedversionoftheclean upprocedure.Theassertionjustbeforethewhileloopisitsinvariant.Variablesx,fxandrxarelocaltoclean upandhencenotsubjecttomodiﬁcationbyotherthreads.Theauxiliaryvariablef0standsfortheinitialvalueoff.clean up(x:X)varfx,rx:X;rx:=root of(x);{rootf0x≤f0rx≤frootfx∧x≤frx}whilex̸=rxdovariant(#(rootsf),distancefxrx){rootf0x≤f0rx≤frootfx∧x<frx}fx:=f[x];{rootf0x≤f0rx≤frootfx∧x<ffx≤frx}f[x]:=rx;{rootf0x≤f0rx≤frootfx∧fx≤frx}x:=fx{rootf0x≤f0rx≤frootfx∧x≤frx}Onecanshowtheassertionsintheaboveprogramarestableunderthestrengthenedrely.Theguaranteeofclean up,thatis,retrf0=retrf,istriviallysatisﬁedbyallcommandsotherthantheassignmenttof[x]becausetheydonotmodifyf.Theassignmenttof[x]mustsatisfytheguarantee,thatis,retrf=retr(f[x�→rx]),wheref[x�→rx]isthesameasfexceptitmapsxtorx(insteadoffx).Expandingthedeﬁnitionofretr(79),thiscorrespondstoshowingusing(77),∀nm.(n≡fm)⇔(n≡f[x�→rx]m)≡∀nm.(rootfn=rootfm)⇔(root(f[x�→rx])n=root(f[x�→rx])m)
Concurrentdatareiﬁcation15Ifthepathsnandmtotheirrootsdonotcontainx,theequivalenceholdstrivially.Ifthepathfromntoitsrootcontainsx,wehaverootfn=rootfxandroot(f[x�→rx])n=rootfrx=rootfxfromthepreconditionrx≤frootfxusing(78),andsimilarlyifthepathfrommtoitsrootcontainsx,andhencetheequivalenceismaintained.7ConclusionDatareiﬁcationisapowerfultechniqueforseparatingtheconcernsofspecifyingadatastructure(ortype)intermsofanabstractrepresentationfromthedetailsoftheconcreterepresentationusedtoimplementthatdatastructureefﬁciently.Therely-guaranteeap-proachmakesthetaskofreasoningaboutconcurrentprogramsmoretractablebypro-vidingacompositionalapproachthatallowsonetoreasonaboutacomponentofaprograminrelativeisolation.Theapproachwehavetakenistoextendaconcurrentrely-guaranteealgebratohandledatareiﬁcation.Ourearlierresearchongeneralisedinvariants[20],andinpar-ticulartheirdistributionproperties,providesthebasisfordatareifyingcommandswithrespecttoacouplinginvariant.Ourresearchonlocalisation[19]providesafoundationfortheproofofTheorem1(data-reiﬁcation)thatjustiﬁesthatareiﬁcationofa(hidden)datastructureguaranteesthatanyprogramusingtheimplementationdatastructureisareﬁnementofthesameprogramusingtheabstractspeciﬁcationofthedatastructure.TheapproachofLiang,FengandFu[18]isalsobasedontherely-guaranteeap-proach.Itusesasimulationtechnique,RGSim,thatfocusesoncomparingexternallyobservablebehavioursonly.References1.Aczel,P.H.G.:Onaninferenceruleforparallelcomposition(1983),privatecom-municationtoCliffJoneshttp://homepages.cs.ncl.ac.uk/cliff.jones/publications/MSs/PHGA-traces.pdf2.Collette,P.,Jones,C.B.:Enhancingthetractabilityofrely/guaranteespeciﬁcationsinthedevelopmentofinterferingoperations.In:Plotkin,G.,Stirling,C.,Tofte,M.(eds.)Proof,LanguageandInteraction,chap.10,pp.277–307.MITPress(2000)3.Colvin,R.J.,Hayes,I.J.,Meinicke,L.A.:Designingasemanticmodelforawide-spectrumlanguagewithconcurrency.FormalAspectsofComputing29,853–875(2016).https://doi.org/10.1007/s00165-017-0416-44.Dijkstra,E.W.:Guardedcommands,nondeterminacy,andaformalderivationofprograms.CACM18,453–458(1975)5.Dijkstra,E.W.:ADisciplineofProgramming.Prentice-Hall(1976)6.Dongol,B.,Hayes,I.J.,Meinicke,L.A.,Struth,G.:CylindricKleenelatticesforprogramconstruction.In:Hutton,G.(ed.)MathematicsofProgramConstruction2019.LectureNotesinComputerScience,vol.11825.SpringerInternationalPublishing,Cham(Oct2019).https://doi.org/10.1007/978-3-030-33636-3_87.Galler,B.A.,Fischer,M.J.:Animprovedequivalencealgorithm.Commun.ACM7(5),301–303(May1964).https://doi.org/10.1145/364099.364331,https://doi.org/10.1145/364099.364331
16LarissaA.Meinicke,IanJ.Hayes,andCliffB.Jones8.Hayes,I.J.,Jones,C.B.,Meinicke,L.A.:Specifyingandreasoningaboutshared-variableconcurrency.In:Bowen,J.P.,Li,Q.,Xu,Q.(eds.)TheoriesofProgrammingandFormalMethods.pp.110–135.No.14080inLNCS,Springer(2023).https://doi.org/10.1007/978-3-031-40436-8_59.Hayes,I.J.,Meinicke,L.A.,Meiring,P.A.:Derivinglawsfordevelopingconcurrentprogramsinarely-guaranteestyle(2021).https://doi.org/10.48550/ARXIV.2103.15292,https://arxiv.org/abs/2103.1529210.Henkin,L.,Monk,J.D.,Tarski,A.:CylindricAlgebras,PartI,Studiesinlogicandthefoun-dationsofmathematics,vol.64.North-HollandPub.Co.(1971)11.Hoare,C.A.R.:Proofofcorrectnessofdatarepresentations.ActaInformatica1,271–281(1972),alsoinProgrammingMethodology,D.Gries(ed)Springer-Verlag(1978)12.Jones,C.B.:SoftwareDevelopment:ARigorousApproach.PrenticeHallInternational(1980),http://portal.acm.org/citation.cfm?id=53977113.Jones,C.B.:DevelopmentMethodsforComputerProgramsincludingaNotionofInterfer-ence.Ph.D.thesis,OxfordUniversity(June1981),availableas:OxfordUniversityComput-ingLaboratory(nowComputerScience)TechnicalMonographPRG-2514.Jones,C.B.:Speciﬁcationanddesignof(parallel)programs.In:ProceedingsofIFIP’83.pp.321–332.North-Holland(1983)15.Jones,C.B.:Tentativestepstowardadevelopmentmethodforinterferingprograms.ACMToPLaS5(4),596–619(1983).https://doi.org/10.1145/69575.6957716.Jones,C.B.:SystematicSoftwareDevelopmentusingVDM.PrenticeHallInterna-tional,secondedn.(1990),http://homepages.cs.ncl.ac.uk/cliff.jones/ftp-stuff/Jones1990.pdf17.Jones,C.B.:Splittingatomssafely.TheoreticalComputerScience375(1–3),109–119(2007)18.Liang,H.,Feng,X.,Fu,M.:Rely-guarantee-basedsimulationforcompositionalveriﬁcationofconcurrentprogramtransformations.ACMTransactionsonProgrammingLanguagesandSystems36(1),3:1–3:55(2014)19.Meinicke,L.A.,Hayes,I.J.:Usingcylindricalgebratosupportlocalvariablesinrely/guaranteeconcurrency.In:2023IEEE/ACM11thInternationalConferenceonFor-malMethodsinSoftwareEngineering(FormaliSE).pp.108–119.IEEE(2023).https://doi.org/10.1109/FormaliSE58978.2023.0001920.Meinicke,L.A.,Hayes,I.J.:Reasoningaboutdistributivelawsinaconcurrentreﬁnementalgebra(2024),arXiv:2403.13425[cs.LO]21.Milner,R.:Analgebraicdeﬁnitionofsimulationbetweenprograms.Tech.Rep.CS-205,ComputerScienceDept,StanfordUniversity(February1971)22.Nipkow,T.,Paulson,L.C.,Wenzel,M.:Isabelle/HOL:AProofAssistantforHigher-OrderLogic,LNCS,vol.2283.Springer(2002)

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

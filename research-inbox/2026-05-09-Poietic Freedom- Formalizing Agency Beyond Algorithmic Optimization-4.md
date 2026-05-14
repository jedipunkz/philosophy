---
source: "https://philarchive.org/rec/FERPFF-3"
title: "Poietic Freedom: Formalizing Agency Beyond Algorithmic Optimization"
author: "Jose Fernández Tamames"
year: "manuscript"
publication: ""
download: "https://philarchive.org/archive/FERPFF-3"
pdf: "https://philpapers.org/archive/FERPFF-3.pdf"
captured_at: "2026-05-09T12:47:17Z"
updated_at: "2026-05-09T12:47:17Z"
capture_tool: "scrapem"
source_name: "philarchive"
keyword: "ハンナ・アーレント"
query: "Hannah Arendt The Human Condition"
tags:
  - "政治哲学"
  - "現代思想"
  - "全体主義"
status: raw
---

# Poietic Freedom: Formalizing Agency Beyond Algorithmic Optimization

- 著者: Jose Fernández Tamames
- 年: manuscript
- 情報源: [philarchive](https://philarchive.org/rec/FERPFF-3)
- ダウンロード: https://philarchive.org/archive/FERPFF-3
- PDF: https://philpapers.org/archive/FERPFF-3.pdf

## Obsidian Links

- 研究動向: [[研究動向/ハンナ・アーレント-現代研究動向|ハンナ・アーレント-現代研究動向]]
- キーワード: [[ハンナ・アーレント]]
- 関連分野: [[政治哲学]]
- 関連分野: [[現代思想]]
- 関連分野: [[全体主義]]
- 関連タグ: #政治哲学 #現代思想 #全体主義

## Abstract

This article reconceptualizes human freedom as poietic freedom: the capacity to initiate realities not deducible from prior rules or conclusive reasons. Against dominant elective models—which reduce agency to selecting among predefined options (Ot)—we argue that paradigm cases of freedom (e.g., founding legal institutions, inventing techniques, creating art) involve ontological expansion rather than optimization. Formalized under the Free∗ operator, poietic freedom satisfies five conditions: (i) relative ontological novelty (NO), (ii) non-necessity (NN), (iii) absence of conclusive reasons (¬RC), (iv) intentional form (IF), and (v) guidance control (Ctrl). We develop a replicable Diagnostic Protocol (FDP) and Poietic Index (L∗) to identify and quantify such acts, validating the model through historical cases. After rebutting determinist and computationalist objections, we explore implications for ethics (post-factum responsibility) and technology, arguing that Generative AI, bound by statistical optimization, structurally lacks poietic capacity.

## Citation

Fernández Tamames, Jose, Poietic Freedom: Formalizing Agency Beyond Algorithmic Optimization.

## BibTeX

```bibtex
@unpublished{FernandezTamamesManuscript-FERPFF-3,
	author = {Jose Fern\
```

## Source Details

Archival history: Archival date: 2026-01-07 View all versions
Keywords: Poietic Freedom   Generative AI   Ontological Novelty   Algorithmic Optimization   Machine Agenc   Machine Agency   Large Language Models (LLMs)   First Cause   Philosophy of Creativity   Hannah Arendt   Non-Necessity

## PDF Text

PoieticFreedom:FormalizingAgencyBeyondAlgorithmicOptimizationNovelty,Non-Necessity,andSelf-DeterminationWithoutPriorRuleJoseFernandezTamamesDepartmentofComputerScienceUNIEUNIVERSITYMadrid,Spainjose.fernandezt@unieuniversidad.comJanuary2026(DraftVersion2.0)AbstractThisarticlereconceptualizeshumanfreedomaspoieticfreedom:thecapacitytoinitiaterealitiesnotdeduciblefrompriorrulesorconclusivereasons.Againstdom-inantelectivemodels—whichreduceagencytoselectingamongpredeﬁnedoptions(Ot)—wearguethatparadigmcasesoffreedom(e.g.,foundinglegalinstitutions,inventingtechniques,creatingart)involveontologicalexpansionratherthanopti-mization.FormalizedundertheFree∗operator,poieticfreedomsatisﬁesﬁvecon-ditions:(i)relativeontologicalnovelty(NO),(ii)non-necessity(NN),(iii)absenceofconclusivereasons(¬RC),(iv)intentionalform(IF),and(v)guidancecontrol(Ctrl).WedevelopareplicableDiagnosticProtocol(FDP)andPoieticIndex(L∗)toidentifyandquantifysuchacts,validatingthemodelthroughhistoricalcases.Af-terrebuttingdeterministandcomputationalistobjections,weexploreimplicationsforethics(post-factumresponsibility)andtechnology,arguingthatGenerativeAI,boundbystatisticaloptimization,structurallylackspoieticcapacity.Keywords:PoieticFreedom,OntologyofAction,ArtiﬁcialIntelligence,Creativity,Agency,Determinism,LogicalFormalization.1Introduction:TheCrisisoftheElectiveModelContemporarytheoriesofagency,particularlyinthecontextofArtiﬁcialIntelligenceanddecisiontheory,predominantlyrelyonan“elective”model.Inthisframework,freedomis1
deﬁnedasthecapacitytoselecttheoptimalpathamongasetofpre-existingalternatives(Ot).Whetherinrationalchoicetheory,gametheory,orreinforcementlearning(RL),theagentisessentiallyanoptimizer:anentitythatnavigatesaclosedontologicalspacetomaximizeautilityfunction.However,thismodelfailstoaccountforthemostdistinctivelyhumanactsoffreedom:thosethatdonotselectfromtheworldasitis,butaddtoit.WhentheHabeasCorpusactwasdraftedin1679,itwasnotaselectionamongexistinglegaloptions;itwastheinstitutionofanewontologicaltypeofprotection.Whenapotterinventsanewglaze,theyarenottraversingasearchspaceofknownchemistry,butinstitutinganewmaterialreality.WetermthiscapacityPoieticFreedom(Free∗).Unlikeelectiveagency(whichAIsystemsnowmaster),poieticfreedomisthecapacitytoinitiatecausalchainsthatareunder-determinedbythepastandover-determinedbytheintenttocreatenovelty.Thispaperaimstoformalizethisdistinction,movingthedebatefromvaguemetaphysicstoarigorous,operationaldeﬁnition.2TheOntologyofPoieticFreedomWedeﬁnehumanagencynotmerelyas“choice,”butas“ontologicalexpansion.”Toformalizethis,weintroducethePoieticFreedomOperator(Free∗).2.1FormalDeﬁnitionDeﬁnition1(ThePoieticFreedomOperator):AnactαperformedbyagentaattimetconstitutesPoieticFreedom(Free∗)ifandonlyif:Free∗(α)⇐⇒NewType(x,t)∧¬□Prior(x)∧¬RC(α)∧IF(α)∧Ctrl(a,α)(1)Where:•NewType(x,t):TheactintroducesanewontologicaltypenotpresentinthesetOt.•¬□Prior(x):Theactisnotnecessitatedbybiological,functional,ormoralan-tecedents(Non-Necessity).•¬RC(α):TheactisnotdeterminedbyaConclusiveReason(itisnottheresultofanoptimizationcalculation).•IF(α):TheactpossessesIntentionalForm(itisnotrandomnoise).2
•Ctrl(a,α):Theagentretainsguidancecontrolovertheexecution.Thisdeﬁnitionallowsustodistinguishstructurallybetweenhumancreativeagencyandalgorithmicoutput.WhileanAIgeneratesvariationsbasedonprobabilisticweight-ings(SecondCause),apoieticagentintroducesaruptureinthecausalchain(FirstCause).OtPoieticAct(α)−−−−−−−−→Ot+=Ot∪{x}(2)3PhilosophicalPrecedentsTheformaldeﬁnitionofFree∗asanactofontologicalexpansionﬁndsrobustsupportinthehistoryofphilosophy.Wehighlightthreedistinctlineagesthatpreﬁgureouroperators:1.Kant:ReﬂectiveJudgment.IntheCritiqueofJudgment(§§46–49),Kantiden-tiﬁes“genius”notasaskillforoptimization,butasthefacultythat“givestheruletoart”(giebtdieRegelderKunst).Thisoperationreliesonreﬂectivejudgment(reﬂektierendeUrteilskraft),whichdoesnotsubsumeparticularsundergivenconcepts,butproducesnewnormativeprincipleswherenoneexisted.“Thebeautifulisthatwhichpleaseswithoutaconcept.”(Kant,1790/2000,§5).RelevanceforFree∗:Kantiancreationexplicitlysatisﬁesourconditionof¬RC(Non-ConclusiveReasons),asitescapespriorconceptualdetermination.2.Arendt:NatalityasEruption.InTheHumanCondition,HannahArendtanchorsfreedomintheconceptofnatality:thehumancapacityto“beginsomethingnew”thatdisruptsforeseeablecausalchains.“Thenewalwayshappensagainsttheoverwhelmingoddsofstatisticallaws...Everybirthisthepossibilityofanontologicalmiracle.”(Arendt,1958,p.178).ConnectiontoFormalization:NatalityisthepoliticalembodimentofOntologicalNovelty(NO),rejectinghistoricaldeterminism.3.Goodman:Worldmaking.NelsonGoodman’sWaysofWorldmaking(1978)demonstratesthatsymbolicsystems(art,science,law)areproductsof“worldmaking”throughoperationslikeweightinganddeformation.Thissupportsourformulaofexpan-sion(Eq.2),showinghowFree∗actsgenerateanexpandedontologyOt+irreducibletotheclosureofOt.3
4FormalComparison:Humanvs.MachineAgencyThedistinctionbetweenhumanpoiesisandmachinegenerationisoftenblurredbythequalityofAIoutputs.However,thediﬀerenceisnotintheproduct,butintheprocessanditslogicalconditions.Table1summarizesthisontologicalgap.Table1:FormalComparison:PoieticAgencyvs.AlgorithmicOptimization
DimensionPoieticAgency(Human/FirstCause)AlgorithmicAgency(AI/SecondCause)
FormalOpera-torFree*(¬RC)ArgMax(�P)OperationalLogicRupture(Institutinganewrule)Optimization(Selectingbestpath)RelationtoRulesLegislative(Createstherule)Executive(Followsthevector)OriginofAc-tionIndigence(Responsetolack)Instruction(Executionofcode)TemporalModeInterval(Hesitation/Angst)Latency(Processingtime)FailureModeTragedy(Ontologicalerror)Hallucination(Statisticalnoise)
ThecrucialdiﬀerentiatoristheabsenceofConclusiveReasons(¬RC).AnAIalwaysactsbecausetheprobabilisticweightoftokenAishigherthantokenB.Thehumanagent,inthepoieticmoment,actsdespitethelackofsuﬃcientreason.5ThePoieticIndex(L∗)andValidationTomovebeyondtheory,weproposeametrictoquantifythedegreeoffreedominanygivenact:thePoieticIndex(L∗).Thisindexaggregatestheﬁvevariablesofourdeﬁnition.5.1TheIndexFormulaL∗(α)=5�i=1(wi·Vi)(3)WhereVi∈{0,1,2}representstheintensityof:Novelty,Non-Necessity,Non-Conclusiveness,Intentionality,andControl.AsshowninTable2,AlphaGo’sfamous“Move37”,whilenovelinthecontextofthegame(NO=1),wasstrictlynecessarygiventhemaximizationfunction(NN=0)anddeterminedbycalculation(¬RC=0).4
Table2:ValidationofHistoricalCasesusingL∗
CaseStudyNONN¬RCIFCtrlTotal(L∗)
HabeasCorpusAct(1679)2222210(Max)ChinesePorcelain(Song)221229(High)AlphaGoMove37100225(Med)ThermostatSwitch000000(Null)
5.2TheFreedomDiagnosticProtocol(FDP)TodistinguishaFree∗actfrommererandomnessorcalculation,weapplythefollowing3-steptest:1.TheNecessityTest:Couldtheacthavebeenderivedlogicallyfromthean-tecedents?(IfYES→Deterministic).2.TheIntentionTest:Doestheactexhibitacoherentformorgoal?(IfNO→RandomNoise).3.TheOptimizationTest:Wastheacttheresultofmaximizingapre-deﬁnedutilityfunction?(IfYES→AlgorithmicAgency;IfNO→PoieticFreedom).6DiscussionandRebuttalsWeaddresstwoprimaryobjectionstotheFree∗model:6.1TheDeterministObjectionObjection:“Everyeventhasacause;therefore,’novelty’isjustignoranceofhiddenvari-ables.”Rebuttal:Weacceptphysicalcausalitybutrejectnomologicalclosure.Poieticactsarephysicallycaused(bythebrain)butlogicallyunder-determined.Theagentmustsupplythe“missingrule”tobridgethegapbetweenstimulusandresponse.This“ﬁllingofthegap”iswhatwecallfreedom.6.2TheComputationalistObjectionObjection:“GenerativeAIcreatesnewthings(images,code);therefore,machineshavepoieticfreedom.”Rebuttal:AIoutputsareprobabilisticinterpolationswithinthelatentspaceoftrainingdata.TheysatisfytheconditionofcombinatorialnoveltybutfailtheconditionofNon-Necessity(¬□).TheAIcannotrefusetooptimize;ithasno“Interval”ofhesitation.5
Its“creativity”isanexecutionofinstructions(SecondCause),notaninstitutionofends(FirstCause).7ImplicationsEthics:Post-FactumResponsibility.Iffreedomisaninstitutionofnewrules,thenresponsibilityisnotjustadherencetoexistinglaws,butaccountabilityforthenewrealitieswebirth.Inthe“Interval”ofdecision,theagentassumestheriskoftheunknown.Anthropology:HomoCreator.Thehumanisnotarationalmaximizer(HomoEco-nomicus)butacreatureofindigencethatmustcreateitsownworldtosurvive(HomoCreator).8ConclusionWehavepresentedaformalmodeloffreedomthatrescuesagencyfrombothreductivedeterminismandmysticallibertarianism.Bydeﬁningfreedomaspoieticrupture(Free∗),weprovideacriteriontodistinguishthehumanactofcreationfromthealgorithmicactofgeneration.Inaneraofautomatedcognition,recognizingthespeciﬁcnatureofhumanfree-dom—itsgroundinginnon-necessityandrisk—isnotjustaphilosophicalexercise,butapoliticalimperative.Wearenotfreebecausewecalculatebetter;wearefreebecausewemustdecidewhencalculationends.DataandCodeAvailabilityThelogicalmodels,historicaldatasets,anddiagnosticprotocolsusedinthisresearchareavailableatthefollowingGitHubrepository:https://github.com/jftmames/anima-10-poietic-freedomForreproducibilitypurposes,apermanentarchivalversionofthesoftwarehasbeende-positedatZenodo:DOI:10.5281/zenodo.18173584.References[1]Arendt,H.(1958).TheHumanCondition.UniversityofChicagoPress.[2]Bender,E.M.,Gebru,T.,McMillan-Major,A.,&Shmitchell,S.(2021).Onthedangersofstochasticparrots:Canlanguagemodelsbetoobig?Proceedingsof6
the2021ACMConferenceonFairness,Accountability,andTransparency,610–623.https://doi.org/10.1145/3442188.3445922[3]FernándezTamames,J.(2026).TheFirstCause:APhilosophicalSymphonyinFiveMovements[Computersoftware].Zenodo.https://doi.org/10.5281/zenodo.18173584[4]Floridi,L.(2023).TheEthicsofArtiﬁcialIntelligence:Principles,Challenges,andOpportunities.OxfordUniversityPress.[5]Goodman,N.(1978).WaysofWorldmaking.HackettPublishing.[6]Heidegger,M.(1962).BeingandTime(J.Macquarrie&E.Robinson,Trans.).Harper&Row.(Originalworkpublished1927).[7]Kant,I.(2000).CritiqueofthePowerofJudgment(P.Guyer,Trans.).CambridgeUniversityPress.(Originalworkpublished1790).[8]Kronfeldner,M.(2023).Creativitynaturalized:Howtobridgethegapbetweenthenewandthevaluable.PhilosophicalPsychology,36(5),1029–1055.[9]List,C.(2019).WhyFreeWillisReal.HarvardUniversityPress.[10]Simondon,G.(1958).OntheModeofExistenceofTechnicalObjects.Paris:Aubier.7

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

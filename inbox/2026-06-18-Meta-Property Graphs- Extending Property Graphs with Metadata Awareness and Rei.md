---
source: "https://arxiv.org/abs/2410.13813v2"
title: "Meta-Property Graphs: Extending Property Graphs with Metadata Awareness and Reification"
author: "Sepehr Sadoughi, Nikolay Yakovets, George Fletcher"
year: "2024"
publication: "arXiv preprint / cs.DB"
download: "https://arxiv.org/pdf/2410.13813v2"
pdf: "https://arxiv.org/pdf/2410.13813v2"
captured_at: "2026-06-18T11:16:19Z"
updated_at: "2026-06-18T11:16:19Z"
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

# Meta-Property Graphs: Extending Property Graphs with Metadata Awareness and Reification

- 著者: Sepehr Sadoughi, Nikolay Yakovets, George Fletcher
- 年: 2024
- 掲載情報: arXiv preprint / cs.DB
- 情報源: [arxiv](https://arxiv.org/abs/2410.13813v2)
- ダウンロード: https://arxiv.org/pdf/2410.13813v2
- PDF: https://arxiv.org/pdf/2410.13813v2

## Obsidian Links

- 研究動向: [[研究動向/ルカーチ・ジェルジュ-現代研究動向|ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代思想]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[物象化論]]
- 関連分野: [[西洋マルクス主義]]
- 関連タグ: #現代思想 #マルクス主義 #物象化論 #西洋マルクス主義

## Abstract

The ISO standard Property Graph model has become increasingly popular for representing complex, interconnected data. However, it lacks native support for querying metadata and reification, which limits its abilities to deal with the demands of modern applications. We introduce the vision of Meta-Property Graphs, a backwards compatible extension of the property graph model addressing these limitations. Our approach enables first-class treatment of labels and properties as queryable objects and supports reification of substructures in a graph. We propose MetaGPML, a backwards compatible extension of the Graph Pattern Matching Language forming the core of the ISO standard GQL, to query these enhanced graphs. We demonstrate how these foundations pave the way for advanced data analytics and governance tasks that are challenging or impossible with current property graph systems.

## PDF Text

Meta-PropertyGraphs:ExtendingPropertyGraphswithMetadataAwarenessandReiﬁcationSepehrSadoughi,NikolayYakovets,GeorgeFletcherEindhovenUniversityofTechnology,TheNetherlands{s.sadoughi,n.yakovets,g.h.l.fletcher}@tue.nlAbstractTheISOstandardPropertyGraphmodelhasbecomeincreasinglypopularforrepresentingcomplex,intercon-necteddata.However,itlacksnativesupportforqueryingmetadataandreiﬁcation,whichlimitsitsabilitiestodealwiththedemandsofmodernapplications.WeintroducethevisionofMeta-PropertyGraph,abackwardscompatibleextensionofthepropertygraphmodeladdressingtheselimitations.Ourapproachenablesﬁrst-classtreatmentofla-belsandpropertiesasqueryableobjectsandsupportsreiﬁcationofsubstructuresinagraph.WeproposeMetaGPML,abackwardscompatibleextensionoftheGraphPatternMatchingLanguageformingthecoreoftheISOstandardGQL,toquerytheseenhancedgraphs.Wedemonstratehowthesefoundationspavethewayforadvanceddataana-lyticsandgovernancetasksthatarechallengingorimpossiblewithcurrentpropertygraphsystems.1IntroductionModerndataengineeringapplicationsincreasinglydemandsﬂexibilityandagilityinhandlingdataandmetadatamorethaneverbefore.Duringexploratoryanalytics,structureemergesgraduallywithoutapredeﬁnedschema.Heterogene-ityinwhatisadatavalueversuswhatisanattributenameisinherentduringdiscovery,proﬁling,andexploration.Indatasourcesintegrationscenarios,suchasbuildingknowledgegraphs,dataandschemaheterogeneityisinevitable:whatappearsasadatavalueinonesourcemightbeanodelabelorpropertynameinanother,whileasubgraphinonesourcemightcorrespondtoasinglenodeelsewhere.Moderndatamanagementsolutionsmustthereforefullysupportheterogeneityandﬂuidboundariesbetweendataandmetadata.Metadataiscommonlyunderstoodintwodistinctways.Theﬁrstisattributemetadata-characteristicsassociatedwithdataobjects,suchascolumnnamesortablenamesinrelationaldatabasesornodelabelsingraphs.Forinstance,whenstoringcontactinformation,attributeslikenameandemailaremetadata.Thesecondformisreiﬁcation-rep-resentinganaggregationofcomplexrelationshipsorstructuresintoanewentityordataobject,makingiteasiertomanageandanalyze.Inthecontextofentity-relationshipmodeling,reiﬁcationisoftenusedtotransformrelationshipsetsintoentitysets,allowingformoreeffectivedatamodelingandmanipulation.Forexample,inane-commercesystem,reifyingtherelationshipbetweencustomersandorderscreatesanorderhistoryentitythatcanbeanalyzedandqueried.1TheISOstandardPropertyGraph(PGs)modelhasgainedpopularityingraphdatamanagementandiswidelyadopted,e.g.,ingraphDBsystemssuchasNeo4j,Tigergraph,andAmazonNeptune,aswellasrelationalDBsystemssuchasDuckDBwhichimplementtheISOextensionstoSQLforPGquerying.Inapropertygraphnodesandedgesarelabeledandalsohaveassociatedsetsofpropertyname/valuepairs(e.g.,anodelabeledPersonwithpropertyBirthdatehavingdatavalue11-11-2001;herePersonandBirthdateareattributemetadata).WhilePGsofferamodelcloselyalignedwithconceptualdomainrepresentations,themodelmakes(1)astrictdistinctionbetweenmetadataanddata,and(2)hasnosupportforreiﬁcation.Incontrast,theW3C’sResourceDescriptionFramework(RDF)formodelinggraphdatatreatseverythingasdata,includingattributesofnodesandedges,andnativelysupportsreiﬁcation.ThiscapabilitymakesRDFparticularlyusefulforapplicationsrequiringsophisticateddatavalidation,
1Othermorecomplextypesofmetadata,suchasreﬂectiveoractivedata[7,9,18],arebeyondthescopeofourcurrentdiscussion.1
arXiv:2410.13813v2 [cs.DB] 15 Dec 2025
veriﬁcation,querying,andgovernanceofmetadata.ThereisanopportunitytobringthesepowerfulfeaturestoPGswhilepreservingtheircoredesignprinciplesandinherentstrengthsinhandlingdataheterogeneity.Inthispaperwepresentavisionforovercomingbarrierstoﬂexiblemanagementofdata-metadataheterogeneityinPGdatamanagementapplications.Towardsthis,weintroduceMeta-PropertyGraphs(MPG),afullybackwardscompatibleextensionofthePGmodelthataddresseslimitationsinrepresentingandqueryingmetadata.Ourapproachenablesﬁrst-classtreatmentoflabelsandpropertiesasqueryableobjects,aswellasreiﬁcationofsubgraphs.Onthisfoundation,wefurtherproposeMetaGPML,afullybackwardscompatibleextensionoftheGraphPatternMatchingLanguage(GPML),thecorelanguageattheheartoftheISOstandardGQLforPGquerying.WegiveacompleteformalspeciﬁcationofMPGandMetaGPML,providingthefoundationsforourvision.Furthermore,wedemonstratehowthesecontributionsfacilitateadvanceddataanalytics,integration,andgovernancetasksthatarechallengingorimpossiblewithcurrentPGsystems.2RelatedWorkThePGmodelrepresentsadesignpointonthecontinuumbetweentheRelational(ﬁxedmetadata)andRDF(noﬁxedmetadata)datamodels.Whilebydefaultschema-less,supportingagilemodelinganddevelopment,PGsnonethelessmakeastrictdistinctionbetweendatavaluesandmetadatavalues.Ofcourse,allpointsalongthisdesigncontinuumareequallyimportant,ﬁndingtheirapplicationsanduse-cases.Andindeed,datacleaning,wrangling,integration,andexchangeareveryoftenaboutmoving(meta)dataalongthecontinuum.WhathasbeenmissingisanappropriatedesignforPGstofullymeetmodelingandqueryingneedsofdatamanagementscenariossuchasthese(seeChapter2of[5]forasurveyofdesignapproachesforPGmodeling).Intheresearchliterature,dealingwiththechallengesofdata-metadataheterogeneitywasstudiedinthecontextofrelationaldataintegrationanddataintegrationontheweb,leadingtosolutionsforrelationalandXMLdata-metadatamappingandexchange(e.g.,[4,10,15])andseminalworksuchastheSchemaSQLandFISQLdata-metadataquerylanguagesforrelationaldatabases[16,21].OurworkapproachesthesechallengesinthenewcontextoftheISOstandardsforgraphdatamanagement.Inanotherdirection,thereisagrowingbodyofworkonmappingsbetweengraphdatamodels[1,6,14,20,19]andgeneralizingbothPGandRDF[13,17,3].Ourvisionisorthogonalandcomplementarytotheseinvestigations,aimingatextendingthePGmodelandthestandardizedGQLandSQL/PGQquerylanguageswithseamlessdata-metadatafunctionalities.3AguidedtourofworkingwithMPGsFigure1illustratesasimpleMPG,demonstratinghowtheMPGmodelenablesdirectqueryingofmetadataasdataobjects.Thesampledatabaserepresentspublications,indexingdatabases,andpersonswiththeirrelationships,mod-eledusingfourtypesofdataobjects:edges(EG),nodes(NG),properties(PG),andlabels(LG).Notethatpropertiesandlabelsarenowﬁrstclasscitizens,unlikeinthecurrentPGstandard.Table1introducesthepatternnotationusedinourexamples,whichwillbeformallydeﬁnedinSection5.2
Figure1:Examplemeta-propertygraphG
Table1:Dataobjectspatternnotation
Nodes
(x:l)nodevariable(nv.)xwithlabell(x:l).znv.xwithlabellandpropertyvariablez(x::π)nv.xwithapatternπinit’sreiﬁedsubstructure
Edges
-[x:l]->edgevariablexwithlabell-[x:?y]->edgevariablexwithlabelsetvariabley-[x].z->edgevariablexwithpropertyvariablez
Properties
{x}propertyvariablexKEY(x),VAL(x)keyandvalueofpropertyvariablex
Labels
|x|labelsetvariablexcELEMENTOFxcheckiflabelcexistsinlabelsetvariablex
3.1WorkingwithData,Labels,andPropertiesMeta-PropertyGraphenablesustoquerydirectlyonlabelsetsasrecognizeddataobjectswithoutconsideringthenodesoredgesthattheyareassociatedwith.Thiswillfacilitateanalyticsoverclassesinthedatabaseandworkwiththisdataindependently.Asasimpleinstance,inthefollowingexampleMetaGPMLqueryQ1,weretrievethelabelsetthatcontainslabelsPublicationinthedatabasetoﬁgureouttheotherco-occurringtagsregardlessofwhichnodetheybelongto.
Q1:WhichlabelsetscontainthelabelPublication?
MATCH|l|WHERE"Publication"ELEMENTOFlRETURNlAS"Publication_Co_Tags"
3
ResulttableofQ1
Publication
Co
Tags
{”Publication”,”Journal”}
{”Publication”,”Conference”}
Herelisavariablewhichbindstolabelsetobjects.WithMetaGPMLwecanalsoquerypropertiesasindependentdataobjects.QueryQ2demonstratesthisbyretrievingthevaluesofallNameproperties,regardlessofwhetherthepropertyisattachedtoanodeoranedge.
Q2:WhatarethevaluesofNameproperties?
MATCH{p}WHEREKEY(p)="Name"RETURNVAL(p)AS"Names"
ResulttableofQ2
Names
Lee
Scopus
Rose
PubMed
Ingeneral,thiscapabilityofMetaPGMLtobindpropertynamesandlabelsetstovariablesallowstheseamlesspromotionofdatatometadata,andviceversa,asiscriticalinapplicationssuchasdataintegration,cleaning,andgovernance.TheQ3andQ4examplesdemonstratesuchpromotionsanddemotionsof(meta)data.Thefollowingexampleaimstomatchtherelationshipsbetweenpublicationsandtheindexingdatabasesandreturntheresultsinastructuretoshowinwhichindexingdatabaseeachpublicationisbeingarchivedorindexed.
Q3:Whichrelationshipsdoeseachpublicationhavewitheachindexingdatabases?
MATCH(x:Publication)-[:?y]->(z:Indexing_DB)RETURNx.TitleAS"Title",LABEL(y)ASz.Name
Queryresultsareshownbelow,whereLABELresolveslabelsetsboundtoy,whilez.Namevaluesformcolumnheaders:(forconvenienceweswitchbetweentabularandbindingsetrepresentationsofqueryresults)
ResultbindingsofQ3
{(Title�→”NatureStudies”,Scopus�→{”Archived”}),(Title�→”NatureStudies”,PubMed�→{”Indexed”}),(Title�→”BiologyAdvancements”,PubMed�→{”Indexed”})}
Ournextquery(Q4)demonstrateshowtheMetaGPMLcanmanipulatemetadataasregulardata.Thisqueryﬁndspotentialreviewersforpublicationsbasedontheirresearchﬁeldsandthepublication’sacceptancecriteria,showcasinghowMPGtreatsmetadata(propertykeys)asqueryabledata,enablingmoreﬂexibleandpowerfulqueriesthanthePGmodel.4
Q4:Whoarepotentialreviewercandidatesforeachpublicationbasedonthecandidate’sresearchﬁelds?
MATCH(x:Person),(y:Publication).zWHEREx.ResearchField=KEY(z)RETURNx.NameAS"Reviewercandidate",y.NameAS"Publicationvenue",KEY(z)AS"Researchfield"
ResulttableofQ4
Reviewercandidate
Publicationvenue
Researchﬁeld
Lee
NatureStudies
Biology
Lee
Biologyadvancement
Biology
Rose
NatureStudies
Ecology
ThestructureofthegraphrequiresdemotingthemetadataofthepropertykeyinPublication,indicatingwhenitstartedpublishinginacertainresearchﬁeld.ThisallowscomparisonwiththeResearchFieldpropertyofPersonentities,enablingidentiﬁcationofreviewercandidatesfordifferentpublications.3.2WorkingwithReiﬁcationMPGandMetaGPMLenablereiﬁcationasthesecondtypeofmetadataonpropertygraphs.Asub-structureofanMPGcanbereiﬁedasanode,therebymakingitaﬁrst-classcitizen.Thisenablesannotatingpartofthegraphbyassigningpropertiesandlabelstoanodethatreiﬁesthatpart.Itisworthmentioningthatthepartofthegraphthatisbeingreiﬁedmayormaynotbeapropergraph,whichmeansthatanodecanreifysomedataobjectswithouttheirassignedrelationshipsinthegraph,e.g.,asetofspeciﬁcpropertiesorrelationships,canbereiﬁedwithoutthenodestheyareassociatedwith.InFigure1,werepresentthestatement“RoseassignedLeeasarevieweron5thNovember2024”usingnoden4thatinducesasub-structurecontainingnoden1,somerelevantproperties,andoutgoingreviewedges.QueryQ5retrievesthenameofthepersonwhoassignsareviewer:
Q5:WhoassignedLeeasareviewerandwhen?
MATCH(x:Person)-[:assigns]->(y::(z:Person)-[:reviews]->())WHEREz.Name="Lee"RETURNz.NameAS"reviewername",y.DateAS"Date",x.NameAS"Assigningeditor"
NotethatintheMATCHclausewehaveagraphpatternembeddedinanodepattern,todenoteaquerytobeexecutedonthesub-structurereiﬁedbythenodewhichisboundtovariabley.
ResulttableofQ5
Reviewername
Date
Assigningeditor
Lee
05-11-2024
Rose
WhileitmaylookthatMPGisahypergraphmodelit’simportanttonotethathypergraphsfocusoncomplexrelationships,whilereiﬁcationfocusesoncomplexobjects.InMPG,edgesstillconnectexactlytwonodes.Nodescan5
actasmeta-nodes,withotherdataobjects(nodes,edges,properties,labelsets)assignedforreiﬁcation.4Meta-PropertyGraphModelWenextformalizeMeta-PropertyGraphs.LetI,L,K,andVbepairwisedisjointsetsofobjectidentiﬁers,labels,propertykeys,andpropertyvalues,respectively.Deﬁnition4.1Ameta-propertygraphisadirectedandundirectedvertex-andedge-labeledgraphG=(N,E,P,L,λ,µ,σ,υ,η,ρ),where:•N,E,P,L⊆Iareﬁnite,pairwisedisjointsets,•µ:L→2Lassignsaﬁnitesetoflabelstoeachlabelidentiﬁer,•λ:N∪E→Lisabijectivelabelingfunctionassigningalabelsetidentiﬁertoeachnodeandedge,•υ:P→K×Vassignskey-valuepairstoproperties,•σ:N∪E→Cassignscompatiblepropertysetstonodesandedges,suchthatforeachpairofdistinctobjectso1,o2∈N∪E,itholdsthatσ(o1)∩σ(o2)=∅and�o∈N∪Eσ(o)=P,i.e.,everyandeachpropertyp∈Pisassignedtoexactlyonenodeoredge.•η=(ηs,ηt,ηu)where:–ηs,ηt:Ed→Nassignsourceandtargetnodestodirectededges,–ηu:Eu→{{u,v}|u,v∈N}assignsnodepairstoundirectededges,•ρ:N→2N∪E∪P∪Lassignsﬁnitesetsofobjectstonodes,suchthatforeachn∈N,itholdsthatn/∈ρ∗(n),whereρ∗(n)istheclosureofρ(n),2ensuringthatthesub-structureassociatedwitheachnodeiswell-founded.whereE=Ed∪EuthesetofcompatiblepropertysetsC={C⊆P|∀p1,p2∈C,p1̸=p2⇒Key(p1)̸=Key(p2)},Key(p)=π1(υ(p)),Val(p)=π2(υ(p)).Designdecisions.TheconstraintswehaveimposedinourdesignofMPGsareveryminimalistic,ensuringonlythatallpropertiesareuniquelyassignedtoedgesandnodes,andthatreiﬁcationiswell-founded.Wehavedeliberatelyavoidedimposingunwarrantedconstraints.Forinstance,thereisnoconstraintpreventinganodefromhavinganoutgoingedgetoasubgraphcontainingoneofthenode’sownproperties.Thisﬂexibilityallowsforrepresentingscenariossuchas“Marylegallychangedherﬁrstname”or“Maryauditsallnamesinthedatabase(includingherown)”.Similarly,thereisnoconstraintrequiringthatforeachedgeinρ(n),thesourceandtargetoftheedgemustalsobeinρ(n),which,forinstance,canbeparticularlyvaluableinapplicationswithprivacypreservationconcerns.Asanexample,Thispermitsrepresentationslike“MaryisconﬁdentthatBobeditedabook,butnotconﬁdentaboutwhichbook”,whereonlythenoderepresentingBobandBob’soutgoingedgelabeled’edited’areincluded,withoutthetargetofthisedge.Ofcourse,suchadditionalconstraintscanbeadded(orexistingconstraintscanberemoved)asappropriatetothespeciﬁcapplicationdomain.ItisworthnotingthatMPGssupportmeta-properties,i.e.,propertiesonproperties.Forinstance,wecanrepresent“John’sbirthdatewasenteredon23March2020”sinceρcanconsistofjustasingletoncontainingaproperty.Forsimplicity,wehaveomittedspecialtreatmentofmeta-propertiesinourmainpresentation.Wenextformallydeﬁnethegraphsubstructureassociatedwithanode.Deﬁnition4.2(Sub-Structure)LetG=(N,E,P,L,λ,µ,σ,υ,η,ρ)beameta-propertygraphandn∈N.Thesub-structureofGinducedbynisGn=(Nn,En,Pn,Ln,λn,µn,σn,υn,ηn,ρn)where:•Nn=N∩ρ(n),En=E∩ρ(n),Ln=L∩ρ(n),Pn=P∩ρ(n)
2Formally,ρ0(n)=ρ(n),ρi(n)=ρi−1(n)∪�n′∈N∩ρi−1(n)ρ(n′),andρ∗(n)=�∞i=0ρi(n).6
•υn=υ|Pn,µn=µ|Ln(domainrestrictions)•λn(o)=λ(o)foro∈Nn∪Enifλ(o)∈ρ(n),undeﬁnedotherwise•σn(o)=σ(o)∩ρ(n)foro∈Nn∪En•Edn,Eun⊆Enaredirectedandundirectededgesubsets•ηn=(ηsn,ηtn,ηun)where:–Fore∈Edn:*ηsn(e)=ηs(e)ifηs(e)∈ρ(n),undeﬁnedotherwise*ηtn(e)=ηt(e)ifηt(e)∈ρ(n),undeﬁnedotherwise–ηun(e)=ηu(e)∩ρ(n)fore∈Eun•ρn(n′)=ρ(n)∩ρ(n′)forn′∈Nn,satisfyingthesamewell-foundednesspropertyasinDeﬁnition4.1.Intuitively,thesub-structureGncanbethoughtofasa”slice”or”view”oftheoriginalmeta-propertygraphG,focusedonthenodenanditsassociatedelements.Thissubstructureincludesonlythosenodes,edges,properties,andlabelsthatareexplicitlylinkedtonthroughthereiﬁcationfunctionρ.Itmaintainstheoriginalgraph’sstructureandrelationships,butonlyfortheelementswithinthissubset.Thisallowsustozoominonspeciﬁcpartsofthegraph,makingiteasiertoanalyzeandquerycomplex,nestedstructureswithoutlosingthecontextofhowtheyﬁtintothelargergraph.5Meta-PropertyGraphPatternMatchingLanguageWenextformalizeourGraphPatternMatchingLanguageforMetapropertyGraphs(MetaGPML)asanextensiontotheGraphPatternMatchingLanguage(GPML)[8,12,11].GPMListhecoreelementofbothSQL/PGQandGQLlanguagesstandardizedbyISO.ByextendingGPMLweensurebackwardcompatibilityinthesensethateveryGPMLqueryisaMetaGPMLquery.5.1SyntaxLetXbeacountableinﬁnitesetofvariables.Weintroducedescriptorsδ1,δ2,andδ3inMetaGPML,consistentwithGPML[12].Forx,y,z∈X,ℓ∈L,k∈K,andv∈V,wehave
Descriptors:
δnodes:=δedges|::π|:ℓ::π|x::π|x:ℓ::π|:?y::π|x:?y::πδedges:=x|:ℓ|:?y|x:ℓ|x:?y(1)
Thedescriptorsin(1)willbeusedinthefollowingpatternstoassignvariablesandconstantstodifferentdataobjects.ApatternπinMetaGPMLandoperationsonitcanbeformulatedasfollows:7
Patterns:
π:=(δnodes)|()(nodepattern)|(δnodes).x|().x|-[δedges]->|<-[δedges]-(edgepattern)|-[δedges]-|-[]->|<-[]-|-[]-|-[δedges].x->|<-[δedges].x-|-[δedges].x-|{x}|{}(propertypattern)||x||||(labelpattern)|ππ(concatenation)|π+π(union)|πWHEREΦ(conditioning)Π:=π|Π,Π(graphpattern)(2)
WedeﬁneEastheexpressionforconditionΦthatcanbeusedinapatternaftertheWHEREkeywordasfollows,consideringc∈K∪V∪L.
Expressionsandconditions:
E:=x|x.k|v|k|ℓ(expressions)|KEY(x)|VAL(x)|LABEL(x)Φ:=E=E|E<E|x:ℓ(conditions)|SUBSETEQ(x,y)|cELEMENTOFy|ΦANDΦ|ΦORΦ|NOTΦ(3)
ItisimportanttonotethatKEY(x)andVAL(x)operateonpropertiesasﬁrst-classdataobjects,similartonodesandedges.ThesameappliestoLABEL(x),whichoperatesonlabelsetsasdistinguishabledataobjectsandretrievesthesetoflabelstringsitcontains.ThesefunctionsshouldnotbeconfusedwithsimilarfunctionsinsomeotherPGquerylanguageslikeOpenCypher,whichonlysimulatesthesecapabilitiesatthequerylanguagelevelwithintheconventionalPGdatamodel.Finally,basedon(2)and(3),weformulatethedeﬁnitionofclausesandqueries.
Clausesandqueries:
C:=MATCHΠ|FILTERΦQ:=CQ|RETURNE1ASx1,...,EnASxn(4)
ThesyntaxofMetaGPMLisdesignedtobeintuitiveandexpressive,buildinguponfamiliarconceptsfromtheGQLstandardquerylanguage.Nodeandedgepatternsuseparenthesesandbracketsrespectively,withlabelspreﬁxedbycolons.Thedouble-colonsyntax(::)introducessubgraphpatterns,enablingqueryingofreiﬁedstructures.Todistinguishthemfromconstants,variablesarepreﬁxedwithquestionmarkswhenusedtobindlabelsorproperties.ThelanguageincludesspecializedpredicateslikeELEMENTOFandSUBSETEQforworkingwithsetsoflabelsandproperties.5.2SemanticsNext,wepresentthesemanticsofourlanguageasaclear,unambiguousfoundationforfurtherstudyandimplementa-tion.Toavoidunnecessarycomplexity,wepresentedthemostvitalsemanticstoillustratetheconcepttounderstandhowMetaGPMLworks.ThecompleteandmoredetailedsemanticscoveringallsituationsarepresentedinappendixA.8
Deﬁnition5.1(Bindings)Abindingβ:X→Vassignsvariablesx∈Xtovaluesv∈V,whereV=I∪L∪K∪V.Wedenotethisas(x1�→v1,...,xn�→vn),wherex1,...,xnarethevariablesinthedomainofβ(Dom(β))andv1,...,vnaretheircorrespondingvalues.Wedenotetheemptybindingas().Deﬁnition5.2(CompatibilityofBindingsandtheirJoin)Twobindingsβ1,β2arecompatible,denotedbyβ1∼β2,iftheyagreeontheirsharedvariables.Speciﬁcally,foreveryx∈Dom(β1)∩Dom(β2),itholdsthatβ1(x)=β2(x).Whenβ1∼β2,theirjoinβ1⋊⋉β2isdeﬁnedasfollows:Dom(β1⋊⋉β2)=Dom(β1)∪Dom(β2).Foranyvariablex,(β1⋊⋉β2)(x)=β1(x)ifx∈Dom(β1)\Dom(β2),and(β1⋊⋉β2)(x)=β2(x)ifx∈Dom(β2).Apatternmatchingsemantic�π�Ginmeta-propertygraphisasetofbindingsβthatassignvariablestovalues.
Nodepatternmatching:
�()�G={()|n∈N}(5)�(x)�G={(x�→n)|n∈N}(6)�(x:ℓ)�G={(x�→n)|n∈N,ℓ∈µ(λ(n))}(7)�(x:?y)�G={(x�→n,y�→l)|n∈N,l=λ(n)}(8)�(x).z�G={(x�→n,z�→p)|n∈N,p∈σ(n)}(9)
Equation5matchesanynoden∈N.Equation6assignsanodetovariablex.Equation7matchesnodeswithspeciﬁclabels,whereµ(λ(n))determinesthenode’slabels.Equation8assignsanode’slabelsettovariabley.FinallyEquation9marchesnodeswithpropertypassignedtovariablez.Meta-nodesareakeyconceptinMetaGPML,deﬁnedasnodestowhichdiverseelementsareassignedbasedontheconceptofsub-structure.ThisfeatureenablesreiﬁcationwithintheMeta-PropertyGraph,allowingforamorenuancedrepresentationofcomplexrelationshipsandstructures.Equation10demonstratesthematchingofnodeswithlabelstopatternswithintheirassociatedsub-structure:
Meta-nodepatternmatching:
�(x:ℓ::π)�G=�(x:ℓ)�G⋊⋉�π�Gn(10)
Themeta-nodesemanticdeﬁnitioninvolvesafulljoinoperationbetweentwopatternsdividedby::inthenodepatterndescriptor.Here,�(x:ℓ)�Grepresentsasimplenodepatternmatching,while�π�GnisapatternmatchingwithintheGnsub-structure.Wedeﬁnethesemanticsofedgesinourmodel,supportingbothdirectedandundirectededges.
Edgepatternmatching:
�-[]->�G={()|e∈Ed}(11)�-[x]->�G={(x�→e)|e∈Ed}(12)�-[:ℓ]->�G={()|e∈Ed,ℓ∈µ(λ(e))}(13)�-[x:ℓ]-�G={(x�→e)|e∈Eu,ℓ∈µ(λ(e))}(14)�-[x:?y]->�G={(x�→e,y�→l)|e∈Ed,l=λ(e)}(15)�-[x].z->�G={(x�→e,z�→p)|e∈Ed,p∈σ(e)}(16)
Equations11-59matchvariousedgepatterns,includingdirectedandundirectededgeswithorwithoutlabels.Equation15allowsassigningtheedge’slabelsettovariabley,while16enablesassigningtheedge’spropertiestovariablez.9
Themeta-propertygraphmodelenablesdeﬁningspeciﬁcsemanticsforpropertyandlabelobjects,enhancingquerycapabilitiesontheseelements.
Propertyandlabelobjects:
�{x}�G={(x�→p)|p∈P}(17)�|x|�G={(x�→l)|l∈L}(18)
Patternconcatenation,union,andconditioning:
�π1π2�G={β1⋊⋉β2|βi∈�πi�G,i=1,2andβ1∼β2}(19)�π1+π2�G={β∪β′|β∈�π1�G∪�π2�G}(20)�πWHEREΦ�G={β∈�π�G|�Φ�βG=True}(21)
In(20),β′mapsanyvariablein�π1+π2�Gnotinβ’sdomaintoNull.
Graphpatterns:
�Π1,Π2�G={β1⋊⋉β2|βi∈�Πi�G,i=1,2andβ1∼β2}(22)
Thesemantics�E�βGofanexpressionEiscomputedwithrespecttobindingβovergraphG.Foravariablexinβ’sdomain,�x�βG=β(x).
Nodeandpropertyvalues:
�x.k�βG=�Val(p)ifβ(x)∈N∪E,p∈σ(β(x)),Key(p)=kNullotherwise(23)
PropertyandLabelsetoperations:
�KEY(x)�βG=Key(�x�βG)for�x�βG∈P(24)�VAL(x)�βG=Val(�x�βG)for�x�βG∈P(25)�LABEL(x)�βG=µ(�x�βG)for�x�βG∈L(26)
Thesemantics�Φ�βGofaconditionΦisanelementin{True,False,Null},computedwithrespecttoβasfollows.10
Conditions:
�E1=E2�βG=Nullif�E1�βG=Nullor�E2�βG=NullTrueif�E1�βG=�E2�βGFalseotherwise(27)�E1<E2�βG=Nullif�E1�βG=Nullor�E2�βG=NullTrueif�E1�βG<�E2�βGFalseotherwise(28)�x:ℓ�βG=�Trueif�x�βG∈N∪Eandℓ∈µ(λ(�x�βG))Falseif�x�βG∈N∪Eandℓ/∈µ(λ(�x�βG))(29)
Logicaloperations:
�Φ1ANDΦ2�βG=�Φ1�βG∧�Φ2�βG(30)�Φ1ORΦ2�βG=�Φ1�βG∨�Φ2�βG(31)�NOTΦ�βG=¬�Φ�βG(32)
Setoperation:
�cELEMENTOFy�βG=�Truec∈µ(�y�βG)Falseotherwise(33)�SUBSETEQ(x,y)�βG=�Trueµ(�x�βG)⊆µ(�y�βG)Falseotherwise(34)
Deﬁnition5.3MetaGPMLevaluatesqueriesandclausesusingworkingtables,ensuringconsistencywithGQLstan-dards.AtableTcomprisesbindingswithshareddomains.Finally,wedeﬁnethesemanticsofclauseCandqueryQasfunctionsoperatingontableTforgraphG.
Clausesandqueries:
�MATCHΠ�G(T)=�β∈T{β⋊⋉β′|β′∈�Π�G,β∼β′}(35)�FILTERΦ�G(T)={β∈T|�Φ�βG=True}(36)�CQ�G(T)=�Q�G(�C�G(T))(37)�RETURNE1ASx1,...,EnASxn�G(T)=�β∈T{(x1�→�E1�βG,...,xn�→�En�βG)}(38)
ThesemanticsofMetaGPMLbuilduponandextendthoseofGPML,providingarobustframeworkforqueryingmeta-propertygraphs.Keyfeaturesincludetheabilitytomatchnodesandedgeswithspeciﬁclabelsorproperties,handlemeta-nodesthatencapsulatesubgraphs,andperformoperationsonpropertyandlabelobjects.Thelanguagesupportspatternconcatenation,union,andconditioning,aswellassetoperationslikeELEMENTOFandSUBSETEQ.11
ClausessuchasMATCHandFILTER,alongwiththeRETURNstatement,allowforcomplexqueryconstruction.Thesefoundationsenablepowerfulandﬂexiblequeryingcapabilities,facilitatingadvancedanalyticsandmetadatamanagementinPGdatabases.6NextstepstowardstheMPGvisionInthispaper,wehighlightedthecriticalneedtobreakdownbarriersbetweendataandmetadatainpropertygraphmanagement.Astheﬁrstconcretestepstowardsthisvision,weintroducedafullyspeciﬁeddatamodelandquerylanguageformeta-propertygraphs,enablingseamlessmodelingandinteroperationofdataandmetadata.Whilethisworklaysthefoundationforﬂexiblepropertygraphmanagementforcontemporaryapplications,furtherresearchisneededtofullyrealizethisvision.(1)Physicalimplementationandtechnicalchallenges.AkeychallengeisimplementingtheMPGdatamodelefﬁciently.Sincelabelsetsandpropertiesaretreatedasdataobjectswithidentiﬁers,theymayneeddedicatedstorageandindexingstrategies.Alternatively,approacheslikeconcatenatedIDscouldmaintainexistingstoragestructuresbutmayimpactqueryevaluation.ResearchisneededonphysicalrepresentationsandindexingstrategiesthatoptimizeMPGperformance.(2)Meta-PropertyGraphsinpractice.WeneedtoinvestigatehowMeta-PropertyGraphcanenhanceknowledgeengineeringandmanagementinpractice.Understandinghowmetadataawarenessandsubstructurereiﬁcationcancon-tributetoimprovingtaskssuchasauditingandhuman-in-the-loopvalidationofknowledgegraphsordatacleaning,wrangling,integration,andexchangeiscrucial.Furthermore,MPGscanfacilitateknowledgereasoningandanalysisonknowledgegraphsbyintroducinginherentsubgraphannotationandqueryingdifferentformsofmetadata.Addi-tionally,developingeffectiveeducationalapproachesandtrainingresourcesforstudentsandprofessionalsworkingwithMPGsandMetaGPMLrequiresfurtherstudy.(3)Improvementsandintegration.Meta-PropertyGraphandMetaGPMLcanbeenhancedthrough:(1)extendingMetaGPMLwithadditionalfunctionstobetterleveragemetadataawarenessandexistingpropertygraphcapabilitieslikepaths,(2)developingschemaandconstraintlanguagesforMPGbuildingonPG-SCHEMA[2],and(3)incorpo-ratingotherformsofmetadatalikereﬂectiontoexpandthemetadataawarenessinMPG.References[1]GhadeerAbuoda,DanieleDell’Aglio,ArthurKeen,andKatjaHose.TransformingRDF-startopropertygraphs:Apreliminaryanalysisoftransformationapproaches–extendedversion.InQuWeDaWorkshoponStoring,QueryingandBenchmarkingKnowledgeGraphs,2022.[2]RenzoAngles,AngelaBonifati,StefaniaDumbrava,GeorgeFletcher,AlastairGreen,JanHidders,BeiLi,LeonidLibkin,VictorMarsault,WimMartens,FilipMurlak,StefanPlantikow,OgnjenSavkovic,MichaelSchmidt,JuanSequeda,SlawekStaworko,DominikTomaszuk,HannesVoigt,DomagojVrgoc,MingxiWu,andDusanZivkovic.Pg-schema:Schemasforpropertygraphs.Proc.ACMManag.Data,1(2):198:1–198:25,2023.[3]RenzoAngles,AidanHogan,OraLassila,CarlosRojas,DanielSchwabe,PedroSzekely,andDomagojVrgoˇc.Multilayergraphs:Auniﬁeddatamodelforgraphdatabases.InProceedingsofthe5thACMSIGMODJointInternationalWorkshoponGraphDataManagementExperiences&Systems(GRADES)andNetworkDataAnalytics(NDA).AssociationforComputingMachinery,2022.[4]AngelaBonifati,ElaineQingChang,TerenceHo,LaksV.S.Lakshmanan,RachelPottinger,andYongikChung.SchemamappingandquerytranslationinheterogeneousP2PXMLdatabases.VLDBJ.,19(2):231–256,2010.12
[5]AngelaBonifati,GeorgeFletcher,HannesVoigt,andNikolayYakovets.Queryinggraphs.Morgan&Claypool,2018.[6]JulianBruyat,Pierre-AntoineChampin,LionelM´edini,andFr´ed´eriqueLaforest.Prec:semantictranslationofpropertygraphs.1stworkshoponSquaringtheCirclesonGraphs,2021.[7]JanVandenBussche,DirkVanGucht,andStijnVansummeren.Acrashcourseondatabasequeries.InProceed-ingsoftheTwenty-SixthACMSIGACT-SIGMOD-SIGARTSymposiumonPrinciplesofDatabaseSystems,pages143–154.ACM,2007.[8]AlinDeutsch,NadimeFrancis,AlastairGreen,KeithHare,BeiLi,LeonidLibkin,TobiasLindaaker,VictorMarsault,WimMartens,JanMichels,FilipMurlak,StefanPlantikow,PetraSelmer,OskarvanRest,HannesVoigt,DomagojVrgoˇc,MingxiWu,andFredZemke.GraphpatternmatchinginGQLandSQL/PGQ.InProceedingsofthe2022InternationalConferenceonManagementofData,SIGMOD’22,pages2246–2258.AssociationforComputingMachinery,2022.[9]GeorgeFletcher.Onreﬂectioninlinkeddatamanagement.InWorkshopsProceedingsofthe30thInternationalConferenceonDataEngineeringWorkshops,ICDE2014,Chicago,IL,USA,pages269–271.IEEEComputerSociety,2014.[10]GeorgeFletcherandCatharineWyss.Towardsageneralframeworkforeffectivesolutionstothedatamappingproblem.JournalonDataSemanticsXIV,pages37–73,2009.[11]NadimeFrancis,AmelieGheerbrant,PaoloGuagliardo,LibkinLeonid,VictorMarsault,WimMartens,FilipMurlak,LiatPeterfreund,AlexandraRogova,andDomagojVrgoˇc.Aresearcher’sdigestofGQL.In26thInternationalConferenceonDatabaseTheory(ICDT2023),volume255,2023.[12]NadimeFrancis,Am´elieGheerbrant,PaoloGuagliardo,LeonidLibkin,VictorMarsault,WimMartens,FilipMurlak,LiatPeterfreund,AlexandraRogova,andDomagojVrgoc.Gpc:Apatterncalculusforpropertygraphs.InProceedingsofthe42ndACMSIGMOD-SIGACT-SIGAISymposiumonPrinciplesofDatabaseSystems,PODS’23,pages241–250.AssociationforComputingMachinery,2023.[13]EwoutGelling,GeorgeFletcher,andMichaelSchmidt.Statementgraphs:Unifyingthegraphdatamodelland-scape.InDatabaseSystemsforAdvancedApplications-29thInternationalConference,DASFAA2024,pages364–376,2024.[14]OlafHartig.ReconciliationofRDF*andpropertygraphs.arXiv:1406.3399v3,2014.[15]MauricioA.Hern´andez,PaoloPapotti,andWangChiewTan.Dataexchangewithdata-metadatatranslations.Proc.VLDBEndow.,1(1):260–273,2008.[16]LaksV.S.Lakshmanan,FereidoonSadri,andSubbuN.Subramanian.SchemaSQL:AnextensiontoSQLformultidatabaseinteroperability.ACMTransactionsonDatabaseSystems(TODS),26(4):476–519,2001.[17]O.Lassila,M.Schmidt,B.Bebee,D.Bechberger,W.Broekema,A.Khandelwal,K.Lawrence,R.Sharda,andB.Thompson.TheOneGraphvision:Challengesofbreakingthegraphmodellock-in.SemanticWebJournal,14,2023.[18]DiveshSrivastavaandYannisVelegrakis.Intensionalassociationsbetweendataandmetadata.InProceedingsoftheACMSIGMODInternationalConferenceonManagementofData,pages401–412.ACM,2007.[19]RaduStoica,GeorgeFletcher,andJuanF.Sequeda.Ondirectlymappingrelationaldatabasestopropertygraphs.InAMW,2019.[20]DominikTomaszuk,RenzoAngles,andHarshThakkar.PGO:DescribingpropertygraphsinRDF.IEEEAccess,8:118355–118369,2020.[21]CatharineM.WyssandEdwardL.Robertson.Relationallanguagesformetadataintegration.ACMTransactionsonDatabaseSystems,30(2),2005.13
AAppendixThisappendixprovidesanin-depthexplorationofthesemanticsmissingfromsection5.2,supplementingtheconciseoverviewpresentedinthemainbodyofthepaper.Forthoseinterestedinamoredetailedunderstanding,thefollowingsectionsoffercomprehensiveinsightsintotheformalstructuresandrulesthatunderpinthequerylanguage’sdesignandfunctionality.Thisadditionalmaterialaimstoenhancethediscussion’sclarityandcompletenesswithoutdisrupt-ingthemaintext’snarrativeﬂow.�(:ℓ)�G={()|n∈N,ℓ∈µ(λ(n))}(39)�(::π)�G=�()�G⋊⋉�π�Gn(40)�(:ℓ::π)�G=�(:ℓ)�G⋊⋉�π�Gn(41)�(x::π)�G=�(x)�G⋊⋉�π�Gn(42)�(:?y)�G={(y�→l)|n∈N,l=λ(n)}(43)�(:?y::π)�G=�(:?y)�G⋊⋉�π�Gn(44)�(x:?y::π)�G=�(x:?y)�G⋊⋉�π�Gn(45)�().z�G={(z�→p)|n∈N,p∈σ(n)}(46)�(:ℓ).z�G={(z�→p)|n∈N,ℓ∈µ(λ(n)),p∈σ(n)}(47)�(::π).z�G=�().z�G⋊⋉�π�Gn(48)�(x:ℓ).z�G={(x�→n,z�→p)|n∈N,ℓ∈µ(λ(n)),p∈σ(n)}(49)�(:ℓ::π).z�G=�(:ℓ).z�G⋊⋉�π�Gn(50)�(x::π).z�G=�(x).z�G⋊⋉�π�Gn(51)�(x:ℓ::π).z�G=�(x:ℓ).z�G⋊⋉�π�Gn(52)�(:?y).z�G={(y�→l,z�→p)|n∈N,l=λ(n),p∈σ(n)}(53)�(x:?y).z�G={(x�→n,y�→l,z�→p)|n∈N,l=λ(n),p∈σ(n)}(54)�(:?y::π).z�G=�(:?y).z�G⋊⋉�π�Gn(55)�(x:?y::π).z�G=�(x:?y).z�G⋊⋉�π�Gn(56)�-[]-�G={()|e∈Eu}(57)�-[x]-�G={(x�→e)|e∈Eu}(58)�-[:ℓ]-�G={()|e∈Eu,ℓ∈µ(λ(e))}(59)�-[:?y]-�G={(y�→l)|e∈Eu,l=λ(e)}(60)�-[x:?y]-�G={(x�→e,y�→l)|e∈Eu,l=λ(e)}(61)�-[].z-�G={(z�→p)|e∈Eu,p∈σ(e)}(62)14
�-[x].z-�G={(x�→e,z�→p)|e∈Eu,p∈σ(e)}(63)�-[:ℓ].z-�G={(z�→p)|e∈Eu,ℓ∈µ(λ(e)),p∈σ(e)}(64)�-[x:ℓ].z-�G={(x�→e,z�→p)|e∈Eu,ℓ∈µ(λ(e)),p∈σ(e)}(65)�-[:?y].z-�G={(y�→l,z�→p)|e∈Eu,l=λ(e),p∈σ(e)}(66)�-[x:?y].z-�G={(x�→e,y�→l,z�→p)|e∈Eu,l=λ(e),p∈σ(e)}(67)�-[x:ℓ]->�G={(x�→e)|e∈Ed,ℓ∈µ(λ(e))}(68)�-[:?y]->�G={(y�→l)|e∈Ed,l=λ(e)}(69)�-[].z->�G={(z�→p)|e∈Ed,p∈σ(e)}(70)�-[:ℓ].z->�G={(z�→p)|e∈Ed,ℓ∈µ(λ(e)),p∈σ(e)}(71)�-[x:ℓ].z->�G={(x�→e,z�→p)|e∈Ed,ℓ∈µ(λ(e)),p∈σ(e)}(72)�-[:?y].z->�G={(y�→l,z�→p)|e∈Ed,l=λ(e),p∈σ(e)}(73)�-[x:?y].z->�G={(x�→e,y�→l,z�→p)|e∈Ed,l=λ(e),p∈σ(e)}(74)15

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

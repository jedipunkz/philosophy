---
source: "https://arxiv.org/abs/1406.3399v3"
title: "Foundations of an Alternative Approach to Reification in RDF"
author: "Olaf Hartig, Bryan Thompson"
year: "2014"
publication: "arXiv preprint / cs.DB"
download: "https://arxiv.org/pdf/1406.3399v3"
pdf: "https://arxiv.org/pdf/1406.3399v3"
captured_at: "2026-06-18T11:16:15Z"
updated_at: "2026-06-18T11:16:15Z"
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

# Foundations of an Alternative Approach to Reification in RDF

- 著者: Olaf Hartig, Bryan Thompson
- 年: 2014
- 掲載情報: arXiv preprint / cs.DB
- 情報源: [arxiv](https://arxiv.org/abs/1406.3399v3)
- ダウンロード: https://arxiv.org/pdf/1406.3399v3
- PDF: https://arxiv.org/pdf/1406.3399v3

## Obsidian Links

- 研究動向: [[研究動向/ルカーチ・ジェルジュ-現代研究動向|ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代思想]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[物象化論]]
- 関連分野: [[西洋マルクス主義]]
- 関連タグ: #現代思想 #マルクス主義 #物象化論 #西洋マルクス主義

## Abstract

This document defines extensions of the RDF data model and of the SPARQL query language that capture an alternative approach to represent statement-level metadata. While this alternative approach is backwards compatible with RDF reification as defined by the RDF standard, the approach aims to address usability and data management shortcomings of RDF reification. One of the great advantages of the proposed approach is that it clarifies a means to (i) understand sparse matrices, the property graph model, hypergraphs, and other data structures with an emphasis on link attributes, (ii) map such data onto RDF, and (iii) query such data using SPARQL. Further, the proposal greatly expands both the freedom that database designers enjoy when creating physical indexing schemes and query plans for graph data annotated with link attributes and the interoperability of those database solutions.

## PDF Text

arXiv:1406.3399v3 [cs.DB] 16 Dec 2021
FoundationsofanAlternativeApproachtoReiﬁcationinRDFOlafHartigBryanThompsonLink¨opingUniversityAmazonWebServiceshttp://olafhartig.deDecember17,2021
Thisdocumenthasbecomeobsolete.Theideasandformalizationsinthisdocumenthavebeenpickedupanddevelopedfurtherbyataskforce
1
oftheRDF-DEVcommunitygroup
2
attheWorldWideWebconsortium(W3C).Inthiscontext,theapproachhasbeenrenamedtoRDF-starandSPARQL-star.Themainresultofthegroupisacommunitygroupreport
3
thatservesasaspeciﬁcationoftheapproach.Additionally,thegrouphasdevelopedacomprehensivecollectionoftestsuites
4
fordiﬀerentRDF-starsyntaxes,fortheRDF-starsemantics,andforSPARQL-star.NoticethattherearesomeaspectsinwhichRDF-starandSPARQL-star,asdeﬁnedintheaforementionedspeciﬁcation,diﬀerfromRDF?andSPARQL?asdeﬁnedinthisdocument.Mostnotably,thesemanticsofembeddedtriples—whicharenowcalledquotedtriples—isnotanymoredeﬁnedintermsofstandardRDFreiﬁcationandsuchaquotedtripleisnotanymoreconsideredtobeimplicitlyassertedinanRDF-stargraphthatcontainsnestedtriplesofwhichthequotedtripleisaconstituent.AnotherdiﬀerenceisthattheadditionalnotionofBINDclausesasdeﬁnedforSPARQL?inthisdocument(cf.Example
4
)hasnotbeencarriedovertoSPARQL-star(however,SPARQL-starhasanewbuilt-infunctioncalledTRIPLEwhichcanbeusedforthesamepurpose).
AbstractThisdocumentdeﬁnesextensionsoftheRDFdatamodelandoftheSPARQLquerylanguagethatcaptureanalternativeapproachtorepresentstatement-levelmetadata.Whilethisalter-nativeapproachisbackwardscompatiblewithRDFreiﬁcationasdeﬁnedbytheRDFstandard,theapproachaimstoaddressusabilityanddatamanagementshortcomingsofRDFreiﬁcation.Oneofthegreatadvantagesoftheproposedapproachisthatitclariﬁesameansto(i)under-standsparsematrices,thepropertygraphmodel,hypergraphs,andotherdatastructureswithanemphasisonlinkattributes,(ii)mapsuchdataontoRDF,and(iii)querysuchdatausingSPARQL.Further,theproposalgreatlyexpandsboththefreedomthatdatabasedesignersenjoywhencreatingphysicalindexingschemesandqueryplansforgraphdataannotatedwithlinkattributesandtheinteroperabilityofthosedatabasesolutions.1IntroductionΩTheRDFstandardintroducesthenotionofreiﬁcationasanapproachtoprovideasetofRDFtriplesthatdescribesomeotherRDFtriple[HPS14].Thisformofstatement-levelmetadataabout
1
https://w3c.github.io/rdf-star/
2
https://www.w3.org/community/rdf-dev/
3
https://w3c.github.io/rdf-star/cg-spec
4
https://w3c.github.io/rdf-star/tests/
1
areiﬁedtriplehastoincludefouradditionalRDFtriplestorefertothereiﬁedtriple.Example1.ConsiderthefollowingtwoRDFtriples—giveninTurtlesyntax[PC14]—thatindicatetheageofsomebodynamedBob5:bobfoaf:name"Bob";foaf:age23.TocapturemetadataaboutagivenRDFtripleasperRDFreiﬁcation,wehavetointroduceanIRIorablanknodeandusethisIRIorblanknodeasthesubjectoffourRDFtriplesthatreifythegiventriplebyapplyingtheRDFreiﬁcationvocabulary[HPS14].Then,theIRIorblanknodecanbeusedtoprovidemetadataaboutthereiﬁedtriple.Forinstance,byusingablanknode,say
:s,wemayreifythesecondofthetwoexampletriplesasfollows:_:srdf:typerdf:Statement;rdf:subject:bob; rdf:predicatefoaf:age; rdf:object23.Now,wecanusetheblanknodetoprovidemetadataaboutthetriple:_:sdct:creator<http://example.com/crawlers#c1>;dct:source<http://example.net/homepage-listing.html>.RDFreiﬁcationasdemonstratedintheexamplehastwomajorshortcomings.First,addingfourreiﬁcationtriplesforeveryreiﬁedtripleisineﬃcientforexchangingaswellasformanagingRDFdatathatincludesstatement-levelmetadata.Second,writingqueriestoaccessstatement-levelmetadataiscumbersomebecauseanymetadata-related(sub)expressioninaqueryhastobeac-companiedbyanothersubexpressiontomatchthecorrespondingfourreiﬁcationtriples.Example2.Considerthedata(includingthemetadata)fromExample1.ToquerythisdatawemayuseSPARQL,thestandardquerylanguageforRDF[HSP13].Forinstance,ifweareinterestedintheageofBob,includingthesourceofthisinformation,wemaywriteaSPARQLquerysuchasthefollowing.SELECT?age?srcWHERE{?bobfoaf:name"Bob";foaf:age?age.?rrdf:typerdf:Statement;rdf:subject?b; rdf:predicatefoaf:age; rdf:object?age; dct:source?src.}Notethatthegivenquerycontainsfourtriplepatternstoidentifythereiﬁedtriplewhosemetadatawewanttosee.Ifwewerealsointerestedinpotentialmetadataaboutthecorrespondingfoaf:nametriple,wewouldhavetoaddanotherfour,reiﬁcation-relatedtriplepatterns.
5Preﬁxdeclarationsareomittedinallexamplesinthisdocument.Thepreﬁxesusedaretheusualpreﬁxesascanbefoundviathe http://prefix.cc service.2
DuringaDagstuhlseminaron“SemanticDataManagement”[ACA+12]severalparticipantsoftheseminar—includingBryanThompson(Systap,LLC),OrriErlingandYrj¨an¨aRankka(Open-LinkSoftware),andOlafHartig(thenHumboldtUniversit¨atzuBerlin)—discussedanalternativeapproachtoreiﬁcationthataddressestheaforementionedshortcomings.Thisdocumentprovidesaformalfoundationforthisapproach.Thedocumentisstructuredasfollows:Section2outlinestheapproachinformally.Thereafter,Section3introducesanextensionoftheRDFdatamodelthatmakesmetadatastatementsaﬁrstclasscitizen.Finally,Sections4and5extendsthequerylanguageSPARQLtoenableuserstobeneﬁtfromtheextendeddatamodel.2AnAlternativeApproachtoReiﬁcationinRDFΩThealternativeapproachtoreiﬁcationisbasedontheideaofusingatripledirectlyinthesubjectpositionorobjectpositionof(other)triplesthatrepresentmetadataabouttheembeddedtriple.Example3.AssumeapossibleextensionoftheTurtlesyntaxthatimplementstheideaofembeddingRDFtriplesintootherRDFtriplesbyenclosinganyembeddedtriplein’<<’and’>>’(Section3.3shallintroducesuchanextension).Then,thedatafromExample1(includingthemetadata)couldberepresentedasfollows.:bobfoaf:name"Bob".
<<:bobfoaf:age23>>dct:creator<http://example.com/crawlers#c1>;dct:source<http://example.net/homepage-listing.html>.Embeddingtriplesinto(metadata)triplesasdemonstratedintheexampleachievesamorecom-pactrepresentationofstatement-levelmetadatathanstandardRDFreiﬁcation.SuchacompactrepresentationmayimprovecomprehensibilityforuserswhohavetoinspectRDFdocumentsdi-rectly(e.g.,inatexteditor).SucharepresentationmayalsoreducethesizeofRDFdocumentsthatincludestatement-levelmetadataand,thus,maybeadvantageousfordataexchange.Further-more,embeddedtriples,conceivedofasaformofself-referencingidentiﬁers,correspondsnaturallytotheconceptoftripleidentiﬁersthatsomeRDFdatamanagementsystemssuchasSystap’sBigdata[TPC14]useinternallytoavoidtheoverheadofkeepingaphysicalrepresentationoffourreiﬁcationtriplesperreiﬁedtriple.Giventhattriplesareembeddedintoothertriples,theideaofsuchanembeddingcanbecarriedovertoSPARQLqueries.
Example4.ByadoptingtheextendedsyntaxoutlinedinExample3,wecouldrepresentthequeryfromExample2inthefollowing,morecompactform.SELECT?age?srcWHERE{?bobfoaf:name"Bob".
<<?bobfoaf:age?age>>dct:source?src.}AnalternativeformistouseBINDclausesasfollows.3
SELECT?age?srcWHERE{?bobfoaf:name"Bob".
BIND(<<?bobfoaf:age?age>>AS?t)
?tdct:source?src.}Theremainderofthisdocumentprovidesaformaldeﬁnitionoftheapproachoutlinedinthissec-tion.AnimportantcharacteristicofthisformalizationisitsbackwardcompatibilitywithstandardRDFreiﬁcation.
3RDF?–AMetadataExtensionofRDFThissectionintroducesanextensionoftheRDFdatamodel[CWL+14]thatmakesmetadatastate-mentsaﬁrstclasscitizen.Hereafter,theextendeddatamodelisreferredtoasRDF?.3.1ConceptsΩAssumepairwisedisjointsetsI(allIRIs),B(blanknodes),andL(literals).Asusual,anRDFtripleisatuple(s,p,o)∈(I∪B)×I×(I∪B∪L)andanRDFgraphisasetofRDFtriples.RDF?extendssuchtriplesbypermittingtheembeddingofagiventripleinthesubjectorobjectpositionofanothertriple.Tripleswhosesubjectorobjectisanembeddedtriplerepresentsomeformofmetadata.Anembeddedtriplemayitselfbeametadatatripleand,thus,mayalsocontainembeddedtriples;andsoforth.Thefollowingdeﬁnitioncapturesthisnotion.
Deﬁnition1.LetT?bean(inﬁnite)setoftuplesthatisdeﬁnedrecursivelyasfollows:1.T?includesallRDFtriples,i.e.,T?⊇(I∪B)×I×(I∪B∪L);and2.ift∈T?andt0∈T?,then(t,p,o)∈T?,(s,p,t)∈T?and(t,p,t0)∈T?foralls∈(I∪B),p∈I,ando∈(I∪B∪L).Anytuple(s,p,o)∈T?isanRDF?triple.AsetofRDF?triplesiscalledanRDF?graph.
Hereafter,foranyRDF?triplet∈T?,Elmts+(t)denotesthesetofallRDFtermsandallRDF?triplesmentionedint;i.e.,ift=(s,p,o),thenElmts+(t)={s,p,o}∪x0∈Elmts+(x)
x∈{s,o}∩T? .AnRDF?tripletwithElmts+(t)∩T?6=∅iscalledametadatatriple(notethatanyotherRDF?tripleisanordinaryRDFtriple).OverloadingfunctionElmts+,foranyRDF?graphG?,Elmts+(G?)=St∈G?Elmts+(t).Further-more,Emb+(G?)denotesthesetofallRDF?triplesthatare(recursively)embeddedinRDF?triplesofRDF?graphG?;i.e.,Emb+(G?)=Elmts+(G?)∩T?.Example5.ThedatarepresentedinExample3canbeparsedintothefollowingRDF?graph.G?
ex=(:bob,foaf:name,"Bob"),((:bob,foaf:age,23),dct:creator,http://example.com/crawlers#c1),((:bob,foaf:age,23),dct:source,http://example.net/homepage-listing.html) Hence,thisRDF?graphconsistsofthreeRDF?triples,anditssetofembeddedRDF?triplescontainsasingletriple,thatis,Emb+(G?
ex)=(:bob,foaf:age,23) .4
3.2RDF?SemanticsTosupportamodel-theoreticinterpretationofRDF?graphsintermsofthestandardRDFseman-tics[HPS14]thissectiondeﬁnesatransformationfromRDF?graphstoordinaryRDFgraphs.ThistransformationmayalsobeusedtoenableordinaryRDFdatamanagementsystems(thatdonotsupportRDF?)toprocessdatathatisrepresentedasanRDF?graph.Thetransformationisbasedonthefollowingthreefunctions.First,thetransformationusesafunctionthatassociateseveryembeddedRDF?triplet∈Emb+(G?)inanRDF?graphG?withafreshanduniqueblanknode.Hence,thisfunctioniscalledabnodeassignmentfunction.
Deﬁnition2.AbnodeassignmentfunctionidforanRDF?graphG?isabijectivefunctionid:Emb+(G?)→BsuchthatB⊆Bisasetofblanknodesthathasthefollowingtwoproperties:(i)|B|=
Emb+(G?)
and(ii)B∩Elmts+(G?)=∅.
Second,thetransformationusesareiﬁcationfunctionthatassociateseveryembeddedRDF?tripleinanRDF?graphwithacorrespondingsetoffourreiﬁcationtriples.
Deﬁnition3.LetG?beanRDF?graphandletidbeabnodeassignmentfunctionforG?.Theid-specificreiﬁcationfunctionforG?isafunctionreifid:Emb+(G?)→2T?that,forevery(embedded)RDF?triplet∈Emb+(G?),isdeﬁnedasfollows:reifid t=(id∗(t),rdf:type,rdf:Statement),(id∗(t),rdf:subject,id∗(s)),(id∗(t),rdf:predicate,id∗(p)),(id∗(t),rdf:object,id∗(o)) ,whereid∗(t)=id(t)forallt∈Emb+(G?)andid∗(x)=xforallx/∈Emb+(G?).
Thethirdfunctionforthetransformationunfolds(potentiallynested)RDF?triplesrecursively.
Deﬁnition4.LetG?beanRDF?graphandletidbeabnodeassignmentfunctionforG?.Theid-specificunfoldfunctionforG?isafunctionrdfid: G?∪Emb+(G?)→2Tthat,foreveryRDF?triplet∈ G?∪Emb+(G?)witht=(s,p,o),isdeﬁnedasfollows:rdfid(t)=
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
(id(s),p,o) ∪reifid(s)∪rdfid(s)ifs∈T?ando/∈T?,(s,p,id(o)) ∪reifid(o)∪rdfid(o)ifs/∈T?ando∈T?,(id(s),p,id(o)) ∪reifid(s)∪rdfid(s)ifs∈T?ando∈T?,∪reifid(o)∪rdfid(o)(s,p,o) else.
Giventhesethreefunctions,thetransformationitselfisdeﬁnedasfollows.
Deﬁnition5.LetG?beanRDF?graphandletidbeabnodeassignmentfunctionforG?.Theid-specificunfoldedRDFgraphofG?,denotedbyrdfid(G?),isanRDFgraphthatisdeﬁnedasfollows:rdfid(G?)=[t∈G?rdfid(t).
5
Remark1.Duetothedeﬁnitionoftheunfoldfunction,thetransformationasgiveninDeﬁnition5entailsanyRDFtriplet∈Emb+(G?)thatisembeddedinsomemetadatatripleinanRDF?graphG?.Hence,thegiventransformationcapturestheusecaseofRDFreiﬁcationinwhichRDFgraphsthatcontainareiﬁcationofanRDFtripletalsocontaintitself.Example6.AnunfoldedRDFgraphoftheRDF?graphinExample5isgivenasfollows:rdfidex(G?
ex)=(:bob,foaf:name,"Bob"),(:bob,foaf:age,23),(b,rdf:type,rdf:Statement),(b,rdf:subject,:bob),(b,rdf:predicate,foaf:age),(b,rdf:object,23),(b,dct:creator,http://example.com/crawlers#c1),(b,dct:source,http://example.net/homepage-listing.html) Notethatthisexampleusesabnodeassignmentfunctionidexthatassociatesthe(embedded)RDF?triple(:bob,foaf:age,23)withblanknodeb∈B.3.3Turtle?–AnRDF?ExtensionofTurtleExample3outlinesapossibleextensionoftheTurtlesyntaxtowriteanRDF?graph.Thissectiondeﬁnesthisextension,calledTurtle?.Turtle?extendstheTurtlegrammar(asgivenin[PC14,Section6.5])withthefollowingthreeadditionalproductions.tripleX::=’<<’subjectXpredicateobjectX’>>’subjectX::=iri|BlankNode|tripleXobjectX::=iri|BlankNode|literal|tripleXAnystringthatmatchesproductiontripleXistobemappedtoanRDF?triple(s,p,o)suchthat(i)sistheRDFtermorthe(embedded)RDF?triplethatcanbeobtainedbyparsingthesubstringthatmatchessubjectX,(ii)pistheRDFtermobtainedbyparsingthesubstringthatmatchespredicate,and(iii)oistheRDFtermorthe(embedded)RDF?tripleobtainedbyparsingthesubstringthatmatchesobjectX.Inadditiontoaddingthesethreeproductionstothegrammar,Turtle?extendstheproductionslabeled[10]and[12]inthestandardTurtlegrammarasfollows(theextensiontotheproductionsaregiveninboldfont).subject::=iri|BlankNode|collection|||tripleXobject::=iri|BlankNode|collection|blankNodePropertyList|literal|||tripleXATurtle?parserisaTurtleparserthatisextendedtotakeintoaccounttheproductionsdeﬁnedinthissection.Hence,suchaparserconstructsasetofRDF?triples(i.e.,anRDF?graph)thatcanbeprocessedbyanRDF?-awaresystem.NotethatordinaryRDFdatamanagementsystems(thatdonotsupportRDF?)mayeasilybeenabledtoreadaTurtle?documentandprocessthedata;theyonlyneedtouseaTurtle?parser6
equippedwithatransformationcomponentthattransformstheRDF?graphgiveninthedocumenttoanunfoldedRDFgraphasdeﬁnedinSection3.2.
4SPARQL?–AMetadataExtensionofSPARQLThissectionintroducesSPARQL?,whichisanRDF?-awareextensionoftheRDFquerylanguageSPARQL;i.e.,SPARQL?canbeusedtoqueryRDF?graphs.Tofullybeneﬁtfromtheextendeddatamodel,SPARQL?addsnewfeaturesthatenableuserstodirectlyaccessmetadatatriplesinqueries.Inparticular,SPARQL?introducesthepossibilitytobindRDF?triplestoqueryvariables;suchavariablemaythenbeusedinatriplepatterninordertoaskformatchingmetadatatriples.Furthermore,asashortcut,(recursivelynested)triplepatternsmaybeembeddeddirectlyintriplepatterns(asdemonstratedinExample4).Inthefollowing,Section4.1introducesbasicterminologyandconcepts.Section4.2deﬁnesSPARQL?basedonP´erezetal.’salgebraicsyntaxofSPARQL[PAG09].Thereafter,Section5providesthecorrespondingextensionoftheW3CspeciﬁcationofSPARQL[HSP13].4.1BasicTerminologyandConceptsΩThebasicconceptsfordeﬁningSPARQLqueriesandtheirsemanticsaretriplepatternsandsolutionmappings.Atriplepatternisatupletp∈ V∪I∪B∪L× V∪I× V∪I∪B∪LwhereVisasetofqueryvariablesthatisdisjointfromI,B,andL,respectively.Asolutionmappingisapartialmappingµ:V→ I∪B∪L.SPARQL?extendstheseconceptsbyintroducinganotionoftriple?patternsandsolution?mappings.
Deﬁnition6.LetTP?bean(inﬁnite)setoftuplesthatisdeﬁnedrecursivelyasfollows:1.TP?includesalltriplepattern,i.e.,TP?⊇ V∪I∪B∪L× V∪I× V∪I∪B∪L;and2.iftp∈TP?andtp0∈TP?,then(tp,p,o)∈TP?,(s,p,tp)∈TP?and(tp,p,tp0)∈TP?foralls∈(V∪I∪B∪L),p∈(V∪I),ando∈(V∪I∪B∪L).Anytuple(s,p,o)∈TP?isatriple?pattern.
Deﬁnition7.Asolution?mappingisapartialmappingη:V→ T?∪I∪B∪L.
Notethat,incontrasttostandardsolutionmappingsthatbindvariablesonlytoanIRI,ablanknode,oraliteral,asolution?mappingmaybindavariablealsotoanRDF?triple.Thefollowingthreedeﬁnitionsadaptthestandardnotionsofcompatibilityofsolutionmappings,mergingofsolutionmappings,andapplicationofsolutionmappingstosolution?mappings.
Deﬁnition8.Twosolution?mappingsηandη0arecompatible,denotedbyη∼η0,if,foreveryvariable?v∈ dom(η)∩dom(η0),η(?v)=η0(?v).
7
Deﬁnition9.Letηandη0betwosolution?mappingsthatarecompatible.Themergeofηandη0,denotedbyη∪η0,isasolution?mappingη00thathasthefollowingthreeproperties:1.dom(η00)=dom(η)∪dom(η0),2.η00(?v)=η(?v)forall?v∈dom(η),and3.η00(?v)=η0(?v)forall?v∈dom(η0)\dom(η).
Deﬁnition10.Theapplicationofasolution?mappingηtoatriple?patterntp,denotedbyη[tp],isthetriple?patternthatcanbeobtainedbyreplacingallvariablesintpaccordingtoη(unboundvariablesmustnotbereplaced).
4.2(Algebraic)SyntaxandSemanticsofSPARQL?ThissectiondeﬁnesthesemanticsofthecorefragmentofSPARQL?,whichisrepresentedbasedonanalgebraicsyntaxthatextendsthealgebraicSPARQLsyntaxintroducedbyP´erezetal.[PAG09].
Deﬁnition11.ASPARQL?expressionisdeﬁnedrecursivelyasfollows:1.Anyﬁnitesetoftriple?patternsisaSPARQL?expression,whichiscalledaBGP?.2.Iftpisatriple?patternand?visavariable,then(tpAS?v)isaSPARQL?expression.3.IfP1andP2areSPARQL?expressionsandRisaﬁltercondition6,then(P1ANDP2),(P1UNIONP2),(P1OPTP2),and(P1FILTERR)areSPARQL?expressions.
Example7.TheﬁrstquerypatternofExample4canberepresentedasaBGP?Pex=(?bob,foaf:name,"Bob"),((?bob,foaf:age,?age),dct:source,?src) ,whichconsistsoftwotriple?patterns.ThesecondquerypatternofExample4canberepresentedasasemanticallyequivalentSPARQL?expressionPex2thathasthefollowingform: (?bob,foaf:age,?age)AS?tAND(?bob,foaf:name,"Bob"),(?t,dct:source,?src) .ThebasisfordeﬁningthesemanticsofSPARQL?isanalgebraovermultisetsofsolution?map-pingsthatresemblesthestandardSPARQLalgebra(whichisdeﬁnedovermultisetsofordinarysolutionmappings).Formally,amultisetsolution?mappingsisapairM=(Ω,card)whereΩistheunderlyingset(ofsolution?mappings)andcardisthecorrespondingcardinalityfunction;i.e.,card:Ω→{1,2,...}.Then,theSPARQL?-speciﬁcalgebraoperatorsaredeﬁnedasfollows.
6ForadeﬁnitionofthesyntaxofﬁlterconditionsrefertoP´erezetal.’swork[PAG09].8
Deﬁnition12.LetM1=(Ω1,card1)andM2=(Ω2,card2)bemultisetsofsolution?mappings.•ThejoinofM1andM2,denotedbyM1onM2,isamultisetofsolution?mappings(Ω,card)suchthatΩ=η1∪η2
η1∈Ω1andη2∈Ω2andη1∼η2 and,foreveryη∈Ω,card(η)=X(η1,η2)∈Ωηcard1(η1)·card2(η2),whereΩη=(η1,η2)∈Ω1×Ω2
η1∪η2=η .•The(multiset)unionofM1andM2,denotedbyM1dM2,isamultisetofsolution?mappings(Ω,card)suchthatΩ=Ω1∪Ω2and,foreveryη∈Ω,card(η)=



card1(η)+card2(η)ifη∈(Ω1∪Ω2),card1(η)ifη∈(Ω1\Ω2),card2(η)else.•The(multiset)diﬀerenceofM1andM2,denotedbyM1\\M2,isamultisetofsolution?mappings(Ω,card)suchthatΩ=η∈Ω1
η6∼η0forallη0∈Ω2 andcard(η)=card1(η)forallη∈Ω.•TheleftouterjoinofM1andM2,denotedbyM1
onM2,isamultisetofsolution?mappingsthatisdeﬁnedby:M1
onM2= M1onM2d M1\\M2.•GivenaﬁlterconditionR,theR-speciﬁcselectionofM1,denotedbyσR(M1),isamultisetofsolution?mappings(Ω,card)suchthatΩ=η∈Ω1
ηsatisﬁesR andcard(η)=card1(η)forallη∈Ω,whereasolution?mappingηsatisﬁesﬁlterconditionRifanyofthefollowingholds:1.Risbound(?v)where?v∈Vand?v∈dom(η);2.Ris?v=cwhere?v∈Vandc∈ I∪L,?v∈dom(η),andη(?v)=c;3.Ris?x=?ywhere?x,?y∈V,?x∈dom(η),?y∈dom(η),andη(?x)=η(?y);4.Ris(¬R0)whereR0isaﬁlterconditionandηdoesnotsatisfyR0;5.Ris(R1∨R2)whereR1andR2areﬁlterconditionsandηsatisﬁesR1orR2;or6.Ris(R1∧R2)whereR1andR2areﬁlterconditionsandηsatisﬁesbothR1andR2.
Giventhesealgebraoperators,thesemanticsofanySPARQL?expressionisdeﬁnedbythefollowingevaluationfunction.9
Deﬁnition13.LetPbeaSPARQL?expressionandletG?beanRDF?graph.TheevaluationofPoverG?,denotedby[[P]]G?,isamultisetofsolution?mappings(Ω,card)thatisdeﬁnedrecursivelyasfollows:1.IfPisaBGP?,thenΩ=η
ηρ[P]⊆(Emb+(G?)∪G?)forsomeP-bnodesmappingρ and,foreveryη∈Ω,card(η)=

ρ
ρisaP-bnodesmappingsuchthatηρ[P]⊆(Emb+(G?)∪G?)

,whereaP-bnodesmappingisamappingρ:bn(P)→(T?∪I∪B∪L)andηρ[P]=η[tp]
tpisatriple?patternobtainedbyreplacingallblanknodesinsometriple?patterntp0∈Paccordingtoρ .2.IfPis(tpAS?v),thenΩ=η
∃η0∈Ω0:η0∼ηanddom(η)= dom(η0)∪{?v}andη(?v)=η0[tp] and,foreveryη∈Ω,card(η)=Xη0∈Ω0
ηcard0(η0),where Ω0,card0=[[{tp}]]G?andΩ0
η={η0∈Ω0|η0∼η}forallη∈Ω.3.IfPis(P1ANDP2),then(Ω,card)=[[P1]]G?on[[P2]]G?.4.IfPis(P1UNIONP2),then(Ω,card)=[[P1]]G?d[[P2]]G?.5.IfPis(P1OPTP2),then(Ω,card)=[[P1]]G?
on[[P2]]G?.6.IfPis(P0FILTERR),then(Ω,card)=σR [[P0]]G?.
Example8.TheevaluationofSPARQL?expressionPex(cf.Example7)overRDF?graphG?
ex(cf.Ex-ample5)consistsofasinglesolution?mappingη1,whichhasthefollowingproperties:1.dom(η1)={?bob,?age,?src},2.η1(?bob)=:bob,3.η1(?age)=23,and4.η1(?src)=http://example.net/homepage-listing.html.FortheotherexpressionfromExample7weobtainthesameresult:[[Pex2]]G?
ex=(Ωex2,cardex2)whereΩex2={η1}andcardex2(η1)=1.10
5ExtensionoftheW3CSpeciﬁcationofSPARQLΩAfterdeﬁningSPARQL?basedonanalgebraicsyntax,theremainderofthisdocumentdeﬁnesacorrespondingextensionoftheformalizationofSPARQL1.1thatisgivenbytheW3Cspeci-ﬁcation[HSP13].Thisextensionassumesthatanymentionof“RDFtriple”inthespeciﬁcationisunderstoodasanRDF?triple;similarly,“RDFgraph”,“triplepattern”,“basicgraphpattern”(or“basicgraphpattern”,and“solutionmapping”areunderstoodasRDF?graph,triple?pattern,BGP?,andsolution?mapping,respectively.Furthermore,theunderstandingofa“propertypathpattern”includesthepossibilitytouseatriple?patternassubjectorobjectofsuchapattern.Section5.1introducesthegrammarofSPARQL?asanextensionofthegrammarofSPARQL1.1.Section5.2speciﬁeshowtosupporttheextendedgrammarduringtheconversionofquerystringsintoalgebraexpressions;thisspeciﬁcationincludestheintroductionofanewalgebrasymbolwhichcorrespondstoSPARQL?expressionsoftheform(tpAS?v).Section5.3deﬁnestheevaluationsemanticsfortheresultingalgebraexpressions.
5.1GrammarΩThissectionspeciﬁestheSPARQL?grammarasanextensionofthestandardSPARQL1.1gram-mar[HSP13].Elementsofthegrammarthatarenotspeciﬁedexplicitlyinthissectionaredeﬁnedasgivenin[HSP13,Section19.8].Anembeddedtriplepatternisanewsyntaxelementthatconformstothefollowing,newgram-marrules:EmbTP::=’<<’VarOrBlankNodeOrIriOrLitOrEmbTPVerbVarOrBlankNodeOrIriOrLitOrEmbTP’>>’VarOrBlankNodeOrIriOrLitOrEmbTP::=Var|BlankNode|iri|RDFLiteral|NumericLiteral|BooleanLiteral|EmbTPAsthegivengrammarrulesindicate,anembeddedtriplepatternmaycontainotherembeddedtriplepatterns.Embeddedtriplepatternsmaybeusedinaqueryinthefollowingtwoways:(i)theyarepartofaBINDclause(whichcorrespondstoSPARQL?expressionsoftheform(tpAS?v)),or(ii)theyareembeddedinatriple?patternorinapropertypathpattern.Bothoftheseoptionsarespeciﬁedanddiscussedinthefollowing.SPARQLintroducestheBINDclauseforassigningthevalueofevaluatingagivenexpressiontoavariable.ToenabletheuseofembeddedtriplepatternsinBINDclauses(insteadofanexpression),thefollowingtwoextensionstothegrammararenecessary.First,anewgrammarruleisadded:ExpressionOrEmbTP::=Expression|||EmbTPSecond,theoriginalgrammarrule[60]isredeﬁnedasfollows:7Bind::=’BIND’’(’ExpressionOrEmbTP’AS’Var’)’AnembeddedtriplepatternmaynotonlybeusedinaBINDclausebutitmayalsobeembeddedinapropertypathpatternorinatriple?pattern.Moreprecisely,intheextendedsyntaxthesubject
7Theadjustedpartinwhicharedeﬁnedgrammarrulediﬀersfromtheoriginalrulein[HSP13,Section19.8]isgiveninboldfont.11
orobjectofapropertypathpatterncanbeanembeddedtriplepattern(insteadofanRDFtermoravariable).Similarly,triple?patternsmayhaveanembeddedtriplepatterninthesubjectpositionorintheobjectposition(cf.Deﬁnition6).Tothisend,thegrammarisextendedwithanewrule:VarOrTermOrEmbTP::=Var|GraphTerm|EmbTPMoreover,theoriginalgrammarrules[75],[80],[81],and[105]areredeﬁned:TriplesSameSubject::=VarOrTermOrEmbTPPropertyListNotEmpty|TriplesNodePropertyListObject::=GraphNode|||EmbTPTriplesSameSubjectPath::=VarOrTermOrEmbTPPropertyListPathNotEmpty|TriplesNodePathPropertyListPathGraphNodePath::=VarOrTermOrEmbTP|TriplesNodePath5.2TranslationtotheAlgebraΩBasedontheSPARQLgrammartheSPARQLspeciﬁcation“deﬁnestheprocessofconvertinggraphpatternsandsolutionmodiﬁersinaSPARQLquerystringintoaSPARQLalgebraexpres-sion”[HSP13,Section18.2].Thisprocessmustbeadjustedtoconsidertheextendedgrammarintroducedintheprevioussection.Inthefollowing,anystepoftheconversionprocessthatre-quiresadjustmentisdiscussed.
5.2.1VariableScopeΩAsabasisofthetranslation,theSPARQLspeciﬁcationintroducesanotionofin-scopevari-ables[HSP13,Section18.2.1].TocoverthenewsyntaxelementsintroducedinSection5.1thisnotionmustbeextendedasfollows.•Avariableisin-scopeofaBGP?BifthevariableoccursinB,whichincludesanoccurrenceinanyembeddedtriplepatterninBGP?(independentofthelevelofnesting).•Avariableisin-scopeofapropertypathpatternifthevariableoccursinthatpattern,whichincludesanoccurrenceinanyembeddedtriplepatterninthepattern(independentofthelevelofnesting).•Avariableisin-scopeofaBINDclauseoftheformBIND(TASv)(whereTisanembeddedtriplepattern)ifthevariableisvariablevorthevariableoccursintheembeddedtriplepatternT.AsforstandardBINDclauseswithexpressions,variablevmust“not[be]in-scopefromthepreceedingelementsinthegroupgraphpatterninwhich[theBINDclause]isused”[HSP13,Section18.2.1].5.2.2ExpandSyntaxFormsΩThetranslationprocessstartswithexpanding“abbreviationsforIRIsandtriplepatterns”[HSP13,Section18.2.2.1].Thisstepmustbeextendedintwoways:1.Abbreviationsfortriplepatternswithembeddedtriplepatternsmustbeexpandedasifeachembeddedtriplepatternwasavariable(oranRDFterm).Forinstance,thefollowingsyntaxexpression12
<<?cardfs:Class>>dct:source?src;prov:wasDerivedFrom<<?caowl:Class>>.mustbeexpandedto
<<?cardfs:Class>>dct:source?src.
<<?cardfs:Class>>prov:wasDerivedFrom<<?caowl:Class>>.2.AbbreviationsforIRIsinallembeddedtriplepatternsmustbeexpanded.Forinstance,theembeddedtriplepattern
<<?cardfs:Class>>Ωmustbeexpandedto
<<?c<http://www.w3.org/1999/02/22-rdf-syntax-ns#type><http://www.w3.org/2000/01/rdf-schema#Class>>>5.2.3TranslatePropertyPathPatternsΩThetranslationofpropertypathpatterns(cf.[HSP13,Section18.2.2.4])hastobeadjustedbecausetheextendedgrammarallowsforpropertypathpatternswhosesubjectorobjectisanembeddedtriplepattern(cf.Section5.1).ThetranslationasspeciﬁedintheW3Cspeciﬁcationdistinguishesfourcases.Theﬁrstthreeofthesecasesdonotrequireadjustmentbecausetheyaretakencareofeitherbyrecursionorbytheadjustedtranslationofbasicgraphpatterns(asdeﬁnedinSection5.2.4below).However,thefourthcasemustbeadjustedasfollows.LetXPYbeastringthatcorrespondstothefourthcasein[HSP13,Section18.2.2.4].GiventhegrammarintroducedinSection5.1,XandYmaybeanRDFterm,avariable,oranembeddedtriplepattern,respectively(andPisapropertypathexpression).ThestringXPYistranslatedtothealgebraexpressionPath(X’,P,Y’)whereX’andY’aretheresultofcallingafunctionnamedLiftforXandY,respectively.ForsomeinputstringZ(suchasXorY)thatcanbeanRDFterm,avariable,oranembeddedtriplepattern,thefunctionLiftisdeﬁnedasfollows:IfZisanembeddedtriplepattern<<S,P,O>>Returntriple*pattern(Lift(S),P,Lift(O))ElseReturnZEnd
5.2.4TranslateBasicGraphPatternsΩAftertranslatingpropertypathpatterns,thetranslationprocesscollects“anyadjacenttriplepat-terns[...]toformabasicgraphpattern”[HSP13,Section18.2.2.5].Thisstephastobeadjustedbecausetriplepatternsintheextendedsyntaxmayhaveanembeddedtriplepatternintheirsub-jectpositionorintheirobjectposition(orinboth).ToensurethateveryresultofthisstepisaBGP?,beforeaddingatriplepatterntoitscorrespondingcollection,itssubjectandobjectmustbereplacedbytheresultofcallingfunctionLift(cf.Section5.2.3)forthesubjectandtheobject,respectively.13
5.2.5TranslateBINDClauseswithanEmbeddedTriplePatternΩTheextendedgrammarinSection5.1allowsforBINDclauseswithanembeddedtriplepattern.ThetranslationofsuchaBINDclausetoaSPARQLalgebraexpressionrequiresanewalgebrasymbol:•TR(triple?pattern,variable)NotethatthissymbolcorrespondstoSPARQL?expressionsoftheform(tpAS?v)(cf.Deﬁnition11).Then,anystringoftheformBIND(TASv)withTbeinganembeddedtriplepattern(i.e.,notastandardBINDexpression)istranslatedtothealgebraexpressionTR(T’,v)whereT’istheresultofcallingtheaforementionedfunctionLiftforT.Notice,thetranslationofBINDclauseswithanembeddedtriplepatternasdeﬁnedinthissectionisusedduringthetranslationofgroupgraphpatternsthatisspeciﬁedin[HSP13,Sec-tion18.2.2.6].ThecaseofBINDclauseswithanembeddedtriplepatterniscoveredinthistrans-lationofgroupgraphpatternsbythelast,“catchallother”IFstatement(i.e.,theIFstatementwiththeconditionEisanyotherform)andnotbytheIFstatementforBINDclauseswithanexpression.
5.3EvaluationSemanticsΩTheSPARQLspeciﬁcationdeﬁnesafunction“eval(D(G),algebraexpression)astheevaluationofanalgebraexpressionwithrespecttoadatasetDhavingactivegraphG”[HSP13,Section18.6].RecallthattheactivegraphGinthecontextofSPARQL?isanRDF?graph,andsoisanyothergraphindatasetD.Thedeﬁnitionoffunctionevalisrecursive;thetwobasecasesofthisdeﬁnitionforSPARQL?aregivenasfollows:•ForanyBGP?B,eval D(G),B=[[B]]G(where[[B]]GistheevaluationofBoverRDF?graphGasperDeﬁnition13).•ForanyalgebraexpressionEoftheformTR(tp,?v)wheretpisatriple?patternand?visavariable(asintroducedinSection5.2.5),eval D(G),E=[[(tpAS?v)]]G(where[[(tpAS?v)]]GistheevaluationofSPARQL?expression(tpAS?v)overRDF?graphGasperDeﬁnition13).Foranyotheralgebraexpression,theSPARQLspeciﬁcationdeﬁnesalgebraoperators.Thesedeﬁnitionscanbeextendednaturallytooperateovermultisetsofsolution?mappings(insteadofordinarysolutionmappings).Giventhisextension,therecursivestepsofthedeﬁnitionoffunctionevalforSPARQL?arethesameasintheSPARQLspeciﬁcation.ReferencesΩ[ACA+12]GrigorisAntoniou,OscarCorcho,KarlAberer,ElenaSimperl,andRudiStuder.Se-manticDataManagement(DagstuhlSeminar12171).DagstuhlReports,2(4),2012.[CWL+14]RichardCyganiak,DavidWood,MarkusLanthaler,GrahamKlyne,JeremyJ.Carroll,andBrianMcBride.RDF1.1ConceptsandAbstractSyntax.W3CRecommendation,Onlineat http://www.w3.org/TR/rdf11-concepts/
,February2014.[HPS14]PatrickJ.HayesandPeterF.Patel-Schneider.RDF1.1Semantics.W3CRecommen-dation,Onlineat http://www.w3.org/TR/rdf11-mt/
,February2014.14
[HSP13]SteveHarris,AndySeaborne,andEricPrud’hommeaux.SPARQL1.1QueryLanguage.W3CRecommendation,Onlineat http://www.w3.org/TR/sparql11-query/
,March2013.[PAG09]JorgeP´erez,MarceloArenas,andClaudioGutierrez.SemanticsandComplexityofSPARQL.ACMTransactionsonDatabaseSystems,34(3),2009.[PC14]EricPrud’hommeauxandGavinCarothers.RDF1.1Turtle.W3CRecommendation,Onlineat http://www.w3.org/TR/turtle/
,February2014.[TPC14]BryanThompson,MikePersonick,andMartynCutcher.ThebigdataRDFGraphDatabase.InAndreasHarth,KatjaHose,andRalfSchenkel,editors,LinkedDataManagement.CRCPress,2014.15

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

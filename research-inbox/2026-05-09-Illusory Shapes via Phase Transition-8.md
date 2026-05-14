---
source: "https://arxiv.org/abs/1406.1265v1"
title: "Illusory Shapes via Phase Transition"
author: "Yoon Mo Jung, Jianhong Jackie Shen"
year: "2014"
publication: "arXiv preprint / math.OC"
download: "https://arxiv.org/pdf/1406.1265v1"
pdf: "https://arxiv.org/pdf/1406.1265v1"
captured_at: "2026-05-09T13:02:49Z"
updated_at: "2026-05-09T13:02:49Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代思想"
status: raw
---

# Illusory Shapes via Phase Transition

- 著者: Yoon Mo Jung, Jianhong Jackie Shen
- 年: 2014
- 掲載情報: arXiv preprint / math.OC
- 情報源: [arxiv](https://arxiv.org/abs/1406.1265v1)
- ダウンロード: https://arxiv.org/pdf/1406.1265v1
- PDF: https://arxiv.org/pdf/1406.1265v1

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

We propose a new variational illusory shape (VIS) model via phase fields and phase transitions. It is inspired by the first-order variational illusory contour (VIC) model proposed by Jung and Shen [{\em J. Visual Comm. Image Repres.}, {\bf 19}:42-55, 2008]. Under the new VIS model, illusory shapes are represented by phase values close to 1 while the rest by values close to 0. The 0-1 transition is achieved by an elliptic energy with a double-well potential, as in the theory of $Γ$-convergence. The VIS model is non-convex, with the zero field as its trivial global optimum. To seek visually meaningful local optima that can induce illusory shapes, an iterative algorithm is designed and its convergence behavior is closely studied. Several generic numerical examples confirm the versatility of the model and the algorithm.

## PDF Text

arXiv:1406.1265v1 [math.OC] 5 Jun 2014
NonamemanuscriptNo.
(willbeinsertedbytheeditor)
IllusoryShapesviaPhaseTransition
YoonMoJung·JianhongJackieShenDedicatedtoGilStrangforHis80thBirthday
AbstractWeproposeanewvariationalillusoryshape(VIS)modelviaphaseﬁeldsandphasetransitions.Itisinspiredbytheﬁrst-ordervariationalillusorycontour(VIC)modelproposedbyJungandShen[J.VisualComm.ImageRepres.,19:42-55,2008].UnderthenewVISmodel,illusoryshapesarerepresentedbyphasevaluescloseto1whiletherestbyvaluescloseto0.The0-1transitionisachievedbyanellipticenergywithadouble-wellpotential,asinthetheoryofΓ-convergence.TheVISmodelisnon-convex,withthezeroﬁeldasitstrivialglobaloptimum.Toseekvisuallymeaningfullocaloptimathatcaninduceillusoryshapes,aniterativealgorithmisdesignedanditsconvergencebehavioriscloselystudied.Severalgenericnumericalexamplesconﬁrmtheversatilityofthemodelandthealgorithm.
KeywordsIllusoryShapes·PhaseTransition·NullHypothesis·Convergence1Introduction
Theintriguingphenomenonofillusorycontoursandshapes(ICS)hasbeenactivelypursuedincontemporarybrainscience,neurophysiology,neuroimaging,neuralnet-works,computationalneuroscience,andpsychophysics[35,24,16,15,25,40,23].WeﬁrstquotefromtherecentreviewpaperbyMurrayandHerrmann[29](2013)tohighlightitsimportance:“Theprogressivedevelopmentofthesemodels(forICS)andthedatasupport-ingtheminmanyregardshighlightwhatmaybeconsideredageneraltendency
YoonMoJung
DepartmentofComputationalScienceandEngineering,YonseiUniversity,KoreaE-mail:ymjung@yonsei.ac.kr
JackieShen(correspondingauthor)
DepartmentofIndustrialandSystemsEngineering,UniversityofIllinois,IL61801,USATel.:(646)725-0065
Fax:(217)244-5705
E-mail:jhshen@illinois.edu
2YoonMoJung,JianhongJackieSheninneuroscienceoverthepast30orsoyears,namely,progressionfrom...individualneurons...tothediscretelocalizationofbrainfunctions...”Thatis,thepersistentscientiﬁcinterestinICShasbeenrepresentativeinmodernneuroscience,andhasprofoundlyenhancedthesystemicapproachtostudyingcomplexvisiontasks.ICSisfundamentallyasystem-levelvisionphenomenon,ascontrasttosimplertasks(e.g.,detectionofedgelets)byindividualsimpleorcomplexneurons.Kanizsa’sTriangleillustratedinFig.1isaclassicexampleforICS.Thewhitetriangleinthemiddlepopsoutdespitethatthereexistnomodalbordersagainstthewhitebackground.Neuronalmeasurementsinthevisualcorticesofcatsormonkeysrevealeddirectﬁringresponsesalongsuchmodaledges[38,37].Thisistheverynatureofillusion,namely,theremarkableabilityinreconstructionandre-organizationfromgivenvisualinformationwithinternalstructures.Intermsofsystemsandcontrol,suchICSphenomenahaveprovidedtheidealclassofinputsignalsthathelpidentifythecomplexstructuresandfunctionali-tiesoftheprimatevisionsystem,including,e.g.,thebottom-upandtop-downschema[23,39,18].Treatingthemulti-tiervisualcorticesandpathwaysassys-temsisfundamental.Inparticular,itimpliesthatmodelshavetobebasedonsystemicestimationsanddecisions,insteadofisolatedneuronalﬁringpatterns.Mathematically,thisnaturallyinvitesBayesiandecisionmakingovernetworks,orsystem-wideoptimization[22].Deterministicvariationaloptimizationcouldbeconsideredasthelowtemper-aturelimitoftheBayesiandecisionframework.Asinstatisticalmechanics[11],underthelow-temperaturelimit,geometricregularitiesemergefromstatisticalgaseouspatterns(correspondingtothetexturalpatternsinimagingandvision)[9].ThevariationalPDEmethodologyisparticularlypowerfulinhandlinggeometricregularitiessuchasthedistanceorthecurvature[4,9,5,36].ForICS,therehavebeenseveralvariationalPDEmodelsintheliterature[31,10,41,19],aspartiallyreviewedinourearlierwork[19].In[19],basedonthedecompositionintorealandimaginarycomponents,wehaveproposedaﬁrst-ordervariationalillusorycontour(VIC)model,whichisimplementedbythesupervisedlevel-setmethodofOsherandSethian[30].Amongallthevariational-PDEmodelsforICSthattheauthorsareawareof,thisVICmodelhasthelowestcomplexityandthusallowsdetailedanalysisoftheillusoryshapes(aslocaloptima).ThiswillbebrieﬂyrecappedinSection3.Thecurrentworkhasbeencloselyinspiredby[19]andseeksavariationalillusoryshape(VIS)modelbasedonphasetransitions.Anillusoryshapeisideallyrepresentedbythephasevalue1.0andtherestby0.0.ThemachineryofΓ-convergenceandphasetransitions[2,3]allowsustoconstructaphase-ﬁeldenergythatiscloselyrelatedtotheVICmodelintroducedin[19].ThenewVISmodelismoreself-containedfrommodelingtocomputing.MoredetailsareexplainedinSections2and4.Likemostphase-ﬁeldmodels[26,27,33,34],theproposednonlinearVISmodelisnotconvex.Thezeroﬁeldistheglobaloptimabutuninteresting.Visuallymean-ingfullocaloptimaarehencesoughtafterviaaniterativealgorithm.ThedesignandanalysisofthealgorithmwillbeelaboratedinSection5.InSection6,wesummarizetheentirealgorithmviaapseudocode,andpresentseveralgenericnumericalexamplestoillustratetheversatilityofthemodeland
IllusoryShapesviaPhaseTransition3thealgorithm.TheworkisﬁnallyconcludedinSection7,wherethelimitationsofthemodelarealsoaddressed.
2PhaseFieldRepresentationofIllusoryShapes
GivenaconﬁgurationQinavisualﬁeldΩ,theremayemergeanillusoryshapeSoutsideQ:S⊆Ω\Q.Fig.1showstheclassicexampleofKanizsa’sTriangle,ofwhichQconsistsofthreepac-mandisksandSreferstothewhitetriangleencompassedandinducedbythem.Mathematically,weassumethatΩisaboundedLipschitzconvexopendomaininR2,e.g.,aCartesianrectangleoradisk,andthatQ⊂ΩisacompactLipschitzsub-domain.Deﬁnetheidealbinaryphaseﬁeldz(x),withx=(x1,x2)∈Ω:z(x)=1,x∈S,0,x∈Ω\S.IfSisgiven,z(x)isexactlyitsindicatorfunction,andapixelx∈ΩbelongstoSifandonlyifitsphasevalueis1.
Fig.1Kanizsa’sTriangle.ThewhiteillusorytriangleSinthemiddleisinducedbytheblackconﬁgurationQconsistingofthreepac-mandisks.Inreality,neitherSnorz(x)isknownapriori.Thephaseﬁeldapproachattemptstoreconstructanontrivialphaseﬁeldz(x)’sthatcaninduceillusoryshapes.Sincebinaryﬁeldsarehardtoworkwithboththeoreticallyandcompu-tationally,oneresortstocontinuousﬁeldsthatarethemolliﬁedversionsofthebinaryones.Theycompletethe0-1phasetransitionalongatubularneighbor-hoodofthedesirablesharpboundaries.Suchcontinuousﬁeldsarethuscalledthediﬀusivephaseﬁeldsintheliterature.Forthecurrenttask,weimposethefollowingtwonaturalconditions:(a)z|Q≡0,i.e.,illusoryshapescanonlyemergeoutsideQ.(b)z|∂Ω≡0.Thisisbecause,forvariationalillusionmodelslike[19],onecanprovethatillusoryshapescanonlyemergewithintheconvexhullofQ,andthusalsowithintheinteriorofΩ.Themainchallengethenbecomeshowtoencouragethephase1.0(orcloseto1.0)toemergeforconﬁgurationsthatcaninduceillusoryshapes.
4YoonMoJung,JianhongJackieShen3VariationalContourModelofJungandShen
Inthissection,webrieﬂyreviewtheﬁrst-ordervariationalillusorycontour(VIC)modelbyJungandShen[19].Inthenextsection,thenewphase-ﬁeldmodelistobebuiltuponthisVICmodel.Forsimplicity,belowweassumeallcurvestobepiecewisesmooth.LetΓbeanysimpleclosed(Jordan)curveonΩ\Q◦.In[19],itisdecomposedtotherealandimaginarypartsby:Γre=Γ∩∂Q,Γim=Γ\∂Q.IfΓisagenuineillusorycontoursuchasthethreesidesofKanizsa’sTriangle,theimaginarypartΓimcorrespondstotheillusoryinterpolant.InFigure1,forinstance,Γrecorrespondstothethreecornerturnsalongthediskinducers,andΓimtothethreeillusorymodalsegments.ForanysuchacurveΓ,JungandShenproposedthefollowingvariationalcontourmodelbasedonthedecomposed“energy”[19]:E[Γ]=E[Γ|a,b]=aZΓreds+bZΓimds,(1)foranypairofproperlydeﬁnedweightsaandbwith0<a<b.Illusorycontoursarethendeﬁnedasthelocalminima.Comparedwithseveralothermodels[31,10,41],theVICmodel(1)issimpleandyetpowerfulenough,asdemonstratedbythegenericexamplesin[19].Ana-lytically,theauthorswereabletoestablishdetailedgeometricproperties,amongwhichthemostimportantisTheorem2.11in[19].
Theorem1(CharacterizationofaLocalMinimum)LetQbeagenericcompactconﬁgurationonanopenLipschitzdomainΩ.Assumethat∂Qispiece-wisesmoothwithﬁnitelymanycorners(butnokinks).Letθmin,θ∗min>0denoteitsminimumouterandinnerspans.SupposeagivensimpleclosedJordancurveΓ∈Csatisﬁesthefollowingstructuralconditions:(i)(ImaginaryBehavior)eachconnectedcomponentγimofΓimisastraightlinesegment,andnotwodistinctcomponentsshareacommonhinge;(ii)(JunctionBehavior)atanyjunctionpointz∈J[Γ],theturnφz<π/2,andtheidleangleφidleΩz≥π/2.LetφmaxdenotethemaximumturnonJ[Γ].Thenthereexistsacriticalratiorc=rc(θmin,φmax)<1,suchthatforanyαandβwithr=α/β<rc,ΓisalocalminimumtotheenergyE[·|α,β].Wereferto[19]forthedeﬁnitionsofspans,hinges,turns,andidleangles.Suchdetailedcharacterizationsaremuchhardertoestablishformorecomplexillusionmodels[31,10,41].In[19],illusoryshapes(asthelocalminimatoE[Γ|a,b])havebeencomputedviaasupervisedlevel-setschemeofOsherandSethian[30],asappliedtotheapproximateenergy(Eqn.(7)in[19]):Eσ[Γ|α,β]=αZΓds+βZΓg(|∇χQ,σ(s)|)ds,(2)
IllusoryShapesviaPhaseTransition5
Fig.2AnillusorycontouremergesfromacomplexconﬁgurationQofobjects,ascapturedbythesupervisedlevel-setmethodofOsherandSethian[30]fortheVICmodelinEqn.(1)proposedbyJungandShen[19].
withα=a+b,β=bforthetargetedenergyE[Γ|a,b]inEqn.(1).Hereσ1denotesasmalldiﬀusionormolliﬁcationscale,sothatthebinaryindicatorχQismolliﬁedtoasmoothapproximationχQ,σ.Thefunctiongcouldbeanypositivefunctionsatisfying:g(0+)=1,g(+∞)=0,forexample,g(p)=exp(−p2)org(p)=1/(1+p2).Itwasshownin[19]thatforanyadmissiblecontourΓ,Eσ[Γ|α,β]→E[Γ|a,b]asσ→0+.Thecriticalratiorcfora/binTheorem1nowtransferstoRc=rc/(1−rc)forα/β,accordingly.4APhaseTransitionModelforIllusoryShapes
Inthissection,wedevelopanewphasetransitionmodelforillusoryshapes,basedonthediscussionintheprecedingtwosections.WeﬁrstdeﬁnetheinducingscalarﬁeldviaG(x)=α+βg(|∇χQ,σ(x)|),foranyx∈Ω.Alternatively,onecouldemployanyoptimalphaseﬁeldzMSfromthephasetransitionapproximationtotheMumford-Shahmodel[2,3]:G(x)=α+βzMSσ(x),x∈Ω.ThebasicrequirementsforzMSσ(x)are:(i)thephasetransitionbandwidthissmall:σ1,(ii)u0(x)=χQ(x)isusedastheinputtotheMumford-Shahmodel,and(iii)zMSσ(x)∈[0,1]withzMSσ'0along∂Qand'1awayfrom∂Q.Ineitherway,onehasG(x)'α,xnearoralong∂Q,α+β,otherwise.WecallGacanyonfunctionassociatedwiththegivenQ,sincenear∂QthevaluesofGquicklydropfromα+βtoα.Forthecanyoneﬀecttobemoresalient,itisnaturaltorequirethedroppingsizeβα,whichisalsoinaccordancewiththecriticalratiorequirementα/β<Rcintheprecedingsection.LetH1
0(Ω)denotetheSobolevspaceoffunctionsonΩwithzeroboundarytraces.IthostsalldiﬀusivephaseﬁeldsdiscussedinSection2.Foranyphaseﬁeld
6YoonMoJung,JianhongJackieShenz∈H1
0(Ω),weintroducethefollowingenergy:E[z]=ZΩ
2|∇z|2+(1−z)2z2
2·Gdx+λZΩχQz2
2dx,(3)wherethepositiveweightλisinthesameorderasofG.Whenα<β=O(1)sothatG=O(1),wesimplysetλ=1inallourcomputationalexampleslateron.Thesmall“diﬀusivity”parameterdeﬁnestheintendedtransitionbandwidthbetween0.0and1.0.Thetwotermshavebeenmotivatedasfollows.Thesecondtermis:λZΩχQz2
2dx=λZQz2
2dx.Thusfor1,itactsasasoftwaytoenforcez=0onQ,whichisthecondition(a)imposedinSection2.Alsonoticethatthecondition(b)ismetautomaticallysincez∈H1
0(Ω).FortheﬁrstterminE[z],deﬁnetheBorelmeasureµforanygivenz∈H1
0(Ω)by:dµ=
2|∇z|2+(1−z)2z2
2dx.ForanysmoothJordancurveΓwitharclengthelementds,weshowsemi-heuristicallythatdµ→1
6dsas→0+,ormoregenerallyviathe1-dimensionalHausdorﬀmea-suredH1ΩΓthatdµ→1
6dH1ΩΓ(withH1ΩΓ(A)=H1(Γ∩A))foraproperlydesignedsequenceofphaseﬁelds{z}.AssumethatΓisthe0-levelsetofasmoothfunctionf(x):Γ=f−1(0),andthatfisregularalongΓinthesensethat∇f(ω)6=0forallω∈Γ.Withatmostasignﬂip,onecouldfurtherassumethat∇f(ω)pointstowardstheinsideofΓ.ThenbytheTubularNeighborhoodTheorem[6],thereexistssomeδ0>0suchthatthemap:ϕ:(ω,n)∈Γ×(−2δ0,2δ0)→x=ω+n·∇f
|∇f|∈Ω,isadiﬀeomorphismbetweenΓ×(−2δ0,2δ0)andanopentubularneighborhoodB2δ0ofΓinΩ.LetSdenotethelogisticfunction:S(t)=1
1+e−t,t∈R1.Foranyδ0,wethenconstructaspecialphaseﬁeldz∈H1
0(Ω):z(x)=





S n
,x=ϕ(ω,n)∈Bδ0,Sδ0
,x/∈Bδ0andx∈int(Γ),S−δ0
,x/∈Bδ0andx∈ext(Γ).Hereint(Γ)andext(Γ)denotetheinteriorandexteriordomainsofΓ,whicharewelldeﬁnedaccordingtotherenownedJordanCurveTheoremintopology[17].OnecanshowthatforanycontinuousfunctionφonΩ:ZΩφdµ→1
6ZΩφdH1ΩΓ=1
6ZΓφds,as→0+,(4)
IllusoryShapesviaPhaseTransition7wherethemultiplier1
6isthetotalvariationofF(z)=z2
2−z3
3fromz=0toz=1,andF(z)istheprimitiveofz(1−z)sothatF0(z)=z(1−z).Asaresult,forexample,ifΓrepresentstheillusoryKanizsaTriangleinFig.1andφ=G,asε→0,onehas6ZΩGdµ'aZΓreds+bZΓimds,a=α,b=α+β.Thisgivesaphase-ﬁeldinterpretationofthemodelbyJungandShen[19].Asimilarargumentcouldalsobefoundin[33],forexample.ButrigoroustreatmentisonlyoﬀeredbythetheoryofΓ-convergenceapproximation[7,12].Insummary,thephasetransitionmodelE[z]forz∈H1
0(Ω)proposedinEqn.(3)generalizesthecontourmodelE[Γ]ofJungandShen[19]inEqn.(1)tocontinuousphaseﬁelds.Italsoimplementsthetwoconditions(a)and(b)imposedinSection2.Onenoticesthatz∗≡0istheglobalminimumofE[z]butuninteresting.Wethusdeﬁneanylocalminimumz6=z∗tobeanillusoryphaseﬁeld,andS={x∈Ω:z>1/2}(5)tobetheassociatedillusoryshape.Theorem2SupposezisalocalminimumofE[z]andisnotalways0onΩ.ThentheassociatedillusoryshapeS6=∅.ProofOtherwiseassumeS=∅.Thenz≤1/2foranyx∈Ω.Foranyt∈(0,1),deﬁnez(t)(x)=tz.Thenz(t)∈H1
0(Ω),andfortcloseto1,z(t)isaperturbationtoz(1)=z.Onehas:E[z(t)]=t2ZΩ
2|∇z|2·Gdx+λZΩχQz2
2dx+ZΩΦ(tz)
2·Gdx,(6)whereΦ(z)=(1−z)2z2isthedouble-wellpotential.SinceΦ(z)isdecreasingon[−∞,0]andincreasingon(0,1/2],andz(x)≤1/2foranyx∈Ω,onehas:Φ(tz(x))≤Φ(z(x)),foranyx∈Ωandt∈(0,1).Thus(6)impliesthatforanyt∈(0,1),E[z(t)]≤E[z].Infact,onemusthaveE[z(t)]<E[z].ThisisbecauseZΩ
2|∇z|2·Gdx≥α
2ZΩ|∇z|2dx=α
2|z|2
1,withα=minG>0.NoticethatinH1
0(Ω),theﬁrst-orderhomogeneoussemi-norm|·|1isactuallyanormduetothezerotrace[1].Sinceithasbeenassumedthatz6=0inH1
0(Ω),onemusthave|z|2
1>0.Thustheﬁrsttermalonein(6)showsE[z(t)]<E[z]fort∈(0,1).Thiscontradictstotheassumptionthatzisalocalminimumsincez(t)’sareitssmallperturbationsfort'1.
8YoonMoJung,JianhongJackieShen5NullHypothesesandAnIterativeAlgorithm
Wenowdesignaspeciﬁciterativealgorithmtoseekvisuallymeaningfullocalminimaoftheproposedenergy:E[z]=ZΩ
2|∇z|2+(1−z)2z2
2·Gdx+λZΩχQz2
2dx,(7)forz∈H1
0(Ω),givenG∈C(Ω)andG≥0.Forthecurrentwork,Gisspeciﬁcallyconstructedasintheopeningoftheprecedingsection.Inparticular,weassumethatα=minG>0andα+β=maxG,withβ≥α.(8)Foroptimization,λcouldbe“absorbed”intoGviaG/λ.Thusweassumeβ=O(1),andλ=1.Intheliteratureofnon-convexoptimization,therehasbeenmuchdiscussiononﬁndingtheglobaloptima(e.g.,stochasticordeterministicannealing[21]).Butnouniversalmethodologiesexistforlocatingthelocaloptimathatareofpracticalinterest.Generallyithastobeproblemspeciﬁc,andthetwokeycomponentsare:(i)thechoiceofaninitialguessorstate,and(ii)thedesignofaniterativesearchingalgorithm.Fortheinitialguess,asmotivatedbyhypothesistestinginstatistics[14],weworkwiththefollowingnullhypothesis:“ThereindeedexistsanillusoryshapesomewhereoutsideQ.”(9)Astheinformationoftheillusoryshapeisunknownapriori,westartwithfollowinginitialguess:z0(x)=1·(1−χQ)+0·χQ=1−χQ.(10)Itconservativelyassignsphase1toallpixelsoutsidethegivenconﬁgurationQ.Strictlyspeaking,z0/∈H1
0(Ω).Thisisnotanissuesincez0istobeusedinaniterativealgorithm:zn→zn+1,andzn’sgeneratedafterwordsallbelongtoH1
0(Ω)forn=1,2,···.Alternatively,onemayapplydiﬀusiontoz0:ut=∆u,u|∂Ω=0,u|t=0=z0,anduseu(·,δ)∈H1
0(Ω)insteadastheinitialguessforsomeδ1.Next,todesigntheiterativealgorithm,weﬁrstcomputetheEuler-LagrangeequationofEinEqn.(7):−∇·(G∇)z+G
(z−3z2+2z3)+λχQ
·z=0,whichcanberearrangedas:−∇·(2G∇)z+(G(1+2z2)+λχQ)z=3Gz2.(11)ThisisanonlinearellipticequationonΩwithboundaryconditionz|∂Ω=0.Asdiscussedearlier,theglobalminimumz∗≡0isasolution.Weareinterestedinthenon-zerosolutionsthatareassociatedwiththelocalminimaofE[z].Ourproposediterativealgorithmistosolvethefollowinglinearellipticequa-tionforzn+1,givenacurrentguesszn:−∇·(2G∇)z+gnz=fn,z|∂Ω=0,withgn=G(1+2z2n)+λχQ,andfn=3Gz2n.(12)
IllusoryShapesviaPhaseTransition9Theorem3SupposethatGsatisﬁesEqn.(8)andzn∈L∞(Ω).Thenaweaksolutionz=zn+1∈H1
0(Ω)toEqn.(12)alwaysexistsandisuniqueinthesensethat,foranyu∈H1
0(Ω),(2G∇zn+1,∇u)+(gnzn+1,u)=(fn,u),(13)where(·,·)denotesthecanonicalinnerproductinL2(Ω,Rk).ProofInH1
0(Ω),deﬁnethesymmetricbilinearfunctionh·,·i1via:hu,vi1=(2G∇u,∇v)+(gnu,v),u,v∈H1
0(Ω),anddenotethecanonicalinnerproductinH1
0(Ω)byh·,·i0,whichisdeﬁnedby:hu,vi0=(∇u,∇v)+(u,v).ByEqn.(8),2α(∇u,∇u)≤(2G∇u,∇u)≤2(α+β)(∇u,∇u),andbythedeﬁnitionofgnin(12),α(u,u)≤(Gu,u)≤(gnu,u)≤ (α+β)(1+2kznk2
∞)+λ(u,u).Therefore,h·,·i1isaninnerproductequivalenttohu,vi0.Noticingthat|(fn,u)|≤3(α+β)kznk2
∞(1,|u|)≤3(α+β)kznk2
∞p
|Ω|p hu,ui0,(fn,·)mustbeacontinuouslinearfunctionforh·,·i0andthusalsoh·,·i1.ApplyingRieszrepresentationtheorem[13]toh·,·i1and(fn,·),oneconcludesthatthereexistsauniquezn+1∈H1
0(Ω),suchthathzn+1,ui1≡(f,u),foranyu∈H1
0(Ω),whichispreciselyEqn.(13).Wealsohaveanenergy-formdescriptionforzn+1undergivenzn.Proposition1z=zn+1∈H1
0(Ω)istheuniqueweaksolutiontoEqn.(12)ifandonlyifzn+1=argminzE[z|zn],whereE[z|zn]=ZΩ
2|∇z|2+(1+2z2n)(z−γn)2
2·Gdx+λZΩχQz2
2dx,(14)withγn=3z2n/(1+2z2n).NoticethatE[·|zn]isstrictlyconvex.ProofBydeﬁnition,foranyz,u∈H1
0(Ω),E[z+u|zn]=E[z|zn]+E∗[u|zn]+J[z,u|zn],(15)withE∗[u|zn]=ZΩ
2|∇u|2+(1+2z2n)u2
2·Gdx+λZΩχQu2
2dx,andJ[z,u|zn]=(G·∇z,∇u)+1
(G(1+2z2n)(z−γn),u)+λ
(χQ·z,u).
10YoonMoJung,JianhongJackieShenSinceE∗[u|zn]≥0,E∗[tu|zn]=t2E∗[u|zn],andJ[z,tu|zn]=tJ[z,u|zn],wecon-cludefromEqn.(15)byusingt1thatzn+1=argminzE[z|zn]iﬀJ[zn+1,u|zn]=0,foranyu∈H1
0(Ω).Ontheotherhand,·J[z,u|zn]=(2G∇z,∇u)+(gnz,u)−(fn,u),whichispreciselyEqn.(13)withz=zn+1.Thuszn+1=argminE[z|zn]ifandonlyifitsatisﬁesEqn.(13).Moreover,E∗[u|zn]>0foru6=0guaranteesstrictconvexity.Thiscompletestheproof.Thisleadstothefollowingdesirablepropertyforzn’stofaithfullyapproximate0-1binaryphases.
Proposition2Letzn+1∈H1
0(Ω)betheuniqueweaksolutiontotheiterativealgo-rithminEqn.(12)givenzn.Ifzn(x)∈[0,1]foranyx∈Ω,somustbezn+1.ProofGivenzn+1∈H1
0(Ω),deﬁnetruncationz[0,1]n+1byz[0,1]n+1(x)=

zn+1(x),if0<zn+1<1,1,ifzn+1≥1,0,ifzn+1≤0.Thenz[0,1]n+1∈H1
0(Ω),anditiswellknown[1]thattruncationdoesnotincreasethenormofagradientinSobolevspaces.Thenonemusthave:ZΩ
2

∇z[0,1]n+1

2Gdx≤ZΩ
2|∇zn+1|2Gdx.Sincezn∈[0,1],onehasz2n≤1,andforthescalarﬁeldγninEqn.(14),γn=3z2n
1+2z2n≤1.Theonemusthaveforanyx∈Ω,(z[0,1]n+1(x)−γn)2≤(zn+1(x)−γn)2andz[0,1]n+1(x)2≤(zn+1(x))2.ExaminingtheexpressioninEqn.(14)thusshows:Ehz[0,1]n+1
zni≤E[zn+1|zn].Duetotheuniquenessresultfromtheprecedingproposition,zn+1(x)≡z[0,1]n+1(x)∈[0,1],x∈Ω.Thiscompletestheproof.Thenexttheoremrevealsthatthephaseﬁeldsequence(zn)deﬁnedbytheiterativealgorithmin(12)isindeedanenergydecreasingsequence.
IllusoryShapesviaPhaseTransition11Theorem4Let{zn}∞Ωn=1bethesequenceofphaseﬁeldsgeneratedbythealgorithminEqn.(12),startedfromthenullhypothesisz0inEqn.(10).ThenE[z1]≥E[z2]≥···≥E[zn]≥···Furthermore,forn≥1,E[zn]−E[zn+1]≥ZΩG
2(zn+1−zn)22zn+4zn+1
21−zn+1
2dx,(16)wherezn+1
2=(zn+zn+1)/2.ProofBythenullhypothesisz0andProposition2,wehave0≤zn(x)≤1,forallx∈Ωandn=1,2,···.Thenzn+1
2(x)=(zn(x)+zn+1(x))/2∈[0,1]foranyx∈Ω.ThusEqn.(16)impliesE[zn]≥E[zn+1]forn≥1sincetheintegrandisnonnegative.Itsuﬃcestoestablish(16)forn≥1.ByProposition1,sincezn+1=argminzE[z|zn],onehasE[zn+1|zn]≤E[zn|zn].Thisspellsouttobe:ZΩG
2|∇zn+1|2dx+λZΩχQz2n+1
2dx+ZΩG
2(1+2z2n)z2n+1−6z2nzn+1dx≤ZΩG
2|∇zn|2dx+λZΩχQz2n
2dx+ZΩG
2(1+2z2n)z2n−6z3ndx.ThenbythedeﬁnitionofE[z]inEqn.(7),E[zn]−E[zn+1]≥ZΩG
2(1+2z2n)z2n+1−6z2nzn+1dx−ZΩG
2(1+2z2n)z2n−6z3ndx+ZΩG
2z2n(1−zn)2dx−ZΩG
2z2n+1(1−zn+1)2dx.Expandingandre-organizingtheintegrandsontheright,onearrivesatthein-equality(16).Proposition3FollowingthesameassumptionsofTheorem4,weconcludethattheremustexistsomeE∗≥0,suchthatlimn→∞E[zn]=E∗.Thisisbecauseanyboundedmonotonicsequencemustconverge.Proposition4IfE∗=limn→∞E[zn]=0,{zn}∞Ωn=1mustconvergetotheglobalminimumz∗≡0inH1
0(Ω).ProofBydeﬁnitionofE[·]inEqn.(7),andtheassumptionsonGinEqn.(8),E[z]≥α
2ZΩ|∇z|2dx≥α
2CΩkzk2
H1
0(Ω),whereCΩisapositiveconstantonlydependingonthedomain.Thesecondin-equalityholdssincethetracealong∂Ωvanishes[1].ThusE[zn]→0impliesthatzn→0inH1
0(Ω).
12YoonMoJung,JianhongJackieShenLetρn=E[zn]−E[zn+1]denotethesuccessiveenergyimprovement.Then∞Xn=1ρn=E[z1]−E∗<∞.Theseries(ρn)aresaidtoconvergeinageneralizedquadraticpower-law(GQPL)if∞Xn=1√
ρn<∞.Forinstance,ρn=1/(1+n)2+δor1/((1+n)2log2+δ(1+n))foranyδ>0.Proposition5Suppose{E[zn]}∞Ωn=1convergestosomeE∗>0inGQPL.SupposethereexistsameasurablesetS⊆ΩwithapositiveLebesguemeasure,someconstantC>0andintegerNsuchthatforanyn>N,zn(x)≥C,forx∈S.Then{zn|S}∞Ωn=1convergesinL2(S).ProofForn>N,2zn+4zn+1
2(1−zn+1
2)≥2zn≥2C.ThenEqn.(16)leadsto:ρn≥ZSG
2(zn+1−zn)2(2C)dx≥αC
kzn+1−znk2
L2(S).Thus,kzn+1−znkL2(S)≤r

αC·√
ρn.SinceP∞Ωn=1√
ρn<∞,itimpliesthat{zn}∞Ωn=1isaCauchysequenceinL2(S).Theresultcanbeintuitivelyinterpretedasfollowsforthehalf-waythresholdC=0.5.Onestartswiththenullhypothesis(viaz0)thattheillusoryshapeisS=Ω\Q.Astheiterationprogresses,somepixelsarerejectedifzn(x)<0.5.ButsupposethereexistsapositivesetS,suchthatallzn’saftersomeNconsistentlyvoteforitinthesenseofzn(x)>0.5.Thenthevotingmustbe“directional”orconverging,andShastobepartoftheﬁnalillusoryshape.6NumericalImplementationandExamples
Inthissection,webrieﬂydescribethenumericalschemeoftheproposedmodelandpresentseveralgenericcomputationalexamples.Algorithm1belowoﬀersapseudocodeblockdescribingthemajorcomputa-tionalstepsfortheproposedmodel.Fortheinitialguessz0,wehaveadoptedthenullhypothesis(9).Thecoreiterationformulaforupdatingzn+1fromznfollowsEqn.(12).Convergenceanalysisofthealgorithmhasbeenpartiallygivenintheprecedingsection.
IllusoryShapesviaPhaseTransition13
Algorithm1IllusoryShapesviaPhaseTransition input:theconﬁgurationQrepresentedviaitsindicatorχQ;initialize:z0(x)=1·(1−χQ)+0·χQ=1−χQasthenullhypothesis;pre-process:thecanyonfuntionG(x)=α+βg(|∇χQ,σ(x)|)orα+βzMSσ(x);whilekzn+1−znk>δdoSolveforzn+1from:−∇·(2G∇)z+gnz=fn,withgn=G(1+2z2n)+λχQ,fn=3Gz2n,andz|∂Ω=0.endwhile
Therefore,thecoreofthealgorithmisanellipticsolver,whichcanbefoundinthestandardliteratureofcomputationalPDE’s[32].Sometypicalparametersgeneratingtheexampleshereinaregivenasfollows:α=0.1;β=1.0;λ=1.0;and=(2∼4)·h.Herehdenotesthegrid/pixelsizeandisdeﬁnedinsuchawaythatthelongestsizeoftheimagedomainΩisalwaysnormalizedtotheunitlength.Theconvergencetoleranceissettobeδ=10−6.
(a)KaninzaTriangle
(b)CanyonFunctionG
(c)NullHypothesisz0(x)
(d)20Iterations
(e)100Iterations
(f)FinalIllusoryShapeFig.3ThemodelandalgorithmontheKanizsaTriangle.Fig.3showsthenumericalsimulationontheclassicexampleofKanizsa’sTriangle.Inallthepanels(b-f),thewhitecolorsrepresentphasevaluescloseto1.0whiletheblackonesto0.0.Panel(b)showsthecanyonfunctionthatisfedintothemodelandalgorithm.Panel(c)showsthenullhypothesisz0(x)thatconservativelyassignsphase1.0toallpixelsoutsideQ.Panels(d)and(e)showtwointermediateiterationsbeforeﬁnalconvergence.InPanel(f)theﬁnalphaseﬁeldisplottedafternumericalconvergence.Itsuccessfullycapturestheillusory
14YoonMoJung,JianhongJackieShentriangle(uptothenumericalprecision).Fig.4andFig.5demonstrateanothertwoexampleswithmorecomplexlayouts.
(a)IllusoryDisk
(b)IntermediateSnapshot
(c)FinalResultFig.4Themodelandalgorithmonanexamplewithanillusorydisk.
(a)IllusoryLady
(b)ConvergedResultFig.5TheButterﬂyLady,withamorecomplexconﬁgurationQ.Fig.6showsanexampleconsistingoftwodisjointillusoryshapes:anellipseandatriangle.Asaresult,thephaseﬁeldsequence(zn)isexpectedtoexperienceatopologicalsplittingoperationduringtheiteration.Liketherenownedlevel-setmethodologyofOsherandSethian[30],thephase-ﬁeldapproachisalsoveryversatileinhandlingregionmergingorsplitting.
7Conclusion
Inspiredbytheﬁrst-ordervariationalillusorycontour(VIC)modelproposedin[19],wehaveproposedavariationalillusoryshape(VIS)modelbasedonthetoolofphasetransitions.TheVISmodelrepresentsanillusoryshapeviaphasevaluescloseto1.0,andtherestbyvaluescloseto0.0.Thephasetransitionisachievedbyavariationalenergyformulatedinthecurrentwork.Asformostnon-quadraticphasetransitionmodels[26,27,33,34],theproposedVISmodelisnon-convex.Thezeroﬁeldistheglobaloptimumbutuninteresting.Toseekvisuallymeaningfullocaloptima,wehavedesignedaniterativealgorithmwithasuitableinitialguess,whichcouldbeconsideredasthenullhypothesisin
IllusoryShapesviaPhaseTransition15
(a)IllusoryEllipseandTriangle
(b)ConvergedResultfromtheModelFig.6Likethelevel-setmethodologyofOsherandSethian[30],thephase-ﬁeldapproachisversatileinhandlingtopologicalchangeslikeregionsplitting.statisticaltesting.Thenullhypothesisassumesthatthereexistsanillusoryshapeoutsidethegivenconﬁguration.Theiterativealgorithmletsthepixels“vote”collectively,untilreachingtheﬁnalconsistentandstationarydecision.Somekeybehaviorsofthealgorithmhavebeenrevealedthroughouranalysis.Andseveralgenericnumericalexamplesshowtheversatilityoftheproposedmodelandalgo-rithm.Asin[19],suchlower-ordermodelsallowonetodevelopdetailedanalysis,butarenecessarilylimitedintermsofapplicabilityorperformance.Forexample,illusoryinterpolationisoftendoneviastraightlines.Nevertheless,theyhelppointtowardsmorecomplexhigh-ordermodelsinvolvingthecurvaturefeatureorEuler’selasticas[9,8,28,20],forexample.
AcknowledgementsJunghasbeensupportedbyBasicScienceResearchProgramthroughtheNationalResearchFoundationofKorea(NRF)ofKorea(2012R1A1A1015492).ShenhasbeensupportedbytheNationalScienceFoundation(NSF)ofUSA.References1.R.A.AdamsandJ.J.F.Fournier.SobolevSpaces.AcademicPress,secondedition,2003.2.L.AmbrosioandV.M.Tortorelli.ApproximationoffunctionalsdependingonjumpsbyellipticfunctionalsviaΓ-convergence.Comm.PureAppl.Math.,43:999–1036,1990.3.L.AmbrosioandV.M.Tortorelli.Ontheapproximationoffreediscontinuityproblems.Boll.Un.Mat.Ital.,6-B:105–123,1992.4.G.AubertandP.Kornprobst.MathematicalProblemsinImageProcessing.Springer-Verlag,2001.5.G.AubertandL.Vese.Avariationalmethodinimagerecovery.SIAMJ.Numer.Anal.,34:1948–1979,1997.6.W.M.Boothby.AnintroductiontodiﬀerentiablemanifoldsandRiemanniangeometry,volume120.AcademicPressInc.,secondedition,1986.7.A.Braides.Γ-convergenceforbeginners.OxfordUniversityPress,Oxford,2002.8.T.F.Chan,S.-H.Kang,andJ.Shen.Euler’selasticaandcurvaturebasedinpainting.SIAMJ.Appl.Math.,63(2):564–592,2002.9.T.F.ChanandJ.Shen.ImageProcessingandAnalysis:variational,PDE,wavelet,andstochasticmethods.SIAMPublisher,Philadelphia,2005.10.T.F.ChanandW.Zhu.Captureillusorycontours:Alevelsetbasedapproach.UCLACAMReport03-65,2003.11.D.Chandler.IntroductiontoModernStatisticalMechanics.OxfordUniversityPress,NewYorkandOxford,1987.
16YoonMoJung,JianhongJackieShen12.G.DalMaso.AnIntroductiontoΓ-Convergence.Birkhauser,Boston,1992.13.G.B.Folland.RealAnalysis-ModernTechniquesandTheirApplications.JohnWiley&Sons,Inc.,secondedition,1999.14.D.Freedman,R.Pisani,andR.Purves.Statistics.W.W.NortonandCo.,2007.15.K.Fukushima.Neuralnetworkmodelforcompletingoccludedcontours.NeuralNetworks,23:528–540,2010.16.D.H.Grosof,R.M.Shapley,andM.J.Hawken.MacaqueV1neuronscansignalillusorycontours.Nature,365:550–552,1993.17.R.Hales.Jordan’sproofoftheJordancurvetheorem.StudiesinLogic,GrammarandRhetoric,10(23):45–60,2007.18.F.HanandS.C.Zhu.Bottom-up/top-downimageparsingwithattributegraphgrammars.IEEETrans.PatternAnal.MachineIntelli.,31(1):59–73,2009.19.Y.M.JungandJ.Shen.First-ordermodelingandstabilityanalysisofillusorycontours.J.VisualCommun.ImageRepresentation,19:42–55,2008.20.S.-H.Kang,W.Zhu,andJ.Shen.Illusoryshapesviacornerfusion.SIAMJ.ImagingSci.,toappear,2014.21.S.Kirkpatrick,C.D.Gelatt,andM.P.Vecchi.Optimizationbysimulatedannealing.Science,220:671–680,1983.22.D.C.KnillandW.Richards.PerceptionasBayesianInference.CambridgeUniv.Press,1996.23.T.S.LeeandD.Mumford.HierarchicalBayesianinferenceinthevisualcortex.J.Opt.Soc.Am.(A),20(7):1434–1448,2003.24.T.S.LeeandM.Nguyen.Dynamicsofsubjectivecontourformationintheearlyvisualcortex.Proc.Natl.Acad.Sci.U.S.A,98:1907–1911,2001.25.J.L´eveill´e,M.Versace,andS.Grossberg.Runningasfastasitcan:Howspikingdynamicsformobjectgroupsinthelaminarcircuitsofvisualcortex.J.Comput.Neurosci.,28:323–346,2010.26.R.March.Visualreconstructionwithdiscontinuitiesusingvariationalmethods.ImageVisionComput.,10:30–38,1992.27.R.MarchandM.Dozio.Avariationalmethodfortherecoveryofsmoothboundaries.ImageVisionComput.,15:705–712,1997.28.D.Mumford.Elasticaandcomputervision.InC.L.Bajaj,editor,AlgebraicGeometryanditsApplications,pages491–506.Springer-Verlag,NewYork,1994.29.M.M.MurrayandC.S.Herrmann.Illusorycontours:Awindowontotheneurophysiologyofconstructingperception.TrendsinCognitiveSciences,17(9):471–481,2013.30.S.OsherandJ.A.Sethian.Frontspropagatingwithcurvature-dependentspeed:Algo-rithmsbasedonHamilton-Jacobiformulations.J.Comput.Phys.,79(12):12–49,1988.31.A.Sarti,R.Malladi,andJ.A.Sethian.Subjectivesurfaces:Ageometricmodelforboundarycompletion.Int’lJ.Comp.Vision,46(3):201–221,2002.32.T.Sauer.NumericalAnalysis.Pearson,2011.33.J.Shen.Γ-convergenceapproximationtopiecewiseconstantMumford-Shahsegmentation.Lec.NotesComp.Sci.,3708:499–506,2005.34.J.Shen.Astochastic-variationalmodelforsoftMumford-Shahsegmentation.Int’lJ.Biomed.Imag.,2006(92329):1–14,2006.35.D.A.StanleyandN.Rubin.fMRIactivationinresponsetoillusorycontoursandsalientregionsinthehumanlateraloccipitalcomplex.Neuron,37:323–331,2003.36.L.A.Vese.AstudyintheBVspaceofadenoising-deblurringvariationalproblem.Appl.Math.Optim.,44(2):131–161,2001.37.R.vonderHeydtandE.Peterhans.Mechanismofcontourperceptioninmonkeyvisualcortex.I.Linesofpatterndiscontinuity.J.Neurosci.,9:1731–1748,1989.38.R.vonderHeydt,E.Peterhans,andG.Baumgartner.Illusorycontoursandcorticalneuronresponses.Science,224(1260-1262):1984.39.T.F.WuandS.C.Zhu.Anumericstudyofthebottom-upandtop-downinferencepro-cessesinand-orgraphs.Int’lJ.Comput.Vision,93(2):226–252,2011.40.A.Yoshino,M.Kawamoto,T.Yhoshida,N.Kobayashi,J.Shigemura,Y.Takahashi,andS.Nomura.Activationtimecourseofresponsestoillusorycontoursandsalientregion:Ahigh-densityelectricalmappingcomparison.BrainResearch,1071:137–144,2006.41.W.ZhuandT.Chan.Illusorycontoursusingshapeinformation.UCLACAMTech.Report,03-09,2005.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

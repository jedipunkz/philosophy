---
source: "https://arxiv.org/abs/1809.08021v4"
title: "Convergence to $α$-stable Lévy motion for chaotic billiards with several cusps at flat points"
author: "Paul Jung, Françoise Pène, Hong-Kun Zhang"
year: "2018"
publication: "arXiv preprint / math.DS"
download: "https://arxiv.org/pdf/1809.08021v4"
pdf: "https://arxiv.org/pdf/1809.08021v4"
captured_at: "2026-05-09T12:28:07Z"
updated_at: "2026-05-09T12:28:07Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代思想"
status: raw
---

# Convergence to $α$-stable Lévy motion for chaotic billiards with several cusps at flat points

- 著者: Paul Jung, Françoise Pène, Hong-Kun Zhang
- 年: 2018
- 掲載情報: arXiv preprint / math.DS
- 情報源: [arxiv](https://arxiv.org/abs/1809.08021v4)
- ダウンロード: https://arxiv.org/pdf/1809.08021v4
- PDF: https://arxiv.org/pdf/1809.08021v4

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

We consider billiards with several possibly non-isometric and asymmetric cusps at flat points; the case of a single symmetric cusp was studied previously in Zhang (2017) and Jung & Zhang (2018). In particular, we show that properly normalized Birkhoff sums of Hölder observables, with respect to the billiard map, converge in Skorokhod's $M_1$-topology to an $α$-stable Lévy motion, where $α$ depends on the `curvature' of the flattest points and the skewness parameter $ξ$ depends on the values of the observable at those same points. Previously, Jung & Zhang (2018) proved convergence of the one-point marginals to totally skewed $α$-stable distributions for a single symmetric cusp. The limits we prove here are stronger, since they are in the functional sense, but also allow for more varied behaviour due to the presence of multiple cusps. In particular, the general limits we obtain allow for any skewness parameter, as opposed to just the totally skewed cases. We also show that convergence in the stronger $J_1$-topology is not possible.

## PDF Text

arXiv:1809.08021v4 [math.DS] 20 May 2019
Convergencetoα-stableLévymotionforchaoticbilliardswithseveralcuspsatﬂatpointsPaulJung1andFrançoisePène2andHong-KunZhang31DepartmentofMathematicalSciences,KAIST,Daejeon,SouthKorea2UniversitédeBrest,InstitutUniversitairedeFrance,LMBA,UMRCNRS6205,Brest,France3DepartmentofMathematics,UniversityofMassachusetts,Amherst,MA,USANovember27,2024AbstractWeconsiderbilliardswithseveralpossiblynon-isometricandasymmetriccuspsatﬂatpoints;thecaseofasinglesymmetriccuspwasstudiedpreviouslyin[Zha17]and[JZ18].Inparticular,weshowthatproperlynormalizedBirkhoffsumsofHölderobservables,withrespecttothebilliardmap,convergeinSkorokhod’sM1-topologytoanα-stableLévymotion,whereαdependsonthe‘curvature’oftheﬂattestpointsandtheskewnessparameterξdependsonthevaluesoftheobservableatthosesamepoints.Previously,[JZ18]provedconvergenceoftheone-pointmarginalstototallyskewedα-stabledistributionsforasymmetriccusp.Thelimitsweproveherearestronger,sincetheyareinthefunctionalsense,butalsoallowformorevariedbehaviourduetothepresenceofmultiplecusps.Inparticular,thegenerallimitsweobtainallowforanyskewnessparameter,asopposedtojustthetotallyskewedcases.WealsoshowthatconvergenceinthestrongerJ1-topologyisnotpossible.Contents
1Introduction2
2Mainresult6
3ProofsofTheorems2.1and2.2usinganinducedmap73.1ConvergencefortheM1-metric..................................73.2Non-convergencefortheJ1-metric................................94ProofofTheorem3.1(functionallimittheoremfortheinducedsystem)104.1Simpliﬁcation...........................................104.2Convergenceindistribution....................................124.3AkeyestimateforConditionI..................................131
4.4ConditionI:Pointprocessconvergence..............................144.4.1Proofof(4.9)........................................154.4.2Proofof(4.10)........................................154.5ConditionI:alternativeapproach.................................174.6ConditionII:Vanishingsmallvalues...............................19AAsymmetriccuspswithﬂatnessβ23A.1Thecornerseries.........................................24A.2ProofofLemma4.5........................................30BSkorokhodJ1andM1topologies341Introduction
Theoriginsofthemoderntheoryofhyperbolicdynamicalsystemslieinclassicalandstatisticalmechanicsthroughthestudyofergodicandstatisticalproperties.Indeedunderstandingstatisticalpropertiesandprovingvariousprobabilitylimittheoremsarevitalinthestudyofstatisticalmechanics.Forexample,suchstudiesshedlightontheimportantissueofﬂuctuationsinentropyproduction.ThispaperisacontributioninthisdirectionbyshowingamechanismbywhichanhyberbolicdynamicalsystemcanconvergeontheprocessleveltosomethingotherthanaBrownianmotion.BeginbylettingTbeameasure-preservingtransformationon(M,µ)and{Xn=f◦Tn}thestochasticprocessgeneratedbythesystemforanobservablef.Onetypicalgoalistoprovethe(functional)centrallimittheoremforthenormalizedpartial-sumprocessassociatedto{Xn}whichisitselfgeneratedbyamixinghyperbolicsystem,i.e.,toprovethat,weakly
1
√
n[nt]Xj=1Xj
t≥0→(σB(t))t≥0where(B(t))t≥0isastandardBrownianmotion.Inotherwords,the(normalized)partial-sumprocessconvergesindistribution,inC[0,1],toaBrownianmotionwithdiffusionparameterσ>0.Theﬁnitenessofthevariance/diffusionparameterrequiresthatthecovariancesCov(X1,Xn)aresummableinn∈N.However,whenthecovariancesequenceisnotsummable,theusualcentrallimittheoremforthepartial-sumprocessfails;andthereareveryfewresults,forhyperbolicsystems,whichinvestigatewhathappensundersuchfailure.Billiardswithcuspswereﬁrstintroducedintheearly80’sin[Mac83],butitwasnotuntil[CM07]thatarigorouspolynomialdecayofcorrelationsofO(1/n)wasprovedforMachta’soriginalmodel(wherecuspsareformedbytangentialcircles).Thispolynomialdecayofcorrelationsledtoanonstandardcentrallimittheoremwhereanormalizationof√
nlognwasusedinsteadof√
n(seealso[BCD11]).In[Zha17],theauthorwasabletoconstructahyperbolicbilliardmodelwitharbitrarilyslowdecayratesofcorrelations,oforderCov(X1,Xn)=O(n−a),2
witha∈(0,1)dependingonthecurvatureatthecusp’svertex.In[JZ18],theﬁrstandlastauthorswereabletoprovethatduetotheslowdecayofcorrelations,insteadofacentrallimittheorem,rather,astablelimittheoremholdsforthepartial-sumprocessgeneratedbyabilliardwithasinglesymmetriccuspataﬂatpoint.Inparticular,theyshowedthatthelimitingdistributionisatotallyskewedα-stablelaw.Inthispaper,weﬁrstinvestigatetheconvergencetoastablelaw,forthecasewhenthebilliardtablehasseveral,possiblyasymmetric,cuspsatﬂatpoints.WethenextendtheseresultstoafunctionallimittheoreminSkorokhodspace,i.e.,convergencetoanα-stableLévymotion.Inthedynamicalsystemsliterature,functionalconvergencetoLévymotionhasbeenshownforexpandingmaps(seeforinstance[TK10]),howeverthisistheﬁrstresultofthiskindfor2-dhyperbolicbilliardsthatweareawareof(afterﬁnishingthiswork,weweremadeawareoftheconcurrentwork[MV18]whereafunctionallimittheoremisalsoprovedforthemodelin[JZ18].).Letusnowdescribeindetailthebilliardswithcuspsatﬂatpointsthatweconsiderhere.ThedispersingbilliardtableQisaboundeddomainofR2,theboundaryofwhichconsistsofq≥3dispersingC3smoothcurves∂Q=∪i∈Z/qZΓi,numberedinclockwiseorder.TheintersectionsPi=Γi∩Γi+1oftwoconsecutivecurvesΓiandΓi+1consistofeitherstandardcornerpoints(i.e.thetangentlinesatPitoΓiandΓi+1donotcoincide)or‘cusps’atﬂatpoints.Thismeansthat,inanappropriateeuclideancoordinatesystem(s,z)originatedatsomepointPi,thetwocurvesΓiandΓi+1canberepresentedrespectivelyasz=zi,+(s)andz=zi,−(s)withzi,±differentiableandsatisfyingzi,±(s)=±Ci,±
βisβi+O(s2βi−1)andz0i,±(s)=±Ci,±sβi−1+O(s2βi−2),fors∈[0,](1.1)with>0somesmallﬁxednumber,Ci,−,Ci,+≥0notbothnullandwithβi≥2.WewillsaythatsuchacuspatPihasﬂatnessβi.Wealsodeﬁne¯Ci:=(Ci,++Ci,−)/2.Weassumemoreoverthat,foreverycuspatsomePi,the(unique)tangenttrajectorycomingoutofthecuspatPihits∂Qoutsideofanyanothercusp.Also,allboundarycomponentsareassumedtobedispersingandhavecurvatureboundedawayfromzeroexceptatthecuspsPi’s.Throughoutthepaper,weassumethatβ∗:=max(βi:Piisacusp)>2.(1.2)Thefunctionallimittheoremweproveherewilldependonlyontheobservablenearthe‘maximallyﬂat’pointsPi(theheuristicreasonforthisisthatthelimittheoremisdominatedbylongsequencesofcollisionsinsidethevariouscusps,andtheﬂattestcuspstrapthelongestsequences).LetJbethesetoflabelsisuchthatPiisacuspandletJ∗bethesetoflabelsi∈Jsuchthatβi=β∗.TheusualbilliardﬂowisdeﬁnedontheunitspherebundleQ×S1andpreservestheLiouvillemeasure,see[CM06]fordetails.WeconsiderthenaturalcrosssectionM⊂Q×S1madeofallpost-collisionvectorsbasedattheboundaryofthetable∂Q.Anypost-collisionvectorx∈Mcanberepresentedbyx=(r,ϕ),wherer∈R/|∂Q|Zistheclockwisecurvilinearabscissaalong∂Q(where|∂Q|denotesthelengthof∂Q),andϕ∈[0,π]istheangleformedbythetangentlineoftheboundaryandthecollisionvectorintheclockwisedirection.ThereforethecollisionspaceMisidentiﬁedwithR/|∂Q|Z×[0,π].Thecorrespondingbilliard3
QP1P2P3P4P5P6
Figure1:Abilliardtablewithseveralcuspsatﬂatpoints.mapT:M→Mtakesavectorx∈Mtothenextpost-collisionvectoralongthetrajectoryofx.LetthesetS0consistofallgrazingcollisionvectorswithwallsaswellasallcollisionvectorsatcornerpoints/cusps.ThenS0:=S0∪T−1S0isthesingularsetofT.ThebilliardmapT:M\S0→M\TS0isalocalC2diffeomorphismandpreservesanaturalabsolutelycontinuousprobabilitymeasuredµ=1
2|∂Q|sinϕdrdϕonthecollisionspaceM={(r,ϕ)∈R/|∂Q|Z×[0,π]}.Arecentlypopularapproachforstudyingthestatisticalpropertiesof(T,M)istouseaninducingschemeasintroducedin[Mar04,CZ05].Byremovingspotswithweakhyperbolicityfromthephasespace,oneconsiderstheﬁrstreturnmaponsomesubspaceM⊂M.Foranyx∈M,theﬁrstreturntimefunctionisdeﬁnedbyR(x):=min{n≥1:Tn(x)∈M}andthe
(ﬁrst)
return mapF:M→MisdeﬁnedbyF(x):=TR(x)(x),forallx∈M.(1.3)ThereturnmapFpreservestheconditionalmeasure˜µ:=1
µ(M)µ|Mandwedeﬁnean inducedfunctionby˜f(x):=R(x)−1Xk=0f(Tkx),x∈M.(1.4)HerewedeﬁnethesubsetMtobethesubsetofMwhichconsistsofallcollisionsoccurringon∂Q\(∪i∈JB(Pi))(1.5)4
whereB(Pi)∩∂Qadmitsarepresentationasin(1.1)andwhereissmallenoughsothatabilliardtrajectorycannotgodirectlyfromoneB(Pi)toanotherB(Pi)withouthitting∂Q\(∪k∈JB(Pk)).WedecomposeM=M0∪[i∈JMisothatfori∈J,MiismadeupofthosepointsinMwhoseﬁrstcollisioninside∪j∈JB(Pj)(see(1.5)),underthemapF,occursinsideB(Pi).Moreover,M0=M\∪i∈JMiconsistsoftheremainingpointswhichdonotimmediatelyenteranycusps.Rigorousboundsonthedecayofcorrelationsforbilliardswithﬂatpointswerederivedin[Zha17],whereamoredetaileddescriptionofbilliardswithﬂatpointsisalsogiven.Itwasshownthatiff,gareHöldercontinuousfunctionsonthecollisionspaceM,thenforalln∈Z,µ(f◦Tn·g)−µ(f)µ(g)=O(n−1
β∗−1).(1.6)Hereweusethestandardnotationµ(f)=RMfdµ,whichwehenceforthassumetobe0forfandg.Asalreadymentioned,itistheaboveslowdecayofcorrelationsthatledtoastablelimittheoremin[JZ18],ratherthanthestandardcentrallimittheorem,foranormalizedversionoftheBirkhoffsumsSnf:=f+f◦T+···+f◦Tn−1,wherefisaHölderfunctiononMandf6=0atthecusp’svertex.Ourresultswillshowthatforageneraldispersingbilliardtablewithmultiplecusps,onlythe|J∗|≥1maximallyﬂatcuspswillcontributetothelimitofproperlynormalizedBirkhoffsums.Inotherwords,weshowthatcuspswithsmallerorderﬂatnesswillnot‘contribute’tothelimitingstablelaw(nortothefunctional-levelconvergencetoaLévyprocess).Althoughtheanalysisin[JZ18]canbecarriedtoageneraltabletostudystatisticalpropertiesforobservablessupportedonanyindividualsymmetriccusp,theconvergencetoastablelawjointlyformultipleasymmetriccuspshereisratherdifferentandrequiressigniﬁcantadjustmentstotheoriginalarguments.Inordertoconsiderafunctionallimittheorem,denotetheprocess{Wn(t):t≥0},n≥1byWn(t):=[nt]−1Xj=0f◦Tj n1/α,fort≥0,whereα:=β∗
β∗−1.(1.7)Onecancheckthatα∈(1,2)sinceβ∗>2.In[MZ15],ageneralresultwasprovedforobtainingfunctionalconvergenceofdynamicalsystemstoanα-stableLévymotionbyliftingsuchalimitlawfromaninduceddynamicalsystemtotheoriginalsystem.Ourfunctionallimittheoremfor{Wn(t):t≥0},n≥1,willbeprovedintwosteps:ﬁrstshowingthattheinducedsystemsatisﬁesthefunctionallimittheorem,utilizingstandardmethodsfrom[DR78],andthenapplyingtheliftingprincipleof[MZ15].Therestofthepaperisorganizedasfollows.Inthenextsectionwestateourmainresult.InSection3wegivethetop-levelproofofourmainresultasjustdescribedinthepreviousparagraph.InSection4weprovidethedetailsofthelimittheoremfortheinducedsystemviathemechanismof[DR78].Intheappendixwepresenttheadditionaltechnicaldetailsneededtohandlethegeneralsettingofasymmetriccuspsweconsiderhere,andalsoprovide(foranyreaderwhoisseeingthesefortheﬁrsttime)averybriefintroductiontotheSkorokhodJ1andM1topologies.5
2Mainresult
Afunctionfwithµ(f)=0issaidtobeinthedomainofattractionofa(strictly)α-stablelawifthereexists{bn}suchthat{Snf bn}convergesindistributiontoarandomvariablewithanα-stablelaw.Here, strictlysimplymeansthatµ(f)=0,andweshallhenceforthjustsayα-stable.ContrarytocentrallimitbehaviorandtheGaussiancaseofα=2,itiswell-knownthateventhoughwehaveameanofzero,thelimitingdistributionmaynotbesymmetric.Inparticular,thereexistsa skewness parameterξ∈[−1,1]ands>0suchthatthelimitstablerandomvariableYsatisﬁeslimx→∞xαP(Y>x)=Cαsα1+ξ
2andlimx→−∞|x|αP(Y<x)=Cαsα1−ξ
2,(2.1)withCαgivenbyCα=1
Γ(1−α)cos(πα/2),see[ST94,p.17].WewillhenceforthdenotebyZα,ξ,sastablerandomvariablewithtaildistributionsatisfying(2.1),i.e.withcharacteristicfunctionE eiuZα,ξ,s=exp−|us|α1−iξsign(u)tanπα
2,u∈R.(2.2)Foranyγ∈(0,1),wedenoteHγastheclassofallHöldercontinuousfunctionsf:M\S0→R,withHölderexponentγ.Letuswrite˜riforthecurvilinearabscissaofthecusppositionPi.Wedeﬁne∀ϕ∈[0,π],˜fi,+(ϕ):=limr→˜ri+f(r,ϕ)and˜fi,−(ϕ):=limr→˜ri−f(r,ϕ),sothat˜fi,−(resp.˜fi,+)correspondstothelimitfunctiononPi×[0,π]off|Γi(resp.f|Γi+1).WealsodeﬁneIf,i=If,i,α:=1
4Zπ0(˜fi,−(ϕ)+˜fi,+(ϕ))sin1
αϕdϕ.(2.3)Theorem2.1(Stablelimittheoremforbilliardwithcusps).LetQbeabilliardtableasdescribedabove,withcuspsdeﬁnedby(1.1).Supposef∈Hγforsomeγ>0satisfyingµ(f)=0andsupposethereexistssomei∈J∗suchthatIf,i6=0.Thenasn→∞,Snf n1/αd−→Xi∈J∗,If,i6=0Zα,ξf,i,σf,i
C1/αα,(2.4)wherethelimitisasumofindependentstablevariableswithσαf,i:=2
|∂Q|·Iαf,i
β¯Cα−1iandξf,i:=sign(If,i).6
Theorem2.2(Functionalα-stablelimittheorem).UndertheassumptionsofTheorem2.1if,foreveryi∈J∗,thereexistsaneighbourhoodUiofPi×[0,π]suchthateitherf|Ui≥0orf|Ui≤0,then,inadditiontotheconvergencein(2.4),((Wn(t))t∈[0,1])n,deﬁnedin(1.7),convergesindistributionintheSkorokhodspaceD([0,1])usingtheM1-metric,toanα-stableLévymotion(Yt)t∈[0,1](anα-stableprocesswithstationaryandindependentincrements)suchthatY1hasthesamedistributionastherightsideof(2.4).Moreover,thisconvergencedoesnotholdfortheJ1-metric.Notethatthevaluesoffatthemaximallyﬂatcuspsdeterminethevalueξandσandinfactonecancomputethattherightsideof(2.4)isjustarandomvariableZα,ξf,C−1
αασfwithσαf:=Xi∈J∗2|If,i|α
β¯Cα−1i|∂Q|andξf:=Pi∈J∗sign(If,i)¯C1−αi|If,i|α
Pi∈J∗¯C1−αi|If,i|α.(2.5)Letusremarkthat,asin[JZ18],Theorem2.2(aswellasTheorem3.1below)holdsalsoforfboundedandpiecewiseγ-Hölder,forsomeγ>0,withdiscontinuitiescontainedinthesingularsetofTandsuchthatfisγ-HölderinaneighborhoodoftheregioninMcorrespondingtoeachofthecusps.3ProofsofTheorems2.1and2.2usinganinducedmap
3.1ConvergencefortheM1-metricRecallthatF(x):=TR(x)(x),withR(x):=min{n≥1:Tn(x)∈M}.Foreveryg:M→R,wedenotetheBirkhoffsumsontheinducedspacebySng(x):=n−1Xk=0g◦Fk(x),x∈M,n≥0.Recallalsothat,givenf:M→R,wedeﬁne˜f:M→Rbysetting˜f(x):=SR(x)f(x).(3.1)Wesimilarlydeﬁnetheprocess˜Wn:=(˜Wn(t))t∈[0,1]asfollows:˜Wn(t):=n−1/αS[nt]˜f.(3.2)Anintermediatefunctionallimittheoremfortheinducedmapisasfollows.7
Theorem3.1.Letf:M→RbeasinTheorem2.1.Then(˜Wn)nconvergesindistribution,intheSkorokhodJ1-topology,toanα-stableLévymotionwithα=β
β−1,skewnessparameterξf,andscaleparameters=C−1/αα˜σf:=(Cαµ(M))−1/ασf,withσfandξfdeﬁnedin(2.5)BothTheorems2.1and2.2areconsequencesofTheorem3.1.ProofofTheorem2.1.ByconsideringtheﬁrsthittingtimeoftheM(inlieuoftheﬁrstreturntime)wemayextendthedeﬁnitionsofRandFtoallofM,sothatinparticular˜f◦FisdeﬁnedonallofM.DuetoTheorem3.1,(n−1
αSbntc˜f◦F)t≥0nconvergesindistribution,withrespecttoµandtothemetricJ1,toaLévyprocess(Yt)t≥0.Moreover,foreverypositiveintegerm,wedeﬁneτm(x)asthenumberofvisitsof(Tk(x))k=1,...,mtoMsothatτm(x)−1Xk=0R◦Fk(x)≤m<τm(x)Xk=0R◦Fk(x).Duetotheergodictheorem(τm/m)m≥1convergesalmostsurelytoµ(M)andso(n−1
αSτbntc−1˜f◦F)t≥0nconvergesindistribution,withrespecttoµandtoJ1,to(Ytµ(M))t≥0.Moreover,

Smf(x)−Sτm(x)−1˜f(F(x))

≤kfk∞(R(x)+R−(Tm(x))),whereR−correspondstothelengthoftimesincethelastvisittoM.Thereforebyreversibilityofthedynamics,forevery>0,µn−1
α

Sbntcf−Sτbntc(x)−1˜f◦F

>≤2µn−1
αkfk∞R>ε/2.Thisimpliestheconvergenceinprobabilityof(n−1
α(Sbntcf(x)−Sτbntc(x)−1˜f(F(x))))nto0foreveryt≥0andsotheconvergenceoftheﬁnite-dimensionaldistributions(withrespecttoµ)of(n−1
αSbntcf)t≥0ntotheonesof(Ytµ(M))t≥0.
TheresultofconvergenceindistributionofTheorem2.2willfollowdirectlyfromTheorem3.1andProposition3.2below.Thelatterisaversionof[MZ15,Thm2.2],thestatementofwhich,requiresthefollowingnotion:f∗(x):=max0≤`0≤`≤R(x)(S`0f(x)−S`f(x))∧max0≤`0≤`≤R(x)(S`f(x)−S`0f(x)).8
Proposition3.2(Liftingaweakinvarianceprinciple,MelbourneandZweimüller).Undertheaboveassump-tionsonfand˜Wn,if(˜Wn(t),t∈[0,1])convergesweaklyintheSkorokhodM1-topology,asn→∞,toanα-stableLévymotion(W(t),t∈[0,1])andn−1/αmax0≤k≤nf∗◦Fkd−→0.(3.3)then(Wn(s),s≥0)d−→(W(sµ(M)),s≥0)intheSkorokhodM1-topology.Notethattheversionofthisresultin[MZ15]assumesstrongdistributionalconvergenceoftheinducedprocess,butProposition2.8thereallowsustostateitaswedidabove.ProofofTheorem2.2.Weneedonlychecktheso-calledweakmonotonicitycondition(3.3)inoursetting.SinceftakesasinglesignintheneighborhoodaroundeachcuspPiwithi∈J∗,wehavenearlyfullmonotonicityoftheergodicsumsduringanexcursion,withtheonlyexceptionpossiblycomingfromtheﬁrststep.Thusf∗≤kfk∞R1∪i∈J\J∗Mi+kfk∞1∪i∈J∗Mi≤kfk∞(R1∪i∈J\J∗Mi+1).Nowlet˜α:=min{βi
βi−1:i∈J\J∗},whichbyassumptionsatisﬁes˜α>α.DuetoSection4,Yn(t):=n−1/αS[nt]R1∪i∈J\J∗Mi−˜µ(R1∪i∈J\J∗Mi)tconvergesintheJ1-metric,asngoestoinﬁnity,tothe0function(here˜µ:=µ(·|M)).Thereforen−1
αmax0≤k≤nf∗◦Fk≤kfk∞n−1
αmax0≤k≤nR1∪i∈J\J∗Mi−˜µ(R1∪i∈J\J∗Mi)◦Fk+1+˜µ(R)≤kfk∞

maxt∈[0,1]Yn(t)−mint∈[0,1]Yn(t)

+n−1
α(1+˜µ(R)),whichconvergesindistributionto0.
Thischecksthat(3.3)holdswhichcompletestheproofofTheorem2.2uptoaproofofTheorem3.1whichisgiveninSection4.
3.2Non-convergencefortheJ1-metricInthissubsection,weexplainwhytheconvergenceinTheorem2.2doesnotholdintheJ1-topology(see[TK10,Example2.1]or[AT92]forotherredactionsofthisargument).Let(wn(t))tbethecontinuousprocessobtainedfrom(Wn(t))tbylinearization:wn(t):=Wn(t)+(nt−bntc)f◦Tbntc n1/α,t∈[0,1].9
Sincefisuniformlybounded,wealsohavesupt∈[0,1]|wn(t)−Wn(t)|≤kfk∞
n1/α−→0,asn→∞.Therefore,theconvergenceindistributionof(Wn(t))nintheJ1-metricwouldimplythesameconvergenceindistributionfor(wn(t))nintheJ1-metricandsofortheuniformmetriconeverycompactinterval,thiswouldcontradictthefactthatthelimitingprocessisdiscontinuous.NotethatthereasonthisargumentdoesnotapplytotheinducedprocessinTheorem3.1(ortoJ1-convergenceforgeneralLévyprocesses),isbecauseinthatsettingk˜fk∞=∞.4ProofofTheorem3.1(functionallimittheoremfortheinducedsys-tem)AssumetheassumptionsofTheorem2.1.Recallthat˜µ=µ(·|M).Wedeﬁne˜fi:=˜f·1Mi−˜µ(˜f·1Mi).Notethat,duetotheKactheorem,˜µ(R)<∞.(4.1)4.1Simpliﬁcation
InordertoproveTheorem3.1,weshowthatwecanreplace˜fbythesimplestfunction˜Rgivenby:˜R(x):=Xi∈JIf,i,αi
I1,αi(Ri(x)−˜µ(Ri)),withRi:=R1Mi(4.2)withI1,αi:=I1,i,αi=Zπ/20sin1
αiϕdϕandαi=β∗i
β∗i−1.WesimplywriteI1forI1,α.Thegoalofthissectionistoprovethefollowingresult.Proposition4.1.Wehavesupk≤nSk(˜f−˜R)
n1/α→0,inprobability,asngoestoinﬁnity.Thispropositionwillcomefromthefollowinglemmaswhichrequiresomenotations.Fixasmallδ>0,andsplitMiaccordingtothelow,intermediate,andhighregionsoftheindexmforthesetsMi,m:={x∈Mi:R(x)=m}:MLi:=∪m<δn1
αMi,m,MIi:=∪δn1
α≤m<1
δn1
αMi,mandMHi:=∪m≥1
δn1
αMi,m(4.3)10
whichalldependimplicitlyonnandδ.Wealsodenote˜fi,ζ:=˜f·1Mζi−˜µ(˜f·1Mζi),withζ∈{L,I,H}.Similarlyto[JZ18,Lemma4.3],˜fiisessentiallydeterminedby˜fi,I.Indeed,byconsideringtheinducedsystemrestrictedtoeachMi,theproofofLemma4.3in[JZ18]givesusthefollowingresult:Lemma4.2(Vanishingoftruncatedportions).Foreveryi∈J,wehave,inprobability,limδ→0supk≤nSk1MHi n1
αi=0.(4.4)Moreover,thereexistC00>0,θ∈(0,1)andκ>0suchthat,foreveryi∈J∗andeveryk≥1,wehaveCov˜fi,L,˜fi,L◦Fk≤C00δκθkn2
αi−1−κ.(4.5)Forζ∈{L,I,H},setRi,ζ:=R|MζianddeﬁneEi,ζas:Ei,ζ:=˜fi,ζ−If,i
I1(Ri,ζ(x)−˜µ(Ri,ζ)).(4.6)Againwenotethatthetruncationtotheintermediateregion,denotedbyIinthesubscripts/superscript,implicitlydependsonn.Similarto[JZ18],Lemma4.4andProposition4.2in[JZ18]impliesthatfisγ-Hölderasthefollowinglemmaindicates.Lemma4.3.ThereexistC=C(δ)>0andϑ∈(0,1)suchthat|Ei,I|≤Ckfk∞|Ri,I−˜µ(Ri,I)|1−γ
β∗i−1onMIiandCov Ei,I,Ei,I◦Fk≤Cn2
αi−1−1
αi(αi+1)ϑkTheproofoftheabovelemmaissimilartoargumentsin[JZ18];weindicatewhichestimatesintheproofneedupdatesand/oradjustmentsattheendofAppendixA.2.
ProofofProposition4.1.Lemma4.3impliesthatthereexistsK>0suchthatforallm,˜µ((SmEi,I)2)=Var(SmEi,I)≤2m−1Xk=0(m−k)Cov(Ei,I,Ei,I◦Fk)≤Kmn2
αi−1−1
αi(αi+1)Theexponentofncanberewritten−α2
i−αi−1
αi(αi+1)≤−α2−α−1
α(α+1)sinceαi≥α.So,by[Bil68,Exercise12.5](withγ=2,α=1there,seealso[Ser70,Eq.(2.3)]),forallm,˜µsupk=1,...,m|SkEi,I|2≤(log24m)2Kmn−α2−α−1
α(α+1).11
thereexistK0=K0(δ)>0andε>0suchthat,foralln≥0,˜µsupk=1,...,n|SkEi,I|2≤K0n2
α−ε.(4.7)Proceedinganalogously,duetoLemma4.2appliedto˜f−˜Rinsteadof˜f,thereexistK00>0andκ>0suchthatforalln≥0,˜µsupk=1,...,n|SkEi,L|2≤K00δκn2
α−ε.Fixε>0.DuetotheMarkovinequality,wecanﬁndδ>0smallenoughsothatforalln˜µ supk≤n
Sk Pi∈JEi,L

n1
α>ε
4!<ε
4.Also,by(4.4),wecanchooseδ>0smallenoughsothatforalln˜µ supk≤n
Sk Pi∈JEi,H

n1
α>ε
4!<ε
4.Now,dueto(4.7)andMarkov’sinequality,fortheminimumofthesechoicesofδ,andfornlargeenough,˜µ supk≤n
Sk Pi∈JEi,I

n1/α>ε
2!<ε
2.
4.2Convergenceindistribution
Duetothesimpliﬁcationintheprevioussubsection,itremainsnowtoprovetheconvergenceindistributionasn→∞forthefollowingfunctionsoft(intheJ1-metric)Zn(t):=S[nt]Pi∈JAi(Ri−˜µ(Ri))
n1/αtnforanyﬁxedrealnumbers{Ai,i∈J∗}.Tothisend,weconsiderthefamilyofpointprocesses(Nn)non(0,+∞)×(R\{0})givenbyNn:=nXj=1δj n,Zj−1
n1/α,withZj:="Xi∈JAi(Ri−˜µ(Ri))#◦Fj,andweemploystandardmethodsforprovingfunctionallimittheoremsfrom[DR78,Sec.4].Herewestateaversionfrom[TK10,Thm1.2].12
Proposition4.4.Assumethefollowingtwoconditions.ConditionI.(Pointprocessconvergence).Thesequenceofpointprocesses(Nn)nconvergesindistributiontoaPoissonpointprocessNon(0,+∞)×(R\{0})withintensityγhavingdensityψwithrespecttotheLebesguemeasureon(0,+∞)×(R\{0}),withψ(t,y)=Xi∈J∗2Iα1|Ai|α
β¯Cα−1iµ(M)|∂Q|1{yAi>0}α|y|−α−1.ConditionII(Vanishingsmallvalues).Forallγ>0lim→0limsupn→∞˜µ"max0≤k≤n

k−1Xj=0Zj n1/α·1|Zj|/n1/α<−k˜µZ1
n1/α·1|Z1|/n1/α<

>γ#=0.Then((Zn(t))t)n(andso((1
n1/αSbntc˜f)t)n)convergesindistribution(intheJ1-metric)toanα-stableLévymotion(Yt)tsuchthatY1hasthesamedistributionasZα,ξ,σwithσα:=Xi∈J∗2Iα1|Ai|α
β¯Cα−1iµ(M)|∂Q|andξ:=Pi∈J∗sign(Ai)2Iα1|Ai|α
β¯Cα−1iµ(M)|∂Q|
Pi∈J∗2Iα1|Ai|α
β¯Cα−1iµ(M)|∂Q|.4.3AkeyestimateforConditionI
InordertoprovetheconvergenceindistributionNn→N,weusetheKallenbergtheoremoramethodbasedonthistheorem.Akeyresultinthisstudyisthefollowinglemma,whichwasprovedasLemma2.2[JZ18]onlyforonespeciﬁcsymmetriccusp.Thegeneralizationofthisinourcontextisthekeyestimatethatshowsthatonlycuspsofmaximalﬂatnessarerelevant.Themodiﬁcationstotheproofgivenin[JZ18]isnontrivial,thusweprovidetheproofofthislemmainAppendixA.2.
Lemma4.5.Thereturntimefunctionsatisﬁes∀i∈J,limy→+∞yαi˜µ(x∈Mi:R(x)>y)=2Iαi1
β∗i¯Cαi−1iµ(M)|∂Q|,(4.8)withαi:=β∗i
β∗i−1.Sinceαi>αfori/∈J∗,thislemmaensuresinparticularthatlimy→+∞yα˜µ(x∈Mi:R(x)>y)>0⇔i∈J∗.13
Corollary4.6.Forany>0,limy→+∞yα˜µ(1/αZ−1>y)=limy→+∞yαXi∈J˜µ Mi∩{AiRi−XjAj˜µ(Rj)>−1
αy}!=limy→+∞yαXi∈J:Ai>0˜µ Mi∩(R>−1
αy
Ai+1
AiXjAj˜µ(Rj))!=Xi∈J∗:Ai>02Iα1Aα
i
β¯Cα−1iµ(M)|∂Q|.andanalogouslylimy→+∞yα˜µ(1/αZ−1<−y)=Xi∈J∗:Ai<02Iα1|Ai|α
β¯Cα−1iµ(M)|∂Q|.4.4ConditionI:Pointprocessconvergence
InordertoprovetheconvergenceindistributionNn→N,duetotheKallenbergtheorem[Kal73](seealso[Res87,Prop.3.22]),itisenoughtoprovethatlimn→+∞˜µ(Nn(R))=η(R)(4.9)andlimn→+∞˜µ(Nn(R)=0)=e−η(R),(4.10)foreveryRoftheformR=m[i=1(ai,bi)×Ici,c0
i,withIc,c0=(−∞,−c)∪(c0,+∞),0<ai<bi<1andci,c0
i>0foreveryi.InR\{0},weﬁxasubsetI=Ic,c0withc,c0>0.OurcorrelationboundswilldependonsetsoftheformDn,j:={x∈M:1
α√
n(Rj−c)◦Fj(x)∈I}.Lemma4.7(Exponentialdecayofcorrelationsforq-pointmarginals,[CZ09,JZ18]).ForeveryI,thereisaconstantC>0andθ∈(0,1)suchthat˜µ(Dcn,1∩···∩Dcn,q∩Dcn,q+k+1∩···∩Dcn,2q+k)− ˜µ(Dcn,1∩···∩Dcn,q)2≤Cθk(4.11)forallk,n,q∈Nsatisfying2q+k≤n.Also,thereexistsθ0>0suchthatforall1≤i<j≤n˜µ(Dn,i∩Dn,j)≤o1
n1+θ0.(4.12)Theﬁrstpartabove,(4.11)followsfromTheorem4in[CZ09]whichcoversthesettingwearein.For(4.12),theproofofLemma3.2in[JZ18]canbeusedbyreplacingLemma2.1therewiththeestimatesweprovideinPropositionintheappendix.14
4.4.1Proofof(4.9)Letm≥1,andrealnumbersa1,b1,c1,c0
1,...,am,bm,cm,c0
mbesuchthat,foreveryi,0<ai<biandci<c0
iwithcic0
i>0.Weassumewithoutanylossofgeneralitythatbi≤ai+1.LetR:=Sm i=1(ai,bi)×Ici,c0
i,sothat˜µ(Nn(R))=mXi=1Xj:nai<j<nbiXk∈J˜µ Mk;(AkRk−XiAi˜µ(Ri))◦Fj∈Icin1
α,c0
in1
α!.DuetoCorollary4.6,˜µ(Nn(R))∼2Iα1
β¯Cα−1iµ(M)|∂Q|mXi=1Xk∈J∗(bi−ai)1{Akci>0}|Ak|α.Therefore,wehaveproved(4.9)withηofdensityasinProposition4.4.4.4.2Proofof(4.10).Inordertoeasethenotationbelow,wewillproveonlythecasewherem=1,a=0andb=1.InthisspecialcasewecanconsiderthecanonicalprojectionofNnontoitssecondargument,i.e.,ontoanempiricalmeasureonR\{0}.LetuscallthisprojectionˆNn,andwereplacethesetRwithIasdeﬁnedaboveLemma4.7.Inthegeneralcase,onesimplyneedstowriteRintermsofaunionofdisjointrectangles,andconsidereachofrectanglesintheunionseparately;alsotheprojectionontothesecondcoordinateforeachoftheserectanglesmustbescaledbytheLebesguemeasureofthetimeinterval.Asin[JZ18,Prop.3.4]weuseBernstein’sblockmethod[Ber27].Condition(4.10)wouldfollowiftherandomvariables{Zj}wereindependentsincethenwehaveforZj=(R−˜µ(R))Pi∈JAi1Mi◦Fj,thesplittingintoaproductmeasure˜µ(ˆNn(I)=0)=nYj=1˜µ n−1/αZj6∈Iwhicheasilyimpliestherequiredconvergence.Sincewedonothaveindependence,weinsteaduseasymptoticindependencetogetherwithBernstein’sblockmethodtocontrolthedependencebetweenthe{Zj}.Choose0<v<w<θ0,whereθ0isasin(4.12),andforanyn≥1,divide{1,...,n}intoasequenceofpairsofalternatingbigintervals(blocks)oflength[nw]andsmallblocksoflength[nv].ThenumberofpairsofbigandsmallblocksisB=[n/([nv]+[nw])]sothatlimn→∞B/n1−w=1.TheremaybealeftoverpartialblockLintheendwhichisnegligiblesinceXj∈L˜µ(n−1/αZj∈I)≤Cnw n.15
Thuswemayhenceforthassumen=([nv]+[nw])B.WedenotebyBkandSkfork=1,···,B,theelementsof{1,...,n}inkthbigblockandsmallblock,respectively.LetYn,k=Xj∈Bk1{n−1/αZj∈I},Vn,k=Xj∈Sk1{n−1/αZj∈I}.Foreachn,both{Yn,k}and{Vn,k}aresequencesofidenticallydistributedrandomvariables.LetS0n=PB
k=1Yn,kandS00n=PB
k=1Vn,k.SimilartotheargumentforL,wehave˜µ(S00n)≤CBnv n≤C1
nw−v,(4.13)sothatwecanignoresmallblocks.Thus,itisenoughtoshowthat˜µ(S0n=0)convergestoexp(−λ(I)).Sincebigblocksareseparatedbysmallblocks,wecanpeeloffonefactoratatimeintheproductQ1{Yn,k=0}asfollows:˜µ(S0n=0)=˜µ(Yn,1=0,Yn,2=0,...,Yn,B=0)=˜µ BYk=11Yn,k=0!≤˜µ B−1Yk=11Yn,k=0!˜µ(1Yn,B=0)+Cθnv≤ ˜µ B−2Y
k=11Yn,k=0!˜µ(1Yn,B−1=0)+Cθnv!·˜µ(1Yn,B=0)+Cθnv.whereweused(4.11)inthelasttwolines.Repeatingthis,weobtainforsomeθ1∈(θ,1),˜µ(S0n=0)=(˜µ(Yn,1=0))B+O(θnv1).(4.14)Itremainstoestimate˜µ(Yn,1=0).˜µ(Yn,1=0)≤1−nwXj=1 ˜µ(n−1/αZj∈I)−nwXk=1˜µ({n−1/αZj∈I}∩{n−1/αZk∈I})!≤1−nwXj=1˜µ(n−1/αZj∈I)−o1
n≤1−nw˜µ(n−1/αZ1∈I)+nwo1
nwherethesecondinequalityfollowsfrom(4.12)sincew<θ0.Aneveneasierlowerboundisgivenby˜µ(Yn,1=0)≥1−nwXj=1˜µ(n−1/αZj∈I).16
Puttingthingstogetherwehave˜µ(S0n=0)= 1−[nw]˜µ(n−1/αZ1∈I)+o(nw−1)B+O(θnv1),whichiswhatweneedsinceB/n1−w→1andn˜µ(n−1/αZ1∈I)→λ(I).4.5ConditionI:alternativeapproach
Itisenoughtoprovetheconvergenceof(Nn)ntoNasapointprocesson(0,+∞)×(R\[−a,a])foreverya>0.Weﬁxa>0andsetA:={|Z−1|>a−1/α}.DuetoCorollary4.6,˜µ(A)∼η0(R\[−a,a])withη0themeasureonRwithdensityψ(1,·)withrespecttotheLebesguemeasure.Theconvergenceof(Nn)ntoNon(0,+∞)×(R\[−a,a])followsdirectlyfromthefollowingresult.Lemma4.8.Foreverya>0,thefamilyofpointprocessesPj≥1:Tjx∈Aδ(j˜µ(A),1/αZj−1)convergesindistribution,asε→0,toaPoissonprocessofintensityγ/η0(R\[−a,a])withrespecttotheLebesguemeasureon(0,+∞)×(R\[−a,a]).Proof.Wewillapply[PS18,Theorem2.1](whichusestheKallenbergtheorem)withAasabove,withH:A→V:=R\[−a,a]givenbyH(x):=1/αZ−1(x),withm:=η0(·\[−a,a])
η0(R\[−a,a])andwithW:={(c,c0);a<c<c0<∞}∪{(−c0,−c);a<c<c0<∞}.Toapply[PS18,Theorem2.1],wehavetoproveﬁrstthat˜µ(H−1(·)|A)convergesindistributiontom(thisisensuredbyCorollary4.6)andsecondthat,foreveryK≥1andeverychoiceofintervalsW1,...,WK∈W,∆,1=o(µ(A))(i.e.ino()),(4.15)with∆,n:=supA,B⊂A:A∈G,B∈σ(Sj≥nF−j(G))|˜µ(B∩A)−˜µ(A)˜µ(B)|,andwithG:=G,W1,...,WK:={H−1(Wi);i=1,...,K}.Asintheproofof[PS18,Proposition4.2],thefactthat∆,1=o()willfollowfromthefollowinglemmas.SetτA:=min{n≥1:Fn(·)∈A}.Wewilluseageneralargumentgivenbythenextgenerallemmaanditsgeneralcorollary.Lemma4.9(Generalresult).Foranypositiveintegerpandany>0,∆,1≤∆,p+1+˜µ(A)(˜µ(τA≤p|A)+˜µ(τA≤p)).17
Proof.LetA∈GandB∈σSj≥1H−1(G).Notethatthereexistsafunctiong:({0,1}K)N→{0,1}suchthat1B=g(X1,...),whereXi=1H−1(Wj))j=1,...,K◦Fi.SetCsuchthat1C:=g(0,...,0,Xp+1,...)∈σ [j≥p+1F−j(G)!(4.16)andobservethat|1B−1C|≤1{τA≤p}sothat|Cov(1A,1B)−Cov(1A,1C)|≤˜µ(A,τA≤p)+˜µ(A)˜µ(τA≤p).(4.17)Finally,dueto(4.16),|Cov(1A,1C)|≤∆,p+1.Thiscombinedwith(4.17)andA⊂Aendstheproofofthelemma.
Corollary4.10(Generalresultwhenµ(A)=O()).Ifthereexistsafamilyofpositiveinteger(p)suchthatp=o(−1)and˜µ(τA≤p|A)=o(1),then∆,1≤∆,p+1+o().Proof.ThiscorollarycomesdirectlyfromLemma4.9combinedwithµ(A)=O()and˜µ(τA≤p)=˜µ p[k=1F−k(A)!≤pXk=1˜µ(F−k(A))=p˜µ(A)=o(1).
Weﬁxθ∈(0,θ0)withθ0=εasin(4.12)andsetp=−θsothatp−1andLemma4.11.Wehave˜µ(τA≤p|A)=o(1).Proof.Dueto(4.12),˜µ(τA≤p|A)≤pXk=1˜µ A∩F−kA
˜µ(A)=Op1+θ0
=O −θ+θ0=o(1).
Lemma4.12(Decorrelation).∆,p+1=o().Proof.LetA∈GandB∈σSj≥p+1F−jG.ObservethatA∈GisaﬁniteunionoflevelsetsofR◦F−1intersectedbysomeF(Mi).ThereforeAcanbesmoothlyfoliatedbyaunionofunstablecurves.Moreover,sinceF−j(G)canbesmoothlyfoliatedbystablecurves,thesetBcanberewrittenB=F−pB0withB0smoothlyfoliatedbyaunionofstablecurves.Therefore,dueto[CZ09,Theorem3],thereexistsz∈(0,1)andC0>0suchthat∆,p+1≤C0zp.NotethatC0doesnotdependonAandB,i.e.,itisauniformconstant.
Corollary4.10combinedwithLemmas4.11and4.12ensures(4.15)(foreverya>0,foreveryK≥1andeveryW1,...,WK∈W)andsoLemma4.8.
18
4.6ConditionII:Vanishingsmallvalues
RecallthatZi:="Xj∈JAj(Rj−˜µ(Rj))#◦Fi.ConditionIIwillfollowimmediatelyfrom[Bil99,Thm10.1]togetherwithLemma4.14below.Given∈(0,1)andn≥1,wesetUm(x)=Um,,n(x):=m−1Xi=0(Vi(x)−˜µ(V0)),withVi=Vi,,n:=Zi(x)
n1/α·1{|Zi|/n1/α<}.WewillalsosetWi:=Vi−˜µ(Vi).InordertoproveLemma4.14,weneedanestimateofthevarianceofUm.Lemma4.13.Thereexists˜C>0andθ2>0suchthat,forevery∈(0,1)andforallnlargeenough(moreprecisely,thereexistsnsuchthatforalln>n),Xk≥0|Cov˜µ(V0,Vk)|=Xk≥0|˜µ(W0Wk)|≤˜C 2−αn−1+θ2n−1−θ2.Proof.Dueto[CZ09,Theorem3],|˜µ(W0Wk)|≤C2θk1.(4.18)Thisestimatewillbeusefulforbigk(k≥2logn/|logθ1|).Fork=0,duetoLemma4.5,thereexistsc0>0suchthat˜µ(W20)≤˜µ V20=Z20˜µ(Z0/n1
α)2>tdt≤c0Z20t−α
2n−1dt=c02−αn−1
1−α
2.(4.19)Fixγ∈(0,2
α−1).Noticingthat|Wk|≤nγ−1
α+max|Wk|−nγ−1
α,0weobtainthat|˜µ(W0Wk)|≤˜µ

(nγ−1
α+|W0|−nγ−1
α)Wk

≤nγ−1
α˜µ(|Wk|)+˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0+nγ−1
α≤2nγ−1
α˜µ(|W0|)+˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0.(4.20)DuetothedeﬁnitionofWj,VjandZj,wehavenγ−1
α˜µ(|W0|)≤2nγ−1
α˜µ(|V0|)≤4n−2
α+γK0˜µ(R),(4.21)19
whereK0:=maxi∈J|Ai|.Moreover,˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0=Z(nγ−1
α,3)2˜µ(|W0|>r,|Wk|>s)drds≤Z(nγ−1
α,3)2˜µ|V0|>r−2K0˜µ(R)
n1
α,|Vk|>s−2K0˜µ(R)
n1
αdrds,since|˜µ(V0)|≤2K0˜µ(R)
n1
α.Assumingandnaresuchthat2K0˜µ(R)
n1
α<nγ−1
α
2,weget˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0≤Z(nγ−1
α,3)2˜µ(|V0|>r/2,|Vk|>s/2)drds≤Z(nγ−1
α,3)2˜µR+˜µ(R)>n1
αr/(2K0),R◦Fk+˜µ(R)>n1
αs/(2K0)drds.≤Z(nγ−1
α,3)2˜µR>n1
αr/(4K0),R◦Fk>n1
αs/(4K0)drds.Now,usingtheinequality˜µ R>a,R◦Fk>b≤min ˜µ R>min(a,b),R◦Fk>min(a,b),˜µ(R>max(a,b)),weobtaintheexistenceofK00>0suchthat,forandnasaboveandforeveryk∈[1,(nγ)α],˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0≤Z(nγ−1
α,)2K00min(min(r,s))−α(1+θ0)n−(1+θ0),(nmax(r,s)α)−1drds,using(4.12)(appliedwithnof(4.12)equaltobn(min(r,s))αc≥b(nγ)αc≥kandI=Ic,c0withc=c0=1/(2K0))fortheﬁrstterminthelastlineandLemma4.5forthesecondterminthelastline.Therefore˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0≤2K00Znγ−1
α<r<s<min (nrα)−1−θ0,(nsα)−1drds.Weassumefromnowonwithoutlossofgenerality(uptorestrictingthevalueofθ0ifnecessary)that(α−1)(1+θ0)<1.Observethat(nrα)−1−θ0<(nsα)−1happensifandonlyifs<nθ0
αr1+θ0andthat(nsα)−1<(nrα)−1−θ020
happensifandonlyifr<n−θ0
α(1+θ0)s1
1+θ0,whichleadsto˜µmax|W0|−nγ−1
α,0max|Wk|−nγ−1
α,0≤2K00Z0(nrα)−1−θ0nθ0
αr1+θ0dr+Z0(nsα)−1n−θ0
α(1+θ0)s1
1+θ0ds≤2K00n−1−θ0+θ0
αZ0r−(α−1)(1+θ0)dr+n−1−θ0
α(1+θ0)Z0s1
1+θ0−αds=O1−(α−1)(1+θ0)n−1−θ0+θ0
α+1−(α−1)(1+θ0)
1+θ0n−1−θ0
α(1+θ0)=Oθ00n−1−θ00,(4.22)withθ00>0.Fix>0andconsidernlargeenoughsothatK0˜µ(R)
n1
α<nγ−1
α
2and2logn/|logθ1|≤(nγ)α.Puttingtogetherestimate(4.18)(forthesumoverk≥2logn/|logθ1|),combinedwith(4.19),(4.20),(4.21)and(4.22)(forthesumoverk<2logn/|logθ1|),weobtainthatthereexists˜K>0suchthatforevery>0,foreverynlargeenough,Xk≥0|Cov˜µ(V0,Vk)|≤˜K2n−2+2−αn−1+lognn−2
α+γ+θ00n−1−θ00,(4.23)whichendstheproofofthelemma.
Armedwiththeabovepreliminaryresult,wecannowprovealemmawhichcombinedwith[Bil99,Eq.(10.12)andThm10.1]willdirectlyimpliesConditionII.
Lemma4.14.Thereexists˜C0>0and˜θ,˜θ0>0suchthat,forevery∈(0,1)andforeverynlargeenough(moreprecisely,thereexistsnsuchthatforalln>n),andforevery0<m1<m2≤n,wehave˜µ U2m1(Um2−Um1)2≤˜C0m2
n1+˜θ4−2α+˜θ0n−˜θ0.Proof.Wehave˜µ U2m1(Um2−Um1)2=m1−1Xk1,k2=0m2−1Xk3,k4=m1˜µ 4Yj=1Wkj!≤4X0≤k1≤k2<m1<k3≤k4≤m2−1

˜µ 4Yj=1Wkj!

.(4.24)But,dueto[CZ09,Theorem3],foreveryj0∈{1,2,3},

˜µ 4Yj=1Wkj!

≤

˜µ Yj≤j0Wkj!˜µ Yj>j0Wkj!

+C04θkj0+1−kj01.(4.25)21
Noticethattheﬁrstpartoftheright-handsideof(4.25)vanishesunlessj0=2.Weapply(4.25)tothesumontherightsideof(4.24),bysplittingupthissumaccordingtothelargestofthethreegapsbetweenthefourindicesk1≤k2<k3≤k4.Weapply(4.25)withj0=1(respectively2or3)whenthelargestgapoccursbetweenk1andk2(respectivelyk2andk3ork3andk4),providingthisgapislargerthanorequaltok.Weconcludethat,foreveryk≥1,˜µ U2m1(Um2−Um1)2≤4
XEk˜µ

4Yj=1Wkj

!+ X0≤k1≤k2<m2

˜µ 2Yj=1Wkj!

!2+C04X`≥k`4θ`1
,whereEkisthesetof(k1,k2,k3,k4)satisfying0≤k1≤k2<m1≤k3≤k4<m2suchthatmax(k2−k1,k3−k2,k4−k3)≤k.Here`intheﬁnalsumrepresentsthelargestgap,i.e.,itismax(k2−k1,k3−k2,k4−k3),andthecoefﬁcient`4intheﬁnalsumcomesfromvaryingthekisubjecttotheconstraint0≤k1≤k2<m1≤k3≤k4<m2.Now,Lemma4.13ensuresthat,forevery>0,thereexistsnsuchthat,foreveryn≥nandevery0≤m1<m2,wehave˜µ U2m1(Um2−Um1)2≤4XEk˜µ

4Yj=1Wkj

!+C00(m2)2(4−2αn−2+2θ0n−2−2θ2)+4C04X`≥k`4θ`1.(4.26)Fixγ∈(0,2
α−1).Dueto(4.20),(4.21)and(4.22),thereexistsK00>0andθ00>0suchthat,foreveryandnsatisfyingK0˜µ(R)<nγ
2,forevery(k1,k2,k3,k4)suchthatmaxj=1,2,3(kj+1−kj)∈[1,(nγ)α],˜µ

4Yj=1Wkj

!≤42˜µ(|Wk1Wk2|)≤K00θ00n−1−θ00.(4.27)Fix>0andconsidernlargeenoughsothatK0˜µ(R)<nγ
2andsuchthatk=3logn/|logθ1|≤(nγ)α.Puttingtogether(4.26)with(4.27),weobtainthat˜µ U2m1(Um2−Um1)2≤C0004n−2+m2
2
n2(4−2α+n−2θ2)+(logn)4n−1−θ00≤C0004m2
n2+m2
n2(4−2α+n−2θ2)+(logn)4m2
n1+˜θn−θ00
2with˜θ:=min(1,θ00
2)sincen−1−θ00≤m1+˜θ2
n1+θ00≤ m2
n1+˜θn−θ00
2.ThisendstheproofofLemma4.14.
22
ProofofConditionII.DuetotheMarkovinequalityandtoLemma4.14,forevery0≤i≤j≤k≤n,˜µ(|Uj−Ui|∧|Uk−Uj|≥λ)≤λ−4E˜µ|Uj−Ui|4∧|Uk−Uj|4≤λ−4E˜µ|Uj−Ui|2|Uk−Uj|2≤λ−4E˜µ|Uj−i|2|Uk−i−Uj−i|2≤λ−4˜C0k−i n1+˜θ4−2α+˜θ0n−˜θ0.Therefore,duetoTheorem[Bil99,Theorem10.1](withβ=1andα=(1+˜θ)/2andu`=n−1(˜C0(4−2α+n−˜θ0))1/(1+˜θ)),thereexistsK>0dependingonlyon(α,β)suchthat˜µ(Ln≥λ)≤λ−4K˜C04−2α+˜θ0n−˜θ0,(4.28)withLn:=max0≤i≤j≤k≤n(|Uj−Ui|∧|Uk−Uj|).Asnoticedin[Bil99,(10.3)],maxk≤n|Uk|≤3Ln+maxk≤n|Vk−˜µ(V0)|.Therefore,foreveryγ>0,˜µmax0≤k≤n|Uk|≥γ≤˜µLn≥γ
6+˜µmaxk≤n|Vk−˜µ(V0)|≥γ
2≤K64γ−4˜C04−2α+˜θ0n−˜θ0+˜µmaxk≤n|Vk|≥γ
2−˜µ(V0)≤K64γ−4˜C04−2α+˜θ0n−˜θ0+˜µ≥γ
2−˜µ(|Z0|)
n1
α.Consequently,foreveryγ>0,limε→0limsupn→+∞˜µmax0≤k≤n|Uk|>γ≤limε→0hK64γ−4˜C04−2α+˜µ≥γ
2i=0.(4.29)
AAsymmetriccuspswithﬂatnessβInthissection,weconcentrateononecuspPiwithﬂatnessβfori∈J.Tosimplifynotation,weignoretheindexi,andconsiderherethecuspformedbyΓ,Γ0withacommontangentlineattheendpointP.Tosimplifynotation,weassumethattheﬂatpointPhascurvilinearabscissar=0andr0=|∂Q|.WechooseaCartesiancoordinatesystem(s,z)withoriginatP,andwiththehorizontals-axisbeingthetangentlinetotheboundaryofthebilliardtable.By(1.1),forsomesmall0>0,theboundarypairΓandΓ0adjacenttoPcan23
berepresentedinthe0−neighborhoodofthecuspPas:{(s,z0(s)),s∈[0,0]}∪{(s,−z1(s)),s∈[0,0]},withzi(s)=ciβ−1sβ+O(s2β−1),z0i(s)=cisβ−1+O(s2β−2),∀s∈[0,0],(A.1)wherec0,c1≥0notbothequalto0.Wealsoset¯c=c0+c1
2andα:=β/(β−1).Wewrite˜MforthesetofvectorsinMthatareinthecuspareaB(P)andsuchthatthepreviousreﬂectionoffof∂QwasoutsidethecuspareaB(P).FixN0.ForanyN≥N0,wedeﬁneMNtobethesetofpointsin˜MwhoseforwardtrajectoriesexplorethecuspatPoverNreﬂectionsoffoftheﬁrstcurve,eitherΓorΓ0,thatithits,beforeleavingthecusp.A.1Thecornerseries
Inthissubsection,weinvestigatethegeometryofcornerseries,whichcorrespondtocertainbilliardtrajectoriesenteringanasymmetriccuspofﬂatnessβ>0andexperiencealargenumberofrefectionstherebeforeexiting.ForalargeN≥N0,weconsideracornersequenceenteringthecuspatPwithNreﬂectionofftheﬁrstcurveithitsbeforeleavingthecusp(somakingeither2Nor2N+1reﬂectionsinthecusp).Weassumemoreover(uptopermutingtherolesplayedbyΓandbyΓ0)thattheﬁrstreﬂectionisonΓ.Wedenote(xn)n=((rn,ϕn))nastheconsecutivesequencecollidingwiththeboundaryΓ,and(x0
n)n=((r0n,ϕ0
n))nthesequenceonΓ0.Letsn(resp.s0
n)tobethes-coordinateofthebasepointofxn(resp.x0
n),forn=1,···,N.Bythesmoothnessoftheboundarycurves,|rn|=Zsn0p
1+(z00(u))2du=sn+O(s2β−1n),|r0n|=s0
n+O((s0
n)2β−1).(A.2)ToestimatethetaildistributionofµM(R≥n),wewillﬁxN0(asabove),andonlyconsiderthosecornerseries,suchthatN≥N0.Wewillalsoworkwithmoreconvenientcoordinates:γn=min(ϕn,π−ϕn),γ0n=min(ϕ0
n,π−ϕ0
n)andαn=arctan(z00(sn)),α0n=arctan(z01(s0
n)).Notethatby(A.1),thetangentvectorof∂Qat(sn,zi(sn))is(1,z0i(sn)),whichimplies,byTaylor-expanding,thatαn=arctan(z00(sn))=c0sβ−1n+O(s2β−2n),α0n=arctan(z01(s0
n))=c1(s0
n)β−1+O((s0
n)2β−2),(A.3)whereαn(resp.α0n)standsfortheanglein[0,π
2]ofthetangentlinetoΓat(sn,z0(sn))(resp.Γ0at(s0
n,−z1(s0
n)))madewiththehorizontalaxis,orequivalently,withthetangentlinethroughtheﬂatpointP.Notethatbothαnandγnarepositivefor1≤n≤N−1.Thesequences(αn)nand(α0n)naredecreasingandtakesmallvaluesifNislargeenough.Whiletheγnareinitiallysmall,theyslowlygrowtoaboutπ/2forn∼N/2,andthenagaindecreaseandgetsmall.Weusenotationsimilartothatof[CM07],anddeﬁneN2suchthatαN2:=min{αn:1≤n≤N}.Comparingthetrajectory(TN2−2jx)j=0,...,bN2/2cwiththeoutgoingtrajectory(TN2+2jx)j=0,...,b(N−N2)/2c,weconcludethatonlytwocasescanoccureithersN2−1<sN2+1<sN2−2<sN2+2<...orsN2+1<sN2−1<sN2+2<sN2−2<....24
Thisimpliesthat|N2−N/2|≤2.Wefurthersubdividethecornerseriesintothreesegments.Weﬁxasmallenough¯γandsetN1:=max{n≤N2:γn<¯γ},N3:=max{n≥N2:γn>¯γ}.Wecallthesegmenton[1,N1]theenteringperiodinthecornerseries,thesegment[N1+1;N3−1]theturningperiod,andthesegment[N3,N]theexitingperiod.ThesameargumentasaboveforN2showsthat|(N3−N2)−(N2−N1)|≤2sothat|N3+N1−N|≤|(N3−N2)−(N2−N1)|+|2N2−N|implies|N3+N1−N|≤6.Bythereversibilityofthebilliarddynamics,itisenoughtoconsidertheﬁrsthalfoftheseries,1≤n≤N2.Usingtherelationsabove,wecollectvariousestimatesinthefollowingpropositionforacornerseriesoflengthNgeneratedbyanyx∈MN.WeﬁrstdenoteanimportantfunctionHdeﬁnedonMinthecuspbyH(r,ϕ):=|r|βsinϕ.Recallthathere|r|representsthecurvilineardistanceon∂QbetweenthecuspPandthepointx=(r,ϕ)inthecuspthatweareinterestedin.Foreveryn=1,...,N,wesetHn:=H(rn,ϕn)=|rn|βsinϕnandH0n:=H(r0n,ϕ0
n)=|r0n|βsinϕ0
n.PropositionA.1.Assumeβ0=β1ormax(β0,β1)≥2β−1.Thefollowingaretrue:1(1)N1≈N2−N1≈N3−N2≈N−N3≈N,i.e.allthreesegmentsinthecornerserieshavelengthoforderN;(2)s2≈s0
1≈N−β
(2β−1)(β−1),sn≈s0
n≈n−1
β−1∼N−1
β−1,forn∈[N1,N2];(3)sn≈s0
n≈(nNβ
β−1)−1
2β−1,forn∈[2,N2];(4)γ1≈γ01=O(N−β
2β−1),γ2≈γ02≈N−β
2β−1;(5)vn≈v0n≈γn≈γ0n≈(nN−1)β
2β−1,forn∈[2,N2];(6)ForNsufﬁcientlylarge,thequantity{HN((rn,ϕn)),n=1,···,N−1}and{H0N((r0n,ϕ0
n)),n=1,···,N−1}arebothalmostinvariantalongacornerseriesoflengthN:∀n=1,...,N2,Hn=CN+O(s2β−1n)andH0n=C0N+O(s2β−1n).Moreprecisely,wewillseeinSectionA.2thatCN=¯c−αIα1N−β
β−1+ON−2β−1
β−1lnNandthatC0N=¯c−αIα1N−β
β−1+ON−2β−1
β−1lnNuniformlyinx∈MN.
1wherean≈bnmeansthat,fornlargeenough,thereexist˜c,˜C>0(independentofthecornerseriesandofn)suchthat˜can≤bn≤˜Can.25
Proof.Forn=1,···,N2,both{αn}and{α0n}aredecreasingsequences,and{γn}and{γ0n}areincreasingsequences.Observethatγn+1=γ0n+α0n+αn+1andγ0n=γn+αn+α0n.(A.4)Nowwedenotevn:=γn+αnandv0n:=γ0n+α0n.Observethatvnandv0ncorrespondtotheanglesbetweenthehorizontalline(i.e.thetangentlinetothecusp)andthereﬂecteddirectionscorrespondingto(rn,ϕn)and(r0n,ϕ0
n),respectively.Withthisnotation,(A.4)leadstovn+1=v0n+2αn+1,v0n=vn+2α0n.(A.5)Thisalsoimpliesthatvn=vn−1+2αn+2α0n−1,v0n=v0n−1+2αn+2α0n.(A.6)By(A.5)and(A.6),wehavev2>2α2,N2Xn=1αn≤vN2/2≤π/4(A.7)and
(n−1)2¯c(sβ−1n+O(s2β−2n))≤(n−1)(αn+α0n−1)≤nXk=2(αn+α0n−1)≤vn/2≤minπ
4sinvn,tanvn
2.(A.8)Thisimpliesinparticularthatsn=On−1
β−1.(A.9)Wedenoteτn:=z0(sn)+z1(s0
n)
sinvn,τ0n:=z0(sn+1)+z1(s0
n)
sinv0n.(A.10)Hereτnisthefreepathbetweentwocollisionsbasedatxnandx0
n,whileτ0nisthefreepathbetweentwocollisionsbasedatx0
nandxn+1.Observethatsn=s0
n−1−τ0n−1cosv0n−1=s0
n−1−z0(sn)+z1(s0
n−1)
tanv0n−1,s0
n=sn−τncosvn=sn−z0(sn)+z1(s0
n)
tanvn.(A.11)Thisimpliesthatsn+1−sn=Osβ
n tanvn(A.12)whichcombinedwith(A.8)impliesthatsn+1/sn=1+O(sβ−1n/tanvn)=1+O(1/n)andthatsn/sn+1=O(1).(A.13)26
Moreprecisely,weobtainthatsn+1−sn=−z0(sn+1)+z0(sn)
tanvn−2z1(s0
n)
tanvn+Os2β−1n
(sinvn)2(A.14)=−4¯csβ
n
βtanvn+Os2β−1n
(sinvn)2.(A.15)whereweusedthefactthat1
tanvn−1
tanv0n=O|vn−v0n|
sin2vncombinedwith(A.5)and(A.1).WedenoteAn:=sβ
nsinvn,n=1,···,N2.ThenextlemmatellsthatAn∼AN2isalmostinvariant,aslongasNislarge.Weset˜N2:=max{n<N2:vn<π
2−¯η1}(foranyﬁxed¯η1>0)andprovethefollowingestimates.LemmaA.2.An+1−An=Os3β−2n sinvn.(A.16)Moreover∀n=1,···,˜N2,An=A˜N2+O(s2β−1n),(A.17)andforeveryε∈(0,1),∀n=1,···,N2,An=A˜N2+O(s2β−1−εn).(A.18)Proof.WestartbywritingAn+1−An=(sβ
n+1−sβ
n)(sinvn+sinvn+1−sinvn)+sβ
n(sinvn+1−sinvn).Taylorexpansionsattheﬁrstordercombinedwith(A.15)andwith(A.6)leadtosβ
n+1−sβ
n=βsβ−1n(sn+1−sn)+O sβ−2n(sn+1−sn)2=−4¯cs2β−1n tanvn+Os3β−2n
(sinvn)2,andsinvn+1−sinvn=cosvn(vn+1−vn)+O vn(vn+1−vn)2=2cosvn(αn+1+α0n)+O s2β−2n=4¯csβ−1ncosvn+O s2β−2n/tanvn.ThereforeAn+1−An=Os3β−2n sinvn.(A.19)27
Using(A.15)and(A.13),itcomesthats2β−1n+1−s2β−1n=−4¯c(2β−1)
βs3β−2n tanvn+Os4β−3n sin2vn=−4¯c(2β−1)
βs3β−2n tanvn+Os2β−1n n2,(A.20)dueto(A.8),whichimpliesthat∀n=1,...,N2,N2Xk=ns3β−2k tanvk=O s2β−1n,andsothat∀n=1,...,˜N2,
An−A˜N2
≤˜N2−1Xk=n|Ak+1−Ak|=O s2β−1n.Letu∈(0,1)besuchthatβ+(β−1)(1−u)=2β−1−ε.Foreveryn=1,...,N2,An−AN2=O N2Xk=ns3β−2k sinvk!=O sβ+(β−1)(1−u)nN2Xk=ns(β−1)(1+u)k sinvk!=O sβ+(β−1)(1−u)nN2Xk=n(sinvk/k)1+u sinvk!=O(sβ+(β−1)(1−u)n),using(A.8).
WedeﬁneCN:=A˜N2=sβ˜N2sinv˜N2.(A.21)NowletusprovethekeyestimatesForeveryn=N1,...,N2,sinvn≈1andsosβ
n≈sβ
N1≈sβ
N2≈CN.Therefore∀n=N1,...,N2,sn≈C1
βN.(A.22)1≈vN2−vN1≈PN2k=N1sβ−1k≈(N2−N1)Cβ−1
βN.ThereforeN2−N1≈C−β−1
βN.(A.23)Letusseethatv1=O(sβ−11).Thisisobviouslytrueifs1≥0/2.Assumenowthats1<0/2,thenbydeﬁnitionofs1theincidentlineat(s1,z(z1))intersects[(2s1,z0(2s1)),(2s1,−z1(2s1))],whichimpliesthattan(v1−α1)≤O(sβ
1)/s1=O(sβ−11),andsov1=O(sβ−11).Thiscombinedwith(A.6)and(A.13)impliesthatv2≈sβ−12andsoCN≈sβ
2v2≈s2β−12,andsos2≈CN1
2β−1.(A.24)28
Foreveryn=2,...,N1,sn+1−sn≈sβ
n vn≈sβ
n
CNs−βn≈s2βn
CNandsos−2β+1n+1−s−2β+1n≈s−2βn(sn+1−sn)≈C−1N.Moreover,dueto(A.24),sinces−2β+12≈C−1N.Therefore∀n=2,...,N1,sn≈ s−2β+1n−1
2β−1≈n
CN−1
2β−1≈CN
n1
2β−1.(A.25)Dueto(A.25),sN1≈(CN/N1)1
2β−1.Thisand(A.22)leadsto(CN/N1)1
2β−1≈C1
βN.ThereforeN1≈C−β−1
βN.(A.26)Combiningthiswith(A.23),weobtainN1≈N2−N1≈N≈C−β−1
βN(A.27)andsoinparticularCN≈N−β
β−1.(A.28)Combining(A.28)withrespectively(A.25)and(A.22),weobtain∀n=2,...,N1,sn≈nNβ
β−1−1
2β−1.(A.29)∀n=N1,...,N2,sn≈N−1
β−1≈nNβ
β−1−1
2β−1,(A.30)sinceN1≈N2≈N.Thisfurtherimpliesthatforn∈[2,N2],wehavesβ−1n≈(αn+α0n−1)≈(nβ−1Nβ)−1
2β−1,γn≈vn≈(nN−1)β
2β−1.(A.31)Inparticularly,forn∈[N1,N2],usingtheabovefactthatN1≈N2≈N,wehave(αn+α0n−1)≈sβ−1n≈n−1≈N−1,vn≈1.(A.32)Nowby(A.17),weknowthat∀n=1,...,˜N2,An=sβ
nsinvn=CN+O(s2β−1n).29
Dueto(A.16)andto(A.32),∀n=˜N2,...,N2,AN2−An=O
N2Xk=˜N2s3β−2n
=O s1−βns3β−2n=O s2β−1n.Dueto(A.2),wehaveHn=|rn|βsinϕn=sβ
nsinγn+O(s3β−2nsinγn).(A.33)TheaboveestimationimpliesthatHn=An−sβ
n(sinγn−sinvn)+O(s3β−2nsinγn)=An+2sβ
nsinαn
2cos(γn+αn
2)+O(s3β−2nsinγn)=An+O(s2β−1n)=CN+O(s2β−1n),whereweusedαn=O(sβ−1n).WealsoobservethatH0n=|r0n|βsinϕ0
n=(s0
n)βsin(γ0n)+O(s2βnsinγn).(A.34)TheaboveestimationimpliesthatH0n=CN+O(s2β−1n).
Duetothetimereversibilityofbilliarddynamics,alltheasymptoticformulasobtainedfortheenteringperiodremainvalidfortheexitingperiod.
A.2ProofofLemma4.5
Recallthatα=β
β−1.Wesetν:=2|∂Q|µforthenonnormalizedmeasureonMsimplygivenbydν=sinϕdrdϕ.Write˜MforthesetofvectorsinMthatareinthecuspareaB(P)andsuchthatthepreviousreﬂectionoffQwasoutsidethecuspareaB(P).Thepurposeofthissectionistoprovethatlimy→∞yαν˜M:R>y=4
β¯c−1
β−1Iα1,(A.35)whichwillimplyLemma4.5sinceνisT-invariantandsince˜µ=ν|M
µ(M)|∂Q|.RecallthatwedenoteTnx=(rn,ϕn)andγn=min(ϕn,π−ϕn).DuetoPropositionA.1,forNlargeenoughthefollowingsequencesarealmostconstantforn=1,...,N:Hn=CN+O(s2β−1n),H0n=C0N+O((s0
n)2β−1),(A.36)whereCN=O(N−α),C0N=O(N−α).InordertoestimateCNandC0N,weuseanellipticintegralandintroducewn:=Zγn0(sinu)1−1
βdu,30
forn=1,...,N2.Thenwn+1−wn=Zγn+1γn(sinu)1−1
βdu=(sinγ∗n)1−1
β(γn+1−γn)(A.37)forsomeγ∗n∈[γn,γn+1].Dueto(A.33)and(A.36),wehavesinγn=Hn rβn=Hn sβ
n+O s2β−2n=CN
sβ
n+O(sβ−1n).(A.38)By(A.4)and(A.12),γn+1−γn=2α0n+αn+αn+1=4¯csβ−1n+O s2β−2n/tanvn.(A.39)andγn+1−γn=O(sβ−1n).Nowcombiningtheaboveandrecallingthatα=β/(β−1),werewrite(A.37)aswn+1−wn=(sinβ−1
βγn)(γn+1−γn)+Osin−1
βγn(γn+1−γn)2="CN
sβ
nβ−1
β+Osin−1
βγnsβ−1n#4¯csβ−1n+O s2β−2n/tanvn+Oγ−1
βns2β−2n=4¯c(CN)1
α+O(γ−1
βns2β−2n)=4¯c(CN)1
α+O(N−1n−1),(A.40)duetoPropositionA.1.Recalling(A.36),ifweuseadummyvariableandsum(A.40)from1ton,wegetwn=4¯cnC1
αN+O(lnn/N).(A.41)Inparticular,forn=N2=N/2+O(1)wegetZπ/20(sinu)1−1
βdu=ZγN20(sinu)1−1
βdu+Oγn−π
2=wN2+O(sβ−1N2)=wN2+O(N−1)=2¯cNC1
αN+O(lnN/N).ThusCNNα=(2¯c)−α Zπ/20(sinu)1−1
βdu!α+O(N−α−2lnN)=Iα1
(2¯c)α+O(lnN/N),(A.42)andsown=2I1n
N+O(lnn/N).31
Similarly,onecanshowthatC0NNα=(2¯c)−αIα1+O(lnN/N).(A.43)LetN0bethenumberofreﬂectionsinthecusp(bothonΓandonΓ0).NotethateitherN0=2NorN0=2N+1,sothatCNN0α=¯c−αIα1+O(lnN0/N0).andC0NN0α=¯c−αIα1+O(lnN0/N0).(A.44)Let˜Mmbethesetofpointsin˜MwhoseforwardtrajectoriesexplorethecuspatPduringmreﬂectionsoffΓ∪Γ0,beforeleavingthecusp.SinceνisT-invariant,ν˜Mm=1
mν m[n=1Tn˜Mm!.(A.45)Letybeaninteger.ObservethatXm0≥y1{N0≥m0}=Xm0≥yXm≥m01{N0=m}=Xm≥ymXm0=y1{N0=m}=Xm≥y(m−y+1)1{N0=m},andsoXm0≥yνx∈˜M:N0≥m0=Xm≥y(m−y+1)ν˜Mm=−yν(x∈˜M:N0≥y)+ν [m≥ym[n=0Tn˜Mm!.(A.46)But(A.36)and(A.42)togetherimplythatthesetSm≥ySm n=1Tm˜Mmcorrespondstothesetofpoints(r,ϕ)basedinthecuspareaB(P)suchthatH(r,ϕ)=|r|βsinϕ≤¯c−αIα1y−α+O(r2β−1)+O(y−1−αlny).ThereforethesetSm≥ySm n=1Tm˜Mmcorrespondstothesetofpoints(r,ϕ)basedinthecuspareaB(P)suchthatr≤¯c−α
βIα
β1y−α
β
sin1
βϕ+O y−α(1
β−1)
sin1
βϕ(lny/y+r2β−1)!.Thereforeν [m≥ym[n=0Tm˜Mm!=2Zπ0¯c−α
βIα
β1y−α
β
sin1
β−1ϕdϕ+O y−1=4¯c1−αIα1y1−α+O y−1.(A.47)32
ThistellsusthatXm≥y(m+1)ν(˜Mm)∼4¯c1−αIα1y−1
β−1asygoestoinﬁnity.UnfortunatelywecannotapplydirectlytheclassicalTauberiantheorembecausewedonotknowifmν(˜Mm)isdecreasingornot.Butwedonotwanttoestimateν(˜Mm),wejustwanttoestimatepy:=νx∈˜M:N0≥y=Xm≥yν(˜Mm).Tothisend,weadapttheTauberiantheoremargumentasfollows.Dueto(A.46)and(A.47)(consideringthedifferencebetweenthetermoforderdyeandthetermoforderd(1+ε)ye),weknowthat,foreveryε>0,dyepdye−d(1+ε)yepd(1+ε)ye+d(1+ε)ye−1Xm=dyepm∼4¯c1−αIα1 1−(1+ε)1−αy1−α,(A.48)asygoestoinﬁnity.Nowusingthefactthat(pm)misdecreasing,weobtainthatdye pdye−pd(1+ε)ye≤dyepdye−d(1+ε)yepd(1+ε)ye+d(1+ε)yeXm=dyepm≤d(1+ε)ye pdye−pd(1+ε)ye.(A.49)Fixanarbitraryϑ∈(0,1).Chooseε∈(0,ϑ)smallenoughsothat(α−1)ε(1−ϑ)≤1−(1+ε)1−α≤(α−1)ε(1+ϑ)andαε(1−ϑ)≤1−(1+ε)−α≤αε(1+ϑ).Dueto(A.48),foreveryylargeenough,weknowthat(α−1)ε(1−ϑ)24¯c1−αIα1y1−α≤dyepdye−d(1+ε)yepd(1+ε)ye+d(1+ε)yeXm=ypm≤(α−1)ε(1+ϑ)24¯c1−αIα1y1−α,andso,dueto(A.49),weobtainthat(α−1)ε(1−ϑ)2
1+ϑ4¯c1−αIα1y−α≤pdye−pd(1+ε)ye≤(α−1)ε(1+ϑ)24¯c1−αIα1y−α,foreveryylargeenough.Nowusingthefactthatpy=pdye=Pk≥0 pd(1+ε)kye−pd(1+ε)k+1ye,itfollowsthat,foreveryylargeenough,(α−1)ε(1−ϑ)2
1+ϑ4¯c1−αIα1Xk≥0(1+ε)−kαy−α≤py≤(α−1)ε(1+ϑ)24¯c1−αIα1Xk≥0(1+ε)−kαy−α33
whichcanberewritten(α−1)ε
1−(1+ε)−α(1−ϑ)2
1+ϑ4¯c1−αIα1y−α≤py≤(α−1)ε
1−(1+ε)−α(1+ϑ)24¯c1−αIα1y−α.Usingoursecondconditiononε,thisleadstoα−1
α(1−ϑ)2
(1+ϑ)24¯c1−αIα1y−α≤py≤α−1
α(1+ϑ)2
1−ϑ4¯c1−αIα1y−α,foreveryylargeenough.Sinceϑisarbitrary,wehaveprovedthatpy∼4α−1
α¯c1−αIα1y−α=4β−1¯c1−αIα1y−α,asygoestoinﬁnity,whichcompletestheproofofLemma4.5.ProofofLemma4.3.WealsonowobtaintheproofofLemma4.3byusing(A.41)and(A.44)inplaceof(6.8)and(6.9)intheproofofLemma4.4in[JZ18],andthenmakingtheappropriateobviousadjustments.
BSkorokhodJ1andM1topologiesAstrongerresultthanthelimittheoremin(2.4)isitsfunctionalversion,calledafunctionallimittheoremorweakinvarianceprinciple.TheWn’sareelementsintheSkorokhodspaceD[0,∞),i.e.,thespaceofallfunctionsφon[0,∞)thatareright-continuousandhaveleft-handlimitsφ(t−)foreveryt>0.WewillconsidertwodifferenttopologiesonD[0,∞).ThemostcommonlyusedtopologyintheliteratureisSkorokhod’sJ1-topologywhichisdescribedasfollows:ifφn,φ∈D[0,∞),thenφn→φintheJ1-topologyifandonlyifthereexistsasequenceof{λn}⊂A,suchthatsups|λn(s)−s|→0,sups≤m|φn(λn(s))−φ(s)|→0forallm∈N,whereAisthefamilyofstrictlyincreasing,continuousmappingsλof[0,∞]ontoitselfsuchthatλ(0)=0andλ(∞)=∞.Incontrast,theM1-topologyallowsafunctionφ1withajumpatttobeapproximatedarbitrarilywellbysomecontinuousφ2(withlargeslopeneart).ThemetricdM1thatgeneratestheM1-topologyonD[0,∞)isdeﬁnedusingcompletedgraphs.Forφ∈D[0,∞]thecompletedgraphofφisthesetΓ(φ):={(t,z)∈[0,∞)×R:z=λφ(t−)+(1−λ)φ(t)forsomeλ∈[0,1]}whereφ(t−)istheleftlimitofφatt.Besidesthepointsofthegraph{(t,φ(t)):t∈[0,∞)},thecompletedgraphofφalsocontainstheverticallinesegmentsjoining(t,φ(t))and(t,φ(t−))foralldiscontinuitypointstofφ.WedeﬁneanorderonthegraphΓ(φ)bysayingthat(t1,z1)≤(t2,z2)ifeither(i)t1<t2or(ii)t1=t2and|φ(t1−)−z1|≤|φ(t2−)−z2|.AparametricrepresentationofthecompletedgraphΓ(φ)isa34
continuousnondecreasingfunction(s,y)mapping[0,∞)ontoΓ(φ),withsbeingthetimecomponentandybeingthespatialcomponent.LetΛ(φ)denotethesetofparametricrepresentationsofthegraphΓ(φ).Forφ1,φ2∈D[0,∞)deﬁnedM1(φ1,φ2):=inf{ks1−s2k[0,∞)∨ku1−u2k[0,∞):(si,ui)∈Λ(φi),i=1,2},wherekφk[0,∞)=sup{|φ(t)|:t∈[0,∞)}.ThisdeﬁnitionintroducesdM1asametriconD[0,∞).TheinducedtopologyiscalledSkorokhod’sM1-topologyandisweakerthanthemorefrequentlyusedJ1-topology.Acknowledgements
TheresearchofH.ZhangwassupportedinpartbyNSFCAREERgrantDMS-1151762.TheresearchofP.JungwassupportedinpartbyNRFgrantN01170220fromtheRepublicofKorea.References
[AT92]F.AvramandM.S.Taqqu.Weakconvergenceofsumsofmovingaveragesintheα-stabledomainofattraction.
The
Annals of
Probability,pages483503,1992.[BCD11]PéterBálint,NikolaiChernov,andDmitryDolgopyat.Limittheoremsfordispersingbilliardswithcusps.
Communications in mathematical physics,308(2):479,2011.[Ber27]SergeBernstein.Surl’extensionduthéorèmelimiteducalculdesprobabilitésauxsommesdequantitésdépendantes.
Mathematische
Annalen,97(1):159,1927.[Bil68]P.Billingsley.
Convergenceofprobabilitymeasures.WileyNewYork,1968.[Bil99]P.Billingsley.
Convergence of
Probability
Measures.Wiley,NewYork,1999.[CM06]NikolaiChernovandRobertoMarkarian.
Chaotic billiards.AmericanMathematicalSociety,2006.[CM07]NChernovandRMarkarian.Dispersingbilliardswithcusps:slowdecayofcorrelations.
Communications in mathematical physics,270(3):727758,2007.[CZ05]NikolaiChernovandHong-KunZhang.Billiardswithpolynomialmixingrates.
Nonlinearity,18(4):1527,2005.[CZ09]NikolaiChernovandHong-KunZhang.Onstatisticalpropertiesofhyperbolicsystemswithsin-gularities.
Journal of
Statistical
Physics,136(4):615642,2009.[DR78]RichardDurrettandSidneyIResnick.Functionallimittheoremsfordependentvariables.
The
Annal

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

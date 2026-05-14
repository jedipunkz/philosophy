---
source: "https://arxiv.org/abs/1804.00043v4"
title: "Data-driven Coordination of Distributed Energy Resources for Active Power Provision"
author: "Hanchen Xu, Alejandro D. Domínguez-García, Peter W. Sauer"
year: "2018"
publication: "arXiv preprint / math.OC"
download: "https://arxiv.org/pdf/1804.00043v4"
pdf: "https://arxiv.org/pdf/1804.00043v4"
captured_at: "2026-05-09T13:11:13Z"
updated_at: "2026-05-09T13:11:13Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche will to power"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Data-driven Coordination of Distributed Energy Resources for Active Power Provision

- 著者: Hanchen Xu, Alejandro D. Domínguez-García, Peter W. Sauer
- 年: 2018
- 掲載情報: arXiv preprint / math.OC
- 情報源: [arxiv](https://arxiv.org/abs/1804.00043v4)
- ダウンロード: https://arxiv.org/pdf/1804.00043v4
- PDF: https://arxiv.org/pdf/1804.00043v4

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

In this paper, we propose a framework for coordinating distributed energy resources (DERs) connected to a power distribution system, the model of which is not completely known, so that they collectively provide a specified amount of active power to the bulk power system as quantified by the power exchange between both systems at the bus interconnecting them, while respecting distribution line capacity limits. The proposed framework consists of (i) a linear time-varying input-output (IO) system model that represents the relation between the DER active power injections (inputs), and the total active power exchanged between the distribution and bulk power systems (output); (ii) an estimator that aims to estimate the IO model parameters, and (iii) a controller that determines the optimal DER active power injections so the power exchanged between both systems equals to the specified amount at a minimum generating cost. We formulate the estimation problem as a quadratic program with box constraints and solve it using the projected gradient descent algorithm. To resolve the potential issue of collinearity in the measurements used by the estimator, we introduce random perturbations in the DER active power injections during the estimation process. Using the estimated IO model, the optimal DER coordination problem to be solved by the controller can be formulated as a convex optimization problem, which can be solved easily. The effectiveness of the framework is validated via numerical simulations using the IEEE 123-bus distribution test feeder.

## PDF Text

arXiv:1804.00043v4 [math.OC] 14 Feb 2019
1Data-drivenCoordinationofDistributedEnergyResourcesforActivePowerProvisionHanchenXu,StudentMember,IEEE,AlejandroD.Dom´ınguez-Garc´ıa,Member,IEEEandPeterW.Sauer,LifeFellow,IEEEAbstractInthispaper,weproposeaframeworkforcoordi-natingdistributedenergyresources(DERs)connectedtoapowerdistributionsystem,themodelofwhichisnotcompletelyknown,sothattheycollectivelyprovideaspeciﬁedamountofactivepowertothebulkpowersystem,whilerespectingdistributionlinecapacitylimits.Theproposedframeworkconsistsof(i)alineartime-varyinginput-output(IO)systemmodelthatrepresentstherelationbetweentheDERactivepowerinjections(inputs),andthetotalactivepowerexchangedbetweenthedistributionandbulkpowersystems(output);(ii)anestimatorthataims toestimatetheIOmodelparameters,and(iii)acontrollerthatdeterminestheoptimalDERactivepowerinjectionssothepowerexchangedbetweenbothsystemsequalstothespeciﬁedamountataminimumgeneratingcost.Weformulatetheestimationproblemasabox-constrainedquadraticprogramandsolveitusingtheprojectedgradient descentalgorithm.Toresolvethepotentialissueofcollinearityinthemeasurements,weintroducerandomperturbationsin theDERactivepowerinjectionsduringtheestimationprocess.UsingtheestimatedIOmodel,theoptimalDERcoordination problemtobesolvedbythecontrollercanbeformulatedas aconvexoptimizationproblem,whichcanbesolvedeasily.
TheeffectivenessoftheframeworkisvalidatedvianumericalsimulationsusingtheIEEE123-busdistributiontestfeeder.IndexTermsdata-driven,distributedenergyresource,co-ordination,activepowerprovision,projectedgradientdescent,parameterestimation.I.INTRODUCTIONINthemodernizationthatelectricpowersystemsare currentlyundergoing,onegoalistomassivelyintegratedistributedenergyresources(DERs)intopowerdistributionsystems[1].TheseDERs,whichincludedistributedgenerationresources,energystorage,demandresponseresources,and typicallyhavesmallcapacities,maybecoordinatedsoasto collectivelyprovidegridsupportservices,e.g.,reactivepowersupportforvoltagecontrol[2][4],andactivepowercontrolforfrequencyregulation[5][7].Inthispaper,wefocusontheproblemofcoordinatingtheresponseofasetofDERsinalossypowerdistributionsystem sothattheycollectivelyprovidesomeamountofactivepowertothebulkpowersystem,e.g.,intheformofdemandresponse.Speciﬁcally,theDERswillberequestedtocollectivelyprovideinreal-timeacertainamountofactivepoweratthebuswhereThisworkhasbeenfundedinpartbytheDepartmentofEnergyunderasubcontractfromtheUniversityofTennesseeKnoxville,theSiebelEnergyInstitute,andtheAdvancedResearchProjectsAgency-Energy(ARPA-E)withintheNODESprogram,underAwardDEAR0000695.TheauthorsarewiththeDepartmentofElectricalandComputerEngineer-ingattheUniversityofIllinoisatUrbana-Champaign,Urbana,IL61801,USA.Email:{hxu45,aledan,psauer}@illinois.edu.thepowerdistributionsystemisinterconnectedwiththebulkpowersystem.InorderfortheDERstofulﬁllsuchrequest,it isnecessarytodevelopappropriateschemesthatexplicitlytakeintoconsiderationthelossesincurred.Onewaytoincludethelossesintheproblemformulationistoutilizeapower-ﬂowlikemodelofthepowerdistributionsystemobtainedofﬂine.However,suchamodelmaynotbeavailableorifso,itmay notbeaccurate[8][10].Asanalternativetotheaforementionedmodel-basedap-proach,data-drivenapproacheshavebeendemonstratedtobeveryeffectiveinsuchsituationswheremodelsarenotreadilyavailable[9][16].Thefundamentalideabehinddata-drivenapproachesistodescribethesystembehaviorbyalinear time-varying(LTV)input-output(IO)model,andestimatetheparametersofthismodelviaregressionusingmeasurementsofpertinentvariables[11],[12].Manypreviousworkshaveap-plieddata-drivenapproachestopowersystemproblems,bothinasteady-statesetting[9],[10],[13],[14],andadynamicalsetting[15],[16].Forexample,in[9],theauthorsdevelopedadata-drivenframeworktoestimatelinearsensitivitydistribu-tionfactorssuchasinjectionshiftingfactors[13];theyfurtherproposedadataefﬁcientsparserepresentationtoestimatethesesensitivities[13].ThisframeworkwaslatertailoredtotheproblemofestimatingthepowerﬂowJacobian[10].In[14], theauthorsusedtheestimationframeworkproposedin[9]
tosolvethesecurityconstrainedeconomicdispatchproblem.Data-drivenapproacheshavealsobeenappliedsuccessfullytodeveloppowersystemsstabilizers[15],anddampingcontrols[16].Wereferinterestedreadersto[17]foranoverviewof data-drivenapproachesandtheirapplicationsinavarietyofotherareas.Yet,duetothecollinearityinthemeasurements[18],theregressionproblemmaybeill-conditioned,thusresultinginlargeestimationerrors[19].Thoughnumericalapproaches suchaslocallyweightedridgeregression[18]andnoiseassistedensembleregression[19]canbeusedtomitigatetheimpactsofcollinearity,thereisnotheoreticalguaranteeontheestimationerrors.Inthispaper,wepursuethedata-drivenapproachtodevelopaframeworkforcoordinatingtheresponseofasetofDERs.Theproposedframeworkconsistsofthree components,namely(i)amodelofthesystemdescribingthe relationbetweenthevariablesofinteresttotheproblem,i.e.,DERactivepowerinjectionsandpowerexchangedbetween thedistributionandbulkpowersystems,(ii)anestimator, whichprovidesestimatesoftheparametersthatpopulatethemodelin(i);and(iii)acontrollerthatusesthemodelin
(i)withtheparametersestimatedvia(ii)todeterminethe
2activepowerinjectionset-pointsoftheDERs.Speciﬁcally,anLTVIOmodelisadoptedasthesystemmodeltocapture therelationbetweentheDERactivepowerinjections(inputs),andthetotalactivepowerexchange(output).Theparametersinthismodelareestimatedbytheestimatorviathesolutionofabox-constrainedquadraticprogram, obtainedbyusingtheprojectedgradientdescent(PGD)algorithm.Inspiredbyideasinpowersystemmodelidentiﬁ-
cation[20],[21],weintroducerandomperturbationsinthe
DERactivepowerinjectionsduringtheestimationprocesstoresolvethepotentialissueofcollinearityinthemeasurementsusedbytheestimator.Weshowthattheestimationalgorithm convergesalmostsurely(a.s.)undersomemildconditions,i.e.,theestimatedparametersconvergetotheirtruevalues,and thetotalprovidedactivepoweralsoconvergestotherequiredamount.UsingtheestimatedIOmodel,theoptimalDER
coordinationproblem(ODCP)tobesolvedbythecontroller canbeformulatedasaconvexoptimizationproblem,which canbesolvedeasily.Themajorcontributionsofthispaper arethedata-drivencoordinationframework,thealgorithmtosolvetheestimationproblem,anditsconvergenceanalysis.Theremainderofthispaperisorganizedasfollows.Sec-tionIIdescribesthepowerdistributionsystemmodelandtheODCPofinterest.SectionIIIdescribesthecomponentsofthedata-drivenDERcoordinationframework.Adescriptionofthealgorithmusedintheframework,aswellasitsconvergence analysis,isprovidedinSectionIV.Theproposedframework isillustratedandvalidatedvianumericalsimulationsontheIEEE123-bustestfeederinSectionV.Concludingremarks arepresentedinSectionVI.II.PRELIMINARIESInthissection,weintroducethepowerdistributionsystemmodeladoptedinthiswork.WethendiscusstheDERcoordinationproblemofinterest.
A.PowerDistributionSystemModelConsiderapowerdistributionsystemrepresentedbyadirectedgraphthatconsistsofasetofbusesindexedbytheel-ementsinthesetN={0,1,···,N},andasetofdistributionlinesindexedbytheelementsinthesetL={1,2,···,L}.Eachline`∈Lisassociatedwithatuple(i,j)∈N×N,whereiisthesendingendandjisthereceivingendofaline,withthedirectionfromitojdeﬁnedtobepositive.Assumebus0correspondstoasubstationbus,whichistheonlyconnectionofthepowerdistributionsystemtothebulk powersystem.Further,assumethatbus0isaninﬁnitebusthatmaintainsaconstantvoltagemagnitude.Withoutloss ofgenerality,assumethereisatmostoneDERand/orload ateachbus,exceptbus0,whichdoesnothaveanyDERorload.LetNg={1,···,n}denotetheDERindexset.Throughoutthispaper,DERsareassumedtobecontrollable.
UncontrollableDERsaremodeledasnegativeloads.Letpd iandqdirespectivelydenotetheactiveandreactivepowerloadsatbusi,i∈N,anddeﬁnepd=[pd
1,···,pd
N],andqd=[qd1,···,qdN]>.Letpg iandqgirespectivelydenotetheactiveandreactivepowerinjectionsfromDERi,i∈Ng,anddeﬁnepg=[pg
1,···,pg n]>,andqg=[qg1,···,qgn]>.Letp g
iand pg irespectivelydenotetheminimumandmaximumactivepowerthatcanbeprovidedbyDERi,i∈Ng,anddeﬁnep g=[p g
1,···,p g
n]>,and pg=[
pg
1,···, pg n]>.Similarly,letq g
iand qg irespectivelydenotetheminimumandmaximumreactivepowerthatcanbeprovidedbyDERi,i∈Ng,anddeﬁneq g=[q g
1,···,q g
n]>,and qg=[
qg
1,···, qg n]>.Letpidenotetheactivepowerinjectionatbusi∈N,anddeﬁnep=[p1,···,pN]>;thenp=Cpg−pd,(1)whereC∈RN×nisthematrixthatmapstheDERindicestothebuses.Theentryattheithrow,jthcolumnofCis1ifDERjisatbusi.Let˜M=[Mi`]∈(N+1)×L,withMi`=1andMj`=−1ifline`startsfrombusiandendsatbusj,andallotherentriesequaltozero.LetMdenotethe(N×L)-dimensionalmatrixthatresultsfromremovingtheﬁrstrowin˜M.Weﬁrstmakethefollowingassumption:
Assumption1.Thepowerdistributionsystemisradial.UnderAssumption1,L=N,andMisinvertible.Letf`denotetheactivepowerthatﬂowsfromthesendingendtothe receivingendofline`∈L,anddeﬁnef=[f1,···,fL]>.Let f`denotemaximumpowerﬂowsonline`∈L,anddeﬁne f=[
f1,···, fL]>.Thenfcanbeapproximatelycomputedfrompasfollows:f≈M−1p=M−1(Cpg−pd).(2)WewouldliketopointoutthatAssumption1allowsustoapproximatelycomputetheactivepowerﬂowonthedistributionlines.Itis,however,notanecessaryconditionfortheanalysistobedoneintherestofthepaper.Letydenotetheactivepowerexchangedbetweenthedistributionandbulkpowersystemsviabus0,deﬁnedtobepositiveiftheﬂowisfromthesubstationtothebulkpower system.Conceptually,ycanberepresentedasafunctionofpg,qg,pd,qd.Notealsoqgistypicallysetaccordingtosomespeciﬁcreactivepowercontrolrulestoachievecertainobjectivessuchasconstantvoltagemagnitudeorconstant powerfactor[7],andthusisafunctionofpg,pd,qd.Then,ycanbewrittenasafunctionofpg,pd,qdasfollows:y=h(pg,pd,qd),(3)whereh(·)capturestheimpactsfromboththephysicallawsaswellasreactivepowercontrol.Weemphasizethatalthoughthevoltagecontrolproblemisnotexplicitlymodeledinthispaper,weassumecertainvoltagecontrolschemesexistinthepowerdistributionsystemsuchthatthevoltageproﬁlewill bemaintainedwithinanacceptablerange.Indeed,voltage controlschemesmayhaveasigniﬁcantimpactonh.Wereferinterestedreadersto[7]foradetailedanalysisontheimpactsofvoltagecontrolschemes.Notethattheexplicitformofhisdifﬁculttoobtain;however,wecanmakethefollowingassumptiononh:Assumption2.Thefunctionhisdifferentiableanditsﬁrstorderpartialderivativeswithrespecttopgbelongto[b
1, b1],
3whereb
1, b1aresomeknownconstants.Inaddition,∂h
∂pgisaLipschitzfunction,i.e.,thereexistsb2>0suchthat

∂h
∂pg

a−∂h
∂pg

b

≤b2ka−bk,wherea,b∈[p g, pg],andk·kdenotestheL2-norm.Assumption2impliesthat,forﬁxedloads,therateofchangeinyisboundedforboundedchangesintheDERactivepowerinjections.Inaddition,thetotalactivepowerprovidedtothebulkpowersystemwillincreasewhenmore activepowerisinjectedinthepowerdistributionsystem.Thisassumptionholdswhenthesystemisatanormaloperating conditionwithoutlinecongestions.
B.OptimalDERCoordinationProblemTheDERsinthedistributionsystemcancollectivelypro-videactivepowertothebulkpowersystemasquantiﬁedbythe powerexchangebetweenbothsystemsatthesubstationbus.
Forexample,theDERscanprovidedemandresponseservices orfrequencyregulationservicestothebulkpowersystem;inbothcases,theDERsneedtobecoordinatedinsuchawaythat thetotalactivepowerprovidedtothebulkpowersystem,y,trackssomepre-speciﬁedvalue,denotedbyy?.TheobjectiveoftheODCPistodeterminetheDERactivepowerinjections, pg,thatminimizesomecostfunction,e.g.,onethatreﬂectsthecostofactivepowerprovision,whilerespectingtothe followingconstraints:[C1.]theactivepowerexchangedbetweenthedistributionand bulkpowersystemsviabus0,y,trackssomepre-speciﬁedvaluey?;[C2.]theactivepowerinjectionfromeachDERi,i∈Ng,doesnotexceeditscorrespondingcapacitylimits,i.e., p
g≤pg≤
pg;[C3.]thepowerﬂowoneachline`,`∈L,doesnotexceeditsmaximumcapacity,i.e.,−
f≤f≤
f.NotethatwhileconstraintC2isahardconstraintthatcannotbeviolated,constraintC3maybeallowedtobeviolatedslightlyforashortperiod.TheODCPcanbeformulatedas thefollowingoptimizationproblem:minimizepg∈[p g, pg]c(pg),subjecttoh(pg,pd,qd)=y?,(4a)−
f≤M−1(Cpg−pd)≤
f,(4b)wherec(·)denotesthecostfunctionoftheactivepowerinjections.Thisproblemisdifﬁcult,however,whenthemodeldescribingthepowerexchangewiththebulkpowersystem, ascapturedbyh,isunknown.Inthispaper,wewillresorttoadata-drivenapproachtotacklethisproblem.III.DERCOORDINATIONFRAMEWORKInthissection,wedescribethebuildingblocksoftheproposedframework;namelyanLTVIOmodel,anestimator, andacontroller.
time controller solves ODCP
time iterations in estimation processestimator updates sensitivities
Fig.1.Timescaleseparationofactionstaken.
A.OverviewTheDERcoordinationframeworkconsistsofthreecom-ponents,namely(i)anLTCmodelofthesystemdescribing therelationbetweenyandu,(ii)anestimatorthatprovidesestimatesoftheparameterstheso-calledsensitivityvectorthatpopulatethemodelin(i);and(iii)acontrollerthatusesthemodelin(i)withtheparametersestimatedvia(ii)tosolvetheODCP.Thisframeworkworksontwotimescalesaslowone andafastone,asillustratedinFig.1.Ontheslowtimescale,thecontrollerdeterminestheDERactivepowerinjection set-pointsbysolvingtheODCPperiodically.Thesensitivityvectorisalsoupdatedbytheestimatorperiodically.However,duringeachupdateofthesensitivityvector,theestimator needstotakeactionsinseveraliterationsonafasttimescale.Sincethesensitivityvectormaynotchangesigniﬁcantlyinashorttime,itisusedintheODCPforseveraltimeinstants beforeitisupdatedagain.Beforeweproceedtopresenting thedetailedcomponentsintheproposedframework,wemake thefollowingassumption:
Assumption3.pdandqdareconstantduringtheestimationprocess;therefore,changesinythatoccuracrossiterationsintheestimationprocessdependonlyonchangesinpg.Remark1.Assumption3allowsustodeterminetheimpactsoftheDERactivepowerinjectionsontheoutput.Whentheload variabilityissigniﬁcantenoughsothatitcannotbeneglectedduringtheestimationprocess,itbecomesnecessarytomeasuretheloadsanddeterminetheirimpactsontheoutputaswell.
Thisisbeyondthescopeofthepaper,andthereforewewill leaveitasfuturework.
B.Input-OutputSystemModelFornotationalsimplicityinthelaterdevelopment,deﬁneu=pg,u
=p g, u=
pg,andπ=[(pd)>,(qd)>]>;then,(3)canbewrittenas:y=h(u,π).(5)Unlessotherwisenoted,throughoutthispaper,x[k]denotesthevaluethatsomevariablextakesatiterationk.Itfollowsfrom(5)andAssumption2thaty[k−1]=h(u[k−1],π)andy[k]=h(u[k],π).Then,bytheMeanValueTheorem,there
4existsak∈[0,1]and˜u[k]=aku[k]+(1−ak)u[k−1]suchthaty[k]−y[k−1]=h(u[k],π)−h(u[k−1],π)=φ[k]>(u[k]−u[k−1]),whereφ[k]>=[φ1[k],···,φn[k]]=∂h
∂u

˜u[k],1isreferredtoasthesensitivityvectoratiterationk.ItfollowsfromAssumption3thatφi[k]∈[b
1, b1],i=1,···,n.Therefore,atanyiterationk,(5)canbetransformedintothefollowingequivalentLTVIOmodel:y[k]=y[k−1]+φ[k]>(u[k]−u[k−1]).(6)C.EstimatorOnFastTimescaleAsillustratedinFig.1,theestimatorupdatesthesensitivityvectoracrossseveraliterationsonthefasttimescale.Atitera-tionk,theobjectiveoftheestimatoristoobtainanestimateofφ[k],denotedbyˆφ[k],usingmeasurementscollecteduptoiterationk,i.e.,y[k−1],u[k−1],y[k−2],u[k−2],···;weformulatethisestimationproblemasfollows:ˆφ[k]=argminˆφ∈BJe(ˆφ)=1
2(y[k−1]−ˆy[k−1])2,subjecttoˆy[k−1]=y[k−2]+ˆφ>(u[k−1]−u[k−2]),(7)whereB=[b
1, b1]n,Je(·)isthecostfunctionoftheestimator,andˆy[k−1]isthevalueofy[k−1]estimatedbytheIOmodelatiterationk.Essentially,(7)aimstoﬁndˆφthatminimizesthesquarederrorbetweentheestimatedvalueandthetrue valueofy.Then,ˆφ[k]isusedinthecontrollertodeterminethecontrolforthenexttimeinstant.Duringtheestimationprocess,itisstillnecessarytotracktheoutputtarget.Therefore,ateachiteration,thecontrolissetbasedonthesolutiontothefollowingproblem:u[k]=argminu∈UJc(u)=1
2(y?−ˆy[k])2,subjecttoˆy[k]=y[k−1]+ˆφ[k]>(u−u[k−1]),(8)whereU=[u
, u],andJc(·)isthecostfunction.Notethatˆφ[k]isusedin(8)topredictthevalueofy[k]foragivenu.DifferentfromtheODCP,theobjectiveofthecontrollerduringtheestimationprocessisthattheoutputtracksthetargetandtheremayexistmultiplesolutionstothisproblem.This objectiveischosensuchthattheDERactivepowerinjectionsbehaveinawaythatcanimprovetheestimationaccuracy,as willbeshownlaterinSectionIV.1Weadopttheconventionthatthepartialderivativeofascalarfunctionwithrespecttoavectorisarowvector.D.ControllerOnSlowTimescaleAsillustratedinFig.1,thecontrollersolvestheODCPtodeterminetheleast-costactivepowerset-pointsforDERsontheslowtimescale.Meanwhile,italsoforcestheDERsto injectrandomactivepowerperturbationsateachiterationonthefasttimescale.TheODCPtobesolvedbythecontroller isasfollows:minimize pg∈[p g, pg]c(pg),subjecttoy+ˆφ>(pg−˜pg)=y?,(9a)−
f≤M−1(Cpg−pd)≤
f,(9b)whereyistheoutputand˜pgistheDERactivepowerinjectionvectoratthebeginningofthecurrenttimeinstant,andˆφistheup-to-datesensitivityvector.Whencisaconvexfunction,(9)isaconvexproblemandthereforecanbesolvedeasilywith convergenceguarantees.Notethatthisformulationisobtainedbyreplacing(4a)in(4)bytheestimatedIOmodel.IV.ESTIMATIONALGORITHMANDITSCONVERGENCETheODCPcanbesolvedusingexistingalgorithmsforconvexoptimizationandthuswedonotdiscussitinmore detailshere.Inthissection,wefocusontheproblemfacedbytheestimatorandproposeaPGDbasedalgorithmtosolve it.Wethenprovideconvergenceresultsfortheproposed algorithm.
A.EstimationAlgorithmWeﬁrstdescribethebasicworkﬂowoftheproposedal-gorithm.Eachiterationconsistsofanestimationstepand acontrolstep.Atthebeginningofiterationk,y[k−1]isavailabletotheestimator,whichusesittoupdatetheestimateofthesensitivityvector.Theupdatedestimateofthesensitivityvector,ˆφ[k],isthenusedinthecontrollertodeterminethecontrol,u[k].Then,theDERsareinstructedtochangetheiractivepowerinjectionset-pointsbasedonu[k].Attimeinstantk+1,theestimationandcontroliterationsarerepeatedoncey[k]becomesavailable.Thesequentialprocessdescribedabove,whichhappensonthefasttimescale,isillustratedasfollows:
···u[k−1]→y[k−1]→ˆφ|
{z
}estimationstepcontrolstepz
}|
{[k]→u[k]→y[k]→ˆφ[k+1]···Problems(7)and(8)canbesolvedusingthePGDmethod(see,e.g.,[22]).LetPV1→V2denotetheprojectionoperatorfromavectorspaceV1toits(arbitrary)subspaceV2,i.e.,PV1→V2(v1)=argminv2∈V2kv2−v1k,wherev1∈V1.Foreaseofnotation,whenthevectorspacetowhichv1belongsisunambiguous,wesimplywritePV2(v1)insteadofPV1→V2(v1).Deﬁnethetrackingerroratiterationkase[k]=y[k]−y?.Inaddition,deﬁne∆y[k]=y[k]−y[k−1]and∆u[k]=
5
Algorithm1:EstimationAlgorithm
Input:y:outputy?:outputtrackingtargetδ:maximumallowedtrackingerrorˆφ0:initialestimateofsensitivityvectoru0:initialDERactivepowerinjectionset-pointOutput:u:DERactivepowerinjectionset-pointˆφ:estimateofsensitivityvectorInitialization:setˆφ[0]=ˆφ0,u[−1]=u[0]=u0,obtainmeasurementofy[−1],setk=1while|e[k]|>δdo obtainnewmeasurementofy[k−1]computee[k−1]=y[k−1]−y?compute∆y[k−1]=y[k−1]−y[k−2]compute∆u[k−1]=u[k−1]−u[k−2]updatethesensitivityvectorestimateaccordingtoˆφ[k]=PBˆφ[k−1]−αk∆u[k−1](∆u[k−1]>ˆφ[k−1]−∆y[k−1])updatethecontrolvector,u,accordingtou[k]=PUu[k−1]−βke[k−1]W[k]ˆφ[k]changeDERactivepowerinjectionstou[k]setk=k+1end u[k]−u[k−1].ThepartialderivativevectorofJe(ˆφ)withrespecttoˆφis∂Je(ˆφ)
∂ˆφ=∆u[k−1](∆u[k−1]>ˆφ−∆y[k−1]),(10)andthatofJc(u)withrespecttouis∂Jc(u)
∂u=ˆφ[k](ˆφ[k]>(u−u[k−1])+e[k−1]).(11)Insteadofsolvingboth(7)and(8)tocompletion,weiteratethePGDalgorithmthatwouldsolvethemforonestepat eachiteration.Speciﬁcally,atiterationk,weevaluatethenewgradientatˆφ[k−1]andu[k−1]anditerateonce.Thus,byusing(10)and(11),theupdaterulesforˆφandu,respectively,areˆφ[k]=PBˆφ[k−1]−αk∆u[k−1](∆u[k−1]>ˆφ[k−1]−∆y[k−1]),(12)u[k]=PUu[k−1]−βke[k−1]ˆφ[k],(13)whereαk>0andβk>0aretheestimationandcontrolstepsizesatiterationk.Inordertoresolvethepotentialissueofcollinearityinthemeasurementsusedbytheestimator,weintroducerandom perturbationsduringtheestimationprocess.DeﬁneW[k]=diag(w1[k],...,wn[k]),wherewi[k]’sareindependentran-domvariablesthatfollowaBernoullidistributionwitha probabilityparameterof0.5.Then,thecontrolupdaterulein(13)ismodiﬁed,resultingin:u[k]=PUu[k−1]−βke[k−1]W[k]ˆφ[k].(14)Intuitively,thismeansthatateachiteration,thecontrolofeachDERisupdatedwithaprobabilityof0.5.Therandomperturbationinthecontroliskeytoestablishconvergenceoftheparameterestimationprocess.Theestimationalgorithm,alongwithitsinitialization,issummarizedinAlgorithm1,whereu0isthevectorofDERactivepowerinjectionsatthebeginningofthetimeinstantatwhichtheestimationstarts andˆφ0istheup-to-dateestimateofthesensitivityvectoratthebeginningofthesametimeinstant.
B.ConvergenceAnalysisThemainconvergenceresultforthecontrolstepduringtheestimationprocessisstatednext.
Theorem1.Usingtheestimationupdaterulein(12)andthecontrolupdaterulein(14)withβk∈(
b
2
1,1
n b2
1),where0<<b
2
1
n b2
1isagivenparameter,thesystemattainsoneofthefollowingequilibria:1)e[k]convergesto0a.s.;2)e[k]convergestosomepositiveconstantandu[k]staysatu
;3)e[k]convergestosomenegativeconstantandu[k]staysat u.Inallcases,limk→∞∆u[k]=0n,where0n∈Rnisanall-zerosvector.Theorem1showssomethingintuitive,i.e.,thetrackingerrorwillbepositive(negative)iftherequestedactivepowerisless(more)thantheminimum(maximum)amountofactivepower theDERscanprovide;otherwise,thetrackingerrorgoesto zeroa.s.Wenotethathasadirectimpactontheconvergencerateofthecontrolalgorithm.Thisismoreobviousinadeterministicsetting,whenthecontrolupdaterulein(13)isusedinstead oftheonein(14).Aresultontheconvergencerateisgiven inthefollowingcorollary.
Corollary1.Assumeu[k]6=u andu[k]6=
u,∀k∈N.Usingtheestimationupdaterulein(12)andthecontrolupdaterulein(13)withβk∈(
b
2
1,1
n b2
1),where>0isagivenparameter,e[k]convergesto0ataratesmallerthat1−,i.e.,

e[k]
e[k−1]

<1−.WereferthereaderstoAppendixAfordetailedproofsoftheconvergenceresultsabove.Next,weestablishtheconvergenceoftheestimationstep.Deﬁnetheestimationerrorvectoratiterationkasε[k]=ˆφ[k]−φ[k].Sincebothˆφ[k]andφ[k]arebounded,ε[k]isalsobounded.Deﬁne∆φ[k]=φ[k]−φ[k−1].Theconvergenceresultfortheestimationstepisstatednext.
Theorem2.Usingtheestimationupdaterulein(12)andthecontrolupdaterule(14),withαk+1=2
k∆u[k]k2,βk∈(
nb
2
1,1
n b2
1),where0<<b
2
1
b2
1isagivenparameter,if
6
0
12345678
912141011131516171819202122232425262728293031323334353637383940
414243444546474849505152535455565758596061626364656667686970
717273747576
7778798081828485
838687888990919293949596979899101100102103104105106107108109110111112113114117115118119116121122120
DER
Fig.2.IEEE123-busdistributiontestfeeder.
u[k]∈(u
, u)ande[k]6=0,∀k∈N,thenkε[k]kconvergesto0a.s.Theintuitionisthattheestimationerrorgoestozeroifthesystemcanbecontinuouslyexcited(guaranteedbythe conditionu[k]∈(u
, u)ande[k]6=0,∀k∈N).WereferthereaderstoAppendixBfordetailedproofsoftheconvergenceresultabove.V.NUMERICALSIMULATIONSInthissection,weillustratetheapplicationoftheproposedDERcoordinationframeworkandvalidatethetheoretical resultspresentedearlier.Fromapracticalpointofview, thetimescaleseparationillustratedinFig.1iscriticalfortheapplicabilityoftheproposedframework.Speciﬁcally,theestimationprocessneedstobemuchfasterthanthetimescalegoverningtheloadchanges.TheDERs,whicharetypically powerelectronicsinterfaced,canrespondveryquickly,onatimescaleofmillisecondtosecond[23].Inthissimulation,wesetthedurationbetweentwoiterationstobe100millisecond.Wewillshowlaterthatunderthissetup,therequirementson thetimeseparationcanbereasonablymet.Amodiﬁedthree-phasebalancedIEEE123-busdistributiontestfeederfrom[24](seeFig.2fortheone-linediagram)is usedforallnumericalsimulations.Thisbalancedtestfeederhasatotalactivepowerloadof3000kW,andatotalreactivepowerloadof1575kVAr.DERsareaddedatbuses19,26,38,49,56,64,78,89,99.WeassumeeachDERcanoutput activepowerfrom0kWto100kW.Therefore,themaximumDERactivepowerinjectionsaccountsfor30%ofthenominalloads.Toillustratetheimpactsofreactivepowercontrol,weassumeallDERsoperateatunitypowerfactorexceptDERsat buses78and89,whichareassumedtohaveenoughreactive powercapacityandmaintainaconstantvoltagemagnitudeof
0.95p.u.Yet,wewouldliketoemphasizethattheproposedalgorithmisagnostictotheunderlyingreactivepowercontrolschemeandalsoworksunderotherreactivepowercontrol schemes.Inaddition,tovalidatetheeffectivenessofthe
0.02.04.06.08.010.0time(s)051015DERacitivepower(kW)19
26
3849
56
6478
89
99
Fig.3.DERactivepowerinjectionsforβk=0.02andy?=−3000kWinCaseI.(LegendsindicateDERbuses.)
0.02.04.06.08.010.0time(s)−600−400−2000trackingerror(kW)-3000kW
-2900kW-2800kW
-2700kW-2600kW
-2500kW
Fig.4.Trackingerrorforβk=0.02undervarioustrackingtargetsinCaseI.(Legendsindicatevaluesofy?.)proposedalgorithmunderdifferentoperatingconditionsofthepowerdistributionsystem,weassumetherearesomeuncontrollablerenewableenergyresourcesinthepowerdistributionsystem,whicharemodeledasnegativeloads.Theunderlying nonlinearpowerﬂowproblemissolvedusingMatpower[25].Inallsubsequentsimulations,wesetb
1=0.8, b1=1.2,whicharereasonablevaluesforrealpowersystems.Intuitively,thesevaluesindicatethatthepercentageofactivepowerlosseswillbenolargerthan20%ofthetotalactivepowerinjections.Notethattheexactvalueofb2isnotnecessary.Underthissimulationsetting,asgiveninTheorem1,theupperboundof thecontrolstepsizeis1
n b2
1≈0.0694.Wenotethatcomprehensivesimulationsincludingthetwotimescalescanbedoneusingdatasuchasonesadoptedin[26].However,sincetheODCPtobesolvedontheslowtimescale isastandardproblem,wewillmainlyfocusonsimulations forthefasttimescale,whereourmajorcontributionslie.
A.CaseIInthiscase,weassumethepowerdistributionsystemisimportingenergyfromthebulkpowersystemwithy=−3110kW.Thiscorrespondstothesituationwheretheuncontrollablerenewableenergyresourcesarenotgeneratingmoreactivepowerthanthatneededbytheloads.Inaddition, wesetˆφ0=1nandu0=0n,where1n∈Rnisanall-onesvector.1)TrackingPerformanceDuringEstimation:Fory?=−3000kWandaconstantstepsizeβk=0.02,theDERactive
7
0.02.04.06.08.010.0time(s)−100−75−50−250trackingerror(kW)0.005
0.010.02
0.030.04
0.05
Fig.5.Trackingerrorfory?=−3000kWandvariousconstantcontrolstepsizesinCaseI.(Legendsindicatevaluesofβk.)TABLEIESTIMATEDSENSITIVITIESINCASEIAFTER60ITERATIONS
bus1926384956
true1.03941.04131.04261.04541.0467estimate1.03421.03901.04401.04681.0421
bus64788999
true1.07021.07031.07491.0711estimate1.06961.06971.08171.0702
0.02.04.06.08.010.0time(s)0.000.020.040.06MAE(p.u./p.u.)1.00
1.05
Fig.6.MAEofestimationerrorswithβk=0.02inCaseI.(Legendsindicatethevaluesofinitialestimates.)
powerinjectionsareshowninFig.3.Thenon-smoothnessin theactivepowerproﬁlesiscausedbytherandomperturbationimposedonthecontrolstep.AlsoasshowninFig.4,the convergencerateofthetrackingerrorisnotaffectedbythe trackingtarget,i.e.,thetotalactivepowerrequiredfromthebulkpowersystem.Thetrackingerrore[k]undervariousconstantcontrolstepsizesisshowninFig.5.Asexpected, alargerstepsizewillreducethetrackingerrorfasterthanasmallstepsize.2)EstimationAccuracy:Withβk=0.02andy?=−3000kW,trueandestimatedsensitivitiesarecomparedinTableIandthemeanabsoluteerror(MAE)ofestimation errors,i.e.,Pn i=1|εi[k]|
N,isplottedinFig.6.Theestimatedsen-sitivitiesareveryclosetotheirtruevaluesafter60steps,whichcorrespondsto6s.Notethatˆφ0hasanimportantimpactontheconvergenceofthesensitivityestimationalgorithm.AscanbeseenfromFig.6,whentheinitialvaluesoftheestimated sensitivitiesaresetto1.05,whichisclosertotheirtruevalues,ittakesmuchlesstimetoobtainasmallestimationerror.
0.02.04.06.08.010.0time(s)0.000.020.040.06MAE(p.u./p.u.)0.01
0.030.05
0.070.1
adaptive
Fig.7.MAEofestimationerrorswithβk=0.02undervariousestimationstepsizesinCaseI.(Legendsindicatevaluesofαk.)
0.02.04.06.08.010.0time(s)0.000.020.040.06MAE(p.u./p.u.)0.005
0.010.02
0.030.04
0.05
Fig.8.MAEofestimationerrorsundervariouscontrolstepsizesinCaseI.(Legendsindicatevaluesofβk.)Whiletheestimationstepsize,αk,intheproposedalgo-rithmisadaptive,wealsoinvestigatethecasewhenαkischosentobeconstant.Figure7showstheMAEofestimation errorundervariousconstantestimationstepsizes.AscanbeseenfromFig.7,theMAEofestimationwillconvergeto somenon-zeroconstantunderconstantestimationstepsizes.Theimpactofthecontrolstepsizesontheestimationaccuracyisalsoinvestigated.Figure8showstheMAEof estimationerrorsundervariouscontrolstepsizes.Withalargecontrolstepsize,thetrackingerrorconvergesto0quickly,leadingtoasituationwherethesystemcannotgetsufﬁcient excitationandconsequently,theestimationerrorscannotbefurtherreduced.
B.CaseIIInthiscase,weassumethepowerdistributionsystemisexportingenergytothebulkpowersystemwithy=1000kW.Thiscorrespondstothesituationwheretheuncontrollable renewableenergyresourcesaregeneratingmuchmoreactive powerthanthatneededbytheloads.Wesetˆφ0=1nandu0=0n.1)TrackingPerformanceDuringEstimation:Usingacon-stantstepsizeβk=0.02,theconvergencerateofthetrackingerrorundervarioustrackingtargetisshowninFig.9.SimilartoresultsinCaseI,theconvergencerateisnotaffectedby thetrackingtarget.2)EstimationAccuracy:Withβk=0.02andy?=1100kW,trueandestimatedsensitivitiesarecomparedin
8
0.02.04.06.08.010.0time(s)−600−400−2000trackingerror(kW)1100kW
1200kW1300kW
1400kW1500kW
1600kW
Fig.9.Trackingerrorforβk=0.02undervarioustrackingtargetsinCaseII.(Legendsindicatevaluesofy?.)TABLEIIESTIMATEDSENSITIVITIESINCASEIIAFTER60ITERATIONS
bus1926384956
true0.95330.95260.95180.95090.9254estimate0.95880.95120.94970.94750.9285
bus64788999
true0.88720.84770.84880.8700estimate0.88540.85360.83960.8707
0.02.04.06.08.010.0time(s)0.000.020.040.060.08MAE(p.u./p.u.)
Fig.10.MAEofestimationerrorswithβk=0.02inCaseII.TableIIandtheMAEofestimationerrorsisplottedinFig.10,respectively.SimilartotheresultsinCaseI,theestimatedsen-sitivitiesareveryclosetotheirtruevaluesafter60steps.Thisveriﬁesthattheproposedestimationalgorithmcaneffectivelyestimatethesensitivitiesunderdifferentoperatingconditionsofthepowerdistributionsystem.WenotethattheperformanceoftheproposedestimationalgorithmisindependentofthenumberofDERs.Toseethis, wesimulateacasewheretheDERatbus99getsdisconnected andconsequently,thereare8DERsinthepowerdistributionsystem.Thesensitivitiesatthese8DERbusescanstillbeestimatedeffectively,asisshowninFig.11.
C.CaseIIIInthiscase,weillustratehowtheproposedframeworkhandleslinecongestions.ThesetupisthesameasCaseII
exceptthatthetrackingtargetisy?=1500kW.Wesetthecapacitylimitofline(55,56)to40kWtocreatecongestion.Forsimplicity,allotherlinesareassumedtohaveaninﬁnite
0.02.04.06.08.010.0time(s)0.000.020.040.060.08MAE(p.u./p.u.)
Fig.11.MAEofestimationerrorswithβk=0.02inCaseIIwith8DERs.
192638495664788999DERbus020406080DERacitivepower(kW)before after
Fig.12.DERactivepowerset-pointbeforeandaftersolvingtheODCPinCaseIII.
capacity;yet,wewouldliketoemphasizethattheproposed frameworkcanhandlemultiplelinecongestions.TheobjectivefunctionintheODCPisassumedtobec(pg)=kpg−˜pgk2,where˜pgisthecurrentDERactivepowerinjectionvectorasusedin(9).Intuitively,thisobjectivefunctionwillfavorthesolutionwiththeleastchangeintheDERactivepower injections.Theestimationalgorithmisﬁrstrantoobtainanestimateofthesensitivityvector.Aftertheestimationalgorithmends,theDERatbus56isgenerating51.3kW,whichexceedsthecapacitylimitofline(55,56).TheODCPisrunafterwardstoadjusttheactivepowerset-pointsoftheDERs.Figure
12showstheDERactivepowerset-pointbeforeandafter solvingtheODCP.TheDERatbus56isdispatcheddown to40kW,whichconformswiththecapacitylimitofline(55,56).AllotherDERsaredispatchedupsuchthattheactivepowerexchangedbetweenthedistributionandbulkpower systemsstillequalsto1500kW.Wenotethatline(55,56)isoverloadedforashortperiodduringtheestimationprocessbuttheisquicklyrestoredtoanormalloadinglevelaftertheDERactivepowerset-pointsareadjustedviatheODCP.VI.CONCLUDINGREMARKSInthispaper,weproposedadata-drivencoordinationframe-workforDERsinalossypowerdistributionsystem,the modelofwhichisnotcompletelyknown,tocollectively providesomepre-speciﬁedamountofactivepowertoabulk powersystemataminimumgeneratingcost,whilerespecting distributionlinecapacitylimits.Theproposedframework consistsofaLTVIOmodel,anestimator,andacontroller.
9Weshowedthatusingtheestimationalgorithmproposedintheframework,theestimatedparametersconvergetothe trueparametersa.s.,andthetotalprovidedactivepowercon-vergestotherequiredamountduringtheestimationprocess.Thedata-drivennatureofthisframeworkmakesitadaptive tovarioussystemoperatingconditions.Wevalidatedthe effectivenessoftheproposedframeworkthroughnumerical simulationsonamodiﬁedversionoftheIEEE123-bustest feeder.Therearetwopotentialdirectionsforfuturework.Theﬁrstoneistodevelopefﬁcientestimationalgorithmsforscenariosinwhichtheestimationprocess(withrandomperturbations)cannotbeexecutedonatimescalethatismuchfasterthan theoneonwhichtheloadsvary,andquantiﬁestheimpacts ofloadvariabilityontheperformanceofthealgorithm; thisessentiallyrelaxesAssumption3.Thesecondoneisto extendtheproposedestimationalgorithmstoscenarioswheretherearemultipleconnectionbusesbetweenthedistributionsystemandthebulkpowersystem,andconsequently,multipleoutputs.APPENDIXA.ProofofTheorem1Theconvergenceanalysisofthecontrolstepduringtheestimationprocessreliesonthefollowingtwolemmas.
Lemma1.Thereexists¯φ[k]satisfying0n≤¯φ[k]≤ˆφ[k],suchthat(14)isequivalenttou[k]=u[k−1]−βke[k−1]W[k]¯φ[k].Also,¯φ[k]=0nifandonlyifu[k]=u oru[k]=
u.Furthermore,ifu[k]6=u andu[k]6=
u,thereexistsi∈Ngsuchthat¯φi[k]=ˆφi[k]∈[b
1, b1].Proof.Ifu[k−1]−βke[k−1]W[k]ˆφ[k]∈U,thenwesimplyset¯φ[k]=ˆφ[k].Withoutlossofgenerality,ﬁrstconsiderthecasewherethefollowingholdsforsomei∈Ng:ui[k−1]−βke[k−1]wi[k]ˆφi[k]>
ui.(A.1)Then,e[k−1]<0andwi[k]>0sinceotherwise(A.1)cannotnothold.Therefore,ui[k]=PU(ui[k−1]−βke[k−1]wi[k]ˆφi[k])=
ui.(A.2)Let¯φi[k]=ui[k−1]−
ui
βke[k−1]wi[k];bydeﬁnition,¯φi[k]=0ifandonlyifui[k−1]=
ui.Then,wehavethat:ui[k]=ui[k−1]−βke[k−1]wi[k]¯φi[k].(A.3)Iffollowsfrom(A.1),(A.2),and(A.3)thatβke[k−1]ˆφi[k]wi[k]<βke[k−1]¯φi[k]wi[k],(A.4)whichleadsto0≤¯φi[k]<ˆφi[k].Asimilarargumentcanbeusedtoforthecasewhereui[k−1]−βke[k−1]wi[k]ˆφi[k]<u iandforsomei∈Ng.Ifu[k]6=u andu[k]6=
u,thenthereexistsi∈Ngsuchthatu i<ui[k]<
ui,whichimpliesui[k]=ui[k−1]−βke[k−1]wi[k]ˆφi[k].(A.5)Therefore,¯φi[k]=ˆφi[k].Consequently,¯φi[k]=ˆφi[k]∈[b
1, b1].ItcanbeeasilyseenthatifUissufﬁcientlylargeandnoDERhitsitscapacitylimits,then¯φ[k]=ˆφ[k].
Lemma2.LetXk,k=1,2,···,beindependentlyidenticallydistributed(i.i.d.)randomvariables.AssumeXk>0andE[Xk]∈(0,1),whereEdenotesexpectation.LetYk=Qk i=1Xi.Then,limk→∞Yk=0a.s.Proof.NotethatYk=expnPk i=1logXio.BytheStrongLawofLargeNumbers(seeProposition2.15in[27]),we havethatlimk→∞kXi=11
klogXi=E[logX1],a.s.(A.6)ByJensen’sinequality(seeTheorem2.18in[27]),wehave thatE[logX1]≤logE[X1]<0.(A.7)Therefore,limk→∞kXi=1k1
klogXi=−∞,a.s.,(A.8)whichleadstolimk→∞Yk=limk→∞exp(kXi=1k1
klogXi)=exp(limk→∞kXi=1k1
klogXi)=0,a.s.;(A.9)thiscompletestheproof.
UsingLemma1andLemma2,wecanprovetheTheorem1asfollows:
Proof.By(6),wehavethate[k]−e[k−1]=φ[k]>∆u[k].(A.10)ByLemma1,wehavethat∆u[k]=−βke[k−1]W[k]¯φ[k],(A.11)where0n≤¯φ[k]≤ˆφ[k].Substituting(A.11)into(A.10)leadstoe[k]=(1−βkφ[k]>W[k]¯φ[k])e[k−1].(A.12)Deﬁneρk=1−βkφ[k]>W[k]¯φ[k],thene[k]=e[0]kY
i=1ρi.(A.13)ByAssumption3,0<b
1≤φi[k]≤
b1.Inaddition,itfollowsfromLemma1that0≤¯φi[k]≤ˆφi[k]≤
b1.Therefore,φ[k]>W[k]¯φ[k]canbeboundedasfollows:0≤φ[k]>W[k]¯φ[k]=nXi=1wi[k]φi[k]¯φi[k]≤n b2
1.(A.14)
10Sinceβk<1
n b2
1,thenalle[k]hasthesamesignforallk(positiveife[0]>0,andnegativeotherwise).Asaresult,theentriesof∆u[k]alwayshavethesamesignby(A.11).(a)Ife[k]=0forsomek∈N,then∆u[k+1]=0n.Thecontrolandestimationalgorithmswillstopupdatingaccordingto(12)and(14).Inthiscase,u[k]mayequaltou or uorneither,andthesystemattainsanequilibrium.
(b)Nowsupposee[k]6=0,∀k∈N.Sincetheincrementsofualwayshavethesamesign,theentriesofucannothittheirboundsindifferentdirections,i.e.,somehittheirlowerboundswhileothershittheirupperbounds.
(b.1)Ifu[k]=u forsomeiterationk,thene[k]>0,∀k∈N.By(14),wehavethatu[k+1]=PU(u
−βke[k]W[k+1]ˆφ[k+1])=u
.(A.15)Thus,∆u[k+1]=0n,whichleadstoe[k+1]=e[k]by(A.10).Therefore,uwillequaltou ande[k0]=e[k]>0forallk0>k.Similarly,whenu[k]=
u,uwillbeequalto u,andewillbeequaltoe[k]<0inallfuturetimeintervals.Thesystemattainsanequilibriuminbothcases.
(b.2)Ifu[k]6=u andu[k]6=
u,∀k∈N,byLemma1,thereexistsi∈Ngsuchthat¯φi[k]∈[b
1, b1].Then,φ[k]>W[k]¯φ[k]=nXi=1φi[k]¯φi[k]wi[k]≥b
2
1wi[k].(A.16)Thus,byusing(A.14)and(A.16),itfollowsthatρk∈[1−βkn b2
1,1−βkb
2
1wi[k]].Deﬁne
ρk=1−wi[k],then
ρkequalsto1−or1,eachwithprobability0.5,andE[
ρk]=1−
2∈(0,1).Notethat0<<b
2
1
n b2
1implies
ρk>0.ByLemma2,limk→∞kYi=1
ρi=0,a.s.(A.17)Whenβk∈(
b
2
1,1
n b2
1),0≤ρk≤
ρk.Then,inana.s.sense,limk→∞|e[k]|=|e[0]|limk→∞kYi=1ρi≤|e[0]|limk→∞kYi=1
ρi=0.(A.18)Since|e[k]|≥0,limk→∞|e[k]|=0a.s.Inaddition,by(A.11),limk→∞∆u[k]=0na.s.
Remark2.IfUissufﬁcientlylargeandnoDERhitsthecapacitylimits,then¯φ[k]=ˆφ[k]andφ[k]>W[k]¯φ[k]≥b
2
1Pn i=1wi[k].Followingasimilarargumentasinpart(b.2)intheproofofTheorem1,wecanshowe[k]convergesto0a.s.whenβk∈(
nb
2
1,1
n b2
1),where0<<b
2
1
b2
1.Followingasimilarargument,Corollary1canbeprovedasfollows:
Proof.Whenthecontrolupdaterulein(13)isusedinsteadoftheonein(14),e[k]=(1−βkφ[k]>¯φ[k])e[k−1].(A.19)Ifu[k]6=u andu[k]6=
u,φ[k]>¯φ[k]=φi[k]¯φi[k]≥b
2
1.Deﬁneρk=1−βkφ[k]>¯φ[k].Whenβk∈(
b
2
1,1
n b2
1),ρk<1−.Therefore,

e[k]
e[k−1]

=ρk<1−.
B.ProofofTheorem2Theconvergenceanalysisoftheestimationstepusessomeconvergenceresultsfor∆φ[k],whicharepresentednext.Lemma3.LetXk,k=1,2,···,bei.i.d.randomvariablesthattakevalue1withprobability0.5,orsomeconstantx∈(0,1),alsowithprobability0.5.LetYk=Qk i=1XiandZ=P∞
i=1Yi.Then,Zisboundeda.s.Proof.LetKdenotethemaximumnumberof1’sthatappearscontinuouslyinthesequence{Xk};then,thesequence{Yk}willhaveanew(smaller)valueatmostafterK+1steps.WeclaimZisunboundedonlyifKisinﬁnite.SupposeXj=x,andXk=1fork=j+1,···,j+m,thenYj=Yj+1=···=Yj+mandPj+mi=jYj=(m+1)Yj≤(K+1)Yj.Therefore,Z=∞Xi=1Yi≤(K+1)∞Xi=0xi=K+1
1−x.(A.20)Itfollowsfrom(A.20)thatZisunboundedonlyifKisinﬁnite.However,P{M=∞}≤P{Xi+1=···=Xi+K=1,forsomei}=1
2∞=0,wherePdenotesprobability.Thus,Zisboundeda.s.
Lemma4.Usingestimationupdaterule(12)andcontrolupdaterule(14),withβk∈(
nb
2
1,1
n b2
1),where0<<b
2
1
b2
1isagivenparameter,thenlimk→∞k∆φ[k]k=0,a.s.(A.21)and∞Xk=1k∆φ[k]k<∞,a.s.(A.22)Proof.IffollowsfromtheproofofTheorem1thattheentriesof∆u[k]alwayshavethesamesign.Firstconsiderthecasewhere∆u[k]≥0nforallk∈N.Notethatφ[k]>=∂h
∂u

˜u[k],where˜u[k]=aku[k]+(1−ak)u[k−1]withak∈[0,1],i.e.,u[k−1]≤˜u[k]≤u[k].Similarly,φ[k−1]>=∂h
∂u

˜u[k−1],whereu[k−2]≤˜u[k−1]≤u[k−1].Thus,byAssumption3,wehavethatk∆φ[k]k≤b2k˜u[k]−˜u[k−1]k≤b2ku[k]−u[k−2]k=b2k∆u[k]+∆u[k−1]k≤b2(k∆u[k]k+k∆u[k−1]k).(A.23)Sincelimk→∞k∆u[k]k=0a.s.byTheorem1,asaresult,limk→∞(k∆u[k]k+k∆u[k−1]k)=0a.s.,whichgiveslimk→∞k∆φ[k]k=0,a.s.(A.24)
11Assumeu[k]=0nforallk<0,thenwehavethat∞X
k=1k∆φ[k]k≤∞X
k=1b2(k∆u[k]k+k∆u[k−1]k)≤2b2∞X
k=0k∆u[k]k≤2b2∞X
k=0kβkW[k]ˆφ[k]e[k−1]k≤2b2
n b2
1√
n b1∞X
k=0|e[k−1]|=2b2
√
n b1∞Xk=−1|e[k]|.(A.25)Recallthat
ρkequalsto1−or1,eachwithprobability0.5,where
ρkisdeﬁnedintheproofofTheorem1.Therefore,byLemma3,P∞
k=1Qk i=1
ρiisboundeda.s.Whenβk∈(
b
2
1,1
n b2
1),0≤ρk≤
ρk,and∞X
k=0|e[k]|=|e[0]|(1+∞Xk=1kYi=1ρi)≤|e[0]|(1+∞X
k=1kYi=1
ρi).(A.26)Asaresult,P∞
k=−1|e[k]|isboundeda.s.since|e[−1]|isalsobounded.Thecasewhere∆u[k]≤0nforallk∈Ncanbeprovedsimilarly.
Theconvergenceanalysisoftheestimationstepalsoreliesonthefollowinglemma(seeTheorem1in[28]).
Lemma5.LetXk,Yk,Zk,k=1,2,···,benon-negativevariablesinRsuchthatP∞
k=0Yk<∞,andXk+1≤Xk+Yk−Zk,thenXkconvergesandP∞
k=0Zk<∞.UsingLemma4andLemma5,Theorem2canthenbeprovedasfollows:
Proof.Consideranarbitrarysamplepath.Withoutlossofgenerality,assumee[k]<0,itfollowsfromTheorem1thate[k]<0,∀k∈N.Sinceu[k]∈(u
, u),∀k∈N,(14)becomes∆u[k]=−βke[k−1]W[k]ˆφ[k].(A.27)Itfollowsfrom(6)and(12)thatˆφ[k+1]=PB(ˆφ[k]−αk+1∆u[k]∆u[k]>ε[k]).(A.28)Bydeﬁnition,theestimationerroratiterationkisε[k+1]=PB(ˆφ[k]−αk+1∆u[k]∆u[k]>ε[k])−φ[k+1].(A.29)Sinceφ[k+1]=PB(φ[k+1]),bythenon-expansivenessoftheprojectionoperation(seeProposition1.1.9in[29]),thenkε[k+1]k≤kε[k]−αk+1∆u[k]∆u[k]>ε[k]−∆φ[k+1]k≤kε[k]−αk+1∆u[k]∆u[k]>ε[k]k+k∆φ[k+1]k.(A.30)Letg(αk+1)=kε[k]−αk+1∆u[k]∆u[k]>ε[k]k2;then,gattainsitsminimumatαk+1=1
k∆u[k]k2,whichiskε[k]k2−(ε[k]>∆u[k]
k∆u[k]k)2=kε[k]k2−(ε[k]>W[k]ˆφ[k]
kW[k]ˆφ[k]k)2.(A.31)Deﬁnecosθk=ε[k]>
kε[k]kW[k]ˆφ[k]
kW[k]ˆφ[k]k.Consequently,g(αk+1)=(1−sin2θk)kε[k]k2,andkε[k+1]k≤|sinθk|kε[k]k+k∆φ[k+1]k.(A.32)LetXk=kε[k]k,Yk=k∆φ[k+1]k,andZk=(1−|sinθk|)kε[k]k.Then,Xk+1≤Xk+Yk−Zk.Also,P∞
k=0Yk=P∞
k=1k∆φ[k]k<∞byLemma4.Therefore,byLemma5,kε[k]kconverges,andP∞
k=1(1−|sinθk|)kε[k]k<∞,whichfurtherimplieslimk→∞(1−|sinθk|)kε[k]k=0.Letε?denotethelimitofkε[k]k;then,limk→∞|sinθk|kε[k]k=limk→∞(|sinθk|−1)kε[k]k+limk→∞kε[k]k=ε?.(A.33)Next,weshowε?=0bycontradiction.Assumeε?>0.Sincebothkε[k]kand|sinθk|kε[k]kconvergestoε?,limk→∞|sinθk|=limk→∞|sinθk|kε[k]k limk→∞kε[k]k=1,(A.34)whichimplies|cosθk|convergesto0.Sincekε[k]kandkW[k]ˆφ[k]karebounded,then|ε[k]>W[k]ˆφ[k]|convergesto0.DeﬁneEi[k]={wj[k]=1ifj=i,wj[k]=0otherwise};thenP{Ei[k]}=1
2n.Consequently,P∞
k=1P{Ei[k]}=∞.AlsonotethatEi[k],k∈N,areindependent.BytheBorel-CantelliLemma(seeLemma1.3in[27]),
P{Ei[k]inﬁnitelyoften}=1;therefore,thereareinﬁnitelymanytimeinstantsthatwi[k]=1andwj[k]=0forallj6=i.LetKidenotethesetofsuchtimein-stants.Then|ε[k]>W[k]ˆφ[k]|=|εi[k]ˆφi[k]|fork∈Ki.Thesequence{|εi[k]ˆφi[k]|,k∈Ki}isasubsequenceof{|ε[k]>W[k]ˆφ[k]|};therefore,italsoconvergesto0.Notethatˆφ[k]>0;thus,εi[k]convergesto0.Sinceiisarbitrary,weconcludethatkεk[k]convergesto0,whichimpliesε?=0,contradiction.Sincethisresultholdsforallsamplepaths,thenweconcludethatkε[k]kconvergesto0a.s.
REFERENCES[1]J.DriesenandF.Katiraei,Designfordistributedenergyresources,IEEEPowerEnergyMag.,vol.6,no.3,pp.3040,May2008.[2]A.Kulmala,S.Repo,andP.J¨arventausta,Coordinatedvoltagecontrolindistributionnetworksincludingseveraldistributedenergyresources,IEEETrans.SmartGrid,vol.5,no.4,pp.20102020,July2014.[3]H.Xu,A.D.Dom´ınguez-Garc´ıa,andP.W.Sauer,Adata-drivenvoltagecontrolframeworkforpowerdistributionsystems,inProc.ofIEEEPESGeneralMeeting,Portland,OR,Aug.2018,pp.15.[4]K.Zhang,W.Shi,H.Zhu,E.Dall’Anese,andT.Basar,Dynamicpowerdistributionsystemmanagementwithalocallyconnectedcommunica-tionnetwork,IEEEJ.Sel.TopicsSignalProcess.,vol.12,no.4,pp.673687,Aug2018.[5]A.D.Dom´ınguez-Garc´ıaandC.N.Hadjicostis,Coordinationandcon-trolofdistributedenergyresourcesforprovisionofancillaryservices,inProc.ofIEEEInternationalConferenceonSmartGridCommunications,Oct.2010,pp.537542.[6]E.Dall’Anese,S.Guggilam,A.Simonetto,Y.C.Chen,andS.V.Dhople,Optimalregulationofvirtualpowerplants,IEEETrans.PowerSyst.,vol.33,no.2,pp.18681881,March2018.
12[7]H.Xu,A.D.Dom´ınguez-Garc´ıa,andP.W.Sauer,Adaptivecoordina-tionofdistributedenergyresourcesinlossypowerdistributionsystems,inProc.ofIEEEPESGeneralMeeting,Portland,OR,Aug.2018,pp.15.[8]O.Ardakanian,V.W.Wong,R.Dobbe,S.H.Low,A.vonMeier,C.Tomlin,andY.Yuan,Onidentiﬁcationofdistributiongrids,arXivpreprintarXiv:1711.01526,2017.[9]Y.C.Chen,A.D.Dom´ınguez-Garc´ıa,andP.W.Sauer,Measurement-basedestimationoflinearsensitivitydistributionfactorsandapplica-tions,IEEETrans.PowerSyst.,vol.29,no.3,pp.13721382,May2014.[10]Y.C.Chen,J.Wang,A.D.Dom´ınguez-Garc´ıa,andP.W.Sauer,Measurement-basedestimationofthepowerﬂowjacobianmatrix,IEEETrans.SmartGrid,vol.7,no.5,pp.25072515,Sept.2016.[11]Z.HouandS.Jin,Data-drivenmodel-freeadaptivecontrolforaclassofmimononlineardiscrete-timesystems,IEEETrans.NeuralNetw.,vol.22,no.12,pp.21732188,Dec.2011.[12]J.Zhang,X.Zheng,Z.Wang,L.Guan,andC.Y.Chung,Powersystemsensitivityidentiﬁcationinherentsystempropertiesanddataquality,IEEETrans.PowerSyst.,vol.32,no.4,pp.27562766,July2017.[13]Y.C.Chen,A.D.Dom´ınguez-Garc´ıa,andP.W.Sauer,Asparserep-resentationapproachtoonlineestimationofpowersystemdistributionfactors,IEEETransactionsonPowerSystems,vol.30,no.4,pp.17271738,July2015.[14]K.E.V.Horn,A.D.Dom´ınguez-Garc´ıa,andP.W.Sauer,Measurement-basedreal-timesecurity-constrainedeconomicdispatch,IEEETrans.PowerSyst.,vol.31,no.5,pp.35483560,Sept.2016.[15]C.Lu,Y.Zhao,K.Men,L.Tu,andY.Han,Wide-areapowersystemstabiliserbasedonmodel-freeadaptivecontrol,IETControlTheoryAppl.,vol.9,no.13,pp.19962007,2015.[16]J.Zhang,C.Y.Chung,andY.Han,Onlinedampingratiopredictionusinglocallyweightedlinearregression,IEEETrans.PowerSyst.,vol.31,no.3,pp.19541962,May2016.[17]Z.Hou,R.Chi,andH.Gao,Anoverviewofdynamic-linearization-baseddata-drivencontrolandapplications,IEEETrans.Ind.Electron.,vol.64,no.5,pp.40764090,May2017.[18]J.Zhang,Z.Wang,X.Zheng,L.Guan,andC.Y.Chung,Locallyweightedridgeregressionforpowersystemonlinesensitivityidentiﬁ-cationconsideringdatacollinearity,IEEETrans.PowerSyst.,vol.PP,no.99,pp.11,2017.[19]J.Zhang,C.Y.Chung,andL.Guan,Noiseeffectandnoise-assistedensembleregressioninpowersystemonlinesensitivityidentiﬁcation,IEEETrans.Ind.Informat.,vol.13,no.5,pp.23022310,Oct.2017.[20]J.ZhangandH.Xu,Microperturbationmethodforpowersystemonlinemodelidentiﬁcation,IEEETrans.Ind.Informat.,vol.12,no.3,pp.10551063,June2016.[21]J.ZhangandH.Xu,Onlineidentiﬁcationofpowersystemequivalentinertiaconstant,IEEETrans.Ind.Electron.,vol.64,no.10,pp.80988107,Oct2017.[22]A.Iusem,Ontheconvergencepropertiesoftheprojectedgradientmethodforconvexoptimization,Computational&AppliedMathemat-ics,vol.22,no.1,pp.3752,2003.[23]O.Ajala,M.Almeida,P.W.I.Celanovic,Sauer,andA.D.Dom´ınguez-Garc´ıa,Ahierarchyofmodelsformicrogridswithgrid-feedinginvert-ers,inProc.oftheIREPBulkPowerSystemDynamicsandControlSymposium,Aug.2017.[24]IEEEdistributiontestfeeders.[Online].Available:https://ewh.ieee.org/soc/pes/dsacom/testfeeders/[25]R.D.Zimmerman,C.E.Murillo-Sanchez,andR.J.Thomas,Mat-power:Steady-stateoperations,planning,andanalysistoolsforpowersystemsresearchandeducation,IEEETrans.PowerSyst.,vol.26,no.1,pp.1219,Feb.2011.[26]B.A.RobbinsandA.D.Dom´ınguez-Garc´ıa,Optimalreactivepowerdispatchforvoltageregulationinunbalanceddistributionsystems,IEEETrans.PowerSyst.,vol.31,no.4,pp.29032913,2016.[27]B.Hajek,Randomprocessesforengineers.CambridgeUniversityPress,2015.[28]H.RobbinsandD.Siegmund,Aconvergencetheoremfornonnegativealmostsupermartingalesandsomeapplications,inHerbertRobbinsSelectedPapers.Springer,1985,pp.111135.[29]D.P.Bertsekas,Convexoptimizationtheory.AthenaScientiﬁcBelmont,2009.HanchenXu(S’12)receivedtheB.Eng.andM.S.degreesinelectricalengineeringfromTsinghuaUniversity,Beijing,China,in2012and2014,respectively,andtheM.S.degreeinappliedmathematicsfromtheUniversityofIllinoisatUrbana-Champaign,Urbana,IL,USA,in2017,whereheiscurrentlyworkingtowardthePh.D.degreeattheDepartmentofElectricalandComputerEngineering.Hiscurrentresearchinterestsincludecontrol,optimization,reinforcementlearning,withapplicationstopowersystemsandelectricitymarkets.
AlejandroD.Dom´ınguez-Garc´ıa(S’02,M’07)receivedthedegreeofelectricalengineeringfromtheUniversityofOviedo(Spain)in2001andthePh.D.degreeinelectricalengineeringandcomputersciencefromtheMassachusettsInstituteofTechnology,Cambridge,MA,in2007.HeisProfessorwiththeDepartmentofElectricalandComputerEngineer-ing(ECE),andResearchProfessorwiththeCoordinatedScienceLaboratoryandtheInformationTrustInstitute,allattheUniversityofIllinoisatUrbana-Champaign.HeisafﬁliatedwiththeECEPowerandEnergySystemsarea,andhasbeenaGraingerAssociatesinceAugust2011.Hisresearchinterestsareintheareasofsystemreliabilitytheoryandcontrol,andtheirapplicationstoelectricpowersystems,powerelectronics,andembeddedelectronicsystemsforsafety-critical/fault-tolerantaircraft,aerospace,andautomotiveapplications.Dr.Dom´ınguez-Garc´ıareceivedtheNSFCAREERAwardin2010,andtheYoungEngineerAwardfromtheIEEEPowerandEnergySocietyin2012.In2014,hewasinvitedbytheNationalAcademyofEngineeringtoattendtheUSFrontiersofEngineeringSymposium,andwasselectedbytheUniversityofIllinoisatUrbana-ChampaignProvosttoreceiveaDistinguishedPromotionAward.In2015,hereceivedtheUofICollegeofEngineeringDean’sAwardforExcellenceinResearch.HeiscurrentlyaneditorfortheIEEETransactionsonControlofNetworkSystems;healsoservedasaneditoroftheIEEETransactionsonPowerSystemsandIEEEPowerEngineeringLettersfrom2011to2017.PeterW.Sauer(S’73,M’77,SM’82,F’93,LF’12)obtainedhisBachelorofSciencedegreeinelectricalengineeringfromtheUniversityofMissouriatRollain1969.From1969to1973,hewastheelectricalengineeronadesignassistanceteamfortheTacticalAirCommandatLangleyAirForceBase,Virginia.HeobtainedtheMasterofScienceandPh.D.degreesinElectricalEngineeringfromPurdueUniversityin1974and1977respectively.FromAugust1991toAugust1992heservedastheProgramDirectorforPowerSystemsintheElectricalandCommunicationSystemsDivisionoftheNationalScienceFoundationinWashingtonD.C.HeisacofounderofthePowerSystemsEngineeringResearchCenter(PSERC)andthePow-erWorldCorporation.HeisaregisteredProfessionalEngineerinVirginiaandIllinois,aFellowoftheIEEE,andamemberoftheU.S.NationalAcademyofEngineering.HeiscurrentlytheGraingerChairProfessorofElectricalEngineeringatIllinois.Additionalinformationcanbefoundat:https://ece.illinois.edu/directory/proﬁle/psauer.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

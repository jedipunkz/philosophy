---
source: "https://arxiv.org/abs/2604.21588v1"
title: "Piercing all maximum cliques in hypergraphs"
author: "Andreas Holmsen, Attila Jung, Balázs Keszegh, Dániel G. Simon, Gábor Tardos"
year: "2026"
publication: "arXiv preprint / math.CO"
download: "https://arxiv.org/pdf/2604.21588v1"
pdf: "https://arxiv.org/pdf/2604.21588v1"
captured_at: "2026-05-09T12:28:22Z"
updated_at: "2026-05-09T12:28:22Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代思想"
status: raw
---

# Piercing all maximum cliques in hypergraphs

- 著者: Andreas Holmsen, Attila Jung, Balázs Keszegh, Dániel G. Simon, Gábor Tardos
- 年: 2026
- 掲載情報: arXiv preprint / math.CO
- 情報源: [arxiv](https://arxiv.org/abs/2604.21588v1)
- ダウンロード: https://arxiv.org/pdf/2604.21588v1
- PDF: https://arxiv.org/pdf/2604.21588v1

## Obsidian Links

- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代思想]]
- 関連タグ: #現代思想

## Abstract

Graphs whose maximum clique size exceeds half of the total number of vertices satisfy a classical property: the family of their maximum sized cliques can be pierced by a single vertex. This result dates back to a 1965 theorem by Hajnal. Motivated by this theorem, Jung, Keszegh, Pálvölgyi, and Yuditsky recently conjectured that an analogous result should hold for hypergraphs of larger uniformity, with an appropriate constant replacing the threshold $1/2$. In this paper we refute this conjecture in a strong form. We show that for any constant $c<1$ and integers $k\ge 3$ and $t\ge 1$, there exist $k$-uniform hypergraphs $G$ whose maximum clique size exceeds $c|V(G)|$, yet the family of maximum size cliques of $G$ cannot be pierced by $t$ vertices. This demonstrates that no universal constant threshold guarantees bounded piercing number for maximum cliques in uniform hypergraphs. We discuss further questions concerning the relationship between clique size and piercing maximum cliques in hypergraphs, and introduce a geometric variant of the problem using Helly's Theorem.

## PDF Text

Piercingallmaximumcliquesinhypergraphs∗AndreasHolmsen†AttilaJung‡Bal´azsKeszegh§D´anielG.Simon¶G´aborTardosApril24,2026AbstractGraphswhosemaximumcliquesizeexceedshalfofthetotalnumberofverticessatisfyaclassicalproperty:thefamilyoftheirmaximumsizedcliquescanbepiercedbyasinglevertex.ThisresultdatesbacktoHa-jnal[Haj65].Motivatedbythistheorem,Jung,Keszegh,P´alv¨olgyi,andYuditskyrecentlyconjecturedthatananalogousresultshouldholdforhypergraphsoflargeruniformity,withanappropriateconstantreplacingthethreshold1
2.Inthispaperwerefutethisconjectureinastrongform.Weshowthatforanyconstantc<1andintegersk≥3andt≥1,thereexistk-uniformhypergraphsGwhosemaximumcliquesizeexceedsc|V(G)|,yetthefam-ilyofmaximumsizecliquesofGcannotbepiercedbytvertices.Thisdemonstratesthatnouniversalconstantthresholdguaranteesboundedpiercingnumberformaximumcliquesinuniformhypergraphs.Wediscussfurtherquestionsconcerningtherelationshipbetweencliquesizeandpiercingmaximumcliquesinhypergraphs,andintroduceageo-metricvariantoftheproblemusingHelly’sTheorem.
∗ResearchwassupportedbytheERCAdvancedgrantsGeoScape882971andERMiD,intheframeworkoftheSpecialSemesteronDiscreteandConvexGeometryattheErd˝osCenter,Budapest.†DepartmentofMathematicalScinces,KAIST,Daejeon,SouthKoreaandDiscreteMath-ematicsGroup,IBS,Daejeon,SouthKorea.SupportedbytheInstituteforBasicScience(IBS-R029-C1).‡HUN-RENAlfr´edR´enyiInstituteofMathematicsandELTEE¨otv¨osLor´andUniversity,Budapest,Hungary.ResearchsupportedbytheERCAdvancedGrantno.101054936“ER-MiD”andbytheEXCELLENCE-24projectno.151504CombinatoricsandGeometryoftheNRDIFund.§HUN-RENAlfr´edR´enyiInstituteofMathematicsandELTEE¨otv¨osLor´andUniversity,Budapest,Hungary.ResearchsupportedbytheERCAdvancedGrantno.101054936“ER-MiD”andbytheEXCELLENCE-24projectno.151504CombinatoricsandGeometryoftheNRDIFund.¶HUN-RENAlfr´edR´enyiInstituteofMathematicsandELTEE¨otv¨osLor´andUniversity,Budapest,Hungary.ResearchwassupportedbytheERCAdvancedGrantGeoscape882971.HUN-RENAlfr´edR´enyiInstituteofMathematics,Budapest,Hungary.Researchsup-portedbytheERCAdvancedGrantsno.101054936“ERMiD”andNo.882971“GeoScape”.1
arXiv:2604.21588v1 [math.CO] 23 Apr 2026
1IntroductionForauniformhypergraphG,letv(G)bethenumberofvertices,andω(G)bethesizeofthelargestcliqueinG.ForacollectionQofsubsetsofsomebasesetwedenote∩Q=∩Q∈QQand∪Q=∪Q∈QQ.Theorem1(Hajnal[Haj65]).LetGbeagraphandQthecollectionofitsmaximumsizecliques.Then|∩Q|≥2ω(G)−|∪Q|Corollary2.IfGisagraphwithω(G)>1
2v(G),thenthemaximumsizecliquesofGcanbehitwithasinglevertex.Jung,Keszegh,P´alv¨olgyi,andYuditskyconjecturedthatsomethingsimilarshouldholdforhypergraphs.Conjecture3(Jung,Keszegh,P´alv¨olgyi,Yuditsky[P´al23]).Forallk≥2thereexistsack<1suchthatifH⊂�[n]k�andω(H)>ckv(H),thenthemaximumsizecliquesofHcanbehitwithasinglepoint.Thecomplementofak-uniformmatchingshowsthatck≥1−1
kifitexists.Weshowthatckdoesnotexist,andtheconjectureisfalseinastrongsense.Letη(H)betheminimumnumberofpointsneededtohitallthemaximumsizecliquesofH.Itisalsocalledthepiercingnumberofthefamilyofmaximumsizecliques.Theorem4.Letk≥3andt≥1beintegers,andc<1bearealnumber.Thereexistsak-uniformhypergraphHsuchthatω(H)≥cv(H)anditsmaximumsizecliquescannotbepiercedbytvertices,i.e.,η(H)>t.Weinvestigatefurthertherelationshipbetweenω(G)andη(G),inthespiritofConjecture3.Letfk,tbetheminimalfunction,suchthatifak-uniformhypergraphHhasv(H)−ω(H)≤fk,t(v(H)),thenη(H)≤t.Wegivenon-triviallowerandupperboundsforthisfunction,yettherestillremainsroomforimprovement.Theorem5.Letk≥3,t≥1beintegers.ThereareconstantsC≥4,c>0,suchthat1
Clogk�tn k�1
k≤fk,t(n)≤c22tn logn.WeproveTheorem4viainductiononkandt.First,inSection2,weproveTheorem4inthecaseofk=3andt=1.NextinSection3,weshowhowthisspecialcaseimpliesTheorem4initsfullgenerality.InSection4,weinvestigatethefunctionfk,t(n)andproveTheorem5.InSection5wediscusssomepartialresultsregardingspecial(geometric)hypergraphs.Inthelastsection,wediscussfurtherresearchdirectionsandposesomeopenquestions.2
2Piercingmaximumsizecliquesbyonevertexin3-uniformhypergraphs.InthissectionweshowthatthereexistsasequenceHmof3-uniformhyper-graphswithω(Hm)
v(Hm)→1andη(H)>1.Weconstructthe3-uniformhypergraphsHmforallm≥3bydefiningmmaximumsizecliquesC1,...,Cmfirst.TheedgesofthehypergraphHmwillbeallthetriplescontainedinatleastoneofthesetsCi.ThevertexsetofHmisthefollowing:let1≤s≤m
3,andtakepairwisedisjointsetsAJofsize�m−2s−1�forJ∈�[m]m−1�andAJofsize1forJ∈�[m]s�.Next,foreachi∈[m],letusdefineCi=�i∈JAJ.ThiswaywehavethatthesetsCiarecliquesofequalsizeandthat∩mi=1Ci=∅.ThespecificsizesofthesetsAJarechosensothatthesetsCiremainmaximumcliquesintheconstruction,aswewillseesoon.Morespecifically,thesizesofthesetsAJareanextremalsolutionoftheinequalityinCase1oftheproofofClaim9.ApartfromthesetsCibeingmaximumcliquesofHm,wealsoneedtoshowthatω(Hm)
v(Hm)=|Ci|
|∪Ci|→1asm→∞.Forthiscondition,s∼cmwithanyc∈(0,1)issufficient.Claim6.Ifc∈(0,1),m→∞,ands∼cmthen|C1|
|∪mi=1Ci|→1.Proof.|C1|
|∪mi=1Ci|=�m−1s−1�+�m−1m−2��m−2s−1�
�ms�+�mm−1��m−2s−1�=s m�ms�+s(m−s)
m�ms�
�ms�+s(m−s)
m−1�ms�==s m+s(m−s)
m
1+s(m−s)
m−1∼c+c(1−c)m
1+c(1−c)m→1.
WecallafamilyFofsetsℓ-wiset-intersecting,ifanysubfamilyofatmostℓsetsofFshareatleasttcommonvertices.Claim7.AsubsetKofatleast3verticesofHmformsacliqueifandonlyifK⊂∪J∈FAJwithsome3-wise1-intersectingF⊂�[m]s�∪�[m]m−1�.Proof.KisacliqueifandonlyifeverysetofthreeverticesofKiscontainedinoneofthesetsCi,innotation�K3�⊂∪mi=1�Ci3�.Thus,foranyAJ1,AJ2,AJ3withAJ1∩K̸=∅,AJ2∩K̸=∅,AJ3∩K̸=∅,therehastobeasharedindexi∈J1∩J2∩J3.
WeusethefollowingstrengtheningoftheErd˝os-Ko-Radotheorem.Theorem8(Wilson[Wil84]).Ifm≥(t+1)sandF⊂�[m]s�is2-wiset-intersecting,then|F|≤�n−ts−t�.Claim9.Ifs≤m
3,thenC1,...,CmaremaximumsizecliquesofHm.3
Proof.AnymaximalcliqueKofHmisoftheformKF=�J∈FAJwitha3-wise1-intersectingfamilyF⊂�[m]s�∪�[m]m−1�byClaim7.Case1:IfthereisnoJ,J′∈Fwith|J∩J′|=1,then���F∩�[m]s����≤�m−2s−2�byTheorem8on2-wise2-intersectingfamilies.Wehave|KF|≤�m−2s−2�+�mm−1��m−2s−1�==�m−1s−1�+�m−1m−2��m−2s−1�=|Ci|forthecorrespondingmaximalcliqueKFasdesired.Case2:IfthereisJ,J′∈Fwith|J∩J′|=1,thenJ∩J′={i}forsomei,andasFis3-wise1-intersecting,everymemberofFcontainsiandthusK⊂Ci.
Forconcreteness,sets=�m
3�.ThenthesequenceHmof3-uniformhyper-graphssatisfyω(Hm)
v(Hm)→1andη(H)>1.3Piercingmaximumsizecliquesbytverticesink-uniformhypergraphs.Generalizingforlargerkisstraightforward.Observation10.IfTheorem4istruefork=3,thenitistrueforallk≥3.Proof.Ifwehavea3-uniformhypergraphHwithω(H)>cv(H)andη(H)>t,thenforany3≤k≤ω,thehypergraphHkofthek-cliquesofHhasthesamesetofmaximumsizecliquesandthussatisfiesω(Hk)>cv(Hk)andη(Hk)>t.
Theprooffortrequiresabitmorework.Fortheconstructionweneedavariantofthelexicographicproductofhypergraphswherecertain”mixed”hyperedgesarealsoallowed.Definition.Giventwok-uniformhypergraphsH1andH2,letV(H1◦H2)=V(H1)×V(H2).Letpi:(v1,v2)�→vibetheprojectiontotheithcoordinateandletE′={{x}×e2:x∈V(H1),e2∈E(H2)},E′′=�e∈�V(H1)×V(H2)k�:|p1(e)|≥2and∃e1∈E(H1)withp1(e)⊆e1�,andE(H1◦H2)=E′∪E′′.Theonlydifferencewiththeoriginallexicographicproductisthecondition|p1(e)|≥2insteadof|p1(e)|≥|e|(whichwouldinturnimplyp1(e)=e1)inthedefinitionofE′′.Thisiscrucialforthefollowingclaimtohold.4
Claim11.Wehave1.v(H1◦H2)=v(H1)v(H2)2.ω(H1◦H2)=ω(H1)ω(H2)3.η(H1◦H2)=η(H1)η(H2)Proof.Wehavev(H1◦H2)=|V(H1)×V(H2)|=v(H1)v(H2).Clearly,asetCisacliqueinH1◦H2,ifandonlyifp1(C)isasubsetofacliqueinH1andCx:={y|(x,y)∈C}isacliqueinH2foreveryx∈p1(C).SoCisamaximumsizecliqueifandonlyifp1(C)isamaximumsizecliqueandsoisCxforeveryx∈p1(C).ThusSisahittingsetofmaximumcliquesinH1◦H2ifandonlyif{x|SxisahittingsetinH2}isahittingsetinH1,andforthatweneedatleastη(H1)η(H2)vertices.
FortheproofofTheorem4forlargert,observethatifasequence(Htm)m∈Nof3-uniformhypergraphsissuchthatη(Htm)≥tandω(Htm)
v(Htm)→1,thenforthesequenceHt2m=Htm◦Htmwehaveη(Ht2m)≥t2andω(Ht2m)
v(Ht2m)→1viaClaim11.4OnapossibleweakeningofConjecture3WeconsideraweakerformofConjecture3.Letfk,tbetheminimalfunction,suchthatifak-uniformhypergraphHhasv(H)−ω(H)≤fk,t(v(H)),thenη(H)≤t.TheconstructionfromSection2providesanupperboundonf3,1asfollows.Claim12.ThepreviouslyconstructedsequenceHmof3-uniformhypergraphsissuchthatv(Hm)−ω(Hm)=O�v(Hm)
logv(Hm)�andη(Hm)>1.Hence,f3,1(n)=O�n logn�.ProofofClaim12.Supposec≤1
3ands=cmisaninteger.Thenω(Hm)=�m−1s−1�+�m−1m−2��m−2s−1�=s m�ms�+s(m−s)
m�ms�andv(Hm)=�ms�+�mm−1��m−2s−1�=�ms�+s(m−s)
m−1�ms�.Thuswehavev(Hm)∼c(1−c)m�ms�andv(Hm)−ω(Hm)∼(1−c2)�ms�.Aswehavelogv(Hm)=Θ(m),theinequalityv(Hm)−ω(Hm)=O�v(Hm)
logv(Hm)�follows.
5
Corollary13.Foranyintegersk≥3,t≥1,thefollowingupperboundholds:fk,t(n)=O�22tn logn�.Proof.Weconsidertheinductionstepsforincreasingthevaluesofkandtinourconstruction.Increasingkdoesnotchangethemaximumcliquesizeintheconstruction,andalsodoesnotdecreaset,hencewegetthesameupperboundforfk,t(n)asforf3,t(n)forallk≥3,t≥1.Nowconsiderourmethodofincreasingt.Inthebasestep,westartwithaconstructionofη(H1)>1,soitisatleasttwo.LetHi+1=Hi◦H1foralli≥1integer.Thenη(Hi)≥2iforalliintegersbyClaim11.Fromhereonward,letvi=v(Hi),ωi=ω(Hi),andletsibesuchthatωi≥vi−si.Wecanpicks1=cv1
logv1byClaim12andletc1=c.Nowweprovebyinductionthatsi≤civi logviforcertainci.ViaClaim11,wehavevi+1=v2iandwi+1=w2i=(vi−si)2=vi+1−4civi+1
logvi+1+4c2ivi+1
log2vi+1.Thussi+1=4civi+1
logvi+1−4c2ivi+1
log2vi+1≤4civi+1
logvi+1.Hence,wecansetci+1=4ciandthussi≤22i−2cvi logviforalli≥1.Thus,ifwerequireη(H)=t,thenforanupperboundonfk,t(n)wecantakethegraphstobeofformH=H⌊log2(t)⌋+1deducedfromsomestartingk-uniformconstructionH1avoidingasinglepiercingvertexofthemaximumcliques.Bythepreviouscalculations,thesekindsofhypergraphshaveη(H)>t,andv(H)−ω(H)≤22tcv(H)
logv(H).Thuswecanconcludethatfk,t(n)=O�22tn logn�
Wealsoprovidealowerboundonthebestpossiblefk,t(v(H)).Toproveit,weneedtodefinesunflowers.Definition.Asunflowerwithspetalsisafamilyofshyperedges,whereeachpairofedgesintersectinthesame(possiblyempty)setofvertices.Sun(s,k)denotestheminimuminteger,suchthatforanysetofthatmanyhyperedgesofsizek,somesofthemformsasunflowerwithspetals.Erd˝osandRadoconjecturedthatSun(s,k)≤c(s)k[ER60].ThecurrentbestboundisduetoBell,Chueluecha,andWarnke:Theorem14(Bell,Chueluecha,Warnke[BCW21]).ThereisaconstantC≥4suchthatSun(s,k)≤(Cslogk)kforallintegerss,k≥2.Wearenowwell-equippedtoproveourlowerboundonfk,t(v(H)).Theorem15.Letk≥3,t≥1beintegers.LetHbeak-uniformhypergraphonnverticeswithcliquenumberω(H).ThereisaconstantC≥4,suchthatifv(H)−ω(H)≤1
Clogk�tn k�1
k−1,thenthemaximalcliquesofHcanbepiercedwithtvertices,soη(H)≤t.Hence,fk,t(n)>1
Clogk�tn k�1
k−1.6
Proof.LetHbeak-uniformhypergraphonnvertices,withcliquenumberω(H)=n−s.Let
Hbethek-uniformcomplementhypergraphofH.Thehyperedgesof
Hcannotbepiercedbys−1vertices,asotherwisetheothern−s+1verticeswouldformacliqueinH.LetH′⊂
Hbeaminimalsub-hypergraphonthesamevertexset,suchthatitshyperedgescannotbepiercedbys−1vertices.ThismeansthatifweremoveanyhyperedgefromH′,theremaininghyperedgescanbepiercedbys−1vertices.Observation.H′doesnotcontainasunflowerwiths+1petals.Proofoftheobservation.SupposeforacontradictionthatH′containsasun-flowerwiths+1petals.Byminimality,ifweeraseoneofthehyperedgesofthesunflower,wecanpiercetherestofH′bys−1vertices.Theses−1verticescanonlypiercetheremainingshyperedgesofthesunflower,ifthereisatleastonevertexinitscenter.Butthatvertexpiercestheremovedhyperedgeaswell,sotheses−1verticespiercethewholeH′,whichisacontradiction.
Hence,H′containsatmostSun(s+1,k)manyhyperedges.NowwecanapplyTheorem14:ThereisauniversalconstantC≥4suchthatH′containsatmost(C(s+1)logk)khyperedges.Observation.IfH′containsavertexxthatiscontainedbyatmostt−1hyperedges,thenxtogetherwithonevertexfromeachofthesehyperedgespierceallmaximalcliquesinH.Proofofobservation.SupposeamaximalcliqueofHisnotpiercedbythesevertices,thentheseverticesformasubsetofthesverticesnotinthemaximumclique.BydefinitionofamaximalcliqueofH,thesesverticespierceH′.Butifweremovethevertexx,thes−1verticesstillpierceH′,sincewepickedonevertexfromeachhyperedgecontainingx.Thisisacontradiction.
Wefinishtheproofusingpigeonholeprinciple.Iftn>k(C(s+1)logk)k,thentheremustbeavertexxpiercedbyatmostt−1hyperedgesinH′.Bythepreviousobservation,wecanthenpickatmosttverticespiercingallmaximumcliquesofH.Afterrearrangingwegetthatifs<1
Clogk�tn k�1
k−1,thenallthemaximalcliquesofHcanbepiercedbyasetofatmosttvertices.
Fork=3anevenbetterupperboundisknownonthesunflowerconjecture,sowecanusethatinstead:Theorem16(Abott,Hanson[AH74]).Foranyintegers≥7,Sun(s,3)≤9
5s(s−1)2+8(s−1)2.Corollary17.Fork=3,f3,t(n)>(5tn
9)1
3−o((tn)1
3).7
Proof.Usethesamemethodasbefore,butconcludewiththeimprovedupper-boundofTheorem16.Sobypigeonholeprinciple,tn>9
5s(s−1)2+8(s−1)2impliestheexistenceofapiercingfamilyofsizeatmostt.Fromthat,onecaneasilyconcludethatf3,t(n)≥s≥(5tn
9)1
3−o((tn)1
3)musthold.
5Piercingmaximumcliquesforintersectionhy-pergraphsofconvexsetsInthissectionwepresentsomepartialresultsregardingananalogueofConjec-ture3forintersectionhypergraphsofconvexsets.LetHdbead+1uniformintersectionhypergraphofnconvexsetsinRd.Intheintersectionhypergraph,theconvexsetscorrespondtotheverticesofthegraph,andd+1verticesformahyperedge,ifandonlyifthecorrespondingd+1convexsetsshareacommonpoint.Question1.Isittruethatforeverydimensiondthereexistsacd<1suchthatifHdisad+1uniformintersectionhypergraphofconvexsetsinRd,andω(Hd)≥cdv(Hd),thenη(Hd)=1?NotethatviaHelly’stheorem,themaximalcliquesofad+1uniforminter-sectionhypergraphofconvexsetsinRdcanberepresentedbypointscontainedbyallconvexsetscorrespondingtotheverticesofthemaximalclique.ThecasewhenHdconsistsofthed+1facetsofad-dimensionalsimplex,showsthatcd≥d d+1.Ontheplane,wecanimproveonthatwiththefollowingconstruction.Claim18.Thereexist3-uniformintersectionhypergraphsHmofplanarconvexsetssuchthatω(Hm)
v(Hm)→3/4asmgoestoinfinity,andthemaximumsizecliquesdonothaveacommonintersection.Proof.Takearegularm-gonandlettheconvexsetsbealltheconvexhullsof(m−1)-tuplesofverticesofthem-gonandtheconvexhullsofconsecutive⌈m
2⌉-tuplesofvertices.Thisconstructionhasω(Hm)=m+⌈m
2⌉−1sizedmaximumcliques,andv(Hm)=2mconvexsets,witharatioω(Hm)
v(Hm)=m+⌈m
2⌉−1
2m→3
4asmapproachesinfinity.
ThefollowingresultofKalaihelpstoprovesomepositiveresults.Theorem19(Kalai[Kal84]).Ifω(Hd)≥n−s,thenthereareatmost�d+sd�maximumsizecliques.Thecomplementofamatchingshowsthatthereisnoanalogousresultforgeneralhypergraphs.UsingtheaboveresultofKalai,wecouldprovethefol-lowingdependencebetweenthecliquenumberandη(Hd).Corollary20.Lettbeapositiveinteger.Ifω(Hd)≥n−sands≤Od,t(nt t+d),thenη(Hd)≤t.8
Proof.SupposewehaveahypergraphHdwithω(H)≥n−sandη(Hd)>t.Wewillbounds.Sinceη(Hd)>t,foranytverticesofthehypergraphtheremustbeamaxi-mumsizedcliquethatisdisjointfromit.ByTheorem19,thehypergraphhasatmost�d+sd�maximumsizecliques.Eachcliqueisdisjointfromatmost�st�t-setsofvertices.Thereare�nt�t-setsofvertices,hencewemusthave�d+sd��st�≥�nt�.Usingthewell-knownbounds(a b)b≤�ab�≤(ea b)bonbinomialcoefficients,afterasmallcomputationthisimpliess≥Ω(nt t+d).
Thisresultisanalogoustothelowerboundforfk,t(n)establishedintheprevioussection.WhilethegeneralTheorem15requiresthemorerestrictiveconditions=Od,t(n1
d+1)forintersectionhypergraphs,Corollary20permitss=Od,t(nt d+t).Fort≥2,thisrepresentsapolynomialimprovementintherequiredconditionforthisspecialcase.6ConcludingRemarksSeveraldirectionsforfurtherresearcharisefromourresults.First,itwouldbeinterestingtoimproveeitherthelowerortheupperboundsonfk,t(n)appearinginTheorem5.Determiningthecorrectasymptoticbehaviorofthisfunctionremainsopen.Anotherpromisingdirectionconcernsintersectionhypergraphsofconvexsets.WewereunabletoresolveQuestion1.Itwouldbeveryinterestingtoseeaspecialsubsetofthek-uniformhypergraphs,forwhichsomethingsimilartoConjecture3holds,andtheconvexsetintersectionhypergraphsmightbegoodcandidatesforthatresult.Forconvenience,werestatethequestionhere.Question1.Isittruethatforeverydimensiond,thereexistsaconstantcd<1suchthatifHdisad+1uniformintersectionhypergraphofconvexsetsinRd,andω(Hd)≥cdv(Hd),thenη(Hd)=1?Finally,weproposeanotherrestrictedclassof3-uniformhypergraphsforwhichapositiveresultmightstillhold.Recallthatatournamentisadirectedgraphwithexactlyonedirectededgebetweeneverypairofvertices.Atransitivetripleisaset{a,b,c}ofverticessuchthattheedges⃗ab,⃗ac,⃗bcareallpresent.ForatournamentT,letH(T)denotethe3-uniformhypergraphwhosehyperedgescorrespondpreciselytothetransitivetriplesofT.Thisleadstothefollowingquestion,analogoustothehypergraphversionsdiscussedearlier.Question2.Isthereaconstantc<1suchthatifω(H(T))>cnforatourna-mentT,thenthemaximumsizetransitivesubtournamentsofTcanbepiercedbyasinglevertex?9
References[AH74]H.L.AbbottandD.Hanson.Onfinite△-systems.DiscreteMathe-matics,8(1):1–12,1974.[BCW21]TolsonBell,SuchakreeChueluecha,andLutzWarnke.Noteonsun-flowers.DiscreteMathematics,344(7):112367,2021.[ER60]PaulErd¨osandRichardRado.Intersectiontheoremsforsystemsofsets.JournaloftheLondonMathematicalSociety,1(1):85–90,1960.[Haj65]Andr´asHajnal.Atheoremonk-saturatedgraphs.CanadianJournalofMathematics,17:720–724,1965.[Kal84]GilKalai.Intersectionpatternsofconvexsets.IsraelJournalofMathematics,48(2):161–174,1984.[P´al23]D¨om¨ot¨orP´alv¨olgyi.Intersectionofmaximumcliques,2023.14thEml´ekt´ablaWorkshop,V´ac,Hungary.[Wil84]RichardMWilson.Theexactboundintheerd¨os-ko-radotheorem.Combinatorica,4(2):247–257,1984.10

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

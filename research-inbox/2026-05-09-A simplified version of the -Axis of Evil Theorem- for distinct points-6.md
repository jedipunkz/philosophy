---
source: "https://arxiv.org/abs/1208.1695v1"
title: "A simplified version of the \"Axis of Evil Theorem\" for distinct points"
author: "Michela Ceria"
year: "2012"
publication: "arXiv preprint / math.AC"
download: "https://arxiv.org/pdf/1208.1695v1"
pdf: "https://arxiv.org/pdf/1208.1695v1"
captured_at: "2026-05-09T13:06:58Z"
updated_at: "2026-05-09T13:06:58Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche Beyond Good and Evil"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# A simplified version of the "Axis of Evil Theorem" for distinct points

- 著者: Michela Ceria
- 年: 2012
- 掲載情報: arXiv preprint / math.AC
- 情報源: [arxiv](https://arxiv.org/abs/1208.1695v1)
- ダウンロード: https://arxiv.org/pdf/1208.1695v1
- PDF: https://arxiv.org/pdf/1208.1695v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

Given a finite set $\mathbf{X}$ of distinct points, Marinari-Mora's 'Axis of Evil Theorem' states that a combinatorial algorithm and interpolation enable to find a 'linear' factorization for a lexicographical minimal Groebner basis $\mathcal{G}(I(\mathbf{X}))$ of the zerodimensional radical ideal $I(\mathbf{X})$. In this work we provide such algorithm, showing that it ends in a finite number of steps and that it actually provides the correct result. The 'Axis of Evil' algorithm takes as input the monomial basis of the initial ideal $T(I(\mathbf{X}))$ but its starting point is the (finite) Groebner escalier $N$ (obtained via Cerlienco-Mureddu correspondence) so we will also define the `potential expansion' 's algorithm, a combinatorical algorithm which computes the minimal basis from a finite Groebner escalier.

## PDF Text

arXiv:1208.1695v1 [math.AC] 8 Aug 2012
Asimpliﬁedversionofthe‘AxisofEvilTheorem’fordistinctpoints.MichelaCeriaUniversit`adegliStudidiTorino.michela.ceria@unito.itAbstractGivenaﬁnitesetXofdistinctpoints,Marinari-Mora’s‘AxisofEvilTheorem’statesthatacombinatorialalgorithmandinterpolationenabletoﬁnda‘linear’factorizationforalexicographicalminimalGroebnerbasisG(I(X))ofthezerodimensionalradicalidealI(X).Inthisworkweprovidesuchalgorithm,showingthatitendsinaﬁnitenumberofstepsandthatitactuallyprovidesthecorrectresult.The‘AxisofEvil’algorithmtakesasinputthemonomialbasisoftheinitialidealT(I(X))butitsstartingpointisthe(ﬁnite)GroebnerescalierN(obtainedviaCerlienco-Muredducorrespondence)sowewillalsodeﬁnethe‘potentialexpansion’’salgorithm,acombinatoricalalgorithmwhichcomputestheminimalbasisfromaﬁniteGroebnerescalier.Keywords:Groebnerbasis,Combinatorialalgorithm,Interpolation.1Introduction.
Marinari-Morain[10],[9],[11]gaveadeepdescriptionofthestructureofazero-dimensionalidealIdescribedbygivingitsMacaulaybasisB(I)([16]);inparticulartheyenhancedthedescriptionoftheGrobnerbasisofanidealinK[X,Y]givenbyLazardin[8]provingthatinarestrictedcasewhichincludestheradicalone,foreachmonomialτ:=Xd11···Xdnnbelongingtothemini-malbasisG(I)oftheinitialidealofI,itispossibletoproducelinearfactorsγmδτ:=Xi−f(X1,...,Xi−1),1≤m≤n,1≤δ≤dmsuchthatthepolynomi-alsfτ:=Qn m=1Qdmδ=1γmδτformaminimallexicographicalGroebnerbasisofI;eachsuchfactorswereobtainedbyproducinganappropriatedecompositionofthegivenMacaulaybasisB(I)=Fn m=1Fdmδ=1Smδ(τ)andinterpolatingoverthemonomialsetobtainedapplyingCerlineco-MuredduAlgorithmoverthesetoffunctionalsSmδ(τ).Suchalgorithmisreportedandprovedin[16];laterMorainaseriesoflecturenoteslabelledtherestrictionofthisdecompositionandinterpolationtothecaseofasetofdistinctpointsas‘Axis-of-Evil’theoremandgaveaprecisedescrip-tion,butnosimpleproof,oftheresultstatedin[16];S.SteidelimplementedtheprocedureinSingular[6],[18].
Wegiveheresuchexplicitalgorithmthat,givenaﬁnitesetXofdistinctpoints,providesacompletedecompositionX=Fn m=1Fdmδ=1Smδ(τ)onwhich,applyingCerlienco-Mureddualgorithmandinterpolation,producestherequiredlinear1
factorizationforalexicographicalminimalGroebnerbasisF={f1,...,fr}oftheidealI(X)andthusaverysimpleproofofthe‘Axis-of-Evil’theoreminthisparticularsituation.
ThisalgorithmarrangesthertermstibelongingtoG(I(X))withrespecttolex(t1≤...≤tr)andconstructsthefactorizationofeachfi∈Fthroughasuit-ableinterpolationonasubsetSmδ(ti)ofXdependingontheexponentsofthecorrespondingti.Moreprecisely,Cerlienco-MureddugiveanalgorithmthatenablestoﬁndtheGroebnerescalierN(I(X))andtheminimalbasisG(I(X))ofthemonomialidealT(I(X)).Sincethe‘AxisofEvil’algorithm’sstartingpointaretheelementsofXandthemonomialsoftheﬁniteGroebnerescalierN(computedusingCerlienco-Mureddualgorithm),butthealgorithmrequiresasinputthemonomialbasisofT(I(X)),wealsodeﬁnethe‘potentialexpansion’’salgorithm.IttakesNandcomputestheminimalbasis.InoteherethatMarinari-Moraexplicitlydeduced,astrivialcorollariesoftheir‘Axis-of-Evil’procedure,Lazardtheorem([8]),Eliminationtheorem([2]),Kalk-brenertheorem([13]),partofGianni-Kalkbrenertheorem([7],[12]);theyhow-everremarkedthat,havingbeingstronglyinﬂuencedbyGianni-Kalkbrnerre-sult,theycannotdismissthepossibilitythatGianni-Kalkbrennerargumentisanessentialtooloftheirproofofthe‘Axis-of-Evil’theorem.
2Notation.
LetP:=k[x1,...,xn]=Ld∈NPdbetheringofpolynomialsinnvariablesandcoeﬃcientsinthebaseﬁeldk.ForallM⊆P,Md=M∪Pdisitsdegreedpart.CallTthesemigroupofterms,generatedbytheset{x1,...,xn}:T:={xa11···xann,(a1,...,an)∈Nn}.Lettingα=(α1,...,αn)∈Nn,wewilloftenwritexαinsteadofxα11···xαnn.DeﬁnealsothesetT[m]:=T∩k[x1,...,xm]={xa11···xamm/(a1,...,am)∈Nm}.Foreachsemigroupordering<onT(i.e.atotalorderingsuchthatt1<t2⇒tt1<tt2,∀t,t1,t2∈T)wecanrepresentapolynomialf∈Pasalinearcombination(withcoeﬃcientsink)ofmonomialsarrangedw.r.t.<:f=Xt∈Tc(f,t)t=sXi=1c(f,ti)ti:c(f,ti)∈k∗,ti∈T,t1>...>ts.WewillcallT(f)=Lt(f):=t1theleadingtermoffandtail(f)=f−T(f)thetailoff.Wecanalsoexpressitinauniquewayasf=δXi=0gixi n∈k[x1,...,xn−1][xn],gi∈k[x1,...,xn−1],gδ6=0(whereδ:=degn(f)isthedegreew.r.t.xn).WedenoteLp(f):=gδ,theleadingpolynomialoff.2
Deﬁnition2.1.Foreachmonomialt∈Tandxj|t,theonlyu∈Tsuchthatt=xjuiscalledj-thpredecessoroft.AsubsetN⊆Tisanorderidealift∈N⇒s∈N∀s|t.LetN⊂Tanorderideal,AsubsetN⊆TisanorderidealifandonlyifT\N=Jisasemigroupideal(i.e.τ∈J⇒tτ∈J,∀t∈T).WesetN(J):=T\T(J)=N.ForasemigroupidealJ,G(J)denotesitsminimalbasisandG(J):={τ∈J|eachpredecessorofτ∈N(J)}=={τ∈T|N(J)∪{τ}orderideal,τ/∈N(J)}.ForallsubsetsG⊂P,wedeﬁneT{G}:={T(g),g∈G}andwecallT(G)thesemigroupideal{τT(g),τ∈T,g∈G},generatedbyT{G}.ForanyidealI/PconsiderthesemigroupidealT(I)=T{I},denotingbyabuseofnotationG(I)itsminimalbasisG(I)andtheborderidealofIB(I):={xht,1≤h≤n,t∈N(I)}\N(I)==T(I)∩({1}∪{xht,1≤h≤n,t∈N(I)}).Wewillalwaysconsiderthelexicographicorderinducedbyx1<...<xn,i.e:xa11···xann<xb11···xbnn⇔∃j|aj<bj,ai=bi,∀i>j.Thisisatermorder,thatisasemigrouporderingsuchthat1lowertoeveryvariableor,equivalently,itisawellordering.Lemma/Deﬁnition2.2.Wehave:1.P∼
=I⊕k[N(I)];2.P/I∼
=k[N(I)];3.∀f∈P,∃!g:=Can(f,I)=Pt∈N(I)γ(f,t,<)t∈k[N(I)],calledcanoni-calformoffwithrespecttoI,suchthatf−g∈I.Deﬁnition2.3.Givenatermorder<onT:1.aGroebnerbasisofIisasetG⊂IsuchthatT(G)=T{I},thatisT{G}generatesthesemigroupidealT(I)=T{I};2.aminimalGroebnerbasisisaGroebnerbasissuchthatdivisibilityrelationsamongtheleadingmonomialsofitsmembersdonotexist;3.theuniquereducedGroebnerbasisofIistheset:G(I):={τ−Can(τ,I):τ∈G(I)}.EachmemberofthereducedGroebnerbasishasaleadingtermwhichdoesnotdivideanymonomialofanothermember.3
LetX={P1,...,PN}⊂knbeaﬁnitesetofdistinctpointsPi:=(ai1,...,ain),i=1,...,N.WecallI(X):={f∈P:f(Pi)=0,∀i},theidealofpointsofX.Finally,wedeﬁnetheprojectionmaps:πm:kn→km(X1,..,Xn)7→(X1,...,Xm),πm:kn→kn−m+1(X1,..,Xn)7→(Xm,...,Xn)and,forP∈kn,X⊂kn,letΠs(P,X):={Pi∈X/πs(Pi)=πs(P)},Πs(P,X):={Pi∈X/πs(Pi)=πs(P)},extendingintheobviouswaythemeaningsofπs(d),πs(d),Πs(d,D),Πs(d,D)tod∈Nn⊂kneD⊂Nn⊆Nn.Withthesamenotationπmweindicatealsoπm:T∼
=Nn→Nm∼
=T[m]xa11···xann7→xa11···xamm.3Cerlienco-MuredduCorrespondence.
CerliencoandMureddu([3],[4],[5])providedanalgorithmwhichsolvesthefollowing
Problem:GivenﬁniteorderedsetofdistinctpointsX
:=(P1,...,PN)⊂kn;Pi:=(ai1,...,ain)computeamonomialbasis(w.r.t.thelexicographicorderinducedbyx1<...<xn)ofthequotientk[x1,...,xn]/I(X),whereXdenotesthesupport{P1,...,PN}ofX.~Moreprecisely,they•deﬁnetheoperatorΦ,associatingtoX
anorderedsetΦ(X
):=(d1,...,dN)⊂Nnsuchthat|Φ(X
)|=|X
|=Nandsuchthat,forallm<Nthesubset(d1,..dm)isexactlyΦ((P1,...,Pm)).•deﬁnetheσ-valuew.r.t.Xs=σ(P,X)ofapointP∈Kn\XasthemaximalintegersuchthatΠs−1(P,X)6=∅(byconvention,∀P,X,Π0(P,X)6=∅).ForP/∈X,theydeﬁnethesetΣ(P,X):={Pi∈X/πs−1(Pi)=πs−1(P),s=σ(P,X)}4
containingallthepointsofXhavingtheﬁrsts−1coordinatesequaltothoseofP/∈X.TheyextendthenotationtothecaseP=Pj∈X
inthefollowingway:σ(P,X
):=σ(P,{P1,..,Pj−1})Σ(P,X
):=Σ(P,{P1,..,Pj−1}).Remark3.1.Givenatermorder,amonomialbasisforA:=k[x1,...,xn]/I(X),[xi1],...,[xiN],withxi1...xiNiscalledminimalwithrespecttothetermorderif,foreveryothermonomialbasis[xi0
1],...,[xi0
N],withxi0
1...xi0
NfortheAitholds∀j=1,...,N,xijxi0
j.In[3],theystatethatthecomputedmonomialbasisistheminimalone.Proposition3.2.([3])LetD:=Φ(X).Then{[xd]/d∈D}isamonomialbasisfork[x1,...,xn]/I(X).Suchamonomialbasisisminimalwithrespecttothegiven<.OncetheGroebnerescalierNisknown,itisverysimpletocomputetheminimalbasisGofT(I(X))=T\N.GiventhesetX,theﬁrststeptocomputethelinearfactorizationofaminimalGroebnerbasiswillbetoapplyCerlienco-MureddualgorithmtoXandcomputeN,inordertoobtainG.4Thepotentialexpansion’salgorithm.
Considerthepolynomialringk[x1,...,xn]withusualordering<.GivenaﬁnitesetofdistinctpointsX={P1,...,PN},considertheidealI(X)/k[x1,...,xn]whichisradicalandzerodimensional,soitsGroebnerescalierNisaﬁniteset.WewillcomputetheminimalmonomialbasisGofT(I(X)),giventheGroeb-nerescalier.Thealgorithmactuallyprovidescorrectresultsirrispectiveofthegiventermordering,butsinceweuseCerlienco-Muredducorrespondence,wewillconsideronlyourlexorder.
Inordertomakethereasoningclear,wewillrepresentthemonomialsusingthesamediagramsintroducedin[15]tostudypropertiesofBorelideals.ApplyCerlienco-MuredducorrespondencetoXinordertohaveN(X)={τ1,...,τN}.Itiswellknown(see,forinstance[16])that|N(X)|=|X|.WeﬁrstdeﬁnethepotentialexpansionofasubsetH⊂T,fromwhichthealgortihmbearsitsname.
Deﬁnition4.1.LetH⊆Tjforsomej∈N∗wesetC(0)(H):=Hand,foralll∈N∗C(l)(τ)=Tj+l\{x1,...,xn}·(Tj+l−1\C(l−1)(H)).WethenslicetheGroebnerescalierbydegree,havingN0,N1,···Nh,wherehisthemaximaldegreeoftermsappearinginN.TheminimalmonomialbasisG(I(X))willhaveatmostdegreeh+1.Asamatteroffact,ifτ∈Gwithdeg(τ)=d>h+1itspredecessorswillbelongtoNandhavedegreed−1≥h+1whichisimpossible.5
Algorithm1Cerlienco-Mureddualgorithm.
1:procedureCeMu(X
)→Φ(X
)2:ifN=1then3:Φ(X
):={(0,...,0)}.4:endif5:if1<Nthen.supposetoknowbyinductionhypothesisΦ((P1,...,PN−1))=(d1,...,dN−1)andlookfordN=Φ(PN).6:s=σ(PN,X
).7:fori=nto1do8:ifi>sthen9:dNi=0.10:endif11:ifi=sthen12:m,(1≤m≤n),maximals.tπs−1(Pm)=πs−1(PN),πs+1(dm)=(0,...,0)=πs+1(dN)..Pmistheσ-antecedentofPNw.r.t.(P1,...,PN−1),Φ((P1,...,PN−1)).13:dNs=dms+1.14:endif15:ifi<sthen.weuseinductionhere.16:W(PN,X
):={P∈X
|Φ(P)=d=(∗,...∗,dNs,0,...,0),}={Pj1,...,Pjr}.17:Q:=πs−1(W(PN,X
))..|Q|=|W(PN,X
)|=r<N.Ifh<r=|W(PN,X
)|,thenπs−1(Pjh)6=πs−1(PN).Moreover,sinceΦisinductive,ifh<k≤rthenπs−1(Pjh)6=πs−1(Pjk).18:πs−1(dN)=fdr..Bytheinductionhypothesis,Φ(Q
)=(fd1,..,gdr)and∀1≤i<r,fdi=πs−1(dji).19:break.20:endif21:endfor22:endif23:returnΦ(X
).24:endprocedure
6
ThecomputationofGisperformedasfollows.ConsiderTi∀i=0,...,h+1:itiswellknownthat|Ti|= n+i−1n−1.Foreachi,deﬁneGeni(I):={t∈G(I)|deg(t)≤i}.SinceIisaproperideal,Gen0(I)=∅.LeththeminimalisuchthatGenh(I)6=∅,∀i≥1Geni+h=Genh+i−1∪(Th+i\(Nh+i∪h+i−1[j=h+1C(h+i−j)(Gj))).Wethenhave
Algorithm2Thepotentialexpansion’salgorithm.
1:procedurePotExp(N(I))→I.Iisexpressedusingitsminimalbasis.Require:N=[N0,...,Nh,Nh+1],suchthatNh+1=∅.2:C=[∅]..thepotentialexpansion’slist.3:Gen=∅.4:I=(0).5:fori=0toh+1do6:d= n+deg(Ni[1])−1n−1−|Ni∪C[i]|.7:ifd=0then.nonewgenerators.8:C[i+1]=PotentialExpansion(C[i]).9:Geni=(0)10:else.addingnewgenerators.11:Geni=Ti\(Ni∪C[i]).12:C[i+1]=PotentialExpansion(Geni∪C[i]).13:I=I+Geni.14:endif15:endfor16:returnI17:endprocedure
ThealgorithmusesasubroutinePotentialExpansionsuchthatPotentialExpansion(F)=C(1)(F).WewillalsohaveasubroutineﬁndingTh+i\(Nh+i∪Sh+i−1j=h+1C(h+i−j)(Gj)).WLOGwewillthinkthatthesetsTh+iandNh+i∪Sh+i−1j=h+1C(h+i−j)(Gj)areorderedw.r.t.thesameordering,sinceitisenoughtoperformamergingwiththeGroebnerescalierandthepotentialexpansionpreviouslyordered.Allthesestepsend:thesubroutineﬁndingthecomplementarycanbedevelopedperformingalooponthetwoorderedlistsA:=Ti=[a1,..,am],m≥nandB:=Ni∪C(i)=[b1,...,bn](usingtwoindicesi,j),keepinginmindthatB⊆AorB=AandthatB[j]≥A[i]ateverystep.Startwithb1:ifb1=a1weseti=i+1;j=j+1.Ifweﬁndai6=bjforacertaincouple(i,j),weputA[i]inthecomplementaryandi=i+1withoutmodifyingj.Example4.2.TherearesituationsinwhichNcontainsmonomialsofdegreeatmosth,butalsotheminimalbasissharesthesameproperty.TakeI=(x3,y2,z2,xy)/k[x,y,z],whoseGroebnerescalieris:7
N0={1}N1={x,y,z}N2={yz,xz,x2}N3={x2z}:
Themonomialbasisdoesnotcontainelementsofdegree4.WecallGithesetofi-degreeelementsoftheminimalbasisandIthemono-mialidealwewanttoﬁnd.
Lemma4.3.Foralli=0,...,h+1Ti\(Ni∪i−1[j=1Ci−j(Gj))=Gi.Proof:TheinclusionTi\(Ni∪Si−1j=1Ci−j(Gj))⊇Giistrivial,soweonlyproveTi\(Ni∪Si−1j=1Ci−j(Gj))⊆Gi.Considerτ∈Ti\(Ni∪Si−1j=1Ci−j(Gj)).Clearlyτ∈I.Letσtheithpredecessorofτ;ifσ∈I,∃t∈Gwithσ=t·mforasuitablem∈T.Thenτ=t·m·xii.e.τ∈Si−1j=1Ci−j(Gj).Thislemmaassuresthattheresultobtainedviathepotentialexpansion’salgo-rithmiscorrect.
5TheAxisofEvilAlgorithm.
A0−dimensionalradicalidealI/PiscompletelydeterminedifweknowthesetV(I)ofitszeros.ConsideraﬁnitesetofdistinctpointsX={P1,...,Pr};wewilldenoteindiﬀer-entlytheGroebnerescalieroftheidealI(X)withN(I(X))orN.AvariationofCerlienco-Mureddualgorithm([3])allowsustoﬁnda‘linearfactorization’foreveryelementofalexicographicminimalGroebnerbasisinthesenseoftheTheorem5.1.Letti:=xd11···xdnn,i=1,...,rbethegeneratorsoftheminimalbasisofT(I),whereIisa0−dimensionalradicalideal.Acombinatoricalalgorithmandinterpolationallowustodeducepolynomialsγmδi=xm−gmδi(x1,...,xm−1),∀i,m,δ,with1≤i≤r,1≤m≤n,1≤δ≤dmsuchthatfi=YmYδγmδi∀i8
wherefi,i=1,...,rarethepolynomialsformingaminimalGroebnerbasisofIwithrespecttothelexicographicorderinducedbyx1<...<xn.Inthatalgorithmwewillusetheprojections,aswedeﬁnedinsection3.TheAxisofEvilalgorithmworkstheninthefollowingway:•considerτj:=xd11···xdnn∈G.Therequiredpolynomialf=τj+tail(f)isfactorizedinPn i=1difactors:d1polynomialswhoseleadingtermisx1,d2polynomialssuchthattheirleadingtermisx2andsoon;•considerthemonomialsxa11xd22···xdnnsuchthata1<d1;•everysuchmonomialisassociated,viaCerlienco-MuredduCorrespondence,toapointofoursetX.Projectthesepointswithrespecttotheﬁrstco-ordinate,obtainingd1numbersy1,...,yd1;•x1−yi,i=1,...,d1aretheﬁrstd1factors;•constructthesubsetD20ofXcontainingallthepointsinwhichtheprod-uct(x1−y1)···(x1−yd1)doesnotvanish.IfitisemptythenstopandconsiderthenextmonomialinG;otherwisecontinueasfollows;•ﬁndthesetN2(τj)ofallmonomialsinT[2]suchthatxα11xα22<xd11xd22;•splittheelementsofN2(τj)withrespecttotheexponentsofx2andcon-struct,viaCerlienco-Muredducorrespondence,theset{Φ−1(vxd2−δ2xd33···xdnn)/v∈T[1],vxd2−δ2∈N2(τj)}•intersecttheprevioussetwithD20,projecttheresultingsetofpoints(A2δ(τj))withrespecttotheﬁrsttwocoordinatesandapplyCerlienco-MuredduCorrespondence,obtainingasetE2δτj;•interpolateoverA2δ(τj),ﬁndingd2factorswhoseleadingtermsareallequaltox2.ThemonomialsofE2δτjaretheonesappearinginsuchfactorization;•updatethesetofpointsinwhichthecurrentpolynomialdoesnotvanishandstopifitisempty;•repeatthesestepslettingallthevariablesvaryonebyone;•repeatallthestepsforallτi∈G.Remark5.2.Givenτj=xd11···xdnn∈G,everyvariablexiwillappearonlyditimesintheexecutionofthealgorithm.
Remark5.3.ThesetsNm(τj):={ω∈T[m],τj>ωxdm+1m+1···xdnn∈N}(inparticularform=1wehaveN1(τj):={xi
1/i<d1})areconstructedinordertodetermineinwhichpointsitisnecessarytointerpolate.
Sinceforµ>τjtheCerlienco-MuredducorrespondenceprovidesapointPµ0suchthat∃k∈{1,...,n}:πk(Pµ)=πk(Pµ0),inordertoobtainpolynomialsvanishingonallthepoinstofXitisnotnecessarytointerpolateinthewholeΦ−1(N)asitsuﬃcestoconsideronlythosecorrespondingtoµ∈Nwithµ<τj.9
Algorithm3TheAxisofEvilalgorithm.
1:procedureAoE(X,G(I(X)):={τ1,...,τr})→R.RcontainsafactorizedminimalGroebnerbasisofI.Require:theelementsG(I(X))areinincreasingorderw.r.tthelexicographicalorderw.r.t.x1<...<xr.2:R=∅3:fori=1tordo4:N1(τj):={xi
1/i<d1}={ω∈T[1],τj>ωxd22···xdnn∈N}5:A1(τj):={Φ−1(xi
1xd22···xdnn)/i<di}⊂X.6:B1(τj):=π1(A1(τj))⊂k.7:γ1τj:=Qa∈B1(τj)(x1−a).8:form=2tondo9:ζmτj:=Qm−1ν=1γντj.10:Dm0:={Pi∈X/ζmτj(Pi)6=0}.11:if|Dm0|=0then12:R=[R,ζmτj].13:break.14:endif15:Nm(τj):={ω∈T[m],τj>ωxdm+1m+1···xdnn∈N}.16:forδ=1todmdo17:Amδ(τj):={Φ−1(vxdm−δmxdm+1m+1···xdnn)/v∈T[m−1],vxdm−δm∈Nm(τj)}∩Dm(δ−1)(τj).18:Emδ(τj):=Φ(πm(Amδ(τj))).19:γmδτj:=xm+Xω∈Emδ(τj)c(γmτj,ω)ω,suchthatγmδτj(P)=0,∀P∈Amδ(τj).20:ξmδ:=Qm−1ν=1γντjQδ
d=1γmdτ.21:Dmδ(τj):={Pi∈X/ξmδ(Pi)6=0}⊆X22:if|Dmδ(τj)|=0then23:R=[R,ξmδ].24:break.25:endif26:endfor27:γmτj:=Qδγmδτj.28:endfor29:endfor30:returnR.31:endprocedure
10
Remark5.4.Thetermssmallerthanτjmentionedbeforearefoundreleasingallthevariablesonebyone.
Imaginethemonomialsink[x1,...xn]aspointsinkn,identifyingeverytermtothen-upleofitsexponents.Sowecan‘draw’theminan-dimensionalspaceandwecanthinkourrealeasingsasanincrementbyoneofthe‘directions’wherewecanmovethere.
WepointoutthatNm(τj)⊆Nh(τj)form≤h.Ifω∈Nm(τj),τj>ωxdm+1m+1···xdnn∈N;asωxdh+1h+1···xdnn|xdm+1m+1···xdnnwehaveωxdh+1h+1···xdnn∈Nandωxdh+1h+1···xdnn≤xdm+1m+1···xdnn<τj.Ateachstepwewillcountoutallthepointsinwhichthepolynomialalreadyvanishesandwewillstopthecomputationwhenthecurrentfactorizedpolyno-mialvanishesonthewholeX.Wewillseeanexampleofitlater.
Remark5.5.Ifthenumberofreleasedvariablesis>1,wealsomustsplittheobtainedmonomialsregardingtheexponentofthemaximalvariable.
Considerthenthelooponδand,inparticular,theset:Cmδ(τj):={Φ−1(vxdm−δmxdm+1m+1···xdnn)/v∈T[m−1],vxdm−δm∈Nm(τj)}.WeintersectCmδ(τj)withthesubsetofXcontainingthepointsnotvanishingthecurrentfactorizedpolynomial.
Wecaneasilynoticethat,performingthealgorithm,weonlycomputethesetsCm1(τj),...,Cmdm(τj),butinNm(τj)therearealsomonomialsω=xa11···xam−1m−1xdmmsuchthatτj>ωxdm+1m+1···xdnn∈N,whichwouldbegeneratedconsideringδ=0.Theyarenotconsideredinthealgorithmbecausetheyarerelatedtomonomialsexaminedinthepreviousstep:=xa11···xam−1m−1∈Nm−1,sothecorrespondingpointshavealreadybeentreated.Takingδ=0,..,dm,thesetsCmδ(τj)formapartitionofNm(τj)basingonthedegreeofxm.Asamatteroffact,inordertohaveω∈Nm(τj)wemusthaveτj>ωxdm+1m+1···xdnn,whereωxdm+1m+1···xdnn∈N,thentheexponentofxmwillbetheﬁrstcheckedinthelexicographictestandsoitwillbelimitedbydm.Accordingtothevaluesofthisexponent,theonesassociatedtosmallervariableswillvary.
Remark5.6.Atthebeginningofthealgorithm,weimposedthemonomialsτj,j=1...,rtobeinincreasingorderwithrespect<.Thestepsmadebythealgorithmoneachτjaretotallyindependentbothonthosemadeandonthosetobemadeonamonomialτk(itisindiﬀerentwhetherj≷k)belongingtoG,sowewillobtainthesamefactorizationsevenifwelaunchthecomputationonalistofunorderedmonomials.
Clearly,theresultofourcomputationisnotthereducedGroebnerbasisofthegivenideal,itisonlyoneoftheminimalGroebnerbasesbutwecanobtainthereducedGroebnerbasisviasimplereduction.
Wedecidedtoputthemonomialsinsuchanorderbecausewewanteverypolynomialtobereducedwithrespecttothe‘previous’ones.IffjisoneofourresultingpolynomialsandLt(fj)=τj,thepolynomialsutilizabletoreducefj(thepreviousones)mustbenecessarilyallandonlytheoneshavingasleadingtermselementsinGlowerthanthegivenτj.11
Thealgorithmterminatesbecauseitworkson:1.pointsintheﬁnitesetX;2.monomialsτ∈G(theyareinaﬁnitenumber,[16]);3.aﬁnitesetofvariables.Letusstudythecorrectnessofthealgorithm.
Lemma5.7.ThefactorizedpolynomialsobtainedfromouralgorithmvanishonallthepointsofthesetX.Proof:Supposewewanttoconstructγτwithτ=xα11···xαnn.Letµ=xβ11···xβnn,correspondingtoapointPµ∈XthroughCerlienco-MuredduCorrespondence.
Letµ<τ,thenatleastoneoftheexponentsofthevariablesappearinginµislowerthanthecorrespondinginτ,sayβi<αi,soµislinkedtoanelementofNi(τ)andsoitcan,alternatively:•belongtoAiδ(τ)forsomeδ;•besuchthatthecorrespondingpointalreadyannihilatesthepolynomialfound.Ifµ>τ(sinceτ/∈N,itissurelyimpossiblethatτ=µ)thentherewillbeapointPµ0suchthatπj(Pµ)=πj(Pµ0),correspondingtoaµ0<τ.Wethenuseµ0andwecomebacktotheprevioussituation.Corollary5.8.TheidealgeneratedbythesepolynomialsisexactlyI(X).Proof:Bythepreviouslemma,thepolynomialsvanishonallthepointsofthesetXandtheequalitycomesoutbyreasonsofmultiplicityTheresultingpolynomialsformaminimalGroebnerbasisbecause:•theyvanishonallthepointsofX;•theirheadsformexactlyG(I(X)).NoticethatwecanobtainthecurrentinterpolatingpolynomialapplyingMoelleralgorithmtotheprojectionthroughπmofthepointsofthecurrentAmδ(τ)([14]).
Example5.9.LetX:={(4,0,0),(2,1,4),(2,4,0),(3,0,1),(2,1,3),(1,3,4),(2,4,3),(2,4,2),(1,0,2)}.P1:=(4,0,0):itisasinglepoint,soΦ({(4,0,0)})=(0,0,0)P2:=(2,1,4):s=1,m=1,(1,0,0)P3:=(2,4,0):s=2,m=2,(0,1,0)P4:=(3,0,1):s=1,m=1,(2,0,0)P5:=(2,1,3):s=3,m=2,(0,0,1)12
P6:=(1,3,4):s=1,m=4,(3,0,0)P7:=(2,4,3):s=3,m=3,W={(2,1,3),(2,4,3)},t7=(0,1,1)P8:=(2,4,2):s=3,m=7,(0,0,2)P9:=(1,0,2):s=2,m=6,W={(2,4,0),(1,0,2)},t9=(1,1,0).ThenN:={1,x1,x2,x2
1,x3,x3
1,x2x3,x2
3,x1x2}andsowecaneasilyobtainG={x4
1,x2
1x2,x2
2,x1x3,x2x2
3,x3
3}.ThemonomialsbelongingtoGareexactlytheinputfortheAxisofEvilalgo-rithmandtheyarealreadyorderedwithrespecttoourordering:startingwithτ=x4
1weobtain
N1(τ)={1,x1,x2
1,x3
1};A1(τ)={(4,0,0),(2,1,4),(3,0,1),(1,3,4)}:thesearethecorre-spondingpointsviaCerlienco-Mureddu
Correspondence;
B1(τ)={4,2,3,1}γ1τ=(x1−4)(x1−2)(x1−3)(x1−1):allthelinearfactorsareonlydepending fromx1arecomputedinthesametime.m=2:ζ2τ=γ1τD20(τ)=∅:stophereobtaining,asﬁrstresult,apolynomialhavingaslead-ingtermanelementofG(whiletheothermonomialsbelongtoN)andbelongingtoI(X)sinceitvanishesineverypointofX(soanelementofourminimalGroebnerbasis).
τ=x2
1x2N1(τ)={1,x1};A1(τ)={(2,4,0),(1,0,2)};B1(τ)={2,1}γ1τ=(x1−2)(x1−1)
m=2:ζmτ=γ1τD20(τ)={(4,0,0),(3,0,1)}Wecannotstophere,sincetheobtainedpolynomialdoesnotvanishatall thepointsanditsheadisdiﬀerentfrom
τ∈G.N2(τ)={1,x1,x2
1,x3
1,x2,x1x2};doingso,weﬁndallthemonomialsofthe previousstep(wecomputedtheircorrespondingpoints)andsomenewones.δ=1:A21(τ)={(4,0,0),(3,0,1)}=D20Themonomialsvxdm−δmare1,x1,x2
1,x3
1,correspondingtothepointsP1,P2,P4,P6.ThepolynomialalreadyvanishesonP2,P6,soweconsideronlytheremainingtwopoints.E21(τ)={1,x1}.γ21τ=x2;ξ21=γ1τγ21τ=(x1−2)(x1−1)x2;D21(τ)=∅.13
Remarkthatγ2τisactuallyγ21τ.τ=x2
2N1(τ)=∅;A1(τ)=∅;B1(τ)=∅m=2:D20(τ)=X
N2(τ)={1,x1,x2
1,x3
1,x2,x1x2};δ=1:A21(τ)={(2,4,0),(1,0,2)};E21(τ)={1,x1};γ21τ=x2−4x1+4ξ21=γ1τγ21τ=x2−4x1+4;D21(τ)={(4,0,0),(2,1,4),(3,0,1),(2,1,3),(1,3,4)};δ=2:A22(τ)={(4,0,0),(2,1,4),(3,0,1),(1,3,4)}Thetermsvxdm−δmare1,x1,x2
1,x3
1andtheycorrespondexactlyto
P1,P2,P4,P6.E22(τ)={1,x1,x2
1,x3
1};γ22τ=2x2−x2
1+7x1−12;ξ22=(x2−4x1+4)(2x2−x2
1+7x1−12)D22(τ)=∅;τ=x1x3
N1(τ)={1};A1(τ)={(2,1,3)};B1(τ)={2}γ1τ=(x1−2)m=2:N2(τ)={1}.D20(τ)={(4,0,0),(3,0,1),(1,3,4),(1,0,2)}δ=1:D21(τ)=D20(τ);m=3:N3(τ)={1,x1,x2,x2
1,x3,x3
1,x1x2};ζmτ=(x1−2);D30(τ)={(4,0,0),(3,0,1),(1,3,4),(1,0,2)};
δ=1:A31(τ)={(4,0,0),(3,0,1),(1,3,4),(1,0,2)}Thetermsare1,x1,x2
1,x3
1,x2,x1x2,correspondingtoP1,P2,P3,P4,P6,P9,butwecanneglectP2,P3.E31(τ)={1,x1,x2
1,x2};14
γ31(τ)=6x3−4x2+x2
1−x1−12;ξ31=(x1−2)(6x3−4x2+x2
1−x1−12);D31(τ)=∅.Thedesiredpolynomialisγ3τ=γ31(τ).τ=x2x2
3N1(τ)=∅;A1(τ)=∅;B1(τ)=∅m=2:
N2(τ)={1};D20(τ)=X;δ=1:A21(τ)={(2,4,2)};E21(τ)={1};γ21τ=x2−4ξ21=x2−4;D21(τ)={(4,0,0),(2,1,4),(3,0,1),(2,1,3),(1,3,4),(1,0,2)};m=3:ζ3τ=x2−4D30(τ)=D21(τ);
N3(τ)=N(X);δ=1:A31(τ)={(2,1,3)}.E31(τ)={1};γ21τ=x3−3ξ31=(x2−4)(x3−3);D31(τ)={(4,0,0),(2,1,4),(3,0,1),(1,3,4),(1,0,2)};δ=2:A32(τ)=D31(τ);E32(τ)={1,x1,x2
1,x3
1,x2};γ32τ=x3−4x2−5x3
1+41x2
1−96x1+48;ξ32=(x2−4)(x3−3)(x3−4x2−5x3
1+41x2
1−96x1+48);D32(τ)=∅;γ3τ=(x3−3)(x3−4x2−5x3
1+41x2
1−96x1+48);τ=x3
3N1(τ)=∅;A1(τ)=∅;B1(τ)=∅m=2:D20(τ)=X;N2(τ)=∅;δ=1:A21(τ)=∅;D21(τ)=X;15
m=3:D30=X;
N3(τ)=N(X);δ=1:A31(τ)={(2,4,2)};E31(τ)={1};γ31τ=x3−2;ξ31=x3−2;D31(τ)={(4,0,0),(2,1,4),(2,4,0),(3,0,1),(2,1,3),(1,3,4),(2,4,3)};δ=2:A32(τ)={(2,1,3),(2,4,3)};E32(τ)={1,x2};γ32τ=x3−3;ξ32=(x3−2)(x3−3);D32={(4,0,0),(2,1,4),(2,4,0),(3,0,1),(1,3,4)};δ=3:A33(τ)=D32;E33(τ)={1,x1,x2
1,x3
1,x2};γ33τ=6x3+8x2−5x3
1+35x2
1−54x1+24;ξ33=(x3−2)(x3−3)(6x3+8x2−5x3
1+35x2
1−54x1+24);D33(τ)=∅;Therequiredpolynomialisγ3τ=(x3−2)(x3−3)(6x3+8x2−5x3
1+35x2
1−54x1+24).ThenourminimalGroebnerbasisoftheidealassociatedtoXwithrespecttothegivenorderis:G(I(X))=nx4
1−10x3
1+35x2
1−50x1+24,x2x2
1−3x2x1+2x2,x2
2−2x2x1−x2+2x3−16x2
1+38x1−24,x3x−2x3−2
3x2x1+4
3x2++1
6x3−1
2x2
1−5
3x1+4,x2
3x2−4x2
3−7x3x2+28x3+8
3x2x1++20
3x2−16
3x3+48x2−344
3x1+32,x3
3−5x2
3+8
3x3x2−14
3x3−16
9x2x1−40
9x2+73
9x3
1−197
3x2
1+1358
9x1−72o,obtainedbyourpolynomialsbythereductionsstatedintheAxisofEvilTheorem.Finally,weremarkthat:1.letτj=xd11···xdnn∈G.ThepolynomialwearelookingformustcontainexactlyPn i=1difactors.Itisimpossiblethatthealgorithmstopsbefore,soitisimpossiblethatapartialproductvanishesonthewholeX.Infact,ifso,therewouldbeapolynomialf∈IsuchthatT(f)/∈(G)(weknowtheminimalbasisGbeforestartingtheAxisofEvilprocess);2.ifweotainafactorizedpolynomialfsuchthatitsleadingtermT(f)belongstotheminimalbasisG,thenfvanishesoverallX,becauseof5.7.16
Example5.10.Considerthefollowingideal,givenwithitsprimarydecompo-sition:
J:=(x2
1,x2+x1,x3)∩(x2
1,x2−x1,x3−1)==(x2
1,x1x2,x2
2,x1x3−1
2x1−1
2x2,x2x3−1
2x1−1
2x2,x2
3−x3)/C[x1,x2,x3].Callitsgeneratorsf1,...,f6,consideringtheminthecorrectorder.Itis0-dimensionalbecausex2
1,x2
2,x2
3∈In(J)(see[16]),butitisnotradical:itsradicalis√
J=(x2,x2
3−x3,x1).ForsuchanidealtheAxisofEvildoesnothold.
Considerthepolynomialf4=x1x3−1
2x1−1
2x2.BytheAxisofEviltheorem(5.1),itsfactorizationshouldbeoftheform:(x1+...)(x3+...)andweshouldhavex1x3−1
2x1−1
2x2+Px2
1+Qx1x2+Rx2
2,P,Q,R∈C[x1,x2,x3],sincewecanonlyreducedeletingthemultiplesofx2
1,x1x2,x2
2,inordertoobtainf4.Inordertohavethecorrectproductwemusthave−1
2x2init.Wecannotobtainitthroughreductions,sotheonlychanceisthatwehaveaproductoftheformk∗hx2,withh,kconstantssuchthathk=−1
2,inparticularbothdiﬀerentfrom0.Apriori,wecanhavetwopossibilities:•(x1+k+...)(x3+hx2+...);•(x1+hx2+...)(x3+k+...).Thesecondoneisimpossible:thepolynomialhavingx1asheadcannotcontainvariablesgreaterthanx1,soweconsideronly:(x1+k+...)(x3+hx2+...).Wewillthenobtainx1x3+hx1x2+kx3−1
2x2+...Wecandeletethetermx1x2butitremainskx3whichcannotbereduced.6Corollaries.
Weenumerateheresomefamoustheoremswhichcanbeeasilyprovedascorol-lariesoftheAxisofEvilTheorem.Formoredetailssee,forexample,[16].Hereweprovidethegeneralstatementsoftheseresults,butclearlytheycanonlybededucedunderthehypothesisoftheAxisofEviltheorem
TheﬁrstoneisLazardStructuralTheorem,whichdescribesthestructureofaminimallexicographicalGroebnerbasisofanI/k[x1,x2].TheoriginalproofconsidersP=k[x1,x2]=k[x1][x2]anditisbasedonthefactthatk[x1]isaPrincipalIdealDomain(PID).Norton-S˘al˘agean[17]reformulateditusing,moregenerally,R[x]withRPIR.Webrieﬂyrecallthefollowing17
Deﬁnition6.1.Thecontentrf∈R,withRPIR,ofapolynomialf(x)∈R[x]istheGCDofitscoeﬃcients.Apolynomialf(x)∈R[x]iscalledprimitiveifrf=1.Theprimitivepartoff(x)∈R[x]isthepolynomialp0(x)∈R[x]suchthatf(x)=rfp0(x).LetRbeaPIR,P:=R[x].LetI/PeF:={f0,...,fs}aminimalGroebnerbasisofIorderedinsuchawaythat,calledd(i):=deg(fi),∀i,0≤i≤sd(0)≤...≤d(s).Deﬁnethenci=lc(fi),ri∈R\{0}epi∈Ptheleadingcoeﬃcient,thecontentandtheprimitivepartoffi,forall1≤i≤n.Theorem6.2(Lazard).If,moreover,RisaPID,then:•f0=PG1···Gs+1;•fj=PHjGj+1···Gs+1,1≤j≤s.where1.d(1)<...<d(s);2.Gi∈R,1≤i≤s+1issuchthatci−1=Gici3.P=p0(theprimitivepartoff0∈R[x]);4.Hi∈R[x]isamonicpolynomialofdegreed(i)inx,foralli;5.foralliwehaveHi+1∈(G1···Gi,H1G2···Gi,...,Hi−1Gi,Hi);6.ri=Gi+1···GsTheorem6.3(Norton-S˘al˘agean).Withthepreviousnotation,eachpi∈(fj,j<i):ri.Infact,wehaveri=Qn−1m=1Qdmδ=1γmδtiandpi=Qdnδ=1γnδti.Thesecondwell-knownresultwhichcanbestraightforwardlyderivedfromtheAxisofEvilTheoremisthewellknownEliminationTheorem(see[2]fordetails)Theorem6.4([19]).LetI/k[x1,...,xn]anideal,takethelexicographicalor-deringinducedbyx1<...<xnandcallIjthej-theliminationidealIj=I∩k[x1,...,xj].LetGbeaGroebnerbasisofI,thenGj=G∩k[x1,...,xj]isaGroebnerbasisofIj.Thefollowingresult,Kalkbrenertheorem([13],[16]),isanotherconsequenceoftheAxisofEvilTheoremanditisastrongercharacterizationofthelexico-graphicalordering.
ForeachsubsetL⊂k[x1,...,xn],i=1,...,n,∀δ∈NsetLiδ={p∈L,|p∈k[x1,...,xi],degi(p)≤δ}andLpi,δ={Lp(p),p∈Li,δ}.18
Theorem6.5(Kalkbrenner).Withthepreviousnotations,consideredanidealI/k[x1,...,xn]andaGroebnerbasisGofit,theseformsareequivalent:•GisaGroebnerbasisofIw.r.t,thelexicographicalorder<inducedbyx1<...<xn;•Lpi,δ(G)isaGroebnerbasisofLpi,δ(I),i=1,...,n,∀δ∈N.LetusnowmentionGianni-Kalkbrenertheorem,whosesituationisabitmorecomplicated(see[12],[7],[16]).
Theorem6.6(Gianni-Kalkbrener).LetI/k[x1,...,xn]anidealandGw.r.tthelexicographicalorder<inducedbyx1<....<xn.AsbeforewedeﬁnealsoGd=G∩k[x1,...,xd].Considerα=(b1,...,bd)∈V(Id)anddeﬁnetheprojectionmapΦα:k[x1,...,xn]→k[xd+1,...,xn]f(x1,...,xn)7→f(b1,...,bd,xd+1,...,xn).LetσbetheminimalvaluesuchthatΦα(Lp(gσ))6=0andj,δthevaluessuchthatgσ=Lp(gσ)xδ+1j+...∈k[x1,...,xj]\k[x1,...,xj−1].Then1.j=δ+12.∀g∈Gd,Φα(g)=0;3.∀g∈Gd+δ,Φα(g)=0;4.Φα(gσ)=gcd(Φα(g),g∈Gd+1)∈k[xd+1];5.∀b∈k,(b1,...,b2,b)∈V(Id+1)⇔Φα(gσ)(b)=0.Clearly(1−3)areessentiallyacorollaryoftheorem6.3;ontheotherside,(4−5)apparentlycannotbededucedfromtheAxisofEvilTheorem.7Acknowledgement.
IwishtothankM.G.Marinariforherhelp,ideasandsuggestionswhilestudyingthissubject.
References[1]M.E.Alonso,M.G.Marinari,T.Mora,ThebigMotherofallDualities2:MacaulayBases,ApplicableAlgebrainEngineering,CommunicationandCom-putingarchiveVol.17Issue6,November2006,409−451.[2]BuchbergerB.,Gr¨obnerBases:AnAlgorithmicMethodinPolynomialIdealTheory,inBoseN.K.(Ed.)MultidimensionalSystemsTheory(1985),184–232,Reider19
[3]L.Cerlienco,M.Mureddu,Algoritmicombinatoriperl’interpolazionepolino-mialeindimensione≥2,preprint(1990).[4]L.Cerlienco,M.Mureddu,Fromalgebraicsetstomonomiallinearbasesbymeansofcombinatorialalgorithms,DiscreteMath.139,73−87.[5]L.Cerlienco,M.Mureddu,MultivariateInterpolationandStandardBasesforMacaulayModules,J.Algebra251(2002),686−726.[6]W.Decker,G.-M.Greuel,G.Pﬁster,H.Sch¨onemann:Singu-lar3-1-4—Acomputeralgebrasystemforpolynomialcomputations.http://www.singular.uni-kl.de(2012).[7]GianniP.,PropertiesofGr¨obnerBasesunderSpecialization,L.N.Comp.Sci.378(1987),293–297,Springer[8]D.Lazard,IdealBasisandPrimaryDecomposition:Caseoftwovariables,J.Symb.Comp.1(1985),261−270.[9]M.G.MarinariandTeoMora,Cerlienco-MuredduCorrespondenceandLazardStructuralTheorem.,RevistaInvesticaci`onOperacional,Vol.27,No.2,155-178,2006.[10]M.G.MarinariandTeoMora,AremarkonaremarkbyMacaulayorEnhancingLazardStructuralTheorem.,BulletinoftheIranianMathematicalSocietyVol.29No.1(2003),pagg.1−45.[11]M.G.MarinariandTeoMora,SomeCommentsonCerlienco-MuredduAlgorithmandEnhancedLazardStructuralTheorem,RejectedbyISSAC-2004(2004).[12]M.Kalkbrenner,SolvingSystemsofAlgebraicEquationsbyUsingGroebnerBases,L.N.Comp.Sci.378(1987),pagg.282−292,Springer.[13]KalkbrenerM.,OnthestabilityofGr¨obnerBasesunderspecialization,J.Symb.Comp.24(1997),51–58[14]M.G.Marinari,H.MMoeller,T.Mora,GroebnerBasesofIdealsDeﬁnedbyFunctionalswithanApplicationtoIdealsofProjectivePoints,ApplicableAl-gebrainEngineering,CommunicationandComputing,vol.4,1993,Springer.[15]M.G.Marinari,L.RamellaBorelIdealsinthreevariables,Beitr¨agezurAlgebraundGeometrie.ContributionstoAlgebraandGeometry,Vol47(2006),N.1,195−209.[16]T.Mora,Solvingpolynomialequationsystems:Macaulay’sparadigmandGroebnertechnology,CambridgeUniversityPress,2005.[17]G.H.Norton,A.S˘al˘agean,StrongGr¨obnerbasesforpolynomialsoveraprin-cipalidealring,Bull.Austral.Math.Soc.64(2001),505–528[18]S.Steidel,pointid.lib.ProceduresforcomputingafactorizedlexGBofthevanishingidealofasetofpointsviatheAxis-of-EvilTheorem(M.G.Marinari,T.Mora)(2011).[19]TrinksW.,¨UberB.BuchbergerVerfahren,SystemealgebraischerGleichungenzul¨osen,J.Numb.Th.10(1978),475–48820

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

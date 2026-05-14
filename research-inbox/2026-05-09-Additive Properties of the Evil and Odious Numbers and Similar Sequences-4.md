---
source: "https://arxiv.org/abs/2112.13627v3"
title: "Additive Properties of the Evil and Odious Numbers and Similar Sequences"
author: "Jean-Paul Allouche, Jeffrey Shallit"
year: "2021"
publication: "arXiv preprint / math.NT"
download: "https://arxiv.org/pdf/2112.13627v3"
pdf: "https://arxiv.org/pdf/2112.13627v3"
captured_at: "2026-05-09T12:44:35Z"
updated_at: "2026-05-09T12:44:35Z"
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

# Additive Properties of the Evil and Odious Numbers and Similar Sequences

- 著者: Jean-Paul Allouche, Jeffrey Shallit
- 年: 2021
- 掲載情報: arXiv preprint / math.NT
- 情報源: [arxiv](https://arxiv.org/abs/2112.13627v3)
- ダウンロード: https://arxiv.org/pdf/2112.13627v3
- PDF: https://arxiv.org/pdf/2112.13627v3

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

First we reprove two results in additive number theory due to Dombi and Chen & Wang, respectively, on the number of representations of n as the sum of two odious or evil numbers, using techniques from automata theory and logic. We also use this technique to prove a new result about the numbers represented by five summands. Furthermore, we prove some new results on the tenfold sums of the evil and odious numbers, as well as k-fold sums of similar sequences of integers, by using techniques of analytic number theory involving trigonometric sums associated with the (+-1)-characteristic sequences of these integers.

## PDF Text

AdditivePropertiesoftheEvilandOdiousNumbersandSimilarSequencesJean-PaulAlloucheCNRS,IMJ-PRG,Sorbonne4PlaceJussieuF-75252ParisCedex05Francejean-paul.allouche@imj-prg.frJe�reyShallit∗SchoolofComputerScienceUniversityofWaterlooWaterloo,OntarioN2L3G1Canadashallit@uwaterloo.caDecember21,2022AbstractFirstwereprovetworesultsinadditivenumbertheoryduetoDombiandChen&Wang,respectively,onthenumberofrepresentationsofnasthesumoftwoodiousorevilnumbers,usingtechniquesfromautomatatheoryandlogic.Wealsousethistechniquetoproveanewresultaboutthenumbersrepresentedby�vesummands.Furthermore,weprovesomenewresultsonthetenfoldsumsoftheevilandodiousnumbers,aswellask-foldsumsofsimilarsequencesofintegers,byusingtechniquesofanalyticnumbertheoryinvolvingtrigonometricsumsassociatedwiththe�1char-acteristicsequencesoftheseintegers.1IntroductionLetN=f0;1;2;:::gandletA�N.Ina1984paper,Erd}os,S�ark�ozy,andS�os[7]introducedthreefunctionsbasedonA,asfollows:1R(A)1(n)=jf(x;y)2N�N:x;y2Aandx+y=ngjR(A)2(n)=jf(x;y)2N�N:x;y2Aandx+y=nandx<ygj(1)R(A)3(n)=jf(x;y)2N�N:x;y2Aandx+y=nandx�ygj:
∗ResearchfundedbyagrantfromNSERC,2018-04118.1Infact,Erd}os,S�ark�ozy,andS�osusedadi�erentde�nitionofNthatexcludes0.Itseemsmorenaturaltoinclude0,andso(exceptinthelastsection)weadoptthisconvention.Onecaneasilygetexamplesoverthepositiveintegersbyshiftingthesetsby1,whichresultsinan\o�-by-k"errorwhentakingsumsofkterms.1
arXiv:2112.13627v3 [math.NT] 20 Dec 2022
Alsosee[14].Fori2f1;2;3g,apparentlyS�ark�ozyaskedwhetherthereexisttwosetsofpositiveintegersAandB,within�nitesymmetricdi�erence,forwhichR(A)i(n)=R(B)i(n)forallsu�cientlylargen.AsimpleexampleofsuchsetswasgivenbyDombi[6]in2002,andwedescribeitnext.Actually,thesameresulthadalreadyappearedearlierinapaperofLambekandMoser[9].Lett=t0t1t2���betheThue-Morsesequence,de�nedbyt0=0,t2n=tn,andt2n+1=1�tnforn�0.Itiseasilyseenthattnistheparityofthenumberof1's(orsumofbits)inthebinaryrepresentationofn.LetAandBbede�nedasfollows:A=fn�0:tn=0g=f0;3;5;6;9;10;12;:::gB=fn�0:tn=1g=f1;2;4;7;8;11;13;:::g:TheseformadisjointpartitionofN.Intheliterature,thesetAissometimescalledthesetofevilnumbers,andthesetBissometimescalledthesetofodiousnumbers.Theyare,respectively,sequencesA001969
andA000069
intheOn-LineEncyclopediaofIntegerSequences(OEIS)[15].DombiprovedthatR(A)2(n)=R(B)2(n)forn�0.Hisproofrequired21
2pagesandanumberofcases.InSection2weshowhowtoprovethisusingmore-or-lessroutinecalculationsinvolving�niteautomataandlogic.Chen&Wang[5]provedasimilarresultforthefunctionR3.InsteadoftheThue-Morsesequence,theyusedarelatedsequencet0ncountingtheparityofthenumberof0'sinthebinaryrepresentationofn,sometimescalledthetwistedThue-Morsesequence.2Wehavet00=1,t01=0,andt02n=1�t0nandt02n+1=t0nforn�1.(Uptothe�rsttermitisdo intheOEIS.)Chen&WangprovedthatifwesetC=fn�0:t0n=0g=f1;3;4;7;9;10;12;15;:::g;D=fn�0:t0n=1g=f0;2;5;6;8;11;13;14;:::g:thenR(C)3(n)=R(D)3(n)forn�1.3Thesesequencesare,respectivelyA059010
andA059009
intheOEIS.Theirproofrequired3pagesandcaseanalysis.Inthispaper,inSection3,wereprovetheirresultsusingtechniquesfromautomatatheoryandlogic.ForotherproofsoftheresultsofDombiandChen&Wang,see[13,10,17].WecanalsoconsidergeneralizationsofR(A)1(n)tomorethantwosummands,asfollows:rj(n):=jf(x1;x2;:::;xj):n=X1�i�jxiandtxi=0for1�i�jgj(2)sj(n):=jf(x1;x2;:::;xj):n=X1�i�jxiandtxi=1for1�i�jgj;(3)
2However,insomeformulations,thetwistedThue-Morsesequencehast00=0.3Again,thereisan\o�-by-two"di�erenceinthewaywestatedtheresult,comparedtothewaytheydid.2
wheret=t0t1t2���istheThue-Morsesequence.InSection4,weprovearesultfromcomplexanalysisthatallowsustoshowthatbothr10(n)ands10(n)areeventuallystrictlyincreasingfunctionsofn.Bycontrast,wecanuseourlogicalapproachtoshowthatthisisnotthecaseforr5(n)ands5(n).Thestatusforsumsof6;7;8;and9termsiscurrentlyunknown.InSection5weprovesomerelatedresults.2Automataand�rst-orderlogicOur�rstprooftechniquedependsonthefactthatboth(tn)and(t0n)arek-automaticse-quences.Thismeansthat,foreachsequence,thereexistsadeterministic�niteautomatonwithoutput(DFAO)computingthesequence,inthefollowingsense:whenwefeedthebase-krepresentationofnintotheautomaton,itprocessesthedigitsandendsinastateqwithoutputthen'thtermofthesequence.Forthesesequenceswehavek=2.Foreveryk-automaticsequence(an),thereisalogicaldecisionproceduretodecidethetruthofassertionsaboutthesequencethatarephrasedinthe�rst-orderlogicalstructurehN;+;<;n!ani.Wecallsuchaformulaak-automaticformula.Theresultsaresumma-rizedinthefollowingtwotheorems.Theorem1.Let'beak-automaticformula.Thereisadecisionprocedurethat,if'hasnofreevariables,willeitherproveordisprove'.Furthermore,if'hasfreevariablesi1;:::;ik,thentheprocedureconstructsadeterministic�niteautomatonacceptingthebase-krepresentationofthosetuples(i1;:::;ik)forwhichtheformulaevaluatestotrue.Foraproof,see[3].Wenowde�nethenotionoflinearrepresentationofafunction.Wesayf:N!Qhasalinearrepresentationofrankrifthereexistanintegerk�2,arowvectoru2Qr,acolumnvectorw2Qr,andanr�r-matrix-valuedmorphismsuchthatf(n)=u(x)vforallbase-krepresentationsxofn(includingthosewithleadingzeros).Theorem2.Thereisanalgorithmthat,givenak-automaticformula',withfreevariablesi1;i2;:::;it;n,computesalinearrepresentationforf(n),thenumberoft-tuplesofnaturalnumbers(i1;i2;:::;it)forwhich'(i1;i2;:::;it;n)istrue.Foraproof,see[4].Finally,thereisthenotionofminimallinearrepresentation,whichisarepresentationofsmallestrank.Awell-knownalgorithmofSch�utzenberger,basedonlinearalgebra,takesalinearrepresentationandproducesaminimalonefromit[2,§2.3].Thesearethebasictoolsweusetoprovetheresults.Theorems1and2havebeenimplementedinfreesoftwarecalledWalnut,originallycreatedbyHamoonMousavi[11,16],andavailableathttps://cs.uwaterloo.ca/~shallit/walnut.html.Theorem3.Suppose(an)n�0isak-automaticbinarysequenceandletAbethecorrespond-ingsetfn:an=1g.ThenthereisanalgorithmproducingthelinearrepresentationforeachofthefunctionsR(A)i(n),i=1;2;3.3
Proof.Itsu�cestogive�rst-orderlogicalformulasspecifyingthat(x;y)isanorderedpairwithsumncorrespondingtothepairsinthede�nition(1).Theyareasfollows:R1:n=x+y^ax=1^ay=1R2:n=x+y^x<y^ax=1^ay=1R3:n=x+y^x�y^ax=1^ay=1Here,asusual,thesymbol^denoteslogicalAND.
WenowgiveourproofofDombi'sresult,whichisbasedonroutinecalculationsusingtheresultsabove.Theorem4.(Dombi)R(A)2(n)=R(B)2(n)forn�0.Proof.The�rststepistoexpressthethesetofpairsasa�rst-orderformula.Wecandothisasfollows:'A:n=x+y^x<y^t[x]=0^t[y]=0'B:n=x+y^x<y^t[x]=1^t[y]=1:InWalnutthisistranslatedasevalr2a"n=x+y&x<y&T[x]=@0&T[y]=@0":evalr2b"n=x+y&x<y&T[x]=@1&T[y]=@1":Theresultingautomata,computedbyWalnut,bothhave12states.Next,fromthesematriceswecanimmediatelycomputealinearrepresentationforthenumberofpairs(x;y)makingtheformulatrue.TodosoinWalnutweusethefollowingcommands:evalr2amn"n=x+y&x<y&T[x]=@0&T[y]=@0":evalr2bmn"n=x+y&x<y&T[x]=@1&T[y]=@1":Thesecommandscreaterank-12linearrepresentationsforR(A)2(n)andR(B)2(n),asfollows:R(A)2(n)=(u;;vA)andR(B)2(n)=(u;;vB);whereu=2666666410000000000037777775T(0)=2666666410000000000000011000000000100000000000000000111000001000000010000000000100000010000000000001000000010001000100010010000100000000001000100000110037777775(1)=2666666401100000000000000100000000010011000000000000000100000100001001000000000000100000101000100000011000000000010000000000100000000011000100010000000037777775vA=2666666400000001000037777775vB=2666666400000010000037777775:4
Next,weapplytheminimizationalgorithmtothesetwo(slightly)di�erentlinearrepresen-tations,anddiscoverthattheybothminimizetothesamelinearrepresentation(u0;ˆ;v0)ofrank5,givenasfollows:u0="10000#Tˆ(0)="1000000100000010�1011�2�1310#ˆ(1)="0100000010�10110�21101�1�1021#v0="00010#Sincethesetwolinearrepresentationsarethesame,theresultisproved.
Remark5.ThesequenceR(A)2(n)isgivenassequenceA133009
intheOEIS.OnedistinctadvantagetothisapproachisthatalinearrepresentationforR(A)2(n)canbeusedtoeasilyproveadditionalresultsaboutit.Forexample:Theorem6.Fort�1wehave(a)R(A)2(2t�1)=(0;iftodd;2t�2;ifteven;(b)R(A)2(2t+1)=((2t+8)=12;ifteven;(2t+4)=6;iftodd:Proof.(a)Notethatthebase-2representationof2t�1consistsofthestringtz
}|
{11���1.ThereforeR(A)2(2t�1)=u0ˆ(1)tv0:Bywell-knownresults,theentriesofˆ(1)tsatisfyalinearrecurrence.Thereforesodoesu0ˆ(1)tv0.Bythefundamentaltheoremoflinearrecurrences,u0ˆ(1)tv0canbeexpressedintermsoftherootsoftheminimalpolynomialofˆ(1).ThisminimalpolynomialisX(X�1)(X�2)(X+2),andthereforeR(A)2(2t�1)=A�2t+B�(�2)t+CforsomeconstantsA;B;C.WecannowsolvefortheseconstantswiththevaluesofR(A)2(2t�1)computedfromthelinearrepresentationto�ndthatA=0,B=1=8,C=1=8.WethereforegetR(A)2(2t�1)=2t�3+(�2)t�3,whichprovestheresult.(b)Weusethefactthat2t+1hasbase-2representation1t�1z
}|
{00���01.Soitsu�cestocarryoutthesamecalculationsaswedidinpart(a),exceptnowtheyarebasedontheminimalpolynomialofˆ(0).Itisthesameasforˆ(1),namelyX(X�1)(X�2)(X+2).Wethen�nd(usingthesametechniqueasbefore)thatR(A)2(2t+1)=2=3+2t=8�(�2)t=24.Theresultnowfollows.
5
3OurproofoftheChen-WangresultTheorem7.(Chen-Wang)WithCandDde�nedasabove,wehaveR(C)3(n)=R(D)3(n)forn�1.Proof.Itiseasytoseethatthesequencet0=t00t01t02���=101001101���canbegeneratedbythefollowingDFAO:
Figure1:DFAOcomputingt0nHerethelabelsofthestatesaregivenintheform\statename/outputofthestate".WestartbytranslatingtheDFAOinFigure1intoWalnut,andstoreitasTT.txtinWalnut'sWordAutomataLibrary.msd_2010->01->1100->21->1210->11->2WecanthenprovetheequivalentresultthatR(C)3(n+1)=R(D)3(n+1)forn�0.evalr3cmn"n+1=x+y&x<=y&TT[x]=@0&TT[y]=@0":evalr3dmn"n+1=x+y&x<=y&TT[x]=@1&TT[y]=@1":6
Thisgivesustwolinearrepresentations,bothofrank20.Whenweminimizethese,asbefore,wegettwoidenticalminimizedrepresentations(u;;v),asfollows:u=2666641000000000377775T(0)=26666410000000000010000000000010000000000010000000000010002�20�1�12010�2002100�110�1�21111�110001�20�1�12110�35�300�23�23377775(1)=266664010000000000010000000000010000000000010000000000010�22001�21010�13�2�10�2301002�2�1�1�120200�1�10100110�33�100�3401377775v=2666640101111202377775:andsotheresultisproved.
Remark8.ThesequenceR(C)3(n)issequenceA059451
intheOEIS.4Resultsfor�veandtensummandsInthissectionweshowthatthesequencesr10ands10,de�nedaboveinEqs.(2)and(3),areeventuallystrictlyincreasing.Bycontrast,aswewillseelater,thesequencesr5ands5arenot.Forr10ands10,the\logicalapproach"ofprevioussectionsdoesnotseemtosu�cetoprovethestrictlyincreasingproperty,soweturninsteadtotechniquesofanalyticnumbertheory.Letq=(qn)n�0beasequenceof�1'stakingthevalue+1in�nitelyoften.Forcomplexnumberszandintegersn�0,wede�nethesumsQn(z):=X0�j�nqjzjQ(z):=Xj�0qjzj(forjzj<1):Wealsode�neL=LqbyL=Lq:=fn�1:qn�1=1gandg(L;z):=Pa2Lza.Letr(k;L;n)denotethenumberofsolutionsoftheequationn=x1+���+xkwithxj2Lforallj.Remark9.Notethat,withthenotationabove,wehavethat0=2L.Toseethatthisdoesnotrestrictthegenerality,notethat,ifwewanttorepresenttheintegerswithksummands,then,adding1toeveryelementoftheunderlyingsetjustshiftstherepresentationfunctionbytheadditiveconstantk.Theorem10.SupposethereexistsaconstantC>0andarealexponent�2(0;1)suchthat,forallz2Cwithjzj=1andforalln�1,onehasjQn(z)j�Cn�.Thenthesequencer(k;L;n)iseventuallystrictlyincreasingforeveryintegerksuchthatk>2=(1��).Proof.First,wenotethatthemaximummodulusprincipleimpliesthatjQn(z)j�Cn�forallzwithjzj�1andalln�1.Weclearlyhaveg(L;z)k=Pn2Nr(k;L;n)zn.Since�k;n:=r(k;L;n)�r(k;L;n�1)7
isthecoe�cientofzninthepowerseriesexpansionof(1�z)g(L;z)k,itsu�cestoprovethatthiscoe�cientispositiveforn>n0(k).Itthussu�cestoprovethat�k;n>0whennislargeenough.But,byCauchy'sdi�erentiationformula,�k;nisalsoequalto�k;n=1
2iˇI�(1�z)g(L;z)kdz zn+1where�isa(small)circlecenteredattheorigin.Thus,takingforthiscircleofintegration�=�k;n:=fz:z=re2iˇt;r=e�1=(n�k)g,wehave�k;n=Z10(1�z)g(L;z)kz�ndt;withz=re2iˇtandr=e�1=(n�k):(4)Sinceg(L;z)=Xa2Lza=Xj�11
2(qj�1+1)zj=z
2�1
1�z+Q(z)�;weobtain�k;n=Z10(1�z)�z
2�k�1
1�z+Q(z)�kz�ndt=Z10(1�z)�z
2�k�1
1�z+Qn(z)�kz�ndt:(5)NotethatthetermsinQcorrespondingtoindices>ngiveintegralsequalto0.Hence�k;n=2�kZ10z�(n�k)(1�z) X0�`�k�k`�1
(1�z)k�`Q`n(z)!dt:Nowwesplit�k;nintothreequantities:thetermcorrespondingto`=0,theterm`=k,andthetermcorrespondingto`2[1;k�1].For`=0thecorrespondingtermis2�kZ101
(1�z)k�1z�(n�k)dt=2�kZ10 Xr�0�k+r�2r�zr!z�(n�k)dt=2�k�n�2n�k�=2�k�n�2k�2�˘2�knk�2
(k�2)!�For`=k,weusetheupperboundjQn(z)j�Cn�,thusobtainingthebound����2�kZ10(1�z)z�(n�k)Qkn(z)dt�����21�kCkenk�:NowwelookatthetermsI`:=2�kZ101
(1�z)k�`�1�k`�z�(n�k)Q`ndt8
for`2[1;k�1].UsingtheboundjQn(z)j�Cn�andthefactthatjzj=e�1=(n�k),weobtainjI`j�2�k�k`�C`n�`eZ10����1
(1�z)k�`�1����dt:(6)Now,inordertoevaluatetheintegralin(6),we�rstnotethat(recallthatz=re2iˇt)Z10����1
(1�z)k�`�1����dt=2Z1=20����1
(1�z)k�`�1����dt:Then,mimickingDombi'smethodin[6],wesplittheinterval[0;1=2]into[0;1=2]=J1[J2wherewede�neJ1:=[0;n�(�+")][[1=2�n�(�+");1=2]:andJ2:=[n�(�+");1=2�n�(�+")];sothatZ1=20����1
(1�z)k�`�1����dt=ZJ1����1
(1�z)k�`�1����dt+ZJ2����1
(1�z)k�`�1����dt:ForJ1,sincejzj=r<1,wehavewhenngoestoin�nity(recallthatkis�xed),that����1
1�z�����1
1�jzj=1
1�e�1=(n�k)˘n�k˘n:Thus����1
1�z����k�`�1=O(nk�`�1)andZJ1����1
(1�z)k�`�1����dt=O(nk�`�1���"):ForJ2,wenotethat,forx2[�;ˇ��](with�2(0;ˇ=2)),wehavesinx�sin��(2=ˇ)�.Hence,fort2J2andnlargeenough����1
1�z���������1
=(1�z)����=����1
rsin(2ˇt)����=O(e1=(n�k)n�+")=O(n�+"):Thus����1
1�z����k�`�1=O(n(�+")(k�`�1))andZJ2����1
(1�z)k�`�1����dt=O(n(�+")(k�`�1)):Finally,weobtainjI`j=O(n�`+k�`�1���")+O(n�`+(�+")(k�`�1)):If�<k�2
k�1,i.e.,k>2��
1��,wecanchoose":=k�2
k�1��>0.ItiseasytocheckthatthisimpliesjI`j=O(nk�2�")for`2[1;k�1]:9
namely�(`�1)�(`�1),hence�`�`����1,whichgives�`+k�`�1���"�k�2�",and�`+(�+")(k�`�1)=((k�2)=(k�1)�")`+(k�`�1)(k�2)=(k�1)=k�2�"`�k�2�".GatheringtheboundsforjIkjandjI`jfor`2[1;k�1]wehaveX1�`�kjI`j=O(nk�)+O(nk�2�")providedthatk>2��
1��:Hence�k;n˘I0˘2�knk�2
(k�2)!providedthatk>2��
1��andk�<k�2.Sincetheconditionk�<k�2,i.e.,theinequalityk(1��)>2,impliesthatk>2��
1��,wearedone.
Corollary11.Thesequencesr10ands10areeventuallystrictlyincreasing.Proof.WeapplyTheorem10tor10ands10.Inthiscasewetakeqn=(�1)tn,andusetheknownfact[8,12]thatforthissequencewehavesupjzj=1jQn(z)j�Cn�for�=(log3)=(log4):=0:79248.Since10>2=(1��):=9:63768,wegetthatr10ands10are(eventually)strictlyincreasingfunctionsofn.
Thestatusfor6;7;8;and9summandsiscurrentlyunknown.Basedonnumericalevi-dence,wemakethefollowingconjectures:Conjecture12.(a)Bothr6(n)ands6(n)areeventuallystrictlyincreasing.(a)r6(n)<r6(n+1)forn�37.(b)s6(n)<s6(n+1)forn�5.Nowweturnourattentiontor5ands5.Incontrasttothesituationforr10ands10,wecanuseour\logicalapproach"toshowthatthesesequencesarenotstrictlyincreasing.Forany�xedj,onecaneasilyobtainlinearrepresentationsforrjandsjusingthemethodsexplainedabove.Theorem13.Wehaver5(2n)>r5(2n+1)ands5(2n)>s5(2n+1)forallsu�cientlylargen.Proof.WecanuseWalnuttocomputealinearrepresentationforr5(n),asfollows:evalr5n"n=i+j+k+l+m&T[i]=@0&T[j]=@0&T[k]=@0&T[l]=@0&T[m]=@0":Thisgivesusvectorsv;wandamatrix-valuedmorphismsuchthatv(x)w=r5(n)forallbinarystringsxsuchthat[x]2=n.Therankofthislinearrepresentationis160,andisnotgivenhereforspacereasons.Next,wecomputetheminimalpolynomialof(0)usingMaple.ItisX4(X�1)(X�2)(X�4)(X�8)(X�16)(X+2)(X+4)(X+8)(X2�8)(X2�2X�16):10
Itfollowsthatbothr5(2n)andr5(2n+1)canbewrittenasalinearcombinationofthen'thpowersofthezerosofthispolynomial,andtherefore,soisthedi�erencer5(2n)�r5(2n+1).Whenwesolveforthecoe�cientsofthislinearcombination,we�ndthatthecoe�cientcorrespondingto16nispositive(infactitis1=14039101440).Since16isthedominantroot,thisshowstheexistenceofsomen0suchthatthedi�erencer5(2n)�r5(2n+1)ispositiveforalln�n0.Exactlythesameproof,word-for-word,worksfors5.
Remark14.Theorem10canbeappliedtoseveralothersequencesforwhichtheconditionjQn(z)j�Cn�forsome�in(0;1)holds.Wegivebutonefamilyofexamples|namely,theGolay-Shapiro-Rudinsequences,forwhichitisknownthat�=1=2,andhencek>4.FortheusualGolay-Shapiro-Rudinsequence,thisisexactlythe�rstpartofDombi'sTheorem1[6,p.138];moregenerallythisalsogivesk>4forthegeneralizedRudin-ShapirosequencesofTheorem3.1in[1,p.20],with'andvbeingtheconstantsequence1.5OtherresultsForthefollowingresultandproof,weadopttheIversonnotationwhere,forapropositionP,weset[P]=1ifPistrueand[P]=0otherwise.Theorem15.Forn�0wehaver2(n)�s2(n)=[neven](�1)tn.Proof.Wecan�ndlinearrepresentationsforr2ands2withtheWalnutcommandsevalr2mn"n=x+y&T[x]=@0&T[y]=@0":evals2mn"n=x+y&T[x]=@1&T[y]=@1":Theyare(u2;2;v2)forr2and(u2;2;v02)fors2,whereu2=2664100000003775T2(0)=2664100000000000111000100000000100000101000101100001000000101000110037752(1)=266401110000000000011000101010000110000001000000100000110001010000003775v2=2664100000003775v02=2664000000103775:Fromthiswecaneasilyformalinearrepresentationforr2(n)�s2(n)asfollows:(u2;2;v2�v02).Whenweminimizeit,wegetalinearrepresentation(x2;02;y2)ofrank2,asfollows:x2=�10�T02(0)=�10�10�02(1)=�010�1�y2=�10�:Nowaneasyinductiongivesthat02(x)=�[neven](�1)tn�[nodd](�1)tn�[neven](�1)tn[nodd](�1)tn�forn�1andallstringsxsuchthat[x]2=n.Thiscompletestheproof.
11
Theorem16.Therearein�nitelymanynforwhichr3(n)=s3(n).Someexamplesincluden=4i�2fori�1andn=3�4i�1fori�0.Proof.Wecan�ndlinearrepresentationsforr3(n)ands3(n)usingthefollowingWalnutcommands:evalr3mn"n=x+y+z&T[x]=@0&T[y]=@0&T[z]=@0":evals3mn"n=x+y+z&T[x]=@1&T[y]=@1&T[z]=@1":Itturnsouttheselinearrepresentationsareofrank24andoftheform(u3;3;v3)and(u3;3;v03),respectively.Sowecanformthelinearrepresentationforr3(n)�s3(n)by(u3;3;v3�v03).Whenweminimizeit,wegetalinearrepresentation(x3;03;y3)ofrank6,asfollows:x3=2410000035T03(0)=24100000001000000010�1�1�10�1�1�3100�2125�32433503(1)=240100000001000000013�32�31�1�3�1�22�1�2�120�1�1235y3=241002�3035:Nowthebinaryrepresentationof4i�2isoftheform12i�10,soweknowthatr(4i�2)�s(4i�2)canbeexpressedasalinearcombinationofthe(2i�1)thpowersoftherootsoftheminimalpolynomialof03(1).ThisminimalpolynomialisX2(X+1)(X2�8).Solvingforthislinearcombination,we�ndthatthecoe�cientsareallzero,sor(4i�2)�s(4i�2)=0foralli�1.Actually,withthesametechnique,onecanprovethatr(4i�2)=s(4i�2)=16i�1�4i�1.For3�4i�1,thesameideaswork.
Theorem17.Therearein�nitelymanynforwhichr4(n)=s4(n).Someexamplesincluden=6�4i�1fori�0andn=2�4i�3fori�1.Proof.Wecan�ndlinearrepresentationsforr4(n)ands4(n)usingthefollowingWalnutcommands:evalr4mn"n=x+y+z+w&T[x]=@0&T[y]=@0&T[z]=@0&T[w]=@0":evals4mn"n=x+y+z+w&T[x]=@1&T[y]=@1&T[z]=@1&T[w]=@1":Itturnsouttheselinearrepresentationsareofrank64andoftheform(u4;4;v4)and(u4;4;v04),respectively.Sowecanformthelinearrepresentationforr4(n)�s4(n)by(u4;4;v4�v04).Whenweminimizeit,wegetalinearrepresentation(x4;04;y4)ofrank7,asfollows:x4=2641000000375T04(0)=26641000000001000000001000000001�3100�210�1
25
2�11
223
21
2�5
25
2�7
29
2�2�1
2�3
25
2377504(1)=266640100000000100000000107
2�5
25
2�33
2�1
21
2�5
2�1
2�3
22�1
2�3
21
2�120�1�1207
2�5
23
2�23
2�3
2�1
237775y4=2641004�104375:Nowthebinaryrepresentationof6�4i�1isoftheform1012i+1,soweknowthatr4(6�4i�1)�s4(6�4i�1)canbeexpressedasalinearcombinationofthe(2i+1)thpowersoftherootsoftheminimalpolynomialof04(1).ThisminimalpolynomialisX3(X+1)(X2�8).Solvingforthis12
linearcombination,we�ndthatthecoe�cientsareallzero,sor4(6�4i�1)�s4(6�4i�1)=0foralli�0.Infact,withalittlemorework,andthesametechnique,onecanshowthatr4(6�4i�1)=s4(6�4i�1)=9
464i+16i+4i
8+c1�i1+c2�i2;where�1=18�2p
17,�2=18+2p
17,c1=(7�1�2�2)=288,c2=(7�2�2�1)=288.For2�4i�3,thesametechniqueworks.
AcknowledgmentsWethankEmmanuelLesigneforraisingtheproblemandMichelDekkingfordiscussions.References[1]J.-PAlloucheandP.Liardet.GeneralizedRudin-Shapirosequences.ActaArith.60(1991),1{27.[2]J.BerstelandC.Reutenauer.NoncommutativeRationalSeriesWithApplications,Vol.137ofEncyclopediaofMathematicsandItsApplications.CambridgeUniversityPress,2011.[3]V.Bruy�ere,G.Hansel,C.Michaux,andR.Villemaire.Logicandp-recognizablesetsofintegers.Bull.BelgianMath.Soc.1(1994),191{238.Corrigendum,Bull.BelgianMath.Soc.1(1994),577.[4]E.Charlier,N.Rampersad,andJ.Shallit.Enumerationanddecidablepropertiesofautomaticsequences.Internat.J.Found.Comp.Sci.23(2012),1035{1066.[5]Y.-G.ChenandB.Wang.Onadditivepropertiesoftwospecialsequences.ActaArith.110(2003),299{303.[6]G.Dombi.Additivepropertiesofcertainsets.ActaArith.103(2002),137{146.[7]P.Erd}os,A.S�ark�ozy,andV.T.S�os.ProblemsandresultsonadditivepropertiesofgeneralsequencesIV.InK.Alladi,editor,NumberTheory:Proceedingsofthe4thMatscienceConferenceheldatOotacamund,India,January5{10,1984,Vol.1122ofLectureNotesinComputerScience,pp.85{104.Springer-Verlag,1984.[8]A.O.Gel'fond.Surlesnombresquiontdespropri�et�esadditivesetmultiplicativesdonn�ees.ActaArith.13(1967/1968),259{265.[9]J.LambekandL.Moser.Onsometwowayclassi�cationsofintegers.Canad.Math.Bull.2(1959),85{89.13
[10]V.F.Lev.Reconstructingintegersetsfromtheirrepresentationfunctions.Electron.J.Combin.11(1)(2004),#R78.[11]H.Mousavi.AutomatictheoremprovinginWalnut.Arxivpreprint,arXiv:1603.06017[cs.FL],availableathttp://arxiv.org/abs/1603.06017,2016.[12]D.J.NewmanandM.Slater.Binarydigitdistributionovernaturallyde�nedsequences.Trans.Amer.Math.Soc.213(1975),71{78.[13]C.S�andor.Partitionsofnaturalnumbersandtheirrepresentationfunctions.INTEGERS|Elect.J.Comb.Numb.Theory4(2004),#A18.[14]A.S�ark�ozy.Onthenumberofadditiverepresentationsofintegers.InE.Gy}ori,G.O.H.Katona,andL.Lov�asz,editors,MoreSets,GraphsandNumbers,BolyaiSocietyMathematicalStudies,Vol.15,Springer,2006,pp.329-339.[15]N.J.A.Sloaneetal.Theon-lineencyclopediaofintegersequences.Electronicresource,availableathttps://oeis.org,2021.[16]J.Shallit.TheLogicalApproachToAutomaticSequences:ExploringCombinatoricsonWordswithWalnut,Vol.482ofLondonMath.Soc.LectureNoteSeries.CambridgeUniversityPress,2022.[17]M.Tang.Partitionsofthesetofnaturalnumbersandtheirrepresentationfunctions.DiscreteMath.308(2008),2614{2616.14

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

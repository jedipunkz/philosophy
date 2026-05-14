---
source: "https://arxiv.org/abs/1405.6214v1"
title: "Beyond odious and evil"
author: "Jean-Paul Allouche, Benoit Cloitre, Vladimir Shevelev"
year: "2014"
publication: "arXiv preprint / math.NT"
download: "https://arxiv.org/pdf/1405.6214v1"
pdf: "https://arxiv.org/pdf/1405.6214v1"
captured_at: "2026-05-09T12:44:00Z"
updated_at: "2026-05-09T12:44:00Z"
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

# Beyond odious and evil

- 著者: Jean-Paul Allouche, Benoit Cloitre, Vladimir Shevelev
- 年: 2014
- 掲載情報: arXiv preprint / math.NT
- 情報源: [arxiv](https://arxiv.org/abs/1405.6214v1)
- ダウンロード: https://arxiv.org/pdf/1405.6214v1
- PDF: https://arxiv.org/pdf/1405.6214v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

In a recent post on the Seqfan list the third author proposed a conjecture concerning the summatory function of odious numbers (i.e., of numbers whose sum of binary digits is odd), and its analog for evil numbers (i.e., of numbers whose sum of binary digits is even). We prove these conjectures here. We will also study the sequences of "generalized" odious and evil numbers, and their iterations, giving in particular a characterization of the sequences of usual odious and evil numbers in terms of functional equations satisfied by their compositions.

## PDF Text

arXiv:1405.6214v1 [math.NT] 23 May 2014
BeyondodiousandevilJ.-P.Allouche∗CNRS,InstitutdeMath.deJussieuUniversit´eP.etM.Curie,Case2474PlaceJussieuF-75252ParisCedex05,Franceallouche@math.jussieu.frBenoitCloitre19,rueLouiseMichel92300Levallois-PerretFrancebenoit7848c@yahoo.frV.ShevelevDepartmentofMathematicsBen-GurionUniversityoftheNegevBeersheva,Israelshevelev@bgu.ac.il[...]Thelightshanging fromoakbeamsabovethereaders lightandilluminateeverypage.Eachbookdustedeachday.Originaljackets,noodiousnumbersgluedtospines, notonedecimal,Deweyorotherwise,intheentireplace![...]
(ThomasLux,TheAmbrosianaLibrary)AbstractInarecentpostontheSeqfanlistthethirdauthorproposedaconjectureconcerningthesummatoryfunctionofodiousnumbers(i.e.,ofnumberswhosesumofbinarydigitsisodd),anditsanalogforevilnumbers(i.e.,ofnumberswhosesumofbinarydigitsiseven).Weprovetheseconjectureshere.Wewillalsostudythesequencesof“generalized”odiousandevilnumbers,andtheiriterations,givinginparticularacharacterizationofthesequencesofusualodiousandevilnumbersintermsoffunctionalequationssatisﬁedbytheircompositions.
Keywords:Odiousnumbers;oddnumbers;Thue-Morsesequence;summatoryfunc-tions;iterationofsequences.
MSCClasses:11A63,11B83,11B85,11A07,05A15.
∗TheauthorwaspartiallysupportedbytheANRproject“FAN”(FractalsetNum´eration).1
1Introduction
Thepurposeofthispaper,whosetitleisreminiscentof[9],isrevisitingthestudyoftwofamiliesofintegersrespectivelycalledodiousandevilnumbers,aswellasthestudyofsomegeneralizations.Anaturalintegeriscalledodiousifthesumofitsbinarydigitsisodd.Anaturalintegeriscalledevilifthesumofitsbinarydigitsiseven.Thisterminologywasintroducedbytheauthorsof[3,4],see[4,p.463];thewords“odious”and“evil”werechosenbecausetheybeginrespectivelylike“odd”and“even”.Leta=(a(n))n≥0denotetheincreasingsequenceofodiousnumbers,andb=(b(n))n≥0denotetheincreasingsequenceofevilnumbers.SequencesaandbarerespectivelyA000069andA001969in[10],exceptthatwelettheindexesstartfrom0insteadfrom1.Sequencesaandbbegina=124781113141619212225262831...b=035691012151718202324272930...Remark1Weseemtorememberhavingreadsomewhereontheweb(butwherewasit?)thatmathematiciansprobablydonotlikenumbers,sinceforthemnumbersarenecessarilyeitherevilortheyareodd!Alsonotethatforsomeauthorstheexpression“evilnumbers”hasaquitediﬀerentmeaning.Thisisexplained,e.g.,in[13]:anumberiscalledevilinthatterminologyiftheﬁrstndecimaldigitsofitsfractionalpartsumto666forsomeintegern.Theexpression“evilnumbers”isalsousedwithothermeanings,inspiringinparticularartists,likeFabioMauri(see[6]).Thepurposeofthispaperistwofold.Firsttoprovetheaboveconjecture.Secondtodescribeiterationsofsequencesaandbandgeneralizations.Namelyitwasnotedbythesecondauthorin[10,A000069]thata(a(n))=2a(n).WewillgeneralizethisresultbyprovingTheorem1belowwhichgivesanexpressionofaj,d(ai,d(n)),where(aj,d(n))n≥0denotetheincreasingsequenceofintegerswhosesumofd-arydigitsiscongruenttojmodulod.2IterationofthesequencesofgeneralizedodiousandevilnumbersAsrecalledintheintroduction,itwasnotedbythesecondauthor[10,A000069]thattheincreasingsequenceofodiousnumbersasatisﬁesa(a(n))=2a(n).WewillgeneralizethisresultbyprovingTheorem1below.
Webeginwithadeﬁnitionandalemma.
Deﬁnition1•Letd≥2beaninteger.Foranyintegerxwelet
(x)ddenotetheresiduemodulodofx,i.e.,theintegerbelongingto[0,d−1]andcongruenttoxmodulod.Notethatwehavex=dbx dc+
(x)d.2
•Weletsd(n)denotethesumofthed-arydigitsofn.Welettd=(td(n))n≥0denotethesequenceofintegersdeﬁnedbytd(n)≡sd(n)moddand0≤td(n)≤d−1,i.e.,td(n)=
(sd(n))d.•Forj∈[0,d−1],weletaj,d=(aj,d(n))n≥0denotetheincreasingsequenceofintegersksuchthatsd(k)≡jmodd(i.e.,td(k)=j).Lemma1•Sequencetdistheﬁxedpointofthemorphismdeﬁnedon{0,1,...,d−1}by0→01...d−1,1→12...d−10,...,d−1→d−101...d−2.•Ifαbelongsto[0,d−1],then,foralln≥0,wehavetd(dn+α)=
(td(n)+α)d=(td(n)+αiftd(n)+α≤d−1td(n)+α−diftd(n)+α≥d.•Ifjbelongsto[0,d−1],then,foralln≥0,wehaved−1−td(dn+d−1−j)=(j−td(n)if0≤td(n)≤jd+j−td(n)ifj+1≤td(n)≤d−1.Proof.Theproofoftheﬁrsttwoitemsiseasyandlefttothereader.Thelastitemisaneasyconsequenceoftheseconditem.Nowweproveahelpfulproposition.Proposition1Thesequence(aj,d(n))n≥0satisﬁesaj,d(n)=dn+
(j−td(n))d.Thiscanalsobewrittenaj,d(n)=dn+(j−td(n)if0≤td(n)≤jd+j−td(n)ifj+1≤td(n)≤d−1=dn+d−1−td(dn+d−i−1).Proof.TheseequalitiesareeasyconsequencesofLemma1,whichimpliesinparticularthatsequencetdconsistsofconsecutiveblockstakenfrom(01...d−1),(12...d−10),...,(d−101...d−2),wherether-thblockbeginswithtd(r).Wearereadytostateandprovetheresultofthissection.3
Theorem1Foralln≥0,foralli,j∈[0,d−1],wehaveaj,d(ai,d(n))=dai,d(n)+
(j−i)d=(dai,d(n)+j−iifj≥idai,d(n)+d+j−iifj<iProof.UsingProposition1wecanwriteaj,d(ai,d(n))=dai,d(n)+
j−td(ai,d(n))d.But,fromthedeﬁnitionofai,dwehavetd(ai,d(n))=i.Henceaj,d(ai,d(n))=dai,d(n)+
(j−i)d=(dai,d(n)+j−iifj≥idai,d(n)+d+j−iifj<i.3SummatoryfunctionofgeneralizedodiousandevilnumbersInordertoaddressShevelev’sconjecturerecalledintheintroduction,wehavetostudythesummatoryfunctionofodiousnumbers.Thissectionisdevotedtostudyingthesummatoryfunctionofgeneralizedodiousandevilnumbers.Theﬁrststepisthefollowingproposition.(WekeepthenotationinDeﬁnition1.)
Proposition2Letaandrbeintegersin[0,d−1].ThenrX`=0
(a−`)d=



a(r+1)−r(r+1)
2ifr≤aa(r+1−d)+dr−r(r+1)
2ifr>a.ThiscanalsobewrittenrX`=0
(a−`)d=a(r+1)−r(r+1)
2+dmax{r−a,0}.Proof.Ifr≤a,thenrX`=0
(a−`)d=
(a)d+
(a−1)d+···+
(a−r)d=a+(a−1)+···+(a−r)=a(r+1)−r(r+1)
2·Now,ifr>a,thenrX`=0
(a−`)d=
(a)d+
(a−1)d+···+
(0)d+
(−1)d+
−(r−a)d=a+(a−1)+···+0+(d−1)+···+(d−r+a)=a(r+1−d)+dr−r(r+1)
2·4
UsingProposition1andProposition2weobtainthefollowingtheorem.
Theorem2Thesummatoryfunctionofthesequenceaj,disgivenbyNXk=0aj,d(k)=dN(N+1)
2+bN/dcd(d−1)
2+
(j−td(bN/dc))d(
(N)d+1)−
(N)d(
(N)d+1)
2+dmax{
(N)d−
(j−td(bN/dc))d,0}Proof.Weﬁrstnotethat,foranyintegera,wehavead−1Xk=0
(j−td(k))d=a−1Xj=0(j+1)d−1Xk=jd
(j−td(k))d=a−1Xj=0d(d−1)
2=ad(d−1)
2sinceforeachj,whenkrunsin[jd,(j+1)d−1],td(k)takesexactlyonceeveryvaluein[0,d−1],thus
(j−td(k))dalsotakesexactlyonceeveryvaluein[0,d−1].Now,usingProposition1,wegetNXk=0aj,d(k)=NXk=0(dk+
(j−td(k))d)=dN(N+1)
2+NXk=0
(j−td(k))d=dN(N+1)
2+dbN/dc−1Xk=0
(j−td(k))d+NXk=dbN/dc
(j−td(k))d=dN(N+1)
2+bN/dcd(d−1)
2+NXk=dbN/dc
(j−td(k))d.ButNXk=dbN/dc
(j−td(k))d=N−dbN/dcX`=0
(j−td(`+dbN/dc))d=N−dbN/dcX`=0
(j−
(td(`)+td(bN/dc)d)d=N−dbN/dcX`=0
(j−td(bN/dc)−`)d=N−dbN/dcX`=0
(
(j−td(bN/dc))d−`)d.Now,usingProposition2withareplacedby
(j−td(bN/dc))dandrreplacedbyN−dbN/dcyieldstheresult.NowwegiveacorollaryofTheorem2above,recallinginpassingsomenotationoftheintroduction.5
Corollary1Lett=(t(n))n≥0betheThue-Morsesequencedeﬁnedbyt(0)=0,andforalln≥0,t(2n)=t(n),t(2n+1)=1−t(n).Leta=(a(n))n≥0bethesequenceofodiousnumbers.Letb=(b(n))n≥0bethesequenceofevilnumbers.LetS(n)=a(0)+a(1)+···+a(n)bethesummatoryfunctionofsequencea.LetR(n)=b(0)+b(1)+···+b(n)bethesummatoryfunctionofsequenceb.Then•a(n)=2n+1−t(n)•b(n)=2n+t(n)•S(n)=n2+3n
2+1
2+1+(−1)n
4(1−2t(n))=n2+3n
2+1
2ifnisoddn2+3n
2+1−t(n)ifniseven.•R(n)=n2+3n
2+1
2+1+(−1)n
4(2t(n)−1)=n2+3n
2+1
2ifnisoddn2+3n
2+t(n)ifniseven.Proof.Putd=2andj=0,1inTheorem2.(Alsosee,e.g.,[7,Section8],[10,A000069]and[10,A173209].)Notethatitisalsopossibletogivesummatoryfunctionsofpolynomialexpressionsforsequenceslikeaandb.Forexample,wecanprovethefollowingresult.Corollary2Wehavetherelationsb(n)Xk=0a(k)=b(n)2+b(n)+n+1a(n)Xk=0b(k)=a(n)2+a(n)+n+1a(n)Xk=0a(k)=a(n)2+2a(n)−nb(n)Xk=0b(k)=b(n)2+2b(n)−n2.WehavetherelationsXa(k)≤na(k)=







n2
4−n
4+nt(n)ifn≡0mod4n2
4+n
4−1
2+t(n)ifn≡1mod4n2
4−n
4−1
2+(n+1)t(n)ifn≡2mod4n2
4+n
4ifn≡3mod4Xb(k)≤nb(k)=







n2
4+3n
4−nt(n)ifn≡0mod4n2
4+n
4+1
2−t(n)ifn≡1mod4n2
4+3n
4+1
2−(n+1)t(n)ifn≡2mod4n2
4+n
4ifn≡3mod4.6
Proof.Lefttothereader.Hintforthelasttwoformulas:notethat{k;a(k)≤n}={k;2k+1−t(k)≤n}={k;t(k)=1,k≤bn
2c}∪{k;t(k)=0,k≤bn−1
2c}={0,1,2,...,bn−1
2c}∪ΘnwhereΘn=∅ifnisodd,orifnisevenandt(n)=0{bn
2c}ifnisevenandt(n)=1.ThusXa(k)≤na(k)=Xk≤bn−1
2ca(k)+1+(−1)n
2t(n)a(bn/2c).4ProofofaconjectureofShevelev
BeforestatingShevelev’sconjecture,weneedadeﬁnitionandalemma.Deﬁnition2Ifxandyaretwointegers,wewritex<4y(resp.x≤4y)iftheresiduesmodulo4ofxandy,denotedby xand y,belongingto{0,1,2,3}andconsideredasnaturalintegers,satisfy x<
y(resp.
x≤
y).Example1Forexample17<46because17≡1mod4,6≡2mod4and1<2.Lemma2LetaandtbeasabovetheincreasingsequenceofodiousnumbersandtheThue-Morsesequence.Then,ifnandmarebothoddorbotheven,thena(n)<4a(m)(resp.a(n)≤4a(m))ifandonlyift(m)<t(n)(resp.t(m)≤t(n)).Proof.Sincea(n)=2n+1−t(n),weseethat a(n)=1−t(n)ifniseven,andthat a(n)=3−t(n)ifnisodd.Thestatementinthelemmafollows.Theorem3(Shevelev’sconjecture)LetS(n)bethesummatoryfunctionofodiousnum-bers,i.e.,S(n)=a(0)+a(1)+···+a(n).Wehave,forn≥2,•ifa(n−1)<4a(n+1)anda(n)≤4a(n+2),thenS(n)=a(n)a(n+1)
4•ifa(n−1)>4a(n+1)anda(n)≥4a(n+2),thenS(n)=a(n)a(n+1)
4+1
2•ifa(n−1)<4a(n+1)anda(n)>4a(n+2),thenS(n)=a(n)(a(n+1)−1)
4•ifa(n−1)>4a(n+1)anda(n)<4a(n+2),thenS(n)=(a(n)+1)a(n+1)
4·7
Proof.(i)Proofoftheﬁrstassertion.FromLemma2theconditionsa(n−1)<4a(n+1)anda(n)≤4a(n+2)areequivalenttot(n+1)<t(n−1)andt(n+2)≤t(n).Thust(n+1)=0,t(n−1)=1,andeithert(n)=1ort(n)=t(n+2)=0.Butt(n)=t(n+2)=0isimpossiblebecausethiswouldimply((t(n),t(n+1),t(n+2))=(0,0,0)andtheThue-Morsesequencedoesnotcontainanycube.Thust(n)=1,t(n−1)=1,t(n+1)=0.Furthermoret(n)=t(n−1)(=1)impliesthatnmustbeeven(ifn=2k+1,t(n)=1−t(k)andt(n−1)=t(k)).So,usingCorollary1,S(n)=n2+3n
2+1−t(n)=n2+3n
2.Ontheotherhanda(n)a(n+1)=(2n+1−t(n))(2n+3−t(n+1)=2n(2n+3)=4n2+6n=4(n2+3n
2).(ii)Proofofthesecondassertion.FromLemma2theconditionsa(n−1)>4a(n+1)anda(n)≥4a(n+2)areequivalenttot(n+1)>t(n−1)andt(n+2)≥t(n).Hencet(n−1)=0,t(n+1)=1,andeithert(n)=t(n+2)=1ort(n)=0.Butwecannothavet(n)=t(n+2)=1,becausethiswouldgive(t(n),t(n+1),t(n+2))=(1,1,1)andthiswouldgiveacubeintheThue-Morsesequence.Thust(n−1)=0,t(n)=0,t(n+1)=1.Aspreviouslyt(n−1)=t(n)(=0)impliesthatnmustbeeven.So,usingCorollary1,S(n)=n2+3n
2+1−t(n)=n2+3n
2+1.Ontheotherhanda(n)a(n+1)+2=(2n+1−t(n))(2n+3−t(n+1)+2=(2n+1)(2n+2)+2=4n2+6n+4=4(n2+3n
2+1).(iii)Proofofthethirdassertion.FromLemma2theconditionsa(n−1)<4a(n+1)anda(n)>4a(n+2)areequivalenttot(n+1)<t(n−1)andt(n+2)>t(n).Thust(n−1)=1,t(n)=0,t(n+1)=0,t(n+2)=1.Sincet(n)=t(n+1)(=0),nmustbeodd.So,usingCorollary1,S(n)=n2+3n
2+1
2.Ontheotherhanda(n)(a(n+1)−1)=(2n+1−t(n))(2n+2−t(n+1))=(2n+1)(2n+2)=4(n2+3n
2+1
2).(iv)Proofofthefourthassertion.FromLemma2theconditionsa(n−1)>4a(n+1)anda(n)<4a(n+2)areequivalenttot(n+1)>t(n−1)andt(n+2)<t(n).Hencet(n−1)=0,t(n)=1,t(n+1)=1,t(n+2)=0.Sincet(n)=t(n+1)(=1),nmustbeodd.So,usingCorollary1,S(n)=n2+3n
2+1
2.Ontheotherhand(a(n)+1)a(n+1)=(2n+2−t(n))(2n+3−t(n+1))=(2n+1)(2n+2)=4(n2+3n
2+1
2).5Functionalequationsforsequencesaandb
Severalstudiesaboutiteratingincreasingsequencesofintegerscanbefoundintheliterature(see,e.g.,[1,5,8,11]andreferencestherein,inparticularpartsofReference[12]thatwediscoveredthanksto[11]).Withthepreviousnotation,theincreasingsequencesofodiousandofevilnumberssatisfya(n)=a1,2(n)andb(n)=a0,2(n).Wethushavethefollowingrelations.Corollary38
•(i)a(a(n))=2a(n)•(ii)b(b(n))=2b(n)•(iii)a(b(n))=2b(n)+1•(iv)b(a(n))=2a(n)+1•(v)a(a(n))=b(a(n))−1•(vi)b(b(n))=a(b(n))−1•(vii)a(n)−b(n)=1−2t(n)(inparticulara(n)−b(n)takesonlythevalues±1)•(viii)a(b(n))−b(a(n))=4t(n)−2(inparticulara(b(n))−b(a(n))takesonlythevalues±2).Proof.TheﬁrstfourrelationsareTheorem1forthecased=2.Relations(v)and(vi)areeasyconsequencesofRelations(i)to(iv).Thelasttworelationsareconsequencesoftheexpressionsofa(n)andb(n)giveninCorollary1andofthepropertiest(2n)=t(n)andt(2n+1)=1−t(n).Onemightaskwhichsetofrelationsamongrelations(i)to(vi)suﬃcetocharacterizesequencesaandb.Thenextthreetheoremsyieldthreeanswerstothequestion.Theorem4SupposethatthetwosetsXandYformapartitionoftheintegers.Letx=(xn)n≥0betheincreasingsequenceoftheelementsofX,andlety=(yn)n≥0betheincreasingsequenceoftheelementsofY.Supposethatxandysatisfythefollowingrelations•x(x(n))=2x(n)foralln≥0•y(y(n))=2y(n)foralln≥0•|x(n)−y(n)|=1foralln≥0Then,eitherx=aandy=b,orx=bandy=a.Inparticularthesequence(x(n)−y(n))n≥0mustbeequalto(1−2t(n))n≥0orto(2t(n)−1)n≥0.Proof.Wemusthavethat{0,1}={x(0),y(0)}.Withoutlossofgeneralitywemaysupposethatx(0)=1thusy(0)=0.Wethuswanttoprovethatx=aandy=b.Wewillprovebyinductiononnthat{2n,2n+1}={x(n),y(n)}.Thepropertyistrueforn=0;supposeitistruefornandletuslookat{2n+2,2n+3}.Eitherthereexistsksuchthat2n+2=x(k)orthereexistsksuchthat2n+2=y(k)(XandYformapartitionoftheintegers).If2n+2=x(k)wehavenecessarily2n+3=y(k)(since|x(k)−y(k)|=1).Furthermorek≥n+1(sincexandyareincreasing).Ifwehadk≥n+2thevalues2n+2,2n+3wouldnotbeintherangeofxnorintherangeofy,hencek=n+1.If2n+2=y(k),thesamereasoningshowsthat2n+3=x(k),andk=n+1.9
Wethushave{2n+1,2n+3}={x(n+1),y(n+1)}andtheinductionisproven.Now,deﬁnethesequence(α(n))n≥0byx(n)=2n+1−α(n).Thisimpliesofcourseα(0)=0andy(n)=2n+α(n).Wethennotethat,foranyintegerm,wehave,byapplyingtheformulax(n)=2n+1−α(n)withn=x(m),ononehandx(x(m))=2x(m)+1−α(x(m)),andontheotherhandx(x(m))=2x(m).Thusα(x(m))=1.Inthesamewaywehaveforanyintegerm,usingtherelationy(n)=2n+α(n)forn=y(m),thaty(y(m))=2y(m)+α(y(m),whiley(y(m))=2y(m).Thusα(y(m))=0.SinceXandYformapartitionoftheintegersthisgivesn∈X⇔α(n)=1andn∈Y⇔α(n)=0.Nowweprovethatα(n)=t2(n),i.e.,thatthesequence(α(n))n≥0istheThue-Morsesequencebeginningwith0.Itsuﬃcestoprovethat,forallm≥0,wehaveα(2m)=α(m)andα(2m+1)=1−α(m).IfmbelongstoX,thenthereexistsaksuchthatm=x(k).Wehavejustseenthatα(m)=1.Wehavex(2m)=4m+1−α(2m).Butx(2m)=x(2x(k))=x(xx(k))=xx(x(k))=2xx(k)=4x(k)=4m.Henceα(2m)=1=α(m).Now,sincewethushavethat2mbelongstoX,wemusthave2m+1belongstoY,henceα(2m+1)=0.IfmbelongstoY,thenthereexistsaksuchthatm=y(k).Thusα(m)=0.Wehavey(2m)=4m+α(2m).Buty(2m)=y(2y(k))=y(yy(k))=yy(y(k))=2yy(k)=4y(k)=4m.Henceα(2m)=0.Now,sincewethushavethat2mbelongstoY,wemusthave2m+1belongstoX,henceα(2m+1)=1.Finallywethushavethat(α(n))n≥0=(t2(n))n≥0,andthenx=aandy=b.ThenexttwotheoremscanbeseenasvariationsonTheorem4.Theorem5Letx=(x(n))n≥0andy=(y(n))n≥0beincreasingintegersequencessuchthat{x(n),n≥0}∪{y(n),n≥0}=Nsatisfyingx(0)=1,y(0)=0,and∀n≥0,x(x(n))=y(x(n))−1andy(y(n))=x(y(n))−1.Thenxandyarerespectivelyequaltoaandbthesequencesofodiousandofevilnumbers.Proof.LetX={x(n),n∈N}andY={y(n),n∈N}.Theconditiononx(x(n))andy(y(n))canbewrittenasfollowsifm∈X,thenx(m)=y(m)−1;ifm∈Y,theny(m)=x(m)−1.ThisimpliesinparticularthatX∩Y=∅,thusXandYformapartitionoftheintegers.Nowlet1XbethecharacteristicfunctionofX(i.e.,1X(n)=1ifandonlyifnbelongsto10
X).Thus1−1XisthecharacteristicfunctionofY.Wewillprovebyinductiononnthat,foralln≥0,x(n)=2n+1−1X(n),y(n)=2n+1X(n).Thepropertyistrueforn=0sincex(0)=1andy(0)=0.Ifitistrueupton,weﬁrsthave{x(n),y(n)}={2n,2n+1}.•Ifn+1belongstoX,thenononehandx(n+1)=y(n+1)−1.Sincex(n+1)>max{2n,2n+1},thisgivesx(n+1)≥2n+2andy(n+1)=x(n+1)+1≥2n+3.ButXandYformapartitionoftheintegers,thusx(n+1)mustbeequalto2n+2(otherwise2n+2ismissedbothbyxandbyy),andy(n+1)=x(n+1)+1=2n+3.Thisgivesx(n+1)=2(n+1)+1−1X(n+1)andy(n+1)=2(n+1)+1X(n+1).•Ifn+1belongstoY,thenoneonehandy(n+1)=x(n+1)−1.Sincey(n+1)>max{2n,2n+1},thisgivesy(n+1)≥2n+2andx(n+1)=y(n+1)+1≥2n+3.ButXandYformapartitionoftheintegers,thusy(n+1)mustbeequalto2n+2(otherwise2n+2ismissedbothbyyandbyx),andx(n+1)=y(n+1)+1=2n+3.Thisgivesy(n+1)=2(n+1)+1X(n+1)andx(n+1)=2(n+1)+1−1X(n+1).Wethennotethatx(n)=2n+1−1X(n)foralln,impliesthatx(x(n))=2x(n)+1−1X(x(n))=2nforalln.Similarlyy(n)=2n+1X(n)forallnimpliesthaty(y(n))=2y(n)foralln.ButwehaveseenthataccordingtombeinginXorY,wehavey(m)−x(m)=±1,i.e.,|x(m)−y(m)|=1.WecanthenconcludeusingTheorem4.Theorem6Letx=(xn)n≥0andy=(yn)n≥0betwosequencesofintegersdeﬁnedbyx(0)=1,y(0)=0,andforeachn≥1,x(n)andy(n)arethesmallestintegersthatdidnotoccurbefore(i.e.,thatdonotbelongto{x(k),k≤n−1}∪{y(k),k≤n−1}),withtheconditionsthatforalln≥0•x(x(n))andy(y(n))areeven,•x(y(n))andy(x(n))areodd.Thenx=athesequenceofodiousnumbers,andy=bthesequenceofevilnumbers.Proof.Thehypothesis“thesmallestnumbersthatdidnotoccurbefore”impliesthatxandydonotmissanyinteger.Inotherwords,deﬁningX={x(n),n≥0}andY={y(n),n≥0},wehaveX∪Y=N.OntheotherhandtheintersectionofXandYisempty:ifnbelongsbothtoXandY,thenthereexistk,`withn=x(k)=y(`).Butthenx(n)=x(x(k))iseven,whilex(n)=(x(y(`))shouldbeodd,whichisimpossible.ThusXandYformapartitionoftheintegers.Wewillproveasabovethat,letting1XdenotethecharacteristicfunctionofX,thenx(n)=2n+1−1X(n)andy(n)=2n+1X(n).Thepropertyistrueforn=0.Supposethatitistrueupton,whichimpliesinparticularthat{x(n),y(n)}={2n,2n+1}.11
•Ifn+1belongstoX,i.e.,n+1=x(k)forsomek,thenx(n+1)=x(x(k))mustbeeven,whiley(n+1)=y(x(k)mustbeodd.Allintegervaluesbeingtaken,thisimpliesthatx(n+1)=2n+2andy(n+1)=2n+3.Thiscanalsobewrittenx(n+1)=2(n+1)+1−1X(n+1)andy(n+1)=2(n+1)+1X(n+1).•Ifn+1belongstoY,i.e.,n+1=y(k)forsomek,thenx(n+1)=x(y(k))mustbeodd,whiley(n+1)=y(y(k))mustbeeven.Allintegervaluesbeingtaken,thisimpliesthaty(n+1)=2n+2andx(n+1)=2n+3.Thiscanalsobewritteny(n+1)=2(n+1)+1X(n+1)andx(n+1)=2(n+1)+1−1X(n+1).Sinceforallnweclearlyhave|x(n)−y(n)|=1weconcludeasinTheorem5.6Conclusion
WewouldliketoaddthatallthefunctionalequationsgivenaboveforthesequencesofodiousandofevilnumberscanbetranslatedintermsofcharacterizationsoftheThue-Morsesequence.Furthermoreanalogousresultscanbeprovenforthesequencesad,j.References[1]J.-P.Allouche,N.Rampersad,J.Shallit,Onintegersequenceswhoseﬁrstiteratesarelinear,AequationesMath.69(2005)114–127.[2]J.-P.Allouche,J.Shallit,TheubiquitousProuhet-Thue-Morsesequence,inSequencesandtheirApplications(Singapore,1998),1–16,SpringerSer.DiscreteMath.Theor.Comput.Sci.,Springer,London,1999.[3]E.Berlekamp,J.Conway,R.Guy,WinningWaysforYourMathematicalPlays,vol.1and2,AcademicPress,1982.[4]E.Berlekamp,J.Conway,R.Guy,WinningWaysforYourMathematicalPlays,vol.3,SecondEdition,A.K.Peters,2003.[5]B.Cloitre,N.J.A.Sloane,M.J.Vandermast,NumericalanaloguesofAronson’sse-quence,J.IntegerSeq.6(2003)Art.03.2.2.[6]M.Emmer,NumeriMaleﬁci(EvilNumbers):HomagetoFabioMauri,inImagineMath2:BetweenCultureandMathematics,M.Emmered.,Springer,2013,pp.139–150.[7]R.K.Guy,Impartialgames,inCombinatorialGames,MSRIPublications,Vol.29,1995.Availableathttp://library.msri.org/books/Book29/files/imp.pdf[8]V.Laohakosol,B.Yuttanan,Iteratesofincreasingsequencesofpositiveintegers,Ae-quationesMath.87(2014)89–103.12
[9]F.W.Nietzche,BeyondGoodandEvil:PreludetoaPhilosophyoftheFuture,EditedbyRolf-PeterHorstmann,EditedandtranslatedbyJudithNorman,CambridgeTextsintheHistoryofPhilosophy,Cambridge,2002.[10]On-LineEncyclopediaofIntegerSequences,availableelectronicallyattheURLhttp://oeis.org[11]KarambirS.Sarkaria,Rootsoftranslations,AequationesMath.75(2008)304–307.[12]G.Targo´nski,Topicsiniterationtheory,StudiaMathematica:Skript,6,Vandenhoeck&Ruprecht,G¨ottingen,1981.[13]WolframMathworld.Availableathttp://mathworld.wolfram.com/EvilNumber.html13

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

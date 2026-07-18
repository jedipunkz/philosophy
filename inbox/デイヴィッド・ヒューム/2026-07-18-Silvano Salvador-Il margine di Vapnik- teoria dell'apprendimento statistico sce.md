---
source: "https://archive.org/details/pac-pagenumber-watermark"
title: "Il margine di Vapnik: teoria dell'apprendimento statistico scetticismo induttivo"
author: "Silvano Salvador"
year: ""
captured_at: "2026-07-18T19:27:50Z"
updated_at: "2026-07-18T19:27:50Z"
capture_tool: "scrapem-book"
source_name: "archive"
keyword: "デイヴィッド・ヒューム"
query: "Hume A Treatise of Human Nature"
plain_text_url: "https://archive.org/download/pac-pagenumber-watermark/PAC_pagenumber_watermark_djvu.txt"
public_domain: true
subjects:
tags:
  - "近代哲学"
  - "経験論"
  - "懐疑主義"
status: raw
---

# Il margine di Vapnik: teoria dell'apprendimento statistico scetticismo induttivo

- 著者: Silvano Salvador
- 情報源: [archive](https://archive.org/details/pac-pagenumber-watermark)
- パブリックドメイン: ✓

## Obsidian Links

- キーワード: [[デイヴィッド・ヒューム]]
- 研究動向: [[デイヴィッド・ヒューム-現代研究動向]]

## Full Text

Silvano Salvador

Il margine di Vapnik: teoria
dell'apprendimento statistico e
scetticismo induttivo

1. L'anatomia di un problema millenario

I. L'abisso di Hume

Esiste un momento preciso nella storia della filosofia in cui la ragione prende
coscienza di un proprio limite strutturale: il 1739, anno in cui David Hume pubblica il
Treatise of Human Nature e vi iscrive, quasi di passaggio, una delle osservazioni più
destabilizzanti mai formulate riguardo al pensiero umano. Il problema, nelle sue linee
essenziali, è il seguente. Ogni inferenza che procede dal passato al futuro,
dall'osservato all'inosservato , dal campione alla popolazione — vale a dire ogni
ragionamento che merita il nome di induttivo — poggia su di un principio che non
può essere giustificato né dalla logica né dall'esperienza senza incorrere in una
citcolarità viziosa. Non dalla logica, poiché la negazione di qualsiasi legge empirica è
sempre concepibile senza contraddizione: nulla impedisce, sul piano della mera
possibilità logica, che il sole sorga domani a occidente o che l'acqua, sotto zero gradi,
rimanga liquida. Non dall'esperienza, poiché qualunque tentativo di fondare
empiricamente l'uniformità della natura presuppone già ciò che si vuole dimostrare —
ovvero che il futuro assomigli al passato — e scivola così in una petitio principii di

proporzioni cosmiche.

La risposta humeana non è uno scetticismo paralitico bensì una ridescrizione della
natura  dell'inferenza induttiva: essa non è razionalmente giustificata ma
psicologicamente inevitabile, radicata nell'abitudine e nell'associazione. La mente
umana è, da questo angolo visuale, una macchina che produce aspettative senza
possedere la minima garanzia logica della loro correttezza. Questa conclusione, che
Kant avrebbe tratto dal sonno dogmatico per trasformarla nell'architrave della Critica
della ragion pura, rimane oggi — passati quasi tre secoli — sostanzialmente inconfutata
sul piano strettamente logico. Ciò che è mutato, e radicalmente, è il contesto
concettuale entro cui il problema viene situato.

Il Novecento ha ereditato la questione humeana e l'ha rimodellata attraverso due
movimenti intellettuali quasi contemporanei e assai differenti: il falsificazionismo di
Karl Popper, che proponeva di aggirare l'induzione sostituendola con la deduzione, e il
progetto carnapiano di un'induzione logica rigorosa fondata sulla probabilità. Né l'uno

Silvano Salvador

né l'altro ha risolto il nodo originario. Popper si è trovato a dover ammettere che la
scelta delle ipotesi da sottoporre a test — il momento generativo, creativo, propositivo
della ricerca scientifica — non è governabile da alcun algoritmo deduttivo e resta
irrazionale nel suo senso tecnico. Carnap, dal canto suo, ha visto il proprio programma
naufragare contro il paradosso elaborato da Nelson Goodman, il quale mostrava che
qualunque misura di conferma logicamente ben definita dipende da scelte

convenzionali che la logica sola non può sorreggere.

Il punto di ingresso della presente monografia è precisamente questo nodo: non si
tratta di proporre l'ennesima soluzione filosofica al problema dell'induzione, quanto di
mostrare che la teoria matematica dell'apprendimento statistico — sviluppata in modo
sistematico da Vladimir Vapnik e Alexei Cervonenkis a partire dagli anni Sessanta del
Novecento — costituisce un cambio di registro concettuale di prima grandezza. Essa
non dissolve il problema humeano, che rimane insoluto sul piano trascendentale, ma
lo riformula in termini che ne rivelano la struttura profonda e ne misurano, per la
prima volta in modo preciso, le condizioni di possibilità. Il margine epistemico tra
osservato e inosservato diventa quantificabile; lo scetticismo induttivo si trasforma da
atteggiamento filosofico in disuguaglianza matematica.

II. Il nuovo enigma dell'induzione: Goodman e la perversione dei
predicati

Prima di addentrarci nell'apparato formale, è necessario sostare sul contributo di
Nelson Goodman, poiché esso radicalizza il problema humeano in una direzione che
la teoria PAC deve fronteggiare esplicitamente. Nel 1955, in Fac Fiction, and Forecast,
Goodman introduce un predicato divenuto celebre: "gruc". La definizione è la
seguente: un oggetto è grye se e soltanto se è stato esaminato prima di un certo tempo
# ed è verde, oppure non è stato esaminato prima di 7 ed è blu. Questa definizione è
perfettamente coerente e precisamente formulata; non vi è nulla di logicamente
difettoso in essa.

Si consideri ora la nostra evidenza empirica ordinaria: tutte le smeralde esaminate fino
ad oggi erano verdi. Ma questo insieme di osservazioni è compatibile, in modo
altrettanto perfetto, con la generalizzazione "tutte le smeralde sono verdi" e con la
generalizzazione "tutte le smeralde sono grue" — a patto che il tempo 7 sia nel futuro
rispetto al momento presente. Le due ipotesi sono supportate in misura identica dai
dati disponibili, eppure le nostre aspettative convergono quasi universalmente verso la
prima. La risposta ovvia — che "verde" sia un predicato più semplice o più naturale di
"grue" — risulta, sotto esame, profondamente insoddisfacente: la semplicità non è una
proprietà assoluta dei predicati ma dipende dal linguaggio e dai sistemi di riferimento
adottati. In un linguaggio primitivo in cui "grue" e "bleen" fossero i termini primitivi,
"verde" risulterebbe il predicato disgiuntivo e complesso, e il parlante di quella lingua

avrebbe altrettante ragioni per proiettare "grue" nel futuro.

Silvano Salvador

Il paradosso di Goodman mostra che il problema dell'induzione non è soltanto
epistemologico — non riguarda esclusivamente la giustificazione delle nostre inferenze
— ma semantico e metafisico: investe la questione di quali predicati, quali proprietà,
quali classificazioni siano prosettabili, meritino di essere estese oltre il dominio
dell'osservato. Goodman chiama questa proprietà entrenchwent. i predicati proicttabili
sono quelli radicati nella pratica linguistica e inferenziale di una comunità, quelli che
hanno una storia d'uso sedimentata e inerziale. È una risposta pragmatica e
parzialmente convenzionalista che lascia aperta la domanda su cosa, se non la

convenzione, giustifichi alcune classificazioni come più naturali di altre.

Da un punto di vista formale, il paradosso può essere riformulato con precisione. Sia
X lo spazio delle istanze e sia H uno spazio di ipotesi, ovvero un insieme di funzioni
da X a {0, 1}. Il problema induttivo consiste nello scegliere, sulla base di un campione
finito S di istanze etichettate, un'ipotesi 4 E H che generalizzi correttamente alle
istanze non osservate. Goodman mostra che, senza vincoli su H, questa scelta è
radicalmente sottodeterminata: per qualsiasi campione finito e per qualsiasi
etichettatura di quel campione, esistono infinite ipotesi in H che concordano
perfettamente con il campione e divergono su ogni istanza non campionata. Uno
spazio di ipotesi non vincolato ha capacità infinita, e un agente che dispone di tale
spazio non impara nulla di stabile.

Questa osservazione, espressa in prosa da Goodman, sarà resa precisa dalla teoria VC
tramite la nozione di dimensione di Vapnik-Cervonenkis, oggetto centrale della
seconda parte. La proiettabilità goodmaniana corrisponde, nell'orizzonte formale, a
uno spazio di ipotesi di dimensione VC finita; la mancanza di entrenchment corrisponde
a una dimensione VC infinita, ovvero a una classe di ipotesi che non è
PAC-apprendibile. Il paradosso di Goodman è, tradotto nel lessico della teoria
dell'apprendimento, l'enunciato del fatto che non ogni spazio di ipotesi è apprendibile
da esempi finiti, e che la scelta dello spazio è un atto che precede e condiziona
qualsiasi inferenza induttiva.

III. La versione kripkeana: regole, infinito e uso finito

Un secondo approfondimento del problema humeano, che risuona con il paradosso di
Goodman pur provenendo da una tradizione differente, è quello elaborato da Saul
Kripke nella sua lettura di Wittgenstein, raccolta nel celebre Wirgenstein on Rules and
Private Language del 1982. Il problema kripkeano — designato talvolta "Kripkenstein"
per sottolineare l'ambiguità attributiva — non concerne l'induzione empirica nel suo
senso classico bensì il seguire una regola: in che modo possiamo sostenere che la
nostra pratica matematica o linguistica è governata da una regola determinata, quando
qualsiasi uso finito di un'espressione è compatibile con infinite regole distinte?

L'esempio privilegiato riguarda la funzione di addizione. Supponiamo di non aver mai
calcolato la somma di due numeri superiori a 57; ci si chiede di calcolare 68 + 57. La
risposta attesa è 125. Ma uno scettico potrebbe argomentare che la nostra pratica

Silvano Salvador

passata è compatibile anche con l'ipotesi che abbiamo sempre seguito la funzione 9445,
definita da:

xOy = x+ty se min(x, y) < 57
xO®y = 5 altrimenti

In questo caso la risposta corretta sarebbe 5, non 125. E nessuna evidenza
comportamentale pregressa, nessun fatto psicologico  sull'intenzione, nessuna
disposizione neurologica potrebbe, argomenta Kripke, discriminare in modo definitivo

tra le due interpretazioni. La regola è sottodeterminata dal suo uso finito.

La risonanza con il paradosso goodmaniano è trasparente: sostituendo "verde" con
"addizione" e "grue" con "quaddizione", ci troviamo di fronte alla medesima struttura.
L'uso finito di un predicato — o di una funzione — è compatibile con infinite
estensioni future, alcune delle quali devianti rispetto alle aspettative condivise. Ciò che
Kripke aggiunge è una dimensione normativa che in Goodman è solo latente: non si
tratta unicamente di quale estensione sia epistemicamente giustificata, ma di quale sia

corretta, di quale faccia parte della regola che l'agente seguiva davvero.

Per la teoria dell'apprendimento, questa dimensione normativa trova una traduzione
precisa nel rapporto tra errore empirico ed errore reale di generalizzazione. Un
classificatore che memorizza fedelmente un campione di addizioni ha errore empirico
nullo su quel campione; eppure potrebbe aver "appreso" la quaddizione anziché
l'addizione, manifestando un errore di generalizzazione elevato su istanze estranee al
dominio di addestramento. La teoria PAC quantifica le condizioni sotto le quali questo
disallineamento è controllabile, e lo fa in funzione della complessità dello spazio di
ipotesi — misurata dalla dimensione VC — e della cardinalità del campione.

IV. L'eredità positivista e il progetto di naturalizzazione

Prima di costruire l'apparato matematico, è opportuno collocare questa impresa nel
quadro più ampio del dibattito novecentesco sulla naturalizzazione dell'epistemologia.
Il programma quineano, avanzato nell'omonimo saggio del 1969, sosteneva che
l'epistemologia dovesse diventare un capitolo della psicologia empirica: anziché cercare
fondamenti normativi per la conoscenza, dovremmo descrivere i meccanismi
mediante i quali gli organismi acquisiscono credenze sul mondo. La proposta era
provocatoria quanto metodologicamente vaga: non specificava quali scienze empiriche
fossero pertinenti né in che modo i risultati empirici potessero rimpiazzare o integrare

le norme epistemiche.

La teoria dell'apprendimento statistico costituisce, da questa prospettiva, la versione
più rigorosa e matematicamente articolata del programma di naturalizzazione. Essa
non si limita a descrivere come gli agenti apprendono, ma fornisce condizioni necessarie
e sufficienti perché l'apprendimento avvenga in modo affidabile — ovvero perché il
divario tra errore empirico ed errore reale sia controllabile. Questa duplicità tra

Silvano Salvador

descrittivo e normativo, tra "cosa fa un sistema che apprende" e "cosa deve fare per
apprendere bene", è ciò che distingue la teoria PAC da una mera psicologia della

cognizione e la avvicina a una genuina epistemologia formalizzata.

Va tuttavia dissipato un equivoco fondamentale. Affermare che la teoria PAC
naturalizza l'epistemologia non implica che riduca le norme epistemiche a fatti causali;
significa piuttosto che le condizioni di possibilità dell'apprendimento affidabile —
espresse come teoremi dimostrabili — hanno risonanze strutturali profonde con i
requisiti normativi che la tradizione epistemologica ha storicamente imposto
all'induzione. La generalizzazione PAC è formale e probabilistica; lo scetticismo
humeano è logico e assoluto. Il rapporto tra i due non è di riduzione bensì di
trasposizione: il problema filosofico viene trascritto in un registro che ne consente un

trattamento rigoroso, rivelando al contempo i limiti intrinseci di qualsiasi risposta

soddisfacente.
È in questo spirito — trasposizione piuttosto che riduzione, illuminazione piuttosto
che dissoluzione — che la presente monografia intende procedere. La seconda parte

costruisce il formalismo matematico con la necessaria completezza; la terza ne esplora
le implicazioni filosofiche nel dettaglio; la quarta estende l'analisi ai sistemi biologici e
alle conseguenze per la filosofia della cognizione. L'obiettivo complessivo è mostrare
che il problema dell'induzione, lungi dall'essere una reliquia della speculazione
preanalitica, è la struttura portante di qualsiasi teoria seria dell'apprendimento —
artificiale o biologico — e che la teoria VC ne è, ad oggi, la formulazione più potente e

intellettualmente onesta di cui disponiamo.

2. L'apparato formale:  PAC-apprendibilità e
dimensione di Vapnik-Cervonenkis

V. Il quadro probabilistico: distribuzioni, ertori e il significato
dell'apprendimento

Il punto di partenza dell'intera costruzione è una scelta epistemologica che vale la pena
rendere esplicita prima di introdurre qualunque simbolo. L'apprendimento da esempi
viene modellato come un problema di stima statistica: esiste una distribuzione di
probabilità D fissa e ignota su uno spazio di istanze X, e una funzione obiettivo f: X
— {0, 1} — anch'essa ignota all'apprenditore — che si vuole approssimare. Il fatto
che D sia ignota è decisivo: significa che l'apprenditore non sa come sono distribuiti i
dati nel mondo, conosce solo i campioni che gli vengono presentati. Il fatto che D sia
fissa equivale a postulare che il mondo abbia un'uniformità strutturale sufficiente
perché i campioni passati siano informativi sui campioni futuri — ovvero è

precisamente il presupposto che Hume chiamava in causa come ingiustificato.

Formalmente, il quadro è il seguente. Sia X uno spazio arbitrario (lo spazio delle
istanze) e sia Y = {0, 1} lo spazio delle etichette. Una distribuzione di apprendimento è una

Silvano Salvador

misura di probabilità D su X X Y tale che per ogni (x, y) campionato da D si abbia y =

f(x) per qualche funzione obiettivo fissa. Un campione di addestramento è una sequenza

S= (1, yi), (x2,72), ...,.(&,y)

dove ogni (xi, yi) è estratto in modo indipendente e identicamente distribuito (i.i.d.)
secondo D.

Uno spazio di ipotesi H è un insieme di funzioni h : X > {0, 1}. Un a/grizzo di
apprendimento è una mappa che associa a ogni campione $ un'ipotesi h_S € H. L'errore
reale (o errore di generalizzazione) di un'ipotesi % rispetto alla distribuzione D è la
probabilità che / classifichi erroneamente un'istanza estratta casualmente:

LD) = Pr_fx< DX} [h) #0]

dove D_X denota la marginale di D su X. L'errore empirico sul campione S è invece la
frazione delle istanze di addestramento su cui / sbaglia:

L S®) := (1/m)- |{iE€ {1,....,m}:h@)# vil

La distinzione tra L_D e L_S è il nucleo tecnico di tutta la teoria, e la sua portata
filosofica è immediata: L_D misura quanto bene % si comporterà su dati futuri, ed è
l'unica quantità che ci interessa davvero; L_S è l'unica quantità direttamente
osservabile. Il problema dell'induzione diventa: sotto quali condizioni L_S@)
LD?

VI. PAC-apprendibilità: la definizione e il suo significato

La nozione di PAC-apprendibilità fu introdotta da Leslie Valiant nel 1984 nel suo
articolo fondativo "A Theory of the Learnable". L'acronimo PAC sta per Probably
Approximately Correct. un algoritmo di apprendimento deve produrre un'ipotesi che sia
approssimativamente corretta (con errore reale non superiore a e) con probabilità elevata
(almeno 1 — è), per parametri e e è liberamente scelti.

Definizione (PAC-apprendibilità). Una classe di ipotesi H è PAC-apprendibile se
esiste un algoritmo A tale che: per ogni e € (0, 1), per ogni è € (0, 1), per ogni
distribuzione D su X X Y, per ogni funzione obiettivo frealizzabile in H (ovvero tale
che esista h* E H con L_D(h*) = 0), esiste una dimensione di campione 70(€, 6) tale
che per ogni campione $S di dimensione m 2 mo(e, è), estratto i.i.d. da D, si abbia:

Pr_S[L_D(A6) > e] < è

Silvano Salvador

La funzione mo(e, d) è detta complessità campionaria (sample complexity) della classe H. Si
noti che la condizione imposta è uniforme in D e in f l'algoritmo deve funzionare per
qualsiasi distribuzione e per qualsiasi funzione obiettivo realizzabile in H. Questa
uniformità è filosoficamente fondamentale: corrisponde all'esigenza che la garanzia di
apprendimento non dipenda da assunzioni specifiche sul "tipo di mondo" in cui ci
troviamo, ma valga pet qualsiasi mondo compatibile con lo spazio di ipotesi scelto.

La prima conseguenza immediata è che la dimensione del campione mo deve essere
finita e dipendente solo da e e è, non da Do f Questo è già un risultato non banale:
afferma che per qualsiasi e e è prefissati, esiste un numero finito di esempi sufficiente
a garantire la generalizzazione — indipendentemente da quale sia la vera distribuzione
del mondo. La dipendenza da H attraverso mo è invece essenziale, come vedremo.

VII. La funzione di crescita e il lemma di Sauer-Shelah

Per caratterizzare la PAC-apprendibilità, dobbiamo misurare in modo preciso la
"capacità" di uno spazio di ipotesi, ovvero quanto variegato e flessibile esso sia nel
separare le istanze. La nozione tecnica centrale è quella di shaztering e di funzione di
crescita.

Definizione (Restrizione e shattering). Sia C = {x1, ...,.x} C Xuninsieme finito

di istanze. La restrizione di H a C è l'insieme di tutte le dicotomie che H induce su C:

H_Cc:= {MhC®), hh, .... h&)):hE H} S {0,1}7m

Si dice che H frantuma (shatters) C se H_C = {0,1}®m, ovvero se ogni possibile
etichettatura binaria di C è realizzata da qualche ipotesi in H. In termini di Goodman:
H frantuma C se, per qualsiasi modo di colorare gli elementi di C come "verdi" o

"non-verdi", esiste un predicato in H che realizza esattamente quella colorazione.

Definizione (Funzione di crescita). La /unzione di crescita di H è la funzione 7_H : N
— N definita da:

t_H(m) := max_{C S X, |C|=m} |H_C|

Essa conta il massimo numero di dicotomie distinte che H può produrre su un
insieme di 77 punti. Si ha sempre t_H(m) £ 2°m, con uguaglianza se e solo se H

frantuma alcuni insiemi di cardinalità 7.

Il risultato chiave che governa la funzione di crescita è il cosiddetto Lemma di
Sauer-Shelah (dimostrato indipendentemente da Norbert Sauer e Saharon Shelah nel
1972):

Silvano Salvador

Lemma (Sauer-Shelah). Se la dimensione VC di H (che definiremo tra breve) è d <

9, allora per ogni m 2 d:

t_H(m) < X_{i=0}"{d} C(m,i)

dove C(m, i) = ml! / (il * (m_i)!) è il coefficiente binomiale. In particolare, per m 2 d si
ha:

t_H(m) £ (em/dYd

dove e è la base del logaritmo naturale. Dunque la funzione di crescita, che in principio
potrebbe essere esponenziale in m, diventa polinorziale in m non appena la dimensione
VC è finita.

Dimostrazione del Lemma di Sauer-Shelah. Dimostriamo per induzione su | X| + m che se
VC(H) < d allora | H_C| £ X_{i=0}"®{d} C(|C]|, i) per ogni C finito.

Caso base: se |C| =00d=0, la disuguaglianza vale banalmente. Se d = 0 e |C| 2 1,
allora H non frantuma nessun sottoinsieme non vuoto di C, il che implica che tutte le
ipotesi in H concordano sull'etichettatura di ogni punto — dunque |H_C| = 1 =
C(|C|, 0).

Passo induttivo: sia C = C' U {x} con x € C'. Partizioniamo H_C' in due parti: sia Yo
l'insieme delle dicotomie su C' che non hanno un "gemello" — ovvero l'etichettatura
di C' non cambia al variare dell'etichetta di x — e sia Y1 l'insieme delle dicotomie su C'
che hanno un gemello in H_C' con l'etichetta di x invertita. Allora |H_C| = |Yo| +
2|Ya].

Osserviamo che la classe di ipotesi ristretta a C' ha VC-dimensione al più d (se fosse
maggiore, anche la restrizione a C avrebbe VC-dimensione maggiore di d). Inoltre la
sottoclasse che produce le coppie di gemelli — ristretta a C' — ha VC-dimensione al
più d-1, poiché qualsiasi insieme frantumato da Yi su C' può essere esteso con x a un
insieme frantumato da H (contraddicendo VC(H) £ dì).

Per ipotesi induttiva:
|Yo|] + |Yi|] £ 2_{i=0}7{d} C(|C'|, î

[Ya] < 2_{i=0}7{d-1} C([C'|,i)

Dunque:

|H_C| = |[Yo| +2[Yi] = ([Yo] + [Yi]) + |Yal]
< X_{i=0}7{d} C(|C'|, ij) + X_{i=0}7{d-1} C([C'|,i)
>_{i=0}{d} [C([C'], i) + C(|C'],i-1)]

Il

Silvano Salvador

con la convenzione C(n, —1) = 0. Applicando l'identità di Pascal C(n,j) + C(a,j-1) =
C(n+1,j):

|H_C| < X_fi=0}"{d} C(|C']+1,î) = X_{i=0}"{d} C(|C|,i)

che è quanto si voleva dimostrare. I
La stima (em/d)?d si ottiene dall'analisi asintotica delle somme binomiali: per m 2 d,

®_{i=0}"{d} C(mjj) < (em/d)d

come conseguenza della disuguaglianza C(m,i) < (em/i)i applicata termine pet
termine e del confronto della somma con il termine dominante.

VIII. La dimensione di Vapnik-Cervonenkis

Definizione (Dimensione VC). La dizzensione di Vapnik-Cervonenkis di H, indicata
VC(H) o VCdim(H), è il massimo intero 4 tale che esiste un sottoinsieme C SE X di
cardinalità 4 frantumato da H:

VCdim(H) := supf |C| :C E X, H_C= {0,1}"|C]}

Se tale supremo non è finito, si scrive VCdim(H) = 20.

La dimensione VC è la misura della capacità espressiva di uno spazio di ipotesi che la
teoria ha eletto a parametro fondamentale. Pochi esempi ne illustrano il significato

geometrico:

Esempio 1: semipiani nel piano. Sia X = RR? e sia H la classe degli indicatori dei
semipiani chiusi, ovvero h_{w,b}(x) = 1 se w-x 2 b, 0 altrimenti, per w E R?,b E R.
Tre punti in posizione generale nel piano possono essere frantumati: per qualsiasi
etichettatura binaria di tre punti non collineari esiste un semipiano che separa gli uni
dagli altri. Quattro punti, invece, non possono essere frantumati in modo uniforme da
semipiani (questo è il contenuto geometrico del teorema di Radon). Dunque
VCdim(H) = 3, che coincide con il numero di parametri liberi della classe (dimensione

del vettore w più l'intercetta b, ovvero 3).

Esempio 2: soglie sulla retta reale. Sia x =ReH={h_0:0 E R} conh_0x =
1 se x 2 0. Qualsiasi singolo punto {xo} può essere frantumato: basta scegliere 0 £ xo
per ottenere l'etichetta 1, 0 > xo per l'etichetta 0. Due punti {x1, x2} con x1 < x: non
possono essere frantumati: l'etichettatura (1, 0) — "primo positivo, secondo negativo"
— è irraggiungibile da qualsiasi funzione soglia monotona crescente. Dunque

VCdim(H) = 1.

Silvano Salvador

10

Esempio 3: reti neurali a soglia. Per una rete neurale con unità e » pesi, la
dimensione VC è dell'ordine O(w - log(w/n)), un risultato dovuto a Baum e Haussler
(1989). Questo segnala che la capacità di una rete cresce con il numero di parametri,

ma in modo sublineare.

IX. Il Teorema Fondamentale dell'Apprendimento Statistico

Siamo ora in posizione di enunciare e dimostrare il teorema centrale dell'intera teoria,
che fornisce la risposta formale alla domanda humeana: sotto quali condizioni

l'esperienza passata garantisce la generalizzazione futura?

Teorema (Fondamentale dell'Apprendimento — versione qualitativa). Sia H
una classe di ipotesi. Le seguenti affermazioni sono equivalenti:

(@ H è PAC-apprendibile nel modello realizzabile. (i) H ha la proprietà di
convergenza uniforme delle frequenze empiriche (PUFE). (ili) Ogni algoritmo di
minimizzazione dell'errore empirico (ERM) è un buon apprenditore per H. (iv)

VCdim(H) < 0.

La dimostrazione completa è lunga; percorreremo i passi principali dell'implicazione
(iv) > 0 e 0) > (iv), che costituiscono il cuore dell'equivalenza.

Parte 1: VCdim(H) < co > PAC-apprendibilità (con stima della complessità

campionaria).

Vogliamo mostrare che se VCdim(H) = d < ©, allora con un campione di dimensione
sufficientemente grande, l'errore reale di qualsiasi ipotesi con errore empirico nullo è
piccolo con alta probabilità. Il passaggio cruciale è la tecnica della doppia campionatura e
l'utilizzo del Lemma di Sauer-Shelah.

Sia S un campione di dimensione m estratto i.i.d. da D. Sia h_S l'ipotesi restituita

dall'algoritmo ERM, ovvero qualsiasi h E H con L_S(b) = 0. Vogliamo maggiorare:

Pr S|[3hEH:L Db>e A LS) =0]

Introduciamo un campione fantasma (ghost sample) S' di dimensione m, estratto i.i.d.
dalla stessa distribuzione D. Si ha che:

Pr S[AKEH:LDM>e A LSM =0]
< 2-Pr_{SS}[3hEH:L {S}M)>e/2 A LS&=0]

purché m 2 2/e (questa disuguaglianza si ottiene osservando che se L_D(h) > e, allora
la probabilità che L_{S'}(h) > e/2 è almeno 1/2 per un campione di dimensione m >
2/e, per la legge dei grandi numeri).

Silvano Salvador

11

Ora fissiamo S e S', consideriamo l'unione S U S' come un insieme di 2m punti, e
osserviamo che la condizione "L_{S"} (h) > e/2 e L_S(h) = 0" dipende soltanto dalla
restrizione di h a questi 2m punti. Il numero di restrizioni distinte di H a S U S' è al
più t_H(2m) £ (2em/d)?d per il Lemma di Sauer-Shelah. Dunque:

Pr {SS}[IhEH:L_{S}@®)>e/2 A L SO) =0]
< 1 Hm): Pro[L_{S'_o}h_0)>e/2 /\ L_{S_o}h_o = 0]

dove la probabilità interna è su una permutazione casuale o dei 2m punti tra S e S'
(argomento di simmetria di Rademacher). Questa probabilità, per qualsiasi h fissato
con L_D(h) > e, può essere maggiorata da 2-exp(-em/8) attraverso la disuguaglianza
di Chernoff.

Combinando:

Po S[IhEH:LD@)>e A LS®=0]
2:0tHQm : 2-exp(-em/8)
4: (2em/d)°d - exp(-em/8)

IATA

Vogliamo che questa espressione sia £ è. Prendendo logaritmi e risolvendo per m, si

ottiene che è sufficiente:

m > (8/9): [d - In(em/d) + In(4/8) ]

che (per disuguaglianza autocontenuta, sostituendo m con una stima dall'alto) dà la

complessità campionaria:

mof(e, è) = O((d/e) - In(1/8) + (1/9) - In(1/3))

Parte 2: PAC-apprendibilità = VCdim(H) < 0.

Supponiamo per assurdo che VCdim(H) = 0. Allora per ogni m, esiste un insieme di
m punti frantumato da H. Costruiamo una distribuzione avversariale nel modo
seguente: sia C_m = {x1, ..., x } un insieme frantumato da H, e sia D la distribuzione
uniforme su C_m. Poiché H frantuma C_m, per ogni etichettatura y : C_m — {0,1}
esiste h_y € H con h_y(x)) = y(xj) per ogni i.

Ora l'algoritmo di apprendimento A riceve un campione $S di m punti da D (estratti
con ripetizione). Con probabilità (1 — 1/e) almeno m/e punti di C_m non compaiono
in S. Su questi punti, qualsiasi algoritmo non ha alcuna informazione, e la funzione
obiettivo potrebbe etichettarli in modo arbitrario. Formalmente, per qualsiasi
algoritmo A, esiste una scelta di f (realizzabile in H) tale che:

Pr_S[L_D(AG) > 1/8] > 1/7

Silvano Salvador

12

Questo mostra che per e = 1/8 e è = 1/7, nessun algoritmo può garantire
PAC-apprendibilità, contraddicendo l'ipotesi. Il (schema)

X. La disuguaglianza di Vapnik-Chervonenkis e la convergenza
uniforme

Il teorema precedente fornisce una caratterizzazione qualitativa. La versione
quantitativa è la celebre disuguaglianza VC, che lega la deviazione dell'errore empirico
dall'errore reale in modo uniforme su tutta la classe H.

Teorema (Disuguaglianza VC). Sia H una classe di ipotesi con VCdim(H) = d < 00.
Per ogni distribuzione D e ogni è > 0, con probabilità almeno 1 — è sul campione S di
dimensione m, si ha simultaneamente per ogni h € H:

| LD) -L_S(h) | < V((2d - In(em/d) + 2-In(2/8)) /m)

Chiamiamo il termine a destra margine di generalizzazione o gap di generalizzazione. Esso
dipende da tre fattori: cresce con d (maggiore è la capacità, maggiore è il margine di
incertezza), decresce con m (più esempi, meno incertezza), e decresce con è (per una
garanzia più debole si paga meno in termini di dati).

La disuguaglianza afferma che, con alta probabilità, nessuna ipotesi in H può avere un
errore reale molto peggiore del suo errore empirico. In termini filosofici: se lo spazio
di ipotesi ha capacità finita e il campione è abbastanza grande, l'induzione è affidabile
— nel senso preciso che l'evidenza empirica non inganna in modo sistematico. Questo
non elimina la possibilità che ipotesi singole ingannino in casi particolari (la probabilità
è 1 — è, non 1), né giustifica il principio di uniformità della natura (la distribuzione D
rimane un assunto). Ma quantifica esattamente il prezzo dell'induzione: il numero di
esempi necessari, la tolleranza all'errore, la confidenza statistica.

Schizzo di dimostrazione. Il risultato si ottiene combinando tre ingredienti. Primo, la
tecnica del campione fantasma riduce la questione alla simmetria di Rademacher.
Secondo, il Lemma di Sauer-Shelah limita il numero di classi di equivalenza su cui la
simmetria deve essere verificata. Terzo, la disuguaglianza di McDiarmid (o di
Hoeffding applicata) controlla le fluttuazioni di ogni funzione che dipende da variabili
i.i.d. attraverso le variazioni limitate.

Più precisamente, per ogni h E H, la mappa $S' L_S(b) ha variazioni limitate da 1/m
(cambiare un punto cambia l'errore empirico al più di 1/m). La disuguaglianza di
McDiarmid afferma che per ogni funzione g(z1,...,z) con variazione limitata ci nel
coordinato i-esimo:

Pr g(Z) — E[g(A] 2 t] £ exp(-2t° / Zio?)

Silvano Salvador

13

Applicando questo a g(S) = sup_{hEH} (L_D(h) — L_S(b)) e combinando con il
Lemma di Sauer-Shelah per maggiorare il numero di comportamenti distinti di H su
qualsiasi campione di dimensione m, si perviene alla stima enunciata. Il

XI. Il ruolo del bias induttivo: scelta di H come atto epistemico

La disuguaglianza VC mette in piena luce un fatto che la filosofia dell'induzione aveva
intravisto senza riuscire a quantificarlo: qualsiasi apprendimento affidabile richiede un
bias induttivo, ovvero una restrizione a priori dello spazio delle ipotesi ammissibili.
Questo non è un difetto degli algoritmi di apprendimento — è una necessità logica,
come mostrato dall'implicazione (i) = (iv) del Teorema Fondamentale.

Uno spazio di ipotesi con VCdim(H) = % non è PAC-apprendibile in modo uniforme.
Qualunque algoritmo che abbia accesso a tale spazio può essere ingannato da qualche
distribuzione avversariale: vi saranno sempre dati di addestramento che sembrano
supportare un'ipotesi errata. In altre parole, l'assenza di bias induttivo — l'assunzione
che tutte le ipotesi siano a priori ugualmente plausibili, che corrisponde formalmente
all'accesso a uno spazio di capacità infinita — rende l'apprendimento impossibile.
Questo risultato, spesso citato come "No Free Lunch Theorem", afferma che non
esiste alcun algoritmo universalmente buono per tutti i problemi di apprendimento:
qualsiasi algoritmo che si comporta bene su alcune distribuzioni si comporta male su

altre.

La scelta di H è dunque un atto epistemico di primaria importanza, analogo alla scelta
del sistema di predicati proiettabili di Goodman. Non si tratta di un'operazione
neutrale o puramente tecnica: incorpora assunzioni substantive sul tipo di regolarità
che ci si aspetta di trovare nel mondo. Una classe di ipotesi di piccola dimensione VC
codifica l'assunzione che il mondo sia "semplice" — che la funzione obiettivo possa
essere approssimata da una funzione di bassa complessità. Una classe di grande
dimensione VC codifica maggiore flessibilità, al prezzo di richiedere campioni più

numerosi per garantire la generalizzazione.

Questo è precisamente il nucleo del dilezzzza bias-varianza, che nella sua formulazione
più elegante si esprime come segue. Sia h* = argmin_{h EH} L_D(h) la migliore
ipotesi raggiungibile in H (errore di approssimazione o bias), e sia h_S l'ipotesi

appresa dall'algoritmo ERM sul campione S. L'errore totale si decompone:

L_Dh_S) = [L DAh_S) - LDbh*] + LD)
= Errore di stima + Errore di approssimazione

= Varianza + Bias

AI crescere della ricchezza di H (VCdim(H) — ©), l'errore di approssimazione tende a
zero (o diminuisce) perché H diventa capace di approssimare meglio qualsiasi
funzione obiettivo; ma l'errore di stima aumenta, perché la convergenza uniforme
richiede campioni sempre più grandi. Al diminuire di H, accade il contrario. Il bias

Silvano Salvador

14

induttivo ottimale è quello che bilancia i due termini per la dimensione del campione
disponibile.

La scelta di H è dunque una scommessa sulla struttura del mondo: scommettere su
uno spazio ricco significa credere che il mondo sia complesso e che si disponga di dati
abbondanti; scommettere su uno spazio povero significa credere che il mondo sia
regolare e accettare il rischio di approssimare male funzioni obiettivo complesse.
Nessuna scelta è epistemicamente neutrale, e nessuna scelta può essere giustificata
puramente a posteriori dai dati — poiché è la scelta stessa che determina cosa i dati

possono insegnare.

XII. Dimensione VC e geometria: il Teorema di Radon come
fondamento

Un aspetto che merita approfondimento è il legame tra la dimensione VC e la
geometria degli spazi di istanze, particolarmente suggestivo per i classificatori lineari.
Per la classe degli iperpiani in Rd, ovvero i classificatori della forma h_{w,b}®) =
sign(w-x + b) conw E Rd eb E R, si può dimostrare che:

VCdim(H _lin°d) = d+1

La dimostrazione usa il Teorema di Radon: qualsiasi insieme di d+2 punti in Rd può
essere partizionato in due sottoinsiemi con inviluppi convessi che si intersecano.
Questa proprietà geometrica fondamentale implica che d+2 punti in posizione
generale non possono essere frantumati da iperpiani in Rd, poiché la dicotomia
"convex hull" non è linearmente separabile.

Più in dettaglio: per dimostrare che VCdim(H_lin°d) 2 d+1, si costruisce un insieme
di d+1 punti frantumabile. Siano ei, ..., e_d i vettori della base canonica di R°d e sia
co = 0. L'insieme {co, ei, ..., ed} ha d+1 elementi. Per qualsiasi etichettatura y :
{0,1,...d} —> {-1, +1}, vogliamo trovare w e b tali che sign(w‘ei + b) = y()). Le
condizioni sono:

wer+b20 se yg=+1
wei+b<0 se yd =-1

Per i 2 1, w*ei = wi (la i-esima componente di w). Per i = 0, w-eo = 0, dunque la
condizione è b = 0 o b < 0. Scegliendo wi = y(î) — b e b opportuno (tipicamente b =
2: (max y(i) + min y()), si verifica che il sistema ha sempre soluzione, dunque
l'insieme è frantumato.

Per dimostrare VCdim £ d+1, si usa il Teorema di Radon: per qualsiasi d+2 punti
{xo,....x_{d+1}} in IRd, esiste una bipartizione {S+, S-} con S: U S- =
{x0,....x_{d+1}}, S: N S- = 9, tale che conv(S+) N conv(S-) # 2. Sia z un punto

Silvano Salvador

15

nell'intersezione degli inviluppi convessi. Allora z = X_{xjE S+} Xxj = ®_{x ES}

px conZA = Xp =1,X20,p ZO. Per qualsiasi iperpiano h_{w,b}, se tutti i punti
di S+ fossero etichettati +1 e tutti quelli di S- fossero etichettati —1, si avrebbe:

wz+b = Zi xi (wxi+ b) 2 0 (da S4)
x u(wx +b) < 0 (da S-)

w'z + b

che è contraddittorio. Dunque la dicotomia (S+ > +1, S- + —1) non è realizzabile da
nessun iperpiano, e i d+2 punti non sono frantumati. I

Questo risultato mostra che la dimensione VC di classificatori lineari in Rd coincide
con la dimensione affine dello spazio + 1, ovvero il numero di parametri liberi del
modello. Questa coincidenza, che vale per classificatori lineari ma roy in generale (per
reti neurali, ad esempio, la dimensione VC è proporzionale a w:log w e non a w),
suggerisce che la dimensione VC cattura qualcosa di più sottile del mero conteggio dei
parametri: misura la capacità effettiva di discriminazione geometrica del modello.

3. Il teorema come risposta filosofica: portata,
limiti e risonanze

XIII. Che cosa ha dimostrato Vapnik, filosoficamente parlando

Occorre essere chirurgicamente precisi su ciò che il Teorema Fondamentale
dell'Apprendimento stabilisce e su ciò che non stabilisce, prima di esaminarne la
portata epistemologica. Non aver chiarezza su questo punto ha generato, nella
letteratura filosofica che ha assorbito la teoria PAC, una serie di equivoci che val la

pena dissipare fin dall'inizio.

Il teorema dimostra che la PAC-apprendibilità è equivalente alla finitezza della
dimensione VC. Non dimostra che un agente con dimensione VC finita apprenderà la
funzione obiettivo vera; dimostra che, con campioni sufficientemente grandi, l'errore
reale di qualsiasi ipotesi che si comporta bene sul campione sarà vicino all'errore reale
della migliore ipotesi realizzabile nello spazio. Questa è una garanzia di consistenza
uniforme, non di verità. La distinzione è tutt'altro che sottile: uno spazio di ipotesi con
VCdim = d potrebbe non contenere affatto la funzione obiettivo vera, nel qual caso
anche l'ipotesi ottimale ha un errore reale positivo, e la garanzia PAC riguarda
esclusivamente la distanza tra questa ipotesi ottimale irraggiungibile e quella appresa

dall'algoritmo. Il termine tecnico è errore di approssimazione o bias, e il teorema non ne
dice nulla.

In secondo luogo, il teorema è una garanzia probabilistica, non deterministica.
Afferma che con probabilità almeno 1 — è sul campione casuale l'errore reale è
piccolo; non esclude che esistano campioni sfortunati su cui l'algoritmo fallisce. La

Silvano Salvador

16

probabilità in questione è calcolata rispetto alla distribuzione di campionamento, che
rimane un assunto non derivato — è precisamente il presupposto humeano di
uniformità della natura, vestito in abiti probabilistici.

Terzo, il teorema è uniforme nella distribuzione D solo all'interno di una classe di
problemi: tutti i problemi con funzione obiettivo realizzabile in H e con distribuzione
qualsiasi su X. Ma la scelta di H non è coperta dalla garanzia: è un atto che precede il

teorema e che il teorema non giustifica.

Con queste precisazioni in mente, la portata filosofica del teorema è comunque
straordinaria. Esso afferma che il problema dell'induzione ha una so/uzione condizionata:
se l'agente è disposto a vincolarsii a uno spazio di ipotesi di dimensione VC finita, e se
il mondo obbedisce a qualche distribuzione stabile (non necessariamente nota), allora
l'induzione è affidabile nel senso preciso e quantificabile della PAC-apprendibilità. La
condizionalità non è un difetto della teoria — è il rispecchiamento onesto della
struttura del problema. Hume aveva ragione nel sostenere che l'induzione non può
essere giustificata incondizionatamente; la teoria VC ha il merito di trasformare questa
impossibilità in una precisazione: l'induzione è giustificabile condizionatamente al bias
induttivo e all'uniformità della distribuzione, e il prezzo di questa giustificazione è
esattamente quantificabile.

XIV. Il No Free Lunch come principio trascendentale negativo

Il No Free Lunch Theorem (NFL), dimostrato da Wolpert e Macready nel 1997 nella
sua versione più generale, afferma che nessun algoritmo di apprendimento ha
prestazioni mediamente migliori di qualsiasi altro quando si fa la media su tutte le
possibili funzioni obiettivo. In termini formali, se si denota con A un algoritmo e con
E_f]L D(A(S)] l'errore atteso mediato su una distribuzione uniforme su tutte le
funzioni f : X — {0,1}, allora pet qualsiasi coppia di algoritmi Ai e A°:

Y_f E_S[L_f(A:(S)] = £_f E_S[L_f(A2:(8)]

dove la somma è su tutte le 2%{2"|X|} funzioni binarie su X (supponendo X finito).
Ogni algoritmo che supera il caso in certi domini paga il prezzo delle prestazioni
peggiorate in altri domini; nessuna strategia di apprendimento è universalmente

superiore alle altre.

La connessione con la filosofia dell'induzione è trasparente e merita di essere
formulata in modo esplicito. Il Teorema NFL è la versione formale dell'argomento
humeano: senza assunzioni a priori sulla struttura del problema (senza bias induttivo),
nessun metodo di generalizzazione è preferibile a qualsiasi altro. La giustificazione
dell'induzione richiede necessariamente un impegno strutturale — una scommessa su
quale tipo di regolarità è probabile trovare — che non può esso stesso essere derivato

dall'esperienza senza circolarità.

Silvano Salvador

17

Ciò che la teoria PAC aggiunge rispetto all'argomento di Hume è una caratterizzazione
precisa di che tipo di impegno strutturale è necessario e sufficiente per rendere
l'apprendimento affidabile. Non è sufficiente assumere vagamente che "il futuro
assomigli al passato" — questa formulazione è troppo vaga per consentire garanzie
quantitative. È invece necessario e sufficiente che lo spazio di ipotesi abbia dimensione
VC finita, ovvero che l'agente abbia ristretto a priori la propria attenzione a funzioni
che non frantumo insiemi arbitrariamente grandi di punti. Questa restrizione è
precisamente un'assunzione sulla regolarità del mondo: il mondo è di un tipo che può

essere approssimato da funzioni di complessità controllata.

Il NFL può quindi essere letto come un principio trascendentale negativo, in senso
kantiano: stabilisce le condizioni sotto le quali l'esperienza non può insegnare nulla. La
sua controparte positiva — il ‘Teorema Fondamentale dell'Apprendimento —
stabilisce le condizioni sotto le quali l'esperienza può insegnare qualcosa. Insieme, i
due risultati delimitano con precisione lo spazio in cui il progetto epistemologico
dell'induzione è realizzabile: è lo spazio delle classi di ipotesi di dimensione VC finita,

ovvero lo spazio in cui il bias induttivo è non banale.

Questa dualità ha una profonda analogia con la struttura dell'epistemologia kantiana.
Le Forme pure dell'intuizione e le Categorie dell'intelletto — spazio, tempo, causalità,
sostanza — sono esattamente un bias induttivo a priori: restringono lo spazio delle
esperienze possibili a quelle strutturate in modo tale da consentire generalizzazioni
affidabili. Un'esperienza senza queste forme sarebbe, nella terminologia della teoria
PAC, un'esperienza con spazio di ipotesi di capacità infinita — non apprendibile, non
generalizzabile, radicalmente sottodeterminata da qualsiasi campione finito. Le
categorie kantiane sono, da questa prospettiva, le condizioni di possibilità
dell'apprendimento, non semplicemente dell'esperienza in senso fenomenologico:
sono il bias induttivo che rende la scienza newtoniana possibile.

XV. Goodman rivisitato: l'entrenchment come vincolo VC

La connessione tra il paradosso di Goodman e la dimensione VC può ora essere
sviluppata nella sua piena portata. Ricordiamo che Goodman distingueva tra predicati
proiettabili (euzrezched) e non proiettabili: solo i primi possono essere legittimamente
estesi oltre il dominio dell'osservato in un'inferenza induttiva. La domanda di
Goodman — cosa rende un predicato proiettabile? — rimase senza risposta
soddisfacente nella sua opera: la risposta "l'uso consolidato nella comunità" è

pragmaticamente descrittiva ma normativa mente inerte.

La teoria VC offre una risposta parziale ma strutturalmente illuminante. Un predicato
(o, più precisamente, un insieme di predicati che costituisce uno spazio di ipotesi) è
proiettabile se e solo se ha dimensione VC finita. Questa non è una risposta alla
domanda normativa — non dice perché dobbiamo scegliere predicati con VC finita,
né giustifica la preferenza per "verde" su "grue" dal punto di vista della struttura del
mondo. Ma è una risposta condizionale: se vogliamo che i nostri predicati siano tali da

Silvano Salvador

18

consentite apprendimento affidabile da esempi finiti, allora dobbiamo scegliere

predicati con dimensione VC finita.

Il predicato "grue" ha, considerato singolarmente, dimensione VC = 1 (è una funzione
soglia temporale, analoga alle funzioni soglia sulla retta reale discusse nell'esempio 2
della sezione VIII). Non è dunque più o meno proiettabile di "verde" nel senso VC. Il
paradosso di Goodman emerge non a livello di singoli predicati ma a livello di famiglie
di predicati: se lo spazio di ipotesi H include sia tutte le funzioni "tempo-relative"
come "grue" sia tutte quelle "tempo-indipendenti" come "verde", e in generale tutte le

combinazioni arbitrarie, allora H ha capacità infinita e non è PAC-apprendibile.

La soluzione goodmaniana, tradotta nel linguaggio VC, è dunque questa: dobbiamo
scegliere uno spazio di ipotesi che escluda le funzioni tempo-telative oppure, più
precisamente, che restringa in modo sufficiente la famiglia delle funzioni ammissibili
da renderla di dimensione VC finita. Questa scelta è un atto convenzionale — dipende
dalla nostra teoria del mondo, dai nostri impegni ontologici — ma non è arbitrario nel
senso di essere indifferente rispetto all'obiettivo dell'apprendimento affidabile.
Scegliere male lo spazio di ipotesi — includere troppe funzioni della forma "grue" —
penalizza l'apprendimento richiedendo campioni proporzionalmente più grandi. In
questo senso, il teorema VC fonda una forma di pragmatismo epistemologico: i
predicati  proiettabili sono quelli che, in quanto famiglia, consentono un
apprendimento efficiente, e questa efficienza è misurabile in termini di complessità
campionaria.

Vale la pena spingersi oltre e notare che la teoria VC suggerisce anche una nuova
prospettiva sul realismo sui tipi naturali. La domanda filosofica classica — esistono
classi naturali oggettive, o le nostre classificazioni sono puramente convenzionali? —
riceve dalla teoria PAC una risposta pragmatica-strutturale: le classificazioni che
corrispondono a tipi naturali sono quelle che, come spazio di ipotesi, consentono
apprendimento efficiente nelle distribuzioni effettivamente incontrate. Un tipo
naturale è, in questa prospettiva, una classe che il mondo "premia" statisticamente —
una classe di ipotesi tale che la dimensione VC sia piccola rispetto alla ricchezza delle
regolarità che cattura. Questa non è una risposta metafisica alla domanda sul realismo,
ma è una risposta funzionale: i tipi naturali si guadagnano il proprio posto

nell'economia della conoscenza attraverso la loro efficienza induttiva.

XVI. Kripkenstein e la normatività dell'apprendimento

Il problema kripkeano della regola pone una sfida ulteriore che la teoria PAC affronta
in modo più indiretto ma non meno significativo. Kripke aveva osservato che qualsiasi
uso finito di una funzione è compatibile con infinite regole distinte; la domanda era
quale fatto — psicologico, semantico, platonico — determina quale regola l'agente stia
seguendo. La risposta della teoria PAC non è una risposta a questa domanda
normativa; è piuttosto una dissoluzione parziale della sua urgenza epistemica.

Silvano Salvador

19

Il problema kripkeano presuppone che esista una distinzione netta tra "seguire la
regola dell'addizione" e "seguire la regola della quaddizione", e che questa distinzione
debba essere determinata da qualche fatto. La teoria PAC suggerisce un modo diverso
di inquadrare la questione: invece di chiedere quale regola l'agente seguisse in passato,
possiamo chiedere quale classe di funzioni consente la migliore generalizzazione per i
problemi effettivamente incontrati dall'agente. Se l'agente opera in un ambiente in cui
l'addizione è la struttura corretta e la quaddizione è una deviazione con errore reale
elevato, allora il campione di addestramento, se sufficientemente grande, separa le due
ipotesi: un algoritmo ERM su uno spazio che include sia addizione sia quaddizione
convergerà sull'addizione.

Il punto delicato è che questa risposta funziona solo se lo spazio di ipotesi è
abbastanza ricco da includere entrambe le funzioni ma abbastanza limitato da avere
dimensione VC finita. Se lo spazio include tutte le funzioni possibili da N° a N (incluse
infinità di deformazioni della quaddizione), il problema di generalizzazione diventa
intrattabile nel senso PAC. La soluzione kripkeana nella prospettiva VC è dunque:
l'agente "seguiva l'addizione" se e solo se operava con uno spazio di ipotesi che
includeva l'addizione come ipotesi a basso errore reale rispetto alla distribuzione degli
input effettivi, e tale spazio aveva dimensione VC sufficiente a garantire che il

campione osservato distinguesse l'addizione dalla quaddizione.

Questa risposta è deflazionistica rispetto alla domanda di Kripke: non identifica un
fatto semantico che determina la regola seguita, ma sostituisce la questione della
determinazione con la questione della distinguibilità statistica. Due regole sono "la
stessa" per un agente se nessun campione di dimensione ragionevole può distinguerle
rispetto alla distribuzione degli input rilevanti; sono "diverse" se la distinzione emerge
dall'apprendimento. La normatività dell'apprendimento diventa così una normatività
statistica, non semantica o platonica.

XVII. La struttura probabilistica come presupposto non derivato

Occorre affrontare frontalmente l'obiezione più seria che si può muovere all'intera
costruzione: tutto l'apparato formale presuppone una distribuzione di probabilità
stabile D sullo spazio delle istanze, ma questa stabilità è esattamente il principio di
uniformità della natura che Hume considerava ingiustificabile. Non stiamo
semplicemente riformulando il problema invece di risolverlo?

L'obiezione è corretta nella sua premessa ma errata nella sua conclusione. Il
presupposto probabilistico non è derivato né giustificato dalla teoria PAC: è
un'assunzione di framework, esattamente come in qualsiasi modello statistico. La
teoria non pretende di dimostrare che il mondo è stabile o che le distribuzioni sono
stazionarie; presuppone questa stabilità. e ne deriva le conseguenze per
l'apprendimento. In questo senso, la teoria PAC è condizionata dall'assunto
probabilistico nello stesso modo in cui la geometria euclidea è condizionata dai propri

assiomi: se gli assiomi valgono, i teoremi seguono.

Silvano Salvador

20

Ma questo non rende la teoria epistemologicamente vacua. Il guadagno concettuale
consiste nel seguente: Hume aveva posto un problema qualitativo — l'induzione non
può essere giustificata — senza specificare in quale senso e a quale prezzo. La teoria
PAC trasforma questo problema qualitativo in una struttura quantitativa condizionale:
dando per buona la stabilità della distribuzione, l'induzione è giustificabile a un costo
preciso, misurato in termini di dimensione del campione e dimensione VC. Sappiamo
esattamente quanto dobbiamo pagare, in dati, per ogni unità di garanzia; sappiamo
esattamente quale tipo di struttura a priori dobbiamo assumere. Questo non è
risolvere Hume — è renderlo preciso, il che è un progresso di natura diversa ma non

inferiore.

Vi è, inoltre, un modo di interpretare il presupposto probabilistico che lo avvicina a
qualcosa di più difendibile epistemologicamente. La distribuzione D non deve essere
pensata come una proprietà oggettiva del mondo, ma come un modello dell'agente
sulla propria incertezza rispetto alle istanze che incontrerà. In questa lettura bayesiana
— o meglio, in una lettura de Finettiana — la stabilità di D esprime non un fatto sul
mondo ma una coerenza interna delle credenze dell'agente: l'agente si aspetta di
incontrare istanze distribuite in modo approssimativamente stabile nel tempo, e il
teorema PAC gli dice che questa aspettativa è sufficiente per rendere l'apprendimento
affidabile. La circolarità humeana non è eliminata, ma viene ricondotta alla coerenza
delle credenze piuttosto che alla struttura oggettiva del mondo, il che è filosoficamente

più difendibile.

XVIII. Occam's Razor e la teoria VC: una giustificazione formale?

Un tema che attraversa tutta la storia della filosofia della scienza è il principio di
parsimonia, comunemente attribuito a Guglielmo di Ockham: "entia non sunt
multiplicanda praeter necessitatem". Nella metodologia scientifica, il rasoio di Occam
è invocato come criterio di preferenza per le teorie più semplici a parità di adeguatezza
empirica. Ma cosa significa esattamente "più semplice"? E perché la semplicità

dovrebbe essere una virtù epistemica?

La teoria VC offre una risposta formale parziale a entrambe le domande. Nel contesto
dell'apprendimento statistico, la "semplicità" di uno spazio di ipotesi è misurata dalla
sua dimensione VC: uno spazio più semplice ha dimensione VC più bassa. E la
ragione per cui la semplicità è una virtù epistemica è precisamente la disuguaglianza
VC: a parità di errore empirico, un'ipotesi scelta da uno spazio di dimensione VC
inferiore ha un margine di generalizzazione migliore, ovvero è più probabile che il suo
errore reale sia vicino al suo errore empirico. Il rasoio di Occam ha dunque, in questo
contesto, una giustificazione formale: non è un principio metafisico sulla struttura del
mondo, né un'abitudine cognitiva priva di fondamento razionale, ma una strategia di

minimizzazione della varianza nello spazio delle ipotesi.

Va tuttavia segnalato un limite importante. La teoria VC giustifica la preferenza pet
ipotesi di spazi a più bassa dimensione VC quando queste abbiano errore empirico

Silvano Salvador

21

comparabile; non dice nulla sulla scelta tra un'ipotesi con bassa dimensione VC e alta
errore empirico e un'altra con alta dimensione VC e basso errore empirico. Il
bilanciamento ottimale tra semplicità e adeguatezza ai dati non è risolto dalla
disuguaglianza VC da sola, ma richiede il principio di minimizzazione strutturale del
rischio (SRM) proposto da Vapnik stesso, che seleziona lo spazio di ipotesi che
minimizza il maggiorante dell'errore reale, ovvero la somma dell'errore empirico e del
margine VC:

argmin_{H_k} [L_S@h_{S}k}) + V((@-d_k-In(em/d_b) + 2-In(2/3)) / m)]

dove la minimizzazione è su una sequenza annidata Hi S Ho S ... < H_kS ... di
spazi con dimensioni VC crescenti di £ d2 £ ..., e h_{S,k} è l'ipotesi ERM in H_k. Il

principio SRM seleziona automaticamente il livello di complessità appropriato in

funzione della dimensione del campione disponibile: con pochi dati, favorisce spazi
semplici; con molti dati, può permettersi spazi più ricchi. Questo è il rasoio di Occam

formalizzato come criterio di ottimizzazione.

XIX. Generalizzazione, memoria e la critica al deep learning

Un'applicazione contemporanea e controversa dei concetti VC riguarda il fenomeno
della generalizzazione nelle reti neurali profonde, che ha posto alla teoria
dell'apprendimento statistico classica alcune sfide di prima grandezza. Le reti neurali
moderne hanno tipicamente un numero di parametri molto maggiore del numero di
istanze di addestramento — sono, nella terminologia classica, drasticamente
sovraparametrizzate. Secondo la teoria VC tradizionale, questo dovrebbe comportare
un'elevata varianza e una scarsa generalizzazione. Eppure, empiricamente, tali reti

generalizzano spesso in modo sorprendentemente buono.

Questo fenomeno — noto come "doppia discesa" nella letteratura recente, descritto
da Belkin, Hsu, Ma e Mandal nel 2019 — mostra che la relazione tra complessità del
modello e errore di generalizzazione non è monotona nel regime sovraparametrizzato:
l'errore di test dapprima cresce all'aumentare della complessità (come predice la teoria
VC classica), poi sorprendentemente decresce nel regime di interpolazione, dove il
modello ha abbastanza parametri da memorizzare perfettamente il campione di

addestramento.

Questo fenomeno non contraddice il Teorema Fondamentale dell'Apprendimento, ma
ne segnala i limiti come strumento analitico per le reti neurali moderne. Il Teorema
garantisce che uno spazio di ipotesi di dimensione VC finita è PAC-apprendibile, ma
non dice che l'algoritmo ERM sia l'unico modo di apprendere, né che la dimensione
VC sia la misura di complessità più rilevante in ogni contesto. Le reti neurali addestrate
con discesa stocastica del gradiente non sono algoritmi ERM nel senso classico:
l'ottimizzazione inerziale del gradiente introduce un bias implicito verso soluzioni di

bassa norma o bassa variazione, che non è catturato dalla dimensione VC.

Silvano Salvador

22

In termini filosofici, il fenomeno della doppia discesa suggerisce che il bias induttivo
delle reti neurali non risiede unicamente nello spazio di ipotesi esplicitamente definito
(tutte le reti con una certa architettura) ma nell'interazione tra l'architettura, l'algoritmo
di ottimizzazione e la struttura dei dati. Il bias induttivo reale è implicito, emergente, e
non interamente riconducibile alla dimensione VC. dell'architettura. Questa
osservazione apre una questione filosofica di grande rilievo: se il bias induttivo
effettivo di un sistema di apprendimento non è trasparente all'agente che lo usa — se
emerge dalla dinamica dell'ottimizzazione piuttosto che da scelte architetturali esplicite
— in che senso possiamo dire che l'agente "sceglie" un sistema di predicati proiettabili
nel senso di Goodman? E in che senso le garanzie PAC, che sono garanzie su spazi di
ipotesi espliciti, si applicano a sistemi il cui spazio effettivo di ipotesi è implicito e
contestuale?

Queste domande non hanno ancora risposta soddisfacente nella letteratura.
Costituiscono il fronte aperto più interessante dell'intersezione tra filosofia

dell'induzione e teoria dell'apprendimento.

4. Il bias induttivo nei sistemi biologici e le
conseguenze per la filosofia della cognizione

XX. Il cervello come apprenditore PAC: un'ipotesi strutturale

Il passaggio dalla teoria dell'apprendimento artificiale ai sistemi cognitivi biologici non
è semplicemente una metafora suggestiva: è una tesi strutturale che, se corretta,
impone conseguenze precise sull'architettura della cognizione. L'ipotesi che si intende
argomentare in questa parte è la seguente: qualsiasi sistema cognitivo che apprende
affidabilmente dal mondo deve, per necessità matematica, essere dotato di un bias
induttivo che corrisponde a uno spazio di ipotesi di dimensione VC finita. Questo non
è un'ipotesi empirica sull'architettura neurale — è una conclusione logica dal Teorema
Fondamentale dell'Apprendimento, applicata a qualsiasi agente che si qualifichi come
apprenditore affidabile nel senso PAC.

Il cervello umano è, di fatto, un apprenditore sorprendentemente efficiente: riesce a
generalizzare da pochissimi esempi in molti domini, fenomeno indicato come few-shot
learning nella letteratura cognitiva e computazionale. Un bambino apprende il concetto
di "cane" da meno di una decina di esemplari e lo generalizza correttamente a
esemplari mai visti, incluse razze morfologicamente diverse da quelle incontrate. Un
adulto può acquisire una regola sintattica da pochi esempi e applicarla a strutture
sintattiche nuove. Questo comportamento, dal punto di vista della teoria PAC, è
possibile solo se il sistema cognitivo opera con uno spazio di ipotesi di dimensione VC
molto bassa rispetto al numero di esempi disponibili — ovvero solo se il sistema è
fortemente biasato a priori verso determinate classi di generalizzazioni.

Silvano Salvador

23

Questa considerazione porta direttamente alle teorie nativiste della cognizione.
Chomsky sosteneva che la grammatica universale è un vincolo innato sullo spazio
delle grammatiche possibili che un bambino considera nell'acquisizione linguistica:
senza tale vincolo, nessun bambino potrebbe acquisire una lingua da un input così
ridotto e disordinato. La teoria PAC fornisce una giustificazione formale di questo
argomento — noto in linguistica come "argomento dalla povertà dello stimolo" —
che va ben oltre l'intuizione originale. Non si tratta di un'ipotesi sulla ricchezza della
struttura innata del linguaggio, ma di una necessità logica: qualsiasi acquisizione
linguistica affidabile da campioni finiti e rumorosi richiede un bias induttivo sufficiente
a rendere finita la dimensione VC dello spazio delle grammatiche considerate. La
disputa tra nativismo e empirismo nella psicologia cognitiva è, riformulata in termini
VC, una disputa sulla fonte del bias induttivo: o è innato (geneticamente codificato) o
è acquisito da un apprendimento di ordine superiore (meta-apprendimento). Ma che il
bias debba esistere — e che debba corrispondere a un vincolo sufficiente a garantire la
finitezza della dimensione VC — non è discutibile senza abbandonare l'ipotesi che il
sistema apprenda in modo affidabile.

XXI. Complessità campionaria e sviluppo ontogenetico

La formula della complessità campionaria ottenuta nella sezione IX,

mole, è) = O((d/e) - In(1/8) + (1/9) - In(1/3) )

dove d = VCdim(H), ha implicazioni non banali per lo sviluppo cognitivo se si assume
che il cervello del bambino in sviluppo operi con spazi di ipotesi la cui dimensione VC

aumenta progressivamente con la maturazione neurale.

Si consideri il seguente scenario idealizzato. Nei primi mesi di vita, il sistema visivo
dispone di uno spazio di ipotesi fortemente vincolato — solo pattern di basso ordine
come bordi, orientamenti, contrasti — con dimensione VC bassa, diciamo di. Questo
consente una generalizzazione affidabile da pochi esempi su queste caratteristiche
primitive. Con lo sviluppo, le aree associative ampliano lo spazio di ipotesi disponibile
(da > di), e il sistema può apprendere concetti di ordine superiore, ma a costo di
richiedere più esempi per garantire la stessa precisione e e la stessa confidenza 1 — è.
La progressione temporale dello sviluppo cognitivo sarebbe dunque, da questa
prospettiva, un itinerario ottimale attraverso spazi di ipotesi di dimensione VC
crescente, calibrato sulla disponibilità di esempi in ciascuna fase: prima si apprendono
le strutture di bassa complessità (che richiedono pochi esempi), poi quelle di alta
complessità (che richiedono molti esempi accumulati).

Questa previsione è coerente con diverse osservazioni dello sviluppo cognitivo. Il fatto
che i bambini apprengano prima le categorie di livello basico (nella terminologia di
Rosch) rispetto a quelle di livello superordinato o subordinato è coerente con l'idea
che le prime abbiano una dimensione VC inferiore — richiedono meno esemplari pet
essere apprese in modo affidabile. Il fatto che la sensibilità fonologica al contrasto si

Silvano Salvador

24

restringa nei primi mesi di vita (il bambino smette di discriminare contrasti fonetici
non pertinenti alla propria lingua) può essere interpretato come una riduzione
deliberata della dimensione VC dello spazio delle ipotesi fonologiche: il sistema
"sceglie" uno spazio di ipotesi più piccolo (solo i contrasti pertinenti alla propria

lingua) che richiede meno dati per la generalizzazione fonologica affidabile.

Il modello è ovviamente una semplificazione grossolana: il cervello non ottimizza
esplicitamente la dimensione VC, e la maturazione neurale non è un processo di
espansione graduale di spazi di ipotesi ben definiti. Tuttavia, la struttura formale
suggerisce che esiste un principio di organizzazione dello sviluppo cognitivo che può
essere compreso in termini di bilanciamento bias-varianza: i sistemi cognitivi biologici
sarebbero configurati per operare, in ciascuna fase dello sviluppo, con uno spazio di
ipotesi la cui dimensione VC è approssimativamente ottimale rispetto alla quantità di
dati disponibili in quella fase.

XXII. Meta-apprendimento e il bias come oggetto di
apprendimento

Una delle intuizioni più feconde che emergono dalla congiunzione tra teoria PAC e
neuroscienze cognitive è che il bias induttivo stesso può essere oggetto di
apprendimento — non in una singola sessione di apprendimento su un problema
fisso, ma attraverso un processo di meta-apprendimento che si svolge su scale
temporali diverse.

Formalmente, il meta-apprendimento può essere modellato nel modo seguente.
Supponiamo che l'agente non affronti un singolo problema di apprendimento con
distribuzione fissa D e funzione obiettivo fissa f ma una sequenza di problemi (Di, fi),
D:, fd), ..., (D, f) campionati da qualche distribuzione soprascritta 7 sullo spazio dei
problemi. L'obiettivo del meta-apprendimento è ottimizzare la scelta dello spazio di
ipotesi H — o equivalentemente, il bias induttivo dell'agente — in modo da
minimizzare l'errore atteso su problemi futuri campionati da 7. Questo è il quadro del
meta-apprendimento proposto indipendentemente da Thrun e Pratt (1998) e
formalizzato da Baxter (2000).

Il risultato chiave di Baxter afferma che, se il numero di problemi osservati 7 è
sufficientemente grande e se il numero di esempi per problema 7 è sufficientemente
grande, allora l'errore di meta-gencralizzazione — la differenza tra la performance
dell'agente su nuovi problemi e la performance ottimale con il miglior bias possibile

rispetto a 7 — è maggiorata da:

O(V(VCdim(B) / n) + V(VCdim(H) /m))

dove B è lo spazio dei possibili bias (spazi di ipotesi) considerati, VCdim(B) è la sua
dimensione VC, e H* è il bias ottimale selezionato. Il primo termine misura il costo
dell'incertezza sulla scelta del bias ottimale (decresce con il numero di problemi

Silvano Salvador

25

osservati), il secondo misura il costo dell'apprendimento entro ciascun problema
(decresce con il numero di esempi per problema).

Questa struttura a due livelli ha implicazioni profonde per la filosofia della cognizione.
Suggerisce che il bias induttivo di un sistema cognitivo maturo non è interamente
fissato geneticamente, né interamente acquisito dall'esperienza entro singole sessioni
di apprendimento: è il prodotto di un processo di meta-apprendimento che si svolge
su scale temporali lunghe (la storia evolutiva della specie, lo sviluppo ontogenetico
dell'individuo, l'accumulo culturale della comunità) e che ottimizza progressivamente il
bias rispetto alla distribuzione di problemi effettivamente incontrati. La specie umana,
in questa prospettiva, ha impiegato milioni di anni di pressione evolutiva per acquisire
un bias induttivo — codificato in parte nell'architettura neurale innata, in parte nelle
pratiche culturali tramandate — che è approssimativamente ottimale rispetto alla
distribuzione di problemi cognitivi caratteristici del nostro ambiente.

Questo è un programma naturalista fortemente diverso da quello quineano. Non si
tratta di ridurre l'epistemologia alla psicologia empirica, ma di mostrare che il bias
induttivo appropriato per la cognizione umana è il risultato di un processo di
meta-apprendimento che può essere analizzato con gli strumenti formali della teoria
PAC estesa. L'epistemologia diventa, da questa prospettiva, una teoria normativa delle
condizioni di ottimalità del meta-apprendimento.

XXIII. Il problema della misura e le limiti della formalizzazione

È necessario, a questo punto, introdurre alcune cautele critiche che una trattazione
onesta non può eludere. La teoria PAC è stata sviluppata per problemi di
classificazione binaria su spazi discreti o compatti ben definiti; la sua applicazione alla
cognizione biologica richiede astrazioni che possono distorcere il fenomeno di
interesse.

Il primo problema è quello della misura. La dimensione VC è definita come il
massimo numero di punti frantumabili, ma per molti problemi cognitivi interessanti
non è chiaro come definire lo "spazio delle istanze" X in modo che la dimensione VC
sia calcolabile 0 anche solo finita in modo non banale. Il concetto di "cane" non ha un
dominio di istanze discreto e ben definito; la grammatica di una lingua naturale non è
facilmente rappresentabile come elemento di uno spazio di ipotesi su cui la
dimensione VC abbia un senso computabile. Le applicazioni della teoria PAC alla
cognizione biologica sono, per questa ragione, più spesso qualitative che quantitative:
parlano di "bassa dimensione VC" come di un requisito strutturale senza fornire stime
numeriche.

Il secondo problema è che la teoria PAC è una teoria del caso peggiore: garantisce la
generalizzazione per qualsiasi distribuzione D, mentre i sistemi cognitivi biologici
operano in ambienti che non sono distribuzioni avversariali ma distribuzioni
fortemente strutturate, correlate, non stazionarie. In ambienti reali, l'apprendimento

affidabile può richiedere meno dati di quanto predice la complessità campionaria nel

Silvano Salvador

26

caso peggiore, perché la struttura dell'ambiente riduce effettivamente la complessità
del problema. Questo è il motivo per cui un bambino apprende i concetti in molti
meno esempi di quanto la teoria VC classica suggerirebbe: l'ambiente cognitivo degli
esseri umani è strutturato in modo tale da ridurre il costo dell'induzione al di sotto del

limite teorico del caso peggiore.

Il terzo problema riguarda la stazionarietà. La teoria PAC presuppone che la
distribuzione D sia stabile nel tempo; ma i sistemi cognitivi biologici operano in
ambienti non stazionari, in cui la distribuzione degli input cambia nel tempo, talvolta
in modo radicale. L'apprendimento in ambienti non stazionari — il cosiddetto conzinza!
learning © apprendimento incrementale — non è catturato dal framework PAC
standard e richiede estensioni teoriche che sono ancora oggetto di ricerca attiva. Dal
punto di vista filosofico, questo significa che l'analogia tra la teoria PAC e la
cognizione biologica è un'approssimazione utile ma non una riduzione completa.

XXIV. Razionalità limitata, vincoli computazionali e la complessità
campionaria come risorsa

Un'ultima connessione merita attenzione: quella tra la teoria PAC e il programma della
razionalità limitata di Herbert Simon. Simon sosteneva che gli agenti razionali non
ottimizzano globalmente su spazi di scelta illimitati, ma operano con euristiche che
sono adatte all'ambiente specifico in cui si trovano. Il concetto di bounded rationality è
una risposta alla discrepanza tra i requisiti computazionali dell'ottimizzazione globale e

le risorse cognitive effettivamente disponibili.

La teoria PAC può essere letta come una formalizzazione di questo programma: il bias
induttivo di un agente è precisamente la restrizione della sua razionalità allo spazio di
ipotesi H, e la dimensione VC di H misura la complessità di questa restrizione. Un
agente razionalmente illimitato avrebbe accesso a uno spazio di ipotesi di capacità
infinita — ma, per il NFL Theorem, sarebbe un agente che non impara nulla in modo
affidabile. La razionalità limitata non è dunque una deviazione dall'ideale della
razionalità piena: è una condizione necessaria per l'apprendimento affidabile. Gli
agenti razionalmente limitati — nel senso di operare con spazi di ipotesi di dimensione
VC finita — sono l'unico tipo di agente capace di generalizzare dall'esperienza.

Il costo della limitatezza è la complessità campionaria: un agente con VCdim(H) = d
richiede dell'ordine di O(d/e : log 1/e) esempi per apprendere con precisione e.
Questo costo è una risorsa, non solo una limitazione: quantifica quanta esperienza è
necessaria per formare credenze affidabili. In questo senso, la teoria PAC fonda una
forma di epistemologia della risorsa — un'epistemologia che misura il costo
dell'acquisizione della conoscenza non in termini di giustificazione logica (il progetto
fondazionalista) né in termini di coerenza delle credenze (il progetto coerentista), ma

in termini di dati necessari e capacità di generalizzazione garantita.

Silvano Salvador

27

XXV. Conclusione: il margine come struttura della conoscenza

Siamo in grado, a questo punto, di raccogliere i fili dell'arvomentazione sviluppata in
g ,aq ; g 8 PP

queste pagine e di enunciare la tesi complessiva nella sua forma definitiva.

Il problema dell'induzione, nella sua formulazione humeana originaria, è l'impossibilità
di giustificare logicamente il passaggio dall'osservato all'inosservato . Le sue versioni
successive — goodmaniana, kripkeana — ne hanno rivelato la struttura semantica e
normativa, mostrando che non si tratta solo di un problema epistemologico ma di una
questione che investe la natura dei predicati, delle regole e delle classificazioni. La
teoria VC non dissolve nessuno di questi problemi nel loro aspetto trascendentale; ma
li traspone in un registro formale che ne rivela la struttura con una precisione finora

itraggiunta.

La tesi centrale è questa: la PAC-apprendibilità — ovvero la possibilità dell'induzione
affidabile — è equivalente alla finitezza della dimensione VC dello spazio di ipotesi
dell'agente. Questa equivalenza non è una metafora né un'analogia: è un teorema. Le
sue implicazioni filosofiche sono precise. L'induzione richiede necessariamente un bias
a priori; tale bias è equivalente a un vincolo sulla capacità espressiva dello spazio di
ipotesi; questo vincolo è quantificabile; il costo dell'induzione affidabile — in termini

di dati necessari — è esplicitamente calcolabile in funzione di questo vincolo.
Il margine di generalizzazione

V( (24 + In(em/d) + 2:In(2/8)) /m)

è, letteralmente, la misura dell'incertezza induttiva: è la distanza massima tra ciò che
l'esperienza osservata insegna (errore empirico) e ciò che vale nel mondo più in largo
(errore reale), garantita con probabilità 1 — è dopo aver visto m esempi, per un agente
con spazio di ipotesi di dimensione VC = d. Questo numero è la struttura quantitativa

del problema dell'induzione.

Hume aveva ragione: l'induzione non può essere giustificata incondizionatamente. Il
Teorema Fondamentale dell'Apprendimento dice qualcosa di più preciso: può essere
giustificata condizionatamente, e il prezzo della condizionalità si misura in dimensione
VC, numero di esempi e probabilità di errore. La conoscenza empirica ha un costo; la

teoria VC ne è il bilancio.

5. La complessità di Rademacher: quando il
mondo conta più del caso peggiore

XXVI. I limiti della dimensione VC come misura universale

Silvano Salvador

28

La dimensione di Vapnik-Cervonenkis è una misura di complessità combinatorica:
dipende unicamente dalla struttura dello spazio di ipotesi H e non dalla distribuzione
D dei dati. Questo la rende universalmente applicabile — garantisce la
generalizzazione per qualsiasi distribuzione — ma al tempo stesso la penalizza sul
piano della finezza: quando la distribuzione dei dati è lontana dal caso peggiore
(ovvero quando la natura dei dati effettivi semplifica il problema rispetto allo scenario
avversariale), la disuguaglianza VC è pessimistica, talvolta in modo drastico.

Un esempio illuminante. Sia H la classe di tutti i classificatori lineari in R1000, con
VCdim(H) = 1001. La disuguaglianza VC prescrive una complessità campionaria
dell'ordine di O(1001/e : log 1/6) — migliaia di esempi anche per tolleranze moderate.
Ma se la distribuzione D è concentrata su una sottovarietà a bassa dimensione
intrinseca — poniamo, su una curva unidimensionale nello spazio — allora il
problema effettivo ha la complessità di un problema lineare in R', non in R71000, e
basterebbe dell'ordine di O(1/e : log 1/e) esempi. La teoria VC non vede questa

semplificazione perché ignora la distribuzione.

Questo difetto non è solo tecnico: ha un significato epistemologico preciso. La teoria
VC misura la difficoltà induttiva nel caso peggiore — ovvero per il "mondo più
avversariale" compatibile con lo spazio di ipotesi scelto. Ma la conoscenza empirica
non avviene nel caso peggiore: avviene nel mondo reale, con una distribuzione
specifica e spesso benevolente. Una teoria dell'induzione che sia davvero informativa
deve tenere conto non solo dello spazio di ipotesi ma anche della distribuzione dei dati

su cui si lavora.

La nozione che risponde a questa esigenza è la complessità di Rademacher, introdotta nel
contesto dell'apprendimento statistico da Bartlett e Mendelson nel 2002, seguendo
contributi precedenti di Koltchinskii e Panchenko. Essa misura la capacità di uno
spazio di ipotesi di correlarsi con rumore casuale su/ campione effettivamente disponibile —
una misura dipendente dalla distribuzione che raffina la teoria VC e fornisce garanzie
più strette nelle situazioni pratiche.

XXVII. Variabili di Rademacher e la definizione formale

Prima di introdurre la complessità di Rademacher, è necessario stabilire alcune
premesse formali. Una variabile di Rademacher è una variabile aleatoria o che assume i
valori +1 e —1 con ugual probabilità 1/2. Le variabili di Rademacher sono l'oggetto
probabilistico più elementare che modella il rumore casuale puro: non hanno struttura,

non hanno correlazioni, sono massimamente imprevedibili.

L'idea chiave è la seguente. Dato un campione $ = (xi, ..., x) e uno spazio di ipotesi
H, la complessità di Rademacher misura quanto bene le funzioni in H riescono, in
media, a "fittare" etichette assolutamente casuali 01, ..., 0 applicate agli stessi punti del
campione. Se H è abbastanza ricco da fittare qualsiasi sequenza di +1, allora in
particolare può fittare etichette casuali — ovvero può memorizzare il rumore — e la
generalizzazione sarà scarsa. Se invece H è troppo rigido per fittare il rumore, allora

Silvano Salvador

29

qualunque ipotesi con basso errore empirico sta veramente catturando struttura, non

rumore, e la generalizzazione sarà buona.

Definizione (Complessità di Rademacher empirica). Dato un campione S = (x1,
...,x) estratto da D, la complessità di Rademacher empirica di H rispetto a S è:

R_S(H) := E_o [sup_fh € H} (1/m)- X_{i=1}"{m} o‘ h&)]

dove o = (01, ..., co) è un vettore di variabili di Rademacher indipendenti, e
l'aspettazione è rispetto a o per S fissato.

Definizione (Complessità di Rademacher). La complessità di Rademacher di H
rispetto alla distribuzione D per campioni di dimensione m è:

R_m(H) := E_S [R_S(H)] = E_{S,o} [sup_{h € H} (1/m)- X_{i=1}"7{m} ai:
h(x;) ]

Notare la struttura: si estrae un campione S da D, si campionano etichette casuali o, e
si misura quanto bene la migliore ipotesi in H riesce a fittare queste etichette casuali
sul campione S. L'aspettazione è sia sul campione sia sulle variabili di Rademacher.

La quantità R_S(H) è calcolabile (almeno approssimativamente) dal campione, il che la
rende uno strumento adatto a fornire garanzie data-dipendenti: invece di una
limitazione uniforme che vale per qualsiasi distribuzione, si ottiene una limitazione che
si adatta alla distribuzione effettiva dei dati.

XXVIII. Il teorema di generalizzazione di Rademacher

Il risultato principale che collega la complessità di Rademacher alla generalizzazione è

il seguente:

Teorema (Generalizzazione tramite Rademacher). Sia H una classe di funzioni da
X a {0, 1}. Per ogni distribuzione D su X e per ogni è > 0, con probabilità almeno 1 —
è sul campione $ di dimensione m, si ha per ognih € H:

L_D(A) < L_S®h) + 2-R_m(H) + V(In2/9) / Cm)

e in modo data-dipendente:

L_D(h) < L_S@h) + 2 R_S(H) + 3: V(In(2/8) / @m))

Silvano Salvador

30

La prima disuguaglianza usa la complessità di Rademacher attesa; la seconda, più forte,
usa la complessità empirica sul campione effettivo e non richiede la conoscenza di
R_m(H).

Dimostrazione. Dimostriamo la prima variante; la seconda segue per argomenti analoghi
con il Lemma di McDiarmid.

Definiamo la funzione di eccesso:

(8) := sup_{h € H} (L.D@h)-L_S®)

Osserviamo che E[®(S)] 2 0 (l'eccesso atteso è non negativo). Vogliamo maggiorare
®(S) con alta probabilità. Per la disuguaglianza di McDiarmid, poiché cambiare un
singolo punto xi nel campione modifica L_S(h) di al più 1/m per qualsiasi h, la
funzione ® ha variazione limitata da 1/m in ciascuna coordinata. Dunque:

Pr] D(S) —- E[D(S] > t] £ exp(-2mt)

Il punto cruciale è maggiorare E[D(S)]. Sia S' = (x', ..., x') un campione fantasma
indipendente. Per la convessità dell'estremo superiore e la linearità dell'aspettazione:

ESS]

= E_S[sup_th} (LD) -LSM)]

F_S [ sup_{h} B_tS'} [L_{S'}h) — L_S@b) ]]
E_{S,S'} [ sup_{h} (L_tS'}(h) - L_S®))]

IA

dove l'ultimo passaggio usa il fatto che l'estremo superiore di un'aspettazione è
maggiorato dall'aspettazione dell'estremo superiore. Espandendo:

E_{S,S'} [ sup_{h} (L_{S'}(h) — L_S@))]
= E_{S,S'} [ sup_{h} (1/m) Xi (1_{h(x) # yi} — 1_{h(%) # yi}) ]

Poiché S e S' sono campioni i.i.d., ciascuna coppia (xi, x')) è simmetrica: possiamo
introdurre variabili di Rademacher o; che decidono, per ciascun indice i, se
"scambiare" o meno il contributo di x; e x'; senza alterare la distribuzione del vettore.
Formalmente, per qualsiasi funzione simmetrica G(z1,...,2) e variabili i.i.d. zi:

E z[G(z,,....:)] = E_{z,0} [G(0121,....0z)]

Applicando questo alla funzione differenza:

E_{S,S',0} [ sup_{h} (1/m) Xi oi © (1_{h(x5) # yi} — 1_{h(x) # yi})]
< 2: E_{S,0} [sup_{h} (1/m) Zio; 1_{h(x) # yi} ]

Silvano Salvador

31

dove il fattore 2 emerge dalla disuguaglianza triangolare applicata all'estremo
superiore. Poiché 1_{h(xj) # yi} = (1 — yi © h())/2 nel caso di etichette in {-1, +1}
(oppure direttamente con la funzione indicatrice in {0,1}), l'ultima espressione è
esattamente 2 + R_m(H) per definizione. Dunque:

E_S[®(8)] < 2: R_m(H)

Combinando con la disuguaglianza di McDiarmid (con t = Van(2/ 3)/(2m))):

Pr] ®(S) > 2-R_m(H) + VIn(2/9)/@m)] < 3/2 +8/2 = è

che è la disuguaglianza enunciata. Il

XXIX. Il legame tra Rademacher e dimensione VC: una catena di
disuguaglianze

La complessità di Rademacher e la dimensione VC non sono misure indipendenti:
esiste una catena di disuguaglianze che le collega, mostrando che la teoria di

Rademacher è strettamente più potente (nel senso di più fine) della teoria VC.
Proposizione. Se VCdim(H) = d < ©, allora:

R_m(H) < V(2-d-In(em/d) /m)

Dimostrazione. Per il Lemma di Sauer-Shelah, il numero di dicotomie distinte che H
produce su qualsiasi campione di m punti è al più t_H(m) £ (em/d)°d. La complessità
di Rademacher di una classe finita di funzioni {hi, ..., h_N} su un campione fisso è
maggiorabile tramite la disuguaglianza di Massart:

R_S({h1,....h_N}) £ max_kIh IS V@-InN) / m)

dove Ih IS = \(1/m) Xi h (x?) è la norma empirica. Per funzioni a valori in {0,1},
In I_S < 1. Condizionando su S e applicando il Lemma di Sauer-Shelah (il numero di
funzioni distinte su S è al più t_H(m)):

R_S(H) < V(2- In_H(m)) /m) < V(2-d -In(em/d) /m)

Prendendo l'aspettazione su S si ottiene la stessa stima per R_m(H). I
Confrontando con la disuguaglianza VC, che dava un margine di:

V((@d + In(em/d) + 2:In(2/8)) /m)

Silvano Salvador

32

si vede che i due maggioranti hanno la stessa dipendenza da d e m: la teoria di
Rademacher recupera la teoria VC come caso particolare. Il guadagno emerge quando
R_m(H) è stimabile direttamente dal campione e risulta molto più piccola del limite
VC: in questo caso, il teorema di Rademacher fornisce una garanzia di
generalizzazione molto più stringente di quella garantita dalla teoria VC uniforme.

La disuguaglianza di Massart. Poiché questo risultato è usato sopra, conviene
derivarlo esplicitamente. Sia F = {f1, ..., f_N} una classe finita di funzioni, e siano 01,
...,0 variabili di Rademacher indipendenti. Per qualsiasi vettore a E RR", la funzione t
P Eflexp(t : Xi o; aj)] = Il cosh(t * aj) £ exp(t° lal? / 2) (per l'ineguaglianza cosh(x) £
exp(x°/2)). Dunque:

E _o[exp(t:Zioa;)] = «pl? lal?:/2)

Applicando l'esponente alle N funzioni e usando la disuguaglianza di Jensen

sull'estremo superiore:

_0|max_kX;oi f (x) ]
(1/t) InE o[X_kexp(t:Zioif())]
(1/6) < In(N exp(t? - max_kIfILS?- m/2))
(in N) /t + t max_KIfIS?- m/2

IATA Hi

Ottimizzando su t > 0 (ponendo la derivata uguale a zero):

& = V(2-InN/(m- max_kIfI_S9))

e sostituendo:

E_o [ max_k X; ci f (x) ] £ max _kIFILS: V(2m- In N)

Dividendo per m:

R_S(P) < max_kIfI_S:V(2-InN/m) ll

XXX. Complessità di Rademacher e geometria degli spazi di
funzioni

Uno degli aspetti più eleganti della teoria di Rademacher è il suo legame con la
geometria degli spazi funzionali, in particolare con la teoria dei processi stocastici. La

Silvano Salvador

33

complessità di Rademacher è strettamente connessa alla nozione di larghezza di
Gaussian (Gaussian width) e alla teoria di Dudley sui processi Gaussiani.

Sia G = {g1, ..., g} un vettore di variabili gaussiane standard indipendenti. La

complessità Gaussiana di H rispetto a S è:

G_S(H) := E_g[sup_{h € H} (1/m) Zig h()]

Per il confronto di Sudakov-Fernique, la complessità Gaussiana e quella di

Rademacher sono legate da:

R_S(H) < V(a/2) : G_SH)

e la complessità Gaussiana è a sua volta maggiorabile tramite l'integrale di entropia di
Dudley:

G_SAH) < C- Jo {diam(H)/2} V(In N(H, I-1_S, e) ) de

dove N(H, l-I_S, e) è il wuzzero di copertura € di H rispetto alla norma empirica sul
campione S — il minimo numero di palle di raggio e nella norma l-IS necessarie a coprire

H— e diam(H) = sup{h,h'E H} lb — h'IS.

L'integrale di Dudley unifica il framework VC e il framework di Rademacher in una
prospettiva geometrica: la capacità di uno spazio di ipotesi è misurata dalla velocità
con cui il suo numero di copertura cresce al decrescere di e. Per classi VC di

dimensione d, si ha:

In N(H, l-1_S, e) < d - In(C/a)

per qualche costante C, e l'integrale di Dudley dà:

G_SAH) < C'-V(d /m): (log-factors)

compatibile con la stima VC. Per classi "lisce" — ad esempio, funzioni lipschitziane su
varietà a bassa dimensione intrinseca — il numero di copertura può crescere molto più
lentamente di O(1/e%d), e la complessità di Rademacher risulta corrispondentemente
inferiore al limite VC.

XXXI. Implicazioni epistemologiche della dipendenza dalla
distribuzione

Silvano Salvador

34

Il passaggio dalla teoria VC alla teoria di Rademacher ha un significato epistemologico
che va al di là del miglioramento tecnico. Nella teoria VC, la garanzia di
generalizzazione è 4 priori rispetto alla distribuzione: vale per qualsiasi mondo
compatibile con lo spazio di ipotesi scelto, indipendentemente da quale sia il mondo
effettivo. Nella teoria di Rademacher, la garanzia è 4 posteriori rispetto al campione:
tiene conto di come i dati effettivi si distribuiscono nello spazio delle istanze, e
fornisce garanzie più forti quando i dati rivelano una struttura più semplice di quella

prevista nel caso peggiore.

Questa dualità ha un riflesso diretto nel dibattito epistemologico tra prospettive
aprioristiche e aposterioristiche sulla giustificazione dell'induzione. La teoria VC
corrisponde a una forma di garanzia aprioristica: prima di vedere i dati, possiamo
affermare che qualsiasi agente con VCdim(H) = d ha bisogno di al più mo(e, è) esempi
pet generalizzare. La teoria di Rademacher corrisponde invece a una forma di garanzia
aposterioristica: dopo aver visto il campione S, possiamo stimare R_S(H) ec ottenere

una garanzia di generalizzazione adattata alla struttura specifica dei dati osservati.

In termini filosofici, questo suggerisce che la giustificazione dell'induzione ha due
livelli distinti. Il primo livello — corrispondente alla teoria VC — è trascendentale nel
senso kantiano: stabilisce le condizioni a priori che uno spazio di ipotesi deve
soddisfare perché l'induzione sia in linea di principio possibile, indipendentemente da
qualsiasi evidenza empirica. Il secondo livello — corrispondente alla teoria di
Rademacher — è empirico nel senso baconiano: si aggiorna in funzione dei dati
osservati, sfruttando la struttura del mondo reale per stringere le garanzie teoriche.
Una teoria completa dell'induzione deve incorporare entrambi i livelli: il primo
fornisce il quadro all'interno del quale l'induzione è pensabile; il secondo ne misura la

portata effettiva in funzione dell'evidenza disponibile.

XXXII. Rademacher e la vecchia questione della semplicità: oltre il
conteggio dei parametri

La teoria di Rademacher consente di chiarire un malinteso persistente nella filosofia
della scienza: l'identificazione della semplicità di una teoria con il numero dei suoi
parametri liberi. Questo criterio — usato in modo implicito da molte discussioni
popperiane e bayesiane sulla semplicità — è sistematicamente fuorviante nelle

situazioni in cui la geometria dello spazio parametrico non è uniforme.

Si consideri la classe HB = { 2: X— /-1,1] | Iblo < 1, Ihl_{Lip} < B } delle
funzioni 1-lipschitziane (con costante B) su X = [0,1]. Questa classe ha infiniti
parametri liberi (è uno spazio funzionale infinito-dimensionale) e quindi VCdim(H_B)
= 90, Eppure la sua complessità di Rademacher per campioni di dimensione m è:

R_m(H_B) <£ B/\m

Silvano Salvador

35

che decresce a zero all'aumentare di m e non dipende dalla dimensione dello spazio.
Uno spazio di ipotesi con infiniti parametri ma vincolato dalla condizione di Lipschitz
è PAC-apprendibile — non nel senso VC classico (che richiede VCdim finita), ma nel
senso esteso della convergenza uniforme garantita dalla teoria di Rademacher, purché
l'errore richiesto e sia tale che B/ Vm<<e.

Questo risultato smentisce la naive identificazione tra semplicità e numero di
parametri: la complessità epistemicamente rilevante non è il numero di gradi di libertà
del modello ma la capacità effettiva del modello di fittare il rumore — ovvero la
complessità di Rademacher. Un modello con molti parametri ma fortemente vincolato
(dalla lipschitzianità, dalla norma spettrale, dalla curvatura, o da altri vincoli

geometrici) può essere "semplice" nel senso rilevante per la generalizzazione.

Questa osservazione ha conseguenze dirette per la comprensione del successo delle
reti neurali profonde: la loro generalizzazione non è spiegata dal numero di parametri
(che è enorme) né dalla dimensione VC (che è altrettanto grande), ma da vincoli
impliciti. — imposti dalla struttura dell'architettura, dall'ottimizzatore, dalla
normalizzazione — che riducono la complessità di Rademacher effettiva a valoti
molto inferiori a quanto predetto dalla teoria VC. La "semplicità" che giustifica
epistemicamente queste reti è una semplicità geometrica nello spazio funzionale, non

una semplicità parametrica nel senso del conteggio delle variabili libere.

6. Descrizione minima, complessità algoritmica e
l'epistemologia bayesiana

XXXIII. Il principio di lunghezza di descrizione minima: Occam
come compressione

Esiste un secondo filone della teoria dell'apprendimento, parallelo e in parte
convergente con la teoria VC, che affronta il problema dell'induzione da una
prospettiva radicalmente diversa: non geometrica ma informazionale. Il principio di
lunghezza di descrizione minima (Minimum Description Length, MDI), sviluppato da
Jorma Rissanen a partire dalla fine degli anni Settanta, afferma che la migliore teoria
pet un insieme di dati è quella che comprime i dati nel modo più efficiente. In altri
termini: la migliore ipotesi è quella che fornisce la rappresentazione più concisa
dell'insieme di osservazioni.

Questa idea ha radici filosofiche che risalgono a Leibniz — il quale sosteneva che le
leggi della natura sono le ipotesi più semplici da cui si possono dedurre tutti i
fenomeni — e si connette, come si vedrà, con la complessità di Kolmogorov e con i
fondamenti logici dell'informazione. Ma la sua formalizzazione rigorosa richiede una
precisazione non banale: "comprimere" i dati significa co-descrivere sia il modello che
i residui, ovvero la lunghezza totale di descrizione è:

Silvano Salvador

36

Lh,S) = LO) + LE|h

dove L(h) è la lunghezza in bit della descrizione dell'ipotesi h, e L(S | h) è la lunghezza
in bit della descrizione del campione S una volta codificato tramite h. Il principio

MDL seleziona l'ipotesi che minimizza questa somma:

b*_MDL := argmin_{h € H} [LO)+L6 | b)]

La connessione con la teoria VC emerge quando si osserva che L(h) è una misura della
complessità dell'ipotesi — ovvero di quanto sia "semplice" nel senso della descrizione
algoritmica — mentre L(S | h) misura l'errore residuo (in un'analogia con i residui di
regressione). Minimizzare la loro somma equivale a bilanciare semplicità del modello e

adattamento ai dati, che è esattamente la struttura del dilemma bias-varianza.

Per rendere questo collegamento preciso, si consideri la codifica di Shannon: se le
ipotesi in H sono enumetate con probabilità a priori P(b), allora la lunghezza di
codifica ottimale è L(h) = — log? P(h). Minimizzare L(b) + L(S | h) diventa allora:

argmin_h [| -log? Ph) — log? P(S | b) ]
= argmax_h [ log: P(b) + log? P(S | b)]
= argmax h[Pb)-PS|b)]

= argmax_h P(h | S)

che è la massimizzazione della probabilità a posteriori MAP): il principio MDL, con questa
scelta di codifica, coincide con l'inferenza bayesiana con la distribuzione a priori Pb).
Il rasoio di Occam formalizzato come MDL è il principio bayesiano di preferire
ipotesi con alta probabilità a priori — e la "semplicità" di un'ipotesi è precisamente la

sua probabilità a priori espressa in bit di lunghezza di descrizione.

XXXIV. Complessità di Kolmogorov: la misura algoritmicamente
assoluta

La complessità di Kolmogorov (anche chiamata complessità algoritmica, o complessità
di Kolmogorov-Chaitin-Solomonoff) è la versione assoluta — indipendente dalla
scelta di distribuzione a priori — della nozione di lunghezza di descrizione. È stata
sviluppata indipendentemente da Kolmogorov (1965), Solomonoff (1964) e Chaitin
(1966), e costituisce una delle idee più profonde e destabilizzanti dell'informatica
teorica.

Definizione (Complessità di Kolmogorov). Fissata una macchina di Turing

universale U, la complessità di Kolmogorov di una stringa binaria x è:

K() := minf |p|:U@)=x}

Silvano Salvador

37

ovvero la lunghezza del programma più corto che, eseguito sulla macchina U, produce
x come output. Analogamente, la complessità condizionale è:

K(x | y) := mint |p|:U,y)=x}

la lunghezza del programma più corto che produce x conoscendo y come input
ausiliario.

La complessità di Kolmogorov ha proprietà fondamentali che la distinguono da

qualsiasi misura probabilistica:
La disuguaglianza di sottoadditività. K(x, y) £ K®& + K(y | x) +O(log K@®)

La simmetria dell'informazione. K(x, y) = K® + K(y | x) +O(og(K@& +K@) = Kw)
+ K(x | y) + Ollog(Kx,y)))

L'invarianza rispetto alla macchina: per qualsiasi coppia di macchine universali U1 e U?, le
complessità K_{U1} e K_{U:} differiscono al più di una costante additiva che
dipende dalla macchina ma non dalla stringa. Questo rende K(x) una proprietà
"quasi-assoluta" di x — assoluta a meno di costanti che dipendono dal sistema di
riferimento computazionale.

La incomputabilità: la funzione x > K(x) non è calcolabile da nessun algoritmo. Questo
non è un difetto tecnico ma un fatto fondamentale: se K fosse calcolabile, potremmo
risolvere il problema dell'arresto. La profondità della misura di Kolmogorov risiede
precisamente in questa incomputabilità — è una misura ideale, che indica una

direzione senza essere raggiungibile.

Per la filosofia dell'induzione, la complessità di Kolmogorov offre la risposta
algoritmicamente più soddisfacente alla questione della semplicità: la teoria più
semplice per un insieme di dati è quella che ha la complessità di Kolmogorov minima.
Questa è la "vera" semplicità, non dipendente da nessuna scelta convenzionale di
linguaggio o di codifica (a meno di costanti), e non dipendente da nessuna
distribuzione a priori.

XXXV, La distribuzione di Solomonoff e l'induzione universale

Il risultato più sorprendente che emerge dalla combinazione di complessità di

Kolmogorov e teoria della probabilità è la distribuzione di probabilità universale di

Solomonofi che rappresenta il tentativo più ambizioso — e filosoficamente più
provocatorio — di risolvere il problema dell'induzione in modo algoritmicamente
assoluto.

Silvano Salvador

38

La distribuzione di Solomonoff M è definita su stringhe binarie finite nel modo

seguente. Per ogni stringa x:

M(&) := £_{ p:U(p)inizia conx } 2X{-|p|}

ovvero M(x) è la somma delle probabilità (nella misura uniforme sui programmi
casuali) di tutti i programmi che iniziano producendo x. Questa è una misura
semiprobabilistica: non è normalizzata a 1, ma la sua somma converge (ed è al più 1
per la disuguaglianza di Kraft). La distribuzione normalizzata è chiamata distribuzione

universale a priori.

Il risultato fondamentale è il Teoreza di Solomonoff pet qualsiasi distribuzione di
probabilità computable p e qualsiasi stringa x di lunghezza n, si ha:

MG&) 2 um) 2°0-K(w}

dove K(u) è la complessità di Kolmogorov della macchina che calcola p. In parole: la
distribuzione universale di Solomonoff domina (a meno di una costante moltiplicativa)
qualsiasi distribuzione computabile. Ciò significa che chiunque usi M come
distribuzione a priori non può essere sistematicamente "battuto" nel lungo periodo da
nessun agente che usi una distribuzione computabile — M anticipa ogni regularità
computabile nei dati.

La conseguenza per l'induzione è radicale. Un agente che usa M come a priori e
aggiorna bayesianamente converge alla vera distribuzione dei dati (se questa è
computabile) in un senso molto preciso: il numero di errori di predizione accumulati è
limitato da K(u_true), la complessità della vera distribuzione. Questo è l'agente
induttivo universale di Solomonoff, ed è la risposta algoritmicamente più completa al
problema dell'induzione: impara qualsiasi sequenza generata da qualsiasi processo
computabile, con un numero di errori proporzionale alla complessità del processo.

Perché questo non risolve definitivamente il problema di Hume? Per due ragioni
fondamentali. Prima, M non è computabile: nessun agente fisico può implementarla. È
un ideale regolativo — indica la direzione ottimale senza essere raggiungibile. Seconda,
M presuppone la computabilità della vera distribuzione: se il mondo è governato da
processi non computabili (questione metafisica aperta), M non converge. Il
presupposto di uniformità della natura è ora rimpiazzato da un presupposto di
computabilità della natura — ugualmente ingiustificato sul piano trascendentale, ma
forse più plausibile empiricamente.

Il confronto con la teoria VC è illuminante. Entrambi gli approcci richiedono un
presupposto strutturale sul mondo — dimensione VC finita nella teoria PAC,
computabilità della distribuzione vera nella teoria di Solomonoff. Entrambi
convertono il problema dell'induzione in una stima del costo dell'apprendimento. Ma
mentre la teoria VC opera nel caso peggiore (worst-case) su tutte le distribuzioni,
Solomonoff opera in media (average-case) su tutte le distribuzioni computabili

Silvano Salvador

39

ponderate per la loro complessità. Le due teorie sono complementari: la prima è più
robusta (vale per qualsiasi distribuzione) ma meno informativa (fornisce garanzie
conservative); la seconda è più informativa (converge alla vera distribuzione) ma meno
robusta (richiede computabilità).

XXXVI. I limiti PAC-Bayes: coniugare frequentismo e
bayesianesimo

Un terzo approccio alle garanzie di generalizzazione, che si pone come sintesi tra la
prospettiva frequentista della teoria VC e quella bayesiana dell'MDL, è la teoria dei
limiti PAC-Bayes, sviluppata da McAllester (1998, 1999) e poi approfondita da Seeger,
Langford e molti altri.

L'idea di base è di non fissare una singola ipotesi h E H ma di considerare una
distribuzione Q su H — una "distribuzione a posteriori" sulle ipotesi — e di fornire
garanzie sull'errore reale atteso del classificatore randomizzato che campiona h da Q.
Sia P una distribuzione a priori su H fissata prima di vedere i dati. La divergenza KL
tra QeP è:

KL(Q1P) := E_{h-Q} [log(Q@) / Ph) ) ]

Il Teorema PAC-Bayes di McAllester afferma:

Teorema (McAllester, 1999). Per ogni distribuzione a priori P su H, per ogni è > 0,
con probabilità almeno 1 — è sul campione S di dimensione m, si ha simultaneamente

per ogni distribuzione a posteriori Q su H:

E_{h-Q} [L_D®)] < E_{h-Q} [L_S@h)] + V((KL(Q1 P) + In(m/8)) / (@m —
2))

Questa disuguaglianza è notevole per molteplici ragioni. Prima di tutto, è un limite
uniforme su tutte le distribuzioni a posteriori Q simultaneamente — il che significa che
si può scegliere Q in funzione dei dati osservati (Q è la distribuzione a posteriori
bayesiana, per esempio) senza invalidare il limite. In secondo luogo, il termine KL(Q ll
P) misura quanto la distribuzione a posteriori Q differisce dalla distribuzione a priori
P: se l'evidenza del campione ha fortemente spostato le credenze dell'agente rispetto
all'a priori, il margine di generalizzazione peggiora; se l'evidenza ha confermato l'a
priori, il margine è stretto. Il rasoio di Occam emerge qui nella forma più elegante: le
ipotesi che non si discostano molto dall'a priori (basso KL) hanno garanzie di

generalizzazione migliori.

Schizzo di dimostrazione. Il punto di partenza è il Lemma di Cambio di Misura (o lemma
di Donsker-Varadhan): per qualsiasi funzione misurabile f e distribuzioni Q, P:

Silvano Salvador

40

F_{h-Q} [ f@) ] S KL(QUP) + log E_{h-P} [ exp(f@)) ]

Applicando questo con f(h) = m * (L_Dh) — L_S(b)) e usando la disuguaglianza di
Hoeffding per maggiorare E_{h-P}[exp(m:(L_ Db) — L_Sb))] tramite un
argomento di "peeling" su tutti gli h e poi sulla distribuzione del campione, si ottiene il
termine VKL/m) con le costanti appropriate. Il dettaglio tecnico principale è la
gestione dell'uniformità in Q, che si ottiene applicando il lemma di cambio di misura
prima di prendere l'aspettazione sul campione, in modo da catturare l'eventuale
dipendenza di Q da S nel termine KL. Il

XXXVII. La struttura del limite PAC-Bayes come epistemologia
relazionale

Il limite PAC-Bayes introduce una struttura epistemologica che va al di là sia del
frequentismo puro della teoria VC sia del bayesianesimo soggettivo classico. La
garanzia di generalizzazione dipende da tre elementi: l'errore empirico di Q (quanto
bene la distribuzione a posteriori fitta i dati), la divergenza KL(Q | P) (quanto le
credenze a posteriori differiscono da quelle a priori), e la dimensione del campione.
Non dipende direttamente dalla dimensione VC di H — che può essere infinita, a
condizione che la massa di Q sia concentrata su ipotesi di bassa complessità rispetto a
jp

Questa struttura suggerisce un'epistemologia relazionale: la giustificazione di un'ipotesi
non dipende unicamente dalla sua complessità intrinseca (misurata dalla VC o dalla
lunghezza di Kolmogorov), né unicamente dalla sua adeguatezza empirica, ma dalla
sua distanza dall'insieme delle credenze a priori dell'agente. Un'ipotesi complessa ma
ben anticipata dall'a priori può essere più giustificata di un'ipotesi semplice ma
sorprendente. Questa è la struttura dell'inferenza bayesiana, ma ora con una garanzia
frequentista: non ci si limita ad affermare che Q è la distribuzione a posteriori corretta
dati i dati e l'a priori, ma si dimostra che il classificatore randomizzato che campiona
da Q ha un errore reale atteso vicino all'errore empirico, con una tolleranza
proporzionale a VKL/m).

C'è una tensione filosofica sottile che questa struttura introduce. La distribuzione a
priori P deve essere fissata prima di vedere i dati — questa è la condizione che rende il
limite non circolare. Ma la scelta di P riflette le credenze dell'agente sulla struttura del
mondo prima dell'evidenza, ovvero è essa stessa un atto epistemico che richiede
giustificazione. Il limite PAC-Bayes non dice come scegliere P; dice solo che, una volta
scelto P, le credenze a posteriori Q che non si discostano troppo da P hanno buone
garanzie di generalizzazione. Il regresso epistemico si sposta da "come giustificare le
generalizzazioni" a "come giustificare la scelta dell'a priori" — che è precisamente il
problema della giustificazione delle credenze a priori, il nodo più difficile dell'epistemologia
bayesiana.

Silvano Salvador

41

La teoria dell'apprendimento non risolve questo regresso — lo evidenzia con
precisione. La divergenza KL(Q | P) è il costo dell'aggiornamento; il costo non si
azzera mai, qualunque sia l'evidenza. In termini filosofici: ogni apprendimento ha un
costo rispetto alle credenze precedenti, e questo costo è ineliminabile. Non è possibile
imparare "da zero", senza un punto di partenza — il punto di partenza è sempre una
distribuzione a priori che codifica credenze pre-empiriche sulla struttura del mondo.

XXXVIII. Convergenza delle tre teorie: una mappa concettuale

È istruttivo, a questo punto, tracciare una mappa delle connessioni tra i tre approcci
esaminati — teoria VC, MDL/Kolmogorov, PAC-Bayes — e il problema
dell'induzione, rivelando la struttura unitaria che li sottende.

Tutti e tre gli approcci condividono la medesima architettura: affermano che la
generalizzazione affidabile richiede una misura di complessità dell'ipotesi, e che il
margine di generalizzazione cresce con la complessità e decresce con la dimensione del

campione. La differenza è nella misura di complessità adottata:

La teoria VC misura la complessità di H attraverso la dimensione VC, che è una
proprietà combinatorica dello spazio di ipotesi — il massimo numero di punti

frantumabili. E una misura nel caso peggiore, uniforme sulla distribuzione.

La teoria MDL misura la complessità di h attraverso la lunghezza di descrizione L(h)
— il numero di bit necessari a codificare h. Dipende dalla scelta di codifica ma è
indipendente dalla distribuzione dei dati.

La teoria PAC-Bayes misura la complessità di Q (distribuzione a posteriori) attraverso
KL(Q | P) — la quantità di informazione che il campione ha aggiunto rispetto all'a
priori. È una misura relazionale, dipende sia dall'a priori che dalla distribuzione dei
dati.

Le tre misure sono collegate da disuguaglianze fondamentali. Per qualsiasi h € H e
qualsiasi distribuzione a priori P su H:

—log: Ph) > K(h) — KP) —- O(1)

dove K(P) è la complessità di Kolmogorov della distribuzione P. Dunque la lunghezza
di descrizione sotto P è sempre almeno la complessità di Kolmogorov di h (a meno
della complessità di P e una costante). E per qualsiasi distribuzione a posteriori Q
concentrata su un'ipotesi singola h (distribuzione degenere in h):

KL(QIP) = —log Pb)

Silvano Salvador

42

Dunque il termine di penalizzazione nel limite PAC-Bayes con distribuzione degenere
ridiventa la lunghezza di descrizione under P _— e il limite PAC-Bayes si riduce a
qualcosa di molto vicino al limite MDL. La teoria VC emerge come caso limite
quando la distribuzione P è uniforme su uno spazio finito di cardinalità 2°d, nel qual
caso —log Ph) = d per ogni h, e il termine di penalizzazione diventa indipendente
dall'ipotesi specifica, recuperando il flavour worst-case della teoria VC.

Questa convergenza strutturale non è accidentale. Tutti e tre gli approcci sono, in
fondo, formalizzazioni della stessa intuizione: generalizzare affidabilmente significa
essere capaci di comprimere l'evidenza in modo più efficiente di quanto farebbe il
caso, e il grado di compressione raggiungibile misura sia la capacità del modello sia le
garanzie di generalizzazione. L'unità profonda di queste teorie è l'unità del problema
dell'induzione: qualsiasi agente che generalizzi dall'esperienza deve, in qualche senso
preciso e quantificabile, "scommettere" sulla struttura del mondo — e il costo di
quella scommessa, misurato in dati necessari, bit di descrizione, o divergenza dall'a

priori, è il prezzo dell'induzione.

XXXIX. Verso un'epistemologia della ‘compressionenale:
conclusione provvisoria

L'itinerario percorso nelle ultime tre sezioni converge su di una proposta
epistemologica che si può chiamare, provvisoriamente, epistezzologia della compressione: la
conoscenza empirica è compressione dell'esperienza, e la giustificazione di un'ipotesi è
misurata dalla sua efficienza compressiva rispetto all'evidenza disponibile e alle

credenze precedenti.

Questa proposta è, in un senso preciso, una naturalizzazione del rasoio di Occam.
Non si tratta di un principio metafisico sull'economia della natura, né di un'abitudine
cognitiva priva di fondamento razionale, né di una convenzione epistemica adottata
per. comodità: è una conseguenza dimostrabile delle condizioni formali
dell'apprendimento affidabile. La parsimonia è epistemicamente virtuosa perché, e
nella misura in cui, riduce il margine di generalizzazione — ovvero riduce la distanza
tra ciò che l'evidenza passata insegna e ciò che vale nel futuro. Questo è quantificabile,
e la teoria — VC, MDL, o PAC-Bayes, a seconda del livello di analisi — ne fornisce le
misure esplicite.

Il problema dell'induzione rimane insoluto nel suo aspetto trascendentale: nessun
teorema può giustificare incondizionatamente il salto dall'osservato all'inosservato .
Ma la teoria dell'apprendimento ha trasformato questo problema aperto in una
struttura articolata di condizioni necessarie e sufficienti, di misure di complessità
connesse tra loro, di garanzie quantitative che rivelano il costo preciso dell'induzione
affidabile. Questa trasformazione non è una dissoluzione del problema — è la forma
più onesta di progresso filosofico: non il silenzio del problema, ma la sua traduzione in

una lingua più precisa.

Silvano Salvador

43

7. La stabilità come virtù epistemica: un
fondamento alternativo della generalizzazione

XL. Dalla complessità dello spazio alla sensibilità dell'algoritmo

Tutti gli approcci esaminati fino a questo punto — teoria VC, complessità di
Rademacher, MDL, PAC-Bayes — misurano la generalizzazione attraverso proprietà
dello spazio di ipotesi H: la sua dimensione combinatorica, la sua capacità di fittare il
rumore, la lunghezza delle descrizioni che richiede. L'algoritmo di apprendimento
appare in questo quadro come uno sfondo relativamente passivo: si sceglie H, si
dimostra che ha dimensione VC finita, e si conclude che qualsiasi algoritmo ERM su
H generalizza bene. L'identità dell'algoritmo — il modo specifico in cui seleziona
un'ipotesi da H dati i dati — sembra irrilevante per le garanzie di generalizzazione.

Questa apparenza è fuorviante. Esiste una prospettiva completamente diversa,
sviluppata da Olivier Bousquet e André Elisseeff nel loro lavoro seminale del 2002,
che rovescia la gerarchia concettuale: invece di misurare la complessità dello spazio di
ipotesi, si misura la sensibilità dell'algoritmo alle perturbazioni del campione. La domanda
non è più "quanto è ricco H?" ma "quanto cambia l'ipotesi restituita dall'algoritmo se
si sostituisce un singolo punto del campione?". Questa sensibilità — denominata
stabilità algoritmica — si rivela sufficiente, e in certi contesti necessaria, pet garantire la
generalizzazione.

Il cambio di prospettiva ha conseguenze filosofiche non banali. Sposta il locus della
giustificazione epistemica dall'ontologia (le proprietà dello spazio di ipotesi, che
rappresenta la struttura del mondo che l'agente si aspetta di trovare) alla metodologia
(le proprietà del processo inferenziale, il modo in cui l'agente forma le proprie
credenze). Un'ipotesi è giustificata non in virtù della semplicità dello spazio da cui
proviene, ma in virtù della robustezza del processo che l'ha generata: un processo che
produce conclusioni molto diverse da dati quasi identici non merita fiducia epistemica,

indipendentemente dalla semplicità formale dell'ipotesi prodotta.

XLI. Definizioni di stabilità: una gerarchia di nozioni

La stabilità algoritmica non è una nozione unica ma una famiglia di nozioni, ordinate
per forza, ciascuna delle quali fornisce garanzie di generalizzazione con caratteristiche
diverse. Fissiamo il quadro: sia A un algoritmo di apprendimento deterministico, sia S
= (z1, ...,z) un campione con zi = (xi, yi). Denotiamo con A($) l'ipotesi restituita da A
su S e con S°{(i)} il campione ottenuto da S sostituendo zi con un nuovo punto z';
estratto indipendentemente dalla stessa distribuzione D. Analogamente, S\{i} indica il
campione di dimensione m-1 ottenuto rimuovendo il punto zi.

Silvano Salvador

44

Definizione (Stabilità uniforme). L'algoritmo A è uniformemente stabile con parametro
8 se per ogni campione S di dimensione m, per ogni punto z' E X X Y, e per ogni
indice i E {1,...,m}:

sup_{z E XXY} | L(A(S), 2) — L[AS{O}), 2) | S B
dove L(h, z) = 1_{h(&®) # y} è la funzione di perdita zero-uno (o più in generale
qualsiasi funzione di perdita limitata).

La stabilità uniforme dice che sostituire un singolo punto di addestramento può
cambiare il valore della funzione di perdita di A al più di f, un/forzzezzente su tutti i punti
di valutazione e su tutti i campioni. È una condizione forte — richiede che la
perturbazione di un singolo dato non produca mai un cambiamento drastico nelle

previsioni dell'algoritmo su nessun punto.

Definizione (Stabilità in ipotesi). A è stabile în ipotesi con parametro B se:
FE_STsup_{z} | UA), 2) — LAM»), 2) |] S B

dove l'aspettazione è sul campione S e z è il punto rimosso. Questa è una condizione

più debole: richiede che il cambiamento medio (aspettato) sia piccolo, non il

cambiamento uniforme sul caso peggiore.

Definizione (Stabilità in errore). A è stabile in errore con parametro f se:

| B.S[L_D(A(6)]- E_S[L_S(A(6)]]| £ B

Questa è la condizione più debole: richiede solo che l'errore reale atteso e l'errore

empirico atteso non differiscano troppo, senza alcuna garanzia di concentrazione
attorno alla media.

La gerarchia di implicazioni è:

Stabilità uniforme = Stabilità in ipotesi = Stabilità in errore

con implicazioni strette (nessuna delle frecce si inverte in generale).

XLII. Il teorema di generalizzazione tramite stabilità uniforme

Teorema (Bousquet-Elisseeff, 2002). Sia A un algoritmo uniformemente stabile
con parametro f rispetto a una funzione di perdita limitata da 1. Allora per ogni è > 0,
con probabilità almeno 1 — è sul campione S di dimensione m:

Silvano Salvador

45

L_D(AG) < L_S(A(S) + 28 + (4Am8 +1) V(In(1/8) /Cm))

Se B = O(1/m) — come accade per molti algoritmi regolarizzati — il margine di
generalizzazione è O(n(1/8)/m), indipendente dalla dimensione di H e dalla
dimensione VC.

Dimostrazione. Definiamo D(S) = L_D(A(S)) — L_S(A(S)). Vogliamo mostrare che © si
concentra attorno alla propria media e che la media è piccola.

Passo 1: la media è piccola. Si ha:

ESTESA) ]
= E_S[(1/m Xi {(AG,z)]
= E S[LAS,z)] (per simmetria degli indici)

E_S[L DA(S)]

= E S[E_{2-D}[{(A©,2)]]
= E_f{S,z2}[L(A(),2)]

Costruiamo il campione S' = (2', 22, ...,z) (S conil primo punto sostituito da z'). Per
la stabilità uniforme:

| {(A(S), 2) — (AG), 2) | SB

Dunque:

F_tS'} [C(A6I, 2)] + 8
E_{S'} [ (A(S), 21) ] + 8 (inominando z' come z1' in $')
ES[LS(AS)] + 8 (per simmetria)

F_{S2'} [ (AS), 2) ]
<

Il

Invertendo:

F_S[®(S)] = E_SIL D(AS)] — ES IL S(AG)] =

Si ottiene analogamente E_S[D(S)] 2 —B, dunque |E_S[D(S)]|] “ B — cioè in media
l'errore reale e l'empirico non differiscono più di 8.

Passo 2: concentrazione di ® via McDiarmid. Sia S°{(j}} il campione con zi
sostituito da z' Si ha:

D(S) — DS {0})
= IL _D(AG) — L_D(ASÒH O })] — IL_S(A(S) — L_STO}} (ASTON)

Silvano Salvador

46

Maggioriamo ciascun termine. Per la stabilità uniforme:

IL_D(A) — L_D(AS{0}) |
| F_{27DJ[{(A(S),2] — F_tz°D}[UAS0}),2]|
F_tz-D}[|UA)2) — UASÙHO})2)]] S B

IA_II

Per il termine empirico, la differenza L_S(A(S)) — L_{ST{®}}(A(ST{O})) coinvolge
sia il cambiamento del campione sia il cambiamento dell'ipotesi. Sviluppando
esplicitamente e usando la stabilità uniforme con perdita limitata da 1, si ottiene:

ILS(AS) — L_STIOASHO)] £ 28 + 1/m

Dunque la variazione di ® sostituendo il punto i-esimo è al più:

ID) —- DSUO})| £ 48 + 1/m

Per la disuguaglianza di McDiarmid applicata a ® con variazioni ci = 48 + 1/m:

Pr[ ®(8) — E[D(S)] > t] < exp(-2t / (m- 48 +1/m)))

Imponendo che questo sia £ è e risolvendo per t, e usando E[D(S)] £ B:

Pr] D(S) 2 28 + (Amb + 1)-V(n(1/9/Cm)) ] > |

XLIII. Stabilità e regolarizzazione di Tikhonov: il caso
emblematico

Il risultato precedente sarebbe di scarso interesse pratico se la stabilità uniforme con B
= O(1/m) fosse una proprietà rara e difficile da verificare. Al contrario, essa emerge in
modo naturale e quasi automatico negli algoritmi di apprendimento con regolarizzazione

di Tikhonov, che sono tra le procedure più usate nell'apprendimento supervisionato.

Sia H uno spazio di Hilbert di funzioni X + RR con prodotto scalare (+,‘) e norma Il.

L'algoritmo di regolarizzazione di Tikhonov con parametro 4 > 0 restituisce:

A_MS) := argmin_{h € H} [(1/m) Xi (h, 2) + x Ibl2]

Il termine Alhl? penalizza le ipotesi di grande norma, bilanciando l'aderenza ai dati con
la semplicità della soluzione. La norma IIhl misura la "complessità" dell'ipotesi nello
spazio di Hilbert — ad esempio, la norma del vettore pesi per un classificatore lineare,
o la norma della funzione in un RKHS (spazio di Hilbert a nucleo riproducente) per
kernel methods.

Silvano Salvador

47

Teorema (Stabilità di 'Tikhonov). Se la funzione di perdita l(:, z) è convessa e
IL-lipschitziana (| L(h,z) — L(b',z)| < L-Ih-h'I_X per qualche norma l-I_X su H), e se
-l => II_X (la norma di Hilbert domina la norma nella perdita), allora A_A è
uniformemente stabile con:

8B=1IL2/Am

Dimostrazione. Siano h = A_MS) e h' = A_MS°{@}). Per definizione di minimo,
entrambi soddisfano le condizioni di ottimalità del rispettivo problema regolarizzato.
Poiché l'obiettivo è fortemente convesso con parametro 2) (il termine Ihl? ha hessiana
241), si ha per qualsiasi coppia di minimi di problemi con obiettivi diversi:

In — h'I? < (1/4) - |KV [funzione dati], h — h')|

La differenza tra i due obiettivi (campione S vs campione S°{(i)}) riguarda solo il
termine relativo al punto i-esimo. Poiché la perdita è L-lipschitziana:

|(V [termine i-esimo], h — h')| < L-Ih-h'"l/m

Combinando:

2 <Ih- hl < (L/m) -Ih- N

dunque lh — h'l < L/(2Am), e per la lipschitzianità della perdita:

| t@,z)-L®,z)| < L-Ih-h'"l < L?2/(24m)

che è 8 = O(1/m) come affermato. Il

Il risultato rivela la struttura profonda della regolarizzazione: non è solo un trucco per
evitare l'overfitting (ridurre l'errore empirico sul campione di addestramento), ma
impone stabilità algoritmica nel senso tecnico, il che a sua volta garantisce la
generalizzazione attraverso il teorema di Bousquet-Elisseeff. La regolarizzazione è il
meccanismo che traduce la parsimonia strutturale (piccola norma dell'ipotesi) in

robustezza procedurale (piccola sensibilità alle perturbazioni del campione).

XLIV. Stabilità e la filosofia della robustezza epistemica

Il concetto di stabilità algoritmica offre una risposta alla domanda epistemologica sulla
robustezza della conoscenza: quando possiamo fidarci di una conclusione inferenziale?
La risposta della teoria della stabilità è: quando il processo che ha generato la

conclusione è tale che piccole variazioni nell'evidenza producono solo piccole

Silvano Salvador

48

variazioni nella conclusione. Un'ipotesi è cpistemicamente robusta se il processo che
l'ha prodotta è stabile; è fragile se piccole perturbazioni dell'evidenza avrebbero

portato a conclusioni radicalmente diverse.

Questa intuizione è profondamente pre-teorica — la ritroviamo in Peirce, che
considerava la fixation of beligf come il processo di consolidamento delle credenze
resistente alla perturbazione; in Rawls, che nelle sue riflessioni sull'equilibrio riflessivo
sottolineava l'importanza della coerenza delle intuizioni morali sotto variazioni dei casi
considerati; e nei criteri di sensitivity e safety proposti da Robert Nozick e Ernest Sosa

nell'epistemologia contemporanea. La condizione di sensitivity di Nozick — "se p
fosse falso, S non crederebbe che p" — e la condizione di safety di Sosa — "$ non
potrebbe facilmente credere che p se p fosse falso" — sono versioni qualitative della

nozione di stabilità: richiedono che le credenze dell'agente siano sensibili alle
perturbazioni controfattuali del mondo.

La teoria della stabilità algoritmica formalizza e quantifica queste intuizioni per il caso
dell'apprendimento da dati. La condizione B = O(1/m) è precisamente la versione
quantitativa della condizione di safety: l'algoritmo non potrebbe facilmente giungere a
conclusioni molto diverse se i dati fossero leggermente diversi. E il teorema di
Bousquet-Elisseeff mostra che questa condizione è sufficiente per la generalizzazione
— ovvero per la safety nel senso epistemologico: la conclusione dell'agente non
cambierebbe drasticamente se il campione di addestramento fosse diverso ma
provenisse dalla stessa distribuzione del mondo.

C'è tuttavia una differenza cruciale tra la condizione di safety epistemologica e la
stabilità algoritmica. La safety epistemologica riguarda la relazione tra la credenza e la
verità — richiede che la credenza avrebbe potuto essere falsa se il mondo fosse
diverso. La stabilità algoritmica riguarda la relazione tra la credenza e il campione —
richiede che la credenza non cambierebbe molto se il campione fosse leggermente
diverso. Queste sono condizioni distinte: un algoritmo può essere stabile ma
sistematicamente sbagliato (se il bias è tale da portare lontano dalla verità in modo
stabile), oppure instabile ma in media corretto. La stabilità da sola non garantisce la

correttezza; garantisce solo la coerenza procedurale del processo inferenziale.

XLV. Stabilità, compressione e il lemma di Littlestone-Warmuth

Esiste un ulteriore collegamento tra la teoria della stabilità e un approccio
completamente diverso alla generalizzazione, basato sulla nozione di schezzi di
compressione del campione (sample compression schemes), sviluppata da Littlestone e
Warmuth nel 1986 e approfondita da Floyd e Warmuth nel 1995.

L'idea è la seguente. Un algoritmo di apprendimento ha uno scherza di compressione di
dimensione k se esiste una procedura che, dato qualsiasi campione S di dimensione m,
seleziona un sottoinsieme S_k S S di al più k punti (il campione compresso) e una stringa
di bit ausiliaria b di lunghezza O(k), tale che l'ipotesi A(S) sia completamente
determinata da (S_k, b):

Silvano Salvador

49

AG) = AGk,b)

In parole: l'ipotesi appresa dall'algoritmo può essere ricostruita a partire da soli k
esempi del campione di addestramento (più qualche informazione ausiliaria). Questo è
un senso molto forte di "compressione": il resto del campione è irrilevante per
determinare l'ipotesi finale.

Teorema (Floyd-Warmuth, 1995). Se A ha uno schema di compressione di
dimensione k, allora con probabilità almeno 1 — è:

L_D(AGS)) < L_{S_k}(A(S_k, b)) + O(V(k- In(m) /m) + M(In(1/8) /m))

Notare che L_{S_k}(A(S_k, b)) è tipicamente zero (i k punti compressi sono di solito
classificati correttamente), e il margine di generalizzazione è Ok-In(m) /m) —

dipende solo dalla dimensione del campione compresso k, non dalla dimensione VC di
H.

La connessione con la stabilità è la seguente. Un algoritmo stabile con 8 = O(1/m)
può essere visto, intuitivamente, come un algoritmo che "si affida" a pochi punti critici
per determinare la propria ipotesi: se la sostituzione di un singolo punto quasi non
altera l'ipotesi, allora ciascun punto ha una piccola influenza individuale, e l'ipotesi
finale dipende in modo diffuso e "mediato" da tutto il campione. Questo è il regime
opposto a quello della compressione: la compressione richiede che pochi punti
determinino l'ipotesi; la stabilità richiede che nessun singolo punto abbia troppa

influenza.

Entrambe le proprietà garantiscono la generalizzazione, ma attraverso meccanismi
opposti: la compressione attraverso la semplicità della descrizione dell'ipotesi (pochi
supporti determinano tutto), la stabilità attraverso la robustezza del processo (ogni
singolo dato ha poca influenza). Questa dualità rispecchia la tensione tra approcci
"sparsi" (come gli SVM, che si fondano su pochi vettori di supporto) e approcci
"diffusi" (come la regressione di Tikhonov, in cui tutti i punti contribuiscono in misura
simile).

XLVI. L'algoritmo come agente epistemico: responsabilità e
trasparenza

La prospettiva della stabilità apre infine una questione di epistemologia normativa che
la teoria VC lasciava sullo sfondo: la responsabilità epistemica dell'algoritmo come agente
che forma credenze. Nella teoria VC, l'algoritmo è strumentale: conta solo che
selezioni un'ipotesi di basso errore empirico da uno spazio di dimensione VC finita, e
qualsiasi algoritmo ERM va bene. La teoria della stabilità invece discrimina tra
algoritmi: alcuni sono più stabili di altri, e la stabilità è una virtù epistemica autonoma

— non riducibile alla scelta dello spazio di ipotesi.

Silvano Salvador

50

Questo suggerisce che la valutazione epistemica di un sistema di apprendimento deve
comprendere non solo la domanda "da quale spazio di ipotesi viene selezionata la
credenza?" ma anche "con quale processo viene selezionata?". Due algoritmi che
minimizzano lo stesso errore empirico sullo stesso spazio di ipotesi possono avere
stabilità radicalmente diverse, e quindi garanzie di generalizzazione radicalmente
diverse. La scelta dell'algoritmo — non solo dello spazio — è un atto epistemicamente

responsabile.

In termini di filosofia dell'azione cognitiva, questo è un argomento a favore della
trasparenza procedurale: la giustificazione di una credenza empirica deve includere non
solo la descrizione delle ipotesi considerate e dell'evidenza osservata, ma anche la
descrizione del processo tramite cui si è giunti alla credenza, e in particolare la sua
stabilità rispetto a variazioni dell'evidenza. Un'ipotesi giunta per un processo instabile
— che avrebbe portato a conclusioni opposte con dati appena diversi — non è
epistemicamente giustificata anche se, per caso, corrisponde alla verità. Questa è la
versione formale della tesi epistemologica che la giustificazione è sensibile al processo,

non solo all'esito.

8. L'apprendimento sequenziale, il rimpianto e il
falsificazionismo come algoritmo

XLVII. Oltre la distribuzione fissa: il modello online

Il framework PAC esaminato nelle parti precedenti presuppone una distribuzione D
fissa e stazionaria da cui i punti vengono estratti ii.d. Questo presupposto, già
discusso nelle sue implicazioni filosofiche nella sezione XVII, è violato in molte
situazioni di interesse — dalle previsioni finanziarie all'apprendimento linguistico, dai
sistemi di raccomandazione alle scienze climatiche. In questi domini, la "distribuzione
del mondo" non è fissa: cambia nel tempo, talvolta in risposta alle azioni dell'agente

stesso, talvolta per effetto di dinamiche esterne imprevedibili.

Il modello dell'apprendizento online affronta questa difficoltà abbandonando del tutto il
presupposto probabilistico. Il gioco si svolge in una sequenza di round t = 1,2, ..., T:

al round t, il mondo presenta all'agente un'istanza x E X; l'agente emette una
predizione } E Y; il mondo rivela l'etichetta veray E Y; l'agente subisce una perdita
L@,y) e aggiorna la propria strategia.

Non si assumono distribuzioni, stazionarietà, né realizzabilità: la sequenza (x1, yi), ...,
(x, y) può essere scelta da un avversario arbitrario — anche adattivo, che conosce la
strategia dell'agente e cerca di massimizzarne l'errore. Il criterio di prestazione non è
più la minimizzazione dell'errore reale rispetto a una distribuzione, ma la

minimizzazione del rigpianto (regret) rispetto alla migliore ipotesi fissa in retrospettiva:

Silvano Salvador

51

Regret T(A, H) := X_{t=1}{T} @(A_t, z) — min_{h € H} X_{t=1}"{T}
LA&),y)

Il rimpianto misura quanto l'agente ha perso rispetto all'oracolo che avesse conosciuto
fin dall'inizio la migliore ipotesi fissa in H. Un algoritmo ha rimpianto sublineare se
Regret_T = o(T), ovvero se il rimpianto per round tende a zero: Regret_T / T + 0.
Questo significa che nel lungo periodo l'agente riesce ad avvicinarsi alle prestazioni
della migliore ipotesi fissa, anche senza sapere quale essa sia e anche in assenza di
qualsiasi assunzione probabilistica.

XLVIII. L'algoritmo di Halving e la dimensione di Littlestone

Il più semplice algoritmo online per la classificazione binaria è l'a/goritmo di Halving (o
Majority Vote), il quale mantiene un insieme di ipotesi "consistenti" — quelle che non
hanno ancora commesso errori — e predice la maggioranza di queste ipotesi su

ciascuna istanza.

Sia H finito e sia V_t l'insieme delle ipotesi consistenti fino al round t (quelle che non
hanno mai sbagliato). L'algoritmo predice:

Y = majority{h&):h E V_t}

Ogni volta che l'algoritmo sbaglia, significa che la maggioranza delle ipotesi consistenti
ha sbagliato, dunque almeno la metà di V_t risulta inconsistente e viene eliminata. Il
numero totale di errori è quindi al più log: |H|.

Questo risultato è eleganteissimo: il numero di errori (non il rimpianto, ma proprio gli
errori commessi) è logaritmico nella dimensione dello spazio di ipotesi. Ma per spazi
H infiniti, |H| = © e il risultato diventa inutile. La domanda analoga a quella della
teoria VC si ripropone: esiste una misura di complessità di H che governa il numero di
errori nel modello online, come la dimensione VC governa la complessità campionaria
nel modello offline?

La risposta è affermativa, e la misura corrispondente è la dizzensione di Littlestone,
introdotta da Nick Littlestone nel 1988. Essa è definita in termini di alberi di decisione
binari anziché di insiemi frantumati.

Definizione (Albero di shattering). Un 4/bero di shattering di profondità d per H è un
albero binario completo di profondità d i cui nodi interni sono etichettati con istanze
x; E Xei cui archi sinistri/destri corrispondono alle etichette 0 e 1, tale che per ogni
percorso radice-foglia — che definisce una sequenza di coppie (xi, bi) con bi E {0,1}
— esiste un'ipotesi h € H consistente con tutte le assegnazioni lungo il percorso:
h(xj) = bi per ogni i sul percorso.

Silvano Salvador

52

Definizione (Dimensione di Littlestone). La dizzensione di Littlestone di H, indicata
Ldim(H), è la profondità massima di un albero di shattering per H.

L'algoritmo Standard Optimal Algorithm (SOA) di Littlestone commette al più
Ldim(H) errori su qualsiasi sequenza avversariale — questo è l'analogo online del
teorema PAC:

Teorema (Littlestone, 1988). Esiste un algoritmo online che commette al più
Ldim(H) errori su qualsiasi sequenza di esempi generata da qualsiasi avversario, nel
modello realizzabile (esiste h* € H con zero errori sulla sequenza). Viceversa,
qualsiasi algoritmo online commette almeno Ldim(H) errori sulla migliore sequenza

avversariale.

La struttura del teorema è identica a quella del teorema PAC: una misura di
complessità combinatorica — questa volta la dimensione di Littlestone invece di quella
VC — è sia sufficiente che necessaria per l'apprendimento affidabile nel modello

online.

XLIX. La relazione tra dimensione VC e dimensione di Littlestone

Le due dimensioni combinatoriche sono legate da una catena di disuguaglianze. Ogni
albero di shattering di profondità d definisce un insieme frantumato (le foglie
dell'albero corrispondono a dicotomie su d istanze), dunque:

VCdim(H) < Ldim(H)

La disuguaglianza è in generale stretta. L'esempio canonico è la classe delle soglie sulla
retta reale (H = {h_0:0 € R}, h_0@ = 1_{x 2 0}), per cui VCdim(H) = 1 ma
Ldim(H) = 0%. Questo perché un albero di shattering per questa classe può avere
profondità arbitraria: al nodo radice si sceglie x1, al figlio sinistro (etichetta 0, dunque 0
> xi) si sceglie x: > x1, e così via. Anche se solo un punto alla volta viene frantumato
(VCdim = 1), la struttura sequenziale dell'albero consente di costruire percorsi di
profondità illimitata.

Questa differenza riflette una differenza strutturale fondamentale tra apprendimento
offline e apprendimento online. Nell'apprendimento offline, il campione è dato tutto
insieme e l'algoritmo può scegliere liberamente quale ipotesi restituire.
Nell'apprendimento online, l'algoritmo deve predire prima di vedere l'etichetta, e
l'avversario può scegliere le istanze in modo adattivo in funzione delle predizioni
passate dell'algoritmo. La struttura sequenziale dell'interazione consente all'avversario
di sfruttare la struttura di H in modo più efficace che nel caso offline, richiedendo una
misura di complessità più fine.

La dicotomia VCdim = 1, Ldim = © per le soglie ha una lettura filosofica diretta: nel
modello offline, le soglie sulla retta reale sono facili da apprendere (bastano O(1/e)

Silvano Salvador

53

esempi per generalizzare con precisione e); nel modello online, contro un avversario
adattivo, sono impossibili da apprendere senza errori illimitati. La "apprendibilità" non
è una proprietà assoluta di uno spazio di ipotesi, ma dipende in modo essenziale dal
modello di interazione — da se l'evidenza arriva tutta insieme o sequenzialmente, da

se l'ambiente è fisso o adattivo.

L. Rimpianto, Follow-the-Leader e la proiezione sul simplesso

Per classi di ipotesi con Ldim(H) = % (come le soglie), l'apprendimento senza errori è
impossibile nel modello online realizzabile. Ma nell'impostazione del rimpianto, è
comunque possibile ottenere rimpianto sublineare per molte classi, a condizione di
accettare una piccola perdita media per round. Il risultato fondamentale è che il
rimpianto sublineare — ovvero la capacità di avvicinarsi nel lungo periodo alle
prestazioni della migliore ipotesi fissa — è garantito da un'ampia famiglia di algoritmi,
indipendentemente dalla struttura di H, a condizione che lo spazio delle ipotesi sia

convesso o che si possa ricondurre a un problema convesso.

L'algoritmo più importante in questo contesto è Follow the Regularized Leader (FIRL),
che al round t seleziona:

$ = argmin_ fh € H} [X_{s=1}"{t-1} (&,z) + RO)]

dove R(h) è un termine di regolarizzazione (tipicamente R(h) = (1/n)Ihl?). Questo è
esattamente l'algoritmo di Tikhonov applicato iterativamente: a ogni passo, si

minimizza la perdita cumulata più la penalità di regolarizzazione.

Teorema (Rimpianto di FTRL). Se la funzione di perdita l(:, 2) è convessa e
I-lipschitziana per ogni z, e se R(h) = (1/2n)lIhl?, allora il rimpianto di FIRL dopo T

round è:

Regret_T < Ih*I? / (2) + n 2_ft=1}{T}IV£ ||?

dove h* = argmin_{hEH} 2 ((h,z) è la migliore ipotesi in retrospettiva e Vl è il
gradiente della perdita al round t. Scegliendo n = Ih*l / (VT), il rimpianto è O(L -
Ib*l - VI) = ONT), che è sublineare in T.

Dimostrazione. Denotiamo f (h) = L(h, 2) la perdita al round t e F (b) = ©_{s=1}7{t}
f (bh) + R(b) l'obiettivo cumulato fino al round t. Per definizione, h_ = argmin_{h}
Fb).

Il rimpianto si decompone come:

Regret_T = X [f(h)- f(b*]
= > [f)-f@)] + © FM) FO]

Silvano Salvador

54

Il secondo termine è telescopico: 2 [F (h +) — F(b*] — 2 [F-i(h) — F_i(b9] +
termini di regolarizzazione. Poiché h = argmin F 1, si ha F-:(h) “ F-1(h*), dunque
questo termine è non positivo eccetto per il contributo di R. Il risultato preciso si
ottiene applicando la condizione di ottimalità del primo ordine e la lipschitzianità della

perdita, dando:
Z [fh)-fha)] “n XIVf@)]?2 <a TL?

2 [fh4)— f(b9] £ R(b*) / n = Ib*I? / (2)

Sommando e ottimizzando su n si conclude. Il

LI. Il falsificazionismo come algoritmo online: Popper rivisitato

La struttura del modello online permette una riformulazione del falsificazionismo
popperiano che è, a mio avviso, una delle connessioni concettuali più fertili tra la
filosofia della scienza e la teoria dell'apprendimento. Popper sosteneva che la scienza
progredisce non per accumulazione di conferme ma per eliminazione di falsificazioni:
si propone un'ipotesi, si testa contro l'evidenza, la si abbandona se risulta falsificata, e
si propone un'ipotesi più restrittiva (meno falsificabile nel senso errato, più audace nel
senso corretto).

Questo schema corrisponde precisamente all'algoritmo di Halving nel modello online.
L'insieme V_t delle ipotesi consistenti è l'insieme delle teorie non ancora falsificate;
ogni errore (falsificazione) elimina almeno la metà delle teorie rimanenti; il numero
totale di falsificazioni prima che una singola teoria sopravviva è al più log:|H|. Il
metodo scientifico popperiano è, in questo senso, un algoritmo di apprendimento
online con rimpianto logaritmico.

Ma il parallelo si approfondisce se si considera la dinamica della corroborazione nella
filosofia di Popper. Popper distingueva tra la falsificazione — l'eliminazione definitiva
di un'ipotesi — e la corroborazione — il "credito" provvisorio accumulato da
un'ipotesi che ha resistito a tentativi di falsificazione. Nel framework online, la
corroborazione corrisponde alla sopravvivenza nell'insieme V_t: un'ipotesi ha alta
corroborazione se è sopravvissuta a molti round avversariali, ovvero se ha bassa
perdita cumulata su una lunga sequenza di test. Il valore epistemico della
corroborazione è precisamente il suo ruolo nell'abbassare il rimpianto futuro

dell'agente.

Il punto più interessante — e dove la formalizzazione supera l'intuizione popperiana
— è la questione delle ipotesi or falsificabili. Popper considerava la non falsificabilità
come un vizio epistemico: un'ipotesi che non può essere refutata non appartiene alla

scienza. Nel framework online, la non falsificabilità corrisponde a ipotesi con

Silvano Salvador

55

Ldim(H) = 0 nel modello realizzabile: classi che un avversario può sempre ingannare
senza che l'algoritmo le elimini definitivamente. Ma nel modello del rimpianto, queste
stesse classi ammettono rimpianto OD — quindi, anche senza poter eliminare le
ipotesi false, l'agente può avvicinarsi alle prestazioni della migliore ipotesi nel lungo
periodo. La falsificazione popperiana è necessaria per l'apprendimento senza errori; il
rimpianto sublineare è più debole e richiede solo la convessità del problema di perdita,
non la falsificabilità delle ipotesi.

Questa è una precisazione formale importante: Popper aveva torto nel considerare la
falsificabilità come il confine assoluto tra scienza e non-scienza. La linea più precisa è
tra classi con Ldim finita (apprendibili senza errori) e classi con Ldim infinita (non
apprendibili senza errori ma comunque trattabili con rimpianto sublineare). La
demarcazione epistemologica non è binaria ma stratificata, e la teoria online ne
quantifica la struttura con precisione.

LII. Il gioco con la natura: minimax e teoria delle decisioni

Il modello online si collega naturalmente alla teoria delle decisioni e alla teoria dei
giochi. L'agente che minimizza il rimpianto contro un avversario arbitrario sta, in

sostanza, risolvendo un problema minimax:

min_ {strategia di A} max_ {sequenza avversariale} Regret_T(A, H)

Questo è un gioco a somma zero tra l'agente e la natura (0 l'avversario). Il teorema
minimax di von Neumann afferma che, se lo spazio delle strategie è convesso e
compatto, esiste un equilibrio del gioco — una coppia di strategie miste ottimali per
entrambi i giocatori — e il valore dell'equilibrio è il rimpianto minimax.

Per il problema dell'apprendimento online con perdita convessa, il valore minimax del
rimpianto dopo T round è:

V_T = ®(X(T -log|H]))

pet H finito, e più in generale è determinato dalla dizensione sequenziale di Rademacher
dello spazio H, che è l'analogo online della complessità di Rademacher offline:

R_T"{seq}(H) := E_o|[sup_{h € H} (1/1) X_{t=1}7{T} o - h& (0) ]

dove la dipendenza di x da o rispecchia il fatto che nel modello online le istanze
possono dipendere dalla storia delle predizioni passate — il che è precisamente la

struttura adattiva dell'avversario.

Silvano Salvador

56

La dimensione sequenziale di Rademacher è il "giusto" analogo online della

complessità di Rademacher offline, nel senso che:

Regret_T(A, H) = @(T - R_T"{seq}(H))

per il migliore algoritmo online. Questo risultato (dovuto a Rakhlin, Sridharan e
Tewari, 2010) unifica la teoria offline e quella online in un quadro comune: entrambe
le complessità — di Rademacher offline e di Rademacher sequenziale — misurano la
capacità di H di correlarsi con il rumore, offline con campioni i.i.d. e online con

sequenze adattive.

LIII. Il rimpianto come metrica del progresso conoscitivo

Il concetto di rimpianto merita una riflessione filosofica finale che va al di là della sua
definizione tecnica. Il rimpianto Regret_T misura il divario tra le prestazioni
dell'agente in apprendimento sequenziale e le prestazioni dell'oracolo retrospettivo —
l'agente che avesse saputo fin dall'inizio la migliore ipotesi fissa. È una misura del costo
dell'ignoranza: quanto ha perso l'agente per non aver saputo, a priori, quale fosse la
regola corretta.

Questa nozione ha una struttura duale interessante rispetto alla complessità
campionaria PAC. La complessità campionaria misura quanti dati servono prima di
iniziare a fare previsioni affidabili — è un costo pagato interamente prima dell'uso
della conoscenza acquisita. Il rimpianto misura il costo distribuito nel tempo
dell'apprendimento — è una misura del prezzo pagato istante per istante durante
l'acquisizione della conoscenza. Un algoritmo con rimpianto ONT paga un prezzo
medio per round di O(1/ VD — 0: il costo dell'apprendimento decresce nel tempo, ma

non è mai zero nelle fasi iniziali.

In filosofia della scienza, questa struttura corrisponde a ciò che Peirce chiamava
autorettificazione del metodo scientifico: il metodo scientifico è un processo che, nel lungo
periodo, converge alla verità (o alla migliore approssimazione disponibile), anche se a
breve termine commette errori. Il rimpianto sublineare è la formalizzazione di questa
proprietà: l'algoritmo non converge a un'ipotesi corretta in un numero finito di passi
(non è necessariamente PAC-apprendibile), ma il suo errore medio per round

converge a quello dell'oracolo. Il progresso conoscitivo è asintotico, non finito.

Questa distinzione — tra convergenza in un numero finito di passi (PAC) e
convergenza asintotica (rimpianto) — corrisponde a una distinzione filosofica reale tra
due modelli del progresso scientifico. Il modello PAC è quello dei breaktbrough
scientifici dopo un numero sufficiente di osservazioni, la teoria corretta viene
identificata con alta probabilità e rimane stabile. Il modello del rimpianto è quello
dell'adattamento continuo: la scienza non converge a una teoria definitiva ma si avvicina
progressivamente alle prestazioni del miglior modello disponibile, aggiornandosi a
ogni nuova evidenza. Entrambi i modelli catturano aspetti reali della pratica scientifica;

Silvano Salvador

57

la teoria dell'apprendimento ne fornisce una precisa separazione formale e ne misura
le condizioni di applicabilità.

9. Il caso agnostico, i margini e il realismo
strutturale

LIV. Abbandonare la realizzabilità: il mondo non coopera

Tutto l'apparato sviluppato nelle parti precedenti — nella sua versione più semplice —
presuppone la rea/izzabilità: l'esistenza di un'ipotesi h* E H con errore reale L_D(h*)
= 0. Questa assunzione è matematicamente comoda ma epistemologicamente
problematica. Nella pratica scientifica e cognitiva, la realizzabilità è quasi sempre
violata: il mondo non è mai perfettamente descritto da nessuna delle ipotesi nello
spazio che l'agente considera. Le leggi fisiche sono approssimazioni; i modelli statistici
sono semplificazioni; le categorie cognitive sono proiezioni su una realtà che le eccede.
Come afferiva George Box con una formula divenuta proverbiale, tutti i modelli sono

sbagliati, ma alcuni sono utili.

Nel modello agnostico — detto anche non realizzabile — si abbandona l'assunzione che f
sia realizzabile in H ce si richiede solo che l'algoritmo si avvicini il più possibile alla
migliore ipotesi nello spazio:

L_D(AG)) < min_fh E H}L_D(h) + e

con probabilità almeno 1 — è. La quantità opt(H) = min_{h € H} L_D(b) è l'errore di
approssimazione o bias dello spazio: misura quanto bene la migliore ipotesi in H
approssima la funzione obiettivo vera (che può non essere in H). L'errore totale si

decompone nel modo già incontrato nella sezione XI:

L_D(AG) = [L_D(A8) — opi()] + opi(H)

= Errore di stima + Errore di approssimazione

Nel modello realizzabile, opt(H) = 0 e solo l'errore di stima è rilevante. Nel modello
agnostico, entrambi i termini sono non nulli, e la loro somma misura la distanza tra

l'ipotesi appresa e la migliore possibile nel mondo.

Teorema (PAC-apprendibilità ‘agnostica). Una classe di ipotesi H è
PAC-apprendibile nel modello agnostico se e solo se VCdim(H) < %. La complessità

campionaria nel modello agnostico è:

mof(e, è) = O((d/e?) - In(1/e) + (1/85) - In(1/3) )

Silvano Salvador

58

con d = VCdim(H). Confrontando con il modello realizzabile — dove la complessità
era O((d/e) : ln(1/) + (1/e) - In(1/3)) — si nota che la dipendenza da e è passata da
1/e a 1/82: il modello agnostico è più difficile del modello realizzabile, nel senso che
richiede un numero di esempi quadraticamente maggiore per la stessa precisione.
Questo non è sorprendente: in assenza di realizzabilità, l'agente non può sfruttare la
struttura "c'è un'ipotesi perfetta" per restringere la ricerca, e deve stimare l'errore con

meno leva statistica.

La dimostrazione nel caso agnostico è più sottile: non ci si può limitare a considerare
ipotesi con errore empirico nullo, e la convergenza uniforme deve essere dimostrata
pet tutte le ipotesi in H simultaneamente, incluse quelle con errore reale positivo. La
tecnica del campione fantasma si applica direttamente, e il risultato è la disuguaglianza

VC nella sua forma generale:
sup_{h € H} | L D@b) - L_S@h) | < V((2d - In(em/d) + 2:In(2/8)) /m)
con probabilità 1 — è, che vale sia nel modello realizzabile sia in quello agnostico. Nel

modello agnostico, la dipendenza da 1/62 emerge perché si deve soddisfare:

V((2d + In(em/d) + In(2/8)) /m) < e

risolvendo per m si ottiene m = O(d/e? - log(1/e)).

LV. La geometria dei margini: separazione e confidenza

Il passaggio dal modello realizzabile a quello agnostico apre la strada a una delle idee
più potenti dell'intera teoria dell'apprendimento: la nozione di margine di separazione e la
sua capacità di raffinare la dimensione VC in modo dipendente dalla geometria dei
dati.

Si consideri la classificazione binaria nello spazio X = R°n con classificatori lineari.
Un iperpiano separatore è definito da (w, b) con lwl = 1, e classifica il punto x come
sion(w-x + b). Il argine funzionale del classificatore (w, b) rispetto a un punto (x, y) E
XX {-1,+1} è:

Yu, b, x,y) := y(wx+b)

Il margine geometrico è il margine funzionale normalizzato rispetto alla norma di w:

yy, b, x,y) = y(wx + b) / Il

Il margine geometrico misura la distanza euclidea del punto x dall'iperpiano {x : w-x +

b = 0}. Per un campione S, il vzargine del campione è:

Silvano Salvador

59

y_S(w, b) := min_{i=1,...,m} y(w b, xi, yi)

Un classificatore che separa il campione con margine grande è, intuitivamente, più
"sicuro" delle proprie previsioni: ogni punto è lontano dalla frontiera di decisione, e
piccole perturbazioni dei dati non cambiano la classificazione. Questa intuizione
geometrica ha una precisa conseguenza per la generalizzazione, espressa dal seguente
risultato:

Teorema (Generalizzazione con margine, Bartlett 1998). Sia H la classe dei
classificatori lineari in Rn con Iwl < A e sia la distribuzione D supportata sulla palla
Ixl < R. Per ogni è > 0 e ogni y > 0, con probabilità almeno 1 — è su un campione S
di dimensione m, per ogni classificatore (w, b) che separa S con margine y_S(w, b) 2 y:

L_D(w b) < (4/Y9 : (R2A? / m) - (in(em/d) + 1) + V(In(2/8) / Cm))

dove d è la dimensione VC di H (pari a n + 1 per classificatori lineari in Rn). Ma —
e questo è il punto cruciale — il margine effettivo y entra nella stima come fattore
1/Y?, indipendentemente dalla dimensione dello spazio R°n. Se il margine è grande, la
garanzia di generalizzazione è buona anche se n è grandissimo, perché la classe effettiva di
classificatori con margine 2 y ha dimensione VC effettiva proporzionale a RA/Y, non

an.
Formalmente, si definisce la dizzensione di fat-shattering con soglia y come:

Definizione (fat-shattering). Uno spazio H di funzioni reali fa/-sbaffers un insieme C

= {x1,...,x } con scala y se esistono valori s1, ...,s € RR tali che per ogni etichettatura
b E {Y1}7k esiste h € H con:

h(xj) 2 sit y se bj = +1
h(xj) £ si — y se bj = —1

La dimensione di fat-shattering fat_y(H) è il massimo k tale che H fat-shatters qualche
insieme di k punti con scala y. Per i classificatori lineari con lwl £ A su punti con lxl £
R, si ha:

fat_y(H_lin) < LR?2A2 /y° J

che è indipendente dalla dimensione n dello spazio ed è tanto più piccolo quanto più grande è
il margine y. Questa è la chiave: in alta dimensione, con un margine sufficiente, la
classe effettiva di ipotesi che "contano" ha capacità molto inferiore a quella suggerita
dalla dimensione VC classica.

Silvano Salvador

60

LVI. Le Support Vector Machine come soluzione al problema
agnostico

Le Support Vector Machine (SVM), sviluppate da Vapnik, Cortes e altri a partire dagli
anni Novanta, sono l'incarnazione algoritmica di questa idea: trovare il classificatore
lineare che zzassizzizza il margine di separazione sul campione di addestramento, nella
convinzione che un margine grande corrisponda a una buona generalizzazione.

Formalmente, il problema duale delle SVM (nel caso separabile) è:
mario Lira gx)

soggetto a: ai > 0 perognii, Zioiyi = 0

dove « = (01, ..., a) E Rm sono i moltiplicatori di Lagrange. La soluzione ottimale
a* definisce il vettore pesi:

wf = Li gf yi xi

I punti con «* > 0 sono i veztori di supporto: sono i soli punti del campione che
determinano l'iperpiano di separazione ottimale. Tutti gli altri punti sono
epistemicamente irrilevanti — il classificatore ottimale è completamente determinato
da un piccolo sottoinsieme del campione.

Questa proprietà di sparsià ha un significato filosofico di prim'ordine. In primo luogo,
le SVM sono uno schezza di compressione del campione nel senso di Littlestone-Warmuth
(discusso nella sezione XLV): l'ipotesi finale è determinata dai soli vettori di supporto,
il cui numero è tipicamente molto inferiore a m. Questo garantisce una buona
generalizzazione attraverso il teorema di Floyd-Warmuth, con margine
ON(k-In(m) /m)) dove k è il numero di vettori di supporto. In secondo luogo, i vettori
di supporto sono esattamente i punti più "informativi" del campione — quelli che si
trovano sulla frontiera della regione di incertezza del classificatore. Sono gli esperizzenti
cruciali del campione, per usare un'espressione baconiana: i dati che, se modificati,
cambierebbero la conclusione. Tutti gli altri dati sono ridondanti rispetto all'ipotesi
appresa.

Nel caso non separabile — che è il caso agnostico per eccellenza — le SVM risolvono
il problema primale con variabili di scarto (s/ack variables) & 0:

min_{w, b, &} (1/2)liwl? + C - Li &

soggetto a: yi(w-xi +b) 21-%, &20 perognii

Il parametro C > 0 governa il compromesso tra la massimizzazione del margine
(piccola Iwl) e la minimizzazione degli errori sul campione (piccola 2%). Per C grande,

Silvano Salvador

61

si penalizza fortemente gli errori sul campione — il bias è basso, la varianza è alta. Per
C piccolo, si accetta qualche errore sul campione pur di mantenere un margine grande
— il bias è alto, la varianza è bassa. Questo è il dilemma bias-varianza nella sua forma

geometrica più trasparente.

LVII. Il trucco del nucleo: induzione in spazi infinito-dimensionali

L'aspetto più matematicamente straordinario delle SVM è il kerze/ trick, che permette
di applicare la classificazione lineare in spazi di caratteristiche di dimensione arbitraria

— anche infinita — senza mai calcolare esplicitamente le coordinate in quello spazio.

Un nucleo (kernel) è una funzione k : X X X — R tale che esiste uno spazio di Hilbert
F e una funzione di caratteristiche p : X + F con:

k(x,x) = (90), 0(x))_F

Il teorema di Mercer afferma che k è un nucleo valido se e solo se è simmetrico e

semidefinito positivo: per qualsiasi insieme finito di punti {x1,...x}, la matrice di

Gram K con KU = (xi, x) è semidefinita positiva.

Il problema duale delle SVM dipende dai dati solo attraverso i prodotti scalati (x;, x ).
Sostituendo questi con k(xi, x ), si ottiene la SVM in uno spazio di caratteristiche F che
può essere infinito-dimensionale — ad esempio, lo spazio di tutte le funzioni
polinomiali (nucleo polinomiale), lo spazio di tutte le funzioni gaussiane (nucleo RBF),
o lo spazio di tutte le funzioni continue (nucleo di Mercer generico). Il classificatore in

questo spazio è:

f(x) = sign(Xi a yi kx, x) + b*)

che è una combinazione lineare (nello spazio di caratteristiche) dei k-prodotti con i
vettori di supporto. La funzione f non è lineare in X — può avere frontiere di
decisione arbitrariamente complesse nello spazio originale — ma è lineare nello spazio

di caratteristiche F, che è dove vivono le garanzie teoriche.

La garanzia di generalizzazione per le SVM con nucleo nel caso agnostico prende la
forma:

L_D(f) < L _S(f) + O( x R_g? Az / (mM) + In(1/3)/m))

dove R_y = sup_{x € X} V&k(,») è il raggio nello spazio di caratteristiche, A = lwl_F
è la norma nell'iperpiano di separazione, e y è il margine. Crucialmente, questa stima

non dipende dalla dimensione dello spazio di caratteristiche F — che può essere

Silvano Salvador

62

infinita — ma solo dal margine e dalla norma del classificatore. Il nucleo "comprime"

la complessità dello spazio infinito-dimensionale in due scalari: il raggio e la norma.

LVIII. Verosimiglianza strutturale e realismo: il margine come
misura dell'accordo

Il concetto di margine ha una risonanza diretta con il dibattito contemporaneo sul
realismo scientifico e sulla verosimiglianza — la proprietà di avvicinarsi alla verità.
Popper aveva proposto la nozione di verisizzilitude (Wahrheitsàhnlichkeit) come misura
del progresso scientifico: le teorie più tardi nella storia della scienza non sono
semplicemente più confermate delle precedenti, ma più vicine alla verità — più
verosimili. Ma la formalizzazione di questa intuizione si rivelò filosoficamente
disastrosa: il teorema di Tichy e Miller (1974) mostrò che per teorie false, la
verosimiglianza popperiana non è ordinabile nel senso desiderato — una teoria falsa
può avere verosimiglianza inferiore alla congiunzione di se stessa con una teoria
arbitraria.

Il concetto di margine offre una via alternativa per pensare la verosimiglianza nel
contesto dell'apprendimento statistico. La frontiera di decisione di un classificatore
con margine y separa le classi con una "zona di sicurezza" di ampiezza y: tutti i punti si
trovano a distanza almeno y dalla frontiera. Questo margine misura non solo la
correttezza della classificazione (il classificatore separa correttamente il campione), ma
la confidenza strutturale nella separazione — quanto robusta è la classificazione rispetto a
perturbazioni dei dati.

Un classificatore con margine piccolo è "verosimilmente" meno affidabile di uno con
margine grande, nel senso che le sue previsioni su dati nuovi sono meno stabili.
Questa è una misura della verosimiglianza che evita i paradossi di Tichy-Miller: non si

confrontano teorie false tra loro in astratto, ma si misura la distanza della frontiera di

decisione dai dati osservati — una quantità geometricamente e statisticamente ben
definita.
In questa prospettiva, il programma del realismo strutturale di John Worrall — che

propone di essere realisti non sugli oggetti descritti dalle teorie ma sulla struttura
matematica che esse catturano — riceve un supporto formale dalla teoria dei margini.
La struttura rilevante di un classificatore non è il suo valore su ogni singolo punto (che
è soggetto a overfitting), ma la geometria della sua frontiera di decisione rispetto alla
distribuzione dei dati — il margine, appunto. Un classificatore con grande margine ha
trovato una struttura robusta nei dati; questa struttura è ciò che merita l'aggettivo

"reale" in senso worralliano — non le etichette singole, ma la frontiera che le separa.

LIX. Il caso agnostico come modello dell'approssimazione
scientifica

Silvano Salvador

63

Il modello agnostico dell'apprendimento ha un'ultima implicazione filosofica che vale
la pena articolare esplicitamente: è il modello appropriato per pensare non solo
all'apprendimento statistico, ma all'intera pratica della modellazione scientifica intesa

come approssimazione.

Nella filosofia della scienza standard, una teoria è valutata in termini di verità o falsità:
o descrive correttamente il mondo o no. Il modello realizzabile dell'apprendimento
corrisponde a questa visione: o la funzione obiettivo vera è nel modello (realizzabilità)
oppure no. Ma questa è una visione ingenuamente bivalente della conoscenza
scientifica. La meccanica newtoniana è falsa — non descrive correttamente i fenomeni
relativistici e quantistici — ma è una buona approssimazione in un vasto dominio di
applicazione. La termodinamica classica è falsa — le leggi dei gas ideali non valgono

esattamente per nessun gas reale — ma è utile e epistemicamente preziosa.

Il modello agnostico cattura questa struttura: opt(H) > 0 significa che nessun modello
nel nostro spazio è perfetto, ma il modello migliore nello spazio ha errore opt(H) —
la migliore approssimazione raggiungibile. L'obiettivo dell'apprendimento agnostico è
avvicinarsi a opt(H) con ettore e addizionale, il che corrisponde a trovare il modello
più accurato all'interno della nostra classe di modelli disponibili.

Questa prospettiva si connette con la tradizione epistemologica dell'empiriszzo costruttivo
di Bas van Fraassen: le teorie scientifiche non devono essere vere per essere
epistemicamente accettabili, devono essere empiricamente adeguate — devono salvare i
fenomeni nel dominio di applicazione considerato. L'adeguatezza empirica
corrisponde, nel linguaggio dell'apprendimento agnostico, a un basso errore reale
rispetto alla distribuzione di fenomeni rilevanti — non zero (che richiederebbe
realizzabilità), ma sufficientemente piccolo rispetto a opt(H). La scienza è un processo
di minimizzazione agnostica: non cerca la verità assoluta, ma la migliore
approssimazione all'interno dello spazio di modelli disponibili, con garanzie che tale

approssimazione generalizza a fenomeni non ancora osservati.
Il margine di generalizzazione agnostico:

L DAh_S) £ opt(H) + e conprobabilità 2 1 — è

purché m 2 O(d/e? * log 1/e : log 1/8), è la formula che quantifica le condizioni
dell'empirismo costruttivo: la garanzia che la migliore approssimazione trovata sui dati
(L_S(h_S) ® opt(H)) generalizza ai fenomeni futuri (L_D(h_S) = opt(H) + s), con
probabilità esplicita e dipendenza esplicita dal numero di osservazioni necessarie.
Questa non è filosofia della scienza speculativa — è la struttura matematica
dell'approssimazione scientifica affidabile.

Silvano Salvador

64

Appendice A — Disuguaglianze di concentrazione
della misura

Le garanzie di generalizzazione discusse lungo tutta la monografia poggiano su un
nucleo comune di risultati probabilistici noti collettivamente come disuguaglianze di
concentrazione. Il loro significato comune è il seguente: una funzione che dipende da
molte variabili indipendenti ma non è troppo sensibile a nessuna singola variabile si
concentra attorno alla propria media con probabilità esponenzialmente alta. Questa
sezione raccoglie, in forma autocontenuta e con derivazioni complete, le
disuguaglianze usate nella monografia.

A.1 La disuguaglianza di Hoeffding

Lemma (Hoeffding). Sia Z una variabile aleatoria con a £ Z £ b p.s. e E[Z] = p.
Allora per ogni t > 0:

F[expA(Z = p)] £ exp(X°b7a)? / 8)
Dimostrazione. Poiché Z € [a, b], possiamo scrivere Z = ab + (1-a)a dove a =
(Z-a)/(b-a) € [0,1]. Per la convessità dell'esponenziale:

exp(AZ) = exp(Mab + (1-a)a)) = a'exp(Ab) + (1-0) exp)

Prendendo l'aspettazione e usando E[Z] = p, dunque E[a] = (u-2)/(b-a) =: p:

FlexpZ)] £ prexp(b) + (1-p)'exp(Aa) = exp(Aa) : [p:expA(b=a)) + (1-p)]

Sia h = XMb—a) e sia g(h) = In[p'exp(h) + (1-p)] — ph. Allora:

FlexpA(Z=p)] S exp(Xa — A) + exp(&(h)) + exp(ph)
= exp(g(b)

Sviluppando g in serie di Taylor attorno a h = 0: g(0) = 0, g'(0) = 0, e g"M) = p(1-p)
< 1/4 (per il massimo di p(1-p) in [0,1]). Dunque g(h) £ h?/8 = X2b—-a)?/8, che è la
disuguaglianza di Hocffding, Il

Teorema (Hoeffding per somme). Siano Zi, ..., Z variabili aleatorie indipendenti
con Zi € [ai, bi]. Allora:

Pr| Xi(Zi — E[Zi]) 2 t] £ exp(-2t /Zi(bi — a)?)

Silvano Salvador

65

Dimostrazione. Per la disuguaglianza di Markov applicata alla funzione esponenziale:
Pr[Li(ZiEl[Zi])) 2 t] = Pr[expOXi(ZiE[Zi])) 2 expQ6]

exp(M)  ElexpAXi(Z-EIZi))]

exp(-At) : Ii E[exp(X(Zi-E[Zi]))] (indipendenza)
exp(-At)  exp(/8 - Zi(biaj))

IA HA IV

Ottimizzando su x: X* = 4t / Xi(bimaj)?, e sostituendo:

Pr[...] < exp(-2t / Zibi-a)) Il

Applicazione alla teoria PAC: con Zi = 1_{hG)#y} € [0,1] e t = em, la
disuguaglianza di Hoeffding dà:

Pr[L Sh) - LD) 28] £ exp(-2me?)

pet una singola ipotesi h fissata. Per ottenere una garanzia uniforme su tutto H, si usa

il Lemma di Sauer-Shelah per limitare il numero di dicotomie distinte e si applica un

union bound.

A.2 La disuguaglianza di McDiarmid

Teorema (McDiarmid, 1989). Sia f : Z°m — R una funzione tale che per ogni i €
{Lamp peroni zia Za

fanzine iz) | Sua

(condizione di variazione limitata o condizione di differenza limitata). Allora se Zi, ...,Z sono
variabili aleatorie indipendenti:

Pr[ f(Z1,...,Z) — E[(Za,...Z)]2t] £ exp(-2t2 / Zio?)

Dimostrazione. Si usa il metodo delle differenze martingale. Definiamo la sequenza:

M := E[f(Z1,...Z)|Z...Z] per j=0,1,...,m

con Mo = E[f] e M = f(Z1,...,Z). La sequenza (Mo, Mi, ..., M) è una martingala
rispetto alla filtrazione F_j = o(Z1,...,Z). Le differenze martingala D = M —- M_i
soddisfano E[D | F_{j-1}] = 0.

Per la condizione di variazione limitata, il range condizionale di M_dato F_{j-1} è:

Silvano Salvador

66

sup_{z } M (2) — inf {z}M(z) £ c

(poiché cambiare solo il j-esimo argomento cambia f di al più c, e M è la media
condizionale di f). Dunque D € [L, U] con U — L £ c condizionatamente a
F_{j-1}. Applicando il Lemma di Hoeffding condizionatamente a F_{j-1}:

E[expAD) | F_{j-1}] < expQc 2/8)

Dunque:

E[ exp(X(f — E[f))] = E[expAX D)] £ II exp(2c?/8) = exp(22c 2/8)

e per la disuguaglianza di Markov:

Pr[f- E[fl =t] £ exp(-At)  expQ?Xc 2/8)

Ottimizzando su A: X* = 4t/Xc 2, da cui la disuguaglianza enunciata. Il

La disuguaglianza di McDiarmid generalizza quella di Hoeffding: quest'ultima è il caso
in cui f = X; Zi/m (con ci = (bi-aj)/m per variabili in [ai,bi]).

A.3 La disuguaglianza di Bernstein

La disuguaglianza di Hoeffding usa solo il range delle variabili; quando la varianza è
piccola, la disuguaglianza di Bernstein fornisce limitazioni più strette.

Teorema (Bernstein). Siano Z1,...,Z variabili aleatorie indipendenti con E[Zi] = 0,
Var[Zi] £ 0° e | Zi| £ M p.s. Allora:

Pr| (1/m)Zi Zi>t] < exp(-—mt? / (20? + 2Mt/3))

Dimostrazione. Per variabili centrate limitate, il momento esponenziale soddisfa:
E[expAZj)] < exp(E[Zi] : (e€ {AM} — AM — 1) / M?) £ exp(0?- g(M) / M?)
dove g(u) = e‘ — u — 1 2 0. Usando la disuguaglianza g(u) < u?/(2(1 — u/3)) per u < 3
(derivabile dalla serie di Taylor dell'esponenziale con stima dei resti), si ottiene:

In EfexpAZj)] £ X°0° / (2(1 — AM/3)

Silvano Salvador

67

Sommando sulle variabili indipendenti e applicando la disuguaglianza di Markov:

Pr[]Li Zi/m > {] < exp(-Amt) + exp(mA202 / (2(1-AM/3)))

Ottimizzando su ) E (0, 3/M): X* = t / (02 + Mt/3), si ottiene la disuguaglianza
enunciata. Il

Nella teoria dell'apprendimento, Bernstein è rilevante quando l'errore di
approssimazione opt(H) è piccolo: in questo caso, la varianza dell'errore è
proporzionale a opt(H) (poiché Var[1_{h@#y}] £ E[l_{h©#y}] = LDOb) =
opt(H)), e la complessità campionaria scende da 0(1/8°) nel caso agnostico generico a
O(1/e) quando opt(H) = 0 (caso realizzabile) — recuperando esattamente la
differenza di scala già osservata nel testo.

A.4 L'identità di Pascal e la stima asintotica delle somme
binomiali

Il Lemma di Sauer-Shelah usa l'identità di Pascal e una stima delle somme binomiali

che qui deriviamo esplicitamente.

Identità di Pascal: C(n, k) + C(a, k-1) = C(n+t1, k). Dizzostrazione: C(a,k) + C(a,k-1)
= n!/(Ki(n-k)) + n!/(k-1)I(a-k+1)!) = nl(n+1) / (£!(n+1-k)) = C(n+1,k).1

Stima della somma binomiale: Per m 2 d 2 1:
X_{k=0}"{d} C(m, k) £ (em/d)°d
Dimostrazione. Sia p = d/m € (0,1]. Consideriamo la variabile aleatotia binomiale
Bin(m, p). Si ha:
Pr[Bin(m,p) £ d] = X_{k=0}7{d} C(m,k) pk (1-p)° {mk}
Poiché p°k (1-p)°©{m-k} 2 p°d (1-p)°{m-d} per k £ d (il termine è massimo in k
= d quando p = d/m per la simmetria della binomiale), si ha:
©_{k=0}"{d} C@mJ) < 2_{k=0}£{d} C(mjJ) / p"d (1-p)*{m-d})
S 1/ (pd (1-p){m-d})
Con p = d/m: p°d = (d/m)d e (1-p)°{m-d} = (1-d/m){m-d} 2 e*{-d}
(poiché (1-x)*{1/x} 2 e%{-1} per x € (0,1)). Dunque:

X_{k=0}{d} C(mjk) < (m/d)d - ed = (em/d)Yd I

Silvano Salvador

68

A.5 Una nota sulla misurabilità

Nella formulazione rigorosa della teoria PAC, occorre verificare che le quantità come
sup_{h EH} (L_D(h) — L_S(b)) siano variabili aleatotie misurabili, ovvero che
l'estremo superiore su una classe di funzioni (potenzialmente non numerabile) sia
misurabile. Questo è un punto tecnico che viene spesso ignorato nelle presentazioni
informali ma che ha rilevanza quando H è non numerabile.

La condizione sufficiente standard è che H sia una dasse di Suslin: cioè l'immagine di
uno spazio polacco (completo e separabile) mediante una funzione misurabile. Classi
di funzioni continue su spazi compatti, classi definite da condizioni di lipschitzianità o
vincoli su norme, e classi parametrizzate da vettori in Rd soddisfano tipicamente
questa condizione. Per tali classi, sup_{hE H} f(b) è una variabile aleatoria misurabile
per qualsiasi funzione misurabile f.

Un'alternativa più elementare, sufficiente per tutte le applicazioni della monografia, è
richiedere che H sia separabile rispetto alla topologia della convergenza puntuale:
esiste un sottoinsieme numerabile denso Ho C H tale che per ognih € Heognie >
0 esiste ho E Ho con sup_{x E X} [h(®) — ho(&)| £ e. In questo caso, sup_{h E H}
f(h) = sup_{hoE Ho} f(ho) è un estremo superiore numerabile di funzioni misurabili,
dunque misurabile. Tutte le classi considerate nella monografia soddisfano questa
condizione di separabilità.

Appendice B — Il lemma di cambio di misura di
Donsker-Varadhan

Il Lemma di Donsker-Varadhan, usato nella derivazione del limite PAC-Bayes nella
sezione XXXVI, merita una trattazione autonoma per la sua importanza trasversale

nella teoria dell'informazione e nell'apprendimento statistico.

Lemma (Donsker-Varadhan, variante variadazionale della divergenza KL). Per
qualsiasi coppia di distribuzioni P, Q su uno spazio misurabile S2 e per qualsiasi
funzione misurabile f:Q2 + R con E_P[exp(f)] < 0:

B_Qffl £ KL(QIP) + In E_P[exp(f]
con uguaglianza raggiunta quando dQ/dP © exp(fî) — ovvero quando Q è la
distribuzione di Gibbs/tilted associata a f.

Dimostrazione. Pet la disuguaglianza di Jensen applicata alla convessità dell'entropia
relativa, possiamo scrivere:

Silvano Salvador

69

KL(Q1P) = E_Q[In(d4Q/dP)]
[ In(dQ/dP) + f — f]
[f] + B_Q[In(dQ / (P-e°£/Z))]

EQ
= E_Q

dove Z = E_P[e?f] è la costante di normalizzazione. L'ultimo termine è KL(Q ll P)
con P = P-ef/Z, che è non negativa (essendo una divergenza KI):

IV

KL(QINP) > 0

Dunque:

KL(QIP) = E_Q[f-InZ+KL(QIÒ) > E_Q[f] - In E_P[e?f

Riorganizzando:

E_Q[f] < KL(QIP) + InE_P[ef] Il

L'uguaglianza si raggiunge quando KL(Q | P) = 0, ovvero quando Q = P = P-e%f/Z.
Questa è la distribuzione di Gibbs: concentra la misura sulle regioni di alto valore di f,
in modo proporzionale all'esponenziale di f pesato da P.

Il Lemma ha una lettura informazionale diretta: il guadagno atteso (sotto Q) di una
funzione f non può superare la somma dell'informazione che Q porta rispetto a P
(misurata dal KL) e del log-momento esponenziale di f sotto P. Ogni guadagno
informazionale rispetto alla distribuzione di riferimento P ha un costo in termini di
"distanza" da P misurata dal KL.

Appendice C — Dimensione VC di alcune classi
notevoli

Per rendere concreta la teoria, raccogliamo qui i valori della dimensione VC per alcune
classi di ipotesi standard, con le indicazioni delle dimostrazioni pertinenti.

Soglie in R: VCdim = 1. (Dimostrato nella sezione VIII, Esempio 2.)

Intervalli in R: H = {h_{a,b}:a<b,h_{a,b}@)=1_ {x € [a,b]}}. VCdim = 2. Tre
punti x1 < x2 < xs non sono frantumati: l'etichettatura (1,0,1) non è realizzabile da
nessun intervallo. Due punti qualsiasi sono frantumati: per (1,0) si prende [x1, xi]; pet
(0,1) si prende [x2, x2]; pet (1,1) si prende |[x1, x2]; pet (0,0) si prende [x3+1, x3+2].

Iperpiani in RAn: VCdim = n+1. (Dimostrato nella sezione XII via Teorema di
Radon.)

Silvano Salvador

70

Semisfere in Rn: VCdim = n+1. Si ottiene con argomento analogo agli iperpiani

tramite dualità proiettiva.

Funzioni sinusoidali h_w(x) = sign(sin(wx)) su R: VCdim = 00. Per qualsiasi m, è
possibile scegliere w sufficientemente grande da realizzare qualsiasi sequenza di m

etichette binarie per punti in posizione generale.

Alberi di decisione con k foglie: VCdim = ®(k log k). Questo risultato è non banale
e richiede argomenti di conteggio; è rilevante per la comprensione della complessità
degli ensemble di alberi decisionali.

Reti neurali con n nodi e w pesi: VCdim = O(w log(w/n)) e VCdim = Q(w
log(n/w)). (Baum-Haussler 1989, Goldberg-Jerrum 1995.) La dipendenza logaritmica
da n/w è cruciale: la dimensione VC cresce sublinearmente con w, suggerendo che
l'architettura (la topologia della rete) comprime la capacità rispetto al mero conteggio

dei pesi.

Bibliografia ragionata

La bibliografa che segue non è una lista esaustiva ma una guida per il lettore che
intenda approfondire le connessioni sviluppate nella monografia. I riferimenti sono

organizzati per tema e accompagnati da note che ne indicano la rilevanza specifica.
Sul problema dell'induzione nella sua formulazione filosofica classica

La fonte primaria irrinunciabile è David Hume, A Treatise of Human Nature (1739), in
particolare il libro I, parte III. La formulazione più lucida e accessibile del problema è
in An Enquiry Concerning Human Understanding (1748), sezione IV. L'edizione critica
curata da Beauchamp (Oxford University Press, 2000) include un apparato di note di
notevole utilità comparativa.

Il contributo di Nelson Goodman si trova in Fac Fiction, and Forecast (1955; quarta
edizione Harvard University Press, 1983). Il capitolo III, dedicato al nuovo enigma
dell'induzione, è il testo da leggere; il capitolo IV, sulla proicttabilità e l'entrenchment,
è il tentativo di risposta. La critica più acuta alla risposta goodmaniana è in Douglas
Stalker (a cura di), Grue! The New Riddle of Induction (Open Court, 1994), che raccoglie i
contributi più significativi del dibattito post-goodmaniano.

La lettura kripkeana di Wittgenstein è in Saul Kripke, Wirtgenstein on Rules and Private
Language (Harvard University Press, 1982). Il capitolo II è quello rilevante per la
monografia; il lettore è avvertito che l'attribuzione del paradosso a Wittgenstein è
controversa e che Simon Blackburn, Paul Boghossian e altri hanno sostenuto che la

risposta comunitaria di Kripke tradisce l'intento wittgensteiniano originale.
Sul programma di naturalizzazione dell'epistemologia

W.V.O. Quine, "Epistemology Naturalized" (1969), ristampato in Ownzo/ogical Relativity
and Other Essays (Columbia University Press, 1969), è il testo fondativo. La critica più

Silvano Salvador

71

incisiva rimane quella di Jaegwon Kim, "What is 'Naturalized Epistemology'?" (1988),
in Philosophical Perspectives 2, che mostra l'irriducibilità della normatività epistemica alla
descrizione empirica.

Il programma di naturalizzazione attraverso la teoria dell'apprendimento è discusso in
maniera accessibile in John Holland, Keith Holyoak, Richard Nisbett, Paul Thagard,
Induction: Processes of Inference, Learning, and Discovery (MIT Press, 1986), che anticipa
molti dei temi formali sviluppati nel decennio successivo. Il tentativo più recente e più
formalmente elaborato è in Leslie Valiant, Probably Approximately Correct (Basic Books,
2013), che espande la teoria PAC verso implicazioni cognitive e filosofiche con una
scrittura accessibile a un pubblico non specialistico.

Sulla teoria PAC-apprendibilità

L'articolo fondativo è Leslie Valiant, "A Theory of the Learnable" (1984),
Communications of the ACM 27(11), 1134-1142. È breve, denso e ancora del tutto
leggibile nella formulazione originale.

Il testo di riferimento per una trattazione sistematica e completa è Shai Shalev-Shwartz
e Shai Ben-David, Understanding Machine Learning: From Theory to Algoritbms (Cambridge
University Press, 2014), disponibile liberamente online. I capitoli II-VI coprono il
quadro PAC con derivazioni complete e attenzione alla distinzione tra modello
realizzabile e agnostico. Per un approccio più avanzato, il volume di Mobhti,
Rostamizadeh e Talwalkar, Foundations of Machine Learning (seconda edizione, MIT
Press, 2018) tratta in modo sistematico la complessità di Rademacher, i limiti
PAC-Bayes e l'apprendimento online.

Sulla dimensione di Vapnik-Cervonenkis

Il lavoro originale di Vapnik e Cervonenkis è in V.N. Vapnik e A.Ya. Chervonenkis,
"On the Uniform Convergence of Relative Frequencies of Events to Their
Probabilities" (1971), Theory of Probability and Its Applications 16(2), 264-280. La
dimostrazione del teorema di generalizzazione uniforme in questo articolo è un
capolavoro di precisione tecnica. Il Lemma di Sauer è in N. Sauer, "On the Density of
Families of Sets" (1972), Journal of Combinatorial Theory (A) 13, 145-147; la versione di
Shelah è in S. Shelah, "A Combinatorial Problem; Stability and Order for Models and
Theories in Infinitary Languages" (1972), Pacific Journal of Mathematics 41, 247-261.

Il testo più approfondito sulla dimensione VC dal punto di vista combinatorico è
Michael Anthony e Norman Biggs, Computational Learning Theory (Cambridge
University Press, 1992). Per il legame con la geometria convessa e il Teorema di
Radon, il riferimento è Rolf Schneider, Convex Bodies: The Brunn-Minkowski Theory
(Cambridge University Press, 2014), anche se un'esposizione autocontenuta del solo
Teorema di Radon si trova in qualsiasi testo di geometria combinatorica.

Sulla complessità di Rademacher

Silvano Salvador

72

L'articolo di Peter Bartlett e Shahar Mendelson, "Rademacher and Gaussian
Complexities: Risk Bounds and Structural Results" (2002), Journa/ of Machine Learning
Research 3, 463-482, è il riferimento principale. La disuguaglianza di Massart è in Pascal
Massart, "Some Applications of Concentration Inequalities to Statistics" (2000),
Annales de la Faculté des Sciences de ‘Toulouse 9(2), 245-303, che contiene anche
un'eccellente introduzione alle disuguaglianze di concentrazione in generale.

Il collegamento con la teoria dei processi gaussiani e l'integrale di Dudley è sviluppato
in Richard Dudley, Uniform Central Limit Theorems (seconda edizione, Cambridge
University Press, 2014). Per i lettori interessati alla connessione con la geometria degli
spazi di Banach, il testo di Gilles Pisier, The Volume of Convex Bodies and Banach Space
Geometry (Cambridge University Press, 1989) è il riferimento d'elezione.

Su MDL, complessità di Kolmogorov e distribuzione di Solomonoff

L'esposizione più chiara e completa della complessità di Kolmogorov è Ming Li e Paul
Vitànyi, An Introduction to Kolmogorov Complexity and Its Applications (terza edizione,
Springer, 2008). I capitoli HIV coprono le basi; il capitolo V tratta la distribuzione di
Solomonoff e il teorema di convergenza. Il testo originale di Solomonoff è in R.J.
Solomonoff, "A Formal Theory of Inductive Inference, Parts I and II" (1964),
Information and Control 7T(1-2).

Il principio MDL è esposto sistematicamente in Peter Griinwald, The Minimum
Description Length Princible (MITI Press, 2007), che è il trattamento più completo e
filosoficamente consapevole disponibile. Il capitolo XVII si occupa esplicitamente
delle relazioni tra MDL, bayesianesimo e teoria dell'apprendimento statistico. Per l'uso
del MDL nella filosofia della scienza, Gerhard Schurz e Paul Thorn hanno sviluppato
connessioni interessanti nel contesto dell'epistemologia frequentista in articoli recenti
su British Journal for the Philosophy of Science.

Sui limiti PAC-Bayes

L'articolo originale di David McAllester è "PAC-Bayesian Model Averaging" (1999),
Proceedings of COLT 1999, 164-170. Il risultato più forte è in McAllester, "Some
PAC-Bayesian Theorems" (1998/1999), Machine Learning 37(3), 355-363. Un tutorial
accessibile e aggiornato è John Shawe-Taylor e Robert Williamson, "A PAC Analysis
of a Bayesian Estimator" (1997), già in circolo come preprint, con le versioni rivedute
di Langford e Shawe-Taylor del 2002.

Per le connessioni con la teoria dell'informazione e il lemma di Donsker-Varadhan, il
riferimento classico è M.D. Donsker e S.R.S. Varadhan, "Asymptotic Evaluation of
Certain  Markov Process Expectations for Large "Time" I-IV (1975-1983),
Communications on Pure and Applied Mathematics, che tratta il caso continuo nel contesto
delle grandi deviazioni. La variante discreta usata in questa monografia è in Thomas
Cover e Joy Thomas, E/ezents of Information Theory (seconda edizione, Wiley, 2006),
capitolo XI.

Sulla stabilità algoritmica

Silvano Salvador

73

Il lavoro fondamentale è Olivier Bousquet e André Elisseeff, "Stability and
Generalization" (2002), Journal of Machine Learning Research 2, 499-526. La teoria della
stabilità uniforme è sviluppata in questo articolo con piena generalità e con le
connessioni con la regolarizzazione di Tikhonov che abbiamo discusso nella sezione
XLII.

Per gli schemi di compressione del campione, l'articolo originale è N. Littlestone e
MK. Warmuth, "Relating Data Compression and Learnability" (1986), Technical
Report, University of California Santa Cruz, con versione rivista in Machine Learning
(1995). Il teorema di Floyd-Warmuth è in S. Floyd e M.K. Warmuth, "Sample
Compression, Learnability, and the Vapnik-Chervonenkis Dimension" (1995), Machine
Learning 21(3), 269-304.

Sull'apprendimento online e la dimensione di Littlestone

L'articolo seminale è Nick Littlestone, "Learning Quickly When Irrelevant Attributes
Abound: A New Linear-Threshold Algorithm" (1988), Machine Learning 2(4), 285-318,
che introduce la dimensione di Littlestone nel contesto dell'algoritmo Winnow. La
caratterizzazione completa dell'apprendimento online attraverso la dimensione di
Littlestone è in Shai Shalev-Shwartz, "Online Learning and Online Convex
Optimization" (2012), Foundations and Trends in Machine Learning 4(2), 107-194.

Il risultato che unifica complessità di Rademacher offline e sequenziale è in Alexander
Rakhlin, Karthik Sridharan e Ambuj Tewari, "Online Learning: Random Averages,
Combinatorial Parameters, and Learnability" (2010), Advances in Neural Information
Processing Systems 23. La connessione con la teoria dei giochi minimax è discussa in
Nicolò Cesa-Bianchi e Gabor Lugosi, Prediction, Learning, and Games (Cambridge
University Press, 2006), il testo di riferimento per l'apprendimento avversariale.

Sui margini, le SVM e la dimensione di fat-shattering

Il teorema di generalizzazione tramite margini è sviluppato in Peter Bartlett, "The
Sample Complexity of Pattern Classification with Neural Networks: The Size of the
Weights is More Important than the Size of the Network" (1998), [IEEE Transactions on
Information Theory 44(2), 525-536, che introduce la dimensione di fat-shattering. Le
SVM nel loro sviluppo completo sono in Vladimir Vapnik, The Nature of Statistical
Learning Theory (seconda edizione, Springer, 2000), che è il testo più diretto dell'autore
principale della teoria.

La connessione con il realismo strutturale è esplicitata in James Ladyman, "What Is
Structural Realism?" (1998), Studies in History and Philosophy of Science 29(3), 409-424,
che è il testo fondativo del realismo strutturale ontico. Per la verisimilitudine e i
problemi di Tichy-Miller, il riferimento è Pavel Tichy, "On Popper's Definitions of
Verisimilitude" (1974), British Journal for the Philosophy of Science 25(2), 155-160, e David
Miller, "Popper's Qualitative Theory of Verisimilitude" (1974), nello stesso numero.

Sul falsificazionismo e la filosofia della scienza

Silvano Salvador

74

Karl Popper, Logik der Forschung (1934; traduzione inglese The Logic of Scientific Discovery,
Hutchinson, 1959) rimane il testo di riferimento. La critica più incisiva del
falsificazionismo dall'interno della filosofia analitica è Imre Lakatos, "Falsification and
the Methodology of Scientific Research Programmes" (1970), in Imre Lakatos e Alan
Musgrave (a cura di), Criticism and the Growth of Knowledge (Cambridge University Press,
1970), 91-196.

Il collegamento tra autorettificazione del metodo scientifico e apprendimento
sequenziale viene suggerito dall'analisi dei metodi di fissazione della credenza in
Charles Sanders Peirce, "The Fixation of Belief" (1877), Popular Science Montbby 12,
1-15, e "How to Make Our Ideas Clear" (1878), stesso periodico, 286-302.
L'interpretazione pragmatista della convergenza nella scienza è in Isaac Levi, The
Enterprise of Knowledge (MIT Press, 1980), che sviluppa il tema della revisione razionale
delle credenze in modo formalmente più rigoroso di Popper.

Sull'epistemologia bayesiana e i suoi fondamenti

Il testo di riferimento per il bayesianesimo soggettivista è Bruno de Finetti, Theory of
Probability (2 voll., Wiley, 1974-1975), in particolare il primo volume. Il teorema di
rappresentazione di de Finetti — secondo cui le sequenze scambiabili di eventi hanno
rappresentazione come miscele di Bernoulli, il che fondisce soggettivismo e
frequentismo —è il risultato più importante per la connessione con la teoria PAC,
poiché la scambiabilità è una condizione più debole dell'i.i.d. e più facilmente
difendibile dal punto di vista filosofico.

Per l'aggiornamento bayesiano e la sua relazione con l'apprendimento, il testo di Colin
Howson e Peter Urbach, Scientific Reasoning: The Bayesian Approach (terza edizione, Open
Court, 2006) è il trattamento più accessibile e filosoficamente responsabile. Le
connessioni con l'apprendimento statistico e la loro relazione con il programma di
naturalizzazione sono discusse nel raccoglitore di saggi curato da David Corfield e Jon
Williamson, Foundations of Bayesianism (Kluwer, 2001).

## Notes

- 自動収集された未処理ノート。notes/ フォルダへの統合前に内容と出典を確認する。

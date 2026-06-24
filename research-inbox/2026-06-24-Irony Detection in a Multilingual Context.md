---
source: "https://arxiv.org/abs/2002.02427v1"
title: "Irony Detection in a Multilingual Context"
author: "Bilal Ghanem, Jihen Karoui, Farah Benamara, Paolo Rosso, Véronique Moriceau"
year: "2020"
publication: "arXiv preprint / cs.CL"
download: "https://arxiv.org/pdf/2002.02427v1"
pdf: "https://arxiv.org/pdf/2002.02427v1"
captured_at: "2026-06-24T11:12:01Z"
updated_at: "2026-06-24T11:12:01Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "リチャード・ローティ"
query: "Rorty Contingency Irony and Solidarity"
tags:
  - "現代思想"
  - "プラグマティズム"
  - "ネオプラグマティズム"
  - "反表象主義"
status: raw
---

# Irony Detection in a Multilingual Context

- 著者: Bilal Ghanem, Jihen Karoui, Farah Benamara, Paolo Rosso, Véronique Moriceau
- 年: 2020
- 掲載情報: arXiv preprint / cs.CL
- 情報源: [arxiv](https://arxiv.org/abs/2002.02427v1)
- ダウンロード: https://arxiv.org/pdf/2002.02427v1
- PDF: https://arxiv.org/pdf/2002.02427v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

This paper proposes the first multilingual (French, English and Arabic) and multicultural (Indo-European languages vs. less culturally close languages) irony detection system. We employ both feature-based models and neural architectures using monolingual word representation. We compare the performance of these systems with state-of-the-art systems to identify their capabilities. We show that these monolingual models trained separately on different languages using multilingual word representation or text-based features can open the door to irony detection in languages that lack of annotated data for irony.

## PDF Text

Irony Detection in a Multilingual Context
Bilal Ghanem1 , Jihen Karoui2, Farah Benamara3, Paolo Rosso1 , and
Véronique Moriceau3

arXiv:2002.02427v1 [cs.CL] 6 Feb 2020

1

PRHLT Research Center, Universitat Politècnica de València, Spain
{bigha@doctor, prosso@dsic}.upv.es
2
AUSY R&D, Paris, France jkaroui@ausy.fr
3
IRIT, CNRS, Université de Toulouse, France
{benamara,moriceau}@irit.fr

Abstract. This paper proposes the first multilingual (French, English and Arabic) and multicultural (Indo-European languages vs. less culturally close languages) irony detection system. We employ both featurebased models and neural architectures using monolingual word representation. We compare the performance of these systems with state-of-theart systems to identify their capabilities. We show that these monolingual models trained separately on different languages using multilingual word representation or text-based features can open the door to irony detection in languages that lack of annotated data for irony.
Keywords: Irony Detection · Social Media · Multilingual Embeddings

Motivations
Figurative language makes use of figures of speech to convey non-literal meaning [16,2]. It encompasses a variety of phenomena, including metaphor, humor, and irony. We focus here on irony and uses it as an umbrella term that covers satire, parody and sarcasm.
Irony detection (ID) has gained relevance recently, due to its importance to extract information from texts. For example, to go beyond the literal matches of user queries, Veale enriched information retrieval with new operators to enable the non-literal retrieval of creative expressions [40]. Also, the performances of sentiment analysis systems drastically decrease when applied to ironic texts [5,19].
Most related work concern English [17,21] with some efforts in French [23], Portuguese [7], Italian [14], Dutch [26], Hindi [37], Spanish variants [31] and Arabic
[22,11]. Bilingual ID with one model per language has also been explored, like
English-Czech [32] and English-Chinese [38], but not within a cross-lingual perspective.
In social media, such as Twitter, specific hashtags (#irony, #sarcasm) are often used as gold labels to detect irony in a supervised learning setting. Although recent studies pointed out the issue of false-alarm hashtags in self-labeled data [20], ID via hashtag filtering provides researchers positive examples with

2

Ghanem et al.

high precision. On the other hand, systems are not able to detect irony in languages where such filtering is not always possible. Multilingual prediction (either relying on machine translation or multilingual embedding methods) is a common solution to tackle under-resourced languages [6,33]. While multilinguality has been widely investigated in information retrieval [27,34] and several NLP
tasks (e.g., sentiment analysis [3,4] and named entity recognition [30]), no one explored it for irony.
We aim here to bridge the gap by tackling ID in tweets from both multilingual
(French, English and Arabic) and multicultural perspectives (Indo-European languages whose speakers share quite the same cultural background vs. less culturally close languages). Our approach does not rely either on machine translation or parallel corpora (which are not always available), but rather builds on previous corpus-based studies that show that irony is a universal phenomenon and many languages share similar irony devices. For example, Karoui et. al [24] concluded that their multi-layer annotated schema, initially used to annotate French tweets, is portable to English and Italian, observing relatively the same tendencies in terms of irony categories and markers. Similarly, Chakhachiro [8] studies irony in English and Arabic, and shows that both languages share several similarities in the rhetorical (e.g., overstatement), grammatical (e.g., redundancy)
and lexical (e.g., synonymy) usage of irony devices. The next step now is to show to what extent these observations are still valid from a computational point of view. Our contributions are:
I. A new freely available corpus of Arabic tweets manually annotated for irony detection4 .
II. Monolingual ID : We propose both feature-based models (relying on languagedependent and language-independent features) and neural models to measure to what extent ID is language dependent.
III. Cross-lingual ID : We experiment using cross-lingual word representation by training on one language and testing on another one to measure how the proposed models are culture-dependent. Our results are encouraging and open the door to ID in languages that lack of annotated data for irony.

Data
Arabic dataset (Ar=11, 225 tweets). Our starting point was the corpus built by [22] that we extended to different political issues and events related to the
Middle East and Maghreb that hold during the years 2011 to 2018. Tweets were collected using a set of predefined keywords (which targeted specific political


figures or events) and containing or not Arabic ironic hashtags ( éK Qm #, èQjÓ#,

 #) 5 . The collection process resulted in a set of 6, 809 ironic tweets
ÕºîE#, Z@QîD@

(I) vs. 15, 509 non ironic (N I) written using standard (formal) and different
Arabic language varieties: Egypt, Gulf, Levantine, and Maghrebi dialects.
4
5

The corpus is available at https://github.com/bilalghanem/multilingual irony
All of these words are synonyms where they mean ”Irony”.

Irony Detection in a Multilingual Context

3

To investigate the validity of using the original tweets labels, a sample of
3, 000 I and 3, 000 N I was manually annotated by two Arabic native speakers which resulted in 2, 636 I vs. 2, 876 N I. The inter-annotator agreement using Cohen’s Kappa was 0.76, while the agreement score between the annotators’ labels and the original labels was 0.6. Agreements being relatively good knowing the difficulty of the task, we sampled 5, 713 instances from the original unlabeled dataset to our manually labeled part. The added tweets have been manually checked to remove duplicates, very short tweets and tweets that depend on external links, images or videos to understand their meaning.
French dataset (Fr=7, 307 tweets). We rely on the corpus used for the
DEFT 2017 French shared task on irony [5] which consists of tweets relative to a set of topics discussed in the media between 2014 and 2016 and contains topic keywords and/or French irony hashtags (#ironie, #sarcasme). Tweets have been annotated by three annotators (after removing the original labels) with a reported Cohen’s Kappa of 0.69.
English dataset (En=11, 225 tweets). We use the corpus built by [32] which consists of 100, 000 tweets collected using the hashtag #sarcasm. It was used as benchmark in several works [13,18]. We sliced a subset of approximately 11, 200
tweets to match the sizes of the other languages’ datasets.
Table 1 shows the tweet distribution in all corpora. Across the three languages, we keep a similar number of instances for train and test sets to have fair cross-lingual experiments as well (see Section 1). Also, for French, we use the original dataset without any modification, keeping the same number of records for train and test to better compare with state-of-the-art results. For the classes distribution (ironic vs. non ironic), we do not choose a specific ratio but we use the resulted distribution from the random shuffling process.
Table 1. Tweet distribution in all corpora.

Ar
Fr
En

# Ironic # Not-Ironic Train Test
6, 005
5, 220
10, 219 1, 006
2, 425
4, 882
5, 843 1, 464
5, 602
5, 623
10, 219 1, 006

Monolingual Irony Detection
It is important to note that our aim is not to outperform state-of-the-art models in monolingual ID but to investigate which of the monolingual architectures
(neural or feature-based) can achieve comparable results with existing systems.
The result can show which kind of features works better in the monolingual settings and can be employed to detect irony in a multilingual setting. In addition, it can show us to what extend ID is language dependent by comparing their results to multilingual results. Two models have been built, as explained below.
Prior to learning, basic preprocessing steps were performed for each language
(e.g., removing foreign characters, ironic hashtags, mentions, and URLs).

4

Ghanem et al.

Feature-based models. We used state-of-the-art features that have shown to be useful in ID: some of them are language-independent (e.g., punctuation marks, positive and negative emoticons, quotations, personal pronouns, tweet’s length, named entities) while others are language-dependent relying on dedicated lexicons (e.g., negation, opinion lexicons, opposition words). Several classical machine learning classifiers were tested with several feature combinations, among them Random Forest (RF) achieved the best result with all features.
Neural model with monolingual embeddings. We used Convolutional
Neural Network (CNN) network whose structure is similar to the one proposed by
[25]. For the embeddings, we relied on AraV ec [36] for Arabic, FastText [15] for
French, and Word2vec Google News [29] for English 6 . For the three languages, the size of the embeddings is 300 and the embeddings were fine-tuned during the training process. The CNN network was tuned with 20% of the training corpus using the Hyperopt7 library.
Results. Table 2 shows the results obtained when using train-test configurations for each language. For English, our results, in terms of macro F-score
(F ), were not comparable to those of [32,39], as we used 11% of the original dataset. For French, our scores are in line with those reported in state of the art (cf. best system in the irony shared task achieved F = 78.3 [5]). They outperform those obtained for Arabic (A = 71.7) [22] and are comparable to those recently reported in the irony detection shared task in Arabic tweets [11,12]
(F = 84.4). Overall, the results show that semantic-based information captured by the embedding space are more productive comparing to standard surface and lexicon-based features.
Table 2. Results of the monolingual experiments (in percentage) in terms of accuracy
(A), precision (P), recall (R), and macro F-score (F).
Arabic
French
English
A
P R
F
A
P R
F
A
P R
F
RF 68.0 67.0 82.0 68.0 68.5 71.7 87.3 61.0 61.2 60.0 70.0 61.0
CNN 80.5 79.1 84.9 80.4 77.6 68.2 59.6 73.5 77.9 74.6 84.7 77.8

Cross-lingual Irony Detection
We use the previous CNN architecture with bilingual embedding and the RF
model with surface features (e.g., use of personal pronoun, presence of interjections, emoticon or specific punctuation)8 to verify which pair of the three languages: (a) has similar ironic pragmatic devices, and (b) uses similar text-based pattern in the narrative of the ironic tweets. As continuous word embedding spaces exhibit similar structures across (even distant) languages [28], we use a
6

Other available pretrained embeddings models have also been tested.
https://github.com/hyperopt/hyperopt
8
To avoid language dependencies, we rely on surface features only discarding those that require external semantic resources or morpho-syntactic parsing.

7

Irony Detection in a Multilingual Context

5

multilingual word representation which aims to learn a linear mapping from a source to a target embedding space. Many methods have been proposed to learn this mapping such as parallel data supervision and bilingual dictionaries [28]
or unsupervised methods relying on monolingual corpora [10,1,41]. For our experiments, we use Conneau et al ’s approach as it showed superior results with respect to the literature [10]. We perform several experiments by training on one language (lang1 ) and testing on another one (lang2 ) (henceforth lang1 → lang2 ).
We get 6 configurations, plus two others to evaluate how irony devices are expressed cross-culturally, i.e. in European vs. non European languages. In each experiment, we took 20% from the training to validate the model before the testing process. Table 3 presents the results.
Table 3. Results of the cross-lingual experiments.
CNN
RF
Train→Test A P R
F
A
P R
F
Ar→Fr
60.1 37.2 26.6 51.7 47.03 29.9 43.9 46.0
Fr→Ar
57.8 62.9 45.7 57.3 51.11 61.1 24.0 54.0
Ar→En
48.5 26.5 17.9 34.1 49.67 49.7 66.2 50.0
En→Ar
56.7 57.7 62.3 56.4 52.5 58.6 38.5 53.0
Fr→En
53.0 67.9 11.0 42.9 52.38 52.0 63.6 52.0
En→Fr
56.7 33.5 29.5 50.0 56.44 74.6 52.7 58.0
(En/Fr)→Ar 62.4 66.1 56.8 62.4 55.08 56.7 68.5 62.0
Ar→(En/Fr) 56.3 33.9 09.5 42.7 59.84 60.0 98.7 74.6

From a semantic perspective, despite the language and cultural differences between Arabic and French languages, CNN results show a high performance comparing to the other languages pairs when we train on each of these two languages and test on the other one. Similarly, for the French and English pair, but when we train on French they are quite lower. We have a similar case when we train on Arabic and test on English. We can justify that by, the language presentation of the Arabic and French tweets are quite informal and have many dialect words that may not exist in the pretrained embeddings we used comparing to the English ones (lower embeddings coverage ratio), which become harder for the
CNN to learn a clear semantic pattern. Another point is the presence of Arabic dialects, where some dialect words may not exist in the multilingual pretrained embedding model that we used. On the other hand, from the text-based perspective, the results show that the text-based features can help in the case when the semantic aspect shows weak detection; this is the case for the Ar −→ En configuration. It is worthy to mention that the highest result we get in this experiment is from the En→Fr pair, as both languages use Latin characters. Finally, when investigating the relatedness between European vs. non European languages (cf.
(En/Fr)→Ar), we obtain similar results than those obtained in the monolingual experiment (macro F-score 62.4 vs. 68.0) and best results are achieved by Ar
→(En/Fr). This shows that there are pragmatic devices in common between both sides and, in a similar way, similar text-based patterns in the narrative way of the ironic tweets.

6

Ghanem et al.

Discussions and Conclusion
This paper proposes the first multilingual ID in tweets. We show that simple monolingual architectures (either neural or feature-based) trained separately on each language can be successfully used in a multilingual setting providing a crosslingual word representation or basic surface features. Our monolingual results are comparable to state of the art for the three languages. The CNN architecture trained on cross-lingual word representation shows that irony has a certain similarity between the languages we targeted despite the cultural differences which confirm that irony is a universal phenomena, as already shown in previous linguistic studies [35,24,9]. The manual analysis of the common misclassified tweets across the languages in the multilingual setup, shows that classification errors are due to three main factors. (1) First, the absence of context where writers did not provide sufficient information to capture the ironic sense even in the monolin

gual setting, as in !! ¼PAJ.Ó úæk ¡® ¡® úGAK @YJ.K (Let’s start again, get off

get off Mubarak!! ) where the writer mocks the Egyptian revolution, as the actual president ”Sisi” is viewed as Mubarak’s fellows. (2) Second, the presence of out of vocabulary (OOV) terms because of the weak coverage of the mutlilingual embeddings which make the system fails to generalize when the OOV set of unseen words is large during the training process. We found tweets in all the three languages written in a very informal way, where some characters of the words were deleted, duplicated or written phonetically (e.g phat instead of fat ). (3) Another important issue is the difficulty to deal with the Arabic language. Arabic tweets are often characterized by non-diacritised texts, a large variations of unstandardized dialectal Arabic (recall that our dataset has 4 main varieties, namely
Egypt, Gulf, Levantine, and Maghrebi), presence of transliterated words (e.g. the

word table becomes éÊK. A£ (tabla)), and finally linguistic code switching between
Modern Standard Arabic and several dialects, and between Arabic and other languages like English and French. We found some tweets contain only words from one of the varieties and most of these words do not exist in the Arabic embed

dings model. For example in QåÓ# éK @ èBð àAJ« ñë ..  ÓAÓ ÐñK ÐA¿ éËA®K. ¼PAJ.Ó

(Since many days Mubarak didn’t die .. is he sick or what? #Egypt ), only the words ÐñK (day), ¼PAJ.Ó (Mubarak), and ñë (he) exist in the embeddings. Clearly, considering only these three available words, we are not able to understand the context or the ironic meaning of the tweet.
To conclude, our multilingual experiments confirmed that the door is open towards multilingual approaches for ID. Furthermore, our results showed that
ID can be applied to languages that lack of annotated data. Our next step is to experiment with other languages such as Hindi and Italian.

Acknowledgment
The work of Paolo Rosso was partially funded by the Spanish MICINN under the research project MISMIS-FAKEnHATE (PGC2018-096212-B-C31).

Irony Detection in a Multilingual Context

7

References
1. Artetxe, M., Labaka, G., Agirre, E., Cho, K.: Unsupervised Neural Machine Translation. arXiv preprint (2017)
2. Attardo, S.: Irony as Relevant Inappropriateness. Journal of Pragmatics 32(6),
793–826 (2000)
3. Balahur, A., Turchi, M.: Comparative Experiments Using Supervised Learning and
Machine Translation for Multilingual Sentiment Analysis. Comput. Speech Lang.
28(1), 56–75 (2014)
4. Barnes, J., Klinger, R., Schulte im Walde, S.: Bilingual Sentiment Embeddings:
Joint Projection of Sentiment Across Languages. In: Proceedings of the 56th Annual Meeting of the Association for Computational Linguistics (Volume 1: Long
Papers). pp. 2483–2493. Association for Computational Linguistics (2018)
5. Benamara, F., Grouin, C., Karoui, J., Moriceau, V., Robba, I.: Analyse d’opinion et langage figuratif dans des tweets présentation et résultats du Défi Fouille de
Textes DEFT2017. In: Actes de DEFT@TALN2017. Orléans, France (2017)
6. Bikel, D., Zitouni, I.: Multilingual Natural Language Processing Applications:
From Theory to Practice. IBM Press, 1st edn. (2012)
7. Carvalho, P., Sarmento, L., Silva, M.J., Oliveira, E.D.: Clues for Detecting Irony in User-Generated Contents: Oh...!! It’s ”so easy”;-). In: Proceedings of the 1st international CIKM workshop on Topic-sentiment analysis for mass opinion. pp.
53–56. ACM (2009)
8. Chakhachiro, R.: Translating Irony in Political Commentary Texts from English into Arabic. Babel 53(3), 216–240 (2007)
9. Colston, H.L.: Irony as Indirectness Cross-Linguistically: On the Scope of Generic
Mechanisms. In: Indirect Reports and Pragmatics in the World Languages, pp.
109–131. Springer (2019)
10. Conneau, A., Lample, G., Ranzato, M., Denoyer, L., Jégou, H.: Word Translation
Without Parallel Data. arXiv preprint (2017)
11. Ghanem, B., Karoui, J., Benamara, F., Moriceau, V., Rosso, P.: Idat at fire2019:
Overview of the track on irony detection in arabic tweets. In: Proceedings of the
11th Forum for Information Retrieval Evaluation. pp. 10–13. ACM (2019)
12. Ghanem, B., Karoui, J., Benamara, F., Moriceau, V., Rosso, P.: Idat@fire2019:
Overview of the track on irony detection in arabic tweets. In: Working Notes of the Forum for Information Retrieval Evaluation (FIRE 2019). CEUR Workshop
Proceedings. In: CEUR-WS.org, Kolkata, India. vol. 2517, pp. 380–390 (2019)
13. Ghanem, B., Rangel, F., Rosso, P.: LDR at SemEval-2018 Task 3: A Low Dimensional Text Representation for Irony Detection. In: Proceedings of The 12th
International Workshop on Semantic Evaluation. pp. 531–536 (2018)
14. Gianti, A., Bosco, C., Patti, V., Bolioli, A., Caro, L.D.: Annotating Irony in a Novel
Italian Corpus for Sentiment Analysis. In: Proceedings of the 4th Workshop on
Corpora for Research on Emotion Sentiment and Social Signals, Istanbul, Turkey.
pp. 1–7 (2012)
15. Grave, E., Bojanowski, P., Gupta, P., Joulin, A., Mikolov, T.: Learning Word
Vectors for 157 Languages. CoRR abs/1802.06893 (2018)
16. Grice, H.P.: Logic and Conversation. In: Cole, P., Morgan, J.L. (eds.) Speech Acts.
Syntax and Semantics, Volume 3, pp. 41–58. Academic Press, New York (1975)
17. Hee, C.V., Lefever, E., Hoste, V.: SemEval-2018 Task 3: Irony Detection in English
Tweets. In: Proceedings of The 12th International Workshop on Semantic Evaluation, SemEval@NAACL-HLT, New Orleans, Louisiana, June 5-6, 2018. pp. 39–50
(2018)

8

Ghanem et al.

18. Hernández Farı́as, D.I., Bosco, C., Patti, V., Rosso, P.: Sentiment Polarity Classification of Figurative Language: Exploring the Role of Irony-Aware and Multifaceted Affect Features. In: International Conference on Computational Linguistics and Intelligent Text Processing. pp. 46–57. Springer (2017)
19. Hernández Farı́as, D.I., Patti, V., Rosso, P.: Irony Detection in Twitter: The Role of Affective Content. ACM Transactions on Internet Technology (TOIT) 16(3),
19 (2016)
20. Huang, H.H., Chen, C.C., Chen, H.H.: Disambiguating False-Alarm Hashtag Usages in Tweets for Irony Detection. In: Proceedings of the 56th Annual Meeting of the Association for Computational Linguistics (Volume 2: Short Papers) (2018)
21. Huang, Y.H., Huang, H.H., Chen, H.H.: Irony Detection with Attentive Recurrent
Neural Networks. In: European Conference on Information Retrieval. pp. 534–540.
Springer (2017)
22. Karoui, J., Benamara, F., Moriceau, V.: SOUKHRIA: Towards an Irony Detection
System for Arabic in Social Media. In: Third International Conference On Arabic
Computational Linguistics, ACLING 2017, November 5-6, 2017, Dubai, United
Arab Emirates. pp. 161–168 (2017)
23. Karoui, J., Benamara, F., Moriceau, V., Aussenac-Gilles, N., Belguith, L.H.: Towards a Contextual Pragmatic Model to Detect Irony in Tweets. In: Proceedings of the 53rd Annual Meeting of the Association for Computational Linguistics and the
7th International Joint Conference on Natural Language Processing of the Asian
Federation of Natural Language Processing : Volume 2 short papers. pp. 644–650.
ACL-IJCNLP’15 (2015)
24. Karoui, J., Benamara, F., Moriceau, V., Patti, V., Bosco, C., Aussenac-Gilles, N.:
Exploring the Impact of Pragmatic Phenomena on Irony Detection in Tweets: A
Multilingual Corpus Study. In: Proceedings of the 15th Conference of the European
Chapter of the Association for Computational Linguistics: Volume 1, Long Papers.
pp. 262–272. Association for Computational Linguistics (2017)
25. Kim, Y.: Convolutional Neural Networks for Sentence Classification. In: Proceedings of the 2014 Conference on Empirical Methods in Natural Language Processing
(EMNLP). pp. 1746–1751. Association for Computational Linguistics (2014)
26. Liebrecht, C., Kunneman, F., van den, B.A.: The Perfect Solution for Detecting
Sarcasm in Tweets# Not. In: Proceedings of the 4th Workshop on Computational
Approaches to Subjectivity, Sentiment and Social Media Analysis. pp. 29–37. New
Brunswick, NJ: ACL (2013)
27. Litschko, R., Glavaš, G., Ponzetto, S.P., Vulić, I.: Unsupervised cross-lingual information retrieval using monolingual data only. In: The 41st International ACM
SIGIR Conference on Research & Development in Information Retrieval. pp. 1253–
1256. SIGIR ’18 (2018)
28. Mikolov, T., Chen, K., Corrado, G., Dean, J.: Efficient Estimation of Word Representations in Vector Space. CoRR abs/1301.3781 (2013)
29. Mikolov, T., Yih, W.t., Zweig, G.: Linguistic Regularities in Continuous Space
Word Representations. In: Proceedings of the 2013 Conference of the North American Chapter of the Association for Computational Linguistics: Human Language
Technologies. pp. 746–751 (2013)
30. Ni, J., Florian, R.: Improving Multilingual Named Entity Recognition with
Wikipedia Entity Type Mapping. CoRR abs/1707.02459 (2017)
31. Ortega-Bueno, R., Rangel, F., Hernández Farıas, D., Rosso, P., Montes-y Gómez,
M., Medina Pagola, J.E.: Overview of the Task on Irony Detection in Spanish
Variants. In: Proceedings of the Iberian Languages Evaluation Forum (IberLEF

Irony Detection in a Multilingual Context

9

2019), co-located with 34th Conference of the Spanish Society for Natural Language
Processing (SEPLN 2019). CEUR-WS. org (2019)
32. Ptáček, T., Habernal, I., Hong, J.: Sarcasm Detection on Czech and English Twitter. In: Proceedings of COLING 2014, the 25th International Conference on Computational Linguistics: Technical Papers. pp. 213–223 (2014)
33. Ruder, S.: A survey of cross-lingual embedding models. CoRR abs/1706.04902
(2017)
34. Sasaki, S., Sun, S., Schamoni, S., Duh, K., Inui, K.: Cross-Lingual Learning-toRank with Shared Representations. In: Proceedings of the 2018 Conference of the
North American Chapter of the Association for Computational Linguistics: Human
Language Technologies, Volume 2 (Short Papers). pp. 458–463 (2018)
35. Sigar, A., Taha, Z.: A Contrastive Study of Ironic Expressions in English and
Arabic. College of Basic Education Researchers Journal 12(2), 795–817 (2012)
36. Soliman, A.B., Eissa, K., El-Beltagy, S.R.: AraVec: A set of Arabic Word Embedding Models for use in Arabic NLP. In: Third International Conference On Arabic
Computational Linguistics, ACLING 2017, November 5-6, 2017, Dubai, United
Arab Emirates. pp. 256–265 (2017)
37. Swami, S., Khandelwal, A., Singh, V., Akhtar, S.S., Shrivastava, M.: A Corpus of
English-Hindi Code-Mixed Tweets for Sarcasm Detection. 19th International Conference on Computational Linguistics and Intelligent Text Processing (CICLing)
(2018)
38. Tang, Y., Chen, H.: Chinese Irony Corpus Construction and Ironic Structure Analysis. In: COLING 2014, 25th International Conference on Computational Linguistics, Proceedings of the Conference: Technical Papers, August 23-29, 2014, Dublin,
Ireland. pp. 1269–1278 (2014)
39. Tay, Y., Luu, A.T., Hui, S.C., Su, J.: Reasoning with Sarcasm by Reading InBetween. In: Proceedings of the 56th Annual Meeting of the Association for Computational Linguistics (Volume 1: Long Papers). pp. 1010–1020 (2018)
40. Veale, T.: Creative Language Retrieval: A Robust Hybrid of Information Retrieval and Linguistic Creativity. In: Proceedings of the 49th Annual Meeting of the Association for Computational Linguistics: Human Language Technologies - Volume
1. pp. 278–287. HLT ’11 (2011)
41. Wada, T., Iwata, T.: Unsupervised Cross-lingual Word Embedding by Multilingual
Neural Language Models. arXiv preprint (2018)

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

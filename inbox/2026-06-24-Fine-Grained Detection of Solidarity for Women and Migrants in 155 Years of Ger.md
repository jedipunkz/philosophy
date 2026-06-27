---
source: "https://arxiv.org/abs/2210.04359v3"
title: "Fine-Grained Detection of Solidarity for Women and Migrants in 155 Years of German Parliamentary Debates"
author: "Aida Kostikova, Benjamin Paassen, Dominik Beese, Ole Pütz, Gregor Wiedemann, Steffen Eger"
year: "2022"
publication: "arXiv preprint / cs.CL"
download: "https://arxiv.org/pdf/2210.04359v3"
pdf: "https://arxiv.org/pdf/2210.04359v3"
captured_at: "2026-06-24T12:58:11Z"
updated_at: "2026-06-24T12:58:11Z"
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

# Fine-Grained Detection of Solidarity for Women and Migrants in 155 Years of German Parliamentary Debates

- 著者: Aida Kostikova, Benjamin Paassen, Dominik Beese, Ole Pütz, Gregor Wiedemann, Steffen Eger
- 年: 2022
- 掲載情報: arXiv preprint / cs.CL
- 情報源: [arxiv](https://arxiv.org/abs/2210.04359v3)
- ダウンロード: https://arxiv.org/pdf/2210.04359v3
- PDF: https://arxiv.org/pdf/2210.04359v3

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Solidarity is a crucial concept to understand social relations in societies. In this paper, we explore fine-grained solidarity frames to study solidarity towards women and migrants in German parliamentary debates between 1867 and 2022. Using 2,864 manually annotated text snippets (with a cost exceeding 18k Euro), we evaluate large language models (LLMs) like Llama 3, GPT-3.5, and GPT-4. We find that GPT-4 outperforms other LLMs, approaching human annotation quality. Using GPT-4, we automatically annotate more than 18k further instances (with a cost of around 500 Euro) across 155 years and find that solidarity with migrants outweighs anti-solidarity but that frequencies and solidarity types shift over time. Most importantly, group-based notions of (anti-)solidarity fade in favor of compassionate solidarity, focusing on the vulnerability of migrant groups, and exchange-based anti-solidarity, focusing on the lack of (economic) contribution. Our study highlights the interplay of historical events, socio-economic needs, and political ideologies in shaping migration discourse and social cohesion. We also show that powerful LLMs, if carefully prompted, can be cost-effective alternatives to human annotation for hard social scientific tasks.

## PDF Text

Fine-Grained Detection of Solidarity for Women and Migrants in
155 Years of German Parliamentary Debates
Aida Kostikova1 , Benjamin Paassen1 , Dominik Beese2 ,
Ole Pütz1 , Gregor Wiedemann3 , Steffen Eger4,5
1
Bielefeld University, 2 TU Darmstadt, 3 Hans-Bredow-Institut,
4
University of Mannheim, 5 University of Technology Nuremberg aida.kostikova@uni-bielefeld.de

arXiv:2210.04359v3 [cs.CL] 21 Nov 2024

Abstract
Solidarity is a crucial concept to understand social relations in societies. In this paper, we explore fine-grained solidarity frames to study solidarity towards women and migrants in German parliamentary debates between 1867 and
2022. Using 2,864 manually annotated text snippets (with a cost exceeding 18k Euro), we evaluate large language models (LLMs) like
Llama 3, GPT-3.5, and GPT-4. We find that
GPT-4 outperforms other LLMs, approaching human annotation quality. Using GPT-4, we automatically annotate more than 18k further instances (with a cost of around 500 Euro)
across 155 years and find that solidarity with migrants outweighs anti-solidarity but that frequencies and solidarity types shift over time.
Most importantly, group-based notions of (anti)solidarity fade in favor of compassionate solidarity, focusing on the vulnerability of migrant groups, and exchange-based anti-solidarity, focusing on the lack of (economic) contribution.
Our study highlights the interplay of historical events, socio-economic needs, and political ideologies in shaping migration discourse and social cohesion. We also show that powerful LLMs, if carefully prompted, can be costeffective alternatives to human annotation for hard social scientific tasks.

1

Introduction

Solidarity is a crucial concept for understanding how societies achieve and maintain stability and cohesion (Reynolds, 2014), and it plays a critical role in shaping policies (Laitinen and Pessi, 2014).
Traditionally, solidarity relied on common identity and reciprocity, potentially excluding out-groups like migrants (Hopman and Knijn, 2022) and reinforcing social boundaries and hierarchies (Anthias, 2014). However, growing diversity and increasing socio-economic, political, religious, and cultural complexities of modern societies (Kymlicka, 2020) call such traditional forms of solidarity into question. Simultaneously, recent populist

Immigrants must receive the guaranteed minimum wage just like everyone else and the same support as everyone else.
This text expresses...

Solidar ity

Mixed

None

Anti-solidar ity based on...

based on...
Shared identity

Groupbased

Perceived differences

Valuing for contributions

Exchangebased

Imbalance in contribution

Support for marginalized groups

Compassionate

Denying help to vulnerable

Respect for diversity

Empathic

Disregard of groups'
differences

Figure 1: Annotation scheme based on Thijssen (2012).
The scheme categorizes statements into solidarity, antisolidarity, mixed, and none (high-level). At the finegrained level, solidarity and anti-solidarity are further divided into group-based, exchange-based, compassionate, and empathic subtypes.

movements challenge policies that support equality, such as equal opportunity or reproductive rights of women (Inglehart and Norris, 2016). These evolving complexities motivate a deeper and broader study of social solidarity, namely (i) a fine-grained exploration of different forms of social solidarity to reflect its multifaceted nature (Oosterlynck and
Bouchaute, 2013), and (ii) a broader historical analysis to trace its evolution from the 19th century to today (Banting and Kymlicka, 2017). In this work, we contribute to such a systematic study of solidarity by tracing fine-grained notions of solidarity and anti-solidarity towards two target groups, women and migrants, in political speech, namely German parliamentary debates from 1867 to 2022 (Walter et al., 2021). We employ the social solidarity framework by Thijssen (2012) that incorporates rational
(group-based and exchange-based) and emotive

Gold Standard

Translation of the Original German Text

(1) Compassionate solidarity towards women
(June 29, 1961)

“In connection with § 1708 BGB, the Bundestag has set the age of 18 as the limit for the obligation to provide maintenance. In the transitional provisions, this stipulation has been repealed for those who had already reached the age of 16 on January 1,
1962. My faction finds this regulation unfair, as it would exempt significant groups of people from this maintenance obligation. Especially women who have made great efforts to send their children to higher education, for example, would have to bear these costs alone. [...]”

(2) Exchange-based anti-solidarity towards migrants
(Apr. 19, 2018)

“[...] Let me also add: Migration is not necessarily successful – you always act as if that is great – it can fail, and it fails in particular when the immigrants’
qualifications are low. In 2013, before the so-called refugee wave, 40 percent of immigrants from non-EU countries had no qualifications. Since the wave of refugees, stabbings have increased by 20 percent, and we have imported anti-Semitism in the country. Does this make for an outstandingly successful migration?”

(3) Mixed stance towards migrants
(Feb. 2, 1982)

“[...] We must accept that in a few years we will again need a higher number of foreign workers in the Federal Republic, as Mr. Urbaniak hinted earlier. In reality, therefore, we must commit to effective integration, which admittedly requires
[...] that there can be no exceptions, no alternative, regarding the recruitment stop and the prevention of illegal immigration. [...]”

(4) None case (women)
(June 17, 2015)

“[...] ‘We want to be free people!’ There is probably no better phrase to open today’s debate here in the German Bundestag about the popular uprising of 1953. [...] We remember women and men who, 62 years ago, showed great courage because they wanted to change the course of their country’s development and their own lives, because they wanted to be free people.”

Table 1: Example sentences from our dataset showing (anti-)solidarity towards women/migrants. Bold text is the main sentence, the other sentences are for context. Original German texts, as well as examples of mixed stance and none, are available in Table 3 in the Appendix.

(compassionate and empathic) elements of solidarity (refer to Fig. 1 and Section 4 for more details on the typology). We focus on migrants, central for solidarity discourse in European and German politics (Thränhardt, 1993; Faist, 1994; Fröhlich, 2023;
Lehr, 2015), and women as an “oppressed majority”
historically marginalized from public life (Calloni,
2020). As manual annotation of (anti-)solidarity concepts in all parliamentary proceedings over this
155-year period using traditional sociological methods is practically infeasible, we explore the use of language models for this complex task. In particular, we assess the efficacy of BERT, Llama-3,
GPT-3.5, and GPT-4, to detect expressions of various (anti-)solidarity types in parliamentary texts, aiming to identify the best performing model for our large-scale analysis. From an NLP perspective, this task is semantically and pragmatically challenging because, (i) expressions of (anti-)solidarity are often implied rather than explicitly stated in the text and their meaning is affected by the political and historical context in which they are made (see the examples in Table 1; Sravanthi et al., 2024); (ii)
German data, especially evolving German language over 155 years, is under-represented in common training data sets for LLMs, which may affect performance (Ahuja et al., 2023; Qin et al., 2024; Liu

et al., 2024); and (iii) LLMs might struggle with annotating complex sociological concepts, achieving lower quality and reliability compared to human annotators (Wang et al., 2021; Ding et al., 2022;
Zhu et al., 2023; Pangakis et al., 2023).
Our contributions are: (i) We provide a human annotated training & evaluation dataset of 2,864
text snippets, which required 40+ hours weekly from 4-5 annotators over nine months, totaling an investment of approximately 18k Euro; (ii) we conduct a comparative analysis of LLMs on a complex sociological task in which pre-trained language models (esp. GPT-4) outperform an open-source model Llama-3-70B-Instruct, as well as models fine-tuned for this task (BERT, GPT-3.5 fine tuned);
(iii) we provide fine-grained insights into solidarity discourse concerning migrants in Germany in the last 155 years across different political parties.
We make our code and data available at https:
//github.com/DominikBeese/FairGer.

2

Related work

Our work connects to (i) computational social science (CSS) (ii) analysis of political data (parliamentary debates) and (iii) the emergent field of analysis of social solidarity using NLP approaches.

NLP-based CSS. Recent CSS studies have leveraged LLMs for a variety of complex tasks. Ziems et al. (2024) conduct a comprehensive evaluation of LLMs, pointing out their weaknesses in tasks which require understanding of subjective expert taxonomies that deviate from the training data of LLMs (such as implicit hate and empathy classification). LLMs enhance text-as-data methods in social sciences, particularly in analyzing political ideology (Wu et al., 2023), but struggle with social language understanding, often outperformed by fine-tuned models (Choi et al.,
2023). Zhang et al. (2023) introduced SPARROW, a benchmark showing ChatGPT’s limitations in sociopragmatic understanding across languages.
In exploring German migration debates, Blokker et al. (2020) and Zaberer et al. (2023) utilize finetuning of transformer-based language models to classify claims in German newspapers. Chen et al.
(2022) apply LLM-based classification on German social media posts to study public controversies over the course of one decade. In contrast to these approaches, we apply LLMs to longitudinal historical data and explore it for a new challenging task, fine-grained detection of social solidarity.
Analysis of parliamentary debates using NLP
tools. Abercrombie and Batista-Navarro (2020a)
review 61 studies on sentiment and position-taking within parliamentary contexts, covering dictionarybased sentiment scoring, statistical machine learning, and other conventional NLP methods. In terms of specific methodologies, studies often deploy: (i) shallow classifiers, where Lai et al. (2020)
use SVM, Naïve Bayes, and Logistic Regression for multilingual stance detection; (ii) deep learning approaches, with Abercrombie and BatistaNavarro (2020b) applying BERT, Al Hamoud et al.
(2022) exploring LSTM variants, and Sawhney et al. (2020) introducing GPolS for political speech analysis; (iii) probabilistic models, as in Vilares and He (2017)’s Bayesian approach to identify topics and perspectives in debates. With German political debates, Müller-Hansen et al. (2021) use topic modeling to study shifts in German parliamentary discussions on coal due to changes in energy policy, while Walter et al. (2021) employ diachronic word embeddings to track antisemitic and anti-communist biases in these debates. More recently, Bornheim et al. (2023) apply Llama 2 to automate speaker attribution in German parliamentary debates from 2017-2021. Our research goes

beyond this by adopting recent powerful LLMs to track changes of a specific social concept, solidarity, in plenary debates from three centuries.
Social solidarity in NLP. Previous studies of social solidarity in NLP have largely focused on social media platforms. For example, Santhanam et al. (2019) study how emojis are used to express solidarity in social media during Hurricane Irma in 2017 and Paris terrorist attacks from November
2015. Ils et al. (2021) consider solidarity in European social media discourse around COVID-19.
Eger et al. (2022) extend this work by examining how design choices, like keyword selection and language, affect assessments of solidarity changes over time. Compared to these works, we use a similar methodological setup (annotate data and infer trends), but focus on parliamentary debates instead of social media, employ a much more finegrained sociological framework (Thijssen, 2012), and use LLMs for systematic categorization and examination of solidarity types over time.

3

Data

We obtain data from two sources: (i) Open Data, covering Bundestag (en.: federal diet) protocols from 1949 until today; and (ii) Reichstagsprotokolle covering Reichstag (en.: imperial diet) protocols until 1945.1 We use the OCR-scanned version from Walter et al. (2021). Links to data, models, etc. used are in Appendix D. For the Reichstag data, we apply preprocessing steps similar to Walter et al. (2021) (e.g., removal of OCR artifacts), but keep German umlauts, capitalization, and punctuation. We automatically split the data into individual sittings and collect metadata like the date, period and number of each sitting, which we manually check and correct. Additionally, we removed interjections and split the text into sentences using
NLTK (Bird et al., 2009), resulting in 19.1M sentences. We release this dataset of plenary protocols from German political debates (DeuParl) consisting of 9,923 sittings from 1867 to 2022 on GitHub.2
To select keywords, we (i) train a Word2Vec model (Mikolov et al., 2013) on our dataset to identify words with vector representations similar to
Migrant (en.: migrant) and Frau (en.: woman); (ii)
manually expand this list with intuitively relevant terms; (iii) from both lists, we filter for those which
1
Volkskammer (en.: Eastern German parliament) protocols could not be included due to lack of availability.
2
https://github.com/DominikBeese/DeuParl-v2

appear at least 200 times in the dataset. This resulted in 32 keywords for Migrant and 18 keywords for Frau. These include general terms like Migrant,
Immigrant and Frau to period-specific terms, such as Vertriebene (en.: expellees) and Bürgerkriegsflüchtlinge (en.: civil war refugees), or social roles, such as Mütter (en.: mothers) and Hausfrauen
(en.: housewives). See the full list of keywords, and further preprocessing in Appendix A. For a detailed keyword distribution across the dataset, see
Fig. 11 and Fig. 12 in the Appendix.
Using these keywords, we extract 58k main sentences (instances) for migrants and 131k for women from DeuParl, expanding each with three preceding and three following sentences for context, resulting in a total of (i) 463k sentences
(9.79M tokens) for migrants and (ii) 1.58M sentences (32.82M tokens) for women. Fig. 2 shows the number of instances over time.3 Fig. 7 in the
Appendix shows yearly relative frequencies of sentences with terms related to women and migrants in the entire dataset. It is notable that both Frau and Migrant terms represent a minor fraction of the discourse, typically under 0.02. Periodic spikes in mentions likely align with historical and societal changes, such as post-WWII for Migrant.

 1 X P E H U  R I  6 H Q W H Q F H V

    

 : R P D Q
 0 L J U D Q W

    
    
   
 
    

    

    

    

    

 < H D U

    

    

    

    

Figure 2: Number of instances in the Woman and
Migrant dataset in each year. Fig. 7 in the Appendix illustrates the relative frequency of instances in both datasets.

4

Data annotation

To obtain ground truth data for model training and evaluation, we annotated 2864 instances with five annotators (all student assistants, with specializations in social science or computer science). The annotation was performed over a duration of nine months. In the first three months, we iteratively
3

4.1

Annotation task design

For the manual annotation, we take the target sentence and three preceding and following sentences for context into account. We first select a high-level category (solidarity, anti-solidarity, mixed, none).
Solidarity or anti-solidarity cases are then further distinguished into frames as defined by Thijssen and Verheyen (2022): group-based, compassionate, exchange-based, and empathic. We describe each of the included variables below.
High-level categories. Based on Lahusen and
Grasso (2018) and Ils et al. (2021), we define solidarity as willingness to share resources, directly or indirectly, or support for target groups, and antisolidarity as statements restricting resources, showing unwillingness to support, or implying exclusion of these groups. Texts with both supporting and opposing expressions are labeled mixed, while neutral or unrelated texts are labeled none.
Group-based solidarity is coded for texts emphasizing shared identity and common goals among group members, whereas group-based antisolidarity emphasizes out-group exclusion based on perceived differences.

 1 X P E H U  R I  6 H Q W H Q F H V  S H U  < H D U  D Q G  & D W H J R U \
    

refined the annotation guidelines and monitored the inter-rater agreement (measured by Cohen’s
Kappa) until inter-rater agreement converged (see
Section 4.2 for exact scores) and annotators began annotating independently.

We note that the dataset is sparse in the period from 1933
to 1949, i.e. during the NS dictatorship and the immediate after-war period until the first parliament after the war was elected in 1949.

Compassionate solidarity is coded for texts supporting marginalized groups, emphasizing their need for protection, while compassionate antisolidarity dismisses these groups by considering them already in a good position, minimizing their need support or protection.
Exchange-based solidarity is coded when texts highlight the economic contributions of “exchange partners” and potential rewards or further contributions. Conversely, exchange-based anti-solidarity calls for penalizing groups perceived as receiving more than they contribute.
Empathic solidarity is coded when a speaker expresses respect for individual differences, seeing social diversity as beneficial, while empathic anti-solidarity arises when differences are used as grounds for exclusion or neglect.
Annotation involved elaborate explanations, with identification of (anti-)solidarity resources

Annotation results
110

7

10

187

2

11

20

65

12

8

3

41

7

16

None

17

Soli Anti-soli Mixed

1

Human Label

3

12

7

10

6

38

21

3

5

49

None

Soli

(a) Annotators’ agreement

Anti-soli Mixed
Model Predictions

None

(b) Annotators vs. Model

Figure 3: Fig. 3a shows the confusion matrix between human annotators; Fig. 3b shows agreement between the best model and human annotators on a test set from one of the three splits. The former is aggregated over all pairwise comparisons of annotators, thus the matrix is symmetric.

While initial agreement levels were low, by the time annotators began working independently, they achieved a pairwise agreement with a Cohen’s
Kappa of 0.42 on a fine-grained level and 0.62
on a high level. We observe three main disagreement issues in annotation: misclassification of none cases, confusion between mixed stance and antisolidarity, and overlap within solidarity and antisolidarity subtypes (see Fig. 6 in the Appendix).
This confusion is often due to overlapping characteristics or the presence of multiple subtypes within the text; moreover, this annotation task is inherently subjective, which can lead to differing interpretations. This is further evidenced by our average agreement scores. Table 6 in the Appendix provides examples of annotator divergence, explaining why multiple labels could be correct, which gives insight into more difficult instances. However, there was almost no confusion between solidarity and anti-solidarity. We note less stability in annotator agreement before 1930, stabilizing in subsequent years (see Fig. 14a in the Appendix).
Although these variations can stem from the complexities of historical language and diverse interpretations of past events, they might also stem from the unbalanced distribution of human annotated data over the decades (see Fig. 4).

Our dataset comprises 2864 annotated instances,
1437 for migrants and 1427 for women. We note that anti-solidarity accounts for 13.5% of instances, being more common among migrants (12.1%) than women (1.4%) (see Table 5a in the Appendix). 368
instances in our dataset (referred to as curated)
were reviewed by a social science expert to provide a reliable comparison benchmark for evaluation of our models. Other consensus mechanisms for the final labels in the human-annotated dataset, and their distribution are shown in Table 5b in the Appendix.
Sentence Distribution
600

Number of Sentences

4.2
Mixed Anti-soli Soli

and highlighting expressions of (anti-)solidarity, as well as self- (speaker’s own viewpoint) and otherposition (addressing or criticizing others’ viewpoints). The full annotation process and a detailed example are described in Appendix C and Fig. 15
in the Appendix, respectively.

500
400
300

Woman (Solidarity)
Woman (Anti-soldiarity)
Woman (Mixed)
Migrant (Solidarity)
Migrant (Anti-solidarity)
Migrant (Mixed)

200
100
0

1875

1900

1925

1950

Decade

1975

2000

2025

Figure 4: Distribution of instances in the human annotated dataset across time and target groups. See Fig. 8
in the Appendix for the plots for each group separately; and Table 4 for the actual numbers of instances in the human annotated dataset.

5

Models and experiments

To determine the most cost-effective model
(both in terms of performance and costs) for our large-scale sociological analysis, we evaluate various models, including Llama-3-70B-Instruct, gpt-4-1106-preview, base and instructionfinetuned gpt-3.5-turbo-0125, across high-level and fine-grained (anti-)solidarity annotation tasks in achieving human-level performance. Once the quality of the models is assured, we apply the best performing model — GPT-4 — large-scale to determine trends in Section 7.
Data We use a 70/15/15 train/dev/test split for all
Migrant and Woman annotated data, which gives us approx. 2000 train, 400 dev and 430 test instances.
We ensure reliability of a test set by allocating approximately 40% curated and 45% majority (labels assigned when more than half of annotators agree on the same label) labels. We create 3 random data splits, and calculate performance metrics as the average score of the 3 runs on the test sets.4
4

These sets are fully used for training and evaluating baseline models; for inference-based experiments with Llama-3,

Metrics To evaluate our models, we report the
Macro F1 Score (Macro F1) to account for class imbalance. We also calculate the F1 Score for the classes individually. We report these metrics for both high-level and fine-grained tasks.
5.1

Models

Baseline For our baseline, we use a BERT-based pipeline with 110M parameters, comprising a highlevel category classifier and two subtype models for solidarity and anti-solidarity. All models share a similar architecture, processing inputs with a target (Frau or Migrant) and full text comprising the focus sentence and left and right context. We add a fully connected layer with softmax activation atop the pooled output of the BERT-based models, with
4 output units for each model. To counter class imbalance, minority classes are oversampled to parity with the majority class. We finetune for 20 epochs with a learning rate of 4e-4, a warmup ratio of 0.05, linear decay, AdamW optimizer (Loshchilov and
Hutter, 2017) and categorical crossentropy loss.
GPT-4 We design two prompts (one for each target group) that include several elements: (i) incorporating chain-of-thought reasoning (Wei et al.,
2022) in a few-shot setting by providing examples with desired reasoning and asking to think step by step in a zero-shot setting (Kojima et al., 2022);
(ii) providing precise definitions and insights derived from annotation discussions; (iii) introducing definitions and examples of potentially problematic labels (such as empathic solidarity and empathic anti-solidarity) earlier in the prompt and (iv)
implementing a two-step prompting strategy that initially categorizes texts at a high-level followed by detailed subtype classification (full prompts are provided in Fig. 16 and Fig. 17 in the Appendix).
Prompt-based fine-tuned GPT-3.5 Using the prompt identified for GPT-4’s fine-grained classification, we proceeded to fine-tune GPT-3.5 on instances sampled from our initial train set (114 for migrants; 109 for women5 ), ensuring a balanced distribution across labels. The fine-tuning dataset was structured with the system providing instructions for classifying texts into high-level categories and requesting further sub-categorization; the user
GPT-3.5, GPT-4, only test sets are used (also averaged on 3
runs).
5
The fine-tuning guide by OpenAI recommends using 50 to
100 examples for training: https://platform.openai.com/
docs/guides/fine-tuning/preparing-your-dataset

presenting texts; and the assistant providing classifications as per our two-step reasoning approach, along with explanations generated using GPT-4.
Llama-3 For an open-source model, we select
Llama-3-70B-Instruct with q6_k quantization
(Meta, 2024), tested in a zero-shot setting. We use the same prompt as detailed in Fig. 16 and Fig. 17
in the Appendix, but involving separate calls for high-level and subcategory classifications.
For all GPT experiments, we test the models in few- (where we include ten category examples in the prompt to demonstrate the desired reasoning and categorization setting) and zero-shot settings.

6

Results

Results on the test sets are shown in Table 2 and
Table 7 in the Appendix.
GPT-4 consistently outperforms other models in both fine-grained and high-level tasks for women and migrants. Interestingly, GPT-4 achieves similar performance in zero-shot and few-shot settings –
0.37 (0.60) and 0.37 (0.54) for women, and 0.42
(0.73) and 0.43 (0.63) for migrants, respectively, where the first number represents the fine-grained, and the number in parentheses the high-level score.
This might be attributed to the use of carefully crafted definitions, which eliminates the need for additional examples. The fine-tuned version of
GPT-3.5 demonstrates only marginal improvement over the base model, generally falling short of GPT4’s performance. Most importantly, GPT-4 leads in F1 scores across categories, effectively identifying both solidarity and anti-solidarity (0.65 for women and 0.87 for migrants in zero-shot settings, see Table 7 in the Appendix). However, it scores lowest in the mixed stance category, indicating challenges with more complex and ambiguous instances. Llama-3, on the other hand, outperforms
GPT-4 in this category for women and also leads in the none category for migrants by better identifying contradicting and subtle cues, likely benefiting from the two-step classification approach.
Conversely, BERT achieves lower F1 scores across all categories, performing comparably to the
GPT-3.5 versions on high-level tasks but struggling with fine-grained tasks and ambivalent categories like mixed stance, likely due to challenges from underrepresented labels such as anti-solidarity categories for women, even after over-sampling (see
Fig. 4, as well as Table 5a in the Appendix).

Model
W
M

GPT-3.5 base
0-shot few-shot

Llama-3-70B
0-shot

BERT

Human upper bound

0.37 (0.60) 0.37 (0.54) 0.15 (0.46) 0.12 (0.41)
0.42 (0.73) 0.43 (0.63) 0.19 (0.48) 0.27 (0.50)

0.24 (0.48)
0.27 (0.65)

0.13 (0.26)
0.24 (0.46)

0.48 (0.72)
0.56 (0.78)

0-shot

GPT-4
few-shot

Table 2: Comparative performance (macro F1) of models vs. human upper bound on combined high-level (in parentheses) and fine-grained tasks for both women (W) and migrants (M). Best scores for each target group are highlighted in bold. Macro F1 cores for GPT-3.5 fine-tuned, as well as F1 scores for each category are provided in Table 7 in the Appendix.

Overall, the human upper bound consistently outperforms all models. GPT-4 is the closest but still falls behind by 0.11 to 0.14 and 0.05 to 0.12 points in fine-grained and high-level tasks respectively for women and migrants. Llama-3-70B follows, with gaps of 0.24 to 0.29 and 0.24 to 0.13 points in finegrained and high-level tasks, while BERT shows the largest discrepancies. These results suggest there is a gap compared to human understanding, especially in more complex annotation tasks.
Error analysis For the error analysis, we compare the human annotations and zero-shot predictions of GPT-4 for both target groups on the test set, using the confusion matrices for high-level labels shown in Fig. 3, as well as for fine-grained level labels in Fig. 6 provided in the Appendix. We also consider explanations provided by GPT-4.
Observed errors align with those noted during human annotation, with solidarity and anti-solidarity rarely confused (1% of cases). Confusion primarily occurs between (anti-)solidarity subtypes and none, as well as mixed stance, as the model seems to look for stronger indications of solidarity despite instructions to consider even slight expressions of it (see examples 1 and 2 in Table 9).
There is also notable confusion between the solidarity subtypes, with the most frequent confusion between group-based and compassionate solidarity, likely because of the presence of multiple category characteristics within the texts (see example
3 in Table 9). There is also confusion between compassionate and empathic solidarity, where the model sometimes misinterprets empathic as merely relating to the term empathy, overlooking the full category definition, which involves respecting diversity, as in example 4 in Table 9.

7

Analysis

In this section, we analyze how the solidarity discourse in the German parliament developed over
155 years, using automatic, fine-grained annotations from the best-performing model, GPT-4 zeroshot. However, due to cost constraints, we limit the annotations to a) data concerning migrants, i.e., the target group for which our models achieved higher scores and b) a sample of 18,300 instances from the overall 58k instances concerning migrants. We draw the sample proportionally for the time spans in the original data. This selection includes all records with political party information (see Appendix B for details on political parties data extraction and list of parties included in the analysis).
The cost of this automatic annotation was around
500 Euro.
How does (anti-)solidarity change over time?
As shown in Fig. 5b, throughout the periods analyzed, solidarity consistently surpasses antisolidarity. From 1880 to 1910, solidarity increased from under 20% to 30%, driven by discussions about Eastern and Central European foreign workers’ rights in the context of industrialization and local demographic shifts (Schönwälder, 1999).
Following the NS regime, solidarity surges above 50%, aligning with the influx of Vertriebene
(en.: expellees), who were generally viewed positively in parliamentary debates (Fröhlich, 2023), as well as the arrival of Gastarbeiter (en.: guest workers) from the Mediterranean (Stalker, 2000).
The highest solidarity and low anti-solidarity in
1940s reflect limited discussions (due to sparse data during this period) skewed towards solidarity, primarily focusing on expellees.
Anti-solidarity towards migrants initially rose from about 5% to over 15% between 1870 and
1890, remaining stable until 1920, likely due to fears of “Polonization” from the influx of Polish workers, leading to restrictive policies against perceived threats to German identity (Triadafilopoulos,
2004). From 1960, anti-solidarity resurged with rising anti-migrant sentiments linked to labor migration in the 1960s and 1970s, right-wing opposition to liberal asylum laws in the 1990s (Faist,

80%
60%

group-based exchange-based compassionate empathic

60%
40%

20%

20%
1900

100%

80%

40%
0%

Solidarity Distribution per Decade (Migrant)

Solidarity vs. Anti-Solidarity Frames Distribution per Decade
Solidarity
Anti-Solidarity
100%

Percentage

Percentage of Label

100%

1950
Year

2000

0%

75%
50%
25%
0%

1900

1950
Year

2000

(a) Solidarity and anti-solidarity frames.

Solidarity
Anti-solidarity
Mixed

1900

1950
Year

2000

(b) Solidarity and anti-solidarity.

Figure 5: Fig. 5b shows the fraction of solidarity, anti-solidarity, and mixed stance towards migrants. Fig. 5b shows the fraction of solidarity (left) and anti-solidarity (right) subtypes according to GPT-4, where percentages represent the proportion of each subtype relative to the total counts of solidarity or anti-solidarity labels per decade. Grey shaded areas from 1933 to 1949 indicate sparse data during the NS dictatorship and post-war period.

1994), and the influx of refugees due to the Syrian war around 2015 with the subsequent rise of the extreme right-wing AfD party (Hertner, 2022). By
2022, solidarity stabilizes at around 40%, while anti-solidarity has risen to above 20%. We also note a relative decrease in solidarity compared to anti-solidarity, as shown in Fig. 10 in the Appendix, which illustrates this shift by subtracting the percentage of anti-solidarity from that of solidarity.
How have (anti-)solidarity frames evolved over time? In Fig. 5a, for solidarity (left), it is evident that group-based solidarity (i.e., emphasis on shared national identity and integration) was dominant until the 1990s, peaking to over 50% in
1870 with the founding of the German empire. It drops to below 20% by 1880, surges back above
50% in 1970, and then declines to below 30% by
2020. These trends align with periods of strong
German nationalism (around 1870 and pre-World
War II) and the influx of expellees in the 1950s and 1960s. For anti-solidarity (right), trends show group-based anti-solidarity dominant before WWII, peaking from 70% to over 90%, then declining to below 60% post-war, reflecting a decline in opposition to migration based on national identity in parliamentary debates after the NS era. Instead, anti-solidarity arguments shifted to exchange-based anti-solidarity (from below 10% to above 40% after
World War II), i.e. arguments that migrants are not providing adequate economic contributions. Neither compassionate nor empathic anti-solidarity are frequent at any time. By 2022, group-based solidarity declines, giving way to compassionate solidarity, which rose to above 40%.
How are solidarity and anti-solidarity frames represented across political parties? Our analysis covers the distribution of (anti-)solidarity frames across parties from 1865 to 2022. It should be noted that we assess the data without regard to the specific emergence dates of political parties
(e.g., AfD established in 2013 versus SPD in 1863).
Data subdivision by the speaker’s party shows that all parties, except for the far-right AfD, mainly exhibit solidarity towards migrants (Fig. 9 in the
Appendix, left). Compassionate and group-based solidarity are prevalent, with left-leaning parties
(Linke, Grüne) showing higher levels of compassionate and empathic solidarity compared to the more exchange-focused solidarity of centrist parties (SPD, FDP, CDU/CSU). Conversely, while the right-wing parties, including CDU/CSU, also engage in anti-solidarity rhetoric, it is AfD that predominantly holds such stances, focusing mainly on exchange-based anti-solidarity that suggests migrants contribute less (see Fig. 9 in the Appendix, right). This distribution mirrors findings in Flanders, Belgium, where radical rightists’ discourse predominantly focuses on group-based frames; greens and social democrats emphasize compassionate solidarity, liberals prefer exchange-based approaches, and both greens and liberals advocate for empathic solidarity, highlighting a polarization of partisan discourse (Thijssen and Verheyen,
2022).

8

Concluding remarks

This study set out to i) provide a high-quality dataset of (anti-)solidarity annotations in German parliamentary debates, ii) from an NLP perspective, evaluate the ability of LLMs to assist in largescale social and political analyses over extended periods, and iii) from a CSS perspective, uncover long-term shifts in solidarity trends in Germany,

laying groundwork for future sociological studies.
Concerning i), we invested 1000+ annotation hours and 18k Euro to provide a data set of 2,864
manually annotated text snippets following the sociological framework of Thijssen (2012).
Regarding ii), our findings indicate that GPT-4
outperforms other models (Llama-3-70B, GPT-3.5
base, fine-tuned, and BERT) in reproducing human annotations. While other models demonstrate proficiency in identifying high-level categories, they exhibit limitations in handling more nuanced categories, specific to sociological theory. GPT-4, though challenged by ambiguity, handles complex tasks even in a zero-shot setting.
Contrary to previous research in CSS that suggested smaller fine-tuned models like BERT performed well (Choi et al., 2023; Zhang et al., 2023), our study finds that larger models are more effective our hard CSS task, possibly due to the carefully designed prompts based on human expertise. Our observations align with Ziems et al.
(2024), who noted that while LLMs have not yet reached the quality of human analysis in classification tasks within CSS, large instruction-tuned
LLMs are preferable. However, we observe that the open-source model Llama-3 exhibits inferior performance compared to the non-open-source GPT-4.
In terms of iii), our analysis of German parliamentary debates from 1865 to 2022 reveals shifts in attitudes toward migrants influenced by labor demands, migration waves, and socio-political changes. The evolution from group-based to compassionate solidarity marks a shift towards economic pragmatism and issues of redistribution in migration discourse, likely spurred by rising global humanitarianism (Vollmer and Karakayali, 2018).
Despite traditional solidarity dominance, the resurgence of anti-solidarity since WWII, peaking in
2022, illustrates deepening political polarization.
This underscores the ongoing tension between xenophobic tendencies and liberal ideals in Germany (Joppke, 2007; Lehr, 2015), reflecting the complex challenges of integration and multiculturalism in a major migrant destination.
These findings lay a strong foundation for further sociological research using LLMs. Future work could examine how shifts in political discourse impact policy-making and public opinion, exploring correlations with civic solidarity. Comparative studies across countries with varying migration histories may reveal factors influencing (anti-)solidarity.
Additionally, exploring the effects of such rhetoric

on related policies, like migration or equal opportunity, could yield further insights. From an NLP perspective, future work can explore a broader range of LLMs, including those fine-tuned on German data, and explore the effect of using highlighting and explanations to improve classification.

Limitations
Our study faces several limitations that should be considered when interpreting the results.
• The task of annotating political speech, particularly concepts such as solidarity and antisolidarity, poses significant challenges. These concepts are inherently complex and laden with subtleties that are difficult to capture, both for human annotators and automated models.
• Due to resource constraints, GPT-4 was applied to only a portion of the dataset, potentially limiting the generalizability of our findings. Moreover, the proportional sampling of instances across decades might have led to the underrepresentation of certain periods, which could have further impacted the comprehensiveness of the analysis.
• While our analysis of solidarity and antisolidarity frames distribution across parties covers data from 1865 to 2022 and includes all historical periods, it does not trace the evolution of partisan discourse over time. We assess the data comprehensively without focusing on the specific historical emergence of political parties (e.g., AfD established in 2013
versus SPD in 1863).
• Additional manual annotation features such as free-comment explanations, highlighting, and indications of opposing positions (all available in our human annotation) were not fully explored. These elements hold potential for future studies, which could use them as cues for LLMs to achieve a more nuanced classification.

Acknowledgments
This research was funded by the Ministry of Culture and Science of the State of North RhineWestphalia under the grant no NW21-059A (SAIL).
The NLLG group gratefully acknowledges support from the Federal Ministry of Education and

Research (BMBF) via the research grant “Metrics4NLG” and the German Research Foundation
(DFG) via the Heisenberg Grant EG 375/5-1.

References
Gavin Abercrombie and Riza Batista-Navarro. 2020a.
Sentiment and position-taking analysis of parliamentary debates: a systematic literature review. Journal of Computational Social Science, 3(1):245–270.
Gavin Abercrombie and Riza Theresa Batista-Navarro.
2020b. Parlvote: A corpus for sentiment analysis of political debates. In Proceedings of the Twelfth Language Resources and Evaluation Conference, pages
5073–5078.
Kabir Ahuja, Harshita Diddee, Rishav Hada, Millicent Ochieng, Krithika Ramesh, Prachi Jain, Akshay Nambi, Tanuja Ganu, Sameer Segal, Maxamed
Axmed, et al. 2023. Mega: Multilingual evaluation of generative ai. arXiv preprint arXiv:2303.12528.
Ahmed Al Hamoud, Amber Hoenig, and Kaushik Roy.
2022. Sentence subjectivity analysis of a political and ideological debate dataset using lstm and bilstm with attention and gru models. Journal of King
Saud University-Computer and Information Sciences,
34(10):7974–7987.
Floya Anthias. 2014. Beyond integration: Intersectional issues of social solidarity and social hierarchy.
In Contesting integration, engendering migration: theory and practice, pages 13–36. Springer.
Keith Banting and Will Kymlicka. 2017. 1Introduction:
The Political Sources of Solidarity in Diverse Societies. In The Strains of Commitment: The Political
Sources of Solidarity in Diverse Societies. Oxford
University Press.
Steven Bird, Ewan Klein, and Edward Loper. 2009. Natural language processing with Python: analyzing text with the natural language toolkit. " O’Reilly Media,
Inc.".
Nico Blokker, Erenay Dayanik, Gabriella Lapesa, and
Sebastian Padó. 2020. Swimming with the tide? positional claim detection across political text types. In
Proceedings of the Fourth Workshop on Natural Language Processing and Computational Social Science, pages 24–34, Online. Association for Computational
Linguistics.
Tobias Bornheim, Niklas Grieger, Patrick Gustav Blaneck, and Stephan Bialonski. 2023. Speaker attribution in german parliamentary debates with qloraadapted large language models. arXiv preprint arXiv:2309.09902.
Marina Calloni. 2020. Women, minorities, populism.
Minorities and Populism–Critical Perspectives from
South Asia and Europe, pages 243–264.

Yiyi Chen, Harald Sack, and Mehwish Alam. 2022. Analyzing social media for measuring public attitudes toward controversies and their driving factors: a case study of migration. Soc. Netw. Anal. Min., 12(1):135.
Minje Choi, Jiaxin Pei, Sagar Kumar, Chang Shu, and
David Jurgens. 2023. Do llms understand social knowledge? evaluating the sociability of large language models with socket benchmark. arXiv preprint arXiv:2305.14938.
Bosheng Ding, Chengwei Qin, Linlin Liu, Yew Ken
Chia, Shafiq Joty, Boyang Li, and Lidong Bing. 2022.
Is gpt-3 a good data annotator? arXiv preprint arXiv:2212.10450.
Steffen Eger, Dan Liu, and Daniela Grunow. 2022. Measuring social solidarity during crisis: The role of design choices. Journal of Social Computing.
Thomas Faist. 1994. How to define a foreigner? the symbolic politics of immigration in german partisan discourse, 1978–1992. West European Politics,
17(2):50–71.
Christiane Fröhlich. 2023. Migration as crisis? german migration discourse at critical points of nationbuilding. American Behavioral Scientist, early access.
Isabelle Hertner. 2022. Germany as ‘a country of integration’? the cdu/csu’s policies and discourses on immigration during angela merkel’s chancellorship.
Journal of Ethnic and Migration Studies, 48(2):461–
481.
Marit Hopman and Trudie Knijn. 2022. Understanding solidarity in society: Triggers and barriers for in-and outgroup solidarity. In Solidarity and Social Justice in Contemporary Societies: An Interdisciplinary Approach to Understanding Inequalities, pages 29–39.
Springer.
Alexandra Ils, Dan Liu, Daniela Grunow, and Steffen
Eger. 2021. Changes in European solidarity before and during COVID-19: Evidence from a large crowdand expert-annotated Twitter dataset. In Proceedings of the 59th Annual Meeting of the Association for
Computational Linguistics and the 11th International
Joint Conference on Natural Language Processing
(Volume 1: Long Papers), pages 1623–1637, Online.
Association for Computational Linguistics.
Ronald F Inglehart and Pippa Norris. 2016. Trump, brexit, and the rise of populism: Economic have-nots and cultural backlash.
Christian Joppke. 2007. Beyond national models: Civic integration policies for immigrants in western europe.
West European Politics, 30(1):1–22.
Takeshi Kojima, Shixiang Shane Gu, Machel Reid, Yutaka Matsuo, and Yusuke Iwasawa. 2022. Large language models are zero-shot reasoners. Advances in neural information processing systems, 35:22199–
22213.

Will Kymlicka. 2020. Solidarity in diverse societies:
Beyond neoliberal multiculturalism and welfare chauvinism. Minorities and populism–critical perspectives from South Asia and Europe, pages 41–62.

Sashank Santhanam, Vidhushini Srinivasan, Shaina
Glass, and Samira Shaikh. 2019. I stand with you:
Using emojis to study solidarity in crisis events.
ArXiv, abs/1907.08326.

Christian Lahusen and Maria T. Grasso, editors. 2018.
Solidarity in Europe: Citizens’ Responses in Times of Crisis. Palgrave Studies in European Political
Sociology. Springer International Publishing.

Ramit Sawhney, Arnav Wadhwa, Shivam Agarwal, and
Rajiv Shah. 2020. Gpols: A contextual graph-based language model for analyzing parliamentary debates and political cohesion. In Proceedings of the 28th
International Conference on Computational Linguistics, pages 4847–4859.

Mirko Lai, Alessandra Teresa Cignarella, Delia
Irazú Hernández Farías, Cristina Bosco, Viviana Patti, and Paolo Rosso. 2020. Multilingual stance detection in social media political debates. Computer Speech
& Language, 63:101075.
Arto Laitinen and Anne Birgitta Pessi. 2014. Solidarity: Theory and practice. an introduction. Solidarity:
Theory and practice, pages 1–29.
Sabine Lehr. 2015. Germany as host: Examining ongoing antiimmigration discourse and policy in a country with a high level of non-national residents. Refugee review: Reconceptualizing refugees and forced migration in the 21st century, 2(1):113–31.
Chaoqun Liu, Wenxuan Zhang, Yiran Zhao, Anh Tuan
Luu, and Lidong Bing. 2024. Is translation all you need? a study on solving multilingual tasks with large language models. arXiv preprint arXiv:2403.10258.
Ilya Loshchilov and Frank Hutter. 2017. Decoupled weight decay regularization. arXiv preprint arXiv:1711.05101.
AI Meta. 2024. Introducing meta llama 3: The most capable openly available llm to date. Meta AI.
Tomas Mikolov, Kai Chen, Gregory S. Corrado, and
Jeffrey Dean. 2013. Efficient estimation of word representations in vector space. In ICLR.
Finn Müller-Hansen, Max W Callaghan, Yuan Ting Lee,
Anna Leipprand, Christian Flachsland, and Jan C
Minx. 2021. Who cares about coal? analyzing 70
years of german parliamentary debates on coal with dynamic topic modeling. Energy Research & Social
Science, 72:101869.
Stijn Oosterlynck and B Bouchaute. 2013. Social solidarities: the search for solidarity in sociology. OASeS
Research Centre.
Nicholas Pangakis, Samuel Wolken, and Neil Fasching.
2023. Automated annotation with generative ai requires validation. arXiv preprint arXiv:2306.00176.
Libo Qin, Qiguang Chen, Yuhang Zhou, Zhi Chen,
Yinghui Li, Lizi Liao, Min Li, Wanxiang Che, and
Philip S Yu. 2024. Multilingual large language model: A survey of resources, taxonomy and frontiers. arXiv preprint arXiv:2404.04925.
Paul Reynolds. 2014. Introduction. In Scott H Boyd and Mary Ann Walter, editors, Cultural Difference and Social Solidarity: Solidarities and Social Function, pages 1–14. Cambridge Scholars Publishing,
Newcastle upon Tyne.

Karen Schönwälder. 1999. Invited but unwanted? migrants from the east in germany, 1890–1990. In The
German Lands and Eastern Europe: Essays on the
History of their Social, Cultural and Political Relations, pages 198–216. Springer.
Settaluri Lakshmi Sravanthi, Meet Doshi, Tankala Pavan Kalyan, Rudra Murthy, Pushpak Bhattacharyya, and Raj Dabre. 2024. Pub: A pragmatics understanding benchmark for assessing llms’ pragmatics capabilities. arXiv preprint arXiv:2401.07078.
Peter Stalker. 2000. Workers without frontiers: the impact of globalization on international migration.
International Labour Organization.
Peter Thijssen. 2012. From mechanical to organic solidarity, and back: With honneth beyond durkheim.
European Journal of Social Theory, 15(4):454–470.
Peter Thijssen and Pieter Verheyen. 2022. It’s all about solidarity stupid! how solidarity frames structure the party political sphere. British Journal of Political
Science, 52(1):128–145.
Dietrich Thränhardt. 1993. Die ursprünge von rassismus und fremdenfeindlichkeit in der konkurrenzdemokratie: Ein vergleich der entwicklungen in england, frankreich und deutschland. Leviathan,
21(3):336–357.
Triadafilos Triadafilopoulos. 2004. Building walls, bounding nations: Migration and exclusion in canada and germany, 1870–1939. Journal of Historical Sociology, 17(4):385–427.
David Vilares and Yulan He. 2017. Detecting perspectives in political debates. In EMNLP 2017Conference on Empirical Methods in Natural Language Processing, pages 1573–1582. Association for
Computat

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

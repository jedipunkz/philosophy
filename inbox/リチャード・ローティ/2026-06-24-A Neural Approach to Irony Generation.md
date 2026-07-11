---
source: "https://arxiv.org/abs/1909.06200v2"
title: "A Neural Approach to Irony Generation"
author: "Mengdi Zhu, Zhiwei Yu, Xiaojun Wan"
year: "2019"
publication: "arXiv preprint / cs.CL"
download: "https://arxiv.org/pdf/1909.06200v2"
pdf: "https://arxiv.org/pdf/1909.06200v2"
captured_at: "2026-06-24T11:11:57Z"
updated_at: "2026-06-24T11:11:57Z"
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

# A Neural Approach to Irony Generation

- 著者: Mengdi Zhu, Zhiwei Yu, Xiaojun Wan
- 年: 2019
- 掲載情報: arXiv preprint / cs.CL
- 情報源: [arxiv](https://arxiv.org/abs/1909.06200v2)
- ダウンロード: https://arxiv.org/pdf/1909.06200v2
- PDF: https://arxiv.org/pdf/1909.06200v2

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Ironies can not only express stronger emotions but also show a sense of humor. With the development of social media, ironies are widely used in public. Although many prior research studies have been conducted in irony detection, few studies focus on irony generation. The main challenges for irony generation are the lack of large-scale irony dataset and difficulties in modeling the ironic pattern. In this work, we first systematically define irony generation based on style transfer task. To address the lack of data, we make use of twitter and build a large-scale dataset. We also design a combination of rewards for reinforcement learning to control the generation of ironic sentences. Experimental results demonstrate the effectiveness of our model in terms of irony accuracy, sentiment preservation, and content preservation.

## PDF Text

A N EURAL A PPROACH TO I RONY G ENERATION
A P REPRINT

arXiv:1909.06200v2 [cs.CL] 16 Sep 2019

Mengdi Zhu, Zhiwei Yu, Xiaojun Wan
Peking University
{1600012990,yuzw,wanxiaojun}@pku.edu.cn

September 17, 2019

A BSTRACT
Ironies can not only express stronger emotions but also show a sense of humor. With the development of social media, ironies are widely used in public. Although many prior research studies have been conducted in irony detection, few studies focus on irony generation. The main challenges for irony generation are the lack of large-scale irony dataset and difficulties in modeling the ironic pattern. In this work, we first systematically define irony generation based on style transfer task. To address the lack of data, we make use of twitter and build a large-scale dataset. We also design a combination of rewards for reinforcement learning to control the generation of ironic sentences. Experimental results demonstrate the effectiveness of our model in terms of irony accuracy, sentiment preservation, and content preservation1 .

1

Introduction

The irony is a kind of figurative language, which is widely used on social media [1]. The irony is defined as a clash between the intended meaning of a sentence and its literal meaning [2]. As an important aspect of language, irony plays an essential role in sentiment analysis [3, 1] and opinion mining [4, 5].
Although some previous studies focus on irony detection, little attention is paid to irony generation. As ironies can strengthen sentiments and express stronger emotions, we mainly focus on generating ironic sentences. Given a non-ironic sentence, we implement a neural network to transfer it to an ironic sentence and constrain the sentiment polarity of the two sentences to be the same. For example, the input is “I hate it when my plans get ruined" which is negative in sentiment polarity and the output should be ironic and negative in sentiment as well, such as “I like it when my plans get ruined". The speaker uses “like" to be ironic and express his or her negative sentiment. At the same time, our model can preserve contents which are irrelevant to sentiment polarity and irony. According to the categories mentioned in [6], irony can be classified into 3 classes: verbal irony by means of a polarity contrast, the sentences containing expression whose polarity is inverted between the intended and the literal evaluation; other types of verbal irony, the sentences that show no polarity contrast between the literal and intended meaning but are still ironic; and situational irony, the sentences that describe situations that fail to meet some expectations. As ironies in the latter two categories are obscure and hard to understand, we decide to only focus on ironies in the first category in this work. For example, our work can be specifically described as: given a sentence “I hate to be ignored", we train our model to generate an ironic sentence such as “I love to be ignored". Although there is “love" in the generated sentence, the speaker still expresses his or her negative sentiment by irony. We also make some explorations in the transformation from ironic sentences to non-ironic sentences at the end of our work. Because of the lack of previous work and baselines on irony generation, we implement our model based on style transfer. Our work will not only provide the first large-scale irony dataset but also make our model as a benchmark for the irony generation.
Recently, unsupervised style transfer becomes a very popular topic. Many state-of-the-art studies try to solve the task with sequence-to-sequence (seq2seq) framework. There are three main ways to build up models. The first is to learn a latent style-independent content representation and generate sentences with the content representation and another style [7, 8]. The second is to directly transfer sentences from one style to another under the control of classifiers
1

Our data and code are released at https://github.com/zmd971202/IronyGeneration.

A PREPRINT - S EPTEMBER 17, 2019

and reinforcement learning [9]. The third is to remove style attribute words from the input sentence and combine the remaining content with new style attribute words [10, 11]. The first method usually obtains better performances via adversarial training with discriminators. The style-independent content representation, nevertheless, is not easily obtained [12], which results in poor performances. The second method is suitable for complex styles which are difficult to model and describe. The model can learn the deep semantic features by itself but sometimes the model is sensitive to parameters and hard to train. The third method succeeds to preserve content but cannot work for some complex styles such as democratic and republican. Sentences with those styles usually do not have specific style attribute words.
Unfortunately, due to the lack of large irony dataset and difficulties of modeling ironies, there has been little work trying to generate ironies based on seq2seq framework as far as we know. Inspired by methods for style transfer, we decide to implement a specifically designed model based on unsupervised style transfer to explore irony generation.
In this paper, in order to address the lack of irony data, we first crawl over 2M tweets from twitter to build a dataset with 262,755 ironic and 112,330 non-ironic tweets. Then, due to the lack of parallel data, we propose a novel model to transfer non-ironic sentences to ironic sentences in an unsupervised way. As ironic style is hard to model and describe, we implement our model with the control of classifiers and reinforcement learning. Different from other studies in style transfer, the transformation from non-ironic to ironic sentences has to preserve sentiment polarity as mentioned above.
Therefore, we not only design an irony reward to control the irony accuracy and implement denoising auto-encoder and back-translation to control content preservation but also design a sentiment reward to control sentiment preservation.
Experimental results demonstrate that our model achieves a high irony accuracy with well-preserved sentiment and content. The contributions of our work are as follows:
• To our knowledge, our work is the first attempt to specifically define irony generation and generate ironic sentences via seq2seq neural network.
• We build the first large-scale irony dataset with 262,755 ironic and 112,330 non-ironic tweets.
• We implement well-designed rewards for irony accuracy and sentiment preservation which are different from those in previous work on style transfer
• Our approach yields substantial results on generating ironic sentences with high irony accuracy and wellpreserved content and sentiment.

2

Related Work

Style Transfer: As irony is a complicated style and hard to model with some specific style attribute words, we mainly focus on studies without editing style attribute words.
Some studies are trying to disentangle style representation from content representation. In [13], authors leverage adversarial networks to learn separate content representations and style representations. In [14] and [7], researchers combine variational auto-encoders (VAEs) with style discriminators.
However, some recent studies [12] reveal that the disentanglement of content and style representations may not be achieved in practice. Therefore, some other research studies [10, 11] strive to separate content and style by removing stylistic words. Nonetheless, many non-ironic sentences do not have specific stylistic words and as a result, we find it difficult to transfer non-ironic sentences to ironic sentences through this way in practice.
Besides, some other research studies do not disentangle style from content but directly learn representations of sentences.
In [9], authors propose a dual reinforcement learning framework without separating content and style representations.
In [8], researchers utilize a machine translation model to learn a sentence representation preserving the meaning of the sentence but reducing stylistic properties. In this method, the quality of generated sentences relies on the performance of classifiers to a large extent. Meanwhile, such models are usually sensitive to parameters and difficult to train. In contrast, we combine a pre-training process with reinforcement learning to build up a stable language model and design special rewards for our task.
Irony Detection: With the development of social media, irony detection becomes a more important task. Methods for irony detection can be mainly divided into two categories: methods based on feature engineering and methods based on neural networks.
As for methods based on feature engineering, In [2], authors investigate pragmatic phenomena and various irony markers. In [15], researchers leverage a combination of sentiment, distributional semantic and text surface features.
Those models rely on hand-crafted features and are hard to implement.
When it comes to methods based on neural networks, long short-term memory (LSTM) [16] network is widely used and is very efficient for irony detection. In [17], a tweet is divided into two segments and a subtract layer is implemented to
2

A PREPRINT - S EPTEMBER 17, 2019

Irony
Non-irony

Table 1: Samples in our dataset.
Examples i love this time of year , when all your teacher throw in one more big project before finals .
i just love it when my plans get ruined .
could be my last day in collingwood place . i am so sad .
i hate my roommates for making drink on a tuesday .

calculate the difference between two segments in order to determine whether the tweet is ironic. In [18], authors utilize a recurrent neural network with Bi-LSTM and self-attention without hand-crafted features. In [19], researchers propose a system based on a densely connected LSTM network.

3

Our Dataset

In this section, we describe how we build our dataset with tweets. First, we crawl over 2M tweets from twitter2 using
GetOldTweets-python3 . We crawl English tweets from 04/09/2012 to /12/18/2018. We first remove all re-tweets and use langdetect4 to remove all non-English sentences. Then, we remove hashtags attached at the end of the tweets because they are usually not parts of sentences and will confuse our language model. After that, we utilize Ekphrasis5 to process tweets. We remove URLs and restore remaining hashtags, elongated words, repeated words, and all-capitalized words. To simplify our dataset, We replace all “<money>" and “<time>" tokens with “<number>" token when using Ekphrasis. And we delete sentences whose lengths are less than 10 or greater than 40. In order to restore abbreviations, we download an abbreviation dictionary from webopedia6 and restore abbreviations to normal words or phrases according to the dictionary. Finally, we remove sentences which have more than two rare words (appearing less than three times) in order to constrain the size of vocabulary. Finally, we get 662,530 sentences after pre-processing.
As neural networks are proved effective in irony detection, we decide to implement a neural classifier in order to classify the sentences into ironic and non-ironic sentences. However, the only high-quality irony dataset we can obtain is the dataset of Semeval-2018 Task 3 and the dataset is pretty small, which will cause overfitting to complex models.
Therefore, we just implement a simple one-layer RNN with LSTM cell to classify pre-processed sentences into ironic sentences and non-ironic sentences because LSTM networks are widely used in irony detection. We train the model with the dataset of Semeval-2018 Task 3. After classification, we get 262,755 ironic sentences and 399,775 non-ironic sentences. According to our observation, not all non-ironic sentences are suitable to be transferred into ironic sentences.
For example, “just hanging out . watching . is it monday yet" is hard to transfer because it does not have an explicit sentiment polarity. So we remove all interrogative sentences from the non-ironic sentences and only obtain the sentences which have words expressing strong sentiments. We evaluate the sentiment polarity of each word with TextBlob7 and we view those words with sentiment scores greater than 0.5 or less than -0.5 as words expressing strong sentiments.
Finally, we build our irony dataset with 262,755 ironic sentences and 102,330 non-ironic sentences.

4

Our Method

Given two non-parallel corpora: non-ironic corpus N={n1 , n2 , ..., nn } and ironic corpus I={i1 , i2 , ..., im }, the goal of our irony generation model is to generate an ironic sentence from a non-ironic sentence while preserving the content and sentiment polarity of the source input sentence. We implement an encoder-decoder framework where two encoders are utilized to encode ironic sentences and non-ironic sentences respectively and two decoders are utilized to decode ironic sentences and non-ironic sentences from latent representations respectively. In order to enforce a shared latent space, we share two layers on both the encoder side and the decoder side. Our model architecture is illustrated in Figure
1. We denote irony encoder as Ei , irony decoder as Di and non-irony encoder as En , non-irony decoder as Dn . Their parameters are θEi , θDi , θEn and θDn .
Our irony generation algorithm is shown in Algorithm 1. We first pre-train our model using denoising auto-encoder and back-translation to build up language models for both styles (section 4.1). Then we implement reinforcement
2

https://twitter.com/.
https://github.com/Jefferson-Henrique/GetOldTweets-python.
4
https://github.com/Mimino666/langdetect.
5
https://github.com/cbaziotis/ekphrasis.
6
https://www.webopedia.com/quick_ref/ textmessageabbreviations.asp.
7
https://github.com/sloria/TextBlob.
3

3

A PREPRINT - S EPTEMBER 17, 2019

𝑖𝑖

Irony Cls

𝑝(𝑠|𝑖𝑖 ; 𝜑)

𝑅𝑊𝑖𝑟𝑜𝑛𝑦

𝑝(𝑠|𝑛෥𝑖 ; 𝜑)

Irony Cls

Irony
Senti Cls

𝑝(𝑡|𝑖𝑖 ; 𝜙)

𝑅𝑊𝑠𝑒𝑛𝑡𝑖

𝑝(𝑡|𝑛෥𝑖 ; 𝜓)

NonIrony
Senti Cls

Irony
Encoder

Shared
Encoder

𝑛𝑖

𝑖෡𝑖

𝑛ෝ𝑖

Non-Irony
Decoder

𝑛෥𝑖

Irony
Decoder

𝑖෩𝑖

Shared
Decoder

Non-Irony
Encoder

NonIrony
Senti Cls

𝑝(𝑠|𝑛𝑖 ; 𝜑)

𝑅𝑊𝑖𝑟𝑜𝑛𝑦

𝑝(𝑠|𝑖෩𝑖 ; 𝜑)

Irony
Senti Cls

Irony Cls

𝑝(𝑡|𝑛𝑖 ; 𝜓)

𝑅𝑊𝑠𝑒𝑛𝑡𝑖

𝑝(𝑡|𝑖෩𝑖 ; 𝜙)

Irony Cls

Figure 1: Framework of our model. The solid lines denote the transformation from non-ironic sentences to ironic sentences and the dotted lines denote the transformation from ironic sentences to non-ironic sentences.
learning to train the model to transfer sentences from one style to another (section 4.2). Meanwhile, to achieve content preservation, we utilize back-translation for one time in every p time steps.
4.1
4.1.1

Pretraining
Denoising Auto-Encoder

In order to build up our language model and preserve the content, we apply the auto-encoder model. To prevent the model from simply copying the input sentence, we randomly add some noises in the input sentence. Specifically, for every word in the input sentence, there is 10% chance that we delete it, 10 % chance that we duplicate it, 10% chance that we swap it with the next word, or it remains unchanged. We first encode the input sentence ni or ii with respective encoder En or Ei to obtain its latent representation n̂i = En (ni ) or iˆi = Ei (ii ) and reconstruct the input sentence with the latent representation and respective decoder. So we can get the reconstruction loss for auto-encoder Lossae :

4.1.2

Lossnae = En∼N [− log pDn (n|n̂i ; θDn )]

(1)

Lossiae = Ei∼I [− log pDi (i|iˆi ; θDi )]

(2)

Back Translation

In addition to denoising auto-encoder, we implement back-translation [20] to generate a pseudo-parallel corpus.
Suppose our model takes non-ironic sentence ni as input. We first encode ni with En to obtain its latent representation n̂i = En (ni ) and decode the latent representation with Di to get a transferred sentence iei . Then we encode iei with Ei and decode its latent representation with Dn to reconstruct the original input sentence ni . Therefore, our reconstruction loss for back-translation Lossbt :
Lossnbt = En∼N [− log pDn (n|Ei (iei ); θDn )]

(3)

And if our model takes ironic sentence ii as input, we can get the reconstruction loss for back-translation as:
Lossibt = Ei∼I [− log pDi (i|En (nei ); θDi )]
4

(4)

A PREPRINT - S EPTEMBER 17, 2019

Algorithm 1 Irony Generation Algorithm
. pre-train with auto-encoder
Pre-train En , Dn with N using MLE based on Eq.1
Pre-train Ei , Di with I using MLE based on Eq.2
. pre-train with back-translation
Pre-train En , Dn , Ei , Di with N using MLE based on Eq.3
Pre-train En , Dn , Ei , Di with I using MLE based on Eq.4
. train with RL
for each epoch e = 1, 2, ..., M do
. train non-irony2irony with RL
for ni in N do iei = Di (En (ni ))
update En , Di , using LossRL based on Eq.8
. back-translation if i%p = 0 then iei = Di (En (ni ))
nei = Dn (Ei (iei ))
update En , Dn , Ei , Di using MLE based on Eq.3
end if end for
. train irony2non-irony with RL
for ii in I do nei = Dn (Ei (ii ))
update Ei , Dn , using LossRL similar to Eq.8
. back-translation if i%p = 0 then nei = Dn (Ei (ii ))
iei = Di (En (nei ))
update En , Dn , Ei , Di using MLE based on Eq.4
end if end for end for

4.2

Reinforcement Learning

Since the gold transferred result of input is unavailable, we cannot evaluate the quality of the generated sentence directly.
Therefore, we implement reinforcement learning and elaborately design two rewards to describe the irony accuracy and sentiment preservation, respectively.
4.2.1

Irony Reward

A pre-trained binary irony classifier based on CNN [21] is used to evaluate how ironic a sentence is. We denote the parameter of the classifier as ϕ and it is fixed during the training process.
In order to facilitate the transformation, we design the irony reward as the difference between the irony score of the input sentence and that of the output sentence. Formally, when we input a non-ironic sentence ni and transfer it to an ironic sentence iei , our irony reward is defined as:
RWirony = p(s|iei ; ϕ) − p(s|ni ; ϕ)

(5)

where s denotes ironic style and p(s|x; ϕ) is the probability of that a sentence x is ironic.
4.2.2

Sentiment Reward

To preserve the sentiment polarity of the input sentence, we also need to use classifiers to evaluate the sentiment polarity of the sentences. However, the sentiment analysis of ironic sentences and non-ironic sentences are different. In the case of figurative languages such as irony, sarcasm or metaphor, the sentiment polarity of the literal meaning may differ significantly from that of the intended figurative meaning [1]. As we aim to train our model to transfer sentences from non-ironic to ironic, using only one classifier is not enough. As a result, we implement two pre-trained sentiment classifiers for non-ironic sentences and ironic sentences respectively. We denote the parameter of the sentiment classifier for ironic sentences as φ and that of the sentiment classifier for non-ironic sentences as ψ.
5

A PREPRINT - S EPTEMBER 17, 2019

A challenge, when we implement two classifiers to evaluate the sentiment polarity, is that the two classifiers trained with different datasets may have different distributions of scores. That means we cannot directly calculate the sentiment reward with scores applied by two classifiers. To alleviate this problem and standardize the prediction results of two classifiers, we set a threshold for each classifier and subtract the respective threshold from scores applied by the classifier to obtain the comparative sentiment polarity score. We get the optimal threshold by maximizing the ability of the classifier according to the distribution of our training data.
We denote the threshold of ironic sentiment classifier as thi and the threshold of non-ironic sentiment classifier as thn . The standardized sentiment score is defined as ST D(p(t|ni ; ψ)) = p(t|ni ; ψ) − thn and ST D(p(t|ii ; φ)) =
p(t|ii ; φ) − thi where t denotes the positive sentiment polarity and p(t|x; ·) is the probability of that a sentence is positive in sentiment polarity.
As mentioned above, the input sentence and the generated sentence should express the same sentiment. For example, if we input a non-ironic sentence “I hate to be ignored" which is negative in sentiment polarity, the generated ironic sentence should be also negative, such as “I love to be ignored". To achieve sentiment preservation, we design the sentiment reward as that one minus the absolute value of the difference between the standardized sentiment score of the input sentence and that of the generated sentence. Formally, when we input a non-ironic sentence ni and transfer it to an ironic sentence iei , our sentiment reward is defined as:
RWsenti = 1 − abs(ST D(p(t|ni ; ψ)) − ST D(p(t|iei ; φ)))
4.2.3

(6)

Overall Reward

To encourage our model to focus on both the irony accuracy and the sentiment preservation, we apply the harmonic mean of irony reward and sentiment reward:
RW = (1 + β 2 )
4.3

RWsenti · RWirony
2
(β · RWsenti ) + RWirony

(7)

Policy Gradient

The policy gradient algorithm [22] is a simple but widely-used algorithm in reinforcement learning. It is used to maximize the expected reward E[RW ]. The objective function to minimize is defined as:
K

LossRL = −

1 X
RWi · p(iei |ni ; θEn , θDi )
K i=1

(8)

where iei = Di (En (ni )), RWi is the reward of iei and K is the input size.

5

Experiments

5.1

Training Details

Ex , Ey , Dx and Dy in our model are Transformers [23] with 4 layers and 2 shared layers. The word embeddings of
128 dimensions are learned during the training process. Our maximum sentence length is set as 40. The optimizer is
Adam [24] and the learning rate is 10−5 . The batch size is 32 and harmonic weight β in Eq.9 is 0.5. We set the interval p as 200. The model is pre-trained for 6 epochs and trained for 15 epochs for reinforcement learning.
• Irony Classifier: We implement a CNN classifier trained with our irony dataset. All the CNN classifiers we utilize in this paper use the same parameters as [21].
• Sentiment Classifier for Irony: We first implement a one-layer LSTM network to classify ironic sentences in our dataset into positive and negative ironies. The LSTM network is trained with the dataset of Semeval 2015
Task 11 [1] which is used for the sentiment analysis of figurative language in twitter. Then, we use the positive ironies and negative ironies to train the CNN sentiment classifier for irony.
• Sentiment Classifier for Non-irony: Similar to the training process of the sentiment classifier for irony, we first implement a one-layer LSTM network trained with the dataset for the sentiment analysis of common twitters8 to classify the non-ironies into positive and negative non-ironies. Then we use the positive and negative non-ironies to train the sentiment classifier for non-irony.
8

https://www.kaggle.com/c/twitter-sentiment-analysis2/data.

6

A PREPRINT - S EPTEMBER 17, 2019

Table 2: Automatic evaluation results of the transformation from non-ironic sentences to ironic sentences.*
Senti Delta↓

Senti ACC↑

BLEU↑

G2↑

H2↑

BackTrans
Unpaired
CrossAlign
CPTG
DualRL

0.5417
0.5454
0.5357
0.5174
0.5167

48.83
49.26
49.56
49.43
49.73

1.80
18.78
2.77
0.26
76.38

9.38
30.41
11.72
3.58
61.63

3.47
27.19
5.25
0.52
60.24

Ours

0.5148

49.68

61.78

55.40

55.07

Model

*

We use ↓ to denote that the smaller value is better and ↑ to denote that the larger value is better.

Table 3: Human evaluation results of the transformation from non-ironic sentences to ironic sentences.
Model

5.2

Irony↓

Senti↓

Content↓

BackTrans
Unpaired
CrossAlign
CPTG
DualRL

2.98
4.35
3.72
4.65
2.91

3.86
4.26
4.26
4.84
1.28

3.91
4.05
4.30
5.00
1.35

Ours

2.40

2.51

2.40

Baselines

We compare our model with the following state-of-art generative models:
BackTrans[8]: In [8], authors propose a model using machine translation in order to preserve the meaning of the sentence while reducing stylistic properties.
Unpaired[11]: In [11], researchers implement a method to remove emotional words and add desired sentiment controlled by reinforcement learning.
CrossAlign[7]: In [7], authors leverage refined alignment of latent representations to perform style transfer and a cross-aligned auto-encoder is implemented.
CPTG[25]: An interpolated reconstruction loss is introduced in [25] and a discriminator is implemented to control attributes in this work.
DualRL[9]: In [9], researchers use two reinforcement rewards simultaneously to control style accuracy and content preservation.
5.3

Evaluation Metrics

5.3.1

Automatic Evaluation

In order to evaluate sentiment preservation, we use the absolute value of the difference between the standardized sentiment score of the input sentence and that of the generated sentence9 . We call the value as sentiment delta (senti delta). Besides, we report the sentiment accuracy (Senti ACC) which measures whether the output sentence has the same sentiment polarity as the input sentence based on our standardized sentiment classifiers. The BLEU score [26]
between the input sentences and the output sentences is calculated to evaluate the content preservation performance. In order to evaluate the overall performance of different models, we also report the geometric mean (G2) and harmonic mean (H2) of the sentiment accuracy and the BLEU score. As for the irony accuracy, we only report it in human evaluation results because it is more accurate for the human to evaluate the quality of irony as it is very complicated.
5.3.2

Human Evaluation

We first sample 50 non-ironic input sentences and their corresponding output sentences of different models. Then, we ask four annotators who are proficient in English to evaluate the qualities of the generated sentences of different models. They are required to rank the output sentences of our model and baselines from the best to the worst in terms
9

See the definition in Eq. 6. As sentiment reward, we use one to minus the absolute value and now we only use the absolute value as the metric.

7

A PREPRINT - S EPTEMBER 17, 2019

Table 4: Example outputs of the transformation from non-ironic sentences to ironic sentences and the transformation from ironic sentences to non-ironic sentences. We use red and blue to annotate the clashes in the sentences.
Example 1 (From non-irony to irony)

Example 2 (From non-irony to irony)

source

tried to leave town and my phone died , definition of success .

not even sleepy . tomorrows going to be fun .

BackTrans
Unpaired
CrossAlign
CPTG
DualRL

getting up at <number> in the morning and my life is going to be so much fun .
tried to leave town and my phone died , plowing split xx .
going to work and my house and <number> minutes , very little <UNK> on <user> , & you like cold tower getting good not ppl know .
tried to leave town and my phone died , definition of success .

not being so much fun . i am gonna be a great day .
no gets cool .
can not wait to go . this is going to work .
not people but <user> if you are binge - of way tell he " on festivities right .
not even sleepy . tomorrows going to be fun .

Ours

nice to leave town and my phone died , definition of success .

not even sleepy . depressed going to be fun .

Table 5: Example error outputs of the transformation from non-ironic sentences to ironic sentences. The main errors are colored.
source (non-irony)

output (irony)

No Change

cool thanks for even coming to say hi to me .

cool thanks for even coming to say hi to me .

Word Repetition

not able to extract data from tabs using python path

not able to accountable accountable from accountable using accountable accountable

Improper Words

friday night and here i am playing fifa . wonderful

feel night and here i am playing just . wonderful

of irony accuracy (Irony), Sentiment preservation (Senti) and content preservation (Content). The best output is ranked with 1 and the worst output is ranked with 6. That means that the smaller our human evaluation value is, the better the corresponding model is.
5.4

Results and Discussions

Table 2 shows the automatic evaluation results of the models in the transformation from non-ironic sentences to ironic sentences. From the results, our model obtains the best result in sentiment delta. The DualRL model achieves the highest result in other metrics, but most of its outputs are the almost same as the input sentences. So it is reasonable that
DualRL system outperforms ours in these metrics but it actually does not transfer the non-ironic sentences to ironic sentences at all. From this perspective, we cannot view DualRL as an effective model for irony generation. In contrast, our model gets results close to those of DualRL and obtains a balance between irony accuracy, sentiment preservation, and content preservation if we also consider the irony accuracy discussed below.
And from human evaluation results shown in Table 3, our model gets the best average rank in irony accuracy. And as mentioned above, the DualRL model usually does not change the input sentence and outputs the same sentence.
Therefore, it is reasonable that it obtains the best rank in sentiment and content preservation and ours is the second.
However, it still demonstrates that our model, instead of changing nothing, transfers the style of the input sentence with content and sentiment preservation at the same time.
5.5

Case Study

In the section, we present some example outputs of different models. Table 4 shows the results of the transformation from non-ironic sentences to ironic sentences. We can observe that: (1) The BackTrans system, the Unpaired system, the CrossAlign system and the CPTG system tends to generate sentences which are towards irony but do not preserve content. (2) The DualRL system preserves content and sentiment very well but even does not change the input sentence.
(3) Our model considers both aspects and achieves a better balance among irony accuracy, sentiment and content preservation.
5.6

Error Analysis

Although our model outperforms other style transfer baselines according to automatic and human evaluation results, there are still some failure cases because irony generation is still a very challenging task. We would like to share the issues we meet during our experiments and our solutions to some of them in this section.
• No Change: As mentioned above, many style transfer models, such as DualRL, tend to make few changes to the input sentence and output the same sentence. Actually, this is a common issue for unsupervised style transfer systems and we also meet it during our experiments. The main reason for the issue is that rewards for content preservation are too prominent and rewards for style accuracy cannot work well. In contrast, in
8

A PREPRINT - S EPTEMBER 17, 2019

Table 6: Automatic evaluation results of the transformation from ironic sentences to non-ironic sentences.
Senti Delta↓

Senti ACC↑

BLEU↑

G2↑

H2↑

BackTrans
Unpaired
CrossAlign
CPTG
DualRL

0.6927
0.5169
0.5710
0.5172
0.5534

40.87
49.64
46.77
48.94
47.82

1.98
9.28
4.85
0.49
74.31

9.00
21.46
15.06
4.90
59.61

3.78
15.64
8.79
0.97
58.19

Ours

0.4976

49.09

62.92

57.33

56.64

Model

Table 7: Human evaluation results of the transformation from ironic sentences to non-ironic sentences. Note that the larger value for irony metric is better here as we generate non-ironic sentences.
Model

Irony↑

Senti↓

Content↓

BackTrans
Unpaired
CrossAlign
CPTG
DualRL

3.46
4.00
3.68
4.86
1.93

3.73
3.94
3.99
5.07
1.60

3.94
3.82
4.27
5.00
1.43

Ours

3.07

2.67

2.54

order to guarantee the readability and fluency of the output sentence, we also cannot emphasize too much on rewards for style accuracy because it may cause some other issues such as word repetition mentioned below.
A method to solve the problem is tuning hyperparameters and this is also the method we implement in this work. As for content preservation, maybe MLE methods such as back-translation are not enough because they tend to force models to generate specific words. In the future, we should further design some more suitable methods to control content preservation for models without disentangling style and content representations, such as DualRL and ours.
• Word Repetition: During our experiments, we observe that some of the outputs prefer to repeat the same word as shown in Table 5. This is because reinforcement learning rewards encourage the model to generate words which can get high scores from classifiers and even back-translation cannot stop it. Our solution is that we can lower the probability of decoding a word in decoders if the word has been generated in the previous time steps during testing. We also try to implement this method during training time but obtain worse performances because it may limit the effects of training. Some previous studies utilize language models to control the fluency of the output sentence and we also try this method. Nonetheless, pre-training a language model with tweets and using it to generate rewards is difficult because tweets are more casual and have more noise.
Rewards from that kind of language model are usually not accurate and may confuse the model. In the future, we should come up with better methods to model language fluency with the consideration of irony accuracy, sentiment and content preservation, especially for tweets.
• Improper Words: As ironic style is hard for our model to learn, it may generate some improper words which make the sentence strange. As the example shown in the Table 5, the sentiment word in the input sentence is
“wonderful" and the model should change it into a negative word such as “sad" to make the output sentence ironic. However, the model changes “friday" and “fifa" which are not related to ironic styles. We have not found a very effective method to address this issue and maybe we should further explore stronger models to learn ironic styles better.
5.7

Additional Experiments

In this section, we describe some additional experiments on the transformation from ironic sentences to non-ironic sentences. Sometimes ironies are hard to understand and may cause misunderstanding, for which our task also explores the transformation from ironic sentences to non-ironic sentences.
As shown in Table 6, we also conduct automatic evaluations and the conclusions are similar to those of the transformation from non-ironic sentences to ironic sentences. As for human evaluation results in Table 7, our model still can achieve the second-best results in sentiment and content preservation. Nevertheless, DualRL system and ours get poor performances
9

A PREPRINT - S EPTEMBER 17, 2019

Table 8: Example outputs of the transformation from ironic sentences to non-ironic sentences. We use red and blue to annotate the clashes in the sentences.
From ironic to non-ironic source

wow i love feeling like the biggest dumbass on earth .

BackTrans
Unpaired
CrossAlign
CPTG
DualRL

<user> i am sure you are going to be a good day .
i am good like always dude [UNK] on secret piece my i am so excited for the house of the house .
it is only i the hated senate meet house of this funny right now tonight outage wow i love feeling like the biggest dumbass on earth .

Ours

wow i started feeling like the biggest dumbass on earth .

in irony accuracy. The reason may be that the other four baselines tend to generate common and even not fluent sentences which are irrelevant to the input sentences10 and are hard to be identified as ironies. So annotators usually mark these output sentences as non-ironic sentences, which causes these models to obtain better performances than
DualRL and ours but much poorer results in sentiment and content preservation. Some examples are shown in Table 8.

6

Conclusion and Future Work

In this paper, we first systematically define irony generation based on style transfer. Because of the lack of irony data, we make use of twitter and build a large-scale dataset. In order to control irony accuracy, sentiment preservation and content preservation at the same time, we also design a combination of rewards for reinforcement learning and incorporate reinforcement learning with a pre-training process. Experimental results demonstrate that our model outperforms other generative models and our rewards are effective. Although our model design is effective, there are still many errors and we systematically analyze them. In the future, we are interested in exploring these directions and our work may extend to other kinds of ironies which are more difficult to model.

References
[1] Aniruddha Ghosh, Guofu Li, Tony Veale, Paolo Rosso, Ekaterina Shutova, John A. Barnden, and Antonio Reyes.
Semeval-2015 task 11: Sentiment analysis of figurative language in twitter. In SemEval@NAACL-HLT, 2015.
[2] Jihen Karoui, Farah Benamara, Véronique Moriceau, Viviana Patti, Cristina Bosco, and Nathalie Aussenac-Gilles.
Exploring the impact of pragmatic phenomena on irony detection in tweets: A multilingual corpus study. In EACL
(1), 2017.
[3] Sara Rosenthal, Alan Ritter, Preslav Nakov, and Veselin Stoyanov. Semeval-2014 task 9: Sentiment analysis in twitter. In SemEval@COLING, 2014.
[4] Bo Pang and Lillian Lee. Opinion mining and sentiment analysis. Foundations and Trends in Information
Retrieval, 2007.
[5] Luís Sarmento, Paula Carvalho, Mário J Silva, and Eugénio De Oliveira. Automatic creation of a reference corpus for political opinion mining in user-generated content. In Proceedings of the 1st international CIKM workshop on
Topic-sentiment analysis for mass opinion, 2009.
[6] Cynthia Van Hee, Els Lefever, and Véronique Hoste. Semeval-2018 task 3: Irony detection in english tweets. In
SemEval@NAACL-HLT, 2018.
[7] Tianxiao Shen, Tao Lei, Regina Barzilay, and Tommi S. Jaakkola. Style transfer from non-parallel text by cross-alignment. In NIPS, 2017.
[8] Shrimai Prabhumoye, Yulia Tsvetkov, Ruslan Salakhutdinov, and Alan W. Black. Style transfer through backtranslation. In ACL (1), 2018.
[9] Fuli Luo, Peng Li, Jie Zhou, Pengcheng Yang, Baobao Chang, Xu Sun, and Zhifang Sui. A dual reinforcement learning framework for unsupervised text style transfer. In IJCAI, 2019.
[10] Juncen Li, Robin Jia, He He, and Percy Liang. Delete, retrieve, generate: a simple approach to sentiment and style transfer. In NAACL-HLT, 2018.
10

See the examples in Table 4.

10

A PREPRINT - S EPTEMBER 17, 2019

[11] Jingjing Xu, Xu Sun, Qi Zeng, Xiaodong Zhang, Xuancheng Ren, Houfeng Wang, and Wenjie Li. Unpaired sentiment-to-sentiment translation: A cycled reinforcement learning approach. In ACL (1), 2018.
[12] Guillaume Lample, Sandeep Subramanian, Eric Michael Smith, Ludovic Denoyer, Marc’Aurelio Ranzato, and
Y-Lan Boureau. Multiple-attribute text rewriting. In ICLR (Poster), 2019.
[13] Zhenxin Fu, Xiaoye Tan, Nanyun Peng, Dongyan Zhao, and Rui Yan. Style transfer in text: Exploration and evaluation. In AAAI, 2018.
[14] Zhiting Hu, Zichao Yang, Xiaodan Liang, Ruslan Salakhutdinov, and Eric P. Xing. Toward controlled generation of text. In ICML, 2017.
[15] Omid Rohanian, Shiva Taslimipoor, Richard Evans, and Ruslan Mitkov. WLV at semeval-2018 task 3: Dissecting tweets in search of irony. In SemEval@NAACL-HLT, 2018.
[16] Sepp Hochreiter and Jürgen Schmidhuber. Long short-term memory. Neural Computation, 1997.
[17] Aniruddha Ghosh and Tony Veale. Ironymagnet at semeval-2018 task 3: A siamese network for irony detection in social media. In SemEval@NAACL-HLT, 2018.
[18] Christos Baziotis, Athanasiou Nikolaos, Alexandra Chronopoulou, Athanasia Kolovou, Georgios Paraskevopoulos,
Nikolaos Ellinas, Shrikanth Narayanan, and Alexandros Potamianos. NTUA-SLP at semeval-2018 task 1:
Predicting affective content in tweets with deep attentive rnns and transfer learning. In SemEval@NAACL-HLT,
2018.
[19] Chuhan Wu, Fangzhao Wu, Sixing Wu, Junxin Liu, Zhigang Yuan, and Yongfeng Huang. Thu_ngn at semeval-2018
task 3: Tweet irony detection with densely connected LSTM and multi-task learning. In SemEval@NAACL-HLT,
2018.
[20] Rico Sennrich, Barry Haddow, and Alexandra Birch. Improving neural machine translation models with monolingual data. In ACL (1), 2016.
[21] Yoon Kim. Convolutional neural networks for sentence classification. In EMNLP, 2014.
[22] Ronald J. Williams. Simple statistical gradient-following algorithms for connectionist reinforcement learning.
Machine Learning, 1992.
[23] Ashish Vaswani, Noam Shazeer, Niki Parmar, Jakob Uszkoreit, Llion Jones, Aidan N. Gomez, Lukasz Kaiser, and
Illia Polosukhin. Attention is all you need. In NIPS, 2017.
[24] Diederik P. Kingma and Jimmy Ba. Adam: A method for stochastic optimization. In ICLR (Poster), 2015.
[25] Lajanugen Logeswaran, Honglak Lee, and Samy Bengio. Content preserving text generation with attribute controls. In NeurIPS, 2018.
[26] Kishore Papineni, Salim Roukos, Todd Ward, and Wei-Jing Zhu. Bleu: a method for automatic evaluation of machine translation. In ACL, 2002.

11

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

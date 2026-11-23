---
source: "https://arxiv.org/abs/2611.27541v1"
title: "From Tokens to Pixels: Training-Free Guided Upsampling of Frozen CARL Features for Hyperspectral Imaging"
author: "Dominik Fromme, Ellen Markowski, Tim Streckert, Jörg Thiem"
year: "2026"
publication: "arXiv preprint / eess.IV"
download: "https://arxiv.org/pdf/2611.27541v1"
pdf: "https://arxiv.org/pdf/2611.27541v1"
captured_at: "2026-11-23T06:55:12Z"
updated_at: "2026-11-23T06:55:12Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ユング"
query: "Carl Gustav Jung"
tags:
  - "現代哲学"
  - "深層心理学"
  - "元型論"
status: raw
---

# From Tokens to Pixels: Training-Free Guided Upsampling of Frozen CARL Features for Hyperspectral Imaging

- 著者: Dominik Fromme, Ellen Markowski, Tim Streckert, Jörg Thiem
- 年: 2026
- 掲載情報: arXiv preprint / eess.IV
- 情報源: [arxiv](https://arxiv.org/abs/2611.27541v1)
- ダウンロード: https://arxiv.org/pdf/2611.27541v1
- PDF: https://arxiv.org/pdf/2611.27541v1

## Obsidian Links

- 研究動向: [[ユング-現代研究動向]]
- キーワード: [[ユング]]
- 関連分野: [[現代哲学]]
- 関連分野: [[深層心理学]]
- 関連分野: [[元型論]]
- 関連タグ: #現代哲学 #深層心理学 #元型論

## Abstract

Camera-agnostic foundation models such as CARL map spectral images from arbitrary sensors into a shared feature space, but deliver their features only on a token grid at one eighth of the input resolution. The usual bilinear interpolation blurs material boundaries over about one token and rounds off corners. We investigate how much spatial accuracy image-guided upsampling recovers for frozen CARL features without training the encoder or the upsampler. We compare joint bilateral upsampling (JBU), a guided filter and the pretrained AnyUp, each with the hyperspectral image or an RGB rendering as the guide. A hyperspectral ColorChecker dataset acquired specifically for this study, with model-independent reference geometry measured to sub-pixel precision, serves as the evaluation benchmark. JBU with the hyperspectral guide raises the Boundary IoU from 13 to 67 % and reduces the effective edge width of the features from 38 to 4 px, close to the resolution of the raw spectra. The effect occurs in all 24 patches of the ColorChecker and replicates on an additional recording. On placenta and oral cavity images, the feature edges also become markedly qualitatively sharper. For sharp, high-contrast boundaries, JBU is the most accurate and most economical method. On tissue with softer transitions, the pretrained AnyUp yields qualitatively sharper contours. An analysis of 78 simulated band configurations shows that the gain persists in every configuration as long as the head matches the configuration. Image-guided upsampling is thus a simple, training-free step for applications that use CARL frozen with linear heads.

## PDF Text

From Tokens to Pixels: Training-Free Guided
Upsampling of Frozen CARL Features for
Hyperspectral Imaging
Dominik Fromme, Ellen Markowski, Tim Streckert and Jörg Thiem

arXiv:2611.27541v1 [eess.IV] 20 Nov 2026

University of Applied Sciences and Arts Dortmund
Dortmund, Germany
{dominik.fromme, ellen.markowski, tim.streckert, joerg.thiem}@fh-dortmund.de covers 50 × 50 native camera pixels (Fig. 2). For dense predictions, the reference CARL implementation bilinearly interpolates this grid. Because each token mixes the properties of all materials within its footprint and bilinear weights depend only on distance, material boundaries are blurred over about one token and right-angled corners become rounded (Fig. 1b).
Learnable upsamplers for vision encoders [5]–[7] address this problem but usually have to be trained for each encoder.
In medical hyperspectral imaging, accurate tissue boundaries are particularly important. Examples include tumor margins, blood vessel contours and transitions between perfused and ischemic tissue. Skin and tissue models that assign parameters such as perfusion, optical or thermal properties to individual regions likewise require precise boundaries. At the same time, medical datasets are small and laborious to annotate. A frozen encoder with a linear head is attractive in this setting because it requires few labels, reduces the risk of overfitting and remains interpretable. The key question is how much edge accuracy can be recovered without additional training.
Image-guided upsamplers from image processing [8], [9]
weight neighboring tokens by how similar they are to a highresolution guide image at the target pixel. For hyperspectral data, the input image itself can always serve as such a guide.
It is much finer than the token grid and carries spectral information beyond RGB. Fig. 1a summarizes the pipeline and the comparison. This work contributes:

Abstract—Camera-agnostic foundation models such as CARL
map spectral images from arbitrary sensors into a shared feature space, but deliver their features only on a token grid at one eighth of the input resolution. The usual bilinear interpolation blurs material boundaries over about one token and rounds off corners. We investigate how much spatial accuracy imageguided upsampling recovers for frozen CARL features without training the encoder or the upsampler. We compare joint bilateral upsampling (JBU), a guided filter and the pretrained AnyUp, each with the hyperspectral image or an RGB rendering as the guide.
A hyperspectral ColorChecker dataset acquired specifically for this study, with model-independent reference geometry measured to sub-pixel precision, serves as the evaluation benchmark. JBU
with the hyperspectral guide raises the Boundary IoU from 13
to 67 % and reduces the effective edge width of the features from 38 to 4 px, close to the resolution of the raw spectra. The effect occurs in all 24 patches of the ColorChecker and replicates on an additional recording. On placenta and oral cavity images, the feature edges also become markedly qualitatively sharper.
For sharp, high-contrast boundaries, JBU is the most accurate and most economical method. On tissue with softer transitions, the pretrained AnyUp yields qualitatively sharper contours. An analysis of 78 simulated band configurations shows that the gain persists in every configuration as long as the head matches the configuration. Image-guided upsampling is thus a simple, training-free step for applications that use CARL frozen with linear heads.
Index Terms—Hyperspectral imaging, foundation models, feature upsampling, joint bilateral upsampling, Boundary IoU

I. Introduction
Foundation models for spectral images aim to make multispectral and hyperspectral data usable without task-specific training [1]–[4]. CARL [1] is explicitly designed for changing cameras. It combines a wavelength-aware spectral encoder, which aggregates arbitrary band configurations, with a vision transformer for spatial context. Recordings from different sensors thus end up in a shared feature space and can be analyzed with simple heads such as linear ones, but at the price of spatial resolution.
Like a vision transformer, CARL splits the input image into non-overlapping patches of 8 × 8 pixels and represents each patch by a single token. Its features therefore lie on a token grid at one eighth of the input resolution. How many camera pixels a token covers depends on how the image is resampled before encoding. In our main setting, one token

a systematic evaluation of training-free, image-guided upsamplers (JBU, guided filter, AnyUp) for frozen CARL
features, each with a hyperspectral and an RGB guide and under a pre-specified protocol
• a hyperspectral dataset of a ColorChecker with modelindependent, sub-pixel reference geometry, together with the contamination width 𝑤 90 as a measure of the effective resolution of the features
• a qualitative transfer to placenta and oral cavity images
• a CARL-specific analysis of robustness to wavelength errors and other band sets
•

Code, result tables and access to the ColorChecker data are provided at https://gitlab.fb10.fh-dortmund.de/
fb10-opensource/CARL-tokens-to-pixels.

1

(a)
Ŷ

Bilinear

New HSI dataset

CARL default, no guide

this study

X ∈ ℝC × H × W

CARL

F ∈ ℝ 768 × H/8 × W/8

bands λ1, … , λC

frozen

1 token = 8 × 8 px

0
0
F̂ ∈ ℝ 768 × H × W

all 768 channels

w90 = 38 px

same for both, fitted once

Ŷ

JBU · guided filter · AnyUp
0

Boundary IoU 13 %

Head

Guided upsampler

g ∈ ℝc × H × W

Blurred boundaries

Sharp boundaries
Boundary IoU 67 %
w90 = 4 px
(JBU, HSI guide)

0

guide: RGB or HSI of the same image

(b)

colorful patches dark patches
(black, neutral 3.5)

Bilinear

σ_r = 1/16

σ_r = 1/4

JBU · HSI guide
σ_r = 1/16

σ_r = 1/4

AnyUp
RGB

HSI

max
2 tokens

0
2 tokens

edge strength ‖∇f̂ f̂‖ (scale per image)

JBU · RGB guide
Image (pseudo-RGB)

Figure 1. (a) Pipeline. CARL stays frozen and returns one feature per token of 8 × 8 input pixels. Only the upsampling step is compared. Bilinear interpolation, the default of CARL, uses no guide. JBU, the guided filter and AnyUp take the guide 𝑔 from the same image. All upsamplers bring the 768 feature channels to the output grid 𝐻 ′ × 𝑊 ′ and feed the same head, which is fitted once on tokens. Numbers as in Table I (prototype head). (b) Where do the CARL features change? Sobel gradient magnitude of the unit-normalized features at two patch junctions of the recording rotated by 16°, overlaid on the dimmed image (color scale per row, gamma 0.6, as in Fig. 8). White dashed: measured patch outline. One token covers 50 × 50 native pixels. Top: colorful patches. Bottom: dark patches with low visible contrast to the frame. Bilinear interpolation blurs the edges over about one token. JBU places the edges on the outline but keeps a fringe of about one token at low-contrast boundaries, which widens with 𝜎𝑟 = 𝑠𝑔 /4. With the RGB guide, the fringe partly replaces the edge at Neutral 3.5.
AnyUp also hits the outline, with even narrower edges. With the RGB guide, it additionally traces fine specks on the black patch. 𝜎𝑟 = 𝑠𝑔 /16 is the value selected on the respective other recording.

II. Related Work

C. Edge-Preserving Filters in Hyperspectral Imaging
In spectral–spatial classification, edge-preserving filtering [15] smooths the probability maps of a pixel-wise classifier with a guided or bilateral filter. The first principal components of the hyperspectral image serve as the guide. In contrast, we filter not the output of a classifier but the features of a frozen encoder before the head. We are not aware of any work that systematically examines how much edge fidelity trainingfree, image-guided upsampling recovers for a camera-agnostic hyperspectral encoder and whether the hyperspectral image is a better guide than RGB.

A. Foundation Models for Spectral Images
Besides CARL [1], SpectralGPT [2] for multispectral
Sentinel-2 data, HyperSIGMA [3] for hyperspectral data and
DOFA [4] for multiple Earth observation sensors have been proposed. Only DOFA and CARL explicitly condition on the wavelengths and thus accept changing band configurations. All operate on patch tokens and deliver features at reduced spatial resolution. In medical hyperspectral imaging, public data are scarce [10]. Segmentation networks have so far mostly been trained task-specifically [11], [12].

III. Method
B. Feature Upsampling

A. Problem Statement

Classical image-guided methods are joint bilateral upsampling [8] and the guided filter [9]. For the features of vision encoders, learnable upsamplers such as FeatUp [5], LoftUp [6]
and JAFAR [7] have been proposed. AnyUp [13] and NAF [14]
are encoder-agnostic and can be applied to new features without retraining. These works evaluate almost exclusively
RGB encoders and RGB guides.

CARL maps a spectral image 𝑋 ∈ R𝐶 ×𝐻 ×𝑊 with band wavelengths 𝜆1 , . . . , 𝜆𝐶 to a token map 𝐹 ∈ R𝐷×𝐻/8×𝑊/8 with
𝐷 = 768 channels. Each token feature 𝑓𝑡 ∈ R𝐷 belongs to a footprint Ω𝑡 of 8 × 8 input pixels with center 𝑐 𝑡 . A dense prediction requires a feature 𝑓ˆ(𝑥) ∈ R𝐷 at every pixel 𝑥 of an output grid of 𝐻 ′ × 𝑊 ′ pixels, which is finer than the token grid but otherwise independent of the input size 𝐻 × 𝑊. Here,

2

the output grid is the evaluation grid of 5 × 5-pixel blocks, so one token covers 10 × 10 output pixels (Fig. 2a–d). Positions
𝑥 and 𝑐 𝑡 are given in token units. Except for the guided filter, all methods considered compute a convex combination of the tokens,
∑︁
∑︁
𝑓ˆ(𝑥) =
𝛼(𝑥, 𝑡) 𝑓𝑡 ,
𝛼(𝑥, 𝑡) ≥ 0,
𝛼(𝑥, 𝑡) = 1,
𝑡

JBU uses its first 16 principal components. AnyUp receives the first three principal components as a false-color image, because it is pretrained on three-channel RGB images and more channels would require retraining. As in edge-preserving filtering [15], the guided filter also uses the first three principal components and is evaluated only with this guide. All guides are formed on the evaluation grid of 5 × 5-pixel blocks, ten times finer than the 50-px token grid.

𝑡

(1)
and differ only in the weights. Nearest selects the token whose footprint contains 𝑥. Bilinear interpolation uses tent weights based on the distance to the token centers. Neither set of weights depends on the image content, so the feature edges lie on the token scale. Because the weights form a convex combination, 𝑓ˆ(𝑥) stays within the range of the tokens and their semantics are preserved. For a linear head, it also makes no difference whether features or logits are upsampled. Every method upsamples all 𝐷 feature channels. The guide image only determines the weights.

IV. Experimental Setup
A. Dataset
We acquired the recordings with the setup described in [10].
It consists of a Resonon Pika UV-505 push-broom camera
(CMOS, 12 bit, 1500 pixels across track) with a Jenoptik
UV-VIS-IR 60 mm 1:4 APO Macro lens (apochromatic for
315–1100 nm) on a Resonon benchtop system whose linear stage moves the sample under the camera, inside an enclosure lined with absorbing foil. Illumination is provided by a halogen line light fed via an optical fiber. The UVA fluorescent tube of the setup (UV lamp in the following) was additionally switched on only for the third recording. With 4-fold spectral binning, the camera delivers 270 bands between 325 and 821 nm
(spacing 1.8–1.9 nm, full width at half maximum 2.8 nm).
Patzke et al. [10] use the 255 bands within 330–800 nm.
Exposure time and line rate are 222 ms and 4.5 Hz (232 ms and 4.3 Hz for the third recording) at gain 0. The working distance is about 45 cm. The acquisition software converts raw values 𝐼 into reflectance using white and dark references
𝑊 and 𝐷,
𝐼 (𝜆) − 𝐷 (𝜆)
.
(4)
𝑅(𝜆) =
𝑊 (𝜆) − 𝐷 (𝜆)

B. Image-Guided Upsamplers
Joint bilateral upsampling [8] computes each output pixel from a local neighborhood of tokens. We use a 5 × 5 window centered on the token containing the output pixel, extending two tokens in each direction. Within this window, token contributions are weighted by spatial distance and guide similarity,


∥𝑥 − 𝑐 𝑡 ∥ 2 ∥𝑔(𝑥) − 𝑔¯ 𝑡 ∥ 2
−
,
(2)
𝛼(𝑥, 𝑡) ∝ exp −
2𝜎𝑠2
2𝜎𝑟2
where 𝑐 𝑡 is the center of token 𝑡, 𝑔(𝑥) is the guide at output pixel 𝑥 and 𝑔¯ 𝑡 is the mean guide value over the footprint Ω𝑡
of token 𝑡. As in [8], the spatial bandwidth is set to 𝜎𝑠 =
0.5 token. The guides are scaled to an RMS spread 𝑠𝑔 = 1, so the bandwidth 𝜎𝑟 has the same scale for the RGB and the hyperspectral guide.
As in edge-preserving filtering [15], the guided filter [9]
is applied after bilinear upsampling. Let 𝑝(𝑥) ∈ R𝐷 be the bilinearly upsampled feature and 𝑔(𝑥) ∈ R3 the guide. In every window 𝜔 𝑘 of radius one token centered at pixel 𝑘, the filter fits a linear model 𝑝 ≈ 𝑎 ⊤𝑘 𝑔+𝑏 𝑘 with 𝑎 𝑘 ∈ R3×𝐷 and 𝑏 𝑘 ∈ R𝐷 ,
𝑎 𝑘 = (Σ 𝑘 + 𝜀𝐼) −1 cov 𝑘 (𝑔, 𝑝),
𝑓ˆ(𝑥) = 𝑎¯ ⊤𝑥 𝑔(𝑥) + 𝑏¯ 𝑥 ,

𝑏 𝑘 = 𝑝¯ 𝑘 − 𝑎 ⊤𝑘 𝑔¯ 𝑘 ,

An X-Rite ColorChecker Classic Mini with 24 patches was recorded in two orientations, upright (2429 × 1500 px) and rotated by 16.1° (2600 × 1500 px). In the rotated recording, the image border cuts four patches (Fig. 3). A token of the main setting covers 50 × 50 px, about one seventh of the patch pitch of 340 px. The frame gap between two patches is only
1.4 tokens wide at about 70 px. A third, upright recording with the UV lamp additionally switched on (2700 × 1500 px) was evaluated only after the main analysis, as a replication and as a test of the UV range (Section V-D).

(3)

B. Reference Geometry

R3×3

where Σ 𝑘 ∈
is the covariance of the guide in 𝜔 𝑘 , cov 𝑘 (𝑔, 𝑝) ∈ R3×𝐷 the covariance between guide and feature,
𝑔¯ 𝑘 and 𝑝¯ 𝑘 their means. 𝑎¯ 𝑥 and 𝑏¯ 𝑥 are the averages of 𝑎 𝑘
and 𝑏 𝑘 over all windows that contain 𝑥. The regularization
𝜀 limits how strongly edges of the guide pass into the features. The guided filter is the only method that is not a convex combination. AnyUp [13] computes the weights with a pretrained, encoder-agnostic window attention between guide pixels and tokens. We use the published model without retraining. Fig. 2 shows the grids involved and the weights of bilinear interpolation, JBU and AnyUp for one pixel.
The guide is either an RGB rendering of the cube (440–500,
510–570 and 600–680 nm) or the hyperspectral image itself.

The reference does not depend on any model. Each fully visible patch is measured at native resolution. From pseudoRGB and NIR, we form a normalized contrast that is 0 on the frame and 1 inside the patch. Its 0.5 isoline defines the edge. We fit a total-least-squares line to each side and take the intersections of adjacent lines as the corners. The patch identities follow from comparing the measured spectra with the BabelColor reference spectra [16] over all rotations and reflections of the chart. All 24 patches are measured in the upright recording and 20 in the rotated one. Each side is located to within 0.8 px RMS, far below the 10-px band of the
Boundary IoU. An ideal chart model, i.e., a homography that maps a perfect patch grid into the image, matches the measured

3

CARL input
8 × 8 px per token

Nearest

→

1 px = 6.25 × 6.25
native px

h

Bilinear
(CARL default)

d

e
Output grid
10 × 10 px per token

f

Distances from x to every token

→

1 px = 5 × 5
native px

ft ∈ ℝ 768, shown as qt
(0 = frame, 1 = patch)

i

Guide g on the output grid

JBU
σr = sg

j

JBU
σr = sg/4

HSI, PC 1–3 shown
(JBU: 16 PCs, AnyUp: 3)

k

JBU (selected)
σr = sg/16

white: ‖x − ct‖ in tokens blue: ‖g(x) − ḡt‖ in sg

l

AnyUp
(pretrained)

n

o

p

q

(as q)

Upsampled f̂ f̂

m

r

1 (patch)
0.9 (w90)

0.5

0 (frame)

apparent patch fraction q

Weights α(x, t)

g

CARL output
1 feature per token

CARL

3 × 3 tokens

→

1 px = 1 native px

c

upsampling

b
Camera image
50 × 50 px per token area mean

a

Figure 2. Grids and upsampling weights for one output pixel 𝑥 (red) just inside the lower edge of the orange patch in the rotated recording. Native px are the pixels of the camera. (a–d) The same 3 × 3 tokens at each step of the pipeline. The token that contains 𝑥 is outlined in blue. One token covers 50 × 50 native px (a). For CARL, the image is resampled by area averaging to 8 × 8 input px per token, so one input px corresponds to 6.25 × 6.25 native px (b). CARL
returns one feature 𝑓𝑡 ∈ R768 per token (c). The upsamplers compute 𝑓ˆ ( 𝑥 ) on the output grid with 10 × 10 px per token, so one output px corresponds to
5 × 5 native px (d). The guide and the evaluation use this grid as well. Features are shown as apparent patch fraction 𝑞 (0 = frame, 1 = patch, Section IV-D).
Although 𝑥 lies on the patch, its token mainly covers the frame (𝑞𝑡 = 0.11). (e) Hyperspectral guide on the output grid, shown by its first three principal components. JBU uses the first 16, AnyUp the first three. (f) The two distances that enter JBU (Eq. 2) for every token. White: distance between 𝑥 and the token center 𝑐𝑡 (lines in d) in token units. Blue: distance between the guide at 𝑥 and the mean guide of the token in units of 𝑠𝑔 . The token of 𝑥 is closer in position (0.35 compared with 0.65 for the token above). The token above, however, matches the guide (0.02 compared with 0.98). (g–l) Weights 𝛼( 𝑥, 𝑡 ) as disks with area proportional to the weight. Only tokens in the framed region can contribute. These are the token of 𝑥 (nearest), the four nearest token centers
(bilinear) and the 5 × 5 window centered on the token of 𝑥 (JBU). AnyUp attends to a window that spans ±10 % of the image height and width around 𝑥, here 12 × 7 tokens. It extends beyond the crop and 24 % of the weight lies outside. (m–r) Upsampled features as 𝑞 with the contour 𝑞 = 0.9 (solid) and the measured outline (white dashed). Their distance, averaged along a side, corresponds to the contamination width 𝑤90 . Bilinear weights depend only on the position. The smaller 𝜎𝑟 , the more weight JBU gives to the token above. The edge of 𝑓ˆ then moves to the measured outline. AnyUp spreads its weight over many tokens, mostly on the patch, and reaches 𝑞 ( 𝑥 ) = 0.86.

corners only to 2.2–2.4 px RMS. The largest deviations occur in both recordings at the same four low-contrast patches (Dark
Skin, Black, Foliage, Neutral 3.5), which are measured about
3.5–7 px larger than the median of the other patches (mean side lengths 265–275 px). For three of them, the reference edge follows the NIR channel and lies about 3 px further out than in the visible range. We therefore use this chart model, i.e., the fitted homography of the ideal grid, only where nothing can be measured. It defines the outer boundary of the evaluated frame region, half a frame gap beyond the outer patches. Everything outside is ignored. It also places the four patches in the rotated recording that are cut by the image border. The classes are the
24 patches and the frame. The radiometry is consistent. The median color difference Δ𝐸 00 is 0.2 between the recordings and 1.7–1.8 to the BabelColor average (30 charts, measured before the 2014 reformulation of the colors).

together with their wavelengths, normalized per image to zero mean and unit standard deviation. Bands below 400 nm are not used because their SNR is clearly below that of the visible range (Fig. 3d) and CARL was pretrained on sensors starting at about 420 nm. The input is area-averaged exactly so that a token covers 50 × 50 native pixels. We additionally examine 20
and 25 px. We evaluate the spatial token map (768 channels), which CARL’s segmentation head also reads, on a grid of
5 × 5-pixel blocks.
Tokens whose footprint lies at least 95 % within one class form the support of that class. Two heads are fitted once on these tokens and never see an upsampled feature, so differences are due to the upsampler alone. The first is a training-free prototype head (class means of the support tokens, assignment by the smallest angle). The second is a linear probe (multinomial logistic regression, corresponding to
CARL’s 1 × 1 convolutional head). The prototype rule on the raw spectra with all 229 bands between 400 and 820 nm serves as an upper reference. In this work, training-free refers to the upsampling. CARL stays frozen, JBU and the guided filter

C. Features and Heads
We use the public CARL SSL checkpoint [1] frozen. CARL
receives 32 linearly spaced bands between 400 and 820 nm

4

c
0° · 24/24 measured

16.1° · 20/24 measured

measured (0°)

0.8

reflectance

a

b

BabelColor

red light skin bluish green

0.6

blue

0.4
0.2

neutral 5
black 2

0.0
400

500

600

700

800

wavelength (nm)

d

SNR

200

not used

100
50

patch 19 (white)
patch 22 (neutral)
CARL bands

20
400

500

600

700

800

wavelength (nm)

Figure 3. Dataset and reference. (a, b) Both recordings as pseudo-RGB with measured patch outlines (white dashed) and model-placed, cut-off patches (white dotted). (c) Measured reflectance spectra of selected patches (mean over the patch interior, upright recording) and, dashed, the BabelColor average of the same patches [16], tabulated from 380 to 730 nm. (d) SNR per band in the white and the neutral 5 patch. Bands below 400 nm (gray) are not used. Blue ticks mark the 32 CARL bands.

have no learned parameters and AnyUp is used as published.
Only the heads are fitted once on labeled support tokens.

100-px tokens in the token sweep with 20 px, because 100px tokens are wider than the frame gap. The two parameters without a transferable default, 𝜎𝑟 ∈ {1/16, . . . , 1} · 𝑠𝑔 and
𝜀 ∈ {0.012 , . . . , 0.22 }, are selected across recordings. The best value of one recording is used on the other. Thus, no reported number is optimized on its own labels.
The statistical unit is the physical patch (𝑛 ≤ 24). For each patch, the difference to bilinear interpolation is averaged over the recordings. We report the median of these differences with a 95 % cluster bootstrap confidence interval (10 000 resamples [18]) and test them with the Wilcoxon signed-rank test.
The primary comparison, fixed in advance, is JBU with the hyperspectral guide versus bilinear in Boundary IoU (spatial features, prototype head, 𝛼 = 0.05). All other comparisons are
Holm-corrected within their analysis [19].

D. Metrics
The Boundary IoU [17] assesses class membership only within a band of 𝑑 = 10 native pixels around the contours of reference and prediction, the Corner IoU within circles of
25 px radius around the four reference corners of a patch. We measure the effective resolution of the features directly on the features. For a patch with prototype 𝜋 𝑘 next to the frame with prototype 𝜋 𝑓 ,
𝑞(𝑠) =

⟨ 𝑓ˆ(𝑠) − 𝜋 𝑓 , 𝜋 𝑘 − 𝜋 𝑓 ⟩
∥𝜋 𝑘 − 𝜋 𝑓 ∥ 2

(5)

is the apparent patch fraction at distance 𝑠 from the reference edge (𝑠 > 0 inside the patch), averaged along the side. Since
Eq. 5 is normalized to the prototypes, which are means over all support tokens, single tokens can yield 𝑞 slightly below 0 or above 1. Edge tokens with up to 5 % frame in their footprint add to this, as they shift 𝜋 𝑘 slightly toward the frame. We average only over the central 60 % of each side. Each profile thus stays about 55 px, i.e., more than one token, away from the adjacent sides and captures only the one edge. The corners are assessed by the Corner IoU. The contamination width 𝑤 90
is the distance into the patch beyond which 𝑞 ≥ 0.9 holds, i.e., less than 10 % frame is mixed in, taken as the median over the four sides.

F. Medical Images
Our qualitative evaluation uses two public datasets. The first contains placental images acquired with a Senop HSC-2 camera mounted on an operating microscope (1024 × 1024 pixels) [20]. The second, ODSI-DB [21], contains images of teeth, oral mucosa and lips acquired using either a Specim
IQ (204 bands, 400–1000 nm) or a Nuance EX (51 bands,
450–950 nm). For the placenta, we use the 35 bands in
515–698 nm that all images contain, excluding the bands around 640 nm with a known sensor anomaly. For ODSIDB, we use the common range of 450–950 nm and every third Specim band. CARL receives both at half resolution, so a token covers 16 × 16 native pixels. We examine the edge maps of the upsampled features themselves, without a

E. Protocol and Statistics
All settings were fixed before the first upsampling result.
The only later change, also made before any result, replaced

5

head and without training. They show the Sobel gradient magnitude of the unit-normalized features, as in Fig. 1b. We deliberately keep this part qualitative, because the accuracy of the annotated boundaries in both datasets is unknown and boundary metrics would partly measure label uncertainty.

D. Hyperspectral versus RGB Guide
Across all patches, the hyperspectral guide is not better than the RGB guide, neither for JBU nor for AnyUp (Holm
𝑝 ≥ 0.18). For most patches, the RGB image alone separates patch and frame. For JBU, the gain is concentrated on a few patches with low visible contrast to the dark frame (Neutral
3.5, Black, Dark Skin, Foliage), with up to +29 pp Boundary
IoU. Its size follows the spectral angle between patch and frame, not the color difference Δ𝐸 00 (Appendix B). With reference edges from the visible range alone, this pattern and all main results remain, although the gain on Foliage becomes smaller. The hyperspectral image thus helps exactly where two materials differ spectrally but look similar in the RGB image.
For AnyUp, the differences between the guides scatter in both directions and follow the spectral angle only in 𝑤 90 .
With the third recording, we test whether the bands below
400 nm, which an RGB image does not capture, provide the guide with additional edges. The camera records these bands in every recording. The UV lamp was meant to strengthen their signal. This does not happen. The reflectance equals that of the recording without the lamp (median Δ𝐸 00 0.3). The
SNR below 400 nm also remains unchanged. Accordingly, a guide from the principal components of 325–820 nm is not better than one from 400–820 nm (JBU and AnyUp −0.1 pp,
Holm 𝑝 ≥ 0.3, Table A1). The principal component analysis suppresses the UV only partly. The UV bands carry 3 % of the variance and barely enter the three components for AnyUp, but the 16 components for JBU retain 99.9 % of their variance.
Weighting the UV more strongly in a post hoc test makes the guide worse. With a UV-only guide, JBU loses a median of
17 pp Boundary IoU. With a guide that is half UV, it still loses
4 pp (Holm 𝑝 = 0.02). The UV bands are blurred along the slit and therefore carry less edge information than the visible range. Whether UV information can improve the guide thus remains open.

G. Implementation
The experiments were implemented in Python with PyTorch and run on an NVIDIA H100 GPU. The analysis and figure code was developed with the AI coding assistant Claude Code
(Claude Opus 5.5, Anthropic) under the authors’ direction and checked by them. All numbers and figures are computed by this code from the recorded data. Code and notebooks are published for verification.
V. Results
A. Edges and Corners
JBU with the hyperspectral guide raises the Boundary IoU
over bilinear interpolation by a median of 55 pp (95 % CI
47–61 pp, 𝑝 ≈ 10−7 ), in all 24 patches and equally in both recordings. On the third recording, evaluated later, the gain is
+52 pp, also in all 24 patches. All other guided methods also clearly improve Boundary IoU and Corner IoU, for both heads and both feature maps (Holm 𝑝 < 0.001, Table I, Fig. 4).
The difference is most pronounced at the corners. Bilinear interpolation rounds them off so strongly that it performs worse than the raw token grid (Fig. 5). With JBU, the CARL features reach the level of the raw spectra in Boundary IoU.
B. Effective Resolution of the Features
After bilinear upsampling, frame information reaches a median of 38 px, almost one token, into a patch before less than 10 % is mixed in. With JBU, it is 4 px, barely more than for the raw spectra (Fig. 6). The profiles show two further failure patterns. The guided filter places the edge correctly but leaves a long tail inside the patch, because it corrects the bilinear ramp, which is about two tokens wide, only in windows that contain the edge. AnyUp yields sharp edges but, for some dark patches, still mixes in frame far inside the patch. Edge maps show only where the features change, not how pure they are. In Fig. 1b, the edges of AnyUp even appear the narrowest, whereas JBU keeps a faint fringe at low-contrast boundaries. Only the profiles capture the residual mixing inside the patch. The findings do not depend on the threshold of 0.9 (Appendix D).

E. Robustness and Cost
With heads from the respective other recording, the Boundary IoU drops by at most 4 pp. The ranking of the methods is preserved. The spectral feature map and the linear probe give the same picture (Tables A2 and A3). For a feature map with 768 channels at 520 × 300 px, JBU takes about
40 ms. AnyUp needs about twice the time and memory, the guided filter almost three times the time (Table A4). The guide does not change this. JBU takes equally long with three RGB
channels and with 16 principal components, because the cost is dominated by the 768 feature channels. AnyUp receives a three-channel image in both cases. The principal component analysis for the guide is not included.

C. Token Size and Parameter Selection
The gain grows with the token size, because bilinear interpolation deteriorates with larger tokens while JBU remains nearly constant (𝑤 90 ≈ 4 px, Fig. 7a, b). With an image guide, the effective resolution of the features is thus decoupled from the token size in this range. The Boundary IoU of
JBU decreases monotonically with 𝜎𝑟 . The curves of both recordings nearly coincide (Fig. 7c), so the selection across recordings finds the best value. For the guided filter, it costs about 4 pp in one recording. The spatial bandwidth 𝜎𝑠 is uncritical (Appendix C).

F. Medical Images
On tissue, the same pattern appears qualitatively as on the
ColorChecker (Fig. 8, cf. Fig. 1b). Upsampled bilinearly, the features form a blurred mosaic of token blocks whose edges follow the 16-px grid. Both guided methods produce markedly narrower and stronger gradients at image structures. These

6

Table I
Results on the ColorChecker (spatial CARL features, mean of both recordings, 50-px tokens). IoU values in %, mean over measured patches.
Median 𝑤90 in native px.
Method

Boundary IoU (prototype)

Boundary IoU (linear)

Corner IoU (prototype)

Patch IoU (prototype)

𝑤90 (px)

8.9
13.3
56.6
62.8
67.2
60.4
63.6
63.2

9.9
18.4
66.6
67.1
70.3
64.3
69.2
–

16.5
10.0
55.9
77.0
81.4
75.7
78.3
80.4

70.7
78.1
92.9
94.9
96.2
95.3
94.9
95.7

27.9
38.2
71.0
3.8
3.8
9.2
22.6
3.2

Token grid (nearest)
Bilinear (CARL default)
Guided filter, PCA guide
JBU, RGB guide
JBU, HSI guide
AnyUp, RGB guide
AnyUp, HSI guide
Raw spectra (no CARL)

Value per patch (dots: 24 patches, bar: median)

a

b

c

Token grid (nearest)

9

15

27

Bilinear (CARL default)

12

5

38

GF · PCA guide

67

65

71

JBU · RGB guide

65

88

4

JBU · HSI guide

73

89

4

AnyUp · RGB guide

61

78

9

AnyUp · HSI guide

67

82

20

Raw spectra (no CARL)

67

84

4

0

25

50

75

100

0

25

50

Boundary IoU (%)

75

100

0

25

Corner IoU (%)

50

75

100

w90 (px)

Paired difference to bilinear (median and 95 % CI over patches)

d

e

f

Token grid (nearest)

+0

+6

−12

GF · PCA guide

+46

+53

+33

JBU · RGB guide

+51

+74

−30

JBU · HSI guide

+55

+75

−32

AnyUp · RGB guide

+45

+68

−25

AnyUp · HSI guide

+50

+69

−16

0

20

40

60

80

0

20

Δ Boundary IoU (pp)

40

60

Δ Corner IoU (pp)

80

100

−20

0

20

40

60

Δ w90 (px), < 0 = better

Figure 4. Main result (spatial CARL features, prototype head, 50-px tokens). (a–c) Value of each measured patch (dots, mean of both recordings) and median
(bar) for Boundary IoU, Corner IoU and 𝑤90 , with 𝑤90 censored at 100 px. (d–f) Median of the paired difference to bilinear interpolation with 95 % cluster bootstrap CI. For 𝑤90 , a negative value means narrower edges.

include tooth outlines, the gingival margin and the lip as well as vessel borders and specular highlights of the placenta.
Qualitatively, AnyUp yields the sharpest contours. Its edges are narrower than those of JBU and follow vessel borders and tooth outlines most closely. JBU depends more strongly on the bandwidth. With 𝜎𝑟 = 𝑠𝑔 /16, selected on the ColorChecker, it also takes over texture, specular reflections and noise of the guide. With 𝜎𝑟 = 𝑠𝑔 /4, the strong boundaries remain, but wider than with AnyUp. In low-contrast regions, JBU also traces the token grid as stripes. RGB and hyperspectral guides mostly yield similar edges. On textured skin, AnyUp takes

over the skin texture with the RGB guide but hardly with the hyperspectral guide. The edge maps show where the CARL
features change, not whether a true tissue boundary lies there.
Appendix G shows all images considered.
G. CARL-Specific Analysis: Wavelengths and Band Sets
CARL is camera-agnostic. It encodes each band together with its wavelength, allowing it to handle different band configurations. For users of the frozen encoder, it is therefore important to assess the required wavelength accuracy and whether guided upsampling remains effective across different band sets.

7

b 16° · colorful patches

c 16° · dark patches (black, neutral 3.5)

d

e

f

RGB guide

HSI guide

a 0° · colorful patches

Bilinear (CARL default)

JBU

AnyUp

reference outline

Figure 5. Decision boundaries of the prototype head at three patch junctions, top with the hyperspectral guide (a–c), bottom with the RGB guide (d–f). White dashed: measured outline. Bilinear rounds off corners and pulls edges into the patch. JBU follows the outline with both guides. At the dark patches, JBU with the RGB guide deviates from the outline (f), with the hyperspectral guide it does not (c).

range from −50 to +50 nm, depending on the available spectral range. Second, we investigate alternative band configurations, including 8 to 128 bands, random subsets, a restricted range of
400–700 nm and band positions corresponding to EnMAP and
Sentinel-2. EnMAP [22] is a German hyperspectral satellite.
We use the 69 centers of the EnMAP band grid from CARL’s pretraining that lie in 400–820 nm (spacing 4.7–7.7 nm).
Sentinel-2 [23] is a multispectral satellite of ESA. Its seven bands B1–B7 lie in this range (Sentinel-2A, 443–783 nm).
The hyperspectral guide of JBU always comes from the full cube and stays unchanged. The analysis thus tests robustness at a constant guide, not a real sensor that provides only these bands.

apparent patch fraction q(s)

1.0
0.9

0.8
0.6
Token grid (nearest)
Bilinear (CARL default)
GF · PCA guide
JBU · RGB guide
JBU · HSI guide
AnyUp · RGB guide
AnyUp · HSI guide
Raw spectra (no CARL)

0.4
0.2
0.0
1 token

−20

0

20

40

60

80

100

For each variation, we read out the features with the prototype head in two ways. The transferred head was fitted once on the protocol setting and is reused unchanged. A user with a ready-made head would do the same. The refitted head is determined on the features of the variation itself and shows whether the upsampler still helps independently of the head
(Fig. 9).

distance from reference edge s (px), > 0 inside

Figure 6. Apparent patch fraction 𝑞 (𝑠) over all measured patch edges of both recordings. Gray background: one token (50 px). Dashed: threshold 0.9
for 𝑤90 .

We evaluate 78 configurations, including the reference setting of 32 bands spanning 400–820 nm, on both recordings. First, we examine how the features respond to wavelength shifts.
We shift either the wavelength labels passed to CARL, the spectral bands sampled from the data or both together. Shifts

With the transferred head, wavelength errors up to about
±10 nm are uncritical (at most 5 pp loss in Boundary IoU), larger errors are not. Wrongly stated wavelengths are the most harmful. If data and stated wavelengths are shifted together, the

8

a

b

100
Token grid (nearest)
Bilinear (CARL default)
GF · PCA guide

80

w90 (px)

80

Boundary IoU (%)

c

100

60
40
20

JBU · RGB guide
JBU · HSI guide
AnyUp · RGB guide
AnyUp · HSI guide

60
40
20

0
50

JBU · HSI guide · 16°

60
40
selected
(cross-scene)

20

0
20 25

JBU · RGB guide · 0°
JBU · RGB guide · 16°
JBU · HSI guide · 0°

80

Boundary IoU (%)

100

0
20 25

token size (native px)

50

1/16

1/8

token size (native px)

1/4

1/2

1

σ_r / s_g

Figure 7. (a) Boundary IoU and (b) 𝑤90 as a function of the token size. Dots: mean (𝑤90 : median) over patches, error bars: 95 % bootstrap CI. (c) Boundary
IoU of JBU as a function of the bandwidth 𝜎𝑟 (in units of the RMS spread 𝑠𝑔 of the guide) for both recordings. Dashed: the value selected across recordings.
JBU · RGB guide
Bilinear

σ_r = 1/16

σ_r = 1/4

JBU · HSI guide
σ_r = 1/16

σ_r = 1/4

AnyUp
RGB

HSI

Placenta
P022

RGB (camera)

2 tokens

Placenta
P027

edge strength ‖∇f̂ f̂‖ (scale per image)

max

Oral cavity
690db113

2 tokens

2 tokens

Oral cavity ba9444b8

0

2 tokens

Figure 8. Edge maps of the upsampled CARL features on placenta and oral cavity images, without a head and without training, overlaid on the dimmed image
(color scale per image, gamma 0.6 emphasizes weak edges). Crops of 256 × 256 px (16 × 16 tokens). Placenta as the unprocessed RGB image, oral cavity with the annotated class boundaries (white) for orientation. JBU with the RGB or hyperspectral guide (16 principal components) and 𝜎𝑟 = 𝑠𝑔 /16 or 𝑠𝑔 /4.
AnyUp with the RGB image or the hyperspectral false-color image (principal components 1–3) as the guide. Images adapted from the placenta dataset [20]
and ODSI-DB [21]. This figure is licensed under CC BY-NC-SA 4.0 and is not covered by the license of this article.

head remains largely usable. For other band sets, the spectral coverage matters more than the number of bands. If the range above 700 nm is missing, as for visible-only, Sentinel-2-like or RGB-like bands, the transferred head collapses. With the refitted head, in contrast, JBU outperforms bilinear interpolation by a median of 46 to 57 pp in all 78 variations and all
24 patches (Holm 𝑝 < 10−5 ). With changed bands, it is thus the transferred head that fails, not the upsampler.
The token maps show where the features change (Fig. 10).
They remain nearly unchanged with a 10-nm error, with only eight bands or when data and stated wavelengths are shifted together (median cosine to the protocol map 0.90–0.98). With

a 50-nm error or without the range above 700 nm, mainly the frame changes, whose spectrum is dark in the visible range and rises steeply above 700 nm. The patches remain internally uniform and distinguishable from each other. Because every patch edge borders the frame, the transferred head fails exactly at these edges, whereas a refitted head separates the classes again.
VI. Discussion
Anyone who uses CARL without fine-tuning and reads out the features directly with a linear head or a prototype classifier loses a large part of the edge fidelity through bilinear interpola-

9

Wrong wavelengths

a (JBU · HSI, head transferred)
70

Other band sets

b protocol: 32 bands

JBU · HSI, head refitted (all Δ)

128 bands

Boundary IoU (%)

60

64 bands
16 bands

50

8 bands
40

32 bins, averaged
EnMAP (69 bands)

stated λ shifted

30

data shifted

32 random bands (5×)

20

both shifted

10

bilinear (reference)

32 bands, 400–700 nm only
Sentinel-2 (7 bands)
3 RGB windows

0
−50 −20 −10

0

10

20

50

0

20

wavelength error Δ (nm)

40

60

Boundary IoU (%)
bilinear, head refitted

JBU · HSI, head refitted

JBU · HSI, head transferred

Figure 9. Robustness to wavelengths and band sets (spatial CARL features, prototype head, mean of both recordings). (a) Boundary IoU after JBU with the transferred head when only the wavelengths passed to CARL (stated 𝜆), only the data or both are shifted by Δ. The data were shifted only up to +10 nm, with the top band limited by the cube end at 820.6 nm. Blue band: range with the refitted head. Dashed: bilinear. (b) Other band sets: bilinear (gray) and
JBU (blue) with the refitted head. Open: JBU with the transferred head. Random bands: mean and range over five draws.
stated λ
+10 nm

stated λ
+50 nm

data and stated λ
−20 nm

8 bands

32 bands
400–700 nm only

3 RGB windows

Image (pseudo-RGB)

median 0.98

median 0.65

median 0.95

median 0.90

median 0.59

median 0.62

1.0

0.5

0.0

cosine to the protocol map

principal components 1–3

protocol
32 bands

Figure 10. Spatial CARL token maps of the rotated recording under selected band changes. Top: principal components 1–3, all in the color basis of the protocol map. Bottom: cosine of each token to the protocol map, with the median above. Stated 𝜆: only the wavelengths passed to CARL shifted. Data and stated 𝜆: both shifted. The maps of all 78 variations are published with the code.

tion. Image-guided upsampling recovers it without additional labels or further training. The two scalar parameters can be selected on a different recording. On the ColorChecker, JBU
was the best or joint-best method in all endpoints, robust to the parameter choice and at the same time the most economical.
AnyUp requires about twice the time and memory as well as a pretrained network. It also left frame contributions inside lowcontrast patches. On CARL’s token scale, the guided filter is a compromise between a sharp decision boundary and a clean patch interior.

Which method is preferable thus depends on the edges.
Where boundaries are sharp and high-contrast, so that the bandwidth of JBU can be determined reliably on another recording, as on the ColorChecker, JBU is the most accurate method. Tissue has softer, lower-contrast and textured transitions. There, AnyUp qualitatively yields the sharper contours, presumably because its learned attention handles such transitions better than the fixed Gaussian kernel of JBU, which, depending on the bandwidth, takes over texture or widens edges (Fig. 8). That AnyUp already shows the narrowest

10

edge maps on the ColorChecker (Fig. 1b) but trails JBU
quantitatively is consistent with its residual mixing inside dark patches, which an edge map does not show. For tissue, the choice therefore remains open until a quantitative evaluation is available. Whether sharper feature edges lead to more accurate segmentations there also depends on the head and on the annotation. Since the guide comes from the same recording, the camera independence of CARL is preserved in any case.
A hyperspectral guide helps where the visible contrast is small and the spectral difference is large. On the ColorChecker, this is rare. In tissue, we expect it more often, for example with differences in oxygenation or water content. The UV
recording could not test this specifically, because the lamp does not change the reflectance of the chart. A real test requires materials that differ or fluoresce in the UV. It also requires a recording that preserves this and in which the UV bands are as sharp as the visible ones. The CARL-specific analysis further shows that band centers should be known to within about
±10 nm and that a head belongs to the band configuration on which it was fitted. The upsampler itself can be used unchanged with any configuration.
The study has limitations. The statistics refer to the 24
patches of one chart in two recordings. The third recording was evaluated only afterwards. With its sharp, high-contrast edges, the ColorChecker is a favorable case for any guide.
The selected bandwidth lay at the edge of the pre-specified grid. The heads were fitted only on pure interior tokens, which isolates the effect of the upsampler but does not correspond to the usual training on all pixels. The band configurations are simulated from the same cubes, so the noise, spectral response and imaging properties of a real second camera are missing.
The medical results are qualitative and based on a few images.
Next steps. A natural next step is to test whether the gain holds once heads are trained on all pixels. Bilinear and guided upsampling could be compared under otherwise identical training, with a linear or a nonlinear head on frozen features and with fine-tuning of encoder and head. The upsampler could be applied either at inference only or already during training. A quantitative evaluation on medical datasets with dense annotations, for example in Boundary IoU and
Boundary F1, would show whether the advantage carries over to segmentation.

for sharp, high-contrast edges. On tissue with softer transitions, the qualitative results point to an advantage for AnyUp.
Whether it holds must be shown by future quantitative studies.
A natural next step is to examine whether the gain holds under training and fine-tuning for segmentation tasks or becomes marginal with nonlinear heads. The hyperspectral guide should also be tested on materials that differ only outside the visible range.
Data and Code Availability
Code, result tables and notebooks are available at https://gitlab.fb10.fh-dortmund.de/fb10-opensource/
CARL-tokens-to-pixels. The repository README provides access to the ColorChecker dataset with all acquisition metadata. A permanent DOI-based deposit of the dataset is in preparation. The placenta and ODSI-DB images are not redistributed. They are available from the original sources [20], [21] under CC BY-NC-SA 4.0. The repository lists all images used. All experiments are available as executable scripts and notebooks. Apart from the runtime measurements, both execution paths yield byte-identical result tables.
Acknowledgment
This research was funded by the German Federal Ministry of Research, Technology and Space (BMFTR) under the
Forschung an HAW program (grant number 13HAW021FQ).
Equipment used in this study was purchased as part of a previous research project funded by the BMFTR (grant number
13FH097KX0).
Declaration of Competing Interest
The authors declare that they have no known competing financial interests or personal relationships that could have appeared to influence the work reported in this paper.
Declaration of Generative AI and AI-Assisted
Technologies in the Manuscript Preparation Process
The authors conceived the study, acquired and prepared the ColorChecker data and conducted initial investigations, including comparisons of feature edge maps, Boundary IoU, effective resolution and simple qualitative tests of shifted wavelengths. ChatGPT (OpenAI, GPT-5.1) served as a sounding board for discussing initial ideas, visualizing first concepts and refining final text. Claude (Anthropic, Claude Opus 5.5
via Claude Code) was subsequently used to develop these investigations into a systematic and reproducible experimental framework with more readable, structured and consistent code.
This included substantial generation and refinement of analysis and figure code, evaluation methods and statistical procedures, literature search assistance as well as drafting, translation and editing of manuscript text across all sections based on the authors’ material. The contamination width 𝑤 90 was proposed by Claude in response to the authors’ aim of quantifying contamination near boundaries. The AI-assisted development involved extensive iterative exchanges in which the authors

VII. Conclusion
Image-guided, training-free upsampling makes frozen
CARL features markedly more edge-accurate. With JBU and the hyperspectral image as the guide, the Boundary IoU
increases by 55 pp and the effective resolution of the features improves from about one token to a few pixels, without additional labels or further training. The gain replicates on a third recording, persists under wavelength errors and other band sets as long as the head matches the configuration and also appears qualitatively on tissue. For applications that use
CARL with linear heads without fine-tuning, we recommend guided upsampling instead of bilinear interpolation, with JBU

11

specified requirements, critically evaluated suggestions, revised or rejected proposed approaches and requested additional checks. Throughout this process, the authors maintained the study’s scientific focus and ensured that the work formed a coherent, self-contained contribution. The authors also devoted substantial time and effort to iteratively reviewing and refining the manuscript and figures, with particular attention to readability and visual clarity. The authors defined the scope of the study, made the final scientific decisions and take full responsibility for the results and the manuscript.

[23] M. Drusch et al., “Sentinel-2: ESA’s optical high-resolution mission for GMES operational services,” Remote Sens. Environ., vol. 120, pp. 25–36, 2012.

References
[1] A. Baumann et al., “CARL: Camera-agnostic representation learning for spectral image analysis,” in Proc. Int. Conf. Learn. Represent. (ICLR),
2026.
[2] D. Hong et al., “SpectralGPT: Spectral remote sensing foundation model,” IEEE Trans. Pattern Anal. Mach. Intell., vol. 46, no. 8, pp. 5227–5244, 2024.
[3] D. Wang et al., “HyperSIGMA: Hyperspectral intelligence comprehension foundation model,” IEEE Trans. Pattern Anal. Mach. Intell., vol. 47, no. 8, pp. 6427–6444, 2025.
[4] Z. Xiong et al., “Neural plasticity-inspired multimodal foundation model for Earth observation,” arXiv:2403.15356, 2024.
[5] S. Fu et al., “FeatUp: A model-agnostic framework for features at any resolution,” in Proc. ICLR, 2024.
[6] H. Huang et al., “LoftUp: Learning a coordinate-based feature upsampler for vision foundation models,” in Proc. IEEE/CVF Int. Conf. Comput.
Vis. (ICCV), 2025, pp. 9913–9923.
[7] P. Couairon et al., “JAFAR: Jack up any feature at any resolution,” in
Adv. Neural Inf. Process. Syst. (NeurIPS), 2025.
[8] J. Kopf, M. F. Cohen, D. Lischinski, and M. Uyttendaele, “Joint bilateral upsampling,” ACM Trans. Graph., vol. 26, no. 3, Art. no. 96, 2007.
[9] K. He, J. Sun, and X. Tang, “Guided image filtering,” IEEE Trans.
Pattern Anal. Mach. Intell., vol. 35, no. 6, pp. 1397–1409, 2013.
[10] S. Patzke, C. Wöhle

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

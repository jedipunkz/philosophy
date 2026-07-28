---
source: "https://www.semanticscholar.org/paper/2d2085208f33bd72e001e90f67489296e108e097"
title: "SOCRATES: an online tool leveraging a social contact data sharing initiative to assess mitigation strategies for COVID-19"
author: "L. Willem, Thang Van Hoang, S. Funk, P. Coletti, P. Beutels, N. Hens"
year: "2020"
publication: "BMC Research Notes"
download: "https://www.medrxiv.org/content/medrxiv/early/2020/03/19/2020.03.03.20030627.full.pdf"
pdf: "https://www.medrxiv.org/content/medrxiv/early/2020/03/19/2020.03.03.20030627.full.pdf"
captured_at: "2026-07-28T12:30:44Z"
updated_at: "2026-07-28T12:30:44Z"
capture_tool: "scrapem"
source_name: "semanticscholar"
keyword: "ソクラテス"
query: "Socrates"
tags:
  - "古代哲学"
  - "倫理学"
  - "問答法"
status: raw
---

# SOCRATES: an online tool leveraging a social contact data sharing initiative to assess mitigation strategies for COVID-19

- 著者: L. Willem, Thang Van Hoang, S. Funk, P. Coletti, P. Beutels, N. Hens
- 年: 2020
- 掲載情報: BMC Research Notes
- 情報源: [semanticscholar](https://www.semanticscholar.org/paper/2d2085208f33bd72e001e90f67489296e108e097)
- ダウンロード: https://www.medrxiv.org/content/medrxiv/early/2020/03/19/2020.03.03.20030627.full.pdf
- PDF: https://www.medrxiv.org/content/medrxiv/early/2020/03/19/2020.03.03.20030627.full.pdf

## Obsidian Links

- 研究動向: [[ソクラテス-現代研究動向]]
- キーワード: [[ソクラテス]]
- 関連分野: [[古代哲学]]
- 関連分野: [[倫理学]]
- 関連分野: [[問答法]]
- 関連タグ: #古代哲学 #倫理学 #問答法

## Abstract

Establishing a social contact data sharing initiative and an interactive tool to assess mitigation strategies for COVID-19. We organized data sharing of published social contact surveys via online repositories and formatting guidelines. We analyzed this social contact data in terms of weighted social contact matrices, next generation matrices, relative incidence and R0\documentclass[12pt]{minimal} \usepackage{amsmath} \usepackage{wasysym} \usepackage{amsfonts} \usepackage{amssymb} \usepackage{amsbsy} \usepackage{mathrsfs} \usepackage{upgreek} \setlength{\oddsidemargin}{-69pt} \begin{document}$$_{0}$$\end{document}. We incorporated location-specific physical distancing measures (e.g. school closure or at work) and capture their effect on transmission dynamics. All methods have been implemented in an online application based on R Shiny and applied to COVID-19 with age-specific susceptibility and infectiousness. Using our online tool with the available social contact data, we illustrate that physical distancing could have a considerable impact on reducing transmission for COVID-19. The effect itself depends on assumptions made about disease-specific characteristics and the choice of intervention(s).

## Citation

DOI: 10.1186/s13104-020-05136-9

## PDF Text

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

SOCRATES: An online tool leveraging a social contact data sharing initiative to assess mitigation strategies for COVID-19
Lander Willem1 , Thang Van Hoang2 , Sebastian Funk3 , Pietro Coletti2 , Philippe Beutels1,4 , Niel Hens1,2
1 Centre for Health Economic Research and Modelling Infectious Diseases, University of Antwerp, Belgium.
2 Interuniversity Institute of Biostatistics and statistical Bioinformatics, Data Science Institute, Hasselt University, Belgium
3 Centre for the Mathematical Modelling of Infectious Diseases, London School of Hygiene & Tropical
Medicine, UK.
4 School of Public health and Community Medicine, University of New South Wales, Australia.
* lander.willem@uantwerp.be

Abstract

1

Objective Establishing a social contact data sharing initiative and an interactive tool to assess mitigation

2

strategies for COVID-19.

3

Results We organized data sharing of published social contact surveys via online repositories and formatting

4

guidelines. We analyzed this social contact data in terms of weighted social contact matrices, next generation

5

matrices, relative incidence and R0 . We incorporated location-specific isolation measures (e.g. school closure or

6

telework) and capture their effect on transmission dynamics. All methods have been implemented in an online

7

application based on R Shiny and applied to COVID-19 with age-specific susceptibility and infectiousness.

8

Using our online tool with the available social contact data, we illustrate that social distancing could have a

9

considerable impact on reducing transmission for COVID-19. The effect itself depends on assumptions made

10

about disease-specific characteristics and the choice of intervention(s).

11

Keywords social contact data, user interface, transmission dynamics, infectious diseases, epidemics, social

12

distancing, behavioral changes, data sharing initiative, open-source, COVID-19

13

Introduction

14

Given the pandemic of SARS-CoV-2, which causes COVID-19 disease, it is of great importance to consider

15

intervention strategies to slow down SARS-CoV-2 spread, and thus decrease surge capacity problems arising

16

to health care provision and essential supplies [1]. Social distancing on a large scale, first at the epicentre of

17

the outbreak in Wuhan, and later in other locations was shown to slow down SARS-CoV-2 spread (e.g. in

18

Shanghai) [2]).

19

NOTE: This preprint reports new research that has not been certified
1 by peer review and should not be used to guide clinical practice.

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

Social contact surveys have proven to be an invaluable source of information about how people mix in the

20

population [3, 4, 5] and explained close contact infectious disease data well [6, 7, 8]. For example, adapted

21

social mixing during the the A(H1N1)v2009 pandemic was fundamental to reproduce the observed incidence

22

patterns [9]. In terms of prevention strategies, social contact data from the POLYMOD project [4] have been

23

used to quantify the impact of school closure on the spread of airborne infections [10]. This was done by

24

comparing the basic reproduction number R0 , or the average number of secondary infections caused by a

25

single infectious individual in a completely susceptible population, derived from mixing patterns observed on

26

weekends or during a holiday period with those derived from mixing patterns observed on weekdays.

27

In this research note, we highlight a social contact data sharing initiative and present an online tool to

28

facilitate data access and analyses. Social distancing measures can be mimicked with this tool by excluding

29

the contribution of mixing patterns at specific locations to investigate the impact on disease transmission and

30

guide policy makers. As a case study, we exploit our application to quantify the potential impact of school

31

closure and a shift of workers from a common workplace to teleworking at home in light of COVID-19.

32

Main text

33

Methods

34

Following a systematic literature review [3], corresponding authors were contacted to share their data subject

35

to ethical approvals and GDPR compliance. All data have been refactored according to guidelines we developed

36

during a Social Contact Data Hackaton in 2017 as part of the TransMID project. Each survey is split into

37

multiple files to capture participant, contact, survey day, household and time-use data. For each data type,

38

there is one “common” file and one “extra” file in which more specific variables related to the survey are

39

included. Each data set contains a dictionary to interpret the columns (see socialcontactdata.org for more

40

information).

41

To extrapolate survey data to the country level and obtain social contact rates on a weekly basis, we incorporate

42

participant weights accounting for age and the number of observations during week (5/7) and weekend (2/7)

43

days. We use the United Nation’s World Population Prospects [11] as reference and constrain weights to a

44

maximum of 3 to limit the influence of single participants. The social contact matrix mij can be estimated

45

by:

46

PTi mij =

d t=1 wit yijt
, d
t=1 wit

PTi

(1)

d where wit denotes the weight for participant t of age i who was surveyed on day type d ∈ {weekday, weekend}

47

and yijt denotes the reported number of contacts made by participant t of age i with someone of age j. By

48

nature, contacts are reciprocal and thus mij Ni should be equal to mji Nj . To resolve differences in reporting,

49

reciprocity can be imposed by:

50

mreciprocal
=
ij

mij Ni + mji Nj
,
2Ni
2

(2)

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

with Ni and Nj the population size in age class i and j, respectively [12]. This reciprocal behavior might not

51

be valid for specific contact types, e.g. contacts at work for retail workers are most likely not contacts at work

52

for their customers.

53

Transmission dynamics can be represented by the next generation matrix G with elements gij that indicate the

54

average number of secondary infections in age class i through the introduction of a single infectious individual

55

of age class j into a fully susceptible population [13]. The next generation matrix is defined by:

56

G = DM q,

(3)

with D the mean duration of infectiousness, M the contact matrix and q a proportionality factor [10, 8]. The

57

proportionality factor q combines several disease-specific characteristics that are related to susceptibility and

58

infectiousness. Equation 3 can be reformulated as:

59

gij = D ∗ mij ∗ si ∗ kj ∗ q̂,

(4)

where si denotes the susceptibility of age group i, kj the infectiousness of age group j and q̂ other disease-

60

specific factors. The leading right eigenvector of G is proportional to the expected incidence by age and R0

61

can be calculated as the dominant eigenvalue of G [4].

62

To evaluate intervention strategies, we focus on the relative impact of adjusted social contact patterns on R0

63

in line with the so-called social contact hypothesis [6] by cancelling disease specific features:

64

max(eigen(Ma ∗ S ∗ K))
max(eigen(DMa q)
R0a
=
,
=
R0b max(eigen(DMb q)
max(eigen(Mb ∗ S ∗ K))

(5)

where indices a and b refer to the different conditions, and S and K account for age-specific susceptibility and

65

infectiousness, respectively [10]. Social distancing can be evaluated by the elimination or reduction of location-

66

specific subsets of the social contact data. Contacts reported at multiple locations are assigned to a single

67

location in the following hierarchical order: home, work, school, transport, leisure and other locations. Firstly,

68

we simulate school closure by excluding all contacts reported at school before calculating mij . Secondly, we

69

consider an increase in telework to proportion ptarget telework , by accounting for the observed social contacts at work observed
Mwork and the observed proportion of telework pobserved telework :

observed all
Mwork
= Mwork
∗ (1 − pobserved telework ), target all
Mwork
= Mwork
∗ (1 − ptarget telework ),
(1 − ptarget target observed telework )
Mwork
= Mwork
∗
,
(1 − pobserved telework )

3

71

(6)
(7)
(8)

To combine the effect of telework and school closure, the social contact matrix M is calculated as:

target
M = Mhome + Mwork
+ (Mschool ∗ 0) + Mtransport + Mleisure + Mother

70

72

(9)

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

We developed an interactive application to access and analyze social contact data based on R packages shiny

73

[14] and socialmixr [15]. The user interface enables the selection of country-specific data, age categories,

74

type of day, contact duration, intensity and gender. Using a selection box, the user can opt to disable the

75

assumption of reciprocity and participant weights or to include age-specific transmission parameters. The user

76

can also enable reactive strategies such as school closure and increase the level of telework. Please note that

77

our proportion of telework can only increase given the specified observed proportion.

78

The user interface contains a plot of the social contact matrix and the principal results of the social contact

79

analysis: M , relative incidences, participant statistics and the reference demography. Relative R0 and M

80

ratios are printed if reactive strategies are selected.

81

As COVID-19 case study, we estimate the effect of school closure and telework on disease transmission dy-

82

namics. In order to do this, we use 3 age classes: 0–18 years, 19–60 years and over 60 years of age. For each

83

We fix pobserved telework at 5%, in line

84

with European observations [16, 17] and capture transmission dynamics with 20%, 35% and 50% telework,

85

with and without school closure. As proof of concept, we include the scenario where children are less vulnerable

86

compared to elderly (si = kj = (0.5, 1, 1.5)), instead of uniform susceptibility and infectiousness.

87

Results

88

The socialcontactdata.org initiative, status 12th March 2020, includes data for Belgium, Finland, Germany,

89

Italy, Luxembourg, Netherlands, Poland and the UK from POLYMOD [4], as well as data from other studies

90

on social mixing in France [18], China [19], Hong Kong [20], Peru [21], UK [22], Russia [23], Zimbabwe [24],

91

Vietnam [25], South Africa and Zambia [26]. All data are available on Zenodo [27, 28, 29, 30, 31, 32, 33, 34,

92

35, 36] and can be retrieved within R using the socialmixr package.

93

The SOcial Contact RATES (Socrates) data tool [37, 38] enables quick and convenient generation of social

94

contact matrices, relevant for the spread of infectious diseases. Figure 1 presents a screenshot of the user

95

interface. The potential of using social contact patterns to simulate infectious disease transmission are endless,

96

and we hope with this initiative to support data-driven modeling endeavors. The survey data from France

97

and Zimbabwe contain multiple days per participant, hence we included only the first day for each participant

98

to minimize the effect of reporting fatigue.

99

We demonstrate the effect of telework and school closure on R0 in Figure 2. If we assume uniform susceptibility

100

and infectiousness, we predict for most countries a 10% decrease in R0 with a telework proportion of 50%.

101

For Poland and Hong Kong, the reduction is slightly higher. The analysis for Peru shows little impact of

102

telework since only few contacts were reported “at work”, whereas a substantial proportion of contacts was

103

reported at the market or street. Cultural differences in how “at work” is understood should be considered

104

when interpreting results. The estimated R0 reduction due to school closure is more country-specific, e.g. 10%

105

reduction for Belgium and Vietnam, but 20% for Italy, Luxembourg and France. If we assume that elderly are

106

more vulnerable compared to children, as is the case for COVID-19 [39] the impact of school closure decreases

107

dramatically. The positive effect of telework on R0 remains the same or increases.

108

The predicted relative incidences, as presented in Figure 3, highlight the impact of school closure compared to

109

an increase in telework by age. The relative incidence in people 18–60 years of age decreases with an increasing

110

country, we calculate contact rates after excluding data from holiday periods.

4

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

proportion of telework, which is of interest if this age group is more vulnerable compared to children. The

111

relative incidence in the age group above 60 years of age increases in all situations compared to no intervention.

112

This does not imply that the absolute number of cases in this age group would rise.

113

Limitations

114

Most survey designs were based on the POLYMOD survey though each survey had additional features and

115

objectives which provide useful additional information. At the moment, we do not capture the full potential

116

of each data set yet. Our social contact analyses focus only on adapting school and work contacts and does

117

not capture compensation behavior due to not being at school or work. This might be valid for a pandemic

118

situation but not for regular (school) holidays. Social distancing due to (pandemic) scares are also not included

119

yet.

120

The current application contains a local version of each data set, with some additional data reformatting. Our

121

aim is to enable a direct link to Zenodo repositories. Note that social contact surveys are available on Zenodo

122

but not included in Socrates. E.g., the data from China [19] contains grouped contacts, which require different

123

methodology. We omitted data from the UK [22], Zambia and South Africa [24] from our case study because

124

only infants or adults were recruited. Data from Zimbabwe are (temporary) excluded due to refactoring issues

125

with the contact location, which are key in this case-study.

126

Note that we will continue to develop this open-source tool [38] and thus the input/output/plots/scenarios

127

might change in future editions.

128

Abbreviations

129

Socrates: SOcial Contact RATES

130

Declarations

131

Ethics approval and consent to participate

132

The social contact data sharing initiative is part of the ERC consolidator grant “TransMID” which received

133

ethical approval from the Hasselt University Medical Ethical Committee (CME2016/618)

134

Consent for publication

135

Not applicable.

136

Availability of data and materials

137

All data sets are available on Zenodo [27, 28, 29, 30, 31, 32, 33, 34, 35]. We also share all R-code on Zenodo

138

[38].

139

5

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

Competing interests

140

The authors declare no competing interests.

141

Funding

142

LW gratefully acknowledges funding from the Research Foundation Flanders (Grant number 1234620N). This

143

work is part of a project that has received funding from the European Research Council (ERC) under the

144

European Union’s Horizon 2020 research and innovation programme (grant agreement 682540 — TransMID)

145

(TVH, PC and NH). This work is partially funded by the Epipose project from the European Union’s SC1-

146

PHE-CORONAVIRUS-2020 programme (101003688). SF was funded by a Wellcome Trust Senior Research

147

Fellowship (210758/Z/18/Z)

148

Authors’ contributions

149

NH conceived the study. TVH and PC collected and reformatted social contact data. LW and NH wrote a

150

first draft of the paper. LW, TVH, SF and NH developed the online tool. All authors contributed to the final

151

version of the paper and approved the final manuscript

152

Acknowledgements

153

We acknowledge support from the Antwerp Study Center for Infectious Diseases (ASCID) and are thankful

154

for all survey data that have been made open-source.

155

References

156

[1] Gilbert, M., Pullano, G., Pinotti, F., Valdano, E., Poletto, C., Boëlle, P.-Y., D’Ortenzio, E., Yazdan-

157

panah, Y., Eholie, S.P., Altmann, M., et al.: Preparedness and vulnerability of African countries against

158

importations of COVID-19: a modelling study. Lancet (2020)

159

[2] Lu, H., Ai, J., Shen, Y., Li, Y., Li, T., Zhou, X., Zhang, H., Zhang, Q., Ling, Y., Wang, S., et al.: A

160

descriptive study of the impact of diseases control and prevention on the epidemics dynamics and clini-

161

cal features of SARS-CoV-2 outbreak in Shanghai, lessons learned for metropolis epidemics prevention.

162

medRxiv (2020)

163

[3] Hoang, T.V., Coletti, P., Melegaro, A., Wallinga, J., Grijalva, C.G., Edmunds, J.W., Beutels, P., Hens,

164

N.: A systematic review of social contact surveys to inform transmission models of close-contact infections.

165

Epidemiology 30(5), 723–736 (2019)

166

[4] Mossong, J., Hens, N., Jit, M., Beutels, P., Auranen, K., Mikolajczyk, R., Massari, M., Salmaso, S.,

167

Tomba, G.S., Wallinga, J., Heijne, J., Sadkowska-Todys, M., Rosinska, M., Edmunds, W.J.: Social

168

contacts and mixing patterns relevant to the spread of infectious diseases. PLoS Med 5(3), 74 (2008)

169

[5] Willem, L., Van Kerckhove, K., Chao, D.L., Hens, N., Beutels, P.: A nice day for an infection? weather

170

conditions and social contact patterns relevant to influenza transmission. PLoS One 7(11) (2012)

6

171

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

[6] Wallinga, J., Teunis, P., Kretzschmar, M.: Using data on social contacts to estimate age-specific trans-

172

mission parameters for respiratory-spread infectious agents. Am J Epidemiol 164(10), 936–944 (2006)

173

[7] Ogunjimi, B., Hens, N., Goeyvaerts, N., Aerts, M., Van Damme, P., Beutels, P.: Using empirical social

174

contact data to model person to person infectious disease transmission: an illustration for varicella. Math

175

Biosci 218(2), 80–87 (2009)

176

[8] Goeyvaerts, N., Hens, N., Ogunjimi, B., Aerts, M., Shkedy, Z., Van Damme, P., Beutels, P.: Estimating

177

infectious disease parameters from data on social contacts and serological status. J R Stat Soc Ser C Appl

178

Stat 59(2), 255–277 (2010)

179

[9] Eames, K., Tilston, N., White, P., Adams, E., Edmunds, W.: The impact of illness and the impact of school closure on social contact patterns. Health Technol Assess 14(34), 267–312 (2010)

180
181

[10] Hens, N., Ayele, G.M., Goeyvaerts, N., Aerts, M., Mossong, J., Edmunds, J.W., Beutels, P.: Estimating

182

the impact of school closure on social mixing behaviour and the transmission of close contact infections

183

in eight European countries. BMC Infect Dis 9(1), 187 (2009)

184

[11] Population Division, Department of Economic and Social Affairs, U.N.: wpp2015: World Population
Prospects 2015. The Comprehensive R Archive Network (2019)
[12] Held, L., Hens, N., D O’Neill, P., Wallinga, J.: Handbook of Infectious Disease Data Analysis. Chapman and Hall/CRC, US (2019)

185
186

187
188

[13] Diekmann, O., Heesterbeek, J.A.P., Metz, J.A.J.: On the definition and the computation of the basic

189

reproduction ratio R0 in models for infectious diseases in heterogeneous populations. J Math Biol 28(4),

190

365–382 (1990)

191

[14] Chang, W., Cheng, J., Allaire, J., Xie, Y., McPherson, J., et al.: Shiny: web application framework for r. R package version 1(5) (2017)

192
193

[15] Funk, S.: socialmixr: Social Mixing Matrices for Infectious Disease Modelling. The Comprehensive R
Archive Network (2020)

194
195

[16] EUROSTAT: Your Key to European Statistics. https://ec.europa.eu/eurostat/data/

196

[17] Federale Overheidsdienst Mobiliteit en Vervoer: Kerncijfers Telewerk en mobiliteit in België. Wettelijk

197

depot: D/2018/13.831/4 (2018)

198

[18] Béraud, G., Kazmercziak, S., Beutels, P., Levy-Bruhl, D., Lenne, X., Mielcarek, N., Yazdanpanah, Y.,

199

Boëlle, P.-Y., Hens, N., Dervaux, B.: The French connection: the first large population-based contact

200

survey in France relevant for the spread of infectious diseases. PLoS One 10(7) (2015)

201

[19] Zhang, J., Klepac, P., Read, J.M., Rosello, A., Wang, X., Lai, S., Li, M., Song, Y., Wei, Q., Jiang, H., et

202

al.: patterns of human social contact and contact with animals in Shanghai, China. Sci Rep 9(1), 1–11

203

(2019)

204

[20] Leung, K., Jit, M., Lau, E.H., Wu, J.T.: Social contact patterns relevant to the spread of respiratory infectious diseases in Hong Kong. Sci Rep 7(1), 1–12 (2017)
7

205
206

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

[21] Grijalva, C.G., Goeyvaerts, N., Verastegui, H., Edwards, K.M., Gil, A.I., Lanata, C.F., Hens, N., et al.: A

207

household-based study of contact networks relevant for the spread of infectious diseases in the highlands

208

of Peru. PLoS One 10(3) (2015)

209

[22] van Hoek, A.J., Andrews, N., Campbell, H., Amirthalingam, G., Edmunds, W.J., Miller, E.: The social

210

life of infants in the context of infectious disease transmission; social contacts and mixing patterns of the

211

very young. PLoS One 8(10) (2013)

212

[23] Litvinova, M., Liu, Q.-H., Kulikov, E.S., Ajelli, M.: Reactive school closure weakens the network of social

213

interactions and reduces the spread of influenza. Proc Natl Acad Sci U S A 116(27), 13174–13181 (2019)

214

[24] Melegaro, A., Del Fava, E., Poletti, P., Merler, S., Nyamukapa, C., Williams, J., Gregson, S., Manfredi,

215

P.: Social contact structures and time use patterns in the Manicaland Province of Zimbabwe. PLoS One

216

12(1) (2017)

217

[25] Horby, P., Thai, P.Q., Hens, N., Yen, N.T.T., Mai, L.Q., Thoang, D.D., Linh, N.M., Huong, N.T.,

218

Alexander, N., Edmunds, W.J., et al.: Social contact patterns in vietnam and implications for the control

219

of infectious diseases. PLoS One 6(2) (2011)

220

[26] Dodd, P.J., Looker, C., Plumb, I.D., Bond, V., Schaap, A., Shanaube, K., Muyoyeta, M., Vynnycky, E.,

221

Godfrey-Faussett, P., Corbett, E.L., et al.: Age-and sex-specific social contact patterns and incidence of

222

mycobacterium tuberculosis infection. Am J Epidemiol 183(2), 156–166 (2016)

223

[27] Grijalva, C.G., Goeyvaerts, N., Verastegui, H., Edwards, K.M., Gil, A.I., Lanata, C.F., Hens, N., et al.:
Peruvian Social Contact Data. https://doi.org/10.5281/zenodo.1215891

224
225

[28] Mossong, J., Hens, N., Jit, M., Beutels, P., Auranen, K., Mikolajczyk, R., Massari, M., Salmaso, S.,

226

Tomba, G.S., Wallinga, J., et al.: POLYMOD Social Contact Data. https://doi.org/10.5281/zenodo.

227

1215899

228

[29] Béraud, G., Kazmercziak, S., Beutels, P., Levy-Bruhl, D., Lenne, X., Mielcarek, N., Yazdanpanah, Y.,

229

Boëlle, P.-Y., Hens, N., Dervaux, B.: France Social Contact Data. https://doi.org/10.5281/zenodo.

230

1158452

231

[30] Ajelli, M., Litvinova, M.: Russian Contact Matrices by Age. https://doi.org/10.5281/zenodo.
3232929

232
233

[31] Horby, P., Thai, P.Q., Hens, N., Yen, N.T.T., Mai, L.Q., Thoang, D.D., Linh, N.M., Huong, N.T.,

234

Alexander, N., Edmunds, W.J., et al.: Social Contact Data for Vietnam. https://doi.org/10.5281/

235

zenodo.1289474

236

[32] Leung, K., Jit, M., Lau, E.H., Wu, J.T.: Social Contact Data for Hong Kong. https://doi.org/10.
5281/zenodo.1165562

237
238

[33] Melegaro, A., Del Fava, E., Poletti, P., Merler, S., Nyamukapa, C., Williams, J., Gregson, S., Manfredi,
P.: Zimbabwe Social Contact Data. https://doi.org/10.5281/zenodo.1251944

8

239
240

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

[34] Dodd, P.J., Looker, C., Plumb, I.D., Bond, V., Schaap, A., Shanaube, K., Muyoyeta, M., Vynnycky, E.,

241

Godfrey-Faussett, P., Corbett, E.L., et al.: Social Contact Data for Zambia and South Africa (CODA

242

Dataset). https://doi.org/10.5281/zenodo.2548693

243

[35] Zhang, J., Klepac, P., Read, J.M., Rosello, A., Wang, X., Lai, S., Li, M., Song, Y., Wei, Q., Jiang, H., et al.: Social Contact Data for China Mainland. https://doi.org/10.5281/zenodo.3516113
[36] van Hoek, A.J., Andrews, N., Campbell, H., Amirthalingam, G., Edmunds, W.J., Miller, E.: Social
Contact Data for UK. doi:10.5281/zenodo.1409507. https://doi.org/10.5281/zenodo.1409507
[37] Social Contact Rates (SOCRATES) Data Tool: as part of the socialcontactdata.org initiative. TransMID.
http://www.socialcontactdata.org

244
245

246
247

248
249

[38] Willem, L., Hoang, V.T., Funk, S., Coletti, P., Beutels, P., Hens, N.: Social Contact Rates (SOCRATES)
Data Tool (v1.5). doi:10.5281/zenodo.3706788. https://doi.org/10.5281/zenodo.3706788
[39] Guan, W.-j., Ni, Z.-y., Hu, Y., Liang, W.-h., Ou, C.-q., He, J.-x., Liu, L., Shan, H., Lei, C.-l., Hui, D.S., et al.: Clinical characteristics of coronavirus disease 2019 in China. N Engl J Med (2020)

Figures

250
251

252
253

254

9

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

Figure 1: Screenshot of the online SOCRATES application [37]. The user interface enables the selection of country data in combination with temporal and contact features. The social contact matrix is shown on the right in addition to principal results ans statistics. When users include reactive measures such as school closure and/or increased teleworking, the R0 ratio is added to the output (not shown).

10

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

0.9

0.9

0.8
0.7

0.8
0.7

0.6

0.6

0.5

0.5
20

30

40

R0 ratio

0.9

50

0.7
0.6
0.5

10

20

30

40

50

10

40

Telework (%)

Germany (Mossong 2008)

Hong Kong (Leung 2017)

Italy (Mossong 2008)

0.9

0.8
0.7

0.8
0.7

0.6

0.6

0.5

0.5
20

30

40

R0 ratio

1.0

0.9
R0 ratio

1.0

50

0.7
0.6
0.5

10

20

30

40

50

10

Netherlands (Mossong 2008)

0.9

0.8
0.7

0.6

0.6

0.5

0.5
30

40

R0 ratio

0.9
R0 ratio

0.9

0.7

50

40

0.8
0.7
0.6
0.5

10

20

30

40

50

10

20

30

40

Telework (%)

Telework (%)

Telework (%)

Poland (Mossong 2008)

Russia (Litvinova 2019)

United Kingdom (Mossong 2008)
1.0

0.9

0.9

0.7

0.8
0.7

0.6

0.6

0.5

0.5
20

30

40

50

Telework (%)

R0 ratio

1.0

0.9
R0 ratio

1.0

0.8

50

Peru (Grijalva 2015)
1.0

0.8

30
Telework (%)

1.0

20

20

Telework (%)

Luxembourg (Mossong 2008)

50

0.8

1.0

10

30

Telework (%)

0.9

10

20

Telework (%)

Telework (%)

R0 ratio

0.8

1.0

10

R0 ratio

France (Beraud 2015)
1.0

10

R0 ratio

Finland (Mossong 2008)
1.0

R0 ratio

R0 ratio

Belgium (Mossong 2008)
1.0

50

0.8
0.7
0.6
0.5

10

20

30

40

Telework (%)

50

10

20

30

40

50

Telework (%)

Vietnam (Horby 2007)
1.0

telework only

R0 ratio

0.9
0.8
0.7

telework and school closure

0.6

Age breaks (0,18,60)
Susceptibility (1,1,1)
Infectiousness (1,1,1)
Susceptibility (0.5,1,1.5)
Infectiousness (0.5,1,1.5)

0.5
10

20

30

40

50

Telework (%)

Figure 2: Predicted R0 ratio by country due to increased teleworking and/or school closure. The reference proportion for telework is fixed to 5% to present a relative increase in telework. The impact on R0
is shown with uniform susceptible and infectiousness parameters (1,1,1) and when children are less vulnerable compared to elderly (0.5,1,1.5).

11

medRxiv preprint doi: https://doi.org/10.1101/2020.03.03.20030627; this version posted March 19, 2020. The copyright holder for this preprint
(which was not certified by peer review) is the author/funder, who has granted medRxiv a license to display the preprint in perpetuity.
It is made available under a CC-BY-NC-ND 4.0 International license .

Finland (Mossong 2008)

0.8

0.8

0.8

0.6
0.4

0.6
0.4

0.2

0.2

0.0

0.0
20

35

50

5

20

35

Relative incidence

1.0

5

50

20

35

50

5

20

35

50

5

Hong Kong (Leung 2017)

0.8

0.8

0.8

0.4
0.2

Relative incidence

1.0

0.6

0.6
0.4
0.2

0.0
50

5

20

35

50

Luxembourg (Mossong 2008)

20

35

50

5

20

35

50

5

Netherlands (Mossong 2008)

0.8

0.6
0.4
0.2
0.0

35

50

5

20

35

Relative incidence

0.8
Relative incidence

0.8

20

50

20

35

50

5

20

35

50

5

Russia (Litvinova 2019)

0.8
Relative incidence

0.8
Relative incidence

0.8

0.6
0.4
0.2

50

5

35

20

35

50

Telework (%)

50

5

20

35

50

0.6
0.4
0.2

0.0
35

20

United Kingdom (Mossong 2008)
1.0

20

50

Telework (%)

1.0

0.0

35

0.0
5

Poland (Mossong 2008)

0.2

20

0.4

1.0

0.4

5

0.6

Telework (%)

0.6

50

0.2

Telework (%)

5

35

Peru (Grijalva 2015)
1.0

5

20

Telework (%)

1.0

0.0

50

0.4

1.0

0.2

35

0.6

Telework (%)

0.4

20

0.0
5

Telework (%)

0.6

5

0.2

0.0
35

50

Italy (Mossong 2008)

1.0

20

35

Telework (%)

1.0

5

20

Telework (%)

Relative incidence

Relative incidence

0.4

0.0
5

Germany (Mossong 2008)

Relative incidence

0.6

0.2

Telework (%)

Relative incidence

France (Beraud 2015)

1.0

Relative incidence

Relative incidence

Belgium (Mossong 2008)
1.0

0.0
5

20

35

50

5

20

35

50

5

Telework (%)

20

35

50

5

20

35

50

Telework (%)

Vietnam (Horby 2007)
1.0

Age group (years)
Relative incidence

0.8

[0,18)
[18,60)
60+

0.6
0.4
0.2

Age breaks (0,18,60)
Susceptibility (1,1,1)
Infectiousness (1,1,1)

with school closure

0.0
5

20

35

50

5

20

35

50

Telework (%)

Figure 3: Predicted age-specific relative incidence by country with increased teleworking and/or school closure. The reference proportion for telework is fixed to 5% to present a relative increase in telework.
The analysis presented here does not account for age-specific vulnerability.

12

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

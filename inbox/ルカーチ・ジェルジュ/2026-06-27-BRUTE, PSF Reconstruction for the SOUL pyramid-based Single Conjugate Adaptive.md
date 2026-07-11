---
source: "https://arxiv.org/abs/2209.03278v1"
title: "BRUTE, PSF Reconstruction for the SOUL pyramid-based Single Conjugate Adaptive Optics facility of the LBT"
author: "Carmelo Arcidiacono, Andrea Grazian, Anita Zanella, Benedetta Vulcani, Elisa Portaluri, Fernando Pedichini, Marco Gullieuszik, Matteo Simioni, Roberto Piazzesi, Roland Wagner, Enrico Pinna, Guido Agapito, Fabio Rossi, Cedric Plantet"
year: "2022"
publication: "arXiv preprint / astro-ph.IM"
download: "https://arxiv.org/pdf/2209.03278v1"
pdf: "https://arxiv.org/pdf/2209.03278v1"
captured_at: "2026-06-27T06:15:46Z"
updated_at: "2026-06-27T06:15:46Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács Soul and Form"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# BRUTE, PSF Reconstruction for the SOUL pyramid-based Single Conjugate Adaptive Optics facility of the LBT

- 著者: Carmelo Arcidiacono, Andrea Grazian, Anita Zanella, Benedetta Vulcani, Elisa Portaluri, Fernando Pedichini, Marco Gullieuszik, Matteo Simioni, Roberto Piazzesi, Roland Wagner, Enrico Pinna, Guido Agapito, Fabio Rossi, Cedric Plantet
- 年: 2022
- 掲載情報: arXiv preprint / astro-ph.IM
- 情報源: [arxiv](https://arxiv.org/abs/2209.03278v1)
- ダウンロード: https://arxiv.org/pdf/2209.03278v1
- PDF: https://arxiv.org/pdf/2209.03278v1

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

The astronomical applications greatly benefit from the knowledge of the instrument PSF. We describe the PSF Reconstruction algorithm developed for the LBT LUCI instrument assisted by the SOUL SCAO module. The reconstruction procedure considers only synchronous wavefront sensor telemetry data and a few asynchronous calibrations. We do not compute the Optical Transfer Function and corresponding filters. We compute instead a temporal series of wavefront maps and for each of these the corresponding instantaneous PSF. We tested the algorithm both in laboratory arrangement and in the nighttime for different SOUL configurations, adapting it to the guide star magnitudes and seeing conditions. We nick-named it "BRUTE", Blind Reconstruction Using TElemetry, also recalling the one-to-one approach, one slope-to one instantaneous PSF the algorithm applies.

## PDF Text

BRUTE, PSF Reconstruction for the SOUL pyramid-based
Single Conjugate Adaptive Optics facility of the LBT
Carmelo Arcidiaconoa,g,* , Andrea Graziana , Anita Zanellaa , Benedetta Vulcania , Elisa
Portaluric,a,g , Fernando Pedichinif,g , Marco Gullieuszika , Matteo Simionia , Roberto Piazzesif,g ,
Roland Wagnerb,e , Enrico Pinnad,g , Guido Agapitod,g , Fabio Rossid,g , and Cedric Plantetd,g

arXiv:2209.03278v1 [astro-ph.IM] 7 Sep 2022

a

INAF - Osservatorio Astronomico di Padova, Vicolo dell’Osservatorio 5, Padova, Italy, I-35122
b
Industrial Mathematics Institute, Johannes Kepler University Linz, Altenberger Strasse 69,
Linz, Austria, 4040
c
INAF - Osservatorio Astronomico d’Abruzzo, Via Mentore Maggini, Teramo, Italy, I-64100
d
INAF - Osservatorio Astrofisico di Arcetri, Via E. Fermi 5, Firenze, Italy, I-50125
e
RICAM - Johann Radon Institute for Computational and Applied Mathematics, Altenberger
Strasse 69, Linz, Austria, 4040
f
INAF - Osservatorio Astronomico di Roma, Via Frascati 33, Monte Porzio Catone, Italy,
I-00078
g
Adoni Laboratorio Nazionale di Astrofisica, Italy
ABSTRACT

The astronomical applications greatly benefit from the knowledge of the instrument PSF. We describe the PSF
Reconstruction algorithm developed for the LBT LUCI instrument assisted by the SOUL SCAO module. The reconstruction procedure considers only synchronous wavefront sensor telemetry data and a few asynchronous calibrations. We do not compute the Optical Transfer Function and corresponding filters. We compute instead a temporal series of wavefront maps and for each of these the corresponding instantaneous PSF. We tested the algorithm both in laboratory arrangement and in the nighttime for different SOUL configurations, adapting it to the guide star magnitudes and seeing conditions. We nick-named it “BRUTE”, Blind Reconstruction Using
TElemetry, also recalling the one-to-one approach, one slope-to one instantaneous PSF the algorithm applies.
Keywords: Adaptive Optics; astronomy; infrared imaging; point spread functions; SOUL; LBT; LUCI; BRUTE

1. INTRODUCTION
The thirst for enhanced sensitivity brought to extremely sophisticated systems. The current paradigm for optical instrumentation is to design larger and larger telescopes properly armed with Adaptive Optics (AO)
modules. Such effort caused an analogous complexity of the camera/spectrograph response. The ground-based
PSF, which was seeing dependent, now is time-evolving and spatially varying on minutes and arcsecond scales, respectively. The stochastic behaviour of both the optical turbulence and the telescope vibrations corresponds to an unreliable PSF definition. A long-time average may reduce the divergence with respect to predictions, reducing the problem to an accurate knowledge of the initial conditions. Accurate photometry and morphological studies require accurate PSF knowledge. The PSF information is mandatory, but the availability of point sources is not obvious, since the field of view imaged by AO instruments is typically small (tens of arcsec). We focus on the blind PSF-reconstruction1, 2 (PSF-R). It is “blind” because applies to the cases when the PSF reference is not available: when the only viable way to build a PSF is to use complementary information made available by the observatory (seeing, wind speed) or by the AO instrument (otherwise, i.e., see Massari+ (2020)3 ). These
“telemetry” data offer the opportunity to reduce the PSF object to parameters describing turbulence statistics,
Further author information: (Send correspondence to C.A.)
C.A.: E-mail: carmelo.arcidiacono@inaf.it, Telephone: +39 049 829 3414

AO loop, noise, vibrations, and the optical train. We describe a possible method, we are developing using the
SOUL4, 5 @LBT6 AO module, and the LUCI7 imager as the reference application. SOUL recently upgraded the
FLAO8 systems at the LBT replacing the wavefront sensor camera with an OCAM2k.9 The FLAO and SOUL
are SCAO adaptive optics modules. Thanks to the low readout noise camera, finer pupil resolution and increased readout frequency SOUL improves several merit figures: sky coverage, magnitude limits, and Strehl Ratio values
(optical quality) SOUL, as the First Light Adaptive Optics (FLAO)8 did, feeds the NIR camera LUCI.
The authors are collaborating to the development of MORFEO10 and MICADO11 adaptive optics module and camera for the European ELT.12 MICADO is one of the first light instruments of the ELT. It will be fed by
ELT through the MAORY adaptive optics for Multi-Conjugate Adaptive Optics (MCAO) or through an internal
SCAO13 module. The contribution of the authors to MICADO is in the development of the PSF-R software
(MICADO specific). The important connection is that both MICADO SCAO and SOUL modules use:
• a Pyramid WaveFront Sensor14 (P-WFS);
• a contact-less voice coil deformable mirror controlled in position using capacitive sensors15 for optical turbulence correction
SOUL is the optimal benchmark for the MICADO PSF-R development, see also Simioni+ (2022).16 We develop this activity within the MAST&R working group, and gathered contributions mainly from INAF-Arcetri
(SOUL data) and INAF-Padova (MICADO PSF-R team) producing a radial elongation of the PSF centered on the reference star, see Figure 1.

2. PSF-R FOR SINGLE CONJUGATE ADAPTIVE OPTICS
The test case we considered here is the SOUL adaptive optics system for the LBT telescope, feeding the LUCI17
Near Infrared camera. In particular, we focus on the on-axis (on the direction of the guide star) case. SOUL is a
Single Conjugate Adaptive Optics (SCAO) system. A SCAO modules senses and corrects the turbulence in the direction of the single guide star. and the optical performance peaks in the direction of the guide star. An AO
system squeezes the PSF by moving one or more deformable mirrors against the fast optical disturbance injected by the turbulence.
The SCAO performance ultimately depends on the initial seeing (+telescope vibrations) and on the brightness of the reference star. The wavefront correction is limited by loop frequency, actuators spacing (on the deformable mirror), Wave Front Sensor-errors (noise and misalignment) and by the control of the Non-Common
Path Aberration (NCPA, instrument versus wavefront sensor). Adaptive Optics flattens the incoming wavefront and acts moving speckles from the seeing halo to the aperture mask diffraction pattern.
In the high Strehl Ratio regime (larger than 20-30%) the PSF is quasi diffraction-limited and present typical structures. The number of modes used to control the deformable mirror fixes the highest spatial frequency corrected:
• It is ultimately limited by the number of actuators;
• This highest frequency generates a boundary with not corrected one, the so-called control radius,18 well visible in SCAO mode;
• Control Radius is at the radius of λ/2d where λ is the imaging wavelength, and d is the characteristic distance between two adjacent actuators on the deformable mirror. Smaller the distance larger the control radius, at which the halo of the uncorrected component of the seeing starts.

Figure 1. SCAO correction vs MCAO correction. In the MCAO the spatial variability of the PSF is much lower than in the SCAO case. On the left the SCAO M15 core, 23”x23” data J band PISCES+FLAO @ LBT, inset 6”x5”. On the right the MCAO, on the halo of NGC 6388, 65”x63”, Ks band, MAD @ VLT, inset 6”x5”

3. PSF-R BLIND METHOD “BRUTE - BLIND RECONSTRUCTION USING
TELEMETRY”
3.1 BRUTE PSF-R in short
To distinguish this method from other blind methods using only adaptive optics system and telescope telemetry data,19–22 we nick-named this BRUTE, Blind Reconstruction Using TElemetry, also recalling the one-to-one method, one slope-to one instantaneous PSF, we follow. We build a wavefront array for each iteration of the adaptive optics loop that was saved. We use the signal from the wavefront sensor, in our case strictly related to the XY gradient of the wavefront, to generate a wavefront map, and we add to this that part of the wavefront unseen by the wavefront sensor. BRUTE uses the adaptive optics loop information registered simultaneously with the science frame. BRUTE fits turbulence on the pseudo open loop data and noise models on the slopes statistics.
Pseudo-open loop data is the description of the original open loop (wavefront without adaptive correction) made by summing synchronous wavefront slopes contributions and deformable mirror shape. It applies an instrumental look up table from calibrated P-WFS optical gain curves23, 24 finally modulated for the best estimation of its amplitude. The algorithm applies calibration made for each of the operative modes of the SOUL instrument, which adapts the pupil resolution (CCD binning) and camera read-out frequency to the reference star brightness.
For each wavefront, BRUTE computes the instantaneous PSF corresponding to the modeled wavefront map and considers the actual pupil mask definition. In the following sections, we take as a reference the case of SOUL at
LBT. It to be stressed that in the case of SOUL the optical gain is monitored during the operation and applied to properly perform the NCPA suppression.25 In the algorithm, regardless of the origin of the residuals, all the wavefronts components are described as a combination of shapes (typically mirror modes, or Zernike modes).

3.2 What the telemetry data are
The BRUTE PSF-R algorithm requests data synchronous to the observations. This condition allows accounting for specific incident happening during open shutter time (wind gusts, seeing spikes, and una-tantum vibrations and loop instabilities).
Wavefront Sensor Slopes: the camera-synchronous wavefront sensor measurements history: it can be a large data sets since grows at loop frequency. Each iteration step is large as twice the number of active sub-apertures.
Deformable Mirror History: the camera-synchronous deformable mirror shape history: growing at loop frequency and large as the number of controlled modes. The control loop describes the deformable mirror shape as linear a combination of the ith -shapes composing the deformable mirror modal description.

Loop Gain: if applied, it is a vector of modulation amplitude, long as the number of modes in use, that defines for each mirror mode the fraction of it to be added to the deformable mirror.
Better is the coverage of the observation with the registration of these data, higher is the precision of the reconstructed PSF. To these quantities that varies during the operation, others calibration or pre-processed data are needed for the blind PSF-R.
IM: it is the deformable mirror to wavefront sensor interaction matrix, this and its the pseudo-inverse IM −1
will be used by the BRUTE algorithm
R: is the control matrix, this is known and determined using the calibrated interaction matrix (IM );
KL: is the array containing the wavefront map of each of the modes in the modal base (KL stands for
Karhounen-Loeve26 ).
NCPA is the vector of modal coefficients containing the modal description of the wavefront difference between wavefront sensor and scientific camera.

3.3 Optical Loop Gain
Given the information in the telemetry, ideally, the instantaneous wavefront seen by the wavefront sensor is:
W F SLP = slopes × IM −1 × KL
In a closed loop system, this W F SLP is a low-resolution image of the original residual wavefront incident on the wavefront sensor. It is the so-called “parallel part“ mostly responsible for the shape of the PSF within the control radius. The parallel part of the wavefront, the W F SLP , history is then computed from the slopes records.
However, several corrections need to be implemented, as for the noise component and for the optical loop gain gopt mentioned above. The proper derivation of the wavefront map from the slopes requires the correction of the non-linearity effects. These makes deviating the response amplitude of the wavefront sensor to the proportional relation with incident wavefront amplitude. The deviation depending also on the spatial scale of the measured mode.
W F SLP = slopes × IM −1 · gopt × KL
In the P-WFS case the optical loop gain, should be calibrated using dedicated strategies, see Agapito+ (2021)5
or Deo+ (2022)27 for application to optimal loop control.

3.4 Noise part into the parallel component
From the history of the slopes we can derive both noise and turbulence contributions. The noise component is part of the parallel wavefront. It is added to the estimated wavefront as part of the W FSLP . We derive the modal coefficients vector of the parallel part of the wavefront, m, from the slopes: m = slopes × IM −1 · gopt returns the elements of the linear combination of wavefront shapes composing the instantaneous W FSLP . The amplitude corresponding to the noise part in the m is removed, by scaling the m mode by mode following the diag IM IM T , and the remaining part is associated to the optical turbulence and transferred as is into the W FSLP ..

3.5 Orthogonal component and derived quantities
The BRUTE PSF-R algorithm estimates the characteristics length (r0 ) of the spatial power spectral density of the turbulence, fitting an analytical model onto the deformable mirror history and wavefront sensor data. We used here a modal representation of the power spectral density (PSD) adapting the Noll (1976)28 formula, fitting the power law on the power spectral density of the modal description of pseudo open loop WF. The point is to identify in that PSD the original atmospheric turbulence contribution, separating the effect of the noise. We assumed that tip-tilt modes are dominated by the telescope vibrations, thus excluding these from the power law fit. The fit allows extrapolating the modal amplitude of the uncontrolled components (the modes- or the frequencies higher than the one controlled by the SCAO system), finally estimating the orthogonal part of the wavefront.
I.e. in SOUL data, we had 500 KL modes, and the orthogonal components extend up to Zernike 3000.

Figure 2. Scheme dividing parallel and orthogonal components. Different colours indicate if the quantities in use are calibrated, measured or secondary products.

3.6 Building the PSF
The processes above are listed schematically in Figure 2 The integrated PSF is the sum of a number of realizations, one for each k-slope vector frame:
W F radk = 2π · W F SLP + W F HiOrder + W F N CP A
X
2
P SF =
| pupil · eiW F radk |
k

4. BRUTE PSF-R METHOD FOR SOUL-LUCI
4.1 Optical loop gain
SOUL continuously tracks the optical gain (sensitivity), by measuring the feedback to a sinusoidal modulation injected on a single mode (#30), Figure 3. The slopes history saved are already corrected for this value. However, this estimation can be refined in the post-processing, and the algorithm re-computes it using again the peak of temporal power spectrum corresponding to the continuous modulation of KL mode #30, and now comparing the ratio between the amplitudes of modes #30 as modulated by the deformable mirror over the modes #30 as derived by the wavefront sensor with respect to similar ratio on adjacent modes. This correction is applied mode by mode scaling an empirical relation calibrated using “daytime” data, in which the optical distortion evolution is known, see also Deo+ 2018.24

4.2 Pupil Masks-Cold stop and Parallactic angle
LUCI foresees three different level of zoom. The one designed for Adaptive Optics application is the N30
camera.29 The N30 camera is a reversed telescope design that includes its “secondary mirror” inserted in the optical axis and mounted on three arms: the diffraction pattern due to these arms follows the apparent sky rotation, However, the effective pupil illumination of LUCI depends whether working on daytime (laboratory)
or nighttime (on-sky), since in the laboratory configuration “daytime”30 LUCI does not see the M2 spider arms
(and the primary mirror). Actual size of the telescope is 8.06 meter because of the extended cold stop in LUCI, and central obstruction is 0.314, because of the obstruction imposed by the secondary mirror of the reflective design of the N30 camera, see Figure 4.

4.3 Pupil Masks, cold stop and parallactic angle
Daytime pattern does not rotate, while the M2-M3 spider arms rotate with the parallactic angle as seen from detector. To build the nighttime on-sky pupil we need to rotate the spider pattern mask for the corresponding parallactic angle and add this to the N30 pupil mask.

Figure 3. Temporal power spectral density of the slopes response (black, solid) and of the deformable mirror command for modulated mode # 30,

Figure 4. From left to right. On-sky pupil definition,the LUCI N30 pupil and an image taken in pupil camera mode (not the N30) during the alignment of LUCI. We used this kind of image to calibrate pupil rotation angle with the LUCI
rotator angle movement.

Figure 5. On the left an image taken with LUCI using BIN 2 mode (200modes). The filter is the broadband J. On the central panel the PSF reconstructed using BRUTE. On the right the direction of dispersion is made evident by the arrow.

4.4 Filter width and Atmospheric Dispersion
The PSF resulting from previous steps are monochromatic. The smoothing effect due to the wavelength dispersion along the width of filter is then applied, considering the slight changes of the PSF shape across the filter.
The LUCI camera does not mount any atmospheric dispersion compensator. As a result, the diffractionlimited images are elongated along the telescope elevation axis direction, Figure 5. This elongation depends mainly on the wavelength, the filter width and on the elevation angle, and also on other parameters that were not tracked during the observation we used (humidity, pressure vertical profile) or have a negligible effect as the color of the star (we observed in the Near Infrared: the spectrum of non-peculiar stars is quite flat across the wavelengths in the broad-band filters). So, the last step for the BRUTE PSF-R is the application of a correction for wavelengths dependant effects on the integrated PSF computed above. We use the analytic formula that gives the chromatic dispersion with respect to the elevation, applying small shift of the monochromatic PSFs along the dispersion direction.

5. BRUTE PSF-R APPLICATION TO SOUL-LUCI DATA
We applied the method on 25 LUCI SOUL frames acquired during commissioning time in a single night of observation. Each LUCI frame was simultaneously registered with SOUL wavefront sensor telemetry. For that commissioning run (11-2019) only the LBT left side was available. We had two different targets, with similar magnitudes. We observed 25 sets of 0.3sec × 20 images in the narrow filter FeII (1.645 micron). See Figure 6. The first star presents a quite stable loop, modulated by the seeing variation, while the second had loop instabilities that make an extreme variation of the performance frame by frame. Despite the two having different behavior the BRUTE algorithm can nicely match the average Strehl Ratio and Encircled Energy profile with errors in the stable case is limited to a few points of percentage in line with the dispersion of the original Strehl Ratio measurements itself.

6. CONCLUSIONS
We reported about the successful application of the BRUTE blind PSF reconstruction method to on sky
LUCI@LBT data. The SOUL adaptive optics system offers the telemetry data necessary to extract the wavefront components: open loop turbulence, noise and pyramid optical gain. We successfully matched the high Strehl
Ratio PSF and demonstrated the method valid on high Strehl Ratio conditions. We tested a-posteriori correction for the pyramid wavefront sensor optical gain. The results obtained on the LUCI@SOUL data are very promising for MICADO PSF-R development, and the BRUTE PSF-R system can be offered right now as support for the observations (on our server). We use the history of the deformable mirror shape for the estimation of the turbulence PSD.

Figure 6. 25 LUCI observations, each one simultaneously registered with SOUL wavefront sensor telemetry on the LBT
left side, two different targets, with similar magnitudes. See insets on the left of the graph.

Figure 7. On the left the average image of the 25 frames of the observed SOUL@LUCI PSF and on the right the mean of the BRUTE PSF-R PSF reconstructions. The mean of the 25 frames gives the possibility to compare also the low signal noise ratio halo of the PSF, otherwise marginally visible in the original LUCI frame.

However, this opportunity is less valid in the case that a few modes are controlled, and/or correction is poor
(the optical surface of the deformable mirror is not half of the negative copy of the reference star WF).
Next step is to prove the method on lower Strehl Ratio observations characterized by a larger noise component.

ACKNOWLEDGMENTS
The work has been supported by INAF through the Math, ASTronomy and Research (MAST&R), a working group for mathematical methods for high-resolution imaging.

REFERENCES
[1] Simioni, M., Arcidiacono, C., Grazian, A., Clenet, Y., Davies, R., Gullieuszik, M., Kleijn, G. V., Pedichini,
F., Wagner, R., Ramlau, R., Zeilinger, W. W., Vidal, F., Vulcani, B., Ragazzoni, R., Sevin, A., Salasnich,
B., Baruffolo, A., Busoni, L., Esposito, S., Gendron, E., Piazzesi, R., Portaluri, E., Zanella, A., Helin, T.,
Kuncarayakti, H., Mattila, S., Falomo, R., and Pinna, E., “MICADO PSF-reconstruction work package description,” in [Adaptive Optics Systems VII ], 11448, 724–733, SPIE (Dec. 2020).
[2] Beltramo-Martin, O., Ragland, S., FÃ©tick, R., Correia, C., Dupuy, T., Fiorentino, G., Fusco, T., Jolissaint, L., Kamann, S., Marasco, A., Massari, D., Neichel, B., Schreiber, L., and Wizinowich, P., “Review of
PSF reconstruction methods and application to post-processing,” in [Adaptive Optics Systems VII], 11448,
114480A, International Society for Optics and Photonics (Dec. 2020).
[3] Massari, D., Marasco, A., Beltramo-Martin, O., Milli, J., Fiorentino, G., Tolstoy, E., and Kerber, F.,
“Successful application of PSF-R techniques to the case of the globular cluster NGC 6121 (M 4),” Astronomy
& Astrophysics 634, L5 (Feb. 2020). Publisher: EDP Sciences.
[4] Pinna, E., Esposito, S., Hinz, P., Agapito, G., Bonaglia, M., Puglisi, A., Xompero, M., Riccardi, A.,
Briguglio, R., Arcidiacono, C., Carbonaro, L., Fini, L., Montoya, M., and Durney, O., “SOUL: the Single conjugated adaptive Optics Upgrade for LBT,” in [Adaptive Optics Systems V ], Marchetti, E., Close, L. M., and Véran, J.-P., eds., Society of Photo-Optical Instrumentation Engineers (SPIE) Conference Series 9909,
99093V (July 2016).
[5] Agapito, G., Rossi, F., Plantet, C., Puglisi, A., and Pinna, E., “Advances in control of a pyramid single conjugate adaptive optics system,” Monthly Notices of the Royal Astronomical Society 508, 1745–1755 (Dec.
2021). ADS Bibcode: 2021MNRAS.508.1745A.
[6] Hill, J. M., “The Large Binocular Telescope,” Applied Optics 49, 115–122 (Apr. 2010).
[7] Wagner, R. M., Edwards, M. L., Kuhn, O., Thompson, D., and Veillet, C., “An overview and the current status of instrumentation at the Large Binocular Telescope Observatory,” in [Ground-based and Airborne
Instrumentation for Astronomy V ], 9147, 30–46, SPIE (July 2014).
[8] Esposito, S., Riccardi, R., Fini, L., Puglisi, A. T., Pinna, E., Xompero, M., Briguglio, R., Quirós-Pacheco,
F., Stefanini, P., Guerra, J. C., Busoni, L., Tozzi, A., Pieralli, F., Agapito, G., Brusa-Zappellini, G., Demers,
R., Brynnel, J., Arcidiacono, C., and Salinari, P., “First light AO (FLAO) system for LBT: final integration, acceptance test in Europe, and preliminary on-sky commissioning results,” in [Adaptive Optics Systems II ],
Proc. SPIE 7736 (2010).
[9] Pinna, E., Rossi, F., Puglisi, A., Agapito, G., Bonaglia, M., Plantet, C., Mazzoni, T., Briguglio, R.,
Carbonaro, L., Xompero, M., Grani, P., Riccardi, A., Esposito, S., Hinz, P., Vaz, A., Ertel, S., Montoya,
O. M., Durney, O., Christou, J., Miller, D. L., Taylor, G., Cavallaro, A., and Lefebvre, M., “Bringing SOUL
on sky,” (Jan. 2021). , ADS Bibcode: 2021arXiv210107091P Type: article.
[10] Ciliegi, P., Diolaiti, E., Abicca, R., Agapito, G., Aliverti, M., Arcidiacono, C., Auricchio, N., Balestra, A.,
Baruffolo, A., Bellazzini, M., Bonaglia, M., Bregoli, G., Brissaud, O., Busoni, L., Carlotti, A., Cascone,
E., Correia, J.-J., Cortecchia, F., Cosentino, G., D’Orazi, V., Dall’Ora, M., Caprio, V. D., Rosa, A. D.,
Delboulbé, A., Antonio, I. D., Rico, G. D., Dolci, M., Esposito, S., Fantinel, D., Feautrier, P., Fiorentino, G.,
Foppiani, I., Giro, E., Gluck, L., Grani, P., Greggio, D., Hénault, F., Jocou, L., Penna, P. L., Lafrasse, S.,
Lauria, M., Coarer, E. L., Louarn, M. L., Lombini, M., Magnard, Y., Magrin, D., Maiorano, E., Mannucci,
F., Marchetti, E., Maurel, D., Michaud, L., Moraux, E., Morgante, G., Moulin, T., Oberti, S., Pariani, G.,
Patti, M., Plantet, C., Podio, L., Puglisi, A., Rabou, P., Ragazzoni, R., Redaelli, E., Riva, M., Rochat, S.,

Roussel, F., Roux, A., Salasnich, B., Saracco, P., Schreiber, L., Spavone, M., Stadler, E., Sztefek, M.-H.,
Terenzi, L., Valentini, A., Ventura, N., VÃérinaud, C., and Zaggia, S., “MAORY for ELT: preliminary design overview,” in [Adaptive Optics Systems VI], 10703, 1070311, International Society for Optics and
Photonics (July 2018).
[11] Davies, R., Ageorges, N., Barl, L., Bedin, L. R., Bender, R., Bernardi, P., Chapron, F., Clenet, Y., Deep,
A., Deul, E., Drost, M., Eisenhauer, F., Falomo, R., Fiorentino, G., Förster Schreiber, N. M., Gendron,
E., Genzel, R., Gratadour, D., Greggio, L., Grupp, F., Held, E., Herbst, T., Hess, H.-J., Hubert, Z.,
Jahnke, K., Kuijken, K., Lutz, D., Magrin, D., Muschielok, B., Navarro, R., Noyola, E., Paumard, T.,
Piotto, G., Ragazzoni, R., Renzini, A., Rousset, G., Rix, H.-W., Saglia, R., Tacconi, L., Thiel, M., Tolstoy,
E., Trippe, S., Tromp, N., Valentijn, E. A., Verdoes Kleijn, G., and Wegner, M., “MICADO: the E-ELT
adaptive optics imaging camera,” 7735, 77352A (July 2010). Conference Name: Ground-based and Airborne
Instrumentation for Astronomy III Place: eprint: arXiv:1005.5009.
[12] Gilmozzi, R. and Spyromilio, J., “The 42m European ELT: status,” in [Society of Photo-Optical Instrumentation Engineers (SPIE) Conference Series ], Society of Photo-Optical Instrumentation Engineers (SPIE)
Conference Series 7012 (2008).
[13] Clénet, Y., Buey, T., Gendron, E., Hubert, Z., Vidal, F., Cohen, M., Chapron, F., Sevin, A., Fédou, P.,
Barbary, G., Borgo, B., Huet, J.-M., Blin, A., Dupuis, O., Gaudemard, J., Ben Nejma, S., Gratadour, D.,
Deo, V., Ferreira, F., Thijs, S., Lapeyrére, V., Raffard, J., Chemla, F., Le Ruyet, B., Bertrou-Cantou, A.,
Rozel, M., Younés, Y., Rousset, G., Zins, G., Diolaiti, E., Ciliegi, P., Garrel, V., Rabien, S., Schubert, J.,
Hartl, M., Hörmann, V., and Davies, R., “MICADO-MAORY SCAO Preliminary design, development plan
& calibration strategies,” (Nov. 2019). Conference Name: Proceedings of the AO4ELT6 conference Pages:
E3 ADS Bibcode: 2019elt6.confE...3C.
[14] Ragazzoni, R., “Pupil plane wavefront sensing with an oscillating prism,” Journal of Modern Optics 43,
289–293 (1996).
[15] Salinari, P., Del Vecchio, C., and Biliotti, V., “A Study of an Adaptive Secondary Mirror,” 48, 247, ESO
(Jan. 1994). Conference Name: European Southern Observatory Conference and Workshop Proceedings
ADS Bibcode: 1994ESOC...48..247S.
[16] Simioni, M., Arcidiacono, C., Wagner, R., Grazian, A., Gullieuszik, M., Portaluri, E., Vulcani, B., Zanella,
A., Agapito, G., Davies, R., Helin, T., Pedichini, F., Piazzesi, R., Pinna, E., Ramlau, R., Rossi, F., and Salo,
A., “Point spread function reconstruction for SOUL+LUCI LBT data,” arXiv e-prints , arXiv:2209.01563
(Sept. 2022).
[17] Seifert, W., Appenzeller, I., Baumeister, H., Bizenberger, P., Bomans, D., Dettmar, R.-J., Grimm, B.,
Herbst, T., Hofmann, R., Juette, M., Laun, W., Lehmitz, M., Lemke, R., Lenzen, R., Mandel, H., Polsterer,
K., Rohloff, R.-R., Schuetze, A., Seltmann, A., Thatte, N. A., Weiser, P., and Xu, W., “LUCIFER: a multimode NIR instrument for the LBT,” in [Instrument Design and Performance for Optical/Infrared
Ground-based Telescopes ], 4841, 962–973, International Society for Optics and Photonics (Mar. 2003).
[18] Perrin, M. D., Sivaramakrishnan, A., Makidon, R. B., Oppenheimer, B. R., and Graham, J. R., “The
Structure of High Strehl Ratio Point-Spread Functions,” The Astrophysical Journal 596, 702 (Oct. 2003).
Publisher: IOP Publishing.
[19] Wagner, R., Hofer, C., and Ramlau, R., “Point spread function reconstruction for single-conjugate adaptive optics on extremely large telescopes,” Journal of Astronomical Telescopes, Instruments, and Systems 4,
049003 (Sept. 2018). Publisher: SPIE.
[20] Vèran, J.-P., Rigaut, F., Maı̂tre, H., and Rouan, D., “Estimation of the adaptive optics long-exposure point-spread function using control loop data,” Journal of the Optical Society of America A 14, 3057 (Nov.
1997).
[21] Fusco, T., Bacon, R., Kamann, S., Conseil, S., Neichel, B., Correia, C., Beltramo-Martin, O., Vernet, J.,
Kolb, J., and Madec, P. Y., “Reconstruction of the ground-layer adaptive-optics point spread function for
MUSE wide field mode observations,” A&A 635, A208 (Mar. 2020).
[22] Ragland, S., Jolissaint, L., Wizinowich, P., Dam, M. A. v., Mugnier, L., Bouxin, A., Chock, J., Kwok, S.,
Mader, J., Witzel, G., Do, T., Fitzgerald, M., Ghez, A., Lu, J., Martinez, G., Morris, M. R., and Sitarski,
B., “Point spread function determination for Keck adaptive optics,” in [Adaptive Optics Systems V ], 9909,
573–590, SPIE (July 2016).

[23] Korkiakoski, V., Vérinaud, C., and Louarn, M. L., “Improving the performance of a pyramid wavefront sensor with modal sensitivity compensation,” Applied Optics 47, 79–87 (Jan. 2008). Publisher: Optica
Publishing Group.
[24] Deo, V., Gendron, E., Rousset, G., Vidal, F., and Buey, T., “A modal approach to optical gain compensation for the pyramid wavefront sensor,” in [Adaptive Optics Systems VI ], 10703, 653–670, SPIE (July 2018).
[25] Esposito, S., Puglisi, A., Pinna, E., Agapito, G., Quirós-Pacheco, F., Véran, J. P., and Herriot, G., “On-sky correction of non-common path aberration with the pyramid wavefront sensor,” A&A 636, A88 (Apr. 2020).
[26] Karhunen, K., “Über lineare Methoden in der Wahrscheinlichkeitsrechnung,” Ann. Acad. Sci. Fennicae.
Ser. A. I. Math.-Phys. 37, 1–79 (1947).
[27] Deo, V., Rozel, M., Bertrou-Cantou, A., Moura Ferreira, F., Vidal, F., Gratadour, D., Sevin, A., Clénet,
Y., Rousset, G., and Gendron, E., “CLOSE: a self-regulating, best-performance tracker for modal integrator based AO loops,” in [Adaptive Optics for Extremely Large Telescopes conference, 6th edition ], (Nov. 2019).
[28] Noll, R. J., “Zernike polynomials and atmospheric turbulence*,” JOSA 66, 207–211 (Mar. 1976). Publisher:
Optical Society of America.
[29] Heidt, J., Pramskiy, A., Thompson, D., Seifert, W., Gredel, R., Miller, D., Taylor, G., Esposito, S., Puglisi,
A., Pinna, E., and Quirrenbach, A., “Commissioning of the adaptive optics supported LUCI instruments at the Large Binocular Telescope: results,” in [Ground-based and Airborne Instrumentation for Astronomy
VII], 10702, 63–77, SPIE (July 2018).
[30] Riccardi, A., Xompero, M., Zanotti, D., Busoni, L., Del Vecchio, C., Salinari, P., Ranfagni, P., Brusa
Zappellini, G., Biasi, R., Andrighettoni, M., Gallieni, D., Anaclerio, E., Martin, H. M., and Miller, S. M.,
“The adaptive secondary mirror for the Large Binocular Telescope: results of acceptance laboratory test,”
in [Adaptive Optics Systems ], Proc. SPIE 7015 (2008).

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

---
source: "https://arxiv.org/abs/2403.08455v2"
title: "Interaction of Autonomous and Manually Controlled Vehicles Multiscenario Vehicle Interaction Dataset"
author: "Novel Certad, Enrico del Re, Helena Korndörfer, Gregory Schröder, Walter Morales-Alvarez, Sebastian Tschernuth, Delgermaa Gankhuyag, Luigi del Re, Cristina Olaverri-Monreal"
year: "2024"
publication: "arXiv preprint / cs.RO"
download: "https://arxiv.org/pdf/2403.08455v2"
pdf: "https://arxiv.org/pdf/2403.08455v2"
captured_at: "2026-06-27T06:15:30Z"
updated_at: "2026-06-27T06:15:30Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács Theory of the Novel"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# Interaction of Autonomous and Manually Controlled Vehicles Multiscenario Vehicle Interaction Dataset

- 著者: Novel Certad, Enrico del Re, Helena Korndörfer, Gregory Schröder, Walter Morales-Alvarez, Sebastian Tschernuth, Delgermaa Gankhuyag, Luigi del Re, Cristina Olaverri-Monreal
- 年: 2024
- 掲載情報: arXiv preprint / cs.RO
- 情報源: [arxiv](https://arxiv.org/abs/2403.08455v2)
- ダウンロード: https://arxiv.org/pdf/2403.08455v2
- PDF: https://arxiv.org/pdf/2403.08455v2

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

The acquisition and analysis of high-quality sensor data constitute an essential requirement in shaping the development of fully autonomous driving systems. This process is indispensable for enhancing road safety and ensuring the effectiveness of the technological advancements in the automotive industry. This study introduces the Interaction of Autonomous and Manually-Controlled Vehicles (IAMCV) dataset, a novel and extensive dataset focused on inter-vehicle interactions. The dataset, enriched with a sophisticated array of sensors such as Light Detection and Ranging, cameras, Inertial Measurement Unit/Global Positioning System, and vehicle bus data acquisition, provides a comprehensive representation of real-world driving scenarios that include roundabouts, intersections, country roads, and highways, recorded across diverse locations in Germany. Furthermore, the study shows the versatility of the IAMCV dataset through several proof-of-concept use cases. Firstly, an unsupervised trajectory clustering algorithm illustrates the dataset's capability in categorizing vehicle movements without the need for labeled training data. Secondly, we compare an online camera calibration method with the Robot Operating System-based standard, using images captured in the dataset. Finally, a preliminary test employing the YOLOv8 object-detection model is conducted, augmented by reflections on the transferability of object detection across various LIDAR resolutions. These use cases underscore the practical utility of the collected dataset, emphasizing its potential to advance research and innovation in the area of intelligent vehicles.

## PDF Text

1

arXiv:2403.08455v2 [cs.RO] 19 Mar 2025

Interaction of Autonomous and Manually Controlled
Vehicles Multiscenario Vehicle Interaction Dataset
Novel Certad∗ , Graduate Student Member, IEEE, Enrico del Re∗ , Graduate Student Member, IEEE, Helena
Korndoerfer∗ , Gregory Schroeder∗ , Member, IEEE, Walter Morales-Alvarez∗ , Member, IEEE, Sebastian
Tschernuth∗ , Student Member, IEEE, Delgermaa Gankhuyag∗ , Luigi del Re† , Senior Member, IEEE, Cristina
Olaverri-Monreal∗ , Senior Member, IEEE
∗ Department Intelligent Transport Systems
Johannes Kepler University Linz, 4040 Linz, Austria
† Institute for Design and Control of Mechatronical Systems
Johannes Kepler University Linz, 4040 Linz, Austria

Abstract—The acquisition and analysis of high-quality sensor data plays a pivotal role in the development of advanced autonomous driving systems and in enhancing transportation safety.
In this study, we presented a novel and extensive dataset centered on inter-vehicle interactions. The dataset was recorded across diverse locations in Germany, including roundabouts, intersections, country roads, and highways. Equipped with a sophisticated array of sensors, including LIDAR sensors, cameras, IMU/GPS, and acquisition of the vehicle bus data, our dataset offers a comprehensive representation of real-world driving scenarios.
Additionally, two proof-of-concept use cases were implemented to demonstrate the versatility of the IAMCV dataset. First, an unsupervised trajectory clustering algorithm was used to highlight the capability of the dataset in categorizing vehicle movements without labeled training data. Finally, an online camera calibration method and the ROS-based standard were compared using the images captured in the dataset.
Index Terms—dataset, floating car data, segmented trajectories, point clouds.

I. I NTRODUCTION

T

HE rapid evolution of autonomous driving technology necessitates the availability of robust datasets that accurately reflect the complexities of real-world driving environments. With continuous advancements in sensor technologies, data acquisition systems have become more sophisticated, enabling the collection of detailed information encompassing the spatial, temporal, and contextual aspects of driving scenarios.
In this context, we present a comprehensive dataset recorded using a state-of-the-art vehicle equipped with a multitude of sensors, such as LIDARs, cameras, GPS, and various other sensors [1], [2].
The IAMCV Dataset was acquired as part of the FWF
Austrian Science Fund-funded “Interaction of Autonomous and Manually-Controlled Vehicles” project. It is primarily centered on inter-vehicle interactions and captures a wide range of road scenes in different locations across Germany,

including roundabouts, intersections, and highways. These locations were carefully selected to encompass various traffic scenarios, representative of both urban and rural environments. By simultaneously capturing data from multiple sensor modalities, our dataset provides an in-depth understanding of the road network from a driver-centric perspective, enabling researchers and developers to analyze, model, and evaluate autonomous driving algorithms and systems under diverse conditions.
The primary objective of this study was to introduce this rich dataset and outline its potential applications in advancing the field of autonomous driving. Furthermore, we present a detailed description of the sensor setup, data acquisition methodology, and preprocessing techniques employed to ensure the quality and reliability of the recorded data. The dataset presented here serves as a valuable resource for researchers, engineers, and practitioners working in the fields of autonomous driving, computer vision, and robotics.
The remainder of this paper is organized as follows: Section II provides an overview of datasets pertinent to the advancement of autonomous driving systems. Section III
describes the vehicle used for the data acquisition, sensor arrangement, and calibration procedures. An overview of the
IAMCV dataset is presented in Section IV. Subsequently, the insights of the dataset are analysed in Section V including two proof-of-concept applications to showcase the dataset’s usability. Final conclusions are provided in Section VI.
II. R ELATED W ORK
This section provides a summary of the datasets relevant to the progress of autonomous driving systems, focusing on multimodal sensor data. The datasets are presented in a chronological sequence.
According to the authors in [3], datasets can be categorized into three primary types: datasets from fixed stations, drones, and test vehicles. Datasets from fixed stations are derived from sensors (such as cameras, LIDARs, and radars) affixed to traffic lights, poles, or buildings at specific locations, serving dual roles as surveillance systems and sources of researchoriented traffic data. One of the notable datasets in this

This work was supported by the Austrian Science Fund (FWF), project number P 34485-N. It was additionally supported by the Austrian Ministry for
Climate Action, Environment, Energy, Mobility, Innovation, and Technology
(BMK) Endowed Professorship for Sustainable Transport Logistics 4.0., IAV
France S.A.S.U., IAV GmbH, Austrian Post AG and the UAS Technikum
Wien.
©2024 IEEE. Personal use of this material is permitted. Permission from IEEE must be obtained for all other uses, in any current or future media, including reprinting/republishing this material for advertising or promotional purposes, creating new collective works, for resale or redistribution to servers or lists, or reuse of any copyrighted component of this work in other works. DOI: 10.1109/MITS.2024.3378114

2

category is The Next Generation Simulation (NGSIM) Dataset
[4].
Similar to fixed stations, drone datasets offer a stable and location-specific perspective. Notably, datasets of this kind have seen an increasing prevalence in recent times [3], [5]–[9].
Nevertheless, for the development, training, and analysis of perception and control algorithms for ADS-operated vehicles, data from test vehicles remains the primary source of essential information.
In 2013, the KITTI dataset [10] was a cornerstone in autonomous vehicle research because of its comprehensive and real-world nature. This dataset encompasses a wide array of sensor data, including high-resolution images, LIDAR
point clouds, and GPS information, collected from a vehicle equipped with a variety of sensors, as it navigates urban environments.
The FORD Multi-AV dataset [11] was collected by a fleet of Ford autonomous vehicles (AVs) on different days from
2017 to 2018. The vehicles traversed an average route of
66 km in Michigan (USA), including freeways, airports, city centers, a university campus, and suburban neighborhoods.
The dataset covered seasonal variations in weather, lighting conditions, road occupancy, and traffic experienced in dynamic urban environments over different months. The vehicles were geared with an Applanix POS-LV GNSS/INS system, four
HDL-32E Velodyne 3D-lidar scanners, and six Point Grey 1.3
MP cameras arranged to grant a 360° coverage.
The PREVENTION dataset [12] was introduced in 2019 and offers extensive and precise annotations of vehicle trajectories, categories, lanes, and various events, including cut-in, cut-out, left/right lane changes, and hazardous maneuvers. The dataset was collected by deploying six distinct sensors (LIDAR, radar, and cameras) on an instrumented vehicle operating under naturalistic driving conditions. The PREVENTION dataset encompasses 356 min of data, equivalent to 540 km of travel distance, featuring over 4 million detections and more than
3,000 individual trajectories.
The EU Longterm dataset [13] was gathered in 2020 using the Robocar in the urban and suburban areas of Montbéliard,
France. The vehicle speed was maintained at 50 km/h throughout the data collection process. The recorded driving route spans approximately 5.0 kilometers, recorded in 16-minutelong segments. The data collected at 10Hz came from two stereo cameras (pointing back and front) two industrial cameras (facing the sides of the vehicle), two Velodyne HDL-32E
LIDARS, one 2D LIDAR, one radar, an IMU, and a GPS
receiver, all of which were stored in rosbag files
Further diversification was presented in 2021 with the
Waymo Open Dataset [14]. This dataset contains more than
100,000 scenes, each spanning 20 s and captured at a rate of
10 Hz, resulting in more than 570 h of unique data covering a distance of 1750 km along roadways across six cities within the United States. The dataset provides LIDAR scans, camera images, labels, and 3D bounding boxes.
Other widely used datasets include the Oxford RobotCar dataset [15], [16] introduced in 2016 and updated in 2020, nuScenes [17] introduced in 2020, and ONCE [18] introduced

in 2021. All were recorded with a sensor array, including cameras and LIDARs.
Although existing datasets have proven invaluable, the need for diversification in data sources, sensor configurations, geographical locations, and driving scenarios remains crucial. This diversity is essential for the rigorous training and validation of both new and existing perception and control algorithms. It allows algorithms to generalize better, adapt to novel and complex environments, and enhance their robustness in real-world situations. Moreover, the evolving landscape of autonomous driving, including the introduction of new sensor technologies and the emergence of different driving cultures worldwide, necessitates the continual generation of fresh data. In essence, the production of new datasets serves as the backbone of innovation in autonomous vehicle technology, enabling safer, more efficient, and globally adaptable self-driving systems.
III. DATA C OLLECTION
This section describes the sensors used in the data acquisition process, the vehicle setup, and the calibration procedures employed to record the dataset.
A. Sensor Overview
This section provides a comprehensive description of the sensors used in the data collection. The dataset was acquired using the JKU-ITS vehicle [1]. The vehicle was equipped with different sensors to perceive the surrounding environment and locate the vehicle within it.

(a)

(b)

(c)

Fig. 1: (a) depicts the location of the three Light Detection
And Ranging (LIDAR) sensors and the vertical Field of View
(FoV) associated with each sensor. (b) Shows the horizontal
FoV of the LIDAR sensors. Lastly, (c) depicts the horizontal
FoV covered by the three cameras [1].
The vehicle was equipped with three LIDARs. An OS2 (64
layers and 22.5◦ of vertical FoV) was located in the center of the roof. This sensor has an extended range of view (240m), covering objects far from the vehicle (Fig. 1a). Two OS0 (128
layers, range from 0.3m to 50m, and 90.0◦ of Vertical FoV)

3

detected the objects in the vicinity of the vehicle. One was located in the front-left part of the roof and the other was in the rear-right part. As shown in Fig. 1a, both OS0 sensors were inclined 22.5◦ towards the plane to cover the sides of the vehicle. This configuration with three LIDAR sensors allowed the vehicle to cover the entire area surrounding the car without noticeable dead zones (Fig. 1b).

connect the acquisition software in ROS [19] with the software interface of the Black Panda (based on the open source system
Openpilot). Throughout the entirety of this paper, the variables of the vehicle accessed through this method are referred to as “car state”. A list of the most commonly used car state variables is presented in Table II. A full list of more than 40
variables can be found in the dataset files [22].

TABLE I: Sensors’ specifications.

TABLE II: car state variables

Lidars lidar c lidar l lidar r
Cameras camera c camera l camera r

model
OS2
OS0
OS0
model acA2040-35gc acA2040-35gc acA2040-35gc

range
240m
50m
50mVFoV
22.5◦
90.0◦
90.0◦
VAoV
59.4◦
59.4◦
59.4◦

HFoV
360◦
360◦
360◦
HAoV
79.0◦
79.0◦
79.0◦

layers
64
128
128
resolution
3.5MP
3.5MP
3.5MP

Name vEgo aEgo wheelSpeeds gearShifter steeringAngle gas gasPressed

1) Cameras: Three acA2040-35gc Basler GigE cameras
(3.2MP, Angle of View: 79.0◦ horizontal, 59.4◦ vertical) were located on the roof, covering a horizontal field of view of approximately 200◦ , as shown in Fig. 1c. The cameras were configured to acquire 10 frames per second according to an external 10Hz signal (trigger) generated by the synchronization box.
2) GNSS/INS: The positioning system receiver (3DMGX545) was located beneath the central LIDAR sensor. It offered access to raw data from the Global Positioning System (GPS)
at a refresh rate of 4Hz, data from the Inertial Measurement
Unit (IMU) at a sample rate of 100Hz, and filtered odometry data obtained through a Kalman filter method at a refresh rate of 100Hz. Furthermore, the GNSS/INS established a direct connection to the main computer through a USB cable.
B. Sensor Calibration
Intrinsic calibration of the cameras was achieved using the standard Robot Operating System (ROS) package camera calibration [19]. This camera calibration model uses the
OpenCV library [20] and is based on the pinhole camera model. The model estimates intrinsic camera parameters, such as focal length, principal point coordinates, and lens distortion coefficients. The calibration procedure involves capturing multiple images of a known calibration pattern (chessboard)
from different viewpoints, and using the corresponding image and world coordinates to calculate the camera parameters.
The extrinsic calibration of the cameras and LIDAR sensors was performed following the automatic procedure described in [21] with the available ROS package. The three cameras and the two side-LIDARs were calibrated using the centerLIDAR location as the reference frame.
C. Vehicle Data Bus
The OBDII port is a standardized diagnostic connector that allows access to real-time data from the vehicle engine and other systems. The Black Panda is a universal car interface that provides access to the communication buses of the car through this port and a USB-C interface at the other end to connect a computer. A custom bridge based on Python was developed to

brakePressed steeringTorque

Description vehicle linear speed in m/s vehicle linear acceleration in m/s2
wheel linear speed in m/s (one measure per wheel)
gear level position (park, neutral, drive, etc...)
angle of the steering wheel gas pedal level, where 0 is released and 1 is fully depressed value specifying whether the gas pedal is pressed
() or released value specifying whether the brake pedal is pressed or released torque applied on the steering wheel

D. Synchronization
Synchronizing data from different sensors is an important step in creating high-quality datasets for research on autonomous vehicles. Thus, the accuracy and robustness are improved. In the IAMCV dataset, the synchronization was achieved using two methods:
• Synchronization by hardware: A hardware box generated six trigger signals for the LIDARs and cameras [1].
Through this process, it is guaranteed that the three
LIDAR beams point to the front of the vehicle simultaneously when the cameras are triggered. However, the timestamps are assigned at the moment the frames are received in the computer running ROS and then edited in an offline process to guarantee that for every point cloud, the timestamp corresponds to the moment when the beam was pointing to the front of the vehicle. This is at midscan for the LIDARs in the center and the left and at the beginning of the scan for the one on the right.
• Synchronization by software: In the case of the car state variables and the GPS, it was not possible to connect a trigger signal; thus, the data were acquired and the timestamps were assigned at the moment the frames were received in the computer running ROS. Some latency between 0ms and 30ms was measured for this procedure.
E. Vehicle setup
The dataset was acquired using the JKU-ITS vehicle [1] as shown in Fig. 2. The vehicle was equipped with the sensors described in Section III-A. The cameras and LIDAR sensors were connected to the main computer through a network switch. The data acquisition software was based on ROS [19].
A custom package for launching the ROS-driver for each sensor and the associated transformations was developed. The raw data are primarily stored in rosbag files and later converted to more standard options (CSV files, PNG, etc.).

4

approach by introducing driver-centric insights, that enable the analysis of high-level interaction patterns.
The collected data were anticipated to include various driving patterns and trajectories contingent upon specific scenarios and traffic conditions. In the locations selected for intersections and roundabouts analysis, the vehicle traversed the same reference intersection or roundabout several times; thus, the followed path resembles loops around the reference point.
These driving patterns are particularly useful for loop-closure purposes when testing mapping or localization algorithms.
B. Data Types

Fig. 2: The JKU-ITS vehicle was used for the collection of the data.

IV. IAMCV DATASET
A. Methodology
The dataset was recorded from the 3rd to the 10th of June of 2022. The recordings were performed during the daytime, between 08:00 h to 20:00 h under different natural light levels
(sunny, cloudy, or overcast). Because the dataset aims to record different interactions between the recording vehicle and other vehicles, trucks, motorcycles, and vulnerable road users, the chosen locations were: intersections, roundabouts, roads, and highways.
Three recording locations for intersections were selected in Aachen, Germany. In suburban areas around the city, two locations were chosen for roundabouts and two for country roads. Finally, highway locations were selected along the A3
Federal Motorway between Passau and Cologne. The path in

Fig. 3: Route followed from Linz (Austria) to Aachen (Germany).
which the data were collected, aligns with the route taken from Linz (Austria) to attend the IEEE Intelligent Vehicles
Symposium conference in Aachen (Germany). The entire route is illustrated in Fig. 3. It was also taken into consideration that these sites had already been investigated in other datasets using drones [5]–[7]. THE IAMCV dataset supplements this

The dataset comprises different types of data:
• Point clouds: The point clouds from the three LIDARs were stored using Point Cloud Data (PCD) file format.
All original fields that came from the LIDAR sensors were also stored in the file:
– x,y,z: coordinates of the target in the reference sensor frame.
– Range: distance from the sensor to the target in mm.
– Near-infrared: is the ambient level of infrared sensed by the receiver when the transmitter does not point to that area.
– Intensity: signal level in the receiver.
– Reflectivity: reflectivity level of the target.
– Ring: LIDAR’s layer.
• Images: Similarly, the images from the three cameras were stored using Portable Network Graphic (PNG) file format. The images were stored without rectification and intrinsic parameters are provided. The images in the dataset were anonymized before publishing by blurring faces and license plates. This was done automatically using scripts based on [23] and [24].
Regarding the naming, files from cameras and LIDARs were named using the format “nn sensor tttttttttt.tttt.ext” as follows:
• nn: stands for the number of recordings. Each recording has a unique two-digit identifier.
• sensor: indicates whether the sensor is a camera or a
LIDAR and its position. e.g., “camera c” stands for the camera located in the center of the vehicle, and “lidar l”
stands for the LIDAR located in the left side of the vehicle.
• tttttttttt.tttt: it corresponds to the timestand of the frame.
• ext: file extension .pcd or .png
Fig. 4 depicts the internal organization of the files that comprised the IAMCV dataset.
V. DATASET A NALYSIS
The IAMCV dataset comprises more than 50 segments of 15 min of driving for a total of approximately 15 h of recording. The overall information is presented in Table III.
Hereunder, some examples of the content of the dataset are discussed. For instance, Fig. 5 presents a set of three images simultaneously captured by the camera sensors. To facilitate

5

<nn>
camera camera_c
<nn>_camera_c_<tttttttttt.tttt>.png
...
camera_l camera_r lidar lidar_c
<nn>_lidar_c_<tttttttttt.tttt>.pcd
...
lidar_l lidar_r info nn_nav_odom.csv nn_carstate.csv tf.csv camera_c_intrinsics.yaml camera_l_intrinsics.yaml camera_r_intrinsics.yaml
...
labels

(a)

Fig. 4: Directory tree depicting how the folders and files comprising the dataset are organized.
TABLE III: IAMCV Dataset: Overview of the locations and recording times location

type

segments
#
27
19

duration
[min]
435
295

A3 (Germany)
Aachen city (Frankenberger,
Bendplatz, Charlottenburgerallee)
B57 (Aachen)
Aachen suburbs (Neuweiler,
Kackertstrasse)
Totals

Highway
Intersections
Urban
Country road
Roundabout

2
8

21
138

-

56

888

/

subsequent stitching processes, we ensured a deliberate 20%
overlap between images.
Fig. 6 depicts the recorded GPS trajectory followed by the vehicle during two recordings. The top image represents a recording that focused on an intersection. As described in Section IV-A, the vehicle traversed the same intersection several times. The image at the bottom was recorded on the
A3 highway around Cologne.
One of the key advantages of the IAMCV dataset is the large

(a)

(b)

(c)

Fig. 5: Raw images acquired simultaneously:(a) left camera,
(b) center camera, and (c) right camera.

(b)

Fig. 6: Openstreet map with the GPS trajectory followed in two recordings: (a) Recording around Charlottenburgeralle in
Aachen. (b) Recording in the A3 highway around Cologne.

number of heterogeneous scenarios recorded. The availability of these scenarios is a crucial step for testing Automated
Driving Systems (ADS)-Operated vehicles and improving the reliability and safety of the algorithms. Six of the most relevant scenarios are depicted in Fig. 7; nevertheless, the complete list of available scenarios is considerably larger.
A comparison between the IAMCV dataset and the datasets described in Section II is presented in Table IV. Notably,
IAMCV distinguishes itself by incorporating data not only from external GNSS and IMU sources, but also from the vehicle bus itself. This comprehensive approach provides a holistic and fine-grained understanding of vehicle dynamics, sensor interactions, and control inputs, which is essential for developing robust autonomous systems. Furthermore, the IAMCV dataset scenarios and location diversity ensure that algorithms and models trained on the IAMCV dataset are better equipped to handle real-world driving scenarios, making it an invaluable resource for developing and testing autonomous systems. However, it is important to acknowledge that the IAMCV dataset may benefit from further expansion to encompass a broader spectrum of weather conditions including rain, snow, and fog.
A. Applications
To demonstrate the IAMCV dataset, serves as a versatile tool with practical applications, two compelling use cases are presented next. In the first instance, we successfully

6

TABLE IV: Comparison between IAMCV Dataset and similar available datasets.

7TB

frames
#
530k

fs1
cameras2
[Hz] pixels
10
3x3.5MP

LIDAR3
layers
1x64 + 2x128

sunny & cloudy various weather

<1TB
23TB

15k
-

10
16

1x64
1x4 + 2x1

Waymo Open [14]

various weather

-

230k

nuScenes [17]
A2D2 [25]

various weather various weather

354GB
2.3TB

various weather

Dataset

weather

size

IAMCV (ours)

sunny & cloudy

KITTI [10]
Oxford RobotCar [15], [16]

10

4x0.7MP
1x1.2MP +
3x1MP
5x

5x

400k
40k

12
10

6x1.4MP
6x2.3MP

1x32
5x16

>10TB

-

6

4x32

sunny

600GB

-

10

6x1.3MP +
1.5MP
2x2MP

1x32

Eu Long Term [13]

various weather

870GB

-

8

eurocity [26]
ONCE [18]

various weather various weather

-

47k
1M

0.25
2

2
+
2(stereo)
1x2MP
7

2x32 + 1x4 +
1x1
No
1

Ford Multi-AV
[11]
Prevention [12]

Seasonal

vehicle data
GPS, IMU, state
GPS, IMU
GPS, IMU

location
Car

velocity, angular velocity
GPS, IMU
GPS, IMU, Car state
GPS, IMU
GPS, IMU, state
GPS, IMU
No
No

Car

Urban, highways, country roads
Urban
Urban
Urban
Urban
Urban, highways, country roads
Urban, highways, country roads
Highways
Urban
Urban
Urban, highways, country roads

“-” : data not provided.
1) fs: sample frequency. In the case of several values, the smaller one is listed.
2) AxB where A is the number of cameras and B is the resolution.
3) AxB where A is the number of sensors and B is the number of layers in each sensor.

implemented unsupervised trajectory clustering, showcasing the dataset’s utility in understanding and categorizing vehicle movements without the need for labeled training data. In a second application, IAMCV proved instrumental in camera calibration. Leveraging the diverse scenarios and sensor data provided by the dataset, we achieved precise calibration, crucial for accurate perception and object localization.
1) Unsupervised trajectory segmentation and clustering: In this application, we tried to segment the recorded trajectories and then cluster the resulting segments into different categories related to the dynamic of the vehicle: braking, turning, accelerating/decelerating. This was accomplished by first segmenting the data using a Bayesian modelbased sequence segmentation (BMOSS) algorithm and then clustering the segments by employing an extended latent
Dirichlet allocation model (LDA). By investigating line- and trajectory plots, the respective driving action was then derived for each cluster. The method was adapted from [27].
As parameters linear acceleration (m/s2 ) along the x-axis, i.e. in the direction of motion, and angular velocity (rad/s)
in the z-axis were selected. The data were directly extracted from the IAMCV dataset files without any further preprocessing. With BMOSS, the sequence segmentation problem was formulated as finding the segment boundaries that maximize the overall marginal likelihood of the data. To derive the likelihood of the possible segmentations, the Bayesian piecewise regression model was employed as the underlying framework.
It followed the assumption that the sequence of input-output data is piecewise stationary and within each segment the data follow a Gaussian distribution. Traversing the model’s marginal likelihood under all segmentation choices was done recursively with a forward-backward algorithm. The authors in [27] extend traditional LDA for continuous driving data in a way that instead of inputting the driving samples directly into the model, they are replaced with symbols, where each

symbol is a mixture component. Topics, which are typically distributions over discrete input data, are thus represented as mixtures of these symbols. Similar to traditional LDA, the extended model then returns the topic proportions for each segment. The topic with the highest proportion is then assigned to each segment and the driving action related to each topic can be derived. In Fig. 8 and Fig. 9 the clustering is visualized together with the derived driving actions for one of the recordings performed in roundabouts. It is clear how three distinctive clusters appeared in the trajectory: Strong counterclockwise rotation (purple), braking with clockwise rotation
(dark cyan), and accelerating/decelerating on a straight line
(yellow).
2) Segment detection for intrinsic camera calibration:
The camera intrinsic matrix is a critical element in computer vision, encapsulating internal parameters like focal length for mapping 3D points to a 2D image plane. Calibration, often involving geometric patterns, establishes these parameters.
However, factors such as wear or environmental changes can impact them, necessitating periodic recalibration for precisiondemanding applications like robotics or computer vision.
To address this need, we introduced a robust edge segment detection algorithm. Specifically designed for camera calibration, the algorithm enables online calibration without pre-established geometric patterns.
The methodology involves multiple steps for 3D line detection and segmentation from distorted images. First, acceptable orientation ranges are set to prevent the connection of multiple correct edges. Inspired by the Canny algorithm, edge detection follows, including Gaussian filtering, gradient calculation, and thresholding for identifying strong and weak edge points.
Weak points are included in the final edge map only if connected to strong points, ensuring edge continuity. Edge pixels are then grouped into continuous segments using a raster scan approach. Shape-based filtering is applied, requiring a minimum extent and aspect ratio for reliable circle fitting.

7

Linear Acceleration

(a)

(b)

2

0.2

1

0.0

Angular Velocity

0.8

Cluster
Strong counter-clockwise rotation 0.6
Braking with clockwise rotation
0.4
Accelerating

3

0

0.2

1

0.4
0.6

0.0

50.0

100.0

Time [sec]

150.0

Fig. 8: Line plot of the clustered data with the derived driving action for each cluster as a label.
Cluster
Strong counter-clockwise rotation
Braking with clockwise rotation
Accelerating

50.8930
50.8920

(d)
Latitude

(c)

50.8910
50.8900
50.8890
50.8880
6.1750

(e)

(f)

6.1725

6.1700

6.1675

6.1650

Longitude

6.1625

6.1600

6.1575

6.1550

Fig. 9: Trajectory plot of the clustered data with the derived driving action for each cluster as a label.

VI. C ONCLUSIONS AND F UTURE WORK

(g)

(h)

Fig. 7: Examples of heterogeneous driving scenes presented within the IAMCV dataset: (a) car crash, (b) country roads,
(c) tunnel, (d) traffic jam, (e) overtaking in highway, (f)
highway, (g) pedestrian crossing with vulnerable road users,
(h) Roundabout.

The algorithm addresses 3D straight-line discontinuities by merging associated edge segments, involving circle fitting and residual calculation. Additional shape-based filtering, requiring a minimum length for acceptance, is applied. Finally, outlier filtering eliminates false detections resembling curved
3D lines. This methodology aims to accurately detect and segment 3D lines from distorted images, addressing challenges like noise, discontinuities, and false detections as shown in
Fig. 10.
Ultimately, the distortion parameters are extracted from the detected edge segments. In Table V the distortion parameters extracted using the IAMCV calibration procedure (ROSbased) are compared with this online approach. Only radial distortion parameters were calculated because their effect is often considered dominant.

The IAMCV dataset represents a significant contribution to the field of autonomous driving and sensor-based data collection. Through meticulous data collection and rigorous quality control processes, we have curated a comprehensive and diverse dataset that encapsulates various traffic scenarios across different locations in Germany, including roundabouts, intersections, and highways. This dataset not only serves as a valuable resource for validating object detection, tracking, 3D
mapping, and localization algorithms to advance autonomous vehicle technology but also opens doors to innovative research opportunities in sensor fusion, machine learning, and realworld transportation applications. As we look ahead, we envision this dataset becoming a cornerstone for researchers and engineers working towards safer, more efficient, and more intelligent transportation systems. Future iterations of

(a)

(b)

Fig. 10: Examples of edge detection on two images extracted from the IAMCV dataset.

8

TABLE V: Comparison of the two main radial distortion parameters estimated using the calibration package in ROS
with a chessboard pattern against the online approach using the robust segment detection camera left left center center right right

method
ROS
online
ROS
online
ROS
online

k1
-0.285356
-0.284997
-0.283517
-0.275294
-0.284487
-0.274391

k2
0.082809
0.079993
0.080529
0.073669
0.081967
0.075005

the dataset will take into account weather conditions to further enhance its applicability.
Two distinct use cases were presented to underscore the dataset’s adaptability and effectiveness in addressing multiple challenges within the realm of autonomous systems.
As we continue to explore IAMCV’s potential, it becomes increasingly evident that its comprehensive and varied data make it an invaluable resource for a spectrum of applications, further solidifying its significance in advancing the field of autonomous driving research.
We plan to include annotations in the form of bounding boxes and high-level labels for selected segments, highlighting relevant interactions in the next release. These real-world data and annotations will play a vital role in advancing the development of autonomous driving systems, making them more robust and dependable. Ultimately, this progress will contribute to the enhancement of safety and efficiency in the field of autonomous vehicles.
ACKNOWLEDGMENT
This work was supported by the Austrian Science Fund
(FWF), project number P 34485-N. It was additionally supported by the Austrian Ministry for Climate Action, Environment, Energy, Mobility, Innovation, and Technology (BMK)
Endowed Professorship for Sustainable Transport Logistics
4.0., IAV France S.A.S.U., IAV GmbH, Austrian Post AG and the UAS Technikum Wien.
R EFERENCES
[1] N. Certad, W. Morales-Alvarez, G. Novotny, and C. Olaverri-Monreal,
“JKU-ITS Automobile for Research on Autonomous Vehicles,” in
Computer Aided Systems Theory – EUROCAST 2022, R. MorenoDı́az, F. Pichler, and A. Quesada-Arencibia, Eds.
Cham: Springer
International Publishing, 2023, pp. 3–10.
[2] A. V. Barrio, W. M. Alvarez, C. Olaverri-Monreal, and J. E. N.
Hernández, “Development and validation of an open architecture for autonomous vehicle control,” in 2023 IEEE Intelligent Vehicles Symposium (IV), 2023, pp. 1–6.
[3] P. Tkachenko, N. Certad, G. Singer, C. Olaverri-Monreal, and L. Del Re,
“The JKU DORA Traffic Dataset,” IEEE Access, vol. 10, pp. 92 673–
92 680, 2022.
[4] V. G. Kovvali, V. Alexiadis, and L. Zhang PE, “Video-based vehicle trajectory data collection,” Tech. Rep., 2007.
[5] J. Bock, R. Krajewski, T. Moers, S. Runde, L. Vater, and L. Eckstein,
“The inD Dataset: A Drone Dataset of Naturalistic Road User Trajectories at German Intersections,” in 2020 IEEE Intelligent Vehicles
Symposium (IV), 2020, pp. 1929–1934.
[6] R. Krajewski, T. Moers, J. Bock, L. Vater, and L. Eckstein, “The rounD
Dataset: A Drone Dataset of Road User Trajectories at Roundabouts in
Germany,” in 2020 IEEE 23rd International Conference on Intelligent
Transportation Systems (ITSC), 2020, pp. 1–6.

[7] R. Krajewski, J. Bock, L. Kloeker, and L. Eckstein, “The highD Dataset:
A Drone Dataset of Naturalistic Vehicle Trajectories on German Highways for Validation of Highly Automated Driving Systems,” in 2018 21st
International Conference on Intelligent Transportation Systems (ITSC),
2018, pp. 2118–2125.
[8] T. Moers, L. Vater, R. Krajewski, J. Bock, A. Zlocki, and L. Eckstein,
“The exiD Dataset: A Real-World Trajectory Dataset of Highly Interactive Highway Scenarios in Germany,” in 2022 IEEE Intelligent Vehicles
Symposium (IV), 2022, pp. 958–964.
[9] W. Zhan, L. Sun, D. Wang, H. Shi, A. Clausse, M. Naumann,
J. Kümmerle, H. Königshof, C. Stiller, A. de La Fortelle, and
M. Tomizuka, “INTERACTION Dataset: An INTERnational, Adversarial and Cooperative moTION Dataset in Interactive Driving Scenarios with Semantic Maps,” arXiv:1910.03088 [cs, eess], Sep. 2019.
[10] A. Geiger, P. Lenz, C. Stiller, and R. Urtasun, “Vision meets Robotics:
The KITTI Dataset,” International Journal of Robotics Research (IJRR),
2013.
[11] S. Agarwal, A. Vora, G. Pandey, W. Williams, H. Kourous, and
J. McBride, “Ford Multi-AV Seasonal Dataset,” The International
Journal of Robotics Research, vol. 39, no. 12, pp. 1367–1376, 2020.
[Online]. Available: https://doi.org/10.1177/0278364920961451
[12] R. Izquierdo, A. Quintanar, I. Parra, D. Fernández-Llorca, and M. A.
Sotelo, “The PREVENTION dataset: a novel benchmark for PREdiction of VEhicles iNTentIONs,” in 2019 IEEE Intelligent Transportation
Systems Conference (ITSC), Oct 2019, pp. 3114–3121.
[13] Z. Yan, L. Sun, T. Krajnik, and Y. Ruichek, “EU Long-term Dataset with
Multiple Sensors for Autonomous Driving,” in Proceedings of the 2020
IEEE/RSJ International Conference on Intelligent Robots and Systems
(IROS), 2020.
[14] S. Ettinger, S. Cheng, B. Caine, C. Liu, H. Zhao, S. Pradhan,
Y. Chai, B. Sapp, C. Qi, Y. Zhou, Z. Yang, A. Chouard, P. Sun,
J. Ngiam, V. Vasudevan, A. McCauley, J. Shlens, and D. Anguelov,
“Large Scale Interactive Motion Forecasting for Autonomous Driving :
The Waymo Open Motion Dataset,” in 2021 IEEE/CVF International
Conference on Computer Vision (ICCV). Los Alamitos, CA, USA:
IEEE Computer Society, oct 2021, pp. 9690–9699. [Online]. Available: https://doi.ieeecomputersociety.org/10.1109/ICCV48922.2021.00957
[15] W. Maddern, G. Pascoe, C. Linegar, and P. Newman, “1 Year, 1000km:
The Oxford RobotCar Dataset,” The International Journal of Robotics
Research (IJRR), vol. 36, no. 1, pp. 3–15, 2017. [Online]. Available: http://dx.doi.org/10.1177/0278364916679498
[16] D. Barnes, M. Gadd, P. Murcutt, P. Newman, and I. Posner,
“The Oxford Radar RobotCar Dataset: A Radar Extension to the
Oxford RobotCar Dataset,” in Proceedings of the IEEE International
Conference on Robotics and Automation (ICRA), Paris, 2020. [Online].
Available: https://arxiv.org/abs/1909.01300
[17] H. Caesar, V. Bankiti, A. H. Lang, S. Vora, V. E. Liong, Q. Xu, A. Krishnan, Y. Pan, G. Baldan, and O. Beijbom, “nuScenes: A multimodal dataset for autonomous driving,” in CVPR, 2020.
[18] J. Mao, M. Niu, C. Jiang, H. Liang, J. Chen, X. Liang, Y. Li, C. Ye,
W. Zhang, Z. Li, J. Yu, H. Xu, and C. Xu, “One Million Scenes for
Autonomous Driving: ONCE Dataset,” 2021.
[19] M. Quigley, B. Gerkey, K. Conley, J. Faust, T. Foote, J. Leibs, E. Berger,
R. Wheeler, and A. Ng, “ROS: an open-source Robot Operating System,”
in Proc. of the IEEE Intl. Conf. on Robotics and Automation (ICRA)
Workshop on Open Source Robotics, Kobe, Japan, May 2009.
[20] G. Bradski, “The OpenCV Library,” Dr. Dobb’s Journal of Software
Tools, 2000.
[21] J. Beltrán, C. Guindel, A. de la Escalera, and F. Garcı́a, “Automatic
Extrinsic Calibration Method for LiDAR and Camera Sensor Setups,”
IEEE Transactions on Intelligent Transportation Systems, 2022.
[22] N. Certad, E. Del Re, A. Aghanouri, D. Gankhuyag, W. MoralesAlvarez, S. Tschernuth, L. Del Re, and C. Olaverri-Monreal.
(2023) IAMCV Interaction of Autonomous and Manually Controlled
Vehicles. IEEE Dataport. [Online]. Available: https://dx.doi.org/10.
21227/d1g3-c160
[23] Y. Xu, W. Yan, H. Sun, G. Yang, and J. Luo, “CenterFace: Joint Face
Detection and Alignment Using Face as Point,” in arXiv:1911.03599,
2019.
[24] S. M. Silva and C. R. Jung, “License plate detection and recognition in unconstrained scenarios,” in 2018 European Conference on Computer
Vision (ECCV), Sep 2018, pp. 580–596.
[25] J. Geyer, Y. Kassahun, M. Mahmudi, X. Ricou, R. Durgesh, A. S.
Chung, L. Hauswald, V. H. Pham, M. Mühlegg, S. Dorn, T. Fernandez,
M. Jänicke, S. Mirashi, C. Savani, M. Sturm, O. Vorobiov, M. Oelker,
S. Garreis, and P. Schuberth, “A2D2: Audi Autonomous Driving
Dataset,” 2020. [Online]. Available: https://www.a2d2.audi

9

[26] M. Braun, S. Krebs, F. Flohr, and D. M. Gavrila, “EuroCity Persons:
A Novel Benchmark for Person Detection in Traffic Scenes,” IEEE
Transactions on Pattern Analysis and Machine Intelligence, vol. 41, no. 8, pp. 1844–1861, 2019.
[27] A. Bender, G. Agamennoni, J. R. Ward, S. Worrall, and E. M. Nebot, “An unsupervised approach for inferring driver behavior from naturalistic driving data,” IEEE Transactions on Intelligent Transportation Systems, vol. 16, no. 6, pp. 3325–3336, 2015.

Novel Certad holds a Master of Science in Electronic Engineering from Simon Bolivar University
(USB) in Caracas, Venezuela, awarded in June 2013.
As a research assistant, he is currently pursuing a Ph.D. at the Department of Intelligent Transport
Systems, at the Johannes Kepler University Linz,
Austria. His research areas encompass autonomous vehicles, interactions with road users, signal processing, and probabilistic robotics.

Walter Morales-Alvarez graduated in 2018 from the Simon Bolivar University in Caracas as an Electronics Engineer. He is a PhD student at Johannes
Kepler University, where he is working full-time as a university assistant/lecturer in the Department of
Intelligent Transport Systems. His research focuses on the field of Intelligent Transport Systems. He has published in conferences such as the IEEE International Conference on Intelligent Transport Systems and the IEEE Intelligent Vehicle Symposium. His research interests include autonomous driving humanmachine interaction and deep learning.

Sebastian Tschernuth received a bachelor’s degree in Biological Chemistry at the JKU Linz and the
USB Budweis, and currently studies Artificial Intelligence Masters at the JKU Linz. As a project researcher at the Department of Intelligent Transport
Systems, he works on Data Analysis and Machine
Learning.

Enrico Del Re graduated with a Master of Science in Physics 2019 from ETH Zürich in Switzerland and is currently pursuing a Ph.D. at the Department of Intelligent Transport Systems at the Johannes
Kepler University Linz, Austria. His research there is focused on the safety of autonomous vehicles, especially in mixed-traffic situations.

Helena Korndoerfer received a Bachelor of Arts in Economics & Politics from the University of
Munster, Germany in 2020. She is currently pursuing a Master of Science degree in Economic & Business
Analytics at the Johannes Kepler University Linz,
Austria and is writing her Master’s thesis at the
Department of Intelligent Transport Systems.

Gregory Schroeder received his Master of Science degree in Computational Engineering Science from the Technische Universität Berlin, Germany, in 2018.
Since then, he has been engaged in the automotive industry, focusing on pre-development projects in assisted and autonomous driving technologies. He specializes in Computer Vision, visual Simultaneous
Localization and Mapping (SLAM), and camera models & calibration. His work emphasizes realworld applications, particularly in developing sophisticated and reliable visual systems. His efforts are instrumental in bridging the gap between theoretical computational concepts and their practical implementation in automotive settings.

Delgermaa Gankhuyag received a bachelor’s degree in statistics and economics from the National
University of Mongolia, Ulaanbaatar, Mongolia, in
June 2012. She is currently pursuing a master’s degree in statistics at Johannes Kepler University,
Linz, Austria. She is a project researcher at the
Department of Intelligent Transport Systems, Johannes Kepler University, Linz, Austria. Her current research interest includes autonomous vehicles and road users interaction, autonomous vehicle’s trajectory prediction, and naturalistic driving behavior of drivers at unsignalized intersections.

Luigi del Re (SM’11) received the Doctoral degree in electrical engineering from the Department of
Mechanical Engineering, ETH, Zurich, Switzerland.
He was a Research Associate with the Institute of
Automatic Control, ETH. He left the ETH in 1994 to work for the Swatch group on different automotive issues and he has been a Full Professor with Johannes Kepler University Linz, Austria, since 1998.
His main applications are in the field of engine and vehicle technology, hydraulics, process control, and biomedicine. His current research interests include complex control problems which cannot be solved using purely analytical approaches, and, in this context, both nonlinear control design and nonlinear identification, in particular using approximation classes.

10

Cristina Olaverri-Monreal is professor and head of the Department of Intelligent Transport Systems at the Johannes Kepler University Linz, in Austria.
Prior to this position, she led diverse teams in the industry and in the academia in the US and in distinct countries in Europe.
She is also the president of the IEEE Intelligent Transportation Systems Society (IEEE ITSS), founder and chair of the Austrian IEEE ITSS chapter, and chair of the Technical Activities Committee
(TAC) on Human Factors in ITS.
She received her PhD from the Ludwig-Maximilians University (LMU) in
Munich in cooperation with BMW. Her research aims at studying solutions for efficient and effective transportation focusing on minimizing the barrier between users and road systems. To this end, she relies on the automation, wireless communication, and sensing technologies that pertain to the field of
Intelligent Transportation Systems (ITS).
Prof. Olaverri is a member of the EU-wide platform for coordinating open road tests (Cooperative, Connected and Automated Mobility (CCAM)) as well as a representative for the European technology platform ”Alliance for Logistics Innovation through Collaboration in Europe” (ALICE) for the ”Workgroup
Road Safety” (WG4: EU-CCAM-WG-ROAD-SAFETY@ec.europa.eu). She is additionally a senior/associate editor and editorial board member of several journals in the field, including the IEEE ITS Transactions and IEEE ITS
Magazine.
Furthermore, she is an expert for the European Commission on ”Automated
Road Transport” and consultant and project evaluator in the field of ICT
and ”Connected, Cooperative Autonomous Mobility Systems” for various EU
and national agencies as well as organizations in Germany, Sweden, France,
Ireland, etc. In 2017, she was the general chair of the ”IEEE International
Conference on Vehicles Electronics and Safety” (ICVES’2017). She was awarded the ”IEEE Educational Activities Board Meritorious Achievement
Award in Continuing Education” for her dedicated contribution to continuing education in the field of ITS.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

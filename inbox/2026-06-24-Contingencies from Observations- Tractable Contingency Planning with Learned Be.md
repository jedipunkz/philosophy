---
source: "https://arxiv.org/abs/2104.10558v1"
title: "Contingencies from Observations: Tractable Contingency Planning with Learned Behavior Models"
author: "Nicholas Rhinehart, Jeff He, Charles Packer, Matthew A. Wright, Rowan McAllister, Joseph E. Gonzalez, Sergey Levine"
year: "2021"
publication: "arXiv preprint / cs.RO"
download: "https://arxiv.org/pdf/2104.10558v1"
pdf: "https://arxiv.org/pdf/2104.10558v1"
captured_at: "2026-06-24T11:11:55Z"
updated_at: "2026-06-24T11:11:55Z"
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

# Contingencies from Observations: Tractable Contingency Planning with Learned Behavior Models

- 著者: Nicholas Rhinehart, Jeff He, Charles Packer, Matthew A. Wright, Rowan McAllister, Joseph E. Gonzalez, Sergey Levine
- 年: 2021
- 掲載情報: arXiv preprint / cs.RO
- 情報源: [arxiv](https://arxiv.org/abs/2104.10558v1)
- ダウンロード: https://arxiv.org/pdf/2104.10558v1
- PDF: https://arxiv.org/pdf/2104.10558v1

## Obsidian Links

- 研究動向: [[研究動向/リチャード・ローティ-現代研究動向|リチャード・ローティ-現代研究動向]]
- キーワード: [[リチャード・ローティ]]
- 関連分野: [[現代思想]]
- 関連分野: [[プラグマティズム]]
- 関連分野: [[ネオプラグマティズム]]
- 関連分野: [[反表象主義]]
- 関連タグ: #現代思想 #プラグマティズム #ネオプラグマティズム #反表象主義

## Abstract

Humans have a remarkable ability to make decisions by accurately reasoning about future events, including the future behaviors and states of mind of other agents. Consider driving a car through a busy intersection: it is necessary to reason about the physics of the vehicle, the intentions of other drivers, and their beliefs about your own intentions. If you signal a turn, another driver might yield to you, or if you enter the passing lane, another driver might decelerate to give you room to merge in front. Competent drivers must plan how they can safely react to a variety of potential future behaviors of other agents before they make their next move. This requires contingency planning: explicitly planning a set of conditional actions that depend on the stochastic outcome of future events. In this work, we develop a general-purpose contingency planner that is learned end-to-end using high-dimensional scene observations and low-dimensional behavioral observations. We use a conditional autoregressive flow model to create a compact contingency planning space, and show how this model can tractably learn contingencies from behavioral observations. We developed a closed-loop control benchmark of realistic multi-agent scenarios in a driving simulator (CARLA), on which we compare our method to various noncontingent methods that reason about multi-agent future behavior, including several state-of-the-art deep learning-based planning approaches. We illustrate that these noncontingent planning methods fundamentally fail on this benchmark, and find that our deep contingency planning method achieves significantly superior performance. Code to run our benchmark and reproduce our results is available at https://sites.google.com/view/contingency-planning

## PDF Text

Contingencies from Observations: Tractable Contingency
Planning with Learned Behavior Models
Nicholas Rhinehart∗ , Jeff He∗ , Charles Packer, Matthew A. Wright
Rowan McAllister, Joseph E. Gonzalez, Sergey Levine

arXiv:2104.10558v1 [cs.RO] 21 Apr 2021

UC Berkeley
Abstract—Humans have a remarkable ability to make decisions by accurately reasoning about future events, including the future behaviors and states of mind of other agents. Consider driving a car through a busy intersection: it is necessary to reason about the physics of the vehicle, the intentions of other drivers, and their beliefs about your own intentions. If you signal a turn, another driver might yield to you, or if you enter the passing lane, another driver might decelerate to give you room to merge in front. Competent drivers must plan how they can safely react to a variety of potential future behaviors of other agents before they make their next move. This requires contingency planning: explicitly planning a set of conditional actions that depend on the stochastic outcome of future events. Contingency planning outputs a policy that is a function of future timesteps and observations, whereas standard model predictive controlbased planning outputs a sequence of future actions, which is equivalent to a policy that is only a function of future timesteps. In this work, we develop a general-purpose contingency planner that is learned end-to-end using high-dimensional scene observations and low-dimensional behavioral observations. We use a conditional autoregressive flow model to create a compact contingency planning space, and show how this model can tractably learn contingencies from behavioral observations. We developed a closed-loop control benchmark of realistic multiagent scenarios in a driving simulator (CARLA), on which we compare our method to various noncontingent methods that reason about multi-agent future behavior, including several stateof-the-art deep learning-based planning approaches. We illustrate that these noncontingent planning methods fundamentally fail on this benchmark, and find that our deep contingency planning method achieves significantly superior performance.

I. I NTRODUCTION
The ability of humans to anticipate, probe, and plan to react to what other actors could do or want is central to many social tasks that humans find simple but AI systems find difficult [1]. Effectively driving a car, for example, generally requires (1) predictive dynamics models of how cars move,
(2) predictive models of driver reactions, and (3) the ability to plan around and resolve uncertainty about other drivers’
intentions. Autonomous systems that operate in multi-agent environments typically satisfy 1–3 with separate modules that decouple prediction and planning, yet in practice the modules are highly dependent on accurate perception, an unsolved task [2–4]. Learning-based end-to-end systems that navigate complex multi-agent environments from raw sensory input have been the focus of recent work [5–7], but do not satisfy (3), because they do not model the behaviors of the other agents.
A potential approach to building a system is to first build models to forecast other agents’ behavior, and then construct
“open-loop” action plans using these models [8, 9]. However,

Fig. 1: A robot’s unprotected left, in which it does not have the right of way in the intersection and must either wait for the oncoming car to either pass or yield before turning. An unprotected left is a realistic multi-agent scenario that requires contingent behavior in order to achieve optimal outcomes, as the best a priori robot trajectory suboptimally waits outside the intersection (left branch), i.e. the best sequence of noncontingent future decisions is (Wait Outside, Turn
After Human), because the robot cannot be guaranteed to be able to traverse a fixed spatiotemporal path due to uncertainty about the human’s intention. However, a contingent planner can achieve a better result: the optimal plan is to enter the intersection in order to observe if the other driver will yield to the robot, and either turn or wait contingent upon the other driver’s decision. Top: The decision tree of robot and human behavior. Our approach, Contingencies from
Observations, uses behavioral observations sampled from decision trees as training data to learn a behavioral model used in a contingent planning objective. Bottom: Conceptual diagrams of the outcomes at each leaf. Videos, code, and trained models are available at our website: https://sites.google.com/view/contingency-planning/home.

this fails to account for the co-dependency between robot actions and environment. Specifically, it ignores how other agents would react to robot actions and how the robot’s future actions should be different depending on what the other agents do. This can lead to underconfident behavior, known as the “frozen robot problem” [10]. Thus, in many real-world situations, decoupling the tasks of forecasting human behavior and planning robot behavior is a poor modeling assumption. A
resolution to this problem can be achieved through contingency planning [11–14], which produces a plan that is adaptive to future behaviors of other agents or the environment. This is equivalent to planning a closed-loop policy.
In many partially-observed settings, optimal policies must gather information about the state of the environment (e.g. the intentions of other agents) and then act accordingly. Consider a common scenario encountered when driving on public roads: an unprotected left-hand turn, in which the robot’s goal is to turn left in the presence of another vehicle that may or may not yield, as seen in Fig. 1. The other driver’s intention to yield

Xr1

Xr2

o

Xr1

Xr2

Xh2

πr

Xr1

Xr2

Xh1

Xh2

o

o
Xh1

πr

Xh1

Xh2

Xr1

Xr2

Xh1

Xh2

o

(a) Robot leader planning. The (b) Human leader planning. The (c) Co-leader planning. This ap- (d) Joint planning. The planner planner does not model the hu- planner does not model the robot’s proach plans a policy modelled to erroneously assumes that it can man’s influence on the robot, re- influence on the human, which both influence and be influenced control all agents, potentially resulting in a deterministic noncon- means that the robot does not by the stochastic behavior of the sulting in “overconfident” behavior tingent plan. Also known as MPC, model how its future actions can human. In contrast to our method, than can lead to unrecoverable this plans action trajectories inde- affect decisions of the other agents prior contingency planning meth- errors. We demonstrate this phependent of future behaviors of the [11, 12]. This can result in “under- ods do not use a learned behavioral nomenon in our experiments.
other agents [15–17]
confident” plans.
model [13, 14].
Fig. 2: Various models of multi-agent interaction, based on different assumptions, as probabilistic graphical models. Circles denote random variables we can predict, square nodes denote robot decisions we can plan, shading indicates known values, and thick arrows represent
“carry-forward dependencies” for visual simplicity (any two nodes connected by a chain of thick nodes has an implicit directed edge).
Learned must be revealed before we (the robot) commit to turning left;
Planning Method
Contextual
Contingency
Behavior Model we must actively seek this information to resolve the uncertainty
Hardy and Campbell [11]
3
7
Passive in the other driver’s intention. Using a turn-signal indicator
Bandyopadhyay et al. [20]
7
7
Active
Xu et al. [21]
7
7
None can provoke the other agent to reveal their intention, as can
Zhan et al. [12]
7
7
Passive edging the robot into the intersection. Not only must we plan
Sadigh et al. [15]
7
7
None
Galceran et al. [13]
3
7
Active to be contingent upon the human, we must model the human
Schmerling et al. [16]
7
3
None
Rhinehart et al. [5]
3
3
None as contingent upon us. Let us term this “active contingency” to
Zhou et al. [22]
7
7
None differentiate it from “passively contingent” models that neglect
Fisac et al. [14]
7
7
Active
Zeng et al. [6]
3
3
None the robot’s influence on the other agents. In this case, we must
Tang and Salakhutdinov [17]
3
3
None plan to enter the intersection to indicate our desire to turn left
Cui et al. [23]
3
3
Passive
Bajcsy et al. [24]
7
3
Passive ahead of the other driver and have a contingency planned in
Contingencies from Observations
3
3
Active case the driver does not yield. Because the other driver may not stop, we cannot produce a single state-space plan that is Table I: Comparative summary of recent explicit planning methods guaranteed to quickly cross the intersection without potentially for autonomous navigation tasks (ordered by publication date).
causing a crash, unless we wait outside the intersection.
We demonstrate that autoregressive normalizing flows can
Recently, fully learned planning approaches have shown explicitly represent a compact and rich space of contingent promising results on autonomous navigation benchmarks plans – forking paths of future decisions necessary to operate and settings [5–7]. These approaches resemble model-based successfully in environments with stochastic outcomes. We use reinforcement learning (MBRL) because they model some this model class to design a deep contingency planner that aspect of environment dynamics in order to generate plans scales to complex tasks with high-dimensional observations under reward functions. However, these methods will fail such as autonomous driving. We designed a contingency on tasks that require explicit contingency planning because planning benchmark of common driving tasks in CARLA on they do not represent the future behavior of other agents, and which we compare our method to ablations and other methods. therefore cannot be explicitly contingent; we demonstrate [5]
We show that a single, general deep contingency planning failing in our experiments. In the terminology of Fig. 2, “robot model is quite effective across several realistic scenarios, and leader” methods, commonly referred to as MPC-shooting based that noncontingent planning methods fundamentally fail at methods, plan action trajectories, which means that the actions these common tasks.
will be fixed for all possible future behaviors of the other agents [15–17].
II. R ELATED W ORK
Modular approaches such as those used in the DARPA
Planning algorithms are central to robot navigation [18], and
Grand
Challenge [2, 3] and modern industry systems [4] have planning under uncertainty is especially critical in uncontrolled the potential to be contingent, but are highly dependent on environments like public roads. Various choices exist for imperfect perception pipelines. We are not aware of a published designing autonomous vehicle planning algorithms that consider demonstration of contingency planning in a pipeline approach other vehicles (see Schwarting et al. [19] for a thorough survey).
on a
realistic autonomous navigation task.
We highlight some of the distinguishing concepts and design choices below, and summarize the related work in Table I. In Passive contingency planning: While prior work has used the following, we refer to Fig. 2, which depicts various ways contingency planning for autonomous navigation, it has not to model inter-agent dependencies.
used learned behavioral models. In “human leader” approaches

Learned behavior model, noncontingent planning:

[11, 12, 23], the behavior prediction of the other agents is

2

independent of the future behavior of the robot, which means that the robot does not model how its future actions can affect decisions of the other agents, shown Fig. 2b. Unlike the active “co-leader” (Fig. 2c) approach, “human leader” methods cannot plan to provoke other agents to reveal their intentions.
Galceran et al. [13] and Fisac et al. [14] perform co-leader contingency planning with hand-crafted models (dynamics models, behavior models, reward functions) rather than learned behavior models. Both Cui et al. [23] and Bajcsy et al.
[24] (a reachability analysis-based approach) only consider single-forking contingencies, whereas our method considers contingencies at every timestep, and therefore does not require determining a ‘branching time’.
Model-free methods: An alternative to using the model-based planning approaches described above is to learn a modelfree policy, which has the capacity to implicitly represent contingent plans. Model-free imitation learning (IL) and reinforcement learning (RL) have been used to construct policies in autonomous driving environments with multi-agent interaction [25–30]. Model-free approaches can be successful when the data and task distribution shifts between train and test are small – a well-trained policy will naturally mimic highquality demonstration data exhibited by a contingent expert or evoked by a well-crafted reward function. However, if there is mismatch between the training and test rewards, the policy will produce suboptimal behavior. While goal-conditioning methods can address the suboptimality of adapting model-free methods to test-time data by serving as the representation for the testtime reward function, the goal space of the agents must be specified a priori, requires goal labels, and precludes adapting the learned system to new types of test-time goal objectives
[27, 30]. Furthermore, in contrast to model-free methods, our method explicitly represents the planned future behavior, which offers interpretability and model introspection benefits.
Our proposed method is both contingent and uses a fullylearned behavior model, unlike prior work that either had learned behavior models or was actively contingent, but never both, as shown in Table I.

Fig. 3: Flowchart that describes the components of our approach.
A behavioral model is trained to forecast multi-agent trajectories conditioned on high-dimensional scene observations, and is used during deployment to construct a contingency planner. The contingency planner receives goals G from an outer navigator, as well as highdimensional observations ot at timestep t from the environment. The contingency planner plans a policy, which predicts the immediate next target position xrt+1 . This next position is passed to a P-controller to produce the necessary action art+1 .

(Fig. 2c). When the policy is open-loop, it is noncontingent, and represents the future positional trajectory of the robot, which can result in underconfident behavior (Fig. 2a). The planning criterion incorporates several components: (1) the learned behavioral model to forecast a probability distribution of the likely behaviors of all agents in terms of positional trajectories; the PDF of this distribution is used in the planning criterion. The planning criterion also incorporates (2) goals specified as positions, and optionally incorporates (3) hard or soft constraints on joint behavior. The planner optimizes for a robot policy that has a high expected probability of joint behavior, and a high expected probability of satisfying the provided goal location and optional goal constraints; the expectation is under the other agents’ sources of uncertainty.
Model design and training: The behavioral model is trained with a dataset of trajectories of multi-agent positions paired with high-dimensional sensory observations of one of the agents.
Training maximizes the likelihood of the trajectories, similar to multi-step Behavioral Cloning [31]. Following prior IL work, we sometimes call these positions behavioral observations,
III. D EEP C ONTINGENCY P LANNING
Below, we outline the deployment phase, planner design, in contrast to demonstrations (observations of state-action pairs). We assume that the behaviors in this data are goaland training phase. Our method is depicted in Fig. 3.
directed, but we do not assume that every behavior would
Deployment: The following loop is executed during deployscore well under the planning criterion: some of these behaviors ment: (1) the contingency planner is given a learned behavioral require navigation to different goals than those seen in training, model, goals, and high-dimensional sensory observations as and some of these trajectories exhibit behavior that may be input, and outputs a multi-step plan in the form of a robot undesirable. This somewhat general assumption on the quality policy. (2) The planned robot policy is executed for one step to of the data behavior demands a planning criterion that involves output a target position, which is (3) tracked by a proportional terms of multi-agent trajectories, which necessitates modeling controller that outputs a control that includes robot steering, the likely trajectories of the other agents.
throttle, and braking controls. (4) The control is executed on the robot in the environment, which produces a new observation. A. Preliminaries
We consider multi-agent systems composed of A ≥ 2
agents (e.g., vehicles). We assume agent positions are fully observable by all agents, but the intentions of each agent (e.g., intended destinations) are hidden from each other. We consider continuous positions and discrete time, and define the position

Planner design: The future behavior of the robot is planned in terms of a parametric policy; a subset of the policy’s parameters are used to optimize a planning criterion. Crucially, this policy is closed-loop – it reacts differently to (i.e., is contingent on) different possibilities of the future behavior of all agents

3

of the ath vehicle at time t by its location in a d-dimensional
Cartesian space as xat ∈ Rd . The joint positions of all vehicles are denoted by xt ∈ RA×d , where a lack of superscript indicates all vehicles. Let xrt = x1t ∈ R1×d index the position of the controlled robot vehicle, and xht = x2:A
∈ R(A−1)×d t
index the positions of other human-driven vehicles that the robot can influence but has no direct control over. Let t = 0 and x0 denote the current timestep and joint position of the multi-agent system, and let x≤t denote future joint positions [1, . . . , t]. Uppercase denotes random variables (e.g.
the stochastic sequence of all multi-agent positions X1:A
≤t ), and lowercase denotes realized values. The current observation of
.
the robot is denoted o = {x0 , i0 }; i0 provides high-dimensional information about the environment. In our implementation, i0 ∈ RH×W is a LIDAR range map image [32].

overall system is still stochastic given the contingency plan parameterized by zr . This method of partially parameterizing a policy is appealing because it requires no modifications to the model. Once qθ is trained, θ is fixed, and planning is a matter of choosing the free parameters, zr , of π r . Other parameterizations could certainly be used, but we found this method efficient. [5, 35] employ similar parameterizations in the noncontingent setting.
A natural question that arises from this discussion is: which contingent plans can such a scheme actually represent? While analyzing this in the continuous case is difficult, we can provide a detailed analysis of this question in the case where all variables (x and z) are discrete. In this case, we show in the Appendix that likely contingent plans can be expressed in terms of zr and can represent sequences of n-th best responses under the model qθ (i.e., first-best, second-best, etc.). Key to
B. Model Design and Training Details this analysis and the general richness of the planning space is
The planner requires a learned behavioral model that the invertibility of m in za , which ensures that all xa ∈ Rd are t
t stochastically forecasts multi-agent trajectories X1:A
realizable outputs of π a . In the Appendix (Fig. 6), we also give
≤T given context o, and captures co-dependencies between agents at all example visualizations of a single zr resulting in high-quality
≤T
timesteps. Let qθ (X|o) denote this model. The model factorizes plans contingent on plausible futures of other agents.
N
into a product of autoregressive distributions parameterized by
We train the model, qθ , with a dataset {(x1:A
≤T , o)n }n=1 , neural networks that receive observations and previous states: depicted in Fig 3. As previously mentioned, we assume that these behaviors are goal-directed to diverse goals, which allows
T Y
A
the model to represent many modes of plausible behavior
Y
1:A
qθ (X1:A
q a (Xat = xat ; φat ),
(1) (e.g., turning left, turning right, traveling passively, traveling
≤T = x≤T |o) =
t=1 a=1
aggressively). We defer implementation details of constructing a
where the parameters of q are computed autoregressively and training qθ to Section IV-B.
in order to model reactions dependent only on past data: C. Planner Design Details a
φat = fθa (x1:A
Primary planning objective: Having described our learned
<t , o). We restrict q to the family of distributions that can be sampled by transforming a sample zat ∈ Rd from behavioral model and a method to parameterize a contingent a fixed simple base distribution chosen a priori, q̄ a (zat ), to plan as a closed-loop policy, we turn to constructing a planning xt , using a fixed invertible transformation of observations and objective in order to evaluate the quality of the planned policy learned values xat = m(zat ; φat ). E.g., if q a = N (·|{µat , σ at }), in terms of how plausible its behavior is and how well it then m(zat ; φat ) = µat + σ at zat , and q̄ a (zat ) = N (zat ; 0, I). Be- satisfies goals at test time. In our application, goals G = (g, G)
cause the q̄ a are independent, the model represents multi-agent are provided to the robot in terms of coordinate destinations behavior as a learned coupling of independent variables, zat . g ∈ R2 , as well as a set of constraints, G, on the joint trajectory
Taken together, these conditions specify that the model is a specified as an indicator function δG (x), which is 1 where
(conditional) autoregressive normalizing flow [33, 34].
x ∈ G and 0 otherwise. To evaluate the quality of a plan, we
Now we show how qθ can be used to design a compact take inspiration from the single-agent (non-contingent) imitative and rich space of contingent policies that is amenable to fast model planning objective [5] by formulating planning with the optimization at planning time. First, note that the process model as maximum a posteriori (MAP) estimation. Instead of of generating agent a’s position at time t with qθ can MAP estimation of a robot trajectory (non-contingent plan), be seen as generating a position from (state-space) policy: we perform MAP estimation of the subset of policy parameters a
q a (Xat | x1:A
zr≤T with posterior p(zr≤T |G, o). The resulting lower-bound
<t , o; θ). Next, note that given zt , m can be written as the application of a deterministic policy parameterized by objective, derived on the supplementary website, is: a
both zat and θ: xat = π a (x1:A
<t , o; zt , θ). On the one hand, qθ
r
A
(2)
can be seen as a collection of stochastic policies {q a }a=1 . On LCfO (πz≤Th) =
i

r the other hand, qθ can be seen as a collection of deterministic log qθ (x̄≤T |o)+log N x̄T ; g, I +log δG (x̄≤T ) .
E
| {z }
| {z }
|
{z
}
zh ∼N (0,I)
policies and policy parameter priors {(π a , q̄ a )}A
a=1 . At planning
(1) prior
(3) constraint
(2) destination time, we can specify a contingent robot plan by deciding values for the robot’s policy parameters, zr≤T ∈ R2×T , rather than Eq. 2 consists of several terms: the (1) behavioral prior sampling each zrt from its prior q̄ r . Doing so results in a encourages the planned policy to result in joint behavior that is contingency plan that is reactive to each potential joint behavior likely under the learned model; (2) the destination likelihood of the other agents, X2:A
encourages the policy to end near a goal location; (3) an
≤t , for t ∈ [1, T ]. Because the future behaviors of the other agents are unobserved at test-time, the optional goal constraints penalizes the policy if it does not

4

satisfy some desired constraint. The expectation results from the uncertainty due to the presence of other (uncontrolled) agents.
In practice, we perform stochastic gradient ascent on Eq. 2 w.r.t.
zr≤T and approximate a violated constraint (log δG (x≤T ) = −∞
when x≤T ∈
/ G) with a large negative value.
Alternative planning objectives: Consider instead directly planning xr≤T under uncertain human trajectories x̂h≤T :
Lr (xr≤T ) = E log N (xrT ;g,I)+log δG ([xr,x̂h ])+log q([xr,x̂h ]|o).
x̂∼q

Planning with Lr will result in a noncontingent plan that is underconfident, because the criterion fails to account for the fact that xrt affects xh>t . This is illustrated in Fig. 2b. For example, we show in our experiments how an underconfident planner leads to a frozen robot that prefers not to enter the intersection and indicate that it intends to turn left, because the plan – a single fixed trajectory – is not affected by whether or not the human driver yields to it. The best trajectory under this criterion is a passive waiting trajectory that stays outside the intersection until the human has passed. Finally, consider planning both xr and xh (Fig. 2d), i.e., assuming control of both the robot and the human trajectories:

Fig. 4: Overhead images for each scenario. The circles denote decision points; blue paths being robot, orange paths being “human”. Left and
Right: The robot tries to turn. At (1), robot decides if to enter the intersection as turning signal to human. At (2), robot decides if to aggressively turn or wait. At (3), if robot previously entered the intersection, human may or may not yield. Again at (2), if robot is waiting, it decides when to turn based on if human yielded. At (4), if robot turned aggressively and human didn’t yield, a near-collision occurs. Middle: Robot tries to overtake. At (1), robot decides if to enter the left lane. At (2), robot decides if to overtake aggressively. At
(3), human decides if to yield. Again at (2), if robot wasn’t aggressive, it decides if to overtake. At (4), if robot was aggressive and human didn’t yield, a near-collision occurs.

Ljoint (x) = log N (xrT ; g, I) + log δG (x) + log qθ (x|o).
This is an overconfident planner, because it assumes that the other agents will obey the robot’s plan for them. Thus, this planner is acceptable when the human behaves as planned, but risky when the human does not. In the left-turn scenario, Ljoint causes the robot to assume that the human will always yield to it, resulting in dangerous behavior when the human sometimes does not. Instead, if the prior of joint behavior is the only component of the planning criterion (i.e. just the final term of
Ljoint ), then the planner simply optimizes for the mode of the prior, which will result in an undirected likely joint behavior.

and a right turn at a stop sign scenario. Images from each scenario are shown in Fig. 4. For each scenario, we identified 4
suitable locations, and used 1 of them as training locations and
3 as test locations. At test time, we evaluate each method 10
times in each test location, resulting in a total of 90 evaluation episodes locations episodes per method (3 scenarios method · 3 scenario · 10 location ). The testing episodes use a fixed set of random seeds in the simulator.
In the left turn scenario, the robot car needs to execute a left turn, but the oncoming vehicle may or may not yield. In the
IV. E XPERIMENTS
highway overtaking scenario, the robot car must pass the car
Our main experimental hypothesis is that (i) some common in front of it, but this requires entering the left lane. Similar driving situations that require multi-agent reasoning can be to how, in the left turn scenario, the other vehicle may or may nearly solved with our deep contingency planning approach. not yield, the leading car in the overtaking scenario may slow
This hypothesis is related to the following secondary hypothe- down or maintain a constant speed to facilitate the robot car’s ses, which state that various noncontingent planning methods overtaking, or it may accelerate to prevent it. In the right turn cannot solve these situations: (ii) a noncontingent learning- scenario, the robot car needs to execute a right turn into a lane based underconfident planner will lead to a frozen robot in that contains a vehicle that may or may not yield.
some situations, (iii) a noncontingent learning-based planner of
These scenarios, while common and easily negotiated by joint trajectories will lead to an unsafe robot in some situations, human drivers, have features that make them difficult for many
(iv) online model-free methods will incur a significant cost learning-based autonomous navigation methods. The optimal of failed behaviors before learning to succeed, in contrast to outcomes (making the turns quickly and safely, and successfully our method, and (v) oblivious planning (unable to adapt to overtaking the other car) require cooperation from the human cost function) will lead to an unsafe robot in some situations. agent, but the intent of the human to cooperate is unknown.
To evaluate these hypotheses, we developed a set of common This calls for contingent planning. However, in order to infer driving scenarios in which contingency planning matters.
the human’s intent, we (the robot) must first act in a way that signals our own desire (entering the intersection or the opposite
A. Benchmark Scenarios
We evaluate our method and several others on three multi- lane) and evaluate their response.
agent driving scenarios that we built in the CARLA simulator B. Implementation Details
[25]: a left-turn scenario similar to the one discussed in Fig. 1
We designed these scenarios in suitable locations in the and conceptually similar to that in [11], a highway overtaking CARLA simulator [25]. We generated training data from handscenario conceptually similar to that used in prior work [14, 15], crafted policies designed to generate each of the outcomes in

5

Left Turns

Overtaking

Right Turns

Method

RG

RG*

RG

RG*

RG

RG*

Single-agent [5]
Noncontingent, Lr
Noncontingent, Ljoint
MFP [17]
CfO (Eq. 2) (ours)

30/30
30/30

16/30
0/30
9/30
9/30
28/30

10/30
30/30

8/30
0/30
12/30
9/30
27/30

15/30
30/30

2/30
0/30
9/30
9/30
27/30

30/30
30/30
29/30

29/30
16/30
27/30

23/30
9/30
27/30

eventually reached the goal region, and near-expert success rate (“RG*”), which we define as the vehicle reaching the goal as quickly as an expert would, without any near-collisions.
Examining Table II, we see evidence that supports the hypotheses mentioned at the beginning of the section. First,
(i) the CfO planner (i.e., the one that uses Eq. 2) has near perfect success rate in both scenarios, as well as a nearly perfect rate of near-expert success. The high rate of completing the maneuver in a near-expert way suggests that contingency planning is helpful for efficiently navigating these complex scenarios, which is possible only by probing the human’s intentions and reacting accordingly. Next, (ii) our
“underconfident” noncontingent planner, while having a high success rate, has a much lower count of near-expert successes than the contingent CfO planner. The reason is that the underconfidence leads to “frozen robot” behavior that slowed down the underconfident planner. In the left turn scenario, for example, the underconfident planner will always wait for the oncoming vehicle to pass through the intersection before beginning its own left turn. This means the underconfident planner missed any opportunity to negotiate the intersection as quickly as the expert did in the trials where the human vehicle would have yielded.
For (iii), note that the “overconfident” noncontingent planner always reaches the goal, but sometimes engages in unsafe behavior. Its aggressive planning means that it will try to solve the scenario quickly every time, which leads to near-collisions every time that the other vehicle does not yield.
For (iv), we observe the “oblivious” planner [5] indeed performs suboptimality across scenarios, leading to more nearcollisions than the other planners. Due to the suboptimality present in the training data, its planner, which reasons solely about the ego-vehicle’s behavior, is insufficient in these scenarios. Finally, for (v), we compare our approach to a model-free reinforcement learning (RL) baseline, Proximal
Policy Optimization (PPO), on the left turn scenario.
As seen in Fig. 5, online model-free methods are capable of learning an optimal policy. However, the RL agent incurs a significant number of collisions before reaching a reasonable success rate (in contrast, since CfO learns entirely offline, it need not experience collisions before learning how to complete the task). Additionally, learning a successful policy with PPO
requires careful reward shaping on top of the cost model used in CfO. Further details about the RL experimental setup, as well as example qualitative results of contingent and noncontingent plans, are given in the Appendix.
V. D ISCUSSION
We present a deep contingent planning method for control in multi-agent environments, and demonstrate its efficacy in common scenarios encountered in the urban driving setting in comparison to several alternative methods for data-driven planning, including a single-agent planner and a deep MPCbased trajectory shooting method. Our method is quite tractable to apply – it is learned entirely from demonstrations of plausible multi-agent behavior and efficiently represents contingent plans in terms of components of the learned model.

Table II: Comparison of learning-based planning for control: Our co-Influence model (CfO) learns how agents behave together, and is able to control a vehicle to negotiate multi-agent driving scenarios safely and efficiently. Ablations show that alternative non-contingent planners lead to either overly conservative or unsafe trajectories.

each scenarios: first, the robot chooses whether to begin its maneuver (enter the intersection or the opposite lane). If the robot begins the maneuver, the other vehicle decides whether to accommodate the it (yield to it or allow passing). Then, the robot decides whether to attempt to complete the maneuver.
The robot’s policy can be considered a suboptimal expert, as it sometimes generates expert paths, but other times, it generates conservative paths (by sometimes not entering the intersection or opposite lane) and risky paths (by sometimes attempting to complete the maneuver despite the other agent not yielding).
Examples are given in the Appendix (Fig. 7).
We adapted the ESP model from [36], which uses qta = N
and we used 1.5s of past at 10Hz and 8s of future at 3.75Hz, resulting in pasts of length 15 and futures of length 30. In contrast to the LIDAR featurization in [36], we instead use i0 ∈ RH×W as a LIDAR range map image [32]. Whereas [36]
performed passive forecasting of behaviors of agents, and was not used to decide controls of a robot, our work focuses on building and executing actively contingent plans.
Method and baselines: We trained a single global CfO model, which forces the method to generalize across scenarios and locations. For comparison, we use [17] (MFP) by following the shooting-based planning method described therein. We experimented with a global model for MFP, but were unable to achieve good performance, instead, we trained separate per-scenario MFP models. We use the CfO model with the contingent planning objective (Eq. 2) as well as with the
“overconfident” noncontingent planner with criterion Ljoint from
Section III-B. For the “underconfident” planner, we used the suboptimal expert to create the set of per-agent paths available at each timestep, and used Lr without the prior term, which serves as a best-case underconfident planner (in the case of a perfect model). To measure the importance of multi-agent planning vs. an oblivious planner, we apply [5] by constructing a single-agent version of our model, which cannot incorporate terms that involve predicted positions of other agents. For this baseline, like MFP, we trained separate per-scenario models.
Further details are given in the Appendix.
C. Results
Prior navigational planning work uses a variety of metrics to measure the quality of a planned or predicted trajectory
[17, 36]. In this work, we evaluate our different planners in terms of closed-loop control metrics: the ability of a method to reach a goal (1) successfully, and (2) safely and efficiently.
In Table II, we report results in terms of the goal-reaching success rate (“RG”), which simply means that the robot vehicle

6

Vehicles,” in International Conference on Robotics and Automation
(ICRA), 2019, pp. 9590–9596.
[15] D. Sadigh, S. Sastry, S. A. Seshia, and A. D. Dragan, “Planning for autonomous cars that leverage effects on human actions,” in Robotics:
Science and Systems (RSS), 2016.
[16] E. Schmerling, K. Leung, W. Vollprecht, and M. Pavone, “Multimodal probabilistic model-based planning for human-robot interaction,” in
International Conference on Robotics and Automation (ICRA). IEEE,
2018, pp. 3399–3406.
[17] Y. C. Tang and R. R. Salakhutdinov, “Multiple futures prediction,” in
Neural Information Processing Systems (NeurIPS), vol. 32, 2019, pp.
15 424–15 434.
[18] S. M. LaValle, Planning algorithms. Cambridge University Press, 2006.
[19] W. Schwarting, J. Alonso-Mora, and D. Rus, “Planning and decisionmaking for autonomous vehicles,” Annual Review of Control, Robotics, and Autonomous Systems, 2018.
[20] T. Bandyopadhyay, K. S. Won, E. Frazzoli, D. Hsu, W. S. Lee, and
D. Rus, “Intention-aware motion planning,” in Algorithmic foundations of robotics X. Springer, 2013, pp. 475–491.
[21] W. Xu, J. Pan, J. Wei, and J. M. Dolan, “Motion planning under uncertainty for on-road autonomous driving,” in International Conference on Robotics and Automation (ICRA). IEEE, 2014, pp. 2507–2512.
[22] B. Zhou, W. Schwarting, D. Rus, and J. Alonso-Mora, “Joint multipolicy behavior estimation and receding-horizon trajectory planning for automated urban driving,” in International Conference on Robotics and
Automation (ICRA). IEEE, 2018, pp. 2388–2394.
[23] A. Cui, A. Sadat, S. Casas, R. Liao, and R. Urtasun, “LookOut: Diverse multi-future prediction and planning for self-driving,” arXiv preprint arXiv:2101.06547, 2021.
[24] A. Bajcsy, A. Siththaranjan, C. J. Tomlin, and A. D. Dragan, “Analyzing human models that adapt online,” in International Conference on Robotics and Automation (ICRA). IEEE, 2021.
[25] A. Dosovitskiy, G. Ros, F. Codevilla, A. Lopez, and V. Koltun, “CARLA:
An Open Urban Driving Simulator,” in Conference on Robot Learning
(CoRL), Mountain View, USA, Nov. 2017.
[26] J. Chen, B. Yuan, and M. Tomizuka, “Deep imitation learning for autonomous driving in generic urban scenarios with enhanced safety,”
arXiv preprint arXiv:1903.00640, 2019.
[27] F. Codevilla, E. Santana, A. M. López, and A. Gaidon, “Exploring the limitations of behavior cloning for autonomous driving,” in International
Conference on Computer Vision (ICCV). IEEE, 2019, pp. 9329–9338.
[28] Y. C. Tang, “Towards learning multi-agent negotiations via self-play,” in
International Conference on Computer Vision Workshops, 2019.
[29] P. Palanisamy, “Multi-agent connected autonomous driving using deep reinforcement learning,” arXiv preprint arXiv:1911.04175, 2019.
[30] J. Hawke, R. Shen, C. Gurau, S. Sharma, D. Reda, N. Nikolov, P. Mazur,
S. Micklethwaite, N. Griffiths, A. Shah et al., “Urban driving with conditional imitation learning,” in International Conference on Robotics and Automation (ICRA). IEEE, 2020, pp. 251–257.
[31] D. A. Pomerleau, “ALVINN: An autonomous land vehicle in a neural network,” in Neural Information Processing Systems (NeurIPS), 1989, pp. 305–313.
[32] L. Caccia, H. Van Hoof, A. Courville, and J. Pineau, “Deep generative modeling of lidar data,” arXiv preprint arXiv:1812.01180, 2018.
[33] L. Dinh, J. Sohl-Dickstein, and S. Bengio, “Density estimation using real
NVP,” in International Conference on Learning Representations (ICLR),
2017. [Online]. Available: https://openreview.net/forum?id=HkpbnH9lx
[34] G. Papamakarios, T. Pavlakou, and I. Murray, “Masked autoregressive flow for density estimation,” in Neural Information Processing Systems
(NeurIPS), 2017, pp. 2338–2347.
[35] A. Singh, H. Liu, G. Zhou, A. Yu, N. Rhinehart, and S. Levine, “Parrot:
Data-driven behavioral priors for reinforcement learning,” arXiv preprint arXiv:2011.10024, 2020.
[36] N. Rhinehart, R. McAllister, K. Kitani, and S. Levine, “PRECOG:
PREdiction CONditioned on Goals in Visual Multi-Agent Settings,” in
International Conference on Computer Vision (ICCV). IEEE, 2019, pp.
2821–2830.
[37] A. Hill, A. Raffin, M. Ernestus, A. Gleave, A. Kanervisto, R. Traore,
P. Dhariwal, C. Hesse, O. Klimov, A. Nichol, M. Plappert, A. Radford,
J. Schulman, S. Sidor, and Y. Wu, “Stable baselines,” https://github.com/
hill-a/stable-baselines, 2018.
[38] D. Tran, K. Vafa, K. Agrawal, L. Dinh, and B. Poole, “Discrete flows:
Invertible generative models of discrete data,” in Neural Information
Processing Systems (NeurIPS), 2019, pp. 14 719–14 728.

Fig. 5: Left: Reward vs steps of experience on the left turn scenario
(mean/std/min/max over three runs, best policy plotted in solid).
Middle: Collision rate vs steps. PPO is the only baseline that requires online training (all other methods are horizontal lines). PPO
(acceleration only, similar to [17]) achieves a 0% collision rate. Right:
Success rate vs steps. PPO (acceleration only) achieves a 100% success rate. PPO (throttle, steering and braking) is unable to exceed a 60%.

Acknowledgements: This research was supported by the Office of Naval Research, ARL DCIST CRA W911NF-17-2-0181, JP
Morgan, the Berkeley DeepDrive industry consortium, and the supporters of the Berkeley RISELab.
R EFERENCES
[1] S. Russell, Human compatible: Artificial intelligence and the problem of control. Penguin, 2019.
[2] S. Thrun, M. Montemerlo, H. Dahlkamp, D. Stavens, A. Aron, J. Diebel,
P. Fong, J. Gale, M. Halpenny, G. Hoffmann et al., “Stanley: The robot that won the DARPA Grand Challenge,” Journal of field Robotics, vol. 23, no. 9, pp. 661–692, 2006.
[3] C. Urmson, J. Anhalt, D. Bagnell, C. Baker, R. Bittner, M. Clark, J. Dolan,
D. Duggins, T. Galatali, C. Geyer et al., “Autonomous driving in urban environments: Boss and the urban challenge,” Journal of Field Robotics, vol. 25, no. 8, pp. 425–466, 2008.
[4] B. Paden, M. Čáp, S. Z. Yong, D. Yershov, and E. Frazzoli, “A survey of motion planning and control techniques for self-driving urban vehicles,”
IEEE Transactions on intelligent vehicles, vol. 1, no. 1, pp. 33–55, 2016.
[5] N. Rhinehart, R. McAllister, and S. Levine, “Deep imitative models for flexible inference, planning, and control,” in International Conference on Learning Representations (ICLR), 2020. [Online]. Available: https://openreview.net/forum?id=Skl4mRNYDr
[6] W. Zeng, W. Luo, S. Suo, A. Sadat, B. Yang, S. Casas, and R. Urtasun,
“End-to-end interpretable neural motion planner,” in Computer Vision and Pattern Recognition (CVPR), 2019, pp. 8660–8669.
[7] A. Filos, P. Tigas, R. McAllister, N. Rhinehart, S. Levine, and Y. Gal,
“Can autonomous vehicles identify, recover from, and adapt to distribution shifts?” in International Conference on Machine Learning (ICML), 2020.
[8] S. Thompson, T. Horiuchi, and S. Kagami, “A probabilistic model of human motion and navigation intent for mobile robot path planning,” in
2009 4th International Conference on Autonomous Robots and Agents.
IEEE, 2009, pp. 663–668.
[9] B. D. Ziebart, N. Ratliff, G. Gallagher, C. Mertz, K. Peterson, J. A.
Bagnell, M. Hebert, A. K. Dey, and S. Srinivasa, “Planning-based prediction for pedestrians,” in 2009 IEEE/RSJ International Conference on Intelligent Robots and Systems. IEEE, 2009, pp. 3931–3936.
[10] P. Trautman and A. Krause, “Unfreezing the robot: Navigation in dense, interacting crowds,” in 2010 IEEE/RSJ International Conference on
Intelligent Robots and Systems. Taipei: IEEE, Oct. 2010, pp. 797–803.
[11] J. Hardy and M. Campbell, “Contingency Planning Over Probabilistic
Obstacle Predictions for Autonomous Road Vehicles,” IEEE Transactions on Robotics, vol. 29, no. 4, pp. 913–929, 2013.
[12] W. Zhan, C. Liu, C.-Y. Chan, and M. Tomizuka, “A non-conservatively defensive strategy for urban autonomous driving,” in International
Conference on Intelligent Transportation Systems (ITSC). IEEE, 2016, pp. 459–464.
[13] E. Galceran, A. G. Cunningham, R. M. Eustice, and E. Olson, “Multipolicy decision-making for autonomous driving via changepoint-based behavior prediction: Theory and experiment,” Autonomous Robots, vol. 41, no. 6, pp. 1367–1382, 2017.
[14] J. F. Fisac, E. Bronstein, E. Stefansson, D. Sadigh, S. S. Sastry, and
A. D. Dragan, “Hierarchical Game-Theoretic Planning for Autonomous

7

A PPENDIX

trajectories for all agents, conditioned on the past trajectories of all agents, the context view (the trajectories and visual context
A. Implementing the MFP baseline both processed by an encoder), and the mode: p(Ya |X, I, k),
We were required to make several key changes to the publicly paramaterized by a 2D Gaussian distribution over xy coordireleased MFP codebase to create a usable MFP baseline for nates for each agent at every predicted timestep t, i.e., N (µ, ρ)
2
a,t a,t a,t our contingency planning scenarios in CARLA: a,t represented by the 5-tuple (µa,t
X , µY , σX , σY , ρ ). A given
1) The MFP codebase only includes code for the NGSIM agent a’s history Xa and an attention-based encoding of the dataset, so we added support for training (and visualizing) history of other agents Attn(Xi6=a ) are concatenated in a vector models on CARLA datasets, in addition to adding support [Xa , Attn(Xi6=a ), I] and passed as input to the encoding RNN, for deploying the model in the CARLA simulator. Since which produces a latent encoding and a Multinoulli distribution
MFP is a model for prediction (not planning), the latter over modes per-agent, p(K a |Xa , Attn(Xi6=a ), I). The latent required implementing a planning procedure to turn encoding is then input to each decoding RNN (one per mode)
predictions from the MFP model into control for the to generate K predicted trajectories per agent. All agents share ego vehicle (see Appendix B).
the same set of (encoder and decoder) RNNs, so parameters
2) We added rotation of model inputs based on the each scale linearly with modes, whereas the number of agents scales vehicle’s yaw as an additional preprocessing stage, such the batch/input dimension.
that the reference frame for prediction always starts with
To summarize, the outputs of the trained MFP model the ego vehicle at (0,0) and pointing in the +X direction represent the following distributions:
(the existing repository only shifts the inputs to (0,0)).
p(Ya |X, I, k), a distribution over future coordinates, perWe observed that rotation of inputs was necessary for agent, per-mode, generated by the decoding RNNs, and training on the left and right turn scenarios.
p(K a |X, I), a distribution over modes per-agent, generated by
3) We removed the attention stage of the model architecture the encoding RNN. This means the model generates K futures which significantly improved model performance on our for each agent, i.e., K × A total trajectory distributions.
CARLA datasets (recall that the scenarios we consider
Note that the MFP model outputs do not include an explicit in this paper only include two agents).
joint trajectory distribution p(Y|X, ...); the conditional joint
4) We removed the pretraining stage during model training future is factorized across agents, and across time (using the
(where the RNN/decoder output sequence is generated autoregressive decoder): in a non-interactive fashion) which also significantly
T Y
A
Y
improved model performance.
a p(Y |X, I, k) =
p(yta |Y≤t , X, I, k a )
(3)
A repository containing a version of the MFP codebase with the t=1 a=1
aforementioned changes can be found on the project website. We note that in Eq. 3 (and its original formulation in [17]), the
An important difference between CfO and MFP is that with joint is conditioned on the (ego) agent’s own k, but not on the
CfO, we can train a single model on multiple scenarios, whereas ks of other agents. However, in the reference implementation with MFP, we are limited to training a single model for each of MFP1 (and our own code), the joint is conditioned on all individual scenario. We were unable to train MFP models agents having the same k value.
on multiple scenarios that generated reasonable predictions
Planning using the same cost model as CfO requires explicit
(usable for planning): we found that the mixed-scenario models joint trajectories. Since the MFP model outputs a distribution of would often generate trajectories from the wrong scenario K trajectories for each of the A agents, to generate candidate
(e.g., a left turn on the right turn scenario). Additionally, we joint trajectories Ŷ in a two-agent setting, we can simply found it necessary to tune hyperparameters per-scenario, e.g., consider the set of all K × K possible combinations of preprocessing with rotation is critical for the left and right individual ego and actor trajectories: turn scenarios, but degrades performance significantly on the
Ŷ = [(µYego , µYactor ) for k ego , k actor ∈ K]
overtake scenario.
(4)
ego actor
= [(µego
, µactor
) for k ego , k actor ∈ K]
Y
X , µY , µX
B. Planning with MFP
In our MFP experiments we used 5 modes, so we consider
Similar to our behavior model described in Section III-B,
25 total candidate joint trajectories. To select the best ego
MFP’s forward pass (used during inference) generates predicted trajectory to use for planning, we can select the ego portion trajectories for X≤T . Moreover, MFP generates K different of the joint with the lowest cost (highest reward), weighted by a,k trajectories for each ve

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

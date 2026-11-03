---
source: "https://arxiv.org/abs/2603.10747v2"
title: "Pneuma-Seeker: Relational Reification of Information Needs for Agentic Data Discovery and Preparation"
author: "Muhammad Imam Luthfi Balaka, John Hillesland, Kemal Badur, Raul Castro Fernandez"
year: "2026"
publication: "arXiv preprint / cs.DB"
download: "https://arxiv.org/pdf/2603.10747v2"
pdf: "https://arxiv.org/pdf/2603.10747v2"
captured_at: "2026-11-03T13:57:00Z"
updated_at: "2026-11-03T13:57:00Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Lukács reification"
tags:
  - "現代哲学"
  - "マルクス主義"
  - "西洋マルクス主義"
  - "物象化論"
status: raw
---

# Pneuma-Seeker: Relational Reification of Information Needs for Agentic Data Discovery and Preparation

- 著者: Muhammad Imam Luthfi Balaka, John Hillesland, Kemal Badur, Raul Castro Fernandez
- 年: 2026
- 掲載情報: arXiv preprint / cs.DB
- 情報源: [arxiv](https://arxiv.org/abs/2603.10747v2)
- ダウンロード: https://arxiv.org/pdf/2603.10747v2
- PDF: https://arxiv.org/pdf/2603.10747v2

## Obsidian Links

- 研究動向: [[ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代哲学]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[西洋マルクス主義]]
- 関連分野: [[物象化論]]
- 関連タグ: #現代哲学 #マルクス主義 #西洋マルクス主義 #物象化論

## Abstract

When faced with data problems, many data workers cannot articulate their information need precisely enough for software to help. Although LLMs interpret natural-language requests, they behave brittly when intent is under-specified, e.g., hallucinating fields, assuming join paths, or producing ungrounded answers. We present Pneuma-Seeker, a system built around a central idea: relational reification. Pneuma-Seeker represents a user's evolving information need as a relational schema: a concrete, analysis-ready data model shared between user and system. Rather than answering prompts directly, Pneuma-Seeker iteratively refines this schema, then discovers and prepares relevant sources to construct a relation and executable program that compute the answer. Pneuma-Seeker employs an LLM-powered agentic architecture with conductor-style planning and macro- and micro-level context management to operate effectively over heterogeneous relational corpora.

## PDF Text

Pneuma-Seeker: Relational Reification of Information Needs for
Agentic Data Discovery and Preparation
Muhammad Imam
Luthfi Balaka luthfibalaka@uchicago.edu
The University of Chicago
Chicago, USA

John Hillesland

Kemal Badur

Raul Castro Fernandez

jhillesland@uchicago.edu
The University of Chicago
Chicago, USA

badur@uchicago.edu
The University of Chicago
Chicago, USA

raulcf@uchicago.edu
The University of Chicago
Chicago, USA

ABSTRACT
When faced with data problems, many data workers cannot articulate their information need precisely enough for software to help. Although LLMs interpret natural-language requests, they behave brittly when intent is under-specified, e.g., hallucinating fields, assuming join paths, or producing ungrounded answers.
We present Pneuma-Seeker, a system built around a central idea: relational reification. Pneuma-Seeker represents a user’s evolving information need as a relational schema: a concrete, analysis-ready data model shared between user and system. Rather than answering prompts directly, Pneuma-Seeker iteratively refines this schema, then discovers and prepares relevant sources to construct a relation and executable program that compute the answer. Pneuma-Seeker employs an LLM-powered agentic architecture with conductorstyle planning and macro- and micro-level context management to operate effectively over heterogeneous relational corpora.
We evaluate Pneuma-Seeker across multiple domains against state-of-the-art academic and industrial baselines, demonstrating higher answer quality and significant support for iterative refinement. Real-world deployment highlights trust and inspectability as essential requirements for LLM-mediated data systems.
PVLDB Reference Format:
Muhammad Imam Luthfi Balaka, John Hillesland, Kemal Badur, and Raul
Castro Fernandez. Pneuma-Seeker: Relational Reification of Information
Needs for Agentic Data Discovery and Preparation. PVLDB, 19(13): 5230 5243, 2026.
doi:10.14778/3849398.3849417
PVLDB Artifact Availability:
The source code, data, and/or other artifacts have been made available at https://github.com/TheDataStation/pneuma-seeker.

1

INTRODUCTION

Organizations increasingly make operational, financial, and policy decisions “in the loop” with data: analysts, engineers, and domain experts continuously formulate questions, locate relevant datasets, reconcile definitions, and assemble tables or features that can support a decision. In principle, this workflow should be accelerable: a decision maker expresses an information need, and a
This work is licensed under the Creative Commons BY-NC-ND 4.0 International
License. Visit https://creativecommons.org/licenses/by-nc-nd/4.0/ to view a copy of this license. For any use beyond those covered by this license, obtain permission by emailing info@vldb.org. Copyright is held by the owner/author(s). Publication rights licensed to the VLDB Endowment.
Proceedings of the VLDB Endowment, Vol. 19, No. 13 ISSN 2150-8097.
doi:10.14778/3849398.3849417

system retrieves and prepares the required data. In practice, however, data work remains hampered by two stubborn stages in the data management lifecycle: i) data discovery, identifying and retrieving documents that could satisfy an information need, and; ii) data preparation, transforming those documents into a usable representation for downstream analysis. These stages resist automation because determining intent is expensive: users rarely articulate what they need precisely enough for software systems to help.
Large language models (LLMs) appear to change the landscape.
They can interpret natural-language requests, synthesize explanations, and increasingly orchestrate multi-step workflows. This has led to a tempting conclusion: if LLMs can “understand” a user request, perhaps they can also discover the right datasets, assemble them, and produce the answer. Our experience building and deploying Pneuma-Seeker suggests a more careful view. The bottleneck is not merely that the system cannot execute steps; it is that the user’s information need is typically vague, evolving, and difficult to operationalize. Questions are moving targets, and without a concrete target, even very capable models behave brittly, e.g., hallucinating missing fields, assuming join paths, or producing plausible but ungrounded reasoning. What is needed is a mechanism that lets users iteratively converge on a precise intent while giving the system a representation it can reliably act on.
Use Case: Purchase Order. Consider a procurement office at a large organization trying to cut costs by reducing re-shipping—items that need to be returned and redelivered. A high-level objective
(“reduce re-shipping”) is not directly answerable with data. An analyst may hypothesize that certain categories of items cause more downstream friction and ask a question like: “Do hazardous materials exhibit higher rates of post-purchase order execution friction?”
Yet the organization’s data may not contain a clean “hazardous”
category. Through cross-team conversations the analyst might narrow the concept to radioactive materials; “execution friction” may be clarified into measurable signals such as over-shipment, overinvoicing, or forced invoice matching; and the relevant sources may turn out to be spread across multiple systems maintained by different teams. Even after the question becomes crisp, the analyst still faces a discovery and preparation problem: identify which tables contain the relevant signals, determine join keys and grain, reconcile definitions, and produce an analysis-ready relation.
This vignette illustrates three challenges of real data work:
• First, information needs are often under-specified initially and become precise only through iteration.

• Second, the iteration is socially and operationally expensive: it requires coordination and knowledge transfer across teams, tools, and institutional memory.
• Third, even after intent stabilizes, retrieval and assembly of the needed data artifacts remains non-trivial.
The ideal system would help users sharpen their intent while reducing cross-team coordination cost, and then would automatically discover and prepare the data needed to answer the now-precise question. Such a system is not a replacement for human communication; instead, it catalyzes convergence toward high-value questions and makes the path from intent to data explicit and verifiable.
This paper investigates whether and how that ideal can be approximated in practice. We present Pneuma-Seeker, an LLMpowered system for interactive data discovery and preparation, engineered around relational reification: representing a user’s evolving information need as a relational schema. Instead of treating the user’s question as a single prompt to be answered, PneumaSeeker treats it as a specification that must be negotiated and made concrete. The system proposes a target relation—attributes, types, and semantics that would constitute an analysis-ready table for the question—and iteratively refines that schema through user feedback. The schema becomes a shared artifact: it is concrete enough for users to critique (“’hazardous’ is too broad; use ’radioactive”’), and structured enough for the system to operationalize discovery
(“find columns and join paths that can populate these attributes”).
Relational reification, in the manner we present it here, is new; it changes the nature of the interaction: the user is no longer forced to perfectly describe intent upfront, but can react to a proposed schema, correct it, and progressively expose latent assumptions; the system, in turn, is no longer forced to guess what the answer should look like, but is tasked with producing a relation that matches a specified schema. This aligns the user’s language-level intent with the system’s data-level operations and enables explainability and verification: once the target schema is explicit, Pneuma-Seeker can explain, e.g., which sources contributed to which fields and what transformations were applied. The user can inspect the explanation before trusting the result downstream.
Making relational reification work in practice is itself a systems problem: a schema is only useful if the system can reliably discover sources, populate fields with grounded evidence, and revise the schema as understanding improves. Pneuma-Seeker is the system we engineer to make this possible. It decomposes discovery and preparation into specialized subtasks with narrowly scoped contexts—schema refinement, source retrieval, table materialization, join/path reasoning, and program synthesis—orchestrated by a conductor that revises plans based on intermediate results. It also introduces micro context management: a deliberate databaseprobing action that lets agents pull targeted evidence from tabular data (e.g., value distributions, schema samples, join statistics) on demand, rather than relying solely on static sample values as a table’s representation. This matters because tabular corpora and enterprise catalogs are large and heterogeneous, and placing all context into a single prompt is often ineffective or infeasible.
We evaluate i) the system’s ability to support users in inspecting intermediate states and verifying results during iterative data discovery, ii) the quality of the discovered/prepared relations and

the resulting answers, iii) the contribution of context specialization and conductor-style planning relative to alternative designs and baselines, iv) the performance of the system. We trace observed gains back to Pneuma-Seeker’s architectural components, e.g., extending our prior table-discovery work [3] to reach near-perfect retrieval recall (Section 6.4.3) and showing that ablating relational reification alone degrades answer quality (Section 6.2.3). We also report learnings from deploying Pneuma-Seeker in a real organizational setting (at The University of Chicago), using observed user feedback on explainability, iteration cost, and trust to distill design requirements for this emerging class of LLM-based data systems and, concretely, the design of Pneuma-Seeker.
In summary, this paper makes the following contributions:
• Relational Reification. We introduce the idea of representing an evolving information need as an executable relational schema, allowing users and the system to iteratively negotiate intent.
• Micro Context Management. We introduce a deliberate, database-probing action for agents to acquire targeted evidence from tabular data (e.g., value distributions, schema samples, join statistics), augmenting static sample values as a table’s representation.
This action complements the now-standard practice of macrocontext management: splitting a task across specialized agents.
• The Pneuma-Seeker System. We design and implement an agentic system that operationalizes relational reification through context management strategies to discover data, materialize relations, and support provenance-aware iterative refinement.
The rest of the paper is structured as follows. Section 2 introduces terminology, the problem statement, and requirements derived from our experience deploying Pneuma at our university.
Section 3 presents the idea of relational reification, which serves as a nexus between human and machine. Section 4 then shows how this idea is operationalized through context management techniques and the resulting system architecture. Section 5 describes the details of the system. We then present evaluation results (Section 6), followed by related work (Section 7) and conclusions (Section 8).

2

INFORMATION NEEDS ON TABULAR DATA

This section introduces terminology, presents the problem and its motivation, and states four requirements for a technical solution that we have gathered from deploying the system at our university and gathering feedback from users.

2.1

Key Terminology

Document. A document is a concrete representation of data [8].
For example, data about a customer could be represented as a piece of paper, a table in a database, or a PDF; all three are documents. In this paper, we concentrate on tabular documents, such as tables in databases, CSV and Parquet files, and spreadsheets.
Information Need. An information need characterizes the data required to solve a data-driven task [8]. In this work, we assume a one-to-one correspondence between a task and its associated information need, and are interested in finding a document that satisfies that need, i.e., that represents the data required to solve the task that induced it.

In Use Case: Purchase Order, the information need characterizes comparable signals of post-purchase order execution friction for specific item categories across procurement, logistics, and finance systems, sufficient to assess which categories contribute disproportionately to re-shipping. This information need is fulfilled by surfacing data that aligns item categories (e.g., radioactive materials) and execution frictions (e.g., over-shipment), allowing the analyst to identify categories that drive re-shipping and ultimately inform interventions toward the high-level objective.
However, even when the task is well-defined, users are often unable to fully and precisely articulate the associated information need at the outset. Relevant variables, subpopulations, or confounding factors may only become apparent after inspecting data, and the adequacy of a data specification may depend on evidence uncovered during analysis. We capture the resulting gap between what is ultimately required and what the user can articulate by distinguishing between latent and active information needs.
Latent Information Need. A latent information need I ∗ denotes the complete and precise characterization of the data required to solve the task. I ∗ is typically unknown to the user a priori, as it may depend on relevant structure in the data (e.g., subpopulations or distributional properties) that is not evident before exploration.
Active Information Need. The active information need I + corresponds to the user’s articulation of I ∗ at a given point in time. I +
reflects their current hypotheses and assumptions and may omit relevant constraints. Through interaction with data and intermediate results, the user may revise I + to better approximate I ∗ [5].
In Use Case: Purchase Order, the analyst may initially specify I +
as comparing friction rates between hazardous and non-hazardous products. After observing that hazardous products span several categories (radioactive, toxin, etc.) with substantially different sample sizes, the analyst may refine I + to stratify by category and control for sample size, moving closer to I ∗ .

2.2

Problem Statement and Motivation

Data discovery is the problem of identifying and retrieving documents that satisfy an information need [9, 14]. This problem arises across many domains. For instance, analysts in a retail organization may integrate customer profiles and transaction records to study purchasing behavior [23]. Similarly, domain scientists may explore open data repositories (e.g., Chicago Open Data) to analyze behavioral or socioeconomic phenomena [36]. More formally:
Given a collection of tables C = {𝑇1, . . . ,𝑇𝑛 } and a latent information need I ∗ , the goal is to construct a single document D derived from C such that D fulfills I ∗ .
Iterative Discovery Process. Solving the data discovery problem requires gathering documents across heterogeneous sources [1, 18]
that are relevant to I ∗ . A fundamental challenge, however, is that
I ∗ is rarely directly observable by the discovery mechanism. Instead, discovery is mediated through I + , the user’s current (and potentially incomplete) articulation of I ∗ .
We formalize discovery as an iterative interaction process between a user O and a discovery mechanism Ψ:
𝜎 = ⟨(𝐼 0+, 𝐷 0 ), (𝐼 1+, 𝐷 1 ), . . . ⟩
At each iteration 𝑖:

• Articulation. O expresses an active information need I𝑖 + that approximates I ∗ (edge case: I0+ = I ∗ ).
• Synthesis. Ψ produces a document 𝐷𝑖 = Ψ(I𝑖 +, 𝐶). Importantly,
𝐷𝑖 may not fully satisfy I𝑖 + due to limitations in either C or Ψ
(e.g., relevant tables may not exist or may not be retrieved).
• Inspection and Refinement. O inspects 𝐷𝑖 and may subjectively update their understanding of I ∗ , producing a refined
+ .
articulation I𝑖+1
Towards Convergence. Because I ∗ is latent, I ∗ and I𝑖 + do not inhabit a common, observable metric space. Consequently, Ψ cannot natively compute their distance, nor can it directly optimize for the fulfillment of I ∗ . Instead, Ψ must act as a mechanism that enables
O to systematically minimize this latent discrepancy.
As a baseline, Ψ must prioritize high-fidelity alignment with I𝑖 + , i.e., producing 𝐷𝑖 that fulfills I𝑖 + as much as possible (subject to the available information in C). If this were not the case, O might not be able to discern whether an inadequate outcome stems from an inherent limitation of information in C (e.g., relevant tables do not exist) or a failure of Ψ itself (e.g., relevant tables were not retrieved).
Beyond alignment with I𝑖 + , Ψ acts as a catalyst for convergence, driving I𝑖 + → I ∗ by justifying the synthesis of 𝐷𝑖 . This interaction leverages the theory of External Cognition [44], which posits that concrete external representations facilitate computational offloading by providing the necessary “external anchors” for human reasoning. In this context, Ψ minimizes the cognitive load of the iterative discovery process by providing an interface that exposes the justification for 𝐷𝑖 (i.e., the steps Ψ took to produce it). By presenting the current state of discovery (𝐷𝑖 alongside its justification),
Ψ establishes the necessary feedback loop for O to systematically
+ , ultimately converging toward I ∗ .
refine I𝑖 + into I𝑖+1
Scope of the Formalization. This formalization does not yield convergence guarantees or approximation bounds on I𝑖 + → I ∗ : because I ∗ is never observable to Ψ, no common metric space exists over which such a distance could even be defined, so these guarantees are undefined under this framework, not merely difficult to obtain. Instead, the formalization supplies a vocabulary for justifying concrete design decisions: why Ψ must prioritize highfidelity alignment with I𝑖 + rather than guess at I ∗ directly, and why exposing the justification for 𝐷𝑖 —rather than optimizing some latent objective—is the primary lever available to Ψ. Responsibility for closing the gap toward I ∗ remains with O; Ψ’s role is to make that iteration cheaper and better grounded.
Example Ψ. One common discovery mechanism is iterative, humandriven workflows that combine manual data discovery, ad hoc integration, and consultation with stakeholders. For example, Use Case:
Purchase Order requires analysts to manually locate and combine purchase order data and invoicing or execution records, and to consult cross-functional teams to interpret observed patterns. These workflows are communication-heavy and may take weeks.
The Opportunity. LLMs demonstrate strong capabilities in interpreting natural-language queries and reasoning over humanspecified instructions [24, 52, 55]. These capabilities align with the need to interpret I + accurately and reason about actions that would help fulfill it, creating an opportunity to rethink how LLMs could reduce the time and effort to address the problem.

2.3

Solution Requirements

We outline the requirements of a technical intervention Σ for addressing the problem, informed by user feedback at our university.
• R1: Alignment Fidelity Since I ∗ is unobservable, Σ must ensure high-fidelity alignment with I𝑖 + : 𝐷𝑖 must faithfully follow the constraints articulated in I𝑖 + based on the available data in C, allowing the user to isolate data limitations from system failures.
• R2: Interpretability To ensure that the user can verify (and hence trust) D, Σ should justify the derivation of D from C.
• R3: Refinability Since I ∗ may only be partially specified through I + , Σ should support mechanisms that allow the user to refine I + and communicate clarifications or constraints.
• R4: Efficiency Σ must reduce the time and effort to fulfill an information need compared to current workflows. For example, manual discovery and consultation in procurement scenarios may take weeks and involve multiple stakeholders, whereas Σ
should produce D within minutes to hours.
These requirements delineate the space of feasible solutions; approaches that fail any of them are insufficient for the problem as framed. For example:
• Traditional data catalogs [21] improve visibility and documentation, but leave users to manually identify tables and refine I + , failing R3 and R4.
• Static dashboards or reporting tools (e.g., Tableau [47]) require users to manually rebuild queries or views to refine I + , failing
R3 and often R4.
Collectively, these examples illustrate that an acceptable technical intervention must not only surface the required data correctly and efficiently, but also support interpretability and refinement of partially-specified information needs.

3

RELATIONAL REIFICATION MECHANISM

When an LLM is tasked with directly answering I + , it must implicitly adopt an interpretation of the need and attempt to fulfill it by operating on the underlying tables, entangling the interpretation of I + with the construction of D in a single process. The user has no structured handle on the system’s interpretation beyond the final natural-language answer and any associated reasoning traces.
Without an explicit representation of this interpretation, the system lacks a grounded notion of progress, and the user cannot easily understand D well enough to help the system make progress.
For example, in Use Case: Purchase Order, the LLM might answer simply "yes" or "no", possibly with an explanation of its reasoning. A user seeking to verify correctness may ask follow-up questions, but each response may provide a new explanation, leaving the user uncertain whether to trust it or whether the LLM is hallucinating. Because these explanations are unstructured and inconsistent across queries, the user cannot rely on a uniform method to cross-check answers, making verification ad hoc and error-prone.
We address this by reifying I + as an explicit relational model.
Instead of directly synthesizing D, the system first commits to a target model (T , 𝑆), where T = {𝑇1, . . . ,𝑇𝑘 } is a set of derived views and 𝑆 is an executable transformation over T (e.g., Python code).
Each 𝑇𝑖 is constructed from C via operations, e.g., joins and filters.

Modeling T as a set of views (instead of a single view) reflects that complex information needs may decompose into multiple semantically distinct relations; for example, in Use Case: Purchase
Order, different operationalizations of “hazardous materials” and different signals of execution friction naturally give rise to multiple views. Collapsing these into one view can obscure modeling assumptions and hinder reuse and inspection. This enforces a separation between interpretation and generation: the system must make its interpretation of I + explicit as (T , 𝑆) before producing D, decomposing the task into (i) defining (T , 𝑆) and (ii) materializing
T , after which executing 𝑆 yields D.
Benefits for the User. Reification externalizes the system’s interpretation into concrete, inspectable representations, fulfilling R3.
By examining the schemas and sample data of 𝑇 , the user can detect logical misalignment (e.g., missing constraints), allowing them to provide feedback to the system to refine (T , 𝑆).
Benefits for the System. (T , 𝑆) operationalizes I + , enabling the system to track progress by checking whether 𝑇 has been materialized and whether D has been produced by executing 𝑆. It also constrains the system’s operation space to actions that support defining (T , 𝑆), materializing 𝑇 , and executing 𝑆.
Interpretability and Provenance. While (T , 𝑆) specifies how D
is constructed, users may want to verify the correctness of 𝑇 and its derivation. To satisfy R2, the system records the derivation of 𝑇
from C as a directed acyclic graph (DAG), where nodes represent relational transformations and edges encode data dependencies.
This process provides users end-to-end visibility and control over
D’s construction by ensuring the final output is generated by a deterministic script derived directly from the provenance graph.
Example: Refining the Relational Model. Consider Use Case:
Purchase Order, where the analyst wants to compare post-purchase friction rates for hazardous or radioactive products. The system may initially reify I + into a single view, 𝑇ℎ𝑎𝑧𝑎𝑟𝑑𝑜𝑢𝑠_𝑟𝑎𝑑𝑖𝑜𝑎𝑐𝑡𝑖𝑣𝑒 , that combines all hazardous and radioactive products.
Upon inspecting 𝑇ℎ𝑎𝑧𝑎𝑟𝑑𝑜𝑢𝑠_𝑟𝑎𝑑𝑖𝑜𝑎𝑐𝑡𝑖𝑣𝑒 , the analyst notices that radioactive products are a subset of hazardous products and require separate handling. To enable a more granular analysis, the user refines I + . The system updates T to consist of three views: (1)
radioactive products, (2) hazardous but non-radioactive products, and (3) non-hazardous products. This refinement ensures that the transformation 𝑆 produces a final D that correctly accounts for subpopulation differences, bringing the I + closer to I ∗ .

4

DESIGNING A SYSTEM TO SUPPORT (T , 𝑆)

This section presents the design of Pneuma-Seeker, an LLM-based system that implements relational reification by reifying I + as
(T , 𝑆). We discuss two complementary context-management strategies: macro-level, which shapes the overall system architecture, and micro-level, which governs LLM interactions with individual tables.

4.1

Macro Context Management

A central design principle is macro context management: decomposing the workflow into multiple components with scoped responsibilities, thereby reducing the context each component must process.
This helps mitigate LLM context window limitations, including the

CoreService
Conductor

Knowledge
Store

State

Target Tables (T)

External
DB

Retriever

Script (S)

Web Search
Web Crawler

Action Set

Table
Enumerator

Python
Executor

...

DBService

Dataset Index

Workspace w1

Materializer

OpenAI
LLM

ProvenanceGraph

...

Azure
OpenAI
LLM

Dataset d1 ...
Workspace ...
w2

LMService
OpenAI
Embedding Model
Azure OpenAI
Embedding Model

...

...

Context
Extraction

Parquet p1

Pneuma-Retriever

User

Relational
Operators

CSV c1

Figure 1: Architecture of Pneuma-Seeker (→: utilized by)

degradation in performance observed with longer contexts [12, 31].
Pneuma-Seeker realizes this principle through a multi-component design (Figure 1) with three tasks: (1) retrieval, (2) materialization, and (3) orchestration.
Rather than processing all tables in C, Pneuma-Seeker retrieves a relevant subset by performing tabular data discovery. We encapsulate the retrieval logic in a component called Retriever
(Section 5.3). Other parts of the system interact with Retriever only through its interface, which returns relevant tables. This modularization enables extensions (e.g., non-tabular retrieval).
Materializing T from its specification—rather than the original natural-language phrasing of I + —constitutes the second main task: given a specification of T (schemas and column descriptions), materialize T using tables returned by Retriever. Isolating this functionality enables targeted optimization, e.g., introducing dedicated relational operators to reduce LLM generation errors. This functionality is implemented in Materializer (Section 5.2).
The overall workflow is coordinated by an orchestrator, which implements the relational reification mechanism to transform I +
into (T , 𝑆). The orchestrator retrieves relevant tables via Retriever to constrain feasible (T , 𝑆), invokes Materializer to materialize
T , and executes 𝑆 to produce D. It is decoupled from the internal implementations of Retriever and Materializer, interacting with them only through their interfaces, and focuses on planning, progress tracking, and user interaction. We define Conductor as the component fulfilling this role (Section 5.1).
These three components define the core service, supported by two additional services: (1) DBService (Section 5.4) for persistent storage of (T , 𝑆), ProvenanceGraph, and query execution; and
(2) LMService (Section 5.5) for access to language models.

4.2

protocol between the LLM and DBService: the LLM is prompted to externalize uncertainty about table contents and may create
Python scripts, equipped with APIs to execute SQL queries (part of DBService), to issue queries over retrieved tables. Rather than simply supplying table content (i.e., schemas and sample rows), the system actively scaffolds the LLM to acquire context on demand by executing targeted probes over tables. This mechanism enables the LLM to retrieve precise values, distributions, and structural properties that are directly relevant to I + .
For instance, in Use Case: Purchase Order, the retrieved PO table lacks a unified “hazardous” flag, instead containing columns such as radioactive and toxin, which may be related. To infer their relationships, the LLM can invoke a script to, e.g., observe the value distributions of these columns. This exploratory probing allows the system to determine exactly which combinations of records may constitute the user’s high-level concept before deciding how that subset should be incorporated into T .
Macro context management reduces the contextual load of each component by decomposing the workflow into smaller tasks, while micro context management equips the LLM with an explicit action for fine-grained evidence acquisition within retrieved tables. This explicit support for targeted context extraction directly supports
R1; Section 6.2.2 shows that this action helps produce correct D.
As another example, consider the query from the Legal dataset in
KramaBench [25]: “What is the proportion (round to 3 decimal places)
of fraud reporters who lost between $1–$500 in 2024?”. The retriever returns a table whose content interleaves coarse-grained and finegrained bins; for instance, it includes rows such as “$1–$1,000:
624,110” and later “$1–$100: 243,174” and “$101–$200: 114,336.”
When only a small set of sample rows is exposed, the LLM assumes that the table provides only coarse-grained aggregates and fails to detect the finer-grained breakdown required to answer the query.
Although the LLM is in principle capable of generating scripts to probe such irregularities, in practice it does not reliably externalize or act on this uncertainty when context extraction is not an explicit, supported action. With micro context management,
Pneuma-Seeker exposes context extraction as a deliberate action mediated by DBService, changing the interaction protocol: the
LLM can surface its uncertainty, test it by issuing scripts to probe the table structure (e.g., checking for overlapping bins or finergrained ranges), surface the relevant rows, and construct (T , 𝑆)
accordingly. This enables correct materialization of T and, in turn, the correct D, despite the table’s non-ideal organization.

Micro Context Management

Macro-level decomposition alone does not suffice for effective
LLM–table interaction. Even after narrowing the search space via
Retriever, fully including retrieved tables in context is often infeasible. In Use Case: Purchase Order, the purchase order table has
117 columns and 5 million rows; with o3-2025-04-16, this would require at least 2.1B tokens—far beyond the 200,000-token limit.
We often expect larger and multiple tables to be retrieved at once.
Conversely, sampling only a few rows can omit critical evidence, leading to incomplete reasoning or errors in constructing (T , 𝑆).
To address this limitation, Pneuma-Seeker introduces a complementary micro-level context management strategy. Concretely, micro context management is realized as a structured interaction

5

THE PNEUMA-SEEKER SYSTEM

This section details how each main component of Pneuma-Seeker performs its task to support the overall workflow: reifying I + as
(T , 𝑆) and executing 𝑆 to produce D.
Dynamic Planning. Rather than following fixed, pre-defined pipelines, Conductor and Materializer construct plans dynamically. Our initial prototypes relied on static pipelines, but we found this approach required costly re-engineering as new use cases necessitated revisiting control-flow decisions (e.g., determining which steps are optional, when to skip or restart actions, and how to integrate new tools). In contrast, dynamic planning allows developers to add new capabilities directly to the components’ action

spaces, which the planners leverage immediately without altering the overall pipeline, significantly improving extensibility.
Both components operate within bounded planning loops driven by model inference. At each iteration, the planner selects a sequence of actions (i.e., agentic tool calls, cf. ReAct [55]), beginning with
Situational Analysis to determine subsequent actions in the same iteration. We permit multi-action sequences to balance efficiency and flexibility: a single action per iteration incurs unnecessary model calls and input-token overhead, whereas selecting all actions at once is too rigid because later actions often depend on earlier outcomes (e.g., Conductor defines T only after gathering evidence from retrieved tables).

5.1

Conductor

Conductor orchestrates the overall workflow of the system. Users interact with it through a chat interface. Given a user query, Conductor loops for up to 𝑐 iterations (default 𝑐 = 10). At each iteration, an LLM observes the current state (e.g., (T , 𝑆) and retrieved tables) and outputs a JSON object specifying a sequence of actions to execute. Its action space is as follows:
• Situational Analysis. Executed at the beginning of each iteration to analyze the current model (T , 𝑆) and prior actions, and to decide which actions to take in the iteration.
• Retriever-related Actions. Retrieve the top-𝑘 relevant tables from Pneuma-Retriever, or search and crawl the web.
• Context Extraction. Enable micro context management by issuing scripts to probe retrieved tables for relevant evidence (e.g., testing whether specific values exist in certain columns).
• (T , 𝑆) Manipulation. Update T , 𝑆, or both, such as adding target views, refining columns, or adjusting transformations.
• Executor. Execute 𝑆 over materialized T to produce D.
• Materializer. Invoke Materializer to materialize T , optionally including textual guidance (e.g., to propagate user feedback).
• User-Facing Communication. Present results to the user or ask clarifying questions to further refine (T , 𝑆).
Conductor in Action. Actions are represented as JSON objects containing an action type and corresponding arguments. For example, in the Use Case: Purchase Order, the first iteration may produce:
{ " plan " : [
{ " action " : " Situational Analysis " ,
" a r g s " : { " message " : " No t a b l e s y e t . . . " } } ,
{ " action " : " Table Retrieve " ,
" args " : { " queries " : [ " hazardous . . . " ] } } ] }
In the subsequent iteration, Conductor observes the newly retrieved tables within its context and may combine Situational
Analysis with Context Extraction to probe them for specific evidence related to the query. The loop continues until Conductor selects User-Facing Communication or reaches the iteration limit 𝑐, forcing a final user-facing response.

5.2

Materializer

Materializer constructs the views specified in T , taking into account 𝑆 (if defined by Conductor) and any textual guidance from Conductor. It receives retrieved data as shared context, but

may invoke Situational Analysis, Retriever-related Actions, and Context Extraction to reason about intermediate results and acquire additional evidence. Similar to Conductor, Materializer iteratively selects action sequences for up to 𝑚 iterations (default
𝑚 = 10), producing intermediate tables.
Table-Construction Actions. Beyond these shared actions, Materializer has a set of table-construction actions. Materializer favors structured operators over free-form query or code generation, reducing error rates and token usage. Nevertheless, free-form generation is retained as a fallback for transformations that cannot be expressed using structured operators.
• Relational Operators. Common relational operators (e.g., join and projection) are implemented as dedicated actions. The LLM
specifies key parameters (e.g., table IDs and join keys), and the application code generates and executes the corresponding queries.
• Semantic Operators. We implement several semantic operators.
Semantic join matches tuples between two tables based on semantic relatedness between textual attributes rather than exact equality, following prior work on embedding-based dataset linking (e.g., [10]) and recent declarative formulations of semantic operators such as LOTUS [40]. We compute embeddings over selected columns and rank candidates using cosine similarity and lightweight syntactic similarity, retaining the top-𝑝 matches per left tuple (default 𝑝 = 1). Semantic column generation, inspired by Palimpzest [30], adds a column to an intermediate table whose values are synthesized by the LLM, conditioned on selected existing columns.
• Query Executor. Beyond composing relational and semantic operators, Materializer may define custom SQL queries and persist their results as intermediate tables, increasing flexibility without the brittleness of unrestricted code generation.
• Python Executor. This allows Materializer to generate arbitrary Python code over retrieved and intermediate tables, with access to the DB service APIs (e.g., execute_query()). While maximally expressive, it is used only when preceding operators cannot express the required transformation.
ProvenanceGraph records all transformations performed by
Materializer. Nodes correspond to transformations (e.g., joins), while edges encode dependencies between transformations and relations. The graph is constructed incrementally during execution, enabling Pneuma-Seeker to expose the derivation of T via the sidebar (Section 6.1). This is precisely the justification for 𝐷𝑖 that our formalization identifies as Pneuma-Seeker’s primary lever for supporting the user’s own convergence toward I ∗ . Topologically sorting the graph reconstructs T as an executable script; appending
𝑆 yields the end-to-end script that reproduces the generated answer, supporting R2 and R3.

5.3

Retriever

Retriever provides a unified interface for all retrieval mechanisms used by Pneuma-Seeker. Both Conductor and Materializer may issue up to 𝑞 retrieval queries in parallel (default 𝑞 = 3), each returning top-𝑘 items.
The main mechanism is Pneuma-Retriever, which retrieves the top-𝑘 tables per query (default 𝑘 = 10). It builds on Pneuma [3],

which summarizes table schemas and sample rows for table discovery. Since schema-level summaries miss cell-level signals, we add a content-aware stage inspired by Octopus [28], extracting likely entities from the query, matching them via case-insensitive regex against tables, and scoring matches using damped term frequency log(1 + tf) normalized per keyword. The final score combines content-aware and Pneuma retrieval relevance scores.
Pneuma-Retriever also supports enumeration queries over table identifiers via regular expressions (e.g., water_body_testing_
\d{4} after retrieving water_body_testing_2020), enabling retrieval when many tables share schemas but differ in attributes such as year or location.
Other mechanisms include: Knowledge Store, which indexes reusable knowledge from user interactions [32, 39, 48]; Web Search, which retrieves data via search APIs; and Web Crawl, which fetches user-specified pages when permitted.
Extensibility. Retriever exposes a uniform interface over heterogeneous sources: each result is an AbstractDocument with an
ID, content, and key–value metadata, where content may represent different modalities (e.g., tables or web documents), enabling
Conductor and Materializer to operate uniformly across sources.
New retrievers can be added by implementing this interface, integrating additional data sources and indexing strategies.
In Use Case: Purchase Order, the analyst requests additional information for reporting beyond the retrieved tables. In an actual execution of Pneuma-Seeker, Conductor invokes Materializer to augment the target view with non-tabular attributes through
Web Search and Web Crawl (Figure 2).
• Web Search. After joining the PO and item tables on radioactivematerial items, Materializer performs a web search for each supplier company and extracts the INDUSTRY_SECTOR attribute.
• Web Crawl. The user provides a URL template for a radionuclide safety data sheet repository keyed by RADIONUCLIDE (e.g., 3H).
Materializer crawls one page for each distinct radionuclide and extracts the RADIATION_TYPE attribute.
In both operations, ProvenanceGraph adds a node for each retrieved document upon retrieval. When Materializer integrates the documents through generated Python code, it annotates the code with the documents’ IDs; ProvenanceGraph detects this annotation and adds the corresponding edge, preserving provenance.
PO ⊲⊳ ITEM (augmented)
SUPP_ID SUPP_NAME INDUSTRY_SECTOR RADIONUCLIDE RADIATION_TYPE IS_RADIOACTIVE IS_OVERINVOICED

Figure 2: Web-Augmented Target View

5.4

DBService

DBService provides query execution and persistent storage for system state, including (T , 𝑆). To support concurrency, each user–
chat pair is assigned an isolated workspace database stored as a separate workspace.db file using DuckDB [42], which localizes failures and transactions. Datasets may be ingested as DuckDB
databases (e.g., procurement.db) or accessed in external formats such as CSV or Parquet, or external systems such as PostgreSQL

via connectors. DBService uses DuckDB’s direct file querying and
ATTACH mechanism to unify access across local and external sources.

5.5

LMService

LMService provides a unified interface to LLMs and embedding models, allowing new providers to be integrated without changes to the core system. We currently support OpenAI and Azure OpenAI
endpoints, with Azure used for internal university deployments due to data governance requirements.

6

EVALUATION

In this section, we answer the following questions:
• RQ1 (User Study): Does Pneuma-Seeker’s answer justification help users complete tasks?
• RQ2 (Answer Quality): Does Pneuma-Seeker fulfill users’ information needs?
• RQ3 (Cost and Scalability): What cost–quality and scalability trade-offs does Pneuma-Seeker exhibit?
• RQ4 (Microbenchmarks): How do key design choices in PneumaSeeker shape the trade-offs between answer quality, runtime, memory usage, and retrieval coverage?
Datasets. We evaluate Pneuma-Seeker on KramaBench [25], a benchmark that includes six real-world datasets (Table 1) spanning scientific, legal, and public domains. We filter the questions to retain only those that require reasoning over tabular data. We additionally preprocess several datasets (e.g., extracting CSVs from XLSX files)
to standardize the input format across systems. All questions and processed tables are publicly available in our repository.1
Table 1: Dataset Characteristics and Benchmark Statistics
Name
Archeology
Astronomy∗
Biomedical
Environment∗
Legal∗
Wildfire∗

#Tables #Cols‡ #Rows‡ Size (MB) #Q
5
1,438
24
36
144
20

16
29
43
10
3
14

11,289
1,053
16,297
9,199
25
432,277

5.90
313.87
76.96
29.45
0.20
1040

12
4
9
17
28
18

∗ Filtered to only questions that require tabular data. ‡ Average per table.

We choose KramaBench because its questions are non-trivial and predominantly require multi-table integration. In contrast, many existing benchmarks (e.g., MMTU [53], RADAR [17], DataBench [37])
focus on single-table question answering with simpler queries;
MultiTabQA [38] supports multi-table integration but with simple questions, and TAG [6], which builds on BIRD [27], contains questions with incorrect annotations [22]. These issues make such benchmarks less suitable for evaluating precise, data-grounded question answering over tables.
Baselines. We compare Pneuma-Seeker against these baselines:
• DS-Guru [25]: A system that instructs an LLM to decompose a question into subtasks, reason through each step, and synthesize
Python code implementing the plan.
1 https://github.com/TheDataStation/pneuma-seeker

• smolagents [13]: An agentic baseline implemented using smolagents, an open-source agent framework supporting tool calling and code-based reasoning. The agent is provided with tools to execute Python code, list available tables, and inspect schemas and sample rows (up to 5 rows per table).
System Setup and Configurations. Experiments run on a MacBook Air (M4) with 16 GB of RAM and Python 3.12.12, using OpenAI’s o3-2025-04-16 and text-embedding-3-small as the LLM
and embedding models. To ensure a fair comparison with the baselines, we disable Web Search, Web Crawl, and Knowledge Store.

6.1

RQ1: User Study

We recruit 36 participants who interact with data either daily or several times per week. Most participants (36.11%) are PhD students across various fields (e.g., neuroscience and anthropology), followed by data analysts (11.11%), engineers (5.56%), and others.
We evaluate whether answer justification exposed by PneumaSeeker supports users in inspecting results during iterative discovery. This targets the proxies for convergence: since I ∗ is unobservable to the system, we look for observable signals that answer justification helps users close the gap themselves. Concretely, the system provides a sidebar beside the main chat interface (Figure 3), which exposes information about the intermediate steps that the system takes to produce an answer. We conduct a withinsubjects study, where each participant completes two tasks, one under each condition: Chat-Sidebar (default of Pneuma-Seeker)
and Chat-Only. We counterbalance both condition order and task assignment to mitigate potential order effects.

• “Using the system, determine how many beaches had a higher bacterial exceedance rate in 2013 than in 2012, considering only beaches with samples collected in both years.” (role: data analyst at the Massachusetts Department of Public Health (DPH))
Participants complete assigned tasks using the system, while also being allowed to inspect and explore artifacts beyond directly prompting it. For instance, they may download relevant tables from the sidebar and independently cross-check results. After exploration, participants rate: (1) their confidence in the correctness of the system’s answer and (2) their readiness to act on the result (e.g., to refine information needs or make decisions). These measures serve as observable proxies for convergence toward I ∗ under our formalization, since the distance to I ∗ itself cannot be measured directly. Responses are collected using a 5-point Likert scale [29].
Participants are also asked to think aloud during interactions with both conditions and describe how they inspect answers. Based on these observations, we classify inspection effort as: (1) low (relying only on chat responses without, e.g., checking calculations); (2)
medium (inspecting answer justification through the sidebar); or
(3) high (performing manual audits such as downloading the tables and performing calculations).
Table 2 shows Chat-Sidebar achieves higher median confidence and readiness-to-act scores than Chat-Only. We further analyze the intra-condition rank correlation (Kendall’s 𝜏𝑏 ) between inspection effort and these ratings. In Chat-Only, both correlations are significant and positive (e.g., 𝜏𝑏 = .338, 𝑝 = .0292 for confidence): users need increasingly manual inspection — recalculating statistics, cross-checking source tables — before feeling confident or ready to act. In Chat-Sidebar, both correlations become non-significant
(e.g., 𝜏𝑏 = .063, 𝑝 = .6621 for confidence): confidence and readiness to act no longer depend on how much users manually verify. This indicates that exposing answer justification through the sidebar substitutes for manual verification as a basis for confidence and readiness to act — consistent with the role of justification in our formalization, where Ψ cannot optimize for I ∗ directly and instead supports the user’s own convergence by justifying 𝐷𝑖 , reducing the verification effort O must supply.
Table 2: Ratings (5-Point Likert) and 𝜏𝑏 vs. Inspection Effort

Figure 3: Pneuma-Seeker’s Interface (Chat-Sidebar)

Chat-Only
Chat-Sidebar
Median 𝜏𝑏
𝑝
Median 𝜏𝑏
𝑝
Confidence
3.50 0.338 0.0292 4.00 0.063 0.6621
Readiness to Act 2.50 0.398 0.0095 3.00 0.082 0.5687

Each participant is assigned a role and provided with relevant contextual information, then asked to answer two questions from
KramaBench. One question is answered using Chat-Sidebar, and the other using Chat-Only, with tasks and conditions counterbalanced across participants. The order of conditions is also counterbalanced, such that half of the participants use the Chat-Sidebar first, while the other half begin with Chat-Only.
• “Using the system, determine whether the number of ’frauds and other data’ records for 2024 is consistent across datasets for the
Metropolitan Area of Miami–Fort Lauderdale–West Palm Beach.”
(role: data analyst at the Federal Trade Commission (FTC))

To further contextualize the size of this improvement, we highlight qualitative feedback from participants:
• “So I could kind of like reason through the steps... I can see the source data... I can now see how this leads to the final output... So I could kind of like reason through... so now I’m super confident.”
• “...this interface makes it really easy for me to follow what’s going on... for Miami-Fort Lauderdale, it’s 171 for number reports, and then it’s 21,793 for reports per 100K, and that’s wrong...”
• “The code explanation is very helpful in terms of you can kind

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

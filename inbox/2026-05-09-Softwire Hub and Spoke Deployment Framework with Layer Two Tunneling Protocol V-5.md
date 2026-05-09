---
source: "https://arxiv.org/abs/1904.07960v1"
title: "Softwire Hub and Spoke Deployment Framework with Layer Two Tunneling Protocol Version 2 (L2TPv2)"
author: "Bill Storer, Carlos Pignataro, Maria Alice Dos Santos, Bruno Stévant, Laurent Toutain, Jean-François Tremblay"
year: "2019"
publication: "arXiv preprint / cs.NI"
download: "https://arxiv.org/pdf/1904.07960v1"
pdf: "https://arxiv.org/pdf/1904.07960v1"
captured_at: "2026-05-09T12:53:55Z"
updated_at: "2026-05-09T12:53:55Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche Thus Spoke Zarathustra"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# Softwire Hub and Spoke Deployment Framework with Layer Two Tunneling Protocol Version 2 (L2TPv2)

- 著者: Bill Storer, Carlos Pignataro, Maria Alice Dos Santos, Bruno Stévant, Laurent Toutain, Jean-François Tremblay
- 年: 2019
- 掲載情報: arXiv preprint / cs.NI
- 情報源: [arxiv](https://arxiv.org/abs/1904.07960v1)
- ダウンロード: https://arxiv.org/pdf/1904.07960v1
- PDF: https://arxiv.org/pdf/1904.07960v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

This document describes the framework of the Softwire ''Hub and Spoke'' solution with the Layer Two Tunneling Protocol version 2 (L2TPv2). The implementation details specified in this document should be followed to achieve interoperability among different vendor implementations.

## PDF Text

Network Working Group B. StorerRequest for Comments: 5571 C. Pignataro, Ed.Category: Standards Track M. Dos Santos Cisco Systems B. Stevant, Ed. L. Toutain TELECOM Bretagne J. Tremblay Videotron Ltd. June 2009 Softwire Hub and Spoke Deployment Framework with Layer Two Tunneling Protocol Version 2 (L2TPv2)Status of This Memo This document specifies an Internet standards track protocol for the Internet community, and requests discussion and suggestions for improvements. Please refer to the current edition of the "Internet Official Protocol Standards" (STD 1) for the standardization state and status of this protocol. Distribution of this memo is unlimited.Copyright Notice Copyright (c) 2009 IETF Trust and the persons identified as the document authors. All rights reserved. This document is subject to
BCP 78
and the IETF Trust's Legal Provisions Relating to IETF Documents in effect on the date of publication of this document (
http://trustee.ietf.org/license-info
). Please review these documents carefully, as they describe your rights and restrictions with respect to this document. This document may contain material from IETF Documents or IETF Contributions published or made publicly available before November 10, 2008. The person(s) controlling the copyright in some of this material may not have granted the IETF Trust the right to allow modifications of such material outside the IETF Standards Process. Without obtaining an adequate license from the person(s) controlling the copyright in such materials, this document may not be modified outside the IETF Standards Process, and derivative works of it may not be created outside the IETF Standards Process, except to format it for publication as an RFC or to translate it into languages other than English.Storer, et al. Standards Track [Page 1]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009Abstract This document describes the framework of the Softwire "Hub and Spoke" solution with the Layer Two Tunneling Protocol version 2 (L2TPv2). The implementation details specified in this document should be followed to achieve interoperability among different vendor implementations.Table of Contents
1
. Introduction ....................................................
4

1.1
. Abbreviations ..............................................
5

1.2
. Requirements Language ......................................
5

1.3
. Considerations .............................................
6

2
. Applicability of L2TPv2 for Softwire Requirements ...............
6

2.1
. Traditional Network Address Translation (NAT and NAPT) .....
6

2.2
. Scalability ................................................
7

2.3
. Routing ....................................................
7

2.4
. Multicast ..................................................
7

2.5
. Authentication, Authorization, and Accounting (AAA) ........
7

2.6
. Privacy, Integrity, and Replay Protection ..................
7

2.7
. Operations and Management ..................................
8

2.8
. Encapsulations .............................................
8

3
. Deployment Scenarios ............................................
8

3.1
. IPv6-over-IPv4 Softwires with L2TPv2 .......................
9

3.1.1
. Host CPE as Softwire Initiator ......................
9

3.1.2
. Router CPE as Softwire Initiator ...................
10

3.1.3
. Host behind CPE as Softwire Initiator ..............
11

3.1.4
. Router behind CPE as Softwire Initiator ............
12

3.2
. IPv4-over-IPv6 Softwires with L2TPv2 ......................
14

3.2.1
. Host CPE as Softwire Initiator .....................
14

3.2.2
. Router CPE as Softwire Initiator ...................
15

3.2.3
. Host behind CPE as Softwire Initiator ..............
16

3.2.4
. Router behind CPE as Softwire Initiator ............
16

4
. References to Standardization Documents ........................
17

4.1
. L2TPv2 ....................................................
18

4.2
. Securing the Softwire Transport ...........................
18

4.3
. Authentication, Authorization, and Accounting .............
18

4.4
. MIB .......................................................
18

4.5
. Softwire Payload Related ..................................
19

4.5.1
. For IPv6 Payloads ..................................
19

4.5.2
. For IPv4 Payloads ..................................
19

5
. Softwire Establishment .........................................
20

5.1
. L2TPv2 Tunnel Setup .......................................
22

5.1.1
. Tunnel Establishment ...............................
22

5.1.1.1
. AVPs Required for Softwires ...............
25

5.1.1.2
. AVPs Optional for Softwires ...............
25

5.1.1.3
. AVPs Not Relevant for Softwires ...........
26
Storer, et al. Standards Track [Page 2]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
5.1.2
. Tunnel Maintenance .................................
26

5.1.3
. Tunnel Teardown ....................................
27

5.1.4
. Additional L2TPv2 Considerations ...................
27

5.2
. PPP Connection ............................................
27

5.2.1
. MTU ................................................
27

5.2.2
. LCP ................................................
27

5.2.3
. Authentication .....................................
28

5.2.4
. IPCP ...............................................
28

5.2.4.1
. IPV6CP ....................................
28

5.2.4.2
. IPv4CP ....................................
28

5.3
. Global IPv6 Address Assignment to Endpoints ...............
28

5.4
. DHCP ......................................................
29

5.4.1
. DHCPv6 .............................................
29

5.4.2
. DHCPv4 .............................................
29

6
. Considerations about the Address Provisioning Model ............
30

6.1
. Softwire Endpoints' Addresses .............................
30

6.1.1
. IPv6 ...............................................
30

6.1.2
. IPv4 ...............................................
31

6.2
. Delegated Prefixes ........................................
31

6.2.1
. IPv6 Prefixes ......................................
31

6.2.2
. IPv4 Prefixes ......................................
31

6.3
. Possible Address Provisioning Scenarios ...................
31

6.3.1
. Scenarios for IPv6 .................................
32

6.3.2
. Scenarios for IPv4 .................................
32

7
. Considerations about Address Stability .........................
32

8
. Considerations about RADIUS Integration ........................
33

8.1
. Softwire Endpoints ........................................
33

8.1.1
. IPv6 Softwires .....................................
33

8.1.2
. IPv4 Softwires .....................................
33

8.2
. Delegated Prefixes ........................................
34

8.2.1
. IPv6 Prefixes ......................................
34

8.2.2
. IPv4 Prefixes ......................................
34

9
. Considerations for Maintenance and Statistics ..................
34

9.1
. RADIUS Accounting .........................................
35

9.2
. MIBs ......................................................
35

10
. Security Considerations .......................................
35

11
. Acknowledgements ..............................................
36

12
. References ....................................................
37

12.1
. Normative References .....................................
37

12.2
. Informative References ...................................
38
Storer, et al. Standards Track [Page 3]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
1
. Introduction The Softwires Working Group has selected Layer Two Tunneling Protocol version 2 (L2TPv2) as the phase 1 protocol to be deployed in the Softwire "Hub and Spoke" solution space. This document describes the framework for the L2TPv2 "Hub and Spoke" solution, and the implementation details specified in this document should be followed to achieve interoperability among different vendor implementations. In the "Hub and Spoke" solution space, a Softwire is established to provide the home network with IPv4 connectivity across an IPv6-only access network, or IPv6 connectivity across an IPv4-only access network. When L2TPv2 is used in the Softwire context, the voluntary tunneling model applies. The Softwire Initiator (SI) at the home network takes the role of the L2TP Access Concentrator (LAC) client (initiating both the L2TP tunnel/session and the PPP link) while the Softwire Concentrator (SC) at the ISP takes the role of the L2TP Network Server (LNS). The terms voluntary tunneling and compulsory tunneling are defined in
Section 1.1 of [RFC3193]
. Since the L2TPv2 compulsory tunneling model does not apply to Softwires, it SHOULD NOT be requested or honored. This document identifies all the voluntary tunneling related L2TPv2 attributes that apply to Softwires and specifies the handling mechanism for such attributes in order to avoid ambiguities in implementations. This document also identifies the set of L2TPv2 attributes specific to the compulsory tunneling model that does not apply to Softwires and specifies the mechanism to ignore or nullify their effect within the Softwire context. The SI and SC MUST follow the L2TPv2 operations described in [
RFC2661
] when performing Softwire establishment, teardown, and Operations, Administration, and Management (OAM). With L2TPv2, a Softwire consists of an L2TPv2 Control Connection (also referred to as Control Channel), a single L2TPv2 Session, and the PPP link negotiated over the Session. To establish the Softwire, the SI first initiates an L2TPv2 Control Channel to the SC, which accepts the request and terminates the Control Channel. L2TPv2 supports an optional mutual Control Channel authentication that allows both SI and SC to validate each other's identity at the initial phase of hand-shaking before proceeding with Control Channel establishment. After the L2TPv2 Control Channel is established between the SI and SC, the SI initiates an L2TPv2 Session to the SC. Then the PPP/IP link is negotiated over the L2TPv2 Session between the SI and SC. After the PPP/IP link is established, it acts as the Softwire between the SI and SC for tunneling IP traffic of one Address Family (AF) across the access network of another Address Family.Storer, et al. Standards Track [Page 4]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 During the life of the Softwire, both SI and SC send L2TPv2 keepalive HELLO messages to monitor the health of the Softwire and the peer L2TP Control Connection Endpoint (LCCE), and to potentially refresh the NAT/NAPT (Network Address Translation / Network Address Port Translation) entry at the CPE or at the other end of the access link. Optionally, Link Control Protocol (LCP) ECHO messages can be used as keepalives for the same purposes. In the event of keepalive timeout or administrative shutdown of the Softwire, either the SI or the SC MAY tear down the Softwire by tearing down the L2TPv2 Control Channel and Session as specified in [
RFC2661
].
1.1
. Abbreviations AF Address Family, IPv4 or IPv6. CPE Customer Premises Equipment. LCCE L2TP Control Connection Endpoint, an L2TP node that exists at either end of an L2TP Control Connection. (See [
RFC3931
].) LNS L2TP Network Server, a node that acts as one side of an L2TP tunnel (Control Connection) endpoint. The LNS is the logical termination point of a PPP session that is being tunneled from the remote system by the peer LCCE. (See [
RFC2661
].) SC Softwire Concentrator, the node terminating the Softwire in the service provider network. (See [
RFC4925
].) SI Softwire Initiator, the node initiating the Softwire within the customer network. (See [
RFC4925
].) SPH Softwire Payload Header, the IP headers being carried within a Softwire. (See [
RFC4925
].) STH Softwire Transport Header, the outermost IP header of a Softwire. (See [
RFC4925
].) SW Softwire, a shared-state "tunnel" created between the SC and SI. (See [
RFC4925
].)
1.2
. Requirements Language The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [
RFC2119
].Storer, et al. Standards Track [Page 5]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
1.3
. Considerations Some sections of this document contain considerations that are not required for interoperability and correct operation of Softwire implementations. These sections' titles are marked beginning with "Considerations".
2
. Applicability of L2TPv2 for Softwire Requirements A list of Softwire "Hub and Spoke" requirements has been identified by the Softwire Problem Statement [
RFC4925
]. The following sub- sections describe how L2TPv2 fulfills each of them.
2.1
. Traditional Network Address Translation (NAT and NAPT) A "Hub and Spoke" Softwire must be able to traverse Network Address Translation (NAT) and Network Address Port Translation (NAPT, also referred to as Port Address Translation or PAT) devices [
RFC3022
] in case the scenario in question involves a non-upgradable, preexisting IPv4 home gateway performing NAT/NAPT or some carrier equipment at the other end of the access link performing NAT/NAPT. The L2TPv2 Softwire (i.e., Control Channel and Session) is capable of NAT/NAPT traversal since L2TPv2 can run over UDP. Since L2TPv2 does not detect NAT/NAPT along the path, L2TPv2 does not offer the option of disabling UDP. The UDP encapsulation is present regardless of NAT/NAPT presence. Both NAT/NAPT "autodetect" and UDP "bypass" are optional requirements in
Section 2.3 of [RFC4925]
. As mentioned in
Section 8.1 of [RFC2661]
and
Section 4 of [RFC3193]
, an L2TP Start-Control-Connection-Reply (SCCRP) responder (SC) can chose a different IP address and/or UDP port than those from the initiator's Start-Control-Connection-Request (SCCRQ) (SI). This may or may not traverse a NAT/NAPT depending on the NAT/NAPT's Filtering Behavior (see
Section 5 of [RFC4787]
). Specifically, any IP address and port combination will work with Endpoint-Independent Filtering, but changing the IP address and port will not work through Address- Dependent or Address-and-Port-Dependent Filtering. Given this, responding from a different IP address and/or UDP port is NOT RECOMMENDED.Storer, et al. Standards Track [Page 6]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
2.2
. Scalability In the "Hub and Spoke" model, a carrier must be able to scale the solution to millions of Softwire Initiators by adding more hubs (i.e., Softwire Concentrators). L2TPv2 is a widely deployed protocol in broadband services, and its scalability has been proven in multiple large-scale IPv4 Virtual Private Network deployments, which scale up to millions of subscribers each.
2.3
. Routing There are no dynamic routing protocols between the SC and SI. A default route from the SI to the SC is used.
2.4
. Multicast Multicast protocols simply run transparently over L2TPv2 Softwires together with other regular IP traffic.
2.5
. Authentication, Authorization, and Accounting (AAA) L2TPv2 supports optional mutual Control Channel authentication and leverages the optional mutual PPP per-session authentication. L2TPv2 is well integrated with AAA solutions (such as RADIUS) for both authentication and authorization. Most L2TPv2 implementations available in the market support the logging of authentication and authorization events. L2TPv2 integration with RADIUS accounting (RADIUS Accounting extension for tunnel [
RFC2867
]) allows the collection and reporting of L2TPv2 Softwire usage statistics.
2.6
. Privacy, Integrity, and Replay Protection Since L2TPv2 runs over IP/UDP in the Softwire context, IPsec Encapsulating Security Payload (ESP) can be used in conjunction to provide per-packet authentication, integrity, replay protection, and confidentiality for both L2TPv2 control and data traffic [
RFC3193
] and [
RFC3948
]. For Softwire deployments in which full payload security is not required, the L2TPv2 built-in Control Channel authentication and the inherited PPP authentication and PPP Encryption Control Protocol can be considered.Storer, et al. Standards Track [Page 7]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
2.7
. Operations and Management L2TPv2 supports an optional in-band keepalive mechanism that injects HELLO control messages after a specified period of time has elapsed since the last data or control message was received on a tunnel (see
Section 5.5 of [RFC2661]
). If the HELLO control message is not reliably delivered, then the Control Channel and its Session will be torn down. In the Softwire context, the L2TPv2 keepalive is used to monitor the connectivity status between the SI and SC and/or as a refresh mechanism for any NAT/NAPT translation entry along the access link. LCP ECHO offers a similar mechanism to monitor the connectivity status, as described in [
RFC1661
]. Softwire implementations SHOULD use L2TPv2 Hello keepalives, and in addition MAY use PPP LCP Echo messages to ensure Dead End Detection and/or to refresh NAT/NAPT translation entries. The combination of these two mechanisms can be used as an optimization. The L2TPv2 MIB [
RFC3371
] supports the complete suite of management operations such as configuration of Control Channel and Session, polling of Control Channel and Session status and their traffic statistics and notifications of Control Channel and Session UP/DOWN events.
2.8
. Encapsulations L2TPv2 supports the following encapsulations: o IPv6/PPP/L2TPv2/UDP/IPv4 o IPv4/PPP/L2TPv2/UDP/IPv6 o IPv4/PPP/L2TPv2/UDP/IPv4 o IPv6/PPP/L2TPv2/UDP/IPv6 Note that UDP bypass is not supported by L2TPv2 since L2TPv2 does not support "autodetect" of NAT/NAPT.
3
. Deployment Scenarios For the "Hub and Spoke" problem space, four scenarios have been identified. In each of these four scenarios, different home equipment plays the role of the Softwire Initiator. This section elaborates each scenario with L2TPv2 as the Softwire protocol andStorer, et al. Standards Track [Page 8]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 other possible protocols involved to complete the solution. This section examines the four scenarios for both IPv6-over-IPv4 (
Section 3.1
) and IPv4-over-IPv6 (
Section 3.2
) encapsulations.
3.1
. IPv6-over-IPv4 Softwires with L2TPv2 The following sub-sections cover IPv6 connectivity (SPH) across an IPv4-only access network (STH) using a Softwire.
3.1.1
. Host CPE as Softwire Initiator The Softwire Initiator (SI) is the host CPE (directly connected to a modem), which is dual-stack. There is no other gateway device. The IPv4 traffic SHOULD NOT traverse the Softwire. See Figure 1. IPv6 or dual-stack IPv4-only dual-stack |------------------||-----------------||----------| I SC SI N +-----+ +----------+ T | | | v4/v6 | E <==[ IPv6 ]....|v4/v6|....[IPv4-only]....| host CPE | R [network] | | [ network ] | | N | LNS | |LAC Client| E +-----+ +----------+ T _ _ _ _ _ _ _ _ _ ()_ _ _ _ _ _ _ _ _() <-- IPv6 traffic PPP o L2TPv2 o UDP o IPv4 (SPH) Softwire <------------------> IPV6CP: capable of /64 Intf-Id assignment or uniqueness check |------------------>/64 prefix RA |------------------>DNS, etc. DHCPv6 Figure 1: Host CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, the IPv6 Control Protocol (IPV6CP) negotiates IPv6-over-PPP, which also provides the capability for the ISP to assign the 64-bit Interface-Identifier to the host CPE or perform uniqueness validation for the two interface identifiers at the two PPP ends [
RFC5072
]. After IPv6-over-PPP is up, IPv6Storer, et al. Standards Track [Page 9]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 Stateless Address Autoconfiguration / Neighbor Discovery runs over the IPv6-over-PPP link, and the LNS can inform the host CPE of a prefix to use for stateless address autoconfiguration through a Router Advertisement (RA) while other non-address configuration options (such as DNS [
RFC3646
] or other servers' addresses that might be available) can be conveyed to the host CPE via DHCPv6.
3.1.2
. Router CPE as Softwire Initiator The Softwire Initiator (SI) is the router CPE, which is a dual-stack device. The IPv4 traffic SHOULD NOT traverse the Softwire. See Figure 2. IPv6 or dual-stack IPv4-only dual-stack |------------------||-----------------||---------------------| I SC SI N +-----+ +----------+ T | | | v4/v6 | +-----+ E <==[ IPv6 ]....|v4/v6|....[IPv4-only]....| CPE |----|v4/v6| R [network] | | [ network ] | | | host| N | LNS | |LAC Client| +-----+ E +-----+ +----------+ T _ _ _ _ _ _ _ _ _ ()_ _ _ _ _ _ _ _ _() <-------- IPv6 traffic PPP o L2TPv2 o UDP o IPv4 (SPH) Softwire <------------------> IPV6CP: capable of /64 Intf-Id assignment or uniqueness check |------------------>/64 prefix RA |------------------>/48 prefix, DHCPv6 DNS, etc. |------->/64 prefix RA |-------> DNS, etc. DHCPv4/v6 Figure 2: Router CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, IPV6CP negotiates IPv6-over-PPP, which also provides the capability for the ISP to assign the 64-bitStorer, et al. Standards Track [Page 10]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 Interface-Identifier to the router CPE or perform uniqueness validation for the two interface identifiers at the two PPP ends [
RFC5072
]. After IPv6-over-PPP is up, IPv6 Stateless Address Autoconfiguration / Neighbor Discovery runs over the IPv6-over-PPP link, and the LNS can inform the router CPE of a prefix to use for stateless address autoconfiguration through a Router Advertisement (RA). DHCPv6 can be used to perform IPv6 Prefix Delegation (e.g., delegating a prefix to be used within the home network [
RFC3633
]) and convey other non-address configuration options (such as DNS [
RFC3646
]) to the router CPE.
3.1.3
. Host behind CPE as Softwire Initiator The CPE is IPv4-only. The Softwire Initiator (SI) is a dual-stack host (behind the IPv4-only CPE), which acts as an IPv6 host CPE. The IPv4 traffic SHOULD NOT traverse the Softwire. See Figure 3. IPv6 or dual-stack IPv4-only dual-stack |------------------||----------------------------||----------| I SC SI N +-----+ +----------+ T | | +-------+ | v4/v6 | E <==[ IPv6 ]....|v4/v6|....[IPv4-only]....|v4-only|--| host | R [network] | | [ network ] | CPE | | | N | LNS | +-------+ |LAC Client| E +-----+ +----------+ T _ _ _ _ _ _ _ _ _ _ _ _ _ _ ()_ _ _ _ _ _ _ _ _ _ _ _ _ _() <-- IPv6 PPP o L2TPv2 o UDP o IPv4 traffic Softwire (SPH) <------------------------------> IPV6CP: capable of /64 Intf-Id assignment or uniqueness check |------------------------------>/64 prefix RA |------------------------------>DNS, etc. DHCPv6 Figure 3: Host behind CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, IPV6CP negotiates IPv6-over-PPP, which also provides the capability for the ISP to assign the 64-bitStorer, et al. Standards Track [Page 11]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 Interface-Identifier to the host or perform uniqueness validation for the two interface identifiers at the two PPP ends [
RFC5072
]. After IPv6-over-PPP is up, IPv6 Stateless Address Autoconfiguration / Neighbor Discovery runs over the IPv6-over-PPP link, and the LNS can inform the host of a prefix to use for stateless address autoconfiguration through a Router Advertisement (RA) while other non-address configuration options (such as DNS [
RFC3646
]) can be conveyed to the host via DHCPv6.
3.1.4
. Router behind CPE as Softwire Initiator The CPE is IPv4-only. The Softwire Initiator (SI) is a dual-stack device (behind the IPv4-only CPE) acting as an IPv6 CPE router inside the home network. The IPv4 traffic SHOULD NOT traverse the Softwire. See Figure 4.Storer, et al. Standards Track [Page 12]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 IPv6 or dual-stack IPv4-only dual-stack |------------------||-------------------------||-------------| I SC SI N +-----+ +----------+ T | | +-------+ | v4/v6 | E <==[ IPv6 ]....|v4/v6|..[IPv4-only]..|v4-only|---| router | R [network] | | [ network ] | CPE | | | | N | LNS | +-------+ | |LAC Client| E +-----+ | +----------+ T | ---------+-----+ |v4/v6| | host| _ _ _ _ _ _ _ _ _ _ _ _ _ +-----+ ()_ _ _ _ _ _ _ _ _ _ _ _ _() <-- IPv6 PPP o L2TPv2 o UDP o IPv4 traffic Softwire (SPH) <---------------------------> IPV6CP: capable of /64 Intf-Id assignment or uniqueness check |--------------------------->/64 prefix RA |--------------------------->/48 prefix, DHCPv6 DNS, etc. |----> /64 RA prefix |----> DNS, DHCPv6 etc. Figure 4: Router behind CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, IPV6CP negotiates IPv6-over-PPP, which also provides the capability for the ISP to assign the 64-bit Interface-Identifier to the v4/v6 router or perform uniqueness validation for the two interface identifiers at the two PPP ends [
RFC5072
]. After IPv6-over-PPP is up, IPv6 Stateless Address Autoconfiguration / Neighbor Discovery runs over the IPv6-over-PPP link, and the LNS can inform the v4/v6 router of a prefix to use for stateless address autoconfiguration through a Router AdvertisementStorer, et al. Standards Track [Page 13]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 (RA). DHCPv6 can be used to perform IPv6 Prefix Delegation (e.g., delegating a prefix to be used within the home network [
RFC3633
]) and convey other non-address configuration options (such as DNS [
RFC3646
]) to the v4/v6 router.
3.2
. IPv4-over-IPv6 Softwires with L2TPv2 The following sub-sections cover IPv4 connectivity (SPH) across an IPv6-only access network (STH) using a Softwire.
3.2.1
. Host CPE as Softwire Initiator The Softwire Initiator (SI) is the host CPE (directly connected to a modem), which is dual-stack. There is no other gateway device. The IPv6 traffic SHOULD NOT traverse the Softwire. See Figure 5. IPv4 or dual-stack IPv6-only dual-stack |------------------||-----------------||----------| I SC SI N +-----+ +----------+ T | | | v4/v6 | E <==[ IPv4 ]....|v4/v6|....[IPv6-only]....| host CPE | R [network] | | [ network ] | | N | LNS | |LAC Client| E +-----+ +----------+ T _ _ _ _ _ _ _ _ _ ()_ _ _ _ _ _ _ _ _() <-- IPv4 traffic PPP o L2TPv2 o UDP o IPv6 (SPH) Softwire <------------------> IPCP: capable of global IP assignment and DNS, etc. Figure 5: Host CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, the IP Control Protocol (IPCP) negotiates IPv4-over-PPP, which also provides the capability for the ISP to assign a global IPv4 address to the host CPE. A global IPv4 address can also be assigned via DHCP. Other configuration options (such as DNS) can be conveyed to the host CPE via IPCP [
RFC1877
] or DHCP [
RFC2132
].Storer, et al. Standards Track [Page 14]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
3.2.2
. Router CPE as Softwire Initiator The Softwire Initiator (SI) is the router CPE, which is a dual-stack device. The IPv6 traffic SHOULD NOT traverse the Softwire. See Figure 6. IPv4 or dual-stack IPv6-only dual-stack Home |------------------||-----------------||-------------------| I SC SI N +-----+ +----------+ T | | | v4/v6 | +-----+ E <==[ IPv4 ]....|v4/v6|....[IPv6-only]....| CPE |--|v4/v6| R [network] | | [ network ] | | | host| N | LNS | |LAC Client| +-----+ E +-----+ +----------+ T _ _ _ _ _ _ _ _ _ ()_ _ _ _ _ _ _ _ _() <--------- IPv4 traffic PPP o L2TPv2 o UDP o IPv6 (SPH) Softwire <------------------> IPCP: capable of global IP assignment and DNS, etc. |------------------> DHCPv4: prefix, mask, PD private/ |------> global DHCP IP, DNS, etc. Figure 6: Router CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, IPCP negotiates IPv4-over-PPP, which also provides the capability for the ISP to assign a global IPv4 address to the router CPE. A global IPv4 address can also be assigned via DHCP. Other configuration options (such as DNS) can be conveyed to the router CPE via IPCP [
RFC1877
] or DHCP [
RFC2132
]. For IPv4 Prefix Delegation for the home network, DHCP [
SUBNET-ALL
] can be used.Storer, et al. Standards Track [Page 15]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
3.2.3
. Host behind CPE as Softwire Initiator The CPE is IPv6-only. The Softwire Initiator (SI) is a dual-stack host (behind the IPv6 CPE), which acts as an IPv4 host CPE. The IPv6 traffic SHOULD NOT traverse the Softwire. See Figure 7. IPv4 or dual-stack IPv6-only dual-stack |------------------||----------------------------||----------| I SC SI N +-----+ +----------+ T | | +-------+ | v4/v6 | E <==[ IPv4 ]....|v4/v6|....[IPv6-only]....|v6-only|--| host | R [network] | | [ network ] | CPE | | | N | LNS | +-------+ |LAC Client| E +-----+ +----------+ T _ _ _ _ _ _ _ _ _ _ _ _ _ _ ()_ _ _ _ _ _ _ _ _ _ _ _ _ _() <-- IPv4 PPP o L2TPv2 o UDP o IPv6 traffic Softwire (SPH) <------------------------------> IPCP: capable of global IP assignment and DNS, etc. Figure 7: Host behind CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, IPCP negotiates IPv4-over-PPP, which also provides the capability for the ISP to assign a global IPv4 address to the host. A global IPv4 address can also be assigned via DHCP. Other configuration options (such as DNS) can be conveyed to the host CPE via IPCP [
RFC1877
] or DHCP [
RFC2132
].
3.2.4
. Router behind CPE as Softwire Initiator The CPE is IPv6-only. The Softwire Initiator (SI) is a dual-stack device (behind the IPv6-only CPE) acting as an IPv4 CPE router inside the home network. The IPv6 traffic SHOULD NOT traverse the Softwire. See Figure 8.Storer, et al. Standards Track [Page 16]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 IPv4 or dual-stack IPv6-only dual-stack |------------------||-------------------------||------------| I SC SI N +-----+ +----------+ T | | +-------+ | v4/v6 | E <==[ IPv4 ]....|v4/v6|..[IPv6-only]..|v6-only|---| router | R [network] | | [ network ] | CPE | | | | N | LNS | +-------+ | |LAC Client| E +-----+ | +----------+ T | --------+-----+ |v4/v6| | host| _ _ _ _ _ _ _ _ _ _ _ _ _ +-----+ ()_ _ _ _ _ _ _ _ _ _ _ _ _() <--- IPv4 PPP o L2TPv2 o UDP o IPv4 traffic Softwire (SPH) <---------------------------> IPCP: assigns global IP address and DNS, etc. |---------------------------> DHCPv4: prefix, mask, PD private/ |----> global DHCP IP, DNS, etc. Figure 8: Router behind CPE as Softwire Initiator In this scenario, after the L2TPv2 Control Channel and Session establishment and PPP LCP negotiation (and optionally PPP Authentication) are successful, IPCP negotiates IPv4-over-PPP, which also provides the capability for the ISP to assign a global IPv4 address to the v4/v6 router. A global IPv4 address can also be assigned via DHCP. Other configuration options (such as DNS) can be conveyed to the v4/v6 router via IPCP [
RFC1877
] or DHCP [
RFC2132
]. For IPv4 Prefix Delegation for the home network, DHCP [
SUBNET-ALL
] can be used.
4
. References to Standardization Documents This section lists and groups documents from the Internet standardization describing technologies used to design the framework of the Softwire "Hub and Spoke" solution. This emphasizes the motivation of Softwire to reuse as many existing standards asStorer, et al. Standards Track [Page 17]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 possible. This list contains both Standards Track (Proposed Standard, Draft Standard, and Standard) and Informational documents. The list of documents and their status should only be only used for description purposes.
4.1
. L2TPv2
RFC 2661
"Layer Two Tunneling Protocol 'L2TP'" [
RFC2661
]. * For both IPv4 and IPv6 payloads (SPH), support is complete. * For both IPv4 and IPv6 transports (STH), support is complete.
4.2
. Securing the Softwire Transport
RFC 3193
"Securing L2TP using IPsec" [
RFC3193
].
RFC 3948
"UDP Encapsulation of IPsec ESP Packets" [
RFC3948
]. * IPsec supports both IPv4 and IPv6 transports.
4.3
. Authentication, Authorization, and Accounting
RFC 2865
"Remote Authentication Dial In User Service (RADIUS)" [
RFC2865
]. * Updated by [
RFC2868
], [
RFC3575
], and [
RFC5080
].
RFC 2867
"RADIUS Accounting Modifications for Tunnel Protocol Support" [
RFC2867
].
RFC 2868
"RADIUS Attributes for Tunnel Protocol Support" [
RFC2868
].
RFC 3162
"RADIUS and IPv6" [
RFC3162
].
4.4
. MIB
RFC 1471
"The Definitions of Managed Objects for the Link Control Protocol of the Point-to-Point Protocol" [
RFC1471
].
RFC 1473
"The Definitions of Managed Objects for the IP Network Control Protocol of the Point-to-Point Protocol" [
RFC1473
].Storer, et al. Standards Track [Page 18]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
RFC 3371
"Layer Two Tunneling Protocol "L2TP" Management Information Base" [
RFC3371
].
RFC 4087
"IP Tunnel MIB" [
RFC4087
]. * Both IPv4 and IPv6 transports are supported.
4.5
. Softwire Payload Related
4.5.1
. For IPv6 Payloads
RFC 4861
"Neighbor Discovery for IP version 6 (IPv6)" [
RFC4861
].
RFC 4862
"IPv6 Stateless Address Autoconfiguration" [
RFC4862
].
RFC 5072
"IP Version 6 over PPP" [
RFC5072
].
RFC 3315
"Dynamic Host Configuration Protocol for IPv6 (DHCPv6)" [
RFC3315
].
RFC 3633
"IPv6 Prefix Options for Dynamic Host Configuration Protocol (DHCP) version 6" [
RFC3633
].
RFC 3646
"DNS Configuration options for Dynamic Host Configuration Protocol for IPv6 (DHCPv6)" [
RFC3646
].
RFC 3736
"Stateless Dynamic Host Configuration Protocol (DHCP) Service for IPv6" [
RFC3736
].
4.5.2
. For IPv4 Payloads
RFC 1332
"The PPP Internet Protocol Control Protocol (IPCP)" [
RFC1332
].
RFC 1661
"The Point-to-Point Protocol (PPP)" [
RFC1661
].
RFC 1877
"PPP Internet Protocol Control Protocol Extensions for Name Server Addresses" [
RFC1877
].
RFC 2131
"Dynamic Host Configuration Protocol" [
RFC2131
].
RFC 2132
"DHCP Options and BOOTP Vendor Extensions" [
RFC2132
]. DHCP Subnet Allocation "Subnet Allocation Option". * Work in progress, see [
SUBNET-ALL
].Storer, et al. Standards Track [Page 19]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
5
. Softwire Establishment A Softwire is established in three distinct steps, potentially preceded by an optional IPsec-related step 0 (see Figure 9). First, an L2TPv2 tunnel with a single session is established from the SI to the SC. Second, a PPP session is established over the L2TPv2 session and the SI obtains an address. Third, the SI optionally gets other information through DHCP such as a delegated prefix and DNS servers. SC SI | | |<-------------IKEv1------------->| Step 0 | | IPsec SA establishment | | (optional) | | |<-------------L2TPv2------------>| Step 1 | | L2TPv2 Tunnel establishment | | |<--------------PPP-------------->| Step 2 |<-----Endpoint Configuration---->| PPP and Endpoint | | configuration | | |<------Router Configuration----->| Step 3 | | Additional configuration | | (optional) Figure 9: Steps for the Establishment of a Softwire Figure 10 depicts details of each of these steps required to establish a Softwire.Storer, et al. Standards Track [Page 20]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 SC SI | | | | Step 0 |<------------IKEv1-------------->| = IKEv1 (Optional) | | | | Step 1 |<------------SCCRQ---------------| - |-------------SCCRP-------------->| | |<------------SCCCN---------------| | |<------------ICRQ----------------| | L2TPv2 |-------------ICRP--------------->| | |<------------ICCN----------------| - | | | | Step 2 |<-----Configuration-Request------| - |------Configuration-Request----->| | PPP |--------Configuration-Ack------->| | LCP |<-------Configuration-Ack--------| - | | |-----------Challenge------------>| - PPP Authentication |<----------Response--------------| | (Optional - CHAP) |------------Success------------->| - | | |<-----Configuration-Request------| - |------Configuration-Request----->| | PPP NCP |--------Configuration-Ack------->| | (IPV6CP or IPCP) |<-------Configuration-Ack--------| - | | |<------Router-Solicitation-------| - Neighbor Discovery |-------Router-Advertisement----->| | (IPv6 only) | | - | | | | Step3 | | DHCP (Optional) |<-----------SOLICIT--------------| - |-----------ADVERTISE------------>| | DHCPv6 |<---------- REQUEST--------------| | (IPv6 SW, Optional) |-------------REPLY-------------->| - | | or |<---------DHCPDISCOVER-----------| - |-----------DHCPOFFER------------>| | DHCPv4 |<---------DHCPREQUEST------------| | (IPv4 SW, Optional) |------------DHCPACK------------->| - Figure 10: Detailed Steps in the Establishment of a SoftwireStorer, et al. Standards Track [Page 21]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 The IPsec-related negotiations in step 0 are optional. The L2TPv2 negotiations in step 1 are described in
Section 5.1
. The PPP Network Control Protocol (NCP) negotiations in step 2 use IPV6CP for IPv6- over-IPv4 Softwires, and IPCP for IPv4-over-IPv6 Softwires (see
Section 5.2.4
). The optional DHCP negotiations in step 3 use DHCPv6 for IPv6-over-IPv4 Softwires, and DHCPv4 for IPv4-over-IPv6 Softwires (see
Section 5.4
). Additionally, for IPv6-over-IPv4 Softwires, the DHCPv6 exchange for non-address configuration (such as DNS) can use Stateless DHCPv6, the two-message exchange with Information-Request and Reply messages (see
Section 1.2 of [RFC3315]
and [
RFC3736
]).
5.1
. L2TPv2 Tunnel Setup L2TPv2 [
RFC2661
] was originally designed to provide private network access to end users connected to a public network. In the L2TPv2 incoming call model, the end user makes a connection to an L2TP Access Concentrator (LAC). The LAC then initiates an L2TPv2 tunnel to an L2TP Network Server (LNS). The LNS then transfers end-user traffic between the L2TPv2 tunnel and the private network. In the Softwire "Hub and Spoke" model, the Softwire Initiator (SI) assumes the role of the LAC Client and the Softwire Concentrator (SC) assumes the role of the LNS. In the Softwire model, an L2TPv2 packet MUST be carried over UDP. The underlying version of the IP protocol may be IPv4 or IPv6, depending on the Softwire scenario. In the following sections, the term "Tunnel" follows the definition from
Section 1.2 of [RFC2661]
, namely: "The Tunnel consists of a Control Connection and zero or more L2TP Sessions".
5.1.1
. Tunnel Establishment Figure 11 describes the messages exchanged and Attribute Value Pairs (AVPs) used to establish a tunnel between an SI (LAC) and an SC (LNS). The messages and AVPs described here are only a subset of those defined in [
RFC2661
]. This is because Softwires use only a subset of the L2TPv2 functionality. The subset of L2TP Control Connection Management AVPs that is applicable to Softwires is grouped into Required AVPs and Optional AVPs on a per-control-message basis (see Figure 11). For each control message, Required AVPs include all the "MUST be present" AVPs from [
RFC2661
] for that control message, and Optional AVPs include the "MAY be present" AVPs from [
RFC2661
] that are used in the Softwire context on that control message. Note that in the Softwire environment, the SI always initiates the tunnel. L2TPv2 AVPs SHOULD NOT be hidden.Storer, et al. Standards Track [Page 22]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 SC SI |<--------SCCRQ---------| Required AVPs: Message Type Protocol Version Host Name Framing Capabilities Assigned Tunnel ID Optional AVPs: Receive Window Size Challenge Firmware Revision Vendor Name |---------SCCRP-------->| Required AVPs: Message Type Protocol Version Framing Capabilities Host Name Assigned Tunnel ID Optional AVPs: Firmware Revision Vendor Name Receive Window Size Challenge Challenge Response |<--------SCCCN---------| Required AVPs: Message Type Optional AVPs: Challenge Response Figure 11: Control Connection Establishment In L2TPv2, generally, the tunnel between an LAC and LNS may carry the data of multiple users. Each of these users is represented by an L2TPv2 session within the tunnel. In the Softwire environment, the tunnel carries the information of a single user. Consequently, there is only one L2TPv2 session per tunnel. Figure 12 describes the messages exchanged and the AVPs used to establish a session between an SI (LAC) and an SC (LNS). The messages and AVPs described here are only a subset of those defined in [
RFC2661
]. This is because Softwires use only a subset of the L2TPv2 functionality. The subset of L2TP Call Management (i.e., Session Management) AVPs that is applicable to Softwires is grouped into Required AVPs and Optional AVPs on a per-control-message basis (see Figure 12). For eachStorer, et al. Standards Track [Page 23]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 control message, Required AVPs include all the "MUST be present" AVPs from [
RFC2661
] for that control message, and Optional AVPs include the "MAY be present" AVPs from [
RFC2661
] that are used in the Softwire context on that control message. Note that in the Softwire environment, the SI always initiates the session. An L2TPv2 session setup for a Softwire uses only the incoming call model. No outgoing or analog calls (sessions) are permitted. L2TPv2 AVPs SHOULD NOT be hidden. SC SI |<--------ICRQ---------| Required AVPs: Message Type Assigned Session ID Call Serial Number |---------ICRP-------->| Required AVPs: Message Type Assigned Session ID |<--------ICCN---------| Required AVPs: Message Type (Tx) Connect Speed Framing Type Figure 12: Session Establishment The following sub-sections (5.1.1.1 through 5.1.1.3) describe in more detail the Control Connection and Session establishment AVPs (see message flows in Figures 11 and 12, respectively) that are required, optional and not relevant for the L2TPv2 Tunnel establishment of a Softwire. Specific L2TPv2 protocol messages and flows that are not explicitly described in these sections are handled as defined in [
RFC2661
]. The mechanism for hiding AVP Attribute values is used, as described in
Section 4.3 of [RFC2661]
, to hide sensitive control message data such as usernames, user passwords, or IDs, instead of sending the AVP contents in the clear. Since AVPs used in L2TP messages for the Softwire establishment do not transport such sensitive data, L2TPv2 AVPs SHOULD NOT be hidden.Storer, et al. Standards Track [Page 24]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
5.1.1.1
. AVPs Required for Softwires This section prescribes specific values for AVPs that are required (by [
RFC2661
]) to be present in one or more of the messages used for the Softwire establishment, as they are used in the Softwire context. It combines all the Required AVPs from all the control messages in
Section 5.1.1
, and provides Softwire-specific use guidance. Host Name AVP This AVP is required in SCCRQ and SCCRP messages. This AVP MAY be used to authenticate users, in which case it would contain a user identification. If this AVP is not used to authenticate users, it may be used for logging purposes. Framing Capabilities AVP Both the synchronous (S) and asynchronous (A) bits SHOULD be set to 1. This AVP SHOULD be ignored by the receiver. Framing Type AVP The synchronous bit SHOULD be set to 1 and the asynchronous bit to 0. This AVP SHOULD be ignored by the receiver. (Tx) Connect Speed AVP (Tx) Connect Speed is a required AVP but is not meaningful in the Softwire context. Its value SHOULD be set to 0 and ignored by the receiver. Message Type AVP, Protocol Version AVP, Assigned Tunnel ID AVP, Call Serial Number AVP, and Assigned Session ID AVP As defined in [
RFC2661
].
5.1.1.2
. AVPs Optional for Softwires This section prescribes specific values for AVPs that are Optional (not required by [
RFC2661
]) but used in the Softwire context. It combines all the Optional AVPs from all the control messages in
Section 5.1.1
, and provides Softwire-specific use guidance.Storer, et al. Standards Track [Page 25]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 Challenge AVP and Challenge Response AVP These AVPs are not required, but are necessary to implement tunnel authentication. Since tunnel authentication happens at the beginning of L2TPv2 tunnel creation, it can be helpful in preventing denial-of-service (DoS) attacks. See
Section 5.1.1 of [RFC2661]
. The usage of these AVPs in L2TP messages is OPTIONAL, but SHOULD be implemented in the SC. Receive Window Size AVP, Firmware Revision AVP, and Vendor Name AVP As defined in [
RFC2661
].
5.1.1.3
. AVPs Not Relevant for Softwires L2TPv2 specifies numerous AVPs that, while allowed for a given message, are irrelevant to Softwires. They can be irrelevant to Softwires because they do not apply to the Softwire establishment flow (e.g., they are only used in the Outgoing Call establishment message exchange, while Softwires only use the Incoming Call message flow), or because they are Optional AVPs that are not used. L2TPv2 AVPs that are relevant to Softwires were covered in Sections
5.1.1
, 5.1.1.1, and 5.1.1.2. Softwire implementations SHOULD NOT send AVPs that are not relevant to Softwires. However, they SHOULD ignore them when they are received. This will simplify the creation of Softwire applications that build upon existing L2TPv2 implementations.
5.1.2
. Tunnel Maintenance Periodically, the SI/SC MUST transmit a message to the peer to detect tunnel or peer failure and maintain NAT/NAPT contexts. The L2TPv2 HELLO message provides a simple, low-overhead method of doing this. The default values specified in [
RFC2661
] for L2TPv2 HELLO messages could result in a dead-end detection time of 83 seconds. Although these retransmission timers and counters SHOULD be configurable (see
Section 5.8 of [RFC2661]
), these values may not be adapted for all situations, where a quicker dead-end detection is required, or where NAT/NAPT context needs to be refreshed more frequently. In such cases, the SI/SC MAY use, in combination with L2TPv2 HELLO, LCP ECHO messages (Echo-Request and Echo-Reply codes) described in [
RFC1661
]. When used, LCP ECHO messages SHOULD have a re-emission timer lower than the value for L2TPv2 HELLO messages. The default value recommended in
Section 6.5 of [RFC2661]
for the HELLO message retransmission interval is 60 seconds. When used, a set of suggested values (included here only for guidance) for the LCP ECHO messageStorer, et al. Standards Track [Page 26]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009 request interval is a default of 30 seconds, a minimum of 10 seconds, and a maximum of the lesser of the configured L2TPv2 HELLO retransmission interval and 60 seconds.
5.1.3
. Tunnel Teardown Either the SI or SC can tear down the session and tunnel. This is done as specified in
Section 5.7 of [RFC2661]
, by sending a StopCCN control message. There is no action specific to Softwires in this case.
5.1.4
. Additional L2TPv2 Considerations In the Softwire "Hub and Spoke" framework, L2TPv2 is layered on top of UDP, as part of an IP-in-IP tunnel;
Section 8.1 of [RFC2661]
describes L2TP over UDP/IP. Therefore, the UDP guidelines specified in [
RFC5405
] apply, as they pertain to the UDP tunneling scenarios carrying IP-based traffic.
Section 3.1.3 of [RFC5405]
specifies that for this case, specific congestion control mechanisms for the tunnel are not necessary. Additionally,
Section 3.2 of [RFC5405]
provides message size guidelines for the encapsulating (outer) datagrams, including the recommendation to implement Path MTU Discovery (PMTUD).
5.2
. PPP Connection This section describes the PPP negotiations between the SI and SC in the Softwire context.
5.2.1
. MTU The MTU of the PPP link presented to the SPH SHOULD be the link MTU minus the size of the IP, UDP, L2TPv2, and PPP headers together. On an IPv4 link with an MTU equal to 1500 bytes, this could typically mean a PPP MTU of 1460 bytes. When the link is managed by IPsec, this MTU SHOULD be lowered to take into account the ESP encapsulation (see [
SW-SEC
]). The value for the MTU may also vary according to the size of the L2TP header, as defined by the leading bits of the L2TP message header (see [
RFC2661
]). Additionally, see [
RFC4623
] for a detailed discussion of fragmentation issues.
5.2.2
. LCP Once the L2TPv2 session is established, the SI and SC initiate the PPP connection by negotiating LCP as described in [
RFC1661
]. The Address-and-Control-Field-Compression configuration option (ACFC) [
RFC1661
] MAY be rejected.Storer, et al. Standards Track [Page 27]
RFC 5571
Softwire H & S Framework with L2TPv2 June 2009
5.2.3
. Authentication After completing LCP negotiation, the SI and SC MAY optionally perform authentication. If authentication is chosen, Challenge Handshake Authentication Protocol (CHAP) [
RFC1994
] authentication MUST be supported by both the Softwire Initiator and Softwire Concentrator. Other authentication methods such as Microsoft CHAP version 1 (MS-CHAPv1) [
RFC2433
] and Extensible Authentication Protocol (EAP) [
RFC3748
] MAY be supported. A detailed discussion of Softwire security is contained in [
SW-SEC
].
5.2.4
. IPCP The only Network Control Protocol (NCP) negotiated in the Softwire context is IPV6CP (see
Section 5.2.4.1
) for IPv6 as SPH, and IPCP (see
Section 5.2.4.2
) for IPv4 as SPH.
5.2.4.1
. IPV6CP In the IPv6-over-IPv4 sc

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。

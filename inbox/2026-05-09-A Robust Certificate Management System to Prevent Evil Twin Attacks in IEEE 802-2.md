---
source: "https://arxiv.org/abs/2302.00338v1"
title: "A Robust Certificate Management System to Prevent Evil Twin Attacks in IEEE 802.11 Networks"
author: "Yousri Daldoul"
year: "2023"
publication: "arXiv preprint / cs.CR"
download: "https://arxiv.org/pdf/2302.00338v1"
pdf: "https://arxiv.org/pdf/2302.00338v1"
captured_at: "2026-05-09T12:24:13Z"
updated_at: "2026-05-09T12:24:13Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ニーチェ"
query: "Nietzsche Beyond Good and Evil"
tags:
  - "近代思想"
  - "実存主義"
  - "ニヒリズム"
status: raw
---

# A Robust Certificate Management System to Prevent Evil Twin Attacks in IEEE 802.11 Networks

- 著者: Yousri Daldoul
- 年: 2023
- 掲載情報: arXiv preprint / cs.CR
- 情報源: [arxiv](https://arxiv.org/abs/2302.00338v1)
- ダウンロード: https://arxiv.org/pdf/2302.00338v1
- PDF: https://arxiv.org/pdf/2302.00338v1

## Obsidian Links

- 研究動向: [[研究動向/ニーチェ-現代研究動向|ニーチェ-現代研究動向]]
- キーワード: [[ニーチェ]]
- 関連分野: [[近代思想]]
- 関連分野: [[実存主義]]
- 関連分野: [[ニヒリズム]]
- 関連タグ: #近代思想 #実存主義 #ニヒリズム

## Abstract

The evil twin attack is a major security threat to WLANs. An evil twin is a rogue AP installed by a malicious user to impersonate legitimate APs. It intends to attract victims in order to intercept their credentials, to steal their sensitive information, to eavesdrop on their data, etc. In this paper, we study the security mechanisms of wireless networks and we introduce the different authentication methods, including 802.1X authentication. We show that 802.1X has improved security through the use of digital certificates but does not define any practical technique for the user to check the network certificate. Therefore, it remains vulnerable to the evil twin attack. To repair this vulnerability, we introduce Robust Certificate Management System (RCMS) which takes advantage of the digital certificates of 802.1X to protect the users against rogue APs. RCMS defines a new verification code to allow the user device to check the network certificate. This practical verification combined with the reliability of digital certificates provides a perfect protection against rogue APs. RCMS requires a small software update on the user terminal and does not need any modification of IEEE 802.11. It has a significant flexibility since trusting a single AP is enough to trust all the APs of the extended network. This allows the administrators to extend their networks easily without the need to update any database of trusted APs on the user devices.

## PDF Text

A Robust Certifi cate Management System to Prevent

Evil Twin

Attack s in
IEEE
802.11 Networks

Yousri Daldoul

Faculty of Sciences of Monastir
, University of Monastir
, Monastir, Tunisia

yousri.daldoul@fsm.rnu.tn

AbstractThe
e vil
t win attack

is a major secu rity threat
to

WLANs
. An
e vil
t win is a rogue AP installed by a malicious user
to

impersonate

legitimate A
Ps
.
It intends to attract victim s

in
order
to
intercept their credentials, to
ste a
l their sensitive
information, to

eavesdrop
on
their data
, etc
.

In t his paper, we
study the security mechanisms of wireless networks and we
introduce the d ifferent authentication methods
, including
802.1X

authentication
. We show that 802.1X

has improved security

through
the use of digital certificates

but

does not define a ny

practical technique for the user to check the network certificate.

Therefore, it
remains vulnerable

to

the evil t win attack.
To repair
this vulnerability
, we introduce Robust Certificate Management
S
ystem
(RCMS)
which
takes advantage of the
digital cert ificates

of

802.1X
to
protect
the users
against rogue APs
.
RCMS

defines

a
new
verification
code
to

allow

the user device to
check

the
network

certificate.
This practical verification combined with the
reliability of digital certificates provides a perfect
protection
against rogue APs
.
RCMS

requires a small software update

on
the user terminal and does not
need

any modification
of

IEEE
802.11
.

It has a significant flexibility

since trusting a single AP is
enough to trust all the APs of the extended network.
This allows
the administrators to extend their networks easily without the
need to update any database of trusted APs

on the user devices
.

KeywordsIEEE 802.11 Networks; WLAN Security; 802.1X
Authentication; Evil Twin Attack; Certificate Verification

1.

I
NTR
ODUCTION

IEEE 80
2.11
‎
[1]

networks

are
widely used thanks

to their high
throughput
capacity
and
easy

installation. Due to
the

broadcast
nature of
these

networks
,

any
attacker

can eavesdrop on
their
transmitted data. Therefore, WLANs must provide enough

security
to protect the

user privacy
.
802.11i

is the principal
amendment that intends to improve the security. It defines

several protocols and algorithms to provide authentication,
integrity and confidentiality services. A WLAN that supports
802.11i is called a

Robust Security
Network
(RSN).

Although
802.11i introduces robust mechanisms,
an RSN

is

still
vulnerable to several attacks, such as the
e vil
t win attack. The
principle of this attack is to install a rogue AP
which
impersonates a legitimate AP
. When a new user wants to jo in
the WLAN, he may confuse the rogue AP with the legitimate
one and associate with
the rogue AP
. This allows the
adversary

to perform
several

attacks
, such as

intercept ing

the
user credentials,
steal ing sensitive information and

eavesdrop ping

on the victi m communication
.

WLANs are
suitable for

multiple
environments.
They

can

provide
public access
to open networks
in different areas
, such
as malls, municipalities, libraries and airports. They can also
provide
private

access
to

authorized users, like student s,
employees,
customers
, hotel guests and
family

members. This
is possible thank s

to the different
supported
authentication
methods. In fact, 802.11i defines 3 authentication methods:
Open System Authentication (OSA)
,
PreShared Key (PSK)

and 802.1X.
OSA
d oes not require any password and allows
any user to join the network. PSK requires the users and the
AP
to
sha re the same password. 802.1X

requires an
authentication server (AS) that authenticates the users by
means of their credentials (e.g.
username

and
password). Open
networks are vulnerable to evil twin attacks since there is no
mutual authentication

between the user and the AP
. PSK
allows mutual authentication using the shared password.

A
s
long as the password is protected, the connection is secure

and

the evil twin attack is
im possible
.

PSK is suitable for small
WLANs
,

such as residential networks
, where few users are
able to share the password securely
with each other
. It is not
convenient for public or large networks since the attacker is
able to obt ain the password
which allows the rogue AP to
succeed the mutual authentication with the victims. On the
other hand,
802.1X

‎
[2]

is suitable for large networks

since it
provides every user with his own credentials. It allows the

user
to authenticate the network by means of the digital certificate
of the AS, while the user credentials
allow the AS

to
authenticate the users. This authentication method is widely
used by companies, hotels, shops and universities, such as the
largest
university network

Eduroam
‎
[3]
.

Unfortunately,

the

evil twin attacks are ea sy to perform against 802.1X

and allow
the attacker to steal the user credentials. This is because the
victim ignores the AS certificate and
can

trust any selfsigned
certificate provided by the rogue AP.
T
herefore, he m ay send
his credentials in plaintext
to the attacker
within a

TLS tunnel.

Despite the robust security mechanisms of 802.11i, the evil
twin attack is easy to perform. As a result, a large number of
studies have been carried out to prevent this

attack. However,
most of them do not provide a trustworthy detection since they
may trust rogue APs and alert from legitimate APs. In addition,
several approaches are not practical since they have extensive
requirements (e.g. additional hardware, extensiv e use of the
bandwidth, multiple network interfaces, costly signed
certificates, etc.). Besides, we notice that all the reviewed
proposals do not provide enough security and are easy to
bypass. Therefore, it is necessary to define a practical and
reliable
approach to efficiently prevent the evil twin attack s
.

We believe that a robust
solution for
the
e vil
t win
problem

must

rely

on
digital

certificates.

This is because the rogue AP
cannot

impersonate the legitimate AP without the private key
.
However, it is
essential to provide the user with a practical and
reliable method to
verify

the AS certificate.
This
allows the
secure
association

with trusted
WLAN
s and the
efficient
detection of rogue APs
.

In this paper, we define

a

Robust
Certificate Management
System

(RCMS) to prevent all evil t win
attacks in WLANs.
Our proposal is suitable for both small and large networks
using
802.1X

authentication.

It runs entirely on the user device
and does not require any
protocol
modification. It allows the
user to strongly au thenticate the AS using an additional code

of a limited length
, called the
verification

code
.

Upon the first
association to an SSID, the user is requested to introduce his
credentials (e.g.
certificate or
username
/
password) and the
verification

code.

Once
the AS

certificate is checked correctly
,
the

root

Certification Authority
(CA)

of

the AS certificate

is
considered as
the

trusted
CA

of

the current SSID
. Therefore,
for any subsequent association to
a given

SSID, any

AS
certificate

is

trusted if
its root

C
A

is the trusted CA
of th e

SSID
.

This allows the network administrators to easily extend
their networks and to deploy multiple AS with different
certificates issued by the same CA.

The
user must provide the
verification
code
only

if the information
stored
by RCMS
on

the user device
does not allow
trusting

the AS

(e.g. first
association to the SSID or modified
public key of the root
CA).

RCMS

efficiently prevents evil twin attacks thanks to the
reliable verification of the AS.

Besides, our proposal
is
practi cal
since

it
only

require s

slight

software updates
on the
user
device
.

To summarize, we study the
e vil
t win attack s

in WLANs and
the limitation of existing security mechanisms. Our main
contribution is to introduce a new mechanism, called RCMS,
in

WLANs em ploying
802.1X

authentication. RCMS allows
the reliable
check

of the AS

using a
new verification

code

entered by the user
. Therefore, it prevents all evil twin attacks

and
allows
the secure association
to

legitimate

APs.

The rem a
inder of this pap er is orga nized as follows.
The next
S
ection

introduce s

related work studying
the
rogue AP
detection

in WLANs
. Then, Section
3

presents

the different
authentication methods of 802.11i and their
limitations against
the evil t win attack s
.

Section
4

presents the threat

model.
We
introduce
RCMS

in

Section
5

and
we conclude in S
ection
6
.

2.

R
ELATED
W
ORKS

Extensive research is carried
out

to define
secure
protection
mechanisms against
e vil twin attacks
.
The existing

approaches
can

be classified in to

4

main
families
:
traffic
a nomaly,
location
,
fingerprint

and cryptography

based approaches.

2.1

Traffic a nomaly

based

approaches

A large number of approaches are defined for
the

case of
a
rogue AP relay ing

the communication between the victims and
the legitimate AP.

Since this type
of evil twins

increase s the
number of wireless hop s and the delays, the authors
of

‎
[4]

choose
the i nterpacket
a rrival
time
as the detection

parameter.

In

‎
[5
]
, the
round trip time (RTT)

of ICMP packets is used for
the Evil Twin detection.
We believe that both

method s are

not
precise as the delays may vary s ignificantly in WLANs due to
several factors such as
the buffering delays,
the used data rates,
the number of users and the signal strength. Besides,
bridges
will be considered as rogue APs since they operate as relays
.

Another characteristic of evil twins

acting
as relays is frame
for warding on the wireless channel.
This characteristic is
considered by t he proposal of

‎
[6]

which

continuously
monitors
the medium to capture
and compare
the transmitted frames.
It
classifies
APs with
frame

forwarding as evil twins.
Legal AP
Finder (LAF)

‎
[7]

is a similar approach which relies on the
frame

forwar ding behavior of the
rogue APs
. Instead of
comparing all data frames, it only examines the TCP 3
-
way
handshake packets. We note that both

‎
[6]

and
‎
[7]

have
significant drawbacks and limited accuracy.

For example, they
are not suitable for encrypted networks because the encryption
algorithm makes
the
forwarded frames different
from

the
original frames
.

I
n

‎
[8]
,

the
rogue AP

detection relies on
the
statistics of
the
data tran smitted by the different APs. An evil
twin is identified if it transmits the same amount of data than
another AP during the same time interval. A similar proposal
is presented in
‎
[9]

and
detects

the forwarding behavior
by
monit oring

the arrival time of frames
having

similar lengths. It
requires multiple wireless interfaces (minimum of two) to scan
the different channels, and necessitates a long scan period to
detect any forwarding behavior.
PrAPHunter
‎
[10]

is a
detection mechanism
for network administrators
. It operates
on a dedicated device with two wireless interfaces.
T
he first
interface associates with
an

AP and transmits data, while the
second interface interferes with channels 1 to 11 sequential ly.
If the first interface notices throughput degradation when the
second interface is interfering with specific channels, this
indicates

that the AP is an evil twin forwarding data. It is clear
that this proposal suffers

from

significant drawbacks; not on ly
does it waste

the bandwidth, but also it does not provide a
trustworthy detection. This is because the throughput of
WLANs is variable due to several factors, such as medium
sharing, interference with legitimate devices, collisions and
channel fading. T
herefore,
PrAPHunter

cannot ensure a ny
effective

protection against rogue APs.

Other approaches
consider the case of

rogue APs
having
their
own Internet connection.

In

‎
[11]
,

the

detection method
relies
on the principle that th e different APs of a single Extended
Service Set (ESS) usually use the same gateway. Therefore, it
verifies the gateways

of

the visible

APs
belonging to

the same
SSID
, and
detects the presence of a

rogue AP
if

different
gateways

are used
.

Similarly, Rogue
AP Finder (RAF)
‎
[12]

compares the transmission paths to a given server over the
different APs of a particular ESS. It reports the presence of a
rogue AP if different paths are used.
Th ese

methods

may work
if
both
the rogue AP
and the

legitimate AP are visible.
Otherwise
, the attacker cannot be detected
.

Besides, the
two
proposal s

may detect the presence of an evil twin but
cannot
distinguish between
rogue
and legitimate
AP
s
.

Several o ther approaches
consider both types of

evil twins
:
relays and those having their own ga teways
.
I
n
‎
[13]
,

the
authors combine

the gateway
check

with the
frame

forwarding

verification
.
BiRe

is another detection mechanism defined
in

‎
[14]
.

It requires two wireless interfaces associated with two
different APs
.

Every interface s ends a TCP SYN packet to a
particular server which acknowledges the other interface. The
absence of
an
acknowledgement indicates th at

one of the two
APs is

an evil twin
.
We note that BiRe

cannot detect the
attack
if only

one

AP is available. Besides, it

ha s excessive
requirements which make it
im practical.

The proposal of
‎
[15]

intends to alert the network administrator of any existing evil
twin. It uses a sniffer that captures and analyses the
transmitted frames.
It considers that the attacker necessar il y
send s

deauthentication frames to disconnect the victims from
legal
networks and connect them with the rogue AP.
Therefore,
an attack is detected if excessive Association Response frames
are intercepted. Unfortunately,
this proposal is not suitable for
many types of evil twin s
.

EvilScout
‎
[16]

i s defined for a very
specific case of evil twins when the rogue AP operates on the
same channel
of

the legitimate
AP and impersonates
the MAC
address of the legitimate
AP. In this case, the detection is
based on
anomalies related to MAC
address conflict s
.

2.2

Locationbased approaches

To prevent the
e vil
t win attack, the authors of
‎
[17]

suggest the
connection of the legitimate AP to a display that confirms the
user connection to the right network. This proposal requires an
additional
device for every AP

and a line of sight between the
users and the display. This solution

is expensi ve and not
suitable for large networks

since the simultaneous verification
of multiple
people
using

a single screen

is not practical
.

The principle of crowd sensing is used by CRAD in

‎
[18]
. The
crowd is composed of the mobile u sers connected to a specific
ESS. Every user device should profile the different available
APs by recording their signal strength s

(i.e. RSS
I
) over time.
Then it shares its measurement reports with the other members
of the crowd.

The

ratio of reports conta ining a

significant
variation of the RSS
I

is used as
an

indicator of an existing
rogue AP.
However
,
an attacker can broadcast forged reports
to decrease the ratio
of reports
with RSSI

anomalies
.

A similar
approach based on crowd sensing is proposed in

‎
[19]
. It uses
the CSI (Channel State Information) and AoA (Angle of
Arrival) to improve the detection accuracy
, and detects the
attack if the spatial location of the AP changes
.

Another
proposal based on RSSI is defined

for
residen tial
networks

in

‎
[20]
. It considers that the signal strength is a stable
parameter that can be used to detect evil twins. This
assumption is not valid

due to the user mobility

and cannot
provide
reliable

and precise attack dete ction.

The princip le

of
RSSI monitoring is also used in
‎
[21]

to detect rogue APs
based on their location.

However, instead of using the crow d

collaboration, the authors suggest to install multiple sensors

that record the RSSI e volution. These measurements are
transmitted to and processed by a remote server to detect any
anomaly
.
This proposal
alert s

the network
administrato r

if a
rogue AP is detected

but does not prevent the attack
.

In

‎
[22]
, the auth ors‎use‎the‎principle‎of‎“trust‎by‎location”‎that‎
records all the visible APs upon the first association to an AP.
For subsequent connections, the AP is trusted if the variation
of the neighbor networks does not exceed a given threshold.
Otherwise, it is c lassified as an evil twin. In

‎
[23]
, the detection
system starts by classifying all the visible APs as authorized
and records their parameters in a whitelist. Then, it checks

for

any suspicious modification of different parameters to report a
rogue AP. For example, if a new AP is det ected after the
initialization step, it is considered as an evil twin. It is clear
that this approach is not reliable as it may classify many
legitimate APs as illegal and may trust rogue APs.

2.3

Fingerprintbased approaches

ETGuard
‎
[24]

is an administratorside mechanism which
detects rogue APs based on the recorded fingerprints. It runs
on a dedicated server and continuously records the beacon
frames. Since the fingerprints are calculated from the beacon
frames, any attacker is ab le to spoof these frames and
impersonate

legitimate AP
s.
This affects the reliability of

ETGuard.

Multiple approaches use radiometric signature as a
unique identifier of each device. In

‎
[25]
, the observation of the
clock skew i s used as the AP fingerprint.
In

‎
[26]
, the authors
use the CSI to extract the physical layer information of the
transmitter. They consider that the phase errors depend on the
device

and can be
used

to
create

a unique

fingerprin t

of
any

AP
.

Another approach
‎
[27]

extracts the AP

fingerprints from
the power
amplifier

and frame distribution of the received data.
The mechanism proposed in
‎
[28]

detects rogue APs based on
multip le parameters, namely the

clock skew,
the used
channel,
the
received signal strength and
the
beacon transmission
duration
. These proposals must

be initialized with a fingerprint
list of authorized devices.
Due to

this constraint, any network
extension or m odification requires the update of the fingerprint
list of every user.
We note that the

attacker can obtain
a

device
identical
to the used
AP. This allows the rogue AP to produce
the same fingerprint and to bypass the detector.

2.4

C
ryptographybased appr oaches

VOUCHAP
‎
[29]

is

among a few proposals that use

digital
certificates to
authenticate the legitimate AP and to
prevent the
attacks. The authors
provide each AP with a certificate
issued
by
a

trusted Certification Authorit y
(
CA
).
This certificate
includes

the network SSID and aims to prevent WLAN
impersonation
. Unfortunately, the SSID is not a unique
identifier for WLANs and can be used by different networks
simultaneously. Therefore,
an

attacker
can

obtain a signed
certifi cate from a trusted CA for
any

SSID

and perform the
evil twin attack
.

As a result,
this proposal is not secure enough
and incurs additional costs

related to the purchase of a
signed
certificate

for every AP
.

In

‎
[30]
, the authors show that the use of
WPA2
-
Enterprise (i.e.
802.1X

authentication
) remains vulnerable to the evil twin
attack which allows the adversary to steal user credentials.
This is because the user
ignores the

AS certificate

and cannot
check it.
Therefore, he may accept any certificate
, including
that of the attacker
.

To solve this problem, the authors suggest
to display
WPA2
-
Enterprise networks

in a list of pairs
«
SSID,
AS name

»
. As the authors

recognize
, this solution is not
secure since the a ttacker is able to produce a certificate (either
selfsigned or signed by a trusted CA) containing the same AS
name.

In

‎
[31]
,

the
authors show that the 802.1X

authentication
used in Eduroam networks does not
sufficiently
secure
WLANs since most users do not check the AS certificate.
Therefore, they suggest activating the check of the AS
name

(i.e. displaying the informati on of the AS certificate
in

an
interface and asking
for
the user
permission

before pursuing
the authentication)
. This is not a
reliable

approach since
selfsigned certificates are
widely

used in WLANs, allowing the
attacker to use
any

AS name. In addition,

most users are not
aware about the AS name

and

trust the WLAN based on its
SSID.

A similar study of the Eduroam security
‎
[32]

shows that

the authors

were able to access the user credentials of

61% of
the
tested
devices
which

accepted to a ssociate with a rogue AP.

To prevent
the
evil twin

attacks
, Eduroam provides a
Configuration Assistant Tool

(CAT)

‎
[33]

that configures the
user
device with the Eduroam network profile. This solution
requires t he user to download and execute CAT. Since the use
of CAT is not mandatory, most users may ignore it.

We note
that the created profile does not prevent the

association with a
rogue AP but
allows to inform

the user that the network details
have changed

and requests the user authorization to pursue the
authentication.

Therefore, we believe that CAT is neither
practical nor reliable.

3.

B
ACKGROUND OF
WLAN

S
ECURITY

3.1

Network Discovery

In a WLAN, every AP is identified using a unique identifier
called BSSID (i.e. the MAC address of the AP). To extend the
coverage of a WLAN, the administrator
may

install multiple
APs. The extended network is called Extended Service
Set
(ESS) and is identified using
a string called

SSID.
To join a
WLAN, the user station (STA) follows 3 steps: network
scanning (active or passive), authentication and association.
During the first step, the STA scans the different channels of
the spectru m to find the available networks. Using passive
scanning, the STA receives the beacon frames of the visible
APs
. These frames are broadcasted periodically and contain all
the information about the AP, such as SSID, BSSID and
the
security protocol.

They all ow the user to select the desired
SSID
.
If multiple APs

belonging to the same SSID are visible,
the STA selects the AP with the highest signal strength (i.e.
RSSI) as it is expected to provide the highest throughput.

During the user mobility, the STA may p erform a seamless
handover from one AP to another within the same ESS. This
handover is defined by the IEEE 802.11
standard
and does not
require the user permission.

3.2

Authentication and Association

The second step after network scanning is user authenti cation.
Current networks support 3 authentication methods
: Open
S
ystem
A
uthentication

(OSA)
,
PreShared Key (PSK) and
80
2.1X
.

The first method
does not
use

any password

and does
not provide any authentication. It

allows any user to join the
network

if his
MAC address is not blacklisted
.

This method
does not use encryption and the network frames are
transmitted in the clear. A recent enhancement of open
authentication is defined in
‎
[34]

and allows data encryption in
open networks. As there is no way to authenticate the users
and the
WLAN
, an evil twin attack is easily performed

against
open networks and cannot be detected or
prevented
.

The second authentication method is PSK. It re quires the users
and the AP to share the same password. During authentication,
both the STA and the AP must prove knowledge of the secret.
This ensures
mutual authentication between the
user

and the
network. Without the password, a rogue AP cannot
authenti cate to the STA and cannot establish a connection with
the victim. Since the transmitted frames are encrypted in a
WLAN

protected
with PSK
, the adversary cannot eavesdrop
on the data and cannot perform any attack. We note that PSK
is practical in small WLA
Ns, such as residential networks,
as
long as

the few users are able to
keep the password
confidential
.
PSK is not suitable for public or large networks
since the password is accessible to any user.
For example,
several restaurants and
cafes

provide their c ustomers with free
connections to WLANs protected
by

PSK. Typically, they
provide them with the password within the receipt. This allows
any malicious user
to obtain the secret,
to impersonate the
legitimate AP and to perform the evil twin attack.

The
thir d
aut hentication method is 802.1X
.

It is widely known
as WPA2
-
Enterprise.

It requires an authentication server (AS)
which performs the mutual authentication with the users by
the intermediate of the AP. The AS uses its certificate to
authenticate
itself
to

the user. I
f

the user trusts the
AS
certificate, he uses his credentials (e.g.
certificate or
username
/
password) to authenticate
to

the
server
.

Th e AS

certificate allows the establishment of a secure tunnel between
the user and the
AS

to
perform

the user
authentication
.

802.1X

authentication

is suitable for large networks since every user
has his own credentials

and can identify legitimate APs thanks
to the AS certificate. It
prevents evil twin attacks and
guarantees data confidentiality
in public networks
, even if the
same
username

and password are publicly shared, as long as
the users are able to verify the AS certificate.

It can also be
used in small networks since many commercial APs have
integrated AS and can use selfsigned certificates.

Unfortunately
, a large number of evil twin attacks are
succe ssfully achieved against 802.1X

and allow the adversary
to steal the user credentials and to
eavesdrop on the traffic
.

This is because most users cannot verify the AS certificate and
accept to authenticate wit h rogue APs providing selfsigned
certificates. Therefore, they establish a secure tunnel with the
attacker who
becomes
able to perform multiple attacks.

As aforementioned,
OSA

is a null authentication protocol. It
uses

a twoframe

exchange
. The first
fram e

contains the STA

identity

(i.e. the MAC address)

and requests authentication.
The second
frame

returns the authentication result. If the r esult
is‎“successful,”‎th e STA and the AP are considered

mutually
authenticated.

As depicted in
Figure
1
, t he authentication step
of PSK is either OSA or Simultaneous Authentication of
Equals (SAE).

SAE intends to make PSK resistant to offline
dictionary attacks. It g enerally

uses
the
elliptic curve
cryptography to derive
an inte rmediate key
, called
Pairwise
Master Key (PMK)
,
from the PSK.
When OSA is used with
PSK, PMK is identical to the preshared key (
i.e.
PMK=PSK).
In the case of 802.1X, the authentication step relies on OSA.

The association step is a twoframe
transaction se quence
following the authentication. It is initiated by the STA and
allows the negotiation of the connection parameters.
In the
case of o pen authentication
, no
more

step s

are

required and the
user device is successfully
associated

to the WLAN. But

all
the
frames
of

open networks are transmitted in the clear.

If
802.1X is used, the 802.1X authentication step starts following
the association. It allows the STA and the AS to derive the
PMK

from the TLS master key which is used to establish the
TLS tunnel.

Then
, the AS sends the
PMK

to the AP. This
allows the STA a nd the AP to share the same key
.

I
f PSK or
802.1X

is used, the
AP

and the
STA

start the 4
-
way
handshake to derive the Pairwise Transient Key (PTK) from
the PMK. This handshake intends to confirm that a

live STA
holds th e PMK and to derive a fresh PTK
.

The PTK is used to
encrypt the
transmitted
data frames
.

Figure
1

illustrates the
network access steps using the 3 authentication methods.

OSA

method

PSK
method

802.1X
method

Network scanning
(active or passive)

OSA

OSA

SAE

OSA

Association

802.1X
Authentication

4
-
way handshake

Network access enabled (controlled port unblocked)

Figure
1
.
Network access steps in

IEEE

802.11 WLANs

3.3

802.1X

authentication

802.1X authentication supports different credential types, such
as digital certificates, usernames and passwords, secure tokens,
and mobile network credentials (i.e. GSM and UMTS se crets).

A WLAN
employing

802.1X

typically consists of user devices,
one or multiple APs belonging to the same ESS, and one AS.
For large networks, such as Eduroam
‎
[3]
,
it is possible to use

multiple

servers
.
Figure
2

illustrates

a simple WLAN using
802.1X
.
The most used authentication server is the RADIUS
server which uses the RADIUS protocol to communicate with
the AP.
T
he refore, t he mutual authentication

between the user
and the AS

is performed using two protocols: EAP over LAN
(EAPOL) and
RADIUS. EAPOL is introduced by 802.1X

and

relies on EAP

‎
[35]
. It

defines additional frames to support
wire d and wireless LANs
.
EAP is an authentication protocol
used between the STA and the AS.
T
he EAP messages are
transmitted within
802.11 frames

over the wireless medium,
and within RADIUS packets between the AP and the
AS
.

Figure
2
. Ne twork architecture using 802.1X

authentication

EAP supports multi ple authentication methods

‎
[36]

which can
be classified into two principal categories:
passwordbased
and

TLSbased methods
. However,

not all of them are
compliant with

WLANs. In fact,

802.11i requires the use of
an
EAP

method
capable of

generating
the keying material

‎
[37]
.

Therefore, only TLSbased methods are
compliant with the
RSN requirements
. They establish a secure TLS tunnel
between the user device and the AS using the server certificate.

If the user authenticates using his username and password,

a
p asswordbased

authentication method must be used
through

the
encrypted
tunnel and is called
inner
or tunneled method
.

This inner method may be EAP or nonEAP method,
depending on the used TLSbased EAP method.

The most
popular TLSbased
EAP
methods are
:

-

EAPTLS

‎
[38]
: This method allows
the user and
the
AS

to mutually authenticate using certificates.

Therefore,
both the AS and the user must have certificates.
This
method is mainly used in
large

companies where the
network administrators take care of configuring
the
device of every employee individually
.

We note that
EAPTLS

is supported by all
device s since it is among
the requirements of

WPA2.

-

EAPTTLS

‎
[39]
: This method only requires the server to
hold a certificate and is, therefore
, more practical than
EAPTLS. It allows the user to authenticate using

his

username and password
through

the TLS tunnel. This
method supports multiple
inner methods.

It supports both
nonEAP (e.g. PAP, CHAP and MSCHAPv2) and EAP
methods (e.g. EAPMD5 and
EAPMSCHAPv2
)
.

Like
EAPTLS,
EAPTTLS is
also
supported by any device
compliant with WPA2.

-

PEAP

‎
[40]
: This method is similar to EAPTTLS, but
only suppor ts EAP methods as inner methods
.

On the other hand, the most used inner methods are:

-

PAP

‎
[41]
: This method allows the user authentication
using his username and passwo rd.
T
hese credentials are
transmitted in
plaintext

and are easily accessible if the
TLS tunnel is established with the attacker
.

-

CHAP

‎
[42]
, MSCHAP

‎
[43]

and
EAPMD5

‎
[37]
: These
are
oneway authentication

methods which allow the
server to authenticate the user using challenges.

-

MSCHAPv2

‎
[44]

and
EAPMSCHAPv2

‎
[45]
: These
methods provide mutual authentication
using
challenges;
both the server and the user
must

prov e their

knowledge
of the user
password
.

Figure
3

illustrates a n

example of
802.1X
a uthentication

using
EAPTTLS and EAPMD5

as the inner method
.

STA

(
User

device)

AP

(RADIUS client)

AS (RADIU
S
server)

User (STA)

AP

RA
DIUS Server

1) EAPOLStart

2) EAPRequest/Identity

3) EAPResponse/Identity

4) RADIUS AccessRequest:

EAPResponse/Identity

5) RADIUS AccessChallenge:

EAPRequest/TTLS:
Start

6)
EAPReques t/TTLS:
Start

7)
EAPRespons e/TTLS: ClientHello

8)
RADIUS AccessRequest:

EAPResponse/TTLS: ClientHello

9)
RADIUS AccessChallenge:

EAPRequest/TTLS:
ServerHello
,

Certificate
,

ServerKeyExchange
,

ServerHelloDone

10) EAPRequest/TTLS:
ServerHello
,

Certificate
,

ServerKeyExchange
,

ServerHelloDone

11)
EAPResponse/TTLS: ClientKeyExchange
,

ChangeCipherSpec
,

Finished

12) RADIUS AccessRequest:

EAPResponse/TTLS: ClientKeyExchange
,

ChangeCipherSpec
,

Finished

13)
RADIUS AccessChallenge:

EAPRequest/TTLS: ChangeCipherSpec,
Finished

14)
EAPRequest/TTLS: ChangeCipherSpec,
Finished

15)
EAPResponse/TTLS:
{
{
EAPResponse/Identity}
}

16) RADIUS AccessRequest:

EAPResponse/TTLS:
{
{
EAPRespon se/Identity}
}

17) RADIUS AccessChallenge:

EAPRequest/TTLS:
{
{
EAPRequest/MD5
-
Challenge}
}

18)
EAPRequest/TTLS:
{
{
EAPRequest/MD5
-
Challenge}
}

19)
EAPResponse/TTLS:
{
{
EAPResponse/MD5
-
Challenge}
}

20)
RADIUS AccessRequest:

EAPResponse/TTLS:
{
{
EAPResponse/MD5
-
Challenge
}
}

21) RADIUS AccessAccept:

MSK,
EAPSuccess

22)
EAPSuccess

Figure
3
. 802.1X authentication using EAPTTLS and EAPMD5

4.

T
HREAT
M
ODEL

4.1

Attack objectives

T
his sec tion introduce s
the
different evil twin attacks

against
802.1X authentication

according to

the adversary expectations
.

In fact, we dis tinguish two attack objectives:

-

C
redential
theftD
ata relay

(
maninthemiddle)

In the first case, the attacker intends to

steal
valid
credentials
in order to access the

WLAN as an authorized user. This is
a
typical attack against

several

private and paid networks

(e.g.
university
, airport and Internet provider WLANs) where the
network access is
limited
to authorized users on ly. The
damage s of this attack depend on the access rights of the
victim and vary

from simple

to very harm ful

damages, such as
bandwidth sharing,

paid plan consumption and unauthorized
access to personal documents and sensitive information.

The second obje ctive is to relay the victim
’s

data.
We note that
the attacker can easily provide an Internet connection using
different methods, such as mobile, wireless or wired
networks
.
Once the victim is connected to
the
Internet
through

the rogue
Established TLS tunnel: messages
between «

{{ » and «

}}

» are encrypted

AP,

the adversary
c an

perform
various
passive and active
attacks
, such as data eavesdropping

and

user redirection to
phishing websites
.

Although most sensitive websites use
https
to encrypt and protect the user data,
several

websites still use
http and
can

be spied.
Since ma ny people use the same
username and password to access different accounts, stealing
the credentials
from

unencrypted websites may allow the
attacker to

gain

access
to
the user accounts of sensitive
websites.
In addition, phishing websites may succeed to
st eal
sensitive
data (e.g. passwords and credit card information)
and

to convince the
victim

to download malware. We note that
some malware are very harmful and allow the attacker

to

easily
spy and control the victim device.

To steal the user credentials, th e attacker does
not

need to
provide an Internet connection
or

to be in visibility with a
legitimate AP. He only needs to install a rogue AP, to capture
the required information and to leave. But to relay the user
data, the attacker must provide an Internet

connection either
using his own gateway (mobile network or wired
LAN
) or
using the legitimate WLAN. In the latter case, a legitimate AP
must be visible.

4.2

Credential theft

As previously mentioned,
EAPTLS

allows the user and the
AS to mutually authentic ate using certificates.

This
is the most
secure EAP

method since
the user certificate is useless
to

the
adversary‎without‎the‎user’s‎private‎key‎which‎is‎never‎
transmitted
over the network
. Therefore, EAPTLS is perfectly
s ecure against credential theft. F
or the other TLSbased
methods,

the credentials are safe if the user associates with the
legitimate AP. But if the victim associates with the rogue AP

and
accepts
any

certificate, the encrypted tunnel
is

established
with the attacker who
can decrypt the tu nneled data
.

In this
case, the attacker can obtain the victim credentials using two
possible attacks: downgrade and dictionary attacks.

The downgrade attack allows the
adversary

to negotiate
the
weakest

possible

EAP

method

in order to facilitate the

acce ss

to the credentials. In fact, EAP has
several
methods which are
not necessarily supported by
every STA and AS. In a typical
scenario, the AS suggests
EAP methods from strongest to
weakest
till a method is accepted by the STA
. This
allows

the
selection of

t he strongest metho d

supported by both parties
. In
the case of
a

downgrade attack, the
malicious
AS suggests the
methods from weakest to strongest

in order to use the weakest
possible
.

If EAPTTLS

with PAP
is

used, the attacker receives
the
user

credentials

in plaintext and no more action is required.
But if a
challengeresponse
method is
used, the attacker
performs an off line dictionary attack
using
the challenge and
the
received
response
. This attack succeeds only if the victim
password figures within
the
dictionary of likely passwords

used by the attacker
.
Since the
adversary
’s

purpose is to steal
the user credentials,
he

does not need to succeed the
authentication
step

or to provide
the victim with
an Internet
connection. He only needs the credentials in
plaintext or the
challenge and the corresponding response.

4.3

Data relay: ManInTheMiddle

(MITM)

EAPTLS is not secure against MITM;

if the victim trusts the
rogue AP and accepts its certificate,
the adversary

accepts the
victim certificate and

succeed s

the mutual authentication. In
this case,
the vic tim data
will be

relayed through the
rogue AP
.

Similarly,
the
other TLSbased EAP methods do not prevent
the victim from accepting the
attacker

certificate. Once the
certificate is accepted, the success of t he mutual authentication
depends on the security of the inner method. Hence, if the
adversary succeeds to
negotiate

an inner method that allow s

oneway authenti cation (e.g. PAP, CHAP, EAPMD5

and
EAPMSCHAP),
he succeeds the authentication step

easily
.

I
f

the
victim

refuses all weak inner methods and only
accept s

a

mutual authentication
method
(e.g
. EAPMSCHAPv2), the
attacker

must prove
knowledge of the password

and reply
correctly to the challenge
.

This makes the authentication
more

difficult, but possib le if a legitimate AP is visible

to

the
attacker.

In this case, the attacker impersonates the STA

to the
AS

and
establishes a second tunnel with the AS. Then, he
negotiates the same authentication
method
.

Upon the reception
of the AS challenge,
the attacke r

sends this challenge to the
victim. Then, he forwards the
victim
’s

response

and challenge

to the AS.
Finally
, he
forwards
the AS response to the victim
.
This allows the attacker to succeed

the mutual authentication
with both STA and AS. In the remainder,

the victim
and the
rogue AP derive the same session keys and the victim data are
relayed through the evil twin

as desired by the
attacker
.

In
many

environments (e.g. restaurants,
cafes
,
libraries, etc.),
the network is available for
customers

and is

prote cted using
the same credentials

for all the users
. Therefore, the attacker is
able to obtain the
shared
password and to succeed the mutual
authentication of the inner method without the need to interact
with the legitimate AP. In this case, the only protec tion against
the evil twin attack is the verification of the AS certificate.

4.4

Summary

To summarize, the

evil twin attacks

against 802.1X are
possible
when

the victim

does not

verify the
AS
certificate and
accepts an y
one
.

If the user is able to check th e certificates and
associates with authorized APs
, no evil twin attack is possible.
Table
1

illustrates the security
level
of the most used EAP
methods and the possible attacks to attain the adversary
objectives.

Fo r a given inner method, we consider that this
method is selected following the downgrade attack

and is the
weakest
possible
method that the attacker can negotiate
.

Table
1
. Security level of most used EAP methods

Main EAP
method

In ner method

Adversary objective

Credential
theft

Data relay:
MITM

EAPTLS

null

Impossible

Easy

EAPTTLS

PAP

Easy

Easy

CHAP, EAPMD5,
MSCHAP, EAPMSCHAP

Possible
using
an
offline
dictionary
attack

Easy

MSCHAPv2,

EAPMSCHAPv2

Possible
: requires
legi timate
AP visibility

PEAP

EAPMD5, EAPMSCHAP

Easy

EAPMSCHAPv2

Possible
: requires
legitimate
AP
visibility

5.

R
OBUST
C
ERTIFICATE
M
ANAGEMENT
S
YSTEM

In this section, we introduce our approach to protect WLANs
against any type of evil twin attack
.

Our so lution is called
Robust Certificate Management System (RCMS) and is
defined for 802.1X authentication. It allows the user device to
easily and
precisely
check the AS certificate. Therefore,
RCMS only accepts legitimate certificates and associates with
auth orized networks. It rejects any authentication with rogue
APs thanks to
a
new
code‎called‎“
verification

code
”
.

This
code allows the verification of the server certificate

and the
authentication

with legitimate networks
.

RCMS is suitable
for all types of cr edentials
, but we mainly focus on the case
of username/password pairs which are widely used with
802.1X.

U
pon the first association
to

an SSID
, the user must
provide 3 valu es instead of 2:

username, password and
the
verification
code.

If the code is valid,

the network is trusted
and is added to th e
list

of trusted networks.
For subsequent
associations
to

trusted APs, the code is not requested

unless
the public key of the root certificate is modified
.

To successfully authenticate the servers, RCMS
maintains

the
certificates of the
root
CA instead of the AS certificates.

When
the STA receives a new AS certificate, it checks the root
certificate and accepts the AS certificate if the root CA is
trusted.

This allows large networks to use multiple
servers

with dif ferent certificates having the same root certificate.

In
the case of a
small networ k

having a single AS
and

a selfsigned certificate
, the root certificate is the AS certificate
.

Therefore,
our design is suitable for both small and large
networks
.

We note
that the root CA does not need to be
public as this incurs additional fees without any improved
security. However, it is possible and more practical (i.e. free
and more secure) to use a private CA, i.e. a selfsigned
certificate that is used directly to si gn the AS certificates or to
sign intermediate CA certificates.

In addition,
RCMS associates
a single root certificate

to

an
SSID
. Therefore, a trusted CA is only trusted for the
corresponding SSID.

Since certificates may be renewed or
updated, RCMS is abl e to update the stored root certificate
seamlessly as long as the public key has not changed. But if
the public key of the trusted CA is modified, the user must
provide a new code to check the AS certificate.

We note that
updating the certificate does not
require the modification of its
public key unless the private key is compromised.

Moreover,
the root certificate is generally valid for
many

years (typically
3

to 20 years)
. Hence,
the
verification
code is rarely
requested

after the very first association
to a WLAN
.

5.1

V
erification
code calculation

The
verification
code
is used to check

the

AS certificate. It can
be calculated differently
:

1)

The code is derived
either
from the AS certificate or
from the root certificate
:

These are two possible
options.
In bo th cases, RCMS
save s

the root certificate as
the

trusted CA upon the
first
successfu l verification
.
If
the code is derived from the AS certificate
,
the following

constraint must be satisfied: the user must authenticate
the first time with the AS for which
the code is
generated
.

To get rid of this constraint, we can
derive the code from
the root certificate. This
second
option is more flexible
and

allows the first authentication to occur with any AS
of the ESS.

2)

The code is calculated
either
based on the enti re
certificate or
based on

the public key
:

If the code is
calculated based on the entire certificate
, a particular
authentication failure may occur in the following case:
the code is calculated and then the certificate is updated
before the first connectio n of the user. In this case, the
authentication fails and the user must obtain a new code.
We note that this is a particular case since
a

certificate
is
not modified frequently. However, we can avoid this
particular
case if we calculate the code based on t he
public key. In fact, the public key does not change during
updates and renewals, unless the
corresponding
private
key is compromised.

3)

The code is calculated using
either
a hash or a keyedhash function
:

It is possible to use a hash function to
calculate

the code. In this case, the code is common
to

all
users and is
easily

accessible
by an

adversary. Therefore,
the code must be long enough to be resistant to a brute
force attack.

Th e second option is to calculate the
code
using

a cryptographic hash functi on and

the user
password. Therefore, the code is not a common value and
varies
among
users
. This allows short codes to be more
r esistant to brute force attacks;

since the adversary
ignores the code

and the password
, he cannot generate
asymmetric keys where

the
keyedhash of the public key
is identical to the
verification
code.

Although multiple
options

are possible to calculate the code,
we choose the most flexible and secure

one
. Therefore,
we
calculate to code using the keyedhash of the public key of
the

root CA
. We
suggest the
use
of

HMACSHA256

which
ha s

a n output of 32

bytes. Since this is a

very
long
value for
the user
, we
suggest using

the first
6

bytes

(i.e. 48 bits)

as our
code. We believe that
this

length is long enough to
authenticate the root CA

and is
convenient
for the user. In
addition, we convert the
binary value into base64

and we
obtain a string of
8

alphanumeric characters
,
as follows:

C
ode

=

base64
(first
48
(HMACSHA256
(password,

CA_PubKey)))

(1)

We designed RCMS

to
accept any AS certificat e issued by the
trusted root CA. Therefore,

the network administrators

have
two options

when the private key of any AS is compromised.
The first option is to

provide an online revocation list which
allows

the users to check
for

revoked certificates. This i s
practical for large networks with many permanent users, such
as Eduroam.

We note that

the management of the certificate
revocation list is beyond the scope of th is

paper.

The second
option is to update the public key of the root CA.
This forces

the users

to

contact the administrator
and to request

new codes.

5.2

List of SSID/CA

RCMS maintains a list of
SSIDs and the corresponding
trusted
root CA

(i.e. list of SSID/CA)
. This list
only contains SSIDs

employing

802.1X authentication. In this list, the SSID m ust
be unique but not the root CA. This means that an SSID must
have a unique root CA, but a root CA may be associated to
multiple SSIDs. This
allows the

network administrators to use
the same root CA with different

SSIDs

having

similar
meanings
,‎such‎as‎“
University of Monastir

1
”‎and‎“
University
of
Monastir

2
”.

The list is updated in the following
cases
:

1)

First authentication with an SSID
: a new entry is added
to the list upon the successful
certificate check using

the
verification
code.

2)

The root certificat e has changed, including the public key
:
if the user has the
right

verification code,
the existing
entry
is

updated.
Otherwise
, this is a possible evil twin
attack

and the authentication must be canceled
. If this
AP
is legitimate
, the user must conta ct the

administrator to
obtain the

new
verification
code.

3)

The root certificate is modified but the public key
has

not
change d
: the root certificate is
seamlessly
updated

in the
list of SSID/CA

and no

verification

code is requested.

Table
2

depicts an example of the list of
SSID/
CA. It includes
the columns SSID, the public key of the trusted CA, the root
certificate fingerprint (allows the detection of any update in
the certificate) and the root certificate path (the storag e path of
the certificate on the user
device
). It is possible to include
addition columns to this list in order to provide more details,
such as the date of first association, the update history, etc.

Similar

to
the operating mode of

current user devices,
it is
necessary to store the user credentials in order to provide
seamless authentication to trusted networks.

These credentials
can be saved
either
in this list

or

in a separate
encrypted
list.

Table
2
. List of SSIDs and the
corres ponding
trusted C
A (SSID/CA)

SSID

P
ublic Key

of Root
CA

Root c ertificate
fingerprint

Root c ertificate path

Univ_Monastir

a5f716e894c6…

487ed1fb85c3
…

rootca/univ_m.pem

Hotel_
SBM

3749f2ae752b…

3cb5d0e5edbd
…

rootca/
h sbm.pem

…

…

…

…

5.3
AS certificate che ck

The successful
check
of the

verification

code means that the
root CA is trustworthy. Therefore, the user can accept any
certificate issued by this CA. To perform the code verification,
the STA must receive the entire certificate chain. This chain
includ es the AS certificate, the root certificate and any
intermediate certificate.
It

is received during
the
establishment
of the
TLS tunnel
within‎the‎“Certificate”‎
field

of the TLS
message, as depicted in message 9 of Figure 3
.

Upon the
reception of the certificate chain,

RCMS check s

the validity of

this
chain by inspecting the

different

issuers
. Then,

RCMS
checks if the public key of the root
CA

exists in the list of
SSID/CA and corresponds to the current SSID. If yes, the code

verification is not required since the root CA is already trusted.
Otherwise, an intermediate code is calculated based on the
user password and the public key of the root CA, according to
Equation 1. If this
intermediate
code is identical to the
verificat ion code, the root CA is trusted and the list of
SSID/CA is updated.

Otherwise, the certificate verification
fails and the 80
2.1X authentication is cancel

<!-- PDF text truncated by scrapem max_pdf_chars. -->

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
